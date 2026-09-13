package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_length gopurs_runtime.Value
var once_Main_length sync.Once

func Get_Main_length() gopurs_runtime.Value {
	once_Main_length.Do(func() {
		cache_Main_length = func() gopurs_runtime.Value {
			var Call_local_Main_go__1086705844_0_0_0 func(int64, []gopurs_runtime.Value) int64
			_ = Call_local_Main_go__1086705844_0_0_0
			var go__1086705844_0_0_0 gopurs_runtime.Value
			_ = go__1086705844_0_0_0
			var Call_local_Main_go__go_0_1_1 func(int64, []gopurs_runtime.Value) int64
			_ = Call_local_Main_go__go_0_1_1
			var go__go_0_1_1 gopurs_runtime.Value
			_ = go__go_0_1_1
			Call_local_Main_go__1086705844_0_0_0 = func(acc_1_loop int64, arr_2_loop []gopurs_runtime.Value) int64 {
			go__1086705844_0_0_0:
				for {
					if false {
						continue go__1086705844_0_0_0
					}
					var acc_1 int64 = acc_1_loop
					_ = acc_1
					var arr_2 []gopurs_runtime.Value = arr_2_loop
					_ = arr_2
					var __t2 int64
					{
						if (gopurs_runtime.Int(int64(len(arr_2))).IntVal) == (int64(0)) {
							__t2 = acc_1
							goto end_branch_2
						} else {

						}
					}
					{
						acc_1_loop = (acc_1) + (int64(1))
						arr_2_loop = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(len(arr_2))), gopurs_runtime.Array(arr_2)).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
						continue go__1086705844_0_0_0
						__t2 = func() int64 { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			go__1086705844_0_0_0 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(arr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_go__1086705844_0_0_0(acc_1_loop_val.IntVal, func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(arr_2_loop_val.UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))
				})
			})
			Call_local_Main_go__go_0_1_1 = func(acc_1_loop int64, arr_2_loop []gopurs_runtime.Value) int64 {
			go__go_0_1_1:
				for {
					if false {
						continue go__go_0_1_1
					}
					var acc_1 int64 = acc_1_loop
					_ = acc_1
					var arr_2 []gopurs_runtime.Value = arr_2_loop
					_ = arr_2
					var __t3 int64
					{
						if (gopurs_runtime.Int(int64(len(arr_2))).IntVal) == (int64(0)) {
							__t3 = acc_1
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_go__1086705844_0_0_0((acc_1)+(int64(1)), func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(len(arr_2))), gopurs_runtime.Array(arr_2)).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
					}
				end_branch_3:
					return __t3
				}
			}
			go__go_0_1_1 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(arr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_go__go_0_1_1(acc_1_loop_val.IntVal, func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(arr_2_loop_val.UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))
				})
			})
			return gopurs_runtime.Apply(go__1086705844_0_0_0, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Main_length
}

var cache_Main_length__1126613660 gopurs_runtime.Value
var once_Main_length__1126613660 sync.Once

func Get_Main_length__1126613660() gopurs_runtime.Value {
	once_Main_length__1126613660.Do(func() {
		cache_Main_length__1126613660 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_length__1126613660(func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_length__1126613660
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_length__1126613660(func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(10000))).UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_length__1126613660(__eta_norm_0_0_loop []int64) int64 {
length__1126613660:
	for {
		if false {
			continue length__1126613660
		}
		var __eta_norm_0_0 []int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		var Call_local_Main_go__4172905617_1_0_2 func(int64, []int64) int64
		_ = Call_local_Main_go__4172905617_1_0_2
		var go__4172905617_1_0_2 gopurs_runtime.Value
		_ = go__4172905617_1_0_2
		var Call_local_Main_go__go_1_1_3 func(int64, []int64) int64
		_ = Call_local_Main_go__go_1_1_3
		var go__go_1_1_3 gopurs_runtime.Value
		_ = go__go_1_1_3
		Call_local_Main_go__4172905617_1_0_2 = func(acc_2_loop int64, arr_3_loop []int64) int64 {
		go__4172905617_1_0_2:
			for {
				if false {
					continue go__4172905617_1_0_2
				}
				var acc_2 int64 = acc_2_loop
				_ = acc_2
				var arr_3 []int64 = arr_3_loop
				_ = arr_3
				var __t2 int64
				{
					if (gopurs_runtime.Int(int64(len(arr_3))).IntVal) == (int64(0)) {
						__t2 = acc_2
						goto end_branch_2
					} else {

					}
				}
				{
					acc_2_loop = (acc_2) + (int64(1))
					arr_3_loop = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(arr_3))).IntVal), func() gopurs_runtime.Value {
							arr := arr_3
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()).UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					continue go__4172905617_1_0_2
					__t2 = func() int64 { panic("unreachable") }()
				}
			end_branch_2:
				return __t2
			}
		}
		go__4172905617_1_0_2 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(arr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Main_go__4172905617_1_0_2(acc_2_loop_val.IntVal, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(arr_3_loop_val.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()))
			})
		})
		Call_local_Main_go__go_1_1_3 = func(acc_2_loop int64, arr_3_loop []int64) int64 {
		go__go_1_1_3:
			for {
				if false {
					continue go__go_1_1_3
				}
				var acc_2 int64 = acc_2_loop
				_ = acc_2
				var arr_3 []int64 = arr_3_loop
				_ = arr_3
				var __t3 int64
				{
					if (gopurs_runtime.Int(int64(len(arr_3))).IntVal) == (int64(0)) {
						__t3 = acc_2
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = Call_local_Main_go__4172905617_1_0_2((acc_2)+(int64(1)), func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(arr_3))).IntVal), func() gopurs_runtime.Value {
							arr := arr_3
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()).UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}())
				}
			end_branch_3:
				return __t3
			}
		}
		go__go_1_1_3 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(arr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Main_go__go_1_1_3(acc_2_loop_val.IntVal, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(arr_3_loop_val.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()))
			})
		})
		return Call_local_Main_go__4172905617_1_0_2(int64(0), __eta_norm_0_0)
	}
}
