// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hujson contains a parser and packer for the JWCC format:
// JSON With Commas and Comments (or "human JSON").
//
// JWCC is an extension of standard JSON (as defined in RFC 8259) in order to
// make it more suitable for humans and configuration files. In particular,
// it supports line comments (e.g., //...), block comments (e.g., /*...*/), and
// trailing commas after the last member or element in a JSON object or array.
//
// See https://nigeltao.github.io/blog/2021/json-with-commas-comments.html
//
// # Functionality
//
// The Parse function parses HuJSON input as a Value,
// which is a syntax tree exactly representing the input.
// Comments and whitespace are represented using the Extra type.
// Composite types in JSON are represented using the Object and Array types.
// Primitive types in JSON are represented using the Literal type.
// The Value.Pack method serializes the syntax tree as raw output,
// which is byte-for-byte identical to the input if no transformations
// were performed on the value.
//
// A HuJSON value can be transformed using the Minimize, Standardize, Format,
// or Patch methods. Each of these methods mutate the value in place.
// Call the Clone method beforehand in order to preserve the original value.
// The Minimize and Standardize methods coerces HuJSON into standard JSON.
// The Format method formats the value; it is similar to `go fmt`,
// but instead for the HuJSON and standard JSON format.
// The Patch method applies a JSON Patch (RFC 6902) to the receiving value.
//
// # Grammar
//
// The changes to the JSON grammar are:
//
//	--- grammar.json
//	+++ grammar.hujson
//	@@ -1,13 +1,31 @@
//	 members
//	 	member
//	+	member ',' ws
//	 	member ',' members
//
//	 elements
//	 	element
//	+	element ',' ws
//	 	element ',' elements
//
//	+comments
//	+	"*/"
//	+	comment comments
//	+
//	+comment
//	+	'0000' . '10FFFF'
//	+
//	+linecomments
//	+	'\n'
//	+	linecomment linecomments
//	+
//	+linecomment
//	+	'0000' . '10FFFF' - '\n'
//	+
//	 ws
//	 	""
//	+	"/*" comments
//	+	"//" linecomments
//	 	'0020' ws
//	 	'000A' ws
//	 	'000D' ws
//
// # Use with the Standard Library
//
// This package operates with HuJSON as an AST. In order to parse HuJSON
// into arbitrary Go types, use this package to parse HuJSON input as an AST,
// strip the AST of any HuJSON-specific lexicographical elements, and
// then pack the AST as a standard JSON output.
//
// Example usage:
//
//	b, err := hujson.Standardize(b)
//	if err != nil {
//		... // handle err
//	}
//	if err := json.Unmarshal(b, &v); err != nil {
//		... // handle err
//	}
package hujson

import (
	"iter"
)

// Kind reports the kind of the JSON value.
// It is the first byte of the grammar for that JSON value,
// with the exception that JSON numbers are represented as a '0'.
//
//	'n': null
//	'f': false
//	't': true
//	'"': string
//	'0': number
//	'{': object
//	'[': array
type Kind byte

// Value is an exact syntactic representation of a JSON value.
// The starting and ending byte offsets are populated when parsing,
// but are otherwise ignored when packing.
//
// By convention, code should operate on a non-pointer Value as a soft signal
// that the value should not be mutated, while operating on a pointer to Value
// to indicate that the value may be mutated. A non-pointer Value does not
// provide any language-enforced guarantees that it cannot be mutated.
// The Value.Clone method can be used to produce a deep copy of Value such that
// mutations on it will not be observed in the original Value.
type Value struct {
	// BeforeExtra are the comments and whitespace before Value.
	// This is the extra after the preceding open brace, open bracket,
	// colon, comma, or start of input.
	BeforeExtra Extra
	// StartOffset is the offset of the first byte in Value.
	StartOffset int
	// Value is the JSON value without surrounding whitespace or comments.
	Value ValueTrimmed
	// EndOffset is the offset of the next byte after Value.
	EndOffset int
	// AfterExtra are the comments and whitespace after Value.
	// This is the extra before the succeeding colon, comma, or end of input.
	AfterExtra Extra
}

