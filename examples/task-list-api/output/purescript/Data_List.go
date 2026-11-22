package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_go__map gopurs_runtime.Value
var once_Data_List_go__map sync.Once

func Get_Data_List_go__map() gopurs_runtime.Value {
	once_Data_List_go__map.Do(func() {
		cache_Data_List_go__map = Rebox_Data_List_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_go__map
}

var cache_Data_List_sub gopurs_runtime.Value
var once_Data_List_sub sync.Once

func Get_Data_List_sub() gopurs_runtime.Value {
	once_Data_List_sub.Do(func() {
		cache_Data_List_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_List_sub
}

var cache_Data_List_map1 gopurs_runtime.Value
var once_Data_List_map1 sync.Once

func Get_Data_List_map1() gopurs_runtime.Value {
	once_Data_List_map1.Do(func() {
		cache_Data_List_map1 = Rebox_Data_List_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_map1
}

var cache_Data_List_tailRecM2 gopurs_runtime.Value
var once_Data_List_tailRecM2 sync.Once

func Get_Data_List_tailRecM2() gopurs_runtime.Value {
	once_Data_List_tailRecM2.Do(func() {
		cache_Data_List_tailRecM2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_tailRecM2(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](a_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_2_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_tailRecM2
}

var cache_Data_List_foldr gopurs_runtime.Value
var once_Data_List_foldr sync.Once

func Get_Data_List_foldr() gopurs_runtime.Value {
	once_Data_List_foldr.Do(func() {
		cache_Data_List_foldr = Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2
	})
	return cache_Data_List_foldr
}

var cache_Data_List_eq gopurs_runtime.Value
var once_Data_List_eq sync.Once

func Get_Data_List_eq() gopurs_runtime.Value {
	once_Data_List_eq.Do(func() {
		cache_Data_List_eq = Get_Data_Eq_eq__2123013075()
	})
	return cache_Data_List_eq
}

var cache_Data_List_notEq gopurs_runtime.Value
var once_Data_List_notEq sync.Once

func Get_Data_List_notEq() gopurs_runtime.Value {
	once_Data_List_notEq.Do(func() {
		cache_Data_List_notEq = Get_Data_Eq_notEq__2123013075()
	})
	return cache_Data_List_notEq
}

var cache_Data_List_compose gopurs_runtime.Value
var once_Data_List_compose sync.Once

func Get_Data_List_compose() gopurs_runtime.Value {
	once_Data_List_compose.Do(func() {
		cache_Data_List_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_compose
}

var cache_Data_List_go__append gopurs_runtime.Value
var once_Data_List_go__append sync.Once

func Get_Data_List_go__append() gopurs_runtime.Value {
	once_Data_List_go__append.Do(func() {
		cache_Data_List_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_List_go__append
}

var cache_Data_List_lessThan gopurs_runtime.Value
var once_Data_List_lessThan sync.Once

func Get_Data_List_lessThan() gopurs_runtime.Value {
	once_Data_List_lessThan.Do(func() {
		cache_Data_List_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_List_lessThan
}

var cache_Data_List_map2 gopurs_runtime.Value
var once_Data_List_map2 sync.Once

func Get_Data_List_map2() gopurs_runtime.Value {
	once_Data_List_map2.Do(func() {
		cache_Data_List_map2 = Rebox_Data_List_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_map2
}

var cache_Data_List_sequence gopurs_runtime.Value
var once_Data_List_sequence sync.Once

func Get_Data_List_sequence() gopurs_runtime.Value {
	once_Data_List_sequence.Do(func() {
		cache_Data_List_sequence = Rebox_Data_List_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList())).V2
	})
	return cache_Data_List_sequence
}

var cache_Data_List_eq1 gopurs_runtime.Value
var once_Data_List_eq1 sync.Once

func Get_Data_List_eq1() gopurs_runtime.Value {
	once_Data_List_eq1.Do(func() {
		cache_Data_List_eq1 = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_List_eq1
}

var cache_Data_List_add gopurs_runtime.Value
var once_Data_List_add sync.Once

func Get_Data_List_add() gopurs_runtime.Value {
	once_Data_List_add.Do(func() {
		cache_Data_List_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_List_add
}

var cache_Data_List_greaterThan gopurs_runtime.Value
var once_Data_List_greaterThan sync.Once

func Get_Data_List_greaterThan() gopurs_runtime.Value {
	once_Data_List_greaterThan.Do(func() {
		cache_Data_List_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_List_greaterThan
}

var cache_Data_List_negate gopurs_runtime.Value
var once_Data_List_negate sync.Once

func Get_Data_List_negate() gopurs_runtime.Value {
	once_Data_List_negate.Do(func() {
		cache_Data_List_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_List_negate(a_0_box.IntVal))
		})
	})
	return cache_Data_List_negate
}

var cache_Data_List_foldr1 gopurs_runtime.Value
var once_Data_List_foldr1 sync.Once

func Get_Data_List_foldr1() gopurs_runtime.Value {
	once_Data_List_foldr1.Do(func() {
		cache_Data_List_foldr1 = Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2
	})
	return cache_Data_List_foldr1
}

var cache_Data_List_compose1 gopurs_runtime.Value
var once_Data_List_compose1 sync.Once

func Get_Data_List_compose1() gopurs_runtime.Value {
	once_Data_List_compose1.Do(func() {
		cache_Data_List_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_compose1
}

var cache_Data_List_bimap gopurs_runtime.Value
var once_Data_List_bimap sync.Once

func Get_Data_List_bimap() gopurs_runtime.Value {
	once_Data_List_bimap.Do(func() {
		cache_Data_List_bimap = gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_bifunctorStep()).V0
	})
	return cache_Data_List_bimap
}

var cache_Data_List_map3 gopurs_runtime.Value
var once_Data_List_map3 sync.Once

func Get_Data_List_map3() gopurs_runtime.Value {
	once_Data_List_map3.Do(func() {
		cache_Data_List_map3 = Rebox_Data_List_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_map3
}

var cache_Data_List_map4 gopurs_runtime.Value
var once_Data_List_map4 sync.Once

func Get_Data_List_map4() gopurs_runtime.Value {
	once_Data_List_map4.Do(func() {
		cache_Data_List_map4 = Rebox_Data_List_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_map4
}

var cache_Data_List_compose2 gopurs_runtime.Value
var once_Data_List_compose2 sync.Once

func Get_Data_List_compose2() gopurs_runtime.Value {
	once_Data_List_compose2.Do(func() {
		cache_Data_List_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_compose2
}

var cache_Data_List_compose3 gopurs_runtime.Value
var once_Data_List_compose3 sync.Once

func Get_Data_List_compose3() gopurs_runtime.Value {
	once_Data_List_compose3.Do(func() {
		cache_Data_List_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_compose3
}

var cache_Data_List_map5 gopurs_runtime.Value
var once_Data_List_map5 sync.Once

func Get_Data_List_map5() gopurs_runtime.Value {
	once_Data_List_map5.Do(func() {
		cache_Data_List_map5 = Get_Data_Functor_map__3627393546()
	})
	return cache_Data_List_map5
}

var cache_Data_List_any gopurs_runtime.Value
var once_Data_List_any sync.Once

func Get_Data_List_any() gopurs_runtime.Value {
	once_Data_List_any.Do(func() {
		cache_Data_List_any = gopurs_runtime.Apply(Call_Data_Foldable_any(Rebox_Data_List_1022383170_1680800814(Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())))), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3591112874_2663347022(Rebox_Data_List_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))})
	})
	return cache_Data_List_any
}

var cache_Data_List_not gopurs_runtime.Value
var once_Data_List_not sync.Once

func Get_Data_List_not() gopurs_runtime.Value {
	once_Data_List_not.Do(func() {
		cache_Data_List_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Data_List_not
}

var cache_Data_List_append1 gopurs_runtime.Value
var once_Data_List_append1 sync.Once

func Get_Data_List_append1() gopurs_runtime.Value {
	once_Data_List_append1.Do(func() {
		cache_Data_List_append1 = gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append")
	})
	return cache_Data_List_append1
}

var cache_Data_List_foldl gopurs_runtime.Value
var once_Data_List_foldl sync.Once

func Get_Data_List_foldl() gopurs_runtime.Value {
	once_Data_List_foldl.Do(func() {
		cache_Data_List_foldl = Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V1
	})
	return cache_Data_List_foldl
}

var cache_Data_List_map6 gopurs_runtime.Value
var once_Data_List_map6 sync.Once

func Get_Data_List_map6() gopurs_runtime.Value {
	once_Data_List_map6.Do(func() {
		cache_Data_List_map6 = Rebox_Data_List_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_map6
}

var cache_Data_List_foldl1 gopurs_runtime.Value
var once_Data_List_foldl1 sync.Once

func Get_Data_List_foldl1() gopurs_runtime.Value {
	once_Data_List_foldl1.Do(func() {
		cache_Data_List_foldl1 = Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V1
	})
	return cache_Data_List_foldl1
}

var cache_Data_List_bind gopurs_runtime.Value
var once_Data_List_bind sync.Once

func Get_Data_List_bind() gopurs_runtime.Value {
	once_Data_List_bind.Do(func() {
		cache_Data_List_bind = Rebox_Data_List_2748095225_2183599445(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindList())).V1
	})
	return cache_Data_List_bind
}

var cache_Data_List_identity gopurs_runtime.Value
var once_Data_List_identity sync.Once

func Get_Data_List_identity() gopurs_runtime.Value {
	once_Data_List_identity.Do(func() {
		cache_Data_List_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_identity(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_0_box)))}
		})
	})
	return cache_Data_List_identity
}

var cache_Data_List_map7 gopurs_runtime.Value
var once_Data_List_map7 sync.Once

func Get_Data_List_map7() gopurs_runtime.Value {
	once_Data_List_map7.Do(func() {
		cache_Data_List_map7 = Rebox_Data_List_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_List_map7
}

var cache_Data_List_compose4 gopurs_runtime.Value
var once_Data_List_compose4 sync.Once

func Get_Data_List_compose4() gopurs_runtime.Value {
	once_Data_List_compose4.Do(func() {
		cache_Data_List_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_List_compose4
}

var cache_Data_List_Pattern gopurs_runtime.Value
var once_Data_List_Pattern sync.Once

func Get_Data_List_Pattern() gopurs_runtime.Value {
	once_Data_List_Pattern.Do(func() {
		cache_Data_List_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Pattern(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_0_box)))}
		})
	})
	return cache_Data_List_Pattern
}

var cache_Data_List_updateAt gopurs_runtime.Value
var once_Data_List_updateAt sync.Once

func Get_Data_List_updateAt() gopurs_runtime.Value {
	once_Data_List_updateAt.Do(func() {
		cache_Data_List_updateAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_updateAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_updateAt
}

var cache_Data_List_unzip gopurs_runtime.Value
var once_Data_List_unzip sync.Once

func Get_Data_List_unzip() gopurs_runtime.Value {
	once_Data_List_unzip.Do(func() {
		cache_Data_List_unzip = gopurs_runtime.Apply2(Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
			__local_var_1_0 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
			_ = __local_var_1_0
			// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
			__local_var_2_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
			_ = __local_var_2_1
			return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_1_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			})
		}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_22134120_138441832(Rebox_Data_List_138441832_22134120(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))})
	})
	return cache_Data_List_unzip
}

var cache_Data_List_uncons gopurs_runtime.Value
var once_Data_List_uncons sync.Once

func Get_Data_List_uncons() gopurs_runtime.Value {
	once_Data_List_uncons.Do(func() {
		cache_Data_List_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_uncons(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)})
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_uncons
}

var cache_Data_List_toUnfoldable gopurs_runtime.Value
var once_Data_List_toUnfoldable sync.Once

func Get_Data_List_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_toUnfoldable.Do(func() {
		cache_Data_List_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_List_toUnfoldable
}

var cache_Data_List_tail gopurs_runtime.Value
var once_Data_List_tail sync.Once

func Get_Data_List_tail() gopurs_runtime.Value {
	once_Data_List_tail.Do(func() {
		cache_Data_List_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_tail
}

var cache_Data_List_stripPrefix gopurs_runtime.Value
var once_Data_List_stripPrefix sync.Once

func Get_Data_List_stripPrefix() gopurs_runtime.Value {
	once_Data_List_stripPrefix.Do(func() {
		cache_Data_List_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_stripPrefix(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](s_2_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_stripPrefix
}

var cache_Data_List_span gopurs_runtime.Value
var once_Data_List_span sync.Once

func Get_Data_List_span() gopurs_runtime.Value {
	once_Data_List_span.Do(func() {
		cache_Data_List_span = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_span(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box))
				_ = orig
				return gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.rest)})
			}()
		})
	})
	return cache_Data_List_span
}

var cache_Data_List_span__2097567232 gopurs_runtime.Value
var once_Data_List_span__2097567232 sync.Once

func Get_Data_List_span__2097567232() gopurs_runtime.Value {
	once_Data_List_span__2097567232.Do(func() {
		cache_Data_List_span__2097567232 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_span__2097567232(v_0_box, Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box)))
				_ = orig
				return gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(orig.go__init))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(orig.rest))})
			}()
		})
	})
	return cache_Data_List_span__2097567232
}

var cache_Data_List_snoc gopurs_runtime.Value
var once_Data_List_snoc sync.Once

func Get_Data_List_snoc() gopurs_runtime.Value {
	once_Data_List_snoc.Do(func() {
		cache_Data_List_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_snoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0_box), x_1_box))}
		})
	})
	return cache_Data_List_snoc
}

var cache_Data_List_singleton gopurs_runtime.Value
var once_Data_List_singleton sync.Once

func Get_Data_List_singleton() gopurs_runtime.Value {
	once_Data_List_singleton.Do(func() {
		cache_Data_List_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_singleton(a_0_box))}
		})
	})
	return cache_Data_List_singleton
}

var cache_Data_List_singleton__115143311 gopurs_runtime.Value
var once_Data_List_singleton__115143311 sync.Once

func Get_Data_List_singleton__115143311() gopurs_runtime.Value {
	once_Data_List_singleton__115143311.Do(func() {
		cache_Data_List_singleton__115143311 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(Call_Data_List_singleton__115143311(a_0_box.IntVal)))}
		})
	})
	return cache_Data_List_singleton__115143311
}

var cache_Data_List_sortBy gopurs_runtime.Value
var once_Data_List_sortBy sync.Once

func Get_Data_List_sortBy() gopurs_runtime.Value {
	once_Data_List_sortBy.Do(func() {
		cache_Data_List_sortBy = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_sortBy(cmp_0_box)
		})
	})
	return cache_Data_List_sortBy
}

var cache_Data_List_sort gopurs_runtime.Value
var once_Data_List_sort sync.Once

func Get_Data_List_sort() gopurs_runtime.Value {
	once_Data_List_sort.Do(func() {
		cache_Data_List_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_sort
}

var cache_Data_List_tails gopurs_runtime.Value
var once_Data_List_tails sync.Once

func Get_Data_List_tails() gopurs_runtime.Value {
	once_Data_List_tails.Do(func() {
		cache_Data_List_tails = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Call_Data_List_tails(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))))}
		})
	})
	return cache_Data_List_tails
}

