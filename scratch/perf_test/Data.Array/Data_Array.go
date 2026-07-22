package Data_Array

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Array_ST_Iterator "gopurs/output/Data.Array.ST.Iterator"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var traverse_ gopurs_runtime.Value
var once_traverse_ sync.Once
func Get_traverse_() gopurs_runtime.Value {
	once_traverse_.Do(func() {
		traverse_ = gopurs_runtime.Apply(pkg_Data_Foldable.Get_traverse_(), pkg_Control_Monad_ST_Internal.Get_applicativeST())
	})
	return traverse_
}

var void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		void = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return void
}

var intercalate1 gopurs_runtime.Value
var once_intercalate1 sync.Once
func Get_intercalate1() gopurs_runtime.Value {
	once_intercalate1.Do(func() {
		intercalate1 = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_2_1 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(sep_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldlArray(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (v_5.PtrVal.(map[string]gopurs_runtime.Value)["init"]).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": v1_6})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["acc"]), gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], sep_3), v1_6))})
}
return __t2
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(true), "acc": mempty_2_1})), xs_4).PtrVal.(map[string]gopurs_runtime.Value)["acc"]
})
})
})
	})
	return intercalate1
}

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_zipWithImpl())
	})
	return zipWith
}

var zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		zipWithA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableArray().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_zipWith(), f_2), xs_3), ys_4))
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

var updateAtIndices gopurs_runtime.Value
var once_updateAtIndices sync.Once
func Get_updateAtIndices() gopurs_runtime.Value {
	once_updateAtIndices.Do(func() {
		updateAtIndices = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictFoldable_0)
return gopurs_runtime.Func(func(us_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(traverse_1_1_0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_poke(), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), res_4)
})), us_2)
})), xs_3))
})
})
})
	})
	return updateAtIndices
}

var updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		updateAt = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn5(), Get__updateAt()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return updateAt
}

var unsafeIndex gopurs_runtime.Value
var once_unsafeIndex sync.Once
func Get_unsafeIndex() gopurs_runtime.Value {
	once_unsafeIndex.Do(func() {
		unsafeIndex = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeIndexImpl())
})
	})
	return unsafeIndex
}

var unsafeIndex1 gopurs_runtime.Value
var once_unsafeIndex1 sync.Once
func Get_unsafeIndex1() gopurs_runtime.Value {
	once_unsafeIndex1.Do(func() {
		unsafeIndex1 = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeIndexImpl())
	})
	return unsafeIndex1
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unconsImpl()), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
})), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": x_0, "tail": xs_1})})
})
}))
	})
	return uncons
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
len_2_0 := gopurs_runtime.Apply(Get_length(), xs_1)
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], i_3), len_2_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeIndex1(), xs_1), i_3)
})), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), i_3), gopurs_runtime.Int(1))})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})), gopurs_runtime.Int(0))
})
})
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unconsImpl()), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
})), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": xs_1})
})
}))
	})
	return tail
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_sortByImpl()), comp_0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
} else {
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Int(0)), gopurs_runtime.Int(1))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
}))
})
	})
	return sortBy
}

var sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		sortWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(f_1, x_2)), gopurs_runtime.Apply(f_1, y_3))
})
}))
})
})
	})
	return sortWith
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

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), x_1)), xs_0))
})
})
	})
	return snoc
}

var slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		slice = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_sliceImpl())
	})
	return slice
}

var splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		splitAt = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"before": gopurs_runtime.Array([]gopurs_runtime.Value{}), "after": v1_1})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"before": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), gopurs_runtime.Int(0)), v_0), v1_1), "after": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), v_0), gopurs_runtime.Apply(Get_length(), v1_1)), v1_1)})
}
return __t0
})
})
	})
	return splitAt
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Array([]gopurs_runtime.Value{})
} else {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), gopurs_runtime.Int(0)), n_0), xs_1)
}
return __t0
})
})
	})
	return take
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{a_0})
})
	})
	return singleton
}

