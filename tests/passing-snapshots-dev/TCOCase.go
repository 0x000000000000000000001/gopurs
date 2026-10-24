package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_One gopurs_runtime.Value
var once_Main_One sync.Once

func Get_Main_One() gopurs_runtime.Value {
	once_Main_One.Do(func() {
		cache_Main_One = gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer((*Constructor_Main_More)(nil))}
	})
	return cache_Main_One
}

var cache_Main_More gopurs_runtime.Value
var once_Main_More sync.Once

func Get_Main_More() gopurs_runtime.Value {
	once_Main_More.Do(func() {
		cache_Main_More = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer((&Constructor_Main_More{1, gopurs_runtime.CoerceToStruct[Constructor_Main_More](value0)}))}
		})
	})
	return cache_Main_More
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var Call_local_Main_to_0_0_0 func(float64, *Constructor_Main_More) *Constructor_Main_More
			_ = Call_local_Main_to_0_0_0
			var to_0_0_0 gopurs_runtime.Value
			_ = to_0_0_0
			Call_local_Main_to_0_0_0 = func(v_1_loop float64, v1_2_loop *Constructor_Main_More) *Constructor_Main_More {
			to_0_0_0:
				for {
					if false {
						continue to_0_0_0
					}
					var v_1 float64 = v_1_loop
					_ = v_1
					var v1_2 *Constructor_Main_More = v1_2_loop
					_ = v1_2
					var __t1 *Constructor_Main_More
					{
						if (v_1) == (0.0) {
							__t1 = v1_2
							goto end_branch_1
						} else {

						}
					}
					{
						v_1_loop = (v_1) - (1.0)
						v1_2_loop = (&Constructor_Main_More{1, v1_2})
						continue to_0_0_0
						__t1 = func() *Constructor_Main_More { panic("unreachable") }()
					}
				end_branch_1:
					return __t1
				}
			}
			to_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer(Call_local_Main_to_0_0_0(v_1_loop_val.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Main_More](v1_2_loop_val)))}
				})
			})
			var Call_local_Main_from_1_2_1 func(*Constructor_Main_More) string
			_ = Call_local_Main_from_1_2_1
			var from_1_2_1 gopurs_runtime.Value
			_ = from_1_2_1
			Call_local_Main_from_1_2_1 = func(v_2_loop *Constructor_Main_More) string {
			from_1_2_1:
				for {
					if false {
						continue from_1_2_1
					}
					var v_2 *Constructor_Main_More = v_2_loop
					_ = v_2
					var __t3 string
					{
						if v_2 == nil {
							__t3 = "Done"
							goto end_branch_3
						} else {

						}
					}
					{
						if v_2 != nil {
							v_2_loop = (v_2).V0
							continue from_1_2_1
							__t3 = func() string { panic("unreachable") }()
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = func() string { panic("Failed pattern match") }()
					}
				end_branch_3:
					return __t3
				}
			}
			from_1_2_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(Call_local_Main_from_1_2_1(gopurs_runtime.CoerceToStruct[Constructor_Main_More](v_2_loop_val)))
			})
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_local_Main_from_1_2_1(Call_local_Main_to_0_0_0(10000.0, (*Constructor_Main_More)(nil)))))
		}()
	})
	return cache_Main_main
}

type Constructor_Main_One struct {
	Rc uint32
}

type Constructor_Main_More struct {
	Rc uint32
	V0 *Constructor_Main_More
}
