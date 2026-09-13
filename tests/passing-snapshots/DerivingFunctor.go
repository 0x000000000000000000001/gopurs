package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))
	})
	return cache_Main_show
}

var cache_Main_eqRec gopurs_runtime.Value
var once_Main_eqRec sync.Once

func Get_Main_eqRec() gopurs_runtime.Value {
	once_Main_eqRec.Do(func() {
		cache_Main_eqRec = gopurs_runtime.Apply(Get_Data_Eq_eqRec(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRec
}

var cache_Main_eqRowCons gopurs_runtime.Value
var once_Main_eqRowCons sync.Once

func Get_Main_eqRowCons() gopurs_runtime.Value {
	once_Main_eqRowCons.Do(func() {
		cache_Main_eqRowCons = gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRowCons
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main_eqArray
}

var cache_Main_eqTuple gopurs_runtime.Value
var once_Main_eqTuple sync.Once

func Get_Main_eqTuple() gopurs_runtime.Value {
	once_Main_eqTuple.Do(func() {
		cache_Main_eqTuple = gopurs_runtime.Apply(Get_Data_Tuple_eqTuple(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main_eqTuple
}

var cache_Main_eqArray1 gopurs_runtime.Value
var once_Main_eqArray1 sync.Once

func Get_Main_eqArray1() gopurs_runtime.Value {
	once_Main_eqArray1.Do(func() {
		cache_Main_eqArray1 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main_eqArray1
}

var cache_Main_eqArray2 gopurs_runtime.Value
var once_Main_eqArray2 sync.Once

func Get_Main_eqArray2() gopurs_runtime.Value {
	once_Main_eqArray2.Do(func() {
		cache_Main_eqArray2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray2
}

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_T
}

var cache_Main_M0 gopurs_runtime.Value
var once_Main_M0 sync.Once

func Get_Main_M0() gopurs_runtime.Value {
	once_Main_M0.Do(func() {
		cache_Main_M0 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(value1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()}))}
			})
		})
	})
	return cache_Main_M0
}

var cache_Main_M0__4162883579 gopurs_runtime.Value
var once_Main_M0__4162883579 sync.Once

func Get_Main_M0__4162883579() gopurs_runtime.Value {
	once_Main_M0__4162883579.Do(func() {
		cache_Main_M0__4162883579 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M0__4162883579(__eta_norm_1_0_box.IntVal, func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}())
		})
	})
	return cache_Main_M0__4162883579
}

var cache_Main_M0__203000413 gopurs_runtime.Value
var once_Main_M0__203000413 sync.Once

func Get_Main_M0__203000413() gopurs_runtime.Value {
	once_Main_M0__203000413.Do(func() {
		cache_Main_M0__203000413 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M0__203000413(__eta_norm_1_0_box.StrVal(), func() []string {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}())
		})
	})
	return cache_Main_M0__203000413
}

var cache_Main_M1 gopurs_runtime.Value
var once_Main_M1 sync.Once

func Get_Main_M1() gopurs_runtime.Value {
	once_Main_M1.Do(func() {
		cache_Main_M1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal}))}
		})
	})
	return cache_Main_M1
}

var cache_Main_M1__1238571811 gopurs_runtime.Value
var once_Main_M1__1238571811 sync.Once

func Get_Main_M1__1238571811() gopurs_runtime.Value {
	once_Main_M1__1238571811.Do(func() {
		cache_Main_M1__1238571811 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M1__1238571811(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_M1__1238571811
}

var cache_Main_M1__371339077 gopurs_runtime.Value
var once_Main_M1__371339077 sync.Once

func Get_Main_M1__371339077() gopurs_runtime.Value {
	once_Main_M1__371339077.Do(func() {
		cache_Main_M1__371339077 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M1__371339077(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_M1__371339077
}

var cache_Main_M2 gopurs_runtime.Value
var once_Main_M2 sync.Once

func Get_Main_M2() gopurs_runtime.Value {
	once_Main_M2.Do(func() {
		cache_Main_M2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_M2
}

var cache_Main_M2__400289010 gopurs_runtime.Value
var once_Main_M2__400289010 sync.Once

func Get_Main_M2__400289010() gopurs_runtime.Value {
	once_Main_M2__400289010.Do(func() {
		cache_Main_M2__400289010 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M2__400289010(__eta_norm_0_0_box)
		})
	})
	return cache_Main_M2__400289010
}

var cache_Main_M2__1079474578 gopurs_runtime.Value
var once_Main_M2__1079474578 sync.Once

func Get_Main_M2__1079474578() gopurs_runtime.Value {
	once_Main_M2__1079474578.Do(func() {
		cache_Main_M2__1079474578 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M2__1079474578(__eta_norm_0_0_box)
		})
	})
	return cache_Main_M2__1079474578
}

var cache_Main_M3 gopurs_runtime.Value
var once_Main_M3 sync.Once

func Get_Main_M3() gopurs_runtime.Value {
	once_Main_M3.Do(func() {
		cache_Main_M3 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_M3
}

var cache_Main_M4 gopurs_runtime.Value
var once_Main_M4 sync.Once

func Get_Main_M4() gopurs_runtime.Value {
	once_Main_M4.Do(func() {
		cache_Main_M4 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
				nested gopurs_runtime.Value
			} {
				orig := value0
				_ = orig
				clone := struct {
					nested gopurs_runtime.Value
				}{}
				clone.nested = gopurs_runtime.RecordGet(orig, "nested")
				return clone
			}()}))}
		})
	})
	return cache_Main_M4
}

var cache_Main_M5 gopurs_runtime.Value
var once_Main_M5 sync.Once

func Get_Main_M5() gopurs_runtime.Value {
	once_Main_M5.Do(func() {
		cache_Main_M5 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(value5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(value6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(value7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1, func() []int64 {
											arr := *(*[]gopurs_runtime.Value)(value2.UnsafePtr)
											unboxed := make([]int64, len(arr))
											for i, v := range arr {
												unboxed[i] = v.IntVal
											}
											return unboxed
										}(), func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(value3.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}(), value4, value5, value6, func() struct {
											nested gopurs_runtime.Value
										} {
											orig := value7
											_ = orig
											clone := struct {
												nested gopurs_runtime.Value
											}{}
											clone.nested = gopurs_runtime.RecordGet(orig, "nested")
											return clone
										}()}))}
									})
								})
							})
						})
					})
				})
			})
		})
	})
	return cache_Main_M5
}

var cache_Main_M6 gopurs_runtime.Value
var once_Main_M6 sync.Once

func Get_Main_M6() gopurs_runtime.Value {
	once_Main_M6.Do(func() {
		cache_Main_M6 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() [][][]gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(value0.UnsafePtr)
				unboxed := make([][][]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = func() [][]gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
						unboxed := make([][]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()
						}
						return unboxed
					}()
				}
				return unboxed
			}()}))}
		})
	})
	return cache_Main_M6
}

var cache_Main_M6__1664895013 gopurs_runtime.Value
var once_Main_M6__1664895013 sync.Once

func Get_Main_M6__1664895013() gopurs_runtime.Value {
	once_Main_M6__1664895013.Do(func() {
		cache_Main_M6__1664895013 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M6__1664895013(func() [][][]int64 {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr)
				unboxed := make([][][]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = func() [][]int64 {
						arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
						unboxed := make([][]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
						}
						return unboxed
					}()
				}
				return unboxed
			}())
		})
	})
	return cache_Main_M6__1664895013
}

var cache_Main_M6__2005871109 gopurs_runtime.Value
var once_Main_M6__2005871109 sync.Once

