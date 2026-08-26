package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_applyN gopurs_runtime.Value
var once_Main_applyN sync.Once

func Get_Main_applyN() gopurs_runtime.Value {
	once_Main_applyN.Do(func() {
		cache_Main_applyN = func() gopurs_runtime.Value {
			var Call_local_Main_go__go_0_0_0 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_0_0
			var go__go_0_0_0 gopurs_runtime.Value
			_ = go__go_0_0_0
			Call_local_Main_go__go_0_0_0 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64, v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_0_0_0:
				for {
					if false {
						continue go__go_0_0_0
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var v2_3 gopurs_runtime.Value = v2_3_loop
					_ = v2_3
					var __t1 gopurs_runtime.Value
					{
						if (v1_2) <= (0) {
							__t1 = v_1
							goto end_branch_1
						} else {

						}
					}
					{
						v_1_loop = gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(v2_3, gopurs_runtime.Apply(v_1, x_4))
						})
						v1_2_loop = (v1_2) - (1)
						v2_3_loop = v2_3
						continue go__go_0_0_0
						__t1 = gopurs_runtime.Value{}
					}
				end_branch_1:
					return __t1
				}
			}
			go__go_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_0_0_0(v_1_loop_val, v1_2_loop_val.IntVal, v2_3_loop_val)
					})
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}))
		}()
	})
	return cache_Main_applyN
}

var cache_Main_applyN__1580697421 gopurs_runtime.Value
var once_Main_applyN__1580697421 sync.Once

