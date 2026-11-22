package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_CodePoints_add gopurs_runtime.Value
var once_Data_String_CodePoints_add sync.Once

func Get_Data_String_CodePoints_add() gopurs_runtime.Value {
	once_Data_String_CodePoints_add.Do(func() {
		cache_Data_String_CodePoints_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_String_CodePoints_add
}

var cache_Data_String_CodePoints_mul gopurs_runtime.Value
var once_Data_String_CodePoints_mul sync.Once

func Get_Data_String_CodePoints_mul() gopurs_runtime.Value {
	once_Data_String_CodePoints_mul.Do(func() {
		cache_Data_String_CodePoints_mul = Get_Data_Semiring_mul__4086903470()
	})
	return cache_Data_String_CodePoints_mul
}

var cache_Data_String_CodePoints_sub gopurs_runtime.Value
var once_Data_String_CodePoints_sub sync.Once

func Get_Data_String_CodePoints_sub() gopurs_runtime.Value {
	once_Data_String_CodePoints_sub.Do(func() {
		cache_Data_String_CodePoints_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_String_CodePoints_sub
}

var cache_Data_String_CodePoints_compose gopurs_runtime.Value
var once_Data_String_CodePoints_compose sync.Once

func Get_Data_String_CodePoints_compose() gopurs_runtime.Value {
	once_Data_String_CodePoints_compose.Do(func() {
		cache_Data_String_CodePoints_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_CodePoints_compose
}

var cache_Data_String_CodePoints_compose1 gopurs_runtime.Value
var once_Data_String_CodePoints_compose1 sync.Once

func Get_Data_String_CodePoints_compose1() gopurs_runtime.Value {
	once_Data_String_CodePoints_compose1.Do(func() {
		cache_Data_String_CodePoints_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_CodePoints_compose1
}

var cache_Data_String_CodePoints_fromEnum gopurs_runtime.Value
var once_Data_String_CodePoints_fromEnum sync.Once

func Get_Data_String_CodePoints_fromEnum() gopurs_runtime.Value {
	once_Data_String_CodePoints_fromEnum.Do(func() {
		cache_Data_String_CodePoints_fromEnum = Get_Data_Enum_toCharCode()
	})
	return cache_Data_String_CodePoints_fromEnum
}

var cache_Data_String_CodePoints_go__map gopurs_runtime.Value
var once_Data_String_CodePoints_go__map sync.Once

func Get_Data_String_CodePoints_go__map() gopurs_runtime.Value {
	once_Data_String_CodePoints_go__map.Do(func() {
		cache_Data_String_CodePoints_go__map = Get_Data_Functor_map__1382837546()
	})
	return cache_Data_String_CodePoints_go__map
}

var cache_Data_String_CodePoints_unfoldr gopurs_runtime.Value
var once_Data_String_CodePoints_unfoldr sync.Once

func Get_Data_String_CodePoints_unfoldr() gopurs_runtime.Value {
	once_Data_String_CodePoints_unfoldr.Do(func() {
		cache_Data_String_CodePoints_unfoldr = Get_Data_Unfoldable_unfoldr__2862945647()
	})
	return cache_Data_String_CodePoints_unfoldr
}

var cache_Data_String_CodePoints_go__append gopurs_runtime.Value
var once_Data_String_CodePoints_go__append sync.Once

func Get_Data_String_CodePoints_go__append() gopurs_runtime.Value {
	once_Data_String_CodePoints_go__append.Do(func() {
		cache_Data_String_CodePoints_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_String_CodePoints_go__append
}

var cache_Data_String_CodePoints_map1 gopurs_runtime.Value
var once_Data_String_CodePoints_map1 sync.Once

func Get_Data_String_CodePoints_map1() gopurs_runtime.Value {
	once_Data_String_CodePoints_map1.Do(func() {
		cache_Data_String_CodePoints_map1 = Get_Data_Functor_map__3627393546()
	})
	return cache_Data_String_CodePoints_map1
}

var cache_Data_String_CodePoints_conj gopurs_runtime.Value
var once_Data_String_CodePoints_conj sync.Once

func Get_Data_String_CodePoints_conj() gopurs_runtime.Value {
	once_Data_String_CodePoints_conj.Do(func() {
		cache_Data_String_CodePoints_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_String_CodePoints_conj
}

var cache_Data_String_CodePoints_lessThanOrEq gopurs_runtime.Value
var once_Data_String_CodePoints_lessThanOrEq sync.Once

func Get_Data_String_CodePoints_lessThanOrEq() gopurs_runtime.Value {
	once_Data_String_CodePoints_lessThanOrEq.Do(func() {
		cache_Data_String_CodePoints_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_String_CodePoints_lessThanOrEq
}

var cache_Data_String_CodePoints_fromEnum1 gopurs_runtime.Value
var once_Data_String_CodePoints_fromEnum1 sync.Once

func Get_Data_String_CodePoints_fromEnum1() gopurs_runtime.Value {
	once_Data_String_CodePoints_fromEnum1.Do(func() {
		cache_Data_String_CodePoints_fromEnum1 = Get_Data_Enum_fromEnum__1526214825()
	})
	return cache_Data_String_CodePoints_fromEnum1
}

var cache_Data_String_CodePoints_greaterThan gopurs_runtime.Value
var once_Data_String_CodePoints_greaterThan sync.Once

func Get_Data_String_CodePoints_greaterThan() gopurs_runtime.Value {
	once_Data_String_CodePoints_greaterThan.Do(func() {
		cache_Data_String_CodePoints_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_String_CodePoints_greaterThan
}

var cache_Data_String_CodePoints_div gopurs_runtime.Value
var once_Data_String_CodePoints_div sync.Once

func Get_Data_String_CodePoints_div() gopurs_runtime.Value {
	once_Data_String_CodePoints_div.Do(func() {
		cache_Data_String_CodePoints_div = Get_Data_EuclideanRing_div__478053998()
	})
	return cache_Data_String_CodePoints_div
}

var cache_Data_String_CodePoints_mod gopurs_runtime.Value
var once_Data_String_CodePoints_mod sync.Once

func Get_Data_String_CodePoints_mod() gopurs_runtime.Value {
	once_Data_String_CodePoints_mod.Do(func() {
		cache_Data_String_CodePoints_mod = Get_Data_EuclideanRing_mod__478053998()
	})
	return cache_Data_String_CodePoints_mod
}

var cache_Data_String_CodePoints_lessThan gopurs_runtime.Value
var once_Data_String_CodePoints_lessThan sync.Once

func Get_Data_String_CodePoints_lessThan() gopurs_runtime.Value {
	once_Data_String_CodePoints_lessThan.Do(func() {
		cache_Data_String_CodePoints_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_String_CodePoints_lessThan
}

var cache_Data_String_CodePoints_eq gopurs_runtime.Value
var once_Data_String_CodePoints_eq sync.Once

func Get_Data_String_CodePoints_eq() gopurs_runtime.Value {
	once_Data_String_CodePoints_eq.Do(func() {
		cache_Data_String_CodePoints_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_String_CodePoints_eq
}

var cache_Data_String_CodePoints_compare gopurs_runtime.Value
var once_Data_String_CodePoints_compare sync.Once

func Get_Data_String_CodePoints_compare() gopurs_runtime.Value {
	once_Data_String_CodePoints_compare.Do(func() {
		cache_Data_String_CodePoints_compare = Get_Data_Ord_compare__1925727043()
	})
	return cache_Data_String_CodePoints_compare
}

var cache_Data_String_CodePoints_greaterThanOrEq gopurs_runtime.Value
var once_Data_String_CodePoints_greaterThanOrEq sync.Once

func Get_Data_String_CodePoints_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_String_CodePoints_greaterThanOrEq.Do(func() {
		cache_Data_String_CodePoints_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_String_CodePoints_greaterThanOrEq
}

var cache_Data_String_CodePoints_CodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_CodePoint sync.Once

func Get_Data_String_CodePoints_CodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_CodePoint.Do(func() {
		cache_Data_String_CodePoints_CodePoint = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_String_CodePoints_CodePoint(x_0_box.IntVal))
		})
	})
	return cache_Data_String_CodePoints_CodePoint
}

var cache_Data_String_CodePoints_unsurrogate gopurs_runtime.Value
var once_Data_String_CodePoints_unsurrogate sync.Once

func Get_Data_String_CodePoints_unsurrogate() gopurs_runtime.Value {
	once_Data_String_CodePoints_unsurrogate.Do(func() {
		cache_Data_String_CodePoints_unsurrogate = gopurs_runtime.Func2(func(lead_0_box gopurs_runtime.Value, trail_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_String_CodePoints_unsurrogate(lead_0_box.IntVal, trail_1_box.IntVal))
		})
	})
	return cache_Data_String_CodePoints_unsurrogate
}

var cache_Data_String_CodePoints_uncons gopurs_runtime.Value
var once_Data_String_CodePoints_uncons sync.Once

func Get_Data_String_CodePoints_uncons() gopurs_runtime.Value {
	once_Data_String_CodePoints_uncons.Do(func() {
		cache_Data_String_CodePoints_uncons = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_uncons(s_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_uncons
}

var cache_Data_String_CodePoints_unconsButWithTuple gopurs_runtime.Value
var once_Data_String_CodePoints_unconsButWithTuple sync.Once

func Get_Data_String_CodePoints_unconsButWithTuple() gopurs_runtime.Value {
	once_Data_String_CodePoints_unconsButWithTuple.Do(func() {
		cache_Data_String_CodePoints_unconsButWithTuple = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_unconsButWithTuple(s_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_unconsButWithTuple
}

var cache_Data_String_CodePoints_toCodePointArrayFallback gopurs_runtime.Value
var once_Data_String_CodePoints_toCodePointArrayFallback sync.Once

func Get_Data_String_CodePoints_toCodePointArrayFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_toCodePointArrayFallback.Do(func() {
		cache_Data_String_CodePoints_toCodePointArrayFallback = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_String_CodePoints_toCodePointArrayFallback(s_0_box.StrVal())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_String_CodePoints_toCodePointArrayFallback
}

var cache_Data_String_CodePoints_showCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_showCodePoint sync.Once

func Get_Data_String_CodePoints_showCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_showCodePoint.Do(func() {
		cache_Data_String_CodePoints_showCodePoint = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(CodePoint 0x", Data_String_Common_ToUpper(gopurs_runtime.Apply2(Get_Data_Int_toStringAs(), gopurs_runtime.Int(int64(16)), gopurs_runtime.Int(v_0.IntVal)).StrVal())), ")"))
		})})))}
	})
	return cache_Data_String_CodePoints_showCodePoint
}

var cache_Data_String_CodePoints_lastIndexOf gopurs_runtime.Value
var once_Data_String_CodePoints_lastIndexOf sync.Once

func Get_Data_String_CodePoints_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_CodePoints_lastIndexOf.Do(func() {
		cache_Data_String_CodePoints_lastIndexOf = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_lastIndexOf(p_0_box.StrVal(), s_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_lastIndexOf
}

var cache_Data_String_CodePoints_isTrail gopurs_runtime.Value
var once_Data_String_CodePoints_isTrail sync.Once

func Get_Data_String_CodePoints_isTrail() gopurs_runtime.Value {
	once_Data_String_CodePoints_isTrail.Do(func() {
		cache_Data_String_CodePoints_isTrail = gopurs_runtime.Func(func(cu_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_String_CodePoints_isTrail(cu_0_box.IntVal))
		})
	})
	return cache_Data_String_CodePoints_isTrail
}

var cache_Data_String_CodePoints_isLead gopurs_runtime.Value
var once_Data_String_CodePoints_isLead sync.Once

func Get_Data_String_CodePoints_isLead() gopurs_runtime.Value {
	once_Data_String_CodePoints_isLead.Do(func() {
		cache_Data_String_CodePoints_isLead = gopurs_runtime.Func(func(cu_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_String_CodePoints_isLead(cu_0_box.IntVal))
		})
	})
	return cache_Data_String_CodePoints_isLead
}

var cache_Data_String_CodePoints_unsafeCodePointAt0Fallback gopurs_runtime.Value
var once_Data_String_CodePoints_unsafeCodePointAt0Fallback sync.Once

func Get_Data_String_CodePoints_unsafeCodePointAt0Fallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_unsafeCodePointAt0Fallback.Do(func() {
		cache_Data_String_CodePoints_unsafeCodePointAt0Fallback = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_String_CodePoints_unsafeCodePointAt0Fallback(s_0_box.StrVal()))
		})
	})
	return cache_Data_String_CodePoints_unsafeCodePointAt0Fallback
}

var cache_Data_String_CodePoints_unsafeCodePointAt0 gopurs_runtime.Value
var once_Data_String_CodePoints_unsafeCodePointAt0 sync.Once

func Get_Data_String_CodePoints_unsafeCodePointAt0() gopurs_runtime.Value {
	once_Data_String_CodePoints_unsafeCodePointAt0.Do(func() {
		cache_Data_String_CodePoints_unsafeCodePointAt0 = gopurs_runtime.Apply(Get_Data_String_CodePoints__unsafeCodePointAt0(), Get_Data_String_CodePoints_unsafeCodePointAt0Fallback())
	})
	return cache_Data_String_CodePoints_unsafeCodePointAt0
}

var cache_Data_String_CodePoints_toCodePointArray gopurs_runtime.Value
var once_Data_String_CodePoints_toCodePointArray sync.Once

func Get_Data_String_CodePoints_toCodePointArray() gopurs_runtime.Value {
	once_Data_String_CodePoints_toCodePointArray.Do(func() {
		cache_Data_String_CodePoints_toCodePointArray = gopurs_runtime.Apply2(Get_Data_String_CodePoints__toCodePointArray(), Get_Data_String_CodePoints_toCodePointArrayFallback(), Get_Data_String_CodePoints_unsafeCodePointAt0())
	})
	return cache_Data_String_CodePoints_toCodePointArray
}

var cache_Data_String_CodePoints_indexOf gopurs_runtime.Value
var once_Data_String_CodePoints_indexOf sync.Once

func Get_Data_String_CodePoints_indexOf() gopurs_runtime.Value {
	once_Data_String_CodePoints_indexOf.Do(func() {
		cache_Data_String_CodePoints_indexOf = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_indexOf(p_0_box.StrVal(), s_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_indexOf
}

var cache_Data_String_CodePoints_fromCharCode gopurs_runtime.Value
var once_Data_String_CodePoints_fromCharCode sync.Once

func Get_Data_String_CodePoints_fromCharCode() gopurs_runtime.Value {
	once_Data_String_CodePoints_fromCharCode.Do(func() {
		cache_Data_String_CodePoints_fromCharCode = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__4060784579(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))))
	})
	return cache_Data_String_CodePoints_fromCharCode
}

var cache_Data_String_CodePoints_singletonFallback gopurs_runtime.Value
var once_Data_String_CodePoints_singletonFallback sync.Once

func Get_Data_String_CodePoints_singletonFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_singletonFallback.Do(func() {
		cache_Data_String_CodePoints_singletonFallback = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodePoints_singletonFallback(v_0_box.IntVal))
		})
	})
	return cache_Data_String_CodePoints_singletonFallback
}

var cache_Data_String_CodePoints_fromCodePointArray gopurs_runtime.Value
var once_Data_String_CodePoints_fromCodePointArray sync.Once

func Get_Data_String_CodePoints_fromCodePointArray() gopurs_runtime.Value {
	once_Data_String_CodePoints_fromCodePointArray.Do(func() {
		cache_Data_String_CodePoints_fromCodePointArray = gopurs_runtime.Apply(Get_Data_String_CodePoints__fromCodePointArray(), Get_Data_String_CodePoints_singletonFallback())
	})
	return cache_Data_String_CodePoints_fromCodePointArray
}

var cache_Data_String_CodePoints_singleton gopurs_runtime.Value
var once_Data_String_CodePoints_singleton sync.Once

func Get_Data_String_CodePoints_singleton() gopurs_runtime.Value {
	once_Data_String_CodePoints_singleton.Do(func() {
		cache_Data_String_CodePoints_singleton = gopurs_runtime.Apply(Get_Data_String_CodePoints__singleton(), Get_Data_String_CodePoints_singletonFallback())
	})
	return cache_Data_String_CodePoints_singleton
}

var cache_Data_String_CodePoints_takeFallback gopurs_runtime.Value
var once_Data_String_CodePoints_takeFallback sync.Once

func Get_Data_String_CodePoints_takeFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_takeFallback.Do(func() {
		cache_Data_String_CodePoints_takeFallback = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodePoints_takeFallback(v_0_box.IntVal, v1_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodePoints_takeFallback
}

var cache_Data_String_CodePoints_take gopurs_runtime.Value
var once_Data_String_CodePoints_take sync.Once

func Get_Data_String_CodePoints_take() gopurs_runtime.Value {
	once_Data_String_CodePoints_take.Do(func() {
		cache_Data_String_CodePoints_take = gopurs_runtime.Apply(Get_Data_String_CodePoints__take(), Get_Data_String_CodePoints_takeFallback())
	})
	return cache_Data_String_CodePoints_take
}

var cache_Data_String_CodePoints_lastIndexOf_prime_ gopurs_runtime.Value
var once_Data_String_CodePoints_lastIndexOf_prime_ sync.Once

func Get_Data_String_CodePoints_lastIndexOf_prime_() gopurs_runtime.Value {
	once_Data_String_CodePoints_lastIndexOf_prime_.Do(func() {
		cache_Data_String_CodePoints_lastIndexOf_prime_ = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_lastIndexOf_prime_(p_0_box.StrVal(), i_1_box.IntVal, s_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_lastIndexOf_prime_
}

var cache_Data_String_CodePoints_splitAt gopurs_runtime.Value
var once_Data_String_CodePoints_splitAt sync.Once

func Get_Data_String_CodePoints_splitAt() gopurs_runtime.Value {
	once_Data_String_CodePoints_splitAt.Do(func() {
		cache_Data_String_CodePoints_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_String_CodePoints_splitAt(i_0_box.IntVal, s_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Str(orig.after), gopurs_runtime.Str(orig.before))
			}()
		})
	})
	return cache_Data_String_CodePoints_splitAt
}

var cache_Data_String_CodePoints_eqCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_eqCodePoint sync.Once

func Get_Data_String_CodePoints_eqCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_eqCodePoint.Do(func() {
		cache_Data_String_CodePoints_eqCodePoint = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_1053099733_3790796878((&Constructor_Data_Eq_Eq[int64]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((x_0.IntVal) == (y_1.IntVal))
		})})))}
	})
	return cache_Data_String_CodePoints_eqCodePoint
}

var cache_Data_String_CodePoints_ordCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_ordCodePoint sync.Once

func Get_Data_String_CodePoints_ordCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_ordCodePoint.Do(func() {
		cache_Data_String_CodePoints_ordCodePoint = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_3308271157_4177771502((&Constructor_Data_Ord_Ord[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_1053099733_3790796878(Rebox_Data_String_CodePoints_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_String_CodePoints_eqCodePoint()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Rebox_Data_String_CodePoints_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(x_0.IntVal), gopurs_runtime.Int(y_1.IntVal))
		})})))}
	})
	return cache_Data_String_CodePoints_ordCodePoint
}

