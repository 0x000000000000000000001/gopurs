package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowNumberImpl(3.0))), gopurs_runtime.Value{})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowNumberImpl(2.0))), gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowNumberImpl(-1.0))), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowNumberImpl(-1.0))), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowNumberImpl(0.5))), gopurs_runtime.Value{})
			_ = __local_var_4_4
			var __t7 string
			{
				var __t_tag_6 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 1.0, 2.0)
				_ = __t_tag_6
				if uint32(__t_tag_6.IntVal) == 380165415 {
					__t7 = "true"
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = "false"
			}
		end_branch_7:
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t7)), gopurs_runtime.Value{})
			_ = __local_var_5_5
			var __t10 string
			{
				var __t_tag_9 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 1.0, 2.0)
				_ = __t_tag_9
				if uint32(__t_tag_9.IntVal) == 1527465420 {
					__t10 = "true"
					goto end_branch_10
				} else {

				}
			}
			{
				__t10 = "false"
			}
		end_branch_10:
			__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t10)), gopurs_runtime.Value{})
			_ = __local_var_6_8
			var __t13 string
			{
				var __t_tag_12 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 1.0, 2.0)
				_ = __t_tag_12
				if (uint32(__t_tag_12.IntVal) == 380165415) != (true) {
					__t13 = "true"
					goto end_branch_13
				} else {

				}
			}
			{
				__t13 = "false"
			}
		end_branch_13:
			__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t13)), gopurs_runtime.Value{})
			_ = __local_var_7_11
			var __t16 string
			{
				var __t_tag_15 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 1.0, 2.0)
				_ = __t_tag_15
				if (uint32(__t_tag_15.IntVal) == 1527465420) != (true) {
					__t16 = "true"
					goto end_branch_16
				} else {

				}
			}
			{
				__t16 = "false"
			}
		end_branch_16:
			__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t16)), gopurs_runtime.Value{})
			_ = __local_var_8_14
			__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = __local_var_9_17
			__local_var_10_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = __local_var_10_18
			__local_var_11_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = __local_var_11_19
			__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = __local_var_12_20
			__local_var_13_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = __local_var_13_21
			__local_var_14_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = __local_var_14_22
			__local_var_15_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = __local_var_15_23
			__local_var_16_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowStringImpl("foobar"))), gopurs_runtime.Value{})
			_ = __local_var_16_24
			__local_var_17_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = __local_var_17_25
			__local_var_18_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = __local_var_18_26
			__local_var_19_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = __local_var_19_27
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}
