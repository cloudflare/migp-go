// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/cloudflare/circl/oprf"
)

const (
	// DefaultMIGPVersion gives the version of the MIGP library and
	// parameter set.  Compatibility across versions is not guaranteed.
	DefaultMIGPVersion = 1

	// DefaultBucketIDBitSize is the number of high-order bits of the
	// bucket hash to use for the bucket identifier. The max size of this
	// field is 32 to allow the bucket identifier to be stored as a uint32.
	DefaultBucketIDBitSize = 20

	// Default cryptographic parameters for this version of MIGP
	DefaultBucketHasher    = BucketHasherSHA256
	DefaultSlowHasher      = SlowHasherScrypt
	DefaultBucketEncryptor = BucketEncryptorHKDFSHA256
	DefaultOPRFSuite       = uint16(oprf.OPRFP256)

	// CtxtKeyCheckSize is the size of key check string in bytes. We use this
	// to check if a given bucket entry header matches the derived key.
	CtxtKeyCheckSize = 20

	// HeaderSize is the size of a MIGP entry header in bytes. The header
	// consists of the key check bytes, 1-byte flag, and 4-byte body
	// length.
	HeaderSize = CtxtKeyCheckSize + 5
)

// Config contains MIGP configuration used both clients and servers.
type Config struct {
	Version           uint16       `json:"version"`
	BucketIDBitSize   int          `json:"bucketIDBitSize"`
	BucketHasherID    uint16       `json:"bucketHasher"`
	SlowHasherID      uint16       `json:"slowHasher"`
	BucketEncryptorID uint16       `json:"bucketEncryptor"`
	OPRFSuite         oprf.SuiteID `json:"oprfSuite"`
}

// DefaultConfig returns a new default configuration
func DefaultConfig() Config {
	return Config{
		Version:           DefaultMIGPVersion,
		BucketHasherID:    DefaultBucketHasher,
		BucketEncryptorID: DefaultBucketEncryptor,
		SlowHasherID:      DefaultSlowHasher,
		OPRFSuite:         DefaultOPRFSuite,
		BucketIDBitSize:   DefaultBucketIDBitSize,
	}
}

// Flag represents the type of metadata for a breach item.
type MetadataType uint8

const (
	// Dummy means the metadata was dummy data (used for length-hiding purposes)
	MetadataDummy MetadataType = iota
	// MetadataBreachedPassword means the (username, password) tuple corresponds to a breached password
	MetadataBreachedPassword
	// MetadataSimilarPassword means the (username, password) tuple is similar to a breached password
	MetadataSimilarPassword
	// MetadataBreachedUsername means the username has at least one breached password
	MetadataBreachedUsername
)

// String returns a string representation of a metadata type
func (mt MetadataType) String() string {
	switch mt {
	case MetadataDummy:
		return "dummy metadata"
	case MetadataBreachedPassword:
		return "breached password"
	case MetadataSimilarPassword:
		return "similar password"
	case MetadataBreachedUsername:
		return "breached username"
	default:
		return "Unknown metadata type"
	}
}

// Valid checks if the metadata type is recognized by the library
func (mt MetadataType) Valid() bool {
	switch mt {
	case MetadataDummy, MetadataBreachedPassword, MetadataSimilarPassword, MetadataBreachedUsername:
		return true
	}
	return false
}

// ToBreachStatus converts a metadata type to a breach status
func (mt MetadataType) ToBreachStatus() BreachStatus {
	switch mt {
	case MetadataBreachedPassword:
		return InBreach
	case MetadataSimilarPassword:
		return SimilarInBreach
	case MetadataBreachedUsername:
		return UsernameInBreach
	default:
		return NotInBreach
	}
}

// BreachStatus indicates the status of (username, password) tuple with respect
// to known breaches, e.g., whether or not the pair exists in a known breach,
// a similar password exists in a known breach, or it's not in a breach at all.
type BreachStatus uint8

const (
	// NotInBreach indicates the target tuple was not in a known breach.
	NotInBreach BreachStatus = iota
	// InBreach indicates the target tuple was in a known breach.
	InBreach
	// SimilarInBreach indicates that a pair with a similar password to
	// the target tuple was in a known breach.
	SimilarInBreach
	// UsernameInBreach indicates the target username has at least one
	// associated password in a known breach.
	UsernameInBreach
)

// String returns a string representation of a breach status
func (bs BreachStatus) String() string {
	switch bs {
	case NotInBreach:
		return "password not in breach"
	case InBreach:
		return "password in breach"
	case SimilarInBreach:
		return "similar password in breach"
	case UsernameInBreach:
		return "username in breach"
	default:
		return "unknown breach status"
	}
}

// encodeString encodes input `data` using a two-byte big-endian length prefix.
func encodeString(data string) []byte {
	if len(data) > (1 << 16) {
		panic("Length overflow")
	}
	lengthBuffer := make([]byte, 2)
	bytes := []byte(data)
	binary.BigEndian.PutUint16(lengthBuffer, uint16(len(bytes)))
	return append(lengthBuffer, bytes...)
}

// serializeUserPassword generates a byte string consisting of username and
// password.  We use a simple prefix-free length-based encoding of the username
// and password, where lengths are encoded as 16-bit big-endian unsigned
// integers.  Note that metadata is not included in this serialization.
func serializeUserPassword(username string, password string) []byte {
	usernameBuffer := encodeString(username)
	passwordBuffer := encodeString(password)
	return append(usernameBuffer, passwordBuffer...)
}

// bucketHashToID returns a uint32 bucket ID given a bucket hash and the bucket
// ID bit size
func bucketHashToID(bucketHash []byte, bitSize int) uint32 {
	if bitSize > 32 {
		panic("Bucket ID bit size cannot be greater than 32")
	}
	bucketID := binary.BigEndian.Uint32(bucketHash)
	// shift to remove extraneous low-order bits
	return bucketID >> (32 - bitSize) & ((1 << bitSize) - 1)
}

// BucketIDToHex encodes a uint32 bucket ID to a hex string
func BucketIDToHex(bucketID uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b[0:], bucketID)
	return hex.EncodeToString(b)
}
