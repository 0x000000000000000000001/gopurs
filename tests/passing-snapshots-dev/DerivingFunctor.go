package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
				return gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, (*(*[]gopurs_runtime.Value)((value1).UnsafePtr))}))}
			})
		})
	})
	return cache_Main_M0
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
				/*gopurs:record-v1:1:006e00650073007400650064*/
			} {
				orig := value0
				_ = orig
				clone := struct {
					nested gopurs_runtime.Value
					/*gopurs:record-v1:1:006e00650073007400650064*/
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
										}(), (*(*[]gopurs_runtime.Value)((value3).UnsafePtr)), value4, value5, value6, func() struct {
											nested gopurs_runtime.Value
											/*gopurs:record-v1:1:006e00650073007400650064*/
										} {
											orig := value7
											_ = orig
											clone := struct {
												nested gopurs_runtime.Value
												/*gopurs:record-v1:1:006e00650073007400650064*/
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
							unboxed[i] = (*(*[]gopurs_runtime.Value)((v).UnsafePtr))
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
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
			__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()).V0, gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0)))
			_ = __local_var_2_0
			return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(m_1, x_3))
			})
		})}))}
	})
	return cache_Main_functorFun2
}

var cache_Main_functorFun1 gopurs_runtime.Value
var once_Main_functorFun1 sync.Once

func Get_Main_functorFun1() gopurs_runtime.Value {
	once_Main_functorFun1.Do(func() {
		cache_Main_functorFun1 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=Any
				__local_var_3_0 := gopurs_runtime.Apply(m_1, x_2)
				_ = __local_var_3_0
				return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_0, x_4))
				})
			})
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

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			}{"71", []int64{int64(92), int64(93)}, struct {

				/*gopurs:record-v1:0:*/
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

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			}{int64(71), []int64{int64(92), int64(93)}, struct {

				/*gopurs:record-v1:0:*/
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
						unboxed[i] = (*(*[]gopurs_runtime.Value)((v).UnsafePtr))
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
						unboxed[i] = (*(*[]gopurs_runtime.Value)((v).UnsafePtr))
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
		cache_Main_m5R = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_1302345387_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, string]{1, int64(0), "1", []int64{int64(2), int64(3)}, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := []string{"3", "4"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr)), func() gopurs_runtime.Value {
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

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
		}(), struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				empty       struct {

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
		}()}})))}
	})
	return cache_Main_m5R
}

var cache_Main_m5L gopurs_runtime.Value
var once_Main_m5L sync.Once

func Get_Main_m5L() gopurs_runtime.Value {
	once_Main_m5L.Do(func() {
		cache_Main_m5L = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_2850066319_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, int64]{1, int64(0), int64(1), []int64{int64(2), int64(3)}, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := []int64{int64(3), int64(4)}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr)), func() gopurs_runtime.Value {
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

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueL()
				_ = orig
				clone := struct {
					a           int64
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
		}(), struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           int64
				arrayIgnore []int64
				empty       struct {

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueL()
				_ = orig
				clone := struct {
					a           int64
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
		}()}})))}
	})
	return cache_Main_m5L
}

var cache_Main_m4R gopurs_runtime.Value
var once_Main_m4R sync.Once

func Get_Main_m4R() gopurs_runtime.Value {
	once_Main_m4R.Do(func() {
		cache_Main_m4R = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				empty       struct {

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
		}()}})))}
	})
	return cache_Main_m4R
}

var cache_Main_m4L gopurs_runtime.Value
var once_Main_m4L sync.Once

func Get_Main_m4L() gopurs_runtime.Value {
	once_Main_m4L.Do(func() {
		cache_Main_m4L = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_2124045134_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, int64]{1, struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           int64
				arrayIgnore []int64
				empty       struct {

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueL()
				_ = orig
				clone := struct {
					a           int64
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
		}()}})))}
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

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []int64
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
				zArrayA    []int64
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueL()
				_ = orig
				clone := struct {
					a           int64
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
		cache_Main_m0R = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977((&Constructor_Main_M0[gopurs_runtime.Value, string]{1, "0", (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := []string{"1", "2"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))})))}
	})
	return cache_Main_m0R
}