var cache_Data_String_CodePoints_drop gopurs_runtime.Value
var once_Data_String_CodePoints_drop sync.Once

func Get_Data_String_CodePoints_drop() gopurs_runtime.Value {
	once_Data_String_CodePoints_drop.Do(func() {
		cache_Data_String_CodePoints_drop = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodePoints_drop(n_0_box.IntVal, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodePoints_drop
}

var cache_Data_String_CodePoints_indexOf_prime_ gopurs_runtime.Value
var once_Data_String_CodePoints_indexOf_prime_ sync.Once

func Get_Data_String_CodePoints_indexOf_prime_() gopurs_runtime.Value {
	once_Data_String_CodePoints_indexOf_prime_.Do(func() {
		cache_Data_String_CodePoints_indexOf_prime_ = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_indexOf_prime_(p_0_box.StrVal(), i_1_box.IntVal, s_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_indexOf_prime_
}

var cache_Data_String_CodePoints_countTail gopurs_runtime.Value
var once_Data_String_CodePoints_countTail sync.Once

func Get_Data_String_CodePoints_countTail() gopurs_runtime.Value {
	once_Data_String_CodePoints_countTail.Do(func() {
		cache_Data_String_CodePoints_countTail = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value, accum_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_String_CodePoints_countTail(p_0_box, s_1_box.StrVal(), accum_2_box.IntVal))
		})
	})
	return cache_Data_String_CodePoints_countTail
}

var cache_Data_String_CodePoints_countFallback gopurs_runtime.Value
var once_Data_String_CodePoints_countFallback sync.Once

func Get_Data_String_CodePoints_countFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_countFallback.Do(func() {
		cache_Data_String_CodePoints_countFallback = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_String_CodePoints_countFallback(p_0_box, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodePoints_countFallback
}

var cache_Data_String_CodePoints_countPrefix gopurs_runtime.Value
var once_Data_String_CodePoints_countPrefix sync.Once

func Get_Data_String_CodePoints_countPrefix() gopurs_runtime.Value {
	once_Data_String_CodePoints_countPrefix.Do(func() {
		cache_Data_String_CodePoints_countPrefix = gopurs_runtime.Apply2(Get_Data_String_CodePoints__countPrefix(), Get_Data_String_CodePoints_countFallback(), Get_Data_String_CodePoints_unsafeCodePointAt0())
	})
	return cache_Data_String_CodePoints_countPrefix
}

var cache_Data_String_CodePoints_dropWhile gopurs_runtime.Value
var once_Data_String_CodePoints_dropWhile sync.Once

func Get_Data_String_CodePoints_dropWhile() gopurs_runtime.Value {
	once_Data_String_CodePoints_dropWhile.Do(func() {
		cache_Data_String_CodePoints_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodePoints_dropWhile(p_0_box, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodePoints_dropWhile
}

var cache_Data_String_CodePoints_takeWhile gopurs_runtime.Value
var once_Data_String_CodePoints_takeWhile sync.Once

func Get_Data_String_CodePoints_takeWhile() gopurs_runtime.Value {
	once_Data_String_CodePoints_takeWhile.Do(func() {
		cache_Data_String_CodePoints_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodePoints_takeWhile(p_0_box, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodePoints_takeWhile
}

var cache_Data_String_CodePoints_codePointFromChar gopurs_runtime.Value
var once_Data_String_CodePoints_codePointFromChar sync.Once

func Get_Data_String_CodePoints_codePointFromChar() gopurs_runtime.Value {
	once_Data_String_CodePoints_codePointFromChar.Do(func() {
		cache_Data_String_CodePoints_codePointFromChar = gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Enum_toCharCode(), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}))
	})
	return cache_Data_String_CodePoints_codePointFromChar
}

var cache_Data_String_CodePoints_codePointAtFallback gopurs_runtime.Value
var once_Data_String_CodePoints_codePointAtFallback sync.Once

func Get_Data_String_CodePoints_codePointAtFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_codePointAtFallback.Do(func() {
		cache_Data_String_CodePoints_codePointAtFallback = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_codePointAtFallback(n_0_box.IntVal, s_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_codePointAtFallback
}

var cache_Data_String_CodePoints_codePointAt gopurs_runtime.Value
var once_Data_String_CodePoints_codePointAt sync.Once

func Get_Data_String_CodePoints_codePointAt() gopurs_runtime.Value {
	once_Data_String_CodePoints_codePointAt.Do(func() {
		cache_Data_String_CodePoints_codePointAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_codePointAt(v_0_box.IntVal, v1_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodePoints_codePointAt
}

var cache_Data_String_CodePoints_boundedCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_boundedCodePoint sync.Once

func Get_Data_String_CodePoints_boundedCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_boundedCodePoint.Do(func() {
		cache_Data_String_CodePoints_boundedCodePoint = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_3308271157_4177771502(Rebox_Data_String_CodePoints_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_String_CodePoints_ordCodePoint()))))}
		}), int64(0), int64(1114111)})))}
	})
	return cache_Data_String_CodePoints_boundedCodePoint
}

var cache_Data_String_CodePoints_boundedEnumCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_boundedEnumCodePoint sync.Once

func Get_Data_String_CodePoints_boundedEnumCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_boundedEnumCodePoint.Do(func() {
		cache_Data_String_CodePoints_boundedEnumCodePoint = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_1306125126_123048125((&Constructor_Data_Enum_BoundedEnum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_3764732725_2094947566(Rebox_Data_String_CodePoints_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedCodePoint()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_4060049525_556578094(Rebox_Data_String_CodePoints_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_String_CodePoints_enumCodePoint()))))}
		}), gopurs_runtime.Int(int64(1114112)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return v_0
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if ((n_0.IntVal) >= (int64(0))) && ((n_0.IntVal) <= (int64(1114111))) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{n_0, true}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_0:
			return func() gopurs_runtime.Value {
				_v := __t0
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})})))}
	})
	return cache_Data_String_CodePoints_boundedEnumCodePoint
}

