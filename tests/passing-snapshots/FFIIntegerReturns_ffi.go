package purescript

import "gopurs/output/gopurs_runtime"

func Main_ReturnInt64(value int64) int64 {
	return value
}

func Main_ReturnInt(value int) int {
	return value
}

func Main_ReturnInt64Array(values []int64) []int64 {
	return values
}

func Main_ReturnIntArray(values []int) []int {
	return values
}

func Main_DynamicInt64(value int64) any {
	return value
}

func Main_DynamicInt(value int) any {
	return value
}

func Main_DynamicString(value string) any {
	return value
}

func Main_DynamicBoolean(value bool) any {
	return value
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Main_DynamicBoolean = // TAST: (Func [Boolean] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_res := Main_DynamicBoolean(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_DynamicInt = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Main_DynamicInt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_DynamicInt64 = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_res := Main_DynamicInt64(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_DynamicString = // TAST: (Func [String] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Main_DynamicString(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_ReturnInt = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Main_ReturnInt(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Main_ReturnInt64 = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_res := Main_ReturnInt64(go_arg0)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Main_ReturnInt64Array = // TAST: (Func [(Array Int)] (Array Int))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]int64, len(arg0_arr))
	for i, v := range arg0_arr {
		go_arg0[i] = gopurs_runtime.Unbox[int64](v)
	}
	go_res := Main_ReturnInt64Array(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Int(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Main_ReturnIntArray = // TAST: (Func [(Array Int)] (Array Int))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]int, len(arg0_arr))
	for i, v := range arg0_arr {
		go_arg0[i] = gopurs_runtime.Unbox[int](v)
	}
	go_res := Main_ReturnIntArray(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Int(int64(v))
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
