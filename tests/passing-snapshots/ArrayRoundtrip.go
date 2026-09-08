package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_intAdd()
	})
	return cache_Main_add
}

var cache_Main_go__range gopurs_runtime.Value
var once_Main_go__range sync.Once

func Get_Main_go__range() gopurs_runtime.Value {
	once_Main_go__range.Do(func() {
		cache_Main_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_go__range(start_0_box.IntVal, end_1_box.IntVal)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_go__range
}

var cache_Main_filterEvens gopurs_runtime.Value
var once_Main_filterEvens sync.Once

func Get_Main_filterEvens() gopurs_runtime.Value {
	once_Main_filterEvens.Do(func() {
		cache_Main_filterEvens = gopurs_runtime.Func(func(arr_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_filterEvens(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(arr_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_filterEvens
}

var cache_Main_sumArrayEvens gopurs_runtime.Value
var once_Main_sumArrayEvens sync.Once

func Get_Main_sumArrayEvens() gopurs_runtime.Value {
	once_Main_sumArrayEvens.Do(func() {
		cache_Main_sumArrayEvens = gopurs_runtime.Func(func(values_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_sumArrayEvens(func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(values_0_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_sumArrayEvens
}

var cache_Main_sumEvens gopurs_runtime.Value
var once_Main_sumEvens sync.Once

func Get_Main_sumEvens() gopurs_runtime.Value {
	once_Main_sumEvens.Do(func() {
		cache_Main_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_sumEvens(n_0_box.IntVal))
		})
	})
	return cache_Main_sumEvens
}

var cache_Main_sumRangeEvens gopurs_runtime.Value
var once_Main_sumRangeEvens sync.Once

func Get_Main_sumRangeEvens() gopurs_runtime.Value {
	once_Main_sumRangeEvens.Do(func() {
		cache_Main_sumRangeEvens = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_sumRangeEvens(start_0_box.IntVal, end_1_box.IntVal))
		})
	})
	return cache_Main_sumRangeEvens
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, actual_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(label_0_box.StrVal(), expected_1_box.IntVal, actual_2_box.IntVal)
		})
	})
	return cache_Main_check
}

var cache_Main_checkArray gopurs_runtime.Value
var once_Main_checkArray sync.Once

func Get_Main_checkArray() gopurs_runtime.Value {
	once_Main_checkArray.Do(func() {
		cache_Main_checkArray = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, values_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkArray(label_0_box.StrVal(), expected_1_box.IntVal, func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(values_2_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}())
		})
	})
	return cache_Main_checkArray
}

var cache_Main_checkBenchmark gopurs_runtime.Value
var once_Main_checkBenchmark sync.Once

func Get_Main_checkBenchmark() gopurs_runtime.Value {
	once_Main_checkBenchmark.Do(func() {
		cache_Main_checkBenchmark = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, n_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkBenchmark(label_0_box.StrVal(), expected_1_box.IntVal, n_2_box.IntVal)
		})
	})
	return cache_Main_checkBenchmark
}

var cache_Main_checkRange gopurs_runtime.Value
var once_Main_checkRange sync.Once

