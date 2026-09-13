package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Call_Data_Semiring_add(Rebox_Main_602713622_2826095630(Rebox_Main_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber()))))
	})
	return cache_Main_add
}

var cache_Main_test7 gopurs_runtime.Value
var once_Main_test7 sync.Once

func Get_Main_test7() gopurs_runtime.Value {
	once_Main_test7.Do(func() {
		cache_Main_test7 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test7(x_0_box.FloatVal()))
		})
	})
	return cache_Main_test7
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Float(1.0)
	})
	return cache_Main_test6
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = func() gopurs_runtime.Value {
			var g__512709115_0_0_2 gopurs_runtime.Value
			_ = g__512709115_0_0_2
			var g__512709115_0_0_2_cell *gopurs_runtime.Value
			_ = g__512709115_0_0_2_cell
			// FALLBACK TCO: isLoop=false len=3
			var g_0_1_3 gopurs_runtime.Value
			_ = g_0_1_3
			var g_0_1_3_cell *gopurs_runtime.Value
			_ = g_0_1_3_cell
			// FALLBACK TCO: isLoop=false len=3
			var f_0_2_4 gopurs_runtime.Value
			_ = f_0_2_4
			var f_0_2_4_cell *gopurs_runtime.Value
			_ = f_0_2_4_cell
			// FALLBACK TCO: isLoop=false len=3
			g__512709115_0_0_2 = gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float((gopurs_runtime.Apply((*f_0_2_4_cell), gopurs_runtime.Float((x_1.FloatVal())-(1.0))).FloatVal()) + (1.0))
			})
			g__512709115_0_0_2_cell = &g__512709115_0_0_2
			g_0_1_3 = gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float((gopurs_runtime.Apply((*f_0_2_4_cell), gopurs_runtime.Float((x_1.FloatVal())-(1.0))).FloatVal()) + (1.0))
			})
			g_0_1_3_cell = &g_0_1_3
			f_0_2_4 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 float64
				{
					var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, v_1, gopurs_runtime.Float(0.0))
					_ = __t_tag_3
					if uint32(__t_tag_3.IntVal) == 380165415 {
						__t4 = (gopurs_runtime.Apply((*g__512709115_0_0_2_cell), gopurs_runtime.Float((v_1.FloatVal())/(2.0))).FloatVal()) + (1.0)
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = 0.0
				}
			end_branch_4:
				return gopurs_runtime.Float(__t4)
			})
			f_0_2_4_cell = &f_0_2_4
			return gopurs_runtime.Float(gopurs_runtime.Apply(g__512709115_0_0_2, gopurs_runtime.Float(10.0)).FloatVal())
		}()
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func(func(dictPartial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test4(dictPartial_0_box)
		})
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Float(6.0)
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test2(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test1(x_0_box.FloatVal()))
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(2.0)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(5.0)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(6.0)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_3_0 shape=LitArray bindingType=(Array Number)
					__local_var_3_0 := []float64{1.0, 2.0}
					_ = __local_var_3_0
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.Apply2(Call_Data_Semiring_add(Rebox_Main_602713622_2826095630(Rebox_Main_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber())))), gopurs_runtime.Float(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := __local_var_3_0
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Float(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 0).FloatVal()), gopurs_runtime.Float(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := __local_var_3_0
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Float(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 1).FloatVal())).FloatVal())).StrVal())), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(Get_Main_test5().FloatVal())).StrVal())), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(1.0)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(Call_Main_test7(100.0))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
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

