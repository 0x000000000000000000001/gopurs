package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Set_NonEmpty_coerce gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce sync.Once

func Get_Data_Set_NonEmpty_coerce() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce.Do(func() {
		cache_Data_Set_NonEmpty_coerce = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce
}

var cache_Data_Set_NonEmpty_toUnfoldable2 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable2 sync.Once

func Get_Data_Set_NonEmpty_toUnfoldable2() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable2.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable2 = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_toUnfoldable(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList())))), Get_Data_Set_toList())
	})
	return cache_Data_Set_NonEmpty_toUnfoldable2
}

var cache_Data_Set_NonEmpty_coerce1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce1 sync.Once

func Get_Data_Set_NonEmpty_coerce1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce1.Do(func() {
		cache_Data_Set_NonEmpty_coerce1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce1
}

var cache_Data_Set_NonEmpty_coerce2 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce2 sync.Once

func Get_Data_Set_NonEmpty_coerce2() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce2.Do(func() {
		cache_Data_Set_NonEmpty_coerce2 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce2
}

var cache_Data_Set_NonEmpty_go__append gopurs_runtime.Value
var once_Data_Set_NonEmpty_go__append sync.Once

func Get_Data_Set_NonEmpty_go__append() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_go__append.Do(func() {
		cache_Data_Set_NonEmpty_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Set_NonEmpty_go__append
}

var cache_Data_Set_NonEmpty_coerce3 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce3 sync.Once

func Get_Data_Set_NonEmpty_coerce3() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce3.Do(func() {
		cache_Data_Set_NonEmpty_coerce3 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce3
}

var cache_Data_Set_NonEmpty_fromJust gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromJust sync.Once

func Get_Data_Set_NonEmpty_fromJust() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromJust.Do(func() {
		cache_Data_Set_NonEmpty_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_fromJust
}

var cache_Data_Set_NonEmpty_coerce4 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce4 sync.Once

func Get_Data_Set_NonEmpty_coerce4() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce4.Do(func() {
		cache_Data_Set_NonEmpty_coerce4 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce4
}

var cache_Data_Set_NonEmpty_fromJust1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromJust1 sync.Once

func Get_Data_Set_NonEmpty_fromJust1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromJust1.Do(func() {
		cache_Data_Set_NonEmpty_fromJust1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_fromJust1(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_fromJust1
}

var cache_Data_Set_NonEmpty_coerce5 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce5 sync.Once

func Get_Data_Set_NonEmpty_coerce5() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce5.Do(func() {
		cache_Data_Set_NonEmpty_coerce5 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce5
}

var cache_Data_Set_NonEmpty_coerce6 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce6 sync.Once

func Get_Data_Set_NonEmpty_coerce6() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce6.Do(func() {
		cache_Data_Set_NonEmpty_coerce6 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce6
}

var cache_Data_Set_NonEmpty_coerce7 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce7 sync.Once

func Get_Data_Set_NonEmpty_coerce7() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce7.Do(func() {
		cache_Data_Set_NonEmpty_coerce7 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce7
}

var cache_Data_Set_NonEmpty_compose gopurs_runtime.Value
var once_Data_Set_NonEmpty_compose sync.Once

func Get_Data_Set_NonEmpty_compose() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_compose.Do(func() {
		cache_Data_Set_NonEmpty_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_NonEmpty_compose
}

var cache_Data_Set_NonEmpty_compose1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_compose1 sync.Once

func Get_Data_Set_NonEmpty_compose1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_compose1.Do(func() {
		cache_Data_Set_NonEmpty_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_NonEmpty_compose1
}

var cache_Data_Set_NonEmpty_foldMap1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldMap1 sync.Once

func Get_Data_Set_NonEmpty_foldMap1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldMap1.Do(func() {
		cache_Data_Set_NonEmpty_foldMap1 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1022383170_1680800814(Rebox_Data_Set_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}), "foldMap1")
	})
	return cache_Data_Set_NonEmpty_foldMap1
}

var cache_Data_Set_NonEmpty_compose2 gopurs_runtime.Value
var once_Data_Set_NonEmpty_compose2 sync.Once

func Get_Data_Set_NonEmpty_compose2() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_compose2.Do(func() {
		cache_Data_Set_NonEmpty_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_NonEmpty_compose2
}

var cache_Data_Set_NonEmpty_foldr1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldr1 sync.Once

func Get_Data_Set_NonEmpty_foldr1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldr1.Do(func() {
		cache_Data_Set_NonEmpty_foldr1 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1022383170_1680800814(Rebox_Data_Set_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}), "foldr1")
	})
	return cache_Data_Set_NonEmpty_foldr1
}

var cache_Data_Set_NonEmpty_compose3 gopurs_runtime.Value
var once_Data_Set_NonEmpty_compose3 sync.Once

func Get_Data_Set_NonEmpty_compose3() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_compose3.Do(func() {
		cache_Data_Set_NonEmpty_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_NonEmpty_compose3
}

var cache_Data_Set_NonEmpty_foldl1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldl1 sync.Once

func Get_Data_Set_NonEmpty_foldl1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldl1.Do(func() {
		cache_Data_Set_NonEmpty_foldl1 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1022383170_1680800814(Rebox_Data_Set_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}), "foldl1")
	})
	return cache_Data_Set_NonEmpty_foldl1
}

var cache_Data_Set_NonEmpty_coerce8 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce8 sync.Once

func Get_Data_Set_NonEmpty_coerce8() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce8.Do(func() {
		cache_Data_Set_NonEmpty_coerce8 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce8
}

var cache_Data_Set_NonEmpty_coerce9 gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce9 sync.Once

func Get_Data_Set_NonEmpty_coerce9() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce9.Do(func() {
		cache_Data_Set_NonEmpty_coerce9 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce9
}

var cache_Data_Set_NonEmpty_NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_NonEmptySet sync.Once

func Get_Data_Set_NonEmpty_NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_NonEmptySet = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_NonEmptySet(x_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_NonEmptySet
}

var cache_Data_Set_NonEmpty_unionSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_unionSet sync.Once

func Get_Data_Set_NonEmpty_unionSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_unionSet.Do(func() {
		cache_Data_Set_NonEmpty_unionSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_unionSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_unionSet
}

var cache_Data_Set_NonEmpty_toUnfoldable1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable1 sync.Once

func Get_Data_Set_NonEmpty_toUnfoldable1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable1.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable1 = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), v_1_box)
		})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable1
}

var cache_Data_Set_NonEmpty_toUnfoldable11 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable11 sync.Once

func Get_Data_Set_NonEmpty_toUnfoldable11() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable11.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable11 = gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()))})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable11
}

