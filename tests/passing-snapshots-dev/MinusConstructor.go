package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Tuple gopurs_runtime.Value
var once_Main_Tuple sync.Once

func Get_Main_Tuple() gopurs_runtime.Value {
	once_Main_Tuple.Do(func() {
		cache_Main_Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Tuple
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Int(int64(8))
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Int(int64(7))
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_test3(Rebox_Main_1152654068_4242115998(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))))
		})
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Int(int64(3))
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Str("")
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((Call_Main_test3(Rebox_Main_2291988990_4242115998((&Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, int64], int64]{1, (&Constructor_Main_Tuple[int64, int64]{1, int64(5), int64(10)}), int64(15)})))) == (int64(5)))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Tuple[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}

func Call_Main_test3(v_0_loop *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value]) int64 {
	var v_0 *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return ((v_0).V0).V0
}

func Rebox_Main_1152654068_4242115998(in *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value]{}
	out.V0 = Rebox_Main_1152654068_813227567(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = in.V1
	return out
}

func Rebox_Main_1152654068_813227567(in *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_Tuple[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[int64, gopurs_runtime.Value]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1
	return out
}

func Rebox_Main_2291988990_4242115998(in *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, int64], int64]) *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value]{}
	out.V0 = Rebox_Main_4045719060_813227567(in.V0)
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_4045719060_813227567(in *Constructor_Main_Tuple[int64, int64]) *Constructor_Main_Tuple[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[int64, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}
