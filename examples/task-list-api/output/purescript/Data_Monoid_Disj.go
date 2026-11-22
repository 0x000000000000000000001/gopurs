package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Disj_go__append gopurs_runtime.Value
var once_Data_Monoid_Disj_go__append sync.Once

func Get_Data_Monoid_Disj_go__append() gopurs_runtime.Value {
	once_Data_Monoid_Disj_go__append.Do(func() {
		cache_Data_Monoid_Disj_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Monoid_Disj_go__append
}

var cache_Data_Monoid_Disj_Disj gopurs_runtime.Value
var once_Data_Monoid_Disj_Disj sync.Once

func Get_Data_Monoid_Disj_Disj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_Disj.Do(func() {
		cache_Data_Monoid_Disj_Disj = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_Disj(x_0_box)
		})
	})
	return cache_Data_Monoid_Disj_Disj
}

var cache_Data_Monoid_Disj_showDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_showDisj sync.Once

func Get_Data_Monoid_Disj_showDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_showDisj.Do(func() {
		cache_Data_Monoid_Disj_showDisj = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_showDisj(dictShow_0_box)
		})
	})
	return cache_Data_Monoid_Disj_showDisj
}

var cache_Data_Monoid_Disj_semiringDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_semiringDisj sync.Once

func Get_Data_Monoid_Disj_semiringDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_semiringDisj.Do(func() {
		cache_Data_Monoid_Disj_semiringDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_semiringDisj(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Monoid_Disj_semiringDisj
}

var cache_Data_Monoid_Disj_semigroupDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_semigroupDisj sync.Once

func Get_Data_Monoid_Disj_semigroupDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_semigroupDisj.Do(func() {
		cache_Data_Monoid_Disj_semigroupDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_semigroupDisj(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Monoid_Disj_semigroupDisj
}

var cache_Data_Monoid_Disj_ordDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_ordDisj sync.Once

func Get_Data_Monoid_Disj_ordDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_ordDisj.Do(func() {
		cache_Data_Monoid_Disj_ordDisj = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_ordDisj(dictOrd_0_box)
		})
	})
	return cache_Data_Monoid_Disj_ordDisj
}

var cache_Data_Monoid_Disj_monoidDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_monoidDisj sync.Once

func Get_Data_Monoid_Disj_monoidDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_monoidDisj.Do(func() {
		cache_Data_Monoid_Disj_monoidDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_monoidDisj(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Monoid_Disj_monoidDisj
}

var cache_Data_Monoid_Disj_functorDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_functorDisj sync.Once

func Get_Data_Monoid_Disj_functorDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_functorDisj.Do(func() {
		cache_Data_Monoid_Disj_functorDisj = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, m_1)
		})}))}
	})
	return cache_Data_Monoid_Disj_functorDisj
}

var cache_Data_Monoid_Disj_eqDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_eqDisj sync.Once

func Get_Data_Monoid_Disj_eqDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_eqDisj.Do(func() {
		cache_Data_Monoid_Disj_eqDisj = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_eqDisj(dictEq_0_box)
		})
	})
	return cache_Data_Monoid_Disj_eqDisj
}

var cache_Data_Monoid_Disj_eq1Disj gopurs_runtime.Value
var once_Data_Monoid_Disj_eq1Disj sync.Once

func Get_Data_Monoid_Disj_eq1Disj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_eq1Disj.Do(func() {
		cache_Data_Monoid_Disj_eq1Disj = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))
		})}))}
	})
	return cache_Data_Monoid_Disj_eq1Disj
}

var cache_Data_Monoid_Disj_ord1Disj gopurs_runtime.Value
var once_Data_Monoid_Disj_ord1Disj sync.Once

func Get_Data_Monoid_Disj_ord1Disj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_ord1Disj.Do(func() {
		cache_Data_Monoid_Disj_ord1Disj = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Monoid_Disj_eq1Disj()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))
		})}))}
	})
	return cache_Data_Monoid_Disj_ord1Disj
}

var cache_Data_Monoid_Disj_boundedDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_boundedDisj sync.Once

func Get_Data_Monoid_Disj_boundedDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_boundedDisj.Do(func() {
		cache_Data_Monoid_Disj_boundedDisj = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Disj_boundedDisj(dictBounded_0_box)
		})
	})
	return cache_Data_Monoid_Disj_boundedDisj
}

var cache_Data_Monoid_Disj_applyDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_applyDisj sync.Once

func Get_Data_Monoid_Disj_applyDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_applyDisj.Do(func() {
		cache_Data_Monoid_Disj_applyDisj = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Disj_functorDisj()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v_0, v1_1)
		})}))}
	})
	return cache_Data_Monoid_Disj_applyDisj
}

var cache_Data_Monoid_Disj_bindDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_bindDisj sync.Once

func Get_Data_Monoid_Disj_bindDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_bindDisj.Do(func() {
		cache_Data_Monoid_Disj_bindDisj = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Monoid_Disj_applyDisj()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_1, v_0)
		})}))}
	})
	return cache_Data_Monoid_Disj_bindDisj
}

var cache_Data_Monoid_Disj_applicativeDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_applicativeDisj sync.Once

func Get_Data_Monoid_Disj_applicativeDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_applicativeDisj.Do(func() {
		cache_Data_Monoid_Disj_applicativeDisj = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Monoid_Disj_applyDisj()))}
		}), Get_Data_Monoid_Disj_Disj()}))}
	})
	return cache_Data_Monoid_Disj_applicativeDisj
}

var cache_Data_Monoid_Disj_monadDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_monadDisj sync.Once

func Get_Data_Monoid_Disj_monadDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_monadDisj.Do(func() {
		cache_Data_Monoid_Disj_monadDisj = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Monoid_Disj_applicativeDisj()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Monoid_Disj_bindDisj()))}
		})}))}
	})
	return cache_Data_Monoid_Disj_monadDisj
}

func Call_Data_Monoid_Disj_Disj(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Monoid_Disj_showDisj(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Disj ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Monoid_Disj_semiringDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
	}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")}))}
}

func Call_Data_Monoid_Disj_semigroupDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
	})}))}
}

func Call_Data_Monoid_Disj_ordDisj(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Monoid_Disj_monoidDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	// TAST (Let): semigroupDisj1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a$scope20)])
	semigroupDisj1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Monoid_Disj_semigroupDisj(dictHeytingAlgebra_0))
	_ = semigroupDisj1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDisj1_1_0)}
	}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")}))}
}

func Call_Data_Monoid_Disj_eqDisj(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Monoid_Disj_boundedDisj(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0))}
}
