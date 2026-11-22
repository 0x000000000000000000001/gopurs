package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_sequence gopurs_runtime.Value
var once_Data_Array_sequence sync.Once

func Get_Data_Array_sequence() gopurs_runtime.Value {
	once_Data_Array_sequence.Do(func() {
		cache_Data_Array_sequence = gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V2
	})
	return cache_Data_Array_sequence
}

var cache_Data_Array_traverse_ gopurs_runtime.Value
var once_Data_Array_traverse_ sync.Once

func Get_Data_Array_traverse_() gopurs_runtime.Value {
	once_Data_Array_traverse_.Do(func() {
		cache_Data_Array_traverse_ = Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST()))
	})
	return cache_Data_Array_traverse_
}

var cache_Data_Array_lessThan gopurs_runtime.Value
var once_Data_Array_lessThan sync.Once

func Get_Data_Array_lessThan() gopurs_runtime.Value {
	once_Data_Array_lessThan.Do(func() {
		cache_Data_Array_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_Array_lessThan
}

var cache_Data_Array_add gopurs_runtime.Value
var once_Data_Array_add sync.Once

func Get_Data_Array_add() gopurs_runtime.Value {
	once_Data_Array_add.Do(func() {
		cache_Data_Array_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_Array_add
}

var cache_Data_Array_negate gopurs_runtime.Value
var once_Data_Array_negate sync.Once

func Get_Data_Array_negate() gopurs_runtime.Value {
	once_Data_Array_negate.Do(func() {
		cache_Data_Array_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Array_negate(a_0_box.IntVal))
		})
	})
	return cache_Data_Array_negate
}

var cache_Data_Array_lessThanOrEq gopurs_runtime.Value
var once_Data_Array_lessThanOrEq sync.Once

func Get_Data_Array_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Array_lessThanOrEq.Do(func() {
		cache_Data_Array_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_Array_lessThanOrEq
}

var cache_Data_Array_eq gopurs_runtime.Value
var once_Data_Array_eq sync.Once

func Get_Data_Array_eq() gopurs_runtime.Value {
	once_Data_Array_eq.Do(func() {
		cache_Data_Array_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_Array_eq
}

var cache_Data_Array_traverse_1 gopurs_runtime.Value
var once_Data_Array_traverse_1 sync.Once

func Get_Data_Array_traverse_1() gopurs_runtime.Value {
	once_Data_Array_traverse_1.Do(func() {
		cache_Data_Array_traverse_1 = Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST()))
	})
	return cache_Data_Array_traverse_1
}

var cache_Data_Array_bind gopurs_runtime.Value
var once_Data_Array_bind sync.Once

func Get_Data_Array_bind() gopurs_runtime.Value {
	once_Data_Array_bind.Do(func() {
		cache_Data_Array_bind = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind
}

var cache_Data_Array_bind1 gopurs_runtime.Value
var once_Data_Array_bind1 sync.Once

func Get_Data_Array_bind1() gopurs_runtime.Value {
	once_Data_Array_bind1.Do(func() {
		cache_Data_Array_bind1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind1
}

var cache_Data_Array_discard gopurs_runtime.Value
var once_Data_Array_discard sync.Once

func Get_Data_Array_discard() gopurs_runtime.Value {
	once_Data_Array_discard.Do(func() {
		cache_Data_Array_discard = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_discard
}

var cache_Data_Array_bind2 gopurs_runtime.Value
var once_Data_Array_bind2 sync.Once

func Get_Data_Array_bind2() gopurs_runtime.Value {
	once_Data_Array_bind2.Do(func() {
		cache_Data_Array_bind2 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind2
}

var cache_Data_Array_void gopurs_runtime.Value
var once_Data_Array_void sync.Once

func Get_Data_Array_void() gopurs_runtime.Value {
	once_Data_Array_void.Do(func() {
		cache_Data_Array_void = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_void(__local_var_0_box)
		})
	})
	return cache_Data_Array_void
}

var cache_Data_Array_pure gopurs_runtime.Value
var once_Data_Array_pure sync.Once

func Get_Data_Array_pure() gopurs_runtime.Value {
	once_Data_Array_pure.Do(func() {
		cache_Data_Array_pure = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Data_Array_pure
}

var cache_Data_Array_intercalate1 gopurs_runtime.Value
var once_Data_Array_intercalate1 sync.Once

func Get_Data_Array_intercalate1() gopurs_runtime.Value {
	once_Data_Array_intercalate1.Do(func() {
		cache_Data_Array_intercalate1 = gopurs_runtime.Apply(Get_Data_Foldable_intercalate(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})
	})
	return cache_Data_Array_intercalate1
}

var cache_Data_Array_zero gopurs_runtime.Value
var once_Data_Array_zero sync.Once

func Get_Data_Array_zero() gopurs_runtime.Value {
	once_Data_Array_zero.Do(func() {
		cache_Data_Array_zero = gopurs_runtime.Int(int64(0))
	})
	return cache_Data_Array_zero
}

var cache_Data_Array_sub gopurs_runtime.Value
var once_Data_Array_sub sync.Once

func Get_Data_Array_sub() gopurs_runtime.Value {
	once_Data_Array_sub.Do(func() {
		cache_Data_Array_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Array_sub
}

var cache_Data_Array_one gopurs_runtime.Value
var once_Data_Array_one sync.Once

func Get_Data_Array_one() gopurs_runtime.Value {
	once_Data_Array_one.Do(func() {
		cache_Data_Array_one = gopurs_runtime.Int(int64(1))
	})
	return cache_Data_Array_one
}

var cache_Data_Array_apply gopurs_runtime.Value
var once_Data_Array_apply sync.Once

func Get_Data_Array_apply() gopurs_runtime.Value {
	once_Data_Array_apply.Do(func() {
		cache_Data_Array_apply = Rebox_Data_Array_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1
	})
	return cache_Data_Array_apply
}

var cache_Data_Array_go__map gopurs_runtime.Value
var once_Data_Array_go__map sync.Once

func Get_Data_Array_go__map() gopurs_runtime.Value {
	once_Data_Array_go__map.Do(func() {
		cache_Data_Array_go__map = Rebox_Data_Array_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Array_go__map
}

var cache_Data_Array_bind3 gopurs_runtime.Value
var once_Data_Array_bind3 sync.Once

func Get_Data_Array_bind3() gopurs_runtime.Value {
	once_Data_Array_bind3.Do(func() {
		cache_Data_Array_bind3 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind3
}

var cache_Data_Array_bind4 gopurs_runtime.Value
var once_Data_Array_bind4 sync.Once

func Get_Data_Array_bind4() gopurs_runtime.Value {
	once_Data_Array_bind4.Do(func() {
		cache_Data_Array_bind4 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind4
}

var cache_Data_Array_bind5 gopurs_runtime.Value
var once_Data_Array_bind5 sync.Once

func Get_Data_Array_bind5() gopurs_runtime.Value {
	once_Data_Array_bind5.Do(func() {
		cache_Data_Array_bind5 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind5
}

var cache_Data_Array_discard1 gopurs_runtime.Value
var once_Data_Array_discard1 sync.Once

func Get_Data_Array_discard1() gopurs_runtime.Value {
	once_Data_Array_discard1.Do(func() {
		cache_Data_Array_discard1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_discard1
}

var cache_Data_Array_discard2 gopurs_runtime.Value
var once_Data_Array_discard2 sync.Once

func Get_Data_Array_discard2() gopurs_runtime.Value {
	once_Data_Array_discard2.Do(func() {
		cache_Data_Array_discard2 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_discard2
}

var cache_Data_Array_void1 gopurs_runtime.Value
var once_Data_Array_void1 sync.Once

func Get_Data_Array_void1() gopurs_runtime.Value {
	once_Data_Array_void1.Do(func() {
		cache_Data_Array_void1 = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST()))
	})
	return cache_Data_Array_void1
}

var cache_Data_Array_bind6 gopurs_runtime.Value
var once_Data_Array_bind6 sync.Once

func Get_Data_Array_bind6() gopurs_runtime.Value {
	once_Data_Array_bind6.Do(func() {
		cache_Data_Array_bind6 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind6
}

var cache_Data_Array_bind7 gopurs_runtime.Value
var once_Data_Array_bind7 sync.Once

func Get_Data_Array_bind7() gopurs_runtime.Value {
	once_Data_Array_bind7.Do(func() {
		cache_Data_Array_bind7 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind7
}

var cache_Data_Array_pure1 gopurs_runtime.Value
var once_Data_Array_pure1 sync.Once

func Get_Data_Array_pure1() gopurs_runtime.Value {
	once_Data_Array_pure1.Do(func() {
		cache_Data_Array_pure1 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Data_Array_pure1
}

var cache_Data_Array_map1 gopurs_runtime.Value
var once_Data_Array_map1 sync.Once

func Get_Data_Array_map1() gopurs_runtime.Value {
	once_Data_Array_map1.Do(func() {
		cache_Data_Array_map1 = Get_Data_Functor_arrayMap()
	})
	return cache_Data_Array_map1
}

var cache_Data_Array_bind8 gopurs_runtime.Value
var once_Data_Array_bind8 sync.Once

func Get_Data_Array_bind8() gopurs_runtime.Value {
	once_Data_Array_bind8.Do(func() {
		cache_Data_Array_bind8 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind8
}

var cache_Data_Array_discard3 gopurs_runtime.Value
var once_Data_Array_discard3 sync.Once

func Get_Data_Array_discard3() gopurs_runtime.Value {
	once_Data_Array_discard3.Do(func() {
		cache_Data_Array_discard3 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_discard3
}

var cache_Data_Array_bind9 gopurs_runtime.Value
var once_Data_Array_bind9 sync.Once

func Get_Data_Array_bind9() gopurs_runtime.Value {
	once_Data_Array_bind9.Do(func() {
		cache_Data_Array_bind9 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind9
}

var cache_Data_Array_map2 gopurs_runtime.Value
var once_Data_Array_map2 sync.Once

func Get_Data_Array_map2() gopurs_runtime.Value {
	once_Data_Array_map2.Do(func() {
		cache_Data_Array_map2 = Get_Control_Monad_ST_Internal_map_()
	})
	return cache_Data_Array_map2
}

var cache_Data_Array_compose gopurs_runtime.Value
var once_Data_Array_compose sync.Once

func Get_Data_Array_compose() gopurs_runtime.Value {
	once_Data_Array_compose.Do(func() {
		cache_Data_Array_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_compose
}

var cache_Data_Array_compose1 gopurs_runtime.Value
var once_Data_Array_compose1 sync.Once

func Get_Data_Array_compose1() gopurs_runtime.Value {
	once_Data_Array_compose1.Do(func() {
		cache_Data_Array_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_compose1
}

var cache_Data_Array_fromJust gopurs_runtime.Value
var once_Data_Array_fromJust sync.Once

func Get_Data_Array_fromJust() gopurs_runtime.Value {
	once_Data_Array_fromJust.Do(func() {
		cache_Data_Array_fromJust = gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{})
	})
	return cache_Data_Array_fromJust
}

var cache_Data_Array_when gopurs_runtime.Value
var once_Data_Array_when sync.Once

func Get_Data_Array_when() gopurs_runtime.Value {
	once_Data_Array_when.Do(func() {
		cache_Data_Array_when = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_when((v_0_box.IntVal) != (0), v1_1_box)
		})
	})
	return cache_Data_Array_when
}

var cache_Data_Array_notEq gopurs_runtime.Value
var once_Data_Array_notEq sync.Once

func Get_Data_Array_notEq() gopurs_runtime.Value {
	once_Data_Array_notEq.Do(func() {
		cache_Data_Array_notEq = Get_Data_Eq_notEq__2123013075()
	})
	return cache_Data_Array_notEq
}

var cache_Data_Array_void2 gopurs_runtime.Value
var once_Data_Array_void2 sync.Once

func Get_Data_Array_void2() gopurs_runtime.Value {
	once_Data_Array_void2.Do(func() {
		cache_Data_Array_void2 = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST()))
	})
	return cache_Data_Array_void2
}

var cache_Data_Array_bind10 gopurs_runtime.Value
var once_Data_Array_bind10 sync.Once

func Get_Data_Array_bind10() gopurs_runtime.Value {
	once_Data_Array_bind10.Do(func() {
		cache_Data_Array_bind10 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind10
}

var cache_Data_Array_bind11 gopurs_runtime.Value
var once_Data_Array_bind11 sync.Once

func Get_Data_Array_bind11() gopurs_runtime.Value {
	once_Data_Array_bind11.Do(func() {
		cache_Data_Array_bind11 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind11
}

var cache_Data_Array_discard4 gopurs_runtime.Value
var once_Data_Array_discard4 sync.Once

func Get_Data_Array_discard4() gopurs_runtime.Value {
	once_Data_Array_discard4.Do(func() {
		cache_Data_Array_discard4 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_discard4
}

var cache_Data_Array_void3 gopurs_runtime.Value
var once_Data_Array_void3 sync.Once

func Get_Data_Array_void3() gopurs_runtime.Value {
	once_Data_Array_void3.Do(func() {
		cache_Data_Array_void3 = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_void3(__local_var_0_box)
		})
	})
	return cache_Data_Array_void3
}

var cache_Data_Array_bind12 gopurs_runtime.Value
var once_Data_Array_bind12 sync.Once

func Get_Data_Array_bind12() gopurs_runtime.Value {
	once_Data_Array_bind12.Do(func() {
		cache_Data_Array_bind12 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind12
}

var cache_Data_Array_bind13 gopurs_runtime.Value
var once_Data_Array_bind13 sync.Once

func Get_Data_Array_bind13() gopurs_runtime.Value {
	once_Data_Array_bind13.Do(func() {
		cache_Data_Array_bind13 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind13
}

var cache_Data_Array_discard5 gopurs_runtime.Value
var once_Data_Array_discard5 sync.Once

func Get_Data_Array_discard5() gopurs_runtime.Value {
	once_Data_Array_discard5.Do(func() {
		cache_Data_Array_discard5 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_discard5
}

var cache_Data_Array_bind14 gopurs_runtime.Value
var once_Data_Array_bind14 sync.Once

func Get_Data_Array_bind14() gopurs_runtime.Value {
	once_Data_Array_bind14.Do(func() {
		cache_Data_Array_bind14 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind14
}

var cache_Data_Array_compose2 gopurs_runtime.Value
var once_Data_Array_compose2 sync.Once

func Get_Data_Array_compose2() gopurs_runtime.Value {
	once_Data_Array_compose2.Do(func() {
		cache_Data_Array_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_compose2
}

var cache_Data_Array_eq1 gopurs_runtime.Value
var once_Data_Array_eq1 sync.Once

func Get_Data_Array_eq1() gopurs_runtime.Value {
	once_Data_Array_eq1.Do(func() {
		cache_Data_Array_eq1 = Get_Data_Eq_eq__2123013075()
	})
	return cache_Data_Array_eq1
}

var cache_Data_Array_greaterThanOrEq gopurs_runtime.Value
var once_Data_Array_greaterThanOrEq sync.Once

func Get_Data_Array_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_Array_greaterThanOrEq.Do(func() {
		cache_Data_Array_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_Array_greaterThanOrEq
}

var cache_Data_Array_foldMap1 gopurs_runtime.Value
var once_Data_Array_foldMap1 sync.Once

func Get_Data_Array_foldMap1() gopurs_runtime.Value {
	once_Data_Array_foldMap1.Do(func() {
		cache_Data_Array_foldMap1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()).V0
	})
	return cache_Data_Array_foldMap1
}

var cache_Data_Array_fold1 gopurs_runtime.Value
var once_Data_Array_fold1 sync.Once

func Get_Data_Array_fold1() gopurs_runtime.Value {
	once_Data_Array_fold1.Do(func() {
		cache_Data_Array_fold1 = gopurs_runtime.Apply(Get_Data_Foldable_fold(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})
	})
	return cache_Data_Array_fold1
}

var cache_Data_Array_fromJust1 gopurs_runtime.Value
var once_Data_Array_fromJust1 sync.Once

func Get_Data_Array_fromJust1() gopurs_runtime.Value {
	once_Data_Array_fromJust1.Do(func() {
		cache_Data_Array_fromJust1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_fromJust1(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))))
		})
	})
	return cache_Data_Array_fromJust1
}

var cache_Data_Array_not gopurs_runtime.Value
var once_Data_Array_not sync.Once

func Get_Data_Array_not() gopurs_runtime.Value {
	once_Data_Array_not.Do(func() {
		cache_Data_Array_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Data_Array_not
}

var cache_Data_Array_map3 gopurs_runtime.Value
var once_Data_Array_map3 sync.Once

func Get_Data_Array_map3() gopurs_runtime.Value {
	once_Data_Array_map3.Do(func() {
		cache_Data_Array_map3 = Rebox_Data_Array_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Array_map3
}

var cache_Data_Array_fromJust2 gopurs_runtime.Value
var once_Data_Array_fromJust2 sync.Once

func Get_Data_Array_fromJust2() gopurs_runtime.Value {
	once_Data_Array_fromJust2.Do(func() {
		cache_Data_Array_fromJust2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_fromJust2(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))))
		})
	})
	return cache_Data_Array_fromJust2
}

var cache_Data_Array_go__append gopurs_runtime.Value
var once_Data_Array_go__append sync.Once

func Get_Data_Array_go__append() gopurs_runtime.Value {
	once_Data_Array_go__append.Do(func() {
		cache_Data_Array_go__append = Get_Data_Semigroup_concatArray()
	})
	return cache_Data_Array_go__append
}

var cache_Data_Array_append__3388884614 gopurs_runtime.Value
var once_Data_Array_append__3388884614 sync.Once

func Get_Data_Array_append__3388884614() gopurs_runtime.Value {
	once_Data_Array_append__3388884614.Do(func() {
		cache_Data_Array_append__3388884614 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_append__3388884614(*((*[]gopurs_runtime.Value)(__eta_norm_1_0_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_append__3388884614
}

var cache_Data_Array_compose3 gopurs_runtime.Value
var once_Data_Array_compose3 sync.Once

func Get_Data_Array_compose3() gopurs_runtime.Value {
	once_Data_Array_compose3.Do(func() {
		cache_Data_Array_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_compose3
}

var cache_Data_Array_composeFlipped gopurs_runtime.Value
var once_Data_Array_composeFlipped sync.Once

func Get_Data_Array_composeFlipped() gopurs_runtime.Value {
	once_Data_Array_composeFlipped.Do(func() {
		cache_Data_Array_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_composeFlipped
}

var cache_Data_Array_traverse gopurs_runtime.Value
var once_Data_Array_traverse sync.Once

func Get_Data_Array_traverse() gopurs_runtime.Value {
	once_Data_Array_traverse.Do(func() {
		cache_Data_Array_traverse = gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3
	})
	return cache_Data_Array_traverse
}

var cache_Data_Array_bind15 gopurs_runtime.Value
var once_Data_Array_bind15 sync.Once

func Get_Data_Array_bind15() gopurs_runtime.Value {
	once_Data_Array_bind15.Do(func() {
		cache_Data_Array_bind15 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind15
}

var cache_Data_Array_discard6 gopurs_runtime.Value
var once_Data_Array_discard6 sync.Once

func Get_Data_Array_discard6() gopurs_runtime.Value {
	once_Data_Array_discard6.Do(func() {
		cache_Data_Array_discard6 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_discard6
}

var cache_Data_Array_bind16 gopurs_runtime.Value
var once_Data_Array_bind16 sync.Once

func Get_Data_Array_bind16() gopurs_runtime.Value {
	once_Data_Array_bind16.Do(func() {
		cache_Data_Array_bind16 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Data_Array_bind16
}

var cache_Data_Array_map4 gopurs_runtime.Value
var once_Data_Array_map4 sync.Once

func Get_Data_Array_map4() gopurs_runtime.Value {
	once_Data_Array_map4.Do(func() {
		cache_Data_Array_map4 = Get_Control_Monad_ST_Internal_map_()
	})
	return cache_Data_Array_map4
}

var cache_Data_Array_compose4 gopurs_runtime.Value
var once_Data_Array_compose4 sync.Once

func Get_Data_Array_compose4() gopurs_runtime.Value {
	once_Data_Array_compose4.Do(func() {
		cache_Data_Array_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_compose4
}

var cache_Data_Array_not1 gopurs_runtime.Value
var once_Data_Array_not1 sync.Once

func Get_Data_Array_not1() gopurs_runtime.Value {
	once_Data_Array_not1.Do(func() {
		cache_Data_Array_not1 = Get_Data_HeytingAlgebra_boolNot()
	})
	return cache_Data_Array_not1
}

var cache_Data_Array_when1 gopurs_runtime.Value
var once_Data_Array_when1 sync.Once

func Get_Data_Array_when1() gopurs_runtime.Value {
	once_Data_Array_when1.Do(func() {
		cache_Data_Array_when1 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_when1((v_0_box.IntVal) != (0), v1_1_box)
		})
	})
	return cache_Data_Array_when1
}

var cache_Data_Array_void4 gopurs_runtime.Value
var once_Data_Array_void4 sync.Once

func Get_Data_Array_void4() gopurs_runtime.Value {
	once_Data_Array_void4.Do(func() {
		cache_Data_Array_void4 = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST()))
	})
	return cache_Data_Array_void4
}

var cache_Data_Array_append1 gopurs_runtime.Value
var once_Data_Array_append1 sync.Once

func Get_Data_Array_append1() gopurs_runtime.Value {
	once_Data_Array_append1.Do(func() {
		cache_Data_Array_append1 = Get_Data_Semigroup_concatArray()
	})
	return cache_Data_Array_append1
}

var cache_Data_Array_zipWith gopurs_runtime.Value
var once_Data_Array_zipWith sync.Once

func Get_Data_Array_zipWith() gopurs_runtime.Value {
	once_Data_Array_zipWith.Do(func() {
		cache_Data_Array_zipWith = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_zipWith(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_zipWith
}

var cache_Data_Array_zipWithA gopurs_runtime.Value
var once_Data_Array_zipWithA sync.Once

func Get_Data_Array_zipWithA() gopurs_runtime.Value {
	once_Data_Array_zipWithA.Do(func() {
		cache_Data_Array_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_Array_zipWithA
}

var cache_Data_Array_zip gopurs_runtime.Value
var once_Data_Array_zip sync.Once

func Get_Data_Array_zip() gopurs_runtime.Value {
	once_Data_Array_zip.Do(func() {
		cache_Data_Array_zip = gopurs_runtime.Apply(Get_Data_Array_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_Array_zip
}

var cache_Data_Array_updateAtIndices gopurs_runtime.Value
var once_Data_Array_updateAtIndices sync.Once

func Get_Data_Array_updateAtIndices() gopurs_runtime.Value {
	once_Data_Array_updateAtIndices.Do(func() {
		cache_Data_Array_updateAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_updateAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_Array_updateAtIndices
}

var cache_Data_Array_updateAt gopurs_runtime.Value
var once_Data_Array_updateAt sync.Once

func Get_Data_Array_updateAt() gopurs_runtime.Value {
	once_Data_Array_updateAt.Do(func() {
		cache_Data_Array_updateAt = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_updateAt(__local_var_0_box.IntVal, __local_var_1_box, *((*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_updateAt
}

var cache_Data_Array_unsafeIndex gopurs_runtime.Value
var once_Data_Array_unsafeIndex sync.Once

func Get_Data_Array_unsafeIndex() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex.Do(func() {
		cache_Data_Array_unsafeIndex = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_unsafeIndex(_dollar___unused_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)), __local_var_2_box.IntVal)
		})
	})
	return cache_Data_Array_unsafeIndex
}

var cache_Data_Array_unsafeIndex1 gopurs_runtime.Value
var once_Data_Array_unsafeIndex1 sync.Once

func Get_Data_Array_unsafeIndex1() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex1.Do(func() {
		cache_Data_Array_unsafeIndex1 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_unsafeIndex1(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)), __local_var_1_box.IntVal)
		})
	})
	return cache_Data_Array_unsafeIndex1
}

var cache_Data_Array_unsafeIndex2 gopurs_runtime.Value
var once_Data_Array_unsafeIndex2 sync.Once

func Get_Data_Array_unsafeIndex2() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex2.Do(func() {
		cache_Data_Array_unsafeIndex2 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_unsafeIndex2(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)), __local_var_1_box.IntVal)
		})
	})
	return cache_Data_Array_unsafeIndex2
}

var cache_Data_Array_unsafeIndex3 gopurs_runtime.Value
var once_Data_Array_unsafeIndex3 sync.Once

func Get_Data_Array_unsafeIndex3() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex3.Do(func() {
		cache_Data_Array_unsafeIndex3 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_unsafeIndex3(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)), __local_var_1_box.IntVal)
		})
	})
	return cache_Data_Array_unsafeIndex3
}

var cache_Data_Array_unsafeIndex4 gopurs_runtime.Value
var once_Data_Array_unsafeIndex4 sync.Once

func Get_Data_Array_unsafeIndex4() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex4.Do(func() {
		cache_Data_Array_unsafeIndex4 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_unsafeIndex4(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)), __local_var_1_box.IntVal)
		})
	})
	return cache_Data_Array_unsafeIndex4
}

var cache_Data_Array_uncons gopurs_runtime.Value
var once_Data_Array_uncons sync.Once

func Get_Data_Array_uncons() gopurs_runtime.Value {
	once_Data_Array_uncons.Do(func() {
		cache_Data_Array_uncons = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_uncons(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Array(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_uncons
}

var cache_Data_Array_toUnfoldable gopurs_runtime.Value
var once_Data_Array_toUnfoldable sync.Once

func Get_Data_Array_toUnfoldable() gopurs_runtime.Value {
	once_Data_Array_toUnfoldable.Do(func() {
		cache_Data_Array_toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_toUnfoldable
}

var cache_Data_Array_tail gopurs_runtime.Value
var once_Data_Array_tail sync.Once

func Get_Data_Array_tail() gopurs_runtime.Value {
	once_Data_Array_tail.Do(func() {
		cache_Data_Array_tail = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_tail(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_tail
}

var cache_Data_Array_sortBy gopurs_runtime.Value
var once_Data_Array_sortBy sync.Once

func Get_Data_Array_sortBy() gopurs_runtime.Value {
	once_Data_Array_sortBy.Do(func() {
		cache_Data_Array_sortBy = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_sortBy(comp_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_sortBy
}

var cache_Data_Array_sortWith gopurs_runtime.Value
var once_Data_Array_sortWith sync.Once

func Get_Data_Array_sortWith() gopurs_runtime.Value {
	once_Data_Array_sortWith.Do(func() {
		cache_Data_Array_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
		})
	})
	return cache_Data_Array_sortWith
}

var cache_Data_Array_sortWith1 gopurs_runtime.Value
var once_Data_Array_sortWith1 sync.Once

func Get_Data_Array_sortWith1() gopurs_runtime.Value {
	once_Data_Array_sortWith1.Do(func() {
		cache_Data_Array_sortWith1 = gopurs_runtime.Apply(Get_Data_Array_sortWith(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_3308271157_4177771502(Rebox_Data_Array_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))})
	})
	return cache_Data_Array_sortWith1
}

var cache_Data_Array_sort gopurs_runtime.Value
var once_Data_Array_sort sync.Once

func Get_Data_Array_sort() gopurs_runtime.Value {
	once_Data_Array_sort.Do(func() {
		cache_Data_Array_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_sort
}

var cache_Data_Array_snoc gopurs_runtime.Value
var once_Data_Array_snoc sync.Once

func Get_Data_Array_snoc() gopurs_runtime.Value {
	once_Data_Array_snoc.Do(func() {
		cache_Data_Array_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_snoc(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)), x_1_box))
		})
	})
	return cache_Data_Array_snoc
}

var cache_Data_Array_slice gopurs_runtime.Value
var once_Data_Array_slice sync.Once

func Get_Data_Array_slice() gopurs_runtime.Value {
	once_Data_Array_slice.Do(func() {
		cache_Data_Array_slice = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_slice(__local_var_0_box.IntVal, __local_var_1_box.IntVal, *((*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_slice
}

var cache_Data_Array_splitAt gopurs_runtime.Value
var once_Data_Array_splitAt sync.Once

func Get_Data_Array_splitAt() gopurs_runtime.Value {
	once_Data_Array_splitAt.Do(func() {
		cache_Data_Array_splitAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_Array_splitAt(v_0_box.IntVal, *((*[]gopurs_runtime.Value)(v1_1_box.UnsafePtr)))
				_ = orig
				return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(orig.after), gopurs_runtime.Array(orig.before))
			}()
		})
	})
	return cache_Data_Array_splitAt
}

var cache_Data_Array_take gopurs_runtime.Value
var once_Data_Array_take sync.Once

func Get_Data_Array_take() gopurs_runtime.Value {
	once_Data_Array_take.Do(func() {
		cache_Data_Array_take = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_take(n_0_box.IntVal, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_take
}

var cache_Data_Array_singleton gopurs_runtime.Value
var once_Data_Array_singleton sync.Once

func Get_Data_Array_singleton() gopurs_runtime.Value {
	once_Data_Array_singleton.Do(func() {
		cache_Data_Array_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_singleton(a_0_box))
		})
	})
	return cache_Data_Array_singleton
}

var cache_Data_Array_scanr gopurs_runtime.Value
var once_Data_Array_scanr sync.Once

func Get_Data_Array_scanr() gopurs_runtime.Value {
	once_Data_Array_scanr.Do(func() {
		cache_Data_Array_scanr = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_scanr(__local_var_0_box, __local_var_1_box, *((*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_scanr
}

var cache_Data_Array_scanl gopurs_runtime.Value
var once_Data_Array_scanl sync.Once

func Get_Data_Array_scanl() gopurs_runtime.Value {
	once_Data_Array_scanl.Do(func() {
		cache_Data_Array_scanl = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_scanl(__local_var_0_box, __local_var_1_box, *((*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_scanl
}

var cache_Data_Array_replicate gopurs_runtime.Value
var once_Data_Array_replicate sync.Once

func Get_Data_Array_replicate() gopurs_runtime.Value {
	once_Data_Array_replicate.Do(func() {
		cache_Data_Array_replicate = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_replicate(__local_var_0_box.IntVal, __local_var_1_box))
		})
	})
	return cache_Data_Array_replicate
}

var cache_Data_Array_go__range gopurs_runtime.Value
var once_Data_Array_go__range sync.Once

func Get_Data_Array_go__range() gopurs_runtime.Value {
	once_Data_Array_go__range.Do(func() {
		cache_Data_Array_go__range = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Array_go__range(__local_var_0_box.IntVal, __local_var_1_box.IntVal)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_Array_go__range
}

var cache_Data_Array_partition gopurs_runtime.Value
var once_Data_Array_partition sync.Once

func Get_Data_Array_partition() gopurs_runtime.Value {
	once_Data_Array_partition.Do(func() {
		cache_Data_Array_partition = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_Array_partition(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)))
				_ = orig
				return gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Array(orig.no), gopurs_runtime.Array(orig.yes))
			}()
		})
	})
	return cache_Data_Array_partition
}

var cache_Data_Array_null gopurs_runtime.Value
var once_Data_Array_null sync.Once

func Get_Data_Array_null() gopurs_runtime.Value {
	once_Data_Array_null.Do(func() {
		cache_Data_Array_null = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Array_null(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_null
}

var cache_Data_Array_modifyAtIndices gopurs_runtime.Value
var once_Data_Array_modifyAtIndices sync.Once

func Get_Data_Array_modifyAtIndices() gopurs_runtime.Value {
	once_Data_Array_modifyAtIndices.Do(func() {
		cache_Data_Array_modifyAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_modifyAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_Array_modifyAtIndices
}

var cache_Data_Array_mapWithIndex gopurs_runtime.Value
var once_Data_Array_mapWithIndex sync.Once

func Get_Data_Array_mapWithIndex() gopurs_runtime.Value {
	once_Data_Array_mapWithIndex.Do(func() {
		cache_Data_Array_mapWithIndex = Get_Data_FunctorWithIndex_mapWithIndexArray()
	})
	return cache_Data_Array_mapWithIndex
}

var cache_Data_Array_intersperse gopurs_runtime.Value
var once_Data_Array_intersperse sync.Once

func Get_Data_Array_intersperse() gopurs_runtime.Value {
	once_Data_Array_intersperse.Do(func() {
		cache_Data_Array_intersperse = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_intersperse(a_0_box, *((*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_intersperse
}

var cache_Data_Array_intercalate gopurs_runtime.Value
var once_Data_Array_intercalate sync.Once

func Get_Data_Array_intercalate() gopurs_runtime.Value {
	once_Data_Array_intercalate.Do(func() {
		cache_Data_Array_intercalate = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Data_Array_intercalate
}

var cache_Data_Array_insertAt gopurs_runtime.Value
var once_Data_Array_insertAt sync.Once

func Get_Data_Array_insertAt() gopurs_runtime.Value {
	once_Data_Array_insertAt.Do(func() {
		cache_Data_Array_insertAt = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_insertAt(__local_var_0_box.IntVal, __local_var_1_box, *((*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_insertAt
}

var cache_Data_Array_go__init gopurs_runtime.Value
var once_Data_Array_go__init sync.Once

func Get_Data_Array_go__init() gopurs_runtime.Value {
	once_Data_Array_go__init.Do(func() {
		cache_Data_Array_go__init = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_go__init(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_go__init
}

var cache_Data_Array_index gopurs_runtime.Value
var once_Data_Array_index sync.Once

func Get_Data_Array_index() gopurs_runtime.Value {
	once_Data_Array_index.Do(func() {
		cache_Data_Array_index = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_index(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)), __local_var_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_index
}

var cache_Data_Array_index__2519689635 gopurs_runtime.Value
var once_Data_Array_index__2519689635 sync.Once

func Get_Data_Array_index__2519689635() gopurs_runtime.Value {
	once_Data_Array_index__2519689635.Do(func() {
		cache_Data_Array_index__2519689635 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_index__2519689635(func() []*Constructor_Data_Tuple_Tuple[uint32, float64] {
					arr := *(*[]gopurs_runtime.Value)(__eta_norm_1_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[uint32, float64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Data_Array_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}(), __eta_norm_0_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_index__2519689635
}

var cache_Data_Array_index__1555618915 gopurs_runtime.Value
var once_Data_Array_index__1555618915 sync.Once

func Get_Data_Array_index__1555618915() gopurs_runtime.Value {
	once_Data_Array_index__1555618915.Do(func() {
		cache_Data_Array_index__1555618915 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_index__1555618915(*((*[]gopurs_runtime.Value)(__eta_norm_1_0_box.UnsafePtr)), __eta_norm_0_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_index__1555618915
}

var cache_Data_Array_index__3792059139 gopurs_runtime.Value
var once_Data_Array_index__3792059139 sync.Once

func Get_Data_Array_index__3792059139() gopurs_runtime.Value {
	once_Data_Array_index__3792059139.Do(func() {
		cache_Data_Array_index__3792059139 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_index__3792059139(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__eta_norm_1_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = *((*[]gopurs_runtime.Value)(v.UnsafePtr))
					}
					return unboxed
				}(), __eta_norm_0_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_index__3792059139
}

var cache_Data_Array_last gopurs_runtime.Value
var once_Data_Array_last sync.Once

func Get_Data_Array_last() gopurs_runtime.Value {
	once_Data_Array_last.Do(func() {
		cache_Data_Array_last = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_last(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_last
}

var cache_Data_Array_unsnoc gopurs_runtime.Value
var once_Data_Array_unsnoc sync.Once

func Get_Data_Array_unsnoc() gopurs_runtime.Value {
	once_Data_Array_unsnoc.Do(func() {
		cache_Data_Array_unsnoc = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_unsnoc(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Array(orig.go__init), orig.last)
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_unsnoc
}

var cache_Data_Array_modifyAt gopurs_runtime.Value
var once_Data_Array_modifyAt sync.Once

func Get_Data_Array_modifyAt() gopurs_runtime.Value {
	once_Data_Array_modifyAt.Do(func() {
		cache_Data_Array_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_modifyAt(i_0_box.IntVal, f_1_box, *((*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_modifyAt
}

var cache_Data_Array_unzip gopurs_runtime.Value
var once_Data_Array_unzip sync.Once

func Get_Data_Array_unzip() gopurs_runtime.Value {
	once_Data_Array_unzip.Do(func() {
		cache_Data_Array_unzip = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_unzip(func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr {
						unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v)
					}
					return unboxed
				}())
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Array_unzip
}

var cache_Data_Array_head gopurs_runtime.Value
var once_Data_Array_head sync.Once

func Get_Data_Array_head() gopurs_runtime.Value {
	once_Data_Array_head.Do(func() {
		cache_Data_Array_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_head(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_head
}

var cache_Data_Array_head__3867246685 gopurs_runtime.Value
var once_Data_Array_head__3867246685 sync.Once

func Get_Data_Array_head__3867246685() gopurs_runtime.Value {
	once_Data_Array_head__3867246685.Do(func() {
		cache_Data_Array_head__3867246685 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_head__3867246685(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = *((*[]gopurs_runtime.Value)(v.UnsafePtr))
					}
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_head__3867246685
}

var cache_Data_Array_nubBy gopurs_runtime.Value
var once_Data_Array_nubBy sync.Once

func Get_Data_Array_nubBy() gopurs_runtime.Value {
	once_Data_Array_nubBy.Do(func() {
		cache_Data_Array_nubBy = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_nubBy(comp_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_nubBy
}

var cache_Data_Array_nub gopurs_runtime.Value
var once_Data_Array_nub sync.Once

func Get_Data_Array_nub() gopurs_runtime.Value {
	once_Data_Array_nub.Do(func() {
		cache_Data_Array_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_nub
}

var cache_Data_Array_groupBy gopurs_runtime.Value
var once_Data_Array_groupBy sync.Once

func Get_Data_Array_groupBy() gopurs_runtime.Value {
	once_Data_Array_groupBy.Do(func() {
		cache_Data_Array_groupBy = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Array_groupBy(op_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Array(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_Array_groupBy
}

var cache_Data_Array_groupAllBy gopurs_runtime.Value
var once_Data_Array_groupAllBy sync.Once

func Get_Data_Array_groupAllBy() gopurs_runtime.Value {
	once_Data_Array_groupAllBy.Do(func() {
		cache_Data_Array_groupAllBy = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_groupAllBy(cmp_0_box)
		})
	})
	return cache_Data_Array_groupAllBy
}

var cache_Data_Array_groupAll gopurs_runtime.Value
var once_Data_Array_groupAll sync.Once

func Get_Data_Array_groupAll() gopurs_runtime.Value {
	once_Data_Array_groupAll.Do(func() {
		cache_Data_Array_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_groupAll
}

var cache_Data_Array_group gopurs_runtime.Value
var once_Data_Array_group sync.Once

func Get_Data_Array_group() gopurs_runtime.Value {
	once_Data_Array_group.Do(func() {
		cache_Data_Array_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_group
}

var cache_Data_Array_fromFoldable gopurs_runtime.Value
var once_Data_Array_fromFoldable sync.Once

func Get_Data_Array_fromFoldable() gopurs_runtime.Value {
	once_Data_Array_fromFoldable.Do(func() {
		cache_Data_Array_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_Array_fromFoldable
}

var cache_Data_Array_foldr gopurs_runtime.Value
var once_Data_Array_foldr sync.Once

func Get_Data_Array_foldr() gopurs_runtime.Value {
	once_Data_Array_foldr.Do(func() {
		cache_Data_Array_foldr = Get_Data_Foldable_foldrArray()
	})
	return cache_Data_Array_foldr
}

var cache_Data_Array_foldl gopurs_runtime.Value
var once_Data_Array_foldl sync.Once

func Get_Data_Array_foldl() gopurs_runtime.Value {
	once_Data_Array_foldl.Do(func() {
		cache_Data_Array_foldl = Get_Data_Foldable_foldlArray()
	})
	return cache_Data_Array_foldl
}

var cache_Data_Array_transpose gopurs_runtime.Value
var once_Data_Array_transpose sync.Once

func Get_Data_Array_transpose() gopurs_runtime.Value {
	once_Data_Array_transpose.Do(func() {
		cache_Data_Array_transpose = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Array_transpose(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = *((*[]gopurs_runtime.Value)(v.UnsafePtr))
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Array(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_Array_transpose
}

var cache_Data_Array_foldRecM gopurs_runtime.Value
var once_Data_Array_foldRecM sync.Once

func Get_Data_Array_foldRecM() gopurs_runtime.Value {
	once_Data_Array_foldRecM.Do(func() {
		cache_Data_Array_foldRecM = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_foldRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
		})
	})
	return cache_Data_Array_foldRecM
}

var cache_Data_Array_foldMap gopurs_runtime.Value
var once_Data_Array_foldMap sync.Once

func Get_Data_Array_foldMap() gopurs_runtime.Value {
	once_Data_Array_foldMap.Do(func() {
		cache_Data_Array_foldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Data_Array_foldMap
}

var cache_Data_Array_foldM gopurs_runtime.Value
var once_Data_Array_foldM sync.Once

func Get_Data_Array_foldM() gopurs_runtime.Value {
	once_Data_Array_foldM.Do(func() {
		cache_Data_Array_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), f_1_box, b_2_box, *((*[]gopurs_runtime.Value)(__local_var_3_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_foldM
}

var cache_Data_Array_fold gopurs_runtime.Value
var once_Data_Array_fold sync.Once

func Get_Data_Array_fold() gopurs_runtime.Value {
	once_Data_Array_fold.Do(func() {
		cache_Data_Array_fold = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_fold(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Data_Array_fold
}

var cache_Data_Array_findMap gopurs_runtime.Value
var once_Data_Array_findMap sync.Once

func Get_Data_Array_findMap() gopurs_runtime.Value {
	once_Data_Array_findMap.Do(func() {
		cache_Data_Array_findMap = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_findMap(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_findMap
}

var cache_Data_Array_findLastIndex gopurs_runtime.Value
var once_Data_Array_findLastIndex sync.Once

func Get_Data_Array_findLastIndex() gopurs_runtime.Value {
	once_Data_Array_findLastIndex.Do(func() {
		cache_Data_Array_findLastIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_findLastIndex(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_findLastIndex
}

var cache_Data_Array_insertBy gopurs_runtime.Value
var once_Data_Array_insertBy sync.Once

func Get_Data_Array_insertBy() gopurs_runtime.Value {
	once_Data_Array_insertBy.Do(func() {
		cache_Data_Array_insertBy = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_insertBy(cmp_0_box, x_1_box, *((*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_insertBy
}

var cache_Data_Array_insert gopurs_runtime.Value
var once_Data_Array_insert sync.Once

func Get_Data_Array_insert() gopurs_runtime.Value {
	once_Data_Array_insert.Do(func() {
		cache_Data_Array_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_insert
}

var cache_Data_Array_findIndex gopurs_runtime.Value
var once_Data_Array_findIndex sync.Once

func Get_Data_Array_findIndex() gopurs_runtime.Value {
	once_Data_Array_findIndex.Do(func() {
		cache_Data_Array_findIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_findIndex(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_findIndex
}

var cache_Data_Array_span gopurs_runtime.Value
var once_Data_Array_span sync.Once

func Get_Data_Array_span() gopurs_runtime.Value {
	once_Data_Array_span.Do(func() {
		cache_Data_Array_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_Array_span(p_0_box, *((*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)))
				_ = orig
				return gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(orig.go__init), gopurs_runtime.Array(orig.rest))
			}()
		})
	})
	return cache_Data_Array_span
}

var cache_Data_Array_takeWhile gopurs_runtime.Value
var once_Data_Array_takeWhile sync.Once

func Get_Data_Array_takeWhile() gopurs_runtime.Value {
	once_Data_Array_takeWhile.Do(func() {
		cache_Data_Array_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_takeWhile(p_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_takeWhile
}

var cache_Data_Array_find gopurs_runtime.Value
var once_Data_Array_find sync.Once

func Get_Data_Array_find() gopurs_runtime.Value {
	once_Data_Array_find.Do(func() {
		cache_Data_Array_find = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_find(f_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_find
}

var cache_Data_Array_filter gopurs_runtime.Value
var once_Data_Array_filter sync.Once

func Get_Data_Array_filter() gopurs_runtime.Value {
	once_Data_Array_filter.Do(func() {
		cache_Data_Array_filter = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_filter(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_filter
}

var cache_Data_Array_intersectBy gopurs_runtime.Value
var once_Data_Array_intersectBy sync.Once

func Get_Data_Array_intersectBy() gopurs_runtime.Value {
	once_Data_Array_intersectBy.Do(func() {
		cache_Data_Array_intersectBy = gopurs_runtime.Func3(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_intersectBy(eq2_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_intersectBy
}

var cache_Data_Array_intersect gopurs_runtime.Value
var once_Data_Array_intersect sync.Once

func Get_Data_Array_intersect() gopurs_runtime.Value {
	once_Data_Array_intersect.Do(func() {
		cache_Data_Array_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_intersect
}

var cache_Data_Array_elemLastIndex gopurs_runtime.Value
var once_Data_Array_elemLastIndex sync.Once

func Get_Data_Array_elemLastIndex() gopurs_runtime.Value {
	once_Data_Array_elemLastIndex.Do(func() {
		cache_Data_Array_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_Array_elemLastIndex
}

var cache_Data_Array_elemIndex gopurs_runtime.Value
var once_Data_Array_elemIndex sync.Once

func Get_Data_Array_elemIndex() gopurs_runtime.Value {
	once_Data_Array_elemIndex.Do(func() {
		cache_Data_Array_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_Array_elemIndex
}

var cache_Data_Array_notElem gopurs_runtime.Value
var once_Data_Array_notElem sync.Once

func Get_Data_Array_notElem() gopurs_runtime.Value {
	once_Data_Array_notElem.Do(func() {
		cache_Data_Array_notElem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Array_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), a_1_box, *((*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_notElem
}

var cache_Data_Array_elem gopurs_runtime.Value
var once_Data_Array_elem sync.Once

func Get_Data_Array_elem() gopurs_runtime.Value {
	once_Data_Array_elem.Do(func() {
		cache_Data_Array_elem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Array_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), a_1_box, *((*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_elem
}

var cache_Data_Array_dropWhile gopurs_runtime.Value
var once_Data_Array_dropWhile sync.Once

func Get_Data_Array_dropWhile() gopurs_runtime.Value {
	once_Data_Array_dropWhile.Do(func() {
		cache_Data_Array_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_dropWhile(p_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_dropWhile
}

var cache_Data_Array_dropEnd gopurs_runtime.Value
var once_Data_Array_dropEnd sync.Once

func Get_Data_Array_dropEnd() gopurs_runtime.Value {
	once_Data_Array_dropEnd.Do(func() {
		cache_Data_Array_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_dropEnd(n_0_box.IntVal, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_dropEnd
}

var cache_Data_Array_drop gopurs_runtime.Value
var once_Data_Array_drop sync.Once

func Get_Data_Array_drop() gopurs_runtime.Value {
	once_Data_Array_drop.Do(func() {
		cache_Data_Array_drop = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_drop(n_0_box.IntVal, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_drop
}

var cache_Data_Array_takeEnd gopurs_runtime.Value
var once_Data_Array_takeEnd sync.Once

func Get_Data_Array_takeEnd() gopurs_runtime.Value {
	once_Data_Array_takeEnd.Do(func() {
		cache_Data_Array_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_takeEnd(n_0_box.IntVal, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_takeEnd
}

var cache_Data_Array_deleteAt gopurs_runtime.Value
var once_Data_Array_deleteAt sync.Once

func Get_Data_Array_deleteAt() gopurs_runtime.Value {
	once_Data_Array_deleteAt.Do(func() {
		cache_Data_Array_deleteAt = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_deleteAt(__local_var_0_box.IntVal, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_deleteAt
}

var cache_Data_Array_deleteBy gopurs_runtime.Value
var once_Data_Array_deleteBy sync.Once

func Get_Data_Array_deleteBy() gopurs_runtime.Value {
	once_Data_Array_deleteBy.Do(func() {
		cache_Data_Array_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_deleteBy(v_0_box, v1_1_box, *((*[]gopurs_runtime.Value)(v2_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_deleteBy
}

var cache_Data_Array_go__delete gopurs_runtime.Value
var once_Data_Array_go__delete sync.Once

func Get_Data_Array_go__delete() gopurs_runtime.Value {
	once_Data_Array_go__delete.Do(func() {
		cache_Data_Array_go__delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_go__delete
}

var cache_Data_Array_difference gopurs_runtime.Value
var once_Data_Array_difference sync.Once

func Get_Data_Array_difference() gopurs_runtime.Value {
	once_Data_Array_difference.Do(func() {
		cache_Data_Array_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_difference
}

var cache_Data_Array_cons gopurs_runtime.Value
var once_Data_Array_cons sync.Once

func Get_Data_Array_cons() gopurs_runtime.Value {
	once_Data_Array_cons.Do(func() {
		cache_Data_Array_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_cons(x_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_cons
}

var cache_Data_Array_cons__3391444928 gopurs_runtime.Value
var once_Data_Array_cons__3391444928 sync.Once

func Get_Data_Array_cons__3391444928() gopurs_runtime.Value {
	once_Data_Array_cons__3391444928.Do(func() {
		cache_Data_Array_cons__3391444928 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_cons__3391444928(x_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_cons__3391444928
}

var cache_Data_Array_some gopurs_runtime.Value
var once_Data_Array_some sync.Once

func Get_Data_Array_some() gopurs_runtime.Value {
	once_Data_Array_some.Do(func() {
		cache_Data_Array_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_1_box), v_2_box)
		})
	})
	return cache_Data_Array_some
}

var cache_Data_Array_many gopurs_runtime.Value
var once_Data_Array_many sync.Once

func Get_Data_Array_many() gopurs_runtime.Value {
	once_Data_Array_many.Do(func() {
		cache_Data_Array_many = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_1_box), v_2_box)
		})
	})
	return cache_Data_Array_many
}

var cache_Data_Array_concatMap gopurs_runtime.Value
var once_Data_Array_concatMap sync.Once

func Get_Data_Array_concatMap() gopurs_runtime.Value {
	once_Data_Array_concatMap.Do(func() {
		cache_Data_Array_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_concatMap(b_0_box, *((*[]gopurs_runtime.Value)(a_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_concatMap
}

var cache_Data_Array_mapMaybe gopurs_runtime.Value
var once_Data_Array_mapMaybe sync.Once

func Get_Data_Array_mapMaybe() gopurs_runtime.Value {
	once_Data_Array_mapMaybe.Do(func() {
		cache_Data_Array_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_mapMaybe(f_0_box)
		})
	})
	return cache_Data_Array_mapMaybe
}

var cache_Data_Array_filterA gopurs_runtime.Value
var once_Data_Array_filterA sync.Once

func Get_Data_Array_filterA() gopurs_runtime.Value {
	once_Data_Array_filterA.Do(func() {
		cache_Data_Array_filterA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_filterA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_Array_filterA
}

var cache_Data_Array_catMaybes gopurs_runtime.Value
var once_Data_Array_catMaybes sync.Once

func Get_Data_Array_catMaybes() gopurs_runtime.Value {
	once_Data_Array_catMaybes.Do(func() {
		cache_Data_Array_catMaybes = Call_Data_Array_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}))
	})
	return cache_Data_Array_catMaybes
}

var cache_Data_Array_any gopurs_runtime.Value
var once_Data_Array_any sync.Once

func Get_Data_Array_any() gopurs_runtime.Value {
	once_Data_Array_any.Do(func() {
		cache_Data_Array_any = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Array_any(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_any
}

var cache_Data_Array_nubByEq gopurs_runtime.Value
var once_Data_Array_nubByEq sync.Once

func Get_Data_Array_nubByEq() gopurs_runtime.Value {
	once_Data_Array_nubByEq.Do(func() {
		cache_Data_Array_nubByEq = gopurs_runtime.Func2(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_nubByEq(eq2_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_nubByEq
}

var cache_Data_Array_nubEq gopurs_runtime.Value
var once_Data_Array_nubEq sync.Once

func Get_Data_Array_nubEq() gopurs_runtime.Value {
	once_Data_Array_nubEq.Do(func() {
		cache_Data_Array_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_nubEq
}

var cache_Data_Array_unionBy gopurs_runtime.Value
var once_Data_Array_unionBy sync.Once

func Get_Data_Array_unionBy() gopurs_runtime.Value {
	once_Data_Array_unionBy.Do(func() {
		cache_Data_Array_unionBy = gopurs_runtime.Func3(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_unionBy(eq2_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_unionBy
}

var cache_Data_Array_union gopurs_runtime.Value
var once_Data_Array_union sync.Once

func Get_Data_Array_union() gopurs_runtime.Value {
	once_Data_Array_union.Do(func() {
		cache_Data_Array_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_union
}

var cache_Data_Array_alterAt gopurs_runtime.Value
var once_Data_Array_alterAt sync.Once

func Get_Data_Array_alterAt() gopurs_runtime.Value {
	once_Data_Array_alterAt.Do(func() {
		cache_Data_Array_alterAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_alterAt(i_0_box.IntVal, f_1_box, *((*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_alterAt
}

var cache_Data_Array_all gopurs_runtime.Value
var once_Data_Array_all sync.Once

func Get_Data_Array_all() gopurs_runtime.Value {
	once_Data_Array_all.Do(func() {
		cache_Data_Array_all = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Array_all(__local_var_0_box, *((*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_all
}

var cache_Data_Array_all__3951006479 gopurs_runtime.Value
var once_Data_Array_all__3951006479 sync.Once

func Get_Data_Array_all__3951006479() gopurs_runtime.Value {
	once_Data_Array_all__3951006479.Do(func() {
		cache_Data_Array_all__3951006479 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Array_all__3951006479(__eta_norm_1_0_box, func() []string {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}()))
		})
	})
	return cache_Data_Array_all__3951006479
}

func Call_Data_Array_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Data_Array_void(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Data_Array_when(v_0_loop bool, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 bool = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 gopurs_runtime.Value
	{
		if v_0 {
			__t0 = v1_1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Unit_unit()
		})
	}
end_branch_0:
	return __t0
}

func Call_Data_Array_void3(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Data_Array_fromJust1(v_0_loop *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]) []gopurs_runtime.Value {
	var v_0 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 []gopurs_runtime.Value
	{
		if v_0 != nil {
			__t0 = (v_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() []gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Array_fromJust2(v_0_loop *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]) []gopurs_runtime.Value {
	var v_0 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 []gopurs_runtime.Value
	{
		if v_0 != nil {
			__t0 = (v_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() []gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Array_append__3388884614(__eta_norm_1_0_loop []gopurs_runtime.Value, __eta_norm_0_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
append__3388884614:
	for {
		if false {
			continue append__3388884614
		}
		var __eta_norm_1_0 []gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 []gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(__eta_norm_1_0), gopurs_runtime.Array(__eta_norm_0_1)).UnsafePtr))
	}
}

func Call_Data_Array_when1(v_0_loop bool, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 bool = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 gopurs_runtime.Value
	{
		if v_0 {
			__t0 = v1_1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Unit_unit()
		})
	}
end_branch_0:
	return __t0
}

func Call_Data_Array_zipWith(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
	_ = __local_var_2
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr))
}

func Call_Data_Array_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): sequence1_1_0 shape=App(Var) bindingType=(Func [(Array (TypeApp (TypeVar m$scope17) [(TypeVar c$scope20)]))] (TypeApp (TypeVar m$scope17) [(Array (TypeVar c$scope20))]))
	sequence1_1_0 := Call_Data_Traversable_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()), dictApplicative_0)
	_ = sequence1_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), f_2, xs_3, ys_4).UnsafePtr))))
	})
}

func Call_Data_Array_updateAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): traverse_2_1_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a$scope30)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope33), Boolean])), (TypeApp (TypeVar t$scope29) [(ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a$scope30)])])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope33), Unit]))
	traverse_2_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST())), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)})
	_ = traverse_2_1_0
	return gopurs_runtime.Func2(func(us_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			result_4_1 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), xs_3)
			_ = result_4_1
			_dollar___unused_5_2 := gopurs_runtime.Apply3(traverse_2_1_0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
					__local_var_6_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
					_ = __local_var_6_3
					// TAST (Let): __local_var_7_4 shape=Other bindingType=Any
					__local_var_7_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
					_ = __local_var_7_4
					return gopurs_runtime.UncurriedApp3(Get_Data_Array_ST_pokeImpl(), __local_var_7_4, __local_var_6_3, result_4_1)
				})
			}), us_2, gopurs_runtime.Value{})
			_ = _dollar___unused_5_2
			return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_4_1)
		}))
	})
}

func Call_Data_Array_updateAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 int64 = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
	_ = __local_var_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_unsafeIndex(_dollar___unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	var __local_var_2 int64 = __local_var_2_loop
	_ = __local_var_2
	arrayUnsafe_value_0 := __local_var_1
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := __local_var_2
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_unsafeIndex1(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) gopurs_runtime.Value {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	arrayUnsafe_value_0 := __local_var_0
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := __local_var_1
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_unsafeIndex2(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) gopurs_runtime.Value {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	arrayUnsafe_value_0 := __local_var_0
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := __local_var_1
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_unsafeIndex3(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) gopurs_runtime.Value {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	arrayUnsafe_value_0 := __local_var_0
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := __local_var_1
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_unsafeIndex4(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) gopurs_runtime.Value {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	arrayUnsafe_value_0 := __local_var_0
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := __local_var_1
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_uncons(__local_var_0_loop []gopurs_runtime.Value) struct {
	V0 struct {
		head gopurs_runtime.Value
		tail []gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V1 bool
} {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return func() struct {
		V0 struct {
			head gopurs_runtime.Value
			tail []gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}
		V1 bool
	} {
		_v := gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.RecordDict2("head", "tail", x_1, xs_2), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), gopurs_runtime.Array(__local_var_0))
		if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 struct {
					head gopurs_runtime.Value
					tail []gopurs_runtime.Value
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V1 bool
			}{V0: func() struct {
				head gopurs_runtime.Value
				tail []gopurs_runtime.Value
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			} {
				orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					head gopurs_runtime.Value
					tail []gopurs_runtime.Value
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}
				clone.head = gopurs_runtime.RecordGet(orig, "head")
				clone.tail = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "tail").UnsafePtr))
				return clone
			}(), V1: true}
		}
		return struct {
			V0 struct {
				head gopurs_runtime.Value
				tail []gopurs_runtime.Value
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V1 bool
		}{V0: struct {
			head gopurs_runtime.Value
			tail []gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, V1: false}
	}()
}

func Call_Data_Array_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	// TAST (Let): len_2_0 shape=Other bindingType=Int
	len_2_0 := gopurs_runtime.Int(int64(len(xs_1))).IntVal
	_ = len_2_0
	return gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 gopurs_runtime.Value
		{
			if (i_3.IntVal) < (len_2_0) {
				arrayUnsafe_value_1 := xs_1
				_ = arrayUnsafe_value_1
				arrayUnsafe_source_1 := arrayUnsafe_value_1
				_ = arrayUnsafe_source_1
				arrayUnsafe_index_2 := i_3.IntVal
				_ = arrayUnsafe_index_2
				__t3 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (arrayUnsafe_source_1)[arrayUnsafe_index_2], gopurs_runtime.Int(gopurs_runtime.IntAdd(i_3.IntVal, int64(1)))}))}, true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		}
	end_branch_3:
		return __t3
	}), gopurs_runtime.Int(int64(0)))
}

func Call_Data_Array_tail(__local_var_0_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{xs_2, true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	}), gopurs_runtime.Array(__local_var_0)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_sortBy(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var comp_0 gopurs_runtime.Value = comp_0_loop
	_ = comp_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 gopurs_runtime.Value
		{
			var __t_tag_0 uint32 = uint32(v_2.IntVal)
			_ = __t_tag_0
			if uint32(__t_tag_0) == 380165415 {
				__t3 = gopurs_runtime.Int(int64(1))
				goto end_branch_3
			} else {

			}
		}
		{
			var __t_tag_1 uint32 = uint32(v_2.IntVal)
			_ = __t_tag_1
			if uint32(__t_tag_1) == 902936544 {
				__t3 = gopurs_runtime.Int(int64(0))
				goto end_branch_3
			} else {

			}
		}
		{
			var __t_tag_2 uint32 = uint32(v_2.IntVal)
			_ = __t_tag_2
			if uint32(__t_tag_2) == 1527465420 {
				__t3 = gopurs_runtime.Int(int64(-1))
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
		}
	end_branch_3:
		return __t3
	}), gopurs_runtime.Array(__local_var_1)).UnsafePtr))
}

func Call_Data_Array_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(Get_Data_Array_sortBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictOrd_0.V1, gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
	}))
}

func Call_Data_Array_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): compare_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope96), (TypeVar a$scope96)] (ADT ["Data","Ordering","Ordering"] []))
	compare_1_0 := Call_Data_Ord_compare(dictOrd_0)
	_ = compare_1_0
	return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Array(Call_Data_Array_sortBy(compare_1_0, *((*[]gopurs_runtime.Value)(xs_2.UnsafePtr))))
	})
}

func Call_Data_Array_snoc(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Array_ST_push(), x_1)
		_ = __local_var_2_0
		result_3_1 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), gopurs_runtime.Array(xs_0))
		_ = result_3_1
		_dollar___unused_4_2 := gopurs_runtime.Apply2(__local_var_2_0, result_3_1, gopurs_runtime.Value{})
		_ = _dollar___unused_4_2
		return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_3_1)
	})).UnsafePtr))
}

func Call_Data_Array_slice(__local_var_0_loop int64, __local_var_1_loop int64, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var __local_var_0 int64 = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
	_ = __local_var_2
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr))
}

func Call_Data_Array_splitAt(v_0_loop int64, v1_1_loop []gopurs_runtime.Value) struct {
	after  []gopurs_runtime.Value
	before []gopurs_runtime.Value
	/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
} {
	var v_0 int64 = v_0_loop
	_ = v_0
	var v1_1 []gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var __t0 struct {
		after  []gopurs_runtime.Value
		before []gopurs_runtime.Value
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	}
	{
		if (v_0) <= (int64(0)) {
			__t0 = struct {
				after  []gopurs_runtime.Value
				before []gopurs_runtime.Value
				/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
			}{v1_1, *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr))}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = struct {
			after  []gopurs_runtime.Value
			before []gopurs_runtime.Value
			/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
		}{*((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(int64(len(v1_1))), gopurs_runtime.Array(v1_1)).UnsafePtr)), *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(v_0), gopurs_runtime.Array(v1_1)).UnsafePtr))}
	}
end_branch_0:
	return __t0
}

func Call_Data_Array_take(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	var __t0 []gopurs_runtime.Value
	{
		if (n_0) < (int64(1)) {
			__t0 = *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(n_0), gopurs_runtime.Array(xs_1)).UnsafePtr))
	}
end_branch_0:
	return __t0
}

func Call_Data_Array_singleton(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr))
}

func Call_Data_Array_scanr(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
	_ = __local_var_2
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanrImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr))
}

func Call_Data_Array_scanl(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
	_ = __local_var_2
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanlImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr))
}

func Call_Data_Array_replicate(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var __local_var_0 int64 = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Int(__local_var_0), __local_var_1).UnsafePtr))
}

func Call_Data_Array_go__range(__local_var_0_loop int64, __local_var_1_loop int64) []int64 {
	var __local_var_0 int64 = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1)).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Data_Array_partition(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct {
	no  []gopurs_runtime.Value
	yes []gopurs_runtime.Value
	/*gopurs:record-v1:2:006e006f,007900650073*/
} {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return func() struct {
		no  []gopurs_runtime.Value
		yes []gopurs_runtime.Value
		/*gopurs:record-v1:2:006e006f,007900650073*/
	} {
		orig := gopurs_runtime.UncurriedApp2(Get_Data_Array_partitionImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1))
		_ = orig
		clone := struct {
			no  []gopurs_runtime.Value
			yes []gopurs_runtime.Value
			/*gopurs:record-v1:2:006e006f,007900650073*/
		}{}
		clone.no = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "no").UnsafePtr))
		clone.yes = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "yes").UnsafePtr))
		return clone
	}()
}

func Call_Data_Array_null(xs_0_loop []gopurs_runtime.Value) bool {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	return (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (int64(0))
}

func Call_Data_Array_modifyAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): traverse_2_1_0 shape=App(Var) bindingType=(Func [(Func [Int] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope160), Boolean])), (TypeApp (TypeVar t$scope156) [Int])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope160), Unit]))
	traverse_2_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST())), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)})
	_ = traverse_2_1_0
	return gopurs_runtime.Func3(func(is_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			result_5_1 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_thawImpl(), xs_4)
			_ = result_5_1
			_dollar___unused_6_2 := gopurs_runtime.Apply3(traverse_2_1_0, gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Data_Array_ST_modify(i_6.IntVal, f_3, result_5_1)
			}), is_2, gopurs_runtime.Value{})
			_ = _dollar___unused_6_2
			return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_5_1)
		}))
	})
}

func Call_Data_Array_intersperse(a_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var arr_1 []gopurs_runtime.Value = arr_1_loop
	_ = arr_1
	// TAST (Let): v_2_0 shape=Other bindingType=Int
	v_2_0 := gopurs_runtime.Int(int64(len(arr_1))).IntVal
	_ = v_2_0
	var __t11 []gopurs_runtime.Value
	{
		if (v_2_0) < (int64(2)) {
			__t11 = arr_1
			goto end_branch_11
		} else {

		}
	}
	{
		__t11 = *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			out_3_1 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
			_ = out_3_1
			arrayUnsafe_value_4 := arr_1
			_ = arrayUnsafe_value_4
			arrayUnsafe_source_4 := arrayUnsafe_value_4
			_ = arrayUnsafe_source_4
			arrayUnsafe_index_5 := int64(0)
			_ = arrayUnsafe_index_5
			_dollar___unused_4_3 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), (arrayUnsafe_source_4)[arrayUnsafe_index_5], out_3_1)
			_ = _dollar___unused_4_3
			_dollar___unused_5_6 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_forImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(v_2_0), gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					_dollar___unused_6_7 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), a_0, out_3_1)
					_ = _dollar___unused_6_7
					arrayUnsafe_value_9 := arr_1
					_ = arrayUnsafe_value_9
					arrayUnsafe_source_9 := arrayUnsafe_value_9
					_ = arrayUnsafe_source_9
					arrayUnsafe_index_10 := idx_5.IntVal
					_ = arrayUnsafe_index_10
					__local_var_7_8 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), (arrayUnsafe_source_9)[arrayUnsafe_index_10], out_3_1)
					_ = __local_var_7_8
					return Get_Data_Unit_unit()
				})
			})), gopurs_runtime.Value{})
			_ = _dollar___unused_5_6
			__local_var_4_2 := out_3_1
			_ = __local_var_4_2
			return gopurs_runtime.Apply2(Get_Data_Array_ST_unsafeFreeze(), __local_var_4_2, gopurs_runtime.Value{})
		})).UnsafePtr))
	}
end_branch_11:
	return __t11
}

func Call_Data_Array_intercalate(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	return Call_Data_Foldable_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), dictMonoid_0)
}

func Call_Data_Array_insertAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 int64 = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
	_ = __local_var_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_go__init(xs_0_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var __t0 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
	{
		if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (int64(0)) {
			__t0 = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
		__t0 = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_0))).IntVal, int64(1))), gopurs_runtime.Array(xs_0)), true}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_index(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	arrayIndex_arg_0 := Get_Data_Maybe_Just()
	_ = arrayIndex_arg_0
	arrayIndex_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}())
	_ = arrayIndex_arg_1
	arrayIndex_arg_2 := __local_var_0
	_ = arrayIndex_arg_2
	arrayIndex_arg_3 := __local_var_1
	_ = arrayIndex_arg_3
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := func() gopurs_runtime.Value {
		arrayIndex_value := arrayIndex_arg_2
		_ = arrayIndex_value
		arrayIndex_source_4 := arrayIndex_value
		_ = arrayIndex_source_4
		if ((arrayIndex_arg_3) >= (0)) && ((arrayIndex_arg_3) < (int64(len(arrayIndex_source_4)))) {
			return gopurs_runtime.Apply(arrayIndex_arg_0, (arrayIndex_source_4)[arrayIndex_arg_3])
		} else {

		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_1)}
	}(); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_index__2519689635(__eta_norm_1_0_loop []*Constructor_Data_Tuple_Tuple[uint32, float64], __eta_norm_0_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
index__2519689635:
	for {
		if false {
			continue index__2519689635
		}
		var __eta_norm_1_0 []*Constructor_Data_Tuple_Tuple[uint32, float64] = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		arrayIndex_arg_0 := Get_Data_Maybe_Just()
		_ = arrayIndex_arg_0
		arrayIndex_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
		_ = arrayIndex_arg_1
		arrayIndex_arg_2 := __eta_norm_1_0
		_ = arrayIndex_arg_2
		arrayIndex_arg_3 := __eta_norm_0_1
		_ = arrayIndex_arg_3
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2304530460_3094389156(Rebox_Data_Array_3094389156_2304530460(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			arrayIndex_value := arrayIndex_arg_2
			_ = arrayIndex_value
			arrayIndex_source_4 := arrayIndex_value
			_ = arrayIndex_source_4
			if ((arrayIndex_arg_3) >= (0)) && ((arrayIndex_arg_3) < (int64(len(arrayIndex_source_4)))) {
				return gopurs_runtime.Apply(arrayIndex_arg_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_3132786365_138441832((arrayIndex_source_4)[arrayIndex_arg_3]))})
			} else {

			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_1)}
		}()))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Array_index__1555618915(__eta_norm_1_0_loop []gopurs_runtime.Value, __eta_norm_0_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
index__1555618915:
	for {
		if false {
			continue index__1555618915
		}
		var __eta_norm_1_0 []gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		arrayIndex_arg_0 := Get_Data_Maybe_Just()
		_ = arrayIndex_arg_0
		arrayIndex_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
		_ = arrayIndex_arg_1
		arrayIndex_arg_2 := __eta_norm_1_0
		_ = arrayIndex_arg_2
		arrayIndex_arg_3 := __eta_norm_0_1
		_ = arrayIndex_arg_3
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			arrayIndex_value := arrayIndex_arg_2
			_ = arrayIndex_value
			arrayIndex_source_4 := arrayIndex_value
			_ = arrayIndex_source_4
			if ((arrayIndex_arg_3) >= (0)) && ((arrayIndex_arg_3) < (int64(len(arrayIndex_source_4)))) {
				return gopurs_runtime.Apply(arrayIndex_arg_0, (arrayIndex_source_4)[arrayIndex_arg_3])
			} else {

			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_1)}
		}()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Array_index__3792059139(__eta_norm_1_0_loop [][]gopurs_runtime.Value, __eta_norm_0_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
index__3792059139:
	for {
		if false {
			continue index__3792059139
		}
		var __eta_norm_1_0 [][]gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		arrayIndex_arg_0 := Get_Data_Maybe_Just()
		_ = arrayIndex_arg_0
		arrayIndex_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
		_ = arrayIndex_arg_1
		arrayIndex_arg_2 := __eta_norm_1_0
		_ = arrayIndex_arg_2
		arrayIndex_arg_3 := __eta_norm_0_1
		_ = arrayIndex_arg_3
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			arrayIndex_value := arrayIndex_arg_2
			_ = arrayIndex_value
			arrayIndex_source_4 := arrayIndex_value
			_ = arrayIndex_source_4
			if ((arrayIndex_arg_3) >= (0)) && ((arrayIndex_arg_3) < (int64(len(arrayIndex_source_4)))) {
				return gopurs_runtime.Apply(arrayIndex_arg_0, gopurs_runtime.Array((arrayIndex_source_4)[arrayIndex_arg_3]))
			} else {

			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_1)}
		}()))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Array_last(xs_0_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	arrayIndex_arg_0 := Get_Data_Maybe_Just()
	_ = arrayIndex_arg_0
	arrayIndex_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}())
	_ = arrayIndex_arg_1
	arrayIndex_arg_2 := xs_0
	_ = arrayIndex_arg_2
	arrayIndex_arg_3 := gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_0))).IntVal, int64(1))
	_ = arrayIndex_arg_3
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		arrayIndex_value := arrayIndex_arg_2
		_ = arrayIndex_value
		arrayIndex_source_4 := arrayIndex_value
		_ = arrayIndex_source_4
		if ((arrayIndex_arg_3) >= (0)) && ((arrayIndex_arg_3) < (int64(len(arrayIndex_source_4)))) {
			return gopurs_runtime.Apply(arrayIndex_arg_0, (arrayIndex_source_4)[arrayIndex_arg_3])
		} else {

		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_1)}
	}()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_unsnoc(xs_0_loop []gopurs_runtime.Value) struct {
	V0 struct {
		go__init []gopurs_runtime.Value
		last     gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
	}
	V1 bool
} {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (int64(0)) {
			__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_6
		} else {

		}
	}
	{
		arrayIndex_arg_1 := Get_Data_Maybe_Just()
		_ = arrayIndex_arg_1
		arrayIndex_arg_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
		_ = arrayIndex_arg_2
		arrayIndex_arg_3 := xs_0
		_ = arrayIndex_arg_3
		arrayIndex_arg_4 := gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_0))).IntVal, int64(1))
		_ = arrayIndex_arg_4
		// TAST (Let): __local_var_1_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope209)])
		__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			arrayIndex_value := arrayIndex_arg_3
			_ = arrayIndex_value
			arrayIndex_source_5 := arrayIndex_value
			_ = arrayIndex_source_5
			if ((arrayIndex_arg_4) >= (0)) && ((arrayIndex_arg_4) < (int64(len(arrayIndex_source_5)))) {
				return gopurs_runtime.Apply(arrayIndex_arg_1, (arrayIndex_source_5)[arrayIndex_arg_4])
			} else {

			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_2)}
		}())
		_ = __local_var_1_0
		var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if __local_var_1_0 != nil {
				__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_0))).IntVal, int64(1))), gopurs_runtime.Array(xs_0)), (__local_var_1_0).V0), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				goto end_branch_5
			} else {

			}
		}
		{
			__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	end_branch_5:
		__t6 = __t5
	}
end_branch_6:
	return func() struct {
		V0 struct {
			go__init []gopurs_runtime.Value
			last     gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
		}
		V1 bool
	} {
		_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
		if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 struct {
					go__init []gopurs_runtime.Value
					last     gopurs_runtime.Value
					/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
				}
				V1 bool
			}{V0: func() struct {
				go__init []gopurs_runtime.Value
				last     gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
			} {
				orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					go__init []gopurs_runtime.Value
					last     gopurs_runtime.Value
					/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
				}{}
				clone.go__init = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "init").UnsafePtr))
				clone.last = gopurs_runtime.RecordGet(orig, "last")
				return clone
			}(), V1: true}
		}
		return struct {
			V0 struct {
				go__init []gopurs_runtime.Value
				last     gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
			}
			V1 bool
		}{V0: struct {
			go__init []gopurs_runtime.Value
			last     gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
		}{}, V1: false}
	}()
}

func Call_Data_Array_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var xs_2 []gopurs_runtime.Value = xs_2_loop
	_ = xs_2
	arrayIndex_arg_1 := Get_Data_Maybe_Just()
	_ = arrayIndex_arg_1
	arrayIndex_arg_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}())
	_ = arrayIndex_arg_2
	arrayIndex_arg_3 := xs_2
	_ = arrayIndex_arg_3
	arrayIndex_arg_4 := i_0
	_ = arrayIndex_arg_4
	// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope211)])
	__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		arrayIndex_value := arrayIndex_arg_3
		_ = arrayIndex_value
		arrayIndex_source_5 := arrayIndex_value
		_ = arrayIndex_source_5
		if ((arrayIndex_arg_4) >= (0)) && ((arrayIndex_arg_4) < (int64(len(arrayIndex_source_5)))) {
			return gopurs_runtime.Apply(arrayIndex_arg_1, (arrayIndex_source_5)[arrayIndex_arg_4])
		} else {

		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_2)}
	}())
	_ = __local_var_3_0
	var __t5 gopurs_runtime.Value
	{
		if __local_var_3_0 == nil {
			__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
			goto end_branch_5
		} else {

		}
	}
	{
		if __local_var_3_0 != nil {
			__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}, gopurs_runtime.Int(i_0), gopurs_runtime.Apply(f_1, (__local_var_3_0).V0), gopurs_runtime.Array(xs_2))))))}
			goto end_branch_5
		} else {

		}
	}
	{
		__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_5:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_unzip(xs_0_loop []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var xs_0 []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = xs_0_loop
	_ = xs_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1807945000_138441832(Rebox_Data_Array_138441832_1807945000(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		fsts_1_0 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
		_ = fsts_1_0
		snds_2_1 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
		_ = snds_2_1
		iter_3_2 := gopurs_runtime.Apply(Call_Data_Array_ST_Iterator_iterator(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			arrayIndex_arg_3 := Get_Data_Maybe_Just()
			_ = arrayIndex_arg_3
			arrayIndex_arg_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			_ = arrayIndex_arg_4
			arrayIndex_arg_5 := xs_0
			_ = arrayIndex_arg_5
			arrayIndex_arg_6 := v_3
			_ = arrayIndex_arg_6
			return func() gopurs_runtime.Value {
				arrayIndex_value := arrayIndex_arg_5
				_ = arrayIndex_value
				arrayIndex_source_7 := arrayIndex_value
				_ = arrayIndex_source_7
				if ((arrayIndex_arg_6.IntVal) >= (0)) && ((arrayIndex_arg_6.IntVal) < (int64(len(arrayIndex_source_7)))) {
					return gopurs_runtime.Apply(arrayIndex_arg_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((arrayIndex_source_7)[arrayIndex_arg_6.IntVal])})
				} else {

				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_4)}
			}()
		})), gopurs_runtime.Value{})
		_ = iter_3_2
		_dollar___unused_4_7 := gopurs_runtime.Apply(Call_Data_Array_ST_Iterator_iterate(Rebox_Data_Array_3988677819_394862806(Rebox_Data_Array_394862806_3988677819(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3_2))), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_5_8 shape=Other bindingType=Any
				__local_var_5_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
				_ = __local_var_5_8
				// TAST (Let): __local_var_6_9 shape=Other bindingType=Any
				__local_var_6_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
				_ = __local_var_6_9
				_dollar___unused_7_10 := gopurs_runtime.Apply2(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_5_8, fsts_1_0)
				}), gopurs_runtime.Value{})
				_ = _dollar___unused_7_10
				return gopurs_runtime.Apply2(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_6_9, snds_2_1)
				}), gopurs_runtime.Value{})
			})
		})), gopurs_runtime.Value{})
		_ = _dollar___unused_4_7
		fsts_prime__5_11 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), fsts_1_0)
		_ = fsts_prime__5_11
		snds_prime__6_12 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), snds_2_1)
		_ = snds_prime__6_12
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, fsts_prime__5_11, snds_prime__6_12}))}
	}))))))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Data_Array_head(xs_0_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	arrayIndex_arg_0 := Get_Data_Maybe_Just()
	_ = arrayIndex_arg_0
	arrayIndex_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}())
	_ = arrayIndex_arg_1
	arrayIndex_arg_2 := xs_0
	_ = arrayIndex_arg_2
	arrayIndex_arg_3 := int64(0)
	_ = arrayIndex_arg_3
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		arrayIndex_value := arrayIndex_arg_2
		_ = arrayIndex_value
		arrayIndex_source_4 := arrayIndex_value
		_ = arrayIndex_source_4
		if ((arrayIndex_arg_3) >= (0)) && ((arrayIndex_arg_3) < (int64(len(arrayIndex_source_4)))) {
			return gopurs_runtime.Apply(arrayIndex_arg_0, (arrayIndex_source_4)[arrayIndex_arg_3])
		} else {

		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_1)}
	}()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_head__3867246685(xs_0_loop [][]gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
head__3867246685:
	for {
		if false {
			continue head__3867246685
		}
		var xs_0 [][]gopurs_runtime.Value = xs_0_loop
		_ = xs_0
		arrayIndex_arg_0 := Get_Data_Maybe_Just()
		_ = arrayIndex_arg_0
		arrayIndex_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
		_ = arrayIndex_arg_1
		arrayIndex_arg_2 := xs_0
		_ = arrayIndex_arg_2
		arrayIndex_arg_3 := int64(0)
		_ = arrayIndex_arg_3
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			arrayIndex_value := arrayIndex_arg_2
			_ = arrayIndex_value
			arrayIndex_source_4 := arrayIndex_value
			_ = arrayIndex_source_4
			if ((arrayIndex_arg_3) >= (0)) && ((arrayIndex_arg_3) < (int64(len(arrayIndex_source_4)))) {
				return gopurs_runtime.Apply(arrayIndex_arg_0, gopurs_runtime.Array((arrayIndex_source_4)[arrayIndex_arg_3]))
			} else {

			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_1)}
		}()))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Array_nubBy(comp_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var comp_0 gopurs_runtime.Value = comp_0_loop
	_ = comp_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	// TAST (Let): indexedAndSorted_2_0 shape=App(Var) bindingType=(Array (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a$scope225)]))
	indexedAndSorted_2_0 := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_sortBy(gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(comp_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1)
		}), *((*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
			arr := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_FunctorWithIndex_mapWithIndexArray(), Get_Data_Tuple_Tuple(), gopurs_runtime.Array(xs_1)).UnsafePtr)
				unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Data_Array_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))}
			}
			return gopurs_runtime.Array(boxed)
		}().UnsafePtr)))).UnsafePtr)
		unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
		for i, v := range arr {
			unboxed[i] = Rebox_Data_Array_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
		}
		return unboxed
	}()
	_ = indexedAndSorted_2_0
	arrayIndex_arg_2 := Get_Data_Maybe_Just()
	_ = arrayIndex_arg_2
	arrayIndex_arg_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}())
	_ = arrayIndex_arg_3
	arrayIndex_arg_4 := indexedAndSorted_2_0
	_ = arrayIndex_arg_4
	arrayIndex_arg_5 := int64(0)
	_ = arrayIndex_arg_5
	// TAST (Let): v_3_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a$scope225)])])
	v_3_1 := Rebox_Data_Array_3094389156_1785233618(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		arrayIndex_value := arrayIndex_arg_4
		_ = arrayIndex_value
		arrayIndex_source_6 := arrayIndex_value
		_ = arrayIndex_source_6
		if ((arrayIndex_arg_5) >= (0)) && ((arrayIndex_arg_5) < (int64(len(arrayIndex_source_6)))) {
			return gopurs_runtime.Apply(arrayIndex_arg_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832((arrayIndex_source_6)[arrayIndex_arg_5]))})
		} else {

		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_3)}
	}()))
	_ = v_3_1
	var __t18 []gopurs_runtime.Value
	{
		if v_3_1 == nil {
			__t18 = *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr))
			goto end_branch_18
		} else {

		}
	}
	{
		if v_3_1 != nil {
			arrayIntrinsic_arg_6 := Get_Data_Tuple_snd()
			_ = arrayIntrinsic_arg_6
			_ = arrayIntrinsic_arg_6
			arrayIntrinsic_arg_17 := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Data_Array_sortWith(Rebox_Data_Array_3308271157_4177771502(Rebox_Data_Array_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))), Get_Data_Tuple_fst()), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): __local_var_4_7 shape=App(Var) bindingType=Any
							__local_var_4_7 := gopurs_runtime.Apply(Get_Data_Array_ST_unsafeThaw(), func() gopurs_runtime.Value {
								arr := []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]{(v_3_1).V0}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))}
								}
								return gopurs_runtime.Array(boxed)
							}())
							_ = __local_var_4_7
							result_5_8 := gopurs_runtime.Apply(__local_var_4_7, gopurs_runtime.Value{})
							_ = result_5_8
							_dollar___unused_6_9 := gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_foreach(), func() gopurs_runtime.Value {
								arr := indexedAndSorted_2_0
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))}
								}
								return gopurs_runtime.Array(boxed)
							}(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): __local_var_7_10 shape=Other bindingType=Any
									__local_var_7_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1
									_ = __local_var_7_10
									// TAST (Let): __local_var_8_11 shape=App(Var) bindingType=Any
									__local_var_8_11 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{}), Get_Data_Array_last()))
									_ = __local_var_8_11
									__local_var_9_13 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_5_8)
									_ = __local_var_9_13
									lst_9_12 := gopurs_runtime.Apply(__local_var_8_11, __local_var_9_13)
									_ = lst_9_12
									// TAST (Let): __local_var_10_14 shape=App(Var) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope230), Unit])
									__local_var_10_14 := gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(Rebox_Data_Array_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v1_6))))}, result_5_8)
									}))
									_ = __local_var_10_14
									var __t16 gopurs_runtime.Value
									{
										var __t_tag_15 uint32 = uint32(gopurs_runtime.Apply2(comp_0, lst_9_12, __local_var_7_10).IntVal)
										_ = __t_tag_15
										if (uint32(__t_tag_15) == 902936544) != (true) {
											__t16 = __local_var_10_14
											goto end_branch_16
										} else {

										}
									}
									{
										__t16 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											return Get_Data_Unit_unit()
										})
									}
								end_branch_16:
									return gopurs_runtime.Apply(__t16, gopurs_runtime.Value{})
								})
							}), gopurs_runtime.Value{})
							_ = _dollar___unused_6_9
							return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_5_8)
						})).UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Data_Array_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
				unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Data_Array_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
				}
				return unboxed
			}()
			_ = arrayIntrinsic_arg_17
			_ = arrayIntrinsic_arg_17
			__t18 = *((*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr_val_arrayMap3 := arrayIntrinsic_arg_17
				_ = arr_val_arrayMap3
				_ = arr_val_arrayMap3
				arr_go_arrayMap3 := arr_val_arrayMap3
				_ = arr_go_arrayMap3
				res_go_arrayMap3 := make([]gopurs_runtime.Value, len(arr_go_arrayMap3))
				_ = res_go_arrayMap3
				for i_arrayMap3, v_arrayMap3 := range arr_go_arrayMap3 {
					res_go_arrayMap3[i_arrayMap3] = Call_Data_Tuple_snd(Rebox_Data_Array_1728839155_138441832(v_arrayMap3))
				}
				return res_go_arrayMap3
			}()).UnsafePtr))
			goto end_branch_18
		} else {

		}
	}
	{
		__t18 = func() []gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_18:
	return __t18
}

