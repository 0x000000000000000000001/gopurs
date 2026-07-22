package VendoredVariant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var VariantFRep gopurs_runtime.Value
var once_VariantFRep sync.Once
func Get_VariantFRep() gopurs_runtime.Value {
	once_VariantFRep.Do(func() {
		VariantFRep = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return VariantFRep
}

var on gopurs_runtime.Value
var once_on sync.Once
func Get_on() gopurs_runtime.Value {
	once_on.Do(func() {
		on = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_0 := gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), r_5)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), v_6_0.PtrVal.(map[string]gopurs_runtime.Value)["type"]), gopurs_runtime.Apply(dictIsSymbol_1.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], p_2))).IntVal != 0 {
__t1 = gopurs_runtime.Apply(f_3, v_6_0.PtrVal.(map[string]gopurs_runtime.Value)["value"])
} else {
__t1 = gopurs_runtime.Apply(g_4, gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), r_5))
}
return __t1
})
})
})
})
})
})
	})
	return on
}

var case_ gopurs_runtime.Value
var once_case_ sync.Once
func Get_case_() gopurs_runtime.Value {
	once_case_.Do(func() {
		case_ = gopurs_runtime.Func(func(r_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("failure on ")), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), r_0).PtrVal.(map[string]gopurs_runtime.Value)["type"])
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), __local_var_1_0)
}))
})
	})
	return case_
}


