package Data_Semigroup_Foldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var FoldRight1 gopurs_runtime.Value
var once_FoldRight1 sync.Once
func Get_FoldRight1() gopurs_runtime.Value {
	once_FoldRight1.Do(func() {
		FoldRight1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("FoldRight1"), "value0": value0, "value1": value1})
})
})
	})
	return FoldRight1
}

var mkFoldRight1 gopurs_runtime.Value
var once_mkFoldRight1 sync.Once
func Get_mkFoldRight1() gopurs_runtime.Value {
	once_mkFoldRight1.Do(func() {
		mkFoldRight1 = gopurs_runtime.Apply(Get_FoldRight1(), pkg_Data_Function.Get_const_())
	})
	return mkFoldRight1
}

var foldr1 gopurs_runtime.Value
var once_foldr1 sync.Once
func Get_foldr1() gopurs_runtime.Value {
	once_foldr1.Do(func() {
		foldr1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr1"]
})
	})
	return foldr1
}

var foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		foldl1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl1"]
})
	})
	return foldl1
}

var maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		maximumBy = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cmp_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl1"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_1, x_2), y_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t0 = x_2
} else {
__t0 = y_3
}
return __t0
})
}))
})
})
	})
	return maximumBy
}

var minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		minimumBy = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cmp_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl1"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_1, x_2), y_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = x_2
} else {
__t0 = y_3
}
return __t0
})
}))
})
})
	})
	return minimumBy
}

var foldableTuple gopurs_runtime.Value
var once_foldableTuple sync.Once
func Get_foldableTuple() gopurs_runtime.Value {
	once_foldableTuple.Do(func() {
		foldableTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap1": gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
}), "foldl1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
})})
	})
	return foldableTuple
}

var foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		foldableMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), "foldl1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), "foldMap1": gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
})})
	})
	return foldableMultiplicative
}

var foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		foldableIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap1": gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), "foldl1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), "foldr1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
})})
	})
	return foldableIdentity
}

var foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		foldableDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldr1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), "foldl1": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), "foldMap1": gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
})})
	})
	return foldableDual
}

var foldRight1Semigroup gopurs_runtime.Value
var once_foldRight1Semigroup sync.Once
func Get_foldRight1Semigroup() gopurs_runtime.Value {
	once_foldRight1Semigroup.Do(func() {
		foldRight1Semigroup = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("FoldRight1"), "value0": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_4, __local_var_2_0), gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], a_3), f_4))), f_4)
})
}), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
})})
	})
	return foldRight1Semigroup
}

var semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		semigroupDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("FoldRight1"), "value0": gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_4, __local_var_2_0), gopurs_runtime.Apply(gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], a_3), f_4))), f_4)
})
}), "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
})})
	})
	return semigroupDual
}

var foldMap1DefaultR gopurs_runtime.Value
var once_foldMap1DefaultR sync.Once
func Get_foldMap1DefaultR() gopurs_runtime.Value {
	once_foldMap1DefaultR.Do(func() {
		foldMap1DefaultR = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
append_3_0 := dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(dictFunctor_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
__local_var_6_2 := gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr1"], append_3_0)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
})
})
})
	})
	return foldMap1DefaultR
}

var foldMap1DefaultL gopurs_runtime.Value
var once_foldMap1DefaultL sync.Once
func Get_foldMap1DefaultL() gopurs_runtime.Value {
	once_foldMap1DefaultL.Do(func() {
		foldMap1DefaultL = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
append_3_0 := dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(dictFunctor_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
__local_var_6_2 := gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl1"], append_3_0)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
})
})
})
	})
	return foldMap1DefaultL
}

var foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		foldMap1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"]
})
	})
	return foldMap1
}

var foldl1Default gopurs_runtime.Value
var once_foldl1Default sync.Once
func Get_foldl1Default() gopurs_runtime.Value {
	once_foldl1Default.Do(func() {
		foldl1Default = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], Get_semigroupDual())), Get_mkFoldRight1())
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_1_0, a_3)
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], __local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(x_2, a_6), b_5)
})
}))
})
})
})
	})
	return foldl1Default
}

var foldr1Default gopurs_runtime.Value
var once_foldr1Default sync.Once
func Get_foldr1Default() gopurs_runtime.Value {
	once_foldr1Default.Do(func() {
		foldr1Default = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], Get_foldRight1Semigroup()), Get_mkFoldRight1())
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_1_0, a_3)
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], __local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), b_2)
})
})
})
	})
	return foldr1Default
}

var intercalateMap gopurs_runtime.Value
var once_intercalateMap sync.Once
func Get_intercalateMap() gopurs_runtime.Value {
	once_intercalateMap.Do(func() {
		intercalateMap = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap12_2_0 := gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(v_2, j_4)), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], j_4), gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}))
return gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap12_2_0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_1 := gopurs_runtime.Apply(f_4, x_6)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_1
})
})), foldable_5), j_3)
})
})
})
})
})
	})
	return intercalateMap
}

var intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		intercalate = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap12_2_0 := gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(v_2, j_4)), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], j_4), gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}))
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap12_2_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
})), foldable_4), a_3)
})
})
})
})
	})
	return intercalate
}

var maximum gopurs_runtime.Value
var once_maximum sync.Once
func Get_maximum() gopurs_runtime.Value {
	once_maximum.Do(func() {
		maximum = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMax_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_1), v1_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = v1_2
} else {
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = v_1
} else {
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = v_1
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t2
})
})})
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], semigroupMax_1_0), pkg_Unsafe_Coerce.Get_unsafeCoerce()))
})
})
	})
	return maximum
}

var minimum gopurs_runtime.Value
var once_minimum sync.Once
func Get_minimum() gopurs_runtime.Value {
	once_minimum.Do(func() {
		minimum = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMin_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_1), v1_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = v_1
} else {
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = v_1
} else {
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = v1_2
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t2
})
})})
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], semigroupMin_1_0), pkg_Unsafe_Coerce.Get_unsafeCoerce()))
})
})
	})
	return minimum
}

var traverse1_ gopurs_runtime.Value
var once_traverse1_ sync.Once
func Get_traverse1_() gopurs_runtime.Value {
	once_traverse1_.Do(func() {
		traverse1_ = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictApply_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
foldMap12_3_1 := gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Apply.Get_applySecond(), dictApply_1), v_3), v1_4)
})
})}))
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap12_3_1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
})), t_5))
})
})
})
})
	})
	return traverse1_
}

var for1_ gopurs_runtime.Value
var once_for1_ sync.Once
func Get_for1_() gopurs_runtime.Value {
	once_for1_.Do(func() {
		for1_ = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traverse1_(), dictFoldable1_0), dictApply_1)
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0, a_4), b_3)
})
})
})
})
	})
	return for1_
}

var sequence1_ gopurs_runtime.Value
var once_sequence1_ sync.Once
func Get_sequence1_() gopurs_runtime.Value {
	once_sequence1_.Do(func() {
		sequence1_ = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traverse1_(), dictFoldable1_0), dictApply_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return sequence1_
}

var fold1 gopurs_runtime.Value
var once_fold1 sync.Once
func Get_fold1() gopurs_runtime.Value {
	once_fold1.Do(func() {
		fold1 = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], dictSemigroup_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return fold1
}


