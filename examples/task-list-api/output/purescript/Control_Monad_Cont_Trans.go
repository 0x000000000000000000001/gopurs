package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Cont_Trans_compose gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_compose sync.Once

func Get_Control_Monad_Cont_Trans_compose() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_compose.Do(func() {
		cache_Control_Monad_Cont_Trans_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Cont_Trans_compose
}

var cache_Control_Monad_Cont_Trans_compose1 gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_compose1 sync.Once

func Get_Control_Monad_Cont_Trans_compose1() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_compose1.Do(func() {
		cache_Control_Monad_Cont_Trans_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Cont_Trans_compose1
}

var cache_Control_Monad_Cont_Trans_compose2 gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_compose2 sync.Once

func Get_Control_Monad_Cont_Trans_compose2() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_compose2.Do(func() {
		cache_Control_Monad_Cont_Trans_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Cont_Trans_compose2
}

var cache_Control_Monad_Cont_Trans_compose3 gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_compose3 sync.Once

func Get_Control_Monad_Cont_Trans_compose3() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_compose3.Do(func() {
		cache_Control_Monad_Cont_Trans_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Cont_Trans_compose3
}

var cache_Control_Monad_Cont_Trans_ContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_ContT sync.Once

func Get_Control_Monad_Cont_Trans_ContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_ContT.Do(func() {
		cache_Control_Monad_Cont_Trans_ContT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_ContT(x_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_ContT
}

var cache_Control_Monad_Cont_Trans_Functor_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_Functor_dollar_Dict sync.Once

func Get_Control_Monad_Cont_Trans_Functor_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_Functor_dollar_Dict.Do(func() {
		cache_Control_Monad_Cont_Trans_Functor_dollar_Dict = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v_1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(k_2, gopurs_runtime.Apply(f_0, a_3))
			}))
		})}))}
	})
	return cache_Control_Monad_Cont_Trans_Functor_dollar_Dict
}

var cache_Control_Monad_Cont_Trans_withContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_withContT sync.Once

func Get_Control_Monad_Cont_Trans_withContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_withContT.Do(func() {
		cache_Control_Monad_Cont_Trans_withContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_withContT(f_0_box, v_1_box, k_2_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_withContT
}

var cache_Control_Monad_Cont_Trans_runContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_runContT sync.Once

func Get_Control_Monad_Cont_Trans_runContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_runContT.Do(func() {
		cache_Control_Monad_Cont_Trans_runContT = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_runContT(v_0_box, k_1_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_runContT
}

var cache_Control_Monad_Cont_Trans_newtypeContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_newtypeContT sync.Once

func Get_Control_Monad_Cont_Trans_newtypeContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_newtypeContT.Do(func() {
		cache_Control_Monad_Cont_Trans_newtypeContT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Monad_Cont_Trans_newtypeContT
}

var cache_Control_Monad_Cont_Trans_monadTransContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadTransContT sync.Once

func Get_Control_Monad_Cont_Trans_monadTransContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadTransContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadTransContT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonad_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_1, k_2)
		})}))}
	})
	return cache_Control_Monad_Cont_Trans_monadTransContT
}

var cache_Control_Monad_Cont_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_lift sync.Once

func Get_Control_Monad_Cont_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_lift.Do(func() {
		cache_Control_Monad_Cont_Trans_lift = gopurs_runtime.RecordGet(Get_Control_Monad_Cont_Trans_monadTransContT(), "lift")
	})
	return cache_Control_Monad_Cont_Trans_lift
}

var cache_Control_Monad_Cont_Trans_lift1 gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_lift1 sync.Once

func Get_Control_Monad_Cont_Trans_lift1() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_lift1.Do(func() {
		cache_Control_Monad_Cont_Trans_lift1 = gopurs_runtime.RecordGet(Get_Control_Monad_Cont_Trans_monadTransContT(), "lift")
	})
	return cache_Control_Monad_Cont_Trans_lift1
}

var cache_Control_Monad_Cont_Trans_mapContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_mapContT sync.Once

