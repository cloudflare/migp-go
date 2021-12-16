// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package mutator

import (
	"bytes"
	"encoding/json"
	"errors"
	"unicode"

	"github.com/spaolacci/murmur3"
)

// RDasRule struct is for initizializing the re-ordered Das rules. Rules are of
// form (RuleType, Position, String1, String2) where
// - RuleType is one of 'c' (capitalize), 'i' (insert), 's' (substitute), 'd' (delete prefix/suffix)
// - Position is a relative position (positive starts from 0 at beginning of string, negative
//   starts at -1 last position in string)
// - String1 is the string that is inserted for 'i', or the string that is matched
//   for substitution for 's'
// - String2 is empty except for 's' in which case it is the string that replaces String1
//
// Semantically the rules mean the following:
// - c changes capitalization of the first character (no other capitalization rules in this ruleset)
// - d removes first position characters from beginning (position > 0) or end (position < 0) of string
// - i inserts String1 at position
// - s replaces all occurrences of String1 with String2

type RDasRule struct {
	RuleType string `json:"ruletype"`
	Position int    `json:"position"`
	String1  string `json:"string1"`
	String2  string `json:"string2"`
}

// RDasMutator uses the ordered Das et al. mangling rules defined in dasrules.go
type RDasMutator struct {
	dasRules []RDasRule
}

// NewRDasMutator returns a new RDasMutator
func NewRDasMutator() *RDasMutator {
	m := new(RDasMutator)
	err := json.Unmarshal([]byte(dasRulesJSONString), &m.dasRules)
	if err != nil {
		panic("could not unmarshal Das rules")
	}
	return m
}

// switchCase switches an upper-case letter to a lower-case letter, and vice-versa
func switchCase(b byte) (byte, error) {
	r := rune(b)
	if unicode.IsLetter(r) {
		if unicode.IsUpper(r) {
			return byte(unicode.ToLower(r)), nil
		} else if unicode.IsLower(r) {
			return byte(unicode.ToUpper(r)), nil
		} else {
			return 0, errors.New("invalid rune")
		}
	}
	return 0, errors.New("not a letter")
}

// Mutate generates up to requested number of mutations. Returns a set of
// unique strings.  May return fewer than requested number, caller should
// check.
func (m *RDasMutator) Mutate(password []byte, num int) [][]byte {

	if len(m.dasRules) == 0 {
		panic("RDasMutator used without being initialized")
	}

	mutations := make([][]byte, 0, num)
	seen := make(map[uint32]struct{})
	seen[murmur3.Sum32(password)] = struct{}{}
	for i, j := 0, 0; j < num && i < len(m.dasRules); i++ {
		s := password
		rule := m.dasRules[i]

		// Rules were trained only on ASCII strings. We will anyway apply them
		// here, since if there are non-ASCII characters only other option
		// would be to just generate dummies, and we might nevertheless get
		// some benefit from applying mangling to UTF8 strings. (E.g., because
		// the first few characters are ASCII)

		position := rule.Position

		switch rule.RuleType {
		case "c":
			s = changeCap(s, position)
		case "d":
			s = deletePortion(s, position)
		case "i":
			s = insert(s, position, rule.String1)
		case "s":
			s = substitute(s, position, rule.String1, rule.String2)
		default:
			panic("One of the dasRules unrecognized")
		}

		key := murmur3.Sum32(s)
		if _, ok := seen[key]; !ok {
			j++
			seen[key] = struct{}{}
			mutations = append(mutations, s)
		}
	}
	return mutations
}

// changeCap returns a copy of the buffer with the case switched at the given
// position, if it's a letter
func changeCap(oldBuf []byte, position int) []byte {
	newBuf := make([]byte, len(oldBuf))
	copy(newBuf, oldBuf)
	if position < 0 {
		position = len(oldBuf) + position
	}
	if position >= 0 && position < len(oldBuf) {
		if b, err := switchCase(oldBuf[position]); err == nil {
			newBuf[position] = b
		}
	}
	return newBuf
}

// deletePortion returns a copy of the buffer with a deleted prefix or suffix
func deletePortion(oldBuf []byte, position int) []byte {
	var newBuf []byte
	if position >= 0 && position <= len(oldBuf) {
		newBuf = make([]byte, len(oldBuf)-position)
		copy(newBuf, oldBuf[position:])
	} else if position < 0 && len(oldBuf)+position >= 0 {
		newBuf = make([]byte, len(oldBuf)+position)
		copy(newBuf, oldBuf[:len(oldBuf)+position])
	} else {
		newBuf = make([]byte, len(oldBuf))
		copy(newBuf, oldBuf)
	}
	return newBuf
}

// insert returns a copy of the buffer with a substring inserted at the given
// position
func insert(oldBuf []byte, position int, string1 string) []byte {
	var newBuf []byte
	if len(oldBuf) == 0 && (position == 0 || position == -1) {
		newBuf = make([]byte, len(string1))
		copy(newBuf, string1)
	} else {
		if position < 0 {
			position = position + len(oldBuf) + 1
		}
		if position >= 0 && position <= len(oldBuf) {
			newBuf = make([]byte, len(oldBuf)+len(string1))
			copy(newBuf, oldBuf[:position])
			copy(newBuf[position:], string1)
			copy(newBuf[position+len(string1):], oldBuf[position:])
		} else {
			newBuf = make([]byte, len(oldBuf))
			copy(newBuf, oldBuf)
		}
	}
	return newBuf
}

// substitute returns a copy of the buffer with instances of one substring
// replaced with another substring
func substitute(buf []byte, _ int, string1, string2 string) []byte {
	return bytes.ReplaceAll(buf, []byte(string1), []byte(string2))
}
