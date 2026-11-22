package purescript

import "gopurs/output/gopurs_runtime"

func Data_Array_NonEmpty_Internal_Foldr1Impl(f func(interface{}) func(interface{}) interface{}, xs []interface{}) interface{} {
	acc := xs[len(xs)-1]
	for i := len(xs) - 2; i >= 0; i-- {
		acc = f(xs[i])(acc)
	}
	return acc
}

func Data_Array_NonEmpty_Internal_Foldl1Impl(f func(interface{}) func(interface{}) interface{}, xs []interface{}) interface{} {
	acc := xs[0]
	length := len(xs)
	for i := 1; i < length; i++ {
		acc = f(acc)(xs[i])
	}
	return acc
}

type listNode struct {
	head interface{}
	tail interface{}
}

func Data_Array_NonEmpty_Internal_Traverse1Impl(apply func(interface{}) func(interface{}) interface{}, mapFn func(interface{}) func(interface{}) interface{}, f func(interface{}) interface{}, array []interface{}) interface{} {

	emptyList := &listNode{}

	consList := func(x interface{}) func(interface{}) interface{} {
		return func(xs interface{}) interface{} {
			xsNode := gopurs_runtime.Unbox[*listNode](xs.(gopurs_runtime.Value))
			return &listNode{head: x, tail: xsNode}
		}
	}

	finalCell := func(head interface{}) interface{} {
		return &listNode{head: head, tail: emptyList}
	}

	listToArray := func(list interface{}) interface{} {
		var arr []interface{}
		xs := gopurs_runtime.Unbox[*listNode](list.(gopurs_runtime.Value))
		for xs != emptyList {
			arr = append(arr, xs.head)
			xs = xs.tail.(*listNode)
		}
		if arr == nil {
			return []interface{}{}
		}
		return arr
	}

	buildFrom := func(x interface{}) func(interface{}) interface{} {
		return func(ys interface{}) interface{} {
			return apply(mapFn(consList)(f(x)))(ys)
		}
	}

	var goFn func(interface{}, int, []interface{}) interface{}
	goFn = func(acc interface{}, currentLen int, xs []interface{}) interface{} {
		if currentLen == 0 {
			return acc
		}
		last := xs[currentLen-1]
		return func() interface{} {
			return goFn(buildFrom(last)(acc), currentLen-1, xs)
		}
	}

	acc := mapFn(finalCell)(f(array[len(array)-1]))
	result := goFn(acc, len(array)-1, array)

	for {
		fn, isFunc := result.(func() interface{})
		if !isFunc {
			break
		}
		result = fn()
	}

	return mapFn(listToArray)(result)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Array_NonEmpty_Internal_Foldl1Impl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a), (TypeVar a)] (TypeVar a)), (Array (TypeVar a)), (TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Array_NonEmpty_Internal_Foldl1Impl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_NonEmpty_Internal_Foldr1Impl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a), (TypeVar a)] (TypeVar a)), (Array (TypeVar a)), (TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Array_NonEmpty_Internal_Foldr1Impl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_NonEmpty_Internal_Traverse1Impl = // TAST: (ForAll [m, a, b] (ADT ["Data","Function","Uncurried","Fn3"] [(ForAll [a', b'] (Func [(TypeApp (TypeVar m) [(Func [(TypeVar a')] (TypeVar b'))]), (TypeApp (TypeVar m) [(TypeVar a')])] (TypeApp (TypeVar m) [(TypeVar b')]))), (ForAll [a', b'] (Func [(Func [(TypeVar a')] (TypeVar b')), (TypeApp (TypeVar m) [(TypeVar a')])] (TypeApp (TypeVar m) [(TypeVar b')]))), (Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (Func [(Array (TypeVar a))] (TypeApp (TypeVar m) [(Array (TypeVar b))]))]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg1 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Func(func(p0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				if fn, ok := p0_0.(func(any) any); ok {
					return gopurs_runtime.Box(fn(gopurs_runtime.Unbox[any](p0)))
				}
				return gopurs_runtime.Apply(gopurs_runtime.Box(p0_0), p0)
			}()
		}))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg2 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
	}
	go_arg3 := func() []any {
		_source := arg3
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Array_NonEmpty_Internal_Traverse1Impl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
