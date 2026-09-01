package main

import (
	"fmt"
	"unsafe"
)

type Value struct {
	Type      int
	IntVal    int64
	UnsafePtr unsafe.Pointer
}

type Constructor_Test_Primes_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Test_Primes_Cons[T_a]
}

func Rebox_Test_Primes_359351273_3637802162(in *Constructor_Test_Primes_Cons[Value]) *Constructor_Test_Primes_Cons[int64] {
	if in == nil { return nil }
	out := &Constructor_Test_Primes_Cons[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Test_Primes_359351273_3637802162(in.V1)
	return out
}

func Rebox_Test_Primes_3637802162_359351273(in *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[Value] {
	if in == nil { return nil }
	out := &Constructor_Test_Primes_Cons[Value]{}
	out.V0 = Value{Type: 0, IntVal: in.V0}
	out.V1 = Rebox_Test_Primes_3637802162_359351273(in.V1)
	return out
}

func main() {
    v0 := &Constructor_Test_Primes_Cons[int64]{1, 10, &Constructor_Test_Primes_Cons[int64]{1, 20, nil}}
    
    lst := Rebox_Test_Primes_3637802162_359351273(v0)
    v_3 := Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst)}
    var v1_4 *Constructor_Test_Primes_Cons[Value] = nil
    
    for v_3.UnsafePtr != nil {
        val := (*Constructor_Test_Primes_Cons[Value])(v_3.UnsafePtr).V0.IntVal
        tmp := &Constructor_Test_Primes_Cons[int64]{1, val, Rebox_Test_Primes_359351273_3637802162(v1_4)}
        v1_4 = Rebox_Test_Primes_3637802162_359351273(tmp)
        
        v_3 = Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons[Value])(v_3.UnsafePtr).V1)}
    }
    
    fmt.Println("Finished loop")
}