var scanr gopurs_runtime.Value
var once_scanr sync.Once
func Get_scanr() gopurs_runtime.Value {
	once_scanr.Do(func() {
		scanr = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_scanrImpl())
	})
	return scanr
}

var scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		scanl = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_scanlImpl())
	})
	return scanl
}

var replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		replicate = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_replicateImpl())
	})
	return replicate
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_rangeImpl())
	})
	return range_
}

var partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		partition = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_partitionImpl())
	})
	return partition
}

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(Get_length(), xs_0)), gopurs_runtime.Int(0))
})
	})
	return null
}

var modifyAtIndices gopurs_runtime.Value
var once_modifyAtIndices sync.Once
func Get_modifyAtIndices() gopurs_runtime.Value {
	once_modifyAtIndices.Do(func() {
		modifyAtIndices = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictFoldable_0)
return gopurs_runtime.Func(func(is_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(traverse_1_1_0, gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_modify(), i_6), f_3), res_5)
})), is_2)
})), xs_4))
})
})
})
})
	})
	return modifyAtIndices
}

var mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		mapWithIndex = pkg_Data_FunctorWithIndex.Get_mapWithIndexArray()
	})
	return mapWithIndex
}

var intersperse gopurs_runtime.Value
var once_intersperse sync.Once
func Get_intersperse() gopurs_runtime.Value {
	once_intersperse.Do(func() {
		intersperse = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(arr_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(Get_length(), arr_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2_0), gopurs_runtime.Int(2)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = arr_1
} else {
__t1 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Data_Array_ST.Get_new_()), gopurs_runtime.Func(func(out_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeIndex1(), arr_1), gopurs_runtime.Int(0))
}))), out_3)), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_for_(), gopurs_runtime.Int(1)), v_2_0), gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), a_0), out_3)), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeIndex1(), arr_1), idx_5)
}))), out_3))
}))
}))), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), out_3)
}))
}))
}))), pkg_Data_Array_ST.Get_unsafeFreeze()))
}
return __t1
})
})
	})
	return intersperse
}

var intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		intercalate = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_intercalate1(), dictMonoid_0)
})
	})
	return intercalate
}

var insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		insertAt = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn5(), Get__insertAt()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return insertAt
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(Get_length(), xs_0)), gopurs_runtime.Int(0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), gopurs_runtime.Int(0)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), xs_0)), gopurs_runtime.Int(1))), xs_0)})
}
return __t0
})
	})
	return init_
}

var index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		index = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_indexImpl()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return index
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), xs_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), xs_0)), gopurs_runtime.Int(1)))
})
	})
	return last
}

var unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		unsnoc = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_init_(), xs_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), xs_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), xs_0)), gopurs_runtime.Int(1)))
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "last": __local_var_2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t1 = __t3
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
	})
	return unsnoc
}

var modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		modifyAt = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), xs_2), i_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_updateAt(), i_0), gopurs_runtime.Apply(f_1, __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), xs_2)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
	})
	return modifyAt
}

var span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		span = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(arr_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), arr_1), i_3)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(p_0, v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
__t3 = gopurs_runtime.Apply(go__2_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), i_3), gopurs_runtime.Int(1)))
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": i_3})
}
__t2 = __t3
} else {
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
breakIndex_3_4 := gopurs_runtime.Apply(go__2_0, gopurs_runtime.Int(0))
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(breakIndex_3_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(breakIndex_3_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Array([]gopurs_runtime.Value{}), "rest": arr_1})
} else {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), gopurs_runtime.Int(0)), breakIndex_3_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), arr_1), "rest": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), breakIndex_3_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(Get_length(), arr_1)), arr_1)})
}
__t5 = __t6
} else {
if (gopurs_runtime.Bool(breakIndex_3_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": arr_1, "rest": gopurs_runtime.Array([]gopurs_runtime.Value{})})
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})
	})
	return span
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_span(), p_0), xs_1).PtrVal.(map[string]gopurs_runtime.Value)["init"]
})
})
	})
	return takeWhile
}

var unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		unzip = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Data_Array_ST.Get_new_()), gopurs_runtime.Func(func(fsts_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Data_Array_ST.Get_new_()), gopurs_runtime.Func(func(snds_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), xs_0), v_3)
}))), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_new_(), gopurs_runtime.Int(0)))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_iterate(), iter_3), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), fsts_1))), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), __local_var_5_0), snds_2))
}))
}))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), fsts_1)), gopurs_runtime.Func(func(fsts_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), snds_2)), gopurs_runtime.Func(func(snds_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": fsts_prime_5, "value1": snds_prime_6}))
}))
}))
}))
}))
}))
})))
})
	})
	return unzip
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), xs_0), gopurs_runtime.Int(0))
})
	})
	return head
}

var nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		nubBy = gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
indexedAndSorted_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(comp_0, x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_FunctorWithIndex.Get_mapWithIndexArray(), pkg_Data_Tuple.Get_Tuple()), xs_1))
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), indexedAndSorted_2_0), gopurs_runtime.Int(0))
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Array([]gopurs_runtime.Value{})
} else {
if (gopurs_runtime.Bool(v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), pkg_Data_Tuple.Get_snd()), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_sortWith(), pkg_Data_Ord.Get_ordInt()), pkg_Data_Tuple.Get_fst()), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeThaw(), gopurs_runtime.Array([]gopurs_runtime.Value{v_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))), gopurs_runtime.Func(func(result_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_foreach(), indexedAndSorted_2_0), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__local_var_7_4 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), x_8), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), x_8)), gopurs_runtime.Int(1)))
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_9_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = __local_var_9_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t6
})
}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, x_8).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), result_4))), gopurs_runtime.Func(func(lst_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_7 := gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), v1_5), result_4))
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqBooleanImpl(), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(comp_0, lst_7), __local_var_6_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")), gopurs_runtime.Bool(false))).IntVal != 0 {
__t8 = __local_var_8_7
} else {
__t8 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), pkg_Data_Unit.Get_unit())
}
return __t8
}))
}))), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), result_4)
}))
})))))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
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

var groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		groupBy = gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Data_Array_ST.Get_new_()), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), xs_1), v_3)
}))), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_new_(), gopurs_runtime.Int(0)))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_iterate(), iter_3), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Data_Array_ST.Get_new_()), gopurs_runtime.Func(func(sub1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), x_4), sub1_5)), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_pushWhile(), gopurs_runtime.Apply(op_0, x_4)), iter_3), sub1_5)), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), sub1_5)), gopurs_runtime.Func(func(grp_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), grp_8), result_2)
}))
}))
}))
})))
}))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), result_2)
}))
}))
})))
})
})
	})
	return groupBy
}

var groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		groupAllBy = gopurs_runtime.Func(func(cmp_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_0, x_1), y_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")
})
}))
__local_var_2_1 := gopurs_runtime.Apply(Get_sortBy(), cmp_0)
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
})
	})
	return groupAllBy
}

var groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		groupAll = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_groupAllBy(), dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"])
})
	})
	return groupAll
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
eq2_1_0 := dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_groupBy(), eq2_1_0), xs_2)
})
})
	})
	return group
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_fromFoldableImpl()), dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"])
})
	})
	return fromFoldable
}

var foldr gopurs_runtime.Value
var once_foldr sync.Once
func Get_foldr() gopurs_runtime.Value {
	once_foldr.Do(func() {
		foldr = pkg_Data_Foldable.Get_foldrArray()
	})
	return foldr
}

var foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		foldl = pkg_Data_Foldable.Get_foldlArray()
	})
	return foldl
}

var transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		transpose = gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(idx_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(allArrays_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldlArray(), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), nextArr_5), idx_2)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = acc_4
} else {
if (gopurs_runtime.Bool(__local_var_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(acc_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Array([]gopurs_runtime.Value{__local_var_6_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(acc_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), __local_var_6_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), acc_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __t4})
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t3
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), xs_0)
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = allArrays_3
} else {
if (gopurs_runtime.Bool(v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), idx_2), gopurs_runtime.Int(1))), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), v_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), allArrays_3)))
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Int(0)), gopurs_runtime.Array([]gopurs_runtime.Value{}))
})
	})
	return transpose
}

