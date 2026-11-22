package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Foreign_Object_compose gopurs_runtime.Value
var once_Foreign_Object_compose sync.Once

func Get_Foreign_Object_compose() gopurs_runtime.Value {
	once_Foreign_Object_compose.Do(func() {
		cache_Foreign_Object_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_Object_compose
}

var cache_Foreign_Object_compose1 gopurs_runtime.Value
var once_Foreign_Object_compose1 sync.Once

func Get_Foreign_Object_compose1() gopurs_runtime.Value {
	once_Foreign_Object_compose1.Do(func() {
		cache_Foreign_Object_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_Object_compose1
}

var cache_Foreign_Object_compose2 gopurs_runtime.Value
var once_Foreign_Object_compose2 sync.Once

func Get_Foreign_Object_compose2() gopurs_runtime.Value {
	once_Foreign_Object_compose2.Do(func() {
		cache_Foreign_Object_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_Object_compose2
}

var cache_Foreign_Object_sortWith gopurs_runtime.Value
var once_Foreign_Object_sortWith sync.Once

func Get_Foreign_Object_sortWith() gopurs_runtime.Value {
	once_Foreign_Object_sortWith.Do(func() {
		cache_Foreign_Object_sortWith = gopurs_runtime.Apply(Get_Data_Array_sortWith(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_2406510097_4177771502(Rebox_Foreign_Object_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString()))))})
	})
	return cache_Foreign_Object_sortWith
}

var cache_Foreign_Object_bindFlipped gopurs_runtime.Value
var once_Foreign_Object_bindFlipped sync.Once

func Get_Foreign_Object_bindFlipped() gopurs_runtime.Value {
	once_Foreign_Object_bindFlipped.Do(func() {
		cache_Foreign_Object_bindFlipped = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_bindFlipped(b_0_box, a_1_box)
		})
	})
	return cache_Foreign_Object_bindFlipped
}

var cache_Foreign_Object_go__append gopurs_runtime.Value
var once_Foreign_Object_go__append sync.Once

func Get_Foreign_Object_go__append() gopurs_runtime.Value {
	once_Foreign_Object_go__append.Do(func() {
		cache_Foreign_Object_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Foreign_Object_go__append
}

var cache_Foreign_Object_showTuple gopurs_runtime.Value
var once_Foreign_Object_showTuple sync.Once

func Get_Foreign_Object_showTuple() gopurs_runtime.Value {
	once_Foreign_Object_showTuple.Do(func() {
		cache_Foreign_Object_showTuple = gopurs_runtime.Apply(Get_Data_Tuple_showTuple(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_1514099793_1386611502(Rebox_Foreign_Object_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Foreign_Object_showTuple
}

var cache_Foreign_Object_bind gopurs_runtime.Value
var once_Foreign_Object_bind sync.Once

func Get_Foreign_Object_bind() gopurs_runtime.Value {
	once_Foreign_Object_bind.Do(func() {
		cache_Foreign_Object_bind = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_bind
}

var cache_Foreign_Object_bind1 gopurs_runtime.Value
var once_Foreign_Object_bind1 sync.Once

func Get_Foreign_Object_bind1() gopurs_runtime.Value {
	once_Foreign_Object_bind1.Do(func() {
		cache_Foreign_Object_bind1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_bind1
}

var cache_Foreign_Object_pure gopurs_runtime.Value
var once_Foreign_Object_pure sync.Once

func Get_Foreign_Object_pure() gopurs_runtime.Value {
	once_Foreign_Object_pure.Do(func() {
		cache_Foreign_Object_pure = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Foreign_Object_pure
}

var cache_Foreign_Object_bind2 gopurs_runtime.Value
var once_Foreign_Object_bind2 sync.Once

func Get_Foreign_Object_bind2() gopurs_runtime.Value {
	once_Foreign_Object_bind2.Do(func() {
		cache_Foreign_Object_bind2 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_bind2
}

var cache_Foreign_Object_discard gopurs_runtime.Value
var once_Foreign_Object_discard sync.Once

func Get_Foreign_Object_discard() gopurs_runtime.Value {
	once_Foreign_Object_discard.Do(func() {
		cache_Foreign_Object_discard = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_discard
}

var cache_Foreign_Object_forWithIndex_ gopurs_runtime.Value
var once_Foreign_Object_forWithIndex_ sync.Once

func Get_Foreign_Object_forWithIndex_() gopurs_runtime.Value {
	once_Foreign_Object_forWithIndex_.Do(func() {
		cache_Foreign_Object_forWithIndex_ = Call_Data_FoldableWithIndex_forWithIndex_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST()))
	})
	return cache_Foreign_Object_forWithIndex_
}

var cache_Foreign_Object_pure1 gopurs_runtime.Value
var once_Foreign_Object_pure1 sync.Once

func Get_Foreign_Object_pure1() gopurs_runtime.Value {
	once_Foreign_Object_pure1.Do(func() {
		cache_Foreign_Object_pure1 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Foreign_Object_pure1
}

var cache_Foreign_Object_bind3 gopurs_runtime.Value
var once_Foreign_Object_bind3 sync.Once

func Get_Foreign_Object_bind3() gopurs_runtime.Value {
	once_Foreign_Object_bind3.Do(func() {
		cache_Foreign_Object_bind3 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_bind3
}

var cache_Foreign_Object_discard1 gopurs_runtime.Value
var once_Foreign_Object_discard1 sync.Once

func Get_Foreign_Object_discard1() gopurs_runtime.Value {
	once_Foreign_Object_discard1.Do(func() {
		cache_Foreign_Object_discard1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_discard1
}

var cache_Foreign_Object_for_ gopurs_runtime.Value
var once_Foreign_Object_for_ sync.Once

func Get_Foreign_Object_for_() gopurs_runtime.Value {
	once_Foreign_Object_for_.Do(func() {
		cache_Foreign_Object_for_ = Call_Data_Foldable_for_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST()))
	})
	return cache_Foreign_Object_for_
}

var cache_Foreign_Object_bind4 gopurs_runtime.Value
var once_Foreign_Object_bind4 sync.Once

func Get_Foreign_Object_bind4() gopurs_runtime.Value {
	once_Foreign_Object_bind4.Do(func() {
		cache_Foreign_Object_bind4 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_bind4
}

var cache_Foreign_Object_pure2 gopurs_runtime.Value
var once_Foreign_Object_pure2 sync.Once

func Get_Foreign_Object_pure2() gopurs_runtime.Value {
	once_Foreign_Object_pure2.Do(func() {
		cache_Foreign_Object_pure2 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Foreign_Object_pure2
}

var cache_Foreign_Object_bind5 gopurs_runtime.Value
var once_Foreign_Object_bind5 sync.Once

func Get_Foreign_Object_bind5() gopurs_runtime.Value {
	once_Foreign_Object_bind5.Do(func() {
		cache_Foreign_Object_bind5 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_bind5
}

var cache_Foreign_Object_discard2 gopurs_runtime.Value
var once_Foreign_Object_discard2 sync.Once

func Get_Foreign_Object_discard2() gopurs_runtime.Value {
	once_Foreign_Object_discard2.Do(func() {
		cache_Foreign_Object_discard2 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_discard2
}

var cache_Foreign_Object_void gopurs_runtime.Value
var once_Foreign_Object_void sync.Once

func Get_Foreign_Object_void() gopurs_runtime.Value {
	once_Foreign_Object_void.Do(func() {
		cache_Foreign_Object_void = Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST()))
	})
	return cache_Foreign_Object_void
}

var cache_Foreign_Object_pure3 gopurs_runtime.Value
var once_Foreign_Object_pure3 sync.Once

func Get_Foreign_Object_pure3() gopurs_runtime.Value {
	once_Foreign_Object_pure3.Do(func() {
		cache_Foreign_Object_pure3 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Foreign_Object_pure3
}

var cache_Foreign_Object_foldr gopurs_runtime.Value
var once_Foreign_Object_foldr sync.Once

func Get_Foreign_Object_foldr() gopurs_runtime.Value {
	once_Foreign_Object_foldr.Do(func() {
		cache_Foreign_Object_foldr = Get_Data_Foldable_foldrArray()
	})
	return cache_Foreign_Object_foldr
}

var cache_Foreign_Object_foldr1 gopurs_runtime.Value
var once_Foreign_Object_foldr1 sync.Once

func Get_Foreign_Object_foldr1() gopurs_runtime.Value {
	once_Foreign_Object_foldr1.Do(func() {
		cache_Foreign_Object_foldr1 = Get_Data_Foldable_foldrArray()
	})
	return cache_Foreign_Object_foldr1
}

var cache_Foreign_Object_compose3 gopurs_runtime.Value
var once_Foreign_Object_compose3 sync.Once

func Get_Foreign_Object_compose3() gopurs_runtime.Value {
	once_Foreign_Object_compose3.Do(func() {
		cache_Foreign_Object_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_Object_compose3
}

var cache_Foreign_Object_identity gopurs_runtime.Value
var once_Foreign_Object_identity sync.Once

func Get_Foreign_Object_identity() gopurs_runtime.Value {
	once_Foreign_Object_identity.Do(func() {
		cache_Foreign_Object_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_identity(x_0_box)
		})
	})
	return cache_Foreign_Object_identity
}

var cache_Foreign_Object_pure4 gopurs_runtime.Value
var once_Foreign_Object_pure4 sync.Once

func Get_Foreign_Object_pure4() gopurs_runtime.Value {
	once_Foreign_Object_pure4.Do(func() {
		cache_Foreign_Object_pure4 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Foreign_Object_pure4
}

var cache_Foreign_Object_bind6 gopurs_runtime.Value
var once_Foreign_Object_bind6 sync.Once

func Get_Foreign_Object_bind6() gopurs_runtime.Value {
	once_Foreign_Object_bind6.Do(func() {
		cache_Foreign_Object_bind6 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Foreign_Object_bind6
}

var cache_Foreign_Object_compose4 gopurs_runtime.Value
var once_Foreign_Object_compose4 sync.Once

func Get_Foreign_Object_compose4() gopurs_runtime.Value {
	once_Foreign_Object_compose4.Do(func() {
		cache_Foreign_Object_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Foreign_Object_compose4
}

var cache_Foreign_Object_conj gopurs_runtime.Value
var once_Foreign_Object_conj sync.Once

func Get_Foreign_Object_conj() gopurs_runtime.Value {
	once_Foreign_Object_conj.Do(func() {
		cache_Foreign_Object_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Foreign_Object_conj
}

var cache_Foreign_Object_ordTuple gopurs_runtime.Value
var once_Foreign_Object_ordTuple sync.Once

func Get_Foreign_Object_ordTuple() gopurs_runtime.Value {
	once_Foreign_Object_ordTuple.Do(func() {
		cache_Foreign_Object_ordTuple = Call_Data_Tuple_ordTuple(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_2406510097_4177771502(Rebox_Foreign_Object_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString()))))})
	})
	return cache_Foreign_Object_ordTuple
}

var cache_Foreign_Object_mapFlipped gopurs_runtime.Value
var once_Foreign_Object_mapFlipped sync.Once

func Get_Foreign_Object_mapFlipped() gopurs_runtime.Value {
	once_Foreign_Object_mapFlipped.Do(func() {
		cache_Foreign_Object_mapFlipped = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Foreign_Object_mapFlipped(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](fa_0_box), f_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Foreign_Object_mapFlipped
}

var cache_Foreign_Object_values gopurs_runtime.Value
var once_Foreign_Object_values sync.Once

func Get_Foreign_Object_values() gopurs_runtime.Value {
	once_Foreign_Object_values.Do(func() {
		cache_Foreign_Object_values = gopurs_runtime.Apply(Get_Foreign_Object_toArrayWithKey(), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return v1_1
		}))
	})
	return cache_Foreign_Object_values
}

var cache_Foreign_Object_toUnfoldable gopurs_runtime.Value
var once_Foreign_Object_toUnfoldable sync.Once

func Get_Foreign_Object_toUnfoldable() gopurs_runtime.Value {
	once_Foreign_Object_toUnfoldable.Do(func() {
		cache_Foreign_Object_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Foreign_Object_toUnfoldable
}

var cache_Foreign_Object_toAscUnfoldable gopurs_runtime.Value
var once_Foreign_Object_toAscUnfoldable sync.Once

func Get_Foreign_Object_toAscUnfoldable() gopurs_runtime.Value {
	once_Foreign_Object_toAscUnfoldable.Do(func() {
		cache_Foreign_Object_toAscUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_toAscUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Foreign_Object_toAscUnfoldable
}

var cache_Foreign_Object_toAscArray gopurs_runtime.Value
var once_Foreign_Object_toAscArray sync.Once

func Get_Foreign_Object_toAscArray() gopurs_runtime.Value {
	once_Foreign_Object_toAscArray.Do(func() {
		cache_Foreign_Object_toAscArray = Call_Foreign_Object_toAscUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()))
	})
	return cache_Foreign_Object_toAscArray
}

var cache_Foreign_Object_toArray gopurs_runtime.Value
var once_Foreign_Object_toArray sync.Once

func Get_Foreign_Object_toArray() gopurs_runtime.Value {
	once_Foreign_Object_toArray.Do(func() {
		cache_Foreign_Object_toArray = gopurs_runtime.Apply(Get_Foreign_Object_toArrayWithKey(), Get_Data_Tuple_Tuple())
	})
	return cache_Foreign_Object_toArray
}

var cache_Foreign_Object_thawST gopurs_runtime.Value
var once_Foreign_Object_thawST sync.Once

func Get_Foreign_Object_thawST() gopurs_runtime.Value {
	once_Foreign_Object_thawST.Do(func() {
		cache_Foreign_Object_thawST = Get_Foreign_Object__copyST()
	})
	return cache_Foreign_Object_thawST
}

var cache_Foreign_Object_singleton gopurs_runtime.Value
var once_Foreign_Object_singleton sync.Once

func Get_Foreign_Object_singleton() gopurs_runtime.Value {
	once_Foreign_Object_singleton.Do(func() {
		cache_Foreign_Object_singleton = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_singleton(k_0_box.StrVal(), v_1_box)
		})
	})
	return cache_Foreign_Object_singleton
}

var cache_Foreign_Object_singleton__2505508685 gopurs_runtime.Value
var once_Foreign_Object_singleton__2505508685 sync.Once

func Get_Foreign_Object_singleton__2505508685() gopurs_runtime.Value {
	once_Foreign_Object_singleton__2505508685.Do(func() {
		cache_Foreign_Object_singleton__2505508685 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_singleton__2505508685(k_0_box.StrVal(), v_1_box)
		})
	})
	return cache_Foreign_Object_singleton__2505508685
}

var cache_Foreign_Object_showObject gopurs_runtime.Value
var once_Foreign_Object_showObject sync.Once

func Get_Foreign_Object_showObject() gopurs_runtime.Value {
	once_Foreign_Object_showObject.Do(func() {
		cache_Foreign_Object_showObject = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_showObject(dictShow_0_box)
		})
	})
	return cache_Foreign_Object_showObject
}

var cache_Foreign_Object_mutate gopurs_runtime.Value
var once_Foreign_Object_mutate sync.Once

func Get_Foreign_Object_mutate() gopurs_runtime.Value {
	once_Foreign_Object_mutate.Do(func() {
		cache_Foreign_Object_mutate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_mutate(f_0_box, m_1_box)
		})
	})
	return cache_Foreign_Object_mutate
}

var cache_Foreign_Object_member gopurs_runtime.Value
var once_Foreign_Object_member sync.Once

func Get_Foreign_Object_member() gopurs_runtime.Value {
	once_Foreign_Object_member.Do(func() {
		cache_Foreign_Object_member = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Foreign_Object_member(__local_var_0_box.StrVal(), __local_var_1_box))
		})
	})
	return cache_Foreign_Object_member
}

var cache_Foreign_Object_mapWithKey gopurs_runtime.Value
var once_Foreign_Object_mapWithKey sync.Once

func Get_Foreign_Object_mapWithKey() gopurs_runtime.Value {
	once_Foreign_Object_mapWithKey.Do(func() {
		cache_Foreign_Object_mapWithKey = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_mapWithKey(f_0_box, m_1_box)
		})
	})
	return cache_Foreign_Object_mapWithKey
}

var cache_Foreign_Object_lookup gopurs_runtime.Value
var once_Foreign_Object_lookup sync.Once

func Get_Foreign_Object_lookup() gopurs_runtime.Value {
	once_Foreign_Object_lookup.Do(func() {
		cache_Foreign_Object_lookup = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Foreign_Object_lookup(__local_var_0_box.StrVal(), __local_var_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Foreign_Object_lookup
}

var cache_Foreign_Object_lookup__1777685996 gopurs_runtime.Value
var once_Foreign_Object_lookup__1777685996 sync.Once

func Get_Foreign_Object_lookup__1777685996() gopurs_runtime.Value {
	once_Foreign_Object_lookup__1777685996.Do(func() {
		cache_Foreign_Object_lookup__1777685996 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Foreign_Object_lookup__1777685996(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Foreign_Object_lookup__1777685996
}

var cache_Foreign_Object_lookup__2830175038 gopurs_runtime.Value
var once_Foreign_Object_lookup__2830175038 sync.Once

func Get_Foreign_Object_lookup__2830175038() gopurs_runtime.Value {
	once_Foreign_Object_lookup__2830175038.Do(func() {
		cache_Foreign_Object_lookup__2830175038 = gopurs_runtime.Func2(func(__eta_norm_0_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Foreign_Object_lookup__2830175038(__eta_norm_0_0_box.StrVal(), __local_var_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Foreign_Object_lookup__2830175038
}

var cache_Foreign_Object_isSubmap gopurs_runtime.Value
var once_Foreign_Object_isSubmap sync.Once

func Get_Foreign_Object_isSubmap() gopurs_runtime.Value {
	once_Foreign_Object_isSubmap.Do(func() {
		cache_Foreign_Object_isSubmap = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, m1_1_box gopurs_runtime.Value, m2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Foreign_Object_isSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), m1_1_box, m2_2_box))
		})
	})
	return cache_Foreign_Object_isSubmap
}

var cache_Foreign_Object_isEmpty gopurs_runtime.Value
var once_Foreign_Object_isEmpty sync.Once

func Get_Foreign_Object_isEmpty() gopurs_runtime.Value {
	once_Foreign_Object_isEmpty.Do(func() {
		cache_Foreign_Object_isEmpty = gopurs_runtime.Apply(Get_Foreign_Object_all(), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(false)
		}))
	})
	return cache_Foreign_Object_isEmpty
}

var cache_Foreign_Object_insert gopurs_runtime.Value
var once_Foreign_Object_insert sync.Once

func Get_Foreign_Object_insert() gopurs_runtime.Value {
	once_Foreign_Object_insert.Do(func() {
		cache_Foreign_Object_insert = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_insert(k_0_box.StrVal(), v_1_box)
		})
	})
	return cache_Foreign_Object_insert
}

var cache_Foreign_Object_insert__3487692426 gopurs_runtime.Value
var once_Foreign_Object_insert__3487692426 sync.Once

func Get_Foreign_Object_insert__3487692426() gopurs_runtime.Value {
	once_Foreign_Object_insert__3487692426.Do(func() {
		cache_Foreign_Object_insert__3487692426 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_insert__3487692426(k_0_box.StrVal(), v_1_box)
		})
	})
	return cache_Foreign_Object_insert__3487692426
}

var cache_Foreign_Object_insert__1072177850 gopurs_runtime.Value
var once_Foreign_Object_insert__1072177850 sync.Once

func Get_Foreign_Object_insert__1072177850() gopurs_runtime.Value {
	once_Foreign_Object_insert__1072177850.Do(func() {
		cache_Foreign_Object_insert__1072177850 = gopurs_runtime.Func3(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_insert__1072177850(k_0_box.StrVal(), v_1_box, __eta_norm_0_2_box)
		})
	})
	return cache_Foreign_Object_insert__1072177850
}

var cache_Foreign_Object_functorObject gopurs_runtime.Value
var once_Foreign_Object_functorObject sync.Once

func Get_Foreign_Object_functorObject() gopurs_runtime.Value {
	once_Foreign_Object_functorObject.Do(func() {
		cache_Foreign_Object_functorObject = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.UncurriedApp2(Get_Foreign_Object__fmapObject(), m_1, f_0)
		})}))}
	})
	return cache_Foreign_Object_functorObject
}

var cache_Foreign_Object_functorWithIndexObject gopurs_runtime.Value
var once_Foreign_Object_functorWithIndexObject sync.Once

func Get_Foreign_Object_functorWithIndexObject() gopurs_runtime.Value {
	once_Foreign_Object_functorWithIndexObject.Do(func() {
		cache_Foreign_Object_functorWithIndexObject = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_2872813271_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[string, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Foreign_Object_functorObject()))}
		}), Get_Foreign_Object_mapWithKey()})))}
	})
	return cache_Foreign_Object_functorWithIndexObject
}

var cache_Foreign_Object_fromHomogeneous gopurs_runtime.Value
var once_Foreign_Object_fromHomogeneous sync.Once

func Get_Foreign_Object_fromHomogeneous() gopurs_runtime.Value {
	once_Foreign_Object_fromHomogeneous.Do(func() {
		cache_Foreign_Object_fromHomogeneous = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_fromHomogeneous(_dollar___unused_0_box)
		})
	})
	return cache_Foreign_Object_fromHomogeneous
}

var cache_Foreign_Object_fromFoldableWithIndex gopurs_runtime.Value
var once_Foreign_Object_fromFoldableWithIndex sync.Once

func Get_Foreign_Object_fromFoldableWithIndex() gopurs_runtime.Value {
	once_Foreign_Object_fromFoldableWithIndex.Do(func() {
		cache_Foreign_Object_fromFoldableWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_fromFoldableWithIndex(Rebox_Foreign_Object_3725484264_1767849495(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box)))
		})
	})
	return cache_Foreign_Object_fromFoldableWithIndex
}

