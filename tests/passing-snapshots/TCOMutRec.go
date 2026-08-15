package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_tco4 gopurs_runtime.Value
var once_Main_tco4 sync.Once

func Get_Main_tco4() gopurs_runtime.Value {
	once_Main_tco4.Do(func() {
		cache_Main_tco4 = func() gopurs_runtime.Value {
			var f_0_0_0 gopurs_runtime.Value
			f_0_0_0 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_1_loop int64 = x_1_loop_val.IntVal
						var y_2_loop int64 = y_2_loop_val.IntVal
					f_0_0_0:
						for {
							if false {
								continue f_0_0_0
							}
							var x_1 int64 = x_1_loop
							_ = x_1
							var y_2 int64 = y_2_loop
							_ = y_2
							var __t1 int64
							{
								if (y_2) <= (0) {
									__t1 = x_1
									goto end_branch_1
								} else {

								}
							}
							{
								x_1_loop = (x_1) + (2)
								y_2_loop = (y_2) - (1)
								continue f_0_0_0
								__t1 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_1:
							return gopurs_runtime.Int(__t1)
						}
					}()
				})
			})
			return gopurs_runtime.Apply(f_0_0_0, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_tco4
}

var cache_Main_tco3 gopurs_runtime.Value
var once_Main_tco3 sync.Once

func Get_Main_tco3() gopurs_runtime.Value {
	once_Main_tco3.Do(func() {
		cache_Main_tco3 = gopurs_runtime.Func(func(y0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_tco3(y0_0_box.IntVal))
		})
	})
	return cache_Main_tco3
}

var cache_Main_tco2 gopurs_runtime.Value
var once_Main_tco2 sync.Once

func Get_Main_tco2() gopurs_runtime.Value {
	once_Main_tco2.Do(func() {
		cache_Main_tco2 = func() gopurs_runtime.Value {
			var f_0_0_3 gopurs_runtime.Value
			f_0_0_3 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_1_loop gopurs_runtime.Value = x_1_loop_val
						var y_2_loop gopurs_runtime.Value = y_2_loop_val
					f_0_0_3:
						for {
							if false {
								continue f_0_0_3
							}
							var x_1 gopurs_runtime.Value = x_1_loop
							_ = x_1
							var y_2 gopurs_runtime.Value = y_2_loop
							_ = y_2
							// TAST (Let): __local_var_3_1 -> int64
							__local_var_3_1 := (x_1.IntVal) + (2)
							_ = __local_var_3_1
							// TAST (Let): __local_var_4_2 -> int64
							__local_var_4_2 := (y_2.IntVal) - (1)
							_ = __local_var_4_2
							var __t3 int64
							{
								if (__local_var_4_2) <= (0) {
									__t3 = __local_var_3_1
									goto end_branch_3
								} else {

								}
							}
							{
								x_1_loop = gopurs_runtime.Int(__local_var_3_1)
								y_2_loop = gopurs_runtime.Int(__local_var_4_2)
								continue f_0_0_3
								__t3 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_3:
							return gopurs_runtime.Int(__t3)
						}
					}()
				})
			})
			return gopurs_runtime.Apply(f_0_0_3, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_tco2
}

var cache_Main_tco1 gopurs_runtime.Value
var once_Main_tco1 sync.Once

func Get_Main_tco1() gopurs_runtime.Value {
	once_Main_tco1.Do(func() {
		cache_Main_tco1 = func() gopurs_runtime.Value {
			var f_0_0_4 gopurs_runtime.Value
			f_0_0_4 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_1_loop gopurs_runtime.Value = x_1_loop_val
						var y_2_loop gopurs_runtime.Value = y_2_loop_val
					f_0_0_4:
						for {
							if false {
								continue f_0_0_4
							}
							var x_1 gopurs_runtime.Value = x_1_loop
							_ = x_1
							var y_2 gopurs_runtime.Value = y_2_loop
							_ = y_2
							// TAST (Let): __local_var_3_1 -> int64
							__local_var_3_1 := (x_1.IntVal) + (2)
							_ = __local_var_3_1
							// TAST (Let): __local_var_4_2 -> int64
							__local_var_4_2 := (y_2.IntVal) - (1)
							_ = __local_var_4_2
							var __t3 int64
							{
								if (__local_var_4_2) <= (0) {
									__t3 = __local_var_3_1
									goto end_branch_3
								} else {

								}
							}
							{
								x_1_loop = gopurs_runtime.Int(__local_var_3_1)
								y_2_loop = gopurs_runtime.Int(__local_var_4_2)
								continue f_0_0_4
								__t3 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_3:
							return gopurs_runtime.Int(__t3)
						}
					}()
				})
			})
			return gopurs_runtime.Apply(f_0_0_4, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_tco1
}

