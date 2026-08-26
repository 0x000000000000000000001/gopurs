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
			var Call_local_Main_f_0_0_0 func(int64, int64) gopurs_runtime.Value
			_ = Call_local_Main_f_0_0_0
			var f_0_0_0 gopurs_runtime.Value
			_ = f_0_0_0
			Call_local_Main_f_0_0_0 = func(x_1_loop int64, y_2_loop int64) gopurs_runtime.Value {
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
			}
			f_0_0_0 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_0_0_0(x_1_loop_val.IntVal, y_2_loop_val.IntVal)
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
			var Call_local_Main_f_0_0_3 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_f_0_0_3
			var f_0_0_3 gopurs_runtime.Value
			_ = f_0_0_3
			Call_local_Main_f_0_0_3 = func(x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
			f_0_0_3:
				for {
					if false {
						continue f_0_0_3
					}
					var x_1 gopurs_runtime.Value = x_1_loop
					_ = x_1
					var y_2 gopurs_runtime.Value = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_1 shape=Other bindingType=Int
					__local_var_3_1 := (x_1.IntVal) + (2)
					_ = __local_var_3_1
					// TAST (Let): __local_var_4_2 shape=Other bindingType=Int
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
			}
			f_0_0_3 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_0_0_3(x_1_loop_val, y_2_loop_val)
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
			var Call_local_Main_f_0_0_4 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_f_0_0_4
			var f_0_0_4 gopurs_runtime.Value
			_ = f_0_0_4
			Call_local_Main_f_0_0_4 = func(x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
			f_0_0_4:
				for {
					if false {
						continue f_0_0_4
					}
					var x_1 gopurs_runtime.Value = x_1_loop
					_ = x_1
					var y_2 gopurs_runtime.Value = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_1 shape=Other bindingType=Int
					__local_var_3_1 := (x_1.IntVal) + (2)
					_ = __local_var_3_1
					// TAST (Let): __local_var_4_2 shape=Other bindingType=Int
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
			}
			f_0_0_4 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_0_0_4(x_1_loop_val, y_2_loop_val)
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
			var Call_local_Main_f_0_0_5 func(int64, int64) gopurs_runtime.Value
			_ = Call_local_Main_f_0_0_5
			var f_0_0_5 gopurs_runtime.Value
			_ = f_0_0_5
			Call_local_Main_f_0_0_5 = func(x_1_loop int64, y_2_loop int64) gopurs_runtime.Value {
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
			}
			f_0_0_5 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_0_0_5(x_1_loop_val.IntVal, y_2_loop_val.IntVal)
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
					// TAST (Let): g_3_1 shape=App(Other) bindingType=(Func [Int] Int)
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
			var Call_local_Main_f_0_0_7 func(int64, int64) gopurs_runtime.Value
			_ = Call_local_Main_f_0_0_7
			var f_0_0_7 gopurs_runtime.Value
			_ = f_0_0_7
			Call_local_Main_f_0_0_7 = func(x_1_loop int64, y_2_loop int64) gopurs_runtime.Value {
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
			}
			f_0_0_7 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_0_0_7(x_1_loop_val.IntVal, y_2_loop_val.IntVal)
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
			var Call_local_Main_f_0_1_9 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_f_0_1_9
			var f_0_1_9 gopurs_runtime.Value
			_ = f_0_1_9
			Call_local_Main_f_0_1_9 = func(x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
			f_0_1_9:
				for {
					if false {
						continue f_0_1_9
					}
					var x_1 gopurs_runtime.Value = x_1_loop
					_ = x_1
					var y_2 gopurs_runtime.Value = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_2 shape=Other bindingType=Int
					__local_var_3_2 := (x_1.IntVal) + (2)
					_ = __local_var_3_2
					// TAST (Let): __local_var_4_3 shape=Other bindingType=Int
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
			}
			f_0_1_9 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_0_1_9(x_1_loop_val, y_2_loop_val)
				})
			})
			// TAST (Let): __local_var_0_0 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_0_0 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_local_Main_f_0_1_9(gopurs_runtime.Int(0), gopurs_runtime.Int(100000)).IntVal), gopurs_runtime.Int(200000))
			_ = __local_var_0_0
			// TAST (Let): result_1_5 shape=Other bindingType=Boolean
			result_1_5 := (gopurs_runtime.RecordGet(__local_var_0_0, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_0_0, "expected").IntVal)
			_ = result_1_5
			// TAST (Let): message_2_6 shape=Other bindingType=String
			message_2_6 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_0_0, "actual")).StrVal())
			_ = message_2_6
			// TAST (Let): __local_var_3_7 shape=Let(Let(EffectBind(App(Var)))) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_3_7 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_8 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
				__local_var_3_8 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_2_6))
				_ = __local_var_3_8
				var __t10 gopurs_runtime.Value
				{
					if (result_1_5) != (true) {
						__t10 = __local_var_3_8
						goto end_branch_10
					} else {

					}
				}
				{
					if result_1_5 {
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
				// TAST (Let): __local_var_4_9 shape=Branch(Other, EffectPure, def=Other) bindingType=(TypeApp (TypeVar m) [Unit])
				__local_var_4_9 := __t10
				_ = __local_var_4_9
				_dollar___unused_5_11 := gopurs_runtime.Apply(__local_var_4_9, gopurs_runtime.Value{})
				_ = _dollar___unused_5_11
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_2_6), gopurs_runtime.Bool(result_1_5)), gopurs_runtime.Value{})
			})
			_ = __local_var_3_7
			_dollar___unused_4_12 := gopurs_runtime.Apply(__local_var_3_7, gopurs_runtime.Value{})
			_ = _dollar___unused_4_12
			var Call_local_Main_f_5_14_10 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_f_5_14_10
			var f_5_14_10 gopurs_runtime.Value
			_ = f_5_14_10
			Call_local_Main_f_5_14_10 = func(x_6_loop gopurs_runtime.Value, y_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
			f_5_14_10:
				for {
					if false {
						continue f_5_14_10
					}
					var x_6 gopurs_runtime.Value = x_6_loop
					_ = x_6
					var y_7 gopurs_runtime.Value = y_7_loop
					_ = y_7
					// TAST (Let): __local_var_8_15 shape=Other bindingType=Int
					__local_var_8_15 := (x_6.IntVal) + (2)
					_ = __local_var_8_15
					// TAST (Let): __local_var_9_16 shape=Other bindingType=Int
					__local_var_9_16 := (y_7.IntVal) - (1)
					_ = __local_var_9_16
					var __t17 int64
					{
						if (__local_var_9_16) <= (0) {
							__t17 = __local_var_8_15
							goto end_branch_17
						} else {

						}
					}
					{
						x_6_loop = gopurs_runtime.Int(__local_var_8_15)
						y_7_loop = gopurs_runtime.Int(__local_var_9_16)
						continue f_5_14_10
						__t17 = gopurs_runtime.Value{}.IntVal
					}
				end_branch_17:
					return gopurs_runtime.Int(__t17)
				}
			}
			f_5_14_10 = gopurs_runtime.Func(func(x_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_5_14_10(x_6_loop_val, y_7_loop_val)
				})
			})
			// TAST (Let): __local_var_5_13 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_5_13 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_local_Main_f_5_14_10(gopurs_runtime.Int(0), gopurs_runtime.Int(100000)).IntVal), gopurs_runtime.Int(200000))
			_ = __local_var_5_13
			// TAST (Let): result_6_18 shape=Other bindingType=Boolean
			result_6_18 := (gopurs_runtime.RecordGet(__local_var_5_13, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_5_13, "expected").IntVal)
			_ = result_6_18
			// TAST (Let): message_7_19 shape=Other bindingType=String
			message_7_19 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_5_13, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_5_13, "actual")).StrVal())
			_ = message_7_19
			// TAST (Let): __local_var_8_21 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_8_21 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_7_19))
			_ = __local_var_8_21
			var __t23 gopurs_runtime.Value
			{
				if (result_6_18) != (true) {
					__t23 = __local_var_8_21
					goto end_branch_23
				} else {

				}
			}
			{
				if result_6_18 {
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
			_dollar___unused_9_22 := gopurs_runtime.Apply(__t23, gopurs_runtime.Value{})
			_ = _dollar___unused_9_22
			_dollar___unused_8_20 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_7_19), gopurs_runtime.Bool(result_6_18)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_20
			var Call_local_Main_f_9_25_11 func(int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_f_9_25_11
			var f_9_25_11 gopurs_runtime.Value
			_ = f_9_25_11
			Call_local_Main_f_9_25_11 = func(x_10_loop int64, y_11_loop gopurs_runtime.Value) gopurs_runtime.Value {
			f_9_25_11:
				for {
					if false {
						continue f_9_25_11
					}
					var x_10 int64 = x_10_loop
					_ = x_10
					var y_11 gopurs_runtime.Value = y_11_loop
					_ = y_11
					var Call_local_Main_g_12_26_12 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
					_ = Call_local_Main_g_12_26_12
					var g_12_26_12 gopurs_runtime.Value
					_ = g_12_26_12
					Call_local_Main_g_12_26_12 = func(x_prime__13_loop gopurs_runtime.Value, y_prime__14_loop gopurs_runtime.Value) gopurs_runtime.Value {
					g_12_26_12:
						for {
							if false {
								continue g_12_26_12
							}
							var x_prime__13 gopurs_runtime.Value = x_prime__13_loop
							_ = x_prime__13
							var y_prime__14 gopurs_runtime.Value = y_prime__14_loop
							_ = y_prime__14
							var __t28 int64
							{
								if (y_prime__14.IntVal) <= (0) {
									__t28 = x_prime__13.IntVal
									goto end_branch_28
								} else {

								}
							}
							{
								var __t27 int64
								{
									if (y_prime__14.IntVal) > (50000) {
										x_prime__13_loop = gopurs_runtime.Int((x_prime__13.IntVal) + (3))
										y_prime__14_loop = gopurs_runtime.Int((y_prime__14.IntVal) - (1))
										continue g_12_26_12
										__t27 = gopurs_runtime.Value{}.IntVal
										goto end_branch_27
									} else {

									}
								}
								{
									__t27 = Call_local_Main_f_9_25_11((x_prime__13.IntVal)+(2), gopurs_runtime.Int(y_prime__14.IntVal)).IntVal
								}
							end_branch_27:
								__t28 = __t27
							}
						end_branch_28:
							return gopurs_runtime.Int(__t28)
						}
					}
					g_12_26_12 = gopurs_runtime.Func(func(x_prime__13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(y_prime__14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return Call_local_Main_g_12_26_12(x_prime__13_loop_val, y_prime__14_loop_val)
						})
					})
					return gopurs_runtime.Int(Call_local_Main_g_12_26_12(gopurs_runtime.Int(x_10), gopurs_runtime.Int((y_11.IntVal)-(1))).IntVal)
				}
			}
			f_9_25_11 = gopurs_runtime.Func(func(x_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_9_25_11(x_10_loop_val.IntVal, y_11_loop_val)
				})
			})
			// TAST (Let): __local_var_9_24 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_9_24 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_local_Main_f_9_25_11(0, gopurs_runtime.Int(100000)).IntVal), gopurs_runtime.Int(249997))
			_ = __local_var_9_24
			// TAST (Let): result_10_29 shape=Other bindingType=Boolean
			result_10_29 := (gopurs_runtime.RecordGet(__local_var_9_24, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_9_24, "expected").IntVal)
			_ = result_10_29
			// TAST (Let): message_11_30 shape=Other bindingType=String
			message_11_30 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_9_24, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_9_24, "actual")).StrVal())
			_ = message_11_30
			// TAST (Let): __local_var_12_32 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_12_32 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_11_30))
			_ = __local_var_12_32
			var __t34 gopurs_runtime.Value
			{
				if (result_10_29) != (true) {
					__t34 = __local_var_12_32
					goto end_branch_34
				} else {

				}
			}
			{
				if result_10_29 {
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
			_dollar___unused_13_33 := gopurs_runtime.Apply(__t34, gopurs_runtime.Value{})
			_ = _dollar___unused_13_33
			_dollar___unused_12_31 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_11_30), gopurs_runtime.Bool(result_10_29)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_31
			var Call_local_Main_f_13_36_13 func(int64, int64) gopurs_runtime.Value
			_ = Call_local_Main_f_13_36_13
			var f_13_36_13 gopurs_runtime.Value
			_ = f_13_36_13
			Call_local_Main_f_13_36_13 = func(x_14_loop int64, y_15_loop int64) gopurs_runtime.Value {
			f_13_36_13:
				for {
					if false {
						continue f_13_36_13
					}
					var x_14 int64 = x_14_loop
					_ = x_14
					var y_15 int64 = y_15_loop
					_ = y_15
					var __t37 int64
					{
						if (y_15) <= (0) {
							__t37 = x_14
							goto end_branch_37
						} else {

						}
					}
					{
						x_14_loop = (x_14) + (2)
						y_15_loop = (y_15) - (1)
						continue f_13_36_13
						__t37 = gopurs_runtime.Value{}.IntVal
					}
				end_branch_37:
					return gopurs_runtime.Int(__t37)
				}
			}
			f_13_36_13 = gopurs_runtime.Func(func(x_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_15_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_13_36_13(x_14_loop_val.IntVal, y_15_loop_val.IntVal)
				})
			})
			// TAST (Let): __local_var_13_35 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_13_35 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_local_Main_f_13_36_13(0, 100000).IntVal), gopurs_runtime.Int(200000))
			_ = __local_var_13_35
			// TAST (Let): result_14_38 shape=Other bindingType=Boolean
			result_14_38 := (gopurs_runtime.RecordGet(__local_var_13_35, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_13_35, "expected").IntVal)
			_ = result_14_38
			// TAST (Let): message_15_39 shape=Other bindingType=String
			message_15_39 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_13_35, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_13_35, "actual")).StrVal())
			_ = message_15_39
			// TAST (Let): __local_var_16_41 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_16_41 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_15_39))
			_ = __local_var_16_41
			var __t43 gopurs_runtime.Value
			{
				if (result_14_38) != (true) {
					__t43 = __local_var_16_41
					goto end_branch_43
				} else {

				}
			}
			{
				if result_14_38 {
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
			_dollar___unused_17_42 := gopurs_runtime.Apply(__t43, gopurs_runtime.Value{})
			_ = _dollar___unused_17_42
			_dollar___unused_16_40 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_15_39), gopurs_runtime.Bool(result_14_38)), gopurs_runtime.Value{})
			_ = _dollar___unused_16_40
			var f_17_45_14 gopurs_runtime.Value
			_ = f_17_45_14
			f_17_45_14 = gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t46 gopurs_runtime.Value
				{
					if (x_18.IntVal) > (1000) {
						__t46 = gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((x_18.IntVal) + (v_19.IntVal))
						})
						goto end_branch_46
					} else {

					}
				}
				{
					__t46 = gopurs_runtime.Func(func(y_prime__19 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.Apply2(f_17_45_14, gopurs_runtime.Int((x_18.IntVal)+(10)), gopurs_runtime.Int((y_prime__19.IntVal)-(1))).IntVal)
					})
				}
			end_branch_46:
				return __t46
			})
			// TAST (Let): __local_var_17_44 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_17_44 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_17_45_14, gopurs_runtime.Int(0), gopurs_runtime.Int(100)).IntVal), gopurs_runtime.Int(1009))
			_ = __local_var_17_44
			// TAST (Let): result_18_47 shape=Other bindingType=Boolean
			result_18_47 := (gopurs_runtime.RecordGet(__local_var_17_44, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_17_44, "expected").IntVal)
			_ = result_18_47
			// TAST (Let): message_19_48 shape=Other bindingType=String
			message_19_48 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_17_44, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_17_44, "actual")).StrVal())
			_ = message_19_48
			// TAST (Let): __local_var_20_50 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_20_50 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_19_48))
			_ = __local_var_20_50
			var __t52 gopurs_runtime.Value
			{
				if (result_18_47) != (true) {
					__t52 = __local_var_20_50
					goto end_branch_52
				} else {

				}
			}
			{
				if result_18_47 {
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
			_dollar___unused_21_51 := gopurs_runtime.Apply(__t52, gopurs_runtime.Value{})
			_ = _dollar___unused_21_51
			_dollar___unused_20_49 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_19_48), gopurs_runtime.Bool(result_18_47)), gopurs_runtime.Value{})
			_ = _dollar___unused_20_49
			var Call_local_Main_f_21_54_15 func(int64, int64) gopurs_runtime.Value
			_ = Call_local_Main_f_21_54_15
			var f_21_54_15 gopurs_runtime.Value
			_ = f_21_54_15
			Call_local_Main_f_21_54_15 = func(x_22_loop int64, y_23_loop int64) gopurs_runtime.Value {
			f_21_54_15:
				for {
					if false {
						continue f_21_54_15
					}
					var x_22 int64 = x_22_loop
					_ = x_22
					var y_23 int64 = y_23_loop
					_ = y_23
					var __t55 int64
					{
						if (y_23) <= (0) {
							__t55 = x_22
							goto end_branch_55
						} else {

						}
					}
					{
						x_22_loop = (x_22) + (2)
						y_23_loop = (y_23) - (1)
						continue f_21_54_15
						__t55 = gopurs_runtime.Value{}.IntVal
					}
				end_branch_55:
					return gopurs_runtime.Int(__t55)
				}
			}
			f_21_54_15 = gopurs_runtime.Func(func(x_22_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_23_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_21_54_15(x_22_loop_val.IntVal, y_23_loop_val.IntVal)
				})
			})
			// TAST (Let): __local_var_21_53 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_21_53 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_local_Main_f_21_54_15(0, 100).IntVal), gopurs_runtime.Int(200))
			_ = __local_var_21_53
			// TAST (Let): result_22_56 shape=Other bindingType=Boolean
			result_22_56 := (gopurs_runtime.RecordGet(__local_var_21_53, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_21_53, "expected").IntVal)
			_ = result_22_56
			// TAST (Let): message_23_57 shape=Other bindingType=String
			message_23_57 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_21_53, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_21_53, "actual")).StrVal())
			_ = message_23_57
			// TAST (Let): __local_var_24_59 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_24_59 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_23_57))
			_ = __local_var_24_59
			var __t61 gopurs_runtime.Value
			{
				if (result_22_56) != (true) {
					__t61 = __local_var_24_59
					goto end_branch_61
				} else {

				}
			}
			{
				if result_22_56 {
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
			_dollar___unused_25_60 := gopurs_runtime.Apply(__t61, gopurs_runtime.Value{})
			_ = _dollar___unused_25_60
			_dollar___unused_24_58 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_23_57), gopurs_runtime.Bool(result_22_56)), gopurs_runtime.Value{})
			_ = _dollar___unused_24_58
			var f_25_63_16 gopurs_runtime.Value
			_ = f_25_63_16
			f_25_63_16 = gopurs_runtime.Func(func(x_26 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_27 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): g_28_64 shape=App(Other) bindingType=(Func [Int] Int)
					g_28_64 := gopurs_runtime.Apply(f_25_63_16, gopurs_runtime.Int((x_26.IntVal)+(2)))
					_ = g_28_64
					var __t65 int64
					{
						if (y_27.IntVal) <= (0) {
							__t65 = x_26.IntVal
							goto end_branch_65
						} else {

						}
					}
					{
						__t65 = gopurs_runtime.Apply(g_28_64, gopurs_runtime.Int((y_27.IntVal)-(1))).IntVal
					}
				end_branch_65:
					return gopurs_runtime.Int(__t65)
				})
			})
			// TAST (Let): __local_var_25_62 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_25_62 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(f_25_63_16, gopurs_runtime.Int(0), gopurs_runtime.Int(100)).IntVal), gopurs_runtime.Int(200))
			_ = __local_var_25_62
			// TAST (Let): result_26_66 shape=Other bindingType=Boolean
			result_26_66 := (gopurs_runtime.RecordGet(__local_var_25_62, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_25_62, "expected").IntVal)
			_ = result_26_66
			// TAST (Let): message_27_67 shape=Other bindingType=String
			message_27_67 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_25_62, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_25_62, "actual")).StrVal())
			_ = message_27_67
			// TAST (Let): __local_var_28_69 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_28_69 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_27_67))
			_ = __local_var_28_69
			var __t71 gopurs_runtime.Value
			{
				if (result_26_66) != (true) {
					__t71 = __local_var_28_69
					goto end_branch_71
				} else {

				}
			}
			{
				if result_26_66 {
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
			_dollar___unused_29_70 := gopurs_runtime.Apply(__t71, gopurs_runtime.Value{})
			_ = _dollar___unused_29_70
			_dollar___unused_28_68 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_27_67), gopurs_runtime.Bool(result_26_66)), gopurs_runtime.Value{})
			_ = _dollar___unused_28_68
			var Call_local_Main_f_29_73_17 func(int64, int64) gopurs_runtime.Value
			_ = Call_local_Main_f_29_73_17
			var f_29_73_17 gopurs_runtime.Value
			_ = f_29_73_17
			Call_local_Main_f_29_73_17 = func(x_30_loop int64, y_31_loop int64) gopurs_runtime.Value {
			f_29_73_17:
				for {
					if false {
						continue f_29_73_17
					}
					var x_30 int64 = x_30_loop
					_ = x_30
					var y_31 int64 = y_31_loop
					_ = y_31
					var __t74 int64
					{
						if (y_31) <= (0) {
							__t74 = x_30
							goto end_branch_74
						} else {

						}
					}
					{
						x_30_loop = (x_30) + (2)
						y_31_loop = (y_31) - (1)
						continue f_29_73_17
						__t74 = gopurs_runtime.Value{}.IntVal
					}
				end_branch_74:
					return gopurs_runtime.Int(__t74)
				}
			}
			f_29_73_17 = gopurs_runtime.Func(func(x_30_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_31_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_f_29_73_17(x_30_loop_val.IntVal, y_31_loop_val.IntVal)
				})
			})
			// TAST (Let): __local_var_29_72 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_29_72 := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_local_Main_f_29_73_17(0, 100).IntVal), gopurs_runtime.Int(200))
			_ = __local_var_29_72
			// TAST (Let): result_30_75 shape=Other bindingType=Boolean
			result_30_75 := (gopurs_runtime.RecordGet(__local_var_29_72, "actual").IntVal) == (gopurs_runtime.RecordGet(__local_var_29_72, "expected").IntVal)
			_ = result_30_75
			// TAST (Let): message_31_76 shape=Other bindingType=String
			message_31_76 := ((("Expected: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_29_72, "expected")).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(__local_var_29_72, "actual")).StrVal())
			_ = message_31_76
			// TAST (Let): __local_var_32_78 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_32_78 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_31_76))
			_ = __local_var_32_78
			var __t80 gopurs_runtime.Value
			{
				if (result_30_75) != (true) {
					__t80 = __local_var_32_78
					goto end_branch_80
				} else {

				}
			}
			{
				if result_30_75 {
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
			_dollar___unused_33_79 := gopurs_runtime.Apply(__t80, gopurs_runtime.Value{})
			_ = _dollar___unused_33_79
			_dollar___unused_32_77 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_31_76), gopurs_runtime.Bool(result_30_75)), gopurs_runtime.Value{})
			_ = _dollar___unused_32_77
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_tco3(y0_0_loop int64) int64 {
	var y0_0 int64 = y0_0_loop
	_ = y0_0
	var Call_local_Main_f_1_0_1 func(int64, gopurs_runtime.Value) gopurs_runtime.Value
	_ = Call_local_Main_f_1_0_1
	var f_1_0_1 gopurs_runtime.Value
	_ = f_1_0_1
	Call_local_Main_f_1_0_1 = func(x_2_loop int64, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	f_1_0_1:
		for {
			if false {
				continue f_1_0_1
			}
			var x_2 int64 = x_2_loop
			_ = x_2
			var y_3 gopurs_runtime.Value = y_3_loop
			_ = y_3
			var Call_local_Main_g_4_1_2 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_g_4_1_2
			var g_4_1_2 gopurs_runtime.Value
			_ = g_4_1_2
			Call_local_Main_g_4_1_2 = func(x_prime__5_loop gopurs_runtime.Value, y_prime__6_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
							__t2 = Call_local_Main_f_1_0_1((x_prime__5.IntVal)+(2), gopurs_runtime.Int(y_prime__6.IntVal)).IntVal
						}
					end_branch_2:
						__t3 = __t2
					}
				end_branch_3:
					return gopurs_runtime.Int(__t3)
				}
			}
			g_4_1_2 = gopurs_runtime.Func(func(x_prime__5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_prime__6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_g_4_1_2(x_prime__5_loop_val, y_prime__6_loop_val)
				})
			})
			return gopurs_runtime.Int(Call_local_Main_g_4_1_2(gopurs_runtime.Int(x_2), gopurs_runtime.Int((y_3.IntVal)-(1))).IntVal)
		}
	}
	f_1_0_1 = gopurs_runtime.Func(func(x_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Main_f_1_0_1(x_2_loop_val.IntVal, y_3_loop_val)
		})
	})
	return Call_local_Main_f_1_0_1(0, gopurs_runtime.Int(y0_0)).IntVal
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