var cache_Foreign_Object_fromFoldableWith gopurs_runtime.Value
var once_Foreign_Object_fromFoldableWith sync.Once

func Get_Foreign_Object_fromFoldableWith() gopurs_runtime.Value {
	once_Foreign_Object_fromFoldableWith.Do(func() {
		cache_Foreign_Object_fromFoldableWith = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_fromFoldableWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Foreign_Object_fromFoldableWith
}

var cache_Foreign_Object_fromFoldable gopurs_runtime.Value
var once_Foreign_Object_fromFoldable sync.Once

func Get_Foreign_Object_fromFoldable() gopurs_runtime.Value {
	once_Foreign_Object_fromFoldable.Do(func() {
		cache_Foreign_Object_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Foreign_Object_fromFoldable
}

var cache_Foreign_Object_freezeST gopurs_runtime.Value
var once_Foreign_Object_freezeST sync.Once

func Get_Foreign_Object_freezeST() gopurs_runtime.Value {
	once_Foreign_Object_freezeST.Do(func() {
		cache_Foreign_Object_freezeST = Get_Foreign_Object__copyST()
	})
	return cache_Foreign_Object_freezeST
}

var cache_Foreign_Object_foldMaybe gopurs_runtime.Value
var once_Foreign_Object_foldMaybe sync.Once

func Get_Foreign_Object_foldMaybe() gopurs_runtime.Value {
	once_Foreign_Object_foldMaybe.Do(func() {
		cache_Foreign_Object_foldMaybe = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_foldMaybe(f_0_box, z_1_box, m_2_box)
		})
	})
	return cache_Foreign_Object_foldMaybe
}

var cache_Foreign_Object_foldM gopurs_runtime.Value
var once_Foreign_Object_foldM sync.Once

func Get_Foreign_Object_foldM() gopurs_runtime.Value {
	once_Foreign_Object_foldM.Do(func() {
		cache_Foreign_Object_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Foreign_Object_foldM
}

var cache_Foreign_Object_foldM1 gopurs_runtime.Value
var once_Foreign_Object_foldM1 sync.Once

func Get_Foreign_Object_foldM1() gopurs_runtime.Value {
	once_Foreign_Object_foldM1.Do(func() {
		cache_Foreign_Object_foldM1 = Call_Foreign_Object_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST()))
	})
	return cache_Foreign_Object_foldM1
}

var cache_Foreign_Object_foldM2 gopurs_runtime.Value
var once_Foreign_Object_foldM2 sync.Once

func Get_Foreign_Object_foldM2() gopurs_runtime.Value {
	once_Foreign_Object_foldM2.Do(func() {
		cache_Foreign_Object_foldM2 = Call_Foreign_Object_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST()))
	})
	return cache_Foreign_Object_foldM2
}

