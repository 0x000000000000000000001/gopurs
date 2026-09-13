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
			var Call_local_Main_f__467072791_0_0_0 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_0_0_0
			var f__467072791_0_0_0 gopurs_runtime.Value
			_ = f__467072791_0_0_0
			var Call_local_Main_f_0_1_1 func(int64, int64) int64
			_ = Call_local_Main_f_0_1_1
			var f_0_1_1 gopurs_runtime.Value
			_ = f_0_1_1
			Call_local_Main_f__467072791_0_0_0 = func(x_1_loop int64, y_2_loop int64) int64 {
			f__467072791_0_0_0:
				for {
					if false {
						continue f__467072791_0_0_0
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					var __t2 int64
					{
						if (y_2) <= (int64(0)) {
							__t2 = x_1
							goto end_branch_2
						} else {

						}
					}
					{
						x_1_loop = (x_1) + (int64(2))
						y_2_loop = (y_2) - (int64(1))
						continue f__467072791_0_0_0
						__t2 = func() int64 { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			f__467072791_0_0_0 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_0_0_0(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			Call_local_Main_f_0_1_1 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_1_1:
				for {
					if false {
						continue f_0_1_1
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					var __t3 int64
					{
						if (y_2) <= (int64(0)) {
							__t3 = x_1
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_f__467072791_0_0_0((x_1)+(int64(2)), (y_2)-(int64(1)))
					}
				end_branch_3:
					return __t3
				}
			}
			f_0_1_1 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_1_1(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f__467072791_0_0_0, gopurs_runtime.Int(int64(0)))
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
			var Call_local_Main_f__467072791_0_0_8 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_0_0_8
			var f__467072791_0_0_8 gopurs_runtime.Value
			_ = f__467072791_0_0_8
			var Call_local_Main_f_0_1_9 func(int64, int64) int64
			_ = Call_local_Main_f_0_1_9
			var f_0_1_9 gopurs_runtime.Value
			_ = f_0_1_9
			Call_local_Main_f__467072791_0_0_8 = func(x_1_loop int64, y_2_loop int64) int64 {
			f__467072791_0_0_8:
				for {
					if false {
						continue f__467072791_0_0_8
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_2 shape=Other bindingType=Int
					__local_var_3_2 := (x_1) + (int64(2))
					_ = __local_var_3_2
					// TAST (Let): __local_var_4_3 shape=Other bindingType=Int
					__local_var_4_3 := (y_2) - (int64(1))
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
						continue f__467072791_0_0_8
						__t4 = func() int64 { panic("unreachable") }()
					}
				end_branch_4:
					return __t4
				}
			}
			f__467072791_0_0_8 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_0_0_8(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
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
					// TAST (Let): __local_var_3_5 shape=Other bindingType=Int
					__local_var_3_5 := (x_1) + (int64(2))
					_ = __local_var_3_5
					// TAST (Let): __local_var_4_6 shape=Other bindingType=Int
					__local_var_4_6 := (y_2) - (int64(1))
					_ = __local_var_4_6
					var __t7 int64
					{
						if (__local_var_4_6) <= (int64(0)) {
							__t7 = __local_var_3_5
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = Call_local_Main_f__467072791_0_0_8(__local_var_3_5, __local_var_4_6)
					}
				end_branch_7:
					return __t7
				}
			}
			f_0_1_9 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_1_9(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f__467072791_0_0_8, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_tco2
}

var cache_Main_tco1 gopurs_runtime.Value
var once_Main_tco1 sync.Once

func Get_Main_tco1() gopurs_runtime.Value {
	once_Main_tco1.Do(func() {
		cache_Main_tco1 = func() gopurs_runtime.Value {
			var Call_local_Main_f__467072791_0_0_10 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_0_0_10
			var f__467072791_0_0_10 gopurs_runtime.Value
			_ = f__467072791_0_0_10
			var Call_local_Main_f_0_1_11 func(int64, int64) int64
			_ = Call_local_Main_f_0_1_11
			var f_0_1_11 gopurs_runtime.Value
			_ = f_0_1_11
			Call_local_Main_f__467072791_0_0_10 = func(x_1_loop int64, y_2_loop int64) int64 {
			f__467072791_0_0_10:
				for {
					if false {
						continue f__467072791_0_0_10
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_2 shape=Other bindingType=Int
					__local_var_3_2 := (x_1) + (int64(2))
					_ = __local_var_3_2
					// TAST (Let): __local_var_4_3 shape=Other bindingType=Int
					__local_var_4_3 := (y_2) - (int64(1))
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
						continue f__467072791_0_0_10
						__t4 = func() int64 { panic("unreachable") }()
					}
				end_branch_4:
					return __t4
				}
			}
			f__467072791_0_0_10 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_0_0_10(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			Call_local_Main_f_0_1_11 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_1_11:
				for {
					if false {
						continue f_0_1_11
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_5 shape=Other bindingType=Int
					__local_var_3_5 := (x_1) + (int64(2))
					_ = __local_var_3_5
					// TAST (Let): __local_var_4_6 shape=Other bindingType=Int
					__local_var_4_6 := (y_2) - (int64(1))
					_ = __local_var_4_6
					var __t7 int64
					{
						if (__local_var_4_6) <= (int64(0)) {
							__t7 = __local_var_3_5
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = Call_local_Main_f__467072791_0_0_10(__local_var_3_5, __local_var_4_6)
					}
				end_branch_7:
					return __t7
				}
			}
			f_0_1_11 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_1_11(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f__467072791_0_0_10, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_tco1
}

var cache_Main_ntco4 gopurs_runtime.Value
var once_Main_ntco4 sync.Once

func Get_Main_ntco4() gopurs_runtime.Value {
	once_Main_ntco4.Do(func() {
		cache_Main_ntco4 = func() gopurs_runtime.Value {
			var Call_local_Main_f__467072791_0_0_12 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_0_0_12
			var f__467072791_0_0_12 gopurs_runtime.Value
			_ = f__467072791_0_0_12
			var Call_local_Main_f_0_1_13 func(int64, int64) int64
			_ = Call_local_Main_f_0_1_13
			var f_0_1_13 gopurs_runtime.Value
			_ = f_0_1_13
			Call_local_Main_f__467072791_0_0_12 = func(x_1_loop int64, y_2_loop int64) int64 {
			f__467072791_0_0_12:
				for {
					if false {
						continue f__467072791_0_0_12
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					var __t2 int64
					{
						if (y_2) <= (int64(0)) {
							__t2 = x_1
							goto end_branch_2
						} else {

						}
					}
					{
						x_1_loop = (x_1) + (int64(2))
						y_2_loop = (y_2) - (int64(1))
						continue f__467072791_0_0_12
						__t2 = func() int64 { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			f__467072791_0_0_12 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_0_0_12(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			Call_local_Main_f_0_1_13 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_1_13:
				for {
					if false {
						continue f_0_1_13
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					var __t3 int64
					{
						if (y_2) <= (int64(0)) {
							__t3 = x_1
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_f__467072791_0_0_12((x_1)+(int64(2)), (y_2)-(int64(1)))
					}
				end_branch_3:
					return __t3
				}
			}
			f_0_1_13 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_1_13(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f__467072791_0_0_12, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_ntco4
}

var cache_Main_ntco3 gopurs_runtime.Value
var once_Main_ntco3 sync.Once

func Get_Main_ntco3() gopurs_runtime.Value {
	once_Main_ntco3.Do(func() {
		cache_Main_ntco3 = func() gopurs_runtime.Value {
			var f__467072791_0_0_14 gopurs_runtime.Value
			_ = f__467072791_0_0_14
			var f__467072791_0_0_14_cell *gopurs_runtime.Value
			_ = f__467072791_0_0_14_cell
			// FALLBACK TCO: isLoop=false len=2
			var f_0_1_15 gopurs_runtime.Value
			_ = f_0_1_15
			var f_0_1_15_cell *gopurs_runtime.Value
			_ = f_0_1_15_cell
			// FALLBACK TCO: isLoop=false len=2
			f__467072791_0_0_14 = gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): g__3466805691_3_2 shape=App(Other) bindingType=(Func [Int] Int)
				g__3466805691_3_2 := gopurs_runtime.Apply((*f__467072791_0_0_14_cell), gopurs_runtime.Int((x_1.IntVal)+(int64(2))))
				_ = g__3466805691_3_2
				var __t3 int64
				{
					if (y_2.IntVal) <= (int64(0)) {
						__t3 = x_1.IntVal
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Apply(g__3466805691_3_2, gopurs_runtime.Int((y_2.IntVal)-(int64(1)))).IntVal
				}
			end_branch_3:
				return gopurs_runtime.Int(__t3)
			})
			f__467072791_0_0_14_cell = &f__467072791_0_0_14
			f_0_1_15 = gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): g__3466805691_3_4 shape=App(Other) bindingType=(Func [Int] Int)
				g__3466805691_3_4 := gopurs_runtime.Apply((*f__467072791_0_0_14_cell), gopurs_runtime.Int((x_1.IntVal)+(int64(2))))
				_ = g__3466805691_3_4
				var __t5 int64
				{
					if (y_2.IntVal) <= (int64(0)) {
						__t5 = x_1.IntVal
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = gopurs_runtime.Apply(g__3466805691_3_4, gopurs_runtime.Int((y_2.IntVal)-(int64(1)))).IntVal
				}
			end_branch_5:
				return gopurs_runtime.Int(__t5)
			})
			f_0_1_15_cell = &f_0_1_15
			return gopurs_runtime.Apply(f__467072791_0_0_14, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_ntco3
}

var cache_Main_ntco2 gopurs_runtime.Value
var once_Main_ntco2 sync.Once

func Get_Main_ntco2() gopurs_runtime.Value {
	once_Main_ntco2.Do(func() {
		cache_Main_ntco2 = func() gopurs_runtime.Value {
			var Call_local_Main_f__467072791_0_0_16 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_0_0_16
			var f__467072791_0_0_16 gopurs_runtime.Value
			_ = f__467072791_0_0_16
			var Call_local_Main_f_0_1_17 func(int64, int64) int64
			_ = Call_local_Main_f_0_1_17
			var f_0_1_17 gopurs_runtime.Value
			_ = f_0_1_17
			Call_local_Main_f__467072791_0_0_16 = func(x_1_loop int64, y_2_loop int64) int64 {
			f__467072791_0_0_16:
				for {
					if false {
						continue f__467072791_0_0_16
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					var __t2 int64
					{
						if (y_2) <= (int64(0)) {
							__t2 = x_1
							goto end_branch_2
						} else {

						}
					}
					{
						x_1_loop = (x_1) + (int64(2))
						y_2_loop = (y_2) - (int64(1))
						continue f__467072791_0_0_16
						__t2 = func() int64 { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			f__467072791_0_0_16 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_0_0_16(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			Call_local_Main_f_0_1_17 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_1_17:
				for {
					if false {
						continue f_0_1_17
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					var __t3 int64
					{
						if (y_2) <= (int64(0)) {
							__t3 = x_1
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_f__467072791_0_0_16((x_1)+(int64(2)), (y_2)-(int64(1)))
					}
				end_branch_3:
					return __t3
				}
			}
			f_0_1_17 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_1_17(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply(f__467072791_0_0_16, gopurs_runtime.Int(int64(0)))
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
		cache_Main_main = func() gopurs_runtime.Value {
			var Call_local_Main_f__467072791_0_0_20 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_0_0_20
			var f__467072791_0_0_20 gopurs_runtime.Value
			_ = f__467072791_0_0_20
			var Call_local_Main_f_0_1_21 func(int64, int64) int64
			_ = Call_local_Main_f_0_1_21
			var f_0_1_21 gopurs_runtime.Value
			_ = f_0_1_21
			Call_local_Main_f__467072791_0_0_20 = func(x_1_loop int64, y_2_loop int64) int64 {
			f__467072791_0_0_20:
				for {
					if false {
						continue f__467072791_0_0_20
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_2 shape=Other bindingType=Int
					__local_var_3_2 := (x_1) + (int64(2))
					_ = __local_var_3_2
					// TAST (Let): __local_var_4_3 shape=Other bindingType=Int
					__local_var_4_3 := (y_2) - (int64(1))
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
						continue f__467072791_0_0_20
						__t4 = func() int64 { panic("unreachable") }()
					}
				end_branch_4:
					return __t4
				}
			}
			f__467072791_0_0_20 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_0_0_20(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			Call_local_Main_f_0_1_21 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_1_21:
				for {
					if false {
						continue f_0_1_21
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_5 shape=Other bindingType=Int
					__local_var_3_5 := (x_1) + (int64(2))
					_ = __local_var_3_5
					// TAST (Let): __local_var_4_6 shape=Other bindingType=Int
					__local_var_4_6 := (y_2) - (int64(1))
					_ = __local_var_4_6
					var __t7 int64
					{
						if (__local_var_4_6) <= (int64(0)) {
							__t7 = __local_var_3_5
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = Call_local_Main_f__467072791_0_0_20(__local_var_3_5, __local_var_4_6)
					}
				end_branch_7:
					return __t7
				}
			}
			f_0_1_21 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_1_21(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_f__467072791_0_0_20(int64(0), int64(100000)), int64(200000)}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var Call_local_Main_f__467072791_1_8_22 func(int64, int64) int64
				_ = Call_local_Main_f__467072791_1_8_22
				var f__467072791_1_8_22 gopurs_runtime.Value
				_ = f__467072791_1_8_22
				var Call_local_Main_f_1_9_23 func(int64, int64) int64
				_ = Call_local_Main_f_1_9_23
				var f_1_9_23 gopurs_runtime.Value
				_ = f_1_9_23
				Call_local_Main_f__467072791_1_8_22 = func(x_2_loop int64, y_3_loop int64) int64 {
				f__467072791_1_8_22:
					for {
						if false {
							continue f__467072791_1_8_22
						}
						var x_2 int64 = x_2_loop
						_ = x_2
						var y_3 int64 = y_3_loop
						_ = y_3
						// TAST (Let): __local_var_4_10 shape=Other bindingType=Int
						__local_var_4_10 := (x_2) + (int64(2))
						_ = __local_var_4_10
						// TAST (Let): __local_var_5_11 shape=Other bindingType=Int
						__local_var_5_11 := (y_3) - (int64(1))
						_ = __local_var_5_11
						var __t12 int64
						{
							if (__local_var_5_11) <= (int64(0)) {
								__t12 = __local_var_4_10
								goto end_branch_12
							} else {

							}
						}
						{
							x_2_loop = __local_var_4_10
							y_3_loop = __local_var_5_11
							continue f__467072791_1_8_22
							__t12 = func() int64 { panic("unreachable") }()
						}
					end_branch_12:
						return __t12
					}
				}
				f__467072791_1_8_22 = gopurs_runtime.Func(func(x_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(Call_local_Main_f__467072791_1_8_22(x_2_loop_val.IntVal, y_3_loop_val.IntVal))
					})
				})
				Call_local_Main_f_1_9_23 = func(x_2_loop int64, y_3_loop int64) int64 {
				f_1_9_23:
					for {
						if false {
							continue f_1_9_23
						}
						var x_2 int64 = x_2_loop
						_ = x_2
						var y_3 int64 = y_3_loop
						_ = y_3
						// TAST (Let): __local_var_4_13 shape=Other bindingType=Int
						__local_var_4_13 := (x_2) + (int64(2))
						_ = __local_var_4_13
						// TAST (Let): __local_var_5_14 shape=Other bindingType=Int
						__local_var_5_14 := (y_3) - (int64(1))
						_ = __local_var_5_14
						var __t15 int64
						{
							if (__local_var_5_14) <= (int64(0)) {
								__t15 = __local_var_4_13
								goto end_branch_15
							} else {

							}
						}
						{
							__t15 = Call_local_Main_f__467072791_1_8_22(__local_var_4_13, __local_var_5_14)
						}
					end_branch_15:
						return __t15
					}
				}
				f_1_9_23 = gopurs_runtime.Func(func(x_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(Call_local_Main_f_1_9_23(x_2_loop_val.IntVal, y_3_loop_val.IntVal))
					})
				})
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{Call_local_Main_f__467072791_1_8_22(int64(0), int64(100000)), int64(200000)}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{Call_Main_tco3(int64(100000)), int64(249997)}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
						var Call_local_Main_f__467072791_3_16_24 func(int64, int64) int64
						_ = Call_local_Main_f__467072791_3_16_24
						var f__467072791_3_16_24 gopurs_runtime.Value
						_ = f__467072791_3_16_24
						var Call_local_Main_f_3_17_25 func(int64, int64) int64
						_ = Call_local_Main_f_3_17_25
						var f_3_17_25 gopurs_runtime.Value
						_ = f_3_17_25
						Call_local_Main_f__467072791_3_16_24 = func(x_4_loop int64, y_5_loop int64) int64 {
						f__467072791_3_16_24:
							for {
								if false {
									continue f__467072791_3_16_24
								}
								var x_4 int64 = x_4_loop
								_ = x_4
								var y_5 int64 = y_5_loop
								_ = y_5
								var __t18 int64
								{
									if (y_5) <= (int64(0)) {
										__t18 = x_4
										goto end_branch_18
									} else {

									}
								}
								{
									x_4_loop = (x_4) + (int64(2))
									y_5_loop = (y_5) - (int64(1))
									continue f__467072791_3_16_24
									__t18 = func() int64 { panic("unreachable") }()
								}
							end_branch_18:
								return __t18
							}
						}
						f__467072791_3_16_24 = gopurs_runtime.Func(func(x_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(y_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Int(Call_local_Main_f__467072791_3_16_24(x_4_loop_val.IntVal, y_5_loop_val.IntVal))
							})
						})
						Call_local_Main_f_3_17_25 = func(x_4_loop int64, y_5_loop int64) int64 {
						f_3_17_25:
							for {
								if false {
									continue f_3_17_25
								}
								var x_4 int64 = x_4_loop
								_ = x_4
								var y_5 int64 = y_5_loop
								_ = y_5
								var __t19 int64
								{
									if (y_5) <= (int64(0)) {
										__t19 = x_4
										goto end_branch_19
									} else {

									}
								}
								{
									__t19 = Call_local_Main_f__467072791_3_16_24((x_4)+(int64(2)), (y_5)-(int64(1)))
								}
							end_branch_19:
								return __t19
							}
						}
						f_3_17_25 = gopurs_runtime.Func(func(x_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(y_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Int(Call_local_Main_f_3_17_25(x_4_loop_val.IntVal, y_5_loop_val.IntVal))
							})
						})
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{Call_local_Main_f__467072791_3_16_24(int64(0), int64(100000)), int64(200000)}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
							var f__467072791_4_20_26 gopurs_runtime.Value
							_ = f__467072791_4_20_26
							var f__467072791_4_20_26_cell *gopurs_runtime.Value
							_ = f__467072791_4_20_26_cell
							// FALLBACK TCO: isLoop=false len=2
							var f_4_21_27 gopurs_runtime.Value
							_ = f_4_21_27
							var f_4_21_27_cell *gopurs_runtime.Value
							_ = f_4_21_27_cell
							// FALLBACK TCO: isLoop=false len=2
							f__467072791_4_20_26 = gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t22 gopurs_runtime.Value
								{
									if (x_5.IntVal) > (int64(1000)) {
										__t22 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Int((x_5.IntVal) + (v_6.IntVal))
										})
										goto end_branch_22
									} else {

									}
								}
								{
									__t22 = gopurs_runtime.Func(func(y_prime__6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Int(gopurs_runtime.Apply2((*f__467072791_4_20_26_cell), gopurs_runtime.Int((x_5.IntVal)+(int64(10))), gopurs_runtime.Int((y_prime__6.IntVal)-(int64(1)))).IntVal)
									})
								}
							end_branch_22:
								return __t22
							})
							f__467072791_4_20_26_cell = &f__467072791_4_20_26
							f_4_21_27 = gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t23 gopurs_runtime.Value
								{
									if (x_5.IntVal) > (int64(1000)) {
										__t23 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Int((x_5.IntVal) + (v_6.IntVal))
										})
										goto end_branch_23
									} else {

									}
								}
								{
									__t23 = gopurs_runtime.Func(func(y_prime__6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Int(gopurs_runtime.Apply2((*f__467072791_4_20_26_cell), gopurs_runtime.Int((x_5.IntVal)+(int64(10))), gopurs_runtime.Int((y_prime__6.IntVal)-(int64(1)))).IntVal)
									})
								}
							end_branch_23:
								return __t23
							})
							f_4_21_27_cell = &f_4_21_27
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
								actual   int64
								expected int64
							}{gopurs_runtime.Apply2(f__467072791_4_20_26, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(100))).IntVal, int64(1009)}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
								var Call_local_Main_f__467072791_5_24_28 func(int64, int64) int64
								_ = Call_local_Main_f__467072791_5_24_28
								var f__467072791_5_24_28 gopurs_runtime.Value
								_ = f__467072791_5_24_28
								var Call_local_Main_f_5_25_29 func(int64, int64) int64
								_ = Call_local_Main_f_5_25_29
								var f_5_25_29 gopurs_runtime.Value
								_ = f_5_25_29
								Call_local_Main_f__467072791_5_24_28 = func(x_6_loop int64, y_7_loop int64) int64 {
								f__467072791_5_24_28:
									for {
										if false {
											continue f__467072791_5_24_28
										}
										var x_6 int64 = x_6_loop
										_ = x_6
										var y_7 int64 = y_7_loop
										_ = y_7
										var __t26 int64
										{
											if (y_7) <= (int64(0)) {
												__t26 = x_6
												goto end_branch_26
											} else {

											}
										}
										{
											x_6_loop = (x_6) + (int64(2))
											y_7_loop = (y_7) - (int64(1))
											continue f__467072791_5_24_28
											__t26 = func() int64 { panic("unreachable") }()
										}
									end_branch_26:
										return __t26
									}
								}
								f__467072791_5_24_28 = gopurs_runtime.Func(func(x_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(y_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Int(Call_local_Main_f__467072791_5_24_28(x_6_loop_val.IntVal, y_7_loop_val.IntVal))
									})
								})
								Call_local_Main_f_5_25_29 = func(x_6_loop int64, y_7_loop int64) int64 {
								f_5_25_29:
									for {
										if false {
											continue f_5_25_29
										}
										var x_6 int64 = x_6_loop
										_ = x_6
										var y_7 int64 = y_7_loop
										_ = y_7
										var __t27 int64
										{
											if (y_7) <= (int64(0)) {
												__t27 = x_6
												goto end_branch_27
											} else {

											}
										}
										{
											__t27 = Call_local_Main_f__467072791_5_24_28((x_6)+(int64(2)), (y_7)-(int64(1)))
										}
									end_branch_27:
										return __t27
									}
								}
								f_5_25_29 = gopurs_runtime.Func(func(x_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(y_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Int(Call_local_Main_f_5_25_29(x_6_loop_val.IntVal, y_7_loop_val.IntVal))
									})
								})
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
									actual   int64
									expected int64
								}{Call_local_Main_f__467072791_5_24_28(int64(0), int64(100)), int64(200)}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
									var f__467072791_6_28_30 gopurs_runtime.Value
									_ = f__467072791_6_28_30
									var f__467072791_6_28_30_cell *gopurs_runtime.Value
									_ = f__467072791_6_28_30_cell
									// FALLBACK TCO: isLoop=false len=2
									var f_6_29_31 gopurs_runtime.Value
									_ = f_6_29_31
									var f_6_29_31_cell *gopurs_runtime.Value
									_ = f_6_29_31_cell
									// FALLBACK TCO: isLoop=false len=2
									f__467072791_6_28_30 = gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): g__3466805691_9_30 shape=App(Other) bindingType=(Func [Int] Int)
										g__3466805691_9_30 := gopurs_runtime.Apply((*f__467072791_6_28_30_cell), gopurs_runtime.Int((x_7.IntVal)+(int64(2))))
										_ = g__3466805691_9_30
										var __t31 int64
										{
											if (y_8.IntVal) <= (int64(0)) {
												__t31 = x_7.IntVal
												goto end_branch_31
											} else {

											}
										}
										{
											__t31 = gopurs_runtime.Apply(g__3466805691_9_30, gopurs_runtime.Int((y_8.IntVal)-(int64(1)))).IntVal
										}
									end_branch_31:
										return gopurs_runtime.Int(__t31)
									})
									f__467072791_6_28_30_cell = &f__467072791_6_28_30
									f_6_29_31 = gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): g__3466805691_9_32 shape=App(Other) bindingType=(Func [Int] Int)
										g__3466805691_9_32 := gopurs_runtime.Apply((*f__467072791_6_28_30_cell), gopurs_runtime.Int((x_7.IntVal)+(int64(2))))
										_ = g__3466805691_9_32
										var __t33 int64
										{
											if (y_8.IntVal) <= (int64(0)) {
												__t33 = x_7.IntVal
												goto end_branch_33
											} else {

											}
										}
										{
											__t33 = gopurs_runtime.Apply(g__3466805691_9_32, gopurs_runtime.Int((y_8.IntVal)-(int64(1)))).IntVal
										}
									end_branch_33:
										return gopurs_runtime.Int(__t33)
									})
									f_6_29_31_cell = &f_6_29_31
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
										actual   int64
										expected int64
									}{gopurs_runtime.Apply2(f__467072791_6_28_30, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(100))).IntVal, int64(200)}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
										var Call_local_Main_f__467072791_7_34_32 func(int64, int64) int64
										_ = Call_local_Main_f__467072791_7_34_32
										var f__467072791_7_34_32 gopurs_runtime.Value
										_ = f__467072791_7_34_32
										var Call_local_Main_f_7_35_33 func(int64, int64) int64
										_ = Call_local_Main_f_7_35_33
										var f_7_35_33 gopurs_runtime.Value
										_ = f_7_35_33
										Call_local_Main_f__467072791_7_34_32 = func(x_8_loop int64, y_9_loop int64) int64 {
										f__467072791_7_34_32:
											for {
												if false {
													continue f__467072791_7_34_32
												}
												var x_8 int64 = x_8_loop
												_ = x_8
												var y_9 int64 = y_9_loop
												_ = y_9
												var __t36 int64
												{
													if (y_9) <= (int64(0)) {
														__t36 = x_8
														goto end_branch_36
													} else {

													}
												}
												{
													x_8_loop = (x_8) + (int64(2))
													y_9_loop = (y_9) - (int64(1))
													continue f__467072791_7_34_32
													__t36 = func() int64 { panic("unreachable") }()
												}
											end_branch_36:
												return __t36
											}
										}
										f__467072791_7_34_32 = gopurs_runtime.Func(func(x_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(y_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Int(Call_local_Main_f__467072791_7_34_32(x_8_loop_val.IntVal, y_9_loop_val.IntVal))
											})
										})
										Call_local_Main_f_7_35_33 = func(x_8_loop int64, y_9_loop int64) int64 {
										f_7_35_33:
											for {
												if false {
													continue f_7_35_33
												}
												var x_8 int64 = x_8_loop
												_ = x_8
												var y_9 int64 = y_9_loop
												_ = y_9
												var __t37 int64
												{
													if (y_9) <= (int64(0)) {
														__t37 = x_8
														goto end_branch_37
													} else {

													}
												}
												{
													__t37 = Call_local_Main_f__467072791_7_34_32((x_8)+(int64(2)), (y_9)-(int64(1)))
												}
											end_branch_37:
												return __t37
											}
										}
										f_7_35_33 = gopurs_runtime.Func(func(x_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(y_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Int(Call_local_Main_f_7_35_33(x_8_loop_val.IntVal, y_9_loop_val.IntVal))
											})
										})
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
											actual   int64
											expected int64
										}{Call_local_Main_f__467072791_7_34_32(int64(0), int64(100)), int64(200)}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
										}))
									}))
								}))
							}))
						}))
					}))
				}))
			}))
		}()
	})
	return cache_Main_main
}

