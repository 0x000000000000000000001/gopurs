package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_Trans_compose gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_compose sync.Once

func Get_Control_Monad_Writer_Trans_compose() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_compose.Do(func() {
		cache_Control_Monad_Writer_Trans_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Writer_Trans_compose
}

var cache_Control_Monad_Writer_Trans_compose1 gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_compose1 sync.Once

func Get_Control_Monad_Writer_Trans_compose1() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_compose1.Do(func() {
		cache_Control_Monad_Writer_Trans_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Writer_Trans_compose1
}

var cache_Control_Monad_Writer_Trans_compose2 gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_compose2 sync.Once

func Get_Control_Monad_Writer_Trans_compose2() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_compose2.Do(func() {
		cache_Control_Monad_Writer_Trans_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Writer_Trans_compose2
}

var cache_Control_Monad_Writer_Trans_compose3 gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_compose3 sync.Once

func Get_Control_Monad_Writer_Trans_compose3() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_compose3.Do(func() {
		cache_Control_Monad_Writer_Trans_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Writer_Trans_compose3
}

var cache_Control_Monad_Writer_Trans_WriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_WriterT sync.Once

func Get_Control_Monad_Writer_Trans_WriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_WriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_WriterT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_WriterT(x_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_WriterT
}

var cache_Control_Monad_Writer_Trans_runWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_runWriterT sync.Once

func Get_Control_Monad_Writer_Trans_runWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_runWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_runWriterT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_runWriterT(v_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_runWriterT
}

var cache_Control_Monad_Writer_Trans_newtypeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_newtypeWriterT sync.Once

func Get_Control_Monad_Writer_Trans_newtypeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_newtypeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_newtypeWriterT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Monad_Writer_Trans_newtypeWriterT
}

var cache_Control_Monad_Writer_Trans_monadTransWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadTransWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadTransWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadTransWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadTransWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadTransWriterT
}

var cache_Control_Monad_Writer_Trans_mapWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_mapWriterT sync.Once

func Get_Control_Monad_Writer_Trans_mapWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_mapWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_mapWriterT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_mapWriterT(f_0_box, v_1_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_mapWriterT
}

var cache_Control_Monad_Writer_Trans_functorWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_functorWriterT sync.Once

func Get_Control_Monad_Writer_Trans_functorWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_functorWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_functorWriterT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_functorWriterT(dictFunctor_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_functorWriterT
}

var cache_Control_Monad_Writer_Trans_execWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_execWriterT sync.Once

func Get_Control_Monad_Writer_Trans_execWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_execWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_execWriterT = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_execWriterT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_execWriterT
}

var cache_Control_Monad_Writer_Trans_applyWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_applyWriterT sync.Once

func Get_Control_Monad_Writer_Trans_applyWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_applyWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_applyWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0_box, dictApply_1_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_applyWriterT
}

var cache_Control_Monad_Writer_Trans_bindWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_bindWriterT sync.Once

func Get_Control_Monad_Writer_Trans_bindWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_bindWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_bindWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictBind_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_bindWriterT(dictSemigroup_0_box, dictBind_1_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_bindWriterT
}

var cache_Control_Monad_Writer_Trans_semigroupWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_semigroupWriterT sync.Once

func Get_Control_Monad_Writer_Trans_semigroupWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_semigroupWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_semigroupWriterT = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_semigroupWriterT(dictApply_0_box, dictSemigroup_1_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_semigroupWriterT
}

var cache_Control_Monad_Writer_Trans_applicativeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_applicativeWriterT sync.Once

func Get_Control_Monad_Writer_Trans_applicativeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_applicativeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_applicativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_applicativeWriterT
}

var cache_Control_Monad_Writer_Trans_monadWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadWriterT
}

var cache_Control_Monad_Writer_Trans_monadAskWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadAskWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadAskWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadAskWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadAskWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadAskWriterT
}

var cache_Control_Monad_Writer_Trans_monadReaderWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadReaderWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadReaderWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadReaderWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadReaderWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadReaderWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadReaderWriterT
}

var cache_Control_Monad_Writer_Trans_monadContWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadContWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadContWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadContWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadContWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadContWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadContWriterT
}

var cache_Control_Monad_Writer_Trans_monadEffectWriter gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadEffectWriter sync.Once

func Get_Control_Monad_Writer_Trans_monadEffectWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadEffectWriter.Do(func() {
		cache_Control_Monad_Writer_Trans_monadEffectWriter = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadEffectWriter
}

var cache_Control_Monad_Writer_Trans_monadRecWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadRecWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadRecWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadRecWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadRecWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadRecWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadRecWriterT
}