var cache_Data_Set_NonEmpty_toUnfoldable12 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable12 sync.Once

func Get_Data_Set_NonEmpty_toUnfoldable12() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable12.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable12 = gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Set_NonEmpty___cached_dict_0()))))})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable12
}

var cache_Data_Set_NonEmpty_toUnfoldable13 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable13 sync.Once

func Get_Data_Set_NonEmpty_toUnfoldable13() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable13.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable13 = gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Set_NonEmpty___cached_dict_1()))))})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable13
}

var cache_Data_Set_NonEmpty_toUnfoldable14 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable14 sync.Once

func Get_Data_Set_NonEmpty_toUnfoldable14() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable14.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable14 = gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Set_NonEmpty___cached_dict_2()))))})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable14
}

var cache_Data_Set_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable sync.Once

func Get_Data_Set_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable
}

var cache_Data_Set_NonEmpty_toSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_toSet sync.Once

func Get_Data_Set_NonEmpty_toSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toSet.Do(func() {
		cache_Data_Set_NonEmpty_toSet = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_toSet(v_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_toSet
}

var cache_Data_Set_NonEmpty_subset gopurs_runtime.Value
var once_Data_Set_NonEmpty_subset sync.Once

func Get_Data_Set_NonEmpty_subset() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_subset.Do(func() {
		cache_Data_Set_NonEmpty_subset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_subset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_subset
}

var cache_Data_Set_NonEmpty_size gopurs_runtime.Value
var once_Data_Set_NonEmpty_size sync.Once

func Get_Data_Set_NonEmpty_size() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_size.Do(func() {
		cache_Data_Set_NonEmpty_size = Get_Data_Set_size()
	})
	return cache_Data_Set_NonEmpty_size
}

var cache_Data_Set_NonEmpty_singleton gopurs_runtime.Value
var once_Data_Set_NonEmpty_singleton sync.Once

func Get_Data_Set_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_singleton.Do(func() {
		cache_Data_Set_NonEmpty_singleton = Get_Data_Set_singleton()
	})
	return cache_Data_Set_NonEmpty_singleton
}

var cache_Data_Set_NonEmpty_showNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_showNonEmptySet sync.Once

func Get_Data_Set_NonEmpty_showNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_showNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_showNonEmptySet = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_showNonEmptySet(dictShow_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_showNonEmptySet
}

var cache_Data_Set_NonEmpty_semigroupNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_semigroupNonEmptySet sync.Once

func Get_Data_Set_NonEmpty_semigroupNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_semigroupNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_semigroupNonEmptySet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_semigroupNonEmptySet(dictOrd_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_semigroupNonEmptySet
}

var cache_Data_Set_NonEmpty_properSubset gopurs_runtime.Value
var once_Data_Set_NonEmpty_properSubset sync.Once

func Get_Data_Set_NonEmpty_properSubset() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_properSubset.Do(func() {
		cache_Data_Set_NonEmpty_properSubset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_properSubset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_properSubset
}

var cache_Data_Set_NonEmpty_ordNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_ordNonEmptySet sync.Once

func Get_Data_Set_NonEmpty_ordNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_ordNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_ordNonEmptySet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_ordNonEmptySet(dictOrd_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_ordNonEmptySet
}

var cache_Data_Set_NonEmpty_ord1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_ord1NonEmptySet sync.Once

func Get_Data_Set_NonEmpty_ord1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_ord1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_ord1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_Set_ord1Set()))}
	})
	return cache_Data_Set_NonEmpty_ord1NonEmptySet
}

