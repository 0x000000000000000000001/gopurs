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
		cache_Main_show = Get_Data_Show_showIntImpl()
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
		cache_Main_eqRowCons = gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Eq_eqRowNil()))}, gopurs_runtime.Value{})
	})
	return cache_Main_eqRowCons
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
	})
	return cache_Main_eqArray
}

var cache_Main_eqTuple gopurs_runtime.Value
var once_Main_eqTuple sync.Once

func Get_Main_eqTuple() gopurs_runtime.Value {
	once_Main_eqTuple.Do(func() {
		cache_Main_eqTuple = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqTuple(dictEq1_0_box)
		})
	})
	return cache_Main_eqTuple
}

var cache_Main_eqArray1 gopurs_runtime.Value
var once_Main_eqArray1 sync.Once

func Get_Main_eqArray1() gopurs_runtime.Value {
	once_Main_eqArray1.Do(func() {
		cache_Main_eqArray1 = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl()))
	})
	return cache_Main_eqArray1
}

var cache_Main_eqArray2 gopurs_runtime.Value
var once_Main_eqArray2 sync.Once

func Get_Main_eqArray2() gopurs_runtime.Value {
	once_Main_eqArray2.Do(func() {
		cache_Main_eqArray2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[[]int64]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl())}))}
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
			return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() *struct {
				nested gopurs_runtime.Value
			} {
				orig := value0
				_ = orig
				clone := struct {
					nested gopurs_runtime.Value
				}{}
				clone.nested = gopurs_runtime.RecordGet(orig, "nested")
				return &clone
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
										}(), value4, value5, value6, func() *struct {
											nested gopurs_runtime.Value
										} {
											orig := value7
											_ = orig
											clone := struct {
												nested gopurs_runtime.Value
											}{}
											clone.nested = gopurs_runtime.RecordGet(orig, "nested")
											return &clone
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
		cache_Main_functorFun2 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
				__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()).V0), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0)))
				_ = __local_var_2_0
				return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(m_1, x_3))
				})
			})
		})}))}
	})
	return cache_Main_functorFun2
}

var cache_Main_functorFun1 gopurs_runtime.Value
var once_Main_functorFun1 sync.Once

func Get_Main_functorFun1() gopurs_runtime.Value {
	once_Main_functorFun1.Do(func() {
		cache_Main_functorFun1 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=Any
					__local_var_3_0 := gopurs_runtime.Apply(m_1, x_2)
					_ = __local_var_3_0
					return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_0, x_4))
					})
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
		cache_Main_recordValueR = gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str("71"), func() gopurs_runtime.Value {
			arr := []int64{92, 93}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.RecordDict0(), func() gopurs_runtime.Value {
			arr := []int64{94}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []string{"73"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.Int(91), func() gopurs_runtime.Value {
			arr := []*Constructor_Data_Tuple_Tuple[int64, []string]{(&Constructor_Data_Tuple_Tuple[int64, []string]{1, gopurs_runtime.Int(1), func() gopurs_runtime.Value {
				arr := []string{"1"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()}), (&Constructor_Data_Tuple_Tuple[int64, []string]{1, gopurs_runtime.Int(2), func() gopurs_runtime.Value {
				arr := []string{"2"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []string{"72"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()})
	})
	return cache_Main_recordValueR
}

var cache_Main_recordValueL gopurs_runtime.Value
var once_Main_recordValueL sync.Once

func Get_Main_recordValueL() gopurs_runtime.Value {
	once_Main_recordValueL.Do(func() {
		cache_Main_recordValueL = gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "empty", "fIgnore", "fa", "ignore", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Int(71), func() gopurs_runtime.Value {
			arr := []int64{92, 93}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.RecordDict0(), func() gopurs_runtime.Value {
			arr := []int64{94}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{73}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.Int(91), func() gopurs_runtime.Value {
			arr := []*Constructor_Data_Tuple_Tuple[int64, []int64]{(&Constructor_Data_Tuple_Tuple[int64, []int64]{1, gopurs_runtime.Int(1), func() gopurs_runtime.Value {
				arr := []int64{1}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()}), (&Constructor_Data_Tuple_Tuple[int64, []int64]{1, gopurs_runtime.Int(2), func() gopurs_runtime.Value {
				arr := []int64{2}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()})}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{72}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()})
	})
	return cache_Main_recordValueL
}

var cache_Main_m6R gopurs_runtime.Value
var once_Main_m6R sync.Once

func Get_Main_m6R() gopurs_runtime.Value {
	once_Main_m6R.Do(func() {
		cache_Main_m6R = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, string]{1, func() [][][]gopurs_runtime.Value {
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
		}()}))}
	})
	return cache_Main_m6R
}

var cache_Main_m6L gopurs_runtime.Value
var once_Main_m6L sync.Once

func Get_Main_m6L() gopurs_runtime.Value {
	once_Main_m6L.Do(func() {
		cache_Main_m6L = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, int64]{1, func() [][][]gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := [][][]int64{[][]int64{[]int64{1, 2}}}
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
		}()}))}
	})
	return cache_Main_m6L
}

var cache_Main_m5R gopurs_runtime.Value
var once_Main_m5R sync.Once

func Get_Main_m5R() gopurs_runtime.Value {
	once_Main_m5R.Do(func() {
		cache_Main_m5R = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, string]{1, 0, gopurs_runtime.Str("1"), []int64{2, 3}, func() []gopurs_runtime.Value {
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
			arr := []int64{7, 8}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), Get_Main_recordValueR(), func() *struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", Get_Main_recordValueR())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return &clone
		}()}))}
	})
	return cache_Main_m5R
}

var cache_Main_m5L gopurs_runtime.Value
var once_Main_m5L sync.Once

func Get_Main_m5L() gopurs_runtime.Value {
	once_Main_m5L.Do(func() {
		cache_Main_m5L = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, int64]{1, 0, gopurs_runtime.Int(1), []int64{2, 3}, func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []int64{3, 4}
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
			arr := []int64{5, 6}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{7, 8}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), Get_Main_recordValueL(), func() *struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", Get_Main_recordValueL())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return &clone
		}()}))}
	})
	return cache_Main_m5L
}

var cache_Main_m4R gopurs_runtime.Value
var once_Main_m4R sync.Once

func Get_Main_m4R() gopurs_runtime.Value {
	once_Main_m4R.Do(func() {
		cache_Main_m4R = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() *struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", Get_Main_recordValueR())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return &clone
		}()}))}
	})
	return cache_Main_m4R
}

var cache_Main_m4L gopurs_runtime.Value
var once_Main_m4L sync.Once

