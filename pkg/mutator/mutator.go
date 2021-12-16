// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package mutator

// Mutator interface gives a Mutate() function that takes as input a requested
// number of mutations, and returns a set of mutated passwords
type Mutator interface {
	Mutate([]byte, int) [][]byte
}
