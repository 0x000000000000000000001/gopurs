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
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(3.0)).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(2.0)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(-1.0)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.Float(-1.0).FloatVal())).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(0.5)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			var __t8 string
			{
				var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(1.0), gopurs_runtime.Float(2.0))
				if (gopurs_runtime.Bool((uint32(__t_tag_7.IntVal) == 380165415)).IntVal) != (0) {
					__t8 = "true"
					goto end_branch_8
				} else {

				}
			}
			{
				__t8 = "false"
			}
		end_branch_8:
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t8)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			var __t11 string
			{
				var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(1.0), gopurs_runtime.Float(2.0))
				if (gopurs_runtime.Bool((uint32(__t_tag_10.IntVal) == 1527465420)).IntVal) != (0) {
					__t11 = "true"
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = "false"
			}
		end_branch_11:
			_dollar___unused_7_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t11)), gopurs_runtime.Value{})
			_ = _dollar___unused_7_9
			var __t14 string
			{
				var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(1.0), gopurs_runtime.Float(2.0))
				if (gopurs_runtime.Bool((uint32(__t_tag_13.IntVal) == 380165415) != (true)).IntVal) != (0) {
					__t14 = "true"
					goto end_branch_14
				} else {

				}
			}
			{
				__t14 = "false"
			}
		end_branch_14:
			_dollar___unused_8_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t14)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_12
			var __t17 string
			{
				var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(1.0), gopurs_runtime.Float(2.0))
				if (gopurs_runtime.Bool((uint32(__t_tag_16.IntVal) == 1527465420) != (true)).IntVal) != (0) {
					__t17 = "true"
					goto end_branch_17
				} else {

				}
			}
			{
				__t17 = "false"
			}
		end_branch_17:
			_dollar___unused_9_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t17)), gopurs_runtime.Value{})
			_ = _dollar___unused_9_15
			_dollar___unused_10_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = _dollar___unused_10_18
			_dollar___unused_11_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = _dollar___unused_11_19
			_dollar___unused_12_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = _dollar___unused_12_20
			_dollar___unused_13_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = _dollar___unused_13_21
			_dollar___unused_14_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = _dollar___unused_14_22
			_dollar___unused_15_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = _dollar___unused_15_23
			_dollar___unused_16_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = _dollar___unused_16_24
			_dollar___unused_17_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("foobar")).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_17_25
			_dollar___unused_18_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true")), gopurs_runtime.Value{})
			_ = _dollar___unused_18_26
			_dollar___unused_19_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = _dollar___unused_19_27
			_dollar___unused_20_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("false")), gopurs_runtime.Value{})
			_ = _dollar___unused_20_28
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}
