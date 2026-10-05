package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_surrogatePair gopurs_runtime.Value
var once_Main_surrogatePair sync.Once

func Get_Main_surrogatePair() gopurs_runtime.Value {
	once_Main_surrogatePair.Do(func() {
		cache_Main_surrogatePair = gopurs_runtime.Bool(true)
	})
	return cache_Main_surrogatePair
}

var cache_Main_singleCharacter gopurs_runtime.Value
var once_Main_singleCharacter sync.Once

func Get_Main_singleCharacter() gopurs_runtime.Value {
	once_Main_singleCharacter.Do(func() {
		cache_Main_singleCharacter = gopurs_runtime.Bool(true)
	})
	return cache_Main_singleCharacter
}

var cache_Main_replacement gopurs_runtime.Value
var once_Main_replacement sync.Once

func Get_Main_replacement() gopurs_runtime.Value {
	once_Main_replacement.Do(func() {
		cache_Main_replacement = gopurs_runtime.Str("�")
	})
	return cache_Main_replacement
}

var cache_Main_lowSurrogate gopurs_runtime.Value
var once_Main_lowSurrogate sync.Once

func Get_Main_lowSurrogate() gopurs_runtime.Value {
	once_Main_lowSurrogate.Do(func() {
		cache_Main_lowSurrogate = gopurs_runtime.Str("\xed\xbc\x86")
	})
	return cache_Main_lowSurrogate
}

var cache_Main_highSurrogate gopurs_runtime.Value
var once_Main_highSurrogate sync.Once

func Get_Main_highSurrogate() gopurs_runtime.Value {
	once_Main_highSurrogate.Do(func() {
		cache_Main_highSurrogate = gopurs_runtime.Str("\xed\xa0\xb4")
	})
	return cache_Main_highSurrogate
}

var cache_Main_loneSurrogates gopurs_runtime.Value
var once_Main_loneSurrogates sync.Once

func Get_Main_loneSurrogates() gopurs_runtime.Value {
	once_Main_loneSurrogates.Do(func() {
		cache_Main_loneSurrogates = gopurs_runtime.Bool(true)
	})
	return cache_Main_loneSurrogates
}

var cache_Main_notReplacing gopurs_runtime.Value
var once_Main_notReplacing sync.Once

func Get_Main_notReplacing() gopurs_runtime.Value {
	once_Main_notReplacing.Do(func() {
		cache_Main_notReplacing = gopurs_runtime.Bool(true)
	})
	return cache_Main_notReplacing
}

var cache_Main_outOfOrderSurrogates gopurs_runtime.Value
var once_Main_outOfOrderSurrogates sync.Once

func Get_Main_outOfOrderSurrogates() gopurs_runtime.Value {
	once_Main_outOfOrderSurrogates.Do(func() {
		cache_Main_outOfOrderSurrogates = gopurs_runtime.Bool(true)
	})
	return cache_Main_outOfOrderSurrogates
}

var cache_Main_hex gopurs_runtime.Value
var once_Main_hex sync.Once

func Get_Main_hex() gopurs_runtime.Value {
	once_Main_hex.Do(func() {
		cache_Main_hex = gopurs_runtime.Bool(true)
	})
	return cache_Main_hex
}

var cache_Main_checkConcat gopurs_runtime.Value
var once_Main_checkConcat sync.Once

func Get_Main_checkConcat() gopurs_runtime.Value {
	once_Main_checkConcat.Do(func() {
		cache_Main_checkConcat = gopurs_runtime.Func3(func(left_0_box gopurs_runtime.Value, right_1_box gopurs_runtime.Value, expected_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkConcat(left_0_box.StrVal(), right_1_box.StrVal(), expected_2_box.StrVal())
		})
	})
	return cache_Main_checkConcat
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("single-character escape sequences"), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("hex escape sequences"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("astral code points are represented as a UTF-16 surrogate pair"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("lone surrogates may be combined into a surrogate pair"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("lone surrogates may be combined out of order to remain lone surrogates"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("lone surrogates are not replaced with the Unicode replacement character U+FFFD"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_checkConcat("\xed\xa0\xb4", "\xed\xbc\x86", "𝌆"), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_checkConcat("é\xed\xa0\xb4", "\xed\xbc\x86z", "é𝌆z"), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_checkConcat("\xed\xbc\x86", "\xed\xa0\xb4", "\xed\xbc\x86\xed\xa0\xb4"), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_checkConcat("\xed\xa0\xb4", "\xed\xa0\xb4", "\xed\xa0\xb4\xed\xa0\xb4"), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_checkConcat("\xed\xa0\x80", "\xed\xb0\x80", "𐀀"), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_checkConcat("\xed\xaf\xbf", "\xed\xbf\xbf", "􏿿"), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_checkConcat("\xed\xa0\xb4", "", "\xed\xa0\xb4"), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Main_checkConcat("", "\xed\xbc\x86", "\xed\xbc\x86"), gopurs_runtime.Value{})
			_ = __local_var_14_14
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_checkConcat(left_0_loop string, right_1_loop string, expected_2_loop string) gopurs_runtime.Value {
	var left_0 string = left_0_loop
	_ = left_0
	var right_1 string = right_1_loop
	_ = right_1
	var expected_2 string = expected_2_loop
	_ = expected_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(left_0))
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(right_1)), gopurs_runtime.Value{})
		_ = __local_var_5_2
		__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
		_ = __local_var_6_3
		__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
		_ = __local_var_7_4
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("runtime UTF-16 concatenation"), gopurs_runtime.Bool((gopurs_runtime.ConcatString(__local_var_6_3.StrVal(), __local_var_7_4.StrVal())) == (expected_2))), gopurs_runtime.Value{})
	})
}
