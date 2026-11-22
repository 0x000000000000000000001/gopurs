package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Lazy_Types_go__map gopurs_runtime.Value
var once_Data_List_Lazy_Types_go__map sync.Once

func Get_Data_List_Lazy_Types_go__map() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_go__map.Do(func() {
		cache_Data_List_Lazy_Types_go__map = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_Types_go__map
}

var cache_Data_List_Lazy_Types_unwrap gopurs_runtime.Value
var once_Data_List_Lazy_Types_unwrap sync.Once

func Get_Data_List_Lazy_Types_unwrap() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unwrap.Do(func() {
		cache_Data_List_Lazy_Types_unwrap = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_Types_unwrap
}

var cache_Data_List_Lazy_Types_compose gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose sync.Once

func Get_Data_List_Lazy_Types_compose() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose.Do(func() {
		cache_Data_List_Lazy_Types_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose
}

var cache_Data_List_Lazy_Types_map1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_map1 sync.Once

func Get_Data_List_Lazy_Types_map1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_map1.Do(func() {
		cache_Data_List_Lazy_Types_map1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_Types_map1
}

var cache_Data_List_Lazy_Types_unwrap1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unwrap1 sync.Once

func Get_Data_List_Lazy_Types_unwrap1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unwrap1.Do(func() {
		cache_Data_List_Lazy_Types_unwrap1 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_Types_unwrap1
}

var cache_Data_List_Lazy_Types_map2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_map2 sync.Once

func Get_Data_List_Lazy_Types_map2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_map2.Do(func() {
		cache_Data_List_Lazy_Types_map2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_Types_map2
}

var cache_Data_List_Lazy_Types_eq1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1 sync.Once

func Get_Data_List_Lazy_Types_eq1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1.Do(func() {
		cache_Data_List_Lazy_Types_eq1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Lazy_eq1Lazy()).V0
	})
	return cache_Data_List_Lazy_Types_eq1
}

var cache_Data_List_Lazy_Types_compare1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compare1 sync.Once

func Get_Data_List_Lazy_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compare1.Do(func() {
		cache_Data_List_Lazy_Types_compare1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_Lazy_ord1Lazy()).V1
	})
	return cache_Data_List_Lazy_Types_compare1
}

var cache_Data_List_Lazy_Types_go__append gopurs_runtime.Value
var once_Data_List_Lazy_Types_go__append sync.Once

func Get_Data_List_Lazy_Types_go__append() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_go__append.Do(func() {
		cache_Data_List_Lazy_Types_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_List_Lazy_Types_go__append
}

var cache_Data_List_Lazy_Types_add gopurs_runtime.Value
var once_Data_List_Lazy_Types_add sync.Once

func Get_Data_List_Lazy_Types_add() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_add.Do(func() {
		cache_Data_List_Lazy_Types_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_List_Lazy_Types_add
}

var cache_Data_List_Lazy_Types_sub gopurs_runtime.Value
var once_Data_List_Lazy_Types_sub sync.Once

func Get_Data_List_Lazy_Types_sub() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_sub.Do(func() {
		cache_Data_List_Lazy_Types_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_List_Lazy_Types_sub
}

var cache_Data_List_Lazy_Types_compose1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose1 sync.Once

func Get_Data_List_Lazy_Types_compose1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose1.Do(func() {
		cache_Data_List_Lazy_Types_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose1
}

var cache_Data_List_Lazy_Types_compose2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose2 sync.Once

func Get_Data_List_Lazy_Types_compose2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose2.Do(func() {
		cache_Data_List_Lazy_Types_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose2
}

var cache_Data_List_Lazy_Types_compose3 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose3 sync.Once

func Get_Data_List_Lazy_Types_compose3() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose3.Do(func() {
		cache_Data_List_Lazy_Types_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose3
}

var cache_Data_List_Lazy_Types_compose4 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose4 sync.Once

func Get_Data_List_Lazy_Types_compose4() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose4.Do(func() {
		cache_Data_List_Lazy_Types_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose4
}

var cache_Data_List_Lazy_Types_compose5 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose5 sync.Once

func Get_Data_List_Lazy_Types_compose5() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose5.Do(func() {
		cache_Data_List_Lazy_Types_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose5
}

var cache_Data_List_Lazy_Types_compose6 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose6 sync.Once

func Get_Data_List_Lazy_Types_compose6() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose6.Do(func() {
		cache_Data_List_Lazy_Types_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose6
}

var cache_Data_List_Lazy_Types_identity gopurs_runtime.Value
var once_Data_List_Lazy_Types_identity sync.Once

func Get_Data_List_Lazy_Types_identity() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_identity.Do(func() {
		cache_Data_List_Lazy_Types_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_identity(x_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_identity
}

var cache_Data_List_Lazy_Types_compose7 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose7 sync.Once

func Get_Data_List_Lazy_Types_compose7() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose7.Do(func() {
		cache_Data_List_Lazy_Types_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose7
}

var cache_Data_List_Lazy_Types_map3 gopurs_runtime.Value
var once_Data_List_Lazy_Types_map3 sync.Once

func Get_Data_List_Lazy_Types_map3() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_map3.Do(func() {
		cache_Data_List_Lazy_Types_map3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0
	})
	return cache_Data_List_Lazy_Types_map3
}

var cache_Data_List_Lazy_Types_unwrap2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unwrap2 sync.Once

func Get_Data_List_Lazy_Types_unwrap2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unwrap2.Do(func() {
		cache_Data_List_Lazy_Types_unwrap2 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_List_Lazy_Types_unwrap2
}

var cache_Data_List_Lazy_Types_unwrap3 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unwrap3 sync.Once

func Get_Data_List_Lazy_Types_unwrap3() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unwrap3.Do(func() {
		cache_Data_List_Lazy_Types_unwrap3 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_Types_unwrap3
}

var cache_Data_List_Lazy_Types_compose8 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compose8 sync.Once

func Get_Data_List_Lazy_Types_compose8() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compose8.Do(func() {
		cache_Data_List_Lazy_Types_compose8 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_Lazy_Types_compose8
}

var cache_Data_List_Lazy_Types_List gopurs_runtime.Value
var once_Data_List_Lazy_Types_List sync.Once

func Get_Data_List_Lazy_Types_List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_List.Do(func() {
		cache_Data_List_Lazy_Types_List = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_List(x_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_List
}

var cache_Data_List_Lazy_Types_Nil gopurs_runtime.Value
var once_Data_List_Lazy_Types_Nil sync.Once

func Get_Data_List_Lazy_Types_Nil() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_Nil.Do(func() {
		cache_Data_List_Lazy_Types_Nil = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Data_List_Lazy_Types_Nil
}

var cache_Data_List_Lazy_Types_Cons gopurs_runtime.Value
var once_Data_List_Lazy_Types_Cons sync.Once

func Get_Data_List_Lazy_Types_Cons() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_Cons.Do(func() {
		cache_Data_List_Lazy_Types_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_List_Lazy_Types_Cons
}

var cache_Data_List_Lazy_Types_NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_NonEmptyList sync.Once

func Get_Data_List_Lazy_Types_NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_NonEmptyList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_NonEmptyList(x_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_NonEmptyList
}

var cache_Data_List_Lazy_Types_nil gopurs_runtime.Value
var once_Data_List_Lazy_Types_nil sync.Once

func Get_Data_List_Lazy_Types_nil() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_nil.Do(func() {
		cache_Data_List_Lazy_Types_nil = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
		}))
	})
	return cache_Data_List_Lazy_Types_nil
}

var cache_Data_List_Lazy_Types_newtypeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_newtypeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_List_Lazy_Types_newtypeNonEmptyList
}

var cache_Data_List_Lazy_Types_newtypeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_newtypeList sync.Once

func Get_Data_List_Lazy_Types_newtypeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_newtypeList.Do(func() {
		cache_Data_List_Lazy_Types_newtypeList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_List_Lazy_Types_newtypeList
}

var cache_Data_List_Lazy_Types_step gopurs_runtime.Value
var once_Data_List_Lazy_Types_step sync.Once

func Get_Data_List_Lazy_Types_step() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_step.Do(func() {
		cache_Data_List_Lazy_Types_step = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Lazy_force(), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})))
	})
	return cache_Data_List_Lazy_Types_step
}

var cache_Data_List_Lazy_Types_semigroupList gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupList sync.Once

func Get_Data_List_Lazy_Types_semigroupList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupList.Do(func() {
		cache_Data_List_Lazy_Types_semigroupList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
			__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_0)
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
				__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
				_ = __local_var_4_1
				var __t2 gopurs_runtime.Value
				{
					if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil {
						__t2 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), ys_1)
						goto end_branch_2
					} else {

					}
				}
				{
					if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, ys_1)}))}
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
		})}))}
	})
	return cache_Data_List_Lazy_Types_semigroupList
}

