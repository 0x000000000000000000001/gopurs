package Data_Functor_Product

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var Product gopurs_runtime.Value
var once_Product sync.Once
func Get_Product() gopurs_runtime.Value {
	once_Product.Do(func() {
		Product = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Product
}

var showProduct gopurs_runtime.Value
var once_showProduct sync.Once
func Get_showProduct() gopurs_runtime.Value {
	once_showProduct.Do(func() {
		showProduct = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(product ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
})})
})
})
	})
	return showProduct
}

var product gopurs_runtime.Value
var once_product sync.Once
func Get_product() gopurs_runtime.Value {
	once_product.Do(func() {
		product = gopurs_runtime.Func(func(fa_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ga_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": fa_0, "value1": ga_1})
})
})
	})
	return product
}

var newtypeProduct gopurs_runtime.Value
var once_newtypeProduct sync.Once
func Get_newtypeProduct() gopurs_runtime.Value {
	once_newtypeProduct.Do(func() {
		newtypeProduct = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeProduct
}

var functorProduct gopurs_runtime.Value
var once_functorProduct sync.Once
func Get_functorProduct() gopurs_runtime.Value {
	once_functorProduct.Do(func() {
		functorProduct = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor1_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
})
	})
	return functorProduct
}

var eq1Product gopurs_runtime.Value
var once_eq1Product sync.Once
func Get_eq1Product() gopurs_runtime.Value {
	once_eq1Product.Do(func() {
		eq1Product = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)
eq13_4_1 := gopurs_runtime.Apply(dictEq11_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(eq12_3_0, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(eq13_4_1, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})})
})
})
	})
	return eq1Product
}

var eqProduct gopurs_runtime.Value
var once_eqProduct sync.Once
func Get_eqProduct() gopurs_runtime.Value {
	once_eqProduct.Do(func() {
		eqProduct = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eq1Product(), dictEq1_0), dictEq11_1).PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)})
})
})
})
	})
	return eqProduct
}

var ord1Product gopurs_runtime.Value
var once_ord1Product sync.Once
func Get_ord1Product() gopurs_runtime.Value {
	once_ord1Product.Do(func() {
		ord1Product = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
eq1Product1_1_0 := gopurs_runtime.Apply(Get_eq1Product(), gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq1Product2_3_1 := gopurs_runtime.Apply(eq1Product1_1_0, gopurs_runtime.Apply(dictOrd11_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare12_5_2 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_4)
compare13_6_3 := gopurs_runtime.Apply(dictOrd11_2.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_4)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
v2_9_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(compare12_5_2, v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_9_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(compare13_6_3, v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t5 = v2_9_4
}
return __t5
})
})
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Product2_3_1
})})
})
})
	})
	return ord1Product
}

var ordProduct gopurs_runtime.Value
var once_ordProduct sync.Once
func Get_ordProduct() gopurs_runtime.Value {
	once_ordProduct.Do(func() {
		ordProduct = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
ord1Product1_1_0 := gopurs_runtime.Apply(Get_ord1Product(), dictOrd1_0)
eqProduct1_2_1 := gopurs_runtime.Apply(Get_eqProduct(), gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqProduct2_4_2 := gopurs_runtime.Apply(eqProduct1_2_1, gopurs_runtime.Apply(dictOrd11_3.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
eqProduct3_6_3 := gopurs_runtime.Apply(eqProduct2_4_2, gopurs_runtime.Apply(dictOrd_5.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(gopurs_runtime.Apply(ord1Product1_1_0, dictOrd11_3).PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_5), "Eq0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqProduct3_6_3
})})
})
})
})
	})
	return ordProduct
}

var bihoistProduct gopurs_runtime.Value
var once_bihoistProduct sync.Once
func Get_bihoistProduct() gopurs_runtime.Value {
	once_bihoistProduct.Do(func() {
		bihoistProduct = gopurs_runtime.Func(func(natF_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(natG_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(natF_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(natG_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})
	})
	return bihoistProduct
}

var applyProduct gopurs_runtime.Value
var once_applyProduct sync.Once
func Get_applyProduct() gopurs_runtime.Value {
	once_applyProduct.Do(func() {
		applyProduct = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictApply1_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorProduct2_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply1_2.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_4_2
})})
})
})
	})
	return applyProduct
}

var bindProduct gopurs_runtime.Value
var once_bindProduct sync.Once
func Get_bindProduct() gopurs_runtime.Value {
	once_bindProduct.Do(func() {
		bindProduct = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictBind1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictBind1_3.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_5_3 := gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorProduct2_6_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
applyProduct2_6_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_6_5
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(f_8, x_9)).PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind1_3.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(f_8, x_9)).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
}))})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_6_4
})})
})
})
	})
	return bindProduct
}

var applicativeProduct gopurs_runtime.Value
var once_applicativeProduct sync.Once
func Get_applicativeProduct() gopurs_runtime.Value {
	once_applicativeProduct.Do(func() {
		applicativeProduct = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictApplicative1_3.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_5_3 := gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorProduct2_6_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
applyProduct2_6_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct2_6_5
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], a_7), "value1": gopurs_runtime.Apply(dictApplicative1_3.PtrVal.(map[string]gopurs_runtime.Value)["pure"], a_7)})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyProduct2_6_4
})})
})
})
	})
	return applicativeProduct
}

var monadProduct gopurs_runtime.Value
var once_monadProduct sync.Once
func Get_monadProduct() gopurs_runtime.Value {
	once_monadProduct.Do(func() {
		monadProduct = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeProduct1_1_0 := gopurs_runtime.Apply(Get_applicativeProduct(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
bindProduct1_2_1 := gopurs_runtime.Apply(Get_bindProduct(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonad1_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeProduct2_4_2 := gopurs_runtime.Apply(applicativeProduct1_1_0, gopurs_runtime.Apply(dictMonad1_3.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
bindProduct2_5_3 := gopurs_runtime.Apply(bindProduct1_2_1, gopurs_runtime.Apply(dictMonad1_3.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeProduct2_4_2
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindProduct2_5_3
})})
})
})
	})
	return monadProduct
}


