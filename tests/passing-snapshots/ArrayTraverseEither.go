package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_traverse gopurs_runtime.Value
var once_Main_traverse sync.Once

func Get_Main_traverse() gopurs_runtime.Value {
	once_Main_traverse.Do(func() {
		cache_Main_traverse = Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))
	})
	return cache_Main_traverse
}

var cache_Main_valueIsSymbol gopurs_runtime.Value
var once_Main_valueIsSymbol sync.Once

func Get_Main_valueIsSymbol() gopurs_runtime.Value {
	once_Main_valueIsSymbol.Do(func() {
		cache_Main_valueIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("value")
		}))
	})
	return cache_Main_valueIsSymbol
}

var cache_Main_eqEither gopurs_runtime.Value
var once_Main_eqEither sync.Once

func Get_Main_eqEither() gopurs_runtime.Value {
	once_Main_eqEither.Do(func() {
		cache_Main_eqEither = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main_eqEither
}

var cache_Main_showEither gopurs_runtime.Value
var once_Main_showEither sync.Once

func Get_Main_showEither() gopurs_runtime.Value {
	once_Main_showEither.Do(func() {
		cache_Main_showEither = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}, Call_Data_Show_showArray(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main_showEither
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_eqMaybe gopurs_runtime.Value
var once_Main_eqMaybe sync.Once

func Get_Main_eqMaybe() gopurs_runtime.Value {
	once_Main_eqMaybe.Do(func() {
		cache_Main_eqMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1888010770_3790796878(Rebox_Main_3790796878_1888010770(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Maybe_eqMaybe(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main_eqMaybe
}

var cache_Main_showMaybe gopurs_runtime.Value
var once_Main_showMaybe sync.Once

func Get_Main_showMaybe() gopurs_runtime.Value {
	once_Main_showMaybe.Do(func() {
		cache_Main_showMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2682098930_1386611502(Rebox_Main_1386611502_2682098930(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Maybe_showMaybe(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main_showMaybe
}

var cache_Main_worker gopurs_runtime.Value
var once_Main_worker sync.Once

func Get_Main_worker() gopurs_runtime.Value {
	once_Main_worker.Do(func() {
		cache_Main_worker = func() gopurs_runtime.Value {
			traverseEither_arg_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())
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
	return cache_Main_worker
}

var cache_Main_track gopurs_runtime.Value
var once_Main_track sync.Once

func Get_Main_track() gopurs_runtime.Value {
	once_Main_track.Do(func() {
		cache_Main_track = gopurs_runtime.Func2(func(calls_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_track(calls_0_box, value_1_box.IntVal))
		})
	})
	return cache_Main_track
}

var cache_Main_run gopurs_runtime.Value
var once_Main_run sync.Once

func Get_Main_run() gopurs_runtime.Value {
	once_Main_run.Do(func() {
		cache_Main_run = gopurs_runtime.Func2(func(calls_0_box gopurs_runtime.Value, values_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_run(calls_0_box, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(values_1_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Main_run
}

var cache_Main_produced gopurs_runtime.Value
var once_Main_produced sync.Once

func Get_Main_produced() gopurs_runtime.Value {
	once_Main_produced.Do(func() {
		cache_Main_produced = gopurs_runtime.Func(func(calls_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_produced(calls_0_box)
		})
	})
	return cache_Main_produced
}

var cache_Main_input gopurs_runtime.Value
var once_Main_input sync.Once

func Get_Main_input() gopurs_runtime.Value {
	once_Main_input.Do(func() {
		cache_Main_input = gopurs_runtime.Func(func(calls_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_input(calls_0_box)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_input
}

var cache_Main_generic gopurs_runtime.Value
var once_Main_generic sync.Once

func Get_Main_generic() gopurs_runtime.Value {
	once_Main_generic.Do(func() {
		cache_Main_generic = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_generic(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Main_generic
}

var cache_Main_generic__3263125282 gopurs_runtime.Value
var once_Main_generic__3263125282 sync.Once

func Get_Main_generic__3263125282() gopurs_runtime.Value {
	once_Main_generic__3263125282.Do(func() {
		cache_Main_generic__3263125282 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_generic__3263125282(__eta_norm_1_0_box, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
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
	return cache_Main_generic__3263125282
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___748768337("", struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
			}{func() gopurs_runtime.Value {
				_v := Call_Main_run(__local_var_1_1, []int64{int64(3), int64(7), int64(2), int64(9)})
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}(), func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, func() gopurs_runtime.Value {
					arr := []struct {
						value int64
					}{struct {
						value int64
					}{int64(6)}, struct {
						value int64
					}{int64(14)}, struct {
						value int64
					}{int64(4)}, struct {
						value int64
					}{int64(18)}}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
						}()
					}
					return gopurs_runtime.Array(boxed)
				}(), true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_2 shape=App(Var) bindingType=Any
				__local_var_3_2 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
				_ = __local_var_3_2
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_4_3 := gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Value{})
					_ = __local_var_4_3
					return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
						actual   []int64
						expected []int64
					}{func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(__local_var_4_3.UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}(), []int64{int64(3), int64(7), int64(2), int64(9)}}), gopurs_runtime.Value{})
				}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
						arr := []int64{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___748768337("", struct {
							actual   gopurs_runtime.Value
							expected gopurs_runtime.Value
						}{func() gopurs_runtime.Value {
							_v := Call_Main_run(__local_var_1_1, []int64{int64(3), int64(-1), int64(8), int64(-2), int64(9)})
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}(), func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Str("-1"), gopurs_runtime.Value{}, false}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): __local_var_7_4 shape=App(Var) bindingType=Any
							__local_var_7_4 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
							_ = __local_var_7_4
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_8_5 := gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Value{})
								_ = __local_var_8_5
								return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
									actual   []int64
									expected []int64
								}{func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(__local_var_8_5.UnsafePtr)
									unboxed := make([]int64, len(arr))
									for i, v := range arr {
										unboxed[i] = v.IntVal
									}
									return unboxed
								}(), []int64{int64(3), int64(-1), int64(8), int64(-2), int64(9)}}), gopurs_runtime.Value{})
							}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
									arr := []int64{}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___748768337("", struct {
										actual   gopurs_runtime.Value
										expected gopurs_runtime.Value
									}{func() gopurs_runtime.Value {
										_v := Call_Main_run(__local_var_1_1, []int64{})
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}(), func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 gopurs_runtime.Value
											V2 bool
										}{gopurs_runtime.Value{}, func() gopurs_runtime.Value {
											arr := []struct {
												value int64
											}{}
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = func() gopurs_runtime.Value {
													orig := v
													_ = orig
													return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
												}()
											}
											return gopurs_runtime.Array(boxed)
										}(), true}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}()}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_11_6 shape=App(Var) bindingType=Any
										__local_var_11_6 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
										_ = __local_var_11_6
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											__local_var_12_7 := gopurs_runtime.Apply(__local_var_11_6, gopurs_runtime.Value{})
											_ = __local_var_12_7
											return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
												actual   []int64
												expected []int64
											}{func() []int64 {
												arr := *(*[]gopurs_runtime.Value)(__local_var_12_7.UnsafePtr)
												unboxed := make([]int64, len(arr))
												for i, v := range arr {
													unboxed[i] = v.IntVal
												}
												return unboxed
											}(), []int64{}}), gopurs_runtime.Value{})
										}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												traverseEither_arg_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())
												_ = traverseEither_arg_9
												_ = traverseEither_arg_9
												traverseEither_arg_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither())
												_ = traverseEither_arg_10
												_ = traverseEither_arg_10
												traverseEither_arg_11 := Call_Main_produced(__local_var_1_1)
												_ = traverseEither_arg_11
												_ = traverseEither_arg_11
												__local_var_13_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func(func(traverseEither_input_12 gopurs_runtime.Value) gopurs_runtime.Value {
													traverseEither_values_12 := *((*[]gopurs_runtime.Value)(traverseEither_input_12.UnsafePtr))
													_ = traverseEither_values_12
													traverseEither_output_12 := make([]gopurs_runtime.Value, len(traverseEither_values_12))
													_ = traverseEither_output_12
													traverseEither_firstError_12 := gopurs_runtime.Value{}
													_ = traverseEither_firstError_12
													traverseEither_failed_12 := false
													_ = traverseEither_failed_12
													for traverseEither_index_12, traverseEither_value_12 := range traverseEither_values_12 {
														traverseEither_result_12 := gopurs_runtime.Apply(traverseEither_arg_11, traverseEither_value_12)
														_ = traverseEither_result_12
														if traverseEither_result_12.Type == 9 && traverseEither_result_12.IntVal == 2465973597 {
															traverseEither_output_12[traverseEither_index_12] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseEither_result_12.UnsafePtr).V0
														} else {
															if traverseEither_result_12.Type == 9 && traverseEither_result_12.IntVal == 3711209382 {
																if !(traverseEither_failed_12) {
																	traverseEither_firstError_12 = traverseEither_result_12
																	traverseEither_failed_12 = true
																} else {

																}
															} else {
																panic("Failed pattern match")
															}
														}
													}
													if traverseEither_failed_12 {
														return traverseEither_firstError_12
													} else {

													}
													return func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 gopurs_runtime.Value
															V2 bool
														}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_12), true}
														if _v.V2 {
															return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
														}
														return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
													}()
												})), gopurs_runtime.Value{})
												_ = __local_var_13_8
												// TAST (Let): __local_var_14_13 shape=App(Var) bindingType=Any
												__local_var_14_13 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
												_ = __local_var_14_13
												return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_15_14 := gopurs_runtime.Apply(__local_var_14_13, gopurs_runtime.Value{})
													_ = __local_var_15_14
													return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
														actual   []int64
														expected []int64
													}{func() []int64 {
														arr := *(*[]gopurs_runtime.Value)(__local_var_15_14.UnsafePtr)
														unboxed := make([]int64, len(arr))
														for i, v := range arr {
															unboxed[i] = v.IntVal
														}
														return unboxed
													}(), []int64{int64(200)}}), gopurs_runtime.Value{})
												}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														// TAST (Let): __local_var_16_15 shape=App(Var) bindingType=Any
														__local_var_16_15 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_13_8)
														_ = __local_var_16_15
														__local_var_17_16 := gopurs_runtime.Apply(__local_var_16_15, gopurs_runtime.Value{})
														_ = __local_var_17_16
														return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___748768337("", struct {
															actual   gopurs_runtime.Value
															expected gopurs_runtime.Value
														}{gopurs_runtime.Apply(__local_var_17_16, func() gopurs_runtime.Value {
															arr := []int64{int64(6)}
															boxed := make([]gopurs_runtime.Value, len(arr))
															for i, v := range arr {
																boxed[i] = gopurs_runtime.Int(v)
															}
															return gopurs_runtime.Array(boxed)
														}()), func() gopurs_runtime.Value {
															_v := struct {
																V0 gopurs_runtime.Value
																V1 gopurs_runtime.Value
																V2 bool
															}{gopurs_runtime.Value{}, func() gopurs_runtime.Value {
																arr := []struct {
																	value int64
																}{struct {
																	value int64
																}{int64(6)}}
																boxed := make([]gopurs_runtime.Value, len(arr))
																for i, v := range arr {
																	boxed[i] = func() gopurs_runtime.Value {
																		orig := v
																		_ = orig
																		return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
																	}()
																}
																return gopurs_runtime.Array(boxed)
															}(), true}
															if _v.V2 {
																return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
															}
															return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
														}()}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___748768337("", struct {
																actual   gopurs_runtime.Value
																expected gopurs_runtime.Value
															}{gopurs_runtime.Apply(__local_var_17_16, func() gopurs_runtime.Value {
																arr := []int64{int64(7)}
																boxed := make([]gopurs_runtime.Value, len(arr))
																for i, v := range arr {
																	boxed[i] = gopurs_runtime.Int(v)
																}
																return gopurs_runtime.Array(boxed)
															}()), func() gopurs_runtime.Value {
																_v := struct {
																	V0 gopurs_runtime.Value
																	V1 gopurs_runtime.Value
																	V2 bool
																}{gopurs_runtime.Value{}, func() gopurs_runtime.Value {
																	arr := []struct {
																		value int64
																	}{struct {
																		value int64
																	}{int64(7)}}
																	boxed := make([]gopurs_runtime.Value, len(arr))
																	for i, v := range arr {
																		boxed[i] = func() gopurs_runtime.Value {
																			orig := v
																			_ = orig
																			return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
																		}()
																	}
																	return gopurs_runtime.Array(boxed)
																}(), true}
																if _v.V2 {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																}
																return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
															}()}), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
																// TAST (Let): __local_var_20_17 shape=App(Var) bindingType=Any
																__local_var_20_17 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
																_ = __local_var_20_17
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_21_18 := gopurs_runtime.Apply(__local_var_20_17, gopurs_runtime.Value{})
																	_ = __local_var_21_18
																	return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
																		actual   []int64
																		expected []int64
																	}{func() []int64 {
																		arr := *(*[]gopurs_runtime.Value)(__local_var_21_18.UnsafePtr)
																		unboxed := make([]int64, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v.IntVal
																		}
																		return unboxed
																	}(), []int64{int64(200), int64(6), int64(7)}}), gopurs_runtime.Value{})
																}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
																		arr := []int64{}
																		boxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			boxed[i] = gopurs_runtime.Int(v)
																		}
																		return gopurs_runtime.Array(boxed)
																	}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
																		traverseEither_arg_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())
																		_ = traverseEither_arg_19
																		_ = traverseEither_arg_19
																		traverseEither_arg_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither())
																		_ = traverseEither_arg_20
																		_ = traverseEither_arg_20
																		traverseEither_arg_21 := Call_Main_produced(__local_var_1_1)
																		_ = traverseEither_arg_21
																		_ = traverseEither_arg_21
																		traverseEither_arg_22 := Call_Main_input(__local_var_1_1)
																		_ = traverseEither_arg_22
																		_ = traverseEither_arg_22
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___748768337("", struct {
																			actual   gopurs_runtime.Value
																			expected gopurs_runtime.Value
																		}{func() gopurs_runtime.Value {
																			traverseEither_values_23 := traverseEither_arg_22
																			_ = traverseEither_values_23
																			traverseEither_output_23 := make([]gopurs_runtime.Value, len(traverseEither_values_23))
																			_ = traverseEither_output_23
																			traverseEither_firstError_23 := gopurs_runtime.Value{}
																			_ = traverseEither_firstError_23
																			traverseEither_failed_23 := false
																			_ = traverseEither_failed_23
																			for traverseEither_index_23, traverseEither_value_23 := range traverseEither_values_23 {
																				traverseEither_result_23 := gopurs_runtime.Apply(traverseEither_arg_21, gopurs_runtime.Int(traverseEither_value_23))
																				_ = traverseEither_result_23
																				if traverseEither_result_23.Type == 9 && traverseEither_result_23.IntVal == 2465973597 {
																					traverseEither_output_23[traverseEither_index_23] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseEither_result_23.UnsafePtr).V0
																				} else {
																					if traverseEither_result_23.Type == 9 && traverseEither_result_23.IntVal == 3711209382 {
																						if !(traverseEither_failed_23) {
																							traverseEither_firstError_23 = traverseEither_result_23
																							traverseEither_failed_23 = true
																						} else {

																						}
																					} else {
																						panic("Failed pattern match")
																					}
																				}
																			}
																			if traverseEither_failed_23 {
																				return traverseEither_firstError_23
																			} else {

																			}
																			return func() gopurs_runtime.Value {
																				_v := struct {
																					V0 gopurs_runtime.Value
																					V1 gopurs_runtime.Value
																					V2 bool
																				}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_23), true}
																				if _v.V2 {
																					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																				}
																				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																			}()
																		}(), func() gopurs_runtime.Value {
																			_v := struct {
																				V0 gopurs_runtime.Value
																				V1 gopurs_runtime.Value
																				V2 bool
																			}{gopurs_runtime.Value{}, func() gopurs_runtime.Value {
																				arr := []struct {
																					value int64
																				}{struct {
																					value int64
																				}{int64(4)}, struct {
																					value int64
																				}{int64(5)}}
																				boxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					boxed[i] = func() gopurs_runtime.Value {
																						orig := v
																						_ = orig
																						return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
																					}()
																				}
																				return gopurs_runtime.Array(boxed)
																			}(), true}
																			if _v.V2 {
																				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																			}
																			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																		}()}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																			// TAST (Let): __local_var_24_24 shape=App(Var) bindingType=Any
																			__local_var_24_24 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
																			_ = __local_var_24_24
																			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																				__local_var_25_25 := gopurs_runtime.Apply(__local_var_24_24, gopurs_runtime.Value{})
																				_ = __local_var_25_25
																				return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
																					actual   []int64
																					expected []int64
																				}{func() []int64 {
																					arr := *(*[]gopurs_runtime.Value)(__local_var_25_25.UnsafePtr)
																					unboxed := make([]int64, len(arr))
																					for i, v := range arr {
																						unboxed[i] = v.IntVal
																					}
																					return unboxed
																				}(), []int64{int64(200), int64(300), int64(4), int64(5)}}), gopurs_runtime.Value{})
																			}), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																					// TAST (Let): __local_var_26_26 shape=App(Var) bindingType=Any
																					__local_var_26_26 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Main_worker())
																					_ = __local_var_26_26
																					savedWorker_27_27 := gopurs_runtime.Apply(__local_var_26_26, gopurs_runtime.Value{})
																					_ = savedWorker_27_27
																					__local_var_28_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), savedWorker_27_27), gopurs_runtime.Value{})
																					_ = __local_var_28_28
																					return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___748768337("", struct {
																						actual   gopurs_runtime.Value
																						expected gopurs_runtime.Value
																					}{gopurs_runtime.Apply2(__local_var_28_28, gopurs_runtime.Func(func(value_29 gopurs_runtime.Value) gopurs_runtime.Value {
																						return func() gopurs_runtime.Value {
																							_v := struct {
																								V0 gopurs_runtime.Value
																								V1 struct {
																									value int64
																								}
																								V2 bool
																							}{gopurs_runtime.Value{}, struct {
																								value int64
																							}{value_29.IntVal}, true}
																							if _v.V2 {
																								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																									orig := _v.V1
																									_ = orig
																									return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
																								}()})}
																							}
																							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																						}()
																					}), func() gopurs_runtime.Value {
																						arr := []int64{int64(10)}
																						boxed := make([]gopurs_runtime.Value, len(arr))
																						for i, v := range arr {
																							boxed[i] = gopurs_runtime.Int(v)
																						}
																						return gopurs_runtime.Array(boxed)
																					}()), func() gopurs_runtime.Value {
																						_v := struct {
																							V0 gopurs_runtime.Value
																							V1 gopurs_runtime.Value
																							V2 bool
																						}{gopurs_runtime.Value{}, func() gopurs_runtime.Value {
																							arr := []struct {
																								value int64
																							}{struct {
																								value int64
																							}{int64(10)}}
																							boxed := make([]gopurs_runtime.Value, len(arr))
																							for i, v := range arr {
																								boxed[i] = func() gopurs_runtime.Value {
																									orig := v
																									_ = orig
																									return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
																								}()
																							}
																							return gopurs_runtime.Array(boxed)
																						}(), true}
																						if _v.V2 {
																							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																						}
																						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																					}()}), gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___283401329("", struct {
																							actual   *Constructor_Data_Maybe_Just[[]int64]
																							expected *Constructor_Data_Maybe_Just[[]int64]
																						}{Rebox_Main_3094389156_1495236409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_649684152_1439734649(Rebox_Main_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe()))))}, gopurs_runtime.Func(func(value_30 gopurs_runtime.Value) gopurs_runtime.Value {
																							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
																								_v := struct {
																									V0 gopurs_runtime.Value
																									V1 bool
																								}{gopurs_runtime.Int((value_30.IntVal) + (int64(1))), true}
																								if _v.V1 {
																									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
																								}
																								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
																							}()))))}
																						}), func() gopurs_runtime.Value {
																							arr := []int64{int64(1), int64(2)}
																							boxed := make([]gopurs_runtime.Value, len(arr))
																							for i, v := range arr {
																								boxed[i] = gopurs_runtime.Int(v)
																							}
																							return gopurs_runtime.Array(boxed)
																						}()))), Rebox_Main_3094389156_1495236409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
																							_v := struct {
																								V0 gopurs_runtime.Value
																								V1 bool
																							}{func() gopurs_runtime.Value {
																								arr := []int64{int64(2), int64(3)}
																								boxed := make([]gopurs_runtime.Value, len(arr))
																								for i, v := range arr {
																									boxed[i] = gopurs_runtime.Int(v)
																								}
																								return gopurs_runtime.Array(boxed)
																							}(), true}
																							if _v.V1 {
																								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
																							}
																							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
																						}()))}), gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___283401329("", struct {
																								actual   *Constructor_Data_Maybe_Just[[]int64]
																								expected *Constructor_Data_Maybe_Just[[]int64]
																							}{Rebox_Main_3094389156_1495236409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_649684152_1439734649(Rebox_Main_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe()))))}, gopurs_runtime.Func(func(value_31 gopurs_runtime.Value) gopurs_runtime.Value {
																								var __t29 *Constructor_Data_Maybe_Just[int64]
																								{
																									if (value_31.IntVal) < (int64(0)) {
																										__t29 = Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
																											_v := struct {
																												V0 gopurs_runtime.Value
																												V1 bool
																											}{gopurs_runtime.Value{}, false}
																											if _v.V1 {
																												return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
																											}
																											return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
																										}()))
																										goto end_branch_29
																									} else {

																									}
																								}
																								{
																									__t29 = Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
																										_v := struct {
																											V0 gopurs_runtime.Value
																											V1 bool
																										}{gopurs_runtime.Int(value_31.IntVal), true}
																										if _v.V1 {
																											return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
																										}
																										return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
																									}()))
																								}
																							end_branch_29:
																								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(__t29))}
																							}), func() gopurs_runtime.Value {
																								arr := []int64{int64(1), int64(-1), int64(2)}
																								boxed := make([]gopurs_runtime.Value, len(arr))
																								for i, v := range arr {
																									boxed[i] = gopurs_runtime.Int(v)
																								}
																								return gopurs_runtime.Array(boxed)
																							}()))), Rebox_Main_3094389156_1495236409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}), gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																								return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																							}))
																						}))
																					})), gopurs_runtime.Value{})
																				})
																			}))
																		}))
																	}))
																}))
															}))
														})), gopurs_runtime.Value{})
													})
												})), gopurs_runtime.Value{})
											})
										}))
									}))
								}))
							}))
						}))
					}))
				}))
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_track(calls_0_loop gopurs_runtime.Value, value_1_loop int64) int64 {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var value_1 int64 = value_1_loop
	_ = value_1
	return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), seen_2, func() gopurs_runtime.Value {
					arr := []int64{value_1}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	}), calls_0), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(value_1)
		})
	}))).IntVal
}

