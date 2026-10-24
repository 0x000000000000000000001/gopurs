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
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1}))}
			})
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

var cache_Main_m2 gopurs_runtime.Value
var once_Main_m2 sync.Once

func Get_Main_m2() gopurs_runtime.Value {
	once_Main_m2.Do(func() {
		cache_Main_m2 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})})))}
	})
	return cache_Main_m2
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

var cache_Main_m0 gopurs_runtime.Value
var once_Main_m0 sync.Once

func Get_Main_m0() gopurs_runtime.Value {
	once_Main_m0.Do(func() {
		cache_Main_m0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}
	})
	return cache_Main_m0
}

var cache_Main_foldrStr gopurs_runtime.Value
var once_Main_foldrStr sync.Once

func Get_Main_foldrStr() gopurs_runtime.Value {
	once_Main_foldrStr.Do(func() {
		cache_Main_foldrStr = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldrStr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Main_foldrStr
}

var cache_Main_foldlStr gopurs_runtime.Value
var once_Main_foldlStr sync.Once

func Get_Main_foldlStr() gopurs_runtime.Value {
	once_Main_foldlStr.Do(func() {
		cache_Main_foldlStr = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldlStr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Main_foldlStr
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

var cache_Main_foldMapStr gopurs_runtime.Value
var once_Main_foldMapStr sync.Once

func Get_Main_foldMapStr() gopurs_runtime.Value {
	once_Main_foldMapStr.Do(func() {
		cache_Main_foldMapStr = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldMapStr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Main_foldMapStr
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M0", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_0())), gopurs_runtime.Func2(func(acc_0 gopurs_runtime.Value, next_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_0.StrVal(), "<"), next_1.StrVal()))
			}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal(), "Start"}))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M1", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_1())), gopurs_runtime.Func2(func(acc_2 gopurs_runtime.Value, next_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_2.StrVal(), "<"), next_3.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal(), "Start<a<b<c"})), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M2", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_2())), gopurs_runtime.Func2(func(acc_3 gopurs_runtime.Value, next_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_3.StrVal(), "<"), next_4.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal(), "Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M3", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_3())), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, next_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_4.StrVal(), "<"), next_5.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal(), "Start<a<b<c"})), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M4", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_4())), gopurs_runtime.Func2(func(acc_5 gopurs_runtime.Value, next_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_5.StrVal(), "<"), next_6.StrVal()))
			}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
			}()})))}).StrVal(), "Start<a<b<c"})), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M5", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_5())), gopurs_runtime.Func2(func(acc_6 gopurs_runtime.Value, next_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_6.StrVal(), "<"), next_7.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal(), "Start<a<b<c"})), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M6", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_6())), gopurs_runtime.Func2(func(acc_7 gopurs_runtime.Value, next_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_7.StrVal(), "<"), next_8.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal(), "Start<a<b<c<a<b<c<a<b<c"})), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldl - M7", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_7())), gopurs_runtime.Func2(func(acc_8 gopurs_runtime.Value, next_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_8.StrVal(), "<"), next_9.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal(), "Start<a<b<c"})), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M0", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_8())), gopurs_runtime.Func2(func(next_9 gopurs_runtime.Value, acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_9.StrVal(), ">"), acc_10.StrVal()))
			}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal(), "Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M1", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_9())), gopurs_runtime.Func2(func(next_10 gopurs_runtime.Value, acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_10.StrVal(), ">"), acc_11.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal(), "a>b>c>Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M2", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_10())), gopurs_runtime.Func2(func(next_11 gopurs_runtime.Value, acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_11.StrVal(), ">"), acc_12.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal(), "Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M3", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_11())), gopurs_runtime.Func2(func(next_12 gopurs_runtime.Value, acc_13 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_12.StrVal(), ">"), acc_13.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal(), "a>b>c>Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M4", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_12())), gopurs_runtime.Func2(func(next_13 gopurs_runtime.Value, acc_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_13.StrVal(), ">"), acc_14.StrVal()))
			}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
			}()})))}).StrVal(), "a>b>c>Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			_dollar___unused_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M5", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_13())), gopurs_runtime.Func2(func(next_14 gopurs_runtime.Value, acc_15 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_14.StrVal(), ">"), acc_15.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal(), "a>b>c>Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_14_14
			_dollar___unused_15_15 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M6", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_14())), gopurs_runtime.Func2(func(next_15 gopurs_runtime.Value, acc_16 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_15.StrVal(), ">"), acc_16.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal(), "a>b>c>a>b>c>a>b>c>Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_15_15
			_dollar___unused_16_16 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldr - M7", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_15())), gopurs_runtime.Func2(func(next_16 gopurs_runtime.Value, acc_17 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_16.StrVal(), ">"), acc_17.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal(), "a>b>c>Start"})), gopurs_runtime.Value{})
			_ = _dollar___unused_16_16
			_dollar___unused_17_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M0", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_16())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_17
			}), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal(), ""})), gopurs_runtime.Value{})
			_ = _dollar___unused_17_17
			_dollar___unused_18_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M1", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_17())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_18
			}), Get_Main_m1()).StrVal(), "abc"})), gopurs_runtime.Value{})
			_ = _dollar___unused_18_18
			_dollar___unused_19_19 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M2", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_18())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_19
			}), Get_Main_m2()).StrVal(), ""})), gopurs_runtime.Value{})
			_ = _dollar___unused_19_19
			_dollar___unused_20_20 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M3", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_19())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_20 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_20
			}), Get_Main_m3()).StrVal(), "abc"})), gopurs_runtime.Value{})
			_ = _dollar___unused_20_20
			_dollar___unused_21_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M4", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_20())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_21
			}), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
			}()})))}).StrVal(), "abc"})), gopurs_runtime.Value{})
			_ = _dollar___unused_21_21
			_dollar___unused_22_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M5", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_21())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_22 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_22
			}), Get_Main_m5()).StrVal(), "abc"})), gopurs_runtime.Value{})
			_ = _dollar___unused_22_22
			_dollar___unused_23_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M6", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_22())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_23 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_23
			}), Get_Main_m6()).StrVal(), "abcabcabc"})), gopurs_runtime.Value{})
			_ = _dollar___unused_23_23
			_dollar___unused_24_24 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime_(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), "foldMap - M7", func(record struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}) struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			} {
				return struct {
					actual   gopurs_runtime.Value
					expected gopurs_runtime.Value
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Str(record.actual), gopurs_runtime.Str(record.expected)}
			}(struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main___cached_dict_23())), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_24
			}), Get_Main_m7()).StrVal(), "abc"})), gopurs_runtime.Value{})
			_ = _dollar___unused_24_24
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_11
}

