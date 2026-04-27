// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var gotIndex int
	var gotDigest []byte
	exitCode := run([]string{"-rtmr", "2", "-sha384", strings.Repeat("ab", 48)}, &stdout, &stderr, func(index int, digest []byte) error {
		gotIndex = index
		gotDigest = append([]byte(nil), digest...)
		return nil
	})

	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
	if gotIndex != 2 {
		t.Fatalf("expected RTMR index 2, got %d", gotIndex)
	}
	if len(gotDigest) != 48 {
		t.Fatalf("expected 48-byte digest, got %d bytes", len(gotDigest))
	}
	if stdout.String() != "extended RTMR2\n" {
		t.Fatalf("expected success output, got %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no stderr output, got %q", stderr.String())
	}
}

func TestRunValidationFailures(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantMsg string
	}{
		{
			name:    "missing digest",
			args:    []string{"-rtmr", "2"},
			wantMsg: "missing required -sha384 digest",
		},
		{
			name:    "invalid rtmr index",
			args:    []string{"-rtmr", "1", "-sha384", strings.Repeat("ab", 48)},
			wantMsg: "invalid RTMR index 1",
		},
		{
			name:    "invalid digest hex",
			args:    []string{"-sha384", "zz"},
			wantMsg: "invalid SHA-384 digest",
		},
		{
			name:    "invalid digest length",
			args:    []string{"-sha384", "ab"},
			wantMsg: "invalid SHA-384 digest length 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			wasCalled := false

			exitCode := run(tt.args, &stdout, &stderr, func(_ int, _ []byte) error {
				wasCalled = true
				return nil
			})

			if exitCode != 1 {
				t.Fatalf("expected exit code 1, got %d", exitCode)
			}
			if wasCalled {
				t.Fatal("extendDigest should not be called on validation failures")
			}
			if stdout.Len() != 0 {
				t.Fatalf("expected no stdout output, got %q", stdout.String())
			}
			if !strings.Contains(stderr.String(), tt.wantMsg) {
				t.Fatalf("expected stderr to contain %q, got %q", tt.wantMsg, stderr.String())
			}
		})
	}
}

func TestRunExtendDigestFailure(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"-sha384", strings.Repeat("ab", 48)}, &stdout, &stderr, func(index int, digest []byte) error {
		if index != 3 {
			t.Fatalf("expected default RTMR index 3, got %d", index)
		}
		if len(digest) != 48 {
			t.Fatalf("expected 48-byte digest, got %d bytes", len(digest))
		}
		return errors.New("boom")
	})

	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}
	if stdout.Len() != 0 {
		t.Fatalf("expected no stdout output, got %q", stdout.String())
	}
	if !strings.Contains(stderr.String(), "failed to extend RTMR3: boom") {
		t.Fatalf("expected stderr to contain failure message, got %q", stderr.String())
	}
}
