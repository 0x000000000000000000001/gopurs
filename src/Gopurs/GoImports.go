package Gopurs_GoImports

import (
	"gopurs/output/gopurs_runtime"
)

// ConcatStringArrays flattens an array of string arrays in one native pass.
// PureScript fold alternatives copy the growing accumulator at every step
// (`foldMap` on arrays is quadratic), and summing the imports of a large
// module's declarations dominated code generation.
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
