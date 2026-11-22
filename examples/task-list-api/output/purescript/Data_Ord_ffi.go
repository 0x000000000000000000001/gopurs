package purescript

import "gopurs/output/gopurs_runtime"

func Data_Ord_OrdBooleanImpl(lt interface{}, eq interface{}, gt interface{}, x bool, y bool) interface{} {
	if !x && y {
		return lt
	} else if x == y {
		return eq
	}
	return gt
}
func Data_Ord_OrdIntImpl(lt interface{}, eq interface{}, gt interface{}, x int64, y int64) interface{} {
	if x < y {
		return lt
	}
	if x == y {
		return eq
	}
	return gt
}
func Data_Ord_OrdCharImpl[T any](lt T, eq T, gt T, x string, y string) T {
	if x < y {
		return lt
	}
	if x == y {
		return eq
	}
	return gt
}
func Data_Ord_OrdStringImpl[T any](lt T, eq T, gt T, x string, y string) T {
	if x < y {
		return lt
	}
	if x == y {
		return eq
	}
	return gt
}
func Data_Ord_OrdNumberImpl(lt interface{}, eq interface{}, gt interface{}, x float64, y float64) interface{} {
	if x < y {
		return lt
	}
	if x == y {
		return eq
	}
	return gt
}
func Data_Ord_OrdArrayImpl(f func(interface{}, interface{}) int64, xs []interface{}, ys []interface{}) int64 {
	xlen := len(xs)
	ylen := len(ys)
	for i := 0; i < xlen && i < ylen; i++ {
		o := f(xs[i], ys[i])
		if o != 0 {
			return o
		}
	}
	if xlen == ylen {
		return 0
	} else if xlen > ylen {
		return -1
	} else {
		return 1
	}
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Ord_OrdArrayImpl = // TAST: (ForAll [a] (Func [(Func [(TypeVar a), (TypeVar a)] Int), (Array (TypeVar a)), (Array (TypeVar a))] Int))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) int64 {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[int64](inner_res0)
	}
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_arg2 := func() []any {
		_source := arg2
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_Ord_OrdArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Data_Ord_OrdBooleanImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Boolean, Boolean] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[bool](arg3)
	go_arg4 := gopurs_runtime.Unbox[bool](arg4)
	go_res := Data_Ord_OrdBooleanImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Ord_OrdCharImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Char, Char] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_Ord_OrdCharImpl[gopurs_runtime.Value](go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return go_res
})
var _Gopurs_Data_Ord_OrdIntImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Int, Int] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[int64](arg3)
	go_arg4 := gopurs_runtime.Unbox[int64](arg4)
	go_res := Data_Ord_OrdIntImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Ord_OrdNumberImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), Number, Number] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[float64](arg3)
	go_arg4 := gopurs_runtime.Unbox[float64](arg4)
	go_res := Data_Ord_OrdNumberImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Ord_OrdStringImpl = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), String, String] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_Ord_OrdStringImpl[gopurs_runtime.Value](go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return go_res
})

// --- Boxed-result FFI workers ---
func Data_Ord_OrdBooleanImpl_nativeWorker(arg0 any, arg1 any, arg2 any, arg3 bool, arg4 bool) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Ord_OrdBooleanImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}

func Data_Ord_OrdIntImpl_nativeWorker(arg0 any, arg1 any, arg2 any, arg3 int64, arg4 int64) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Ord_OrdIntImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}

func Data_Ord_OrdNumberImpl_nativeWorker(arg0 any, arg1 any, arg2 any, arg3 float64, arg4 float64) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Ord_OrdNumberImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}
