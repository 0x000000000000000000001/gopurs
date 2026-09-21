package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))
	})
	return cache_Main_pure
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

var cache_Main_eqArray1 gopurs_runtime.Value
var once_Main_eqArray1 sync.Once

func Get_Main_eqArray1() gopurs_runtime.Value {
	once_Main_eqArray1.Do(func() {
		cache_Main_eqArray1 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main_eqArray1
}

var cache_Main_eqArray2 gopurs_runtime.Value
var once_Main_eqArray2 sync.Once

func Get_Main_eqArray2() gopurs_runtime.Value {
	once_Main_eqArray2.Do(func() {
		cache_Main_eqArray2 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main_eqArray2
}

var cache_Main_eqArray3 gopurs_runtime.Value
var once_Main_eqArray3 sync.Once

func Get_Main_eqArray3() gopurs_runtime.Value {
	once_Main_eqArray3.Do(func() {
		cache_Main_eqArray3 = Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("nested")
		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("zArrayA")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("ignore")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fa")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("arrayIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))
	})
	return cache_Main_eqArray3
}

var cache_Main_M0 gopurs_runtime.Value
var once_Main_M0 sync.Once

func Get_Main_M0() gopurs_runtime.Value {
	once_Main_M0.Do(func() {
		cache_Main_M0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Main_M0
}

var cache_Main_M1 gopurs_runtime.Value
var once_Main_M1 sync.Once

func Get_Main_M1() gopurs_runtime.Value {
	once_Main_M1.Do(func() {
		cache_Main_M1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, (*(*[]gopurs_runtime.Value)((value1).UnsafePtr))}))}
			})
		})
	})
	return cache_Main_M1
}

var cache_Main_M1__2269889115 gopurs_runtime.Value
var once_Main_M1__2269889115 sync.Once

func Get_Main_M1__2269889115() gopurs_runtime.Value {
	once_Main_M1__2269889115.Do(func() {
		cache_Main_M1__2269889115 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M1__2269889115(func() []string {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_1_0_box.UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}(), func() [][]string {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
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
			}())
		})
	})
	return cache_Main_M1__2269889115
}

var cache_Main_M1__203000413 gopurs_runtime.Value
var once_Main_M1__203000413 sync.Once

func Get_Main_M1__203000413() gopurs_runtime.Value {
	once_Main_M1__203000413.Do(func() {
		cache_Main_M1__203000413 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M1__203000413(__eta_norm_1_0_box.StrVal(), func() []string {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}())
		})
	})
	return cache_Main_M1__203000413
}

var cache_Main_M2 gopurs_runtime.Value
var once_Main_M2 sync.Once

func Get_Main_M2() gopurs_runtime.Value {
	once_Main_M2.Do(func() {
		cache_Main_M2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal}))}
		})
	})
	return cache_Main_M2
}

var cache_Main_M2__1740217219 gopurs_runtime.Value
var once_Main_M2__1740217219 sync.Once

func Get_Main_M2__1740217219() gopurs_runtime.Value {
	once_Main_M2__1740217219.Do(func() {
		cache_Main_M2__1740217219 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M2__1740217219(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_M2__1740217219
}

var cache_Main_M2__371339077 gopurs_runtime.Value
var once_Main_M2__371339077 sync.Once

func Get_Main_M2__371339077() gopurs_runtime.Value {
	once_Main_M2__371339077.Do(func() {
		cache_Main_M2__371339077 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M2__371339077(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_M2__371339077
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

var cache_Main_M3__2634390930 gopurs_runtime.Value
var once_Main_M3__2634390930 sync.Once

func Get_Main_M3__2634390930() gopurs_runtime.Value {
	once_Main_M3__2634390930.Do(func() {
		cache_Main_M3__2634390930 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M3__2634390930(__eta_norm_0_0_box)
		})
	})
	return cache_Main_M3__2634390930
}

var cache_Main_M3__1079474578 gopurs_runtime.Value
var once_Main_M3__1079474578 sync.Once

func Get_Main_M3__1079474578() gopurs_runtime.Value {
	once_Main_M3__1079474578.Do(func() {
		cache_Main_M3__1079474578 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_M3__1079474578(__eta_norm_0_0_box)
		})
	})
	return cache_Main_M3__1079474578
}

var cache_Main_M4 gopurs_runtime.Value
var once_Main_M4 sync.Once

func Get_Main_M4() gopurs_runtime.Value {
	once_Main_M4.Do(func() {
		cache_Main_M4 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_M4
}

var cache_Main_M5 gopurs_runtime.Value
var once_Main_M5 sync.Once

func Get_Main_M5() gopurs_runtime.Value {
	once_Main_M5.Do(func() {
		cache_Main_M5 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
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
	return cache_Main_M5
}

var cache_Main_M6 gopurs_runtime.Value
var once_Main_M6 sync.Once

func Get_Main_M6() gopurs_runtime.Value {
	once_Main_M6.Do(func() {
		cache_Main_M6 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(value5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(value6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(value7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1, func() []int64 {
											arr := *(*[]gopurs_runtime.Value)(value2.UnsafePtr)
											unboxed := make([]int64, len(arr))
											for i, v := range arr {
												unboxed[i] = v.IntVal
											}
											return unboxed
										}(), (*(*[]gopurs_runtime.Value)((value3).UnsafePtr)), value4, value5, value6, func() struct {
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
	return cache_Main_M6
}

var cache_Main_M7 gopurs_runtime.Value
var once_Main_M7 sync.Once

func Get_Main_M7() gopurs_runtime.Value {
	once_Main_M7.Do(func() {
		cache_Main_M7 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_M7
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

var cache_Main_foldableM gopurs_runtime.Value
var once_Main_foldableM sync.Once

func Get_Main_foldableM() gopurs_runtime.Value {
	once_Main_foldableM.Do(func() {
		cache_Main_foldableM = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldableM(dictFoldable_0_box)
		})
	})
	return cache_Main_foldableM
}

var cache_Main_traversableM gopurs_runtime.Value
var once_Main_traversableM sync.Once

func Get_Main_traversableM() gopurs_runtime.Value {
	once_Main_traversableM.Do(func() {
		cache_Main_traversableM = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_traversableM(dictTraversable_0_box)
		})
	})
	return cache_Main_traversableM
}

var cache_Main_traversableM1 gopurs_runtime.Value
var once_Main_traversableM1 sync.Once

func Get_Main_traversableM1() gopurs_runtime.Value {
	once_Main_traversableM1.Do(func() {
		cache_Main_traversableM1 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main_traversableM1
}

var cache_Main_eqM gopurs_runtime.Value
var once_Main_eqM sync.Once

func Get_Main_eqM() gopurs_runtime.Value {
	once_Main_eqM.Do(func() {
		cache_Main_eqM = gopurs_runtime.Func4(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, dictEq2_2_box gopurs_runtime.Value, dictEq3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqM(dictEq1_0_box, dictEq_1_box, dictEq2_2_box, dictEq3_3_box)
		})
	})
	return cache_Main_eqM
}

var cache_Main_eqArray4 gopurs_runtime.Value
var once_Main_eqArray4 sync.Once

func Get_Main_eqArray4() gopurs_runtime.Value {
	once_Main_eqArray4.Do(func() {
		cache_Main_eqArray4 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1939691112_3790796878(Rebox_Main_3790796878_1939691112(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("nested")
		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("zArrayA")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("ignore")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fa")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("arrayIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("nested")
		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("zArrayA")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("ignore")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fa")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("fIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("arrayIgnore")
		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))))}
	})
	return cache_Main_eqArray4
}

var cache_Main_traverseStr gopurs_runtime.Value
var once_Main_traverseStr sync.Once

func Get_Main_traverseStr() gopurs_runtime.Value {
	once_Main_traverseStr.Do(func() {
		cache_Main_traverseStr = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_traverseStr(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
		})
	})
	return cache_Main_traverseStr
}

var cache_Main_m0 gopurs_runtime.Value
var once_Main_m0 sync.Once

func Get_Main_m0() gopurs_runtime.Value {
	once_Main_m0.Do(func() {
		cache_Main_m0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}
	})
	return cache_Main_m0
}

var cache_Main_traverseStr__163786110 gopurs_runtime.Value
var once_Main_traverseStr__163786110 sync.Once

func Get_Main_traverseStr__163786110() gopurs_runtime.Value {
	once_Main_traverseStr__163786110.Do(func() {
		cache_Main_traverseStr__163786110 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__163786110(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__163786110
}

var cache_Main_m1 gopurs_runtime.Value
var once_Main_m1 sync.Once

func Get_Main_m1() gopurs_runtime.Value {
	once_Main_m1.Do(func() {
		cache_Main_m1 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, "a", (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := []string{"b", "c"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))})))}
	})
	return cache_Main_m1
}

var cache_Main_traverseStr__1598195487 gopurs_runtime.Value
var once_Main_traverseStr__1598195487 sync.Once

func Get_Main_traverseStr__1598195487() gopurs_runtime.Value {
	once_Main_traverseStr__1598195487.Do(func() {
		cache_Main_traverseStr__1598195487 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__1598195487(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__1598195487
}

var cache_Main_m2 gopurs_runtime.Value
var once_Main_m2 sync.Once

func Get_Main_m2() gopurs_runtime.Value {
	once_Main_m2.Do(func() {
		cache_Main_m2 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})))}
	})
	return cache_Main_m2
}

var cache_Main_traverseStr__2241449532 gopurs_runtime.Value
var once_Main_traverseStr__2241449532 sync.Once

func Get_Main_traverseStr__2241449532() gopurs_runtime.Value {
	once_Main_traverseStr__2241449532.Do(func() {
		cache_Main_traverseStr__2241449532 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__2241449532(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__2241449532
}

var cache_Main_m3 gopurs_runtime.Value
var once_Main_m3 sync.Once

func Get_Main_m3() gopurs_runtime.Value {
	once_Main_m3.Do(func() {
		cache_Main_m3 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_2785108781_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			arr := []string{"a", "b", "c"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()})))}
	})
	return cache_Main_m3
}

var cache_Main_traverseStr__2823203293 gopurs_runtime.Value
var once_Main_traverseStr__2823203293 sync.Once

func Get_Main_traverseStr__2823203293() gopurs_runtime.Value {
	once_Main_traverseStr__2823203293.Do(func() {
		cache_Main_traverseStr__2823203293 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__2823203293(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__2823203293
}

var cache_Main_recordValue gopurs_runtime.Value
var once_Main_recordValue sync.Once

func Get_Main_recordValue() gopurs_runtime.Value {
	once_Main_recordValue.Do(func() {
		cache_Main_recordValue = func() gopurs_runtime.Value {
			orig := struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
			}{"a", []int64{int64(2), int64(3)}, []int64{int64(4)}, []string{"b"}, int64(1), []string{"c"}}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}()
	})
	return cache_Main_recordValue
}

var cache_Main_m4 gopurs_runtime.Value
var once_Main_m4 sync.Once

func Get_Main_m4() gopurs_runtime.Value {
	once_Main_m4.Do(func() {
		cache_Main_m4 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
			} {
				orig := Get_Main_recordValue()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
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
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}()})))}
	})
	return cache_Main_m4
}

