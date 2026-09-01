package main

import (
	"fmt"
	"unsafe"
)

type Constructor_Test_Primes_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Test_Primes_Cons[T_a]
}

type Value struct {
	Type      int
	IntVal    int64
	UnsafePtr unsafe.Pointer
}

func main() {
    v := &Constructor_Test_Primes_Cons[int64]{Rc: 1, V0: 42, V1: nil}
    val := Value{Type: 9, IntVal: 0, UnsafePtr: unsafe.Pointer(v)}
    
    // cast it to Value list!
    valList := (*Constructor_Test_Primes_Cons[Value])(val.UnsafePtr)
    
    fmt.Printf("valList.V0: %+v\n", valList.V0)
    fmt.Printf("valList.V1: %p\n", valList.V1)
}