var foldRecM gopurs_runtime.Value
var once_foldRecM sync.Once
func Get_foldRecM() gopurs_runtime.Value {
	once_foldRecM.Do(func() {
		foldRecM = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(array_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(o_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], o_6.PtrVal.(map[string]gopurs_runtime.Value)["b"]), gopurs_runtime.Apply(Get_length(), array_5)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": o_6.PtrVal.(map[string]gopurs_runtime.Value)["a"]}))
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, o_6.PtrVal.(map[string]gopurs_runtime.Value)["a"]), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeIndex1(), array_5), o_6.PtrVal.(map[string]gopurs_runtime.Value)["b"])
})))), gopurs_runtime.Func(func(res_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": res_prime_7, "b": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), o_6.PtrVal.(map[string]gopurs_runtime.Value)["b"]), gopurs_runtime.Int(1))})}))
}))
}
return __t2
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": b_4, "b": gopurs_runtime.Int(0)}))
})
})
})
})
	})
	return foldRecM
}

var foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		foldMap = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableArray().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_0)
})
	})
	return foldMap
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_unconsImpl()), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], b_2)
})), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, b_2), a_3)), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldM(), dictMonad_0), f_1), b_prime_5), as_4)
}))
})
}))
}
}()
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
		fold = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableArray().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return fold
}

var findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		findMap = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_findMapImpl()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), pkg_Data_Maybe.Get_isJust())
	})
	return findMap
}

var findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		findLastIndex = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_findLastIndexImpl()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return findLastIndex
}

var insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		insertBy = gopurs_runtime.Func(func(cmp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(cmp_0, x_1), y_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")
})), ys_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
i_4_1 := __t2
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_insertAt(), i_4_1), x_1), ys_2)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = __local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t4
}))
})
})
})
	})
	return insertBy
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

var findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		findIndex = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_findIndexImpl()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return findIndex
}

var find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		find = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unsafeIndex1(), xs_1)
}))
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_findIndex(), f_0), xs_1)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_2_0, __local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t2
})
})
	})
	return find
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_filterImpl())
	})
	return filter
}

var intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		intersectBy = gopurs_runtime.Func(func(eq2_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_filter(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Apply(eq2_0, x_3)), ys_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(__local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})), xs_1)
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

var notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		notElem = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_3), a_1)
})), arr_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
} else {
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
	})
	return notElem
}

var elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		elem = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_3), a_1)
})), arr_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
	})
	return elem
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_span(), p_0), xs_1).PtrVal.(map[string]gopurs_runtime.Value)["rest"]
})
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
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), xs_1)), n_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_2_0), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{})
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), gopurs_runtime.Int(0)), __local_var_2_0), xs_1)
}
return __t1
})
})
	})
	return dropEnd
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = xs_1
} else {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_slice(), n_0), gopurs_runtime.Apply(Get_length(), xs_1)), xs_1)
}
return __t0
})
})
	})
	return drop
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

var deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		deleteAt = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get__deleteAt()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return deleteAt
}

var deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		deleteBy = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Int(len(v2_2.PtrVal.([]gopurs_runtime.Value))).IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Array([]gopurs_runtime.Value{})
} else {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Apply(v_0, v1_1)), v2_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = v2_2
} else {
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_4_2 := __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t1 = gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_deleteAt(), __local_var_4_2), v2_2)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = __local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t4
}))
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t5 = __t1
}
return __t5
})
})
})
	})
	return deleteBy
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
return gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldrArray(), gopurs_runtime.Apply(Get_delete_(), dictEq_0))
})
	})
	return difference
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0})), xs_1)
})
})
	})
	return cons
}

var some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		some = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictLazy_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_cons()), v_2)), gopurs_runtime.Apply(dictLazy_1.PtrVal.(map[string]gopurs_runtime.Value)["defer"], gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_some(), dictAlternative_0), dictLazy_1), v_2)), gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Array([]gopurs_runtime.Value{})))
})
})
})
	})
	return many
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func(func(b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Bind.Get_arrayBind(), a_1), b_0)
})
})
	})
	return concatMap
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, x_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{})
} else {
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
}))
})
	})
	return mapMaybe
}

