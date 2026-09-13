package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = func() gopurs_runtime.Value {
			var Call_local_Main_f__4286688220_0_0_0 func(float64) bool
			_ = Call_local_Main_f__4286688220_0_0_0
			var f__4286688220_0_0_0 gopurs_runtime.Value
			_ = f__4286688220_0_0_0
			var Call_local_Main_g_0_1_1 func(float64) bool
			_ = Call_local_Main_g_0_1_1
			var g_0_1_1 gopurs_runtime.Value
			_ = g_0_1_1
			var Call_local_Main_f_0_2_2 func(float64) bool
			_ = Call_local_Main_f_0_2_2
			var f_0_2_2 gopurs_runtime.Value
			_ = f_0_2_2
			Call_local_Main_f__4286688220_0_0_0 = func(v_1_loop float64) bool {
			f__4286688220_0_0_0:
				for {
					if false {
						continue f__4286688220_0_0_0
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
						__t3 = Call_local_Main_g_0_1_1((v_1) - (1.0))
					}
				end_branch_3:
					return __t3
				}
			}
			f__4286688220_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(Call_local_Main_f__4286688220_0_0_0(v_1_loop_val.FloatVal()))
			})
			Call_local_Main_g_0_1_1 = func(v_1_loop float64) bool {
			g_0_1_1:
				for {
					if false {
						continue g_0_1_1
					}
					var v_1 float64 = v_1_loop
					_ = v_1
					var __t4 bool
					{
						if (v_1) == (0.0) {
							__t4 = true
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = Call_local_Main_f__4286688220_0_0_0((v_1) - (1.0))
					}
				end_branch_4:
					return __t4
				}
			}
			g_0_1_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(Call_local_Main_g_0_1_1(v_1_loop_val.FloatVal()))
			})
			Call_local_Main_f_0_2_2 = func(v_1_loop float64) bool {
			f_0_2_2:
				for {
					if false {
						continue f_0_2_2
					}
					var v_1 float64 = v_1_loop
					_ = v_1
					var __t5 bool
					{
						if (v_1) == (0.0) {
							__t5 = false
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = Call_local_Main_g_0_1_1((v_1) - (1.0))
					}
				end_branch_5:
					return __t5
				}
			}
			f_0_2_2 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(Call_local_Main_f_0_2_2(v_1_loop_val.FloatVal()))
			})
			return gopurs_runtime.Bool((Call_local_Main_f__4286688220_0_0_0(1.0)) != (true))
		}()
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 string
			{
				if (Get_Main_test().IntVal) != (0) {
					__t0 = "true"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "false"
			}
		end_branch_0:
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t0)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}()
	})
	return cache_Main_main
}