var cache_Data_List_Lazy_Types_append1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_append1 sync.Once

func Get_Data_List_Lazy_Types_append1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_append1.Do(func() {
		cache_Data_List_Lazy_Types_append1 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_Lazy_Types_append1
}

var cache_Data_List_Lazy_Types_append2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_append2 sync.Once

func Get_Data_List_Lazy_Types_append2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_append2.Do(func() {
		cache_Data_List_Lazy_Types_append2 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_Lazy_Types_append2
}

var cache_Data_List_Lazy_Types_append3 gopurs_runtime.Value
var once_Data_List_Lazy_Types_append3 sync.Once

func Get_Data_List_Lazy_Types_append3() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_append3.Do(func() {
		cache_Data_List_Lazy_Types_append3 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_Lazy_Types_append3
}

var cache_Data_List_Lazy_Types_append4 gopurs_runtime.Value
var once_Data_List_Lazy_Types_append4 sync.Once

func Get_Data_List_Lazy_Types_append4() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_append4.Do(func() {
		cache_Data_List_Lazy_Types_append4 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_Lazy_Types_append4
}

var cache_Data_List_Lazy_Types_monoidList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monoidList sync.Once

func Get_Data_List_Lazy_Types_monoidList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monoidList.Do(func() {
		cache_Data_List_Lazy_Types_monoidList = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_List_Lazy_Types_semigroupList()
		}), Get_Data_List_Lazy_Types_nil()}))}
	})
	return cache_Data_List_Lazy_Types_monoidList
}

var cache_Data_List_Lazy_Types_lazyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_lazyList sync.Once

func Get_Data_List_Lazy_Types_lazyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_lazyList.Do(func() {
		cache_Data_List_Lazy_Types_lazyList = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), f_0))
		})}))}
	})
	return cache_Data_List_Lazy_Types_lazyList
}

var cache_Data_List_Lazy_Types_go__defer gopurs_runtime.Value
var once_Data_List_Lazy_Types_go__defer sync.Once

func Get_Data_List_Lazy_Types_go__defer() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_go__defer.Do(func() {
		cache_Data_List_Lazy_Types_go__defer = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer")
	})
	return cache_Data_List_Lazy_Types_go__defer
}

var cache_Data_List_Lazy_Types_defer1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_defer1 sync.Once

func Get_Data_List_Lazy_Types_defer1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_defer1.Do(func() {
		cache_Data_List_Lazy_Types_defer1 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer")
	})
	return cache_Data_List_Lazy_Types_defer1
}

var cache_Data_List_Lazy_Types_defer2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_defer2 sync.Once

func Get_Data_List_Lazy_Types_defer2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_defer2.Do(func() {
		cache_Data_List_Lazy_Types_defer2 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer")
	})
	return cache_Data_List_Lazy_Types_defer2
}

var cache_Data_List_Lazy_Types_functorList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorList sync.Once

func Get_Data_List_Lazy_Types_functorList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorList.Do(func() {
		cache_Data_List_Lazy_Types_functorList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
			__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_1)
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
				__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
				_ = __local_var_4_1
				var __t2 gopurs_runtime.Value
				{
					if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
						goto end_branch_2
					} else {

					}
				}
				{
					if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0, f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)}))}
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
		})}))}
	})
	return cache_Data_List_Lazy_Types_functorList
}

var cache_Data_List_Lazy_Types_map4 gopurs_runtime.Value
var once_Data_List_Lazy_Types_map4 sync.Once

func Get_Data_List_Lazy_Types_map4() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_map4.Do(func() {
		cache_Data_List_Lazy_Types_map4 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}), "map")
	})
	return cache_Data_List_Lazy_Types_map4
}

var cache_Data_List_Lazy_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
			__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}), "map"), f_0)
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1))
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_functorNonEmptyList
}

var cache_Data_List_Lazy_Types_eq1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1List sync.Once

func Get_Data_List_Lazy_Types_eq1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1List.Do(func() {
		cache_Data_List_Lazy_Types_eq1List = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var go__go_3_0_0 gopurs_runtime.Value
			_ = go__go_3_0_0
			var go__go_3_0_0_cell *gopurs_runtime.Value
			_ = go__go_3_0_0_cell
			// FALLBACK TCO: isLoop=false len=1
			go__go_3_0_0 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 bool
				{
					if v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil {
						__t1 = (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil)
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2((*go__go_3_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0))))
				}
			end_branch_1:
				return gopurs_runtime.Bool(__t1)
			})
			go__go_3_0_0_cell = &go__go_3_0_0
			return gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), ys_2)))})
		})}))}
	})
	return cache_Data_List_Lazy_Types_eq1List
}

var cache_Data_List_Lazy_Types_eqNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_eqNonEmpty sync.Once