var cache_Data_List_showPattern gopurs_runtime.Value
var once_Data_List_showPattern sync.Once

func Get_Data_List_showPattern() gopurs_runtime.Value {
	once_Data_List_showPattern.Do(func() {
		cache_Data_List_showPattern = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_showPattern(dictShow_0_box)
		})
	})
	return cache_Data_List_showPattern
}

var cache_Data_List_reverse gopurs_runtime.Value
var once_Data_List_reverse sync.Once

func Get_Data_List_reverse() gopurs_runtime.Value {
	once_Data_List_reverse.Do(func() {
		cache_Data_List_reverse = func() gopurs_runtime.Value {
			var Call_local_Data_List_go__go_0_0_6 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			_ = Call_local_Data_List_go__go_0_0_6
			var go__go_0_0_6 gopurs_runtime.Value
			_ = go__go_0_0_6
			Call_local_Data_List_go__go_0_0_6 = func(v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
			go__go_0_0_6:
				for {
					if false {
						continue go__go_0_0_6
					}
					var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
					_ = v_1
					var v1_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_2_loop
					_ = v1_2
					var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					{
						if v1_2 == nil {
							__t1 = v_1
							goto end_branch_1
						} else {

						}
					}
					{
						if v1_2 != nil {
							v_1_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_2).V0, v_1})
							v1_2_loop = (v1_2).V1
							continue go__go_0_0_6
							__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_1:
					return __t1
				}
			}
			go__go_0_0_6 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_0_0_6(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_2_loop_val)))}
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
		}()
	})
	return cache_Data_List_reverse
}

var cache_Data_List_reverse__4012840474 gopurs_runtime.Value
var once_Data_List_reverse__4012840474 sync.Once

func Get_Data_List_reverse__4012840474() gopurs_runtime.Value {
	once_Data_List_reverse__4012840474.Do(func() {
		cache_Data_List_reverse__4012840474 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(Call_Data_List_reverse__4012840474(Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_0_box)))))}
		})
	})
	return cache_Data_List_reverse__4012840474
}

var cache_Data_List_take gopurs_runtime.Value
var once_Data_List_take sync.Once

func Get_Data_List_take() gopurs_runtime.Value {
	once_Data_List_take.Do(func() {
		cache_Data_List_take = func() gopurs_runtime.Value {
			var Call_local_Data_List_go__go_0_0_9 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			_ = Call_local_Data_List_go__go_0_0_9
			var go__go_0_0_9 gopurs_runtime.Value
			_ = go__go_0_0_9
			Call_local_Data_List_go__go_0_0_9 = func(v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_2_loop int64, v2_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
			go__go_0_0_9:
				for {
					if false {
						continue go__go_0_0_9
					}
					var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
					_ = v_1
					var v1_2 int64 = v1_2_loop
					_ = v1_2
					var v2_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_3_loop
					_ = v2_3
					var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					{
						if (v1_2) < (int64(1)) {
							var Call_local_Data_List_go__go_4_1_10 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							_ = Call_local_Data_List_go__go_4_1_10
							var go__go_4_1_10 gopurs_runtime.Value
							_ = go__go_4_1_10
							Call_local_Data_List_go__go_4_1_10 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
							go__go_4_1_10:
								for {
									if false {
										continue go__go_4_1_10
									}
									var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
									_ = v_5
									var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
									_ = v1_6
									var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
									{
										if v1_6 == nil {
											__t2 = v_5
											goto end_branch_2
										} else {

										}
									}
									{
										if v1_6 != nil {
											v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
											v1_6_loop = (v1_6).V1
											continue go__go_4_1_10
											__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
											goto end_branch_2
										} else {

										}
									}
									{
										__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
									}
								end_branch_2:
									return __t2
								}
							}
							go__go_4_1_10 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_10(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
								})
							})
							__t5 = Call_local_Data_List_go__go_4_1_10((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_1)
							goto end_branch_5
						} else {

						}
					}
					{
						if v2_3 == nil {
							var Call_local_Data_List_go__go_4_3_11 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							_ = Call_local_Data_List_go__go_4_3_11
							var go__go_4_3_11 gopurs_runtime.Value
							_ = go__go_4_3_11
							Call_local_Data_List_go__go_4_3_11 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
							go__go_4_3_11:
								for {
									if false {
										continue go__go_4_3_11
									}
									var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
									_ = v_5
									var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
									_ = v1_6
									var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
									{
										if v1_6 == nil {
											__t4 = v_5
											goto end_branch_4
										} else {

										}
									}
									{
										if v1_6 != nil {
											v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
											v1_6_loop = (v1_6).V1
											continue go__go_4_3_11
											__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
											goto end_branch_4
										} else {

										}
									}
									{
										__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
									}
								end_branch_4:
									return __t4
								}
							}
							go__go_4_3_11 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_3_11(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
								})
							})
							__t5 = Call_local_Data_List_go__go_4_3_11((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_1)
							goto end_branch_5
						} else {

						}
					}
					{
						if v2_3 != nil {
							v_1_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_3).V0, v_1})
							v1_2_loop = gopurs_runtime.IntSub(v1_2, int64(1))
							v2_3_loop = (v2_3).V1
							continue go__go_0_0_9
							__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_5:
					return __t5
				}
			}
			go__go_0_0_9 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_0_0_9(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_loop_val), v1_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_3_loop_val)))}
					})
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_9, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
		}()
	})
	return cache_Data_List_take
}

var cache_Data_List_takeWhile gopurs_runtime.Value
var once_Data_List_takeWhile sync.Once

func Get_Data_List_takeWhile() gopurs_runtime.Value {
	once_Data_List_takeWhile.Do(func() {
		cache_Data_List_takeWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_takeWhile(p_0_box)
		})
	})
	return cache_Data_List_takeWhile
}

var cache_Data_List_unsnoc gopurs_runtime.Value
var once_Data_List_unsnoc sync.Once

func Get_Data_List_unsnoc() gopurs_runtime.Value {
	once_Data_List_unsnoc.Do(func() {
		cache_Data_List_unsnoc = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_unsnoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](lst_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, orig.last)
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_unsnoc
}

var cache_Data_List_zipWith gopurs_runtime.Value
var once_Data_List_zipWith sync.Once

func Get_Data_List_zipWith() gopurs_runtime.Value {
	once_Data_List_zipWith.Do(func() {
		cache_Data_List_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2_box)))}
		})
	})
	return cache_Data_List_zipWith
}

var cache_Data_List_zip gopurs_runtime.Value
var once_Data_List_zip sync.Once

func Get_Data_List_zip() gopurs_runtime.Value {
	once_Data_List_zip.Do(func() {
		cache_Data_List_zip = gopurs_runtime.Apply(Get_Data_List_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_List_zip
}

var cache_Data_List_zipWithA gopurs_runtime.Value
var once_Data_List_zipWithA sync.Once

func Get_Data_List_zipWithA() gopurs_runtime.Value {
	once_Data_List_zipWithA.Do(func() {
		cache_Data_List_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_List_zipWithA
}

var cache_Data_List_go__range gopurs_runtime.Value
var once_Data_List_go__range sync.Once

func Get_Data_List_go__range() gopurs_runtime.Value {
	once_Data_List_go__range.Do(func() {
		cache_Data_List_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(Call_Data_List_go__range(start_0_box.IntVal, end_1_box.IntVal)))}
		})
	})
	return cache_Data_List_go__range
}

var cache_Data_List_partition gopurs_runtime.Value
var once_Data_List_partition sync.Once

func Get_Data_List_partition() gopurs_runtime.Value {
	once_Data_List_partition.Do(func() {
		cache_Data_List_partition = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_List_partition(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box))
				_ = orig
				return gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.no)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.yes)})
			}()
		})
	})
	return cache_Data_List_partition
}

var cache_Data_List_null gopurs_runtime.Value
var once_Data_List_null sync.Once

func Get_Data_List_null() gopurs_runtime.Value {
	once_Data_List_null.Do(func() {
		cache_Data_List_null = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_List_null(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box)))
		})
	})
	return cache_Data_List_null
}

var cache_Data_List_null__1968419687 gopurs_runtime.Value
var once_Data_List_null__1968419687 sync.Once

func Get_Data_List_null__1968419687() gopurs_runtime.Value {
	once_Data_List_null__1968419687.Do(func() {
		cache_Data_List_null__1968419687 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_List_null__1968419687(Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))))
		})
	})
	return cache_Data_List_null__1968419687
}

var cache_Data_List_nubBy gopurs_runtime.Value
var once_Data_List_nubBy sync.Once

func Get_Data_List_nubBy() gopurs_runtime.Value {
	once_Data_List_nubBy.Do(func() {
		cache_Data_List_nubBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_nubBy(p_0_box)
		})
	})
	return cache_Data_List_nubBy
}

var cache_Data_List_nub gopurs_runtime.Value
var once_Data_List_nub sync.Once

func Get_Data_List_nub() gopurs_runtime.Value {
	once_Data_List_nub.Do(func() {
		cache_Data_List_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_nub
}

var cache_Data_List_newtypePattern gopurs_runtime.Value
var once_Data_List_newtypePattern sync.Once

func Get_Data_List_newtypePattern() gopurs_runtime.Value {
	once_Data_List_newtypePattern.Do(func() {
		cache_Data_List_newtypePattern = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1501865320_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_List_newtypePattern
}

var cache_Data_List_mapMaybe gopurs_runtime.Value
var once_Data_List_mapMaybe sync.Once

func Get_Data_List_mapMaybe() gopurs_runtime.Value {
	once_Data_List_mapMaybe.Do(func() {
		cache_Data_List_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_mapMaybe(f_0_box)
		})
	})
	return cache_Data_List_mapMaybe
}

var cache_Data_List_manyRec gopurs_runtime.Value
var once_Data_List_manyRec sync.Once

func Get_Data_List_manyRec() gopurs_runtime.Value {
	once_Data_List_manyRec.Do(func() {
		cache_Data_List_manyRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_manyRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
		})
	})
	return cache_Data_List_manyRec
}

var cache_Data_List_someRec gopurs_runtime.Value
var once_Data_List_someRec sync.Once

func Get_Data_List_someRec() gopurs_runtime.Value {
	once_Data_List_someRec.Do(func() {
		cache_Data_List_someRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_someRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
		})
	})
	return cache_Data_List_someRec
}

var cache_Data_List_some gopurs_runtime.Value
var once_Data_List_some sync.Once

func Get_Data_List_some() gopurs_runtime.Value {
	once_Data_List_some.Do(func() {
		cache_Data_List_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_1_box), v_2_box)
		})
	})
	return cache_Data_List_some
}

var cache_Data_List_many gopurs_runtime.Value
var once_Data_List_many sync.Once

func Get_Data_List_many() gopurs_runtime.Value {
	once_Data_List_many.Do(func() {
		cache_Data_List_many = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_1_box), v_2_box)
		})
	})
	return cache_Data_List_many
}

var cache_Data_List_length gopurs_runtime.Value
var once_Data_List_length sync.Once

func Get_Data_List_length() gopurs_runtime.Value {
	once_Data_List_length.Do(func() {
		cache_Data_List_length = func() gopurs_runtime.Value {
			var Call_local_Data_List_go__go_0_0_23 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
			_ = Call_local_Data_List_go__go_0_0_23
			var go__go_0_0_23 gopurs_runtime.Value
			_ = go__go_0_0_23
			Call_local_Data_List_go__go_0_0_23 = func(b_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
			go__go_0_0_23:
				for {
					if false {
						continue go__go_0_0_23
					}
					var b_1 gopurs_runtime.Value = b_1_loop
					_ = b_1
					var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
					_ = v_2
					var __t1 gopurs_runtime.Value
					{
						if v_2 == nil {
							__t1 = b_1
							goto end_branch_1
						} else {

						}
					}
					{
						if v_2 != nil {
							b_1_loop = gopurs_runtime.Int(gopurs_runtime.IntAdd(b_1.IntVal, int64(1)))
							v_2_loop = (v_2).V1
							continue go__go_0_0_23
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
			go__go_0_0_23 = gopurs_runtime.Func(func(b_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Data_List_go__go_0_0_23(b_1_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val))
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_23, gopurs_runtime.Int(int64(0)))
		}()
	})
	return cache_Data_List_length
}

var cache_Data_List_last gopurs_runtime.Value
var once_Data_List_last sync.Once

func Get_Data_List_last() gopurs_runtime.Value {
	once_Data_List_last.Do(func() {
		cache_Data_List_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_last(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_last
}

var cache_Data_List_insertBy gopurs_runtime.Value
var once_Data_List_insertBy sync.Once

func Get_Data_List_insertBy() gopurs_runtime.Value {
	once_Data_List_insertBy.Do(func() {
		cache_Data_List_insertBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_insertBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box)))}
		})
	})
	return cache_Data_List_insertBy
}

var cache_Data_List_insertAt gopurs_runtime.Value
var once_Data_List_insertAt sync.Once

func Get_Data_List_insertAt() gopurs_runtime.Value {
	once_Data_List_insertAt.Do(func() {
		cache_Data_List_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_insertAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_insertAt
}

var cache_Data_List_insert gopurs_runtime.Value
var once_Data_List_insert sync.Once

func Get_Data_List_insert() gopurs_runtime.Value {
	once_Data_List_insert.Do(func() {
		cache_Data_List_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_insert
}

var cache_Data_List_go__init gopurs_runtime.Value
var once_Data_List_go__init sync.Once

func Get_Data_List_go__init() gopurs_runtime.Value {
	once_Data_List_go__init.Do(func() {
		cache_Data_List_go__init = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_go__init(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](lst_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_go__init
}

var cache_Data_List_index gopurs_runtime.Value
var once_Data_List_index sync.Once

func Get_Data_List_index() gopurs_runtime.Value {
	once_Data_List_index.Do(func() {
		cache_Data_List_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_index(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_index
}

var cache_Data_List_head gopurs_runtime.Value
var once_Data_List_head sync.Once

func Get_Data_List_head() gopurs_runtime.Value {
	once_Data_List_head.Do(func() {
		cache_Data_List_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_head(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_head
}

var cache_Data_List_transpose gopurs_runtime.Value
var once_Data_List_transpose sync.Once

func Get_Data_List_transpose() gopurs_runtime.Value {
	once_Data_List_transpose.Do(func() {
		cache_Data_List_transpose = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Call_Data_List_transpose(Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box)))))}
		})
	})
	return cache_Data_List_transpose
}

var cache_Data_List_groupBy gopurs_runtime.Value
var once_Data_List_groupBy sync.Once

func Get_Data_List_groupBy() gopurs_runtime.Value {
	once_Data_List_groupBy.Do(func() {
		cache_Data_List_groupBy = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_587109416_849153993(Call_Data_List_groupBy(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box))))}
		})
	})
	return cache_Data_List_groupBy
}

var cache_Data_List_groupAllBy gopurs_runtime.Value
var once_Data_List_groupAllBy sync.Once

func Get_Data_List_groupAllBy() gopurs_runtime.Value {
	once_Data_List_groupAllBy.Do(func() {
		cache_Data_List_groupAllBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_groupAllBy(p_0_box)
		})
	})
	return cache_Data_List_groupAllBy
}

