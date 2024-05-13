// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"syscall"

	"github.com/docker/docker/pkg/reexec"
	"github.com/google/go-sev-guest/client"
	"github.com/ultravioletrs/cocos/agent/events"
	"golang.org/x/crypto/sha3"
)

var _ Service = (*agentService)(nil)

const (
	// ReportDataSize is the size of the report data expected by the attestation service.
	ReportDataSize     = 64
	socketPath         = "unix_socket"
	algoFilePermission = 0o700
	newRoot            = "/cocos/newRoot"
	oldRoot            = "./oldRoot"
	hostName           = "cocos-algo"
	filePermission     = 0x755
)

var (
	// ErrMalformedEntity indicates malformed entity specification (e.g.
	// invalid username or password).
	ErrMalformedEntity = errors.New("malformed entity specification")
	// ErrUnauthorizedAccess indicates missing or invalid credentials provided
	// when accessing a protected resource.
	ErrUnauthorizedAccess = errors.New("missing or invalid credentials provided")
	// errUndeclaredAlgorithm indicates algorithm was not declared in computation manifest.
	errUndeclaredAlgorithm = errors.New("algorithm not declared in computation manifest")
	// errUndeclaredAlgorithm indicates algorithm was not declared in computation manifest.
	errUndeclaredDataset = errors.New("dataset not declared in computation manifest")
	// errProviderMissmatch algorithm/dataset provider does not match computation manifest.
	errProviderMissmatch = errors.New("provider does not match declaration on manifest")
	// errAllManifestItemsReceived indicates no new computation manifest items expected.
	errAllManifestItemsReceived = errors.New("all expected manifest Items have been received")
	// errUndeclaredConsumer indicates the consumer requesting results in not declared in computation manifest.
	errUndeclaredConsumer = errors.New("result consumer is undeclared in computation manifest")
	// errResultsNotReady indicates the computation results are not ready.
	errResultsNotReady = errors.New("computation results are not yet ready")
	// errStateNotReady agent received a request in the wrong state.
	errStateNotReady = errors.New("agent not expecting this operation in the current state")
	// errHashMismatch provided algorithm/dataset does not match hash in manifest.
	errHashMismatch = errors.New("malformed data, hash does not match manifest")
)

// Service specifies an API that must be fullfiled by the domain service
// implementation, and all of its decorators (e.g. logging & metrics).
type Service interface {
	Algo(ctx context.Context, algorithm Algorithm) error
	Data(ctx context.Context, dataset Dataset) error
	Result(ctx context.Context, consumer string) ([]byte, error)
	Attestation(ctx context.Context, reportData [ReportDataSize]byte) ([]byte, error)
}

type agentService struct {
	computation Computation    // Holds the current computation request details.
	algorithm   []byte         // Stores the algorithm received for the computation.
	datasets    [][]byte       // Stores the datasets received for the computation.
	result      []byte         // Stores the result of the computation.
	sm          *StateMachine  // Manages the state transitions of the agent service.
	runError    error          // Stores any error encountered during the computation run.
	eventSvc    events.Service // Service for publishing events related to computation.
}

var _ Service = (*agentService)(nil)

// New instantiates the agent service implementation.
func New(ctx context.Context, logger *slog.Logger, eventSvc events.Service, cmp Computation) Service {
	svc := &agentService{
		sm:       NewStateMachine(logger),
		eventSvc: eventSvc,
	}

	go svc.sm.Start(ctx)
	svc.sm.SendEvent(start)
	svc.sm.StateFunctions[idle] = svc.publishEvent("in-progress", json.RawMessage{})
	svc.sm.StateFunctions[receivingManifest] = svc.publishEvent("in-progress", json.RawMessage{})
	svc.sm.StateFunctions[receivingAlgorithm] = svc.publishEvent("in-progress", json.RawMessage{})
	svc.sm.StateFunctions[receivingData] = svc.publishEvent("in-progress", json.RawMessage{})
	svc.sm.StateFunctions[resultsReady] = svc.publishEvent("in-progress", json.RawMessage{})
	svc.sm.StateFunctions[complete] = svc.publishEvent("in-progress", json.RawMessage{})
	svc.sm.StateFunctions[running] = svc.runComputation

	svc.computation = cmp
	svc.sm.SendEvent(manifestReceived)
	return svc
}

func (as *agentService) Algo(ctx context.Context, algorithm Algorithm) error {
	if as.sm.GetState() != receivingAlgorithm {
		return errStateNotReady
	}
	if as.algorithm != nil {
		return errAllManifestItemsReceived
	}

	hash := sha3.Sum256(algorithm.Algorithm)

	if as.computation.Algorithm.ID != algorithm.ID {
		return errUndeclaredAlgorithm
	}

	if as.computation.Algorithm.Provider != algorithm.Provider {
		return errProviderMissmatch
	}

	if hash != as.computation.Algorithm.Hash {
		return errHashMismatch
	}

	as.algorithm = algorithm.Algorithm

	if as.algorithm != nil {
		as.sm.SendEvent(algorithmReceived)
	}

	return nil
}