var cache_Main_traverseStr__2384857594 gopurs_runtime.Value
var once_Main_traverseStr__2384857594 sync.Once

func Get_Main_traverseStr__2384857594() gopurs_runtime.Value {
	once_Main_traverseStr__2384857594.Do(func() {
		cache_Main_traverseStr__2384857594 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__2384857594(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__2384857594
}

var cache_Main_m5 gopurs_runtime.Value
var once_Main_m5 sync.Once

func Get_Main_m5() gopurs_runtime.Value {
	once_Main_m5.Do(func() {
		cache_Main_m5 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_1302345387_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, string]{1, func() struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				} {
					orig := Get_Main_recordValue()
					_ = orig
					clone := struct {
						a           string
						arrayIgnore []int64
						fIgnore     []int64
						fa          []string
						ignore      int64
						zArrayA     []string
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
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
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
	return cache_Main_m5
}

var cache_Main_traverseStr__3819266971 gopurs_runtime.Value
var once_Main_traverseStr__3819266971 sync.Once

func Get_Main_traverseStr__3819266971() gopurs_runtime.Value {
	once_Main_traverseStr__3819266971.Do(func() {
		cache_Main_traverseStr__3819266971 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__3819266971(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__3819266971
}

var cache_Main_m6 gopurs_runtime.Value
var once_Main_m6 sync.Once

func Get_Main_m6() gopurs_runtime.Value {
	once_Main_m6.Do(func() {
		cache_Main_m6 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer(Rebox_Main_1554627816_4224211191((&Constructor_Main_M6[gopurs_runtime.Value, string]{1, int64(1), "a", func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}(), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := []string{"b"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr)), func() gopurs_runtime.Value {
			arr := []string{"c"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.Array([]gopurs_runtime.Value{}), func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
			} {
				orig := Get_Main_recordValue()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
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
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				} {
					orig := Get_Main_recordValue()
					_ = orig
					clone := struct {
						a           string
						arrayIgnore []int64
						fIgnore     []int64
						fa          []string
						ignore      int64
						zArrayA     []string
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
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
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
	return cache_Main_m6
}

var cache_Main_traverseStr__167553720 gopurs_runtime.Value
var once_Main_traverseStr__167553720 sync.Once

func Get_Main_traverseStr__167553720() gopurs_runtime.Value {
	once_Main_traverseStr__167553720.Do(func() {
		cache_Main_traverseStr__167553720 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__167553720(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__167553720
}

var cache_Main_m7 gopurs_runtime.Value
var once_Main_m7 sync.Once

func Get_Main_m7() gopurs_runtime.Value {
	once_Main_m7.Do(func() {
		cache_Main_m7 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer(Rebox_Main_82800937_961717174((&Constructor_Main_M7[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			arr := [][]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				}
			}{[]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				}
			}{struct {
				nested struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				}
			}{func() struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
			} {
				orig := Get_Main_recordValue()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
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
				clone.zArrayA = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				return clone
			}()}}}
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
								return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
									arr := orig.arrayIgnore
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
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
		}()})))}
	})
	return cache_Main_m7
}

var cache_Main_traverseStr__749307481 gopurs_runtime.Value
var once_Main_traverseStr__749307481 sync.Once

func Get_Main_traverseStr__749307481() gopurs_runtime.Value {
	once_Main_traverseStr__749307481.Do(func() {
		cache_Main_traverseStr__749307481 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_traverseStr__749307481(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_traverseStr__749307481
}

var cache_Main_sequenceStr gopurs_runtime.Value
var once_Main_sequenceStr sync.Once

func Get_Main_sequenceStr() gopurs_runtime.Value {
	once_Main_sequenceStr.Do(func() {
		cache_Main_sequenceStr = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceStr(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
		})
	})
	return cache_Main_sequenceStr
}

var cache_Main_m0_prime_ gopurs_runtime.Value
var once_Main_m0_prime_ sync.Once

func Get_Main_m0_prime_() gopurs_runtime.Value {
	once_Main_m0_prime_.Do(func() {
		cache_Main_m0_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3974551048_67812977(nil))}
	})
	return cache_Main_m0_prime_
}

var cache_Main_sequenceStr__20379839 gopurs_runtime.Value
var once_Main_sequenceStr__20379839 sync.Once

func Get_Main_sequenceStr__20379839() gopurs_runtime.Value {
	once_Main_sequenceStr__20379839.Do(func() {
		cache_Main_sequenceStr__20379839 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__20379839(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__20379839
}

var cache_Main_m1_prime_ gopurs_runtime.Value
var once_Main_m1_prime_ sync.Once

func Get_Main_m1_prime_() gopurs_runtime.Value {
	once_Main_m1_prime_.Do(func() {
		cache_Main_m1_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_1951009097_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, []string]{1, []string{"a"}, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := [][]string{[]string{"b"}, []string{"c"}}
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
		}()).UnsafePtr))})))}
	})
	return cache_Main_m1_prime_
}

var cache_Main_sequenceStr__2491708894 gopurs_runtime.Value
var once_Main_sequenceStr__2491708894 sync.Once

func Get_Main_sequenceStr__2491708894() gopurs_runtime.Value {
	once_Main_sequenceStr__2491708894.Do(func() {
		cache_Main_sequenceStr__2491708894 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__2491708894(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__2491708894
}

var cache_Main_m2_prime_ gopurs_runtime.Value
var once_Main_m2_prime_ sync.Once

func Get_Main_m2_prime_() gopurs_runtime.Value {
	once_Main_m2_prime_.Do(func() {
		cache_Main_m2_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_2857385098_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, []string]{1, int64(0)})))}
	})
	return cache_Main_m2_prime_
}

var cache_Main_sequenceStr__2672135165 gopurs_runtime.Value
var once_Main_sequenceStr__2672135165 sync.Once

func Get_Main_sequenceStr__2672135165() gopurs_runtime.Value {
	once_Main_sequenceStr__2672135165.Do(func() {
		cache_Main_sequenceStr__2672135165 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__2672135165(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__2672135165
}

var cache_Main_m3_prime_ gopurs_runtime.Value
var once_Main_m3_prime_ sync.Once

func Get_Main_m3_prime_() gopurs_runtime.Value {
	once_Main_m3_prime_.Do(func() {
		cache_Main_m3_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_1019686859_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, []string]{1, func() gopurs_runtime.Value {
			arr := [][]string{[]string{"a"}, []string{"b"}, []string{"c"}}
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
		}()})))}
	})
	return cache_Main_m3_prime_
}

var cache_Main_sequenceStr__1137381404 gopurs_runtime.Value
var once_Main_sequenceStr__1137381404 sync.Once

func Get_Main_sequenceStr__1137381404() gopurs_runtime.Value {
	once_Main_sequenceStr__1137381404.Do(func() {
		cache_Main_sequenceStr__1137381404 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__1137381404(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__1137381404
}

var cache_Main_recordValue_prime_ gopurs_runtime.Value
var once_Main_recordValue_prime_ sync.Once

func Get_Main_recordValue_prime_() gopurs_runtime.Value {
	once_Main_recordValue_prime_.Do(func() {
		cache_Main_recordValue_prime_ = func() gopurs_runtime.Value {
			orig := struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
			}{[]string{"a"}, []int64{int64(2), int64(3)}, []int64{int64(4)}, [][]string{[]string{"b"}}, int64(1), [][]string{[]string{"c"}}}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
				arr := orig.a
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.zArrayA
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
			}()})
		}()
	})
	return cache_Main_recordValue_prime_
}

