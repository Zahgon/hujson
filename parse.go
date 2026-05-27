// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hujson

import (
	"errors"
)

func lineColumn(b []byte, n int) (line, column int) { _ = "STUB: not implemented"; return 0, 0 }

// Parse parses a HuJSON value as a Value.
// Extra and Literal values in v will alias the provided input buffer.
func Parse(b []byte) (Value, error) { _ = "STUB: not implemented"; return *new(Value), nil }

// parseNext parses the next value with surrounding whitespace and comments.
func parseNext(n int, b []byte) (v Value, _ int, err error) {
	_ = "STUB: not implemented"

	// Consume leading whitespace and comments.
	return *new(Value), 0, nil
}

// Parse the next value.

// Consume trailing whitespace and comments.

var (
	errInvalidObjectEnd = errors.New("invalid character '}' at start of value")
	errInvalidArrayEnd  = errors.New("invalid character ']' at start of value")
)

// parseNextTrimmed parses the next value without surrounding whitespace and comments.
func parseNextTrimmed(n int, b []byte) (ValueTrimmed, int, error) {
	_ = "STUB: not implemented"
	return *new(ValueTrimmed), 0, nil
}

// Parse objects.

// Parse the name.

// Parse the colon.

// Parse the value.

// Move AfterExtra from last value to AfterExtra of the object.

// Parse arrays.

// Move AfterExtra from last value to AfterExtra of the array.

// Parse strings.

// Parse null, booleans, and numbers.

var (
	lineCommentStart  = []byte("//")
	lineCommentEnd    = []byte("\n")
	blockCommentStart = []byte("/*")
	blockCommentEnd   = []byte("*/")
)

// consumeExtra consumes leading whitespace and comments.
func consumeExtra(n int, b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Skip past whitespace.

// Skip past comments.

func consumeWhitespace(b []byte) (n int) { _ = "STUB: not implemented"; return 0 }

// consumeComment consumes a line or block comment start in b.
// It returns the length of the comment if valid, otherwise
// it returns 0 if it is not a comment and -1 if it is invalid.
func consumeComment(b []byte) (n int) { _ = "STUB: not implemented"; return 0 }

func newInvalidCharacterError(prefix []byte, where string) error {
	_ = "STUB: not implemented"
	return nil
}
