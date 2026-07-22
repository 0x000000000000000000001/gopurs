package Data_FoldableWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		monoidEndo = func() gopurs_runtime.Value {
semigroupEndo1_0_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
})})
}()
	})
	return monoidEndo
}

var monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		monoidDual = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(Get_monoidEndo().PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
semigroupDual1_1_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_0_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v1_2), v_1)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": Get_monoidEndo().PtrVal.(map[string]gopurs_runtime.Value)["mempty"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_1_1
})})
}()
	})
	return monoidDual
}

var foldrWithIndex gopurs_runtime.Value
var once_foldrWithIndex sync.Once
func Get_foldrWithIndex() gopurs_runtime.Value {
	once_foldrWithIndex.Do(func() {
		foldrWithIndex = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"]
})
	})
	return foldrWithIndex
}

var traverseWithIndex_ gopurs_runtime.Value
var once_traverseWithIndex_ sync.Once
func Get_traverseWithIndex_() gopurs_runtime.Value {
	once_traverseWithIndex_.Do(func() {
		traverseWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
applySecond_1_0 := gopurs_runtime.Apply(pkg_Control_Apply.Get_applySecond(), gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_2.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(f_3, i_4)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(__local_var_5_1, x_6))
})
})), gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit()))
})
})
})
	})
	return traverseWithIndex_
}

var forWithIndex_ gopurs_runtime.Value
var once_forWithIndex_ sync.Once
func Get_forWithIndex_() gopurs_runtime.Value {
	once_forWithIndex_.Do(func() {
		forWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex_1_1_0 := gopurs_runtime.Apply(Get_traverseWithIndex_(), dictApplicative_0)
return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(traverseWithIndex_1_1_0, dictFoldableWithIndex_2)
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1, a_5), b_4)
})
})
})
})
	})
	return forWithIndex_
}

var foldrDefault gopurs_runtime.Value
var once_foldrDefault sync.Once
func Get_foldrDefault() gopurs_runtime.Value {
	once_foldrDefault.Do(func() {
		foldrDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
})
})
	})
	return foldrDefault
}

var foldlWithIndex gopurs_runtime.Value
var once_foldlWithIndex sync.Once
func Get_foldlWithIndex() gopurs_runtime.Value {
	once_foldlWithIndex.Do(func() {
		foldlWithIndex = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"]
})
	})
	return foldlWithIndex
}

var foldlDefault gopurs_runtime.Value
var once_foldlDefault sync.Once
func Get_foldlDefault() gopurs_runtime.Value {
	once_foldlDefault.Do(func() {
		foldlDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
})
})
	})
	return foldlDefault
}

var foldableWithIndexTuple gopurs_runtime.Value
var once_foldableWithIndexTuple sync.Once
func Get_foldableWithIndexTuple() gopurs_runtime.Value {
	once_foldableWithIndexTuple.Do(func() {
		foldableWithIndexTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), z_1)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), z_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
})})
	})
	return foldableWithIndexTuple
}

var foldableWithIndexMultiplicative gopurs_runtime.Value
var once_foldableWithIndexMultiplicative sync.Once
func Get_foldableWithIndexMultiplicative() gopurs_runtime.Value {
	once_foldableWithIndexMultiplicative.Do(func() {
		foldableWithIndexMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v_3), z_2)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit())
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
})})
	})
	return foldableWithIndexMultiplicative
}

var foldableWithIndexMaybe gopurs_runtime.Value
var once_foldableWithIndexMaybe sync.Once
func Get_foldableWithIndexMaybe() gopurs_runtime.Value {
	once_foldableWithIndexMaybe.Do(func() {
		foldableWithIndexMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = v1_2
} else {
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_2)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = v1_2
} else {
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_2, v1_2), v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_4 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = mempty_1_4
} else {
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(__local_var_3_5, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
})})
	})
	return foldableWithIndexMaybe
}

var foldableWithIndexLast gopurs_runtime.Value
var once_foldableWithIndexLast sync.Once
func Get_foldableWithIndexLast() gopurs_runtime.Value {
	once_foldableWithIndexLast.Do(func() {
		foldableWithIndexLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = z_2
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), z_2)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = z_2
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_2, z_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_4 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = mempty_1_4
} else {
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(__local_var_3_5, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableLast()
})})
	})
	return foldableWithIndexLast
}

