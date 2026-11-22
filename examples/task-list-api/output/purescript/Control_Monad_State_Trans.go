package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_State_Trans_compose gopurs_runtime.Value
var once_Control_Monad_State_Trans_compose sync.Once

func Get_Control_Monad_State_Trans_compose() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_compose.Do(func() {
		cache_Control_Monad_State_Trans_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_Trans_compose
}

var cache_Control_Monad_State_Trans_compose1 gopurs_runtime.Value
var once_Control_Monad_State_Trans_compose1 sync.Once

func Get_Control_Monad_State_Trans_compose1() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_compose1.Do(func() {
		cache_Control_Monad_State_Trans_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_Trans_compose1
}

var cache_Control_Monad_State_Trans_compose2 gopurs_runtime.Value
var once_Control_Monad_State_Trans_compose2 sync.Once

func Get_Control_Monad_State_Trans_compose2() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_compose2.Do(func() {
		cache_Control_Monad_State_Trans_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_Trans_compose2
}

var cache_Control_Monad_State_Trans_compose3 gopurs_runtime.Value
var once_Control_Monad_State_Trans_compose3 sync.Once

func Get_Control_Monad_State_Trans_compose3() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_compose3.Do(func() {
		cache_Control_Monad_State_Trans_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_Trans_compose3
}

var cache_Control_Monad_State_Trans_compose4 gopurs_runtime.Value
var once_Control_Monad_State_Trans_compose4 sync.Once

func Get_Control_Monad_State_Trans_compose4() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_compose4.Do(func() {
		cache_Control_Monad_State_Trans_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_Trans_compose4
}

var cache_Control_Monad_State_Trans_compose5 gopurs_runtime.Value
var once_Control_Monad_State_Trans_compose5 sync.Once

func Get_Control_Monad_State_Trans_compose5() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_compose5.Do(func() {
		cache_Control_Monad_State_Trans_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_Trans_compose5
}

var cache_Control_Monad_State_Trans_compose6 gopurs_runtime.Value
var once_Control_Monad_State_Trans_compose6 sync.Once

func Get_Control_Monad_State_Trans_compose6() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_compose6.Do(func() {
		cache_Control_Monad_State_Trans_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_Trans_compose6
}

var cache_Control_Monad_State_Trans_StateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_StateT sync.Once

func Get_Control_Monad_State_Trans_StateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_StateT.Do(func() {
		cache_Control_Monad_State_Trans_StateT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_StateT(x_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_StateT
}

var cache_Control_Monad_State_Trans_withStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_withStateT sync.Once

func Get_Control_Monad_State_Trans_withStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_withStateT.Do(func() {
		cache_Control_Monad_State_Trans_withStateT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_withStateT(f_0_box, v_1_box)
		})
	})
	return cache_Control_Monad_State_Trans_withStateT
}

var cache_Control_Monad_State_Trans_runStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_runStateT sync.Once

func Get_Control_Monad_State_Trans_runStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_runStateT.Do(func() {
		cache_Control_Monad_State_Trans_runStateT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_runStateT(v_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_runStateT
}

var cache_Control_Monad_State_Trans_newtypeStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_newtypeStateT sync.Once

func Get_Control_Monad_State_Trans_newtypeStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_newtypeStateT.Do(func() {
		cache_Control_Monad_State_Trans_newtypeStateT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Monad_State_Trans_newtypeStateT
}

var cache_Control_Monad_State_Trans_monadTransStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadTransStateT sync.Once

func Get_Control_Monad_State_Trans_monadTransStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadTransStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadTransStateT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope92), (TypeVar s$scope87)])] (TypeApp (TypeVar m$scope91) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope92), (TypeVar s$scope87)])]))
			pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
			_ = pure_1_0
			return gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{x_4, s_3}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))})
				}))
			})
		})}))}
	})
	return cache_Control_Monad_State_Trans_monadTransStateT
}

var cache_Control_Monad_State_Trans_lift gopurs_runtime.Value
var once_Control_Monad_State_Trans_lift sync.Once

func Get_Control_Monad_State_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_lift.Do(func() {
		cache_Control_Monad_State_Trans_lift = gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift")
	})
	return cache_Control_Monad_State_Trans_lift
}

