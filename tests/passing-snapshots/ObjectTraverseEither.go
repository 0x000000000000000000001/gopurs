package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sort "sort"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_eqEither gopurs_runtime.Value
var once_Main_eqEither sync.Once

func Get_Main_eqEither() gopurs_runtime.Value {
	once_Main_eqEither.Do(func() {
		cache_Main_eqEither = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}, Call_Foreign_Object_eqObject(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))}
	})
	return cache_Main_eqEither
}

var cache_Main_showEither gopurs_runtime.Value
var once_Main_showEither sync.Once

func Get_Main_showEither() gopurs_runtime.Value {
	once_Main_showEither.Do(func() {
		cache_Main_showEither = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}, Call_Foreign_Object_showObject(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))}
	})
	return cache_Main_showEither
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_131790935_3790796878(Rebox_Main_3790796878_131790935(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1953100407_1386611502(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_eqMaybe gopurs_runtime.Value
var once_Main_eqMaybe sync.Once

func Get_Main_eqMaybe() gopurs_runtime.Value {
	once_Main_eqMaybe.Do(func() {
		cache_Main_eqMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3960201844_3790796878(Rebox_Main_3790796878_3960201844(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqMaybe
}

var cache_Main_showMaybe gopurs_runtime.Value
var once_Main_showMaybe sync.Once

func Get_Main_showMaybe() gopurs_runtime.Value {
	once_Main_showMaybe.Do(func() {
		cache_Main_showMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2818770644_1386611502(Rebox_Main_1386611502_2818770644(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showMaybe
}

var cache_Main_track gopurs_runtime.Value
var once_Main_track sync.Once

func Get_Main_track() gopurs_runtime.Value {
	once_Main_track.Do(func() {
		cache_Main_track = gopurs_runtime.Func3(func(calls_0_box gopurs_runtime.Value, key_1_box gopurs_runtime.Value, value_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_track(calls_0_box, key_1_box.StrVal(), value_2_box.IntVal))
		})
	})
	return cache_Main_track
}

var cache_Main_run gopurs_runtime.Value
var once_Main_run sync.Once

func Get_Main_run() gopurs_runtime.Value {
	once_Main_run.Do(func() {
		cache_Main_run = gopurs_runtime.Func(func(calls_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_run(calls_0_box)
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
				arr := []string{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}())
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			// TAST (Let): input_2_2 shape=App(Var) bindingType=(ADT ["Foreign","Object","Object"] [Int])
			input_2_2 := Call_Foreign_Object_fromFoldable__1025470229(func() gopurs_runtime.Value {
				arr := []*Constructor_Data_Tuple_Tuple[string, int64]{Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Str("a"), gopurs_runtime.Int(int64(3))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}())), Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Str("b"), gopurs_runtime.Int(int64(7))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1830029836_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}())
			_ = input_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4024920945("", struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
			}{gopurs_runtime.Apply(Call_Main_run(__local_var_1_1), input_2_2), func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Value{}, Call_Foreign_Object_fromFoldable__1025470229(func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple[string, int64]{Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("a"), gopurs_runtime.Int(int64(6))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}())), Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Str("b"), gopurs_runtime.Int(int64(14))}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1830029836_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}()), true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_3 shape=App(Var) bindingType=Any
				__local_var_4_3 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
				_ = __local_var_4_3
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_5_4 := gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Value{})
					_ = __local_var_5_4
					return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2499841393("", struct {
						actual   []string
						expected []string
					}{func() []string {
						arr := *(*[]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}(), []string{"a", "b"}}), gopurs_runtime.Value{})
				}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
						arr := []string{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4024920945("", struct {
							actual   gopurs_runtime.Value
							expected gopurs_runtime.Value
						}{gopurs_runtime.Apply(Call_Main_run(__local_var_1_1), Call_Foreign_Object_fromFoldable__1025470229(func() gopurs_runtime.Value {
							arr := []*Constructor_Data_Tuple_Tuple[string, int64]{Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{gopurs_runtime.Str("a"), gopurs_runtime.Int(int64(-1))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}())), Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{gopurs_runtime.Str("b"), gopurs_runtime.Int(int64(8))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}())), Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{gopurs_runtime.Str("c"), gopurs_runtime.Int(int64(-2))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}()))}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1830029836_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}())), func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Str("a"), gopurs_runtime.Value{}, false}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): __local_var_8_5 shape=App(Var) bindingType=Any
							__local_var_8_5 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
							_ = __local_var_8_5
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_9_6 := gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Value{})
								_ = __local_var_9_6
								return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2499841393("", struct {
									actual   []string
									expected []string
								}{func() []string {
									arr := *(*[]gopurs_runtime.Value)(__local_var_9_6.UnsafePtr)
									unboxed := make([]string, len(arr))
									for i, v := range arr {
										unboxed[i] = v.StrVal()
									}
									return unboxed
								}(), []string{"a", "b", "c"}}), gopurs_runtime.Value{})
							}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
									arr := []string{}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), __local_var_1_1), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4024920945("", struct {
										actual   gopurs_runtime.Value
										expected gopurs_runtime.Value
									}{gopurs_runtime.Apply(Call_Main_run(__local_var_1_1), Get_Foreign_Object_empty()), func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 gopurs_runtime.Value
											V2 bool
										}{gopurs_runtime.Value{}, Get_Foreign_Object_empty(), true}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}()}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_12_7 shape=App(Var) bindingType=Any
										__local_var_12_7 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
										_ = __local_var_12_7
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											__local_var_13_8 := gopurs_runtime.Apply(__local_var_12_7, gopurs_runtime.Value{})
											_ = __local_var_13_8
											return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2499841393("", struct {
												actual   []string
												expected []string
											}{func() []string {
												arr := *(*[]gopurs_runtime.Value)(__local_var_13_8.UnsafePtr)
												unboxed := make([]string, len(arr))
												for i, v := range arr {
													unboxed[i] = v.StrVal()
												}
												return unboxed
											}(), []string{}}), gopurs_runtime.Value{})
										}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												__local_var_14_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_partial(__local_var_1_1, int64(10))), gopurs_runtime.Value{})
												_ = __local_var_14_9
												// TAST (Let): __local_var_15_10 shape=App(Var) bindingType=Any
												__local_var_15_10 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
												_ = __local_var_15_10
												return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_16_11 := gopurs_runtime.Apply(__local_var_15_10, gopurs_runtime.Value{})
													_ = __local_var_16_11
													return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2499841393("", struct {
														actual   []string
														expected []string
													}{func() []string {
														arr := *(*[]gopurs_runtime.Value)(__local_var_16_11.UnsafePtr)
														unboxed := make([]string, len(arr))
														for i, v := range arr {
															unboxed[i] = v.StrVal()
														}
														return unboxed
													}(), []string{"capture"}}), gopurs_runtime.Value{})
												}), gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														// TAST (Let): __local_var_17_12 shape=App(Var) bindingType=Any
														__local_var_17_12 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_14_9)
														_ = __local_var_17_12
														__local_var_18_13 := gopurs_runtime.Apply(__local_var_17_12, gopurs_runtime.Value{})
														_ = __local_var_18_13
														return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4024920945("", struct {
															actual   gopurs_runtime.Value
															expected gopurs_runtime.Value
														}{gopurs_runtime.Apply(__local_var_18_13, gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST())), Get_Foreign_Object_ST_go__new(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str("a"), gopurs_runtime.Int(int64(3)))))), func() gopurs_runtime.Value {
															_v := struct {
																V0 gopurs_runtime.Value
																V1 gopurs_runtime.Value
																V2 bool
															}{gopurs_runtime.Value{}, gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST())), Get_Foreign_Object_ST_go__new(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str("a"), gopurs_runtime.Int(int64(13))))), true}
															if _v.V2 {
																return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
															}
															return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
														}()}), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4024920945("", struct {
																actual   gopurs_runtime.Value
																expected gopurs_runtime.Value
															}{gopurs_runtime.Apply(__local_var_18_13, gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST())), Get_Foreign_Object_ST_go__new(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str("a"), gopurs_runtime.Int(int64(7)))))), func() gopurs_runtime.Value {
																_v := struct {
																	V0 gopurs_runtime.Value
																	V1 gopurs_runtime.Value
																	V2 bool
																}{gopurs_runtime.Value{}, gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST())), Get_Foreign_Object_ST_go__new(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str("a"), gopurs_runtime.Int(int64(17))))), true}
																if _v.V2 {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
																}
																return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
															}()}), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
																// TAST (Let): __local_var_21_14 shape=App(Var) bindingType=Any
																__local_var_21_14 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)
																_ = __local_var_21_14
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_22_15 := gopurs_runtime.Apply(__local_var_21_14, gopurs_runtime.Value{})
																	_ = __local_var_22_15
																	return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2499841393("", struct {
																		actual   []string
																		expected []string
																	}{func() []string {
																		arr := *(*[]gopurs_runtime.Value)(__local_var_22_15.UnsafePtr)
																		unboxed := make([]string, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v.StrVal()
																		}
																		return unboxed
																	}(), []string{"capture"}}), gopurs_runtime.Value{})
																}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___692899441("", struct {
																		actual   *Constructor_Data_Maybe_Just[int64]
																		expected *Constructor_Data_Maybe_Just[int64]
																	}{Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Foreign_Object__lookup(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}, Get_Data_Maybe_Just(), gopurs_runtime.Str("a"), input_2_2))), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
																		_v := struct {
																			V0 gopurs_runtime.Value
																			V1 bool
																		}{gopurs_runtime.Int(int64(3)), true}
																		if _v.V1 {
																			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
																		}
																		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
																	}()))}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
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