func Get_Main_M6__2005871109() gopurs_runtime.Value {
	once_Main_M6__2005871109.Do(func() {
		cache_Main_M6__2005871109 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M6__2005871109(func() [][][]string {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr)
				unboxed := make([][][]string, len(arr))
				for i, v := range arr {
					unboxed[i] = func() [][]string {
						arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
						unboxed := make([][]string, len(arr))
						for i, v := range arr {
							unboxed[i] = func() []string {
								arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
								unboxed := make([]string, len(arr))
								for i, v := range arr {
									unboxed[i] = v.StrVal()
								}
								return unboxed
							}()
						}
						return unboxed
					}()
				}
				return unboxed
			}())
		})
	})
	return cache_Main_M6__2005871109
}

var cache_Main_Fun3 gopurs_runtime.Value
var once_Main_Fun3 sync.Once

func Get_Main_Fun3() gopurs_runtime.Value {
	once_Main_Fun3.Do(func() {
		cache_Main_Fun3 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Fun3
}

var cache_Main_Fun2 gopurs_runtime.Value
var once_Main_Fun2 sync.Once

func Get_Main_Fun2() gopurs_runtime.Value {
	once_Main_Fun2.Do(func() {
		cache_Main_Fun2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Fun2
}

var cache_Main_Fun2__2445298257 gopurs_runtime.Value
var once_Main_Fun2__2445298257 sync.Once

func Get_Main_Fun2__2445298257() gopurs_runtime.Value {
	once_Main_Fun2__2445298257.Do(func() {
		cache_Main_Fun2__2445298257 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Fun2__2445298257(__eta_norm_0_0_box)
		})
	})
	return cache_Main_Fun2__2445298257
}

var cache_Main_Fun1 gopurs_runtime.Value
var once_Main_Fun1 sync.Once

func Get_Main_Fun1() gopurs_runtime.Value {
	once_Main_Fun1.Do(func() {
		cache_Main_Fun1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Fun1
}

var cache_Main_Fun1__2419099218 gopurs_runtime.Value
var once_Main_Fun1__2419099218 sync.Once

func Get_Main_Fun1__2419099218() gopurs_runtime.Value {
	once_Main_Fun1__2419099218.Do(func() {
		cache_Main_Fun1__2419099218 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Fun1__2419099218(__eta_norm_0_0_box)
		})
	})
	return cache_Main_Fun1__2419099218
}

var cache_Main_functorFun3 gopurs_runtime.Value
var once_Main_functorFun3 sync.Once

func Get_Main_functorFun3() gopurs_runtime.Value {
	once_Main_functorFun3.Do(func() {
		cache_Main_functorFun3 = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_functorFun3(dictFunctor_0_box)
		})
	})
	return cache_Main_functorFun3
}

var cache_Main_functorFun2 gopurs_runtime.Value
var once_Main_functorFun2 sync.Once

func Get_Main_functorFun2() gopurs_runtime.Value {
	once_Main_functorFun2.Do(func() {
		cache_Main_functorFun2 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0))), m_1)
		})}))}
	})
	return cache_Main_functorFun2
}

var cache_Main_functorFun1 gopurs_runtime.Value
var once_Main_functorFun1 sync.Once

func Get_Main_functorFun1() gopurs_runtime.Value {
	once_Main_functorFun1.Do(func() {
		cache_Main_functorFun1 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), f_0), m_1)
		})}))}
	})
	return cache_Main_functorFun1
}

var cache_Main_recordValueR gopurs_runtime.Value
var once_Main_recordValueR sync.Once

func Get_Main_recordValueR() gopurs_runtime.Value {
	once_Main_recordValueR.Do(func() {
		cache_Main_recordValueR = func() gopurs_runtime.Value {
			orig := struct {
				a           string
				arrayIgnore []int64
				empty       struct {
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
			}{"71", []int64{int64(92), int64(93)}, struct {
			}{}, []int64{int64(94)}, []string{"73"}, int64(91), []*Constructor_Data_Tuple_Tuple[int64, []string]{Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int(int64(1)), func() gopurs_runtime.Value {
					arr := []string{"1"}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}())), Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int(int64(2)), func() gopurs_runtime.Value {
					arr := []string{"2"}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))}, []string{"72"}}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				orig := orig.empty
				_ = orig
				return gopurs_runtime.RecordDict0()
			}(), func() gopurs_runtime.Value {
				arr := orig.fIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.fa
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.recursiveA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}()
	})
	return cache_Main_recordValueR
}

var cache_Main_recordValueL gopurs_runtime.Value
var once_Main_recordValueL sync.Once

func Get_Main_recordValueL() gopurs_runtime.Value {
	once_Main_recordValueL.Do(func() {
		cache_Main_recordValueL = func() gopurs_runtime.Value {
			orig := struct {
				a           int64
				arrayIgnore []int64
				empty       struct {
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
			}{int64(71), []int64{int64(92), int64(93)}, struct {
			}{}, []int64{int64(94)}, []int64{int64(73)}, int64(91), []*Constructor_Data_Tuple_Tuple[int64, []int64]{Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int(int64(1)), func() gopurs_runtime.Value {
					arr := []int64{int64(1)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}())), Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int(int64(2)), func() gopurs_runtime.Value {
					arr := []int64{int64(2)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))}, []int64{int64(72)}}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				orig := orig.empty
				_ = orig
				return gopurs_runtime.RecordDict0()
			}(), func() gopurs_runtime.Value {
				arr := orig.fIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.fa
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.recursiveA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_690815662_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}()
	})
	return cache_Main_recordValueL
}

var cache_Main_m6R gopurs_runtime.Value
var once_Main_m6R sync.Once

func Get_Main_m6R() gopurs_runtime.Value {
	once_Main_m6R.Do(func() {
		cache_Main_m6R = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer(Rebox_Main_1554627816_4224211191((&Constructor_Main_M6[gopurs_runtime.Value, string]{1, func() [][][]gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := [][][]string{[][]string{[]string{"1", "2"}}}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								arr := v
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Str(v)
								}
								return gopurs_runtime.Array(boxed)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([][][]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
					}
					return unboxed
				}()
			}
			return unboxed
		}()})))}
	})
	return cache_Main_m6R
}

var cache_Main_m6L gopurs_runtime.Value
var once_Main_m6L sync.Once

func Get_Main_m6L() gopurs_runtime.Value {
	once_Main_m6L.Do(func() {
		cache_Main_m6L = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer(Rebox_Main_2628211660_4224211191((&Constructor_Main_M6[gopurs_runtime.Value, int64]{1, func() [][][]gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := [][][]int64{[][]int64{[]int64{int64(1), int64(2)}}}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								arr := v
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([][][]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
					}
					return unboxed
				}()
			}
			return unboxed
		}()})))}
	})
	return cache_Main_m6L
}

var cache_Main_m5R gopurs_runtime.Value
var once_Main_m5R sync.Once

