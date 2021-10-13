// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package mutator

import (
	"testing"
)

// TestRdasMutate tests that the mutator produces the expected variants
func TestRdasMutate(t *testing.T) {
	tests := []struct {
		inPw   string
		outPws []string
	}{
		{inPw: "hello", outPws: []string{"Hello", "13hello", "hello1", "777hello", "helloN"}},
		{inPw: "asdf1234asdf", outPws: []string{"zxcv1234zxcv", "asdf1234asdf1", "asdf1234a", "asdf1234as", "Asdf1234asdf", "asdf1234asdfN"}},
		{inPw: "", outPws: []string{"1", "13", "777", "N"}},
	}

	for _, test := range tests {
		mutator := NewRDasMutator()
		variants := mutator.Mutate(test.inPw, 1000)
		if len(variants) < 100 {
			t.Errorf("RDasMutator didn't give back at least 100 variants, %s got only: %v",
				test.inPw, variants)
		}
		for _, outPw := range test.outPws {
			found := false
			for _, variant := range variants {
				if outPw == variant {
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
	testPassword := "password1"
	for i := 0; i < b.N; i++ {
		_ = m.Mutate(testPassword, 100)
	}
}

// BenchmarkRdasMutator10 benchmarks the first 10 mutator rules
func BenchmarkRdasMutator10(b *testing.B) {
	m := NewRDasMutator()
	testPassword := "password1"
	for i := 0; i < b.N; i++ {
		_ = m.Mutate(testPassword, 10)
	}
}