func Call_Data_Array_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Array_nubBy(), Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_Array_groupBy(op_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
	var op_0 gopurs_runtime.Value = op_0_loop
	_ = op_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return func() [][]gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			result_2_0 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
			_ = result_2_0
			iter_3_1 := gopurs_runtime.Apply(Call_Data_Array_ST_Iterator_iterator(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				arrayIndex_arg_2 := Get_Data_Maybe_Just()
				_ = arrayIndex_arg_2
				arrayIndex_arg_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				_ = arrayIndex_arg_3
				arrayIndex_arg_4 := xs_1
				_ = arrayIndex_arg_4
				arrayIndex_arg_5 := v_3
				_ = arrayIndex_arg_5
				return func() gopurs_runtime.Value {
					arrayIndex_value := arrayIndex_arg_4
					_ = arrayIndex_value
					arrayIndex_source_6 := arrayIndex_value
					_ = arrayIndex_source_6
					if ((arrayIndex_arg_5.IntVal) >= (0)) && ((arrayIndex_arg_5.IntVal) < (int64(len(arrayIndex_source_6)))) {
						return gopurs_runtime.Apply(arrayIndex_arg_2, (arrayIndex_source_6)[arrayIndex_arg_5.IntVal])
					} else {

					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_3)}
				}()
			})), gopurs_runtime.Value{})
			_ = iter_3_1
			_dollar___unused_4_6 := gopurs_runtime.Apply(Call_Data_Array_ST_Iterator_iterate(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3_1), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					sub1_5_8 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
					_ = sub1_5_8
					_dollar___unused_6_9 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), x_4, sub1_5_8)
					_ = _dollar___unused_6_9
					_dollar___unused_7_10 := gopurs_runtime.Apply(Call_Data_Array_ST_Iterator_pushWhile(gopurs_runtime.Apply(op_0, x_4), gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3_1), sub1_5_8), gopurs_runtime.Value{})
					_ = _dollar___unused_7_10
					grp_8_11 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), sub1_5_8)
					_ = grp_8_11
					__local_var_5_7 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), grp_8_11, result_2_0)
					_ = __local_var_5_7
					return Get_Data_Unit_unit()
				})
			})), gopurs_runtime.Value{})
			_ = _dollar___unused_4_6
			return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2_0)
		})).UnsafePtr)
		unboxed := make([][]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = *((*[]gopurs_runtime.Value)(v.UnsafePtr))
		}
		return unboxed
	}()
}

