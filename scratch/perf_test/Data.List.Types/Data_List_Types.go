package Data_List_Types

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
)

var Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		Nil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
	})
	return Nil
}

var Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": value0, "value1": value1})
})
})
	})
	return Cons
}

var NonEmptyList gopurs_runtime.Value
var once_NonEmptyList sync.Once
func Get_NonEmptyList() gopurs_runtime.Value {
	once_NonEmptyList.Do(func() {
		NonEmptyList = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NonEmptyList
}

var toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		toList = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
	})
	return toList
}

var newtypeNonEmptyList gopurs_runtime.Value
var once_newtypeNonEmptyList sync.Once
func Get_newtypeNonEmptyList() gopurs_runtime.Value {
	once_newtypeNonEmptyList.Do(func() {
		newtypeNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeNonEmptyList
}

var nelCons gopurs_runtime.Value
var once_nelCons sync.Once
func Get_nelCons() gopurs_runtime.Value {
	once_nelCons.Do(func() {
		nelCons = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": a_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
})
})
	})
	return nelCons
}

var listMap gopurs_runtime.Value
var once_listMap sync.Once
func Get_listMap() gopurs_runtime.Value {
	once_listMap.Do(func() {
		listMap = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var chunkedRevMap_1_0 gopurs_runtime.Value
chunkedRevMap_1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0).IntVal != 0)).IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(chunkedRevMap_1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3, "value1": v_2})), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
var reverseUnrolledMap_4_1 gopurs_runtime.Value
reverseUnrolledMap_4_1 = gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0).IntVal != 0).IntVal != 0)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(reverseUnrolledMap_4_1, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v3_6})})}))
} else {
__t2 = v3_6
}
return __t2
})
})
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})})
} else {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
}
__t4 = __t5
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
} else {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
}
}
__t3 = __t4
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
}
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(reverseUnrolledMap_4_1, v_2), __t3)
}
return __t6
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
	})
	return listMap
}

var functorList gopurs_runtime.Value
var once_functorList sync.Once
func Get_functorList() gopurs_runtime.Value {
	once_functorList.Do(func() {
		functorList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": Get_listMap()})
	})
	return functorList
}

var functorNonEmptyList gopurs_runtime.Value
var once_functorNonEmptyList sync.Once
func Get_functorNonEmptyList() gopurs_runtime.Value {
	once_functorNonEmptyList.Do(func() {
		functorNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(f_0, m_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_listMap(), f_0), m_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
	})
	return functorNonEmptyList
}

var foldableList gopurs_runtime.Value
var once_foldableList sync.Once
func Get_foldableList() gopurs_runtime.Value {
	once_foldableList.Do(func() {
		foldableList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, a_3), b_2)
})
})), b_1)
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = v_4
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4})), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
__local_var_4_3 := gopurs_runtime.Apply(go__3_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_4_3, x_5))
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_4 gopurs_runtime.Value
go__1_4 = gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = b_2
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_4, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, b_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})
return go__1_4
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_6 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], acc_3)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_7, gopurs_runtime.Apply(f_2, x_5))
})
})), mempty_1_6)
})
})})
	})
	return foldableList
}

var intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		intercalate = gopurs_runtime.Func(func(sep_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = b_3
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (b_3.PtrVal.(map[string]gopurs_runtime.Value)["init"]).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), b_3.PtrVal.(map[string]gopurs_runtime.Value)["acc"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), sep_0), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))})
}
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, __t2), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(true), "acc": gopurs_runtime.Str("")})), xs_1).PtrVal.(map[string]gopurs_runtime.Value)["acc"]
})
})
	})
	return intercalate
}

var foldableNonEmptyList gopurs_runtime.Value
var once_foldableNonEmptyList sync.Once
func Get_foldableNonEmptyList() gopurs_runtime.Value {
	once_foldableNonEmptyList.Do(func() {
		foldableNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_2, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_1_0, f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = b_4
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_1, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, b_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_1, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, b_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_0), b_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})})
	})
	return foldableNonEmptyList
}

