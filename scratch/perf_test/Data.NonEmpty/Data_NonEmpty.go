package Data_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var NonEmpty gopurs_runtime.Value
var once_NonEmpty sync.Once
func Get_NonEmpty() gopurs_runtime.Value {
	once_NonEmpty.Do(func() {
		NonEmpty = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": value0, "value1": value1})
})
})
	})
	return NonEmpty
}

var unfoldable1NonEmpty gopurs_runtime.Value
var once_unfoldable1NonEmpty sync.Once
func Get_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_unfoldable1NonEmpty.Do(func() {
		unfoldable1NonEmpty = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr1": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(f_1, b_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(f_1, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})), __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
	})
	return unfoldable1NonEmpty
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return tail
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
empty_1_0 := dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": a_2, "value1": empty_1_0})
})
})
	})
	return singleton
}

var showNonEmpty gopurs_runtime.Value
var once_showNonEmpty sync.Once
func Get_showNonEmpty() gopurs_runtime.Value {
	once_showNonEmpty.Do(func() {
		showNonEmpty = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(NonEmpty ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
})})
})
})
	})
	return showNonEmpty
}

var semigroupNonEmpty gopurs_runtime.Value
var once_semigroupNonEmpty sync.Once
func Get_semigroupNonEmpty() gopurs_runtime.Value {
	once_semigroupNonEmpty.Do(func() {
		semigroupNonEmpty = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))})
})
})})
})
})
	})
	return semigroupNonEmpty
}

var oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		oneOf = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
	})
	return oneOf
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
	})
	return head
}

var functorNonEmpty gopurs_runtime.Value
var once_functorNonEmpty sync.Once
func Get_functorNonEmpty() gopurs_runtime.Value {
	once_functorNonEmpty.Do(func() {
		functorNonEmpty = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(f_1, m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1), m_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
	})
	return functorNonEmpty
}

var functorWithIndex gopurs_runtime.Value
var once_functorWithIndex sync.Once
func Get_functorWithIndex() gopurs_runtime.Value {
	once_functorWithIndex.Do(func() {
		functorWithIndex = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorNonEmpty1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(f_2, m_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), m_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_5}))
})), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_1
})})
})
	})
	return functorWithIndex
}

var fromNonEmpty gopurs_runtime.Value
var once_fromNonEmpty sync.Once
func Get_fromNonEmpty() gopurs_runtime.Value {
	once_fromNonEmpty.Do(func() {
		fromNonEmpty = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
	})
	return fromNonEmpty
}

var foldableNonEmpty gopurs_runtime.Value
var once_foldableNonEmpty sync.Once
func Get_foldableNonEmpty() gopurs_runtime.Value {
	once_foldableNonEmpty.Do(func() {
		foldableNonEmpty = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap": gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_3, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_2_0, f_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "foldl": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_1), gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, b_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_1), b_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})})
})
	})
	return foldableNonEmpty
}

var foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		foldableWithIndexNonEmpty = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{})
foldableNonEmpty1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap": gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_3_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_4, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_3_2, f_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "foldl": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2), gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, b_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_2), b_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_4_3 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_3)
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex1_4_3, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_7}))
})), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_6}))
})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), b_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldrWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_6}))
})), b_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_2_1
})})
})
	})
	return foldableWithIndexNonEmpty
}

var traversableNonEmpty gopurs_runtime.Value
var once_traversableNonEmpty sync.Once
func Get_traversableNonEmpty() gopurs_runtime.Value {
	once_traversableNonEmpty.Do(func() {
		traversableNonEmpty = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorNonEmpty1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(f_2, m_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), m_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
__local_var_3_2 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable1"], gopurs_runtime.Value{})
foldableNonEmpty1_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_4)
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_4.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_6, v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_5_4, f_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "foldl": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_4), gopurs_runtime.Apply(gopurs_runtime.Apply(f_4, b_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_4, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_4), b_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sequence": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_6_5 := gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
sequence1_7_6 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["sequence"], dictApplicative_5)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_6_5.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_6_5.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_NonEmpty()), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(sequence1_7_6, v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
}), "traverse": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_6_7 := gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
traverse1_7_8 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_5)
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_6_7.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_6_7.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_NonEmpty()), gopurs_runtime.Apply(f_8, v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(traverse1_7_8, f_8), v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_1
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_4_3
})})
})
	})
	return traversableNonEmpty
}

