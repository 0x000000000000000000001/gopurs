package purescript

import "gopurs/output/gopurs_runtime"

import "sort"

func Data_Array_ST_NewImpl(s interface{}) *[]interface{} {
	arr := make([]interface{}, 0)
	return &arr
}

func Data_Array_ST_PeekImpl(just func(interface{}) interface{}, nothing interface{}, i int64, arr *[]interface{}) interface{} {
	if i >= 0 && i < int64(len(*arr)) {
		return just((*arr)[i])
	}
	return nothing
}

func Data_Array_ST_PokeImpl(i int64, a interface{}, arr *[]interface{}) interface{} {
	if i >= 0 && i < int64(len(*arr)) {
		(*arr)[i] = a
		return true
	}
	return false
}

func Data_Array_ST_LengthImpl(arr *[]interface{}) interface{} {
	return int64(len(*arr))
}

func Data_Array_ST_PopImpl(just func(interface{}) interface{}, nothing interface{}, arr *[]interface{}) interface{} {
	if len(*arr) > 0 {
		last := (*arr)[len(*arr)-1]
		*arr = (*arr)[:len(*arr)-1]
		return just(last)
	}
	return nothing
}

func Data_Array_ST_PushAllImpl(xs []interface{}, arr *[]interface{}) interface{} {
	*arr = append(*arr, xs...)
	return int64(len(*arr))
}

func Data_Array_ST_PushImpl(x interface{}, arr *[]interface{}) interface{} {
	*arr = append(*arr, x)
	return int64(len(*arr))
}

func Data_Array_ST_ShiftImpl(just func(interface{}) interface{}, nothing interface{}, arr *[]interface{}) interface{} {
	if len(*arr) > 0 {
		first := (*arr)[0]
		*arr = (*arr)[1:]
		return just(first)
	}
	return nothing
}

func Data_Array_ST_UnshiftAllImpl(xs []interface{}, arr *[]interface{}) interface{} {
	*arr = append(xs, *arr...)
	return int64(len(*arr))
}

func Data_Array_ST_UnshiftImpl(x interface{}, arr *[]interface{}) interface{} {
	*arr = append([]interface{}{x}, *arr...)
	return int64(len(*arr))
}

func Data_Array_ST_SpliceImpl(start int64, count int64, xs []interface{}, arr *[]interface{}) interface{} {
	removed := make([]interface{}, count)
	copy(removed, (*arr)[start:start+count])

	newArr := make([]interface{}, 0, len(*arr)-int(count)+len(xs))
	newArr = append(newArr, (*arr)[:start]...)
	newArr = append(newArr, xs...)
	newArr = append(newArr, (*arr)[start+count:]...)
	*arr = newArr
	return removed
}

func Data_Array_ST_UnsafeFreezeImpl(arr *[]interface{}) []interface{} {
	return *arr
}

func Data_Array_ST_UnsafeThawImpl(xs []interface{}) *[]interface{} {
	return &xs
}

func Data_Array_ST_FreezeImpl(arr *[]interface{}) []interface{} {
	return *arr
}

func Data_Array_ST_ThawImpl(xs []interface{}) *[]interface{} {
	return &xs
}

func Data_Array_ST_CloneImpl(arr *[]interface{}) *[]interface{} {
	res := make([]interface{}, len(*arr))
	copy(res, *arr)
	return &res
}

// stSortSlice avoids sort.SliceStable's reflection-based swaps.
type stSortSlice struct {
	xs   []interface{}
	less func(a, b interface{}) bool
}

func (s stSortSlice) Len() int           { return len(s.xs) }
func (s stSortSlice) Less(i, j int) bool { return s.less(s.xs[i], s.xs[j]) }
func (s stSortSlice) Swap(i, j int)      { s.xs[i], s.xs[j] = s.xs[j], s.xs[i] }

func Data_Array_ST_SortByImpl(f func(interface{}, interface{}) interface{}, toInt func(interface{}) int64, arr *[]interface{}) interface{} {
	sort.Stable(stSortSlice{
		xs: *arr,
		less: func(a, b interface{}) bool {
			return toInt(f(a, b)) < 0
		},
	})
	return arr
}