var cache_Data_List_group gopurs_runtime.Value
var once_Data_List_group sync.Once

func Get_Data_List_group() gopurs_runtime.Value {
	once_Data_List_group.Do(func() {
		cache_Data_List_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_group
}

var cache_Data_List_groupAll gopurs_runtime.Value
var once_Data_List_groupAll sync.Once

func Get_Data_List_groupAll() gopurs_runtime.Value {
	once_Data_List_groupAll.Do(func() {
		cache_Data_List_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_List_groupAll
}

var cache_Data_List_fromFoldable gopurs_runtime.Value
var once_Data_List_fromFoldable sync.Once

func Get_Data_List_fromFoldable() gopurs_runtime.Value {
	once_Data_List_fromFoldable.Do(func() {
		cache_Data_List_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_List_fromFoldable
}

var cache_Data_List_foldM gopurs_runtime.Value
var once_Data_List_foldM sync.Once

func Get_Data_List_foldM() gopurs_runtime.Value {
	once_Data_List_foldM.Do(func() {
		cache_Data_List_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value, v2_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), v_1_box, v1_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_3_box))
		})
	})
	return cache_Data_List_foldM
}

var cache_Data_List_findIndex gopurs_runtime.Value
var once_Data_List_findIndex sync.Once

func Get_Data_List_findIndex() gopurs_runtime.Value {
	once_Data_List_findIndex.Do(func() {
		cache_Data_List_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_findIndex(fn_0_box)
		})
	})
	return cache_Data_List_findIndex
}

var cache_Data_List_findLastIndex gopurs_runtime.Value
var once_Data_List_findLastIndex sync.Once

func Get_Data_List_findLastIndex() gopurs_runtime.Value {
	once_Data_List_findLastIndex.Do(func() {
		cache_Data_List_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_findLastIndex(fn_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_findLastIndex
}

var cache_Data_List_filterM gopurs_runtime.Value
var once_Data_List_filterM sync.Once

func Get_Data_List_filterM() gopurs_runtime.Value {
	once_Data_List_filterM.Do(func() {
		cache_Data_List_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Data_List_filterM
}

var cache_Data_List_filter gopurs_runtime.Value
var once_Data_List_filter sync.Once

func Get_Data_List_filter() gopurs_runtime.Value {
	once_Data_List_filter.Do(func() {
		cache_Data_List_filter = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_filter(p_0_box)
		})
	})
	return cache_Data_List_filter
}

var cache_Data_List_intersectBy gopurs_runtime.Value
var once_Data_List_intersectBy sync.Once

func Get_Data_List_intersectBy() gopurs_runtime.Value {
	once_Data_List_intersectBy.Do(func() {
		cache_Data_List_intersectBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_intersectBy(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box)))}
		})
	})
	return cache_Data_List_intersectBy
}

var cache_Data_List_intersect gopurs_runtime.Value
var once_Data_List_intersect sync.Once

func Get_Data_List_intersect() gopurs_runtime.Value {
	once_Data_List_intersect.Do(func() {
		cache_Data_List_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_intersect
}

var cache_Data_List_nubByEq gopurs_runtime.Value
var once_Data_List_nubByEq sync.Once

func Get_Data_List_nubByEq() gopurs_runtime.Value {
	once_Data_List_nubByEq.Do(func() {
		cache_Data_List_nubByEq = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box)))}
		})
	})
	return cache_Data_List_nubByEq
}

var cache_Data_List_nubEq gopurs_runtime.Value
var once_Data_List_nubEq sync.Once

func Get_Data_List_nubEq() gopurs_runtime.Value {
	once_Data_List_nubEq.Do(func() {
		cache_Data_List_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_nubEq
}

var cache_Data_List_eqPattern gopurs_runtime.Value
var once_Data_List_eqPattern sync.Once

func Get_Data_List_eqPattern() gopurs_runtime.Value {
	once_Data_List_eqPattern.Do(func() {
		cache_Data_List_eqPattern = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_eqPattern(dictEq_0_box)
		})
	})
	return cache_Data_List_eqPattern
}

var cache_Data_List_ordPattern gopurs_runtime.Value
var once_Data_List_ordPattern sync.Once

func Get_Data_List_ordPattern() gopurs_runtime.Value {
	once_Data_List_ordPattern.Do(func() {
		cache_Data_List_ordPattern = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_ordPattern(dictOrd_0_box)
		})
	})
	return cache_Data_List_ordPattern
}

var cache_Data_List_elemLastIndex gopurs_runtime.Value
var once_Data_List_elemLastIndex sync.Once

func Get_Data_List_elemLastIndex() gopurs_runtime.Value {
	once_Data_List_elemLastIndex.Do(func() {
		cache_Data_List_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_List_elemLastIndex
}

var cache_Data_List_elemIndex gopurs_runtime.Value
var once_Data_List_elemIndex sync.Once

func Get_Data_List_elemIndex() gopurs_runtime.Value {
	once_Data_List_elemIndex.Do(func() {
		cache_Data_List_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_List_elemIndex
}

var cache_Data_List_dropWhile gopurs_runtime.Value
var once_Data_List_dropWhile sync.Once

func Get_Data_List_dropWhile() gopurs_runtime.Value {
	once_Data_List_dropWhile.Do(func() {
		cache_Data_List_dropWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_dropWhile(p_0_box)
		})
	})
	return cache_Data_List_dropWhile
}

var cache_Data_List_dropEnd gopurs_runtime.Value
var once_Data_List_dropEnd sync.Once

func Get_Data_List_dropEnd() gopurs_runtime.Value {
	once_Data_List_dropEnd.Do(func() {
		cache_Data_List_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_dropEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box)))}
		})
	})
	return cache_Data_List_dropEnd
}

var cache_Data_List_drop gopurs_runtime.Value
var once_Data_List_drop sync.Once

func Get_Data_List_drop() gopurs_runtime.Value {
	once_Data_List_drop.Do(func() {
		cache_Data_List_drop = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box)))}
		})
	})
	return cache_Data_List_drop
}

var cache_Data_List_slice gopurs_runtime.Value
var once_Data_List_slice sync.Once

func Get_Data_List_slice() gopurs_runtime.Value {
	once_Data_List_slice.Do(func() {
		cache_Data_List_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_slice(start_0_box.IntVal, end_1_box.IntVal, xs_2_box)
		})
	})
	return cache_Data_List_slice
}

var cache_Data_List_takeEnd gopurs_runtime.Value
var once_Data_List_takeEnd sync.Once

func Get_Data_List_takeEnd() gopurs_runtime.Value {
	once_Data_List_takeEnd.Do(func() {
		cache_Data_List_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_takeEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box)))}
		})
	})
	return cache_Data_List_takeEnd
}

var cache_Data_List_deleteBy gopurs_runtime.Value
var once_Data_List_deleteBy sync.Once

func Get_Data_List_deleteBy() gopurs_runtime.Value {
	once_Data_List_deleteBy.Do(func() {
		cache_Data_List_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box)))}
		})
	})
	return cache_Data_List_deleteBy
}

var cache_Data_List_unionBy gopurs_runtime.Value
var once_Data_List_unionBy sync.Once

func Get_Data_List_unionBy() gopurs_runtime.Value {
	once_Data_List_unionBy.Do(func() {
		cache_Data_List_unionBy = gopurs_runtime.Func3(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_unionBy(eq2_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2_box)))}
		})
	})
	return cache_Data_List_unionBy
}

var cache_Data_List_union gopurs_runtime.Value
var once_Data_List_union sync.Once

func Get_Data_List_union() gopurs_runtime.Value {
	once_Data_List_union.Do(func() {
		cache_Data_List_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_union
}

var cache_Data_List_deleteAt gopurs_runtime.Value
var once_Data_List_deleteAt sync.Once

func Get_Data_List_deleteAt() gopurs_runtime.Value {
	once_Data_List_deleteAt.Do(func() {
		cache_Data_List_deleteAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_deleteAt(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_deleteAt
}

var cache_Data_List_go__delete gopurs_runtime.Value
var once_Data_List_go__delete sync.Once

func Get_Data_List_go__delete() gopurs_runtime.Value {
	once_Data_List_go__delete.Do(func() {
		cache_Data_List_go__delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_go__delete
}

var cache_Data_List_difference gopurs_runtime.Value
var once_Data_List_difference sync.Once

func Get_Data_List_difference() gopurs_runtime.Value {
	once_Data_List_difference.Do(func() {
		cache_Data_List_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_List_difference
}

var cache_Data_List_concatMap gopurs_runtime.Value
var once_Data_List_concatMap sync.Once

func Get_Data_List_concatMap() gopurs_runtime.Value {
	once_Data_List_concatMap.Do(func() {
		cache_Data_List_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_concatMap(b_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](a_1_box)))}
		})
	})
	return cache_Data_List_concatMap
}

var cache_Data_List_concat gopurs_runtime.Value
var once_Data_List_concat sync.Once

func Get_Data_List_concat() gopurs_runtime.Value {
	once_Data_List_concat.Do(func() {
		cache_Data_List_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_concat(Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))))}
		})
	})
	return cache_Data_List_concat
}

var cache_Data_List_catMaybes gopurs_runtime.Value
var once_Data_List_catMaybes sync.Once

func Get_Data_List_catMaybes() gopurs_runtime.Value {
	once_Data_List_catMaybes.Do(func() {
		cache_Data_List_catMaybes = func() gopurs_runtime.Value {
			var Call_local_Data_List_go__go_0_0_47 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			_ = Call_local_Data_List_go__go_0_0_47
			var go__go_0_0_47 gopurs_runtime.Value
			_ = go__go_0_0_47
			Call_local_Data_List_go__go_0_0_47 = func(v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
			go__go_0_0_47:
				for {
					if false {
						continue go__go_0_0_47
					}
					var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
					_ = v_1
					var v1_2 *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] = v1_2_loop
					_ = v1_2
					var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					{
						if v1_2 == nil {
							var Call_local_Data_List_go__go_3_1_48 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							_ = Call_local_Data_List_go__go_3_1_48
							var go__go_3_1_48 gopurs_runtime.Value
							_ = go__go_3_1_48
							Call_local_Data_List_go__go_3_1_48 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
							go__go_3_1_48:
								for {
									if false {
										continue go__go_3_1_48
									}
									var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
									_ = v_4
									var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
									_ = v1_5
									var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
									{
										if v1_5 == nil {
											__t2 = v_4
											goto end_branch_2
										} else {

										}
									}
									{
										if v1_5 != nil {
											v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
											v1_5_loop = (v1_5).V1
											continue go__go_3_1_48
											__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
											goto end_branch_2
										} else {

										}
									}
									{
										__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
									}
								end_branch_2:
									return __t2
								}
							}
							go__go_3_1_48 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_1_48(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
								})
							})
							__t6 = Call_local_Data_List_go__go_3_1_48((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_1)
							goto end_branch_6
						} else {

						}
					}
					{
						if v1_2 != nil {
							var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							{
								var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v1_2).V0
								_ = __t_tag_3
								if __t_tag_3 == nil {
									v_1_loop = v_1
									v1_2_loop = (v1_2).V1
									continue go__go_0_0_47
									__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
									goto end_branch_5
								} else {

								}
							}
							{
								var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v1_2).V0
								_ = __t_tag_4
								if __t_tag_4 != nil {
									v_1_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v1_2).V0).V0, v_1})
									v1_2_loop = (v1_2).V1
									continue go__go_0_0_47
									__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
									goto end_branch_5
								} else {

								}
							}
							{
								__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
							}
						end_branch_5:
							__t6 = __t5
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_6:
					return __t6
				}
			}
			go__go_0_0_47 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_0_0_47(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_loop_val), Rebox_Data_List_849153993_1220287592(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_2_loop_val))))}
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_47, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
		}()
	})
	return cache_Data_List_catMaybes
}

var cache_Data_List_alterAt gopurs_runtime.Value
var once_Data_List_alterAt sync.Once

func Get_Data_List_alterAt() gopurs_runtime.Value {
	once_Data_List_alterAt.Do(func() {
		cache_Data_List_alterAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_List_alterAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_List_alterAt
}

var cache_Data_List_modifyAt gopurs_runtime.Value
var once_Data_List_modifyAt sync.Once

func Get_Data_List_modifyAt() gopurs_runtime.Value {
	once_Data_List_modifyAt.Do(func() {
		cache_Data_List_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_List_modifyAt(n_0_box.IntVal, f_1_box)
		})
	})
	return cache_Data_List_modifyAt
}

func Call_Data_List_tailRecM2(f_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], b_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = a_1_loop
	_ = a_1
	var b_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = b_2_loop
	_ = b_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_4130553207_1542299734(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecMaybe())).V1, gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(o_3, "a"), gopurs_runtime.RecordGet(o_3, "b"))
	}), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			/*gopurs:record-v1:2:0061,0062*/
		}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}}
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

func Call_Data_List_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Data_List_identity(x_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var x_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_Pattern(x_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var x_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_List_updateAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
updateAt:
	for {
		if false {
			continue updateAt
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
		_ = v2_2
		var __t3 gopurs_runtime.Value
		{
			if v2_2 != nil {
				var __t2 gopurs_runtime.Value
				{
					if (v_0) == (int64(0)) {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, (v2_2).V1}))}, true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))))}
						goto end_branch_2
					} else {

					}
				}
				{
					// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope1)])])
					__local_var_3_0 := Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := Call_Data_List_updateAt(gopurs_runtime.IntSub(v_0, int64(1)), v1_1, (v2_2).V1)
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					_ = __local_var_3_0
					var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					{
						if __local_var_3_0 != nil {
							__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, (__local_var_3_0).V0}))}, true}
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
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				}
			end_branch_2:
				__t3 = __t2
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
		}
	end_branch_3:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_List_uncons(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 struct {
		head gopurs_runtime.Value
		tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}
	V1 bool
} {
	var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[struct {
		head gopurs_runtime.Value
		tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]
	{
		if v_0 == nil {
			__t0 = Rebox_Data_List_3094389156_2681346401(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						head gopurs_runtime.Value
						tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
					}
					V1 bool
				}{struct {
					head gopurs_runtime.Value
					tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)})
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 != nil {
			__t0 = Rebox_Data_List_3094389156_2681346401(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						head gopurs_runtime.Value
						tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
					}
					V1 bool
				}{struct {
					head gopurs_runtime.Value
					tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{(v_0).V0, (v_0).V1}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)})
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Data_Maybe_Just[struct {
			head gopurs_runtime.Value
			tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}] {
			panic("Failed pattern match")
		}()
	}
