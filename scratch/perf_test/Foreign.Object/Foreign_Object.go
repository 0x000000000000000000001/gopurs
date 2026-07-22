package Foreign_Object

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Foreign_Object_ST "gopurs/output/Foreign.Object.ST"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var forWithIndex_ gopurs_runtime.Value
var once_forWithIndex_ sync.Once
func Get_forWithIndex_() gopurs_runtime.Value {
	once_forWithIndex_.Do(func() {
		forWithIndex_ = gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_forWithIndex_(), pkg_Control_Monad_ST_Internal.Get_applicativeST())
	})
	return forWithIndex_
}

var for_ gopurs_runtime.Value
var once_for_ sync.Once
func Get_for_() gopurs_runtime.Value {
	once_for_.Do(func() {
		for_ = gopurs_runtime.Apply(pkg_Data_Foldable.Get_for_(), pkg_Control_Monad_ST_Internal.Get_applicativeST())
	})
	return for_
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

var ordTuple gopurs_runtime.Value
var once_ordTuple sync.Once
func Get_ordTuple() gopurs_runtime.Value {
	once_ordTuple.Do(func() {
		ordTuple = gopurs_runtime.Apply(pkg_Data_Tuple.Get_ordTuple(), pkg_Data_Ord.Get_ordString())
	})
	return ordTuple
}

var values gopurs_runtime.Value
var once_values sync.Once
func Get_values() gopurs_runtime.Value {
	once_values.Do(func() {
		values = gopurs_runtime.Apply(Get_toArrayWithKey(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return values
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_toArrayWithKey(), pkg_Data_Tuple.Get_Tuple())
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array.Get_toUnfoldable(), dictUnfoldable_0), gopurs_runtime.Apply(__local_var_1_0, x_2))
})
})
	})
	return toUnfoldable
}

var toAscUnfoldable gopurs_runtime.Value
var once_toAscUnfoldable sync.Once
func Get_toAscUnfoldable() gopurs_runtime.Value {
	once_toAscUnfoldable.Do(func() {
		toAscUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array.Get_sortWith(), pkg_Data_Ord.Get_ordString()), pkg_Data_Tuple.Get_fst())
__local_var_2_2 := gopurs_runtime.Apply(Get_toArrayWithKey(), pkg_Data_Tuple.Get_Tuple())
__local_var_2_1 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_2, x_3))
})
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array.Get_toUnfoldable(), dictUnfoldable_0), gopurs_runtime.Apply(__local_var_2_1, x_3))
})
})
	})
	return toAscUnfoldable
}

var toAscArray gopurs_runtime.Value
var once_toAscArray sync.Once
func Get_toAscArray() gopurs_runtime.Value {
	once_toAscArray.Do(func() {
		toAscArray = gopurs_runtime.Apply(Get_toAscUnfoldable(), pkg_Data_Unfoldable.Get_unfoldableArray())
	})
	return toAscArray
}

var toArray gopurs_runtime.Value
var once_toArray sync.Once
func Get_toArray() gopurs_runtime.Value {
	once_toArray.Do(func() {
		toArray = gopurs_runtime.Apply(Get_toArrayWithKey(), pkg_Data_Tuple.Get_Tuple())
	})
	return toArray
}

var thawST gopurs_runtime.Value
var once_thawST sync.Once
func Get_thawST() gopurs_runtime.Value {
	once_thawST.Do(func() {
		thawST = Get__copyST()
	})
	return thawST
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_runST(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Foreign_Object_ST.Get_new_()), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_0), v_1)))
})
})
	})
	return singleton
}

var showObject gopurs_runtime.Value
var once_showObject sync.Once
func Get_showObject() gopurs_runtime.Value {
	once_showObject.Do(func() {
		showObject = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
show_1_0 := gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Tuple.Get_showTuple(), pkg_Data_Show.Get_showString()), dictShow_0).PtrVal.(map[string]gopurs_runtime.Value)["show"])
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(fromFoldable ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(show_1_0, gopurs_runtime.Apply(Get_toArray(), m_2))), gopurs_runtime.Str(")")))
})})
})
	})
	return showObject
}

