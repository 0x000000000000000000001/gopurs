package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_negate gopurs_runtime.Value
var once_Main_negate sync.Once

func Get_Main_negate() gopurs_runtime.Value {
	once_Main_negate.Do(func() {
		cache_Main_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_negate(a_0_box.FloatVal()))
		})
	})
	return cache_Main_negate
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Float(7.0)
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Float(1.0)
	})
	return cache_Main_test2
}

var cache_Main_test19 gopurs_runtime.Value
var once_Main_test19 sync.Once

func Get_Main_test19() gopurs_runtime.Value {
	once_Main_test19.Do(func() {
		cache_Main_test19 = gopurs_runtime.Float(gopurs_runtime.Float(-1.0).FloatVal())
	})
	return cache_Main_test19
}

var cache_Main_test18 gopurs_runtime.Value
var once_Main_test18 sync.Once

func Get_Main_test18() gopurs_runtime.Value {
	once_Main_test18.Do(func() {
		cache_Main_test18 = gopurs_runtime.Float(gopurs_runtime.Float(1.0).FloatVal())
	})
	return cache_Main_test18
}

var cache_Main_test17 gopurs_runtime.Value
var once_Main_test17 sync.Once

func Get_Main_test17() gopurs_runtime.Value {
	once_Main_test17.Do(func() {
		cache_Main_test17 = gopurs_runtime.Float(gopurs_runtime.Float(1.0).FloatVal())
	})
	return cache_Main_test17
}

var cache_Main_test14 gopurs_runtime.Value
var once_Main_test14 sync.Once

func Get_Main_test14() gopurs_runtime.Value {
	once_Main_test14.Do(func() {
		cache_Main_test14 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_test14(a_0_box.FloatVal(), b_1_box.FloatVal()))
		})
	})
	return cache_Main_test14
}

var cache_Main_test15 gopurs_runtime.Value
var once_Main_test15 sync.Once

func Get_Main_test15() gopurs_runtime.Value {
	once_Main_test15.Do(func() {
		cache_Main_test15 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_test15(a_0_box.FloatVal(), b_1_box.FloatVal()))
		})
	})
	return cache_Main_test15
}

var cache_Main_test10 gopurs_runtime.Value
var once_Main_test10 sync.Once

func Get_Main_test10() gopurs_runtime.Value {
	once_Main_test10.Do(func() {
		cache_Main_test10 = gopurs_runtime.Str("Hello")
	})
	return cache_Main_test10
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func4(func(dictSemiring_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_0_box), x_1_box, y_2_box, z_3_box)
		})
	})
	return cache_Main_test1
}

var cache_Main_op5 gopurs_runtime.Value
var once_Main_op5 sync.Once

func Get_Main_op5() gopurs_runtime.Value {
	once_Main_op5.Do(func() {
		cache_Main_op5 = gopurs_runtime.Func2(func(as_0_box gopurs_runtime.Value, bs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_op5(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(as_0_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}(), func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(bs_1_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_op5
}

var cache_Main_test11 gopurs_runtime.Value
var once_Main_test11 sync.Once

func Get_Main_test11() gopurs_runtime.Value {
	once_Main_test11.Do(func() {
		cache_Main_test11 = func() gopurs_runtime.Value {
			arr := func() []float64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []float64{1.0, 2.0, 0.0}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Float(v)
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()).UnsafePtr)
				unboxed := make([]float64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.FloatVal()
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Float(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
	return cache_Main_test11
}

var cache_Main_op4 gopurs_runtime.Value
var once_Main_op4 sync.Once

func Get_Main_op4() gopurs_runtime.Value {
	once_Main_op4.Do(func() {
		cache_Main_op4 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_op4(f_0_box, x_1_box)
		})
	})
	return cache_Main_op4
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Str("Hello World")
	})
	return cache_Main_test8
}

var cache_Main_test9 gopurs_runtime.Value
var once_Main_test9 sync.Once

func Get_Main_test9() gopurs_runtime.Value {
	once_Main_test9.Do(func() {
		cache_Main_test9 = gopurs_runtime.Str("Hello World")
	})
	return cache_Main_test9
}

var cache_Main_op3 gopurs_runtime.Value
var once_Main_op3 sync.Once

func Get_Main_op3() gopurs_runtime.Value {
	once_Main_op3.Do(func() {
		cache_Main_op3 = gopurs_runtime.Func2(func(s1_0_box gopurs_runtime.Value, s2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_op3(s1_0_box.StrVal(), s2_1_box.StrVal()))
		})
	})
	return cache_Main_op3
}

var cache_Main_test7 gopurs_runtime.Value
var once_Main_test7 sync.Once

func Get_Main_test7() gopurs_runtime.Value {
	once_Main_test7.Do(func() {
		cache_Main_test7 = gopurs_runtime.Str(("Hello") + ("World!"))
	})
	return cache_Main_test7
}

var cache_Main_op2 gopurs_runtime.Value
var once_Main_op2 sync.Once

func Get_Main_op2() gopurs_runtime.Value {
	once_Main_op2.Do(func() {
		cache_Main_op2 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_op2(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_op2
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Float(15.0)
	})
	return cache_Main_test5
}

var cache_Main_op1 gopurs_runtime.Value
var once_Main_op1 sync.Once

func Get_Main_op1() gopurs_runtime.Value {
	once_Main_op1.Do(func() {
		cache_Main_op1 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_op1(x_0_box, v_1_box)
		})
	})
	return cache_Main_op1
}

var cache_Main_k gopurs_runtime.Value
var once_Main_k sync.Once

func Get_Main_k() gopurs_runtime.Value {
	once_Main_k.Do(func() {
		cache_Main_k = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_k(x_0_box, y_1_box)
		})
	})
	return cache_Main_k
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)
	})
	return cache_Main_test4
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Float(3.0)
	})
	return cache_Main_test6
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_op4__369910300 gopurs_runtime.Value
var once_Main_op4__369910300 sync.Once