func Get_Main_m4L() gopurs_runtime.Value {
	once_Main_m4L.Do(func() {
		cache_Main_m4L = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, int64]{1, func() *struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", Get_Main_recordValueL())
			_ = orig
			clone := struct {
				nested gopurs_runtime.Value
			}{}
			clone.nested = gopurs_runtime.RecordGet(orig, "nested")
			return &clone
		}()}))}
	})
	return cache_Main_m4L
}

var cache_Main_m3R gopurs_runtime.Value
var once_Main_m3R sync.Once

func Get_Main_m3R() gopurs_runtime.Value {
	once_Main_m3R.Do(func() {
		cache_Main_m3R = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, string]{1, Get_Main_recordValueR()}))}
	})
	return cache_Main_m3R
}

var cache_Main_m3L gopurs_runtime.Value
var once_Main_m3L sync.Once

func Get_Main_m3L() gopurs_runtime.Value {
	once_Main_m3L.Do(func() {
		cache_Main_m3L = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, int64]{1, Get_Main_recordValueL()}))}
	})
	return cache_Main_m3L
}

var cache_Main_m2R gopurs_runtime.Value
var once_Main_m2R sync.Once

func Get_Main_m2R() gopurs_runtime.Value {
	once_Main_m2R.Do(func() {
		cache_Main_m2R = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			arr := []string{"0", "1"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m2R
}

var cache_Main_m2L gopurs_runtime.Value
var once_Main_m2L sync.Once

func Get_Main_m2L() gopurs_runtime.Value {
	once_Main_m2L.Do(func() {
		cache_Main_m2L = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, int64]{1, func() gopurs_runtime.Value {
			arr := []int64{0, 1}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m2L
}

var cache_Main_m1R gopurs_runtime.Value
var once_Main_m1R sync.Once

func Get_Main_m1R() gopurs_runtime.Value {
	once_Main_m1R.Do(func() {
		cache_Main_m1R = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, 0}))}
	})
	return cache_Main_m1R
}

var cache_Main_m1L gopurs_runtime.Value
var once_Main_m1L sync.Once

func Get_Main_m1L() gopurs_runtime.Value {
	once_Main_m1L.Do(func() {
		cache_Main_m1L = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, int64]{1, 0}))}
	})
	return cache_Main_m1L
}

var cache_Main_m0R gopurs_runtime.Value
var once_Main_m0R sync.Once

func Get_Main_m0R() gopurs_runtime.Value {
	once_Main_m0R.Do(func() {
		cache_Main_m0R = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, string]{1, gopurs_runtime.Str("0"), func() []gopurs_runtime.Value {
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
		}()}))}
	})
	return cache_Main_m0R
}

var cache_Main_m0L gopurs_runtime.Value
var once_Main_m0L sync.Once

func Get_Main_m0L() gopurs_runtime.Value {
	once_Main_m0L.Do(func() {
		cache_Main_m0L = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, int64]{1, gopurs_runtime.Int(0), func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []int64{1, 2}
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
		}()}))}
	})
	return cache_Main_m0L
}

var cache_Main_functorT gopurs_runtime.Value
var once_Main_functorT sync.Once

func Get_Main_functorT() gopurs_runtime.Value {
	once_Main_functorT.Do(func() {
		cache_Main_functorT = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(dictShow_2 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=Any
					__local_var_3_0 := gopurs_runtime.Apply(m_1, dictShow_2)
					_ = __local_var_3_0
					return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_0, x_4))
					})
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
		cache_Main_taTests = gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map show T"), gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(42)).StrVal()) == ("42")))
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
		cache_Main_functorM1 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if m_1.Type == 9 && m_1.IntVal == 3852365315 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_0, v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
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
					if m_1.Type == 9 && m_1.IntVal == 769986722 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0}))}
						goto end_branch_4
					} else {

					}
				}
				{
					if m_1.Type == 9 && m_1.IntVal == 2727978561 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_0, v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())}))}
						goto end_branch_4
					} else {

					}
				}
				{
					if m_1.Type == 9 && m_1.IntVal == 1830062304 {
						// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
						__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0)
						_ = __local_var_2_0
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, "a")), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap6 := gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, "fa")
								_ = arr_val_arrayMap6
								arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
								_ = arr_go_arrayMap6
								res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
								_ = res_go_arrayMap6
								for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()), func() gopurs_runtime.Value {
							arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap6 := func() gopurs_runtime.Value {
											arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, "recursiveA").UnsafePtr)
												unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
												for i, v := range arr {
													unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
												}
												return unboxed
											}()
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
											}
											return gopurs_runtime.Array(boxed)
										}()
										_ = arr_val_arrayMap6
										arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
										_ = arr_go_arrayMap6
										res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
										_ = res_go_arrayMap6
										for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
											res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)}))}
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
									unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
								}
								return unboxed
							}()
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, "zArrayA").UnsafePtr)
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
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
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
					if m_1.Type == 9 && m_1.IntVal == 3190619783 {
						// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
						__local_var_2_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0)
						_ = __local_var_2_1
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() *struct {
							nested gopurs_runtime.Value
						} {
							clone := *((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)
							clone.nested = gopurs_runtime.RecordUpdateDict((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0.nested, "a")), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0.nested, "fa")
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap7 := func() gopurs_runtime.Value {
												arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
													unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
													for i, v := range arr {
														unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
													}
													return unboxed
												}()
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
												}
												return gopurs_runtime.Array(boxed)
											}()
											_ = arr_val_arrayMap7
											arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
											_ = arr_go_arrayMap7
											res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
											_ = res_go_arrayMap7
											for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
												res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_2_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)}))}
												}), v_arrayMap7)
											}
											return gopurs_runtime.Array(res_go_arrayMap7)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()).UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())})
							return &clone
						}()}))}
						goto end_branch_4
					} else {

					}
				}
				{
					if m_1.Type == 9 && m_1.IntVal == 108241190 {
						// TAST (Let): __local_var_2_2 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
						__local_var_2_2 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0)
						_ = __local_var_2_2
						// TAST (Let): __local_var_2_3 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
						__local_var_2_3 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0)
						_ = __local_var_2_3
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V2, func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V3)
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_0, v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V4
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_0, v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V5, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V6, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V6, "a")), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap6 := gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V6, "fa")
								_ = arr_val_arrayMap6
								arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
								_ = arr_go_arrayMap6
								res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
								_ = res_go_arrayMap6
								for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()), func() gopurs_runtime.Value {
							arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap6 := func() gopurs_runtime.Value {
											arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V6, "recursiveA").UnsafePtr)
												unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
												for i, v := range arr {
													unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
												}
												return unboxed
											}()
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
											}
											return gopurs_runtime.Array(boxed)
										}()
										_ = arr_val_arrayMap6
										arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
										_ = arr_go_arrayMap6
										res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
										_ = res_go_arrayMap6
										for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
											res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_2_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)}))}
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
									unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
								}
								return unboxed
							}()
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V6, "zArrayA").UnsafePtr)
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
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())}), func() *struct {
							nested gopurs_runtime.Value
						} {
							clone := *((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V7)
							clone.nested = gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V7.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V7.nested, "a")), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V7.nested, "fa")
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap7 := func() gopurs_runtime.Value {
												arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
													unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
													for i, v := range arr {
														unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
													}
													return unboxed
												}()
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
												}
												return gopurs_runtime.Array(boxed)
											}()
											_ = arr_val_arrayMap7
											arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
											_ = arr_go_arrayMap7
											res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
											_ = res_go_arrayMap7
											for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
												res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_2_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)}))}
												}), v_arrayMap7)
											}
											return gopurs_runtime.Array(res_go_arrayMap7)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()).UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())})
							return &clone
						}()}))}
						goto end_branch_4
					} else {

					}
				}
				{
					if m_1.Type == 9 && m_1.IntVal == 2066233029 {
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() [][][]gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap5 := func() gopurs_runtime.Value {
										arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0
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
									_ = arr_val_arrayMap5
									arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
									_ = arr_go_arrayMap5
									res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
									_ = res_go_arrayMap5
									for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
										res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_0)), v_arrayMap5)
									}
									return gopurs_runtime.Array(res_go_arrayMap5)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()).UnsafePtr)
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
			})
		})}))}
	})
	return cache_Main_functorM1
}

