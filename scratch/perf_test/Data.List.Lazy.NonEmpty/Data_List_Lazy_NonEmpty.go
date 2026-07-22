package Data_List_Lazy_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_List_Lazy "gopurs/output/Data.List.Lazy"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "tail": v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
	})
	return uncons
}

var toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		toList = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
__local_var_2_1 := v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_3_2 := v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_2_1, "value1": __local_var_3_2})
}))
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
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_uncons(), xs_1)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["head"], "value1": __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["tail"]})})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t2
}))
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(Get_toList(), x_2))
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
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return tail
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = pkg_Data_List_Lazy_Types.Get_applicativeNonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["pure"]
	})
	return singleton
}

var repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		repeat = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": x_0, "value1": gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_fix(), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": x_0, "value1": xs_2})
}))
}))})
}))
})
	})
	return repeat
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_last(), v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
if (gopurs_runtime.Bool(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
	})
	return last
}

var iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		iterate = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": x_1, "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_iterate(), f_0), gopurs_runtime.Apply(f_0, x_1))})
}))
})
})
	})
	return iterate
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_init_(), v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = pkg_Data_List_Lazy_Types.Get_nil()
} else {
if (gopurs_runtime.Bool(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_3_3 := __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t2 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_3})
}))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
	})
	return init_
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
	})
	return head
}

var fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		fromList = gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
v_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), l_0))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__local_var_2_2 := v_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_3_3 := v_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": __local_var_2_2, "value1": __local_var_3_3})
}))})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
	})
	return fromList
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Lazy_Types.Get_cons()), pkg_Data_List_Lazy_Types.Get_nil())
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_fromList(), gopurs_runtime.Apply(__local_var_1_0, x_2))
})
})
	})
	return fromFoldable
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1)
__local_var_4_1 := v2_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_5_2 := v2_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": y_0, "value1": gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_4_1, "value1": __local_var_5_2})
}))})
}))
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_bindNonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_1), b_0)
})
})
	})
	return concatMap
}

var appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		appendFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
fromFoldable1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], pkg_Data_List_Lazy_Types.Get_cons()), pkg_Data_List_Lazy_Types.Get_nil())
return gopurs_runtime.Func(func(nel_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_2).PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_semigroupList().PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_2).PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Apply(fromFoldable1_1_0, ys_3))})
}))
})
})
})
	})
	return appendFoldable
}


