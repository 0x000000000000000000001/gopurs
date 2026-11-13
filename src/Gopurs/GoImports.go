// Handwritten FFI for Gopurs.GoImports; keep in sync with the .purs foreign imports.
package Gopurs_GoImports

// Borrow the input slices read-only, then copy their elements into an
// independent buffer. The generic element preserves the caller's representation.
// Preserve order and duplicates: collectImports owns deduplication and sorting.
func ConcatStringArrays[T any](arrays [][]T) []T {
	total := 0
	for _, inner := range arrays {
		total += len(inner)
	}
	out := make([]T, 0, total)
	for _, inner := range arrays {
		out = append(out, inner...)
	}
	return out
}
