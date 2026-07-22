package Data_List

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_any(), pkg_Data_List_Types.Get_foldableList()), pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean())
	})
	return any
}

var Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		Pattern = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Pattern
}

var updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		updateAt = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1, "value1": v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_updateAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t3 = __t2
}
__t0 = __t3
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}
}()
})
})
})
	})
	return updateAt
}

var unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		unzip = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_2_1 := v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_1_0, "value1": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_2_1, "value1": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})}))
	})
	return unzip
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "tail": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return uncons
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(xs_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(xs_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": xs_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": xs_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}))
})
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return tail
}

var stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		stripPrefix = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Value{PtrVal: func(prefix_3 gopurs_runtime.Value, input_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(input_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(prefix_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], prefix_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), input_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": prefix_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "b": input_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t2 = __t3
} else {
if (gopurs_runtime.Bool(prefix_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": input_4})})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
}
__t1 = __t2
} else {
if (gopurs_runtime.Bool(prefix_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": input_4})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
}
return __t1
}}
__local_var_4_4 := gopurs_runtime.Value{PtrVal: func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})})
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": __local_var_3_0.PtrVal.(func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value)(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["a"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["b"])})
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t5 = __t6
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
}}
var go__5_7 gopurs_runtime.Value
go__5_7 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t8 = gopurs_runtime.Apply(go__5_7, __local_var_4_4.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t8 = v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t8
})
return gopurs_runtime.Apply(go__5_7, __local_var_4_4.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(__local_var_3_0.PtrVal.(func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value)(v_1, s_2)))
})
})
})
	})
	return stripPrefix
}

var span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		span = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(v_0, v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)).IntVal != 0 {
v2_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_span(), v_0), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_2_1.PtrVal.(map[string]gopurs_runtime.Value)["init"]}), "rest": v2_2_1.PtrVal.(map[string]gopurs_runtime.Value)["rest"]})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "rest": v1_1})
}
return __t0
}
}()
})
})
	})
	return span
}

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": x_1, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), xs_0)
})
})
	})
	return snoc
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": a_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
})
	})
	return singleton
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func(func(cmp_0 gopurs_runtime.Value) gopurs_runtime.Value {
var merge_1_0 gopurs_runtime.Value
merge_1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(merge_1_0, v_2), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(merge_1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_3)})
}
__t2 = __t3
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = v_2
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t2
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v1_3
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v_2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
var mergePairs_2_4 gopurs_runtime.Value
mergePairs_2_4 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0)).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(merge_1_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(mergePairs_2_4, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
} else {
__t5 = v_3
}
return __t5
})
var mergeAll_3_6 gopurs_runtime.Value
mergeAll_3_6 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil").IntVal != 0)).IntVal != 0 {
__t7 = v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t7 = gopurs_runtime.Apply(mergeAll_3_6, gopurs_runtime.Apply(mergePairs_2_4, v_4))
}
return __t7
})
var sequences_4_8 gopurs_runtime.Value
var descending_4_9 gopurs_runtime.Value
var ascending_4_10 gopurs_runtime.Value
sequences_4_8 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0)).IntVal != 0 {
var __t13 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_0, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(descending_4_9, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__local_var_6_12 := v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t13 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(ascending_4_10, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_6_12, "value1": v1_7})
})), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t11 = __t13
} else {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
}
return __t11
})
descending_4_9 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v2_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_0, v_5), v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal != 0)).IntVal != 0 {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(descending_4_9, v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5, "value1": v1_6})), v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5, "value1": v1_6}), "value1": gopurs_runtime.Apply(sequences_4_8, v2_7)})
}
return __t14
})
})
})
ascending_4_10 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v2_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqBooleanImpl(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_0, v_5), v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")), gopurs_runtime.Bool(false)).IntVal != 0)).IntVal != 0 {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(ascending_4_10, v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5, "value1": ys_8}))
})), v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t15 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(v1_6, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_5, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), "value1": gopurs_runtime.Apply(sequences_4_8, v2_7)})
}
return __t15
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_6, gopurs_runtime.Apply(sequences_4_8, x_5))
})
})
	})
	return sortBy
}

var sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		sort = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_sortBy(), compare_1_0), xs_2)
})
})
	})
	return sort
}

var showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		showPattern = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Pattern ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_showList(), dictShow_0).PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showPattern
}

var reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		reverse = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v_1
} else {
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__0_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1})), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
}()
	})
	return reverse
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v1_2), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_1)
} else {
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_1)
} else {
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__0_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1})), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v1_2), gopurs_runtime.Int(1))), v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
})
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
}()
	})
	return take
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(p_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2})), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_2)
}
return __t3
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
	})
	return takeWhile
}

var unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		unsnoc = gopurs_runtime.Func(func(lst_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"revInit": v1_3, "last": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_3}))
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, lst_0), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var go__3_5 gopurs_runtime.Value
go__3_5 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t6 = v_4
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4})), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), __local_var_2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["revInit"]), "last": __local_var_2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["last"]})})
} else {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t4
})
	})
	return unsnoc
}

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v2_6
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v2_6
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v2_6}))
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
})
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, xs_1), ys_2), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))
})
})
})
	})
	return zipWith
}

var zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		zip = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v2_5
} else {
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v2_5
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}), "value1": v2_5}))
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
})
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = v_4
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4})), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, xs_0), ys_1), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))
})
})
	})
	return zip
}

var zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		zipWithA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_traversableList().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_1 gopurs_runtime.Value
go__5_1 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = v2_8
} else {
if (gopurs_runtime.Bool(v1_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = v2_8
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_1, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v2_8}))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t2
})
})
})
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = v_7
} else {
if (gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_7})), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_1, xs_3), ys_4), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))))
})
})
})
})
	})
	return zipWithA
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func(func(start_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(end_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), start_0), end_1)).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": start_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
} else {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(step_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rest_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), s_3), e_4)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": s_3, "value1": rest_6})
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), s_3), step_5)), e_4), step_5), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": s_3, "value1": rest_6}))
}
return __t1
})
})
})
})
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], start_0), end_1).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.Int(1)
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Int(0)), gopurs_runtime.Int(1))
}
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, end_1), start_0), __t2), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
}
return __t3
})
})
	})
	return range_
}

var partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		partition = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Apply(p_0, x_2)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"no": v_3.PtrVal.(map[string]gopurs_runtime.Value)["no"], "yes": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": x_2, "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["yes"]})})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"no": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": x_2, "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["no"]}), "yes": v_3.PtrVal.(map[string]gopurs_runtime.Value)["yes"]})
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"no": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "yes": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), xs_1)
})
})
	})
	return partition
}

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")
})
	})
	return null
}

var nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		nubBy = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v1_3
} else {
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v3_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_2)
var __t3 gopurs_runtime.Value
if (v3_5_2.PtrVal.(map[string]gopurs_runtime.Value)["found"]).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, v3_5_2.PtrVal.(map[string]gopurs_runtime.Value)["result"]), v1_3), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, v3_5_2.PtrVal.(map[string]gopurs_runtime.Value)["result"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_3})), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t1 = __t3
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_5 gopurs_runtime.Value
go__4_5 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t6 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Apply(__local_var_2_4, x_3))
})
})
	})
	return nubBy
}

var nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		nub = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubBy(), dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"])
})
	})
	return nub
}

var newtypePattern gopurs_runtime.Value
var once_newtypePattern sync.Once
func Get_newtypePattern() gopurs_runtime.Value {
	once_newtypePattern.Do(func() {
		newtypePattern = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypePattern
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_2)
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v2_4_2 := gopurs_runtime.Apply(f_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, v_2), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
if (gopurs_runtime.Bool(v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2})), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t3
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
	})
	return mapMaybe
}

var manyRec gopurs_runtime.Value
var once_manyRec sync.Once
func Get_manyRec() gopurs_runtime.Value {
	once_manyRec.Do(func() {
		manyRec = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictAlternative_1 gopurs_runtime.Value) gopurs_runtime.Value {
Alt0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_1.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_3_1 := gopurs_runtime.Apply(dictAlternative_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(p_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(Alt0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Alt0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Control_Monad_Rec_Class.Get_Loop()), p_4)), gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": pkg_Data_Unit.Get_unit()})))), gopurs_runtime.Func(func(aa_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(aa_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": aa_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": acc_5})})
} else {
if (gopurs_runtime.Bool(aa_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
var go__7_3 gopurs_runtime.Value
go__7_3 = gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = v_8
} else {
if (gopurs_runtime.Bool(v1_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__7_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_8})), v1_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(go__7_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), acc_5)})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t2)
}))
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
})
})
	})
	return manyRec
}

