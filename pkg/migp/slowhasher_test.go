// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import "testing"

// BenchmarkScryptSlowHasher runs benchmark tests for the scrypt slow hasher
func BenchmarkScryptSlowHasher(b *testing.B) {
	input := []byte{32}

	slowHasher := NewScryptSlowHasher()
	for i := 0; i < b.N; i++ {
		_ = slowHasher.Hash(input)
	}
}
