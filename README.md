# MIGP library

This contains a library for the MIGP (Might I Get Pwned) protocol. MIGP can be
used to build privacy-preserving compromised credential checking services.
Read [the paper](https://arxiv.org/pdf/2109.14490.pdf) for more details.

## Quick start

### Build

	mkdir -p bin && go build -o bin/ ./cmd/...

### Test

	go test ./...

### Generate server configuration and start MIGP server

Start a server that processes and stores breach entries from the input file.

	cat testdata/test_breach.txt | bin/server &
	
### Query MIGP server

Read entries in from the input file and query a MIGP server.  By default, the
target is set to a locally-running MIGP server, but the target flag can be used
to target production MIGP servers such as https://migp.cloudflare.com.

	cat testdata/test_queries.txt | bin/client [--target <target-server>]

## Advanced usage

Run the client and server commands with `--help` for more options, including
custom configuration support.