var cache_Data_String_CodePoints_enumCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_enumCodePoint sync.Once

func Get_Data_String_CodePoints_enumCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_enumCodePoint.Do(func() {
		cache_Data_String_CodePoints_enumCodePoint = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_3308271157_4177771502(Rebox_Data_String_CodePoints_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_String_CodePoints_ordCodePoint()))))}
		}), gopurs_runtime.Apply2(Get_Data_Enum_defaultPred__2406132389(), Rebox_Data_String_CodePoints_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V4, Rebox_Data_String_CodePoints_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V3), gopurs_runtime.Apply2(Get_Data_Enum_defaultSucc__2406132389(), Rebox_Data_String_CodePoints_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V4, Rebox_Data_String_CodePoints_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V3)})))}
	})
	return cache_Data_String_CodePoints_enumCodePoint
}

func Call_Data_String_CodePoints_CodePoint(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_String_CodePoints_unsurrogate(lead_0_loop int64, trail_1_loop int64) int64 {
	var lead_0 int64 = lead_0_loop
	_ = lead_0
	var trail_1 int64 = trail_1_loop
	_ = trail_1
	return gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntMul(gopurs_runtime.IntSub(lead_0, int64(55296)), int64(1024)), gopurs_runtime.IntSub(trail_1, int64(56320))), int64(65536))
}

