package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Const_go__append gopurs_runtime.Value
var once_Data_Const_go__append sync.Once

func Get_Data_Const_go__append() gopurs_runtime.Value {
	once_Data_Const_go__append.Do(func() {
		cache_Data_Const_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Const_go__append
}

var cache_Data_Const_Const gopurs_runtime.Value
var once_Data_Const_Const sync.Once

func Get_Data_Const_Const() gopurs_runtime.Value {
	once_Data_Const_Const.Do(func() {
		cache_Data_Const_Const = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_Const(x_0_box)
		})
	})
	return cache_Data_Const_Const
}

var cache_Data_Const_showConst gopurs_runtime.Value
var once_Data_Const_showConst sync.Once

func Get_Data_Const_showConst() gopurs_runtime.Value {
	once_Data_Const_showConst.Do(func() {
		cache_Data_Const_showConst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_showConst(dictShow_0_box)
		})
	})
	return cache_Data_Const_showConst
}

var cache_Data_Const_semiringConst gopurs_runtime.Value
var once_Data_Const_semiringConst sync.Once

func Get_Data_Const_semiringConst() gopurs_runtime.Value {
	once_Data_Const_semiringConst.Do(func() {
		cache_Data_Const_semiringConst = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_semiringConst(dictSemiring_0_box)
		})
	})
	return cache_Data_Const_semiringConst
}

var cache_Data_Const_semigroupoidConst gopurs_runtime.Value
var once_Data_Const_semigroupoidConst sync.Once

func Get_Data_Const_semigroupoidConst() gopurs_runtime.Value {
	once_Data_Const_semigroupoidConst.Do(func() {
		cache_Data_Const_semigroupoidConst = gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer((&Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return v1_1
		})}))}
	})
	return cache_Data_Const_semigroupoidConst
}

var cache_Data_Const_semigroupConst gopurs_runtime.Value
var once_Data_Const_semigroupConst sync.Once

func Get_Data_Const_semigroupConst() gopurs_runtime.Value {
	once_Data_Const_semigroupConst.Do(func() {
		cache_Data_Const_semigroupConst = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_semigroupConst(dictSemigroup_0_box)
		})
	})
	return cache_Data_Const_semigroupConst
}

var cache_Data_Const_ringConst gopurs_runtime.Value
var once_Data_Const_ringConst sync.Once

func Get_Data_Const_ringConst() gopurs_runtime.Value {
	once_Data_Const_ringConst.Do(func() {
		cache_Data_Const_ringConst = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_ringConst(dictRing_0_box)
		})
	})
	return cache_Data_Const_ringConst
}

var cache_Data_Const_ordConst gopurs_runtime.Value
var once_Data_Const_ordConst sync.Once

func Get_Data_Const_ordConst() gopurs_runtime.Value {
	once_Data_Const_ordConst.Do(func() {
		cache_Data_Const_ordConst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_ordConst(dictOrd_0_box)
		})
	})
	return cache_Data_Const_ordConst
}

var cache_Data_Const_newtypeConst gopurs_runtime.Value
var once_Data_Const_newtypeConst sync.Once

func Get_Data_Const_newtypeConst() gopurs_runtime.Value {
	once_Data_Const_newtypeConst.Do(func() {
		cache_Data_Const_newtypeConst = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Const_newtypeConst
}

var cache_Data_Const_monoidConst gopurs_runtime.Value
var once_Data_Const_monoidConst sync.Once

func Get_Data_Const_monoidConst() gopurs_runtime.Value {
	once_Data_Const_monoidConst.Do(func() {
		cache_Data_Const_monoidConst = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_monoidConst(dictMonoid_0_box)
		})
	})
	return cache_Data_Const_monoidConst
}

var cache_Data_Const_heytingAlgebraConst gopurs_runtime.Value
var once_Data_Const_heytingAlgebraConst sync.Once

func Get_Data_Const_heytingAlgebraConst() gopurs_runtime.Value {
	once_Data_Const_heytingAlgebraConst.Do(func() {
		cache_Data_Const_heytingAlgebraConst = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_heytingAlgebraConst(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Const_heytingAlgebraConst
}

var cache_Data_Const_functorConst gopurs_runtime.Value
var once_Data_Const_functorConst sync.Once

func Get_Data_Const_functorConst() gopurs_runtime.Value {
	once_Data_Const_functorConst.Do(func() {
		cache_Data_Const_functorConst = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return m_1
		})}))}
	})
	return cache_Data_Const_functorConst
}

var cache_Data_Const_invariantConst gopurs_runtime.Value
var once_Data_Const_invariantConst sync.Once

func Get_Data_Const_invariantConst() gopurs_runtime.Value {
	once_Data_Const_invariantConst.Do(func() {
		cache_Data_Const_invariantConst = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), Get_Data_Const_functorConst())}))}
	})
	return cache_Data_Const_invariantConst
}

var cache_Data_Const_euclideanRingConst gopurs_runtime.Value
var once_Data_Const_euclideanRingConst sync.Once

func Get_Data_Const_euclideanRingConst() gopurs_runtime.Value {
	once_Data_Const_euclideanRingConst.Do(func() {
		cache_Data_Const_euclideanRingConst = gopurs_runtime.Func(func(dictEuclideanRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_euclideanRingConst(dictEuclideanRing_0_box)
		})
	})
	return cache_Data_Const_euclideanRingConst
}

var cache_Data_Const_eqConst gopurs_runtime.Value
var once_Data_Const_eqConst sync.Once