func Call_Data_Array_groupAllBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var cmp_0 gopurs_runtime.Value = cmp_0_loop
	_ = cmp_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Array_groupBy(), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_2).IntVal)
		_ = __t_tag_0
		return gopurs_runtime.Bool((uint32(__t_tag_0) == 902936544))
	})), gopurs_runtime.Apply(Get_Data_Array_sortBy(), cmp_0))
}

func Call_Data_Array_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Array_groupAllBy(Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_Array_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): eq2_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope249), (TypeVar a$scope249)] Boolean)
	eq2_1_0 := Call_Data_Eq_eq(dictEq_0)
	_ = eq2_1_0
	return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			arr := Call_Data_Array_groupBy(eq2_1_0, *((*[]gopurs_runtime.Value)(xs_2.UnsafePtr)))
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Array(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
}

func Call_Data_Array_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope253), (TypeVar b$scope260)] (TypeVar b$scope260)), (TypeVar b$scope260), (TypeApp (TypeVar f$scope252) [(TypeVar a$scope253)])] (TypeVar b$scope260))
	__local_var_1_0 := Call_Data_Foldable_foldr(dictFoldable_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_0, __local_var_2)
	})
}

func Call_Data_Array_transpose(xs_0_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
	var xs_0 [][]gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var Call_local_Data_Array_go__go_1_0_0 func(int64, [][]gopurs_runtime.Value) [][]gopurs_runtime.Value
	_ = Call_local_Data_Array_go__go_1_0_0
	var go__go_1_0_0 gopurs_runtime.Value
	_ = go__go_1_0_0
	Call_local_Data_Array_go__go_1_0_0 = func(idx_2_loop int64, allArrays_3_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
	go__go_1_0_0:
		for {
			if false {
				continue go__go_1_0_0
			}
			var idx_2 int64 = idx_2_loop
			_ = idx_2
			var allArrays_3 [][]gopurs_runtime.Value = allArrays_3_loop
			_ = allArrays_3
			// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a$scope270))])
			v_4_1 := Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Array_foldl(), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
				arrayIndex_arg_3 := Get_Data_Maybe_Just()
				_ = arrayIndex_arg_3
				arrayIndex_arg_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				_ = arrayIndex_arg_4
				arrayIndex_arg_5 := nextArr_5
				_ = arrayIndex_arg_5
				arrayIndex_arg_6 := idx_2
				_ = arrayIndex_arg_6
				// TAST (Let): __local_var_6_2 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope270)])
				__local_var_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					arrayIndex_value := arrayIndex_arg_5
					_ = arrayIndex_value
					arrayIndex_source_7 := (*[]gopurs_runtime.Value)(arrayIndex_value.UnsafePtr)
					_ = arrayIndex_source_7
					if ((arrayIndex_arg_6) >= (0)) && ((arrayIndex_arg_6) < (int64(len(*arrayIndex_source_7)))) {
						return gopurs_runtime.Apply(arrayIndex_arg_3, (*arrayIndex_source_7)[arrayIndex_arg_6])
					} else {

					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_4)}
				}())
				_ = __local_var_6_2
				var __t8 gopurs_runtime.Value
				{
					if __local_var_6_2 == nil {
						__t8 = acc_4
						goto end_branch_8
					} else {

					}
				}
				{
					if __local_var_6_2 != nil {
						var __t7 gopurs_runtime.Value
						{
							if acc_4.Type == 9 && acc_4.IntVal == 930809136 && acc_4.UnsafePtr == nil {
								__t7 = gopurs_runtime.Array([]gopurs_runtime.Value{(__local_var_6_2).V0})
								goto end_branch_7
							} else {

							}
						}
						{
							if acc_4.Type == 9 && acc_4.IntVal == 930809136 && acc_4.UnsafePtr != nil {
								__t7 = gopurs_runtime.Array(Call_Data_Array_snoc(*((*[]gopurs_runtime.Value)((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(acc_4.UnsafePtr).V0.UnsafePtr)), (__local_var_6_2).V0))
								goto end_branch_7
							} else {

							}
						}
						{
							__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_7:
						__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, __t7}))}
						goto end_branch_8
					} else {

					}
				}
				{
					__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_8:
				return __t8
			}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}, func() gopurs_runtime.Value {
				arr := xs_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Array(v)
				}
				return gopurs_runtime.Array(boxed)
			}())))
			_ = v_4_1
			var __t9 [][]gopurs_runtime.Value
			{
				if v_4_1 == nil {
					__t9 = allArrays_3
					goto end_branch_9
				} else {

				}
			}
			{
				if v_4_1 != nil {
					idx_2_loop = gopurs_runtime.IntAdd(idx_2, int64(1))
					allArrays_3_loop = func() [][]gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc(*((*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr := allArrays_3
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Array(v)
							}
							return gopurs_runtime.Array(boxed)
						}().UnsafePtr)), gopurs_runtime.Array((v_4_1).V0))).UnsafePtr)
						unboxed := make([][]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = *((*[]gopurs_runtime.Value)(v.UnsafePtr))
						}
						return unboxed
					}()
					continue go__go_1_0_0
					__t9 = func() [][]gopurs_runtime.Value { panic("unreachable") }()
					goto end_branch_9
				} else {

				}
			}
			{
				__t9 = func() [][]gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_9:
			return __t9
		}
	}
	go__go_1_0_0 = gopurs_runtime.Func(func(idx_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(allArrays_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_local_Data_Array_go__go_1_0_0(idx_2_loop_val.IntVal, func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(allArrays_3_loop_val.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = *((*[]gopurs_runtime.Value)(v.UnsafePtr))
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Array(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return Call_local_Data_Array_go__go_1_0_0(int64(0), [][]gopurs_runtime.Value{})
}

func Call_Data_Array_foldRecM(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
	__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
	_ = __local_var_2_1
	return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, array_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if (gopurs_runtime.RecordGet(o_6, "b").IntVal) >= (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(array_5))).IntVal) {
					__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2206451173_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct {
						a gopurs_runtime.Value
						b int64
						/*gopurs:record-v1:2:0061,0062*/
					}, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_6, "a")})))})
					goto end_branch_4
				} else {

				}
			}
			{
				arrayUnsafe_value_2 := array_5
				_ = arrayUnsafe_value_2
				arrayUnsafe_source_2 := (*[]gopurs_runtime.Value)(arrayUnsafe_value_2.UnsafePtr)
				_ = arrayUnsafe_source_2
				arrayUnsafe_index_3 := gopurs_runtime.RecordGet(o_6, "b").IntVal
				_ = arrayUnsafe_index_3
				__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet(o_6, "a"), (*arrayUnsafe_source_2)[arrayUnsafe_index_3]), gopurs_runtime.Func(func(res_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2660862937_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct {
						a gopurs_runtime.Value
						b int64
						/*gopurs:record-v1:2:0061,0062*/
					}, gopurs_runtime.Value]{1, struct {
						a gopurs_runtime.Value
						b int64
						/*gopurs:record-v1:2:0061,0062*/
					}{res_prime__7, gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(o_6, "b").IntVal, int64(1))}})))})
				}))
			}
		end_branch_4:
			return __t4
		}), func() gopurs_runtime.Value {
			orig := struct {
				a gopurs_runtime.Value
				b gopurs_runtime.Value
				/*gopurs:record-v1:2:0061,0062*/
			}{b_4, gopurs_runtime.Int(int64(0))}
			_ = orig
			return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
		}())
	})
}

