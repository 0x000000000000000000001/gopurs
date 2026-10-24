package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = func() gopurs_runtime.Value {
			var Call_local_Main_g_0_0_0 func(float64) bool
			_ = Call_local_Main_g_0_0_0
			var g_0_0_0 gopurs_runtime.Value
			_ = g_0_0_0
			var Call_local_Main_f_0_1_1 func(float64) bool
			_ = Call_local_Main_f_0_1_1
			var f_0_1_1 gopurs_runtime.Value
			_ = f_0_1_1
			Call_local_Main_g_0_0_0 = func(v_1_loop float64) bool {
			g_0_0_0:
				for {
					if false {
						continue g_0_0_0
					}
					var v_1 float64 = v_1_loop
					_ = v_1
					var __t2 bool
					{
						if (v_1) == (0.0) {
							__t2 = true
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = Call_local_Main_f_0_1_1((v_1) - (1.0))
					}
				end_branch_2:
					return __t2
				}
			}
			g_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(Call_local_Main_g_0_0_0(v_1_loop_val.FloatVal()))
			})
			Call_local_Main_f_0_1_1 = func(v_1_loop float64) bool {
			f_0_1_1:
				for {
					if false {
						continue f_0_1_1
					}
					var v_1 float64 = v_1_loop
					_ = v_1
					var __t3 bool
					{
						if (v_1) == (0.0) {
							__t3 = false
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_g_0_0_0((v_1) - (1.0))
					}
				end_branch_3:
					return __t3
				}
			}
			f_0_1_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(Call_local_Main_f_0_1_1(v_1_loop_val.FloatVal()))
			})
			return gopurs_runtime.Bool((Call_local_Main_f_0_1_1(1.0)) != (true))
		}()
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((Get_Main_test().IntVal) != (0)))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