// Clone returns a deep copy of the value.
func (v Value) Clone() Value { _ = "STUB: not implemented"; return *new(Value) }

// Range iterates through a Value in depth-first order and
// calls f for each value (including the root value).
// It stops iteration when f returns false.
//
// Deprecated: Use [All] instead.
func (v *Value) Range(f func(v *Value) bool) bool { _ = "STUB: not implemented"; return false }

// All returns an iterator over all values in depth-first order,
// starting with v itself.
func (v *Value) All() iter.Seq[*Value] { _ = "STUB: not implemented"; return nil }

// ValueTrimmed is a JSON value without surrounding whitespace or comments.
// This is a sum type consisting of Literal, *Object, or *Array.
type ValueTrimmed interface {
	// Kind reports the kind of the JSON value.
	Kind() Kind
	// clone returns a deep copy of the value.
	clone() ValueTrimmed

	isValueTrimmed()
}

var (
	_ ValueTrimmed = Literal(nil)
	_ ValueTrimmed = (*Object)(nil)
	_ ValueTrimmed = (*Array)(nil)
)

// Literal is the raw bytes for a JSON null, boolean, string, or number.
// It contains no surrounding whitespace or comments.
type Literal []byte // e.g., null, false, true, "string", or 3.14159

// Bool constructs a JSON literal for a boolean.
func Bool(v bool) Literal { _ = "STUB: not implemented"; return *new(Literal) }

// String constructs a JSON literal for string.
// Invalid UTF-8 is mangled with the Unicode replacement character.
func String(v string) Literal {
	_ = "STUB: not implemented"
	// Format according to RFC 8785, section 3.2.2.2.
	return *new(Literal)
}

// Int construct a JSON literal for a signed integer.
func Int(v int64) Literal { _ = "STUB: not implemented"; return *new(Literal) }

// Uint construct a JSON literal for an unsigned integer.
func Uint(v uint64) Literal { _ = "STUB: not implemented"; return *new(Literal) }

// Float construct a JSON literal for a floating-point number.
// The values NaN, +Inf, and -Inf will be represented as a JSON string
// with the values "NaN", "Infinity", and "-Infinity".
func Float(v float64) Literal { _ = "STUB: not implemented"; return *new(Literal) }

func (b Literal) clone() ValueTrimmed { _ = "STUB: not implemented"; return *new(ValueTrimmed) }

// Kind represents each possible JSON literal kind with a single byte,
// which is conveniently the first byte of that kind's grammar
// with the restriction that numbers always be represented with '0'.
func (b Literal) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

// IsValid reports whether b is a valid JSON null, boolean, string, or number.
// The literal must not have surrounding whitespace.
func (b Literal) IsValid() bool {
	_ = "STUB: not implemented"
	// NOTE: The v1 json package is non-compliant with RFC 8259, section 8.1
	// in that it does not enforce the use of valid UTF-8.
	return false
}

// Bool returns the value for a JSON boolean.
// It returns false if the literal is not a JSON boolean.
func (b Literal) Bool() bool { _ = "STUB: not implemented"; return false }

// String returns the unescaped string value for a JSON string.
// For other JSON kinds, this returns the raw JSON represention.
func (b Literal) String() (s string) { _ = "STUB: not implemented"; return "" }

// Int returns the signed integer value for a JSON number.
// It returns 0 if the literal is not a signed integer.
func (b Literal) Int() (n int64) { _ = "STUB: not implemented"; return 0 }

// Uin returns the unsigned integer value for a JSON number.
// It returns 0 if the literal is not an unsigned integer.
func (b Literal) Uint() (n uint64) { _ = "STUB: not implemented"; return 0 }

// Float returns the floating-point value for a JSON number.
// It returns a NaN, +Inf, or -Inf value for any JSON string with the values
// "NaN", "Infinity", or "-Infinity".
// It returns 0 for all other cases.
func (b Literal) Float() (n float64) { _ = "STUB: not implemented"; return 0 }

