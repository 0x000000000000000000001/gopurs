package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_indexIsSymbol gopurs_runtime.Value
var once_Main_indexIsSymbol sync.Once

func Get_Main_indexIsSymbol() gopurs_runtime.Value {
	once_Main_indexIsSymbol.Do(func() {
		cache_Main_indexIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("index")
		}))
	})
	return cache_Main_indexIsSymbol
}

var cache_Main_eqEither gopurs_runtime.Value
var once_Main_eqEither sync.Once

func Get_Main_eqEither() gopurs_runtime.Value {
	once_Main_eqEither.Do(func() {
		cache_Main_eqEither = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_indexIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main_eqEither
}

var cache_Main_showEither gopurs_runtime.Value
var once_Main_showEither sync.Once

func Get_Main_showEither() gopurs_runtime.Value {
	once_Main_showEither.Do(func() {
		cache_Main_showEither = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}, Call_Data_Show_showArray(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_indexIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
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

var cache_Main_track gopurs_runtime.Value
var once_Main_track sync.Once

func Get_Main_track() gopurs_runtime.Value {
	once_Main_track.Do(func() {
		cache_Main_track = gopurs_runtime.Func3(func(calls_0_box gopurs_runtime.Value, marker_1_box gopurs_runtime.Value, value_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_track(calls_0_box, marker_1_box.IntVal, value_2_box.IntVal))
		})
	})
	return cache_Main_track
}

var cache_Main_staged gopurs_runtime.Value
var once_Main_staged sync.Once

func Get_Main_staged() gopurs_runtime.Value {
	once_Main_staged.Do(func() {
		cache_Main_staged = gopurs_runtime.Func2(func(calls_0_box gopurs_runtime.Value, values_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_staged(calls_0_box, func() []int64 {
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
	return cache_Main_staged
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

var cache_Main_partial gopurs_runtime.Value
var once_Main_partial sync.Once

func Get_Main_partial() gopurs_runtime.Value {
	once_Main_partial.Do(func() {
		cache_Main_partial = gopurs_runtime.Func2(func(calls_0_box gopurs_runtime.Value, offset_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_partial(calls_0_box, offset_1_box.IntVal)
		})
	})
	return cache_Main_partial
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
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1693432849("", struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
			}{func() gopurs_runtime.Value {
				_v := Call_Main_run(__local_var_1_1, []int64{int64(3), int64(7)})
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
						index int64
						value int64
					}{struct {
						index int64
						value int64
					}{int64(0), int64(3)}, struct {
						index int64
						value int64
					}{int64(1), int64(7)}}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
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
					}(), []int64{int64(0), int64(1)}}), gopurs_runtime.Value{})
				}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
						arr := []int64{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1693432849("", struct {
							actual   gopurs_runtime.Value
							expected gopurs_runtime.Value
						}{func() gopurs_runtime.Value {
							_v := Call_Main_run(__local_var_1_1, []int64{int64(-1), int64(8), int64(-2)})
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}(), func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Str("0"), gopurs_runtime.Value{}, false}
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
								}(), []int64{int64(0), int64(1), int64(2)}}), gopurs_runtime.Value{})
							}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
									arr := []int64{}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1693432849("", struct {
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
												index int64
												value int64
											}{}
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = func() gopurs_runtime.Value {
													orig := v
													_ = orig
													return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
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
												__local_var_13_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_partial(__local_var_1_1, int64(10))), gopurs_runtime.Value{})
												_ = __local_var_13_8
												// TAST (Let): __local_var_14_9 shape=App(Var) bindingType=Any
												__local_var_14_9 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
												_ = __local_var_14_9
												return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_15_10 := gopurs_runtime.Apply(__local_var_14_9, gopurs_runtime.Value{})
													_ = __local_var_15_10
													return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
														actual   []int64
														expected []int64
													}{func() []int64 {
														arr := *(*[]gopurs_runtime.Value)(__local_var_15_10.UnsafePtr)
														unboxed := make([]int64, len(arr))
														for i, v := range arr {
															unboxed[i] = v.IntVal
														}
														return unboxed
													}(), []int64{int64(100)}}), gopurs_runtime.Value{})
												}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														// TAST (Let): __local_var_16_11 shape=App(Var) bindingType=Any
														__local_var_16_11 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_13_8)
														_ = __local_var_16_11
														__local_var_17_12 := gopurs_runtime.Apply(__local_var_16_11, gopurs_runtime.Value{})
														_ = __local_var_17_12
														return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1693432849("", struct {
															actual   gopurs_runtime.Value
															expected gopurs_runtime.Value
														}{gopurs_runtime.Apply(__local_var_17_12, func() gopurs_runtime.Value {
															arr := []int64{int64(3)}
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
																	index int64
																	value int64
																}{struct {
																	index int64
																	value int64
																}{int64(0), int64(13)}}
																boxed := make([]gopurs_runtime.Value, len(arr))
																for i, v := range arr {
																	boxed[i] = func() gopurs_runtime.Value {
																		orig := v
																		_ = orig
																		return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
																	}()
																}
																return gopurs_runtime.Array(boxed)
															}(), true}
															if _v.V2 {
																return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
															}
															return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
														}()}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1693432849("", struct {
																actual   gopurs_runtime.Value
																expected gopurs_runtime.Value
															}{gopurs_runtime.Apply(__local_var_17_12, func() gopurs_runtime.Value {
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
																		index int64
																		value int64
																	}{struct {
																		index int64
																		value int64
																	}{int64(0), int64(17)}}
																	boxed := make([]gopurs_runtime.Value, len(arr))
																	for i, v := range arr {
																		boxed[i] = func() gopurs_runtime.Value {
																			orig := v
																			_ = orig
																			return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
																		}()
																	}
																	return gopurs_runtime.Array(boxed)
																}(), true}
																if _v.V2 {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																}
																return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
															}()}), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
																// TAST (Let): __local_var_20_13 shape=App(Var) bindingType=Any
																__local_var_20_13 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
																_ = __local_var_20_13
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_21_14 := gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Value{})
																	_ = __local_var_21_14
																	return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
																		actual   []int64
																		expected []int64
																	}{func() []int64 {
																		arr := *(*[]gopurs_runtime.Value)(__local_var_21_14.UnsafePtr)
																		unboxed := make([]int64, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v.IntVal
																		}
																		return unboxed
																	}(), []int64{int64(100)}}), gopurs_runtime.Value{})
																}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
																		arr := []int64{}
																		boxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			boxed[i] = gopurs_runtime.Int(v)
																		}
																		return gopurs_runtime.Array(boxed)
																	}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1693432849("", struct {
																			actual   gopurs_runtime.Value
																			expected gopurs_runtime.Value
																		}{func() gopurs_runtime.Value {
																			_v := Call_Main_staged(__local_var_1_1, []int64{int64(3), int64(-1), int64(-2)})
																			if _v.V2 {
																				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																			}
																			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																		}(), func() gopurs_runtime.Value {
																			_v := struct {
																				V0 gopurs_runtime.Value
																				V1 gopurs_runtime.Value
																				V2 bool
																			}{gopurs_runtime.Str("1"), gopurs_runtime.Value{}, false}
																			if _v.V2 {
																				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																			}
																			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																		}()}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																			// TAST (Let): __local_var_24_15 shape=App(Var) bindingType=Any
																			__local_var_24_15 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
																			_ = __local_var_24_15
																			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																				__local_var_25_16 := gopurs_runtime.Apply(__local_var_24_15, gopurs_runtime.Value{})
																				_ = __local_var_25_16
																				return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
																					actual   []int64
																					expected []int64
																				}{func() []int64 {
																					arr := *(*[]gopurs_runtime.Value)(__local_var_25_16.UnsafePtr)
																					unboxed := make([]int64, len(arr))
																					for i, v := range arr {
																						unboxed[i] = v.IntVal
																					}
																					return unboxed
																				}(), []int64{int64(100), int64(101), int64(102)}}), gopurs_runtime.Value{})
																			}), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
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

