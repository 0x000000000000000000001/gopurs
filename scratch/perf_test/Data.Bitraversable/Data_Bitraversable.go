package Data_Bitraversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor_Clown "gopurs/output/Data.Functor.Clown"
	pkg_Data_Functor_Flip "gopurs/output/Data.Functor.Flip"
	pkg_Data_Functor_Joker "gopurs/output/Data.Functor.Joker"
	pkg_Data_Functor_Product2 "gopurs/output/Data.Functor.Product2"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var bitraverse gopurs_runtime.Value
var once_bitraverse sync.Once
func Get_bitraverse() gopurs_runtime.Value {
	once_bitraverse.Do(func() {
		bitraverse = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"]
})
	})
	return bitraverse
}

var lfor gopurs_runtime.Value
var once_lfor sync.Once
func Get_lfor() gopurs_runtime.Value {
	once_lfor.Do(func() {
		lfor = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_1)
pure_3_1 := dictApplicative_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"]
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse2_2_0, f_5), pure_3_1), t_4)
})
})
})
})
	})
	return lfor
}

var ltraverse gopurs_runtime.Value
var once_ltraverse sync.Once
func Get_ltraverse() gopurs_runtime.Value {
	once_ltraverse.Do(func() {
		ltraverse = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_1)
pure_3_1 := dictApplicative_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"]
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse2_2_0, f_4), pure_3_1)
})
})
})
	})
	return ltraverse
}

var rfor gopurs_runtime.Value
var once_rfor sync.Once
func Get_rfor() gopurs_runtime.Value {
	once_rfor.Do(func() {
		rfor = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_1)
pure_3_1 := dictApplicative_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"]
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse2_2_0, pure_3_1), f_5), t_4)
})
})
})
})
	})
	return rfor
}

var rtraverse gopurs_runtime.Value
var once_rtraverse sync.Once
func Get_rtraverse() gopurs_runtime.Value {
	once_rtraverse.Do(func() {
		rtraverse = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_1), dictApplicative_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"])
})
})
	})
	return rtraverse
}

var bitraversableTuple gopurs_runtime.Value
var once_bitraversableTuple sync.Once
func Get_bitraversableTuple() gopurs_runtime.Value {
	once_bitraversableTuple.Do(func() {
		bitraversableTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_Tuple()), gopurs_runtime.Apply(f_2, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(g_3, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_1 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_Tuple()), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableTuple()
})})
	})
	return bitraversableTuple
}

var bitraversableJoker gopurs_runtime.Value
var once_bitraversableJoker sync.Once
func Get_bitraversableJoker() gopurs_runtime.Value {
	once_bitraversableJoker.Do(func() {
		bitraversableJoker = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
bifunctorJoker_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_3), v1_4)
})
})
})})
__local_var_3_2 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable1"], gopurs_runtime.Value{})
bifoldableJoker_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], r_5), u_6), v1_7)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], r_5), u_6), v1_7)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_4)
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_5_4, r_7), v1_8)
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_5 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Joker.Get_Joker()), gopurs_runtime.Apply(gopurs_runtime.Apply(traverse1_6_5, r_8), v1_9))
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_6 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["sequence"], dictApplicative_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Joker.Get_Joker()), gopurs_runtime.Apply(sequence1_6_6, v_7))
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker_2_1
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableJoker_4_3
})})
})
	})
	return bitraversableJoker
}

var bitraversableEither gopurs_runtime.Value
var once_bitraversableEither sync.Once
func Get_bitraversableEither() gopurs_runtime.Value {
	once_bitraversableEither.Do(func() {
		bitraversableEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Left()), gopurs_runtime.Apply(v_2, v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Right()), gopurs_runtime.Apply(v1_3, v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Left()), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Right()), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorEither()
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableEither()
})})
	})
	return bitraversableEither
}

var bitraversableConst gopurs_runtime.Value
var once_bitraversableConst sync.Once
func Get_bitraversableConst() gopurs_runtime.Value {
	once_bitraversableConst.Do(func() {
		bitraversableConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Const.Get_Const()), gopurs_runtime.Apply(f_1, v1_3))
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Const.Get_Const()), v_1)
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorConst()
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableConst()
})})
	})
	return bitraversableConst
}

