package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3467499472_1386611502(Rebox_Main_1386611502_3467499472(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_logShow gopurs_runtime.Value
var once_Main_logShow sync.Once

func Get_Main_logShow() gopurs_runtime.Value {
	once_Main_logShow.Do(func() {
		cache_Main_logShow = gopurs_runtime.Apply(Get_Effect_Console_logShow(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))})
	})
	return cache_Main_logShow
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_g(__local_var_0_box, __local_var_1_box)
		})
	})
	return cache_Main_g
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(__local_var_0_box, __local_var_1_box)
		})
	})
	return cache_Main_f
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Data_Function_Uncurried_runFn0(), gopurs_runtime.Apply(Get_Data_Function_Uncurried_mkFn0(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(0.0)).StrVal()))
		}))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(gopurs_runtime.Func(func(__local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), __local_var_1).StrVal()))
			}), gopurs_runtime.Float(0.0)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
					arr := []float64{0.0, 0.0}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Float(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
						arr := []float64{0.0, 0.0, 0.0}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Float(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
							arr := []float64{0.0, 0.0, 0.0, 0.0}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Float(v)
							}
							return gopurs_runtime.Array(boxed)
						}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
								arr := []float64{0.0, 0.0, 0.0, 0.0, 0.0}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Float(v)
								}
								return gopurs_runtime.Array(boxed)
							}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
									arr := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Float(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
										arr := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Float(v)
										}
										return gopurs_runtime.Array(boxed)
									}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
											arr := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Float(v)
											}
											return gopurs_runtime.Array(boxed)
										}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
												arr := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Float(v)
												}
												return gopurs_runtime.Array(boxed)
											}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), func() gopurs_runtime.Value {
													arr := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
													boxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														boxed[i] = gopurs_runtime.Float(v)
													}
													return gopurs_runtime.Array(boxed)
												}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))), gopurs_runtime.Float(gopurs_runtime.UncurriedApp2(Get_Main_g(), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0)).FloatVal())), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
													}))
												}))
											}))
										}))
									}))
								}))
							}))
						}))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main_g(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	var __t3 float64
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_0, gopurs_runtime.Float(0.0))
		_ = __t_tag_0
		var __t_or_2 bool = true
		if !((uint32(__t_tag_0.IntVal) == 380165415) != (true)) {

			var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_1, gopurs_runtime.Float(0.0))
			_ = __t_tag_1
			__t_or_2 = (uint32(__t_tag_1.IntVal) == 380165415) != (true)
		}
		if __t_or_2 {
			__t3 = __local_var_1.FloatVal()
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = gopurs_runtime.UncurriedApp2(Get_Main_f(), gopurs_runtime.Float((__local_var_0.FloatVal())-(0.0)), gopurs_runtime.Float((__local_var_1.FloatVal())-(0.0))).FloatVal()
	}
end_branch_3:
	return gopurs_runtime.Float(__t3)
}

func Call_Main_f(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return gopurs_runtime.Float((gopurs_runtime.UncurriedApp2(Get_Main_g(), gopurs_runtime.Float(__local_var_0.FloatVal()), gopurs_runtime.Float(__local_var_1.FloatVal())).FloatVal()) + (gopurs_runtime.UncurriedApp2(Get_Main_g(), gopurs_runtime.Float(__local_var_1.FloatVal()), gopurs_runtime.Float(__local_var_0.FloatVal())).FloatVal()))
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_3467499472(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3467499472_1386611502(in *Constructor_Data_Show_Show[[]float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
