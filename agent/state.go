// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package agent

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

//go:generate stringer -type=State
type State uint8

const (
	Idle State = iota
	ReceivingManifest
	ReceivingAlgorithm
	ReceivingData
	Running
	ConsumingResults
	Complete
	Failed
	AlgorithmRun
)

//go:generate stringer -type=Status
type Status uint8

const (
	IdleState Status = iota
	InProgress
	Ready
	Completed
	Terminated
	Warning
)

type event uint8

const (
	start event = iota
	manifestReceived
	algorithmReceived
	dataReceived
	runComplete
	resultsConsumed
	runFailed
)

// StateMachine represents the state machine.
type StateMachine struct {
	mu             sync.Mutex
	State          State
	EventChan      chan event
	Transitions    map[State]map[event]State
	StateFunctions map[State]func()
	logger         *slog.Logger
	wg             *sync.WaitGroup
}

// NewStateMachine creates a new StateMachine.
func NewStateMachine(logger *slog.Logger, cmp Computation) *StateMachine {
	sm := &StateMachine{
		State:          Idle,
		EventChan:      make(chan event),
		Transitions:    make(map[State]map[event]State),
		StateFunctions: make(map[State]func()),
		logger:         logger,
		wg:             &sync.WaitGroup{},
	}

	sm.Transitions[Idle] = make(map[event]State)
	sm.Transitions[Idle][start] = ReceivingManifest

	sm.Transitions[ReceivingManifest] = make(map[event]State)
	sm.Transitions[ReceivingManifest][manifestReceived] = ReceivingAlgorithm

	sm.Transitions[ReceivingAlgorithm] = make(map[event]State)
	switch len(cmp.Datasets) {
	case 0:
		sm.Transitions[ReceivingAlgorithm][algorithmReceived] = Running
	default:
		sm.Transitions[ReceivingAlgorithm][algorithmReceived] = ReceivingData
	}

	sm.Transitions[ReceivingData] = make(map[event]State)
	sm.Transitions[ReceivingData][dataReceived] = Running

	sm.Transitions[Running] = make(map[event]State)
	sm.Transitions[Running][runComplete] = ConsumingResults
	sm.Transitions[Running][runFailed] = Failed

	sm.Transitions[ConsumingResults] = make(map[event]State)
	sm.Transitions[ConsumingResults][resultsConsumed] = Complete

	return sm
}

// Start the state machine.
func (sm *StateMachine) Start(ctx context.Context) {
	sm.wg.Add(1)
	defer sm.wg.Done()
	for {
		select {
		case event := <-sm.EventChan:
			currentState := sm.GetState()
			var nextState State
			var stateFunc func()
			var valid bool

			sm.mu.Lock()
			nextState, valid = sm.Transitions[sm.State][event]
			if valid {
				sm.State = nextState
				stateFunc = sm.StateFunctions[nextState]
			}
			sm.mu.Unlock()

			if valid {
				sm.logger.Debug(fmt.Sprintf("Transition: %v -> %v\n", currentState, nextState))
				if stateFunc != nil {
					go stateFunc()
				}
			} else {
				sm.logger.Error(fmt.Sprintf("Invalid transition: %v -> ???\n", sm.State))
			}

		case <-ctx.Done():
			return
		}
	}
}

// SendEvent sends an event to the state machine.
func (sm *StateMachine) SendEvent(event event) {
	sm.EventChan <- event
}

func (sm *StateMachine) GetState() State {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.State
}

func (sm *StateMachine) SetState(state State) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.State = state
}