var foldableWithIndexList gopurs_runtime.Value
var once_foldableWithIndexList sync.Once
func Get_foldableWithIndexList() gopurs_runtime.Value {
	once_foldableWithIndexList.Do(func() {
		foldableWithIndexList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = b_4
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), b_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1)), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
v_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Int(0), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), xs_2)
var go__5_3 gopurs_runtime.Value
go__5_3 = gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = b_6
} else {
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), b_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1)), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), b_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), b_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_1})), v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_5 gopurs_runtime.Value
go__2_5 = gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t6 = b_3
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), b_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1)), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, b_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), b_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
__local_var_3_7 := gopurs_runtime.Apply(go__2_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Int(0), "value1": acc_1}))
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_7, x_4).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_8 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexList().PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], acc_4)
__local_var_6_10 := gopurs_runtime.Apply(f_2, i_3)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_9, gopurs_runtime.Apply(__local_var_6_10, x_7))
})
})
})), mempty_1_8)
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
})})
	})
	return foldableWithIndexList
}

var foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		foldableWithIndexNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableWithIndexNonEmpty(), Get_foldableWithIndexList())
	})
	return foldableWithIndexNonEmpty
}

var foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_foldableWithIndexNonEmptyList sync.Once
func Get_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyList.Do(func() {
		foldableWithIndexNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(Get_foldableWithIndexNonEmpty().PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1))
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_4_1, x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Apply(f_2, __t2)
})), v_3)
})
})
}), "foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_3 := gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1))
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexNonEmpty().PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(__local_var_3_3, x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Apply(f_0, __t4)
})), b_1), v_2)
})
})
}), "foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1))
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexNonEmpty().PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(__local_var_3_5, x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Apply(f_0, __t6)
})), b_1), v_2)
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
})})
	})
	return foldableWithIndexNonEmptyList
}

var functorWithIndexList gopurs_runtime.Value
var once_functorWithIndexList sync.Once
func Get_functorWithIndexList() gopurs_runtime.Value {
	once_functorWithIndexList.Do(func() {
		functorWithIndexList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexList().PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, i_1), x_2), "value1": acc_3})
})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
})})
	})
	return functorWithIndexList
}

var mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		mapWithIndex = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexList().PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": i_2})), x_3), "value1": acc_4})
})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
	})
	return mapWithIndex
}

var functorWithIndexNonEmptyList gopurs_runtime.Value
var once_functorWithIndexNonEmptyList sync.Once
func Get_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList.Do(func() {
		functorWithIndexNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1))
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mapWithIndex(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, x_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Apply(fn_0, __t1)
})), v_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
})})
	})
	return functorWithIndexNonEmptyList
}

var semigroupList gopurs_runtime.Value
var once_semigroupList sync.Once
func Get_semigroupList() gopurs_runtime.Value {
	once_semigroupList.Do(func() {
		semigroupList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], Get_Cons()), ys_1), xs_0)
})
})})
	})
	return semigroupList
}

var monoidList gopurs_runtime.Value
var once_monoidList sync.Once
func Get_monoidList() gopurs_runtime.Value {
	once_monoidList.Do(func() {
		monoidList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupList()
})})
	})
	return monoidList
}

var semigroupNonEmptyList gopurs_runtime.Value
var once_semigroupNonEmptyList sync.Once
func Get_semigroupNonEmptyList() gopurs_runtime.Value {
	once_semigroupNonEmptyList.Do(func() {
		semigroupNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": as_prime_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": as_prime_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
	})
	return semigroupNonEmptyList
}

var showList gopurs_runtime.Value
var once_showList sync.Once
func Get_showList() gopurs_runtime.Value {
	once_showList.Do(func() {
		showList = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
show_1_0 := dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Str("Nil")
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_intercalate(), gopurs_runtime.Str(" : ")), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_listMap(), show_1_0), v_2))), gopurs_runtime.Str(" : Nil)")))
}
return __t1
})})
})
	})
	return showList
}

var showNonEmptyList gopurs_runtime.Value
var once_showNonEmptyList sync.Once
func Get_showNonEmptyList() gopurs_runtime.Value {
	once_showNonEmptyList.Do(func() {
		showNonEmptyList = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(NonEmptyList ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_showNonEmpty(), dictShow_0), gopurs_runtime.Apply(Get_showList(), dictShow_0)).PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showNonEmptyList
}

var traversableList gopurs_runtime.Value
var once_traversableList sync.Once
func Get_traversableList() gopurs_runtime.Value {
	once_traversableList.Do(func() {
		traversableList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = b_4
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_4})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = b_5
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_4, gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": a_8, "value1": b_7})
})
})), b_5)), gopurs_runtime.Apply(f_2, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})
__local_var_5_6 := gopurs_runtime.Apply(go__4_4, gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(__local_var_5_6, x_6))
})
})
}), "sequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableList().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
})})
	})
	return traversableList
}

