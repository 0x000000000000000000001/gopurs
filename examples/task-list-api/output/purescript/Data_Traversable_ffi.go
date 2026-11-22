package purescript

import "gopurs/output/gopurs_runtime"

func Data_Traversable_TraverseArrayImpl(
	apply func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value,
	mapFn func(func(gopurs_runtime.Value) gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value,
	pure func(gopurs_runtime.Value) gopurs_runtime.Value,
	concat2 func(gopurs_runtime.Value) func(gopurs_runtime.Value) gopurs_runtime.Value,
	f func(gopurs_runtime.Value) gopurs_runtime.Value,
	arrayVal []gopurs_runtime.Value,
) gopurs_runtime.Value {
	array1 := func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Array([]gopurs_runtime.Value{a})
	}
	array2 := func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array([]gopurs_runtime.Value{a, b})
		})
	}
	array3 := func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Array([]gopurs_runtime.Value{a, b, c})
			})
		})
	}
	var goFn func(int, int) gopurs_runtime.Value
	goFn = func(bot, top int) gopurs_runtime.Value {
		switch top - bot {
		case 0:
			return pure(gopurs_runtime.Array([]gopurs_runtime.Value{}))
		case 1:
			return mapFn(array1, f(arrayVal[bot]))
		case 2:
			return apply(mapFn(array2, f(arrayVal[bot])), f(arrayVal[bot+1]))
		case 3:
			return apply(apply(mapFn(array3, f(arrayVal[bot])), f(arrayVal[bot+1])), f(arrayVal[bot+2]))
		default:
			pivot := bot + ((top-bot)/4)*2
			return apply(mapFn(func(x gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(concat2(x))
			}, goFn(bot, pivot)), goFn(pivot, top))
		}
	}
	return goFn(0, len(arrayVal))
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Traversable_TraverseArrayImpl = // TAST: (ForAll [m, a, b] (Func [(ForAll [x, y] (Func [(TypeApp (TypeVar m) [(Func [(TypeVar x)] (TypeVar y))]), (TypeApp (TypeVar m) [(TypeVar x)])] (TypeApp (TypeVar m) [(TypeVar y)]))), (ForAll [x, y] (Func [(Func [(TypeVar x)] (TypeVar y)), (TypeApp (TypeVar m) [(TypeVar x)])] (TypeApp (TypeVar m) [(TypeVar y)]))), (ForAll [x] (Func [(TypeVar x)] (TypeApp (TypeVar m) [(TypeVar x)]))), (Func [(Array (TypeVar b)), (Array (TypeVar b))] (Array (TypeVar b))), (Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (Array (TypeVar a))] (TypeApp (TypeVar m) [(Array (TypeVar b))])))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 gopurs_runtime.Value, p0_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(arg0, p0_0, p0_1)
	}
	go_arg1 := func(p0_0 func(gopurs_runtime.Value) gopurs_runtime.Value, p0_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(arg1, gopurs_runtime.Func(p0_0), p0_1)
	}
	go_arg2 := func(p0_0 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(arg2, p0_0)
	}
	go_arg3 := func(p0_0 gopurs_runtime.Value) func(gopurs_runtime.Value) gopurs_runtime.Value {
		inner_res0 := gopurs_runtime.Apply(arg3, p0_0)
		return func(p1_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(inner_res0, p1_0)
		}
	}
	go_arg4 := func(p0_0 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(arg4, p0_0)
	}
	go_arg5 := gopurs_runtime.Unbox[[]gopurs_runtime.Value](arg5)
	go_res := Data_Traversable_TraverseArrayImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return go_res
})
