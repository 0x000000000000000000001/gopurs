package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Map_Internal_fromFoldable1 gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldable1 sync.Once

func Get_Data_Map_Internal_fromFoldable1() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldable1.Do(func() {
		cache_Data_Map_Internal_fromFoldable1 = gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_Map_Internal_fromFoldable1
}

var cache_Data_Map_Internal_go__map gopurs_runtime.Value
var once_Data_Map_Internal_go__map sync.Once

func Get_Data_Map_Internal_go__map() gopurs_runtime.Value {
	once_Data_Map_Internal_go__map.Do(func() {
		cache_Data_Map_Internal_go__map = Rebox_Data_Map_Internal_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Map_Internal_go__map
}

var cache_Data_Map_Internal_map__870244361 gopurs_runtime.Value
var once_Data_Map_Internal_map__870244361 sync.Once

func Get_Data_Map_Internal_map__870244361() gopurs_runtime.Value {
	once_Data_Map_Internal_map__870244361.Do(func() {
		cache_Data_Map_Internal_map__870244361 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_map__870244361(__eta_norm_1_0_box, Rebox_Data_Map_Internal_3094389156_2304530460(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Map_Internal_map__870244361
}

var cache_Data_Map_Internal_add gopurs_runtime.Value
var once_Data_Map_Internal_add sync.Once

func Get_Data_Map_Internal_add() gopurs_runtime.Value {
	once_Data_Map_Internal_add.Do(func() {
		cache_Data_Map_Internal_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_Map_Internal_add
}

var cache_Data_Map_Internal_go__append gopurs_runtime.Value
var once_Data_Map_Internal_go__append sync.Once

func Get_Data_Map_Internal_go__append() gopurs_runtime.Value {
	once_Data_Map_Internal_go__append.Do(func() {
		cache_Data_Map_Internal_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Map_Internal_go__append
}

var cache_Data_Map_Internal_conj gopurs_runtime.Value
var once_Data_Map_Internal_conj sync.Once

func Get_Data_Map_Internal_conj() gopurs_runtime.Value {
	once_Data_Map_Internal_conj.Do(func() {
		cache_Data_Map_Internal_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Map_Internal_conj
}

var cache_Data_Map_Internal_eq gopurs_runtime.Value
var once_Data_Map_Internal_eq sync.Once

func Get_Data_Map_Internal_eq() gopurs_runtime.Value {
	once_Data_Map_Internal_eq.Do(func() {
		cache_Data_Map_Internal_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_Map_Internal_eq
}

var cache_Data_Map_Internal_negate gopurs_runtime.Value
var once_Data_Map_Internal_negate sync.Once

func Get_Data_Map_Internal_negate() gopurs_runtime.Value {
	once_Data_Map_Internal_negate.Do(func() {
		cache_Data_Map_Internal_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Map_Internal_negate(a_0_box.IntVal))
		})
	})
	return cache_Data_Map_Internal_negate
}

var cache_Data_Map_Internal_compose gopurs_runtime.Value
var once_Data_Map_Internal_compose sync.Once

func Get_Data_Map_Internal_compose() gopurs_runtime.Value {
	once_Data_Map_Internal_compose.Do(func() {
		cache_Data_Map_Internal_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Map_Internal_compose
}

var cache_Data_Map_Internal_disj gopurs_runtime.Value
var once_Data_Map_Internal_disj sync.Once

func Get_Data_Map_Internal_disj() gopurs_runtime.Value {
	once_Data_Map_Internal_disj.Do(func() {
		cache_Data_Map_Internal_disj = Get_Data_HeytingAlgebra_disj__2421783737()
	})
	return cache_Data_Map_Internal_disj
}

var cache_Data_Map_Internal_values gopurs_runtime.Value
var once_Data_Map_Internal_values sync.Once

func Get_Data_Map_Internal_values() gopurs_runtime.Value {
	once_Data_Map_Internal_values.Do(func() {
		cache_Data_Map_Internal_values = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_values(m_0_box))}
		})
	})
	return cache_Data_Map_Internal_values
}

var cache_Data_Map_Internal_toUnfoldable gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable sync.Once

func Get_Data_Map_Internal_toUnfoldable() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable.Do(func() {
		cache_Data_Map_Internal_toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), m_1_box)
		})
	})
	return cache_Data_Map_Internal_toUnfoldable
}

var cache_Data_Map_Internal_toUnfoldable__3206239856 gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable__3206239856 sync.Once

func Get_Data_Map_Internal_toUnfoldable__3206239856() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable__3206239856.Do(func() {
		cache_Data_Map_Internal_toUnfoldable__3206239856 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2442833393_849153993(Call_Data_Map_Internal_toUnfoldable__3206239856(__eta0_0_box)))}
		})
	})
	return cache_Data_Map_Internal_toUnfoldable__3206239856
}

var cache_Data_Map_Internal_toUnfoldable1 gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable1 sync.Once

func Get_Data_Map_Internal_toUnfoldable1() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable1.Do(func() {
		cache_Data_Map_Internal_toUnfoldable1 = gopurs_runtime.Apply(Get_Data_Map_Internal_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()))})
	})
	return cache_Data_Map_Internal_toUnfoldable1
}

var cache_Data_Map_Internal_toUnfoldable2 gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldable2 sync.Once

func Get_Data_Map_Internal_toUnfoldable2() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldable2.Do(func() {
		cache_Data_Map_Internal_toUnfoldable2 = gopurs_runtime.Apply(Get_Data_Map_Internal_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_132451362_2738507278(Rebox_Data_Map_Internal_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})
	})
	return cache_Data_Map_Internal_toUnfoldable2
}

var cache_Data_Map_Internal_toUnfoldableUnordered gopurs_runtime.Value
var once_Data_Map_Internal_toUnfoldableUnordered sync.Once

func Get_Data_Map_Internal_toUnfoldableUnordered() gopurs_runtime.Value {
	once_Data_Map_Internal_toUnfoldableUnordered.Do(func() {
		cache_Data_Map_Internal_toUnfoldableUnordered = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_toUnfoldableUnordered(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_Map_Internal_toUnfoldableUnordered
}

var cache_Data_Map_Internal_size gopurs_runtime.Value
var once_Data_Map_Internal_size sync.Once

func Get_Data_Map_Internal_size() gopurs_runtime.Value {
	once_Data_Map_Internal_size.Do(func() {
		cache_Data_Map_Internal_size = Get_Data_Map_Internal_sizeImpl()
	})
	return cache_Data_Map_Internal_size
}

var cache_Data_Map_Internal_size__2179923709 gopurs_runtime.Value
var once_Data_Map_Internal_size__2179923709 sync.Once

func Get_Data_Map_Internal_size__2179923709() gopurs_runtime.Value {
	once_Data_Map_Internal_size__2179923709.Do(func() {
		cache_Data_Map_Internal_size__2179923709 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Map_Internal_size__2179923709(__eta_norm_0_0_box))
		})
	})
	return cache_Data_Map_Internal_size__2179923709
}

var cache_Data_Map_Internal_showMap gopurs_runtime.Value
var once_Data_Map_Internal_showMap sync.Once

func Get_Data_Map_Internal_showMap() gopurs_runtime.Value {
	once_Data_Map_Internal_showMap.Do(func() {
		cache_Data_Map_Internal_showMap = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_showMap(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Data_Map_Internal_showMap
}

var cache_Data_Map_Internal_lookupLT gopurs_runtime.Value
var once_Data_Map_Internal_lookupLT sync.Once

func Get_Data_Map_Internal_lookupLT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupLT.Do(func() {
		cache_Data_Map_Internal_lookupLT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_lookupLT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
		})
	})
	return cache_Data_Map_Internal_lookupLT
}