var cache_Main_m4_prime_ gopurs_runtime.Value
var once_Main_m4_prime_ sync.Once

func Get_Main_m4_prime_() gopurs_runtime.Value {
	once_Main_m4_prime_.Do(func() {
		cache_Main_m4_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_2381891596_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, []string]{1, func() gopurs_runtime.Value {
			orig := func() struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
			} {
				orig := Get_Main_recordValue_prime_()
				_ = orig
				clone := struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				}{}
				clone.a = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
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
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.zArrayA = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
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
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
				arr := orig.a
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.zArrayA
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
			}()})
		}()})))}
	})
	return cache_Main_m4_prime_
}

var cache_Main_sequenceStr__2102650939 gopurs_runtime.Value
var once_Main_sequenceStr__2102650939 sync.Once

func Get_Main_sequenceStr__2102650939() gopurs_runtime.Value {
	once_Main_sequenceStr__2102650939.Do(func() {
		cache_Main_sequenceStr__2102650939 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__2102650939(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__2102650939
}

var cache_Main_m5_prime_ gopurs_runtime.Value
var once_Main_m5_prime_ sync.Once

func Get_Main_m5_prime_() gopurs_runtime.Value {
	once_Main_m5_prime_.Do(func() {
		cache_Main_m5_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_358349645_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, []string]{1, func() struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				} {
					orig := Get_Main_recordValue_prime_()
					_ = orig
					clone := struct {
						a           []string
						arrayIgnore []int64
						fIgnore     []int64
						fa          [][]string
						ignore      int64
						zArrayA     [][]string
					}{}
					clone.a = func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()
					clone.arrayIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fa = func() [][]string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
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
					clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
					clone.zArrayA = func() [][]string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
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
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					arr := orig.a
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
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
				}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
					arr := orig.zArrayA
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
	return cache_Main_m5_prime_
}

var cache_Main_sequenceStr__279012698 gopurs_runtime.Value
var once_Main_sequenceStr__279012698 sync.Once

func Get_Main_sequenceStr__279012698() gopurs_runtime.Value {
	once_Main_sequenceStr__279012698.Do(func() {
		cache_Main_sequenceStr__279012698 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__279012698(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__279012698
}

var cache_Main_m6_prime_ gopurs_runtime.Value
var once_Main_m6_prime_ sync.Once

func Get_Main_m6_prime_() gopurs_runtime.Value {
	once_Main_m6_prime_.Do(func() {
		cache_Main_m6_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer(Rebox_Main_1264725646_4224211191((&Constructor_Main_M6[gopurs_runtime.Value, []string]{1, int64(1), []string{"a"}, func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}(), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := [][]string{[]string{"b"}}
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
		}()).UnsafePtr)), func() gopurs_runtime.Value {
			arr := [][]string{[]string{"c"}}
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
		}(), gopurs_runtime.Array([]gopurs_runtime.Value{}), func() gopurs_runtime.Value {
			orig := func() struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
			} {
				orig := Get_Main_recordValue_prime_()
				_ = orig
				clone := struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				}{}
				clone.a = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
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
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.zArrayA = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
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
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
				arr := orig.a
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.zArrayA
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
			}()})
		}(), func() struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				} {
					orig := Get_Main_recordValue_prime_()
					_ = orig
					clone := struct {
						a           []string
						arrayIgnore []int64
						fIgnore     []int64
						fa          [][]string
						ignore      int64
						zArrayA     [][]string
					}{}
					clone.a = func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()
					clone.arrayIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fIgnore = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					clone.fa = func() [][]string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
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
					clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
					clone.zArrayA = func() [][]string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
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
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					arr := orig.a
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.arrayIgnore
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
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
				}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
					arr := orig.zArrayA
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
	return cache_Main_m6_prime_
}

var cache_Main_sequenceStr__459438969 gopurs_runtime.Value
var once_Main_sequenceStr__459438969 sync.Once

func Get_Main_sequenceStr__459438969() gopurs_runtime.Value {
	once_Main_sequenceStr__459438969.Do(func() {
		cache_Main_sequenceStr__459438969 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__459438969(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__459438969
}

var cache_Main_m7_prime_ gopurs_runtime.Value
var once_Main_m7_prime_ sync.Once

func Get_Main_m7_prime_() gopurs_runtime.Value {
	once_Main_m7_prime_.Do(func() {
		cache_Main_m7_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer(Rebox_Main_3721994703_961717174((&Constructor_Main_M7[gopurs_runtime.Value, []string]{1, func() gopurs_runtime.Value {
			arr := [][]struct {
				nested struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				}
			}{[]struct {
				nested struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				}
			}{struct {
				nested struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				}
			}{func() struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
			} {
				orig := Get_Main_recordValue_prime_()
				_ = orig
				clone := struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				}{}
				clone.a = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
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
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.zArrayA = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
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
				return clone
			}()}}}
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
								return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
									arr := orig.a
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := orig.arrayIgnore
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
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
								}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
									arr := orig.zArrayA
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
								}()})
							}())
						}()
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()})))}
	})
	return cache_Main_m7_prime_
}

var cache_Main_sequenceStr__3219652504 gopurs_runtime.Value
var once_Main_sequenceStr__3219652504 sync.Once

