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
			var Call_local_Main_to__805151591_0_0_0 func(float64, *Constructor_Main_More) *Constructor_Main_More
			_ = Call_local_Main_to__805151591_0_0_0
			var to__805151591_0_0_0 gopurs_runtime.Value
			_ = to__805151591_0_0_0
			var Call_local_Main_to_0_1_1 func(float64, *Constructor_Main_More) *Constructor_Main_More
			_ = Call_local_Main_to_0_1_1
			var to_0_1_1 gopurs_runtime.Value
			_ = to_0_1_1
			Call_local_Main_to__805151591_0_0_0 = func(v_1_loop float64, v1_2_loop *Constructor_Main_More) *Constructor_Main_More {
			to__805151591_0_0_0:
				for {
					if false {
						continue to__805151591_0_0_0
					}
					var v_1 float64 = v_1_loop
					_ = v_1
					var v1_2 *Constructor_Main_More = v1_2_loop
					_ = v1_2
					var __t2 *Constructor_Main_More
					{
						if (v_1) == (0.0) {
							__t2 = v1_2
							goto end_branch_2
						} else {

						}
					}
					{
						v_1_loop = (v_1) - (1.0)
						v1_2_loop = (&Constructor_Main_More{1, v1_2})
						continue to__805151591_0_0_0
						__t2 = func() *Constructor_Main_More { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}
			to__805151591_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer(Call_local_Main_to__805151591_0_0_0(v_1_loop_val.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Main_More](v1_2_loop_val)))}
				})
			})
			Call_local_Main_to_0_1_1 = func(v_1_loop float64, v1_2_loop *Constructor_Main_More) *Constructor_Main_More {
			to_0_1_1:
				for {
					if false {
						continue to_0_1_1
					}
					var v_1 float64 = v_1_loop
					_ = v_1
					var v1_2 *Constructor_Main_More = v1_2_loop
					_ = v1_2
					var __t3 *Constructor_Main_More
					{
						if (v_1) == (0.0) {
							__t3 = v1_2
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = Call_local_Main_to__805151591_0_0_0((v_1)-(1.0), (&Constructor_Main_More{1, v1_2}))
					}
				end_branch_3:
					return __t3
				}
			}
			to_0_1_1 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer(Call_local_Main_to_0_1_1(v_1_loop_val.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Main_More](v1_2_loop_val)))}
				})
			})
			var Call_local_Main_from__3402454404_1_4_2 func(*Constructor_Main_More) string
			_ = Call_local_Main_from__3402454404_1_4_2
			var from__3402454404_1_4_2 gopurs_runtime.Value
			_ = from__3402454404_1_4_2
			var Call_local_Main_from_1_5_3 func(*Constructor_Main_More) string
			_ = Call_local_Main_from_1_5_3
			var from_1_5_3 gopurs_runtime.Value
			_ = from_1_5_3
			Call_local_Main_from__3402454404_1_4_2 = func(v_2_loop *Constructor_Main_More) string {
			from__3402454404_1_4_2:
				for {
					if false {
						continue from__3402454404_1_4_2
					}
					var v_2 *Constructor_Main_More = v_2_loop
					_ = v_2
					var __t6 string
					{
						if v_2 == nil {
							__t6 = "Done"
							goto end_branch_6
						} else {

						}
					}
					{
						if v_2 != nil {
							v_2_loop = (v_2).V0
							continue from__3402454404_1_4_2
							__t6 = func() string { panic("unreachable") }()
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = func() string { panic("Failed pattern match") }()
					}
				end_branch_6:
					return __t6
				}
			}
			from__3402454404_1_4_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(Call_local_Main_from__3402454404_1_4_2(gopurs_runtime.CoerceToStruct[Constructor_Main_More](v_2_loop_val)))
			})
			Call_local_Main_from_1_5_3 = func(v_2_loop *Constructor_Main_More) string {
			from_1_5_3:
				for {
					if false {
						continue from_1_5_3
					}
					var v_2 *Constructor_Main_More = v_2_loop
					_ = v_2
					var __t7 string
					{
						if v_2 == nil {
							__t7 = "Done"
							goto end_branch_7
						} else {

						}
					}
					{
						if v_2 != nil {
							__t7 = Call_local_Main_from__3402454404_1_4_2((v_2).V0)
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = func() string { panic("Failed pattern match") }()
					}
				end_branch_7:
					return __t7
				}
			}
			from_1_5_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(Call_local_Main_from_1_5_3(gopurs_runtime.CoerceToStruct[Constructor_Main_More](v_2_loop_val)))
			})
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_local_Main_from__3402454404_1_4_2(Call_local_Main_to__805151591_0_0_0(10000.0, (*Constructor_Main_More)(nil)))))
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
