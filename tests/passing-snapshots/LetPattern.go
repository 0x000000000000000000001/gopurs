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
		cache_Main_pure = Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))
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
			return gopurs_runtime.Int(Call_Main_X(x_0_box.IntVal))
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

var cache_Main_Cons__41240261 gopurs_runtime.Value
var once_Main_Cons__41240261 sync.Once

func Get_Main_Cons__41240261() gopurs_runtime.Value {
	once_Main_Cons__41240261.Do(func() {
		cache_Main_Cons__41240261 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803(Call_Main_Cons__41240261(__eta_norm_1_0_box.IntVal, Rebox_Main_176455803_2737593216(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](__eta_norm_0_unused_1_box)))))}
		})
	})
	return cache_Main_Cons__41240261
}

var cache_Main_Cons__2678532228 gopurs_runtime.Value
var once_Main_Cons__2678532228 sync.Once

func Get_Main_Cons__2678532228() gopurs_runtime.Value {
	once_Main_Cons__2678532228.Do(func() {
		cache_Main_Cons__2678532228 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803(Call_Main_Cons__2678532228(__eta_norm_1_0_box.IntVal, Rebox_Main_176455803_2737593216(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box)))))}
		})
	})
	return cache_Main_Cons__2678532228
}

var cache_Main_patternWithParens gopurs_runtime.Value
var once_Main_patternWithParens sync.Once

func Get_Main_patternWithParens() gopurs_runtime.Value {
	once_Main_patternWithParens.Do(func() {
		cache_Main_patternWithParens = gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})).IntVal) != (0))
	})
	return cache_Main_patternWithParens
}

var cache_Main_patternWithNamedBinder gopurs_runtime.Value
var once_Main_patternWithNamedBinder sync.Once

func Get_Main_patternWithNamedBinder() gopurs_runtime.Value {
	once_Main_patternWithNamedBinder.Do(func() {
		cache_Main_patternWithNamedBinder = gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})).IntVal) != (0))
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
		cache_Main_patternMultipleWithNormal = gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})).IntVal) != (0))
	})
	return cache_Main_patternMultipleWithNormal
}

var cache_Main_patternMultiple gopurs_runtime.Value
var once_Main_patternMultiple sync.Once

func Get_Main_patternMultiple() gopurs_runtime.Value {
	once_Main_patternMultiple.Do(func() {
		cache_Main_patternMultiple = gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})).IntVal) != (0))
	})
	return cache_Main_patternMultiple
}

var cache_Main_patternDoWithParens gopurs_runtime.Value
var once_Main_patternDoWithParens sync.Once

