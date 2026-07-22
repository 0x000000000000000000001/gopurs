package Data_Show_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var genericShowArgsNoArguments gopurs_runtime.Value
var once_genericShowArgsNoArguments sync.Once
func Get_genericShowArgsNoArguments() gopurs_runtime.Value {
	once_genericShowArgsNoArguments.Do(func() {
		genericShowArgsNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericShowArgs": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{})
})})
	})
	return genericShowArgsNoArguments
}

var genericShowArgsArgument gopurs_runtime.Value
var once_genericShowArgsArgument sync.Once
func Get_genericShowArgsArgument() gopurs_runtime.Value {
	once_genericShowArgsArgument.Do(func() {
		genericShowArgsArgument = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericShowArgs": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)})
})})
})
	})
	return genericShowArgsArgument
}

var genericShowArgs gopurs_runtime.Value
var once_genericShowArgs sync.Once
func Get_genericShowArgs() gopurs_runtime.Value {
	once_genericShowArgs.Do(func() {
		genericShowArgs = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericShowArgs"]
})
	})
	return genericShowArgs
}

var genericShowArgsProduct gopurs_runtime.Value
var once_genericShowArgsProduct sync.Once
func Get_genericShowArgsProduct() gopurs_runtime.Value {
	once_genericShowArgsProduct.Do(func() {
		genericShowArgsProduct = gopurs_runtime.Func(func(dictGenericShowArgs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericShowArgs1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericShowArgs": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(dictGenericShowArgs_0.PtrVal.(map[string]gopurs_runtime.Value)["genericShowArgs"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(dictGenericShowArgs1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericShowArgs"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})})
})
})
	})
	return genericShowArgsProduct
}

var genericShowConstructor gopurs_runtime.Value
var once_genericShowConstructor sync.Once
func Get_genericShowConstructor() gopurs_runtime.Value {
	once_genericShowConstructor.Do(func() {
		genericShowConstructor = gopurs_runtime.Func(func(dictGenericShowArgs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericShow'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
ctor_3_0 := gopurs_runtime.Apply(dictIsSymbol_1.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
v1_4_1 := gopurs_runtime.Apply(dictGenericShowArgs_0.PtrVal.(map[string]gopurs_runtime.Value)["genericShowArgs"], v_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Int(len(v1_4_1.PtrVal.([]gopurs_runtime.Value))).IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = ctor_3_0
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_intercalate(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{ctor_3_0})), v1_4_1))), gopurs_runtime.Str(")")))
}
return __t2
})})
})
})
	})
	return genericShowConstructor
}

var genericShow_prime gopurs_runtime.Value
var once_genericShow_prime sync.Once
func Get_genericShow_prime() gopurs_runtime.Value {
	once_genericShow_prime.Do(func() {
		genericShow_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericShow'"]
})
	})
	return genericShow_prime
}

var genericShowNoConstructors gopurs_runtime.Value
var once_genericShowNoConstructors sync.Once
func Get_genericShowNoConstructors() gopurs_runtime.Value {
	once_genericShowNoConstructors.Do(func() {
		genericShowNoConstructors = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericShow'": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_genericShowNoConstructors().PtrVal.(map[string]gopurs_runtime.Value)["genericShow'"], a_0)
})})
	})
	return genericShowNoConstructors
}

var genericShowSum gopurs_runtime.Value
var once_genericShowSum sync.Once
func Get_genericShowSum() gopurs_runtime.Value {
	once_genericShowSum.Do(func() {
		genericShowSum = gopurs_runtime.Func(func(dictGenericShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericShow'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(dictGenericShow_0.PtrVal.(map[string]gopurs_runtime.Value)["genericShow'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(dictGenericShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericShow'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})})
})
})
	})
	return genericShowSum
}

var genericShow gopurs_runtime.Value
var once_genericShow sync.Once
func Get_genericShow() gopurs_runtime.Value {
	once_genericShow.Do(func() {
		genericShow = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGenericShow_1.PtrVal.(map[string]gopurs_runtime.Value)["genericShow'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2))
})
})
})
	})
	return genericShow
}

func Get_intercalate() gopurs_runtime.Value {
	return Intercalate
}
