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
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_identity(x_0_box.StrVal()))
		})
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
			return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() *struct {
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
			orig := func() *struct {
				a           string
				arrayIgnore []int64
				fIgnore     []int64
				fa          []string
				ignore      int64
				zArrayA     []string
			} {
				orig := gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str("a"), func() gopurs_runtime.Value {
					arr := []int64{2, 3}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := []int64{4}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := []string{"b"}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(1), func() gopurs_runtime.Value {
					arr := []string{"c"}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
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
				return &clone
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
		}()
	})
	return cache_Main_recordValue
}

var cache_Main_m7 gopurs_runtime.Value
var once_Main_m7 sync.Once

func Get_Main_m7() gopurs_runtime.Value {
	once_Main_m7.Do(func() {
		cache_Main_m7 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			arr := [][]*struct {
				nested *struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				}
			}{[]*struct {
				nested *struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				}
			}{func() *struct {
				nested *struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				}
			} {
				orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
					orig := func() *struct {
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
						return &clone
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
					nested *struct {
						a           string
						arrayIgnore []int64
						fIgnore     []int64
						fa          []string
						ignore      int64
						zArrayA     []string
					}
				}{}
				clone.nested = func() *struct {
					a           string
					arrayIgnore []int64
					fIgnore     []int64
					fa          []string
					ignore      int64
					zArrayA     []string
				} {
					orig := gopurs_runtime.RecordGet(orig, "nested")
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
					return &clone
				}()
				return &clone
			}()}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordDict([]string{"nested"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
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
							}()})
						}()
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m7
}

var cache_Main_m6 gopurs_runtime.Value
var once_Main_m6 sync.Once

func Get_Main_m6() gopurs_runtime.Value {
	once_Main_m6.Do(func() {
		cache_Main_m6 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6[gopurs_runtime.Value, string]{1, 1, gopurs_runtime.Str("a"), func() []int64 {
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
			orig := func() *struct {
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
				return &clone
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
		}(), func() *struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() *struct {
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
					return &clone
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
			return &clone
		}()}))}
	})
	return cache_Main_m6
}

var cache_Main_m5 gopurs_runtime.Value
var once_Main_m5 sync.Once

func Get_Main_m5() gopurs_runtime.Value {
	once_Main_m5.Do(func() {
		cache_Main_m5 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5[gopurs_runtime.Value, string]{1, func() *struct {
			nested gopurs_runtime.Value
		} {
			orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
				orig := func() *struct {
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
					return &clone
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
			return &clone
		}()}))}
	})
	return cache_Main_m5
}

var cache_Main_m4 gopurs_runtime.Value
var once_Main_m4 sync.Once

func Get_Main_m4() gopurs_runtime.Value {
	once_Main_m4.Do(func() {
		cache_Main_m4 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			orig := func() *struct {
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
				return &clone
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
		}()}))}
	})
	return cache_Main_m4
}

var cache_Main_m3 gopurs_runtime.Value
var once_Main_m3 sync.Once