func Get_Main_m5R() gopurs_runtime.Value {
	once_Main_m5R.Do(func() {
		cache_Main_m5R = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_1302345387_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, string]{1, int64(0), "1", []int64{int64(2), int64(3)}, func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []string{"3", "4"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}(), func() gopurs_runtime.Value {
			arr := []string{"5", "6"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{int64(7), int64(8)}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				empty       struct {
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.empty = func() struct {
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {
					}{}

					return clone
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []string] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []string], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}()
				clone.zArrayA = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				orig := orig.empty
				_ = orig
				return gopurs_runtime.RecordDict0()
			}(), func() gopurs_runtime.Value {
				arr := orig.fIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.fa
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.recursiveA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}(), func() struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				} {
					orig := Get_Main_recordValueR()
					_ = orig
					clone := struct {
						a           string
						arrayIgnore []int64
						empty       struct {
						}
						fIgnore    []int64
						fa         []string
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
						zArrayA    []string
					}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
					clone.arrayIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.empty = func() struct {
					} {
						orig := gopurs_runtime.RecordGet(orig, "empty")
						_ = orig
						clone := struct {
						}{}

						return clone
					}()
					clone.fIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fa = func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()
					clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
					clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []string] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []string], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					clone.zArrayA = func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					orig := orig.empty
					_ = orig
					return gopurs_runtime.RecordDict0()
				}(), func() gopurs_runtime.Value {
					arr := orig.fIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.fa
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
					arr := orig.recursiveA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.zArrayA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return clone
		}()})))}
	})
	return cache_Main_m5R
}

var cache_Main_m5L gopurs_runtime.Value
var once_Main_m5L sync.Once

func Get_Main_m5L() gopurs_runtime.Value {
	once_Main_m5L.Do(func() {
		cache_Main_m5L = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_2850066319_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, int64]{1, int64(0), int64(1), []int64{int64(2), int64(3)}, func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []int64{int64(3), int64(4)}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}(), func() gopurs_runtime.Value {
			arr := []int64{int64(5), int64(6)}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{int64(7), int64(8)}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			orig := func() struct {
				a           int64
				arrayIgnore []int64
				empty       struct {
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
			} {
				orig := Get_Main_recordValueL()
				_ = orig
				clone := struct {
					a           int64
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.empty = func() struct {
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {
					}{}

					return clone
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []int64] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}()
				clone.zArrayA = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				orig := orig.empty
				_ = orig
				return gopurs_runtime.RecordDict0()
			}(), func() gopurs_runtime.Value {
				arr := orig.fIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.fa
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.recursiveA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_690815662_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}(), func() struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           int64
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
				} {
					orig := Get_Main_recordValueL()
					_ = orig
					clone := struct {
						a           int64
						arrayIgnore []int64
						empty       struct {
						}
						fIgnore    []int64
						fa         []int64
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
						zArrayA    []int64
					}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
					clone.arrayIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.empty = func() struct {
					} {
						orig := gopurs_runtime.RecordGet(orig, "empty")
						_ = orig
						clone := struct {
						}{}

						return clone
					}()
					clone.fIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fa = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
					clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []int64] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []int64], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					clone.zArrayA = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					orig := orig.empty
					_ = orig
					return gopurs_runtime.RecordDict0()
				}(), func() gopurs_runtime.Value {
					arr := orig.fIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.fa
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
					arr := orig.recursiveA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_690815662_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.zArrayA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return clone
		}()})))}
	})
	return cache_Main_m5L
}

var cache_Main_m4R gopurs_runtime.Value
var once_Main_m4R sync.Once

func Get_Main_m4R() gopurs_runtime.Value {
	once_Main_m4R.Do(func() {
		cache_Main_m4R = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				} {
					orig := Get_Main_recordValueR()
					_ = orig
					clone := struct {
						a           string
						arrayIgnore []int64
						empty       struct {
						}
						fIgnore    []int64
						fa         []string
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
						zArrayA    []string
					}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
					clone.arrayIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.empty = func() struct {
					} {
						orig := gopurs_runtime.RecordGet(orig, "empty")
						_ = orig
						clone := struct {
						}{}

						return clone
					}()
					clone.fIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fa = func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()
					clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
					clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []string] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []string], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					clone.zArrayA = func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					orig := orig.empty
					_ = orig
					return gopurs_runtime.RecordDict0()
				}(), func() gopurs_runtime.Value {
					arr := orig.fIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.fa
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
					arr := orig.recursiveA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.zArrayA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return clone
		}()})))}
	})
	return cache_Main_m4R
}

var cache_Main_m4L gopurs_runtime.Value
var once_Main_m4L sync.Once

func Get_Main_m4L() gopurs_runtime.Value {
	once_Main_m4L.Do(func() {
		cache_Main_m4L = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_2124045134_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, int64]{1, func() struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           int64
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
				} {
					orig := Get_Main_recordValueL()
					_ = orig
					clone := struct {
						a           int64
						arrayIgnore []int64
						empty       struct {
						}
						fIgnore    []int64
						fa         []int64
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
						zArrayA    []int64
					}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
					clone.arrayIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.empty = func() struct {
					} {
						orig := gopurs_runtime.RecordGet(orig, "empty")
						_ = orig
						clone := struct {
						}{}

						return clone
					}()
					clone.fIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fa = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
					clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []int64] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []int64], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					clone.zArrayA = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					orig := orig.empty
					_ = orig
					return gopurs_runtime.RecordDict0()
				}(), func() gopurs_runtime.Value {
					arr := orig.fIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.fa
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
					arr := orig.recursiveA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_690815662_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.zArrayA
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return clone
		}()})))}
	})
	return cache_Main_m4L
}

var cache_Main_m3R gopurs_runtime.Value
var once_Main_m3R sync.Once

func Get_Main_m3R() gopurs_runtime.Value {
	once_Main_m3R.Do(func() {
		cache_Main_m3R = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_2785108781_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				empty       struct {
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.empty = func() struct {
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {
					}{}

					return clone
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []string] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []string], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}()
				clone.zArrayA = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				orig := orig.empty
				_ = orig
				return gopurs_runtime.RecordDict0()
			}(), func() gopurs_runtime.Value {
				arr := orig.fIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.fa
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.recursiveA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}()})))}
	})
	return cache_Main_m3R
}

var cache_Main_m3L gopurs_runtime.Value
var once_Main_m3L sync.Once

func Get_Main_m3L() gopurs_runtime.Value {
	once_Main_m3L.Do(func() {
		cache_Main_m3L = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_837768713_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, int64]{1, func() gopurs_runtime.Value {
			orig := func() struct {
				a           int64
				arrayIgnore []int64
				empty       struct {
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
			} {
				orig := Get_Main_recordValueL()
				_ = orig
				clone := struct {
					a           int64
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.empty = func() struct {
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {
					}{}

					return clone
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []int64] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}()
				clone.zArrayA = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				orig := orig.empty
				_ = orig
				return gopurs_runtime.RecordDict0()
			}(), func() gopurs_runtime.Value {
				arr := orig.fIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.fa
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.recursiveA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_690815662_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}()})))}
	})
	return cache_Main_m3L
}

var cache_Main_m2R gopurs_runtime.Value
var once_Main_m2R sync.Once

func Get_Main_m2R() gopurs_runtime.Value {
	once_Main_m2R.Do(func() {
		cache_Main_m2R = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			arr := []string{"0", "1"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()})))}
	})
	return cache_Main_m2R
}

var cache_Main_m2L gopurs_runtime.Value
var once_Main_m2L sync.Once

func Get_Main_m2L() gopurs_runtime.Value {
	once_Main_m2L.Do(func() {
		cache_Main_m2L = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4220871112_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, int64]{1, func() gopurs_runtime.Value {
			arr := []int64{int64(0), int64(1)}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()})))}
	})
	return cache_Main_m2L
}

var cache_Main_m1R gopurs_runtime.Value
var once_Main_m1R sync.Once

func Get_Main_m1R() gopurs_runtime.Value {
	once_Main_m1R.Do(func() {
		cache_Main_m1R = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, int64(0)})))}
	})
	return cache_Main_m1R
}

