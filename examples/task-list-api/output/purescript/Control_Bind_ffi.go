package purescript

import "gopurs/output/gopurs_runtime"

// ConcatMap natif sur les tableaux PureScript : entree et sortie en `Value`
// (aucune copie de conversion), le callback est applique element par element.
func Control_Bind_ArrayBind(arr []gopurs_runtime.Value, f func(gopurs_runtime.Value) []gopurs_runtime.Value) gopurs_runtime.Value {
	result := make([]gopurs_runtime.Value, 0, len(arr))
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return gopurs_runtime.Array(result)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Control_Bind_ArrayBind = // TAST: (ForAll [a, b] (Func [(Array (TypeVar a)), (Func [(TypeVar a)] (Array (TypeVar b)))] (Array (TypeVar b))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[[]gopurs_runtime.Value](arg0)
	go_arg1 := func(p0_0 gopurs_runtime.Value) []gopurs_runtime.Value {
		inner_res0 := gopurs_runtime.Apply(arg1, p0_0)
		return gopurs_runtime.Unbox[[]gopurs_runtime.Value](inner_res0)
	}
	go_res := Control_Bind_ArrayBind(go_arg0, go_arg1)
	return go_res
})