var someRec gopurs_runtime.Value
var once_someRec sync.Once
func Get_someRec() gopurs_runtime.Value {
	once_someRec.Do(func() {
		someRec = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictAlternative_1 gopurs_runtime.Value) gopurs_runtime.Value {
manyRec2_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_manyRec(), dictMonadRec_0), dictAlternative_1)
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_1.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_List_Types.Get_Cons()), v_3)), gopurs_runtime.Apply(manyRec2_2_0, v_3))
})
})
})
	})
	return someRec
}

var some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		some = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictLazy_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_List_Types.Get_Cons()), v_2)), gopurs_runtime.Apply(dictLazy_1.PtrVal.(map[string]gopurs_runtime.Value)["defer"], gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_many(), dictAlternative_0), dictLazy_1), v_2)
})))
})
})
})
	})
	return some
}

var many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		many = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictLazy_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_some(), dictAlternative_0), dictLazy_1), v_2)), gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))
})
})
})
	})
	return many
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = b_1
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__0_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), b_1), gopurs_runtime.Int(1))), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Int(0))
}()
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = gopurs_runtime.Apply(Get_last(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t0 = __t1
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}
}()
})
	})
	return last
}

var insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		insertBy = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_insertBy(), v_0), v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1, "value1": v2_2})
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}
}()
})
})
})
	})
	return insertBy
}

var insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		insertAt = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1, "value1": v2_2})})
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_insertAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t0 = __t2
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
}
return __t0
}
}()
})
})
})
	})
	return insertAt
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_insertBy(), dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"])
})
	})
	return insert
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(lst_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_unsnoc(), lst_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["init"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
	})
	return init_
}

var index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		index = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v1_1), gopurs_runtime.Int(1)))
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}
}()
})
})
	})
	return index
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return head
}

var transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		transpose = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(Get_transpose(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var go__1_2 gopurs_runtime.Value
go__1_2 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__4_5 gopurs_runtime.Value
go__4_5 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t6 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_2)
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_2, v_2), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2})), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t3 = __t4
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
var go__1_7 gopurs_runtime.Value
go__1_7 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__4_10 gopurs_runtime.Value
go__4_10 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t11 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_10, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t11
})
})
__t8 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_10, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_2)
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t9 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_7, v_2), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_7, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v_2})), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t8 = __t9
} else {
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t8
})
})
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])}), "value1": gopurs_runtime.Apply(Get_transpose(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_7, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])}))})
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
}
}()
})
	})
	return transpose
}

var groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		groupBy = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v2_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_span(), gopurs_runtime.Apply(v_0, v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_2_1.PtrVal.(map[string]gopurs_runtime.Value)["init"]}), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_groupBy(), v_0), v2_2_1.PtrVal.(map[string]gopurs_runtime.Value)["rest"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}
}()
})
})
	})
	return groupBy
}

var groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		groupAllBy = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(p_0, x_1), y_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")
})
}))
__local_var_2_1 := gopurs_runtime.Apply(Get_sortBy(), p_0)
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
})
	})
	return groupAllBy
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_groupBy(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])
})
	})
	return group
}

var groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		groupAll = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_group(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
compare_2_1 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(Get_sortBy(), compare_2_1), x_3))
})
})
	})
	return groupAll
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
	})
	return fromFoldable
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], v1_2)
} else {
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__local_var_4_1 := v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, v1_2), v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldM(), dictMonad_0), v_1), b_prime_5), __local_var_4_1)
}))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}
}()
})
})
})
})
	})
	return foldM
}

var findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		findIndex = gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(fn_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2})
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v_2), gopurs_runtime.Int(1))), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t1 = __t2
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Int(0))
})
	})
	return findIndex
}

var findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		findLastIndex = gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_1 gopurs_runtime.Value
go__2_1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = v_3
} else {
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_3})), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_findIndex(), fn_0), gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), xs_1))
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), xs_1)), gopurs_runtime.Int(1))), __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t3
})
})
	})
	return findLastIndex
}

var filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		filterM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
} else {
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__local_var_5_3 := v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_6_4 := v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v_3, __local_var_5_3)), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_filterM(), dictMonad_0), v_3), __local_var_6_4)), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (b_7).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_5_3, "value1": xs_prime_8})
} else {
__t5 = xs_prime_8
}
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t5)
}))
}))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
}
}()
})
	})
	return filterM
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__4_3 gopurs_runtime.Value
go__4_3 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = v_5
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_2)
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(p_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2})), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, v_2), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
	})
	return filter
}

var intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		intersectBy = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = v_7
} else {
if (gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_7})), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_4)
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(Get_any(), gopurs_runtime.Apply(v_0, v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v2_2)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4})), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, v_4), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v1_1)
}
}
return __t5
})
})
})
	})
	return intersectBy
}

var intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		intersect = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_intersectBy(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])
})
	})
	return intersect
}

var nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		nubByEq = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__local_var_2_1 := v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__6_5 gopurs_runtime.Value
go__6_5 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t6 = v_7
} else {
if (gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_7})), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
})
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__6_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_4)
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, __local_var_2_1), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4})), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_2, v_4), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t3 = __t4
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_2_1, "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_nubByEq(), v_0), gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}
}()
})
})
	})
	return nubByEq
}

var nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		nubEq = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubByEq(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])
})
	})
	return nubEq
}

var eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		eqPattern = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_eq1List().PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_0), x_1), y_2)
})
})})
})
	})
	return eqPattern
}

var ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		ordPattern = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eqPattern1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_eq1List().PtrVal.(map[string]gopurs_runtime.Value)["eq1"], __local_var_1_0), x_2), y_3)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_ordList(), dictOrd_0).PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_3), y_4)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_2_1
})})
})
	})
	return ordPattern
}

var elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		elemLastIndex = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_2), x_1)
}))
})
})
	})
	return elemLastIndex
}

var elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		elemIndex = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_2), x_1)
}))
})
})
	})
	return elemIndex
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(p_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = v_2
}
return __t1
})
return go__1_0
})
	})
	return dropWhile
}

var dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		dropEnd = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), xs_1)), n_0)), xs_1)
})
})
	})
	return dropEnd
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = v1_1
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_drop(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
}
}()
})
})
	})
	return drop
}

var slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		slice = gopurs_runtime.Func(func(start_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(end_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), end_1), start_0)), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_drop(), start_0), xs_2))
})
})
})
	})
	return slice
}

var takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		takeEnd = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_drop(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), xs_1)), n_0)), xs_1)
})
})
	})
	return takeEnd
}

var deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		deleteBy = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t1 = v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_deleteBy(), v_0), v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}
}()
})
})
})
	})
	return deleteBy
}

var unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		unionBy = gopurs_runtime.Func(func(eq2_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = b_4
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_deleteBy(), eq2_0), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), b_4)), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(Get_nubByEq(), eq2_0), ys_2)), xs_1)), xs_1)
})
})
})
	})
	return unionBy
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unionBy(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])
})
	})
	return union
}

var deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		deleteAt = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
} else {
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_deleteAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t3 = __t2
}
__t0 = __t3
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}
}()
})
})
	})
	return deleteAt
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_deleteBy(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])
})
	})
	return delete_
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = b_2
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_deleteBy(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), b_2)), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return go__1_0
})
	})
	return difference
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func(func(b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_bindList().PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_1), b_0)
})
})
	})
	return concatMap
}

var concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		concat = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_bindList().PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return concat
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__3_3 gopurs_runtime.Value
go__3_3 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t4 = v_4
} else {
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4})), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), v_1)
} else {
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__0_0, v_1), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__0_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1})), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
}()
	})
	return catMaybes
}

var alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		alterAt = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
v3_3_4 := gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_3_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
} else {
if (gopurs_runtime.Bool(v3_3_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v3_3_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __t5})
} else {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_alterAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t3 = __t2
}
__t0 = __t3
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}
}()
})
})
})
	})
	return alterAt
}

var modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		modifyAt = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_alterAt(), n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(f_1, x_2)})
}))
})
})
	})
	return modifyAt
}


