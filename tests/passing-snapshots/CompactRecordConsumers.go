package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_showArray1 gopurs_runtime.Value
var once_Main_showArray1 sync.Once

func Get_Main_showArray1() gopurs_runtime.Value {
	once_Main_showArray1.Do(func() {
		cache_Main_showArray1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showArray1
}

var cache_Main_makeEntry gopurs_runtime.Value
var once_Main_makeEntry sync.Once

func Get_Main_makeEntry() gopurs_runtime.Value {
	once_Main_makeEntry.Do(func() {
		cache_Main_makeEntry = gopurs_runtime.Func2(func(count_0_box gopurs_runtime.Value, label_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeEntry(count_0_box.IntVal, label_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
			}()
		})
	})
	return cache_Main_makeEntry
}

var cache_Main_produceEntry gopurs_runtime.Value
var once_Main_produceEntry sync.Once

func Get_Main_produceEntry() gopurs_runtime.Value {
	once_Main_produceEntry.Do(func() {
		cache_Main_produceEntry = gopurs_runtime.Func(func(trace_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_produceEntry(trace_0_box)
		})
	})
	return cache_Main_produceEntry
}

var cache_Main_inspectProduced gopurs_runtime.Value
var once_Main_inspectProduced sync.Once

func Get_Main_inspectProduced() gopurs_runtime.Value {
	once_Main_inspectProduced.Do(func() {
		cache_Main_inspectProduced = gopurs_runtime.Func2(func(inspect_0_box gopurs_runtime.Value, produce_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inspectProduced(inspect_0_box, produce_1_box)
		})
	})
	return cache_Main_inspectProduced
}

var cache_Main_inspectProduced__354691939 gopurs_runtime.Value
var once_Main_inspectProduced__354691939 sync.Once

func Get_Main_inspectProduced__354691939() gopurs_runtime.Value {
	once_Main_inspectProduced__354691939.Do(func() {
		cache_Main_inspectProduced__354691939 = gopurs_runtime.Func2(func(inspect_unused_0_box gopurs_runtime.Value, produce_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inspectProduced__354691939(inspect_unused_0_box, produce_unused_1_box)
		})
	})
	return cache_Main_inspectProduced__354691939
}

var cache_Main_inspectProduced__1003637247 gopurs_runtime.Value
var once_Main_inspectProduced__1003637247 sync.Once

func Get_Main_inspectProduced__1003637247() gopurs_runtime.Value {
	once_Main_inspectProduced__1003637247.Do(func() {
		cache_Main_inspectProduced__1003637247 = gopurs_runtime.Func2(func(inspect_unused_0_box gopurs_runtime.Value, produce_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inspectProduced__1003637247(inspect_unused_0_box, produce_unused_1_box)
		})
	})
	return cache_Main_inspectProduced__1003637247
}

var cache_Main_inspectProduced__2901793856 gopurs_runtime.Value
var once_Main_inspectProduced__2901793856 sync.Once

func Get_Main_inspectProduced__2901793856() gopurs_runtime.Value {
	once_Main_inspectProduced__2901793856.Do(func() {
		cache_Main_inspectProduced__2901793856 = gopurs_runtime.Func2(func(inspect_unused_0_box gopurs_runtime.Value, produce_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inspectProduced__2901793856(inspect_unused_0_box, produce_1_box)
		})
	})
	return cache_Main_inspectProduced__2901793856
}

var cache_Main_inspectProduced__1824334758 gopurs_runtime.Value
var once_Main_inspectProduced__1824334758 sync.Once

func Get_Main_inspectProduced__1824334758() gopurs_runtime.Value {
	once_Main_inspectProduced__1824334758.Do(func() {
		cache_Main_inspectProduced__1824334758 = gopurs_runtime.Func2(func(inspect_unused_0_box gopurs_runtime.Value, produce_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inspectProduced__1824334758(inspect_unused_0_box, produce_1_box)
		})
	})
	return cache_Main_inspectProduced__1824334758
}

var cache_Main_inspectProduced__3949581232 gopurs_runtime.Value
var once_Main_inspectProduced__3949581232 sync.Once

func Get_Main_inspectProduced__3949581232() gopurs_runtime.Value {
	once_Main_inspectProduced__3949581232.Do(func() {
		cache_Main_inspectProduced__3949581232 = gopurs_runtime.Func2(func(inspect_unused_0_box gopurs_runtime.Value, produce_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inspectProduced__3949581232(inspect_unused_0_box, produce_unused_1_box)
		})
	})
	return cache_Main_inspectProduced__3949581232
}

var cache_Main_inspectProduced__1784222911 gopurs_runtime.Value
var once_Main_inspectProduced__1784222911 sync.Once

func Get_Main_inspectProduced__1784222911() gopurs_runtime.Value {
	once_Main_inspectProduced__1784222911.Do(func() {
		cache_Main_inspectProduced__1784222911 = gopurs_runtime.Func2(func(inspect_unused_0_box gopurs_runtime.Value, produce_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inspectProduced__1784222911(inspect_unused_0_box, produce_unused_1_box)
		})
	})
	return cache_Main_inspectProduced__1784222911
}

var cache_Main_checkEffects gopurs_runtime.Value
var once_Main_checkEffects sync.Once

func Get_Main_checkEffects() gopurs_runtime.Value {
	once_Main_checkEffects.Do(func() {
		cache_Main_checkEffects = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
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
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func(func(entry_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr := func() []string {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_3, func() gopurs_runtime.Value {
								arr := []string{("consume:") + (gopurs_runtime.Apply(Get_Main_describeEntryMap(), entry_2).StrVal())}
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
				}), __local_var_1_1)
			})), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_2), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_4, func() gopurs_runtime.Value {
							arr := []string{"before"}
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
			}), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Main_produceEntry(__local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_3, __local_var_5_5), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_7, func() gopurs_runtime.Value {
							arr := []string{"after"}
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
			}), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2499841393("", struct {
				actual   []string
				expected []string
			}{func() []string {
				arr := *(*[]gopurs_runtime.Value)(__local_var_8_8.UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}(), []string{"before", "produce", "consume:9:created", "after"}}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("effect order: ")+(gopurs_runtime.Apply(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showArray())).V0, __local_var_8_8).StrVal()))), gopurs_runtime.Value{})
		})
	})
	return cache_Main_checkEffects
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