func Get_Main_m3() gopurs_runtime.Value {
	once_Main_m3.Do(func() {
		cache_Main_m3 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
			arr := []string{"a", "b", "c"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m3
}

var cache_Main_m2 gopurs_runtime.Value
var once_Main_m2 sync.Once

func Get_Main_m2() gopurs_runtime.Value {
	once_Main_m2.Do(func() {
		cache_Main_m2 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2[gopurs_runtime.Value, string]{1, 0, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_m2
}

var cache_Main_m1 gopurs_runtime.Value
var once_Main_m1 sync.Once

func Get_Main_m1() gopurs_runtime.Value {
	once_Main_m1.Do(func() {
		cache_Main_m1 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1[gopurs_runtime.Value, string]{1, gopurs_runtime.Str("a"), func() []gopurs_runtime.Value {
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
		}()}))}
	})
	return cache_Main_m1
}

var cache_Main_m0 gopurs_runtime.Value
var once_Main_m0 sync.Once

func Get_Main_m0() gopurs_runtime.Value {
	once_Main_m0.Do(func() {
		cache_Main_m0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
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

var cache_Main_foldrStr__1016676313 gopurs_runtime.Value
var once_Main_foldrStr__1016676313 sync.Once

func Get_Main_foldrStr__1016676313() gopurs_runtime.Value {
	once_Main_foldrStr__1016676313.Do(func() {
		cache_Main_foldrStr__1016676313 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldrStr__1016676313(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Main_foldrStr__1016676313
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

var cache_Main_foldlStr__1016676313 gopurs_runtime.Value
var once_Main_foldlStr__1016676313 sync.Once

func Get_Main_foldlStr__1016676313() gopurs_runtime.Value {
	once_Main_foldlStr__1016676313.Do(func() {
		cache_Main_foldlStr__1016676313 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldlStr__1016676313(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Main_foldlStr__1016676313
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
		cache_Main_foldableM1 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar m)
			mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
			_ = mempty_1_0
			// TAST (Let): Semigroup0_2_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
			Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_2_1
			return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t19 gopurs_runtime.Value
					{
						if m_4.Type == 9 && m_4.IntVal == 3852365315 {
							__t19 = mempty_1_0
							goto end_branch_19
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 769986722 {
							// TAST (Let): Semigroup0_5_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_2
							__t19 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_2.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)))
							goto end_branch_19
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 2727978561 {
							__t19 = mempty_1_0
							goto end_branch_19
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 1830062304 {
							// TAST (Let): Semigroup0_5_3 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_3
							__t19 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_3.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0)
							goto end_branch_19
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 3190619783 {
							// TAST (Let): Semigroup0_5_4 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_4
							// TAST (Let): Semigroup0_5_5 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_5
							__t19 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_4.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_5.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))
							goto end_branch_19
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 108241190 {
							// TAST (Let): Semigroup0_5_6 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_6
							// TAST (Let): Semigroup0_5_7 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_7
							__t19 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0.nested, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_7.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))
							goto end_branch_19
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 2066233029 {
							// TAST (Let): Semigroup0_5_8 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_8
							// TAST (Let): Semigroup0_5_9 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_9
							// TAST (Let): Semigroup0_5_10 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_10
							// TAST (Let): Semigroup0_5_11 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_11
							// TAST (Let): Semigroup0_5_12 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_12
							// TAST (Let): Semigroup0_5_13 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_13
							__t19 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_8.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V3)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_9.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_10.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_11.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_12.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V7.nested, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_13.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))))))))
							goto end_branch_19
						} else {

						}
					}
					{
						if m_4.Type == 9 && m_4.IntVal == 1168316772 {
							// TAST (Let): Semigroup0_5_14 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_5_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_5_14
							// TAST (Let): Semigroup0_6_16 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_6_16
							// TAST (Let): __local_var_6_15 shape=Let(App(Var)) bindingType=(Func [(TypeApp (TypeVar f) [(Record (Row [nested: (Record (Row [a: (TypeVar a), fa: (TypeApp (TypeVar f) [(TypeVar a)]), zArrayA: (Array (TypeVar a)), arrayIgnore: (Array Int), fIgnore: (TypeApp (TypeVar f) [Int]), ignore: Int] Any))] Any))])] (TypeVar m))
							__local_var_6_15 := gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): Semigroup0_9_17 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
									Semigroup0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
									_ = Semigroup0_9_17
									// TAST (Let): Semigroup0_9_18 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
									Semigroup0_9_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
									_ = Semigroup0_9_18
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_16.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_7, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_17.V0), gopurs_runtime.Apply(f_3, x_10), acc_11)
										})
									}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_7, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_18.V0), gopurs_runtime.Apply(f_3, x_10), acc_11)
										})
									}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_7, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())))), acc_8)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
							_ = __local_var_6_15
							__t19 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_14.V0), gopurs_runtime.Apply(__local_var_6_15, x_7), acc_8)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0)
							goto end_branch_19
						} else {

						}
					}
					{
						__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_19:
					return __t19
				})
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t20 gopurs_runtime.Value
					{
						if m_2.Type == 9 && m_2.IntVal == 3852365315 {
							__t20 = z_1
							goto end_branch_20
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 769986722 {
							__t20 = func() gopurs_runtime.Value {
								arr_val_foldlArray5 := gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)
								_ = arr_val_foldlArray5
								res_go_foldlArray5 := gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)
								_ = res_go_foldlArray5
								arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
								_ = arr_go_foldlArray5
								for _, v_foldlArray5 := range *arr_go_foldlArray5 {
									res_go_foldlArray5 = gopurs_runtime.Apply2(f_0, res_go_foldlArray5, v_foldlArray5)
								}
								return res_go_foldlArray5
							}()
							goto end_branch_20
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 2727978561 {
							__t20 = z_1
							goto end_branch_20
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 1830062304 {
							__t20 = func() gopurs_runtime.Value {
								arr_val_foldlArray5 := (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0
								_ = arr_val_foldlArray5
								res_go_foldlArray5 := z_1
								_ = res_go_foldlArray5
								arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
								_ = arr_go_foldlArray5
								for _, v_foldlArray5 := range *arr_go_foldlArray5 {
									res_go_foldlArray5 = gopurs_runtime.Apply2(f_0, res_go_foldlArray5, v_foldlArray5)
								}
								return res_go_foldlArray5
							}()
							goto end_branch_20
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 3190619783 {
							__t20 = func() gopurs_runtime.Value {
								arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray5
								res_go_foldlArray5 := func() gopurs_runtime.Value {
									arr_val_foldlArray6 := gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "fa")
									_ = arr_val_foldlArray6
									res_go_foldlArray6 := gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a"))
									_ = res_go_foldlArray6
									arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
									_ = arr_go_foldlArray6
									for _, v_foldlArray6 := range *arr_go_foldlArray6 {
										res_go_foldlArray6 = gopurs_runtime.Apply2(f_0, res_go_foldlArray6, v_foldlArray6)
									}
									return res_go_foldlArray6
								}()
								_ = res_go_foldlArray5
								arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
								_ = arr_go_foldlArray5
								for _, v_foldlArray5 := range *arr_go_foldlArray5 {
									res_go_foldlArray5 = gopurs_runtime.Apply2(f_0, res_go_foldlArray5, v_foldlArray5)
								}
								return res_go_foldlArray5
							}()
							goto end_branch_20
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 108241190 {
							__t20 = func() gopurs_runtime.Value {
								arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray5
								res_go_foldlArray5 := func() gopurs_runtime.Value {
									arr_val_foldlArray6 := gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "fa")
									_ = arr_val_foldlArray6
									res_go_foldlArray6 := gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a"))
									_ = res_go_foldlArray6
									arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
									_ = arr_go_foldlArray6
									for _, v_foldlArray6 := range *arr_go_foldlArray6 {
										res_go_foldlArray6 = gopurs_runtime.Apply2(f_0, res_go_foldlArray6, v_foldlArray6)
									}
									return res_go_foldlArray6
								}()
								_ = res_go_foldlArray5
								arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
								_ = arr_go_foldlArray5
								for _, v_foldlArray5 := range *arr_go_foldlArray5 {
									res_go_foldlArray5 = gopurs_runtime.Apply2(f_0, res_go_foldlArray5, v_foldlArray5)
								}
								return res_go_foldlArray5
							}()
							goto end_branch_20
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 2066233029 {
							__t20 = func() gopurs_runtime.Value {
								arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray5
								res_go_foldlArray5 := func() gopurs_runtime.Value {
									arr_val_foldlArray6 := gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "fa")
									_ = arr_val_foldlArray6
									res_go_foldlArray6 := gopurs_runtime.Apply2(f_0, func() gopurs_runtime.Value {
										arr_val_foldlArray8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())
										_ = arr_val_foldlArray8
										res_go_foldlArray8 := func() gopurs_runtime.Value {
											arr_val_foldlArray9 := gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "fa")
											_ = arr_val_foldlArray9
											res_go_foldlArray9 := gopurs_runtime.Apply2(f_0, func() gopurs_runtime.Value {
												arr_val_foldlArray11 := (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4
												_ = arr_val_foldlArray11
												res_go_foldlArray11 := func() gopurs_runtime.Value {
													arr_val_foldlArray12 := gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3)
													_ = arr_val_foldlArray12
													res_go_foldlArray12 := gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)
													_ = res_go_foldlArray12
													arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
													_ = arr_go_foldlArray12
													for _, v_foldlArray12 := range *arr_go_foldlArray12 {
														res_go_foldlArray12 = gopurs_runtime.Apply2(f_0, res_go_foldlArray12, v_foldlArray12)
													}
													return res_go_foldlArray12
												}()
												_ = res_go_foldlArray11
												arr_go_foldlArray11 := (*[]gopurs_runtime.Value)(arr_val_foldlArray11.UnsafePtr)
												_ = arr_go_foldlArray11
												for _, v_foldlArray11 := range *arr_go_foldlArray11 {
													res_go_foldlArray11 = gopurs_runtime.Apply2(f_0, res_go_foldlArray11, v_foldlArray11)
												}
												return res_go_foldlArray11
											}(), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a"))
											_ = res_go_foldlArray9
											arr_go_foldlArray9 := (*[]gopurs_runtime.Value)(arr_val_foldlArray9.UnsafePtr)
											_ = arr_go_foldlArray9
											for _, v_foldlArray9 := range *arr_go_foldlArray9 {
												res_go_foldlArray9 = gopurs_runtime.Apply2(f_0, res_go_foldlArray9, v_foldlArray9)
											}
											return res_go_foldlArray9
										}()
										_ = res_go_foldlArray8
										arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
										_ = arr_go_foldlArray8
										for _, v_foldlArray8 := range *arr_go_foldlArray8 {
											res_go_foldlArray8 = gopurs_runtime.Apply2(f_0, res_go_foldlArray8, v_foldlArray8)
										}
										return res_go_foldlArray8
									}(), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a"))
									_ = res_go_foldlArray6
									arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
									_ = arr_go_foldlArray6
									for _, v_foldlArray6 := range *arr_go_foldlArray6 {
										res_go_foldlArray6 = gopurs_runtime.Apply2(f_0, res_go_foldlArray6, v_foldlArray6)
									}
									return res_go_foldlArray6
								}()
								_ = res_go_foldlArray5
								arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
								_ = arr_go_foldlArray5
								for _, v_foldlArray5 := range *arr_go_foldlArray5 {
									res_go_foldlArray5 = gopurs_runtime.Apply2(f_0, res_go_foldlArray5, v_foldlArray5)
								}
								return res_go_foldlArray5
							}()
							goto end_branch_20
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 1168316772 {
							__t20 = func() gopurs_runtime.Value {
								arr_val_foldlArray5 := (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0
								_ = arr_val_foldlArray5
								res_go_foldlArray5 := z_1
								_ = res_go_foldlArray5
								arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
								_ = arr_go_foldlArray5
								for _, v_foldlArray5 := range *arr_go_foldlArray5 {
									res_go_foldlArray5 = gopurs_runtime.Apply2(gopurs_runtime.Apply(Get_Data_Foldable_foldlArray(), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
											return func() gopurs_runtime.Value {
												arr_val_foldlArray9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_4, "nested"), "zArrayA").UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}())
												_ = arr_val_foldlArray9
												res_go_foldlArray9 := func() gopurs_runtime.Value {
													arr_val_foldlArray10 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_4, "nested"), "fa")
													_ = arr_val_foldlArray10
													res_go_foldlArray10 := gopurs_runtime.Apply2(f_0, v1_3, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_4, "nested"), "a"))
													_ = res_go_foldlArray10
													arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
													_ = arr_go_foldlArray10
													for _, v_foldlArray10 := range *arr_go_foldlArray10 {
														res_go_foldlArray10 = gopurs_runtime.Apply2(f_0, res_go_foldlArray10, v_foldlArray10)
													}
													return res_go_foldlArray10
												}()
												_ = res_go_foldlArray9
												arr_go_foldlArray9 := (*[]gopurs_runtime.Value)(arr_val_foldlArray9.UnsafePtr)
												_ = arr_go_foldlArray9
												for _, v_foldlArray9 := range *arr_go_foldlArray9 {
													res_go_foldlArray9 = gopurs_runtime.Apply2(f_0, res_go_foldlArray9, v_foldlArray9)
												}
												return res_go_foldlArray9
											}()
										})
									})), res_go_foldlArray5, v_foldlArray5)
								}
								return res_go_foldlArray5
							}()
							goto end_branch_20
						} else {

						}
					}
					{
						__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_20:
					return __t20
				})
			})
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t22 gopurs_runtime.Value
					{
						if m_2.Type == 9 && m_2.IntVal == 3852365315 {
							__t22 = z_1
							goto end_branch_22
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 769986722 {
							__t22 = gopurs_runtime.Apply2(f_0, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, z_1, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)))
							goto end_branch_22
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 2727978561 {
							__t22 = z_1
							goto end_branch_22
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 1830062304 {
							__t22 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, z_1, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)
							goto end_branch_22
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 3190619783 {
							__t22 = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, z_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, "fa")))
							goto end_branch_22
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 108241190 {
							__t22 = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, z_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested, "fa")))
							goto end_branch_22
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 2066233029 {
							__t22 = gopurs_runtime.Apply2(f_0, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, z_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V7.nested, "fa"))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V6, "fa"))), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V4), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V3)))
							goto end_branch_22
						} else {

						}
					}
					{
						if m_2.Type == 9 && m_2.IntVal == 1168316772 {
							// TAST (Let): __local_var_3_21 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(Record (Row [nested: (Record (Row [a: (TypeVar a), zArrayA: (Array (TypeVar a)), fa: (TypeApp (TypeVar f) [(TypeVar a)]), arrayIgnore: (Array Int), fIgnore: (TypeApp (TypeVar f) [Int]), ignore: Int] Any))] Any))])] (TypeVar b))
							__local_var_3_21 := gopurs_runtime.Apply(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, v2_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa")))
								})
							}))
							_ = __local_var_3_21
							__t22 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(__local_var_3_21, a_5, b_4)
								})
							}), z_1, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)
							goto end_branch_22
						} else {

						}
					}
					{
						__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_22:
					return __t22
				})
			})
		})}))}
	})
	return cache_Main_foldableM1
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

