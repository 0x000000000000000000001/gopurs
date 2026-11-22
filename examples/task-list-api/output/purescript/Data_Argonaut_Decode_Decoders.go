package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Decode_Decoders_pure gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_pure sync.Once

func Get_Data_Argonaut_Decode_Decoders_pure() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_pure.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_pure = Get_Data_Either_Right()
	})
	return cache_Data_Argonaut_Decode_Decoders_pure
}

var cache_Data_Argonaut_Decode_Decoders_pure1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_pure1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_pure1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_pure1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_pure1 = Get_Data_Either_Right()
	})
	return cache_Data_Argonaut_Decode_Decoders_pure1
}

var cache_Data_Argonaut_Decode_Decoders_bindFlipped gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_bindFlipped sync.Once

func Get_Data_Argonaut_Decode_Decoders_bindFlipped() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_bindFlipped.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_bindFlipped = Get_Control_Bind_bindFlipped__79681002()
	})
	return cache_Data_Argonaut_Decode_Decoders_bindFlipped
}

var cache_Data_Argonaut_Decode_Decoders_go__map gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_go__map sync.Once

func Get_Data_Argonaut_Decode_Decoders_go__map() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_go__map.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_go__map = Get_Data_Functor_map__2013991078()
	})
	return cache_Data_Argonaut_Decode_Decoders_go__map
}

var cache_Data_Argonaut_Decode_Decoders_pure2 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_pure2 sync.Once

func Get_Data_Argonaut_Decode_Decoders_pure2() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_pure2.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_pure2 = Get_Data_Either_Right()
	})
	return cache_Data_Argonaut_Decode_Decoders_pure2
}

var cache_Data_Argonaut_Decode_Decoders_map1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map1 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Decoders_map1
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped(f_0_box, g_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped
}

var cache_Data_Argonaut_Decode_Decoders_compose gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose
}

var cache_Data_Argonaut_Decode_Decoders_lmap gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap
}

var cache_Data_Argonaut_Decode_Decoders_traverse gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverse sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverse() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverse.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverse = gopurs_runtime.Apply(Rebox_Data_Argonaut_Decode_Decoders_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))})
	})
	return cache_Data_Argonaut_Decode_Decoders_traverse
}

var cache_Data_Argonaut_Decode_Decoders_map2 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map2 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map2() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map2.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map2 = Get_Data_Functor_map__4275211926()
	})
	return cache_Data_Argonaut_Decode_Decoders_map2
}

var cache_Data_Argonaut_Decode_Decoders_map3 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map3 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map3() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map3.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map3 = Get_Data_Functor_map__3659863366()
	})
	return cache_Data_Argonaut_Decode_Decoders_map3
}

var cache_Data_Argonaut_Decode_Decoders_fromFoldable gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_fromFoldable sync.Once

func Get_Data_Argonaut_Decode_Decoders_fromFoldable() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_fromFoldable.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_fromFoldable = Call_Data_List_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))
	})
	return cache_Data_Argonaut_Decode_Decoders_fromFoldable
}

var cache_Data_Argonaut_Decode_Decoders_compose1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose1
}

var cache_Data_Argonaut_Decode_Decoders_map4 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map4 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map4() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map4.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map4 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Decoders_map4
}

var cache_Data_Argonaut_Decode_Decoders_fromFoldable1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_fromFoldable1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_fromFoldable1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_fromFoldable1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_fromFoldable1 = gopurs_runtime.Apply(Get_Data_Set_fromFoldable(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_Decoders_1022383170_1680800814(Rebox_Data_Argonaut_Decode_Decoders_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))})
	})
	return cache_Data_Argonaut_Decode_Decoders_fromFoldable1
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1(f_0_box, g_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1
}

var cache_Data_Argonaut_Decode_Decoders_compose2 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose2 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose2() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose2.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose2
}

var cache_Data_Argonaut_Decode_Decoders_lmap1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap1 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap1(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap1
}

var cache_Data_Argonaut_Decode_Decoders_traverse1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverse1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverse1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverse1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverse1 = func() gopurs_runtime.Value {
			traverseEither_arg_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())
			_ = traverseEither_arg_0
			_ = traverseEither_arg_0
			traverseEither_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither())
			_ = traverseEither_arg_1
			_ = traverseEither_arg_1
			return gopurs_runtime.Func2(func(traverseEither_callback_2 gopurs_runtime.Value, traverseEither_input_2 gopurs_runtime.Value) gopurs_runtime.Value {
				traverseEither_values_2 := *((*[]gopurs_runtime.Value)(traverseEither_input_2.UnsafePtr))
				_ = traverseEither_values_2
				traverseEither_output_2 := make([]gopurs_runtime.Value, len(traverseEither_values_2))
				_ = traverseEither_output_2
				traverseEither_firstError_2 := gopurs_runtime.Value{}
				_ = traverseEither_firstError_2
				traverseEither_failed_2 := false
				_ = traverseEither_failed_2
				for traverseEither_index_2, traverseEither_value_2 := range traverseEither_values_2 {
					traverseEither_result_2 := gopurs_runtime.Apply(traverseEither_callback_2, traverseEither_value_2)
					_ = traverseEither_result_2
					if traverseEither_result_2.Type == 9 && traverseEither_result_2.IntVal == 2465973597 {
						traverseEither_output_2[traverseEither_index_2] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseEither_result_2.UnsafePtr).V0
					} else {
						if traverseEither_result_2.Type == 9 && traverseEither_result_2.IntVal == 3711209382 {
							if !(traverseEither_failed_2) {
								traverseEither_firstError_2 = traverseEither_result_2
								traverseEither_failed_2 = true
							} else {

							}
						} else {
							panic("Failed pattern match")
						}
					}
				}
				if traverseEither_failed_2 {
					return traverseEither_firstError_2
				} else {

				}
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_2), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
			})
		}()
	})
	return cache_Data_Argonaut_Decode_Decoders_traverse1
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped2 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped2 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped2() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped2.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped2 = Get_Control_Bind_composeKleisliFlipped__647481931()
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped2
}

var cache_Data_Argonaut_Decode_Decoders_compose3 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose3 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose3() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose3.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose3
}

var cache_Data_Argonaut_Decode_Decoders_map5 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map5 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map5() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map5.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map5 = Get_Data_Functor_map__3477210790()
	})
	return cache_Data_Argonaut_Decode_Decoders_map5
}

var cache_Data_Argonaut_Decode_Decoders_compose4 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose4 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose4() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose4.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose4
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3(f_0_box, g_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3
}

var cache_Data_Argonaut_Decode_Decoders_compose5 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose5 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose5() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose5.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose5
}

var cache_Data_Argonaut_Decode_Decoders_lmap2 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap2 sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap2() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap2.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap2 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap2(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap2
}

var cache_Data_Argonaut_Decode_Decoders_traverse2 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverse2 sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverse2() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverse2.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverse2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_Decoders_3037784642_3043886126(Rebox_Data_Argonaut_Decode_Decoders_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))})
	})
	return cache_Data_Argonaut_Decode_Decoders_traverse2
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped4 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped4 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped4() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped4.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped4 = Get_Control_Bind_composeKleisliFlipped__1244140651()
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped4
}

var cache_Data_Argonaut_Decode_Decoders_compose6 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose6 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose6() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose6.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose6
}

var cache_Data_Argonaut_Decode_Decoders_map6 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map6 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map6() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map6.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map6 = Get_Data_Functor_map__1891958886()
	})
	return cache_Data_Argonaut_Decode_Decoders_map6
}

var cache_Data_Argonaut_Decode_Decoders_compose7 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose7 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose7() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose7.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose7
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5(f_0_box, g_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5
}

var cache_Data_Argonaut_Decode_Decoders_compose8 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose8 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose8() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose8.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose8 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose8
}

var cache_Data_Argonaut_Decode_Decoders_lmap3 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap3 sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap3() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap3.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap3 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap3(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap3
}

