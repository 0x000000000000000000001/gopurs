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
	Vals []Value
}

func RecordDict(keys []string, vals []Value) Value {
	r := RecordData{keys, vals}
	return Value{Type: 25, UnsafePtr: unsafe.Pointer(&r)}
}

func main() {
	v := RecordDict([]string{"1", "2"}, []Value{})
	// overwrite stack
	func() {
		a := [10]int{9,9,9,9,9,9,9,9,9,9}
		_ = a
	}()
	r := (*RecordData)(v.UnsafePtr)
	fmt.Println(r.Keys)
}
