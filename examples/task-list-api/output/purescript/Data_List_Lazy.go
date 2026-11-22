package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Lazy_apply gopurs_runtime.Value
var once_Data_List_Lazy_apply sync.Once

func Get_Data_List_Lazy_apply() gopurs_runtime.Value {
	once_Data_List_Lazy_apply.Do(func() {
		cache_Data_List_Lazy_apply = gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Lazy_applyLazy()).V1
	})
	return cache_Data_List_Lazy_apply
}

var cache_Data_List_Lazy_go__map gopurs_runtime.Value
var once_Data_List_Lazy_go__map sync.Once

func Get_Data_List_Lazy_go__map() gopurs_runtime.Value {
	once_Data_List_Lazy_go__map.Do(func() {
		cache_Data_List_Lazy_go__map = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_go__map
}

var cache_Data_List_Lazy_unwrap gopurs_runtime.Value
var once_Data_List_Lazy_unwrap sync.Once

func Get_Data_List_Lazy_unwrap() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap.Do(func() {
		cache_Data_List_Lazy_unwrap = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap
}

var cache_Data_List_Lazy_unwrap1 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap1 sync.Once

func Get_Data_List_Lazy_unwrap1() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap1.Do(func() {
		cache_Data_List_Lazy_unwrap1 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap1
}

var cache_Data_List_Lazy_sequence gopurs_runtime.Value
var once_Data_List_Lazy_sequence sync.Once

func Get_Data_List_Lazy_sequence() gopurs_runtime.Value {
	once_Data_List_Lazy_sequence.Do(func() {
		cache_Data_List_Lazy_sequence = gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()).V2
	})
	return cache_Data_List_Lazy_sequence
}

var cache_Data_List_Lazy_sub gopurs_runtime.Value
var once_Data_List_Lazy_sub sync.Once

func Get_Data_List_Lazy_sub() gopurs_runtime.Value {
	once_Data_List_Lazy_sub.Do(func() {
		cache_Data_List_Lazy_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_List_Lazy_sub
}

var cache_Data_List_Lazy_map1 gopurs_runtime.Value
var once_Data_List_Lazy_map1 sync.Once

func Get_Data_List_Lazy_map1() gopurs_runtime.Value {
	once_Data_List_Lazy_map1.Do(func() {
		cache_Data_List_Lazy_map1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map1
}

var cache_Data_List_Lazy_unwrap2 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap2 sync.Once

func Get_Data_List_Lazy_unwrap2() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap2.Do(func() {
		cache_Data_List_Lazy_unwrap2 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap2
}

var cache_Data_List_Lazy_map2 gopurs_runtime.Value
var once_Data_List_Lazy_map2 sync.Once

func Get_Data_List_Lazy_map2() gopurs_runtime.Value {
	once_Data_List_Lazy_map2.Do(func() {
		cache_Data_List_Lazy_map2 = Rebox_Data_List_Lazy_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_Lazy_map2
}

var cache_Data_List_Lazy_compose gopurs_runtime.Value
var once_Data_List_Lazy_compose sync.Once

func Get_Data_List_Lazy_compose() gopurs_runtime.Value {
	once_Data_List_Lazy_compose.Do(func() {
		cache_Data_List_Lazy_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose
}

var cache_Data_List_Lazy_compose1 gopurs_runtime.Value
var once_Data_List_Lazy_compose1 sync.Once

func Get_Data_List_Lazy_compose1() gopurs_runtime.Value {
	once_Data_List_Lazy_compose1.Do(func() {
		cache_Data_List_Lazy_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose1
}

var cache_Data_List_Lazy_map3 gopurs_runtime.Value
var once_Data_List_Lazy_map3 sync.Once

func Get_Data_List_Lazy_map3() gopurs_runtime.Value {
	once_Data_List_Lazy_map3.Do(func() {
		cache_Data_List_Lazy_map3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map3
}

var cache_Data_List_Lazy_unwrap3 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap3 sync.Once

func Get_Data_List_Lazy_unwrap3() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap3.Do(func() {
		cache_Data_List_Lazy_unwrap3 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap3
}

var cache_Data_List_Lazy_lessThanOrEq gopurs_runtime.Value
var once_Data_List_Lazy_lessThanOrEq sync.Once

func Get_Data_List_Lazy_lessThanOrEq() gopurs_runtime.Value {
	once_Data_List_Lazy_lessThanOrEq.Do(func() {
		cache_Data_List_Lazy_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_List_Lazy_lessThanOrEq
}

var cache_Data_List_Lazy_compose2 gopurs_runtime.Value
var once_Data_List_Lazy_compose2 sync.Once

func Get_Data_List_Lazy_compose2() gopurs_runtime.Value {
	once_Data_List_Lazy_compose2.Do(func() {
		cache_Data_List_Lazy_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose2
}

var cache_Data_List_Lazy_compose3 gopurs_runtime.Value
var once_Data_List_Lazy_compose3 sync.Once

func Get_Data_List_Lazy_compose3() gopurs_runtime.Value {
	once_Data_List_Lazy_compose3.Do(func() {
		cache_Data_List_Lazy_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose3
}

var cache_Data_List_Lazy_map4 gopurs_runtime.Value
var once_Data_List_Lazy_map4 sync.Once

func Get_Data_List_Lazy_map4() gopurs_runtime.Value {
	once_Data_List_Lazy_map4.Do(func() {
		cache_Data_List_Lazy_map4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map4
}

var cache_Data_List_Lazy_unwrap4 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap4 sync.Once

func Get_Data_List_Lazy_unwrap4() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap4.Do(func() {
		cache_Data_List_Lazy_unwrap4 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap4
}

var cache_Data_List_Lazy_map5 gopurs_runtime.Value
var once_Data_List_Lazy_map5 sync.Once

func Get_Data_List_Lazy_map5() gopurs_runtime.Value {
	once_Data_List_Lazy_map5.Do(func() {
		cache_Data_List_Lazy_map5 = Rebox_Data_List_Lazy_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_Lazy_map5
}

var cache_Data_List_Lazy_tailRecM2 gopurs_runtime.Value
var once_Data_List_Lazy_tailRecM2 sync.Once

func Get_Data_List_Lazy_tailRecM2() gopurs_runtime.Value {
	once_Data_List_Lazy_tailRecM2.Do(func() {
		cache_Data_List_Lazy_tailRecM2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_tailRecM2(f_0_box, a_1_box, b_2_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_Lazy_tailRecM2
}

var cache_Data_List_Lazy_foldr gopurs_runtime.Value
var once_Data_List_Lazy_foldr sync.Once

func Get_Data_List_Lazy_foldr() gopurs_runtime.Value {
	once_Data_List_Lazy_foldr.Do(func() {
		cache_Data_List_Lazy_foldr = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2
	})
	return cache_Data_List_Lazy_foldr
}

var cache_Data_List_Lazy_go__append gopurs_runtime.Value
var once_Data_List_Lazy_go__append sync.Once

func Get_Data_List_Lazy_go__append() gopurs_runtime.Value {
	once_Data_List_Lazy_go__append.Do(func() {
		cache_Data_List_Lazy_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_List_Lazy_go__append
}

var cache_Data_List_Lazy_map6 gopurs_runtime.Value
var once_Data_List_Lazy_map6 sync.Once

func Get_Data_List_Lazy_map6() gopurs_runtime.Value {
	once_Data_List_Lazy_map6.Do(func() {
		cache_Data_List_Lazy_map6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map6
}

var cache_Data_List_Lazy_unwrap5 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap5 sync.Once

func Get_Data_List_Lazy_unwrap5() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap5.Do(func() {
		cache_Data_List_Lazy_unwrap5 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap5
}

var cache_Data_List_Lazy_go__defer gopurs_runtime.Value
var once_Data_List_Lazy_go__defer sync.Once

func Get_Data_List_Lazy_go__defer() gopurs_runtime.Value {
	once_Data_List_Lazy_go__defer.Do(func() {
		cache_Data_List_Lazy_go__defer = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer")
	})
	return cache_Data_List_Lazy_go__defer
}

var cache_Data_List_Lazy_foldl gopurs_runtime.Value
var once_Data_List_Lazy_foldl sync.Once

func Get_Data_List_Lazy_foldl() gopurs_runtime.Value {
	once_Data_List_Lazy_foldl.Do(func() {
		cache_Data_List_Lazy_foldl = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1
	})
	return cache_Data_List_Lazy_foldl
}

var cache_Data_List_Lazy_lessThan gopurs_runtime.Value
var once_Data_List_Lazy_lessThan sync.Once

func Get_Data_List_Lazy_lessThan() gopurs_runtime.Value {
	once_Data_List_Lazy_lessThan.Do(func() {
		cache_Data_List_Lazy_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_List_Lazy_lessThan
}

var cache_Data_List_Lazy_one gopurs_runtime.Value
var once_Data_List_Lazy_one sync.Once

func Get_Data_List_Lazy_one() gopurs_runtime.Value {
	once_Data_List_Lazy_one.Do(func() {
		cache_Data_List_Lazy_one = gopurs_runtime.Int(int64(1))
	})
	return cache_Data_List_Lazy_one
}

var cache_Data_List_Lazy_fix gopurs_runtime.Value
var once_Data_List_Lazy_fix sync.Once

func Get_Data_List_Lazy_fix() gopurs_runtime.Value {
	once_Data_List_Lazy_fix.Do(func() {
		cache_Data_List_Lazy_fix = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_fix(f_0_box)
		})
	})
	return cache_Data_List_Lazy_fix
}

var cache_Data_List_Lazy_greaterThan gopurs_runtime.Value
var once_Data_List_Lazy_greaterThan sync.Once

func Get_Data_List_Lazy_greaterThan() gopurs_runtime.Value {
	once_Data_List_Lazy_greaterThan.Do(func() {
		cache_Data_List_Lazy_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_List_Lazy_greaterThan
}

var cache_Data_List_Lazy_greaterThanOrEq gopurs_runtime.Value
var once_Data_List_Lazy_greaterThanOrEq sync.Once

func Get_Data_List_Lazy_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_List_Lazy_greaterThanOrEq.Do(func() {
		cache_Data_List_Lazy_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_List_Lazy_greaterThanOrEq
}

var cache_Data_List_Lazy_unfoldr gopurs_runtime.Value
var once_Data_List_Lazy_unfoldr sync.Once

func Get_Data_List_Lazy_unfoldr() gopurs_runtime.Value {
	once_Data_List_Lazy_unfoldr.Do(func() {
		cache_Data_List_Lazy_unfoldr = gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldableList()).V1
	})
	return cache_Data_List_Lazy_unfoldr
}

var cache_Data_List_Lazy_add gopurs_runtime.Value
var once_Data_List_Lazy_add sync.Once

func Get_Data_List_Lazy_add() gopurs_runtime.Value {
	once_Data_List_Lazy_add.Do(func() {
		cache_Data_List_Lazy_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_List_Lazy_add
}

var cache_Data_List_Lazy_foldr1 gopurs_runtime.Value
var once_Data_List_Lazy_foldr1 sync.Once

func Get_Data_List_Lazy_foldr1() gopurs_runtime.Value {
	once_Data_List_Lazy_foldr1.Do(func() {
		cache_Data_List_Lazy_foldr1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2
	})
	return cache_Data_List_Lazy_foldr1
}

var cache_Data_List_Lazy_map7 gopurs_runtime.Value
var once_Data_List_Lazy_map7 sync.Once

func Get_Data_List_Lazy_map7() gopurs_runtime.Value {
	once_Data_List_Lazy_map7.Do(func() {
		cache_Data_List_Lazy_map7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map7
}

var cache_Data_List_Lazy_compose4 gopurs_runtime.Value
var once_Data_List_Lazy_compose4 sync.Once

func Get_Data_List_Lazy_compose4() gopurs_runtime.Value {
	once_Data_List_Lazy_compose4.Do(func() {
		cache_Data_List_Lazy_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose4
}

var cache_Data_List_Lazy_compose5 gopurs_runtime.Value
var once_Data_List_Lazy_compose5 sync.Once

func Get_Data_List_Lazy_compose5() gopurs_runtime.Value {
	once_Data_List_Lazy_compose5.Do(func() {
		cache_Data_List_Lazy_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose5
}

var cache_Data_List_Lazy_map8 gopurs_runtime.Value
var once_Data_List_Lazy_map8 sync.Once

func Get_Data_List_Lazy_map8() gopurs_runtime.Value {
	once_Data_List_Lazy_map8.Do(func() {
		cache_Data_List_Lazy_map8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map8
}

var cache_Data_List_Lazy_unwrap6 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap6 sync.Once

func Get_Data_List_Lazy_unwrap6() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap6.Do(func() {
		cache_Data_List_Lazy_unwrap6 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap6
}

var cache_Data_List_Lazy_fix1 gopurs_runtime.Value
var once_Data_List_Lazy_fix1 sync.Once

func Get_Data_List_Lazy_fix1() gopurs_runtime.Value {
	once_Data_List_Lazy_fix1.Do(func() {
		cache_Data_List_Lazy_fix1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_fix1(f_0_box)
		})
	})
	return cache_Data_List_Lazy_fix1
}

var cache_Data_List_Lazy_map9 gopurs_runtime.Value
var once_Data_List_Lazy_map9 sync.Once

func Get_Data_List_Lazy_map9() gopurs_runtime.Value {
	once_Data_List_Lazy_map9.Do(func() {
		cache_Data_List_Lazy_map9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0
	})
	return cache_Data_List_Lazy_map9
}

var cache_Data_List_Lazy_map10 gopurs_runtime.Value
var once_Data_List_Lazy_map10 sync.Once

func Get_Data_List_Lazy_map10() gopurs_runtime.Value {
	once_Data_List_Lazy_map10.Do(func() {
		cache_Data_List_Lazy_map10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map10
}

var cache_Data_List_Lazy_unwrap7 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap7 sync.Once

func Get_Data_List_Lazy_unwrap7() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap7.Do(func() {
		cache_Data_List_Lazy_unwrap7 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap7
}

var cache_Data_List_Lazy_map11 gopurs_runtime.Value
var once_Data_List_Lazy_map11 sync.Once

func Get_Data_List_Lazy_map11() gopurs_runtime.Value {
	once_Data_List_Lazy_map11.Do(func() {
		cache_Data_List_Lazy_map11 = Rebox_Data_List_Lazy_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_Lazy_map11
}

var cache_Data_List_Lazy_map12 gopurs_runtime.Value
var once_Data_List_Lazy_map12 sync.Once

func Get_Data_List_Lazy_map12() gopurs_runtime.Value {
	once_Data_List_Lazy_map12.Do(func() {
		cache_Data_List_Lazy_map12 = Rebox_Data_List_Lazy_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_Lazy_map12
}

var cache_Data_List_Lazy_compose6 gopurs_runtime.Value
var once_Data_List_Lazy_compose6 sync.Once

func Get_Data_List_Lazy_compose6() gopurs_runtime.Value {
	once_Data_List_Lazy_compose6.Do(func() {
		cache_Data_List_Lazy_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose6
}

var cache_Data_List_Lazy_compose7 gopurs_runtime.Value
var once_Data_List_Lazy_compose7 sync.Once

func Get_Data_List_Lazy_compose7() gopurs_runtime.Value {
	once_Data_List_Lazy_compose7.Do(func() {
		cache_Data_List_Lazy_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose7
}

var cache_Data_List_Lazy_map13 gopurs_runtime.Value
var once_Data_List_Lazy_map13 sync.Once

func Get_Data_List_Lazy_map13() gopurs_runtime.Value {
	once_Data_List_Lazy_map13.Do(func() {
		cache_Data_List_Lazy_map13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map13
}

var cache_Data_List_Lazy_unwrap8 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap8 sync.Once

func Get_Data_List_Lazy_unwrap8() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap8.Do(func() {
		cache_Data_List_Lazy_unwrap8 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap8
}

var cache_Data_List_Lazy_map14 gopurs_runtime.Value
var once_Data_List_Lazy_map14 sync.Once

func Get_Data_List_Lazy_map14() gopurs_runtime.Value {
	once_Data_List_Lazy_map14.Do(func() {
		cache_Data_List_Lazy_map14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map14
}

var cache_Data_List_Lazy_unwrap9 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap9 sync.Once

func Get_Data_List_Lazy_unwrap9() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap9.Do(func() {
		cache_Data_List_Lazy_unwrap9 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap9
}

var cache_Data_List_Lazy_bind gopurs_runtime.Value
var once_Data_List_Lazy_bind sync.Once

func Get_Data_List_Lazy_bind() gopurs_runtime.Value {
	once_Data_List_Lazy_bind.Do(func() {
		cache_Data_List_Lazy_bind = Rebox_Data_List_Lazy_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1
	})
	return cache_Data_List_Lazy_bind
}

var cache_Data_List_Lazy_pure gopurs_runtime.Value
var once_Data_List_Lazy_pure sync.Once

func Get_Data_List_Lazy_pure() gopurs_runtime.Value {
	once_Data_List_Lazy_pure.Do(func() {
		cache_Data_List_Lazy_pure = Get_Control_Applicative_pure__3917743074()
	})
	return cache_Data_List_Lazy_pure
}

var cache_Data_List_Lazy_map15 gopurs_runtime.Value
var once_Data_List_Lazy_map15 sync.Once

func Get_Data_List_Lazy_map15() gopurs_runtime.Value {
	once_Data_List_Lazy_map15.Do(func() {
		cache_Data_List_Lazy_map15 = Get_Data_Functor_map__3627393546()
	})
	return cache_Data_List_Lazy_map15
}

var cache_Data_List_Lazy_compose8 gopurs_runtime.Value
var once_Data_List_Lazy_compose8 sync.Once

func Get_Data_List_Lazy_compose8() gopurs_runtime.Value {
	once_Data_List_Lazy_compose8.Do(func() {
		cache_Data_List_Lazy_compose8 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose8
}

var cache_Data_List_Lazy_compose9 gopurs_runtime.Value
var once_Data_List_Lazy_compose9 sync.Once

func Get_Data_List_Lazy_compose9() gopurs_runtime.Value {
	once_Data_List_Lazy_compose9.Do(func() {
		cache_Data_List_Lazy_compose9 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose9
}

var cache_Data_List_Lazy_map16 gopurs_runtime.Value
var once_Data_List_Lazy_map16 sync.Once

func Get_Data_List_Lazy_map16() gopurs_runtime.Value {
	once_Data_List_Lazy_map16.Do(func() {
		cache_Data_List_Lazy_map16 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map16
}

var cache_Data_List_Lazy_unwrap10 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap10 sync.Once

func Get_Data_List_Lazy_unwrap10() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap10.Do(func() {
		cache_Data_List_Lazy_unwrap10 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap10
}

var cache_Data_List_Lazy_any gopurs_runtime.Value
var once_Data_List_Lazy_any sync.Once

func Get_Data_List_Lazy_any() gopurs_runtime.Value {
	once_Data_List_Lazy_any.Do(func() {
		cache_Data_List_Lazy_any = gopurs_runtime.Apply(Call_Data_Foldable_any(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList())), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3591112874_2663347022(Rebox_Data_List_Lazy_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))})
	})
	return cache_Data_List_Lazy_any
}

var cache_Data_List_Lazy_not gopurs_runtime.Value
var once_Data_List_Lazy_not sync.Once

func Get_Data_List_Lazy_not() gopurs_runtime.Value {
	once_Data_List_Lazy_not.Do(func() {
		cache_Data_List_Lazy_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Data_List_Lazy_not
}

var cache_Data_List_Lazy_compose10 gopurs_runtime.Value
var once_Data_List_Lazy_compose10 sync.Once

func Get_Data_List_Lazy_compose10() gopurs_runtime.Value {
	once_Data_List_Lazy_compose10.Do(func() {
		cache_Data_List_Lazy_compose10 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose10
}

var cache_Data_List_Lazy_compose11 gopurs_runtime.Value
var once_Data_List_Lazy_compose11 sync.Once

func Get_Data_List_Lazy_compose11() gopurs_runtime.Value {
	once_Data_List_Lazy_compose11.Do(func() {
		cache_Data_List_Lazy_compose11 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose11
}

var cache_Data_List_Lazy_map17 gopurs_runtime.Value
var once_Data_List_Lazy_map17 sync.Once

func Get_Data_List_Lazy_map17() gopurs_runtime.Value {
	once_Data_List_Lazy_map17.Do(func() {
		cache_Data_List_Lazy_map17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map17
}

var cache_Data_List_Lazy_unwrap11 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap11 sync.Once

func Get_Data_List_Lazy_unwrap11() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap11.Do(func() {
		cache_Data_List_Lazy_unwrap11 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap11
}

var cache_Data_List_Lazy_compose12 gopurs_runtime.Value
var once_Data_List_Lazy_compose12 sync.Once

func Get_Data_List_Lazy_compose12() gopurs_runtime.Value {
	once_Data_List_Lazy_compose12.Do(func() {
		cache_Data_List_Lazy_compose12 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose12
}

var cache_Data_List_Lazy_compose13 gopurs_runtime.Value
var once_Data_List_Lazy_compose13 sync.Once

func Get_Data_List_Lazy_compose13() gopurs_runtime.Value {
	once_Data_List_Lazy_compose13.Do(func() {
		cache_Data_List_Lazy_compose13 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose13
}

var cache_Data_List_Lazy_compose14 gopurs_runtime.Value
var once_Data_List_Lazy_compose14 sync.Once

func Get_Data_List_Lazy_compose14() gopurs_runtime.Value {
	once_Data_List_Lazy_compose14.Do(func() {
		cache_Data_List_Lazy_compose14 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose14
}

var cache_Data_List_Lazy_map18 gopurs_runtime.Value
var once_Data_List_Lazy_map18 sync.Once

func Get_Data_List_Lazy_map18() gopurs_runtime.Value {
	once_Data_List_Lazy_map18.Do(func() {
		cache_Data_List_Lazy_map18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map18
}

var cache_Data_List_Lazy_unwrap12 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap12 sync.Once

func Get_Data_List_Lazy_unwrap12() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap12.Do(func() {
		cache_Data_List_Lazy_unwrap12 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap12
}

var cache_Data_List_Lazy_map19 gopurs_runtime.Value
var once_Data_List_Lazy_map19 sync.Once

func Get_Data_List_Lazy_map19() gopurs_runtime.Value {
	once_Data_List_Lazy_map19.Do(func() {
		cache_Data_List_Lazy_map19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map19
}

var cache_Data_List_Lazy_unwrap13 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap13 sync.Once

func Get_Data_List_Lazy_unwrap13() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap13.Do(func() {
		cache_Data_List_Lazy_unwrap13 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap13
}

var cache_Data_List_Lazy_append1 gopurs_runtime.Value
var once_Data_List_Lazy_append1 sync.Once

func Get_Data_List_Lazy_append1() gopurs_runtime.Value {
	once_Data_List_Lazy_append1.Do(func() {
		cache_Data_List_Lazy_append1 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_Lazy_append1
}

var cache_Data_List_Lazy_foldl1 gopurs_runtime.Value
var once_Data_List_Lazy_foldl1 sync.Once

func Get_Data_List_Lazy_foldl1() gopurs_runtime.Value {
	once_Data_List_Lazy_foldl1.Do(func() {
		cache_Data_List_Lazy_foldl1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1
	})
	return cache_Data_List_Lazy_foldl1
}

var cache_Data_List_Lazy_map20 gopurs_runtime.Value
var once_Data_List_Lazy_map20 sync.Once

func Get_Data_List_Lazy_map20() gopurs_runtime.Value {
	once_Data_List_Lazy_map20.Do(func() {
		cache_Data_List_Lazy_map20 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map20
}

var cache_Data_List_Lazy_unwrap14 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap14 sync.Once

func Get_Data_List_Lazy_unwrap14() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap14.Do(func() {
		cache_Data_List_Lazy_unwrap14 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap14
}

var cache_Data_List_Lazy_foldl2 gopurs_runtime.Value
var once_Data_List_Lazy_foldl2 sync.Once

func Get_Data_List_Lazy_foldl2() gopurs_runtime.Value {
	once_Data_List_Lazy_foldl2.Do(func() {
		cache_Data_List_Lazy_foldl2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1
	})
	return cache_Data_List_Lazy_foldl2
}

var cache_Data_List_Lazy_fix2 gopurs_runtime.Value
var once_Data_List_Lazy_fix2 sync.Once

func Get_Data_List_Lazy_fix2() gopurs_runtime.Value {
	once_Data_List_Lazy_fix2.Do(func() {
		cache_Data_List_Lazy_fix2 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_fix2(f_0_box)
		})
	})
	return cache_Data_List_Lazy_fix2
}

var cache_Data_List_Lazy_append2 gopurs_runtime.Value
var once_Data_List_Lazy_append2 sync.Once

func Get_Data_List_Lazy_append2() gopurs_runtime.Value {
	once_Data_List_Lazy_append2.Do(func() {
		cache_Data_List_Lazy_append2 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_Lazy_append2
}

var cache_Data_List_Lazy_bind1 gopurs_runtime.Value
var once_Data_List_Lazy_bind1 sync.Once

func Get_Data_List_Lazy_bind1() gopurs_runtime.Value {
	once_Data_List_Lazy_bind1.Do(func() {
		cache_Data_List_Lazy_bind1 = gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1
	})
	return cache_Data_List_Lazy_bind1
}

var cache_Data_List_Lazy_identity gopurs_runtime.Value
var once_Data_List_Lazy_identity sync.Once

func Get_Data_List_Lazy_identity() gopurs_runtime.Value {
	once_Data_List_Lazy_identity.Do(func() {
		cache_Data_List_Lazy_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_identity(x_0_box)
		})
	})
	return cache_Data_List_Lazy_identity
}

var cache_Data_List_Lazy_map21 gopurs_runtime.Value
var once_Data_List_Lazy_map21 sync.Once

func Get_Data_List_Lazy_map21() gopurs_runtime.Value {
	once_Data_List_Lazy_map21.Do(func() {
		cache_Data_List_Lazy_map21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_map21
}

var cache_Data_List_Lazy_unwrap15 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap15 sync.Once

func Get_Data_List_Lazy_unwrap15() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap15.Do(func() {
		cache_Data_List_Lazy_unwrap15 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_unwrap15
}

var cache_Data_List_Lazy_compose15 gopurs_runtime.Value
var once_Data_List_Lazy_compose15 sync.Once

func Get_Data_List_Lazy_compose15() gopurs_runtime.Value {
	once_Data_List_Lazy_compose15.Do(func() {
		cache_Data_List_Lazy_compose15 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_compose15
}

var cache_Data_List_Lazy_Pattern gopurs_runtime.Value
var once_Data_List_Lazy_Pattern sync.Once

func Get_Data_List_Lazy_Pattern() gopurs_runtime.Value {
	once_Data_List_Lazy_Pattern.Do(func() {
		cache_Data_List_Lazy_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Pattern(x_0_box)
		})
	})
	return cache_Data_List_Lazy_Pattern
}

var cache_Data_List_Lazy_zipWith gopurs_runtime.Value
var once_Data_List_Lazy_zipWith sync.Once

func Get_Data_List_Lazy_zipWith() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWith.Do(func() {
		cache_Data_List_Lazy_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_zipWith(f_0_box, xs_1_box, ys_2_box)
		})
	})
	return cache_Data_List_Lazy_zipWith
}

var cache_Data_List_Lazy_zipWithA gopurs_runtime.Value
var once_Data_List_Lazy_zipWithA sync.Once

func Get_Data_List_Lazy_zipWithA() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWithA.Do(func() {
		cache_Data_List_Lazy_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, xs_2_box, ys_3_box)
		})
	})
	return cache_Data_List_Lazy_zipWithA
}

var cache_Data_List_Lazy_zip gopurs_runtime.Value
var once_Data_List_Lazy_zip sync.Once

func Get_Data_List_Lazy_zip() gopurs_runtime.Value {
	once_Data_List_Lazy_zip.Do(func() {
		cache_Data_List_Lazy_zip = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_zip(xs_0_box, ys_1_box)
		})
	})
	return cache_Data_List_Lazy_zip
}

var cache_Data_List_Lazy_updateAt gopurs_runtime.Value
var once_Data_List_Lazy_updateAt sync.Once

func Get_Data_List_Lazy_updateAt() gopurs_runtime.Value {
	once_Data_List_Lazy_updateAt.Do(func() {
		cache_Data_List_Lazy_updateAt = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_updateAt(n_0_box.IntVal, x_1_box, xs_2_box)
		})
	})
	return cache_Data_List_Lazy_updateAt
}

var cache_Data_List_Lazy_unzip gopurs_runtime.Value
var once_Data_List_Lazy_unzip sync.Once

func Get_Data_List_Lazy_unzip() gopurs_runtime.Value {
	once_Data_List_Lazy_unzip.Do(func() {
		cache_Data_List_Lazy_unzip = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
			__local_var_1_0 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
			_ = __local_var_1_0
			// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
			__local_var_2_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
			_ = __local_var_2_1
			return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
				__local_var_4_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
				_ = __local_var_4_2
				// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
				__local_var_5_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
				_ = __local_var_5_3
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_1_0, __local_var_4_2}))}
					})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_1, __local_var_5_3}))}
					}))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			})
		}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))})
	})
	return cache_Data_List_Lazy_unzip
}

var cache_Data_List_Lazy_uncons gopurs_runtime.Value
var once_Data_List_Lazy_uncons sync.Once

func Get_Data_List_Lazy_uncons() gopurs_runtime.Value {
	once_Data_List_Lazy_uncons.Do(func() {
		cache_Data_List_Lazy_uncons = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_Lazy_uncons
}

var cache_Data_List_Lazy_toUnfoldable gopurs_runtime.Value
var once_Data_List_Lazy_toUnfoldable sync.Once

func Get_Data_List_Lazy_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_toUnfoldable.Do(func() {
		cache_Data_List_Lazy_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_List_Lazy_toUnfoldable
}

var cache_Data_List_Lazy_takeWhile gopurs_runtime.Value
var once_Data_List_Lazy_takeWhile sync.Once

func Get_Data_List_Lazy_takeWhile() gopurs_runtime.Value {
	once_Data_List_Lazy_takeWhile.Do(func() {
		cache_Data_List_Lazy_takeWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_takeWhile(p_0_box)
		})
	})
	return cache_Data_List_Lazy_takeWhile
}

var cache_Data_List_Lazy_take gopurs_runtime.Value
var once_Data_List_Lazy_take sync.Once

func Get_Data_List_Lazy_take() gopurs_runtime.Value {
	once_Data_List_Lazy_take.Do(func() {
		cache_Data_List_Lazy_take = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_take(n_0_box.IntVal)
		})
	})
	return cache_Data_List_Lazy_take
}

var cache_Data_List_Lazy_tail gopurs_runtime.Value
var once_Data_List_Lazy_tail sync.Once

func Get_Data_List_Lazy_tail() gopurs_runtime.Value {
	once_Data_List_Lazy_tail.Do(func() {
		cache_Data_List_Lazy_tail = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_tail(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_Lazy_tail
}

var cache_Data_List_Lazy_stripPrefix gopurs_runtime.Value
var once_Data_List_Lazy_stripPrefix sync.Once

func Get_Data_List_Lazy_stripPrefix() gopurs_runtime.Value {
	once_Data_List_Lazy_stripPrefix.Do(func() {
		cache_Data_List_Lazy_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_stripPrefix(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), v_1_box, s_2_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_Lazy_stripPrefix
}

var cache_Data_List_Lazy_span gopurs_runtime.Value
var once_Data_List_Lazy_span sync.Once

func Get_Data_List_Lazy_span() gopurs_runtime.Value {
	once_Data_List_Lazy_span.Do(func() {
		cache_Data_List_Lazy_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_Lazy_span(p_0_box, xs_1_box)
				_ = orig
				return gopurs_runtime.RecordDict2("init", "rest", orig.go__init, orig.rest)
			}()
		})
	})
	return cache_Data_List_Lazy_span
}

var cache_Data_List_Lazy_snoc gopurs_runtime.Value
var once_Data_List_Lazy_snoc sync.Once

func Get_Data_List_Lazy_snoc() gopurs_runtime.Value {
	once_Data_List_Lazy_snoc.Do(func() {
		cache_Data_List_Lazy_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_snoc(xs_0_box, x_1_box)
		})
	})
	return cache_Data_List_Lazy_snoc
}

var cache_Data_List_Lazy_singleton gopurs_runtime.Value
var once_Data_List_Lazy_singleton sync.Once

func Get_Data_List_Lazy_singleton() gopurs_runtime.Value {
	once_Data_List_Lazy_singleton.Do(func() {
		cache_Data_List_Lazy_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_singleton(a_0_box)
		})
	})
	return cache_Data_List_Lazy_singleton
}

var cache_Data_List_Lazy_showPattern gopurs_runtime.Value
var once_Data_List_Lazy_showPattern sync.Once

func Get_Data_List_Lazy_showPattern() gopurs_runtime.Value {
	once_Data_List_Lazy_showPattern.Do(func() {
		cache_Data_List_Lazy_showPattern = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_showPattern(dictShow_0_box)
		})
	})
	return cache_Data_List_Lazy_showPattern
}

var cache_Data_List_Lazy_scanlLazy gopurs_runtime.Value
var once_Data_List_Lazy_scanlLazy sync.Once

func Get_Data_List_Lazy_scanlLazy() gopurs_runtime.Value {
	once_Data_List_Lazy_scanlLazy.Do(func() {
		cache_Data_List_Lazy_scanlLazy = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_scanlLazy(f_0_box, acc_1_box, xs_2_box)
		})
	})
	return cache_Data_List_Lazy_scanlLazy
}

var cache_Data_List_Lazy_reverse gopurs_runtime.Value
var once_Data_List_Lazy_reverse sync.Once

func Get_Data_List_Lazy_reverse() gopurs_runtime.Value {
	once_Data_List_Lazy_reverse.Do(func() {
		cache_Data_List_Lazy_reverse = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_reverse(xs_0_box)
		})
	})
	return cache_Data_List_Lazy_reverse
}

var cache_Data_List_Lazy_replicateM gopurs_runtime.Value
var once_Data_List_Lazy_replicateM sync.Once

func Get_Data_List_Lazy_replicateM() gopurs_runtime.Value {
	once_Data_List_Lazy_replicateM.Do(func() {
		cache_Data_List_Lazy_replicateM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_replicateM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Data_List_Lazy_replicateM
}

var cache_Data_List_Lazy_repeat gopurs_runtime.Value
var once_Data_List_Lazy_repeat sync.Once

func Get_Data_List_Lazy_repeat() gopurs_runtime.Value {
	once_Data_List_Lazy_repeat.Do(func() {
		cache_Data_List_Lazy_repeat = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_repeat(x_0_box)
		})
	})
	return cache_Data_List_Lazy_repeat
}

var cache_Data_List_Lazy_replicate gopurs_runtime.Value
var once_Data_List_Lazy_replicate sync.Once

func Get_Data_List_Lazy_replicate() gopurs_runtime.Value {
	once_Data_List_Lazy_replicate.Do(func() {
		cache_Data_List_Lazy_replicate = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_replicate(i_0_box.IntVal, xs_1_box)
		})
	})
	return cache_Data_List_Lazy_replicate
}

var cache_Data_List_Lazy_go__range gopurs_runtime.Value
var once_Data_List_Lazy_go__range sync.Once

func Get_Data_List_Lazy_go__range() gopurs_runtime.Value {
	once_Data_List_Lazy_go__range.Do(func() {
		cache_Data_List_Lazy_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_go__range(start_0_box.IntVal, end_1_box.IntVal)
		})
	})
	return cache_Data_List_Lazy_go__range
}

var cache_Data_List_Lazy_partition gopurs_runtime.Value
var once_Data_List_Lazy_partition sync.Once

func Get_Data_List_Lazy_partition() gopurs_runtime.Value {
	once_Data_List_Lazy_partition.Do(func() {
		cache_Data_List_Lazy_partition = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_partition(f_0_box)
		})
	})
	return cache_Data_List_Lazy_partition
}

var cache_Data_List_Lazy_null gopurs_runtime.Value
var once_Data_List_Lazy_null sync.Once

func Get_Data_List_Lazy_null() gopurs_runtime.Value {
	once_Data_List_Lazy_null.Do(func() {
		cache_Data_List_Lazy_null = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isNothing(), Get_Data_List_Lazy_uncons())
	})
	return cache_Data_List_Lazy_null
}

var cache_Data_List_Lazy_nubBy gopurs_runtime.Value
var once_Data_List_Lazy_nubBy sync.Once

func Get_Data_List_Lazy_nubBy() gopurs_runtime.Value {
	once_Data_List_Lazy_nubBy.Do(func() {
		cache_Data_List_Lazy_nubBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_nubBy(p_0_box)
		})
	})
	return cache_Data_List_Lazy_nubBy
}

var cache_Data_List_Lazy_nub gopurs_runtime.Value
var once_Data_List_Lazy_nub sync.Once

func Get_Data_List_Lazy_nub() gopurs_runtime.Value {
	once_Data_List_Lazy_nub.Do(func() {
		cache_Data_List_Lazy_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_Lazy_nub
}

var cache_Data_List_Lazy_newtypePattern gopurs_runtime.Value
var once_Data_List_Lazy_newtypePattern sync.Once

func Get_Data_List_Lazy_newtypePattern() gopurs_runtime.Value {
	once_Data_List_Lazy_newtypePattern.Do(func() {
		cache_Data_List_Lazy_newtypePattern = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_List_Lazy_newtypePattern
}

var cache_Data_List_Lazy_mapMaybe gopurs_runtime.Value
var once_Data_List_Lazy_mapMaybe sync.Once

func Get_Data_List_Lazy_mapMaybe() gopurs_runtime.Value {
	once_Data_List_Lazy_mapMaybe.Do(func() {
		cache_Data_List_Lazy_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_mapMaybe(f_0_box)
		})
	})
	return cache_Data_List_Lazy_mapMaybe
}

var cache_Data_List_Lazy_some gopurs_runtime.Value
var once_Data_List_Lazy_some sync.Once

func Get_Data_List_Lazy_some() gopurs_runtime.Value {
	once_Data_List_Lazy_some.Do(func() {
		cache_Data_List_Lazy_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_1_box), v_2_box)
		})
	})
	return cache_Data_List_Lazy_some
}

var cache_Data_List_Lazy_many gopurs_runtime.Value
var once_Data_List_Lazy_many sync.Once

func Get_Data_List_Lazy_many() gopurs_runtime.Value {
	once_Data_List_Lazy_many.Do(func() {
		cache_Data_List_Lazy_many = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_1_box), v_2_box)
		})
	})
	return cache_Data_List_Lazy_many
}

var cache_Data_List_Lazy_length gopurs_runtime.Value
var once_Data_List_Lazy_length sync.Once

func Get_Data_List_Lazy_length() gopurs_runtime.Value {
	once_Data_List_Lazy_length.Do(func() {
		cache_Data_List_Lazy_length = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(l_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(l_0.IntVal, int64(1)))
		}), gopurs_runtime.Int(int64(0)))
	})
	return cache_Data_List_Lazy_length
}

var cache_Data_List_Lazy_last gopurs_runtime.Value
var once_Data_List_Lazy_last sync.Once

func Get_Data_List_Lazy_last() gopurs_runtime.Value {
	once_Data_List_Lazy_last.Do(func() {
		cache_Data_List_Lazy_last = func() gopurs_runtime.Value {
			var Call_local_Data_List_Lazy_go__go_0_0_7 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
			_ = Call_local_Data_List_Lazy_go__go_0_0_7
			var go__go_0_0_7 gopurs_runtime.Value
			_ = go__go_0_0_7
			Call_local_Data_List_Lazy_go__go_0_0_7 = func(v_1_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
			go__go_0_0_7:
				for {
					if false {
						continue go__go_0_0_7
					}
					var v_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_1_loop
					_ = v_1
					var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					{
						if v_1 != nil {
							var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
							{
								if (gopurs_runtime.Apply(Get_Data_List_Lazy_null(), (v_1).V1).IntVal) != (0) {
									__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 bool
										}{(v_1).V0, true}
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
								v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_1).V1))
								continue go__go_0_0_7
								__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
							}
						end_branch_1:
							__t2 = __t1
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
					}
				end_branch_2:
					return __t2
				}
			}
			go__go_0_0_7 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_0_0_7(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1_loop_val)))}
			})
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), go__go_0_0_7, Get_Data_List_Lazy_Types_step())
		}()
	})
	return cache_Data_List_Lazy_last
}

var cache_Data_List_Lazy_iterate gopurs_runtime.Value
var once_Data_List_Lazy_iterate sync.Once

func Get_Data_List_Lazy_iterate() gopurs_runtime.Value {
	once_Data_List_Lazy_iterate.Do(func() {
		cache_Data_List_Lazy_iterate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_iterate(f_0_box, x_1_box)
		})
	})
	return cache_Data_List_Lazy_iterate
}

var cache_Data_List_Lazy_insertAt gopurs_runtime.Value
var once_Data_List_Lazy_insertAt sync.Once

func Get_Data_List_Lazy_insertAt() gopurs_runtime.Value {
	once_Data_List_Lazy_insertAt.Do(func() {
		cache_Data_List_Lazy_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_insertAt(v_0_box.IntVal, v1_1_box, v2_2_box)
		})
	})
	return cache_Data_List_Lazy_insertAt
}

var cache_Data_List_Lazy_go__init gopurs_runtime.Value
var once_Data_List_Lazy_go__init sync.Once

func Get_Data_List_Lazy_go__init() gopurs_runtime.Value {
	once_Data_List_Lazy_go__init.Do(func() {
		cache_Data_List_Lazy_go__init = func() gopurs_runtime.Value {
			var go__go_0_0_9 gopurs_runtime.Value
			_ = go__go_0_0_9
			var go__go_0_0_9_cell *gopurs_runtime.Value
			_ = go__go_0_0_9_cell
			// FALLBACK TCO: isLoop=false len=1
			go__go_0_0_9 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t6 struct {
					V0 gopurs_runtime.Value
					V1 bool
				}
				{
					if v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil {
						var __t5 struct {
							V0 gopurs_runtime.Value
							V1 bool
						}
						{
							if (gopurs_runtime.Apply(Get_Data_List_Lazy_null(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1).IntVal) != (0) {
								__t5 = struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{Get_Data_List_Lazy_Types_nil(), true}
								goto end_branch_5
							} else {

							}
						}
						{
							// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope177)])])])
							__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply((*go__go_0_0_9_cell), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)))}))
							_ = __local_var_2_1
							var __t4 struct {
								V0 gopurs_runtime.Value
								V1 bool
							}
							{
								if __local_var_2_1 != nil {
									// TAST (Let): __local_var_3_2 shape=Other bindingType=Any
									__local_var_3_2 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
									_ = __local_var_3_2
									// TAST (Let): __local_var_4_3 shape=Other bindingType=Any
									__local_var_4_3 := (__local_var_2_1).V0
									_ = __local_var_4_3
									__t4 = struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_4_3}))}
									})), true}
									goto end_branch_4
								} else {

								}
							}
							{
								__t4 = struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
							}
						end_branch_4:
							__t5 = __t4
						}
					end_branch_5:
						__t6 = __t5
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
				}
			end_branch_6:
				return func() gopurs_runtime.Value {
					_v := __t6
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
			})
			go__go_0_0_9_cell = &go__go_0_0_9
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), go__go_0_0_9, Get_Data_List_Lazy_Types_step())
		}()
	})
	return cache_Data_List_Lazy_go__init
}

var cache_Data_List_Lazy_index gopurs_runtime.Value
var once_Data_List_Lazy_index sync.Once

func Get_Data_List_Lazy_index() gopurs_runtime.Value {
	once_Data_List_Lazy_index.Do(func() {
		cache_Data_List_Lazy_index = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_index(xs_0_box)
		})
	})
	return cache_Data_List_Lazy_index
}

var cache_Data_List_Lazy_head gopurs_runtime.Value
var once_Data_List_Lazy_head sync.Once

func Get_Data_List_Lazy_head() gopurs_runtime.Value {
	once_Data_List_Lazy_head.Do(func() {
		cache_Data_List_Lazy_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_head(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_Lazy_head
}

var cache_Data_List_Lazy_transpose gopurs_runtime.Value
var once_Data_List_Lazy_transpose sync.Once

func Get_Data_List_Lazy_transpose() gopurs_runtime.Value {
	once_Data_List_Lazy_transpose.Do(func() {
		cache_Data_List_Lazy_transpose = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_transpose(xs_0_box)
		})
	})
	return cache_Data_List_Lazy_transpose
}

var cache_Data_List_Lazy_groupBy gopurs_runtime.Value
var once_Data_List_Lazy_groupBy sync.Once

func Get_Data_List_Lazy_groupBy() gopurs_runtime.Value {
	once_Data_List_Lazy_groupBy.Do(func() {
		cache_Data_List_Lazy_groupBy = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_groupBy(eq_0_box)
		})
	})
	return cache_Data_List_Lazy_groupBy
}

var cache_Data_List_Lazy_group gopurs_runtime.Value
var once_Data_List_Lazy_group sync.Once

func Get_Data_List_Lazy_group() gopurs_runtime.Value {
	once_Data_List_Lazy_group.Do(func() {
		cache_Data_List_Lazy_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_Lazy_group
}

var cache_Data_List_Lazy_fromStep gopurs_runtime.Value
var once_Data_List_Lazy_fromStep sync.Once

func Get_Data_List_Lazy_fromStep() gopurs_runtime.Value {
	once_Data_List_Lazy_fromStep.Do(func() {
		cache_Data_List_Lazy_fromStep = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Lazy_applicativeLazy()).V1)
	})
	return cache_Data_List_Lazy_fromStep
}

var cache_Data_List_Lazy_insertBy gopurs_runtime.Value
var once_Data_List_Lazy_insertBy sync.Once

func Get_Data_List_Lazy_insertBy() gopurs_runtime.Value {
	once_Data_List_Lazy_insertBy.Do(func() {
		cache_Data_List_Lazy_insertBy = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_insertBy(cmp_0_box, x_1_box, xs_2_box)
		})
	})
	return cache_Data_List_Lazy_insertBy
}

var cache_Data_List_Lazy_insert gopurs_runtime.Value
var once_Data_List_Lazy_insert sync.Once

func Get_Data_List_Lazy_insert() gopurs_runtime.Value {
	once_Data_List_Lazy_insert.Do(func() {
		cache_Data_List_Lazy_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_Lazy_insert
}

var cache_Data_List_Lazy_fromFoldable gopurs_runtime.Value
var once_Data_List_Lazy_fromFoldable sync.Once

func Get_Data_List_Lazy_fromFoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_fromFoldable.Do(func() {
		cache_Data_List_Lazy_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_List_Lazy_fromFoldable
}

var cache_Data_List_Lazy_foldrLazy gopurs_runtime.Value
var once_Data_List_Lazy_foldrLazy sync.Once

func Get_Data_List_Lazy_foldrLazy() gopurs_runtime.Value {
	once_Data_List_Lazy_foldrLazy.Do(func() {
		cache_Data_List_Lazy_foldrLazy = gopurs_runtime.Func3(func(dictLazy_0_box gopurs_runtime.Value, op_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_foldrLazy(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_0_box), op_1_box, z_2_box)
		})
	})
	return cache_Data_List_Lazy_foldrLazy
}

var cache_Data_List_Lazy_foldM gopurs_runtime.Value
var once_Data_List_Lazy_foldM sync.Once

func Get_Data_List_Lazy_foldM() gopurs_runtime.Value {
	once_Data_List_Lazy_foldM.Do(func() {
		cache_Data_List_Lazy_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), f_1_box, b_2_box, xs_3_box)
		})
	})
	return cache_Data_List_Lazy_foldM
}

var cache_Data_List_Lazy_findIndex gopurs_runtime.Value
var once_Data_List_Lazy_findIndex sync.Once

func Get_Data_List_Lazy_findIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_findIndex.Do(func() {
		cache_Data_List_Lazy_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_findIndex(fn_0_box)
		})
	})
	return cache_Data_List_Lazy_findIndex
}

