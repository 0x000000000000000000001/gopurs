package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_monoidAdditive gopurs_runtime.Value
var once_Main_monoidAdditive sync.Once

func Get_Main_monoidAdditive() gopurs_runtime.Value {
	once_Main_monoidAdditive.Do(func() {
		cache_Main_monoidAdditive = func() gopurs_runtime.Value {
			// TAST (Let): semigroupAdditive1_0_0 -> *Constructor_Data_Semigroup_Semigroup
			semigroupAdditive1_0_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((v_0.IntVal) + (v1_1.IntVal))
				})
			})}
			_ = semigroupAdditive1_0_0
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_0_0)}
			}), gopurs_runtime.Int(0)})}
		}()
	})
	return cache_Main_monoidAdditive
}

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(&Constructor_Main_Foo{1, value0.IntVal})}
		})
	})
	return cache_Main_Foo
}

var cache_Main_Bar gopurs_runtime.Value
var once_Main_Bar sync.Once

func Get_Main_Bar() gopurs_runtime.Value {
	once_Main_Bar.Do(func() {
		cache_Main_Bar = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2512729583, UnsafePtr: unsafe.Pointer(&Constructor_Main_Bar{1, value0.IntVal})}
		})
	})
	return cache_Main_Bar
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(v_0_box))
		})
	})
	return cache_Main_g
}

type Constructor_Main_Foo struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_Bar struct {
	Rc uint32
	V0 int64
}

func Call_Main_test(v_0_loop *Constructor_Data_Maybe_Just) int64 {
	var v_0 *Constructor_Data_Maybe_Just = v_0_loop
	_ = v_0
	var __t3 int64
	{
		if v_0 != nil {
			__t3 = (v_0).V0.IntVal
			goto end_branch_3
		} else {

		}
	}
	{
		// TAST (Let): semigroupAdditive1_1_1 -> *Constructor_Data_Semigroup_Semigroup
		semigroupAdditive1_1_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((v_1.IntVal) + (v1_2.IntVal))
			})
		})}
		_ = semigroupAdditive1_1_1
		// TAST (Let): __local_var_1_0 -> *Constructor_Data_Monoid_Monoid
		var __local_var_1_0 *Constructor_Data_Monoid_Monoid = gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_1_1)}
		}), gopurs_runtime.Int(0)})})
		// TAST (Let): Semigroup0_2_2 -> *Constructor_Data_Semigroup_Semigroup
		Semigroup0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_0.V0), gopurs_runtime.Value{}))
		_ = Semigroup0_2_2
		__t3 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_2.V0), x_3, acc_4)
			})
		}), gopurs_runtime.Box(__local_var_1_0.V1), func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
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
		}()).IntVal
	}
end_branch_3:
	return __t3
}

func Call_Main_g(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 int64
	{
		if v_0.Type == 9 && v_0.IntVal == 2512729583 {
			__t0 = (*Constructor_Main_Bar)(v_0.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 2763139640 {
			__t0 = (*Constructor_Main_Foo)(v_0.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = 42
	}
end_branch_0:
	return __t0
}
