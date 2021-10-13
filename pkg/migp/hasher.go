// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
)

const (
	BucketHasherSHA256 uint16 = 0x0001

	BucketHashSalt = "MIGP bucket"
)

// BucketHasher is a generic interface for a cryptographic hash algorithm
// that computes a bucket identifier
type BucketHasher interface {
	ID() uint16
	Hash([]byte) []byte
}

// sha256BucketHasher implements BucketHasher with SHA256
type sha256BucketHasher struct {
	salt string
}

// NewSHA256BucketHasher returns a BucketHasher that uses SHA256 with a fixed salt for
// computing a hash of a bucket.
func NewSHA256BucketHasher() sha256BucketHasher {
	return sha256BucketHasher{BucketHashSalt}
}

// ID returns the identifier of this hasher
func (h sha256BucketHasher) ID() uint16 {
	return BucketHasherSHA256
}

// Hash implements the Hash function for sha256BucketHasher
func (h sha256BucketHasher) Hash(buf []byte) []byte {
	if len(h.salt) > (1<<16) || len(buf) > (1<<16) {
		panic("Invalid input")
	}

	acc := make([]byte, 2+len(h.salt)+2+len(buf))
	binary.BigEndian.PutUint16(acc[:2], uint16(len(h.salt)))
	copy(acc[2:], h.salt)
	binary.BigEndian.PutUint16(acc[2+len(h.salt):], uint16(len(buf)))
	copy(acc[2+len(h.salt)+2:], buf)
	temp := sha256.Sum256(acc)
	return temp[:]
}

// NewBucketHasher returns an hasher given its ID
func NewBucketHasher(id uint16) (BucketHasher, error) {
	switch id {
	case BucketHasherSHA256:
		return NewSHA256BucketHasher(), nil
	default:
		return nil, errors.New("unsupported bucket hasher")
	}
}
