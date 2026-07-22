package Data_Foldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		monoidEndo = func() gopurs_runtime.Value {
semigroupEndo1_0_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
})})
}()
	})
	return monoidEndo
}

var Empty gopurs_runtime.Value
var once_Empty sync.Once
func Get_Empty() gopurs_runtime.Value {
	once_Empty.Do(func() {
		Empty = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")})
	})
	return Empty
}

var Node gopurs_runtime.Value
var once_Node sync.Once
func Get_Node() gopurs_runtime.Value {
	once_Node.Do(func() {
		Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": value0})
})
	})
	return Node
}

var Append gopurs_runtime.Value
var once_Append sync.Once
func Get_Append() gopurs_runtime.Value {
	once_Append.Do(func() {
		Append = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Append"), "value0": value0, "value1": value1})
})
})
	})
	return Append
}

var semigroupFreeMonoidTree gopurs_runtime.Value
var once_semigroupFreeMonoidTree sync.Once
func Get_semigroupFreeMonoidTree() gopurs_runtime.Value {
	once_semigroupFreeMonoidTree.Do(func() {
		semigroupFreeMonoidTree = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": Get_Append()})
	})
	return semigroupFreeMonoidTree
}

var monoidFreeMonoidTree gopurs_runtime.Value
var once_monoidFreeMonoidTree sync.Once
func Get_monoidFreeMonoidTree() gopurs_runtime.Value {
	once_monoidFreeMonoidTree.Do(func() {
		monoidFreeMonoidTree = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFreeMonoidTree()
})})
	})
	return monoidFreeMonoidTree
}

var foldr gopurs_runtime.Value
var once_foldr sync.Once
func Get_foldr() gopurs_runtime.Value {
	once_foldr.Do(func() {
		foldr = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"]
})
	})
	return foldr
}

var indexr gopurs_runtime.Value
var once_indexr sync.Once
func Get_indexr() gopurs_runtime.Value {
	once_indexr.Do(func() {
		indexr = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(idx_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cursor_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(cursor_3.PtrVal.(map[string]gopurs_runtime.Value)["elem"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = cursor_3
} else {
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), cursor_3.PtrVal.(map[string]gopurs_runtime.Value)["pos"]), idx_1)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"elem": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": a_2}), "pos": cursor_3.PtrVal.(map[string]gopurs_runtime.Value)["pos"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pos": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), cursor_3.PtrVal.(map[string]gopurs_runtime.Value)["pos"]), gopurs_runtime.Int(1)), "elem": cursor_3.PtrVal.(map[string]gopurs_runtime.Value)["elem"]})
}
}
return __t1
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"elem": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "pos": gopurs_runtime.Int(0)}))
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3).PtrVal.(map[string]gopurs_runtime.Value)["elem"]
})
})
})
	})
	return indexr
}

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
})), gopurs_runtime.Bool(true))
})
	})
	return null
}

var oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		oneOf = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictPlus_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Apply(dictPlus_1.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alt"]), dictPlus_1.PtrVal.(map[string]gopurs_runtime.Value)["empty"])
})
})
	})
	return oneOf
}

var oneOfMap gopurs_runtime.Value
var once_oneOfMap sync.Once
func Get_oneOfMap() gopurs_runtime.Value {
	once_oneOfMap.Do(func() {
		oneOfMap = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictPlus_1 gopurs_runtime.Value) gopurs_runtime.Value {
empty_2_0 := dictPlus_1.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictPlus_1.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(f_3, x_4))
})), empty_2_0)
})
})
})
	})
	return oneOfMap
}

var traverse_ gopurs_runtime.Value
var once_traverse_ sync.Once
func Get_traverse_() gopurs_runtime.Value {
	once_traverse_.Do(func() {
		traverse_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
applySecond_1_0 := gopurs_runtime.Apply(pkg_Control_Apply.Get_applySecond(), gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_2.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
})), gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit()))
})
})
})
	})
	return traverse_
}

