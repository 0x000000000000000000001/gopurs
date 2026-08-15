package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((&Constructor_Main_B{1, value0, value1}))}
			})
		})
	})
	return cache_Main_B
}

var cache_Main_memptyB gopurs_runtime.Value
var once_Main_memptyB sync.Once

func Get_Main_memptyB() gopurs_runtime.Value {
	once_Main_memptyB.Do(func() {
		cache_Main_memptyB = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_memptyB(dictMonoid_0_box)
		})
	})
	return cache_Main_memptyB
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 string
			{
				// TAST (Let): __local_var_0_1 -> *Constructor_Main_B
				__local_var_0_1 := (&Constructor_Main_B{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())
				}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())
				})})
				_ = __local_var_0_1
				if (gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(true)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply((__local_var_0_1).V0, gopurs_runtime.Int(0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply((__local_var_0_1).V1, gopurs_runtime.Int(0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0)).IntVal) != (0) {
					__t2 = "true"
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = "false"
			}
		end_branch_2:
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t2))
			_ = __local_var_0_0
			_dollar___unused_1_3 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_3
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_B struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func Call_Main_memptyB(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((&Constructor_Main_B{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	})}))}
}