var cache_Control_Monad_State_Trans_lift1 gopurs_runtime.Value
var once_Control_Monad_State_Trans_lift1 sync.Once

func Get_Control_Monad_State_Trans_lift1() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_lift1.Do(func() {
		cache_Control_Monad_State_Trans_lift1 = gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift")
	})
	return cache_Control_Monad_State_Trans_lift1
}

var cache_Control_Monad_State_Trans_lift2 gopurs_runtime.Value
var once_Control_Monad_State_Trans_lift2 sync.Once

func Get_Control_Monad_State_Trans_lift2() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_lift2.Do(func() {
		cache_Control_Monad_State_Trans_lift2 = gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift")
	})
	return cache_Control_Monad_State_Trans_lift2
}

var cache_Control_Monad_State_Trans_mapStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_mapStateT sync.Once

func Get_Control_Monad_State_Trans_mapStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_mapStateT.Do(func() {
		cache_Control_Monad_State_Trans_mapStateT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_mapStateT(f_0_box, v_1_box)
		})
	})
	return cache_Control_Monad_State_Trans_mapStateT
}

var cache_Control_Monad_State_Trans_lazyStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_lazyStateT sync.Once

func Get_Control_Monad_State_Trans_lazyStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_lazyStateT.Do(func() {
		cache_Control_Monad_State_Trans_lazyStateT = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_0, Get_Data_Unit_unit(), s_1)
		})}))}
	})
	return cache_Control_Monad_State_Trans_lazyStateT
}

var cache_Control_Monad_State_Trans_functorStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_functorStateT sync.Once

func Get_Control_Monad_State_Trans_functorStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_functorStateT.Do(func() {
		cache_Control_Monad_State_Trans_functorStateT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_functorStateT(dictFunctor_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_functorStateT
}

var cache_Control_Monad_State_Trans_execStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_execStateT sync.Once

func Get_Control_Monad_State_Trans_execStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_execStateT.Do(func() {
		cache_Control_Monad_State_Trans_execStateT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_execStateT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box, s_2_box)
		})
	})
	return cache_Control_Monad_State_Trans_execStateT
}

var cache_Control_Monad_State_Trans_evalStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_evalStateT sync.Once

func Get_Control_Monad_State_Trans_evalStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_evalStateT.Do(func() {
		cache_Control_Monad_State_Trans_evalStateT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_evalStateT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box, s_2_box)
		})
	})
	return cache_Control_Monad_State_Trans_evalStateT
}

var cache_Control_Monad_State_Trans_monadStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadStateT sync.Once

func Get_Control_Monad_State_Trans_monadStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadStateT
}

var cache_Control_Monad_State_Trans_bindStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_bindStateT sync.Once

