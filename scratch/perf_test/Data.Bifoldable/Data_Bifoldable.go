package Data_Bifoldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Category "gopurs/output/Control.Category"
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

var bifoldr gopurs_runtime.Value
var once_bifoldr sync.Once
func Get_bifoldr() gopurs_runtime.Value {
	once_bifoldr.Do(func() {
		bifoldr = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldr"]
})
	})
	return bifoldr
}

var bitraverse_ gopurs_runtime.Value
var once_bitraverse_ sync.Once
func Get_bitraverse_() gopurs_runtime.Value {
	once_bitraverse_.Do(func() {
		bitraverse_ = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
applySecond_2_0 := gopurs_runtime.Apply(pkg_Control_Apply.Get_applySecond(), gopurs_runtime.Apply(dictApplicative_1.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldr"], gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_2_0, gopurs_runtime.Apply(f_3, x_5))
})), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_2_0, gopurs_runtime.Apply(g_4, x_5))
})), gopurs_runtime.Apply(dictApplicative_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit()))
})
})
})
})
	})
	return bitraverse_
}

var bifor_ gopurs_runtime.Value
var once_bifor_ sync.Once
func Get_bifor_() gopurs_runtime.Value {
	once_bifor_.Do(func() {
		bifor_ = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse_2_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bitraverse_(), dictBifoldable_0), dictApplicative_1)
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse_2_2_0, f_4), g_5), t_3)
})
})
})
})
})
	})
	return bifor_
}

var bisequence_ gopurs_runtime.Value
var once_bisequence_ sync.Once
func Get_bisequence_() gopurs_runtime.Value {
	once_bisequence_.Do(func() {
		bisequence_ = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bitraverse_(), dictBifoldable_0), dictApplicative_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return bisequence_
}

var bifoldl gopurs_runtime.Value
var once_bifoldl sync.Once
func Get_bifoldl() gopurs_runtime.Value {
	once_bifoldl.Do(func() {
		bifoldl = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldl"]
})
	})
	return bifoldl
}

var bifoldableTuple gopurs_runtime.Value
var once_bifoldableTuple sync.Once
func Get_bifoldableTuple() gopurs_runtime.Value {
	once_bifoldableTuple.Do(func() {
		bifoldableTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(g_2, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
}), "bifoldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), z_2))
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
})
})})
	})
	return bifoldableTuple
}

var bifoldableJoker gopurs_runtime.Value
var once_bifoldableJoker sync.Once
func Get_bifoldableJoker() gopurs_runtime.Value {
	once_bifoldableJoker.Do(func() {
		bifoldableJoker = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], r_2), u_3), v1_4)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], r_2), u_3), v1_4)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_1)
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_2_0, r_4), v1_5)
})
})
})
})})
})
	})
	return bifoldableJoker
}

var bifoldableEither gopurs_runtime.Value
var once_bifoldableEither sync.Once
func Get_bifoldableEither() gopurs_runtime.Value {
	once_bifoldableEither.Do(func() {
		bifoldableEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v2_2)
} else {
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v2_2)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v2_2), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, v2_2), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(v_1, v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(v1_2, v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
})
})})
	})
	return bifoldableEither
}

var bifoldableConst gopurs_runtime.Value
var once_bifoldableConst sync.Once
func Get_bifoldableConst() gopurs_runtime.Value {
	once_bifoldableConst.Do(func() {
		bifoldableConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v1_3), z_2)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_2), v1_3)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v1_3)
})
})
})
})})
	})
	return bifoldableConst
}

var bifoldableClown gopurs_runtime.Value
var once_bifoldableClown sync.Once
func Get_bifoldableClown() gopurs_runtime.Value {
	once_bifoldableClown.Do(func() {
		bifoldableClown = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], l_1), u_3), v1_4)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], l_1), u_3), v1_4)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_1)
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_2_0, l_3), v1_5)
})
})
})
})})
})
	})
	return bifoldableClown
}

var bifoldMapDefaultR gopurs_runtime.Value
var once_bifoldMapDefaultR sync.Once
func Get_bifoldMapDefaultR() gopurs_runtime.Value {
	once_bifoldMapDefaultR.Do(func() {
		bifoldMapDefaultR = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_3_1 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldr"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_4, x_6))
})), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(g_5, x_6))
})), mempty_3_1)
})
})
})
})
	})
	return bifoldMapDefaultR
}