func Call_Data_String_CodePoints_uncons(s_0_loop string) struct {
	V0 struct {
		head int64
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V1 bool
} {
	var s_0 string = s_0_loop
	_ = s_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	v_1_0 := Rebox_Data_String_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Data_String_CodePoints__CodePointAt_nativeWorker(gopurs_runtime.Func2(func(v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	}), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Enum_toCharCode(), Data_String_Unsafe_CharAt_nativeWorker(gopurs_runtime.Int(int64(0))))), int64(0), s_0)))
	_ = v_1_0
	var __t1 *Constructor_Data_Maybe_Just[struct {
		head int64
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]
	{
		if v_1_0 == nil {
			__t1 = Rebox_Data_String_CodePoints_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						head int64
						tail string
						/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
					}
					V1 bool
				}{struct {
					head int64
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_1
		} else {

		}
	}
	{
		if v_1_0 != nil {
			__t1 = Rebox_Data_String_CodePoints_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						head int64
						tail string
						/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
					}
					V1 bool
				}{struct {
					head int64
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{(v_1_0).V0, gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(Data_String_CodePoints__Take(gopurs_runtime.Func2(func(v1_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.StrWithCodeUnits("", 0)
				}), int64(1), s_0))).IntVal), gopurs_runtime.Str(s_0)).StrVal()}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Data_Maybe_Just[struct {
			head int64
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}] {
			panic("Failed pattern match")
		}()
	}
