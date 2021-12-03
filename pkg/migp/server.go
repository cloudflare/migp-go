// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/cloudflare/circl/oprf"
)

// Server implements the server-side functionality of MIGP, with
// two primary functionalities: FullEvaluate, to evaluate a
// (username, password) tuple and store it in the backing database,
// and HandleRequest, to process a Client request and return the
// corresponding bucket data.
type Server struct {
	version         uint16
	bucketIDBitSize int
	bucketHasher    BucketHasher
	bucketEncryptor BucketEncryptor
	slowHasher      SlowHasher
	oprfServer      *oprf.Server
	oprfSuite       oprf.SuiteID
	privateKey      *oprf.PrivateKey
}

// ServerConfig stores all version information associated with a given server.
// ServerConfig implements the json.Marshal and json.Unmarshal interfaces.
type ServerConfig struct {
	Config
	PrivateKey *oprf.PrivateKey
}

// auxServerConfig is used for custom JSON (un)marshaling of ServerConfig
type auxServerConfig struct {
	Config
	PrivateKey []byte `json:"privateKey"`
}

// MarshalJSON serializes a server configuration to JSON
func (c *ServerConfig) MarshalJSON() ([]byte, error) {
	serializedPrivateKey, err := c.PrivateKey.Serialize()
	if err != nil {
		panic(err)
	}
	return json.Marshal(&auxServerConfig{
		Config:     c.Config,
		PrivateKey: serializedPrivateKey,
	})
}

// UnmarshalJSON deserializes a server configuration from JSON
func (c *ServerConfig) UnmarshalJSON(data []byte) error {
	var aux auxServerConfig
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Config = aux.Config
	c.PrivateKey = new(oprf.PrivateKey)
	if err := c.PrivateKey.Deserialize(aux.OPRFSuite, aux.PrivateKey); err != nil {
		return err
	}
	return nil
}

// Config returns an inspectable ServerConfig associated
// with the given server.
func (s *Server) Config() *ServerConfig {
	return &ServerConfig{
		Config: Config{
			Version:           s.version,
			BucketIDBitSize:   s.bucketIDBitSize,
			BucketHasherID:    s.bucketHasher.ID(),
			SlowHasherID:      s.slowHasher.ID(),
			BucketEncryptorID: s.bucketEncryptor.ID(),
			OPRFSuite:         s.oprfSuite,
		},
		PrivateKey: s.privateKey,
	}
}

// DefaultServerConfig generates a new default server state with a freshly keyed OPRF instance.
func DefaultServerConfig() ServerConfig {
	privateKey, err := oprf.GenerateKey(DefaultOPRFSuite, rand.Reader)
	if err != nil {
		// This will only occur in the event of developer error as we
		// supply working defaults.
		panic(err)
	}

	return ServerConfig{
		Config:     DefaultConfig(),
		PrivateKey: privateKey,
	}
}