var traversableNonEmptyList gopurs_runtime.Value
var once_traversableNonEmptyList sync.Once
func Get_traversableNonEmptyList() gopurs_runtime.Value {
	once_traversableNonEmptyList.Do(func() {
		traversableNonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableNonEmpty(), Get_traversableList())
	})
	return traversableNonEmptyList
}

var traversableWithIndexList gopurs_runtime.Value
var once_traversableWithIndexList sync.Once
func Get_traversableWithIndexList() gopurs_runtime.Value {
	once_traversableWithIndexList.Do(func() {
		traversableWithIndexList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = b_4
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_4})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableWithIndexList().PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(f_2, i_4)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": a_9, "value1": b_8})
})
})), acc_5)), gopurs_runtime.Apply(__local_var_6_5, x_7))
})
})
})), gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(__local_var_4_4, x_5))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexList()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexList()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableList()
})})
	})
	return traversableWithIndexList
}

var traverseWithIndex gopurs_runtime.Value
var once_traverseWithIndex sync.Once
func Get_traverseWithIndex() gopurs_runtime.Value {
	once_traverseWithIndex.Do(func() {
		traverseWithIndex = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableWithIndexNonEmpty(), Get_traversableWithIndexList()).PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"]
	})
	return traverseWithIndex
}

var traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_traversableWithIndexNonEmptyList sync.Once
func Get_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_traversableWithIndexNonEmptyList.Do(func() {
		traversableWithIndexNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex1_1_0 := gopurs_runtime.Apply(Get_traverseWithIndex(), dictApplicative_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1))
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_NonEmptyList()), gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex1_1_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_4_1, x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Apply(f_2, __t2)
})), v_3))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexNonEmptyList()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexNonEmptyList()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
})})
	})
	return traversableWithIndexNonEmptyList
}

var unfoldable1List gopurs_runtime.Value
var once_unfoldable1List sync.Once
func Get_unfoldable1List() gopurs_runtime.Value {
	once_unfoldable1List.Do(func() {
		unfoldable1List = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr1": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(source_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(f_0, source_3)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": memo_4}))
} else {
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = b_7
} else {
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_7})), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": memo_4}))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, b_1), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
})})
	})
	return unfoldable1List
}

var unfoldableList gopurs_runtime.Value
var once_unfoldableList sync.Once
func Get_unfoldableList() gopurs_runtime.Value {
	once_unfoldableList.Do(func() {
		unfoldableList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(source_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(f_0, source_3)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = b_7
} else {
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_7})), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), memo_4)
} else {
if (gopurs_runtime.Bool(v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": memo_4}))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, b_1), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
}), "Unfoldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_unfoldable1List()
})})
	})
	return unfoldableList
}

var unfoldable1NonEmptyList gopurs_runtime.Value
var once_unfoldable1NonEmptyList sync.Once
func Get_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_unfoldable1NonEmptyList.Do(func() {
		unfoldable1NonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr1": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, b_1)
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(source_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(source_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_1, gopurs_runtime.Apply(f_0, source_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, source_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": memo_5}))
} else {
var go__6_2 gopurs_runtime.Value
go__6_2 = gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = b_7
} else {
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_7})), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), memo_5)
}
return __t4
})
})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_1, __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))})
})
})})
	})
	return unfoldable1NonEmptyList
}

var foldable1NonEmptyList gopurs_runtime.Value
var once_foldable1NonEmptyList sync.Once
func Get_foldable1NonEmptyList() gopurs_runtime.Value {
	once_foldable1NonEmptyList.Do(func() {
		foldable1NonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), Get_foldableList())
	})
	return foldable1NonEmptyList
}

var extendNonEmptyList gopurs_runtime.Value
var once_extendNonEmptyList sync.Once
func Get_extendNonEmptyList() gopurs_runtime.Value {
	once_extendNonEmptyList.Do(func() {
		extendNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(f_0, v_1), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"val": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(f_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": a_2, "value1": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["acc"]})), "value1": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["val"]}), "acc": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": a_2, "value1": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["acc"]})})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"val": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "acc": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["val"]})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
})})
	})
	return extendNonEmptyList
}