end_branch_1:
	return func() struct {
		V0 struct {
			head int64
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}
		V1 bool
	} {
		_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_1387998409_3094389156(__t1))}
		if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 struct {
					head int64
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V1 bool
			}{V0: func() struct {
				head int64
				tail string
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			} {
				orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					head int64
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}
				clone.head = gopurs_runtime.RecordGet(orig, "head").IntVal
				clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
				return clone
			}(), V1: true}
		}
		return struct {
			V0 struct {
				head int64
				tail string
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V1 bool
		}{V0: struct {
			head int64
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, V1: false}
	}()
}

func Call_Data_String_CodePoints_unconsButWithTuple(s_0_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var s_0 string = s_0_loop
	_ = s_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Data_String_CodePoints_uncons(s_0)
	_ = __local_var_1_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_1_0.V1 {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Int(__local_var_1_0.V0.head), gopurs_runtime.Str(__local_var_1_0.V0.tail)}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_CodePoints_toCodePointArrayFallback(s_0_loop string) []int64 {
	var s_0 string = s_0_loop
	_ = s_0
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()).V1, Get_Data_String_CodePoints_unconsButWithTuple(), gopurs_runtime.Str(s_0)).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Data_String_CodePoints_lastIndexOf(p_0_loop string, s_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var p_0 string = p_0_loop
	_ = p_0
	var s_1 string = s_1_loop
	_ = s_1
	// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
	__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_lastIndexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1))
	_ = __local_var_2_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(Data_String_CodePoints_Length(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Str(s_1)).StrVal())), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_CodePoints_isTrail(cu_0_loop int64) bool {
	var cu_0 int64 = cu_0_loop
	_ = cu_0
	return ((int64(56320)) <= (cu_0)) && ((cu_0) <= (int64(57343)))
}