var cache_Main_foldMapStr__1016676313 gopurs_runtime.Value
var once_Main_foldMapStr__1016676313 sync.Once

func Get_Main_foldMapStr__1016676313() gopurs_runtime.Value {
	once_Main_foldMapStr__1016676313.Do(func() {
		cache_Main_foldMapStr__1016676313 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldMapStr__1016676313(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Main_foldMapStr__1016676313
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_0_0 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_1 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_0.StrVal()) + ("<")) + (next_1.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}).StrVal()), gopurs_runtime.Str("Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_0_0
			// TAST (Let): result_1_1 shape=Other bindingType=Boolean
			result_1_1 := (__local_var_0_0.actual) == (gopurs_runtime.Str(__local_var_0_0.expected).StrVal())
			_ = result_1_1
			// TAST (Let): message_2_2 shape=Other bindingType=String
			message_2_2 := ((("foldl - M0\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_0_0.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_0_0.actual)).StrVal())
			_ = message_2_2
			// TAST (Let): __local_var_3_3 shape=Let(Let(EffectBind(App(Var)))) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_3_3 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_3_4 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
				__local_var_3_4 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_2_2))
				_ = __local_var_3_4
				var __t6 gopurs_runtime.Value
				{
					if (result_1_1) != (true) {
						__t6 = __local_var_3_4
						goto end_branch_6
					} else {

					}
				}
				{
					if result_1_1 {
						__t6 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							return Get_Data_Unit_unit()
						})
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_6:
				// TAST (Let): __local_var_4_5 shape=Branch(Other, EffectPure, def=Other) bindingType=(TypeApp (TypeVar m) [Unit])
				__local_var_4_5 := __t6
				_ = __local_var_4_5
				_dollar___unused_5_7 := gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Value{})
				_ = _dollar___unused_5_7
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_2_2), gopurs_runtime.Bool(result_1_1)), gopurs_runtime.Value{})
			})
			_ = __local_var_3_3
			_dollar___unused_4_8 := gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{})
			_ = _dollar___unused_4_8
			// TAST (Let): __local_var_5_9 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_5_9 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_5.StrVal()) + ("<")) + (next_6.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal()), gopurs_runtime.Str("Start<a<b<c"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_5_9
			// TAST (Let): result_6_10 shape=Other bindingType=Boolean
			result_6_10 := (__local_var_5_9.actual) == (gopurs_runtime.Str(__local_var_5_9.expected).StrVal())
			_ = result_6_10
			// TAST (Let): message_7_11 shape=Other bindingType=String
			message_7_11 := ((("foldl - M1\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_5_9.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_5_9.actual)).StrVal())
			_ = message_7_11
			// TAST (Let): __local_var_8_13 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_8_13 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_7_11))
			_ = __local_var_8_13
			var __t15 gopurs_runtime.Value
			{
				if (result_6_10) != (true) {
					__t15 = __local_var_8_13
					goto end_branch_15
				} else {

				}
			}
			{
				if result_6_10 {
					__t15 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_15
				} else {

				}
			}
			{
				__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_15:
			_dollar___unused_9_14 := gopurs_runtime.Apply(__t15, gopurs_runtime.Value{})
			_ = _dollar___unused_9_14
			_dollar___unused_8_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_7_11), gopurs_runtime.Bool(result_6_10)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_12
			// TAST (Let): __local_var_9_16 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_9_16 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_9.StrVal()) + ("<")) + (next_10.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal()), gopurs_runtime.Str("Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_9_16
			// TAST (Let): result_10_17 shape=Other bindingType=Boolean
			result_10_17 := (__local_var_9_16.actual) == (gopurs_runtime.Str(__local_var_9_16.expected).StrVal())
			_ = result_10_17
			// TAST (Let): message_11_18 shape=Other bindingType=String
			message_11_18 := ((("foldl - M2\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_9_16.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_9_16.actual)).StrVal())
			_ = message_11_18
			// TAST (Let): __local_var_12_20 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_12_20 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_11_18))
			_ = __local_var_12_20
			var __t22 gopurs_runtime.Value
			{
				if (result_10_17) != (true) {
					__t22 = __local_var_12_20
					goto end_branch_22
				} else {

				}
			}
			{
				if result_10_17 {
					__t22 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_22
				} else {

				}
			}
			{
				__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_22:
			_dollar___unused_13_21 := gopurs_runtime.Apply(__t22, gopurs_runtime.Value{})
			_ = _dollar___unused_13_21
			_dollar___unused_12_19 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_11_18), gopurs_runtime.Bool(result_10_17)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_19
			// TAST (Let): __local_var_13_23 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_13_23 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_13.StrVal()) + ("<")) + (next_14.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal()), gopurs_runtime.Str("Start<a<b<c"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_13_23
			// TAST (Let): result_14_24 shape=Other bindingType=Boolean
			result_14_24 := (__local_var_13_23.actual) == (gopurs_runtime.Str(__local_var_13_23.expected).StrVal())
			_ = result_14_24
			// TAST (Let): message_15_25 shape=Other bindingType=String
			message_15_25 := ((("foldl - M3\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_13_23.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_13_23.actual)).StrVal())
			_ = message_15_25
			// TAST (Let): __local_var_16_27 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_16_27 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_15_25))
			_ = __local_var_16_27
			var __t29 gopurs_runtime.Value
			{
				if (result_14_24) != (true) {
					__t29 = __local_var_16_27
					goto end_branch_29
				} else {

				}
			}
			{
				if result_14_24 {
					__t29 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_29
				} else {

				}
			}
			{
				__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_29:
			_dollar___unused_17_28 := gopurs_runtime.Apply(__t29, gopurs_runtime.Value{})
			_ = _dollar___unused_17_28
			_dollar___unused_16_26 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_15_25), gopurs_runtime.Bool(result_14_24)), gopurs_runtime.Value{})
			_ = _dollar___unused_16_26
			// TAST (Let): __local_var_17_30 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_17_30 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_18 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_17.StrVal()) + ("<")) + (next_18.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
					orig := func() *struct {
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
						return &clone
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
				}()}))}).StrVal()), gopurs_runtime.Str("Start<a<b<c"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_17_30
			// TAST (Let): result_18_31 shape=Other bindingType=Boolean
			result_18_31 := (__local_var_17_30.actual) == (gopurs_runtime.Str(__local_var_17_30.expected).StrVal())
			_ = result_18_31
			// TAST (Let): message_19_32 shape=Other bindingType=String
			message_19_32 := ((("foldl - M4\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_17_30.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_17_30.actual)).StrVal())
			_ = message_19_32
			// TAST (Let): __local_var_20_34 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_20_34 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_19_32))
			_ = __local_var_20_34
			var __t36 gopurs_runtime.Value
			{
				if (result_18_31) != (true) {
					__t36 = __local_var_20_34
					goto end_branch_36
				} else {

				}
			}
			{
				if result_18_31 {
					__t36 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_36
				} else {

				}
			}
			{
				__t36 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_36:
			_dollar___unused_21_35 := gopurs_runtime.Apply(__t36, gopurs_runtime.Value{})
			_ = _dollar___unused_21_35
			_dollar___unused_20_33 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_19_32), gopurs_runtime.Bool(result_18_31)), gopurs_runtime.Value{})
			_ = _dollar___unused_20_33
			// TAST (Let): __local_var_21_37 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_21_37 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_21 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_22 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_21.StrVal()) + ("<")) + (next_22.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal()), gopurs_runtime.Str("Start<a<b<c"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_21_37
			// TAST (Let): result_22_38 shape=Other bindingType=Boolean
			result_22_38 := (__local_var_21_37.actual) == (gopurs_runtime.Str(__local_var_21_37.expected).StrVal())
			_ = result_22_38
			// TAST (Let): message_23_39 shape=Other bindingType=String
			message_23_39 := ((("foldl - M5\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_21_37.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_21_37.actual)).StrVal())
			_ = message_23_39
			// TAST (Let): __local_var_24_41 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_24_41 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_23_39))
			_ = __local_var_24_41
			var __t43 gopurs_runtime.Value
			{
				if (result_22_38) != (true) {
					__t43 = __local_var_24_41
					goto end_branch_43
				} else {

				}
			}
			{
				if result_22_38 {
					__t43 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_43
				} else {

				}
			}
			{
				__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_43:
			_dollar___unused_25_42 := gopurs_runtime.Apply(__t43, gopurs_runtime.Value{})
			_ = _dollar___unused_25_42
			_dollar___unused_24_40 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_23_39), gopurs_runtime.Bool(result_22_38)), gopurs_runtime.Value{})
			_ = _dollar___unused_24_40
			// TAST (Let): __local_var_25_44 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_25_44 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_25 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_26 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_25.StrVal()) + ("<")) + (next_26.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal()), gopurs_runtime.Str("Start<a<b<c<a<b<c<a<b<c"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_25_44
			// TAST (Let): result_26_45 shape=Other bindingType=Boolean
			result_26_45 := (__local_var_25_44.actual) == (gopurs_runtime.Str(__local_var_25_44.expected).StrVal())
			_ = result_26_45
			// TAST (Let): message_27_46 shape=Other bindingType=String
			message_27_46 := ((("foldl - M6\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_25_44.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_25_44.actual)).StrVal())
			_ = message_27_46
			// TAST (Let): __local_var_28_48 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_28_48 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_27_46))
			_ = __local_var_28_48
			var __t50 gopurs_runtime.Value
			{
				if (result_26_45) != (true) {
					__t50 = __local_var_28_48
					goto end_branch_50
				} else {

				}
			}
			{
				if result_26_45 {
					__t50 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_50
				} else {

				}
			}
			{
				__t50 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_50:
			_dollar___unused_29_49 := gopurs_runtime.Apply(__t50, gopurs_runtime.Value{})
			_ = _dollar___unused_29_49
			_dollar___unused_28_47 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_27_46), gopurs_runtime.Bool(result_26_45)), gopurs_runtime.Value{})
			_ = _dollar___unused_28_47
			// TAST (Let): __local_var_29_51 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_29_51 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V1), gopurs_runtime.Func(func(acc_29 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(next_30 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((acc_29.StrVal()) + ("<")) + (next_30.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal()), gopurs_runtime.Str("Start<a<b<c"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_29_51
			// TAST (Let): result_30_52 shape=Other bindingType=Boolean
			result_30_52 := (__local_var_29_51.actual) == (gopurs_runtime.Str(__local_var_29_51.expected).StrVal())
			_ = result_30_52
			// TAST (Let): message_31_53 shape=Other bindingType=String
			message_31_53 := ((("foldl - M7\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_29_51.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_29_51.actual)).StrVal())
			_ = message_31_53
			// TAST (Let): __local_var_32_55 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_32_55 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_31_53))
			_ = __local_var_32_55
			var __t57 gopurs_runtime.Value
			{
				if (result_30_52) != (true) {
					__t57 = __local_var_32_55
					goto end_branch_57
				} else {

				}
			}
			{
				if result_30_52 {
					__t57 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_57
				} else {

				}
			}
			{
				__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_57:
			_dollar___unused_33_56 := gopurs_runtime.Apply(__t57, gopurs_runtime.Value{})
			_ = _dollar___unused_33_56
			_dollar___unused_32_54 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_31_53), gopurs_runtime.Bool(result_30_52)), gopurs_runtime.Value{})
			_ = _dollar___unused_32_54
			// TAST (Let): __local_var_33_58 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_33_58 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_33 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_34 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_33.StrVal()) + (">")) + (acc_34.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}).StrVal()), gopurs_runtime.Str("Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_33_58
			// TAST (Let): result_34_59 shape=Other bindingType=Boolean
			result_34_59 := (__local_var_33_58.actual) == (gopurs_runtime.Str(__local_var_33_58.expected).StrVal())
			_ = result_34_59
			// TAST (Let): message_35_60 shape=Other bindingType=String
			message_35_60 := ((("foldr - M0\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_33_58.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_33_58.actual)).StrVal())
			_ = message_35_60
			// TAST (Let): __local_var_36_62 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_36_62 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_35_60))
			_ = __local_var_36_62
			var __t64 gopurs_runtime.Value
			{
				if (result_34_59) != (true) {
					__t64 = __local_var_36_62
					goto end_branch_64
				} else {

				}
			}
			{
				if result_34_59 {
					__t64 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_64
				} else {

				}
			}
			{
				__t64 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_64:
			_dollar___unused_37_63 := gopurs_runtime.Apply(__t64, gopurs_runtime.Value{})
			_ = _dollar___unused_37_63
			_dollar___unused_36_61 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_35_60), gopurs_runtime.Bool(result_34_59)), gopurs_runtime.Value{})
			_ = _dollar___unused_36_61
			// TAST (Let): __local_var_37_65 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_37_65 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_37 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_38 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_37.StrVal()) + (">")) + (acc_38.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m1()).StrVal()), gopurs_runtime.Str("a>b>c>Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_37_65
			// TAST (Let): result_38_66 shape=Other bindingType=Boolean
			result_38_66 := (__local_var_37_65.actual) == (gopurs_runtime.Str(__local_var_37_65.expected).StrVal())
			_ = result_38_66
			// TAST (Let): message_39_67 shape=Other bindingType=String
			message_39_67 := ((("foldr - M1\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_37_65.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_37_65.actual)).StrVal())
			_ = message_39_67
			// TAST (Let): __local_var_40_69 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_40_69 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_39_67))
			_ = __local_var_40_69
			var __t71 gopurs_runtime.Value
			{
				if (result_38_66) != (true) {
					__t71 = __local_var_40_69
					goto end_branch_71
				} else {

				}
			}
			{
				if result_38_66 {
					__t71 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_71
				} else {

				}
			}
			{
				__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_71:
			_dollar___unused_41_70 := gopurs_runtime.Apply(__t71, gopurs_runtime.Value{})
			_ = _dollar___unused_41_70
			_dollar___unused_40_68 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_39_67), gopurs_runtime.Bool(result_38_66)), gopurs_runtime.Value{})
			_ = _dollar___unused_40_68
			// TAST (Let): __local_var_41_72 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_41_72 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_41 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_42 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_41.StrVal()) + (">")) + (acc_42.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m2()).StrVal()), gopurs_runtime.Str("Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_41_72
			// TAST (Let): result_42_73 shape=Other bindingType=Boolean
			result_42_73 := (__local_var_41_72.actual) == (gopurs_runtime.Str(__local_var_41_72.expected).StrVal())
			_ = result_42_73
			// TAST (Let): message_43_74 shape=Other bindingType=String
			message_43_74 := ((("foldr - M2\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_41_72.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_41_72.actual)).StrVal())
			_ = message_43_74
			// TAST (Let): __local_var_44_76 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_44_76 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_43_74))
			_ = __local_var_44_76
			var __t78 gopurs_runtime.Value
			{
				if (result_42_73) != (true) {
					__t78 = __local_var_44_76
					goto end_branch_78
				} else {

				}
			}
			{
				if result_42_73 {
					__t78 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_78
				} else {

				}
			}
			{
				__t78 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_78:
			_dollar___unused_45_77 := gopurs_runtime.Apply(__t78, gopurs_runtime.Value{})
			_ = _dollar___unused_45_77
			_dollar___unused_44_75 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_43_74), gopurs_runtime.Bool(result_42_73)), gopurs_runtime.Value{})
			_ = _dollar___unused_44_75
			// TAST (Let): __local_var_45_79 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_45_79 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_45 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_46 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_45.StrVal()) + (">")) + (acc_46.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m3()).StrVal()), gopurs_runtime.Str("a>b>c>Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_45_79
			// TAST (Let): result_46_80 shape=Other bindingType=Boolean
			result_46_80 := (__local_var_45_79.actual) == (gopurs_runtime.Str(__local_var_45_79.expected).StrVal())
			_ = result_46_80
			// TAST (Let): message_47_81 shape=Other bindingType=String
			message_47_81 := ((("foldr - M3\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_45_79.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_45_79.actual)).StrVal())
			_ = message_47_81
			// TAST (Let): __local_var_48_83 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_48_83 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_47_81))
			_ = __local_var_48_83
			var __t85 gopurs_runtime.Value
			{
				if (result_46_80) != (true) {
					__t85 = __local_var_48_83
					goto end_branch_85
				} else {

				}
			}
			{
				if result_46_80 {
					__t85 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_85
				} else {

				}
			}
			{
				__t85 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_85:
			_dollar___unused_49_84 := gopurs_runtime.Apply(__t85, gopurs_runtime.Value{})
			_ = _dollar___unused_49_84
			_dollar___unused_48_82 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_47_81), gopurs_runtime.Bool(result_46_80)), gopurs_runtime.Value{})
			_ = _dollar___unused_48_82
			// TAST (Let): __local_var_49_86 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_49_86 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_49 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_50 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_49.StrVal()) + (">")) + (acc_50.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
					orig := func() *struct {
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
						return &clone
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
				}()}))}).StrVal()), gopurs_runtime.Str("a>b>c>Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_49_86
			// TAST (Let): result_50_87 shape=Other bindingType=Boolean
			result_50_87 := (__local_var_49_86.actual) == (gopurs_runtime.Str(__local_var_49_86.expected).StrVal())
			_ = result_50_87
			// TAST (Let): message_51_88 shape=Other bindingType=String
			message_51_88 := ((("foldr - M4\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_49_86.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_49_86.actual)).StrVal())
			_ = message_51_88
			// TAST (Let): __local_var_52_90 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_52_90 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_51_88))
			_ = __local_var_52_90
			var __t92 gopurs_runtime.Value
			{
				if (result_50_87) != (true) {
					__t92 = __local_var_52_90
					goto end_branch_92
				} else {

				}
			}
			{
				if result_50_87 {
					__t92 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_92
				} else {

				}
			}
			{
				__t92 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_92:
			_dollar___unused_53_91 := gopurs_runtime.Apply(__t92, gopurs_runtime.Value{})
			_ = _dollar___unused_53_91
			_dollar___unused_52_89 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_51_88), gopurs_runtime.Bool(result_50_87)), gopurs_runtime.Value{})
			_ = _dollar___unused_52_89
			// TAST (Let): __local_var_53_93 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_53_93 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_53 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_54 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_53.StrVal()) + (">")) + (acc_54.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m5()).StrVal()), gopurs_runtime.Str("a>b>c>Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_53_93
			// TAST (Let): result_54_94 shape=Other bindingType=Boolean
			result_54_94 := (__local_var_53_93.actual) == (gopurs_runtime.Str(__local_var_53_93.expected).StrVal())
			_ = result_54_94
			// TAST (Let): message_55_95 shape=Other bindingType=String
			message_55_95 := ((("foldr - M5\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_53_93.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_53_93.actual)).StrVal())
			_ = message_55_95
			// TAST (Let): __local_var_56_97 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_56_97 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_55_95))
			_ = __local_var_56_97
			var __t99 gopurs_runtime.Value
			{
				if (result_54_94) != (true) {
					__t99 = __local_var_56_97
					goto end_branch_99
				} else {

				}
			}
			{
				if result_54_94 {
					__t99 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_99
				} else {

				}
			}
			{
				__t99 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_99:
			_dollar___unused_57_98 := gopurs_runtime.Apply(__t99, gopurs_runtime.Value{})
			_ = _dollar___unused_57_98
			_dollar___unused_56_96 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_55_95), gopurs_runtime.Bool(result_54_94)), gopurs_runtime.Value{})
			_ = _dollar___unused_56_96
			// TAST (Let): __local_var_57_100 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_57_100 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_57 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_58 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_57.StrVal()) + (">")) + (acc_58.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m6()).StrVal()), gopurs_runtime.Str("a>b>c>a>b>c>a>b>c>Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_57_100
			// TAST (Let): result_58_101 shape=Other bindingType=Boolean
			result_58_101 := (__local_var_57_100.actual) == (gopurs_runtime.Str(__local_var_57_100.expected).StrVal())
			_ = result_58_101
			// TAST (Let): message_59_102 shape=Other bindingType=String
			message_59_102 := ((("foldr - M6\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_57_100.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_57_100.actual)).StrVal())
			_ = message_59_102
			// TAST (Let): __local_var_60_104 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_60_104 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_59_102))
			_ = __local_var_60_104
			var __t106 gopurs_runtime.Value
			{
				if (result_58_101) != (true) {
					__t106 = __local_var_60_104
					goto end_branch_106
				} else {

				}
			}
			{
				if result_58_101 {
					__t106 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_106
				} else {

				}
			}
			{
				__t106 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_106:
			_dollar___unused_61_105 := gopurs_runtime.Apply(__t106, gopurs_runtime.Value{})
			_ = _dollar___unused_61_105
			_dollar___unused_60_103 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_59_102), gopurs_runtime.Bool(result_58_101)), gopurs_runtime.Value{})
			_ = _dollar___unused_60_103
			// TAST (Let): __local_var_61_107 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_61_107 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V2), gopurs_runtime.Func(func(next_61 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(acc_62 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((next_61.StrVal()) + (">")) + (acc_62.StrVal()))
					})
				}), gopurs_runtime.Str("Start"), Get_Main_m7()).StrVal()), gopurs_runtime.Str("a>b>c>Start"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_61_107
			// TAST (Let): result_62_108 shape=Other bindingType=Boolean
			result_62_108 := (__local_var_61_107.actual) == (gopurs_runtime.Str(__local_var_61_107.expected).StrVal())
			_ = result_62_108
			// TAST (Let): message_63_109 shape=Other bindingType=String
			message_63_109 := ((("foldr - M7\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_61_107.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_61_107.actual)).StrVal())
			_ = message_63_109
			// TAST (Let): __local_var_64_111 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_64_111 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_63_109))
			_ = __local_var_64_111
			var __t113 gopurs_runtime.Value
			{
				if (result_62_108) != (true) {
					__t113 = __local_var_64_111
					goto end_branch_113
				} else {

				}
			}
			{
				if result_62_108 {
					__t113 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_113
				} else {

				}
			}
			{
				__t113 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_113:
			_dollar___unused_65_112 := gopurs_runtime.Apply(__t113, gopurs_runtime.Value{})
			_ = _dollar___unused_65_112
			_dollar___unused_64_110 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_63_109), gopurs_runtime.Bool(result_62_108)), gopurs_runtime.Value{})
			_ = _dollar___unused_64_110
			// TAST (Let): __local_var_65_114 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_65_114 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_65 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_65
				}), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}).StrVal()), gopurs_runtime.Str(""))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_65_114
			// TAST (Let): result_66_115 shape=Other bindingType=Boolean
			result_66_115 := (__local_var_65_114.actual) == (gopurs_runtime.Str(__local_var_65_114.expected).StrVal())
			_ = result_66_115
			// TAST (Let): message_67_116 shape=Other bindingType=String
			message_67_116 := ((("foldMap - M0\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_65_114.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_65_114.actual)).StrVal())
			_ = message_67_116
			// TAST (Let): __local_var_68_118 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_68_118 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_67_116))
			_ = __local_var_68_118
			var __t120 gopurs_runtime.Value
			{
				if (result_66_115) != (true) {
					__t120 = __local_var_68_118
					goto end_branch_120
				} else {

				}
			}
			{
				if result_66_115 {
					__t120 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_120
				} else {

				}
			}
			{
				__t120 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_120:
			_dollar___unused_69_119 := gopurs_runtime.Apply(__t120, gopurs_runtime.Value{})
			_ = _dollar___unused_69_119
			_dollar___unused_68_117 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_67_116), gopurs_runtime.Bool(result_66_115)), gopurs_runtime.Value{})
			_ = _dollar___unused_68_117
			// TAST (Let): __local_var_69_121 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_69_121 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_69 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_69
				}), Get_Main_m1()).StrVal()), gopurs_runtime.Str("abc"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_69_121
			// TAST (Let): result_70_122 shape=Other bindingType=Boolean
			result_70_122 := (__local_var_69_121.actual) == (gopurs_runtime.Str(__local_var_69_121.expected).StrVal())
			_ = result_70_122
			// TAST (Let): message_71_123 shape=Other bindingType=String
			message_71_123 := ((("foldMap - M1\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_69_121.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_69_121.actual)).StrVal())
			_ = message_71_123
			// TAST (Let): __local_var_72_125 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_72_125 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_71_123))
			_ = __local_var_72_125
			var __t127 gopurs_runtime.Value
			{
				if (result_70_122) != (true) {
					__t127 = __local_var_72_125
					goto end_branch_127
				} else {

				}
			}
			{
				if result_70_122 {
					__t127 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_127
				} else {

				}
			}
			{
				__t127 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_127:
			_dollar___unused_73_126 := gopurs_runtime.Apply(__t127, gopurs_runtime.Value{})
			_ = _dollar___unused_73_126
			_dollar___unused_72_124 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_71_123), gopurs_runtime.Bool(result_70_122)), gopurs_runtime.Value{})
			_ = _dollar___unused_72_124
			// TAST (Let): __local_var_73_128 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_73_128 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_73 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_73
				}), Get_Main_m2()).StrVal()), gopurs_runtime.Str(""))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_73_128
			// TAST (Let): result_74_129 shape=Other bindingType=Boolean
			result_74_129 := (__local_var_73_128.actual) == (gopurs_runtime.Str(__local_var_73_128.expected).StrVal())
			_ = result_74_129
			// TAST (Let): message_75_130 shape=Other bindingType=String
			message_75_130 := ((("foldMap - M2\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_73_128.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_73_128.actual)).StrVal())
			_ = message_75_130
			// TAST (Let): __local_var_76_132 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_76_132 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_75_130))
			_ = __local_var_76_132
			var __t134 gopurs_runtime.Value
			{
				if (result_74_129) != (true) {
					__t134 = __local_var_76_132
					goto end_branch_134
				} else {

				}
			}
			{
				if result_74_129 {
					__t134 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_134
				} else {

				}
			}
			{
				__t134 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_134:
			_dollar___unused_77_133 := gopurs_runtime.Apply(__t134, gopurs_runtime.Value{})
			_ = _dollar___unused_77_133
			_dollar___unused_76_131 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_75_130), gopurs_runtime.Bool(result_74_129)), gopurs_runtime.Value{})
			_ = _dollar___unused_76_131
			// TAST (Let): __local_var_77_135 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_77_135 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_77 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_77
				}), Get_Main_m3()).StrVal()), gopurs_runtime.Str("abc"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_77_135
			// TAST (Let): result_78_136 shape=Other bindingType=Boolean
			result_78_136 := (__local_var_77_135.actual) == (gopurs_runtime.Str(__local_var_77_135.expected).StrVal())
			_ = result_78_136
			// TAST (Let): message_79_137 shape=Other bindingType=String
			message_79_137 := ((("foldMap - M3\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_77_135.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_77_135.actual)).StrVal())
			_ = message_79_137
			// TAST (Let): __local_var_80_139 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_80_139 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_79_137))
			_ = __local_var_80_139
			var __t141 gopurs_runtime.Value
			{
				if (result_78_136) != (true) {
					__t141 = __local_var_80_139
					goto end_branch_141
				} else {

				}
			}
			{
				if result_78_136 {
					__t141 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_141
				} else {

				}
			}
			{
				__t141 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_141:
			_dollar___unused_81_140 := gopurs_runtime.Apply(__t141, gopurs_runtime.Value{})
			_ = _dollar___unused_81_140
			_dollar___unused_80_138 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_79_137), gopurs_runtime.Bool(result_78_136)), gopurs_runtime.Value{})
			_ = _dollar___unused_80_138
			// TAST (Let): __local_var_81_142 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_81_142 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_81 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_81
				}), gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4[gopurs_runtime.Value, string]{1, func() gopurs_runtime.Value {
					orig := func() *struct {
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
						return &clone
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
				}()}))}).StrVal()), gopurs_runtime.Str("abc"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_81_142
			// TAST (Let): result_82_143 shape=Other bindingType=Boolean
			result_82_143 := (__local_var_81_142.actual) == (gopurs_runtime.Str(__local_var_81_142.expected).StrVal())
			_ = result_82_143
			// TAST (Let): message_83_144 shape=Other bindingType=String
			message_83_144 := ((("foldMap - M4\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_81_142.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_81_142.actual)).StrVal())
			_ = message_83_144
			// TAST (Let): __local_var_84_146 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_84_146 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_83_144))
			_ = __local_var_84_146
			var __t148 gopurs_runtime.Value
			{
				if (result_82_143) != (true) {
					__t148 = __local_var_84_146
					goto end_branch_148
				} else {

				}
			}
			{
				if result_82_143 {
					__t148 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_148
				} else {

				}
			}
			{
				__t148 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_148:
			_dollar___unused_85_147 := gopurs_runtime.Apply(__t148, gopurs_runtime.Value{})
			_ = _dollar___unused_85_147
			_dollar___unused_84_145 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_83_144), gopurs_runtime.Bool(result_82_143)), gopurs_runtime.Value{})
			_ = _dollar___unused_84_145
			// TAST (Let): __local_var_85_149 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_85_149 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_85 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_85
				}), Get_Main_m5()).StrVal()), gopurs_runtime.Str("abc"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_85_149
			// TAST (Let): result_86_150 shape=Other bindingType=Boolean
			result_86_150 := (__local_var_85_149.actual) == (gopurs_runtime.Str(__local_var_85_149.expected).StrVal())
			_ = result_86_150
			// TAST (Let): message_87_151 shape=Other bindingType=String
			message_87_151 := ((("foldMap - M5\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_85_149.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_85_149.actual)).StrVal())
			_ = message_87_151
			// TAST (Let): __local_var_88_153 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_88_153 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_87_151))
			_ = __local_var_88_153
			var __t155 gopurs_runtime.Value
			{
				if (result_86_150) != (true) {
					__t155 = __local_var_88_153
					goto end_branch_155
				} else {

				}
			}
			{
				if result_86_150 {
					__t155 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_155
				} else {

				}
			}
			{
				__t155 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_155:
			_dollar___unused_89_154 := gopurs_runtime.Apply(__t155, gopurs_runtime.Value{})
			_ = _dollar___unused_89_154
			_dollar___unused_88_152 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_87_151), gopurs_runtime.Bool(result_86_150)), gopurs_runtime.Value{})
			_ = _dollar___unused_88_152
			// TAST (Let): __local_var_89_156 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_89_156 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_89 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_89
				}), Get_Main_m6()).StrVal()), gopurs_runtime.Str("abcabcabc"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_89_156
			// TAST (Let): result_90_157 shape=Other bindingType=Boolean
			result_90_157 := (__local_var_89_156.actual) == (gopurs_runtime.Str(__local_var_89_156.expected).StrVal())
			_ = result_90_157
			// TAST (Let): message_91_158 shape=Other bindingType=String
			message_91_158 := ((("foldMap - M6\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_89_156.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_89_156.actual)).StrVal())
			_ = message_91_158
			// TAST (Let): __local_var_92_160 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_92_160 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_91_158))
			_ = __local_var_92_160
			var __t162 gopurs_runtime.Value
			{
				if (result_90_157) != (true) {
					__t162 = __local_var_92_160
					goto end_branch_162
				} else {

				}
			}
			{
				if result_90_157 {
					__t162 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_162
				} else {

				}
			}
			{
				__t162 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_162:
			_dollar___unused_93_161 := gopurs_runtime.Apply(__t162, gopurs_runtime.Value{})
			_ = _dollar___unused_93_161
			_dollar___unused_92_159 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_91_158), gopurs_runtime.Bool(result_90_157)), gopurs_runtime.Value{})
			_ = _dollar___unused_92_159
			// TAST (Let): __local_var_93_163 shape=LitRecord bindingType=(Record (Row [actual: String, expected: String] Any))
			__local_var_93_163 := func() *struct {
				actual   string
				expected string
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableM1()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_93 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_93
				}), Get_Main_m7()).StrVal()), gopurs_runtime.Str("abc"))
				_ = orig
				clone := struct {
					actual   string
					expected string
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").StrVal()
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").StrVal()
				return &clone
			}()
			_ = __local_var_93_163
			// TAST (Let): result_94_164 shape=Other bindingType=Boolean
			result_94_164 := (__local_var_93_163.actual) == (gopurs_runtime.Str(__local_var_93_163.expected).StrVal())
			_ = result_94_164
			// TAST (Let): message_95_165 shape=Other bindingType=String
			message_95_165 := ((("foldMap - M7\x0aExpected: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_93_163.expected)).StrVal())) + ("\x0aActual:   ")) + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__local_var_93_163.actual)).StrVal())
			_ = message_95_165
			// TAST (Let): __local_var_96_167 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_96_167 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_95_165))
			_ = __local_var_96_167
			var __t169 gopurs_runtime.Value
			{
				if (result_94_164) != (true) {
					__t169 = __local_var_96_167
					goto end_branch_169
				} else {

				}
			}
			{
				if result_94_164 {
					__t169 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return Get_Data_Unit_unit()
					})
					goto end_branch_169
				} else {

				}
			}
			{
				__t169 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_169:
			_dollar___unused_97_168 := gopurs_runtime.Apply(__t169, gopurs_runtime.Value{})
			_ = _dollar___unused_97_168
			_dollar___unused_96_166 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_95_165), gopurs_runtime.Bool(result_94_164)), gopurs_runtime.Value{})
			_ = _dollar___unused_96_166
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_M0[T_f any, T_a any] struct {
	Rc uint32
}

