// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package main

import "sync"

// kvStore is a wrapper for a KV store. For now just use a simple dynamically
// allocated in-memory go map This won't scale properly, but ok for testing.
// Implements migp.Getter
type kvStore struct {
	store map[string][]byte
	lock  sync.RWMutex
}

// newKVStore initializes a new bucket store. Just using a simple map for now.
func newKVStore() (*kvStore, error) {
	return &kvStore{
		store: make(map[string][]byte),
	}, nil
}

// Put a value at key id and replace any existing value.
func (kv *kvStore) Put(id string, value []byte) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	kv.store[id] = value
	return nil
}

// Append a value to any existing value at key id.
func (kv *kvStore) Append(id string, value []byte) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	kv.store[id] = append(kv.store[id], value...)
	return nil
}

// Get returns the value in the key identified by id.
func (kv *kvStore) Get(id string) ([]byte, error) {
	kv.lock.RLock()
	defer kv.lock.RUnlock()
	return kv.store[id], nil
}