var cache_Main___cached_dict_12 gopurs_runtime.Value
var once_Main___cached_dict_12 sync.Once

func Get_Main___cached_dict_12() gopurs_runtime.Value {
	once_Main___cached_dict_12.Do(func() {
		cache_Main___cached_dict_12 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_12
}

var cache_Main___cached_dict_13 gopurs_runtime.Value
var once_Main___cached_dict_13 sync.Once

func Get_Main___cached_dict_13() gopurs_runtime.Value {
	once_Main___cached_dict_13.Do(func() {
		cache_Main___cached_dict_13 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_13
}

var cache_Main___cached_dict_14 gopurs_runtime.Value
var once_Main___cached_dict_14 sync.Once

func Get_Main___cached_dict_14() gopurs_runtime.Value {
	once_Main___cached_dict_14.Do(func() {
		cache_Main___cached_dict_14 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_14
}

var cache_Main___cached_dict_15 gopurs_runtime.Value
var once_Main___cached_dict_15 sync.Once

func Get_Main___cached_dict_15() gopurs_runtime.Value {
	once_Main___cached_dict_15.Do(func() {
		cache_Main___cached_dict_15 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_15
}

var cache_Main___cached_dict_16 gopurs_runtime.Value
var once_Main___cached_dict_16 sync.Once

func Get_Main___cached_dict_16() gopurs_runtime.Value {
	once_Main___cached_dict_16.Do(func() {
		cache_Main___cached_dict_16 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_16
}

var cache_Main___cached_dict_17 gopurs_runtime.Value
var once_Main___cached_dict_17 sync.Once

func Get_Main___cached_dict_17() gopurs_runtime.Value {
	once_Main___cached_dict_17.Do(func() {
		cache_Main___cached_dict_17 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_17
}

var cache_Main___cached_dict_18 gopurs_runtime.Value
var once_Main___cached_dict_18 sync.Once

func Get_Main___cached_dict_18() gopurs_runtime.Value {
	once_Main___cached_dict_18.Do(func() {
		cache_Main___cached_dict_18 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_18
}

var cache_Main___cached_dict_19 gopurs_runtime.Value
var once_Main___cached_dict_19 sync.Once

func Get_Main___cached_dict_19() gopurs_runtime.Value {
	once_Main___cached_dict_19.Do(func() {
		cache_Main___cached_dict_19 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_19
}

var cache_Main___cached_dict_20 gopurs_runtime.Value
var once_Main___cached_dict_20 sync.Once

func Get_Main___cached_dict_20() gopurs_runtime.Value {
	once_Main___cached_dict_20.Do(func() {
		cache_Main___cached_dict_20 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_20
}

var cache_Main___cached_dict_21 gopurs_runtime.Value
var once_Main___cached_dict_21 sync.Once

func Get_Main___cached_dict_21() gopurs_runtime.Value {
	once_Main___cached_dict_21.Do(func() {
		cache_Main___cached_dict_21 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_21
}

var cache_Main___cached_dict_22 gopurs_runtime.Value
var once_Main___cached_dict_22 sync.Once

func Get_Main___cached_dict_22() gopurs_runtime.Value {
	once_Main___cached_dict_22.Do(func() {
		cache_Main___cached_dict_22 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_22
}

var cache_Main___cached_dict_23 gopurs_runtime.Value
var once_Main___cached_dict_23 sync.Once

func Get_Main___cached_dict_23() gopurs_runtime.Value {
	once_Main___cached_dict_23.Do(func() {
		cache_Main___cached_dict_23 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main___cached_dict_23
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
	V1 gopurs_runtime.Value
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

func Call_Main_foldrStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(Call_Data_Foldable_foldr(dictFoldable_0), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(next_1.StrVal(), ">"), acc_2.StrVal()))
	}), gopurs_runtime.Str("Start"))
}

func Call_Main_foldlStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(Call_Data_Foldable_foldl(dictFoldable_0), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(acc_1.StrVal(), "<"), next_2.StrVal()))
	}), gopurs_runtime.Str("Start"))
}

func Call_Main_foldableM(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_2_0 shape=App(Var) bindingType=(TypeVar m$scope39)
		mempty_2_0 := Call_Data_Monoid_mempty(dictMonoid_1)
		_ = mempty_2_0
		// TAST (Let): Semigroup0_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope39)])
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

func Call_Main_foldMapStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(Call_Data_Foldable_foldMap(dictFoldable_0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Rebox_Main_1039524714_2770120821(in *Constructor_Main_M4[gopurs_runtime.Value, string]) *Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1201789390_1950344881(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[string]{}
	out.V0 = in.V0
	out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Main_1302345387_2067946548(in *Constructor_Main_M5[gopurs_runtime.Value, string]) *Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
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

func Rebox_Main_1950344881_1201789390(in *Constructor_Data_Monoid_Monoid[string]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Main_2785108781_2554376626(in *Constructor_Main_M3[gopurs_runtime.Value, string]) *Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3741832558_67812977(in *Constructor_Main_M0[gopurs_runtime.Value, string]) *Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_M0[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
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
