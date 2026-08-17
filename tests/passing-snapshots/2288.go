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
			var go__go_0_0_0 gopurs_runtime.Value
			go__go_0_0_0 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(arr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var acc_1_loop int64 = acc_1_loop_val.IntVal
						var arr_2_loop []gopurs_runtime.Value = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(arr_2_loop_val.UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
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
								if (gopurs_runtime.Int(int64(len(arr_2))).IntVal) == (0) {
									__t1 = acc_1
									goto end_branch_1
								} else {

								}
							}
							{
								acc_1_loop = (acc_1) + (1)
								arr_2_loop = func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(int64(len(arr_2))), gopurs_runtime.Array(arr_2)).UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()
								continue go__go_0_0_0
								__t1 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_1:
							return gopurs_runtime.Int(__t1)
						}
					}()
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_0, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_length
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_length(), func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(10000)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal)).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_length__4151727363 gopurs_runtime.Value
var once_Main_length__4151727363 sync.Once

func Get_Main_length__4151727363() gopurs_runtime.Value {
	once_Main_length__4151727363.Do(func() {
		cache_Main_length__4151727363 = func() gopurs_runtime.Value {
			var go__go_0_0_1 gopurs_runtime.Value
			go__go_0_0_1 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(arr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var acc_1_loop int64 = acc_1_loop_val.IntVal
						var arr_2_loop []gopurs_runtime.Value = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(arr_2_loop_val.UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
					go__go_0_0_1:
						for {
							if false {
								continue go__go_0_0_1
							}
							var acc_1 int64 = acc_1_loop
							_ = acc_1
							var arr_2 []gopurs_runtime.Value = arr_2_loop
							_ = arr_2
							var __t1 int64
							{
								if (gopurs_runtime.Int(int64(len(arr_2))).IntVal) == (0) {
									__t1 = acc_1
									goto end_branch_1
								} else {

								}
							}
							{
								acc_1_loop = (acc_1) + (1)
								arr_2_loop = func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(int64(len(arr_2))), gopurs_runtime.Array(arr_2)).UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()
								continue go__go_0_0_1
								__t1 = gopurs_runtime.Value{}.IntVal
							}
						end_branch_1:
							return gopurs_runtime.Int(__t1)
						}
					}()
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_1, gopurs_runtime.Int(0))
		}()
	})
	return cache_Main_length__4151727363
}