var cache_Data_Argonaut_Decode_Decoders_traverse3 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverse3 sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverse3() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverse3.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverse3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))})
	})
	return cache_Data_Argonaut_Decode_Decoders_traverse3
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped6 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped6 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped6() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped6.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped6 = Get_Control_Bind_composeKleisliFlipped__2774842251()
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped6
}

var cache_Data_Argonaut_Decode_Decoders_compose9 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose9 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose9() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose9.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose9 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose9
}

var cache_Data_Argonaut_Decode_Decoders_map7 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map7 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map7() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map7.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map7 = Get_Data_Functor_map__1962508006()
	})
	return cache_Data_Argonaut_Decode_Decoders_map7
}

var cache_Data_Argonaut_Decode_Decoders_compose10 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose10 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose10() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose10.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose10 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose10
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7(f_0_box, g_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7
}

var cache_Data_Argonaut_Decode_Decoders_compose11 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose11 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose11() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose11.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose11 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose11
}

var cache_Data_Argonaut_Decode_Decoders_lmap4 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap4 sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap4() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap4.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap4 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap4(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap4
}

var cache_Data_Argonaut_Decode_Decoders_traverse4 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverse4 sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverse4() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverse4.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverse4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_Decoders_3037784642_3043886126(Rebox_Data_Argonaut_Decode_Decoders_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))})
	})
	return cache_Data_Argonaut_Decode_Decoders_traverse4
}

var cache_Data_Argonaut_Decode_Decoders_compose12 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose12 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose12() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose12.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose12 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose12
}

var cache_Data_Argonaut_Decode_Decoders_compose13 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose13 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose13() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose13.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose13 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose13
}

var cache_Data_Argonaut_Decode_Decoders_map8 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map8 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map8() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map8.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map8 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Decoders_map8
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8(f_0_box, g_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8
}

var cache_Data_Argonaut_Decode_Decoders_compose14 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose14 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose14() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose14.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose14 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose14
}

var cache_Data_Argonaut_Decode_Decoders_lmap5 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap5 sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap5() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap5.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap5 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap5(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap5
}

var cache_Data_Argonaut_Decode_Decoders_traverse5 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverse5 sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverse5() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverse5.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverse5 = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Rebox_Data_Argonaut_Decode_Decoders_1812164904_3467869399(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Object_traversableWithIndexObject())).V3, Get_Data_Either_applicativeEither()), Get_Data_Function_go__const())
	})
	return cache_Data_Argonaut_Decode_Decoders_traverse5
}

var cache_Data_Argonaut_Decode_Decoders_lmap6 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap6 sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap6() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap6.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap6 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap6(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap6
}

var cache_Data_Argonaut_Decode_Decoders_bind gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_bind sync.Once

func Get_Data_Argonaut_Decode_Decoders_bind() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_bind.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_bind = gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind")
	})
	return cache_Data_Argonaut_Decode_Decoders_bind
}

var cache_Data_Argonaut_Decode_Decoders_bind1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_bind1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_bind1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_bind1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_bind1 = gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind")
	})
	return cache_Data_Argonaut_Decode_Decoders_bind1
}

var cache_Data_Argonaut_Decode_Decoders_map9 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map9 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map9() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map9.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map9 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Decoders_map9
}

var cache_Data_Argonaut_Decode_Decoders_map10 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map10 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map10() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map10.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map10 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Decoders_map10
}

var cache_Data_Argonaut_Decode_Decoders_bindFlipped1 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_bindFlipped1 sync.Once

func Get_Data_Argonaut_Decode_Decoders_bindFlipped1() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_bindFlipped1.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_bindFlipped1 = Get_Control_Bind_bindFlipped__2935978538()
	})
	return cache_Data_Argonaut_Decode_Decoders_bindFlipped1
}

var cache_Data_Argonaut_Decode_Decoders_map11 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map11 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map11() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map11.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map11 = Get_Data_Functor_map__1881518470()
	})
	return cache_Data_Argonaut_Decode_Decoders_map11
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9(f_0_box, g_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9
}

var cache_Data_Argonaut_Decode_Decoders_decodeJArray gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeJArray sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeJArray() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeJArray.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeJArray = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Either_note__2066121466(gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "Array"}))}), Call_Data_Argonaut_Core_toJsonType__3906794974(Get_Data_Argonaut_Core_caseJsonArray()))
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeJArray
}

var cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745 sync.Once

func Get_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_unused_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745(__eta_norm_2_0_box, __eta_norm_1_unused_1_box, __eta_norm_0_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745
}

var cache_Data_Argonaut_Decode_Decoders_compose15 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose15 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose15() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose15.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose15 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose15
}

var cache_Data_Argonaut_Decode_Decoders_compose15__3452995720 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose15__3452995720 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose15__3452995720() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose15__3452995720.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose15__3452995720 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_compose15__3452995720(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_compose15__3452995720
}

var cache_Data_Argonaut_Decode_Decoders_lmap7 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_lmap7 sync.Once

func Get_Data_Argonaut_Decode_Decoders_lmap7() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_lmap7.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_lmap7 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_lmap7(f_0_box, v2_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_lmap7
}

var cache_Data_Argonaut_Decode_Decoders_traverseWithIndex gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverseWithIndex sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverseWithIndex() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverseWithIndex.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverseWithIndex = func() gopurs_runtime.Value {
			traverseEither_arg_0 := Rebox_Data_Argonaut_Decode_Decoders_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))
			_ = traverseEither_arg_0
			_ = traverseEither_arg_0
			traverseEither_arg_1 := Get_Data_Either_applicativeEither()
			_ = traverseEither_arg_1
			_ = traverseEither_arg_1
			return gopurs_runtime.Func2(func(traverseEither_callback_2 gopurs_runtime.Value, traverseEither_input_2 gopurs_runtime.Value) gopurs_runtime.Value {
				traverseEither_values_2 := *((*[]gopurs_runtime.Value)(traverseEither_input_2.UnsafePtr))
				_ = traverseEither_values_2
				traverseEither_output_2 := make([]gopurs_runtime.Value, len(traverseEither_values_2))
				_ = traverseEither_output_2
				traverseEither_firstError_2 := gopurs_runtime.Value{}
				_ = traverseEither_firstError_2
				traverseEither_failed_2 := false
				_ = traverseEither_failed_2
				for traverseEither_index_2, traverseEither_value_2 := range traverseEither_values_2 {
					traverseEither_result_2 := gopurs_runtime.Apply2(traverseEither_callback_2, gopurs_runtime.Int(int64(traverseEither_index_2)), traverseEither_value_2)
					_ = traverseEither_result_2
					if traverseEither_result_2.Type == 9 && traverseEither_result_2.IntVal == 2465973597 {
						traverseEither_output_2[traverseEither_index_2] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseEither_result_2.UnsafePtr).V0
					} else {
						if traverseEither_result_2.Type == 9 && traverseEither_result_2.IntVal == 3711209382 {
							if !(traverseEither_failed_2) {
								traverseEither_firstError_2 = traverseEither_result_2
								traverseEither_failed_2 = true
							} else {

							}
						} else {
							panic("Failed pattern match")
						}
					}
				}
				if traverseEither_failed_2 {
					return traverseEither_firstError_2
				} else {

				}
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_2), true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
			})
		}()
	})
	return cache_Data_Argonaut_Decode_Decoders_traverseWithIndex
}

var cache_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703 sync.Once

func Get_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703(__eta_norm_0_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703
}

var cache_Data_Argonaut_Decode_Decoders_apply gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_apply sync.Once

func Get_Data_Argonaut_Decode_Decoders_apply() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_apply.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_apply = gopurs_runtime.RecordGet(Get_Data_Either_applyEither(), "apply")
	})
	return cache_Data_Argonaut_Decode_Decoders_apply
}

var cache_Data_Argonaut_Decode_Decoders_map12 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map12 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map12() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map12.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map12 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Decoders_map12
}

