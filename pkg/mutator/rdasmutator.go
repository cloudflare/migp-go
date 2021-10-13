// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package mutator

import (
	"encoding/json"
	"errors"
	"strings"
	"unicode"
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
			newc := unicode.ToLower(r)
			return byte(newc), nil
		} else if unicode.IsLower(r) {
			newc := unicode.ToUpper(r)
			return byte(newc), nil
		} else {
			return 0, errors.New("invalid rune")
		}
	}
	return 0, errors.New("not a letter")
}

// Mutate generates up to requested number of mutations. Returns a set of
// unique strings.  May return fewer than requested number, caller should
// check.
func (m *RDasMutator) Mutate(password string, num int) []string {
	var mutations []string

	if len(m.dasRules) == 0 {
		panic("RDasMutator used without being initialized")
	}

	seen := make(map[string]struct{})
	seen[password] = struct{}{}
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

		if _, ok := seen[s]; !ok {
			j++
			seen[s] = struct{}{}
			mutations = append(mutations, s)
		}
	}
	return mutations
}

// changeCap switches the case at the given position, if it's a letter
func changeCap(oldStr string, position int) string {
	newStr := oldStr
	if position < 0 {
		position = len(oldStr) + position
	}
	if position >= 0 && position < len(oldStr) {
		b, err := switchCase(oldStr[position])
		if err == nil {
			newStr = oldStr[:position] + string(b) + oldStr[position+1:]
		}
	}
	return newStr
}

// deletePortion deletes a prefix or suffix
func deletePortion(oldStr string, position int) string {
	newStr := oldStr
	if position >= 0 && position <= len(oldStr) {
		newStr = oldStr[position:]
	} else if position < 0 && len(oldStr)+position >= 0 {
		newStr = oldStr[:len(oldStr)+position]
	}
	return newStr
}

// insert inserts a substring at the given position
func insert(oldStr string, position int, string1 string) string {
	newStr := oldStr
	if len(oldStr) == 0 && (position == 0 || position == -1) {
		newStr = string1
	} else {
		if position < 0 {
			position = position + len(oldStr) + 1
		}
		if position >= 0 && position <= len(oldStr) {
			newStr = oldStr[:position] + string1 + oldStr[position:]
		}
	}
	return newStr
}

// substitute swaps out instances of one substring with another substring
func substitute(oldStr string, _ int, string1 string, string2 string) string {
	newStr := oldStr
	if strings.Contains(oldStr, string1) {
		newStr = strings.Replace(oldStr, string1, string2, -1)
	}
	return newStr
}
