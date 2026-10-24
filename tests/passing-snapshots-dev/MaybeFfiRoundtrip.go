package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_makeMaybe gopurs_runtime.Value
var once_Main_makeMaybe sync.Once

func Get_Main_makeMaybe() gopurs_runtime.Value {
	once_Main_makeMaybe.Do(func() {
		cache_Main_makeMaybe = gopurs_runtime.Func2(func(present_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_makeMaybe((present_0_box.IntVal) != (0), value_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Main_makeMaybe
}

var cache_Main_checkSpan gopurs_runtime.Value
var once_Main_checkSpan sync.Once

func Get_Main_checkSpan() gopurs_runtime.Value {
	once_Main_checkSpan.Do(func() {
		cache_Main_checkSpan = gopurs_runtime.Func5(func(label_0_box gopurs_runtime.Value, predicate_1_box gopurs_runtime.Value, values_2_box gopurs_runtime.Value, expectedInit_3_box gopurs_runtime.Value, expectedRest_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkSpan(label_0_box.StrVal(), predicate_1_box, func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(values_2_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(expectedInit_3_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(expectedRest_4_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}())
		})
	})
	return cache_Main_checkSpan
}

var cache_Main_checkMaybe gopurs_runtime.Value
var once_Main_checkMaybe sync.Once

func Get_Main_checkMaybe() gopurs_runtime.Value {
	once_Main_checkMaybe.Do(func() {
		cache_Main_checkMaybe = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, value_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkMaybe(label_0_box.StrVal(), expected_1_box.StrVal(), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](value_2_box)))
		})
	})
	return cache_Main_checkMaybe
}

var cache_Main_checkFindIndex gopurs_runtime.Value
var once_Main_checkFindIndex sync.Once

func Get_Main_checkFindIndex() gopurs_runtime.Value {
	once_Main_checkFindIndex.Do(func() {
		cache_Main_checkFindIndex = gopurs_runtime.Func4(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, predicate_2_box gopurs_runtime.Value, values_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkFindIndex(label_0_box.StrVal(), expected_1_box.StrVal(), predicate_2_box, func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(values_3_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}())
		})
	})
	return cache_Main_checkFindIndex
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Main_checkMaybe("Nothing roundtrip", "Nothing", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Main_checkMaybe("Just zero roundtrip", "(Just 0)", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(int64(0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(Call_Main_checkMaybe("Just payload roundtrip", "(Just 7)", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(int64(7)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			savedCases_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				orig := struct {
					absent  bool
					payload int64
					present bool
					zero    int64
					/*gopurs:record-v1:4:0061006200730065006e0074,007000610079006c006f00610064,00700072006500730065006e0074,007a00650072006f*/
				}{false, int64(7), true, int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict4("absent", "payload", "present", "zero", gopurs_runtime.Bool(orig.absent), gopurs_runtime.Int(orig.payload), gopurs_runtime.Bool(orig.present), gopurs_runtime.Int(orig.zero))
			}()), gopurs_runtime.Value{})
			_ = savedCases_4_4
			cases_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), savedCases_4_4), gopurs_runtime.Value{})
			_ = cases_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(Call_Main_checkMaybe("constructed Nothing", "Nothing", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Main_makeMaybe((gopurs_runtime.RecordGet(cases_5_5, "absent").IntVal) != (0), gopurs_runtime.RecordGet(cases_5_5, "payload").IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(Call_Main_checkMaybe("constructed Just zero", "(Just 0)", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Main_makeMaybe((gopurs_runtime.RecordGet(cases_5_5, "present").IntVal) != (0), gopurs_runtime.RecordGet(cases_5_5, "zero").IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(Call_Main_checkMaybe("constructed Just payload", "(Just 7)", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Main_makeMaybe((gopurs_runtime.RecordGet(cases_5_5, "present").IntVal) != (0), gopurs_runtime.RecordGet(cases_5_5, "payload").IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(Call_Main_checkFindIndex("empty, false", "Nothing", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(false)
			}), []int64{}), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(Call_Main_checkFindIndex("empty, true", "Nothing", gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{}), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(Call_Main_checkFindIndex("singleton, false", "Nothing", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(false)
			}), []int64{int64(0)}), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply(Call_Main_checkFindIndex("singleton, true", "(Just 0)", gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{int64(0)}), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply(Call_Main_checkFindIndex("later match", "(Just 2)", gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_13.IntVal) == (int64(3)))
			}), []int64{int64(1), int64(2), int64(3)}), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			_dollar___unused_14_14 := gopurs_runtime.Apply(Call_Main_checkFindIndex("no match", "Nothing", gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_14.IntVal) == (int64(4)))
			}), []int64{int64(1), int64(2), int64(3)}), gopurs_runtime.Value{})
			_ = _dollar___unused_14_14
			_dollar___unused_15_15 := gopurs_runtime.Apply(Call_Main_checkSpan("span empty", gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{}, []int64{}, []int64{}), gopurs_runtime.Value{})
			_ = _dollar___unused_15_15
			_dollar___unused_16_16 := gopurs_runtime.Apply(Call_Main_checkSpan("span singleton all", gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{int64(1)}, []int64{int64(1)}, []int64{}), gopurs_runtime.Value{})
			_ = _dollar___unused_16_16
			_dollar___unused_17_17 := gopurs_runtime.Apply(Call_Main_checkSpan("span singleton none", gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(false)
			}), []int64{int64(1)}, []int64{}, []int64{int64(1)}), gopurs_runtime.Value{})
			_ = _dollar___unused_17_17
			_dollar___unused_18_18 := gopurs_runtime.Apply(Call_Main_checkSpan("span all", gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{int64(1), int64(2), int64(3)}, []int64{int64(1), int64(2), int64(3)}, []int64{}), gopurs_runtime.Value{})
			_ = _dollar___unused_18_18
			_dollar___unused_19_19 := gopurs_runtime.Apply(Call_Main_checkSpan("span prefix", gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_19.IntVal) < (int64(3)))
			}), []int64{int64(1), int64(2), int64(3)}, []int64{int64(1), int64(2)}, []int64{int64(3)}), gopurs_runtime.Value{})
			_ = _dollar___unused_19_19
			saved_20_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				arr := Call_Data_Array_go__range(int64(1), int64(10000))
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()), gopurs_runtime.Value{})
			_ = saved_20_20
			input_21_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), saved_20_20), gopurs_runtime.Value{})
			_ = input_21_21
			// TAST (Let): result_22_22 shape=App(Var) bindingType=(Record (Row [init: (Array Int), rest: (Array Int)] Empty))
			result_22_22 := func(record struct {
				go__init []gopurs_runtime.Value
				rest     []gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			}) struct {
				go__init []int64
				rest     []int64
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			} {
				return struct {
					go__init []int64
					rest     []int64
					/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.go__init).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.rest).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()}
			}(Call_Data_Array_span(gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(input_21_21.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).UnsafePtr))))
			_ = result_22_22
			_dollar___unused_23_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Int(int64(len(result_22_22.go__init))).IntVal, int64(10000)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_23_23
			_dollar___unused_24_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Int(int64(len(result_22_22.rest))).IntVal, int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_24_24
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2818770644_1386611502(Rebox_Main_1386611502_2818770644(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_4
}

func Call_Main_makeMaybe(present_0_loop bool, value_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var present_0 bool = present_0_loop
	_ = present_0
	var value_1 int64 = value_1_loop
	_ = value_1
	var __t0 *Constructor_Data_Maybe_Just[int64]
	{
		if present_0 {
			__t0 = Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(value_1), true}
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
		__t0 = Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Main_checkSpan(label_0_loop string, predicate_1_loop gopurs_runtime.Value, values_2_loop []int64, expectedInit_3_loop []int64, expectedRest_4_loop []int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var predicate_1 gopurs_runtime.Value = predicate_1_loop
	_ = predicate_1
	var values_2 []int64 = values_2_loop
	_ = values_2
	var expectedInit_3 []int64 = expectedInit_3_loop
	_ = expectedInit_3
	var expectedRest_4 []int64 = expectedRest_4_loop
	_ = expectedRest_4
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_5_0 shape=App(Var) bindingType=Any
		__local_var_5_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
			arr := values_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())
		_ = __local_var_5_0
		saved_6_1 := gopurs_runtime.Apply(__local_var_5_0, gopurs_runtime.Value{})
		_ = saved_6_1
		input_7_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), saved_6_1), gopurs_runtime.Value{})
		_ = input_7_2
		// TAST (Let): actual_8_3 shape=App(Var) bindingType=(Record (Row [init: (Array Int), rest: (Array Int)] Empty))
		actual_8_3 := func(record struct {
			go__init []gopurs_runtime.Value
			rest     []gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
		}) struct {
			go__init []int64
			rest     []int64
			/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
		} {
			return struct {
				go__init []int64
				rest     []int64
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.go__init).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.rest).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()}
		}(Call_Data_Array_span(predicate_1, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(input_7_2.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))))
		_ = actual_8_3
		_dollar___unused_9_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{actual_8_3.go__init, expectedInit_3}
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
		_ = _dollar___unused_9_4
		_dollar___unused_10_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{actual_8_3.rest, expectedRest_4}
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
		_ = _dollar___unused_10_5
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(label_0)), gopurs_runtime.Value{})
	})
}

