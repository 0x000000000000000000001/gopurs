package Data_Eq_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var genericEqNoConstructors gopurs_runtime.Value
var once_genericEqNoConstructors sync.Once
func Get_genericEqNoConstructors() gopurs_runtime.Value {
	once_genericEqNoConstructors.Do(func() {
		genericEqNoConstructors = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericEq'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})})
	})
	return genericEqNoConstructors
}

var genericEqNoArguments gopurs_runtime.Value
var once_genericEqNoArguments sync.Once
func Get_genericEqNoArguments() gopurs_runtime.Value {
	once_genericEqNoArguments.Do(func() {
		genericEqNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericEq'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})})
	})
	return genericEqNoArguments
}

var genericEqArgument gopurs_runtime.Value
var once_genericEqArgument sync.Once
func Get_genericEqArgument() gopurs_runtime.Value {
	once_genericEqArgument.Do(func() {
		genericEqArgument = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericEq'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_1), v1_2)
})
})})
})
	})
	return genericEqArgument
}

var genericEq_prime gopurs_runtime.Value
var once_genericEq_prime sync.Once
func Get_genericEq_prime() gopurs_runtime.Value {
	once_genericEq_prime.Do(func() {
		genericEq_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericEq'"]
})
	})
	return genericEq_prime
}

var genericEqConstructor gopurs_runtime.Value
var once_genericEqConstructor sync.Once
func Get_genericEqConstructor() gopurs_runtime.Value {
	once_genericEqConstructor.Do(func() {
		genericEqConstructor = gopurs_runtime.Func(func(dictGenericEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericEq'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericEq_0.PtrVal.(map[string]gopurs_runtime.Value)["genericEq'"], v_1), v1_2)
})
})})
})
	})
	return genericEqConstructor
}

var genericEqProduct gopurs_runtime.Value
var once_genericEqProduct sync.Once
func Get_genericEqProduct() gopurs_runtime.Value {
	once_genericEqProduct.Do(func() {
		genericEqProduct = gopurs_runtime.Func(func(dictGenericEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericEq'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericEq_0.PtrVal.(map[string]gopurs_runtime.Value)["genericEq'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericEq'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})})
})
})
	})
	return genericEqProduct
}

var genericEqSum gopurs_runtime.Value
var once_genericEqSum sync.Once
func Get_genericEqSum() gopurs_runtime.Value {
	once_genericEqSum.Do(func() {
		genericEqSum = gopurs_runtime.Func(func(dictGenericEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericEq'": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericEq_0.PtrVal.(map[string]gopurs_runtime.Value)["genericEq'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["genericEq'"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t0
})
})})
})
})
	})
	return genericEqSum
}

var genericEq gopurs_runtime.Value
var once_genericEq sync.Once
func Get_genericEq() gopurs_runtime.Value {
	once_genericEq.Do(func() {
		genericEq = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictGenericEq_1.PtrVal.(map[string]gopurs_runtime.Value)["genericEq'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2)), gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], y_3))
})
})
})
})
	})
	return genericEq
}


