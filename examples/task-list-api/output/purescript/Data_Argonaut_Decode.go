package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Decode_composeKleisli gopurs_runtime.Value
var once_Data_Argonaut_Decode_composeKleisli sync.Once

func Get_Data_Argonaut_Decode_composeKleisli() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_composeKleisli.Do(func() {
		cache_Data_Argonaut_Decode_composeKleisli = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_composeKleisli(f_0_box, g_1_box, a_2_box.StrVal())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_composeKleisli
}

var cache_Data_Argonaut_Decode_composeKleisli__4221964054 gopurs_runtime.Value
var once_Data_Argonaut_Decode_composeKleisli__4221964054 sync.Once

func Get_Data_Argonaut_Decode_composeKleisli__4221964054() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_composeKleisli__4221964054.Do(func() {
		cache_Data_Argonaut_Decode_composeKleisli__4221964054 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_composeKleisli__4221964054(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Argonaut_Decode_composeKleisli__4221964054
}

var cache_Data_Argonaut_Decode_composeKleisli__396278538 gopurs_runtime.Value
var once_Data_Argonaut_Decode_composeKleisli__396278538 sync.Once

func Get_Data_Argonaut_Decode_composeKleisli__396278538() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_composeKleisli__396278538.Do(func() {
		cache_Data_Argonaut_Decode_composeKleisli__396278538 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_composeKleisli__396278538(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Argonaut_Decode_composeKleisli__396278538
}

var cache_Data_Argonaut_Decode_fromJsonString gopurs_runtime.Value
var once_Data_Argonaut_Decode_fromJsonString sync.Once

func Get_Data_Argonaut_Decode_fromJsonString() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_fromJsonString.Do(func() {
		cache_Data_Argonaut_Decode_fromJsonString = gopurs_runtime.Func(func(dictDecodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_fromJsonString(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](dictDecodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Decode_fromJsonString
}

var cache_Data_Argonaut_Decode_fromJsonString__2316388992 gopurs_runtime.Value
var once_Data_Argonaut_Decode_fromJsonString__2316388992 sync.Once

func Get_Data_Argonaut_Decode_fromJsonString__2316388992() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_fromJsonString__2316388992.Do(func() {
		cache_Data_Argonaut_Decode_fromJsonString__2316388992 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_fromJsonString__2316388992(__eta0_0_box.StrVal())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict1("completed", gopurs_runtime.Bool(orig.completed))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_fromJsonString__2316388992
}

var cache_Data_Argonaut_Decode_fromJsonString__2215306652 gopurs_runtime.Value
var once_Data_Argonaut_Decode_fromJsonString__2215306652 sync.Once

func Get_Data_Argonaut_Decode_fromJsonString__2215306652() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_fromJsonString__2215306652.Do(func() {
		cache_Data_Argonaut_Decode_fromJsonString__2215306652 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_fromJsonString__2215306652(__eta0_0_box.StrVal())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict1("title", gopurs_runtime.Str(orig.title))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_fromJsonString__2215306652
}

func Call_Data_Argonaut_Decode_composeKleisli(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop string) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 string = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope48) [(TypeVar b$scope46)])
	__local_var_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Str(a_2))
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(g_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
}

func Call_Data_Argonaut_Decode_composeKleisli__4221964054(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeKleisli__4221964054:
	for {
		if false {
			continue composeKleisli__4221964054
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Record (Row [completed: Boolean] Empty))]))
		__local_var_2_0 := Call_Data_Argonaut_Decode_Class_decodeJson(Rebox_Data_Argonaut_Decode_2648973148_1277236813(Rebox_Data_Argonaut_Decode_1277236813_2648973148(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_3459698345_1277236813(Rebox_Data_Argonaut_Decode_1277236813_3459698345(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonBoolean()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_completedIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","Argonaut","Core","Json"] [])])
			__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), gopurs_runtime.NormalizeStringValue(a_3))
			_ = __local_var_4_1
			var __t2 gopurs_runtime.Value
			{
				if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 3711209382 {
					__t2 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_2
				} else {

				}
			}
			{
				if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 2465973597 {
					__t2 = gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0)
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		})
	}
}

func Call_Data_Argonaut_Decode_composeKleisli__396278538(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeKleisli__396278538:
	for {
		if false {
			continue composeKleisli__396278538
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Record (Row [title: String] Empty))]))
		__local_var_2_0 := Call_Data_Argonaut_Decode_Class_decodeJson(Rebox_Data_Argonaut_Decode_2400514730_1277236813(Rebox_Data_Argonaut_Decode_1277236813_2400514730(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_48201874_1277236813(Rebox_Data_Argonaut_Decode_1277236813_48201874(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonString()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_titleIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","Argonaut","Core","Json"] [])])
			__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), gopurs_runtime.NormalizeStringValue(a_3))
			_ = __local_var_4_1
			var __t2 gopurs_runtime.Value
			{
				if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 3711209382 {
					__t2 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_2
				} else {

				}
			}
			{
				if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 2465973597 {
					__t2 = gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0)
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		})
	}
}

func Call_Data_Argonaut_Decode_fromJsonString(dictDecodeJson_0_loop *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDecodeJson_0 *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] = dictDecodeJson_0_loop
	_ = dictDecodeJson_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar json$scope1)]))
	__local_var_1_0 := Call_Data_Argonaut_Decode_Class_decodeJson(dictDecodeJson_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope48) [(TypeVar b$scope46)])
		__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), gopurs_runtime.NormalizeStringValue(a_2))
		_ = __local_var_3_1
		var __t2 gopurs_runtime.Value
		{
			if __local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 3711209382 {
				__t2 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_2
			} else {

			}
		}
		{
			if __local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 2465973597 {
				__t2 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V0)
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_2:
		return __t2
	})
}