var foldableWithIndexIdentity gopurs_runtime.Value
var once_foldableWithIndexIdentity sync.Once
func Get_foldableWithIndexIdentity() gopurs_runtime.Value {
	once_foldableWithIndexIdentity.Do(func() {
		foldableWithIndexIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), v_2), z_1)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), z_1), v_2)
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()), v_2)
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
})})
	})
	return foldableWithIndexIdentity
}

var foldableWithIndexFirst gopurs_runtime.Value
var once_foldableWithIndexFirst sync.Once
func Get_foldableWithIndexFirst() gopurs_runtime.Value {
	once_foldableWithIndexFirst.Do(func() {
		foldableWithIndexFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = z_2
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), z_2)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = z_2
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_2, z_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_4 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = mempty_1_4
} else {
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(__local_var_3_5, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableFirst()
})})
	})
	return foldableWithIndexFirst
}

var foldableWithIndexEither gopurs_runtime.Value
var once_foldableWithIndexEither sync.Once
func Get_foldableWithIndexEither() gopurs_runtime.Value {
	once_foldableWithIndexEither.Do(func() {
		foldableWithIndexEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = v1_1
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit()), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_1)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = v1_1
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit()), v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = mempty_1_2
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_2, pkg_Data_Unit.Get_unit()), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableEither()
})})
	})
	return foldableWithIndexEither
}

var foldableWithIndexDual gopurs_runtime.Value
var once_foldableWithIndexDual sync.Once
func Get_foldableWithIndexDual() gopurs_runtime.Value {
	once_foldableWithIndexDual.Do(func() {
		foldableWithIndexDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v_3), z_2)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit())
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
})})
	})
	return foldableWithIndexDual
}

var foldableWithIndexDisj gopurs_runtime.Value
var once_foldableWithIndexDisj sync.Once
func Get_foldableWithIndexDisj() gopurs_runtime.Value {
	once_foldableWithIndexDisj.Do(func() {
		foldableWithIndexDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v_3), z_2)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit())
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDisj()
})})
	})
	return foldableWithIndexDisj
}

var foldableWithIndexConst gopurs_runtime.Value
var once_foldableWithIndexConst sync.Once
func Get_foldableWithIndexConst() gopurs_runtime.Value {
	once_foldableWithIndexConst.Do(func() {
		foldableWithIndexConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConst()
})})
	})
	return foldableWithIndexConst
}

var foldableWithIndexConj gopurs_runtime.Value
var once_foldableWithIndexConj sync.Once
func Get_foldableWithIndexConj() gopurs_runtime.Value {
	once_foldableWithIndexConj.Do(func() {
		foldableWithIndexConj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v_3), z_2)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit())
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConj()
})})
	})
	return foldableWithIndexConj
}

var foldableWithIndexAdditive gopurs_runtime.Value
var once_foldableWithIndexAdditive sync.Once
func Get_foldableWithIndexAdditive() gopurs_runtime.Value {
	once_foldableWithIndexAdditive.Do(func() {
		foldableWithIndexAdditive = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0, v_3), z_2)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit())
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableAdditive()
})})
	})
	return foldableWithIndexAdditive
}

var foldWithIndexM gopurs_runtime.Value
var once_foldWithIndexM sync.Once
func Get_foldWithIndexM() gopurs_runtime.Value {
	once_foldWithIndexM.Do(func() {
		foldWithIndexM = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ma_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_0 := gopurs_runtime.Apply(f_2, i_4)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], ma_5), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_7_0, a_8), b_6)
}))
})
})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], a0_3))
})
})
})
})
	})
	return foldWithIndexM
}

var foldMapWithIndexDefaultR gopurs_runtime.Value
var once_foldMapWithIndexDefaultR sync.Once
func Get_foldMapWithIndexDefaultR() gopurs_runtime.Value {
	once_foldMapWithIndexDefaultR.Do(func() {
		foldMapWithIndexDefaultR = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_2_0 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, i_4), x_5)), acc_6)
})
})
})), mempty_2_0)
})
})
})
	})
	return foldMapWithIndexDefaultR
}

