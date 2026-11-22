package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Store_Trans_composeFlipped gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_composeFlipped sync.Once

func Get_Control_Comonad_Store_Trans_composeFlipped() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_composeFlipped.Do(func() {
		cache_Control_Comonad_Store_Trans_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Store_Trans_composeFlipped
}

var cache_Control_Comonad_Store_Trans_StoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_StoreT sync.Once

func Get_Control_Comonad_Store_Trans_StoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_StoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_StoreT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Store_Trans_StoreT(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Store_Trans_StoreT
}

var cache_Control_Comonad_Store_Trans_runStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_runStoreT sync.Once

func Get_Control_Comonad_Store_Trans_runStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_runStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_runStoreT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Store_Trans_runStoreT(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Store_Trans_runStoreT
}

var cache_Control_Comonad_Store_Trans_newtypeStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_newtypeStoreT sync.Once

func Get_Control_Comonad_Store_Trans_newtypeStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_newtypeStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_newtypeStoreT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Trans_1417950280_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Control_Comonad_Store_Trans_newtypeStoreT
}

var cache_Control_Comonad_Store_Trans_functorStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_functorStoreT sync.Once

func Get_Control_Comonad_Store_Trans_functorStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_functorStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_functorStoreT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Trans_functorStoreT(dictFunctor_0_box)
		})
	})
	return cache_Control_Comonad_Store_Trans_functorStoreT
}

var cache_Control_Comonad_Store_Trans_extendStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_extendStoreT sync.Once

func Get_Control_Comonad_Store_Trans_extendStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_extendStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_extendStoreT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Trans_extendStoreT(dictExtend_0_box)
		})
	})
	return cache_Control_Comonad_Store_Trans_extendStoreT
}

var cache_Control_Comonad_Store_Trans_comonadTransStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_comonadTransStoreT sync.Once

func Get_Control_Comonad_Store_Trans_comonadTransStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_comonadTransStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_comonadTransStoreT = gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Trans_34561274_3351420535((&Constructor_Control_Comonad_Trans_Class_ComonadTrans[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(dictComonad_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
			__local_var_2_0 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
			_ = __local_var_2_0
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(v1_3, __local_var_2_0)
			}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)
		})})))}
	})
	return cache_Control_Comonad_Store_Trans_comonadTransStoreT
}

var cache_Control_Comonad_Store_Trans_comonadStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Trans_comonadStoreT sync.Once

func Get_Control_Comonad_Store_Trans_comonadStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Trans_comonadStoreT.Do(func() {
		cache_Control_Comonad_Store_Trans_comonadStoreT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Trans_comonadStoreT(dictComonad_0_box)
		})
	})
	return cache_Control_Comonad_Store_Trans_comonadStoreT
}

func Call_Control_Comonad_Store_Trans_StoreT(x_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
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

func Call_Control_Comonad_Store_Trans_runStoreT(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
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

func Call_Control_Comonad_Store_Trans_functorStoreT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Trans_2363162019_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), h_3, f_1)
		}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1}))}
	})})))}
}

func Call_Control_Comonad_Store_Trans_extendStoreT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
	_ = dictExtend_0
	// TAST (Let): functorStoreT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w$scope32) [(Func [(TypeVar s$scope33)] (TypeVar a))]), (TypeVar s$scope33)])])
	functorStoreT1_1_0 := Rebox_Control_Comonad_Store_Trans_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Comonad_Store_Trans_functorStoreT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))))
	_ = functorStoreT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Trans_3108778612_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Trans_2363162019_2812149806(functorStoreT1_1_0))}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func2(func(w_prime__4 gopurs_runtime.Value, s_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{w_prime__4, s_prime__5}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
		}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1}))}
	})})))}
}

func Call_Control_Comonad_Store_Trans_comonadStoreT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): extendStoreT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Extend","Extend"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w$scope58) [(Func [(TypeVar s$scope59)] (TypeVar a))]), (TypeVar s$scope59)])])
	extendStoreT1_1_0 := Rebox_Control_Comonad_Store_Trans_3290176857_3108778612(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Call_Control_Comonad_Store_Trans_extendStoreT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))))
	_ = extendStoreT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Trans_3056445460_2550391993((&Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Trans_3108778612_3290176857(extendStoreT1_1_0))}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
	})})))}
}

func Rebox_Control_Comonad_Store_Trans_1417950280_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Trans_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Trans_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Trans_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Trans_3108778612_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Trans_3290176857_3108778612(in *Constructor_Control_Extend_Extend[gopurs_runtime.Value]) *Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Trans_34561274_3351420535(in *Constructor_Control_Comonad_Trans_Class_ComonadTrans[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value])(unsafe.Pointer(in))
}
