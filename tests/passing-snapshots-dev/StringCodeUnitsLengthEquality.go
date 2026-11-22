package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_shorter gopurs_runtime.Value
var once_Main_shorter sync.Once

func Get_Main_shorter() gopurs_runtime.Value {
	once_Main_shorter.Do(func() {
		cache_Main_shorter = gopurs_runtime.Func2(func(left_0_box gopurs_runtime.Value, right_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_shorter(left_0_box.StrVal(), right_1_box.StrVal()))
		})
	})
	return cache_Main_shorter
}

var cache_Main_sameLength gopurs_runtime.Value
var once_Main_sameLength sync.Once

func Get_Main_sameLength() gopurs_runtime.Value {
	once_Main_sameLength.Do(func() {
		cache_Main_sameLength = gopurs_runtime.Func2(func(left_0_box gopurs_runtime.Value, right_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_sameLength(left_0_box.StrVal(), right_1_box.StrVal()))
		})
	})
	return cache_Main_sameLength
}

var cache_Main_differentLength gopurs_runtime.Value
var once_Main_differentLength sync.Once

func Get_Main_differentLength() gopurs_runtime.Value {
	once_Main_differentLength.Do(func() {
		cache_Main_differentLength = gopurs_runtime.Func2(func(left_0_box gopurs_runtime.Value, right_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_differentLength(left_0_box.StrVal(), right_1_box.StrVal()))
		})
	})
	return cache_Main_differentLength
}

var cache_Main_compareWith gopurs_runtime.Value
var once_Main_compareWith sync.Once

func Get_Main_compareWith() gopurs_runtime.Value {
	once_Main_compareWith.Do(func() {
		cache_Main_compareWith = gopurs_runtime.Func3(func(measure_0_box gopurs_runtime.Value, left_1_box gopurs_runtime.Value, right_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_compareWith(measure_0_box, left_1_box.StrVal(), right_2_box.StrVal()))
		})
	})
	return cache_Main_compareWith
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func3(func(expected_0_box gopurs_runtime.Value, left_1_box gopurs_runtime.Value, right_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check((expected_0_box.IntVal) != (0), left_1_box.StrVal(), right_2_box.StrVal())
		})
	})
	return cache_Main_check
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := Call_Main_check(true, "", "")
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Main_check(false, "", "a"), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(Call_Main_check(true, "abc", "xyz"), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(Call_Main_check(false, "abc", "ab"), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(Call_Main_check(true, "é", "a"), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(Call_Main_check(false, "é", "ab"), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(Call_Main_check(true, "😀", "ab"), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(Call_Main_check(false, "😀", "abcd"), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(Call_Main_check(true, "é😀", "abc"), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(Call_Main_check(true, "\xed\xa0\x80", "a"), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(Call_Main_check(true, "\xed\xbf\xbf", "é"), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply(Call_Main_check(true, "𝌆", "ab"), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply(Call_Main_check(false, "\xed\xa0\x80", "ab"), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			reference_14_14 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("x", Data_String_CodeUnits_FromCharArray(func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_replicate(int64(4096), gopurs_runtime.StrWithCodeUnits("é", 1))).UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}())), "😀")), gopurs_runtime.Value{})
			_ = reference_14_14
			source_15_15 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), reference_14_14, gopurs_runtime.Value{})
			_ = source_15_15
			// TAST (Let): suffix_16_16 shape=App(Var) bindingType=String
			suffix_16_16 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(int64(1)), gopurs_runtime.NormalizeStringValue(source_15_15)).StrVal()
			_ = suffix_16_16
			// TAST (Let): copied_17_17 shape=App(Var) bindingType=String
			copied_17_17 := Data_String_CodeUnits_FromCharArray(Data_String_CodeUnits_ToCharArray(suffix_16_16))
			_ = copied_17_17
			_dollar___unused_18_18 := gopurs_runtime.Apply(Call_Main_check(true, suffix_16_16, suffix_16_16), gopurs_runtime.Value{})
			_ = _dollar___unused_18_18
			_dollar___unused_19_19 := gopurs_runtime.Apply(Call_Main_check(true, suffix_16_16, copied_17_17), gopurs_runtime.Value{})
			_ = _dollar___unused_19_19
			_dollar___unused_20_20 := gopurs_runtime.Apply(Call_Main_check(false, source_15_15.StrVal(), suffix_16_16), gopurs_runtime.Value{})
			_ = _dollar___unused_20_20
			_dollar___unused_21_21 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_shorter(suffix_16_16, source_15_15.StrVal()), true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_21_21
			_dollar___unused_22_22 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_shorter(source_15_15.StrVal(), suffix_16_16), false}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_22_22
			partialRef_23_23 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Main_sameLength(), gopurs_runtime.Str(suffix_16_16)), gopurs_runtime.Value{})
			_ = partialRef_23_23
			partial_24_24 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), partialRef_23_23, gopurs_runtime.Value{})
			_ = partial_24_24
			_dollar___unused_25_25 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.Apply(partial_24_24, gopurs_runtime.Str(copied_17_17)).IntVal) != (0), true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_25_25
			_dollar___unused_26_26 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.Apply(partial_24_24, gopurs_runtime.NormalizeStringValue(source_15_15)).IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_26_26
			measureRef_27_27 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), Get_Data_String_CodeUnits_length(), gopurs_runtime.Value{})
			_ = measureRef_27_27
			measure_28_28 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), measureRef_27_27, gopurs_runtime.Value{})
			_ = measure_28_28
			_dollar___unused_29_29 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_compareWith(measure_28_28, suffix_16_16, copied_17_17), true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_29_29
			_dollar___unused_30_30 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_compareWith(measure_28_28, source_15_15.StrVal(), suffix_16_16), false}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_30_30
			_dollar___unused_31_31 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_compareWith(gopurs_runtime.Func(func(v_31 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(int64(0))
				}), source_15_15.StrVal(), suffix_16_16), true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_31_31
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_shorter(left_0_loop string, right_1_loop string) bool {
	var left_0 string = left_0_loop
	_ = left_0
	var right_1 string = right_1_loop
	_ = right_1
	return (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(left_0)).IntVal) < (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(right_1)).IntVal)
}