var extendList gopurs_runtime.Value
var once_extendList sync.Once
func Get_extendList() gopurs_runtime.Value {
	once_extendList.Do(func() {
		extendList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(v_0, v1_1), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(a_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"val": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(v_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": a_prime_2, "value1": v2_3.PtrVal.(map[string]gopurs_runtime.Value)["acc"]})), "value1": v2_3.PtrVal.(map[string]gopurs_runtime.Value)["val"]}), "acc": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": a_prime_2, "value1": v2_3.PtrVal.(map[string]gopurs_runtime.Value)["acc"]})})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"val": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "acc": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["val"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
})})
	})
	return extendList
}

var eq1List gopurs_runtime.Value
var once_eq1List sync.Once
func Get_eq1List() gopurs_runtime.Value {
	once_eq1List.Do(func() {
		eq1List = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_6.IntVal == 0)).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil").IntVal != 0 && v2_6.IntVal != 0)
} else {
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), v2_6), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))).IntVal != 0).IntVal != 0)
}
}
return __t1
})
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, xs_1), ys_2), gopurs_runtime.Bool(true))
})
})
})})
	})
	return eq1List
}

var eq1NonEmptyList gopurs_runtime.Value
var once_eq1NonEmptyList sync.Once
func Get_eq1NonEmptyList() gopurs_runtime.Value {
	once_eq1NonEmptyList.Do(func() {
		eq1NonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_eqNonEmpty(), Get_eq1List()), dictEq_0).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1NonEmptyList
}

var eqList gopurs_runtime.Value
var once_eqList sync.Once
func Get_eqList() gopurs_runtime.Value {
	once_eqList.Do(func() {
		eqList = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(Get_eq1List().PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_0)})
})
	})
	return eqList
}

var eqNonEmptyList gopurs_runtime.Value
var once_eqNonEmptyList sync.Once
func Get_eqNonEmptyList() gopurs_runtime.Value {
	once_eqNonEmptyList.Do(func() {
		eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_eqNonEmpty(), Get_eq1List()), dictEq_0)
})
	})
	return eqNonEmptyList
}

var ord1List gopurs_runtime.Value
var once_ord1List sync.Once
func Get_ord1List() gopurs_runtime.Value {
	once_ord1List.Do(func() {
		ord1List = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t1 = __t4
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0)).IntVal != 0 {
v2_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = v2_6_2
}
__t1 = __t3
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, xs_1), ys_2)
})
})
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1List()
})})
	})
	return ord1List
}

var ordNonEmpty gopurs_runtime.Value
var once_ordNonEmpty sync.Once
func Get_ordNonEmpty() gopurs_runtime.Value {
	once_ordNonEmpty.Do(func() {
		ordNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_ordNonEmpty(), Get_ord1List())
	})
	return ordNonEmpty
}

var ord1NonEmptyList gopurs_runtime.Value
var once_ord1NonEmptyList sync.Once
func Get_ord1NonEmptyList() gopurs_runtime.Value {
	once_ord1NonEmptyList.Do(func() {
		ord1NonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_ord1NonEmpty(), Get_ord1List())
	})
	return ord1NonEmptyList
}

var ordList gopurs_runtime.Value
var once_ordList sync.Once
func Get_ordList() gopurs_runtime.Value {
	once_ordList.Do(func() {
		ordList = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t1 = __t4
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0)).IntVal != 0 {
v2_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = v2_6_2
}
__t1 = __t3
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, xs_1), ys_2)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(Get_eq1List().PtrVal.(map[string]gopurs_runtime.Value)["eq1"], gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))})
})})
})
	})
	return ordList
}

var ordNonEmptyList gopurs_runtime.Value
var once_ordNonEmptyList sync.Once
func Get_ordNonEmptyList() gopurs_runtime.Value {
	once_ordNonEmptyList.Do(func() {
		ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd_0)
})
	})
	return ordNonEmptyList
}

var comonadNonEmptyList gopurs_runtime.Value
var once_comonadNonEmptyList sync.Once
func Get_comonadNonEmptyList() gopurs_runtime.Value {
	once_comonadNonEmptyList.Do(func() {
		comonadNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
}), "Extend0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendNonEmptyList()
})})
	})
	return comonadNonEmptyList
}

var applyList gopurs_runtime.Value
var once_applyList sync.Once
func Get_applyList() gopurs_runtime.Value {
	once_applyList.Do(func() {
		applyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], Get_Cons()), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applyList().PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_1)), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_listMap(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_1))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
})})
	})
	return applyList
}

