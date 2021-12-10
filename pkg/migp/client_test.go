// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"bytes"
	"testing"
)

// KVMock is a simple KV store implementation
type KVMock struct {
	store map[string][]byte
}

// Get returns the value associated with the requested key
func (kv *KVMock) Get(key string) ([]byte, error) {
	return kv.store[key], nil
}

// TestQuery spins up a MIGP server and runs a series of client requests
// against it
func TestQuery(t *testing.T) {

	testCases := []struct {
		Username     []byte
		Password     []byte
		MetadataFlag MetadataType
		Metadata     []byte
		OutputValue  BreachStatus
	}{
		{[]byte("test@mail.com"), []byte("password1234"), MetadataBreachedPassword, []byte("my favorite breach"), InBreach},
		{[]byte("TESTt@mail.com"), []byte("password1234"), MetadataBreachedPassword, []byte("my favorite breach"), InBreach},
		{[]byte("test@mail.com"), []byte("password3214"), MetadataBreachedPassword, []byte("my favorite breach, flipped"), InBreach},
		{[]byte{}, []byte{}, MetadataSimilarPassword, []byte("my favorite breach"), SimilarInBreach},
		{[]byte("!!#%test"), []byte("Password1231"), MetadataSimilarPassword, nil, SimilarInBreach},
		{[]byte("test2"), []byte("!!&*F(DSbjklzd"), MetadataBreachedPassword, []byte("$#!!%BVAF"), InBreach},
		{[]byte("username"), []byte("password5678"), MetadataBreachedUsername, []byte("my favorite breach"), UsernameInBreach},
	}

	// initialize server with fresh random key
	server, err := NewServer(DefaultServerConfig())
	if err != nil {
		t.Fatal(err)
	}

	kv := new(KVMock)
	kv.store = make(map[string][]byte)
	for _, test := range testCases {
		bucketIDHex := BucketIDToHex(server.BucketID(test.Username))

		newEntry, err := server.EncryptBucketEntry(test.Username, test.Password, test.MetadataFlag, test.Metadata)
		if err != nil {
			t.Error(err)
		}

		kv.store[bucketIDHex] = append(kv.store[bucketIDHex], newEntry...)
	}

	client, err := NewClient(DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	// test for inserted entries
	for _, test := range testCases {

		request, clientFinalize, err := client.Request(test.Username, test.Password)
		if err != nil {
			t.Error(err)
		}

		response, err := server.HandleRequest(request, kv)
		if err != nil {
			t.Error(err)
		}

		result, mdString, err := clientFinalize.Finalize(response)
		if err != nil {
			t.Error(err)
		}

		// Compare with encEntry
		if result != test.OutputValue || !bytes.Equal(mdString, test.Metadata) {
			t.Errorf("result for %s incorrect. Got %d '%s' (expected: %d '%s')",
				test.Password, result, mdString, test.OutputValue, test.Metadata)
		}
	}

	// test for an uninserted entry
	username, password := []byte("username"), []byte("password")
	request, clientFinalize, err := client.Request(username, password)
	if err != nil {
		t.Error(err)
	}

	response, err := server.HandleRequest(request, kv)
	if err != nil {
		t.Error(err)
	}

	result, mdString, err := clientFinalize.Finalize(response)
	if err != nil {
		t.Error(err)
	}

	// Compare with encEntry
	if result != NotInBreach || mdString != nil {
		t.Errorf("result for %s incorrect. Got %d (expected: %d)",
			password, result, NotInBreach)
	}
}