var cache_Data_Set_NonEmpty_min gopurs_runtime.Value
var once_Data_Set_NonEmpty_min sync.Once

func Get_Data_Set_NonEmpty_min() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_min.Do(func() {
		cache_Data_Set_NonEmpty_min = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_min(v_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_min
}

var cache_Data_Set_NonEmpty_member gopurs_runtime.Value
var once_Data_Set_NonEmpty_member sync.Once

func Get_Data_Set_NonEmpty_member() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_member.Do(func() {
		cache_Data_Set_NonEmpty_member = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_member
}

var cache_Data_Set_NonEmpty_max gopurs_runtime.Value
var once_Data_Set_NonEmpty_max sync.Once

func Get_Data_Set_NonEmpty_max() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_max.Do(func() {
		cache_Data_Set_NonEmpty_max = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_max(v_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_max
}

var cache_Data_Set_NonEmpty_mapMaybe gopurs_runtime.Value
var once_Data_Set_NonEmpty_mapMaybe sync.Once

func Get_Data_Set_NonEmpty_mapMaybe() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_mapMaybe.Do(func() {
		cache_Data_Set_NonEmpty_mapMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_mapMaybe
}

var cache_Data_Set_NonEmpty_go__map gopurs_runtime.Value
var once_Data_Set_NonEmpty_go__map sync.Once

func Get_Data_Set_NonEmpty_go__map() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_go__map.Do(func() {
		cache_Data_Set_NonEmpty_go__map = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_go__map
}

var cache_Data_Set_NonEmpty_insert gopurs_runtime.Value
var once_Data_Set_NonEmpty_insert sync.Once

func Get_Data_Set_NonEmpty_insert() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_insert.Do(func() {
		cache_Data_Set_NonEmpty_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_insert
}

var cache_Data_Set_NonEmpty_fromSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromSet sync.Once

func Get_Data_Set_NonEmpty_fromSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromSet.Do(func() {
		cache_Data_Set_NonEmpty_fromSet = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Set_NonEmpty_fromSet(s_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Set_NonEmpty_fromSet
}

var cache_Data_Set_NonEmpty_intersection gopurs_runtime.Value
var once_Data_Set_NonEmpty_intersection sync.Once

func Get_Data_Set_NonEmpty_intersection() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_intersection.Do(func() {
		cache_Data_Set_NonEmpty_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_intersection
}

var cache_Data_Set_NonEmpty_fromFoldable1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromFoldable1 sync.Once

func Get_Data_Set_NonEmpty_fromFoldable1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromFoldable1.Do(func() {
		cache_Data_Set_NonEmpty_fromFoldable1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
		})
	})
	return cache_Data_Set_NonEmpty_fromFoldable1
}

var cache_Data_Set_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromFoldable sync.Once

func Get_Data_Set_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromFoldable.Do(func() {
		cache_Data_Set_NonEmpty_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
		})
	})
	return cache_Data_Set_NonEmpty_fromFoldable
}

var cache_Data_Set_NonEmpty_foldableNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldableNonEmptySet sync.Once

func Get_Data_Set_NonEmpty_foldableNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldableNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_foldableNonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Set_foldableSet()))}
	})
	return cache_Data_Set_NonEmpty_foldableNonEmptySet
}

