package Data_Divisible

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Divide "gopurs/output/Data.Divide"
)

var divisiblePredicate gopurs_runtime.Value
var once_divisiblePredicate sync.Once
func Get_divisiblePredicate() gopurs_runtime.Value {
	once_divisiblePredicate.Do(func() {
		divisiblePredicate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"conquer": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_dividePredicate()
})})
	})
	return divisiblePredicate
}

var divisibleOp gopurs_runtime.Value
var once_divisibleOp sync.Once
func Get_divisibleOp() gopurs_runtime.Value {
	once_divisibleOp.Do(func() {
		divisibleOp = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
divideOp_1_0 := gopurs_runtime.Apply(pkg_Data_Divide.Get_divideOp(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
__local_var_2_1 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"conquer": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return divideOp_1_0
})})
})
	})
	return divisibleOp
}

var divisibleEquivalence gopurs_runtime.Value
var once_divisibleEquivalence sync.Once
func Get_divisibleEquivalence() gopurs_runtime.Value {
	once_divisibleEquivalence.Do(func() {
		divisibleEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"conquer": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideEquivalence()
})})
	})
	return divisibleEquivalence
}

var divisibleComparison gopurs_runtime.Value
var once_divisibleComparison sync.Once
func Get_divisibleComparison() gopurs_runtime.Value {
	once_divisibleComparison.Do(func() {
		divisibleComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"conquer": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideComparison()
})})
	})
	return divisibleComparison
}

var conquer gopurs_runtime.Value
var once_conquer sync.Once
func Get_conquer() gopurs_runtime.Value {
	once_conquer.Do(func() {
		conquer = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["conquer"]
})
	})
	return conquer
}


