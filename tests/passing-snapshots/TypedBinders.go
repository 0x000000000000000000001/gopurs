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
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(Call_Main_test4(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[int64, int64]](v_0_box)))}
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

var cache_Main_put__693193917 gopurs_runtime.Value
var once_Main_put__693193917 sync.Once

func Get_Main_put__693193917() gopurs_runtime.Value {
	once_Main_put__693193917.Do(func() {
		cache_Main_put__693193917 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_put__693193917(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_put__693193917
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
		cache_Main_get1 = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_0, s_0}))}
		})
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
		cache_Main_functorState = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindState()).V1), a_1, gopurs_runtime.Func(func(a_prime__2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeState()).V1), gopurs_runtime.Apply(f_0, a_prime__2))
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
		cache_Main_bindState = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, (v_3_0).V1), (v_3_0).V0)))}
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
		cache_Main_applyState = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					// TAST (Let): v_4_1 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(a_1, (v_3_0).V0))
					_ = v_4_1
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeState()).V1), gopurs_runtime.Apply((v_3_0).V1, (v_4_1).V1)), (v_4_1).V0)))}
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

var cache_Main_functorState__2186776544 gopurs_runtime.Value
var once_Main_functorState__2186776544 sync.Once

func Get_Main_functorState__2186776544() gopurs_runtime.Value {
	once_Main_functorState__2186776544.Do(func() {
		cache_Main_functorState__2186776544 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(a_1, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_3_0).V0, gopurs_runtime.Apply(f_0, (v_3_0).V1)}))}
				})
			})
		})}))}
	})
	return cache_Main_functorState__2186776544
}

var cache_Main_bindState__1148353140 gopurs_runtime.Value
var once_Main_bindState__1148353140 sync.Once

func Get_Main_bindState__1148353140() gopurs_runtime.Value {
	once_Main_bindState__1148353140.Do(func() {
		cache_Main_bindState__1148353140 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, (v_3_0).V1), (v_3_0).V0)))}
				})
			})
		})}))}
	})
	return cache_Main_bindState__1148353140
}

var cache_Main_bindState__1233139095 gopurs_runtime.Value
var once_Main_bindState__1233139095 sync.Once

func Get_Main_bindState__1233139095() gopurs_runtime.Value {
	once_Main_bindState__1233139095.Do(func() {
		cache_Main_bindState__1233139095 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, (v_3_0).V1), (v_3_0).V0)))}
				})
			})
		})}))}
	})
	return cache_Main_bindState__1233139095
}

var cache_Main_applyState__3856467351 gopurs_runtime.Value
var once_Main_applyState__3856467351 sync.Once

func Get_Main_applyState__3856467351() gopurs_runtime.Value {
	once_Main_applyState__3856467351.Do(func() {
		cache_Main_applyState__3856467351 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorState()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, s_2))
					_ = v_3_0
					// TAST (Let): v_4_1 shape=App(Other) bindingType=(ADT ["Main","Tuple"] [(TypeVar s58), (TypeVar a59)])
					v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(a_1, (v_3_0).V0))
					_ = v_4_1
					return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_4_1).V0, gopurs_runtime.Apply((v_3_0).V1, (v_4_1).V1)}))}
				})
			})
		})}))}
	})
	return cache_Main_applyState__3856467351
}

var cache_Main_applicativeState__512641780 gopurs_runtime.Value
var once_Main_applicativeState__512641780 sync.Once

func Get_Main_applicativeState__512641780() gopurs_runtime.Value {
	once_Main_applicativeState__512641780.Do(func() {
		cache_Main_applicativeState__512641780 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_1, a_0}))}
			})
		})}))}
	})
	return cache_Main_applicativeState__512641780
}

var cache_Main_applicativeState__4174879383 gopurs_runtime.Value
var once_Main_applicativeState__4174879383 sync.Once

func Get_Main_applicativeState__4174879383() gopurs_runtime.Value {
	once_Main_applicativeState__4174879383.Do(func() {
		cache_Main_applicativeState__4174879383 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyState()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, s_1, a_0}))}
			})
		})}))}
	})
	return cache_Main_applicativeState__4174879383
}

var cache_Main_modify1 gopurs_runtime.Value
var once_Main_modify1 sync.Once

func Get_Main_modify1() gopurs_runtime.Value {
	once_Main_modify1.Do(func() {
		cache_Main_modify1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_modify1(f_0_box)
		})
	})
	return cache_Main_modify1
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=Other bindingType=Any
			__local_var_0_0 := gopurs_runtime.Str(("Hello, ") + ("World!"))
			_ = __local_var_0_0
			return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0_0, gopurs_runtime.Str(__local_var_0_0.StrVal())}))}
		}()
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
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
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
		c := (*Constructor_Main_MonadState[any, any])(ptr)
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

func Call_Main_MonadState_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
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
	return (n_0) == (0)
}

func Call_Main_test2(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, gopurs_runtime.Int(10)).IntVal
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
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_put__693193917(dict_0_loop *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_get(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "get")
}

func Call_Main_modify(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
	_ = Bind1_1_0
	return gopurs_runtime.Func(func(dictMonadState_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): get2_3_1 shape=Other bindingType=(TypeApp (TypeVar m) [(TypeVar s)])
		get2_3_1 := gopurs_runtime.RecordGet(dictMonadState_2, "get")
		_ = get2_3_1
		return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), get2_3_1, gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_2, "put"), gopurs_runtime.Apply(f_4, s_5))
			}))
		})
	})
}

func Call_Main_modify1(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3562159846, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, s_1), Get_Data_Unit_unit()}))}
	})
}
