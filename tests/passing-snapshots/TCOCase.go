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
			var to_0_0_0 gopurs_runtime.Value
			to_0_0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						var v_1_loop gopurs_runtime.Value = v_1_loop_val
						var v1_2_loop *Constructor_Main_More = gopurs_runtime.CoerceToStruct[Constructor_Main_More](v1_2_loop_val)
					to_0_0_0:
						for {
							if false {
								continue to_0_0_0
							}
							var v_1 gopurs_runtime.Value = v_1_loop
							_ = v_1
							var v1_2 *Constructor_Main_More = v1_2_loop
							_ = v1_2
							var __t1 *Constructor_Main_More
							{
								if (v_1.FloatVal()) == (0.0) {
									__t1 = v1_2
									goto end_branch_1
								} else {

								}
							}
							{
								v_1_loop = gopurs_runtime.Float((v_1.FloatVal()) - (1.0))
								v1_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Main_More](gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer((&Constructor_Main_More{1, v1_2}))})
								continue to_0_0_0
								__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_More](gopurs_runtime.Value{})
							}
						end_branch_1:
							return gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer(__t1)}
						}
					}()
				})
			})
			var from_1_2_1 gopurs_runtime.Value
			from_1_2_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					var v_2_loop gopurs_runtime.Value = v_2_loop_val
				from_1_2_1:
					for {
						if false {
							continue from_1_2_1
						}
						var v_2 gopurs_runtime.Value = v_2_loop
						_ = v_2
						var __t3 string
						{
							if v_2.Type == 9 && v_2.IntVal == 2472542475 && v_2.UnsafePtr == nil {
								__t3 = "Done"
								goto end_branch_3
							} else {

							}
						}
						{
							if v_2.Type == 9 && v_2.IntVal == 2472542475 && v_2.UnsafePtr != nil {
								v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer((*Constructor_Main_More)(v_2.UnsafePtr).V0)}
								continue from_1_2_1
								__t3 = gopurs_runtime.Value{}.StrVal()
								goto end_branch_3
							} else {

							}
						}
						{
							__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
						}
					end_branch_3:
						return gopurs_runtime.Str(__t3)
					}
				}()
			})
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(from_1_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_More](gopurs_runtime.Apply2(to_0_0_0, gopurs_runtime.Float(10000.0), gopurs_runtime.Value{Type: 9, IntVal: 2472542475, UnsafePtr: unsafe.Pointer((*Constructor_Main_More)(nil))})))}).StrVal()))
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