var cache_Main_ntco4 gopurs_runtime.Value
var once_Main_ntco4 sync.Once

func Get_Main_ntco4() gopurs_runtime.Value {
	once_Main_ntco4.Do(func() {
		cache_Main_ntco4 = func() gopurs_runtime.Value {
			var f_0_0_5 gopurs_runtime.Value
			f_0_0_5 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_1_loop int64 = x_1_loop_val.IntVal
						var y_2_loop int64 = y_2_loop_val.IntVal
					f_0_0_5:
						for {
							if false {
								continue f_0_0_5
							}
							var x_1 int64 = x_1_loop
							_ = x_1
							var y_2 int64 = y_2_loop
							_ = y_2
							var __t1 int64
							{
								if (y_2) <= (0) {
									__t1 = x_1
									goto end_branch_1
								} else {

								}
							}
							{
								x_1_loop = (x_1) + (2)
								y_2_loop = (y_2) - (1)
								continue f_0_0_5
								__t1 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_1:
							return gopurs_runtime.Int(__t1)
						}
					}()
				})
			})
			return gopurs_runtime.Apply(f_0_0_5, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_ntco4
}

var cache_Main_ntco3 gopurs_runtime.Value
var once_Main_ntco3 sync.Once

func Get_Main_ntco3() gopurs_runtime.Value {
	once_Main_ntco3.Do(func() {
		cache_Main_ntco3 = func() gopurs_runtime.Value {
			var f_0_0_6 gopurs_runtime.Value
			_ = f_0_0_6
			f_0_0_6 = gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): g_3_1 -> gopurs_runtime.Value
					g_3_1 := gopurs_runtime.Apply(f_0_0_6, gopurs_runtime.Int((x_1.IntVal)+(2)))
					_ = g_3_1
					var __t2 int64
					{
						if (y_2.IntVal) <= (0) {
							__t2 = x_1.IntVal
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Apply(g_3_1, gopurs_runtime.Int((y_2.IntVal)-(1))).IntVal
					}
				end_branch_2:
					return gopurs_runtime.Int(__t2)
				})
			})
			return gopurs_runtime.Apply(f_0_0_6, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_ntco3
}

var cache_Main_ntco2 gopurs_runtime.Value
var once_Main_ntco2 sync.Once

func Get_Main_ntco2() gopurs_runtime.Value {
	once_Main_ntco2.Do(func() {
		cache_Main_ntco2 = func() gopurs_runtime.Value {
			var f_0_0_7 gopurs_runtime.Value
			f_0_0_7 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_1_loop int64 = x_1_loop_val.IntVal
						var y_2_loop int64 = y_2_loop_val.IntVal
					f_0_0_7:
						for {
							if false {
								continue f_0_0_7
							}
							var x_1 int64 = x_1_loop
							_ = x_1
							var y_2 int64 = y_2_loop
							_ = y_2
							var __t1 int64
							{
								if (y_2) <= (0) {
									__t1 = x_1
									goto end_branch_1
								} else {

								}
							}
							{
								x_1_loop = (x_1) + (2)
								y_2_loop = (y_2) - (1)
								continue f_0_0_7
								__t1 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_1:
							return gopurs_runtime.Int(__t1)
						}
					}()
				})
			})
			return gopurs_runtime.Apply(f_0_0_7, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_ntco2
}