var cache_Data_List_Lazy_findLastIndex gopurs_runtime.Value
var once_Data_List_Lazy_findLastIndex sync.Once

func Get_Data_List_Lazy_findLastIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_findLastIndex.Do(func() {
		cache_Data_List_Lazy_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_findLastIndex(fn_0_box, xs_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_Lazy_findLastIndex
}

var cache_Data_List_Lazy_filterM gopurs_runtime.Value
var once_Data_List_Lazy_filterM sync.Once

func Get_Data_List_Lazy_filterM() gopurs_runtime.Value {
	once_Data_List_Lazy_filterM.Do(func() {
		cache_Data_List_Lazy_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Data_List_Lazy_filterM
}

var cache_Data_List_Lazy_filter gopurs_runtime.Value
var once_Data_List_Lazy_filter sync.Once

func Get_Data_List_Lazy_filter() gopurs_runtime.Value {
	once_Data_List_Lazy_filter.Do(func() {
		cache_Data_List_Lazy_filter = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_filter(p_0_box)
		})
	})
	return cache_Data_List_Lazy_filter
}

var cache_Data_List_Lazy_intersectBy gopurs_runtime.Value
var once_Data_List_Lazy_intersectBy sync.Once

func Get_Data_List_Lazy_intersectBy() gopurs_runtime.Value {
	once_Data_List_Lazy_intersectBy.Do(func() {
		cache_Data_List_Lazy_intersectBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_intersectBy(eq_0_box, xs_1_box, ys_2_box)
		})
	})
	return cache_Data_List_Lazy_intersectBy
}

