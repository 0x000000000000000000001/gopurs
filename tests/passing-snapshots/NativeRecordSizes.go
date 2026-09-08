package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502((&Constructor_Data_Show_Show[[]int64]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), Get_Data_Show_showIntImpl())})))}
	})
	return cache_Main_showArray
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_131790935_3790796878((&Constructor_Data_Eq_Eq[[]string]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray1 gopurs_runtime.Value
var once_Main_showArray1 sync.Once

func Get_Main_showArray1() gopurs_runtime.Value {
	once_Main_showArray1.Do(func() {
		cache_Main_showArray1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1953100407_1386611502((&Constructor_Data_Show_Show[[]string]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), Get_Data_Show_showStringImpl())})))}
	})
	return cache_Main_showArray1
}

var cache_Main_shapeZero gopurs_runtime.Value
var once_Main_shapeZero sync.Once

func Get_Main_shapeZero() gopurs_runtime.Value {
	once_Main_shapeZero.Do(func() {
		cache_Main_shapeZero = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_shapeZero(func() struct {
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}()))
		})
	})
	return cache_Main_shapeZero
}

var cache_Main_shapeTwo gopurs_runtime.Value
var once_Main_shapeTwo sync.Once

func Get_Main_shapeTwo() gopurs_runtime.Value {
	once_Main_shapeTwo.Do(func() {
		cache_Main_shapeTwo = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_shapeTwo(func() struct {
				count int64
				label string
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}()))
		})
	})
	return cache_Main_shapeTwo
}

var cache_Main_shapeThree gopurs_runtime.Value
var once_Main_shapeThree sync.Once

func Get_Main_shapeThree() gopurs_runtime.Value {
	once_Main_shapeThree.Do(func() {
		cache_Main_shapeThree = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_shapeThree(func() struct {
				count  int64
				flag   bool
				number float64
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
					count  int64
					flag   bool
					number float64
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
				clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
				return clone
			}()))
		})
	})
	return cache_Main_shapeThree
}

var cache_Main_shapeSix gopurs_runtime.Value
var once_Main_shapeSix sync.Once

func Get_Main_shapeSix() gopurs_runtime.Value {
	once_Main_shapeSix.Do(func() {
		cache_Main_shapeSix = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_shapeSix(func() struct {
				child struct {
					count int64
					label string
				}
				count  int64
				flag   bool
				label  string
				number float64
				values []int64
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
					child struct {
						count int64
						label string
					}
					count  int64
					flag   bool
					label  string
					number float64
					values []int64
				}{}
				clone.child = func() struct {
					count int64
					label string
				} {
					orig := gopurs_runtime.RecordGet(orig, "child")
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
				clone.values = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}()))
		})
	})
	return cache_Main_shapeSix
}

var cache_Main_shapeOne gopurs_runtime.Value
var once_Main_shapeOne sync.Once

func Get_Main_shapeOne() gopurs_runtime.Value {
	once_Main_shapeOne.Do(func() {
		cache_Main_shapeOne = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_shapeOne(func() struct {
				count int64
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
					count int64
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				return clone
			}()))
		})
	})
	return cache_Main_shapeOne
}

var cache_Main_shapeFour gopurs_runtime.Value
var once_Main_shapeFour sync.Once

func Get_Main_shapeFour() gopurs_runtime.Value {
	once_Main_shapeFour.Do(func() {
		cache_Main_shapeFour = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_shapeFour(func() struct {
				child struct {
					count int64
					label string
				}
				count  int64
				label  string
				values []int64
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
					child struct {
						count int64
						label string
					}
					count  int64
					label  string
					values []int64
				}{}
				clone.child = func() struct {
					count int64
					label string
				} {
					orig := gopurs_runtime.RecordGet(orig, "child")
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.values = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}()))
		})
	})
	return cache_Main_shapeFour
}

var cache_Main_shapeFive gopurs_runtime.Value
var once_Main_shapeFive sync.Once

