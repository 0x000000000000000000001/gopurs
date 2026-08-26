package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_MonadWriter_dollar_Dict gopurs_runtime.Value
var once_Main_MonadWriter_dollar_Dict sync.Once

func Get_Main_MonadWriter_dollar_Dict() gopurs_runtime.Value {
	once_Main_MonadWriter_dollar_Dict.Do(func() {
		cache_Main_MonadWriter_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MonadWriter_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MonadWriter_dollar_Dict
}

var cache_Main_MTrace gopurs_runtime.Value
var once_Main_MTrace sync.Once

func Get_Main_MTrace() gopurs_runtime.Value {
	once_Main_MTrace.Do(func() {
		cache_Main_MTrace = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_MTrace
}

var cache_Main_testFunctor gopurs_runtime.Value
var once_Main_testFunctor sync.Once

func Get_Main_testFunctor() gopurs_runtime.Value {
	once_Main_testFunctor.Do(func() {
		cache_Main_testFunctor = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testFunctor(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_testFunctor
}

var cache_Main_tell gopurs_runtime.Value
var once_Main_tell sync.Once

func Get_Main_tell() gopurs_runtime.Value {
	once_Main_tell.Do(func() {
		cache_Main_tell = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_tell(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_tell
}

var cache_Main_tell__2495349857 gopurs_runtime.Value
var once_Main_tell__2495349857 sync.Once

func Get_Main_tell__2495349857() gopurs_runtime.Value {
	once_Main_tell__2495349857.Do(func() {
		cache_Main_tell__2495349857 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_tell__2495349857(gopurs_runtime.CoerceToStruct[Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_tell__2495349857
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_test
}

var cache_Main_test__952106832 gopurs_runtime.Value
var once_Main_test__952106832 sync.Once

func Get_Main_test__952106832() gopurs_runtime.Value {
	once_Main_test__952106832.Do(func() {
		cache_Main_test__952106832 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test__952106832(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_test__952106832
}

var cache_Main_runMTrace gopurs_runtime.Value
var once_Main_runMTrace sync.Once

func Get_Main_runMTrace() gopurs_runtime.Value {
	once_Main_runMTrace.Do(func() {
		cache_Main_runMTrace = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runMTrace(v_0_box)
		})
	})
	return cache_Main_runMTrace
}

var cache_Main_runMTrace__426026716 gopurs_runtime.Value
var once_Main_runMTrace__426026716 sync.Once

func Get_Main_runMTrace__426026716() gopurs_runtime.Value {
	once_Main_runMTrace__426026716.Do(func() {
		cache_Main_runMTrace__426026716 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runMTrace__426026716(v_0_box)
		})
	})
	return cache_Main_runMTrace__426026716
}

var cache_Main_runMTrace__474729148 gopurs_runtime.Value
var once_Main_runMTrace__474729148 sync.Once

func Get_Main_runMTrace__474729148() gopurs_runtime.Value {
	once_Main_runMTrace__474729148.Do(func() {
		cache_Main_runMTrace__474729148 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runMTrace__474729148(v_0_box)
		})
	})
	return cache_Main_runMTrace__474729148
}

var cache_Main_runMTrace__823161127 gopurs_runtime.Value
var once_Main_runMTrace__823161127 sync.Once

func Get_Main_runMTrace__823161127() gopurs_runtime.Value {
	once_Main_runMTrace__823161127.Do(func() {
		cache_Main_runMTrace__823161127 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runMTrace__823161127(v_0_box)
		})
	})
	return cache_Main_runMTrace__823161127
}

var cache_Main_monadMTrace gopurs_runtime.Value
var once_Main_monadMTrace sync.Once

func Get_Main_monadMTrace() gopurs_runtime.Value {
	once_Main_monadMTrace.Do(func() {
		cache_Main_monadMTrace = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeMTrace()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindMTrace()))}
		})}))}
	})
	return cache_Main_monadMTrace
}

