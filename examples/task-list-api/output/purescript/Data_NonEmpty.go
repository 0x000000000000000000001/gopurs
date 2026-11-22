package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_NonEmpty_go__map gopurs_runtime.Value
var once_Data_NonEmpty_go__map sync.Once

func Get_Data_NonEmpty_go__map() gopurs_runtime.Value {
	once_Data_NonEmpty_go__map.Do(func() {
		cache_Data_NonEmpty_go__map = gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map")
	})
	return cache_Data_NonEmpty_go__map
}

var cache_Data_NonEmpty_map1 gopurs_runtime.Value
var once_Data_NonEmpty_map1 sync.Once

func Get_Data_NonEmpty_map1() gopurs_runtime.Value {
	once_Data_NonEmpty_map1.Do(func() {
		cache_Data_NonEmpty_map1 = Rebox_Data_NonEmpty_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_NonEmpty_map1
}

var cache_Data_NonEmpty_go__append gopurs_runtime.Value
var once_Data_NonEmpty_go__append sync.Once

func Get_Data_NonEmpty_go__append() gopurs_runtime.Value {
	once_Data_NonEmpty_go__append.Do(func() {
		cache_Data_NonEmpty_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_NonEmpty_go__append
}

var cache_Data_NonEmpty_compose gopurs_runtime.Value
var once_Data_NonEmpty_compose sync.Once

func Get_Data_NonEmpty_compose() gopurs_runtime.Value {
	once_Data_NonEmpty_compose.Do(func() {
		cache_Data_NonEmpty_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_NonEmpty_compose
}

var cache_Data_NonEmpty_compose1 gopurs_runtime.Value
var once_Data_NonEmpty_compose1 sync.Once

func Get_Data_NonEmpty_compose1() gopurs_runtime.Value {
	once_Data_NonEmpty_compose1.Do(func() {
		cache_Data_NonEmpty_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_NonEmpty_compose1
}

var cache_Data_NonEmpty_compose2 gopurs_runtime.Value
var once_Data_NonEmpty_compose2 sync.Once

func Get_Data_NonEmpty_compose2() gopurs_runtime.Value {
	once_Data_NonEmpty_compose2.Do(func() {
		cache_Data_NonEmpty_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_NonEmpty_compose2
}

var cache_Data_NonEmpty_compose3 gopurs_runtime.Value
var once_Data_NonEmpty_compose3 sync.Once

func Get_Data_NonEmpty_compose3() gopurs_runtime.Value {
	once_Data_NonEmpty_compose3.Do(func() {
		cache_Data_NonEmpty_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_NonEmpty_compose3
}

var cache_Data_NonEmpty_compose4 gopurs_runtime.Value
var once_Data_NonEmpty_compose4 sync.Once

func Get_Data_NonEmpty_compose4() gopurs_runtime.Value {
	once_Data_NonEmpty_compose4.Do(func() {
		cache_Data_NonEmpty_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_NonEmpty_compose4
}

var cache_Data_NonEmpty_compose5 gopurs_runtime.Value
var once_Data_NonEmpty_compose5 sync.Once

func Get_Data_NonEmpty_compose5() gopurs_runtime.Value {
	once_Data_NonEmpty_compose5.Do(func() {
		cache_Data_NonEmpty_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_NonEmpty_compose5
}

var cache_Data_NonEmpty_conj gopurs_runtime.Value
var once_Data_NonEmpty_conj sync.Once

func Get_Data_NonEmpty_conj() gopurs_runtime.Value {
	once_Data_NonEmpty_conj.Do(func() {
		cache_Data_NonEmpty_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_NonEmpty_conj
}

var cache_Data_NonEmpty_NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_NonEmpty sync.Once

func Get_Data_NonEmpty_NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_NonEmpty.Do(func() {
		cache_Data_NonEmpty_NonEmpty = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_NonEmpty_NonEmpty
}

var cache_Data_NonEmpty_NonEmpty__2689777921 gopurs_runtime.Value
var once_Data_NonEmpty_NonEmpty__2689777921 sync.Once

func Get_Data_NonEmpty_NonEmpty__2689777921() gopurs_runtime.Value {
	once_Data_NonEmpty_NonEmpty__2689777921.Do(func() {
		cache_Data_NonEmpty_NonEmpty__2689777921 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_3123684004_1293498952(Call_Data_NonEmpty_NonEmpty__2689777921(__eta_norm_1_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box))))}
		})
	})
	return cache_Data_NonEmpty_NonEmpty__2689777921
}

var cache_Data_NonEmpty_NonEmpty__2838833473 gopurs_runtime.Value
var once_Data_NonEmpty_NonEmpty__2838833473 sync.Once

func Get_Data_NonEmpty_NonEmpty__2838833473() gopurs_runtime.Value {
	once_Data_NonEmpty_NonEmpty__2838833473.Do(func() {
		cache_Data_NonEmpty_NonEmpty__2838833473 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_NonEmpty_NonEmpty__2838833473(__eta_norm_1_0_box, __eta_norm_0_1_box))}
		})
	})
	return cache_Data_NonEmpty_NonEmpty__2838833473
}

var cache_Data_NonEmpty_NonEmpty__3692804292 gopurs_runtime.Value
var once_Data_NonEmpty_NonEmpty__3692804292 sync.Once

func Get_Data_NonEmpty_NonEmpty__3692804292() gopurs_runtime.Value {
	once_Data_NonEmpty_NonEmpty__3692804292.Do(func() {
		cache_Data_NonEmpty_NonEmpty__3692804292 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_3123684004_1293498952(Call_Data_NonEmpty_NonEmpty__3692804292(__eta_norm_1_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box))))}
		})
	})
	return cache_Data_NonEmpty_NonEmpty__3692804292
}

var cache_Data_NonEmpty_unfoldable1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_unfoldable1NonEmpty sync.Once

func Get_Data_NonEmpty_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_unfoldable1NonEmpty.Do(func() {
		cache_Data_NonEmpty_unfoldable1NonEmpty = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_unfoldable1NonEmpty(dictUnfoldable_0_box)
		})
	})
	return cache_Data_NonEmpty_unfoldable1NonEmpty
}

var cache_Data_NonEmpty_tail gopurs_runtime.Value
var once_Data_NonEmpty_tail sync.Once

func Get_Data_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_NonEmpty_tail.Do(func() {
		cache_Data_NonEmpty_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_NonEmpty_tail
}

var cache_Data_NonEmpty_singleton gopurs_runtime.Value
var once_Data_NonEmpty_singleton sync.Once

func Get_Data_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_NonEmpty_singleton.Do(func() {
		cache_Data_NonEmpty_singleton = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_singleton(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_0_box))
		})
	})
	return cache_Data_NonEmpty_singleton
}