func Call_Main_track(calls_0_loop gopurs_runtime.Value, key_1_loop string, value_2_loop int64) int64 {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var key_1 string = key_1_loop
	_ = key_1
	var value_2 int64 = value_2_loop
	_ = value_2
	return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			arr := func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), seen_3, func() gopurs_runtime.Value {
					arr := []string{key_1}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr))).UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	}), calls_0), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(value_2)
		})
	}))).IntVal
}

func Call_Main_run(calls_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	traverseObjectEither_arg_0 := gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Main_3467869399_1812164904(Rebox_Main_1812164904_3467869399(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Object_traversableWithIndexObject()))))}
	_ = traverseObjectEither_arg_0
	_ = traverseObjectEither_arg_0
	traverseObjectEither_arg_1 := gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))}
	_ = traverseObjectEither_arg_1
	_ = traverseObjectEither_arg_1
	traverseObjectEither_arg_4 := gopurs_runtime.Func2(func(key_1 gopurs_runtime.Value, value_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): observed_3_2 shape=App(Var) bindingType=Int
		observed_3_2 := Call_Main_track(calls_0, key_1.StrVal(), value_2.IntVal)
		_ = observed_3_2
		var __t3 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}
		{
			if (observed_3_2) < (int64(0)) {
				__t3 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{gopurs_runtime.Str(key_1.StrVal()), gopurs_runtime.Value{}, false}
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Int((observed_3_2) * (int64(2))), true}
		}
	end_branch_3:
		return func() gopurs_runtime.Value {
			_v := __t3
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	})
	_ = traverseObjectEither_arg_4
	_ = traverseObjectEither_arg_4
	return gopurs_runtime.Func(func(traverseObjectEither_input_5 gopurs_runtime.Value) gopurs_runtime.Value {
		traverseObjectEither_values_5 := gopurs_runtime.UnboxObject(traverseObjectEither_input_5)
		_ = traverseObjectEither_values_5
		traverseObjectEither_keys_5 := make([]string, 0, len(traverseObjectEither_values_5))
		_ = traverseObjectEither_keys_5
		for traverseObjectEither_key_5 := range traverseObjectEither_values_5 {
			traverseObjectEither_keys_5 = append(traverseObjectEither_keys_5, traverseObjectEither_key_5)
		}
		sort.Strings(traverseObjectEither_keys_5)
		traverseObjectEither_output_5 := make(map[string]any, len(traverseObjectEither_values_5))
		_ = traverseObjectEither_output_5
		traverseObjectEither_firstError_5 := gopurs_runtime.Value{}
		_ = traverseObjectEither_firstError_5
		traverseObjectEither_failed_5 := false
		_ = traverseObjectEither_failed_5
		for _, traverseObjectEither_key_5 := range traverseObjectEither_keys_5 {
			traverseObjectEither_result_5 := gopurs_runtime.Apply2(traverseObjectEither_arg_4, gopurs_runtime.Str(traverseObjectEither_key_5), gopurs_runtime.Box((traverseObjectEither_values_5)[traverseObjectEither_key_5]))
			_ = traverseObjectEither_result_5
			if traverseObjectEither_result_5.Type == 9 && traverseObjectEither_result_5.IntVal == 2465973597 {
				traverseObjectEither_output_5[traverseObjectEither_key_5] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseObjectEither_result_5.UnsafePtr).V0
			} else {
				if traverseObjectEither_result_5.Type == 9 && traverseObjectEither_result_5.IntVal == 3711209382 {
					if !(traverseObjectEither_failed_5) {
						traverseObjectEither_firstError_5 = traverseObjectEither_result_5
						traverseObjectEither_failed_5 = true
					} else {

					}
				} else {
					panic("Failed pattern match")
				}
			}
		}
		if traverseObjectEither_failed_5 {
			return traverseObjectEither_firstError_5
		} else {

		}
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Any(traverseObjectEither_output_5), true}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	})
}

