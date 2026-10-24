package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{[]struct {
				nested struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{struct {
				nested struct {
					a           []string
					arrayIgnore []int64
					fIgnore     []int64
					fa          [][]string
					ignore      int64
					zArrayA     [][]string
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{func() struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{[]struct {
				nested struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{struct {
				nested struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{func() struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
		}(), struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
		}()}})))}
	})
	return cache_Main_m6_prime_
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
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
		}(), struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
		}()}})))}
	})
	return cache_Main_m6
}

var cache_Main_m5_prime_ gopurs_runtime.Value
var once_Main_m5_prime_ sync.Once

func Get_Main_m5_prime_() gopurs_runtime.Value {
	once_Main_m5_prime_.Do(func() {
		cache_Main_m5_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_358349645_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, []string]{1, struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           []string
				arrayIgnore []int64
				fIgnore     []int64
				fa          [][]string
				ignore      int64
				zArrayA     [][]string
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
		}()}})))}
	})
	return cache_Main_m5_prime_
}

var cache_Main_m5 gopurs_runtime.Value
var once_Main_m5 sync.Once

func Get_Main_m5() gopurs_runtime.Value {
	once_Main_m5.Do(func() {
		cache_Main_m5 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer(Rebox_Main_1302345387_2067946548((&Constructor_Main_M5[gopurs_runtime.Value, string]{1, struct {
			nested gopurs_runtime.Value
			/*gopurs:record-v1:1:006e00650073007400650064*/
		}{func() gopurs_runtime.Value {
			orig := func() struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
		}()}})))}
	})
	return cache_Main_m5
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
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
				/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
					/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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

var cache_Main_m2_prime_ gopurs_runtime.Value
var once_Main_m2_prime_ sync.Once

func Get_Main_m2_prime_() gopurs_runtime.Value {
	once_Main_m2_prime_.Do(func() {
		cache_Main_m2_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_2857385098_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, []string]{1, int64(0)})))}
	})
	return cache_Main_m2_prime_
}

var cache_Main_m2 gopurs_runtime.Value
var once_Main_m2 sync.Once

func Get_Main_m2() gopurs_runtime.Value {
	once_Main_m2.Do(func() {
		cache_Main_m2 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})))}
	})
	return cache_Main_m2
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

var cache_Main_m0_prime_ gopurs_runtime.Value
var once_Main_m0_prime_ sync.Once

func Get_Main_m0_prime_() gopurs_runtime.Value {
	once_Main_m0_prime_.Do(func() {
		cache_Main_m0_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3974551048_67812977(nil))}
	})
	return cache_Main_m0_prime_
}

var cache_Main_m0 gopurs_runtime.Value
var once_Main_m0 sync.Once

