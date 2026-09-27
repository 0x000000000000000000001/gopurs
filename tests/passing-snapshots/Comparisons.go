package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_checkOrdering gopurs_runtime.Value
var once_Main_checkOrdering sync.Once

func Get_Main_checkOrdering() gopurs_runtime.Value {
	once_Main_checkOrdering.Do(func() {
		cache_Main_checkOrdering = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Main_checkOrdering
}

var cache_Main_checkOrdering__39504560 gopurs_runtime.Value
var once_Main_checkOrdering__39504560 sync.Once

func Get_Main_checkOrdering__39504560() gopurs_runtime.Value {
	once_Main_checkOrdering__39504560.Do(func() {
		cache_Main_checkOrdering__39504560 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__39504560(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__39504560
}

var cache_Main_checkOrdering__4119744823 gopurs_runtime.Value
var once_Main_checkOrdering__4119744823 sync.Once

func Get_Main_checkOrdering__4119744823() gopurs_runtime.Value {
	once_Main_checkOrdering__4119744823.Do(func() {
		cache_Main_checkOrdering__4119744823 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__4119744823(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__4119744823
}

var cache_Main_checkOrdering__2650964060 gopurs_runtime.Value
var once_Main_checkOrdering__2650964060 sync.Once

func Get_Main_checkOrdering__2650964060() gopurs_runtime.Value {
	once_Main_checkOrdering__2650964060.Do(func() {
		cache_Main_checkOrdering__2650964060 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__2650964060(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__2650964060
}

var cache_Main_checkOrdering__325519965 gopurs_runtime.Value
var once_Main_checkOrdering__325519965 sync.Once

func Get_Main_checkOrdering__325519965() gopurs_runtime.Value {
	once_Main_checkOrdering__325519965.Do(func() {
		cache_Main_checkOrdering__325519965 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__325519965(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__325519965
}

var cache_Main_checkOrdering__2503009178 gopurs_runtime.Value
var once_Main_checkOrdering__2503009178 sync.Once

func Get_Main_checkOrdering__2503009178() gopurs_runtime.Value {
	once_Main_checkOrdering__2503009178.Do(func() {
		cache_Main_checkOrdering__2503009178 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__2503009178(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__2503009178
}

var cache_Main_checkOrdering__3515330801 gopurs_runtime.Value
var once_Main_checkOrdering__3515330801 sync.Once

func Get_Main_checkOrdering__3515330801() gopurs_runtime.Value {
	once_Main_checkOrdering__3515330801.Do(func() {
		cache_Main_checkOrdering__3515330801 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__3515330801(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__3515330801
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var __t_tag_1 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 1.0, 2.0)
			_ = __t_tag_1
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_1.IntVal) == 1527465420)))
			_ = __local_var_0_0
			__local_var_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_2
			__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_2_3
			var __t_tag_5 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 3.0, 1.0)
			_ = __t_tag_5
			__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_5.IntVal) == 380165415))), gopurs_runtime.Value{})
			_ = __local_var_3_4
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("a"), gopurs_runtime.Str("b"))
			_ = __t_tag_7
			__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_7.IntVal) == 1527465420))), gopurs_runtime.Value{})
			_ = __local_var_4_6
			__local_var_5_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_5_8
			var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("z"), gopurs_runtime.Str("a"))
			_ = __t_tag_10
			__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_10.IntVal) == 380165415))), gopurs_runtime.Value{})
			_ = __local_var_6_9
			__local_var_7_11 := gopurs_runtime.Apply(Call_Main_checkOrdering__2650964060("a", "b", 1527465420), gopurs_runtime.Value{})
			_ = __local_var_7_11
			__local_var_8_12 := gopurs_runtime.Apply(Call_Main_checkOrdering__39504560("é", "é", 902936544), gopurs_runtime.Value{})
			_ = __local_var_8_12
			__local_var_9_13 := gopurs_runtime.Apply(Call_Main_checkOrdering__4119744823("Ā", "ÿ", 380165415), gopurs_runtime.Value{})
			_ = __local_var_9_13
			__local_var_10_14 := gopurs_runtime.Apply(Call_Main_checkOrdering__2650964060("ÿ", "Ā", 1527465420), gopurs_runtime.Value{})
			_ = __local_var_10_14
			__local_var_11_15 := gopurs_runtime.Apply(Call_Main_checkOrdering__3515330801("a", "b", 1527465420), gopurs_runtime.Value{})
			_ = __local_var_11_15
			__local_var_12_16 := gopurs_runtime.Apply(Call_Main_checkOrdering__325519965("é", "é", 902936544), gopurs_runtime.Value{})
			_ = __local_var_12_16
			__local_var_13_17 := gopurs_runtime.Apply(Call_Main_checkOrdering__2503009178("éa", "é", 380165415), gopurs_runtime.Value{})
			_ = __local_var_13_17
			__local_var_14_18 := gopurs_runtime.Apply(Call_Main_checkOrdering__3515330801("é", "éa", 1527465420), gopurs_runtime.Value{})
			_ = __local_var_14_18
			__local_var_15_19 := gopurs_runtime.Apply(Call_Main_checkOrdering__3515330801("😀", "😁", 1527465420), gopurs_runtime.Value{})
			_ = __local_var_15_19
			__local_var_16_20 := gopurs_runtime.Apply(Call_Main_checkOrdering__2503009178("😁", "😀", 380165415), gopurs_runtime.Value{})
			_ = __local_var_16_20
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_checkOrdering(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): Eq0_1_0 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeVar a$scope1)])
	Eq0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(dictOrd_0.V0, gopurs_runtime.Value{}))
	_ = Eq0_1_0
	return gopurs_runtime.Func3(func(left_2 gopurs_runtime.Value, right_3 gopurs_runtime.Value, expected_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=Any
			__local_var_5_1 := gopurs_runtime.Apply(Get_Effect_Ref__new(), left_2)
			_ = __local_var_5_1
			leftRef_6_2 := gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Value{})
			_ = leftRef_6_2
			rightRef_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), right_3), gopurs_runtime.Value{})
			_ = rightRef_7_3
			x_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), leftRef_6_2), gopurs_runtime.Value{})
			_ = x_8_4
			y_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), rightRef_7_3), gopurs_runtime.Value{})
			_ = y_9_5
			__local_var_10_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Value{})
			_ = __local_var_10_6
			var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_8
			var __t_tag_9 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_9
			__local_var_11_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_8.IntVal) == 1527465420), (uint32(__t_tag_9) == 1527465420)}), gopurs_runtime.Value{})
			_ = __local_var_11_7
			var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_11
			var __t_tag_12 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_12
			__local_var_12_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_11.IntVal) == 380165415) != (true), (uint32(__t_tag_12) == 380165415) != (true)}), gopurs_runtime.Value{})
			_ = __local_var_12_10
			var __t_tag_14 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_14
			__local_var_13_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply2(Eq0_1_0.V0, x_8_4, y_9_5).IntVal) != (0), (uint32(__t_tag_14) == 902936544)}), gopurs_runtime.Value{})
			_ = __local_var_13_13
			var __t_tag_16 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_16
			__local_var_14_15 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{((gopurs_runtime.Apply2(Eq0_1_0.V0, x_8_4, y_9_5).IntVal) != (0)) != (true), (uint32(__t_tag_16) == 902936544) != (true)}), gopurs_runtime.Value{})
			_ = __local_var_14_15
			var __t_tag_18 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_18
			var __t_tag_19 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_19
			__local_var_15_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_18.IntVal) == 1527465420) != (true), (uint32(__t_tag_19) == 1527465420) != (true)}), gopurs_runtime.Value{})
			_ = __local_var_15_17
			var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_21
			var __t_tag_22 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_22
			__local_var_16_20 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_21.IntVal) == 380165415), (uint32(__t_tag_22) == 380165415)}), gopurs_runtime.Value{})
			_ = __local_var_16_20
			comparatorRef_17_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), dictOrd_0.V1), gopurs_runtime.Value{})
			_ = comparatorRef_17_23
			comparator_18_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), comparatorRef_17_23), gopurs_runtime.Value{})
			_ = comparator_18_24
			__local_var_19_25 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(comparator_18_24, x_8_4, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Value{})
			_ = __local_var_19_25
			partialRef_20_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(comparator_18_24, x_8_4)), gopurs_runtime.Value{})
			_ = partialRef_20_26
			partial_21_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), partialRef_20_26), gopurs_runtime.Value{})
			_ = partial_21_27
			__local_var_22_28 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(partial_21_27, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Value{})
			_ = __local_var_22_28
			__local_var_23_29 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(partial_21_27, x_8_4).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_23_29
			__local_var_24_30 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(partial_21_27, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Value{})
			_ = __local_var_24_30
			lessThanRef_25_31 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func(func(a2_25 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_32 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, a2_25)
				_ = __t_tag_32
				return gopurs_runtime.Bool((uint32(__t_tag_32.IntVal) == 1527465420))
			})), gopurs_runtime.Value{})
			_ = lessThanRef_25_31
			lessThan_26_33 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), lessThanRef_25_31), gopurs_runtime.Value{})
			_ = lessThan_26_33
			var __t_tag_35 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_35
			__local_var_27_34 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(lessThan_26_33, y_9_5).IntVal) != (0), (uint32(__t_tag_35) == 1527465420)}), gopurs_runtime.Value{})
			_ = __local_var_27_34
			__local_var_28_36 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(lessThan_26_33, x_8_4).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_28_36
			var __t_tag_37 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_37
			return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(lessThan_26_33, y_9_5).IntVal) != (0), (uint32(__t_tag_37) == 1527465420)}), gopurs_runtime.Value{})
		})
	})
}

