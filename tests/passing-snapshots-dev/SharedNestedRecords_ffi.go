package purescript

import "gopurs/output/gopurs_runtime"

func Main_EchoPayload(value gopurs_runtime.Value) gopurs_runtime.Value {
	return value
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Main_EchoPayload = // TAST: (ForAll [a] (Func [(TypeVar a)] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Main_EchoPayload(go_arg0)
	return go_res
})
