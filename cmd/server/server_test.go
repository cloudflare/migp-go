// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/migp-go/pkg/migp"
)

// TestServer spins up a MIGP server and runs a series of tests
func TestServer(t *testing.T) {

	testUsername := []byte("username1")
	testPassword := []byte("password1")
	testMetadata := []byte("test metadata")

	s, err := newServer(migp.DefaultServerConfig())
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(s.handler())
	defer httpServer.Close()

	cfg := migp.DefaultConfig()

	// query test record before insertion
	status, metadata, err := migp.Query(cfg, httpServer.URL+"/evaluate", testUsername, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	if status != migp.NotInBreach {
		t.Fatalf("status: want %s, got %s", migp.NotInBreach, status)
	}
	if len(metadata) != 0 {
		t.Fatalf("metadata: want %d, got %d", 0, len(metadata))
	}

	// insert test record
	err = s.insert(testUsername, testPassword, testMetadata, 9, true)
	if err != nil {
		t.Fatal(err)
	}

	// query test record after insertion
	status, metadata, err = migp.Query(cfg, httpServer.URL+"/evaluate", testUsername, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	if status != migp.InBreach {
		t.Fatalf("status: want %s, got %s", migp.InBreach, status)
	}
	if !bytes.Equal(metadata, testMetadata) {
		t.Fatalf("metadata: want %s, got %s", testMetadata, string(metadata))
	}
}
