package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_cIsSymbol gopurs_runtime.Value
var once_Main_cIsSymbol sync.Once

func Get_Main_cIsSymbol() gopurs_runtime.Value {
	once_Main_cIsSymbol.Do(func() {
		cache_Main_cIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("c")
		}))
	})
	return cache_Main_cIsSymbol
}

var cache_Main_bIsSymbol gopurs_runtime.Value
var once_Main_bIsSymbol sync.Once

func Get_Main_bIsSymbol() gopurs_runtime.Value {
	once_Main_bIsSymbol.Do(func() {
		cache_Main_bIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("b")
		}))
	})
	return cache_Main_bIsSymbol
}

var cache_Main_aIsSymbol gopurs_runtime.Value
var once_Main_aIsSymbol sync.Once

func Get_Main_aIsSymbol() gopurs_runtime.Value {
	once_Main_aIsSymbol.Do(func() {
		cache_Main_aIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		}))
	})
	return cache_Main_aIsSymbol
}

var cache_Main_headIsSymbol gopurs_runtime.Value
var once_Main_headIsSymbol sync.Once

func Get_Main_headIsSymbol() gopurs_runtime.Value {
	once_Main_headIsSymbol.Do(func() {
		cache_Main_headIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("head")
		}))
	})
	return cache_Main_headIsSymbol
}

var cache_Main_outerIsSymbol gopurs_runtime.Value
var once_Main_outerIsSymbol sync.Once

func Get_Main_outerIsSymbol() gopurs_runtime.Value {
	once_Main_outerIsSymbol.Do(func() {
		cache_Main_outerIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("outer")
		}))
	})
	return cache_Main_outerIsSymbol
}

var cache_Main_requiredIsSymbol gopurs_runtime.Value
var once_Main_requiredIsSymbol sync.Once

func Get_Main_requiredIsSymbol() gopurs_runtime.Value {
	once_Main_requiredIsSymbol.Do(func() {
		cache_Main_requiredIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("required")
		}))
	})
	return cache_Main_requiredIsSymbol
}

var cache_Main_tailIsSymbol gopurs_runtime.Value
var once_Main_tailIsSymbol sync.Once

func Get_Main_tailIsSymbol() gopurs_runtime.Value {
	once_Main_tailIsSymbol.Do(func() {
		cache_Main_tailIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("tail")
		}))
	})
	return cache_Main_tailIsSymbol
}

var cache_Main_Probe gopurs_runtime.Value
var once_Main_Probe sync.Once

func Get_Main_Probe() gopurs_runtime.Value {
	once_Main_Probe.Do(func() {
		cache_Main_Probe = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_Probe(x_0_box.IntVal))
		})
	})
	return cache_Main_Probe
}

var cache_Main_showProbe gopurs_runtime.Value
var once_Main_showProbe sync.Once

func Get_Main_showProbe() gopurs_runtime.Value {
	once_Main_showProbe.Do(func() {
		cache_Main_showProbe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}
	})
	return cache_Main_showProbe
}

var cache_Main_sharedTail gopurs_runtime.Value
var once_Main_sharedTail sync.Once

func Get_Main_sharedTail() gopurs_runtime.Value {
	once_Main_sharedTail.Do(func() {
		cache_Main_sharedTail = func() gopurs_runtime.Value {
			orig := struct {
				tail int64
				/*gopurs:record-v1:1:007400610069006c*/
			}{int64(41)}
			_ = orig
			return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
		}()
	})
	return cache_Main_sharedTail
}

var cache_Main_object gopurs_runtime.Value
var once_Main_object sync.Once

func Get_Main_object() gopurs_runtime.Value {
	once_Main_object.Do(func() {
		cache_Main_object = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Argonaut_Core_fromObject(), gopurs_runtime.Apply(Get_Foreign_Object_fromFoldable(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}))
	})
	return cache_Main_object
}

var cache_Main_eqProbe gopurs_runtime.Value
var once_Main_eqProbe sync.Once

