package main

import (
	"fmt"
	"time"
)

type Value interface{}

type Func func(Value) Value

func Apply(f Value, a Value) Value {
	return f.(Func)(a)
}

func main() {
	var n int64 = 100000000

	// 1. Boxed Closure via Apply
	f_boxed := Func(func(a Value) Value {
		return a.(int64) + 1
	})

	start1 := time.Now()
	var res1 int64 = 0
	for i := int64(0); i < n; i++ {
		res1 = Apply(f_boxed, res1).(int64)
	}
	fmt.Println("Boxed Apply:", time.Since(start1), res1)

	// 2. Local Typed Closure Variable
	f_native := func(a int64) int64 {
		return a + 1
	}

	start2 := time.Now()
	var res2 int64 = 0
	for i := int64(0); i < n; i++ {
		res2 = f_native(res2)
	}
	fmt.Println("Local Typed Closure:", time.Since(start2), res2)

    // 3. Global typed closure
    start3 := time.Now()
	var res3 int64 = 0
	for i := int64(0); i < n; i++ {
		res3 = global_f_native(res3)
	}
	fmt.Println("Global Typed Closure:", time.Since(start3), res3)
}

func global_f_native(a int64) int64 {
    return a + 1
}