var cache_Main_m1L gopurs_runtime.Value
var once_Main_m1L sync.Once

func Get_Main_m1L() gopurs_runtime.Value {
	once_Main_m1L.Do(func() {
		cache_Main_m1L = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_147758475_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, int64]{1, int64(0)})))}
	})
	return cache_Main_m1L
}

var cache_Main_m0R gopurs_runtime.Value
var once_Main_m0R sync.Once

func Get_Main_m0R() gopurs_runtime.Value {
	once_Main_m0R.Do(func() {
		cache_Main_m0R = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977((&Constructor_Main_M0[gopurs_runtime.Value, string]{1, "0", func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []string{"1", "2"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}()})))}
	})
	return cache_Main_m0R
}

var cache_Main_m0L gopurs_runtime.Value
var once_Main_m0L sync.Once

func Get_Main_m0L() gopurs_runtime.Value {
	once_Main_m0L.Do(func() {
		cache_Main_m0L = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3716704586_67812977((&Constructor_Main_M0[gopurs_runtime.Value, int64]{1, int64(0), func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []int64{int64(1), int64(2)}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}()})))}
	})
	return cache_Main_m0L
}

var cache_Main_functorT gopurs_runtime.Value
var once_Main_functorT sync.Once

func Get_Main_functorT() gopurs_runtime.Value {
	once_Main_functorT.Do(func() {
		cache_Main_functorT = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(dictShow_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), f_0, gopurs_runtime.Apply(m_1, dictShow_2))
			})
		})}))}
	})
	return cache_Main_functorT
}

var cache_Main_taTests gopurs_runtime.Value
var once_Main_taTests sync.Once

func Get_Main_taTests() gopurs_runtime.Value {
	once_Main_taTests.Do(func() {
		cache_Main_taTests = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(Func [Int] String)
			__local_var_0_0 := Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))
			_ = __local_var_0_0
			return gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map show T"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), __local_var_0_0, gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(int64(42))
				}))
			}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}, gopurs_runtime.Str("hello")).StrVal()) == ("42")))
		}()
	})
	return cache_Main_taTests
}

var cache_Main_functorM gopurs_runtime.Value
var once_Main_functorM sync.Once

func Get_Main_functorM() gopurs_runtime.Value {
	once_Main_functorM.Do(func() {
		cache_Main_functorM = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_functorM(dictFunctor_0_box)
		})
	})
	return cache_Main_functorM
}

var cache_Main_functorM1 gopurs_runtime.Value
var once_Main_functorM1 sync.Once

func Get_Main_functorM1() gopurs_runtime.Value {
	once_Main_functorM1.Do(func() {
		cache_Main_functorM1 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main_functorM1
}

var cache_Main_f3Test gopurs_runtime.Value
var once_Main_f3Test sync.Once

func Get_Main_f3Test() gopurs_runtime.Value {
	once_Main_f3Test.Do(func() {
		cache_Main_f3Test = gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - Fun3"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("nested")
		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("zArrayA")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("recursiveA")
		})), Call_Data_Eq_eqArray(Call_Data_Tuple_eqTuple(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("ignore")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fa")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("empty")
		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Get_Data_Eq_eqRowNil())), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("arrayIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))))), "eq"), func() gopurs_runtime.Value {
			arr := func() [][][]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}
			} {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorFun3(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr := [][][]struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
							}
						}{[][]struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
							}
						}{[]struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
							}
						}{struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
							}
						}{func() struct {
							a           int64
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    []int64
							fa         []int64
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
							zArrayA    []int64
						} {
							orig := Get_Main_recordValueL()
							_ = orig
							clone := struct {
								a           int64
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
							}{}
							clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
							clone.arrayIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.empty = func() struct {
							} {
								orig := gopurs_runtime.RecordGet(orig, "empty")
								_ = orig
								clone := struct {
								}{}

								return clone
							}()
							clone.fIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.fa = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
							clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []int64] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []int64], len(arr))
								for i, v := range arr {
									unboxed[i] = Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
								}
								return unboxed
							}()
							clone.zArrayA = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							return clone
						}()}}}}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								arr := v
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = func() gopurs_runtime.Value {
										arr := v
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = func() gopurs_runtime.Value {
												orig := v
												_ = orig
												return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
													orig := orig.nested
													_ = orig
													return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
														arr := orig.arrayIgnore
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = gopurs_runtime.Int(v)
														}
														return gopurs_runtime.Array(boxed)
													}(), func() gopurs_runtime.Value {
														orig := orig.empty
														_ = orig
														return gopurs_runtime.RecordDict0()
													}(), func() gopurs_runtime.Value {
														arr := orig.fIgnore
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = gopurs_runtime.Int(v)
														}
														return gopurs_runtime.Array(boxed)
													}(), func() gopurs_runtime.Value {
														arr := orig.fa
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = gopurs_runtime.Int(v)
														}
														return gopurs_runtime.Array(boxed)
													}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
														arr := orig.recursiveA
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_690815662_138441832(v))}
														}
														return gopurs_runtime.Array(boxed)
													}(), func() gopurs_runtime.Value {
														arr := orig.zArrayA
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = gopurs_runtime.Int(v)
														}
														return gopurs_runtime.Array(boxed)
													}()})
												}())
											}()
										}
										return gopurs_runtime.Array(boxed)
									}()
								}
								return gopurs_runtime.Array(boxed)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()
				})), Get_Data_Unit_unit()).UnsafePtr)
				unboxed := make([][][]struct {
					nested struct {
						a           string
						arrayIgnore []int64
						empty       struct {
						}
						fIgnore    []int64
						fa         []string
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
						zArrayA    []string
					}
				}, len(arr))
				for i, v := range arr {
					unboxed[i] = func() [][]struct {
						nested struct {
							a           string
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    []int64
							fa         []string
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
							zArrayA    []string
						}
					} {
						arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
						unboxed := make([][]struct {
							nested struct {
								a           string
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []string
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
								zArrayA    []string
							}
						}, len(arr))
						for i, v := range arr {
							unboxed[i] = func() []struct {
								nested struct {
									a           string
									arrayIgnore []int64
									empty       struct {
									}
									fIgnore    []int64
									fa         []string
									ignore     int64
									recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
									zArrayA    []string
								}
							} {
								arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
								unboxed := make([]struct {
									nested struct {
										a           string
										arrayIgnore []int64
										empty       struct {
										}
										fIgnore    []int64
										fa         []string
										ignore     int64
										recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
										zArrayA    []string
									}
								}, len(arr))
								for i, v := range arr {
									unboxed[i] = func() struct {
										nested struct {
											a           string
											arrayIgnore []int64
											empty       struct {
											}
											fIgnore    []int64
											fa         []string
											ignore     int64
											recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
											zArrayA    []string
										}
									} {
										orig := v
										_ = orig
										clone := struct {
											nested struct {
												a           string
												arrayIgnore []int64
												empty       struct {
												}
												fIgnore    []int64
												fa         []string
												ignore     int64
												recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
												zArrayA    []string
											}
										}{}
										clone.nested = func() struct {
											a           string
											arrayIgnore []int64
											empty       struct {
											}
											fIgnore    []int64
											fa         []string
											ignore     int64
											recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
											zArrayA    []string
										} {
											orig := gopurs_runtime.RecordGet(orig, "nested")
											_ = orig
											clone := struct {
												a           string
												arrayIgnore []int64
												empty       struct {
												}
												fIgnore    []int64
												fa         []string
												ignore     int64
												recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
												zArrayA    []string
											}{}
											clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
											clone.arrayIgnore = func() []int64 {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
												unboxed := make([]int64, len(arr))
												for i, v := range arr {
													unboxed[i] = v.IntVal
												}
												return unboxed
											}()
											clone.empty = func() struct {
											} {
												orig := gopurs_runtime.RecordGet(orig, "empty")
												_ = orig
												clone := struct {
												}{}

												return clone
											}()
											clone.fIgnore = func() []int64 {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
												unboxed := make([]int64, len(arr))
												for i, v := range arr {
													unboxed[i] = v.IntVal
												}
												return unboxed
											}()
											clone.fa = func() []string {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
												unboxed := make([]string, len(arr))
												for i, v := range arr {
													unboxed[i] = v.StrVal()
												}
												return unboxed
											}()
											clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
											clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []string] {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
												unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []string], len(arr))
												for i, v := range arr {
													unboxed[i] = Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
												}
												return unboxed
											}()
											clone.zArrayA = func() []string {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
												unboxed := make([]string, len(arr))
												for i, v := range arr {
													unboxed[i] = v.StrVal()
												}
												return unboxed
											}()
											return clone
										}()
										return clone
									}()
								}
								return unboxed
							}()
						}
						return unboxed
					}()
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							arr := v
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
										orig := orig.nested
										_ = orig
										return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
											arr := orig.arrayIgnore
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), func() gopurs_runtime.Value {
											orig := orig.empty
											_ = orig
											return gopurs_runtime.RecordDict0()
										}(), func() gopurs_runtime.Value {
											arr := orig.fIgnore
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), func() gopurs_runtime.Value {
											arr := orig.fa
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Str(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
											arr := orig.recursiveA
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
											}
											return gopurs_runtime.Array(boxed)
										}(), func() gopurs_runtime.Value {
											arr := orig.zArrayA
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Str(v)
											}
											return gopurs_runtime.Array(boxed)
										}()})
									}())
								}()
							}
							return gopurs_runtime.Array(boxed)
						}()
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := [][][]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}
			}{[][]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}
			}{[]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}
			}{struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}
			}{func() struct {
				a           string
				arrayIgnore []int64
				empty       struct {
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.empty = func() struct {
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {
					}{}

					return clone
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []string] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []string], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
					}
					return unboxed
				}()
				clone.zArrayA = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				return clone
			}()}}}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							arr := v
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
										orig := orig.nested
										_ = orig
										return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
											arr := orig.arrayIgnore
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), func() gopurs_runtime.Value {
											orig := orig.empty
											_ = orig
											return gopurs_runtime.RecordDict0()
										}(), func() gopurs_runtime.Value {
											arr := orig.fIgnore
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), func() gopurs_runtime.Value {
											arr := orig.fa
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Str(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
											arr := orig.recursiveA
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
											}
											return gopurs_runtime.Array(boxed)
										}(), func() gopurs_runtime.Value {
											arr := orig.zArrayA
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Str(v)
											}
											return gopurs_runtime.Array(boxed)
										}()})
									}())
								}()
							}
							return gopurs_runtime.Array(boxed)
						}()
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal) != (0)))
	})
	return cache_Main_f3Test
}

