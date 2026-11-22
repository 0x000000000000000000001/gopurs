package purescript

import "gopurs/output/gopurs_runtime"

func Data_Array_ST_Partial_PeekImpl(i int64, xs *[]interface{}) interface{} {
	return (*xs)[i]
}
func Data_Array_ST_Partial_PokeImpl(i int64, a interface{}, xs *[]interface{}) bool {
	(*xs)[i] = a
	return true
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Array_ST_Partial_PeekImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn2"] [Int, (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	var go_arg1 *[]any
	if arg1.Type == gopurs_runtime.TypeAny {
		go_arg1 = arg1.AnyVal().(*[]any)
	} else {
		go_arg1 = *(**[]any)(arg1.UnsafePtr)
	}
	go_res := Data_Array_ST_Partial_PeekImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_Partial_PokeImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn3"] [Int, (TypeVar a), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Unit]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := arg1
	var go_arg2 *[]any
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*[]any)
	} else {
		go_arg2 = *(**[]any)(arg2.UnsafePtr)
	}
	go_res := Data_Array_ST_Partial_PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