var cache_Main_f3Test gopurs_runtime.Value
var once_Main_f3Test sync.Once

func Get_Main_f3Test() gopurs_runtime.Value {
	once_Main_f3Test.Do(func() {
		cache_Main_f3Test = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=LitRecord bindingType=(Record (Row [eqRecord: (Func [(TypeApp (ADT ["Type","Proxy","Proxy"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)])]), (Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_0_0 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray(), "eq"), gopurs_runtime.RecordGet(ra_1, "zArrayA"), gopurs_runtime.RecordGet(rb_2, "zArrayA")).IntVal) != (0))
					})
				})
			}))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_2 shape=LitRecord bindingType=(Record (Row [eqRecord: (Func [(TypeApp (ADT ["Type","Proxy","Proxy"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)])]), (Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_1_2 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray(), "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
							})
						}), gopurs_runtime.RecordGet(ra_2, "recursiveA"), gopurs_runtime.RecordGet(rb_3, "recursiveA")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_0_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_2, rb_3).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_1_2
			// TAST (Let): __local_var_2_3 shape=LitRecord bindingType=(Record (Row [eqRecord: (Func [(TypeApp (ADT ["Type","Proxy","Proxy"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)])]), (Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_2_3 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_3, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_4, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_2, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_3, rb_4).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_2_3
			// TAST (Let): __local_var_3_4 shape=LitRecord bindingType=(Record (Row [eqRecord: (Func [(TypeApp (ADT ["Type","Proxy","Proxy"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)])]), (Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_3_4 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray(), "eq"), gopurs_runtime.RecordGet(ra_4, "fa"), gopurs_runtime.RecordGet(rb_5, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_4, rb_5).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_3_4
			// TAST (Let): __local_var_4_5 shape=LitRecord bindingType=(Record (Row [eqRecord: (Func [(TypeApp (ADT ["Type","Proxy","Proxy"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)])]), (Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_4_5 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_5, "fIgnore"), gopurs_runtime.RecordGet(rb_6, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_4, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_4_5
			// TAST (Let): __local_var_5_6 shape=LitRecord bindingType=(Record (Row [eqRecord: (Func [(TypeApp (ADT ["Type","Proxy","Proxy"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)])]), (Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_5_6 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_5, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_6, rb_7).IntVal) != (0))
					})
				})
			}))
			_ = __local_var_5_6
			// TAST (Let): __local_var_6_7 shape=LitRecord bindingType=(Record (Row [eqRecord: (Func [(TypeApp (ADT ["Type","Proxy","Proxy"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)])]), (Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_6_7 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_7, "arrayIgnore"), gopurs_runtime.RecordGet(rb_8, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_6, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_6_7
			// TAST (Let): __local_var_1_1 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Record (Row [] (TypeVar row))), (Record (Row [] (TypeVar row)))] Boolean)] Any))
			__local_var_1_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_7, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_8, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_7, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8).IntVal) != (0)))
				})
			}))
			_ = __local_var_1_1
			return gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - Fun3"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), gopurs_runtime.RecordGet(ra_2, "nested"), gopurs_runtime.RecordGet(rb_3, "nested")).IntVal) != (0))
				})
			}))), func() gopurs_runtime.Value {
				arr := func() [][][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap2 := func() gopurs_runtime.Value {
							arr := [][][]gopurs_runtime.Value{[][]gopurs_runtime.Value{func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordDict1("nested", Get_Main_recordValueL())}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()}}
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
						_ = arr_val_arrayMap2
						arr_go_arrayMap2 := (*[]gopurs_runtime.Value)(arr_val_arrayMap2.UnsafePtr)
						_ = arr_go_arrayMap2
						res_go_arrayMap2 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap2))
						_ = res_go_arrayMap2
						for i_arrayMap2, v_arrayMap2 := range *arr_go_arrayMap2 {
							res_go_arrayMap2[i_arrayMap2] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Func(func(v1_0 gopurs_runtime.Value) gopurs_runtime.Value {
								// TAST (Let): __local_var_1_8 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
								__local_var_1_8 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Show_showIntImpl())
								_ = __local_var_1_8
								return gopurs_runtime.RecordUpdate1(v1_0, "nested", gopurs_runtime.RecordUpdateDict(gopurs_runtime.RecordGet(v1_0, "nested"), []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_0, "nested"), "a")).StrVal()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap8 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_0, "nested"), "fa")
										_ = arr_val_arrayMap8
										arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
										_ = arr_go_arrayMap8
										res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
										_ = res_go_arrayMap8
										for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
											res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), v_arrayMap8)
										}
										return gopurs_runtime.Array(res_go_arrayMap8)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), func() gopurs_runtime.Value {
									arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
												arr_val_arrayMap8 := func() gopurs_runtime.Value {
													arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
														arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_0, "nested"), "recursiveA").UnsafePtr)
														unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
														for i, v := range arr {
															unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
														}
														return unboxed
													}()
													boxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
													}
													return gopurs_runtime.Array(boxed)
												}()
												_ = arr_val_arrayMap8
												arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
												_ = arr_go_arrayMap8
												res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
												_ = res_go_arrayMap8
												for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
													res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_1_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)}))}
													}), v_arrayMap8)
												}
												return gopurs_runtime.Array(res_go_arrayMap8)
											}().UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()).UnsafePtr)
										unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
										for i, v := range arr {
											unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
									}
									return gopurs_runtime.Array(boxed)
								}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_0, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())
										_ = arr_val_arrayMap8
										arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
										_ = arr_go_arrayMap8
										res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
										_ = res_go_arrayMap8
										for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
											res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), v_arrayMap8)
										}
										return gopurs_runtime.Array(res_go_arrayMap8)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())}))
							}))), v_arrayMap2)
						}
						return gopurs_runtime.Array(res_go_arrayMap2)
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
				}()
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
				arr := [][][]gopurs_runtime.Value{[][]gopurs_runtime.Value{func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordDict1("nested", Get_Main_recordValueR())}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()}}
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
		}()
	})
	return cache_Main_f3Test
}

