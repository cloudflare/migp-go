// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"errors"

	"golang.org/x/crypto/scrypt"
)

const (
	SlowHasherNull   uint16 = 0x0000
	SlowHasherScrypt uint16 = 0x0001
)

const (
	SlowHashSalt = "MIGP slow hash"
	SlowHashLen  = 32    // scrypt number of bytes of output to request
	ScryptN      = 16384 // scrypt N
	Scryptr      = 8     // scrypt r
	Scryptp      = 1     // scrypt p
)

// SlowHasher is a generic interface for a slow (memory hard) hash algorithm
type SlowHasher interface {
	ID() uint16
	Hash([]byte) []byte
}

// scryptSlowHasher implements SlowHasher using scrypt
type scryptSlowHasher struct {
	salt string
	N    int
	r    int
	p    int
	L    int
}

// NewScryptSlowHasher returns a SlowHasher instance using Scrypt with the
// following parameters from Google's mundane:
// - N: 16384
// - r: 8
// - p: 1
// See: https://github.com/google/mundane/blob/master/src/password.rs#L68
func NewScryptSlowHasher() scryptSlowHasher {
	return scryptSlowHasher{
		salt: SlowHashSalt,
		N:    ScryptN,
		r:    Scryptr,
		p:    Scryptp,
		L:    SlowHashLen,
	}
}

// ID returns the identifier of this particular hash function
func (h scryptSlowHasher) ID() uint16 {
	return SlowHasherScrypt
}

// Hash applies scrypt, with the corresponding parameters, to the input buf
func (h scryptSlowHasher) Hash(buf []byte) []byte {
	temp, err := scrypt.Key(buf, []byte(h.salt), h.N, h.r, h.p, h.L)
	if err != nil {
		// This should not happen since we use valid parameters by default.
		// Invalid parameters are therefore due to programmer error, and
		// should yield panics.
		panic(err)
	}
	return temp[:]
}

// nullSlowHasher implements SlowHasher using a no-op
type nullSlowHasher struct{}

// NewNullSlowHasher returns a no-op implementation of the SlowHasher interface
func NewNullSlowHasher() nullSlowHasher {
	return nullSlowHasher{}
}

// ID returns the identifier of this particular hash function
func (h nullSlowHasher) ID() uint16 {
	return SlowHasherNull
}

// Hash is a no-op, returning the input buf unmodified
func (h nullSlowHasher) Hash(buf []byte) []byte {
	return buf
}

// NewHasher returns an slow hasher given its ID
func NewSlowHasher(id uint16) (SlowHasher, error) {
	switch id {
	case SlowHasherNull:
		return NewNullSlowHasher(), nil
	case SlowHasherScrypt:
		return NewScryptSlowHasher(), nil
	default:
		return nil, errors.New("Unsupported slow hasher")
	}
}
