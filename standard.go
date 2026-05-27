// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hujson

// IsStandard reports whether this is standard JSON
// by checking that there are no comments and no trailing commas.
func (v Value) IsStandard() bool { _ = "STUB: not implemented"; return false }

func (v *Value) isStandard() bool { _ = "STUB: not implemented"; return false }

// IsStandard reports whether this is standard JSON whitespace.
func (b Extra) IsStandard() bool { _ = "STUB: not implemented"; return false }

func (b Extra) hasComment() bool { _ = "STUB: not implemented"; return false }

// Minimize removes all whitespace, comments, and trailing commas from v,
// making it compliant with standard JSON per RFC 8259.
func (v *Value) Minimize() { _ = "STUB: not implemented"; return }

func (v *Value) minimize() { _ = "STUB: not implemented"; return }

// Standardize strips any features specific to HuJSON from v,
// making it compliant with standard JSON per RFC 8259.
// All comments and trailing commas are replaced with a space character
// in order to preserve the original line numbers and byte offsets.
func (v *Value) Standardize() { _ = "STUB: not implemented"; return }

// should be noop if offsets are already correct

func (v *Value) standardize() { _ = "STUB: not implemented"; return }

func (b *Extra) standardize() { _ = "STUB: not implemented"; return }

// NOTE: Avoid changing '\n' to keep line numbers the same.