func Call_Main_track(calls_0_loop gopurs_runtime.Value, marker_1_loop int64, value_2_loop int64) int64 {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var marker_1 int64 = marker_1_loop
	_ = marker_1
	var value_2 int64 = value_2_loop
	_ = value_2
	return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), seen_3, func() gopurs_runtime.Value {
					arr := []int64{marker_1}
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
	}), calls_0), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(value_2)
		})
	}))).IntVal
}

func Call_Main_staged(calls_0_loop gopurs_runtime.Value, values_1_loop []int64) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var values_1 []int64 = values_1_loop
	_ = values_1
	traverseEither_arg_0 := Rebox_Main_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))
	_ = traverseEither_arg_0
	_ = traverseEither_arg_0
	traverseEither_arg_1 := Get_Data_Either_applicativeEither()
	_ = traverseEither_arg_1
	_ = traverseEither_arg_1
	traverseEither_arg_4 := gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): captured_3_2 shape=App(Var) bindingType=Int
		captured_3_2 := Call_Main_track(calls_0, (int64(100))+(i_2.IntVal), i_2.IntVal)
		_ = captured_3_2
		return gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 struct {
				V0 gopurs_runtime.Value
				V1 struct {
					index int64
					value int64
				}
				V2 bool
			}
			{
				if (value_4.IntVal) < (int64(0)) {
					__t3 = struct {
						V0 gopurs_runtime.Value
						V1 struct {
							index int64
							value int64
						}
						V2 bool
					}{gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(captured_3_2)).StrVal()), struct {
						index int64
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
						index int64
						value int64
					}
					V2 bool
				}{gopurs_runtime.Value{}, struct {
					index int64
					value int64
				}{captured_3_2, value_4.IntVal}, true}
			}
		end_branch_3:
			return func() gopurs_runtime.Value {
				_v := __t3
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
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
			traverseEither_result_6 := gopurs_runtime.Apply2(traverseEither_arg_4, gopurs_runtime.Int(int64(traverseEither_index_6)), gopurs_runtime.Int(traverseEither_value_6))
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

func Call_Main_run(calls_0_loop gopurs_runtime.Value, values_1_loop []int64) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var values_1 []int64 = values_1_loop
	_ = values_1
	traverseEither_arg_0 := Rebox_Main_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))
	_ = traverseEither_arg_0
	_ = traverseEither_arg_0
	traverseEither_arg_1 := Get_Data_Either_applicativeEither()
	_ = traverseEither_arg_1
	_ = traverseEither_arg_1
	traverseEither_arg_4 := func(i_2 int64, value_3 gopurs_runtime.Value) struct {
		V0 gopurs_runtime.Value
		V1 struct {
			index int64
			value int64
		}
		V2 bool
	} {
		// TAST (Let): observed_4_2 shape=App(Var) bindingType=Int
		observed_4_2 := Call_Main_track(calls_0, i_2, value_3.IntVal)
		_ = observed_4_2
		var __t3 struct {
			V0 gopurs_runtime.Value
			V1 struct {
				index int64
				value int64
			}
			V2 bool
		}
		{
			if (observed_4_2) < (int64(0)) {
				__t3 = struct {
					V0 gopurs_runtime.Value
					V1 struct {
						index int64
						value int64
					}
					V2 bool
				}{gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(i_2)).StrVal()), struct {
					index int64
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
					index int64
					value int64
				}
				V2 bool
			}{gopurs_runtime.Value{}, struct {
				index int64
				value int64
			}{i_2, observed_4_2}, true}
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
			traverseEither_result_6 := traverseEither_arg_4(int64(traverseEither_index_6), gopurs_runtime.Int(traverseEither_value_6))
			_ = traverseEither_result_6
			if traverseEither_result_6.V2 {
				traverseEither_output_6[traverseEither_index_6] = func() gopurs_runtime.Value {
					orig := traverseEither_result_6.V1
					_ = orig
					return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
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
									return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
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

func Call_Main_partial(calls_0_loop gopurs_runtime.Value, offset_1_loop int64) gopurs_runtime.Value {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var offset_1 int64 = offset_1_loop
	_ = offset_1
	// TAST (Let): captured_2_0 shape=App(Var) bindingType=Int
	captured_2_0 := Call_Main_track(calls_0, int64(100), offset_1)
	_ = captured_2_0
	traverseEither_arg_1 := Rebox_Main_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))
	_ = traverseEither_arg_1
	_ = traverseEither_arg_1
	traverseEither_arg_2 := Get_Data_Either_applicativeEither()
	_ = traverseEither_arg_2
	_ = traverseEither_arg_2
	traverseEither_arg_3 := func(i_3 int64, value_4 gopurs_runtime.Value) struct {
		V0 gopurs_runtime.Value
		V1 struct {
			index int64
			value int64
		}
		V2 bool
	} {
		return struct {
			V0 gopurs_runtime.Value
			V1 struct {
				index int64
				value int64
			}
			V2 bool
		}{gopurs_runtime.Value{}, struct {
			index int64
			value int64
		}{i_3, (value_4.IntVal) + (captured_2_0)}, true}
	}
	_ = traverseEither_arg_3
	_ = traverseEither_arg_3
	return gopurs_runtime.Func(func(traverseEither_input_4 gopurs_runtime.Value) gopurs_runtime.Value {
		traverseEither_values_4 := *((*[]gopurs_runtime.Value)(traverseEither_input_4.UnsafePtr))
		_ = traverseEither_values_4
		traverseEither_output_4 := make([]gopurs_runtime.Value, len(traverseEither_values_4))
		_ = traverseEither_output_4
		traverseEither_firstError_4 := gopurs_runtime.Value{}
		_ = traverseEither_firstError_4
		traverseEither_failed_4 := false
		_ = traverseEither_failed_4
		for traverseEither_index_4, traverseEither_value_4 := range traverseEither_values_4 {
			traverseEither_result_4 := traverseEither_arg_3(int64(traverseEither_index_4), traverseEither_value_4)
			_ = traverseEither_result_4
			if traverseEither_result_4.V2 {
				traverseEither_output_4[traverseEither_index_4] = func() gopurs_runtime.Value {
					orig := traverseEither_result_4.V1
					_ = orig
					return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
				}()
			} else {
				if !traverseEither_result_4.V2 {
					if !(traverseEither_failed_4) {
						traverseEither_firstError_4 = func() gopurs_runtime.Value {
							_v := traverseEither_result_4
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
									orig := _v.V1
									_ = orig
									return gopurs_runtime.RecordDict2("index", "value", gopurs_runtime.Int(orig.index), gopurs_runtime.Int(orig.value))
								}()})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()
						traverseEither_failed_4 = true
					} else {

					}
				} else {
					panic("Failed pattern match")
				}
			}
		}
		if traverseEither_failed_4 {
			return traverseEither_firstError_4
		} else {

		}
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Array(traverseEither_output_4), true}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	})
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

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
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

func Rebox_Main_1812164904_2955889203(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	out.V3 = in.V3
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

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]int64]{}
	out.V0 = in.V0
	return out
}