var cache_Data_NonEmpty_showNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_showNonEmpty sync.Once

func Get_Data_NonEmpty_showNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_showNonEmpty.Do(func() {
		cache_Data_NonEmpty_showNonEmpty = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_showNonEmpty(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Data_NonEmpty_showNonEmpty
}

var cache_Data_NonEmpty_semigroupNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_semigroupNonEmpty sync.Once

func Get_Data_NonEmpty_semigroupNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_semigroupNonEmpty.Do(func() {
		cache_Data_NonEmpty_semigroupNonEmpty = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_semigroupNonEmpty(dictApplicative_0_box, dictSemigroup_1_box)
		})
	})
	return cache_Data_NonEmpty_semigroupNonEmpty
}

var cache_Data_NonEmpty_oneOf gopurs_runtime.Value
var once_Data_NonEmpty_oneOf sync.Once

func Get_Data_NonEmpty_oneOf() gopurs_runtime.Value {
	once_Data_NonEmpty_oneOf.Do(func() {
		cache_Data_NonEmpty_oneOf = gopurs_runtime.Func2(func(dictAlternative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
		})
	})
	return cache_Data_NonEmpty_oneOf
}

var cache_Data_NonEmpty_head gopurs_runtime.Value
var once_Data_NonEmpty_head sync.Once

func Get_Data_NonEmpty_head() gopurs_runtime.Value {
	once_Data_NonEmpty_head.Do(func() {
		cache_Data_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_head(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_NonEmpty_head
}

var cache_Data_NonEmpty_functorNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_functorNonEmpty sync.Once

func Get_Data_NonEmpty_functorNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_functorNonEmpty.Do(func() {
		cache_Data_NonEmpty_functorNonEmpty = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_functorNonEmpty(dictFunctor_0_box)
		})
	})
	return cache_Data_NonEmpty_functorNonEmpty
}

var cache_Data_NonEmpty_functorWithIndex gopurs_runtime.Value
var once_Data_NonEmpty_functorWithIndex sync.Once

