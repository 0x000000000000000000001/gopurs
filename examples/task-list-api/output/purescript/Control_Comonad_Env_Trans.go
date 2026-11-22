package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Env_Trans_go__map gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_go__map sync.Once

func Get_Control_Comonad_Env_Trans_go__map() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_go__map.Do(func() {
		cache_Control_Comonad_Env_Trans_go__map = gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()).V0
	})
	return cache_Control_Comonad_Env_Trans_go__map
}

var cache_Control_Comonad_Env_Trans_map1 gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_map1 sync.Once

func Get_Control_Comonad_Env_Trans_map1() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_map1.Do(func() {
		cache_Control_Comonad_Env_Trans_map1 = gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()).V0
	})
	return cache_Control_Comonad_Env_Trans_map1
}

var cache_Control_Comonad_Env_Trans_map2 gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_map2 sync.Once

func Get_Control_Comonad_Env_Trans_map2() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_map2.Do(func() {
		cache_Control_Comonad_Env_Trans_map2 = gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()).V0
	})
	return cache_Control_Comonad_Env_Trans_map2
}

var cache_Control_Comonad_Env_Trans_composeFlipped gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_composeFlipped sync.Once

func Get_Control_Comonad_Env_Trans_composeFlipped() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_composeFlipped.Do(func() {
		cache_Control_Comonad_Env_Trans_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Env_Trans_composeFlipped
}

var cache_Control_Comonad_Env_Trans_EnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_EnvT sync.Once

func Get_Control_Comonad_Env_Trans_EnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_EnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_EnvT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Env_Trans_EnvT(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Env_Trans_EnvT
}

var cache_Control_Comonad_Env_Trans_withEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_withEnvT sync.Once

func Get_Control_Comonad_Env_Trans_withEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_withEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_withEnvT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Env_Trans_withEnvT(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Env_Trans_withEnvT
}

var cache_Control_Comonad_Env_Trans_runEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_runEnvT sync.Once

func Get_Control_Comonad_Env_Trans_runEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_runEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_runEnvT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Env_Trans_runEnvT(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Env_Trans_runEnvT
}

var cache_Control_Comonad_Env_Trans_newtypeEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_newtypeEnvT sync.Once

func Get_Control_Comonad_Env_Trans_newtypeEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_newtypeEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_newtypeEnvT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_1417950280_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Control_Comonad_Env_Trans_newtypeEnvT
}

var cache_Control_Comonad_Env_Trans_mapEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_mapEnvT sync.Once

func Get_Control_Comonad_Env_Trans_mapEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_mapEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_mapEnvT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Env_Trans_mapEnvT(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Env_Trans_mapEnvT
}

var cache_Control_Comonad_Env_Trans_functorEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_functorEnvT sync.Once

func Get_Control_Comonad_Env_Trans_functorEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_functorEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_functorEnvT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_functorEnvT(dictFunctor_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_functorEnvT
}

var cache_Control_Comonad_Env_Trans_functorWithIndexEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_functorWithIndexEnvT sync.Once

func Get_Control_Comonad_Env_Trans_functorWithIndexEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_functorWithIndexEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_functorWithIndexEnvT = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_functorWithIndexEnvT(dictFunctorWithIndex_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_functorWithIndexEnvT
}

var cache_Control_Comonad_Env_Trans_foldableEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_foldableEnvT sync.Once

func Get_Control_Comonad_Env_Trans_foldableEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_foldableEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_foldableEnvT = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_foldableEnvT(dictFoldable_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_foldableEnvT
}

var cache_Control_Comonad_Env_Trans_foldableWithIndexEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_foldableWithIndexEnvT sync.Once

func Get_Control_Comonad_Env_Trans_foldableWithIndexEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_foldableWithIndexEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_foldableWithIndexEnvT = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_foldableWithIndexEnvT(dictFoldableWithIndex_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_foldableWithIndexEnvT
}

var cache_Control_Comonad_Env_Trans_traversableEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_traversableEnvT sync.Once

func Get_Control_Comonad_Env_Trans_traversableEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_traversableEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_traversableEnvT = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_traversableEnvT(dictTraversable_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_traversableEnvT
}

var cache_Control_Comonad_Env_Trans_traversableWithIndexEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_traversableWithIndexEnvT sync.Once

func Get_Control_Comonad_Env_Trans_traversableWithIndexEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_traversableWithIndexEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_traversableWithIndexEnvT = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_traversableWithIndexEnvT(dictTraversableWithIndex_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_traversableWithIndexEnvT
}

var cache_Control_Comonad_Env_Trans_extendEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_extendEnvT sync.Once