func Call_Data_Array_foldMap(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	return Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), dictMonoid_0)
}

func Call_Data_Array_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, __local_var_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
foldM:
	for {
		if false {
			continue foldM
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var b_2 gopurs_runtime.Value = b_2_loop
		_ = b_2
		var __local_var_3 []gopurs_runtime.Value = __local_var_3_loop
		_ = __local_var_3
		return gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), b_2)
		}), gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, as_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, a_4), gopurs_runtime.Func(func(b_prime__6 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Data_Array_foldM(dictMonad_0, f_1, b_prime__6, *((*[]gopurs_runtime.Value)(as_5.UnsafePtr)))
			}))
		}), gopurs_runtime.Array(__local_var_3))
	}
}

func Call_Data_Array_fold(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	return Call_Data_Foldable_fold(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), dictMonoid_0)
}

func Call_Data_Array_findMap(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, Get_Data_Maybe_isJust(), __local_var_0, gopurs_runtime.Array(__local_var_1)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_findLastIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, __local_var_0, gopurs_runtime.Array(__local_var_1)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var cmp_0 gopurs_runtime.Value = cmp_0_loop
	_ = cmp_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var ys_2 []gopurs_runtime.Value = ys_2_loop
	_ = ys_2
	// TAST (Let): __local_var_3_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	__local_var_3_1 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t_tag_2 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_3).IntVal)
		_ = __t_tag_2
		return gopurs_runtime.Bool((uint32(__t_tag_2) == 380165415))
	}), gopurs_runtime.Array(ys_2))))
	_ = __local_var_3_1
	var __t3 int64
	{
		if __local_var_3_1 == nil {
			__t3 = int64(0)
			goto end_branch_3
		} else {

		}
	}
	{
		if __local_var_3_1 != nil {
			__t3 = gopurs_runtime.IntAdd((__local_var_3_1).V0, int64(1))
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_3:
	// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a$scope336))])
	__local_var_3_0 := Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Int(__t3), x_1, gopurs_runtime.Array(ys_2))))
	_ = __local_var_3_0
	var __t4 []gopurs_runtime.Value
	{
		if __local_var_3_0 != nil {
			__t4 = (__local_var_3_0).V0
			goto end_branch_4
		} else {

		}
	}
	{
		__t4 = func() []gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_4:
	return __t4
}

