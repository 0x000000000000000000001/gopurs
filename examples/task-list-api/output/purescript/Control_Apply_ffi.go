package purescript

import "gopurs/output/gopurs_runtime"

func Control_Apply_ArrayApply(fs []func(interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, 0, len(fs)*len(xs))
	for _, f := range fs {
		for _, x := range xs {
			result = append(result, f(x))
		}
	}
	return result
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Control_Apply_ArrayApply = // TAST: (ForAll [a, b] (Func [(Array (Func [(TypeVar a)] (TypeVar b))), (Array (TypeVar a))] (Array (TypeVar b))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() []func(any) any {
		_source := arg0
		_out := make([]func(any) any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = func(p1_0 any) any {
				return gopurs_runtime.Apply(gopurs_runtime.ArrayAccess(_source, _i), gopurs_runtime.Box(p1_0))
			}
		}
		return _out
	}()
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Control_Apply_ArrayApply(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