var for_ gopurs_runtime.Value
var once_for_ sync.Once
func Get_for_() gopurs_runtime.Value {
	once_for_.Do(func() {
		for_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictApplicative_0)
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(traverse_1_1_0, dictFoldable_2)
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1, a_5), b_4)
})
})
})
})
	})
	return for_
}

var sequence_ gopurs_runtime.Value
var once_sequence_ sync.Once
func Get_sequence_() gopurs_runtime.Value {
	once_sequence_.Do(func() {
		sequence_ = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictApplicative_0)
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(traverse_1_1_0, dictFoldable_2), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return sequence_
}

var foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		foldl = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"]
})
	})
	return foldl
}

var indexl gopurs_runtime.Value
var once_indexl sync.Once
func Get_indexl() gopurs_runtime.Value {
	once_indexl.Do(func() {
		indexl = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(idx_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(cursor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(cursor_2.PtrVal.(map[string]gopurs_runtime.Value)["elem"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = cursor_2
} else {
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), cursor_2.PtrVal.(map[string]gopurs_runtime.Value)["pos"]), idx_1)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"elem": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": a_3}), "pos": cursor_2.PtrVal.(map[string]gopurs_runtime.Value)["pos"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pos": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), cursor_2.PtrVal.(map[string]gopurs_runtime.Value)["pos"]), gopurs_runtime.Int(1)), "elem": cursor_2.PtrVal.(map[string]gopurs_runtime.Value)["elem"]})
}
}
return __t1
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"elem": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "pos": gopurs_runtime.Int(0)}))
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3).PtrVal.(map[string]gopurs_runtime.Value)["elem"]
})
})
})
	})
	return indexl
}

var intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		intercalate = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_3_1 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(sep_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (v_6.PtrVal.(map[string]gopurs_runtime.Value)["init"]).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": v1_7})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["acc"]), gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], sep_4), v1_7))})
}
return __t2
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(true), "acc": mempty_3_1})), xs_5).PtrVal.(map[string]gopurs_runtime.Value)["acc"]
})
})
})
})
	})
	return intercalate
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
one_2_0 := dictSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["one"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["add"], one_2_0), c_3)
})
})), dictSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["zero"])
})
})
	})
	return length
}

var maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		maximumBy = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cmp_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_3})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = v1_3
}
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __t1})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return maximumBy
}

var maximum gopurs_runtime.Value
var once_maximum sync.Once
func Get_maximum() gopurs_runtime.Value {
	once_maximum.Do(func() {
		maximum = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_3})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = v1_3
}
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __t1})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return maximum
}

var minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		minimumBy = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cmp_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_3})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = v1_3
}
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __t1})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return minimumBy
}

var minimum gopurs_runtime.Value
var once_minimum sync.Once
func Get_minimum() gopurs_runtime.Value {
	once_minimum.Do(func() {
		minimum = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_3})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = v1_3
}
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __t1})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return minimum
}

var product gopurs_runtime.Value
var once_product sync.Once
func Get_product() gopurs_runtime.Value {
	once_product.Do(func() {
		product = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], dictSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["mul"]), dictSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["one"])
})
})
	})
	return product
}

var sum gopurs_runtime.Value
var once_sum sync.Once
func Get_sum() gopurs_runtime.Value {
	once_sum.Do(func() {
		sum = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemiring_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], dictSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["add"]), dictSemiring_1.PtrVal.(map[string]gopurs_runtime.Value)["zero"])
})
})
	})
	return sum
}

var foldableTuple gopurs_runtime.Value
var once_foldableTuple sync.Once
func Get_foldableTuple() gopurs_runtime.Value {
	once_foldableTuple.Do(func() {
		foldableTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
})})
	})
	return foldableTuple
}

var foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		foldableMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
})})
	})
	return foldableMultiplicative
}

