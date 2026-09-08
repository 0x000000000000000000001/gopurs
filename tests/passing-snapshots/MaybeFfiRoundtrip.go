package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878((&Constructor_Data_Eq_Eq[[]int64]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl())})))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502((&Constructor_Data_Show_Show[[]int64]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), Get_Data_Show_showIntImpl())})))}
	})
	return cache_Main_showArray
}

var cache_Main_showMaybe gopurs_runtime.Value
var once_Main_showMaybe sync.Once

func Get_Main_showMaybe() gopurs_runtime.Value {
	once_Main_showMaybe.Do(func() {
		cache_Main_showMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2818770644_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 string
			{
				var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
				if __t_tag_0 != nil {
					__t2 = (("(Just ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0).StrVal())) + (")")
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
				if __t_tag_1 == nil {
					__t2 = "Nothing"
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() string { panic("Failed pattern match") }()
			}
		end_branch_2:
			return gopurs_runtime.Str(__t2)
		})})))}
	})
	return cache_Main_showMaybe
}

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
			return Call_Main_checkMaybe(label_0_box.StrVal(), expected_1_box.StrVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](value_2_box))
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
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Main_checkMaybe("Just zero roundtrip", "(Just 0)", Rebox_Main_3094389156_1170268447((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(0))}))), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Main_checkMaybe("Just payload roundtrip", "(Just 7)", Rebox_Main_3094389156_1170268447((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(7))}))), gopurs_runtime.Value{})
			_ = __local_var_3_3
			savedCases_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					absent  bool
					payload int64
					present bool
					zero    int64
				}{false, int64(7), true, int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict4("absent", "payload", "present", "zero", gopurs_runtime.Bool(orig.absent), gopurs_runtime.Int(orig.payload), gopurs_runtime.Bool(orig.present), gopurs_runtime.Int(orig.zero))
			}()), gopurs_runtime.Value{})
			_ = savedCases_4_4
			cases_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), savedCases_4_4), gopurs_runtime.Value{})
			_ = cases_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Main_checkMaybe("constructed Nothing", "Nothing", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Main_makeMaybe((gopurs_runtime.RecordGet(cases_5_5, "absent").IntVal) != (0), gopurs_runtime.RecordGet(cases_5_5, "payload").IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_checkMaybe("constructed Just zero", "(Just 0)", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Main_makeMaybe((gopurs_runtime.RecordGet(cases_5_5, "present").IntVal) != (0), gopurs_runtime.RecordGet(cases_5_5, "zero").IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_checkMaybe("constructed Just payload", "(Just 7)", Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Main_makeMaybe((gopurs_runtime.RecordGet(cases_5_5, "present").IntVal) != (0), gopurs_runtime.RecordGet(cases_5_5, "payload").IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_checkFindIndex("empty, false", "Nothing", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(false)
			}), []int64{}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_checkFindIndex("empty, true", "Nothing", gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{}), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_checkFindIndex("singleton, false", "Nothing", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(false)
			}), []int64{int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_checkFindIndex("singleton, true", "(Just 0)", gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_checkFindIndex("later match", "(Just 2)", gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_13.IntVal) == (int64(3)))
			}), []int64{int64(1), int64(2), int64(3)}), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Main_checkFindIndex("no match", "Nothing", gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_14.IntVal) == (int64(4)))
			}), []int64{int64(1), int64(2), int64(3)}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Main_checkSpan("span empty", gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{}, []int64{}, []int64{}), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Main_checkSpan("span singleton all", gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{int64(1)}, []int64{int64(1)}, []int64{}), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Main_checkSpan("span singleton none", gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(false)
			}), []int64{int64(1)}, []int64{}, []int64{int64(1)}), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Main_checkSpan("span all", gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), []int64{int64(1), int64(2), int64(3)}, []int64{int64(1), int64(2), int64(3)}, []int64{}), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_checkSpan("span prefix", gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_19.IntVal) < (int64(3)))
			}), []int64{int64(1), int64(2), int64(3)}, []int64{int64(1), int64(2)}, []int64{int64(3)}), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(10000)))), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_20_20), gopurs_runtime.Value{})
			_ = __local_var_21_21
			// TAST (Let): result_22_22 shape=App(Var) bindingType=Any
			result_22_22 := gopurs_runtime.Apply2(Get_Data_Array_span(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__local_var_21_21.UnsafePtr)
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
			}())
			_ = result_22_22
			__local_var_23_23 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.RecordGet(result_22_22, "init")))).IntVal), gopurs_runtime.Int(int64(10000)))), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.RecordGet(result_22_22, "rest")))).IntVal), gopurs_runtime.Int(int64(0)))), gopurs_runtime.Value{})
			_ = __local_var_24_24
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_makeMaybe(present_0_loop bool, value_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var present_0 bool = present_0_loop
	_ = present_0
	var value_1 int64 = value_1_loop
	_ = value_1
	var __t0 gopurs_runtime.Value
	{
		if present_0 {
			__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(value_1)}))}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))}
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
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
		__local_var_5_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			arr := values_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())
		_ = __local_var_5_0
		__local_var_6_1 := gopurs_runtime.Apply(__local_var_5_0, gopurs_runtime.Value{})
		_ = __local_var_6_1
		__local_var_7_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_1), gopurs_runtime.Value{})
		_ = __local_var_7_2
		// TAST (Let): actual_8_3 shape=App(Var) bindingType=Any
		actual_8_3 := gopurs_runtime.Apply2(Get_Data_Array_span(), predicate_1, func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(__local_var_7_2.UnsafePtr)
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
		}())
		_ = actual_8_3
		__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(actual_8_3, "init").UnsafePtr)
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
		}(), func() gopurs_runtime.Value {
			arr := expectedInit_3
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())), gopurs_runtime.Value{})
		_ = __local_var_9_4
		__local_var_10_5 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(actual_8_3, "rest").UnsafePtr)
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
		}(), func() gopurs_runtime.Value {
			arr := expectedRest_4
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())), gopurs_runtime.Value{})
		_ = __local_var_10_5
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
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(value_2))})
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
		_ = __local_var_5_2
		__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]](Get_Main_showMaybe()).V0), __local_var_5_2).StrVal()), gopurs_runtime.Str(expected_1))), gopurs_runtime.Value{})
		_ = __local_var_6_3
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
		__local_var_4_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			arr := values_3
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())
		_ = __local_var_4_0
		__local_var_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
		_ = __local_var_5_1
		__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_1), gopurs_runtime.Value{})
		_ = __local_var_6_2
		return gopurs_runtime.Apply(Call_Main_checkMaybe(label_0, expected_1, Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}, predicate_2, __local_var_6_2)))), gopurs_runtime.Value{})
	})
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

func Rebox_Main_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
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

func Rebox_Main_2818770644_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
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
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
