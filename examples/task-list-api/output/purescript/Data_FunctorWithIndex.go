package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_FunctorWithIndex_go__map gopurs_runtime.Value
var once_Data_FunctorWithIndex_go__map sync.Once

func Get_Data_FunctorWithIndex_go__map() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_go__map.Do(func() {
		cache_Data_FunctorWithIndex_go__map = gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map")
	})
	return cache_Data_FunctorWithIndex_go__map
}

var cache_Data_FunctorWithIndex_bimap gopurs_runtime.Value
var once_Data_FunctorWithIndex_bimap sync.Once

func Get_Data_FunctorWithIndex_bimap() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_bimap.Do(func() {
		cache_Data_FunctorWithIndex_bimap = Rebox_Data_FunctorWithIndex_1688994542_1495429347(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorTuple())).V0
	})
	return cache_Data_FunctorWithIndex_bimap
}

var cache_Data_FunctorWithIndex_compose gopurs_runtime.Value
var once_Data_FunctorWithIndex_compose sync.Once

func Get_Data_FunctorWithIndex_compose() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_compose.Do(func() {
		cache_Data_FunctorWithIndex_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FunctorWithIndex_compose
}

var cache_Data_FunctorWithIndex_compose1 gopurs_runtime.Value
var once_Data_FunctorWithIndex_compose1 sync.Once

func Get_Data_FunctorWithIndex_compose1() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_compose1.Do(func() {
		cache_Data_FunctorWithIndex_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FunctorWithIndex_compose1
}

var cache_Data_FunctorWithIndex_map1 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map1 sync.Once

func Get_Data_FunctorWithIndex_map1() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map1.Do(func() {
		cache_Data_FunctorWithIndex_map1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Multiplicative_functorMultiplicative()).V0
	})
	return cache_Data_FunctorWithIndex_map1
}

var cache_Data_FunctorWithIndex_map2 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map2 sync.Once

func Get_Data_FunctorWithIndex_map2() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map2.Do(func() {
		cache_Data_FunctorWithIndex_map2 = Rebox_Data_FunctorWithIndex_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_FunctorWithIndex_map2
}

var cache_Data_FunctorWithIndex_map3 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map3 sync.Once

func Get_Data_FunctorWithIndex_map3() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map3.Do(func() {
		cache_Data_FunctorWithIndex_map3 = Rebox_Data_FunctorWithIndex_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_FunctorWithIndex_map3
}

var cache_Data_FunctorWithIndex_map4 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map4 sync.Once

func Get_Data_FunctorWithIndex_map4() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map4.Do(func() {
		cache_Data_FunctorWithIndex_map4 = Rebox_Data_FunctorWithIndex_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_FunctorWithIndex_map4
}

var cache_Data_FunctorWithIndex_map5 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map5 sync.Once

func Get_Data_FunctorWithIndex_map5() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map5.Do(func() {
		cache_Data_FunctorWithIndex_map5 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_FunctorWithIndex_map5
}

var cache_Data_FunctorWithIndex_map6 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map6 sync.Once

func Get_Data_FunctorWithIndex_map6() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map6.Do(func() {
		cache_Data_FunctorWithIndex_map6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Dual_functorDual()).V0
	})
	return cache_Data_FunctorWithIndex_map6
}

var cache_Data_FunctorWithIndex_map7 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map7 sync.Once

func Get_Data_FunctorWithIndex_map7() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map7.Do(func() {
		cache_Data_FunctorWithIndex_map7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Disj_functorDisj()).V0
	})
	return cache_Data_FunctorWithIndex_map7
}

var cache_Data_FunctorWithIndex_bimap1 gopurs_runtime.Value
var once_Data_FunctorWithIndex_bimap1 sync.Once

func Get_Data_FunctorWithIndex_bimap1() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_bimap1.Do(func() {
		cache_Data_FunctorWithIndex_bimap1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorEither()).V0
	})
	return cache_Data_FunctorWithIndex_bimap1
}

var cache_Data_FunctorWithIndex_compose2 gopurs_runtime.Value
var once_Data_FunctorWithIndex_compose2 sync.Once