var cache_Data_List_Lazy_intersect gopurs_runtime.Value
var once_Data_List_Lazy_intersect sync.Once

func Get_Data_List_Lazy_intersect() gopurs_runtime.Value {
	once_Data_List_Lazy_intersect.Do(func() {
		cache_Data_List_Lazy_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_Lazy_intersect
}

var cache_Data_List_Lazy_nubByEq gopurs_runtime.Value
var once_Data_List_Lazy_nubByEq sync.Once

func Get_Data_List_Lazy_nubByEq() gopurs_runtime.Value {
	once_Data_List_Lazy_nubByEq.Do(func() {
		cache_Data_List_Lazy_nubByEq = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_nubByEq(eq_0_box)
		})
	})
	return cache_Data_List_Lazy_nubByEq
}

var cache_Data_List_Lazy_nubEq gopurs_runtime.Value
var once_Data_List_Lazy_nubEq sync.Once

func Get_Data_List_Lazy_nubEq() gopurs_runtime.Value {
	once_Data_List_Lazy_nubEq.Do(func() {
		cache_Data_List_Lazy_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_Lazy_nubEq
}

var cache_Data_List_Lazy_eqPattern gopurs_runtime.Value
var once_Data_List_Lazy_eqPattern sync.Once

func Get_Data_List_Lazy_eqPattern() gopurs_runtime.Value {
	once_Data_List_Lazy_eqPattern.Do(func() {
		cache_Data_List_Lazy_eqPattern = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_eqPattern(dictEq_0_box)
		})
	})
	return cache_Data_List_Lazy_eqPattern
}

