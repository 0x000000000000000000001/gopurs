package Data_Equivalence

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var Equivalence gopurs_runtime.Value
var once_Equivalence sync.Once
func Get_Equivalence() gopurs_runtime.Value {
	once_Equivalence.Do(func() {
		Equivalence = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Equivalence
}

var semigroupEquivalence gopurs_runtime.Value
var once_semigroupEquivalence sync.Once
func Get_semigroupEquivalence() gopurs_runtime.Value {
	once_semigroupEquivalence.Do(func() {
		semigroupEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, a_2), b_3)), gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, a_2), b_3))
})
})
})
})})
	})
	return semigroupEquivalence
}

var newtypeEquivalence gopurs_runtime.Value
var once_newtypeEquivalence sync.Once
func Get_newtypeEquivalence() gopurs_runtime.Value {
	once_newtypeEquivalence.Do(func() {
		newtypeEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeEquivalence
}

var monoidEquivalence gopurs_runtime.Value
var once_monoidEquivalence sync.Once
func Get_monoidEquivalence() gopurs_runtime.Value {
	once_monoidEquivalence.Do(func() {
		monoidEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupEquivalence()
})})
	})
	return monoidEquivalence
}

var defaultEquivalence gopurs_runtime.Value
var once_defaultEquivalence sync.Once
func Get_defaultEquivalence() gopurs_runtime.Value {
	once_defaultEquivalence.Do(func() {
		defaultEquivalence = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})
	})
	return defaultEquivalence
}

var contravariantEquivalence gopurs_runtime.Value
var once_contravariantEquivalence sync.Once
func Get_contravariantEquivalence() gopurs_runtime.Value {
	once_contravariantEquivalence.Do(func() {
		contravariantEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cmap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2)), gopurs_runtime.Apply(f_0, y_3))
})
})
})
})})
	})
	return contravariantEquivalence
}

var comparisonEquivalence gopurs_runtime.Value
var once_comparisonEquivalence sync.Once
func Get_comparisonEquivalence() gopurs_runtime.Value {
	once_comparisonEquivalence.Do(func() {
		comparisonEquivalence = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, a_1), b_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")
})
})
})
	})
	return comparisonEquivalence
}