var cache_Data_Set_NonEmpty_foldable1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldable1NonEmptySet sync.Once

func Get_Data_Set_NonEmpty_foldable1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldable1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_foldable1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Set_foldableSet()))}
		}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMap11_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope68)] (TypeVar m$scope69)), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a$scope68)])] (TypeVar m$scope69))
			foldMap11_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1022383170_1680800814(Rebox_Data_Set_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))})
			_ = foldMap11_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(foldMap11_1_0, f_2), gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Set_NonEmpty___cached_dict_3()))))}))
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1022383170_1680800814(Rebox_Data_Set_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}), "foldl1"), f_0), gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Set_NonEmpty___cached_dict_5()))))}))
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1022383170_1680800814(Rebox_Data_Set_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}), "foldr1"), f_0), gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Set_NonEmpty___cached_dict_4()))))}))
		})}))}
	})
	return cache_Data_Set_NonEmpty_foldable1NonEmptySet
}

var cache_Data_Set_NonEmpty_filter gopurs_runtime.Value
var once_Data_Set_NonEmpty_filter sync.Once

func Get_Data_Set_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_filter.Do(func() {
		cache_Data_Set_NonEmpty_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_filter
}

var cache_Data_Set_NonEmpty_eqNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_eqNonEmptySet sync.Once

func Get_Data_Set_NonEmpty_eqNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_eqNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_eqNonEmptySet = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_eqNonEmptySet(dictEq_0_box)
		})
	})
	return cache_Data_Set_NonEmpty_eqNonEmptySet
}

var cache_Data_Set_NonEmpty_eq1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_eq1NonEmptySet sync.Once

func Get_Data_Set_NonEmpty_eq1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_eq1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_eq1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Set_eq1Set()))}
	})
	return cache_Data_Set_NonEmpty_eq1NonEmptySet
}

var cache_Data_Set_NonEmpty_difference gopurs_runtime.Value
var once_Data_Set_NonEmpty_difference sync.Once

func Get_Data_Set_NonEmpty_difference() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_difference.Do(func() {
		cache_Data_Set_NonEmpty_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_difference
}

var cache_Data_Set_NonEmpty_go__delete gopurs_runtime.Value
var once_Data_Set_NonEmpty_go__delete sync.Once

func Get_Data_Set_NonEmpty_go__delete() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_go__delete.Do(func() {
		cache_Data_Set_NonEmpty_go__delete = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_go__delete
}

var cache_Data_Set_NonEmpty_cons gopurs_runtime.Value
var once_Data_Set_NonEmpty_cons sync.Once

func Get_Data_Set_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_cons.Do(func() {
		cache_Data_Set_NonEmpty_cons = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_NonEmpty_cons(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_NonEmpty_cons
}

var cache_Data_Set_NonEmpty___cached_dict_0 gopurs_runtime.Value
var once_Data_Set_NonEmpty___cached_dict_0 sync.Once

func Get_Data_Set_NonEmpty___cached_dict_0() gopurs_runtime.Value {
	once_Data_Set_NonEmpty___cached_dict_0.Do(func() {
		cache_Data_Set_NonEmpty___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})))))}
	})
	return cache_Data_Set_NonEmpty___cached_dict_0
}

var cache_Data_Set_NonEmpty___cached_dict_1 gopurs_runtime.Value
var once_Data_Set_NonEmpty___cached_dict_1 sync.Once

func Get_Data_Set_NonEmpty___cached_dict_1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty___cached_dict_1.Do(func() {
		cache_Data_Set_NonEmpty___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})))))}
	})
	return cache_Data_Set_NonEmpty___cached_dict_1
}

var cache_Data_Set_NonEmpty___cached_dict_2 gopurs_runtime.Value
var once_Data_Set_NonEmpty___cached_dict_2 sync.Once

func Get_Data_Set_NonEmpty___cached_dict_2() gopurs_runtime.Value {
	once_Data_Set_NonEmpty___cached_dict_2.Do(func() {
		cache_Data_Set_NonEmpty___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})))))}
	})
	return cache_Data_Set_NonEmpty___cached_dict_2
}