var cache_Data_Argonaut_Decode_Decoders_bind2 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_bind2 sync.Once

func Get_Data_Argonaut_Decode_Decoders_bind2() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_bind2.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_bind2 = gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind")
	})
	return cache_Data_Argonaut_Decode_Decoders_bind2
}

var cache_Data_Argonaut_Decode_Decoders_compose16 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_compose16 sync.Once

func Get_Data_Argonaut_Decode_Decoders_compose16() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_compose16.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_compose16 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Decode_Decoders_compose16
}

var cache_Data_Argonaut_Decode_Decoders_map13 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_map13 sync.Once

func Get_Data_Argonaut_Decode_Decoders_map13() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_map13.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_map13 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Decoders_map13
}

var cache_Data_Argonaut_Decode_Decoders_getFieldOptional_prime_ gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_getFieldOptional_prime_ sync.Once

func Get_Data_Argonaut_Decode_Decoders_getFieldOptional_prime_() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_getFieldOptional_prime_.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_getFieldOptional_prime_ = gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Internal_Record_getFieldOptionalNullableImpl(), gopurs_runtime.Func3(func(decoder_0 gopurs_runtime.Value, obj_1 gopurs_runtime.Value, str_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Argonaut","Core","Json"] [])])
			__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.NormalizeStringValue(str_2), obj_1))
			_ = __local_var_3_0
			var __t4 gopurs_runtime.Value
			{
				if __local_var_3_0 == nil {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}))}
					goto end_branch_4
				} else {

				}
			}
			{
				if __local_var_3_0 != nil {
					var __t3 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}
					{
						if (Call_Data_Argonaut_Core_caseJsonNull(gopurs_runtime.Bool(false), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(true)
						}), (__local_var_3_0).V0).IntVal) != (0) {
							__t3 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, true}
							goto end_branch_3
						} else {

						}
					}
					{
						// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope1)])
						v_4_1 := gopurs_runtime.Apply2(Get_Data_Argonaut_Decode_Internal_Record_decodeFieldFast(), decoder_0, (__local_var_3_0).V0)
						_ = v_4_1
						var __t2 struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}
						{
							if v_4_1.Type == 9 && v_4_1.IntVal == 3711209382 {
								__t2 = struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
									V2 bool
								}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, str_2.StrVal(), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0}))}, gopurs_runtime.Value{}, false}
								goto end_branch_2
							} else {

							}
						}
						{
							if v_4_1.Type == 9 && v_4_1.IntVal == 2465973597 {
								__t2 = struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
									V2 bool
								}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0}))}, true}
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = func() struct {
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
					end_branch_2:
						__t3 = __t2
					}
				end_branch_3:
					__t4 = func() gopurs_runtime.Value {
						_v := __t3
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
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
	})
	return cache_Data_Argonaut_Decode_Decoders_getFieldOptional_prime_
}

var cache_Data_Argonaut_Decode_Decoders_getFieldOptional gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_getFieldOptional sync.Once

func Get_Data_Argonaut_Decode_Decoders_getFieldOptional() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_getFieldOptional.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_getFieldOptional = gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Internal_Record_getFieldOptionalImpl(), gopurs_runtime.Func3(func(decoder_0 gopurs_runtime.Value, obj_1 gopurs_runtime.Value, str_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Argonaut","Core","Json"] [])])
			__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.NormalizeStringValue(str_2), obj_1))
			_ = __local_var_3_0
			var __t3 gopurs_runtime.Value
			{
				if __local_var_3_0 == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}))}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_3_0 != nil {
					// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope12)])
					v_4_1 := gopurs_runtime.Apply2(Get_Data_Argonaut_Decode_Internal_Record_decodeFieldFast(), decoder_0, (__local_var_3_0).V0)
					_ = v_4_1
					var __t2 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}
					{
						if v_4_1.Type == 9 && v_4_1.IntVal == 3711209382 {
							__t2 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, str_2.StrVal(), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0}))}, gopurs_runtime.Value{}, false}
							goto end_branch_2
						} else {

						}
					}
					{
						if v_4_1.Type == 9 && v_4_1.IntVal == 2465973597 {
							__t2 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_1.UnsafePtr).V0}))}, true}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = func() struct {
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
				end_branch_2:
					__t3 = func() gopurs_runtime.Value {
						_v := __t2
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		}))
	})
	return cache_Data_Argonaut_Decode_Decoders_getFieldOptional
}

var cache_Data_Argonaut_Decode_Decoders_getField gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_getField sync.Once

func Get_Data_Argonaut_Decode_Decoders_getField() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_getField.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_getField = gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Internal_Record_getFieldImpl(), gopurs_runtime.Func3(func(decoder_0 gopurs_runtime.Value, obj_1 gopurs_runtime.Value, str_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Argonaut","Core","Json"] [])])
			v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.NormalizeStringValue(str_2), obj_1))
			_ = v_3_0
			var __t3 gopurs_runtime.Value
			{
				if v_3_0 == nil {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, str_2.StrVal(), gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}}))}
					goto end_branch_3
				} else {

				}
			}
			{
				if v_3_0 != nil {
					// TAST (Let): v1_4_1 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope22)])
					v1_4_1 := gopurs_runtime.Apply2(Get_Data_Argonaut_Decode_Internal_Record_decodeFieldFast(), decoder_0, (v_3_0).V0)
					_ = v1_4_1
					var __t2 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}
					{
						if v1_4_1.Type == 9 && v1_4_1.IntVal == 3711209382 {
							__t2 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, str_2.StrVal(), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0}))}, gopurs_runtime.Value{}, false}
							goto end_branch_2
						} else {

						}
					}
					{
						if v1_4_1.Type == 9 && v1_4_1.IntVal == 2465973597 {
							__t2 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, true}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = func() struct {
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
				end_branch_2:
					__t3 = func() gopurs_runtime.Value {
						_v := __t2
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		}))
	})
	return cache_Data_Argonaut_Decode_Decoders_getField
}

var cache_Data_Argonaut_Decode_Decoders_decodeVoid gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeVoid sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeVoid() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeVoid.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeVoid = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeVoid(v_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeVoid
}

var cache_Data_Argonaut_Decode_Decoders_decodeString gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeString sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeString() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeString.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeString = gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonString(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "String"}))}, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeString
}

var cache_Data_Argonaut_Decode_Decoders_decodeNumber gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeNumber sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeNumber() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeNumber.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeNumber = gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonNumber(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "Number"}))}, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeNumber
}

var cache_Data_Argonaut_Decode_Decoders_decodeNull gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeNull sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeNull() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeNull.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeNull = gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonNull(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "null"}))}, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), gopurs_runtime.Apply(Get_Data_Function_go__const(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()))
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeNull
}

var cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyString gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeNonEmptyString sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeNonEmptyString() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeNonEmptyString.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyString = gopurs_runtime.Func(func(json_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeNonEmptyString(json_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyString
}

var cache_Data_Argonaut_Decode_Decoders_decodeMaybe gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeMaybe sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeMaybe() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeMaybe.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeMaybe = gopurs_runtime.Func2(func(decoder_0_box gopurs_runtime.Value, json_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeMaybe(decoder_0_box, json_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeMaybe
}

var cache_Data_Argonaut_Decode_Decoders_decodeJObject gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeJObject sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeJObject() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeJObject.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeJObject = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Either_note__3211399674(gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "Object"}))}), Call_Data_Argonaut_Core_toJsonType__1322236059(Get_Data_Argonaut_Core_caseJsonObject()))
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeJObject
}

var cache_Data_Argonaut_Decode_Decoders_decodeList gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeList sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeList() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeList.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeList = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeList(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeList
}

var cache_Data_Argonaut_Decode_Decoders_decodeSet gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeSet sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeSet() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeSet.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeSet
}

var cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray
}

var cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyList gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeNonEmptyList sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeNonEmptyList() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeNonEmptyList.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyList = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeNonEmptyList(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeNonEmptyList
}

var cache_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array
}

var cache_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List
}