var cache_Main_m0L gopurs_runtime.Value
var once_Main_m0L sync.Once

func Get_Main_m0L() gopurs_runtime.Value {
	once_Main_m0L.Do(func() {
		cache_Main_m0L = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3716704586_67812977((&Constructor_Main_M0[gopurs_runtime.Value, int64]{1, int64(0), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := []int64{int64(1), int64(2)}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))})))}
	})
	return cache_Main_m0L
}

var cache_Main_functorT gopurs_runtime.Value
var once_Main_functorT sync.Once

func Get_Main_functorT() gopurs_runtime.Value {
	once_Main_functorT.Do(func() {
		cache_Main_functorT = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(dictShow_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=Any
				__local_var_3_0 := gopurs_runtime.Apply(m_1, dictShow_2)
				_ = __local_var_3_0
				return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_0, x_4))
				})
			})
		})}))}
	})
	return cache_Main_functorT
}

var cache_Main_taTests gopurs_runtime.Value
var once_Main_taTests sync.Once

func Get_Main_taTests() gopurs_runtime.Value {
	once_Main_taTests.Do(func() {
		cache_Main_taTests = gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map show T"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Func2(func(dictShow_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Data_Show_ShowIntImpl(gopurs_runtime.Int(int64(42)).IntVal))
		}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}, gopurs_runtime.Str("hello")).StrVal()) == ("42")))
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

var cache_Main_f3Test gopurs_runtime.Value
var once_Main_f3Test sync.Once