var cache_Main_checkEntry gopurs_runtime.Value
var once_Main_checkEntry sync.Once

func Get_Main_checkEntry() gopurs_runtime.Value {
	once_Main_checkEntry.Do(func() {
		cache_Main_checkEntry = gopurs_runtime.Func(func(reverse_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkEntry((reverse_0_box.IntVal) != (0))
		})
	})
	return cache_Main_checkEntry
}

var cache_Main_checkPayload gopurs_runtime.Value
var once_Main_checkPayload sync.Once

func Get_Main_checkPayload() gopurs_runtime.Value {
	once_Main_checkPayload.Do(func() {
		cache_Main_checkPayload = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Get_Main_compactPayload(), gopurs_runtime.Value{})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), __local_var_0_0), gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Get_Main_compactPayload(), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Get_Main_compactKeys(), __local_var_2_2)
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Main_check("payload compact keys", "values|child", __local_var_3_3.StrVal()), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Main_check("nested count access", "2", Data_Show_ShowIntImpl(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_0_0, "child"), "count").IntVal)), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Main_check("nested label access", "nested", gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_0_0, "child"), "label").StrVal()), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_check("array access", "[-3,0,7]", gopurs_runtime.Apply(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showArray1())).V0, gopurs_runtime.RecordGet(__local_var_0_0, "values")).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_check("payload FFI map", "2:nested:[-3 0 7]", gopurs_runtime.Apply(Get_Main_describePayloadMap(), __local_var_0_0).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Get_Main_compactPayload(), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Get_Main_describePayloadMap(), __local_var_9_9)
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_check("compact payload FFI map", "2:nested:[-3 0 7]", __local_var_10_10.StrVal()), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_check("empty array update", "2:nested:[]", gopurs_runtime.Apply(Get_Main_describePayloadMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					values []int64
				} {
					orig := gopurs_runtime.RecordUpdate1(__local_var_0_0, "values", func() gopurs_runtime.Value {
						arr := []int64{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
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
				return gopurs_runtime.RecordDict2("child", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_check("child update", "9:replacement:[-3 0 7]", gopurs_runtime.Apply(Get_Main_describePayloadMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					values []int64
				} {
					orig := gopurs_runtime.RecordUpdate1(__local_var_0_0, "child", func() gopurs_runtime.Value {
						orig := struct {
							count int64
							label string
						}{int64(9), "replacement"}
						_ = orig
						return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
					}())
					_ = orig
					clone := struct {
						child struct {
							count int64
							label string
						}
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
				return gopurs_runtime.RecordDict2("child", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Main_check("nested field update", "2:updated:[-3 0 7]", gopurs_runtime.Apply(Get_Main_describePayloadMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					child struct {
						count int64
						label string
					}
					values []int64
				} {
					orig := gopurs_runtime.RecordUpdate1(__local_var_0_0, "child", func() gopurs_runtime.Value {
						orig := func() struct {
							count int64
							label string
						} {
							orig := gopurs_runtime.RecordUpdate1(gopurs_runtime.RecordGet(__local_var_0_0, "child"), "label", gopurs_runtime.Str("updated"))
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
				return gopurs_runtime.RecordDict2("child", "values", func() gopurs_runtime.Value {
					orig := orig.child
					_ = orig
					return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
				}(), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Main_check("payload after updates", "2:nested:[-3 0 7]", gopurs_runtime.Apply(Get_Main_describePayloadMap(), __local_var_0_0).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_16_16
			return gopurs_runtime.Apply(Call_Main_check("retained payload", "2:nested:[-3 0 7]", gopurs_runtime.Apply(Get_Main_describePayloadMap(), __local_var_16_16).StrVal()), gopurs_runtime.Value{})
		})
	})
	return cache_Main_checkPayload
}

var cache_Main_checkScalars gopurs_runtime.Value
var once_Main_checkScalars sync.Once

func Get_Main_checkScalars() gopurs_runtime.Value {
	once_Main_checkScalars.Do(func() {
		cache_Main_checkScalars = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Get_Main_compactScalars(), gopurs_runtime.Value{})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(Get_Main_compactScalars(), gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Get_Main_compactKeys(), __local_var_1_1)
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Main_check("scalar compact keys", "number|flag", __local_var_2_2.StrVal()), gopurs_runtime.Value{})
			_ = __local_var_3_3
			var __t5 string
			{
				if (gopurs_runtime.RecordGet(__local_var_0_0, "flag").IntVal) != (0) {
					__t5 = "true"
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = "false"
			}
		end_branch_5:
			__local_var_4_4 := gopurs_runtime.Apply(Call_Main_check("boolean access", "true", __t5), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_6 := gopurs_runtime.Apply(Call_Main_check("number access", "-1.25", Data_Show_ShowNumberImpl(gopurs_runtime.RecordGet(__local_var_0_0, "number").FloatVal())), gopurs_runtime.Value{})
			_ = __local_var_5_6
			__local_var_6_7 := gopurs_runtime.Apply(Call_Main_check("scalar FFI map", "true:-1.25", gopurs_runtime.Apply(Get_Main_describeScalarsMap(), __local_var_0_0).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_6_7
			__local_var_7_8 := gopurs_runtime.Apply(Get_Main_compactScalars(), gopurs_runtime.Value{})
			_ = __local_var_7_8
			__local_var_8_9 := gopurs_runtime.Apply(Get_Main_describeScalarsMap(), __local_var_7_8)
			_ = __local_var_8_9
			__local_var_9_10 := gopurs_runtime.Apply(Call_Main_check("compact scalar FFI map", "true:-1.25", __local_var_8_9.StrVal()), gopurs_runtime.Value{})
			_ = __local_var_9_10
			__local_var_10_11 := gopurs_runtime.Apply(Call_Main_check("scalar updates", "false:2.5", gopurs_runtime.Apply(Get_Main_describeScalarsMap(), func() gopurs_runtime.Value {
				orig := func() struct {
					flag   bool
					number float64
				} {
					orig := gopurs_runtime.RecordUpdate2(__local_var_0_0, "flag", gopurs_runtime.Bool(false), "number", gopurs_runtime.Float(2.5))
					_ = orig
					clone := struct {
						flag   bool
						number float64
					}{}
					clone.flag = (gopurs_runtime.RecordGet(orig, "flag").IntVal) != (0)
					clone.number = gopurs_runtime.RecordGet(orig, "number").FloatVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("flag", "number", gopurs_runtime.Bool(orig.flag), gopurs_runtime.Float(orig.number))
			}()).StrVal()), gopurs_runtime.Value{})
			_ = __local_var_10_11
			return gopurs_runtime.Apply(Call_Main_check("retained scalars", "true:-1.25", gopurs_runtime.Apply(Get_Main_describeScalarsMap(), __local_var_0_0).StrVal()), gopurs_runtime.Value{})
		})
	})
	return cache_Main_checkScalars
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Main_checkEntry(false)
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Main_checkEntry(true), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Get_Main_checkPayload(), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Get_Main_checkScalars(), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Get_Main_checkEffects(), gopurs_runtime.Value{})
			_ = __local_var_5_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_makeEntry(count_0_loop int64, label_1_loop string) struct {
	count int64
	label string
} {
	var count_0 int64 = count_0_loop
	_ = count_0
	var label_1 string = label_1_loop
	_ = label_1
	return struct {
		count int64
		label string
	}{count_0, label_1}
}

func Call_Main_produceEntry(trace_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var trace_0 gopurs_runtime.Value = trace_0_loop
	_ = trace_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_1, func() gopurs_runtime.Value {
						arr := []string{"produce"}
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
		}), trace_0)
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = __local_var_2_1
		return func() gopurs_runtime.Value {
			orig := Call_Main_makeEntry(int64(9), "created")
			_ = orig
			return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
		}()
	})
}

func Call_Main_inspectProduced(inspect_0_loop gopurs_runtime.Value, produce_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var inspect_0 gopurs_runtime.Value = inspect_0_loop
	_ = inspect_0
	var produce_1 gopurs_runtime.Value = produce_1_loop
	_ = produce_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		value_2_0 := gopurs_runtime.Apply(produce_1, gopurs_runtime.Value{})
		_ = value_2_0
		return gopurs_runtime.Str(gopurs_runtime.Apply(inspect_0, value_2_0).StrVal())
	})
}

func Call_Main_inspectProduced__354691939(inspect_unused_0_loop gopurs_runtime.Value, produce_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
inspectProduced__354691939:
	for {
		if false {
			continue inspectProduced__354691939
		}
		var inspect_unused_0 gopurs_runtime.Value = inspect_unused_0_loop
		_ = inspect_unused_0
		var produce_unused_1 gopurs_runtime.Value = produce_unused_1_loop
		_ = produce_unused_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(Get_Main_compactPayload(), gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Main_compactKeys(), __local_var_2_0)
		})
	}
}

func Call_Main_inspectProduced__1003637247(inspect_unused_0_loop gopurs_runtime.Value, produce_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
inspectProduced__1003637247:
	for {
		if false {
			continue inspectProduced__1003637247
		}
		var inspect_unused_0 gopurs_runtime.Value = inspect_unused_0_loop
		_ = inspect_unused_0
		var produce_unused_1 gopurs_runtime.Value = produce_unused_1_loop
		_ = produce_unused_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(Get_Main_compactPayload(), gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Main_describePayloadMap(), __local_var_2_0)
		})
	}
}

func Call_Main_inspectProduced__2901793856(inspect_unused_0_loop gopurs_runtime.Value, produce_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
inspectProduced__2901793856:
	for {
		if false {
			continue inspectProduced__2901793856
		}
		var inspect_unused_0 gopurs_runtime.Value = inspect_unused_0_loop
		_ = inspect_unused_0
		var produce_1 gopurs_runtime.Value = produce_1_loop
		_ = produce_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(produce_1, gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Main_compactKeys(), __local_var_2_0)
		})
	}
}

func Call_Main_inspectProduced__1824334758(inspect_unused_0_loop gopurs_runtime.Value, produce_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
inspectProduced__1824334758:
	for {
		if false {
			continue inspectProduced__1824334758
		}
		var inspect_unused_0 gopurs_runtime.Value = inspect_unused_0_loop
		_ = inspect_unused_0
		var produce_1 gopurs_runtime.Value = produce_1_loop
		_ = produce_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(produce_1, gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Main_describeEntryMap(), __local_var_2_0)
		})
	}
}

