// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"code.cfops.it/crypto/migp/pkg/migp"
	"code.cfops.it/crypto/migp/pkg/mutator"
)

// newServer returns a new server initialized using the provided configuration
func newServer(cfg migp.ServerConfig) (*server, error) {
	migpServer, err := migp.NewServer(cfg)
	if err != nil {
		return nil, err
	}

	kv, err := newKVStore()
	if err != nil {
		return nil, err
	}

	return &server{
		migpServer: migpServer,
		kv:         kv,
	}, nil
}

// server wraps a MIGP server and backing KV store
type server struct {
	migpServer *migp.Server
	kv         *kvStore
}

// handler handles client requests
func (s *server) handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/evaluate", s.handleEvaluate)
	return mux
}

// insert encrypts a credential pair and stores it in the configured KV store
func (s *server) insert(username, password, metadata string, numVariants int, includeUsernameVariant bool) error {

	bucketIDHex := migp.BucketIDToHex(s.migpServer.BucketID(username))
	newEntry, err := s.migpServer.EncryptBucketEntry(username, password, migp.MetadataBreachedPassword, []byte(metadata))
	if err != nil {
		return err
	}
	err = s.kv.Append(bucketIDHex, newEntry)
	if err != nil {
		return err
	}

	passwordVariants := mutator.NewRDasMutator().Mutate(password, numVariants)
	for _, variant := range passwordVariants {
		newEntry, err = s.migpServer.EncryptBucketEntry(username, variant, migp.MetadataSimilarPassword, []byte(metadata))
		if err != nil {
			return err
		}
		err = s.kv.Append(bucketIDHex, newEntry)
		if err != nil {
			return err
		}
	}

	if includeUsernameVariant {
		newEntry, err = s.migpServer.EncryptBucketEntry(username, "", migp.MetadataBreachedUsername, []byte(metadata))
		if err != nil {
			return err
		}
		err = s.kv.Append(bucketIDHex, newEntry)
		if err != nil {
			return err
		}
	}

	return nil
}

// handleIndex returns a welcome message
func (s *server) handleIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the MIGP demo server\n")
}

// handleEvaluate serves a request from a MIGP client
func (s *server) handleEvaluate(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Request body reading failed:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var request migp.ClientRequest
	if err := json.Unmarshal(body, &request); err != nil {
		log.Println("Request body unmarshal failed:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	migpResponse, err := s.migpServer.HandleRequest(request, s.kv)
	if err != nil {
		log.Println("HandleRequest failed:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")

	respBody, err := migpResponse.MarshalBinary()
	if err != nil {
		log.Println("Response serialization failed:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if _, err := w.Write(respBody); err != nil {
		log.Println("Writing response failed:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
