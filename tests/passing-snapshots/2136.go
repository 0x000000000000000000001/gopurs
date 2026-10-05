package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_bottom gopurs_runtime.Value
var once_Main_bottom sync.Once

func Get_Main_bottom() gopurs_runtime.Value {
	once_Main_bottom.Do(func() {
		cache_Main_bottom = gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)
	})
	return cache_Main_bottom
}

var cache_Main_top gopurs_runtime.Value
var once_Main_top sync.Once

func Get_Main_top() gopurs_runtime.Value {
	once_Main_top.Do(func() {
		cache_Main_top = gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal)
	})
	return cache_Main_top
}

var cache_Main_negate gopurs_runtime.Value
var once_Main_negate sync.Once

func Get_Main_negate() gopurs_runtime.Value {
	once_Main_negate.Do(func() {
		cache_Main_negate = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_negate(__local_var_0_box.IntVal))
		})
	})
	return cache_Main_negate
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_intAdd()
	})
	return cache_Main_add
}

var cache_Main_sub gopurs_runtime.Value
var once_Main_sub sync.Once

func Get_Main_sub() gopurs_runtime.Value {
	once_Main_sub.Do(func() {
		cache_Main_sub = Get_Data_Ring_intSub()
	})
	return cache_Main_sub
}

var cache_Main_mul gopurs_runtime.Value
var once_Main_mul sync.Once

func Get_Main_mul() gopurs_runtime.Value {
	once_Main_mul.Do(func() {
		cache_Main_mul = Get_Data_Semiring_intMul()
	})
	return cache_Main_mul
}

var cache_Main_div gopurs_runtime.Value
var once_Main_div sync.Once

func Get_Main_div() gopurs_runtime.Value {
	once_Main_div.Do(func() {
		cache_Main_div = Get_Data_EuclideanRing_intDiv()
	})
	return cache_Main_div
}

var cache_Main_mod gopurs_runtime.Value
var once_Main_mod sync.Once

func Get_Main_mod() gopurs_runtime.Value {
	once_Main_mod.Do(func() {
		cache_Main_mod = Get_Data_EuclideanRing_intMod()
	})
	return cache_Main_mod
}

var cache_Main_literalBottom gopurs_runtime.Value
var once_Main_literalBottom sync.Once

func Get_Main_literalBottom() gopurs_runtime.Value {
	once_Main_literalBottom.Do(func() {
		cache_Main_literalBottom = gopurs_runtime.Int(int64(-2147483648))
	})
	return cache_Main_literalBottom
}

var cache_Main_foldedSubtract gopurs_runtime.Value
var once_Main_foldedSubtract sync.Once

func Get_Main_foldedSubtract() gopurs_runtime.Value {
	once_Main_foldedSubtract.Do(func() {
		cache_Main_foldedSubtract = gopurs_runtime.Int(gopurs_runtime.IntSub(int64(-2147483648), int64(1)))
	})
	return cache_Main_foldedSubtract
}

var cache_Main_foldedShift gopurs_runtime.Value
var once_Main_foldedShift sync.Once

func Get_Main_foldedShift() gopurs_runtime.Value {
	once_Main_foldedShift.Do(func() {
		cache_Main_foldedShift = gopurs_runtime.Int(int64(1))
	})
	return cache_Main_foldedShift
}

var cache_Main_foldedNegate gopurs_runtime.Value
var once_Main_foldedNegate sync.Once

func Get_Main_foldedNegate() gopurs_runtime.Value {
	once_Main_foldedNegate.Do(func() {
		cache_Main_foldedNegate = gopurs_runtime.Int(int64(-2147483648))
	})
	return cache_Main_foldedNegate
}

var cache_Main_foldedMultiply gopurs_runtime.Value
var once_Main_foldedMultiply sync.Once

func Get_Main_foldedMultiply() gopurs_runtime.Value {
	once_Main_foldedMultiply.Do(func() {
		cache_Main_foldedMultiply = gopurs_runtime.Int(gopurs_runtime.IntMul(int64(2147483647), int64(2147483647)))
	})
	return cache_Main_foldedMultiply
}

var cache_Main_foldedAdd gopurs_runtime.Value
var once_Main_foldedAdd sync.Once

func Get_Main_foldedAdd() gopurs_runtime.Value {
	once_Main_foldedAdd.Do(func() {
		cache_Main_foldedAdd = gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(2147483647), int64(1)))
	})
	return cache_Main_foldedAdd
}

var cache_Main_check2 gopurs_runtime.Value
var once_Main_check2 sync.Once

func Get_Main_check2() gopurs_runtime.Value {
	once_Main_check2.Do(func() {
		cache_Main_check2 = gopurs_runtime.Func5(func(label_0_box gopurs_runtime.Value, op_1_box gopurs_runtime.Value, left_2_box gopurs_runtime.Value, right_3_box gopurs_runtime.Value, expected_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check2(label_0_box.StrVal(), op_1_box, left_2_box.IntVal, right_3_box.IntVal, expected_4_box.IntVal)
		})
	})
	return cache_Main_check2
}

var cache_Main_check1 gopurs_runtime.Value
var once_Main_check1 sync.Once