func Get_Data_Const_eqConst() gopurs_runtime.Value {
	once_Data_Const_eqConst.Do(func() {
		cache_Data_Const_eqConst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_eqConst(dictEq_0_box)
		})
	})
	return cache_Data_Const_eqConst
}

var cache_Data_Const_eq1Const gopurs_runtime.Value
var once_Data_Const_eq1Const sync.Once

func Get_Data_Const_eq1Const() gopurs_runtime.Value {
	once_Data_Const_eq1Const.Do(func() {
		cache_Data_Const_eq1Const = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_eq1Const(dictEq_0_box)
		})
	})
	return cache_Data_Const_eq1Const
}

var cache_Data_Const_ord1Const gopurs_runtime.Value
var once_Data_Const_ord1Const sync.Once

func Get_Data_Const_ord1Const() gopurs_runtime.Value {
	once_Data_Const_ord1Const.Do(func() {
		cache_Data_Const_ord1Const = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_ord1Const(dictOrd_0_box)
		})
	})
	return cache_Data_Const_ord1Const
}

var cache_Data_Const_commutativeRingConst gopurs_runtime.Value
var once_Data_Const_commutativeRingConst sync.Once

func Get_Data_Const_commutativeRingConst() gopurs_runtime.Value {
	once_Data_Const_commutativeRingConst.Do(func() {
		cache_Data_Const_commutativeRingConst = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_commutativeRingConst(dictCommutativeRing_0_box)
		})
	})
	return cache_Data_Const_commutativeRingConst
}

var cache_Data_Const_boundedConst gopurs_runtime.Value
var once_Data_Const_boundedConst sync.Once

func Get_Data_Const_boundedConst() gopurs_runtime.Value {
	once_Data_Const_boundedConst.Do(func() {
		cache_Data_Const_boundedConst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_boundedConst(dictBounded_0_box)
		})
	})
	return cache_Data_Const_boundedConst
}

var cache_Data_Const_booleanAlgebraConst gopurs_runtime.Value
var once_Data_Const_booleanAlgebraConst sync.Once

func Get_Data_Const_booleanAlgebraConst() gopurs_runtime.Value {
	once_Data_Const_booleanAlgebraConst.Do(func() {
		cache_Data_Const_booleanAlgebraConst = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_booleanAlgebraConst(dictBooleanAlgebra_0_box)
		})
	})
	return cache_Data_Const_booleanAlgebraConst
}

var cache_Data_Const_applyConst gopurs_runtime.Value
var once_Data_Const_applyConst sync.Once

func Get_Data_Const_applyConst() gopurs_runtime.Value {
	once_Data_Const_applyConst.Do(func() {
		cache_Data_Const_applyConst = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_applyConst(dictSemigroup_0_box)
		})
	})
	return cache_Data_Const_applyConst
}

var cache_Data_Const_applicativeConst gopurs_runtime.Value
var once_Data_Const_applicativeConst sync.Once

func Get_Data_Const_applicativeConst() gopurs_runtime.Value {
	once_Data_Const_applicativeConst.Do(func() {
		cache_Data_Const_applicativeConst = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Const_applicativeConst(dictMonoid_0_box)
		})
	})
	return cache_Data_Const_applicativeConst
}

func Call_Data_Const_Const(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Const_showConst(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Const ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Const_semiringConst(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
	_ = dictSemiring_0
	return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_0))}
}

func Call_Data_Const_semigroupConst(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))}
}

func Call_Data_Const_ringConst(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
	_ = dictRing_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dictRing_0))}
}

func Call_Data_Const_ordConst(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Const_monoidConst(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}
}

func Call_Data_Const_heytingAlgebraConst(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	return gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_0))}
}

func Call_Data_Const_euclideanRingConst(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
	_ = dictEuclideanRing_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_0))}
}

func Call_Data_Const_eqConst(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Const_eq1Const(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): eq_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar a) [(TypeVar a$scope78), (TypeVar a$scope81)]), (TypeApp (TypeVar a) [(TypeVar a$scope78), (TypeVar a$scope81)])] Boolean)
	eq_1_0 := Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))
	_ = eq_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return eq_1_0
	})}))}
}

func Call_Data_Const_ord1Const(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): compare_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar a) [(TypeVar a$scope33), (TypeVar a$scope36)]), (TypeApp (TypeVar a) [(TypeVar a$scope33), (TypeVar a$scope36)])] (ADT ["Data","Ordering","Ordering"] []))
	compare_1_0 := Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))
	_ = compare_1_0
	// TAST (Let): eq1Const1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (TypeVar a) [(TypeVar a$scope33)])])
	eq1Const1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_Const_eq1Const(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
	_ = eq1Const1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Const1_2_1)}
	}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return compare_1_0
	})}))}
}

func Call_Data_Const_commutativeRingConst(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
	_ = dictCommutativeRing_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]](dictCommutativeRing_0))}
}

func Call_Data_Const_boundedConst(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0))}
}

func Call_Data_Const_booleanAlgebraConst(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
	_ = dictBooleanAlgebra_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3257204378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]](dictBooleanAlgebra_0))}
}

func Call_Data_Const_applyConst(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Const_functorConst()
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
	})}))}
}

func Call_Data_Const_applicativeConst(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar a$scope105)
	mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_1_0
	// TAST (Let): applyConst1_2_1 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar a) [(TypeVar a$scope105)])])
	applyConst1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Const_applyConst(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
	_ = applyConst1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyConst1_2_1)}
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return mempty_1_0
	})}))}
}
