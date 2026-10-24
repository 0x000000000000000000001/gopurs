package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_length gopurs_runtime.Value
var once_Main_length sync.Once

func Get_Main_length() gopurs_runtime.Value {
	once_Main_length.Do(func() {
		cache_Main_length = func() gopurs_runtime.Value {
			var Call_local_Main_go__go_0_0_0 func(int64, []gopurs_runtime.Value) int64
			_ = Call_local_Main_go__go_0_0_0
			var go__go_0_0_0 gopurs_runtime.Value
			_ = go__go_0_0_0
			Call_local_Main_go__go_0_0_0 = func(acc_1_loop int64, arr_2_loop []gopurs_runtime.Value) int64 {
			go__go_0_0_0:
				for {
					if false {
						continue go__go_0_0_0
					}
					var acc_1 int64 = acc_1_loop
					_ = acc_1
					var arr_2 []gopurs_runtime.Value = arr_2_loop
					_ = arr_2
					var __t1 int64
					{
						if Call_Data_Array_null(arr_2) {
							__t1 = acc_1
							goto end_branch_1
						} else {

						}
					}
					{
						acc_1_loop = gopurs_runtime.IntAdd(acc_1, int64(1))
						arr_2_loop = (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(Get_Data_Array_Partial_tail(), gopurs_runtime.Value{})
						}), gopurs_runtime.Array(arr_2))).UnsafePtr))
						continue go__go_0_0_0
						__t1 = func() int64 { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			go__go_0_0_0 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(arr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_go__go_0_0_0(acc_1_loop_val.IntVal, (*(*[]gopurs_runtime.Value)((arr_2_loop_val).UnsafePtr))))
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_0, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_length
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_length(), func() gopurs_runtime.Value {
				arr := Call_Data_Array_go__range(int64(1), int64(10000))
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
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
