package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_NonEmpty_CodePoints_go__append gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_go__append sync.Once

func Get_Data_String_NonEmpty_CodePoints_go__append() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_go__append.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_String_NonEmpty_CodePoints_go__append
}

var cache_Data_String_NonEmpty_CodePoints_compose gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_compose sync.Once

func Get_Data_String_NonEmpty_CodePoints_compose() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_compose.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodePoints_compose
}

var cache_Data_String_NonEmpty_CodePoints_compose1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_compose1 sync.Once

func Get_Data_String_NonEmpty_CodePoints_compose1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_compose1.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodePoints_compose1
}

var cache_Data_String_NonEmpty_CodePoints_compose2 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_compose2 sync.Once

func Get_Data_String_NonEmpty_CodePoints_compose2() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_compose2.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodePoints_compose2
}

var cache_Data_String_NonEmpty_CodePoints_lessThan gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_lessThan sync.Once

func Get_Data_String_NonEmpty_CodePoints_lessThan() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_lessThan.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_String_NonEmpty_CodePoints_lessThan
}

var cache_Data_String_NonEmpty_CodePoints_fromJust gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromJust sync.Once

func Get_Data_String_NonEmpty_CodePoints_fromJust() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromJust.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromJust = Get_Data_Maybe_fromJust__123856731()
	})
	return cache_Data_String_NonEmpty_CodePoints_fromJust
}

var cache_Data_String_NonEmpty_CodePoints_fromJust1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromJust1 sync.Once

func Get_Data_String_NonEmpty_CodePoints_fromJust1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromJust1.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromJust1 = Get_Data_Maybe_fromJust__3338668507()
	})
	return cache_Data_String_NonEmpty_CodePoints_fromJust1
}

var cache_Data_String_NonEmpty_CodePoints_fromJust2 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromJust2 sync.Once

func Get_Data_String_NonEmpty_CodePoints_fromJust2() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromJust2.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromJust2 = Get_Data_Maybe_fromJust__702591195()
	})
	return cache_Data_String_NonEmpty_CodePoints_fromJust2
}

var cache_Data_String_NonEmpty_CodePoints_compose3 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_compose3 sync.Once

func Get_Data_String_NonEmpty_CodePoints_compose3() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_compose3.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodePoints_compose3
}

var cache_Data_String_NonEmpty_CodePoints_greaterThanOrEq gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_greaterThanOrEq sync.Once

func Get_Data_String_NonEmpty_CodePoints_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_greaterThanOrEq.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_String_NonEmpty_CodePoints_greaterThanOrEq
}

var cache_Data_String_NonEmpty_CodePoints_toNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toNonEmptyString sync.Once

func Get_Data_String_NonEmpty_CodePoints_toNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toNonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_toNonEmptyString(x_0_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_toNonEmptyString
}

var cache_Data_String_NonEmpty_CodePoints_snoc gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_snoc sync.Once

func Get_Data_String_NonEmpty_CodePoints_snoc() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_snoc.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_snoc(c_0_box.IntVal, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_snoc
}

var cache_Data_String_NonEmpty_CodePoints_singleton gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_singleton sync.Once

func Get_Data_String_NonEmpty_CodePoints_singleton() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_singleton.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_singleton = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), Get_Data_String_CodePoints_singleton())
	})
	return cache_Data_String_NonEmpty_CodePoints_singleton
}

var cache_Data_String_NonEmpty_CodePoints_liftS gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS sync.Once

func Get_Data_String_NonEmpty_CodePoints_liftS() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodePoints_liftS(f_0_box, v_1_box.StrVal())
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS
}

var cache_Data_String_NonEmpty_CodePoints_liftS__2946459259 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS__2946459259 sync.Once

func Get_Data_String_NonEmpty_CodePoints_liftS__2946459259() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS__2946459259.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS__2946459259 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_liftS__2946459259(f_0_box, v_1_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS__2946459259
}

var cache_Data_String_NonEmpty_CodePoints_takeWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_takeWhile sync.Once

func Get_Data_String_NonEmpty_CodePoints_takeWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_takeWhile.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_takeWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodePoints_takeWhile(f_0_box)
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_takeWhile
}

var cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_ sync.Once

func Get_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_ = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_(pat_0_box.StrVal())
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_
}

var cache_Data_String_NonEmpty_CodePoints_lastIndexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_lastIndexOf sync.Once

func Get_Data_String_NonEmpty_CodePoints_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_lastIndexOf.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_lastIndexOf = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodePoints_liftS(), Get_Data_String_CodePoints_lastIndexOf())
	})
	return cache_Data_String_NonEmpty_CodePoints_lastIndexOf
}

var cache_Data_String_NonEmpty_CodePoints_indexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_indexOf_prime_ sync.Once

