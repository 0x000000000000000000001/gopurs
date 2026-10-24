package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
			_dollar___unused_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_2
			_dollar___unused_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_3
			var __t_tag_5 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 3.0, 1.0)
			_ = __t_tag_5
			_dollar___unused_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_5.IntVal) == 380165415))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_4
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("a"), gopurs_runtime.Str("b"))
			_ = __t_tag_7
			_dollar___unused_4_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_7.IntVal) == 1527465420))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_6
			_dollar___unused_5_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_5_8
			var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("z"), gopurs_runtime.Str("a"))
			_ = __t_tag_10
			_dollar___unused_6_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_10.IntVal) == 380165415))), gopurs_runtime.Value{})
			_ = _dollar___unused_6_9
			_dollar___unused_7_11 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar())))), gopurs_runtime.Str("a"), gopurs_runtime.Str("b"), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_7_11
			_dollar___unused_8_12 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar())))), gopurs_runtime.Str("é"), gopurs_runtime.Str("é"), gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_8_12
			_dollar___unused_9_13 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar())))), gopurs_runtime.Str("Ā"), gopurs_runtime.Str("ÿ"), gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_9_13
			_dollar___unused_10_14 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar())))), gopurs_runtime.Str("ÿ"), gopurs_runtime.Str("Ā"), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_10_14
			_dollar___unused_11_15 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())))), gopurs_runtime.Str("a"), gopurs_runtime.Str("b"), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_11_15
			_dollar___unused_12_16 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())))), gopurs_runtime.Str("é"), gopurs_runtime.Str("é"), gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_12_16
			_dollar___unused_13_17 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())))), gopurs_runtime.Str("éa"), gopurs_runtime.Str("é"), gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_13_17
			_dollar___unused_14_18 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())))), gopurs_runtime.Str("é"), gopurs_runtime.Str("éa"), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_14_18
			_dollar___unused_15_19 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())))), gopurs_runtime.Str("😀"), gopurs_runtime.Str("😁"), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_15_19
			_dollar___unused_16_20 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Main_checkOrdering(Rebox_Main_2406510097_4177771502(Rebox_Main_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())))), gopurs_runtime.Str("😁"), gopurs_runtime.Str("😀"), gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}), gopurs_runtime.Value{})
			_ = _dollar___unused_16_20
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
			__local_var_5_1 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), left_2)
			_ = __local_var_5_1
			leftRef_6_2 := gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Value{})
			_ = leftRef_6_2
			rightRef_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), right_3), gopurs_runtime.Value{})
			_ = rightRef_7_3
			x_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), leftRef_6_2), gopurs_runtime.Value{})
			_ = x_8_4
			y_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), rightRef_7_3), gopurs_runtime.Value{})
			_ = y_9_5
			_dollar___unused_10_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Ordering_eqOrdering()))), Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Ordering_showOrdering())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   uint32
					expected uint32
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5).IntVal), uint32(expected_4.IntVal)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.actual), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.expected), UnsafePtr: nil})
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_10_6
			var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_8
			var __t_tag_9 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_9
			_dollar___unused_11_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(uint32(__t_tag_8.IntVal) == 1527465420), (uint32(__t_tag_9) == 1527465420)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_11_7
			var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_11
			var __t_tag_12 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_12
			_dollar___unused_12_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(uint32(__t_tag_11.IntVal) == 380165415) != (true), (uint32(__t_tag_12) == 380165415) != (true)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_12_10
			var __t_tag_14 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_14
			_dollar___unused_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.Apply2(Eq0_1_0.V0, x_8_4, y_9_5).IntVal) != (0), (uint32(__t_tag_14) == 902936544)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			var __t_tag_16 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_16
			_dollar___unused_14_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{((gopurs_runtime.Apply2(Eq0_1_0.V0, x_8_4, y_9_5).IntVal) != (0)) != (true), (uint32(__t_tag_16) == 902936544) != (true)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_14_15
			var __t_tag_18 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_18
			var __t_tag_19 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_19
			_dollar___unused_15_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(uint32(__t_tag_18.IntVal) == 1527465420) != (true), (uint32(__t_tag_19) == 1527465420) != (true)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_15_17
			var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5)
			_ = __t_tag_21
			var __t_tag_22 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_22
			_dollar___unused_16_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(uint32(__t_tag_21.IntVal) == 380165415), (uint32(__t_tag_22) == 380165415)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_16_20
			comparatorRef_17_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), dictOrd_0.V1), gopurs_runtime.Value{})
			_ = comparatorRef_17_23
			comparator_18_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), comparatorRef_17_23), gopurs_runtime.Value{})
			_ = comparator_18_24
			_dollar___unused_19_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Ordering_eqOrdering()))), Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Ordering_showOrdering())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   uint32
					expected uint32
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{uint32(gopurs_runtime.Apply2(comparator_18_24, x_8_4, y_9_5).IntVal), uint32(expected_4.IntVal)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.actual), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.expected), UnsafePtr: nil})
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_19_25
			partialRef_20_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(comparator_18_24, x_8_4)), gopurs_runtime.Value{})
			_ = partialRef_20_26
			partial_21_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), partialRef_20_26), gopurs_runtime.Value{})
			_ = partial_21_27
			_dollar___unused_22_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Ordering_eqOrdering()))), Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Ordering_showOrdering())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   uint32
					expected uint32
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{uint32(gopurs_runtime.Apply(partial_21_27, y_9_5).IntVal), uint32(expected_4.IntVal)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.actual), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.expected), UnsafePtr: nil})
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_22_28
			_dollar___unused_23_29 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Ordering_eqOrdering()))), Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Ordering_showOrdering())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   uint32
					expected uint32
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{uint32(gopurs_runtime.Apply(partial_21_27, x_8_4).IntVal), 902936544}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.actual), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.expected), UnsafePtr: nil})
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_23_29
			_dollar___unused_24_30 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Ordering_eqOrdering()))), Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Ordering_showOrdering())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   uint32
					expected uint32
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{uint32(gopurs_runtime.Apply(partial_21_27, y_9_5).IntVal), uint32(expected_4.IntVal)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.actual), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.expected), UnsafePtr: nil})
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_24_30
			lessThanRef_25_31 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Func(func(a2_25 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_32 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, a2_25)
				_ = __t_tag_32
				return gopurs_runtime.Bool((uint32(__t_tag_32.IntVal) == 1527465420))
			})), gopurs_runtime.Value{})
			_ = lessThanRef_25_31
			lessThan_26_33 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), lessThanRef_25_31), gopurs_runtime.Value{})
			_ = lessThan_26_33
			var __t_tag_35 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_35
			_dollar___unused_27_34 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.Apply(lessThan_26_33, y_9_5).IntVal) != (0), (uint32(__t_tag_35) == 1527465420)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_27_34
			_dollar___unused_28_36 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.Apply(lessThan_26_33, x_8_4).IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_28_36
			var __t_tag_37 uint32 = uint32(expected_4.IntVal)
			_ = __t_tag_37
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.Apply(lessThan_26_33, y_9_5).IntVal) != (0), (uint32(__t_tag_37) == 1527465420)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
		})
	})
}

func Rebox_Main_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1209612131(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[uint32])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[uint32])(unsafe.Pointer(in))
}

func Rebox_Main_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}
