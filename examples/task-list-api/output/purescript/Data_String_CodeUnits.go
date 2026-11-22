package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_CodeUnits_zero gopurs_runtime.Value
var once_Data_String_CodeUnits_zero sync.Once

func Get_Data_String_CodeUnits_zero() gopurs_runtime.Value {
	once_Data_String_CodeUnits_zero.Do(func() {
		cache_Data_String_CodeUnits_zero = gopurs_runtime.Int(int64(0))
	})
	return cache_Data_String_CodeUnits_zero
}

var cache_Data_String_CodeUnits_one gopurs_runtime.Value
var once_Data_String_CodeUnits_one sync.Once

func Get_Data_String_CodeUnits_one() gopurs_runtime.Value {
	once_Data_String_CodeUnits_one.Do(func() {
		cache_Data_String_CodeUnits_one = gopurs_runtime.Int(int64(1))
	})
	return cache_Data_String_CodeUnits_one
}

var cache_Data_String_CodeUnits_sub gopurs_runtime.Value
var once_Data_String_CodeUnits_sub sync.Once

func Get_Data_String_CodeUnits_sub() gopurs_runtime.Value {
	once_Data_String_CodeUnits_sub.Do(func() {
		cache_Data_String_CodeUnits_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_String_CodeUnits_sub
}

var cache_Data_String_CodeUnits_eq gopurs_runtime.Value
var once_Data_String_CodeUnits_eq sync.Once

func Get_Data_String_CodeUnits_eq() gopurs_runtime.Value {
	once_Data_String_CodeUnits_eq.Do(func() {
		cache_Data_String_CodeUnits_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Data_String_CodeUnits_eq
}

var cache_Data_String_CodeUnits_compose gopurs_runtime.Value
var once_Data_String_CodeUnits_compose sync.Once

func Get_Data_String_CodeUnits_compose() gopurs_runtime.Value {
	once_Data_String_CodeUnits_compose.Do(func() {
		cache_Data_String_CodeUnits_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_CodeUnits_compose
}

var cache_Data_String_CodeUnits_compose1 gopurs_runtime.Value
var once_Data_String_CodeUnits_compose1 sync.Once

func Get_Data_String_CodeUnits_compose1() gopurs_runtime.Value {
	once_Data_String_CodeUnits_compose1.Do(func() {
		cache_Data_String_CodeUnits_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_CodeUnits_compose1
}

var cache_Data_String_CodeUnits_compose2 gopurs_runtime.Value
var once_Data_String_CodeUnits_compose2 sync.Once

func Get_Data_String_CodeUnits_compose2() gopurs_runtime.Value {
	once_Data_String_CodeUnits_compose2.Do(func() {
		cache_Data_String_CodeUnits_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_String_CodeUnits_compose2
}

var cache_Data_String_CodeUnits_uncons gopurs_runtime.Value
var once_Data_String_CodeUnits_uncons sync.Once

func Get_Data_String_CodeUnits_uncons() gopurs_runtime.Value {
	once_Data_String_CodeUnits_uncons.Do(func() {
		cache_Data_String_CodeUnits_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodeUnits_uncons(v_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Str(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodeUnits_uncons
}

var cache_Data_String_CodeUnits_toChar gopurs_runtime.Value
var once_Data_String_CodeUnits_toChar sync.Once

func Get_Data_String_CodeUnits_toChar() gopurs_runtime.Value {
	once_Data_String_CodeUnits_toChar.Do(func() {
		cache_Data_String_CodeUnits_toChar = gopurs_runtime.Apply2(Get_Data_String_CodeUnits__toChar(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_CodeUnits_toChar
}

var cache_Data_String_CodeUnits_takeWhile gopurs_runtime.Value
var once_Data_String_CodeUnits_takeWhile sync.Once

func Get_Data_String_CodeUnits_takeWhile() gopurs_runtime.Value {
	once_Data_String_CodeUnits_takeWhile.Do(func() {
		cache_Data_String_CodeUnits_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodeUnits_takeWhile(p_0_box, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodeUnits_takeWhile
}

var cache_Data_String_CodeUnits_takeRight gopurs_runtime.Value
var once_Data_String_CodeUnits_takeRight sync.Once

func Get_Data_String_CodeUnits_takeRight() gopurs_runtime.Value {
	once_Data_String_CodeUnits_takeRight.Do(func() {
		cache_Data_String_CodeUnits_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodeUnits_takeRight(i_0_box.IntVal, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodeUnits_takeRight
}

var cache_Data_String_CodeUnits_stripSuffix gopurs_runtime.Value
var once_Data_String_CodeUnits_stripSuffix sync.Once

func Get_Data_String_CodeUnits_stripSuffix() gopurs_runtime.Value {
	once_Data_String_CodeUnits_stripSuffix.Do(func() {
		cache_Data_String_CodeUnits_stripSuffix = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, str_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodeUnits_stripSuffix(v_0_box.StrVal(), str_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodeUnits_stripSuffix
}

var cache_Data_String_CodeUnits_stripPrefix gopurs_runtime.Value
var once_Data_String_CodeUnits_stripPrefix sync.Once

func Get_Data_String_CodeUnits_stripPrefix() gopurs_runtime.Value {
	once_Data_String_CodeUnits_stripPrefix.Do(func() {
		cache_Data_String_CodeUnits_stripPrefix = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, str_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_CodeUnits_stripPrefix(v_0_box.StrVal(), str_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_String_CodeUnits_stripPrefix
}

var cache_Data_String_CodeUnits_startsWith gopurs_runtime.Value
var once_Data_String_CodeUnits_startsWith sync.Once

func Get_Data_String_CodeUnits_startsWith() gopurs_runtime.Value {
	once_Data_String_CodeUnits_startsWith.Do(func() {
		cache_Data_String_CodeUnits_startsWith = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_CodeUnits_startsWith(pat_0_box.StrVal())
		})
	})
	return cache_Data_String_CodeUnits_startsWith
}

var cache_Data_String_CodeUnits_lastIndexOf_prime_ gopurs_runtime.Value
var once_Data_String_CodeUnits_lastIndexOf_prime_ sync.Once

func Get_Data_String_CodeUnits_lastIndexOf_prime_() gopurs_runtime.Value {
	once_Data_String_CodeUnits_lastIndexOf_prime_.Do(func() {
		cache_Data_String_CodeUnits_lastIndexOf_prime_ = gopurs_runtime.Apply2(Get_Data_String_CodeUnits__lastIndexOfStartingAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_CodeUnits_lastIndexOf_prime_
}

var cache_Data_String_CodeUnits_lastIndexOf gopurs_runtime.Value
var once_Data_String_CodeUnits_lastIndexOf sync.Once

func Get_Data_String_CodeUnits_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_CodeUnits_lastIndexOf.Do(func() {
		cache_Data_String_CodeUnits_lastIndexOf = gopurs_runtime.Apply2(Get_Data_String_CodeUnits__lastIndexOf(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_CodeUnits_lastIndexOf
}

var cache_Data_String_CodeUnits_indexOf_prime_ gopurs_runtime.Value
var once_Data_String_CodeUnits_indexOf_prime_ sync.Once

func Get_Data_String_CodeUnits_indexOf_prime_() gopurs_runtime.Value {
	once_Data_String_CodeUnits_indexOf_prime_.Do(func() {
		cache_Data_String_CodeUnits_indexOf_prime_ = gopurs_runtime.Apply2(Get_Data_String_CodeUnits__indexOfStartingAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_CodeUnits_indexOf_prime_
}

var cache_Data_String_CodeUnits_indexOf gopurs_runtime.Value
var once_Data_String_CodeUnits_indexOf sync.Once

func Get_Data_String_CodeUnits_indexOf() gopurs_runtime.Value {
	once_Data_String_CodeUnits_indexOf.Do(func() {
		cache_Data_String_CodeUnits_indexOf = gopurs_runtime.Apply2(Get_Data_String_CodeUnits__indexOf(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_CodeUnits_indexOf
}

var cache_Data_String_CodeUnits_endsWith gopurs_runtime.Value
var once_Data_String_CodeUnits_endsWith sync.Once

func Get_Data_String_CodeUnits_endsWith() gopurs_runtime.Value {
	once_Data_String_CodeUnits_endsWith.Do(func() {
		cache_Data_String_CodeUnits_endsWith = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_CodeUnits_endsWith(pat_0_box.StrVal())
		})
	})
	return cache_Data_String_CodeUnits_endsWith
}

var cache_Data_String_CodeUnits_dropWhile gopurs_runtime.Value
var once_Data_String_CodeUnits_dropWhile sync.Once

func Get_Data_String_CodeUnits_dropWhile() gopurs_runtime.Value {
	once_Data_String_CodeUnits_dropWhile.Do(func() {
		cache_Data_String_CodeUnits_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodeUnits_dropWhile(p_0_box, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodeUnits_dropWhile
}

var cache_Data_String_CodeUnits_dropRight gopurs_runtime.Value
var once_Data_String_CodeUnits_dropRight sync.Once

func Get_Data_String_CodeUnits_dropRight() gopurs_runtime.Value {
	once_Data_String_CodeUnits_dropRight.Do(func() {
		cache_Data_String_CodeUnits_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CodeUnits_dropRight(i_0_box.IntVal, s_1_box.StrVal()))
		})
	})
	return cache_Data_String_CodeUnits_dropRight
}

var cache_Data_String_CodeUnits_contains gopurs_runtime.Value
var once_Data_String_CodeUnits_contains sync.Once

func Get_Data_String_CodeUnits_contains() gopurs_runtime.Value {
	once_Data_String_CodeUnits_contains.Do(func() {
		cache_Data_String_CodeUnits_contains = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_CodeUnits_contains(pat_0_box.StrVal())
		})
	})
	return cache_Data_String_CodeUnits_contains
}

var cache_Data_String_CodeUnits_charAt gopurs_runtime.Value
var once_Data_String_CodeUnits_charAt sync.Once

func Get_Data_String_CodeUnits_charAt() gopurs_runtime.Value {
	once_Data_String_CodeUnits_charAt.Do(func() {
		cache_Data_String_CodeUnits_charAt = gopurs_runtime.Apply2(Get_Data_String_CodeUnits__charAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_CodeUnits_charAt
}

func Call_Data_String_CodeUnits_uncons(v_0_loop string) struct {
	V0 struct {
		head string
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V1 bool
} {
	var v_0 string = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[struct {
		head string
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]
	{
		if (v_0) == ("") {
			__t0 = Rebox_Data_String_CodeUnits_3094389156_2377819597(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						head string
						tail string
						/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
					}
					V1 bool
				}{struct {
					head string
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Str(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Rebox_Data_String_CodeUnits_3094389156_2377819597(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 struct {
					head string
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V1 bool
			}{struct {
				head string
				tail string
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}{gopurs_runtime.Apply(Data_String_Unsafe_CharAt_nativeWorker(gopurs_runtime.Int(int64(0))), gopurs_runtime.Str(v_0)).StrVal(), gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Str(v_0)).StrVal()}, true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Str(orig.tail))
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 struct {
			head string
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}
		V1 bool
	} {
		_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodeUnits_2377819597_3094389156(__t0))}
		if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 struct {
					head string
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V1 bool
			}{V0: func() struct {
				head string
				tail string
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			} {
				orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					head string
					tail string
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}
				clone.head = gopurs_runtime.RecordGet(orig, "head").StrVal()
				clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
				return clone
			}(), V1: true}
		}
		return struct {
			V0 struct {
				head string
				tail string
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V1 bool
		}{V0: struct {
			head string
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, V1: false}
	}()
}

func Call_Data_String_CodeUnits_takeWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(Data_String_CodeUnits_CountPrefix_nativeWorker(p_0, s_1).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodeUnits_takeRight(i_0_loop int64, s_1_loop string) string {
	var i_0 int64 = i_0_loop
	_ = i_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(s_1)).IntVal, i_0)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodeUnits_stripSuffix(v_0_loop string, str_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 string = v_0_loop
	_ = v_0
	var str_1 string = str_1_loop
	_ = str_1
	// TAST (Let): v1_2_0 shape=App(Var) bindingType=(Record (Row [before: String, after: String] Empty))
	v1_2_0 := func() struct {
		after  string
		before string
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	} {
		orig := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_splitAt(), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(str_1)).IntVal, gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(v_0)).IntVal)), gopurs_runtime.Str(str_1))
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
	_ = v1_2_0
	var __t1 *Constructor_Data_Maybe_Just[string]
	{
		if (v1_2_0.after) == (v_0) {
			__t1 = Rebox_Data_String_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Str(v1_2_0.before), true}
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
		__t1 = Rebox_Data_String_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodeUnits_742090555_3094389156(__t1))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_CodeUnits_stripPrefix(v_0_loop string, str_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 string = v_0_loop
	_ = v_0
	var str_1 string = str_1_loop
	_ = str_1
	// TAST (Let): v1_2_0 shape=App(Var) bindingType=(Record (Row [before: String, after: String] Empty))
	v1_2_0 := func() struct {
		after  string
		before string
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	} {
		orig := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_splitAt(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(v_0)).IntVal), gopurs_runtime.Str(str_1))
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
	_ = v1_2_0
	var __t1 *Constructor_Data_Maybe_Just[string]
	{
		if (v1_2_0.before) == (v_0) {
			__t1 = Rebox_Data_String_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Str(v1_2_0.after), true}
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
		__t1 = Rebox_Data_String_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CodeUnits_742090555_3094389156(__t1))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_String_CodeUnits_startsWith(pat_0_loop string) gopurs_runtime.Value {
	var pat_0 string = pat_0_loop
	_ = pat_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_stripPrefix(), gopurs_runtime.Str(pat_0)))
}

func Call_Data_String_CodeUnits_endsWith(pat_0_loop string) gopurs_runtime.Value {
	var pat_0 string = pat_0_loop
	_ = pat_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_stripSuffix(), gopurs_runtime.Str(pat_0)))
}

func Call_Data_String_CodeUnits_dropWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(Data_String_CodeUnits_CountPrefix_nativeWorker(p_0, s_1).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodeUnits_dropRight(i_0_loop int64, s_1_loop string) string {
	var i_0 int64 = i_0_loop
	_ = i_0
	var s_1 string = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(s_1)).IntVal, i_0)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodeUnits_contains(pat_0_loop string) gopurs_runtime.Value {
	var pat_0 string = pat_0_loop
	_ = pat_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str(pat_0)))
}

func Rebox_Data_String_CodeUnits_2377819597_3094389156(in *Constructor_Data_Maybe_Just[struct {
	head string
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
		return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Str(orig.tail))
	}()
	return out
}

func Rebox_Data_String_CodeUnits_3094389156_2377819597(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	head string
	tail string
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		head string
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]{}
	out.V0 = func() struct {
		head string
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			head string
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}
		clone.head = gopurs_runtime.RecordGet(orig, "head").StrVal()
		clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
		return clone
	}()
	return out
}

func Rebox_Data_String_CodeUnits_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_String_CodeUnits_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Get_Data_String_CodeUnits__charAt() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits__CharAt
}

func Get_Data_String_CodeUnits__indexOf() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits__IndexOf
}

func Get_Data_String_CodeUnits__indexOfStartingAt() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits__IndexOfStartingAt
}

func Get_Data_String_CodeUnits__lastIndexOf() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits__LastIndexOf
}

func Get_Data_String_CodeUnits__lastIndexOfStartingAt() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits__LastIndexOfStartingAt
}

func Get_Data_String_CodeUnits__toChar() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits__ToChar
}

func Get_Data_String_CodeUnits_countPrefix() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_CountPrefix
}

func Get_Data_String_CodeUnits_drop() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_Drop
}

func Get_Data_String_CodeUnits_fromCharArray() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_FromCharArray
}

func Get_Data_String_CodeUnits_length() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_Length
}

func Get_Data_String_CodeUnits_singleton() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_Singleton
}

func Get_Data_String_CodeUnits_slice() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_Slice
}

func Get_Data_String_CodeUnits_splitAt() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_SplitAt
}

func Get_Data_String_CodeUnits_take() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_Take
}

func Get_Data_String_CodeUnits_toCharArray() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodeUnits_ToCharArray
}
