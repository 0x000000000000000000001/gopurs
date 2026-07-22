package Data_Enum

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		guard = gopurs_runtime.Apply(pkg_Control_Alternative.Get_guard(), pkg_Data_Maybe.Get_alternativeMaybe())
	})
	return guard
}

var Cardinality gopurs_runtime.Value
var once_Cardinality sync.Once
func Get_Cardinality() gopurs_runtime.Value {
	once_Cardinality.Do(func() {
		Cardinality = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Cardinality
}

var toEnum gopurs_runtime.Value
var once_toEnum sync.Once
func Get_toEnum() gopurs_runtime.Value {
	once_toEnum.Do(func() {
		toEnum = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["toEnum"]
})
	})
	return toEnum
}

var succ gopurs_runtime.Value
var once_succ sync.Once
func Get_succ() gopurs_runtime.Value {
	once_succ.Do(func() {
		succ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["succ"]
})
	})
	return succ
}

var upFromIncluding gopurs_runtime.Value
var once_upFromIncluding sync.Once
func Get_upFromIncluding() gopurs_runtime.Value {
	once_upFromIncluding.Do(func() {
		upFromIncluding = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictUnfoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictUnfoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_2, "value1": gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["succ"], x_2)})
}))
})
})
	})
	return upFromIncluding
}

var showCardinality gopurs_runtime.Value
var once_showCardinality sync.Once
func Get_showCardinality() gopurs_runtime.Value {
	once_showCardinality.Do(func() {
		showCardinality = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Cardinality ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0)), gopurs_runtime.Str(")")))
})})
	})
	return showCardinality
}

var pred gopurs_runtime.Value
var once_pred sync.Once
func Get_pred() gopurs_runtime.Value {
	once_pred.Do(func() {
		pred = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"]
})
	})
	return pred
}

var ordCardinality gopurs_runtime.Value
var once_ordCardinality sync.Once
func Get_ordCardinality() gopurs_runtime.Value {
	once_ordCardinality.Do(func() {
		ordCardinality = pkg_Data_Ord.Get_ordInt()
	})
	return ordCardinality
}

var newtypeCardinality gopurs_runtime.Value
var once_newtypeCardinality sync.Once
func Get_newtypeCardinality() gopurs_runtime.Value {
	once_newtypeCardinality.Do(func() {
		newtypeCardinality = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeCardinality
}

var fromEnum gopurs_runtime.Value
var once_fromEnum sync.Once
func Get_fromEnum() gopurs_runtime.Value {
	once_fromEnum.Do(func() {
		fromEnum = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"]
})
	})
	return fromEnum
}

var toEnumWithDefaults gopurs_runtime.Value
var once_toEnumWithDefaults sync.Once
func Get_toEnumWithDefaults() gopurs_runtime.Value {
	once_toEnumWithDefaults.Do(func() {
		toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["Bounded0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], x_4)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4), gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"], bottom2_1_0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = low_2
} else {
__t3 = high_3
}
__t2 = __t3
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
})
})
	})
	return toEnumWithDefaults
}

var eqCardinality gopurs_runtime.Value
var once_eqCardinality sync.Once
func Get_eqCardinality() gopurs_runtime.Value {
	once_eqCardinality.Do(func() {
		eqCardinality = pkg_Data_Eq.Get_eqInt()
	})
	return eqCardinality
}

var enumUnit gopurs_runtime.Value
var once_enumUnit sync.Once
func Get_enumUnit() gopurs_runtime.Value {
	once_enumUnit.Do(func() {
		enumUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}), "pred": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordUnit()
})})
	})
	return enumUnit
}

