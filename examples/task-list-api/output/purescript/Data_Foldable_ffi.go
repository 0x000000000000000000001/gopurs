package purescript

import "gopurs/output/gopurs_runtime"

func Data_Foldable_FoldrArray(f func(interface{}, interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i], acc)
	}
	return acc
}

func Data_Foldable_FoldlArray(f func(interface{}, interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := 0; i < len(xs); i++ {
		acc = f(acc, xs[i])
	}
	return acc
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Foldable_FoldlArray = // TAST: (ForAll [a, b] (Func [(Func [(TypeVar b), (TypeVar a)] (TypeVar b)), (TypeVar b), (Array (TypeVar a))] (TypeVar b)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
		return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
	}
	go_arg1 := arg1
	go_arg2 := func() []any {
		_source := arg2
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Foldable_FoldlArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Foldable_FoldrArray = // TAST: (ForAll [a, b] (Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (Array (TypeVar a))] (TypeVar b)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
		return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
	}
	go_arg1 := arg1
	go_arg2 := func() []any {
		_source := arg2
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Foldable_FoldrArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
