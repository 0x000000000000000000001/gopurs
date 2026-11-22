package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_NonEmpty_sequence1 gopurs_runtime.Value
var once_Data_List_NonEmpty_sequence1 sync.Once

func Get_Data_List_NonEmpty_sequence1() gopurs_runtime.Value {
	once_Data_List_NonEmpty_sequence1.Do(func() {
		cache_Data_List_NonEmpty_sequence1 = Rebox_Data_List_NonEmpty_306175789_3842311788(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_List_Types_traversable1NonEmptyList())).V2
	})
	return cache_Data_List_NonEmpty_sequence1
}

var cache_Data_List_NonEmpty_go__append gopurs_runtime.Value
var once_Data_List_NonEmpty_go__append sync.Once

func Get_Data_List_NonEmpty_go__append() gopurs_runtime.Value {
	once_Data_List_NonEmpty_go__append.Do(func() {
		cache_Data_List_NonEmpty_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_List_NonEmpty_go__append
}

var cache_Data_List_NonEmpty_eq gopurs_runtime.Value
var once_Data_List_NonEmpty_eq sync.Once

func Get_Data_List_NonEmpty_eq() gopurs_runtime.Value {
	once_Data_List_NonEmpty_eq.Do(func() {
		cache_Data_List_NonEmpty_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_List_NonEmpty_eq
}

var cache_Data_List_NonEmpty_go__map gopurs_runtime.Value
var once_Data_List_NonEmpty_go__map sync.Once

func Get_Data_List_NonEmpty_go__map() gopurs_runtime.Value {
	once_Data_List_NonEmpty_go__map.Do(func() {
		cache_Data_List_NonEmpty_go__map = Rebox_Data_List_NonEmpty_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_NonEmpty_go__map
}

var cache_Data_List_NonEmpty_compose gopurs_runtime.Value
var once_Data_List_NonEmpty_compose sync.Once

func Get_Data_List_NonEmpty_compose() gopurs_runtime.Value {
	once_Data_List_NonEmpty_compose.Do(func() {
		cache_Data_List_NonEmpty_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_NonEmpty_compose
}

var cache_Data_List_NonEmpty_sub gopurs_runtime.Value
var once_Data_List_NonEmpty_sub sync.Once

func Get_Data_List_NonEmpty_sub() gopurs_runtime.Value {
	once_Data_List_NonEmpty_sub.Do(func() {
		cache_Data_List_NonEmpty_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_List_NonEmpty_sub
}

var cache_Data_List_NonEmpty_map1 gopurs_runtime.Value
var once_Data_List_NonEmpty_map1 sync.Once

func Get_Data_List_NonEmpty_map1() gopurs_runtime.Value {
	once_Data_List_NonEmpty_map1.Do(func() {
		cache_Data_List_NonEmpty_map1 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3262795586_2812149806(Rebox_Data_List_NonEmpty_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}), "map")
	})
	return cache_Data_List_NonEmpty_map1
}

var cache_Data_List_NonEmpty_map2 gopurs_runtime.Value
var once_Data_List_NonEmpty_map2 sync.Once

func Get_Data_List_NonEmpty_map2() gopurs_runtime.Value {
	once_Data_List_NonEmpty_map2.Do(func() {
		cache_Data_List_NonEmpty_map2 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3262795586_2812149806(Rebox_Data_List_NonEmpty_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}), "map")
	})
	return cache_Data_List_NonEmpty_map2
}

var cache_Data_List_NonEmpty_compose1 gopurs_runtime.Value
var once_Data_List_NonEmpty_compose1 sync.Once

func Get_Data_List_NonEmpty_compose1() gopurs_runtime.Value {
	once_Data_List_NonEmpty_compose1.Do(func() {
		cache_Data_List_NonEmpty_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_NonEmpty_compose1
}

var cache_Data_List_NonEmpty_map3 gopurs_runtime.Value
var once_Data_List_NonEmpty_map3 sync.Once

func Get_Data_List_NonEmpty_map3() gopurs_runtime.Value {
	once_Data_List_NonEmpty_map3.Do(func() {
		cache_Data_List_NonEmpty_map3 = Rebox_Data_List_NonEmpty_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_NonEmpty_map3
}

var cache_Data_List_NonEmpty_map4 gopurs_runtime.Value
var once_Data_List_NonEmpty_map4 sync.Once

func Get_Data_List_NonEmpty_map4() gopurs_runtime.Value {
	once_Data_List_NonEmpty_map4.Do(func() {
		cache_Data_List_NonEmpty_map4 = Rebox_Data_List_NonEmpty_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_NonEmpty_map4
}

var cache_Data_List_NonEmpty_compose2 gopurs_runtime.Value
var once_Data_List_NonEmpty_compose2 sync.Once

func Get_Data_List_NonEmpty_compose2() gopurs_runtime.Value {
	once_Data_List_NonEmpty_compose2.Do(func() {
		cache_Data_List_NonEmpty_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_NonEmpty_compose2
}

var cache_Data_List_NonEmpty_add gopurs_runtime.Value
var once_Data_List_NonEmpty_add sync.Once

func Get_Data_List_NonEmpty_add() gopurs_runtime.Value {
	once_Data_List_NonEmpty_add.Do(func() {
		cache_Data_List_NonEmpty_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_List_NonEmpty_add
}

var cache_Data_List_NonEmpty_map5 gopurs_runtime.Value
var once_Data_List_NonEmpty_map5 sync.Once

func Get_Data_List_NonEmpty_map5() gopurs_runtime.Value {
	once_Data_List_NonEmpty_map5.Do(func() {
		cache_Data_List_NonEmpty_map5 = Rebox_Data_List_NonEmpty_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_NonEmpty_map5
}

var cache_Data_List_NonEmpty_compose3 gopurs_runtime.Value
var once_Data_List_NonEmpty_compose3 sync.Once

func Get_Data_List_NonEmpty_compose3() gopurs_runtime.Value {
	once_Data_List_NonEmpty_compose3.Do(func() {
		cache_Data_List_NonEmpty_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_NonEmpty_compose3
}

var cache_Data_List_NonEmpty_compose4 gopurs_runtime.Value
var once_Data_List_NonEmpty_compose4 sync.Once

func Get_Data_List_NonEmpty_compose4() gopurs_runtime.Value {
	once_Data_List_NonEmpty_compose4.Do(func() {
		cache_Data_List_NonEmpty_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_NonEmpty_compose4
}

var cache_Data_List_NonEmpty_map6 gopurs_runtime.Value
var once_Data_List_NonEmpty_map6 sync.Once

func Get_Data_List_NonEmpty_map6() gopurs_runtime.Value {
	once_Data_List_NonEmpty_map6.Do(func() {
		cache_Data_List_NonEmpty_map6 = Get_Data_Functor_map__3627393546()
	})
	return cache_Data_List_NonEmpty_map6
}

var cache_Data_List_NonEmpty_compose5 gopurs_runtime.Value
var once_Data_List_NonEmpty_compose5 sync.Once

func Get_Data_List_NonEmpty_compose5() gopurs_runtime.Value {
	once_Data_List_NonEmpty_compose5.Do(func() {
		cache_Data_List_NonEmpty_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_NonEmpty_compose5
}

var cache_Data_List_NonEmpty_bind gopurs_runtime.Value
var once_Data_List_NonEmpty_bind sync.Once

func Get_Data_List_NonEmpty_bind() gopurs_runtime.Value {
	once_Data_List_NonEmpty_bind.Do(func() {
		cache_Data_List_NonEmpty_bind = Rebox_Data_List_NonEmpty_2748095225_3054666744(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindNonEmptyList())).V1
	})
	return cache_Data_List_NonEmpty_bind
}

var cache_Data_List_NonEmpty_identity gopurs_runtime.Value
var once_Data_List_NonEmpty_identity sync.Once

func Get_Data_List_NonEmpty_identity() gopurs_runtime.Value {
	once_Data_List_NonEmpty_identity.Do(func() {
		cache_Data_List_NonEmpty_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_identity(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box)))))}
		})
	})
	return cache_Data_List_NonEmpty_identity
}

var cache_Data_List_NonEmpty_append1 gopurs_runtime.Value
var once_Data_List_NonEmpty_append1 sync.Once

func Get_Data_List_NonEmpty_append1() gopurs_runtime.Value {
	once_Data_List_NonEmpty_append1.Do(func() {
		cache_Data_List_NonEmpty_append1 = gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append")
	})
	return cache_Data_List_NonEmpty_append1
}

var cache_Data_List_NonEmpty_zipWith gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWith sync.Once

func Get_Data_List_NonEmpty_zipWith() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWith.Do(func() {
		cache_Data_List_NonEmpty_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_zipWith(f_0_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)), Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v1_2_box)))))}
		})
	})
	return cache_Data_List_NonEmpty_zipWith
}

var cache_Data_List_NonEmpty_zipWithA gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWithA sync.Once

func Get_Data_List_NonEmpty_zipWithA() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWithA.Do(func() {
		cache_Data_List_NonEmpty_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_List_NonEmpty_zipWithA
}

var cache_Data_List_NonEmpty_zip gopurs_runtime.Value
var once_Data_List_NonEmpty_zip sync.Once

func Get_Data_List_NonEmpty_zip() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zip.Do(func() {
		cache_Data_List_NonEmpty_zip = gopurs_runtime.Apply(Get_Data_List_NonEmpty_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_List_NonEmpty_zip
}

var cache_Data_List_NonEmpty_wrappedOperation2 gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation2 sync.Once

func Get_Data_List_NonEmpty_wrappedOperation2() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation2.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation2 = gopurs_runtime.Func4(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_wrappedOperation2(name_0_box.StrVal(), f_1_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)), Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v1_3_box)))))}
		})
	})
	return cache_Data_List_NonEmpty_wrappedOperation2
}

