package Data_Bounded_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericTopNoArguments gopurs_runtime.Value
var once_genericTopNoArguments sync.Once
func Get_genericTopNoArguments() gopurs_runtime.Value {
	once_genericTopNoArguments.Do(func() {
		genericTopNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericTop'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})})
	})
	return genericTopNoArguments
}

var genericTopArgument gopurs_runtime.Value
var once_genericTopArgument sync.Once
func Get_genericTopArgument() gopurs_runtime.Value {
	once_genericTopArgument.Do(func() {
		genericTopArgument = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericTop'": dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["top"]})
})
	})
	return genericTopArgument
}

var genericTop_prime gopurs_runtime.Value
var once_genericTop_prime sync.Once
func Get_genericTop_prime() gopurs_runtime.Value {
	once_genericTop_prime.Do(func() {
		genericTop_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"]
})
	})
	return genericTop_prime
}

var genericTopConstructor gopurs_runtime.Value
var once_genericTopConstructor sync.Once
func Get_genericTopConstructor() gopurs_runtime.Value {
	once_genericTopConstructor.Do(func() {
		genericTopConstructor = gopurs_runtime.Func(func(dictGenericTop_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericTop'": dictGenericTop_0.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"]})
})
	})
	return genericTopConstructor
}

var genericTopProduct gopurs_runtime.Value
var once_genericTopProduct sync.Once
func Get_genericTopProduct() gopurs_runtime.Value {
	once_genericTopProduct.Do(func() {
		genericTopProduct = gopurs_runtime.Func(func(dictGenericTop_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericTop_prime1_1_0 := dictGenericTop_0.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"]
return gopurs_runtime.Func(func(dictGenericTop1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericTop'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": genericTop_prime1_1_0, "value1": dictGenericTop1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"]})})
})
})
	})
	return genericTopProduct
}

var genericTopSum gopurs_runtime.Value
var once_genericTopSum sync.Once
func Get_genericTopSum() gopurs_runtime.Value {
	once_genericTopSum.Do(func() {
		genericTopSum = gopurs_runtime.Func(func(dictGenericTop_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericTop'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inr"), "value0": dictGenericTop_0.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"]})})
})
	})
	return genericTopSum
}

var genericTop gopurs_runtime.Value
var once_genericTop sync.Once
func Get_genericTop() gopurs_runtime.Value {
	once_genericTop.Do(func() {
		genericTop = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericTop_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], dictGenericTop_1.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"])
})
})
	})
	return genericTop
}

var genericBottomNoArguments gopurs_runtime.Value
var once_genericBottomNoArguments sync.Once
func Get_genericBottomNoArguments() gopurs_runtime.Value {
	once_genericBottomNoArguments.Do(func() {
		genericBottomNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericBottom'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})})
	})
	return genericBottomNoArguments
}

var genericBottomArgument gopurs_runtime.Value
var once_genericBottomArgument sync.Once
func Get_genericBottomArgument() gopurs_runtime.Value {
	once_genericBottomArgument.Do(func() {
		genericBottomArgument = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericBottom'": dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]})
})
	})
	return genericBottomArgument
}

var genericBottom_prime gopurs_runtime.Value
var once_genericBottom_prime sync.Once
func Get_genericBottom_prime() gopurs_runtime.Value {
	once_genericBottom_prime.Do(func() {
		genericBottom_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"]
})
	})
	return genericBottom_prime
}

var genericBottomConstructor gopurs_runtime.Value
var once_genericBottomConstructor sync.Once
func Get_genericBottomConstructor() gopurs_runtime.Value {
	once_genericBottomConstructor.Do(func() {
		genericBottomConstructor = gopurs_runtime.Func(func(dictGenericBottom_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericBottom'": dictGenericBottom_0.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"]})
})
	})
	return genericBottomConstructor
}

var genericBottomProduct gopurs_runtime.Value
var once_genericBottomProduct sync.Once
func Get_genericBottomProduct() gopurs_runtime.Value {
	once_genericBottomProduct.Do(func() {
		genericBottomProduct = gopurs_runtime.Func(func(dictGenericBottom_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime1_1_0 := dictGenericBottom_0.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"]
return gopurs_runtime.Func(func(dictGenericBottom1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericBottom'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": genericBottom_prime1_1_0, "value1": dictGenericBottom1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"]})})
})
})
	})
	return genericBottomProduct
}

var genericBottomSum gopurs_runtime.Value
var once_genericBottomSum sync.Once
func Get_genericBottomSum() gopurs_runtime.Value {
	once_genericBottomSum.Do(func() {
		genericBottomSum = gopurs_runtime.Func(func(dictGenericBottom_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericBottom'": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inl"), "value0": dictGenericBottom_0.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"]})})
})
	})
	return genericBottomSum
}

var genericBottom gopurs_runtime.Value
var once_genericBottom sync.Once
func Get_genericBottom() gopurs_runtime.Value {
	once_genericBottom.Do(func() {
		genericBottom = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericBottom_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], dictGenericBottom_1.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"])
})
})
	})
	return genericBottom
}