end_branch_0:
	return func() struct {
		V0 struct {
			head gopurs_runtime.Value
			tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}
		V1 bool
	} {
		_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2681346401_3094389156(__t0))}
		if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 struct {
					head gopurs_runtime.Value
					tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}
				V1 bool
			}{V0: func() struct {
				head gopurs_runtime.Value
				tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			} {
				orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					head gopurs_runtime.Value
					tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
				}{}
				clone.head = gopurs_runtime.RecordGet(orig, "head")
				clone.tail = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "tail"))
				return clone
			}(), V1: true}
		}
		return struct {
			V0 struct {
				head gopurs_runtime.Value
				tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}
			V1 bool
		}{V0: struct {
			head gopurs_runtime.Value
			tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}, V1: false}
	}()
}

func Call_Data_List_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
	}))
}

func Call_Data_List_tail(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
	{
		if v_0 == nil {
			__t0 = Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
			__t0 = Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}, true}
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
		__t0 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
			panic("Failed pattern match")
		}()
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_stripPrefix(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], s_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
	_ = v_1
	var s_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = s_2_loop
	_ = s_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_4130553207_1542299734(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecMaybe())).V1, gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t6 gopurs_runtime.Value
		{
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "b")
			_ = __t_tag_0
			if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr != nil {
				var __t4 gopurs_runtime.Value
				{
					var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
					_ = __t_tag_1
					if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil {
						var __t2 gopurs_runtime.Value
						{
							if (gopurs_runtime.Apply2(dictEq_0.V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "a").UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "b").UnsafePtr).V0).IntVal) != (0) {
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "a").UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(o_3, "b").UnsafePtr).V1)})}))}}))}
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
						}
					end_branch_2:
						__t4 = __t2
						goto end_branch_4
					} else {

					}
				}
				{
					var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
					_ = __t_tag_3
					if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr == nil {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_3, "b")}))}}))}
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}
			end_branch_4:
				__t6 = __t4
				goto end_branch_6
			} else {

			}
		}
		{
			var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
			_ = __t_tag_5
			if __t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr == nil {
				__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_3, "b")}))}}))}
				goto end_branch_6
			} else {

			}
		}
		{
			__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		}
	end_branch_6:
		return __t6
	}), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			/*gopurs:record-v1:2:0061,0062*/
		}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(s_2)}}
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

func Call_Data_List_span(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest     *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
} {
span:
	for {
		if false {
			continue span
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
		_ = v1_1
		var __t1 struct {
			go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			rest     *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
		}
		{
			if (v1_1 != nil) && ((gopurs_runtime.Apply(v_0, (v1_1).V0).IntVal) != (0)) {
				// TAST (Let): v2_2_0 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a$scope43)]), rest: (ADT ["Data","List","Types","List"] [(TypeVar a$scope43)])] Empty))
				v2_2_0 := Call_Data_List_span(v_0, (v1_1).V1)
				_ = v2_2_0
				__t1 = struct {
					go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					rest     *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
				}{(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_1).V0, v2_2_0.go__init}), v2_2_0.rest}
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = struct {
				go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				rest     *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v1_1}
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_List_span__2097567232(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) struct {
	go__init *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	rest     *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
} {
span__2097567232:
	for {
		if false {
			continue span__2097567232
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v1_1_loop
		_ = v1_1
		var __t1 struct {
			go__init *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
			rest     *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
			/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
		}
		{
			if (v1_1 != nil) && ((gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3132786365_138441832((v1_1).V0))}).IntVal) != (0)) {
				// TAST (Let): v2_2_0 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Interval","Duration","DurationComponent"] []), Number])]), rest: (ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Interval","Duration","DurationComponent"] []), Number])])] Empty))
				v2_2_0 := Call_Data_List_span__2097567232(v_0, (v1_1).V1)
				_ = v2_2_0
				__t1 = struct {
					go__init *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
					rest     *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
					/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
				}{(&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, (v1_1).V0, v2_2_0.go__init}), v2_2_0.rest}
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = struct {
				go__init *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
				rest     *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
				/*gopurs:record-v1:2:0069006e00690074,0072006500730074*/
			}{(*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]])(nil), v1_1}
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_List_snoc(xs_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var xs_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_0_loop
	_ = xs_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, x_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_Data_List_singleton(a_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
}

func Call_Data_List_singleton__115143311(a_0_loop int64) *Constructor_Data_List_Types_Cons[int64] {
singleton__115143311:
	for {
		if false {
			continue singleton__115143311
		}
		var a_0 int64 = a_0_loop
		_ = a_0
		return (&Constructor_Data_List_Types_Cons[int64]{1, a_0, (*Constructor_Data_List_Types_Cons[int64])(nil)})
	}
}

func Call_Data_List_sortBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var cmp_0 gopurs_runtime.Value = cmp_0_loop
	_ = cmp_0
	var merge_1_0_0 gopurs_runtime.Value
	_ = merge_1_0_0
	var merge_1_0_0_cell *gopurs_runtime.Value
	_ = merge_1_0_0_cell
	// FALLBACK TCO: isLoop=false len=1
	merge_1_0_0 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t4 gopurs_runtime.Value
		{
			if v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil {
				var __t3 gopurs_runtime.Value
				{
					if v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil {
						var __t2 gopurs_runtime.Value
						{
							var __t_tag_1 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal)
							_ = __t_tag_1
							if uint32(__t_tag_1) == 380165415 {
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2((*merge_1_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}))}
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2((*merge_1_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3))}))}))}
						}
					end_branch_2:
						__t3 = __t2
						goto end_branch_3
					} else {

					}
				}
				{
					if v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil {
						__t3 = v_2
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
				}
			end_branch_3:
				__t4 = __t3
				goto end_branch_4
			} else {

			}
		}
		{
			if v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil {
				__t4 = v1_3
				goto end_branch_4
			} else {

			}
		}
		{
			if v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil {
				__t4 = v_2
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }())}
		}
	end_branch_4:
		return __t4
	})
	merge_1_0_0_cell = &merge_1_0_0
	var mergePairs_2_5_1 gopurs_runtime.Value
	_ = mergePairs_2_5_1
	var mergePairs_2_5_1_cell *gopurs_runtime.Value
	_ = mergePairs_2_5_1_cell
	// FALLBACK TCO: isLoop=false len=1
	mergePairs_2_5_1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t8 gopurs_runtime.Value
		{
			var __t_and_7 bool = false
			if v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil {

				var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1
				_ = __t_tag_6
				__t_and_7 = (__t_tag_6 != nil)
			}
			if __t_and_7 {
				__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(merge_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1).V0))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply((*mergePairs_2_5_1_cell), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1).V1)}))}))}
				goto end_branch_8
			} else {

			}
		}
		{
			__t8 = v_3
		}
	end_branch_8:
		return __t8
	})
	mergePairs_2_5_1_cell = &mergePairs_2_5_1
	var Call_local_Data_List_mergeAll_3_9_2 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_mergeAll_3_9_2
	var mergeAll_3_9_2 gopurs_runtime.Value
	_ = mergeAll_3_9_2
	Call_local_Data_List_mergeAll_3_9_2 = func(v_4_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	mergeAll_3_9_2:
		for {
			if false {
				continue mergeAll_3_9_2
			}
			var v_4 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_4_loop
			_ = v_4
			var __t12 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				var __t_and_11 bool = false
				if v_4 != nil {

					var __t_tag_10 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = (v_4).V1
					_ = __t_tag_10
					__t_and_11 = (__t_tag_10 == nil)
				}
				if __t_and_11 {
					__t12 = (v_4).V0
					goto end_branch_12
				} else {

				}
			}
			{
				v_4_loop = Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(mergePairs_2_5_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(v_4))})))
				continue mergeAll_3_9_2
				__t12 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
			}
		end_branch_12:
			return __t12
		}
	}
	mergeAll_3_9_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_mergeAll_3_9_2(Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))))}
	})
	var Call_local_Data_List_sequences_4_13_3 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
	_ = Call_local_Data_List_sequences_4_13_3
	var sequences_4_13_3 gopurs_runtime.Value
	_ = sequences_4_13_3
	var Call_local_Data_List_descending_4_14_4 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
	_ = Call_local_Data_List_descending_4_14_4
	var descending_4_14_4 gopurs_runtime.Value
	_ = descending_4_14_4
	var Call_local_Data_List_ascending_4_15_5 func(gopurs_runtime.Value, gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
	_ = Call_local_Data_List_ascending_4_15_5
	var ascending_4_15_5 gopurs_runtime.Value
	_ = ascending_4_15_5
	Call_local_Data_List_sequences_4_13_3 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	sequences_4_13_3:
		for {
			if false {
				continue sequences_4_13_3
			}
			var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
			_ = v_5
			var __t21 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
			{
				var __t_and_17 bool = false
				if v_5 != nil {

					var __t_tag_16 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_5).V1
					_ = __t_tag_16
					__t_and_17 = (__t_tag_16 != nil)
				}
				if __t_and_17 {
					var __t20 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
					{
						var __t_tag_19 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, (v_5).V0, ((v_5).V1).V0).IntVal)
						_ = __t_tag_19
						if uint32(__t_tag_19) == 380165415 {
							__t20 = Call_local_Data_List_descending_4_14_4(((v_5).V1).V0, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}), ((v_5).V1).V1)
							goto end_branch_20
						} else {

						}
					}
					{
						// TAST (Let): __local_var_6_18 shape=Other bindingType=Any
						__local_var_6_18 := (v_5).V0
						_ = __local_var_6_18
						__t20 = Call_local_Data_List_ascending_4_15_5(((v_5).V1).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_18, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7)}))}
						}), ((v_5).V1).V1)
					}
				end_branch_20:
					__t21 = __t20
					goto end_branch_21
				} else {

				}
			}
			{
				__t21 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, v_5, (*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(nil)})
			}
		end_branch_21:
			return __t21
		}
	}
	sequences_4_13_3 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Call_local_Data_List_sequences_4_13_3(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))))}
	})
	Call_local_Data_List_descending_4_14_4 = func(v_5_loop gopurs_runtime.Value, v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	descending_4_14_4:
		for {
			if false {
				continue descending_4_14_4
			}
			var v_5 gopurs_runtime.Value = v_5_loop
			_ = v_5
			var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
			_ = v1_6
			var v2_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_7_loop
			_ = v2_7
			var __t24 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
			{
				var __t_and_23 bool = false
				if v2_7 != nil {

					var __t_tag_22 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, v_5, (v2_7).V0).IntVal)
					_ = __t_tag_22
					__t_and_23 = (uint32(__t_tag_22) == 380165415)
				}
				if __t_and_23 {
					v_5_loop = (v2_7).V0
					v1_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v_5, v1_6})
					v2_7_loop = (v2_7).V1
					continue descending_4_14_4
					__t24 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
						panic("unreachable")
					}()
					goto end_branch_24
				} else {

				}
			}
			{
				__t24 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v_5, v1_6}), Call_local_Data_List_sequences_4_13_3(v2_7)})
			}
		end_branch_24:
			return __t24
		}
	}
	descending_4_14_4 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Call_local_Data_List_descending_4_14_4(v_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_7_loop_val))))}
			})
		})
	})
	Call_local_Data_List_ascending_4_15_5 = func(v_5_loop gopurs_runtime.Value, v1_6_loop gopurs_runtime.Value, v2_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	ascending_4_15_5:
		for {
			if false {
				continue ascending_4_15_5
			}
			var v_5 gopurs_runtime.Value = v_5_loop
			_ = v_5
			var v1_6 gopurs_runtime.Value = v1_6_loop
			_ = v1_6
			var v2_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_7_loop
			_ = v2_7
			var __t27 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
			{
				var __t_and_26 bool = false
				if v2_7 != nil {

					var __t_tag_25 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, v_5, (v2_7).V0).IntVal)
					_ = __t_tag_25
					__t_and_26 = (uint32(__t_tag_25) == 380165415) != (true)
				}
				if __t_and_26 {
					v_5_loop = (v2_7).V0
					v1_6_loop = gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_8)}))})
					})
					v2_7_loop = (v2_7).V1
					continue ascending_4_15_5
					__t27 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
						panic("unreachable")
					}()
					goto end_branch_27
				} else {

				}
			}
			{
				__t27 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v_5, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))})), Call_local_Data_List_sequences_4_13_3(v2_7)})
			}
		end_branch_27:
			return __t27
		}
	}
	ascending_4_15_5 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Call_local_Data_List_ascending_4_15_5(v_5_loop_val, v1_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_7_loop_val))))}
			})
		})
	})
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), mergeAll_3_9_2, sequences_4_13_3)
}

func Call_Data_List_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): compare_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope69), (TypeVar a$scope69)] (ADT ["Data","Ordering","Ordering"] []))
	compare_1_0 := Call_Data_Ord_compare(dictOrd_0)
	_ = compare_1_0
	return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Call_Data_List_sortBy(compare_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))})
	})
}

func Call_Data_List_tails(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
tails:
	for {
		if false {
			continue tails
		}
		var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
		_ = v_0
		var __t0 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
		{
			if v_0 == nil {
				__t0 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), (*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(nil)})
				goto end_branch_0
			} else {

			}
		}
		{
			if v_0 != nil {
				__t0 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, v_0, Call_Data_List_tails((v_0).V1)})
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
				panic("Failed pattern match")
			}()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Data_List_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3351995458_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Pattern ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_List_Types_showList(dictShow_0), "show"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1))}).StrVal()), ")"))
	})})))}
}

func Call_Data_List_reverse__4012840474(__eta_norm_0_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
reverse__4012840474:
	for {
		if false {
			continue reverse__4012840474
		}
		var __eta_norm_0_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		var Call_local_Data_List_go__952457181_1_0_7 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
		_ = Call_local_Data_List_go__952457181_1_0_7
		var go__952457181_1_0_7 gopurs_runtime.Value
		_ = go__952457181_1_0_7
		var Call_local_Data_List_go__go_1_1_8 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
		_ = Call_local_Data_List_go__go_1_1_8
		var go__go_1_1_8 gopurs_runtime.Value
		_ = go__go_1_1_8
		Call_local_Data_List_go__952457181_1_0_7 = func(v_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], v1_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
		go__952457181_1_0_7:
			for {
				if false {
					continue go__952457181_1_0_7
				}
				var v_2 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v_2_loop
				_ = v_2
				var v1_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v1_3_loop
				_ = v1_3
				var __t2 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
				{
					if v1_3 == nil {
						__t2 = v_2
						goto end_branch_2
					} else {

					}
				}
				{
					if v1_3 != nil {
						v_2_loop = (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, (v1_3).V0, v_2})
						v1_3_loop = (v1_3).V1
						continue go__952457181_1_0_7
						__t2 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
							panic("unreachable")
						}()
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
						panic("Failed pattern match")
					}()
				}
			end_branch_2:
				return __t2
			}
		}
		go__952457181_1_0_7 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(Call_local_Data_List_go__952457181_1_0_7(Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)), Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))))}
			})
		})
		Call_local_Data_List_go__go_1_1_8 = func(v_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], v1_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
		go__go_1_1_8:
			for {
				if false {
					continue go__go_1_1_8
				}
				var v_2 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v_2_loop
				_ = v_2
				var v1_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v1_3_loop
				_ = v1_3
				var __t3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
				{
					if v1_3 == nil {
						__t3 = v_2
						goto end_branch_3
					} else {

					}
				}
				{
					if v1_3 != nil {
						__t3 = Call_local_Data_List_go__952457181_1_0_7((&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, (v1_3).V0, v_2}), (v1_3).V1)
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
						panic("Failed pattern match")
					}()
				}
			end_branch_3:
				return __t3
			}
		}
		go__go_1_1_8 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(Call_local_Data_List_go__go_1_1_8(Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)), Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))))}
			})
		})
		return Call_local_Data_List_go__952457181_1_0_7((*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]])(nil), __eta_norm_0_0)
	}
}