func Get_Control_Monad_State_Trans_bindStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_bindStateT.Do(func() {
		cache_Control_Monad_State_Trans_bindStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_bindStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_bindStateT
}

var cache_Control_Monad_State_Trans_applyStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_applyStateT sync.Once

func Get_Control_Monad_State_Trans_applyStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_applyStateT.Do(func() {
		cache_Control_Monad_State_Trans_applyStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_applyStateT
}

var cache_Control_Monad_State_Trans_applicativeStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_applicativeStateT sync.Once

func Get_Control_Monad_State_Trans_applicativeStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_applicativeStateT.Do(func() {
		cache_Control_Monad_State_Trans_applicativeStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_applicativeStateT
}

var cache_Control_Monad_State_Trans_semigroupStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_semigroupStateT sync.Once

func Get_Control_Monad_State_Trans_semigroupStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_semigroupStateT.Do(func() {
		cache_Control_Monad_State_Trans_semigroupStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_semigroupStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_semigroupStateT
}

var cache_Control_Monad_State_Trans_monadAskStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadAskStateT sync.Once

func Get_Control_Monad_State_Trans_monadAskStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadAskStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadAskStateT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadAskStateT(dictMonadAsk_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadAskStateT
}

var cache_Control_Monad_State_Trans_monadReaderStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadReaderStateT sync.Once

func Get_Control_Monad_State_Trans_monadReaderStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadReaderStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadReaderStateT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadReaderStateT(dictMonadReader_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadReaderStateT
}

var cache_Control_Monad_State_Trans_monadContStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadContStateT sync.Once

func Get_Control_Monad_State_Trans_monadContStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadContStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadContStateT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadContStateT(dictMonadCont_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadContStateT
}

var cache_Control_Monad_State_Trans_monadEffectState gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadEffectState sync.Once

func Get_Control_Monad_State_Trans_monadEffectState() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadEffectState.Do(func() {
		cache_Control_Monad_State_Trans_monadEffectState = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadEffectState(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadEffectState
}

var cache_Control_Monad_State_Trans_monadRecStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadRecStateT sync.Once

func Get_Control_Monad_State_Trans_monadRecStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadRecStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadRecStateT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadRecStateT(dictMonadRec_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadRecStateT
}

var cache_Control_Monad_State_Trans_monadStateStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadStateStateT sync.Once

func Get_Control_Monad_State_Trans_monadStateStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadStateStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadStateStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadStateStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadStateStateT
}

var cache_Control_Monad_State_Trans_monadTellStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadTellStateT sync.Once

func Get_Control_Monad_State_Trans_monadTellStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadTellStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadTellStateT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadTellStateT(dictMonadTell_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadTellStateT
}

var cache_Control_Monad_State_Trans_monadWriterStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadWriterStateT sync.Once

func Get_Control_Monad_State_Trans_monadWriterStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadWriterStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadWriterStateT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadWriterStateT(dictMonadWriter_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadWriterStateT
}

var cache_Control_Monad_State_Trans_monadThrowStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadThrowStateT sync.Once

func Get_Control_Monad_State_Trans_monadThrowStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadThrowStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadThrowStateT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadThrowStateT(dictMonadThrow_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadThrowStateT
}

var cache_Control_Monad_State_Trans_monadErrorStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadErrorStateT sync.Once

func Get_Control_Monad_State_Trans_monadErrorStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadErrorStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadErrorStateT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadErrorStateT(dictMonadError_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadErrorStateT
}

var cache_Control_Monad_State_Trans_monadSTStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadSTStateT sync.Once

func Get_Control_Monad_State_Trans_monadSTStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadSTStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadSTStateT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadSTStateT(dictMonadST_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadSTStateT
}

var cache_Control_Monad_State_Trans_monoidStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monoidStateT sync.Once

func Get_Control_Monad_State_Trans_monoidStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monoidStateT.Do(func() {
		cache_Control_Monad_State_Trans_monoidStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monoidStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monoidStateT
}

var cache_Control_Monad_State_Trans_altStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_altStateT sync.Once

func Get_Control_Monad_State_Trans_altStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_altStateT.Do(func() {
		cache_Control_Monad_State_Trans_altStateT = gopurs_runtime.Func2(func(dictMonad_0_box gopurs_runtime.Value, dictAlt_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_altStateT(dictMonad_0_box, dictAlt_1_box)
		})
	})
	return cache_Control_Monad_State_Trans_altStateT
}

var cache_Control_Monad_State_Trans_plusStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_plusStateT sync.Once

func Get_Control_Monad_State_Trans_plusStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_plusStateT.Do(func() {
		cache_Control_Monad_State_Trans_plusStateT = gopurs_runtime.Func2(func(dictMonad_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_plusStateT(dictMonad_0_box, dictPlus_1_box)
		})
	})
	return cache_Control_Monad_State_Trans_plusStateT
}

var cache_Control_Monad_State_Trans_alternativeStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_alternativeStateT sync.Once

func Get_Control_Monad_State_Trans_alternativeStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_alternativeStateT.Do(func() {
		cache_Control_Monad_State_Trans_alternativeStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_alternativeStateT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_alternativeStateT
}

var cache_Control_Monad_State_Trans_monadPlusStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadPlusStateT sync.Once

func Get_Control_Monad_State_Trans_monadPlusStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadPlusStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadPlusStateT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Trans_monadPlusStateT(dictMonadPlus_0_box)
		})
	})
	return cache_Control_Monad_State_Trans_monadPlusStateT
}

func Call_Control_Monad_State_Trans_StateT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_State_Trans_withStateT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_1, f_0)
}

