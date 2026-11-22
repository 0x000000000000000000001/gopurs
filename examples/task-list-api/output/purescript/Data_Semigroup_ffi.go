package purescript

import "gopurs/output/gopurs_runtime"

func Data_Semigroup_ConcatString(s1 string, s2 string) string {
	return gopurs_runtime.ConcatString(s1, s2)
}
func Data_Semigroup_ConcatArray(xs []gopurs_runtime.Value, ys []gopurs_runtime.Value) []gopurs_runtime.Value {
	if len(xs) == 0 {
		return ys
	}
	if len(ys) == 0 {
		return xs
	}
	res := make([]gopurs_runtime.Value, 0, len(xs)+len(ys))
	res = append(res, xs...)
	res = append(res, ys...)
	return res
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Semigroup_ConcatArray = // TAST: (ForAll [a] (Func [(Array (TypeVar a)), (Array (TypeVar a))] (Array (TypeVar a))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[[]gopurs_runtime.Value](arg0)
	go_arg1 := gopurs_runtime.Unbox[[]gopurs_runtime.Value](arg1)
	go_res := Data_Semigroup_ConcatArray(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Semigroup_ConcatString = // TAST: (Func [String, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_Semigroup_ConcatString(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
