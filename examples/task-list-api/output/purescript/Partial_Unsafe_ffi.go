package purescript

import "gopurs/output/gopurs_runtime"

func Partial_Unsafe__UnsafePartial(f func(interface{}) interface{}) interface{} {
	return f(nil)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Partial_Unsafe__UnsafePartial = // TAST: (ForAll [a, b] (Func [(TypeVar a)] (TypeVar b)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_res := Partial_Unsafe__UnsafePartial(go_arg0)
	return gopurs_runtime.Box(go_res)
})
