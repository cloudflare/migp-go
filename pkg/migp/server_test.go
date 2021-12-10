// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/cloudflare/circl/oprf"
)

// TestSerialization tests that a server configuration can be correctly
// serialized and deserialized
func TestSerialization(t *testing.T) {

	testUsername := []byte("username1")
	testPassword := []byte("password1")
	testMetadata := []byte("test metadata")

	// initialize server with fresh random key
	server, err := NewServer(DefaultServerConfig())
	if err != nil {
		t.Fatal(err)
	}

	buf, err := json.Marshal(server.Config())
	if err != nil {
		t.Error(err)
	}

	var cfg ServerConfig
	err = json.Unmarshal(buf, &cfg)
	if err != nil {
		t.Error(err)
	}
	server2, err := NewServer(cfg)
	if err != nil {
		t.Error(err)
	}
	if server.version != server2.version {
		t.Error("serialization failed: version not equal")
	}

	ciphertextA, err := server.EncryptBucketEntry(testUsername, testPassword, MetadataSimilarPassword, testMetadata)
	if err != nil {
		t.Error(err)
	}

	ciphertextB, err := server2.EncryptBucketEntry(testUsername, testPassword, MetadataSimilarPassword, testMetadata)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(ciphertextA, ciphertextB) {
		t.Error("serialization test failed")
	}
}

// TestServerResponseSerialization tests the serialization of a MIGP server response
func TestServerResponseSerialization(t *testing.T) {
	sizes, err := oprf.GetSizes(DefaultOPRFSuite)
	if err != nil {
		t.Fatal(err)
	}
	r1 := ServerResponse{
		123,
		make([]byte, sizes.SerializedElementLength),
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	if _, err := rand.Read(r1.EvaluatedElement); err != nil {
		t.Fatal(err)
	}
	data, err := r1.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	expectedLen := 4 + len(r1.EvaluatedElement) + len(r1.BucketContents)
	if len(data) != expectedLen {
		t.Fatalf("want %d, got %d", expectedLen, len(data))
	}

	r2 := ServerResponse{}

	if err = r2.UnmarshalBinary(data); err != nil {
		t.Fatal(err)
	}

	if r1.Version != r2.Version || !bytes.Equal(r1.EvaluatedElement, r2.EvaluatedElement) || !bytes.Equal(r1.BucketContents, r2.BucketContents) {
		t.Fatal("mismatch")
	}
}
