package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_identity(x_0_box.IntVal))
		})
	})
	return cache_Main_identity
}

var cache_Main_sumTCObug_prime_ gopurs_runtime.Value
var once_Main_sumTCObug_prime_ sync.Once

func Get_Main_sumTCObug_prime_() gopurs_runtime.Value {
	once_Main_sumTCObug_prime_.Do(func() {
		cache_Main_sumTCObug_prime_ = func() gopurs_runtime.Value {
			var Call_local_Main_go__2680466838_0_0_0 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__2680466838_0_0_0
			var go__2680466838_0_0_0 gopurs_runtime.Value
			_ = go__2680466838_0_0_0
			var Call_local_Main_go__go_0_1_1 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_1_1
			var go__go_0_1_1 gopurs_runtime.Value
			_ = go__go_0_1_1
			Call_local_Main_go__2680466838_0_0_0 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64) gopurs_runtime.Value {
			go__2680466838_0_0_0:
				for {
					if false {
						continue go__2680466838_0_0_0
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var __t2 gopurs_runtime.Value
					{
						if (v1_2) == (int64(0)) {
							__t2 = v_1
							goto end_branch_2
						} else {

						}
					}
					{
						v_1_loop = gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_2) + (a_3.IntVal))
						})
						v1_2_loop = int64(0)
						continue go__2680466838_0_0_0
						__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			go__2680466838_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__2680466838_0_0_0(v_1_loop_val, v1_2_loop_val.IntVal)
				})
			})
			Call_local_Main_go__go_0_1_1 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64) gopurs_runtime.Value {
			go__go_0_1_1:
				for {
					if false {
						continue go__go_0_1_1
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var __t3 gopurs_runtime.Value
					{
						if (v1_2) == (int64(0)) {
							__t3 = v_1
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_go__2680466838_0_0_0(gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_2) + (a_3.IntVal))
						}), int64(0))
					}
				end_branch_3:
					return __t3
				}
			}
			go__go_0_1_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_0_1_1(v_1_loop_val, v1_2_loop_val.IntVal)
				})
			})
			return gopurs_runtime.Apply(go__2680466838_0_0_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}))
		}()
	})
	return cache_Main_sumTCObug_prime_
}

var cache_Main_sumTCObug gopurs_runtime.Value
var once_Main_sumTCObug sync.Once

func Get_Main_sumTCObug() gopurs_runtime.Value {
	once_Main_sumTCObug.Do(func() {
		cache_Main_sumTCObug = func() gopurs_runtime.Value {
			var Call_local_Main_go__2680466838_0_0_2 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__2680466838_0_0_2
			var go__2680466838_0_0_2 gopurs_runtime.Value
			_ = go__2680466838_0_0_2
			var Call_local_Main_go__go_0_1_3 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_1_3
			var go__go_0_1_3 gopurs_runtime.Value
			_ = go__go_0_1_3
			Call_local_Main_go__2680466838_0_0_2 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64) gopurs_runtime.Value {
			go__2680466838_0_0_2:
				for {
					if false {
						continue go__2680466838_0_0_2
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var __t2 gopurs_runtime.Value
					{
						if (v1_2) == (int64(0)) {
							__t2 = v_1
							goto end_branch_2
						} else {

						}
					}
					{
						v_1_loop = gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_2) + (a_3.IntVal))
						})
						v1_2_loop = int64(0)
						continue go__2680466838_0_0_2
						__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			go__2680466838_0_0_2 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__2680466838_0_0_2(v_1_loop_val, v1_2_loop_val.IntVal)
				})
			})
			Call_local_Main_go__go_0_1_3 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64) gopurs_runtime.Value {
			go__go_0_1_3:
				for {
					if false {
						continue go__go_0_1_3
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var __t3 gopurs_runtime.Value
					{
						if (v1_2) == (int64(0)) {
							__t3 = v_1
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_go__2680466838_0_0_2(gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_2) + (a_3.IntVal))
						}), int64(0))
					}
				end_branch_3:
					return __t3
				}
			}
			go__go_0_1_3 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_0_1_3(v_1_loop_val, v1_2_loop_val.IntVal)
				})
			})
			return gopurs_runtime.Apply(go__2680466838_0_0_2, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}))
		}()
	})
	return cache_Main_sumTCObug
}

