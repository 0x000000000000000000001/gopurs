package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Dual_go__append gopurs_runtime.Value
var once_Data_Monoid_Dual_go__append sync.Once

func Get_Data_Monoid_Dual_go__append() gopurs_runtime.Value {
	once_Data_Monoid_Dual_go__append.Do(func() {
		cache_Data_Monoid_Dual_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Monoid_Dual_go__append
}

var cache_Data_Monoid_Dual_Dual gopurs_runtime.Value
var once_Data_Monoid_Dual_Dual sync.Once

func Get_Data_Monoid_Dual_Dual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_Dual.Do(func() {
		cache_Data_Monoid_Dual_Dual = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Dual_Dual(x_0_box)
		})
	})
	return cache_Data_Monoid_Dual_Dual
}

var cache_Data_Monoid_Dual_showDual gopurs_runtime.Value
var once_Data_Monoid_Dual_showDual sync.Once

func Get_Data_Monoid_Dual_showDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_showDual.Do(func() {
		cache_Data_Monoid_Dual_showDual = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Dual_showDual(dictShow_0_box)
		})
	})
	return cache_Data_Monoid_Dual_showDual
}

var cache_Data_Monoid_Dual_semigroupDual gopurs_runtime.Value
var once_Data_Monoid_Dual_semigroupDual sync.Once

func Get_Data_Monoid_Dual_semigroupDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_semigroupDual.Do(func() {
		cache_Data_Monoid_Dual_semigroupDual = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Dual_semigroupDual(dictSemigroup_0_box)
		})
	})
	return cache_Data_Monoid_Dual_semigroupDual
}

var cache_Data_Monoid_Dual_ordDual gopurs_runtime.Value
var once_Data_Monoid_Dual_ordDual sync.Once

func Get_Data_Monoid_Dual_ordDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_ordDual.Do(func() {
		cache_Data_Monoid_Dual_ordDual = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Dual_ordDual(dictOrd_0_box)
		})
	})
	return cache_Data_Monoid_Dual_ordDual
}

var cache_Data_Monoid_Dual_monoidDual gopurs_runtime.Value
var once_Data_Monoid_Dual_monoidDual sync.Once

func Get_Data_Monoid_Dual_monoidDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_monoidDual.Do(func() {
		cache_Data_Monoid_Dual_monoidDual = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Dual_monoidDual(dictMonoid_0_box)
		})
	})
	return cache_Data_Monoid_Dual_monoidDual
}

var cache_Data_Monoid_Dual_functorDual gopurs_runtime.Value
var once_Data_Monoid_Dual_functorDual sync.Once

func Get_Data_Monoid_Dual_functorDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_functorDual.Do(func() {
		cache_Data_Monoid_Dual_functorDual = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, m_1)
		})}))}
	})
	return cache_Data_Monoid_Dual_functorDual
}

var cache_Data_Monoid_Dual_eqDual gopurs_runtime.Value
var once_Data_Monoid_Dual_eqDual sync.Once

func Get_Data_Monoid_Dual_eqDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_eqDual.Do(func() {
		cache_Data_Monoid_Dual_eqDual = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Dual_eqDual(dictEq_0_box)
		})
	})
	return cache_Data_Monoid_Dual_eqDual
}

var cache_Data_Monoid_Dual_eq1Dual gopurs_runtime.Value
var once_Data_Monoid_Dual_eq1Dual sync.Once

func Get_Data_Monoid_Dual_eq1Dual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_eq1Dual.Do(func() {
		cache_Data_Monoid_Dual_eq1Dual = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))
		})}))}
	})
	return cache_Data_Monoid_Dual_eq1Dual
}

var cache_Data_Monoid_Dual_ord1Dual gopurs_runtime.Value
var once_Data_Monoid_Dual_ord1Dual sync.Once

func Get_Data_Monoid_Dual_ord1Dual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_ord1Dual.Do(func() {
		cache_Data_Monoid_Dual_ord1Dual = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Monoid_Dual_eq1Dual()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))
		})}))}
	})
	return cache_Data_Monoid_Dual_ord1Dual
}

var cache_Data_Monoid_Dual_boundedDual gopurs_runtime.Value
var once_Data_Monoid_Dual_boundedDual sync.Once

func Get_Data_Monoid_Dual_boundedDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_boundedDual.Do(func() {
		cache_Data_Monoid_Dual_boundedDual = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Dual_boundedDual(dictBounded_0_box)
		})
	})
	return cache_Data_Monoid_Dual_boundedDual
}

var cache_Data_Monoid_Dual_applyDual gopurs_runtime.Value
var once_Data_Monoid_Dual_applyDual sync.Once

func Get_Data_Monoid_Dual_applyDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_applyDual.Do(func() {
		cache_Data_Monoid_Dual_applyDual = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Dual_functorDual()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v_0, v1_1)
		})}))}
	})
	return cache_Data_Monoid_Dual_applyDual
}

var cache_Data_Monoid_Dual_bindDual gopurs_runtime.Value
var once_Data_Monoid_Dual_bindDual sync.Once

func Get_Data_Monoid_Dual_bindDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_bindDual.Do(func() {
		cache_Data_Monoid_Dual_bindDual = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Monoid_Dual_applyDual()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_1, v_0)
		})}))}
	})
	return cache_Data_Monoid_Dual_bindDual
}

var cache_Data_Monoid_Dual_applicativeDual gopurs_runtime.Value
var once_Data_Monoid_Dual_applicativeDual sync.Once

func Get_Data_Monoid_Dual_applicativeDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_applicativeDual.Do(func() {
		cache_Data_Monoid_Dual_applicativeDual = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Monoid_Dual_applyDual()))}
		}), Get_Data_Monoid_Dual_Dual()}))}
	})
	return cache_Data_Monoid_Dual_applicativeDual
}

var cache_Data_Monoid_Dual_monadDual gopurs_runtime.Value
var once_Data_Monoid_Dual_monadDual sync.Once

func Get_Data_Monoid_Dual_monadDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_monadDual.Do(func() {
		cache_Data_Monoid_Dual_monadDual = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Monoid_Dual_applicativeDual()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Monoid_Dual_bindDual()))}
		})}))}
	})
	return cache_Data_Monoid_Dual_monadDual
}

func Call_Data_Monoid_Dual_Dual(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Monoid_Dual_showDual(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Dual ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Monoid_Dual_semigroupDual(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v1_2, v_1)
	})}))}
}

func Call_Data_Monoid_Dual_ordDual(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Monoid_Dual_monoidDual(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): semigroupDual1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a$scope13)])
	semigroupDual1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Monoid_Dual_semigroupDual(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = semigroupDual1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDual1_1_0)}
	}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))}
}

func Call_Data_Monoid_Dual_eqDual(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Monoid_Dual_boundedDual(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0))}
}
