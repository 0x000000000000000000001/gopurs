package Data_Semigroup_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericSemigroupNoConstructors gopurs_runtime.Value
var once_genericSemigroupNoConstructors sync.Once
func Get_genericSemigroupNoConstructors() gopurs_runtime.Value {
	once_genericSemigroupNoConstructors.Do(func() {
		genericSemigroupNoConstructors = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAppend'": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
})})
	})
	return genericSemigroupNoConstructors
}

var genericSemigroupNoArguments gopurs_runtime.Value
var once_genericSemigroupNoArguments sync.Once
func Get_genericSemigroupNoArguments() gopurs_runtime.Value {
	once_genericSemigroupNoArguments.Do(func() {
		genericSemigroupNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAppend'": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
})})
	})
	return genericSemigroupNoArguments
}

var genericSemigroupArgument gopurs_runtime.Value
var once_genericSemigroupArgument sync.Once
func Get_genericSemigroupArgument() gopurs_runtime.Value {
	once_genericSemigroupArgument.Do(func() {
		genericSemigroupArgument = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAppend'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_1), v1_2)
})
})})
})
	})
	return genericSemigroupArgument
}

var genericAppend_prime gopurs_runtime.Value
var once_genericAppend_prime sync.Once
func Get_genericAppend_prime() gopurs_runtime.Value {
	once_genericAppend_prime.Do(func() {
		genericAppend_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericAppend'"]
})
	})
	return genericAppend_prime
}

var genericSemigroupConstructor gopurs_runtime.Value
var once_genericSemigroupConstructor sync.Once
func Get_genericSemigroupConstructor() gopurs_runtime.Value {
	once_genericSemigroupConstructor.Do(func() {
		genericSemigroupConstructor = gopurs_runtime.Func(func(dictGenericSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAppend'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["genericAppend'"], v_1), v1_2)
})
})})
})
	})
	return genericSemigroupConstructor
}

var genericSemigroupProduct gopurs_runtime.Value
var once_genericSemigroupProduct sync.Once
func Get_genericSemigroupProduct() gopurs_runtime.Value {
	once_genericSemigroupProduct.Do(func() {
		genericSemigroupProduct = gopurs_runtime.Func(func(dictGenericSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericSemigroup1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericAppend'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["genericAppend'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemigroup1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericAppend'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
})
	})
	return genericSemigroupProduct
}

var genericAppend gopurs_runtime.Value
var once_genericAppend sync.Once
func Get_genericAppend() gopurs_runtime.Value {
	once_genericAppend.Do(func() {
		genericAppend = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["genericAppend'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3)))
})
})
})
})
	})
	return genericAppend
}