func Call_Data_Array_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Array_insertBy(), Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_Array_findIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, __local_var_0, gopurs_runtime.Array(__local_var_1)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_span(p_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) struct {
	go__init []gopurs_runtime.Value
	rest     []gopurs_runtime.Value
	/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
} {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var arr_1 []gopurs_runtime.Value = arr_1_loop
	_ = arr_1
	// TAST (Let): breakIndex_2_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	breakIndex_2_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((gopurs_runtime.Apply(p_0, x_2).IntVal) != (0)) != (true))
	}), gopurs_runtime.Array(arr_1))))
	_ = breakIndex_2_0
	var __t2 struct {
		go__init []gopurs_runtime.Value
		rest     []gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
	}
	{
		if breakIndex_2_0 != nil {
			var __t1 struct {
				go__init []gopurs_runtime.Value
				rest     []gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			}
			{
				if ((breakIndex_2_0).V0) == (int64(0)) {
					__t1 = struct {
						go__init []gopurs_runtime.Value
						rest     []gopurs_runtime.Value
						/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
					}{*((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)), arr_1}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					go__init []gopurs_runtime.Value
					rest     []gopurs_runtime.Value
					/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
				}{*((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int((breakIndex_2_0).V0), gopurs_runtime.Array(arr_1)).UnsafePtr)), *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int((breakIndex_2_0).V0), gopurs_runtime.Int(int64(len(arr_1))), gopurs_runtime.Array(arr_1)).UnsafePtr))}
			}
		end_branch_1:
			__t2 = __t1
			goto end_branch_2
		} else {

		}
	}
	{
		if breakIndex_2_0 == nil {
			__t2 = struct {
				go__init []gopurs_runtime.Value
				rest     []gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			}{arr_1, *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr))}
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() struct {
			go__init []gopurs_runtime.Value
			rest     []gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
		} {
			panic("Failed pattern match")
		}()
	}
