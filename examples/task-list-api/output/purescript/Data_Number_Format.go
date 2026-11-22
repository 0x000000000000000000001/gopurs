package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Number_Format_clamp gopurs_runtime.Value
var once_Data_Number_Format_clamp sync.Once

func Get_Data_Number_Format_clamp() gopurs_runtime.Value {
	once_Data_Number_Format_clamp.Do(func() {
		cache_Data_Number_Format_clamp = Get_Data_Ord_clamp__1405043522()
	})
	return cache_Data_Number_Format_clamp
}

var cache_Data_Number_Format_Precision gopurs_runtime.Value
var once_Data_Number_Format_Precision sync.Once

func Get_Data_Number_Format_Precision() gopurs_runtime.Value {
	once_Data_Number_Format_Precision.Do(func() {
		cache_Data_Number_Format_Precision = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1786680275, UnsafePtr: unsafe.Pointer((&Constructor_Data_Number_Format_Precision{1, value0.IntVal}))}
		})
	})
	return cache_Data_Number_Format_Precision
}

var cache_Data_Number_Format_Fixed gopurs_runtime.Value
var once_Data_Number_Format_Fixed sync.Once

func Get_Data_Number_Format_Fixed() gopurs_runtime.Value {
	once_Data_Number_Format_Fixed.Do(func() {
		cache_Data_Number_Format_Fixed = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1390206259, UnsafePtr: unsafe.Pointer((&Constructor_Data_Number_Format_Fixed{1, value0.IntVal}))}
		})
	})
	return cache_Data_Number_Format_Fixed
}

var cache_Data_Number_Format_Exponential gopurs_runtime.Value
var once_Data_Number_Format_Exponential sync.Once

func Get_Data_Number_Format_Exponential() gopurs_runtime.Value {
	once_Data_Number_Format_Exponential.Do(func() {
		cache_Data_Number_Format_Exponential = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1734244434, UnsafePtr: unsafe.Pointer((&Constructor_Data_Number_Format_Exponential{1, value0.IntVal}))}
		})
	})
	return cache_Data_Number_Format_Exponential
}

var cache_Data_Number_Format_toStringWith gopurs_runtime.Value
var once_Data_Number_Format_toStringWith sync.Once

func Get_Data_Number_Format_toStringWith() gopurs_runtime.Value {
	once_Data_Number_Format_toStringWith.Do(func() {
		cache_Data_Number_Format_toStringWith = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Number_Format_toStringWith(v_0_box)
		})
	})
	return cache_Data_Number_Format_toStringWith
}

var cache_Data_Number_Format_precision gopurs_runtime.Value
var once_Data_Number_Format_precision sync.Once

func Get_Data_Number_Format_precision() gopurs_runtime.Value {
	once_Data_Number_Format_precision.Do(func() {
		cache_Data_Number_Format_precision = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Number_Format_Precision(), gopurs_runtime.Apply2(Get_Data_Ord_clamp__1405043522(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(21))))
	})
	return cache_Data_Number_Format_precision
}

var cache_Data_Number_Format_fixed gopurs_runtime.Value
var once_Data_Number_Format_fixed sync.Once

func Get_Data_Number_Format_fixed() gopurs_runtime.Value {
	once_Data_Number_Format_fixed.Do(func() {
		cache_Data_Number_Format_fixed = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Number_Format_Fixed(), gopurs_runtime.Apply2(Get_Data_Ord_clamp__1405043522(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(20))))
	})
	return cache_Data_Number_Format_fixed
}

var cache_Data_Number_Format_exponential gopurs_runtime.Value
var once_Data_Number_Format_exponential sync.Once

func Get_Data_Number_Format_exponential() gopurs_runtime.Value {
	once_Data_Number_Format_exponential.Do(func() {
		cache_Data_Number_Format_exponential = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Number_Format_Exponential(), gopurs_runtime.Apply2(Get_Data_Ord_clamp__1405043522(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(20))))
	})
	return cache_Data_Number_Format_exponential
}

type Constructor_Data_Number_Format_Precision struct {
	Rc uint32
	V0 int64
}

type Constructor_Data_Number_Format_Fixed struct {
	Rc uint32
	V0 int64
}

type Constructor_Data_Number_Format_Exponential struct {
	Rc uint32
	V0 int64
}

func Call_Data_Number_Format_toStringWith(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 gopurs_runtime.Value
	{
		if v_0.Type == 9 && v_0.IntVal == 1786680275 {
			__t0 = gopurs_runtime.Apply(Get_Data_Number_Format_toPrecisionNative(), gopurs_runtime.Int((*Constructor_Data_Number_Format_Precision)(v_0.UnsafePtr).V0))
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 1390206259 {
			__t0 = gopurs_runtime.Apply(Get_Data_Number_Format_toFixedNative(), gopurs_runtime.Int((*Constructor_Data_Number_Format_Fixed)(v_0.UnsafePtr).V0))
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 1734244434 {
			__t0 = gopurs_runtime.Apply(Get_Data_Number_Format_toExponentialNative(), gopurs_runtime.Int((*Constructor_Data_Number_Format_Exponential)(v_0.UnsafePtr).V0))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Get_Data_Number_Format_toExponentialNative() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Format_ToExponentialNative
}

func Get_Data_Number_Format_toFixedNative() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Format_ToFixedNative
}

func Get_Data_Number_Format_toPrecisionNative() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Format_ToPrecisionNative
}

func Get_Data_Number_Format_toString() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Format_ToString
}
