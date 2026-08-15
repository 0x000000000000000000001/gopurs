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
			var go__go_0_0_0 gopurs_runtime.Value
			go__go_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var v_1_loop gopurs_runtime.Value = v_1_loop_val
						var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
					go__go_0_0_0:
						for {
							if false {
								continue go__go_0_0_0
							}
							var v_1 gopurs_runtime.Value = v_1_loop
							_ = v_1
							var v1_2 gopurs_runtime.Value = v1_2_loop
							_ = v1_2
							var __t1 gopurs_runtime.Value
							{
								if (v1_2.IntVal) == (0) {
									__t1 = v_1
									goto end_branch_1
								} else {

								}
							}
							{
								v_1_loop = gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Int((v1_2.IntVal) + (a_3.IntVal))
								})
								v1_2_loop = gopurs_runtime.Int(0)
								continue go__go_0_0_0
								__t1 = gopurs_runtime.Value{}
							}
						end_branch_1:
							return __t1
						}
					}()
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
			var go__go_0_0_1 gopurs_runtime.Value
			go__go_0_0_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var v_1_loop gopurs_runtime.Value = v_1_loop_val
						var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
					go__go_0_0_1:
						for {
							if false {
								continue go__go_0_0_1
							}
							var v_1 gopurs_runtime.Value = v_1_loop
							_ = v_1
							var v1_2 gopurs_runtime.Value = v1_2_loop
							_ = v1_2
							var __t1 gopurs_runtime.Value
							{
								if (v1_2.IntVal) == (0) {
									__t1 = v_1
									goto end_branch_1
								} else {

								}
							}
							{
								v_1_loop = gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Int((v1_2.IntVal) + (a_3.IntVal))
								})
								v1_2_loop = gopurs_runtime.Int(0)
								continue go__go_0_0_1
								__t1 = gopurs_runtime.Value{}
							}
						end_branch_1:
							return __t1
						}
					}()
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
			// TAST (Let): z_0_0 -> int64
			z_0_0 := gopurs_runtime.Apply(Call_Main_count(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v_0.IntVal) > (0))
			})), func() gopurs_runtime.Value {
				arr := []int64{gopurs_runtime.Int(-1).IntVal, 0, 1}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal
			_ = z_0_0
			var go__go_1_2_3 gopurs_runtime.Value
			go__go_1_2_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var v_2_loop gopurs_runtime.Value = v_2_loop_val
						var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
					go__go_1_2_3:
						for {
							if false {
								continue go__go_1_2_3
							}
							var v_2 gopurs_runtime.Value = v_2_loop
							_ = v_2
							var v1_3 gopurs_runtime.Value = v1_3_loop
							_ = v1_3
							var __t3 gopurs_runtime.Value
							{
								if (v1_3.IntVal) == (0) {
									__t3 = v_2
									goto end_branch_3
								} else {

								}
							}
							{
								v_2_loop = gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Int((v1_3.IntVal) + (a_4.IntVal))
								})
								v1_3_loop = gopurs_runtime.Int(0)
								continue go__go_1_2_3
								__t3 = gopurs_runtime.Value{}
							}
						end_branch_3:
							return __t3
						}
					}()
				})
			})
			// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
			__local_var_1_1 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply3(go__go_1_2_3, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_2
			}), gopurs_runtime.Int(7), gopurs_runtime.Int(3)).IntVal)).StrVal()))
			_ = __local_var_1_1
			_dollar___unused_2_4 := gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Value{})
			_ = _dollar___unused_2_4
			var go__go_3_6_4 gopurs_runtime.Value
			go__go_3_6_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var v_4_loop gopurs_runtime.Value = v_4_loop_val
						var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
					go__go_3_6_4:
						for {
							if false {
								continue go__go_3_6_4
							}
							var v_4 gopurs_runtime.Value = v_4_loop
							_ = v_4
							var v1_5 gopurs_runtime.Value = v1_5_loop
							_ = v1_5
							var __t7 gopurs_runtime.Value
							{
								if (v1_5.IntVal) == (0) {
									__t7 = v_4
									goto end_branch_7
								} else {

								}
							}
							{
								v_4_loop = gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Int((v1_5.IntVal) + (a_6.IntVal))
								})
								v1_5_loop = gopurs_runtime.Int(0)
								continue go__go_3_6_4
								__t7 = gopurs_runtime.Value{}
							}
						end_branch_7:
							return __t7
						}
					}()
				})
			})
			_dollar___unused_3_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply3(go__go_3_6_4, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_4
			}), gopurs_runtime.Int(7), gopurs_runtime.Int(3)).IntVal)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_3_5
			_dollar___unused_4_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(z_0_0)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_4_8
			var __t14 gopurs_runtime.Value
			{
				var go__go_5_9_5 gopurs_runtime.Value
				go__go_5_9_5 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return func() gopurs_runtime.Value {
							var v_6_loop gopurs_runtime.Value = v_6_loop_val
							var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
						go__go_5_9_5:
							for {
								if false {
									continue go__go_5_9_5
								}
								var v_6 gopurs_runtime.Value = v_6_loop
								_ = v_6
								var v1_7 gopurs_runtime.Value = v1_7_loop
								_ = v1_7
								var __t10 gopurs_runtime.Value
								{
									if (v1_7.IntVal) == (0) {
										__t10 = v_6
										goto end_branch_10
									} else {

									}
								}
								{
									v_6_loop = gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Int((v1_7.IntVal) + (a_8.IntVal))
									})
									v1_7_loop = gopurs_runtime.Int(0)
									continue go__go_5_9_5
									__t10 = gopurs_runtime.Value{}
								}
							end_branch_10:
								return __t10
							}
						}()
					})
				})
				var __t_and_13 bool = false
				if (gopurs_runtime.Apply3(go__go_5_9_5, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_6
				}), gopurs_runtime.Int(7), gopurs_runtime.Int(3)).IntVal) == (10) {

					var go__go_6_11_6 gopurs_runtime.Value
					go__go_6_11_6 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								var v_7_loop gopurs_runtime.Value = v_7_loop_val
								var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
							go__go_6_11_6:
								for {
									if false {
										continue go__go_6_11_6
									}
									var v_7 gopurs_runtime.Value = v_7_loop
									_ = v_7
									var v1_8 gopurs_runtime.Value = v1_8_loop
									_ = v1_8
									var __t12 gopurs_runtime.Value
									{
										if (v1_8.IntVal) == (0) {
											__t12 = v_7
											goto end_branch_12
										} else {

										}
									}
									{
										v_7_loop = gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Int((v1_8.IntVal) + (a_9.IntVal))
										})
										v1_8_loop = gopurs_runtime.Int(0)
										continue go__go_6_11_6
										__t12 = gopurs_runtime.Value{}
									}
								end_branch_12:
									return __t12
								}
							}()
						})
					})
					__t_and_13 = ((gopurs_runtime.Apply3(go__go_6_11_6, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return x_7
					}), gopurs_runtime.Int(7), gopurs_runtime.Int(3)).IntVal) == (10)) && ((z_0_0) == (1))
				}
				if __t_and_13 {
					__t14 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_14
				} else {

				}
			}
			{
				__t14 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail"))
			}
		end_branch_14:
			return gopurs_runtime.Apply(__t14, gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_count__3600013392 gopurs_runtime.Value
var once_Main_count__3600013392 sync.Once

func Get_Main_count__3600013392() gopurs_runtime.Value {
	once_Main_count__3600013392.Do(func() {
		cache_Main_count__3600013392 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_count__3600013392(p_0_box)
		})
	})
	return cache_Main_count__3600013392
}