func Call_Main_partial(calls_0_loop gopurs_runtime.Value, offset_1_loop int64) gopurs_runtime.Value {
	var calls_0 gopurs_runtime.Value = calls_0_loop
	_ = calls_0
	var offset_1 int64 = offset_1_loop
	_ = offset_1
	// TAST (Let): captured_2_0 shape=App(Var) bindingType=Int
	captured_2_0 := Call_Main_track(calls_0, "capture", offset_1)
	_ = captured_2_0
	traverseObjectEither_arg_1 := gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Main_3467869399_1812164904(Rebox_Main_1812164904_3467869399(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Foreign_Object_traversableWithIndexObject()))))}
	_ = traverseObjectEither_arg_1
	_ = traverseObjectEither_arg_1
	traverseObjectEither_arg_2 := gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))}
	_ = traverseObjectEither_arg_2
	_ = traverseObjectEither_arg_2
	traverseObjectEither_arg_3 := gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, value_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Int((value_4.IntVal) + (captured_2_0)), true}
			if _v.V2 {
				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
		}()
	})
	_ = traverseObjectEither_arg_3
	_ = traverseObjectEither_arg_3
	return gopurs_runtime.Func(func(traverseObjectEither_input_4 gopurs_runtime.Value) gopurs_runtime.Value {
		traverseObjectEither_values_4 := gopurs_runtime.UnboxObject(traverseObjectEither_input_4)
		_ = traverseObjectEither_values_4
		traverseObjectEither_keys_4 := make([]string, 0, len(traverseObjectEither_values_4))
		_ = traverseObjectEither_keys_4
		for traverseObjectEither_key_4 := range traverseObjectEither_values_4 {
			traverseObjectEither_keys_4 = append(traverseObjectEither_keys_4, traverseObjectEither_key_4)
		}
		sort.Strings(traverseObjectEither_keys_4)
		traverseObjectEither_output_4 := make(map[string]any, len(traverseObjectEither_values_4))
		_ = traverseObjectEither_output_4
		traverseObjectEither_firstError_4 := gopurs_runtime.Value{}
		_ = traverseObjectEither_firstError_4
		traverseObjectEither_failed_4 := false
		_ = traverseObjectEither_failed_4
		for _, traverseObjectEither_key_4 := range traverseObjectEither_keys_4 {
			traverseObjectEither_result_4 := gopurs_runtime.Apply2(traverseObjectEither_arg_3, gopurs_runtime.Str(traverseObjectEither_key_4), gopurs_runtime.Box((traverseObjectEither_values_4)[traverseObjectEither_key_4]))
			_ = traverseObjectEither_result_4
			if traverseObjectEither_result_4.Type == 9 && traverseObjectEither_result_4.IntVal == 2465973597 {
				traverseObjectEither_output_4[traverseObjectEither_key_4] = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(traverseObjectEither_result_4.UnsafePtr).V0
			} else {
				if traverseObjectEither_result_4.Type == 9 && traverseObjectEither_result_4.IntVal == 3711209382 {
					if !(traverseObjectEither_failed_4) {
						traverseObjectEither_firstError_4 = traverseObjectEither_result_4
						traverseObjectEither_failed_4 = true
					} else {

					}
				} else {
					panic("Failed pattern match")
				}
			}
		}
		if traverseObjectEither_failed_4 {
			return traverseObjectEither_firstError_4
		} else {

		}
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, gopurs_runtime.Any(traverseObjectEither_output_4), true}
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

func Rebox_Main_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_131790935_3790796878(in *Constructor_Data_Eq_Eq[[]string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_138441832_1830029836(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, int64]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1.IntVal
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

func Rebox_Main_1386611502_1953100407(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2818770644(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]{}
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

func Rebox_Main_1812164904_3467869399(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	out.V3 = in.V3
	return out
}

func Rebox_Main_1830029836_138441832(in *Constructor_Data_Tuple_Tuple[string, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_1953100407_1386611502(in *Constructor_Data_Show_Show[[]string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2818770644_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
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

func Rebox_Main_3467869399_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[string, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2
	out.V3 = in.V3
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

func Rebox_Main_3790796878_131790935(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_3960201844(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[int64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[int64]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3960201844_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