func Get_Main_f3Test() gopurs_runtime.Value {
	once_Main_f3Test.Do(func() {
		cache_Main_f3Test = gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - Fun3"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			} {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_0())), Get_Data_Show_showIntImpl(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr := [][][]struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {

									/*gopurs:record-v1:0:*/
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
								/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
							}
							/*gopurs:record-v1:1:006e00650073007400650064*/
						}{[][]struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {

									/*gopurs:record-v1:0:*/
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
								/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
							}
							/*gopurs:record-v1:1:006e00650073007400650064*/
						}{[]struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {

									/*gopurs:record-v1:0:*/
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
								/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
							}
							/*gopurs:record-v1:1:006e00650073007400650064*/
						}{struct {
							nested struct {
								a           int64
								arrayIgnore []int64
								empty       struct {

									/*gopurs:record-v1:0:*/
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
								/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
							}
							/*gopurs:record-v1:1:006e00650073007400650064*/
						}{func() struct {
							a           int64
							arrayIgnore []int64
							empty       struct {

								/*gopurs:record-v1:0:*/
							}
							fIgnore    []int64
							fa         []int64
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
							zArrayA    []int64
							/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
						} {
							orig := Get_Main_recordValueL()
							_ = orig
							clone := struct {
								a           int64
								arrayIgnore []int64
								empty       struct {

									/*gopurs:record-v1:0:*/
								}
								fIgnore    []int64
								fa         []int64
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
								zArrayA    []int64
								/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

								/*gopurs:record-v1:0:*/
							} {
								orig := gopurs_runtime.RecordGet(orig, "empty")
								_ = orig
								clone := struct {

									/*gopurs:record-v1:0:*/
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

							/*gopurs:record-v1:0:*/
						}
						fIgnore    []int64
						fa         []string
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
						zArrayA    []string
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
				}, len(arr))
				for i, v := range arr {
					unboxed[i] = func() [][]struct {
						nested struct {
							a           string
							arrayIgnore []int64
							empty       struct {

								/*gopurs:record-v1:0:*/
							}
							fIgnore    []int64
							fa         []string
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
							zArrayA    []string
							/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
						}
						/*gopurs:record-v1:1:006e00650073007400650064*/
					} {
						arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
						unboxed := make([][]struct {
							nested struct {
								a           string
								arrayIgnore []int64
								empty       struct {

									/*gopurs:record-v1:0:*/
								}
								fIgnore    []int64
								fa         []string
								ignore     int64
								recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
								zArrayA    []string
								/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
							}
							/*gopurs:record-v1:1:006e00650073007400650064*/
						}, len(arr))
						for i, v := range arr {
							unboxed[i] = func() []struct {
								nested struct {
									a           string
									arrayIgnore []int64
									empty       struct {

										/*gopurs:record-v1:0:*/
									}
									fIgnore    []int64
									fa         []string
									ignore     int64
									recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
									zArrayA    []string
									/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
								}
								/*gopurs:record-v1:1:006e00650073007400650064*/
							} {
								arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
								unboxed := make([]struct {
									nested struct {
										a           string
										arrayIgnore []int64
										empty       struct {

											/*gopurs:record-v1:0:*/
										}
										fIgnore    []int64
										fa         []string
										ignore     int64
										recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
										zArrayA    []string
										/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
									}
									/*gopurs:record-v1:1:006e00650073007400650064*/
								}, len(arr))
								for i, v := range arr {
									unboxed[i] = func() struct {
										nested struct {
											a           string
											arrayIgnore []int64
											empty       struct {

												/*gopurs:record-v1:0:*/
											}
											fIgnore    []int64
											fa         []string
											ignore     int64
											recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
											zArrayA    []string
											/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
										}
										/*gopurs:record-v1:1:006e00650073007400650064*/
									} {
										orig := v
										_ = orig
										clone := struct {
											nested struct {
												a           string
												arrayIgnore []int64
												empty       struct {

													/*gopurs:record-v1:0:*/
												}
												fIgnore    []int64
												fa         []string
												ignore     int64
												recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
												zArrayA    []string
												/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
											}
											/*gopurs:record-v1:1:006e00650073007400650064*/
										}{}
										clone.nested = func() struct {
											a           string
											arrayIgnore []int64
											empty       struct {

												/*gopurs:record-v1:0:*/
											}
											fIgnore    []int64
											fa         []string
											ignore     int64
											recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
											zArrayA    []string
											/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
										} {
											orig := gopurs_runtime.RecordGet(orig, "nested")
											_ = orig
											clone := struct {
												a           string
												arrayIgnore []int64
												empty       struct {

													/*gopurs:record-v1:0:*/
												}
												fIgnore    []int64
												fa         []string
												ignore     int64
												recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
												zArrayA    []string
												/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

												/*gopurs:record-v1:0:*/
											} {
												orig := gopurs_runtime.RecordGet(orig, "empty")
												_ = orig
												clone := struct {

													/*gopurs:record-v1:0:*/
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

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{[][]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{[]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{struct {
				nested struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{func() struct {
				a           string
				arrayIgnore []int64
				empty       struct {

					/*gopurs:record-v1:0:*/
				}
				fIgnore    []int64
				fa         []string
				ignore     int64
				recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
				zArrayA    []string
				/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
			} {
				orig := Get_Main_recordValueR()
				_ = orig
				clone := struct {
					a           string
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

					/*gopurs:record-v1:0:*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "empty")
					_ = orig
					clone := struct {

						/*gopurs:record-v1:0:*/
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
			arrayIntrinsic_arg_0 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Show_showIntImpl())
			_ = arrayIntrinsic_arg_0
			arrayIntrinsic_arg_1 := func() gopurs_runtime.Value {
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
			_ = arrayIntrinsic_arg_1
			arrayIntrinsic_arg_2 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Show_showIntImpl())
			_ = arrayIntrinsic_arg_2
			arrayIntrinsic_arg_3 := func() gopurs_runtime.Value {
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
			_ = arrayIntrinsic_arg_3
			return gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - Fun2"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), func() gopurs_runtime.Value {
				arr := func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap2 := arrayIntrinsic_arg_1
						_ = arr_val_arrayMap2
						arr_go_arrayMap2 := (*[]gopurs_runtime.Value)(arr_val_arrayMap2.UnsafePtr)
						_ = arr_go_arrayMap2
						res_go_arrayMap2 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap2))
						_ = res_go_arrayMap2
						for i_arrayMap2, v_arrayMap2 := range *arr_go_arrayMap2 {
							res_go_arrayMap2[i_arrayMap2] = gopurs_runtime.Apply(arrayIntrinsic_arg_0, v_arrayMap2)
						}
						return gopurs_runtime.Array(res_go_arrayMap2)
					}().UnsafePtr)
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
			}(), func() gopurs_runtime.Value {
				arr := func() [][]string {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap2 := arrayIntrinsic_arg_3
						_ = arr_val_arrayMap2
						arr_go_arrayMap2 := (*[]gopurs_runtime.Value)(arr_val_arrayMap2.UnsafePtr)
						_ = arr_go_arrayMap2
						res_go_arrayMap2 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap2))
						_ = res_go_arrayMap2
						for i_arrayMap2, v_arrayMap2 := range *arr_go_arrayMap2 {
							res_go_arrayMap2[i_arrayMap2] = gopurs_runtime.Apply(arrayIntrinsic_arg_2, v_arrayMap2)
						}
						return gopurs_runtime.Array(res_go_arrayMap2)
					}().UnsafePtr)
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
		cache_Main_f1Test = gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - Fun1"), gopurs_runtime.Bool((Data_Show_ShowIntImpl(gopurs_runtime.Int(int64(3)).IntVal)) == (Data_Show_ShowIntImpl(int64(3)))))
	})
	return cache_Main_f1Test
}

var cache_Main_funTests gopurs_runtime.Value
var once_Main_funTests sync.Once

func Get_Main_funTests() gopurs_runtime.Value {
	once_Main_funTests.Do(func() {
		cache_Main_funTests = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			_dollar___unused_0_0 := gopurs_runtime.Apply(Get_Main_f1Test(), gopurs_runtime.Value{})
			_ = _dollar___unused_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(Get_Main_f2Test(), gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Get_Main_f3Test(), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			return gopurs_runtime.Apply(Get_Main_taTests(), gopurs_runtime.Value{})
		})
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

var cache_Main_maTests gopurs_runtime.Value
var once_Main_maTests sync.Once

func Get_Main_maTests() gopurs_runtime.Value {
	once_Main_maTests.Do(func() {
		cache_Main_maTests = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - M0"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_1())), Get_Data_Show_showIntImpl(), Get_Main_m0L()), Get_Main_m0R()).IntVal) != (0)))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - M1"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_2())), Get_Data_Show_showIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_147758475_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, int64]{1, int64(0)})))}), gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, int64(0)})))}).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - M2"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_3())), Get_Data_Show_showIntImpl(), Get_Main_m2L()), Get_Main_m2R()).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - M3"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_4())), Get_Data_Show_showIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer(Rebox_Main_837768713_2554376626((&Constructor_Main_M3[gopurs_runtime.Value, int64]{1, func() gopurs_runtime.Value {
				orig := func() struct {
					a           int64
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []int64
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
					zArrayA    []int64
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				} {
					orig := Get_Main_recordValueL()
					_ = orig
					clone := struct {
						a           int64
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    []int64
						fa         []int64
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []int64]
						zArrayA    []int64
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

						/*gopurs:record-v1:0:*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "empty")
						_ = orig
						clone := struct {

							/*gopurs:record-v1:0:*/
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

						/*gopurs:record-v1:0:*/
					}
					fIgnore    []int64
					fa         []string
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
					zArrayA    []string
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				} {
					orig := Get_Main_recordValueR()
					_ = orig
					clone := struct {
						a           string
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    []int64
						fa         []string
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []string]
						zArrayA    []string
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

						/*gopurs:record-v1:0:*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "empty")
						_ = orig
						clone := struct {

							/*gopurs:record-v1:0:*/
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
			}()})))}).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - M4"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_5())), Get_Data_Show_showIntImpl(), Get_Main_m4L()), Get_Main_m4R()).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - M5"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_6())), Get_Data_Show_showIntImpl(), Get_Main_m5L()), Get_Main_m5R()).IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("map - M6"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main___cached_dict_7())), Get_Data_Show_showIntImpl(), Get_Main_m6L()), Get_Main_m6R()).IntVal) != (0))), gopurs_runtime.Value{})
		})
	})
	return cache_Main_maTests
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			_dollar___unused_0_0 := gopurs_runtime.Apply(Get_Main_maTests(), gopurs_runtime.Value{})
			_ = _dollar___unused_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(Get_Main_funTests(), gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorFun3(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Main_functorM(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})))}
	})
	return cache_Main___cached_dict_7
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
		/*gopurs:record-v1:1:006e00650073007400650064*/
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
		/*gopurs:record-v1:1:006e00650073007400650064*/
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

