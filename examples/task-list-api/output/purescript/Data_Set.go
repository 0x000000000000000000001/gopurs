package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Set_coerce gopurs_runtime.Value
var once_Data_Set_coerce sync.Once

func Get_Data_Set_coerce() gopurs_runtime.Value {
	once_Data_Set_coerce.Do(func() {
		cache_Data_Set_coerce = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_coerce
}

var cache_Data_Set_compose gopurs_runtime.Value
var once_Data_Set_compose sync.Once

func Get_Data_Set_compose() gopurs_runtime.Value {
	once_Data_Set_compose.Do(func() {
		cache_Data_Set_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_compose
}

var cache_Data_Set_go__append gopurs_runtime.Value
var once_Data_Set_go__append sync.Once

func Get_Data_Set_go__append() gopurs_runtime.Value {
	once_Data_Set_go__append.Do(func() {
		cache_Data_Set_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Set_go__append
}

var cache_Data_Set_coerce1 gopurs_runtime.Value
var once_Data_Set_coerce1 sync.Once

func Get_Data_Set_coerce1() gopurs_runtime.Value {
	once_Data_Set_coerce1.Do(func() {
		cache_Data_Set_coerce1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_coerce1
}

var cache_Data_Set_coerce2 gopurs_runtime.Value
var once_Data_Set_coerce2 sync.Once

func Get_Data_Set_coerce2() gopurs_runtime.Value {
	once_Data_Set_coerce2.Do(func() {
		cache_Data_Set_coerce2 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_coerce2
}

var cache_Data_Set_compose1 gopurs_runtime.Value
var once_Data_Set_compose1 sync.Once

func Get_Data_Set_compose1() gopurs_runtime.Value {
	once_Data_Set_compose1.Do(func() {
		cache_Data_Set_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_compose1
}

var cache_Data_Set_foldMap gopurs_runtime.Value
var once_Data_Set_foldMap sync.Once

func Get_Data_Set_foldMap() gopurs_runtime.Value {
	once_Data_Set_foldMap.Do(func() {
		cache_Data_Set_foldMap = Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V0
	})
	return cache_Data_Set_foldMap
}

var cache_Data_Set_compose2 gopurs_runtime.Value
var once_Data_Set_compose2 sync.Once

func Get_Data_Set_compose2() gopurs_runtime.Value {
	once_Data_Set_compose2.Do(func() {
		cache_Data_Set_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_compose2
}

var cache_Data_Set_foldl gopurs_runtime.Value
var once_Data_Set_foldl sync.Once

func Get_Data_Set_foldl() gopurs_runtime.Value {
	once_Data_Set_foldl.Do(func() {
		cache_Data_Set_foldl = Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V1
	})
	return cache_Data_Set_foldl
}

var cache_Data_Set_compose3 gopurs_runtime.Value
var once_Data_Set_compose3 sync.Once

func Get_Data_Set_compose3() gopurs_runtime.Value {
	once_Data_Set_compose3.Do(func() {
		cache_Data_Set_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Set_compose3
}

var cache_Data_Set_foldr gopurs_runtime.Value
var once_Data_Set_foldr sync.Once

func Get_Data_Set_foldr() gopurs_runtime.Value {
	once_Data_Set_foldr.Do(func() {
		cache_Data_Set_foldr = Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2
	})
	return cache_Data_Set_foldr
}

var cache_Data_Set_map1 gopurs_runtime.Value
var once_Data_Set_map1 sync.Once

func Get_Data_Set_map1() gopurs_runtime.Value {
	once_Data_Set_map1.Do(func() {
		cache_Data_Set_map1 = Rebox_Data_Set_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Set_map1
}

var cache_Data_Set_map2 gopurs_runtime.Value
var once_Data_Set_map2 sync.Once

func Get_Data_Set_map2() gopurs_runtime.Value {
	once_Data_Set_map2.Do(func() {
		cache_Data_Set_map2 = Rebox_Data_Set_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Set_map2
}

var cache_Data_Set_coerce3 gopurs_runtime.Value
var once_Data_Set_coerce3 sync.Once

func Get_Data_Set_coerce3() gopurs_runtime.Value {
	once_Data_Set_coerce3.Do(func() {
		cache_Data_Set_coerce3 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_coerce3
}

var cache_Data_Set_coerce4 gopurs_runtime.Value
var once_Data_Set_coerce4 sync.Once

func Get_Data_Set_coerce4() gopurs_runtime.Value {
	once_Data_Set_coerce4.Do(func() {
		cache_Data_Set_coerce4 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_coerce4
}

var cache_Data_Set_conj gopurs_runtime.Value
var once_Data_Set_conj sync.Once

func Get_Data_Set_conj() gopurs_runtime.Value {
	once_Data_Set_conj.Do(func() {
		cache_Data_Set_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Set_conj
}

var cache_Data_Set_notEq gopurs_runtime.Value
var once_Data_Set_notEq sync.Once

func Get_Data_Set_notEq() gopurs_runtime.Value {
	once_Data_Set_notEq.Do(func() {
		cache_Data_Set_notEq = Get_Data_Eq_notEq__3020190905()
	})
	return cache_Data_Set_notEq
}

var cache_Data_Set_coerce5 gopurs_runtime.Value
var once_Data_Set_coerce5 sync.Once

func Get_Data_Set_coerce5() gopurs_runtime.Value {
	once_Data_Set_coerce5.Do(func() {
		cache_Data_Set_coerce5 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_coerce5
}

var cache_Data_Set_coerce6 gopurs_runtime.Value
var once_Data_Set_coerce6 sync.Once

func Get_Data_Set_coerce6() gopurs_runtime.Value {
	once_Data_Set_coerce6.Do(func() {
		cache_Data_Set_coerce6 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_coerce6
}

var cache_Data_Set_identity gopurs_runtime.Value
var once_Data_Set_identity sync.Once

func Get_Data_Set_identity() gopurs_runtime.Value {
	once_Data_Set_identity.Do(func() {
		cache_Data_Set_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Set_identity(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Set_identity
}

var cache_Data_Set_Set gopurs_runtime.Value
var once_Data_Set_Set sync.Once

func Get_Data_Set_Set() gopurs_runtime.Value {
	once_Data_Set_Set.Do(func() {
		cache_Data_Set_Set = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_Set(x_0_box)
		})
	})
	return cache_Data_Set_Set
}

var cache_Data_Set_union gopurs_runtime.Value
var once_Data_Set_union sync.Once

func Get_Data_Set_union() gopurs_runtime.Value {
	once_Data_Set_union.Do(func() {
		cache_Data_Set_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_union
}

var cache_Data_Set_toggle gopurs_runtime.Value
var once_Data_Set_toggle sync.Once

func Get_Data_Set_toggle() gopurs_runtime.Value {
	once_Data_Set_toggle.Do(func() {
		cache_Data_Set_toggle = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_toggle(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_toggle
}

var cache_Data_Set_toMap gopurs_runtime.Value
var once_Data_Set_toMap sync.Once

func Get_Data_Set_toMap() gopurs_runtime.Value {
	once_Data_Set_toMap.Do(func() {
		cache_Data_Set_toMap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_toMap(v_0_box)
		})
	})
	return cache_Data_Set_toMap
}

var cache_Data_Set_toList gopurs_runtime.Value
var once_Data_Set_toList sync.Once

func Get_Data_Set_toList() gopurs_runtime.Value {
	once_Data_Set_toList.Do(func() {
		cache_Data_Set_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Set_toList(v_0_box))}
		})
	})
	return cache_Data_Set_toList
}

var cache_Data_Set_toUnfoldable gopurs_runtime.Value
var once_Data_Set_toUnfoldable sync.Once

func Get_Data_Set_toUnfoldable() gopurs_runtime.Value {
	once_Data_Set_toUnfoldable.Do(func() {
		cache_Data_Set_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_Set_toUnfoldable
}

var cache_Data_Set_toUnfoldable1 gopurs_runtime.Value
var once_Data_Set_toUnfoldable1 sync.Once

func Get_Data_Set_toUnfoldable1() gopurs_runtime.Value {
	once_Data_Set_toUnfoldable1.Do(func() {
		cache_Data_Set_toUnfoldable1 = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray())), Get_Data_Set_toList())
	})
	return cache_Data_Set_toUnfoldable1
}

var cache_Data_Set_size gopurs_runtime.Value
var once_Data_Set_size sync.Once

func Get_Data_Set_size() gopurs_runtime.Value {
	once_Data_Set_size.Do(func() {
		cache_Data_Set_size = Get_Data_Map_Internal_size()
	})
	return cache_Data_Set_size
}

var cache_Data_Set_singleton gopurs_runtime.Value
var once_Data_Set_singleton sync.Once

func Get_Data_Set_singleton() gopurs_runtime.Value {
	once_Data_Set_singleton.Do(func() {
		cache_Data_Set_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_singleton(a_0_box)
		})
	})
	return cache_Data_Set_singleton
}

var cache_Data_Set_showSet gopurs_runtime.Value
var once_Data_Set_showSet sync.Once

func Get_Data_Set_showSet() gopurs_runtime.Value {
	once_Data_Set_showSet.Do(func() {
		cache_Data_Set_showSet = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_showSet(dictShow_0_box)
		})
	})
	return cache_Data_Set_showSet
}

var cache_Data_Set_semigroupSet gopurs_runtime.Value
var once_Data_Set_semigroupSet sync.Once

func Get_Data_Set_semigroupSet() gopurs_runtime.Value {
	once_Data_Set_semigroupSet.Do(func() {
		cache_Data_Set_semigroupSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_semigroupSet(dictOrd_0_box)
		})
	})
	return cache_Data_Set_semigroupSet
}

var cache_Data_Set_go__range gopurs_runtime.Value
var once_Data_Set_go__range sync.Once

func Get_Data_Set_go__range() gopurs_runtime.Value {
	once_Data_Set_go__range.Do(func() {
		cache_Data_Set_go__range = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_go__range(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_go__range
}

var cache_Data_Set_member gopurs_runtime.Value
var once_Data_Set_member sync.Once

func Get_Data_Set_member() gopurs_runtime.Value {
	once_Data_Set_member.Do(func() {
		cache_Data_Set_member = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_member
}

var cache_Data_Set_isEmpty gopurs_runtime.Value
var once_Data_Set_isEmpty sync.Once

func Get_Data_Set_isEmpty() gopurs_runtime.Value {
	once_Data_Set_isEmpty.Do(func() {
		cache_Data_Set_isEmpty = Get_Data_Map_Internal_isEmpty()
	})
	return cache_Data_Set_isEmpty
}

var cache_Data_Set_intersection gopurs_runtime.Value
var once_Data_Set_intersection sync.Once

func Get_Data_Set_intersection() gopurs_runtime.Value {
	once_Data_Set_intersection.Do(func() {
		cache_Data_Set_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_intersection
}

var cache_Data_Set_insert gopurs_runtime.Value
var once_Data_Set_insert sync.Once

func Get_Data_Set_insert() gopurs_runtime.Value {
	once_Data_Set_insert.Do(func() {
		cache_Data_Set_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_insert
}

var cache_Data_Set_fromMap gopurs_runtime.Value
var once_Data_Set_fromMap sync.Once

func Get_Data_Set_fromMap() gopurs_runtime.Value {
	once_Data_Set_fromMap.Do(func() {
		cache_Data_Set_fromMap = Get_Data_Set_Set()
	})
	return cache_Data_Set_fromMap
}

var cache_Data_Set_foldableSet gopurs_runtime.Value
var once_Data_Set_foldableSet sync.Once

func Get_Data_Set_foldableSet() gopurs_runtime.Value {
	once_Data_Set_foldableSet.Do(func() {
		cache_Data_Set_foldableSet = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMap1_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope59)] (TypeVar m$scope60)), (ADT ["Data","List","Types","List"] [(TypeVar a$scope59)])] (TypeVar m$scope60))
			foldMap1_1_0 := gopurs_runtime.Apply(Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
			_ = foldMap1_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(foldMap1_1_0, f_2), Get_Data_Set_toList())
			})
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Data_Set_go__go_2_1_0 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
			_ = Call_local_Data_Set_go__go_2_1_0
			var go__go_2_1_0 gopurs_runtime.Value
			_ = go__go_2_1_0
			Call_local_Data_Set_go__go_2_1_0 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
			go__go_2_1_0:
				for {
					if false {
						continue go__go_2_1_0
					}
					var b_3 gopurs_runtime.Value = b_3_loop
					_ = b_3
					var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
					_ = v_4
					var __t2 gopurs_runtime.Value
					{
						if v_4 == nil {
							__t2 = b_3
							goto end_branch_2
						} else {

						}
					}
					{
						if v_4 != nil {
							b_3_loop = gopurs_runtime.Apply2(f_0, b_3, (v_4).V0)
							v_4_loop = (v_4).V1
							continue go__go_2_1_0
							__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_2:
					return __t2
				}
			}
			go__go_2_1_0 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Data_Set_go__go_2_1_0(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
				})
			})
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_1_0, x_1), Get_Data_Set_toList())
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, f_0, x_1), Get_Data_Set_toList())
		})}))}
	})
	return cache_Data_Set_foldableSet
}

var cache_Data_Set_foldl1 gopurs_runtime.Value
var once_Data_Set_foldl1 sync.Once

func Get_Data_Set_foldl1() gopurs_runtime.Value {
	once_Data_Set_foldl1.Do(func() {
		cache_Data_Set_foldl1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Set_foldableSet()).V1
	})
	return cache_Data_Set_foldl1
}

var cache_Data_Set_foldr1 gopurs_runtime.Value
var once_Data_Set_foldr1 sync.Once

func Get_Data_Set_foldr1() gopurs_runtime.Value {
	once_Data_Set_foldr1.Do(func() {
		cache_Data_Set_foldr1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Set_foldableSet()).V2
	})
	return cache_Data_Set_foldr1
}

var cache_Data_Set_findMin gopurs_runtime.Value
var once_Data_Set_findMin sync.Once

func Get_Data_Set_findMin() gopurs_runtime.Value {
	once_Data_Set_findMin.Do(func() {
		cache_Data_Set_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Set_findMin(v_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Set_findMin
}

var cache_Data_Set_findMax gopurs_runtime.Value
var once_Data_Set_findMax sync.Once

func Get_Data_Set_findMax() gopurs_runtime.Value {
	once_Data_Set_findMax.Do(func() {
		cache_Data_Set_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Set_findMax(v_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Set_findMax
}

var cache_Data_Set_filter gopurs_runtime.Value
var once_Data_Set_filter sync.Once

func Get_Data_Set_filter() gopurs_runtime.Value {
	once_Data_Set_filter.Do(func() {
		cache_Data_Set_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_filter
}

var cache_Data_Set_eqSet gopurs_runtime.Value
var once_Data_Set_eqSet sync.Once

func Get_Data_Set_eqSet() gopurs_runtime.Value {
	once_Data_Set_eqSet.Do(func() {
		cache_Data_Set_eqSet = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_eqSet(dictEq_0_box)
		})
	})
	return cache_Data_Set_eqSet
}

var cache_Data_Set_ordSet gopurs_runtime.Value
var once_Data_Set_ordSet sync.Once

func Get_Data_Set_ordSet() gopurs_runtime.Value {
	once_Data_Set_ordSet.Do(func() {
		cache_Data_Set_ordSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_ordSet(dictOrd_0_box)
		})
	})
	return cache_Data_Set_ordSet
}

var cache_Data_Set_eq1Set gopurs_runtime.Value
var once_Data_Set_eq1Set sync.Once

func Get_Data_Set_eq1Set() gopurs_runtime.Value {
	once_Data_Set_eq1Set.Do(func() {
		cache_Data_Set_eq1Set = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Set_eqSet(dictEq_0)))
		})}))}
	})
	return cache_Data_Set_eq1Set
}

var cache_Data_Set_ord1Set gopurs_runtime.Value
var once_Data_Set_ord1Set sync.Once

func Get_Data_Set_ord1Set() gopurs_runtime.Value {
	once_Data_Set_ord1Set.Do(func() {
		cache_Data_Set_ord1Set = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Set_eq1Set()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Set_ordSet(dictOrd_0)))
		})}))}
	})
	return cache_Data_Set_ord1Set
}

