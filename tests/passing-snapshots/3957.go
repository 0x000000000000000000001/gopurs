package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Just
}

var cache_Main_weirdsum gopurs_runtime.Value
var once_Main_weirdsum sync.Once

func Get_Main_weirdsum() gopurs_runtime.Value {
	once_Main_weirdsum.Do(func() {
		cache_Main_weirdsum = gopurs_runtime.Func3(func(accum_0_box gopurs_runtime.Value, f1_1_box gopurs_runtime.Value, n_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_weirdsum(accum_0_box.IntVal, f1_1_box, n_2_box.IntVal))
		})
	})
	return cache_Main_weirdsum
}

var cache_Main_tricksyinners gopurs_runtime.Value
var once_Main_tricksyinners sync.Once

func Get_Main_tricksyinners() gopurs_runtime.Value {
	once_Main_tricksyinners.Do(func() {
		cache_Main_tricksyinners = gopurs_runtime.Func2(func(accum_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_tricksyinners(accum_0_box.IntVal, x_1_box.IntVal))
		})
	})
	return cache_Main_tricksyinners
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(x_0_box.IntVal))
		})
	})
	return cache_Main_g
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_f(x_0_box.IntVal))
		})
	})
	return cache_Main_f
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_0_0 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_f(99999)), gopurs_runtime.Int(0))
			_ = __local_var_0_0
			// TAST (Let): result_1_1 shape=Other bindingType=Boolean
			result_1_1 := (gopurs_runtime.RecordGet(__local_var_0_0, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_0_0, "expected").IntVal)
			_ = result_1_1
			// TAST (Let): message_2_2 shape=Other bindingType=String
			message_2_2 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "actual")).StrVal())
			_ = message_2_2
			// TAST (Let): __local_var_3_3 shape=Let(Let(EffectBind(App(Var)))) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_3_3 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_4 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
				__local_var_3_4 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_2_2))
				_ = __local_var_3_4
				var __t6 gopurs_runtime.Value
				{
					if (result_1_1) != (true) {
						__t6 = __local_var_3_4
						goto end_branch_6
					} else {

					}
				}
				{
					if result_1_1 {
						__t6 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							return Get_Data_Unit_unit()
						})
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_6:
				// TAST (Let): __local_var_4_5 shape=Branch(Other, EffectPure, def=Other) bindingType=(TypeApp (TypeVar m) [Unit])
				__local_var_4_5 := __t6
				_ = __local_var_4_5
				_dollar___unused_5_7 := gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Value{})
				_ = _dollar___unused_5_7
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_2_2), gopurs_runtime.Bool(result_1_1)), gopurs_runtime.Value{})
			})
			_ = __local_var_3_3
			_dollar___unused_4_8 := gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{})
			_ = _dollar___unused_4_8
			// TAST (Let): __local_var_5_9 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_5_9 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_g(100000)), gopurs_runtime.Int(0))
			_ = __local_var_5_9
			// TAST (Let): result_6_10 shape=Other bindingType=Boolean
			result_6_10 := (gopurs_runtime.RecordGet(__local_var_5_9, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_5_9, "expected").IntVal)
			_ = result_6_10
			// TAST (Let): message_7_11 shape=Other bindingType=String
			message_7_11 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_5_9, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_5_9, "actual")).StrVal())
			_ = message_7_11
			// TAST (Let): __local_var_8_13 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_8_13 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_7_11))
			_ = __local_var_8_13
			var __t15 gopurs_runtime.Value
			{
				if (result_6_10) != (true) {
					__t15 = __local_var_8_13
					goto end_branch_15
				} else {

				}
			}
			{
				if result_6_10 {
					__t15 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_15
				} else {

				}
			}
			{
				__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_15:
			_dollar___unused_9_14 := gopurs_runtime.Apply(__t15, gopurs_runtime.Value{})
			_ = _dollar___unused_9_14
			_dollar___unused_8_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_7_11), gopurs_runtime.Bool(result_6_10)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_12
			// TAST (Let): __local_var_9_16 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_9_16 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_weirdsum(0, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t17 *Constructor_Main_Just[int64]
				{
					if (x_9.IntVal) < (5) {
						__t17 = (&Constructor_Main_Just[int64]{1, gopurs_runtime.Int((2) * (x_9.IntVal))})
						goto end_branch_17
					} else {

					}
				}
				{
					__t17 = (*Constructor_Main_Just[int64])(nil)
				}
			end_branch_17:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t17)}
			}), 100000)), gopurs_runtime.Int(20))
			_ = __local_var_9_16
			// TAST (Let): result_10_18 shape=Other bindingType=Boolean
			result_10_18 := (gopurs_runtime.RecordGet(__local_var_9_16, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_9_16, "expected").IntVal)
			_ = result_10_18
			// TAST (Let): message_11_19 shape=Other bindingType=String
			message_11_19 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_9_16, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_9_16, "actual")).StrVal())
			_ = message_11_19
			// TAST (Let): __local_var_12_21 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_12_21 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_11_19))
			_ = __local_var_12_21
			var __t23 gopurs_runtime.Value
			{
				if (result_10_18) != (true) {
					__t23 = __local_var_12_21
					goto end_branch_23
				} else {

				}
			}
			{
				if result_10_18 {
					__t23 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_23
				} else {

				}
			}
			{
				__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_23:
			_dollar___unused_13_22 := gopurs_runtime.Apply(__t23, gopurs_runtime.Value{})
			_ = _dollar___unused_13_22
			_dollar___unused_12_20 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_11_19), gopurs_runtime.Bool(result_10_18)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_20
			// TAST (Let): __local_var_13_24 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_13_24 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_tricksyinners(0, 100000)), gopurs_runtime.Int(200009))
			_ = __local_var_13_24
			// TAST (Let): result_14_25 shape=Other bindingType=Boolean
			result_14_25 := (gopurs_runtime.RecordGet(__local_var_13_24, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_13_24, "expected").IntVal)
			_ = result_14_25
			// TAST (Let): message_15_26 shape=Other bindingType=String
			message_15_26 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_13_24, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_13_24, "actual")).StrVal())
			_ = message_15_26
			// TAST (Let): __local_var_16_28 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_16_28 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_15_26))
			_ = __local_var_16_28
			var __t30 gopurs_runtime.Value
			{
				if (result_14_25) != (true) {
					__t30 = __local_var_16_28
					goto end_branch_30
				} else {

				}
			}
			{
				if result_14_25 {
					__t30 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_30
				} else {

				}
			}
			{
				__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_30:
			_dollar___unused_17_29 := gopurs_runtime.Apply(__t30, gopurs_runtime.Value{})
			_ = _dollar___unused_17_29
			_dollar___unused_16_27 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_15_26), gopurs_runtime.Bool(result_14_25)), gopurs_runtime.Value{})
			_ = _dollar___unused_16_27
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_weirdsum(accum_0_loop int64, f1_1_loop gopurs_runtime.Value, n_2_loop int64) int64 {
weirdsum:
	for {
		if false {
			continue weirdsum
		}
		var accum_0 int64 = accum_0_loop
		_ = accum_0
		var f1_1 gopurs_runtime.Value = f1_1_loop
		_ = f1_1
		var n_2 int64 = n_2_loop
		_ = n_2
		var __t2 int64
		{
			if (n_2) == (0) {
				__t2 = accum_0
				goto end_branch_2
			} else {

			}
		}
		{
			// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(ADT ["Main","Maybe"] [Int])
			__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Just[int64]](gopurs_runtime.Apply(f1_1, gopurs_runtime.Int(n_2)))
			_ = __local_var_3_0
			var __t1 int64
			{
				if __local_var_3_0 != nil {
					accum_0_loop = (accum_0) + ((__local_var_3_0).V0.IntVal)
					f1_1_loop = f1_1
					n_2_loop = (n_2) - (1)
					continue weirdsum
					__t1 = gopurs_runtime.Value{}.IntVal
					goto end_branch_1
				} else {

				}
			}
			{
				accum_0_loop = accum_0
				f1_1_loop = f1_1
				n_2_loop = (n_2) - (1)
				continue weirdsum
				__t1 = gopurs_runtime.Value{}.IntVal
			}
		end_branch_1:
			__t2 = __t1
		}
	end_branch_2:
		return __t2
	}
}