func Call_Main_functorFun3(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(Func [(Array (TypeVar a$scope5))] (Array (TypeVar b$scope6)))
			__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
			_ = __local_var_4_1
			arrayIntrinsic_arg_2 := gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_4_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1)}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			})
			_ = arrayIntrinsic_arg_2
			arrayIntrinsic_arg_3 := func() gopurs_runtime.Value {
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
			_ = arrayIntrinsic_arg_3
			arrayIntrinsic_arg_4 := f_1
			_ = arrayIntrinsic_arg_4
			arrayIntrinsic_arg_5 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA")).UnsafePtr)))
			_ = arrayIntrinsic_arg_5
			return gopurs_runtime.RecordUpdate1(v1_3, "nested", func() gopurs_runtime.Value {
				orig := func() struct {
					a           gopurs_runtime.Value
					arrayIgnore []int64
					empty       struct {

						/*gopurs:record-v1:0:*/
					}
					fIgnore    gopurs_runtime.Value
					fa         gopurs_runtime.Value
					ignore     int64
					recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
					zArrayA    []gopurs_runtime.Value
					/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
				} {
					orig := gopurs_runtime.RecordUpdateDict(gopurs_runtime.RecordGet(v1_3, "nested"), []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a")), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa")), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
								arr_val_arrayMap10 := arrayIntrinsic_arg_3
								_ = arr_val_arrayMap10
								arr_go_arrayMap10 := (*[]gopurs_runtime.Value)(arr_val_arrayMap10.UnsafePtr)
								_ = arr_go_arrayMap10
								res_go_arrayMap10 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap10))
								_ = res_go_arrayMap10
								for i_arrayMap10, v_arrayMap10 := range *arr_go_arrayMap10 {
									res_go_arrayMap10[i_arrayMap10] = gopurs_runtime.Apply(arrayIntrinsic_arg_2, v_arrayMap10)
								}
								return gopurs_runtime.Array(res_go_arrayMap10)
							}()).UnsafePtr))).UnsafePtr)
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
					}(), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
						arr_val_arrayMap10 := arrayIntrinsic_arg_5
						_ = arr_val_arrayMap10
						arr_go_arrayMap10 := (*[]gopurs_runtime.Value)(arr_val_arrayMap10.UnsafePtr)
						_ = arr_go_arrayMap10
						res_go_arrayMap10 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap10))
						_ = res_go_arrayMap10
						for i_arrayMap10, v_arrayMap10 := range *arr_go_arrayMap10 {
							res_go_arrayMap10[i_arrayMap10] = gopurs_runtime.Apply(arrayIntrinsic_arg_4, v_arrayMap10)
						}
						return gopurs_runtime.Array(res_go_arrayMap10)
					}()).UnsafePtr)))})
					_ = orig
					clone := struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

						/*gopurs:record-v1:0:*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "empty")
						_ = orig
						clone := struct {

							/*gopurs:record-v1:0:*/
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
					clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
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
		}))))
		_ = __local_var_3_0
		return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Apply(m_2, x_4))
		})
	})}))}
}

