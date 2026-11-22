package purescript

import "gopurs/output/gopurs_runtime"

func Data_Semiring_IntAdd(x int64, y int64) int64 {
	return gopurs_runtime.IntAdd(x, y)
}
func Data_Semiring_IntMul(x int64, y int64) int64 {
	return gopurs_runtime.IntMul(x, y)
}
func Data_Semiring_NumAdd(n1 float64, n2 float64) float64 {
	return n1 + n2
}
func Data_Semiring_NumMul(n1 float64, n2 float64) float64 {
	return n1 * n2
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Semiring_IntAdd = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := Data_Semiring_IntAdd(go_arg0, go_arg1)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Data_Semiring_IntMul = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := Data_Semiring_IntMul(go_arg0, go_arg1)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Data_Semiring_NumAdd = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Semiring_NumAdd(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Semiring_NumMul = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Semiring_NumMul(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
