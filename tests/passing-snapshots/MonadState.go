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
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity(x_0_box)
		})
	})
	return cache_Main_identity
}

var cache_Main_Tuple gopurs_runtime.Value
var once_Main_Tuple sync.Once

func Get_Main_Tuple() gopurs_runtime.Value {
	once_Main_Tuple.Do(func() {
		cache_Main_Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, value0, value1}))}
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
			return Call_Main_MonadState_dollar_Dict(x_0_box)
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
			return Call_Main_put(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState](dict_0_box))
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
			return Call_Main_modify(dictBind_0_box, dictMonadState_1_box)
		})
	})
	return cache_Main_modify
}

var cache_Main_monadState gopurs_runtime.Value
var once_Main_monadState sync.Once

func Get_Main_monadState() gopurs_runtime.Value {
	once_Main_monadState.Do(func() {
		cache_Main_monadState = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeState()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindState()))}
		})}))}
	})
	return cache_Main_monadState
}

var cache_Main_functorState gopurs_runtime.Value
var once_Main_functorState sync.Once

func Get_Main_functorState() gopurs_runtime.Value {
	once_Main_functorState.Do(func() {
		cache_Main_functorState = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindState()).V1), a_1, gopurs_runtime.Func(func(a_prime__2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeState()).V1), gopurs_runtime.Apply(f_0, a_prime__2))
				}))
			})
		})}))}
	})
	return cache_Main_functorState
}

var cache_Main_bindState gopurs_runtime.Value
var once_Main_bindState sync.Once

func Get_Main_bindState() gopurs_runtime.Value {
	once_Main_bindState.Do(func() {
		cache_Main_bindState = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, (v_3_0).V1), (v_3_0).V0)))}
				})
			})
		})}))}
	})
	return cache_Main_bindState
}

var cache_Main_applyState gopurs_runtime.Value
var once_Main_applyState sync.Once

func Get_Main_applyState() gopurs_runtime.Value {
	once_Main_applyState.Do(func() {
		cache_Main_applyState = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					// TAST (Let): v_4_1 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(a_1, (v_3_0).V0))
					_ = v_4_1
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeState()).V1), gopurs_runtime.Apply((v_3_0).V1, (v_4_1).V1)), (v_4_1).V0)))}
				})
			})
		})}))}
	})
	return cache_Main_applyState
}

var cache_Main_applicativeState gopurs_runtime.Value
var once_Main_applicativeState sync.Once

func Get_Main_applicativeState() gopurs_runtime.Value {
	once_Main_applicativeState.Do(func() {
		cache_Main_applicativeState = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, s_1, a_0}))}
			})
		})}))}
	})
	return cache_Main_applicativeState
}

var cache_Main_monadStateState gopurs_runtime.Value
var once_Main_monadStateState sync.Once

func Get_Main_monadStateState() gopurs_runtime.Value {
	once_Main_monadStateState.Do(func() {
		cache_Main_monadStateState = gopurs_runtime.Value{Type: 9, IntVal: 2980279296, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadState{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Main_monadState()))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, s_0, s_0}))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, s_0, Get_Data_Unit_unit()}))}
			})
		})}))}
	})
	return cache_Main_monadStateState
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("(")+(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(1)).StrVal()))+(", unit)")))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_applicativeState__1389141106 gopurs_runtime.Value
var once_Main_applicativeState__1389141106 sync.Once

func Get_Main_applicativeState__1389141106() gopurs_runtime.Value {
	once_Main_applicativeState__1389141106.Do(func() {
		cache_Main_applicativeState__1389141106 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, s_1, a_0}))}
			})
		})}))}
	})
	return cache_Main_applicativeState__1389141106
}

var cache_Main_applicativeState__4174879383 gopurs_runtime.Value
var once_Main_applicativeState__4174879383 sync.Once

func Get_Main_applicativeState__4174879383() gopurs_runtime.Value {
	once_Main_applicativeState__4174879383.Do(func() {
		cache_Main_applicativeState__4174879383 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, s_1, a_0}))}
			})
		})}))}
	})
	return cache_Main_applicativeState__4174879383
}

var cache_Main_applyState__3033139442 gopurs_runtime.Value
var once_Main_applyState__3033139442 sync.Once

func Get_Main_applyState__3033139442() gopurs_runtime.Value {
	once_Main_applyState__3033139442.Do(func() {
		cache_Main_applyState__3033139442 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					// TAST (Let): v_4_1 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(a_1, (v_3_0).V0))
					_ = v_4_1
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, (v_4_1).V0, gopurs_runtime.Apply((v_3_0).V1, (v_4_1).V1)}))}
				})
			})
		})}))}
	})
	return cache_Main_applyState__3033139442
}

var cache_Main_applyState__3856467351 gopurs_runtime.Value
var once_Main_applyState__3856467351 sync.Once

func Get_Main_applyState__3856467351() gopurs_runtime.Value {
	once_Main_applyState__3856467351.Do(func() {
		cache_Main_applyState__3856467351 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					// TAST (Let): v_4_1 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(a_1, (v_3_0).V0))
					_ = v_4_1
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, (v_4_1).V0, gopurs_runtime.Apply((v_3_0).V1, (v_4_1).V1)}))}
				})
			})
		})}))}
	})
	return cache_Main_applyState__3856467351
}

