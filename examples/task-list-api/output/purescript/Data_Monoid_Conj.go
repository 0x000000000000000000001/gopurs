package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Conj_go__append gopurs_runtime.Value
var once_Data_Monoid_Conj_go__append sync.Once

func Get_Data_Monoid_Conj_go__append() gopurs_runtime.Value {
	once_Data_Monoid_Conj_go__append.Do(func() {
		cache_Data_Monoid_Conj_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Monoid_Conj_go__append
}

var cache_Data_Monoid_Conj_Conj gopurs_runtime.Value
var once_Data_Monoid_Conj_Conj sync.Once

func Get_Data_Monoid_Conj_Conj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_Conj.Do(func() {
		cache_Data_Monoid_Conj_Conj = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_Conj(x_0_box)
		})
	})
	return cache_Data_Monoid_Conj_Conj
}

var cache_Data_Monoid_Conj_showConj gopurs_runtime.Value
var once_Data_Monoid_Conj_showConj sync.Once

func Get_Data_Monoid_Conj_showConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_showConj.Do(func() {
		cache_Data_Monoid_Conj_showConj = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_showConj(dictShow_0_box)
		})
	})
	return cache_Data_Monoid_Conj_showConj
}

var cache_Data_Monoid_Conj_semiringConj gopurs_runtime.Value
var once_Data_Monoid_Conj_semiringConj sync.Once

func Get_Data_Monoid_Conj_semiringConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_semiringConj.Do(func() {
		cache_Data_Monoid_Conj_semiringConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_semiringConj(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Monoid_Conj_semiringConj
}

var cache_Data_Monoid_Conj_semigroupConj gopurs_runtime.Value
var once_Data_Monoid_Conj_semigroupConj sync.Once

func Get_Data_Monoid_Conj_semigroupConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_semigroupConj.Do(func() {
		cache_Data_Monoid_Conj_semigroupConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_semigroupConj(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Monoid_Conj_semigroupConj
}

var cache_Data_Monoid_Conj_ordConj gopurs_runtime.Value
var once_Data_Monoid_Conj_ordConj sync.Once

func Get_Data_Monoid_Conj_ordConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_ordConj.Do(func() {
		cache_Data_Monoid_Conj_ordConj = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_ordConj(dictOrd_0_box)
		})
	})
	return cache_Data_Monoid_Conj_ordConj
}

var cache_Data_Monoid_Conj_monoidConj gopurs_runtime.Value
var once_Data_Monoid_Conj_monoidConj sync.Once

func Get_Data_Monoid_Conj_monoidConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_monoidConj.Do(func() {
		cache_Data_Monoid_Conj_monoidConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_monoidConj(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Monoid_Conj_monoidConj
}

var cache_Data_Monoid_Conj_functorConj gopurs_runtime.Value
var once_Data_Monoid_Conj_functorConj sync.Once

func Get_Data_Monoid_Conj_functorConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_functorConj.Do(func() {
		cache_Data_Monoid_Conj_functorConj = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, m_1)
		})}))}
	})
	return cache_Data_Monoid_Conj_functorConj
}

var cache_Data_Monoid_Conj_eqConj gopurs_runtime.Value
var once_Data_Monoid_Conj_eqConj sync.Once

func Get_Data_Monoid_Conj_eqConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_eqConj.Do(func() {
		cache_Data_Monoid_Conj_eqConj = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_eqConj(dictEq_0_box)
		})
	})
	return cache_Data_Monoid_Conj_eqConj
}

var cache_Data_Monoid_Conj_eq1Conj gopurs_runtime.Value
var once_Data_Monoid_Conj_eq1Conj sync.Once

func Get_Data_Monoid_Conj_eq1Conj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_eq1Conj.Do(func() {
		cache_Data_Monoid_Conj_eq1Conj = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))
		})}))}
	})
	return cache_Data_Monoid_Conj_eq1Conj
}

var cache_Data_Monoid_Conj_ord1Conj gopurs_runtime.Value
var once_Data_Monoid_Conj_ord1Conj sync.Once

func Get_Data_Monoid_Conj_ord1Conj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_ord1Conj.Do(func() {
		cache_Data_Monoid_Conj_ord1Conj = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Monoid_Conj_eq1Conj()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))
		})}))}
	})
	return cache_Data_Monoid_Conj_ord1Conj
}

var cache_Data_Monoid_Conj_boundedConj gopurs_runtime.Value
var once_Data_Monoid_Conj_boundedConj sync.Once

func Get_Data_Monoid_Conj_boundedConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_boundedConj.Do(func() {
		cache_Data_Monoid_Conj_boundedConj = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Conj_boundedConj(dictBounded_0_box)
		})
	})
	return cache_Data_Monoid_Conj_boundedConj
}

var cache_Data_Monoid_Conj_applyConj gopurs_runtime.Value
var once_Data_Monoid_Conj_applyConj sync.Once

func Get_Data_Monoid_Conj_applyConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_applyConj.Do(func() {
		cache_Data_Monoid_Conj_applyConj = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Conj_functorConj()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v_0, v1_1)
		})}))}
	})
	return cache_Data_Monoid_Conj_applyConj
}

var cache_Data_Monoid_Conj_bindConj gopurs_runtime.Value
var once_Data_Monoid_Conj_bindConj sync.Once

func Get_Data_Monoid_Conj_bindConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_bindConj.Do(func() {
		cache_Data_Monoid_Conj_bindConj = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Monoid_Conj_applyConj()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_1, v_0)
		})}))}
	})
	return cache_Data_Monoid_Conj_bindConj
}

var cache_Data_Monoid_Conj_applicativeConj gopurs_runtime.Value
var once_Data_Monoid_Conj_applicativeConj sync.Once

func Get_Data_Monoid_Conj_applicativeConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_applicativeConj.Do(func() {
		cache_Data_Monoid_Conj_applicativeConj = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Monoid_Conj_applyConj()))}
		}), Get_Data_Monoid_Conj_Conj()}))}
	})
	return cache_Data_Monoid_Conj_applicativeConj
}

var cache_Data_Monoid_Conj_monadConj gopurs_runtime.Value
var once_Data_Monoid_Conj_monadConj sync.Once

func Get_Data_Monoid_Conj_monadConj() gopurs_runtime.Value {
	once_Data_Monoid_Conj_monadConj.Do(func() {
		cache_Data_Monoid_Conj_monadConj = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Monoid_Conj_applicativeConj()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Monoid_Conj_bindConj()))}
		})}))}
	})
	return cache_Data_Monoid_Conj_monadConj
}

func Call_Data_Monoid_Conj_Conj(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Monoid_Conj_showConj(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Conj ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Monoid_Conj_semiringConj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
	}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")}))}
}

func Call_Data_Monoid_Conj_semigroupConj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
	})}))}
}

func Call_Data_Monoid_Conj_ordConj(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Monoid_Conj_monoidConj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	// TAST (Let): semigroupConj1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a$scope20)])
	semigroupConj1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Monoid_Conj_semigroupConj(dictHeytingAlgebra_0))
	_ = semigroupConj1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupConj1_1_0)}
	}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")}))}
}

func Call_Data_Monoid_Conj_eqConj(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Monoid_Conj_boundedConj(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0))}
}