func Call_Control_Monad_State_Trans_runStateT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Control_Monad_State_Trans_mapStateT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, v_1)
}

func Call_Control_Monad_State_Trans_functorStateT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.Apply(v_2, s_3))
	})}))}
}

func Call_Control_Monad_State_Trans_execStateT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, Get_Data_Tuple_snd(), gopurs_runtime.Apply(v_1, s_2))
}

func Call_Control_Monad_State_Trans_evalStateT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, Get_Data_Tuple_fst(), gopurs_runtime.Apply(v_1, s_2))
}

func Call_Control_Monad_State_Trans_monadStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
	}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_State_Trans_bindStateT(dictMonad_0)
	})}))}
}

func Call_Control_Monad_State_Trans_bindStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
	}), gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(v_1, s_3), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1)
		}))
	})}))}
}

func Call_Control_Monad_State_Trans_applyStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): functorStateT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar s$scope283)] (TypeApp (TypeVar m$scope282) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope283)])]))])
	functorStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_functorStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
	_ = functorStateT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStateT1_1_0)}
	}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(dictMonad_0)))}))}
}

func Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope295), (TypeVar s$scope292)])] (TypeApp (TypeVar m$scope291) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope295), (TypeVar s$scope292)])]))
	pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
	}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{a_2, s_3}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))})
	})}))}
}

func Call_Control_Monad_State_Trans_semigroupStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope25), (TypeVar a$scope25)] (TypeVar a$scope25))
		__local_var_3_1 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_2))
		_ = __local_var_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_3_1, a_4), b_5)
		})}))}
	})
}

func Call_Control_Monad_State_Trans_monadAskStateT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
	_ = dictMonadAsk_0
	// TAST (Let): monadStateT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope204)] (TypeApp (TypeVar m$scope203) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope204)])]))])
	monadStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})))
	_ = monadStateT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_1_0)}
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})))}, Call_Control_Monad_Reader_Class_ask(dictMonadAsk_0))}))}
}

func Call_Control_Monad_State_Trans_monadReaderStateT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
	_ = dictMonadReader_0
	// TAST (Let): monadAskStateT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r$scope224), (Func [(TypeVar s$scope226)] (TypeApp (TypeVar m$scope225) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope226)])]))])
	monadAskStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadAskStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})))
	_ = monadAskStateT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskStateT1_1_0)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Control_Monad_State_Trans_mapStateT(), Call_Control_Monad_Reader_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadReader_0)))}))}
}

func Call_Control_Monad_State_Trans_monadContStateT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
	_ = dictMonadCont_0
	// TAST (Let): monadStateT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope185)] (TypeApp (TypeVar m$scope184) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope185)])]))])
	monadStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})))
	_ = monadStateT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_1_0)}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, s_prime__6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{a_5, s_prime__6}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))})
			}), s_3)
		}))
	})}))}
}

func Call_Control_Monad_State_Trans_monadEffectState(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadStateT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope177)] (TypeApp (TypeVar m$scope176) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope177)])]))])
	monadStateT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(Monad0_1_0))
	_ = monadStateT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), Monad0_1_0), Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0)))}))}
}

func Call_Control_Monad_State_Trans_monadRecStateT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadStateT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope137)] (TypeApp (TypeVar m$scope136) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope137)])]))])
	monadStateT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(Monad0_1_0))
	_ = monadStateT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_2_1)}
	}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0
					_ = __t_tag_2
					if __t_tag_2.Type == 9 && __t_tag_2.IntVal == 525585346 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_State_Trans_3131450224_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V1})})))}
						goto end_branch_4
					} else {

					}
				}
				{
					var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0
					_ = __t_tag_3
					if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 60402430 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_State_Trans_2886445004_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V1})})))}
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t4)
			}))
		}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{a_4, s_5}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))})
	})}))}
}

func Call_Control_Monad_State_Trans_monadStateStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope129), (TypeVar s$scope125)])] (TypeApp (TypeVar m$scope124) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope129), (TypeVar s$scope125)])]))
	pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_1_0
	// TAST (Let): monadStateT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope125)] (TypeApp (TypeVar m$scope124) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope125)])]))])
	monadStateT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(dictMonad_0))
	_ = monadStateT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_2_1)}
	}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), pure_1_0, f_3)
	})}))}
}

