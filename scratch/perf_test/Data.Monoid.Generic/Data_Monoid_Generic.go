package Data_Monoid_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericMonoidNoArguments gopurs_runtime.Value
var once_genericMonoidNoArguments sync.Once
func Get_genericMonoidNoArguments() gopurs_runtime.Value {
	once_genericMonoidNoArguments.Do(func() {
		genericMonoidNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericMempty'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})})
	})
	return genericMonoidNoArguments
}

var genericMonoidArgument gopurs_runtime.Value
var once_genericMonoidArgument sync.Once
func Get_genericMonoidArgument() gopurs_runtime.Value {
	once_genericMonoidArgument.Do(func() {
		genericMonoidArgument = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericMempty'": dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]})
})
	})
	return genericMonoidArgument
}

var genericMempty_prime gopurs_runtime.Value
var once_genericMempty_prime sync.Once
func Get_genericMempty_prime() gopurs_runtime.Value {
	once_genericMempty_prime.Do(func() {
		genericMempty_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericMempty'"]
})
	})
	return genericMempty_prime
}

var genericMonoidConstructor gopurs_runtime.Value
var once_genericMonoidConstructor sync.Once
func Get_genericMonoidConstructor() gopurs_runtime.Value {
	once_genericMonoidConstructor.Do(func() {
		genericMonoidConstructor = gopurs_runtime.Func(func(dictGenericMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericMempty'": dictGenericMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["genericMempty'"]})
})
	})
	return genericMonoidConstructor
}

var genericMonoidProduct gopurs_runtime.Value
var once_genericMonoidProduct sync.Once
func Get_genericMonoidProduct() gopurs_runtime.Value {
	once_genericMonoidProduct.Do(func() {
		genericMonoidProduct = gopurs_runtime.Func(func(dictGenericMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericMempty_prime1_1_0 := dictGenericMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["genericMempty'"]
return gopurs_runtime.Func(func(dictGenericMonoid1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericMempty'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": genericMempty_prime1_1_0, "value1": dictGenericMonoid1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericMempty'"]})})
})
})
	})
	return genericMonoidProduct
}

var genericMempty gopurs_runtime.Value
var once_genericMempty sync.Once
func Get_genericMempty() gopurs_runtime.Value {
	once_genericMempty.Do(func() {
		genericMempty = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], dictGenericMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["genericMempty'"])
})
})
	})
	return genericMempty
}


