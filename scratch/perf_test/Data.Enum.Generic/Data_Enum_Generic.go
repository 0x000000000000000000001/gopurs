package Data_Enum_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var genericToEnum_prime gopurs_runtime.Value
var once_genericToEnum_prime sync.Once
func Get_genericToEnum_prime() gopurs_runtime.Value {
	once_genericToEnum_prime.Do(func() {
		genericToEnum_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericToEnum'"]
})
	})
	return genericToEnum_prime
}

var genericToEnum gopurs_runtime.Value
var once_genericToEnum sync.Once
func Get_genericToEnum() gopurs_runtime.Value {
	once_genericToEnum.Do(func() {
		genericToEnum = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericBoundedEnum_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(dictGenericBoundedEnum_1.PtrVal.(map[string]gopurs_runtime.Value)["genericToEnum'"], x_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
})
	})
	return genericToEnum
}

var genericSucc_prime gopurs_runtime.Value
var once_genericSucc_prime sync.Once
func Get_genericSucc_prime() gopurs_runtime.Value {
	once_genericSucc_prime.Do(func() {
		genericSucc_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericSucc'"]
})
	})
	return genericSucc_prime
}

var genericSucc gopurs_runtime.Value
var once_genericSucc sync.Once
func Get_genericSucc() gopurs_runtime.Value {
	once_genericSucc.Do(func() {
		genericSucc = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericEnum_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(dictGenericEnum_1.PtrVal.(map[string]gopurs_runtime.Value)["genericSucc'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
})
	})
	return genericSucc
}

var genericPred_prime gopurs_runtime.Value
var once_genericPred_prime sync.Once
func Get_genericPred_prime() gopurs_runtime.Value {
	once_genericPred_prime.Do(func() {
		genericPred_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericPred'"]
})
	})
	return genericPred_prime
}

var genericPred gopurs_runtime.Value
var once_genericPred sync.Once
func Get_genericPred() gopurs_runtime.Value {
	once_genericPred.Do(func() {
		genericPred = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericEnum_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(dictGenericEnum_1.PtrVal.(map[string]gopurs_runtime.Value)["genericPred'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["to"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
})
	})
	return genericPred
}

var genericFromEnum_prime gopurs_runtime.Value
var once_genericFromEnum_prime sync.Once
func Get_genericFromEnum_prime() gopurs_runtime.Value {
	once_genericFromEnum_prime.Do(func() {
		genericFromEnum_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericFromEnum'"]
})
	})
	return genericFromEnum_prime
}

var genericFromEnum gopurs_runtime.Value
var once_genericFromEnum sync.Once
func Get_genericFromEnum() gopurs_runtime.Value {
	once_genericFromEnum.Do(func() {
		genericFromEnum = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericBoundedEnum_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGenericBoundedEnum_1.PtrVal.(map[string]gopurs_runtime.Value)["genericFromEnum'"], gopurs_runtime.Apply(dictGeneric_0.PtrVal.(map[string]gopurs_runtime.Value)["from"], x_2))
})
})
})
	})
	return genericFromEnum
}

var genericEnumSum gopurs_runtime.Value
var once_genericEnumSum sync.Once
func Get_genericEnumSum() gopurs_runtime.Value {
	once_genericEnumSum.Do(func() {
		genericEnumSum = gopurs_runtime.Func(func(dictGenericEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericTop_1 gopurs_runtime.Value) gopurs_runtime.Value {
genericTop_prime_2_0 := dictGenericTop_1.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"]
return gopurs_runtime.Func(func(dictGenericEnum1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericBottom_4 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_5_1 := dictGenericBottom_4.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericPred'": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__local_var_7_5 := gopurs_runtime.Apply(dictGenericEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericPred'"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_7_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inl"), "value0": __local_var_7_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t2 = __t6
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
v1_7_3 := gopurs_runtime.Apply(dictGenericEnum1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericPred'"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inl"), "value0": genericTop_prime_2_0})})
} else {
if (gopurs_runtime.Bool(v1_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inr"), "value0": v1_7_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t2 = __t4
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
}), "genericSucc'": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
v1_7_10 := gopurs_runtime.Apply(dictGenericEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericSucc'"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t11 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_7_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inr"), "value0": genericBottom_prime_5_1})})
} else {
if (gopurs_runtime.Bool(v1_7_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inl"), "value0": v1_7_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t7 = __t11
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
__local_var_7_8 := gopurs_runtime.Apply(dictGenericEnum1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericSucc'"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t9 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_7_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inr"), "value0": __local_var_7_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t7 = __t9
} else {
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t7
})})
})
})
})
})
	})
	return genericEnumSum
}