var cache_Data_Map_Internal_lookupLE gopurs_runtime.Value
var once_Data_Map_Internal_lookupLE sync.Once

func Get_Data_Map_Internal_lookupLE() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupLE.Do(func() {
		cache_Data_Map_Internal_lookupLE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_lookupLE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
		})
	})
	return cache_Data_Map_Internal_lookupLE
}

var cache_Data_Map_Internal_lookupGT gopurs_runtime.Value
var once_Data_Map_Internal_lookupGT sync.Once

func Get_Data_Map_Internal_lookupGT() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGT.Do(func() {
		cache_Data_Map_Internal_lookupGT = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_lookupGT(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
		})
	})
	return cache_Data_Map_Internal_lookupGT
}

var cache_Data_Map_Internal_lookupGE gopurs_runtime.Value
var once_Data_Map_Internal_lookupGE sync.Once

func Get_Data_Map_Internal_lookupGE() gopurs_runtime.Value {
	once_Data_Map_Internal_lookupGE.Do(func() {
		cache_Data_Map_Internal_lookupGE = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_lookupGE(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
		})
	})
	return cache_Data_Map_Internal_lookupGE
}

var cache_Data_Map_Internal_keys gopurs_runtime.Value
var once_Data_Map_Internal_keys sync.Once

func Get_Data_Map_Internal_keys() gopurs_runtime.Value {
	once_Data_Map_Internal_keys.Do(func() {
		cache_Data_Map_Internal_keys = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Map_Internal_keys(m_0_box))}
		})
	})
	return cache_Data_Map_Internal_keys
}

var cache_Data_Map_Internal_functorMap gopurs_runtime.Value
var once_Data_Map_Internal_functorMap sync.Once

func Get_Data_Map_Internal_functorMap() gopurs_runtime.Value {
	once_Data_Map_Internal_functorMap.Do(func() {
		cache_Data_Map_Internal_functorMap = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Get_Data_Map_Internal_mapImpl()}))}
	})
	return cache_Data_Map_Internal_functorMap
}

var cache_Data_Map_Internal_foldableMap gopurs_runtime.Value
var once_Data_Map_Internal_foldableMap sync.Once

func Get_Data_Map_Internal_foldableMap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableMap.Do(func() {
		cache_Data_Map_Internal_foldableMap = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope152)
			mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
			_ = mempty_1_0
			return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Map_Internal_foldableMap(), "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_4, gopurs_runtime.Apply(f_2, v_5))
				}), mempty_1_0, m_3)
			})
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(f_0, acc_3, v1_5)
			}), z_1, m_2)
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldrImpl(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(f_0, v1_5, acc_3)
			}), z_1, m_2)
		})}))}
	})
	return cache_Data_Map_Internal_foldableMap
}

var cache_Data_Map_Internal_foldableWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_foldableWithIndexMap sync.Once

func Get_Data_Map_Internal_foldableWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_foldableWithIndexMap.Do(func() {
		cache_Data_Map_Internal_foldableWithIndexMap = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Map_Internal_foldableMap()
		}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope128)
			mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
			_ = mempty_1_0
			return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_4, gopurs_runtime.Apply2(f_2, k_5, v_6))
				}), mempty_1_0, m_3)
			})
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply3(f_0, k_4, acc_3, v_5)
			}), z_1, m_2)
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldrImpl(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply3(f_0, k_4, v_5, acc_3)
			}), z_1, m_2)
		})}))}
	})
	return cache_Data_Map_Internal_foldableWithIndexMap
}

var cache_Data_Map_Internal_findMin gopurs_runtime.Value
var once_Data_Map_Internal_findMin sync.Once

func Get_Data_Map_Internal_findMin() gopurs_runtime.Value {
	once_Data_Map_Internal_findMin.Do(func() {
		cache_Data_Map_Internal_findMin = gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if acc_0.Type == 9 && acc_0.IntVal == 930809136 && acc_0.UnsafePtr == nil {
					__t0 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.RecordDict2("key", "value", k_1, v_2), true}
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
				if acc_0.Type == 9 && acc_0.IntVal == 930809136 && acc_0.UnsafePtr != nil {
					__t0 = acc_0
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(func() *Constructor_Data_Maybe_Just[struct {
					key   gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
				}] {
					panic("Failed pattern match")
				}()))}
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 struct {
					key   gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
				}
				V1 bool
			}{struct {
				key   gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
			}{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))})
	})
	return cache_Data_Map_Internal_findMin
}

var cache_Data_Map_Internal_findMax gopurs_runtime.Value
var once_Data_Map_Internal_findMax sync.Once

func Get_Data_Map_Internal_findMax() gopurs_runtime.Value {
	once_Data_Map_Internal_findMax.Do(func() {
		cache_Data_Map_Internal_findMax = gopurs_runtime.Apply2(Get_Data_Map_Internal_foldrImpl(), gopurs_runtime.Func3(func(acc_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if acc_0.Type == 9 && acc_0.IntVal == 930809136 && acc_0.UnsafePtr == nil {
					__t0 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.RecordDict2("key", "value", k_1, v_2), true}
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
				if acc_0.Type == 9 && acc_0.IntVal == 930809136 && acc_0.UnsafePtr != nil {
					__t0 = acc_0
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(func() *Constructor_Data_Maybe_Just[struct {
					key   gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
				}] {
					panic("Failed pattern match")
				}()))}
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 struct {
					key   gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
				}
				V1 bool
			}{struct {
				key   gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
			}{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
				}()})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))})
	})
	return cache_Data_Map_Internal_findMax
}

var cache_Data_Map_Internal_filterKeys gopurs_runtime.Value
var once_Data_Map_Internal_filterKeys sync.Once

func Get_Data_Map_Internal_filterKeys() gopurs_runtime.Value {
	once_Data_Map_Internal_filterKeys.Do(func() {
		cache_Data_Map_Internal_filterKeys = Get_Data_Map_Internal_filterKeysImpl()
	})
	return cache_Data_Map_Internal_filterKeys
}

var cache_Data_Map_Internal_eqMap gopurs_runtime.Value
var once_Data_Map_Internal_eqMap sync.Once