func Call_Main_tco3(y0_0_loop int64) int64 {
	var y0_0 int64 = y0_0_loop
	_ = y0_0
	var Call_local_Main_f__467072791_1_0_2 func(int64, int64) int64
	_ = Call_local_Main_f__467072791_1_0_2
	var f__467072791_1_0_2 gopurs_runtime.Value
	_ = f__467072791_1_0_2
	var Call_local_Main_f_1_1_3 func(int64, int64) int64
	_ = Call_local_Main_f_1_1_3
	var f_1_1_3 gopurs_runtime.Value
	_ = f_1_1_3
	Call_local_Main_f__467072791_1_0_2 = func(x_2_loop int64, y_3_loop int64) int64 {
	f__467072791_1_0_2:
		for {
			if false {
				continue f__467072791_1_0_2
			}
			var x_2 int64 = x_2_loop
			_ = x_2
			var y_3 int64 = y_3_loop
			_ = y_3
			var Call_local_Main_g__467072791_4_2_4 func(int64, int64) int64
			_ = Call_local_Main_g__467072791_4_2_4
			var g__467072791_4_2_4 gopurs_runtime.Value
			_ = g__467072791_4_2_4
			var Call_local_Main_g_4_3_5 func(int64, int64) int64
			_ = Call_local_Main_g_4_3_5
			var g_4_3_5 gopurs_runtime.Value
			_ = g_4_3_5
			Call_local_Main_g__467072791_4_2_4 = func(x_prime__5_loop int64, y_prime__6_loop int64) int64 {
			g__467072791_4_2_4:
				for {
					if false {
						continue g__467072791_4_2_4
					}
					var x_prime__5 int64 = x_prime__5_loop
					_ = x_prime__5
					var y_prime__6 int64 = y_prime__6_loop
					_ = y_prime__6
					var __t5 int64
					{
						if (y_prime__6) <= (int64(0)) {
							__t5 = x_prime__5
							goto end_branch_5
						} else {

						}
					}
					{
						var __t4 int64
						{
							if (y_prime__6) > (gopurs_runtime.IntDiv(y0_0, int64(2))) {
								x_prime__5_loop = (x_prime__5) + (int64(3))
								y_prime__6_loop = (y_prime__6) - (int64(1))
								continue g__467072791_4_2_4
								__t4 = func() int64 { panic("unreachable") }()
								goto end_branch_4
							} else {

							}
						}
						{
							__t4 = Call_local_Main_f__467072791_1_0_2((x_prime__5)+(int64(2)), y_prime__6)
						}
					end_branch_4:
						__t5 = __t4
					}
				end_branch_5:
					return __t5
				}
			}
			g__467072791_4_2_4 = gopurs_runtime.Func(func(x_prime__5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_prime__6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_g__467072791_4_2_4(x_prime__5_loop_val.IntVal, y_prime__6_loop_val.IntVal))
				})
			})
			Call_local_Main_g_4_3_5 = func(x_prime__5_loop int64, y_prime__6_loop int64) int64 {
			g_4_3_5:
				for {
					if false {
						continue g_4_3_5
					}
					var x_prime__5 int64 = x_prime__5_loop
					_ = x_prime__5
					var y_prime__6 int64 = y_prime__6_loop
					_ = y_prime__6
					var __t7 int64
					{
						if (y_prime__6) <= (int64(0)) {
							__t7 = x_prime__5
							goto end_branch_7
						} else {

						}
					}
					{
						var __t6 int64
						{
							if (y_prime__6) > (gopurs_runtime.IntDiv(y0_0, int64(2))) {
								__t6 = Call_local_Main_g__467072791_4_2_4((x_prime__5)+(int64(3)), (y_prime__6)-(int64(1)))
								goto end_branch_6
							} else {

							}
						}
						{
							__t6 = Call_local_Main_f__467072791_1_0_2((x_prime__5)+(int64(2)), y_prime__6)
						}
					end_branch_6:
						__t7 = __t6
					}
				end_branch_7:
					return __t7
				}
			}
			g_4_3_5 = gopurs_runtime.Func(func(x_prime__5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_prime__6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_g_4_3_5(x_prime__5_loop_val.IntVal, y_prime__6_loop_val.IntVal))
				})
			})
			return Call_local_Main_g__467072791_4_2_4(x_2, (y_3)-(int64(1)))
		}
	}
	f__467072791_1_0_2 = gopurs_runtime.Func(func(x_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_f__467072791_1_0_2(x_2_loop_val.IntVal, y_3_loop_val.IntVal))
		})
	})
	Call_local_Main_f_1_1_3 = func(x_2_loop int64, y_3_loop int64) int64 {
	f_1_1_3:
		for {
			if false {
				continue f_1_1_3
			}
			var x_2 int64 = x_2_loop
			_ = x_2
			var y_3 int64 = y_3_loop
			_ = y_3
			var Call_local_Main_g__467072791_4_8_6 func(int64, int64) int64
			_ = Call_local_Main_g__467072791_4_8_6
			var g__467072791_4_8_6 gopurs_runtime.Value
			_ = g__467072791_4_8_6
			var Call_local_Main_g_4_9_7 func(int64, int64) int64
			_ = Call_local_Main_g_4_9_7
			var g_4_9_7 gopurs_runtime.Value
			_ = g_4_9_7
			Call_local_Main_g__467072791_4_8_6 = func(x_prime__5_loop int64, y_prime__6_loop int64) int64 {
			g__467072791_4_8_6:
				for {
					if false {
						continue g__467072791_4_8_6
					}
					var x_prime__5 int64 = x_prime__5_loop
					_ = x_prime__5
					var y_prime__6 int64 = y_prime__6_loop
					_ = y_prime__6
					var __t11 int64
					{
						if (y_prime__6) <= (int64(0)) {
							__t11 = x_prime__5
							goto end_branch_11
						} else {

						}
					}
					{
						var __t10 int64
						{
							if (y_prime__6) > (gopurs_runtime.IntDiv(y0_0, int64(2))) {
								x_prime__5_loop = (x_prime__5) + (int64(3))
								y_prime__6_loop = (y_prime__6) - (int64(1))
								continue g__467072791_4_8_6
								__t10 = func() int64 { panic("unreachable") }()
								goto end_branch_10
							} else {

							}
						}
						{
							__t10 = Call_local_Main_f__467072791_1_0_2((x_prime__5)+(int64(2)), y_prime__6)
						}
					end_branch_10:
						__t11 = __t10
					}
				end_branch_11:
					return __t11
				}
			}
			g__467072791_4_8_6 = gopurs_runtime.Func(func(x_prime__5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_prime__6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_g__467072791_4_8_6(x_prime__5_loop_val.IntVal, y_prime__6_loop_val.IntVal))
				})
			})
			Call_local_Main_g_4_9_7 = func(x_prime__5_loop int64, y_prime__6_loop int64) int64 {
			g_4_9_7:
				for {
					if false {
						continue g_4_9_7
					}
					var x_prime__5 int64 = x_prime__5_loop
					_ = x_prime__5
					var y_prime__6 int64 = y_prime__6_loop
					_ = y_prime__6
					var __t13 int64
					{
						if (y_prime__6) <= (int64(0)) {
							__t13 = x_prime__5
							goto end_branch_13
						} else {

						}
					}
					{
						var __t12 int64
						{
							if (y_prime__6) > (gopurs_runtime.IntDiv(y0_0, int64(2))) {
								__t12 = Call_local_Main_g__467072791_4_8_6((x_prime__5)+(int64(3)), (y_prime__6)-(int64(1)))
								goto end_branch_12
							} else {

							}
						}
						{
							__t12 = Call_local_Main_f__467072791_1_0_2((x_prime__5)+(int64(2)), y_prime__6)
						}
					end_branch_12:
						__t13 = __t12
					}
				end_branch_13:
					return __t13
				}
			}
			g_4_9_7 = gopurs_runtime.Func(func(x_prime__5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_prime__6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_g_4_9_7(x_prime__5_loop_val.IntVal, y_prime__6_loop_val.IntVal))
				})
			})
			return Call_local_Main_g__467072791_4_8_6(x_2, (y_3)-(int64(1)))
		}
	}
	f_1_1_3 = gopurs_runtime.Func(func(x_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_f_1_1_3(x_2_loop_val.IntVal, y_3_loop_val.IntVal))
		})
	})
	return Call_local_Main_f__467072791_1_0_2(int64(0), y0_0)
}