func Get_Main_applyN__1580697421() gopurs_runtime.Value {
	once_Main_applyN__1580697421.Do(func() {
		cache_Main_applyN__1580697421 = func() gopurs_runtime.Value {
			var Call_local_Main_go__go_0_0_1 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_0_1
			var go__go_0_0_1 gopurs_runtime.Value
			_ = go__go_0_0_1
			Call_local_Main_go__go_0_0_1 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64, v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_0_0_1:
				for {
					if false {
						continue go__go_0_0_1
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var v2_3 gopurs_runtime.Value = v2_3_loop
					_ = v2_3
					var __t1 gopurs_runtime.Value
					{
						if (v1_2) <= (0) {
							__t1 = v_1
							goto end_branch_1
						} else {

						}
					}
					{
						v_1_loop = gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(v2_3, gopurs_runtime.Apply(v_1, x_4))
						})
						v1_2_loop = (v1_2) - (1)
						v2_3_loop = v2_3
						continue go__go_0_0_1
						__t1 = gopurs_runtime.Value{}
					}
				end_branch_1:
					return __t1
				}
			}
			go__go_0_0_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_0_0_1(v_1_loop_val, v1_2_loop_val.IntVal, v2_3_loop_val)
					})
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}))
		}()
	})
	return cache_Main_applyN__1580697421
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Main_go__go_0_1_2 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_1_2
			var go__go_0_1_2 gopurs_runtime.Value
			_ = go__go_0_1_2
			Call_local_Main_go__go_0_1_2 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64, v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_0_1_2:
				for {
					if false {
						continue go__go_0_1_2
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var v2_3 gopurs_runtime.Value = v2_3_loop
					_ = v2_3
					var __t2 gopurs_runtime.Value
					{
						if (v1_2) <= (0) {
							__t2 = v_1
							goto end_branch_2
						} else {

						}
					}
					{
						v_1_loop = gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(v2_3, gopurs_runtime.Apply(v_1, x_4))
						})
						v1_2_loop = (v1_2) - (1)
						v2_3_loop = v2_3
						continue go__go_0_1_2
						__t2 = gopurs_runtime.Value{}
					}
				end_branch_2:
					return __t2
				}
			}
			go__go_0_1_2 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_0_1_2(v_1_loop_val, v1_2_loop_val.IntVal, v2_3_loop_val)
					})
				})
			})
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Call_local_Main_go__go_0_1_2(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}), 0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_1.IntVal) + (1))
			})), gopurs_runtime.Int(0))).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_3 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_3
			var Call_local_Main_go__go_2_5_3 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_2_5_3
			var go__go_2_5_3 gopurs_runtime.Value
			_ = go__go_2_5_3
			Call_local_Main_go__go_2_5_3 = func(v_3_loop gopurs_runtime.Value, v1_4_loop int64, v2_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_2_5_3:
				for {
					if false {
						continue go__go_2_5_3
					}
					var v_3 gopurs_runtime.Value = v_3_loop
					_ = v_3
					var v1_4 int64 = v1_4_loop
					_ = v1_4
					var v2_5 gopurs_runtime.Value = v2_5_loop
					_ = v2_5
					var __t6 gopurs_runtime.Value
					{
						if (v1_4) <= (0) {
							__t6 = v_3
							goto end_branch_6
						} else {

						}
					}
					{
						v_3_loop = gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(v2_5, gopurs_runtime.Apply(v_3, x_6))
						})
						v1_4_loop = (v1_4) - (1)
						v2_5_loop = v2_5
						continue go__go_2_5_3
						__t6 = gopurs_runtime.Value{}
					}
				end_branch_6:
					return __t6
				}
			}
			go__go_2_5_3 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_2_5_3(v_3_loop_val, v1_4_loop_val.IntVal, v2_5_loop_val)
					})
				})
			})
			_dollar___unused_2_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Call_local_Main_go__go_2_5_3(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_3
			}), 1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_3.IntVal) + (1))
			})), gopurs_runtime.Int(0))).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_2_4
			var Call_local_Main_go__go_3_8_4 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_3_8_4
			var go__go_3_8_4 gopurs_runtime.Value
			_ = go__go_3_8_4
			Call_local_Main_go__go_3_8_4 = func(v_4_loop gopurs_runtime.Value, v1_5_loop int64, v2_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_3_8_4:
				for {
					if false {
						continue go__go_3_8_4
					}
					var v_4 gopurs_runtime.Value = v_4_loop
					_ = v_4
					var v1_5 int64 = v1_5_loop
					_ = v1_5
					var v2_6 gopurs_runtime.Value = v2_6_loop
					_ = v2_6
					var __t9 gopurs_runtime.Value
					{
						if (v1_5) <= (0) {
							__t9 = v_4
							goto end_branch_9
						} else {

						}
					}
					{
						v_4_loop = gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(v2_6, gopurs_runtime.Apply(v_4, x_7))
						})
						v1_5_loop = (v1_5) - (1)
						v2_6_loop = v2_6
						continue go__go_3_8_4
						__t9 = gopurs_runtime.Value{}
					}
				end_branch_9:
					return __t9
				}
			}
			go__go_3_8_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_3_8_4(v_4_loop_val, v1_5_loop_val.IntVal, v2_6_loop_val)
					})
				})
			})
			_dollar___unused_3_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Call_local_Main_go__go_3_8_4(gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_4
			}), 2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_4.IntVal) + (1))
			})), gopurs_runtime.Int(0))).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_3_7
			var Call_local_Main_go__go_4_11_5 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_4_11_5
			var go__go_4_11_5 gopurs_runtime.Value
			_ = go__go_4_11_5
			Call_local_Main_go__go_4_11_5 = func(v_5_loop gopurs_runtime.Value, v1_6_loop int64, v2_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_4_11_5:
				for {
					if false {
						continue go__go_4_11_5
					}
					var v_5 gopurs_runtime.Value = v_5_loop
					_ = v_5
					var v1_6 int64 = v1_6_loop
					_ = v1_6
					var v2_7 gopurs_runtime.Value = v2_7_loop
					_ = v2_7
					var __t12 gopurs_runtime.Value
					{
						if (v1_6) <= (0) {
							__t12 = v_5
							goto end_branch_12
						} else {

						}
					}
					{
						v_5_loop = gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(v2_7, gopurs_runtime.Apply(v_5, x_8))
						})
						v1_6_loop = (v1_6) - (1)
						v2_7_loop = v2_7
						continue go__go_4_11_5
						__t12 = gopurs_runtime.Value{}
					}
				end_branch_12:
					return __t12
				}
			}
			go__go_4_11_5 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_4_11_5(v_5_loop_val, v1_6_loop_val.IntVal, v2_7_loop_val)
					})
				})
			})
			_dollar___unused_4_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Call_local_Main_go__go_4_11_5(gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_5
			}), 3, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_5.IntVal) + (1))
			})), gopurs_runtime.Int(0))).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_4_10
			var Call_local_Main_go__go_5_14_6 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_5_14_6
			var go__go_5_14_6 gopurs_runtime.Value
			_ = go__go_5_14_6
			Call_local_Main_go__go_5_14_6 = func(v_6_loop gopurs_runtime.Value, v1_7_loop int64, v2_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_5_14_6:
				for {
					if false {
						continue go__go_5_14_6
					}
					var v_6 gopurs_runtime.Value = v_6_loop
					_ = v_6
					var v1_7 int64 = v1_7_loop
					_ = v1_7
					var v2_8 gopurs_runtime.Value = v2_8_loop
					_ = v2_8
					var __t15 gopurs_runtime.Value
					{
						if (v1_7) <= (0) {
							__t15 = v_6
							goto end_branch_15
						} else {

						}
					}
					{
						v_6_loop = gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(v2_8, gopurs_runtime.Apply(v_6, x_9))
						})
						v1_7_loop = (v1_7) - (1)
						v2_8_loop = v2_8
						continue go__go_5_14_6
						__t15 = gopurs_runtime.Value{}
					}
				end_branch_15:
					return __t15
				}
			}
			go__go_5_14_6 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_5_14_6(v_6_loop_val, v1_7_loop_val.IntVal, v2_8_loop_val)
					})
				})
			})
			_dollar___unused_5_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Call_local_Main_go__go_5_14_6(gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_6
			}), 4, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_6.IntVal) + (1))
			})), gopurs_runtime.Int(0))).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_5_13
			_dollar___unused_6_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Array_span(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(10000)).UnsafePtr)
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
			}()), "init")))).IntVal)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_6_16
			var Call_local_Main_go__go_7_18_7 func(gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_7_18_7
			var go__go_7_18_7 gopurs_runtime.Value
			_ = go__go_7_18_7
			Call_local_Main_go__go_7_18_7 = func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_7_18_7:
				for {
					if false {
						continue go__go_7_18_7
					}
					var v_8 gopurs_runtime.Value = v_8_loop
					_ = v_8
					var __t20 gopurs_runtime.Value
					{
						if v_8.Type == 9 && v_8.IntVal == 525585346 {
							var __t19 gopurs_runtime.Value
							{
								if ((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0.IntVal) < (10000) {
									__t19 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[int64, int64]{1, gopurs_runtime.Int(((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0.IntVal) + (1))}))}
									goto end_branch_19
								} else {

								}
							}
							{
								__t19 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[int64, int64]{1, gopurs_runtime.Int(42)}))}
							}
						end_branch_19:
							v_8_loop = __t19
							continue go__go_7_18_7
							__t20 = gopurs_runtime.Value{}
							goto end_branch_20
						} else {

						}
					}
					{
						if v_8.Type == 9 && v_8.IntVal == 60402430 {
							__t20 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
							goto end_branch_20
						} else {

						}
					}
					{
						__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_20:
					return __t20
				}
			}
			go__go_7_18_7 = gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_local_Main_go__go_7_18_7(v_8_loop_val)
			})
			_dollar___unused_7_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), Call_local_Main_go__go_7_18_7(gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[int64, int64]{1, gopurs_runtime.Int(1)}))})).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_7_17
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}