func Get_Data_Map_Internal_eqMap() gopurs_runtime.Value {
	once_Data_Map_Internal_eqMap.Do(func() {
		cache_Data_Map_Internal_eqMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_eqMap(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Data_Map_Internal_eqMap
}

var cache_Data_Map_Internal_ordMap gopurs_runtime.Value
var once_Data_Map_Internal_ordMap sync.Once

func Get_Data_Map_Internal_ordMap() gopurs_runtime.Value {
	once_Data_Map_Internal_ordMap.Do(func() {
		cache_Data_Map_Internal_ordMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_ordMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_ordMap
}

var cache_Data_Map_Internal_compareInt gopurs_runtime.Value
var once_Data_Map_Internal_compareInt sync.Once

func Get_Data_Map_Internal_compareInt() gopurs_runtime.Value {
	once_Data_Map_Internal_compareInt.Do(func() {
		cache_Data_Map_Internal_compareInt = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Map_Internal_compareInt(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, b_2_box))
		})
	})
	return cache_Data_Map_Internal_compareInt
}

var cache_Data_Map_Internal_go__delete gopurs_runtime.Value
var once_Data_Map_Internal_go__delete sync.Once

func Get_Data_Map_Internal_go__delete() gopurs_runtime.Value {
	once_Data_Map_Internal_go__delete.Do(func() {
		cache_Data_Map_Internal_go__delete = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_go__delete
}

var cache_Data_Map_Internal_difference gopurs_runtime.Value
var once_Data_Map_Internal_difference sync.Once

func Get_Data_Map_Internal_difference() gopurs_runtime.Value {
	once_Data_Map_Internal_difference.Do(func() {
		cache_Data_Map_Internal_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_difference
}

var cache_Data_Map_Internal_insert gopurs_runtime.Value
var once_Data_Map_Internal_insert sync.Once

func Get_Data_Map_Internal_insert() gopurs_runtime.Value {
	once_Data_Map_Internal_insert.Do(func() {
		cache_Data_Map_Internal_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_insert
}

var cache_Data_Map_Internal_filterWithKey gopurs_runtime.Value
var once_Data_Map_Internal_filterWithKey sync.Once

func Get_Data_Map_Internal_filterWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_filterWithKey.Do(func() {
		cache_Data_Map_Internal_filterWithKey = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_filterWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_filterWithKey
}

var cache_Data_Map_Internal_filter gopurs_runtime.Value
var once_Data_Map_Internal_filter sync.Once

func Get_Data_Map_Internal_filter() gopurs_runtime.Value {
	once_Data_Map_Internal_filter.Do(func() {
		cache_Data_Map_Internal_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_filter
}

var cache_Data_Map_Internal_submap gopurs_runtime.Value
var once_Data_Map_Internal_submap sync.Once

func Get_Data_Map_Internal_submap() gopurs_runtime.Value {
	once_Data_Map_Internal_submap.Do(func() {
		cache_Data_Map_Internal_submap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_submap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_submap
}

var cache_Data_Map_Internal_fromFoldable gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldable sync.Once

func Get_Data_Map_Internal_fromFoldable() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldable.Do(func() {
		cache_Data_Map_Internal_fromFoldable = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_fromFoldable
}

var cache_Data_Map_Internal_fromFoldableWithIndex gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldableWithIndex sync.Once

func Get_Data_Map_Internal_fromFoldableWithIndex() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldableWithIndex.Do(func() {
		cache_Data_Map_Internal_fromFoldableWithIndex = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_fromFoldableWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_fromFoldableWithIndex
}

var cache_Data_Map_Internal_functorWithIndexMap gopurs_runtime.Value
var once_Data_Map_Internal_functorWithIndexMap sync.Once

func Get_Data_Map_Internal_functorWithIndexMap() gopurs_runtime.Value {
	once_Data_Map_Internal_functorWithIndexMap.Do(func() {
		cache_Data_Map_Internal_functorWithIndexMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_functorWithIndexMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_functorWithIndexMap
}

var cache_Data_Map_Internal_mapMaybeWithKey gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybeWithKey sync.Once

func Get_Data_Map_Internal_mapMaybeWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybeWithKey.Do(func() {
		cache_Data_Map_Internal_mapMaybeWithKey = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_mapMaybeWithKey
}

var cache_Data_Map_Internal_mapMaybe gopurs_runtime.Value
var once_Data_Map_Internal_mapMaybe sync.Once

func Get_Data_Map_Internal_mapMaybe() gopurs_runtime.Value {
	once_Data_Map_Internal_mapMaybe.Do(func() {
		cache_Data_Map_Internal_mapMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_mapMaybe
}

var cache_Data_Map_Internal_insertWith gopurs_runtime.Value
var once_Data_Map_Internal_insertWith sync.Once

func Get_Data_Map_Internal_insertWith() gopurs_runtime.Value {
	once_Data_Map_Internal_insertWith.Do(func() {
		cache_Data_Map_Internal_insertWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_insertWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_insertWith
}

var cache_Data_Map_Internal_fromFoldableWith gopurs_runtime.Value
var once_Data_Map_Internal_fromFoldableWith sync.Once

func Get_Data_Map_Internal_fromFoldableWith() gopurs_runtime.Value {
	once_Data_Map_Internal_fromFoldableWith.Do(func() {
		cache_Data_Map_Internal_fromFoldableWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_fromFoldableWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_fromFoldableWith
}

var cache_Data_Map_Internal_intersectionWith gopurs_runtime.Value
var once_Data_Map_Internal_intersectionWith sync.Once

func Get_Data_Map_Internal_intersectionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_intersectionWith.Do(func() {
		cache_Data_Map_Internal_intersectionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_intersectionWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_intersectionWith
}

var cache_Data_Map_Internal_intersection gopurs_runtime.Value
var once_Data_Map_Internal_intersection sync.Once

func Get_Data_Map_Internal_intersection() gopurs_runtime.Value {
	once_Data_Map_Internal_intersection.Do(func() {
		cache_Data_Map_Internal_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_intersection
}

var cache_Data_Map_Internal_lookup gopurs_runtime.Value
var once_Data_Map_Internal_lookup sync.Once

func Get_Data_Map_Internal_lookup() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup.Do(func() {
		cache_Data_Map_Internal_lookup = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_lookup(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_lookup
}

var cache_Data_Map_Internal_lookup__4110433260 gopurs_runtime.Value
var once_Data_Map_Internal_lookup__4110433260 sync.Once

func Get_Data_Map_Internal_lookup__4110433260() gopurs_runtime.Value {
	once_Data_Map_Internal_lookup__4110433260.Do(func() {
		cache_Data_Map_Internal_lookup__4110433260 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_lookup__4110433260(uint32(__eta0_0_box.IntVal), __eta1_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Map_Internal_lookup__4110433260
}

var cache_Data_Map_Internal_isSubmap gopurs_runtime.Value
var once_Data_Map_Internal_isSubmap sync.Once

func Get_Data_Map_Internal_isSubmap() gopurs_runtime.Value {
	once_Data_Map_Internal_isSubmap.Do(func() {
		cache_Data_Map_Internal_isSubmap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_isSubmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_isSubmap
}

var cache_Data_Map_Internal_member gopurs_runtime.Value
var once_Data_Map_Internal_member sync.Once

func Get_Data_Map_Internal_member() gopurs_runtime.Value {
	once_Data_Map_Internal_member.Do(func() {
		cache_Data_Map_Internal_member = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_member
}

var cache_Data_Map_Internal_pop gopurs_runtime.Value
var once_Data_Map_Internal_pop sync.Once

func Get_Data_Map_Internal_pop() gopurs_runtime.Value {
	once_Data_Map_Internal_pop.Do(func() {
		cache_Data_Map_Internal_pop = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_pop(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_pop
}

var cache_Data_Map_Internal_unionWith gopurs_runtime.Value
var once_Data_Map_Internal_unionWith sync.Once

func Get_Data_Map_Internal_unionWith() gopurs_runtime.Value {
	once_Data_Map_Internal_unionWith.Do(func() {
		cache_Data_Map_Internal_unionWith = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_unionWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_unionWith
}

var cache_Data_Map_Internal_unionWith__2491407586 gopurs_runtime.Value
var once_Data_Map_Internal_unionWith__2491407586 sync.Once

func Get_Data_Map_Internal_unionWith__2491407586() gopurs_runtime.Value {
	once_Data_Map_Internal_unionWith__2491407586.Do(func() {
		cache_Data_Map_Internal_unionWith__2491407586 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_unionWith__2491407586(__eta0_0_box, __eta1_1_box, __eta2_2_box)
		})
	})
	return cache_Data_Map_Internal_unionWith__2491407586
}

var cache_Data_Map_Internal_semigroupMap gopurs_runtime.Value
var once_Data_Map_Internal_semigroupMap sync.Once

func Get_Data_Map_Internal_semigroupMap() gopurs_runtime.Value {
	once_Data_Map_Internal_semigroupMap.Do(func() {
		cache_Data_Map_Internal_semigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_semigroupMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_semigroupMap
}

var cache_Data_Map_Internal_monoidMap gopurs_runtime.Value
var once_Data_Map_Internal_monoidMap sync.Once

func Get_Data_Map_Internal_monoidMap() gopurs_runtime.Value {
	once_Data_Map_Internal_monoidMap.Do(func() {
		cache_Data_Map_Internal_monoidMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_monoidMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_monoidMap
}

var cache_Data_Map_Internal_union gopurs_runtime.Value
var once_Data_Map_Internal_union sync.Once

func Get_Data_Map_Internal_union() gopurs_runtime.Value {
	once_Data_Map_Internal_union.Do(func() {
		cache_Data_Map_Internal_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_union
}

var cache_Data_Map_Internal_checkValid gopurs_runtime.Value
var once_Data_Map_Internal_checkValid sync.Once

func Get_Data_Map_Internal_checkValid() gopurs_runtime.Value {
	once_Data_Map_Internal_checkValid.Do(func() {
		cache_Data_Map_Internal_checkValid = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Map_Internal_checkValid(v_0_box))
		})
	})
	return cache_Data_Map_Internal_checkValid
}

var cache_Data_Map_Internal_catMaybes gopurs_runtime.Value
var once_Data_Map_Internal_catMaybes sync.Once

func Get_Data_Map_Internal_catMaybes() gopurs_runtime.Value {
	once_Data_Map_Internal_catMaybes.Do(func() {
		cache_Data_Map_Internal_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_catMaybes
}

var cache_Data_Map_Internal_applyMap gopurs_runtime.Value
var once_Data_Map_Internal_applyMap sync.Once

func Get_Data_Map_Internal_applyMap() gopurs_runtime.Value {
	once_Data_Map_Internal_applyMap.Do(func() {
		cache_Data_Map_Internal_applyMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_applyMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_applyMap
}

var cache_Data_Map_Internal_bindMap gopurs_runtime.Value
var once_Data_Map_Internal_bindMap sync.Once

func Get_Data_Map_Internal_bindMap() gopurs_runtime.Value {
	once_Data_Map_Internal_bindMap.Do(func() {
		cache_Data_Map_Internal_bindMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_bindMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_bindMap
}

var cache_Data_Map_Internal_anyWithKey gopurs_runtime.Value
var once_Data_Map_Internal_anyWithKey sync.Once

func Get_Data_Map_Internal_anyWithKey() gopurs_runtime.Value {
	once_Data_Map_Internal_anyWithKey.Do(func() {
		cache_Data_Map_Internal_anyWithKey = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_anyWithKey(f_0_box)
		})
	})
	return cache_Data_Map_Internal_anyWithKey
}

var cache_Data_Map_Internal_any gopurs_runtime.Value
var once_Data_Map_Internal_any sync.Once

func Get_Data_Map_Internal_any() gopurs_runtime.Value {
	once_Data_Map_Internal_any.Do(func() {
		cache_Data_Map_Internal_any = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_any(f_0_box)
		})
	})
	return cache_Data_Map_Internal_any
}

var cache_Data_Map_Internal_alter gopurs_runtime.Value
var once_Data_Map_Internal_alter sync.Once

func Get_Data_Map_Internal_alter() gopurs_runtime.Value {
	once_Data_Map_Internal_alter.Do(func() {
		cache_Data_Map_Internal_alter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_alter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_alter
}

var cache_Data_Map_Internal_update gopurs_runtime.Value
var once_Data_Map_Internal_update sync.Once

func Get_Data_Map_Internal_update() gopurs_runtime.Value {
	once_Data_Map_Internal_update.Do(func() {
		cache_Data_Map_Internal_update = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_update(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Map_Internal_update
}

var cache_Data_Map_Internal_altMap gopurs_runtime.Value
var once_Data_Map_Internal_altMap sync.Once

func Get_Data_Map_Internal_altMap() gopurs_runtime.Value {
	once_Data_Map_Internal_altMap.Do(func() {
		cache_Data_Map_Internal_altMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_altMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_altMap
}

var cache_Data_Map_Internal_plusMap gopurs_runtime.Value
var once_Data_Map_Internal_plusMap sync.Once

func Get_Data_Map_Internal_plusMap() gopurs_runtime.Value {
	once_Data_Map_Internal_plusMap.Do(func() {
		cache_Data_Map_Internal_plusMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Map_Internal_plusMap(dictOrd_0_box)
		})
	})
	return cache_Data_Map_Internal_plusMap
}

func Call_Data_Map_Internal_map__870244361(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__870244361:
	for {
		if false {
			continue map__870244361
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		{
			if __eta_norm_0_1 != nil {
				__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3132786365_138441832((__eta_norm_0_1).V0))}), true}
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
}

func Call_Data_Map_Internal_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Data_Map_Internal_values(m_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_valuesImpl(), m_0).UnsafePtr)))))
}

func Call_Data_Map_Internal_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	// TAST (Let): entries_2_0 shape=App(Var) bindingType=(Array (ADT ["Data","Tuple","Tuple"] [(TypeVar k$scope10), (TypeVar v$scope11)]))
	entries_2_0 := func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Map_Internal_toArrayImpl(), Get_Data_Tuple_Tuple(), m_1).UnsafePtr)
		unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
		for i, v := range arr {
			unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v)
		}
		return unboxed
	}()
	_ = entries_2_0
	return gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
		arrayIndex_arg_4 := entries_2_0
		_ = arrayIndex_arg_4
		arrayIndex_arg_5 := i_3
		_ = arrayIndex_arg_5
		// TAST (Let): __local_var_4_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar k$scope10), (TypeVar v$scope11)])])
		__local_var_4_1 := Rebox_Data_Map_Internal_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			arrayIndex_value := arrayIndex_arg_4
			_ = arrayIndex_value
			arrayIndex_source_6 := arrayIndex_value
			_ = arrayIndex_source_6
			if ((arrayIndex_arg_5.IntVal) >= (0)) && ((arrayIndex_arg_5.IntVal) < (int64(len(arrayIndex_source_6)))) {
				return gopurs_runtime.Apply(arrayIndex_arg_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((arrayIndex_source_6)[arrayIndex_arg_5.IntVal])})
			} else {

			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_3)}
		}()))
		_ = __local_var_4_1
		var __t6 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if __local_var_4_1 != nil {
				__t6 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((__local_var_4_1).V0)}, gopurs_runtime.Int(gopurs_runtime.IntAdd(i_3.IntVal, int64(1)))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}(), true}
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
	}), gopurs_runtime.Int(int64(0)))
}

func Call_Data_Map_Internal_toUnfoldable__3206239856(__eta0_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
toUnfoldable__3206239856:
	for {
		if false {
			continue toUnfoldable__3206239856
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		// TAST (Let): entries_1_0 shape=App(Var) bindingType=(Array (ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Interval","Duration","DurationComponent"] []), Number]))
		entries_1_0 := func() []*Constructor_Data_Tuple_Tuple[uint32, float64] {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Map_Internal_toArrayImpl(), Get_Data_Tuple_Tuple(), __eta0_0).UnsafePtr)
			unboxed := make([]*Constructor_Data_Tuple_Tuple[uint32, float64], len(arr))
			for i, v := range arr {
				unboxed[i] = Rebox_Data_Map_Internal_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
			}
			return unboxed
		}()
		_ = entries_1_0
		var Call_local_Data_Map_Internal_go__go_2_1_0 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
		_ = Call_local_Data_Map_Internal_go__go_2_1_0
		var go__go_2_1_0 gopurs_runtime.Value
		_ = go__go_2_1_0
		Call_local_Data_Map_Internal_go__go_2_1_0 = func(source_3_loop gopurs_runtime.Value, memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
		go__go_2_1_0:
			for {
				if false {
					continue go__go_2_1_0
				}
				var source_3 gopurs_runtime.Value = source_3_loop
				_ = source_3
				var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
				_ = memo_4
				arrayIndex_arg_4 := Get_Data_Maybe_Just()
				_ = arrayIndex_arg_4
				arrayIndex_arg_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
				_ = arrayIndex_arg_5
				arrayIndex_arg_6 := entries_1_0
				_ = arrayIndex_arg_6
				arrayIndex_arg_7 := source_3.IntVal
				_ = arrayIndex_arg_7
				// TAST (Let): __local_var_5_3 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Interval","Duration","DurationComponent"] []), Number])])
				__local_var_5_3 := Rebox_Data_Map_Internal_3094389156_2304530460(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					arrayIndex_value := arrayIndex_arg_6
					_ = arrayIndex_value
					arrayIndex_source_8 := arrayIndex_value
					_ = arrayIndex_source_8
					if ((arrayIndex_arg_7) >= (0)) && ((arrayIndex_arg_7) < (int64(len(arrayIndex_source_8)))) {
						return gopurs_runtime.Apply(arrayIndex_arg_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3132786365_138441832((arrayIndex_source_8)[arrayIndex_arg_7]))})
					} else {

					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(arrayIndex_arg_5)}
				}()))
				_ = __local_var_5_3
				var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				{
					if __local_var_5_3 != nil {
						__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3132786365_138441832((__local_var_5_3).V0))}, gopurs_runtime.Int(gopurs_runtime.IntAdd(source_3.IntVal, int64(1)))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}(), true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}())
						goto end_branch_8
					} else {

					}
				}
				{
					__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
			end_branch_8:
				// TAST (Let): v_5_2 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope3), (TypeVar b$scope4)])])
				var v_5_2 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Data_Map_Internal_3094389156_4010058633(__t8)
				var __t11 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				{
					if v_5_2 == nil {
						var Call_local_Data_Map_Internal_go__go_6_9_1 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
						_ = Call_local_Data_Map_Internal_go__go_6_9_1
						var go__go_6_9_1 gopurs_runtime.Value
						_ = go__go_6_9_1
						Call_local_Data_Map_Internal_go__go_6_9_1 = func(b_7_loop gopurs_runtime.Value, v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
						go__go_6_9_1:
							for {
								if false {
									continue go__go_6_9_1
								}
								var b_7 gopurs_runtime.Value = b_7_loop
								_ = b_7
								var v_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_8_loop
								_ = v_8
								var __t10 gopurs_runtime.Value
								{
									if v_8 == nil {
										__t10 = b_7
										goto end_branch_10
									} else {

									}
								}
								{
									if v_8 != nil {
										b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_8).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_7)}))}
										v_8_loop = (v_8).V1
										continue go__go_6_9_1
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
						go__go_6_9_1 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return Call_local_Data_Map_Internal_go__go_6_9_1(b_7_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val))
							})
						})
						__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_Map_Internal_go__go_6_9_1(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, memo_4))
						goto end_branch_11
					} else {

					}
				}
				{
					if v_5_2 != nil {
						source_3_loop = ((v_5_2).V0).V1
						memo_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v_5_2).V0).V0, memo_4})
						continue go__go_2_1_0
						__t11 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
						goto end_branch_11
					} else {

					}
				}
				{
					__t11 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
				}
			end_branch_11:
				return __t11
			}
		}
		go__go_2_1_0 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_Map_Internal_go__go_2_1_0(source_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)))}
			})
		})
		return Rebox_Data_Map_Internal_849153993_2442833393(Call_local_Data_Map_Internal_go__go_2_1_0(gopurs_runtime.Int(int64(0)), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))
	}
}