var traversableWithIndexNonEmpty gopurs_runtime.Value
var once_traversableWithIndexNonEmpty sync.Once
func Get_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_traversableWithIndexNonEmpty.Do(func() {
		traversableWithIndexNonEmpty = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorNonEmpty1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(f_3, m_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3), m_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
functorWithIndex1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_6}))
})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_3_3
})})
foldableWithIndexNonEmpty1_4_4 := gopurs_runtime.Apply(Get_foldableWithIndexNonEmpty(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableNonEmpty1_5_5 := gopurs_runtime.Apply(Get_traversableNonEmpty(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_6 := gopurs_runtime.Apply(dictApplicative_6.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
traverseWithIndex1_8_7 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_6)
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_7_6.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_7_6.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_NonEmpty()), gopurs_runtime.Apply(gopurs_runtime.Apply(f_9, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), v_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex1_8_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_11}))
})), v_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndex1_3_2
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexNonEmpty1_4_4
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableNonEmpty1_5_5
})})
})
	})
	return traversableWithIndexNonEmpty
}

var foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		foldable1NonEmpty = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldableNonEmpty1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap": gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_1 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_3, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_2_1, f_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "foldl": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_1), gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, b_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_1), b_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap1": gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], s_5), gopurs_runtime.Apply(f_3, a1_6))
})
})), gopurs_runtime.Apply(f_3, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr1": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(f_2, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(f_2, a1_5)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = a1_5
} else {
if (gopurs_runtime.Bool(x_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_6_4, x_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __t5})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(__local_var_4_2, __local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
}), "foldl1": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_1_0
})})
})
	})
	return foldable1NonEmpty
}

var foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		foldl1 = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_foldable1NonEmpty(), dictFoldable_0).PtrVal.(map[string]gopurs_runtime.Value)["foldl1"]
})
	})
	return foldl1
}

var eqNonEmpty gopurs_runtime.Value
var once_eqNonEmpty sync.Once
func Get_eqNonEmpty() gopurs_runtime.Value {
	once_eqNonEmpty.Do(func() {
		eqNonEmpty = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_2_0 := gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_1)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(eq11_2_0, x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})})
})
})
	})
	return eqNonEmpty
}

var ordNonEmpty gopurs_runtime.Value
var once_ordNonEmpty sync.Once
func Get_ordNonEmpty() gopurs_runtime.Value {
	once_ordNonEmpty.Do(func() {
		ordNonEmpty = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqNonEmpty1_1_0 := gopurs_runtime.Apply(Get_eqNonEmpty(), gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_2)
eqNonEmpty2_4_2 := gopurs_runtime.Apply(eqNonEmpty1_1_0, gopurs_runtime.Apply(dictOrd_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(compare11_3_1, x_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
}
return __t4
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqNonEmpty2_4_2
})})
})
})
	})
	return ordNonEmpty
}

var eq1NonEmpty gopurs_runtime.Value
var once_eq1NonEmpty sync.Once
func Get_eq1NonEmpty() gopurs_runtime.Value {
	once_eq1NonEmpty.Do(func() {
		eq1NonEmpty = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqNonEmpty(), dictEq1_0), dictEq_1).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
})
	})
	return eq1NonEmpty
}

var ord1NonEmpty gopurs_runtime.Value
var once_ord1NonEmpty sync.Once
func Get_ord1NonEmpty() gopurs_runtime.Value {
	once_ord1NonEmpty.Do(func() {
		ord1NonEmpty = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordNonEmpty1_1_0 := gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd1_0)
__local_var_2_1 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
eq1NonEmpty1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqNonEmpty(), __local_var_2_1), dictEq_3).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ordNonEmpty1_1_0, dictOrd_4).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1NonEmpty1_3_2
})})
})
	})
	return ord1NonEmpty
}


