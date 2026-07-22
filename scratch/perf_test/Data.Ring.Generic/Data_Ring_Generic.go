package Data_Ring_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericSub_prime gopurs_runtime.Value
var once_genericSub_prime sync.Once
func Get_genericSub_prime() gopurs_runtime.Value {
	once_genericSub_prime.Do(func() {
		genericSub_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericSub'"]
})
	})
	return genericSub_prime
}

var genericSub gopurs_runtime.Value
var once_genericSub sync.Once
func Get_genericSub() gopurs_runtime.Value {
	once_genericSub.Do(func() {
		genericSub = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericRing_1.PtrVal.(map[string]gopurs_runtime.Value)["genericSub'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3)))
})
})
})
})
	})
	return genericSub
}

var genericRingProduct gopurs_runtime.Value
var once_genericRingProduct sync.Once
func Get_genericRingProduct() gopurs_runtime.Value {
	once_genericRingProduct.Do(func() {
		genericRingProduct = gopurs_runtime.Func(func(dictGenericRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericRing1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericSub'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericRing_0.PtrVal.(map[string]gopurs_runtime.Value)["genericSub'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericRing1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericSub'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
})
	})
	return genericRingProduct
}

var genericRingNoArguments gopurs_runtime.Value
var once_genericRingNoArguments sync.Once
func Get_genericRingNoArguments() gopurs_runtime.Value {
	once_genericRingNoArguments.Do(func() {
		genericRingNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericSub'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})
})
})})
	})
	return genericRingNoArguments
}

var genericRingConstructor gopurs_runtime.Value
var once_genericRingConstructor sync.Once
func Get_genericRingConstructor() gopurs_runtime.Value {
	once_genericRingConstructor.Do(func() {
		genericRingConstructor = gopurs_runtime.Func(func(dictGenericRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericSub'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericRing_0.PtrVal.(map[string]gopurs_runtime.Value)["genericSub'"], v_1), v1_2)
})
})})
})
	})
	return genericRingConstructor
}

var genericRingArgument gopurs_runtime.Value
var once_genericRingArgument sync.Once
func Get_genericRingArgument() gopurs_runtime.Value {
	once_genericRingArgument.Do(func() {
		genericRingArgument = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericSub'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["sub"], v_1), v1_2)
})
})})
})
	})
	return genericRingArgument
}


