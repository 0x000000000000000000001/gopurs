package Gopurs_Cache_TypeMemo

import "gopurs/output/gopurs_runtime"

// ExprType is an immutable sum represented by a constructor tag and an optional
// payload pointer. Keep the whole Value as the key so the GC retains its owner.
type memo map[gopurs_runtime.Value]int64

func Create(_ any) gopurs_runtime.Value {
	return gopurs_runtime.Any(memo{})
}

func Lookup(handle, value gopurs_runtime.Value, _ any) int64 {
	if index, found := handle.AnyVal().(memo)[value]; found {
		return index
	}
	return -1
}

func Insert(handle, value gopurs_runtime.Value, index int64, _ any) {
	handle.AnyVal().(memo)[value] = index
}