var cache_Foreign_Object_foldM3 gopurs_runtime.Value
var once_Foreign_Object_foldM3 sync.Once

func Get_Foreign_Object_foldM3() gopurs_runtime.Value {
	once_Foreign_Object_foldM3.Do(func() {
		cache_Foreign_Object_foldM3 = Call_Foreign_Object_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST()))
	})
	return cache_Foreign_Object_foldM3
}

var cache_Foreign_Object_union gopurs_runtime.Value
var once_Foreign_Object_union sync.Once

func Get_Foreign_Object_union() gopurs_runtime.Value {
	once_Foreign_Object_union.Do(func() {
		cache_Foreign_Object_union = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_union(m_0_box)
		})
	})
	return cache_Foreign_Object_union
}

var cache_Foreign_Object_unions gopurs_runtime.Value
var once_Foreign_Object_unions sync.Once

func Get_Foreign_Object_unions() gopurs_runtime.Value {
	once_Foreign_Object_unions.Do(func() {
		cache_Foreign_Object_unions = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_unions(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Foreign_Object_unions
}

var cache_Foreign_Object_unionWith gopurs_runtime.Value
var once_Foreign_Object_unionWith sync.Once

func Get_Foreign_Object_unionWith() gopurs_runtime.Value {
	once_Foreign_Object_unionWith.Do(func() {
		cache_Foreign_Object_unionWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, m1_1_box gopurs_runtime.Value, m2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_unionWith(f_0_box, m1_1_box, m2_2_box)
		})
	})
	return cache_Foreign_Object_unionWith
}

var cache_Foreign_Object_semigroupObject gopurs_runtime.Value
var once_Foreign_Object_semigroupObject sync.Once

func Get_Foreign_Object_semigroupObject() gopurs_runtime.Value {
	once_Foreign_Object_semigroupObject.Do(func() {
		cache_Foreign_Object_semigroupObject = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_semigroupObject(dictSemigroup_0_box)
		})
	})
	return cache_Foreign_Object_semigroupObject
}

