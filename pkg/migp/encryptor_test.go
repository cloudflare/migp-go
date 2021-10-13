// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package migp

import (
	"bytes"
	"testing"

	"github.com/cloudflare/circl/oprf"
)

// TestEncryptDecrypt tests that bucket entry can be encrypted and then
// correctly decrypted
func TestEncryptDecrypt(t *testing.T) {
	testCases := []struct {
		secret       []byte
		metadataFlag MetadataType
		metadata     []byte
	}{
		{[]byte("test"), MetadataDummy, []byte("helloworld")},
		{[]byte("124jbZC"), MetadataSimilarPassword, nil},
		{[]byte("1ujkbjkfdb09fjdzvzjkdfA!#"), 100, []byte("12")},
		{[]byte("abc345fds!#"), MetadataBreachedUsername, []byte("abc")},
	}

	privateKey, err := oprf.GenerateKey(DefaultOPRFSuite)
	if err != nil {
		t.Fatal(err)
	}

	oprfServer, err := oprf.NewServer(DefaultOPRFSuite, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	oprfClient, err := oprf.NewClient(DefaultOPRFSuite)
	if err != nil {
		t.Fatal(err)
	}

	bucketEncryptor, err := NewBucketEncryptor(DefaultBucketEncryptor)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range testCases {

		// Client generates blinded element
		oprfRequest, err := oprfClient.Request([][]byte{test.secret})
		if err != nil {
			t.Fatal(err)
		}

		// Server evaluates blinded element
		evaluatedMessage, err := oprfServer.Evaluate(oprfRequest.BlindedElements())
		if err != nil {
			t.Fatal(err)
		}

		// Client and server finalize to derive shared secret
		clientSecrets, err := oprfClient.Finalize(oprfRequest, evaluatedMessage)
		if err != nil {
			t.Fatal(err)
		}
		if len(clientSecrets) < 1 {
			t.Fatal("invalid Finalize response")
		}
		clientSecret := clientSecrets[0]
		serverSecret, err := oprfServer.FullEvaluate(test.secret)
		if err != nil {
			t.Fatal(err)
		}

		// Server encrypts bucket using shared secret
		ciphertext, err := bucketEncryptor.Encrypt(serverSecret, test.metadataFlag, test.metadata)
		if err != nil {
			t.Errorf("encryption failed with error %s", err.Error())
		}

		// Client decrypts header
		valid, flag, bodyLength, err := bucketEncryptor.DecryptHeader(clientSecret, ciphertext[:HeaderSize])
		if !valid {
			t.Errorf("header key check invalid")
		}
		if err != nil {
			t.Errorf("header decryption failed with error %s", err.Error())
		}
		if flag != test.metadataFlag {
			t.Errorf("decryption got mdFlag of %d (expected %d)", flag, test.metadataFlag)
		}
		if len(test.metadata) != bodyLength {
			t.Errorf("header decryption failed to recover length. got %d, expected %d", bodyLength, len(test.metadata))
		}

		// Client decrypts body
		metadata, err := bucketEncryptor.DecryptBody(clientSecret, ciphertext[HeaderSize:HeaderSize+bodyLength])
		if err != nil {
			t.Errorf("header decryption failed with error %s", err.Error())
		}
		if !bytes.Equal(metadata, test.metadata) {
			t.Errorf("decryption got mdString of '%s' (expected '%s')", metadata, test.metadata)
		}
	}
}

// BenchmarkHKDFSHA256Encryptor runs benchmark tests for the bucket encryptor
func BenchmarkHKDFSHA256Encryptor(b *testing.B) {
	secret := []byte{32}
	metadataFlag := MetadataDummy
	metadata := []byte{32}

	encryptor := NewHKDFSHA256BucketEncryptor()
	for i := 0; i < b.N; i++ {
		_, err := encryptor.Encrypt(secret, metadataFlag, metadata)
		if err != nil {
			b.Fatal(err)
		}
	}
}