var bitraversableClown gopurs_runtime.Value
var once_bitraversableClown sync.Once
func Get_bitraversableClown() gopurs_runtime.Value {
	once_bitraversableClown.Do(func() {
		bitraversableClown = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
bifunctorClown_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v1_4)
})
})
})})
__local_var_3_2 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable1"], gopurs_runtime.Value{})
bifoldableClown_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], l_4), u_6), v1_7)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], l_4), u_6), v1_7)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_4)
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_5_4, l_6), v1_8)
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_5 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_5)
return gopurs_runtime.Func(func(l_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Clown.Get_Clown()), gopurs_runtime.Apply(gopurs_runtime.Apply(traverse1_6_5, l_7), v1_9))
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_6 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["sequence"], dictApplicative_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Clown.Get_Clown()), gopurs_runtime.Apply(sequence1_6_6, v_7))
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown_2_1
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableClown_4_3
})})
})
	})
	return bitraversableClown
}

var bisequenceDefault gopurs_runtime.Value
var once_bisequenceDefault sync.Once
func Get_bisequenceDefault() gopurs_runtime.Value {
	once_bisequenceDefault.Do(func() {
		bisequenceDefault = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return bisequenceDefault
}

var bisequence gopurs_runtime.Value
var once_bisequence sync.Once
func Get_bisequence() gopurs_runtime.Value {
	once_bisequence.Do(func() {
		bisequence = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bisequence"]
})
	})
	return bisequence
}

var bitraversableFlip gopurs_runtime.Value
var once_bitraversableFlip sync.Once
func Get_bitraversableFlip() gopurs_runtime.Value {
	once_bitraversableFlip.Do(func() {
		bitraversableFlip = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorFlip_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], g_3), f_2), v_4)
})
})
})})
__local_var_3_2 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifoldable1"], gopurs_runtime.Value{})
bifoldableFlip_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bifoldr"], l_5), r_4), u_6), v_7)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bifoldl"], l_5), r_4), u_6), v_7)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap2_5_4 := gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], dictMonoid_4)
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap2_5_4, l_7), r_6), v_8)
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_6_5 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_5)
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Flip.Get_Flip()), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse2_6_5, l_8), r_7), v_9))
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
bisequence2_6_6 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bisequence"], dictApplicative_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Flip.Get_Flip()), gopurs_runtime.Apply(bisequence2_6_6, v_7))
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip_2_1
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableFlip_4_3
})})
})
	})
	return bitraversableFlip
}

var bitraversableProduct2 gopurs_runtime.Value
var once_bitraversableProduct2 sync.Once
func Get_bitraversableProduct2() gopurs_runtime.Value {
	once_bitraversableProduct2.Do(func() {
		bitraversableProduct2 = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifoldableProduct2_2_1 := gopurs_runtime.Apply(pkg_Data_Bifoldable.Get_bifoldableProduct2(), gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifoldable1"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictBitraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictBitraversable1_3.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorProduct21_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_5), g_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_5), g_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})})
bifoldableProduct21_6_4 := gopurs_runtime.Apply(bifoldableProduct2_2_1, gopurs_runtime.Apply(dictBitraversable1_3.PtrVal.(map[string]gopurs_runtime.Value)["Bifoldable1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_8_5 := gopurs_runtime.Apply(dictApplicative_7.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
bitraverse3_9_6 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_7)
bitraverse4_10_7 := gopurs_runtime.Apply(dictBitraversable1_3.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_7)
return gopurs_runtime.Func(func(l_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_8_5.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_8_5.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Product2.Get_Product2()), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse3_9_6, l_11), r_12), v_13.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse4_10_7, l_11), r_12), v_13.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_8_8 := gopurs_runtime.Apply(dictApplicative_7.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
bisequence3_9_9 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bisequence"], dictApplicative_7)
bisequence4_10_10 := gopurs_runtime.Apply(dictBitraversable1_3.PtrVal.(map[string]gopurs_runtime.Value)["bisequence"], dictApplicative_7)
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_8_8.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_8_8.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Functor_Product2.Get_Product2()), gopurs_runtime.Apply(bisequence3_9_9, v_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(bisequence4_10_10, v_11.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct21_5_3
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableProduct21_6_4
})})
})
})
	})
	return bitraversableProduct2
}

var bitraverseDefault gopurs_runtime.Value
var once_bitraverseDefault sync.Once
func Get_bitraverseDefault() gopurs_runtime.Value {
	once_bitraverseDefault.Do(func() {
		bitraverseDefault = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bisequence2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bisequence"], dictApplicative_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(bisequence2_2_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_3), g_4), t_5))
})
})
})
})
})
	})
	return bitraverseDefault
}

var bifor gopurs_runtime.Value
var once_bifor sync.Once
func Get_bifor() gopurs_runtime.Value {
	once_bifor.Do(func() {
		bifor = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_1)
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse2_2_0, f_4), g_5), t_3)
})
})
})
})
})
	})
	return bifor
}


