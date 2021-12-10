package migp

import (
	"bytes"
	"testing"
)

func TestSerializeUsernamePassword(t *testing.T) {
	tests := []struct {
		inUser []byte
		inPass []byte
		out    []byte
	}{
		{nil, nil, []byte{0, 0, 0, 0}},
		{[]byte{1, 2, 3}, []byte{4, 5, 6, 7}, []byte{0, 3, 1, 2, 3, 0, 4, 4, 5, 6, 7}},
	}
	for i, test := range tests {
		result := serializeUsernamePassword(test.inUser, test.inPass)
		if !bytes.Equal(result, test.out) {
			t.Errorf("failed test %d: want %v, got %v", i, test.out, result)
		}
	}
}

func TestBucketHashToID(t *testing.T) {
	tests := []struct {
		hash    []byte
		bitSize int
		out     uint32
	}{
		{[]byte{1, 2, 3, 4, 5, 6}, 0, 0},
		{[]byte{1, 2, 3, 4, 5, 6}, 15, 0x81},
		{[]byte{1, 2, 3, 4, 5, 6}, 16, 0x102},
		{[]byte{1, 2, 3, 4, 5, 6}, 32, 0x1020304},
	}
	for i, test := range tests {
		result := bucketHashToID(test.hash, test.bitSize)
		if result != test.out {
			t.Errorf("failed test %d: want %d, got %d", i, test.out, result)
		}
	}
}

func TestBucketIDToHex(t *testing.T) {
	tests := []struct {
		in  uint32
		out string
	}{
		{0x01020304, "01020304"},
		{0xff, "000000ff"},
		{0xffffffff, "ffffffff"},
	}
	for i, test := range tests {
		result := BucketIDToHex(test.in)
		if result != test.out {
			t.Errorf("failed test %d: want %q, got %q", i, test.out, result)
		}
	}
}
