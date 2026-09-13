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
			var Call_local_Main_go__3739387282_0_0_0 func(*Constructor_Main_C[gopurs_runtime.Value], *Constructor_Main_C[gopurs_runtime.Value]) *Constructor_Main_C[gopurs_runtime.Value]
			_ = Call_local_Main_go__3739387282_0_0_0
			var go__3739387282_0_0_0 gopurs_runtime.Value
			_ = go__3739387282_0_0_0
			var Call_local_Main_go__go_0_1_1 func(*Constructor_Main_C[gopurs_runtime.Value], *Constructor_Main_C[gopurs_runtime.Value]) *Constructor_Main_C[gopurs_runtime.Value]
			_ = Call_local_Main_go__go_0_1_1
			var go__go_0_1_1 gopurs_runtime.Value
			_ = go__go_0_1_1
			Call_local_Main_go__3739387282_0_0_0 = func(v_1_loop *Constructor_Main_C[gopurs_runtime.Value], v1_2_loop *Constructor_Main_C[gopurs_runtime.Value]) *Constructor_Main_C[gopurs_runtime.Value] {
			go__3739387282_0_0_0:
				for {
					if false {
						continue go__3739387282_0_0_0
					}
					var v_1 *Constructor_Main_C[gopurs_runtime.Value] = v_1_loop
					_ = v_1
					var v1_2 *Constructor_Main_C[gopurs_runtime.Value] = v1_2_loop
					_ = v1_2
					var __t4 *Constructor_Main_C[gopurs_runtime.Value]
					{
						if v1_2 == nil {
							__t4 = v_1
							goto end_branch_4
						} else {

						}
					}
					{
						if v1_2 != nil {
							var __t3 *Constructor_Main_C[gopurs_runtime.Value]
							{
								var __t_tag_2 gopurs_runtime.Value = (v1_2).V0
								_ = __t_tag_2
								if __t_tag_2.Type == 9 && __t_tag_2.IntVal == 4133812178 {
									v_1_loop = (&Constructor_Main_C[gopurs_runtime.Value]{1, (*Constructor_Main_L[gopurs_runtime.Value, gopurs_runtime.Value])((v1_2).V0.UnsafePtr).V0, v_1})
									v1_2_loop = (v1_2).V1
									continue go__3739387282_0_0_0
									__t3 = func() *Constructor_Main_C[gopurs_runtime.Value] { panic("unreachable") }()
									goto end_branch_3
								} else {

								}
							}
							{
								v_1_loop = v_1
								v1_2_loop = (v1_2).V1
								continue go__3739387282_0_0_0
								__t3 = func() *Constructor_Main_C[gopurs_runtime.Value] { panic("unreachable") }()
							}
						end_branch_3:
							__t4 = __t3
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = func() *Constructor_Main_C[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_4:
					return __t4
				}
			}
			go__3739387282_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer(Call_local_Main_go__3739387282_0_0_0(gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](v_1_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](v1_2_loop_val)))}
				})
			})
			Call_local_Main_go__go_0_1_1 = func(v_1_loop *Constructor_Main_C[gopurs_runtime.Value], v1_2_loop *Constructor_Main_C[gopurs_runtime.Value]) *Constructor_Main_C[gopurs_runtime.Value] {
			go__go_0_1_1:
				for {
					if false {
						continue go__go_0_1_1
					}
					var v_1 *Constructor_Main_C[gopurs_runtime.Value] = v_1_loop
					_ = v_1
					var v1_2 *Constructor_Main_C[gopurs_runtime.Value] = v1_2_loop
					_ = v1_2
					var __t7 *Constructor_Main_C[gopurs_runtime.Value]
					{
						if v1_2 == nil {
							__t7 = v_1
							goto end_branch_7
						} else {

						}
					}
					{
						if v1_2 != nil {
							var __t6 *Constructor_Main_C[gopurs_runtime.Value]
							{
								var __t_tag_5 gopurs_runtime.Value = (v1_2).V0
								_ = __t_tag_5
								if __t_tag_5.Type == 9 && __t_tag_5.IntVal == 4133812178 {
									__t6 = Call_local_Main_go__3739387282_0_0_0((&Constructor_Main_C[gopurs_runtime.Value]{1, (*Constructor_Main_L[gopurs_runtime.Value, gopurs_runtime.Value])((v1_2).V0.UnsafePtr).V0, v_1}), (v1_2).V1)
									goto end_branch_6
								} else {

								}
							}
							{
								__t6 = Call_local_Main_go__3739387282_0_0_0(v_1, (v1_2).V1)
							}
						end_branch_6:
							__t7 = __t6
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = func() *Constructor_Main_C[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_7:
					return __t7
				}
			}
			go__go_0_1_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer(Call_local_Main_go__go_0_1_1(gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](v_1_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](v1_2_loop_val)))}
				})
			})
			return gopurs_runtime.Apply(go__3739387282_0_0_0, gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((*Constructor_Main_C[gopurs_runtime.Value])(nil))})
		}()
	})
	return cache_Main_lefts
}

type Constructor_Main_C[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Main_C[T_a]
}

type Constructor_Main_N[T_a any] struct {
	Rc uint32
}

type Constructor_Main_L[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_R[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}