var cache_Data_Set_NonEmpty___cached_dict_3 gopurs_runtime.Value
var once_Data_Set_NonEmpty___cached_dict_3 sync.Once

func Get_Data_Set_NonEmpty___cached_dict_3() gopurs_runtime.Value {
	once_Data_Set_NonEmpty___cached_dict_3.Do(func() {
		cache_Data_Set_NonEmpty___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})))))}
	})
	return cache_Data_Set_NonEmpty___cached_dict_3
}

var cache_Data_Set_NonEmpty___cached_dict_4 gopurs_runtime.Value
var once_Data_Set_NonEmpty___cached_dict_4 sync.Once

func Get_Data_Set_NonEmpty___cached_dict_4() gopurs_runtime.Value {
	once_Data_Set_NonEmpty___cached_dict_4.Do(func() {
		cache_Data_Set_NonEmpty___cached_dict_4 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})))))}
	})
	return cache_Data_Set_NonEmpty___cached_dict_4
}

var cache_Data_Set_NonEmpty___cached_dict_5 gopurs_runtime.Value
var once_Data_Set_NonEmpty___cached_dict_5 sync.Once

func Get_Data_Set_NonEmpty___cached_dict_5() gopurs_runtime.Value {
	once_Data_Set_NonEmpty___cached_dict_5.Do(func() {
		cache_Data_Set_NonEmpty___cached_dict_5 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2435023311_2187088110(Rebox_Data_Set_NonEmpty_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})))))}
	})
	return cache_Data_Set_NonEmpty___cached_dict_5
}

func Call_Data_Set_NonEmpty_fromJust(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 gopurs_runtime.Value
	{
		if v_0 != nil {
			__t0 = (v_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Set_NonEmpty_fromJust1(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 gopurs_runtime.Value
	{
		if v_0 != nil {
			__t0 = (v_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Set_NonEmpty_NonEmptySet(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Set_NonEmpty_unionSet(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Set_semigroupSet(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})))
}

func Call_Data_Set_NonEmpty_toUnfoldable1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
	_ = dictUnfoldable1_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(lst_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 gopurs_runtime.Value
		{
			if lst_2.Type == 9 && lst_2.IntVal == 1358893437 && lst_2.UnsafePtr == nil {
				__t2 = gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.StrWithCodeUnits("toUnfoldable1: empty set", 24))
				goto end_branch_2
			} else {

			}
		}
		{
			if lst_2.Type == 9 && lst_2.IntVal == 1358893437 && lst_2.UnsafePtr != nil {
				var __t1 gopurs_runtime.Value
				{
					var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(lst_2.UnsafePtr).V1
					_ = __t_tag_0
					if __t_tag_0 == nil {
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(lst_2.UnsafePtr).V1)}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()
				}
			end_branch_1:
				__t2 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(lst_2.UnsafePtr).V0, __t1}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2892990469_138441832(func() *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]] {
				panic("Failed pattern match")
			}()))}
		}
	end_branch_2:
		return __t2
	}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_toUnfoldable(Rebox_Data_Set_NonEmpty_132451362_2738507278(Rebox_Data_Set_NonEmpty_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList())))), Get_Data_Set_toList(), v_1)))})
}

func Call_Data_Set_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return Call_Data_Set_toUnfoldable(dictUnfoldable_0)
}

func Call_Data_Set_NonEmpty_toSet(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Data_Set_NonEmpty_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_subset(dictOrd_0)
}

func Call_Data_Set_NonEmpty_showNonEmptySet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(fromFoldable1 ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Array_NonEmpty_Internal_showNonEmptyArray(dictShow_0), "show"), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(Call_Data_Set_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()), s_1).UnsafePtr)))).StrVal()), ")"))
	})}))}
}

func Call_Data_Set_NonEmpty_semigroupNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Set_semigroupSet(dictOrd_0)))}
}

func Call_Data_Set_NonEmpty_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_properSubset(dictOrd_0)
}

func Call_Data_Set_NonEmpty_ordNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Set_ordSet(dictOrd_0)))}
}

func Call_Data_Set_NonEmpty_min(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a$scope32), value: Unit] Empty))])
	__local_var_1_0 := Rebox_Data_Set_NonEmpty_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), v_0)))
	_ = __local_var_1_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_1_0 != nil {
			__t1 = (__local_var_1_0).V0.key
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Data_Set_NonEmpty_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_member(dictOrd_0)
}

