package purescript

import "gopurs/output/gopurs_runtime"

func Data_Unfoldable1_Unfoldr1ArrayImpl(isNothing func(interface{}) bool, fromJust func(interface{}) interface{}, fst func(interface{}) interface{}, snd func(interface{}) interface{}, f func(interface{}) interface{}, b interface{}) []interface{} {
	var result []interface{}
	value := b
	for {
		tuple := f(value)
		result = append(result, fst(tuple))
		maybe := snd(tuple)
		if isNothing(maybe) {
			break
		}
		value = fromJust(maybe)
	}
	return result
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Unfoldable1_Unfoldr1ArrayImpl = // TAST: (ForAll [a, b] (Func [(ForAll [x] (Func [(ADT ["Data","Maybe","Maybe"] [(TypeVar x)])] Boolean)), (ForAll [x] (Func [(ADT ["Data","Maybe","Maybe"] [(TypeVar x)])] (TypeVar x))), (ForAll [x, y] (Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar x), (TypeVar y)])] (TypeVar x))), (ForAll [x, y] (Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar x), (TypeVar y)])] (TypeVar y))), (Func [(TypeVar b)] (ADT ["Data","Tuple","Tuple"] [(TypeVar a), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])])), (TypeVar b)] (Array (TypeVar a))))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) bool {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return gopurs_runtime.Unbox[bool](inner_res0)
	}
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
	}
	go_arg3 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg3, gopurs_runtime.Box(p0_0))
	}
	go_arg4 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg4, gopurs_runtime.Box(p0_0))
	}
	go_arg5 := arg5
	go_res := Data_Unfoldable1_Unfoldr1ArrayImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