var cache_Data_List_NonEmpty_wrappedOperation gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation sync.Once

func Get_Data_List_NonEmpty_wrappedOperation() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_wrappedOperation(name_0_box.StrVal(), f_1_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))))}
		})
	})
	return cache_Data_List_NonEmpty_wrappedOperation
}

var cache_Data_List_NonEmpty_updateAt gopurs_runtime.Value
var once_Data_List_NonEmpty_updateAt sync.Once

func Get_Data_List_NonEmpty_updateAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_updateAt.Do(func() {
		cache_Data_List_NonEmpty_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_updateAt(i_0_box.IntVal, a_1_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_updateAt
}

var cache_Data_List_NonEmpty_unzip gopurs_runtime.Value
var once_Data_List_NonEmpty_unzip sync.Once

func Get_Data_List_NonEmpty_unzip() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unzip.Do(func() {
		cache_Data_List_NonEmpty_unzip = gopurs_runtime.Func(func(ts_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_unzip(Rebox_Data_List_NonEmpty_1293498952_4231696009(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](ts_0_box)))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_List_NonEmpty_unzip
}

var cache_Data_List_NonEmpty_unsnoc gopurs_runtime.Value
var once_Data_List_NonEmpty_unsnoc sync.Once

func Get_Data_List_NonEmpty_unsnoc() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unsnoc.Do(func() {
		cache_Data_List_NonEmpty_unsnoc = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_NonEmpty_unsnoc(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
				_ = orig
				return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, orig.last)
			}()
		})
	})
	return cache_Data_List_NonEmpty_unsnoc
}

var cache_Data_List_NonEmpty_unionBy gopurs_runtime.Value
var once_Data_List_NonEmpty_unionBy sync.Once

func Get_Data_List_NonEmpty_unionBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unionBy.Do(func() {
		cache_Data_List_NonEmpty_unionBy = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.StrWithCodeUnits("unionBy", 7)), Get_Data_List_unionBy())
	})
	return cache_Data_List_NonEmpty_unionBy
}

var cache_Data_List_NonEmpty_union gopurs_runtime.Value
var once_Data_List_NonEmpty_union sync.Once

func Get_Data_List_NonEmpty_union() gopurs_runtime.Value {
	once_Data_List_NonEmpty_union.Do(func() {
		cache_Data_List_NonEmpty_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_NonEmpty_union
}

var cache_Data_List_NonEmpty_uncons gopurs_runtime.Value
var once_Data_List_NonEmpty_uncons sync.Once

func Get_Data_List_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_List_NonEmpty_uncons.Do(func() {
		cache_Data_List_NonEmpty_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_NonEmpty_uncons(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)})
			}()
		})
	})
	return cache_Data_List_NonEmpty_uncons
}

var cache_Data_List_NonEmpty_toList gopurs_runtime.Value
var once_Data_List_NonEmpty_toList sync.Once

func Get_Data_List_NonEmpty_toList() gopurs_runtime.Value {
	once_Data_List_NonEmpty_toList.Do(func() {
		cache_Data_List_NonEmpty_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_toList(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))))}
		})
	})
	return cache_Data_List_NonEmpty_toList
}

var cache_Data_List_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_toUnfoldable sync.Once

func Get_Data_List_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_List_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_List_NonEmpty_toUnfoldable
}

var cache_Data_List_NonEmpty_tail gopurs_runtime.Value
var once_Data_List_NonEmpty_tail sync.Once

func Get_Data_List_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_List_NonEmpty_tail.Do(func() {
		cache_Data_List_NonEmpty_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_tail(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))))}
		})
	})
	return cache_Data_List_NonEmpty_tail
}

var cache_Data_List_NonEmpty_sortBy gopurs_runtime.Value
var once_Data_List_NonEmpty_sortBy sync.Once