var cache_Control_Monad_Writer_Trans_monadStateWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadStateWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadStateWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadStateWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadStateWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadStateWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadStateWriterT
}

var cache_Control_Monad_Writer_Trans_monadTellWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadTellWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadTellWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadTellWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadTellWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadTellWriterT
}

var cache_Control_Monad_Writer_Trans_monadWriterWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadWriterWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadWriterWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadWriterWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadWriterWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadWriterWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadWriterWriterT
}

var cache_Control_Monad_Writer_Trans_monadThrowWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadThrowWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadThrowWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadThrowWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadThrowWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadThrowWriterT
}

var cache_Control_Monad_Writer_Trans_monadErrorWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadErrorWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadErrorWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadErrorWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadErrorWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadErrorWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadErrorWriterT
}

var cache_Control_Monad_Writer_Trans_monadSTWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadSTWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadSTWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadSTWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadSTWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadSTWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadSTWriterT
}

var cache_Control_Monad_Writer_Trans_monoidWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monoidWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monoidWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monoidWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monoidWriterT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monoidWriterT(dictApplicative_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monoidWriterT
}

var cache_Control_Monad_Writer_Trans_altWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_altWriterT sync.Once

func Get_Control_Monad_Writer_Trans_altWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_altWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_altWriterT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_altWriterT(dictAlt_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_altWriterT
}

var cache_Control_Monad_Writer_Trans_plusWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_plusWriterT sync.Once

func Get_Control_Monad_Writer_Trans_plusWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_plusWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_plusWriterT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_plusWriterT(dictPlus_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_plusWriterT
}

var cache_Control_Monad_Writer_Trans_alternativeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_alternativeWriterT sync.Once

func Get_Control_Monad_Writer_Trans_alternativeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_alternativeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_alternativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_alternativeWriterT
}

var cache_Control_Monad_Writer_Trans_monadPlusWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadPlusWriterT sync.Once

func Get_Control_Monad_Writer_Trans_monadPlusWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadPlusWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadPlusWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_Trans_monadPlusWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Monad_Writer_Trans_monadPlusWriterT
}

func Call_Control_Monad_Writer_Trans_WriterT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_Writer_Trans_runWriterT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar w$scope78)
	mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): pure_3_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope83), (TypeVar w$scope78)])] (TypeApp (TypeVar m$scope82) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope83), (TypeVar w$scope78)])]))
		pure_3_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{})))
		_ = pure_3_1
		return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "bind"), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{a_5, mempty_1_0}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))})
			}))
		})
	})}))}
}

func Call_Control_Monad_Writer_Trans_mapWriterT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Writer_Trans_functorWriterT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_mapWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})))
	})}))}
}

func Call_Control_Monad_Writer_Trans_execWriterT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(dictFunctor_0.V0, Get_Data_Tuple_snd(), v_1)
}

func Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
	_ = dictApply_1
	// TAST (Let): functorWriterT1_2_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m$scope254) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope253)])])])
	functorWriterT1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_functorWriterT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{})))
	_ = functorWriterT1_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_2_0)}
	}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(v3_5 gopurs_runtime.Value, v4_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_5.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_5.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_6.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), v_3), v1_4)
	})}))}
}

func Call_Control_Monad_Writer_Trans_bindWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictBind_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var dictBind_1 gopurs_runtime.Value = dictBind_1_loop
	_ = dictBind_1
	// TAST (Let): Apply0_2_0 shape=App(Other) bindingType=Any
	Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_1, "Apply0"), gopurs_runtime.Value{})
	_ = Apply0_2_0
	// TAST (Let): applyWriterT2_3_1 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m$scope240) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope239)])])])
	applyWriterT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0, Apply0_2_0))
	_ = applyWriterT2_3_1
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_1)}
	}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_1, "bind"), v_4, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_7_2 shape=Other bindingType=Any
			__local_var_7_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1
			_ = __local_var_7_2
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v3_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_7_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1)}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			}), gopurs_runtime.Apply(k_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0))
		}))
	})}))}
}

