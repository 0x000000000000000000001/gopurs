package Data_Array_NonEmpty_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var NonEmptyArray gopurs_runtime.Value
var once_NonEmptyArray sync.Once
func Get_NonEmptyArray() gopurs_runtime.Value {
	once_NonEmptyArray.Do(func() {
		NonEmptyArray = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NonEmptyArray
}

var unfoldable1NonEmptyArray gopurs_runtime.Value
var once_unfoldable1NonEmptyArray sync.Once
func Get_unfoldable1NonEmptyArray() gopurs_runtime.Value {
	once_unfoldable1NonEmptyArray.Do(func() {
		unfoldable1NonEmptyArray = pkg_Data_Unfoldable1.Get_unfoldable1Array()
	})
	return unfoldable1NonEmptyArray
}

var traversableWithIndexNonEmptyArray gopurs_runtime.Value
var once_traversableWithIndexNonEmptyArray sync.Once
func Get_traversableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_traversableWithIndexNonEmptyArray.Do(func() {
		traversableWithIndexNonEmptyArray = pkg_Data_TraversableWithIndex.Get_traversableWithIndexArray()
	})
	return traversableWithIndexNonEmptyArray
}

var traversableNonEmptyArray gopurs_runtime.Value
var once_traversableNonEmptyArray sync.Once
func Get_traversableNonEmptyArray() gopurs_runtime.Value {
	once_traversableNonEmptyArray.Do(func() {
		traversableNonEmptyArray = pkg_Data_Traversable.Get_traversableArray()
	})
	return traversableNonEmptyArray
}

var showNonEmptyArray gopurs_runtime.Value
var once_showNonEmptyArray sync.Once
func Get_showNonEmptyArray() gopurs_runtime.Value {
	once_showNonEmptyArray.Do(func() {
		showNonEmptyArray = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(NonEmptyArray ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"]), v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showNonEmptyArray
}

var semigroupNonEmptyArray gopurs_runtime.Value
var once_semigroupNonEmptyArray sync.Once
func Get_semigroupNonEmptyArray() gopurs_runtime.Value {
	once_semigroupNonEmptyArray.Do(func() {
		semigroupNonEmptyArray = pkg_Data_Semigroup.Get_semigroupArray()
	})
	return semigroupNonEmptyArray
}

var ordNonEmptyArray gopurs_runtime.Value
var once_ordNonEmptyArray sync.Once
func Get_ordNonEmptyArray() gopurs_runtime.Value {
	once_ordNonEmptyArray.Do(func() {
		ordNonEmptyArray = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Ord.Get_ordArray(), dictOrd_0)
})
	})
	return ordNonEmptyArray
}

var ord1NonEmptyArray gopurs_runtime.Value
var once_ord1NonEmptyArray sync.Once
func Get_ord1NonEmptyArray() gopurs_runtime.Value {
	once_ord1NonEmptyArray.Do(func() {
		ord1NonEmptyArray = pkg_Data_Ord.Get_ord1Array()
	})
	return ord1NonEmptyArray
}

var monadNonEmptyArray gopurs_runtime.Value
var once_monadNonEmptyArray sync.Once
func Get_monadNonEmptyArray() gopurs_runtime.Value {
	once_monadNonEmptyArray.Do(func() {
		monadNonEmptyArray = pkg_Control_Monad.Get_monadArray()
	})
	return monadNonEmptyArray
}

var functorWithIndexNonEmptyArray gopurs_runtime.Value
var once_functorWithIndexNonEmptyArray sync.Once
func Get_functorWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyArray.Do(func() {
		functorWithIndexNonEmptyArray = pkg_Data_FunctorWithIndex.Get_functorWithIndexArray()
	})
	return functorWithIndexNonEmptyArray
}

var functorNonEmptyArray gopurs_runtime.Value
var once_functorNonEmptyArray sync.Once
func Get_functorNonEmptyArray() gopurs_runtime.Value {
	once_functorNonEmptyArray.Do(func() {
		functorNonEmptyArray = pkg_Data_Functor.Get_functorArray()
	})
	return functorNonEmptyArray
}

var foldableWithIndexNonEmptyArray gopurs_runtime.Value
var once_foldableWithIndexNonEmptyArray sync.Once
func Get_foldableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyArray.Do(func() {
		foldableWithIndexNonEmptyArray = pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray()
	})
	return foldableWithIndexNonEmptyArray
}

