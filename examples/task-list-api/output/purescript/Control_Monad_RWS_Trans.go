package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_RWS_Trans_compose gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_compose sync.Once

func Get_Control_Monad_RWS_Trans_compose() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_compose.Do(func() {
		cache_Control_Monad_RWS_Trans_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_RWS_Trans_compose
}

var cache_Control_Monad_RWS_Trans_compose1 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_compose1 sync.Once

func Get_Control_Monad_RWS_Trans_compose1() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_compose1.Do(func() {
		cache_Control_Monad_RWS_Trans_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_RWS_Trans_compose1
}

var cache_Control_Monad_RWS_Trans_RWSResult gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_RWSResult sync.Once

func Get_Control_Monad_RWS_Trans_RWSResult() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_RWSResult.Do(func() {
		cache_Control_Monad_RWS_Trans_RWSResult = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2}))}
				})
			})
		})
	})
	return cache_Control_Monad_RWS_Trans_RWSResult
}

var cache_Control_Monad_RWS_Trans_RWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_RWST sync.Once

func Get_Control_Monad_RWS_Trans_RWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_RWST.Do(func() {
		cache_Control_Monad_RWS_Trans_RWST = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_RWST(x_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_RWST
}

var cache_Control_Monad_RWS_Trans_withRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_withRWST sync.Once

func Get_Control_Monad_RWS_Trans_withRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_withRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_withRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_withRWST(f_0_box, m_1_box, r_2_box, s_3_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_withRWST
}

var cache_Control_Monad_RWS_Trans_runRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_runRWST sync.Once

func Get_Control_Monad_RWS_Trans_runRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_runRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_runRWST = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_runRWST(v_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_runRWST
}

var cache_Control_Monad_RWS_Trans_newtypeRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_newtypeRWST sync.Once

func Get_Control_Monad_RWS_Trans_newtypeRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_newtypeRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_newtypeRWST = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Monad_RWS_Trans_newtypeRWST
}

var cache_Control_Monad_RWS_Trans_monadTransRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadTransRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadTransRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadTransRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadTransRWST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadTransRWST
}

var cache_Control_Monad_RWS_Trans_mapRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_mapRWST sync.Once

func Get_Control_Monad_RWS_Trans_mapRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_mapRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_mapRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_mapRWST(f_0_box, v_1_box, r_2_box, s_3_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_mapRWST
}

var cache_Control_Monad_RWS_Trans_lazyRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_lazyRWST sync.Once

func Get_Control_Monad_RWS_Trans_lazyRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_lazyRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_lazyRWST = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, r_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), r_1, s_2)
		})}))}
	})
	return cache_Control_Monad_RWS_Trans_lazyRWST
}

var cache_Control_Monad_RWS_Trans_functorRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_functorRWST sync.Once

func Get_Control_Monad_RWS_Trans_functorRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_functorRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_functorRWST = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_functorRWST(dictFunctor_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_functorRWST
}

var cache_Control_Monad_RWS_Trans_execRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_execRWST sync.Once

func Get_Control_Monad_RWS_Trans_execRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_execRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_execRWST = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_execRWST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), v_1_box, r_2_box, s_3_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_execRWST
}

var cache_Control_Monad_RWS_Trans_evalRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_evalRWST sync.Once

func Get_Control_Monad_RWS_Trans_evalRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_evalRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_evalRWST = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_evalRWST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), v_1_box, r_2_box, s_3_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_evalRWST
}

var cache_Control_Monad_RWS_Trans_applyRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_applyRWST sync.Once

func Get_Control_Monad_RWS_Trans_applyRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_applyRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_applyRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_applyRWST
}

var cache_Control_Monad_RWS_Trans_bindRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_bindRWST sync.Once

func Get_Control_Monad_RWS_Trans_bindRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_bindRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_bindRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_bindRWST(dictBind_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_bindRWST
}

var cache_Control_Monad_RWS_Trans_semigroupRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_semigroupRWST sync.Once

