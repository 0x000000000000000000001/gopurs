package Data_Predicate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_BooleanAlgebra "gopurs/output/Data.BooleanAlgebra"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var Predicate gopurs_runtime.Value
var once_Predicate sync.Once
func Get_Predicate() gopurs_runtime.Value {
	once_Predicate.Do(func() {
		Predicate = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Predicate
}

var newtypePredicate gopurs_runtime.Value
var once_newtypePredicate sync.Once
func Get_newtypePredicate() gopurs_runtime.Value {
	once_newtypePredicate.Do(func() {
		newtypePredicate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypePredicate
}

var heytingAlgebraPredicate gopurs_runtime.Value
var once_heytingAlgebraPredicate sync.Once
func Get_heytingAlgebraPredicate() gopurs_runtime.Value {
	once_heytingAlgebraPredicate.Do(func() {
		heytingAlgebraPredicate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), "tt": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), "implies": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(f_0, a_2))), gopurs_runtime.Apply(g_1, a_2))
})
})
}), "conj": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(f_0, a_2)), gopurs_runtime.Apply(g_1, a_2))
})
})
}), "disj": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(f_0, a_2)), gopurs_runtime.Apply(g_1, a_2))
})
})
}), "not": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(f_0, a_1))
})
})})
	})
	return heytingAlgebraPredicate
}

var contravariantPredicate gopurs_runtime.Value
var once_contravariantPredicate sync.Once
func Get_contravariantPredicate() gopurs_runtime.Value {
	once_contravariantPredicate.Do(func() {
		contravariantPredicate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cmap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})})
	})
	return contravariantPredicate
}

var booleanAlgebraPredicate gopurs_runtime.Value
var once_booleanAlgebraPredicate sync.Once
func Get_booleanAlgebraPredicate() gopurs_runtime.Value {
	once_booleanAlgebraPredicate.Do(func() {
		booleanAlgebraPredicate = gopurs_runtime.Apply(pkg_Data_BooleanAlgebra.Get_booleanAlgebraFn(), pkg_Data_BooleanAlgebra.Get_booleanAlgebraBoolean())
	})
	return booleanAlgebraPredicate
}