end_branch_2:
	return __t2
}

func Call_Data_Array_takeWhile(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return Call_Data_Array_span(p_0, xs_1).go__init
}

func Call_Data_Array_find(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	// TAST (Let): __local_var_2_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	__local_var_2_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, f_0, gopurs_runtime.Array(xs_1))))
	_ = __local_var_2_0
	var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_2_0 != nil {
			arrayUnsafe_value_1 := xs_1
			_ = arrayUnsafe_value_1
			arrayUnsafe_source_1 := arrayUnsafe_value_1
			_ = arrayUnsafe_source_1
			arrayUnsafe_index_2 := gopurs_runtime.Int((__local_var_2_0).V0).IntVal
			_ = arrayUnsafe_index_2
			__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(arrayUnsafe_source_1)[arrayUnsafe_index_2], true}
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
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_filter(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	arrayIntrinsic_arg_0 := __local_var_0
	_ = arrayIntrinsic_arg_0
	_ = arrayIntrinsic_arg_0
	arrayIntrinsic_arg_1 := __local_var_1
	_ = arrayIntrinsic_arg_1
	_ = arrayIntrinsic_arg_1
	return func() []gopurs_runtime.Value {
		arr_val_filterImpl0 := arrayIntrinsic_arg_1
		_ = arr_val_filterImpl0
		_ = arr_val_filterImpl0
		arrayIntrinsic_callback_2 := arrayIntrinsic_arg_0
		_ = arrayIntrinsic_callback_2
		arr_go_filterImpl0 := arr_val_filterImpl0
		_ = arr_go_filterImpl0
		res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
		_ = res_go_filterImpl0
		for _, v_filterImpl0 := range arr_go_filterImpl0 {
			if (gopurs_runtime.Apply(arrayIntrinsic_callback_2, v_filterImpl0).IntVal) != (0) {
				if (len(res_go_filterImpl0)) == (0) {
					res_go_filterImpl0 = make([]gopurs_runtime.Value, 0, min(256, len(arr_go_filterImpl0)))
				} else {

				}
				res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
			} else {

			}
		}
		return res_go_filterImpl0
	}()
}

func Call_Data_Array_intersectBy(eq2_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var eq2_0 gopurs_runtime.Value = eq2_0_loop
	_ = eq2_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	var ys_2 []gopurs_runtime.Value = ys_2_loop
	_ = ys_2
	arrayIntrinsic_arg_2 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
		__local_var_4_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}, gopurs_runtime.Apply(eq2_0, x_3), gopurs_runtime.Array(ys_2))))
		_ = __local_var_4_0
		var __t1 gopurs_runtime.Value
		{
			if __local_var_4_0 == nil {
				__t1 = gopurs_runtime.Bool(false)
				goto end_branch_1
			} else {

			}
		}
		{
			if __local_var_4_0 != nil {
				__t1 = gopurs_runtime.Bool(true)
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Bool(func() bool { panic("Failed pattern match") }())
		}
	end_branch_1:
		return __t1
	})
	_ = arrayIntrinsic_arg_2
	_ = arrayIntrinsic_arg_2
	arrayIntrinsic_arg_3 := xs_1
	_ = arrayIntrinsic_arg_3
	_ = arrayIntrinsic_arg_3
	return func() []gopurs_runtime.Value {
		arr_val_filterImpl0 := arrayIntrinsic_arg_3
		_ = arr_val_filterImpl0
		_ = arr_val_filterImpl0
		arrayIntrinsic_callback_4 := arrayIntrinsic_arg_2
		_ = arrayIntrinsic_callback_4
		arr_go_filterImpl0 := arr_val_filterImpl0
		_ = arr_go_filterImpl0
		res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
		_ = res_go_filterImpl0
		for _, v_filterImpl0 := range arr_go_filterImpl0 {
			if (gopurs_runtime.Apply(arrayIntrinsic_callback_4, v_filterImpl0).IntVal) != (0) {
				if (len(res_go_filterImpl0)) == (0) {
					res_go_filterImpl0 = make([]gopurs_runtime.Value, 0, min(256, len(arr_go_filterImpl0)))
				} else {

				}
				res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
			} else {

			}
		}
		return res_go_filterImpl0
	}()
}