var cache_Data_List_Lazy_ordPattern gopurs_runtime.Value
var once_Data_List_Lazy_ordPattern sync.Once

func Get_Data_List_Lazy_ordPattern() gopurs_runtime.Value {
	once_Data_List_Lazy_ordPattern.Do(func() {
		cache_Data_List_Lazy_ordPattern = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_ordPattern(dictOrd_0_box)
		})
	})
	return cache_Data_List_Lazy_ordPattern
}

var cache_Data_List_Lazy_elemLastIndex gopurs_runtime.Value
var once_Data_List_Lazy_elemLastIndex sync.Once

func Get_Data_List_Lazy_elemLastIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_elemLastIndex.Do(func() {
		cache_Data_List_Lazy_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_List_Lazy_elemLastIndex
}

var cache_Data_List_Lazy_elemIndex gopurs_runtime.Value
var once_Data_List_Lazy_elemIndex sync.Once

func Get_Data_List_Lazy_elemIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_elemIndex.Do(func() {
		cache_Data_List_Lazy_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_List_Lazy_elemIndex
}

var cache_Data_List_Lazy_dropWhile gopurs_runtime.Value
var once_Data_List_Lazy_dropWhile sync.Once

func Get_Data_List_Lazy_dropWhile() gopurs_runtime.Value {
	once_Data_List_Lazy_dropWhile.Do(func() {
		cache_Data_List_Lazy_dropWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_dropWhile(p_0_box)
		})
	})
	return cache_Data_List_Lazy_dropWhile
}

var cache_Data_List_Lazy_drop gopurs_runtime.Value
var once_Data_List_Lazy_drop sync.Once

func Get_Data_List_Lazy_drop() gopurs_runtime.Value {
	once_Data_List_Lazy_drop.Do(func() {
		cache_Data_List_Lazy_drop = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_drop(n_0_box.IntVal)
		})
	})
	return cache_Data_List_Lazy_drop
}

var cache_Data_List_Lazy_slice gopurs_runtime.Value
var once_Data_List_Lazy_slice sync.Once