func Get_Data_FunctorWithIndex_compose2() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_compose2.Do(func() {
		cache_Data_FunctorWithIndex_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FunctorWithIndex_compose2
}

var cache_Data_FunctorWithIndex_compose3 gopurs_runtime.Value
var once_Data_FunctorWithIndex_compose3 sync.Once

func Get_Data_FunctorWithIndex_compose3() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_compose3.Do(func() {
		cache_Data_FunctorWithIndex_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FunctorWithIndex_compose3
}

var cache_Data_FunctorWithIndex_map8 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map8 sync.Once

func Get_Data_FunctorWithIndex_map8() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map8.Do(func() {
		cache_Data_FunctorWithIndex_map8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Conj_functorConj()).V0
	})
	return cache_Data_FunctorWithIndex_map8
}

var cache_Data_FunctorWithIndex_compose4 gopurs_runtime.Value
var once_Data_FunctorWithIndex_compose4 sync.Once

func Get_Data_FunctorWithIndex_compose4() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_compose4.Do(func() {
		cache_Data_FunctorWithIndex_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_FunctorWithIndex_compose4
}

var cache_Data_FunctorWithIndex_map9 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map9 sync.Once

func Get_Data_FunctorWithIndex_map9() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map9.Do(func() {
		cache_Data_FunctorWithIndex_map9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Additive_functorAdditive()).V0
	})
	return cache_Data_FunctorWithIndex_map9
}

var cache_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict gopurs_runtime.Value
var once_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict sync.Once

func Get_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict.Do(func() {
		cache_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Call_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict(func() struct {
				Functor0     gopurs_runtime.Value
				mapWithIndex gopurs_runtime.Value
				/*gopurs:record-v1:2:00460075006e00630074006f00720030,006d0061007000570069007400680049006e006400650078*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Functor0     gopurs_runtime.Value
					mapWithIndex gopurs_runtime.Value
					/*gopurs:record-v1:2:00460075006e00630074006f00720030,006d0061007000570069007400680049006e006400650078*/
				}{}
				clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
				clone.mapWithIndex = gopurs_runtime.RecordGet(orig, "mapWithIndex")
				return clone
			}()))}
		})
	})
	return cache_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict
}

var cache_Data_FunctorWithIndex_mapWithIndex gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex sync.Once

func Get_Data_FunctorWithIndex_mapWithIndex() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FunctorWithIndex_mapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex
}

var cache_Data_FunctorWithIndex_mapDefault gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapDefault sync.Once

func Get_Data_FunctorWithIndex_mapDefault() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapDefault.Do(func() {
		cache_Data_FunctorWithIndex_mapDefault = gopurs_runtime.Func2(func(dictFunctorWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FunctorWithIndex_mapDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFunctorWithIndex_0_box), f_1_box)
		})
	})
	return cache_Data_FunctorWithIndex_mapDefault
}

var cache_Data_FunctorWithIndex_functorWithIndexTuple gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexTuple sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexTuple() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexTuple.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexTuple = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_935379557_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Tuple_functorTuple()
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope18)] (TypeVar b$scope19))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			})
		})})))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexTuple
}

var cache_Data_FunctorWithIndex_functorWithIndexProduct gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexProduct sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexProduct() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexProduct.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexProduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FunctorWithIndex_functorWithIndexProduct(dictFunctorWithIndex_0_box)
		})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexProduct
}

var cache_Data_FunctorWithIndex_functorWithIndexMultiplicative gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexMultiplicative sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexMultiplicative() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexMultiplicative.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Multiplicative_functorMultiplicative()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope48)] (TypeVar b$scope49))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_1_0, m_2)
			})
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexMultiplicative
}

var cache_Data_FunctorWithIndex_functorWithIndexMaybe gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexMaybe sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexMaybe() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexMaybe.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_3136246921_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_3689823567_2812149806(Rebox_Data_FunctorWithIndex_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope54)] (TypeVar b$scope55))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 struct {
					V0 gopurs_runtime.Value
					V1 bool
				}
				{
					if v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil {
						__t1 = struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
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
			})
		})})))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexMaybe
}

var cache_Data_FunctorWithIndex_functorWithIndexLast gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexLast sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexLast() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexLast.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexLast = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_3136246921_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_3689823567_2812149806(Rebox_Data_FunctorWithIndex_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope60)] (TypeVar b$scope61))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 struct {
					V0 gopurs_runtime.Value
					V1 bool
				}
				{
					if v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil {
						__t1 = struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
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
			})
		})})))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexLast
}