func Call_Control_Monad_State_Trans_monadTellStateT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
	_ = dictMonadTell_0
	// TAST (Let): Monad1_1_0 shape=App(Other) bindingType=Any
	Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
	_ = Monad1_1_0
	// TAST (Let): Semigroup0_2_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w$scope111)])
	Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{}))
	_ = Semigroup0_2_1
	// TAST (Let): monadStateT1_3_2 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope113)] (TypeApp (TypeVar m$scope112) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope113)])]))])
	monadStateT1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(Monad1_1_0))
	_ = monadStateT1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_3_2)}
	}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), Monad1_1_0), Call_Control_Monad_Writer_Class_tell(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadTell_0)))}))}
}

func Call_Control_Monad_State_Trans_monadWriterStateT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
	_ = dictMonadWriter_0
	// TAST (Let): MonadTell1_1_0 shape=App(Other) bindingType=Any
	MonadTell1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{})
	_ = MonadTell1_1_0
	// TAST (Let): Monad1_2_1 shape=App(Other) bindingType=Any
	Monad1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
	_ = Monad1_2_1
	// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope64)])
	Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{}))
	_ = Bind1_3_2
	// TAST (Let): Applicative0_4_3 shape=App(Other) bindingType=Any
	Applicative0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{})
	_ = Applicative0_4_3
	// TAST (Let): pure_5_4 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope70), (TypeVar w$scope63)]), (TypeVar s$scope65)])] (TypeApp (TypeVar m$scope64) [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope70), (TypeVar w$scope63)]), (TypeVar s$scope65)])]))
	pure_5_4 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_4_3))
	_ = pure_5_4
	// TAST (Let): pure1_6_5 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope78), (TypeVar s$scope65)]), (Func [(TypeVar w$scope63)] (TypeVar w$scope63))])] (TypeApp (TypeVar m$scope64) [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope78), (TypeVar s$scope65)]), (Func [(TypeVar w$scope63)] (TypeVar w$scope63))])]))
	pure1_6_5 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_4_3))
	_ = pure1_6_5
	// TAST (Let): Monoid0_7_6 shape=App(Other) bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar w$scope63)])
	Monoid0_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{}))
	_ = Monoid0_7_6
	// TAST (Let): monadTellStateT1_8_7 shape=App(Var) bindingType=(ADT ["Control","Monad","Writer","Class","MonadTell"] [(TypeVar w$scope63), (Func [(TypeVar s$scope65)] (TypeApp (TypeVar m$scope64) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope65)])]))])
	monadTellStateT1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadTellStateT(MonadTell1_1_0))
	_ = monadTellStateT1_8_7
	return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellStateT1_8_7)}
	}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Monoid0_7_6)}
	}), gopurs_runtime.Func2(func(m_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), gopurs_runtime.Apply(m_9, s_10)), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_5_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_State_Trans_1785332133_138441832(Rebox_Control_Monad_State_Trans_138441832_1785332133(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1}))}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))})
		}))
	}), gopurs_runtime.Func2(func(m_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply(m_9, s_10), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure1_6_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_State_Trans_1785332133_138441832(Rebox_Control_Monad_State_Trans_138441832_1785332133(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1}))}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))})
		})))
	})}))}
}

func Call_Control_Monad_State_Trans_monadThrowStateT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
	_ = dictMonadThrow_0
	// TAST (Let): lift3_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope102) [(TypeVar a$scope107)]), (TypeVar s$scope103)] (TypeApp (TypeVar m$scope102) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope107), (TypeVar s$scope103)])]))
	lift3_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})))})
	_ = lift3_1_0
	// TAST (Let): monadStateT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope103)] (TypeApp (TypeVar m$scope102) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope103)])]))])
	monadStateT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})))
	_ = monadStateT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_2_1)}
	}), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(lift3_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_3))
	})}))}
}

func Call_Control_Monad_State_Trans_monadErrorStateT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
	_ = dictMonadError_0
	// TAST (Let): monadThrowStateT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e$scope164), (Func [(TypeVar s$scope166)] (TypeApp (TypeVar m$scope165) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope166)])]))])
	monadThrowStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadThrowStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{})))
	_ = monadThrowStateT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowStateT1_1_0)}
	}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply(v_2, s_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(h_3, e_5, s_4)
		}))
	})}))}
}

