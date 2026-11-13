// Handwritten FFI for Gopurs.Cache.TypeMemo; keep in sync with the .purs foreign imports.
package Gopurs_Cache_TypeMemo

// Keep each complete comparable key, including any pointers it owns. The FFI
// bridge supplies the ExprType representation without exposing it to this cache.
// All operations use the same key type; only the map handle is type-erased.
type memo struct {
	indices any
}

func Create[T comparable]() *memo {
	return &memo{indices: make(map[T]int64)}
}

func Lookup[T comparable](handle *memo, value T, _ any) int64 {
	if index, found := handle.indices.(map[T]int64)[value]; found {
		return index
	}
	return -1
}

func Insert[T comparable](handle *memo, value T, index int64, _ any) {
	handle.indices.(map[T]int64)[value] = index
}
