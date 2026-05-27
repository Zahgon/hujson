// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hujson

// TODO(dsnet): Insert/remove operations on an array has O(n) complexity
// where n is the length of the array. We could improve this with more clever
// data structure that has efficient insertion, deletion, and indexing.
// One possibility is an "order statistic tree", which provide O(log n)
// behavior for the necessary operations.
// See https://en.wikipedia.org/wiki/Order_statistic_tree.

// TODO(dsnet): Name lookup on an object has O(n) complexity performing
// a linear search through all names. This can be alleviated by building
// a map of names to indexes for relevant objects. Currently, we always insert
// a new member at the end of the members list, so that operation carries an
// amortized cost of O(1).

// TODO(dsnet): Cache intermediate lookups when resolving a JSON pointer.
// Patch operations tend to operate on paths that are related.
// Caching can reduce pointer lookup from O(n) to be closer to O(1)
// where n is the number of path segments in the JSON pointer.

// TODO(dsnet): Batch sequential insert/remove operations performed
// on the same object or array. This handles the possibly common case of batch
// inserting or removing a number of consecutive members/elements.
// Pointer caching may make this optimization unnecessary.

// Patch patches the value according to the provided patch file (per RFC 6902).
// The patch file may be in the HuJSON format where comments around and within
// a value being inserted are preserved. If the patch fails to fully apply,
// the receiver value will be left in a partially mutated state.
// Use Clone to preserve the original value.
//
// It does not format the value. It is recommended that Format be called after
// applying a patch.
func (v *Value) Patch(patch []byte) error { _ = "STUB: not implemented"; return nil }

type patchOperation struct {
	op    string // "add" | "remove" | "replace" | "move" | "copy" | "test"
	path  string // used by all operations
	from  string // used by "move" and "copy"
	value Value  // used by "add", "replace", and "test"
}

func parsePatch(patch []byte) ([]patchOperation, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *Value) patchAdd(i int, op patchOperation) error { _ = "STUB: not implemented"; return nil }

// only occurs for root

func (v *Value) patchRemoveOrReplace(i int, op patchOperation) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Value) patchMoveOrCopy(i int, op patchOperation) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(dsnet): For a move operation within the same object,
// we should simplify this as just a rename or replace.

func (v *Value) patchTest(i int, op patchOperation) error { _ = "STUB: not implemented"; return nil }

// hasPathPrefix is a stricter version of strings.HasPrefix where
// the prefix must end on a path segment boundary.
func hasPathPrefix(s, prefix string) bool { _ = "STUB: not implemented"; return false }

func equalValue(x, y Value) bool {
	_ = "STUB: not implemented"
	// TODO(dsnet): This definition of equality is both naive and slow.
	//   - It fails to properly compare strings with invalid UTF-8.
	//   - It fails to precisely compare integers beyond ±2⁵³.
	//   - It cannot handle values greater than ±math.MaxFloat64.
	//   - Comparison of objects with duplicate names has undefined behavior.
	return false
}

func (obj *Object) getAt(i int) ValueTrimmed { _ = "STUB: not implemented"; return *new(ValueTrimmed) }

func (obj *Object) setAt(i int, v ValueTrimmed) { _ = "STUB: not implemented"; return }

func (obj *Object) insertAt(i int, v ValueTrimmed) {
	_ = "STUB: not implemented"
	// TODO(dsnet): Use slices.Insert. See https://golang.org/issue/45955.
	return
}

func (obj *Object) removeAt(i int) ValueTrimmed {
	_ = "STUB: not implemented"
	// TODO(dsnet): Use slices.Delete. See https://golang.org/issue/45955.
	return *new(ValueTrimmed)
}

func (arr *Array) getAt(i int) ValueTrimmed { _ = "STUB: not implemented"; return *new(ValueTrimmed) }

func (arr *Array) setAt(i int, v ValueTrimmed) { _ = "STUB: not implemented"; return }

func (arr *Array) insertAt(i int, v ValueTrimmed) {
	_ = "STUB: not implemented"
	// TODO(dsnet): Use slices.Insert. See https://golang.org/issue/45955.
	return
}

func (arr *Array) removeAt(i int) ValueTrimmed {
	_ = "STUB: not implemented"
	// TODO(dsnet): Use slices.Delete. See https://golang.org/issue/45955.
	return *new(ValueTrimmed)
}

// Preserving and moving comments is impossible to perform reasonably in all
// conceivable situations given that the placement of comments is more
// a matter of human taste than it is a matter of mathematical rigor.
//
// We assume that:
//	* comments do not appear between the object member name and the colon
//	  (i.e., ObjectMember.Name.AfterExtra is nil),
//	* comments do not appear between the colon and the object member value
//	  (i.e., ObjectMember.Value.BeforeExtra is nil), and
//	* comments do not appear between the value and the comma
//	  (i.e., ObjectMember.Value.AfterExtra and ArrayElement.AfterExtra are nil).
// Such comments will be lost when patching.
//
// We further assume that:
//	* comments before an object member name and before an array element value
//	  are strongly associated with that member/element, and
//	* comments immediately after an object member value and after an
//	  array element value are strongly associated with that member/element.
// Such comments will be moved along with the member/element.
//
// Consider the following example:
//	{
//		...
//		// Comment1
//
//		// Comment2
//		"name": "value", // Comment3
//		// Comment4
//
//		// Comment5
//		...
//	}
//
// Moving "/name" will move only Comment2, Comment3, and Comment4.
// All other comments will be left alone.
//
// The above approach may perform contrary to expectation in this example:
//	{
//		// Comment1
//		"name1": "value1",
//		"name2": "value2",
//		"name3": "value3",
//	}
//
// Moving "/name1" will move Comment1. It is unclear whether Comment1 is
// strongly associated with just "name1" or the entire sequence of members
// from "name1" to "name2".

func copyAt(comp composite, i int) (v Value) { _ = "STUB: not implemented"; return *new(Value) }

func replaceAt(comp composite, i int, v Value) { _ = "STUB: not implemented"; return }

func insertAt(comp composite, i int, v Value) { _ = "STUB: not implemented"; return }

func removeAt(comp composite, i int) (v Value) { _ = "STUB: not implemented"; return *new(Value) }

// injectLeadingComments injects leading comments into the bottom of b.
func (b *Extra) injectLeadingComments(leading Extra) { _ = "STUB: not implemented"; return }

// extractLeadingComments extracts leading comments from the bottom of b.
// If readonly, then the source is not mutated.
func (b *Extra) extractLeadingComments(readonly bool) (leading Extra) {
	_ = "STUB: not implemented"
	return *new(Extra)
}

// injectTrailingComments injects trailing comments into the top of b.
func (b *Extra) injectTrailingComments(trailing Extra) { _ = "STUB: not implemented"; return }

// preserve trailing newline

// drop trailing newline

// extractTrailingcomments extracts trailing comments from the top of b.
// If readonly, then the source is not mutated.
func (b *Extra) extractTrailingcomments(readonly bool) (trailing Extra) {
	_ = "STUB: not implemented"
	return *new(Extra)
}

// preserve trailing newline

// classifyComments classifies comments as belonging to the previous element
// or belonging to the current element such that:
//   - b[:prevEnd] belongs to the previous element, and
//   - b[currStart:] belongs to the current element.
//
// Invariant: prevEnd <= currStart
func (b Extra) classifyComments() (prevEnd, currStart int) {
	_ = "STUB: not implemented"
	// Scan for dividers between comment blocks.
	return 0, 0
}

// adjust newline accounting for next iteration

// Without dividers, a line comment starting on the first line belongs
// to the previous element.

// Ownership is more clear when there is at least one divider.