func Get_Data_NonEmpty_functorWithIndex() gopurs_runtime.Value {
	once_Data_NonEmpty_functorWithIndex.Do(func() {
		cache_Data_NonEmpty_functorWithIndex = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_functorWithIndex(dictFunctorWithIndex_0_box)
		})
	})
	return cache_Data_NonEmpty_functorWithIndex
}

var cache_Data_NonEmpty_fromNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_fromNonEmpty sync.Once

func Get_Data_NonEmpty_fromNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_fromNonEmpty.Do(func() {
		cache_Data_NonEmpty_fromNonEmpty = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_fromNonEmpty(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
		})
	})
	return cache_Data_NonEmpty_fromNonEmpty
}

var cache_Data_NonEmpty_foldableNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_foldableNonEmpty sync.Once

func Get_Data_NonEmpty_foldableNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_foldableNonEmpty.Do(func() {
		cache_Data_NonEmpty_foldableNonEmpty = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_foldableNonEmpty(dictFoldable_0_box)
		})
	})
	return cache_Data_NonEmpty_foldableNonEmpty
}

var cache_Data_NonEmpty_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_foldableWithIndexNonEmpty sync.Once

func Get_Data_NonEmpty_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_NonEmpty_foldableWithIndexNonEmpty = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_box)
		})
	})
	return cache_Data_NonEmpty_foldableWithIndexNonEmpty
}

var cache_Data_NonEmpty_traversableNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_traversableNonEmpty sync.Once

func Get_Data_NonEmpty_traversableNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_traversableNonEmpty.Do(func() {
		cache_Data_NonEmpty_traversableNonEmpty = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_traversableNonEmpty(dictTraversable_0_box)
		})
	})
	return cache_Data_NonEmpty_traversableNonEmpty
}

var cache_Data_NonEmpty_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_traversableWithIndexNonEmpty sync.Once

func Get_Data_NonEmpty_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_NonEmpty_traversableWithIndexNonEmpty = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_traversableWithIndexNonEmpty(dictTraversableWithIndex_0_box)
		})
	})
	return cache_Data_NonEmpty_traversableWithIndexNonEmpty
}

var cache_Data_NonEmpty_foldable1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_foldable1NonEmpty sync.Once

func Get_Data_NonEmpty_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_foldable1NonEmpty.Do(func() {
		cache_Data_NonEmpty_foldable1NonEmpty = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_foldable1NonEmpty(dictFoldable_0_box)
		})
	})
	return cache_Data_NonEmpty_foldable1NonEmpty
}

var cache_Data_NonEmpty_foldl1 gopurs_runtime.Value
var once_Data_NonEmpty_foldl1 sync.Once

func Get_Data_NonEmpty_foldl1() gopurs_runtime.Value {
	once_Data_NonEmpty_foldl1.Do(func() {
		cache_Data_NonEmpty_foldl1 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_NonEmpty_foldl1
}

var cache_Data_NonEmpty_eqNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_eqNonEmpty sync.Once

func Get_Data_NonEmpty_eqNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_eqNonEmpty.Do(func() {
		cache_Data_NonEmpty_eqNonEmpty = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_eqNonEmpty(dictEq1_0_box, dictEq_1_box)
		})
	})
	return cache_Data_NonEmpty_eqNonEmpty
}

var cache_Data_NonEmpty_ordNonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_ordNonEmpty sync.Once

func Get_Data_NonEmpty_ordNonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_ordNonEmpty.Do(func() {
		cache_Data_NonEmpty_ordNonEmpty = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_ordNonEmpty(dictOrd1_0_box)
		})
	})
	return cache_Data_NonEmpty_ordNonEmpty
}

var cache_Data_NonEmpty_eq1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_eq1NonEmpty sync.Once

func Get_Data_NonEmpty_eq1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_eq1NonEmpty.Do(func() {
		cache_Data_NonEmpty_eq1NonEmpty = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_eq1NonEmpty(dictEq1_0_box)
		})
	})
	return cache_Data_NonEmpty_eq1NonEmpty
}

var cache_Data_NonEmpty_ord1NonEmpty gopurs_runtime.Value
var once_Data_NonEmpty_ord1NonEmpty sync.Once

func Get_Data_NonEmpty_ord1NonEmpty() gopurs_runtime.Value {
	once_Data_NonEmpty_ord1NonEmpty.Do(func() {
		cache_Data_NonEmpty_ord1NonEmpty = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_ord1NonEmpty(dictOrd1_0_box)
		})
	})
	return cache_Data_NonEmpty_ord1NonEmpty
}

type Constructor_Data_NonEmpty_NonEmpty[T_f any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}