func Get_Data_List_Lazy_Types_eqNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eqNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_eqNonEmpty = gopurs_runtime.Apply(Get_Data_NonEmpty_eqNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))})
	})
	return cache_Data_List_Lazy_Types_eqNonEmpty
}

var cache_Data_List_Lazy_Types_eq11 gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq11 sync.Once

func Get_Data_List_Lazy_Types_eq11() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq11.Do(func() {
		cache_Data_List_Lazy_Types_eq11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()).V0
	})
	return cache_Data_List_Lazy_Types_eq11
}

var cache_Data_List_Lazy_Types_eq1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eq1NonEmptyList sync.Once

func Get_Data_List_Lazy_Types_eq1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eq1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_eq1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): eq12_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope162)])]), (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope162)])])] Boolean)
			eq12_1_0 := Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Lazy_eqLazy(Call_Data_NonEmpty_eqNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))}, dictEq_0))))
			_ = eq12_1_0
			return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(eq12_1_0, v_2, v1_3)
			})
		})}))}
	})
	return cache_Data_List_Lazy_Types_eq1NonEmptyList
}

var cache_Data_List_Lazy_Types_eqList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eqList sync.Once

func Get_Data_List_Lazy_Types_eqList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eqList.Do(func() {
		cache_Data_List_Lazy_Types_eqList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_eqList(dictEq_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_eqList
}

var cache_Data_List_Lazy_Types_eqNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_eqNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_eqNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_eqNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_eqNonEmptyList(dictEq_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_eqNonEmptyList
}

var cache_Data_List_Lazy_Types_ord1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_ord1List sync.Once

func Get_Data_List_Lazy_Types_ord1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ord1List.Do(func() {
		cache_Data_List_Lazy_Types_ord1List = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))}
		}), gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Data_List_Lazy_Types_go__go_3_0_1 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32
			_ = Call_local_Data_List_Lazy_Types_go__go_3_0_1
			var go__go_3_0_1 gopurs_runtime.Value
			_ = go__go_3_0_1
			Call_local_Data_List_Lazy_Types_go__go_3_0_1 = func(v_4_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32 {
			go__go_3_0_1:
				for {
					if false {
						continue go__go_3_0_1
					}
					var v_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_4_loop
					_ = v_4
					var v1_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_5_loop
					_ = v1_5
					var __t4 uint32
					{
						if v_4 == nil {
							var __t1 uint32
							{
								if v1_5 == nil {
									__t1 = 902936544
									goto end_branch_1
								} else {

								}
							}
							{
								__t1 = 1527465420
							}
						end_branch_1:
							__t4 = __t1
							goto end_branch_4
						} else {

						}
					}
					{
						if v1_5 == nil {
							__t4 = 380165415
							goto end_branch_4
						} else {

						}
					}
					{
						if (v_4 != nil) && (v1_5 != nil) {
							// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
							v2_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v_4).V0, (v1_5).V0).IntVal)
							_ = v2_6_2
							var __t3 uint32
							{
								if v2_6_2 == 902936544 {
									v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_4).V1))
									v1_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v1_5).V1))
									continue go__go_3_0_1
									__t3 = func() uint32 { panic("unreachable") }()
									goto end_branch_3
								} else {

								}
							}
							{
								__t3 = v2_6_2
							}
						end_branch_3:
							__t4 = __t3
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = func() uint32 { panic("Failed pattern match") }()
					}
				end_branch_4:
					return __t4
				}
			}
			go__go_3_0_1 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_3_0_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val))), UnsafePtr: nil}
				})
			})
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_Types_go__go_3_0_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_1)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), ys_2)))), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_List_Lazy_Types_ord1List
}

var cache_Data_List_Lazy_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordNonEmpty sync.Once

func Get_Data_List_Lazy_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_ordNonEmpty = Call_Data_NonEmpty_ordNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()))})
	})
	return cache_Data_List_Lazy_Types_ordNonEmpty
}

var cache_Data_List_Lazy_Types_compare11 gopurs_runtime.Value
var once_Data_List_Lazy_Types_compare11 sync.Once

func Get_Data_List_Lazy_Types_compare11() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_compare11.Do(func() {
		cache_Data_List_Lazy_Types_compare11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()).V1
	})
	return cache_Data_List_Lazy_Types_compare11
}

var cache_Data_List_Lazy_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ord1NonEmptyList sync.Once

func Get_Data_List_Lazy_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_ord1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1NonEmptyList()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): compare12_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope61)])]), (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope61)])])] (ADT ["Data","Ordering","Ordering"] []))
			compare12_1_0 := Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Lazy_ordLazy(gopurs_runtime.Apply(Get_Data_List_Lazy_Types_ordNonEmpty(), dictOrd_0))))
			_ = compare12_1_0
			return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(compare12_1_0, v_2, v1_3)
			})
		})}))}
	})
	return cache_Data_List_Lazy_Types_ord1NonEmptyList
}

var cache_Data_List_Lazy_Types_ordList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordList sync.Once

func Get_Data_List_Lazy_Types_ordList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordList.Do(func() {
		cache_Data_List_Lazy_Types_ordList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_ordList(dictOrd_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_ordList
}

var cache_Data_List_Lazy_Types_ordNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_ordNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_ordNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_ordNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_ordNonEmptyList(dictOrd_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_ordNonEmptyList
}

var cache_Data_List_Lazy_Types_cons gopurs_runtime.Value
var once_Data_List_Lazy_Types_cons sync.Once

func Get_Data_List_Lazy_Types_cons() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_cons.Do(func() {
		cache_Data_List_Lazy_Types_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_cons(x_0_box, xs_1_box)
		})
	})
	return cache_Data_List_Lazy_Types_cons
}

var cache_Data_List_Lazy_Types_foldableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableList sync.Once

func Get_Data_List_Lazy_Types_foldableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableList.Do(func() {
		cache_Data_List_Lazy_Types_foldableList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope202)
			mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
			_ = mempty_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), b_3, gopurs_runtime.Apply(f_2, a_4))
				}), mempty_1_0)
			})
		}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Data_List_Lazy_Types_go__go_1_1_2 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
			_ = Call_local_Data_List_Lazy_Types_go__go_1_1_2
			var go__go_1_1_2 gopurs_runtime.Value
			_ = go__go_1_1_2
			Call_local_Data_List_Lazy_Types_go__go_1_1_2 = func(b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			go__go_1_1_2:
				for {
					if false {
						continue go__go_1_1_2
					}
					var b_2 gopurs_runtime.Value = b_2_loop
					_ = b_2
					var xs_3 gopurs_runtime.Value = xs_3_loop
					_ = xs_3
					// TAST (Let): v_4_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
					v_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
					_ = v_4_2
					var __t3 gopurs_runtime.Value
					{
						if v_4_2 == nil {
							__t3 = b_2
							goto end_branch_3
						} else {

						}
					}
					{
						if v_4_2 != nil {
							b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (v_4_2).V0)
							xs_3_loop = (v_4_2).V1
							continue go__go_1_1_2
							__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_3:
					return __t3
				}
			}
			go__go_1_1_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Data_List_Lazy_Types_go__go_1_1_2(b_2_loop_val, xs_3_loop_val)
				})
			})
			return go__go_1_1_2
		}), gopurs_runtime.Func3(func(op_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(op_0, a_4, b_3)
			}), z_1, gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Data_List_Lazy_Types_cons(a_4, b_3)
			}), Get_Data_List_Lazy_Types_nil(), xs_2))
		})}))}
	})
	return cache_Data_List_Lazy_Types_foldableList
}