func Call_Control_Monad_Writer_Trans_semigroupWriterT(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
	_ = dictSemigroup_1
	// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
	__local_var_2_0 := Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_1, dictApply_0)
	_ = __local_var_2_0
	return gopurs_runtime.Func(func(dictSemigroup1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope16), (TypeVar a$scope16)] (TypeVar a$scope16))
		__local_var_4_1 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup1_3))
		_ = __local_var_4_1
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_4_1, a_5), b_6)
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar w$scope267)
	mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_1_0
	// TAST (Let): applyWriterT1_2_1 shape=App(Var) bindingType=Any
	applyWriterT1_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_applyWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
	_ = applyWriterT1_2_1
	return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): pure_4_2 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope271), (TypeVar w$scope267)])] (TypeApp (TypeVar m$scope268) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope271), (TypeVar w$scope267)])]))
		pure_4_2 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))
		_ = pure_4_2
		// TAST (Let): applyWriterT2_5_3 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m$scope268) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope267)])])])
		applyWriterT2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
		_ = applyWriterT2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_5_3)}
		}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_4_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{a_6, mempty_1_0}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): applicativeWriterT1_1_0 shape=App(Var) bindingType=Any
	applicativeWriterT1_1_0 := Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0)
	_ = applicativeWriterT1_1_0
	// TAST (Let): bindWriterT1_2_1 shape=App(Var) bindingType=Any
	bindWriterT1_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_bindWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
	_ = bindWriterT1_2_1
	return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applicativeWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m$scope75) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope74)])])])
		applicativeWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{})))
		_ = applicativeWriterT2_4_2
		// TAST (Let): bindWriterT2_5_3 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m$scope75) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope74)])])])
		bindWriterT2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})))
		_ = bindWriterT2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_4_2)}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_5_3)}
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadWriterT1_1_0 shape=App(Var) bindingType=Any
	monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_1_0
	return gopurs_runtime.Func(func(dictMonadAsk_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadWriterT2_3_1 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope192) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope190)])])])
		monadWriterT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_2, "Monad0"), gopurs_runtime.Value{})))
		_ = monadWriterT2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_3_1)}
		}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_2, "Monad0"), gopurs_runtime.Value{})))}, Call_Control_Monad_Reader_Class_ask(dictMonadAsk_2))}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadReaderWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadAskWriterT1_1_0 shape=App(Var) bindingType=Any
	monadAskWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0)
	_ = monadAskWriterT1_1_0
	return gopurs_runtime.Func(func(dictMonadReader_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadAskWriterT2_3_1 shape=App(Other) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r$scope214), (TypeApp (TypeVar m$scope215) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope213)])])])
		monadAskWriterT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadAskWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "MonadAsk0"), gopurs_runtime.Value{})))
		_ = monadAskWriterT2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskWriterT2_3_1)}
		}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope215) [(TypeVar a)])] (TypeApp (TypeVar m$scope215) [(TypeVar a)]))
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "local"), f_4)
			_ = __local_var_5_2
			return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_5_2, v_6)
			})
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadContWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar w$scope172)
	mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_1_0
	// TAST (Let): monadWriterT1_2_1 shape=App(Var) bindingType=Any
	monadWriterT1_2_1 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_2_1
	return gopurs_runtime.Func(func(dictMonadCont_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope173) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope172)])])])
		monadWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_3, "Monad0"), gopurs_runtime.Value{})))
		_ = monadWriterT2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_2)}
		}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_3, "callCC"), gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f_5, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(c_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{a_7, mempty_1_0}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))})
				}))
			}))
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): lift_1_0 shape=App(Var) bindingType=Any
	lift_1_0 := Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0)))
	_ = lift_1_0
	// TAST (Let): monadWriterT1_2_1 shape=App(Var) bindingType=Any
	monadWriterT1_2_1 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_2_1
	return gopurs_runtime.Func(func(dictMonadEffect_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Monad0_4_2 shape=App(Other) bindingType=Any
		Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_3, "Monad0"), gopurs_runtime.Value{})
		_ = Monad0_4_2
		// TAST (Let): monadWriterT2_5_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope165) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope164)])])])
		monadWriterT2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2))
		_ = monadWriterT2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_5_3)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(lift_1_0, Monad0_4_2), Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dictMonadEffect_3)))}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadRecWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar w$scope124)
	mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_2_1
	// TAST (Let): monadWriterT1_3_2 shape=App(Var) bindingType=Any
	monadWriterT1_3_2 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_3_2
	return gopurs_runtime.Func(func(dictMonadRec_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Monad0_5_3 shape=App(Other) bindingType=Any
		Monad0_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_4, "Monad0"), gopurs_runtime.Value{})
		_ = Monad0_5_3
		// TAST (Let): monadWriterT2_6_4 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope125) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope124)])])])
		monadWriterT2_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_3_2, Monad0_5_3))
		_ = monadWriterT2_6_4
		return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_6_4)}
		}), gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_4, "tailRecM"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_10_5 shape=Other bindingType=Any
				__local_var_10_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
				_ = __local_var_10_5
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t8 gopurs_runtime.Value
					{
						var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0
						_ = __t_tag_6
						if __t_tag_6.Type == 9 && __t_tag_6.IntVal == 525585346 {
							__t8 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Writer_Trans_3131450224_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), __local_var_10_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V1)})})))}
							goto end_branch_8
						} else {

						}
					}
					{
						var __t_tag_7 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0
						_ = __t_tag_7
						if __t_tag_7.Type == 9 && __t_tag_7.IntVal == 60402430 {
							__t8 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Writer_Trans_2886445004_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), __local_var_10_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V1)})})))}
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_8:
					return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t8)
				}))
			}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{a_8, mempty_2_1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadStateWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadWriterT1_1_0 shape=App(Var) bindingType=Any
	monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_1_0
	return gopurs_runtime.Func(func(dictMonadState_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): lift1_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope117) [(TypeVar a$scope121)])] (TypeApp (TypeVar m$scope117) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope121), (TypeVar w$scope115)])]))
		lift1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_2, "Monad0"), gopurs_runtime.Value{})))})
		_ = lift1_3_1
		// TAST (Let): monadWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope117) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope115)])])])
		monadWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_2, "Monad0"), gopurs_runtime.Value{})))
		_ = monadWriterT2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_2)}
		}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(lift1_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_2, "state"), f_5))
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w$scope101)])
	Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
	_ = Semigroup0_1_0
	// TAST (Let): monadWriterT1_2_1 shape=App(Var) bindingType=Any
	monadWriterT1_2_1 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_2_1
	return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope102) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope101)])])])
		monadWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_2_1, dictMonad_3))
		_ = monadWriterT2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_2)}
		}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_1_0)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Control_Monad_Writer_Trans_WriterT(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), Get_Data_Unit_unit())))}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadWriterWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadTellWriterT1_1_0 shape=App(Var) bindingType=Any
	monadTellWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0)
	_ = monadTellWriterT1_1_0
	return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Bind1_3_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope53)])
		Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
		_ = Bind1_3_1
		// TAST (Let): Applicative0_4_2 shape=App(Other) bindingType=Any
		Applicative0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{})
		_ = Applicative0_4_2
		// TAST (Let): pure_5_3 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope58), (TypeVar w$scope52)]), (TypeVar w$scope52)])] (TypeApp (TypeVar m$scope53) [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope58), (TypeVar w$scope52)]), (TypeVar w$scope52)])]))
		pure_5_3 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_4_2))
		_ = pure_5_3
		// TAST (Let): pure1_6_4 shape=App(Var) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope66), (TypeVar w$scope52)])] (TypeApp (TypeVar m$scope53) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope66), (TypeVar w$scope52)])]))
		pure1_6_4 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Applicative0_4_2))
		_ = pure1_6_4
		// TAST (Let): monadTellWriterT2_7_5 shape=App(Other) bindingType=(ADT ["Control","Monad","Writer","Class","MonadTell"] [(TypeVar w$scope52), (TypeApp (TypeVar m$scope53) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope52)])])])
		monadTellWriterT2_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadTellWriterT1_1_0, dictMonad_2))
		_ = monadTellWriterT2_7_5
		return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellWriterT2_7_5)}
		}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return dictMonoid_0
		}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Bind1_3_1.V1, v_8, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(pure_5_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Writer_Trans_1785332133_138441832(Rebox_Control_Monad_Writer_Trans_138441832_1785332133(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1}))}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))))})
			}))
		}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Bind1_3_1.V1, v_8, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(pure1_6_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1)}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))})
			}))
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadWriterT1_1_0 shape=App(Var) bindingType=Any
	monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_1_0
	return gopurs_runtime.Func(func(dictMonadThrow_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): lift1_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope94) [(TypeVar a$scope98)])] (TypeApp (TypeVar m$scope94) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope98), (TypeVar w$scope92)])]))
		lift1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_2, "Monad0"), gopurs_runtime.Value{})))})
		_ = lift1_3_1
		// TAST (Let): monadWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope94) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope92)])])])
		monadWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_2, "Monad0"), gopurs_runtime.Value{})))
		_ = monadWriterT2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_2)}
		}), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(lift1_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_2, "throwError"), e_5))
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadErrorWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadThrowWriterT1_1_0 shape=App(Var) bindingType=Any
	monadThrowWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0)
	_ = monadThrowWriterT1_1_0
	return gopurs_runtime.Func(func(dictMonadError_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadThrowWriterT2_3_1 shape=App(Other) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e$scope153), (TypeApp (TypeVar m$scope154) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope152)])])])
		monadThrowWriterT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadThrowWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_2, "MonadThrow0"), gopurs_runtime.Value{})))
		_ = monadThrowWriterT2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowWriterT2_3_1)}
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, h_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_2, "catchError"), v_4, gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(h_5, e_6)
			}))
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadSTWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): lift_1_0 shape=App(Var) bindingType=Any
	lift_1_0 := Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0)))
	_ = lift_1_0
	// TAST (Let): monadWriterT1_2_1 shape=App(Var) bindingType=Any
	monadWriterT1_2_1 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_2_1
	return gopurs_runtime.Func(func(dictMonadST_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Monad0_4_2 shape=App(Other) bindingType=Any
		Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_3, "Monad0"), gopurs_runtime.Value{})
		_ = Monad0_4_2
		// TAST (Let): monadWriterT2_5_3 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope5) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope3)])])])
		monadWriterT2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2))
		_ = monadWriterT2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_5_3)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(lift_1_0, Monad0_4_2), Call_Control_Monad_ST_Class_liftST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadST_3)))}))}
	})
}