var cache_Foreign_Object_monoidObject gopurs_runtime.Value
var once_Foreign_Object_monoidObject sync.Once

func Get_Foreign_Object_monoidObject() gopurs_runtime.Value {
	once_Foreign_Object_monoidObject.Do(func() {
		cache_Foreign_Object_monoidObject = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_monoidObject(dictSemigroup_0_box)
		})
	})
	return cache_Foreign_Object_monoidObject
}

var cache_Foreign_Object_fold gopurs_runtime.Value
var once_Foreign_Object_fold sync.Once

func Get_Foreign_Object_fold() gopurs_runtime.Value {
	once_Foreign_Object_fold.Do(func() {
		cache_Foreign_Object_fold = gopurs_runtime.Apply(Get_Foreign_Object__foldM(), Get_Data_Function_applyFlipped())
	})
	return cache_Foreign_Object_fold
}

var cache_Foreign_Object_foldMap gopurs_runtime.Value
var once_Foreign_Object_foldMap sync.Once

func Get_Foreign_Object_foldMap() gopurs_runtime.Value {
	once_Foreign_Object_foldMap.Do(func() {
		cache_Foreign_Object_foldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Foreign_Object_foldMap
}

var cache_Foreign_Object_foldableObject gopurs_runtime.Value
var once_Foreign_Object_foldableObject sync.Once

func Get_Foreign_Object_foldableObject() gopurs_runtime.Value {
	once_Foreign_Object_foldableObject.Do(func() {
		cache_Foreign_Object_foldableObject = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): foldMap1_1_0 shape=App(Var) bindingType=(Func [(Func [String, (TypeVar a$scope258)] (TypeVar m$scope259)), (ADT ["Foreign","Object","Object"] [(TypeVar a$scope258)])] (TypeVar m$scope259))
			foldMap1_1_0 := Call_Foreign_Object_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
			_ = foldMap1_1_0
			return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(foldMap1_1_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return f_2
				}))
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Foreign_Object_fold(), gopurs_runtime.Func2(func(z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f_0, z_1)
			}))
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, z_1, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Foreign_Object_values(), m_2).UnsafePtr))))
		})}))}
	})
	return cache_Foreign_Object_foldableObject
}