func Call_Data_Map_Internal_toUnfoldableUnordered(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply(Get_Data_Map_Internal_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)})
}

func Call_Data_Map_Internal_size__2179923709(__eta_norm_0_0_loop gopurs_runtime.Value) int64 {
size__2179923709:
	for {
		if false {
			continue size__2179923709
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(Get_Data_Map_Internal_sizeImpl(), __eta_norm_0_0).IntVal
	}
}

func Call_Data_Map_Internal_showMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(fromFoldable ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(Call_Data_Tuple_showTuple(dictShow_0, dictShow1_1)), "show"), func() gopurs_runtime.Value {
			arr := func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(Call_Data_Map_Internal_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()), m_2).UnsafePtr)
				unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v)
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
			}
			return gopurs_runtime.Array(boxed)
		}()).StrVal()), ")"))
	})}))}
}

func Call_Data_Map_Internal_lookupLT(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var k_1 gopurs_runtime.Value = k_1_loop
	_ = k_1
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_2 gopurs_runtime.Value, k_prime__3 gopurs_runtime.Value, v_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_prime__3, k_1)
			_ = __t_tag_0
			if uint32(__t_tag_0.IntVal) == 1527465420 {
				__t1 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.RecordDict2("key", "value", k_prime__3, v_prime__4), true}
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
			__t1 = acc_2
		}
	end_branch_1:
		return __t1
	}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 struct {
				key   gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
			}
			V1 bool
		}{struct {
			key   gopurs_runtime.Value
			value gopurs_runtime.Value
			/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
		}{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))))})
}

