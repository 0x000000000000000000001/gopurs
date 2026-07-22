package Data_Profunctor_Join

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

var semigroupJoin gopurs_runtime.Value
var once_semigroupJoin sync.Once
func Get_semigroupJoin() gopurs_runtime.Value {
	once_semigroupJoin.Do(func() {
		semigroupJoin = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupoid_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], v_1), v1_2)
})
})})
})
	})
	return semigroupJoin
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

var monoidJoin gopurs_runtime.Value
var once_monoidJoin sync.Once
func Get_monoidJoin() gopurs_runtime.Value {
	once_monoidJoin.Do(func() {
		monoidJoin = gopurs_runtime.Func(func(dictCategory_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictCategory_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroupoid0"], gopurs_runtime.Value{})
semigroupJoin1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], v_2), v1_3)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictCategory_0.PtrVal.(map[string]gopurs_runtime.Value)["identity"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupJoin1_2_1
})})
})
	})
	return monoidJoin
}

var invariantJoin gopurs_runtime.Value
var once_invariantJoin sync.Once
func Get_invariantJoin() gopurs_runtime.Value {
	once_invariantJoin.Do(func() {
		invariantJoin = gopurs_runtime.Func(func(dictProfunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], g_2), f_1), v_3)
})
})
})})
})
	})
	return invariantJoin
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