var genericEnumProduct gopurs_runtime.Value
var once_genericEnumProduct sync.Once
func Get_genericEnumProduct() gopurs_runtime.Value {
	once_genericEnumProduct.Do(func() {
		genericEnumProduct = gopurs_runtime.Func(func(dictGenericEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericTop_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericBottom_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericEnum1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericTop1_4 gopurs_runtime.Value) gopurs_runtime.Value {
genericTop_prime_5_0 := dictGenericTop1_4.PtrVal.(map[string]gopurs_runtime.Value)["genericTop'"]
return gopurs_runtime.Func(func(dictGenericBottom1_6 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_7_1 := dictGenericBottom1_6.PtrVal.(map[string]gopurs_runtime.Value)["genericBottom'"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericPred'": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
v1_9_2 := gopurs_runtime.Apply(dictGenericEnum1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericPred'"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_9_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_9_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
if (gopurs_runtime.Bool(v1_9_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__local_var_10_4 := gopurs_runtime.Apply(dictGenericEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericPred'"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_10_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": __local_var_10_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": genericTop_prime_5_0})})
} else {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t3 = __t5
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
}), "genericSucc'": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
v1_9_6 := gopurs_runtime.Apply(dictGenericEnum1_3.PtrVal.(map[string]gopurs_runtime.Value)["genericSucc'"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_9_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_9_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
if (gopurs_runtime.Bool(v1_9_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__local_var_10_8 := gopurs_runtime.Apply(dictGenericEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericSucc'"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t9 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_10_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": __local_var_10_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": genericBottom_prime_7_1})})
} else {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t7 = __t9
} else {
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t7
})})
})
})
})
})
})
})
	})
	return genericEnumProduct
}

var genericEnumNoArguments gopurs_runtime.Value
var once_genericEnumNoArguments sync.Once
func Get_genericEnumNoArguments() gopurs_runtime.Value {
	once_genericEnumNoArguments.Do(func() {
		genericEnumNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericPred'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}), "genericSucc'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
})})
	})
	return genericEnumNoArguments
}

var genericEnumConstructor gopurs_runtime.Value
var once_genericEnumConstructor sync.Once
func Get_genericEnumConstructor() gopurs_runtime.Value {
	once_genericEnumConstructor.Do(func() {
		genericEnumConstructor = gopurs_runtime.Func(func(dictGenericEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericPred'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictGenericEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericPred'"], v_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}), "genericSucc'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply(dictGenericEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericSucc'"], v_1)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t3
})})
})
	})
	return genericEnumConstructor
}

var genericEnumArgument gopurs_runtime.Value
var once_genericEnumArgument sync.Once
func Get_genericEnumArgument() gopurs_runtime.Value {
	once_genericEnumArgument.Do(func() {
		genericEnumArgument = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericPred'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"], v_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}), "genericSucc'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["succ"], v_1)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t3
})})
})
	})
	return genericEnumArgument
}

var genericCardinality_prime gopurs_runtime.Value
var once_genericCardinality_prime sync.Once
func Get_genericCardinality_prime() gopurs_runtime.Value {
	once_genericCardinality_prime.Do(func() {
		genericCardinality_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCardinality'"]
})
	})
	return genericCardinality_prime
}

var genericCardinality gopurs_runtime.Value
var once_genericCardinality sync.Once
func Get_genericCardinality() gopurs_runtime.Value {
	once_genericCardinality.Do(func() {
		genericCardinality = gopurs_runtime.Func(func(dictGeneric_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictGenericBoundedEnum_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), dictGenericBoundedEnum_1.PtrVal.(map[string]gopurs_runtime.Value)["genericCardinality'"])
})
})
	})
	return genericCardinality
}