var cache_Main_functorMTrace gopurs_runtime.Value
var once_Main_functorMTrace sync.Once

func Get_Main_functorMTrace() gopurs_runtime.Value {
	once_Main_functorMTrace.Do(func() {
		cache_Main_functorMTrace = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindMTrace()).V1), a_1, gopurs_runtime.Func(func(a_prime__2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeMTrace()).V1), gopurs_runtime.Apply(f_0, a_prime__2))
				}))
			})
		})}))}
	})
	return cache_Main_functorMTrace
}

var cache_Main_bindMTrace gopurs_runtime.Value
var once_Main_bindMTrace sync.Once

func Get_Main_bindMTrace() gopurs_runtime.Value {
	once_Main_bindMTrace.Do(func() {
		cache_Main_bindMTrace = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyMTrace()))}
		}), gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
					__local_var_2_0 := m_0
					_ = __local_var_2_0
					x_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
					_ = x_3_1
					return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, x_3_1), gopurs_runtime.Value{})
				})
			})
		})}))}
	})
	return cache_Main_bindMTrace
}

var cache_Main_applyMTrace gopurs_runtime.Value
var once_Main_applyMTrace sync.Once

func Get_Main_applyMTrace() gopurs_runtime.Value {
	once_Main_applyMTrace.Do(func() {
		cache_Main_applyMTrace = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorMTrace()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
					__local_var_2_0 := f_0
					_ = __local_var_2_0
					x_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
					_ = x_3_1
					x_4_2 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
					_ = x_4_2
					return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeMTrace()).V1), gopurs_runtime.Apply(x_3_1, x_4_2)), gopurs_runtime.Value{})
				})
			})
		})}))}
	})
	return cache_Main_applyMTrace
}

var cache_Main_applicativeMTrace gopurs_runtime.Value
var once_Main_applicativeMTrace sync.Once

func Get_Main_applicativeMTrace() gopurs_runtime.Value {
	once_Main_applicativeMTrace.Do(func() {
		cache_Main_applicativeMTrace = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyMTrace()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return x_0
			})
		})}))}
	})
	return cache_Main_applicativeMTrace
}

var cache_Main_writerMTrace gopurs_runtime.Value
var once_Main_writerMTrace sync.Once

func Get_Main_writerMTrace() gopurs_runtime.Value {
	once_Main_writerMTrace.Do(func() {
		cache_Main_writerMTrace = gopurs_runtime.Value{Type: 9, IntVal: 2544837208, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadWriter[string, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Main_monadMTrace()))}
		}), gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(s_0.StrVal()))
		})}))}
	})
	return cache_Main_writerMTrace
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal()))
			_ = __local_var_0_0
			x_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = x_1_1
			x_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal())), gopurs_runtime.Value{})
			_ = x_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_MTrace[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_MonadWriter[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2544837208] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadWriter[any, any])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "tell":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_MonadWriter: " + key)
		}
	}
}

func Call_Main_MonadWriter_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_testFunctor(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func(func(n_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(__local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((1.0) + (__local_var_3.FloatVal()))
		}), n_2)
	})
}

func Call_Main_tell(dict_0_loop *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_tell__2495349857(dict_0_loop *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_test(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
	_ = Bind1_1_0
	return gopurs_runtime.Func(func(dictMonadWriter_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_2, "tell"), w_3), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_2, "tell"), w_3), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_2, "tell"), w_3)
				}))
			}))
		})
	})
}

func Call_Main_test__952106832(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
	_ = Bind1_1_0
	return gopurs_runtime.Func(func(dictMonadWriter_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_2, "tell"), w_3), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_2, "tell"), w_3), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_2, "tell"), w_3)
				}))
			}))
		})
	})
}

func Call_Main_runMTrace(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_runMTrace__426026716(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_runMTrace__474729148(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_runMTrace__823161127(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}