func Call_Data_String_CodePoints_isLead(cu_0_loop int64) bool {
	var cu_0 int64 = cu_0_loop
	_ = cu_0
	return ((int64(55296)) <= (cu_0)) && ((cu_0) <= (int64(56319)))
}

func Call_Data_String_CodePoints_unsafeCodePointAt0Fallback(s_0_loop string) int64 {
	var s_0 string = s_0_loop
	_ = s_0
	// TAST (Let): cu0_1_0 shape=App(Var) bindingType=Int
	cu0_1_0 := gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(Data_String_Unsafe_CharAt_nativeWorker(gopurs_runtime.Int(int64(0))), gopurs_runtime.Str(s_0)))).IntVal
	_ = cu0_1_0
	var __t3 int64
	{
		if (((int64(55296)) <= (cu0_1_0)) && ((cu0_1_0) <= (int64(56319)))) && ((gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(s_0)).IntVal) > (int64(1))) {
			// TAST (Let): cu1_2_1 shape=App(Var) bindingType=Int
			cu1_2_1 := gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(Data_String_Unsafe_CharAt_nativeWorker(gopurs_runtime.Int(int64(1))), gopurs_runtime.Str(s_0)))).IntVal
			_ = cu1_2_1
			var __t2 int64
			{
				if ((int64(56320)) <= (cu1_2_1)) && ((cu1_2_1) <= (int64(57343))) {
					__t2 = gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntMul(gopurs_runtime.IntSub(cu0_1_0, int64(55296)), int64(1024)), gopurs_runtime.IntSub(cu1_2_1, int64(56320))), int64(65536))
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = cu0_1_0
			}
		end_branch_2:
			__t3 = __t2
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = cu0_1_0
	}