var cache_Foreign_Object_foldableWithIndexObject gopurs_runtime.Value
var once_Foreign_Object_foldableWithIndexObject sync.Once

func Get_Foreign_Object_foldableWithIndexObject() gopurs_runtime.Value {
	once_Foreign_Object_foldableWithIndexObject.Do(func() {
		cache_Foreign_Object_foldableWithIndexObject = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_1767849495_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[string, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Foreign_Object_foldableObject()))}
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Foreign_Object_fold(), gopurs_runtime.Func2(func(b_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(f_0, gopurs_runtime.NormalizeStringValue(a_2), b_1)
			}))
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(f_0, gopurs_runtime.NormalizeStringValue((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
			}), z_1, func() gopurs_runtime.Value {
				arr := func() []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Foreign_Object_toArrayWithKey(), Get_Data_Tuple_Tuple(), m_2).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Foreign_Object_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_3103076375_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}())
		})})))}
	})
	return cache_Foreign_Object_foldableWithIndexObject
}

var cache_Foreign_Object_traversableWithIndexObject gopurs_runtime.Value
var once_Foreign_Object_traversableWithIndexObject sync.Once

func Get_Foreign_Object_traversableWithIndexObject() gopurs_runtime.Value {
	once_Foreign_Object_traversableWithIndexObject.Do(func() {
		cache_Foreign_Object_traversableWithIndexObject = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_3467869399_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_1767849495_3725484264(Rebox_Foreign_Object_3725484264_1767849495(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Object_foldableWithIndexObject()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_2872813271_2412140840(Rebox_Foreign_Object_2412140840_2872813271(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Object_functorWithIndexObject()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Foreign_Object_traversableObject()))}
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, ms_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Foreign_Object_fold(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_Foreign_Object_mutate(gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.NormalizeStringValue(k_4), a_7), b_6)
				}), acc_3), gopurs_runtime.Apply2(f_1, gopurs_runtime.NormalizeStringValue(k_4), v_5))
			}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Foreign_Object_empty()), ms_2)
		})})))}
	})
	return cache_Foreign_Object_traversableWithIndexObject
}

var cache_Foreign_Object_traversableObject gopurs_runtime.Value
var once_Foreign_Object_traversableObject sync.Once

func Get_Foreign_Object_traversableObject() gopurs_runtime.Value {
	once_Foreign_Object_traversableObject.Do(func() {
		cache_Foreign_Object_traversableObject = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Foreign_Object_foldableObject()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Foreign_Object_functorObject()))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Foreign_Object_traversableObject()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_1
			}))
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Rebox_Foreign_Object_1812164904_3467869399(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Object_traversableWithIndexObject())).V3, dictApplicative_0), Get_Data_Function_go__const())
		})}))}
	})
	return cache_Foreign_Object_traversableObject
}

var cache_Foreign_Object_filterWithKey gopurs_runtime.Value
var once_Foreign_Object_filterWithKey sync.Once

func Get_Foreign_Object_filterWithKey() gopurs_runtime.Value {
	once_Foreign_Object_filterWithKey.Do(func() {
		cache_Foreign_Object_filterWithKey = gopurs_runtime.Func2(func(predicate_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_filterWithKey(predicate_0_box, m_1_box)
		})
	})
	return cache_Foreign_Object_filterWithKey
}

var cache_Foreign_Object_filterKeys gopurs_runtime.Value
var once_Foreign_Object_filterKeys sync.Once

func Get_Foreign_Object_filterKeys() gopurs_runtime.Value {
	once_Foreign_Object_filterKeys.Do(func() {
		cache_Foreign_Object_filterKeys = gopurs_runtime.Func(func(predicate_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_filterKeys(predicate_0_box)
		})
	})
	return cache_Foreign_Object_filterKeys
}

var cache_Foreign_Object_filter gopurs_runtime.Value
var once_Foreign_Object_filter sync.Once

func Get_Foreign_Object_filter() gopurs_runtime.Value {
	once_Foreign_Object_filter.Do(func() {
		cache_Foreign_Object_filter = gopurs_runtime.Func(func(predicate_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_filter(predicate_0_box)
		})
	})
	return cache_Foreign_Object_filter
}

var cache_Foreign_Object_eqObject gopurs_runtime.Value
var once_Foreign_Object_eqObject sync.Once

func Get_Foreign_Object_eqObject() gopurs_runtime.Value {
	once_Foreign_Object_eqObject.Do(func() {
		cache_Foreign_Object_eqObject = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_eqObject(dictEq_0_box)
		})
	})
	return cache_Foreign_Object_eqObject
}

var cache_Foreign_Object_ordObject gopurs_runtime.Value
var once_Foreign_Object_ordObject sync.Once

func Get_Foreign_Object_ordObject() gopurs_runtime.Value {
	once_Foreign_Object_ordObject.Do(func() {
		cache_Foreign_Object_ordObject = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_ordObject(dictOrd_0_box)
		})
	})
	return cache_Foreign_Object_ordObject
}

var cache_Foreign_Object_eq1Object gopurs_runtime.Value
var once_Foreign_Object_eq1Object sync.Once