var enumTuple gopurs_runtime.Value
var once_enumTuple sync.Once
func Get_enumTuple() gopurs_runtime.Value {
	once_enumTuple.Do(func() {
		enumTuple = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordTuple_1_0 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_ordTuple(), gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
Bounded0_3_1 := gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["Bounded0"], gopurs_runtime.Value{})
bottom2_4_2 := Bounded0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
Enum1_5_3 := gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["Enum1"], gopurs_runtime.Value{})
top2_6_4 := Bounded0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["top"]
ordTuple1_7_5 := gopurs_runtime.Apply(ordTuple_1_0, gopurs_runtime.Apply(Enum1_5_3.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_6 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["succ"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_9_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": bottom2_4_2})})
} else {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__local_var_10_7 := __t8
__local_var_11_9 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
__local_var_12_10 := gopurs_runtime.Apply(Enum1_5_3.PtrVal.(map[string]gopurs_runtime.Value)["succ"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t11 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_12_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t11 = __local_var_10_7
} else {
if (gopurs_runtime.Bool(__local_var_12_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_11_9, __local_var_12_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t11
}), "pred": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_12 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t14 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_12.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_9_12.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": top2_6_4})})
} else {
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__local_var_10_13 := __t14
__local_var_11_15 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
__local_var_12_16 := gopurs_runtime.Apply(Enum1_5_3.PtrVal.(map[string]gopurs_runtime.Value)["pred"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t17 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_12_16.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t17 = __local_var_10_13
} else {
if (gopurs_runtime.Bool(__local_var_12_16.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t17 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_11_15, __local_var_12_16.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t17
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple1_7_5
})})
})
})
	})
	return enumTuple
}

var enumOrdering gopurs_runtime.Value
var once_enumOrdering sync.Once
func Get_enumOrdering() gopurs_runtime.Value {
	once_enumOrdering.Do(func() {
		enumOrdering = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
}), "pred": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordOrdering()
})})
	})
	return enumOrdering
}

var enumMaybe gopurs_runtime.Value
var once_enumMaybe sync.Once
func Get_enumMaybe() gopurs_runtime.Value {
	once_enumMaybe.Do(func() {
		enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["Bounded0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
Enum1_2_1 := gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["Enum1"], gopurs_runtime.Value{})
__local_var_3_2 := gopurs_runtime.Apply(Enum1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{})
__local_var_4_3 := gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eqMaybe1_5_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Bool(y_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")
} else {
__t6 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t6
})
})})
ordMaybe_5_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(y_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t7 = __t8
} else {
if (gopurs_runtime.Bool(y_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t7
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_5_5
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": bottom2_1_0})})
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_10 := gopurs_runtime.Apply(Enum1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["succ"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t11 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_7_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_7_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t9 = __t11
} else {
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t9
}), "pred": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(Enum1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["pred"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t12
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe_5_4
})})
})
	})
	return enumMaybe
}

var enumInt gopurs_runtime.Value
var once_enumInt sync.Once
func Get_enumInt() gopurs_runtime.Value {
	once_enumInt.Do(func() {
		enumInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), pkg_Data_Bounded.Get_topInt()).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), n_0), gopurs_runtime.Int(1))})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}), "pred": gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), pkg_Data_Bounded.Get_bottomInt()).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), n_0), gopurs_runtime.Int(1))})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return enumInt
}

var enumFromTo gopurs_runtime.Value
var once_enumFromTo sync.Once
func Get_enumFromTo() gopurs_runtime.Value {
	once_enumFromTo.Do(func() {
		enumFromTo = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
Ord0_1_0 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Ord0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_3), v1_4)).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Unfoldable1.Get_replicate1(), dictUnfoldable1_2), gopurs_runtime.Int(1)), v_3)
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(Ord0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_3), v1_4).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable1_2.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["succ"], a_5)
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(Get_guard(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(Ord0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_6_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_4).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_6_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t6 = __t7
} else {
if (gopurs_runtime.Bool(__local_var_6_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_5, "value1": __t6})
})), v_3)
} else {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable1_2.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"], a_5)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(Get_guard(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(Ord0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_4).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t2 = __t3
} else {
if (gopurs_runtime.Bool(__local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_5, "value1": __t2})
})), v_3)
}
}
return __t4
})
})
})
})
	})
	return enumFromTo
}

var enumFromThenTo gopurs_runtime.Value
var once_enumFromThenTo sync.Once
func Get_enumFromThenTo() gopurs_runtime.Value {
	once_enumFromThenTo.Do(func() {
		enumFromThenTo = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_7_0 := gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"], a_4)
__local_var_8_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"], b_5)), a_prime_7_0)
__local_var_9_4 := gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"], c_6)
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_1 := gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], x_8)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = __local_var_9_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t2
})), gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(e_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], e_10), __local_var_9_4).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": e_10, "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), e_10), __local_var_8_3)})})
} else {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t5
})), a_prime_7_0))
})
})
})
}))
})
})
})
	})
	return enumFromThenTo
}