func Get_Main_shapeFive() gopurs_runtime.Value {
	once_Main_shapeFive.Do(func() {
		cache_Main_shapeFive = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_shapeFive(func() struct {
				child struct {
					count int64
					label string
				}
				count   int64
				flag    bool
				label   string
				payload gopurs_runtime.Value
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
					child struct {
						count int64
						label string
					}
					count   int64
					flag    bool
					label   string
					payload gopurs_runtime.Value
				}{}
				clone.child = func() struct {
					count int64
					label string
				} {
					orig := gopurs_runtime.RecordGet(orig, "child")
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.payload = gopurs_runtime.RecordGet(orig, "payload")
				return clone
			}()))
		})
	})
	return cache_Main_shapeFive
}

var cache_Main_makeFour gopurs_runtime.Value
var once_Main_makeFour sync.Once

func Get_Main_makeFour() gopurs_runtime.Value {
	once_Main_makeFour.Do(func() {
		cache_Main_makeFour = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeFour(count_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()
		})
	})
	return cache_Main_makeFour
}

var cache_Main_produceFour gopurs_runtime.Value
var once_Main_produceFour sync.Once

func Get_Main_produceFour() gopurs_runtime.Value {
	once_Main_produceFour.Do(func() {
		cache_Main_produceFour = gopurs_runtime.Func2(func(trace_0_box gopurs_runtime.Value, count_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_produceFour(trace_0_box, count_1_box.IntVal)
		})
	})
	return cache_Main_produceFour
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, actual_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(label_0_box.StrVal(), expected_1_box.StrVal(), actual_2_box.StrVal())
		})
	})
	return cache_Main_check
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(5)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Get_Main_opaqueValue(), gopurs_runtime.Value{})
			_ = __local_var_3_3
			oneRef_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					count int64
				}{__local_var_2_2.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("count", gopurs_runtime.Int(orig.count))
			}()), gopurs_runtime.Value{})
			_ = oneRef_4_4
			twoRef_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					count int64
					label string
				}{__local_var_2_2.IntVal, "alpha"}
				_ = orig
				return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
			}()), gopurs_runtime.Value{})
			_ = twoRef_5_5
			threeRef_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					count  int64
					flag   bool
					number float64
				}{__local_var_2_2.IntVal, true, -1.25}
				_ = orig
				return gopurs_runtime.RecordDict3("count", "flag", "number", gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Float(orig.number))
			}()), gopurs_runtime.Value{})
			_ = threeRef_6_6
			fourRef_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := Call_Main_makeFour(__local_var_2_2.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()), gopurs_runtime.Value{})
			_ = fourRef_7_7
			fiveRef_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					child struct {
						count int64
						label string
					}
					count   int64
					flag    bool
					label   string
					payload gopurs_runtime.Value
				}{struct {
					count int64
					label string
				}{int64(2), "nested"}, __local_var_2_2.IntVal, true, "five", __local_var_3_3}
				_ = orig
				return gopurs_runtime.RecordDict5("child", "count", "flag", "label", "payload", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), orig.payload)
			}()), gopurs_runtime.Value{})
			_ = fiveRef_8_8
			sixRef_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					child struct {
						count int64
						label string
					}
					count  int64
					flag   bool
					label  string
					number float64
					values []int64
				}{struct {
					count int64
					label string
				}{int64(2), "nested"}, __local_var_2_2.IntVal, true, "six", -1.25, []int64{int64(-3), int64(0), int64(7)}}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"child", "count", "flag", "label", "number", "values"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), gopurs_runtime.Float(orig.number), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}()), gopurs_runtime.Value{})
			_ = sixRef_9_9
			one_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), oneRef_4_4), gopurs_runtime.Value{})
			_ = one_10_10
			two_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), twoRef_5_5), gopurs_runtime.Value{})
			_ = two_11_11
			three_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), threeRef_6_6), gopurs_runtime.Value{})
			_ = three_12_12
			four_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), fourRef_7_7), gopurs_runtime.Value{})
			_ = four_13_13
			five_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), fiveRef_8_8), gopurs_runtime.Value{})
			_ = five_14_14
			six_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), sixRef_9_9), gopurs_runtime.Value{})
			_ = six_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Main_check("shape 0", "compact0:", Call_Main_shapeZero(struct {
			}{})), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Main_check("shape 1", "compact1:count", Call_Main_shapeOne(func() struct {
				count int64
			} {
				orig := one_10_10
				_ = orig
				clone := struct {
					count int64
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Main_check("shape 2", "compact2:count|label", Call_Main_shapeTwo(func() struct {
				count int64
				label string
			} {
				orig := two_11_11
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_check("shape 3", "compact3:count|flag|number", Call_Main_shapeThree(func() struct {
				count  int64
				flag   bool
				number float64
			} {
				orig := three_12_12
				_ = orig
				clone := struct {
					count  int64
					flag   bool
					number float64
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
				clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Main_check("shape 4", "compact4:child|count|label|values", Call_Main_shapeFour(func() struct {
				child struct {
					count int64
					label string
				}
				count  int64
				label  string
				values []int64
			} {
				orig := four_13_13
				_ = orig
				clone := struct {
					child struct {
						count int64
						label string
					}
					count  int64
					label  string
					values []int64
				}{}
				clone.child = func() struct {
					count int64
					label string
				} {
					orig := gopurs_runtime.RecordGet(orig, "child")
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.values = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Main_check("shape 5", "compact5:child|count|flag|label|payload", Call_Main_shapeFive(func() struct {
				child struct {
					count int64
					label string
				}
				count   int64
				flag    bool
				label   string
				payload gopurs_runtime.Value
			} {
				orig := five_14_14
				_ = orig
				clone := struct {
					child struct {
						count int64
						label string
					}
					count   int64
					flag    bool
					label   string
					payload gopurs_runtime.Value
				}{}
				clone.child = func() struct {
					count int64
					label string
				} {
					orig := gopurs_runtime.RecordGet(orig, "child")
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.payload = gopurs_runtime.RecordGet(orig, "payload")
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(Call_Main_check("shape 6", "generic6:child|count|flag|label|number|values", Call_Main_shapeSix(func() struct {
				child struct {
					count int64
					label string
				}
				count  int64
				flag   bool
				label  string
				number float64
				values []int64
			} {
				orig := six_15_15
				_ = orig
				clone := struct {
					child struct {
						count int64
						label string
					}
					count  int64
					flag   bool
					label  string
					number float64
					values []int64
				}{}
				clone.child = func() struct {
					count int64
					label string
				} {
					orig := gopurs_runtime.RecordGet(orig, "child")
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
				clone.values = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(Call_Main_check("map 0", "{}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := struct {
				}{}
				_ = orig
				return gopurs_runtime.RecordDict0()
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(Call_Main_check("map 1", "{count=i:5}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count int64
				} {
					orig := one_10_10
					_ = orig
					clone := struct {
						count int64
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict1("count", gopurs_runtime.Int(orig.count))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply(Call_Main_check("map 2", "{count=i:5,label=s:alpha}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count int64
					label string
				} {
					orig := two_11_11
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply(Call_Main_check("map 3", "{count=i:5,flag=b:true,number=n:-1.25}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count  int64
					flag   bool
					number float64
				} {
					orig := three_12_12
					_ = orig
					clone := struct {
						count  int64
						flag   bool
						number float64
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict3("count", "flag", "number", gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Float(orig.number))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_26_26
			__local_var_27_27 := gopurs_runtime.Apply(Call_Main_check("map 4", "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					label  string
					values []int64
				} {
					orig := four_13_13
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						label  string
						values []int64
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_27_27
			__local_var_28_28 := gopurs_runtime.Apply(Call_Main_check("map 5", "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:five,payload=s:opaque}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count   int64
					flag    bool
					label   string
					payload gopurs_runtime.Value
				} {
					orig := five_14_14
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count   int64
						flag    bool
						label   string
						payload gopurs_runtime.Value
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.payload = gopurs_runtime.RecordGet(orig, "payload")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict5("child", "count", "flag", "label", "payload", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), orig.payload)
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_28_28
			__local_var_29_29 := gopurs_runtime.Apply(Call_Main_check("map 6", "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:six,number=n:-1.25,values=[i:-3,i:0,i:7]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					flag   bool
					label  string
					number float64
					values []int64
				} {
					orig := six_15_15
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						flag   bool
						label  string
						number float64
						values []int64
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"child", "count", "flag", "label", "number", "values"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), gopurs_runtime.Float(orig.number), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_29_29
			__local_var_30_30 := gopurs_runtime.Apply(Call_Main_check("update 1", "{count=i:-7}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count int64
				} {
					orig := gopurs_runtime.RecordUpdate1(one_10_10, "count", gopurs_runtime.Int(int64(-7)))
					_ = orig
					clone := struct {
						count int64
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict1("count", gopurs_runtime.Int(orig.count))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_30_30
			__local_var_31_31 := gopurs_runtime.Apply(Call_Main_check("update 2", "{count=i:5,label=s:}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count int64
					label string
				} {
					orig := gopurs_runtime.RecordUpdate1(two_11_11, "label", gopurs_runtime.Str(""))
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_31_31
			__local_var_32_32 := gopurs_runtime.Apply(Call_Main_check("update 3", "{count=i:5,flag=b:false,number=n:2.5}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count  int64
					flag   bool
					number float64
				} {
					orig := gopurs_runtime.RecordUpdate2(three_12_12, "flag", gopurs_runtime.Bool(false), "number", gopurs_runtime.Float(2.5))
					_ = orig
					clone := struct {
						count  int64
						flag   bool
						number float64
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict3("count", "flag", "number", gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Float(orig.number))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_32_32
			__local_var_33_33 := gopurs_runtime.Apply(Call_Main_check("update 4", "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					label  string
					values []gopurs_runtime.Value
				} {
					orig := gopurs_runtime.RecordUpdate1(four_13_13, "values", gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						label  string
						values []gopurs_runtime.Value
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.values = func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), gopurs_runtime.Array(orig.values))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_33_33
			__local_var_34_34 := gopurs_runtime.Apply(Call_Main_check("update 5", "{child={count=i:2,label=s:nested},count=i:5,flag=b:false,label=s:five,payload=s:opaque}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count   int64
					flag    bool
					label   string
					payload gopurs_runtime.Value
				} {
					orig := gopurs_runtime.RecordUpdate1(five_14_14, "flag", gopurs_runtime.Bool(false))
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count   int64
						flag    bool
						label   string
						payload gopurs_runtime.Value
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.payload = gopurs_runtime.RecordGet(orig, "payload")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict5("child", "count", "flag", "label", "payload", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), orig.payload)
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_34_34
			__local_var_35_35 := gopurs_runtime.Apply(Call_Main_check("update 6", "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:,number=n:2.5,values=[i:-3,i:0,i:7]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					flag   bool
					label  string
					number float64
					values []int64
				} {
					orig := gopurs_runtime.RecordUpdate2(six_15_15, "label", gopurs_runtime.Str(""), "number", gopurs_runtime.Float(2.5))
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						flag   bool
						label  string
						number float64
						values []int64
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"child", "count", "flag", "label", "number", "values"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), gopurs_runtime.Float(orig.number), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_35_35
			savedOne_36_36 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), oneRef_4_4), gopurs_runtime.Value{})
			_ = savedOne_36_36
			savedTwo_37_37 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), twoRef_5_5), gopurs_runtime.Value{})
			_ = savedTwo_37_37
			savedThree_38_38 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), threeRef_6_6), gopurs_runtime.Value{})
			_ = savedThree_38_38
			savedFour_39_39 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), fourRef_7_7), gopurs_runtime.Value{})
			_ = savedFour_39_39
			savedFive_40_40 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), fiveRef_8_8), gopurs_runtime.Value{})
			_ = savedFive_40_40
			savedSix_41_41 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), sixRef_9_9), gopurs_runtime.Value{})
			_ = savedSix_41_41
			__local_var_42_42 := gopurs_runtime.Apply(Call_Main_check("retained 1", "{count=i:5}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count int64
				} {
					orig := savedOne_36_36
					_ = orig
					clone := struct {
						count int64
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict1("count", gopurs_runtime.Int(orig.count))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_42_42
			__local_var_43_43 := gopurs_runtime.Apply(Call_Main_check("retained 2", "{count=i:5,label=s:alpha}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count int64
					label string
				} {
					orig := savedTwo_37_37
					_ = orig
					clone := struct {
						count int64
						label string
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_43_43
			__local_var_44_44 := gopurs_runtime.Apply(Call_Main_check("retained 3", "{count=i:5,flag=b:true,number=n:-1.25}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					count  int64
					flag   bool
					number float64
				} {
					orig := savedThree_38_38
					_ = orig
					clone := struct {
						count  int64
						flag   bool
						number float64
					}{}
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict3("count", "flag", "number", gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Float(orig.number))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_44_44
			__local_var_45_45 := gopurs_runtime.Apply(Call_Main_check("retained 4", "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					label  string
					values []int64
				} {
					orig := savedFour_39_39
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						label  string
						values []int64
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_45_45
			__local_var_46_46 := gopurs_runtime.Apply(Call_Main_check("retained 5", "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:five,payload=s:opaque}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count   int64
					flag    bool
					label   string
					payload gopurs_runtime.Value
				} {
					orig := savedFive_40_40
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count   int64
						flag    bool
						label   string
						payload gopurs_runtime.Value
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.payload = gopurs_runtime.RecordGet(orig, "payload")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict5("child", "count", "flag", "label", "payload", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), orig.payload)
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_46_46
			__local_var_47_47 := gopurs_runtime.Apply(Call_Main_check("retained 6", "{child={count=i:2,label=s:nested},count=i:5,flag=b:true,label=s:six,number=n:-1.25,values=[i:-3,i:0,i:7]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					flag   bool
					label  string
					number float64
					values []int64
				} {
					orig := savedSix_41_41
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						flag   bool
						label  string
						number float64
						values []int64
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"child", "count", "flag", "label", "number", "values"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), gopurs_runtime.Float(orig.number), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_47_47
			var __t49 string
			{
				if (gopurs_runtime.RecordGet(three_12_12, "flag").IntVal) != (0) {
					__t49 = ((((("true:") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.RecordGet(three_12_12, "number")).StrVal())) + (":")) + (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(four_13_13, "child"), "label").StrVal())) + (":")) + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray()).V0), gopurs_runtime.RecordGet(four_13_13, "values")).StrVal())
					goto end_branch_49
				} else {

				}
			}
			{
				__t49 = ((((("false:") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.RecordGet(three_12_12, "number")).StrVal())) + (":")) + (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(four_13_13, "child"), "label").StrVal())) + (":")) + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray()).V0), gopurs_runtime.RecordGet(four_13_13, "values")).StrVal())
			}
		end_branch_49:
			__local_var_48_48 := gopurs_runtime.Apply(Call_Main_check("native field access", "5:true:-1.25:nested:[-3,0,7]", ((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(three_12_12, "count")).StrVal())+(":"))+(__t49)), gopurs_runtime.Value{})
			_ = __local_var_48_48
			__local_var_49_50 := gopurs_runtime.Apply(Call_Main_check("nested update", "{child={count=i:2,label=s:changed},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					label  string
					values []int64
				} {
					orig := gopurs_runtime.RecordUpdate1(four_13_13, "child", func() gopurs_runtime.Value {
						orig := func() struct {
							count int64
							label string
						} {
							orig := gopurs_runtime.RecordUpdate1(gopurs_runtime.RecordGet(four_13_13, "child"), "label", gopurs_runtime.Str("changed"))
							_ = orig
							clone := struct {
								count int64
								label string
							}{}
							clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
							clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
					}())
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						label  string
						values []int64
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_49_50
			__local_var_50_51 := gopurs_runtime.Apply(Call_Main_check("nested original", "{child={count=i:2,label=s:nested},count=i:5,label=s:four,values=[i:-3,i:0,i:7]}", gopurs_runtime.Apply(Get_Main_nativeMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					count  int64
					label  string
					values []int64
				} {
					orig := four_13_13
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
						count  int64
						label  string
						values []int64
					}{}
					clone.child = func() struct {
						count int64
						label string
					} {
						orig := gopurs_runtime.RecordGet(orig, "child")
						_ = orig
						clone := struct {
							count int64
							label string
						}{}
						clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
					clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_50_51
			__local_var_51_52 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Array([]gopurs_runtime.Value{})), gopurs_runtime.Value{})
			_ = __local_var_51_52
			__local_var_52_53 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_52 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_52, func() gopurs_runtime.Value {
								arr := []string{"before"}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Str(v)
								}
								return gopurs_runtime.Array(boxed)
							}()).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
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
			}), __local_var_51_52), gopurs_runtime.Value{})
			_ = __local_var_52_53
			produced_53_54 := gopurs_runtime.Apply(Call_Main_produceFour(__local_var_51_52, __local_var_2_2.IntVal), gopurs_runtime.Value{})
			_ = produced_53_54
			__local_var_54_55 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_54 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_54, func() gopurs_runtime.Value {
								arr := []string{("consume:") + (Call_Main_shapeFour(func() struct {
									child struct {
										count int64
										label string
									}
									count  int64
									label  string
									values []int64
								} {
									orig := produced_53_54
									_ = orig
									clone := struct {
										child struct {
											count int64
											label string
										}
										count  int64
										label  string
										values []int64
									}{}
									clone.child = func() struct {
										count int64
										label string
									} {
										orig := gopurs_runtime.RecordGet(orig, "child")
										_ = orig
										clone := struct {
											count int64
											label string
										}{}
										clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
										clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
										return clone
									}()
									clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
									clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
									clone.values = func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
										unboxed := make([]int64, len(arr))
										for i, v := range arr {
											unboxed[i] = v.IntVal
										}
										return unboxed
									}()
									return clone
								}()))}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Str(v)
								}
								return gopurs_runtime.Array(boxed)
							}()).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
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
			}), __local_var_51_52), gopurs_runtime.Value{})
			_ = __local_var_54_55
			__local_var_55_56 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_55 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_55, func() gopurs_runtime.Value {
								arr := []string{"after"}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Str(v)
								}
								return gopurs_runtime.Array(boxed)
							}()).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()).UnsafePtr)
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
			}), __local_var_51_52), gopurs_runtime.Value{})
			_ = __local_var_55_56
			__local_var_56_57 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_51_52), gopurs_runtime.Value{})
			_ = __local_var_56_57
			__local_var_57_58 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_131790935_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]string]](Get_Main_eqArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1953100407_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]string]](Get_Main_showArray1())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
				arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(__local_var_56_57.UnsafePtr)
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
			}(), func() gopurs_runtime.Value {
				arr := []string{"before", "produce", "consume:compact4:child|count|label|values", "after"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}())), gopurs_runtime.Value{})
			_ = __local_var_57_58
			__local_var_58_59 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("effect order: ")+(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]string]](Get_Main_showArray1()).V0), __local_var_56_57).StrVal()))), gopurs_runtime.Value{})
			_ = __local_var_58_59
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_shapeZero(row_0_loop struct {
}) string {
	var row_0 struct {
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.Apply(Get_Main_nativeShape(), func() gopurs_runtime.Value {
		orig := row_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()).StrVal()
}

func Call_Main_shapeTwo(row_0_loop struct {
	count int64
	label string
}) string {
	var row_0 struct {
		count int64
		label string
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.Apply(Get_Main_nativeShape(), func() gopurs_runtime.Value {
		orig := row_0
		_ = orig
		return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
	}()).StrVal()
}

func Call_Main_shapeThree(row_0_loop struct {
	count  int64
	flag   bool
	number float64
}) string {
	var row_0 struct {
		count  int64
		flag   bool
		number float64
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.Apply(Get_Main_nativeShape(), func() gopurs_runtime.Value {
		orig := row_0
		_ = orig
		return gopurs_runtime.RecordDict3("count", "flag", "number", gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Float(orig.number))
	}()).StrVal()
}

func Call_Main_shapeSix(row_0_loop struct {
	child struct {
		count int64
		label string
	}
	count  int64
	flag   bool
	label  string
	number float64
	values []int64
}) string {
	var row_0 struct {
		child struct {
			count int64
			label string
		}
		count  int64
		flag   bool
		label  string
		number float64
		values []int64
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.Apply(Get_Main_nativeShape(), func() gopurs_runtime.Value {
		orig := row_0
		_ = orig
		return gopurs_runtime.RecordDict([]string{"child", "count", "flag", "label", "number", "values"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
			orig := orig.child
			_ = orig
			return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
		}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), gopurs_runtime.Float(orig.number), func() gopurs_runtime.Value {
			arr := orig.values
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()})
	}()).StrVal()
}

func Call_Main_shapeOne(row_0_loop struct {
	count int64
}) string {
	var row_0 struct {
		count int64
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.Apply(Get_Main_nativeShape(), func() gopurs_runtime.Value {
		orig := row_0
		_ = orig
		return gopurs_runtime.RecordDict1("count", gopurs_runtime.Int(orig.count))
	}()).StrVal()
}

func Call_Main_shapeFour(row_0_loop struct {
	child struct {
		count int64
		label string
	}
	count  int64
	label  string
	values []int64
}) string {
	var row_0 struct {
		child struct {
			count int64
			label string
		}
		count  int64
		label  string
		values []int64
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.Apply(Get_Main_nativeShape(), func() gopurs_runtime.Value {
		orig := row_0
		_ = orig
		return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
			orig := orig.child
			_ = orig
			return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
		}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
			arr := orig.values
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}())
	}()).StrVal()
}

func Call_Main_shapeFive(row_0_loop struct {
	child struct {
		count int64
		label string
	}
	count   int64
	flag    bool
	label   string
	payload gopurs_runtime.Value
}) string {
	var row_0 struct {
		child struct {
			count int64
			label string
		}
		count   int64
		flag    bool
		label   string
		payload gopurs_runtime.Value
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.Apply(Get_Main_nativeShape(), func() gopurs_runtime.Value {
		orig := row_0
		_ = orig
		return gopurs_runtime.RecordDict5("child", "count", "flag", "label", "payload", func() gopurs_runtime.Value {
			orig := orig.child
			_ = orig
			return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
		}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Bool(orig.flag), gopurs_runtime.Str(orig.label), orig.payload)
	}()).StrVal()
}

func Call_Main_makeFour(count_0_loop int64) struct {
	child struct {
		count int64
		label string
	}
	count  int64
	label  string
	values []int64
} {
	var count_0 int64 = count_0_loop
	_ = count_0
	return struct {
		child struct {
			count int64
			label string
		}
		count  int64
		label  string
		values []int64
	}{struct {
		count int64
		label string
	}{int64(2), "nested"}, count_0, "four", []int64{int64(-3), int64(0), int64(7)}}
}

func Call_Main_produceFour(trace_0_loop gopurs_runtime.Value, count_1_loop int64) gopurs_runtime.Value {
	var trace_0 gopurs_runtime.Value = trace_0_loop
	_ = trace_0
	var count_1 int64 = count_1_loop
	_ = count_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_2, func() gopurs_runtime.Value {
							arr := []string{"produce"}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Str(v)
							}
							return gopurs_runtime.Array(boxed)
						}()).UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()).UnsafePtr)
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
		}), trace_0)
		_ = __local_var_2_0
		_dollar___unused_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
		_ = _dollar___unused_3_1
		return func() gopurs_runtime.Value {
			orig := Call_Main_makeFour(count_1)
			_ = orig
			return gopurs_runtime.RecordDict4("child", "count", "label", "values", func() gopurs_runtime.Value {
				orig := orig.child
				_ = orig
				return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
			}(), gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
				arr := orig.values
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())
		}()
	})
}

func Call_Main_check(label_0_loop string, expected_1_loop string, actual_2_loop string) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 string = expected_1_loop
	_ = expected_1
	var actual_2 string = actual_2_loop
	_ = actual_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(actual_2), gopurs_runtime.Str(expected_1)))
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(((label_0)+(": "))+(actual_2))), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
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

func Rebox_Main_1953100407_1386611502(in *Constructor_Data_Show_Show[[]string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Get_Main_nativeMap() gopurs_runtime.Value {
	return _Gopurs_Main_NativeMap
}

func Get_Main_nativeShape() gopurs_runtime.Value {
	return _Gopurs_Main_NativeShape
}

func Get_Main_opaqueValue() gopurs_runtime.Value {
	return _Gopurs_Main_OpaqueValue
}
