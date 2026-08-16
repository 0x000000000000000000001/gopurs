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
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just)(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, value0}))}
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
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_f(99999)), gopurs_runtime.Int(0))
			_ = __local_var_0_0
			// TAST (Let): result_1_1 -> bool
			result_1_1 := (gopurs_runtime.RecordGet(__local_var_0_0, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_0_0, "expected").IntVal)
			_ = result_1_1
			// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
			__local_var_2_2 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): message_2_3 -> string
				message_2_3 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "actual")).StrVal())
				_ = message_2_3
				// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
				__local_var_3_4 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_2_3))
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
				// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
				__local_var_4_5 := __t6
				_ = __local_var_4_5
				_dollar___unused_5_7 := gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Value{})
				_ = _dollar___unused_5_7
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_2_3), gopurs_runtime.Bool(result_1_1)), gopurs_runtime.Value{})
			})
			_ = __local_var_2_2
			_dollar___unused_3_8 := gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Value{})
			_ = _dollar___unused_3_8
			// TAST (Let): __local_var_4_9 -> gopurs_runtime.Value
			__local_var_4_9 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_g(100000)), gopurs_runtime.Int(0))
			_ = __local_var_4_9
			// TAST (Let): result_5_10 -> bool
			result_5_10 := (gopurs_runtime.RecordGet(__local_var_4_9, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_4_9, "expected").IntVal)
			_ = result_5_10
			// TAST (Let): message_6_12 -> string
			message_6_12 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_4_9, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_4_9, "actual")).StrVal())
			_ = message_6_12
			// TAST (Let): __local_var_7_13 -> gopurs_runtime.Value
			__local_var_7_13 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_6_12))
			_ = __local_var_7_13
			var __t15 gopurs_runtime.Value
			{
				if (result_5_10) != (true) {
					__t15 = __local_var_7_13
					goto end_branch_15
				} else {

				}
			}
			{
				if result_5_10 {
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
			_dollar___unused_8_14 := gopurs_runtime.Apply(__t15, gopurs_runtime.Value{})
			_ = _dollar___unused_8_14
			_dollar___unused_6_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_6_12), gopurs_runtime.Bool(result_5_10)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_11
			// TAST (Let): __local_var_7_16 -> gopurs_runtime.Value
			__local_var_7_16 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_weirdsum(0, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t17 *Constructor_Main_Just
				{
					if (x_7.IntVal) < (5) {
						__t17 = (&Constructor_Main_Just{1, gopurs_runtime.Int((2) * (x_7.IntVal))})
						goto end_branch_17
					} else {

					}
				}
				{
					__t17 = (*Constructor_Main_Just)(nil)
				}
			end_branch_17:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t17)}
			}), 100000)), gopurs_runtime.Int(20))
			_ = __local_var_7_16
			// TAST (Let): result_8_18 -> bool
			result_8_18 := (gopurs_runtime.RecordGet(__local_var_7_16, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_7_16, "expected").IntVal)
			_ = result_8_18
			// TAST (Let): message_9_20 -> string
			message_9_20 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_7_16, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_7_16, "actual")).StrVal())
			_ = message_9_20
			// TAST (Let): __local_var_10_21 -> gopurs_runtime.Value
			__local_var_10_21 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_9_20))
			_ = __local_var_10_21
			var __t23 gopurs_runtime.Value
			{
				if (result_8_18) != (true) {
					__t23 = __local_var_10_21
					goto end_branch_23
				} else {

				}
			}
			{
				if result_8_18 {
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
			_dollar___unused_11_22 := gopurs_runtime.Apply(__t23, gopurs_runtime.Value{})
			_ = _dollar___unused_11_22
			_dollar___unused_9_19 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_9_20), gopurs_runtime.Bool(result_8_18)), gopurs_runtime.Value{})
			_ = _dollar___unused_9_19
			// TAST (Let): __local_var_10_24 -> gopurs_runtime.Value
			__local_var_10_24 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_tricksyinners(0, 100000)), gopurs_runtime.Int(200009))
			_ = __local_var_10_24
			// TAST (Let): result_11_25 -> bool
			result_11_25 := (gopurs_runtime.RecordGet(__local_var_10_24, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_10_24, "expected").IntVal)
			_ = result_11_25
			// TAST (Let): message_12_27 -> string
			message_12_27 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_10_24, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_10_24, "actual")).StrVal())
			_ = message_12_27
			// TAST (Let): __local_var_13_28 -> gopurs_runtime.Value
			__local_var_13_28 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_12_27))
			_ = __local_var_13_28
			var __t30 gopurs_runtime.Value
			{
				if (result_11_25) != (true) {
					__t30 = __local_var_13_28
					goto end_branch_30
				} else {

				}
			}
			{
				if result_11_25 {
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
			_dollar___unused_14_29 := gopurs_runtime.Apply(__t30, gopurs_runtime.Value{})
			_ = _dollar___unused_14_29
			_dollar___unused_12_26 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_12_27), gopurs_runtime.Bool(result_11_25)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_26
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Nothing struct {
	Rc uint32
}

type Constructor_Main_Just struct {
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
			// TAST (Let): __local_var_3_0 -> *Constructor_Main_Just
			__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Apply(f1_1, gopurs_runtime.Int(n_2)))
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
