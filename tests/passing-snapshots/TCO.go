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
						if (v1_2) <= (int64(0)) {
							__t1 = v_1
							goto end_branch_1
						} else {

						}
					}
					{
						v_1_loop = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), v_1, v2_3)
						v1_2_loop = (v1_2) - (int64(1))
						v2_3_loop = v2_3
						continue go__go_0_0_0
						__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
			return gopurs_runtime.Apply(go__go_0_0_0, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
		}()
	})
	return cache_Main_applyN
}

var cache_Main_applyN__1066644181 gopurs_runtime.Value
var once_Main_applyN__1066644181 sync.Once

func Get_Main_applyN__1066644181() gopurs_runtime.Value {
	once_Main_applyN__1066644181.Do(func() {
		cache_Main_applyN__1066644181 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_applyN__1066644181(__eta_norm_2_0_box.IntVal, __eta_norm_1_1_box, __eta_norm_0_2_box.IntVal))
		})
	})
	return cache_Main_applyN__1066644181
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_applyN__1066644181(int64(0), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int((x_0.IntVal) + (int64(1)))
		}), int64(0)))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_applyN__1066644181(int64(1), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_1.IntVal) + (int64(1)))
			}), int64(0)))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_applyN__1066644181(int64(2), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((x_2.IntVal) + (int64(1)))
				}), int64(0)))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_applyN__1066644181(int64(3), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int((x_3.IntVal) + (int64(1)))
					}), int64(0)))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_applyN__1066644181(int64(4), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((x_4.IntVal) + (int64(1)))
						}), int64(0)))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(int64(len(Call_Data_Array_span__2676604243(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Bool(true)
							}), func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(10000))).UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()).go__init)))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Control_Monad_Rec_Class_tailRec__196765010(gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t0 gopurs_runtime.Value
									{
										if (n_6.IntVal) < (int64(10000)) {
											__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Main_2795810992_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[int64, int64]{1, (n_6.IntVal) + (int64(1))})))}
											goto end_branch_0
										} else {

										}
									}
									{
										__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Main_1462327052_3603546092((&Constructor_Control_Monad_Rec_Class_Done[int64, int64]{1, int64(42)})))}
									}
								end_branch_0:
									return __t0
								}), int64(0)))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Main_applyN__1066644181(__eta_norm_2_0_loop int64, __eta_norm_1_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop int64) int64 {
applyN__1066644181:
	for {
		if false {
			continue applyN__1066644181
		}
		var __eta_norm_2_0 int64 = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_2 int64 = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		var Call_local_Main_go__1633702519_3_0_1 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
		_ = Call_local_Main_go__1633702519_3_0_1
		var go__1633702519_3_0_1 gopurs_runtime.Value
		_ = go__1633702519_3_0_1
		var Call_local_Main_go__go_3_1_2 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
		_ = Call_local_Main_go__go_3_1_2
		var go__go_3_1_2 gopurs_runtime.Value
		_ = go__go_3_1_2
		Call_local_Main_go__1633702519_3_0_1 = func(v_4_loop gopurs_runtime.Value, v1_5_loop int64, v2_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
		go__1633702519_3_0_1:
			for {
				if false {
					continue go__1633702519_3_0_1
				}
				var v_4 gopurs_runtime.Value = v_4_loop
				_ = v_4
				var v1_5 int64 = v1_5_loop
				_ = v1_5
				var v2_6 gopurs_runtime.Value = v2_6_loop
				_ = v2_6
				var __t2 gopurs_runtime.Value
				{
					if (v1_5) <= (int64(0)) {
						__t2 = v_4
						goto end_branch_2
					} else {

					}
				}
				{
					v_4_loop = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), v_4, v2_6)
					v1_5_loop = (v1_5) - (int64(1))
					v2_6_loop = v2_6
					continue go__1633702519_3_0_1
					__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
				}
			end_branch_2:
				return __t2
			}
		}
		go__1633702519_3_0_1 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__1633702519_3_0_1(v_4_loop_val, v1_5_loop_val.IntVal, v2_6_loop_val)
				})
			})
		})
		Call_local_Main_go__go_3_1_2 = func(v_4_loop gopurs_runtime.Value, v1_5_loop int64, v2_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
		go__go_3_1_2:
			for {
				if false {
					continue go__go_3_1_2
				}
				var v_4 gopurs_runtime.Value = v_4_loop
				_ = v_4
				var v1_5 int64 = v1_5_loop
				_ = v1_5
				var v2_6 gopurs_runtime.Value = v2_6_loop
				_ = v2_6
				var __t3 gopurs_runtime.Value
				{
					if (v1_5) <= (int64(0)) {
						__t3 = v_4
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = Call_local_Main_go__1633702519_3_0_1(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), v_4, v2_6), (v1_5)-(int64(1)), v2_6)
				}
			end_branch_3:
				return __t3
			}
		}
		go__go_3_1_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_3_1_2(v_4_loop_val, v1_5_loop_val.IntVal, v2_6_loop_val)
				})
			})
		})
		return gopurs_runtime.Apply(Call_local_Main_go__1633702519_3_0_1(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), __eta_norm_2_0, __eta_norm_1_1), gopurs_runtime.Int(__eta_norm_0_2)).IntVal
	}
}

func Rebox_Main_1462327052_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[int64, int64]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_2795810992_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[int64, int64]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}
