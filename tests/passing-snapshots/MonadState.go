package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

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
				Monad0 gopurs_runtime.Value
				get    gopurs_runtime.Value
				put    gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0 gopurs_runtime.Value
					get    gopurs_runtime.Value
					put    gopurs_runtime.Value
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.get = gopurs_runtime.RecordGet(orig, "get")
				clone.put = gopurs_runtime.RecordGet(orig, "put")
				return clone
			}()))}
		})
	})
	return cache_Main_MonadState_dollar_Dict
}

var cache_Main_showTuple gopurs_runtime.Value
var once_Main_showTuple sync.Once

func Get_Main_showTuple() gopurs_runtime.Value {
	once_Main_showTuple.Do(func() {
		cache_Main_showTuple = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showTuple(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Main_showTuple
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

var cache_Main_modify gopurs_runtime.Value
var once_Main_modify sync.Once

func Get_Main_modify() gopurs_runtime.Value {
	once_Main_modify.Do(func() {
		cache_Main_modify = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, dictMonadState_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_modify(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_1_box))
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
				// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Main","Tuple"] [(TypeVar s$scope39), (TypeVar a$scope43)])
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

var cache_Main_monadStateState gopurs_runtime.Value
var once_Main_monadStateState sync.Once

func Get_Main_monadStateState() gopurs_runtime.Value {
	once_Main_monadStateState.Do(func() {
		cache_Main_monadStateState = gopurs_runtime.Value{Type: 9, IntVal: 2980279296, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadState()))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_0, s_0}))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_0, Get_Data_Unit_unit()}))}
			})
		})}))}
	})
	return cache_Main_monadStateState
}

var cache_Main_modify__86672813 gopurs_runtime.Value
var once_Main_modify__86672813 sync.Once

func Get_Main_modify__86672813() gopurs_runtime.Value {
	once_Main_modify__86672813.Do(func() {
		cache_Main_modify__86672813 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_modify__86672813(__eta_norm_0_0_box)
		})
	})
	return cache_Main_modify__86672813
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Main_showTuple(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showUnit()))})), gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Call_Main_runState(gopurs_runtime.Int(int64(0)), Call_Main_modify__86672813(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int((v_0.IntVal) + (int64(1)))
		}))))}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
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
	V2 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2980279296] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "get":
			return gopurs_runtime.Box(c.V1)
		case "put":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Main_MonadState: " + key)
		}
	}
}

func Call_Main_MonadState_dollar_Dict(x_0_loop struct {
	Monad0 gopurs_runtime.Value
	get    gopurs_runtime.Value
	put    gopurs_runtime.Value
}) *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Monad0 gopurs_runtime.Value
		get    gopurs_runtime.Value
		put    gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Monad0", "get", "put", orig.Monad0, orig.get, orig.put)
	}())
}

func Call_Main_showTuple(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1668656255_1386611502((&Constructor_Data_Show_Show[*Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str((((("(") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (", ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
	})})))}
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
	return dict_0.V2
}

func Call_Main_get(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "get")
}

func Call_Main_modify(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], dictMonadState_1_loop *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
	_ = dictBind_0
	var dictMonadState_1 *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_1_loop
	_ = dictMonadState_1
	// TAST (Let): get1_2_0 shape=App(Var) bindingType=(TypeApp Any [Any])
	get1_2_0 := Call_Main_get(gopurs_runtime.Value{Type: 9, IntVal: 2980279296, UnsafePtr: unsafe.Pointer(dictMonadState_1)})
	_ = get1_2_0
	return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictBind_0.V1, get1_2_0, gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(dictMonadState_1.V2, gopurs_runtime.Apply2(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), f_3, s_4))
		}))
	})
}

func Call_Main_modify__86672813(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
modify__86672813:
	for {
		if false {
			continue modify__86672813
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(TypeApp Any [Any])
		__local_var_1_0 := Call_Main_get(gopurs_runtime.Value{Type: 9, IntVal: 2980279296, UnsafePtr: unsafe.Pointer(Rebox_Main_1994721641_934605138(Rebox_Main_934605138_1994721641(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_monadStateState()))))})
		_ = __local_var_1_0
		return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_3_1 shape=App(Var) bindingType=(ADT ["Main","Tuple"] [(TypeVar s$scope39), (TypeVar a$scope43)])
			v_3_1 := Call_Main_runState(s_2, __local_var_1_0)
			_ = v_3_1
			// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=Any
			__local_var_4_2 := gopurs_runtime.Apply2(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), __eta_norm_0_0, (v_3_1).V1)
			_ = __local_var_4_2
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Call_Main_runState((v_3_1).V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_4_2, Get_Data_Unit_unit()}))}
			})))}
		})
	}
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1668656255_1386611502(in *Constructor_Data_Show_Show[*Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1994721641_934605138(in *Constructor_Main_MonadState[int64, gopurs_runtime.Value]) *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	return out
}

func Rebox_Main_934605138_1994721641(in *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Main_MonadState[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_MonadState[int64, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	return out
}