var cache_Data_Set_empty gopurs_runtime.Value
var once_Data_Set_empty sync.Once

func Get_Data_Set_empty() gopurs_runtime.Value {
	once_Data_Set_empty.Do(func() {
		cache_Data_Set_empty = Get_Data_Map_Internal_empty()
	})
	return cache_Data_Set_empty
}

var cache_Data_Set_fromFoldable gopurs_runtime.Value
var once_Data_Set_fromFoldable sync.Once

func Get_Data_Set_fromFoldable() gopurs_runtime.Value {
	once_Data_Set_fromFoldable.Do(func() {
		cache_Data_Set_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
		})
	})
	return cache_Data_Set_fromFoldable
}

var cache_Data_Set_go__map gopurs_runtime.Value
var once_Data_Set_go__map sync.Once

func Get_Data_Set_go__map() gopurs_runtime.Value {
	once_Data_Set_go__map.Do(func() {
		cache_Data_Set_go__map = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_go__map
}

var cache_Data_Set_mapMaybe gopurs_runtime.Value
var once_Data_Set_mapMaybe sync.Once

func Get_Data_Set_mapMaybe() gopurs_runtime.Value {
	once_Data_Set_mapMaybe.Do(func() {
		cache_Data_Set_mapMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_mapMaybe
}

var cache_Data_Set_monoidSet gopurs_runtime.Value
var once_Data_Set_monoidSet sync.Once

func Get_Data_Set_monoidSet() gopurs_runtime.Value {
	once_Data_Set_monoidSet.Do(func() {
		cache_Data_Set_monoidSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_monoidSet(dictOrd_0_box)
		})
	})
	return cache_Data_Set_monoidSet
}

var cache_Data_Set_unions gopurs_runtime.Value
var once_Data_Set_unions sync.Once

func Get_Data_Set_unions() gopurs_runtime.Value {
	once_Data_Set_unions.Do(func() {
		cache_Data_Set_unions = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_unions(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
		})
	})
	return cache_Data_Set_unions
}

var cache_Data_Set_difference gopurs_runtime.Value
var once_Data_Set_difference sync.Once

func Get_Data_Set_difference() gopurs_runtime.Value {
	once_Data_Set_difference.Do(func() {
		cache_Data_Set_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_difference
}

var cache_Data_Set_subset gopurs_runtime.Value
var once_Data_Set_subset sync.Once

func Get_Data_Set_subset() gopurs_runtime.Value {
	once_Data_Set_subset.Do(func() {
		cache_Data_Set_subset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_subset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_subset
}

var cache_Data_Set_properSubset gopurs_runtime.Value
var once_Data_Set_properSubset sync.Once

func Get_Data_Set_properSubset() gopurs_runtime.Value {
	once_Data_Set_properSubset.Do(func() {
		cache_Data_Set_properSubset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_properSubset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_properSubset
}

var cache_Data_Set_go__delete gopurs_runtime.Value
var once_Data_Set_go__delete sync.Once

func Get_Data_Set_go__delete() gopurs_runtime.Value {
	once_Data_Set_go__delete.Do(func() {
		cache_Data_Set_go__delete = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_go__delete
}

var cache_Data_Set_checkValid gopurs_runtime.Value
var once_Data_Set_checkValid sync.Once

func Get_Data_Set_checkValid() gopurs_runtime.Value {
	once_Data_Set_checkValid.Do(func() {
		cache_Data_Set_checkValid = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_checkValid(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_checkValid
}

var cache_Data_Set_catMaybes gopurs_runtime.Value
var once_Data_Set_catMaybes sync.Once

func Get_Data_Set_catMaybes() gopurs_runtime.Value {
	once_Data_Set_catMaybes.Do(func() {
		cache_Data_Set_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Set_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Set_catMaybes
}

func Call_Data_Set_identity(x_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var x_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(x_0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Set_Set(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Set_union(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Map_Internal_union(dictOrd_0)
}

func Call_Data_Set_toggle(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): alter_1_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Maybe","Maybe"] [Unit])] (ADT ["Data","Maybe","Maybe"] [Unit])), (TypeVar a$scope3), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope3), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope3), Unit]))
	alter_1_0 := Call_Data_Map_Internal_alter(dictOrd_0)
	_ = alter_1_0
	return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(alter_1_0, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr == nil {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, Get_Data_Unit_unit()}))}
					goto end_branch_1
				} else {

				}
			}
			{
				if v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr != nil {
					__t1 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		}), a_2, v_3)
	})
}

func Call_Data_Set_toMap(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Data_Set_toList(v_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_keysImpl(), v_0).UnsafePtr)))))
}

func Call_Data_Set_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_toUnfoldable(dictUnfoldable_0), Get_Data_Set_toList())
}

func Call_Data_Set_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), a_0, Get_Data_Unit_unit())
}