func Get_Data_List_NonEmpty_sortBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_sortBy.Do(func() {
		cache_Data_List_NonEmpty_sortBy = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("sortBy", 6)), Get_Data_List_sortBy())
	})
	return cache_Data_List_NonEmpty_sortBy
}

var cache_Data_List_NonEmpty_sort gopurs_runtime.Value
var once_Data_List_NonEmpty_sort sync.Once

func Get_Data_List_NonEmpty_sort() gopurs_runtime.Value {
	once_Data_List_NonEmpty_sort.Do(func() {
		cache_Data_List_NonEmpty_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_NonEmpty_sort
}

var cache_Data_List_NonEmpty_snoc gopurs_runtime.Value
var once_Data_List_NonEmpty_snoc sync.Once

func Get_Data_List_NonEmpty_snoc() gopurs_runtime.Value {
	once_Data_List_NonEmpty_snoc.Do(func() {
		cache_Data_List_NonEmpty_snoc = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_snoc(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)), y_1_box)))}
		})
	})
	return cache_Data_List_NonEmpty_snoc
}

var cache_Data_List_NonEmpty_singleton gopurs_runtime.Value
var once_Data_List_NonEmpty_singleton sync.Once

func Get_Data_List_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_List_NonEmpty_singleton.Do(func() {
		cache_Data_List_NonEmpty_singleton = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), Call_Data_NonEmpty_singleton(Rebox_Data_List_NonEmpty_2173899317_3706288089(Rebox_Data_List_NonEmpty_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList())))))
	})
	return cache_Data_List_NonEmpty_singleton
}

var cache_Data_List_NonEmpty_singleton__1307584590 gopurs_runtime.Value
var once_Data_List_NonEmpty_singleton__1307584590 sync.Once

func Get_Data_List_NonEmpty_singleton__1307584590() gopurs_runtime.Value {
	once_Data_List_NonEmpty_singleton__1307584590.Do(func() {
		cache_Data_List_NonEmpty_singleton__1307584590 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_singleton__1307584590(__eta_norm_0_0_box)))}
		})
	})
	return cache_Data_List_NonEmpty_singleton__1307584590
}

var cache_Data_List_NonEmpty_snoc_prime_ gopurs_runtime.Value
var once_Data_List_NonEmpty_snoc_prime_ sync.Once

func Get_Data_List_NonEmpty_snoc_prime_() gopurs_runtime.Value {
	once_Data_List_NonEmpty_snoc_prime_.Do(func() {
		cache_Data_List_NonEmpty_snoc_prime_ = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_snoc_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)))}
		})
	})
	return cache_Data_List_NonEmpty_snoc_prime_
}

var cache_Data_List_NonEmpty_reverse gopurs_runtime.Value
var once_Data_List_NonEmpty_reverse sync.Once

func Get_Data_List_NonEmpty_reverse() gopurs_runtime.Value {
	once_Data_List_NonEmpty_reverse.Do(func() {
		cache_Data_List_NonEmpty_reverse = gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("reverse", 7), Get_Data_List_reverse())
	})
	return cache_Data_List_NonEmpty_reverse
}

var cache_Data_List_NonEmpty_nubEq gopurs_runtime.Value
var once_Data_List_NonEmpty_nubEq sync.Once

func Get_Data_List_NonEmpty_nubEq() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nubEq.Do(func() {
		cache_Data_List_NonEmpty_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_NonEmpty_nubEq
}

var cache_Data_List_NonEmpty_nubByEq gopurs_runtime.Value
var once_Data_List_NonEmpty_nubByEq sync.Once

func Get_Data_List_NonEmpty_nubByEq() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nubByEq.Do(func() {
		cache_Data_List_NonEmpty_nubByEq = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("nubByEq", 7)), Get_Data_List_nubByEq())
	})
	return cache_Data_List_NonEmpty_nubByEq
}

var cache_Data_List_NonEmpty_nubBy gopurs_runtime.Value
var once_Data_List_NonEmpty_nubBy sync.Once

func Get_Data_List_NonEmpty_nubBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nubBy.Do(func() {
		cache_Data_List_NonEmpty_nubBy = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("nubBy", 5)), Get_Data_List_nubBy())
	})
	return cache_Data_List_NonEmpty_nubBy
}

var cache_Data_List_NonEmpty_nub gopurs_runtime.Value
var once_Data_List_NonEmpty_nub sync.Once

func Get_Data_List_NonEmpty_nub() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nub.Do(func() {
		cache_Data_List_NonEmpty_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_NonEmpty_nub
}

var cache_Data_List_NonEmpty_modifyAt gopurs_runtime.Value
var once_Data_List_NonEmpty_modifyAt sync.Once

func Get_Data_List_NonEmpty_modifyAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_modifyAt.Do(func() {
		cache_Data_List_NonEmpty_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_modifyAt(i_0_box.IntVal, f_1_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_modifyAt
}

var cache_Data_List_NonEmpty_lift gopurs_runtime.Value
var once_Data_List_NonEmpty_lift sync.Once

func Get_Data_List_NonEmpty_lift() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift.Do(func() {
		cache_Data_List_NonEmpty_lift = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_lift(f_0_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))
		})
	})
	return cache_Data_List_NonEmpty_lift
}

var cache_Data_List_NonEmpty_mapMaybe gopurs_runtime.Value
var once_Data_List_NonEmpty_mapMaybe sync.Once

func Get_Data_List_NonEmpty_mapMaybe() gopurs_runtime.Value {
	once_Data_List_NonEmpty_mapMaybe.Do(func() {
		cache_Data_List_NonEmpty_mapMaybe = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_mapMaybe())
	})
	return cache_Data_List_NonEmpty_mapMaybe
}

var cache_Data_List_NonEmpty_partition gopurs_runtime.Value
var once_Data_List_NonEmpty_partition sync.Once

func Get_Data_List_NonEmpty_partition() gopurs_runtime.Value {
	once_Data_List_NonEmpty_partition.Do(func() {
		cache_Data_List_NonEmpty_partition = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_partition())
	})
	return cache_Data_List_NonEmpty_partition
}

var cache_Data_List_NonEmpty_span gopurs_runtime.Value
var once_Data_List_NonEmpty_span sync.Once

func Get_Data_List_NonEmpty_span() gopurs_runtime.Value {
	once_Data_List_NonEmpty_span.Do(func() {
		cache_Data_List_NonEmpty_span = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_span())
	})
	return cache_Data_List_NonEmpty_span
}

var cache_Data_List_NonEmpty_take gopurs_runtime.Value
var once_Data_List_NonEmpty_take sync.Once

func Get_Data_List_NonEmpty_take() gopurs_runtime.Value {
	once_Data_List_NonEmpty_take.Do(func() {
		cache_Data_List_NonEmpty_take = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_take())
	})
	return cache_Data_List_NonEmpty_take
}

