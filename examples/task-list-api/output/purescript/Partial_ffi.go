package purescript

import "gopurs/output/gopurs_runtime"

func Partial__CrashWith(msg string) interface{} {
	panic(msg)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Partial__CrashWith = // TAST: (ForAll [a] (Func [String] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Partial__CrashWith(go_arg0)
	return gopurs_runtime.Box(go_res)
})
