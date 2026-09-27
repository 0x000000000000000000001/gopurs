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
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Main_f__467072791_0_1_20 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_0_1_20
			var f__467072791_0_1_20 gopurs_runtime.Value
			_ = f__467072791_0_1_20
			var Call_local_Main_f_0_2_21 func(int64, int64) int64
			_ = Call_local_Main_f_0_2_21
			var f_0_2_21 gopurs_runtime.Value
			_ = f_0_2_21
			Call_local_Main_f__467072791_0_1_20 = func(x_1_loop int64, y_2_loop int64) int64 {
			f__467072791_0_1_20:
				for {
					if false {
						continue f__467072791_0_1_20
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_3 shape=Other bindingType=Int
					__local_var_3_3 := (x_1) + (int64(2))
					_ = __local_var_3_3
					// TAST (Let): __local_var_4_4 shape=Other bindingType=Int
					__local_var_4_4 := (y_2) - (int64(1))
					_ = __local_var_4_4
					var __t5 int64
					{
						if (__local_var_4_4) <= (int64(0)) {
							__t5 = __local_var_3_3
							goto end_branch_5
						} else {

						}
					}
					{
						x_1_loop = __local_var_3_3
						y_2_loop = __local_var_4_4
						continue f__467072791_0_1_20
						__t5 = func() int64 { panic("unreachable") }()
					}
				end_branch_5:
					return __t5
				}
			}
			f__467072791_0_1_20 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_0_1_20(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			Call_local_Main_f_0_2_21 = func(x_1_loop int64, y_2_loop int64) int64 {
			f_0_2_21:
				for {
					if false {
						continue f_0_2_21
					}
					var x_1 int64 = x_1_loop
					_ = x_1
					var y_2 int64 = y_2_loop
					_ = y_2
					// TAST (Let): __local_var_3_6 shape=Other bindingType=Int
					__local_var_3_6 := (x_1) + (int64(2))
					_ = __local_var_3_6
					// TAST (Let): __local_var_4_7 shape=Other bindingType=Int
					__local_var_4_7 := (y_2) - (int64(1))
					_ = __local_var_4_7
					var __t8 int64
					{
						if (__local_var_4_7) <= (int64(0)) {
							__t8 = __local_var_3_6
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = Call_local_Main_f__467072791_0_1_20(__local_var_3_6, __local_var_4_7)
					}
				end_branch_8:
					return __t8
				}
			}
			f_0_2_21 = gopurs_runtime.Func(func(x_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_0_2_21(x_1_loop_val.IntVal, y_2_loop_val.IntVal))
				})
			})
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_f__467072791_0_1_20(int64(0), int64(100000)), int64(200000)})
			_ = __local_var_0_0
			__local_var_1_9 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_9
			var Call_local_Main_f__467072791_2_11_22 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_2_11_22
			var f__467072791_2_11_22 gopurs_runtime.Value
			_ = f__467072791_2_11_22
			var Call_local_Main_f_2_12_23 func(int64, int64) int64
			_ = Call_local_Main_f_2_12_23
			var f_2_12_23 gopurs_runtime.Value
			_ = f_2_12_23
			Call_local_Main_f__467072791_2_11_22 = func(x_3_loop int64, y_4_loop int64) int64 {
			f__467072791_2_11_22:
				for {
					if false {
						continue f__467072791_2_11_22
					}
					var x_3 int64 = x_3_loop
					_ = x_3
					var y_4 int64 = y_4_loop
					_ = y_4
					// TAST (Let): __local_var_5_13 shape=Other bindingType=Int
					__local_var_5_13 := (x_3) + (int64(2))
					_ = __local_var_5_13
					// TAST (Let): __local_var_6_14 shape=Other bindingType=Int
					__local_var_6_14 := (y_4) - (int64(1))
					_ = __local_var_6_14
					var __t15 int64
					{
						if (__local_var_6_14) <= (int64(0)) {
							__t15 = __local_var_5_13
							goto end_branch_15
						} else {

						}
					}
					{
						x_3_loop = __local_var_5_13
						y_4_loop = __local_var_6_14
						continue f__467072791_2_11_22
						__t15 = func() int64 { panic("unreachable") }()
					}
				end_branch_15:
					return __t15
				}
			}
			f__467072791_2_11_22 = gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_2_11_22(x_3_loop_val.IntVal, y_4_loop_val.IntVal))
				})
			})
			Call_local_Main_f_2_12_23 = func(x_3_loop int64, y_4_loop int64) int64 {
			f_2_12_23:
				for {
					if false {
						continue f_2_12_23
					}
					var x_3 int64 = x_3_loop
					_ = x_3
					var y_4 int64 = y_4_loop
					_ = y_4
					// TAST (Let): __local_var_5_16 shape=Other bindingType=Int
					__local_var_5_16 := (x_3) + (int64(2))
					_ = __local_var_5_16
					// TAST (Let): __local_var_6_17 shape=Other bindingType=Int
					__local_var_6_17 := (y_4) - (int64(1))
					_ = __local_var_6_17
					var __t18 int64
					{
						if (__local_var_6_17) <= (int64(0)) {
							__t18 = __local_var_5_16
							goto end_branch_18
						} else {

						}
					}
					{
						__t18 = Call_local_Main_f__467072791_2_11_22(__local_var_5_16, __local_var_6_17)
					}
				end_branch_18:
					return __t18
				}
			}
			f_2_12_23 = gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_2_12_23(x_3_loop_val.IntVal, y_4_loop_val.IntVal))
				})
			})
			__local_var_2_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_f__467072791_2_11_22(int64(0), int64(100000)), int64(200000)}), gopurs_runtime.Value{})
			_ = __local_var_2_10
			__local_var_3_19 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_tco3(int64(100000)), int64(249997)}), gopurs_runtime.Value{})
			_ = __local_var_3_19
			var Call_local_Main_f__467072791_4_21_24 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_4_21_24
			var f__467072791_4_21_24 gopurs_runtime.Value
			_ = f__467072791_4_21_24
			var Call_local_Main_f_4_22_25 func(int64, int64) int64
			_ = Call_local_Main_f_4_22_25
			var f_4_22_25 gopurs_runtime.Value
			_ = f_4_22_25
			Call_local_Main_f__467072791_4_21_24 = func(x_5_loop int64, y_6_loop int64) int64 {
			f__467072791_4_21_24:
				for {
					if false {
						continue f__467072791_4_21_24
					}
					var x_5 int64 = x_5_loop
					_ = x_5
					var y_6 int64 = y_6_loop
					_ = y_6
					var __t23 int64
					{
						if (y_6) <= (int64(0)) {
							__t23 = x_5
							goto end_branch_23
						} else {

						}
					}
					{
						x_5_loop = (x_5) + (int64(2))
						y_6_loop = (y_6) - (int64(1))
						continue f__467072791_4_21_24
						__t23 = func() int64 { panic("unreachable") }()
					}
				end_branch_23:
					return __t23
				}
			}
			f__467072791_4_21_24 = gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_4_21_24(x_5_loop_val.IntVal, y_6_loop_val.IntVal))
				})
			})
			Call_local_Main_f_4_22_25 = func(x_5_loop int64, y_6_loop int64) int64 {
			f_4_22_25:
				for {
					if false {
						continue f_4_22_25
					}
					var x_5 int64 = x_5_loop
					_ = x_5
					var y_6 int64 = y_6_loop
					_ = y_6
					var __t24 int64
					{
						if (y_6) <= (int64(0)) {
							__t24 = x_5
							goto end_branch_24
						} else {

						}
					}
					{
						__t24 = Call_local_Main_f__467072791_4_21_24((x_5)+(int64(2)), (y_6)-(int64(1)))
					}
				end_branch_24:
					return __t24
				}
			}
			f_4_22_25 = gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_4_22_25(x_5_loop_val.IntVal, y_6_loop_val.IntVal))
				})
			})
			__local_var_4_20 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_f__467072791_4_21_24(int64(0), int64(100000)), int64(200000)}), gopurs_runtime.Value{})
			_ = __local_var_4_20
			var f__467072791_5_26_26 gopurs_runtime.Value
			_ = f__467072791_5_26_26
			var f__467072791_5_26_26_cell *gopurs_runtime.Value
			_ = f__467072791_5_26_26_cell
			// FALLBACK TCO: isLoop=false len=2
			var f_5_27_27 gopurs_runtime.Value
			_ = f_5_27_27
			var f_5_27_27_cell *gopurs_runtime.Value
			_ = f_5_27_27_cell
			// FALLBACK TCO: isLoop=false len=2
			f__467072791_5_26_26 = gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t28 gopurs_runtime.Value
				{
					if (x_6.IntVal) > (int64(1000)) {
						__t28 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((x_6.IntVal) + (v_7.IntVal))
						})
						goto end_branch_28
					} else {

					}
				}
				{
					__t28 = gopurs_runtime.Func(func(y_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.Apply2((*f__467072791_5_26_26_cell), gopurs_runtime.Int((x_6.IntVal)+(int64(10))), gopurs_runtime.Int((y_prime__7.IntVal)-(int64(1)))).IntVal)
					})
				}
			end_branch_28:
				return __t28
			})
			f__467072791_5_26_26_cell = &f__467072791_5_26_26
			f_5_27_27 = gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t29 gopurs_runtime.Value
				{
					if (x_6.IntVal) > (int64(1000)) {
						__t29 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((x_6.IntVal) + (v_7.IntVal))
						})
						goto end_branch_29
					} else {

					}
				}
				{
					__t29 = gopurs_runtime.Func(func(y_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.Apply2((*f__467072791_5_26_26_cell), gopurs_runtime.Int((x_6.IntVal)+(int64(10))), gopurs_runtime.Int((y_prime__7.IntVal)-(int64(1)))).IntVal)
					})
				}
			end_branch_29:
				return __t29
			})
			f_5_27_27_cell = &f_5_27_27
			__local_var_5_25 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply2(f__467072791_5_26_26, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(100))).IntVal, int64(1009)}), gopurs_runtime.Value{})
			_ = __local_var_5_25
			var Call_local_Main_f__467072791_6_31_28 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_6_31_28
			var f__467072791_6_31_28 gopurs_runtime.Value
			_ = f__467072791_6_31_28
			var Call_local_Main_f_6_32_29 func(int64, int64) int64
			_ = Call_local_Main_f_6_32_29
			var f_6_32_29 gopurs_runtime.Value
			_ = f_6_32_29
			Call_local_Main_f__467072791_6_31_28 = func(x_7_loop int64, y_8_loop int64) int64 {
			f__467072791_6_31_28:
				for {
					if false {
						continue f__467072791_6_31_28
					}
					var x_7 int64 = x_7_loop
					_ = x_7
					var y_8 int64 = y_8_loop
					_ = y_8
					var __t33 int64
					{
						if (y_8) <= (int64(0)) {
							__t33 = x_7
							goto end_branch_33
						} else {

						}
					}
					{
						x_7_loop = (x_7) + (int64(2))
						y_8_loop = (y_8) - (int64(1))
						continue f__467072791_6_31_28
						__t33 = func() int64 { panic("unreachable") }()
					}
				end_branch_33:
					return __t33
				}
			}
			f__467072791_6_31_28 = gopurs_runtime.Func(func(x_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_6_31_28(x_7_loop_val.IntVal, y_8_loop_val.IntVal))
				})
			})
			Call_local_Main_f_6_32_29 = func(x_7_loop int64, y_8_loop int64) int64 {
			f_6_32_29:
				for {
					if false {
						continue f_6_32_29
					}
					var x_7 int64 = x_7_loop
					_ = x_7
					var y_8 int64 = y_8_loop
					_ = y_8
					var __t34 int64
					{
						if (y_8) <= (int64(0)) {
							__t34 = x_7
							goto end_branch_34
						} else {

						}
					}
					{
						__t34 = Call_local_Main_f__467072791_6_31_28((x_7)+(int64(2)), (y_8)-(int64(1)))
					}
				end_branch_34:
					return __t34
				}
			}
			f_6_32_29 = gopurs_runtime.Func(func(x_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_6_32_29(x_7_loop_val.IntVal, y_8_loop_val.IntVal))
				})
			})
			__local_var_6_30 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_f__467072791_6_31_28(int64(0), int64(100)), int64(200)}), gopurs_runtime.Value{})
			_ = __local_var_6_30
			var f__467072791_7_36_30 gopurs_runtime.Value
			_ = f__467072791_7_36_30
			var f__467072791_7_36_30_cell *gopurs_runtime.Value
			_ = f__467072791_7_36_30_cell
			// FALLBACK TCO: isLoop=false len=2
			var f_7_37_31 gopurs_runtime.Value
			_ = f_7_37_31
			var f_7_37_31_cell *gopurs_runtime.Value
			_ = f_7_37_31_cell
			// FALLBACK TCO: isLoop=false len=2
			f__467072791_7_36_30 = gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): g__3466805691_10_38 shape=App(Other) bindingType=(Func [Int] Int)
				g__3466805691_10_38 := gopurs_runtime.Apply((*f__467072791_7_36_30_cell), gopurs_runtime.Int((x_8.IntVal)+(int64(2))))
				_ = g__3466805691_10_38
				var __t39 int64
				{
					if (y_9.IntVal) <= (int64(0)) {
						__t39 = x_8.IntVal
						goto end_branch_39
					} else {

					}
				}
				{
					__t39 = gopurs_runtime.Apply(g__3466805691_10_38, gopurs_runtime.Int((y_9.IntVal)-(int64(1)))).IntVal
				}
			end_branch_39:
				return gopurs_runtime.Int(__t39)
			})
			f__467072791_7_36_30_cell = &f__467072791_7_36_30
			f_7_37_31 = gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): g__3466805691_10_40 shape=App(Other) bindingType=(Func [Int] Int)
				g__3466805691_10_40 := gopurs_runtime.Apply((*f__467072791_7_36_30_cell), gopurs_runtime.Int((x_8.IntVal)+(int64(2))))
				_ = g__3466805691_10_40
				var __t41 int64
				{
					if (y_9.IntVal) <= (int64(0)) {
						__t41 = x_8.IntVal
						goto end_branch_41
					} else {

					}
				}
				{
					__t41 = gopurs_runtime.Apply(g__3466805691_10_40, gopurs_runtime.Int((y_9.IntVal)-(int64(1)))).IntVal
				}
			end_branch_41:
				return gopurs_runtime.Int(__t41)
			})
			f_7_37_31_cell = &f_7_37_31
			__local_var_7_35 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply2(f__467072791_7_36_30, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(100))).IntVal, int64(200)}), gopurs_runtime.Value{})
			_ = __local_var_7_35
			var Call_local_Main_f__467072791_8_43_32 func(int64, int64) int64
			_ = Call_local_Main_f__467072791_8_43_32
			var f__467072791_8_43_32 gopurs_runtime.Value
			_ = f__467072791_8_43_32
			var Call_local_Main_f_8_44_33 func(int64, int64) int64
			_ = Call_local_Main_f_8_44_33
			var f_8_44_33 gopurs_runtime.Value
			_ = f_8_44_33
			Call_local_Main_f__467072791_8_43_32 = func(x_9_loop int64, y_10_loop int64) int64 {
			f__467072791_8_43_32:
				for {
					if false {
						continue f__467072791_8_43_32
					}
					var x_9 int64 = x_9_loop
					_ = x_9
					var y_10 int64 = y_10_loop
					_ = y_10
					var __t45 int64
					{
						if (y_10) <= (int64(0)) {
							__t45 = x_9
							goto end_branch_45
						} else {

						}
					}
					{
						x_9_loop = (x_9) + (int64(2))
						y_10_loop = (y_10) - (int64(1))
						continue f__467072791_8_43_32
						__t45 = func() int64 { panic("unreachable") }()
					}
				end_branch_45:
					return __t45
				}
			}
			f__467072791_8_43_32 = gopurs_runtime.Func(func(x_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f__467072791_8_43_32(x_9_loop_val.IntVal, y_10_loop_val.IntVal))
				})
			})
			Call_local_Main_f_8_44_33 = func(x_9_loop int64, y_10_loop int64) int64 {
			f_8_44_33:
				for {
					if false {
						continue f_8_44_33
					}
					var x_9 int64 = x_9_loop
					_ = x_9
					var y_10 int64 = y_10_loop
					_ = y_10
					var __t46 int64
					{
						if (y_10) <= (int64(0)) {
							__t46 = x_9
							goto end_branch_46
						} else {

						}
					}
					{
						__t46 = Call_local_Main_f__467072791_8_43_32((x_9)+(int64(2)), (y_10)-(int64(1)))
					}
				end_branch_46:
					return __t46
				}
			}
			f_8_44_33 = gopurs_runtime.Func(func(x_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_f_8_44_33(x_9_loop_val.IntVal, y_10_loop_val.IntVal))
				})
			})
			__local_var_8_42 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_f__467072791_8_43_32(int64(0), int64(100)), int64(200)}), gopurs_runtime.Value{})
			_ = __local_var_8_42
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
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
