package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_applyN gopurs_runtime.Value
var once_Main_applyN sync.Once

func Get_Main_applyN() gopurs_runtime.Value {
	once_Main_applyN.Do(func() {
		cache_Main_applyN = func() gopurs_runtime.Value {
			var Call_local_Main_go__go_0_0_0 func(gopurs_runtime.Value, int64, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_0_0
			var go__go_0_0_0 gopurs_runtime.Value
			_ = go__go_0_0_0
			Call_local_Main_go__go_0_0_0 = func(v_1_loop gopurs_runtime.Value, v1_2_loop int64, v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_0_0_0:
				for {
					if false {
						continue go__go_0_0_0
					}
					var v_1 gopurs_runtime.Value = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var v2_3 gopurs_runtime.Value = v2_3_loop
					_ = v2_3
					var __t1 gopurs_runtime.Value
					{
						if (v1_2) <= (int64(0)) {
							__t1 = v_1
							goto end_branch_1
						} else {

						}
					}
					{
						v_1_loop = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), v_1, v2_3)
						v1_2_loop = gopurs_runtime.IntSub(v1_2, int64(1))
						v2_3_loop = v2_3
						continue go__go_0_0_0
						__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			go__go_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_local_Main_go__go_0_0_0(v_1_loop_val, v1_2_loop_val.IntVal, v2_3_loop_val)
					})
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}))
		}()
	})
	return cache_Main_applyN
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Main_applyN(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_0.IntVal, int64(1)))
			}), gopurs_runtime.Int(int64(0))).IntVal))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Main_applyN(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_2.IntVal, int64(1)))
			}), gopurs_runtime.Int(int64(0))).IntVal)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Main_applyN(), gopurs_runtime.Int(int64(2)), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_3.IntVal, int64(1)))
			}), gopurs_runtime.Int(int64(0))).IntVal)), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Main_applyN(), gopurs_runtime.Int(int64(3)), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_4.IntVal, int64(1)))
			}), gopurs_runtime.Int(int64(0))).IntVal)), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Main_applyN(), gopurs_runtime.Int(int64(4)), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_5.IntVal, int64(1)))
			}), gopurs_runtime.Int(int64(0))).IntVal)), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(Call_Data_Array_span(gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
				arr := Call_Data_Array_go__range(int64(1), int64(10000))
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).UnsafePtr))).go__init))).IntVal)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply(Call_Control_Monad_Rec_Class_tailRec(gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t8 gopurs_runtime.Value
				{
					if (n_7.IntVal) < (int64(10000)) {
						__t8 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.IntAdd(n_7.IntVal, int64(1)))}))}
						goto end_branch_8
					} else {

					}
				}
				{
					__t8 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(42))}))}
				}
			end_branch_8:
				return __t8
			})), gopurs_runtime.Int(int64(0))).IntVal)), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
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
