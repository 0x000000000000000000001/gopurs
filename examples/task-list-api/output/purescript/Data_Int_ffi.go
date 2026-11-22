package purescript

import "gopurs/output/gopurs_runtime"

import (
	"math"
	"strconv"
)

func Data_Int_FromNumberImpl(just func(int) interface{}, nothing interface{}, n float64) interface{} {
	if math.IsNaN(n) || math.IsInf(n, 0) || math.Trunc(n) != n {
		return nothing
	}
	return just(int(n))
}

func Data_Int_ToNumber(n int) float64 {
	return float64(n)
}

func Data_Int_FromStringAsImpl(just func(int) interface{}, nothing interface{}, radix int, s string) interface{} {
	val, err := strconv.ParseInt(s, radix, 32)
	if err != nil {
		return nothing
	}
	return just(int(val))
}

func Data_Int_ToStringAs(radix int, i int) string {
	return strconv.FormatInt(int64(i), radix)
}

func Data_Int_Quot(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}

func Data_Int_Rem(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x % y
}

func Data_Int_Pow(x int, y int) int {
	if y < 0 {
		if x == 1 {
			return 1
		}
		if x == -1 {
			if y%2 == 0 {
				return 1
			}
			return -1
		}
		return 0
	}
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Int_FromNumberImpl = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), Number] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[float64](arg2)
	go_res := Data_Int_FromNumberImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Int_FromStringAsImpl = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_Int_FromStringAsImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Int_Pow = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Pow(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Quot = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Quot(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_Rem = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_Rem(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_Int_ToNumber = // TAST: (Func [Int] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Data_Int_ToNumber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Int_ToStringAs = // TAST: (Func [Int, Int] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Data_Int_ToStringAs(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

// --- Boxed-result FFI workers ---
func Data_Int_FromNumberImpl_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 float64) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Int_FromNumberImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}

func Data_Int_FromStringAsImpl_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 int, arg3 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_Int_FromStringAsImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