func Call_Control_Monad_State_Trans_monadSTStateT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
	_ = dictMonadST_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadStateT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s'$scope5)] (TypeApp (TypeVar m$scope4) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s'$scope5)])]))])
	monadStateT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(Monad0_1_0))
	_ = monadStateT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), Monad0_1_0), Call_Control_Monad_ST_Class_liftST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadST_0)))}))}
}

func Call_Control_Monad_State_Trans_monoidStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): semigroupStateT1_1_0 shape=App(Var) bindingType=Any
	semigroupStateT1_1_0 := Call_Control_Monad_State_Trans_semigroupStateT(dictMonad_0)
	_ = semigroupStateT1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupStateT2_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar s$scope58)] (TypeApp (TypeVar m$scope56) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope57), (TypeVar s$scope58)])]))])
		semigroupStateT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupStateT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})))
		_ = semigroupStateT2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupStateT2_3_1)}
		}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0), "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty"))}))}
	})
}

func Call_Control_Monad_State_Trans_altStateT(dictMonad_0_loop gopurs_runtime.Value, dictAlt_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	var dictAlt_1 gopurs_runtime.Value = dictAlt_1_loop
	_ = dictAlt_1
	// TAST (Let): functorStateT1_2_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar s$scope309)] (TypeApp (TypeVar m$scope308) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope309)])]))])
	functorStateT1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_functorStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_1, "Functor0"), gopurs_runtime.Value{})))
	_ = functorStateT1_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStateT1_2_0)}
	}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_1, "alt"), gopurs_runtime.Apply(v_3, s_5), gopurs_runtime.Apply(v1_4, s_5))
	})}))}
}

func Call_Control_Monad_State_Trans_plusStateT(dictMonad_0_loop gopurs_runtime.Value, dictPlus_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	var dictPlus_1 gopurs_runtime.Value = dictPlus_1_loop
	_ = dictPlus_1
	// TAST (Let): empty_2_0 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope36) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope40), (TypeVar s$scope37)])])
	empty_2_0 := Call_Control_Plus_empty(dictPlus_1)
	_ = empty_2_0
	// TAST (Let): altStateT2_3_1 shape=App(Var) bindingType=(ADT ["Control","Alt","Alt"] [(Func [(TypeVar s$scope37)] (TypeApp (TypeVar m$scope36) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope37)])]))])
	altStateT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_altStateT(dictMonad_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{})))
	_ = altStateT2_3_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altStateT2_3_1)}
	}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return empty_2_0
	})}))}
}

func Call_Control_Monad_State_Trans_alternativeStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): applicativeStateT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar s$scope304)] (TypeApp (TypeVar m$scope303) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope304)])]))])
	applicativeStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0))
	_ = applicativeStateT1_1_0
	return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): plusStateT2_3_1 shape=App(Var) bindingType=(ADT ["Control","Plus","Plus"] [(Func [(TypeVar s$scope304)] (TypeApp (TypeVar m$scope303) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope304)])]))])
		plusStateT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_plusStateT(dictMonad_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})))
		_ = plusStateT2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeStateT1_1_0)}
		}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusStateT2_3_1)}
		})}))}
	})
}

func Call_Control_Monad_State_Trans_monadPlusStateT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
	_ = dictMonadPlus_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadStateT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar s$scope159)] (TypeApp (TypeVar m$scope158) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope159)])]))])
	monadStateT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadStateT(Monad0_1_0))
	_ = monadStateT1_2_1
	// TAST (Let): alternativeStateT1_3_2 shape=App(Var) bindingType=(ADT ["Control","Alternative","Alternative"] [(Func [(TypeVar s$scope159)] (TypeApp (TypeVar m$scope158) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope159)])]))])
	alternativeStateT1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_State_Trans_alternativeStateT(Monad0_1_0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{})))
	_ = alternativeStateT1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeStateT1_3_2)}
	}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStateT1_2_1)}
	})}))}
}

func Rebox_Control_Monad_State_Trans_138441832_1785332133(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_State_Trans_1785332133_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_State_Trans_2886445004_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_State_Trans_3131450224_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}
