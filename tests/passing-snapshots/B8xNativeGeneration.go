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

var cache_Main_Packed__592525336 gopurs_runtime.Value
var once_Main_Packed__592525336 sync.Once

func Get_Main_Packed__592525336() gopurs_runtime.Value {
	once_Main_Packed__592525336.Do(func() {
		cache_Main_Packed__592525336 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 214700806, UnsafePtr: unsafe.Pointer(Rebox_Main_3409422644_1106105172(Call_Main_Packed__592525336(__eta_norm_1_0_box, __eta_norm_0_1_box.IntVal)))}
		})
	})
	return cache_Main_Packed__592525336
}

var cache_Main_Marker_dollar_Dict gopurs_runtime.Value
var once_Main_Marker_dollar_Dict sync.Once

func Get_Main_Marker_dollar_Dict() gopurs_runtime.Value {
	once_Main_Marker_dollar_Dict.Do(func() {
		cache_Main_Marker_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Marker_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Marker_dollar_Dict
}

var cache_Main_Marker_dollar_Dict__1545153923 gopurs_runtime.Value
var once_Main_Marker_dollar_Dict__1545153923 sync.Once

func Get_Main_Marker_dollar_Dict__1545153923() gopurs_runtime.Value {
	once_Main_Marker_dollar_Dict__1545153923.Do(func() {
		cache_Main_Marker_dollar_Dict__1545153923 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Marker_dollar_Dict__1545153923(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Marker_dollar_Dict__1545153923
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

var cache_Main_runErased__2704168187 gopurs_runtime.Value
var once_Main_runErased__2704168187 sync.Once

func Get_Main_runErased__2704168187() gopurs_runtime.Value {
	once_Main_runErased__2704168187.Do(func() {
		cache_Main_runErased__2704168187 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runErased__2704168187(__eta_norm_0_0_box))
		})
	})
	return cache_Main_runErased__2704168187
}

var cache_Main_pack gopurs_runtime.Value
var once_Main_pack sync.Once

func Get_Main_pack() gopurs_runtime.Value {
	once_Main_pack.Do(func() {
		cache_Main_pack = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_pack
}

var cache_Main_pack__3800196234 gopurs_runtime.Value
var once_Main_pack__3800196234 sync.Once

func Get_Main_pack__3800196234() gopurs_runtime.Value {
	once_Main_pack__3800196234.Do(func() {
		cache_Main_pack__3800196234 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_pack__3800196234(Rebox_Main_1106105172_3409422644(gopurs_runtime.CoerceToStruct[Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
		})
	})
	return cache_Main_pack__3800196234
}

var cache_Main_markerInt gopurs_runtime.Value
var once_Main_markerInt sync.Once

func Get_Main_markerInt() gopurs_runtime.Value {
	once_Main_markerInt.Do(func() {
		cache_Main_markerInt = func() gopurs_runtime.Value {
			orig := struct {
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

var cache_Main_mapErased__4236014793 gopurs_runtime.Value
var once_Main_mapErased__4236014793 sync.Once

func Get_Main_mapErased__4236014793() gopurs_runtime.Value {
	once_Main_mapErased__4236014793.Do(func() {
		cache_Main_mapErased__4236014793 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_mapErased__4236014793(f_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Main_mapErased__4236014793
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(1)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_sumEleven(__local_var_2_2.IntVal, int64(2), int64(3), int64(4), int64(5), int64(6), int64(7), int64(8), int64(9), int64(10), int64(11)), int64(66)}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_4_3 shape=App(Var) bindingType=Any
					__local_var_4_3 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Main_sumEleven())
					_ = __local_var_4_3
					fnRef_5_4 := gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Value{})
					_ = fnRef_5_4
					__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), fnRef_5_4), gopurs_runtime.Value{})
					_ = __local_var_6_5
					return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{gopurs_runtime.Apply(gopurs_runtime.Apply10(__local_var_6_5, gopurs_runtime.Int(__local_var_2_2.IntVal), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6)), gopurs_runtime.Int(int64(7)), gopurs_runtime.Int(int64(8)), gopurs_runtime.Int(int64(9)), gopurs_runtime.Int(int64(10))), gopurs_runtime.Int(int64(11))).IntVal, int64(66)}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): __local_var_8_6 shape=Other bindingType=(ADT ["Main","Erased"] [Int])
						__local_var_8_6 := (&Constructor_Main_Packed[int64, int64]{1, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v_8.IntVal) + (int64(1)))
						}), int64(21)})
						_ = __local_var_8_6
						// TAST (Let): __local_var_9_7 shape=Other bindingType=(ADT ["Main","Erased"] [(TypeVar b$scope13)])
						__local_var_9_7 := (&Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v_9.IntVal) * (int64(2)))
						}), (__local_var_8_6).V0), gopurs_runtime.Int((__local_var_8_6).V1)})
						_ = __local_var_9_7
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{gopurs_runtime.Apply((__local_var_9_7).V0, (__local_var_9_7).V1).IntVal, int64(44)}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
						}))
					})), gopurs_runtime.Value{})
				})
			})), gopurs_runtime.Value{})
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

func Call_Main_Packed__592525336(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop int64) *Constructor_Main_Packed[int64, int64] {
Packed__592525336:
	for {
		if false {
			continue Packed__592525336
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Packed[int64, int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_Marker_dollar_Dict(x_0_loop struct {
}) gopurs_runtime.Value {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_Marker_dollar_Dict__1545153923(x_0_loop struct {
}) gopurs_runtime.Value {
Marker_dollar_Dict__1545153923:
	for {
		if false {
			continue Marker_dollar_Dict__1545153923
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
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
	return ((((((((((a_0) + (b_1)) + (c_2)) + (d_3)) + (e_4)) + (f_5)) + (g_6)) + (h_7)) + (i_8)) + (j_9)) + (k_10)
}

func Call_Main_runErased(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply((*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
}

func Call_Main_runErased__2704168187(__eta_norm_0_0_loop gopurs_runtime.Value) int64 {
runErased__2704168187:
	for {
		if false {
			continue runErased__2704168187
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply((*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(__eta_norm_0_0.UnsafePtr).V0, (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(__eta_norm_0_0.UnsafePtr).V1).IntVal
	}
}

func Call_Main_pack__3800196234(__eta_norm_0_0_loop *Constructor_Main_Packed[int64, int64]) gopurs_runtime.Value {
pack__3800196234:
	for {
		if false {
			continue pack__3800196234
		}
		var __eta_norm_0_0 *Constructor_Main_Packed[int64, int64] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 214700806, UnsafePtr: unsafe.Pointer(Rebox_Main_3409422644_1106105172(__eta_norm_0_0))}
	}
}

func Call_Main_mapErased(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Value{Type: 9, IntVal: 214700806, UnsafePtr: unsafe.Pointer((&Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1}))}
}

func Call_Main_mapErased__4236014793(f_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapErased__4236014793:
	for {
		if false {
			continue mapErased__4236014793
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 214700806, UnsafePtr: unsafe.Pointer((&Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(__eta_norm_0_1.UnsafePtr).V0), (*Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value])(__eta_norm_0_1.UnsafePtr).V1}))}
	}
}

func Rebox_Main_1106105172_3409422644(in *Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_Packed[int64, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Packed[int64, int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_3409422644_1106105172(in *Constructor_Main_Packed[int64, int64]) *Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Packed[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}
