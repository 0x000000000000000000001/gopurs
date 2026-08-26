package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_testIndentation gopurs_runtime.Value
var once_Main_testIndentation sync.Once

func Get_Main_testIndentation() gopurs_runtime.Value {
	once_Main_testIndentation.Do(func() {
		cache_Main_testIndentation = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_testIndentation(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_testIndentation
}

var cache_Main_min gopurs_runtime.Value
var once_Main_min sync.Once

func Get_Main_min() gopurs_runtime.Value {
	once_Main_min.Do(func() {
		cache_Main_min = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_min(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), n_1_box, m_2_box)
		})
	})
	return cache_Main_min
}

var cache_Main_min__2767602680 gopurs_runtime.Value
var once_Main_min__2767602680 sync.Once

func Get_Main_min__2767602680() gopurs_runtime.Value {
	once_Main_min__2767602680.Do(func() {
		cache_Main_min__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_min__2767602680(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), n_1_box, m_2_box)
		})
	})
	return cache_Main_min__2767602680
}

var cache_Main_max gopurs_runtime.Value
var once_Main_max sync.Once

func Get_Main_max() gopurs_runtime.Value {
	once_Main_max.Do(func() {
		cache_Main_max = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_max(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), n_1_box, m_2_box)
		})
	})
	return cache_Main_max
}

var cache_Main_max__2767602680 gopurs_runtime.Value
var once_Main_max__2767602680 sync.Once

func Get_Main_max__2767602680() gopurs_runtime.Value {
	once_Main_max__2767602680.Do(func() {
		cache_Main_max__2767602680 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_max__2767602680(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), n_1_box, m_2_box)
		})
	})
	return cache_Main_max__2767602680
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("Done"), gopurs_runtime.Str("ZZZZ"))
				if uint32(__t_tag_0.IntVal) == 1527465420 {
					__t1 = gopurs_runtime.Str("Done")
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Str("ZZZZ")
			}
		end_branch_1:
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t1.StrVal()))
		}()
	})
	return cache_Main_main
}

var cache_Main_collatz2 gopurs_runtime.Value
var once_Main_collatz2 sync.Once

func Get_Main_collatz2() gopurs_runtime.Value {
	once_Main_collatz2.Do(func() {
		cache_Main_collatz2 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_collatz2(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_collatz2
}

var cache_Main_collatz gopurs_runtime.Value
var once_Main_collatz sync.Once

func Get_Main_collatz() gopurs_runtime.Value {
	once_Main_collatz.Do(func() {
		cache_Main_collatz = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_collatz(x_0_box.FloatVal()))
		})
	})
	return cache_Main_collatz
}

var cache_Main_clunky_case2 gopurs_runtime.Value
var once_Main_clunky_case2 sync.Once

func Get_Main_clunky_case2() gopurs_runtime.Value {
	once_Main_clunky_case2.Do(func() {
		cache_Main_clunky_case2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_clunky_case2(a_0_box.IntVal, b_1_box.IntVal))
		})
	})
	return cache_Main_clunky_case2
}

var cache_Main_clunky_case1 gopurs_runtime.Value
var once_Main_clunky_case1 sync.Once

func Get_Main_clunky_case1() gopurs_runtime.Value {
	once_Main_clunky_case1.Do(func() {
		cache_Main_clunky_case1 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_clunky_case1(a_0_box.IntVal, b_1_box.IntVal))
		})
	})
	return cache_Main_clunky_case1
}

var cache_Main_clunky2 gopurs_runtime.Value
var once_Main_clunky2 sync.Once

func Get_Main_clunky2() gopurs_runtime.Value {
	once_Main_clunky2.Do(func() {
		cache_Main_clunky2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_clunky2(a_0_box.IntVal, b_1_box.IntVal))
		})
	})
	return cache_Main_clunky2
}

var cache_Main_clunky1_refutable gopurs_runtime.Value
var once_Main_clunky1_refutable sync.Once

func Get_Main_clunky1_refutable() gopurs_runtime.Value {
	once_Main_clunky1_refutable.Do(func() {
		cache_Main_clunky1_refutable = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_clunky1_refutable(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_clunky1_refutable
}

var cache_Main_clunky1 gopurs_runtime.Value
var once_Main_clunky1 sync.Once

func Get_Main_clunky1() gopurs_runtime.Value {
	once_Main_clunky1.Do(func() {
		cache_Main_clunky1 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_clunky1(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_clunky1
}

func Call_Main_testIndentation(x_0_loop float64, y_1_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	var __t1 float64
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(x_0), gopurs_runtime.Float(0.0))
		if uint32(__t_tag_0.IntVal) == 380165415 {
			__t1 = (x_0) + (y_1)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = (y_1) - (x_0)
	}
end_branch_1:
	return __t1
}

func Call_Main_min(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], n_1_loop gopurs_runtime.Value, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var n_1 gopurs_runtime.Value = n_1_loop
	_ = n_1
	var m_2 gopurs_runtime.Value = m_2_loop
	_ = m_2
	var __t1 gopurs_runtime.Value
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), n_1, m_2)
		if uint32(__t_tag_0.IntVal) == 1527465420 {
			__t1 = n_1
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = m_2
	}
end_branch_1:
	return __t1
}

func Call_Main_min__2767602680(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], n_1_loop gopurs_runtime.Value, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var n_1 gopurs_runtime.Value = n_1_loop
	_ = n_1
	var m_2 gopurs_runtime.Value = m_2_loop
	_ = m_2
	var __t1 gopurs_runtime.Value
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), n_1, m_2)
		if uint32(__t_tag_0.IntVal) == 1527465420 {
			__t1 = n_1
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = m_2
	}
