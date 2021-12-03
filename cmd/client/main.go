// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

// client implements the client-side logic to interact with an OPRF service and decrypt a bucket.

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloudflare/migp-go/pkg/migp"
)

func main() {
	var targetURL, configFile, inputFilename string
	var dumpConfig, showPassword bool
	var err error

	flag.StringVar(&configFile, "config", "", "Client configuration file")
	flag.BoolVar(&dumpConfig, "dump-config", false, "Dump the client configuration to stdout and exit")
	flag.BoolVar(&showPassword, "show-password", false, "Show the password in the output")
	flag.StringVar(&inputFilename, "infile", "-", "input file of credentials to query in the format <username>:<password> ('-' for stdin)")
	flag.StringVar(&targetURL, "target", "http://localhost:8080", "target MIGP server")

	flag.Parse()

	var cfg migp.Config
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
		cfg = migp.DefaultConfig()
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

	inputFile := os.Stdin
	if inputFilename != "-" {
		if inputFile, err = os.Open(inputFilename); err != nil {
			log.Fatal(err)
		}
		defer inputFile.Close()
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		fields := strings.SplitN(scanner.Text(), ":", 2)
		if len(fields) < 2 {
			continue
		}
		username, password := fields[0], fields[1]
		if status, metadata, err := migp.Query(cfg, targetURL+"/evaluate", username, password); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {

			if !showPassword {
				password = ""
			}
			out, err := json.Marshal(struct {
				Username string `json:"username"`
				Password string `json:"password,omitempty"`
				Status   string `json:"status"`
				Metadata string `json:"metadata,omitempty"`
			}{
				Username: username,
				Password: password,
				Status:   status.String(),
				Metadata: string(metadata),
			})
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println(string(out))
		}
	}
}