func Call_Data_Array_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_intersectBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Data_Array_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1)
	}))
}

func Call_Data_Array_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Data_Array_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1)
	}))
}

func Call_Data_Array_notElem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var arr_2 []gopurs_runtime.Value = arr_2_loop
	_ = arr_2
	// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	__local_var_3_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_3, a_1)
	}), gopurs_runtime.Array(arr_2))))
	_ = __local_var_3_0
	var __t1 bool
	{
		if __local_var_3_0 == nil {
			__t1 = true
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0 != nil {
			__t1 = false
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() bool { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Data_Array_elem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var arr_2 []gopurs_runtime.Value = arr_2_loop
	_ = arr_2
	// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	__local_var_3_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_3, a_1)
	}), gopurs_runtime.Array(arr_2))))
	_ = __local_var_3_0
	var __t1 bool
	{
		if __local_var_3_0 == nil {
			__t1 = false
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0 != nil {
			__t1 = true
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() bool { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Data_Array_dropWhile(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return Call_Data_Array_span(p_0, xs_1).rest
}

func Call_Data_Array_dropEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	// TAST (Let): __local_var_2_0 shape=Other bindingType=Int
	__local_var_2_0 := gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_1))).IntVal, n_0)
	_ = __local_var_2_0
	var __t1 []gopurs_runtime.Value
	{
		if (__local_var_2_0) < (int64(1)) {
			__t1 = *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr))
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Array(xs_1)).UnsafePtr))
	}
