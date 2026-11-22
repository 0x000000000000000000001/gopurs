package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Product_go__append gopurs_runtime.Value
var once_Data_Functor_Product_go__append sync.Once

func Get_Data_Functor_Product_go__append() gopurs_runtime.Value {
	once_Data_Functor_Product_go__append.Do(func() {
		cache_Data_Functor_Product_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Functor_Product_go__append
}

var cache_Data_Functor_Product_bimap gopurs_runtime.Value
var once_Data_Functor_Product_bimap sync.Once

func Get_Data_Functor_Product_bimap() gopurs_runtime.Value {
	once_Data_Functor_Product_bimap.Do(func() {
		cache_Data_Functor_Product_bimap = Rebox_Data_Functor_Product_1688994542_1495429347(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorTuple())).V0
	})
	return cache_Data_Functor_Product_bimap
}

var cache_Data_Functor_Product_conj gopurs_runtime.Value
var once_Data_Functor_Product_conj sync.Once

func Get_Data_Functor_Product_conj() gopurs_runtime.Value {
	once_Data_Functor_Product_conj.Do(func() {
		cache_Data_Functor_Product_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Functor_Product_conj
}

var cache_Data_Functor_Product_bimap1 gopurs_runtime.Value
var once_Data_Functor_Product_bimap1 sync.Once

func Get_Data_Functor_Product_bimap1() gopurs_runtime.Value {
	once_Data_Functor_Product_bimap1.Do(func() {
		cache_Data_Functor_Product_bimap1 = Rebox_Data_Functor_Product_1688994542_1495429347(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorTuple())).V0
	})
	return cache_Data_Functor_Product_bimap1
}

var cache_Data_Functor_Product_compose gopurs_runtime.Value
var once_Data_Functor_Product_compose sync.Once

func Get_Data_Functor_Product_compose() gopurs_runtime.Value {
	once_Data_Functor_Product_compose.Do(func() {
		cache_Data_Functor_Product_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Product_compose
}

var cache_Data_Functor_Product_compose1 gopurs_runtime.Value
var once_Data_Functor_Product_compose1 sync.Once

func Get_Data_Functor_Product_compose1() gopurs_runtime.Value {
	once_Data_Functor_Product_compose1.Do(func() {
		cache_Data_Functor_Product_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Product_compose1
}

var cache_Data_Functor_Product_unwrap gopurs_runtime.Value
var once_Data_Functor_Product_unwrap sync.Once

func Get_Data_Functor_Product_unwrap() gopurs_runtime.Value {
	once_Data_Functor_Product_unwrap.Do(func() {
		cache_Data_Functor_Product_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Functor_Product_unwrap
}

var cache_Data_Functor_Product_compose2 gopurs_runtime.Value
var once_Data_Functor_Product_compose2 sync.Once

func Get_Data_Functor_Product_compose2() gopurs_runtime.Value {
	once_Data_Functor_Product_compose2.Do(func() {
		cache_Data_Functor_Product_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Product_compose2
}

var cache_Data_Functor_Product_compose3 gopurs_runtime.Value
var once_Data_Functor_Product_compose3 sync.Once

func Get_Data_Functor_Product_compose3() gopurs_runtime.Value {
	once_Data_Functor_Product_compose3.Do(func() {
		cache_Data_Functor_Product_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Product_compose3
}

var cache_Data_Functor_Product_unwrap1 gopurs_runtime.Value
var once_Data_Functor_Product_unwrap1 sync.Once

func Get_Data_Functor_Product_unwrap1() gopurs_runtime.Value {
	once_Data_Functor_Product_unwrap1.Do(func() {
		cache_Data_Functor_Product_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Functor_Product_unwrap1
}

var cache_Data_Functor_Product_Product gopurs_runtime.Value
var once_Data_Functor_Product_Product sync.Once

func Get_Data_Functor_Product_Product() gopurs_runtime.Value {
	once_Data_Functor_Product_Product.Do(func() {
		cache_Data_Functor_Product_Product = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_Product_Product(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Functor_Product_Product
}

var cache_Data_Functor_Product_showProduct gopurs_runtime.Value
var once_Data_Functor_Product_showProduct sync.Once

func Get_Data_Functor_Product_showProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_showProduct.Do(func() {
		cache_Data_Functor_Product_showProduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_showProduct(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Data_Functor_Product_showProduct
}

var cache_Data_Functor_Product_product gopurs_runtime.Value
var once_Data_Functor_Product_product sync.Once

func Get_Data_Functor_Product_product() gopurs_runtime.Value {
	once_Data_Functor_Product_product.Do(func() {
		cache_Data_Functor_Product_product = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_Product_product(fa_0_box, ga_1_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Functor_Product_product
}

var cache_Data_Functor_Product_newtypeProduct gopurs_runtime.Value
var once_Data_Functor_Product_newtypeProduct sync.Once

func Get_Data_Functor_Product_newtypeProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_newtypeProduct.Do(func() {
		cache_Data_Functor_Product_newtypeProduct = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_1417950280_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Functor_Product_newtypeProduct
}

var cache_Data_Functor_Product_functorProduct gopurs_runtime.Value
var once_Data_Functor_Product_functorProduct sync.Once

func Get_Data_Functor_Product_functorProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_functorProduct.Do(func() {
		cache_Data_Functor_Product_functorProduct = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_functorProduct(dictFunctor_0_box, dictFunctor1_1_box)
		})
	})
	return cache_Data_Functor_Product_functorProduct
}

var cache_Data_Functor_Product_eq1Product gopurs_runtime.Value
var once_Data_Functor_Product_eq1Product sync.Once

func Get_Data_Functor_Product_eq1Product() gopurs_runtime.Value {
	once_Data_Functor_Product_eq1Product.Do(func() {
		cache_Data_Functor_Product_eq1Product = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_eq1Product(dictEq1_0_box, dictEq11_1_box)
		})
	})
	return cache_Data_Functor_Product_eq1Product
}

var cache_Data_Functor_Product_eqProduct gopurs_runtime.Value
var once_Data_Functor_Product_eqProduct sync.Once

func Get_Data_Functor_Product_eqProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_eqProduct.Do(func() {
		cache_Data_Functor_Product_eqProduct = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_eqProduct(dictEq1_0_box, dictEq11_1_box)
		})
	})
	return cache_Data_Functor_Product_eqProduct
}

var cache_Data_Functor_Product_ord1Product gopurs_runtime.Value
var once_Data_Functor_Product_ord1Product sync.Once

func Get_Data_Functor_Product_ord1Product() gopurs_runtime.Value {
	once_Data_Functor_Product_ord1Product.Do(func() {
		cache_Data_Functor_Product_ord1Product = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_ord1Product(dictOrd1_0_box)
		})
	})
	return cache_Data_Functor_Product_ord1Product
}

var cache_Data_Functor_Product_ordProduct gopurs_runtime.Value
var once_Data_Functor_Product_ordProduct sync.Once

func Get_Data_Functor_Product_ordProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_ordProduct.Do(func() {
		cache_Data_Functor_Product_ordProduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_ordProduct(dictOrd1_0_box)
		})
	})
	return cache_Data_Functor_Product_ordProduct
}

var cache_Data_Functor_Product_bihoistProduct gopurs_runtime.Value
var once_Data_Functor_Product_bihoistProduct sync.Once

func Get_Data_Functor_Product_bihoistProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_bihoistProduct.Do(func() {
		cache_Data_Functor_Product_bihoistProduct = gopurs_runtime.Func3(func(natF_0_box gopurs_runtime.Value, natG_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_bihoistProduct(natF_0_box, natG_1_box, v_2_box)
		})
	})
	return cache_Data_Functor_Product_bihoistProduct
}

var cache_Data_Functor_Product_applyProduct gopurs_runtime.Value
var once_Data_Functor_Product_applyProduct sync.Once

func Get_Data_Functor_Product_applyProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_applyProduct.Do(func() {
		cache_Data_Functor_Product_applyProduct = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_applyProduct(dictApply_0_box)
		})
	})
	return cache_Data_Functor_Product_applyProduct
}

var cache_Data_Functor_Product_bindProduct gopurs_runtime.Value
var once_Data_Functor_Product_bindProduct sync.Once

func Get_Data_Functor_Product_bindProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_bindProduct.Do(func() {
		cache_Data_Functor_Product_bindProduct = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_bindProduct(dictBind_0_box)
		})
	})
	return cache_Data_Functor_Product_bindProduct
}

var cache_Data_Functor_Product_applicativeProduct gopurs_runtime.Value
var once_Data_Functor_Product_applicativeProduct sync.Once

func Get_Data_Functor_Product_applicativeProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_applicativeProduct.Do(func() {
		cache_Data_Functor_Product_applicativeProduct = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_applicativeProduct(dictApplicative_0_box)
		})
	})
	return cache_Data_Functor_Product_applicativeProduct
}

