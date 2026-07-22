package Data_Ord_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericOrdNoConstructors gopurs_runtime.Value
var once_genericOrdNoConstructors sync.Once
func Get_genericOrdNoConstructors() gopurs_runtime.Value {
	once_genericOrdNoConstructors.Do(func() {
		genericOrdNoConstructors = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCompare'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
})})
	})
	return genericOrdNoConstructors
}

var genericOrdNoArguments gopurs_runtime.Value
var once_genericOrdNoArguments sync.Once
func Get_genericOrdNoArguments() gopurs_runtime.Value {
	once_genericOrdNoArguments.Do(func() {
		genericOrdNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCompare'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
})})
	})
	return genericOrdNoArguments
}

var genericOrdArgument gopurs_runtime.Value
var once_genericOrdArgument sync.Once
func Get_genericOrdArgument() gopurs_runtime.Value {
	once_genericOrdArgument.Do(func() {
		genericOrdArgument = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCompare'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_1), v1_2)
})
})})
})
	})
	return genericOrdArgument
}

var genericCompare_prime gopurs_runtime.Value
var once_genericCompare_prime sync.Once
func Get_genericCompare_prime() gopurs_runtime.Value {
	once_genericCompare_prime.Do(func() {
		genericCompare_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCompare'"]
})
	})
	return genericCompare_prime
}

var genericOrdConstructor gopurs_runtime.Value
var once_genericOrdConstructor sync.Once
func Get_genericOrdConstructor() gopurs_runtime.Value {
	once_genericOrdConstructor.Do(func() {
		genericOrdConstructor = gopurs_runtime.Func(func(dictGenericOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCompare'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCompare'"], v_1), v1_2)
})
})})
})
	})
	return genericOrdConstructor
}

var genericOrdProduct gopurs_runtime.Value
var once_genericOrdProduct sync.Once
func Get_genericOrdProduct() gopurs_runtime.Value {
	once_genericOrdProduct.Do(func() {
		genericOrdProduct = gopurs_runtime.Func(func(dictGenericOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericOrd1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCompare'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCompare'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericOrd1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericCompare'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = v2_4_0
}
return __t1
})
})})
})
})
	})
	return genericOrdProduct
}

var genericOrdSum gopurs_runtime.Value
var once_genericOrdSum sync.Once
func Get_genericOrdSum() gopurs_runtime.Value {
	once_genericOrdSum.Do(func() {
		genericOrdSum = gopurs_runtime.Func(func(dictGenericOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericOrd1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCompare'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCompare'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t0 = __t2
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericOrd1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericCompare'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})})
})
})
	})
	return genericOrdSum
}

var genericCompare gopurs_runtime.Value
var once_genericCompare sync.Once
func Get_genericCompare() gopurs_runtime.Value {
	once_genericCompare.Do(func() {
		genericCompare = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericOrd_1.PtrVal.(map[string]gopurs_runtime.Value)["genericCompare'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3))
})
})
})
})
	})
	return genericCompare
}