end_branch_1:
	return __t1
}

func Call_Data_Array_drop(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	var __t0 []gopurs_runtime.Value
	{
		if (n_0) < (int64(1)) {
			__t0 = xs_1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Array(xs_1)).UnsafePtr))
	}
end_branch_0:
	return __t0
}

func Call_Data_Array_takeEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	// TAST (Let): __local_var_2_0 shape=Other bindingType=Int
	__local_var_2_0 := gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_1))).IntVal, n_0)
	_ = __local_var_2_0
	var __t1 []gopurs_runtime.Value
	{
		if (__local_var_2_0) < (int64(1)) {
			__t1 = xs_1
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Array(xs_1)).UnsafePtr))
	}
end_branch_1:
	return __t1
}

func Call_Data_Array_deleteAt(__local_var_0_loop int64, __local_var_1_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 int64 = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Int(__local_var_0), gopurs_runtime.Array(__local_var_1)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 []gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t5 []gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(v2_2))).IntVal) == (int64(0)) {
			__t5 = *((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr))
			goto end_branch_5
		} else {

		}
	}
	{
		// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
		__local_var_3_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Array(v2_2))))
		_ = __local_var_3_0
		var __t4 gopurs_runtime.Value
		{
			if __local_var_3_0 == nil {
				__t4 = gopurs_runtime.Array(v2_2)
				goto end_branch_4
			} else {

			}
		}
		{
			if __local_var_3_0 != nil {
				// TAST (Let): __local_var_4_1 shape=Other bindingType=Int
				__local_var_4_1 := (__local_var_3_0).V0
				_ = __local_var_4_1
				__t4 = gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_6_2 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a$scope404))])
					__local_var_6_2 := Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))}, gopurs_runtime.Int(__local_var_4_1), gopurs_runtime.Array(v2_2))))
					_ = __local_var_6_2
					var __t3 gopurs_runtime.Value
					{
						if __local_var_6_2 != nil {
							__t3 = gopurs_runtime.Array((__local_var_6_2).V0)
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = gopurs_runtime.Array(func() []gopurs_runtime.Value { panic("Failed pattern match") }())
					}
				end_branch_3:
					return __t3
				}))
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_4:
		__t5 = *((*[]gopurs_runtime.Value)(__t4.UnsafePtr))
	}
end_branch_5:
	return __t5
}

func Call_Data_Array_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_deleteBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Foldable_foldrArray(), gopurs_runtime.Apply(Get_Data_Array_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_Array_cons(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr))
}

func Call_Data_Array_cons__3391444928(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
cons__3391444928:
	for {
		if false {
			continue cons__3391444928
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		var xs_1 []gopurs_runtime.Value = xs_1_loop
		_ = xs_1
		return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr))
	}
}

func Call_Data_Array_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], dictLazy_1_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var dictLazy_1 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_1_loop
	_ = dictLazy_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Array_cons(), v_2), gopurs_runtime.Apply(dictLazy_1.V0, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Array_many(dictAlternative_0, dictLazy_1, v_2)
	})))
}

func Call_Data_Array_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], dictLazy_1_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var dictLazy_1 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_1_loop
	_ = dictLazy_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_Data_Array_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "pure"), gopurs_runtime.Array([]gopurs_runtime.Value{})))
}

func Call_Data_Array_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 []gopurs_runtime.Value = a_1_loop
	_ = a_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Bind_arrayBind(), gopurs_runtime.Array(a_1), b_0).UnsafePtr))
}

func Call_Data_Array_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(Get_Data_Array_concatMap(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if v2_1.Type == 9 && v2_1.IntVal == 930809136 && v2_1.UnsafePtr == nil {
				__t0 = gopurs_runtime.Array([]gopurs_runtime.Value{})
				goto end_branch_0
			} else {

			}
		}
		{
			if v2_1.Type == 9 && v2_1.IntVal == 930809136 && v2_1.UnsafePtr != nil {
				__t0 = gopurs_runtime.Array(Call_Data_Array_singleton((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_1.UnsafePtr).V0))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	}), f_0))
}

func Call_Data_Array_filterA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): traverse1_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope434)] (TypeApp (TypeVar f$scope435) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope434), Boolean])])), (Array (TypeVar a$scope434))] (TypeApp (TypeVar f$scope435) [(Array (ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope434), Boolean]))]))
	traverse1_1_0 := gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
	_ = traverse1_1_0
	// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope435)])
	Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_2_1
	return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(traverse1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), x_4), gopurs_runtime.Apply(p_3, x_4))
		})), gopurs_runtime.Apply(Functor0_2_1.V0, Call_Data_Array_mapMaybe(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) != (0) {
					__t2 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, true}
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
		}))))
	})
}

func Call_Data_Array_any(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return (gopurs_runtime.UncurriedApp2(Get_Data_Array_anyImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_Data_Array_nubByEq(eq2_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var eq2_0 gopurs_runtime.Value = eq2_0_loop
	_ = eq2_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		arr_2_0 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
		_ = arr_2_0
		_dollar___unused_3_1 := gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_foreach(), gopurs_runtime.Array(xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=Any
				__local_var_4_2 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_HeytingAlgebra_boolNot(), gopurs_runtime.Apply(Get_Data_Array_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(eq2_0, v_4, x_3)
				})))
				_ = __local_var_4_2
				__local_var_5_4 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), arr_2_0)
				_ = __local_var_5_4
				e_5_3 := gopurs_runtime.Apply(__local_var_4_2, __local_var_5_4)
				_ = e_5_3
				// TAST (Let): __local_var_6_5 shape=App(Var) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope454), Unit])
				__local_var_6_5 := gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), x_3, arr_2_0)
				}))
				_ = __local_var_6_5
				var __t6 gopurs_runtime.Value
				{
					if (e_5_3.IntVal) != (0) {
						__t6 = __local_var_6_5
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
				}
			end_branch_6:
				return gopurs_runtime.Apply(__t6, gopurs_runtime.Value{})
			})
		}), gopurs_runtime.Value{})
		_ = _dollar___unused_3_1
		return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), arr_2_0)
	})).UnsafePtr))
}

func Call_Data_Array_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_nubByEq(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_unionBy(eq2_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var eq2_0 gopurs_runtime.Value = eq2_0_loop
	_ = eq2_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	var ys_2 []gopurs_runtime.Value = ys_2_loop
	_ = ys_2
	arrayIntrinsic_arg_0 := gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Array(Call_Data_Array_deleteBy(eq2_0, a_4, *((*[]gopurs_runtime.Value)(b_3.UnsafePtr))))
	})
	_ = arrayIntrinsic_arg_0
	arrayIntrinsic_arg_1 := gopurs_runtime.Array(Call_Data_Array_nubByEq(eq2_0, ys_2))
	_ = arrayIntrinsic_arg_1
	arrayIntrinsic_arg_2 := gopurs_runtime.Array(xs_1)
	_ = arrayIntrinsic_arg_2
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_1), func() gopurs_runtime.Value {
		arr_val_foldlArray1 := arrayIntrinsic_arg_2
		_ = arr_val_foldlArray1
		res_go_foldlArray1 := arrayIntrinsic_arg_1
		_ = res_go_foldlArray1
		arr_go_foldlArray1 := (*[]gopurs_runtime.Value)(arr_val_foldlArray1.UnsafePtr)
		_ = arr_go_foldlArray1
		for _, v_foldlArray1 := range *arr_go_foldlArray1 {
			res_go_foldlArray1 = gopurs_runtime.Apply2(arrayIntrinsic_arg_0, res_go_foldlArray1, v_foldlArray1)
		}
		return res_go_foldlArray1
	}()).UnsafePtr))
}

func Call_Data_Array_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_unionBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var xs_2 []gopurs_runtime.Value = xs_2_loop
	_ = xs_2
	arrayIndex_arg_1 := Get_Data_Maybe_Just()
	_ = arrayIndex_arg_1
	arrayIndex_arg_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}())
	_ = arrayIndex_arg_2
	arrayIndex_arg_3 := xs_2
	_ = arrayIndex_arg_3
	arrayIndex_arg_4 := i_0
	_ = arrayIndex_arg_4
	// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope465)])
	__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		arrayIndex_value := arrayIndex_arg_3
		_ = arrayIndex_value
		arrayIndex_source_5 := arrayIndex_value
		_ = arrayIndex_source_5
		if ((arrayIndex_arg_4) >= (0)) && ((arrayIndex_arg_4) < (int64(len(arrayIndex_source_5)))) {
			return gopurs_runtime.Apply(arrayIndex_arg_1, (arrayIndex_source_5)[arrayIndex_arg_4])
		} else {

		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_2)}
	}())
	_ = __local_var_3_0
	var __t7 gopurs_runtime.Value
	{
		if __local_var_3_0 == nil {
			__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
			goto end_branch_7
		} else {

		}
	}
	{
		if __local_var_3_0 != nil {
			// TAST (Let): v_4_5 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope465)])
			v_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (__local_var_3_0).V0))
			_ = v_4_5
			var __t6 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
			{
				if v_4_5 == nil {
					__t6 = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_2))))
					goto end_branch_6
				} else {

				}
			}
			{
				if v_4_5 != nil {
					__t6 = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))}, gopurs_runtime.Int(i_0), (v_4_5).V0, gopurs_runtime.Array(xs_2))))
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_6:
			__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(__t6))}
			goto end_branch_7
		} else {

		}
	}
	{
		__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_7:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t7))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_all(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return (gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_Data_Array_all__3951006479(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop []string) bool {
all__3951006479:
	for {
		if false {
			continue all__3951006479
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 []string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), __eta_norm_1_0, func() gopurs_runtime.Value {
			arr := __eta_norm_0_1
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal) != (0)
	}
}

func Rebox_Data_Array_138441832_1728839155(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1
	return out
}

func Rebox_Data_Array_138441832_1807945000(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value]{}
	out.V0 = *((*[]gopurs_runtime.Value)(in.V0.UnsafePtr))
	out.V1 = *((*[]gopurs_runtime.Value)(in.V1.UnsafePtr))
	return out
}

func Rebox_Data_Array_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
	out.V0 = uint32(in.V0.IntVal)
	out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Array_1728839155_138441832(in *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Data_Array_1807945000_138441832(in *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Array(in.V0)
	out.V1 = gopurs_runtime.Array(in.V1)
	return out
}

func Rebox_Data_Array_2206451173_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[struct {
	a gopurs_runtime.Value
	b int64
	/*gopurs:record-v1:2:0061,0062*/
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_2304530460_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_3132786365_138441832(in.V0))}
	return out
}

func Rebox_Data_Array_2660862937_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct {
	a gopurs_runtime.Value
	b int64
	/*gopurs:record-v1:2:0061,0062*/
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("a", "b", orig.a, gopurs_runtime.Int(orig.b))
	}()
	return out
}

func Rebox_Data_Array_2807397954_3094389156(in *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Array(in.V0)
	return out
}

func Rebox_Data_Array_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Array_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Array_3094389156_1785233618(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]]{}
	out.V0 = Rebox_Data_Array_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Array_3094389156_2304530460(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
	out.V0 = Rebox_Data_Array_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Array_3094389156_2807397954(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]{}
	out.V0 = *((*[]gopurs_runtime.Value)(in.V0.UnsafePtr))
	return out
}

func Rebox_Data_Array_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Array_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_3741347833_3552963512(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Array_394862806_3988677819(in *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Array_3988677819_394862806(in *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Get_Data_Array__deleteAt() gopurs_runtime.Value {
	return _Gopurs_Data_Array__DeleteAt
}

func Get_Data_Array__insertAt() gopurs_runtime.Value {
	return _Gopurs_Data_Array__InsertAt
}

func Get_Data_Array__updateAt() gopurs_runtime.Value {
	return _Gopurs_Data_Array__UpdateAt
}

func Get_Data_Array_allImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_AllImpl
}

func Get_Data_Array_anyImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_AnyImpl
}

func Get_Data_Array_concat() gopurs_runtime.Value {
	return _Gopurs_Data_Array_Concat
}

func Get_Data_Array_filterImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FilterImpl
}

func Get_Data_Array_findIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FindIndexImpl
}

func Get_Data_Array_findLastIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FindLastIndexImpl
}

func Get_Data_Array_findMapImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FindMapImpl
}

func Get_Data_Array_fromFoldableImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FromFoldableImpl
}

func Get_Data_Array_indexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_IndexImpl
}

func Get_Data_Array_length() gopurs_runtime.Value {
	return _Gopurs_Data_Array_Length
}

func Get_Data_Array_partitionImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_PartitionImpl
}

func Get_Data_Array_rangeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_RangeImpl
}

func Get_Data_Array_replicateImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ReplicateImpl
}

func Get_Data_Array_reverse() gopurs_runtime.Value {
	return _Gopurs_Data_Array_Reverse
}

func Get_Data_Array_scanlImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ScanlImpl
}

func Get_Data_Array_scanrImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ScanrImpl
}

func Get_Data_Array_sliceImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_SliceImpl
}

func Get_Data_Array_sortByImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_SortByImpl
}

func Get_Data_Array_unconsImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_UnconsImpl
}

func Get_Data_Array_unsafeIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_UnsafeIndexImpl
}

func Get_Data_Array_zipWithImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ZipWithImpl
}