func Call_Data_Set_showSet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(fromFoldable ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(dictShow_0), "show"), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray())), Get_Data_Set_toList(), s_1).UnsafePtr)))).StrVal()), ")"))
	})}))}
}

func Call_Data_Set_semigroupSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Call_Data_Set_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}))}
}

func Call_Data_Set_go__range(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): submap_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope30)]), (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope30)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope30), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope30), Unit]))
	submap_1_0 := Call_Data_Map_Internal_submap(dictOrd_0)
	_ = submap_1_0
	return gopurs_runtime.Func3(func(kmin_2 gopurs_runtime.Value, kmax_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(submap_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmin_2))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](kmax_3))}, v_4)
	})
}

func Call_Data_Set_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Map_Internal_member(dictOrd_0)
}

func Call_Data_Set_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Map_Internal_intersection(dictOrd_0)
}

func Call_Data_Set_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope47), Unit, (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope47), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope47), Unit]))
	insert1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insert1_1_0
	return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(insert1_1_0, a_2, Get_Data_Unit_unit(), v_3)
	})
}

func Call_Data_Set_findMin(v_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a$scope69), value: Unit] Empty))])
	__local_var_1_0 := Rebox_Data_Set_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), v_0)))
	_ = __local_var_1_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_1_0 != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(__local_var_1_0).V0.key, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Set_findMax(v_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a$scope71), value: Unit] Empty))])
	__local_var_1_0 := Rebox_Data_Set_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), v_0)))
	_ = __local_var_1_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_1_0 != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(__local_var_1_0).V0.key, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Set_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Get_Data_Map_Internal_filterKeys()
}