func Get_Data_List_Lazy_slice() gopurs_runtime.Value {
	once_Data_List_Lazy_slice.Do(func() {
		cache_Data_List_Lazy_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_slice(start_0_box.IntVal, end_1_box.IntVal, xs_2_box)
		})
	})
	return cache_Data_List_Lazy_slice
}

var cache_Data_List_Lazy_deleteBy gopurs_runtime.Value
var once_Data_List_Lazy_deleteBy sync.Once

func Get_Data_List_Lazy_deleteBy() gopurs_runtime.Value {
	once_Data_List_Lazy_deleteBy.Do(func() {
		cache_Data_List_Lazy_deleteBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_deleteBy(eq_0_box, x_1_box, xs_2_box)
		})
	})
	return cache_Data_List_Lazy_deleteBy
}

var cache_Data_List_Lazy_unionBy gopurs_runtime.Value
var once_Data_List_Lazy_unionBy sync.Once

func Get_Data_List_Lazy_unionBy() gopurs_runtime.Value {
	once_Data_List_Lazy_unionBy.Do(func() {
		cache_Data_List_Lazy_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_unionBy(eq_0_box, xs_1_box, ys_2_box)
		})
	})
	return cache_Data_List_Lazy_unionBy
}

var cache_Data_List_Lazy_union gopurs_runtime.Value
var once_Data_List_Lazy_union sync.Once

func Get_Data_List_Lazy_union() gopurs_runtime.Value {
	once_Data_List_Lazy_union.Do(func() {
		cache_Data_List_Lazy_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_Lazy_union
}

var cache_Data_List_Lazy_deleteAt gopurs_runtime.Value
var once_Data_List_Lazy_deleteAt sync.Once

func Get_Data_List_Lazy_deleteAt() gopurs_runtime.Value {
	once_Data_List_Lazy_deleteAt.Do(func() {
		cache_Data_List_Lazy_deleteAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_deleteAt(n_0_box.IntVal, xs_1_box)
		})
	})
	return cache_Data_List_Lazy_deleteAt
}

var cache_Data_List_Lazy_go__delete gopurs_runtime.Value
var once_Data_List_Lazy_go__delete sync.Once

func Get_Data_List_Lazy_go__delete() gopurs_runtime.Value {
	once_Data_List_Lazy_go__delete.Do(func() {
		cache_Data_List_Lazy_go__delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_Lazy_go__delete
}

var cache_Data_List_Lazy_difference gopurs_runtime.Value
var once_Data_List_Lazy_difference sync.Once

func Get_Data_List_Lazy_difference() gopurs_runtime.Value {
	once_Data_List_Lazy_difference.Do(func() {
		cache_Data_List_Lazy_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_Lazy_difference
}

var cache_Data_List_Lazy_cycle gopurs_runtime.Value
var once_Data_List_Lazy_cycle sync.Once

func Get_Data_List_Lazy_cycle() gopurs_runtime.Value {
	once_Data_List_Lazy_cycle.Do(func() {
		cache_Data_List_Lazy_cycle = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_cycle(xs_0_box)
		})
	})
	return cache_Data_List_Lazy_cycle
}

var cache_Data_List_Lazy_concatMap gopurs_runtime.Value
var once_Data_List_Lazy_concatMap sync.Once

func Get_Data_List_Lazy_concatMap() gopurs_runtime.Value {
	once_Data_List_Lazy_concatMap.Do(func() {
		cache_Data_List_Lazy_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_concatMap(b_0_box, a_1_box)
		})
	})
	return cache_Data_List_Lazy_concatMap
}

var cache_Data_List_Lazy_concat gopurs_runtime.Value
var once_Data_List_Lazy_concat sync.Once

func Get_Data_List_Lazy_concat() gopurs_runtime.Value {
	once_Data_List_Lazy_concat.Do(func() {
		cache_Data_List_Lazy_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_concat(v_0_box)
		})
	})
	return cache_Data_List_Lazy_concat
}

var cache_Data_List_Lazy_catMaybes gopurs_runtime.Value
var once_Data_List_Lazy_catMaybes sync.Once

func Get_Data_List_Lazy_catMaybes() gopurs_runtime.Value {
	once_Data_List_Lazy_catMaybes.Do(func() {
		cache_Data_List_Lazy_catMaybes = Call_Data_List_Lazy_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}))
	})
	return cache_Data_List_Lazy_catMaybes
}

var cache_Data_List_Lazy_alterAt gopurs_runtime.Value
var once_Data_List_Lazy_alterAt sync.Once

func Get_Data_List_Lazy_alterAt() gopurs_runtime.Value {
	once_Data_List_Lazy_alterAt.Do(func() {
		cache_Data_List_Lazy_alterAt = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_alterAt(n_0_box.IntVal, f_1_box, xs_2_box)
		})
	})
	return cache_Data_List_Lazy_alterAt
}

var cache_Data_List_Lazy_modifyAt gopurs_runtime.Value
var once_Data_List_Lazy_modifyAt sync.Once

func Get_Data_List_Lazy_modifyAt() gopurs_runtime.Value {
	once_Data_List_Lazy_modifyAt.Do(func() {
		cache_Data_List_Lazy_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_modifyAt(n_0_box.IntVal, f_1_box)
		})
	})
	return cache_Data_List_Lazy_modifyAt
}

func Call_Data_List_Lazy_tailRecM2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Lazy_4130553207_1542299734(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecMaybe())).V1, gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(o_3, "a"), gopurs_runtime.RecordGet(o_3, "b"))
	}), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			/*gopurs:record-v1:2:0061,0062*/
		}{a_1, b_2}
		_ = orig
		return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
	}())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_Lazy_fix(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var go__go_1_0_0 gopurs_runtime.Value
	_ = go__go_1_0_0
	var go__go_1_0_0_cell *gopurs_runtime.Value
	_ = go__go_1_0_0_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_1_0_0 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f_0, (*go__go_1_0_0_cell))
	})))
	go__go_1_0_0_cell = &go__go_1_0_0
	return go__go_1_0_0
}

func Call_Data_List_Lazy_fix1(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var go__go_1_0_1 gopurs_runtime.Value
	_ = go__go_1_0_1
	var go__go_1_0_1_cell *gopurs_runtime.Value
	_ = go__go_1_0_1_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_1_0_1 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f_0, (*go__go_1_0_1_cell))
	})))
	go__go_1_0_1_cell = &go__go_1_0_1
	return go__go_1_0_1
}

func Call_Data_List_Lazy_fix2(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var go__go_1_0_2 gopurs_runtime.Value
	_ = go__go_1_0_2
	var go__go_1_0_2_cell *gopurs_runtime.Value
	_ = go__go_1_0_2_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_1_0_2 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f_0, (*go__go_1_0_2_cell))
	})))
	go__go_1_0_2_cell = &go__go_1_0_2
	return go__go_1_0_2
}

func Call_Data_List_Lazy_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_Lazy_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_Lazy_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
zipWith:
	for {
		if false {
			continue zipWith
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var xs_1 gopurs_runtime.Value = xs_1_loop
		_ = xs_1
		var ys_2 gopurs_runtime.Value = ys_2_loop
		_ = ys_2
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_1)
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), ys_2)
		_ = __local_var_4_1
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_7_2 shape=App(Var) bindingType=(TypeVar a$scope106)
				__local_var_7_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
				_ = __local_var_7_2
				return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t3 gopurs_runtime.Value
					{
						if __local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 218341868 && __local_var_7_2.UnsafePtr == nil {
							__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
							goto end_branch_3
						} else {

						}
					}
					{
						if v1_8.Type == 9 && v1_8.IntVal == 218341868 && v1_8.UnsafePtr == nil {
							__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
							goto end_branch_3
						} else {

						}
					}
					{
						if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 218341868 && __local_var_7_2.UnsafePtr != nil) && (v1_8.Type == 9 && v1_8.IntVal == 218341868 && v1_8.UnsafePtr != nil) {
							__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_2.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_8.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_2.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_8.UnsafePtr).V1)}))}
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
					}
				end_branch_3:
					return __t3
				})
			})), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1))
		}))
	}
}

func Call_Data_List_Lazy_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value, ys_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var xs_2 gopurs_runtime.Value = xs_2_loop
	_ = xs_2
	var ys_3 gopurs_runtime.Value = ys_3_loop
	_ = ys_3
	// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=Any
	__local_var_4_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
	_ = __local_var_4_0
	// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=Any
	__local_var_5_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), ys_3)
	_ = __local_var_5_1
	return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_4
	}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_8_2 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_8_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_0)
			_ = __local_var_8_2
			return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if __local_var_8_2.Type == 9 && __local_var_8_2.IntVal == 218341868 && __local_var_8_2.UnsafePtr == nil {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
						goto end_branch_3
					} else {

					}
				}
				{
					if v1_9.Type == 9 && v1_9.IntVal == 218341868 && v1_9.UnsafePtr == nil {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
						goto end_branch_3
					} else {

					}
				}
				{
					if (__local_var_8_2.Type == 9 && __local_var_8_2.IntVal == 218341868 && __local_var_8_2.UnsafePtr != nil) && (v1_9.Type == 9 && v1_9.IntVal == 218341868 && v1_9.UnsafePtr != nil) {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_2.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_9.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_2.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_9.UnsafePtr).V1)}))}
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
				}
			end_branch_3:
				return __t3
			})
		})), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_1))
	})))
}

func Call_Data_List_Lazy_zip(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var ys_1 gopurs_runtime.Value = ys_1_loop
	_ = ys_1
	// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
	__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_0)
	_ = __local_var_2_0
	// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=Any
	__local_var_3_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), ys_1)
	_ = __local_var_3_1
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
			_ = __local_var_6_2
			return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if __local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
						goto end_branch_3
					} else {

					}
				}
				{
					if v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
						goto end_branch_3
					} else {

					}
				}
				{
					if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) && (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr != nil) {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
							}{(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V0}
							return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
						}(), Call_Data_List_Lazy_zipWith(Get_Data_Tuple_Tuple(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V1)}))}
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
				}
			end_branch_3:
				return __t3
			})
		})), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1))
	}))
}

func Call_Data_List_Lazy_updateAt(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateAt:
	for {
		if false {
			continue updateAt
		}
		var n_0 int64 = n_0_loop
		_ = n_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		var xs_2 gopurs_runtime.Value = xs_2_loop
		_ = xs_2
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
		_ = __local_var_3_0
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
			_ = __local_var_5_1
			var __t3 gopurs_runtime.Value
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil {
					var __t2 gopurs_runtime.Value
					{
						if (n_0) == (int64(0)) {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1}))}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_updateAt(gopurs_runtime.IntSub(n_0, int64(1)), x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)}))}
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_3:
			return __t3
		}))
	}
}

func Call_Data_List_Lazy_uncons(xs_0_loop gopurs_runtime.Value) struct {
	V0 struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V1 bool
} {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope43)])
	v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_0))
	_ = v_1_0
	var __t1 *Constructor_Data_Maybe_Just[struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]
	{
		if v_1_0 == nil {
			__t1 = Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						head gopurs_runtime.Value
						tail gopurs_runtime.Value
						/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
					}
					V1 bool
				}{struct {
					head gopurs_runtime.Value
					tail gopurs_runtime.Value
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_1
		} else {

		}
	}
	{
		if v_1_0 != nil {
			__t1 = Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						head gopurs_runtime.Value
						tail gopurs_runtime.Value
						/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
					}
					V1 bool
				}{struct {
					head gopurs_runtime.Value
					tail gopurs_runtime.Value
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{(v_1_0).V0, (v_1_0).V1}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Data_Maybe_Just[struct {
			head gopurs_runtime.Value
			tail gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}] {
			panic("Failed pattern match")
		}()
	}
end_branch_1:
	return func() struct {
		V0 struct {
			head gopurs_runtime.Value
			tail gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}
		V1 bool
	} {
		_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3833657837_3094389156(__t1))}
		if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 struct {
					head gopurs_runtime.Value
					tail gopurs_runtime.Value
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V1 bool
			}{V0: func() struct {
				head gopurs_runtime.Value
				tail gopurs_runtime.Value
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			} {
				orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					head gopurs_runtime.Value
					tail gopurs_runtime.Value
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}
				clone.head = gopurs_runtime.RecordGet(orig, "head")
				clone.tail = gopurs_runtime.RecordGet(orig, "tail")
				return clone
			}(), V1: true}
		}
		return struct {
			V0 struct {
				head gopurs_runtime.Value
				tail gopurs_runtime.Value
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V1 bool
		}{V0: struct {
			head gopurs_runtime.Value
			tail gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, V1: false}
	}()
}

func Call_Data_List_Lazy_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope49), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope49)])])] Empty))])
		__local_var_2_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_Lazy_uncons(xs_1)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = __local_var_2_0
		var __t1 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if __local_var_2_0 != nil {
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(__local_var_2_0).V0.head, (__local_var_2_0).V0.tail}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}(), true}
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
		}
	end_branch_1:
		return func() gopurs_runtime.Value {
			_v := __t1
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	}))
}

func Call_Data_List_Lazy_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
takeWhile:
	for {
		if false {
			continue takeWhile
		}
		var p_0 gopurs_runtime.Value = p_0_loop
		_ = p_0
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).IntVal) != (0)) {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_takeWhile(p_0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
			}
		end_branch_0:
			return __t0
		})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	}
}

