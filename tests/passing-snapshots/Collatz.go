package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_void gopurs_runtime.Value
var once_Main_void sync.Once

func Get_Main_void() gopurs_runtime.Value {
	once_Main_void.Do(func() {
		cache_Main_void = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST()))
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
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(Call_Main_collatz(int64(1000)))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_collatz(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(n_0))
		_ = __local_var_1_0
		r_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = r_2_1
		count_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(int64(0))), gopurs_runtime.Value{})
		_ = count_3_2
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST())), gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_4_3 := (*(r_2_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = __local_var_4_3
			return gopurs_runtime.Bool(((__local_var_4_3.IntVal) == (int64(1))) != (true))
		}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_4 shape=App(Var) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope2), Int])
			__local_var_4_4 := gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_modifyImpl(), gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): s_prime__5_5 shape=Other bindingType=(TypeVar a$scope52)
				s_prime__5_5 := gopurs_runtime.Int((s_4.IntVal) + (int64(1)))
				_ = s_prime__5_5
				return func() gopurs_runtime.Value {
					orig := struct {
						state gopurs_runtime.Value
						value gopurs_runtime.Value
					}{s_prime__5_5, s_prime__5_5}
					_ = orig
					return gopurs_runtime.RecordDict2("state", "value", orig.state, orig.value)
				}()
			}), count_3_2)
			_ = __local_var_4_4
			_dollar___unused_5_6 := gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Value{})
			_ = _dollar___unused_5_6
			m_6_7 := (*(r_2_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			_ = m_6_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				var __t9 int64
				{
					if (gopurs_runtime.IntMod(m_6_7.IntVal, int64(2))) == (int64(0)) {
						__t9 = gopurs_runtime.IntDiv(m_6_7.IntVal, int64(2))
						goto end_branch_9
					} else {

					}
				}
				{
					__t9 = ((int64(3)) * (m_6_7.IntVal)) + (int64(1))
				}
			end_branch_9:
				// TAST (Let): __local_var_7_8 shape=Branch(Other, def=Other) bindingType=Any
				__local_var_7_8 := __t9
				_ = __local_var_7_8
				*(r_2_1.PtrVal().(*interface{})) = gopurs_runtime.Int(__local_var_7_8)
				return gopurs_runtime.Int(__local_var_7_8)
			})), gopurs_runtime.Value{})
		})), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return (*(count_3_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
			})
		})), gopurs_runtime.Value{})
	})).IntVal
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
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
