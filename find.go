// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hujson

import (
	"fmt"
)

var errNotFound = fmt.Errorf("value not found")

// Find locates the value specified by the JSON pointer (see RFC 6901).
// It returns nil if the value does not exist or the pointer is invalid.
// If a JSON object has multiple members matching a given name,
// the first is returned. Object names are matched exactly,
// rather than with a case-insensitive match.
func (v *Value) Find(ptr string) *Value { _ = "STUB: not implemented"; return nil }

type findState struct {
	pointer string // pointer[:offset] is the current value, pointer[offset:] is the remainder
	offset  int

	parent composite // nil for root pointer
	name   string    // name into parent to obtain current value
	idx    int       // idx into parent to obtain current value
	value  *Value    // the current value
}

func (v *Value) find(s findState) (findState, error) {
	_ = "STUB: not implemented"
	// An empty pointer denotes the value itself.
	return *new(findState), nil
}

// There must be one or more fragments.

// Unescape the name if necessary (section 4).

// Index into the object or array.

func (b Literal) equalString(s string) bool {
	_ = "STUB: not implemented"
	// Fast-path: Assume there are no escape characters.
	return false
}

// Slow-path: Unescape the string and then compare it.
// TODO(dsnet): Implement allocation-free comparison.