var cache_Main_f2Test gopurs_runtime.Value
var once_Main_f2Test sync.Once

func Get_Main_f2Test() gopurs_runtime.Value {
	once_Main_f2Test.Do(func() {
		cache_Main_f2Test = func() gopurs_runtime.Value {
			// TAST (Let): fn1_0_0 shape=App(Var) bindingType=(Func [Int] String)
			fn1_0_0 := Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))
			_ = fn1_0_0
			return gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - Fun2"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), fn1_0_0))), gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := [][]int64{[]int64{(a_1.IntVal) + (b_2.IntVal)}}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							arr := v
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()
					}
					return gopurs_runtime.Array(boxed)
				}()
			}), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
				arr := func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap3 := func() gopurs_runtime.Value {
								arr := [][]int64{[]int64{int64(3)}}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = func() gopurs_runtime.Value {
										arr := v
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Int(v)
										}
										return gopurs_runtime.Array(boxed)
									}()
								}
								return gopurs_runtime.Array(boxed)
							}()
							_ = arr_val_arrayMap3
							arr_go_arrayMap3 := (*[]gopurs_runtime.Value)(arr_val_arrayMap3.UnsafePtr)
							_ = arr_go_arrayMap3
							res_go_arrayMap3 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap3))
							_ = res_go_arrayMap3
							for i_arrayMap3, v_arrayMap3 := range *arr_go_arrayMap3 {
								res_go_arrayMap3[i_arrayMap3] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), fn1_0_0), v_arrayMap3)
							}
							return gopurs_runtime.Array(res_go_arrayMap3)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()).UnsafePtr)
					unboxed := make([][]string, len(arr))
					for i, v := range arr {
						unboxed[i] = func() []string {
							arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
							unboxed := make([]string, len(arr))
							for i, v := range arr {
								unboxed[i] = v.StrVal()
							}
							return unboxed
						}()
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal) != (0)))
		}()
	})
	return cache_Main_f2Test
}

var cache_Main_f1Test gopurs_runtime.Value
var once_Main_f1Test sync.Once

func Get_Main_f1Test() gopurs_runtime.Value {
	once_Main_f1Test.Do(func() {
		cache_Main_f1Test = func() gopurs_runtime.Value {
			// TAST (Let): fn1_0_0 shape=App(Var) bindingType=(Func [Int] String)
			fn1_0_0 := Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))
			_ = fn1_0_0
			return gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - Fun1"), gopurs_runtime.Bool((gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), fn1_0_0), gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((a_1.IntVal) + (b_2.IntVal))
			}), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))).StrVal()) == (gopurs_runtime.Apply(fn1_0_0, gopurs_runtime.Int(int64(3))).StrVal())))
		}()
	})
	return cache_Main_f1Test
}

var cache_Main_funTests gopurs_runtime.Value
var once_Main_funTests sync.Once

func Get_Main_funTests() gopurs_runtime.Value {
	once_Main_funTests.Do(func() {
		cache_Main_funTests = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_f1Test(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_f2Test(), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_f3Test(), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Main_taTests()
				}))
			}))
		}))
	})
	return cache_Main_funTests
}

var cache_Main_eqM gopurs_runtime.Value
var once_Main_eqM sync.Once

func Get_Main_eqM() gopurs_runtime.Value {
	once_Main_eqM.Do(func() {
		cache_Main_eqM = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqM(dictEq1_0_box, dictEq_1_box)
		})
	})
	return cache_Main_eqM
}

var cache_Main_eqM1 gopurs_runtime.Value
var once_Main_eqM1 sync.Once

func Get_Main_eqM1() gopurs_runtime.Value {
	once_Main_eqM1.Do(func() {
		cache_Main_eqM1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))}
	})
	return cache_Main_eqM1
}

var cache_Main_maTests gopurs_runtime.Value
var once_Main_maTests sync.Once