func Get_Foreign_Object_eq1Object() gopurs_runtime.Value {
	once_Foreign_Object_eq1Object.Do(func() {
		cache_Foreign_Object_eq1Object = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Foreign_Object_eqObject(dictEq_0)))
		})}))}
	})
	return cache_Foreign_Object_eq1Object
}

var cache_Foreign_Object_go__delete gopurs_runtime.Value
var once_Foreign_Object_go__delete sync.Once

func Get_Foreign_Object_go__delete() gopurs_runtime.Value {
	once_Foreign_Object_go__delete.Do(func() {
		cache_Foreign_Object_go__delete = gopurs_runtime.Func(func(k_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_go__delete(k_0_box.StrVal())
		})
	})
	return cache_Foreign_Object_go__delete
}

var cache_Foreign_Object_pop gopurs_runtime.Value
var once_Foreign_Object_pop sync.Once

func Get_Foreign_Object_pop() gopurs_runtime.Value {
	once_Foreign_Object_pop.Do(func() {
		cache_Foreign_Object_pop = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Foreign_Object_pop(k_0_box.StrVal(), m_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Foreign_Object_pop
}

var cache_Foreign_Object_alter gopurs_runtime.Value
var once_Foreign_Object_alter sync.Once

func Get_Foreign_Object_alter() gopurs_runtime.Value {
	once_Foreign_Object_alter.Do(func() {
		cache_Foreign_Object_alter = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_alter(f_0_box, k_1_box.StrVal(), m_2_box)
		})
	})
	return cache_Foreign_Object_alter
}

var cache_Foreign_Object_update gopurs_runtime.Value
var once_Foreign_Object_update sync.Once

func Get_Foreign_Object_update() gopurs_runtime.Value {
	once_Foreign_Object_update.Do(func() {
		cache_Foreign_Object_update = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, m_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Foreign_Object_update(f_0_box, k_1_box.StrVal(), m_2_box)
		})
	})
	return cache_Foreign_Object_update
}

func Call_Foreign_Object_bindFlipped(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_2_0 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
		_ = __local_var_2_0
		return gopurs_runtime.Apply2(b_0, __local_var_2_0, gopurs_runtime.Value{})
	})
}

func Call_Foreign_Object_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Foreign_Object_mapFlipped(fa_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var fa_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = fa_0_loop
	_ = fa_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if fa_0 != nil {
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Apply(f_1, (fa_0).V0), true}
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
			}{gopurs_runtime.Value{}, false}
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

func Call_Foreign_Object_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Array_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)}), gopurs_runtime.Apply(Get_Foreign_Object_toArrayWithKey(), Get_Data_Tuple_Tuple()))
}

func Call_Foreign_Object_toAscUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Array_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Array_sortWith(Rebox_Foreign_Object_2406510097_4177771502(Rebox_Foreign_Object_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString()))), Get_Data_Tuple_fst()), gopurs_runtime.Apply(Get_Foreign_Object_toArrayWithKey(), Get_Data_Tuple_Tuple())))
}

func Call_Foreign_Object_singleton(k_0_loop string, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var k_0 string = k_0_loop
	_ = k_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(ADT ["Foreign","Object","ST","STObject"] [(TypeVar r$scope55), (TypeVar a$scope52)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope55), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r$scope55), (TypeVar a$scope52)])]))
		__local_var_2_0 := gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str(k_0), v_1)
		_ = __local_var_2_0
		__local_var_3_1 := gopurs_runtime.Apply(Get_Foreign_Object_ST_go__new(), gopurs_runtime.Value{})
		_ = __local_var_3_1
		return gopurs_runtime.Apply2(__local_var_2_0, __local_var_3_1, gopurs_runtime.Value{})
	}))
}

func Call_Foreign_Object_singleton__2505508685(k_0_loop string, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
singleton__2505508685:
	for {
		if false {
			continue singleton__2505508685
		}
		var k_0 string = k_0_loop
		_ = k_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(ADT ["Foreign","Object","ST","STObject"] [(TypeVar r$scope55), (ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope55), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r$scope55), (ADT ["Data","Argonaut","Core","Json"] [])])]))
			__local_var_2_0 := gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str(k_0), v_1)
			_ = __local_var_2_0
			__local_var_3_1 := gopurs_runtime.Apply(Get_Foreign_Object_ST_go__new(), gopurs_runtime.Value{})
			_ = __local_var_3_1
			return gopurs_runtime.Apply2(__local_var_2_0, __local_var_3_1, gopurs_runtime.Value{})
		}))
	}
}

func Call_Foreign_Object_showObject(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(fromFoldable ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(Call_Data_Tuple_showTuple(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_1514099793_1386611502(Rebox_Foreign_Object_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}, dictShow_0)), "show"), func() gopurs_runtime.Value {
			arr := func() []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Foreign_Object_toArray(), m_1).UnsafePtr)
				unboxed := make([]*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Foreign_Object_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_3103076375_138441832(v))}
			}
			return gopurs_runtime.Array(boxed)
		}()).StrVal()), ")"))
	})}))}
}

func Call_Foreign_Object_mutate(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	return gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Foreign_Object__copyST(), m_1)
		_ = __local_var_2_0
		s_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
		_ = s_3_1
		_dollar___unused_4_2 := gopurs_runtime.Apply2(f_0, s_3_1, gopurs_runtime.Value{})
		_ = _dollar___unused_4_2
		return s_3_1
	}))
}

func Call_Foreign_Object_member(__local_var_0_loop string, __local_var_1_loop gopurs_runtime.Value) bool {
	var __local_var_0 string = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return (gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Bool(false), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(true)
	}), gopurs_runtime.Str(__local_var_0), __local_var_1).IntVal) != (0)
}

func Call_Foreign_Object_mapWithKey(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	return gopurs_runtime.UncurriedApp2(Get_Foreign_Object__mapWithKey(), m_1, f_0)
}

func Call_Foreign_Object_lookup(__local_var_0_loop string, __local_var_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var __local_var_0 string = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.Str(__local_var_0), __local_var_1); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Foreign_Object_lookup__1777685996(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
lookup__1777685996:
	for {
		if false {
			continue lookup__1777685996
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.Str(__eta_norm_1_0), __eta_norm_0_1)))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Foreign_Object_lookup__2830175038(__eta_norm_0_0_loop string, __local_var_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
lookup__2830175038:
	for {
		if false {
			continue lookup__2830175038
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
		_ = __local_var_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_742090555_3094389156(Rebox_Foreign_Object_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}, Get_Data_Maybe_Just(), gopurs_runtime.Str(__eta_norm_0_0), __local_var_1); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Foreign_Object_isSubmap(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], m1_1_loop gopurs_runtime.Value, m2_2_loop gopurs_runtime.Value) bool {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var m1_1 gopurs_runtime.Value = m1_1_loop
	_ = m1_1
	var m2_2 gopurs_runtime.Value = m2_2_loop
	_ = m2_2
	return (gopurs_runtime.Apply2(Get_Foreign_Object_all(), gopurs_runtime.Func2(func(k_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Bool(false), gopurs_runtime.Apply(dictEq_0.V0, v_4), gopurs_runtime.NormalizeStringValue(k_3), m2_2)
	}), m1_1).IntVal) != (0)
}