func Call_Data_List_Lazy_take(n_0_loop int64) gopurs_runtime.Value {
take:
	for {
		if false {
			continue take
		}
		var n_0 int64 = n_0_loop
		_ = n_0
		var __t1 gopurs_runtime.Value
		{
			if (n_0) <= (int64(0)) {
				__t1 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_List_Lazy_Types_nil()
				})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr == nil {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
						goto end_branch_0
					} else {

					}
				}
				{
					if v1_1.Type == 9 && v1_1.IntVal == 218341868 && v1_1.UnsafePtr != nil {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_take(gopurs_runtime.IntSub(n_0, int64(1))), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)}))}
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
				}
			end_branch_0:
				return __t0
			})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_List_Lazy_tail(xs_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope69)])]), head: (TypeVar a$scope69)] Empty))])
	__local_var_1_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := Call_Data_List_Lazy_uncons(xs_0)
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))
	_ = __local_var_1_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_1_0 != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(__local_var_1_0).V0.tail, true}
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

func Call_Data_List_Lazy_stripPrefix(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Lazy_4130553207_1542299734(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecMaybe())).V1, gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v1_4_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope71)])
		v1_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.RecordGet(o_3, "a")))
		_ = v1_4_0
		var __t3 gopurs_runtime.Value
		{
			if v1_4_0 == nil {
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_3, "b")}))}}))}
				goto end_branch_3
			} else {

			}
		}
		{
			if v1_4_0 != nil {
				// TAST (Let): v2_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope71)])
				v2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.RecordGet(o_3, "b")))
				_ = v2_5_1
				var __t2 gopurs_runtime.Value
				{
					if (v2_5_1 != nil) && ((gopurs_runtime.Apply2(dictEq_0.V0, (v1_4_0).V0, (v2_5_1).V0).IntVal) != (0)) {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", (v1_4_0).V1, (v2_5_1).V1)}))}}))}
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}
			end_branch_2:
				__t3 = __t2
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
		}
	end_branch_3:
		return __t3
	}), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			/*gopurs:record-v1:2:0061,0062*/
		}{v_1, s_2}
		_ = orig
		return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
	}())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_Lazy_span(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) struct {
	go__init gopurs_runtime.Value
	rest     gopurs_runtime.Value
	/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
} {
span:
	for {
		if false {
			continue span
		}
		var p_0 gopurs_runtime.Value = p_0_loop
		_ = p_0
		var xs_1 gopurs_runtime.Value = xs_1_loop
		_ = xs_1
		// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope80), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope80)])])] Empty))])
		v_2_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_Lazy_uncons(xs_1)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_2_0
		var __t4 struct {
			go__init gopurs_runtime.Value
			rest     gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
		}
		{
			if (v_2_0 != nil) && ((gopurs_runtime.Apply(p_0, (v_2_0).V0.head).IntVal) != (0)) {
				// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
				__local_var_3_1 := (v_2_0).V0.head
				_ = __local_var_3_1
				// TAST (Let): v1_4_2 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope80)])]), rest: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope80)])])] Empty))
				v1_4_2 := Call_Data_List_Lazy_span(p_0, (v_2_0).V0.tail)
				_ = v1_4_2
				// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
				__local_var_5_3 := v1_4_2.go__init
				_ = __local_var_5_3
				__t4 = struct {
					go__init gopurs_runtime.Value
					rest     gopurs_runtime.Value
					/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
				}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_5_3}))}
				})), v1_4_2.rest}
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = struct {
				go__init gopurs_runtime.Value
				rest     gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			}{Get_Data_List_Lazy_Types_nil(), xs_1}
		}
	end_branch_4:
		return __t4
	}
}

func Call_Data_List_Lazy_snoc(xs_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2, Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, Get_Data_List_Lazy_Types_nil()}))}
	})), xs_0)
}

func Call_Data_List_Lazy_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()}))}
	}))
}

func Call_Data_List_Lazy_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Pattern ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_List_Lazy_Types_showList(dictShow_0), "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_List_Lazy_scanlLazy(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
scanlLazy:
	for {
		if false {
			continue scanlLazy
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var acc_1 gopurs_runtime.Value = acc_1_loop
		_ = acc_1
		var xs_2 gopurs_runtime.Value = xs_2_loop
		_ = xs_2
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
		_ = __local_var_3_0
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
			_ = __local_var_5_1
			var __t3 gopurs_runtime.Value
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil {
					// TAST (Let): acc_prime__6_2 shape=App(Other) bindingType=(TypeVar b$scope92)
					acc_prime__6_2 := gopurs_runtime.Apply2(f_0, acc_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0)
					_ = acc_prime__6_2
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, acc_prime__6_2, Call_Data_List_Lazy_scanlLazy(f_0, acc_prime__6_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)}))}
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_3:
			return __t3
		}))
	}
}

func Call_Data_List_Lazy_reverse(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_cons(a_3, b_2)
		}), Get_Data_List_Lazy_Types_nil(), xs_0)
	})))
}

func Call_Data_List_Lazy_replicateM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
replicateM:
	for {
		if false {
			continue replicateM
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope102)])
		Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
		_ = Bind1_2_1
		return gopurs_runtime.Func2(func(n_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if (n_3.IntVal) < (int64(1)) {
					__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), Get_Data_List_Lazy_Types_nil())
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.Apply2(Bind1_2_1.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_Data_List_Lazy_replicateM(dictMonad_0), gopurs_runtime.Int(gopurs_runtime.IntSub(n_3.IntVal, int64(1))), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_5, as_6}))}
						})))
					}))
				}))
			}
		end_branch_2:
			return __t2
		})
	}
}

func Call_Data_List_Lazy_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var go__go_1_0_3 gopurs_runtime.Value
	_ = go__go_1_0_3
	var go__go_1_0_3_cell *gopurs_runtime.Value
	_ = go__go_1_0_3_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_1_0_3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, (*go__go_1_0_3_cell)}))}
		}))
	})))
	go__go_1_0_3_cell = &go__go_1_0_3
	return go__go_1_0_3
}

func Call_Data_List_Lazy_replicate(i_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	var xs_1 gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return gopurs_runtime.Apply(Call_Data_List_Lazy_take(i_0), Call_Data_List_Lazy_repeat(xs_1))
}

func Call_Data_List_Lazy_go__range(start_0_loop int64, end_1_loop int64) gopurs_runtime.Value {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	var __t2 gopurs_runtime.Value
	{
		if (start_0) > (end_1) {
			__t2 = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldableList()).V1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if (x_2.IntVal) >= (end_1) {
						__t1 = func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{x_2, gopurs_runtime.Int(gopurs_runtime.IntSub(x_2.IntVal, int64(1)))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}(), true}
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
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}
			end_branch_1:
				return __t1
			}), gopurs_runtime.Int(start_0))
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldableList()).V1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if (x_2.IntVal) <= (end_1) {
					__t0 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
							}{x_2, gopurs_runtime.Int(gopurs_runtime.IntAdd(x_2.IntVal, int64(1)))}
							return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
						}(), true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Int(start_0))
	}
end_branch_2:
	return __t2
}

func Call_Data_List_Lazy_partition(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if (gopurs_runtime.Apply(f_0, x_1).IntVal) != (0) {
				__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordGet(v_2, "no"), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.RecordGet(v_2, "yes")}))}
				})))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.RecordGet(v_2, "no")}))}
			})), gopurs_runtime.RecordGet(v_2, "yes"))
		}
	end_branch_0:
		return __t0
	}), func() gopurs_runtime.Value {
		orig := struct {
			no  gopurs_runtime.Value
			yes gopurs_runtime.Value
			/*gopurs:record-v1:2:006e006f,007900650073*/
		}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
		_ = orig
		return gopurs_runtime.RecordDict2("no", "yes", orig.no, orig.yes)
	}())
}

func Call_Data_List_Lazy_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var goStep_1_0_4 gopurs_runtime.Value
	_ = goStep_1_0_4
	var goStep_1_0_4_cell *gopurs_runtime.Value
	_ = goStep_1_0_4_cell
	// FALLBACK TCO: isLoop=false len=2
	var go__go_1_1_5 gopurs_runtime.Value
	_ = go__go_1_1_5
	var go__go_1_1_5_cell *gopurs_runtime.Value
	_ = go__go_1_1_5_cell
	// FALLBACK TCO: isLoop=false len=2
	goStep_1_0_4 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t4 gopurs_runtime.Value
		{
			if v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
				goto end_branch_4
			} else {

			}
		}
		{
			if v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil {
				// TAST (Let): v2_4_2 shape=App(Var) bindingType=(Record (Row [found: Boolean, result: (ADT ["Data","List","Internal","Set"] [(TypeVar a$scope127)])] Empty))
				v2_4_2 := Call_Data_List_Internal_insertAndLookupBy(p_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2)
				_ = v2_4_2
				var __t3 gopurs_runtime.Value
				{
					if v2_4_2.found {
						__t3 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply2((*go__go_1_1_5_cell), v2_4_2.result, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1))
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2((*go__go_1_1_5_cell), v2_4_2.result, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
				}
			end_branch_3:
				__t4 = __t3
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
		}
	end_branch_4:
		return __t4
	})
	goStep_1_0_4_cell = &goStep_1_0_4
	go__go_1_1_5 = gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=(Func [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope127)])] (ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope127)]))
		__local_var_4_5 := gopurs_runtime.Apply((*goStep_1_0_4_cell), s_2)
		_ = __local_var_4_5
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
		}))
	})
	go__go_1_1_5_cell = &go__go_1_1_5
	return gopurs_runtime.Apply(go__go_1_1_5, Get_Data_List_Internal_emptySet())
}

func Call_Data_List_Lazy_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_List_Lazy_nubBy(Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_List_Lazy_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
	for {
		if false {
			continue mapMaybe
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var Call_local_Data_List_Lazy_go__go_1_0_6 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
		_ = Call_local_Data_List_Lazy_go__go_1_0_6
		var go__go_1_0_6 gopurs_runtime.Value
		_ = go__go_1_0_6
		Call_local_Data_List_Lazy_go__go_1_0_6 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
		go__go_1_0_6:
			for {
				if false {
					continue go__go_1_0_6
				}
				var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
				_ = v_2
				var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
				{
					if v_2 == nil {
						__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
						goto end_branch_3
					} else {

					}
				}
				{
					if v_2 != nil {
						// TAST (Let): v1_3_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope141)])
						v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (v_2).V0))
						_ = v1_3_1
						var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
						{
							if v1_3_1 == nil {
								v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
								continue go__go_1_0_6
								__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
								goto end_branch_2
							} else {

							}
						}
						{
							if v1_3_1 != nil {
								__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (v_2).V1)})
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
						}
					end_branch_2:
						__t3 = __t2
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			}
		}
		go__go_1_0_6 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_6(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
		})
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, go__go_1_0_6), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	}
}

func Call_Data_List_Lazy_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], dictLazy_1_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var dictLazy_1 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_1_loop
	_ = dictLazy_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Lazy_Types_cons(), v_2), gopurs_runtime.Apply(dictLazy_1.V0, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_List_Lazy_many(dictAlternative_0, dictLazy_1, v_2)
	})))
}

func Call_Data_List_Lazy_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], dictLazy_1_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var dictLazy_1 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_1_loop
	_ = dictLazy_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_Data_List_Lazy_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "pure"), Get_Data_List_Lazy_Types_nil()))
}

func Call_Data_List_Lazy_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var go__go_2_0_8 gopurs_runtime.Value
	_ = go__go_2_0_8
	var go__go_2_0_8_cell *gopurs_runtime.Value
	_ = go__go_2_0_8_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_2_0_8 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0, f_0, (*go__go_2_0_8_cell))
		_ = __local_var_4_1
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, __local_var_4_1}))}
		}))
	})))
	go__go_2_0_8_cell = &go__go_2_0_8
	return go__go_2_0_8
}

func Call_Data_List_Lazy_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertAt:
	for {
		if false {
			continue insertAt
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t3 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, v1_1, v2_2}))}
				}))
				goto end_branch_3
			} else {

			}
		}
		{
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), v2_2)
			_ = __local_var_3_0
			__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
				__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
				_ = __local_var_5_1
				var __t2 gopurs_runtime.Value
				{
					if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, v1_1, Get_Data_List_Lazy_Types_nil()}))}
						goto end_branch_2
					} else {

					}
				}
				{
					if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_insertAt(gopurs_runtime.IntSub(v_0, int64(1)), v1_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)}))}
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
				}
			end_branch_2:
				return __t2
			}))
		}
	end_branch_3:
		return __t3
	}
}

func Call_Data_List_Lazy_index(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var Call_local_Data_List_Lazy_go__go_1_0_10 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], int64) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	_ = Call_local_Data_List_Lazy_go__go_1_0_10
	var go__go_1_0_10 gopurs_runtime.Value
	_ = go__go_1_0_10
	Call_local_Data_List_Lazy_go__go_1_0_10 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_3_loop int64) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	go__go_1_0_10:
		for {
			if false {
				continue go__go_1_0_10
			}
			var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
			_ = v_2
			var v1_3 int64 = v1_3_loop
			_ = v1_3
			var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
			{
				if v_2 == nil {
					__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
					goto end_branch_2
				} else {

				}
			}
			{
				if v_2 != nil {
					var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					{
						if (v1_3) == (int64(0)) {
							__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{(v_2).V0, true}
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
						v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
						v1_3_loop = gopurs_runtime.IntSub(v1_3, int64(1))
						continue go__go_1_0_10
						__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
					}
				end_branch_1:
					__t2 = __t1
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}
	}
	go__go_1_0_10 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_10(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), v1_3_loop_val.IntVal))}
		})
	})
	return gopurs_runtime.Apply(go__go_1_0_10, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_0)))})
}