func Get_Main_checkRange() gopurs_runtime.Value {
	once_Main_checkRange.Do(func() {
		cache_Main_checkRange = gopurs_runtime.Func4(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, start_2_box gopurs_runtime.Value, end_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkRange(label_0_box.StrVal(), expected_1_box.IntVal, start_2_box.IntVal, end_3_box.IntVal)
		})
	})
	return cache_Main_checkRange
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Main_checkArray("empty", int64(0), []int64{})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Main_checkArray("singleton odd", int64(0), []int64{int64(7)}), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Main_checkArray("singleton even", int64(8), []int64{int64(8)}), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Main_checkArray("singleton zero", int64(0), []int64{int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Main_checkArray("even only", int64(12), []int64{int64(2), int64(4), int64(6)}), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Main_checkArray("odd only", int64(0), []int64{int64(1), int64(3), int64(5)}), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_checkArray("mixed parity", int64(12), []int64{int64(1), int64(2), int64(3), int64(4), int64(5), int64(6)}), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_checkArray("negative values", int64(-6), []int64{int64(-5), int64(-4), int64(-3), int64(-2), int64(-1)}), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_checkArray("mixed signs", int64(-2), []int64{int64(-5), int64(-4), int64(-1), int64(0), int64(2), int64(7)}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_checkArray("duplicates", int64(-4), []int64{int64(2), int64(-4), int64(2), int64(-4)}), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_checkArray("minimum int", int64(-2147483648), []int64{int64(-2147483648)}), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_checkArray("maximum int is odd", int64(0), []int64{int64(2147483647)}), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_checkArray("maximum even int", int64(2147483646), []int64{int64(2147483646)}), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Main_checkArray("bounds with cancellation", int64(0), []int64{int64(2147483646), int64(-2147483648), int64(2)}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Main_checkRange("ascending range", int64(12), int64(1), int64(6)), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Main_checkRange("descending range", int64(12), int64(6), int64(1)), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Main_checkRange("negative ascending range", int64(-6), int64(-5), int64(-1)), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Main_checkRange("negative descending range", int64(-6), int64(-1), int64(-5)), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_checkRange("range crossing zero", int64(0), int64(-3), int64(3)), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Main_checkRange("minimum boundary range", int64(-2147483648), int64(-2147483648), int64(-2147483647)), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Main_checkRange("minimum boundary reversed", int64(-2147483648), int64(-2147483647), int64(-2147483648)), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(Call_Main_checkRange("maximum boundary range", int64(2147483646), int64(2147483646), int64(2147483647)), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(Call_Main_checkRange("maximum boundary reversed", int64(2147483646), int64(2147483647), int64(2147483646)), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(Call_Main_checkBenchmark("benchmark n=0 descends", int64(0), int64(0)), gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply(Call_Main_checkBenchmark("benchmark n=1 filters to empty", int64(0), int64(1)), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply(Call_Main_checkBenchmark("benchmark n=2", int64(2), int64(2)), gopurs_runtime.Value{})
			_ = __local_var_26_26
			__local_var_27_27 := gopurs_runtime.Apply(Call_Main_checkBenchmark("benchmark negative n", int64(-6), int64(-4)), gopurs_runtime.Value{})
			_ = __local_var_27_27
			__local_var_28_28 := gopurs_runtime.Apply(Call_Main_checkBenchmark("benchmark n=900", int64(202950), int64(900)), gopurs_runtime.Value{})
			_ = __local_var_28_28
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_go__range(start_0_loop int64, end_1_loop int64) []int64 {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Main_filterEvens(arr_0_loop []int64) []int64 {
	var arr_0 []int64 = arr_0_loop
	_ = arr_0
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr_val_filterImpl0 := func() gopurs_runtime.Value {
				arr := arr_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
			_ = arr_val_filterImpl0
			_ = arr_val_filterImpl0
			arr_go_filterImpl0 := (*[]gopurs_runtime.Value)(arr_val_filterImpl0.UnsafePtr)
			_ = arr_go_filterImpl0
			res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
			_ = res_go_filterImpl0
			for _, v_filterImpl0 := range *arr_go_filterImpl0 {
				if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool(((x_1.IntVal) % (int64(2))) == (int64(0)))
				}), v_filterImpl0).BoolVal() {
					res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
				} else {

				}
			}
			return res_go_filterImpl0
		}()).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Main_sumArrayEvens(values_0_loop []int64) int64 {
	var values_0 []int64 = values_0_loop
	_ = values_0
	return func() gopurs_runtime.Value {
		arr_val_foldlArray0 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr_val_filterImpl1 := func() gopurs_runtime.Value {
				arr := values_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
			_ = arr_val_filterImpl1
			_ = arr_val_filterImpl1
			arr_go_filterImpl1 := (*[]gopurs_runtime.Value)(arr_val_filterImpl1.UnsafePtr)
			_ = arr_go_filterImpl1
			res_go_filterImpl1 := make([]gopurs_runtime.Value, 0)
			_ = res_go_filterImpl1
			for _, v_filterImpl1 := range *arr_go_filterImpl1 {
				if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool(((x_1.IntVal) % (int64(2))) == (int64(0)))
				}), v_filterImpl1).BoolVal() {
					res_go_filterImpl1 = append(res_go_filterImpl1, v_filterImpl1)
				} else {

				}
			}
			return res_go_filterImpl1
		}())
		_ = arr_val_foldlArray0
		res_go_foldlArray0 := gopurs_runtime.Int(int64(0))
		_ = res_go_foldlArray0
		arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
		_ = arr_go_foldlArray0
		for _, v_foldlArray0 := range *arr_go_foldlArray0 {
			res_go_foldlArray0 = gopurs_runtime.Apply2(Get_Data_Semiring_intAdd(), res_go_foldlArray0, v_foldlArray0)
		}
		return res_go_foldlArray0
	}().IntVal
}

func Call_Main_sumEvens(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	return func() gopurs_runtime.Value {
		arr_val_foldlArray0 := func() gopurs_runtime.Value {
			source_int_array_foldlArray0_0 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr_val_filterImpl1 := gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(n_0))
				_ = arr_val_filterImpl1
				_ = arr_val_filterImpl1
				arr_go_filterImpl1 := (*[]gopurs_runtime.Value)(arr_val_filterImpl1.UnsafePtr)
				_ = arr_go_filterImpl1
				res_go_filterImpl1 := make([]gopurs_runtime.Value, 0)
				_ = res_go_filterImpl1
				for _, v_filterImpl1 := range *arr_go_filterImpl1 {
					if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((x_1.IntVal) % (int64(2))) == (int64(0)))
					}), v_filterImpl1).BoolVal() {
						res_go_filterImpl1 = append(res_go_filterImpl1, v_filterImpl1)
					} else {

					}
				}
				return res_go_filterImpl1
			}())
			_ = source_int_array_foldlArray0_0
			items_int_array_foldlArray0_0 := (*[]gopurs_runtime.Value)(source_int_array_foldlArray0_0.UnsafePtr)
			_ = items_int_array_foldlArray0_0
			for i_int_array_foldlArray0_0, v_int_array_foldlArray0_0 := range *items_int_array_foldlArray0_0 {
				(*items_int_array_foldlArray0_0)[i_int_array_foldlArray0_0] = gopurs_runtime.Int(v_int_array_foldlArray0_0.IntVal)
			}
			return source_int_array_foldlArray0_0
		}()
		_ = arr_val_foldlArray0
		res_go_foldlArray0 := gopurs_runtime.Int(int64(0))
		_ = res_go_foldlArray0
		arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
		_ = arr_go_foldlArray0
		for _, v_foldlArray0 := range *arr_go_foldlArray0 {
			res_go_foldlArray0 = gopurs_runtime.Apply2(Get_Data_Semiring_intAdd(), res_go_foldlArray0, v_foldlArray0)
		}
		return res_go_foldlArray0
	}().IntVal
}

