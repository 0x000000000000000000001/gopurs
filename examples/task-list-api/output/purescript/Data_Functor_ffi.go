package purescript

import "gopurs/output/gopurs_runtime"

func Data_Functor_ArrayMap(f func(interface{}) interface{}, arr []interface{}) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Functor_ArrayMap = // TAST: (ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (Array (TypeVar a))] (Array (TypeVar b))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Functor_ArrayMap(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