func Call_Main_tricksyinners(accum_0_loop int64, x_1_loop int64) int64 {
tricksyinners:
	for {
		if false {
			continue tricksyinners
		}
		var accum_0 int64 = accum_0_loop
		_ = accum_0
		var x_1 int64 = x_1_loop
		_ = x_1
		var __t0 int64
		{
			if (x_1) == (0) {
				__t0 = (accum_0) + (((x_1) + (3)) * ((x_1) + (3)))
				goto end_branch_0
			} else {

			}
		}
		{
			accum_0_loop = (accum_0) + (2)
			x_1_loop = (x_1) - (1)
			continue tricksyinners
			__t0 = gopurs_runtime.Value{}.IntVal
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_g(x_0_loop int64) int64 {
g:
	for {
		if false {
			continue g
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var __t0 int64
		{
			if (x_0) == (0) {
				__t0 = 0
				goto end_branch_0
			} else {

			}
		}
		{
			if (x_0) == (x_0) {
				x_0_loop = (x_0) - (1)
				continue g
				__t0 = gopurs_runtime.Value{}.IntVal
				goto end_branch_0
			} else {

			}
		}
		{
			x_0_loop = (x_0) - (2)
			continue g
			__t0 = gopurs_runtime.Value{}.IntVal
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_f(x_0_loop int64) int64 {
f:
	for {
		if false {
			continue f
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var __t0 int64
		{
			if (x_0) == (0) {
				__t0 = 0
				goto end_branch_0
			} else {

			}
		}
		{
			x_0_loop = (x_0) - (1)
			continue f
			__t0 = gopurs_runtime.Value{}.IntVal
		}
	end_branch_0:
		return __t0
	}
}