var cache_Data_List_Lazy_Types_foldr gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldr sync.Once

func Get_Data_List_Lazy_Types_foldr() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldr.Do(func() {
		cache_Data_List_Lazy_Types_foldr = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2
	})
	return cache_Data_List_Lazy_Types_foldr
}

var cache_Data_List_Lazy_Types_foldr1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldr1 sync.Once

func Get_Data_List_Lazy_Types_foldr1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldr1.Do(func() {
		cache_Data_List_Lazy_Types_foldr1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2
	})
	return cache_Data_List_Lazy_Types_foldr1
}

var cache_Data_List_Lazy_Types_foldableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmpty sync.Once

func Get_Data_List_Lazy_Types_foldableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_3071895939_1680800814(Rebox_Data_List_Lazy_Types_1680800814_3071895939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))})))))}
	})
	return cache_Data_List_Lazy_Types_foldableNonEmpty
}

var cache_Data_List_Lazy_Types_foldr2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldr2 sync.Once

func Get_Data_List_Lazy_Types_foldr2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldr2.Do(func() {
		cache_Data_List_Lazy_Types_foldr2 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldr")
	})
	return cache_Data_List_Lazy_Types_foldr2
}

var cache_Data_List_Lazy_Types_foldl gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldl sync.Once

func Get_Data_List_Lazy_Types_foldl() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldl.Do(func() {
		cache_Data_List_Lazy_Types_foldl = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldl")
	})
	return cache_Data_List_Lazy_Types_foldl
}

var cache_Data_List_Lazy_Types_foldMap gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldMap sync.Once

func Get_Data_List_Lazy_Types_foldMap() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldMap.Do(func() {
		cache_Data_List_Lazy_Types_foldMap = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldMap")
	})
	return cache_Data_List_Lazy_Types_foldMap
}

var cache_Data_List_Lazy_Types_foldl1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldl1 sync.Once

func Get_Data_List_Lazy_Types_foldl1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldl1.Do(func() {
		cache_Data_List_Lazy_Types_foldl1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1
	})
	return cache_Data_List_Lazy_Types_foldl1
}

var cache_Data_List_Lazy_Types_foldl2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldl2 sync.Once

func Get_Data_List_Lazy_Types_foldl2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldl2.Do(func() {
		cache_Data_List_Lazy_Types_foldl2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1
	})
	return cache_Data_List_Lazy_Types_foldl2
}

var cache_Data_List_Lazy_Types_foldl3 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldl3 sync.Once

func Get_Data_List_Lazy_Types_foldl3() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldl3.Do(func() {
		cache_Data_List_Lazy_Types_foldl3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1
	})
	return cache_Data_List_Lazy_Types_foldl3
}

var cache_Data_List_Lazy_Types_foldl4 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldl4 sync.Once

func Get_Data_List_Lazy_Types_foldl4() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldl4.Do(func() {
		cache_Data_List_Lazy_Types_foldl4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1
	})
	return cache_Data_List_Lazy_Types_foldl4
}

var cache_Data_List_Lazy_Types_foldr3 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldr3 sync.Once

func Get_Data_List_Lazy_Types_foldr3() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldr3.Do(func() {
		cache_Data_List_Lazy_Types_foldr3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2
	})
	return cache_Data_List_Lazy_Types_foldr3
}

var cache_Data_List_Lazy_Types_extendList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendList sync.Once

func Get_Data_List_Lazy_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendList.Do(func() {
		cache_Data_List_Lazy_Types_extendList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope175)])
			v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), l_1))
			_ = v_2_0
			var __t7 gopurs_runtime.Value
			{
				if v_2_0 == nil {
					__t7 = Get_Data_List_Lazy_Types_nil()
					goto end_branch_7
				} else {

				}
			}
			{
				if v_2_0 != nil {
					// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(TypeVar b$scope174)
					__local_var_3_1 := gopurs_runtime.Apply(f_0, l_1)
					_ = __local_var_3_1
					// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
					__local_var_4_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
						__local_var_6_3 := gopurs_runtime.RecordGet(v_5, "acc")
						_ = __local_var_6_3
						// TAST (Let): __local_var_7_4 shape=Other bindingType=Any
						__local_var_7_4 := gopurs_runtime.RecordGet(v_5, "val")
						_ = __local_var_7_4
						// TAST (Let): acc_prime__8_5 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope175)])])
						acc_prime__8_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, __local_var_6_3}))}
						}))
						_ = acc_prime__8_5
						// TAST (Let): __local_var_9_6 shape=App(Other) bindingType=(TypeVar b$scope174)
						__local_var_9_6 := gopurs_runtime.Apply(f_0, acc_prime__8_5)
						_ = __local_var_9_6
						return gopurs_runtime.RecordDict2("acc", "val", acc_prime__8_5, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_9_6, __local_var_7_4}))}
						})))
					}), func() gopurs_runtime.Value {
						orig := struct {
							acc gopurs_runtime.Value
							val gopurs_runtime.Value
							/*gopurs:record-v1:2:006100630063,00760061006c*/
						}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
						_ = orig
						return gopurs_runtime.RecordDict2("acc", "val", orig.acc, orig.val)
					}(), (v_2_0).V1), "val")
					_ = __local_var_4_2
					__t7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_4_2}))}
					}))
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_7:
			return __t7
		})}))}
	})
	return cache_Data_List_Lazy_Types_extendList
}

var cache_Data_List_Lazy_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_extendNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_extendNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
			__local_var_2_0 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1).UnsafePtr).V1
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_6_1 shape=Other bindingType=Any
					__local_var_6_1 := gopurs_runtime.RecordGet(v1_5, "acc")
					_ = __local_var_6_1
					// TAST (Let): __local_var_7_2 shape=Other bindingType=Any
					__local_var_7_2 := gopurs_runtime.RecordGet(v1_5, "val")
					_ = __local_var_7_2
					// TAST (Let): __local_var_8_3 shape=App(Other) bindingType=(TypeVar b$scope179)
					__local_var_8_3 := gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, __local_var_6_1}))}
					})))
					_ = __local_var_8_3
					return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_4, __local_var_6_1}))}
					})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_8_3, __local_var_7_2}))}
					})))
				}), func() gopurs_runtime.Value {
					orig := struct {
						acc gopurs_runtime.Value
						val gopurs_runtime.Value
						/*gopurs:record-v1:2:006100630063,00760061006c*/
					}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
					_ = orig
					return gopurs_runtime.RecordDict2("acc", "val", orig.acc, orig.val)
				}(), __local_var_2_0), "val")}))}
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_extendNonEmptyList
}