var cache_Main_f2Test gopurs_runtime.Value
var once_Main_f2Test sync.Once

func Get_Main_f2Test() gopurs_runtime.Value {
	once_Main_f2Test.Do(func() {
		cache_Main_f2Test = gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - Fun2"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(Get_Main_eqArray(), "eq"), func() gopurs_runtime.Value {
			arr := func() [][]string {
				arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr_val_arrayMap2 := func() gopurs_runtime.Value {
						arr := [][]int64{[]int64{3}}
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
					_ = arr_val_arrayMap2
					arr_go_arrayMap2 := (*[]gopurs_runtime.Value)(arr_val_arrayMap2.UnsafePtr)
					_ = arr_go_arrayMap2
					res_go_arrayMap2 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap2))
					_ = res_go_arrayMap2
					for i_arrayMap2, v_arrayMap2 := range *arr_go_arrayMap2 {
						res_go_arrayMap2[i_arrayMap2] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Show_showIntImpl()), v_arrayMap2)
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
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap2 := func() gopurs_runtime.Value {
							arr := [][]int64{[]int64{3}}
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
						_ = arr_val_arrayMap2
						arr_go_arrayMap2 := (*[]gopurs_runtime.Value)(arr_val_arrayMap2.UnsafePtr)
						_ = arr_go_arrayMap2
						res_go_arrayMap2 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap2))
						_ = res_go_arrayMap2
						for i_arrayMap2, v_arrayMap2 := range *arr_go_arrayMap2 {
							res_go_arrayMap2[i_arrayMap2] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Show_showIntImpl()), v_arrayMap2)
						}
						return gopurs_runtime.Array(res_go_arrayMap2)
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
	})
	return cache_Main_f2Test
}

var cache_Main_f1Test gopurs_runtime.Value
var once_Main_f1Test sync.Once

func Get_Main_f1Test() gopurs_runtime.Value {
	once_Main_f1Test.Do(func() {
		cache_Main_f1Test = gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - Fun1"), gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(3)).StrVal()) == (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(3)).StrVal())))
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

var cache_Main_eqM1 gopurs_runtime.Value
var once_Main_eqM1 sync.Once

func Get_Main_eqM1() gopurs_runtime.Value {
	once_Main_eqM1.Do(func() {
		cache_Main_eqM1 = func() gopurs_runtime.Value {
			// TAST (Let): eqArray3_0_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
			eqArray3_0_0 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray3_0_0
			// TAST (Let): eqArray4_1_1 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
			eqArray4_1_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
			_ = eqArray4_1_1
			// TAST (Let): eqArray5_2_2 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
			eqArray5_2_2 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_1_1, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
				})
			}))})
			_ = eqArray5_2_2
			// TAST (Let): eqArray6_3_3 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a))))])
			eqArray6_3_3 := (&Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(eqArray4_1_1, "eq")))})
			_ = eqArray6_3_3
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t4 bool
					{
						if x_4.Type == 9 && x_4.IntVal == 3852365315 {
							__t4 = (y_5.Type == 9 && y_5.IntVal == 3852365315) && ((((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_0.V0), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1)).IntVal) != (0)))
							goto end_branch_4
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 769986722 {
							__t4 = (y_5.Type == 9 && y_5.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0) == ((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0))
							goto end_branch_4
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 2727978561 {
							__t4 = (y_5.Type == 9 && y_5.IntVal == 2727978561) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0))
							goto end_branch_4
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 1830062304 {
							__t4 = (y_5.Type == 9 && y_5.IntVal == 1830062304) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_2.V0), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}(), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
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
						if x_4.Type == 9 && x_4.IntVal == 3190619783 {
							__t4 = (y_5.Type == 9 && y_5.IntVal == 3190619783) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_2.V0), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}(), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
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
						if x_4.Type == 9 && x_4.IntVal == 108241190 {
							__t4 = (y_5.Type == 9 && y_5.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)) && (((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V2
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V2
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_0.V0), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_2.V0), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V6, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}(), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V6, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0)))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V7.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V7.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_2.V0), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}(), func() gopurs_runtime.Value {
								arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
									unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
									for i, v := range arr {
										unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
									}
									return unboxed
								}()
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
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
						__t4 = (x_4.Type == 9 && x_4.IntVal == 2066233029) && ((y_5.Type == 9 && y_5.IntVal == 2066233029) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray6_3_3.V0), func() gopurs_runtime.Value {
							arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0
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
							arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0
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
				})
			})}))}
		}()
	})
	return cache_Main_eqM1
}

var cache_Main_maTests gopurs_runtime.Value
var once_Main_maTests sync.Once