func (Literal) isValueTrimmed() {
	_ = "STUB: not implemented"

	// Object is an exact syntactic representation of a JSON object.
	return
}

type Object struct {
	// Members are the members of a JSON object.
	// A trailing comma is emitted only if the Value.AfterExtra
	// on the last value is non-nil. Otherwise it is omitted.
	Members []ObjectMember
	// AfterExtra are the comments and whitespace
	// after the preceding open brace or comma and before the closing brace.
	AfterExtra Extra
}

type ObjectMember struct {
	Name, Value Value
}

func (obj Object) length() int { _ = "STUB: not implemented"; return 0 }

func (obj Object) firstValue() *Value { _ = "STUB: not implemented"; return nil }

// allValues iterates all members of the object,
// interleaved between the member name and the member value.
func (obj Object) allValues() iter.Seq[*Value] { _ = "STUB: not implemented"; return nil }

func (obj Object) lastValue() *Value { _ = "STUB: not implemented"; return nil }

func (obj *Object) beforeExtraAt(i int) *Extra { _ = "STUB: not implemented"; return nil }

func (obj *Object) afterExtra() *Extra { _ = "STUB: not implemented"; return nil }

func (obj Object) clone() ValueTrimmed { _ = "STUB: not implemented"; return *new(ValueTrimmed) }

func (Object) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (*Object) isValueTrimmed() {
	_ = "STUB: not implemented"

	// Array is an exact syntactic representation of a JSON array.
	return
}

type Array struct {
	// Elements are the elements of a JSON array.
	// A trailing comma is emitted only if the Value.AfterExtra
	// on the last value is non-nil. Otherwise it is omitted.
	Elements []ArrayElement
	// AfterExtra are the comments and whitespace
	// after the preceding open bracket or comma and before the closing bracket.
	AfterExtra Extra
}

type ArrayElement = Value

func (arr Array) length() int { _ = "STUB: not implemented"; return 0 }

func (arr Array) firstValue() *Value { _ = "STUB: not implemented"; return nil }

// allValues iterates all elements of the array.
func (arr Array) allValues() iter.Seq[*Value] { _ = "STUB: not implemented"; return nil }

func (arr Array) lastValue() *Value { _ = "STUB: not implemented"; return nil }

func (arr *Array) beforeExtraAt(i int) *Extra { _ = "STUB: not implemented"; return nil }

func (arr *Array) afterExtra() *Extra { _ = "STUB: not implemented"; return nil }

func (arr Array) clone() ValueTrimmed { _ = "STUB: not implemented"; return *new(ValueTrimmed) }

func (Array) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (*Array) isValueTrimmed() {
	_ = "STUB: not implemented"

	// composite are the common methods of Object and Array.
	return
}

type composite interface {
	Kind() Kind
	length() int

	firstValue() *Value
	allValues() iter.Seq[*Value]
	lastValue() *Value

	getAt(int) ValueTrimmed
	setAt(int, ValueTrimmed)
	insertAt(int, ValueTrimmed)
	removeAt(int) ValueTrimmed

	beforeExtraAt(int) *Extra
	afterExtra() *Extra
}

func hasTrailingComma(comp composite) bool { _ = "STUB: not implemented"; return false }

func setTrailingComma(comp composite, v bool) { _ = "STUB: not implemented"; return }

var (
	_ composite = (*Object)(nil)
	_ composite = (*Array)(nil)
)

// Extra is the raw bytes for whitespace and comments.
// Whitespace per RFC 8259, section 2 are permitted.
// Line comments that start with "//" and end with "\n" are permitted.
// Block comments that start with "/*" and end with "*/" are permitted.
type Extra []byte

// IsValid reports whether the whitespace and comments are valid
// according to the HuJSON grammar.
func (b Extra) IsValid() bool { _ = "STUB: not implemented"; return false }

func copyBytes(b []byte) []byte { _ = "STUB: not implemented"; return nil }