func Call_Data_List_Lazy_head(xs_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope186), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope186)])])] Empty))])
	__local_var_1_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := Call_Data_List_Lazy_uncons(xs_0)
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))
	_ = __local_var_1_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_1_0 != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{(__local_var_1_0).V0.head, true}
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

func Call_Data_List_Lazy_transpose(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
transpose:
	for {
		if false {
			continue transpose
		}
		var xs_0 gopurs_runtime.Value = xs_0_loop
		_ = xs_0
		// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope188)])]), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope188)])])])])] Empty))])
		v_1_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_Lazy_uncons(xs_0)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_1_0
		var __t8 gopurs_runtime.Value
		{
			if v_1_0 == nil {
				__t8 = xs_0
				goto end_branch_8
			} else {

			}
		}
		{
			if v_1_0 != nil {
				// TAST (Let): v1_2_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope188), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope188)])])] Empty))])
				v1_2_1 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := Call_Data_List_Lazy_uncons((v_1_0).V0.head)
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
							orig := _v.V0
							_ = orig
							return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
						}()})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))
				_ = v1_2_1
				var __t7 gopurs_runtime.Value
				{
					if v1_2_1 == nil {
						xs_0_loop = (v_1_0).V0.tail
						continue transpose
						__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
						goto end_branch_7
					} else {

					}
				}
				{
					if v1_2_1 != nil {
						// TAST (Let): __local_var_3_2 shape=Other bindingType=Any
						__local_var_3_2 := (v1_2_1).V0.head
						_ = __local_var_3_2
						// TAST (Let): __local_var_4_3 shape=Other bindingType=Any
						__local_var_4_3 := (v1_2_1).V0.tail
						_ = __local_var_4_3
						// TAST (Let): __local_var_5_4 shape=App(Var) bindingType=Any
						__local_var_5_4 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_head()), (v_1_0).V0.tail)
						_ = __local_var_5_4
						// TAST (Let): __local_var_6_6 shape=App(Var) bindingType=Any
						__local_var_6_6 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_tail()), (v_1_0).V0.tail)
						_ = __local_var_6_6
						// TAST (Let): __local_var_6_5 shape=App(Var) bindingType=Any
						__local_var_6_5 := Call_Data_List_Lazy_transpose(gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_3, __local_var_6_6}))}
						})))
						_ = __local_var_6_5
						__t7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_5_4}))}
							})), __local_var_6_5}))}
						}))
						goto end_branch_7
					} else {

					}
				}
				{
					__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_7:
				__t8 = __t7
				goto end_branch_8
			} else {

			}
		}
		{
			__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_8:
		return __t8
	}
}

func Call_Data_List_Lazy_groupBy(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
groupBy:
	for {
		if false {
			continue groupBy
		}
		var eq_0 gopurs_runtime.Value = eq_0_loop
		_ = eq_0
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				if v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_3
				} else {

				}
			}
			{
				if v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil {
					// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
					__local_var_2_0 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
					_ = __local_var_2_0
					// TAST (Let): v1_3_1 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope192)])]), rest: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope192)])])] Empty))
					v1_3_1 := Call_Data_List_Lazy_span(gopurs_runtime.Apply(eq_0, __local_var_2_0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
					_ = v1_3_1
					// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
					__local_var_4_2 := v1_3_1.go__init
					_ = __local_var_4_2
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_0, __local_var_4_2}))}
					})), gopurs_runtime.Apply(Call_Data_List_Lazy_groupBy(eq_0), v1_3_1.rest)}))}
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_3:
			return __t3
		})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	}
}

func Call_Data_List_Lazy_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return Call_Data_List_Lazy_groupBy(Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_Lazy_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertBy:
	for {
		if false {
			continue insertBy
		}
		var cmp_0 gopurs_runtime.Value = cmp_0_loop
		_ = cmp_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		var xs_2 gopurs_runtime.Value = xs_2_loop
		_ = xs_2
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
		_ = __local_var_3_0
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
			_ = __local_var_5_1
			var __t4 gopurs_runtime.Value
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, Get_Data_List_Lazy_Types_nil()}))}
					goto end_branch_4
				} else {

				}
			}
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil {
					var __t3 gopurs_runtime.Value
					{
						var __t_tag_2 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0).IntVal)
						_ = __t_tag_2
						if uint32(__t_tag_2) == 380165415 {
							__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)}))}
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply(Get_Data_List_Lazy_fromStep(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__local_var_5_1))})}))}
					}
				end_branch_3:
					__t4 = __t3
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_4:
			return __t4
		}))
	}
}

func Call_Data_List_Lazy_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope217), (TypeVar a$scope217)] (ADT ["Data","Ordering","Ordering"] []))
	__local_var_1_0 := Call_Data_Ord_compare(dictOrd_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_3)
		_ = __local_var_4_1
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
			_ = __local_var_6_2
			var __t5 gopurs_runtime.Value
			{
				if __local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil {
					__t5 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_2, Get_Data_List_Lazy_Types_nil()}))}
					goto end_branch_5
				} else {

				}
			}
			{
				if __local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil {
					var __t4 gopurs_runtime.Value
					{
						var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0)
						_ = __t_tag_3
						if uint32(__t_tag_3.IntVal) == 380165415 {
							__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)}))}
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.Apply(Get_Data_List_Lazy_fromStep(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__local_var_6_2))})}))}
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_5:
			return __t5
		}))
	})
}

func Call_Data_List_Lazy_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(dictFoldable_0.V2, Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil())
}

func Call_Data_List_Lazy_foldrLazy(dictLazy_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], op_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictLazy_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
	_ = dictLazy_0
	var op_1 gopurs_runtime.Value = op_1_loop
	_ = op_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var go__go_3_0_11 gopurs_runtime.Value
	_ = go__go_3_0_11
	var go__go_3_0_11_cell *gopurs_runtime.Value
	_ = go__go_3_0_11_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_3_0_11 = gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope224)])
		v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
		_ = v_5_1
		var __t4 gopurs_runtime.Value
		{
			if v_5_1 != nil {
				// TAST (Let): __local_var_6_2 shape=Other bindingType=Any
				__local_var_6_2 := (v_5_1).V0
				_ = __local_var_6_2
				// TAST (Let): __local_var_7_3 shape=Other bindingType=Any
				__local_var_7_3 := (v_5_1).V1
				_ = __local_var_7_3
				__t4 = gopurs_runtime.Apply(dictLazy_0.V0, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(op_1, __local_var_6_2, gopurs_runtime.Apply((*go__go_3_0_11_cell), __local_var_7_3))
				}))
				goto end_branch_4
			} else {

			}
		}
		{
			if v_5_1 == nil {
				__t4 = z_2
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_4:
		return __t4
	})
	go__go_3_0_11_cell = &go__go_3_0_11
	return go__go_3_0_11
}

func Call_Data_List_Lazy_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
		var xs_3 gopurs_runtime.Value = xs_3_loop
		_ = xs_3
		// TAST (Let): v_4_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope230), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope230)])])] Empty))])
		v_4_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_List_Lazy_uncons(xs_3)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_4_0
		var __t2 gopurs_runtime.Value
		{
			if v_4_0 == nil {
				__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), b_2)
				goto end_branch_2
			} else {

			}
		}
		{
			if v_4_0 != nil {
				// TAST (Let): __local_var_5_1 shape=Other bindingType=Any
				__local_var_5_1 := (v_4_0).V0.tail
				_ = __local_var_5_1
				__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, (v_4_0).V0.head), gopurs_runtime.Func(func(b_prime__6 gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_Data_List_Lazy_foldM(dictMonad_0, f_1, b_prime__6, __local_var_5_1)
				}))
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

func Call_Data_List_Lazy_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var fn_0 gopurs_runtime.Value = fn_0_loop
	_ = fn_0
	var Call_local_Data_List_Lazy_go__2988133926_1_0_12 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_Lazy_go__2988133926_1_0_12
	var go__2988133926_1_0_12 gopurs_runtime.Value
	_ = go__2988133926_1_0_12
	var Call_local_Data_List_Lazy_go__go_1_1_13 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_Lazy_go__go_1_1_13
	var go__go_1_1_13 gopurs_runtime.Value
	_ = go__go_1_1_13
	Call_local_Data_List_Lazy_go__2988133926_1_0_12 = func(n_2_loop int64, list_3_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[int64] {
	go__2988133926_1_0_12:
		for {
			if false {
				continue go__2988133926_1_0_12
			}
			var n_2 int64 = n_2_loop
			_ = n_2
			var list_3 gopurs_runtime.Value = list_3_loop
			_ = list_3
			// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=Any
			__local_var_4_2 := Call_Data_List_Lazy_uncons(list_3)
			_ = __local_var_4_2
			var __t4 *Constructor_Data_Maybe_Just[int64]
			{
				if __local_var_4_2.V1 {
					var __t3 *Constructor_Data_Maybe_Just[int64]
					{
						if (gopurs_runtime.Apply(fn_0, __local_var_4_2.V0.head).IntVal) != (0) {
							__t3 = Rebox_Data_List_Lazy_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Int(n_2), true}
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
						n_2_loop = gopurs_runtime.IntAdd(n_2, int64(1))
						list_3_loop = __local_var_4_2.V0.tail
						continue go__2988133926_1_0_12
						__t3 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
					}
				end_branch_3:
					__t4 = __t3
					goto end_branch_4
				} else {

				}
			}
			{
				if !__local_var_4_2.V1 {
					__t4 = Rebox_Data_List_Lazy_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
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
				__t4 = Rebox_Data_List_Lazy_3094389156_1170268447(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())
			}
		end_branch_4:
			return __t4
		}
	}
	go__2988133926_1_0_12 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(list_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1170268447_3094389156(Call_local_Data_List_Lazy_go__2988133926_1_0_12(n_2_loop_val.IntVal, list_3_loop_val)))}
		})
	})
	Call_local_Data_List_Lazy_go__go_1_1_13 = func(n_2_loop int64, list_3_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[int64] {
	go__go_1_1_13:
		for {
			if false {
				continue go__go_1_1_13
			}
			var n_2 int64 = n_2_loop
			_ = n_2
			var list_3 gopurs_runtime.Value = list_3_loop
			_ = list_3
			// TAST (Let): __local_var_4_5 shape=App(Var) bindingType=Any
			__local_var_4_5 := Call_Data_List_Lazy_uncons(list_3)
			_ = __local_var_4_5
			var __t7 *Constructor_Data_Maybe_Just[int64]
			{
				if __local_var_4_5.V1 {
					var __t6 *Constructor_Data_Maybe_Just[int64]
					{
						if (gopurs_runtime.Apply(fn_0, __local_var_4_5.V0.head).IntVal) != (0) {
							__t6 = Rebox_Data_List_Lazy_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Int(n_2), true}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}()))
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = Call_local_Data_List_Lazy_go__2988133926_1_0_12(gopurs_runtime.IntAdd(n_2, int64(1)), __local_var_4_5.V0.tail)
					}
				end_branch_6:
					__t7 = __t6
					goto end_branch_7
				} else {

				}
			}
			{
				if !__local_var_4_5.V1 {
					__t7 = Rebox_Data_List_Lazy_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = Rebox_Data_List_Lazy_3094389156_1170268447(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())
			}
		end_branch_7:
			return __t7
		}
	}
	go__go_1_1_13 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(list_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1170268447_3094389156(Call_local_Data_List_Lazy_go__go_1_1_13(n_2_loop_val.IntVal, list_3_loop_val)))}
		})
	})
	return gopurs_runtime.Apply(go__2988133926_1_0_12, gopurs_runtime.Int(int64(0)))
}

func Call_Data_List_Lazy_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var fn_0 gopurs_runtime.Value = fn_0_loop
	_ = fn_0
	var xs_1 gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
	__local_var_2_0 := gopurs_runtime.Apply(Call_Data_List_Lazy_findIndex(fn_0), Call_Data_List_Lazy_reverse(xs_1))
	_ = __local_var_2_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_List_Lazy_length(), xs_1).IntVal, int64(1)), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0.IntVal)), true}
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

func Call_Data_List_Lazy_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
filterM:
	for {
		if false {
			continue filterM
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope242)])
		Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
		_ = Bind1_2_1
		return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, list_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_5_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope241), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope241)])])] Empty))])
			v_5_2 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(list_4)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			_ = v_5_2
			var __t6 gopurs_runtime.Value
			{
				if v_5_2 == nil {
					__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), Get_Data_List_Lazy_Types_nil())
					goto end_branch_6
				} else {

				}
			}
			{
				if v_5_2 != nil {
					// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
					__local_var_6_3 := (v_5_2).V0.head
					_ = __local_var_6_3
					// TAST (Let): __local_var_7_4 shape=Other bindingType=Any
					__local_var_7_4 := (v_5_2).V0.tail
					_ = __local_var_7_4
					__t6 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(p_3, __local_var_6_3), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_Data_List_Lazy_filterM(dictMonad_0), p_3, __local_var_7_4), gopurs_runtime.Func(func(xs_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
							var __t5 gopurs_runtime.Value
							{
								if (b_8.IntVal) != (0) {
									__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_3, xs_prime__9}))}
									}))
									goto end_branch_5
								} else {

								}
							}
							{
								__t5 = xs_prime__9
							}
						end_branch_5:
							return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), __t5)
						}))
					}))
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_6:
			return __t6
		})
	}
}