var cache_Data_List_NonEmpty_takeWhile gopurs_runtime.Value
var once_Data_List_NonEmpty_takeWhile sync.Once

func Get_Data_List_NonEmpty_takeWhile() gopurs_runtime.Value {
	once_Data_List_NonEmpty_takeWhile.Do(func() {
		cache_Data_List_NonEmpty_takeWhile = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_takeWhile())
	})
	return cache_Data_List_NonEmpty_takeWhile
}

var cache_Data_List_NonEmpty_length gopurs_runtime.Value
var once_Data_List_NonEmpty_length sync.Once

func Get_Data_List_NonEmpty_length() gopurs_runtime.Value {
	once_Data_List_NonEmpty_length.Do(func() {
		cache_Data_List_NonEmpty_length = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_List_NonEmpty_length(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))))
		})
	})
	return cache_Data_List_NonEmpty_length
}

var cache_Data_List_NonEmpty_last gopurs_runtime.Value
var once_Data_List_NonEmpty_last sync.Once

func Get_Data_List_NonEmpty_last() gopurs_runtime.Value {
	once_Data_List_NonEmpty_last.Do(func() {
		cache_Data_List_NonEmpty_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_last(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
		})
	})
	return cache_Data_List_NonEmpty_last
}

var cache_Data_List_NonEmpty_intersectBy gopurs_runtime.Value
var once_Data_List_NonEmpty_intersectBy sync.Once

func Get_Data_List_NonEmpty_intersectBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_intersectBy.Do(func() {
		cache_Data_List_NonEmpty_intersectBy = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.StrWithCodeUnits("intersectBy", 11)), Get_Data_List_intersectBy())
	})
	return cache_Data_List_NonEmpty_intersectBy
}

var cache_Data_List_NonEmpty_intersect gopurs_runtime.Value
var once_Data_List_NonEmpty_intersect sync.Once

func Get_Data_List_NonEmpty_intersect() gopurs_runtime.Value {
	once_Data_List_NonEmpty_intersect.Do(func() {
		cache_Data_List_NonEmpty_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_NonEmpty_intersect
}

var cache_Data_List_NonEmpty_insertAt gopurs_runtime.Value
var once_Data_List_NonEmpty_insertAt sync.Once

func Get_Data_List_NonEmpty_insertAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_insertAt.Do(func() {
		cache_Data_List_NonEmpty_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_insertAt(i_0_box.IntVal, a_1_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_insertAt
}

var cache_Data_List_NonEmpty_go__init gopurs_runtime.Value
var once_Data_List_NonEmpty_go__init sync.Once

func Get_Data_List_NonEmpty_go__init() gopurs_runtime.Value {
	once_Data_List_NonEmpty_go__init.Do(func() {
		cache_Data_List_NonEmpty_go__init = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_go__init(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))))}
		})
	})
	return cache_Data_List_NonEmpty_go__init
}

var cache_Data_List_NonEmpty_index gopurs_runtime.Value
var once_Data_List_NonEmpty_index sync.Once

func Get_Data_List_NonEmpty_index() gopurs_runtime.Value {
	once_Data_List_NonEmpty_index.Do(func() {
		cache_Data_List_NonEmpty_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_index(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)), i_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_index
}

var cache_Data_List_NonEmpty_head gopurs_runtime.Value
var once_Data_List_NonEmpty_head sync.Once

func Get_Data_List_NonEmpty_head() gopurs_runtime.Value {
	once_Data_List_NonEmpty_head.Do(func() {
		cache_Data_List_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_head(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
		})
	})
	return cache_Data_List_NonEmpty_head
}

var cache_Data_List_NonEmpty_groupBy gopurs_runtime.Value
var once_Data_List_NonEmpty_groupBy sync.Once

func Get_Data_List_NonEmpty_groupBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_groupBy.Do(func() {
		cache_Data_List_NonEmpty_groupBy = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("groupBy", 7)), Get_Data_List_groupBy())
	})
	return cache_Data_List_NonEmpty_groupBy
}

var cache_Data_List_NonEmpty_groupAllBy gopurs_runtime.Value
var once_Data_List_NonEmpty_groupAllBy sync.Once

func Get_Data_List_NonEmpty_groupAllBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_groupAllBy.Do(func() {
		cache_Data_List_NonEmpty_groupAllBy = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("groupAllBy", 10)), Get_Data_List_groupAllBy())
	})
	return cache_Data_List_NonEmpty_groupAllBy
}

var cache_Data_List_NonEmpty_groupAll gopurs_runtime.Value
var once_Data_List_NonEmpty_groupAll sync.Once