func Call_Data_Set_eqSet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Map_Internal_eqMap(dictEq_0, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqUnit()))}), "eq"), v_1, v1_2)
	})}))}
}

func Call_Data_Set_ordSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqSet1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope33), Unit])])
	eqSet1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Set_eqSet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
	_ = eqSet1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqSet1_1_0)}
	}), gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Types_ordList(dictOrd_0), "compare"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_keysImpl(), s1_2).UnsafePtr)))), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_keysImpl(), s2_3).UnsafePtr)))))
	})}))}
}

func Call_Data_Set_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
	_ = dictOrd_1
	// TAST (Let): insert1_2_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope88), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope88), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope88), Unit]))
	insert1_2_0 := Call_Data_Set_insert(dictOrd_1)
	_ = insert1_2_0
	return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(m_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(insert1_2_0, a_4, m_3)
	}), Get_Data_Set_empty())
}

func Call_Data_Set_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar b$scope92), (ADT ["Data","Map","Internal","Map"] [(TypeVar b$scope92), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar b$scope92), Unit]))
	insert1_1_0 := Call_Data_Set_insert(dictOrd_0)
	_ = insert1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Set_foldableSet()).V1, gopurs_runtime.Func2(func(m_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(insert1_1_0, gopurs_runtime.Apply(f_2, a_4), m_3)
		}), Get_Data_Set_empty())
	})
}