func Get_Main_sequenceStr__3219652504() gopurs_runtime.Value {
	once_Main_sequenceStr__3219652504.Do(func() {
		cache_Main_sequenceStr__3219652504 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_sequenceStr__3219652504(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_sequenceStr__3219652504
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m0"))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Boolean
			__local_var_1_1 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("nested")
			})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("zArrayA")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("ignore")
			})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("fa")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("fIgnore")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("arrayIgnore")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("a")
			})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("nested")
			})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("zArrayA")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("ignore")
			})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("fa")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("fIgnore")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("arrayIgnore")
			})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("a")
			})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))})).UnsafePtr))), func() gopurs_runtime.Value {
				arr := []*Constructor_Main_M0[gopurs_runtime.Value, string]{nil}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(v))}
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal) != (0)
			_ = __local_var_1_1
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Bool(__local_var_1_1)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_2 shape=App(Var) bindingType=Any
				__local_var_2_2 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m1"))
				_ = __local_var_2_2
				// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Boolean
				__local_var_3_3 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("nested")
				})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("zArrayA")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("ignore")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fa")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("arrayIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("a")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("nested")
				})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("zArrayA")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("ignore")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fa")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("arrayIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("a")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m1())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m1()})).IntVal) != (0)
				_ = __local_var_3_3
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Bool(__local_var_3_3)), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_4_4 shape=App(Var) bindingType=Any
					__local_var_4_4 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m2"))
					_ = __local_var_4_4
					// TAST (Let): __local_var_5_5 shape=App(Other) bindingType=Boolean
					__local_var_5_5 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("nested")
					})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("zArrayA")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("ignore")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fa")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("arrayIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("a")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("nested")
					})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("zArrayA")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("ignore")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fa")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("arrayIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("a")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})))})).UnsafePtr))), func() gopurs_runtime.Value {
						arr := []*Constructor_Main_M2[gopurs_runtime.Value, string]{(&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347(v))}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0)
					_ = __local_var_5_5
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Bool(__local_var_5_5)), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): __local_var_6_6 shape=App(Var) bindingType=Any
						__local_var_6_6 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m3"))
						_ = __local_var_6_6
						// TAST (Let): __local_var_7_7 shape=App(Other) bindingType=Boolean
						__local_var_7_7 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("nested")
						})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("zArrayA")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("ignore")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fa")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("arrayIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("a")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("nested")
						})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("zArrayA")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("ignore")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fa")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("arrayIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("a")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m3())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m3()})).IntVal) != (0)
						_ = __local_var_7_7
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Bool(__local_var_7_7)), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): __local_var_8_8 shape=App(Var) bindingType=Any
							__local_var_8_8 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m4"))
							_ = __local_var_8_8
							// TAST (Let): __local_var_9_9 shape=App(Other) bindingType=Boolean
							__local_var_9_9 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("nested")
							})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("zArrayA")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("ignore")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fa")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("arrayIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("a")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("nested")
							})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("zArrayA")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("ignore")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fa")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("arrayIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("a")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
								orig := func() struct {
									a           string
									arrayIgnore []int64
									fIgnore     []int64
									fa          []string
									ignore      int64
									zArrayA     []string
								} {
									orig := Get_Main_recordValue()
									_ = orig
									clone := struct {
										a           string
										arrayIgnore []int64
										fIgnore     []int64
										fa          []string
										ignore      int64
										zArrayA     []string
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
								return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
									arr := orig.arrayIgnore
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
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
									arr := orig.zArrayA
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Str(v)
									}
									return gopurs_runtime.Array(boxed)
								}()})
							}()})))})).UnsafePtr))), func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M4[gopurs_runtime.Value, string]{(&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
									orig := func() struct {
										a           string
										arrayIgnore []int64
										fIgnore     []int64
										fa          []string
										ignore      int64
										zArrayA     []string
									} {
										orig := Get_Main_recordValue()
										_ = orig
										clone := struct {
											a           string
											arrayIgnore []int64
											fIgnore     []int64
											fa          []string
											ignore      int64
											zArrayA     []string
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
									return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
										arr := orig.arrayIgnore
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Int(v)
										}
										return gopurs_runtime.Array(boxed)
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
										arr := orig.zArrayA
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Str(v)
										}
										return gopurs_runtime.Array(boxed)
									}()})
								}()})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821(v))}
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0)
							_ = __local_var_9_9
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_8_8, gopurs_runtime.Bool(__local_var_9_9)), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								// TAST (Let): __local_var_10_10 shape=App(Var) bindingType=Any
								__local_var_10_10 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m5"))
								_ = __local_var_10_10
								// TAST (Let): __local_var_11_11 shape=App(Other) bindingType=Boolean
								__local_var_11_11 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("nested")
								})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("zArrayA")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("ignore")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fa")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("arrayIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("a")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("nested")
								})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("zArrayA")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("ignore")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fa")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("arrayIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("a")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m5())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m5()})).IntVal) != (0)
								_ = __local_var_11_11
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_10_10, gopurs_runtime.Bool(__local_var_11_11)), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): __local_var_12_12 shape=App(Var) bindingType=Any
									__local_var_12_12 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m6"))
									_ = __local_var_12_12
									// TAST (Let): __local_var_13_13 shape=App(Other) bindingType=Boolean
									__local_var_13_13 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("nested")
									})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("zArrayA")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("ignore")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fa")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("arrayIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("a")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("nested")
									})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("zArrayA")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("ignore")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fa")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("arrayIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("a")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m6())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m6()})).IntVal) != (0)
									_ = __local_var_13_13
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_12_12, gopurs_runtime.Bool(__local_var_13_13)), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_14_14 shape=App(Var) bindingType=Any
										__local_var_14_14 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m7"))
										_ = __local_var_14_14
										// TAST (Let): __local_var_15_15 shape=App(Other) bindingType=Boolean
										__local_var_15_15 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("nested")
										})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("zArrayA")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("ignore")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fa")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("arrayIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("a")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("nested")
										})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("zArrayA")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("ignore")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fa")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("arrayIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("a")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m7())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m7()})).IntVal) != (0)
										_ = __local_var_15_15
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_14_14, gopurs_runtime.Bool(__local_var_15_15)), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_16_16 shape=App(Var) bindingType=Any
											__local_var_16_16 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m0"))
											_ = __local_var_16_16
											// TAST (Let): __local_var_17_17 shape=App(Other) bindingType=Boolean
											__local_var_17_17 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("nested")
											})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("zArrayA")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("ignore")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fa")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("arrayIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("a")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("nested")
											})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("zArrayA")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("ignore")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fa")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("arrayIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("a")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3974551048_67812977(nil))})).UnsafePtr))), func() gopurs_runtime.Value {
												arr := []*Constructor_Main_M0[gopurs_runtime.Value, string]{nil}
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(v))}
												}
												return gopurs_runtime.Array(boxed)
											}()).IntVal) != (0)
											_ = __local_var_17_17
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_16_16, gopurs_runtime.Bool(__local_var_17_17)), gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
												// TAST (Let): __local_var_18_18 shape=App(Var) bindingType=Any
												__local_var_18_18 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m1"))
												_ = __local_var_18_18
												// TAST (Let): __local_var_19_19 shape=App(Other) bindingType=Boolean
												__local_var_19_19 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("nested")
												})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("zArrayA")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("ignore")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fa")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("arrayIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("a")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("nested")
												})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("zArrayA")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("ignore")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fa")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("arrayIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("a")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m1_prime_())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m1()})).IntVal) != (0)
												_ = __local_var_19_19
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_18_18, gopurs_runtime.Bool(__local_var_19_19)), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
													// TAST (Let): __local_var_20_20 shape=App(Var) bindingType=Any
													__local_var_20_20 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m2"))
													_ = __local_var_20_20
													// TAST (Let): __local_var_21_21 shape=App(Other) bindingType=Boolean
													__local_var_21_21 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("nested")
													})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("zArrayA")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("ignore")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fa")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("arrayIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("a")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("nested")
													})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("zArrayA")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("ignore")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fa")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("arrayIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("a")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_2857385098_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, []string]{1, int64(0)})))})).UnsafePtr))), func() gopurs_runtime.Value {
														arr := []*Constructor_Main_M2[gopurs_runtime.Value, string]{(&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})}
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347(v))}
														}
														return gopurs_runtime.Array(boxed)
													}()).IntVal) != (0)
													_ = __local_var_21_21
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_20_20, gopurs_runtime.Bool(__local_var_21_21)), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														// TAST (Let): __local_var_22_22 shape=App(Var) bindingType=Any
														__local_var_22_22 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m3"))
														_ = __local_var_22_22
														// TAST (Let): __local_var_23_23 shape=App(Other) bindingType=Boolean
														__local_var_23_23 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("nested")
														})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("zArrayA")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("ignore")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fa")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("arrayIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("a")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("nested")
														})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("zArrayA")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("ignore")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fa")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("arrayIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("a")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m3_prime_())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m3()})).IntVal) != (0)
														_ = __local_var_23_23
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_22_22, gopurs_runtime.Bool(__local_var_23_23)), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															// TAST (Let): __local_var_24_24 shape=App(Var) bindingType=Any
															__local_var_24_24 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m4"))
															_ = __local_var_24_24
															// TAST (Let): __local_var_25_25 shape=App(Other) bindingType=Boolean
															__local_var_25_25 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("nested")
															})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("zArrayA")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("ignore")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fa")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("arrayIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("a")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("nested")
															})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("zArrayA")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("ignore")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fa")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("arrayIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("a")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_2381891596_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, []string]{1, func() gopurs_runtime.Value {
																orig := func() struct {
																	a           []string
																	arrayIgnore []int64
																	fIgnore     []int64
																	fa          [][]string
																	ignore      int64
																	zArrayA     [][]string
																} {
																	orig := Get_Main_recordValue_prime_()
																	_ = orig
																	clone := struct {
																		a           []string
																		arrayIgnore []int64
																		fIgnore     []int64
																		fa          [][]string
																		ignore      int64
																		zArrayA     [][]string
																	}{}
																	clone.a = func() []string {
																		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
																		unboxed := make([]string, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v.StrVal()
																		}
																		return unboxed
																	}()
																	clone.arrayIgnore = func() []int64 {
																		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
																		unboxed := make([]int64, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v.IntVal
																		}
																		return unboxed
																	}()
																	clone.fIgnore = func() []int64 {
																		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
																		unboxed := make([]int64, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v.IntVal
																		}
																		return unboxed
																	}()
																	clone.fa = func() [][]string {
																		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
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
																	clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
																	clone.zArrayA = func() [][]string {
																		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
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
																	return clone
																}()
																_ = orig
																return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
																	arr := orig.a
																	boxed := make([]gopurs_runtime.Value, len(arr))
																	for i, v := range arr {
																		boxed[i] = gopurs_runtime.Str(v)
																	}
																	return gopurs_runtime.Array(boxed)
																}(), func() gopurs_runtime.Value {
																	arr := orig.arrayIgnore
																	boxed := make([]gopurs_runtime.Value, len(arr))
																	for i, v := range arr {
																		boxed[i] = gopurs_runtime.Int(v)
																	}
																	return gopurs_runtime.Array(boxed)
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
																}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
																	arr := orig.zArrayA
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
																}()})
															}()})))})).UnsafePtr))), func() gopurs_runtime.Value {
																arr := []*Constructor_Main_M4[gopurs_runtime.Value, string]{(&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
																	orig := func() struct {
																		a           string
																		arrayIgnore []int64
																		fIgnore     []int64
																		fa          []string
																		ignore      int64
																		zArrayA     []string
																	} {
																		orig := Get_Main_recordValue()
																		_ = orig
																		clone := struct {
																			a           string
																			arrayIgnore []int64
																			fIgnore     []int64
																			fa          []string
																			ignore      int64
																			zArrayA     []string
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
																	return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
																		arr := orig.arrayIgnore
																		boxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			boxed[i] = gopurs_runtime.Int(v)
																		}
																		return gopurs_runtime.Array(boxed)
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
																		arr := orig.zArrayA
																		boxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			boxed[i] = gopurs_runtime.Str(v)
																		}
																		return gopurs_runtime.Array(boxed)
																	}()})
																}()})}
																boxed := make([]gopurs_runtime.Value, len(arr))
																for i, v := range arr {
																	boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821(v))}
																}
																return gopurs_runtime.Array(boxed)
															}()).IntVal) != (0)
															_ = __local_var_25_25
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_24_24, gopurs_runtime.Bool(__local_var_25_25)), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																// TAST (Let): __local_var_26_26 shape=App(Var) bindingType=Any
																__local_var_26_26 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m5"))
																_ = __local_var_26_26
																// TAST (Let): __local_var_27_27 shape=App(Other) bindingType=Boolean
																__local_var_27_27 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("nested")
																})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("zArrayA")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("ignore")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fa")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("arrayIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("a")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("nested")
																})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("zArrayA")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("ignore")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fa")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("arrayIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("a")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m5_prime_())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m5()})).IntVal) != (0)
																_ = __local_var_27_27
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_26_26, gopurs_runtime.Bool(__local_var_27_27)), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																	// TAST (Let): __local_var_28_28 shape=App(Var) bindingType=Any
																	__local_var_28_28 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m6"))
																	_ = __local_var_28_28
																	// TAST (Let): __local_var_29_29 shape=App(Other) bindingType=Boolean
																	__local_var_29_29 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("nested")
																	})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("zArrayA")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("ignore")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fa")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("arrayIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("a")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("nested")
																	})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("zArrayA")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("ignore")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fa")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("arrayIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("a")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m6_prime_())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m6()})).IntVal) != (0)
																	_ = __local_var_29_29
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_28_28, gopurs_runtime.Bool(__local_var_29_29)), gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																		// TAST (Let): __local_var_30_30 shape=App(Var) bindingType=Any
																		__local_var_30_30 := gopurs_runtime.Apply(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m7"))
																		_ = __local_var_30_30
																		// TAST (Let): __local_var_31_31 shape=App(Other) bindingType=Boolean
																		__local_var_31_31 := (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("nested")
																		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("zArrayA")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("ignore")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fa")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("arrayIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("a")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("nested")
																		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("zArrayA")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("ignore")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fa")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("arrayIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("a")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m7_prime_())).UnsafePtr))), gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m7()})).IntVal) != (0)
																		_ = __local_var_31_31
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(__local_var_30_30, gopurs_runtime.Bool(__local_var_31_31)), gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																		}))
																	}))
																}))
															}))
														}))
													}))
												}))
											}))
										}))
									}))
								}))
							}))
						}))
					}))
				}))
			}))
		}()
	})
	return cache_Main_main
}