func Get_Main_maTests() gopurs_runtime.Value {
	once_Main_maTests.Do(func() {
		cache_Main_maTests = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M0"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), Get_Main_m0L()), Get_Main_m0R()).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M1"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_147758475_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, int64]{1, int64(0)})))}), gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, int64(0)})))}).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M2"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), Get_Main_m2L()), Get_Main_m2R()).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M3"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_837768713_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, int64]{1, func() gopurs_runtime.Value {
						orig := func() struct {
							a           int64
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    []int64
							fa         []int64
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
							zArrayA    []int64
						} {
							orig := Get_Main_recordValueL()
							_ = orig
							clone := struct {
								a           int64
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
							}{}
							clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
							clone.arrayIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.empty = func() struct {
							} {
								orig := gopurs_runtime.RecordGet(orig, "empty")
								_ = orig
								clone := struct {
								}{}

								return clone
							}()
							clone.fIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.fa = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
							clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []int64] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []int64], len(arr))
								for i, v := range arr {
									unboxed[i] = Rebox_Main_138441832_690815662(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
								}
								return unboxed
							}()
							clone.zArrayA = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
							arr := orig.arrayIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							orig := orig.empty
							_ = orig
							return gopurs_runtime.RecordDict0()
						}(), func() gopurs_runtime.Value {
							arr := orig.fIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							arr := orig.fa
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
							arr := orig.recursiveA
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_690815662_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							arr := orig.zArrayA
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()})
					}()})))}), gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_2785108781_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
						orig := func() struct {
							a           string
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    []int64
							fa         []string
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
							zArrayA    []string
						} {
							orig := Get_Main_recordValueR()
							_ = orig
							clone := struct {
								a           string
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    []int64
								fa         []string
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
								zArrayA    []string
							}{}
							clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
							clone.arrayIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.empty = func() struct {
							} {
								orig := gopurs_runtime.RecordGet(orig, "empty")
								_ = orig
								clone := struct {
								}{}

								return clone
							}()
							clone.fIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.fa = func() []string {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
								unboxed := make([]string, len(arr))
								for i, v := range arr {
									unboxed[i] = v.StrVal()
								}
								return unboxed
							}()
							clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
							clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []string] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []string], len(arr))
								for i, v := range arr {
									unboxed[i] = Rebox_Main_138441832_1200068938(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
								}
								return unboxed
							}()
							clone.zArrayA = func() []string {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
								unboxed := make([]string, len(arr))
								for i, v := range arr {
									unboxed[i] = v.StrVal()
								}
								return unboxed
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
							arr := orig.arrayIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							orig := orig.empty
							_ = orig
							return gopurs_runtime.RecordDict0()
						}(), func() gopurs_runtime.Value {
							arr := orig.fIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							arr := orig.fa
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Str(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
							arr := orig.recursiveA
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_1200068938_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							arr := orig.zArrayA
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Str(v)
							}
							return gopurs_runtime.Array(boxed)
						}()})
					}()})))}).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M4"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), Get_Main_m4L()), Get_Main_m4R()).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M5"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), Get_Main_m5L()), Get_Main_m5R()).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M6"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}), "map"), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), Get_Main_m6L()), Get_Main_m6R()).IntVal) != (0)))
							}))
						}))
					}))
				}))
			}))
		}))
	})
	return cache_Main_maTests
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_maTests(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_funTests(), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_T[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M0[T_f any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 []gopurs_runtime.Value
}

type Constructor_Main_M1[T_f any, T_a any] struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_M2[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M3[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M4[T_f any, T_a any] struct {
	Rc uint32
	V0 struct {
		nested gopurs_runtime.Value
	}
}

type Constructor_Main_M5[T_f any, T_a any] struct {
	Rc uint32
	V0 int64
	V1 T_a
	V2 []int64
	V3 []gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
	V6 gopurs_runtime.Value
	V7 struct {
		nested gopurs_runtime.Value
	}
}

type Constructor_Main_M6[T_f any, T_a any] struct {
	Rc uint32
	V0 [][][]gopurs_runtime.Value
}

type Constructor_Main_Fun3[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Fun2[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Fun1[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_M0__4162883579(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop []int64) gopurs_runtime.Value {
M0__4162883579:
	for {
		if false {
			continue M0__4162883579
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 []int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3716704586_67812977((&Constructor_Main_M0[gopurs_runtime.Value, int64]{1, __eta_norm_1_0, func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := __eta_norm_0_1
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}()})))}
	}
}

func Call_Main_M0__203000413(__eta_norm_1_0_loop string, __eta_norm_0_1_loop []string) gopurs_runtime.Value {
M0__203000413:
	for {
		if false {
			continue M0__203000413
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 []string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977((&Constructor_Main_M0[gopurs_runtime.Value, string]{1, __eta_norm_1_0, func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := __eta_norm_0_1
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}()})))}
	}
}

func Call_Main_M1__1238571811(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
M1__1238571811:
	for {
		if false {
			continue M1__1238571811
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_147758475_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_M1__371339077(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
M1__371339077:
	for {
		if false {
			continue M1__371339077
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_M2__400289010(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
M2__400289010:
	for {
		if false {
			continue M2__400289010
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4220871112_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_M2__1079474578(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
M2__1079474578:
	for {
		if false {
			continue M2__1079474578
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_M6__1664895013(__eta_norm_0_0_loop [][][]int64) gopurs_runtime.Value {
M6__1664895013:
	for {
		if false {
			continue M6__1664895013
		}
		var __eta_norm_0_0 [][][]int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer(Rebox_Main_2628211660_4224211191((&Constructor_Main_M6[gopurs_runtime.Value, int64]{1, func() [][][]gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := __eta_norm_0_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								arr := v
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([][][]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
					}
					return unboxed
				}()
			}
			return unboxed
		}()})))}
	}
}

func Call_Main_M6__2005871109(__eta_norm_0_0_loop [][][]string) gopurs_runtime.Value {
M6__2005871109:
	for {
		if false {
			continue M6__2005871109
		}
		var __eta_norm_0_0 [][][]string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer(Rebox_Main_1554627816_4224211191((&Constructor_Main_M6[gopurs_runtime.Value, string]{1, func() [][][]gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := __eta_norm_0_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								arr := v
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Str(v)
								}
								return gopurs_runtime.Array(boxed)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([][][]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
					}
					return unboxed
				}()
			}
			return unboxed
		}()})))}
	}
}

func Call_Main_Fun2__2445298257(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Fun2__2445298257:
	for {
		if false {
			continue Fun2__2445298257
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Call_Main_Fun1__2419099218(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Fun1__2419099218:
	for {
		if false {
			continue Fun1__2419099218
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Call_Main_functorFun3(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=Any
			__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
			_ = __local_var_4_0
			return func() gopurs_runtime.Value {
				orig := func() struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
					}
				} {
					orig := gopurs_runtime.RecordUpdate1(v1_3, "nested", func() gopurs_runtime.Value {
						orig := func() struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
						} {
							orig := gopurs_runtime.RecordUpdateDict(gopurs_runtime.RecordGet(v1_3, "nested"), []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa")), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap10 := func() gopurs_runtime.Value {
												arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "recursiveA").UnsafePtr)
													unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
													for i, v := range arr {
														unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
													}
													return unboxed
												}()
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
												}
												return gopurs_runtime.Array(boxed)
											}()
											_ = arr_val_arrayMap10
											arr_go_arrayMap10 := (*[]gopurs_runtime.Value)(arr_val_arrayMap10.UnsafePtr)
											_ = arr_go_arrayMap10
											res_go_arrayMap10 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap10))
											_ = res_go_arrayMap10
											for i_arrayMap10, v_arrayMap10 := range *arr_go_arrayMap10 {
												res_go_arrayMap10[i_arrayMap10] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 gopurs_runtime.Value
														}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_4_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1)}
														return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
													}()))}
												}), v_arrayMap10)
											}
											return gopurs_runtime.Array(res_go_arrayMap10)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()).UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
								}
								return gopurs_runtime.Array(boxed)
							}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap10 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap10
									arr_go_arrayMap10 := (*[]gopurs_runtime.Value)(arr_val_arrayMap10.UnsafePtr)
									_ = arr_go_arrayMap10
									res_go_arrayMap10 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap10))
									_ = res_go_arrayMap10
									for i_arrayMap10, v_arrayMap10 := range *arr_go_arrayMap10 {
										res_go_arrayMap10[i_arrayMap10] = gopurs_runtime.Apply(f_1, v_arrayMap10)
									}
									return gopurs_runtime.Array(res_go_arrayMap10)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())})
							_ = orig
							clone := struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    gopurs_runtime.Value
								fa         gopurs_runtime.Value
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
								zArrayA    []gopurs_runtime.Value
							}{}
							clone.a = gopurs_runtime.RecordGet(orig, "a")
							clone.arrayIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.empty = func() struct {
							} {
								orig := gopurs_runtime.RecordGet(orig, "empty")
								_ = orig
								clone := struct {
								}{}

								return clone
							}()
							clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
							clone.fa = gopurs_runtime.RecordGet(orig, "fa")
							clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
							clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
								for i, v := range arr {
									unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
								}
								return unboxed
							}()
							clone.zArrayA = func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
							arr := orig.arrayIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							orig := orig.empty
							_ = orig
							return gopurs_runtime.RecordDict0()
						}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
							arr := orig.recursiveA
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Array(orig.zArrayA)})
					}())
					_ = orig
					clone := struct {
						nested struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
						}
					}{}
					clone.nested = func() struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
					} {
						orig := gopurs_runtime.RecordGet(orig, "nested")
						_ = orig
						clone := struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
						}{}
						clone.a = gopurs_runtime.RecordGet(orig, "a")
						clone.arrayIgnore = func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
							unboxed := make([]int64, len(arr))
							for i, v := range arr {
								unboxed[i] = v.IntVal
							}
							return unboxed
						}()
						clone.empty = func() struct {
						} {
							orig := gopurs_runtime.RecordGet(orig, "empty")
							_ = orig
							clone := struct {
							}{}

							return clone
						}()
						clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
						clone.fa = gopurs_runtime.RecordGet(orig, "fa")
						clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
						clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
							}
							return unboxed
						}()
						clone.zArrayA = func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()
						return clone
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
					orig := orig.nested
					_ = orig
					return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
						arr := orig.arrayIgnore
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						orig := orig.empty
						_ = orig
						return gopurs_runtime.RecordDict0()
					}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
						arr := orig.recursiveA
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
						}
						return gopurs_runtime.Array(boxed)
					}(), gopurs_runtime.Array(orig.zArrayA)})
				}())
			}()
		})))), m_2)
	})}))}
}

