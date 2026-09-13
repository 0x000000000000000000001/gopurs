package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_ArrayBox gopurs_runtime.Value
var once_Main_ArrayBox sync.Once

func Get_Main_ArrayBox() gopurs_runtime.Value {
	once_Main_ArrayBox.Do(func() {
		cache_Main_ArrayBox = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_ArrayBox
}

var cache_Main_nil gopurs_runtime.Value
var once_Main_nil sync.Once

func Get_Main_nil() gopurs_runtime.Value {
	once_Main_nil.Do(func() {
		cache_Main_nil = gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())
	})
	return cache_Main_nil
}

var cache_Main_cons_prime_ gopurs_runtime.Value
var once_Main_cons_prime_ sync.Once

func Get_Main_cons_prime_() gopurs_runtime.Value {
	once_Main_cons_prime_.Do(func() {
		cache_Main_cons_prime_ = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons_prime_(x_0_box, v_1_box)
		})
	})
	return cache_Main_cons_prime_
}

var cache_Main_cons_prime___498012883 gopurs_runtime.Value
var once_Main_cons_prime___498012883 sync.Once

func Get_Main_cons_prime___498012883() gopurs_runtime.Value {
	once_Main_cons_prime___498012883.Do(func() {
		cache_Main_cons_prime___498012883 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons_prime___498012883(x_0_box.IntVal, v_unused_1_box)
		})
	})
	return cache_Main_cons_prime___498012883
}

var cache_Main_cons_prime___1736767744 gopurs_runtime.Value
var once_Main_cons_prime___1736767744 sync.Once

func Get_Main_cons_prime___1736767744() gopurs_runtime.Value {
	once_Main_cons_prime___1736767744.Do(func() {
		cache_Main_cons_prime___1736767744 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons_prime___1736767744(x_0_box.IntVal, v_1_box)
		})
	})
	return cache_Main_cons_prime___1736767744
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): v_0_0 shape=App(Var) bindingType=(TypeApp (ADT ["Main","ArrayBox"] []) [(TypeApp (ADT ["Main","S"] []) [(TypeApp (ADT ["Main","S"] []) [(TypeApp (ADT ["Main","S"] []) [(ADT ["Main","Z"] [])])])]), Int])
			v_0_0 := gopurs_runtime.Apply2(Get_Main_cons_prime_(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Apply2(Get_Main_cons_prime_(), gopurs_runtime.Int(int64(2)), gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
								arr := []int64{int64(3)}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), Get_Main_nil()).UnsafePtr)
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
				}().UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())))
			_ = v_0_0
			var __t1 gopurs_runtime.Value
			{
				if ((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v_0_0))).IntVal) == (int64(3))) && (((gopurs_runtime.ArrayAccess(v_0_0, 0).IntVal) == (int64(1))) && (((gopurs_runtime.ArrayAccess(v_0_0, 1).IntVal) == (int64(2))) && ((gopurs_runtime.ArrayAccess(v_0_0, 2).IntVal) == (int64(3))))) {
					__t1 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Failed"), gopurs_runtime.Bool(false))
			}
		end_branch_1:
			return __t1
		}()
	})
	return cache_Main_main
}

type Constructor_Main_ArrayBox[T_n any, T_a any] struct {
	Rc uint32
	V0 []gopurs_runtime.Value
}

func Call_Main_cons_prime_(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Array(func() []gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), v_1).UnsafePtr)
		unboxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = v
		}
		return unboxed
	}())
}

func Call_Main_cons_prime___498012883(x_0_loop int64, v_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
cons_prime___498012883:
	for {
		if false {
			continue cons_prime___498012883
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var v_unused_1 gopurs_runtime.Value = v_unused_1_loop
		_ = v_unused_1
		return gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
							arr := []int64{x_0}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), Get_Main_nil()).UnsafePtr)
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
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())
	}
}

func Call_Main_cons_prime___1736767744(x_0_loop int64, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
cons_prime___1736767744:
	for {
		if false {
			continue cons_prime___1736767744
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
							arr := []int64{x_0}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), v_1).UnsafePtr)
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
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())
	}
}