var cache_Data_Functor_Product_monadProduct gopurs_runtime.Value
var once_Data_Functor_Product_monadProduct sync.Once

func Get_Data_Functor_Product_monadProduct() gopurs_runtime.Value {
	once_Data_Functor_Product_monadProduct.Do(func() {
		cache_Data_Functor_Product_monadProduct = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Product_monadProduct(dictMonad_0_box)
		})
	})
	return cache_Data_Functor_Product_monadProduct
}

func Call_Data_Functor_Product_Product(x_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
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

func Call_Data_Functor_Product_showProduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_1829442467_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(product ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), " "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), ")"))
	})})))}
}

func Call_Data_Functor_Product_product(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var fa_0 gopurs_runtime.Value = fa_0_loop
	_ = fa_0
	var ga_1 gopurs_runtime.Value = ga_1_loop
	_ = ga_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{fa_0, ga_1}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Data_Functor_Product_functorProduct(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
	_ = dictFunctor1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2363162019_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})})))}
}

func Call_Data_Functor_Product_eq1Product(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
	_ = dictEq11_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2909557234_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eq12_3_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope69) [(TypeVar a$scope73)]), (TypeApp (TypeVar f$scope69) [(TypeVar a$scope73)])] Boolean)
		eq12_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))})
		_ = eq12_3_0
		// TAST (Let): eq13_4_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope70) [(TypeVar a$scope73)]), (TypeApp (TypeVar g$scope70) [(TypeVar a$scope73)])] Boolean)
		eq13_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))})
		_ = eq13_4_1
		return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(((gopurs_runtime.Apply2(eq12_3_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(eq13_4_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1).IntVal) != (0)))
		})
	})})))}
}