func Get_Control_Monad_RWS_Trans_semigroupRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_semigroupRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_semigroupRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_semigroupRWST(dictBind_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_semigroupRWST
}

var cache_Control_Monad_RWS_Trans_applicativeRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_applicativeRWST sync.Once

func Get_Control_Monad_RWS_Trans_applicativeRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_applicativeRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_applicativeRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_applicativeRWST
}

var cache_Control_Monad_RWS_Trans_monadRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadRWST
}

var cache_Control_Monad_RWS_Trans_monadAskRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadAskRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadAskRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadAskRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadAskRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadAskRWST
}

var cache_Control_Monad_RWS_Trans_monadReaderRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadReaderRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadReaderRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadReaderRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadReaderRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadReaderRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadReaderRWST
}

var cache_Control_Monad_RWS_Trans_monadEffectRWS gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadEffectRWS sync.Once

func Get_Control_Monad_RWS_Trans_monadEffectRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadEffectRWS.Do(func() {
		cache_Control_Monad_RWS_Trans_monadEffectRWS = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadEffectRWS(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadEffectRWS
}

var cache_Control_Monad_RWS_Trans_monadRecRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadRecRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadRecRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadRecRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadRecRWST = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadRecRWST(dictMonadRec_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadRecRWST
}

var cache_Control_Monad_RWS_Trans_monadStateRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadStateRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadStateRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadStateRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadStateRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadStateRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadStateRWST
}

var cache_Control_Monad_RWS_Trans_monadTellRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadTellRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadTellRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadTellRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadTellRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadTellRWST
}

var cache_Control_Monad_RWS_Trans_monadWriterRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadWriterRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadWriterRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadWriterRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadWriterRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadWriterRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadWriterRWST
}

var cache_Control_Monad_RWS_Trans_monadThrowRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadThrowRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadThrowRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadThrowRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadThrowRWST = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadThrowRWST(dictMonadThrow_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadThrowRWST
}

var cache_Control_Monad_RWS_Trans_monadErrorRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadErrorRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadErrorRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadErrorRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadErrorRWST = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadErrorRWST(dictMonadError_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadErrorRWST
}

var cache_Control_Monad_RWS_Trans_monadSTRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadSTRWST sync.Once

func Get_Control_Monad_RWS_Trans_monadSTRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadSTRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadSTRWST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monadSTRWST(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monadSTRWST
}

var cache_Control_Monad_RWS_Trans_monoidRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monoidRWST sync.Once

func Get_Control_Monad_RWS_Trans_monoidRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monoidRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monoidRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_monoidRWST(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_monoidRWST
}

var cache_Control_Monad_RWS_Trans_altRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_altRWST sync.Once

func Get_Control_Monad_RWS_Trans_altRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_altRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_altRWST = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_altRWST(dictAlt_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_altRWST
}

var cache_Control_Monad_RWS_Trans_plusRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_plusRWST sync.Once

func Get_Control_Monad_RWS_Trans_plusRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_plusRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_plusRWST = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_plusRWST(dictPlus_0_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_plusRWST
}

var cache_Control_Monad_RWS_Trans_alternativeRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_alternativeRWST sync.Once

func Get_Control_Monad_RWS_Trans_alternativeRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_alternativeRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_alternativeRWST = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_Trans_alternativeRWST(dictMonoid_0_box, dictAlternative_1_box)
		})
	})
	return cache_Control_Monad_RWS_Trans_alternativeRWST
}

type Constructor_Control_Monad_RWS_Trans_RWSResult[T_state any, T_result any, T_writer any] struct {
	Rc uint32
	V0 T_state
	V1 T_result
	V2 T_writer
}

func Call_Control_Monad_RWS_Trans_RWST(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_RWS_Trans_withRWST(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var s_3 gopurs_runtime.Value = s_3_loop
	_ = s_3
	// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar r1$scope21), (TypeVar s$scope24)])
	__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, r_2, s_3))
	_ = __local_var_4_0
	return gopurs_runtime.Apply2(m_1, (__local_var_4_0).V0, (__local_var_4_0).V1)
}

