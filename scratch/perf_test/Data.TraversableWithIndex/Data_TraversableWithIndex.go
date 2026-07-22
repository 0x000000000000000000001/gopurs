package Data_TraversableWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor_App "gopurs/output/Data.Functor.App"
	pkg_Data_Functor_Compose "gopurs/output/Data.Functor.Compose"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Traversable_Accum_Internal "gopurs/output/Data.Traversable.Accum.Internal"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var traverseWithIndexDefault gopurs_runtime.Value
var once_traverseWithIndexDefault sync.Once
func Get_traverseWithIndexDefault() gopurs_runtime.Value {
	once_traverseWithIndexDefault.Do(func() {
		traverseWithIndexDefault = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["sequence"], dictApplicative_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], f_3)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
})
})
	})
	return traverseWithIndexDefault
}

var traverseWithIndex gopurs_runtime.Value
var once_traverseWithIndex sync.Once
func Get_traverseWithIndex() gopurs_runtime.Value {
	once_traverseWithIndex.Do(func() {
		traverseWithIndex = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"]
})
	})
	return traverseWithIndex
}

var traverseDefault gopurs_runtime.Value
var once_traverseDefault sync.Once
func Get_traverseDefault() gopurs_runtime.Value {
	once_traverseDefault.Do(func() {
		traverseDefault = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex2_2_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverseWithIndex2_2_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return f_3
}))
})
})
})
	})
	return traverseDefault
}

var traversableWithIndexTuple gopurs_runtime.Value
var once_traversableWithIndexTuple sync.Once
func Get_traversableWithIndexTuple() gopurs_runtime.Value {
	once_traversableWithIndexTuple.Do(func() {
		traversableWithIndexTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexTuple()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexTuple()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableTuple()
})})
	})
	return traversableWithIndexTuple
}

var traversableWithIndexProduct gopurs_runtime.Value
var once_traversableWithIndexProduct sync.Once
func Get_traversableWithIndexProduct() gopurs_runtime.Value {
	once_traversableWithIndexProduct.Do(func() {
		traversableWithIndexProduct = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
foldableWithIndexProduct_3_2 := gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_foldableWithIndexProduct(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableProduct_4_3 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableProduct(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictTraversableWithIndex1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{})
__local_var_7_5 := gopurs_runtime.Apply(__local_var_6_4.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorProduct1_8_7 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_8), v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_7_5.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_8), v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
functorWithIndexProduct1_8_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_11}))
})), v_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_6_4.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_11}))
})), v_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_8_7
})})
foldableWithIndexProduct1_9_8 := gopurs_runtime.Apply(foldableWithIndexProduct_3_2, gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableProduct1_10_9 := gopurs_runtime.Apply(traversableProduct_4_3, gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_11 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_12_10 := gopurs_runtime.Apply(dictApplicative_11.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
traverseWithIndex3_13_11 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_11)
traverseWithIndex4_14_12 := gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_11)
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_12_10.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_12_10.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Product.Get_product()), gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex3_13_11, gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_15, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_17}))
})), v_16.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex4_14_12, gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_15, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_17}))
})), v_16.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexProduct1_8_6
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexProduct1_9_8
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableProduct1_10_9
})})
})
})
	})
	return traversableWithIndexProduct
}

var traversableWithIndexMultiplicative gopurs_runtime.Value
var once_traversableWithIndexMultiplicative sync.Once
func Get_traversableWithIndexMultiplicative() gopurs_runtime.Value {
	once_traversableWithIndexMultiplicative.Do(func() {
		traversableWithIndexMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableMultiplicative().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexMultiplicative()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexMultiplicative()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableMultiplicative()
})})
	})
	return traversableWithIndexMultiplicative
}

var traversableWithIndexMaybe gopurs_runtime.Value
var once_traversableWithIndexMaybe sync.Once
func Get_traversableWithIndexMaybe() gopurs_runtime.Value {
	once_traversableWithIndexMaybe.Do(func() {
		traversableWithIndexMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableMaybe().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexMaybe()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexMaybe()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableMaybe()
})})
	})
	return traversableWithIndexMaybe
}

var traversableWithIndexLast gopurs_runtime.Value
var once_traversableWithIndexLast sync.Once
func Get_traversableWithIndexLast() gopurs_runtime.Value {
	once_traversableWithIndexLast.Do(func() {
		traversableWithIndexLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableLast().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexLast()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexLast()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableLast()
})})
	})
	return traversableWithIndexLast
}

var traversableWithIndexIdentity gopurs_runtime.Value
var once_traversableWithIndexIdentity sync.Once
func Get_traversableWithIndexIdentity() gopurs_runtime.Value {
	once_traversableWithIndexIdentity.Do(func() {
		traversableWithIndexIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Identity.Get_Identity()), gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()), v_2))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexIdentity()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexIdentity()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableIdentity()
})})
	})
	return traversableWithIndexIdentity
}

var traversableWithIndexFirst gopurs_runtime.Value
var once_traversableWithIndexFirst sync.Once
func Get_traversableWithIndexFirst() gopurs_runtime.Value {
	once_traversableWithIndexFirst.Do(func() {
		traversableWithIndexFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableFirst().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexFirst()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexFirst()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableFirst()
})})
	})
	return traversableWithIndexFirst
}

