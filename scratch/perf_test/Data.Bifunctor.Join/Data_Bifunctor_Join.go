package Data_Bifunctor_Join

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Join gopurs_runtime.Value
var once_Join sync.Once
func Get_Join() gopurs_runtime.Value {
	once_Join.Do(func() {
		Join = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Join
}

var showJoin gopurs_runtime.Value
var once_showJoin sync.Once
func Get_showJoin() gopurs_runtime.Value {
	once_showJoin.Do(func() {
		showJoin = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Join ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showJoin
}

var ordJoin gopurs_runtime.Value
var once_ordJoin sync.Once
func Get_ordJoin() gopurs_runtime.Value {
	once_ordJoin.Do(func() {
		ordJoin = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordJoin
}

var newtypeJoin gopurs_runtime.Value
var once_newtypeJoin sync.Once
func Get_newtypeJoin() gopurs_runtime.Value {
	once_newtypeJoin.Do(func() {
		newtypeJoin = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeJoin
}

var eqJoin gopurs_runtime.Value
var once_eqJoin sync.Once
func Get_eqJoin() gopurs_runtime.Value {
	once_eqJoin.Do(func() {
		eqJoin = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqJoin
}

var bifunctorJoin gopurs_runtime.Value
var once_bifunctorJoin sync.Once
func Get_bifunctorJoin() gopurs_runtime.Value {
	once_bifunctorJoin.Do(func() {
		bifunctorJoin = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_1), f_1), v_2)
})
})})
})
	})
	return bifunctorJoin
}

var biapplyJoin gopurs_runtime.Value
var once_biapplyJoin sync.Once
func Get_biapplyJoin() gopurs_runtime.Value {
	once_biapplyJoin.Do(func() {
		biapplyJoin = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorJoin1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_2), f_2), v_3)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_3), v1_4)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoin1_2_1
})})
})
	})
	return biapplyJoin
}

var biapplicativeJoin gopurs_runtime.Value
var once_biapplicativeJoin sync.Once
func Get_biapplicativeJoin() gopurs_runtime.Value {
	once_biapplicativeJoin.Do(func() {
		biapplicativeJoin = gopurs_runtime.Func(func(dictBiapplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBiapplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Biapply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{})
bifunctorJoin1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_3), f_3), v_4)
})
})})
biapplyJoin1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], v_4), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoin1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["bipure"], a_4), a_4)
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyJoin1_3_2
})})
})
	})
	return biapplicativeJoin
}