func Get_Control_Monad_Cont_Trans_mapContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_mapContT.Do(func() {
		cache_Control_Monad_Cont_Trans_mapContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_mapContT(f_0_box, v_1_box, k_2_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_mapContT
}

var cache_Control_Monad_Cont_Trans_functorContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_functorContT sync.Once

func Get_Control_Monad_Cont_Trans_functorContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_functorContT.Do(func() {
		cache_Control_Monad_Cont_Trans_functorContT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_functorContT(dictFunctor_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_functorContT
}

var cache_Control_Monad_Cont_Trans_applyContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_applyContT sync.Once

func Get_Control_Monad_Cont_Trans_applyContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_applyContT.Do(func() {
		cache_Control_Monad_Cont_Trans_applyContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_applyContT(dictApply_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_applyContT
}

var cache_Control_Monad_Cont_Trans_bindContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_bindContT sync.Once

func Get_Control_Monad_Cont_Trans_bindContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_bindContT.Do(func() {
		cache_Control_Monad_Cont_Trans_bindContT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_bindContT(dictBind_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_bindContT
}

var cache_Control_Monad_Cont_Trans_semigroupContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_semigroupContT sync.Once

func Get_Control_Monad_Cont_Trans_semigroupContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_semigroupContT.Do(func() {
		cache_Control_Monad_Cont_Trans_semigroupContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_semigroupContT(dictApply_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_semigroupContT
}

var cache_Control_Monad_Cont_Trans_applicativeContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_applicativeContT sync.Once

func Get_Control_Monad_Cont_Trans_applicativeContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_applicativeContT.Do(func() {
		cache_Control_Monad_Cont_Trans_applicativeContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_applicativeContT(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_applicativeContT
}

var cache_Control_Monad_Cont_Trans_monadContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadContT sync.Once

func Get_Control_Monad_Cont_Trans_monadContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monadContT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monadContT
}

var cache_Control_Monad_Cont_Trans_monadAskContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadAskContT sync.Once

func Get_Control_Monad_Cont_Trans_monadAskContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadAskContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadAskContT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monadAskContT(dictMonadAsk_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monadAskContT
}

var cache_Control_Monad_Cont_Trans_monadReaderContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadReaderContT sync.Once

func Get_Control_Monad_Cont_Trans_monadReaderContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadReaderContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadReaderContT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monadReaderContT(dictMonadReader_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monadReaderContT
}

var cache_Control_Monad_Cont_Trans_monadContContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadContContT sync.Once

func Get_Control_Monad_Cont_Trans_monadContContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadContContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadContContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monadContContT(dictMonad_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monadContContT
}

var cache_Control_Monad_Cont_Trans_monadEffectContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadEffectContT sync.Once

func Get_Control_Monad_Cont_Trans_monadEffectContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadEffectContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadEffectContT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monadEffectContT(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monadEffectContT
}

var cache_Control_Monad_Cont_Trans_monadStateContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadStateContT sync.Once

func Get_Control_Monad_Cont_Trans_monadStateContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadStateContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadStateContT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monadStateContT(dictMonadState_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monadStateContT
}

var cache_Control_Monad_Cont_Trans_monadSTContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadSTContT sync.Once

func Get_Control_Monad_Cont_Trans_monadSTContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadSTContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadSTContT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monadSTContT(dictMonadST_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monadSTContT
}

var cache_Control_Monad_Cont_Trans_monoidContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monoidContT sync.Once

func Get_Control_Monad_Cont_Trans_monoidContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monoidContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monoidContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Cont_Trans_monoidContT(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_Cont_Trans_monoidContT
}

func Call_Control_Monad_Cont_Trans_ContT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_Cont_Trans_withContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var k_2 gopurs_runtime.Value = k_2_loop
	_ = k_2
	return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, k_2))
}

func Call_Control_Monad_Cont_Trans_runContT(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var k_1 gopurs_runtime.Value = k_1_loop
	_ = k_1
	return gopurs_runtime.Apply(v_0, k_1)
}

func Call_Control_Monad_Cont_Trans_mapContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var k_2 gopurs_runtime.Value = k_2_loop
	_ = k_2
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_Control_Monad_Cont_Trans_functorContT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_Cont_Trans_Functor_dollar_Dict()))}
}

func Call_Control_Monad_Cont_Trans_applyContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): functorContT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope166), (TypeVar m$scope165)])])
	functorContT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_functorContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})))
	_ = functorContT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_0)}
	}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
			}))
		}))
	})}))}
}

func Call_Control_Monad_Cont_Trans_bindContT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): applyContT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope153), (TypeVar m$scope152)])])
	applyContT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_applyContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})))
	_ = applyContT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_0)}
	}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, k_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(k_3, a_5, k_prime__4)
		}))
	})}))}
}