func Call_Control_Monad_RWS_Trans_runRWST(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar w$scope128)
	mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): pure_3_1 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope130), (TypeVar a$scope135), (TypeVar w$scope128)])] (TypeApp (TypeVar m$scope134) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope130), (TypeVar a$scope135), (TypeVar w$scope128)])]))
		pure_3_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{})))
		_ = pure_3_1
		return gopurs_runtime.Func3(func(m_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "bind"), m_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_6, a_7, mempty_1_0}))})
			}))
		})
	})}))}
}

func Call_Control_Monad_RWS_Trans_mapRWST(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var s_3 gopurs_runtime.Value = s_3_loop
	_ = s_3
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(v_1, r_2, s_3))
}

func Call_Control_Monad_RWS_Trans_functorRWST(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2}))}
		}), gopurs_runtime.Apply2(v_2, r_3, s_4))
	})}))}
}

func Call_Control_Monad_RWS_Trans_execRWST(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var s_3 gopurs_runtime.Value = s_3_loop
	_ = s_3
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, r_2, s_3), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{(*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))})
	}))
}

func Call_Control_Monad_RWS_Trans_evalRWST(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var s_3 gopurs_runtime.Value = s_3_loop
	_ = s_3
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, r_2, s_3), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{(*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))})
	}))
}

func Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=Any
	Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
	_ = Apply0_1_0
	// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
	__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_2_1
	// TAST (Let): functorRWST1_3_2 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r$scope411), (TypeVar s$scope412)] (TypeApp (TypeVar m$scope409) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope412), (TypeVar a), (TypeVar w$scope410)])]))])
	functorRWST1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_functorRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})))
	_ = functorRWST1_3_2
	return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_3_2)}
		}), gopurs_runtime.Func4(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_5, r_7, s_8), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_10_3 shape=Other bindingType=Any
				__local_var_10_3 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V2
				_ = __local_var_10_3
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}), "append"), __local_var_10_3, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V2)}))}
				}), gopurs_runtime.Apply2(v1_6, r_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V0))
			}))
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_bindRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	// TAST (Let): applyRWST1_2_1 shape=App(Var) bindingType=Any
	applyRWST1_2_1 := Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0)
	_ = applyRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applyRWST2_4_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r$scope390), (TypeVar s$scope391)] (TypeApp (TypeVar m$scope388) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope391), (TypeVar a), (TypeVar w$scope389)])]))])
		applyRWST2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3))
		_ = applyRWST2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_4_2)}
		}), gopurs_runtime.Func4(func(v_5 gopurs_runtime.Value, f_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_5, r_7, s_8), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_10_3 shape=Other bindingType=Any
				__local_var_10_3 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V2
				_ = __local_var_10_3
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append"), __local_var_10_3, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V2)}))}
				}), gopurs_runtime.Apply3(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1, r_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0))
			}))
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_semigroupRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): applyRWST1_1_0 shape=App(Var) bindingType=Any
	applyRWST1_1_0 := Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0)
	_ = applyRWST1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
		__local_var_3_1 := gopurs_runtime.Apply(applyRWST1_1_0, dictMonoid_2)
		_ = __local_var_3_1
		return gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_2 shape=App(Var) bindingType=(Func [(TypeVar a$scope39), (TypeVar a$scope39)] (TypeVar a$scope39))
			__local_var_5_2 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_4))
			_ = __local_var_5_2
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_5_2, a_6), b_7)
			})}))}
		})
	})
}

func Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope433), (TypeVar a$scope436), (TypeVar w$scope431)])] (TypeApp (TypeVar m$scope430) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope433), (TypeVar a$scope436), (TypeVar w$scope431)])]))
	pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_1_0
	// TAST (Let): applyRWST1_2_1 shape=App(Var) bindingType=Any
	applyRWST1_2_1 := Call_Control_Monad_RWS_Trans_applyRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
	_ = applyRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_4_2 shape=Other bindingType=(TypeVar w$scope431)
		mempty_4_2 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
		_ = mempty_4_2
		// TAST (Let): applyRWST2_5_3 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r$scope432), (TypeVar s$scope433)] (TypeApp (TypeVar m$scope430) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope433), (TypeVar a), (TypeVar w$scope431)])]))])
		applyRWST2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3))
		_ = applyRWST2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_5_3)}
		}), gopurs_runtime.Func3(func(a_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_6, mempty_4_2}))})
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): applicativeRWST1_1_0 shape=App(Var) bindingType=Any
	applicativeRWST1_1_0 := Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0)
	_ = applicativeRWST1_1_0
	// TAST (Let): bindRWST1_2_1 shape=App(Var) bindingType=Any
	bindRWST1_2_1 := Call_Control_Monad_RWS_Trans_bindRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
	_ = bindRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applicativeRWST2_4_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r$scope250), (TypeVar s$scope251)] (TypeApp (TypeVar m$scope248) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope251), (TypeVar a), (TypeVar w$scope249)])]))])
		applicativeRWST2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3))
		_ = applicativeRWST2_4_2
		// TAST (Let): bindRWST2_5_3 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r$scope250), (TypeVar s$scope251)] (TypeApp (TypeVar m$scope248) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope251), (TypeVar a), (TypeVar w$scope249)])]))])
		bindRWST2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1_2_1, dictMonoid_3))
		_ = bindRWST2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_4_2)}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_5_3)}
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope291), (TypeVar r$scope290), (TypeVar w$scope289)])] (TypeApp (TypeVar m$scope288) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope291), (TypeVar r$scope290), (TypeVar w$scope289)])]))
	pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_1_0
	// TAST (Let): monadRWST1_2_1 shape=App(Var) bindingType=Any
	monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0)
	_ = monadRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_4_2 shape=Other bindingType=(TypeVar w$scope289)
		mempty_4_2 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
		_ = mempty_4_2
		// TAST (Let): monadRWST2_5_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r$scope290), (TypeVar s$scope291)] (TypeApp (TypeVar m$scope288) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope291), (TypeVar a), (TypeVar w$scope289)])]))])
		monadRWST2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3))
		_ = monadRWST2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_5_3)}
		}), gopurs_runtime.Func2(func(r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_7, r_6, mempty_4_2}))})
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadReaderRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): monadAskRWST1_1_0 shape=App(Var) bindingType=Any
	monadAskRWST1_1_0 := Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0)
	_ = monadAskRWST1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadAskRWST2_3_1 shape=App(Other) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r$scope233), (Func [(TypeVar r$scope233), (TypeVar s$scope234)] (TypeApp (TypeVar m$scope231) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope234), (TypeVar a), (TypeVar w$scope232)])]))])
		monadAskRWST2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadAskRWST1_1_0, dictMonoid_2))
		_ = monadAskRWST2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskRWST2_3_1)}
		}), gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(m_5, gopurs_runtime.Apply(f_4, r_6), s_7)
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadEffectRWS(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): lift_1_0 shape=App(Var) bindingType=Any
	lift_1_0 := Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_0)))
	_ = lift_1_0
	return gopurs_runtime.Func(func(dictMonadEffect_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Monad0_3_1 shape=App(Other) bindingType=Any
		Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "Monad0"), gopurs_runtime.Value{})
		_ = Monad0_3_1
		// TAST (Let): monadRWST1_4_2 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r$scope278), (TypeVar s$scope279)] (TypeApp (TypeVar m$scope277) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope279), (TypeVar a), (TypeVar w$scope276)])]))])
		monadRWST1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_monadRWST(Monad0_3_1), dictMonoid_0))
		_ = monadRWST1_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST1_4_2)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(lift_1_0, Monad0_3_1), Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_2)))}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadRecRWST(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadRWST1_2_1 shape=App(Var) bindingType=Any
	monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(Monad0_1_0)
	_ = monadRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=Any
		__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})
		_ = __local_var_4_2
		// TAST (Let): mempty_5_3 shape=Other bindingType=(TypeVar w$scope201)
		mempty_5_3 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
		_ = mempty_5_3
		// TAST (Let): monadRWST2_6_4 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r$scope202), (TypeVar s$scope203)] (TypeApp (TypeVar m$scope200) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope203), (TypeVar a), (TypeVar w$scope201)])]))])
		monadRWST2_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3))
		_ = monadRWST2_6_4
		return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_6_4)}
		}), gopurs_runtime.Func4(func(k_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_12_5 shape=Other bindingType=Any
				__local_var_12_5 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V2
				_ = __local_var_12_5
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply3(k_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1, r_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t8 gopurs_runtime.Value
					{
						var __t_tag_6 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1
						_ = __t_tag_6
						if __t_tag_6.Type == 9 && __t_tag_6.IntVal == 525585346 {
							__t8 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_RWS_Trans_2567023824_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "append"), __local_var_12_5, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2)})})))}
							goto end_branch_8
						} else {

						}
					}
					{
						var __t_tag_7 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1
						_ = __t_tag_7
						if __t_tag_7.Type == 9 && __t_tag_7.IntVal == 60402430 {
							__t8 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_RWS_Trans_4258103980_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "append"), __local_var_12_5, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2)})})))}
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_8:
					return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t8)
				}))
			}), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_10, a_8, mempty_5_3}))})
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadStateRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope182), (TypeVar a$scope187), (TypeVar w$scope181)])] (TypeApp (TypeVar m$scope180) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope182), (TypeVar a$scope187), (TypeVar w$scope181)])]))
	pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_1_0
	// TAST (Let): monadRWST1_2_1 shape=App(Var) bindingType=Any
	monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0)
	_ = monadRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_4_2 shape=Other bindingType=(TypeVar w$scope181)
		mempty_4_2 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
		_ = mempty_4_2
		// TAST (Let): monadRWST2_5_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r$scope183), (TypeVar s$scope182)] (TypeApp (TypeVar m$scope180) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope182), (TypeVar a), (TypeVar w$scope181)])]))])
		monadRWST2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3))
		_ = monadRWST2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_5_3)}
		}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_9_4 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope187), (TypeVar s$scope182)])
			v1_9_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_6, s_8))
			_ = v1_9_4
			return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_9_4).V1, (v1_9_4).V0, mempty_4_2}))})
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope165), Unit, (TypeVar w$scope163)])] (TypeApp (TypeVar m$scope162) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope165), Unit, (TypeVar w$scope163)])]))
	pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_1_0
	// TAST (Let): monadRWST1_2_1 shape=App(Var) bindingType=Any
	monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0)
	_ = monadRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Semigroup0_4_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w$scope163)])
		Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_4_2
		// TAST (Let): monadRWST2_5_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r$scope164), (TypeVar s$scope165)] (TypeApp (TypeVar m$scope162) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope165), (TypeVar a), (TypeVar w$scope163)])]))])
		monadRWST2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3))
		_ = monadRWST2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_5_3)}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_4_2)}
		}), gopurs_runtime.Func3(func(w_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, Get_Data_Unit_unit(), w_6}))})
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadWriterRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope97)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
	_ = Bind1_1_0
	// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=Any
	Applicative0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
	_ = Applicative0_2_1
	// TAST (Let): pure_3_2 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope100), (ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope105), (TypeVar w$scope98)]), (TypeVar w$scope98)])] (TypeApp (TypeVar m$scope97) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope100), (ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope105), (TypeVar w$scope98)]), (TypeVar w$scope98)])]))
	pure_3_2 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_2_1))
	_ = pure_3_2
	// TAST (Let): pure1_4_3 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope100), (TypeVar a$scope116), (TypeVar w$scope98)])] (TypeApp (TypeVar m$scope97) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope100), (TypeVar a$scope116), (TypeVar w$scope98)])]))
	pure1_4_3 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_2_1))
	_ = pure1_4_3
	// TAST (Let): monadTellRWST1_5_4 shape=App(Var) bindingType=Any
	monadTellRWST1_5_4 := Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0)
	_ = monadTellRWST1_5_4
	return gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadTellRWST2_7_5 shape=App(Other) bindingType=(ADT ["Control","Monad","Writer","Class","MonadTell"] [(TypeVar w$scope98), (Func [(TypeVar r$scope99), (TypeVar s$scope100)] (TypeApp (TypeVar m$scope97) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope100), (TypeVar a), (TypeVar w$scope98)])]))])
		monadTellRWST2_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadTellRWST1_5_4, dictMonoid_6))
		_ = monadTellRWST2_7_5
		return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellRWST2_7_5)}
		}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return dictMonoid_6
		}), gopurs_runtime.Func3(func(m_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(m_8, r_9, s_10), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_RWS_Trans_1947641166_2326785955((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V2}), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V2})))})
			}))
		}), gopurs_runtime.Func3(func(m_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(m_8, r_9, s_10), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(pure1_4_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V2)}))})
			}))
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadThrowRWST(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
	_ = dictMonadThrow_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m$scope150)])
	Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
	_ = Monad0_1_0
	// TAST (Let): monadRWST1_2_1 shape=App(Var) bindingType=Any
	monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
	_ = monadRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): lift_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope150) [(TypeVar a$scope157)]), (TypeVar r$scope152), (TypeVar s$scope153)] (TypeApp (TypeVar m$scope150) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope153), (TypeVar a$scope157), (TypeVar w$scope151)])]))
		lift_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_3), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)})
		_ = lift_4_2
		// TAST (Let): monadRWST2_5_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r$scope152), (TypeVar s$scope153)] (TypeApp (TypeVar m$scope150) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope153), (TypeVar a), (TypeVar w$scope151)])]))])
		monadRWST2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3))
		_ = monadRWST2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_5_3)}
		}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(lift_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_6))
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadErrorRWST(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
	_ = dictMonadError_0
	// TAST (Let): monadThrowRWST1_1_0 shape=App(Var) bindingType=Any
	monadThrowRWST1_1_0 := Call_Control_Monad_RWS_Trans_monadThrowRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
	_ = monadThrowRWST1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadThrowRWST2_3_1 shape=App(Other) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e$scope258), (Func [(TypeVar r$scope261), (TypeVar s$scope262)] (TypeApp (TypeVar m$scope259) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope262), (TypeVar a), (TypeVar w$scope260)])]))])
		monadThrowRWST2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadThrowRWST1_1_0, dictMonoid_2))
		_ = monadThrowRWST2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowRWST2_3_1)}
		}), gopurs_runtime.Func4(func(m_4 gopurs_runtime.Value, h_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply2(m_4, r_6, s_7), gopurs_runtime.Func(func(e_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply3(h_5, e_8, r_6, s_7)
			}))
		})}))}
	})
}

