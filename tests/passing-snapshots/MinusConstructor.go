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

var cache_Main_Tuple__158284501 gopurs_runtime.Value
var once_Main_Tuple__158284501 sync.Once

func Get_Main_Tuple__158284501() gopurs_runtime.Value {
	once_Main_Tuple__158284501.Do(func() {
		cache_Main_Tuple__158284501 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Rebox_Main_2291988990_1152654068(Call_Main_Tuple__158284501(Rebox_Main_1152654068_4045719060(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_1_0_box)), __eta_norm_0_1_box.IntVal)))}
		})
	})
	return cache_Main_Tuple__158284501
}

var cache_Main_Tuple__3273409525 gopurs_runtime.Value
var once_Main_Tuple__3273409525 sync.Once

func Get_Main_Tuple__3273409525() gopurs_runtime.Value {
	once_Main_Tuple__3273409525.Do(func() {
		cache_Main_Tuple__3273409525 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Rebox_Main_4045719060_1152654068(Call_Main_Tuple__3273409525(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.IntVal)))}
		})
	})
	return cache_Main_Tuple__3273409525
}

var cache_Main_Tuple__1512896373 gopurs_runtime.Value
var once_Main_Tuple__1512896373 sync.Once

func Get_Main_Tuple__1512896373() gopurs_runtime.Value {
	once_Main_Tuple__1512896373.Do(func() {
		cache_Main_Tuple__1512896373 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Rebox_Main_2590881780_1152654068(Call_Main_Tuple__1512896373(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box.StrVal())))}
		})
	})
	return cache_Main_Tuple__1512896373
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
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((Call_Main_test3(Rebox_Main_2291988990_4242115998((&Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, int64], int64]{1, (&Constructor_Main_Tuple[int64, int64]{1, int64(5), int64(10)}), int64(15)})))) == (int64(5)))), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
						}))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Tuple[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}

func Call_Main_Tuple__158284501(__eta_norm_1_0_loop *Constructor_Main_Tuple[int64, int64], __eta_norm_0_1_loop int64) *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, int64], int64] {
Tuple__158284501:
	for {
		if false {
			continue Tuple__158284501
		}
		var __eta_norm_1_0 *Constructor_Main_Tuple[int64, int64] = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, int64], int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_Tuple__3273409525(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop int64) *Constructor_Main_Tuple[int64, int64] {
Tuple__3273409525:
	for {
		if false {
			continue Tuple__3273409525
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Tuple[int64, int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_Tuple__1512896373(__eta_norm_1_0_loop string, __eta_norm_0_1_loop string) *Constructor_Main_Tuple[string, string] {
Tuple__1512896373:
	for {
		if false {
			continue Tuple__1512896373
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Tuple[string, string]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_test3(v_0_loop *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value]) int64 {
	var v_0 *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return ((v_0).V0).V0
}

func Rebox_Main_1152654068_4045719060(in *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_Tuple[int64, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[int64, int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1.IntVal
	return out
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

func Rebox_Main_2291988990_1152654068(in *Constructor_Main_Tuple[*Constructor_Main_Tuple[int64, int64], int64]) *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Rebox_Main_4045719060_1152654068(in.V0))}
	out.V1 = gopurs_runtime.Int(in.V1)
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

func Rebox_Main_2590881780_1152654068(in *Constructor_Main_Tuple[string, string]) *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Main_4045719060_1152654068(in *Constructor_Main_Tuple[int64, int64]) *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
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