func Get_Main_maTests() gopurs_runtime.Value {
	once_Main_maTests.Do(func() {
		cache_Main_maTests = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): eqArray3_0_1 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
			eqArray3_0_1 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray3_0_1
			// TAST (Let): eqArray4_1_2 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
			eqArray4_1_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
			_ = eqArray4_1_2
			// TAST (Let): eqArray5_2_3 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
			eqArray5_2_3 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_1_2, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
				})
			}))})
			_ = eqArray5_2_3
			// TAST (Let): eqArray6_3_4 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a))))])
			eqArray6_3_4 := (&Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(eqArray4_1_2, "eq")))})
			_ = eqArray6_3_4
			// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=(ADT ["Main","M"] [(ADT ["Prim","Array"] []), String])
			__local_var_4_5 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorM1()).V0), Get_Data_Show_showIntImpl(), Get_Main_m0L())
			_ = __local_var_4_5
			var __t14 bool
			{
				if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 3852365315 {
					var __t_tag_8 gopurs_runtime.Value = Get_Main_m0R()
					__t14 = (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3852365315) && ((((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_1.V0), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V1)).IntVal) != (0)))
					goto end_branch_14
				} else {

				}
			}
			{
				if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 769986722 {
					var __t_tag_9 gopurs_runtime.Value = Get_Main_m0R()
					__t14 = (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0) == ((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0))
					goto end_branch_14
				} else {

				}
			}
			{
				if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 2727978561 {
					var __t_tag_10 gopurs_runtime.Value = Get_Main_m0R()
					__t14 = (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 2727978561) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0).IntVal) != (0))
					goto end_branch_14
				} else {

				}
			}
			{
				if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 1830062304 {
					var __t_tag_11 gopurs_runtime.Value = Get_Main_m0R()
					__t14 = (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 1830062304) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_3.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_1.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_14
				} else {

				}
			}
			{
				if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 3190619783 {
					var __t_tag_12 gopurs_runtime.Value = Get_Main_m0R()
					__t14 = (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 3190619783) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_3.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_1.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_14
				} else {

				}
			}
			{
				if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 108241190 {
					var __t_tag_13 gopurs_runtime.Value = Get_Main_m0R()
					__t14 = (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0)) && (((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_1.V0), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_3.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_1.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V7.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V7.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_2_3.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_0_1.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0))))
					goto end_branch_14
				} else {

				}
			}
			{
				var __t_and_7 bool = false
				if __local_var_4_5.Type == 9 && __local_var_4_5.IntVal == 2066233029 {

					var __t_tag_6 gopurs_runtime.Value = Get_Main_m0R()
					__t_and_7 = (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2066233029) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray6_3_4.V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_5.UnsafePtr).V0
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
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m0R().UnsafePtr).V0
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
					}()).IntVal) != (0))
				}
				__t14 = __t_and_7
			}
		end_branch_14:
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M0"), gopurs_runtime.Bool(__t14))
			_ = __local_var_0_0
			_dollar___unused_1_15 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_15
			// TAST (Let): __local_var_2_17 shape=App(Other) bindingType=(ADT ["Main","M"] [(ADT ["Prim","Array"] []), String])
			__local_var_2_17 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorM1()).V0), Get_Data_Show_showIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, int64]{1, 0}))})
			_ = __local_var_2_17
			var __t18 bool
			{
				if __local_var_2_17.Type == 9 && __local_var_2_17.IntVal == 3852365315 {
					__t18 = false
					goto end_branch_18
				} else {

				}
			}
			{
				__t18 = (__local_var_2_17.Type == 9 && __local_var_2_17.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_17.UnsafePtr).V0) == (0))
			}
		end_branch_18:
			_dollar___unused_2_16 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M1"), gopurs_runtime.Bool(__t18)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_16
			// TAST (Let): eqArray3_3_20 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
			eqArray3_3_20 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray3_3_20
			// TAST (Let): eqArray4_4_21 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
			eqArray4_4_21 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
			_ = eqArray4_4_21
			// TAST (Let): eqArray5_5_22 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
			eqArray5_5_22 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_4_21, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V1).IntVal) != (0)))
				})
			}))})
			_ = eqArray5_5_22
			// TAST (Let): eqArray6_6_23 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a))))])
			eqArray6_6_23 := (&Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(eqArray4_4_21, "eq")))})
			_ = eqArray6_6_23
			// TAST (Let): __local_var_7_24 shape=App(Other) bindingType=(ADT ["Main","M"] [(ADT ["Prim","Array"] []), String])
			__local_var_7_24 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorM1()).V0), Get_Data_Show_showIntImpl(), Get_Main_m2L())
			_ = __local_var_7_24
			var __t33 bool
			{
				if __local_var_7_24.Type == 9 && __local_var_7_24.IntVal == 3852365315 {
					var __t_tag_27 gopurs_runtime.Value = Get_Main_m2R()
					__t33 = (__t_tag_27.Type == 9 && __t_tag_27.IntVal == 3852365315) && ((((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_3_20.V0), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V1)).IntVal) != (0)))
					goto end_branch_33
				} else {

				}
			}
			{
				if __local_var_7_24.Type == 9 && __local_var_7_24.IntVal == 769986722 {
					var __t_tag_28 gopurs_runtime.Value = Get_Main_m2R()
					__t33 = (__t_tag_28.Type == 9 && __t_tag_28.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0) == ((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0))
					goto end_branch_33
				} else {

				}
			}
			{
				if __local_var_7_24.Type == 9 && __local_var_7_24.IntVal == 2727978561 {
					var __t_tag_29 gopurs_runtime.Value = Get_Main_m2R()
					__t33 = (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 2727978561) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0).IntVal) != (0))
					goto end_branch_33
				} else {

				}
			}
			{
				if __local_var_7_24.Type == 9 && __local_var_7_24.IntVal == 1830062304 {
					var __t_tag_30 gopurs_runtime.Value = Get_Main_m2R()
					__t33 = (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 1830062304) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_22.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_3_20.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_33
				} else {

				}
			}
			{
				if __local_var_7_24.Type == 9 && __local_var_7_24.IntVal == 3190619783 {
					var __t_tag_31 gopurs_runtime.Value = Get_Main_m2R()
					__t33 = (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 3190619783) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_22.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_3_20.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_33
				} else {

				}
			}
			{
				if __local_var_7_24.Type == 9 && __local_var_7_24.IntVal == 108241190 {
					var __t_tag_32 gopurs_runtime.Value = Get_Main_m2R()
					__t33 = (__t_tag_32.Type == 9 && __t_tag_32.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0)) && (((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_3_20.V0), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_22.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_3_20.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V7.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V7.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_22.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_3_20.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0))))
					goto end_branch_33
				} else {

				}
			}
			{
				var __t_and_26 bool = false
				if __local_var_7_24.Type == 9 && __local_var_7_24.IntVal == 2066233029 {

					var __t_tag_25 gopurs_runtime.Value = Get_Main_m2R()
					__t_and_26 = (__t_tag_25.Type == 9 && __t_tag_25.IntVal == 2066233029) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray6_6_23.V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_24.UnsafePtr).V0
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
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m2R().UnsafePtr).V0
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
					}()).IntVal) != (0))
				}
				__t33 = __t_and_26
			}
		end_branch_33:
			_dollar___unused_3_19 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M2"), gopurs_runtime.Bool(__t33)), gopurs_runtime.Value{})
			_ = _dollar___unused_3_19
			// TAST (Let): eqArray3_4_35 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
			eqArray3_4_35 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray3_4_35
			// TAST (Let): eqArray4_5_36 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
			eqArray4_5_36 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
			_ = eqArray4_5_36
			// TAST (Let): eqArray5_6_37 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
			eqArray5_6_37 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_5_36, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V1).IntVal) != (0)))
				})
			}))})
			_ = eqArray5_6_37
			// TAST (Let): __local_var_7_38 shape=App(Other) bindingType=(ADT ["Main","M"] [(ADT ["Prim","Array"] []), String])
			__local_var_7_38 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorM1()).V0), Get_Data_Show_showIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, int64]{1, Get_Main_recordValueL()}))})
			_ = __local_var_7_38
			var __t39 bool
			{
				if __local_var_7_38.Type == 9 && __local_var_7_38.IntVal == 3852365315 {
					__t39 = false
					goto end_branch_39
				} else {

				}
			}
			{
				if __local_var_7_38.Type == 9 && __local_var_7_38.IntVal == 769986722 {
					__t39 = false
					goto end_branch_39
				} else {

				}
			}
			{
				if __local_var_7_38.Type == 9 && __local_var_7_38.IntVal == 2727978561 {
					__t39 = false
					goto end_branch_39
				} else {

				}
			}
			{
				__t39 = (__local_var_7_38.Type == 9 && __local_var_7_38.IntVal == 1830062304) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_38.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.Str("71").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_38.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Get_Main_recordValueR(), "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_38.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet(Get_Main_recordValueR(), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_38.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet(Get_Main_recordValueR(), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_38.UnsafePtr).V0, "ignore").IntVal) == (91))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_6_37.V0), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_38.UnsafePtr).V0, "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Get_Main_recordValueR(), "recursiveA").UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_4_35.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_7_38.UnsafePtr).V0, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Get_Main_recordValueR(), "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0)))
			}
		end_branch_39:
			_dollar___unused_4_34 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M3"), gopurs_runtime.Bool(__t39)), gopurs_runtime.Value{})
			_ = _dollar___unused_4_34
			// TAST (Let): eqArray3_5_41 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
			eqArray3_5_41 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray3_5_41
			// TAST (Let): eqArray4_6_42 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
			eqArray4_6_42 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
			_ = eqArray4_6_42
			// TAST (Let): eqArray5_7_43 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
			eqArray5_7_43 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_7.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_6_42, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_8.UnsafePtr).V1).IntVal) != (0)))
				})
			}))})
			_ = eqArray5_7_43
			// TAST (Let): eqArray6_8_44 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a))))])
			eqArray6_8_44 := (&Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(eqArray4_6_42, "eq")))})
			_ = eqArray6_8_44
			// TAST (Let): __local_var_9_45 shape=App(Other) bindingType=(ADT ["Main","M"] [(ADT ["Prim","Array"] []), String])
			__local_var_9_45 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorM1()).V0), Get_Data_Show_showIntImpl(), Get_Main_m4L())
			_ = __local_var_9_45
			var __t54 bool
			{
				if __local_var_9_45.Type == 9 && __local_var_9_45.IntVal == 3852365315 {
					var __t_tag_48 gopurs_runtime.Value = Get_Main_m4R()
					__t54 = (__t_tag_48.Type == 9 && __t_tag_48.IntVal == 3852365315) && ((((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_5_41.V0), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V1)).IntVal) != (0)))
					goto end_branch_54
				} else {

				}
			}
			{
				if __local_var_9_45.Type == 9 && __local_var_9_45.IntVal == 769986722 {
					var __t_tag_49 gopurs_runtime.Value = Get_Main_m4R()
					__t54 = (__t_tag_49.Type == 9 && __t_tag_49.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0) == ((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0))
					goto end_branch_54
				} else {

				}
			}
			{
				if __local_var_9_45.Type == 9 && __local_var_9_45.IntVal == 2727978561 {
					var __t_tag_50 gopurs_runtime.Value = Get_Main_m4R()
					__t54 = (__t_tag_50.Type == 9 && __t_tag_50.IntVal == 2727978561) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0).IntVal) != (0))
					goto end_branch_54
				} else {

				}
			}
			{
				if __local_var_9_45.Type == 9 && __local_var_9_45.IntVal == 1830062304 {
					var __t_tag_51 gopurs_runtime.Value = Get_Main_m4R()
					__t54 = (__t_tag_51.Type == 9 && __t_tag_51.IntVal == 1830062304) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_43.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_5_41.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_54
				} else {

				}
			}
			{
				if __local_var_9_45.Type == 9 && __local_var_9_45.IntVal == 3190619783 {
					var __t_tag_52 gopurs_runtime.Value = Get_Main_m4R()
					__t54 = (__t_tag_52.Type == 9 && __t_tag_52.IntVal == 3190619783) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_43.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_5_41.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_54
				} else {

				}
			}
			{
				if __local_var_9_45.Type == 9 && __local_var_9_45.IntVal == 108241190 {
					var __t_tag_53 gopurs_runtime.Value = Get_Main_m4R()
					__t54 = (__t_tag_53.Type == 9 && __t_tag_53.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0)) && (((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_5_41.V0), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_43.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_5_41.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V7.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V7.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_43.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_5_41.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0))))
					goto end_branch_54
				} else {

				}
			}
			{
				var __t_and_47 bool = false
				if __local_var_9_45.Type == 9 && __local_var_9_45.IntVal == 2066233029 {

					var __t_tag_46 gopurs_runtime.Value = Get_Main_m4R()
					__t_and_47 = (__t_tag_46.Type == 9 && __t_tag_46.IntVal == 2066233029) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray6_8_44.V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_9_45.UnsafePtr).V0
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
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m4R().UnsafePtr).V0
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
					}()).IntVal) != (0))
				}
				__t54 = __t_and_47
			}
		end_branch_54:
			_dollar___unused_5_40 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M4"), gopurs_runtime.Bool(__t54)), gopurs_runtime.Value{})
			_ = _dollar___unused_5_40
			// TAST (Let): eqArray3_6_56 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
			eqArray3_6_56 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray3_6_56
			// TAST (Let): eqArray4_7_57 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
			eqArray4_7_57 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
			_ = eqArray4_7_57
			// TAST (Let): eqArray5_8_58 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
			eqArray5_8_58 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_8.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_7_57, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_9.UnsafePtr).V1).IntVal) != (0)))
				})
			}))})
			_ = eqArray5_8_58
			// TAST (Let): eqArray6_9_59 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a))))])
			eqArray6_9_59 := (&Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(eqArray4_7_57, "eq")))})
			_ = eqArray6_9_59
			// TAST (Let): __local_var_10_60 shape=App(Other) bindingType=(ADT ["Main","M"] [(ADT ["Prim","Array"] []), String])
			__local_var_10_60 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorM1()).V0), Get_Data_Show_showIntImpl(), Get_Main_m5L())
			_ = __local_var_10_60
			var __t69 bool
			{
				if __local_var_10_60.Type == 9 && __local_var_10_60.IntVal == 3852365315 {
					var __t_tag_63 gopurs_runtime.Value = Get_Main_m5R()
					__t69 = (__t_tag_63.Type == 9 && __t_tag_63.IntVal == 3852365315) && ((((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_6_56.V0), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V1)).IntVal) != (0)))
					goto end_branch_69
				} else {

				}
			}
			{
				if __local_var_10_60.Type == 9 && __local_var_10_60.IntVal == 769986722 {
					var __t_tag_64 gopurs_runtime.Value = Get_Main_m5R()
					__t69 = (__t_tag_64.Type == 9 && __t_tag_64.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0) == ((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0))
					goto end_branch_69
				} else {

				}
			}
			{
				if __local_var_10_60.Type == 9 && __local_var_10_60.IntVal == 2727978561 {
					var __t_tag_65 gopurs_runtime.Value = Get_Main_m5R()
					__t69 = (__t_tag_65.Type == 9 && __t_tag_65.IntVal == 2727978561) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0).IntVal) != (0))
					goto end_branch_69
				} else {

				}
			}
			{
				if __local_var_10_60.Type == 9 && __local_var_10_60.IntVal == 1830062304 {
					var __t_tag_66 gopurs_runtime.Value = Get_Main_m5R()
					__t69 = (__t_tag_66.Type == 9 && __t_tag_66.IntVal == 1830062304) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_58.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_6_56.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_69
				} else {

				}
			}
			{
				if __local_var_10_60.Type == 9 && __local_var_10_60.IntVal == 3190619783 {
					var __t_tag_67 gopurs_runtime.Value = Get_Main_m5R()
					__t69 = (__t_tag_67.Type == 9 && __t_tag_67.IntVal == 3190619783) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_58.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_6_56.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_69
				} else {

				}
			}
			{
				if __local_var_10_60.Type == 9 && __local_var_10_60.IntVal == 108241190 {
					var __t_tag_68 gopurs_runtime.Value = Get_Main_m5R()
					__t69 = (__t_tag_68.Type == 9 && __t_tag_68.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0)) && (((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_6_56.V0), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_58.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_6_56.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V7.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V7.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_58.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_6_56.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0))))
					goto end_branch_69
				} else {

				}
			}
			{
				var __t_and_62 bool = false
				if __local_var_10_60.Type == 9 && __local_var_10_60.IntVal == 2066233029 {

					var __t_tag_61 gopurs_runtime.Value = Get_Main_m5R()
					__t_and_62 = (__t_tag_61.Type == 9 && __t_tag_61.IntVal == 2066233029) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray6_9_59.V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_10_60.UnsafePtr).V0
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
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m5R().UnsafePtr).V0
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
					}()).IntVal) != (0))
				}
				__t69 = __t_and_62
			}
		end_branch_69:
			_dollar___unused_6_55 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M5"), gopurs_runtime.Bool(__t69)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_55
			// TAST (Let): eqArray3_7_70 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
			eqArray3_7_70 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray3_7_70
			// TAST (Let): eqArray4_8_71 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
			eqArray4_8_71 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
			_ = eqArray4_8_71
			// TAST (Let): eqArray5_9_72 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
			eqArray5_9_72 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_9.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_8_71, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_9.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_10.UnsafePtr).V1).IntVal) != (0)))
				})
			}))})
			_ = eqArray5_9_72
			// TAST (Let): eqArray6_10_73 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a))))])
			eqArray6_10_73 := (&Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(eqArray4_8_71, "eq")))})
			_ = eqArray6_10_73
			// TAST (Let): __local_var_11_74 shape=App(Other) bindingType=(ADT ["Main","M"] [(ADT ["Prim","Array"] []), String])
			__local_var_11_74 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorM1()).V0), Get_Data_Show_showIntImpl(), Get_Main_m6L())
			_ = __local_var_11_74
			var __t83 bool
			{
				if __local_var_11_74.Type == 9 && __local_var_11_74.IntVal == 3852365315 {
					var __t_tag_77 gopurs_runtime.Value = Get_Main_m6R()
					__t83 = (__t_tag_77.Type == 9 && __t_tag_77.IntVal == 3852365315) && ((((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_7_70.V0), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V1)).IntVal) != (0)))
					goto end_branch_83
				} else {

				}
			}
			{
				if __local_var_11_74.Type == 9 && __local_var_11_74.IntVal == 769986722 {
					var __t_tag_78 gopurs_runtime.Value = Get_Main_m6R()
					__t83 = (__t_tag_78.Type == 9 && __t_tag_78.IntVal == 769986722) && (((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0) == ((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0))
					goto end_branch_83
				} else {

				}
			}
			{
				if __local_var_11_74.Type == 9 && __local_var_11_74.IntVal == 2727978561 {
					var __t_tag_79 gopurs_runtime.Value = Get_Main_m6R()
					__t83 = (__t_tag_79.Type == 9 && __t_tag_79.IntVal == 2727978561) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0).IntVal) != (0))
					goto end_branch_83
				} else {

				}
			}
			{
				if __local_var_11_74.Type == 9 && __local_var_11_74.IntVal == 1830062304 {
					var __t_tag_80 gopurs_runtime.Value = Get_Main_m6R()
					__t83 = (__t_tag_80.Type == 9 && __t_tag_80.IntVal == 1830062304) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_72.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_7_70.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_83
				} else {

				}
			}
			{
				if __local_var_11_74.Type == 9 && __local_var_11_74.IntVal == 3190619783 {
					var __t_tag_81 gopurs_runtime.Value = Get_Main_m6R()
					__t83 = (__t_tag_81.Type == 9 && __t_tag_81.IntVal == 3190619783) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_72.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_7_70.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_83
				} else {

				}
			}
			{
				if __local_var_11_74.Type == 9 && __local_var_11_74.IntVal == 108241190 {
					var __t_tag_82 gopurs_runtime.Value = Get_Main_m6R()
					__t83 = (__t_tag_82.Type == 9 && __t_tag_82.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0)) && (((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_7_70.V0), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_72.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_7_70.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))) && ((((((((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V7.nested, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V7.nested, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_72.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_7_70.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0))))
					goto end_branch_83
				} else {

				}
			}
			{
				var __t_and_76 bool = false
				if __local_var_11_74.Type == 9 && __local_var_11_74.IntVal == 2066233029 {

					var __t_tag_75 gopurs_runtime.Value = Get_Main_m6R()
					__t_and_76 = (__t_tag_75.Type == 9 && __t_tag_75.IntVal == 2066233029) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray6_10_73.V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_11_74.UnsafePtr).V0
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
						arr := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(Get_Main_m6R().UnsafePtr).V0
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
					}()).IntVal) != (0))
				}
				__t83 = __t_and_76
			}
		end_branch_83:
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("map - M6"), gopurs_runtime.Bool(__t83)), gopurs_runtime.Value{})
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