var cache_Data_Argonaut_Decode_Decoders_decodeInt gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeInt sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeInt() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeInt.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeInt = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(Func [Number] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), Int]))
			__local_var_0_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Either_note__271183866(gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "Integer"}))}), Get_Data_Int_fromNumber())
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
				__local_var_2_1 := Data_Argonaut_Core__CaseJsonNumber(gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "Number"}))}}))}, Get_Data_Either_Right(), a_1)
				_ = __local_var_2_1
				var __t2 gopurs_runtime.Value
				{
					if __local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 3711209382 {
						__t2 = func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()
						goto end_branch_2
					} else {

					}
				}
				{
					if __local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 2465973597 {
						__t2 = gopurs_runtime.Apply(__local_var_0_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0)
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_2:
				return __t2
			})
		}()
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeInt
}

var cache_Data_Argonaut_Decode_Decoders_decodeIdentity gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeIdentity sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeIdentity() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeIdentity.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeIdentity = gopurs_runtime.Func2(func(decoder_0_box gopurs_runtime.Value, json_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeIdentity(decoder_0_box, json_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeIdentity
}

var cache_Data_Argonaut_Decode_Decoders_decodeForeignObject gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeForeignObject sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeForeignObject() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeForeignObject.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeForeignObject = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeForeignObject(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeForeignObject
}

var cache_Data_Argonaut_Decode_Decoders_decodeEither gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeEither sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeEither() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeEither.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeEither = gopurs_runtime.Func3(func(decoderA_0_box gopurs_runtime.Value, decoderB_1_box gopurs_runtime.Value, json_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeEither(decoderA_0_box, decoderB_1_box, json_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeEither
}

var cache_Data_Argonaut_Decode_Decoders_decodeCodePoint gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeCodePoint sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeCodePoint() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeCodePoint.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeCodePoint = gopurs_runtime.Func(func(json_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeCodePoint(json_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeCodePoint
}

var cache_Data_Argonaut_Decode_Decoders_decodeBoolean gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeBoolean sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeBoolean() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeBoolean.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeBoolean = gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonBoolean(), func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "Boolean"}))}, gopurs_runtime.Value{}, false}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}(), Get_Data_Either_Right())
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeBoolean
}

var cache_Data_Argonaut_Decode_Decoders_decodeArray gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeArray sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeArray() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeArray.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeArray = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeArray(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeArray
}

var cache_Data_Argonaut_Decode_Decoders_decodeArray__1478683895 gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeArray__1478683895 sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeArray__1478683895() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeArray__1478683895.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeArray__1478683895 = gopurs_runtime.Func2(func(decoder_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeArray__1478683895(decoder_unused_0_box, __eta_norm_0_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeArray__1478683895
}

var cache_Data_Argonaut_Decode_Decoders_decodeTuple gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeTuple sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeTuple() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeTuple.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeTuple = gopurs_runtime.Func3(func(decoderA_0_box gopurs_runtime.Value, decoderB_1_box gopurs_runtime.Value, json_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Decoders_decodeTuple(decoderA_0_box, decoderB_1_box, json_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeTuple
}

var cache_Data_Argonaut_Decode_Decoders_decodeMap gopurs_runtime.Value
var once_Data_Argonaut_Decode_Decoders_decodeMap sync.Once

func Get_Data_Argonaut_Decode_Decoders_decodeMap() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Decoders_decodeMap.Do(func() {
		cache_Data_Argonaut_Decode_Decoders_decodeMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Decoders_decodeMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Argonaut_Decode_Decoders_decodeMap
}

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
	__local_var_3_0 := gopurs_runtime.Apply(g_1, a_2)
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_lmap(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped1(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
	__local_var_3_0 := gopurs_runtime.Apply(g_1, a_2)
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_lmap1(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped3(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
	__local_var_3_0 := gopurs_runtime.Apply(g_1, a_2)
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_lmap2(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped5(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
	__local_var_3_0 := gopurs_runtime.Apply(g_1, a_2)
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_lmap3(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped7(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
	__local_var_3_0 := gopurs_runtime.Apply(g_1, a_2)
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_lmap4(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped8(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
	__local_var_3_0 := gopurs_runtime.Apply(g_1, a_2)
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_lmap5(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_lmap6(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
	__local_var_3_0 := gopurs_runtime.Apply(g_1, a_2)
	_ = __local_var_3_0
	var __t1 gopurs_runtime.Value
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
			__t1 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745(__eta_norm_2_0_loop gopurs_runtime.Value, __eta_norm_1_unused_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
composeKleisliFlipped9__2030527745:
	for {
		if false {
			continue composeKleisliFlipped9__2030527745
		}
		var __eta_norm_2_0 gopurs_runtime.Value = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_unused_1 gopurs_runtime.Value = __eta_norm_1_unused_1_loop
		_ = __eta_norm_1_unused_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJArray(), __eta_norm_0_2)
		_ = __local_var_3_0
		var __t1 gopurs_runtime.Value
		{
			if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 3711209382 {
				__t1 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_1
			} else {

			}
		}
		{
			if __local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 2465973597 {
				__t1 = gopurs_runtime.Apply(__eta_norm_2_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_1:
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := __t1; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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
}

func Call_Data_Argonaut_Decode_Decoders_compose15__3452995720(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose15__3452995720:
	for {
		if false {
			continue compose15__3452995720
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Data_Argonaut_Decode_Decoders_lmap7(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v2_1 gopurs_runtime.Value = v2_1_loop
	_ = v2_1
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 3711209382 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_1.Type == 9 && v2_1.IntVal == 2465973597 {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() struct {
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
end_branch_0:
	return __t0
}

func Call_Data_Argonaut_Decode_Decoders_traverseWithIndex__1902959703(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex__1902959703:
	for {
		if false {
			continue traverseWithIndex__1902959703
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		traverseEither_arg_0 := Rebox_Data_Argonaut_Decode_Decoders_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))
		_ = traverseEither_arg_0
		_ = traverseEither_arg_0
		traverseEither_arg_1 := Get_Data_Either_applicativeEither()
		_ = traverseEither_arg_1
		_ = traverseEither_arg_1
		traverseEither_arg_2 := __eta_norm_0_0
		_ = traverseEither_arg_2
		_ = traverseEither_arg_2
		return gopurs_runtime.Func(func(traverseEither_input_3 gopurs_runtime.Value) gopurs_runtime.Value {
			traverseEither_values_3 := *((*[]gopurs_runtime.Value)(traverseEither_input_3.UnsafePtr))
			_ = traverseEither_values_3
			traverseEither_output_3 := make([]gopurs_runtime.Value, len(traverseEither_values_3))
			_ = traverseEither_output_3
			traverseEither_firstError_3 := gopurs_runtime.Value{}
			_ = traverseEither_firstError_3
			traverseEither_failed_3 := false
			_ = traverseEither_failed_3
			for traverseEither_index_3, traverseEither_value_3 := range traverseEither_values_3 {
				traverseEither_result_3 := gopurs_runtime.Apply2(traverseEither_arg_2, gopurs_runtime.Int(int64(traverseEither_index_3)), traverseEither_value_3)
				_ = traverseEither_result_3
				if traverseEither_result_3.Type == 9 && traverseEither_result_3.IntVal == 2465973597 {
					traverseEither_output_3[traverseEither_index_3] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseEither_result_3.UnsafePtr).V0
				} else {
					if traverseEither_result_3.Type == 9 && traverseEither_result_3.IntVal == 3711209382 {
						if !(traverseEither_failed_3) {
							traverseEither_firstError_3 = traverseEither_result_3
							traverseEither_failed_3 = true
						} else {

						}
					} else {
						panic("Failed pattern match")
					}
				}
			}
			if traverseEither_failed_3 {
				return traverseEither_firstError_3
			} else {

			}
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_3), true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	}
}

func Call_Data_Argonaut_Decode_Decoders_decodeVoid(v_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}{gopurs_runtime.Value{Type: 9, IntVal: 2101242581, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_UnexpectedValue{1, Data_Argonaut_Core_FromString_nativeWorker(gopurs_runtime.StrWithCodeUnits("Value cannot be Void", 20))}))}, gopurs_runtime.Value{}, false}
}

func Call_Data_Argonaut_Decode_Decoders_decodeNonEmptyString(json_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var json_0 gopurs_runtime.Value = json_0_loop
	_ = json_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Data_Argonaut_Core__CaseJsonString(gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "String"}))}}))}, Get_Data_Either_Right(), json_0)
	_ = __local_var_1_0
	var __t2 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 3711209382 {
			__t2 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
			goto end_branch_2
		} else {

		}
	}
	{
		if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 2465973597 {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if ((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0.StrVal()) == ("") {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 2718035288, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_Named{1, "NonEmptyString", gopurs_runtime.Value{Type: 9, IntVal: 2101242581, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_UnexpectedValue{1, json_0}))}}))}, gopurs_runtime.Value{}, false}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0, true}
			}
		end_branch_1:
			__t2 = __t1
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() struct {
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
end_branch_2:
	return __t2
}

func Call_Data_Argonaut_Decode_Decoders_decodeMaybe(decoder_0_loop gopurs_runtime.Value, json_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	var json_1 gopurs_runtime.Value = json_1_loop
	_ = json_1
	var __t2 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if (Call_Data_Argonaut_Core_caseJsonNull(gopurs_runtime.Bool(false), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		}), json_1).IntVal) != (0) {
			__t2 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, true}
			goto end_branch_2
		} else {

		}
	}
	{
		// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope46)])
		__local_var_2_0 := gopurs_runtime.Apply(decoder_0, json_1)
		_ = __local_var_2_0
		var __t1 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3711209382 {
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				goto end_branch_1
			} else {

			}
		}
		{
			if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 2465973597 {
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0}))}, true}
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
		__t2 = __t1
	}
end_branch_2:
	return __t2
}

func Call_Data_Argonaut_Decode_Decoders_decodeList(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_1_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("List", 4))
	_ = __local_var_1_1
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","List","Types","List"] [(ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","List","Types","List"] [(TypeVar a$scope52)])]))
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Apply(__local_var_1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, true}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
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
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}), gopurs_runtime.Apply2(Rebox_Data_Argonaut_Decode_Decoders_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))}, decoder_0))
	_ = __local_var_1_0
	// TAST (Let): __local_var_2_3 shape=App(Var) bindingType=Any
	__local_var_2_3 := gopurs_runtime.Apply(Get_Data_Functor_map__3659863366(), Call_Data_List_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray())))
	_ = __local_var_2_3
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_4_4 := gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJArray(), a_3))
		_ = __local_var_4_4
		var __t5 gopurs_runtime.Value
		{
			if __local_var_4_4.Type == 9 && __local_var_4_4.IntVal == 3711209382 {
				__t5 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_4.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_5
			} else {

			}
		}
		{
			if __local_var_4_4.Type == 9 && __local_var_4_4.IntVal == 2465973597 {
				__t5 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_4.UnsafePtr).V0)
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