func Call_Data_Set_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar b$scope96), (ADT ["Data","Map","Internal","Map"] [(TypeVar b$scope96), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar b$scope96), Unit]))
	insert1_1_0 := Call_Data_Set_insert(dictOrd_0)
	_ = insert1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope96)])
			__local_var_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_2, a_3))
			_ = __local_var_5_1
			var __t2 gopurs_runtime.Value
			{
				if __local_var_5_1 == nil {
					__t2 = acc_4
					goto end_branch_2
				} else {

				}
			}
			{
				if __local_var_5_1 != nil {
					__t2 = gopurs_runtime.Apply2(insert1_1_0, (__local_var_5_1).V0, acc_4)
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}), Get_Data_Set_empty()), Get_Data_Set_toList())
	})
}

func Call_Data_Set_monoidSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): semigroupSet1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope98), Unit])])
	semigroupSet1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Set_semigroupSet(dictOrd_0))
	_ = semigroupSet1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupSet1_1_0)}
	}), Get_Data_Map_Internal_empty()}))}
}

func Call_Data_Set_unions(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
	_ = dictOrd_1
	return gopurs_runtime.Apply2(dictFoldable_0.V1, Call_Data_Set_union(dictOrd_1), Get_Data_Set_empty())
}

func Call_Data_Set_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Map_Internal_difference(dictOrd_0)
}

