package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = Get_Effect_pureE()
	})
	return cache_Main_pure
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1682951303, UnsafePtr: unsafe.Pointer((&Constructor_Main_Y{1, value0.IntVal, value1.StrVal(), (value2.IntVal) != (0)}))}
				})
			})
		})
	})
	return cache_Main_Y
}

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_X(x_0_box)
		})
	})
	return cache_Main_X
}

var cache_Main_Nil gopurs_runtime.Value
var once_Main_Nil sync.Once

func Get_Main_Nil() gopurs_runtime.Value {
	once_Main_Nil.Do(func() {
		cache_Main_Nil = gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nil
}

var cache_Main_Cons gopurs_runtime.Value
var once_Main_Cons sync.Once

func Get_Main_Cons() gopurs_runtime.Value {
	once_Main_Cons.Do(func() {
		cache_Main_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Main_Cons
}

var cache_Main_patternWithParens gopurs_runtime.Value
var once_Main_patternWithParens sync.Once

func Get_Main_patternWithParens() gopurs_runtime.Value {
	once_Main_patternWithParens.Do(func() {
		cache_Main_patternWithParens = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternWithParens
}

var cache_Main_patternWithNamedBinder gopurs_runtime.Value
var once_Main_patternWithNamedBinder sync.Once

func Get_Main_patternWithNamedBinder() gopurs_runtime.Value {
	once_Main_patternWithNamedBinder.Do(func() {
		cache_Main_patternWithNamedBinder = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternWithNamedBinder
}

var cache_Main_patternSimple gopurs_runtime.Value
var once_Main_patternSimple sync.Once

func Get_Main_patternSimple() gopurs_runtime.Value {
	once_Main_patternSimple.Do(func() {
		cache_Main_patternSimple = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternSimple
}

var cache_Main_patternNewtype gopurs_runtime.Value
var once_Main_patternNewtype sync.Once

func Get_Main_patternNewtype() gopurs_runtime.Value {
	once_Main_patternNewtype.Do(func() {
		cache_Main_patternNewtype = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternNewtype
}

var cache_Main_patternMultipleWithNormal gopurs_runtime.Value
var once_Main_patternMultipleWithNormal sync.Once

func Get_Main_patternMultipleWithNormal() gopurs_runtime.Value {
	once_Main_patternMultipleWithNormal.Do(func() {
		cache_Main_patternMultipleWithNormal = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternMultipleWithNormal
}

var cache_Main_patternMultiple gopurs_runtime.Value
var once_Main_patternMultiple sync.Once

func Get_Main_patternMultiple() gopurs_runtime.Value {
	once_Main_patternMultiple.Do(func() {
		cache_Main_patternMultiple = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternMultiple
}

var cache_Main_patternDoWithParens gopurs_runtime.Value
var once_Main_patternDoWithParens sync.Once

func Get_Main_patternDoWithParens() gopurs_runtime.Value {
	once_Main_patternDoWithParens.Do(func() {
		cache_Main_patternDoWithParens = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoWithParens
}

var cache_Main_patternDoWithNamedBinder gopurs_runtime.Value
var once_Main_patternDoWithNamedBinder sync.Once

func Get_Main_patternDoWithNamedBinder() gopurs_runtime.Value {
	once_Main_patternDoWithNamedBinder.Do(func() {
		cache_Main_patternDoWithNamedBinder = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoWithNamedBinder
}

var cache_Main_patternDoSimple gopurs_runtime.Value
var once_Main_patternDoSimple sync.Once

func Get_Main_patternDoSimple() gopurs_runtime.Value {
	once_Main_patternDoSimple.Do(func() {
		cache_Main_patternDoSimple = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoSimple
}

var cache_Main_patternDoNewtype gopurs_runtime.Value
var once_Main_patternDoNewtype sync.Once

func Get_Main_patternDoNewtype() gopurs_runtime.Value {
	once_Main_patternDoNewtype.Do(func() {
		cache_Main_patternDoNewtype = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoNewtype
}

var cache_Main_patternDoMultipleWithNormal gopurs_runtime.Value
var once_Main_patternDoMultipleWithNormal sync.Once

func Get_Main_patternDoMultipleWithNormal() gopurs_runtime.Value {
	once_Main_patternDoMultipleWithNormal.Do(func() {
		cache_Main_patternDoMultipleWithNormal = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoMultipleWithNormal
}

var cache_Main_patternDoMultiple gopurs_runtime.Value
var once_Main_patternDoMultiple sync.Once

func Get_Main_patternDoMultiple() gopurs_runtime.Value {
	once_Main_patternDoMultiple.Do(func() {
		cache_Main_patternDoMultiple = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoMultiple
}

var cache_Main_patternDoDataIgnored gopurs_runtime.Value
var once_Main_patternDoDataIgnored sync.Once

func Get_Main_patternDoDataIgnored() gopurs_runtime.Value {
	once_Main_patternDoDataIgnored.Do(func() {
		cache_Main_patternDoDataIgnored = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoDataIgnored
}

var cache_Main_patternDoData gopurs_runtime.Value
var once_Main_patternDoData sync.Once

func Get_Main_patternDoData() gopurs_runtime.Value {
	once_Main_patternDoData.Do(func() {
		cache_Main_patternDoData = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoData
}

var cache_Main_patternDoArray gopurs_runtime.Value
var once_Main_patternDoArray sync.Once

func Get_Main_patternDoArray() gopurs_runtime.Value {
	once_Main_patternDoArray.Do(func() {
		cache_Main_patternDoArray = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_patternDoArray
}

var cache_Main_patternDataIgnored gopurs_runtime.Value
var once_Main_patternDataIgnored sync.Once

func Get_Main_patternDataIgnored() gopurs_runtime.Value {
	once_Main_patternDataIgnored.Do(func() {
		cache_Main_patternDataIgnored = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternDataIgnored
}

var cache_Main_patternData gopurs_runtime.Value
var once_Main_patternData sync.Once

func Get_Main_patternData() gopurs_runtime.Value {
	once_Main_patternData.Do(func() {
		cache_Main_patternData = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternData
}

var cache_Main_patternArray gopurs_runtime.Value
var once_Main_patternArray sync.Once

func Get_Main_patternArray() gopurs_runtime.Value {
	once_Main_patternArray.Do(func() {
		cache_Main_patternArray = gopurs_runtime.Bool(true)
	})
	return cache_Main_patternArray
}

var cache_Main_eqList gopurs_runtime.Value
var once_Main_eqList sync.Once

func Get_Main_eqList() gopurs_runtime.Value {
	once_Main_eqList.Do(func() {
		cache_Main_eqList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqList(dictEq_0_box)
		})
	})
	return cache_Main_eqList
}

var cache_Main_eqList1 gopurs_runtime.Value
var once_Main_eqList1 sync.Once

func Get_Main_eqList1() gopurs_runtime.Value {
	once_Main_eqList1.Do(func() {
		cache_Main_eqList1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Main_Cons[int64]]{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var go__go_2_0_1 gopurs_runtime.Value
				_ = go__go_2_0_1
				go__go_2_0_1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
							var __t1 bool
							{
								if ((v2_5.IntVal) != (0)) != (true) {
									__t1 = false
									goto end_branch_1
								} else {

								}
							}
							{
								if v_3.Type == 9 && v_3.IntVal == 322902991 && v_3.UnsafePtr == nil {
									__t1 = (v1_4.Type == 9 && v1_4.IntVal == 322902991 && v1_4.UnsafePtr == nil) && ((v2_5.IntVal) != (0))
									goto end_branch_1
								} else {

								}
							}
							{
								__t1 = (v_3.Type == 9 && v_3.IntVal == 322902991 && v_3.UnsafePtr != nil) && ((v1_4.Type == 9 && v1_4.IntVal == 322902991 && v1_4.UnsafePtr != nil) && ((gopurs_runtime.Apply3(go__go_2_0_1, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_5.IntVal) != (0)) && (((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0.IntVal) == ((*Constructor_Main_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal)))).IntVal) != (0)))
							}
						end_branch_1:
							return gopurs_runtime.Bool(__t1)
						})
					})
				})
				return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_2_0_1, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](xs_0))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Bool(true)).IntVal) != (0))
			})
		})}))}
	})
	return cache_Main_eqList1
}

var cache_Main_patternDoWithInfixOp gopurs_runtime.Value
var once_Main_patternDoWithInfixOp sync.Once

func Get_Main_patternDoWithInfixOp() gopurs_runtime.Value {
	once_Main_patternDoWithInfixOp.Do(func() {
		cache_Main_patternDoWithInfixOp = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var go__go_0_0_2 gopurs_runtime.Value
			_ = go__go_0_0_2
			go__go_0_0_2 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t1 bool
						{
							if ((v2_3.IntVal) != (0)) != (true) {
								__t1 = false
								goto end_branch_1
							} else {

							}
						}
						{
							if v_1.Type == 9 && v_1.IntVal == 322902991 && v_1.UnsafePtr == nil {
								__t1 = (v1_2.Type == 9 && v1_2.IntVal == 322902991 && v1_2.UnsafePtr == nil) && ((v2_3.IntVal) != (0))
								goto end_branch_1
							} else {

							}
						}
						{
							__t1 = (v_1.Type == 9 && v_1.IntVal == 322902991 && v_1.UnsafePtr != nil) && ((v1_2.Type == 9 && v1_2.IntVal == 322902991 && v1_2.UnsafePtr != nil) && ((gopurs_runtime.Apply3(go__go_0_0_2, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_3.IntVal) != (0)) && (((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0.IntVal) == ((*Constructor_Main_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0.IntVal)))).IntVal) != (0)))
						}
					end_branch_1:
						return gopurs_runtime.Bool(__t1)
					})
				})
			})
			return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_0_0_2, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(2), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(3), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(4), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)})})}))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[int64]{1, gopurs_runtime.Int(2), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(3), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(4), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)})})}))}, gopurs_runtime.Bool(true)).IntVal) != (0))
		})
	})
	return cache_Main_patternDoWithInfixOp
}