func Call_Data_Argonaut_Decode_Decoders_decodeSet(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): fromFoldable2_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a$scope55)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope55), Unit]))
	fromFoldable2_1_0 := Call_Data_Set_fromFoldable(Rebox_Data_Argonaut_Decode_Decoders_1022383170_1680800814(Rebox_Data_Argonaut_Decode_Decoders_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))), dictOrd_0)
	_ = fromFoldable2_1_0
	return gopurs_runtime.Func(func(decoder_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if m_3.Type == 9 && m_3.IntVal == 3711209382 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					goto end_branch_1
				} else {

				}
			}
			{
				if m_3.Type == 9 && m_3.IntVal == 2465973597 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.Apply(fromFoldable2_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), true}
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
		}), Call_Data_Argonaut_Decode_Decoders_decodeList(decoder_2))
	})
}

func Call_Data_Argonaut_Decode_Decoders_decodeNonEmptyArray(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_1_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("NonEmptyArray", 13))
	_ = __local_var_1_1
	traverseEither_arg_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())
	_ = traverseEither_arg_3
	_ = traverseEither_arg_3
	traverseEither_arg_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither())
	_ = traverseEither_arg_4
	_ = traverseEither_arg_4
	traverseEither_arg_5 := decoder_0
	_ = traverseEither_arg_5
	_ = traverseEither_arg_5
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(Array (ADT ["Data","Argonaut","Core","Json"] []))] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Array (TypeVar a$scope58))]))
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Apply(__local_var_1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, true}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
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
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}), gopurs_runtime.Func(func(traverseEither_input_6 gopurs_runtime.Value) gopurs_runtime.Value {
		traverseEither_values_6 := *((*[]gopurs_runtime.Value)(traverseEither_input_6.UnsafePtr))
		_ = traverseEither_values_6
		traverseEither_output_6 := make([]gopurs_runtime.Value, len(traverseEither_values_6))
		_ = traverseEither_output_6
		traverseEither_firstError_6 := gopurs_runtime.Value{}
		_ = traverseEither_firstError_6
		traverseEither_failed_6 := false
		_ = traverseEither_failed_6
		for traverseEither_index_6, traverseEither_value_6 := range traverseEither_values_6 {
			traverseEither_result_6 := gopurs_runtime.Apply(traverseEither_arg_5, traverseEither_value_6)
			_ = traverseEither_result_6
			if traverseEither_result_6.Type == 9 && traverseEither_result_6.IntVal == 2465973597 {
				traverseEither_output_6[traverseEither_index_6] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseEither_result_6.UnsafePtr).V0
			} else {
				if traverseEither_result_6.Type == 9 && traverseEither_result_6.IntVal == 3711209382 {
					if !(traverseEither_failed_6) {
						traverseEither_firstError_6 = traverseEither_result_6
						traverseEither_failed_6 = true
					} else {

					}
				} else {
					panic("Failed pattern match")
				}
			}
		}
		if traverseEither_failed_6 {
			return traverseEither_firstError_6
		} else {

		}
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_6), true}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}))
	_ = __local_var_1_0
	// TAST (Let): __local_var_2_7 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Array (ADT ["Data","Argonaut","Core","Json"] []))]))
	__local_var_2_7 := gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__647481931(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Functor_map__3477210790(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(x_2, "head")}), gopurs_runtime.RecordGet(x_2, "tail"))
	})), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Either_note__67430010(gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "NonEmptyArray"}))}), Get_Data_Array_uncons())), Get_Data_Argonaut_Decode_Decoders_decodeJArray())
	_ = __local_var_2_7
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_8 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_4_8 := gopurs_runtime.Apply(__local_var_2_7, a_3)
		_ = __local_var_4_8
		var __t9 gopurs_runtime.Value
		{
			if __local_var_4_8.Type == 9 && __local_var_4_8.IntVal == 3711209382 {
				__t9 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_9
			} else {

			}
		}
		{
			if __local_var_4_8.Type == 9 && __local_var_4_8.IntVal == 2465973597 {
				__t9 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_8.UnsafePtr).V0)
				goto end_branch_9
			} else {

			}
		}
		{
			__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_9:
		return __t9
	})
}

