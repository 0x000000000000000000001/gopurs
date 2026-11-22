package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_ZipList_go__append gopurs_runtime.Value
var once_Data_List_ZipList_go__append sync.Once

func Get_Data_List_ZipList_go__append() gopurs_runtime.Value {
	once_Data_List_ZipList_go__append.Do(func() {
		cache_Data_List_ZipList_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_List_ZipList_go__append
}

var cache_Data_List_ZipList_append1 gopurs_runtime.Value
var once_Data_List_ZipList_append1 sync.Once

func Get_Data_List_ZipList_append1() gopurs_runtime.Value {
	once_Data_List_ZipList_append1.Do(func() {
		cache_Data_List_ZipList_append1 = gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append")
	})
	return cache_Data_List_ZipList_append1
}

var cache_Data_List_ZipList_ZipList gopurs_runtime.Value
var once_Data_List_ZipList_ZipList sync.Once

func Get_Data_List_ZipList_ZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_ZipList.Do(func() {
		cache_Data_List_ZipList_ZipList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_ZipList_ZipList(x_0_box)
		})
	})
	return cache_Data_List_ZipList_ZipList
}

var cache_Data_List_ZipList_traversableZipList gopurs_runtime.Value
var once_Data_List_ZipList_traversableZipList sync.Once

func Get_Data_List_ZipList_traversableZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_traversableZipList.Do(func() {
		cache_Data_List_ZipList_traversableZipList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_traversableList()))}
	})
	return cache_Data_List_ZipList_traversableZipList
}

var cache_Data_List_ZipList_showZipList gopurs_runtime.Value
var once_Data_List_ZipList_showZipList sync.Once

func Get_Data_List_ZipList_showZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_showZipList.Do(func() {
		cache_Data_List_ZipList_showZipList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_ZipList_showZipList(dictShow_0_box)
		})
	})
	return cache_Data_List_ZipList_showZipList
}

var cache_Data_List_ZipList_semigroupZipList gopurs_runtime.Value
var once_Data_List_ZipList_semigroupZipList sync.Once

func Get_Data_List_ZipList_semigroupZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_semigroupZipList.Do(func() {
		cache_Data_List_ZipList_semigroupZipList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_semigroupList()))}
	})
	return cache_Data_List_ZipList_semigroupZipList
}

var cache_Data_List_ZipList_ordZipList gopurs_runtime.Value
var once_Data_List_ZipList_ordZipList sync.Once

func Get_Data_List_ZipList_ordZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_ordZipList.Do(func() {
		cache_Data_List_ZipList_ordZipList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_ZipList_ordZipList(dictOrd_0_box)
		})
	})
	return cache_Data_List_ZipList_ordZipList
}

var cache_Data_List_ZipList_newtypeZipList gopurs_runtime.Value
var once_Data_List_ZipList_newtypeZipList sync.Once

func Get_Data_List_ZipList_newtypeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_newtypeZipList.Do(func() {
		cache_Data_List_ZipList_newtypeZipList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_List_ZipList_newtypeZipList
}

var cache_Data_List_ZipList_monoidZipList gopurs_runtime.Value
var once_Data_List_ZipList_monoidZipList sync.Once

func Get_Data_List_ZipList_monoidZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_monoidZipList.Do(func() {
		cache_Data_List_ZipList_monoidZipList = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_monoidList()))}
	})
	return cache_Data_List_ZipList_monoidZipList
}

var cache_Data_List_ZipList_functorZipList gopurs_runtime.Value
var once_Data_List_ZipList_functorZipList sync.Once

func Get_Data_List_ZipList_functorZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_functorZipList.Do(func() {
		cache_Data_List_ZipList_functorZipList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
	})
	return cache_Data_List_ZipList_functorZipList
}

var cache_Data_List_ZipList_foldableZipList gopurs_runtime.Value
var once_Data_List_ZipList_foldableZipList sync.Once

func Get_Data_List_ZipList_foldableZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_foldableZipList.Do(func() {
		cache_Data_List_ZipList_foldableZipList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList()))}
	})
	return cache_Data_List_ZipList_foldableZipList
}

var cache_Data_List_ZipList_eqZipList gopurs_runtime.Value
var once_Data_List_ZipList_eqZipList sync.Once