type Constructor_Main_M0[T_f any, T_a any] struct {
	Rc uint32
}

type Constructor_Main_M1[T_f any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 []gopurs_runtime.Value
}

type Constructor_Main_M2[T_f any, T_a any] struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_M3[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M4[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M5[T_f any, T_a any] struct {
	Rc uint32
	V0 struct {
		nested gopurs_runtime.Value
	}
}

type Constructor_Main_M6[T_f any, T_a any] struct {
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

type Constructor_Main_M7[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_M1__2269889115(__eta_norm_1_0_loop []string, __eta_norm_0_1_loop [][]string) gopurs_runtime.Value {
M1__2269889115:
	for {
		if false {
			continue M1__2269889115
		}
		var __eta_norm_1_0 []string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 [][]string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_1951009097_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, []string]{1, __eta_norm_1_0, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := __eta_norm_0_1
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
		}()).UnsafePtr))})))}
	}
}

func Call_Main_M1__203000413(__eta_norm_1_0_loop string, __eta_norm_0_1_loop []string) gopurs_runtime.Value {
M1__203000413:
	for {
		if false {
			continue M1__203000413
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 []string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, __eta_norm_1_0, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := __eta_norm_0_1
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))})))}
	}
}

func Call_Main_M2__1740217219(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
M2__1740217219:
	for {
		if false {
			continue M2__1740217219
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_2857385098_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, []string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_M2__371339077(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
M2__371339077:
	for {
		if false {
			continue M2__371339077
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_M3__2634390930(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
M3__2634390930:
	for {
		if false {
			continue M3__2634390930
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_1019686859_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, []string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_M3__1079474578(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
M3__1079474578:
	for {
		if false {
			continue M3__1079474578
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_2785108781_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_functorM(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if m_2.Type == 9 && m_2.IntVal == 3852365315 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 769986722 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(f_1, v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr))}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2727978561 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 1830062304 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 3190619783 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap5 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "zArrayA")).UnsafePtr)))
					_ = arr_val_arrayMap5
					arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
					_ = arr_go_arrayMap5
					res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
					_ = res_go_arrayMap5
					for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
						res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_1, v_arrayMap5)
					}
					return gopurs_runtime.Array(res_go_arrayMap5)
				}()).UnsafePtr))))}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 108241190 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
					nested gopurs_runtime.Value
				} {
					orig := func() gopurs_runtime.Value {
						orig := func() struct {
							nested struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								fIgnore     gopurs_runtime.Value
								fa          gopurs_runtime.Value
								ignore      int64
								zArrayA     []gopurs_runtime.Value
							}
						} {
							originalRecord := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0
							_ = originalRecord
							_ = originalRecord
							var clone struct {
								nested struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
								}
							}
							clone.nested = func() struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								fIgnore     gopurs_runtime.Value
								fa          gopurs_runtime.Value
								ignore      int64
								zArrayA     []gopurs_runtime.Value
							} {
								orig := gopurs_runtime.RecordUpdate3((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
									arr_val_arrayMap6 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_1, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}()).UnsafePtr))))
								_ = orig
								clone := struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
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
								clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
								clone.fa = gopurs_runtime.RecordGet(orig, "fa")
								clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
								clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
								return clone
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
							orig := orig.nested
							_ = orig
							return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
								arr := orig.arrayIgnore
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
						}())
					}()
					_ = orig
					clone := struct {
						nested gopurs_runtime.Value
					}{}
					clone.nested = gopurs_runtime.RecordGet(orig, "nested")
					return clone
				}()}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2066233029 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3)
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(f_1, v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap5 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "zArrayA")).UnsafePtr)))
					_ = arr_val_arrayMap5
					arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
					_ = arr_go_arrayMap5
					res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
					_ = res_go_arrayMap5
					for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
						res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_1, v_arrayMap5)
					}
					return gopurs_runtime.Array(res_go_arrayMap5)
				}()).UnsafePtr)))), func() struct {
					nested gopurs_runtime.Value
				} {
					orig := func() gopurs_runtime.Value {
						orig := func() struct {
							nested struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								fIgnore     gopurs_runtime.Value
								fa          gopurs_runtime.Value
								ignore      int64
								zArrayA     []gopurs_runtime.Value
							}
						} {
							originalRecord := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7
							_ = originalRecord
							_ = originalRecord
							var clone struct {
								nested struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
								}
							}
							clone.nested = func() struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								fIgnore     gopurs_runtime.Value
								fa          gopurs_runtime.Value
								ignore      int64
								zArrayA     []gopurs_runtime.Value
							} {
								orig := gopurs_runtime.RecordUpdate3((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
									arr_val_arrayMap6 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_1, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}()).UnsafePtr))))
								_ = orig
								clone := struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
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
								clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
								clone.fa = gopurs_runtime.RecordGet(orig, "fa")
								clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
								clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
								return clone
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
							orig := orig.nested
							_ = orig
							return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
								arr := orig.arrayIgnore
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
						}())
					}()
					_ = orig
					clone := struct {
						nested gopurs_runtime.Value
					}{}
					clone.nested = gopurs_runtime.RecordGet(orig, "nested")
					return clone
				}()}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 1168316772 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						orig := func() struct {
							nested struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								fIgnore     gopurs_runtime.Value
								fa          gopurs_runtime.Value
								ignore      int64
								zArrayA     []gopurs_runtime.Value
							}
						} {
							orig := gopurs_runtime.RecordUpdate1(v1_3, "nested", func() gopurs_runtime.Value {
								orig := func() struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
								} {
									orig := gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_3, "nested"), "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
										arr_val_arrayMap9 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA")).UnsafePtr)))
										_ = arr_val_arrayMap9
										arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
										_ = arr_go_arrayMap9
										res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
										_ = res_go_arrayMap9
										for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
											res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(f_1, v_arrayMap9)
										}
										return gopurs_runtime.Array(res_go_arrayMap9)
									}()).UnsafePtr))))
									_ = orig
									clone := struct {
										a           gopurs_runtime.Value
										arrayIgnore []int64
										fIgnore     gopurs_runtime.Value
										fa          gopurs_runtime.Value
										ignore      int64
										zArrayA     []gopurs_runtime.Value
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
									clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
									clone.fa = gopurs_runtime.RecordGet(orig, "fa")
									clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
									clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
									arr := orig.arrayIgnore
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
							}())
							_ = orig
							clone := struct {
								nested struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
								}
							}{}
							clone.nested = func() struct {
								a           gopurs_runtime.Value
								arrayIgnore []int64
								fIgnore     gopurs_runtime.Value
								fa          gopurs_runtime.Value
								ignore      int64
								zArrayA     []gopurs_runtime.Value
							} {
								orig := gopurs_runtime.RecordGet(orig, "nested")
								_ = orig
								clone := struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
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
								clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
								clone.fa = gopurs_runtime.RecordGet(orig, "fa")
								clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
								clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
								return clone
							}()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
							orig := orig.nested
							_ = orig
							return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
								arr := orig.arrayIgnore
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
						}())
					}()
				})), (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	})}))}
}