func Call_Data_NonEmpty_NonEmpty__2689777921(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
NonEmpty__2689777921:
	for {
		if false {
			continue NonEmpty__2689777921
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, __eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__eta_norm_0_1)}})
	}
}

func Call_Data_NonEmpty_NonEmpty__2838833473(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
NonEmpty__2838833473:
	for {
		if false {
			continue NonEmpty__2838833473
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Data_NonEmpty_NonEmpty__3692804292(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
NonEmpty__3692804292:
	for {
		if false {
			continue NonEmpty__3692804292
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, __eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__eta_norm_0_1)}})
	}
}

func Call_Data_NonEmpty_unfoldable1NonEmpty(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_2751875267_2187088110((&Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope5), (ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope6)])])
		__local_var_3_1 := Rebox_Data_NonEmpty_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, b_2)))
		_ = __local_var_3_1
		// TAST (Let): __local_var_3_0 shape=Let(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope5), (TypeApp (TypeVar f$scope1) [(TypeVar a$scope5)])])
		__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{(__local_var_3_1).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 struct {
					V0 gopurs_runtime.Value
					V1 bool
				}
				{
					if v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil {
						__t2 = struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), true}
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
				}
			end_branch_2:
				return func() gopurs_runtime.Value {
					_v := __t2
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
			}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((__local_var_3_1).V1)})}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}())
		_ = __local_var_3_0
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_3_0).V0, (__local_var_3_0).V1}))}
	})})))}
}

func Call_Data_NonEmpty_tail(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (v_0).V1
}

func Call_Data_NonEmpty_singleton(dictPlus_0_loop *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictPlus_0 *Constructor_Control_Plus_Plus[gopurs_runtime.Value] = dictPlus_0_loop
	_ = dictPlus_0
	// TAST (Let): empty_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar f$scope49) [(TypeVar a$scope50)])
	empty_1_0 := Call_Control_Plus_empty(gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(dictPlus_0)})
	_ = empty_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, empty_1_0}))}
	})
}

func Call_Data_NonEmpty_showNonEmpty(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1507176707_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(NonEmpty ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), " "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), ")"))
	})})))}
}

func Call_Data_NonEmpty_semigroupNonEmpty(dictApplicative_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
	_ = dictSemigroup_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1577339395_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))}))}
	})})))}
}

func Call_Data_NonEmpty_oneOf(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], v_1_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var v_1 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "pure"), (v_1).V0), (v_1).V1)
}

func Call_Data_NonEmpty_head(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (v_0).V0
}

func Call_Data_NonEmpty_functorNonEmpty(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1468200963_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): map3_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope99) [(TypeVar a$scope103)])] (TypeApp (TypeVar f$scope99) [(TypeVar b$scope104)]))
		map3_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
		_ = map3_2_0
		return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), gopurs_runtime.Apply(map3_2_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)}))}
		})
	})})))}
}

func Call_Data_NonEmpty_functorWithIndex(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
	_ = dictFunctorWithIndex_0
	// TAST (Let): functorNonEmpty1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope86)])])
	functorNonEmpty1_1_0 := Rebox_Data_NonEmpty_2812149806_1468200963(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))))
	_ = functorNonEmpty1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_475922212_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1468200963_2812149806(functorNonEmpty1_1_0))}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, Get_Data_Maybe_Just()), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)}))}
	})})))}
}

func Call_Data_NonEmpty_fromNonEmpty(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1)
}

func Call_Data_NonEmpty_foldableNonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_3071895939_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldMap1_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope153)] (TypeVar m$scope154)), (TypeApp (TypeVar f$scope145) [(TypeVar a$scope153)])] (TypeVar m$scope154))
		foldMap1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))})
		_ = foldMap1_2_0
		return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(foldMap1_2_0, f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
		})
	}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
	}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
	})})))}
}

func Call_Data_NonEmpty_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): foldableNonEmpty1_1_0 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope120)])])
	foldableNonEmpty1_1_0 := Rebox_Data_NonEmpty_1680800814_3071895939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))))
	_ = foldableNonEmpty1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1595536100_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_3071895939_1680800814(foldableNonEmpty1_1_0))}
	}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): foldMapWithIndex1_3_1 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope119), (TypeVar a$scope129)] (TypeVar m$scope130)), (TypeApp (TypeVar f$scope120) [(TypeVar a$scope129)])] (TypeVar m$scope130))
		foldMapWithIndex1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))})
		_ = foldMapWithIndex1_3_1
		return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(foldMapWithIndex1_3_1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Maybe_Just()), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
		})
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, Get_Data_Maybe_Just()), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}, b_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, Get_Data_Maybe_Just()), b_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
	})})))}
}

