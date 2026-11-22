package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Flip_go__append gopurs_runtime.Value
var once_Data_Functor_Flip_go__append sync.Once

func Get_Data_Functor_Flip_go__append() gopurs_runtime.Value {
	once_Data_Functor_Flip_go__append.Do(func() {
		cache_Data_Functor_Flip_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Functor_Flip_go__append
}

var cache_Data_Functor_Flip_Flip gopurs_runtime.Value
var once_Data_Functor_Flip_Flip sync.Once

func Get_Data_Functor_Flip_Flip() gopurs_runtime.Value {
	once_Data_Functor_Flip_Flip.Do(func() {
		cache_Data_Functor_Flip_Flip = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_Flip(x_0_box)
		})
	})
	return cache_Data_Functor_Flip_Flip
}

var cache_Data_Functor_Flip_showFlip gopurs_runtime.Value
var once_Data_Functor_Flip_showFlip sync.Once

func Get_Data_Functor_Flip_showFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_showFlip.Do(func() {
		cache_Data_Functor_Flip_showFlip = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_showFlip(dictShow_0_box)
		})
	})
	return cache_Data_Functor_Flip_showFlip
}

var cache_Data_Functor_Flip_semigroupoidFlip gopurs_runtime.Value
var once_Data_Functor_Flip_semigroupoidFlip sync.Once

func Get_Data_Functor_Flip_semigroupoidFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_semigroupoidFlip.Do(func() {
		cache_Data_Functor_Flip_semigroupoidFlip = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_semigroupoidFlip(dictSemigroupoid_0_box)
		})
	})
	return cache_Data_Functor_Flip_semigroupoidFlip
}

var cache_Data_Functor_Flip_ordFlip gopurs_runtime.Value
var once_Data_Functor_Flip_ordFlip sync.Once

func Get_Data_Functor_Flip_ordFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_ordFlip.Do(func() {
		cache_Data_Functor_Flip_ordFlip = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_ordFlip(dictOrd_0_box)
		})
	})
	return cache_Data_Functor_Flip_ordFlip
}

var cache_Data_Functor_Flip_newtypeFlip gopurs_runtime.Value
var once_Data_Functor_Flip_newtypeFlip sync.Once

func Get_Data_Functor_Flip_newtypeFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_newtypeFlip.Do(func() {
		cache_Data_Functor_Flip_newtypeFlip = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Functor_Flip_newtypeFlip
}

var cache_Data_Functor_Flip_functorFlip gopurs_runtime.Value
var once_Data_Functor_Flip_functorFlip sync.Once

func Get_Data_Functor_Flip_functorFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_functorFlip.Do(func() {
		cache_Data_Functor_Flip_functorFlip = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_functorFlip(dictBifunctor_0_box)
		})
	})
	return cache_Data_Functor_Flip_functorFlip
}

var cache_Data_Functor_Flip_eqFlip gopurs_runtime.Value
var once_Data_Functor_Flip_eqFlip sync.Once

func Get_Data_Functor_Flip_eqFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_eqFlip.Do(func() {
		cache_Data_Functor_Flip_eqFlip = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_eqFlip(dictEq_0_box)
		})
	})
	return cache_Data_Functor_Flip_eqFlip
}

var cache_Data_Functor_Flip_contravariantFlip gopurs_runtime.Value
var once_Data_Functor_Flip_contravariantFlip sync.Once

func Get_Data_Functor_Flip_contravariantFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_contravariantFlip.Do(func() {
		cache_Data_Functor_Flip_contravariantFlip = gopurs_runtime.Func(func(dictProfunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_contravariantFlip(dictProfunctor_0_box)
		})
	})
	return cache_Data_Functor_Flip_contravariantFlip
}

var cache_Data_Functor_Flip_categoryFlip gopurs_runtime.Value
var once_Data_Functor_Flip_categoryFlip sync.Once

func Get_Data_Functor_Flip_categoryFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_categoryFlip.Do(func() {
		cache_Data_Functor_Flip_categoryFlip = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_categoryFlip(dictCategory_0_box)
		})
	})
	return cache_Data_Functor_Flip_categoryFlip
}

var cache_Data_Functor_Flip_bifunctorFlip gopurs_runtime.Value
var once_Data_Functor_Flip_bifunctorFlip sync.Once