var mutate gopurs_runtime.Value
var once_mutate sync.Once
func Get_mutate() gopurs_runtime.Value {
	once_mutate.Do(func() {
		mutate = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_runST(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(Get__copyST(), m_1)), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(f_0, s_2)), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), s_2)
}))
})))
})
})
	})
	return mutate
}

var member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		member = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get__lookup()), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return member
}

var mapWithKey gopurs_runtime.Value
var once_mapWithKey sync.Once
func Get_mapWithKey() gopurs_runtime.Value {
	once_mapWithKey.Do(func() {
		mapWithKey = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get__mapWithKey()), m_1), f_0)
})
})
	})
	return mapWithKey
}

var lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		lookup = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get__lookup()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), pkg_Data_Maybe.Get_Just())
	})
	return lookup
}

var isSubmap gopurs_runtime.Value
var once_isSubmap sync.Once
func Get_isSubmap() gopurs_runtime.Value {
	once_isSubmap.Do(func() {
		isSubmap = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_all(), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get__lookup()), gopurs_runtime.Bool(false)), gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_4)), k_3), m2_2)
})
})), m1_1)
})
})
})
	})
	return isSubmap
}

var isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		isEmpty = gopurs_runtime.Apply(Get_all(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
}))
	})
	return isEmpty
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_0), v_1))
})
})
	})
	return insert
}

var functorObject gopurs_runtime.Value
var once_functorObject sync.Once
func Get_functorObject() gopurs_runtime.Value {
	once_functorObject.Do(func() {
		functorObject = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get__fmapObject()), m_1), f_0)
})
})})
	})
	return functorObject
}

var functorWithIndexObject gopurs_runtime.Value
var once_functorWithIndexObject sync.Once
func Get_functorWithIndexObject() gopurs_runtime.Value {
	once_functorWithIndexObject.Do(func() {
		functorWithIndexObject = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": Get_mapWithKey(), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorObject()
})})
	})
	return functorWithIndexObject
}

var fromHomogeneous gopurs_runtime.Value
var once_fromHomogeneous sync.Once
func Get_fromHomogeneous() gopurs_runtime.Value {
	once_fromHomogeneous.Do(func() {
		fromHomogeneous = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
})
	})
	return fromHomogeneous
}

var fromFoldableWithIndex gopurs_runtime.Value
var once_fromFoldableWithIndex sync.Once
func Get_fromFoldableWithIndex() gopurs_runtime.Value {
	once_fromFoldableWithIndex.Do(func() {
		fromFoldableWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
forWithIndex_1_1_0 := gopurs_runtime.Apply(Get_forWithIndex_(), dictFoldableWithIndex_0)
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_runST(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Foreign_Object_ST.Get_new_()), gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(forWithIndex_1_1_0, l_2), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_4), v_5), s_3)
})
}))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), s_3)
}))
})))
})
})
	})
	return fromFoldableWithIndex
}

var fromFoldableWith gopurs_runtime.Value
var once_fromFoldableWith sync.Once
func Get_fromFoldableWith() gopurs_runtime.Value {
	once_fromFoldableWith.Do(func() {
		fromFoldableWith = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
for_1_1_0 := gopurs_runtime.Apply(Get_for_(), dictFoldable_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_runST(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Foreign_Object_ST.Get_new_()), gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(for_1_1_0, l_3), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get__lookupST()), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Apply(f_2, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), __local_var_6_1), s_4)), gopurs_runtime.Func(func(v_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), __local_var_6_1), v_prime_7), s_4)
}))
}))), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), s_4)
}))
})))
})
})
})
	})
	return fromFoldableWith
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
fromFoldable1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), pkg_Data_Array.Get_fromFoldableImpl()), dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"])
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_runST(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Foreign_Object_ST.Get_new_()), gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_foreach(), gopurs_runtime.Apply(fromFoldable1_1_0, l_2)), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), s_3))
}))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), s_3)
}))
})))
})
})
	})
	return fromFoldable
}

var freezeST gopurs_runtime.Value
var once_freezeST sync.Once
func Get_freezeST() gopurs_runtime.Value {
	once_freezeST.Do(func() {
		freezeST = Get__copyST()
	})
	return freezeST
}