func Get_Main_op4__369910300() gopurs_runtime.Value {
	once_Main_op4__369910300.Do(func() {
		cache_Main_op4__369910300 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_op4__369910300(f_0_box, x_1_box)
		})
	})
	return cache_Main_op4__369910300
}

var cache_Main_op5__39834991 gopurs_runtime.Value
var once_Main_op5__39834991 sync.Once

func Get_Main_op5__39834991() gopurs_runtime.Value {
	once_Main_op5__39834991.Do(func() {
		cache_Main_op5__39834991 = gopurs_runtime.Func2(func(as_0_box gopurs_runtime.Value, bs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_op5__39834991(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(as_0_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}(), func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(bs_1_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_op5__39834991
}

var cache_Main_test1__638868095 gopurs_runtime.Value
var once_Main_test1__638868095 sync.Once

func Get_Main_test1__638868095() gopurs_runtime.Value {
	once_Main_test1__638868095.Do(func() {
		cache_Main_test1__638868095 = gopurs_runtime.Func4(func(dictSemiring_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test1__638868095(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_0_box), x_1_box, y_2_box, z_3_box)
		})
	})
	return cache_Main_test1__638868095
}

func Call_Main_negate(a_0_loop float64) float64 {
	var a_0 float64 = a_0_loop
	_ = a_0
	return gopurs_runtime.Float(-(gopurs_runtime.Float(a_0).FloatVal())).FloatVal()
}

func Call_Main_test14(a_0_loop float64, b_1_loop float64) bool {
	var a_0 float64 = a_0_loop
	_ = a_0
	var b_1 float64 = b_1_loop
	_ = b_1
	var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(a_0), gopurs_runtime.Float(b_1))
	return (uint32(__t_tag_0.IntVal) == 1527465420)
}

func Call_Main_test15(a_0_loop float64, b_1_loop float64) bool {
	var a_0 float64 = a_0_loop
	_ = a_0
	var b_1 float64 = b_1_loop
	_ = b_1
	return (gopurs_runtime.Bool(false).IntVal) != (0)
}

func Call_Main_test1(dictSemiring_0_loop *Constructor_Data_Semiring_Semiring, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 *Constructor_Data_Semiring_Semiring = dictSemiring_0_loop
	_ = dictSemiring_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var y_2 gopurs_runtime.Value = y_2_loop
	_ = y_2
	var z_3 gopurs_runtime.Value = z_3_loop
	_ = z_3
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_0.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_0.V1), x_1, y_2), gopurs_runtime.Apply2(z_3, x_1, y_2))
}

func Call_Main_op5(as_0_loop []gopurs_runtime.Value, bs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var as_0 []gopurs_runtime.Value = as_0_loop
	_ = as_0
	var bs_1 []gopurs_runtime.Value = bs_1_loop
	_ = bs_1
	return as_0
}

func Call_Main_op4(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(f_0, x_1)
}

func Call_Main_op3(s1_0_loop string, s2_1_loop string) string {
	var s1_0 string = s1_0_loop
	_ = s1_0
	var s2_1 string = s2_1_loop
	_ = s2_1
	return (s1_0) + (s2_1)
}

func Call_Main_op2(x_0_loop float64, y_1_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	return ((x_0) * (y_1)) + (y_1)
}

func Call_Main_op1(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return x_0
}

func Call_Main_k(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var y_1 gopurs_runtime.Value = y_1_loop
	_ = y_1
	return x_0
}

func Call_Main_op4__369910300(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(f_0, x_1)
}

func Call_Main_op5__39834991(as_0_loop []gopurs_runtime.Value, bs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var as_0 []gopurs_runtime.Value = as_0_loop
	_ = as_0
	var bs_1 []gopurs_runtime.Value = bs_1_loop
	_ = bs_1
	return as_0
}

func Call_Main_test1__638868095(dictSemiring_0_loop *Constructor_Data_Semiring_Semiring, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 *Constructor_Data_Semiring_Semiring = dictSemiring_0_loop
	_ = dictSemiring_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var y_2 gopurs_runtime.Value = y_2_loop
	_ = y_2
	var z_3 gopurs_runtime.Value = z_3_loop
	_ = z_3
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_0.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_0.V1), x_1, y_2), gopurs_runtime.Apply2(z_3, x_1, y_2))
}