func Call_Main_functorM(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t4 gopurs_runtime.Value
		{
			if m_2.Type == 9 && m_2.IntVal == 3852365315 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap4 := gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)
						_ = arr_val_arrayMap4
						arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
						_ = arr_go_arrayMap4
						res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
						_ = res_go_arrayMap4
						for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
							res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(f_1, v_arrayMap4)
						}
						return gopurs_runtime.Array(res_go_arrayMap4)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 769986722 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2727978561 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 1830062304 {
				// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
				__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_0
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "fa")), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := func() gopurs_runtime.Value {
									arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "recursiveA").UnsafePtr)
										unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
										for i, v := range arr {
											unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
									}
									return gopurs_runtime.Array(boxed)
								}()
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
											_v := struct {
												V0 gopurs_runtime.Value
												V1 gopurs_runtime.Value
											}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
											return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
										}()))}
									}), v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						_ = arr_val_arrayMap5
						arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
						_ = arr_go_arrayMap5
						res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
						_ = res_go_arrayMap5
						for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
							res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_1, v_arrayMap5)
						}
						return gopurs_runtime.Array(res_go_arrayMap5)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())})}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 3190619783 {
				// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=Any
				__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_1
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
					nested gopurs_runtime.Value
				} {
					clone := (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0
					clone.nested = func() gopurs_runtime.Value {
						orig := func() struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
						} {
							orig := gopurs_runtime.RecordUpdateDict((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "fa")), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap6 := func() gopurs_runtime.Value {
												arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
													unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
													for i, v := range arr {
														unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
													}
													return unboxed
												}()
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
												}
												return gopurs_runtime.Array(boxed)
											}()
											_ = arr_val_arrayMap6
											arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
											_ = arr_go_arrayMap6
											res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
											_ = res_go_arrayMap6
											for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
												res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 gopurs_runtime.Value
														}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
														return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
													}()))}
												}), v_arrayMap6)
											}
											return gopurs_runtime.Array(res_go_arrayMap6)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()).UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
								}
								return gopurs_runtime.Array(boxed)
							}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_1, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())})
							_ = orig
							clone := struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    gopurs_runtime.Value
								fa         gopurs_runtime.Value
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
								zArrayA    []gopurs_runtime.Value
							}{}
							clone.a = gopurs_runtime.RecordGet(orig, "a")
							clone.arrayIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.empty = func() struct {
							} {
								orig := gopurs_runtime.RecordGet(orig, "empty")
								_ = orig
								clone := struct {
								}{}

								return clone
							}()
							clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
							clone.fa = gopurs_runtime.RecordGet(orig, "fa")
							clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
							clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
								for i, v := range arr {
									unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
								}
								return unboxed
							}()
							clone.zArrayA = func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
							arr := orig.arrayIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							orig := orig.empty
							_ = orig
							return gopurs_runtime.RecordDict0()
						}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
							arr := orig.recursiveA
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Array(orig.zArrayA)})
					}()
					return clone
				}()}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 108241190 {
				// TAST (Let): __local_var_3_2 shape=App(Var) bindingType=Any
				__local_var_3_2 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_2
				// TAST (Let): __local_var_3_3 shape=App(Var) bindingType=Any
				__local_var_3_3 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_3
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap4 := gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3)
						_ = arr_val_arrayMap4
						arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
						_ = arr_go_arrayMap4
						res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
						_ = res_go_arrayMap4
						for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
							res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(f_1, v_arrayMap4)
						}
						return gopurs_runtime.Array(res_go_arrayMap4)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "fa")), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := func() gopurs_runtime.Value {
									arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "recursiveA").UnsafePtr)
										unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
										for i, v := range arr {
											unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
									}
									return gopurs_runtime.Array(boxed)
								}()
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
											_v := struct {
												V0 gopurs_runtime.Value
												V1 gopurs_runtime.Value
											}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
											return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
										}()))}
									}), v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						_ = arr_val_arrayMap5
						arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
						_ = arr_go_arrayMap5
						res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
						_ = res_go_arrayMap5
						for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
							res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_1, v_arrayMap5)
						}
						return gopurs_runtime.Array(res_go_arrayMap5)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())}), func() struct {
					nested gopurs_runtime.Value
				} {
					clone := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7
					clone.nested = func() gopurs_runtime.Value {
						orig := func() struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
						} {
							orig := gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "fa")), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap6 := func() gopurs_runtime.Value {
												arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
													unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
													for i, v := range arr {
														unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
													}
													return unboxed
												}()
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
												}
												return gopurs_runtime.Array(boxed)
											}()
											_ = arr_val_arrayMap6
											arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
											_ = arr_go_arrayMap6
											res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
											_ = res_go_arrayMap6
											for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
												res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 gopurs_runtime.Value
														}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
														return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
													}()))}
												}), v_arrayMap6)
											}
											return gopurs_runtime.Array(res_go_arrayMap6)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()).UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
								}
								return gopurs_runtime.Array(boxed)
							}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_1, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())})
							_ = orig
							clone := struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								empty       struct {
								}
								fIgnore    gopurs_runtime.Value
								fa         gopurs_runtime.Value
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
								zArrayA    []gopurs_runtime.Value
							}{}
							clone.a = gopurs_runtime.RecordGet(orig, "a")
							clone.arrayIgnore = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.empty = func() struct {
							} {
								orig := gopurs_runtime.RecordGet(orig, "empty")
								_ = orig
								clone := struct {
								}{}

								return clone
							}()
							clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
							clone.fa = gopurs_runtime.RecordGet(orig, "fa")
							clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
							clone.recursiveA = func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "recursiveA").UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
								for i, v := range arr {
									unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
								}
								return unboxed
							}()
							clone.zArrayA = func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
							arr := orig.arrayIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), func() gopurs_runtime.Value {
							orig := orig.empty
							_ = orig
							return gopurs_runtime.RecordDict0()
						}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
							arr := orig.recursiveA
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Array(orig.zArrayA)})
					}()
					return clone
				}()}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2066233029 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() [][][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap4 := func() gopurs_runtime.Value {
							arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									arr := v
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Array(v)
									}
									return gopurs_runtime.Array(boxed)
								}()
							}
							return gopurs_runtime.Array(boxed)
						}()
						_ = arr_val_arrayMap4
						arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
						_ = arr_go_arrayMap4
						res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
						_ = res_go_arrayMap4
						for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
							res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)), v_arrayMap4)
						}
						return gopurs_runtime.Array(res_go_arrayMap4)
					}().UnsafePtr)
					unboxed := make([][][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = func() [][]gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
							unboxed := make([][]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()
							}
							return unboxed
						}()
					}
					return unboxed
				}()}))}
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_4:
		return __t4
	})}))}
}