var foldMaybe gopurs_runtime.Value
var once_foldMaybe sync.Once
func Get_foldMaybe() gopurs_runtime.Value {
	once_foldMaybe.Do(func() {
		foldMaybe = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get__foldSCObject()), m_2), z_1), f_0), pkg_Data_Maybe.Get_fromMaybe())
})
})
})
	})
	return foldMaybe
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
bind1_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get__foldM(), bind1_1_0), f_2), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], z_3))
})
})
})
	})
	return foldM
}

var foldM1 gopurs_runtime.Value
var once_foldM1 sync.Once
func Get_foldM1() gopurs_runtime.Value {
	once_foldM1.Do(func() {
		foldM1 = gopurs_runtime.Apply(Get_foldM(), pkg_Control_Monad_ST_Internal.Get_monadST())
	})
	return foldM1
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldM1(), gopurs_runtime.Func(func(s_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_3), v_4), s_prime_2)
})
})
})), s_1), m_0)
}))
})
	})
	return union
}

var unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		unions = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], Get_union()), Get_empty())
})
	})
	return unions
}

var unionWith gopurs_runtime.Value
var once_unionWith sync.Once
func Get_unionWith() gopurs_runtime.Value {
	once_unionWith.Do(func() {
		unionWith = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Func(func(s1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldM1(), gopurs_runtime.Func(func(s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_5), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get__lookup()), v1_6), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v1_6), v2_7)
})), k_5), m2_2)), s2_4)
})
})
})), s1_3), m1_1)
})), m2_2)
})
})
})
	})
	return unionWith
}

var semigroupObject gopurs_runtime.Value
var once_semigroupObject sync.Once
func Get_semigroupObject() gopurs_runtime.Value {
	once_semigroupObject.Do(func() {
		semigroupObject = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_unionWith(), dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"])})
})
	})
	return semigroupObject
}

var monoidObject gopurs_runtime.Value
var once_monoidObject sync.Once
func Get_monoidObject() gopurs_runtime.Value {
	once_monoidObject.Do(func() {
		monoidObject = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": Get_empty(), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_unionWith(), dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"])})
})})
})
	})
	return monoidObject
}

var fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		fold = gopurs_runtime.Apply(Get__foldM(), pkg_Data_Function.Get_applyFlipped())
	})
	return fold
}

var foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		foldMap = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fold(), gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], acc_3), gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, k_4), v_5))
})
})
})), mempty_1_0)
})
})
	})
	return foldMap
}

var foldableObject gopurs_runtime.Value
var once_foldableObject sync.Once
func Get_foldableObject() gopurs_runtime.Value {
	once_foldableObject.Do(func() {
		foldableObject = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_fold(), gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, z_1)
})
}))
}), "foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldrArray(), f_0), z_1), gopurs_runtime.Apply(Get_values(), m_2))
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(Get_foldMap(), dictMonoid_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap1_1_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return f_2
}))
})
})})
	})
	return foldableObject
}

var foldableWithIndexObject gopurs_runtime.Value
var once_foldableWithIndexObject sync.Once
func Get_foldableWithIndexObject() gopurs_runtime.Value {
	once_foldableWithIndexObject.Do(func() {
		foldableWithIndexObject = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldlWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_fold(), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, a_2), b_1)
})
}))
}), "foldrWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldrArray(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})), z_1), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_toArrayWithKey(), pkg_Data_Tuple.Get_Tuple()), m_2))
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_foldMap(), dictMonoid_0)
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableObject()
})})
	})
	return foldableWithIndexObject
}

var traversableWithIndexObject gopurs_runtime.Value
var once_traversableWithIndexObject sync.Once
func Get_traversableWithIndexObject() gopurs_runtime.Value {
	once_traversableWithIndexObject.Do(func() {
		traversableWithIndexObject = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ms_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fold(), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_5), a_8)), b_7)
})
})), acc_4)), gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, k_5), v_6))
})
})
})), gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], Get_empty())), ms_3)
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexObject()
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexObject()
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableObject()
})})
	})
	return traversableWithIndexObject
}

var traversableObject gopurs_runtime.Value
var once_traversableObject sync.Once
func Get_traversableObject() gopurs_runtime.Value {
	once_traversableObject.Do(func() {
		traversableObject = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_traversableWithIndexObject().PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_0)
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
})
}), "sequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableObject().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorObject()
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableObject()
})})
	})
	return traversableObject
}