var cache_Main_patternWithInfixOp gopurs_runtime.Value
var once_Main_patternWithInfixOp sync.Once

func Get_Main_patternWithInfixOp() gopurs_runtime.Value {
	once_Main_patternWithInfixOp.Do(func() {
		cache_Main_patternWithInfixOp = func() gopurs_runtime.Value {
			var go__go_0_0_3 gopurs_runtime.Value
			_ = go__go_0_0_3
			go__go_0_0_3 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t1 bool
						{
							if ((v2_3.IntVal) != (0)) != (true) {
								__t1 = false
								goto end_branch_1
							} else {

							}
						}
						{
							if v_1.Type == 9 && v_1.IntVal == 322902991 && v_1.UnsafePtr == nil {
								__t1 = (v1_2.Type == 9 && v1_2.IntVal == 322902991 && v1_2.UnsafePtr == nil) && ((v2_3.IntVal) != (0))
								goto end_branch_1
							} else {

							}
						}
						{
							__t1 = (v_1.Type == 9 && v_1.IntVal == 322902991 && v_1.UnsafePtr != nil) && ((v1_2.Type == 9 && v1_2.IntVal == 322902991 && v1_2.UnsafePtr != nil) && ((gopurs_runtime.Apply3(go__go_0_0_3, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_3.IntVal) != (0)) && (((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0.IntVal) == ((*Constructor_Main_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0.IntVal)))).IntVal) != (0)))
						}
					end_branch_1:
						return gopurs_runtime.Bool(__t1)
					})
				})
			})
			return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_0_0_3, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(2), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(3), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(4), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)})})}))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[int64]{1, gopurs_runtime.Int(2), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(3), (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Int(4), (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)})})}))}, gopurs_runtime.Bool(true)).IntVal) != (0))
		}()
	})
	return cache_Main_patternWithInfixOp
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("simple variable pattern"), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			__local_var_2_3 := gopurs_runtime.Bool(true)
			_ = __local_var_2_3
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("simple variable pattern with do"), __local_var_2_3), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (newtype)"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_3_4
			__local_var_4_6 := gopurs_runtime.Bool(true)
			_ = __local_var_4_6
			_dollar___unused_4_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (newtype) with do"), __local_var_4_6), gopurs_runtime.Value{})
			_ = _dollar___unused_4_5
			_dollar___unused_5_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (data)"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_5_7
			_dollar___unused_6_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern with ignorances"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_8
			__local_var_7_10 := gopurs_runtime.Bool(true)
			_ = __local_var_7_10
			_dollar___unused_7_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (data) with do"), __local_var_7_10), gopurs_runtime.Value{})
			_ = _dollar___unused_7_9
			__local_var_8_12 := gopurs_runtime.Bool(true)
			_ = __local_var_8_12
			_dollar___unused_8_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern with ignorances and do"), __local_var_8_12), gopurs_runtime.Value{})
			_ = _dollar___unused_8_11
			_dollar___unused_9_13 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("array pattern"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_9_13
			__local_var_10_15 := gopurs_runtime.Bool(true)
			_ = __local_var_10_15
			_dollar___unused_10_14 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("array pattern with do"), __local_var_10_15), gopurs_runtime.Value{})
			_ = _dollar___unused_10_14
			_dollar___unused_11_16 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_11_16
			__local_var_12_18 := gopurs_runtime.Bool(true)
			_ = __local_var_12_18
			_dollar___unused_12_17 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with do"), __local_var_12_18), gopurs_runtime.Value{})
			_ = _dollar___unused_12_17
			_dollar___unused_13_19 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with normal let's"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_13_19
			__local_var_14_21 := gopurs_runtime.Bool(true)
			_ = __local_var_14_21
			_dollar___unused_14_20 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with normal let's and do"), __local_var_14_21), gopurs_runtime.Value{})
			_ = _dollar___unused_14_20
			_dollar___unused_15_22 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with parens"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_15_22
			__local_var_16_24 := gopurs_runtime.Bool(true)
			_ = __local_var_16_24
			_dollar___unused_16_23 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with parens and do"), __local_var_16_24), gopurs_runtime.Value{})
			_ = _dollar___unused_16_23
			_dollar___unused_17_25 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with named binder"), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_17_25
			__local_var_18_27 := gopurs_runtime.Bool(true)
			_ = __local_var_18_27
			_dollar___unused_18_26 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with named binder and do"), __local_var_18_27), gopurs_runtime.Value{})
			_ = _dollar___unused_18_26
			_dollar___unused_19_28 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("pattern with infix operator"), gopurs_runtime.Bool((Get_Main_patternWithInfixOp().IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_19_28
			__local_var_20_30 := gopurs_runtime.Apply(Get_Main_patternDoWithInfixOp(), gopurs_runtime.Value{})
			_ = __local_var_20_30
			_dollar___unused_20_29 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("pattern with infix operator and do"), __local_var_20_30), gopurs_runtime.Value{})
			_ = _dollar___unused_20_29
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Y struct {
	Rc uint32
	V0 int64
	V1 string
	V2 bool
}

type Constructor_Main_Nil[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Cons[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_Cons[gopurs_runtime.Value]
}

func Call_Main_X(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Main_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var go__go_3_0_0 gopurs_runtime.Value
			_ = go__go_3_0_0
			go__go_3_0_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t1 bool
						{
							if ((v2_6.IntVal) != (0)) != (true) {
								__t1 = false
								goto end_branch_1
							} else {

							}
						}
						{
							if v_4.Type == 9 && v_4.IntVal == 322902991 && v_4.UnsafePtr == nil {
								__t1 = (v1_5.Type == 9 && v1_5.IntVal == 322902991 && v1_5.UnsafePtr == nil) && ((v2_6.IntVal) != (0))
								goto end_branch_1
							} else {

							}
						}
						{
							__t1 = (v_4.Type == 9 && v_4.IntVal == 322902991 && v_4.UnsafePtr != nil) && ((v1_5.Type == 9 && v1_5.IntVal == 322902991 && v1_5.UnsafePtr != nil) && ((gopurs_runtime.Apply3(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Main_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0)))
						}
					end_branch_1:
						return gopurs_runtime.Bool(__t1)
					})
				})
			})
			return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
		})
	})}))}
}
