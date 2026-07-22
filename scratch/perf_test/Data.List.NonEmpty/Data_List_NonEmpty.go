package Data_List_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__3_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))})
})
})
})
	})
	return zipWith
}

var zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		zipWithA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence11_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_traversable1NonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["traverse1"], gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence11_1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_zipWith(), f_2), xs_3), ys_4))
})
})
})
})
	})
	return zipWithA
}

var zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		zip = gopurs_runtime.Apply(Get_zipWith(), pkg_Data_Tuple.Get_Tuple())
	})
	return zip
}

var wrappedOperation2 gopurs_runtime.Value
var once_wrappedOperation2 sync.Once
func Get_wrappedOperation2() gopurs_runtime.Value {
	once_wrappedOperation2.Do(func() {
		wrappedOperation2 = gopurs_runtime.Func(func(name_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v2_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
} else {
if (gopurs_runtime.Bool(v2_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList ")), name_0)
__t1 = gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), __local_var_5_2)
}))
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
})
	})
	return wrappedOperation2
}

var wrappedOperation gopurs_runtime.Value
var once_wrappedOperation sync.Once
func Get_wrappedOperation() gopurs_runtime.Value {
	once_wrappedOperation.Do(func() {
		wrappedOperation = gopurs_runtime.Func(func(name_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
v1_3_0 := gopurs_runtime.Apply(f_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v1_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
} else {
if (gopurs_runtime.Bool(v1_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList ")), name_0)
__t1 = gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), __local_var_4_2)
}))
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
	})
	return wrappedOperation
}

var updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		updateAt = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), i_0), gopurs_runtime.Int(0))).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": a_1, "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_updateAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), i_0), gopurs_runtime.Int(1))), a_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t2 = __t1
}
return __t2
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
		unzip = gopurs_runtime.Func(func(ts_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": ts_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_listMap(), pkg_Data_Tuple.Get_fst()), ts_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": ts_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_listMap(), pkg_Data_Tuple.Get_snd()), ts_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})})
})
	})
	return unzip
}

var unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		unsnoc = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}), "last": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["init"]}), "last": v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["last"]})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
	})
	return unsnoc
}

var unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		unionBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation2(), gopurs_runtime.Str("unionBy")), gopurs_runtime.Apply(pkg_Data_List.Get_unionBy(), x_0))
})
	})
	return unionBy
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation2(), gopurs_runtime.Str("union")), gopurs_runtime.Apply(pkg_Data_List.Get_union(), dictEq_0))
})
	})
	return union
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "tail": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
	})
	return uncons
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

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(xs_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(xs_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": xs_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": xs_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
}))
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return tail
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("sortBy")), gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), x_0))
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("sortBy")), gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), compare_1_0)), xs_2)
})
})
	})
	return sort
}

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": y_1, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
	})
	return snoc
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": x_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
})
	})
	return singleton
}

var snoc_prime gopurs_runtime.Value
var once_snoc_prime sync.Once
func Get_snoc_prime() gopurs_runtime.Value {
	once_snoc_prime.Do(func() {
		snoc_prime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v1_1, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
	})
	return snoc_prime
}

var reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		reverse = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("reverse")), pkg_Data_List.Get_reverse())
	})
	return reverse
}

var nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		nubEq = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("nubEq")), gopurs_runtime.Apply(pkg_Data_List.Get_nubEq(), dictEq_0))
})
	})
	return nubEq
}

var nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		nubByEq = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("nubByEq")), gopurs_runtime.Apply(pkg_Data_List.Get_nubByEq(), x_0))
})
	})
	return nubByEq
}

var nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		nubBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("nubBy")), gopurs_runtime.Apply(pkg_Data_List.Get_nubBy(), x_0))
})
	})
	return nubBy
}

var nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		nub = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("nub")), gopurs_runtime.Apply(pkg_Data_List.Get_nubBy(), dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]))
})
	})
	return nub
}

var modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		modifyAt = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), i_0), gopurs_runtime.Int(0))).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_alterAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), i_0), gopurs_runtime.Int(1))), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(f_1, x_3)})
})), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t2 = __t1
}
return __t2
})
})
})
	})
	return modifyAt
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
v2_4_2 := gopurs_runtime.Apply(x_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
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
__local_var_2_6 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return mapMaybe
}

var partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		partition = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_partition(), x_0), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return partition
}

var span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		span = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_span(), x_0), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return span
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_take(), x_0)
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return take
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(x_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)).IntVal != 0 {
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
__local_var_2_4 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return takeWhile
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), gopurs_runtime.Apply(pkg_Data_List.Get_length(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_List.Get_last(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_List.Get_last(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(pkg_Data_List.Get_last(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]).PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
__t0 = __t1
} else {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
}
return __t0
})
	})
	return last
}

var intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		intersectBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation2(), gopurs_runtime.Str("intersectBy")), gopurs_runtime.Apply(pkg_Data_List.Get_intersectBy(), x_0))
})
	})
	return intersectBy
}

var intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		intersect = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation2(), gopurs_runtime.Str("intersect")), gopurs_runtime.Apply(pkg_Data_List.Get_intersect(), dictEq_0))
})
	})
	return intersect
}

var insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		insertAt = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), i_0), gopurs_runtime.Int(0))).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": a_1, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})})
} else {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_insertAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), i_0), gopurs_runtime.Int(1))), a_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t2 = __t1
}
return __t2
})
})
})
	})
	return insertAt
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["init"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
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
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), i_1), gopurs_runtime.Int(0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_index(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), i_1), gopurs_runtime.Int(1)))
}
return __t0
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
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
	})
	return head
}

var groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		groupBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("groupBy")), gopurs_runtime.Apply(pkg_Data_List.Get_groupBy(), x_0))
})
	})
	return groupBy
}

var groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		groupAllBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("groupAllBy")), gopurs_runtime.Apply(pkg_Data_List.Get_groupAllBy(), x_0))
})
	})
	return groupAllBy
}

var groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		groupAll = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("groupAll")), gopurs_runtime.Apply(pkg_Data_List.Get_groupAll(), dictOrd_0))
})
	})
	return groupAll
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_wrappedOperation(), gopurs_runtime.Str("group")), gopurs_runtime.Apply(pkg_Data_List.Get_group(), dictEq_0))
})
	})
	return group
}

var fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		fromList = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return fromList
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": __local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
	})
	return fromFoldable
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, b_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_foldM(), dictMonad_0), f_1), b_prime_5), __local_var_4_0)
}))
})
})
})
})
	})
	return foldM
}

var findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		findLastIndex = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_findLastIndex(), f_0), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))})
} else {
if (gopurs_runtime.Bool(v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Int(0)})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
	})
	return findLastIndex
}

var findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		findIndex = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Int(0)})
} else {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_findIndex(), f_0), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t2 = __t1
}
return __t2
})
})
	})
	return findIndex
}

var filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		filterM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_filterM(), dictMonad_0)
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
})
	})
	return filterM
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Apply(x_0, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
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
__local_var_2_5 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return filter
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
		dropWhile = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(x_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t1 = v_2
}
return __t1
})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return dropWhile
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List.Get_drop(), x_0), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
})
	})
	return drop
}

var cons_prime gopurs_runtime.Value
var once_cons_prime sync.Once
func Get_cons_prime() gopurs_runtime.Value {
	once_cons_prime.Do(func() {
		cons_prime = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": x_0, "value1": xs_1})
})
})
	})
	return cons_prime
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": y_0, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
})
})
	})
	return cons
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func(func(b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_bindNonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_1), b_0)
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_bindNonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return concat
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, v_2), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_2})), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})
	})
	return catMaybes
}

var appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		appendFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
fromFoldable1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Types.Get_Cons()), gopurs_runtime.Apply(fromFoldable1_1_0, ys_3)), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})
	})
	return appendFoldable
}