func Call_Data_Map_Internal_lookupLE(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var k_1 gopurs_runtime.Value = k_1_loop
	_ = k_1
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_2 gopurs_runtime.Value, k_prime__3 gopurs_runtime.Value, v_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_prime__3, k_1)
			_ = __t_tag_0
			if (uint32(__t_tag_0.IntVal) == 380165415) != (true) {
				__t1 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.RecordDict2("key", "value", k_prime__3, v_prime__4), true}
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
			__t1 = acc_2
		}
	end_branch_1:
		return __t1
	}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 struct {
				key   gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
			}
			V1 bool
		}{struct {
			key   gopurs_runtime.Value
			value gopurs_runtime.Value
			/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
		}{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))))})
}

func Call_Data_Map_Internal_lookupGT(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var k_1 gopurs_runtime.Value = k_1_loop
	_ = k_1
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldrImpl(), gopurs_runtime.Func3(func(acc_2 gopurs_runtime.Value, k_prime__3 gopurs_runtime.Value, v_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_prime__3, k_1)
			_ = __t_tag_0
			if uint32(__t_tag_0.IntVal) == 380165415 {
				__t1 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.RecordDict2("key", "value", k_prime__3, v_prime__4), true}
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
			__t1 = acc_2
		}
	end_branch_1:
		return __t1
	}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 struct {
				key   gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
			}
			V1 bool
		}{struct {
			key   gopurs_runtime.Value
			value gopurs_runtime.Value
			/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
		}{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))))})
}