var cache_Main_bindState__3056551538 gopurs_runtime.Value
var once_Main_bindState__3056551538 sync.Once

func Get_Main_bindState__3056551538() gopurs_runtime.Value {
	once_Main_bindState__3056551538.Do(func() {
		cache_Main_bindState__3056551538 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, (v_3_0).V1), (v_3_0).V0)))}
				})
			})
		})}))}
	})
	return cache_Main_bindState__3056551538
}

var cache_Main_bindState__1233139095 gopurs_runtime.Value
var once_Main_bindState__1233139095 sync.Once

func Get_Main_bindState__1233139095() gopurs_runtime.Value {
	once_Main_bindState__1233139095.Do(func() {
		cache_Main_bindState__1233139095 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, (v_3_0).V1), (v_3_0).V0)))}
				})
			})
		})}))}
	})
	return cache_Main_bindState__1233139095
}

var cache_Main_functorState__1056004645 gopurs_runtime.Value
var once_Main_functorState__1056004645 sync.Once

func Get_Main_functorState__1056004645() gopurs_runtime.Value {
	once_Main_functorState__1056004645.Do(func() {
		cache_Main_functorState__1056004645 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(a_1, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, (v_3_0).V0, gopurs_runtime.Apply(f_0, (v_3_0).V1)}))}
				})
			})
		})}))}
	})
	return cache_Main_functorState__1056004645
}

var cache_Main_functorState__2186776544 gopurs_runtime.Value
var once_Main_functorState__2186776544 sync.Once

func Get_Main_functorState__2186776544() gopurs_runtime.Value {
	once_Main_functorState__2186776544.Do(func() {
		cache_Main_functorState__2186776544 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) expectedFromAst=*Constructor_Main_Tuple actual=*Constructor_Main_Tuple bindingType=(ADT ["Main","Tuple"] [(TypeVar s64), (TypeVar a65)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(a_1, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, (v_3_0).V0, gopurs_runtime.Apply(f_0, (v_3_0).V1)}))}
				})
			})
		})}))}
	})
	return cache_Main_functorState__2186776544
}

var cache_Main_monadState__3046931442 gopurs_runtime.Value
var once_Main_monadState__3046931442 sync.Once

func Get_Main_monadState__3046931442() gopurs_runtime.Value {
	once_Main_monadState__3046931442.Do(func() {
		cache_Main_monadState__3046931442 = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeState()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindState()))}
		})}))}
	})
	return cache_Main_monadState__3046931442
}

var cache_Main_monadState__1513115415 gopurs_runtime.Value
var once_Main_monadState__1513115415 sync.Once

func Get_Main_monadState__1513115415() gopurs_runtime.Value {
	once_Main_monadState__1513115415.Do(func() {
		cache_Main_monadState__1513115415 = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeState()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindState()))}
		})}))}
	})
	return cache_Main_monadState__1513115415
}

var cache_Main_monadStateState__111897651 gopurs_runtime.Value
var once_Main_monadStateState__111897651 sync.Once

func Get_Main_monadStateState__111897651() gopurs_runtime.Value {
	once_Main_monadStateState__111897651.Do(func() {
		cache_Main_monadStateState__111897651 = gopurs_runtime.Value{Type: 9, IntVal: 2980279296, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadState{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Main_monadState()))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, s_0, s_0}))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple{1, s_0, Get_Data_Unit_unit()}))}
			})
		})}))}
	})
	return cache_Main_monadStateState__111897651
}

var cache_Main_put__693193917 gopurs_runtime.Value
var once_Main_put__693193917 sync.Once

func Get_Main_put__693193917() gopurs_runtime.Value {
	once_Main_put__693193917.Do(func() {
		cache_Main_put__693193917 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_put__693193917(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState](dict_0_box))
		})
	})
	return cache_Main_put__693193917
}

type Constructor_Main_Tuple struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Main_State struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_MonadState struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2980279296] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadState)(ptr)
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

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MonadState_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_showTuple(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str((((("(") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Main_Tuple)(v_2.UnsafePtr).V0).StrVal())) + (", ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Main_Tuple)(v_2.UnsafePtr).V1).StrVal())) + (")"))
	})}))}
}

func Call_Main_runState(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *Constructor_Main_Tuple {
	var s_0 gopurs_runtime.Value = s_0_loop
	_ = s_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple](gopurs_runtime.Apply(v_1, s_0))
}

func Call_Main_put(dict_0_loop *Constructor_Main_MonadState) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadState = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V2)
}

func Call_Main_get(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "get")
}

func Call_Main_modify(dictBind_0_loop gopurs_runtime.Value, dictMonadState_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	var dictMonadState_1 gopurs_runtime.Value = dictMonadState_1_loop
	_ = dictMonadState_1
	// TAST (Let): get1_2_0 shape=Other expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(TypeApp Any [Any])
	get1_2_0 := gopurs_runtime.RecordGet(dictMonadState_1, "get")
	_ = get1_2_0
	return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), get1_2_0, gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_1, "put"), gopurs_runtime.Apply(f_3, s_4))
		}))
	})
}

func Call_Main_put__693193917(dict_0_loop *Constructor_Main_MonadState) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadState = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V2)
}