func Call_Main_functorM(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t26 gopurs_runtime.Value
		{
			if m_2.Type == 9 && m_2.IntVal == 3852365315 {
				arrayIntrinsic_arg_0 := f_1
				_ = arrayIntrinsic_arg_0
				arrayIntrinsic_arg_1 := gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)
				_ = arrayIntrinsic_arg_1
				__t26 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := arrayIntrinsic_arg_1
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(arrayIntrinsic_arg_0, v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr))}))}
				goto end_branch_26
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 769986722 {
				__t26 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0}))}
				goto end_branch_26
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2727978561 {
				__t26 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
				goto end_branch_26
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 1830062304 {
				// TAST (Let): __local_var_3_2 shape=App(Var) bindingType=(Func [(Array (TypeVar a$scope75))] (Array (TypeVar b$scope76)))
				__local_var_3_2 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_2
				arrayIntrinsic_arg_3 := gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				})
				_ = arrayIntrinsic_arg_3
				arrayIntrinsic_arg_4 := func() gopurs_runtime.Value {
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
				_ = arrayIntrinsic_arg_4
				arrayIntrinsic_arg_5 := f_1
				_ = arrayIntrinsic_arg_5
				arrayIntrinsic_arg_6 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_6
				__t26 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "fa")), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap5 := arrayIntrinsic_arg_4
							_ = arr_val_arrayMap5
							arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
							_ = arr_go_arrayMap5
							res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
							_ = res_go_arrayMap5
							for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
								res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(arrayIntrinsic_arg_3, v_arrayMap5)
							}
							return gopurs_runtime.Array(res_go_arrayMap5)
						}()).UnsafePtr))).UnsafePtr)
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
				}(), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap5 := arrayIntrinsic_arg_6
					_ = arr_val_arrayMap5
					arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
					_ = arr_go_arrayMap5
					res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
					_ = res_go_arrayMap5
					for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
						res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(arrayIntrinsic_arg_5, v_arrayMap5)
					}
					return gopurs_runtime.Array(res_go_arrayMap5)
				}()).UnsafePtr)))})}))}
				goto end_branch_26
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 3190619783 {
				// TAST (Let): __local_var_3_7 shape=App(Var) bindingType=(Func [(Array (TypeVar a$scope75))] (Array (TypeVar b$scope76)))
				__local_var_3_7 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_7
				arrayIntrinsic_arg_8 := gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_7, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				})
				_ = arrayIntrinsic_arg_8
				arrayIntrinsic_arg_9 := func() gopurs_runtime.Value {
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
				_ = arrayIntrinsic_arg_9
				arrayIntrinsic_arg_10 := f_1
				_ = arrayIntrinsic_arg_10
				arrayIntrinsic_arg_11 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_11
				__t26 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, func(record struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
				}) struct {
					nested gopurs_runtime.Value
					/*gopurs:record-v1:1:006e00650073007400650064*/
				} {
					return struct {
						nested gopurs_runtime.Value
						/*gopurs:record-v1:1:006e00650073007400650064*/
					}{func() gopurs_runtime.Value {
						orig := record.nested
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
					}()}
				}(func() struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
				} {
					originalRecord := (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0
					_ = originalRecord
					_ = originalRecord
					var clone struct {
						nested struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {

								/*gopurs:record-v1:0:*/
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
							/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
						}
						/*gopurs:record-v1:1:006e00650073007400650064*/
					}
					clone.nested = func() struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
					} {
						orig := gopurs_runtime.RecordUpdateDict((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "fa")), func() gopurs_runtime.Value {
							arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
									arr_val_arrayMap6 := arrayIntrinsic_arg_9
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(arrayIntrinsic_arg_8, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}()).UnsafePtr))).UnsafePtr)
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
						}(), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap6 := arrayIntrinsic_arg_11
							_ = arr_val_arrayMap6
							arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
							_ = arr_go_arrayMap6
							res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
							_ = res_go_arrayMap6
							for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
								res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(arrayIntrinsic_arg_10, v_arrayMap6)
							}
							return gopurs_runtime.Array(res_go_arrayMap6)
						}()).UnsafePtr)))})
						_ = orig
						clone := struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {

								/*gopurs:record-v1:0:*/
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
							/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

							/*gopurs:record-v1:0:*/
						} {
							orig := gopurs_runtime.RecordGet(orig, "empty")
							_ = orig
							clone := struct {

								/*gopurs:record-v1:0:*/
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
						clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
						return clone
					}()
					return clone
				}())}))}
				goto end_branch_26
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 108241190 {
				arrayIntrinsic_arg_12 := f_1
				_ = arrayIntrinsic_arg_12
				arrayIntrinsic_arg_13 := gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3)
				_ = arrayIntrinsic_arg_13
				// TAST (Let): __local_var_3_14 shape=App(Var) bindingType=(Func [(Array (TypeVar a$scope75))] (Array (TypeVar b$scope76)))
				__local_var_3_14 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_14
				arrayIntrinsic_arg_15 := gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_14, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				})
				_ = arrayIntrinsic_arg_15
				arrayIntrinsic_arg_16 := func() gopurs_runtime.Value {
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
				_ = arrayIntrinsic_arg_16
				arrayIntrinsic_arg_17 := f_1
				_ = arrayIntrinsic_arg_17
				arrayIntrinsic_arg_18 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_18
				// TAST (Let): __local_var_3_19 shape=App(Var) bindingType=(Func [(Array (TypeVar a$scope75))] (Array (TypeVar b$scope76)))
				__local_var_3_19 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_3_19
				arrayIntrinsic_arg_20 := gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_19, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				})
				_ = arrayIntrinsic_arg_20
				arrayIntrinsic_arg_21 := func() gopurs_runtime.Value {
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
				_ = arrayIntrinsic_arg_21
				arrayIntrinsic_arg_22 := f_1
				_ = arrayIntrinsic_arg_22
				arrayIntrinsic_arg_23 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_23
				__t26 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := arrayIntrinsic_arg_13
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(arrayIntrinsic_arg_12, v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr)), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "fa")), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap5 := arrayIntrinsic_arg_16
							_ = arr_val_arrayMap5
							arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
							_ = arr_go_arrayMap5
							res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
							_ = res_go_arrayMap5
							for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
								res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(arrayIntrinsic_arg_15, v_arrayMap5)
							}
							return gopurs_runtime.Array(res_go_arrayMap5)
						}()).UnsafePtr))).UnsafePtr)
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
				}(), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap5 := arrayIntrinsic_arg_18
					_ = arr_val_arrayMap5
					arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
					_ = arr_go_arrayMap5
					res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
					_ = res_go_arrayMap5
					for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
						res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(arrayIntrinsic_arg_17, v_arrayMap5)
					}
					return gopurs_runtime.Array(res_go_arrayMap5)
				}()).UnsafePtr)))}), func(record struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
				}) struct {
					nested gopurs_runtime.Value
					/*gopurs:record-v1:1:006e00650073007400650064*/
				} {
					return struct {
						nested gopurs_runtime.Value
						/*gopurs:record-v1:1:006e00650073007400650064*/
					}{func() gopurs_runtime.Value {
						orig := record.nested
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
					}()}
				}(func() struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
				} {
					originalRecord := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7
					_ = originalRecord
					_ = originalRecord
					var clone struct {
						nested struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {

								/*gopurs:record-v1:0:*/
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
							/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
						}
						/*gopurs:record-v1:1:006e00650073007400650064*/
					}
					clone.nested = func() struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						empty       struct {

							/*gopurs:record-v1:0:*/
						}
						fIgnore    gopurs_runtime.Value
						fa         gopurs_runtime.Value
						ignore     int64
						recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
						zArrayA    []gopurs_runtime.Value
						/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
					} {
						orig := gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "fa")), func() gopurs_runtime.Value {
							arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
									arr_val_arrayMap6 := arrayIntrinsic_arg_21
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(arrayIntrinsic_arg_20, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}()).UnsafePtr))).UnsafePtr)
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
						}(), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap6 := arrayIntrinsic_arg_23
							_ = arr_val_arrayMap6
							arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
							_ = arr_go_arrayMap6
							res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
							_ = res_go_arrayMap6
							for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
								res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(arrayIntrinsic_arg_22, v_arrayMap6)
							}
							return gopurs_runtime.Array(res_go_arrayMap6)
						}()).UnsafePtr)))})
						_ = orig
						clone := struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							empty       struct {

								/*gopurs:record-v1:0:*/
							}
							fIgnore    gopurs_runtime.Value
							fa         gopurs_runtime.Value
							ignore     int64
							recursiveA []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]
							zArrayA    []gopurs_runtime.Value
							/*gopurs:record-v1:8:0061,0061007200720061007900490067006e006f00720065,0065006d007000740079,006600490067006e006f00720065,00660061,00690067006e006f00720065,0072006500630075007200730069007600650041,007a004100720072006100790041*/
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

							/*gopurs:record-v1:0:*/
						} {
							orig := gopurs_runtime.RecordGet(orig, "empty")
							_ = orig
							clone := struct {

								/*gopurs:record-v1:0:*/
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
						clone.zArrayA = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "zArrayA")).UnsafePtr))
						return clone
					}()
					return clone
				}())}))}
				goto end_branch_26
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2066233029 {
				arrayIntrinsic_arg_24 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1))
				_ = arrayIntrinsic_arg_24
				arrayIntrinsic_arg_25 := func() gopurs_runtime.Value {
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
				_ = arrayIntrinsic_arg_25
				__t26 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() [][][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap4 := arrayIntrinsic_arg_25
						_ = arr_val_arrayMap4
						arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
						_ = arr_go_arrayMap4
						res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
						_ = res_go_arrayMap4
						for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
							res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(arrayIntrinsic_arg_24, v_arrayMap4)
						}
						return gopurs_runtime.Array(res_go_arrayMap4)
					}().UnsafePtr)
					unboxed := make([][][]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = func() [][]gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
							unboxed := make([][]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = (*(*[]gopurs_runtime.Value)((v).UnsafePtr))
							}
							return unboxed
						}()
					}
					return unboxed
				}()}))}
				goto end_branch_26
			} else {

			}
		}
		{
			__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_26:
		return __t26
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
				__t4 = (y_7.Type == 9 && y_7.IntVal == 2727978561) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0).IntVal) != (0))
				goto end_branch_4
			} else {

			}
		}
		{
			if x_6.Type == 9 && x_6.IntVal == 1830062304 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 1830062304) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
				}(), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "zArrayA")).UnsafePtr)))).IntVal) != (0)))
				goto end_branch_4
			} else {

			}
		}
		{
			if x_6.Type == 9 && x_6.IntVal == 3190619783 {
				__t4 = (y_7.Type == 9 && y_7.IntVal == 3190619783) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
				}(), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))).IntVal) != (0)))
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
				}(), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "zArrayA")).UnsafePtr)))).IntVal) != (0)))) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
				}(), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_2.V0, func() gopurs_runtime.Value {
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray3_2_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))).IntVal) != (0))))
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
	return (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
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
	out.V1 = (*(*[]gopurs_runtime.Value)((in.V1).UnsafePtr))
	return out
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_147758475_3660606000(in *Constructor_Main_M1[gopurs_runtime.Value, int64]) *Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1554627816_4224211191(in *Constructor_Main_M6[gopurs_runtime.Value, string]) *Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2124045134_2770120821(in *Constructor_Main_M4[gopurs_runtime.Value, int64]) *Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2628211660_4224211191(in *Constructor_Main_M6[gopurs_runtime.Value, int64]) *Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2785108781_2554376626(in *Constructor_Main_M3[gopurs_runtime.Value, string]) *Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
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

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1161550504(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1871481144(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1939691112(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4004653231_3660606000(in *Constructor_Main_M1[gopurs_runtime.Value, string]) *Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4220871112_1521903347(in *Constructor_Main_M2[gopurs_runtime.Value, int64]) *Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4256935660_1521903347(in *Constructor_Main_M2[gopurs_runtime.Value, string]) *Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
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
	return (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