func Call_Data_Map_Internal_lookupGE(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var k_1 gopurs_runtime.Value = k_1_loop
	_ = k_1
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldrImpl(), gopurs_runtime.Func3(func(acc_2 gopurs_runtime.Value, k_prime__3 gopurs_runtime.Value, v_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_prime__3, k_1)
			_ = __t_tag_0
			if (uint32(__t_tag_0.IntVal) == 1527465420) != (true) {
				__t1 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.RecordDict2("key", "value", k_prime__3, v_prime__4), true}
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
			__t1 = acc_2
		}
	end_branch_1:
		return __t1
	}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_2758605161_3094389156(Rebox_Data_Map_Internal_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 struct {
				key   gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
			}
			V1 bool
		}{struct {
			key   gopurs_runtime.Value
			value gopurs_runtime.Value
			/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
		}{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
				orig := _v.V0
				_ = orig
				return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
			}()})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))))})
}

func Call_Data_Map_Internal_keys(m_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_keysImpl(), m_0).UnsafePtr)))))
}

func Call_Data_Map_Internal_eqMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((gopurs_runtime.Apply(Get_Data_Map_Internal_sizeImpl(), m1_2).IntVal) == (gopurs_runtime.Apply(Get_Data_Map_Internal_sizeImpl(), m2_3).IntVal)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Types_eqList(dictEq_0), "eq"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_keysImpl(), m1_2).UnsafePtr)))), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_keysImpl(), m2_3).UnsafePtr))))).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Types_eqList(dictEq1_1), "eq"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_valuesImpl(), m1_2).UnsafePtr)))), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Map_Internal_valuesImpl(), m2_3).UnsafePtr))))).IntVal) != (0))))
	})}))}
}

func Call_Data_Map_Internal_ordMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): ordTuple_1_0 shape=App(Var) bindingType=Any
	ordTuple_1_0 := Call_Data_Tuple_ordTuple(dictOrd_0)
	_ = ordTuple_1_0
	// TAST (Let): eqMap1_2_1 shape=App(Var) bindingType=Any
	eqMap1_2_1 := gopurs_runtime.Apply(Get_Data_Map_Internal_eqMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqMap1_2_1
	return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eqMap2_4_2 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope45), (TypeVar v$scope46)])])
		eqMap2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqMap1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})))
		_ = eqMap2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMap2_4_2)}
		}), gopurs_runtime.Func2(func(m1_5 gopurs_runtime.Value, m2_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_List_Types_ordList(gopurs_runtime.Apply(ordTuple_1_0, dictOrd1_3)), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3266529732_849153993(Rebox_Data_Map_Internal_849153993_3266529732(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_Data_Map_Internal_toUnfoldable(Rebox_Data_Map_Internal_132451362_2738507278(Rebox_Data_Map_Internal_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))), m1_5)))))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3266529732_849153993(Rebox_Data_Map_Internal_849153993_3266529732(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_Data_Map_Internal_toUnfoldable(Rebox_Data_Map_Internal_132451362_2738507278(Rebox_Data_Map_Internal_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))), m2_6)))))})
		})}))}
	})
}

func Call_Data_Map_Internal_compareInt(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) int64 {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
	v_3_0 := uint32(gopurs_runtime.Apply2(dictOrd_0.V1, a_1, b_2).IntVal)
	_ = v_3_0
	var __t1 int64
	{
		if v_3_0 == 1527465420 {
			__t1 = int64(-1)
			goto end_branch_1
		} else {

		}
	}
	{
		if v_3_0 == 902936544 {
			__t1 = int64(0)
			goto end_branch_1
		} else {

		}
	}
	{
		if v_3_0 == 380165415 {
			__t1 = int64(1)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Data_Map_Internal_go__delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Map_Internal_deleteImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_Map_Internal_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Map_Internal_differenceImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_Map_Internal_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_Map_Internal_filterWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope207), (TypeVar v$scope208), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope207), (TypeVar v$scope208)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope207), (TypeVar v$scope208)]))
	insert1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insert1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if (gopurs_runtime.Apply2(f_2, k_4, v_5).IntVal) != (0) {
					__t1 = gopurs_runtime.Apply3(insert1_1_0, k_4, v_5, acc_3)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = acc_3
			}
		end_branch_1:
			return __t1
		}), Get_Data_Map_Internal_empty())
	})
}

func Call_Data_Map_Internal_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): filterWithKey1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar k$scope216), (TypeVar v$scope217)] Boolean), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope216), (TypeVar v$scope217)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope216), (TypeVar v$scope217)]))
	filterWithKey1_1_0 := Call_Data_Map_Internal_filterWithKey(dictOrd_0)
	_ = filterWithKey1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(filterWithKey1_1_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return f_2
		}))
	})
}

func Call_Data_Map_Internal_submap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): filterWithKey1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar k$scope220), (TypeVar v$scope221)] Boolean), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope220), (TypeVar v$scope221)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope220), (TypeVar v$scope221)]))
	filterWithKey1_1_0 := Call_Data_Map_Internal_filterWithKey(dictOrd_0)
	_ = filterWithKey1_1_0
	return gopurs_runtime.Func2(func(k1_2 gopurs_runtime.Value, k2_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(filterWithKey1_1_0, gopurs_runtime.Func2(func(k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 bool
			{
				if k1_2.Type == 9 && k1_2.IntVal == 930809136 && k1_2.UnsafePtr == nil {
					__t2 = true
					goto end_branch_2
				} else {

				}
			}
			{
				if k1_2.Type == 9 && k1_2.IntVal == 930809136 && k1_2.UnsafePtr != nil {
					var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_4, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(k1_2.UnsafePtr).V0)
					_ = __t_tag_1
					__t2 = (uint32(__t_tag_1.IntVal) == 1527465420) != (true)
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
			}
		end_branch_2:
			var __t_and_5 bool = false
			if __t2 {

				var __t4 bool
				{
					if k2_3.Type == 9 && k2_3.IntVal == 930809136 && k2_3.UnsafePtr == nil {
						__t4 = true
						goto end_branch_4
					} else {

					}
				}
				{
					if k2_3.Type == 9 && k2_3.IntVal == 930809136 && k2_3.UnsafePtr != nil {
						var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, k_4, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(k2_3.UnsafePtr).V0)
						_ = __t_tag_3
						__t4 = (uint32(__t_tag_3.IntVal) == 380165415) != (true)
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() bool { panic("Failed pattern match") }()
				}
			end_branch_4:
				__t_and_5 = __t4
			}
			return gopurs_runtime.Bool(__t_and_5)
		}))
	})
}

func Call_Data_Map_Internal_fromFoldable(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope226), (TypeVar v$scope227), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope226), (TypeVar v$scope227)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope226), (TypeVar v$scope227)]))
	insert1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insert1_1_0
	return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), gopurs_runtime.Func2(func(m_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(insert1_1_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, m_3)
		}), Get_Data_Map_Internal_empty())
	})
}

func Call_Data_Map_Internal_fromFoldableWithIndex(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope234), (TypeVar v$scope235), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope234), (TypeVar v$scope235)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope234), (TypeVar v$scope235)]))
	insert1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insert1_1_0
	return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_2, "foldlWithIndex"), gopurs_runtime.Func3(func(k_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(insert1_1_0, k_3, v_5, m_4)
		}), Get_Data_Map_Internal_empty())
	})
}

func Call_Data_Map_Internal_functorWithIndexMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope239), (TypeVar b$scope245), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope239), (TypeVar b$scope245)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope239), (TypeVar b$scope245)]))
	insert1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), dictOrd_0))
	_ = insert1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Map_Internal_functorMap()
	}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(insert1_1_0, k_4, gopurs_runtime.Apply2(f_2, k_4, v_5), acc_3)
		}), Get_Data_Map_Internal_empty())
	})}))}
}

func Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope254), (TypeVar b$scope256), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope254), (TypeVar b$scope256)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope254), (TypeVar b$scope256)]))
	insert1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insert1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v1_6_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope256)])
			v1_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_2, k_4, v_5))
			_ = v1_6_1
			var __t2 gopurs_runtime.Value
			{
				if v1_6_1 != nil {
					__t2 = gopurs_runtime.Apply3(insert1_1_0, k_4, (v1_6_1).V0, acc_3)
					goto end_branch_2
				} else {

				}
			}
			{
				if v1_6_1 == nil {
					__t2 = acc_3
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}), Get_Data_Map_Internal_empty())
	})
}

func Call_Data_Map_Internal_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): mapMaybeWithKey1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar k$scope265), (TypeVar a$scope266)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope267)])), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope265), (TypeVar a$scope266)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope265), (TypeVar b$scope267)]))
	mapMaybeWithKey1_1_0 := Call_Data_Map_Internal_mapMaybeWithKey(dictOrd_0)
	_ = mapMaybeWithKey1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(mapMaybeWithKey1_1_0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return f_2
		}))
	})
}

func Call_Data_Map_Internal_insertWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Map_Internal_insertWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_Map_Internal_fromFoldableWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insertWith1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar v$scope279), (TypeVar v$scope279)] (TypeVar v$scope279)), (TypeVar k$scope278), (TypeVar v$scope279), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope278), (TypeVar v$scope279)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope278), (TypeVar v$scope279)]))
	insertWith1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insertWith1_1_0
	return gopurs_runtime.Func2(func(dictFoldable_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): f_prime__4_1 shape=App(Other) bindingType=(Func [(TypeVar k$scope278), (TypeVar v$scope279), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope278), (TypeVar v$scope279)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope278), (TypeVar v$scope279)]))
		f_prime__4_1 := gopurs_runtime.Apply(insertWith1_1_0, gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_3, a_5, b_4)
		}))
		_ = f_prime__4_1
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(f_prime__4_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1, m_5)
		}), Get_Data_Map_Internal_empty())
	})
}

func Call_Data_Map_Internal_intersectionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Map_Internal_intersectionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_Map_Internal_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_intersectionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}), Get_Data_Function_go__const())
}

func Call_Data_Map_Internal_lookup(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply3(Get_Data_Map_Internal_lookupImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_Map_Internal_lookup__4110433260(__eta0_0_loop uint32, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
lookup__4110433260:
	for {
		if false {
			continue lookup__4110433260
		}
		var __eta0_0 uint32 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3240988860_3094389156(Rebox_Data_Map_Internal_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply5(Get_Data_Map_Internal_lookupImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3240988860_3094389156(Rebox_Data_Map_Internal_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))}, gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3730953251_4177771502(Rebox_Data_Map_Internal_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Interval_Duration_ordDurationComponent()))))}), gopurs_runtime.Value{Type: 9, IntVal: int64(__eta0_0), UnsafePtr: nil}, __eta1_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Map_Internal_isSubmap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): lookup1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope310), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope310), (TypeVar v$scope311)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope311)]))
	lookup1_1_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_lookupImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = lookup1_1_0
	return gopurs_runtime.Func3(func(dictEq_2 gopurs_runtime.Value, m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(((acc_5.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_eqMaybe(dictEq_2), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(lookup1_1_0, k_6, m2_4)))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{v_7, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}).IntVal) != (0)))
		}), gopurs_runtime.Bool(true), m1_3)
	})
}

func Call_Data_Map_Internal_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): lookup1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope318), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope318), (TypeVar v$scope319)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope319)]))
	lookup1_1_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_lookupImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = lookup1_1_0
	return gopurs_runtime.Func2(func(k_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope319)])
		v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(lookup1_1_0, k_2, m_3))
		_ = v_4_1
		var __t2 gopurs_runtime.Value
		{
			if v_4_1 == nil {
				__t2 = gopurs_runtime.Bool(false)
				goto end_branch_2
			} else {

			}
		}
		{
			if v_4_1 != nil {
				__t2 = gopurs_runtime.Bool(true)
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.Bool(func() bool { panic("Failed pattern match") }())
		}
	end_branch_2:
		return __t2
	})
}

func Call_Data_Map_Internal_pop(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): lookup1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope322), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope322), (TypeVar v$scope323)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope323)]))
	lookup1_1_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_lookupImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = lookup1_1_0
	// TAST (Let): delete1_2_1 shape=App(Var) bindingType=(Func [(TypeVar k$scope322), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope322), (TypeVar v$scope323)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope322), (TypeVar v$scope323)]))
	delete1_2_1 := gopurs_runtime.Apply(Get_Data_Map_Internal_deleteImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = delete1_2_1
	return gopurs_runtime.Func2(func(k_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v_5_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope323)])
		v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(lookup1_1_0, k_3, m_4))
		_ = v_5_2
		var __t3 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if v_5_2 == nil {
				__t3 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				goto end_branch_3
			} else {

			}
		}
		{
			if v_5_2 != nil {
				__t3 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_5_2).V0, gopurs_runtime.Apply2(delete1_2_1, k_3, m_4)}))}, true}
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() struct {
				V0 gopurs_runtime.Value
				V1 bool
			} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_4010058633_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
				panic("Failed pattern match")
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
	end_branch_3:
		return func() gopurs_runtime.Value {
			_v := __t3
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	})
}

func Call_Data_Map_Internal_unionWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply(Get_Data_Map_Internal_unionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_Map_Internal_unionWith__2491407586(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unionWith__2491407586:
	for {
		if false {
			continue unionWith__2491407586
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
		_ = __eta2_2
		return gopurs_runtime.Apply4(Get_Data_Map_Internal_unionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3730953251_4177771502(Rebox_Data_Map_Internal_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Interval_Duration_ordDurationComponent()))))}), __eta0_0, __eta1_1, __eta2_2)
	}
}

func Call_Data_Map_Internal_semigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): unionWith1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar v$scope337), (TypeVar v$scope337)] (TypeVar v$scope337)), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope336), (TypeVar v$scope337)]), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope336), (TypeVar v$scope337)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope336), (TypeVar v$scope337)]))
	unionWith1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_unionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), dictOrd_0))
	_ = unionWith1_1_0
	return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Apply(unionWith1_1_0, Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_2)))}))}
	})
}

func Call_Data_Map_Internal_monoidMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): semigroupMap1_1_0 shape=App(Var) bindingType=Any
	semigroupMap1_1_0 := Call_Data_Map_Internal_semigroupMap(dictOrd_0)
	_ = semigroupMap1_1_0
	return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupMap2_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope50), (TypeVar v$scope51)])])
		semigroupMap2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupMap1_1_0, dictSemigroup_2))
		_ = semigroupMap2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMap2_3_1)}
		}), Get_Data_Map_Internal_empty()}))}
	})
}

func Call_Data_Map_Internal_union(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_unionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}), Get_Data_Function_go__const())
}

func Call_Data_Map_Internal_checkValid(v_0_loop gopurs_runtime.Value) bool {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return true
}

func Call_Data_Map_Internal_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): insert1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope349), (TypeVar v$scope350), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope349), (TypeVar v$scope350)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope349), (TypeVar v$scope350)]))
	insert1_1_0 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insert1_1_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil {
				__t1 = gopurs_runtime.Apply3(insert1_1_0, k_3, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0, acc_2)
				goto end_branch_1
			} else {

			}
		}
		{
			if v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil {
				__t1 = acc_2
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}), Get_Data_Map_Internal_empty())
}

