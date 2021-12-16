// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package mutator

import (
	"bytes"
	"testing"
)

// TestRdasMutate tests that the mutator produces the expected variants
func TestRdasMutate(t *testing.T) {
	tests := []struct {
		inPw   []byte
		outPws [][]byte
	}{
		{inPw: []byte("hello"), outPws: [][]byte{[]byte("Hello"), []byte("13hello"), []byte("hello1"), []byte("777hello"), []byte("helloN")}},
		{inPw: []byte("asdf1234asdf"), outPws: [][]byte{[]byte("zxcv1234zxcv"), []byte("asdf1234asdf1"), []byte("asdf1234a"), []byte("asdf1234as"), []byte("Asdf1234asdf"), []byte("asdf1234asdfN")}},
		{inPw: nil, outPws: [][]byte{[]byte("1"), []byte("13"), []byte("777"), []byte("N")}},
	}

	for _, test := range tests {
		mutator := NewRDasMutator()
		variants := mutator.Mutate(test.inPw, 1000)
		if len(variants) < 100 {
			t.Errorf("RDasMutator didn't give back at least 100 variants, %s got only: %v",
				test.inPw, len(variants))
		}
		for _, outPw := range test.outPws {
			found := false
			for _, variant := range variants {
				if bytes.Equal(outPw, variant) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("RDasMutator didn't give back for %s the expected mutation %s",
					test.inPw, outPw)
			}
		}
	}
}

// BenchmarkRdasMutator100 benchmarks the first 100 mutator rules
func BenchmarkRdasMutator100(b *testing.B) {
	m := NewRDasMutator()
	testPassword := []byte("password1")
	for i := 0; i < b.N; i++ {
		_ = m.Mutate(testPassword, 100)
	}
}

// BenchmarkRdasMutator10 benchmarks the first 10 mutator rules
func BenchmarkRdasMutator10(b *testing.B) {
	m := NewRDasMutator()
	testPassword := []byte("password1")
	for i := 0; i < b.N; i++ {
		_ = m.Mutate(testPassword, 10)
	}
}