func Get_Control_Comonad_Env_Trans_extendEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_extendEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_extendEnvT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_extendEnvT(dictExtend_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_extendEnvT
}

var cache_Control_Comonad_Env_Trans_comonadTransEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_comonadTransEnvT sync.Once

func Get_Control_Comonad_Env_Trans_comonadTransEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_comonadTransEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_comonadTransEnvT = gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_34561274_3351420535((&Constructor_Control_Comonad_Trans_Class_ComonadTrans[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(dictComonad_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
		})})))}
	})
	return cache_Control_Comonad_Env_Trans_comonadTransEnvT
}

var cache_Control_Comonad_Env_Trans_comonadEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_comonadEnvT sync.Once

func Get_Control_Comonad_Env_Trans_comonadEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_comonadEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_comonadEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Trans_comonadEnvT(dictComonad_0_box)
		})
	})
	return cache_Control_Comonad_Env_Trans_comonadEnvT
}

func Call_Control_Comonad_Env_Trans_EnvT(x_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var x_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(x_0)}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Control_Comonad_Env_Trans_withEnvT(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
	_ = v_1
	return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{gopurs_runtime.Value{}, gopurs_runtime.Value{}}
}

func Call_Control_Comonad_Env_Trans_runEnvT(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var v_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Control_Comonad_Env_Trans_mapEnvT(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
	_ = v_1
	return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{gopurs_runtime.Value{}, gopurs_runtime.Value{}}
}

func Call_Control_Comonad_Env_Trans_functorEnvT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_2363162019_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)}))}
	})})))}
}

func Call_Control_Comonad_Env_Trans_functorWithIndexEnvT(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
	_ = dictFunctorWithIndex_0
	// TAST (Let): functorEnvT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope97), (TypeVar w$scope96)])])
	functorEnvT1_1_0 := Rebox_Control_Comonad_Env_Trans_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_functorEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))))
	_ = functorEnvT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_935379557_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_2363162019_2812149806(functorEnvT1_1_0))}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)}))}
	})})))}
}

func Call_Control_Comonad_Env_Trans_foldableEnvT(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_4173511203_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldMap1_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope160)] (TypeVar m$scope161)), (TypeApp (TypeVar f$scope147) [(TypeVar a$scope160)])] (TypeVar m$scope161))
		foldMap1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))})
		_ = foldMap1_2_0
		return gopurs_runtime.Func2(func(fn_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(foldMap1_2_0, fn_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
		})
	}), gopurs_runtime.Func3(func(fn_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), fn_1, a_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
	}), gopurs_runtime.Func3(func(fn_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), fn_1, a_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
	})})))}
}

func Call_Control_Comonad_Env_Trans_foldableWithIndexEnvT(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldableEnvT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope130), (TypeVar w$scope129)])])
	foldableEnvT1_1_0 := Rebox_Control_Comonad_Env_Trans_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_foldableEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))))
	_ = foldableEnvT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_3374046885_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_4173511203_1680800814(foldableEnvT1_1_0))}
	}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldMapWithIndex1_3_1 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope128), (TypeVar a$scope143)] (TypeVar m$scope144)), (TypeApp (TypeVar w$scope129) [(TypeVar a$scope143)])] (TypeVar m$scope144))
		foldMapWithIndex1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))})
		_ = foldMapWithIndex1_3_1
		return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(foldMapWithIndex1_3_1, f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
		})
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_2, a_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_2, a_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
	})})))}
}

func Call_Control_Comonad_Env_Trans_traversableEnvT(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): functorEnvT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope37), (TypeVar f$scope36)])])
	functorEnvT1_1_0 := Rebox_Control_Comonad_Env_Trans_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_functorEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))))
	_ = functorEnvT1_1_0
	// TAST (Let): foldableEnvT1_2_1 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope37), (TypeVar f$scope36)])])
	foldableEnvT1_2_1 := Rebox_Control_Comonad_Env_Trans_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_foldableEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))))
	_ = foldableEnvT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_3543431075_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_4173511203_1680800814(foldableEnvT1_2_1))}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_2363162019_2812149806(functorEnvT1_1_0))}
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): sequence1_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope36) [(TypeApp (TypeVar m$scope45) [(TypeVar a$scope44)])])] (TypeApp (TypeVar m$scope45) [(TypeApp (TypeVar f$scope36) [(TypeVar a$scope44)])]))
		sequence1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = sequence1_4_2
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_3 shape=App(Var) bindingType=Any
			__local_var_6_3 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
			_ = __local_var_6_3
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_6_3, x_7)
			}), gopurs_runtime.Apply(sequence1_4_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
		})
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverse1_4_4 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope52)] (TypeApp (TypeVar m$scope54) [(TypeVar b$scope53)])), (TypeApp (TypeVar f$scope36) [(TypeVar a$scope52)])] (TypeApp (TypeVar m$scope54) [(TypeApp (TypeVar f$scope36) [(TypeVar b$scope53)])]))
		traverse1_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = traverse1_4_4
		return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_7_5 shape=App(Var) bindingType=Any
			__local_var_7_5 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
			_ = __local_var_7_5
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_7_5, x_8)
			}), gopurs_runtime.Apply2(traverse1_4_4, f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
		})
	})})))}
}