var foldableMaybe gopurs_runtime.Value
var once_foldableMaybe sync.Once
func Get_foldableMaybe() gopurs_runtime.Value {
	once_foldableMaybe.Do(func() {
		foldableMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v1_1
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_1)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
}), "foldl": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = v1_1
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = mempty_1_2
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
})})
	})
	return foldableMaybe
}

var foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		foldableIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
})})
	})
	return foldableIdentity
}

var foldableEither gopurs_runtime.Value
var once_foldableEither sync.Once
func Get_foldableEither() gopurs_runtime.Value {
	once_foldableEither.Do(func() {
		foldableEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = v1_1
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_1)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
}), "foldl": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = v1_1
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = mempty_1_2
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
})})
	})
	return foldableEither
}

var foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		foldableDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
})})
	})
	return foldableDual
}

var foldableDisj gopurs_runtime.Value
var once_foldableDisj sync.Once
func Get_foldableDisj() gopurs_runtime.Value {
	once_foldableDisj.Do(func() {
		foldableDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
})})
	})
	return foldableDisj
}

var foldableConst gopurs_runtime.Value
var once_foldableConst sync.Once
func Get_foldableConst() gopurs_runtime.Value {
	once_foldableConst.Do(func() {
		foldableConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}), "foldl": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
})
})})
	})
	return foldableConst
}

var foldableConj gopurs_runtime.Value
var once_foldableConj sync.Once
func Get_foldableConj() gopurs_runtime.Value {
	once_foldableConj.Do(func() {
		foldableConj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
})})
	})
	return foldableConj
}

var foldableAdditive gopurs_runtime.Value
var once_foldableAdditive sync.Once
func Get_foldableAdditive() gopurs_runtime.Value {
	once_foldableAdditive.Do(func() {
		foldableAdditive = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2), z_1)
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
})})
	})
	return foldableAdditive
}

var foldMapDefaultR gopurs_runtime.Value
var once_foldMapDefaultR sync.Once
func Get_foldMapDefaultR() gopurs_runtime.Value {
	once_foldMapDefaultR.Do(func() {
		foldMapDefaultR = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_2_0 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_3, x_4)), acc_5)
})
})), mempty_2_0)
})
})
})
	})
	return foldMapDefaultR
}

var foldableArray gopurs_runtime.Value
var once_foldableArray sync.Once
func Get_foldableArray() gopurs_runtime.Value {
	once_foldableArray.Do(func() {
		foldableArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": Get_foldrArray(), "foldl": Get_foldlArray(), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableArray().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_2, x_3)), acc_4)
})
})), mempty_1_0)
})
})})
	})
	return foldableArray
}

var foldableFreeMonoidTree gopurs_runtime.Value
var once_foldableFreeMonoidTree sync.Once
func Get_foldableFreeMonoidTree() gopurs_runtime.Value {
	once_foldableFreeMonoidTree.Do(func() {
		foldableFreeMonoidTree = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldl": gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(fn_0, acc_2), lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), rhs_4), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")}))
} else {
if (gopurs_runtime.Bool(lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Append")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, acc_2), lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), rhs_4)
} else {
if (gopurs_runtime.Bool(rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, acc_2), lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, acc_2), lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Append"), "value0": lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": rhs_4}))
}
}
__t1 = __t3
} else {
if (gopurs_runtime.Bool(lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
__t2 = acc_2
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, acc_2), rhs_4), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")}))
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, a_2), b_3), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")}))
})
})
}), "foldr": gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_4 gopurs_runtime.Value
go__1_4 = gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_4, gopurs_runtime.Apply(gopurs_runtime.Apply(fn_0, rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), acc_2)), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")})), lhs_3)
} else {
if (gopurs_runtime.Bool(rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Append")).IntVal != 0 {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_4, acc_2), lhs_3), rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
if (gopurs_runtime.Bool(lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_4, acc_2), rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
} else {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_4, acc_2), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Append"), "value0": lhs_3, "value1": rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})), rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
}
__t5 = __t7
} else {
if (gopurs_runtime.Bool(rhs_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(lhs_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Empty")).IntVal != 0 {
__t6 = acc_2
} else {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_4, acc_2), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")})), lhs_3)
}
__t5 = __t6
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t5
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_4, a_2), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Empty")})), b_3)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_8 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableFreeMonoidTree().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_2, x_3)), acc_4)
})
})), mempty_1_8)
})
})})
	})
	return foldableFreeMonoidTree
}

