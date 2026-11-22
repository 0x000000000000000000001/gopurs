package purescript

import "gopurs/output/gopurs_runtime"

import r "gopurs/output/gopurs_runtime"

func Data_FoldableWithIndex_FoldlWithIndexImpl(f func(r.Value, r.Value, r.Value) r.Value, z r.Value, xs []r.Value) r.Value {
	acc := z
	for i, v := range xs {
		acc = f(r.Int(int64(i)), acc, v)
	}
	return acc
}

func Data_FoldableWithIndex_FoldrWithIndexImpl(f func(r.Value, r.Value, r.Value) r.Value, z r.Value, xs []r.Value) r.Value {
	acc := z
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(r.Int(int64(i)), xs[i], acc)
	}
	return acc
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_FoldableWithIndex_FoldlWithIndexImpl = // TAST: (ForAll [a, b] (Func [(Func [Int, (TypeVar b), (TypeVar a)] (TypeVar b)), (TypeVar b), (Array (TypeVar a))] (TypeVar b)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 r.Value, p0_1 r.Value, p0_2 r.Value) r.Value {
		inner_res0 := gopurs_runtime.Apply3(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2))
		return gopurs_runtime.Unbox[r.Value](inner_res0)
	}
	go_arg1 := gopurs_runtime.Unbox[r.Value](arg1)
	go_arg2 := func() []r.Value {
		_source := arg2
		_out := make([]r.Value, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.Unbox[r.Value](gopurs_runtime.ArrayAccess(_source, _i))
		}
		return _out
	}()
	go_res := Data_FoldableWithIndex_FoldlWithIndexImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_FoldableWithIndex_FoldrWithIndexImpl = // TAST: (ForAll [a, b] (Func [(Func [Int, (TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (Array (TypeVar a))] (TypeVar b)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 r.Value, p0_1 r.Value, p0_2 r.Value) r.Value {
		inner_res0 := gopurs_runtime.Apply3(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2))
		return gopurs_runtime.Unbox[r.Value](inner_res0)
	}
	go_arg1 := gopurs_runtime.Unbox[r.Value](arg1)
	go_arg2 := func() []r.Value {
		_source := arg2
		_out := make([]r.Value, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.Unbox[r.Value](gopurs_runtime.ArrayAccess(_source, _i))
		}
		return _out
	}()
	go_res := Data_FoldableWithIndex_FoldrWithIndexImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