var filterWithKey gopurs_runtime.Value
var once_filterWithKey sync.Once
func Get_filterWithKey() gopurs_runtime.Value {
	once_filterWithKey.Do(func() {
		filterWithKey = gopurs_runtime.Func(func(predicate_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_runST(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), pkg_Foreign_Object_ST.Get_new_()), gopurs_runtime.Func(func(m_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldM1(), gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(predicate_0, k_4), v_5)).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_4), v_5), acc_3)
} else {
__t0 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), acc_3)
}
return __t0
})
})
})), m_prime_2), m_1)
})))
})
})
	})
	return filterWithKey
}

var filterKeys gopurs_runtime.Value
var once_filterKeys sync.Once
func Get_filterKeys() gopurs_runtime.Value {
	once_filterKeys.Do(func() {
		filterKeys = gopurs_runtime.Func(func(predicate_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_filterWithKey(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(predicate_0, x_1)
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_0
})
}))
})
	})
	return filterKeys
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(predicate_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_filterWithKey(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return predicate_0
}))
})
	})
	return filter
}

var eqObject gopurs_runtime.Value
var once_eqObject sync.Once
func Get_eqObject() gopurs_runtime.Value {
	once_eqObject.Do(func() {
		eqObject = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(m1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_isSubmap(), dictEq_0), m1_1), m2_2)), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_isSubmap(), dictEq_0), m2_2), m1_1))
})
})})
})
	})
	return eqObject
}

var ordObject gopurs_runtime.Value
var once_ordObject sync.Once
func Get_ordObject() gopurs_runtime.Value {
	once_ordObject.Do(func() {
		ordObject = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqObject1_1_0 := gopurs_runtime.Apply(Get_eqObject(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordArray(), gopurs_runtime.Apply(Get_ordTuple(), dictOrd_0)).PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(Get_toAscArray(), m1_2)), gopurs_runtime.Apply(Get_toAscArray(), m2_3))
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqObject1_1_0
})})
})
	})
	return ordObject
}

var eq1Object gopurs_runtime.Value
var once_eq1Object sync.Once
func Get_eq1Object() gopurs_runtime.Value {
	once_eq1Object.Do(func() {
		eq1Object = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_eqObject(), dictEq_0).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1Object
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_delete_(), k_0))
})
	})
	return delete_
}

var pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		pop = gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_lookup(), k_0), m_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_delete_(), k_0)), m_1)})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
	})
	return pop
}

var alter gopurs_runtime.Value
var once_alter sync.Once
func Get_alter() gopurs_runtime.Value {
	once_alter.Do(func() {
		alter = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(gopurs_runtime.Apply(Get_lookup(), k_1), m_2))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_delete_(), k_1)), m_2)
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mutate(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Foreign_Object_ST.Get_poke(), k_1), v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), m_2)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
	})
	return alter
}

var update gopurs_runtime.Value
var once_update sync.Once
func Get_update() gopurs_runtime.Value {
	once_update.Do(func() {
		update = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_alter(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_0, v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})), k_1), m_2)
})
})
})
	})
	return update
}

func Get__copyST() gopurs_runtime.Value {
	return X_CopyST
}

func Get__fmapObject() gopurs_runtime.Value {
	return X_FmapObject
}

func Get__foldM() gopurs_runtime.Value {
	return X_FoldM
}

func Get__foldSCObject() gopurs_runtime.Value {
	return X_FoldSCObject
}

func Get__lookup() gopurs_runtime.Value {
	return X_Lookup
}

func Get__lookupST() gopurs_runtime.Value {
	return X_LookupST
}

func Get__mapWithKey() gopurs_runtime.Value {
	return X_MapWithKey
}

func Get_all() gopurs_runtime.Value {
	return All
}

func Get_empty() gopurs_runtime.Value {
	return Empty
}

func Get_keys() gopurs_runtime.Value {
	return Keys
}

func Get_runST() gopurs_runtime.Value {
	return RunST
}

func Get_size() gopurs_runtime.Value {
	return Size
}

func Get_toArrayWithKey() gopurs_runtime.Value {
	return ToArrayWithKey
}