func Get_Main_patternDoWithParens() gopurs_runtime.Value {
	once_Main_patternDoWithParens.Do(func() {
		cache_Main_patternDoWithParens = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoWithParens
}

var cache_Main_patternDoWithNamedBinder gopurs_runtime.Value
var once_Main_patternDoWithNamedBinder sync.Once

func Get_Main_patternDoWithNamedBinder() gopurs_runtime.Value {
	once_Main_patternDoWithNamedBinder.Do(func() {
		cache_Main_patternDoWithNamedBinder = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoWithNamedBinder
}

var cache_Main_patternDoSimple gopurs_runtime.Value
var once_Main_patternDoSimple sync.Once

func Get_Main_patternDoSimple() gopurs_runtime.Value {
	once_Main_patternDoSimple.Do(func() {
		cache_Main_patternDoSimple = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoSimple
}

var cache_Main_patternDoNewtype gopurs_runtime.Value
var once_Main_patternDoNewtype sync.Once

func Get_Main_patternDoNewtype() gopurs_runtime.Value {
	once_Main_patternDoNewtype.Do(func() {
		cache_Main_patternDoNewtype = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoNewtype
}

var cache_Main_patternDoMultipleWithNormal gopurs_runtime.Value
var once_Main_patternDoMultipleWithNormal sync.Once

func Get_Main_patternDoMultipleWithNormal() gopurs_runtime.Value {
	once_Main_patternDoMultipleWithNormal.Do(func() {
		cache_Main_patternDoMultipleWithNormal = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoMultipleWithNormal
}

var cache_Main_patternDoMultiple gopurs_runtime.Value
var once_Main_patternDoMultiple sync.Once

func Get_Main_patternDoMultiple() gopurs_runtime.Value {
	once_Main_patternDoMultiple.Do(func() {
		cache_Main_patternDoMultiple = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoMultiple
}

var cache_Main_patternDoDataIgnored gopurs_runtime.Value
var once_Main_patternDoDataIgnored sync.Once

func Get_Main_patternDoDataIgnored() gopurs_runtime.Value {
	once_Main_patternDoDataIgnored.Do(func() {
		cache_Main_patternDoDataIgnored = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoDataIgnored
}

var cache_Main_patternDoData gopurs_runtime.Value
var once_Main_patternDoData sync.Once

func Get_Main_patternDoData() gopurs_runtime.Value {
	once_Main_patternDoData.Do(func() {
		cache_Main_patternDoData = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
	})
	return cache_Main_patternDoData
}

var cache_Main_patternDoArray gopurs_runtime.Value
var once_Main_patternDoArray sync.Once

func Get_Main_patternDoArray() gopurs_runtime.Value {
	once_Main_patternDoArray.Do(func() {
		cache_Main_patternDoArray = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool(true))
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
		cache_Main_patternArray = gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})).IntVal) != (0))
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
		cache_Main_eqList1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2495748075_3790796878(Rebox_Main_3790796878_2495748075(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqList1
}

var cache_Main_patternDoWithInfixOp gopurs_runtime.Value
var once_Main_patternDoWithInfixOp sync.Once

func Get_Main_patternDoWithInfixOp() gopurs_runtime.Value {
	once_Main_patternDoWithInfixOp.Do(func() {
		cache_Main_patternDoWithInfixOp = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803((&Constructor_Main_Cons[int64]{1, int64(2), (&Constructor_Main_Cons[int64]{1, int64(3), (&Constructor_Main_Cons[int64]{1, int64(4), (*Constructor_Main_Cons[int64])(nil)})})})))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803((&Constructor_Main_Cons[int64]{1, int64(2), (&Constructor_Main_Cons[int64]{1, int64(3), (&Constructor_Main_Cons[int64]{1, int64(4), (*Constructor_Main_Cons[int64])(nil)})})})))}).IntVal) != (0)))
	})
	return cache_Main_patternDoWithInfixOp
}

var cache_Main_patternWithInfixOp gopurs_runtime.Value
var once_Main_patternWithInfixOp sync.Once

func Get_Main_patternWithInfixOp() gopurs_runtime.Value {
	once_Main_patternWithInfixOp.Do(func() {
		cache_Main_patternWithInfixOp = gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803((&Constructor_Main_Cons[int64]{1, int64(2), (&Constructor_Main_Cons[int64]{1, int64(3), (&Constructor_Main_Cons[int64]{1, int64(4), (*Constructor_Main_Cons[int64])(nil)})})})))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803((&Constructor_Main_Cons[int64]{1, int64(2), (&Constructor_Main_Cons[int64]{1, int64(3), (&Constructor_Main_Cons[int64]{1, int64(4), (*Constructor_Main_Cons[int64])(nil)})})})))}).IntVal) != (0))
		})).IntVal) != (0))
	})
	return cache_Main_patternWithInfixOp
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("simple variable pattern"), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoSimple(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("simple variable pattern with do"))), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (newtype)"), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoNewtype(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (newtype) with do"))), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (data)"), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern with ignorances"), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoData(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern (data) with do"))), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoDataIgnored(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("constructor pattern with ignorances and do"))), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("array pattern"), gopurs_runtime.Bool((Get_Main_patternArray().IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoArray(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("array pattern with do"))), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns"), gopurs_runtime.Bool((Get_Main_patternMultiple().IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoMultiple(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with do"))), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with normal let's"), gopurs_runtime.Bool((Get_Main_patternMultipleWithNormal().IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoMultipleWithNormal(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with normal let's and do"))), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with parens"), gopurs_runtime.Bool((Get_Main_patternWithParens().IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoWithParens(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with parens and do"))), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with named binder"), gopurs_runtime.Bool((Get_Main_patternWithNamedBinder().IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoWithNamedBinder(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("multiple patterns with named binder and do"))), gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
																				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("pattern with infix operator"), gopurs_runtime.Bool((Get_Main_patternWithInfixOp().IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
																					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_patternDoWithInfixOp(), gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("pattern with infix operator and do"))), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
																						return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																					}))
																				}))
																			}))
																		}))
																	}))
																}))
															}))
														}))
													}))
												}))
											}))
										}))
									}))
								}))
							}))
						}))
					}))
				}))
			}))
		}))
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
	V0 T_a
	V1 *Constructor_Main_Cons[T_a]
}

