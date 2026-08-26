package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_numAdd()
	})
	return cache_Main_add
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test8(x_0_box.FloatVal()))
		})
	})
	return cache_Main_test8
}

var cache_Main_test7 gopurs_runtime.Value
var once_Main_test7 sync.Once

func Get_Main_test7() gopurs_runtime.Value {
	once_Main_test7.Do(func() {
		cache_Main_test7 = gopurs_runtime.Float(gopurs_runtime.Float(1.0).FloatVal())
	})
	return cache_Main_test7
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = func() gopurs_runtime.Value {
			var g_0_0_1 gopurs_runtime.Value
			_ = g_0_0_1
			var f_0_1_2 gopurs_runtime.Value
			_ = f_0_1_2
			g_0_0_1 = gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float((gopurs_runtime.Apply(f_0_1_2, gopurs_runtime.Float((x_1.FloatVal())-(1.0))).FloatVal()) + (1.0))
			})
			f_0_1_2 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 float64
				{
					var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(v_1.FloatVal()), gopurs_runtime.Float(0.0))
					if uint32(__t_tag_2.IntVal) == 380165415 {
						__t3 = (gopurs_runtime.Apply(g_0_0_1, gopurs_runtime.Float((v_1.FloatVal())/(2.0))).FloatVal()) + (1.0)
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = 0.0
				}
			end_branch_3:
				return gopurs_runtime.Float(__t3)
			})
			return gopurs_runtime.Float(gopurs_runtime.Apply(f_0_1_2, gopurs_runtime.Float(10.0)).FloatVal())
		}()
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func(func(dictPartial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test4(dictPartial_0_box)
		})
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Float(6.0)
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test2(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_test2
}

var cache_Main_test10 gopurs_runtime.Value
var once_Main_test10 sync.Once

func Get_Main_test10() gopurs_runtime.Value {
	once_Main_test10.Do(func() {
		cache_Main_test10 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test10(v_0_box))
		})
	})
	return cache_Main_test10
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test1(x_0_box.FloatVal()))
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(2.0)).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(5.0)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(6.0)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(3.0)).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(Get_Main_test5().FloatVal())).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.Float(1.0).FloatVal())).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(Call_Main_test8(100.0))).StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_test8(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	var Call_local_Main_go__go_1_0_0 func(float64) gopurs_runtime.Value
	_ = Call_local_Main_go__go_1_0_0
	var go__go_1_0_0 gopurs_runtime.Value
	_ = go__go_1_0_0
	Call_local_Main_go__go_1_0_0 = func(v_2_loop float64) gopurs_runtime.Value {
	go__go_1_0_0:
		for {
			if false {
				continue go__go_1_0_0
			}
			var v_2 float64 = v_2_loop
			_ = v_2
			var __t4 float64
			{
				var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((x_0)-(0.1)), gopurs_runtime.Float((v_2)*(v_2)))
				var __t_and_3 bool = false
				if uint32(__t_tag_1.IntVal) == 1527465420 {

					var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((v_2)*(v_2)), gopurs_runtime.Float((x_0)+(0.1)))
					__t_and_3 = (uint32(__t_tag_2.IntVal) == 1527465420)
				}
				if __t_and_3 {
					__t4 = v_2
					goto end_branch_4
				} else {

				}
			}
			{
				v_2_loop = ((v_2) + ((x_0) / (v_2))) / (2.0)
				continue go__go_1_0_0
				__t4 = gopurs_runtime.Value{}.FloatVal()
			}
		end_branch_4:
			return gopurs_runtime.Float(__t4)
		}
	}
	go__go_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Main_go__go_1_0_0(v_2_loop_val.FloatVal())
	})
	return Call_local_Main_go__go_1_0_0(x_0).FloatVal()
}

func Call_Main_test4(dictPartial_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPartial_0 gopurs_runtime.Value = dictPartial_0_loop
	_ = dictPartial_0
	return gopurs_runtime.Float(3.0)
}

func Call_Main_test2(x_0_loop float64, y_1_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	return (((x_0) + (1.0)) + (y_1)) + (1.0)
}

func Call_Main_test10(v_0_loop gopurs_runtime.Value) float64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var g_1_0_3 gopurs_runtime.Value
	_ = g_1_0_3
	var f_1_1_4 gopurs_runtime.Value
	_ = f_1_1_4
	g_1_0_3 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply(f_1_1_4, gopurs_runtime.Float(x_2.FloatVal())).FloatVal()) / (2.0))
	})
	f_1_1_4 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply(g_1_0_3, gopurs_runtime.Float(x_2.FloatVal())).FloatVal()) * (3.0))
	})
	return gopurs_runtime.Apply(f_1_1_4, gopurs_runtime.Float(10.0)).FloatVal()
}

func Call_Main_test1(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return (x_0) + (1.0)
}