var cache_Data_List_Lazy_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMap1_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope152)] (TypeVar m$scope153)), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope152)])] (TypeVar m$scope153))
			foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
			_ = foldMap1_1_0
			return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(foldMap1_1_0, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))})
			})
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldl"), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}), "foldr"), f_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
		})}))}
	})
	return cache_Data_List_Lazy_Types_foldableNonEmptyList
}

var cache_Data_List_Lazy_Types_showList gopurs_runtime.Value
var once_Data_List_Lazy_Types_showList sync.Once

func Get_Data_List_Lazy_Types_showList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_showList.Do(func() {
		cache_Data_List_Lazy_Types_showList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_showList(dictShow_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_showList
}

var cache_Data_List_Lazy_Types_showNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_showNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_showNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_showNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_showNonEmptyList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_showNonEmptyList(dictShow_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_showNonEmptyList
}

var cache_Data_List_Lazy_Types_showStep gopurs_runtime.Value
var once_Data_List_Lazy_Types_showStep sync.Once

func Get_Data_List_Lazy_Types_showStep() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_showStep.Do(func() {
		cache_Data_List_Lazy_Types_showStep = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_showStep(dictShow_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_showStep
}

var cache_Data_List_Lazy_Types_foldableWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexList sync.Once

func Get_Data_List_Lazy_Types_foldableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope229)
			mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
			_ = mempty_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())).V2, gopurs_runtime.Func2(func(i_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_4), gopurs_runtime.Apply(f_2, gopurs_runtime.Int(i_3.IntVal)))
				}), mempty_1_0)
			})
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
				__local_var_3_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
				_ = __local_var_3_1
				// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
				__local_var_4_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
				_ = __local_var_4_2
				return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Int(gopurs_runtime.IntAdd(__local_var_4_2.IntVal, int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, a_5)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				})
			}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int(int64(0)), acc_1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}))
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_3_3 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope211)])])])
			v_3_3 := Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_4 shape=Other bindingType=Any
				__local_var_4_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
				_ = __local_var_4_4
				// TAST (Let): __local_var_5_5 shape=Other bindingType=Any
				__local_var_5_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
				_ = __local_var_5_5
				return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Int(gopurs_runtime.IntAdd(__local_var_5_5.IntVal, int64(1))), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_6, __local_var_4_4}))}
						}))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				})
			}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int(int64(0)), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, xs_2)))
			_ = v_3_3
			return Call_Data_Tuple_snd(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_5_6 shape=Other bindingType=Any
				__local_var_5_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
				_ = __local_var_5_6
				// TAST (Let): __local_var_6_7 shape=Other bindingType=Any
				__local_var_6_7 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
				_ = __local_var_6_7
				return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Int(gopurs_runtime.IntSub(__local_var_6_7.IntVal, int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(gopurs_runtime.IntSub(__local_var_6_7.IntVal, int64(1))), a_7, __local_var_5_6)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				})
			}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_1728839155_138441832(Rebox_Data_List_Lazy_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int((v_3_3).V0), b_1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, (v_3_3).V1)))))
		})})))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexList
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmpty sync.Once

func Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2048287455_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2048287455(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))})))))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Lazy_Types_foldMapWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldMapWithIndex sync.Once

func Get_Data_List_Lazy_Types_foldMapWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldMapWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_foldMapWithIndex = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldMapWithIndex")
	})
	return cache_Data_List_Lazy_Types_foldMapWithIndex
}

var cache_Data_List_Lazy_Types_foldlWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldlWithIndex sync.Once

func Get_Data_List_Lazy_Types_foldlWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldlWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_foldlWithIndex = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldlWithIndex")
	})
	return cache_Data_List_Lazy_Types_foldlWithIndex
}

var cache_Data_List_Lazy_Types_foldrWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldrWithIndex sync.Once

func Get_Data_List_Lazy_Types_foldrWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldrWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_foldrWithIndex = gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldrWithIndex")
	})
	return cache_Data_List_Lazy_Types_foldrWithIndex
}

var cache_Data_List_Lazy_Types_foldrWithIndex1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldrWithIndex1 sync.Once

func Get_Data_List_Lazy_Types_foldrWithIndex1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldrWithIndex1.Do(func() {
		cache_Data_List_Lazy_Types_foldrWithIndex1 = Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())).V3
	})
	return cache_Data_List_Lazy_Types_foldrWithIndex1
}

var cache_Data_List_Lazy_Types_foldrWithIndex2 gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldrWithIndex2 sync.Once

func Get_Data_List_Lazy_Types_foldrWithIndex2() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldrWithIndex2.Do(func() {
		cache_Data_List_Lazy_Types_foldrWithIndex2 = Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())).V3
	})
	return cache_Data_List_Lazy_Types_foldrWithIndex2
}

var cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMapWithIndex1_1_0 shape=App(Other) bindingType=(Func [(Func [(ADT ["Data","Maybe","Maybe"] [Int]), (TypeVar a$scope127)] (TypeVar m$scope128)), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope127)])] (TypeVar m$scope128))
			foldMapWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
			_ = foldMapWithIndex1_1_0
			return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(foldMapWithIndex1_1_0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t1 gopurs_runtime.Value
					{
						if v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr == nil {
							__t1 = gopurs_runtime.Int(int64(0))
							goto end_branch_1
						} else {

						}
					}
					{
						if v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr != nil {
							__t1 = gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_4.UnsafePtr).V0.IntVal))
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
					}
				end_branch_1:
					return __t1
				})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))})
			})
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil {
						__t2 = gopurs_runtime.Int(int64(0))
						goto end_branch_2
					} else {

					}
				}
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil {
						__t2 = gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
				}
			end_branch_2:
				return __t2
			})), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}), "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil {
						__t3 = gopurs_runtime.Int(int64(0))
						goto end_branch_3
					} else {

					}
				}
				{
					if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil {
						__t3 = gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
				}
			end_branch_3:
				return __t3
			})), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))})
		})})))}
	})
	return cache_Data_List_Lazy_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexList sync.Once

func Get_Data_List_Lazy_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())).V3, gopurs_runtime.Func3(func(i_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(TypeVar b$scope236)
				__local_var_4_0 := gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2)
				_ = __local_var_4_0
				return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_0, acc_3}))}
				}))
			}), Get_Data_List_Lazy_Types_nil())
		})})))}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexList
}