func Get_Main_m0() gopurs_runtime.Value {
	once_Main_m0.Do(func() {
		cache_Main_m0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}
	})
	return cache_Main_m0
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m0"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
			})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_0())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))})).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}})).UnsafePtr)))).IntVal) != (0)))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_1 shape=EffectBind(Let(Let(Other))) bindingType=Any
			__local_var_1_1 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_2
				// TAST (Let): __local_var_2_3 shape=App(Var) bindingType=Any
				__local_var_2_3 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m1"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("nested")
				})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("zArrayA")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("ignore")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fa")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("arrayIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("a")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("nested")
				})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("zArrayA")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("ignore")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fa")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("fIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("arrayIgnore")
				})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str("a")
				})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_1())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, Get_Main_m1())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m1()})).UnsafePtr)))).IntVal) != (0)))
				_ = __local_var_2_3
				// TAST (Let): __local_var_3_4 shape=EffectBind(Let(Let(Other))) bindingType=Any
				__local_var_3_4 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					_dollar___unused_3_5 := gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Value{})
					_ = _dollar___unused_3_5
					// TAST (Let): __local_var_4_6 shape=App(Var) bindingType=Any
					__local_var_4_6 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m2"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("nested")
					})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("zArrayA")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("ignore")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fa")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("arrayIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("a")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("nested")
					})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("zArrayA")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("ignore")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fa")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("fIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("arrayIgnore")
					})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str("a")
					})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_2())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})))})).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})))}})).UnsafePtr)))).IntVal) != (0)))
					_ = __local_var_4_6
					// TAST (Let): __local_var_5_7 shape=EffectBind(Let(Let(Other))) bindingType=Any
					__local_var_5_7 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						_dollar___unused_5_8 := gopurs_runtime.Apply(__local_var_4_6, gopurs_runtime.Value{})
						_ = _dollar___unused_5_8
						// TAST (Let): __local_var_6_9 shape=App(Var) bindingType=Any
						__local_var_6_9 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m3"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("nested")
						})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("zArrayA")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("ignore")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fa")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("arrayIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("a")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("nested")
						})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("zArrayA")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("ignore")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fa")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("fIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("arrayIgnore")
						})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str("a")
						})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_3())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, Get_Main_m3())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m3()})).UnsafePtr)))).IntVal) != (0)))
						_ = __local_var_6_9
						// TAST (Let): __local_var_7_10 shape=EffectBind(Let(Let(Other))) bindingType=Any
						__local_var_7_10 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							_dollar___unused_7_11 := gopurs_runtime.Apply(__local_var_6_9, gopurs_runtime.Value{})
							_ = _dollar___unused_7_11
							// TAST (Let): __local_var_8_12 shape=App(Var) bindingType=Any
							__local_var_8_12 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m4"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("nested")
							})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("zArrayA")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("ignore")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fa")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("arrayIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("a")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("nested")
							})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("zArrayA")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("ignore")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fa")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("fIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("arrayIgnore")
							})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str("a")
							})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_4())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
								orig := func() struct {
									a           string
									arrayIgnore []int64
									fIgnore     []int64
									fa          []string
									ignore      int64
									zArrayA     []string
									/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
							}()})))})).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
								orig := func() struct {
									a           string
									arrayIgnore []int64
									fIgnore     []int64
									fa          []string
									ignore      int64
									zArrayA     []string
									/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
							}()})))}})).UnsafePtr)))).IntVal) != (0)))
							_ = __local_var_8_12
							// TAST (Let): __local_var_9_13 shape=EffectBind(Let(Let(Other))) bindingType=Any
							__local_var_9_13 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								_dollar___unused_9_14 := gopurs_runtime.Apply(__local_var_8_12, gopurs_runtime.Value{})
								_ = _dollar___unused_9_14
								// TAST (Let): __local_var_10_15 shape=App(Var) bindingType=Any
								__local_var_10_15 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m5"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("nested")
								})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("zArrayA")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("ignore")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fa")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("arrayIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("a")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("nested")
								})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("zArrayA")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("ignore")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fa")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("fIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("arrayIgnore")
								})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str("a")
								})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_5())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, Get_Main_m5())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m5()})).UnsafePtr)))).IntVal) != (0)))
								_ = __local_var_10_15
								// TAST (Let): __local_var_11_16 shape=EffectBind(Let(Let(Other))) bindingType=Any
								__local_var_11_16 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									_dollar___unused_11_17 := gopurs_runtime.Apply(__local_var_10_15, gopurs_runtime.Value{})
									_ = _dollar___unused_11_17
									// TAST (Let): __local_var_12_18 shape=App(Var) bindingType=Any
									__local_var_12_18 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m6"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("nested")
									})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("zArrayA")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("ignore")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fa")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("arrayIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("a")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("nested")
									})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("zArrayA")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("ignore")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fa")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("fIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("arrayIgnore")
									})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str("a")
									})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_6())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, Get_Main_m6())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m6()})).UnsafePtr)))).IntVal) != (0)))
									_ = __local_var_12_18
									// TAST (Let): __local_var_13_19 shape=EffectBind(Let(Let(Other))) bindingType=Any
									__local_var_13_19 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										_dollar___unused_13_20 := gopurs_runtime.Apply(__local_var_12_18, gopurs_runtime.Value{})
										_ = _dollar___unused_13_20
										// TAST (Let): __local_var_14_21 shape=App(Var) bindingType=Any
										__local_var_14_21 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("traverse - m7"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("nested")
										})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("zArrayA")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("ignore")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fa")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("arrayIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("a")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("nested")
										})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("zArrayA")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("ignore")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fa")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("fIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("arrayIgnore")
										})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str("a")
										})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main___cached_dict_7())), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1, Get_Main_m7())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m7()})).UnsafePtr)))).IntVal) != (0)))
										_ = __local_var_14_21
										// TAST (Let): __local_var_15_22 shape=EffectBind(Let(Let(Other))) bindingType=Any
										__local_var_15_22 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											_dollar___unused_15_23 := gopurs_runtime.Apply(__local_var_14_21, gopurs_runtime.Value{})
											_ = _dollar___unused_15_23
											// TAST (Let): __local_var_16_24 shape=App(Var) bindingType=Any
											__local_var_16_24 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m0"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("nested")
											})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("zArrayA")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("ignore")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fa")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("arrayIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("a")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("nested")
											})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("zArrayA")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("ignore")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fa")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("fIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("arrayIgnore")
											})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str("a")
											})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3974551048_67812977(nil))})).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}})).UnsafePtr)))).IntVal) != (0)))
											_ = __local_var_16_24
											// TAST (Let): __local_var_17_25 shape=EffectBind(Let(Let(Other))) bindingType=Any
											__local_var_17_25 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												_dollar___unused_17_26 := gopurs_runtime.Apply(__local_var_16_24, gopurs_runtime.Value{})
												_ = _dollar___unused_17_26
												// TAST (Let): __local_var_18_27 shape=App(Var) bindingType=Any
												__local_var_18_27 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m1"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("nested")
												})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("zArrayA")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("ignore")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fa")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("arrayIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("a")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("nested")
												})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("zArrayA")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("ignore")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fa")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("fIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("arrayIgnore")
												})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str("a")
												})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m1_prime_())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m1()})).UnsafePtr)))).IntVal) != (0)))
												_ = __local_var_18_27
												// TAST (Let): __local_var_19_28 shape=EffectBind(Let(Let(Other))) bindingType=Any
												__local_var_19_28 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													_dollar___unused_19_29 := gopurs_runtime.Apply(__local_var_18_27, gopurs_runtime.Value{})
													_ = _dollar___unused_19_29
													// TAST (Let): __local_var_20_30 shape=App(Var) bindingType=Any
													__local_var_20_30 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m2"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("nested")
													})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("zArrayA")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("ignore")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fa")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("arrayIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("a")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("nested")
													})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("zArrayA")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("ignore")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fa")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("fIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("arrayIgnore")
													})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str("a")
													})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_2857385098_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, []string]{1, int64(0)})))})).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0)})))}})).UnsafePtr)))).IntVal) != (0)))
													_ = __local_var_20_30
													// TAST (Let): __local_var_21_31 shape=EffectBind(Let(Let(Other))) bindingType=Any
													__local_var_21_31 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
														_dollar___unused_21_32 := gopurs_runtime.Apply(__local_var_20_30, gopurs_runtime.Value{})
														_ = _dollar___unused_21_32
														// TAST (Let): __local_var_22_33 shape=App(Var) bindingType=Any
														__local_var_22_33 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m3"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("nested")
														})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("zArrayA")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("ignore")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fa")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("arrayIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("a")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("nested")
														})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("zArrayA")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("ignore")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fa")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("fIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("arrayIgnore")
														})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str("a")
														})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m3_prime_())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m3()})).UnsafePtr)))).IntVal) != (0)))
														_ = __local_var_22_33
														// TAST (Let): __local_var_23_34 shape=EffectBind(Let(Let(Other))) bindingType=Any
														__local_var_23_34 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
															_dollar___unused_23_35 := gopurs_runtime.Apply(__local_var_22_33, gopurs_runtime.Value{})
															_ = _dollar___unused_23_35
															// TAST (Let): __local_var_24_36 shape=App(Var) bindingType=Any
															__local_var_24_36 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m4"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("nested")
															})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("zArrayA")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("ignore")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fa")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("arrayIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("a")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("nested")
															})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("zArrayA")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("ignore")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fa")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("fIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("arrayIgnore")
															})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str("a")
															})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_2381891596_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, []string]{1, func() gopurs_runtime.Value {
																orig := func() struct {
																	a           []string
																	arrayIgnore []int64
																	fIgnore     []int64
																	fa          [][]string
																	ignore      int64
																	zArrayA     [][]string
																	/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
																		/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
															}()})))})).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
																orig := func() struct {
																	a           string
																	arrayIgnore []int64
																	fIgnore     []int64
																	fa          []string
																	ignore      int64
																	zArrayA     []string
																	/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
																		/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
															}()})))}})).UnsafePtr)))).IntVal) != (0)))
															_ = __local_var_24_36
															// TAST (Let): __local_var_25_37 shape=EffectBind(Let(Let(Other))) bindingType=Any
															__local_var_25_37 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																_dollar___unused_25_38 := gopurs_runtime.Apply(__local_var_24_36, gopurs_runtime.Value{})
																_ = _dollar___unused_25_38
																// TAST (Let): __local_var_26_39 shape=App(Var) bindingType=Any
																__local_var_26_39 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m5"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("nested")
																})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("zArrayA")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("ignore")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fa")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("arrayIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("a")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("nested")
																})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("zArrayA")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("ignore")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fa")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("fIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("arrayIgnore")
																})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str("a")
																})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m5_prime_())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m5()})).UnsafePtr)))).IntVal) != (0)))
																_ = __local_var_26_39
																// TAST (Let): __local_var_27_40 shape=EffectBind(Let(Let(Other))) bindingType=Any
																__local_var_27_40 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	_dollar___unused_27_41 := gopurs_runtime.Apply(__local_var_26_39, gopurs_runtime.Value{})
																	_ = _dollar___unused_27_41
																	// TAST (Let): __local_var_28_42 shape=App(Var) bindingType=Any
																	__local_var_28_42 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m6"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("nested")
																	})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("zArrayA")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("ignore")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fa")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("arrayIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("a")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("nested")
																	})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("zArrayA")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("ignore")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fa")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("fIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("arrayIgnore")
																	})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str("a")
																	})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m6_prime_())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m6()})).UnsafePtr)))).IntVal) != (0)))
																	_ = __local_var_28_42
																	// TAST (Let): __local_var_29_43 shape=EffectBind(Let(Let(Other))) bindingType=Any
																	__local_var_29_43 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																		_dollar___unused_29_44 := gopurs_runtime.Apply(__local_var_28_42, gopurs_runtime.Value{})
																		_ = _dollar___unused_29_44
																		// TAST (Let): __local_var_30_45 shape=App(Var) bindingType=Any
																		__local_var_30_45 := gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("sequence - m7"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(Call_Main_eqM(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}, Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("nested")
																		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("zArrayA")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("ignore")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fa")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("arrayIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("a")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("nested")
																		})), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("zArrayA")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("ignore")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fa")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("fIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("arrayIgnore")
																		})), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Str("a")
																		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, Get_Main_m7_prime_())).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m7()})).UnsafePtr)))).IntVal) != (0)))
																		_ = __local_var_30_45
																		// TAST (Let): __local_var_31_46 shape=EffectBind(App(Var)) bindingType=Any
																		__local_var_31_46 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																			_dollar___unused_31_47 := gopurs_runtime.Apply(__local_var_30_45, gopurs_runtime.Value{})
																			_ = _dollar___unused_31_47
																			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
																		})
																		_ = __local_var_31_46
																		return gopurs_runtime.Apply(__local_var_31_46, gopurs_runtime.Value{})
																	})
																	_ = __local_var_29_43
																	return gopurs_runtime.Apply(__local_var_29_43, gopurs_runtime.Value{})
																})
																_ = __local_var_27_40
																return gopurs_runtime.Apply(__local_var_27_40, gopurs_runtime.Value{})
															})
															_ = __local_var_25_37
															return gopurs_runtime.Apply(__local_var_25_37, gopurs_runtime.Value{})
														})
														_ = __local_var_23_34
														return gopurs_runtime.Apply(__local_var_23_34, gopurs_runtime.Value{})
													})
													_ = __local_var_21_31
													return gopurs_runtime.Apply(__local_var_21_31, gopurs_runtime.Value{})
												})
												_ = __local_var_19_28
												return gopurs_runtime.Apply(__local_var_19_28, gopurs_runtime.Value{})
											})
											_ = __local_var_17_25
											return gopurs_runtime.Apply(__local_var_17_25, gopurs_runtime.Value{})
										})
										_ = __local_var_15_22
										return gopurs_runtime.Apply(__local_var_15_22, gopurs_runtime.Value{})
									})
									_ = __local_var_13_19
									return gopurs_runtime.Apply(__local_var_13_19, gopurs_runtime.Value{})
								})
								_ = __local_var_11_16
								return gopurs_runtime.Apply(__local_var_11_16, gopurs_runtime.Value{})
							})
							_ = __local_var_9_13
							return gopurs_runtime.Apply(__local_var_9_13, gopurs_runtime.Value{})
						})
						_ = __local_var_7_10
						return gopurs_runtime.Apply(__local_var_7_10, gopurs_runtime.Value{})
					})
					_ = __local_var_5_7
					return gopurs_runtime.Apply(__local_var_5_7, gopurs_runtime.Value{})
				})
				_ = __local_var_3_4
				return gopurs_runtime.Apply(__local_var_3_4, gopurs_runtime.Value{})
			})
			_ = __local_var_1_1
			return __local_var_1_1
		}()
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))})))}
	})
	return cache_Main___cached_dict_7
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
		/*gopurs:record-v1:1:006e00650073007400650064*/
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
		/*gopurs:record-v1:1:006e00650073007400650064*/
	}
}

