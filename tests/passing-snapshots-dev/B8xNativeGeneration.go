package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Packed gopurs_runtime.Value
var once_Main_Packed sync.Once

func Get_Main_Packed() gopurs_runtime.Value {
	once_Main_Packed.Do(func() {
		cache_Main_Packed = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 214700806, UnsafePtr: unsafe.Pointer((&Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Packed
}

var cache_Main_unpack gopurs_runtime.Value
var once_Main_unpack sync.Once

func Get_Main_unpack() gopurs_runtime.Value {
	once_Main_unpack.Do(func() {
		cache_Main_unpack = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_unpack
}

var cache_Main_sumEleven gopurs_runtime.Value
var once_Main_sumEleven sync.Once

func Get_Main_sumEleven() gopurs_runtime.Value {
	once_Main_sumEleven.Do(func() {
		cache_Main_sumEleven = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(d_3_box gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(e_4_box gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(g_6_box gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(h_7_box gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(i_8_box gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(j_9_box gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(k_10_box gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Int(Call_Main_sumEleven(a_0_box.IntVal, b_1_box.IntVal, c_2_box.IntVal, d_3_box.IntVal, e_4_box.IntVal, f_5_box.IntVal, g_6_box.IntVal, h_7_box.IntVal, i_8_box.IntVal, j_9_box.IntVal, k_10_box.IntVal))
												})
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})
	return cache_Main_sumEleven
}

var cache_Main_runErased gopurs_runtime.Value
var once_Main_runErased sync.Once

func Get_Main_runErased() gopurs_runtime.Value {
	once_Main_runErased.Do(func() {
		cache_Main_runErased = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runErased(v_0_box)
		})
	})
	return cache_Main_runErased
}

var cache_Main_pack gopurs_runtime.Value
var once_Main_pack sync.Once

func Get_Main_pack() gopurs_runtime.Value {
	once_Main_pack.Do(func() {
		cache_Main_pack = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_pack
}

var cache_Main_markerInt gopurs_runtime.Value
var once_Main_markerInt sync.Once

func Get_Main_markerInt() gopurs_runtime.Value {
	once_Main_markerInt.Do(func() {
		cache_Main_markerInt = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_markerInt
}

var cache_Main_mapErased gopurs_runtime.Value
var once_Main_mapErased sync.Once

func Get_Main_mapErased() gopurs_runtime.Value {
	once_Main_mapErased.Do(func() {
		cache_Main_mapErased = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_mapErased(f_0_box, v_1_box)
		})
	})
	return cache_Main_mapErased
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(1)))
			_ = __local_var_0_0
			input_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = input_1_1
			first_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), input_1_1), gopurs_runtime.Value{})
			_ = first_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_sumEleven(first_2_2.IntVal, int64(2), int64(3), int64(4), int64(5), int64(6), int64(7), int64(8), int64(9), int64(10), int64(11)), int64(66)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			fnRef_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Get_Main_sumEleven()), gopurs_runtime.Value{})
			_ = fnRef_4_4
			fn_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), fnRef_4_4), gopurs_runtime.Value{})
			_ = fn_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(gopurs_runtime.Apply10(fn_5_5, gopurs_runtime.Int(first_2_2.IntVal), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6)), gopurs_runtime.Int(int64(7)), gopurs_runtime.Int(int64(8)), gopurs_runtime.Int(int64(9)), gopurs_runtime.Int(int64(10))), gopurs_runtime.Int(int64(11))).IntVal, int64(66)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			// TAST (Let): __local_var_7_8 shape=Other bindingType=(ADT ["Main","Erased"] [(TypeVar b$scope13)])
			__local_var_7_8 := (&Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntMul(v_7.IntVal, int64(2)))
			}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(v_7.IntVal, int64(1)))
			})), gopurs_runtime.Int(int64(21))})
			_ = __local_var_7_8
			_dollar___unused_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply((__local_var_7_8).V0, (__local_var_7_8).V1).IntVal, int64(44)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Packed[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}

type Constructor_Main_Marker[T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d00610072006b00650072
func init() {
	gopurs_runtime.StructGetters[185911708] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Marker[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Marker: " + key)
		}
	}
}

func Call_Main_sumEleven(a_0_loop int64, b_1_loop int64, c_2_loop int64, d_3_loop int64, e_4_loop int64, f_5_loop int64, g_6_loop int64, h_7_loop int64, i_8_loop int64, j_9_loop int64, k_10_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	var b_1 int64 = b_1_loop
	_ = b_1
	var c_2 int64 = c_2_loop
	_ = c_2
	var d_3 int64 = d_3_loop
	_ = d_3
	var e_4 int64 = e_4_loop
	_ = e_4
	var f_5 int64 = f_5_loop
	_ = f_5
	var g_6 int64 = g_6_loop
	_ = g_6
	var h_7 int64 = h_7_loop
	_ = h_7
	var i_8 int64 = i_8_loop
	_ = i_8
	var j_9 int64 = j_9_loop
	_ = j_9
	var k_10 int64 = k_10_loop
	_ = k_10
	return gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(a_0, b_1), c_2), d_3), e_4), f_5), g_6), h_7), i_8), j_9), k_10)
}

func Call_Main_runErased(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply((*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
}

func Call_Main_mapErased(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Value{Type: 9, IntVal: 214700806, UnsafePtr: unsafe.Pointer((&Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1}))}
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}