func Call_Main_inspectProduced__3949581232(inspect_unused_0_loop gopurs_runtime.Value, produce_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
inspectProduced__3949581232:
	for {
		if false {
			continue inspectProduced__3949581232
		}
		var inspect_unused_0 gopurs_runtime.Value = inspect_unused_0_loop
		_ = inspect_unused_0
		var produce_unused_1 gopurs_runtime.Value = produce_unused_1_loop
		_ = produce_unused_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(Get_Main_compactScalars(), gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Main_compactKeys(), __local_var_2_0)
		})
	}
}

func Call_Main_inspectProduced__1784222911(inspect_unused_0_loop gopurs_runtime.Value, produce_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
inspectProduced__1784222911:
	for {
		if false {
			continue inspectProduced__1784222911
		}
		var inspect_unused_0 gopurs_runtime.Value = inspect_unused_0_loop
		_ = inspect_unused_0
		var produce_unused_1 gopurs_runtime.Value = produce_unused_1_loop
		_ = produce_unused_1
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.Apply(Get_Main_compactScalars(), gopurs_runtime.Value{})
			_ = __local_var_2_0
			return gopurs_runtime.Apply(Get_Main_describeScalarsMap(), __local_var_2_0)
		})
	}
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
		__local_var_3_0 := Call_Test_Assert_assertEqual_prime___4176622598("", struct {
			actual   string
			expected string
		}{actual_2, expected_1})
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(((label_0)+(": "))+(actual_2))), gopurs_runtime.Value{})
	})
}