var enumEither gopurs_runtime.Value
var once_enumEither sync.Once
func Get_enumEither() gopurs_runtime.Value {
	once_enumEither.Do(func() {
		enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
Enum1_1_0 := gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["Enum1"], gopurs_runtime.Value{})
top2_2_1 := gopurs_runtime.Apply(dictBoundedEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["Bounded0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["top"]
ordEither_3_2 := gopurs_runtime.Apply(pkg_Data_Either.Get_ordEither(), gopurs_runtime.Apply(Enum1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictBoundedEnum1_4 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_5_3 := gopurs_runtime.Apply(dictBoundedEnum1_4.PtrVal.(map[string]gopurs_runtime.Value)["Bounded0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
Enum11_6_4 := gopurs_runtime.Apply(dictBoundedEnum1_4.PtrVal.(map[string]gopurs_runtime.Value)["Enum1"], gopurs_runtime.Value{})
ordEither1_7_5 := gopurs_runtime.Apply(ordEither_3_2, gopurs_runtime.Apply(Enum11_6_4.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__local_var_9_9 := gopurs_runtime.Apply(Enum1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["succ"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t10 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": bottom2_5_3})})
} else {
if (gopurs_runtime.Bool(__local_var_9_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": __local_var_9_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t6 = __t10
} else {
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__local_var_9_7 := gopurs_runtime.Apply(Enum11_6_4.PtrVal.(map[string]gopurs_runtime.Value)["succ"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(__local_var_9_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": __local_var_9_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t6 = __t8
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
}), "pred": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__local_var_9_14 := gopurs_runtime.Apply(Enum1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t15 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_14.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t15 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(__local_var_9_14.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t15 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": __local_var_9_14.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t11 = __t15
} else {
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__local_var_9_12 := gopurs_runtime.Apply(Enum11_6_4.PtrVal.(map[string]gopurs_runtime.Value)["pred"], v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t13 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_12.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": top2_2_1})})
} else {
if (gopurs_runtime.Bool(__local_var_9_12.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": __local_var_9_12.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t11 = __t13
} else {
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t11
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither1_7_5
})})
})
})
	})
	return enumEither
}

var enumBoolean gopurs_runtime.Value
var once_enumBoolean sync.Once
func Get_enumBoolean() gopurs_runtime.Value {
	once_enumBoolean.Do(func() {
		enumBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Bool(true)})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}), "pred": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (v_0).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Bool(false)})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordBoolean()
})})
	})
	return enumBoolean
}

var downFromIncluding gopurs_runtime.Value
var once_downFromIncluding sync.Once
func Get_downFromIncluding() gopurs_runtime.Value {
	once_downFromIncluding.Do(func() {
		downFromIncluding = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictUnfoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictUnfoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_2, "value1": gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"], x_2)})
}))
})
})
	})
	return downFromIncluding
}

var downFrom gopurs_runtime.Value
var once_downFrom sync.Once
func Get_downFrom() gopurs_runtime.Value {
	once_downFrom.Do(func() {
		downFrom = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictUnfoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictUnfoldable_1.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"], x_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}))
})
})
	})
	return downFrom
}

var upFrom gopurs_runtime.Value
var once_upFrom sync.Once
func Get_upFrom() gopurs_runtime.Value {
	once_upFrom.Do(func() {
		upFrom = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictUnfoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictUnfoldable_1.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["succ"], x_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}))
})
})
	})
	return upFrom
}

var defaultToEnum gopurs_runtime.Value
var once_defaultToEnum sync.Once
func Get_defaultToEnum() gopurs_runtime.Value {
	once_defaultToEnum.Do(func() {
		defaultToEnum = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), i_5), gopurs_runtime.Int(0))).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_6})
} else {
v_7_2 := gopurs_runtime.Apply(dictEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["succ"], x_6)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_7_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_1, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), i_5), gopurs_runtime.Int(1))), v_7_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_7_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t4 = __t3
}
return __t4
})
})
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], i_prime_3), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_1, i_prime_3), bottom2_1_0)
}
return __t5
})
})
})
	})
	return defaultToEnum
}

var defaultSucc gopurs_runtime.Value
var once_defaultSucc sync.Once
func Get_defaultSucc() gopurs_runtime.Value {
	once_defaultSucc.Do(func() {
		defaultSucc = gopurs_runtime.Func(func(toEnum_prime_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fromEnum_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(fromEnum_prime_1, a_2)), gopurs_runtime.Int(1)))
})
})
})
	})
	return defaultSucc
}