func Call_Main_checkMaybe(label_0_loop string, expected_1_loop string, value_2_loop *Constructor_Data_Maybe_Just[int64]) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 string = expected_1_loop
	_ = expected_1
	var value_2 *Constructor_Data_Maybe_Just[int64] = value_2_loop
	_ = value_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(value_2))})
		_ = __local_var_3_0
		saved_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = saved_4_1
		actual_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), saved_4_1), gopurs_runtime.Value{})
		_ = actual_5_2
		_dollar___unused_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(Call_Data_Show_show(Rebox_Main_2818770644_1386611502(Rebox_Main_1386611502_2818770644(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_4())))), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](actual_5_2))))}).StrVal(), expected_1}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_6_3
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(label_0)), gopurs_runtime.Value{})
	})
}

func Call_Main_checkFindIndex(label_0_loop string, expected_1_loop string, predicate_2_loop gopurs_runtime.Value, values_3_loop []int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 string = expected_1_loop
	_ = expected_1
	var predicate_2 gopurs_runtime.Value = predicate_2_loop
	_ = predicate_2
	var values_3 []int64 = values_3_loop
	_ = values_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=Any
		__local_var_4_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
			arr := values_3
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())
		_ = __local_var_4_0
		saved_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
		_ = saved_5_1
		input_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), saved_5_1), gopurs_runtime.Value{})
		_ = input_6_2
		return gopurs_runtime.Apply(Call_Main_checkMaybe(label_0, expected_1, Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_Array_findIndex(predicate_2, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(input_6_2.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).UnsafePtr)))
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2818770644(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]])(unsafe.Pointer(in))
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
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

func Rebox_Main_2818770644_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
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

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]int64])(unsafe.Pointer(in))
}
