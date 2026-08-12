package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

type Value struct {
	Type      int
	IntVal    int64
	UnsafePtr unsafe.Pointer
}

type RecordData1 struct {
	K0 string
	V0 Value
}

func RecordToMap(v Value) map[string]Value {
	r := (*RecordData1)(v.UnsafePtr)
	return map[string]Value{r.K0: r.V0}
}

type Constructor_Eq struct {
	Rc uint32
	V0 Value
}

func CoerceToStruct[T any](val Value) *T {
	if val.Type == 9 {
		return (*T)(val.UnsafePtr)
	}
	res := new(T)
	resVal := reflect.ValueOf(res).Elem()
	m := RecordToMap(val)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	for i, k := range keys {
		resVal.Field(i + 1).Set(reflect.ValueOf(m[k]))
	}
	return res
}

func main() {
	v0 := Value{Type: 6, IntVal: 42}
	r := &RecordData1{K0: "eq", V0: v0}
	val := Value{Type: 20, UnsafePtr: unsafe.Pointer(r)}
	
	fmt.Printf("Before coerce\n")
	res := CoerceToStruct[Constructor_Eq](val)
	fmt.Printf("After coerce: %+v\n", res)
}
