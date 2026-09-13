package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Tuple gopurs_runtime.Value
var once_Main_Tuple sync.Once

func Get_Main_Tuple() gopurs_runtime.Value {
	once_Main_Tuple.Do(func() {
		cache_Main_Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Tuple
}

var cache_Main_Tuple__3273409525 gopurs_runtime.Value
var once_Main_Tuple__3273409525 sync.Once

func Get_Main_Tuple__3273409525() gopurs_runtime.Value {
	once_Main_Tuple__3273409525.Do(func() {
		cache_Main_Tuple__3273409525 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Rebox_Main_4045719060_1152654068(Call_Main_Tuple__3273409525(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.IntVal)))}
		})
	})
	return cache_Main_Tuple__3273409525
}

var cache_Main_State gopurs_runtime.Value
var once_Main_State sync.Once

func Get_Main_State() gopurs_runtime.Value {
	once_Main_State.Do(func() {
		cache_Main_State = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_State
}

var cache_Main_MonadState_dollar_Dict gopurs_runtime.Value
var once_Main_MonadState_dollar_Dict sync.Once

func Get_Main_MonadState_dollar_Dict() gopurs_runtime.Value {
	once_Main_MonadState_dollar_Dict.Do(func() {
		cache_Main_MonadState_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2980279296, UnsafePtr: unsafe.Pointer(Call_Main_MonadState_dollar_Dict(func() struct {
				get gopurs_runtime.Value
				put gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					get gopurs_runtime.Value
					put gopurs_runtime.Value
				}{}
				clone.get = gopurs_runtime.RecordGet(orig, "get")
				clone.put = gopurs_runtime.RecordGet(orig, "put")
				return clone
			}()))}
		})
	})
	return cache_Main_MonadState_dollar_Dict
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_test5(v_0_box.IntVal))
		})
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Rebox_Main_4045719060_1152654068(Call_Main_test4(Rebox_Main_1152654068_4045719060(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))))}
		})
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_test3(n_0_box.IntVal))
		})
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_test2(v_0_box))
		})
	})
	return cache_Main_test2
}

var cache_Main_runState gopurs_runtime.Value
var once_Main_runState sync.Once

func Get_Main_runState() gopurs_runtime.Value {
	once_Main_runState.Do(func() {
		cache_Main_runState = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Call_Main_runState(s_0_box, v_1_box))}
		})
	})
	return cache_Main_runState
}

var cache_Main_put gopurs_runtime.Value
var once_Main_put sync.Once

func Get_Main_put() gopurs_runtime.Value {
	once_Main_put.Do(func() {
		cache_Main_put = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_put(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_put
}

var cache_Main_monadStateState gopurs_runtime.Value
var once_Main_monadStateState sync.Once

func Get_Main_monadStateState() gopurs_runtime.Value {
	once_Main_monadStateState.Do(func() {
		cache_Main_monadStateState = gopurs_runtime.Value{Type: 9, IntVal: 2980279296, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_0, s_0}))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_0, Get_Data_Unit_unit()}))}
			})
		})}))}
	})
	return cache_Main_monadStateState
}

var cache_Main_get gopurs_runtime.Value
var once_Main_get sync.Once

func Get_Main_get() gopurs_runtime.Value {
	once_Main_get.Do(func() {
		cache_Main_get = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get(dict_0_box)
		})
	})
	return cache_Main_get
}

var cache_Main_get1 gopurs_runtime.Value
var once_Main_get1 sync.Once

func Get_Main_get1() gopurs_runtime.Value {
	once_Main_get1.Do(func() {
		cache_Main_get1 = Call_Main_get(Get_Main_monadStateState())
	})
	return cache_Main_get1
}

var cache_Main_modify gopurs_runtime.Value
var once_Main_modify sync.Once

func Get_Main_modify() gopurs_runtime.Value {
	once_Main_modify.Do(func() {
		cache_Main_modify = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_modify(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_modify
}

var cache_Main_monadState gopurs_runtime.Value
var once_Main_monadState sync.Once

func Get_Main_monadState() gopurs_runtime.Value {
	once_Main_monadState.Do(func() {
		cache_Main_monadState = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeState()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindState()))}
		})}))}
	})
	return cache_Main_monadState
}

var cache_Main_functorState gopurs_runtime.Value
var once_Main_functorState sync.Once

func Get_Main_functorState() gopurs_runtime.Value {
	once_Main_functorState.Do(func() {
		cache_Main_functorState = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Call_Control_Monad_liftM1(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadState()))}))}
	})
	return cache_Main_functorState
}

var cache_Main_bindState gopurs_runtime.Value
var once_Main_bindState sync.Once