func Get_Main_eqProbe() gopurs_runtime.Value {
	once_Main_eqProbe.Do(func() {
		cache_Main_eqProbe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878((&Constructor_Data_Eq_Eq[int64]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((x_0.IntVal) == (y_1.IntVal))
		})})))}
	})
	return cache_Main_eqProbe
}

var cache_Main_calls gopurs_runtime.Value
var once_Main_calls sync.Once

func Get_Main_calls() gopurs_runtime.Value {
	once_Main_calls.Do(func() {
		cache_Main_calls = gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
			arr := []int64{}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()))
	})
	return cache_Main_calls
}

var cache_Main_observe gopurs_runtime.Value
var once_Main_observe sync.Once

func Get_Main_observe() gopurs_runtime.Value {
	once_Main_observe.Do(func() {
		cache_Main_observe = gopurs_runtime.Func(func(value_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_observe(value_0_box.IntVal))
		})
	})
	return cache_Main_observe
}

var cache_Main_decodeProbe gopurs_runtime.Value
var once_Main_decodeProbe sync.Once

func Get_Main_decodeProbe() gopurs_runtime.Value {
	once_Main_decodeProbe.Do(func() {
		cache_Main_decodeProbe = gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813((&Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64]{1, gopurs_runtime.Func(func(json_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), Int])
			__local_var_1_0 := gopurs_runtime.Apply(Data_Argonaut_Decode_Internal_Record_TypedInt(func() gopurs_runtime.Value {
				orig := func() struct {
					atIndex      gopurs_runtime.Value
					atKey        gopurs_runtime.Value
					isRight      gopurs_runtime.Value
					just         gopurs_runtime.Value
					left         gopurs_runtime.Value
					leftOf       gopurs_runtime.Value
					missingValue gopurs_runtime.Value
					named        gopurs_runtime.Value
					nothing      *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					right        gopurs_runtime.Value
					rightValue   gopurs_runtime.Value
					typeMismatch gopurs_runtime.Value
					/*gopurs:record-v1:12:006100740049006e006400650078,00610074004b00650079,0069007300520069006700680074,006a007500730074,006c006500660074,006c006500660074004f0066,006d0069007300730069006e006700560061006c00750065,006e0061006d00650064,006e006f007400680069006e0067,00720069006700680074,0072006900670068007400560061006c00750065,0074007900700065004d00690073006d0061007400630068*/
				} {
					orig := Get_Data_Argonaut_Decode_Class_recordErrorSupport()
					_ = orig
					clone := struct {
						atIndex      gopurs_runtime.Value
						atKey        gopurs_runtime.Value
						isRight      gopurs_runtime.Value
						just         gopurs_runtime.Value
						left         gopurs_runtime.Value
						leftOf       gopurs_runtime.Value
						missingValue gopurs_runtime.Value
						named        gopurs_runtime.Value
						nothing      *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
						right        gopurs_runtime.Value
						rightValue   gopurs_runtime.Value
						typeMismatch gopurs_runtime.Value
						/*gopurs:record-v1:12:006100740049006e006400650078,00610074004b00650079,0069007300520069006700680074,006a007500730074,006c006500660074,006c006500660074004f0066,006d0069007300730069006e006700560061006c00750065,006e0061006d00650064,006e006f007400680069006e0067,00720069006700680074,0072006900670068007400560061006c00750065,0074007900700065004d00690073006d0061007400630068*/
					}{}
					clone.atIndex = gopurs_runtime.RecordGet(orig, "atIndex")
					clone.atKey = gopurs_runtime.RecordGet(orig, "atKey")
					clone.isRight = gopurs_runtime.RecordGet(orig, "isRight")
					clone.just = gopurs_runtime.RecordGet(orig, "just")
					clone.left = gopurs_runtime.RecordGet(orig, "left")
					clone.leftOf = gopurs_runtime.RecordGet(orig, "leftOf")
					clone.missingValue = gopurs_runtime.RecordGet(orig, "missingValue")
					clone.named = gopurs_runtime.RecordGet(orig, "named")
					clone.nothing = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "nothing"))
					clone.right = gopurs_runtime.RecordGet(orig, "right")
					clone.rightValue = gopurs_runtime.RecordGet(orig, "rightValue")
					clone.typeMismatch = gopurs_runtime.RecordGet(orig, "typeMismatch")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"atIndex", "atKey", "isRight", "just", "left", "leftOf", "missingValue", "named", "nothing", "right", "rightValue", "typeMismatch"}, []gopurs_runtime.Value{orig.atIndex, orig.atKey, orig.isRight, orig.just, orig.left, orig.leftOf, orig.missingValue, orig.named, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.nothing)}, orig.right, orig.rightValue, orig.typeMismatch})
			}(), Get_Data_Argonaut_Decode_Decoders_decodeInt()), json_0)
			_ = __local_var_1_0
			var __t5 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 3711209382 {
					// TAST (Let): __local_var_2_1 shape=Other bindingType=(TypeVar e$scope161)
					__local_var_2_1 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
					_ = __local_var_2_1
					__t5 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{__local_var_2_1, gopurs_runtime.Value{}, false}
					goto end_branch_5
				} else {

				}
			}
			{
				if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 2465973597 {
					// TAST (Let): __local_var_2_2 shape=Other bindingType=(TypeVar b$scope119)
					__local_var_2_2 := (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
					_ = __local_var_2_2
					// TAST (Let): observed_5_3 shape=App(Var) bindingType=Int
					observed_5_3 := Call_Main_observe(__local_var_2_2.IntVal)
					_ = observed_5_3
					var __t4 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}
					{
						if (observed_5_3) < (int64(0)) {
							__t4 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "negative probe"}))}, gopurs_runtime.Value{}, false}
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, gopurs_runtime.Int(observed_5_3), true}
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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
		end_branch_5:
			return func() gopurs_runtime.Value {
				_v := __t5
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})})))}
	})
	return cache_Main_decodeProbe
}

