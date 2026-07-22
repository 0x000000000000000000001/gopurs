package Data_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var showLazy gopurs_runtime.Value
var once_showLazy sync.Once
func Get_showLazy() gopurs_runtime.Value {
	once_showLazy.Do(func() {
		showLazy = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(defer \\_ -> ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], gopurs_runtime.Apply(Get_force(), x_1))), gopurs_runtime.Str(")")))
})})
})
	})
	return showLazy
}

var semiringLazy gopurs_runtime.Value
var once_semiringLazy sync.Once
func Get_semiringLazy() gopurs_runtime.Value {
	once_semiringLazy.Do(func() {
		semiringLazy = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
zero_1_0 := dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
one_2_1 := dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["add"], gopurs_runtime.Apply(Get_force(), a_3)), gopurs_runtime.Apply(Get_force(), b_4))
}))
})
}), "zero": gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return zero_1_0
})), "mul": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], gopurs_runtime.Apply(Get_force(), a_3)), gopurs_runtime.Apply(Get_force(), b_4))
}))
})
}), "one": gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return one_2_1
}))})
})
	})
	return semiringLazy
}

var semigroupLazy gopurs_runtime.Value
var once_semigroupLazy sync.Once
func Get_semigroupLazy() gopurs_runtime.Value {
	once_semigroupLazy.Do(func() {
		semigroupLazy = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(Get_force(), a_1)), gopurs_runtime.Apply(Get_force(), b_2))
}))
})
})})
})
	})
	return semigroupLazy
}

var ringLazy gopurs_runtime.Value
var once_ringLazy sync.Once
func Get_ringLazy() gopurs_runtime.Value {
	once_ringLazy.Do(func() {
		ringLazy = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
semiringLazy1_1_0 := gopurs_runtime.Apply(Get_semiringLazy(), gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["sub"], gopurs_runtime.Apply(Get_force(), a_2)), gopurs_runtime.Apply(Get_force(), b_3))
}))
})
}), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringLazy1_1_0
})})
})
	})
	return ringLazy
}

var monoidLazy gopurs_runtime.Value
var once_monoidLazy sync.Once
func Get_monoidLazy() gopurs_runtime.Value {
	once_monoidLazy.Do(func() {
		monoidLazy = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
semigroupLazy1_2_1 := gopurs_runtime.Apply(Get_semigroupLazy(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupLazy1_2_1
})})
})
	})
	return monoidLazy
}

var lazyLazy gopurs_runtime.Value
var once_lazyLazy sync.Once
func Get_lazyLazy() gopurs_runtime.Value {
	once_lazyLazy.Do(func() {
		lazyLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"defer": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
})})
	})
	return lazyLazy
}

var functorLazy gopurs_runtime.Value
var once_functorLazy sync.Once
func Get_functorLazy() gopurs_runtime.Value {
	once_functorLazy.Do(func() {
		functorLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_force(), l_1))
}))
})
})})
	})
	return functorLazy
}

var functorWithIndexLazy gopurs_runtime.Value
var once_functorWithIndexLazy sync.Once
func Get_functorWithIndexLazy() gopurs_runtime.Value {
	once_functorWithIndexLazy.Do(func() {
		functorWithIndexLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_functorLazy().PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
})})
	})
	return functorWithIndexLazy
}

var invariantLazy gopurs_runtime.Value
var once_invariantLazy sync.Once
func Get_invariantLazy() gopurs_runtime.Value {
	once_invariantLazy.Do(func() {
		invariantLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_functorLazy().PtrVal.(map[string]gopurs_runtime.Value)["map"], f_0)
})
})})
	})
	return invariantLazy
}

var foldableLazy gopurs_runtime.Value
var once_foldableLazy sync.Once
func Get_foldableLazy() gopurs_runtime.Value {
	once_foldableLazy.Do(func() {
		foldableLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_force(), l_2)), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
})})
	})
	return foldableLazy
}

var foldableWithIndexLazy gopurs_runtime.Value
var once_foldableWithIndexLazy sync.Once
func Get_foldableWithIndexLazy() gopurs_runtime.Value {
	once_foldableWithIndexLazy.Do(func() {
		foldableWithIndexLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_foldableLazy().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_foldableLazy().PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableLazy().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
})})
	})
	return foldableWithIndexLazy
}

