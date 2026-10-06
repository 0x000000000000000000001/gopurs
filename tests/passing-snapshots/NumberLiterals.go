package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	math "math"
	sync "sync"
)

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func2(func(str_0_box gopurs_runtime.Value, num_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(str_0_box.StrVal(), num_1_box.FloatVal())
		})
	})
	return cache_Main_check
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(str_0_box gopurs_runtime.Value, num_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(str_0_box.StrVal(), num_1_box.FloatVal())
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Main_test("0.17", 0.17)
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Main_test("0.25996181067141905", 0.25996181067141905), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Main_test("0.3572019862807257", 0.3572019862807257), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Main_test("0.46817723004874223", 0.46817723004874223), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Main_test("0.9640035681058178", 0.9640035681058178), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Main_test("4.23808622486133", 4.23808622486133), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_test("4.540362294799751", 4.540362294799751), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_test("5.212384849884261", 5.212384849884261), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_test("13.958257048123212", 13.958257048123212), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_test("32.96176575630599", 32.96176575630599), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_test("38.47735512322269", 38.47735512322269), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_test("10000000000.0", 10000000000.0), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_test("10000000000.0", 10000000000.0), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Main_test("0.00001", 0.00001), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Main_test("0.00001", 0.00001), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Main_test("1.5339794352098402e-118", 1.5339794352098402e-118), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Main_test("2.108934760892056e-59", 2.108934760892056e-59), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Main_test("2.250634744599241e-19", 2.250634744599241e-19), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_test("5.960464477539063e-8", 5.960464477539063e-8), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Main_test("5e-324", 5e-324), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Main_test("5e-324", 5e-324), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(Call_Main_test("0.0", 0.0), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(Call_Main_test("0.0", gopurs_runtime.NegativeZero()), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(Call_Main_test("-5e-324", -5e-324), gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply(Call_Main_test("1e-7", 1e-7), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply(Call_Main_test("-1e-7", -1e-7), gopurs_runtime.Value{})
			_ = __local_var_26_26
			__local_var_27_27 := gopurs_runtime.Apply(Call_Main_test("1e-8", 1e-8), gopurs_runtime.Value{})
			_ = __local_var_27_27
			__local_var_28_28 := gopurs_runtime.Apply(Call_Main_test("1e-9", 1e-9), gopurs_runtime.Value{})
			_ = __local_var_28_28
			__local_var_29_29 := gopurs_runtime.Apply(Call_Main_test("0.000001", 0.000001), gopurs_runtime.Value{})
			_ = __local_var_29_29
			__local_var_30_30 := gopurs_runtime.Apply(Call_Main_test("9.999999999999997e-7", 9.999999999999997e-7), gopurs_runtime.Value{})
			_ = __local_var_30_30
			__local_var_31_31 := gopurs_runtime.Apply(Call_Main_test("100000000000000000000.0", 100000000000000000000.0), gopurs_runtime.Value{})
			_ = __local_var_31_31
			__local_var_32_32 := gopurs_runtime.Apply(Call_Main_test("999999999999999900000.0", 999999999999999900000.0), gopurs_runtime.Value{})
			_ = __local_var_32_32
			__local_var_33_33 := gopurs_runtime.Apply(Call_Main_test("1e+21", 1e+21), gopurs_runtime.Value{})
			_ = __local_var_33_33
			__local_var_34_34 := gopurs_runtime.Apply(Call_Main_test("-1e+21", -1e+21), gopurs_runtime.Value{})
			_ = __local_var_34_34
			__local_var_35_35 := gopurs_runtime.Apply(Call_Main_test("9007199254740991.0", 9007199254740991.0), gopurs_runtime.Value{})
			_ = __local_var_35_35
			__local_var_36_36 := gopurs_runtime.Apply(Call_Main_test("1000000000000000100.0", 1000000000000000100.0), gopurs_runtime.Value{})
			_ = __local_var_36_36
			__local_var_37_37 := gopurs_runtime.Apply(Call_Main_test("2.2250738585072014e-308", 2.2250738585072014e-308), gopurs_runtime.Value{})
			_ = __local_var_37_37
			__local_var_38_38 := gopurs_runtime.Apply(Call_Main_test("1.7976931348623157e+308", 1.7976931348623157e+308), gopurs_runtime.Value{})
			_ = __local_var_38_38
			__local_var_39_39 := gopurs_runtime.Apply(Call_Main_test("Infinity", math.Inf(1)), gopurs_runtime.Value{})
			_ = __local_var_39_39
			__local_var_40_40 := gopurs_runtime.Apply(Call_Main_test("-Infinity", math.Inf(-1)), gopurs_runtime.Value{})
			_ = __local_var_40_40
			__local_var_41_41 := gopurs_runtime.Apply(Call_Main_test("NaN", math.NaN()), gopurs_runtime.Value{})
			_ = __local_var_41_41
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_check(str_0_loop string, num_1_loop float64) gopurs_runtime.Value {
	var str_0 string = str_0_loop
	_ = str_0
	var num_1 float64 = num_1_loop
	_ = num_1
	var __t0 gopurs_runtime.Value
	{
		if (Data_Show_ShowNumberImpl(num_1)) == (str_0) {
			__t0 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return Get_Data_Unit_unit()
			})
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("Expected ", Data_Show_ShowStringImpl(str_0)), ", got "), Data_Show_ShowStringImpl(Data_Show_ShowNumberImpl(num_1))), ".")), gopurs_runtime.Bool(false))
	}
end_branch_0:
	return __t0
}

func Call_Main_test(str_0_loop string, num_1_loop float64) gopurs_runtime.Value {
	var str_0 string = str_0_loop
	_ = str_0
	var num_1 float64 = num_1_loop
	_ = num_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := Call_Main_check(str_0, num_1)
		_ = __local_var_2_0
		__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
		_ = __local_var_3_1
		__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Float(num_1)), gopurs_runtime.Value{})
		_ = __local_var_4_2
		__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_2), gopurs_runtime.Value{})
		_ = __local_var_5_3
		return gopurs_runtime.Apply(Call_Main_check(str_0, __local_var_5_3.FloatVal()), gopurs_runtime.Value{})
	})
}