var cache_Main_decodeTracked gopurs_runtime.Value
var once_Main_decodeTracked sync.Once

func Get_Main_decodeTracked() gopurs_runtime.Value {
	once_Main_decodeTracked.Do(func() {
		cache_Main_decodeTracked = Call_Data_Argonaut_Decode_Class_decodeJson(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main_decodeProbe()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main_decodeProbe()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main_decodeProbe()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_cIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_bIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_aIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))
	})
	return cache_Main_decodeTracked
}

var cache_Main_tailResult gopurs_runtime.Value
var once_Main_tailResult sync.Once

func Get_Main_tailResult() gopurs_runtime.Value {
	once_Main_tailResult.Do(func() {
		cache_Main_tailResult = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_tailResult(v_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Main_tailResult
}

var cache_Main_customTailDecoder gopurs_runtime.Value
var once_Main_customTailDecoder sync.Once

func Get_Main_customTailDecoder() gopurs_runtime.Value {
	once_Main_customTailDecoder.Do(func() {
		cache_Main_customTailDecoder = gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(value_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_tailResult(value_0)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})}))}
	})
	return cache_Main_customTailDecoder
}

var cache_Main_decodeCustom gopurs_runtime.Value
var once_Main_decodeCustom sync.Once

func Get_Main_decodeCustom() gopurs_runtime.Value {
	once_Main_decodeCustom.Do(func() {
		cache_Main_decodeCustom = gopurs_runtime.Func(func(value_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_decodeCustom(value_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Main_decodeCustom
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), Get_Main_calls())
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_1()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_2())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(Get_Main_decodeTracked(), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("c"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(3.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("a"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(1.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("b"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(2.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())), func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3))), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			actual_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = actual_3_4
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_3()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_4())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(actual_3_4.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(1), int64(2), int64(3)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = _dollar___unused_4_5
			_dollar___unused_5_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_5()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_6())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(Get_Main_decodeTracked(), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("a"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(1.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("b"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(-2.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("c"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(-3.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())), func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "b", gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "negative probe"}))}}))}, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_5_6
			actual_6_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = actual_6_8
			_dollar___unused_6_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_7()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_8())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(actual_6_8.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(1), int64(-2)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_6_7
			_dollar___unused_7_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = _dollar___unused_7_9
			_dollar___unused_8_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_9()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_10())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(Get_Main_decodeTracked(), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("a"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(1.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("c"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(3.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())), func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "b", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_8_10
			actual_9_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = actual_9_12
			_dollar___unused_9_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_11()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_12())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(actual_9_12.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(1)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_9_11
			_dollar___unused_10_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_13()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_14())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeJson(Rebox_Main_2334244318_1277236813(Rebox_Main_1277236813_2334244318(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main___cached_dict_15())))), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("a"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(7.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())), func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(int64(7)), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_10_13
			_dollar___unused_11_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_16()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_17())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeJson(Rebox_Main_2334244318_1277236813(Rebox_Main_1277236813_2334244318(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main___cached_dict_18())))), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("a"), Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(7.0))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("b"), Get_Data_Argonaut_Core_jsonNull()}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())), func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(int64(7)), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_11_14
			_dollar___unused_12_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_19()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_20())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeJson(Rebox_Main_252997862_1277236813(Rebox_Main_1277236813_252997862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main___cached_dict_21())))), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("outer"), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
							arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}())}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())), func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "outer", gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "required", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}}))}, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_12_15
			_dollar___unused_13_16 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = _dollar___unused_13_16
			first_14_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				_v := Call_Main_decodeCustom(Call_Foreign_Object_singleton("head", Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(1.0))))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()), gopurs_runtime.Value{})
			_ = first_14_17
			second_15_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				_v := Call_Main_decodeCustom(Call_Foreign_Object_singleton("head", Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(2.0))))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()), gopurs_runtime.Value{})
			_ = second_15_18
			actual_16_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), first_14_17), gopurs_runtime.Value{})
			_ = actual_16_20
			_dollar___unused_16_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_22()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_23())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{actual_16_20, func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(41))), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_16_19
			actual_17_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), second_15_18), gopurs_runtime.Value{})
			_ = actual_17_22
			_dollar___unused_17_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_24()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_25())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{actual_17_22, func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(41))), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_17_21
			_dollar___unused_18_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_381738941_3790796878(Rebox_Main_3790796878_381738941(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_26()))), Rebox_Main_2815936861_1386611502(Rebox_Main_1386611502_2815936861(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_27())))), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						tail int64
						/*gopurs:record-v1:1:007400610069006c*/
					}
					expected struct {
						tail int64
						/*gopurs:record-v1:1:007400610069006c*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() struct {
					tail int64
					/*gopurs:record-v1:1:007400610069006c*/
				} {
					orig := Get_Main_sharedTail()
					_ = orig
					clone := struct {
						tail int64
						/*gopurs:record-v1:1:007400610069006c*/
					}{}
					clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
					return clone
				}(), struct {
					tail int64
					/*gopurs:record-v1:1:007400610069006c*/
				}{int64(41)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_18_23
			_dollar___unused_19_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeHas(), gopurs_runtime.Str("head"), func() gopurs_runtime.Value {
					orig := func() struct {
						tail int64
						/*gopurs:record-v1:1:007400610069006c*/
					} {
						orig := Get_Main_sharedTail()
						_ = orig
						clone := struct {
							tail int64
							/*gopurs:record-v1:1:007400610069006c*/
						}{}
						clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
				}()).IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_19_24
			actual_20_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = actual_20_26
			_dollar___unused_20_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_28()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_29())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(actual_20_26.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(100), int64(100)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_20_25
			_dollar___unused_21_27 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = _dollar___unused_21_27
			_dollar___unused_22_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_30()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_31())), func() gopurs_runtime.Value {
				orig := struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() gopurs_runtime.Value {
					_v := Call_Main_decodeCustom(Get_Foreign_Object_empty())
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
							orig := _v.V1
							_ = orig
							return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
						}()})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}(), func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "head", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", orig.actual, orig.expected)
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_22_28
			actual_23_30 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls()), gopurs_runtime.Value{})
			_ = actual_23_30
			_dollar___unused_23_29 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_32()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_33())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(actual_23_30.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_23_29
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_customTailDecoder()))}, Get_Main_headIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_bIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_aIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))})))))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_aIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_bIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_bIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_aIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))})))))}
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_aIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_bIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_bIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_aIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))})))))}
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_aIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_bIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_11
}