var traversableLazy gopurs_runtime.Value
var once_traversableLazy sync.Once
func Get_traversableLazy() gopurs_runtime.Value {
	once_traversableLazy.Do(func() {
		traversableLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
})), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2)))
})
})
}), "sequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
})), gopurs_runtime.Apply(Get_force(), l_1))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
})})
	})
	return traversableLazy
}

var traversableWithIndexLazy gopurs_runtime.Value
var once_traversableWithIndexLazy sync.Once
func Get_traversableWithIndexLazy() gopurs_runtime.Value {
	once_traversableWithIndexLazy.Do(func() {
		traversableWithIndexLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableLazy().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexLazy()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexLazy()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableLazy()
})})
	})
	return traversableWithIndexLazy
}

var foldable1Lazy gopurs_runtime.Value
var once_foldable1Lazy sync.Once
func Get_foldable1Lazy() gopurs_runtime.Value {
	once_foldable1Lazy.Do(func() {
		foldable1Lazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap1": gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), "foldr1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
})
}), "foldl1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
})})
	})
	return foldable1Lazy
}

var traversable1Lazy gopurs_runtime.Value
var once_traversable1Lazy sync.Once
func Get_traversable1Lazy() gopurs_runtime.Value {
	once_traversable1Lazy.Do(func() {
		traversable1Lazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
})), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2)))
})
})
}), "sequence1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
})), gopurs_runtime.Apply(Get_force(), l_1))
})
}), "Foldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1Lazy()
}), "Traversable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableLazy()
})})
	})
	return traversable1Lazy
}

var extendLazy gopurs_runtime.Value
var once_extendLazy sync.Once
func Get_extendLazy() gopurs_runtime.Value {
	once_extendLazy.Do(func() {
		extendLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
})})
	})
	return extendLazy
}

var eqLazy gopurs_runtime.Value
var once_eqLazy sync.Once
func Get_eqLazy() gopurs_runtime.Value {
	once_eqLazy.Do(func() {
		eqLazy = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], gopurs_runtime.Apply(Get_force(), x_1)), gopurs_runtime.Apply(Get_force(), y_2))
})
})})
})
	})
	return eqLazy
}

var ordLazy gopurs_runtime.Value
var once_ordLazy sync.Once
func Get_ordLazy() gopurs_runtime.Value {
	once_ordLazy.Do(func() {
		ordLazy = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eqLazy1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], gopurs_runtime.Apply(Get_force(), x_2)), gopurs_runtime.Apply(Get_force(), y_3))
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(Get_force(), x_3)), gopurs_runtime.Apply(Get_force(), y_4))
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqLazy1_2_1
})})
})
	})
	return ordLazy
}

var eq1Lazy gopurs_runtime.Value
var once_eq1Lazy sync.Once
func Get_eq1Lazy() gopurs_runtime.Value {
	once_eq1Lazy.Do(func() {
		eq1Lazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], gopurs_runtime.Apply(Get_force(), x_1)), gopurs_runtime.Apply(Get_force(), y_2))
})
})
})})
	})
	return eq1Lazy
}

var ord1Lazy gopurs_runtime.Value
var once_ord1Lazy sync.Once
func Get_ord1Lazy() gopurs_runtime.Value {
	once_ord1Lazy.Do(func() {
		ord1Lazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_ordLazy(), dictOrd_0).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Lazy()
})})
	})
	return ord1Lazy
}

var comonadLazy gopurs_runtime.Value
var once_comonadLazy sync.Once
func Get_comonadLazy() gopurs_runtime.Value {
	once_comonadLazy.Do(func() {
		comonadLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": Get_force(), "Extend0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendLazy()
})})
	})
	return comonadLazy
}

var commutativeRingLazy gopurs_runtime.Value
var once_commutativeRingLazy sync.Once
func Get_commutativeRingLazy() gopurs_runtime.Value {
	once_commutativeRingLazy.Do(func() {
		commutativeRingLazy = gopurs_runtime.Func(func(dictCommutativeRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
ringLazy1_1_0 := gopurs_runtime.Apply(Get_ringLazy(), gopurs_runtime.Apply(dictCommutativeRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
})})
})
	})
	return commutativeRingLazy
}

