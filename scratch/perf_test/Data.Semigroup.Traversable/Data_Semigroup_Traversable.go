package Data_Semigroup_Traversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var traverse1 gopurs_runtime.Value
var once_traverse1 sync.Once
func Get_traverse1() gopurs_runtime.Value {
	once_traverse1.Do(func() {
		traverse1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["traverse1"]
})
	})
	return traverse1
}

var traversableTuple gopurs_runtime.Value
var once_traversableTuple sync.Once
func Get_traversableTuple() gopurs_runtime.Value {
	once_traversableTuple.Do(func() {
		traversableTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "sequence1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
}), "Foldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableTuple()
}), "Traversable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableTuple()
})})
	})
	return traversableTuple
}

var traversableIdentity gopurs_runtime.Value
var once_traversableIdentity sync.Once
func Get_traversableIdentity() gopurs_runtime.Value {
	once_traversableIdentity.Do(func() {
		traversableIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Identity.Get_Identity()), gopurs_runtime.Apply(f_1, v_2))
})
})
}), "sequence1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Identity.Get_Identity()), v_1)
})
}), "Foldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableIdentity()
}), "Traversable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableIdentity()
})})
	})
	return traversableIdentity
}

var sequence1Default gopurs_runtime.Value
var once_sequence1Default sync.Once
func Get_sequence1Default() gopurs_runtime.Value {
	once_sequence1Default.Do(func() {
		sequence1Default = gopurs_runtime.Func(func(dictTraversable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictTraversable1_0.PtrVal.(map[string]gopurs_runtime.Value)["traverse1"], dictApply_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return sequence1Default
}

var traversableDual gopurs_runtime.Value
var once_traversableDual sync.Once
func Get_traversableDual() gopurs_runtime.Value {
	once_traversableDual.Do(func() {
		traversableDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Monoid_Dual.Get_Dual()), gopurs_runtime.Apply(f_1, v_2))
})
})
}), "sequence1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableDual().PtrVal.(map[string]gopurs_runtime.Value)["traverse1"], dictApply_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Foldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableDual()
}), "Traversable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableDual()
})})
	})
	return traversableDual
}

var traversableMultiplicative gopurs_runtime.Value
var once_traversableMultiplicative sync.Once
func Get_traversableMultiplicative() gopurs_runtime.Value {
	once_traversableMultiplicative.Do(func() {
		traversableMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Monoid_Multiplicative.Get_Multiplicative()), gopurs_runtime.Apply(f_1, v_2))
})
})
}), "sequence1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableMultiplicative().PtrVal.(map[string]gopurs_runtime.Value)["traverse1"], dictApply_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Foldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup_Foldable.Get_foldableMultiplicative()
}), "Traversable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableMultiplicative()
})})
	})
	return traversableMultiplicative
}

var sequence1 gopurs_runtime.Value
var once_sequence1 sync.Once
func Get_sequence1() gopurs_runtime.Value {
	once_sequence1.Do(func() {
		sequence1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["sequence1"]
})
	})
	return sequence1
}

var traverse1Default gopurs_runtime.Value
var once_traverse1Default sync.Once
func Get_traverse1Default() gopurs_runtime.Value {
	once_traverse1Default.Do(func() {
		traverse1Default = gopurs_runtime.Func(func(dictTraversable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
sequence12_2_0 := gopurs_runtime.Apply(dictTraversable1_0.PtrVal.(map[string]gopurs_runtime.Value)["sequence1"], dictApply_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence12_2_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictTraversable1_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3), ta_4))
})
})
})
})
	})
	return traverse1Default
}