var genericBoundedEnumSum gopurs_runtime.Value
var once_genericBoundedEnumSum sync.Once
func Get_genericBoundedEnumSum() gopurs_runtime.Value {
	once_genericBoundedEnumSum.Do(func() {
		genericBoundedEnumSum = gopurs_runtime.Func(func(dictGenericBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericCardinality_prime1_1_0 := dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCardinality'"]
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCardinality'": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), genericCardinality_prime1_1_0)), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), dictGenericBoundedEnum1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericCardinality'"])), "genericToEnum'": gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_3), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_3), genericCardinality_prime1_1_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT"))).IntVal != 0 {
__local_var_4_4 := gopurs_runtime.Apply(dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericToEnum'"], n_3)
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_4_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inl"), "value0": __local_var_4_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t3 = __t5
} else {
__local_var_4_1 := gopurs_runtime.Apply(dictGenericBoundedEnum1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericToEnum'"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), n_3), genericCardinality_prime1_1_0))
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inr"), "value0": __local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t3 = __t2
}
return __t3
}), "genericFromEnum'": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericFromEnum'"], v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(dictGenericBoundedEnum1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericFromEnum'"], v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), genericCardinality_prime1_1_0))
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})})
})
})
	})
	return genericBoundedEnumSum
}

var genericBoundedEnumProduct gopurs_runtime.Value
var once_genericBoundedEnumProduct sync.Once
func Get_genericBoundedEnumProduct() gopurs_runtime.Value {
	once_genericBoundedEnumProduct.Do(func() {
		genericBoundedEnumProduct = gopurs_runtime.Func(func(dictGenericBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericCardinality_prime1_1_0 := dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCardinality'"]
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
genericCardinality_prime2_3_1 := dictGenericBoundedEnum1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericCardinality'"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCardinality'": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intMul(), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), genericCardinality_prime1_1_0)), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), genericCardinality_prime2_3_1)), "genericToEnum'": gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericToEnum'"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intDiv(), n_4), genericCardinality_prime2_3_1))
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_5_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_4 := gopurs_runtime.Apply(dictGenericBoundedEnum1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericToEnum'"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), n_4), genericCardinality_prime2_3_1))
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": __local_var_5_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_6_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t3 = __t5
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t3
}), "genericFromEnum'": gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intMul(), gopurs_runtime.Apply(dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericFromEnum'"], v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), genericCardinality_prime2_3_1)), gopurs_runtime.Apply(dictGenericBoundedEnum1_2.PtrVal.(map[string]gopurs_runtime.Value)["genericFromEnum'"], v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})})
})
})
	})
	return genericBoundedEnumProduct
}

var genericBoundedEnumNoArguments gopurs_runtime.Value
var once_genericBoundedEnumNoArguments sync.Once
func Get_genericBoundedEnumNoArguments() gopurs_runtime.Value {
	once_genericBoundedEnumNoArguments.Do(func() {
		genericBoundedEnumNoArguments = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCardinality'": gopurs_runtime.Int(1), "genericToEnum'": gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), i_0), gopurs_runtime.Int(0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}), "genericFromEnum'": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
})})
	})
	return genericBoundedEnumNoArguments
}

var genericBoundedEnumConstructor gopurs_runtime.Value
var once_genericBoundedEnumConstructor sync.Once
func Get_genericBoundedEnumConstructor() gopurs_runtime.Value {
	once_genericBoundedEnumConstructor.Do(func() {
		genericBoundedEnumConstructor = gopurs_runtime.Func(func(dictGenericBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCardinality'": gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericCardinality'"]), "genericToEnum'": gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericToEnum'"], i_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}), "genericFromEnum'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictGenericBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["genericFromEnum'"], v_1)
})})
})
	})
	return genericBoundedEnumConstructor
}

var genericBoundedEnumArgument gopurs_runtime.Value
var once_genericBoundedEnumArgument sync.Once
func Get_genericBoundedEnumArgument() gopurs_runtime.Value {
	once_genericBoundedEnumArgument.Do(func() {
		genericBoundedEnumArgument = gopurs_runtime.Func(func(dictBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"genericCardinality'": gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["cardinality"]), "genericToEnum'": gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], i_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}), "genericFromEnum'": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"], v_1)
})})
})
	})
	return genericBoundedEnumArgument
}