func Call_Data_List_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var Call_local_Data_List_go__go_1_0_12 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_go__go_1_0_12
	var go__go_1_0_12 gopurs_runtime.Value
	_ = go__go_1_0_12
	Call_local_Data_List_go__go_1_0_12 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	go__go_1_0_12:
		for {
			if false {
				continue go__go_1_0_12
			}
			var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				if (v1_3 != nil) && ((gopurs_runtime.Apply(p_0, (v1_3).V0).IntVal) != (0)) {
					v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, v_2})
					v1_3_loop = (v1_3).V1
					continue go__go_1_0_12
					__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
					goto end_branch_3
				} else {

				}
			}
			{
				var Call_local_Data_List_go__go_4_1_13 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				_ = Call_local_Data_List_go__go_4_1_13
				var go__go_4_1_13 gopurs_runtime.Value
				_ = go__go_4_1_13
				Call_local_Data_List_go__go_4_1_13 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
				go__go_4_1_13:
					for {
						if false {
							continue go__go_4_1_13
						}
						var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
						_ = v_5
						var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
						_ = v1_6
						var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						{
							if v1_6 == nil {
								__t2 = v_5
								goto end_branch_2
							} else {

							}
						}
						{
							if v1_6 != nil {
								v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
								v1_6_loop = (v1_6).V1
								continue go__go_4_1_13
								__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
						}
					end_branch_2:
						return __t2
					}
				}
				go__go_4_1_13 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_13(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
					})
				})
				__t3 = Call_local_Data_List_go__go_4_1_13((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
			}
		end_branch_3:
			return __t3
		}
	}
	go__go_1_0_12 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_12(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
		})
	})
	return gopurs_runtime.Apply(go__go_1_0_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_unsnoc(lst_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 struct {
		go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		last     gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
	}
	V1 bool
} {
	var lst_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = lst_0_loop
	_ = lst_0
	var Call_local_Data_List_go__go_1_0_14 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
		last    gopurs_runtime.Value
		revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
	}]
	_ = Call_local_Data_List_go__go_1_0_14
	var go__go_1_0_14 gopurs_runtime.Value
	_ = go__go_1_0_14
	Call_local_Data_List_go__go_1_0_14 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
		last    gopurs_runtime.Value
		revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
	}] {
	go__go_1_0_14:
		for {
			if false {
				continue go__go_1_0_14
			}
			var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t3 *Constructor_Data_Maybe_Just[struct {
				last    gopurs_runtime.Value
				revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
			}]
			{
				if v_2 == nil {
					__t3 = Rebox_Data_List_3094389156_3270864776(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
					goto end_branch_3
				} else {

				}
			}
			{
				if v_2 != nil {
					var __t2 *Constructor_Data_Maybe_Just[struct {
						last    gopurs_runtime.Value
						revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
					}]
					{
						var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_2).V1
						_ = __t_tag_1
						if __t_tag_1 == nil {
							__t2 = Rebox_Data_List_3094389156_3270864776(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 struct {
										last    gopurs_runtime.Value
										revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
										/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
									}
									V1 bool
								}{struct {
									last    gopurs_runtime.Value
									revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
									/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
								}{(v_2).V0, v1_3}, true}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
										orig := _v.V0
										_ = orig
										return gopurs_runtime.RecordDict2("last", "revInit", orig.last, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.revInit)})
									}()})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}()))
							goto end_branch_2
						} else {

						}
					}
					{
						v_2_loop = (v_2).V1
						v1_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, v1_3})
						continue go__go_1_0_14
						__t2 = func() *Constructor_Data_Maybe_Just[struct {
							last    gopurs_runtime.Value
							revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
						}] {
							panic("unreachable")
						}()
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() *Constructor_Data_Maybe_Just[struct {
					last    gopurs_runtime.Value
					revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
				}] {
					panic("Failed pattern match")
				}()
			}
		end_branch_3:
			return __t3
		}
	}
	go__go_1_0_14 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3270864776_3094389156(Call_local_Data_List_go__go_1_0_14(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val))))}
		})
	})
	// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [revInit: (ADT ["Data","List","Types","List"] [(TypeVar a$scope91)]), last: (TypeVar a$scope91)] Empty))])
	__local_var_2_4 := Call_local_Data_List_go__go_1_0_14(lst_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))
	_ = __local_var_2_4
	var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_2_4 != nil {
			var Call_local_Data_List_go__go_3_5_15 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			_ = Call_local_Data_List_go__go_3_5_15
			var go__go_3_5_15 gopurs_runtime.Value
			_ = go__go_3_5_15
			Call_local_Data_List_go__go_3_5_15 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
			go__go_3_5_15:
				for {
					if false {
						continue go__go_3_5_15
					}
					var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
					_ = v_4
					var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
					_ = v1_5
					var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					{
						if v1_5 == nil {
							__t6 = v_4
							goto end_branch_6
						} else {

						}
					}
					{
						if v1_5 != nil {
							v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
							v1_5_loop = (v1_5).V1
							continue go__go_3_5_15
							__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_6:
					return __t6
				}
			}
			go__go_3_5_15 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_5_15(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
				})
			})
			__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_5_15((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (__local_var_2_4).V0.revInit))}, (__local_var_2_4).V0.last), true}
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
	return func() struct {
		V0 struct {
			go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			last     gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
		}
		V1 bool
	} {
		_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
		if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 struct {
					go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					last     gopurs_runtime.Value
					/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
				}
				V1 bool
			}{V0: func() struct {
				go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				last     gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
			} {
				orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					last     gopurs_runtime.Value
					/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
				}{}
				clone.go__init = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "init"))
				clone.last = gopurs_runtime.RecordGet(orig, "last")
				return clone
			}(), V1: true}
		}
		return struct {
			V0 struct {
				go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				last     gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
			}
			V1 bool
		}{V0: struct {
			go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			last     gopurs_runtime.Value
			/*gopurs:record-v1:2:0069006e00690074,006c006100730074*/
		}{}, V1: false}
	}()
}

func Call_Data_List_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	var ys_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_2_loop
	_ = ys_2
	var Call_local_Data_List_go__go_3_0_16 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_go__go_3_0_16
	var go__go_3_0_16 gopurs_runtime.Value
	_ = go__go_3_0_16
	Call_local_Data_List_go__go_3_0_16 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	go__go_3_0_16:
		for {
			if false {
				continue go__go_3_0_16
			}
			var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
			_ = v_4
			var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
			_ = v1_5
			var v2_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_6_loop
			_ = v2_6
			var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				if v_4 == nil {
					__t1 = v2_6
					goto end_branch_1
				} else {

				}
			}
			{
				if v1_5 == nil {
					__t1 = v2_6
					goto end_branch_1
				} else {

				}
			}
			{
				if (v_4 != nil) && (v1_5 != nil) {
					v_4_loop = (v_4).V1
					v1_5_loop = (v1_5).V1
					v2_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (v_4).V0, (v1_5).V0), v2_6})
					continue go__go_3_0_16
					__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		}
	}
	go__go_3_0_16 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_0_16(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_6_loop_val)))}
			})
		})
	})
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_reverse(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_0_16(xs_1, ys_2, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))}))
}

func Call_Data_List_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): sequence1_1_0 shape=App(Other) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeApp (TypeVar m$scope114) [(TypeVar c$scope117)])])] (TypeApp (TypeVar m$scope114) [(ADT ["Data","List","Types","List"] [(TypeVar c$scope117)])]))
	sequence1_1_0 := gopurs_runtime.Apply2(Rebox_Data_List_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
	_ = sequence1_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith(f_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_3), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_4)))})
	})
}

func Call_Data_List_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Data_List_Types_Cons[int64] {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	var __t5 *Constructor_Data_List_Types_Cons[int64]
	{
		if (start_0) == (end_1) {
			__t5 = (&Constructor_Data_List_Types_Cons[int64]{1, start_0, (*Constructor_Data_List_Types_Cons[int64])(nil)})
			goto end_branch_5
		} else {

		}
	}
	{
		var Call_local_Data_List_go__535684919_2_0_17 func(int64, int64, int64, *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64]
		_ = Call_local_Data_List_go__535684919_2_0_17
		var go__535684919_2_0_17 gopurs_runtime.Value
		_ = go__535684919_2_0_17
		var Call_local_Data_List_go__go_2_1_18 func(int64, int64, int64, *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64]
		_ = Call_local_Data_List_go__go_2_1_18
		var go__go_2_1_18 gopurs_runtime.Value
		_ = go__go_2_1_18
		Call_local_Data_List_go__535684919_2_0_17 = func(s_3_loop int64, e_4_loop int64, step_5_loop int64, rest_6_loop *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64] {
		go__535684919_2_0_17:
			for {
				if false {
					continue go__535684919_2_0_17
				}
				var s_3 int64 = s_3_loop
				_ = s_3
				var e_4 int64 = e_4_loop
				_ = e_4
				var step_5 int64 = step_5_loop
				_ = step_5
				var rest_6 *Constructor_Data_List_Types_Cons[int64] = rest_6_loop
				_ = rest_6
				var __t2 *Constructor_Data_List_Types_Cons[int64]
				{
					if (s_3) == (e_4) {
						__t2 = (&Constructor_Data_List_Types_Cons[int64]{1, s_3, rest_6})
						goto end_branch_2
					} else {

					}
				}
				{
					s_3_loop = gopurs_runtime.IntAdd(s_3, step_5)
					e_4_loop = e_4
					step_5_loop = step_5
					rest_6_loop = (&Constructor_Data_List_Types_Cons[int64]{1, s_3, rest_6})
					continue go__535684919_2_0_17
					__t2 = func() *Constructor_Data_List_Types_Cons[int64] { panic("unreachable") }()
				}
			end_branch_2:
				return __t2
			}
		}
		go__535684919_2_0_17 = gopurs_runtime.Func(func(s_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(e_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(step_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rest_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(Call_local_Data_List_go__535684919_2_0_17(s_3_loop_val.IntVal, e_4_loop_val.IntVal, step_5_loop_val.IntVal, Rebox_Data_List_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](rest_6_loop_val)))))}
					})
				})
			})
		})
		Call_local_Data_List_go__go_2_1_18 = func(s_3_loop int64, e_4_loop int64, step_5_loop int64, rest_6_loop *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64] {
		go__go_2_1_18:
			for {
				if false {
					continue go__go_2_1_18
				}
				var s_3 int64 = s_3_loop
				_ = s_3
				var e_4 int64 = e_4_loop
				_ = e_4
				var step_5 int64 = step_5_loop
				_ = step_5
				var rest_6 *Constructor_Data_List_Types_Cons[int64] = rest_6_loop
				_ = rest_6
				var __t3 *Constructor_Data_List_Types_Cons[int64]
				{
					if (s_3) == (e_4) {
						__t3 = (&Constructor_Data_List_Types_Cons[int64]{1, s_3, rest_6})
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = Call_local_Data_List_go__535684919_2_0_17(gopurs_runtime.IntAdd(s_3, step_5), e_4, step_5, (&Constructor_Data_List_Types_Cons[int64]{1, s_3, rest_6}))
				}
			end_branch_3:
				return __t3
			}
		}
		go__go_2_1_18 = gopurs_runtime.Func(func(s_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(e_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(step_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rest_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(Call_local_Data_List_go__go_2_1_18(s_3_loop_val.IntVal, e_4_loop_val.IntVal, step_5_loop_val.IntVal, Rebox_Data_List_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](rest_6_loop_val)))))}
					})
				})
			})
		})
		var __t4 int64
		{
			if (start_0) > (end_1) {
				__t4 = int64(1)
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = int64(-1)
		}
	end_branch_4:
		__t5 = Call_local_Data_List_go__535684919_2_0_17(end_1, start_0, __t4, (*Constructor_Data_List_Types_Cons[int64])(nil))
	}
end_branch_5:
	return __t5
}

func Call_Data_List_partition(p_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	no  *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	/*gopurs:record-v1:2:006e006f,007900650073*/
} {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	return func() struct {
		no  *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:006e006f,007900650073*/
	} {
		orig := gopurs_runtime.Apply3(Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if (gopurs_runtime.Apply(p_0, x_2).IntVal) != (0) {
					__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordGet(v_3, "no"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "yes"))}))})
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "no"))}))}, gopurs_runtime.RecordGet(v_3, "yes"))
			}
		end_branch_0:
			return __t0
		}), func() gopurs_runtime.Value {
			orig := struct {
				no  *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:006e006f,007900650073*/
			}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}
			_ = orig
			return gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.no)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.yes)})
		}(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})
		_ = orig
		clone := struct {
			no  *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:006e006f,007900650073*/
		}{}
		clone.no = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "no"))
		clone.yes = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "yes"))
		return clone
	}()
}

func Call_Data_List_null(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) bool {
	var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (v_0 == nil)
}

func Call_Data_List_null__1968419687(v_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) bool {
null__1968419687:
	for {
		if false {
			continue null__1968419687
		}
		var v_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v_0_loop
		_ = v_0
		return (v_0 == nil)
	}
}