end_branch_3:
	return __t3
}

func Call_Data_String_CodePoints_indexOf(p_0_loop string, s_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var p_0 string = p_0_loop
	_ = p_0
	var s_1 string = s_1_loop
	_ = s_1
	// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
	__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1))
	_ = __local_var_2_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(Data_String_CodePoints_Length(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Str(s_1)).StrVal())), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_CodePoints_singletonFallback(v_0_loop int64) string {
	var v_0 int64 = v_0_loop
	_ = v_0
	var __t0 string
	{
		if (v_0) <= (int64(65535)) {
			__t0 = gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCharCode(), gopurs_runtime.Int(v_0)).StrVal()
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.ConcatString(gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCharCode(), gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntDiv(gopurs_runtime.IntSub(v_0, int64(65536)), int64(1024)), int64(55296)))).StrVal(), gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCharCode(), gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntMod(gopurs_runtime.IntSub(v_0, int64(65536)), int64(1024)), int64(56320)))).StrVal())
	}
end_branch_0:
	return __t0
}

func Call_Data_String_CodePoints_takeFallback(v_0_loop int64, v1_1_loop string) string {
takeFallback:
	for {
		if false {
			continue takeFallback
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 string = v1_1_loop
		_ = v1_1
		var __t2 string
		{
			if (v_0) < (int64(1)) {
				__t2 = ""
				goto end_branch_2
			} else {

			}
		}
		{
			// TAST (Let): v2_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: Int, tail: String] Empty))])
			v2_2_0 := Rebox_Data_String_CodePoints_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_String_CodePoints_uncons(v1_1)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			_ = v2_2_0
			var __t1 string
			{
				if v2_2_0 != nil {
					__t1 = gopurs_runtime.ConcatString(gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int((v2_2_0).V0.head)).StrVal(), Call_Data_String_CodePoints_takeFallback(gopurs_runtime.IntSub(v_0, int64(1)), (v2_2_0).V0.tail))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = v1_1
			}
		end_branch_1:
			__t2 = __t1
		}
	end_branch_2:
		return __t2
	}
}

func Call_Data_String_CodePoints_lastIndexOf_prime_(p_0_loop string, i_1_loop int64, s_2_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var p_0 string = p_0_loop
	_ = p_0
	var i_1 int64 = i_1_loop
	_ = i_1
	var s_2 string = s_2_loop
	_ = s_2
	// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
	__local_var_3_0 := gopurs_runtime.Apply3(Get_Data_String_CodeUnits_lastIndexOf_prime_(), gopurs_runtime.Str(p_0), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2)))).IntVal), gopurs_runtime.Str(s_2))
	_ = __local_var_3_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(Data_String_CodePoints_Length(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0.IntVal), gopurs_runtime.Str(s_2)).StrVal())), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_CodePoints_splitAt(i_0_loop int64, s_1_loop string) struct {
	after  string
	before string
	/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var s_1 string = s_1_loop
	_ = s_1
	// TAST (Let): before_2_0 shape=App(Var) bindingType=String
	before_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(s_1)).StrVal()
	_ = before_2_0
	return struct {
		after  string
		before string
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	}{gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(before_2_0)).IntVal), gopurs_runtime.Str(s_1)).StrVal(), before_2_0}
}

func Call_Data_String_CodePoints_drop(n_0_loop int64, s_1_loop string) string {
	var n_0 int64 = n_0_loop
	_ = n_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(n_0), gopurs_runtime.Str(s_1)))).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodePoints_indexOf_prime_(p_0_loop string, i_1_loop int64, s_2_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var p_0 string = p_0_loop
	_ = p_0
	var i_1 int64 = i_1_loop
	_ = i_1
	var s_2 string = s_2_loop
	_ = s_2
	// TAST (Let): s_prime__3_0 shape=App(Var) bindingType=String
	s_prime__3_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2)))).IntVal), gopurs_runtime.Str(s_2)).StrVal()
	_ = s_prime__3_0
	// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
	__local_var_4_1 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_prime__3_0))
	_ = __local_var_4_1
	var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136 && __local_var_4_1.UnsafePtr != nil {
			__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(gopurs_runtime.IntAdd(i_1, Data_String_CodePoints_Length(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0.IntVal), gopurs_runtime.Str(s_prime__3_0)).StrVal()))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	}