func Call_Main_sumRangeEvens(start_0_loop int64, end_1_loop int64) int64 {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	return func() gopurs_runtime.Value {
		arr_val_foldlArray0 := func() gopurs_runtime.Value {
			source_int_array_foldlArray0_0 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr_val_filterImpl1 := gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1))
				_ = arr_val_filterImpl1
				_ = arr_val_filterImpl1
				arr_go_filterImpl1 := (*[]gopurs_runtime.Value)(arr_val_filterImpl1.UnsafePtr)
				_ = arr_go_filterImpl1
				res_go_filterImpl1 := make([]gopurs_runtime.Value, 0)
				_ = res_go_filterImpl1
				for _, v_filterImpl1 := range *arr_go_filterImpl1 {
					if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((x_2.IntVal) % (int64(2))) == (int64(0)))
					}), v_filterImpl1).BoolVal() {
						res_go_filterImpl1 = append(res_go_filterImpl1, v_filterImpl1)
					} else {

					}
				}
				return res_go_filterImpl1
			}())
			_ = source_int_array_foldlArray0_0
			items_int_array_foldlArray0_0 := (*[]gopurs_runtime.Value)(source_int_array_foldlArray0_0.UnsafePtr)
			_ = items_int_array_foldlArray0_0
			for i_int_array_foldlArray0_0, v_int_array_foldlArray0_0 := range *items_int_array_foldlArray0_0 {
				(*items_int_array_foldlArray0_0)[i_int_array_foldlArray0_0] = gopurs_runtime.Int(v_int_array_foldlArray0_0.IntVal)
			}
			return source_int_array_foldlArray0_0
		}()
		_ = arr_val_foldlArray0
		res_go_foldlArray0 := gopurs_runtime.Int(int64(0))
		_ = res_go_foldlArray0
		arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
		_ = arr_go_foldlArray0
		for _, v_foldlArray0 := range *arr_go_foldlArray0 {
			res_go_foldlArray0 = gopurs_runtime.Apply2(Get_Data_Semiring_intAdd(), res_go_foldlArray0, v_foldlArray0)
		}
		return res_go_foldlArray0
	}().IntVal
}