func Call_Control_Monad_Writer_Trans_monoidWriterT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): semigroupWriterT1_1_0 shape=App(Var) bindingType=Any
	semigroupWriterT1_1_0 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_semigroupWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
	_ = semigroupWriterT1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupWriterT2_3_1 shape=App(Other) bindingType=Any
		semigroupWriterT2_3_1 := gopurs_runtime.Apply(semigroupWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
		_ = semigroupWriterT2_3_1
		return gopurs_runtime.Func(func(dictMonoid1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): semigroupWriterT3_5_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar m$scope46) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope48), (TypeVar w$scope47)])])])
			semigroupWriterT3_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupWriterT2_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_4, "Semigroup0"), gopurs_runtime.Value{})))
			_ = semigroupWriterT3_5_2
			return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupWriterT3_5_2)}
			}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_2), dictApplicative_0), "pure"), gopurs_runtime.RecordGet(dictMonoid1_4, "mempty"))}))}
		})
	})
}

func Call_Control_Monad_Writer_Trans_altWriterT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
	_ = dictAlt_0
	// TAST (Let): functorWriterT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m$scope284) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope285)])])])
	functorWriterT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_functorWriterT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})))
	_ = functorWriterT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_2, v1_3)
	})}))}
}

func Call_Control_Monad_Writer_Trans_plusWriterT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
	_ = dictPlus_0
	// TAST (Let): altWriterT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar m$scope26) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope27)])])])
	altWriterT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_altWriterT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})))
	_ = altWriterT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altWriterT1_1_0)}
	}), Call_Control_Plus_empty(dictPlus_0)}))}
}

func Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): applicativeWriterT1_1_0 shape=App(Var) bindingType=Any
	applicativeWriterT1_1_0 := Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0)
	_ = applicativeWriterT1_1_0
	return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applicativeWriterT2_3_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m$scope280) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope279)])])])
		applicativeWriterT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{})))
		_ = applicativeWriterT2_3_1
		// TAST (Let): plusWriterT1_4_2 shape=App(Var) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar m$scope280) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope279)])])])
		plusWriterT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_plusWriterT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})))
		_ = plusWriterT1_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_3_1)}
		}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusWriterT1_4_2)}
		})}))}
	})
}

func Call_Control_Monad_Writer_Trans_monadPlusWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): monadWriterT1_1_0 shape=App(Var) bindingType=Any
	monadWriterT1_1_0 := Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0)
	_ = monadWriterT1_1_0
	// TAST (Let): alternativeWriterT1_2_1 shape=App(Var) bindingType=Any
	alternativeWriterT1_2_1 := Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0)
	_ = alternativeWriterT1_2_1
	return gopurs_runtime.Func(func(dictMonadPlus_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m$scope147) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope146)])])])
		monadWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Monad0"), gopurs_runtime.Value{})))
		_ = monadWriterT2_4_2
		// TAST (Let): alternativeWriterT2_5_3 shape=App(Other) bindingType=(ADT ["Control","Alternative","Alternative"] [(TypeApp (TypeVar m$scope147) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope146)])])])
		alternativeWriterT2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](gopurs_runtime.Apply(alternativeWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Alternative1"), gopurs_runtime.Value{})))
		_ = alternativeWriterT2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeWriterT2_5_3)}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_2)}
		})}))}
	})
}

func Rebox_Control_Monad_Writer_Trans_138441832_1785332133(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Writer_Trans_1785332133_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Writer_Trans_2886445004_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_Writer_Trans_3131450224_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}