func Call_Main_foldableM(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_2_0 shape=App(Var) bindingType=(TypeVar m$scope67)
		mempty_2_0 := Call_Data_Monoid_mempty(dictMonoid_1)
		_ = mempty_2_0
		// TAST (Let): Semigroup0_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope67)])
		Semigroup0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_3_1
		return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if m_5.Type == 9 && m_5.IntVal == 3852365315 {
					__t2 = mempty_2_0
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 769986722 {
					__t2 = gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1)))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 2727978561 {
					__t2 = mempty_2_0
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 1830062304 {
					__t2 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0)
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 3190619783 {
					__t2 = gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "zArrayA")).UnsafePtr))))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 108241190 {
					__t2 = gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr))))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 2066233029 {
					__t2 = gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V3)), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V4), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "zArrayA")).UnsafePtr)))), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr))))))))))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 1168316772 {
					__t2 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "zArrayA")).UnsafePtr))))))
					})), (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0)
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
	}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 gopurs_runtime.Value
		{
			if m_3.Type == 9 && m_3.IntVal == 3852365315 {
				__t3 = z_2
				goto end_branch_3
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 769986722 {
				__t3 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(f_1, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_3
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2727978561 {
				__t3 = z_2
				goto end_branch_3
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1830062304 {
				__t3 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, z_2, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				goto end_branch_3
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3190619783 {
				__t3 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA")).UnsafePtr)))
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "fa"))
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(f_1, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_3
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 108241190 {
				__t3 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "fa"))
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(f_1, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_3
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2066233029 {
				__t3 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, func() gopurs_runtime.Value {
						arr_val_foldlArray6 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "zArrayA")).UnsafePtr)))
						_ = arr_val_foldlArray6
						res_go_foldlArray6 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, func() gopurs_runtime.Value {
							arr_val_foldlArray10 := gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)
							_ = arr_val_foldlArray10
							res_go_foldlArray10 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)
							_ = res_go_foldlArray10
							arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
							_ = arr_go_foldlArray10
							for _, v_foldlArray10 := range *arr_go_foldlArray10 {
								res_go_foldlArray10 = gopurs_runtime.Apply2(f_1, res_go_foldlArray10, v_foldlArray10)
							}
							return res_go_foldlArray10
						}(), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V4), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "fa"))
						_ = res_go_foldlArray6
						arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
						_ = arr_go_foldlArray6
						for _, v_foldlArray6 := range *arr_go_foldlArray6 {
							res_go_foldlArray6 = gopurs_runtime.Apply2(f_1, res_go_foldlArray6, v_foldlArray6)
						}
						return res_go_foldlArray6
					}(), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "fa"))
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(f_1, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_3
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1168316772 {
				__t3 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr_val_foldlArray6 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "zArrayA")).UnsafePtr)))
						_ = arr_val_foldlArray6
						res_go_foldlArray6 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, v1_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "fa"))
						_ = res_go_foldlArray6
						arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
						_ = arr_go_foldlArray6
						for _, v_foldlArray6 := range *arr_go_foldlArray6 {
							res_go_foldlArray6 = gopurs_runtime.Apply2(f_1, res_go_foldlArray6, v_foldlArray6)
						}
						return res_go_foldlArray6
					}()
				})), z_2, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_3:
		return __t3
	}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t5 gopurs_runtime.Value
		{
			if m_3.Type == 9 && m_3.IntVal == 3852365315 {
				__t5 = z_2
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 769986722 {
				__t5 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)))
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2727978561 {
				__t5 = z_2
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1830062304 {
				__t5 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, z_2, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3190619783 {
				__t5 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "fa")))
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 108241190 {
				__t5 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "fa")))
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2066233029 {
				__t5 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "fa"))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "fa"))), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V4), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)))
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1168316772 {
				// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
				__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, v2_5, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "fa")))
				}))
				_ = __local_var_4_4
				__t5 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(__local_var_4_4, a_6, b_5)
				}), z_2, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				goto end_branch_5
			} else {

			}
		}
		{
			__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_5:
		return __t5
	})}))}
}

