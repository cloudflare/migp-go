// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

// server implements a MIGP server. It supports encrypting and uploading a
// database of breach entries to buckets, and serving those buckets to clients
// via the MIGP protocol.

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"code.cfops.it/crypto/migp/pkg/migp"
)

func main() {

	var configFile, inputFilename, metadata, listenAddr string
	var dumpConfig, includeUsernameVariant bool
	var numVariants int

	flag.StringVar(&configFile, "config", "", "Server configuration file")
	flag.StringVar(&listenAddr, "listen", "localhost:8080", "Server listen address")
	flag.BoolVar(&dumpConfig, "dump-config", false, "Dump the server configuration to stdout and exit")
	flag.StringVar(&inputFilename, "infile", "-", "input file of credentials to insert in the format <username>:<password> ('-' for stdin)")
	flag.StringVar(&metadata, "metadata", "", "optional metadata string to store alongside breach entries")
	flag.IntVar(&numVariants, "num-variants", 9, "number of password variants to include")
	flag.BoolVar(&includeUsernameVariant, "username-variant", true, "include a username-only variant")

	flag.Parse()

	var cfg migp.ServerConfig
	if configFile != "" {
		data, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(data, &cfg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cfg = migp.DefaultServerConfig()
	}

	if dumpConfig {
		data, err := json.Marshal(&cfg)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Stdout.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	s, err := newServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	inputFile := os.Stdin
	if inputFilename != "-" {
		if inputFile, err = os.Open(inputFilename); err != nil {
			log.Fatal(err)
		}
		defer inputFile.Close()
	}

	successCount, failureCount := 0, 0
	log.Printf("Encrypting breach entries: %d successes, %d failures", successCount, failureCount)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		fields := strings.SplitN(scanner.Text(), ":", 2)
		if len(fields) < 2 {
			failureCount += 1
			continue
		}
		username, password := fields[0], fields[1]
		if err := s.insert(username, password, metadata, numVariants, includeUsernameVariant); err != nil {
			failureCount += 1
			continue
		}
		successCount += 1
		log.Printf("\rEncrypting breach entries: %d successes, %d failures", successCount, failureCount)
	}

	log.Printf("\nStarting MIGP server")
	log.Fatal(http.ListenAndServe(listenAddr, s.handler()))
}