var cache_Main_ntco1 gopurs_runtime.Value
var once_Main_ntco1 sync.Once

func Get_Main_ntco1() gopurs_runtime.Value {
	once_Main_ntco1.Do(func() {
		cache_Main_ntco1 = gopurs_runtime.Func(func(y0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_ntco1(y0_0_box.IntVal))
		})
	})
	return cache_Main_ntco1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var f_0_1_9 gopurs_runtime.Value
			f_0_1_9 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_1_loop gopurs_runtime.Value = x_1_loop_val
						var y_2_loop gopurs_runtime.Value = y_2_loop_val
					f_0_1_9:
						for {
							if false {
								continue f_0_1_9
							}
							var x_1 gopurs_runtime.Value = x_1_loop
							_ = x_1
							var y_2 gopurs_runtime.Value = y_2_loop
							_ = y_2
							// TAST (Let): __local_var_3_2 -> int64
							__local_var_3_2 := (x_1.IntVal) + (2)
							_ = __local_var_3_2
							// TAST (Let): __local_var_4_3 -> int64
							__local_var_4_3 := (y_2.IntVal) - (1)
							_ = __local_var_4_3
							var __t4 int64
							{
								if (__local_var_4_3) <= (0) {
									__t4 = __local_var_3_2
									goto end_branch_4
								} else {

								}
							}
							{
								x_1_loop = gopurs_runtime.Int(__local_var_3_2)
								y_2_loop = gopurs_runtime.Int(__local_var_4_3)
								continue f_0_1_9
								__t4 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_4:
							return gopurs_runtime.Int(__t4)
						}
					}()
				})
			})
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_0_1_9, gopurs_runtime.Int(0), gopurs_runtime.Int(100000)).IntVal), gopurs_runtime.Int(200000))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_5 -> gopurs_runtime.Value
			__local_var_1_5 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): result_1_6 -> bool
				result_1_6 := (gopurs_runtime.RecordGet(__local_var_0_0, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_0_0, "expected").IntVal)
				_ = result_1_6
				// TAST (Let): message_2_7 -> string
				message_2_7 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "actual")).StrVal())
				_ = message_2_7
				// TAST (Let): __local_var_3_8 -> gopurs_runtime.Value
				__local_var_3_8 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_2_7))
				_ = __local_var_3_8
				var __t10 gopurs_runtime.Value
				{
					if (result_1_6) != (true) {
						__t10 = __local_var_3_8
						goto end_branch_10
					} else {

					}
				}
				{
					if result_1_6 {
						__t10 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							return Get_Data_Unit_unit()
						})
						goto end_branch_10
					} else {

					}
				}
				{
					__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_10:
				// TAST (Let): __local_var_4_9 -> gopurs_runtime.Value
				__local_var_4_9 := __t10
				_ = __local_var_4_9
				_dollar___unused_5_11 := gopurs_runtime.Apply(__local_var_4_9, gopurs_runtime.Value{})
				_ = _dollar___unused_5_11
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_2_7), gopurs_runtime.Bool(result_1_6)), gopurs_runtime.Value{})
			})
			_ = __local_var_1_5
			_dollar___unused_2_12 := gopurs_runtime.Apply(__local_var_1_5, gopurs_runtime.Value{})
			_ = _dollar___unused_2_12
			var f_3_14_10 gopurs_runtime.Value
			f_3_14_10 = gopurs_runtime.Func(func(x_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_4_loop gopurs_runtime.Value = x_4_loop_val
						var y_5_loop gopurs_runtime.Value = y_5_loop_val
					f_3_14_10:
						for {
							if false {
								continue f_3_14_10
							}
							var x_4 gopurs_runtime.Value = x_4_loop
							_ = x_4
							var y_5 gopurs_runtime.Value = y_5_loop
							_ = y_5
							// TAST (Let): __local_var_6_15 -> int64
							__local_var_6_15 := (x_4.IntVal) + (2)
							_ = __local_var_6_15
							// TAST (Let): __local_var_7_16 -> int64
							__local_var_7_16 := (y_5.IntVal) - (1)
							_ = __local_var_7_16
							var __t17 int64
							{
								if (__local_var_7_16) <= (0) {
									__t17 = __local_var_6_15
									goto end_branch_17
								} else {

								}
							}
							{
								x_4_loop = gopurs_runtime.Int(__local_var_6_15)
								y_5_loop = gopurs_runtime.Int(__local_var_7_16)
								continue f_3_14_10
								__t17 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_17:
							return gopurs_runtime.Int(__t17)
						}
					}()
				})
			})
			// TAST (Let): __local_var_3_13 -> gopurs_runtime.Value
			__local_var_3_13 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_3_14_10, gopurs_runtime.Int(0), gopurs_runtime.Int(100000)).IntVal), gopurs_runtime.Int(200000))
			_ = __local_var_3_13
			// TAST (Let): result_4_19 -> bool
			result_4_19 := (gopurs_runtime.RecordGet(__local_var_3_13, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_3_13, "expected").IntVal)
			_ = result_4_19
			// TAST (Let): message_5_20 -> string
			message_5_20 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_3_13, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_3_13, "actual")).StrVal())
			_ = message_5_20
			// TAST (Let): __local_var_6_21 -> gopurs_runtime.Value
			__local_var_6_21 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_5_20))
			_ = __local_var_6_21
			var __t23 gopurs_runtime.Value
			{
				if (result_4_19) != (true) {
					__t23 = __local_var_6_21
					goto end_branch_23
				} else {

				}
			}
			{
				if result_4_19 {
					__t23 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_23
				} else {

				}
			}
			{
				__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_23:
			_dollar___unused_7_22 := gopurs_runtime.Apply(__t23, gopurs_runtime.Value{})
			_ = _dollar___unused_7_22
			_dollar___unused_4_18 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_5_20), gopurs_runtime.Bool(result_4_19)), gopurs_runtime.Value{})
			_ = _dollar___unused_4_18
			var f_5_25_11 gopurs_runtime.Value
			f_5_25_11 = gopurs_runtime.Func(func(x_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_6_loop int64 = x_6_loop_val.IntVal
						var y_7_loop gopurs_runtime.Value = y_7_loop_val
					f_5_25_11:
						for {
							if false {
								continue f_5_25_11
							}
							var x_6 int64 = x_6_loop
							_ = x_6
							var y_7 gopurs_runtime.Value = y_7_loop
							_ = y_7
							var g_8_26_12 gopurs_runtime.Value
							g_8_26_12 = gopurs_runtime.Func(func(x_prime__9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(y_prime__10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return func() gopurs_runtime.Value {
										var x_prime__9_loop gopurs_runtime.Value = x_prime__9_loop_val
										var y_prime__10_loop gopurs_runtime.Value = y_prime__10_loop_val
									g_8_26_12:
										for {
											if false {
												continue g_8_26_12
											}
											var x_prime__9 gopurs_runtime.Value = x_prime__9_loop
											_ = x_prime__9
											var y_prime__10 gopurs_runtime.Value = y_prime__10_loop
											_ = y_prime__10
											var __t28 int64
											{
												if (y_prime__10.IntVal) <= (0) {
													__t28 = x_prime__9.IntVal
													goto end_branch_28
												} else {

												}
											}
											{
												var __t27 int64
												{
													if (y_prime__10.IntVal) > (50000) {
														x_prime__9_loop = gopurs_runtime.Int((x_prime__9.IntVal) + (3))
														y_prime__10_loop = gopurs_runtime.Int((y_prime__10.IntVal) - (1))
														continue g_8_26_12
														__t27 = gopurs_runtime.Value{}.IntVal
														goto end_branch_27
													} else {

													}
												}
												{
													__t27 = gopurs_runtime.Apply2(f_5_25_11, gopurs_runtime.Int((x_prime__9.IntVal)+(2)), gopurs_runtime.Int(y_prime__10.IntVal)).IntVal
												}
											end_branch_27:
												__t28 = __t27
											}
										end_branch_28:
											return gopurs_runtime.Int(__t28)
										}
									}()
								})
							})
							return gopurs_runtime.Int(gopurs_runtime.Apply2(g_8_26_12, gopurs_runtime.Int(x_6), gopurs_runtime.Int((y_7.IntVal)-(1))).IntVal)
						}
					}()
				})
			})
			// TAST (Let): __local_var_5_24 -> gopurs_runtime.Value
			__local_var_5_24 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_5_25_11, gopurs_runtime.Int(0), gopurs_runtime.Int(100000)).IntVal), gopurs_runtime.Int(249997))
			_ = __local_var_5_24
			// TAST (Let): result_6_30 -> bool
			result_6_30 := (gopurs_runtime.RecordGet(__local_var_5_24, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_5_24, "expected").IntVal)
			_ = result_6_30
			// TAST (Let): message_7_31 -> string
			message_7_31 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_5_24, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_5_24, "actual")).StrVal())
			_ = message_7_31
			// TAST (Let): __local_var_8_32 -> gopurs_runtime.Value
			__local_var_8_32 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_7_31))
			_ = __local_var_8_32
			var __t34 gopurs_runtime.Value
			{
				if (result_6_30) != (true) {
					__t34 = __local_var_8_32
					goto end_branch_34
				} else {

				}
			}
			{
				if result_6_30 {
					__t34 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_34
				} else {

				}
			}
			{
				__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_34:
			_dollar___unused_9_33 := gopurs_runtime.Apply(__t34, gopurs_runtime.Value{})
			_ = _dollar___unused_9_33
			_dollar___unused_6_29 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_7_31), gopurs_runtime.Bool(result_6_30)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_29
			var f_7_36_13 gopurs_runtime.Value
			f_7_36_13 = gopurs_runtime.Func(func(x_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_8_loop int64 = x_8_loop_val.IntVal
						var y_9_loop int64 = y_9_loop_val.IntVal
					f_7_36_13:
						for {
							if false {
								continue f_7_36_13
							}
							var x_8 int64 = x_8_loop
							_ = x_8
							var y_9 int64 = y_9_loop
							_ = y_9
							var __t37 int64
							{
								if (y_9) <= (0) {
									__t37 = x_8
									goto end_branch_37
								} else {

								}
							}
							{
								x_8_loop = (x_8) + (2)
								y_9_loop = (y_9) - (1)
								continue f_7_36_13
								__t37 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_37:
							return gopurs_runtime.Int(__t37)
						}
					}()
				})
			})
			// TAST (Let): __local_var_7_35 -> gopurs_runtime.Value
			__local_var_7_35 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_7_36_13, gopurs_runtime.Int(0), gopurs_runtime.Int(100000)).IntVal), gopurs_runtime.Int(200000))
			_ = __local_var_7_35
			// TAST (Let): result_8_39 -> bool
			result_8_39 := (gopurs_runtime.RecordGet(__local_var_7_35, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_7_35, "expected").IntVal)
			_ = result_8_39
			// TAST (Let): message_9_40 -> string
			message_9_40 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_7_35, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_7_35, "actual")).StrVal())
			_ = message_9_40
			// TAST (Let): __local_var_10_41 -> gopurs_runtime.Value
			__local_var_10_41 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_9_40))
			_ = __local_var_10_41
			var __t43 gopurs_runtime.Value
			{
				if (result_8_39) != (true) {
					__t43 = __local_var_10_41
					goto end_branch_43
				} else {

				}
			}
			{
				if result_8_39 {
					__t43 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_43
				} else {

				}
			}
			{
				__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_43:
			_dollar___unused_11_42 := gopurs_runtime.Apply(__t43, gopurs_runtime.Value{})
			_ = _dollar___unused_11_42
			_dollar___unused_8_38 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_9_40), gopurs_runtime.Bool(result_8_39)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_38
			var f_9_45_14 gopurs_runtime.Value
			_ = f_9_45_14
			f_9_45_14 = gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t46 gopurs_runtime.Value
				{
					if (x_10.IntVal) > (1000) {
						__t46 = gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((x_10.IntVal) + (v_11.IntVal))
						})
						goto end_branch_46
					} else {

					}
				}
				{
					__t46 = gopurs_runtime.Func(func(y_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.Apply2(f_9_45_14, gopurs_runtime.Int((x_10.IntVal)+(10)), gopurs_runtime.Int((y_prime__11.IntVal)-(1))).IntVal)
					})
				}
			end_branch_46:
				return __t46
			})
			// TAST (Let): __local_var_9_44 -> gopurs_runtime.Value
			__local_var_9_44 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_9_45_14, gopurs_runtime.Int(0), gopurs_runtime.Int(100)).IntVal), gopurs_runtime.Int(1009))
			_ = __local_var_9_44
			// TAST (Let): result_10_48 -> bool
			result_10_48 := (gopurs_runtime.RecordGet(__local_var_9_44, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_9_44, "expected").IntVal)
			_ = result_10_48
			// TAST (Let): message_11_49 -> string
			message_11_49 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_9_44, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_9_44, "actual")).StrVal())
			_ = message_11_49
			// TAST (Let): __local_var_12_50 -> gopurs_runtime.Value
			__local_var_12_50 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_11_49))
			_ = __local_var_12_50
			var __t52 gopurs_runtime.Value
			{
				if (result_10_48) != (true) {
					__t52 = __local_var_12_50
					goto end_branch_52
				} else {

				}
			}
			{
				if result_10_48 {
					__t52 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_52
				} else {

				}
			}
			{
				__t52 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_52:
			_dollar___unused_13_51 := gopurs_runtime.Apply(__t52, gopurs_runtime.Value{})
			_ = _dollar___unused_13_51
			_dollar___unused_10_47 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_11_49), gopurs_runtime.Bool(result_10_48)), gopurs_runtime.Value{})
			_ = _dollar___unused_10_47
			var f_11_54_15 gopurs_runtime.Value
			f_11_54_15 = gopurs_runtime.Func(func(x_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_12_loop int64 = x_12_loop_val.IntVal
						var y_13_loop int64 = y_13_loop_val.IntVal
					f_11_54_15:
						for {
							if false {
								continue f_11_54_15
							}
							var x_12 int64 = x_12_loop
							_ = x_12
							var y_13 int64 = y_13_loop
							_ = y_13
							var __t55 int64
							{
								if (y_13) <= (0) {
									__t55 = x_12
									goto end_branch_55
								} else {

								}
							}
							{
								x_12_loop = (x_12) + (2)
								y_13_loop = (y_13) - (1)
								continue f_11_54_15
								__t55 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_55:
							return gopurs_runtime.Int(__t55)
						}
					}()
				})
			})
			// TAST (Let): __local_var_11_53 -> gopurs_runtime.Value
			__local_var_11_53 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_11_54_15, gopurs_runtime.Int(0), gopurs_runtime.Int(100)).IntVal), gopurs_runtime.Int(200))
			_ = __local_var_11_53
			// TAST (Let): result_12_57 -> bool
			result_12_57 := (gopurs_runtime.RecordGet(__local_var_11_53, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_11_53, "expected").IntVal)
			_ = result_12_57
			// TAST (Let): message_13_58 -> string
			message_13_58 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_11_53, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_11_53, "actual")).StrVal())
			_ = message_13_58
			// TAST (Let): __local_var_14_59 -> gopurs_runtime.Value
			__local_var_14_59 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_13_58))
			_ = __local_var_14_59
			var __t61 gopurs_runtime.Value
			{
				if (result_12_57) != (true) {
					__t61 = __local_var_14_59
					goto end_branch_61
				} else {

				}
			}
			{
				if result_12_57 {
					__t61 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_61
				} else {

				}
			}
			{
				__t61 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_61:
			_dollar___unused_15_60 := gopurs_runtime.Apply(__t61, gopurs_runtime.Value{})
			_ = _dollar___unused_15_60
			_dollar___unused_12_56 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_13_58), gopurs_runtime.Bool(result_12_57)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_56
			var f_13_63_16 gopurs_runtime.Value
			_ = f_13_63_16
			f_13_63_16 = gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_15 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): g_16_64 -> gopurs_runtime.Value
					g_16_64 := gopurs_runtime.Apply(f_13_63_16, gopurs_runtime.Int((x_14.IntVal)+(2)))
					_ = g_16_64
					var __t65 int64
					{
						if (y_15.IntVal) <= (0) {
							__t65 = x_14.IntVal
							goto end_branch_65
						} else {

						}
					}
					{
						__t65 = gopurs_runtime.Apply(g_16_64, gopurs_runtime.Int((y_15.IntVal)-(1))).IntVal
					}
				end_branch_65:
					return gopurs_runtime.Int(__t65)
				})
			})
			// TAST (Let): __local_var_13_62 -> gopurs_runtime.Value
			__local_var_13_62 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_13_63_16, gopurs_runtime.Int(0), gopurs_runtime.Int(100)).IntVal), gopurs_runtime.Int(200))
			_ = __local_var_13_62
			// TAST (Let): result_14_67 -> bool
			result_14_67 := (gopurs_runtime.RecordGet(__local_var_13_62, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_13_62, "expected").IntVal)
			_ = result_14_67
			// TAST (Let): message_15_68 -> string
			message_15_68 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_13_62, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_13_62, "actual")).StrVal())
			_ = message_15_68
			// TAST (Let): __local_var_16_69 -> gopurs_runtime.Value
			__local_var_16_69 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_15_68))
			_ = __local_var_16_69
			var __t71 gopurs_runtime.Value
			{
				if (result_14_67) != (true) {
					__t71 = __local_var_16_69
					goto end_branch_71
				} else {

				}
			}
			{
				if result_14_67 {
					__t71 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_71
				} else {

				}
			}
			{
				__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_71:
			_dollar___unused_17_70 := gopurs_runtime.Apply(__t71, gopurs_runtime.Value{})
			_ = _dollar___unused_17_70
			_dollar___unused_14_66 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_15_68), gopurs_runtime.Bool(result_14_67)), gopurs_runtime.Value{})
			_ = _dollar___unused_14_66
			var f_15_73_17 gopurs_runtime.Value
			f_15_73_17 = gopurs_runtime.Func(func(x_16_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_17_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var x_16_loop int64 = x_16_loop_val.IntVal
						var y_17_loop int64 = y_17_loop_val.IntVal
					f_15_73_17:
						for {
							if false {
								continue f_15_73_17
							}
							var x_16 int64 = x_16_loop
							_ = x_16
							var y_17 int64 = y_17_loop
							_ = y_17
							var __t74 int64
							{
								if (y_17) <= (0) {
									__t74 = x_16
									goto end_branch_74
								} else {

								}
							}
							{
								x_16_loop = (x_16) + (2)
								y_17_loop = (y_17) - (1)
								continue f_15_73_17
								__t74 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_74:
							return gopurs_runtime.Int(__t74)
						}
					}()
				})
			})
			// TAST (Let): __local_var_15_72 -> gopurs_runtime.Value
			__local_var_15_72 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_15_73_17, gopurs_runtime.Int(0), gopurs_runtime.Int(100)).IntVal), gopurs_runtime.Int(200))
			_ = __local_var_15_72
			// TAST (Let): result_16_76 -> bool
			result_16_76 := (gopurs_runtime.RecordGet(__local_var_15_72, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_15_72, "expected").IntVal)
			_ = result_16_76
			// TAST (Let): message_17_77 -> string
			message_17_77 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_15_72, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_15_72, "actual")).StrVal())
			_ = message_17_77
			// TAST (Let): __local_var_18_78 -> gopurs_runtime.Value
			__local_var_18_78 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_17_77))
			_ = __local_var_18_78
			var __t80 gopurs_runtime.Value
			{
				if (result_16_76) != (true) {
					__t80 = __local_var_18_78
					goto end_branch_80
				} else {

				}
			}
			{
				if result_16_76 {
					__t80 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_80
				} else {

				}
			}
			{
				__t80 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_80:
			_dollar___unused_19_79 := gopurs_runtime.Apply(__t80, gopurs_runtime.Value{})
			_ = _dollar___unused_19_79
			_dollar___unused_16_75 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_17_77), gopurs_runtime.Bool(result_16_76)), gopurs_runtime.Value{})
			_ = _dollar___unused_16_75
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_tco3(y0_0_loop int64) int64 {
	var y0_0 int64 = y0_0_loop
	_ = y0_0
	var f_1_0_1 gopurs_runtime.Value
	f_1_0_1 = gopurs_runtime.Func(func(x_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				var x_2_loop int64 = x_2_loop_val.IntVal
				var y_3_loop gopurs_runtime.Value = y_3_loop_val
			f_1_0_1:
				for {
					if false {
						continue f_1_0_1
					}
					var x_2 int64 = x_2_loop
					_ = x_2
					var y_3 gopurs_runtime.Value = y_3_loop
					_ = y_3
					var g_4_1_2 gopurs_runtime.Value
					g_4_1_2 = gopurs_runtime.Func(func(x_prime__5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(y_prime__6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								var x_prime__5_loop gopurs_runtime.Value = x_prime__5_loop_val
								var y_prime__6_loop gopurs_runtime.Value = y_prime__6_loop_val
							g_4_1_2:
								for {
									if false {
										continue g_4_1_2
									}
									var x_prime__5 gopurs_runtime.Value = x_prime__5_loop
									_ = x_prime__5
									var y_prime__6 gopurs_runtime.Value = y_prime__6_loop
									_ = y_prime__6
									var __t3 int64
									{
										if (y_prime__6.IntVal) <= (0) {
											__t3 = x_prime__5.IntVal
											goto end_branch_3
										} else {

										}
									}
									{
										var __t2 int64
										{
											if (y_prime__6.IntVal) > ((y0_0) / (2)) {
												x_prime__5_loop = gopurs_runtime.Int((x_prime__5.IntVal) + (3))
												y_prime__6_loop = gopurs_runtime.Int((y_prime__6.IntVal) - (1))
												continue g_4_1_2
												__t2 = gopurs_runtime.Value{}.IntVal
												goto end_branch_2
											} else {

											}
										}
										{
											__t2 = gopurs_runtime.Apply2(f_1_0_1, gopurs_runtime.Int((x_prime__5.IntVal)+(2)), gopurs_runtime.Int(y_prime__6.IntVal)).IntVal
										}
									end_branch_2:
										__t3 = __t2
									}
								end_branch_3:
									return gopurs_runtime.Int(__t3)
								}
							}()
						})
					})
					return gopurs_runtime.Int(gopurs_runtime.Apply2(g_4_1_2, gopurs_runtime.Int(x_2), gopurs_runtime.Int((y_3.IntVal)-(1))).IntVal)
				}
			}()
		})
	})
	return gopurs_runtime.Apply2(f_1_0_1, gopurs_runtime.Int(0), gopurs_runtime.Int(y0_0)).IntVal
}

func Call_Main_ntco1(y0_0_loop int64) int64 {
	var y0_0 int64 = y0_0_loop
	_ = y0_0
	var f_1_0_8 gopurs_runtime.Value
	_ = f_1_0_8
	f_1_0_8 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (x_2.IntVal) > ((10) * (y0_0)) {
				__t1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((x_2.IntVal) + (v_3.IntVal))
				})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Func(func(y_prime__3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply2(f_1_0_8, gopurs_runtime.Int((x_2.IntVal)+(10)), gopurs_runtime.Int((y_prime__3.IntVal)-(1))).IntVal)
			})
		}
	end_branch_1:
		return __t1
	})
	return gopurs_runtime.Apply2(f_1_0_8, gopurs_runtime.Int(0), gopurs_runtime.Int(y0_0)).IntVal
}