func Get_Data_List_NonEmpty_groupAll() gopurs_runtime.Value {
	once_Data_List_NonEmpty_groupAll.Do(func() {
		cache_Data_List_NonEmpty_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_NonEmpty_groupAll
}

var cache_Data_List_NonEmpty_group gopurs_runtime.Value
var once_Data_List_NonEmpty_group sync.Once

func Get_Data_List_NonEmpty_group() gopurs_runtime.Value {
	once_Data_List_NonEmpty_group.Do(func() {
		cache_Data_List_NonEmpty_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_NonEmpty_group
}

var cache_Data_List_NonEmpty_fromList gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList sync.Once

func Get_Data_List_NonEmpty_fromList() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList.Do(func() {
		cache_Data_List_NonEmpty_fromList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_fromList(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_fromList
}

var cache_Data_List_NonEmpty_fromList__4030873370 gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList__4030873370 sync.Once

func Get_Data_List_NonEmpty_fromList__4030873370() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList__4030873370.Do(func() {
		cache_Data_List_NonEmpty_fromList__4030873370 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_fromList__4030873370(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_fromList__4030873370
}

var cache_Data_List_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_fromFoldable sync.Once

func Get_Data_List_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromFoldable.Do(func() {
		cache_Data_List_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_List_NonEmpty_fromFoldable
}

var cache_Data_List_NonEmpty_foldM gopurs_runtime.Value
var once_Data_List_NonEmpty_foldM sync.Once

func Get_Data_List_NonEmpty_foldM() gopurs_runtime.Value {
	once_Data_List_NonEmpty_foldM.Do(func() {
		cache_Data_List_NonEmpty_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), f_1_box, b_2_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_box)))
		})
	})
	return cache_Data_List_NonEmpty_foldM
}

var cache_Data_List_NonEmpty_findLastIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_findLastIndex sync.Once

func Get_Data_List_NonEmpty_findLastIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findLastIndex.Do(func() {
		cache_Data_List_NonEmpty_findLastIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_findLastIndex(f_0_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_findLastIndex
}

var cache_Data_List_NonEmpty_findIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_findIndex sync.Once

func Get_Data_List_NonEmpty_findIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findIndex.Do(func() {
		cache_Data_List_NonEmpty_findIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_findIndex(f_0_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_NonEmpty_findIndex
}

var cache_Data_List_NonEmpty_filterM gopurs_runtime.Value
var once_Data_List_NonEmpty_filterM sync.Once

func Get_Data_List_NonEmpty_filterM() gopurs_runtime.Value {
	once_Data_List_NonEmpty_filterM.Do(func() {
		cache_Data_List_NonEmpty_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Data_List_NonEmpty_filterM
}

var cache_Data_List_NonEmpty_filter gopurs_runtime.Value
var once_Data_List_NonEmpty_filter sync.Once

func Get_Data_List_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_List_NonEmpty_filter.Do(func() {
		cache_Data_List_NonEmpty_filter = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_filter())
	})
	return cache_Data_List_NonEmpty_filter
}

var cache_Data_List_NonEmpty_elemLastIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_elemLastIndex sync.Once

func Get_Data_List_NonEmpty_elemLastIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_elemLastIndex.Do(func() {
		cache_Data_List_NonEmpty_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_List_NonEmpty_elemLastIndex
}

var cache_Data_List_NonEmpty_elemIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_elemIndex sync.Once

func Get_Data_List_NonEmpty_elemIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_elemIndex.Do(func() {
		cache_Data_List_NonEmpty_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_List_NonEmpty_elemIndex
}

var cache_Data_List_NonEmpty_dropWhile gopurs_runtime.Value
var once_Data_List_NonEmpty_dropWhile sync.Once

func Get_Data_List_NonEmpty_dropWhile() gopurs_runtime.Value {
	once_Data_List_NonEmpty_dropWhile.Do(func() {
		cache_Data_List_NonEmpty_dropWhile = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_dropWhile())
	})
	return cache_Data_List_NonEmpty_dropWhile
}

var cache_Data_List_NonEmpty_drop gopurs_runtime.Value
var once_Data_List_NonEmpty_drop sync.Once

func Get_Data_List_NonEmpty_drop() gopurs_runtime.Value {
	once_Data_List_NonEmpty_drop.Do(func() {
		cache_Data_List_NonEmpty_drop = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Get_Data_List_drop())
	})
	return cache_Data_List_NonEmpty_drop
}

var cache_Data_List_NonEmpty_cons_prime_ gopurs_runtime.Value
var once_Data_List_NonEmpty_cons_prime_ sync.Once

func Get_Data_List_NonEmpty_cons_prime_() gopurs_runtime.Value {
	once_Data_List_NonEmpty_cons_prime_.Do(func() {
		cache_Data_List_NonEmpty_cons_prime_ = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_cons_prime_(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box))))}
		})
	})
	return cache_Data_List_NonEmpty_cons_prime_
}

var cache_Data_List_NonEmpty_cons_prime___2689777921 gopurs_runtime.Value
var once_Data_List_NonEmpty_cons_prime___2689777921 sync.Once

func Get_Data_List_NonEmpty_cons_prime___2689777921() gopurs_runtime.Value {
	once_Data_List_NonEmpty_cons_prime___2689777921.Do(func() {
		cache_Data_List_NonEmpty_cons_prime___2689777921 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_cons_prime___2689777921(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box))))}
		})
	})
	return cache_Data_List_NonEmpty_cons_prime___2689777921
}

var cache_Data_List_NonEmpty_cons gopurs_runtime.Value
var once_Data_List_NonEmpty_cons sync.Once

func Get_Data_List_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_List_NonEmpty_cons.Do(func() {
		cache_Data_List_NonEmpty_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_cons(y_0_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))))}
		})
	})
	return cache_Data_List_NonEmpty_cons
}

var cache_Data_List_NonEmpty_concatMap gopurs_runtime.Value
var once_Data_List_NonEmpty_concatMap sync.Once

func Get_Data_List_NonEmpty_concatMap() gopurs_runtime.Value {
	once_Data_List_NonEmpty_concatMap.Do(func() {
		cache_Data_List_NonEmpty_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_concatMap(b_0_box, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](a_1_box)))))}
		})
	})
	return cache_Data_List_NonEmpty_concatMap
}

var cache_Data_List_NonEmpty_concat gopurs_runtime.Value
var once_Data_List_NonEmpty_concat sync.Once

func Get_Data_List_NonEmpty_concat() gopurs_runtime.Value {
	once_Data_List_NonEmpty_concat.Do(func() {
		cache_Data_List_NonEmpty_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_concat(Rebox_Data_List_NonEmpty_1293498952_2209796965(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))))}
		})
	})
	return cache_Data_List_NonEmpty_concat
}

var cache_Data_List_NonEmpty_catMaybes gopurs_runtime.Value
var once_Data_List_NonEmpty_catMaybes sync.Once

func Get_Data_List_NonEmpty_catMaybes() gopurs_runtime.Value {
	once_Data_List_NonEmpty_catMaybes.Do(func() {
		cache_Data_List_NonEmpty_catMaybes = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_catMaybes(Rebox_Data_List_NonEmpty_1293498952_1656330533(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))))}
		})
	})
	return cache_Data_List_NonEmpty_catMaybes
}

var cache_Data_List_NonEmpty_appendFoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_appendFoldable sync.Once

func Get_Data_List_NonEmpty_appendFoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_appendFoldable.Do(func() {
		cache_Data_List_NonEmpty_appendFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_NonEmpty_appendFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_List_NonEmpty_appendFoldable
}

func Call_Data_List_NonEmpty_identity(x_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var x_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_NonEmpty_zipWith(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
	_ = v_1
	var v1_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_2_loop
	_ = v1_2
	return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (v_1).V0, (v1_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith(f_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v1_2).V1)))}})
}

func Call_Data_List_NonEmpty_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): sequence11_1_0 shape=App(Other) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeApp (TypeVar m$scope13) [(TypeVar c$scope16)])])] (TypeApp (TypeVar m$scope13) [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar c$scope16)])]))
	sequence11_1_0 := gopurs_runtime.Apply2(Rebox_Data_List_NonEmpty_306175789_3842311788(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_List_Types_traversable1NonEmptyList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{})))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
	_ = sequence11_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(sequence11_1_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_zipWith(f_2, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3)), Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4)))))})
	})
}

func Call_Data_List_NonEmpty_wrappedOperation2(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_3_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var name_0 string = name_0_loop
	_ = name_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
	_ = v_2
	var v1_3 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_3_loop
	_ = v1_3
	// TAST (Let): v2_4_0 shape=App(Other) bindingType=(ADT ["Data","List","Types","List"] [(TypeVar c$scope28)])
	v2_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v1_3).V1)}))}))
	_ = v2_4_0
	var __t1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
	{
		if v2_4_0 != nil {
			__t1 = (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v2_4_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v2_4_0).V1)}})
			goto end_branch_1
		} else {

		}
	}
	{
		if v2_4_0 == nil {
			__t1 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(gopurs_runtime.ConcatString("Impossible: empty list in NonEmptyList ", name_0)))))
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
			panic("Failed pattern match")
		}()
	}