var cache_Data_FunctorWithIndex_functorWithIndexIdentity gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexIdentity sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexIdentity() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexIdentity.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexIdentity = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_0, Get_Data_Unit_unit(), v_1)
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexIdentity
}

var cache_Data_FunctorWithIndex_functorWithIndexFirst gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexFirst sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexFirst() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexFirst.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexFirst = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_3136246921_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_3689823567_2812149806(Rebox_Data_FunctorWithIndex_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope73)] (TypeVar b$scope74))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 struct {
					V0 gopurs_runtime.Value
					V1 bool
				}
				{
					if v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil {
						__t1 = struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
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
			})
		})})))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexFirst
}

var cache_Data_FunctorWithIndex_functorWithIndexEither gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexEither sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexEither() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexEither.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexEither = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_functorEither()
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope81)] (TypeVar b$scope82))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}
				{
					if m_2.Type == 9 && m_2.IntVal == 3711209382 {
						__t1 = struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Value{}, false}
						goto end_branch_1
					} else {

					}
				}
				{
					if m_2.Type == 9 && m_2.IntVal == 2465973597 {
						__t1 = struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), true}
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
						return struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
					}; return struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
				}
			end_branch_1:
				return func() gopurs_runtime.Value {
					_v := __t1
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
			})
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexEither
}

var cache_Data_FunctorWithIndex_functorWithIndexDual gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexDual sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexDual() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexDual.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexDual = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Dual_functorDual()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope87)] (TypeVar b$scope88))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_1_0, m_2)
			})
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexDual
}

var cache_Data_FunctorWithIndex_functorWithIndexDisj gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexDisj sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexDisj() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexDisj.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexDisj = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Disj_functorDisj()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope93)] (TypeVar b$scope94))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_1_0, m_2)
			})
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexDisj
}

var cache_Data_FunctorWithIndex_functorWithIndexCoproduct gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexCoproduct sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexCoproduct() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexCoproduct.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexCoproduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FunctorWithIndex_functorWithIndexCoproduct(dictFunctorWithIndex_0_box)
		})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexCoproduct
}

var cache_Data_FunctorWithIndex_functorWithIndexConst gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexConst sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexConst() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexConst.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexConst = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Const_functorConst()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return v1_1
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexConst
}

var cache_Data_FunctorWithIndex_functorWithIndexConj gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexConj sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexConj() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexConj.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexConj = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Conj_functorConj()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope134)] (TypeVar b$scope135))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_1_0, m_2)
			})
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexConj
}

var cache_Data_FunctorWithIndex_functorWithIndexCompose gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexCompose sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexCompose() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexCompose.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexCompose = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FunctorWithIndex_functorWithIndexCompose(dictFunctorWithIndex_0_box)
		})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexCompose
}

var cache_Data_FunctorWithIndex_functorWithIndexArray gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexArray sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexArray() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexArray.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexArray = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_2773701683_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
		}), Get_Data_FunctorWithIndex_mapWithIndexArray()})))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexArray
}

var cache_Data_FunctorWithIndex_functorWithIndexApp gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexApp sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexApp() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexApp.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexApp = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_FunctorWithIndex_functorWithIndexApp(dictFunctorWithIndex_0_box)
		})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexApp
}

var cache_Data_FunctorWithIndex_functorWithIndexAdditive gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexAdditive sync.Once

func Get_Data_FunctorWithIndex_functorWithIndexAdditive() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexAdditive.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexAdditive = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Additive_functorAdditive()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar a$scope181)] (TypeVar b$scope182))
			__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
			_ = __local_var_1_0
			return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_1_0, m_2)
			})
		})}))}
	})
	return cache_Data_FunctorWithIndex_functorWithIndexAdditive
}

type Constructor_Data_FunctorWithIndex_FunctorWithIndex[T_i any, T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00460075006e00630074006f007200570069007400680049006e006400650078:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00460075006e00630074006f007200570069007400680049006e006400650078005f00460075006e00630074006f007200570069007400680049006e006400650078
func init() {
	gopurs_runtime.StructGetters[4077743418] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "mapWithIndex":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_FunctorWithIndex_FunctorWithIndex: " + key)
		}
	}
}