func Call_Control_Monad_Cont_Trans_semigroupContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Monad_Cont_Trans_applyContT(dictApply_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope27), (TypeVar a$scope27)] (TypeVar a$scope27))
		__local_var_3_1 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_2))
		_ = __local_var_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_3_1, a_4), b_5)
		})}))}
	})
}

func Call_Control_Monad_Cont_Trans_applicativeContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): applyContT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope179), (TypeVar m$scope178)])])
	applyContT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_applyContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})))
	_ = applyContT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_0)}
	}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(k_3, a_2)
	})}))}
}

func Call_Control_Monad_Cont_Trans_monadContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): applicativeContT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope100), (TypeVar m$scope99)])])
	applicativeContT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_applicativeContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = applicativeContT1_1_0
	// TAST (Let): bindContT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope100), (TypeVar m$scope99)])])
	bindContT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_bindContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})))
	_ = bindContT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_2_1)}
	})}))}
}

func Call_Control_Monad_Cont_Trans_monadAskContT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
	_ = dictMonadAsk_0
	// TAST (Let): monadContT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope124), (TypeVar m$scope123)])])
	monadContT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})))
	_ = monadContT1_1_0
	// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
	__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
	_ = __local_var_2_1
	// TAST (Let): __local_var_3_2 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope123) [(TypeVar r1$scope122)])
	__local_var_3_2 := Call_Control_Monad_Reader_Class_ask(dictMonadAsk_0)
	_ = __local_var_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_1_0)}
	}), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}), "bind"), __local_var_3_2, k_4)
	})}))}
}

func Call_Control_Monad_Cont_Trans_monadReaderContT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
	_ = dictMonadReader_0
	// TAST (Let): MonadAsk0_1_0 shape=App(Other) bindingType=Any
	MonadAsk0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
	_ = MonadAsk0_1_0
	// TAST (Let): ask_2_1 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope78) [(TypeVar r1$scope77)])
	ask_2_1 := Call_Control_Monad_Reader_Class_ask(MonadAsk0_1_0)
	_ = ask_2_1
	// TAST (Let): monadAskContT1_3_2 shape=App(Var) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r1$scope77), (TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope79), (TypeVar m$scope78)])])
	monadAskContT1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadAskContT(MonadAsk0_1_0))
	_ = monadAskContT1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskContT1_3_2)}
	}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadAsk0_1_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), ask_2_1, gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_4, gopurs_runtime.Apply(v_5, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return r_7
			})), k_6)))
		}))
	})}))}
}

func Call_Control_Monad_Cont_Trans_monadContContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): monadContT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope105), (TypeVar m$scope104)])])
	monadContT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadContT(dictMonad_0))
	_ = monadContT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_1_0)}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(k_3, a_4)
		}), k_3)
	})}))}
}

func Call_Control_Monad_Cont_Trans_monadEffectContT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadContT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope92), (TypeVar m$scope91)])])
	monadContT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadContT(Monad0_1_0))
	_ = monadContT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func2(func(m_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_3, k_4)
	}), Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_0)))}))}
}

func Call_Control_Monad_Cont_Trans_monadStateContT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
	_ = dictMonadState_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadContT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope68), (TypeVar m$scope67)])])
	monadContT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadContT(Monad0_1_0))
	_ = monadContT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func2(func(m_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_3, k_4)
	}), Call_Control_Monad_State_Class_state(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0)))}))}
}

func Call_Control_Monad_Cont_Trans_monadSTContT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
	_ = dictMonadST_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): monadContT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope5), (TypeVar m$scope4)])])
	monadContT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadContT(Monad0_1_0))
	_ = monadContT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func2(func(m_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_3, k_4)
	}), Call_Control_Monad_ST_Class_liftST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadST_0)))}))}
}

func Call_Control_Monad_Cont_Trans_monoidContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): semigroupContT1_1_0 shape=App(Var) bindingType=Any
	semigroupContT1_1_0 := Call_Control_Monad_Cont_Trans_semigroupContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
	_ = semigroupContT1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupContT2_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope50), (TypeVar m$scope48), (TypeVar a$scope49)])])
		semigroupContT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupContT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})))
		_ = semigroupContT2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupContT2_3_1)}
		}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Cont_Trans_applicativeContT(dictApplicative_0), "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty"))}))}
	})
}