func Call_Data_Argonaut_Decode_fromJsonString__2316388992(__eta0_0_loop string) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		completed bool
		/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
	}
	V2 bool
} {
fromJsonString__2316388992:
	for {
		if false {
			continue fromJsonString__2316388992
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 struct {
				completed bool
				/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
			}
			V2 bool
		} {
			_v := gopurs_runtime.Apply(Call_Data_Argonaut_Decode_composeKleisli__4221964054(Get_Data_Argonaut_Decode_Parser_parseJson(), Call_Data_Argonaut_Decode_Class_decodeJson(Rebox_Data_Argonaut_Decode_2648973148_1277236813(Rebox_Data_Argonaut_Decode_1277236813_2648973148(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_3459698345_1277236813(Rebox_Data_Argonaut_Decode_1277236813_3459698345(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonBoolean()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_completedIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))), gopurs_runtime.Str(__eta0_0))
			if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 struct {
						completed bool
						/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
					}
					V2 bool
				}{V0: gopurs_runtime.Value{}, V1: func() struct {
					completed bool
					/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
				} {
					orig := (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0
					_ = orig
					clone := struct {
						completed bool
						/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
					}{}
					clone.completed = (gopurs_runtime.RecordGet(orig, "completed").IntVal) != (0)
					return clone
				}(), V2: true}
			}
			return struct {
				V0 gopurs_runtime.Value
				V1 struct {
					completed bool
					/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
				}
				V2 bool
			}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: struct {
				completed bool
				/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
			}{}, V2: false}
		}()
	}
}

func Call_Data_Argonaut_Decode_fromJsonString__2215306652(__eta0_0_loop string) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		title string
		/*gopurs:record-v1:1:007400690074006c0065*/
	}
	V2 bool
} {
fromJsonString__2215306652:
	for {
		if false {
			continue fromJsonString__2215306652
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 struct {
				title string
				/*gopurs:record-v1:1:007400690074006c0065*/
			}
			V2 bool
		} {
			_v := gopurs_runtime.Apply(Call_Data_Argonaut_Decode_composeKleisli__396278538(Get_Data_Argonaut_Decode_Parser_parseJson(), Call_Data_Argonaut_Decode_Class_decodeJson(Rebox_Data_Argonaut_Decode_2400514730_1277236813(Rebox_Data_Argonaut_Decode_1277236813_2400514730(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_48201874_1277236813(Rebox_Data_Argonaut_Decode_1277236813_48201874(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonString()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_titleIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))), gopurs_runtime.Str(__eta0_0))
			if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 struct {
						title string
						/*gopurs:record-v1:1:007400690074006c0065*/
					}
					V2 bool
				}{V0: gopurs_runtime.Value{}, V1: func() struct {
					title string
					/*gopurs:record-v1:1:007400690074006c0065*/
				} {
					orig := (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0
					_ = orig
					clone := struct {
						title string
						/*gopurs:record-v1:1:007400690074006c0065*/
					}{}
					clone.title = gopurs_runtime.RecordGet(orig, "title").StrVal()
					return clone
				}(), V2: true}
			}
			return struct {
				V0 gopurs_runtime.Value
				V1 struct {
					title string
					/*gopurs:record-v1:1:007400690074006c0065*/
				}
				V2 bool
			}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: struct {
				title string
				/*gopurs:record-v1:1:007400690074006c0065*/
			}{}, V2: false}
		}()
	}
}

func Rebox_Data_Argonaut_Decode_1277236813_2400514730(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	title string
	/*gopurs:record-v1:1:007400690074006c0065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		title string
		/*gopurs:record-v1:1:007400690074006c0065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_1277236813_2648973148(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	completed bool
	/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		completed bool
		/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
	}])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_1277236813_3459698345(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[bool])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_1277236813_48201874(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[string])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_2400514730_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	title string
	/*gopurs:record-v1:1:007400690074006c0065*/
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_2648973148_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	completed bool
	/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_3459698345_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[bool]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_48201874_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[string]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}