func Call_Main_eqM(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	// TAST (Let): eqArray3_2_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a$scope98))])
	eqArray3_2_0 := Rebox_Main_3790796878_1939691112(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(dictEq_1)))
	_ = eqArray3_2_0
	// TAST (Let): eqArray4_3_1 shape=App(Var) bindingType=Any
	eqArray4_3_1 := Call_Data_Eq_eqArray(dictEq_1)
	_ = eqArray4_3_1
	// TAST (Let): eqArray5_4_2 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a$scope98))]))])
	eqArray5_4_2 := Rebox_Main_3790796878_1871481144(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(Call_Data_Tuple_eqTuple(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, eqArray4_3_1))))
	_ = eqArray5_4_2
	// TAST (Let): eqArray6_5_3 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a$scope98))))])
	eqArray6_5_3 := Rebox_Main_3790796878_1161550504(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(eqArray4_3_1))))
	_ = eqArray6_5_3
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t4 bool
		{
			if x_6.Type == 9 && x_6.IntVal == 3852365315 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 3852365315) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V1)).IntVal) != (0)))
				goto end_branch_4
			} else {

			}
		}
		{
			if x_6.Type == 9 && x_6.IntVal == 769986722 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0) == ((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0))
				goto end_branch_4
			} else {

			}
		}
		{
			if x_6.Type == 9 && x_6.IntVal == 2727978561 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 2727978561) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0).IntVal) != (0))
				goto end_branch_4
			} else {

			}
		}
		{
			if x_6.Type == 9 && x_6.IntVal == 1830062304 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 1830062304) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0)))
				goto end_branch_4
			} else {

			}
		}
		{
			if x_6.Type == 9 && x_6.IntVal == 3190619783 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 3190619783) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0)))
				goto end_branch_4
			} else {

			}
		}
		{
			if x_6.Type == 9 && x_6.IntVal == 108241190 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0)))) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = Rebox_Main_138441832_749979413(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v))
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_749979413_138441832(v))}
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))))
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = (x_6.Type == 9 && x_6.IntVal == 2066233029) && ((y_7.Type == 9 && y_7.IntVal == 2066233029) && ((gopurs_runtime.Apply2(eqArray6_5_3.V0, func() gopurs_runtime.Value {
				arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Array(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Array(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal) != (0)))
		}
	end_branch_4:
		return gopurs_runtime.Bool(__t4)
	})}))}
}

func Rebox_Main_1039524714_2770120821(in *Constructor_Main_M4[gopurs_runtime.Value, string]) *Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
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

func Rebox_Main_1200068938_138441832(in *Constructor_Data_Tuple_Tuple[int64, []string]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = func() gopurs_runtime.Value {
		arr := in.V1
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Str(v)
		}
		return gopurs_runtime.Array(boxed)
	}()
	return out
}

func Rebox_Main_1302345387_2067946548(in *Constructor_Main_M5[gopurs_runtime.Value, string]) *Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Str(in.V1)
	out.V2 = in.V2
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = in.V5
	out.V6 = in.V6
	out.V7 = in.V7
	return out
}

func Rebox_Main_138441832_1200068938(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, []string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, []string]{}
	out.V0 = in.V0.IntVal
	out.V1 = func() []string {
		arr := *(*[]gopurs_runtime.Value)(in.V1.UnsafePtr)
		unboxed := make([]string, len(arr))
		for i, v := range arr {
			unboxed[i] = v.StrVal()
		}
		return unboxed
	}()
	return out
}

func Rebox_Main_138441832_690815662(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, []int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, []int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(in.V1.UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
	return out
}

func Rebox_Main_138441832_749979413(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]{}
	out.V0 = in.V0.IntVal
	out.V1 = func() []gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(in.V1.UnsafePtr)
		unboxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = v
		}
		return unboxed
	}()
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

func Rebox_Main_147758475_3660606000(in *Constructor_Main_M1[gopurs_runtime.Value, int64]) *Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{}
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

func Rebox_Main_1554627816_4224211191(in *Constructor_Main_M6[gopurs_runtime.Value, string]) *Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{}
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

func Rebox_Main_2124045134_2770120821(in *Constructor_Main_M4[gopurs_runtime.Value, int64]) *Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2628211660_4224211191(in *Constructor_Main_M6[gopurs_runtime.Value, int64]) *Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2785108781_2554376626(in *Constructor_Main_M3[gopurs_runtime.Value, string]) *Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2850066319_2067946548(in *Constructor_Main_M5[gopurs_runtime.Value, int64]) *Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	out.V2 = in.V2
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = in.V5
	out.V6 = in.V6
	out.V7 = in.V7
	return out
}

func Rebox_Main_3716704586_67812977(in *Constructor_Main_M0[gopurs_runtime.Value, int64]) *Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_3741832558_67812977(in *Constructor_Main_M0[gopurs_runtime.Value, string]) *Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
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

func Rebox_Main_3790796878_1161550504(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1871481144(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1939691112(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{}
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

func Rebox_Main_4004653231_3660606000(in *Constructor_Main_M1[gopurs_runtime.Value, string]) *Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4220871112_1521903347(in *Constructor_Main_M2[gopurs_runtime.Value, int64]) *Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4256935660_1521903347(in *Constructor_Main_M2[gopurs_runtime.Value, string]) *Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_690815662_138441832(in *Constructor_Data_Tuple_Tuple[int64, []int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = func() gopurs_runtime.Value {
		arr := in.V1
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Int(v)
		}
		return gopurs_runtime.Array(boxed)
	}()
	return out
}

func Rebox_Main_749979413_138441832(in *Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Array(in.V1)
	return out
}

func Rebox_Main_837768713_2554376626(in *Constructor_Main_M3[gopurs_runtime.Value, int64]) *Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
