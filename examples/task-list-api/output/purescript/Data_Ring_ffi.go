package purescript

import "gopurs/output/gopurs_runtime"

func Data_Ring_IntSub(x int64, y int64) int64 {
	return gopurs_runtime.IntSub(x, y)
}
func Data_Ring_NumSub(x float64, y float64) float64 {
	return x - y
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Ring_IntSub = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := Data_Ring_IntSub(go_arg0, go_arg1)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Data_Ring_NumSub = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Ring_NumSub(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