type Constructor_Main_M7[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_functorM(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t14 gopurs_runtime.Value
		{
			if m_2.Type == 9 && m_2.IntVal == 3852365315 {
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
				goto end_branch_14
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 769986722 {
				arrayIntrinsic_arg_0 := f_1
				_ = arrayIntrinsic_arg_0
				arrayIntrinsic_arg_1 := gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)
				_ = arrayIntrinsic_arg_1
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
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
				goto end_branch_14
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2727978561 {
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0}))}
				goto end_branch_14
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 1830062304 {
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
				goto end_branch_14
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 3190619783 {
				arrayIntrinsic_arg_2 := f_1
				_ = arrayIntrinsic_arg_2
				arrayIntrinsic_arg_3 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_3
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a")), "fa", gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap5 := arrayIntrinsic_arg_3
					_ = arr_val_arrayMap5
					arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
					_ = arr_go_arrayMap5
					res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
					_ = res_go_arrayMap5
					for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
						res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(arrayIntrinsic_arg_2, v_arrayMap5)
					}
					return gopurs_runtime.Array(res_go_arrayMap5)
				}()).UnsafePtr))))}))}
				goto end_branch_14
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 108241190 {
				arrayIntrinsic_arg_4 := f_1
				_ = arrayIntrinsic_arg_4
				arrayIntrinsic_arg_5 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_5
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, func(record struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						fIgnore     gopurs_runtime.Value
						fa          gopurs_runtime.Value
						ignore      int64
						zArrayA     []gopurs_runtime.Value
						/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
						return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
							arr := orig.arrayIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
					}()}
				}(func() struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						fIgnore     gopurs_runtime.Value
						fa          gopurs_runtime.Value
						ignore      int64
						zArrayA     []gopurs_runtime.Value
						/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
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
							/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
						}
						/*gopurs:record-v1:1:006e00650073007400650064*/
					}
					clone.nested = func() struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						fIgnore     gopurs_runtime.Value
						fa          gopurs_runtime.Value
						ignore      int64
						zArrayA     []gopurs_runtime.Value
						/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
					} {
						orig := gopurs_runtime.RecordUpdate3((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a")), "fa", gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap6 := arrayIntrinsic_arg_5
							_ = arr_val_arrayMap6
							arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
							_ = arr_go_arrayMap6
							res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
							_ = res_go_arrayMap6
							for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
								res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(arrayIntrinsic_arg_4, v_arrayMap6)
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
							/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
				}())}))}
				goto end_branch_14
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2066233029 {
				arrayIntrinsic_arg_6 := f_1
				_ = arrayIntrinsic_arg_6
				arrayIntrinsic_arg_7 := gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3)
				_ = arrayIntrinsic_arg_7
				arrayIntrinsic_arg_8 := f_1
				_ = arrayIntrinsic_arg_8
				arrayIntrinsic_arg_9 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_9
				arrayIntrinsic_arg_10 := f_1
				_ = arrayIntrinsic_arg_10
				arrayIntrinsic_arg_11 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_11
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V2, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := arrayIntrinsic_arg_7
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(arrayIntrinsic_arg_6, v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr)), gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a")), "fa", gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap5 := arrayIntrinsic_arg_9
					_ = arr_val_arrayMap5
					arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
					_ = arr_go_arrayMap5
					res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
					_ = res_go_arrayMap5
					for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
						res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(arrayIntrinsic_arg_8, v_arrayMap5)
					}
					return gopurs_runtime.Array(res_go_arrayMap5)
				}()).UnsafePtr)))), func(record struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						fIgnore     gopurs_runtime.Value
						fa          gopurs_runtime.Value
						ignore      int64
						zArrayA     []gopurs_runtime.Value
						/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
						return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
							arr := orig.arrayIgnore
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
					}()}
				}(func() struct {
					nested struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						fIgnore     gopurs_runtime.Value
						fa          gopurs_runtime.Value
						ignore      int64
						zArrayA     []gopurs_runtime.Value
						/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
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
							/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
						}
						/*gopurs:record-v1:1:006e00650073007400650064*/
					}
					clone.nested = func() struct {
						a           gopurs_runtime.Value
						arrayIgnore []int64
						fIgnore     gopurs_runtime.Value
						fa          gopurs_runtime.Value
						ignore      int64
						zArrayA     []gopurs_runtime.Value
						/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
					} {
						orig := gopurs_runtime.RecordUpdate3((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a")), "fa", gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
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
						}()).UnsafePtr))))
						_ = orig
						clone := struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							fIgnore     gopurs_runtime.Value
							fa          gopurs_runtime.Value
							ignore      int64
							zArrayA     []gopurs_runtime.Value
							/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
				}())}))}
				goto end_branch_14
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 1168316772 {
				__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), gopurs_runtime.Apply(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
					arrayIntrinsic_arg_12 := f_1
					_ = arrayIntrinsic_arg_12
					arrayIntrinsic_arg_13 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA")).UnsafePtr)))
					_ = arrayIntrinsic_arg_13
					return gopurs_runtime.RecordUpdate1(v1_3, "nested", func() gopurs_runtime.Value {
						orig := func() struct {
							a           gopurs_runtime.Value
							arrayIgnore []int64
							fIgnore     gopurs_runtime.Value
							fa          gopurs_runtime.Value
							ignore      int64
							zArrayA     []gopurs_runtime.Value
							/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
						} {
							orig := gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_3, "nested"), "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a")), "fa", gopurs_runtime.Apply2(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0)), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa")), "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
								arr_val_arrayMap9 := arrayIntrinsic_arg_13
								_ = arr_val_arrayMap9
								arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
								_ = arr_go_arrayMap9
								res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
								_ = res_go_arrayMap9
								for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
									res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(arrayIntrinsic_arg_12, v_arrayMap9)
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
								/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
				})), (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
				goto end_branch_14
			} else {

			}
		}
		{
			__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_14:
		return __t14
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
					__t2 = gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply(f_4, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1)))
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
					__t2 = gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0)
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 3190619783 {
					__t2 = gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "zArrayA")).UnsafePtr))))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 108241190 {
					__t2 = gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr))))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 2066233029 {
					__t2 = gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply(f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V3)), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V4), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "zArrayA")).UnsafePtr)))), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr))))))))))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 1168316772 {
					__t2 = gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Apply2(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "a")), gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(Semigroup0_3_1), gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "zArrayA")).UnsafePtr))))))
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
		var __t24 gopurs_runtime.Value
		{
			if m_3.Type == 9 && m_3.IntVal == 3852365315 {
				__t24 = z_2
				goto end_branch_24
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 769986722 {
				arrayIntrinsic_arg_3 := f_1
				_ = arrayIntrinsic_arg_3
				arrayIntrinsic_arg_4 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				_ = arrayIntrinsic_arg_4
				arrayIntrinsic_arg_5 := gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)
				_ = arrayIntrinsic_arg_5
				__t24 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := arrayIntrinsic_arg_5
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := arrayIntrinsic_arg_4
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(arrayIntrinsic_arg_3, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_24
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2727978561 {
				__t24 = z_2
				goto end_branch_24
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1830062304 {
				__t24 = gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, z_2, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				goto end_branch_24
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3190619783 {
				arrayIntrinsic_arg_6 := f_1
				_ = arrayIntrinsic_arg_6
				arrayIntrinsic_arg_7 := gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "fa"))
				_ = arrayIntrinsic_arg_7
				arrayIntrinsic_arg_8 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_8
				__t24 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := arrayIntrinsic_arg_8
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := arrayIntrinsic_arg_7
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(arrayIntrinsic_arg_6, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_24
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 108241190 {
				arrayIntrinsic_arg_9 := f_1
				_ = arrayIntrinsic_arg_9
				arrayIntrinsic_arg_10 := gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "fa"))
				_ = arrayIntrinsic_arg_10
				arrayIntrinsic_arg_11 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_11
				__t24 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := arrayIntrinsic_arg_11
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := arrayIntrinsic_arg_10
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(arrayIntrinsic_arg_9, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_24
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2066233029 {
				arrayIntrinsic_arg_12 := f_1
				_ = arrayIntrinsic_arg_12
				arrayIntrinsic_arg_13 := f_1
				_ = arrayIntrinsic_arg_13
				arrayIntrinsic_arg_14 := f_1
				_ = arrayIntrinsic_arg_14
				arrayIntrinsic_arg_15 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)
				_ = arrayIntrinsic_arg_15
				arrayIntrinsic_arg_16 := gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)
				_ = arrayIntrinsic_arg_16
				arrayIntrinsic_arg_17 := gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, func() gopurs_runtime.Value {
					arr_val_foldlArray10 := arrayIntrinsic_arg_16
					_ = arr_val_foldlArray10
					res_go_foldlArray10 := arrayIntrinsic_arg_15
					_ = res_go_foldlArray10
					arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
					_ = arr_go_foldlArray10
					for _, v_foldlArray10 := range *arr_go_foldlArray10 {
						res_go_foldlArray10 = gopurs_runtime.Apply2(arrayIntrinsic_arg_14, res_go_foldlArray10, v_foldlArray10)
					}
					return res_go_foldlArray10
				}(), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V4), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "fa"))
				_ = arrayIntrinsic_arg_17
				arrayIntrinsic_arg_18 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_18
				arrayIntrinsic_arg_19 := gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply2(f_1, func() gopurs_runtime.Value {
					arr_val_foldlArray6 := arrayIntrinsic_arg_18
					_ = arr_val_foldlArray6
					res_go_foldlArray6 := arrayIntrinsic_arg_17
					_ = res_go_foldlArray6
					arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
					_ = arr_go_foldlArray6
					for _, v_foldlArray6 := range *arr_go_foldlArray6 {
						res_go_foldlArray6 = gopurs_runtime.Apply2(arrayIntrinsic_arg_13, res_go_foldlArray6, v_foldlArray6)
					}
					return res_go_foldlArray6
				}(), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "fa"))
				_ = arrayIntrinsic_arg_19
				arrayIntrinsic_arg_20 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))
				_ = arrayIntrinsic_arg_20
				__t24 = func() gopurs_runtime.Value {
					arr_val_foldlArray3 := arrayIntrinsic_arg_20
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := arrayIntrinsic_arg_19
					_ = res_go_foldlArray3
					arr_go_foldlArray3 := (*[]gopurs_runtime.Value)(arr_val_foldlArray3.UnsafePtr)
					_ = arr_go_foldlArray3
					for _, v_foldlArray3 := range *arr_go_foldlArray3 {
						res_go_foldlArray3 = gopurs_runtime.Apply2(arrayIntrinsic_arg_12, res_go_foldlArray3, v_foldlArray3)
					}
					return res_go_foldlArray3
				}()
				goto end_branch_24
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1168316772 {
				__t24 = gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Apply(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Func2(func(v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
					arrayIntrinsic_arg_21 := f_1
					_ = arrayIntrinsic_arg_21
					arrayIntrinsic_arg_22 := gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply2(f_1, v1_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "fa"))
					_ = arrayIntrinsic_arg_22
					arrayIntrinsic_arg_23 := gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "zArrayA")).UnsafePtr)))
					_ = arrayIntrinsic_arg_23
					return func() gopurs_runtime.Value {
						arr_val_foldlArray6 := arrayIntrinsic_arg_23
						_ = arr_val_foldlArray6
						res_go_foldlArray6 := arrayIntrinsic_arg_22
						_ = res_go_foldlArray6
						arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
						_ = arr_go_foldlArray6
						for _, v_foldlArray6 := range *arr_go_foldlArray6 {
							res_go_foldlArray6 = gopurs_runtime.Apply2(arrayIntrinsic_arg_21, res_go_foldlArray6, v_foldlArray6)
						}
						return res_go_foldlArray6
					}()
				})), z_2, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				goto end_branch_24
			} else {

			}
		}
		{
			__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_24:
		return __t24
	}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t26 gopurs_runtime.Value
		{
			if m_3.Type == 9 && m_3.IntVal == 3852365315 {
				__t26 = z_2
				goto end_branch_26
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 769986722 {
				__t26 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)))
				goto end_branch_26
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2727978561 {
				__t26 = z_2
				goto end_branch_26
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1830062304 {
				__t26 = gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, z_2, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				goto end_branch_26
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3190619783 {
				__t26 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "fa")))
				goto end_branch_26
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 108241190 {
				__t26 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "a"), gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "fa")))
				goto end_branch_26
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2066233029 {
				__t26 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "a"), gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "fa"))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "fa"))), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V4), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)))
				goto end_branch_26
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1168316772 {
				// TAST (Let): __local_var_4_25 shape=App(Var) bindingType=Any
				__local_var_4_25 := gopurs_runtime.Apply(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Func2(func(v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "a"), gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, v2_5, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "zArrayA")).UnsafePtr)))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "fa")))
				}))
				_ = __local_var_4_25
				__t26 = gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0)), gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(__local_var_4_25, a_6, b_5)
				}), z_2, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
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
			return gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Main_traversableM(dictTraversable_0))), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_5
			}), v_4)
		}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_4_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope10)])
			Apply0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_4_2
			// TAST (Let): Functor0_5_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope10)])
			Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_5_3
			return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t18 gopurs_runtime.Value
				{
					if m_7.Type == 9 && m_7.IntVal == 3852365315 {
						__t18 = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)})
						goto end_branch_18
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 769986722 {
						// TAST (Let): Apply0_8_4 shape=App(Other) bindingType=Any
						Apply0_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_8_4
						__t18 = gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_5_3), gopurs_runtime.Func2(func(v2_8 gopurs_runtime.Value, v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_8, (*(*[]gopurs_runtime.Value)((v3_9).UnsafePtr))}))}
						}), gopurs_runtime.Apply(f_6, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0)), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_8_4)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_4, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1)))
						goto end_branch_18
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 2727978561 {
						__t18 = gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0}))})
						goto end_branch_18
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 1830062304 {
						__t18 = gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_5_3), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{1, v1_8}))}
						}), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0))
						goto end_branch_18
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 3190619783 {
						// TAST (Let): __local_var_8_5 shape=Other bindingType=Any
						__local_var_8_5 := (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
						_ = __local_var_8_5
						// TAST (Let): Apply0_9_6 shape=App(Other) bindingType=Any
						Apply0_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_9_6
						__t18 = gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_5_3), gopurs_runtime.Func3(func(v1_9 gopurs_runtime.Value, v2_10 gopurs_runtime.Value, v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordUpdate3(__local_var_8_5, "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v3_11).UnsafePtr))))}))}
						}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_8_5, "a"))), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_8_5, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_9_6)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_6, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_8_5, "zArrayA")).UnsafePtr)))))
						goto end_branch_18
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 108241190 {
						// TAST (Let): __local_var_8_7 shape=Other bindingType=Any
						__local_var_8_7 := (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
						_ = __local_var_8_7
						// TAST (Let): Apply0_9_8 shape=App(Other) bindingType=Any
						Apply0_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_9_8
						__t18 = gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_5_3), gopurs_runtime.Func3(func(v1_9 gopurs_runtime.Value, v2_10 gopurs_runtime.Value, v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, func(record struct {
								nested struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
									/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
									return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
										arr := orig.arrayIgnore
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Int(v)
										}
										return gopurs_runtime.Array(boxed)
									}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
								}()}
							}(func() struct {
								nested struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
									/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
								}
								/*gopurs:record-v1:1:006e00650073007400650064*/
							} {
								originalRecord := __local_var_8_7
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
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
									}
									/*gopurs:record-v1:1:006e00650073007400650064*/
								}
								clone.nested = func() struct {
									a           gopurs_runtime.Value
									arrayIgnore []int64
									fIgnore     gopurs_runtime.Value
									fa          gopurs_runtime.Value
									ignore      int64
									zArrayA     []gopurs_runtime.Value
									/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
								} {
									orig := gopurs_runtime.RecordUpdate3(__local_var_8_7.nested, "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v3_11).UnsafePtr))))
									_ = orig
									clone := struct {
										a           gopurs_runtime.Value
										arrayIgnore []int64
										fIgnore     gopurs_runtime.Value
										fa          gopurs_runtime.Value
										ignore      int64
										zArrayA     []gopurs_runtime.Value
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
							}())}))}
						}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_8_7.nested, "a"))), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_8_7.nested, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_9_8)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_8, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_8_7.nested, "zArrayA")).UnsafePtr)))))
						goto end_branch_18
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 2066233029 {
						// TAST (Let): __local_var_8_9 shape=Other bindingType=Any
						__local_var_8_9 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0
						_ = __local_var_8_9
						// TAST (Let): __local_var_9_10 shape=Other bindingType=Any
						__local_var_9_10 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V2
						_ = __local_var_9_10
						// TAST (Let): __local_var_10_11 shape=Other bindingType=Any
						__local_var_10_11 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V5
						_ = __local_var_10_11
						// TAST (Let): __local_var_11_12 shape=Other bindingType=Any
						__local_var_11_12 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V6
						_ = __local_var_11_12
						// TAST (Let): __local_var_12_13 shape=Other bindingType=Any
						__local_var_12_13 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V7
						_ = __local_var_12_13
						// TAST (Let): Apply0_13_14 shape=App(Other) bindingType=Any
						Apply0_13_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_13_14
						// TAST (Let): Apply0_13_15 shape=App(Other) bindingType=Any
						Apply0_13_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_13_15
						// TAST (Let): Apply0_13_16 shape=App(Other) bindingType=Any
						Apply0_13_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_13_16
						__t18 = gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_5_3), gopurs_runtime.Func5(func(v8_13 gopurs_runtime.Value, v9_14 gopurs_runtime.Value, v10_15 gopurs_runtime.Value, v11_16 gopurs_runtime.Value, v12_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func4(func(v13_18 gopurs_runtime.Value, v14_19 gopurs_runtime.Value, v15_20 gopurs_runtime.Value, v16_21 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_8_9, v8_13, __local_var_9_10, (*(*[]gopurs_runtime.Value)((v9_14).UnsafePtr)), v10_15, __local_var_10_11, gopurs_runtime.RecordUpdate3(__local_var_11_12, "a", v11_16, "fa", v12_17, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v13_18).UnsafePtr)))), func(record struct {
									nested struct {
										a           gopurs_runtime.Value
										arrayIgnore []int64
										fIgnore     gopurs_runtime.Value
										fa          gopurs_runtime.Value
										ignore      int64
										zArrayA     []gopurs_runtime.Value
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
										return gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{orig.a, func() gopurs_runtime.Value {
											arr := orig.arrayIgnore
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}(), orig.fIgnore, orig.fa, gopurs_runtime.Int(orig.ignore), gopurs_runtime.Array(orig.zArrayA)})
									}()}
								}(func() struct {
									nested struct {
										a           gopurs_runtime.Value
										arrayIgnore []int64
										fIgnore     gopurs_runtime.Value
										fa          gopurs_runtime.Value
										ignore      int64
										zArrayA     []gopurs_runtime.Value
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
									}
									/*gopurs:record-v1:1:006e00650073007400650064*/
								} {
									originalRecord := __local_var_12_13
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
											/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
										}
										/*gopurs:record-v1:1:006e00650073007400650064*/
									}
									clone.nested = func() struct {
										a           gopurs_runtime.Value
										arrayIgnore []int64
										fIgnore     gopurs_runtime.Value
										fa          gopurs_runtime.Value
										ignore      int64
										zArrayA     []gopurs_runtime.Value
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
									} {
										orig := gopurs_runtime.RecordUpdate3(__local_var_12_13.nested, "a", v14_19, "fa", v15_20, "zArrayA", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v16_21).UnsafePtr))))
										_ = orig
										clone := struct {
											a           gopurs_runtime.Value
											arrayIgnore []int64
											fIgnore     gopurs_runtime.Value
											fa          gopurs_runtime.Value
											ignore      int64
											zArrayA     []gopurs_runtime.Value
											/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
								}())}))}
							})
						}), gopurs_runtime.Apply(f_6, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V1)), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_13_14)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_13_14, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V3))), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V4)), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_11_12, "a"))), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_11_12, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_13_15)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_13_15, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_11_12, "zArrayA")).UnsafePtr))))), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_12_13.nested, "a"))), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_12_13.nested, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_13_16)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_13_16, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_12_13.nested, "zArrayA")).UnsafePtr)))))
						goto end_branch_18
					} else {

					}
				}
				{
					if m_7.Type == 9 && m_7.IntVal == 1168316772 {
						__t18 = gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_5_3), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value]{1, v1_8}))}
						}), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Apply2(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): Apply0_9_17 shape=App(Other) bindingType=Any
							Apply0_9_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
							_ = Apply0_9_17
							return gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Control_Apply_apply(Apply0_4_2), gopurs_runtime.Apply2(Call_Data_Functor_go__map(Functor0_5_3), gopurs_runtime.Func3(func(v2_9 gopurs_runtime.Value, v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.RecordUpdate1(v1_8, "nested", func() gopurs_runtime.Value {
									orig := func() struct {
										a           gopurs_runtime.Value
										arrayIgnore []int64
										fIgnore     gopurs_runtime.Value
										fa          gopurs_runtime.Value
										ignore      int64
										zArrayA     []gopurs_runtime.Value
										/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
											/*gopurs:record-v1:6:0061,0061007200720061007900490067006e006f00720065,006600490067006e006f00720065,00660061,00690067006e006f00720065,007a004100720072006100790041*/
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
							}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "a"))), gopurs_runtime.Apply3(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0)), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_9_17)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_17, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3)), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "zArrayA")).UnsafePtr)))))
						})), (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0))
						goto end_branch_18
					} else {

					}
				}
				{
					__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_18:
				return __t18
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
				__t1 = (y_6.Type == 9 && y_6.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal) != (0))
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 3190619783 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 3190619783) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0, "zArrayA")).UnsafePtr)))).IntVal) != (0)))
				goto end_branch_1
			} else {

			}
		}
		{
			if x_5.Type == 9 && x_5.IntVal == 108241190 {
				__t1 = (y_6.Type == 9 && y_6.IntVal == 108241190) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0.nested, "zArrayA")).UnsafePtr)))).IntVal) != (0)))
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V5, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V6, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V6, "zArrayA")).UnsafePtr)))).IntVal) != (0)))) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "arrayIgnore").UnsafePtr)
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
				}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "ignore").IntVal))) && ((gopurs_runtime.Apply2(eqArray5_4_0.V0, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V7.nested, "zArrayA")).UnsafePtr)))).IntVal) != (0))))
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = (x_5.Type == 9 && x_5.IntVal == 1168316772) && ((y_6.Type == 9 && y_6.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0)), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq2_2))}, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal) != (0)))
		}
	end_branch_1:
		return gopurs_runtime.Bool(__t1)
	})}))}
}

func Call_Main_traverseStr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply2(Call_Data_Traversable_traverse(dictTraversable_0), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()).V1)
}

func Call_Main_sequenceStr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply(Call_Data_Traversable_sequence(dictTraversable_0), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))})
}

func Rebox_Main_1019686859_2554376626(in *Constructor_Main_M3[gopurs_runtime.Value, []string]) *Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
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
	return (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
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
	return (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2785108781_2554376626(in *Constructor_Main_M3[gopurs_runtime.Value, string]) *Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2857385098_1521903347(in *Constructor_Main_M2[gopurs_runtime.Value, []string]) *Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_358349645_2067946548(in *Constructor_Main_M5[gopurs_runtime.Value, []string]) *Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3721994703_961717174(in *Constructor_Main_M7[gopurs_runtime.Value, []string]) *Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3741832558_67812977(in *Constructor_Main_M0[gopurs_runtime.Value, string]) *Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
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

func Rebox_Main_3790796878_1939691112(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3974551048_67812977(in *Constructor_Main_M0[gopurs_runtime.Value, []string]) *Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
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
	return (*Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_82800937_961717174(in *Constructor_Main_M7[gopurs_runtime.Value, string]) *Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
