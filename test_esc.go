package main

import (
	"fmt"
	"unsafe"
)

type Value struct {
	Type      int
	UnsafePtr unsafe.Pointer
}

type RecordData struct {
	Keys []string
}

//go:noinline
func RecordDict(keys []string) Value {
	r := &RecordData{keys}
	return Value{Type: 25, UnsafePtr: unsafe.Pointer(r)}
}

//go:noinline
func main() {
	v := RecordDict([]string{"A"})
	func() {
		a := [100]int{}
		_ = a
	}()
	r := (*RecordData)(v.UnsafePtr)
	fmt.Println(r.Keys)
}