func Call_Data_Functor_Product_eqProduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
	_ = dictEq11_1
	// TAST (Let): eq1_2_0 shape=App(Var) bindingType=Any
	eq1_2_0 := Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_Functor_Product_eq1Product(dictEq1_0, dictEq11_1)))
	_ = eq1_2_0
	return gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2450374787_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Apply(eq1_2_0, dictEq_3)})))}
	})
}

func Call_Data_Functor_Product_ord1Product(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): eq1Product1_1_0 shape=App(Var) bindingType=Any
	eq1Product1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Product_eq1Product(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
	_ = eq1Product1_1_0
	return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eq1Product2_3_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope28), (TypeVar g$scope29)])])
		eq1Product2_3_1 := Rebox_Data_Functor_Product_1766074591_2909557234(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](gopurs_runtime.Apply(eq1Product1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{}))))
		_ = eq1Product2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_1697277074_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2909557234_1766074591(eq1Product2_3_1))}
		}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): compare12_5_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope28) [(TypeVar a$scope32)]), (TypeApp (TypeVar f$scope28) [(TypeVar a$scope32)])] (ADT ["Data","Ordering","Ordering"] []))
			compare12_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))})
			_ = compare12_5_2
			// TAST (Let): compare13_6_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope29) [(TypeVar a$scope32)]), (TypeApp (TypeVar g$scope29) [(TypeVar a$scope32)])] (ADT ["Data","Ordering","Ordering"] []))
			compare13_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))})
			_ = compare13_6_3
			return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v2_9_4 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
				v2_9_4 := uint32(gopurs_runtime.Apply2(compare12_5_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0).IntVal)
				_ = v2_9_4
				var __t5 gopurs_runtime.Value
				{
					if v2_9_4 == 902936544 {
						__t5 = gopurs_runtime.Apply2(compare13_6_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1)
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(v2_9_4), UnsafePtr: nil}
				}
			end_branch_5:
				return __t5
			})
		})})))}
	})
}

func Call_Data_Functor_Product_ordProduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): ord1Product1_1_0 shape=App(Var) bindingType=Any
	ord1Product1_1_0 := Call_Data_Functor_Product_ord1Product(dictOrd1_0)
	_ = ord1Product1_1_0
	// TAST (Let): eqProduct1_2_1 shape=App(Var) bindingType=Any
	eqProduct1_2_1 := gopurs_runtime.Apply(Get_Data_Functor_Product_eqProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
	_ = eqProduct1_2_1
	return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): compare1_4_2 shape=App(Var) bindingType=Any
		compare1_4_2 := Call_Data_Ord_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](gopurs_runtime.Apply(ord1Product1_1_0, dictOrd11_3)))
		_ = compare1_4_2
		// TAST (Let): eqProduct2_5_3 shape=App(Other) bindingType=Any
		eqProduct2_5_3 := gopurs_runtime.Apply(eqProduct1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{}))
		_ = eqProduct2_5_3
		return gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): eqProduct3_7_4 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope22), (TypeVar g$scope23), (TypeVar a$scope24)])])
			eqProduct3_7_4 := Rebox_Data_Functor_Product_3790796878_2450374787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqProduct2_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_6, "Eq0"), gopurs_runtime.Value{}))))
			_ = eqProduct3_7_4
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_1535415139_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2450374787_3790796878(eqProduct3_7_4))}
			}), gopurs_runtime.Apply(compare1_4_2, dictOrd_6)})))}
		})
	})
}