type Constructor_Main_T[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M0[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
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
	V0 *struct {
		nested gopurs_runtime.Value
	}
}

type Constructor_Main_M5[T_f any, T_a any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 []int64
	V3 []gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
	V6 gopurs_runtime.Value
	V7 *struct {
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

func Call_Main_eqTuple(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_0, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1).IntVal) != (0)))
		})
	})}))}
}

func Call_Main_functorFun3(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
				__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
				_ = __local_var_4_1
				return gopurs_runtime.RecordUpdate1(v1_3, "nested", gopurs_runtime.RecordUpdateDict(gopurs_runtime.RecordGet(v1_3, "nested"), []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa")), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap11 := func() gopurs_runtime.Value {
									arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "recursiveA").UnsafePtr)
										unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
										for i, v := range arr {
											unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
										}
										return unboxed
									}()
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
									}
									return gopurs_runtime.Array(boxed)
								}()
								_ = arr_val_arrayMap11
								arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
								_ = arr_go_arrayMap11
								res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
								_ = res_go_arrayMap11
								for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
									res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_4_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1)}))}
									}), v_arrayMap11)
								}
								return gopurs_runtime.Array(res_go_arrayMap11)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
						unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
						for i, v := range arr {
							unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
						}
						return unboxed
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap11 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						_ = arr_val_arrayMap11
						arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
						_ = arr_go_arrayMap11
						res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
						_ = res_go_arrayMap11
						for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
							res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(f_1, v_arrayMap11)
						}
						return gopurs_runtime.Array(res_go_arrayMap11)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())}))
			}))))
			_ = __local_var_3_0
			return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Apply(m_2, x_4))
			})
		})
	})}))}
}