func Call_Main_X(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Cons__41240261(__eta_norm_1_0_loop int64, __eta_norm_0_unused_1_loop *Constructor_Main_Cons[int64]) *Constructor_Main_Cons[int64] {
Cons__41240261:
	for {
		if false {
			continue Cons__41240261
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 *Constructor_Main_Cons[int64] = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Main_Cons[int64]{1, __eta_norm_1_0, (*Constructor_Main_Cons[int64])(nil)})
	}
}

func Call_Main_Cons__2678532228(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Main_Cons[int64]) *Constructor_Main_Cons[int64] {
Cons__2678532228:
	for {
		if false {
			continue Cons__2678532228
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Main_Cons[int64] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Cons[int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2465999344_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Main_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var go__go_3_0_0 gopurs_runtime.Value
		_ = go__go_3_0_0
		var go__go_3_0_0_cell *gopurs_runtime.Value
		_ = go__go_3_0_0_cell
		// FALLBACK TCO: isLoop=false len=1
		go__go_3_0_0 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t6 bool
			{
				if ((v2_6.IntVal) != (0)) != (true) {
					__t6 = false
					goto end_branch_6
				} else {

				}
			}
			{
				var __t_tag_4 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v_4)
				_ = __t_tag_4
				if __t_tag_4 == nil {
					var __t_tag_5 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v1_5)
					_ = __t_tag_5
					__t6 = (__t_tag_5 == nil) && ((v2_6.IntVal) != (0))
					goto end_branch_6
				} else {

				}
			}
			{
				var __t_tag_1 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v_4)
				_ = __t_tag_1
				var __t_and_3 bool = false
				if __t_tag_1 != nil {

					var __t_tag_2 *Constructor_Main_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v1_5)
					_ = __t_tag_2
					__t_and_3 = (__t_tag_2 != nil) && ((gopurs_runtime.Bool((gopurs_runtime.Apply3((*go__go_3_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Main_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0)).IntVal) != (0))
				}
				__t6 = __t_and_3
			}
		end_branch_6:
			return gopurs_runtime.Bool(__t6)
		})
		go__go_3_0_0_cell = &go__go_3_0_0
		return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
	})})))}
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_176455803_2737593216(in *Constructor_Main_Cons[gopurs_runtime.Value]) *Constructor_Main_Cons[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Cons[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Main_176455803_2737593216(in.V1)
	return out
}

func Rebox_Main_2465999344_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Cons[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2495748075_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Cons[int64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2737593216_176455803(in *Constructor_Main_Cons[int64]) *Constructor_Main_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = Rebox_Main_2737593216_176455803(in.V1)
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2495748075(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Main_Cons[int64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Main_Cons[int64]]{}
	out.V0 = in.V0
	return out
}