func Call_Data_Set_NonEmpty_max(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a$scope36), value: Unit] Empty))])
	__local_var_1_0 := Rebox_Data_Set_NonEmpty_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), v_0)))
	_ = __local_var_1_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_1_0 != nil {
			__t1 = (__local_var_1_0).V0.key
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Data_Set_NonEmpty_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_mapMaybe(dictOrd_0)
}

func Call_Data_Set_NonEmpty_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_go__map(dictOrd_0)
}

func Call_Data_Set_NonEmpty_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_insert(dictOrd_0)
}

func Call_Data_Set_NonEmpty_fromSet(s_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var s_0 gopurs_runtime.Value = s_0_loop
	_ = s_0
	var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if (gopurs_runtime.Apply(Get_Data_Map_Internal_isEmpty(), s_0).IntVal) != (0) {
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{s_0, true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Set_NonEmpty_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): intersection1_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope53), Unit]), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope53), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope53), Unit]))
	intersection1_1_0 := Call_Data_Map_Internal_intersection(dictOrd_0)
	_ = intersection1_1_0
	return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply2(intersection1_1_0, v_2, v1_3)
		_ = __local_var_4_1
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if (gopurs_runtime.Apply(Get_Data_Map_Internal_isEmpty(), __local_var_4_1).IntVal) != (0) {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{__local_var_4_1, true}
		}
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	})
}

func Call_Data_Set_NonEmpty_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
	_ = dictFoldable1_0
	var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
	_ = dictOrd_1
	return gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Set_NonEmpty_semigroupNonEmptySet(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_1)})))}, Get_Data_Set_NonEmpty_singleton())
}

func Call_Data_Set_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
	_ = dictOrd_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Set_NonEmpty_fromSet(), Call_Data_Set_fromFoldable(dictFoldable_0, dictOrd_1))
}

func Call_Data_Set_NonEmpty_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_filter(dictOrd_0)
}

func Call_Data_Set_NonEmpty_eqNonEmptySet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Set_eqSet(dictEq_0)))}
}

func Call_Data_Set_NonEmpty_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): difference1_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope86), Unit]), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope86), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope86), Unit]))
	difference1_1_0 := Call_Data_Map_Internal_difference(dictOrd_0)
	_ = difference1_1_0
	return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply2(difference1_1_0, v_2, v1_3)
		_ = __local_var_4_1
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if (gopurs_runtime.Apply(Get_Data_Map_Internal_isEmpty(), __local_var_4_1).IntVal) != (0) {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{__local_var_4_1, true}
		}
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	})
}

func Call_Data_Set_NonEmpty_go__delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): delete1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope88), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope88), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope88), Unit]))
	delete1_1_0 := Call_Data_Map_Internal_go__delete(dictOrd_0)
	_ = delete1_1_0
	return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply2(delete1_1_0, a_2, v_3)
		_ = __local_var_4_1
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if (gopurs_runtime.Apply(Get_Data_Map_Internal_isEmpty(), __local_var_4_1).IntVal) != (0) {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{__local_var_4_1, true}
		}
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	})
}

func Call_Data_Set_NonEmpty_cons(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Set_insert(dictOrd_0)
}

func Rebox_Data_Set_NonEmpty_1022383170_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Set_NonEmpty_132451362_2738507278(in *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Set_NonEmpty_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Set_NonEmpty_1978018568_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Set_NonEmpty_2187088110_2435023311(in *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Set_NonEmpty_2435023311_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Set_NonEmpty_2738507278_132451362(in *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Set_NonEmpty_2892990469_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1978018568_3094389156(in.V1))}
	return out
}

func Rebox_Data_Set_NonEmpty_3094389156_2758605161(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	key   gopurs_runtime.Value
	value gopurs_runtime.Value
	/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		key   gopurs_runtime.Value
		value gopurs_runtime.Value
		/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
	}]{}
	out.V0 = func() struct {
		key   gopurs_runtime.Value
		value gopurs_runtime.Value
		/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			key   gopurs_runtime.Value
			value gopurs_runtime.Value
			/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
		}{}
		clone.key = gopurs_runtime.RecordGet(orig, "key")
		clone.value = gopurs_runtime.RecordGet(orig, "value")
		return clone
	}()
	return out
}