func Call_Main_identity(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return gopurs_runtime.Int(x_0).IntVal
}

func Call_Main_count(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var count_prime__1_0_2 gopurs_runtime.Value
	count_prime__1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				var v_2_loop int64 = v_2_loop_val.IntVal
				var v1_3_loop []gopurs_runtime.Value = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v1_3_loop_val.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()
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
						if (gopurs_runtime.Int(int64(len(v1_3))).IntVal) == (0) {
							__t2 = v_2
							goto end_branch_2
						} else {

						}
					}
					{
						var __t1 int64
						{
							if (gopurs_runtime.Apply(p_0, v1_3[0]).IntVal) != (0) {
								__t1 = (v_2) + (1)
								goto end_branch_1
							} else {

							}
						}
						{
							__t1 = (v_2) + (0)
						}
					end_branch_1:
						v_2_loop = __t1
						v1_3_loop = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(int64(len(v1_3))), gopurs_runtime.Array(v1_3)).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
						continue count_prime__1_0_2
						__t2 = gopurs_runtime.Value{}.IntVal
					}
				end_branch_2:
					return gopurs_runtime.Int(__t2)
				}
			}()
		})
	})
	return gopurs_runtime.Apply(count_prime__1_0_2, gopurs_runtime.Int(0))
}

func Call_Main_count__3600013392(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var count_prime__1_0_7 gopurs_runtime.Value
	count_prime__1_0_7 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				var v_2_loop int64 = v_2_loop_val.IntVal
				var v1_3_loop []gopurs_runtime.Value = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v1_3_loop_val.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()
			count_prime__1_0_7:
				for {
					if false {
						continue count_prime__1_0_7
					}
					var v_2 int64 = v_2_loop
					_ = v_2
					var v1_3 []gopurs_runtime.Value = v1_3_loop
					_ = v1_3
					var __t2 int64
					{
						if (gopurs_runtime.Int(int64(len(v1_3))).IntVal) == (0) {
							__t2 = v_2
							goto end_branch_2
						} else {

						}
					}
					{
						var __t1 int64
						{
							if (gopurs_runtime.Apply(p_0, v1_3[0]).IntVal) != (0) {
								__t1 = (v_2) + (1)
								goto end_branch_1
							} else {

							}
						}
						{
							__t1 = (v_2) + (0)
						}
					end_branch_1:
						v_2_loop = __t1
						v1_3_loop = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(int64(len(v1_3))), gopurs_runtime.Array(v1_3)).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
						continue count_prime__1_0_7
						__t2 = gopurs_runtime.Value{}.IntVal
					}
				end_branch_2:
					return gopurs_runtime.Int(__t2)
				}
			}()
		})
	})
	return gopurs_runtime.Apply(count_prime__1_0_7, gopurs_runtime.Int(0))
}