end_branch_2:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_CodePoints_countTail(p_0_loop gopurs_runtime.Value, s_1_loop string, accum_2_loop int64) int64 {
countTail:
	for {
		if false {
			continue countTail
		}
		var p_0 gopurs_runtime.Value = p_0_loop
		_ = p_0
		var s_1 string = s_1_loop
		_ = s_1
		var accum_2 int64 = accum_2_loop
		_ = accum_2
		// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: Int, tail: String] Empty))])
		v_3_0 := Rebox_Data_String_CodePoints_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_String_CodePoints_uncons(s_1)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_3_0
		var __t1 int64
		{
			if (v_3_0 != nil) && ((gopurs_runtime.Apply(p_0, gopurs_runtime.Int((v_3_0).V0.head)).IntVal) != (0)) {
				p_0_loop = p_0
				s_1_loop = (v_3_0).V0.tail
				accum_2_loop = gopurs_runtime.IntAdd(accum_2, int64(1))
				continue countTail
				__t1 = func() int64 { panic("unreachable") }()
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = accum_2
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_String_CodePoints_countFallback(p_0_loop gopurs_runtime.Value, s_1_loop string) int64 {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var s_1 string = s_1_loop
	_ = s_1
	return Call_Data_String_CodePoints_countTail(p_0, s_1, int64(0))
}

func Call_Data_String_CodePoints_dropWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodePoints_countPrefix(), p_0, gopurs_runtime.Str(s_1)).IntVal), gopurs_runtime.Str(s_1)))).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodePoints_takeWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodePoints_countPrefix(), p_0, gopurs_runtime.Str(s_1)).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodePoints_codePointAtFallback(n_0_loop int64, s_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
codePointAtFallback:
	for {
		if false {
			continue codePointAtFallback
		}
		var n_0 int64 = n_0_loop
		_ = n_0
		var s_1 string = s_1_loop
		_ = s_1
		// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: Int, tail: String] Empty))])
		v_2_0 := Rebox_Data_String_CodePoints_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_String_CodePoints_uncons(s_1)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_2_0
		var __t2 *Constructor_Data_Maybe_Just[int64]
		{
			if v_2_0 != nil {
				var __t1 *Constructor_Data_Maybe_Just[int64]
				{
					if (n_0) == (int64(0)) {
						__t1 = Rebox_Data_String_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Int((v_2_0).V0.head), true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))
						goto end_branch_1
					} else {

					}
				}
				{
					n_0_loop = gopurs_runtime.IntSub(n_0, int64(1))
					s_1_loop = (v_2_0).V0.tail
					continue codePointAtFallback
					__t1 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = Rebox_Data_String_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
		}
	end_branch_2:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_1170268447_3094389156(__t2))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_String_CodePoints_codePointAt(v_0_loop int64, v1_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 int64 = v_0_loop
	_ = v_0
	var v1_1 string = v1_1_loop
	_ = v1_1
	var __t1 *Constructor_Data_Maybe_Just[int64]
	{
		if (v_0) < (int64(0)) {
			__t1 = Rebox_Data_String_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_1
		} else {

		}
	}
	{
		if (v_0) == (int64(0)) {
			var __t0 *Constructor_Data_Maybe_Just[int64]
			{
				if (v1_1) == ("") {
					__t0 = Rebox_Data_String_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = Rebox_Data_String_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Apply(Get_Data_String_CodePoints_unsafeCodePointAt0(), gopurs_runtime.Str(v1_1)), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))
			}
		end_branch_0:
			__t1 = __t0
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = Rebox_Data_String_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Data_String_CodePoints__CodePointAt_nativeWorker(Get_Data_String_CodePoints_codePointAtFallback(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}, Get_Data_String_CodePoints_unsafeCodePointAt0(), v_0, v1_1)))
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodePoints_1170268447_3094389156(__t1))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Rebox_Data_String_CodePoints_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_String_CodePoints_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_1387998409_3094389156(in *Constructor_Data_Maybe_Just[struct {
	head int64
	tail string
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
	}()
	return out
}

func Rebox_Data_String_CodePoints_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_2094947566_3764732725(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	out.V2 = in.V2.IntVal
	return out
}

func Rebox_Data_String_CodePoints_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_String_CodePoints_3094389156_1387998409(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	head int64
	tail string
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		head int64
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]{}
	out.V0 = func() struct {
		head int64
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			head int64
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}
		clone.head = gopurs_runtime.RecordGet(orig, "head").IntVal
		clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
		return clone
	}()
	return out
}

func Rebox_Data_String_CodePoints_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_String_CodePoints_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_4060049525_556578094(in *Constructor_Data_Enum_Enum[int64]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Rebox_Data_String_CodePoints_556578094_4060049525(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[int64])(unsafe.Pointer(in))
}

func Get_Data_String_CodePoints__codePointAt() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__CodePointAt
}

func Get_Data_String_CodePoints__countPrefix() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__CountPrefix
}

func Get_Data_String_CodePoints__fromCodePointArray() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__FromCodePointArray
}

func Get_Data_String_CodePoints__singleton() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__Singleton
}

func Get_Data_String_CodePoints__take() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__Take
}

func Get_Data_String_CodePoints__toCodePointArray() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__ToCodePointArray
}

func Get_Data_String_CodePoints__unsafeCodePointAt0() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__UnsafeCodePointAt0
}

func Get_Data_String_CodePoints_length() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints_Length
}