func Call_Main_traversableM(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
traversableM:
	for {
		if false {
			continue traversableM
		}
		var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
		_ = dictTraversable_0
		// TAST (Let): functorM1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(ADT ["Main","M"] [(TypeVar f$scope1)])])
		functorM1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})))
		_ = functorM1_1_0
		// TAST (Let): foldableM1_2_1 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(ADT ["Main","M"] [(TypeVar f$scope1)])])
		foldableM1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})))
		_ = foldableM1_2_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableM1_2_1)}
		}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorM1_1_0)}
		}), gopurs_runtime.Func2(func(dictApplicative_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(dictTraversable_0), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), v_4)
		}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_4_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope10)])
			Apply0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_4_2
			// TAST (Let): Functor0_5_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope10)])
			Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_5_3
			return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t11 gopurs_runtime.Value
				{
					if m_7.Type == 9 && m_7.IntVal == 3852365315 {
						__t11 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)})
						goto end_branch_11
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 769986722 {
						__t11 = gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func2(func(v2_8 gopurs_runtime.Value, v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_8, (*(*[]gopurs_runtime.Value)((v3_9).UnsafePtr))}))}
						}), gopurs_runtime.Apply(f_6, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1)))
						goto end_branch_11
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 2727978561 {
						__t11 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0}))})
						goto end_branch_11
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 1830062304 {
						__t11 = gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, v1_8}))}
						}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0))
						goto end_branch_11
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 3190619783 {
						// TAST (Let): __local_var_8_4 shape=Other bindingType=Any
						__local_var_8_4 := (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
						_ = __local_var_8_4
						__t11 = gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func3(func(v1_9 gopurs_runtime.Value, v2_10 gopurs_runtime.Value, v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdate3(__local_var_8_4, "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v3_11).UnsafePtr))))}))}
						}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_8_4, "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_8_4, "fa"))), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_8_4, "zArrayA")).UnsafePtr)))))
						goto end_branch_11
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 108241190 {
						// TAST (Let): __local_var_8_5 shape=Other bindingType=Any
						__local_var_8_5 := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
						_ = __local_var_8_5
						__t11 = gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func3(func(v1_9 gopurs_runtime.Value, v2_10 gopurs_runtime.Value, v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
								nested gopurs_runtime.Value
							} {
								orig := func() gopurs_runtime.Value {
									orig := func() struct {
										nested struct {
											a           gopurs_runtime.Value
											arrayIgnore []int64
											fIgnore     gopurs_runtime.Value
											fa          gopurs_runtime.Value
											ignore      int64
											zArrayA     []gopurs_runtime.Value
										}
									} {
										originalRecord := __local_var_8_5
										_ = originalRecord
										_ = originalRecord
										var clone struct {
											nested struct {
												a           gopurs_runtime.Value
												arrayIgnore []int64
												fIgnore     gopurs_runtime.Value
												fa          gopurs_runtime.Value
												ignore      int64
												zArrayA     []gopurs_runtime.Value
											}
										}
										clone.nested = func() struct {
											a           gopurs_runtime.Value
											arrayIgnore []int64
											fIgnore     gopurs_runtime.Value
											fa          gopurs_runtime.Value
											ignore      int64
											zArrayA     []gopurs_runtime.Value
										} {
											orig := gopurs_runtime.RecordUpdate3(__local_var_8_5.nested, "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v3_11).UnsafePtr))))
											_ = orig
											clone := struct {
												a           gopurs_runtime.Value
												arrayIgnore []int64
												fIgnore     gopurs_runtime.Value
												fa          gopurs_runtime.Value
												ignore      int64
												zArrayA     []gopurs_runtime.Value
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
											clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
											clone.fa = gopurs_runtime.RecordGet(orig, "fa")
											clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
											clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
											return clone
										}()
										return clone
									}()
									_ = orig
									return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
										orig := orig.nested
										_ = orig
										return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
											arr := orig.arrayIgnore
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
									}())
								}()
								_ = orig
								clone := struct {
									nested gopurs_runtime.Value
								}{}
								clone.nested = gopurs_runtime.RecordGet(orig, "nested")
								return clone
							}()}))}
						}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_8_5.nested, "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_8_5.nested, "fa"))), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_8_5.nested, "zArrayA")).UnsafePtr)))))
						goto end_branch_11
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 2066233029 {
						// TAST (Let): __local_var_8_6 shape=Other bindingType=Any
						__local_var_8_6 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
						_ = __local_var_8_6
						// TAST (Let): __local_var_9_7 shape=Other bindingType=Any
						__local_var_9_7 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V2
						_ = __local_var_9_7
						// TAST (Let): __local_var_10_8 shape=Other bindingType=Any
						__local_var_10_8 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V5
						_ = __local_var_10_8
						// TAST (Let): __local_var_11_9 shape=Other bindingType=Any
						__local_var_11_9 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V6
						_ = __local_var_11_9
						// TAST (Let): __local_var_12_10 shape=Other bindingType=Any
						__local_var_12_10 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V7
						_ = __local_var_12_10
						__t11 = gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func5(func(v8_13 gopurs_runtime.Value, v9_14 gopurs_runtime.Value, v10_15 gopurs_runtime.Value, v11_16 gopurs_runtime.Value, v12_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func4(func(v13_18 gopurs_runtime.Value, v14_19 gopurs_runtime.Value, v15_20 gopurs_runtime.Value, v16_21 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_8_6, v8_13, __local_var_9_7, (*(*[]gopurs_runtime.Value)((v9_14).UnsafePtr)), v10_15, __local_var_10_8, gopurs_runtime.RecordUpdate3(__local_var_11_9, "a", v11_16, "fa", v12_17, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v13_18).UnsafePtr)))), func() struct {
									nested gopurs_runtime.Value
								} {
									orig := func() gopurs_runtime.Value {
										orig := func() struct {
											nested struct {
												a           gopurs_runtime.Value
												arrayIgnore []int64
												fIgnore     gopurs_runtime.Value
												fa          gopurs_runtime.Value
												ignore      int64
												zArrayA     []gopurs_runtime.Value
											}
										} {
											originalRecord := __local_var_12_10
											_ = originalRecord
											_ = originalRecord
											var clone struct {
												nested struct {
													a           gopurs_runtime.Value
													arrayIgnore []int64
													fIgnore     gopurs_runtime.Value
													fa          gopurs_runtime.Value
													ignore      int64
													zArrayA     []gopurs_runtime.Value
												}
											}
											clone.nested = func() struct {
												a           gopurs_runtime.Value
												arrayIgnore []int64
												fIgnore     gopurs_runtime.Value
												fa          gopurs_runtime.Value
												ignore      int64
												zArrayA     []gopurs_runtime.Value
											} {
												orig := gopurs_runtime.RecordUpdate3(__local_var_12_10.nested, "a", v14_19, "fa", v15_20, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v16_21).UnsafePtr))))
												_ = orig
												clone := struct {
													a           gopurs_runtime.Value
													arrayIgnore []int64
													fIgnore     gopurs_runtime.Value
													fa          gopurs_runtime.Value
													ignore      int64
													zArrayA     []gopurs_runtime.Value
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
												clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
												clone.fa = gopurs_runtime.RecordGet(orig, "fa")
												clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
												clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
												return clone
											}()
											return clone
										}()
										_ = orig
										return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
											orig := orig.nested
											_ = orig
											return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
												arr := orig.arrayIgnore
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Int(v)
												}
												return gopurs_runtime.Array(boxed)
											}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
										}())
									}()
									_ = orig
									clone := struct {
										nested gopurs_runtime.Value
									}{}
									clone.nested = gopurs_runtime.RecordGet(orig, "nested")
									return clone
								}()}))}
							})
						}), gopurs_runtime.Apply(f_6, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1)), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V3))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V4)), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_11_9, "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_11_9, "fa"))), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_11_9, "zArrayA")).UnsafePtr))))), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_12_10.nested, "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_12_10.nested, "fa"))), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_12_10.nested, "zArrayA")).UnsafePtr)))))
						goto end_branch_11
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 1168316772 {
						__t11 = gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value]{1, v1_8}))}
						}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Apply0_4_2.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Func3(func(v2_9 gopurs_runtime.Value, v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return func() gopurs_runtime.Value {
									orig := func() struct {
										nested struct {
											a           gopurs_runtime.Value
											arrayIgnore []int64
											fIgnore     gopurs_runtime.Value
											fa          gopurs_runtime.Value
											ignore      int64
											zArrayA     []gopurs_runtime.Value
										}
									} {
										orig := gopurs_runtime.RecordUpdate1(v1_8, "nested", func() gopurs_runtime.Value {
											orig := func() struct {
												a           gopurs_runtime.Value
												arrayIgnore []int64
												fIgnore     gopurs_runtime.Value
												fa          gopurs_runtime.Value
												ignore      int64
												zArrayA     []gopurs_runtime.Value
											} {
												orig := gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_8, "nested"), "a", v2_9, "fa", v3_10, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v4_11).UnsafePtr))))
												_ = orig
												clone := struct {
													a           gopurs_runtime.Value
													arrayIgnore []int64
													fIgnore     gopurs_runtime.Value
													fa          gopurs_runtime.Value
													ignore      int64
													zArrayA     []gopurs_runtime.Value
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
												clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
												clone.fa = gopurs_runtime.RecordGet(orig, "fa")
												clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
												clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
												return clone
											}()
											_ = orig
											return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
												arr := orig.arrayIgnore
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Int(v)
												}
												return gopurs_runtime.Array(boxed)
											}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
										}())
										_ = orig
										clone := struct {
											nested struct {
												a           gopurs_runtime.Value
												arrayIgnore []int64
												fIgnore     gopurs_runtime.Value
												fa          gopurs_runtime.Value
												ignore      int64
												zArrayA     []gopurs_runtime.Value
											}
										}{}
										clone.nested = func() struct {
											a           gopurs_runtime.Value
											arrayIgnore []int64
											fIgnore     gopurs_runtime.Value
											fa          gopurs_runtime.Value
											ignore      int64
											zArrayA     []gopurs_runtime.Value
										} {
											orig := gopurs_runtime.RecordGet(orig, "nested")
											_ = orig
											clone := struct {
												a           gopurs_runtime.Value
												arrayIgnore []int64
												fIgnore     gopurs_runtime.Value
												fa          gopurs_runtime.Value
												ignore      int64
												zArrayA     []gopurs_runtime.Value
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
											clone.fIgnore = gopurs_runtime.RecordGet(orig, "fIgnore")
											clone.fa = gopurs_runtime.RecordGet(orig, "fa")
											clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
											clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
											return clone
										}()
										return clone
									}()
									_ = orig
									return gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
										orig := orig.nested
										_ = orig
										return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
											arr := orig.arrayIgnore
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
									}())
								}()
							}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "fa"))), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "zArrayA")).UnsafePtr)))))
						})), (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0))
						goto end_branch_11
					} else {

					}
				}
				{
					__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_11:
				return __t11
			})
		})}))}
	}
}