end_branch_1:
	return __t1
}

func Call_Data_List_NonEmpty_wrappedOperation(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var name_0 string = name_0_loop
	_ = name_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
	_ = v_2
	// TAST (Let): v1_3_0 shape=App(Other) bindingType=(ADT ["Data","List","Types","List"] [(TypeVar b$scope37)])
	v1_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1)}))}))
	_ = v1_3_0
	var __t1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
	{
		if v1_3_0 != nil {
			__t1 = (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v1_3_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_3_0).V1)}})
			goto end_branch_1
		} else {

		}
	}
	{
		if v1_3_0 == nil {
			__t1 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(gopurs_runtime.ConcatString("Impossible: empty list in NonEmptyList ", name_0)))))
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
			panic("Failed pattern match")
		}()
	}
end_branch_1:
	return __t1
}

func Call_Data_List_NonEmpty_updateAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
	_ = v_2
	var __t4 gopurs_runtime.Value
	{
		if (i_0) == (int64(0)) {
			__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, (v_2).V1}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
			goto end_branch_4
		} else {

		}
	}
	{
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_2).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_0, v1_4}))}
		}))
		_ = __local_var_4_1
		// TAST (Let): __local_var_5_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope43)])])
		__local_var_5_2 := Rebox_Data_List_NonEmpty_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_updateAt(gopurs_runtime.IntSub(i_0, int64(1)), a_1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1))
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = __local_var_5_2
		var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if __local_var_5_2 != nil {
				__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_5_2).V0)}), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
		}
	end_branch_3:
		__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
	}
end_branch_4:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_NonEmpty_unzip(ts_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var ts_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = ts_0_loop
	_ = ts_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1673906728_138441832(Rebox_Data_List_NonEmpty_138441832_1673906728(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3262795586_2812149806(Rebox_Data_List_NonEmpty_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}), "map"), Get_Data_Tuple_fst(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_4231696009_1293498952(ts_0))}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3262795586_2812149806(Rebox_Data_List_NonEmpty_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}), "map"), Get_Data_Tuple_snd(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_4231696009_1293498952(ts_0))})}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Data_List_NonEmpty_unsnoc(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last     gopurs_runtime.Value
	/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
} {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	// TAST (Let): v1_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a$scope59)]), last: (TypeVar a$scope59)] Empty))])
	v1_1_0 := Rebox_Data_List_NonEmpty_3094389156_3371309921(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := Call_Data_List_unsnoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, orig.last)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))
	_ = v1_1_0
	var __t1 struct {
		go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		last     gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
	}
	{
		if v1_1_0 == nil {
			__t1 = struct {
				go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				last     gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
			}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (v_0).V0}
			goto end_branch_1
		} else {

		}
	}
	{
		if v1_1_0 != nil {
			__t1 = struct {
				go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				last     gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
			}{(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_0).V0, (v1_1_0).V0.go__init}), (v1_1_0).V0.last}
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() struct {
			go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			last     gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
		} {
			panic("Failed pattern match")
		}()
	}
end_branch_1:
	return __t1
}

func Call_Data_List_NonEmpty_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.StrWithCodeUnits("union", 5), Call_Data_List_union(dictEq_0))
}

func Call_Data_List_NonEmpty_uncons(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
} {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return struct {
		head gopurs_runtime.Value
		tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}{(v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)}
}

func Call_Data_List_NonEmpty_toList(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)})
}

func Call_Data_List_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if xs_1.Type == 9 && xs_1.IntVal == 1358893437 && xs_1.UnsafePtr == nil {
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				goto end_branch_0
			} else {

			}
		}
		{
			if xs_1.Type == 9 && xs_1.IntVal == 1358893437 && xs_1.UnsafePtr != nil {
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V1)}}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}(), true}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() struct {
				V0 gopurs_runtime.Value
				V1 bool
			} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
			}; return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: gopurs_runtime.Value{}, V1: false} }()
		}
	end_branch_0:
		return func() gopurs_runtime.Value {
			_v := __t0
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	})), Get_Data_List_NonEmpty_toList())
}

func Call_Data_List_NonEmpty_tail(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)
}

func Call_Data_List_NonEmpty_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): compare_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope85), (TypeVar a$scope85)] (ADT ["Data","Ordering","Ordering"] []))
	compare_1_0 := Call_Data_Ord_compare(dictOrd_0)
	_ = compare_1_0
	return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_sortBy(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2))))})
	})
}

func Call_Data_List_NonEmpty_snoc(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], y_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var y_1 gopurs_runtime.Value = y_1_loop
	_ = y_1
	return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Apply3(Rebox_Data_List_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, y_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))})})
}

func Call_Data_List_NonEmpty_singleton__1307584590(__eta_norm_0_0_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
singleton__1307584590:
	for {
		if false {
			continue singleton__1307584590
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), Call_Data_NonEmpty_singleton(Rebox_Data_List_NonEmpty_2173899317_3706288089(Rebox_Data_List_NonEmpty_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList())))), __eta_norm_0_0)))
	}
}

func Call_Data_List_NonEmpty_snoc_prime_(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
	{
		if v_0 != nil {
			__t0 = (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Apply3(Rebox_Data_List_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})})
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == nil {
			__t0 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_NonEmpty_singleton(), v1_1)))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
			panic("Failed pattern match")
		}()
	}
end_branch_0:
	return __t0
}

func Call_Data_List_NonEmpty_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("nubEq", 5), Call_Data_List_nubEq(dictEq_0))
}

func Call_Data_List_NonEmpty_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("nub", 3), Call_Data_List_nub(dictOrd_0))
}

func Call_Data_List_NonEmpty_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
	_ = v_2
	var __t4 gopurs_runtime.Value
	{
		if (i_0) == (int64(0)) {
			__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (v_2).V0), (v_2).V1}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
			goto end_branch_4
		} else {

		}
	}
	{
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_2).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_0, v1_4}))}
		}))
		_ = __local_var_4_1
		// TAST (Let): __local_var_5_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope113)])])
		__local_var_5_2 := Rebox_Data_List_NonEmpty_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_alterAt(gopurs_runtime.IntSub(i_0, int64(1)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), f_1), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1))
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = __local_var_5_2
		var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if __local_var_5_2 != nil {
				__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_5_2).V0)}), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
		}
	end_branch_3:
		__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
	}
end_branch_4:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_NonEmpty_lift(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))})
}

