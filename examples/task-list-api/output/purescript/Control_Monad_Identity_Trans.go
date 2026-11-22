package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Identity_Trans_compose gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_compose sync.Once

func Get_Control_Monad_Identity_Trans_compose() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_compose.Do(func() {
		cache_Control_Monad_Identity_Trans_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Identity_Trans_compose
}

var cache_Control_Monad_Identity_Trans_compose1 gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_compose1 sync.Once

func Get_Control_Monad_Identity_Trans_compose1() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_compose1.Do(func() {
		cache_Control_Monad_Identity_Trans_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Identity_Trans_compose1
}

var cache_Control_Monad_Identity_Trans_IdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_IdentityT sync.Once

func Get_Control_Monad_Identity_Trans_IdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_IdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_IdentityT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_IdentityT(x_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_IdentityT
}

var cache_Control_Monad_Identity_Trans_monadSTIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadSTIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadSTIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadSTIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadSTIdentityT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadSTIdentityT(dictMonadST_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadSTIdentityT
}

var cache_Control_Monad_Identity_Trans_traversableIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_traversableIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_traversableIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_traversableIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_traversableIdentityT = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_traversableIdentityT(dictTraversable_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_traversableIdentityT
}

var cache_Control_Monad_Identity_Trans_runIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_runIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_runIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_runIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_runIdentityT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_runIdentityT(v_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_runIdentityT
}

var cache_Control_Monad_Identity_Trans_plusIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_plusIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_plusIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_plusIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_plusIdentityT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_plusIdentityT(dictPlus_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_plusIdentityT
}

var cache_Control_Monad_Identity_Trans_newtypeIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_newtypeIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_newtypeIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_newtypeIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_newtypeIdentityT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Monad_Identity_Trans_newtypeIdentityT
}

var cache_Control_Monad_Identity_Trans_monadWriterIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadWriterIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadWriterIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadWriterIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadWriterIdentityT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadWriterIdentityT(dictMonadWriter_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadWriterIdentityT
}

var cache_Control_Monad_Identity_Trans_monadTransIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadTransIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadTransIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadTransIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadTransIdentityT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_Identity_Trans_IdentityT()
		})}))}
	})
	return cache_Control_Monad_Identity_Trans_monadTransIdentityT
}

var cache_Control_Monad_Identity_Trans_monadThrowIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadThrowIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadThrowIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadThrowIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadThrowIdentityT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadThrowIdentityT(dictMonadThrow_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadThrowIdentityT
}

var cache_Control_Monad_Identity_Trans_monadTellIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadTellIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadTellIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadTellIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadTellIdentityT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadTellIdentityT(dictMonadTell_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadTellIdentityT
}

var cache_Control_Monad_Identity_Trans_monadStateIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadStateIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadStateIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadStateIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadStateIdentityT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadStateIdentityT(dictMonadState_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadStateIdentityT
}

var cache_Control_Monad_Identity_Trans_monadRecIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadRecIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadRecIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadRecIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadRecIdentityT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadRecIdentityT(dictMonadRec_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadRecIdentityT
}

var cache_Control_Monad_Identity_Trans_monadReaderIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadReaderIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadReaderIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadReaderIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadReaderIdentityT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadReaderIdentityT(dictMonadReader_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadReaderIdentityT
}

var cache_Control_Monad_Identity_Trans_monadPlusIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadPlusIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadPlusIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadPlusIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadPlusIdentityT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadPlusIdentityT(dictMonadPlus_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadPlusIdentityT
}

var cache_Control_Monad_Identity_Trans_monadIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadIdentityT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadIdentityT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadIdentityT
}

var cache_Control_Monad_Identity_Trans_monadErrorIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadErrorIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadErrorIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadErrorIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadErrorIdentityT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadErrorIdentityT(dictMonadError_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadErrorIdentityT
}

var cache_Control_Monad_Identity_Trans_monadEffectIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadEffectIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadEffectIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadEffectIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadEffectIdentityT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadEffectIdentityT(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadEffectIdentityT
}

var cache_Control_Monad_Identity_Trans_monadContIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadContIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadContIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadContIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadContIdentityT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadContIdentityT(dictMonadCont_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadContIdentityT
}

var cache_Control_Monad_Identity_Trans_monadAskIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_monadAskIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_monadAskIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_monadAskIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_monadAskIdentityT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_monadAskIdentityT(dictMonadAsk_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_monadAskIdentityT
}

var cache_Control_Monad_Identity_Trans_mapIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_mapIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_mapIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_mapIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_mapIdentityT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_mapIdentityT(f_0_box, v_1_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_mapIdentityT
}