var foldableWithIndexArray gopurs_runtime.Value
var once_foldableWithIndexArray sync.Once
func Get_foldableWithIndexArray() gopurs_runtime.Value {
	once_foldableWithIndexArray.Do(func() {
		foldableWithIndexArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldrArray(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_4_2 := v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, __local_var_3_1), __local_var_4_2), y_5)
})
})), z_1)
__local_var_3_3 := gopurs_runtime.Apply(pkg_Data_FunctorWithIndex.Get_mapWithIndexArray(), pkg_Data_Tuple.Get_Tuple())
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldlArray(), gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})), z_1)
__local_var_3_5 := gopurs_runtime.Apply(pkg_Data_FunctorWithIndex.Get_mapWithIndexArray(), pkg_Data_Tuple.Get_Tuple())
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(__local_var_3_5, x_4))
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_6 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexArray().PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, i_3), x_4)), acc_5)
})
})
})), mempty_1_6)
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
})})
	})
	return foldableWithIndexArray
}

var foldMapWithIndexDefaultL gopurs_runtime.Value
var once_foldMapWithIndexDefaultL sync.Once
func Get_foldMapWithIndexDefaultL() gopurs_runtime.Value {
	once_foldMapWithIndexDefaultL.Do(func() {
		foldMapWithIndexDefaultL = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_2_0 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], acc_5), gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, i_4), x_6))
})
})
})), mempty_2_0)
})
})
})
	})
	return foldMapWithIndexDefaultL
}

var foldMapWithIndex gopurs_runtime.Value
var once_foldMapWithIndex sync.Once
func Get_foldMapWithIndex() gopurs_runtime.Value {
	once_foldMapWithIndex.Do(func() {
		foldMapWithIndex = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"]
})
	})
	return foldMapWithIndex
}

var foldableWithIndexApp gopurs_runtime.Value
var once_foldableWithIndexApp sync.Once
func Get_foldableWithIndexApp() gopurs_runtime.Value {
	once_foldableWithIndexApp.Do(func() {
		foldableWithIndexApp = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{})
foldableApp_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_2), i_3), v_4)
})
})
}), "foldl": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2), i_3), v_4)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], f_3), z_4), v_5)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], f_3), z_4), v_5)
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_3)
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableApp_2_1
})})
})
	})
	return foldableWithIndexApp
}

var foldableWithIndexCompose gopurs_runtime.Value
var once_foldableWithIndexCompose sync.Once
func Get_foldableWithIndexCompose() gopurs_runtime.Value {
	once_foldableWithIndexCompose.Do(func() {
		foldableWithIndexCompose = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{})
foldableCompose1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_3 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_4)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_7_3, a_9), b_8)
})
})), i_5), v_6)
})
})
}), "foldl": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_4)), i_5), v_6)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_5_4 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_4)
foldMap5_6_5 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_4)
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap4_5_4, gopurs_runtime.Apply(foldMap5_6_5, f_7)), v_8)
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_6 := gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_8, "value1": b_9}))
}))
return gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_9_6, a_11), b_10)
})
})
})), i_6), v_7)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_8, "value1": b_9}))
}))
})), i_6), v_7)
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex3_6_7 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_5)
foldMapWithIndex4_7_8 := gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_5)
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex3_6_7, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex4_7_8, gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_8, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_10, "value1": b_11}))
}))
})), v_9)
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCompose1_4_2
})})
})
})
	})
	return foldableWithIndexCompose
}

var foldableWithIndexCoproduct gopurs_runtime.Value
var once_foldableWithIndexCoproduct sync.Once
func Get_foldableWithIndexCoproduct() gopurs_runtime.Value {
	once_foldableWithIndexCoproduct.Do(func() {
		foldableWithIndexCoproduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldableCoproduct_1_0 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableCoproduct(), gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldableCoproduct1_3_1 := gopurs_runtime.Apply(foldableCoproduct_1_0, gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_6}))
})), z_5)
__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_7}))
})), z_5)
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(__local_var_6_2, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(__local_var_7_3, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_6}))
})), z_5)
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_7}))
})), z_5)
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(__local_var_6_5, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(__local_var_7_6, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t7
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex3_5_8 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_4)
foldMapWithIndex4_6_9 := gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_4)
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_10 := gopurs_runtime.Apply(foldMapWithIndex3_5_8, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_8}))
}))
__local_var_9_11 := gopurs_runtime.Apply(foldMapWithIndex4_6_9, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_9}))
}))
return gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t12 = gopurs_runtime.Apply(__local_var_8_10, v2_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t12 = gopurs_runtime.Apply(__local_var_9_11, v2_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t12
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCoproduct1_3_1
})})
})
})
	})
	return foldableWithIndexCoproduct
}