func Call_Data_FunctorWithIndex_FunctorWithIndex_dollar_Dict(x_0_loop struct {
	Functor0     gopurs_runtime.Value
	mapWithIndex gopurs_runtime.Value
	/*gopurs:record-v1:2:00460075006e00630074006f00720030,006d0061007000570069007400680049006e006400650078*/
}) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Functor0     gopurs_runtime.Value
		mapWithIndex gopurs_runtime.Value
		/*gopurs:record-v1:2:00460075006e00630074006f00720030,006d0061007000570069007400680049006e006400650078*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", orig.Functor0, orig.mapWithIndex)
	}())
}

func Call_Data_FunctorWithIndex_mapWithIndex(dict_0_loop *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_FunctorWithIndex_mapDefault(dictFunctorWithIndex_0_loop *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctorWithIndex_0 *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFunctorWithIndex_0_loop
	_ = dictFunctorWithIndex_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(dictFunctorWithIndex_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return f_1
	}))
}

func Call_Data_FunctorWithIndex_functorWithIndexProduct(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
	_ = dictFunctorWithIndex_0
	// TAST (Let): functorProduct_1_0 shape=App(Var) bindingType=Any
	functorProduct_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Product_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorProduct_1_0
	return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): functorProduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope25), (TypeVar g$scope27)])])
		functorProduct1_3_1 := Rebox_Data_FunctorWithIndex_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{}))))
		_ = functorProduct1_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_935379557_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_2363162019_2812149806(functorProduct1_3_1))}
		}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
}

func Call_Data_FunctorWithIndex_functorWithIndexCoproduct(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
	_ = dictFunctorWithIndex_0
	// TAST (Let): functorCoproduct_1_0 shape=App(Var) bindingType=Any
	functorCoproduct_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorCoproduct_1_0
	return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): functorCoproduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope100), (TypeVar g$scope102)])])
		functorCoproduct1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})))
		_ = functorCoproduct1_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCoproduct1_3_1)}
		}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope100) [(TypeVar a)])] (TypeApp (TypeVar f$scope100) [(TypeVar b)]))
			__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()))
			_ = __local_var_6_2
			// TAST (Let): __local_var_7_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope102) [(TypeVar a$scope107)])] (TypeApp (TypeVar g$scope102) [(TypeVar b$scope108)]))
			__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()))
			_ = __local_var_7_3
			var __t4 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if v_5.Type == 9 && v_5.IntVal == 3711209382 {
					__t4 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Apply(__local_var_6_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Value{}, false}
					goto end_branch_4
				} else {

				}
			}
			{
				if v_5.Type == 9 && v_5.IntVal == 2465973597 {
					__t4 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.Apply(__local_var_7_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), true}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
			}
		end_branch_4:
			return func() gopurs_runtime.Value {
				_v := __t4
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})}))}
	})
}

func Call_Data_FunctorWithIndex_functorWithIndexCompose(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
	_ = dictFunctorWithIndex_0
	// TAST (Let): functorCompose_1_0 shape=App(Var) bindingType=Any
	functorCompose_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Compose_functorCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorCompose_1_0
	return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mapWithIndex2_3_1 shape=App(Var) bindingType=(Func [(Func [(TypeVar b$scope142), (TypeVar a$scope148)] (TypeVar b$scope149)), (TypeApp (TypeVar g$scope143) [(TypeVar a$scope148)])] (TypeApp (TypeVar g$scope143) [(TypeVar b$scope149)]))
		mapWithIndex2_3_1 := Call_Data_FunctorWithIndex_mapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFunctorWithIndex1_2))
		_ = mapWithIndex2_3_1
		// TAST (Let): functorCompose1_4_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope141), (TypeVar g$scope143)])])
		functorCompose1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})))
		_ = functorCompose1_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_FunctorWithIndex_1961781733_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose1_4_2)}
		}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), mapWithIndex2_3_1, gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{a_7, b_8}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))})
			})), v_6)
		})})))}
	})
}

func Call_Data_FunctorWithIndex_functorWithIndexApp(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
	_ = dictFunctorWithIndex_0
	// TAST (Let): functorApp_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope167)])])
	functorApp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorApp_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorApp_1_0)}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), f_2, v_3)
	})}))}
}

func Rebox_Data_FunctorWithIndex_1688994542_1495429347(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_1961781733_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_2773701683_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_3136246921_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_3689823567_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_FunctorWithIndex_935379557_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Data_FunctorWithIndex_mapWithIndexArray() gopurs_runtime.Value {
	return _Gopurs_Data_FunctorWithIndex_MapWithIndexArray
}