var bifoldMapDefaultL gopurs_runtime.Value
var once_bifoldMapDefaultL sync.Once
func Get_bifoldMapDefaultL() gopurs_runtime.Value {
	once_bifoldMapDefaultL.Do(func() {
		bifoldMapDefaultL = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_3_1 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldl"], gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], m_6), gopurs_runtime.Apply(f_4, a_7))
})
})), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], m_6), gopurs_runtime.Apply(g_5, b_7))
})
})), mempty_3_1)
})
})
})
})
	})
	return bifoldMapDefaultL
}

var bifoldMap gopurs_runtime.Value
var once_bifoldMap sync.Once
func Get_bifoldMap() gopurs_runtime.Value {
	once_bifoldMap.Do(func() {
		bifoldMap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"]
})
	})
	return bifoldMap
}

var bifoldableFlip gopurs_runtime.Value
var once_bifoldableFlip sync.Once
func Get_bifoldableFlip() gopurs_runtime.Value {
	once_bifoldableFlip.Do(func() {
		bifoldableFlip = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldr"], l_2), r_1), u_3), v_4)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldl"], l_2), r_1), u_3), v_4)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap2_2_0 := gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], dictMonoid_1)
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap2_2_0, l_4), r_3), v_5)
})
})
})
})})
})
	})
	return bifoldableFlip
}

var bifoldlDefault gopurs_runtime.Value
var once_bifoldlDefault sync.Once
func Get_bifoldlDefault() gopurs_runtime.Value {
	once_bifoldlDefault.Do(func() {
		bifoldlDefault = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap1_1_0 := gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], Get_monoidDual())
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap1_1_0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, a_7), x_6)
})
})), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(g_3, a_7), x_6)
})
})), p_5))), z_4)
})
})
})
})
})
	})
	return bifoldlDefault
}

var bifoldrDefault gopurs_runtime.Value
var once_bifoldrDefault sync.Once
func Get_bifoldrDefault() gopurs_runtime.Value {
	once_bifoldrDefault.Do(func() {
		bifoldrDefault = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap1_1_0 := gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], Get_monoidEndo())
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap1_1_0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, x_6)
})), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_3, x_6)
})), p_5)), z_4)
})
})
})
})
})
	})
	return bifoldrDefault
}

var bifoldableProduct2 gopurs_runtime.Value
var once_bifoldableProduct2 sync.Once
func Get_bifoldableProduct2() gopurs_runtime.Value {
	once_bifoldableProduct2.Do(func() {
		bifoldableProduct2 = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBifoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldr": gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bifoldrDefault(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bifoldableProduct2(), dictBifoldable_0), dictBifoldable1_1)), l_2), r_3), u_4), m_5)
})
})
})
}), "bifoldl": gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bifoldlDefault(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bifoldableProduct2(), dictBifoldable_0), dictBifoldable1_1)), l_2), r_3), u_4), m_5)
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap3_3_0 := gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], dictMonoid_2)
bifoldMap4_4_1 := gopurs_runtime.Apply(dictBifoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], dictMonoid_2)
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap3_3_0, l_5), r_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap4_4_1, l_5), r_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
})})
}
}()
})
})
	})
	return bifoldableProduct2
}

var bifold gopurs_runtime.Value
var once_bifold sync.Once
func Get_bifold() gopurs_runtime.Value {
	once_bifold.Do(func() {
		bifold = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], dictMonoid_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return bifold
}

var biany gopurs_runtime.Value
var once_biany sync.Once
func Get_biany() gopurs_runtime.Value {
	once_biany.Do(func() {
		biany = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBooleanAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(dictBooleanAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebra0"], gopurs_runtime.Value{})
semigroupDisj1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_3), v1_4)
})
})})
bifoldMap2_2_0 := gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["ff"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_3_2
})}))
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(q_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap2_2_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
})), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
})
})
})
	})
	return biany
}

var biall gopurs_runtime.Value
var once_biall sync.Once
func Get_biall() gopurs_runtime.Value {
	once_biall.Do(func() {
		biall = gopurs_runtime.Func(func(dictBifoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBooleanAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(dictBooleanAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebra0"], gopurs_runtime.Value{})
semigroupConj1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["conj"], v_3), v1_4)
})
})})
bifoldMap2_2_0 := gopurs_runtime.Apply(dictBifoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["bifoldMap"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["tt"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_3_2
})}))
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(q_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(bifoldMap2_2_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
})), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
})
})
})
	})
	return biall
}