var foldableNonEmptyArray gopurs_runtime.Value
var once_foldableNonEmptyArray sync.Once
func Get_foldableNonEmptyArray() gopurs_runtime.Value {
	once_foldableNonEmptyArray.Do(func() {
		foldableNonEmptyArray = pkg_Data_Foldable.Get_foldableArray()
	})
	return foldableNonEmptyArray
}

var foldable1NonEmptyArray gopurs_runtime.Value
var once_foldable1NonEmptyArray sync.Once
func Get_foldable1NonEmptyArray() gopurs_runtime.Value {
	once_foldable1NonEmptyArray.Do(func() {
		foldable1NonEmptyArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap1": gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
append_1_0 := dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), f_2)
__local_var_4_2 := gopurs_runtime.Apply(Get_foldable1NonEmptyArray().PtrVal.(map[string]gopurs_runtime.Value)["foldl1"], append_1_0)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}), "foldr1": gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_foldr1Impl()), "foldl1": gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_foldl1Impl()), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
})})
	})
	return foldable1NonEmptyArray
}

var traversable1NonEmptyArray gopurs_runtime.Value
var once_traversable1NonEmptyArray sync.Once
func Get_traversable1NonEmptyArray() gopurs_runtime.Value {
	once_traversable1NonEmptyArray.Do(func() {
		traversable1NonEmptyArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
apply_1_0 := dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"]
map__2_1 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_traverse1Impl()), apply_1_0), map__2_1), f_3)
})
}), "sequence1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversable1NonEmptyArray().PtrVal.(map[string]gopurs_runtime.Value)["traverse1"], dictApply_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Foldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1NonEmptyArray()
}), "Traversable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray()
})})
	})
	return traversable1NonEmptyArray
}

var eqNonEmptyArray gopurs_runtime.Value
var once_eqNonEmptyArray sync.Once
func Get_eqNonEmptyArray() gopurs_runtime.Value {
	once_eqNonEmptyArray.Do(func() {
		eqNonEmptyArray = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])})
})
	})
	return eqNonEmptyArray
}

var eq1NonEmptyArray gopurs_runtime.Value
var once_eq1NonEmptyArray sync.Once
func Get_eq1NonEmptyArray() gopurs_runtime.Value {
	once_eq1NonEmptyArray.Do(func() {
		eq1NonEmptyArray = pkg_Data_Eq.Get_eq1Array()
	})
	return eq1NonEmptyArray
}

var bindNonEmptyArray gopurs_runtime.Value
var once_bindNonEmptyArray sync.Once
func Get_bindNonEmptyArray() gopurs_runtime.Value {
	once_bindNonEmptyArray.Do(func() {
		bindNonEmptyArray = pkg_Control_Bind.Get_bindArray()
	})
	return bindNonEmptyArray
}

var applyNonEmptyArray gopurs_runtime.Value
var once_applyNonEmptyArray sync.Once
func Get_applyNonEmptyArray() gopurs_runtime.Value {
	once_applyNonEmptyArray.Do(func() {
		applyNonEmptyArray = pkg_Control_Apply.Get_applyArray()
	})
	return applyNonEmptyArray
}

var applicativeNonEmptyArray gopurs_runtime.Value
var once_applicativeNonEmptyArray sync.Once
func Get_applicativeNonEmptyArray() gopurs_runtime.Value {
	once_applicativeNonEmptyArray.Do(func() {
		applicativeNonEmptyArray = pkg_Control_Applicative.Get_applicativeArray()
	})
	return applicativeNonEmptyArray
}

var altNonEmptyArray gopurs_runtime.Value
var once_altNonEmptyArray sync.Once
func Get_altNonEmptyArray() gopurs_runtime.Value {
	once_altNonEmptyArray.Do(func() {
		altNonEmptyArray = pkg_Control_Alt.Get_altArray()
	})
	return altNonEmptyArray
}

func Get_foldl1Impl() gopurs_runtime.Value {
	return Foldl1Impl
}

func Get_foldr1Impl() gopurs_runtime.Value {
	return Foldr1Impl
}

func Get_traverse1Impl() gopurs_runtime.Value {
	return Traverse1Impl
}