func Call_Data_List_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var Call_local_Data_List_go__go_1_0_19 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_go__go_1_0_19
	var go__go_1_0_19 gopurs_runtime.Value
	_ = go__go_1_0_19
	Call_local_Data_List_go__go_1_0_19 = func(v_2_loop gopurs_runtime.Value, v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	go__go_1_0_19:
		for {
			if false {
				continue go__go_1_0_19
			}
			var v_2 gopurs_runtime.Value = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var v2_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_4_loop
			_ = v2_4
			var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				if v2_4 == nil {
					__t3 = v1_3
					goto end_branch_3
				} else {

				}
			}
			{
				if v2_4 != nil {
					// TAST (Let): v3_5_1 shape=App(Var) bindingType=(Record (Row [found: Boolean, result: (ADT ["Data","List","Internal","Set"] [(TypeVar a$scope133)])] Empty))
					v3_5_1 := Call_Data_List_Internal_insertAndLookupBy(p_0, (v2_4).V0, v_2)
					_ = v3_5_1
					var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					{
						if v3_5_1.found {
							v_2_loop = v3_5_1.result
							v1_3_loop = v1_3
							v2_4_loop = (v2_4).V1
							continue go__go_1_0_19
							__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_2
						} else {

						}
					}
					{
						v_2_loop = v3_5_1.result
						v1_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_4).V0, v1_3})
						v2_4_loop = (v2_4).V1
						continue go__go_1_0_19
						__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		}
	}
	go__go_1_0_19 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_19(v_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_4_loop_val)))}
			})
		})
	})
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_reverse(), gopurs_runtime.Apply2(go__go_1_0_19, Get_Data_List_Internal_emptySet(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_Data_List_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_List_nubBy(Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_List_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var Call_local_Data_List_go__go_1_0_20 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_go__go_1_0_20
	var go__go_1_0_20 gopurs_runtime.Value
	_ = go__go_1_0_20
	Call_local_Data_List_go__go_1_0_20 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	go__go_1_0_20:
		for {
			if false {
				continue go__go_1_0_20
			}
			var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				if v1_3 == nil {
					var Call_local_Data_List_go__go_4_1_21 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					_ = Call_local_Data_List_go__go_4_1_21
					var go__go_4_1_21 gopurs_runtime.Value
					_ = go__go_4_1_21
					Call_local_Data_List_go__go_4_1_21 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
					go__go_4_1_21:
						for {
							if false {
								continue go__go_4_1_21
							}
							var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
							_ = v_5
							var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
							_ = v1_6
							var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							{
								if v1_6 == nil {
									__t2 = v_5
									goto end_branch_2
								} else {

								}
							}
							{
								if v1_6 != nil {
									v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
									v1_6_loop = (v1_6).V1
									continue go__go_4_1_21
									__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
									goto end_branch_2
								} else {

								}
							}
							{
								__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
							}
						end_branch_2:
							return __t2
						}
					}
					go__go_4_1_21 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_21(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
						})
					})
					__t5 = Call_local_Data_List_go__go_4_1_21((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
					goto end_branch_5
				} else {

				}
			}
			{
				if v1_3 != nil {
					// TAST (Let): v2_4_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope147)])
					v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (v1_3).V0))
					_ = v2_4_3
					var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					{
						if v2_4_3 == nil {
							v_2_loop = v_2
							v1_3_loop = (v1_3).V1
							continue go__go_1_0_20
							__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_4
						} else {

						}
					}
					{
						if v2_4_3 != nil {
							v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_4_3).V0, v_2})
							v1_3_loop = (v1_3).V1
							continue go__go_1_0_20
							__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_5:
			return __t5
		}
	}
	go__go_1_0_20 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_20(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
		})
	})
	return gopurs_runtime.Apply(go__go_1_0_20, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_manyRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictAlternative_1_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictAlternative_1 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
	_ = dictAlternative_1
	// TAST (Let): Plus1_2_0 shape=App(Other) bindingType=Any
	Plus1_2_0 := gopurs_runtime.Apply(dictAlternative_1.V1, gopurs_runtime.Value{})
	_ = Plus1_2_0
	// TAST (Let): pure1_3_1 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope153)]), (ADT ["Data","List","Types","List"] [(TypeVar a$scope153)])])] (TypeApp (TypeVar f$scope152) [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope153)]), (ADT ["Data","List","Types","List"] [(TypeVar a$scope153)])])]))
	pure1_3_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_1.V0, gopurs_runtime.Value{})))
	_ = pure1_3_1
	return gopurs_runtime.Func(func(p_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_2_0, "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_2_0, "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Control_Monad_Rec_Class_Loop(), p_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_1.V0, gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()}))})), gopurs_runtime.Func(func(aa_6 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if aa_6.Type == 9 && aa_6.IntVal == 525585346 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(aa_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_5)}))}}))}
						goto end_branch_4
					} else {

					}
				}
				{
					if aa_6.Type == 9 && aa_6.IntVal == 60402430 {
						var Call_local_Data_List_go__go_7_2_22 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						_ = Call_local_Data_List_go__go_7_2_22
						var go__go_7_2_22 gopurs_runtime.Value
						_ = go__go_7_2_22
						Call_local_Data_List_go__go_7_2_22 = func(v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_9_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
						go__go_7_2_22:
							for {
								if false {
									continue go__go_7_2_22
								}
								var v_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_8_loop
								_ = v_8
								var v1_9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_9_loop
								_ = v1_9
								var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
								{
									if v1_9 == nil {
										__t3 = v_8
										goto end_branch_3
									} else {

									}
								}
								{
									if v1_9 != nil {
										v_8_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_9).V0, v_8})
										v1_9_loop = (v1_9).V1
										continue go__go_7_2_22
										__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
										goto end_branch_3
									} else {

									}
								}
								{
									__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
								}
							end_branch_3:
								return __t3
							}
						}
						go__go_7_2_22 = gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v1_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_7_2_22(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_9_loop_val)))}
							})
						})
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_7_2_22((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_5)))}}))}
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return gopurs_runtime.Apply(pure1_3_1, __t4)
			}))
		}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
}

func Call_Data_List_someRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictAlternative_1_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictAlternative_1 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
	_ = dictAlternative_1
	// TAST (Let): manyRec2_2_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope162) [(TypeVar a$scope163)])] (TypeApp (TypeVar f$scope162) [(ADT ["Data","List","Types","List"] [(TypeVar a$scope163)])]))
	manyRec2_2_0 := Call_Data_List_manyRec(dictMonadRec_0, dictAlternative_1)
	_ = manyRec2_2_0
	return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_1.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_1.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Types_Cons(), v_3), gopurs_runtime.Apply(manyRec2_2_0, v_3))
	})
}

func Call_Data_List_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], dictLazy_1_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var dictLazy_1 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_1_loop
	_ = dictLazy_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Types_Cons(), v_2), gopurs_runtime.Apply(dictLazy_1.V0, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_List_many(dictAlternative_0, dictLazy_1, v_2)
	})))
}

func Call_Data_List_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], dictLazy_1_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var dictLazy_1 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_1_loop
	_ = dictLazy_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_Data_List_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))}))
}

func Call_Data_List_last(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
last:
	for {
		if false {
			continue last
		}
		var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
		_ = v_0
		var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if v_0 != nil {
				var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				{
					var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_0).V1
					_ = __t_tag_0
					if __t_tag_0 == nil {
						__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{(v_0).V0, true}
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
					v_0_loop = (v_0).V1
					continue last
					__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	end_branch_2:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_List_insertBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
insertBy:
	for {
		if false {
			continue insertBy
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
		_ = v2_2
		var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		{
			if v2_2 == nil {
				__t2 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2 != nil {
				var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				{
					var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0).IntVal)
					_ = __t_tag_0
					if uint32(__t_tag_0) == 380165415 {
						__t1 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, Call_Data_List_insertBy(v_0, v1_1, (v2_2).V1)})
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
		}
	end_branch_2:
		return __t2
	}
}

func Call_Data_List_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
insertAt:
	for {
		if false {
			continue insertAt
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
		_ = v2_2
		var __t2 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, v2_2}))}, true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))))}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2 != nil {
				// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope192)])])
				__local_var_3_0 := Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := Call_Data_List_insertAt(gopurs_runtime.IntSub(v_0, int64(1)), v1_1, (v2_2).V1)
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))
				_ = __local_var_3_0
				var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				{
					if __local_var_3_0 != nil {
						__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, (__local_var_3_0).V0}))}, true}
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
				__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
		}
	end_branch_2:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_List_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_List_insertBy(), Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_List_go__init(lst_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var lst_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = lst_0_loop
	_ = lst_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a$scope201)]), last: (TypeVar a$scope201)] Empty))])
	__local_var_1_0 := Rebox_Data_List_3094389156_3371309921(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := Call_Data_List_unsnoc(lst_0)
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
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_1_0 != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_1_0).V0.go__init)}, true}
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

func Call_Data_List_index(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
index:
	for {
		if false {
			continue index
		}
		var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if v_0 == nil {
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
				goto end_branch_1
			} else {

			}
		}
		{
			if v_0 != nil {
				var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				{
					if (v1_1) == (int64(0)) {
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
					v_0_loop = (v_0).V1
					v1_1_loop = gopurs_runtime.IntSub(v1_1, int64(1))
					continue index
					__t0 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
				}
			end_branch_0:
				__t1 = __t0
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
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
}

func Call_Data_List_head(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if v_0 == nil {
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
		if v_0 != nil {
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
		__t0 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
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

func Call_Data_List_transpose(v_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
transpose:
	for {
		if false {
			continue transpose
		}
		var v_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_0_loop
		_ = v_0
		var __t15 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
		{
			if v_0 == nil {
				__t15 = (*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(nil)
				goto end_branch_15
			} else {

			}
		}
		{
			if v_0 != nil {
				var __t14 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
				{
					var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_0).V0
					_ = __t_tag_0
					if __t_tag_0 == nil {
						v_0_loop = (v_0).V1
						continue transpose
						__t14 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
							panic("unreachable")
						}()
						goto end_branch_14
					} else {

					}
				}
				{
					var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_0).V0
					_ = __t_tag_1
					if __t_tag_1 != nil {
						var Call_local_Data_List_go__go_1_2_24 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						_ = Call_local_Data_List_go__go_1_2_24
						var go__go_1_2_24 gopurs_runtime.Value
						_ = go__go_1_2_24
						Call_local_Data_List_go__go_1_2_24 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
						go__go_1_2_24:
							for {
								if false {
									continue go__go_1_2_24
								}
								var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
								_ = v_2
								var v1_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v1_3_loop
								_ = v1_3
								var __t7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
								{
									if v1_3 == nil {
										var Call_local_Data_List_go__go_4_3_25 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
										_ = Call_local_Data_List_go__go_4_3_25
										var go__go_4_3_25 gopurs_runtime.Value
										_ = go__go_4_3_25
										Call_local_Data_List_go__go_4_3_25 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
										go__go_4_3_25:
											for {
												if false {
													continue go__go_4_3_25
												}
												var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
												_ = v_5
												var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
												_ = v1_6
												var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
												{
													if v1_6 == nil {
														__t4 = v_5
														goto end_branch_4
													} else {

													}
												}
												{
													if v1_6 != nil {
														v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
														v1_6_loop = (v1_6).V1
														continue go__go_4_3_25
														__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
														goto end_branch_4
													} else {

													}
												}
												{
													__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
												}
											end_branch_4:
												return __t4
											}
										}
										go__go_4_3_25 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_3_25(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
											})
										})
										__t7 = Call_local_Data_List_go__go_4_3_25((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
										goto end_branch_7
									} else {

									}
								}
								{
									if v1_3 != nil {
										// TAST (Let): v2_4_5 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope147)])
										v2_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
											_v := Call_Data_List_head((v1_3).V0)
											if _v.V1 {
												return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
											}
											return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
										}())
										_ = v2_4_5
										var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
										{
											if v2_4_5 == nil {
												v_2_loop = v_2
												v1_3_loop = (v1_3).V1
												continue go__go_1_2_24
												__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
												goto end_branch_6
											} else {

											}
										}
										{
											if v2_4_5 != nil {
												v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_4_5).V0, v_2})
												v1_3_loop = (v1_3).V1
												continue go__go_1_2_24
												__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
												goto end_branch_6
											} else {

											}
										}
										{
											__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
										}
									end_branch_6:
										__t7 = __t6
										goto end_branch_7
									} else {

									}
								}
								{
									__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
								}
							end_branch_7:
								return __t7
							}
						}
						go__go_1_2_24 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_2_24(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val))))}
							})
						})
						var Call_local_Data_List_go__go_1_8_26 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						_ = Call_local_Data_List_go__go_1_8_26
						var go__go_1_8_26 gopurs_runtime.Value
						_ = go__go_1_8_26
						Call_local_Data_List_go__go_1_8_26 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
						go__go_1_8_26:
							for {
								if false {
									continue go__go_1_8_26
								}
								var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
								_ = v_2
								var v1_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v1_3_loop
								_ = v1_3
								var __t13 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
								{
									if v1_3 == nil {
										var Call_local_Data_List_go__go_4_9_27 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
										_ = Call_local_Data_List_go__go_4_9_27
										var go__go_4_9_27 gopurs_runtime.Value
										_ = go__go_4_9_27
										Call_local_Data_List_go__go_4_9_27 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
										go__go_4_9_27:
											for {
												if false {
													continue go__go_4_9_27
												}
												var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
												_ = v_5
												var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
												_ = v1_6
												var __t10 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
												{
													if v1_6 == nil {
														__t10 = v_5
														goto end_branch_10
													} else {

													}
												}
												{
													if v1_6 != nil {
														v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
														v1_6_loop = (v1_6).V1
														continue go__go_4_9_27
														__t10 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
														goto end_branch_10
													} else {

													}
												}
												{
													__t10 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
												}
											end_branch_10:
												return __t10
											}
										}
										go__go_4_9_27 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_9_27(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
											})
										})
										__t13 = Call_local_Data_List_go__go_4_9_27((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
										goto end_branch_13
									} else {

									}
								}
								{
									if v1_3 != nil {
										// TAST (Let): v2_4_11 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope147)])
										v2_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
											_v := Call_Data_List_tail((v1_3).V0)
											if _v.V1 {
												return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
											}
											return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
										}())
										_ = v2_4_11
										var __t12 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
										{
											if v2_4_11 == nil {
												v_2_loop = v_2
												v1_3_loop = (v1_3).V1
												continue go__go_1_8_26
												__t12 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
												goto end_branch_12
											} else {

											}
										}
										{
											if v2_4_11 != nil {
												v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_4_11).V0, v_2})
												v1_3_loop = (v1_3).V1
												continue go__go_1_8_26
												__t12 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
												goto end_branch_12
											} else {

											}
										}
										{
											__t12 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
										}
									end_branch_12:
										__t13 = __t12
										goto end_branch_13
									} else {

									}
								}
								{
									__t13 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
								}
							end_branch_13:
								return __t13
							}
						}
						go__go_1_8_26 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_8_26(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val))))}
							})
						})
						__t14 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v_0).V0).V0, Call_local_Data_List_go__go_1_2_24((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (v_0).V1)}), Call_Data_List_transpose((&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, ((v_0).V0).V1, Rebox_Data_List_849153993_1116310629(Call_local_Data_List_go__go_1_8_26((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (v_0).V1))}))})
						goto end_branch_14
					} else {

					}
				}
				{
					__t14 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
						panic("Failed pattern match")
					}()
				}
			end_branch_14:
				__t15 = __t14
				goto end_branch_15
			} else {

			}
		}
		{
			__t15 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
				panic("Failed pattern match")
			}()
		}
	end_branch_15:
		return __t15
	}
}

