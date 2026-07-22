package Data_Functor_Product2

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Product2 gopurs_runtime.Value
var once_Product2 sync.Once
func Get_Product2() gopurs_runtime.Value {
	once_Product2.Do(func() {
		Product2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": value0, "value1": value1})
})
})
	})
	return Product2
}

var showProduct2 gopurs_runtime.Value
var once_showProduct2 sync.Once
func Get_showProduct2() gopurs_runtime.Value {
	once_showProduct2.Do(func() {
		showProduct2 = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Product2 ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
})})
})
})
	})
	return showProduct2
}

var profunctorProduct2 gopurs_runtime.Value
var once_profunctorProduct2 sync.Once
func Get_profunctorProduct2() gopurs_runtime.Value {
	once_profunctorProduct2.Do(func() {
		profunctorProduct2 = gopurs_runtime.Func(func(dictProfunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictProfunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], f_2), g_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor1_1.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], f_2), g_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})})
})
})
	})
	return profunctorProduct2
}

var functorProduct2 gopurs_runtime.Value
var once_functorProduct2 sync.Once
func Get_functorProduct2() gopurs_runtime.Value {
	once_functorProduct2.Do(func() {
		functorProduct2 = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor1_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
})
	})
	return functorProduct2
}

var eqProduct2 gopurs_runtime.Value
var once_eqProduct2 sync.Once
func Get_eqProduct2() gopurs_runtime.Value {
	once_eqProduct2.Do(func() {
		eqProduct2 = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})})
})
})
	})
	return eqProduct2
}

var ordProduct2 gopurs_runtime.Value
var once_ordProduct2 sync.Once
func Get_ordProduct2() gopurs_runtime.Value {
	once_ordProduct2.Do(func() {
		ordProduct2 = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqProduct21_1_0 := gopurs_runtime.Apply(Get_eqProduct2(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqProduct22_3_1 := gopurs_runtime.Apply(eqProduct21_1_0, gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
if (gopurs_runtime.Bool(v_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
}
return __t3
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct22_3_1
})})
})
})
	})
	return ordProduct2
}

var bifunctorProduct2 gopurs_runtime.Value
var once_bifunctorProduct2 sync.Once
func Get_bifunctorProduct2() gopurs_runtime.Value {
	once_bifunctorProduct2.Do(func() {
		bifunctorProduct2 = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBifunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_2), g_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifunctor1_1.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_2), g_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})})
})
})
	})
	return bifunctorProduct2
}

var biapplyProduct2 gopurs_runtime.Value
var once_biapplyProduct2 sync.Once
func Get_biapplyProduct2() gopurs_runtime.Value {
	once_biapplyProduct2.Do(func() {
		biapplyProduct2 = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictBiapply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictBiapply1_2.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorProduct22_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_4), g_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_4), g_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"biapply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply1_2.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_4_2
})})
})
})
	})
	return biapplyProduct2
}

var biapplicativeProduct2 gopurs_runtime.Value
var once_biapplicativeProduct2 sync.Once
func Get_biapplicativeProduct2() gopurs_runtime.Value {
	once_biapplicativeProduct2.Do(func() {
		biapplicativeProduct2 = gopurs_runtime.Func(func(dictBiapplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBiapplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Biapply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictBiapplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictBiapplicative1_3.PtrVal.(map[string]gopurs_runtime.Value)["Biapply0"], gopurs_runtime.Value{})
__local_var_5_3 := gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorProduct22_6_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_6), g_7), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_6), g_7), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})})
biapplyProduct22_6_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"biapply": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct22_6_5
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bipure": gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product2"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["bipure"], a_7), b_8), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapplicative1_3.PtrVal.(map[string]gopurs_runtime.Value)["bipure"], a_7), b_8)})
})
}), "Biapply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyProduct22_6_4
})})
})
})
	})
	return biapplicativeProduct2
}


