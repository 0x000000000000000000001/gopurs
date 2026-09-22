package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_eqRec gopurs_runtime.Value
var once_Main_eqRec sync.Once

func Get_Main_eqRec() gopurs_runtime.Value {
	once_Main_eqRec.Do(func() {
		cache_Main_eqRec = gopurs_runtime.Apply(Get_Data_Eq_eqRec(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRec
}

var cache_Main_eqRowCons gopurs_runtime.Value
var once_Main_eqRowCons sync.Once

func Get_Main_eqRowCons() gopurs_runtime.Value {
	once_Main_eqRowCons.Do(func() {
		cache_Main_eqRowCons = gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRowCons
}

var cache_Main_typeIsSymbol gopurs_runtime.Value
var once_Main_typeIsSymbol sync.Once

func Get_Main_typeIsSymbol() gopurs_runtime.Value {
	once_Main_typeIsSymbol.Do(func() {
		cache_Main_typeIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("type")
		}))
	})
	return cache_Main_typeIsSymbol
}

var cache_Main_limitIsSymbol gopurs_runtime.Value
var once_Main_limitIsSymbol sync.Once

func Get_Main_limitIsSymbol() gopurs_runtime.Value {
	once_Main_limitIsSymbol.Do(func() {
		cache_Main_limitIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("limit")
		}))
	})
	return cache_Main_limitIsSymbol
}

var cache_Main_closeIsSymbol gopurs_runtime.Value
var once_Main_closeIsSymbol sync.Once

func Get_Main_closeIsSymbol() gopurs_runtime.Value {
	once_Main_closeIsSymbol.Do(func() {
		cache_Main_closeIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("close")
		}))
	})
	return cache_Main_closeIsSymbol
}

var cache_Main_afterIsSymbol gopurs_runtime.Value
var once_Main_afterIsSymbol sync.Once

func Get_Main_afterIsSymbol() gopurs_runtime.Value {
	once_Main_afterIsSymbol.Do(func() {
		cache_Main_afterIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("after")
		}))
	})
	return cache_Main_afterIsSymbol
}

var cache_Main_identifierIsSymbol gopurs_runtime.Value
var once_Main_identifierIsSymbol sync.Once

func Get_Main_identifierIsSymbol() gopurs_runtime.Value {
	once_Main_identifierIsSymbol.Do(func() {
		cache_Main_identifierIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("identifier")
		}))
	})
	return cache_Main_identifierIsSymbol
}

var cache_Main_eqRec1 gopurs_runtime.Value
var once_Main_eqRec1 sync.Once

func Get_Main_eqRec1() gopurs_runtime.Value {
	once_Main_eqRec1.Do(func() {
		cache_Main_eqRec1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3039522309_3790796878(Rebox_Main_3790796878_3039522309(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_typeIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), gopurs_runtime.Value{}, Get_Main_limitIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_closeIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}), gopurs_runtime.Value{}, Get_Main_afterIsSymbol(), Call_Data_Maybe_eqMaybe(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_identifierIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))))))}
	})
	return cache_Main_eqRec1
}

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Apply2(Get_Data_Show_showRecord(), gopurs_runtime.Value{}, gopurs_runtime.Value{})
	})
	return cache_Main_showRecord
}

var cache_Main_showRecord1 gopurs_runtime.Value
var once_Main_showRecord1 sync.Once

func Get_Main_showRecord1() gopurs_runtime.Value {
	once_Main_showRecord1.Do(func() {
		cache_Main_showRecord1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2664792997_1386611502(Rebox_Main_1386611502_2664792997(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_afterIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_closeIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_limitIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_typeIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}), Call_Data_Maybe_showMaybe(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_identifierIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))))))))}
	})
	return cache_Main_showRecord1
}

var cache_Main_convert gopurs_runtime.Value
var once_Main_convert sync.Once