var applyNonEmptyList gopurs_runtime.Value
var once_applyNonEmptyList sync.Once
func Get_applyNonEmptyList() gopurs_runtime.Value {
	once_applyNonEmptyList.Do(func() {
		applyNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], Get_Cons()), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applyList().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applyList().PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})))})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
})})
	})
	return applyNonEmptyList
}

var bindList gopurs_runtime.Value
var once_bindList sync.Once
func Get_bindList() gopurs_runtime.Value {
	once_bindList.Do(func() {
		bindList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], Get_Cons()), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bindList().PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_1)), gopurs_runtime.Apply(v1_1, v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
})})
	})
	return bindList
}

var bindNonEmptyList gopurs_runtime.Value
var once_bindNonEmptyList sync.Once
func Get_bindNonEmptyList() gopurs_runtime.Value {
	once_bindNonEmptyList.Do(func() {
		bindNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(f_1, v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], Get_Cons()), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bindList().PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))), v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
})})
	})
	return bindNonEmptyList
}

var applicativeList gopurs_runtime.Value
var once_applicativeList sync.Once
func Get_applicativeList() gopurs_runtime.Value {
	once_applicativeList.Do(func() {
		applicativeList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": a_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
})})
	})
	return applicativeList
}

var monadList gopurs_runtime.Value
var once_monadList sync.Once
func Get_monadList() gopurs_runtime.Value {
	once_monadList.Do(func() {
		monadList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindList()
})})
	})
	return monadList
}

var altNonEmptyList gopurs_runtime.Value
var once_altNonEmptyList sync.Once
func Get_altNonEmptyList() gopurs_runtime.Value {
	once_altNonEmptyList.Do(func() {
		altNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": Get_semigroupNonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["append"], "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
})})
	})
	return altNonEmptyList
}

var altList gopurs_runtime.Value
var once_altList sync.Once
func Get_altList() gopurs_runtime.Value {
	once_altList.Do(func() {
		altList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": Get_semigroupList().PtrVal.(map[string]gopurs_runtime.Value)["append"], "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
})})
	})
	return altList
}

var plusList gopurs_runtime.Value
var once_plusList sync.Once
func Get_plusList() gopurs_runtime.Value {
	once_plusList.Do(func() {
		plusList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altList()
})})
	})
	return plusList
}

var alternativeList gopurs_runtime.Value
var once_alternativeList sync.Once
func Get_alternativeList() gopurs_runtime.Value {
	once_alternativeList.Do(func() {
		alternativeList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusList()
})})
	})
	return alternativeList
}

var monadPlusList gopurs_runtime.Value
var once_monadPlusList sync.Once
func Get_monadPlusList() gopurs_runtime.Value {
	once_monadPlusList.Do(func() {
		monadPlusList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadList()
}), "Alternative1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_alternativeList()
})})
	})
	return monadPlusList
}

var applicativeNonEmptyList gopurs_runtime.Value
var once_applicativeNonEmptyList sync.Once
func Get_applicativeNonEmptyList() gopurs_runtime.Value {
	once_applicativeNonEmptyList.Do(func() {
		applicativeNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": x_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
})})
	})
	return applicativeNonEmptyList
}

var monadNonEmptyList gopurs_runtime.Value
var once_monadNonEmptyList sync.Once
func Get_monadNonEmptyList() gopurs_runtime.Value {
	once_monadNonEmptyList.Do(func() {
		monadNonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeNonEmptyList()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindNonEmptyList()
})})
	})
	return monadNonEmptyList
}

var traversable1NonEmptyList gopurs_runtime.Value
var once_traversable1NonEmptyList sync.Once
func Get_traversable1NonEmptyList() gopurs_runtime.Value {
	once_traversable1NonEmptyList.Do(func() {
		traversable1NonEmptyList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_3 gopurs_runtime.Value
go__4_3 = gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = b_5
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_3, gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": a_8, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": b_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
})
})), b_5)), gopurs_runtime.Apply(f_2, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_1 gopurs_runtime.Value
go__5_1 = gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = b_6
} else {
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": b_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": b_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})), gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_3, gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_applicativeNonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["pure"]), gopurs_runtime.Apply(f_2, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "sequence1": gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversable1NonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["traverse1"], dictApply_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Foldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1NonEmptyList()
}), "Traversable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
})})
	})
	return traversable1NonEmptyList
}