func Call_Main_checkOrdering__39504560(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__39504560:
	for {
		if false {
			continue checkOrdering__39504560
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			__local_var_8_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_8_5
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_7
			__local_var_9_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_7.IntVal) == 1527465420), false}), gopurs_runtime.Value{})
			_ = __local_var_9_6
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_9
			__local_var_10_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_9.IntVal) == 380165415) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_10_8
			__local_var_11_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), true}), gopurs_runtime.Value{})
			_ = __local_var_11_10
			__local_var_12_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), false}), gopurs_runtime.Value{})
			_ = __local_var_12_11
			var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_13
			__local_var_13_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_13.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_13_12
			var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_15
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_15.IntVal) == 380165415), false}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2320923292()), gopurs_runtime.Value{})
			_ = __local_var_15_16
			__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_16), gopurs_runtime.Value{})
			_ = __local_var_16_17
			__local_var_17_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_17_18
			__local_var_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_18_19
			__local_var_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_19), gopurs_runtime.Value{})
			_ = __local_var_19_20
			__local_var_20_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_20_21
			__local_var_21_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_21_22
			__local_var_22_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_22_23
			__local_var_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__3343604442(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_23_24
			__local_var_24_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_23_24), gopurs_runtime.Value{})
			_ = __local_var_24_25
			__local_var_25_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_25_26
			__local_var_26_27 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_26_27
			return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__4119744823(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__4119744823:
	for {
		if false {
			continue checkOrdering__4119744823
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			__local_var_8_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_8_5
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_7
			__local_var_9_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_7.IntVal) == 1527465420), false}), gopurs_runtime.Value{})
			_ = __local_var_9_6
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_9
			__local_var_10_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_9.IntVal) == 380165415) != (true), false}), gopurs_runtime.Value{})
			_ = __local_var_10_8
			__local_var_11_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Value{})
			_ = __local_var_11_10
			__local_var_12_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_12_11
			var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_13
			__local_var_13_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_13.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_13_12
			var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_15
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_15.IntVal) == 380165415), true}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2320923292()), gopurs_runtime.Value{})
			_ = __local_var_15_16
			__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_16), gopurs_runtime.Value{})
			_ = __local_var_16_17
			__local_var_17_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_17_18
			__local_var_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_18_19
			__local_var_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_19), gopurs_runtime.Value{})
			_ = __local_var_19_20
			__local_var_20_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_20_21
			__local_var_21_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_21_22
			__local_var_22_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_22_23
			__local_var_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__3343604442(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_23_24
			__local_var_24_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_23_24), gopurs_runtime.Value{})
			_ = __local_var_24_25
			__local_var_25_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_25_26
			__local_var_26_27 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_26_27
			return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__2650964060(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__2650964060:
	for {
		if false {
			continue checkOrdering__2650964060
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			__local_var_8_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_8_5
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_7
			__local_var_9_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_7.IntVal) == 1527465420), true}), gopurs_runtime.Value{})
			_ = __local_var_9_6
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_9
			__local_var_10_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_9.IntVal) == 380165415) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_10_8
			__local_var_11_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Value{})
			_ = __local_var_11_10
			__local_var_12_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_12_11
			var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_13
			__local_var_13_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_13.IntVal) == 1527465420) != (true), false}), gopurs_runtime.Value{})
			_ = __local_var_13_12
			var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_15
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_15.IntVal) == 380165415), false}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2320923292()), gopurs_runtime.Value{})
			_ = __local_var_15_16
			__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_16), gopurs_runtime.Value{})
			_ = __local_var_16_17
			__local_var_17_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_17_18
			__local_var_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_18_19
			__local_var_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_19), gopurs_runtime.Value{})
			_ = __local_var_19_20
			__local_var_20_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_20_21
			__local_var_21_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_21_22
			__local_var_22_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_22_23
			__local_var_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__3343604442(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_23_24
			__local_var_24_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_23_24), gopurs_runtime.Value{})
			_ = __local_var_24_25
			__local_var_25_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true}), gopurs_runtime.Value{})
			_ = __local_var_25_26
			__local_var_26_27 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_26_27
			return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true}), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__325519965(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__325519965:
	for {
		if false {
			continue checkOrdering__325519965
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			__local_var_8_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_8_5
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_7
			__local_var_9_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_7.IntVal) == 1527465420), false}), gopurs_runtime.Value{})
			_ = __local_var_9_6
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_9
			__local_var_10_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_9.IntVal) == 380165415) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_10_8
			__local_var_11_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), true}), gopurs_runtime.Value{})
			_ = __local_var_11_10
			__local_var_12_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), false}), gopurs_runtime.Value{})
			_ = __local_var_12_11
			var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_13
			__local_var_13_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_13.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_13_12
			var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_15
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_15.IntVal) == 380165415), false}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2340924753()), gopurs_runtime.Value{})
			_ = __local_var_15_16
			__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_16), gopurs_runtime.Value{})
			_ = __local_var_16_17
			__local_var_17_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_17_18
			__local_var_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_18_19
			__local_var_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_19), gopurs_runtime.Value{})
			_ = __local_var_19_20
			__local_var_20_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_20_21
			__local_var_21_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_21_22
			__local_var_22_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_22_23
			__local_var_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__2428873527(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_23_24
			__local_var_24_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_23_24), gopurs_runtime.Value{})
			_ = __local_var_24_25
			__local_var_25_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_25_26
			__local_var_26_27 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_26_27
			return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__2503009178(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__2503009178:
	for {
		if false {
			continue checkOrdering__2503009178
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			__local_var_8_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_8_5
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_7
			__local_var_9_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_7.IntVal) == 1527465420), false}), gopurs_runtime.Value{})
			_ = __local_var_9_6
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_9
			__local_var_10_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_9.IntVal) == 380165415) != (true), false}), gopurs_runtime.Value{})
			_ = __local_var_10_8
			__local_var_11_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Value{})
			_ = __local_var_11_10
			__local_var_12_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_12_11
			var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_13
			__local_var_13_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_13.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_13_12
			var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_15
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_15.IntVal) == 380165415), true}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2340924753()), gopurs_runtime.Value{})
			_ = __local_var_15_16
			__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_16), gopurs_runtime.Value{})
			_ = __local_var_16_17
			__local_var_17_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_17_18
			__local_var_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_18_19
			__local_var_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_19), gopurs_runtime.Value{})
			_ = __local_var_19_20
			__local_var_20_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_20_21
			__local_var_21_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_21_22
			__local_var_22_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Value{})
			_ = __local_var_22_23
			__local_var_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__2428873527(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_23_24
			__local_var_24_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_23_24), gopurs_runtime.Value{})
			_ = __local_var_24_25
			__local_var_25_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_25_26
			__local_var_26_27 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_26_27
			return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__3515330801(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__3515330801:
	for {
		if false {
			continue checkOrdering__3515330801
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			__local_var_8_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_8_5
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_7
			__local_var_9_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_7.IntVal) == 1527465420), true}), gopurs_runtime.Value{})
			_ = __local_var_9_6
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_9
			__local_var_10_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_9.IntVal) == 380165415) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_10_8
			__local_var_11_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Value{})
			_ = __local_var_11_10
			__local_var_12_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Value{})
			_ = __local_var_12_11
			var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_13
			__local_var_13_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_13.IntVal) == 1527465420) != (true), false}), gopurs_runtime.Value{})
			_ = __local_var_13_12
			var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
			_ = __t_tag_15
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(uint32(__t_tag_15.IntVal) == 380165415), false}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2340924753()), gopurs_runtime.Value{})
			_ = __local_var_15_16
			__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_16), gopurs_runtime.Value{})
			_ = __local_var_16_17
			__local_var_17_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_17_18
			__local_var_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_16_17, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_18_19
			__local_var_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_19), gopurs_runtime.Value{})
			_ = __local_var_19_20
			__local_var_20_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_20_21
			__local_var_21_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Value{})
			_ = __local_var_21_22
			__local_var_22_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply(__local_var_19_20, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Value{})
			_ = __local_var_22_23
			__local_var_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__2428873527(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_23_24
			__local_var_24_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_23_24), gopurs_runtime.Value{})
			_ = __local_var_24_25
			__local_var_25_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true}), gopurs_runtime.Value{})
			_ = __local_var_25_26
			__local_var_26_27 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_26_27
			return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(gopurs_runtime.Apply(__local_var_24_25, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true}), gopurs_runtime.Value{})
		})
	}
}
