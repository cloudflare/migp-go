// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import "testing"

// BenchmarkSHA256BucketHasher runs benchmark tests for the bucket hasher
func BenchmarkSHA256BucketHasher(b *testing.B) {
	input := []byte{32}

	hasher := NewSHA256BucketHasher()
	for i := 0; i < b.N; i++ {
		_ = hasher.Hash(input)
	}
}