func Call_Data_List_groupBy(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
groupBy:
	for {
		if false {
			continue groupBy
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
		_ = v1_1
		var __t1 *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
		{
			if v1_1 == nil {
				__t1 = (*Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]])(nil)
				goto end_branch_1
			} else {

			}
		}
		{
			if v1_1 != nil {
				// TAST (Let): v2_2_0 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a$scope220)]), rest: (ADT ["Data","List","Types","List"] [(TypeVar a$scope220)])] Empty))
				v2_2_0 := Call_Data_List_span(gopurs_runtime.Apply(v_0, (v1_1).V0), (v1_1).V1)
				_ = v2_2_0
				__t1 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v1_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2_0.go__init)}}), Call_Data_List_groupBy(v_0, v2_2_0.rest)})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
				panic("Failed pattern match")
			}()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_List_groupAllBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_List_groupBy(), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(p_0, x_1, y_2).IntVal)
		_ = __t_tag_0
		return gopurs_runtime.Bool((uint32(__t_tag_0) == 902936544))
	})), Call_Data_List_sortBy(p_0))
}

func Call_Data_List_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_List_groupBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(dictOrd_0.V0, gopurs_runtime.Value{}))), Call_Data_List_sort(dictOrd_0))
}

func Call_Data_List_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(dictFoldable_0.V2, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value, v2_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
foldM:
	for {
		if false {
			continue foldM
		}
		var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
		_ = dictMonad_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		var v1_2 gopurs_runtime.Value = v1_2_loop
		_ = v1_2
		var v2_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_3_loop
		_ = v2_3
		var __t1 gopurs_runtime.Value
		{
			if v2_3 == nil {
				__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), v1_2)
				goto end_branch_1
			} else {

			}
		}
		{
			if v2_3 != nil {
				// TAST (Let): __local_var_4_0 shape=Other bindingType=Any
				__local_var_4_0 := (v2_3).V1
				_ = __local_var_4_0
				__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, v1_2, (v2_3).V0), gopurs_runtime.Func(func(b_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_Data_List_foldM(dictMonad_0, v_1, b_prime__5, __local_var_4_0)
				}))
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

func Call_Data_List_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var fn_0 gopurs_runtime.Value = fn_0_loop
	_ = fn_0
	var Call_local_Data_List_go__816245670_1_0_28 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_go__816245670_1_0_28
	var go__816245670_1_0_28 gopurs_runtime.Value
	_ = go__816245670_1_0_28
	var Call_local_Data_List_go__go_1_1_29 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_go__go_1_1_29
	var go__go_1_1_29 gopurs_runtime.Value
	_ = go__go_1_1_29
	Call_local_Data_List_go__816245670_1_0_28 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	go__816245670_1_0_28:
		for {
			if false {
				continue go__816245670_1_0_28
			}
			var v_2 int64 = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t3 *Constructor_Data_Maybe_Just[int64]
			{
				if v1_3 != nil {
					var __t2 *Constructor_Data_Maybe_Just[int64]
					{
						if (gopurs_runtime.Apply(fn_0, (v1_3).V0).IntVal) != (0) {
							__t2 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Int(v_2), true}
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
						v_2_loop = gopurs_runtime.IntAdd(v_2, int64(1))
						v1_3_loop = (v1_3).V1
						continue go__816245670_1_0_28
						__t2 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
					}
				end_branch_2:
					__t3 = __t2
					goto end_branch_3
				} else {

				}
			}
			{
				if v1_3 == nil {
					__t3 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	go__816245670_1_0_28 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_1_0_28(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val))))}
		})
	})
	Call_local_Data_List_go__go_1_1_29 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	go__go_1_1_29:
		for {
			if false {
				continue go__go_1_1_29
			}
			var v_2 int64 = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t5 *Constructor_Data_Maybe_Just[int64]
			{
				if v1_3 != nil {
					var __t4 *Constructor_Data_Maybe_Just[int64]
					{
						if (gopurs_runtime.Apply(fn_0, (v1_3).V0).IntVal) != (0) {
							__t4 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Int(v_2), true}
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
						__t4 = Call_local_Data_List_go__816245670_1_0_28(gopurs_runtime.IntAdd(v_2, int64(1)), (v1_3).V1)
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				if v1_3 == nil {
					__t5 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	go__go_1_1_29 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__go_1_1_29(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val))))}
		})
	})
	return gopurs_runtime.Apply(go__816245670_1_0_28, gopurs_runtime.Int(int64(0)))
}

func Call_Data_List_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var fn_0 gopurs_runtime.Value = fn_0_loop
	_ = fn_0
	var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	var Call_local_Data_List_go__816245670_2_0_30 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_go__816245670_2_0_30
	var go__816245670_2_0_30 gopurs_runtime.Value
	_ = go__816245670_2_0_30
	var Call_local_Data_List_go__go_2_1_31 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_go__go_2_1_31
	var go__go_2_1_31 gopurs_runtime.Value
	_ = go__go_2_1_31
	Call_local_Data_List_go__816245670_2_0_30 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	go__816245670_2_0_30:
		for {
			if false {
				continue go__816245670_2_0_30
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
						if (gopurs_runtime.Apply(fn_0, (v1_4).V0).IntVal) != (0) {
							__t2 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
						continue go__816245670_2_0_30
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
					__t3 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	go__816245670_2_0_30 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_2_0_30(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
		})
	})
	Call_local_Data_List_go__go_2_1_31 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	go__go_2_1_31:
		for {
			if false {
				continue go__go_2_1_31
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
						if (gopurs_runtime.Apply(fn_0, (v1_4).V0).IntVal) != (0) {
							__t4 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
						__t4 = Call_local_Data_List_go__816245670_2_0_30(gopurs_runtime.IntAdd(v_3, int64(1)), (v1_4).V1)
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				if v1_4 == nil {
					__t5 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	go__go_2_1_31 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__go_2_1_31(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
		})
	})
	var Call_local_Data_List_go__go_3_7_32 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_go__go_3_7_32
	var go__go_3_7_32 gopurs_runtime.Value
	_ = go__go_3_7_32
	Call_local_Data_List_go__go_3_7_32 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	go__go_3_7_32:
		for {
			if false {
				continue go__go_3_7_32
			}
			var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
			_ = v_4
			var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
			_ = v1_5
			var __t8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				if v1_5 == nil {
					__t8 = v_4
					goto end_branch_8
				} else {

				}
			}
			{
				if v1_5 != nil {
					v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
					v1_5_loop = (v1_5).V1
					continue go__go_3_7_32
					__t8 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
					goto end_branch_8
				} else {

				}
			}
			{
				__t8 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_8:
			return __t8
		}
	}
	go__go_3_7_32 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_7_32(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
		})
	})
	// TAST (Let): __local_var_3_6 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
	__local_var_3_6 := Call_local_Data_List_go__816245670_2_0_30(int64(0), Call_local_Data_List_go__go_3_7_32((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), xs_1))
	_ = __local_var_3_6
	var __t11 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_3_6 != nil {
			var Call_local_Data_List_go__go_4_9_33 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
			_ = Call_local_Data_List_go__go_4_9_33
			var go__go_4_9_33 gopurs_runtime.Value
			_ = go__go_4_9_33
			Call_local_Data_List_go__go_4_9_33 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
			go__go_4_9_33:
				for {
					if false {
						continue go__go_4_9_33
					}
					var b_5 gopurs_runtime.Value = b_5_loop
					_ = b_5
					var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
					_ = v_6
					var __t10 gopurs_runtime.Value
					{
						if v_6 == nil {
							__t10 = b_5
							goto end_branch_10
						} else {

						}
					}
					{
						if v_6 != nil {
							b_5_loop = gopurs_runtime.Int(gopurs_runtime.IntAdd(b_5.IntVal, int64(1)))
							v_6_loop = (v_6).V1
							continue go__go_4_9_33
							__t10 = func() gopurs_runtime.Value { panic("unreachable") }()
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_10:
					return __t10
				}
			}
			go__go_4_9_33 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Data_List_go__go_4_9_33(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
				})
			})
			__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.IntSub(Call_local_Data_List_go__go_4_9_33(gopurs_runtime.Int(int64(0)), xs_1).IntVal, int64(1)), gopurs_runtime.Int((__local_var_3_6).V0).IntVal)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_11
		} else {

		}
	}
	{
		__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
end_branch_11:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_List_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
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
		// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope259)])
		Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
		_ = Bind1_2_1
		return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t5 gopurs_runtime.Value
			{
				if v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil {
					__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))})
					goto end_branch_5
				} else {

				}
			}
			{
				if v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil {
					// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
					__local_var_5_2 := (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0
					_ = __local_var_5_2
					// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
					__local_var_6_3 := (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1
					_ = __local_var_6_3
					__t5 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(v_3, __local_var_5_2), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_Data_List_filterM(dictMonad_0), v_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_3)}), gopurs_runtime.Func(func(xs_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
							var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							{
								if (b_7.IntVal) != (0) {
									__t4 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_prime__8)})
									goto end_branch_4
								} else {

								}
							}
							{
								__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_prime__8)
							}
						end_branch_4:
							return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)})
						}))
					}))
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_5:
			return __t5
		})
	}
}

func Call_Data_List_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var Call_local_Data_List_go__go_1_0_34 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_go__go_1_0_34
	var go__go_1_0_34 gopurs_runtime.Value
	_ = go__go_1_0_34
	Call_local_Data_List_go__go_1_0_34 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	go__go_1_0_34:
		for {
			if false {
				continue go__go_1_0_34
			}
			var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
			_ = v1_3
			var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				if v1_3 == nil {
					var Call_local_Data_List_go__go_4_1_35 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					_ = Call_local_Data_List_go__go_4_1_35
					var go__go_4_1_35 gopurs_runtime.Value
					_ = go__go_4_1_35
					Call_local_Data_List_go__go_4_1_35 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
					go__go_4_1_35:
						for {
							if false {
								continue go__go_4_1_35
							}
							var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
							_ = v_5
							var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
							_ = v1_6
							var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
							{
								if v1_6 == nil {
									__t2 = v_5
									goto end_branch_2
								} else {

								}
							}
							{
								if v1_6 != nil {
									v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
									v1_6_loop = (v1_6).V1
									continue go__go_4_1_35
									__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
									goto end_branch_2
								} else {

								}
							}
							{
								__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
							}
						end_branch_2:
							return __t2
						}
					}
					go__go_4_1_35 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_35(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
						})
					})
					__t4 = Call_local_Data_List_go__go_4_1_35((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
					goto end_branch_4
				} else {

				}
			}
			{
				if v1_3 != nil {
					var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					{
						if (gopurs_runtime.Apply(p_0, (v1_3).V0).IntVal) != (0) {
							v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, v_2})
							v1_3_loop = (v1_3).V1
							continue go__go_1_0_34
							__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_3
						} else {

						}
					}
					{
						v_2_loop = v_2
						v1_3_loop = (v1_3).V1
						continue go__go_1_0_34
						__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
					}
				end_branch_3:
					__t4 = __t3
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_4:
			return __t4
		}
	}
	go__go_1_0_34 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_34(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
		})
	})
	return gopurs_runtime.Apply(go__go_1_0_34, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_intersectBy(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
	_ = v1_1
	var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
	_ = v2_2
	var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	{
		if v1_1 == nil {
			__t5 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
			goto end_branch_5
		} else {

		}
	}
	{
		if v2_2 == nil {
			__t5 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
			goto end_branch_5
		} else {

		}
	}
	{
		var Call_local_Data_List_go__go_3_0_36 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		_ = Call_local_Data_List_go__go_3_0_36
		var go__go_3_0_36 gopurs_runtime.Value
		_ = go__go_3_0_36
		Call_local_Data_List_go__go_3_0_36 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
		go__go_3_0_36:
			for {
				if false {
					continue go__go_3_0_36
				}
				var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
				_ = v_4
				var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
				_ = v1_5
				var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				{
					if v1_5 == nil {
						var Call_local_Data_List_go__go_6_1_37 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						_ = Call_local_Data_List_go__go_6_1_37
						var go__go_6_1_37 gopurs_runtime.Value
						_ = go__go_6_1_37
						Call_local_Data_List_go__go_6_1_37 = func(v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
						go__go_6_1_37:
							for {
								if false {
									continue go__go_6_1_37
								}
								var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
								_ = v_7
								var v1_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_8_loop
								_ = v1_8
								var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
								{
									if v1_8 == nil {
										__t2 = v_7
										goto end_branch_2
									} else {

									}
								}
								{
									if v1_8 != nil {
										v_7_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_8).V0, v_7})
										v1_8_loop = (v1_8).V1
										continue go__go_6_1_37
										__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
										goto end_branch_2
									} else {

									}
								}
								{
									__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
								}
							end_branch_2:
								return __t2
							}
						}
						go__go_6_1_37 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_6_1_37(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_8_loop_val)))}
							})
						})
						__t4 = Call_local_Data_List_go__go_6_1_37((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_4)
						goto end_branch_4
					} else {

					}
				}
				{
					if v1_5 != nil {
						var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						{
							if (gopurs_runtime.Apply3(Call_Data_Foldable_any(Rebox_Data_List_1022383170_1680800814(Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())))), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3591112874_2663347022(Rebox_Data_List_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))}, gopurs_runtime.Apply(v_0, (v1_5).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0) {
								v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
								v1_5_loop = (v1_5).V1
								continue go__go_3_0_36
								__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
								goto end_branch_3
							} else {

							}
						}
						{
							v_4_loop = v_4
							v1_5_loop = (v1_5).V1
							continue go__go_3_0_36
							__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
						}
					end_branch_3:
						__t4 = __t3
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			}
		}
		go__go_3_0_36 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_0_36(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
			})
		})
		__t5 = Call_local_Data_List_go__go_3_0_36((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v1_1)
	}
end_branch_5:
	return __t5
}

func Call_Data_List_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_List_intersectBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_nubByEq(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
nubByEq:
	for {
		if false {
			continue nubByEq
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
		_ = v1_1
		var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		{
			if v1_1 == nil {
				__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
				goto end_branch_6
			} else {

			}
		}
		{
			if v1_1 != nil {
				// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
				__local_var_2_0 := (v1_1).V0
				_ = __local_var_2_0
				var Call_local_Data_List_go__go_3_1_38 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				_ = Call_local_Data_List_go__go_3_1_38
				var go__go_3_1_38 gopurs_runtime.Value
				_ = go__go_3_1_38
				Call_local_Data_List_go__go_3_1_38 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
				go__go_3_1_38:
					for {
						if false {
							continue go__go_3_1_38
						}
						var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
						_ = v_4
						var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
						_ = v1_5
						var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
						{
							if v1_5 == nil {
								var Call_local_Data_List_go__go_6_2_39 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
								_ = Call_local_Data_List_go__go_6_2_39
								var go__go_6_2_39 gopurs_runtime.Value
								_ = go__go_6_2_39
								Call_local_Data_List_go__go_6_2_39 = func(v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
								go__go_6_2_39:
									for {
										if false {
											continue go__go_6_2_39
										}
										var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
										_ = v_7
										var v1_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_8_loop
										_ = v1_8
										var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
										{
											if v1_8 == nil {
												__t3 = v_7
												goto end_branch_3
											} else {

											}
										}
										{
											if v1_8 != nil {
												v_7_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_8).V0, v_7})
												v1_8_loop = (v1_8).V1
												continue go__go_6_2_39
												__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
												goto end_branch_3
											} else {

											}
										}
										{
											__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
										}
									end_branch_3:
										return __t3
									}
								}
								go__go_6_2_39 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_6_2_39(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_8_loop_val)))}
									})
								})
								__t5 = Call_local_Data_List_go__go_6_2_39((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_4)
								goto end_branch_5
							} else {

							}
						}
						{
							if v1_5 != nil {
								var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
								{
									if ((gopurs_runtime.Apply2(v_0, __local_var_2_0, (v1_5).V0).IntVal) != (0)) != (true) {
										v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
										v1_5_loop = (v1_5).V1
										continue go__go_3_1_38
										__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
										goto end_branch_4
									} else {

									}
								}
								{
									v_4_loop = v_4
									v1_5_loop = (v1_5).V1
									continue go__go_3_1_38
									__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
								}
							end_branch_4:
								__t5 = __t4
								goto end_branch_5
							} else {

							}
						}
						{
							__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
						}
					end_branch_5:
						return __t5
					}
				}
				go__go_3_1_38 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_1_38(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
					})
				})
				__t6 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_0, Call_Data_List_nubByEq(v_0, Call_local_Data_List_go__go_3_1_38((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (v1_1).V1))})
				goto end_branch_6
			} else {

			}
		}
		{
			__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
		}
	end_branch_6:
		return __t6
	}
}