var euclideanRingLazy gopurs_runtime.Value
var once_euclideanRingLazy sync.Once
func Get_euclideanRingLazy() gopurs_runtime.Value {
	once_euclideanRingLazy.Do(func() {
		euclideanRingLazy = gopurs_runtime.Func(func(dictEuclideanRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
ringLazy1_1_0 := gopurs_runtime.Apply(Get_ringLazy(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEuclideanRing_0.PtrVal.(map[string]gopurs_runtime.Value)["CommutativeRing0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"degree": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictEuclideanRing_0.PtrVal.(map[string]gopurs_runtime.Value)["degree"], gopurs_runtime.Apply(Get_force(), x_2))
}), "div": gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEuclideanRing_0.PtrVal.(map[string]gopurs_runtime.Value)["div"], gopurs_runtime.Apply(Get_force(), a_2)), gopurs_runtime.Apply(Get_force(), b_3))
}))
})
}), "mod": gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEuclideanRing_0.PtrVal.(map[string]gopurs_runtime.Value)["mod"], gopurs_runtime.Apply(Get_force(), a_2)), gopurs_runtime.Apply(Get_force(), b_3))
}))
})
}), "CommutativeRing0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
})})
})})
})
	})
	return euclideanRingLazy
}

var boundedLazy gopurs_runtime.Value
var once_boundedLazy sync.Once
func Get_boundedLazy() gopurs_runtime.Value {
	once_boundedLazy.Do(func() {
		boundedLazy = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
top_1_0 := dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["top"]
bottom_2_1 := dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
ordLazy1_3_2 := gopurs_runtime.Apply(Get_ordLazy(), gopurs_runtime.Apply(dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return top_1_0
})), "bottom": gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bottom_2_1
})), "Ord0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordLazy1_3_2
})})
})
	})
	return boundedLazy
}

var applyLazy gopurs_runtime.Value
var once_applyLazy sync.Once
func Get_applyLazy() gopurs_runtime.Value {
	once_applyLazy.Do(func() {
		applyLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_force(), f_0), gopurs_runtime.Apply(Get_force(), x_1))
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
})})
	})
	return applyLazy
}

var bindLazy gopurs_runtime.Value
var once_bindLazy sync.Once
func Get_bindLazy() gopurs_runtime.Value {
	once_bindLazy.Do(func() {
		bindLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_0)))
}))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
})})
	})
	return bindLazy
}

var heytingAlgebraLazy gopurs_runtime.Value
var once_heytingAlgebraLazy sync.Once
func Get_heytingAlgebraLazy() gopurs_runtime.Value {
	once_heytingAlgebraLazy.Do(func() {
		heytingAlgebraLazy = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
ff_1_0 := dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"]
tt_2_1 := dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["tt"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ff_1_0
})), "tt": gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return tt_2_1
})), "implies": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["implies"], gopurs_runtime.Apply(Get_force(), a_3))
}))
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_force(), __local_var_5_2), gopurs_runtime.Apply(Get_force(), b_4))
}))
})
}), "conj": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["conj"], gopurs_runtime.Apply(Get_force(), a_3))
}))
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_force(), __local_var_5_3), gopurs_runtime.Apply(Get_force(), b_4))
}))
})
}), "disj": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], gopurs_runtime.Apply(Get_force(), a_3))
}))
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_force(), __local_var_5_4), gopurs_runtime.Apply(Get_force(), b_4))
}))
})
}), "not": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["not"], gopurs_runtime.Apply(Get_force(), a_3))
}))
})})
})
	})
	return heytingAlgebraLazy
}

var booleanAlgebraLazy gopurs_runtime.Value
var once_booleanAlgebraLazy sync.Once
func Get_booleanAlgebraLazy() gopurs_runtime.Value {
	once_booleanAlgebraLazy.Do(func() {
		booleanAlgebraLazy = gopurs_runtime.Func(func(dictBooleanAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraLazy1_1_0 := gopurs_runtime.Apply(Get_heytingAlgebraLazy(), gopurs_runtime.Apply(dictBooleanAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebra0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebra0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraLazy1_1_0
})})
})
	})
	return booleanAlgebraLazy
}

var applicativeLazy gopurs_runtime.Value
var once_applicativeLazy sync.Once
func Get_applicativeLazy() gopurs_runtime.Value {
	once_applicativeLazy.Do(func() {
		applicativeLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
})})
	})
	return applicativeLazy
}

var monadLazy gopurs_runtime.Value
var once_monadLazy sync.Once
func Get_monadLazy() gopurs_runtime.Value {
	once_monadLazy.Do(func() {
		monadLazy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeLazy()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindLazy()
})})
	})
	return monadLazy
}

func Get_defer_() gopurs_runtime.Value {
	return Defer_
}

func Get_force() gopurs_runtime.Value {
	return Force
}
