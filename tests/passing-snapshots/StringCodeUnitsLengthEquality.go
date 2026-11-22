package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_lessThan gopurs_runtime.Value
var once_Main_lessThan sync.Once

func Get_Main_lessThan() gopurs_runtime.Value {
	once_Main_lessThan.Do(func() {
		cache_Main_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Main_lessThan
}

var cache_Main_eq gopurs_runtime.Value
var once_Main_eq sync.Once

func Get_Main_eq() gopurs_runtime.Value {
	once_Main_eq.Do(func() {
		cache_Main_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Main_eq
}

var cache_Main_notEq gopurs_runtime.Value
var once_Main_notEq sync.Once

func Get_Main_notEq() gopurs_runtime.Value {
	once_Main_notEq.Do(func() {
		cache_Main_notEq = Get_Data_Eq_notEq__3020190905()
	})
	return cache_Main_notEq
}

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = Get_Control_Bind_bind__3833792468()
	})
	return cache_Main_bind
}

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = Get_Control_Bind_bind__777602484()
	})
	return cache_Main_bind1
}

var cache_Main_discard gopurs_runtime.Value
var once_Main_discard sync.Once

func Get_Main_discard() gopurs_runtime.Value {
	once_Main_discard.Do(func() {
		cache_Main_discard = Get_Control_Bind_discard__143037513()
	})
	return cache_Main_discard
}

var cache_Main_assertEqual gopurs_runtime.Value
var once_Main_assertEqual sync.Once

func Get_Main_assertEqual() gopurs_runtime.Value {
	once_Main_assertEqual.Do(func() {
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__2127306453()
	})
	return cache_Main_assertEqual
}

var cache_Main_not gopurs_runtime.Value
var once_Main_not sync.Once

func Get_Main_not() gopurs_runtime.Value {
	once_Main_not.Do(func() {
		cache_Main_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Main_not
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__2010590964()
	})
	return cache_Main_bind2
}

var cache_Main_go__append gopurs_runtime.Value
var once_Main_go__append sync.Once

func Get_Main_go__append() gopurs_runtime.Value {
	once_Main_go__append.Do(func() {
		cache_Main_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Main_go__append
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__1178656404()
	})
	return cache_Main_bind3
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__691505748()
	})
	return cache_Main_bind4
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Bind_bind__3124274164()
	})
	return cache_Main_bind5
}

var cache_Main_bind6 gopurs_runtime.Value
var once_Main_bind6 sync.Once

func Get_Main_bind6() gopurs_runtime.Value {
	once_Main_bind6.Do(func() {
		cache_Main_bind6 = Get_Control_Bind_bind__850991572()
	})
	return cache_Main_bind6
}

var cache_Main_bind7 gopurs_runtime.Value
var once_Main_bind7 sync.Once

func Get_Main_bind7() gopurs_runtime.Value {
	once_Main_bind7.Do(func() {
		cache_Main_bind7 = Get_Control_Bind_bind__1022999252()
	})
	return cache_Main_bind7
}

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
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Main_check(false, "", "a"), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Main_check(true, "abc", "xyz"), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Main_check(false, "abc", "ab"), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Main_check(true, "é", "a"), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Main_check(false, "é", "ab"), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_check(true, "😀", "ab"), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_check(false, "😀", "abcd"), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_check(true, "é😀", "abc"), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_check(true, "\xed\xa0\x80", "a"), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_check(true, "\xed\xbf\xbf", "é"), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_check(true, "𝌆", "ab"), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_check(false, "\xed\xa0\x80", "ab"), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("x", Data_String_CodeUnits_FromCharArray(func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Int(int64(4096)), gopurs_runtime.StrWithCodeUnits("é", 1)).UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}())), "😀")), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_14_14, gopurs_runtime.Value{})
			_ = __local_var_15_15
			// TAST (Let): suffix_16_16 shape=App(Var) bindingType=String
			suffix_16_16 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(int64(1)), gopurs_runtime.NormalizeStringValue(__local_var_15_15)).StrVal()
			_ = suffix_16_16
			// TAST (Let): copied_17_17 shape=App(Var) bindingType=String
			copied_17_17 := Data_String_CodeUnits_FromCharArray(Data_String_CodeUnits_ToCharArray(suffix_16_16))
			_ = copied_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Main_check(true, suffix_16_16, suffix_16_16), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_check(true, suffix_16_16, copied_17_17), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Main_check(false, __local_var_15_15.StrVal(), suffix_16_16), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_shorter(suffix_16_16, __local_var_15_15.StrVal()), true}), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_shorter(__local_var_15_15.StrVal(), suffix_16_16), false}), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Main_sameLength(), gopurs_runtime.Str(suffix_16_16)), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_23_23, gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.Apply(__local_var_24_24, gopurs_runtime.Str(copied_17_17)).IntVal) != (0), true}), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.Apply(__local_var_24_24, gopurs_runtime.NormalizeStringValue(__local_var_15_15)).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_26_26
			__local_var_27_27 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), Get_Data_String_CodeUnits_length(), gopurs_runtime.Value{})
			_ = __local_var_27_27
			__local_var_28_28 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_27_27, gopurs_runtime.Value{})
			_ = __local_var_28_28
			__local_var_29_29 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_compareWith(__local_var_28_28, suffix_16_16, copied_17_17), true}), gopurs_runtime.Value{})
			_ = __local_var_29_29
			__local_var_30_30 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_compareWith(__local_var_28_28, __local_var_15_15.StrVal(), suffix_16_16), false}), gopurs_runtime.Value{})
			_ = __local_var_30_30
			__local_var_31_31 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_compareWith(gopurs_runtime.Func(func(v_31 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(int64(0))
			}), __local_var_15_15.StrVal(), suffix_16_16), true}), gopurs_runtime.Value{})
			_ = __local_var_31_31
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
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			orig := struct {
				left  string
				right string
				/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
			}{left_1, right_2}
			_ = orig
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Str(orig.left), gopurs_runtime.Str(orig.right))
		}())
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_4_1, gopurs_runtime.Value{})
		_ = __local_var_5_2
		__local_var_6_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
			actual   bool
			expected bool
			/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
		}{Call_Main_sameLength(gopurs_runtime.RecordGet(__local_var_5_2, "left").StrVal(), gopurs_runtime.RecordGet(__local_var_5_2, "right").StrVal()), expected_0}), gopurs_runtime.Value{})
		_ = __local_var_6_3
		return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
			actual   bool
			expected bool
			/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
		}{Call_Main_differentLength(gopurs_runtime.RecordGet(__local_var_5_2, "left").StrVal(), gopurs_runtime.RecordGet(__local_var_5_2, "right").StrVal()), (expected_0) != (true)}), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys2{K0: "left", K1: "right"}