func Call_Main_test7(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	var Call_local_Main_go__512709115_1_0_0 func(float64) float64
	_ = Call_local_Main_go__512709115_1_0_0
	var go__512709115_1_0_0 gopurs_runtime.Value
	_ = go__512709115_1_0_0
	var Call_local_Main_go__go_1_1_1 func(float64) float64
	_ = Call_local_Main_go__go_1_1_1
	var go__go_1_1_1 gopurs_runtime.Value
	_ = go__go_1_1_1
	Call_local_Main_go__512709115_1_0_0 = func(v_2_loop float64) float64 {
	go__512709115_1_0_0:
		for {
			if false {
				continue go__512709115_1_0_0
			}
			var v_2 float64 = v_2_loop
			_ = v_2
			var __t5 float64
			{
				var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((x_0)-(0.1)), gopurs_runtime.Float((v_2)*(v_2)))
				_ = __t_tag_2
				var __t_and_4 bool = false
				if uint32(__t_tag_2.IntVal) == 1527465420 {

					var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((v_2)*(v_2)), gopurs_runtime.Float((x_0)+(0.1)))
					_ = __t_tag_3
					__t_and_4 = (uint32(__t_tag_3.IntVal) == 1527465420)
				}
				if __t_and_4 {
					__t5 = v_2
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = gopurs_runtime.Float(Call_local_Main_go__go_1_1_1(((v_2) + ((x_0) / (v_2))) / (2.0))).FloatVal()
			}
		end_branch_5:
			return __t5
		}
	}
	go__512709115_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(Call_local_Main_go__512709115_1_0_0(v_2_loop_val.FloatVal()))
	})
	Call_local_Main_go__go_1_1_1 = func(v_2_loop float64) float64 {
	go__go_1_1_1:
		for {
			if false {
				continue go__go_1_1_1
			}
			var v_2 float64 = v_2_loop
			_ = v_2
			var __t9 float64
			{
				var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((x_0)-(0.1)), gopurs_runtime.Float((v_2)*(v_2)))
				_ = __t_tag_6
				var __t_and_8 bool = false
				if uint32(__t_tag_6.IntVal) == 1527465420 {

					var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((v_2)*(v_2)), gopurs_runtime.Float((x_0)+(0.1)))
					_ = __t_tag_7
					__t_and_8 = (uint32(__t_tag_7.IntVal) == 1527465420)
				}
				if __t_and_8 {
					__t9 = v_2
					goto end_branch_9
				} else {

				}
			}
			{
				v_2_loop = ((v_2) + ((x_0) / (v_2))) / (2.0)
				continue go__go_1_1_1
				__t9 = func() gopurs_runtime.Value { panic("unreachable") }().FloatVal()
			}
		end_branch_9:
			return __t9
		}
	}
	go__go_1_1_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(Call_local_Main_go__go_1_1_1(v_2_loop_val.FloatVal()))
	})
	return Call_local_Main_go__512709115_1_0_0(x_0)
}

func Call_Main_test4(dictPartial_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPartial_0 gopurs_runtime.Value = dictPartial_0_loop
	_ = dictPartial_0
	// TAST (Let): __local_var_1_0 shape=LitArray bindingType=(Array Number)
	__local_var_1_0 := []float64{1.0, 2.0}
	_ = __local_var_1_0
	return gopurs_runtime.Float(gopurs_runtime.Apply2(Call_Data_Semiring_add(Rebox_Main_602713622_2826095630(Rebox_Main_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber())))), gopurs_runtime.Float(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
		arr := __local_var_1_0
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Float(v)
		}
		return gopurs_runtime.Array(boxed)
	}(), 0).FloatVal()), gopurs_runtime.Float(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
		arr := __local_var_1_0
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Float(v)
		}
		return gopurs_runtime.Array(boxed)
	}(), 1).FloatVal())).FloatVal())
}

func Call_Main_test2(x_0_loop float64, y_1_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	return (((x_0) + (1.0)) + (y_1)) + (1.0)
}

func Call_Main_test1(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return (x_0) + (1.0)
}

func Rebox_Main_2826095630_602713622(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[float64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.FloatVal()
	out.V3 = in.V3.FloatVal()
	return out
}

func Rebox_Main_602713622_2826095630(in *Constructor_Data_Semiring_Semiring[float64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Float(in.V2)
	out.V3 = gopurs_runtime.Float(in.V3)
	return out
}
