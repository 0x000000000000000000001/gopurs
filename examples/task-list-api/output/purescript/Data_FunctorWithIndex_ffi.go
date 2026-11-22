package purescript

import "gopurs/output/gopurs_runtime"

func Data_FunctorWithIndex_MapWithIndexArray(f func(int64, interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, len(xs))
	for i, x := range xs {
		result[i] = f(int64(i), x)
	}
	return result
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_FunctorWithIndex_MapWithIndexArray = // TAST: (ForAll [a, b] (Func [(Func [Int, (TypeVar a)] (TypeVar b)), (Array (TypeVar a))] (Array (TypeVar b))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int64, p0_1 any) any {
		return gopurs_runtime.Apply2(arg0, gopurs_runtime.Int(p0_0), gopurs_runtime.Box(p0_1))
	}
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_FunctorWithIndex_MapWithIndexArray(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