var cache_Main_count gopurs_runtime.Value
var once_Main_count sync.Once

func Get_Main_count() gopurs_runtime.Value {
	once_Main_count.Do(func() {
		cache_Main_count = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_count(p_0_box)
		})
	})
	return cache_Main_count
}

var cache_Main_count__3963975448 gopurs_runtime.Value
var once_Main_count__3963975448 sync.Once

func Get_Main_count__3963975448() gopurs_runtime.Value {
	once_Main_count__3963975448.Do(func() {
		cache_Main_count__3963975448 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_count__3963975448(p_0_box, func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_count__3963975448
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): z_0_0 shape=App(Var) bindingType=Int
			z_0_0 := Call_Main_count__3963975448(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_0.IntVal) > (int64(0)))
			}), []int64{int64(-1), int64(0), int64(1)})
			_ = z_0_0
			var Call_local_Main_go__2680466838_1_1_8 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__2680466838_1_1_8
			var go__2680466838_1_1_8 gopurs_runtime.Value
			_ = go__2680466838_1_1_8
			var Call_local_Main_go__go_1_2_9 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_1_2_9
			var go__go_1_2_9 gopurs_runtime.Value
			_ = go__go_1_2_9
			Call_local_Main_go__2680466838_1_1_8 = func(v_2_loop gopurs_runtime.Value, v1_3_loop int64) gopurs_runtime.Value {
			go__2680466838_1_1_8:
				for {
					if false {
						continue go__2680466838_1_1_8
					}
					var v_2 gopurs_runtime.Value = v_2_loop
					_ = v_2
					var v1_3 int64 = v1_3_loop
					_ = v1_3
					var __t3 gopurs_runtime.Value
					{
						if (v1_3) == (int64(0)) {
							__t3 = v_2
							goto end_branch_3
						} else {

						}
					}
					{
						v_2_loop = gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_3) + (a_4.IntVal))
						})
						v1_3_loop = int64(0)
						continue go__2680466838_1_1_8
						__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_3:
					return __t3
				}
			}
			go__2680466838_1_1_8 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__2680466838_1_1_8(v_2_loop_val, v1_3_loop_val.IntVal)
				})
			})
			Call_local_Main_go__go_1_2_9 = func(v_2_loop gopurs_runtime.Value, v1_3_loop int64) gopurs_runtime.Value {
			go__go_1_2_9:
				for {
					if false {
						continue go__go_1_2_9
					}
					var v_2 gopurs_runtime.Value = v_2_loop
					_ = v_2
					var v1_3 int64 = v1_3_loop
					_ = v1_3
					var __t4 gopurs_runtime.Value
					{
						if (v1_3) == (int64(0)) {
							__t4 = v_2
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = Call_local_Main_go__2680466838_1_1_8(gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_3) + (a_4.IntVal))
						}), int64(0))
					}
				end_branch_4:
					return __t4
				}
			}
			go__go_1_2_9 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_1_2_9(v_2_loop_val, v1_3_loop_val.IntVal)
				})
			})
			// TAST (Let): y_2_5 shape=App(Other) bindingType=Int
			y_2_5 := gopurs_runtime.Apply(Call_local_Main_go__2680466838_1_1_8(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_2
			}), int64(7)), gopurs_runtime.Int(int64(3))).IntVal
			_ = y_2_5
			var Call_local_Main_go__2680466838_3_6_10 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__2680466838_3_6_10
			var go__2680466838_3_6_10 gopurs_runtime.Value
			_ = go__2680466838_3_6_10
			var Call_local_Main_go__go_3_7_11 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_3_7_11
			var go__go_3_7_11 gopurs_runtime.Value
			_ = go__go_3_7_11
			Call_local_Main_go__2680466838_3_6_10 = func(v_4_loop gopurs_runtime.Value, v1_5_loop int64) gopurs_runtime.Value {
			go__2680466838_3_6_10:
				for {
					if false {
						continue go__2680466838_3_6_10
					}
					var v_4 gopurs_runtime.Value = v_4_loop
					_ = v_4
					var v1_5 int64 = v1_5_loop
					_ = v1_5
					var __t8 gopurs_runtime.Value
					{
						if (v1_5) == (int64(0)) {
							__t8 = v_4
							goto end_branch_8
						} else {

						}
					}
					{
						v_4_loop = gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_5) + (a_6.IntVal))
						})
						v1_5_loop = int64(0)
						continue go__2680466838_3_6_10
						__t8 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_8:
					return __t8
				}
			}
			go__2680466838_3_6_10 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__2680466838_3_6_10(v_4_loop_val, v1_5_loop_val.IntVal)
				})
			})
			Call_local_Main_go__go_3_7_11 = func(v_4_loop gopurs_runtime.Value, v1_5_loop int64) gopurs_runtime.Value {
			go__go_3_7_11:
				for {
					if false {
						continue go__go_3_7_11
					}
					var v_4 gopurs_runtime.Value = v_4_loop
					_ = v_4
					var v1_5 int64 = v1_5_loop
					_ = v1_5
					var __t9 gopurs_runtime.Value
					{
						if (v1_5) == (int64(0)) {
							__t9 = v_4
							goto end_branch_9
						} else {

						}
					}
					{
						__t9 = Call_local_Main_go__2680466838_3_6_10(gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((v1_5) + (a_6.IntVal))
						}), int64(0))
					}
				end_branch_9:
					return __t9
				}
			}
			go__go_3_7_11 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_3_7_11(v_4_loop_val, v1_5_loop_val.IntVal)
				})
			})
			// TAST (Let): x_4_10 shape=App(Other) bindingType=Int
			x_4_10 := gopurs_runtime.Apply(Call_local_Main_go__2680466838_3_6_10(gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_4
			}), int64(7)), gopurs_runtime.Int(int64(3))).IntVal
			_ = x_4_10
			// TAST (Let): __local_var_5_11 shape=App(Var) bindingType=Any
			__local_var_5_11 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowIntImpl(x_4_10)))
			_ = __local_var_5_11
			__local_var_6_12 := gopurs_runtime.Apply(__local_var_5_11, gopurs_runtime.Value{})
			_ = __local_var_6_12
			__local_var_7_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowIntImpl(y_2_5))), gopurs_runtime.Value{})
			_ = __local_var_7_13
			__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowIntImpl(z_0_0))), gopurs_runtime.Value{})
			_ = __local_var_8_14
			var __t15 gopurs_runtime.Value
			{
				if ((x_4_10) == (int64(10))) && (((y_2_5) == (int64(10))) && ((z_0_0) == (int64(1)))) {
					__t15 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_15
				} else {

				}
			}
			{
				__t15 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail"))
			}
		end_branch_15:
			return gopurs_runtime.Apply(__t15, gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_identity(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return gopurs_runtime.Int(x_0).IntVal
}

func Call_Main_count(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var Call_local_Main_count_prime___1086705844_1_0_4 func(int64, []gopurs_runtime.Value) int64
	_ = Call_local_Main_count_prime___1086705844_1_0_4
	var count_prime___1086705844_1_0_4 gopurs_runtime.Value
	_ = count_prime___1086705844_1_0_4
	var Call_local_Main_count_prime__1_1_5 func(int64, []gopurs_runtime.Value) int64
	_ = Call_local_Main_count_prime__1_1_5
	var count_prime__1_1_5 gopurs_runtime.Value
	_ = count_prime__1_1_5
	Call_local_Main_count_prime___1086705844_1_0_4 = func(v_2_loop int64, v1_3_loop []gopurs_runtime.Value) int64 {
	count_prime___1086705844_1_0_4:
		for {
			if false {
				continue count_prime___1086705844_1_0_4
			}
			var v_2 int64 = v_2_loop
			_ = v_2
			var v1_3 []gopurs_runtime.Value = v1_3_loop
			_ = v1_3
			var __t5 int64
			{
				if (gopurs_runtime.Int(int64(len(v1_3))).IntVal) == (int64(0)) {
					__t5 = v_2
					goto end_branch_5
				} else {

				}
			}
			{
				var __t4 int64
				{
					arrayUnsafe_value_2 := v1_3
					_ = arrayUnsafe_value_2
					arrayUnsafe_source_2 := arrayUnsafe_value_2
					_ = arrayUnsafe_source_2
					arrayUnsafe_index_3 := int64(0)
					_ = arrayUnsafe_index_3
					if (gopurs_runtime.Apply(p_0, (arrayUnsafe_source_2)[arrayUnsafe_index_3]).IntVal) != (0) {
						__t4 = (v_2) + (int64(1))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = (v_2) + (int64(0))
				}
			end_branch_4:
				v_2_loop = __t4
				v1_3_loop = (*(*[]gopurs_runtime.Value)((gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(len(v1_3))), gopurs_runtime.Array(v1_3))).UnsafePtr))
				continue count_prime___1086705844_1_0_4
				__t5 = func() int64 { panic("unreachable") }()
			}
		end_branch_5:
			return __t5
		}
	}
	count_prime___1086705844_1_0_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_count_prime___1086705844_1_0_4(v_2_loop_val.IntVal, (*(*[]gopurs_runtime.Value)((v1_3_loop_val).UnsafePtr))))
		})
	})
	Call_local_Main_count_prime__1_1_5 = func(v_2_loop int64, v1_3_loop []gopurs_runtime.Value) int64 {
	count_prime__1_1_5:
		for {
			if false {
				continue count_prime__1_1_5
			}
			var v_2 int64 = v_2_loop
			_ = v_2
			var v1_3 []gopurs_runtime.Value = v1_3_loop
			_ = v1_3
			var __t9 int64
			{
				if (gopurs_runtime.Int(int64(len(v1_3))).IntVal) == (int64(0)) {
					__t9 = v_2
					goto end_branch_9
				} else {

				}
			}
			{
				var __t8 int64
				{
					arrayUnsafe_value_6 := v1_3
					_ = arrayUnsafe_value_6
					arrayUnsafe_source_6 := arrayUnsafe_value_6
					_ = arrayUnsafe_source_6
					arrayUnsafe_index_7 := int64(0)
					_ = arrayUnsafe_index_7
					if (gopurs_runtime.Apply(p_0, (arrayUnsafe_source_6)[arrayUnsafe_index_7]).IntVal) != (0) {
						__t8 = (v_2) + (int64(1))
						goto end_branch_8
					} else {

					}
				}
				{
					__t8 = (v_2) + (int64(0))
				}
			end_branch_8:
				__t9 = Call_local_Main_count_prime___1086705844_1_0_4(__t8, (*(*[]gopurs_runtime.Value)((gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(len(v1_3))), gopurs_runtime.Array(v1_3))).UnsafePtr)))
			}
		end_branch_9:
			return __t9
		}
	}
	count_prime__1_1_5 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_count_prime__1_1_5(v_2_loop_val.IntVal, (*(*[]gopurs_runtime.Value)((v1_3_loop_val).UnsafePtr))))
		})
	})
	return gopurs_runtime.Apply(count_prime___1086705844_1_0_4, gopurs_runtime.Int(int64(0)))
}