func (as *agentService) Data(ctx context.Context, dataset Dataset) error {
	if as.sm.GetState() != receivingData {
		return errStateNotReady
	}
	if len(as.computation.Datasets) == 0 {
		return errAllManifestItemsReceived
	}

	hash := sha3.Sum256(dataset.Dataset)

	index := containsID(as.computation.Datasets, dataset.ID)
	switch index {
	case -1:
		return errUndeclaredDataset
	default:
		if as.computation.Datasets[index].Provider != dataset.Provider {
			return errProviderMissmatch
		}
		if hash != as.computation.Datasets[index].Hash {
			return errHashMismatch
		}
		as.computation.Datasets = slices.Delete(as.computation.Datasets, index, index+1)
	}

	as.datasets = append(as.datasets, dataset.Dataset)

	if len(as.computation.Datasets) == 0 {
		as.sm.SendEvent(dataReceived)
	}

	return nil
}

func (as *agentService) Result(ctx context.Context, consumer string) ([]byte, error) {
	if as.sm.GetState() != resultsReady {
		return []byte{}, errResultsNotReady
	}
	if len(as.computation.ResultConsumers) == 0 {
		return []byte{}, errAllManifestItemsReceived
	}
	index := slices.Index(as.computation.ResultConsumers, consumer)
	switch index {
	case -1:
		return []byte{}, errUndeclaredConsumer
	default:
		as.computation.ResultConsumers = slices.Delete(as.computation.ResultConsumers, index, index+1)
	}

	if len(as.computation.ResultConsumers) == 0 {
		as.sm.SendEvent(resultsConsumed)
	}
	// Return the result file or an error
	return as.result, as.runError
}

func (as *agentService) Attestation(ctx context.Context, reportData [ReportDataSize]byte) ([]byte, error) {
	provider, err := client.GetQuoteProvider()
	if err != nil {
		return []byte{}, err
	}
	rawQuote, err := provider.GetRawQuote(reportData)
	if err != nil {
		return []byte{}, err
	}

	return rawQuote, nil
}

func (as *agentService) runComputation() {
	as.publishEvent("starting", json.RawMessage{})()
	as.sm.logger.Debug("computation run started")
	defer as.sm.SendEvent(runComplete)
	as.publishEvent("in-progress", json.RawMessage{})()
	result, err := run(as.algorithm, as.datasets[0])
	if err != nil {
		as.runError = err
		as.sm.logger.Warn(fmt.Sprintf("computation failed with error: %s", err.Error()))
		as.publishEvent("failed", json.RawMessage{})()
		return
	}
	as.publishEvent("complete", json.RawMessage{})()
	as.result = result
}

func (as *agentService) publishEvent(status string, details json.RawMessage) func() {
	return func() {
		if err := as.eventSvc.SendEvent(as.sm.State.String(), status, details); err != nil {
			as.sm.logger.Warn(err.Error())
		}
	}
}

func run(algoContent, dataContent []byte) ([]byte, error) {
	var resultBuffer bytes.Buffer
	var errorBuffer bytes.Buffer

	if err := os.MkdirAll(newRoot, fs.FileMode(filePermission)); err != nil {
		return nil, fmt.Errorf("error could not create new root dir: %v", err)
	}
	defer os.RemoveAll(newRoot)

	f, err := os.Create(path.Join(newRoot, "algorithm"))
	if err != nil {
		return nil, fmt.Errorf("error creating algorithm file: %v", err)
	}
	defer os.Remove(f.Name())

	if _, err := f.Write(algoContent); err != nil {
		return nil, fmt.Errorf("error writing algorithm to file: %v", err)
	}

	if err := os.Chmod(f.Name(), algoFilePermission); err != nil {
		return nil, fmt.Errorf("error changing file permissions: %v", err)
	}

	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("error closing file: %v", err)
	}

	// Construct the executable with CSV data as a command-line argument
	data := string(dataContent)
	reexec.Register("namespaceInit", namespaceInit)
	if reexec.Init() {
		return nil, fmt.Errorf("error while initializing namespace: %v", err)
	}

	cmd := reexec.Command("namespaceInit", "algorithm", data)
	cmd.Stdout = &resultBuffer
	cmd.Stderr = &errorBuffer
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("error starting reexec: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("reexec execution error: %v", err)
	}

	if errorBuffer.Len() > 0 {
		return nil, fmt.Errorf("error occurred during algorithm run: %v", err)
	}

	return resultBuffer.Bytes(), nil
}

func namespaceInit() {
	// Mount /proc
	newProc := filepath.Join(newRoot, "/proc")
	if err := os.MkdirAll(newProc, fs.FileMode(filePermission)); err != nil {
		os.Exit(1)
	}

	if err := syscall.Mount("proc", newProc, "proc", 0, ""); err != nil {
		os.Exit(1)
	}

	if err := pivotRoot(newRoot); err != nil {
		os.Exit(1)
	}

	if err := syscall.Sethostname([]byte(hostName)); err != nil {
		os.Exit(1)
	}

	if err := namespaceRun(); err != nil {
		os.Exit(1)
	}
}

func pivotRoot(newRootPath string) error {
	oldRootPath := filepath.Join(newRootPath, oldRoot)

	if err := syscall.Mount(newRootPath, newRootPath, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}

	if err := os.MkdirAll(oldRootPath, algoFilePermission); err != nil {
		return err
	}

	if err := syscall.PivotRoot(newRootPath, oldRootPath); err != nil {
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	if err := syscall.Unmount(oldRoot, syscall.MNT_DETACH); err != nil {
		return err
	}

	if err := os.RemoveAll(oldRoot); err != nil {
		return err
	}

	return nil
}

func namespaceRun() error {
	cmd := exec.Command(os.Args[1], os.Args[2])

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting algorithm: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("algorithm execution error: %v", err)
	}

	return nil
}
