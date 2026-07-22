package Data_Functor_Flip

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Flip gopurs_runtime.Value
var once_Flip sync.Once
func Get_Flip() gopurs_runtime.Value {
	once_Flip.Do(func() {
		Flip = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Flip
}

var showFlip gopurs_runtime.Value
var once_showFlip sync.Once
func Get_showFlip() gopurs_runtime.Value {
	once_showFlip.Do(func() {
		showFlip = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Flip ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showFlip
}

var semigroupoidFlip gopurs_runtime.Value
var once_semigroupoidFlip sync.Once
func Get_semigroupoidFlip() gopurs_runtime.Value {
	once_semigroupoidFlip.Do(func() {
		semigroupoidFlip = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupoid_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], v1_2), v_1)
})
})})
})
	})
	return semigroupoidFlip
}

var ordFlip gopurs_runtime.Value
var once_ordFlip sync.Once
func Get_ordFlip() gopurs_runtime.Value {
	once_ordFlip.Do(func() {
		ordFlip = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordFlip
}

var newtypeFlip gopurs_runtime.Value
var once_newtypeFlip sync.Once
func Get_newtypeFlip() gopurs_runtime.Value {
	once_newtypeFlip.Do(func() {
		newtypeFlip = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeFlip
}

var functorFlip gopurs_runtime.Value
var once_functorFlip sync.Once
func Get_functorFlip() gopurs_runtime.Value {
	once_functorFlip.Do(func() {
		functorFlip = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), v_2)
})
})})
})
	})
	return functorFlip
}

var eqFlip gopurs_runtime.Value
var once_eqFlip sync.Once
func Get_eqFlip() gopurs_runtime.Value {
	once_eqFlip.Do(func() {
		eqFlip = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqFlip
}

var contravariantFlip gopurs_runtime.Value
var once_contravariantFlip sync.Once
func Get_contravariantFlip() gopurs_runtime.Value {
	once_contravariantFlip.Do(func() {
		contravariantFlip = gopurs_runtime.Func(func(dictProfunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cmap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], f_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), v_2)
})
})})
})
	})
	return contravariantFlip
}

var categoryFlip gopurs_runtime.Value
var once_categoryFlip sync.Once
func Get_categoryFlip() gopurs_runtime.Value {
	once_categoryFlip.Do(func() {
		categoryFlip = gopurs_runtime.Func(func(dictCategory_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictCategory_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroupoid0"], gopurs_runtime.Value{})
semigroupoidFlip1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], v1_3), v_2)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"identity": dictCategory_0.PtrVal.(map[string]gopurs_runtime.Value)["identity"], "Semigroupoid0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidFlip1_2_1
})})
})
	})
	return categoryFlip
}

var bifunctorFlip gopurs_runtime.Value
var once_bifunctorFlip sync.Once
func Get_bifunctorFlip() gopurs_runtime.Value {
	once_bifunctorFlip.Do(func() {
		bifunctorFlip = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], g_2), f_1), v_3)
})
})
})})
})
	})
	return bifunctorFlip
}

var biapplyFlip gopurs_runtime.Value
var once_biapplyFlip sync.Once
func Get_biapplyFlip() gopurs_runtime.Value {
	once_biapplyFlip.Do(func() {
		biapplyFlip = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorFlip1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], g_3), f_2), v_4)
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"biapply": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_3), v1_4)
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip1_2_1
})})
})
	})
	return biapplyFlip
}

var biapplicativeFlip gopurs_runtime.Value
var once_biapplicativeFlip sync.Once
func Get_biapplicativeFlip() gopurs_runtime.Value {
	once_biapplicativeFlip.Do(func() {
		biapplicativeFlip = gopurs_runtime.Func(func(dictBiapplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBiapplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Biapply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorFlip1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], g_4), f_3), v_5)
})
})
})})
biapplyFlip1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"biapply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_4), v1_5)
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bipure": gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["bipure"], b_5), a_4)
})
}), "Biapply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyFlip1_3_2
})})
})
	})
	return biapplicativeFlip
}