func Call_Data_Map_Internal_applyMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Map_Internal_functorMap()
	}), gopurs_runtime.Apply2(Get_Data_Map_Internal_intersectionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), dictOrd_0), gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f_1, x_2)
	}))}))}
}

func Call_Data_Map_Internal_bindMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): mapMaybeWithKey1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar k$scope357), (TypeVar a$scope361)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope362)])), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope357), (TypeVar a$scope361)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope357), (TypeVar b$scope362)]))
	mapMaybeWithKey1_1_0 := Call_Data_Map_Internal_mapMaybeWithKey(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))
	_ = mapMaybeWithKey1_1_0
	// TAST (Let): lookup1_2_1 shape=App(Var) bindingType=(Func [(TypeVar k$scope357), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope357), (TypeVar b$scope362)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope362)]))
	lookup1_2_1 := gopurs_runtime.Apply3(Get_Data_Map_Internal_lookupImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), dictOrd_0))
	_ = lookup1_2_1
	// TAST (Let): applyMap1_3_2 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope357)])])
	applyMap1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Map_Internal_applyMap(dictOrd_0))
	_ = applyMap1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyMap1_3_2)}
	}), gopurs_runtime.Func2(func(m_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(mapMaybeWithKey1_1_0, gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(lookup1_2_1, k_6), f_5)
		}), m_4)
	})}))}
}

func Call_Data_Map_Internal_anyWithKey(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((acc_1.IntVal) != (0)) || ((gopurs_runtime.Apply2(f_0, k_2, v_3).IntVal) != (0)))
	}), gopurs_runtime.Bool(false))
}

func Call_Data_Map_Internal_any(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_foldlImpl(), gopurs_runtime.Func3(func(acc_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((acc_1.IntVal) != (0)) || ((gopurs_runtime.Apply(f_0, v_3).IntVal) != (0)))
	}), gopurs_runtime.Bool(false))
}

func Call_Data_Map_Internal_alter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): lookup1_1_0 shape=App(Var) bindingType=(Func [(TypeVar k$scope384), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope384), (TypeVar v$scope385)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope385)]))
	lookup1_1_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_lookupImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))}, gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = lookup1_1_0
	// TAST (Let): delete1_2_1 shape=App(Var) bindingType=(Func [(TypeVar k$scope384), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope384), (TypeVar v$scope385)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope384), (TypeVar v$scope385)]))
	delete1_2_1 := gopurs_runtime.Apply(Get_Data_Map_Internal_deleteImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = delete1_2_1
	// TAST (Let): insert1_3_2 shape=App(Var) bindingType=(Func [(TypeVar k$scope384), (TypeVar v$scope385), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope384), (TypeVar v$scope385)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope384), (TypeVar v$scope385)]))
	insert1_3_2 := gopurs_runtime.Apply(Get_Data_Map_Internal_insertImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
	_ = insert1_3_2
	return gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, m_6 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v_7_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope385)])
		v_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(lookup1_1_0, k_5, m_6)))}))
		_ = v_7_3
		var __t4 gopurs_runtime.Value
		{
			if v_7_3 == nil {
				__t4 = gopurs_runtime.Apply2(delete1_2_1, k_5, m_6)
				goto end_branch_4
			} else {

			}
		}
		{
			if v_7_3 != nil {
				__t4 = gopurs_runtime.Apply3(insert1_3_2, k_5, (v_7_3).V0, m_6)
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
}

func Call_Data_Map_Internal_update(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): alter1_1_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope389)])] (ADT ["Data","Maybe","Maybe"] [(TypeVar v$scope389)])), (TypeVar k$scope388), (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope388), (TypeVar v$scope389)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope388), (TypeVar v$scope389)]))
	alter1_1_0 := Call_Data_Map_Internal_alter(dictOrd_0)
	_ = alter1_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(alter1_1_0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr == nil {
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
				if v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr != nil {
					__t1 = gopurs_runtime.Apply(f_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_1:
			return __t1
		}), k_3, m_4)
	})
}

func Call_Data_Map_Internal_altMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Map_Internal_functorMap()
	}), Call_Data_Map_Internal_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}))}
}

func Call_Data_Map_Internal_plusMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): altMap1_1_0 shape=App(Var) bindingType=(ADT ["Control","Alt","Alt"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar k$scope37)])])
	altMap1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Call_Data_Map_Internal_altMap(dictOrd_0))
	_ = altMap1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altMap1_1_0)}
	}), Get_Data_Map_Internal_empty()}))}
}

func Rebox_Data_Map_Internal_132451362_2738507278(in *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Internal_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
	out.V0 = uint32(in.V0.IntVal)
	out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Map_Internal_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Map_Internal_3132786365_138441832(in.V0))}
	out.V1 = Rebox_Data_Map_Internal_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Map_Internal_2738507278_132451362(in *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Internal_2758605161_3094389156(in *Constructor_Data_Maybe_Just[struct {
	key   gopurs_runtime.Value
	value gopurs_runtime.Value
	/*gopurs:record-v1:2:006b00650079,00760061006c00750065*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("key", "value", orig.key, orig.value)
	}()
	return out
}

func Rebox_Data_Map_Internal_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Internal_3094389156_2304530460(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
	out.V0 = Rebox_Data_Map_Internal_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Map_Internal_3094389156_2758605161(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
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

func Rebox_Data_Map_Internal_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[float64]{}
	out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Data_Map_Internal_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Map_Internal_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Map_Internal_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Data_Map_Internal_3266529732_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = Rebox_Data_Map_Internal_3266529732_849153993(in.V1)
	return out
}

func Rebox_Data_Map_Internal_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Internal_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Map_Internal_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Map_Internal_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
	out.V0 = Rebox_Data_Map_Internal_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = Rebox_Data_Map_Internal_849153993_2442833393(in.V1)
	return out
}

func Rebox_Data_Map_Internal_849153993_3266529732(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	out.V1 = Rebox_Data_Map_Internal_849153993_3266529732(in.V1)
	return out
}

func Get_Data_Map_Internal_deleteImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_DeleteImpl
}

func Get_Data_Map_Internal_differenceImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_DifferenceImpl
}

func Get_Data_Map_Internal_empty() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_Empty
}

func Get_Data_Map_Internal_filterKeysImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_FilterKeysImpl
}

func Get_Data_Map_Internal_foldlImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_FoldlImpl
}

func Get_Data_Map_Internal_foldrImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_FoldrImpl
}

func Get_Data_Map_Internal_insertImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_InsertImpl
}

func Get_Data_Map_Internal_insertNativeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_InsertNativeImpl
}

func Get_Data_Map_Internal_insertWithImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_InsertWithImpl
}

func Get_Data_Map_Internal_intersectionWithImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_IntersectionWithImpl
}

func Get_Data_Map_Internal_isEmpty() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_IsEmpty
}

func Get_Data_Map_Internal_keysImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_KeysImpl
}

func Get_Data_Map_Internal_lookupImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_LookupImpl
}

func Get_Data_Map_Internal_lookupNativeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_LookupNativeImpl
}

func Get_Data_Map_Internal_mapImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_MapImpl
}

func Get_Data_Map_Internal_singleton() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_Singleton
}

func Get_Data_Map_Internal_sizeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_SizeImpl
}

func Get_Data_Map_Internal_toArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_ToArrayImpl
}

func Get_Data_Map_Internal_unionWithImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_UnionWithImpl
}

func Get_Data_Map_Internal_valuesImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Map_Internal_ValuesImpl
}