func Get_Data_String_NonEmpty_CodePoints_indexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_indexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_indexOf_prime_ = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodePoints_indexOf_prime_(pat_0_box.StrVal())
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_indexOf_prime_
}

var cache_Data_String_NonEmpty_CodePoints_indexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_indexOf sync.Once

func Get_Data_String_NonEmpty_CodePoints_indexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_indexOf.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_indexOf = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodePoints_liftS(), Get_Data_String_CodePoints_indexOf())
	})
	return cache_Data_String_NonEmpty_CodePoints_indexOf
}

var cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromNonEmptyString sync.Once

func Get_Data_String_NonEmpty_CodePoints_fromNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_fromNonEmptyString(v_0_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString
}

var cache_Data_String_NonEmpty_CodePoints_length gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_length sync.Once

func Get_Data_String_NonEmpty_CodePoints_length() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_length.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_length = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_CodePoints_length(), Get_Data_String_NonEmpty_CodePoints_fromNonEmptyString())
	})
	return cache_Data_String_NonEmpty_CodePoints_length
}

var cache_Data_String_NonEmpty_CodePoints_splitAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_splitAt sync.Once

func Get_Data_String_NonEmpty_CodePoints_splitAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_splitAt.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodePoints_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(orig.after))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(orig.before))})
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_splitAt
}

var cache_Data_String_NonEmpty_CodePoints_take gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_take sync.Once

func Get_Data_String_NonEmpty_CodePoints_take() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_take.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_take(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_take
}

var cache_Data_String_NonEmpty_CodePoints_toCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toCodePointArray sync.Once

func Get_Data_String_NonEmpty_CodePoints_toCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toCodePointArray = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_CodePoints_toCodePointArray(), Get_Data_String_NonEmpty_CodePoints_fromNonEmptyString())
	})
	return cache_Data_String_NonEmpty_CodePoints_toCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray sync.Once

func Get_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_fromJust__123856731(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_fromArray(), Get_Data_String_NonEmpty_CodePoints_toCodePointArray()))
	})
	return cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_uncons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_uncons sync.Once

func Get_Data_String_NonEmpty_CodePoints_uncons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_uncons.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodePoints_uncons(nes_0_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(orig.tail))})
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_uncons
}

var cache_Data_String_NonEmpty_CodePoints_fromFoldable1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromFoldable1 sync.Once

func Get_Data_String_NonEmpty_CodePoints_fromFoldable1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromFoldable1.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodePoints_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromFoldable1
}

var cache_Data_String_NonEmpty_CodePoints_fromCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromCodePointArray sync.Once

func Get_Data_String_NonEmpty_CodePoints_fromCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromCodePointArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_fromCodePointArray(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray sync.Once

func Get_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_fromJust__702591195(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodePoints_fromCodePointArray(), Get_Data_Array_NonEmpty_toArray()))
	})
	return cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_dropWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_dropWhile sync.Once

func Get_Data_String_NonEmpty_CodePoints_dropWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_dropWhile.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_dropWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodePoints_dropWhile(f_0_box)
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_dropWhile
}

var cache_Data_String_NonEmpty_CodePoints_drop gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_drop sync.Once

func Get_Data_String_NonEmpty_CodePoints_drop() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_drop.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_drop(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_drop
}

var cache_Data_String_NonEmpty_CodePoints_countPrefix gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_countPrefix sync.Once

func Get_Data_String_NonEmpty_CodePoints_countPrefix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_countPrefix.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_countPrefix = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodePoints_liftS(), Get_Data_String_CodePoints_countPrefix())
	})
	return cache_Data_String_NonEmpty_CodePoints_countPrefix
}

var cache_Data_String_NonEmpty_CodePoints_cons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_cons sync.Once

func Get_Data_String_NonEmpty_CodePoints_cons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_cons.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_cons(c_0_box.IntVal, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodePoints_cons
}

var cache_Data_String_NonEmpty_CodePoints_codePointAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_codePointAt sync.Once

func Get_Data_String_NonEmpty_CodePoints_codePointAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_codePointAt.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_codePointAt = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodePoints_liftS(), Get_Data_String_CodePoints_codePointAt())
	})
	return cache_Data_String_NonEmpty_CodePoints_codePointAt
}

func Call_Data_String_NonEmpty_CodePoints_toNonEmptyString(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_String_NonEmpty_CodePoints_snoc(c_0_loop int64, s_1_loop string) string {
	var c_0 int64 = c_0_loop
	_ = c_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.ConcatString(s_1, gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(c_0)).StrVal())
}

func Call_Data_String_NonEmpty_CodePoints_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 string = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_CodePoints_liftS__2946459259(f_0_loop gopurs_runtime.Value, v_1_loop string) string {
liftS__2946459259:
	for {
		if false {
			continue liftS__2946459259
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var v_1 string = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).StrVal()
	}
}