// NewServer initializes and returns a new MIGP server from the given
// configuration
func NewServer(cfg ServerConfig) (*Server, error) {
	var err error

	s := new(Server)
	s.version = cfg.Version
	s.bucketIDBitSize = cfg.BucketIDBitSize

	s.bucketHasher, err = NewBucketHasher(cfg.BucketHasherID)
	if err != nil {
		return nil, err
	}

	s.slowHasher, err = NewSlowHasher(cfg.SlowHasherID)
	if err != nil {
		return nil, err
	}

	s.bucketEncryptor, err = NewBucketEncryptor(cfg.BucketEncryptorID)
	if err != nil {
		return nil, err
	}

	s.oprfSuite = cfg.OPRFSuite
	s.privateKey = cfg.PrivateKey

	s.oprfServer, err = oprf.NewServer(s.oprfSuite, s.privateKey)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// deriveBucketEntryKey derives a bucket entry key from a credential pair
func (s *Server) deriveBucketEntryKey(username string, password string) ([]byte, error) {
	input := s.slowHasher.Hash(serializeUserPassword(username, password))
	return s.oprfServer.FullEvaluate(input, OprfInfo)
}

// BucketID returns the bucket ID for the given username
func (s *Server) BucketID(username string) uint32 {
	return bucketHashToID(s.bucketHasher.Hash([]byte(username)), s.bucketIDBitSize)
}

// EncryptBucketEntry performs the full OPRF and encryption of metadata, without any
// blinding steps. This is useful for precomputing the buckets of encrypted
// items. The return value is the bucket ID (2 byte hash of username) as well
// as the ciphertext, both encoded as byte slices.
func (s *Server) EncryptBucketEntry(username string, password string, metadataFlag MetadataType, metadata []byte) ([]byte, error) {
	if !metadataFlag.Valid() {
		return nil, errors.New("invalid metadata flag value: " + string(metadataFlag))
	}
	key, err := s.deriveBucketEntryKey(username, password)
	if err != nil {
		return nil, err
	}

	return s.bucketEncryptor.Encrypt(key, metadataFlag, metadata)
}

// ServerResponse wraps up the server's response state.
type ServerResponse struct {
	Version          uint32 `json:"version"`
	EvaluatedElement []byte `json:"evaluatedElement"`
	BucketContents   []byte `json:"bucketContents"`
}

// MarshalBinary marshals the server response in the following binary format:
// <32-bit version>|<evaluated-element>|<bucket-contents>
func (r *ServerResponse) MarshalBinary() ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, r.Version); err != nil {
		return nil, err
	}
	if _, err := buffer.Write(r.EvaluatedElement); err != nil {
		return nil, err
	}
	if _, err := buffer.Write(r.BucketContents); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// UnmarshalBinary unmarshals the server response from the following binary format:
// <32-bit version>|<evaluated-element>|<bucket-contents>
func (r *ServerResponse) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &r.Version); err != nil {
		return err
	}
	sizes, err := oprf.GetSizes(DefaultOPRFSuite)
	if err != nil {
		return err
	}
	r.EvaluatedElement = make([]byte, sizes.SerializedElementLength)
	if n, err := buffer.Read(r.EvaluatedElement); err != nil {
		return err
	} else if n != len(r.EvaluatedElement) {
		return errors.New("too few bytes to deserialize EvaluatedElement")
	}
	r.BucketContents = buffer.Bytes()
	return nil
}

// Getter defines the interface needed for fetching bucket items to insert into
// a response. The caller should define an implementation of this interface
// appropriate for their deployment.
type Getter interface {
	Get(id string) ([]byte, error)
}

// HandleRequest takes as input a client request buffer and kv that implements
// the Getter interface. The request is a JSON encoding of a bucket
// identifier and oprf.IntValue  (a blinded group element) Should return a new
// IntValue (input group element multiplied by server's secret key) plus the
// bucket contents associated to the bucket identifier Returns a byte string
// that is a protobuf encoding of an oprf.IntValue (the Eval'd blinded value)
// plus the associated bucket
func (s *Server) HandleRequest(request ClientRequest, kv Getter) (ServerResponse, error) {
	if uint16(request.Version) != s.version {
		return ServerResponse{}, errors.New("requested version doesn't match server version")
	}

	evaluation, err := s.oprfServer.Evaluate([]oprf.Blinded{request.BlindElement}, OprfInfo)
	if err != nil {
		return ServerResponse{}, err
	}
	if len(evaluation.Elements) < 1 {
		return ServerResponse{}, errors.New("invalid Evaluation response")
	}

	_, err = hex.DecodeString(request.BucketID)
	if err != nil {
		return ServerResponse{}, errors.New("bucket ID not valid hex")
	}

	bucketContents, err := kv.Get(request.BucketID)
	if err != nil {
		return ServerResponse{}, err
	}

	return ServerResponse{
		Version:          request.Version,
		EvaluatedElement: evaluation.Elements[0],
		BucketContents:   bucketContents,
	}, nil
}