func Call_Data_Functor_Product_bihoistProduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var natF_0 gopurs_runtime.Value = natF_0_loop
	_ = natF_0
	var natG_1 gopurs_runtime.Value = natG_1_loop
	_ = natG_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{gopurs_runtime.Apply(natF_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(natG_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))}
}

func Call_Data_Functor_Product_applyProduct(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): functorProduct1_1_0 shape=App(Var) bindingType=Any
	functorProduct1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Product_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorProduct1_1_0
	return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): functorProduct2_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope106), (TypeVar g$scope107)])])
		functorProduct2_3_1 := Rebox_Data_Functor_Product_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_2, "Functor0"), gopurs_runtime.Value{}))))
		_ = functorProduct2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2708463828_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2363162019_2812149806(functorProduct2_3_1))}
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Func2(func(fa_6 gopurs_runtime.Value, ga_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{fa_6, ga_7}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply1_2, "apply"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1))
		})})))}
	})
}

func Call_Data_Functor_Product_bindProduct(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): applyProduct1_1_0 shape=App(Var) bindingType=Any
	applyProduct1_1_0 := Call_Data_Functor_Product_applyProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}))
	_ = applyProduct1_1_0
	return gopurs_runtime.Func(func(dictBind1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applyProduct2_3_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope76), (TypeVar g$scope77)])])
		applyProduct2_3_1 := Rebox_Data_Functor_Product_3741347833_2708463828(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind1_2, "Apply0"), gopurs_runtime.Value{}))))
		_ = applyProduct2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2495815764_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2708463828_3741347833(applyProduct2_3_1))}
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Func2(func(fa_6 gopurs_runtime.Value, ga_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{fa_6, ga_7}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_fst(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), f_5))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind1_2, "bind"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), f_5))))
		})})))}
	})
}

func Call_Data_Functor_Product_applicativeProduct(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): applyProduct1_1_0 shape=App(Var) bindingType=Any
	applyProduct1_1_0 := Call_Data_Functor_Product_applyProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
	_ = applyProduct1_1_0
	return gopurs_runtime.Func(func(dictApplicative1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applyProduct2_3_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope116), (TypeVar g$scope117)])])
		applyProduct2_3_1 := Rebox_Data_Functor_Product_3741347833_2708463828(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_2, "Apply0"), gopurs_runtime.Value{}))))
		_ = applyProduct2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_894579924_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2708463828_3741347833(applyProduct2_3_1))}
		}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Func2(func(fa_5 gopurs_runtime.Value, ga_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{fa_5, ga_6}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_2, "pure"), a_4))
		})})))}
	})
}

func Call_Data_Functor_Product_monadProduct(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): applicativeProduct1_1_0 shape=App(Var) bindingType=Any
	applicativeProduct1_1_0 := Call_Data_Functor_Product_applicativeProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = applicativeProduct1_1_0
	// TAST (Let): bindProduct1_2_1 shape=App(Var) bindingType=Any
	bindProduct1_2_1 := Call_Data_Functor_Product_bindProduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
	_ = bindProduct1_2_1
	return gopurs_runtime.Func(func(dictMonad1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applicativeProduct2_4_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope44), (TypeVar g$scope45)])])
		applicativeProduct2_4_2 := Rebox_Data_Functor_Product_1439734649_894579924(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeProduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad1_3, "Applicative0"), gopurs_runtime.Value{}))))
		_ = applicativeProduct2_4_2
		// TAST (Let): bindProduct2_5_3 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope44), (TypeVar g$scope45)])])
		bindProduct2_5_3 := Rebox_Data_Functor_Product_2748095225_2495815764(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindProduct1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad1_3, "Bind1"), gopurs_runtime.Value{}))))
		_ = bindProduct2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2157788756_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_894579924_1439734649(applicativeProduct2_4_2))}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_Product_2495815764_2748095225(bindProduct2_5_3))}
		})})))}
	})
}

func Rebox_Data_Functor_Product_1417950280_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_1439734649_894579924(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_1535415139_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_1688994542_1495429347(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_1697277074_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_1766074591_2909557234(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_1829442467_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2157788756_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Monad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2450374787_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2495815764_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2708463828_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2748095225_2495815764(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_2909557234_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_3741347833_2708463828(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_3790796878_2450374787(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_Product_894579924_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[gopurs_runtime.Value])(unsafe.Pointer(in))
}
