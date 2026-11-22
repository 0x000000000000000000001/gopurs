package purescript

import "gopurs/output/gopurs_runtime"

func Data_Symbol_UnsafeCoerce(x interface{}) interface{} { return x }

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Symbol_UnsafeCoerce = // TAST: (ForAll [a, b] (Func [(TypeVar a)] (TypeVar b)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Symbol_UnsafeCoerce(go_arg0)
	return gopurs_runtime.Box(go_res)
})