var cache_Data_List_Lazy_Types_mapWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_mapWithIndex sync.Once

func Get_Data_List_Lazy_Types_mapWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_mapWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_mapWithIndex = gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorWithIndex(gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexList()))))}), "mapWithIndex")
	})
	return cache_Data_List_Lazy_Types_mapWithIndex
}

var cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_functorWithIndexNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorWithIndex(gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexList()))))}), "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t0 gopurs_runtime.Value
					{
						if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil {
							__t0 = gopurs_runtime.Int(int64(0))
							goto end_branch_0
						} else {

						}
					}
					{
						if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil {
							__t0 = gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
							goto end_branch_0
						} else {

						}
					}
					{
						__t0 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
					}
				end_branch_0:
					return __t0
				})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_1)))})
			}))
		})})))}
	})
	return cache_Data_List_Lazy_Types_functorWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_toList gopurs_runtime.Value
var once_Data_List_Lazy_Types_toList sync.Once

func Get_Data_List_Lazy_Types_toList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_toList.Do(func() {
		cache_Data_List_Lazy_Types_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_Lazy_Types_toList(v_0_box)
		})
	})
	return cache_Data_List_Lazy_Types_toList
}

var cache_Data_List_Lazy_Types_semigroupNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_semigroupNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_semigroupNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_semigroupNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_semigroupNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime__1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope241)])
			v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
			_ = v1_2_0
			// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
			__local_var_3_1 := (v1_2_0).V0
			_ = __local_var_3_1
			// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
			__local_var_4_2 := (v1_2_0).V1
			_ = __local_var_4_2
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), __local_var_4_2, Call_Data_List_Lazy_Types_toList(as_prime__1))}))}
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_semigroupNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableList sync.Once

func Get_Data_List_Lazy_Types_traversableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableList.Do(func() {
		cache_Data_List_Lazy_Types_traversableList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}))
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V2, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_1, a_2)), b_3)
			}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
		})}))}
	})
	return cache_Data_List_Lazy_Types_traversableList
}

var cache_Data_List_Lazy_Types_traversableNonEmpty gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmpty sync.Once

func Get_Data_List_Lazy_Types_traversableNonEmpty() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmpty.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_4249989635_3043886126(Rebox_Data_List_Lazy_Types_3043886126_4249989635(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))})))))}
	})
	return cache_Data_List_Lazy_Types_traversableNonEmpty
}

var cache_Data_List_Lazy_Types_traverse gopurs_runtime.Value
var once_Data_List_Lazy_Types_traverse sync.Once

func Get_Data_List_Lazy_Types_traverse() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traverse.Do(func() {
		cache_Data_List_Lazy_Types_traverse = gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "traverse")
	})
	return cache_Data_List_Lazy_Types_traverse
}

var cache_Data_List_Lazy_Types_sequence gopurs_runtime.Value
var once_Data_List_Lazy_Types_sequence sync.Once

func Get_Data_List_Lazy_Types_sequence() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_sequence.Do(func() {
		cache_Data_List_Lazy_Types_sequence = gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "sequence")
	})
	return cache_Data_List_Lazy_Types_sequence
}

var cache_Data_List_Lazy_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableNonEmptyList()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): sequence1_1_0 shape=App(Other) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeApp (TypeVar m$scope32) [(TypeVar a$scope31)])])] (TypeApp (TypeVar m$scope32) [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope31)])]))
			sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))})
			_ = sequence1_1_0
			return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(xxs_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return xxs_3
					}))
				}), gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_2)))}))
			})
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): traverse1_1_1 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope25)] (TypeApp (TypeVar m$scope27) [(TypeVar b$scope26)])), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope25)])] (TypeApp (TypeVar m$scope27) [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar b$scope26)])]))
			traverse1_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))})
			_ = traverse1_1_1
			return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return xxs_4
					}))
				}), gopurs_runtime.Apply2(traverse1_1_1, f_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
			})
		})}))}
	})
	return cache_Data_List_Lazy_Types_traversableNonEmptyList
}

var cache_Data_List_Lazy_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexList sync.Once

func Get_Data_List_Lazy_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexList()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexList())).V3, gopurs_runtime.Func3(func(i_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply2(f_1, gopurs_runtime.Int(i_2.IntVal), a_3)), b_4)
			}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
		})})))}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexList
}

var cache_Data_List_Lazy_Types_traverseWithIndex gopurs_runtime.Value
var once_Data_List_Lazy_Types_traverseWithIndex sync.Once

func Get_Data_List_Lazy_Types_traverseWithIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traverseWithIndex.Do(func() {
		cache_Data_List_Lazy_Types_traverseWithIndex = gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904(Rebox_Data_List_Lazy_Types_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableWithIndexList()))))}), "traverseWithIndex")
	})
	return cache_Data_List_Lazy_Types_traverseWithIndex
}

var cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2491554675_3725484264(Rebox_Data_List_Lazy_Types_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmptyList()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2773701683_2412140840(Rebox_Data_List_Lazy_Types_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorWithIndexNonEmptyList()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableNonEmptyList()))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): traverseWithIndex1_1_0 shape=App(Other) bindingType=(Func [(Func [(ADT ["Data","Maybe","Maybe"] [Int]), (TypeVar a$scope12)] (TypeApp (TypeVar m$scope14) [(TypeVar b$scope13)])), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope12)])] (TypeApp (TypeVar m$scope14) [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar b$scope13)])]))
			traverseWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_2955889203_1812164904(Rebox_Data_List_Lazy_Types_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableWithIndexList()))))}), "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))})
			_ = traverseWithIndex1_1_0
			return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(xxs_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return xxs_4
					}))
				}), gopurs_runtime.Apply2(traverseWithIndex1_1_0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t1 gopurs_runtime.Value
					{
						if v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr == nil {
							__t1 = gopurs_runtime.Int(int64(0))
							goto end_branch_1
						} else {

						}
					}
					{
						if v2_4.Type == 9 && v2_4.IntVal == 930809136 && v2_4.UnsafePtr != nil {
							__t1 = gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_4.UnsafePtr).V0.IntVal))
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
					}
				end_branch_1:
					return __t1
				})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3)))}))
			})
		})})))}
	})
	return cache_Data_List_Lazy_Types_traversableWithIndexNonEmptyList
}

var cache_Data_List_Lazy_Types_unfoldable1List gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1List sync.Once