func Call_Data_List_NonEmpty_length(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) int64 {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var Call_local_Data_List_NonEmpty_go__go_1_0_0 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
	_ = Call_local_Data_List_NonEmpty_go__go_1_0_0
	var go__go_1_0_0 gopurs_runtime.Value
	_ = go__go_1_0_0
	Call_local_Data_List_NonEmpty_go__go_1_0_0 = func(b_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	go__go_1_0_0:
		for {
			if false {
				continue go__go_1_0_0
			}
			var b_2 gopurs_runtime.Value = b_2_loop
			_ = b_2
			var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
			_ = v_3
			var __t1 gopurs_runtime.Value
			{
				if v_3 == nil {
					__t1 = b_2
					goto end_branch_1
				} else {

				}
			}
			{
				if v_3 != nil {
					b_2_loop = gopurs_runtime.Int(gopurs_runtime.IntAdd(b_2.IntVal, int64(1)))
					v_3_loop = (v_3).V1
					continue go__go_1_0_0
					__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
	}
	go__go_1_0_0 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_List_NonEmpty_go__go_1_0_0(b_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))
		})
	})
	return gopurs_runtime.IntAdd(int64(1), Call_local_Data_List_NonEmpty_go__go_1_0_0(gopurs_runtime.Int(int64(0)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)).IntVal)
}

func Call_Data_List_NonEmpty_last(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t6 gopurs_runtime.Value
	{
		var __t_tag_0 gopurs_runtime.Value = (v_0).V1
		_ = __t_tag_0
		if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr != nil {
			var __t5 gopurs_runtime.Value
			{
				var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1
				_ = __t_tag_4
				if __t_tag_4 == nil {
					__t5 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V0
					goto end_branch_5
				} else {

				}
			}
			{
				var __t3 gopurs_runtime.Value
				{
					var __t_tag_1 struct {
						V0 gopurs_runtime.Value
						V1 bool
					} = Call_Data_List_last((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1)
					_ = __t_tag_1
					if !__t_tag_1.V1 {
						__t3 = (v_0).V0
						goto end_branch_3
					} else {

					}
				}
				{
					var __t_tag_2 struct {
						V0 gopurs_runtime.Value
						V1 bool
					} = Call_Data_List_last((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1)
					_ = __t_tag_2
					if __t_tag_2.V1 {
						__t3 = Call_Data_List_last((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1).V0
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				__t5 = __t3
			}
		end_branch_5:
			__t6 = __t5
			goto end_branch_6
		} else {

		}
	}
	{
		__t6 = (v_0).V0
	}
end_branch_6:
	return __t6
}

func Call_Data_List_NonEmpty_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.StrWithCodeUnits("intersect", 9), Call_Data_List_intersect(dictEq_0))
}

func Call_Data_List_NonEmpty_insertAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
	_ = v_2
	var __t4 gopurs_runtime.Value
	{
		if (i_0) == (int64(0)) {
			__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1)}))}}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
			goto end_branch_4
		} else {

		}
	}
	{
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_2).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_0, v1_4}))}
		}))
		_ = __local_var_4_1
		// TAST (Let): __local_var_5_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope154)])])
		__local_var_5_2 := Rebox_Data_List_NonEmpty_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_insertAt(gopurs_runtime.IntSub(i_0, int64(1)), a_1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1))
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = __local_var_5_2
		var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if __local_var_5_2 != nil {
				__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_5_2).V0)}), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
		}
	end_branch_3:
		__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
	}
end_branch_4:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_NonEmpty_go__init(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a$scope165)]), last: (TypeVar a$scope165)] Empty))])
	__local_var_1_0 := Rebox_Data_List_NonEmpty_3094389156_3371309921(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := Call_Data_List_unsnoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, orig.last)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))
	_ = __local_var_1_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_1_0 != nil {
			__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_0).V0, (__local_var_1_0).V0.go__init}))}
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
	}
end_branch_1:
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t1)
}

func Call_Data_List_NonEmpty_index(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], i_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var i_1 int64 = i_1_loop
	_ = i_1
	var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if (i_1) == (int64(0)) {
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(v_0).V0, true}
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
			_v := Call_Data_List_index(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1), gopurs_runtime.IntSub(i_1, int64(1)))
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

func Call_Data_List_NonEmpty_head(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (v_0).V0
}

func Call_Data_List_NonEmpty_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("groupAll", 8), Call_Data_List_groupAll(dictOrd_0))
}

func Call_Data_List_NonEmpty_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.StrWithCodeUnits("group", 5), Call_Data_List_group(dictEq_0))
}

func Call_Data_List_NonEmpty_fromList(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
	{
		if v_0 == nil {
			__t0 = Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 != nil {
			__t0 = Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
			panic("Failed pattern match")
		}()
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_NonEmpty_fromList__4030873370(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
fromList__4030873370:
	for {
		if false {
			continue fromList__4030873370
		}
		var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
		_ = v_0
		var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
		{
			if v_0 == nil {
				__t0 = Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0 != nil {
				__t0 = Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}}))}, true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
				panic("Failed pattern match")
			}()
		}
	end_branch_0:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_List_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_fromList(), Call_Data_List_fromFoldable(dictFoldable_0))
}

func Call_Data_List_NonEmpty_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	var v_3 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_3_loop
	_ = v_3
	// TAST (Let): __local_var_4_0 shape=Other bindingType=Any
	__local_var_4_0 := (v_3).V1
	_ = __local_var_4_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, (v_3).V0), gopurs_runtime.Func(func(b_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_List_foldM(dictMonad_0, f_1, b_prime__5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_0))
	}))
}