func Call_Main_functorM(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if m_2.Type == 9 && m_2.IntVal == 3852365315 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer((&Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap5 := gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)
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
					// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
					__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
					_ = __local_var_3_0
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "fa")), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap6 := func() gopurs_runtime.Value {
										arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "recursiveA").UnsafePtr)
											unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
											for i, v := range arr {
												unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
											}
											return unboxed
										}()
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
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
											return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}))}
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
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "zArrayA").UnsafePtr)
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
					}())})}))}
					goto end_branch_4
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 3190619783 {
					// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
					__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
					_ = __local_var_3_1
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() *struct {
						nested gopurs_runtime.Value
					} {
						clone := *((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)
						clone.nested = gopurs_runtime.RecordUpdateDict((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "fa")), func() gopurs_runtime.Value {
							arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap7 := func() gopurs_runtime.Value {
											arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
												unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
												for i, v := range arr {
													unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
												}
												return unboxed
											}()
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
											}
											return gopurs_runtime.Array(boxed)
										}()
										_ = arr_val_arrayMap7
										arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
										_ = arr_go_arrayMap7
										res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
										_ = res_go_arrayMap7
										for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
											res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}))}
											}), v_arrayMap7)
										}
										return gopurs_runtime.Array(res_go_arrayMap7)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()).UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
								for i, v := range arr {
									unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
								}
								return unboxed
							}()
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_1, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())})
						return &clone
					}()}))}
					goto end_branch_4
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 108241190 {
					// TAST (Let): __local_var_3_2 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
					__local_var_3_2 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
					_ = __local_var_3_2
					// TAST (Let): __local_var_3_3 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
					__local_var_3_3 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)
					_ = __local_var_3_3
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2, func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap5 := gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3)
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
					}(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5, gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "fa")), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap6 := func() gopurs_runtime.Value {
										arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "recursiveA").UnsafePtr)
											unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
											for i, v := range arr {
												unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
											}
											return unboxed
										}()
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
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
											return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}))}
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
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "zArrayA").UnsafePtr)
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
					}())}), func() *struct {
						nested gopurs_runtime.Value
					} {
						clone := *((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7)
						clone.nested = gopurs_runtime.RecordUpdateDict((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, []string{"a", "fa", "recursiveA", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "fa")), func() gopurs_runtime.Value {
							arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap7 := func() gopurs_runtime.Value {
											arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
												unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
												for i, v := range arr {
													unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
												}
												return unboxed
											}()
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
											}
											return gopurs_runtime.Array(boxed)
										}()
										_ = arr_val_arrayMap7
										arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
										_ = arr_go_arrayMap7
										res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
										_ = res_go_arrayMap7
										for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
											res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_3_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)}))}
											}), v_arrayMap7)
										}
										return gopurs_runtime.Array(res_go_arrayMap7)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()).UnsafePtr)
								unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
								for i, v := range arr {
									unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
								}
								return unboxed
							}()
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_1, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())})
						return &clone
					}()}))}
					goto end_branch_4
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 2066233029 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() [][][]gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap5 := func() gopurs_runtime.Value {
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
								_ = arr_val_arrayMap5
								arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
								_ = arr_go_arrayMap5
								res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
								_ = res_go_arrayMap5
								for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
									res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_1)), v_arrayMap5)
								}
								return gopurs_runtime.Array(res_go_arrayMap5)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
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
		})
	})}))}
}