var cache_Control_Monad_Identity_Trans_functorIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_functorIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_functorIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_functorIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_functorIdentityT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_functorIdentityT(dictFunctor_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_functorIdentityT
}

var cache_Control_Monad_Identity_Trans_foldableIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_foldableIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_foldableIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_foldableIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_foldableIdentityT = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_foldableIdentityT(dictFoldable_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_foldableIdentityT
}

var cache_Control_Monad_Identity_Trans_extendIdentityI gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_extendIdentityI sync.Once

func Get_Control_Monad_Identity_Trans_extendIdentityI() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_extendIdentityI.Do(func() {
		cache_Control_Monad_Identity_Trans_extendIdentityI = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_extendIdentityI(dictExtend_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_extendIdentityI
}

var cache_Control_Monad_Identity_Trans_eqIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_eqIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_eqIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_eqIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_eqIdentityT = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_eqIdentityT(dictEq1_0_box, dictEq_1_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_eqIdentityT
}

var cache_Control_Monad_Identity_Trans_ordIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_ordIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_ordIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_ordIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_ordIdentityT = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_ordIdentityT(dictOrd1_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_ordIdentityT
}

var cache_Control_Monad_Identity_Trans_eq1IdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_eq1IdentityT sync.Once

func Get_Control_Monad_Identity_Trans_eq1IdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_eq1IdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_eq1IdentityT = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_eq1IdentityT(dictEq1_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_eq1IdentityT
}

var cache_Control_Monad_Identity_Trans_ord1IdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_ord1IdentityT sync.Once

func Get_Control_Monad_Identity_Trans_ord1IdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_ord1IdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_ord1IdentityT = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_ord1IdentityT(dictOrd1_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_ord1IdentityT
}

var cache_Control_Monad_Identity_Trans_comonadIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_comonadIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_comonadIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_comonadIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_comonadIdentityT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_comonadIdentityT(dictComonad_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_comonadIdentityT
}

var cache_Control_Monad_Identity_Trans_bindIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_bindIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_bindIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_bindIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_bindIdentityT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_bindIdentityT(dictBind_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_bindIdentityT
}

var cache_Control_Monad_Identity_Trans_applyIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_applyIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_applyIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_applyIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_applyIdentityT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_applyIdentityT(dictApply_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_applyIdentityT
}

var cache_Control_Monad_Identity_Trans_applicativeIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_applicativeIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_applicativeIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_applicativeIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_applicativeIdentityT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_applicativeIdentityT(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_applicativeIdentityT
}

var cache_Control_Monad_Identity_Trans_alternativeIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_alternativeIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_alternativeIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_alternativeIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_alternativeIdentityT = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_alternativeIdentityT(dictAlternative_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_alternativeIdentityT
}

var cache_Control_Monad_Identity_Trans_altIdentityT gopurs_runtime.Value
var once_Control_Monad_Identity_Trans_altIdentityT sync.Once

func Get_Control_Monad_Identity_Trans_altIdentityT() gopurs_runtime.Value {
	once_Control_Monad_Identity_Trans_altIdentityT.Do(func() {
		cache_Control_Monad_Identity_Trans_altIdentityT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Identity_Trans_altIdentityT(dictAlt_0_box)
		})
	})
	return cache_Control_Monad_Identity_Trans_altIdentityT
}

func Call_Control_Monad_Identity_Trans_IdentityT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_Identity_Trans_monadSTIdentityT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
	_ = dictMonadST_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadST_0))}
}

func Call_Control_Monad_Identity_Trans_traversableIdentityT(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0))}
}

func Call_Control_Monad_Identity_Trans_runIdentityT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Control_Monad_Identity_Trans_plusIdentityT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
	_ = dictPlus_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_0))}
}

func Call_Control_Monad_Identity_Trans_monadWriterIdentityT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
	_ = dictMonadWriter_0
	return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadWriter_0))}
}

func Call_Control_Monad_Identity_Trans_monadThrowIdentityT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
	_ = dictMonadThrow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadThrow_0))}
}

func Call_Control_Monad_Identity_Trans_monadTellIdentityT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
	_ = dictMonadTell_0
	return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadTell_0))}
}

func Call_Control_Monad_Identity_Trans_monadStateIdentityT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
	_ = dictMonadState_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0))}
}

func Call_Control_Monad_Identity_Trans_monadRecIdentityT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0))}
}

func Call_Control_Monad_Identity_Trans_monadReaderIdentityT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
	_ = dictMonadReader_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadReader_0))}
}

func Call_Control_Monad_Identity_Trans_monadPlusIdentityT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
	_ = dictMonadPlus_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]](dictMonadPlus_0))}
}