var foldableWithIndexProduct gopurs_runtime.Value
var once_foldableWithIndexProduct sync.Once
func Get_foldableWithIndexProduct() gopurs_runtime.Value {
	once_foldableWithIndexProduct.Do(func() {
		foldableWithIndexProduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldableProduct_1_0 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableProduct(), gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldableProduct1_3_1 := gopurs_runtime.Apply(foldableProduct_1_0, gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_7}))
})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_7}))
})), z_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_7}))
})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_7}))
})), z_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex3_5_2 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_4)
foldMapWithIndex4_6_3 := gopurs_runtime.Apply(dictFoldableWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_4)
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_4.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex3_5_2, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_9}))
})), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex4_6_3, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_9}))
})), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableProduct1_3_1
})})
})
})
	})
	return foldableWithIndexProduct
}

var foldlWithIndexDefault gopurs_runtime.Value
var once_foldlWithIndexDefault sync.Once
func Get_foldlWithIndexDefault() gopurs_runtime.Value {
	once_foldlWithIndexDefault.Do(func() {
		foldlWithIndexDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], Get_monoidDual())
return gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(c_2, i_5)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_6_1, a_8), x_7)
})
})
})), xs_4))), u_3)
})
})
})
})
	})
	return foldlWithIndexDefault
}

var foldrWithIndexDefault gopurs_runtime.Value
var once_foldrWithIndexDefault sync.Once
func Get_foldrWithIndexDefault() gopurs_runtime.Value {
	once_foldrWithIndexDefault.Do(func() {
		foldrWithIndexDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], Get_monoidEndo())
return gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_2, i_5)
})), xs_4)), u_3)
})
})
})
})
	})
	return foldrWithIndexDefault
}

var surroundMapWithIndex gopurs_runtime.Value
var once_surroundMapWithIndex sync.Once
func Get_surroundMapWithIndex() gopurs_runtime.Value {
	once_surroundMapWithIndex.Do(func() {
		surroundMapWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], Get_monoidEndo())
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], d_3), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(t_4, i_6), a_7)), m_8))
})
})
})), f_5)), d_3)
})
})
})
})
})
	})
	return surroundMapWithIndex
}

var foldMapDefault gopurs_runtime.Value
var once_foldMapDefault sync.Once
func Get_foldMapDefault() gopurs_runtime.Value {
	once_foldMapDefault.Do(func() {
		foldMapDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex2_2_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return f_3
}))
})
})
})
	})
	return foldMapDefault
}

var findWithIndex gopurs_runtime.Value
var once_findWithIndex sync.Once
func Get_findWithIndex() gopurs_runtime.Value {
	once_findWithIndex.Do(func() {
		findWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(p_1, v_2), v2_4).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"index": v_2, "value": v2_4})})
} else {
__t0 = v1_3
}
return __t0
})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return findWithIndex
}

var findMapWithIndex gopurs_runtime.Value
var once_findMapWithIndex sync.Once
func Get_findMapWithIndex() gopurs_runtime.Value {
	once_findMapWithIndex.Do(func() {
		findMapWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, v_2), v2_4)
} else {
__t0 = v1_3
}
return __t0
})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return findMapWithIndex
}

var anyWithIndex gopurs_runtime.Value
var once_anyWithIndex sync.Once
func Get_anyWithIndex() gopurs_runtime.Value {
	once_anyWithIndex.Do(func() {
		anyWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_2), v1_3)
})
})})
foldMapWithIndex2_2_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["ff"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_1
})}))
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(t_3, i_4)
}))
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
})
})
	})
	return anyWithIndex
}

var allWithIndex gopurs_runtime.Value
var once_allWithIndex sync.Once
func Get_allWithIndex() gopurs_runtime.Value {
	once_allWithIndex.Do(func() {
		allWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupConj1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["conj"], v_2), v1_3)
})
})})
foldMapWithIndex2_2_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["tt"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_1
})}))
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(t_3, i_4)
}))
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
})
})
	})
	return allWithIndex
}