func Call_Data_NonEmpty_traversableNonEmpty(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): functorNonEmpty1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope27)])])
	functorNonEmpty1_1_0 := Rebox_Data_NonEmpty_2812149806_1468200963(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))))
	_ = functorNonEmpty1_1_0
	// TAST (Let): foldableNonEmpty1_2_1 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope27)])])
	foldableNonEmpty1_2_1 := Rebox_Data_NonEmpty_1680800814_3071895939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))))
	_ = foldableNonEmpty1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_4249989635_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_3071895939_1680800814(foldableNonEmpty1_2_1))}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1468200963_2812149806(functorNonEmpty1_1_0))}
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): sequence1_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope27) [(TypeApp (TypeVar m$scope35) [(TypeVar a$scope34)])])] (TypeApp (TypeVar m$scope35) [(TypeApp (TypeVar f$scope27) [(TypeVar a$scope34)])]))
		sequence1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = sequence1_4_2
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply(sequence1_4_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
		})
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverse1_4_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope38)] (TypeApp (TypeVar m$scope40) [(TypeVar b$scope39)])), (TypeApp (TypeVar f$scope27) [(TypeVar a$scope38)])] (TypeApp (TypeVar m$scope40) [(TypeApp (TypeVar f$scope27) [(TypeVar b$scope39)])]))
		traverse1_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = traverse1_4_3
		return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)), gopurs_runtime.Apply2(traverse1_4_3, f_5, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
		})
	})})))}
}

func Call_Data_NonEmpty_traversableWithIndexNonEmpty(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
	_ = dictTraversableWithIndex_0
	// TAST (Let): functorWithIndex1_1_0 shape=App(Var) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar i$scope11)]), (ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope12)])])
	functorWithIndex1_1_0 := Rebox_Data_NonEmpty_2412140840_475922212(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_functorWithIndex(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))))
	_ = functorWithIndex1_1_0
	// TAST (Let): foldableWithIndexNonEmpty1_2_1 shape=App(Var) bindingType=(ADT ["Data","FoldableWithIndex","FoldableWithIndex"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar i$scope11)]), (ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope12)])])
	foldableWithIndexNonEmpty1_2_1 := Rebox_Data_NonEmpty_3725484264_1595536100(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))))
	_ = foldableWithIndexNonEmpty1_2_1
	// TAST (Let): traversableNonEmpty1_3_2 shape=App(Var) bindingType=(ADT ["Data","Traversable","Traversable"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope12)])])
	traversableNonEmpty1_3_2 := Rebox_Data_NonEmpty_3043886126_4249989635(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))))
	_ = traversableNonEmpty1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_705974564_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1595536100_3725484264(foldableWithIndexNonEmpty1_2_1))}
	}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_475922212_2412140840(functorWithIndex1_1_0))}
	}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_4249989635_3043886126(traversableNonEmpty1_3_2))}
	}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverseWithIndex1_5_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar i$scope11), (TypeVar a$scope18)] (TypeApp (TypeVar m$scope20) [(TypeVar b$scope19)])), (TypeApp (TypeVar f$scope12) [(TypeVar a$scope18)])] (TypeApp (TypeVar m$scope20) [(TypeApp (TypeVar f$scope12) [(TypeVar b$scope19)])]))
		traverseWithIndex1_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_4))})
		_ = traverseWithIndex1_5_3
		return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply2(f_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)), gopurs_runtime.Apply2(traverseWithIndex1_5_3, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_6, Get_Data_Maybe_Just()), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
		})
	})})))}
}