func Get_Data_List_Lazy_Types_unfoldable1List() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1List.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1List = func() gopurs_runtime.Value {
			var go__go_0_0_3 gopurs_runtime.Value
			_ = go__go_0_0_3
			var go__go_0_0_3_cell *gopurs_runtime.Value
			_ = go__go_0_0_3_cell
			// FALLBACK TCO: isLoop=false len=1
			go__go_0_0_3 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope271), (ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope272)])])
					v1_4_1 := Rebox_Data_List_Lazy_Types_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, b_2)))
					_ = v1_4_1
					var __t7 gopurs_runtime.Value
					{
						var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v1_4_1).V1
						_ = __t_tag_2
						if __t_tag_2 != nil {
							// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
							__local_var_5_3 := (v1_4_1).V0
							_ = __local_var_5_3
							// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope271)])])
							__local_var_6_4 := gopurs_runtime.Apply2((*go__go_0_0_3_cell), f_1, ((v1_4_1).V1).V0)
							_ = __local_var_6_4
							__t7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_3, __local_var_6_4}))}
							}))
							goto end_branch_7
						} else {

						}
					}
					{
						var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v1_4_1).V1
						_ = __t_tag_5
						if __t_tag_5 == nil {
							// TAST (Let): __local_var_5_6 shape=Other bindingType=Any
							__local_var_5_6 := (v1_4_1).V0
							_ = __local_var_5_6
							__t7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_6, Get_Data_List_Lazy_Types_nil()}))}
							}))
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_7:
					return __t7
				})))
			})
			go__go_0_0_3_cell = &go__go_0_0_3
			return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, go__go_0_0_3}))}
		}()
	})
	return cache_Data_List_Lazy_Types_unfoldable1List
}

var cache_Data_List_Lazy_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldableList sync.Once

func Get_Data_List_Lazy_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldableList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldableList = func() gopurs_runtime.Value {
			var go__go_0_0_4 gopurs_runtime.Value
			_ = go__go_0_0_4
			var go__go_0_0_4_cell *gopurs_runtime.Value
			_ = go__go_0_0_4_cell
			// FALLBACK TCO: isLoop=false len=1
			go__go_0_0_4 = gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope276), (TypeVar b$scope277)])])
					v1_4_1 := Rebox_Data_List_Lazy_Types_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, b_2)))
					_ = v1_4_1
					var __t4 gopurs_runtime.Value
					{
						if v1_4_1 == nil {
							__t4 = Get_Data_List_Lazy_Types_nil()
							goto end_branch_4
						} else {

						}
					}
					{
						if v1_4_1 != nil {
							// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
							__local_var_5_2 := ((v1_4_1).V0).V0
							_ = __local_var_5_2
							// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope276)])])
							__local_var_6_3 := gopurs_runtime.Apply2((*go__go_0_0_4_cell), f_1, ((v1_4_1).V0).V1)
							_ = __local_var_6_3
							__t4 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_2, __local_var_6_3}))}
							}))
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_4:
					return __t4
				})))
			})
			go__go_0_0_4_cell = &go__go_0_0_4
			return gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldable1List()))}
			}), go__go_0_0_4}))}
		}()
	})
	return cache_Data_List_Lazy_Types_unfoldableList
}

var cache_Data_List_Lazy_Types_unfoldr1 gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldr1 sync.Once

func Get_Data_List_Lazy_Types_unfoldr1() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldr1.Do(func() {
		cache_Data_List_Lazy_Types_unfoldr1 = gopurs_runtime.RecordGet(Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldableList()))}), "unfoldr1")
	})
	return cache_Data_List_Lazy_Types_unfoldr1
}

var cache_Data_List_Lazy_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_unfoldable1NonEmptyList sync.Once

func Get_Data_List_Lazy_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_unfoldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer((&Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_unfoldableList()))}), "unfoldr1"), f_0, b_1)
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Lazy_Types_comonadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_comonadNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_comonadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_comonadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_comonadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_extendNonEmptyList()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_NonEmpty_head(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)))
		})}))}
	})
	return cache_Data_List_Lazy_Types_comonadNonEmptyList
}

var cache_Data_List_Lazy_Types_monadList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadList sync.Once

func Get_Data_List_Lazy_Types_monadList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadList.Do(func() {
		cache_Data_List_Lazy_Types_monadList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applicativeList()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()))}
		})}))}
	})
	return cache_Data_List_Lazy_Types_monadList
}

var cache_Data_List_Lazy_Types_bindList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindList sync.Once

func Get_Data_List_Lazy_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindList.Do(func() {
		cache_Data_List_Lazy_Types_bindList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyList()))}
		}), gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
			__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_0)
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
				__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
				_ = __local_var_4_1
				var __t2 gopurs_runtime.Value
				{
					if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))}
						goto end_branch_2
					} else {

					}
				}
				{
					if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil {
						__t2 = gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, f_1)))
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
		})}))}
	})
	return cache_Data_List_Lazy_Types_bindList
}

var cache_Data_List_Lazy_Types_applyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyList sync.Once

func Get_Data_List_Lazy_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyList.Do(func() {
		cache_Data_List_Lazy_Types_applyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
		}), Call_Control_Monad_ap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList()))}))}
	})
	return cache_Data_List_Lazy_Types_applyList
}

var cache_Data_List_Lazy_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeList sync.Once

func Get_Data_List_Lazy_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyList()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()}))}
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_applicativeList
}

var cache_Data_List_Lazy_Types_apply gopurs_runtime.Value
var once_Data_List_Lazy_Types_apply sync.Once

func Get_Data_List_Lazy_Types_apply() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_apply.Do(func() {
		cache_Data_List_Lazy_Types_apply = gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyList()).V1
	})
	return cache_Data_List_Lazy_Types_apply
}

var cache_Data_List_Lazy_Types_bind gopurs_runtime.Value
var once_Data_List_Lazy_Types_bind sync.Once

func Get_Data_List_Lazy_Types_bind() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bind.Do(func() {
		cache_Data_List_Lazy_Types_bind = gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1
	})
	return cache_Data_List_Lazy_Types_bind
}

var cache_Data_List_Lazy_Types_applyNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applyNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_applyNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applyNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applyNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope304)])
			v2_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v1_1))
			_ = v2_2_0
			// TAST (Let): v3_3_1 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (Func [(TypeVar a$scope304)] (TypeVar b$scope305))])
			v3_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
			_ = v3_3_1
			// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
			__local_var_4_2 := (v2_2_0).V0
			_ = __local_var_4_2
			// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
			__local_var_5_3 := (v2_2_0).V1
			_ = __local_var_5_3
			// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
			__local_var_6_4 := (v3_3_1).V0
			_ = __local_var_6_4
			// TAST (Let): __local_var_7_5 shape=Other bindingType=Any
			__local_var_7_5 := (v3_3_1).V1
			_ = __local_var_7_5
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_4, __local_var_4_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyList()).V1, __local_var_7_5, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_2, Get_Data_List_Lazy_Types_nil()}))}
				}))), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyList()).V1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_4, __local_var_7_5}))}
				})), __local_var_5_3))}))}
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_applyNonEmptyList
}

