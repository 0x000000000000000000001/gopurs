package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_singleton gopurs_runtime.Value
var once_Main_singleton sync.Once

func Get_Main_singleton() gopurs_runtime.Value {
	once_Main_singleton.Do(func() {
		cache_Main_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_singleton(x_0_box))
		})
	})
	return cache_Main_singleton
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_foldMap gopurs_runtime.Value
var once_Main_foldMap sync.Once

func Get_Main_foldMap() gopurs_runtime.Value {
	once_Main_foldMap.Do(func() {
		cache_Main_foldMap = gopurs_runtime.Func3(func(dictSemigroup_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0_box), v_1_box, func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v1_2_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})
	})
	return cache_Main_foldMap
}

var cache_Main_empty gopurs_runtime.Value
var once_Main_empty sync.Once

func Get_Main_empty() gopurs_runtime.Value {
	once_Main_empty.Do(func() {
		cache_Main_empty = gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())
	})
	return cache_Main_empty
}

var cache_Main_regression gopurs_runtime.Value
var once_Main_regression sync.Once

func Get_Main_regression() gopurs_runtime.Value {
	once_Main_regression.Do(func() {
		cache_Main_regression = func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(2), gopurs_runtime.Int(3)}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()).UnsafePtr)
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
		}()
	})
	return cache_Main_regression
}

var cache_Main_empty__4047862030 gopurs_runtime.Value
var once_Main_empty__4047862030 sync.Once

func Get_Main_empty__4047862030() gopurs_runtime.Value {
	once_Main_empty__4047862030.Do(func() {
		cache_Main_empty__4047862030 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())
	})
	return cache_Main_empty__4047862030
}

var cache_Main_foldMap__1307344740 gopurs_runtime.Value
var once_Main_foldMap__1307344740 sync.Once

func Get_Main_foldMap__1307344740() gopurs_runtime.Value {
	once_Main_foldMap__1307344740.Do(func() {
		cache_Main_foldMap__1307344740 = gopurs_runtime.Func3(func(dictSemigroup_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldMap__1307344740(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0_box), v_1_box, func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v1_2_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})
	})
	return cache_Main_foldMap__1307344740
}

var cache_Main_singleton__2286220742 gopurs_runtime.Value
var once_Main_singleton__2286220742 sync.Once

func Get_Main_singleton__2286220742() gopurs_runtime.Value {
	once_Main_singleton__2286220742.Do(func() {
		cache_Main_singleton__2286220742 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_singleton__2286220742(x_0_box))
		})
	})
	return cache_Main_singleton__2286220742
}

func Call_Main_singleton(x_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return func() []gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
		unboxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = v
		}
		return unboxed
	}()
}

func Call_Main_foldMap(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup, v_1_loop gopurs_runtime.Value, v1_2_loop []gopurs_runtime.Value) gopurs_runtime.Value {
foldMap:
	for {
		if false {
			continue foldMap
		}
		var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_0_loop
		_ = dictSemigroup_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		var v1_2 []gopurs_runtime.Value = v1_2_loop
		_ = v1_2
		var __t0 gopurs_runtime.Value
		{
			if (gopurs_runtime.Int(int64(len(v1_2))).IntVal) == (5) {
				__t0 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 0)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 1)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 2)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 3)), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 4))))))
				goto end_branch_0
			} else {

			}
		}
		{
			dictSemigroup_0_loop = dictSemigroup_0
			v_1_loop = v_1
			v1_2_loop = v1_2
			continue foldMap
			__t0 = gopurs_runtime.Value{}
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_foldMap__1307344740(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup, v_1_loop gopurs_runtime.Value, v1_2_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var v1_2 []gopurs_runtime.Value = v1_2_loop
	_ = v1_2
	var __t1 gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(v1_2))).IntVal) == (5) {
			__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 0)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 1)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 2)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 3)), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 4))))))
			goto end_branch_1
		} else {

		}
	}
	{
		var __t0 gopurs_runtime.Value
		{
			if (gopurs_runtime.Int(int64(len(v1_2))).IntVal) == (5) {
				__t0 = gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 0)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 1)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 2)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 3)), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 4))))))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = Call_Main_foldMap(dictSemigroup_0, v_1, v1_2)
		}
	end_branch_0:
		__t1 = __t0
	}
end_branch_1:
	return __t1
}

func Call_Main_singleton__2286220742(x_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return func() []gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
		unboxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = v
		}
		return unboxed
	}()
}