var traversableWithIndexEither gopurs_runtime.Value
var once_traversableWithIndexEither sync.Once
func Get_traversableWithIndexEither() gopurs_runtime.Value {
	once_traversableWithIndexEither.Do(func() {
		traversableWithIndexEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
} else {
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Right()), gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, pkg_Data_Unit.Get_unit()), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexEither()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexEither()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableEither()
})})
	})
	return traversableWithIndexEither
}

var traversableWithIndexDual gopurs_runtime.Value
var once_traversableWithIndexDual sync.Once
func Get_traversableWithIndexDual() gopurs_runtime.Value {
	once_traversableWithIndexDual.Do(func() {
		traversableWithIndexDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableDual().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexDual()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexDual()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableDual()
})})
	})
	return traversableWithIndexDual
}

var traversableWithIndexDisj gopurs_runtime.Value
var once_traversableWithIndexDisj sync.Once
func Get_traversableWithIndexDisj() gopurs_runtime.Value {
	once_traversableWithIndexDisj.Do(func() {
		traversableWithIndexDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableDisj().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexDisj()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexDisj()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableDisj()
})})
	})
	return traversableWithIndexDisj
}

var traversableWithIndexCoproduct gopurs_runtime.Value
var once_traversableWithIndexCoproduct sync.Once
func Get_traversableWithIndexCoproduct() gopurs_runtime.Value {
	once_traversableWithIndexCoproduct.Do(func() {
		traversableWithIndexCoproduct = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorWithIndexCoproduct_1_0 := gopurs_runtime.Apply(pkg_Data_FunctorWithIndex.Get_functorWithIndexCoproduct(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{}))
foldableWithIndexCoproduct_2_1 := gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_foldableWithIndexCoproduct(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableCoproduct_3_2 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableCoproduct(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
functorWithIndexCoproduct1_5_3 := gopurs_runtime.Apply(functorWithIndexCoproduct_1_0, gopurs_runtime.Apply(dictTraversableWithIndex1_4.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{}))
foldableWithIndexCoproduct1_6_4 := gopurs_runtime.Apply(foldableWithIndexCoproduct_2_1, gopurs_runtime.Apply(dictTraversableWithIndex1_4.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableCoproduct1_7_5 := gopurs_runtime.Apply(traversableCoproduct_3_2, gopurs_runtime.Apply(dictTraversableWithIndex1_4.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_8.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
traverseWithIndex3_10_7 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_8)
traverseWithIndex4_11_8 := gopurs_runtime.Apply(dictTraversableWithIndex1_4.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_8)
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_13_9 := gopurs_runtime.Apply(__local_var_9_6.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_13})
}))
__local_var_14_10 := gopurs_runtime.Apply(traverseWithIndex3_10_7, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_14}))
}))
__local_var_15_11 := gopurs_runtime.Apply(__local_var_9_6.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_15})
}))
__local_var_16_12 := gopurs_runtime.Apply(traverseWithIndex4_11_8, gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_16}))
}))
return gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_17.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t13 = gopurs_runtime.Apply(__local_var_13_9, gopurs_runtime.Apply(__local_var_14_10, v2_17.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v2_17.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t13 = gopurs_runtime.Apply(__local_var_15_11, gopurs_runtime.Apply(__local_var_16_12, v2_17.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t13
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexCoproduct1_5_3
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexCoproduct1_6_4
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableCoproduct1_7_5
})})
})
})
	})
	return traversableWithIndexCoproduct
}

var traversableWithIndexConst gopurs_runtime.Value
var once_traversableWithIndexConst sync.Once
func Get_traversableWithIndexConst() gopurs_runtime.Value {
	once_traversableWithIndexConst.Do(func() {
		traversableWithIndexConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], v1_2)
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexConst()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexConst()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableConst()
})})
	})
	return traversableWithIndexConst
}

var traversableWithIndexConj gopurs_runtime.Value
var once_traversableWithIndexConj sync.Once
func Get_traversableWithIndexConj() gopurs_runtime.Value {
	once_traversableWithIndexConj.Do(func() {
		traversableWithIndexConj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableConj().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexConj()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexConj()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableConj()
})})
	})
	return traversableWithIndexConj
}

var traversableWithIndexCompose gopurs_runtime.Value
var once_traversableWithIndexCompose sync.Once
func Get_traversableWithIndexCompose() gopurs_runtime.Value {
	once_traversableWithIndexCompose.Do(func() {
		traversableWithIndexCompose = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
foldableWithIndexCompose_3_2 := gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_foldableWithIndexCompose(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableCompose_4_3 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableCompose(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictTraversableWithIndex1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{})
__local_var_7_5 := gopurs_runtime.Apply(__local_var_6_4.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCompose1_8_7 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(__local_var_7_5.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_8)), v_9)
})
})})
functorWithIndexCompose1_8_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_11, "value1": b_12}))
}))
})), v_10)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_8_7
})})
foldableWithIndexCompose1_9_8 := gopurs_runtime.Apply(foldableWithIndexCompose_3_2, gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableCompose1_10_9 := gopurs_runtime.Apply(traversableCompose_4_3, gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_11 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex3_12_10 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_11)
traverseWithIndex4_13_11 := gopurs_runtime.Apply(dictTraversableWithIndex1_5.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_11)
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_11.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Compose.Get_Compose()), gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex3_12_10, gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverseWithIndex4_13_11, gopurs_runtime.Func(func(b_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_14, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_16, "value1": b_17}))
}))
})), v_15))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexCompose1_8_6
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexCompose1_9_8
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableCompose1_10_9
})})
})
})
	})
	return traversableWithIndexCompose
}