func Call_Control_Monad_RWS_Trans_monadSTRWST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): lift_1_0 shape=App(Var) bindingType=Any
	lift_1_0 := Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_0)))
	_ = lift_1_0
	return gopurs_runtime.Func(func(dictMonadST_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Monad0_3_1 shape=App(Other) bindingType=Any
		Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "Monad0"), gopurs_runtime.Value{})
		_ = Monad0_3_1
		// TAST (Let): monadRWST1_4_2 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r$scope8), (TypeVar s'$scope9)] (TypeApp (TypeVar m$scope7) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s'$scope9), (TypeVar a), (TypeVar w$scope5)])]))])
		monadRWST1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_monadRWST(Monad0_3_1), dictMonoid_0))
		_ = monadRWST1_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST1_4_2)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(lift_1_0, Monad0_3_1), Call_Control_Monad_ST_Class_liftST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadST_2)))}))}
	})
}

func Call_Control_Monad_RWS_Trans_monoidRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): applicativeRWST1_1_0 shape=App(Var) bindingType=Any
	applicativeRWST1_1_0 := Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0)
	_ = applicativeRWST1_1_0
	// TAST (Let): semigroupRWST1_2_1 shape=App(Var) bindingType=Any
	semigroupRWST1_2_1 := Call_Control_Monad_RWS_Trans_semigroupRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
	_ = semigroupRWST1_2_1
	return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupRWST2_4_2 shape=App(Other) bindingType=Any
		semigroupRWST2_4_2 := gopurs_runtime.Apply(semigroupRWST1_2_1, dictMonoid_3)
		_ = semigroupRWST2_4_2
		return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): semigroupRWST3_6_3 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar r$scope90), (TypeVar s$scope91)] (TypeApp (TypeVar m$scope87) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope91), (TypeVar a$scope89), (TypeVar w$scope88)])]))])
			semigroupRWST3_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupRWST2_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_5, "Semigroup0"), gopurs_runtime.Value{})))
			_ = semigroupRWST3_6_3
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupRWST3_6_3)}
			}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3), "pure"), gopurs_runtime.RecordGet(dictMonoid1_5, "mempty"))}))}
		})
	})
}

