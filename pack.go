// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hujson

// UpdateOffsets iterates through v and updates all
// Value.StartOffset and Value.EndOffset fields so that they are accurate.
func (v *Value) UpdateOffsets() { _ = "STUB: not implemented"; return }

func (v *Value) updateOffsets(n int) int { _ = "STUB: not implemented"; return 0 }

// Pack serializes the value as HuJSON.
// The output is valid so long as every Extra and Literal in the Value is valid.
// The output does not alias the memory of any buffers referenced by v.
func (v Value) Pack() []byte { _ = "STUB: not implemented"; return nil }

// String is a string representation of v.
func (v Value) String() string { _ = "STUB: not implemented"; return "" }

func (v Value) append(b []byte) []byte { _ = "STUB: not implemented"; return nil }
