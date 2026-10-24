package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_sumTCObug_prime_ gopurs_runtime.Value
var once_Main_sumTCObug_prime_ sync.Once

func Get_Main_sumTCObug_prime_() gopurs_runtime.Value {
	once_Main_sumTCObug_prime_.Do(func() {
		cache_Main_sumTCObug_prime_ = func() gopurs_runtime.Value {
			var Call_local_Main_go__go_0_0_0 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_0_0
			var go__go_0_0_0 gopurs_runtime.Value
			_ = go__go_0_0_0
			Call_local_Main_go__go_0_0_0 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64) gopurs_runtime.Value {
			go__go_0_0_0:
				for {
					if false {
						continue go__go_0_0_0
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var __t1 gopurs_runtime.Value
					{
						if (v1_2) == (int64(0)) {
							__t1 = v_1
							goto end_branch_1
						} else {

						}
					}
					{
						v_1_loop = gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_2, a_3.IntVal))
						})
						v1_2_loop = int64(0)
						continue go__go_0_0_0
						__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			go__go_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_0_0_0(v_1_loop_val, v1_2_loop_val.IntVal)
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
			var Call_local_Main_go__go_0_0_1 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_0_1
			var go__go_0_0_1 gopurs_runtime.Value
			_ = go__go_0_0_1
			Call_local_Main_go__go_0_0_1 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64) gopurs_runtime.Value {
			go__go_0_0_1:
				for {
					if false {
						continue go__go_0_0_1
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var __t1 gopurs_runtime.Value
					{
						if (v1_2) == (int64(0)) {
							__t1 = v_1
							goto end_branch_1
						} else {

						}
					}
					{
						v_1_loop = gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_2, a_3.IntVal))
						})
						v1_2_loop = int64(0)
						continue go__go_0_0_1
						__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			go__go_0_0_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_0_0_1(v_1_loop_val, v1_2_loop_val.IntVal)
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
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

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): z_0_0 shape=App(Var) bindingType=Int
			z_0_0 := gopurs_runtime.Apply(Call_Main_count(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_0.IntVal) > (int64(0)))
			})), func() gopurs_runtime.Value {
				arr := []int64{int64(-1), int64(0), int64(1)}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal
			_ = z_0_0
			var Call_local_Main_go__go_1_1_3 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_1_1_3
			var go__go_1_1_3 gopurs_runtime.Value
			_ = go__go_1_1_3
			Call_local_Main_go__go_1_1_3 = func(v_2_loop gopurs_runtime.Value, v1_3_loop int64) gopurs_runtime.Value {
			go__go_1_1_3:
				for {
					if false {
						continue go__go_1_1_3
					}
					var v_2 gopurs_runtime.Value = v_2_loop
					_ = v_2
					var v1_3 int64 = v1_3_loop
					_ = v1_3
					var __t2 gopurs_runtime.Value
					{
						if (v1_3) == (int64(0)) {
							__t2 = v_2
							goto end_branch_2
						} else {

						}
					}
					{
						v_2_loop = gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_3, a_4.IntVal))
						})
						v1_3_loop = int64(0)
						continue go__go_1_1_3
						__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			go__go_1_1_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_1_1_3(v_2_loop_val, v1_3_loop_val.IntVal)
				})
			})
			// TAST (Let): y_2_3 shape=App(Other) bindingType=Int
			y_2_3 := gopurs_runtime.Apply(Call_local_Main_go__go_1_1_3(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_2
			}), int64(7)), gopurs_runtime.Int(int64(3))).IntVal
			_ = y_2_3
			var Call_local_Main_go__go_3_4_4 func(gopurs_runtime.Value, int64) gopurs_runtime.Value
			_ = Call_local_Main_go__go_3_4_4
			var go__go_3_4_4 gopurs_runtime.Value
			_ = go__go_3_4_4
			Call_local_Main_go__go_3_4_4 = func(v_4_loop gopurs_runtime.Value, v1_5_loop int64) gopurs_runtime.Value {
			go__go_3_4_4:
				for {
					if false {
						continue go__go_3_4_4
					}
					var v_4 gopurs_runtime.Value = v_4_loop
					_ = v_4
					var v1_5 int64 = v1_5_loop
					_ = v1_5
					var __t5 gopurs_runtime.Value
					{
						if (v1_5) == (int64(0)) {
							__t5 = v_4
							goto end_branch_5
						} else {

						}
					}
					{
						v_4_loop = gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_5, a_6.IntVal))
						})
						v1_5_loop = int64(0)
						continue go__go_3_4_4
						__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_5:
					return __t5
				}
			}
			go__go_3_4_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_3_4_4(v_4_loop_val, v1_5_loop_val.IntVal)
				})
			})
			// TAST (Let): x_4_6 shape=App(Other) bindingType=Int
			x_4_6 := gopurs_runtime.Apply(Call_local_Main_go__go_3_4_4(gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_4
			}), int64(7)), gopurs_runtime.Int(int64(3))).IntVal
			_ = x_4_6
			// TAST (Let): __local_var_5_7 shape=App(Var) bindingType=Any
			__local_var_5_7 := Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(x_4_6))
			_ = __local_var_5_7
			_dollar___unused_6_8 := gopurs_runtime.Apply(__local_var_5_7, gopurs_runtime.Value{})
			_ = _dollar___unused_6_8
			_dollar___unused_7_9 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(y_2_3)), gopurs_runtime.Value{})
			_ = _dollar___unused_7_9
			_dollar___unused_8_10 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(z_0_0)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_10
			var __t11 gopurs_runtime.Value
			{
				if ((x_4_6) == (int64(10))) && (((y_2_3) == (int64(10))) && ((z_0_0) == (int64(1)))) {
					__t11 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail"))
			}
		end_branch_11:
			return gopurs_runtime.Apply(__t11, gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_count(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var Call_local_Main_count_prime__1_0_2 func(int64, []gopurs_runtime.Value) int64
	_ = Call_local_Main_count_prime__1_0_2
	var count_prime__1_0_2 gopurs_runtime.Value
	_ = count_prime__1_0_2
	Call_local_Main_count_prime__1_0_2 = func(v_2_loop int64, v1_3_loop []gopurs_runtime.Value) int64 {
	count_prime__1_0_2:
		for {
			if false {
				continue count_prime__1_0_2
			}
			var v_2 int64 = v_2_loop
			_ = v_2
			var v1_3 []gopurs_runtime.Value = v1_3_loop
			_ = v1_3
			var __t2 int64
			{
				if (gopurs_runtime.Int(int64(len(v1_3))).IntVal) == (int64(0)) {
					__t2 = v_2
					goto end_branch_2
				} else {

				}
			}
			{
				var __t1 int64
				{
					if (gopurs_runtime.Apply(p_0, gopurs_runtime.Apply2(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Data_Array_Partial_head(), gopurs_runtime.Value{})
					}), gopurs_runtime.Array(v1_3))).IntVal) != (0) {
						__t1 = gopurs_runtime.IntAdd(v_2, int64(1))
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = gopurs_runtime.IntAdd(v_2, int64(0))
				}
			end_branch_1:
				v_2_loop = __t1
				v1_3_loop = (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Data_Array_Partial_tail(), gopurs_runtime.Value{})
				}), gopurs_runtime.Array(v1_3))).UnsafePtr))
				continue count_prime__1_0_2
				__t2 = func() int64 { panic("unreachable") }()
			}
		end_branch_2:
			return __t2
		}
	}
	count_prime__1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_count_prime__1_0_2(v_2_loop_val.IntVal, (*(*[]gopurs_runtime.Value)((v1_3_loop_val).UnsafePtr))))
		})
	})
	return gopurs_runtime.Apply(count_prime__1_0_2, gopurs_runtime.Int(int64(0)))
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
