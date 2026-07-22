package Data_Semiring_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericZero_prime gopurs_runtime.Value
var once_genericZero_prime sync.Once
func Get_genericZero_prime() gopurs_runtime.Value {
	once_genericZero_prime.Do(func() {
		genericZero_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericZero'"]
})
	})
	return genericZero_prime
}

var genericZero gopurs_runtime.Value
var once_genericZero sync.Once
func Get_genericZero() gopurs_runtime.Value {
	once_genericZero.Do(func() {
		genericZero = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], dictGenericSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["genericZero'"])
})
})
	})
	return genericZero
}

var genericSemiringNoArguments gopurs_runtime.Value
var once_genericSemiringNoArguments sync.Once
func Get_genericSemiringNoArguments() gopurs_runtime.Value {
	once_genericSemiringNoArguments.Do(func() {
		genericSemiringNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAdd'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})
})
}), "genericZero'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")}), "genericMul'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})
})
}), "genericOne'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})})
	})
	return genericSemiringNoArguments
}

var genericSemiringArgument gopurs_runtime.Value
var once_genericSemiringArgument sync.Once
func Get_genericSemiringArgument() gopurs_runtime.Value {
	once_genericSemiringArgument.Do(func() {
		genericSemiringArgument = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAdd'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["add"], v_1), v1_2)
})
}), "genericZero'": dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"], "genericMul'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], v_1), v1_2)
})
}), "genericOne'": dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]})
})
	})
	return genericSemiringArgument
}

var genericOne_prime gopurs_runtime.Value
var once_genericOne_prime sync.Once
func Get_genericOne_prime() gopurs_runtime.Value {
	once_genericOne_prime.Do(func() {
		genericOne_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericOne'"]
})
	})
	return genericOne_prime
}

var genericOne gopurs_runtime.Value
var once_genericOne sync.Once
func Get_genericOne() gopurs_runtime.Value {
	once_genericOne.Do(func() {
		genericOne = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], dictGenericSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["genericOne'"])
})
})
	})
	return genericOne
}

var genericMul_prime gopurs_runtime.Value
var once_genericMul_prime sync.Once
func Get_genericMul_prime() gopurs_runtime.Value {
	once_genericMul_prime.Do(func() {
		genericMul_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericMul'"]
})
	})
	return genericMul_prime
}

var genericMul gopurs_runtime.Value
var once_genericMul sync.Once
func Get_genericMul() gopurs_runtime.Value {
	once_genericMul.Do(func() {
		genericMul = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["genericMul'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3)))
})
})
})
})
	})
	return genericMul
}

var genericAdd_prime gopurs_runtime.Value
var once_genericAdd_prime sync.Once
func Get_genericAdd_prime() gopurs_runtime.Value {
	once_genericAdd_prime.Do(func() {
		genericAdd_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericAdd'"]
})
	})
	return genericAdd_prime
}

var genericSemiringConstructor gopurs_runtime.Value
var once_genericSemiringConstructor sync.Once
func Get_genericSemiringConstructor() gopurs_runtime.Value {
	once_genericSemiringConstructor.Do(func() {
		genericSemiringConstructor = gopurs_runtime.Func(func(dictGenericSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAdd'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericAdd'"], v_1), v1_2)
})
}), "genericZero'": dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericZero'"], "genericMul'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericMul'"], v_1), v1_2)
})
}), "genericOne'": dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericOne'"]})
})
	})
	return genericSemiringConstructor
}

var genericSemiringProduct gopurs_runtime.Value
var once_genericSemiringProduct sync.Once
func Get_genericSemiringProduct() gopurs_runtime.Value {
	once_genericSemiringProduct.Do(func() {
		genericSemiringProduct = gopurs_runtime.Func(func(dictGenericSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericZero_prime1_1_0 := dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericZero'"]
genericOne_prime1_2_1 := dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericOne'"]
return gopurs_runtime.Func(func(dictGenericSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAdd'": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericAdd'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericAdd'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "genericZero'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": genericZero_prime1_1_0, "value1": dictGenericSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericZero'"]}), "genericMul'": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["genericMul'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericMul'"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "genericOne'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": genericOne_prime1_2_1, "value1": dictGenericSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericOne'"]})})
})
})
	})
	return genericSemiringProduct
}

var genericAdd gopurs_runtime.Value
var once_genericAdd sync.Once
func Get_genericAdd() gopurs_runtime.Value {
	once_genericAdd.Do(func() {
		genericAdd = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["genericAdd'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3)))
})
})
})
})
	})
	return genericAdd
}