func Data_Array_ST_ToAssocArrayImpl(arr *[]interface{}) interface{} {
	res := make([]interface{}, len(*arr))
	for i, v := range *arr {
		res[i] = map[string]interface{}{
			"value": v,
			"index": int64(i),
		}
	}
	return res
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Array_ST_CloneImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *[]any
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*[]any)
	} else {
		go_arg0 = *(**[]any)(arg0.UnsafePtr)
	}
	go_res := Data_Array_ST_CloneImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_FreezeImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array (TypeVar a))]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *[]any
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*[]any)
	} else {
		go_arg0 = *(**[]any)(arg0.UnsafePtr)
	}
	go_res := Data_Array_ST_FreezeImpl(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Data_Array_ST_LengthImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *[]any
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*[]any)
	} else {
		go_arg0 = *(**[]any)(arg0.UnsafePtr)
	}
	go_res := Data_Array_ST_LengthImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_NewImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Array_ST_NewImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_PeekImpl = // TAST: (ForAll [h, a, r] (ADT ["Control","Monad","ST","Uncurried","STFn4"] [(Func [(TypeVar a)] (TypeVar r)), (TypeVar r), Int, (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (TypeVar r)]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	var go_arg3 *[]any
	if arg3.Type == gopurs_runtime.TypeAny {
		go_arg3 = arg3.AnyVal().(*[]any)
	} else {
		go_arg3 = *(**[]any)(arg3.UnsafePtr)
	}
	go_res := Data_Array_ST_PeekImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_PokeImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn3"] [Int, (TypeVar a), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Boolean]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := arg1
	var go_arg2 *[]any
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*[]any)
	} else {
		go_arg2 = *(**[]any)(arg2.UnsafePtr)
	}
	go_res := Data_Array_ST_PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_PopImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(ForAll [b] (Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]))), (ForAll [b] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	var go_arg2 *[]any
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*[]any)
	} else {
		go_arg2 = *(**[]any)(arg2.UnsafePtr)
	}
	go_res := Data_Array_ST_PopImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_PushAllImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(Array (TypeVar a)), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() []any {
		_source := arg0
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	var go_arg1 *[]any
	if arg1.Type == gopurs_runtime.TypeAny {
		go_arg1 = arg1.AnyVal().(*[]any)
	} else {
		go_arg1 = *(**[]any)(arg1.UnsafePtr)
	}
	go_res := Data_Array_ST_PushAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_PushImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(TypeVar a), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	var go_arg1 *[]any
	if arg1.Type == gopurs_runtime.TypeAny {
		go_arg1 = arg1.AnyVal().(*[]any)
	} else {
		go_arg1 = *(**[]any)(arg1.UnsafePtr)
	}
	go_res := Data_Array_ST_PushImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_ShiftImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(ForAll [b] (Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]))), (ForAll [b] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	var go_arg2 *[]any
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*[]any)
	} else {
		go_arg2 = *(**[]any)(arg2.UnsafePtr)
	}
	go_res := Data_Array_ST_ShiftImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_SortByImpl = // TAST: (ForAll [a, h] (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] [])), (Func [(ADT ["Data","Ordering","Ordering"] [])] Int), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
		return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
	}
	go_arg1 := func(p0_0 any) int64 {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		return gopurs_runtime.Unbox[int64](inner_res0)
	}
	var go_arg2 *[]any
	if arg2.Type == gopurs_runtime.TypeAny {
		go_arg2 = arg2.AnyVal().(*[]any)
	} else {
		go_arg2 = *(**[]any)(arg2.UnsafePtr)
	}
	go_res := Data_Array_ST_SortByImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_SpliceImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn4"] [Int, Int, (Array (TypeVar a)), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array (TypeVar a))]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := func() []any {
		_source := arg2
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	var go_arg3 *[]any
	if arg3.Type == gopurs_runtime.TypeAny {
		go_arg3 = arg3.AnyVal().(*[]any)
	} else {
		go_arg3 = *(**[]any)(arg3.UnsafePtr)
	}
	go_res := Data_Array_ST_SpliceImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_ThawImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(Array (TypeVar a)), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() []any {
		_source := arg0
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Array_ST_ThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_ToAssocArrayImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array (Record (Row [value: (TypeVar a), index: Int] Empty)))]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *[]any
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*[]any)
	} else {
		go_arg0 = *(**[]any)(arg0.UnsafePtr)
	}
	go_res := Data_Array_ST_ToAssocArrayImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_UnsafeFreezeImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array (TypeVar a))]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	var go_arg0 *[]any
	if arg0.Type == gopurs_runtime.TypeAny {
		go_arg0 = arg0.AnyVal().(*[]any)
	} else {
		go_arg0 = *(**[]any)(arg0.UnsafePtr)
	}
	go_res := Data_Array_ST_UnsafeFreezeImpl(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Data_Array_ST_UnsafeThawImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(Array (TypeVar a)), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() []any {
		_source := arg0
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Array_ST_UnsafeThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ST_UnshiftAllImpl = // TAST: (ForAll [h, a] (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(Array (TypeVar a)), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() []any {
		_source := arg0
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	var go_arg1 *[]any
	if arg1.Type == gopurs_runtime.TypeAny {
		go_arg1 = arg1.AnyVal().(*[]any)
	} else {
		go_arg1 = *(**[]any)(arg1.UnsafePtr)
	}
	go_res := Data_Array_ST_UnshiftAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
