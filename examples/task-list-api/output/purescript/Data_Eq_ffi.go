package purescript

import "gopurs/output/gopurs_runtime"

import r "gopurs/output/gopurs_runtime"

func refEq(r1 interface{}, r2 interface{}) bool {
	return r1 == r2
}
func Data_Eq_EqBooleanImpl(r1 bool, r2 bool) bool {
	return r1 == r2
}
func Data_Eq_EqIntImpl(r1 int64, r2 int64) bool {
	return r1 == r2
}
func Data_Eq_EqNumberImpl(r1 float64, r2 float64) bool {
	return r1 == r2
}
func Data_Eq_EqCharImpl(r1 string, r2 string) bool {
	return r1 == r2
}
func Data_Eq_EqStringImpl(r1 string, r2 string) bool {
	return r1 == r2
}
func Data_Eq_EqArrayImpl(f func(r.Value, r.Value) r.Value, xs []r.Value, ys []r.Value) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if !r.Unbox[bool](f(xs[i], ys[i])) {
			return false
		}
	}
	return true
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Eq_EqArrayImpl = // TAST: (ForAll [a] (Func [(Func [(TypeVar a), (TypeVar a)] Boolean), (Array (TypeVar a)), (Array (TypeVar a))] Boolean))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 r.Value, p0_1 r.Value) r.Value {
		inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		return gopurs_runtime.Unbox[r.Value](inner_res0)
	}
	go_arg1 := func() []r.Value {
		_source := arg1
		_out := make([]r.Value, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.Unbox[r.Value](gopurs_runtime.ArrayAccess(_source, _i))
		}
		return _out
	}()
	go_arg2 := func() []r.Value {
		_source := arg2
		_out := make([]r.Value, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.Unbox[r.Value](gopurs_runtime.ArrayAccess(_source, _i))
		}
		return _out
	}()
	go_res := Data_Eq_EqArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Eq_EqBooleanImpl = // TAST: (Func [Boolean, Boolean] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := Data_Eq_EqBooleanImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Eq_EqCharImpl = // TAST: (Func [Char, Char] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_Eq_EqCharImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Eq_EqIntImpl = // TAST: (Func [Int, Int] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := Data_Eq_EqIntImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Eq_EqNumberImpl = // TAST: (Func [Number, Number] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Data_Eq_EqNumberImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Eq_EqStringImpl = // TAST: (Func [String, String] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_Eq_EqStringImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