var foldMapDefaultL gopurs_runtime.Value
var once_foldMapDefaultL sync.Once
func Get_foldMapDefaultL() gopurs_runtime.Value {
	once_foldMapDefaultL.Do(func() {
		foldMapDefaultL = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_2_0 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], acc_4), gopurs_runtime.Apply(f_3, x_5))
})
})), mempty_2_0)
})
})
})
	})
	return foldMapDefaultL
}

var foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		foldMap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"]
})
	})
	return foldMap
}

var foldableApp gopurs_runtime.Value
var once_foldableApp sync.Once
func Get_foldableApp() gopurs_runtime.Value {
	once_foldableApp.Do(func() {
		foldableApp = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_1), i_2), v_3)
})
})
}), "foldl": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_1), i_2), v_3)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_1)
})})
})
	})
	return foldableApp
}

var foldableCompose gopurs_runtime.Value
var once_foldableCompose sync.Once
func Get_foldableCompose() gopurs_runtime.Value {
	once_foldableCompose.Do(func() {
		foldableCompose = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_2)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_0, a_7), b_6)
})
})), i_3), v_4)
})
})
}), "foldl": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2)), i_3), v_4)
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_1 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
foldMap5_4_2 := gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap4_3_1, gopurs_runtime.Apply(foldMap5_4_2, f_5)), v_6)
})
})
})})
})
})
	})
	return foldableCompose
}

var foldableCoproduct gopurs_runtime.Value
var once_foldableCoproduct sync.Once
func Get_foldableCoproduct() gopurs_runtime.Value {
	once_foldableCoproduct.Do(func() {
		foldableCoproduct = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_2), z_3)
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_2), z_3)
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_4_0, v2_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_5_1, v2_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
}), "foldl": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2), z_3)
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2), z_3)
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_4_3, v2_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_5_4, v2_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_6 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
foldMap5_4_7 := gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_8 := gopurs_runtime.Apply(foldMap4_3_6, f_5)
__local_var_7_9 := gopurs_runtime.Apply(foldMap5_4_7, f_5)
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t10 = gopurs_runtime.Apply(__local_var_6_8, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t10 = gopurs_runtime.Apply(__local_var_7_9, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t10
})
})
})})
})
})
	})
	return foldableCoproduct
}

var foldableFirst gopurs_runtime.Value
var once_foldableFirst sync.Once
func Get_foldableFirst() gopurs_runtime.Value {
	once_foldableFirst.Do(func() {
		foldableFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = z_1
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), z_1)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = z_1
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = mempty_1_2
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
})})
	})
	return foldableFirst
}

var foldableLast gopurs_runtime.Value
var once_foldableLast sync.Once
func Get_foldableLast() gopurs_runtime.Value {
	once_foldableLast.Do(func() {
		foldableLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = z_1
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), z_1)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = z_1
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, z_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = mempty_1_2
} else {
if (gopurs_runtime.Bool(v1_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(v_2, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})
})})
	})
	return foldableLast
}

var foldableProduct gopurs_runtime.Value
var once_foldableProduct sync.Once
func Get_foldableProduct() gopurs_runtime.Value {
	once_foldableProduct.Do(func() {
		foldableProduct = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_2), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_2), z_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
})
})
}), "foldl": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], f_2), z_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_3_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
foldMap5_4_1 := gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap4_3_0, f_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap5_4_1, f_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})})
})
})
	})
	return foldableProduct
}

