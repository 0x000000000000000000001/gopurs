package Data_HeytingAlgebra_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericTT_prime gopurs_runtime.Value
var once_genericTT_prime sync.Once
func Get_genericTT_prime() gopurs_runtime.Value {
	once_genericTT_prime.Do(func() {
		genericTT_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericTT'"]
})
	})
	return genericTT_prime
}

var genericTT gopurs_runtime.Value
var once_genericTT sync.Once
func Get_genericTT() gopurs_runtime.Value {
	once_genericTT.Do(func() {
		genericTT = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], dictGenericHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["genericTT'"])
})
})
	})
	return genericTT
}

var genericNot_prime gopurs_runtime.Value
var once_genericNot_prime sync.Once
func Get_genericNot_prime() gopurs_runtime.Value {
	once_genericNot_prime.Do(func() {
		genericNot_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericNot'"]
})
	})
	return genericNot_prime
}

var genericNot gopurs_runtime.Value
var once_genericNot sync.Once
func Get_genericNot() gopurs_runtime.Value {
	once_genericNot.Do(func() {
		genericNot = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(dictGenericHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["genericNot'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)))
})
})
})
	})
	return genericNot
}

var genericImplies_prime gopurs_runtime.Value
var once_genericImplies_prime sync.Once
func Get_genericImplies_prime() gopurs_runtime.Value {
	once_genericImplies_prime.Do(func() {
		genericImplies_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericImplies'"]
})
	})
	return genericImplies_prime
}

var genericImplies gopurs_runtime.Value
var once_genericImplies sync.Once
func Get_genericImplies() gopurs_runtime.Value {
	once_genericImplies.Do(func() {
		genericImplies = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["genericImplies'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3)))
})
})
})
})
	})
	return genericImplies
}

var genericHeytingAlgebraNoArguments gopurs_runtime.Value
var once_genericHeytingAlgebraNoArguments sync.Once
func Get_genericHeytingAlgebraNoArguments() gopurs_runtime.Value {
	once_genericHeytingAlgebraNoArguments.Do(func() {
		genericHeytingAlgebraNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericFF'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")}), "genericTT'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")}), "genericImplies'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})
})
}), "genericConj'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})
})
}), "genericDisj'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})
})
}), "genericNot'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})
})})
	})
	return genericHeytingAlgebraNoArguments
}

var genericHeytingAlgebraArgument gopurs_runtime.Value
var once_genericHeytingAlgebraArgument sync.Once
func Get_genericHeytingAlgebraArgument() gopurs_runtime.Value {
	once_genericHeytingAlgebraArgument.Do(func() {
		genericHeytingAlgebraArgument = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericFF'": dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"], "genericTT'": dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["tt"], "genericImplies'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["implies"], v_1), v1_2)
})
}), "genericConj'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["conj"], v_1), v1_2)
})
}), "genericDisj'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_1), v1_2)
})
}), "genericNot'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["not"], v_1)
})})
})
	})
	return genericHeytingAlgebraArgument
}

var genericFF_prime gopurs_runtime.Value
var once_genericFF_prime sync.Once
func Get_genericFF_prime() gopurs_runtime.Value {
	once_genericFF_prime.Do(func() {
		genericFF_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericFF'"]
})
	})
	return genericFF_prime
}

var genericFF gopurs_runtime.Value
var once_genericFF sync.Once
func Get_genericFF() gopurs_runtime.Value {
	once_genericFF.Do(func() {
		genericFF = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], dictGenericHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["genericFF'"])
})
})
	})
	return genericFF
}

var genericDisj_prime gopurs_runtime.Value
var once_genericDisj_prime sync.Once
func Get_genericDisj_prime() gopurs_runtime.Value {
	once_genericDisj_prime.Do(func() {
		genericDisj_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericDisj'"]
})
	})
	return genericDisj_prime
}

var genericDisj gopurs_runtime.Value
var once_genericDisj sync.Once
func Get_genericDisj() gopurs_runtime.Value {
	once_genericDisj.Do(func() {
		genericDisj = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["genericDisj'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3)))
})
})
})
})
	})
	return genericDisj
}

var genericConj_prime gopurs_runtime.Value
var once_genericConj_prime sync.Once
func Get_genericConj_prime() gopurs_runtime.Value {
	once_genericConj_prime.Do(func() {
		genericConj_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericConj'"]
})
	})
	return genericConj_prime
}

var genericHeytingAlgebraConstructor gopurs_runtime.Value
var once_genericHeytingAlgebraConstructor sync.Once
func Get_genericHeytingAlgebraConstructor() gopurs_runtime.Value {
	once_genericHeytingAlgebraConstructor.Do(func() {
		genericHeytingAlgebraConstructor = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericFF'": dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericFF'"], "genericTT'": dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericTT'"], "genericImplies'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericImplies'"], v_1), v1_2)
})
}), "genericConj'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericConj'"], v_1), v1_2)
})
}), "genericDisj'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericDisj'"], v_1), v1_2)
})
}), "genericNot'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericNot'"], v_1)
})})
})
	})
	return genericHeytingAlgebraConstructor
}

var genericHeytingAlgebraProduct gopurs_runtime.Value
var once_genericHeytingAlgebraProduct sync.Once
func Get_genericHeytingAlgebraProduct() gopurs_runtime.Value {
	once_genericHeytingAlgebraProduct.Do(func() {
		genericHeytingAlgebraProduct = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericFF_prime1_1_0 := dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericFF'"]
genericTT_prime1_2_1 := dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericTT'"]
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericFF'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": genericFF_prime1_1_0, "value1": dictGenericHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericFF'"]}), "genericTT'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": genericTT_prime1_2_1, "value1": dictGenericHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericTT'"]}), "genericImplies'": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericImplies'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericImplies'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "genericConj'": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericConj'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericConj'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "genericDisj'": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericDisj'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericDisj'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "genericNot'": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(dictGenericHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["genericNot'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(dictGenericHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericNot'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})})
})
})
	})
	return genericHeytingAlgebraProduct
}

var genericConj gopurs_runtime.Value
var once_genericConj sync.Once
func Get_genericConj() gopurs_runtime.Value {
	once_genericConj.Do(func() {
		genericConj = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["genericConj'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3)))
})
})
})
})
	})
	return genericConj
}