func Call_Main_run(calls_0_loop gopurs_runtime.Value, values_1_loop []int64) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var values_1 []int64 = values_1_loop
	_ = values_1
	traverseEither_arg_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())
	_ = traverseEither_arg_0
	_ = traverseEither_arg_0
	traverseEither_arg_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither())
	_ = traverseEither_arg_1
	_ = traverseEither_arg_1
	traverseEither_arg_4 := func(value_2 gopurs_runtime.Value) struct {
		V0 gopurs_runtime.Value
		V1 struct {
			value int64
		}
		V2 bool
	} {
		// TAST (Let): observed_3_2 shape=App(Var) bindingType=Int
		observed_3_2 := Call_Main_track(calls_0, value_2.IntVal)
		_ = observed_3_2
		var __t3 struct {
			V0 gopurs_runtime.Value
			V1 struct {
				value int64
			}
			V2 bool
		}
		{
			if (observed_3_2) < (int64(0)) {
				__t3 = struct {
					V0 gopurs_runtime.Value
					V1 struct {
						value int64
					}
					V2 bool
				}{gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(observed_3_2)).StrVal()), struct {
					value int64
				}{}, false}
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = struct {
				V0 gopurs_runtime.Value
				V1 struct {
					value int64
				}
				V2 bool
			}{gopurs_runtime.Value{}, struct {
				value int64
			}{(observed_3_2) * (int64(2))}, true}
		}
	end_branch_3:
		return __t3
	}
	_ = traverseEither_arg_4
	_ = traverseEither_arg_4
	traverseEither_arg_5 := values_1
	_ = traverseEither_arg_5
	_ = traverseEither_arg_5
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := func() gopurs_runtime.Value {
		traverseEither_values_6 := traverseEither_arg_5
		_ = traverseEither_values_6
		traverseEither_output_6 := make([]gopurs_runtime.Value, len(traverseEither_values_6))
		_ = traverseEither_output_6
		traverseEither_firstError_6 := gopurs_runtime.Value{}
		_ = traverseEither_firstError_6
		traverseEither_failed_6 := false
		_ = traverseEither_failed_6
		for traverseEither_index_6, traverseEither_value_6 := range traverseEither_values_6 {
			traverseEither_result_6 := traverseEither_arg_4(gopurs_runtime.Int(traverseEither_value_6))
			_ = traverseEither_result_6
			if traverseEither_result_6.V2 {
				traverseEither_output_6[traverseEither_index_6] = func() gopurs_runtime.Value {
					orig := traverseEither_result_6.V1
					_ = orig
					return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
				}()
			} else {
				if !traverseEither_result_6.V2 {
					if !(traverseEither_failed_6) {
						traverseEither_firstError_6 = func() gopurs_runtime.Value {
							_v := traverseEither_result_6
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
									orig := _v.V1
									_ = orig
									return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
								}()})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()
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
	}(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Main_produced(calls_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), seen_1, func() gopurs_runtime.Value {
					arr := []int64{int64(200)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	}), calls_0), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 struct {
							value int64
						}
						V2 bool
					}{gopurs_runtime.Value{}, struct {
						value int64
					}{Call_Main_track(calls_0, value_2.IntVal)}, true}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
							orig := _v.V1
							_ = orig
							return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
						}()})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
			})
		})
	})))
}

func Call_Main_input(calls_0_loop gopurs_runtime.Value) []int64 {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), seen_1, func() gopurs_runtime.Value {
						arr := []int64{int64(300)}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())).UnsafePtr))).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		}), calls_0), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := []int64{int64(4), int64(5)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			})
		}))).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Main_generic(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	return gopurs_runtime.Apply(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
}

func Call_Main_generic__3263125282(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop []int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
generic__3263125282:
	for {
		if false {
			continue generic__3263125282
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 []int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1495236409_3094389156(Rebox_Main_3094389156_1495236409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_649684152_1439734649(Rebox_Main_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe()))))}, __eta_norm_1_0, func() gopurs_runtime.Value {
			arr := __eta_norm_0_1
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2682098930(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[[]int64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[[]int64]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1439734649_649684152(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1495236409_3094389156(in *Constructor_Data_Maybe_Just[[]int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		arr := in.V0
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Int(v)
		}
		return gopurs_runtime.Array(boxed)
	}()
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1888010770_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[[]int64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2682098930_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[[]int64]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Main_3094389156_1495236409(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[[]int64]{}
	out.V0 = func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(in.V0.UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
	return out
}

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1888010770(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[[]int64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[[]int64]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_649684152_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}