var foldlDefault gopurs_runtime.Value
var once_foldlDefault sync.Once
func Get_foldlDefault() gopurs_runtime.Value {
	once_foldlDefault.Do(func() {
		foldlDefault = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], Get_monoidFreeMonoidTree())
return gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableFreeMonoidTree().PtrVal.(map[string]gopurs_runtime.Value)["foldl"], c_2), u_3), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap2_1_0, Get_Node()), xs_4))
})
})
})
})
	})
	return foldlDefault
}

var foldrDefault gopurs_runtime.Value
var once_foldrDefault sync.Once
func Get_foldrDefault() gopurs_runtime.Value {
	once_foldrDefault.Do(func() {
		foldrDefault = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], Get_monoidFreeMonoidTree())
return gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableFreeMonoidTree().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], c_2), u_3), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap2_1_0, Get_Node()), xs_4))
})
})
})
})
	})
	return foldrDefault
}

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], pkg_Data_Maybe_First.Get_monoidFirst())
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(foldMap2_1_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_2.PtrVal.(map[string]gopurs_runtime.Value)["eq"], a_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t2
}))
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
})
})
	})
	return lookup
}

var surroundMap gopurs_runtime.Value
var once_surroundMap sync.Once
func Get_surroundMap() gopurs_runtime.Value {
	once_surroundMap.Do(func() {
		surroundMap = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap2_1_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], Get_monoidEndo())
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap2_1_0, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], d_3), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(t_4, a_6)), m_7))
})
})), f_5)), d_3)
})
})
})
})
})
	})
	return surroundMap
}

var surround gopurs_runtime.Value
var once_surround sync.Once
func Get_surround() gopurs_runtime.Value {
	once_surround.Do(func() {
		surround = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
surroundMap1_1_0 := gopurs_runtime.Apply(Get_surroundMap(), dictFoldable_0)
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
surroundMap2_3_1 := gopurs_runtime.Apply(surroundMap1_1_0, dictSemigroup_2)
return gopurs_runtime.Func(func(d_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(surroundMap2_3_1, d_4), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
})
	})
	return surround
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], b_4), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, a_6), a_5)
}))
})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], b0_3))
})
})
})
})
	})
	return foldM
}

var fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		fold = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return fold
}

var findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		findMap = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(p_1, v1_3)
} else {
__t0 = v_2
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return findMap
}

var find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		find = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing").IntVal != 0 && gopurs_runtime.Apply(p_1, v1_3).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_3})
} else {
__t0 = v_2
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
})
})
	})
	return find
}

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_2_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_2), v1_3)
})
})})
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["ff"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
})})))
})
})
	})
	return any
}

var elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		elem = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
any1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_any(), dictFoldable_0), pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean())
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(dictEq_2.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_3))
})
})
})
	})
	return elem
}

var notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		notElem = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
elem1_1_0 := gopurs_runtime.Apply(Get_elem(), dictFoldable_0)
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
elem2_3_1 := gopurs_runtime.Apply(elem1_1_0, dictEq_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(elem2_3_1, x_4)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(__local_var_5_2, x_6))
})
})
})
})
	})
	return notElem
}

var or gopurs_runtime.Value
var once_or sync.Once
func Get_or() gopurs_runtime.Value {
	once_or.Do(func() {
		or = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_any(), dictFoldable_0), dictHeytingAlgebra_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return or
}

var all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		all = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupConj1_2_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["conj"], v_2), v1_3)
})
})})
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictHeytingAlgebra_1.PtrVal.(map[string]gopurs_runtime.Value)["tt"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
})})))
})
})
	})
	return all
}

var and gopurs_runtime.Value
var once_and sync.Once
func Get_and() gopurs_runtime.Value {
	once_and.Do(func() {
		and = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_all(), dictFoldable_0), dictHeytingAlgebra_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return and
}

func Get_foldlArray() gopurs_runtime.Value {
	return FoldlArray
}

func Get_foldrArray() gopurs_runtime.Value {
	return FoldrArray
}