var cache_Main___cached_dict_12 gopurs_runtime.Value
var once_Main___cached_dict_12 sync.Once

func Get_Main___cached_dict_12() gopurs_runtime.Value {
	once_Main___cached_dict_12.Do(func() {
		cache_Main___cached_dict_12 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_12
}

var cache_Main___cached_dict_13 gopurs_runtime.Value
var once_Main___cached_dict_13 sync.Once

func Get_Main___cached_dict_13() gopurs_runtime.Value {
	once_Main___cached_dict_13.Do(func() {
		cache_Main___cached_dict_13 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_bIsSymbol(), Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_aIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_13
}

var cache_Main___cached_dict_14 gopurs_runtime.Value
var once_Main___cached_dict_14 sync.Once

func Get_Main___cached_dict_14() gopurs_runtime.Value {
	once_Main___cached_dict_14.Do(func() {
		cache_Main___cached_dict_14 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_aIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_bIsSymbol(), Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_14
}

var cache_Main___cached_dict_15 gopurs_runtime.Value
var once_Main___cached_dict_15 sync.Once

func Get_Main___cached_dict_15() gopurs_runtime.Value {
	once_Main___cached_dict_15.Do(func() {
		cache_Main___cached_dict_15 = gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_2334244318_1277236813(Rebox_Main_1277236813_2334244318(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_bIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_aIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main___cached_dict_15
}

var cache_Main___cached_dict_16 gopurs_runtime.Value
var once_Main___cached_dict_16 sync.Once

func Get_Main___cached_dict_16() gopurs_runtime.Value {
	once_Main___cached_dict_16.Do(func() {
		cache_Main___cached_dict_16 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_bIsSymbol(), Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_aIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_16
}

var cache_Main___cached_dict_17 gopurs_runtime.Value
var once_Main___cached_dict_17 sync.Once

func Get_Main___cached_dict_17() gopurs_runtime.Value {
	once_Main___cached_dict_17.Do(func() {
		cache_Main___cached_dict_17 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_aIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_bIsSymbol(), Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_17
}

var cache_Main___cached_dict_18 gopurs_runtime.Value
var once_Main___cached_dict_18 sync.Once

func Get_Main___cached_dict_18() gopurs_runtime.Value {
	once_Main___cached_dict_18.Do(func() {
		cache_Main___cached_dict_18 = gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_2334244318_1277236813(Rebox_Main_1277236813_2334244318(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_bIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_aIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main___cached_dict_18
}

var cache_Main___cached_dict_19 gopurs_runtime.Value
var once_Main___cached_dict_19 sync.Once

func Get_Main___cached_dict_19() gopurs_runtime.Value {
	once_Main___cached_dict_19.Do(func() {
		cache_Main___cached_dict_19 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_outerIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_requiredIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))))}
	})
	return cache_Main___cached_dict_19
}

var cache_Main___cached_dict_20 gopurs_runtime.Value
var once_Main___cached_dict_20 sync.Once

func Get_Main___cached_dict_20() gopurs_runtime.Value {
	once_Main___cached_dict_20.Do(func() {
		cache_Main___cached_dict_20 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_outerIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_requiredIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))))}
	})
	return cache_Main___cached_dict_20
}

var cache_Main___cached_dict_21 gopurs_runtime.Value
var once_Main___cached_dict_21 sync.Once

func Get_Main___cached_dict_21() gopurs_runtime.Value {
	once_Main___cached_dict_21.Do(func() {
		cache_Main___cached_dict_21 = gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_252997862_1277236813(Rebox_Main_1277236813_252997862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_requiredIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{}))), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_outerIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main___cached_dict_21
}

var cache_Main___cached_dict_22 gopurs_runtime.Value
var once_Main___cached_dict_22 sync.Once

func Get_Main___cached_dict_22() gopurs_runtime.Value {
	once_Main___cached_dict_22.Do(func() {
		cache_Main___cached_dict_22 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_headIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_22
}

var cache_Main___cached_dict_23 gopurs_runtime.Value
var once_Main___cached_dict_23 sync.Once

func Get_Main___cached_dict_23() gopurs_runtime.Value {
	once_Main___cached_dict_23.Do(func() {
		cache_Main___cached_dict_23 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_headIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_23
}

var cache_Main___cached_dict_24 gopurs_runtime.Value
var once_Main___cached_dict_24 sync.Once

func Get_Main___cached_dict_24() gopurs_runtime.Value {
	once_Main___cached_dict_24.Do(func() {
		cache_Main___cached_dict_24 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_headIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_24
}

var cache_Main___cached_dict_25 gopurs_runtime.Value
var once_Main___cached_dict_25 sync.Once

func Get_Main___cached_dict_25() gopurs_runtime.Value {
	once_Main___cached_dict_25.Do(func() {
		cache_Main___cached_dict_25 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_headIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_25
}

var cache_Main___cached_dict_26 gopurs_runtime.Value
var once_Main___cached_dict_26 sync.Once

func Get_Main___cached_dict_26() gopurs_runtime.Value {
	once_Main___cached_dict_26.Do(func() {
		cache_Main___cached_dict_26 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_381738941_3790796878(Rebox_Main_3790796878_381738941(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main___cached_dict_26
}

var cache_Main___cached_dict_27 gopurs_runtime.Value
var once_Main___cached_dict_27 sync.Once

func Get_Main___cached_dict_27() gopurs_runtime.Value {
	once_Main___cached_dict_27.Do(func() {
		cache_Main___cached_dict_27 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2815936861_1386611502(Rebox_Main_1386611502_2815936861(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main___cached_dict_27
}

var cache_Main___cached_dict_28 gopurs_runtime.Value
var once_Main___cached_dict_28 sync.Once

func Get_Main___cached_dict_28() gopurs_runtime.Value {
	once_Main___cached_dict_28.Do(func() {
		cache_Main___cached_dict_28 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_28
}

var cache_Main___cached_dict_29 gopurs_runtime.Value
var once_Main___cached_dict_29 sync.Once

func Get_Main___cached_dict_29() gopurs_runtime.Value {
	once_Main___cached_dict_29.Do(func() {
		cache_Main___cached_dict_29 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_29
}

var cache_Main___cached_dict_30 gopurs_runtime.Value
var once_Main___cached_dict_30 sync.Once

func Get_Main___cached_dict_30() gopurs_runtime.Value {
	once_Main___cached_dict_30.Do(func() {
		cache_Main___cached_dict_30 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_headIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_30
}

var cache_Main___cached_dict_31 gopurs_runtime.Value
var once_Main___cached_dict_31 sync.Once

func Get_Main___cached_dict_31() gopurs_runtime.Value {
	once_Main___cached_dict_31.Do(func() {
		cache_Main___cached_dict_31 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_headIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_31
}

var cache_Main___cached_dict_32 gopurs_runtime.Value
var once_Main___cached_dict_32 sync.Once

func Get_Main___cached_dict_32() gopurs_runtime.Value {
	once_Main___cached_dict_32.Do(func() {
		cache_Main___cached_dict_32 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_32
}

var cache_Main___cached_dict_33 gopurs_runtime.Value
var once_Main___cached_dict_33 sync.Once

func Get_Main___cached_dict_33() gopurs_runtime.Value {
	once_Main___cached_dict_33.Do(func() {
		cache_Main___cached_dict_33 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_33
}

func Call_Main_Probe(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_observe(value_0_loop int64) int64 {
	var value_0 int64 = value_0_loop
	_ = value_0
	return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), seen_1, func() gopurs_runtime.Value {
				arr := []int64{value_0}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())
		}), Get_Main_calls())
		_ = __local_var_1_0
		_dollar___unused_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = _dollar___unused_2_1
		return gopurs_runtime.Int(value_0)
	})).IntVal
}

func Call_Main_tailResult(v_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		tail int64
		/*gopurs:record-v1:1:007400610069006c*/
	}
	V2 bool
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 struct {
			tail int64
			/*gopurs:record-v1:1:007400610069006c*/
		}
		V2 bool
	}
	{
		if (Call_Main_observe(int64(100))) == (int64(100)) {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 struct {
					tail int64
					/*gopurs:record-v1:1:007400610069006c*/
				}
				V2 bool
			}{gopurs_runtime.Value{}, func() struct {
				tail int64
				/*gopurs:record-v1:1:007400610069006c*/
			} {
				orig := Get_Main_sharedTail()
				_ = orig
				clone := struct {
					tail int64
					/*gopurs:record-v1:1:007400610069006c*/
				}{}
				clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
				return clone
			}(), true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = struct {
			V0 gopurs_runtime.Value
			V1 struct {
				tail int64
				/*gopurs:record-v1:1:007400610069006c*/
			}
			V2 bool
		}{gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}, struct {
			tail int64
			/*gopurs:record-v1:1:007400610069006c*/
		}{}, false}
	}
end_branch_0:
	return __t0
}

func Call_Main_decodeCustom(value_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		head int64
		tail int64
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V2 bool
} {
	var value_0 gopurs_runtime.Value = value_0_loop
	_ = value_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 struct {
			head int64
			tail int64
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}
		V2 bool
	} {
		_v := gopurs_runtime.Apply2(Call_Data_Argonaut_Decode_Class_gDecodeJson(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main___cached_dict_0())), value_0, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
		if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 struct {
					head int64
					tail int64
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: func() struct {
				head int64
				tail int64
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			} {
				orig := (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					head int64
					tail int64
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}
				clone.head = gopurs_runtime.RecordGet(orig, "head").IntVal
				clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
				return clone
			}(), V2: true}
		}
		return struct {
			V0 gopurs_runtime.Value
			V1 struct {
				head int64
				tail int64
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: struct {
			head int64
			tail int64
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, V2: false}
	}()
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1277236813_1405234230(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1277236813_2334244318(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	a int64
	b *Constructor_Data_Maybe_Just[int64]
	/*gopurs:record-v1:2:0061,0062*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		a int64
		b *Constructor_Data_Maybe_Just[int64]
		/*gopurs:record-v1:2:0061,0062*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_1277236813_252997862(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	outer struct {
		required int64
		/*gopurs:record-v1:1:00720065007100750069007200650064*/
	}
	/*gopurs:record-v1:1:006f0075007400650072*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		outer struct {
			required int64
			/*gopurs:record-v1:1:00720065007100750069007200650064*/
		}
		/*gopurs:record-v1:1:006f0075007400650072*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_138441832_3103076375(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1
	return out
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2815936861(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	tail int64
	/*gopurs:record-v1:1:007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[struct {
		tail int64
		/*gopurs:record-v1:1:007400610069006c*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_1405234230_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2334244318_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	a int64
	b *Constructor_Data_Maybe_Just[int64]
	/*gopurs:record-v1:2:0061,0062*/
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_252997862_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	outer struct {
		required int64
		/*gopurs:record-v1:1:00720065007100750069007200650064*/
	}
	/*gopurs:record-v1:1:006f0075007400650072*/
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2815936861_1386611502(in *Constructor_Data_Show_Show[struct {
	tail int64
	/*gopurs:record-v1:1:007400610069006c*/
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3103076375_138441832(in *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_381738941(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	tail int64
	/*gopurs:record-v1:1:007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[struct {
		tail int64
		/*gopurs:record-v1:1:007400610069006c*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_381738941_3790796878(in *Constructor_Data_Eq_Eq[struct {
	tail int64
	/*gopurs:record-v1:1:007400610069006c*/
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