func Call_Data_List_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_List_nubByEq(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1636902754_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Types_eqList(dictEq_0), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](y_2))})
	})})))}
}

func Call_Data_List_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqPattern1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope128)])])
	eqPattern1_1_0 := Rebox_Data_List_3790796878_1636902754(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_eqPattern(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))))
	_ = eqPattern1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_List_4210054658_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1636902754_3790796878(eqPattern1_1_0))}
	}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Types_ordList(dictOrd_0), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](y_3))})
	})})))}
}

func Call_Data_List_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Data_List_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1)
	}))
}

func Call_Data_List_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var Call_local_Data_List_go__816245670_2_0_40 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_go__816245670_2_0_40
	var go__816245670_2_0_40 gopurs_runtime.Value
	_ = go__816245670_2_0_40
	var Call_local_Data_List_go__go_2_1_41 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
	_ = Call_local_Data_List_go__go_2_1_41
	var go__go_2_1_41 gopurs_runtime.Value
	_ = go__go_2_1_41
	Call_local_Data_List_go__816245670_2_0_40 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	go__816245670_2_0_40:
		for {
			if false {
				continue go__816245670_2_0_40
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
						if (gopurs_runtime.Apply2(dictEq_0.V0, (v1_4).V0, x_1).IntVal) != (0) {
							__t2 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
						continue go__816245670_2_0_40
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
					__t3 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	go__816245670_2_0_40 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_2_0_40(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
		})
	})
	Call_local_Data_List_go__go_2_1_41 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	go__go_2_1_41:
		for {
			if false {
				continue go__go_2_1_41
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
						if (gopurs_runtime.Apply2(dictEq_0.V0, (v1_4).V0, x_1).IntVal) != (0) {
							__t4 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
						__t4 = Call_local_Data_List_go__816245670_2_0_40(gopurs_runtime.IntAdd(v_3, int64(1)), (v1_4).V1)
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				if v1_4 == nil {
					__t5 = Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	go__go_2_1_41 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__go_2_1_41(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
		})
	})
	return gopurs_runtime.Apply(go__816245670_2_0_40, gopurs_runtime.Int(int64(0)))
}

func Call_Data_List_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	var Call_local_Data_List_go__go_1_0_42 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	_ = Call_local_Data_List_go__go_1_0_42
	var go__go_1_0_42 gopurs_runtime.Value
	_ = go__go_1_0_42
	Call_local_Data_List_go__go_1_0_42 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	go__go_1_0_42:
		for {
			if false {
				continue go__go_1_0_42
			}
			var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
			_ = v_2
			var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			{
				if (v_2 != nil) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
					v_2_loop = (v_2).V1
					continue go__go_1_0_42
					__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = v_2
			}
		end_branch_1:
			return __t1
		}
	}
	go__go_1_0_42 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_42(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
	})
	return go__go_1_0_42
}

func Call_Data_List_dropEnd(n_0_loop int64, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var n_0 int64 = n_0_loop
	_ = n_0
	var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	var Call_local_Data_List_go__go_2_0_43 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
	_ = Call_local_Data_List_go__go_2_0_43
	var go__go_2_0_43 gopurs_runtime.Value
	_ = go__go_2_0_43
	Call_local_Data_List_go__go_2_0_43 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	go__go_2_0_43:
		for {
			if false {
				continue go__go_2_0_43
			}
			var b_3 gopurs_runtime.Value = b_3_loop
			_ = b_3
			var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
			_ = v_4
			var __t1 gopurs_runtime.Value
			{
				if v_4 == nil {
					__t1 = b_3
					goto end_branch_1
				} else {

				}
			}
			{
				if v_4 != nil {
					b_3_loop = gopurs_runtime.Int(gopurs_runtime.IntAdd(b_3.IntVal, int64(1)))
					v_4_loop = (v_4).V1
					continue go__go_2_0_43
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
	go__go_2_0_43 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_List_go__go_2_0_43(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
		})
	})
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_take(), gopurs_runtime.Int(gopurs_runtime.IntSub(Call_local_Data_List_go__go_2_0_43(gopurs_runtime.Int(int64(0)), xs_1).IntVal, n_0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}))
}

func Call_Data_List_drop(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
drop:
	for {
		if false {
			continue drop
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
		_ = v1_1
		var __t0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		{
			if (v_0) < (int64(1)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			if v1_1 == nil {
				__t0 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
				goto end_branch_0
			} else {

			}
		}
		{
			if v1_1 != nil {
				v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
				v1_1_loop = (v1_1).V1
				continue drop
				__t0 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Data_List_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	var xs_2 gopurs_runtime.Value = xs_2_loop
	_ = xs_2
	return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_take(), gopurs_runtime.Int(gopurs_runtime.IntSub(end_1, start_0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop(start_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2)))})))}
}

func Call_Data_List_takeEnd(n_0_loop int64, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var n_0 int64 = n_0_loop
	_ = n_0
	var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	var Call_local_Data_List_go__go_2_0_44 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
	_ = Call_local_Data_List_go__go_2_0_44
	var go__go_2_0_44 gopurs_runtime.Value
	_ = go__go_2_0_44
	Call_local_Data_List_go__go_2_0_44 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	go__go_2_0_44:
		for {
			if false {
				continue go__go_2_0_44
			}
			var b_3 gopurs_runtime.Value = b_3_loop
			_ = b_3
			var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
			_ = v_4
			var __t1 gopurs_runtime.Value
			{
				if v_4 == nil {
					__t1 = b_3
					goto end_branch_1
				} else {

				}
			}
			{
				if v_4 != nil {
					b_3_loop = gopurs_runtime.Int(gopurs_runtime.IntAdd(b_3.IntVal, int64(1)))
					v_4_loop = (v_4).V1
					continue go__go_2_0_44
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
	go__go_2_0_44 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_List_go__go_2_0_44(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
		})
	})
	return Call_Data_List_drop(gopurs_runtime.IntSub(Call_local_Data_List_go__go_2_0_44(gopurs_runtime.Int(int64(0)), xs_1).IntVal, n_0), xs_1)
}

func Call_Data_List_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
deleteBy:
	for {
		if false {
			continue deleteBy
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
		_ = v2_2
		var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		{
			if v2_2 == nil {
				__t1 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
				goto end_branch_1
			} else {

			}
		}
		{
			if v2_2 != nil {
				var __t0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				{
					if (gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0).IntVal) != (0) {
						__t0 = (v2_2).V1
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, Call_Data_List_deleteBy(v_0, v1_1, (v2_2).V1)})
				}
			end_branch_0:
				__t1 = __t0
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_List_unionBy(eq2_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var eq2_0 gopurs_runtime.Value = eq2_0_loop
	_ = eq2_0
	var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
	_ = xs_1
	var ys_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_2_loop
	_ = ys_2
	var Call_local_Data_List_go__go_3_0_45 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
	_ = Call_local_Data_List_go__go_3_0_45
	var go__go_3_0_45 gopurs_runtime.Value
	_ = go__go_3_0_45
	Call_local_Data_List_go__go_3_0_45 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	go__go_3_0_45:
		for {
			if false {
				continue go__go_3_0_45
			}
			var b_4 gopurs_runtime.Value = b_4_loop
			_ = b_4
			var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
			_ = v_5
			var __t1 gopurs_runtime.Value
			{
				if v_5 == nil {
					__t1 = b_4
					goto end_branch_1
				} else {

				}
			}
			{
				if v_5 != nil {
					b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(eq2_0, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)))}
					v_5_loop = (v_5).V1
					continue go__go_3_0_45
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
	go__go_3_0_45 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_List_go__go_3_0_45(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
		})
	})
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Rebox_Data_List_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_go__go_3_0_45(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(eq2_0, ys_2))}, xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}))
}

func Call_Data_List_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_List_unionBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_deleteAt(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
deleteAt:
	for {
		if false {
			continue deleteAt
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
		_ = v1_1
		var __t3 gopurs_runtime.Value
		{
			if v1_1 != nil {
				var __t2 gopurs_runtime.Value
				{
					if (v_0) == (int64(0)) {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1).V1)}, true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))))}
						goto end_branch_2
					} else {

					}
				}
				{
					// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope310)])])
					__local_var_2_0 := Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := Call_Data_List_deleteAt(gopurs_runtime.IntSub(v_0, int64(1)), (v1_1).V1)
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					_ = __local_var_2_0
					var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					{
						if __local_var_2_0 != nil {
							__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_1).V0, (__local_var_2_0).V0}))}, true}
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
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				}
			end_branch_2:
				__t3 = __t2
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
		}
	end_branch_3:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_List_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_List_deleteBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var Call_local_Data_List_go__go_1_0_46 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
	_ = Call_local_Data_List_go__go_1_0_46
	var go__go_1_0_46 gopurs_runtime.Value
	_ = go__go_1_0_46
	Call_local_Data_List_go__go_1_0_46 = func(b_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
	go__go_1_0_46:
		for {
			if false {
				continue go__go_1_0_46
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
					b_2_loop = gopurs_runtime.Apply3(Get_Data_List_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}, (v_3).V0, b_2)
					v_3_loop = (v_3).V1
					continue go__go_1_0_46
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
	go__go_1_0_46 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_List_go__go_1_0_46(b_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))
		})
	})
	return go__go_1_0_46
}

func Call_Data_List_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = a_1_loop
	_ = a_1
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_2748095225_2183599445(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_1)}, b_0))
}

func Call_Data_List_concat(v_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var v_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_2748095225_2183599445(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(v_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	})))
}

func Call_Data_List_alterAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
alterAt:
	for {
		if false {
			continue alterAt
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
		_ = v2_2
		var __t5 gopurs_runtime.Value
		{
			if v2_2 != nil {
				var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				{
					if (v_0) == (int64(0)) {
						// TAST (Let): v3_3_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope330)])
						v3_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (v2_2).V0))
						_ = v3_3_2
						var __t3 gopurs_runtime.Value
						{
							if v3_3_2 == nil {
								__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v2_2).V1)}
								goto end_branch_3
							} else {

							}
						}
						{
							if v3_3_2 != nil {
								__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v3_3_2).V0, (v2_2).V1}))}
								goto end_branch_3
							} else {

							}
						}
						{
							__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_3:
						__t4 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, __t3})
						goto end_branch_4
					} else {

					}
				}
				{
					// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope330)])])
					__local_var_3_0 := Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := Call_Data_List_alterAt(gopurs_runtime.IntSub(v_0, int64(1)), v1_1, (v2_2).V1)
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					_ = __local_var_3_0
					var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					{
						if __local_var_3_0 != nil {
							__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, (__local_var_3_0).V0}))}, true}
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
					__t4 = __t1
				}
			end_branch_4:
				__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
				goto end_branch_5
			} else {

			}
		}
		{
			__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
		}
	end_branch_5:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_List_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply2(Get_Data_List_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), f_1))
}

func Rebox_Data_List_1022383170_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_1116310629_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = Rebox_Data_List_1116310629_849153993(in.V1)
	return out
}

func Rebox_Data_List_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_List_138441832_22134120(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
	out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_List_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
	out.V0 = uint32(in.V0.IntVal)
	out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_List_1501865320_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_1636902754_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_1978018568_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_List_22134120_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_List_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3132786365_138441832(in.V0))}
	out.V1 = Rebox_Data_List_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_List_2663347022_3591112874(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool] {
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

func Rebox_Data_List_2681346401_3094389156(in *Constructor_Data_Maybe_Just[struct {
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)})
	}()
	return out
}

func Rebox_Data_List_2748095225_2183599445(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_3043886126_3037784642(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_List_3094389156_1978018568(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_3094389156_2681346401(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		head gopurs_runtime.Value
		tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]{}
	out.V0 = func() struct {
		head gopurs_runtime.Value
		tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			head gopurs_runtime.Value
			tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}
		clone.head = gopurs_runtime.RecordGet(orig, "head")
		clone.tail = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "tail"))
		return clone
	}()
	return out
}

func Rebox_Data_List_3094389156_3270864776(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	last    gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		last    gopurs_runtime.Value
		revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
	}]{}
	out.V0 = func() struct {
		last    gopurs_runtime.Value
		revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			last    gopurs_runtime.Value
			revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
			/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
		}{}
		clone.last = gopurs_runtime.RecordGet(orig, "last")
		clone.revInit = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "revInit"))
		return clone
	}()
	return out
}

func Rebox_Data_List_3094389156_3371309921(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
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

func Rebox_Data_List_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_List_3270864776_3094389156(in *Constructor_Data_Maybe_Just[struct {
	last    gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	/*gopurs:record-v1:2:006c006100730074,0072006500760049006e00690074*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("last", "revInit", orig.last, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.revInit)})
	}()
	return out
}

func Rebox_Data_List_3351995458_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_3591112874_2663347022(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] {
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

func Rebox_Data_List_3704040722_849153993(in *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = Rebox_Data_List_3704040722_849153993(in.V1)
	return out
}

func Rebox_Data_List_3790796878_1636902754(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_4130553207_1542299734(in *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_List_4210054658_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_List_587109416_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3123684004_1293498952(in.V0))}
	out.V1 = Rebox_Data_List_587109416_849153993(in.V1)
	return out
}

func Rebox_Data_List_849153993_1116310629(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
	out.V1 = Rebox_Data_List_849153993_1116310629(in.V1)
	return out
}

func Rebox_Data_List_849153993_1220287592(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0)
	out.V1 = Rebox_Data_List_849153993_1220287592(in.V1)
	return out
}

func Rebox_Data_List_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
	out.V0 = Rebox_Data_List_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = Rebox_Data_List_849153993_2442833393(in.V1)
	return out
}

func Rebox_Data_List_849153993_3704040722(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Data_List_849153993_3704040722(in.V1)
	return out
}