func Call_Data_List_NonEmpty_findLastIndex(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
	_ = v_1
	// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	v1_2_0 := Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := Call_Data_List_findLastIndex(f_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1))
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))
	_ = v1_2_0
	var __t2 *Constructor_Data_Maybe_Just[int64]
	{
		if v1_2_0 != nil {
			__t2 = Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(gopurs_runtime.IntAdd((v1_2_0).V0, int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_2
		} else {

		}
	}
	{
		if v1_2_0 == nil {
			var __t1 *Constructor_Data_Maybe_Just[int64]
			{
				if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
					__t1 = Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Int(int64(0)), true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))
			}
		end_branch_1:
			__t2 = __t1
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()
	}
end_branch_2:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(__t2))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_NonEmpty_findIndex(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
	_ = v_1
	var __t8 gopurs_runtime.Value
	{
		if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
			__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(int64(0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
			goto end_branch_8
		} else {

		}
	}
	{
		var Call_local_Data_List_NonEmpty_go__816245670_2_0_1 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
		_ = Call_local_Data_List_NonEmpty_go__816245670_2_0_1
		var go__816245670_2_0_1 gopurs_runtime.Value
		_ = go__816245670_2_0_1
		var Call_local_Data_List_NonEmpty_go__go_2_1_2 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
		_ = Call_local_Data_List_NonEmpty_go__go_2_1_2
		var go__go_2_1_2 gopurs_runtime.Value
		_ = go__go_2_1_2
		Call_local_Data_List_NonEmpty_go__816245670_2_0_1 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
		go__816245670_2_0_1:
			for {
				if false {
					continue go__816245670_2_0_1
				}
				var v_3 int64 = v_3_loop
				_ = v_3
				var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
				_ = v1_4
				var __t3 *Constructor_Data_Maybe_Just[int64]
				{
					if v1_4 != nil {
						var __t2 *Constructor_Data_Maybe_Just[int64]
						{
							if (gopurs_runtime.Apply(f_0, (v1_4).V0).IntVal) != (0) {
								__t2 = Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Int(v_3), true}
									if _v.V1 {
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
								}()))
								goto end_branch_2
							} else {

							}
						}
						{
							v_3_loop = gopurs_runtime.IntAdd(v_3, int64(1))
							v1_4_loop = (v1_4).V1
							continue go__816245670_2_0_1
							__t2 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
						}
					end_branch_2:
						__t3 = __t2
						goto end_branch_3
					} else {

					}
				}
				{
					if v1_4 == nil {
						__t3 = Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			}
		}
		go__816245670_2_0_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Call_local_Data_List_NonEmpty_go__816245670_2_0_1(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
			})
		})
		Call_local_Data_List_NonEmpty_go__go_2_1_2 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
		go__go_2_1_2:
			for {
				if false {
					continue go__go_2_1_2
				}
				var v_3 int64 = v_3_loop
				_ = v_3
				var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
				_ = v1_4
				var __t5 *Constructor_Data_Maybe_Just[int64]
				{
					if v1_4 != nil {
						var __t4 *Constructor_Data_Maybe_Just[int64]
						{
							if (gopurs_runtime.Apply(f_0, (v1_4).V0).IntVal) != (0) {
								__t4 = Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Int(v_3), true}
									if _v.V1 {
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
								}()))
								goto end_branch_4
							} else {

							}
						}
						{
							__t4 = Call_local_Data_List_NonEmpty_go__816245670_2_0_1(gopurs_runtime.IntAdd(v_3, int64(1)), (v1_4).V1)
						}
					end_branch_4:
						__t5 = __t4
						goto end_branch_5
					} else {

					}
				}
				{
					if v1_4 == nil {
						__t5 = Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()
				}
			end_branch_5:
				return __t5
			}
		}
		go__go_2_1_2 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Call_local_Data_List_NonEmpty_go__go_2_1_2(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
			})
		})
		// TAST (Let): __local_var_3_6 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
		__local_var_3_6 := Call_local_Data_List_NonEmpty_go__816245670_2_0_1(int64(0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1))
		_ = __local_var_3_6
		var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if __local_var_3_6 != nil {
				__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(gopurs_runtime.IntAdd((__local_var_3_6).V0, int64(1))), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				goto end_branch_7
			} else {

			}
		}
		{
			__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
		}
	end_branch_7:
		__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
	}
end_branch_8:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t8))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_NonEmpty_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_NonEmpty_lift(), Call_Data_List_filterM(dictMonad_0))
}

func Call_Data_List_NonEmpty_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Data_List_NonEmpty_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1)
	}))
}

func Call_Data_List_NonEmpty_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Data_List_NonEmpty_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1)
	}))
}

func Call_Data_List_NonEmpty_cons_prime_(x_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}})
}

func Call_Data_List_NonEmpty_cons_prime___2689777921(x_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
cons_prime___2689777921:
	for {
		if false {
			continue cons_prime___2689777921
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
		_ = xs_1
		return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}})
	}
}

func Call_Data_List_NonEmpty_cons(y_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var y_0 gopurs_runtime.Value = y_0_loop
	_ = y_0
	var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
	_ = v_1
	return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, y_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))}})
}

func Call_Data_List_NonEmpty_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = a_1_loop
	_ = a_1
	return Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_NonEmpty_2748095225_3054666744(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindNonEmptyList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(a_1))}, b_0)))
}

func Call_Data_List_NonEmpty_concat(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] = v_0_loop
	_ = v_0
	return Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_NonEmpty_2748095225_3054666744(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindNonEmptyList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_2209796965_1293498952(v_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))))
}

func Call_Data_List_NonEmpty_catMaybes(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_catMaybes(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1220287592_849153993((&Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, (v_0).V0, Rebox_Data_List_NonEmpty_849153993_1220287592(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))})))}))
}

func Call_Data_List_NonEmpty_appendFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): fromFoldable1__1127528225_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t$scope249) [(TypeVar a$scope250)])] (ADT ["Data","List","Types","List"] [(TypeVar a$scope250)]))
	fromFoldable1__1127528225_1_0 := gopurs_runtime.Apply2(dictFoldable_0.V2, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	_ = fromFoldable1__1127528225_1_0
	return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply3(Rebox_Data_List_NonEmpty_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(fromFoldable1__1127528225_1_0, ys_3)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))})}))}
	})
}

func Rebox_Data_List_NonEmpty_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_List_NonEmpty_1220287592_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = Rebox_Data_List_NonEmpty_1220287592_849153993(in.V1)
	return out
}

func Rebox_Data_List_NonEmpty_1293498952_1656330533(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Data_List_NonEmpty_1293498952_2209796965(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
	out.V0 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = in.V1
	return out
}

func Rebox_Data_List_NonEmpty_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_1293498952_4231696009(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Data_List_NonEmpty_138441832_1673906728(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
	out.V0 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](in.V1))
	return out
}

func Rebox_Data_List_NonEmpty_1673906728_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(in.V0))}
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(in.V1))}
	return out
}

func Rebox_Data_List_NonEmpty_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_2173899317_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_2209796965_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(in.V0))}
	out.V1 = in.V1
	return out
}

func Rebox_Data_List_NonEmpty_2748095225_3054666744(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_2812149806_3262795586(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_306175789_3842311788(in *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_List_NonEmpty_3094389156_1978018568(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_NonEmpty_3094389156_3371309921(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last     gopurs_runtime.Value
	/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		last     gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
	}]{}
	out.V0 = func() struct {
		go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		last     gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			last     gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
		}{}
		clone.go__init = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "init"))
		clone.last = gopurs_runtime.RecordGet(orig, "last")
		return clone
	}()
	return out
}

func Rebox_Data_List_NonEmpty_3094389156_470062885(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
	out.V0 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_List_NonEmpty_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_3262795586_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_3706288089_2173899317(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_NonEmpty_4231696009_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = in.V1
	return out
}

func Rebox_Data_List_NonEmpty_470062885_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(in.V0))}
	return out
}

func Rebox_Data_List_NonEmpty_849153993_1220287592(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0)
	out.V1 = Rebox_Data_List_NonEmpty_849153993_1220287592(in.V1)
	return out
}
