package Control_Monad_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup_Last "gopurs/output/Data.Semigroup.Last"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var monoidAdditive gopurs_runtime.Value
var once_monoidAdditive sync.Once
func Get_monoidAdditive() gopurs_runtime.Value {
	once_monoidAdditive.Do(func() {
		monoidAdditive = func() gopurs_runtime.Value {
semigroupAdditive1_0_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), v_0), v1_1)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Float(0.0), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
})})
}()
	})
	return monoidAdditive
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

var Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		Nil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
	})
	return Nil
}

var unfoldable gopurs_runtime.Value
var once_unfoldable sync.Once
func Get_unfoldable() gopurs_runtime.Value {
	once_unfoldable.Do(func() {
		unfoldable = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
__local_var_3_1 := gopurs_runtime.Apply(Monad0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
Bind1_4_2 := gopurs_runtime.Apply(Monad0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictUnfoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t7 = gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
} else {
__local_var_8_5 := v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_9_6 := v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gen_6), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": x_10, "value1": __local_var_8_5}), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), __local_var_9_6), gopurs_runtime.Int(1))})}))
}))
}
return __t7
}))
__local_var_8_8 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(dictUnfoldable_5.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
}))), gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["sized"], gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(__local_var_8_8, x_9))
})))
})
})
})
})
	})
	return unfoldable
}

var semigroupFreqSemigroup gopurs_runtime.Value
var once_semigroupFreqSemigroup sync.Once
func Get_semigroupFreqSemigroup() gopurs_runtime.Value {
	once_semigroupFreqSemigroup.Do(func() {
		semigroupFreqSemigroup = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.Apply(v_0, pos_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(v1_1, v2_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = v2_3_0
}
return __t1
})
})
})})
	})
	return semigroupFreqSemigroup
}

var fromIndex gopurs_runtime.Value
var once_fromIndex sync.Once
func Get_fromIndex() gopurs_runtime.Value {
	once_fromIndex.Do(func() {
		fromIndex = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], pkg_Data_Semigroup_Last.Get_semigroupLast())
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_5), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t3 = v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_1, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_5), gopurs_runtime.Int(1))), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
}
__t2 = __t3
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_1_0, pkg_Data_Semigroup_Last.Get_Last()), xs_3))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_1, i_2), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["foldr"], Get_Cons()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), xs_3))
})
})
})
	})
	return fromIndex
}

var oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		oneOf = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
length_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), c_2)
})
})), gopurs_runtime.Int(0))
fromIndex1_3_1 := gopurs_runtime.Apply(Get_fromIndex(), dictFoldable1_1)
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(0)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(length_2_0, xs_4)), gopurs_runtime.Int(1)))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(fromIndex1_3_1, n_5), xs_4)
}))
})
})
})
	})
	return oneOf
}

var freqSemigroup gopurs_runtime.Value
var once_freqSemigroup sync.Once
func Get_freqSemigroup() gopurs_runtime.Value {
	once_freqSemigroup.Do(func() {
		freqSemigroup = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__local_var_2_1 := v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], pos_3), __local_var_1_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), pos_3), __local_var_1_0)}), "value1": __local_var_2_1})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "value1": __local_var_2_1})
}
return __t2
})
})
	})
	return freqSemigroup
}

var frequency gopurs_runtime.Value
var once_frequency sync.Once
func Get_frequency() gopurs_runtime.Value {
	once_frequency.Do(func() {
		frequency = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], Get_monoidAdditive())
foldMap1_3_1 := gopurs_runtime.Apply(dictFoldable1_1.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], Get_semigroupFreqSemigroup())
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_3_1, Get_freqSemigroup()), xs_4)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseFloat"], gopurs_runtime.Float(0.0)), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), foldMap_2_0), pkg_Data_Tuple.Get_fst()), xs_4))), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, x_6).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
}))
})
})
})
	})
	return frequency
}

var filtered gopurs_runtime.Value
var once_filtered sync.Once
func Get_filtered() gopurs_runtime.Value {
	once_filtered.Do(func() {
		filtered = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(a_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": pkg_Data_Unit.Get_unit()})
} else {
if (gopurs_runtime.Bool(a_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": a_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})), gen_3)
})), pkg_Data_Unit.Get_unit())
})
})
})
	})
	return filtered
}

var suchThat gopurs_runtime.Value
var once_suchThat sync.Once
func Get_suchThat() gopurs_runtime.Value {
	once_suchThat.Do(func() {
		suchThat = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
filtered2_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_filtered(), dictMonadRec_0), dictMonadGen_1)
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(pred_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(filtered2_2_0, gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(pred_5, a_6)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": a_6})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t2
})), gen_4))
})
})
})
})
	})
	return suchThat
}

var elements gopurs_runtime.Value
var once_elements sync.Once
func Get_elements() gopurs_runtime.Value {
	once_elements.Do(func() {
		elements = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
length_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_2.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), c_3)
})
})), gopurs_runtime.Int(0))
fromIndex1_4_2 := gopurs_runtime.Apply(Get_fromIndex(), dictFoldable1_2)
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(0)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(length_3_1, xs_5)), gopurs_runtime.Int(1)))), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(gopurs_runtime.Apply(fromIndex1_4_2, n_6), xs_5))
}))
})
})
})
	})
	return elements
}

var choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		choose = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
chooseBool_1_0 := dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseBool"]
return gopurs_runtime.Func(func(genA_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], chooseBool_1_0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (v_4).IntVal != 0 {
__t1 = genA_2
} else {
__t1 = genB_3
}
return __t1
}))
})
})
})
	})
	return choose
}