func Call_Main_checkEntry(reverse_0_loop bool) gopurs_runtime.Value {
	var reverse_0 bool = reverse_0_loop
	_ = reverse_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Main_compactEntry(), gopurs_runtime.Bool(reverse_0))
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = __local_var_2_1
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_3_2
		var __t4 string
		{
			if reverse_0 {
				__t4 = "reversed "
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = "ordered "
		}
	end_branch_4:
		// TAST (Let): prefix_4_3 shape=Branch(LitString, def=LitString) bindingType=Any
		prefix_4_3 := __t4
		_ = prefix_4_3
		__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Main_compactEntry(), gopurs_runtime.Bool(reverse_0)), gopurs_runtime.Value{})
		_ = __local_var_5_6
		__local_var_5_5 := gopurs_runtime.Apply(Get_Main_compactKeys(), __local_var_5_6)
		_ = __local_var_5_5
		var __t8 string
		{
			if reverse_0 {
				__t8 = "label|count"
				goto end_branch_8
			} else {

			}
		}
		{
			__t8 = "count|label"
		}
	end_branch_8:
		__local_var_6_7 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("compact keys"), __t8, __local_var_5_5.StrVal()), gopurs_runtime.Value{})
		_ = __local_var_6_7
		__local_var_7_9 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("count access"), "5", Data_Show_ShowIntImpl(gopurs_runtime.RecordGet(__local_var_2_1, "count").IntVal)), gopurs_runtime.Value{})
		_ = __local_var_7_9
		__local_var_8_10 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("label access"), "alpha", gopurs_runtime.RecordGet(__local_var_2_1, "label").StrVal()), gopurs_runtime.Value{})
		_ = __local_var_8_10
		__local_var_9_11 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("FFI map"), "5:alpha", gopurs_runtime.Apply(Get_Main_describeEntryMap(), __local_var_2_1).StrVal()), gopurs_runtime.Value{})
		_ = __local_var_9_11
		__local_var_10_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Main_compactEntry(), gopurs_runtime.Bool(reverse_0)), gopurs_runtime.Value{})
		_ = __local_var_10_13
		__local_var_10_12 := gopurs_runtime.Apply(Get_Main_describeEntryMap(), __local_var_10_13)
		_ = __local_var_10_12
		__local_var_11_14 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("compact FFI map"), "5:alpha", __local_var_10_12.StrVal()), gopurs_runtime.Value{})
		_ = __local_var_11_14
		// TAST (Let): changedCount_12_15 shape=Other bindingType=(Record (Row [count: Int, label: String] Empty))
		changedCount_12_15 := func() struct {
			count int64
			label string
		} {
			orig := gopurs_runtime.RecordUpdate1(__local_var_2_1, "count", gopurs_runtime.Int(int64(-7)))
			_ = orig
			clone := struct {
				count int64
				label string
			}{}
			clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
			clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
			return clone
		}()
		_ = changedCount_12_15
		__local_var_13_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			orig := changedCount_12_15
			_ = orig
			return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
		}()), gopurs_runtime.Value{})
		_ = __local_var_13_16
		__local_var_14_17 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("count update"), "-7:alpha", gopurs_runtime.Apply(Get_Main_describeEntryMap(), func() gopurs_runtime.Value {
			orig := changedCount_12_15
			_ = orig
			return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label))
		}()).StrVal()), gopurs_runtime.Value{})
		_ = __local_var_14_17
		__local_var_15_18 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("label update"), "5:beta", gopurs_runtime.Apply(Get_Main_describeEntryMap(), func() gopurs_runtime.Value {
			orig := func() struct {
				count int64
				label string
			} {
				orig := gopurs_runtime.RecordUpdate1(__local_var_2_1, "label", gopurs_runtime.Str("beta"))
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
		_ = __local_var_15_18
		__local_var_16_19 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("both updates"), "0:", gopurs_runtime.Apply(Get_Main_describeEntryMap(), func() gopurs_runtime.Value {
			orig := func() struct {
				count int64
				label string
			} {
				orig := gopurs_runtime.RecordUpdate2(__local_var_2_1, "label", gopurs_runtime.Str(""), "count", gopurs_runtime.Int(int64(0)))
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
		_ = __local_var_16_19
		__local_var_17_20 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("original after updates"), "5:alpha", gopurs_runtime.Apply(Get_Main_describeEntryMap(), __local_var_2_1).StrVal()), gopurs_runtime.Value{})
		_ = __local_var_17_20
		__local_var_18_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_2), gopurs_runtime.Value{})
		_ = __local_var_18_21
		__local_var_19_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_13_16), gopurs_runtime.Value{})
		_ = __local_var_19_22
		__local_var_20_23 := gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("retained original"), "5:alpha", gopurs_runtime.Apply(Get_Main_describeEntryMap(), __local_var_18_21).StrVal()), gopurs_runtime.Value{})
		_ = __local_var_20_23
		return gopurs_runtime.Apply(Call_Main_check((prefix_4_3)+("retained count update"), "-7:alpha", gopurs_runtime.Apply(Get_Main_describeEntryMap(), __local_var_19_22).StrVal()), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_131790935_3790796878(in *Constructor_Data_Eq_Eq[[]string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1953100407(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]string])(unsafe.Pointer(in))
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1953100407_1386611502(in *Constructor_Data_Show_Show[[]string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_131790935(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]string])(unsafe.Pointer(in))
}

func Get_Main_compactEntry() gopurs_runtime.Value {
	return _Gopurs_Main_CompactEntry
}

func Get_Main_compactKeys() gopurs_runtime.Value {
	return _Gopurs_Main_CompactKeys
}

func Get_Main_compactPayload() gopurs_runtime.Value {
	return _Gopurs_Main_CompactPayload
}

func Get_Main_compactScalars() gopurs_runtime.Value {
	return _Gopurs_Main_CompactScalars
}

func Get_Main_describeEntryMap() gopurs_runtime.Value {
	return _Gopurs_Main_DescribeEntryMap
}

func Get_Main_describePayloadMap() gopurs_runtime.Value {
	return _Gopurs_Main_DescribePayloadMap
}

func Get_Main_describeScalarsMap() gopurs_runtime.Value {
	return _Gopurs_Main_DescribeScalarsMap
}
