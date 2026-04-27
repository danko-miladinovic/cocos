// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	tdxrtmr "github.com/google/go-tdx-guest/rtmr"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, tdxrtmr.ExtendDigest))
}

func run(args []string, stdout, stderr io.Writer, extendDigest func(int, []byte) error) int {
	flags := flag.NewFlagSet("tdx-rtmr-extend", flag.ContinueOnError)
	flags.SetOutput(stderr)

	rtmrIndex := flags.Int("rtmr", 3, "TDX RTMR index to extend; userspace may extend RTMR2 or RTMR3")
	sha384Hex := flags.String("sha384", "", "hex-encoded 48-byte SHA-384 digest")

	if err := flags.Parse(args); err != nil {
		return 2
	}

	if *sha384Hex == "" {
		fmt.Fprintln(stderr, "missing required -sha384 digest")
		return 1
	}

	if *rtmrIndex != 2 && *rtmrIndex != 3 {
		fmt.Fprintf(stderr, "invalid RTMR index %d: userspace can extend only RTMR2 or RTMR3\n", *rtmrIndex)
		return 1
	}

	digest, err := hex.DecodeString(*sha384Hex)
	if err != nil {
		fmt.Fprintf(stderr, "invalid SHA-384 digest: %v\n", err)
		return 1
	}

	if len(digest) != 48 {
		fmt.Fprintf(stderr, "invalid SHA-384 digest length %d: expected 48 bytes\n", len(digest))
		return 1
	}

	if err := extendDigest(*rtmrIndex, digest); err != nil {
		fmt.Fprintf(stderr, "failed to extend RTMR%d: %v\n", *rtmrIndex, err)
		return 1
	}

	fmt.Fprintf(stdout, "extended RTMR%d\n", *rtmrIndex)
	return 0
}
