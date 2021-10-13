// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/binary"
	"errors"

	"golang.org/x/crypto/hkdf"
)

const (
	BucketEncryptorHKDFSHA256 uint16 = 0x0001
)

var (
	DerivePadHeaderSalt = []byte("MIGP derive pad header")
	DerivePadBodySalt   = []byte("MIGP derive pad body")
)

// BucketEncryptor is a generic interface for a bucket encryption algorithm.
type BucketEncryptor interface {
	ID() uint16
	Encrypt(secret []byte, metadataFlag MetadataType, metadata []byte) (ciphertext []byte, err error)
	DecryptHeader(secret []byte, ciphertext []byte) (keyCheck bool, flag MetadataType, bodyLength int, err error)
	DecryptBody(secret []byte, ciphertext []byte) (body []byte, err error)
}

// hkdfSHA256BucketEncryptor implements BucketEncryptor using HKDF-SHA256
type hkdfSHA256BucketEncryptor struct{}

// NewHKDFSHA256BucketEncryptor returns a new key-commiting AEAD based on HKDF-SHA256
// key derivation and XOR-based encryption
func NewHKDFSHA256BucketEncryptor() hkdfSHA256BucketEncryptor {
	return hkdfSHA256BucketEncryptor{}
}

// ID returns the hkdfSHA256BucketEncryptor identifier
func (h hkdfSHA256BucketEncryptor) ID() uint16 {
	return BucketEncryptorHKDFSHA256
}

// Encrypt encrypts the input (metadataFlag || metadata) using the input secret using
// a key-committing AEAD based on HKDF-SHA256 key derivation and XOR-based encryption
// Output format:
//   XOR(<20-byte all-zero key check> | <1-byte flag>, <headerPad>) | <4-byte body length> | XOR(<body>, <bodyPad>)
func (h hkdfSHA256BucketEncryptor) Encrypt(secret []byte, flag MetadataType, body []byte) ([]byte, error) {

	headerPad, err := derivePad(secret, DerivePadHeaderSalt, CtxtKeyCheckSize+1)
	if err != nil {
		return nil, err
	}

	header := make([]byte, CtxtKeyCheckSize+1)
	header[CtxtKeyCheckSize] = byte(flag)
	encryptedHeader := xorBytes(header, headerPad)

	bodyPad, err := derivePad(secret, DerivePadBodySalt, len(body))
	if err != nil {
		return nil, err
	}

	encryptedBody := xorBytes(body, bodyPad)

	ciphertext := make([]byte, len(encryptedHeader)+4+len(encryptedBody))
	copy(ciphertext, encryptedHeader)
	binary.BigEndian.PutUint32(ciphertext[len(encryptedHeader):], uint32(len(encryptedBody)))
	copy(ciphertext[len(encryptedHeader)+4:], encryptedBody)

	return ciphertext, nil
}

// DecryptHeader decrypts the input (key check || metadataFlag) using the input secret using
// a key-committing AEAD based on HKDF-SHA256 key derivation and XOR-based encryption
func (h hkdfSHA256BucketEncryptor) DecryptHeader(secret []byte, ciphertext []byte) (bool, MetadataType, int, error) {
	// key check bytes + 1-byte flag + 4-byte metadata length
	if len(ciphertext) < HeaderSize {
		return false, 0, 0, errors.New("ciphertext of insufficient length to parse header")
	}

	// derive header pad, which encrypts the key check and flag
	headerPad, err := derivePad(secret, DerivePadHeaderSalt, CtxtKeyCheckSize+1)
	if err != nil {
		return false, 0, 0, err
	}

	keyCheck := (subtle.ConstantTimeCompare(headerPad[:CtxtKeyCheckSize], ciphertext[:CtxtKeyCheckSize]) == 1)
	flag := MetadataType(headerPad[CtxtKeyCheckSize] ^ ciphertext[CtxtKeyCheckSize])

	// body length is in plaintext
	bodyLength := int(binary.BigEndian.Uint32(ciphertext[CtxtKeyCheckSize+1 : CtxtKeyCheckSize+5]))

	return keyCheck, flag, bodyLength, nil
}

// DecryptBody decrypts the input (optional entry metadata) using the input
// secret using a key-committing AEAD based on HKDF-SHA256 key derivation and
// XOR-based encryption
func (h hkdfSHA256BucketEncryptor) DecryptBody(secret []byte, ciphertext []byte) ([]byte, error) {

	// derive body pad
	bodyPad, err := derivePad(secret, DerivePadBodySalt, len(ciphertext))
	if err != nil {
		return nil, err
	}

	return xorBytes(ciphertext, bodyPad), nil

}

// xorBytes is a helper function that computes the XOR of two byte slices that
// must be of the same length.
func xorBytes(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		panic("bytes are of mismatched length")
	}
	b3 := make([]byte, len(b1))
	for i := range b1 {
		b3[i] = b1[i] ^ b2[i]
	}
	return b3
}

// derivePad is a helper function for a collision-resistant pseudorandom
// generator.  We currently support using HKDF-SHA256 for this.
func derivePad(secret, salt []byte, length int) ([]byte, error) {
	r := hkdf.New(sha256.New, secret, salt, nil)
	pad := make([]byte, length)
	n, err := r.Read(pad)
	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, errors.New("HKDF failed to read requested bytes")
	}
	return pad, nil
}

// NewBucketEncryptor returns a bucket encryptor given its ID
func NewBucketEncryptor(id uint16) (BucketEncryptor, error) {
	switch id {
	case BucketEncryptorHKDFSHA256:
		return NewHKDFSHA256BucketEncryptor(), nil
	default:
		return nil, errors.New("unsupported bucket encryptor")
	}
}
