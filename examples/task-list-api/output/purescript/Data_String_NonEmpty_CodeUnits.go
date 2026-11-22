package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_NonEmpty_CodeUnits_go__append gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_go__append sync.Once

func Get_Data_String_NonEmpty_CodeUnits_go__append() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_go__append.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_String_NonEmpty_CodeUnits_go__append
}

var cache_Data_String_NonEmpty_CodeUnits_compose gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_compose sync.Once

func Get_Data_String_NonEmpty_CodeUnits_compose() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_compose.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodeUnits_compose
}

var cache_Data_String_NonEmpty_CodeUnits_compose1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_compose1 sync.Once

func Get_Data_String_NonEmpty_CodeUnits_compose1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_compose1.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodeUnits_compose1
}

var cache_Data_String_NonEmpty_CodeUnits_compose2 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_compose2 sync.Once

func Get_Data_String_NonEmpty_CodeUnits_compose2() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_compose2.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodeUnits_compose2
}

var cache_Data_String_NonEmpty_CodeUnits_lessThan gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_lessThan sync.Once

func Get_Data_String_NonEmpty_CodeUnits_lessThan() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_lessThan.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_String_NonEmpty_CodeUnits_lessThan
}

var cache_Data_String_NonEmpty_CodeUnits_fromJust gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromJust sync.Once

func Get_Data_String_NonEmpty_CodeUnits_fromJust() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromJust.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromJust = Get_Data_Maybe_fromJust__1329272123()
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromJust
}

var cache_Data_String_NonEmpty_CodeUnits_fromJust1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromJust1 sync.Once

func Get_Data_String_NonEmpty_CodeUnits_fromJust1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromJust1.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromJust1 = Get_Data_Maybe_fromJust__702591195()
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromJust1
}

var cache_Data_String_NonEmpty_CodeUnits_compose3 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_compose3 sync.Once

func Get_Data_String_NonEmpty_CodeUnits_compose3() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_compose3.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_NonEmpty_CodeUnits_compose3
}

var cache_Data_String_NonEmpty_CodeUnits_greaterThanOrEq gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_greaterThanOrEq sync.Once

func Get_Data_String_NonEmpty_CodeUnits_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_greaterThanOrEq.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_String_NonEmpty_CodeUnits_greaterThanOrEq
}

var cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toNonEmptyString sync.Once

func Get_Data_String_NonEmpty_CodeUnits_toNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_toNonEmptyString(x_0_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString
}

var cache_Data_String_NonEmpty_CodeUnits_snoc gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_snoc sync.Once

func Get_Data_String_NonEmpty_CodeUnits_snoc() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_snoc.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_snoc(c_0_box.StrVal(), s_1_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_snoc
}

var cache_Data_String_NonEmpty_CodeUnits_singleton gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_singleton sync.Once

func Get_Data_String_NonEmpty_CodeUnits_singleton() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_singleton.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_singleton = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), Get_Data_String_CodeUnits_singleton())
	})
	return cache_Data_String_NonEmpty_CodeUnits_singleton
}

var cache_Data_String_NonEmpty_CodeUnits_liftS gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS sync.Once

func Get_Data_String_NonEmpty_CodeUnits_liftS() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodeUnits_liftS(f_0_box, v_1_box.StrVal())
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS
}

var cache_Data_String_NonEmpty_CodeUnits_liftS__2946459259 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS__2946459259 sync.Once

func Get_Data_String_NonEmpty_CodeUnits_liftS__2946459259() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS__2946459259.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS__2946459259 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_liftS__2946459259(f_0_box, v_1_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS__2946459259
}

var cache_Data_String_NonEmpty_CodeUnits_takeWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_takeWhile sync.Once

func Get_Data_String_NonEmpty_CodeUnits_takeWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_takeWhile.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_takeWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodeUnits_takeWhile(f_0_box)
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_takeWhile
}

var cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_ sync.Once

func Get_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_ = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_(pat_0_box.StrVal())
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_
}

var cache_Data_String_NonEmpty_CodeUnits_lastIndexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_lastIndexOf sync.Once

func Get_Data_String_NonEmpty_CodeUnits_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_lastIndexOf.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_lastIndexOf = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodeUnits_liftS(), Get_Data_String_CodeUnits_lastIndexOf())
	})
	return cache_Data_String_NonEmpty_CodeUnits_lastIndexOf
}

var cache_Data_String_NonEmpty_CodeUnits_indexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_indexOf_prime_ sync.Once

func Get_Data_String_NonEmpty_CodeUnits_indexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_indexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_indexOf_prime_ = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodeUnits_indexOf_prime_(pat_0_box.StrVal())
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_indexOf_prime_
}

var cache_Data_String_NonEmpty_CodeUnits_indexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_indexOf sync.Once

func Get_Data_String_NonEmpty_CodeUnits_indexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_indexOf.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_indexOf = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodeUnits_liftS(), Get_Data_String_CodeUnits_indexOf())
	})
	return cache_Data_String_NonEmpty_CodeUnits_indexOf
}

var cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromNonEmptyString sync.Once

func Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(v_0_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString
}

var cache_Data_String_NonEmpty_CodeUnits_length gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_length sync.Once

func Get_Data_String_NonEmpty_CodeUnits_length() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_length.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_length = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_CodeUnits_length(), Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString())
	})
	return cache_Data_String_NonEmpty_CodeUnits_length
}

var cache_Data_String_NonEmpty_CodeUnits_splitAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_splitAt sync.Once

func Get_Data_String_NonEmpty_CodeUnits_splitAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_splitAt.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodeUnits_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(orig.after))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(orig.before))})
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_splitAt
}

var cache_Data_String_NonEmpty_CodeUnits_take gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_take sync.Once

func Get_Data_String_NonEmpty_CodeUnits_take() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_take.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_take(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_take
}

var cache_Data_String_NonEmpty_CodeUnits_takeRight gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_takeRight sync.Once

func Get_Data_String_NonEmpty_CodeUnits_takeRight() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_takeRight.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_takeRight(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_takeRight
}

var cache_Data_String_NonEmpty_CodeUnits_toChar gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toChar sync.Once

func Get_Data_String_NonEmpty_CodeUnits_toChar() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toChar.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toChar = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_CodeUnits_toChar(), Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString())
	})
	return cache_Data_String_NonEmpty_CodeUnits_toChar
}

var cache_Data_String_NonEmpty_CodeUnits_toCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toCharArray sync.Once

func Get_Data_String_NonEmpty_CodeUnits_toCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toCharArray = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_CodeUnits_toCharArray(), Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString())
	})
	return cache_Data_String_NonEmpty_CodeUnits_toCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray sync.Once

func Get_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_fromJust__1329272123(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_fromArray(), Get_Data_String_NonEmpty_CodeUnits_toCharArray()))
	})
	return cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_uncons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_uncons sync.Once

func Get_Data_String_NonEmpty_CodeUnits_uncons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_uncons.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodeUnits_uncons(nes_0_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(orig.tail))})
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_uncons
}

var cache_Data_String_NonEmpty_CodeUnits_fromFoldable1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromFoldable1 sync.Once

func Get_Data_String_NonEmpty_CodeUnits_fromFoldable1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromFoldable1.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodeUnits_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromFoldable1
}

var cache_Data_String_NonEmpty_CodeUnits_fromCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromCharArray sync.Once

func Get_Data_String_NonEmpty_CodeUnits_fromCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromCharArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_fromCharArray(func() []string {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
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
	return cache_Data_String_NonEmpty_CodeUnits_fromCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray sync.Once

func Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_fromJust__702591195(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodeUnits_fromCharArray(), Get_Data_Array_NonEmpty_toArray()))
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_dropWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_dropWhile sync.Once

func Get_Data_String_NonEmpty_CodeUnits_dropWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_dropWhile.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_dropWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_NonEmpty_CodeUnits_dropWhile(f_0_box)
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_dropWhile
}

var cache_Data_String_NonEmpty_CodeUnits_dropRight gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_dropRight sync.Once

func Get_Data_String_NonEmpty_CodeUnits_dropRight() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_dropRight.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_dropRight(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_dropRight
}

var cache_Data_String_NonEmpty_CodeUnits_drop gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_drop sync.Once

func Get_Data_String_NonEmpty_CodeUnits_drop() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_drop.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_drop(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_drop
}

var cache_Data_String_NonEmpty_CodeUnits_countPrefix gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_countPrefix sync.Once

func Get_Data_String_NonEmpty_CodeUnits_countPrefix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_countPrefix.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_countPrefix = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodeUnits_liftS(), Get_Data_String_CodeUnits_countPrefix())
	})
	return cache_Data_String_NonEmpty_CodeUnits_countPrefix
}

var cache_Data_String_NonEmpty_CodeUnits_cons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_cons sync.Once

func Get_Data_String_NonEmpty_CodeUnits_cons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_cons.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_cons(c_0_box.StrVal(), s_1_box.StrVal()))
		})
	})
	return cache_Data_String_NonEmpty_CodeUnits_cons
}

var cache_Data_String_NonEmpty_CodeUnits_charAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_charAt sync.Once

func Get_Data_String_NonEmpty_CodeUnits_charAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_charAt.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_charAt = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodeUnits_liftS(), Get_Data_String_CodeUnits_charAt())
	})
	return cache_Data_String_NonEmpty_CodeUnits_charAt
}