func Call_Foreign_Object_insert(k_0_loop string, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var k_0 string = k_0_loop
	_ = k_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(Get_Foreign_Object_mutate(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str(k_0), v_1))
}

func Call_Foreign_Object_insert__3487692426(k_0_loop string, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
insert__3487692426:
	for {
		if false {
			continue insert__3487692426
		}
		var k_0 string = k_0_loop
		_ = k_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply(Get_Foreign_Object_mutate(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str(k_0), v_1))
	}
}

func Call_Foreign_Object_insert__1072177850(k_0_loop string, v_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insert__1072177850:
	for {
		if false {
			continue insert__1072177850
		}
		var k_0 string = k_0_loop
		_ = k_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return Call_Foreign_Object_mutate(gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str(k_0), v_1), __eta_norm_0_2)
	}
}

func Call_Foreign_Object_fromHomogeneous(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Unsafe_Coerce_unsafeCoerce()
}

func Call_Foreign_Object_fromFoldableWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[string, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[string, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
	_ = dictFoldableWithIndex_0
	// TAST (Let): forWithIndex_1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope147) [(TypeVar a$scope148)]), (Func [String, (TypeVar a$scope148)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope151), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r$scope151), (TypeVar a$scope148)])]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope151), Unit]))
	forWithIndex_1_1_0 := gopurs_runtime.Apply(Call_Data_FoldableWithIndex_forWithIndex_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST())), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_1767849495_3725484264(dictFoldableWithIndex_0))})
	_ = forWithIndex_1_1_0
	return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			s_3_1 := gopurs_runtime.Apply(Get_Foreign_Object_ST_go__new(), gopurs_runtime.Value{})
			_ = s_3_1
			_dollar___unused_4_2 := gopurs_runtime.Apply3(forWithIndex_1_1_0, l_2, gopurs_runtime.Func2(func(k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply3(Get_Foreign_Object_ST_poke(), gopurs_runtime.NormalizeStringValue(k_4), v_5, s_3_1)
			}), gopurs_runtime.Value{})
			_ = _dollar___unused_4_2
			return s_3_1
		}))
	})
}

func Call_Foreign_Object_fromFoldableWith(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): for_1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope156) [(ADT ["Data","Tuple","Tuple"] [String, (TypeVar a$scope157)])]), (Func [(ADT ["Data","Tuple","Tuple"] [String, (TypeVar a$scope157)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope160), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r$scope160), (TypeVar a$scope157)])]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope160), Unit]))
	for_1_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_for_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST())), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)})
	_ = for_1_1_0
	return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, l_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			s_4_1 := gopurs_runtime.Apply(Get_Foreign_Object_ST_go__new(), gopurs_runtime.Value{})
			_ = s_4_1
			_dollar___unused_5_2 := gopurs_runtime.Apply3(for_1_1_0, l_3, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
					__local_var_6_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
					_ = __local_var_6_3
					// TAST (Let): __local_var_7_4 shape=UncurriedApp(Var) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r$scope160), (TypeVar a$scope157)])
					__local_var_7_4 := gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookupST(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1), gopurs_runtime.NormalizeStringValue(__local_var_6_3), s_4_1)
					_ = __local_var_7_4
					v_prime__8_5 := gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Value{})
					_ = v_prime__8_5
					return gopurs_runtime.Apply(gopurs_runtime.Apply3(Get_Foreign_Object_ST_poke(), gopurs_runtime.NormalizeStringValue(__local_var_6_3), v_prime__8_5, s_4_1), gopurs_runtime.Value{})
				})
			}), gopurs_runtime.Value{})
			_ = _dollar___unused_5_2
			return s_4_1
		}))
	})
}

func Call_Foreign_Object_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope253), (TypeVar b$scope260)] (TypeVar b$scope260)), (TypeVar b$scope260), (TypeApp (TypeVar f$scope173) [(TypeVar a$scope253)])] (TypeVar b$scope260))
	__local_var_1_0 := Call_Data_Foldable_foldr(dictFoldable_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			s_3_1 := gopurs_runtime.Apply(Get_Foreign_Object_ST_go__new(), gopurs_runtime.Value{})
			_ = s_3_1
			_dollar___unused_4_2 := gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_foreach(), func() gopurs_runtime.Value {
				arr := func() []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_0, l_2).UnsafePtr))).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Foreign_Object_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_3103076375_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_functorST())), gopurs_runtime.Apply3(Get_Foreign_Object_ST_poke(), gopurs_runtime.NormalizeStringValue((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, s_3_1))
			}), gopurs_runtime.Value{})
			_ = _dollar___unused_4_2
			return s_3_1
		}))
	})
}

func Call_Foreign_Object_foldMaybe(f_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var z_1 gopurs_runtime.Value = z_1_loop
	_ = z_1
	var m_2 gopurs_runtime.Value = m_2_loop
	_ = m_2
	return gopurs_runtime.UncurriedApp4(Get_Foreign_Object__foldSCObject(), m_2, z_1, f_0, Get_Data_Maybe_fromMaybe())
}

func Call_Foreign_Object_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): bind7_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope206) [(TypeVar z$scope207)]), (Func [(TypeVar z$scope207)] (TypeApp (TypeVar m$scope206) [(TypeVar z$scope207)]))] (TypeApp (TypeVar m$scope206) [(TypeVar z$scope207)]))
	bind7_1_0 := Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{})))
	_ = bind7_1_0
	return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Get_Foreign_Object__foldM(), bind7_1_0, f_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}), "pure"), z_3))
	})
}

func Call_Foreign_Object_union(m_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	return gopurs_runtime.Apply(Get_Foreign_Object_mutate(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Foreign_Object_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST())), gopurs_runtime.Func3(func(s_prime__2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Foreign_Object_ST_poke(), gopurs_runtime.NormalizeStringValue(k_3), v_4, s_prime__2)
		}), s_1, m_0)
	}))
}

func Call_Foreign_Object_unions(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(dictFoldable_0.V1, Get_Foreign_Object_union(), Get_Foreign_Object_empty())
}

func Call_Foreign_Object_unionWith(f_0_loop gopurs_runtime.Value, m1_1_loop gopurs_runtime.Value, m2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var m1_1 gopurs_runtime.Value = m1_1_loop
	_ = m1_1
	var m2_2 gopurs_runtime.Value = m2_2_loop
	_ = m2_2
	return Call_Foreign_Object_mutate(gopurs_runtime.Func(func(s1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Call_Foreign_Object_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST())), gopurs_runtime.Func3(func(s2_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Foreign_Object_ST_poke(), gopurs_runtime.NormalizeStringValue(k_5), gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), v1_6, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(f_0, v1_6, v2_7)
			}), gopurs_runtime.NormalizeStringValue(k_5), m2_2), s2_4)
		}), s1_3, m1_1)
	}), m2_2)
}

