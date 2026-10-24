package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_tco4 gopurs_runtime.Value
var once_Main_tco4 sync.Once

func Get_Main_tco4() gopurs_runtime.Value {
	once_Main_tco4.Do(func() {
		cache_Main_tco4 = func() gopurs_runtime.Value {
			var Call_local_Main_f_0_0_0 func(int64, int64) int64
			_ = Call_local_Main_f_0_0_0
			var f_0_0_0 gopurs_runtime.Value
			_ = f_0_0_0
			Call_local_Main_f_0_0_0 = func(x_1_loop int64, y_2_loop int64) int64 {
				if ((x_1_loop) == (int64(int32(x_1_loop)))) && ((y_2_loop) == (int64(int32(y_2_loop)))) {
					return func() int64 {
						x_1_loop := int32(x_1_loop)
						_ = x_1_loop
						y_2_loop := int32(y_2_loop)
						_ = y_2_loop
					f_0_0_0:
						for {
							if false {
								continue f_0_0_0
							}
							x_1 := x_1_loop
							_ = x_1
							y_2 := y_2_loop
							_ = y_2
							var __t1 int64
							{
								if (y_2) <= (int32(0)) {
									__t1 = int64(x_1)
									goto end_branch_1
								} else {

								}
							}
							{
								x_1_loop = (x_1) + (int32(2))
								y_2_loop = (y_2) - (int32(1))
								continue f_0_0_0
								__t1 = func() int64 { panic("unreachable") }()
							}
						end_branch_1:
							return __t1
						}
					}()
				} else {

				}
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
						if (y_2) <= (int64(0)) {
							__t1 = x_1
							goto end_branch_1
						} else {

						}
					}
					{
						x_1_loop = gopurs_runtime.IntAdd(x_1, int64(2))
						y_2_loop = gopurs_runtime.IntSub(y_2, int64(1))
						continue f_0_0_0
						__t1 = func() int64 { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			f_0_0_0 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_0_0(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f_0_0_0, gopurs_runtime.Int(int64(0)))
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
			var Call_local_Main_f_0_0_3 func(int64, int64) int64
			_ = Call_local_Main_f_0_0_3
			var f_0_0_3 gopurs_runtime.Value
			_ = f_0_0_3
			Call_local_Main_f_0_0_3 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_0_3:
				for {
					if false {
						continue f_0_0_3
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_1 shape=Other bindingType=Int
					__local_var_3_1 := gopurs_runtime.IntAdd(x_1, int64(2))
					_ = __local_var_3_1
					// TAST (Let): __local_var_4_2 shape=Other bindingType=Int
					__local_var_4_2 := gopurs_runtime.IntSub(y_2, int64(1))
					_ = __local_var_4_2
					var __t3 int64
					{
						if (__local_var_4_2) <= (int64(0)) {
							__t3 = __local_var_3_1
							goto end_branch_3
						} else {

						}
					}
					{
						x_1_loop = __local_var_3_1
						y_2_loop = __local_var_4_2
						continue f_0_0_3
						__t3 = func() int64 { panic("unreachable") }()
					}
				end_branch_3:
					return __t3
				}
			}
			f_0_0_3 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_0_3(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f_0_0_3, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_tco2
}

var cache_Main_tco1 gopurs_runtime.Value
var once_Main_tco1 sync.Once

func Get_Main_tco1() gopurs_runtime.Value {
	once_Main_tco1.Do(func() {
		cache_Main_tco1 = func() gopurs_runtime.Value {
			var Call_local_Main_f_0_0_4 func(int64, int64) int64
			_ = Call_local_Main_f_0_0_4
			var f_0_0_4 gopurs_runtime.Value
			_ = f_0_0_4
			Call_local_Main_f_0_0_4 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_0_4:
				for {
					if false {
						continue f_0_0_4
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_1 shape=Other bindingType=Int
					__local_var_3_1 := gopurs_runtime.IntAdd(x_1, int64(2))
					_ = __local_var_3_1
					// TAST (Let): __local_var_4_2 shape=Other bindingType=Int
					__local_var_4_2 := gopurs_runtime.IntSub(y_2, int64(1))
					_ = __local_var_4_2
					var __t3 int64
					{
						if (__local_var_4_2) <= (int64(0)) {
							__t3 = __local_var_3_1
							goto end_branch_3
						} else {

						}
					}
					{
						x_1_loop = __local_var_3_1
						y_2_loop = __local_var_4_2
						continue f_0_0_4
						__t3 = func() int64 { panic("unreachable") }()
					}
				end_branch_3:
					return __t3
				}
			}
			f_0_0_4 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_0_4(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f_0_0_4, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_tco1
}

var cache_Main_ntco4 gopurs_runtime.Value
var once_Main_ntco4 sync.Once

func Get_Main_ntco4() gopurs_runtime.Value {
	once_Main_ntco4.Do(func() {
		cache_Main_ntco4 = func() gopurs_runtime.Value {
			var Call_local_Main_f_0_0_5 func(int64, int64) int64
			_ = Call_local_Main_f_0_0_5
			var f_0_0_5 gopurs_runtime.Value
			_ = f_0_0_5
			Call_local_Main_f_0_0_5 = func(x_1_loop int64, y_2_loop int64) int64 {
				if ((x_1_loop) == (int64(int32(x_1_loop)))) && ((y_2_loop) == (int64(int32(y_2_loop)))) {
					return func() int64 {
						x_1_loop := int32(x_1_loop)
						_ = x_1_loop
						y_2_loop := int32(y_2_loop)
						_ = y_2_loop
					f_0_0_5:
						for {
							if false {
								continue f_0_0_5
							}
							x_1 := x_1_loop
							_ = x_1
							y_2 := y_2_loop
							_ = y_2
							var __t1 int64
							{
								if (y_2) <= (int32(0)) {
									__t1 = int64(x_1)
									goto end_branch_1
								} else {

								}
							}
							{
								x_1_loop = (x_1) + (int32(2))
								y_2_loop = (y_2) - (int32(1))
								continue f_0_0_5
								__t1 = func() int64 { panic("unreachable") }()
							}
						end_branch_1:
							return __t1
						}
					}()
				} else {

				}
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
						if (y_2) <= (int64(0)) {
							__t1 = x_1
							goto end_branch_1
						} else {

						}
					}
					{
						x_1_loop = gopurs_runtime.IntAdd(x_1, int64(2))
						y_2_loop = gopurs_runtime.IntSub(y_2, int64(1))
						continue f_0_0_5
						__t1 = func() int64 { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			f_0_0_5 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_0_5(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f_0_0_5, gopurs_runtime.Int(int64(0)))
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
			var f_0_0_6_cell *gopurs_runtime.Value
			_ = f_0_0_6_cell
			// FALLBACK TCO: isLoop=false len=1
			f_0_0_6 = gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): g_3_1 shape=App(Other) bindingType=(Func [Int] Int)
				g_3_1 := gopurs_runtime.Apply((*f_0_0_6_cell), gopurs_runtime.Int(gopurs_runtime.IntAdd(x_1.IntVal, int64(2))))
				_ = g_3_1
				var __t2 gopurs_runtime.Value
				{
					if (y_2.IntVal) <= (int64(0)) {
						__t2 = x_1
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = gopurs_runtime.Apply(g_3_1, gopurs_runtime.Int(gopurs_runtime.IntSub(y_2.IntVal, int64(1))))
				}
			end_branch_2:
				return __t2
			})
			f_0_0_6_cell = &f_0_0_6
			return gopurs_runtime.Apply(f_0_0_6, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_ntco3
}

var cache_Main_ntco2 gopurs_runtime.Value
var once_Main_ntco2 sync.Once

func Get_Main_ntco2() gopurs_runtime.Value {
	once_Main_ntco2.Do(func() {
		cache_Main_ntco2 = func() gopurs_runtime.Value {
			var Call_local_Main_f_0_0_7 func(int64, int64) int64
			_ = Call_local_Main_f_0_0_7
			var f_0_0_7 gopurs_runtime.Value
			_ = f_0_0_7
			Call_local_Main_f_0_0_7 = func(x_1_loop int64, y_2_loop int64) int64 {
				if ((x_1_loop) == (int64(int32(x_1_loop)))) && ((y_2_loop) == (int64(int32(y_2_loop)))) {
					return func() int64 {
						x_1_loop := int32(x_1_loop)
						_ = x_1_loop
						y_2_loop := int32(y_2_loop)
						_ = y_2_loop
					f_0_0_7:
						for {
							if false {
								continue f_0_0_7
							}
							x_1 := x_1_loop
							_ = x_1
							y_2 := y_2_loop
							_ = y_2
							var __t1 int64
							{
								if (y_2) <= (int32(0)) {
									__t1 = int64(x_1)
									goto end_branch_1
								} else {

								}
							}
							{
								x_1_loop = (x_1) + (int32(2))
								y_2_loop = (y_2) - (int32(1))
								continue f_0_0_7
								__t1 = func() int64 { panic("unreachable") }()
							}
						end_branch_1:
							return __t1
						}
					}()
				} else {

				}
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
						if (y_2) <= (int64(0)) {
							__t1 = x_1
							goto end_branch_1
						} else {

						}
					}
					{
						x_1_loop = gopurs_runtime.IntAdd(x_1, int64(2))
						y_2_loop = gopurs_runtime.IntSub(y_2, int64(1))
						continue f_0_0_7
						__t1 = func() int64 { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			f_0_0_7 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_0_7(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f_0_0_7, gopurs_runtime.Int(int64(0)))
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
			var Call_local_Main_f_0_1_9 func(int64, int64) int64
			_ = Call_local_Main_f_0_1_9
			var f_0_1_9 gopurs_runtime.Value
			_ = f_0_1_9
			Call_local_Main_f_0_1_9 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_1_9:
				for {
					if false {
						continue f_0_1_9
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_2 shape=Other bindingType=Int
					__local_var_3_2 := gopurs_runtime.IntAdd(x_1, int64(2))
					_ = __local_var_3_2
					// TAST (Let): __local_var_4_3 shape=Other bindingType=Int
					__local_var_4_3 := gopurs_runtime.IntSub(y_2, int64(1))
					_ = __local_var_4_3
					var __t4 int64
					{
						if (__local_var_4_3) <= (int64(0)) {
							__t4 = __local_var_3_2
							goto end_branch_4
						} else {

						}
					}
					{
						x_1_loop = __local_var_3_2
						y_2_loop = __local_var_4_3
						continue f_0_1_9
						__t4 = func() int64 { panic("unreachable") }()
					}
				end_branch_4:
					return __t4
				}
			}
			f_0_1_9 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_1_9(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_local_Main_f_0_1_9(int64(0), int64(100000)), int64(200000)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}())
			_ = __local_var_0_0
			_dollar___unused_1_5 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_5
			var Call_local_Main_f_2_7_10 func(int64, int64) int64
			_ = Call_local_Main_f_2_7_10
			var f_2_7_10 gopurs_runtime.Value
			_ = f_2_7_10
			Call_local_Main_f_2_7_10 = func(x_3_loop int64, y_4_loop int64) int64 {
			f_2_7_10:
				for {
					if false {
						continue f_2_7_10
					}
					var x_3 int64 = x_3_loop
					_ = x_3
					var y_4 int64 = y_4_loop
					_ = y_4
					// TAST (Let): __local_var_5_8 shape=Other bindingType=Int
					__local_var_5_8 := gopurs_runtime.IntAdd(x_3, int64(2))
					_ = __local_var_5_8
					// TAST (Let): __local_var_6_9 shape=Other bindingType=Int
					__local_var_6_9 := gopurs_runtime.IntSub(y_4, int64(1))
					_ = __local_var_6_9
					var __t10 int64
					{
						if (__local_var_6_9) <= (int64(0)) {
							__t10 = __local_var_5_8
							goto end_branch_10
						} else {

						}
					}
					{
						x_3_loop = __local_var_5_8
						y_4_loop = __local_var_6_9
						continue f_2_7_10
						__t10 = func() int64 { panic("unreachable") }()
					}
				end_branch_10:
					return __t10
				}
			}
			f_2_7_10 = gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_2_7_10(x_3_loop_val.IntVal, y_4_loop_val.IntVal))
				})
			})
			_dollar___unused_2_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_local_Main_f_2_7_10(int64(0), int64(100000)), int64(200000)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_2_6
			var Call_local_Main_f_3_12_11 func(int64, int64) int64
			_ = Call_local_Main_f_3_12_11
			var f_3_12_11 gopurs_runtime.Value
			_ = f_3_12_11
			Call_local_Main_f_3_12_11 = func(x_4_loop int64, y_5_loop int64) int64 {
			f_3_12_11:
				for {
					if false {
						continue f_3_12_11
					}
					var x_4 int64 = x_4_loop
					_ = x_4
					var y_5 int64 = y_5_loop
					_ = y_5
					var Call_local_Main_g_6_13_12 func(int64, int64) int64
					_ = Call_local_Main_g_6_13_12
					var g_6_13_12 gopurs_runtime.Value
					_ = g_6_13_12
					Call_local_Main_g_6_13_12 = func(x_prime__7_loop int64, y_prime__8_loop int64) int64 {
					g_6_13_12:
						for {
							if false {
								continue g_6_13_12
							}
							var x_prime__7 int64 = x_prime__7_loop
							_ = x_prime__7
							var y_prime__8 int64 = y_prime__8_loop
							_ = y_prime__8
							var __t15 int64
							{
								if (y_prime__8) <= (int64(0)) {
									__t15 = x_prime__7
									goto end_branch_15
								} else {

								}
							}
							{
								var __t14 int64
								{
									if (y_prime__8) > (int64(50000)) {
										x_prime__7_loop = gopurs_runtime.IntAdd(x_prime__7, int64(3))
										y_prime__8_loop = gopurs_runtime.IntSub(y_prime__8, int64(1))
										continue g_6_13_12
										__t14 = func() int64 { panic("unreachable") }()
										goto end_branch_14
									} else {

									}
								}
								{
									__t14 = Call_local_Main_f_3_12_11(gopurs_runtime.IntAdd(x_prime__7, int64(2)), y_prime__8)
								}
							end_branch_14:
								__t15 = __t14
							}
						end_branch_15:
							return __t15
						}
					}
					g_6_13_12 = gopurs_runtime.Func(func(x_prime__7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(y_prime__8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int(Call_local_Main_g_6_13_12(x_prime__7_loop_val.IntVal, y_prime__8_loop_val.IntVal))
						})
					})
					return Call_local_Main_g_6_13_12(x_4, gopurs_runtime.IntSub(y_5, int64(1)))
				}
			}
			f_3_12_11 = gopurs_runtime.Func(func(x_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_3_12_11(x_4_loop_val.IntVal, y_5_loop_val.IntVal))
				})
			})
			_dollar___unused_3_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_local_Main_f_3_12_11(int64(0), int64(100000)), int64(249997)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_3_11
			var Call_local_Main_f_4_17_13 func(int64, int64) int64
			_ = Call_local_Main_f_4_17_13
			var f_4_17_13 gopurs_runtime.Value
			_ = f_4_17_13
			Call_local_Main_f_4_17_13 = func(x_5_loop int64, y_6_loop int64) int64 {
				if ((x_5_loop) == (int64(int32(x_5_loop)))) && ((y_6_loop) == (int64(int32(y_6_loop)))) {
					return func() int64 {
						x_5_loop := int32(x_5_loop)
						_ = x_5_loop
						y_6_loop := int32(y_6_loop)
						_ = y_6_loop
					f_4_17_13:
						for {
							if false {
								continue f_4_17_13
							}
							x_5 := x_5_loop
							_ = x_5
							y_6 := y_6_loop
							_ = y_6
							var __t18 int64
							{
								if (y_6) <= (int32(0)) {
									__t18 = int64(x_5)
									goto end_branch_18
								} else {

								}
							}
							{
								x_5_loop = (x_5) + (int32(2))
								y_6_loop = (y_6) - (int32(1))
								continue f_4_17_13
								__t18 = func() int64 { panic("unreachable") }()
							}
						end_branch_18:
							return __t18
						}
					}()
				} else {

				}
			f_4_17_13:
				for {
					if false {
						continue f_4_17_13
					}
					var x_5 int64 = x_5_loop
					_ = x_5
					var y_6 int64 = y_6_loop
					_ = y_6
					var __t18 int64
					{
						if (y_6) <= (int64(0)) {
							__t18 = x_5
							goto end_branch_18
						} else {

						}
					}
					{
						x_5_loop = gopurs_runtime.IntAdd(x_5, int64(2))
						y_6_loop = gopurs_runtime.IntSub(y_6, int64(1))
						continue f_4_17_13
						__t18 = func() int64 { panic("unreachable") }()
					}
				end_branch_18:
					return __t18
				}
			}
			f_4_17_13 = gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_4_17_13(x_5_loop_val.IntVal, y_6_loop_val.IntVal))
				})
			})
			_dollar___unused_4_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_local_Main_f_4_17_13(int64(0), int64(100000)), int64(200000)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_4_16
			var f_5_20_14 gopurs_runtime.Value
			_ = f_5_20_14
			var f_5_20_14_cell *gopurs_runtime.Value
			_ = f_5_20_14_cell
			// FALLBACK TCO: isLoop=false len=1
			f_5_20_14 = gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t21 gopurs_runtime.Value
				{
					if (x_6.IntVal) > (int64(1000)) {
						__t21 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_6.IntVal, v_7.IntVal))
						})
						goto end_branch_21
					} else {

					}
				}
				{
					__t21 = gopurs_runtime.Func(func(y_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2((*f_5_20_14_cell), gopurs_runtime.Int(gopurs_runtime.IntAdd(x_6.IntVal, int64(10))), gopurs_runtime.Int(gopurs_runtime.IntSub(y_prime__7.IntVal, int64(1))))
					})
				}
			end_branch_21:
				return __t21
			})
			f_5_20_14_cell = &f_5_20_14
			_dollar___unused_5_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply2(f_5_20_14, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(100))).IntVal, int64(1009)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_5_19
			var Call_local_Main_f_6_23_15 func(int64, int64) int64
			_ = Call_local_Main_f_6_23_15
			var f_6_23_15 gopurs_runtime.Value
			_ = f_6_23_15
			Call_local_Main_f_6_23_15 = func(x_7_loop int64, y_8_loop int64) int64 {
				if ((x_7_loop) == (int64(int32(x_7_loop)))) && ((y_8_loop) == (int64(int32(y_8_loop)))) {
					return func() int64 {
						x_7_loop := int32(x_7_loop)
						_ = x_7_loop
						y_8_loop := int32(y_8_loop)
						_ = y_8_loop
					f_6_23_15:
						for {
							if false {
								continue f_6_23_15
							}
							x_7 := x_7_loop
							_ = x_7
							y_8 := y_8_loop
							_ = y_8
							var __t24 int64
							{
								if (y_8) <= (int32(0)) {
									__t24 = int64(x_7)
									goto end_branch_24
								} else {

								}
							}
							{
								x_7_loop = (x_7) + (int32(2))
								y_8_loop = (y_8) - (int32(1))
								continue f_6_23_15
								__t24 = func() int64 { panic("unreachable") }()
							}
						end_branch_24:
							return __t24
						}
					}()
				} else {

				}
			f_6_23_15:
				for {
					if false {
						continue f_6_23_15
					}
					var x_7 int64 = x_7_loop
					_ = x_7
					var y_8 int64 = y_8_loop
					_ = y_8
					var __t24 int64
					{
						if (y_8) <= (int64(0)) {
							__t24 = x_7
							goto end_branch_24
						} else {

						}
					}
					{
						x_7_loop = gopurs_runtime.IntAdd(x_7, int64(2))
						y_8_loop = gopurs_runtime.IntSub(y_8, int64(1))
						continue f_6_23_15
						__t24 = func() int64 { panic("unreachable") }()
					}
				end_branch_24:
					return __t24
				}
			}
			f_6_23_15 = gopurs_runtime.Func(func(x_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_6_23_15(x_7_loop_val.IntVal, y_8_loop_val.IntVal))
				})
			})
			_dollar___unused_6_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_local_Main_f_6_23_15(int64(0), int64(100)), int64(200)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_6_22
			var f_7_26_16 gopurs_runtime.Value
			_ = f_7_26_16
			var f_7_26_16_cell *gopurs_runtime.Value
			_ = f_7_26_16_cell
			// FALLBACK TCO: isLoop=false len=1
			f_7_26_16 = gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): g_10_27 shape=App(Other) bindingType=(Func [Int] Int)
				g_10_27 := gopurs_runtime.Apply((*f_7_26_16_cell), gopurs_runtime.Int(gopurs_runtime.IntAdd(x_8.IntVal, int64(2))))
				_ = g_10_27
				var __t28 gopurs_runtime.Value
				{
					if (y_9.IntVal) <= (int64(0)) {
						__t28 = x_8
						goto end_branch_28
					} else {

					}
				}
				{
					__t28 = gopurs_runtime.Apply(g_10_27, gopurs_runtime.Int(gopurs_runtime.IntSub(y_9.IntVal, int64(1))))
				}
			end_branch_28:
				return __t28
			})
			f_7_26_16_cell = &f_7_26_16
			_dollar___unused_7_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply2(f_7_26_16, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(100))).IntVal, int64(200)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_7_25
			var Call_local_Main_f_8_30_17 func(int64, int64) int64
			_ = Call_local_Main_f_8_30_17
			var f_8_30_17 gopurs_runtime.Value
			_ = f_8_30_17
			Call_local_Main_f_8_30_17 = func(x_9_loop int64, y_10_loop int64) int64 {
				if ((x_9_loop) == (int64(int32(x_9_loop)))) && ((y_10_loop) == (int64(int32(y_10_loop)))) {
					return func() int64 {
						x_9_loop := int32(x_9_loop)
						_ = x_9_loop
						y_10_loop := int32(y_10_loop)
						_ = y_10_loop
					f_8_30_17:
						for {
							if false {
								continue f_8_30_17
							}
							x_9 := x_9_loop
							_ = x_9
							y_10 := y_10_loop
							_ = y_10
							var __t31 int64
							{
								if (y_10) <= (int32(0)) {
									__t31 = int64(x_9)
									goto end_branch_31
								} else {

								}
							}
							{
								x_9_loop = (x_9) + (int32(2))
								y_10_loop = (y_10) - (int32(1))
								continue f_8_30_17
								__t31 = func() int64 { panic("unreachable") }()
							}
						end_branch_31:
							return __t31
						}
					}()
				} else {

				}
			f_8_30_17:
				for {
					if false {
						continue f_8_30_17
					}
					var x_9 int64 = x_9_loop
					_ = x_9
					var y_10 int64 = y_10_loop
					_ = y_10
					var __t31 int64
					{
						if (y_10) <= (int64(0)) {
							__t31 = x_9
							goto end_branch_31
						} else {

						}
					}
					{
						x_9_loop = gopurs_runtime.IntAdd(x_9, int64(2))
						y_10_loop = gopurs_runtime.IntSub(y_10, int64(1))
						continue f_8_30_17
						__t31 = func() int64 { panic("unreachable") }()
					}
				end_branch_31:
					return __t31
				}
			}
			f_8_30_17 = gopurs_runtime.Func(func(x_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_8_30_17(x_9_loop_val.IntVal, y_10_loop_val.IntVal))
				})
			})
			_dollar___unused_8_29 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_local_Main_f_8_30_17(int64(0), int64(100)), int64(200)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_8_29
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_tco3(y0_0_loop int64) int64 {
	var y0_0 int64 = y0_0_loop
	_ = y0_0
	var Call_local_Main_f_1_0_1 func(int64, int64) int64
	_ = Call_local_Main_f_1_0_1
	var f_1_0_1 gopurs_runtime.Value
	_ = f_1_0_1
	Call_local_Main_f_1_0_1 = func(x_2_loop int64, y_3_loop int64) int64 {
	f_1_0_1:
		for {
			if false {
				continue f_1_0_1
			}
			var x_2 int64 = x_2_loop
			_ = x_2
			var y_3 int64 = y_3_loop
			_ = y_3
			var Call_local_Main_g_4_1_2 func(int64, int64) int64
			_ = Call_local_Main_g_4_1_2
			var g_4_1_2 gopurs_runtime.Value
			_ = g_4_1_2
			Call_local_Main_g_4_1_2 = func(x_prime__5_loop int64, y_prime__6_loop int64) int64 {
			g_4_1_2:
				for {
					if false {
						continue g_4_1_2
					}
					var x_prime__5 int64 = x_prime__5_loop
					_ = x_prime__5
					var y_prime__6 int64 = y_prime__6_loop
					_ = y_prime__6
					var __t3 int64
					{
						if (y_prime__6) <= (int64(0)) {
							__t3 = x_prime__5
							goto end_branch_3
						} else {

						}
					}
					{
						var __t2 int64
						{
							if (y_prime__6) > (gopurs_runtime.IntDiv(y0_0, int64(2))) {
								x_prime__5_loop = gopurs_runtime.IntAdd(x_prime__5, int64(3))
								y_prime__6_loop = gopurs_runtime.IntSub(y_prime__6, int64(1))
								continue g_4_1_2
								__t2 = func() int64 { panic("unreachable") }()
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = Call_local_Main_f_1_0_1(gopurs_runtime.IntAdd(x_prime__5, int64(2)), y_prime__6)
						}
					end_branch_2:
						__t3 = __t2
					}
				end_branch_3:
					return __t3
				}
			}
			g_4_1_2 = gopurs_runtime.Func(func(x_prime__5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_prime__6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_g_4_1_2(x_prime__5_loop_val.IntVal, y_prime__6_loop_val.IntVal))
				})
			})
			return Call_local_Main_g_4_1_2(x_2, gopurs_runtime.IntSub(y_3, int64(1)))
		}
	}
	f_1_0_1 = gopurs_runtime.Func(func(x_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_f_1_0_1(x_2_loop_val.IntVal, y_3_loop_val.IntVal))
		})
	})
	return Call_local_Main_f_1_0_1(int64(0), y0_0)
}

func Call_Main_ntco1(y0_0_loop int64) int64 {
	var y0_0 int64 = y0_0_loop
	_ = y0_0
	var f_1_0_8 gopurs_runtime.Value
	_ = f_1_0_8
	var f_1_0_8_cell *gopurs_runtime.Value
	_ = f_1_0_8_cell
	// FALLBACK TCO: isLoop=false len=1
	f_1_0_8 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (x_2.IntVal) > (gopurs_runtime.IntMul(int64(10), y0_0)) {
				__t1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_2.IntVal, v_3.IntVal))
				})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Func(func(y_prime__3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2((*f_1_0_8_cell), gopurs_runtime.Int(gopurs_runtime.IntAdd(x_2.IntVal, int64(10))), gopurs_runtime.Int(gopurs_runtime.IntSub(y_prime__3.IntVal, int64(1))))
			})
		}
	end_branch_1:
		return __t1
	})
	f_1_0_8_cell = &f_1_0_8
	return gopurs_runtime.Apply2(f_1_0_8, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(y0_0)).IntVal
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
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

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}