func Call_Data_String_NonEmpty_CodePoints_takeWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodePoints_countPrefix(), f_0, gopurs_runtime.NormalizeStringValue(v_1)).IntVal), gopurs_runtime.NormalizeStringValue(v_1))
	}))
}

func Call_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_(pat_0_loop string) gopurs_runtime.Value {
	var pat_0 string = pat_0_loop
	_ = pat_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodePoints_liftS(), gopurs_runtime.Apply(Get_Data_String_CodePoints_lastIndexOf_prime_(), gopurs_runtime.Str(pat_0)))
}

func Call_Data_String_NonEmpty_CodePoints_indexOf_prime_(pat_0_loop string) gopurs_runtime.Value {
	var pat_0 string = pat_0_loop
	_ = pat_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodePoints_liftS(), gopurs_runtime.Apply(Get_Data_String_CodePoints_indexOf_prime_(), gopurs_runtime.Str(pat_0)))
}

func Call_Data_String_NonEmpty_CodePoints_fromNonEmptyString(v_0_loop string) string {
	var v_0 string = v_0_loop
	_ = v_0
	return v_0
}

func Call_Data_String_NonEmpty_CodePoints_splitAt(i_0_loop int64, nes_1_loop string) struct {
	after  *Constructor_Data_Maybe_Just[string]
	before *Constructor_Data_Maybe_Just[string]
	/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var nes_1 string = nes_1_loop
	_ = nes_1
	// TAST (Let): v_2_0 shape=App(Var) bindingType=(Record (Row [before: String, after: String] Empty))
	v_2_0 := Call_Data_String_CodePoints_splitAt(i_0, nes_1)
	_ = v_2_0
	var __t1 *Constructor_Data_Maybe_Just[string]
	{
		if (v_2_0.after) == ("") {
			__t1 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t1 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Str(v_2_0.after), true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_1:
	var __t2 *Constructor_Data_Maybe_Just[string]
	{
		if (v_2_0.before) == ("") {
			__t2 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Str(v_2_0.before), true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_2:
	return struct {
		after  *Constructor_Data_Maybe_Just[string]
		before *Constructor_Data_Maybe_Just[string]
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	}{__t1, __t2}
}

func Call_Data_String_NonEmpty_CodePoints_take(i_0_loop int64, nes_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var nes_1 string = nes_1_loop
	_ = nes_1
	var __t0 *Constructor_Data_Maybe_Just[string]
	{
		if (i_0) < (int64(1)) {
			__t0 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)), true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodePoints_uncons(nes_0_loop string) struct {
	head int64
	tail *Constructor_Data_Maybe_Just[string]
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
} {
	var nes_0 string = nes_0_loop
	_ = nes_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Data_String_CodePoints_codePointAt(int64(0), nes_0)
	_ = __local_var_1_0
	var __t1 int64
	{
		if __local_var_1_0.V1 {
			__t1 = __local_var_1_0.V0.IntVal
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_1:
	// TAST (Let): __local_var_1_2 shape=App(Var) bindingType=Any
	__local_var_1_2 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Str(nes_0)))).IntVal), gopurs_runtime.Str(nes_0))
	_ = __local_var_1_2
	var __t3 *Constructor_Data_Maybe_Just[string]
	{
		if (__local_var_1_2.StrVal()) == ("") {
			__t3 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{__local_var_1_2, true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_3:
	return struct {
		head int64
		tail *Constructor_Data_Maybe_Just[string]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}{__t1, __t3}
}

func Call_Data_String_NonEmpty_CodePoints_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
	_ = dictFoldable1_0
	return gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_443971153_4179793454(Rebox_Data_String_NonEmpty_CodePoints_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}, Get_Data_String_NonEmpty_CodePoints_singleton())
}

func Call_Data_String_NonEmpty_CodePoints_fromCodePointArray(v_0_loop []int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 []int64 = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[string]
	{
		if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (int64(0)) {
			__t0 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCodePointArray(), func() gopurs_runtime.Value {
				arr := v_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()), true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodePoints_dropWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(Call_Data_String_CodePoints_dropWhile(f_0, v_1.StrVal()))
	}))
}

func Call_Data_String_NonEmpty_CodePoints_drop(i_0_loop int64, nes_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var nes_1 string = nes_1_loop
	_ = nes_1
	var __t0 *Constructor_Data_Maybe_Just[string]
	{
		if (i_0) >= (Data_String_CodePoints_Length(nes_1)) {
			__t0 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)))).IntVal), gopurs_runtime.Str(nes_1)), true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodePoints_cons(c_0_loop int64, s_1_loop string) string {
	var c_0 int64 = c_0_loop
	_ = c_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.ConcatString(gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(c_0)).StrVal(), s_1)
}

func Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_String_NonEmpty_CodePoints_4179793454_443971153(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[string])(unsafe.Pointer(in))
}

func Rebox_Data_String_NonEmpty_CodePoints_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}