func Get_Main_check1() gopurs_runtime.Value {
	once_Main_check1.Do(func() {
		cache_Main_check1 = gopurs_runtime.Func4(func(label_0_box gopurs_runtime.Value, op_1_box gopurs_runtime.Value, value_2_box gopurs_runtime.Value, expected_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check1(label_0_box.StrVal(), op_1_box, value_2_box.IntVal, expected_3_box.IntVal)
		})
	})
	return cache_Main_check1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("2136: negating bottom does not exceed top"), gopurs_runtime.Bool(((gopurs_runtime.IntNegate(Get_Data_Bounded_bottomInt().IntVal)) > (Get_Data_Bounded_topInt().IntVal)) != (true)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("folded negation wraps"), gopurs_runtime.Bool((int64(-2147483648)) == (Get_Data_Bounded_bottomInt().IntVal))), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("folded addition wraps"), gopurs_runtime.Bool((Get_Main_foldedAdd().IntVal) == (Get_Data_Bounded_bottomInt().IntVal))), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("folded subtraction wraps"), gopurs_runtime.Bool((Get_Main_foldedSubtract().IntVal) == (Get_Data_Bounded_topInt().IntVal))), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiplication follows the JS Number then Int conversion"), gopurs_runtime.Bool((Get_Main_foldedMultiply().IntVal) == (int64(0)))), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("folded shift masks the count"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_check1("runtime negation of bottom", gopurs_runtime.Func(func(__local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntNegate(__local_var_7.IntVal))
			}), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_check1("runtime negation of top", gopurs_runtime.Func(func(__local_var_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntNegate(__local_var_8.IntVal))
			}), Get_Data_Bounded_topInt().IntVal, int64(-2147483647)), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_check1("runtime complement", Get_Data_Int_Bits_complement(), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_topInt().IntVal), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_check2("runtime addition", Get_Data_Semiring_intAdd(), Get_Data_Bounded_topInt().IntVal, int64(1), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_check2("runtime subtraction", Get_Data_Ring_intSub(), Get_Data_Bounded_bottomInt().IntVal, int64(1), Get_Data_Bounded_topInt().IntVal), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_check2("runtime multiplication", Get_Data_Semiring_intMul(), Get_Data_Bounded_topInt().IntVal, int64(2), int64(-2)), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_check2("runtime multiplication with Number rounding", Get_Data_Semiring_intMul(), Get_Data_Bounded_topInt().IntVal, Get_Data_Bounded_topInt().IntVal, int64(0)), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Main_check2("runtime shift overflow", Get_Data_Int_Bits_shl(), int64(1), int64(31), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Main_check2("runtime shift count 32", Get_Data_Int_Bits_shl(), int64(1), int64(32), int64(1)), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Main_check2("runtime negative shift count", Get_Data_Int_Bits_shl(), int64(1), int64(-1), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Main_check2("runtime arithmetic shift", Get_Data_Int_Bits_shr(), Get_Data_Bounded_bottomInt().IntVal, int64(32), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Main_check2("runtime unsigned shift", Get_Data_Int_Bits_zshr(), Get_Data_Bounded_bottomInt().IntVal, int64(1), int64(1073741824)), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_check2("runtime unsigned shift with count 32", Get_Data_Int_Bits_zshr(), Get_Data_Bounded_bottomInt().IntVal, int64(32), gopurs_runtime.Zshr(gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Main_check2("runtime Euclidean division boundary", Get_Data_EuclideanRing_intDiv(), Get_Data_Bounded_bottomInt().IntVal, int64(-1), gopurs_runtime.Zshr(gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Main_check2("runtime modulo boundary", Get_Data_EuclideanRing_intMod(), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_topInt().IntVal, int64(2147483646)), gopurs_runtime.Value{})
			_ = __local_var_21_21
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_negate(__local_var_0_loop int64) int64 {
	var __local_var_0 int64 = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.IntNegate(__local_var_0)
}

func Call_Main_check2(label_0_loop string, op_1_loop gopurs_runtime.Value, left_2_loop int64, right_3_loop int64, expected_4_loop int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var op_1 gopurs_runtime.Value = op_1_loop
	_ = op_1
	var left_2 int64 = left_2_loop
	_ = left_2
	var right_3 int64 = right_3_loop
	_ = right_3
	var expected_4 int64 = expected_4_loop
	_ = expected_4
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_5_0 shape=App(Var) bindingType=Any
		__local_var_5_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(left_2))
		_ = __local_var_5_0
		__local_var_6_1 := gopurs_runtime.Apply(__local_var_5_0, gopurs_runtime.Value{})
		_ = __local_var_6_1
		__local_var_7_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(right_3)), gopurs_runtime.Value{})
		_ = __local_var_7_2
		__local_var_8_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_1), gopurs_runtime.Value{})
		_ = __local_var_8_3
		__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_7_2), gopurs_runtime.Value{})
		_ = __local_var_9_4
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(label_0), gopurs_runtime.Bool((gopurs_runtime.Apply2(op_1, gopurs_runtime.Int(__local_var_8_3.IntVal), gopurs_runtime.Int(__local_var_9_4.IntVal)).IntVal) == (expected_4))), gopurs_runtime.Value{})
	})
}

func Call_Main_check1(label_0_loop string, op_1_loop gopurs_runtime.Value, value_2_loop int64, expected_3_loop int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var op_1 gopurs_runtime.Value = op_1_loop
	_ = op_1
	var value_2 int64 = value_2_loop
	_ = value_2
	var expected_3 int64 = expected_3_loop
	_ = expected_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=Any
		__local_var_4_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(value_2))
		_ = __local_var_4_0
		__local_var_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
		_ = __local_var_5_1
		__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_1), gopurs_runtime.Value{})
		_ = __local_var_6_2
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(label_0), gopurs_runtime.Bool((gopurs_runtime.Apply(op_1, gopurs_runtime.Int(__local_var_6_2.IntVal)).IntVal) == (expected_3))), gopurs_runtime.Value{})
	})
}
