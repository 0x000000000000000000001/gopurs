package purescript

import "gopurs/output/gopurs_runtime"

func Effect_Unsafe_UnsafePerformEffect(f func(interface{}) interface{}) interface{} {
	return f(nil)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Effect_Unsafe_UnsafePerformEffect = // TAST: (ForAll [a] (Func [(ADT ["Effect","Effect"] [(TypeVar a)])] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_res := Effect_Unsafe_UnsafePerformEffect(go_arg0)
	return gopurs_runtime.Box(go_res)
})
