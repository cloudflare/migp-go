// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/migp-go/pkg/migp"
)

// TestServer spins up a MIGP server and runs a series of tests
func TestServer(t *testing.T) {

	testUsername := "username1"
	testPassword := "password1"
	testMetadata := "test metadata"

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
	if string(metadata) != "" {
		t.Fatalf("metadata: want %s, got %s", "", string(metadata))
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
	if string(metadata) != testMetadata {
		t.Fatalf("metadata: want %s, got %s", testMetadata, string(metadata))
	}
}