var traversableWithIndexArray gopurs_runtime.Value
var once_traversableWithIndexArray sync.Once
func Get_traversableWithIndexArray() gopurs_runtime.Value {
	once_traversableWithIndexArray.Do(func() {
		traversableWithIndexArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableWithIndexArray().PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["sequence"], dictApplicative_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableWithIndexArray().PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], f_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexArray()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray()
})})
	})
	return traversableWithIndexArray
}

var traversableWithIndexApp gopurs_runtime.Value
var once_traversableWithIndexApp sync.Once
func Get_traversableWithIndexApp() gopurs_runtime.Value {
	once_traversableWithIndexApp.Do(func() {
		traversableWithIndexApp = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWithIndexApp_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], f_3), v_4)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
})})
__local_var_4_3 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{})
__local_var_5_4 := gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{})
foldableApp_6_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_4.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_6), i_7), v_8)
})
})
}), "foldl": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_4.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_6), i_7), v_8)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_6)
})})
foldableWithIndexApp_6_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], f_7), z_8), v_9)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], f_7), z_8), v_9)
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_7)
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableApp_6_6
})})
traversableApp_7_7 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableApp(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex2_9_8 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_8)
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_8.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_App.Get_App()), gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex2_9_8, f_10), v_11))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexApp_3_2
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexApp_6_5
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableApp_7_7
})})
})
	})
	return traversableWithIndexApp
}

var traversableWithIndexAdditive gopurs_runtime.Value
var once_traversableWithIndexAdditive sync.Once
func Get_traversableWithIndexAdditive() gopurs_runtime.Value {
	once_traversableWithIndexAdditive.Do(func() {
		traversableWithIndexAdditive = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableAdditive().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexAdditive()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexAdditive()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableAdditive()
})})
	})
	return traversableWithIndexAdditive
}

var mapAccumRWithIndex gopurs_runtime.Value
var once_mapAccumRWithIndex sync.Once
func Get_mapAccumRWithIndex() gopurs_runtime.Value {
	once_mapAccumRWithIndex.Do(func() {
		mapAccumRWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex1_1_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR())
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, i_5), s_7), a_6)
})
})
})), xs_4), s0_3)
})
})
})
})
	})
	return mapAccumRWithIndex
}

var scanrWithIndex gopurs_runtime.Value
var once_scanrWithIndex sync.Once
func Get_scanrWithIndex() gopurs_runtime.Value {
	once_scanrWithIndex.Do(func() {
		scanrWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
mapAccumRWithIndex1_1_0 := gopurs_runtime.Apply(Get_mapAccumRWithIndex(), dictTraversableWithIndex_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(mapAccumRWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_8_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, i_5), a_7), b_6)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": b_prime_8_1, "value": b_prime_8_1})
})
})
})), b0_3), xs_4).PtrVal.(map[string]gopurs_runtime.Value)["value"]
})
})
})
})
	})
	return scanrWithIndex
}

var mapAccumLWithIndex gopurs_runtime.Value
var once_mapAccumLWithIndex sync.Once
func Get_mapAccumLWithIndex() gopurs_runtime.Value {
	once_mapAccumLWithIndex.Do(func() {
		mapAccumLWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex1_1_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL())
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, i_5), s_7), a_6)
})
})
})), xs_4), s0_3)
})
})
})
})
	})
	return mapAccumLWithIndex
}

var scanlWithIndex gopurs_runtime.Value
var once_scanlWithIndex sync.Once
func Get_scanlWithIndex() gopurs_runtime.Value {
	once_scanlWithIndex.Do(func() {
		scanlWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
mapAccumLWithIndex1_1_0 := gopurs_runtime.Apply(Get_mapAccumLWithIndex(), dictTraversableWithIndex_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(mapAccumLWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_8_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, i_5), b_6), a_7)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": b_prime_8_1, "value": b_prime_8_1})
})
})
})), b0_3), xs_4).PtrVal.(map[string]gopurs_runtime.Value)["value"]
})
})
})
})
	})
	return scanlWithIndex
}

var forWithIndex gopurs_runtime.Value
var once_forWithIndex sync.Once
func Get_forWithIndex() gopurs_runtime.Value {
	once_forWithIndex.Do(func() {
		forWithIndex = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTraversableWithIndex_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictTraversableWithIndex_1.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_0)
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0, a_4), b_3)
})
})
})
})
	})
	return forWithIndex
}