var defaultPred gopurs_runtime.Value
var once_defaultPred sync.Once
func Get_defaultPred() gopurs_runtime.Value {
	once_defaultPred.Do(func() {
		defaultPred = gopurs_runtime.Func(func(toEnum_prime_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fromEnum_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(fromEnum_prime_1, a_2)), gopurs_runtime.Int(1)))
})
})
})
	})
	return defaultPred
}

var defaultFromEnum gopurs_runtime.Value
var once_defaultFromEnum sync.Once
func Get_defaultFromEnum() gopurs_runtime.Value {
	once_defaultFromEnum.Do(func() {
		defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.Apply(dictEnum_0.PtrVal.(map[string]gopurs_runtime.Value)["pred"], x_3)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), i_2), gopurs_runtime.Int(1))), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = i_2
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Int(0))
})
	})
	return defaultFromEnum
}

var defaultCardinality gopurs_runtime.Value
var once_defaultCardinality sync.Once
func Get_defaultCardinality() gopurs_runtime.Value {
	once_defaultCardinality.Do(func() {
		defaultCardinality = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply(dictEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["succ"], x_5)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_1, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), i_4), gopurs_runtime.Int(1))), v_6_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = i_4
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_1, gopurs_runtime.Int(1)), bottom2_1_0)
})
})
	})
	return defaultCardinality
}

var charToEnum gopurs_runtime.Value
var once_charToEnum sync.Once
func Get_charToEnum() gopurs_runtime.Value {
	once_charToEnum.Do(func() {
		charToEnum = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_bottomChar())).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_topChar())).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(Get_fromCharCode(), v_0)})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
})
	})
	return charToEnum
}

var enumChar gopurs_runtime.Value
var once_enumChar sync.Once
func Get_enumChar() gopurs_runtime.Value {
	once_enumChar.Do(func() {
		enumChar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_defaultSucc(), Get_charToEnum()), Get_toCharCode()), "pred": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_defaultPred(), Get_charToEnum()), Get_toCharCode()), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordChar()
})})
	})
	return enumChar
}

var cardinality gopurs_runtime.Value
var once_cardinality sync.Once
func Get_cardinality() gopurs_runtime.Value {
	once_cardinality.Do(func() {
		cardinality = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["cardinality"]
})
	})
	return cardinality
}

var boundedEnumUnit gopurs_runtime.Value
var once_boundedEnumUnit sync.Once
func Get_boundedEnumUnit() gopurs_runtime.Value {
	once_boundedEnumUnit.Do(func() {
		boundedEnumUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Int(1), "toEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": pkg_Data_Unit.Get_unit()})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedUnit()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumUnit()
})})
	})
	return boundedEnumUnit
}

var boundedEnumOrdering gopurs_runtime.Value
var once_boundedEnumOrdering sync.Once
func Get_boundedEnumOrdering() gopurs_runtime.Value {
	once_boundedEnumOrdering.Do(func() {
		boundedEnumOrdering = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Int(3), "toEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})})
} else {
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})})
} else {
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
}
}
return __t0
}), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedOrdering()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumOrdering()
})})
	})
	return boundedEnumOrdering
}

var boundedEnumChar gopurs_runtime.Value
var once_boundedEnumChar sync.Once
func Get_boundedEnumChar() gopurs_runtime.Value {
	once_boundedEnumChar.Do(func() {
		boundedEnumChar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_topChar())), gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_bottomChar())), "toEnum": Get_charToEnum(), "fromEnum": Get_toCharCode(), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedChar()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumChar()
})})
	})
	return boundedEnumChar
}

var boundedEnumBoolean gopurs_runtime.Value
var once_boundedEnumBoolean sync.Once
func Get_boundedEnumBoolean() gopurs_runtime.Value {
	once_boundedEnumBoolean.Do(func() {
		boundedEnumBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Int(2), "toEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Bool(false)})
} else {
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Bool(true)})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
}
return __t0
}), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == 0)).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
} else {
if (v_0).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedBoolean()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumBoolean()
})})
	})
	return boundedEnumBoolean
}

func Get_fromCharCode() gopurs_runtime.Value {
	return FromCharCode
}

func Get_toCharCode() gopurs_runtime.Value {
	return ToCharCode
}