func Call_Data_Argonaut_Decode_Decoders_decodeNonEmptyList(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_1_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("NonEmptyList", 12))
	_ = __local_var_1_1
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a$scope63)])]))
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Apply(__local_var_1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, true}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
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
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_Decoders_3037784642_3043886126(Rebox_Data_Argonaut_Decode_Decoders_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))}, decoder_0))
	_ = __local_var_1_0
	// TAST (Let): __local_var_2_4 shape=App(Var) bindingType=Any
	__local_var_2_4 := gopurs_runtime.Apply(Get_Data_Functor_map__3659863366(), Call_Data_List_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray())))
	_ = __local_var_2_4
	// TAST (Let): __local_var_2_3 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Data","Argonaut","Core","Json"] [])])]))
	__local_var_2_3 := gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__1244140651(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Functor_map__1891958886(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(x_2, "head"), gopurs_runtime.RecordGet(x_2, "tail")}))}
	})), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Either_note__2268946106(gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "NonEmptyList"}))}), Get_Data_List_uncons())), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJArray(), x_3))
	}))
	_ = __local_var_2_3
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_4_5 := gopurs_runtime.Apply(__local_var_2_3, a_3)
		_ = __local_var_4_5
		var __t6 gopurs_runtime.Value
		{
			if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 3711209382 {
				__t6 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_6
			} else {

			}
		}
		{
			if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 2465973597 {
				__t6 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0)
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

func Call_Data_Argonaut_Decode_Decoders_decodeNonEmpty_Array(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_1_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("NonEmpty Array", 14))
	_ = __local_var_1_1
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Prim","Array"] []), (ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Prim","Array"] []), (TypeVar a$scope68)])]))
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Apply(__local_var_1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, true}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
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
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))}, decoder_0))
	_ = __local_var_1_0
	// TAST (Let): __local_var_2_3 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Prim","Array"] []), (ADT ["Data","Argonaut","Core","Json"] [])])]))
	__local_var_2_3 := gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__2774842251(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Functor_map__1962508006(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(x_2, "head"), gopurs_runtime.RecordGet(x_2, "tail")}))}
	})), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Either_note__67430010(gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "NonEmpty Array"}))}), Get_Data_Array_uncons())), Get_Data_Argonaut_Decode_Decoders_decodeJArray())
	_ = __local_var_2_3
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_4_4 := gopurs_runtime.Apply(__local_var_2_3, a_3)
		_ = __local_var_4_4
		var __t5 gopurs_runtime.Value
		{
			if __local_var_4_4.Type == 9 && __local_var_4_4.IntVal == 3711209382 {
				__t5 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_4.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_5
			} else {

			}
		}
		{
			if __local_var_4_4.Type == 9 && __local_var_4_4.IntVal == 2465973597 {
				__t5 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_4.UnsafePtr).V0)
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

func Call_Data_Argonaut_Decode_Decoders_decodeNonEmpty_List(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_1_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("NonEmpty List", 13))
	_ = __local_var_1_1
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a$scope75)])]))
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Apply(__local_var_1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, true}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
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
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_Decoders_3037784642_3043886126(Rebox_Data_Argonaut_Decode_Decoders_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))}, decoder_0))
	_ = __local_var_1_0
	// TAST (Let): __local_var_2_4 shape=App(Var) bindingType=Any
	__local_var_2_4 := gopurs_runtime.Apply(Get_Data_Functor_map__3659863366(), Call_Data_List_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray())))
	_ = __local_var_2_4
	// TAST (Let): __local_var_2_3 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Data","Argonaut","Core","Json"] [])])]))
	__local_var_2_3 := gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__1244140651(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Functor_map__1891958886(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(x_2, "head"), gopurs_runtime.RecordGet(x_2, "tail")}))}
	})), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Either_note__2268946106(gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "NonEmpty List"}))}), Get_Data_List_uncons())), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJArray(), x_3))
	}))
	_ = __local_var_2_3
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_4_5 := gopurs_runtime.Apply(__local_var_2_3, a_3)
		_ = __local_var_4_5
		var __t6 gopurs_runtime.Value
		{
			if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 3711209382 {
				__t6 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_6
			} else {

			}
		}
		{
			if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 2465973597 {
				__t6 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0)
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

func Call_Data_Argonaut_Decode_Decoders_decodeIdentity(decoder_0_loop gopurs_runtime.Value, json_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	var json_1 gopurs_runtime.Value = json_1_loop
	_ = json_1
	// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope83)])
	__local_var_2_0 := gopurs_runtime.Apply(decoder_0, json_1)
	_ = __local_var_2_0
	var __t1 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3711209382 {
			__t1 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 2465973597 {
			__t1 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0, true}
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
	return __t1
}

func Call_Data_Argonaut_Decode_Decoders_decodeForeignObject(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_1_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("ForeignObject", 13))
	_ = __local_var_1_1
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Foreign","Object","Object"] [(ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Foreign","Object","Object"] [(TypeVar a$scope86)])]))
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Apply(__local_var_1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, true}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
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
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}), gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Rebox_Data_Argonaut_Decode_Decoders_1812164904_3467869399(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Object_traversableWithIndexObject())).V3, Get_Data_Either_applicativeEither()), Get_Data_Function_go__const(), decoder_0))
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_3 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_3_3 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJObject(), a_2)
		_ = __local_var_3_3
		var __t4 gopurs_runtime.Value
		{
			if __local_var_3_3.Type == 9 && __local_var_3_3.IntVal == 3711209382 {
				__t4 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_4
			} else {

			}
		}
		{
			if __local_var_3_3.Type == 9 && __local_var_3_3.IntVal == 2465973597 {
				__t4 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_3.UnsafePtr).V0)
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