type Constructor_Main_M1[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
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
	V0 *struct {
		nested gopurs_runtime.Value
	}
}

type Constructor_Main_M6[T_f any, T_a any] struct {
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

type Constructor_Main_M7[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_identity(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return gopurs_runtime.Str(x_0).StrVal()
}

func Call_Main_foldrStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(next_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		})
	}), gopurs_runtime.Str("Start")).StrVal())
}

func Call_Main_foldrStr__1016676313(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(next_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((next_1.StrVal()) + (">")) + (acc_2.StrVal()))
		})
	}), gopurs_runtime.Str("Start")).StrVal())
}

func Call_Main_foldlStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		})
	}), gopurs_runtime.Str("Start")).StrVal())
}

func Call_Main_foldlStr__1016676313(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(next_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((acc_1.StrVal()) + ("<")) + (next_2.StrVal()))
		})
	}), gopurs_runtime.Str("Start")).StrVal())
}

func Call_Main_foldableM(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_2_0 shape=Other bindingType=(TypeVar m)
		mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
		_ = mempty_2_0
		// TAST (Let): Semigroup0_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
		Semigroup0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_3_1
		return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t9 gopurs_runtime.Value
				{
					if m_5.Type == 9 && m_5.IntVal == 3852365315 {
						__t9 = mempty_2_0
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 769986722 {
						// TAST (Let): Semigroup0_6_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
						Semigroup0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_2
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_2.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1)))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 2727978561 {
						__t9 = mempty_2_0
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 1830062304 {
						__t9 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0)
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 3190619783 {
						// TAST (Let): Semigroup0_6_3 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
						Semigroup0_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_3
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_3.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 108241190 {
						// TAST (Let): Semigroup0_6_4 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
						Semigroup0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_4
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_4.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 2066233029 {
						// TAST (Let): Semigroup0_6_5 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
						Semigroup0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_5
						// TAST (Let): Semigroup0_6_6 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
						Semigroup0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_6
						// TAST (Let): Semigroup0_6_7 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
						Semigroup0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_7
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_5.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V3)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_6.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_7.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))))))))))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 1168316772 {
						__t9 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): Semigroup0_7_8 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
							Semigroup0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_7_8
							return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_8.V0), gopurs_runtime.Apply(f_4, x_8), acc_9)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))
						})), (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0)
						goto end_branch_9
					} else {

					}
				}
				{
					__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_9:
				return __t9
			})
		})
	}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t10 gopurs_runtime.Value
				{
					if m_3.Type == 9 && m_3.IntVal == 3852365315 {
						__t10 = z_2
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 769986722 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2727978561 {
						__t10 = z_2
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1830062304 {
						__t10 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, z_2, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 3190619783 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "fa"))
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 108241190 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "fa"))
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2066233029 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, func() gopurs_runtime.Value {
								arr_val_foldlArray8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray8
								res_go_foldlArray8 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, func() gopurs_runtime.Value {
									arr_val_foldlArray12 := gopurs_runtime.Array((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)
									_ = arr_val_foldlArray12
									res_go_foldlArray12 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)
									_ = res_go_foldlArray12
									arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
									_ = arr_go_foldlArray12
									for _, v_foldlArray12 := range *arr_go_foldlArray12 {
										res_go_foldlArray12 = gopurs_runtime.Apply2(f_1, res_go_foldlArray12, v_foldlArray12)
									}
									return res_go_foldlArray12
								}(), (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V4), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "fa"))
								_ = res_go_foldlArray8
								arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
								_ = arr_go_foldlArray8
								for _, v_foldlArray8 := range *arr_go_foldlArray8 {
									res_go_foldlArray8 = gopurs_runtime.Apply2(f_1, res_go_foldlArray8, v_foldlArray8)
								}
								return res_go_foldlArray8
							}(), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "fa"))
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1168316772 {
						__t10 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return func() gopurs_runtime.Value {
									arr_val_foldlArray9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_foldlArray9
									res_go_foldlArray9 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, v1_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "fa"))
									_ = res_go_foldlArray9
									arr_go_foldlArray9 := (*[]gopurs_runtime.Value)(arr_val_foldlArray9.UnsafePtr)
									_ = arr_go_foldlArray9
									for _, v_foldlArray9 := range *arr_go_foldlArray9 {
										res_go_foldlArray9 = gopurs_runtime.Apply2(f_1, res_go_foldlArray9, v_foldlArray9)
									}
									return res_go_foldlArray9
								}()
							})
						})), z_2, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
						goto end_branch_10
					} else {

					}
				}
				{
					__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_10:
				return __t10
			})
		})
	}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t12 gopurs_runtime.Value
				{
					if m_3.Type == 9 && m_3.IntVal == 3852365315 {
						__t12 = z_2
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 769986722 {
						__t12 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*Constructor_Main_M1[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)))
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2727978561 {
						__t12 = z_2
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1830062304 {
						__t12 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, z_2, (*Constructor_Main_M3[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 3190619783 {
						__t12 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.RecordGet((*Constructor_Main_M4[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, "fa")))
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 108241190 {
						__t12 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.RecordGet((*Constructor_Main_M5[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested, "fa")))
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2066233029 {
						__t12 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V7.nested, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1168316772 {
						// TAST (Let): __local_var_4_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(Record (Row [nested: (Record (Row [a: (TypeVar a), zArrayA: (Array (TypeVar a)), fa: (TypeApp (TypeVar f) [(TypeVar a)]), arrayIgnore: (Array Int), fIgnore: (TypeApp (TypeVar f) [Int]), ignore: Int] Any))] Any))])] (TypeVar b))
						__local_var_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, v2_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "fa")))
							})
						}))
						_ = __local_var_4_11
						__t12 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(__local_var_4_11, a_6, b_5)
							})
						}), z_2, (*Constructor_Main_M7[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
						goto end_branch_12
					} else {

					}
				}
				{
					__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_12:
				return __t12
			})
		})
	})}))}
}

func Call_Main_foldMapStr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	})).StrVal())
}

func Call_Main_foldMapStr__1016676313(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[string]](Get_Data_Monoid_monoidString()))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	})).StrVal())
}
