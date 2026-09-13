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

var cache_Main_singleton__2309289500 gopurs_runtime.Value
var once_Main_singleton__2309289500 sync.Once

func Get_Main_singleton__2309289500() gopurs_runtime.Value {
	once_Main_singleton__2309289500.Do(func() {
		cache_Main_singleton__2309289500 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_singleton__2309289500(x_0_box.IntVal)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_singleton__2309289500
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
			return Call_Main_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box), v_1_box, func() []gopurs_runtime.Value {
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

var cache_Main_foldMap__869368821 gopurs_runtime.Value
var once_Main_foldMap__869368821 sync.Once

func Get_Main_foldMap__869368821() gopurs_runtime.Value {
	once_Main_foldMap__869368821.Do(func() {
		cache_Main_foldMap__869368821 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_foldMap__869368821(v_0_box, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v1_1_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_foldMap__869368821
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
			arr := Call_Main_foldMap__869368821(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 []int64
				{
					if ((int64(1)) < (x_0.IntVal)) && ((x_0.IntVal) < (int64(4))) {
						__t0 = []int64{x_0.IntVal}
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = []int64{}
				}
			end_branch_0:
				return func() gopurs_runtime.Value {
					arr := __t0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}), []int64{int64(1), int64(2), int64(3), int64(4), int64(5)})
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
	return cache_Main_regression
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

func Call_Main_singleton__2309289500(x_0_loop int64) []int64 {
singleton__2309289500:
	for {
		if false {
			continue singleton__2309289500
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return []int64{x_0}
	}
}

func Call_Main_foldMap(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, v1_2_loop []gopurs_runtime.Value) gopurs_runtime.Value {
foldMap:
	for {
		if false {
			continue foldMap
		}
		var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
		_ = dictSemigroup_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		var v1_2 []gopurs_runtime.Value = v1_2_loop
		_ = v1_2
		var __t0 gopurs_runtime.Value
		{
			if (gopurs_runtime.Int(int64(len(v1_2))).IntVal) == (int64(5)) {
				__t0 = gopurs_runtime.Apply2(dictSemigroup_0.V0, gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 0)), gopurs_runtime.Apply2(dictSemigroup_0.V0, gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 1)), gopurs_runtime.Apply2(dictSemigroup_0.V0, gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 2)), gopurs_runtime.Apply2(dictSemigroup_0.V0, gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 3)), gopurs_runtime.Apply(v_1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v1_2), 4))))))
				goto end_branch_0
			} else {

			}
		}
		{
			dictSemigroup_0_loop = dictSemigroup_0
			v_1_loop = v_1
			v1_2_loop = v1_2
			continue foldMap
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_foldMap__869368821(v_0_loop gopurs_runtime.Value, v1_1_loop []int64) []int64 {
foldMap__869368821:
	for {
		if false {
			continue foldMap__869368821
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 []int64 = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (gopurs_runtime.Int(int64(len(v1_1))).IntVal) == (int64(5)) {
				__t0 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply(v_0, gopurs_runtime.Int(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := v1_1
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 0).IntVal)), gopurs_runtime.Apply(v_0, gopurs_runtime.Int(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := v1_1
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 1).IntVal))), gopurs_runtime.Apply(v_0, gopurs_runtime.Int(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := v1_1
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 2).IntVal))), gopurs_runtime.Apply(v_0, gopurs_runtime.Int(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := v1_1
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 3).IntVal))), gopurs_runtime.Apply(v_0, gopurs_runtime.Int(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := v1_1
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), 4).IntVal))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = Call_Main_foldMap(Rebox_Main_4291402899_4179793454(Rebox_Main_4179793454_4291402899(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupArray()))), v_0, func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := v1_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		}
	end_branch_0:
		return func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(__t0.UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()
	}
}

func Rebox_Main_4179793454_4291402899(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semigroup_Semigroup[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4291402899_4179793454(in *Constructor_Data_Semigroup_Semigroup[[]int64]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
