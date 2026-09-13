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
				return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, func() []gopurs_runtime.Value {
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
	return cache_Main_M1
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
		}(), func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []string{"b"}
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
		cache_Main_m2 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(Rebox_Main_4256935660_1521903347((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, int64(0), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})})))}
	})
	return cache_Main_m2
}

var cache_Main_m1 gopurs_runtime.Value
var once_Main_m1 sync.Once

func Get_Main_m1() gopurs_runtime.Value {
	once_Main_m1.Do(func() {
		cache_Main_m1 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, "a", func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []string{"b", "c"}
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

var cache_Main_foldableM1 gopurs_runtime.Value
var once_Main_foldableM1 sync.Once

func Get_Main_foldableM1() gopurs_runtime.Value {
	once_Main_foldableM1.Do(func() {
		cache_Main_foldableM1 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))}
	})
	return cache_Main_foldableM1
}

var cache_Main_foldrStr__2116068004 gopurs_runtime.Value
var once_Main_foldrStr__2116068004 sync.Once

func Get_Main_foldrStr__2116068004() gopurs_runtime.Value {
	once_Main_foldrStr__2116068004.Do(func() {
		cache_Main_foldrStr__2116068004 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__2116068004(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__2116068004
}

var cache_Main_foldrStr__325071877 gopurs_runtime.Value
var once_Main_foldrStr__325071877 sync.Once

func Get_Main_foldrStr__325071877() gopurs_runtime.Value {
	once_Main_foldrStr__325071877.Do(func() {
		cache_Main_foldrStr__325071877 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__325071877(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__325071877
}

var cache_Main_foldrStr__412824550 gopurs_runtime.Value
var once_Main_foldrStr__412824550 sync.Once

func Get_Main_foldrStr__412824550() gopurs_runtime.Value {
	once_Main_foldrStr__412824550.Do(func() {
		cache_Main_foldrStr__412824550 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__412824550(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__412824550
}

var cache_Main_foldrStr__3602154567 gopurs_runtime.Value
var once_Main_foldrStr__3602154567 sync.Once

func Get_Main_foldrStr__3602154567() gopurs_runtime.Value {
	once_Main_foldrStr__3602154567.Do(func() {
		cache_Main_foldrStr__3602154567 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__3602154567(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__3602154567
}

var cache_Main_foldrStr__3287222048 gopurs_runtime.Value
var once_Main_foldrStr__3287222048 sync.Once

func Get_Main_foldrStr__3287222048() gopurs_runtime.Value {
	once_Main_foldrStr__3287222048.Do(func() {
		cache_Main_foldrStr__3287222048 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__3287222048(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__3287222048
}

var cache_Main_foldrStr__1496225921 gopurs_runtime.Value
var once_Main_foldrStr__1496225921 sync.Once

func Get_Main_foldrStr__1496225921() gopurs_runtime.Value {
	once_Main_foldrStr__1496225921.Do(func() {
		cache_Main_foldrStr__1496225921 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__1496225921(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__1496225921
}

var cache_Main_foldrStr__1583978594 gopurs_runtime.Value
var once_Main_foldrStr__1583978594 sync.Once

func Get_Main_foldrStr__1583978594() gopurs_runtime.Value {
	once_Main_foldrStr__1583978594.Do(func() {
		cache_Main_foldrStr__1583978594 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__1583978594(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__1583978594
}

var cache_Main_foldrStr__478341315 gopurs_runtime.Value
var once_Main_foldrStr__478341315 sync.Once

func Get_Main_foldrStr__478341315() gopurs_runtime.Value {
	once_Main_foldrStr__478341315.Do(func() {
		cache_Main_foldrStr__478341315 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldrStr__478341315(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldrStr__478341315
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

var cache_Main_foldlStr__2116068004 gopurs_runtime.Value
var once_Main_foldlStr__2116068004 sync.Once

func Get_Main_foldlStr__2116068004() gopurs_runtime.Value {
	once_Main_foldlStr__2116068004.Do(func() {
		cache_Main_foldlStr__2116068004 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__2116068004(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__2116068004
}

var cache_Main_foldlStr__325071877 gopurs_runtime.Value
var once_Main_foldlStr__325071877 sync.Once

func Get_Main_foldlStr__325071877() gopurs_runtime.Value {
	once_Main_foldlStr__325071877.Do(func() {
		cache_Main_foldlStr__325071877 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__325071877(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__325071877
}

var cache_Main_foldlStr__412824550 gopurs_runtime.Value
var once_Main_foldlStr__412824550 sync.Once

func Get_Main_foldlStr__412824550() gopurs_runtime.Value {
	once_Main_foldlStr__412824550.Do(func() {
		cache_Main_foldlStr__412824550 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__412824550(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__412824550
}

var cache_Main_foldlStr__3602154567 gopurs_runtime.Value
var once_Main_foldlStr__3602154567 sync.Once

func Get_Main_foldlStr__3602154567() gopurs_runtime.Value {
	once_Main_foldlStr__3602154567.Do(func() {
		cache_Main_foldlStr__3602154567 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__3602154567(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__3602154567
}

var cache_Main_foldlStr__3287222048 gopurs_runtime.Value
var once_Main_foldlStr__3287222048 sync.Once

func Get_Main_foldlStr__3287222048() gopurs_runtime.Value {
	once_Main_foldlStr__3287222048.Do(func() {
		cache_Main_foldlStr__3287222048 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__3287222048(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__3287222048
}

var cache_Main_foldlStr__1496225921 gopurs_runtime.Value
var once_Main_foldlStr__1496225921 sync.Once

func Get_Main_foldlStr__1496225921() gopurs_runtime.Value {
	once_Main_foldlStr__1496225921.Do(func() {
		cache_Main_foldlStr__1496225921 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__1496225921(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__1496225921
}

var cache_Main_foldlStr__1583978594 gopurs_runtime.Value
var once_Main_foldlStr__1583978594 sync.Once

func Get_Main_foldlStr__1583978594() gopurs_runtime.Value {
	once_Main_foldlStr__1583978594.Do(func() {
		cache_Main_foldlStr__1583978594 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__1583978594(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__1583978594
}

var cache_Main_foldlStr__478341315 gopurs_runtime.Value
var once_Main_foldlStr__478341315 sync.Once

func Get_Main_foldlStr__478341315() gopurs_runtime.Value {
	once_Main_foldlStr__478341315.Do(func() {
		cache_Main_foldlStr__478341315 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldlStr__478341315(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldlStr__478341315
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

var cache_Main_foldMapStr__2116068004 gopurs_runtime.Value
var once_Main_foldMapStr__2116068004 sync.Once

func Get_Main_foldMapStr__2116068004() gopurs_runtime.Value {
	once_Main_foldMapStr__2116068004.Do(func() {
		cache_Main_foldMapStr__2116068004 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__2116068004(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__2116068004
}

var cache_Main_foldMapStr__325071877 gopurs_runtime.Value
var once_Main_foldMapStr__325071877 sync.Once

func Get_Main_foldMapStr__325071877() gopurs_runtime.Value {
	once_Main_foldMapStr__325071877.Do(func() {
		cache_Main_foldMapStr__325071877 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__325071877(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__325071877
}

var cache_Main_foldMapStr__412824550 gopurs_runtime.Value
var once_Main_foldMapStr__412824550 sync.Once

func Get_Main_foldMapStr__412824550() gopurs_runtime.Value {
	once_Main_foldMapStr__412824550.Do(func() {
		cache_Main_foldMapStr__412824550 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__412824550(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__412824550
}

var cache_Main_foldMapStr__3602154567 gopurs_runtime.Value
var once_Main_foldMapStr__3602154567 sync.Once

func Get_Main_foldMapStr__3602154567() gopurs_runtime.Value {
	once_Main_foldMapStr__3602154567.Do(func() {
		cache_Main_foldMapStr__3602154567 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__3602154567(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__3602154567
}

var cache_Main_foldMapStr__3287222048 gopurs_runtime.Value
var once_Main_foldMapStr__3287222048 sync.Once

func Get_Main_foldMapStr__3287222048() gopurs_runtime.Value {
	once_Main_foldMapStr__3287222048.Do(func() {
		cache_Main_foldMapStr__3287222048 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__3287222048(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__3287222048
}

var cache_Main_foldMapStr__1496225921 gopurs_runtime.Value
var once_Main_foldMapStr__1496225921 sync.Once

func Get_Main_foldMapStr__1496225921() gopurs_runtime.Value {
	once_Main_foldMapStr__1496225921.Do(func() {
		cache_Main_foldMapStr__1496225921 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__1496225921(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__1496225921
}

var cache_Main_foldMapStr__1583978594 gopurs_runtime.Value
var once_Main_foldMapStr__1583978594 sync.Once

func Get_Main_foldMapStr__1583978594() gopurs_runtime.Value {
	once_Main_foldMapStr__1583978594.Do(func() {
		cache_Main_foldMapStr__1583978594 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__1583978594(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__1583978594
}

var cache_Main_foldMapStr__478341315 gopurs_runtime.Value
var once_Main_foldMapStr__478341315 sync.Once

func Get_Main_foldMapStr__478341315() gopurs_runtime.Value {
	once_Main_foldMapStr__478341315.Do(func() {
		cache_Main_foldMapStr__478341315 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foldMapStr__478341315(__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_foldMapStr__478341315
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M0", struct {
			actual   string
			expected string
		}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_0 gopurs_runtime.Value, next_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_0.StrVal()) + ("<")) + (next_1.StrVal()))
		}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal(), "Start"}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M1", struct {
				actual   string
				expected string
			}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
			}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal(), "Start<a<b<c"}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M2", struct {
					actual   string
					expected string
				}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_2 gopurs_runtime.Value, next_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str(((acc_2.StrVal()) + ("<")) + (next_3.StrVal()))
				}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal(), "Start"}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M3", struct {
						actual   string
						expected string
					}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_3 gopurs_runtime.Value, next_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_3.StrVal()) + ("<")) + (next_4.StrVal()))
					}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal(), "Start<a<b<c"}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M4", struct {
							actual   string
							expected string
						}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, next_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str(((acc_4.StrVal()) + ("<")) + (next_5.StrVal()))
						}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
						}()})))}).StrVal(), "Start<a<b<c"}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M5", struct {
								actual   string
								expected string
							}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_5 gopurs_runtime.Value, next_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str(((acc_5.StrVal()) + ("<")) + (next_6.StrVal()))
							}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal(), "Start<a<b<c"}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M6", struct {
									actual   string
									expected string
								}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_6 gopurs_runtime.Value, next_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str(((acc_6.StrVal()) + ("<")) + (next_7.StrVal()))
								}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal(), "Start<a<b<c<a<b<c<a<b<c"}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldl - M7", struct {
										actual   string
										expected string
									}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_7 gopurs_runtime.Value, next_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Str(((acc_7.StrVal()) + ("<")) + (next_8.StrVal()))
									}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal(), "Start<a<b<c"}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M0", struct {
											actual   string
											expected string
										}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_8 gopurs_runtime.Value, acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Str(((next_8.StrVal()) + (">")) + (acc_9.StrVal()))
										}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal(), "Start"}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M1", struct {
												actual   string
												expected string
											}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_9 gopurs_runtime.Value, acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Str(((next_9.StrVal()) + (">")) + (acc_10.StrVal()))
											}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal(), "a>b>c>Start"}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M2", struct {
													actual   string
													expected string
												}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_10 gopurs_runtime.Value, acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Str(((next_10.StrVal()) + (">")) + (acc_11.StrVal()))
												}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal(), "Start"}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M3", struct {
														actual   string
														expected string
													}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_11 gopurs_runtime.Value, acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Str(((next_11.StrVal()) + (">")) + (acc_12.StrVal()))
													}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal(), "a>b>c>Start"}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M4", struct {
															actual   string
															expected string
														}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_12 gopurs_runtime.Value, acc_13 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Str(((next_12.StrVal()) + (">")) + (acc_13.StrVal()))
														}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
														}()})))}).StrVal(), "a>b>c>Start"}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M5", struct {
																actual   string
																expected string
															}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_13 gopurs_runtime.Value, acc_14 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Str(((next_13.StrVal()) + (">")) + (acc_14.StrVal()))
															}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal(), "a>b>c>Start"}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M6", struct {
																	actual   string
																	expected string
																}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_14 gopurs_runtime.Value, acc_15 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Str(((next_14.StrVal()) + (">")) + (acc_15.StrVal()))
																}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal(), "a>b>c>a>b>c>a>b>c>Start"}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldr - M7", struct {
																		actual   string
																		expected string
																	}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_15 gopurs_runtime.Value, acc_16 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Str(((next_15.StrVal()) + (">")) + (acc_16.StrVal()))
																	}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal(), "a>b>c>Start"}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M0", struct {
																			actual   string
																			expected string
																		}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal(), ""}), gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M1", struct {
																				actual   string
																				expected string
																			}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m1()).StrVal(), "abc"}), gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
																				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M2", struct {
																					actual   string
																					expected string
																				}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m2()).StrVal(), ""}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
																					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M3", struct {
																						actual   string
																						expected string
																					}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m3()).StrVal(), "abc"}), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
																						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M4", struct {
																							actual   string
																							expected string
																						}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
																						}()})))}).StrVal(), "abc"}), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
																							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M5", struct {
																								actual   string
																								expected string
																							}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m5()).StrVal(), "abc"}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
																								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M6", struct {
																									actual   string
																									expected string
																								}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m6()).StrVal(), "abcabcabc"}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
																									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("foldMap - M7", struct {
																										actual   string
																										expected string
																									}{gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m7()).StrVal(), "abc"}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
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
									}))
								}))
							}))
						}))
					}))
				}))
			}))
		}))
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
		return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer(Rebox_Main_4004653231_3660606000((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, __eta_norm_1_0, func() []gopurs_runtime.Value {
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

func Call_Main_foldrStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
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
					__t2 = gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 108241190 {
					__t2 = gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 2066233029 {
					__t2 = gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V3)), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V4), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))))))))))
					goto end_branch_2
				} else {

				}
			}
			{
				if m_5.Type == 9 && m_5.IntVal == 1168316772 {
					__t2 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "a")), gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "fa")), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1)), f_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))))
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
					arr_val_foldlArray3 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())
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
					arr_val_foldlArray3 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())
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
					arr_val_foldlArray3 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())
					_ = arr_val_foldlArray3
					res_go_foldlArray3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, func() gopurs_runtime.Value {
						arr_val_foldlArray6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
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
						arr_val_foldlArray6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
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
				__t5 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "fa")))
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 108241190 {
				__t5 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "fa")))
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2066233029 {
				__t5 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "fa"))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "fa"))), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V4), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)))
				goto end_branch_5
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1168316772 {
				// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
				__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, v2_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "fa")))
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

