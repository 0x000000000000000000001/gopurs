package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func(func(expected_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(expected_0_box.StrVal())
		})
	})
	return cache_Main_check
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Main_check("")
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Main_check("quotes\" slash\\"), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Main_check("\x00\x08\x09\x0a\x0d\x1f\x7f"), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Main_check("é漢€💻𝄞  "), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Main_check("𝌆"), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Main_check("\xed\xa0\x80"), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_check("\xed\xbf\xbf"), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_check("a\xed\xa0\x80z"), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_check("\xed\xa0\x80𐀀\xed\xbf\xbf"), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_check("math.Abs(1); unsafe.Pointer(nil); sync.Once{}; gopurs_runtime.Value{}"), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_check("/* sync.Once */ // unsafe.Pointer\x0a\"math.Abs\""), gopurs_runtime.Value{})
			_ = __local_var_11_11
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_check(expected_0_loop string) gopurs_runtime.Value {
	var expected_0 string = expected_0_loop
	_ = expected_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(expected_0))
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = __local_var_2_1
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_3_2
		return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___4176622598("", struct {
			actual   string
			expected string
		}{__local_var_3_2.StrVal(), expected_0}), gopurs_runtime.Value{})
	})
}