func Call_Data_List_Lazy_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
	for {
		if false {
			continue filter
		}
		var p_0 gopurs_runtime.Value = p_0_loop
		_ = p_0
		var Call_local_Data_List_Lazy_go__go_1_0_14 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
		_ = Call_local_Data_List_Lazy_go__go_1_0_14
		var go__go_1_0_14 gopurs_runtime.Value
		_ = go__go_1_0_14
		Call_local_Data_List_Lazy_go__go_1_0_14 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
		go__go_1_0_14:
			for {
				if false {
					continue go__go_1_0_14
				}
				var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
				_ = v_2
				var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
				{
					if v_2 == nil {
						__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
						goto end_branch_2
					} else {

					}
				}
				{
					if v_2 != nil {
						var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
						{
							if (gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0) {
								__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_filter(p_0), (v_2).V1)})
								goto end_branch_1
							} else {

							}
						}
						{
							v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
							continue go__go_1_0_14
							__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
						}
					end_branch_1:
						__t2 = __t1
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
				}
			end_branch_2:
				return __t2
			}
		}
		go__go_1_0_14 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_14(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
		})
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, go__go_1_0_14), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	}
}

func Call_Data_List_Lazy_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var eq_0 gopurs_runtime.Value = eq_0_loop
	_ = eq_0
	var xs_1 gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	var ys_2 gopurs_runtime.Value = ys_2_loop
	_ = ys_2
	return gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Data_Foldable_any(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList())), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3591112874_2663347022(Rebox_Data_List_Lazy_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))}, gopurs_runtime.Apply(eq_0, x_3), ys_2)
	})), xs_1)
}

func Call_Data_List_Lazy_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_List_Lazy_intersectBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_Lazy_nubByEq(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
nubByEq:
	for {
		if false {
			continue nubByEq
		}
		var eq_0 gopurs_runtime.Value = eq_0_loop
		_ = eq_0
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_1
				} else {

				}
			}
			{
				if v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil {
					// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
					__local_var_2_0 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
					_ = __local_var_2_0
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_0, gopurs_runtime.Apply(Call_Data_List_Lazy_nubByEq(eq_0), gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(eq_0, __local_var_2_0, y_3).IntVal) != (0)) != (true))
					})), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))}))}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_1:
			return __t1
		})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	}
}

func Call_Data_List_Lazy_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return Call_Data_List_Lazy_nubByEq(Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_Lazy_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Lazy_Types_eqList(dictEq_0), "eq"), x_1, y_2)
	})}))}
}

func Call_Data_List_Lazy_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqPattern1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope121)])])])
	eqPattern1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Lazy_eqPattern(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
	_ = eqPattern1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqPattern1_1_0)}
	}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Lazy_Types_ordList(dictOrd_0), "compare"), x_2, y_3)
	})}))}
}

func Call_Data_List_Lazy_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Data_List_Lazy_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1)
	}))
}

func Call_Data_List_Lazy_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_List_Lazy_findIndex(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1)
	}))
}

func Call_Data_List_Lazy_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var Call_local_Data_List_Lazy_go__go_1_0_15 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
	_ = Call_local_Data_List_Lazy_go__go_1_0_15
	var go__go_1_0_15 gopurs_runtime.Value
	_ = go__go_1_0_15
	Call_local_Data_List_Lazy_go__go_1_0_15 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	go__go_1_0_15:
		for {
			if false {
				continue go__go_1_0_15
			}
			var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
			_ = v_2
			var __t1 gopurs_runtime.Value
			{
				if (v_2 != nil) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
					v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
					continue go__go_1_0_15
					__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Apply(Get_Data_List_Lazy_fromStep(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2)})
			}
		end_branch_1:
			return __t1
		}
	}
	go__go_1_0_15 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Data_List_Lazy_go__go_1_0_15(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val))
	})
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), go__go_1_0_15, Get_Data_List_Lazy_Types_step())
}

func Call_Data_List_Lazy_drop(n_0_loop int64) gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	var Call_local_Data_List_Lazy_go__2018343127_1_0_16 func(int64, *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_Lazy_go__2018343127_1_0_16
	var go__2018343127_1_0_16 gopurs_runtime.Value
	_ = go__2018343127_1_0_16
	var Call_local_Data_List_Lazy_go__go_1_1_17 func(int64, *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_Lazy_go__go_1_1_17
	var go__go_1_1_17 gopurs_runtime.Value
	_ = go__go_1_1_17
	Call_local_Data_List_Lazy_go__2018343127_1_0_16 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
	go__2018343127_1_0_16:
		for {
			if false {
				continue go__2018343127_1_0_16
			}
			var v_2 int64 = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
			{
				if (v_2) == (int64(0)) {
					__t2 = v1_3
					goto end_branch_2
				} else {

				}
			}
			{
				if v1_3 == nil {
					__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
					goto end_branch_2
				} else {

				}
			}
			{
				if v1_3 != nil {
					v_2_loop = gopurs_runtime.IntSub(v_2, int64(1))
					v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v1_3).V1))
					continue go__2018343127_1_0_16
					__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}
	}
	go__2018343127_1_0_16 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__2018343127_1_0_16(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
		})
	})
	Call_local_Data_List_Lazy_go__go_1_1_17 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
	go__go_1_1_17:
		for {
			if false {
				continue go__go_1_1_17
			}
			var v_2 int64 = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
			{
				if (v_2) == (int64(0)) {
					__t3 = v1_3
					goto end_branch_3
				} else {

				}
			}
			{
				if v1_3 == nil {
					__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
					goto end_branch_3
				} else {

				}
			}
			{
				if v1_3 != nil {
					__t3 = Call_local_Data_List_Lazy_go__2018343127_1_0_16(gopurs_runtime.IntSub(v_2, int64(1)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v1_3).V1)))
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		}
	}
	go__go_1_1_17 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_1_17(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
		})
	})
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Apply(go__2018343127_1_0_16, gopurs_runtime.Int(n_0))), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
}

func Call_Data_List_Lazy_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	var xs_2 gopurs_runtime.Value = xs_2_loop
	_ = xs_2
	return gopurs_runtime.Apply(Call_Data_List_Lazy_take(gopurs_runtime.IntSub(end_1, start_0)), gopurs_runtime.Apply(Call_Data_List_Lazy_drop(start_0), xs_2))
}

func Call_Data_List_Lazy_deleteBy(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteBy:
	for {
		if false {
			continue deleteBy
		}
		var eq_0 gopurs_runtime.Value = eq_0_loop
		_ = eq_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		var xs_2 gopurs_runtime.Value = xs_2_loop
		_ = xs_2
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
		_ = __local_var_3_0
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
			_ = __local_var_5_1
			var __t3 gopurs_runtime.Value
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil {
					var __t2 gopurs_runtime.Value
					{
						if (gopurs_runtime.Apply2(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0).IntVal) != (0) {
							__t2 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)}))}
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_3:
			return __t3
		}))
	}
}

func Call_Data_List_Lazy_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var eq_0 gopurs_runtime.Value = eq_0_loop
	_ = eq_0
	var xs_1 gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	var ys_2 gopurs_runtime.Value = ys_2_loop
	_ = ys_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), xs_1, gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_5_0 shape=App(Var) bindingType=Any
		__local_var_5_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), b_3)
		_ = __local_var_5_0
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_7_1 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_7_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_0)
			_ = __local_var_7_1
			var __t3 gopurs_runtime.Value
			{
				if __local_var_7_1.Type == 9 && __local_var_7_1.IntVal == 218341868 && __local_var_7_1.UnsafePtr == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_7_1.Type == 9 && __local_var_7_1.IntVal == 218341868 && __local_var_7_1.UnsafePtr != nil {
					var __t2 gopurs_runtime.Value
					{
						if (gopurs_runtime.Apply2(eq_0, a_4, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_1.UnsafePtr).V0).IntVal) != (0) {
							__t2 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_1.UnsafePtr).V1)
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_1.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, a_4, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_1.UnsafePtr).V1)}))}
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_3:
			return __t3
		}))
	}), gopurs_runtime.Apply(Call_Data_List_Lazy_nubByEq(eq_0), ys_2), xs_1))
}

func Call_Data_List_Lazy_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_List_Lazy_unionBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_Lazy_deleteAt(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteAt:
	for {
		if false {
			continue deleteAt
		}
		var n_0 int64 = n_0_loop
		_ = n_0
		var xs_1 gopurs_runtime.Value = xs_1_loop
		_ = xs_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_1)
		_ = __local_var_2_0
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
			_ = __local_var_4_1
			var __t3 gopurs_runtime.Value
			{
				if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil {
					var __t2 gopurs_runtime.Value
					{
						if (n_0) == (int64(0)) {
							__t2 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, Call_Data_List_Lazy_deleteAt(gopurs_runtime.IntSub(n_0, int64(1)), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)}))}
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_3:
			return __t3
		}))
	}
}

func Call_Data_List_Lazy_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope302), (TypeVar a$scope302)] Boolean)
	__local_var_1_0 := Call_Data_Eq_eq(dictEq_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_3)
		_ = __local_var_4_1
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
			_ = __local_var_6_2
			var __t4 gopurs_runtime.Value
			{
				if __local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_4
				} else {

				}
			}
			{
				if __local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil {
					var __t3 gopurs_runtime.Value
					{
						if (gopurs_runtime.Apply2(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0).IntVal) != (0) {
							__t3 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)}))}
					}
				end_branch_3:
					__t4 = __t3
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_4:
			return __t4
		}))
	})
}

func Call_Data_List_Lazy_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_Lazy_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)})
	_ = __local_var_1_0
	return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_1_0, a_3, b_2)
	}))
}

func Call_Data_List_Lazy_cycle(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var xs_0 gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var go__go_1_0_18 gopurs_runtime.Value
	_ = go__go_1_0_18
	var go__go_1_0_18_cell *gopurs_runtime.Value
	_ = go__go_1_0_18_cell
	// FALLBACK TCO: isLoop=false len=1
	go__go_1_0_18 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), xs_0, (*go__go_1_0_18_cell))
	})))
	go__go_1_0_18_cell = &go__go_1_0_18
	return go__go_1_0_18
}

func Call_Data_List_Lazy_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1, a_1, b_0)
}

func Call_Data_List_Lazy_concat(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1, v_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Data_List_Lazy_alterAt(n_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
alterAt:
	for {
		if false {
			continue alterAt
		}
		var n_0 int64 = n_0_loop
		_ = n_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var xs_2 gopurs_runtime.Value = xs_2_loop
		_ = xs_2
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
		_ = __local_var_3_0
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
			_ = __local_var_5_1
			var __t5 gopurs_runtime.Value
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil {
					__t5 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_5
				} else {

				}
			}
			{
				if __local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil {
					var __t4 gopurs_runtime.Value
					{
						if (n_0) == (int64(0)) {
							// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope318)])
							v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0))
							_ = v2_6_2
							var __t3 gopurs_runtime.Value
							{
								if v2_6_2 == nil {
									__t3 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)
									goto end_branch_3
								} else {

								}
							}
							{
								if v2_6_2 != nil {
									__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v2_6_2).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1}))}
									goto end_branch_3
								} else {

								}
							}
							{
								__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
							}
						end_branch_3:
							__t4 = __t3
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_alterAt(gopurs_runtime.IntSub(n_0, int64(1)), f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)}))}
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_5:
			return __t5
		}))
	}
}

func Call_Data_List_Lazy_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope325)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope325)]))
	__local_var_2_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), f_1)
	_ = __local_var_2_0
	return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_3)
		_ = __local_var_4_1
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
			__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
			_ = __local_var_6_2
			var __t6 gopurs_runtime.Value
			{
				if __local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil {
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
					goto end_branch_6
				} else {

				}
			}
			{
				if __local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil {
					var __t5 gopurs_runtime.Value
					{
						if (n_0) == (int64(0)) {
							// TAST (Let): v2_7_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope325)])
							v2_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0))
							_ = v2_7_3
							var __t4 gopurs_runtime.Value
							{
								if v2_7_3 == nil {
									__t4 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)
									goto end_branch_4
								} else {

								}
							}
							{
								if v2_7_3 != nil {
									__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v2_7_3).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1}))}
									goto end_branch_4
								} else {

								}
							}
							{
								__t4 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
							}
						end_branch_4:
							__t5 = __t4
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, Call_Data_List_Lazy_alterAt(gopurs_runtime.IntSub(n_0, int64(1)), __local_var_2_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)}))}
					}
				end_branch_5:
					__t6 = __t5
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_6:
			return __t6
		}))
	})
}

func Rebox_Data_List_Lazy_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_List_Lazy_2663347022_3591112874(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = (in.V2.IntVal) != (0)
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = (in.V5.IntVal) != (0)
	return out
}

func Rebox_Data_List_Lazy_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_List_Lazy_3094389156_3833657837(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]{}
	out.V0 = func() struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			head gopurs_runtime.Value
			tail gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}
		clone.head = gopurs_runtime.RecordGet(orig, "head")
		clone.tail = gopurs_runtime.RecordGet(orig, "tail")
		return clone
	}()
	return out
}

func Rebox_Data_List_Lazy_3591112874_2663347022(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Bool(in.V2)
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = gopurs_runtime.Bool(in.V5)
	return out
}

func Rebox_Data_List_Lazy_3833657837_3094389156(in *Constructor_Data_Maybe_Just[struct {
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
	}()
	return out
}

func Rebox_Data_List_Lazy_4130553207_1542299734(in *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}