func Call_Main_count__3963975448(p_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop []int64) int64 {
count__3963975448:
	for {
		if false {
			continue count__3963975448
		}
		var p_0 gopurs_runtime.Value = p_0_loop
		_ = p_0
		var __eta_norm_0_1 []int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var Call_local_Main_count_prime___4172905617_2_0_6 func(int64, []int64) int64
		_ = Call_local_Main_count_prime___4172905617_2_0_6
		var count_prime___4172905617_2_0_6 gopurs_runtime.Value
		_ = count_prime___4172905617_2_0_6
		var Call_local_Main_count_prime__2_1_7 func(int64, []int64) int64
		_ = Call_local_Main_count_prime__2_1_7
		var count_prime__2_1_7 gopurs_runtime.Value
		_ = count_prime__2_1_7
		Call_local_Main_count_prime___4172905617_2_0_6 = func(v_3_loop int64, v1_4_loop []int64) int64 {
		count_prime___4172905617_2_0_6:
			for {
				if false {
					continue count_prime___4172905617_2_0_6
				}
				var v_3 int64 = v_3_loop
				_ = v_3
				var v1_4 []int64 = v1_4_loop
				_ = v1_4
				var __t5 int64
				{
					if (gopurs_runtime.Int(int64(len(v1_4))).IntVal) == (int64(0)) {
						__t5 = v_3
						goto end_branch_5
					} else {

					}
				}
				{
					var __t4 int64
					{
						arrayUnsafe_value_2 := v1_4
						_ = arrayUnsafe_value_2
						arrayUnsafe_source_2 := arrayUnsafe_value_2
						_ = arrayUnsafe_source_2
						arrayUnsafe_index_3 := int64(0)
						_ = arrayUnsafe_index_3
						if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int(gopurs_runtime.Int((arrayUnsafe_source_2)[arrayUnsafe_index_3]).IntVal)).IntVal) != (0) {
							__t4 = (v_3) + (int64(1))
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = (v_3) + (int64(0))
					}
				end_branch_4:
					v_3_loop = __t4
					v1_4_loop = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(v1_4))).IntVal), func() gopurs_runtime.Value {
							arr := v1_4
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()).UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					continue count_prime___4172905617_2_0_6
					__t5 = func() int64 { panic("unreachable") }()
				}
			end_branch_5:
				return __t5
			}
		}
		count_prime___4172905617_2_0_6 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Main_count_prime___4172905617_2_0_6(v_3_loop_val.IntVal, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v1_4_loop_val.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()))
			})
		})
		Call_local_Main_count_prime__2_1_7 = func(v_3_loop int64, v1_4_loop []int64) int64 {
		count_prime__2_1_7:
			for {
				if false {
					continue count_prime__2_1_7
				}
				var v_3 int64 = v_3_loop
				_ = v_3
				var v1_4 []int64 = v1_4_loop
				_ = v1_4
				var __t9 int64
				{
					if (gopurs_runtime.Int(int64(len(v1_4))).IntVal) == (int64(0)) {
						__t9 = v_3
						goto end_branch_9
					} else {

					}
				}
				{
					var __t8 int64
					{
						arrayUnsafe_value_6 := v1_4
						_ = arrayUnsafe_value_6
						arrayUnsafe_source_6 := arrayUnsafe_value_6
						_ = arrayUnsafe_source_6
						arrayUnsafe_index_7 := int64(0)
						_ = arrayUnsafe_index_7
						if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int(gopurs_runtime.Int((arrayUnsafe_source_6)[arrayUnsafe_index_7]).IntVal)).IntVal) != (0) {
							__t8 = (v_3) + (int64(1))
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = (v_3) + (int64(0))
					}
				end_branch_8:
					__t9 = Call_local_Main_count_prime___4172905617_2_0_6(__t8, func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(v1_4))).IntVal), func() gopurs_runtime.Value {
							arr := v1_4
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()).UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}())
				}
			end_branch_9:
				return __t9
			}
		}
		count_prime__2_1_7 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Main_count_prime__2_1_7(v_3_loop_val.IntVal, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v1_4_loop_val.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()))
			})
		})
		return Call_local_Main_count_prime___4172905617_2_0_6(int64(0), __eta_norm_0_1)
	}
}