func Get_Data_Functor_Flip_bifunctorFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_bifunctorFlip.Do(func() {
		cache_Data_Functor_Flip_bifunctorFlip = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_bifunctorFlip(dictBifunctor_0_box)
		})
	})
	return cache_Data_Functor_Flip_bifunctorFlip
}

var cache_Data_Functor_Flip_biapplyFlip gopurs_runtime.Value
var once_Data_Functor_Flip_biapplyFlip sync.Once

func Get_Data_Functor_Flip_biapplyFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_biapplyFlip.Do(func() {
		cache_Data_Functor_Flip_biapplyFlip = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_biapplyFlip(dictBiapply_0_box)
		})
	})
	return cache_Data_Functor_Flip_biapplyFlip
}

var cache_Data_Functor_Flip_biapplicativeFlip gopurs_runtime.Value
var once_Data_Functor_Flip_biapplicativeFlip sync.Once

func Get_Data_Functor_Flip_biapplicativeFlip() gopurs_runtime.Value {
	once_Data_Functor_Flip_biapplicativeFlip.Do(func() {
		cache_Data_Functor_Flip_biapplicativeFlip = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Flip_biapplicativeFlip(dictBiapplicative_0_box)
		})
	})
	return cache_Data_Functor_Flip_biapplicativeFlip
}

func Call_Data_Functor_Flip_Flip(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Functor_Flip_showFlip(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Flip ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Functor_Flip_semigroupoidFlip(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
	_ = dictSemigroupoid_0
	// TAST (Let): compose_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar p$scope8) [(TypeVar c$scope15), (TypeVar b$scope14)]), (TypeApp (TypeVar p$scope8) [(TypeVar d$scope16), (TypeVar c$scope15)])] (TypeApp (TypeVar p$scope8) [(TypeVar d$scope16), (TypeVar b$scope14)]))
	compose_1_0 := Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0))
	_ = compose_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer((&Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(compose_1_0, v1_3, v_2)
	})}))}
}

func Call_Data_Functor_Flip_ordFlip(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Functor_Flip_functorFlip(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
	_ = dictBifunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_3
		}), v_2)
	})}))}
}

func Call_Data_Functor_Flip_eqFlip(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Functor_Flip_contravariantFlip(dictProfunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
	_ = dictProfunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), f_1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_3
		}), v_2)
	})}))}
}

func Call_Data_Functor_Flip_categoryFlip(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
	_ = dictCategory_0
	// TAST (Let): semigroupoidFlip1_1_0 shape=App(Var) bindingType=(TypeApp (ADT ["Control","Semigroupoid","Semigroupoid"] []) [(TypeApp (TypeApp (TypeVar p) [(TypeVar b), (TypeVar a)]) [(TypeVar p$scope73)])])
	semigroupoidFlip1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Call_Data_Functor_Flip_semigroupoidFlip(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})))
	_ = semigroupoidFlip1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer((&Constructor_Control_Category_Category[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(semigroupoidFlip1_1_0)}
	}), Call_Control_Category_identity(dictCategory_0)}))}
}

func Call_Data_Functor_Flip_bifunctorFlip(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
	_ = dictBifunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), g_2, f_1, v_3)
	})}))}
}

func Call_Data_Functor_Flip_biapplyFlip(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
	_ = dictBiapply_0
	// TAST (Let): bifunctorFlip1_1_0 shape=App(Var) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar b), (TypeVar a)]) [(TypeVar p$scope101)])])
	bifunctorFlip1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Call_Data_Functor_Flip_bifunctorFlip(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})))
	_ = bifunctorFlip1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer((&Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorFlip1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), v_2, v1_3)
	})}))}
}

func Call_Data_Functor_Flip_biapplicativeFlip(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
	_ = dictBiapplicative_0
	// TAST (Let): biapplyFlip1_1_0 shape=App(Var) bindingType=(ADT ["Control","Biapply","Biapply"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar b), (TypeVar a)]) [(TypeVar p$scope117)])])
	biapplyFlip1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](Call_Data_Functor_Flip_biapplyFlip(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})))
	_ = biapplyFlip1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3949191309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(biapplyFlip1_1_0)}
	}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), b_3, a_2)
	})}))}
}