end_branch_1:
	return __t1
}

func Call_Main_max(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], n_1_loop gopurs_runtime.Value, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var n_1 gopurs_runtime.Value = n_1_loop
	_ = n_1
	var m_2 gopurs_runtime.Value = m_2_loop
	_ = m_2
	var __t1 gopurs_runtime.Value
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), m_2, n_1)
		if uint32(__t_tag_0.IntVal) == 1527465420 {
			__t1 = n_1
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = m_2
	}
end_branch_1:
	return __t1
}

func Call_Main_max__2767602680(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], n_1_loop gopurs_runtime.Value, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var n_1 gopurs_runtime.Value = n_1_loop
	_ = n_1
	var m_2 gopurs_runtime.Value = m_2_loop
	_ = m_2
	var __t1 gopurs_runtime.Value
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), m_2, n_1)
		if uint32(__t_tag_0.IntVal) == 1527465420 {
			__t1 = n_1
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = m_2
	}
end_branch_1:
	return __t1
}

func Call_Main_collatz2(x_0_loop float64, y_1_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	var __t1 float64
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(y_1), gopurs_runtime.Float(0.0))
		if uint32(__t_tag_0.IntVal) == 380165415 {
			__t1 = (x_0) / (2.0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = ((x_0) * (3.0)) + (1.0)
	}
end_branch_1:
	return __t1
}

func Call_Main_collatz(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return (x_0) / (2.0)
}

func Call_Main_clunky_case2(a_0_loop int64, b_1_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	var b_1 int64 = b_1_loop
	_ = b_1
	var __t1 gopurs_runtime.Value
	{
		if (b_1) < (a_0) {
			__t1 = gopurs_runtime.Int(a_0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.Int(b_1)
	}
end_branch_1:
	// TAST (Let): __local_var_2_0 shape=Branch(Other, def=Other) bindingType=Int
	__local_var_2_0 := __t1.IntVal
	_ = __local_var_2_0
	var __t2 int64
	{
		if (__local_var_2_0) > (5) {
			__t2 = __local_var_2_0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = (a_0) + (b_1)
	}
end_branch_2:
	return __t2
}

func Call_Main_clunky_case1(a_0_loop int64, b_1_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	var b_1 int64 = b_1_loop
	_ = b_1
	var __t1 gopurs_runtime.Value
	{
		if (b_1) < (a_0) {
			__t1 = gopurs_runtime.Int(a_0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.Int(b_1)
	}
end_branch_1:
	// TAST (Let): __local_var_2_0 shape=Branch(Other, def=Other) bindingType=Int
	__local_var_2_0 := __t1.IntVal
	_ = __local_var_2_0
	var __t2 int64
	{
		if (__local_var_2_0) > (5) {
			__t2 = __local_var_2_0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = (a_0) + (b_1)
	}
end_branch_2:
	return __t2
}

func Call_Main_clunky2(a_0_loop int64, b_1_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	var b_1 int64 = b_1_loop
	_ = b_1
	var __t1 gopurs_runtime.Value
	{
		if (b_1) < (a_0) {
			__t1 = gopurs_runtime.Int(a_0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.Int(b_1)
	}
end_branch_1:
	// TAST (Let): __local_var_2_0 shape=Branch(Other, def=Other) bindingType=Int
	__local_var_2_0 := __t1.IntVal
	_ = __local_var_2_0
	var __t2 int64
	{
		if (__local_var_2_0) > (5) {
			__t2 = __local_var_2_0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = (a_0) + (b_1)
	}
end_branch_2:
	return __t2
}

func Call_Main_clunky1_refutable(v_0_loop int64, v1_1_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	var v1_1 int64 = v1_1_loop
	_ = v1_1
	var __t3 int64
	{
		if (v_0) == (0) {
			var __t1 gopurs_runtime.Value
			{
				if (v1_1) < (v1_1) {
					__t1 = gopurs_runtime.Int(v1_1)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Int(v1_1)
			}
		end_branch_1:
			// TAST (Let): __local_var_2_0 shape=Branch(Other, def=Other) bindingType=Int
			__local_var_2_0 := __t1.IntVal
			_ = __local_var_2_0
			var __t2 int64
			{
				if (__local_var_2_0) > (5) {
					__t2 = __local_var_2_0
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = v_0
			}
		end_branch_2:
			__t3 = __t2
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = v_0
	}
end_branch_3:
	return __t3
}

func Call_Main_clunky1(v_0_loop int64, v1_1_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	var v1_1 int64 = v1_1_loop
	_ = v1_1
	var __t1 gopurs_runtime.Value
	{
		if (v1_1) < (v_0) {
			__t1 = gopurs_runtime.Int(v_0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.Int(v1_1)
	}
end_branch_1:
	// TAST (Let): __local_var_2_0 shape=Branch(Other, def=Other) bindingType=Int
	__local_var_2_0 := __t1.IntVal
	_ = __local_var_2_0
	var __t2 int64
	{
		if (__local_var_2_0) > (5) {
			__t2 = __local_var_2_0
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = v_0
	}
end_branch_2:
	return __t2
}
