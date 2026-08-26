package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_C gopurs_runtime.Value
var once_Main_C sync.Once

func Get_Main_C() gopurs_runtime.Value {
	once_Main_C.Do(func() {
		cache_Main_C = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((&Constructor_Main_C[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Main_C
}

var cache_Main_N gopurs_runtime.Value
var once_Main_N sync.Once

func Get_Main_N() gopurs_runtime.Value {
	once_Main_N.Do(func() {
		cache_Main_N = gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((*Constructor_Main_C[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_N
}

var cache_Main_L gopurs_runtime.Value
var once_Main_L sync.Once

func Get_Main_L() gopurs_runtime.Value {
	once_Main_L.Do(func() {
		cache_Main_L = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4133812178, UnsafePtr: unsafe.Pointer((&Constructor_Main_L[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_L
}

var cache_Main_R gopurs_runtime.Value
var once_Main_R sync.Once

func Get_Main_R() gopurs_runtime.Value {
	once_Main_R.Do(func() {
		cache_Main_R = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3558538316, UnsafePtr: unsafe.Pointer((&Constructor_Main_R[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_R
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_lefts gopurs_runtime.Value
var once_Main_lefts sync.Once

func Get_Main_lefts() gopurs_runtime.Value {
	once_Main_lefts.Do(func() {
		cache_Main_lefts = func() gopurs_runtime.Value {
			var Call_local_Main_go__go_0_0_0 func(*Constructor_Main_C[gopurs_runtime.Value], gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Main_go__go_0_0_0
			var go__go_0_0_0 gopurs_runtime.Value
			_ = go__go_0_0_0
			Call_local_Main_go__go_0_0_0 = func(v_1_loop *Constructor_Main_C[gopurs_runtime.Value], v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_0_0_0:
				for {
					if false {
						continue go__go_0_0_0
					}
					var v_1 *Constructor_Main_C[gopurs_runtime.Value] = v_1_loop
					_ = v_1
					var v1_2 gopurs_runtime.Value = v1_2_loop
					_ = v1_2
					var __t3 *Constructor_Main_C[gopurs_runtime.Value]
					{
						if v1_2.Type == 9 && v1_2.IntVal == 2167983901 && v1_2.UnsafePtr == nil {
							__t3 = v_1
							goto end_branch_3
						} else {

						}
					}
					{
						if v1_2.Type == 9 && v1_2.IntVal == 2167983901 && v1_2.UnsafePtr != nil {
							var __t2 *Constructor_Main_C[gopurs_runtime.Value]
							{
								var __t_tag_1 gopurs_runtime.Value = (*Constructor_Main_C[gopurs_runtime.Value])(v1_2.UnsafePtr).V0
								if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 4133812178 {
									v_1_loop = (&Constructor_Main_C[gopurs_runtime.Value]{1, (*Constructor_Main_L[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Main_C[gopurs_runtime.Value])(v1_2.UnsafePtr).V0.UnsafePtr).V0, v_1})
									v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((*Constructor_Main_C[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
									continue go__go_0_0_0
									__t2 = gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](gopurs_runtime.Value{})
									goto end_branch_2
								} else {

								}
							}
							{
								v_1_loop = v_1
								v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((*Constructor_Main_C[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
								continue go__go_0_0_0
								__t2 = gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](gopurs_runtime.Value{})
							}
						end_branch_2:
							__t3 = __t2
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
					}
				end_branch_3:
					return gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer(__t3)}
				}
			}
			go__go_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Main_go__go_0_0_0(gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](v_1_loop_val), v1_2_loop_val)
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_0, gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((*Constructor_Main_C[gopurs_runtime.Value])(nil))})
		}()
	})
	return cache_Main_lefts
}

type Constructor_Main_C[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_C[gopurs_runtime.Value]
}

type Constructor_Main_N[T_a any] struct {
	Rc uint32
}

type Constructor_Main_L[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_R[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}