func Call_Main_sameLength(left_0_loop string, right_1_loop string) bool {
	var left_0 string = left_0_loop
	_ = left_0
	var right_1 string = right_1_loop
	_ = right_1
	return gopurs_runtime.StringCodeUnitsLengthsEqual(gopurs_runtime.Str(left_0), gopurs_runtime.Str(right_1))
}

func Call_Main_differentLength(left_0_loop string, right_1_loop string) bool {
	var left_0 string = left_0_loop
	_ = left_0
	var right_1 string = right_1_loop
	_ = right_1
	return (gopurs_runtime.StringCodeUnitsLengthsEqual(gopurs_runtime.Str(left_0), gopurs_runtime.Str(right_1))) != (true)
}

func Call_Main_compareWith(measure_0_loop gopurs_runtime.Value, left_1_loop string, right_2_loop string) bool {
	var measure_0 gopurs_runtime.Value = measure_0_loop
	_ = measure_0
	var left_1 string = left_1_loop
	_ = left_1
	var right_2 string = right_2_loop
	_ = right_2
	return (gopurs_runtime.Apply(measure_0, gopurs_runtime.Str(left_1)).IntVal) == (gopurs_runtime.Apply(measure_0, gopurs_runtime.Str(right_2)).IntVal)
}

func Call_Main_check(expected_0_loop bool, left_1_loop string, right_2_loop string) gopurs_runtime.Value {
	var expected_0 bool = expected_0_loop
	_ = expected_0
	var left_1 string = left_1_loop
	_ = left_1
	var right_2 string = right_2_loop
	_ = right_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
			orig := struct {
				left  string
				right string
				/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
			}{left_1, right_2}
			_ = orig
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.left), gopurs_runtime.Str(orig.right))
		}())
		_ = __local_var_3_0
		reference_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = reference_4_1
		input_5_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), reference_4_1, gopurs_runtime.Value{})
		_ = input_5_2
		_dollar___unused_6_3 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_sameLength(gopurs_runtime.RecordGet(input_5_2, "left").StrVal(), gopurs_runtime.RecordGet(input_5_2, "right").StrVal()), expected_0}
			_ = orig
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
		}(), gopurs_runtime.Value{})
		_ = _dollar___unused_6_3
		return gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_differentLength(gopurs_runtime.RecordGet(input_5_2, "left").StrVal(), gopurs_runtime.RecordGet(input_5_2, "right").StrVal()), (expected_0) != (true)}
			_ = orig
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
		}(), gopurs_runtime.Value{})
	})
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

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
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

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys2{K0: "left", K1: "right"}