var cache_Data_List_Lazy_Types_bindNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_bindNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_bindNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_bindNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_bindNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope285)])
			v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
			_ = v1_2_0
			// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
			__local_var_3_1 := (v1_2_0).V1
			_ = __local_var_3_1
			// TAST (Let): v2_4_2 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar b$scope286)])
			v2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply(f_1, (v1_2_0).V0))))
			_ = v2_4_2
			// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
			__local_var_5_3 := (v2_4_2).V0
			_ = __local_var_5_3
			// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
			__local_var_6_4 := (v2_4_2).V1
			_ = __local_var_6_4
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), __local_var_6_4, gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1, __local_var_3_1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_toList(), f_1)))}))}
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_bindNonEmptyList
}

var cache_Data_List_Lazy_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_altNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorNonEmptyList()))}
		}), gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupNonEmptyList(), "append")}))}
	})
	return cache_Data_List_Lazy_Types_altNonEmptyList
}

var cache_Data_List_Lazy_Types_altList gopurs_runtime.Value
var once_Data_List_Lazy_Types_altList sync.Once

func Get_Data_List_Lazy_Types_altList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_altList.Do(func() {
		cache_Data_List_Lazy_Types_altList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
		}), gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")}))}
	})
	return cache_Data_List_Lazy_Types_altList
}

var cache_Data_List_Lazy_Types_plusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_plusList sync.Once

func Get_Data_List_Lazy_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_plusList.Do(func() {
		cache_Data_List_Lazy_Types_plusList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_altList()))}
		}), Get_Data_List_Lazy_Types_nil()}))}
	})
	return cache_Data_List_Lazy_Types_plusList
}

var cache_Data_List_Lazy_Types_singleton gopurs_runtime.Value
var once_Data_List_Lazy_Types_singleton sync.Once

func Get_Data_List_Lazy_Types_singleton() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_singleton.Do(func() {
		cache_Data_List_Lazy_Types_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_Types_singleton(a_0_box))}
		})
	})
	return cache_Data_List_Lazy_Types_singleton
}

var cache_Data_List_Lazy_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Lazy_Types_alternativeList sync.Once

func Get_Data_List_Lazy_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_alternativeList.Do(func() {
		cache_Data_List_Lazy_Types_alternativeList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applicativeList()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_plusList()))}
		})}))}
	})
	return cache_Data_List_Lazy_Types_alternativeList
}

var cache_Data_List_Lazy_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadPlusList sync.Once

func Get_Data_List_Lazy_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadPlusList.Do(func() {
		cache_Data_List_Lazy_Types_monadPlusList = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_alternativeList()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monadList()))}
		})}))}
	})
	return cache_Data_List_Lazy_Types_monadPlusList
}

var cache_Data_List_Lazy_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_applicativeNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_applicativeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applyNonEmptyList()))}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()}))}
			}))
		})}))}
	})
	return cache_Data_List_Lazy_Types_applicativeNonEmptyList
}

var cache_Data_List_Lazy_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Lazy_Types_monadNonEmptyList sync.Once

func Get_Data_List_Lazy_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Lazy_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Lazy_Types_monadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_applicativeNonEmptyList()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindNonEmptyList()))}
		})}))}
	})
	return cache_Data_List_Lazy_Types_monadNonEmptyList
}

type Constructor_Data_List_Lazy_Types_Nil[T_a any] struct {
	Rc uint32
}

type Constructor_Data_List_Lazy_Types_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}

func Call_Data_List_Lazy_Types_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_Lazy_Types_List(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_Lazy_Types_NonEmptyList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_Lazy_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()).V0, dictEq_0)}))}
}

func Call_Data_List_Lazy_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Lazy_eqLazy(Call_Data_NonEmpty_eqNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_eq1List()))}, dictEq_0))))}
}

func Call_Data_List_Lazy_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqList1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope57)])])])
	eqList1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_eqList(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
	_ = eqList1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_0)}
	}), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_ord1List()).V1, dictOrd_0)}))}
}

func Call_Data_List_Lazy_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Lazy_ordLazy(gopurs_runtime.Apply(Get_Data_List_Lazy_Types_ordNonEmpty(), dictOrd_0))))}
}

func Call_Data_List_Lazy_Types_cons(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var xs_1 gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, xs_1}))}
	}))
}

func Call_Data_List_Lazy_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope47)])
		v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_1))
		_ = v_2_0
		var __t1 string
		{
			if v_2_0 == nil {
				__t1 = "(fromFoldable [])"
				goto end_branch_1
			} else {

			}
		}
		{
			if v_2_0 != nil {
				__t1 = gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(fromFoldable [", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_0).V0).StrVal()), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()).V1, gopurs_runtime.Func2(func(shown_3 gopurs_runtime.Value, x_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(shown_3.StrVal(), ","), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), x_prime__4).StrVal()))
				}), gopurs_runtime.StrWithCodeUnits("", 0), (v_2_0).V1).StrVal()), "])")
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() string { panic("Failed pattern match") }()
		}
	end_branch_1:
		return gopurs_runtime.Str(__t1)
	})}))}
}

func Call_Data_List_Lazy_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(NonEmptyList ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Lazy_showLazy(Call_Data_NonEmpty_showNonEmpty(dictShow_0, Call_Data_List_Lazy_Types_showList(dictShow_0))), "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_List_Lazy_Types_showStep(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_Types_928333203_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr == nil {
				__t0 = gopurs_runtime.StrWithCodeUnits("Nil", 3)
				goto end_branch_0
			} else {

			}
		}
		{
			if v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil {
				__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).StrVal()), " : "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_List_Lazy_Types_showList(dictShow_0), "show"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1).StrVal()), ")"))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
		}
	end_branch_0:
		return __t0
	})})))}
}

func Call_Data_List_Lazy_Types_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v2_2_0 shape=App(Var) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a$scope238)])
		v2_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0))
		_ = v2_2_0
		// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
		__local_var_3_1 := (v2_2_0).V0
		_ = __local_var_3_1
		// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
		__local_var_4_2 := (v2_2_0).V1
		_ = __local_var_4_2
		return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_4_2}))}
		}))
	})))
}

func Call_Data_List_Lazy_Types_singleton(a_0_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return (&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()})
}

func Rebox_Data_List_Lazy_Types_138441832_1728839155(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_138441832_3804580809(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_List_Lazy_Types_1680800814_3071895939(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_1728839155_138441832(in *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Data_List_Lazy_Types_1812164904_2955889203(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_2048287455_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_2412140840_2773701683(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_2491554675_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_2773701683_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_2955889203_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_3043886126_4249989635(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_3071895939_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_Lazy_Types_3725484264_2048287455(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_3725484264_2491554675(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_4249989635_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_Lazy_Types_928333203_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
