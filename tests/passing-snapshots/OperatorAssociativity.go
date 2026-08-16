package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_bug gopurs_runtime.Value
var once_Main_bug sync.Once

func Get_Main_bug() gopurs_runtime.Value {
	once_Main_bug.Do(func() {
		cache_Main_bug = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_bug(a_0_box.FloatVal(), b_1_box.FloatVal()))
		})
	})
	return cache_Main_bug
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_7.IntVal) == 1527465420) != (true))), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.Float(-1.0).FloatVal()), gopurs_runtime.Float(0.0))
			_dollar___unused_7_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_9.IntVal) == 1527465420))), gopurs_runtime.Value{})
			_ = _dollar___unused_7_8
			_dollar___unused_8_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_10
			_dollar___unused_9_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_9_11
			_dollar___unused_10_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_10_12
			_dollar___unused_11_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_11_13
			_dollar___unused_12_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_14
			_dollar___unused_13_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_13_15
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_bug(a_0_loop float64, b_1_loop float64) float64 {
	var a_0 float64 = a_0_loop
	_ = a_0
	var b_1 float64 = b_1_loop
	_ = b_1
	return -((a_0) - (b_1))
}
