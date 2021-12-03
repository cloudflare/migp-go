// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudflare/circl/oprf"
)

// Client wraps the relevant context needed to generate MIGP requests.
type Client struct {
	version         uint16
	bucketIDBitSize int
	bucketHasher    BucketHasher
	bucketEncryptor BucketEncryptor
	slowHasher      SlowHasher
	oprfClient      *oprf.Client
	oprfSuite       oprf.SuiteID
}

// ClientRequest carries the information the server needs to perform an
// evaluation
type ClientRequest struct {
	Version      uint32 `json:"version"`
	BucketID     string `json:"bucketID"`
	BlindElement []byte `json:"blindElement"`
}

// ClientRequestContext wraps the context needed to process MIGP responses
// to produce the request (username, password) breach status and associated
// metadata (if available). Not all breach entries will have metadata.
type ClientRequestContext struct {
	client      Client
	oprfRequest *oprf.ClientRequest
}

func NewClient(cfg Config) (*Client, error) {
	var err error

	c := new(Client)
	c.version = cfg.Version
	c.bucketIDBitSize = cfg.BucketIDBitSize

	c.bucketHasher, err = NewBucketHasher(cfg.BucketHasherID)
	if err != nil {
		return nil, err
	}

	c.slowHasher, err = NewSlowHasher(cfg.SlowHasherID)
	if err != nil {
		return nil, err
	}

	c.bucketEncryptor, err = NewBucketEncryptor(cfg.BucketEncryptorID)
	if err != nil {
		return nil, err
	}

	c.oprfSuite = cfg.OPRFSuite
	c.oprfClient, err = oprf.NewClient(c.oprfSuite)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// BucketID returns the bucket ID for the given username
func (c *Client) BucketID(username string) uint32 {
	return bucketHashToID(c.bucketHasher.Hash([]byte(username)), c.bucketIDBitSize)
}

// Request generates a client request byte string and a ClientRequest struct,
// given a username and password
func (c Client) Request(username string, pw string) (ClientRequest, ClientRequestContext, error) {
	bucketIDHex := BucketIDToHex(c.BucketID(username))

	input := c.slowHasher.Hash(serializeUserPassword(username, pw))

	oprfRequest, err := c.oprfClient.Request([][]byte{input})
	if err != nil {
		return ClientRequest{}, ClientRequestContext{}, err
	}
	blindedElements := oprfRequest.BlindedElements()
	if len(blindedElements) < 1 {
		return ClientRequest{}, ClientRequestContext{}, errors.New("invalid BlindedElements response")
	}

	request := ClientRequest{
		Version:      uint32(c.version),
		BucketID:     bucketIDHex,
		BlindElement: blindedElements[0],
	}
	context := ClientRequestContext{
		client:      c,
		oprfRequest: oprfRequest,
	}

	return request, context, nil
}

// Finalize parses a response message from server, completes the computation of
// the OPRF value, determines if it is in the received bucket, and decrypts the
// associated ciphertext
func (ctx ClientRequestContext) Finalize(response ServerResponse) (BreachStatus, []byte, error) {
	if uint16(response.Version) != ctx.client.version {
		return NotInBreach, nil, errors.New("wrong version in reply")
	}

	oprfOutput, err := ctx.client.oprfClient.Finalize(ctx.oprfRequest, &oprf.Evaluation{
		Elements: []oprf.SerializedElement{response.EvaluatedElement},
	}, OprfInfo)
	if err != nil {
		return NotInBreach, nil, err
	}
	if len(oprfOutput) < 1 {
		return NotInBreach, nil, errors.New("invalid Finalize response")
	}
	secret := oprfOutput[0]

	offset := 0

	for {
		if (offset + HeaderSize) > len(response.BucketContents) {
			// Note(caw): we could return an error here, but bail out to the default case
			break
		}

		valid, flag, bodyLength, err := ctx.client.bucketEncryptor.DecryptHeader(secret, response.BucketContents[offset:])
		if err != nil {
			return NotInBreach, nil, err
		}
		offset += HeaderSize
		if offset+bodyLength > len(response.BucketContents) {
			return NotInBreach, nil, errors.New("parsing error in bucket")
		}
		if valid {
			metadata, err := ctx.client.bucketEncryptor.DecryptBody(secret, response.BucketContents[offset:offset+bodyLength])
			if err != nil {
				return NotInBreach, nil, err
			}
			return flag.ToBreachStatus(), metadata, err
		}

		// Skip to the next entry
		offset += bodyLength
	}

	return NotInBreach, nil, nil
}

// Query submits a MIGP query to the target MIGP server.
func Query(cfg Config, targetURL, username, password string) (BreachStatus, []byte, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return 0, nil, err
	}

	migpRequest, context, err := client.Request(username, password)
	if err != nil {
		return 0, nil, err
	}

	serializedRequestPayload, err := json.Marshal(migpRequest)
	if err != nil {
		return 0, nil, err
	}

	request, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(serializedRequestPayload))
	if err != nil {
		return 0, nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, nil, err
	}
	if response.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf("Request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, nil, err
	}

	var responsePayload ServerResponse
	if err := responsePayload.UnmarshalBinary(body); err != nil {
		return 0, nil, err
	}

	return context.Finalize(responsePayload)
}
