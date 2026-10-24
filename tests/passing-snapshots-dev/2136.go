package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

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
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("2136: negating bottom does not exceed top"), gopurs_runtime.Bool(((gopurs_runtime.IntNegate(Get_Data_Bounded_bottomInt().IntVal)) > (Get_Data_Bounded_topInt().IntVal)) != (true)))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("folded negation wraps"), gopurs_runtime.Bool((int64(-2147483648)) == (Get_Data_Bounded_bottomInt().IntVal))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("folded addition wraps"), gopurs_runtime.Bool((Get_Main_foldedAdd().IntVal) == (Get_Data_Bounded_bottomInt().IntVal))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("folded subtraction wraps"), gopurs_runtime.Bool((Get_Main_foldedSubtract().IntVal) == (Get_Data_Bounded_topInt().IntVal))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("multiplication follows the JS Number then Int conversion"), gopurs_runtime.Bool((Get_Main_foldedMultiply().IntVal) == (int64(0)))), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("folded shift masks the count"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(Call_Main_check1("runtime negation of bottom", gopurs_runtime.Func(func(__local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntNegate(__local_var_7.IntVal))
			}), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(Call_Main_check1("runtime negation of top", gopurs_runtime.Func(func(__local_var_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntNegate(__local_var_8.IntVal))
			}), Get_Data_Bounded_topInt().IntVal, int64(-2147483647)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(Call_Main_check1("runtime complement", Get_Data_Int_Bits_complement(), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_topInt().IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(Call_Main_check2("runtime addition", Get_Data_Semiring_intAdd(), Get_Data_Bounded_topInt().IntVal, int64(1), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(Call_Main_check2("runtime subtraction", Get_Data_Ring_intSub(), Get_Data_Bounded_bottomInt().IntVal, int64(1), Get_Data_Bounded_topInt().IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply(Call_Main_check2("runtime multiplication", Get_Data_Semiring_intMul(), Get_Data_Bounded_topInt().IntVal, int64(2), int64(-2)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply(Call_Main_check2("runtime multiplication with Number rounding", Get_Data_Semiring_intMul(), Get_Data_Bounded_topInt().IntVal, Get_Data_Bounded_topInt().IntVal, int64(0)), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			_dollar___unused_14_14 := gopurs_runtime.Apply(Call_Main_check2("runtime shift overflow", Get_Data_Int_Bits_shl(), int64(1), int64(31), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_14_14
			_dollar___unused_15_15 := gopurs_runtime.Apply(Call_Main_check2("runtime shift count 32", Get_Data_Int_Bits_shl(), int64(1), int64(32), int64(1)), gopurs_runtime.Value{})
			_ = _dollar___unused_15_15
			_dollar___unused_16_16 := gopurs_runtime.Apply(Call_Main_check2("runtime negative shift count", Get_Data_Int_Bits_shl(), int64(1), int64(-1), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_16_16
			_dollar___unused_17_17 := gopurs_runtime.Apply(Call_Main_check2("runtime arithmetic shift", Get_Data_Int_Bits_shr(), Get_Data_Bounded_bottomInt().IntVal, int64(32), Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_17_17
			_dollar___unused_18_18 := gopurs_runtime.Apply(Call_Main_check2("runtime unsigned shift", Get_Data_Int_Bits_zshr(), Get_Data_Bounded_bottomInt().IntVal, int64(1), int64(1073741824)), gopurs_runtime.Value{})
			_ = _dollar___unused_18_18
			_dollar___unused_19_19 := gopurs_runtime.Apply(Call_Main_check2("runtime unsigned shift with count 32", Get_Data_Int_Bits_zshr(), Get_Data_Bounded_bottomInt().IntVal, int64(32), gopurs_runtime.Zshr(gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_19_19
			_dollar___unused_20_20 := gopurs_runtime.Apply(Call_Main_check2("runtime Euclidean division boundary", Get_Data_EuclideanRing_intDiv(), Get_Data_Bounded_bottomInt().IntVal, int64(-1), gopurs_runtime.Zshr(gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_20_20
			_dollar___unused_21_21 := gopurs_runtime.Apply(Call_Main_check2("runtime modulo boundary", Get_Data_EuclideanRing_intMod(), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_topInt().IntVal, int64(2147483646)), gopurs_runtime.Value{})
			_ = _dollar___unused_21_21
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
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
		__local_var_5_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(left_2))
		_ = __local_var_5_0
		leftRef_6_1 := gopurs_runtime.Apply(__local_var_5_0, gopurs_runtime.Value{})
		_ = leftRef_6_1
		rightRef_7_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(right_3)), gopurs_runtime.Value{})
		_ = rightRef_7_2
		a_8_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), leftRef_6_1), gopurs_runtime.Value{})
		_ = a_8_3
		b_9_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), rightRef_7_2), gopurs_runtime.Value{})
		_ = b_9_4
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str(label_0), gopurs_runtime.Bool((gopurs_runtime.Apply2(op_1, gopurs_runtime.Int(a_8_3.IntVal), gopurs_runtime.Int(b_9_4.IntVal)).IntVal) == (expected_4))), gopurs_runtime.Value{})
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
		__local_var_4_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(value_2))
		_ = __local_var_4_0
		ref_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
		_ = ref_5_1
		a_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), ref_5_1), gopurs_runtime.Value{})
		_ = a_6_2
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str(label_0), gopurs_runtime.Bool((gopurs_runtime.Apply(op_1, gopurs_runtime.Int(a_6_2.IntVal)).IntVal) == (expected_3))), gopurs_runtime.Value{})
	})
}