func Call_Control_Monad_RWS_Trans_altRWST(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
	_ = dictAlt_0
	// TAST (Let): functorRWST1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r$scope459), (TypeVar s$scope461)] (TypeApp (TypeVar m$scope458) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope461), (TypeVar a), (TypeVar w$scope460)])]))])
	functorRWST1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_functorRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})))
	_ = functorRWST1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_1_0)}
	}), gopurs_runtime.Func4(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply2(v_2, r_4, s_5), gopurs_runtime.Apply2(v1_3, r_4, s_5))
	})}))}
}

func Call_Control_Monad_RWS_Trans_plusRWST(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
	_ = dictPlus_0
	// TAST (Let): empty_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope57) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope60), (TypeVar a$scope63), (TypeVar w$scope59)])])
	empty_1_0 := Call_Control_Plus_empty(dictPlus_0)
	_ = empty_1_0
	// TAST (Let): altRWST1_2_1 shape=App(Var) bindingType=(ADT ["Control","Alt","Alt"] [(Func [(TypeVar r$scope58), (TypeVar s$scope60)] (TypeApp (TypeVar m$scope57) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope60), (TypeVar a), (TypeVar w$scope59)])]))])
	altRWST1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_altRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})))
	_ = altRWST1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altRWST1_2_1)}
	}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return empty_1_0
	})}))}
}

func Call_Control_Monad_RWS_Trans_alternativeRWST(dictMonoid_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
	_ = dictAlternative_1
	// TAST (Let): plusRWST1_2_0 shape=App(Var) bindingType=(ADT ["Control","Plus","Plus"] [(Func [(TypeVar r$scope451), (TypeVar s$scope452)] (TypeApp (TypeVar m$scope450) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope452), (TypeVar a), (TypeVar w$scope449)])]))])
	plusRWST1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_plusRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{})))
	_ = plusRWST1_2_0
	return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applicativeRWST1_4_1 shape=App(Var) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r$scope451), (TypeVar s$scope452)] (TypeApp (TypeVar m$scope450) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope452), (TypeVar a), (TypeVar w$scope449)])]))])
		applicativeRWST1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_3), dictMonoid_0))
		_ = applicativeRWST1_4_1
		return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST1_4_1)}
		}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusRWST1_2_0)}
		})}))}
	})
}

func Rebox_Control_Monad_RWS_Trans_1947641166_2326785955(in *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V1)}
	out.V2 = in.V2
	return out
}

func Rebox_Control_Monad_RWS_Trans_2567023824_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_RWS_Trans_4258103980_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}
