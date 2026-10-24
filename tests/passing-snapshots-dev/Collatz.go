package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(Call_Main_collatz(int64(1000))))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_collatz(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_ref_1 := gopurs_runtime.Int(n_0)
		_ = __local_ref_1
		var __local_iface_1 interface{} = __local_ref_1
		r_1_0 := gopurs_runtime.Any(&__local_iface_1)
		_ = r_1_0
		__local_ref_3 := gopurs_runtime.Int(int64(0))
		_ = __local_ref_3
		var __local_iface_3 interface{} = __local_ref_3
		count_2_2 := gopurs_runtime.Any(&__local_iface_3)
		_ = count_2_2
		_dollar___unused_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_3_5 := (*(r_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = __local_var_3_5
			return gopurs_runtime.Bool(((__local_var_3_5.IntVal) == (int64(1))) != (true))
		}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_3_7 := (*(count_2_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = __local_var_3_7
			*(count_2_2.PtrVal().(*interface{})) = gopurs_runtime.Int(gopurs_runtime.IntAdd(__local_var_3_7.IntVal, int64(1)))
			_dollar___unused_3_6 := gopurs_runtime.Int(gopurs_runtime.IntAdd(__local_var_3_7.IntVal, int64(1)))
			_ = _dollar___unused_3_6
			m_4_8 := (*(r_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = m_4_8
			var __t10 int64
			{
				if (gopurs_runtime.IntMod(m_4_8.IntVal, int64(2))) == (int64(0)) {
					__t10 = gopurs_runtime.IntDiv(m_4_8.IntVal, int64(2))
					goto end_branch_10
				} else {

				}
			}
			{
				__t10 = gopurs_runtime.IntAdd(gopurs_runtime.IntMul(int64(3), m_4_8.IntVal), int64(1))
			}
		end_branch_10:
			// TAST (Let): __local_var_5_9 shape=Branch(Other, def=Other) bindingType=Any
			__local_var_5_9 := __t10
			_ = __local_var_5_9
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				*(r_1_0.PtrVal().(*interface{})) = gopurs_runtime.Int(__local_var_5_9)
				return gopurs_runtime.Int(__local_var_5_9)
			})), gopurs_runtime.Value{})
		})), gopurs_runtime.Value{})
		_ = _dollar___unused_3_4
		return (*(count_2_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
	})).IntVal
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