func Call_Main_check(label_0_loop string, expected_1_loop int64, actual_2_loop int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 int64 = expected_1_loop
	_ = expected_1
	var actual_2 int64 = actual_2_loop
	_ = actual_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(actual_2), gopurs_runtime.Int(expected_1)))
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(((label_0)+(": "))+(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(actual_2)).StrVal()))), gopurs_runtime.Value{})
	})
}

func Call_Main_checkArray(label_0_loop string, expected_1_loop int64, values_2_loop []int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 int64 = expected_1_loop
	_ = expected_1
	var values_2 []int64 = values_2_loop
	_ = values_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			arr := values_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
		_ = __local_var_5_2
		return gopurs_runtime.Apply(Call_Main_check(label_0, expected_1, Call_Main_sumArrayEvens(func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(__local_var_5_2.UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}())), gopurs_runtime.Value{})
	})
}

func Call_Main_checkBenchmark(label_0_loop string, expected_1_loop int64, n_2_loop int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 int64 = expected_1_loop
	_ = expected_1
	var n_2 int64 = n_2_loop
	_ = n_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(n_2))
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
		_ = __local_var_5_2
		return gopurs_runtime.Apply(Call_Main_check(label_0, expected_1, Call_Main_sumEvens(__local_var_5_2.IntVal)), gopurs_runtime.Value{})
	})
}

func Call_Main_checkRange(label_0_loop string, expected_1_loop int64, start_2_loop int64, end_3_loop int64) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 int64 = expected_1_loop
	_ = expected_1
	var start_2 int64 = start_2_loop
	_ = start_2
	var end_3 int64 = end_3_loop
	_ = end_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=Any
		__local_var_4_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			orig := struct {
				end   int64
				start int64
			}{end_3, start_2}
			_ = orig
			return gopurs_runtime.RecordDict2("end", "start", gopurs_runtime.Int(orig.end), gopurs_runtime.Int(orig.start))
		}())
		_ = __local_var_4_0
		inputRef_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
		_ = inputRef_5_1
		input_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), inputRef_5_1), gopurs_runtime.Value{})
		_ = input_6_2
		return gopurs_runtime.Apply(Call_Main_check(label_0, expected_1, Call_Main_sumRangeEvens(gopurs_runtime.RecordGet(input_6_2, "start").IntVal, gopurs_runtime.RecordGet(input_6_2, "end").IntVal)), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