func Call_Data_String_NonEmpty_CodeUnits_toNonEmptyString(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_String_NonEmpty_CodeUnits_snoc(c_0_loop string, s_1_loop string) string {
	var c_0 string = c_0_loop
	_ = c_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.ConcatString(s_1, Data_String_CodeUnits_Singleton(c_0))
}

func Call_Data_String_NonEmpty_CodeUnits_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 string = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_CodeUnits_liftS__2946459259(f_0_loop gopurs_runtime.Value, v_1_loop string) string {
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

func Call_Data_String_NonEmpty_CodeUnits_takeWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(Data_String_CodeUnits_CountPrefix_nativeWorker(f_0, v_1.StrVal()).IntVal), gopurs_runtime.NormalizeStringValue(v_1))
	}))
}

func Call_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_(pat_0_loop string) gopurs_runtime.Value {
	var pat_0 string = pat_0_loop
	_ = pat_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodeUnits_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_lastIndexOf_prime_(), gopurs_runtime.Str(pat_0)))
}

func Call_Data_String_NonEmpty_CodeUnits_indexOf_prime_(pat_0_loop string) gopurs_runtime.Value {
	var pat_0 string = pat_0_loop
	_ = pat_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_CodeUnits_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf_prime_(), gopurs_runtime.Str(pat_0)))
}

func Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(v_0_loop string) string {
	var v_0 string = v_0_loop
	_ = v_0
	return v_0
}

func Call_Data_String_NonEmpty_CodeUnits_splitAt(i_0_loop int64, nes_1_loop string) struct {
	after  *Constructor_Data_Maybe_Just[string]
	before *Constructor_Data_Maybe_Just[string]
	/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var nes_1 string = nes_1_loop
	_ = nes_1
	// TAST (Let): v_2_0 shape=App(Var) bindingType=(Record (Row [before: String, after: String] Empty))
	v_2_0 := func() struct {
		after  string
		before string
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	} {
		orig := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
		_ = orig
		clone := struct {
			after  string
			before string
			/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
		}{}
		clone.after = gopurs_runtime.RecordGet(orig, "after").StrVal()
		clone.before = gopurs_runtime.RecordGet(orig, "before").StrVal()
		return clone
	}()
	_ = v_2_0
	var __t1 *Constructor_Data_Maybe_Just[string]
	{
		if (v_2_0.after) == ("") {
			__t1 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t1 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
			__t2 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t2 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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

func Call_Data_String_NonEmpty_CodeUnits_take(i_0_loop int64, nes_1_loop string) struct {
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
			__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)), true}
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
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodeUnits_takeRight(i_0_loop int64, nes_1_loop string) struct {
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
			__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal, i_0)), gopurs_runtime.Str(nes_1)), true}
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
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodeUnits_uncons(nes_0_loop string) struct {
	head string
	tail *Constructor_Data_Maybe_Just[string]
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
} {
	var nes_0 string = nes_0_loop
	_ = nes_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Str(nes_0))
	_ = __local_var_1_0
	var __t1 *Constructor_Data_Maybe_Just[string]
	{
		if (__local_var_1_0.StrVal()) == ("") {
			__t1 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t1 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{__local_var_1_0, true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_1:
	return struct {
		head string
		tail *Constructor_Data_Maybe_Just[string]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}{gopurs_runtime.Apply(Data_String_Unsafe_CharAt_nativeWorker(gopurs_runtime.Int(int64(0))), gopurs_runtime.Str(nes_0)).StrVal(), __t1}
}

func Call_Data_String_NonEmpty_CodeUnits_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
	_ = dictFoldable1_0
	return gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_443971153_4179793454(Rebox_Data_String_NonEmpty_CodeUnits_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}, Get_Data_String_NonEmpty_CodeUnits_singleton())
}

func Call_Data_String_NonEmpty_CodeUnits_fromCharArray(v_0_loop []string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 []string = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[string]
	{
		if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (int64(0)) {
			__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Str(Data_String_CodeUnits_FromCharArray(v_0)), true}
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
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodeUnits_dropWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(Data_String_CodeUnits_CountPrefix_nativeWorker(f_0, v_1.StrVal()).IntVal), gopurs_runtime.NormalizeStringValue(v_1))
	}))
}

func Call_Data_String_NonEmpty_CodeUnits_dropRight(i_0_loop int64, nes_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var nes_1 string = nes_1_loop
	_ = nes_1
	var __t0 *Constructor_Data_Maybe_Just[string]
	{
		if (i_0) >= (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) {
			__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal, i_0)), gopurs_runtime.Str(nes_1)), true}
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
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodeUnits_drop(i_0_loop int64, nes_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var nes_1 string = nes_1_loop
	_ = nes_1
	var __t0 *Constructor_Data_Maybe_Just[string]
	{
		if (i_0) >= (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) {
			__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)), true}
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
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_NonEmpty_CodeUnits_cons(c_0_loop string, s_1_loop string) string {
	var c_0 string = c_0_loop
	_ = c_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.ConcatString(Data_String_CodeUnits_Singleton(c_0), s_1)
}

func Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_String_NonEmpty_CodeUnits_4179793454_443971153(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[string])(unsafe.Pointer(in))
}

func Rebox_Data_String_NonEmpty_CodeUnits_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}