func Call_Data_Set_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): difference1_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope107), Unit]), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope107), Unit])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope107), Unit]))
	difference1_1_0 := Call_Data_Map_Internal_difference(dictOrd_0)
	_ = difference1_1_0
	return gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Data_Set_isEmpty(), gopurs_runtime.Apply2(difference1_1_0, s1_2, s2_3))
	})
}

func Call_Data_Set_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): subset1_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope109), Unit]), (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope109), Unit])] Boolean)
	subset1_1_0 := Call_Data_Set_subset(dictOrd_0)
	_ = subset1_1_0
	return gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool((((gopurs_runtime.Apply(Get_Data_Map_Internal_size(), s1_2).IntVal) == (gopurs_runtime.Apply(Get_Data_Map_Internal_size(), s2_3).IntVal)) != (true)) && ((gopurs_runtime.Apply2(subset1_1_0, s1_2, s2_3).IntVal) != (0)))
	})
}

func Call_Data_Set_go__delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Map_Internal_go__delete(dictOrd_0)
}

func Call_Data_Set_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Get_Data_Map_Internal_checkValid()
}

func Call_Data_Set_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Call_Data_Set_mapMaybe(dictOrd_0), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Rebox_Data_Set_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Set_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Set_3094389156_2758605161(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
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
