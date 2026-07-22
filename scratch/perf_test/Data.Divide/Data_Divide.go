package Data_Divide

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Comparison "gopurs/output/Data.Comparison"
	pkg_Data_Equivalence "gopurs/output/Data.Equivalence"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Op "gopurs/output/Data.Op"
	pkg_Data_Predicate "gopurs/output/Data.Predicate"
)

var dividePredicate gopurs_runtime.Value
var once_dividePredicate sync.Once
func Get_dividePredicate() gopurs_runtime.Value {
	once_dividePredicate.Do(func() {
		dividePredicate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"divide": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply(f_0, a_3)
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(v_1, v2_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(v1_2, v2_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
}), "Contravariant0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Predicate.Get_contravariantPredicate()
})})
	})
	return dividePredicate
}

var divideOp gopurs_runtime.Value
var once_divideOp sync.Once
func Get_divideOp() gopurs_runtime.Value {
	once_divideOp.Do(func() {
		divideOp = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"divide": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_1, a_4)
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(v_2, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(v1_3, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
}), "Contravariant0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Op.Get_contravariantOp()
})})
})
	})
	return divideOp
}

var divideEquivalence gopurs_runtime.Value
var once_divideEquivalence sync.Once
func Get_divideEquivalence() gopurs_runtime.Value {
	once_divideEquivalence.Do(func() {
		divideEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"divide": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v3_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(v1_2, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v3_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
})
}), "Contravariant0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Equivalence.Get_contravariantEquivalence()
})})
	})
	return divideEquivalence
}

var divideComparison gopurs_runtime.Value
var once_divideComparison sync.Once
func Get_divideComparison() gopurs_runtime.Value {
	once_divideComparison.Do(func() {
		divideComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"divide": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
v3_6_1 := gopurs_runtime.Apply(f_0, b_4)
__local_var_7_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v3_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
__local_var_8_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(v1_2, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v3_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_7_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
if (gopurs_runtime.Bool(__local_var_7_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(__local_var_7_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t4 = __local_var_8_3
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t4
})
})
})
})
}), "Contravariant0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Comparison.Get_contravariantComparison()
})})
	})
	return divideComparison
}

var divide gopurs_runtime.Value
var once_divide sync.Once
func Get_divide() gopurs_runtime.Value {
	once_divide.Do(func() {
		divide = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["divide"]
})
	})
	return divide
}

var divided gopurs_runtime.Value
var once_divided sync.Once
func Get_divided() gopurs_runtime.Value {
	once_divided.Do(func() {
		divided = gopurs_runtime.Func(func(dictDivide_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictDivide_0.PtrVal.(map[string]gopurs_runtime.Value)["divide"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return divided
}