func Get_Main_convert() gopurs_runtime.Value {
	once_Main_convert.Do(func() {
		cache_Main_convert = gopurs_runtime.Func(func(options_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_convert(func() struct {
					after     *Constructor_Data_Maybe_Just[string]
					go__close bool
					limit     int64
					go__type  string
				} {
					orig := options_0_box
					_ = orig
					clone := struct {
						after     *Constructor_Data_Maybe_Just[string]
						go__close bool
						limit     int64
						go__type  string
					}{}
					clone.after = Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "after")))
					clone.go__close = (gopurs_runtime.RecordGet(orig, "close").IntVal) != (0)
					clone.limit = gopurs_runtime.RecordGet(orig, "limit").IntVal
					clone.go__type = gopurs_runtime.RecordGet(orig, "type").StrVal()
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict4("after", "close", "limit", "type", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_3543310304_3094389156(orig.after))}, gopurs_runtime.Bool(orig.go__close), gopurs_runtime.Int(orig.limit), gopurs_runtime.Str(orig.go__type))
			}()
		})
	})
	return cache_Main_convert
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					after     *Constructor_Data_Maybe_Just[string]
					go__close bool
					limit     int64
					go__type  string
				}{Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Str("book-42"), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())), true, int64(10), "book"}
				_ = orig
				return gopurs_runtime.RecordDict4("after", "close", "limit", "type", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_742090555_3094389156(orig.after))}, gopurs_runtime.Bool(orig.go__close), gopurs_runtime.Int(orig.limit), gopurs_runtime.Str(orig.go__type))
			}())
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1043860040("", struct {
				actual struct {
					after *Constructor_Data_Maybe_Just[struct {
						identifier string
					}]
					go__close bool
					limit     int64
					go__type  string
				}
				expected struct {
					after *Constructor_Data_Maybe_Just[struct {
						identifier string
					}]
					go__close bool
					limit     int64
					go__type  string
				}
			}{Call_Main_convert(func() struct {
				after     *Constructor_Data_Maybe_Just[string]
				go__close bool
				limit     int64
				go__type  string
			} {
				orig := __local_var_2_2
				_ = orig
				clone := struct {
					after     *Constructor_Data_Maybe_Just[string]
					go__close bool
					limit     int64
					go__type  string
				}{}
				clone.after = Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "after")))
				clone.go__close = (gopurs_runtime.RecordGet(orig, "close").IntVal) != (0)
				clone.limit = gopurs_runtime.RecordGet(orig, "limit").IntVal
				clone.go__type = gopurs_runtime.RecordGet(orig, "type").StrVal()
				return clone
			}()), struct {
				after *Constructor_Data_Maybe_Just[struct {
					identifier string
				}]
				go__close bool
				limit     int64
				go__type  string
			}{Rebox_Main_3094389156_3543310304(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						identifier string
					}
					V1 bool
				}{struct {
					identifier string
				}{"book-42"}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict1("identifier", gopurs_runtime.Str(orig.identifier))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())), true, int64(10), "book"}}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_convert(options_0_loop struct {
	after     *Constructor_Data_Maybe_Just[string]
	go__close bool
	limit     int64
	go__type  string
}) struct {
	after *Constructor_Data_Maybe_Just[struct {
		identifier string
	}]
	go__close bool
	limit     int64
	go__type  string
} {
	var options_0 struct {
		after     *Constructor_Data_Maybe_Just[string]
		go__close bool
		limit     int64
		go__type  string
	} = options_0_loop
	_ = options_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		var __t_tag_0 *Constructor_Data_Maybe_Just[string] = options_0.after
		_ = __t_tag_0
		if __t_tag_0 != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{func() gopurs_runtime.Value {
					orig := struct {
						identifier string
					}{gopurs_runtime.Str((options_0.after).V0).StrVal()}
					_ = orig
					return gopurs_runtime.RecordDict1("identifier", gopurs_runtime.Str(orig.identifier))
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
		after *Constructor_Data_Maybe_Just[struct {
			identifier string
		}]
		go__close bool
		limit     int64
		go__type  string
	} {
		originalRecord := options_0
		_ = originalRecord
		_ = originalRecord
		var clone struct {
			after *Constructor_Data_Maybe_Just[struct {
				identifier string
			}]
			go__close bool
			limit     int64
			go__type  string
		}
		clone.go__close = originalRecord.go__close
		clone.limit = originalRecord.limit
		clone.go__type = originalRecord.go__type
		clone.after = Rebox_Main_3094389156_3543310304(__t1)
		return clone
	}()
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2664792997(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	after *Constructor_Data_Maybe_Just[struct {
		identifier string
	}]
	go__close bool
	limit     int64
	go__type  string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		after *Constructor_Data_Maybe_Just[struct {
			identifier string
		}]
		go__close bool
		limit     int64
		go__type  string
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2664792997_1386611502(in *Constructor_Data_Show_Show[struct {
	after *Constructor_Data_Maybe_Just[struct {
		identifier string
	}]
	go__close bool
	limit     int64
	go__type  string
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3039522309_3790796878(in *Constructor_Data_Eq_Eq[struct {
	after *Constructor_Data_Maybe_Just[struct {
		identifier string
	}]
	go__close bool
	limit     int64
	go__type  string
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3094389156_3543310304(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	identifier string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		identifier string
	}]{}
	out.V0 = func() struct {
		identifier string
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			identifier string
		}{}
		clone.identifier = gopurs_runtime.RecordGet(orig, "identifier").StrVal()
		return clone
	}()
	return out
}

func Rebox_Main_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Main_3543310304_3094389156(in *Constructor_Data_Maybe_Just[struct {
	identifier string
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict1("identifier", gopurs_runtime.Str(orig.identifier))
	}()
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

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_3039522309(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	after *Constructor_Data_Maybe_Just[struct {
		identifier string
	}]
	go__close bool
	limit     int64
	go__type  string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		after *Constructor_Data_Maybe_Just[struct {
			identifier string
		}]
		go__close bool
		limit     int64
		go__type  string
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}