func Call_Main_ntco1(y0_0_loop int64) int64 {
	var y0_0 int64 = y0_0_loop
	_ = y0_0
	var f__467072791_1_0_18 gopurs_runtime.Value
	_ = f__467072791_1_0_18
	var f__467072791_1_0_18_cell *gopurs_runtime.Value
	_ = f__467072791_1_0_18_cell
	// FALLBACK TCO: isLoop=false len=2
	var f_1_1_19 gopurs_runtime.Value
	_ = f_1_1_19
	var f_1_1_19_cell *gopurs_runtime.Value
	_ = f_1_1_19_cell
	// FALLBACK TCO: isLoop=false len=2
	f__467072791_1_0_18 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 gopurs_runtime.Value
		{
			if (x_2.IntVal) > ((int64(10)) * (y0_0)) {
				__t2 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((x_2.IntVal) + (v_3.IntVal))
				})
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.Func(func(y_prime__3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply2((*f__467072791_1_0_18_cell), gopurs_runtime.Int((x_2.IntVal)+(int64(10))), gopurs_runtime.Int((y_prime__3.IntVal)-(int64(1)))).IntVal)
			})
		}
	end_branch_2:
		return __t2
	})
	f__467072791_1_0_18_cell = &f__467072791_1_0_18
	f_1_1_19 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 gopurs_runtime.Value
		{
			if (x_2.IntVal) > ((int64(10)) * (y0_0)) {
				__t3 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((x_2.IntVal) + (v_3.IntVal))
				})
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Func(func(y_prime__3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply2((*f__467072791_1_0_18_cell), gopurs_runtime.Int((x_2.IntVal)+(int64(10))), gopurs_runtime.Int((y_prime__3.IntVal)-(int64(1)))).IntVal)
			})
		}
	end_branch_3:
		return __t3
	})
	f_1_1_19_cell = &f_1_1_19
	return gopurs_runtime.Apply2(f__467072791_1_0_18, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(y0_0)).IntVal
}