func Call_Data_NonEmpty_foldable1NonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): foldableNonEmpty1_1_0 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope160)])])
	foldableNonEmpty1_1_0 := Rebox_Data_NonEmpty_1680800814_3071895939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_NonEmpty_foldableNonEmpty(dictFoldable_0)))
	_ = foldableNonEmpty1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_4217626592_4151366573((&Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_3071895939_1680800814(foldableNonEmpty1_1_0))}
	}), gopurs_runtime.Func3(func(dictSemigroup_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(s_5 gopurs_runtime.Value, a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), s_5, gopurs_runtime.Apply(f_3, a1_6))
		}), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(Func [(TypeVar a$scope168)] (TypeVar a$scope168))
		__local_var_4_1 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)
		_ = __local_var_4_1
		// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope168)])
		__local_var_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=(Func [(TypeVar a$scope168)] (TypeVar a$scope168))
			__local_var_6_3 := gopurs_runtime.Apply(f_2, a1_5)
			_ = __local_var_6_3
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr == nil {
						__t4 = a1_5
						goto end_branch_4
					} else {

					}
				}
				{
					if v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr != nil {
						__t4 = gopurs_runtime.Apply(__local_var_6_3, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			}))
		}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
		_ = __local_var_5_2
		var __t5 gopurs_runtime.Value
		{
			if __local_var_5_2 == nil {
				__t5 = (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
				goto end_branch_5
			} else {

			}
		}
		{
			if __local_var_5_2 != nil {
				__t5 = gopurs_runtime.Apply(__local_var_4_1, (__local_var_5_2).V0)
				goto end_branch_5
			} else {

			}
		}
		{
			__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_5:
		return __t5
	})})))}
}

func Call_Data_NonEmpty_foldl1(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return Call_Data_Semigroup_Foldable_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)})))
}

func Call_Data_NonEmpty_eqNonEmpty(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	// TAST (Let): eq11_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope175) [(TypeVar a$scope176)]), (TypeApp (TypeVar f$scope175) [(TypeVar a$scope176)])] Boolean)
	eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))})
	_ = eq11_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_2176830691_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(eq11_2_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
	})})))}
}

func Call_Data_NonEmpty_ordNonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): eqNonEmpty1_1_0 shape=App(Var) bindingType=Any
	eqNonEmpty1_1_0 := gopurs_runtime.Apply(Get_Data_NonEmpty_eqNonEmpty(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
	_ = eqNonEmpty1_1_0
	return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): compare11_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope67) [(TypeVar a$scope68)]), (TypeApp (TypeVar f$scope67) [(TypeVar a$scope68)])] (ADT ["Data","Ordering","Ordering"] []))
		compare11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_2))})
		_ = compare11_3_1
		// TAST (Let): eqNonEmpty2_4_2 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope67), (TypeVar a$scope68)])])
		eqNonEmpty2_4_2 := Rebox_Data_NonEmpty_3790796878_2176830691(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqNonEmpty1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))))
		_ = eqNonEmpty2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_167870147_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_2176830691_3790796878(eqNonEmpty2_4_2))}
		}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_7_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
			v_7_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_2, "compare"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal)
			_ = v_7_3
			var __t4 gopurs_runtime.Value
			{
				if v_7_3 == 1527465420 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					goto end_branch_4
				} else {

				}
			}
			{
				if v_7_3 == 380165415 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = gopurs_runtime.Apply2(compare11_3_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V1)
			}
		end_branch_4:
			return __t4
		})})))}
	})
}

func Call_Data_NonEmpty_eq1NonEmpty(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1625813138_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_NonEmpty_eqNonEmpty(dictEq1_0, dictEq_1)))
	})})))}
}

func Call_Data_NonEmpty_ord1NonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): ordNonEmpty1_1_0 shape=App(Var) bindingType=Any
	ordNonEmpty1_1_0 := Call_Data_NonEmpty_ordNonEmpty(dictOrd1_0)
	_ = ordNonEmpty1_1_0
	// TAST (Let): eq1NonEmpty1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq1"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f$scope71)])])
	eq1NonEmpty1_2_1 := Rebox_Data_NonEmpty_1766074591_1625813138(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_NonEmpty_eq1NonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))))
	_ = eq1NonEmpty1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_937092146_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_NonEmpty_1625813138_1766074591(eq1NonEmpty1_2_1))}
	}), gopurs_runtime.Func(func(dictOrd_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordNonEmpty1_1_0, dictOrd_3)))
	})})))}
}

func Rebox_Data_NonEmpty_138441832_3804580809(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_NonEmpty_1468200963_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_1507176707_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_1577339395_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_1595536100_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_1625813138_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_167870147_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_1680800814_3071895939(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_1766074591_1625813138(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_2176830691_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_2412140840_475922212(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_2751875267_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_2812149806_1468200963(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_3043886126_4249989635(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_3071895939_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_3725484264_1595536100(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_3790796878_2176830691(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_4217626592_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_4249989635_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_475922212_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_705974564_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_NonEmpty_937092146_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord1[gopurs_runtime.Value])(unsafe.Pointer(in))
}