func Call_Control_Monad_Identity_Trans_monadIdentityT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))}
}

func Call_Control_Monad_Identity_Trans_monadErrorIdentityT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
	_ = dictMonadError_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadError_0))}
}

func Call_Control_Monad_Identity_Trans_monadEffectIdentityT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0))}
}

func Call_Control_Monad_Identity_Trans_monadContIdentityT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
	_ = dictMonadCont_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dictMonadCont_0))}
}

func Call_Control_Monad_Identity_Trans_monadAskIdentityT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
	_ = dictMonadAsk_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadAsk_0))}
}

func Call_Control_Monad_Identity_Trans_mapIdentityT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Identity_Trans_functorIdentityT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0))}
}

func Call_Control_Monad_Identity_Trans_foldableIdentityT(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0))}
}

func Call_Control_Monad_Identity_Trans_extendIdentityI(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
	_ = dictExtend_0
	// TAST (Let): functorIdentityT1_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar m) [(TypeVar a)]) [(TypeVar w$scope91)])])
	functorIdentityT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorIdentityT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorIdentityT1_1_0)}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, Get_Control_Monad_Identity_Trans_IdentityT()), v_3)
	})}))}
}

func Call_Control_Monad_Identity_Trans_eqIdentityT(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	// TAST (Let): eq11_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope106) [(TypeVar a$scope107)]), (TypeApp (TypeVar m$scope106) [(TypeVar a$scope107)])] Boolean)
	eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))})
	_ = eq11_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(eq11_2_0, x_3, y_4)
	})}))}
}

func Call_Control_Monad_Identity_Trans_ordIdentityT(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): eqIdentityT1_1_0 shape=App(Var) bindingType=Any
	eqIdentityT1_1_0 := gopurs_runtime.Apply(Get_Control_Monad_Identity_Trans_eqIdentityT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
	_ = eqIdentityT1_1_0
	return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): compare11_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope14) [(TypeVar a$scope15)]), (TypeApp (TypeVar m$scope14) [(TypeVar a$scope15)])] (ADT ["Data","Ordering","Ordering"] []))
		compare11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_2))})
		_ = compare11_3_1
		// TAST (Let): eqIdentityT2_4_2 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar m) [(TypeVar a)]) [(TypeVar m$scope14), (TypeVar a$scope15)])])
		eqIdentityT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqIdentityT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{})))
		_ = eqIdentityT2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqIdentityT2_4_2)}
		}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(compare11_3_1, x_5, y_6)
		})}))}
	})
}

func Call_Control_Monad_Identity_Trans_eq1IdentityT(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Control_Monad_Identity_Trans_eqIdentityT(dictEq1_0, dictEq_1)))
	})}))}
}

func Call_Control_Monad_Identity_Trans_ord1IdentityT(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): ordIdentityT1_1_0 shape=App(Var) bindingType=Any
	ordIdentityT1_1_0 := Call_Control_Monad_Identity_Trans_ordIdentityT(dictOrd1_0)
	_ = ordIdentityT1_1_0
	// TAST (Let): eq1IdentityT1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (TypeApp (TypeVar m) [(TypeVar a)]) [(TypeVar m$scope18)])])
	eq1IdentityT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Control_Monad_Identity_Trans_eq1IdentityT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})))
	_ = eq1IdentityT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1IdentityT1_2_1)}
	}), gopurs_runtime.Func(func(dictOrd_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordIdentityT1_1_0, dictOrd_3)))
	})}))}
}

func Call_Control_Monad_Identity_Trans_comonadIdentityT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): extendIdentityI1_1_0 shape=App(Var) bindingType=(ADT ["Control","Extend","Extend"] [(TypeApp (TypeApp (TypeVar m) [(TypeVar a)]) [(TypeVar w$scope115)])])
	extendIdentityI1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Call_Control_Monad_Identity_Trans_extendIdentityI(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})))
	_ = extendIdentityI1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendIdentityI1_1_0)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Comonad_extract(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dictComonad_0)), Get_Control_Monad_Identity_Trans_runIdentityT())}))}
}

func Call_Control_Monad_Identity_Trans_bindIdentityT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0))}
}

func Call_Control_Monad_Identity_Trans_applyIdentityT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}
}

func Call_Control_Monad_Identity_Trans_applicativeIdentityT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}
}

func Call_Control_Monad_Identity_Trans_alternativeIdentityT(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
	_ = dictAlternative_0
	return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0))}
}

func Call_Control_Monad_Identity_Trans_altIdentityT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
	_ = dictAlt_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dictAlt_0))}
}