func Call_Main_foldrStr__2116068004(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__2116068004:
	for {
		if false {
			continue foldrStr__2116068004
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal()
	}
}

func Call_Main_foldrStr__325071877(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__325071877:
	for {
		if false {
			continue foldrStr__325071877
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal()
	}
}

func Call_Main_foldrStr__412824550(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__412824550:
	for {
		if false {
			continue foldrStr__412824550
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal()
	}
}

func Call_Main_foldrStr__3602154567(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__3602154567:
	for {
		if false {
			continue foldrStr__3602154567
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal()
	}
}

func Call_Main_foldrStr__3287222048(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__3287222048:
	for {
		if false {
			continue foldrStr__3287222048
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
		}()})))}).StrVal()
	}
}

func Call_Main_foldrStr__1496225921(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__1496225921:
	for {
		if false {
			continue foldrStr__1496225921
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal()
	}
}

func Call_Main_foldrStr__1583978594(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__1583978594:
	for {
		if false {
			continue foldrStr__1583978594
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal()
	}
}

func Call_Main_foldrStr__478341315(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldrStr__478341315:
	for {
		if false {
			continue foldrStr__478341315
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldr"), gopurs_runtime.Func2(func(next_1 gopurs_runtime.Value, acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal()
	}
}

func Call_Main_foldlStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
	}), gopurs_runtime.Str("Start"))
}