var filterA gopurs_runtime.Value
var once_filterA sync.Once
func Get_filterA() gopurs_runtime.Value {
	once_filterA.Do(func() {
		filterA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_1_0 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableArray().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0)
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(traverse1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), x_4)), gopurs_runtime.Apply(p_3, x_4))
}))
__local_var_5_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(Get_mapMaybe(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t4
})))
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(__local_var_4_2, x_6))
})
})
})
	})
	return filterA
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Apply(Get_mapMaybe(), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
	})
	return catMaybes
}

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_anyImpl())
	})
	return any
}

var nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		nubByEq = gopurs_runtime.Func(func(eq2_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Data_Array_ST.Get_new_()), gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_foreach(), xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(Get_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(eq2_0, v_4), x_3)
}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(__local_var_4_0, x_5))
})), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), arr_2))), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), x_3), arr_2))
var __t2 gopurs_runtime.Value
if (e_4).IntVal != 0 {
__t2 = __local_var_5_1
} else {
__t2 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), pkg_Data_Unit.Get_unit())
}
return __t2
}))
}))), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Array_ST.Get_unsafeFreeze(), arr_2)
}))
})))
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

var unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		unionBy = gopurs_runtime.Func(func(eq2_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), xs_1), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldlArray(), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_deleteBy(), eq2_0), a_4), b_3)
})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_nubByEq(), eq2_0), ys_2)), xs_1))
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

var alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		alterAt = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_index(), xs_2), i_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
v_4_2 := gopurs_runtime.Apply(f_1, __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_deleteAt(), i_0), xs_2)
} else {
if (gopurs_runtime.Bool(v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_updateAt(), i_0), v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), xs_2)
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
})
	})
	return alterAt
}

var all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		all = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_allImpl())
	})
	return all
}

func Get__deleteAt() gopurs_runtime.Value {
	return X_DeleteAt
}

func Get__insertAt() gopurs_runtime.Value {
	return X_InsertAt
}

func Get__updateAt() gopurs_runtime.Value {
	return X_UpdateAt
}

func Get_allImpl() gopurs_runtime.Value {
	return AllImpl
}

func Get_anyImpl() gopurs_runtime.Value {
	return AnyImpl
}

func Get_concat() gopurs_runtime.Value {
	return Concat
}

func Get_filterImpl() gopurs_runtime.Value {
	return FilterImpl
}

func Get_findIndexImpl() gopurs_runtime.Value {
	return FindIndexImpl
}

func Get_findLastIndexImpl() gopurs_runtime.Value {
	return FindLastIndexImpl
}

func Get_findMapImpl() gopurs_runtime.Value {
	return FindMapImpl
}

func Get_fromFoldableImpl() gopurs_runtime.Value {
	return FromFoldableImpl
}

func Get_indexImpl() gopurs_runtime.Value {
	return IndexImpl
}

func Get_length() gopurs_runtime.Value {
	return Length
}

func Get_partitionImpl() gopurs_runtime.Value {
	return PartitionImpl
}

func Get_rangeImpl() gopurs_runtime.Value {
	return RangeImpl
}

func Get_replicateImpl() gopurs_runtime.Value {
	return ReplicateImpl
}

func Get_reverse() gopurs_runtime.Value {
	return Reverse
}

func Get_scanlImpl() gopurs_runtime.Value {
	return ScanlImpl
}

func Get_scanrImpl() gopurs_runtime.Value {
	return ScanrImpl
}

func Get_sliceImpl() gopurs_runtime.Value {
	return SliceImpl
}

func Get_sortByImpl() gopurs_runtime.Value {
	return SortByImpl
}

func Get_unconsImpl() gopurs_runtime.Value {
	return UnconsImpl
}

func Get_unsafeIndexImpl() gopurs_runtime.Value {
	return UnsafeIndexImpl
}

func Get_zipWithImpl() gopurs_runtime.Value {
	return ZipWithImpl
}