func Call_Data_Argonaut_Decode_Decoders_decodeEither(decoderA_0_loop gopurs_runtime.Value, decoderB_1_loop gopurs_runtime.Value, json_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var decoderA_0 gopurs_runtime.Value = decoderA_0_loop
	_ = decoderA_0
	var decoderB_1 gopurs_runtime.Value = decoderB_1_loop
	_ = decoderB_1
	var json_2 gopurs_runtime.Value = json_2_loop
	_ = json_2
	// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("Either", 6))
	_ = __local_var_3_0
	// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
	__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJObject(), json_2)
	_ = __local_var_4_1
	var __t16 gopurs_runtime.Value
	{
		if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 3711209382 {
			// TAST (Let): __local_var_5_3 shape=Other bindingType=(TypeVar e$scope161)
			__local_var_5_3 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0
			_ = __local_var_5_3
			__t16 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{__local_var_5_3, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_16
		} else {

		}
	}
	{
		if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 2465973597 {
			// TAST (Let): __local_var_5_4 shape=Other bindingType=(TypeVar b$scope119)
			__local_var_5_4 := (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0
			_ = __local_var_5_4
			// TAST (Let): __local_var_11_5 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Argonaut","Core","Json"] [])])
			__local_var_11_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.StrWithCodeUnits("tag", 3), __local_var_5_4))
			_ = __local_var_11_5
			var __t15 gopurs_runtime.Value
			{
				if __local_var_11_5 == nil {
					__t15 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "tag", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_15
				} else {

				}
			}
			{
				if __local_var_11_5 != nil {
					// TAST (Let): __local_var_12_6 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Argonaut","Core","Json"] [])])
					__local_var_12_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just(), gopurs_runtime.StrWithCodeUnits("value", 5), __local_var_5_4))
					_ = __local_var_12_6
					var __t14 gopurs_runtime.Value
					{
						if __local_var_12_6 == nil {
							__t14 = func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
									V2 bool
								}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "value", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{}, false}
								if _v.V2 {
									return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
							}()
							goto end_branch_14
						} else {

						}
					}
					{
						if __local_var_12_6 != nil {
							// TAST (Let): v_13_7 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [String])
							v_13_7 := Rebox_Data_Argonaut_Decode_Decoders_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_Data_Argonaut_Core_caseJsonString(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_Decoders_742090555_3094389156(Rebox_Data_Argonaut_Decode_Decoders_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}, Get_Data_Maybe_Just(), (__local_var_11_5).V0)))
							_ = v_13_7
							var __t13 gopurs_runtime.Value
							{
								if v_13_7 != nil {
									var __t12 gopurs_runtime.Value
									{
										if ((v_13_7).V0) == ("Right") {
											// TAST (Let): __local_var_14_8 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar b$scope91)])
											__local_var_14_8 := gopurs_runtime.Apply(decoderB_1, (__local_var_12_6).V0)
											_ = __local_var_14_8
											var __t9 struct {
												V0 gopurs_runtime.Value
												V1 gopurs_runtime.Value
												V2 bool
											}
											{
												if __local_var_14_8.Type == 9 && __local_var_14_8.IntVal == 3711209382 {
													__t9 = struct {
														V0 gopurs_runtime.Value
														V1 gopurs_runtime.Value
														V2 bool
													}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_14_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
													goto end_branch_9
												} else {

												}
											}
											{
												if __local_var_14_8.Type == 9 && __local_var_14_8.IntVal == 2465973597 {
													__t9 = struct {
														V0 gopurs_runtime.Value
														V1 gopurs_runtime.Value
														V2 bool
													}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_14_8.UnsafePtr).V0}))}, true}
													goto end_branch_9
												} else {

												}
											}
											{
												__t9 = func() struct {
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
										end_branch_9:
											__t12 = func() gopurs_runtime.Value {
												_v := __t9
												if _v.V2 {
													return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
												}
												return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
											}()
											goto end_branch_12
										} else {

										}
									}
									{
										if ((v_13_7).V0) == ("Left") {
											// TAST (Let): __local_var_15_10 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope90)])
											__local_var_15_10 := gopurs_runtime.Apply(decoderA_0, (__local_var_12_6).V0)
											_ = __local_var_15_10
											var __t11 struct {
												V0 gopurs_runtime.Value
												V1 gopurs_runtime.Value
												V2 bool
											}
											{
												if __local_var_15_10.Type == 9 && __local_var_15_10.IntVal == 3711209382 {
													__t11 = struct {
														V0 gopurs_runtime.Value
														V1 gopurs_runtime.Value
														V2 bool
													}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_15_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
													goto end_branch_11
												} else {

												}
											}
											{
												if __local_var_15_10.Type == 9 && __local_var_15_10.IntVal == 2465973597 {
													__t11 = struct {
														V0 gopurs_runtime.Value
														V1 gopurs_runtime.Value
														V2 bool
													}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_15_10.UnsafePtr).V0}))}, true}
													goto end_branch_11
												} else {

												}
											}
											{
												__t11 = func() struct {
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
										end_branch_11:
											__t12 = func() gopurs_runtime.Value {
												_v := __t11
												if _v.V2 {
													return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
												}
												return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
											}()
											goto end_branch_12
										} else {

										}
									}
									{
										__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "tag", gopurs_runtime.Value{Type: 9, IntVal: 2101242581, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_UnexpectedValue{1, (__local_var_11_5).V0}))}}))}}))}
									}
								end_branch_12:
									__t13 = __t12
									goto end_branch_13
								} else {

								}
							}
							{
								__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "tag", gopurs_runtime.Value{Type: 9, IntVal: 2101242581, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_UnexpectedValue{1, (__local_var_11_5).V0}))}}))}}))}
							}
						end_branch_13:
							__t14 = __t13
							goto end_branch_14
						} else {

						}
					}
					{
						__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_14:
					__t15 = __t14
					goto end_branch_15
				} else {

				}
			}
			{
				__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_15:
			__t16 = __t15
			goto end_branch_16
		} else {

		}
	}
	{
		__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_16:
	// TAST (Let): __local_var_5_2 shape=Branch(Let(Other), Let(Let(Branch(Other, Let(Branch(Other, Let(Branch(Branch(Let(Branch(Other, Other, def=Other)), Let(Branch(Other, Other, def=Other)), def=Other), def=Other)), def=Other)), def=Other))), def=Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (ADT ["Data","Either","Either"] [(TypeVar a$scope90), (TypeVar b$scope91)])])
	__local_var_5_2 := __t16
	_ = __local_var_5_2
	var __t17 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if __local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 3711209382 {
			__t17 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Apply(__local_var_3_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_5_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
			goto end_branch_17
		} else {

		}
	}
	{
		if __local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 2465973597 {
			__t17 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_5_2.UnsafePtr).V0, true}
			goto end_branch_17
		} else {

		}
	}
	{
		__t17 = func() struct {
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
end_branch_17:
	return __t17
}

func Call_Data_Argonaut_Decode_Decoders_decodeCodePoint(json_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var json_0 gopurs_runtime.Value = json_0_loop
	_ = json_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Data_Argonaut_Core__CaseJsonString(gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "String"}))}}))}, Get_Data_Either_Right(), json_0)
	_ = __local_var_1_0
	var __t3 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 3711209382 {
			__t3 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
			goto end_branch_3
		} else {

		}
	}
	{
		if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 2465973597 {
			// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=Any
			__local_var_2_1 := Call_Data_String_CodePoints_codePointAt(int64(0), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0.StrVal())
			_ = __local_var_2_1
			var __t2 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if !__local_var_2_1.V1 {
					__t2 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 2718035288, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_Named{1, "CodePoint", gopurs_runtime.Value{Type: 9, IntVal: 2101242581, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_UnexpectedValue{1, json_0}))}}))}, gopurs_runtime.Value{}, false}
					goto end_branch_2
				} else {

				}
			}
			{
				if __local_var_2_1.V1 {
					__t2 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, __local_var_2_1.V0, true}
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() struct {
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
		end_branch_2:
			__t3 = __t2
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = func() struct {
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
end_branch_3:
	return __t3
}

func Call_Data_Argonaut_Decode_Decoders_decodeArray(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
	__local_var_1_1 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("Array", 5))
	_ = __local_var_1_1
	traverseEither_arg_3 := Rebox_Data_Argonaut_Decode_Decoders_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))
	_ = traverseEither_arg_3
	_ = traverseEither_arg_3
	traverseEither_arg_4 := Get_Data_Either_applicativeEither()
	_ = traverseEither_arg_4
	_ = traverseEither_arg_4
	traverseEither_arg_7 := func(i_1 int64, json_2 gopurs_runtime.Value) struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} {
	// TAST (Let): v_3_5 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope103)])
	v_3_5 := gopurs_runtime.Apply(decoder_0, json_2); _ = v_3_5; var __t6 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}; {
		if v_3_5.Type == 9 && v_3_5.IntVal == 3711209382 {
			__t6 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{Type: 9, IntVal: 1044667600, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtIndex{1, i_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_5.UnsafePtr).V0}))}, gopurs_runtime.Value{}, false}
			goto end_branch_6
		} else {

		}
	}; {
		if v_3_5.Type == 9 && v_3_5.IntVal == 2465973597 {
			__t6 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_5.UnsafePtr).V0, true}
			goto end_branch_6
		} else {

		}
	}; {
		__t6 = func() struct {
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
	}; end_branch_6:
	return __t6 }
	_ = traverseEither_arg_7
	_ = traverseEither_arg_7
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(Array (ADT ["Data","Argonaut","Core","Json"] []))] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Array (TypeVar a$scope103))]))
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3711209382 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Apply(__local_var_1_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				goto end_branch_2
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 2465973597 {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, true}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() struct {
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
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}), gopurs_runtime.Func(func(traverseEither_input_8 gopurs_runtime.Value) gopurs_runtime.Value {
		traverseEither_values_8 := *((*[]gopurs_runtime.Value)(traverseEither_input_8.UnsafePtr))
		_ = traverseEither_values_8
		traverseEither_output_8 := make([]gopurs_runtime.Value, len(traverseEither_values_8))
		_ = traverseEither_output_8
		traverseEither_firstError_8 := gopurs_runtime.Value{}
		_ = traverseEither_firstError_8
		traverseEither_failed_8 := false
		_ = traverseEither_failed_8
		for traverseEither_index_8, traverseEither_value_8 := range traverseEither_values_8 {
			traverseEither_result_8 := traverseEither_arg_7(int64(traverseEither_index_8), traverseEither_value_8)
			_ = traverseEither_result_8
			if traverseEither_result_8.V2 {
				traverseEither_output_8[traverseEither_index_8] = traverseEither_result_8.V1
			} else {
				if !traverseEither_result_8.V2 {
					if !(traverseEither_failed_8) {
						traverseEither_firstError_8 = func() gopurs_runtime.Value {
							_v := traverseEither_result_8
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()
						traverseEither_failed_8 = true
					} else {

					}
				} else {
					panic("Failed pattern match")
				}
			}
		}
		if traverseEither_failed_8 {
			return traverseEither_firstError_8
		} else {

		}
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_8), true}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	}))
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_9 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope40) [(TypeVar b$scope38)])
		__local_var_3_9 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJArray(), a_2)
		_ = __local_var_3_9
		var __t10 gopurs_runtime.Value
		{
			if __local_var_3_9.Type == 9 && __local_var_3_9.IntVal == 3711209382 {
				__t10 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_10
			} else {

			}
		}
		{
			if __local_var_3_9.Type == 9 && __local_var_3_9.IntVal == 2465973597 {
				__t10 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_9.UnsafePtr).V0)
				goto end_branch_10
			} else {

			}
		}
		{
			__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_10:
		return __t10
	})
}