func Call_Main_foldlStr__2116068004(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__2116068004:
	for {
		if false {
			continue foldlStr__2116068004
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal()
	}
}

func Call_Main_foldlStr__325071877(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__325071877:
	for {
		if false {
			continue foldlStr__325071877
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal()
	}
}

func Call_Main_foldlStr__412824550(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__412824550:
	for {
		if false {
			continue foldlStr__412824550
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal()
	}
}

func Call_Main_foldlStr__3602154567(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__3602154567:
	for {
		if false {
			continue foldlStr__3602154567
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal()
	}
}

func Call_Main_foldlStr__3287222048(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__3287222048:
	for {
		if false {
			continue foldlStr__3287222048
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
		}()})))}).StrVal()
	}
}

func Call_Main_foldlStr__1496225921(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__1496225921:
	for {
		if false {
			continue foldlStr__1496225921
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal()
	}
}

func Call_Main_foldlStr__1583978594(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__1583978594:
	for {
		if false {
			continue foldlStr__1583978594
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal()
	}
}

func Call_Main_foldlStr__478341315(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldlStr__478341315:
	for {
		if false {
			continue foldlStr__478341315
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldl"), gopurs_runtime.Func2(func(acc_1 gopurs_runtime.Value, next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal()
	}
}

func Call_Main_foldMapStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Main_foldMapStr__2116068004(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__2116068004:
	for {
		if false {
			continue foldMapStr__2116068004
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(Rebox_Main_3741832558_67812977(nil))}).StrVal()
	}
}

func Call_Main_foldMapStr__325071877(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__325071877:
	for {
		if false {
			continue foldMapStr__325071877
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m1()).StrVal()
	}
}

func Call_Main_foldMapStr__412824550(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__412824550:
	for {
		if false {
			continue foldMapStr__412824550
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m2()).StrVal()
	}
}

func Call_Main_foldMapStr__3602154567(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__3602154567:
	for {
		if false {
			continue foldMapStr__3602154567
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m3()).StrVal()
	}
}

func Call_Main_foldMapStr__3287222048(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__3287222048:
	for {
		if false {
			continue foldMapStr__3287222048
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(Rebox_Main_1039524714_2770120821((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
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
		}()})))}).StrVal()
	}
}

func Call_Main_foldMapStr__1496225921(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__1496225921:
	for {
		if false {
			continue foldMapStr__1496225921
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m5()).StrVal()
	}
}

func Call_Main_foldMapStr__1583978594(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__1583978594:
	for {
		if false {
			continue foldMapStr__1583978594
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m6()).StrVal()
	}
}

func Call_Main_foldMapStr__478341315(__eta_norm_0_unused_0_loop gopurs_runtime.Value) string {
foldMapStr__478341315:
	for {
		if false {
			continue foldMapStr__478341315
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_foldableM(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Get_Main_m7()).StrVal()
	}
}

func Rebox_Main_1039524714_2770120821(in *Constructor_Main_M4[gopurs_runtime.Value, string]) *Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
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
	out := &Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value]{}
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
	out.V1 = in.V1
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