func Call_Main_eqM(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value, dictEq2_2_loop gopurs_runtime.Value, dictEq3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	var dictEq2_2 gopurs_runtime.Value = dictEq2_2_loop
	_ = dictEq2_2
	var dictEq3_3 gopurs_runtime.Value = dictEq3_3_loop
	_ = dictEq3_3
	// TAST (Let): eqArray5_4_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a$scope71))])
	eqArray5_4_0 := Rebox_Main_3790796878_1939691112(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(dictEq3_3)))
	_ = eqArray5_4_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 bool
		{
			if x_5.Type == 9 && x_5.IntVal == 3852365315 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 3852365315)
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 769986722 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 769986722) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V1)).IntVal) != (0)))
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 2727978561 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 2727978561) && (((*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0) == ((*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0))
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 1830062304 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 1830062304) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal) != (0))
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 3190619783 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 3190619783) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "zArrayA")).UnsafePtr)))).IntVal) != (0)))
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 108241190 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 108241190) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))).IntVal) != (0)))
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 2066233029 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0) == ((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V5, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "zArrayA")).UnsafePtr)))).IntVal) != (0)))) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "arrayIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "arrayIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))).IntVal) != (0))))
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = (x_5.Type == 9 && x_5.IntVal == 1168316772) && ((y_6.Type == 9 && y_6.IntVal == 1168316772) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq2_2))}, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal) != (0)))
		}
	end_branch_1:
		return gopurs_runtime.Bool(__t1)
	})}))}
}

func Call_Main_traverseStr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply2(dictTraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())))
}

func Call_Main_traverseStr__163786110(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__163786110:
	for {
		if false {
			continue traverseStr__163786110
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))})).UnsafePtr))
	}
}

func Call_Main_traverseStr__1598195487(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__1598195487:
	for {
		if false {
			continue traverseStr__1598195487
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m1())).UnsafePtr))
	}
}

func Call_Main_traverseStr__2241449532(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__2241449532:
	for {
		if false {
			continue traverseStr__2241449532
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})))})).UnsafePtr))
	}
}

func Call_Main_traverseStr__2823203293(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__2823203293:
	for {
		if false {
			continue traverseStr__2823203293
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m3())).UnsafePtr))
	}
}

func Call_Main_traverseStr__2384857594(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__2384857594:
	for {
		if false {
			continue traverseStr__2384857594
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
			} {
				orig := Get_Main_recordValue()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
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
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.a), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
				arr := orig.zArrayA
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})
		}()})))})).UnsafePtr))
	}
}

func Call_Main_traverseStr__3819266971(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__3819266971:
	for {
		if false {
			continue traverseStr__3819266971
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m5())).UnsafePtr))
	}
}

func Call_Main_traverseStr__167553720(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__167553720:
	for {
		if false {
			continue traverseStr__167553720
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m6())).UnsafePtr))
	}
}

func Call_Main_traverseStr__749307481(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
traverseStr__749307481:
	for {
		if false {
			continue traverseStr__749307481
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray())), Get_Main_m7())).UnsafePtr))
	}
}

func Call_Main_sequenceStr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply(Call_Data_Traversable_sequence(dictTraversable_0), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))})
}

func Call_Main_sequenceStr__20379839(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__20379839:
	for {
		if false {
			continue sequenceStr__20379839
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3974551048_67812977(nil))})).UnsafePtr))
	}
}

func Call_Main_sequenceStr__2491708894(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__2491708894:
	for {
		if false {
			continue sequenceStr__2491708894
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m1_prime_())).UnsafePtr))
	}
}

func Call_Main_sequenceStr__2672135165(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__2672135165:
	for {
		if false {
			continue sequenceStr__2672135165
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_2857385098_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, []string]{1, int64(0)})))})).UnsafePtr))
	}
}

func Call_Main_sequenceStr__1137381404(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__1137381404:
	for {
		if false {
			continue sequenceStr__1137381404
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m3_prime_())).UnsafePtr))
	}
}

func Call_Main_sequenceStr__2102650939(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__2102650939:
	for {
		if false {
			continue sequenceStr__2102650939
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_2381891596_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, []string]{1, func() gopurs_runtime.Value {
			orig := func() struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
			} {
				orig := Get_Main_recordValue_prime_()
				_ = orig
				clone := struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
				}{}
				clone.a = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				clone.arrayIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "arrayIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fIgnore = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fIgnore").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				clone.fa = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "fa").UnsafePtr)
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
				clone.ignore = gopurs_runtime.RecordGet(orig, "ignore").IntVal
				clone.zArrayA = func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "zArrayA").UnsafePtr)
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
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
				arr := orig.a
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := orig.arrayIgnore
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
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
			}(), gopurs_runtime.Int(orig.ignore), func() gopurs_runtime.Value {
				arr := orig.zArrayA
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
			}()})
		}()})))})).UnsafePtr))
	}
}

func Call_Main_sequenceStr__279012698(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__279012698:
	for {
		if false {
			continue sequenceStr__279012698
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m5_prime_())).UnsafePtr))
	}
}

func Call_Main_sequenceStr__459438969(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__459438969:
	for {
		if false {
			continue sequenceStr__459438969
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m6_prime_())).UnsafePtr))
	}
}

func Call_Main_sequenceStr__3219652504(__eta_norm_0_unused_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
sequenceStr__3219652504:
	for {
		if false {
			continue sequenceStr__3219652504
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m7_prime_())).UnsafePtr))
	}
}

func Rebox_Main_1019686859_2554376626(in *Constructor_Main_M3[gopurs_runtime.Value, []string]) *Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
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

func Rebox_Main_1264725646_4224211191(in *Constructor_Main_M6[gopurs_runtime.Value, []string]) *Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = func() gopurs_runtime.Value {
		arr := in.V1
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Str(v)
		}
		return gopurs_runtime.Array(boxed)
	}()
	out.V2 = in.V2
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = in.V5
	out.V6 = in.V6
	out.V7 = in.V7
	return out
}

func Rebox_Main_1302345387_2067946548(in *Constructor_Main_M5[gopurs_runtime.Value, string]) *Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1554627816_4224211191(in *Constructor_Main_M6[gopurs_runtime.Value, string]) *Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{}
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

func Rebox_Main_1939691112_3790796878(in *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1951009097_3660606000(in *Constructor_Main_M1[gopurs_runtime.Value, []string]) *Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		arr := in.V0
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Str(v)
		}
		return gopurs_runtime.Array(boxed)
	}()
	out.V1 = in.V1
	return out
}

func Rebox_Main_2381891596_2770120821(in *Constructor_Main_M4[gopurs_runtime.Value, []string]) *Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{}
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

func Rebox_Main_2857385098_1521903347(in *Constructor_Main_M2[gopurs_runtime.Value, []string]) *Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_358349645_2067946548(in *Constructor_Main_M5[gopurs_runtime.Value, []string]) *Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3721994703_961717174(in *Constructor_Main_M7[gopurs_runtime.Value, []string]) *Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3741832558_67812977(in *Constructor_Main_M0[gopurs_runtime.Value, string]) *Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{}

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

func Rebox_Main_3974551048_67812977(in *Constructor_Main_M0[gopurs_runtime.Value, []string]) *Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{}

	return out
}

func Rebox_Main_4004653231_3660606000(in *Constructor_Main_M1[gopurs_runtime.Value, string]) *Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
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

func Rebox_Main_82800937_961717174(in *Constructor_Main_M7[gopurs_runtime.Value, string]) *Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