func Call_Data_Argonaut_Decode_Decoders_decodeArray__1478683895(decoder_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
decodeArray__1478683895:
	for {
		if false {
			continue decodeArray__1478683895
		}
		var decoder_unused_0 gopurs_runtime.Value = decoder_unused_0_loop
		_ = decoder_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] [])] (ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []))
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_Named(), gopurs_runtime.StrWithCodeUnits("Array", 5))
		_ = __local_var_2_0
		traverseEither_arg_2 := Rebox_Data_Argonaut_Decode_Decoders_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))
		_ = traverseEither_arg_2
		_ = traverseEither_arg_2
		traverseEither_arg_3 := Get_Data_Either_applicativeEither()
		_ = traverseEither_arg_3
		_ = traverseEither_arg_3
		traverseEither_arg_4 := func(i_3 int64, json_4 gopurs_runtime.Value) struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{gopurs_runtime.Value{}, json_4, true} }
		_ = traverseEither_arg_4
		_ = traverseEither_arg_4
		return Call_Data_Argonaut_Decode_Decoders_composeKleisliFlipped9__2030527745(gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if v2_3.Type == 9 && v2_3.IntVal == 3711209382 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0), gopurs_runtime.Value{}, false}
					goto end_branch_1
				} else {

				}
			}
			{
				if v2_3.Type == 9 && v2_3.IntVal == 2465973597 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0, true}
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
		}), gopurs_runtime.Func(func(traverseEither_input_5 gopurs_runtime.Value) gopurs_runtime.Value {
			traverseEither_values_5 := *((*[]gopurs_runtime.Value)(traverseEither_input_5.UnsafePtr))
			_ = traverseEither_values_5
			traverseEither_output_5 := make([]gopurs_runtime.Value, len(traverseEither_values_5))
			_ = traverseEither_output_5
			traverseEither_firstError_5 := gopurs_runtime.Value{}
			_ = traverseEither_firstError_5
			traverseEither_failed_5 := false
			_ = traverseEither_failed_5
			for traverseEither_index_5, traverseEither_value_5 := range traverseEither_values_5 {
				traverseEither_result_5 := traverseEither_arg_4(int64(traverseEither_index_5), traverseEither_value_5)
				_ = traverseEither_result_5
				if traverseEither_result_5.V2 {
					traverseEither_output_5[traverseEither_index_5] = traverseEither_result_5.V1
				} else {
					if !traverseEither_result_5.V2 {
						if !(traverseEither_failed_5) {
							traverseEither_firstError_5 = func() gopurs_runtime.Value {
								_v := traverseEither_result_5
								if _v.V2 {
									return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
							}()
							traverseEither_failed_5 = true
						} else {

						}
					} else {
						panic("Failed pattern match")
					}
				}
			}
			if traverseEither_failed_5 {
				return traverseEither_firstError_5
			} else {

			}
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_5), true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})), Get_Data_Argonaut_Decode_Decoders_decodeJArray(), __eta_norm_0_1)
	}
}

func Call_Data_Argonaut_Decode_Decoders_decodeTuple(decoderA_0_loop gopurs_runtime.Value, decoderB_1_loop gopurs_runtime.Value, json_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var decoderA_0 gopurs_runtime.Value = decoderA_0_loop
	_ = decoderA_0
	var decoderB_1 gopurs_runtime.Value = decoderB_1_loop
	_ = decoderB_1
	var json_2 gopurs_runtime.Value = json_2_loop
	_ = json_2
	// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
	__local_var_3_0 := Call_Data_Argonaut_Decode_Decoders_decodeArray__1478683895(Get_Data_Either_Right(), json_2)
	_ = __local_var_3_0
	var __t8 gopurs_runtime.Value
	{
		if !__local_var_3_0.V2 {
			// TAST (Let): __local_var_4_1 shape=Other bindingType=(TypeVar e$scope161)
			__local_var_4_1 := __local_var_3_0.V0
			_ = __local_var_4_1
			__t8 = func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{__local_var_4_1, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
			goto end_branch_8
		} else {

		}
	}
	{
		if __local_var_3_0.V2 {
			// TAST (Let): __local_var_4_2 shape=Other bindingType=(TypeVar b$scope119)
			__local_var_4_2 := __local_var_3_0.V1
			_ = __local_var_4_2
			var __t7 gopurs_runtime.Value
			{
				if (gopurs_runtime.Int(int64(len(*((*[]gopurs_runtime.Value)(__local_var_4_2.UnsafePtr))))).IntVal) == (int64(2)) {
					// TAST (Let): __local_var_8_3 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar a$scope111)])
					__local_var_8_3 := gopurs_runtime.Apply(decoderA_0, gopurs_runtime.ArrayAccess(__local_var_4_2, 0))
					_ = __local_var_8_3
					var __t6 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}
					{
						if __local_var_8_3.Type == 9 && __local_var_8_3.IntVal == 3711209382 {
							__t6 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_8_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
							goto end_branch_6
						} else {

						}
					}
					{
						if __local_var_8_3.Type == 9 && __local_var_8_3.IntVal == 2465973597 {
							// TAST (Let): __local_var_9_4 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (TypeVar b$scope112)])
							__local_var_9_4 := gopurs_runtime.Apply(decoderB_1, gopurs_runtime.ArrayAccess(__local_var_4_2, 1))
							_ = __local_var_9_4
							var __t5 struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}
							{
								if __local_var_9_4.Type == 9 && __local_var_9_4.IntVal == 3711209382 {
									__t5 = struct {
										V0 gopurs_runtime.Value
										V1 gopurs_runtime.Value
										V2 bool
									}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_4.UnsafePtr).V0, gopurs_runtime.Value{}, false}
									goto end_branch_5
								} else {

								}
							}
							{
								if __local_var_9_4.Type == 9 && __local_var_9_4.IntVal == 2465973597 {
									__t5 = struct {
										V0 gopurs_runtime.Value
										V1 gopurs_runtime.Value
										V2 bool
									}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_8_3.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_4.UnsafePtr).V0}))}, true}
									goto end_branch_5
								} else {

								}
							}
							{
								__t5 = func() struct {
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
						end_branch_5:
							__t6 = __t5
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = func() struct {
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
				end_branch_6:
					__t7 = func() gopurs_runtime.Value {
						_v := __t6
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}()
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "Tuple"}))}}))}
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
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := __t8; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Argonaut_Decode_Decoders_decodeMap(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): fromFoldable2_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope121), (TypeVar b$scope122)])])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope121), (TypeVar b$scope122)]))
	fromFoldable2_1_0 := gopurs_runtime.Apply(Call_Data_Map_Internal_fromFoldable(dictOrd_0), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Decode_Decoders_1022383170_1680800814(Rebox_Data_Argonaut_Decode_Decoders_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))})
	_ = fromFoldable2_1_0
	return gopurs_runtime.Func2(func(decoderA_2 gopurs_runtime.Value, decoderB_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if m_4.Type == 9 && m_4.IntVal == 3711209382 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					goto end_branch_1
				} else {

				}
			}
			{
				if m_4.Type == 9 && m_4.IntVal == 2465973597 {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, gopurs_runtime.Apply(fromFoldable2_1_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0), true}
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
		}), Call_Data_Argonaut_Decode_Decoders_decodeList(gopurs_runtime.Apply2(Get_Data_Argonaut_Decode_Decoders_decodeTuple(), decoderA_2, decoderB_3)))
	})
}

func Rebox_Data_Argonaut_Decode_Decoders_1022383170_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_Decoders_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_Decoders_1812164904_2955889203(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_Decoders_1812164904_3467869399(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_Decoders_3037784642_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_Decoders_3043886126_3037784642(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Decode_Decoders_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Argonaut_Decode_Decoders_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}