func Call_Foreign_Object_semigroupObject(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Foreign_Object_unionWith(), Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0)))}))}
}

func Call_Foreign_Object_monoidObject(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): semigroupObject1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Foreign","Object","Object"] [(TypeVar a$scope73)])])
	semigroupObject1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Foreign_Object_semigroupObject(dictSemigroup_0))
	_ = semigroupObject1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupObject1_1_0)}
	}), Get_Foreign_Object_empty()}))}
}

func Call_Foreign_Object_foldMap(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope246)
	mempty_1_0 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)})
	_ = mempty_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Foreign_Object_fold(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonoid_0.V0, gopurs_runtime.Value{}), "append"), acc_3, gopurs_runtime.Apply2(f_2, gopurs_runtime.NormalizeStringValue(k_4), v_5))
		}), mempty_1_0)
	})
}

func Call_Foreign_Object_filterWithKey(predicate_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var predicate_0 gopurs_runtime.Value = predicate_0_loop
	_ = predicate_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	return gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		m_prime__2_0 := gopurs_runtime.Apply(Get_Foreign_Object_ST_go__new(), gopurs_runtime.Value{})
		_ = m_prime__2_0
		return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Foreign_Object_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST())), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if (gopurs_runtime.Apply2(predicate_0, gopurs_runtime.NormalizeStringValue(k_4), v_5).IntVal) != (0) {
					__t1 = gopurs_runtime.Apply3(Get_Foreign_Object_ST_poke(), gopurs_runtime.NormalizeStringValue(k_4), v_5, acc_3)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return acc_3
				})
			}
		end_branch_1:
			return __t1
		}), m_prime__2_0, m_1), gopurs_runtime.Value{})
	}))
}

func Call_Foreign_Object_filterKeys(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var predicate_0 gopurs_runtime.Value = predicate_0_loop
	_ = predicate_0
	return gopurs_runtime.Apply(Get_Foreign_Object_filterWithKey(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Function_go__const(), predicate_0))
}

func Call_Foreign_Object_filter(predicate_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var predicate_0 gopurs_runtime.Value = predicate_0_loop
	_ = predicate_0
	return gopurs_runtime.Apply(Get_Foreign_Object_filterWithKey(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return predicate_0
	}))
}

func Call_Foreign_Object_eqObject(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(m1_1 gopurs_runtime.Value, m2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool((Call_Foreign_Object_isSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0), m1_1, m2_2)) && (Call_Foreign_Object_isSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0), m2_2, m1_1)))
	})}))}
}

func Call_Foreign_Object_ordObject(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqObject1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Foreign","Object","Object"] [(TypeVar a$scope62)])])
	eqObject1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Foreign_Object_eqObject(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
	_ = eqObject1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqObject1_1_0)}
	}), gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Ord_ordArray(gopurs_runtime.Apply(Get_Foreign_Object_ordTuple(), dictOrd_0)), "compare"), func() gopurs_runtime.Value {
			arr := func() []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Foreign_Object_toAscUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray())), m1_2).UnsafePtr)
				unboxed := make([]*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Foreign_Object_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_3103076375_138441832(v))}
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := func() []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Foreign_Object_toAscUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray())), m2_3).UnsafePtr)
				unboxed := make([]*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Foreign_Object_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Foreign_Object_3103076375_138441832(v))}
			}
			return gopurs_runtime.Array(boxed)
		}())
	})}))}
}

func Call_Foreign_Object_go__delete(k_0_loop string) gopurs_runtime.Value {
	var k_0 string = k_0_loop
	_ = k_0
	return gopurs_runtime.Apply(Get_Foreign_Object_mutate(), gopurs_runtime.Apply(Get_Foreign_Object_ST_deleteImpl(), gopurs_runtime.Str(k_0)))
}

func Call_Foreign_Object_pop(k_0_loop string, m_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var k_0 string = k_0_loop
	_ = k_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	// TAST (Let): __local_var_2_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope311)])
	__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.Str(k_0), m_1))
	_ = __local_var_2_0
	var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if __local_var_2_0 != nil {
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(__local_var_2_0).V0, Call_Foreign_Object_mutate(gopurs_runtime.Apply(Get_Foreign_Object_ST_deleteImpl(), gopurs_runtime.Str(k_0)), m_1)}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}(), true}
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

func Call_Foreign_Object_alter(f_0_loop gopurs_runtime.Value, k_1_loop string, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var k_1 string = k_1_loop
	_ = k_1
	var m_2 gopurs_runtime.Value = m_2_loop
	_ = m_2
	// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope315)])
	v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.Str(k_1), m_2)))}))
	_ = v_3_0
	var __t1 gopurs_runtime.Value
	{
		if v_3_0 == nil {
			__t1 = Call_Foreign_Object_mutate(gopurs_runtime.Apply(Get_Foreign_Object_ST_deleteImpl(), gopurs_runtime.Str(k_1)), m_2)
			goto end_branch_1
		} else {

		}
	}
	{
		if v_3_0 != nil {
			__t1 = Call_Foreign_Object_mutate(gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str(k_1), (v_3_0).V0), m_2)
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

func Call_Foreign_Object_update(f_0_loop gopurs_runtime.Value, k_1_loop string, m_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var k_1 string = k_1_loop
	_ = k_1
	var m_2 gopurs_runtime.Value = m_2_loop
	_ = m_2
	return Call_Foreign_Object_alter(gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				goto end_branch_0
			} else {

			}
		}
		{
			if v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil {
				__t0 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	}), k_1, m_2)
}

func Rebox_Foreign_Object_138441832_3103076375(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1
	return out
}

func Rebox_Foreign_Object_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_1767849495_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[string, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_1812164904_3467869399(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_2412140840_2872813271(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[string, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_2872813271_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[string, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Foreign_Object_3103076375_138441832(in *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Foreign_Object_3467869399_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_3725484264_1767849495(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[string, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}

func Rebox_Foreign_Object_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Get_Foreign_Object__copyST() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object__CopyST
}

func Get_Foreign_Object__fmapObject() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object__FmapObject
}

func Get_Foreign_Object__foldM() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object__FoldM
}

func Get_Foreign_Object__foldSCObject() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object__FoldSCObject
}

func Get_Foreign_Object__lookup() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object__Lookup
}

func Get_Foreign_Object__lookupST() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object__LookupST
}

func Get_Foreign_Object__mapWithKey() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object__MapWithKey
}

func Get_Foreign_Object_all() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_All
}

func Get_Foreign_Object_empty() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_Empty
}

func Get_Foreign_Object_keys() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_Keys
}

func Get_Foreign_Object_runST() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_RunST
}

func Get_Foreign_Object_size() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_Size
}

func Get_Foreign_Object_toArrayWithKey() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_ToArrayWithKey
}