func Call_Control_Comonad_Env_Trans_traversableWithIndexEnvT(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
	_ = dictTraversableWithIndex_0
	// TAST (Let): functorWithIndexEnvT1_1_0 shape=App(Var) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(TypeVar i$scope17), (TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope19), (TypeVar w$scope18)])])
	functorWithIndexEnvT1_1_0 := Rebox_Control_Comonad_Env_Trans_2412140840_935379557(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_functorWithIndexEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))))
	_ = functorWithIndexEnvT1_1_0
	// TAST (Let): foldableWithIndexEnvT1_2_1 shape=App(Var) bindingType=(ADT ["Data","FoldableWithIndex","FoldableWithIndex"] [(TypeVar i$scope17), (TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope19), (TypeVar w$scope18)])])
	foldableWithIndexEnvT1_2_1 := Rebox_Control_Comonad_Env_Trans_3725484264_3374046885(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_foldableWithIndexEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))))
	_ = foldableWithIndexEnvT1_2_1
	// TAST (Let): traversableEnvT1_3_2 shape=App(Var) bindingType=(ADT ["Data","Traversable","Traversable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope19), (TypeVar w$scope18)])])
	traversableEnvT1_3_2 := Rebox_Control_Comonad_Env_Trans_3043886126_3543431075(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_traversableEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))))
	_ = traversableEnvT1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_2190796645_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_3374046885_3725484264(foldableWithIndexEnvT1_2_1))}
	}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_935379557_2412140840(functorWithIndexEnvT1_1_0))}
	}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_3543431075_3043886126(traversableEnvT1_3_2))}
	}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverseWithIndex1_5_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope17), (TypeVar a$scope25)] (TypeApp (TypeVar m$scope27) [(TypeVar b$scope26)])), (TypeApp (TypeVar w$scope18) [(TypeVar a$scope25)])] (TypeApp (TypeVar m$scope27) [(TypeApp (TypeVar w$scope18) [(TypeVar b$scope26)])]))
		traverseWithIndex1_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_4))})
		_ = traverseWithIndex1_5_3
		return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_8_4 shape=App(Var) bindingType=Any
			__local_var_8_4 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)
			_ = __local_var_8_4
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_8_4, x_9)
			}), gopurs_runtime.Apply2(traverseWithIndex1_5_3, f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
		})
	})})))}
}

func Call_Control_Comonad_Env_Trans_extendEnvT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
	_ = dictExtend_0
	// TAST (Let): functorEnvT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope165), (TypeVar w$scope164)])])
	functorEnvT1_1_0 := Rebox_Control_Comonad_Env_Trans_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_functorEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))))
	_ = functorEnvT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_3108778612_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_2363162019_2812149806(functorEnvT1_1_0))}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}), "map"), f_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), Get_Control_Comonad_Env_Trans_EnvT()), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))}
	})})))}
}

func Call_Control_Comonad_Env_Trans_comonadEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): extendEnvT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Extend","Extend"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope194), (TypeVar w$scope193)])])
	extendEnvT1_1_0 := Rebox_Control_Comonad_Env_Trans_3290176857_3108778612(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_extendEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))))
	_ = extendEnvT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_3056445460_2550391993((&Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Trans_3108778612_3290176857(extendEnvT1_1_0))}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
	})})))}
}

func Rebox_Control_Comonad_Env_Trans_1417950280_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_1680800814_4173511203(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_2190796645_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_2412140840_935379557(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_3043886126_3543431075(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_3108778612_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_3290176857_3108778612(in *Constructor_Control_Extend_Extend[gopurs_runtime.Value]) *Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_3374046885_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_34561274_3351420535(in *Constructor_Control_Comonad_Trans_Class_ComonadTrans[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_3543431075_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_3725484264_3374046885(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_4173511203_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Trans_935379557_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
