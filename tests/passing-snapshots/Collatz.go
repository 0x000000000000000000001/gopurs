package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_void gopurs_runtime.Value
var once_Main_void sync.Once

func Get_Main_void() gopurs_runtime.Value {
	once_Main_void.Do(func() {
		cache_Main_void = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_void(__local_var_0_box)
		})
	})
	return cache_Main_void
}

var cache_Main_collatz gopurs_runtime.Value
var once_Main_collatz sync.Once

func Get_Main_collatz() gopurs_runtime.Value {
	once_Main_collatz.Do(func() {
		cache_Main_collatz = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_collatz(n_0_box.IntVal))
		})
	})
	return cache_Main_collatz
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_collatz(1000))).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_void(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Main_collatz(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(n_0))
		_ = __local_var_1_0
		r_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = r_2_1
		count_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0)), gopurs_runtime.Value{})
		_ = count_3_2
		_dollar___unused_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_4_4 := (*(r_2_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = __local_var_4_4
			return gopurs_runtime.Bool(((__local_var_4_4.IntVal) == (gopurs_runtime.Int(1).IntVal)) != (true))
		}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_4_5 := (*(count_3_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = __local_var_4_5
			*(count_3_2.PtrVal().(*interface{})) = gopurs_runtime.Int((__local_var_4_5.IntVal) + (1))
			_dollar___unused_5_6 := gopurs_runtime.Int((__local_var_4_5.IntVal) + (1))
			_ = _dollar___unused_5_6
			m_6_7 := (*(r_2_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = m_6_7
			var __t9 int64
			{
				if (gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(m_6_7.IntVal), gopurs_runtime.Int(2)).IntVal) == (0) {
					__t9 = (m_6_7.IntVal) / (2)
					goto end_branch_9
				} else {

				}
			}
			{
				__t9 = ((3) * (m_6_7.IntVal)) + (1)
			}
		end_branch_9:
			*(r_2_1.PtrVal().(*interface{})) = gopurs_runtime.Int(__t9)
			__local_var_7_8 := gopurs_runtime.Int(__t9)
			_ = __local_var_7_8
			return Get_Data_Unit_unit()
		})), gopurs_runtime.Value{})
		_ = _dollar___unused_4_3
		return (*(count_3_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
	})).IntVal
}
