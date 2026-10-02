package Gopurs_GoImports

import (
	"gopurs/output/gopurs_runtime"
)

// The PureScript signature guarantees TypeArray at both levels and TypeString
// elements. Borrow the input slices read-only, then copy their Values into an
// independent buffer. Packed string bytes remain immutable and may be shared.
// Preserve order and duplicates: collectImports owns deduplication and sorting.
func ConcatStringArrays(arrays gopurs_runtime.Value) gopurs_runtime.Value {
	outer := *(*[]gopurs_runtime.Value)(arrays.UnsafePtr)
	total := 0
	for _, inner := range outer {
		total += gopurs_runtime.ArrayLength(inner)
	}
	out := make([]gopurs_runtime.Value, 0, total)
	for _, inner := range outer {
		out = append(out, *(*[]gopurs_runtime.Value)(inner.UnsafePtr)...)
	}
	return gopurs_runtime.Array(out)
}