func Get_Data_List_ZipList_eqZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_eqZipList.Do(func() {
		cache_Data_List_ZipList_eqZipList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_ZipList_eqZipList(dictEq_0_box)
		})
	})
	return cache_Data_List_ZipList_eqZipList
}

var cache_Data_List_ZipList_applyZipList gopurs_runtime.Value
var once_Data_List_ZipList_applyZipList sync.Once

func Get_Data_List_ZipList_applyZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_applyZipList.Do(func() {
		cache_Data_List_ZipList_applyZipList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
			__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), v_0)
			_ = __local_var_2_0
			// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=Any
			__local_var_3_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), v1_1)
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
								__t3 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, Call_Data_Function_apply((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(Get_Data_Function_apply(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V1)}))}
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
		})}))}
	})
	return cache_Data_List_ZipList_applyZipList
}

var cache_Data_List_ZipList_Bind_dollar_Dict gopurs_runtime.Value
var once_Data_List_ZipList_Bind_dollar_Dict sync.Once

func Get_Data_List_ZipList_Bind_dollar_Dict() gopurs_runtime.Value {
	once_Data_List_ZipList_Bind_dollar_Dict.Do(func() {
		cache_Data_List_ZipList_Bind_dollar_Dict = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_ZipList_applyZipList()))}
		}), gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.StrWithCodeUnits("bind: unreachable", 17))}))}
	})
	return cache_Data_List_ZipList_Bind_dollar_Dict
}

var cache_Data_List_ZipList_zipListIsNotBind gopurs_runtime.Value
var once_Data_List_ZipList_zipListIsNotBind sync.Once

func Get_Data_List_ZipList_zipListIsNotBind() gopurs_runtime.Value {
	once_Data_List_ZipList_zipListIsNotBind.Do(func() {
		cache_Data_List_ZipList_zipListIsNotBind = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_ZipList_zipListIsNotBind(_dollar___unused_0_box)
		})
	})
	return cache_Data_List_ZipList_zipListIsNotBind
}

var cache_Data_List_ZipList_applicativeZipList gopurs_runtime.Value
var once_Data_List_ZipList_applicativeZipList sync.Once

func Get_Data_List_ZipList_applicativeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_applicativeZipList.Do(func() {
		cache_Data_List_ZipList_applicativeZipList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_ZipList_applyZipList()))}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_ZipList_ZipList(), Get_Data_List_Lazy_repeat())}))}
	})
	return cache_Data_List_ZipList_applicativeZipList
}

var cache_Data_List_ZipList_altZipList gopurs_runtime.Value
var once_Data_List_ZipList_altZipList sync.Once

func Get_Data_List_ZipList_altZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_altZipList.Do(func() {
		cache_Data_List_ZipList_altZipList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), v_0, gopurs_runtime.Apply(Call_Data_List_Lazy_drop(gopurs_runtime.Apply(Get_Data_List_Lazy_length(), v_0).IntVal), v1_1))
		})}))}
	})
	return cache_Data_List_ZipList_altZipList
}

var cache_Data_List_ZipList_plusZipList gopurs_runtime.Value
var once_Data_List_ZipList_plusZipList sync.Once

func Get_Data_List_ZipList_plusZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_plusZipList.Do(func() {
		cache_Data_List_ZipList_plusZipList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Data_List_ZipList_altZipList()))}
		}), gopurs_runtime.RecordGet(Get_Data_List_ZipList_monoidZipList(), "mempty")}))}
	})
	return cache_Data_List_ZipList_plusZipList
}

var cache_Data_List_ZipList_alternativeZipList gopurs_runtime.Value
var once_Data_List_ZipList_alternativeZipList sync.Once

func Get_Data_List_ZipList_alternativeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_alternativeZipList.Do(func() {
		cache_Data_List_ZipList_alternativeZipList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_ZipList_applicativeZipList()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_ZipList_plusZipList()))}
		})}))}
	})
	return cache_Data_List_ZipList_alternativeZipList
}

func Call_Data_List_ZipList_ZipList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_ZipList_showZipList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(ZipList ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_List_Lazy_Types_showList(dictShow_0), "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_List_ZipList_ordZipList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_ordList(dictOrd_0)))}
}

func Call_Data_List_ZipList_eqZipList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_eqList(dictEq_0)))}
}

func Call_Data_List_ZipList_zipListIsNotBind(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_ZipList_Bind_dollar_Dict()))}
}