func Get_Main_bindState() gopurs_runtime.Value {
	once_Main_bindState.Do(func() {
		cache_Main_bindState = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyState()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Main","Tuple"] [(TypeVar s$scope38), (TypeVar a$scope42)])
				v_3_0 := Call_Main_runState(s_2, f_0)
				_ = v_3_0
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Call_Main_runState((v_3_0).V0, gopurs_runtime.Apply(g_1, (v_3_0).V1)))}
			})
		})}))}
	})
	return cache_Main_bindState
}

var cache_Main_applyState gopurs_runtime.Value
var once_Main_applyState sync.Once

func Get_Main_applyState() gopurs_runtime.Value {
	once_Main_applyState.Do(func() {
		cache_Main_applyState = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorState()))}
		}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadState()))}))}
	})
	return cache_Main_applyState
}

var cache_Main_applicativeState gopurs_runtime.Value
var once_Main_applicativeState sync.Once

func Get_Main_applicativeState() gopurs_runtime.Value {
	once_Main_applicativeState.Do(func() {
		cache_Main_applicativeState = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_1, a_0}))}
			})
		})}))}
	})
	return cache_Main_applicativeState
}

var cache_Main_modify1 gopurs_runtime.Value
var once_Main_modify1 sync.Once

func Get_Main_modify1() gopurs_runtime.Value {
	once_Main_modify1.Do(func() {
		cache_Main_modify1 = gopurs_runtime.Apply(Call_Main_modify(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadState())), Get_Main_monadStateState())
	})
	return cache_Main_modify1
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Call_Main_runState(gopurs_runtime.Str(""), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindState())), gopurs_runtime.Apply2(Call_Main_modify(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadState())), Get_Main_monadStateState(), gopurs_runtime.Func(func(__local_var_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(("World!") + (__local_var_0.StrVal()))
		})), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindState())), gopurs_runtime.Apply2(Call_Main_modify(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadState())), Get_Main_monadStateState(), gopurs_runtime.Func(func(__local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(("Hello, ") + (__local_var_1.StrVal()))
			})), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Main","State"] [String, String])
				__local_var_2_0 := Call_Main_get(Get_Main_monadStateState())
				_ = __local_var_2_0
				return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Main","Tuple"] [(TypeVar s$scope38), (TypeVar a$scope42)])
					v_4_1 := Call_Main_runState(s_3, __local_var_2_0)
					_ = v_4_1
					// TAST (Let): __local_var_5_2 shape=Other bindingType=(TypeVar a$scope42)
					__local_var_5_2 := (v_4_1).V1
					_ = __local_var_5_2
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Call_Main_runState((v_4_1).V0, gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_6, gopurs_runtime.Str(__local_var_5_2.StrVal())}))}
					})))}
				})
			}))
		}))))}
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

type Constructor_Main_Tuple[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}

type Constructor_Main_State[T_s any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_MonadState[T_s any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2980279296] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "get":
			return gopurs_runtime.Box(c.V0)
		case "put":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_MonadState: " + key)
		}
	}
}

func Call_Main_Tuple__3273409525(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop int64) *Constructor_Main_Tuple[int64, int64] {
Tuple__3273409525:
	for {
		if false {
			continue Tuple__3273409525
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Tuple[int64, int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_MonadState_dollar_Dict(x_0_loop struct {
	get gopurs_runtime.Value
	put gopurs_runtime.Value
}) *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		get gopurs_runtime.Value
		put gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("get", "put", orig.get, orig.put)
	}())
}

func Call_Main_test5(v_0_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_test4(v_0_loop *Constructor_Main_Tuple[int64, int64]) *Constructor_Main_Tuple[int64, int64] {
	var v_0 *Constructor_Main_Tuple[int64, int64] = v_0_loop
	_ = v_0
	return (&Constructor_Main_Tuple[int64, int64]{1, (v_0).V1, (v_0).V0})
}

func Call_Main_test3(n_0_loop int64) bool {
	var n_0 int64 = n_0_loop
	_ = n_0
	return (n_0) == (int64(0))
}

func Call_Main_test2(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, gopurs_runtime.Int(int64(10))).IntVal
}

func Call_Main_runState(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	var s_0 gopurs_runtime.Value = s_0_loop
	_ = s_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(v_1, s_0))
}

func Call_Main_put(dict_0_loop *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_get(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "get")
}

func Call_Main_modify(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope23)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
	_ = Bind1_1_0
	return gopurs_runtime.Func(func(dictMonadState_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): get2_3_1 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope23) [(TypeVar s$scope24)])
		get2_3_1 := Call_Main_get(dictMonadState_2)
		_ = get2_3_1
		return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Bind1_1_0.V1, get2_3_1, gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_2, "put"), gopurs_runtime.Apply(f_4, s_5))
			}))
		})
	})
}

func Rebox_Main_1152654068_4045719060(in *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_Tuple[int64, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[int64, int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_4045719060_1152654068(in *Constructor_Main_Tuple[int64, int64]) *Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}