func Call_Main_eqM(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	// TAST (Let): eqArray3_2_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a))])
	eqArray3_2_0 := (&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_1, "eq"))})
	_ = eqArray3_2_0
	// TAST (Let): eqArray4_3_1 shape=LitRecord bindingType=(Record (Row [eq: (Func [(Array (TypeVar a)), (Array (TypeVar a))] Boolean)] Any))
	eqArray4_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_1, "eq")))
	_ = eqArray4_3_1
	// TAST (Let): eqArray5_4_2 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (ADT ["Data","Tuple","Tuple"] [Int, (Array (TypeVar a))]))])
	eqArray5_4_2 := (&Constructor_Data_Eq_Eq[[]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0.IntVal)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(eqArray4_3_1, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
		})
	}))})
	_ = eqArray5_4_2
	// TAST (Let): eqArray6_5_3 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Array (Array (Array (TypeVar a))))])
	eqArray6_5_3 := (&Constructor_Data_Eq_Eq[[][][]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(eqArray4_3_1, "eq")))})
	_ = eqArray6_5_3
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 bool
			{
				if x_6.Type == 9 && x_6.IntVal == 3852365315 {
					__t4 = (y_7.Type == 9 && y_7.IntVal == 3852365315) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_2_0.V0), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V1)).IntVal) != (0)))
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
					__t4 = (y_7.Type == 9 && y_7.IntVal == 1830062304) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_2.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_2_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
					__t4 = (y_7.Type == 9 && y_7.IntVal == 3190619783) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_2.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_2_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
					__t4 = (y_7.Type == 9 && y_7.IntVal == 108241190) && ((((((((((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0) == ((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V1, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_2_0.V0), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V4, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt()))}, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V5, (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V5).IntVal) != (0))) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_2.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V6, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_2_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
					}())).IntVal) != (0)))) && ((((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray2()).V0), func() gopurs_runtime.Value {
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_2.V0), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := func() []*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value] {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V7.nested, "recursiveA").UnsafePtr)
							unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value], len(arr))
							for i, v := range arr {
								unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, []gopurs_runtime.Value]](v)
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray3_2_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
				__t4 = (x_6.Type == 9 && x_6.IntVal == 2066233029) && ((y_7.Type == 9 && y_7.IntVal == 2066233029) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray6_5_3.V0), func() gopurs_runtime.Value {
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
		})
	})}))}
}
