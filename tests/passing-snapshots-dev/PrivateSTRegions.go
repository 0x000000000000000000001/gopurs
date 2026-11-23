package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_valuesIsSymbol gopurs_runtime.Value
var once_Main_valuesIsSymbol sync.Once

func Get_Main_valuesIsSymbol() gopurs_runtime.Value {
	once_Main_valuesIsSymbol.Do(func() {
		cache_Main_valuesIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("values", 6)
		}))
	})
	return cache_Main_valuesIsSymbol
}

var cache_Main_nextIsSymbol gopurs_runtime.Value
var once_Main_nextIsSymbol sync.Once

func Get_Main_nextIsSymbol() gopurs_runtime.Value {
	once_Main_nextIsSymbol.Do(func() {
		cache_Main_nextIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("next", 4)
		}))
	})
	return cache_Main_nextIsSymbol
}

var cache_Main_failedIsSymbol gopurs_runtime.Value
var once_Main_failedIsSymbol sync.Once

func Get_Main_failedIsSymbol() gopurs_runtime.Value {
	once_Main_failedIsSymbol.Do(func() {
		cache_Main_failedIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("failed", 6)
		}))
	})
	return cache_Main_failedIsSymbol
}

var cache_Main_beforeIsSymbol gopurs_runtime.Value
var once_Main_beforeIsSymbol sync.Once

func Get_Main_beforeIsSymbol() gopurs_runtime.Value {
	once_Main_beforeIsSymbol.Do(func() {
		cache_Main_beforeIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("before", 6)
		}))
	})
	return cache_Main_beforeIsSymbol
}

var cache_Main_afterIsSymbol gopurs_runtime.Value
var once_Main_afterIsSymbol sync.Once

func Get_Main_afterIsSymbol() gopurs_runtime.Value {
	once_Main_afterIsSymbol.Do(func() {
		cache_Main_afterIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("after", 5)
		}))
	})
	return cache_Main_afterIsSymbol
}

var cache_Main_Emit gopurs_runtime.Value
var once_Main_Emit sync.Once

func Get_Main_Emit() gopurs_runtime.Value {
	once_Main_Emit.Do(func() {
		cache_Main_Emit = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, value0, value1.IntVal}))}
			})
		})
	})
	return cache_Main_Emit
}

var cache_Main_Stop gopurs_runtime.Value
var once_Main_Stop sync.Once

func Get_Main_Stop() gopurs_runtime.Value {
	once_Main_Stop.Do(func() {
		cache_Main_Stop = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, value0.IntVal}))}
		})
	})
	return cache_Main_Stop
}

var cache_Main_Failed gopurs_runtime.Value
var once_Main_Failed sync.Once

func Get_Main_Failed() gopurs_runtime.Value {
	once_Main_Failed.Do(func() {
		cache_Main_Failed = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 270449853, UnsafePtr: unsafe.Pointer((&Constructor_Main_Failed[gopurs_runtime.Value]{1, value0.IntVal}))}
		})
	})
	return cache_Main_Failed
}

var cache_Main_snapshotRead gopurs_runtime.Value
var once_Main_snapshotRead sync.Once

func Get_Main_snapshotRead() gopurs_runtime.Value {
	once_Main_snapshotRead.Do(func() {
		cache_Main_snapshotRead = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_snapshotRead(initial_0_box.IntVal)
		})
	})
	return cache_Main_snapshotRead
}

var cache_Main_skipLoop gopurs_runtime.Value
var once_Main_skipLoop sync.Once

func Get_Main_skipLoop() gopurs_runtime.Value {
	once_Main_skipLoop.Do(func() {
		cache_Main_skipLoop = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_skipLoop(initial_0_box.IntVal))
		})
	})
	return cache_Main_skipLoop
}

var cache_Main_operationResults gopurs_runtime.Value
var once_Main_operationResults sync.Once

func Get_Main_operationResults() gopurs_runtime.Value {
	once_Main_operationResults.Do(func() {
		cache_Main_operationResults = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_operationResults(initial_0_box.IntVal)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_operationResults
}

var cache_Main_freezeSnapshot gopurs_runtime.Value
var once_Main_freezeSnapshot sync.Once

func Get_Main_freezeSnapshot() gopurs_runtime.Value {
	once_Main_freezeSnapshot.Do(func() {
		cache_Main_freezeSnapshot = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_freezeSnapshot(initial_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
					arr := orig.after
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.before
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()
		})
	})
	return cache_Main_freezeSnapshot
}

var cache_Main_collect gopurs_runtime.Value
var once_Main_collect sync.Once

func Get_Main_collect() gopurs_runtime.Value {
	once_Main_collect.Do(func() {
		cache_Main_collect = gopurs_runtime.Func2(func(step_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_collect(step_0_box, start_1_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), gopurs_runtime.Array(orig.values))
			}()
		})
	})
	return cache_Main_collect
}

var cache_Main_aliasRead gopurs_runtime.Value
var once_Main_aliasRead sync.Once

func Get_Main_aliasRead() gopurs_runtime.Value {
	once_Main_aliasRead.Do(func() {
		cache_Main_aliasRead = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_aliasRead(initial_0_box.IntVal))
		})
	})
	return cache_Main_aliasRead
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if (n_0.IntVal) < (int64(4)) {
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.IntMul(n_0.IntVal, int64(3))), gopurs_runtime.IntAdd(n_0.IntVal, int64(1))}))}
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, gopurs_runtime.IntAdd(n_0.IntVal, int64(7))}))}
				}
			end_branch_1:
				return __t1
			}))
			_ = __local_var_0_0
			stepRef_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = stepRef_1_2
			step_2_3 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), stepRef_1_2, gopurs_runtime.Value{})
			_ = step_2_3
			__callback_fn_5 := Get_Effect_Ref_go__new()
			_ = __callback_fn_5
			__callback_arg_6 := func(record struct {
				failed bool
				next   int64
				values []gopurs_runtime.Value
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}) struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			} {
				return struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{record.failed, record.next, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.values).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()}
			}(Call_Main_collect(step_2_3, int64(1)))
			_ = __callback_arg_6
			keptRef_3_4 := gopurs_runtime.Apply(func(__native_callback func(struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
				if __native_callback_ok {
					return __native_callback(__callback_arg_6)
				} else {

				}
				return gopurs_runtime.Apply(__callback_fn_5, func() gopurs_runtime.Value {
					orig := __callback_arg_6
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(gopurs_runtime.FunctionData[func(struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}) gopurs_runtime.Value](__callback_fn_5)), gopurs_runtime.Value{})
			_ = keptRef_3_4
			first_4_7 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), keptRef_3_4, gopurs_runtime.Value{})
			_ = first_4_7
			_dollar___unused_5_8 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					expected struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				} {
					orig := first_4_7
					_ = orig
					clone := struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}{}
					clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
					clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}(), struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{false, int64(11), []int64{int64(3), int64(6), int64(9)}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_5_8
			_dollar___unused_6_9 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					expected struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func(record struct {
					failed bool
					next   int64
					values []gopurs_runtime.Value
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}) struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				} {
					return struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}{record.failed, record.next, func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.values).UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()}
				}(Call_Main_collect(step_2_3, int64(3))), struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{false, int64(11), []int64{int64(9)}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_6_9
			_dollar___unused_7_10 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_4()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_5())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					expected struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func(record struct {
					failed bool
					next   int64
					values []gopurs_runtime.Value
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}) struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				} {
					return struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}{record.failed, record.next, func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.values).UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()}
				}(Call_Main_collect(step_2_3, int64(5))), struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{false, int64(12), []int64{}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_7_10
			preserved_8_11 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), keptRef_3_4, gopurs_runtime.Value{})
			_ = preserved_8_11
			_dollar___unused_9_12 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_6()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_7())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					expected struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				} {
					orig := preserved_8_11
					_ = orig
					clone := struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}{}
					clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
					clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}(), struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{false, int64(11), []int64{int64(3), int64(6), int64(9)}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_9_12
			_dollar___unused_10_13 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_8()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_9())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					expected struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func(record struct {
					failed bool
					next   int64
					values []gopurs_runtime.Value
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}) struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				} {
					return struct {
						failed bool
						next   int64
						values []int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}{record.failed, record.next, func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.values).UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()}
				}(Call_Main_collect(gopurs_runtime.Func(func(n_10 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t14 gopurs_runtime.Value
					{
						if (n_10.IntVal) < (int64(2)) {
							__t14 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, n_10, gopurs_runtime.IntAdd(n_10.IntVal, int64(1))}))}
							goto end_branch_14
						} else {

						}
					}
					{
						__t14 = gopurs_runtime.Value{Type: 9, IntVal: 270449853, UnsafePtr: unsafe.Pointer((&Constructor_Main_Failed[gopurs_runtime.Value]{1, int64(91)}))}
					}
				end_branch_14:
					return __t14
				}), int64(0))), struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{true, int64(91), []int64{}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_10_13
			_dollar___unused_11_15 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_10()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_11())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						failed bool
						next   int64
						values []string
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					expected struct {
						failed bool
						next   int64
						values []string
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func(record struct {
					failed bool
					next   int64
					values []gopurs_runtime.Value
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}) struct {
					failed bool
					next   int64
					values []string
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				} {
					return struct {
						failed bool
						next   int64
						values []string
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}{record.failed, record.next, func() []string {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.values).UnsafePtr)
						unboxed := make([]string, len(arr))
						for i, v := range arr {
							unboxed[i] = v.StrVal()
						}
						return unboxed
					}()}
				}(Call_Main_collect(gopurs_runtime.Func(func(n_11 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t16 gopurs_runtime.Value
					{
						if (n_11.IntVal) < (int64(2)) {
							__t16 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, gopurs_runtime.StrWithCodeUnits("λ😀", 3), gopurs_runtime.IntAdd(n_11.IntVal, int64(1))}))}
							goto end_branch_16
						} else {

						}
					}
					{
						__t16 = gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, n_11.IntVal}))}
					}
				end_branch_16:
					return __t16
				}), int64(0))), struct {
					failed bool
					next   int64
					values []string
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{false, int64(2), []string{"λ😀", "λ😀"}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_11_15
			_dollar___unused_12_17 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_12()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_13())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						failed bool
						next   int64
						values [][]int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					expected struct {
						failed bool
						next   int64
						values [][]int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func(record struct {
					failed bool
					next   int64
					values []gopurs_runtime.Value
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}) struct {
					failed bool
					next   int64
					values [][]int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				} {
					return struct {
						failed bool
						next   int64
						values [][]int64
						/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
					}{record.failed, record.next, func() [][]int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(record.values).UnsafePtr)
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
					}()}
				}(Call_Main_collect(gopurs_runtime.Func(func(n_12 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t18 gopurs_runtime.Value
					{
						if (n_12.IntVal) < (int64(2)) {
							__t18 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
								arr := []int64{n_12.IntVal, gopurs_runtime.IntAdd(n_12.IntVal, int64(1))}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), gopurs_runtime.IntAdd(n_12.IntVal, int64(1))}))}
							goto end_branch_18
						} else {

						}
					}
					{
						__t18 = gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, n_12.IntVal}))}
					}
				end_branch_18:
					return __t18
				}), int64(0))), struct {
					failed bool
					next   int64
					values [][]int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{false, int64(2), [][]int64{[]int64{int64(0), int64(1)}, []int64{int64(1), int64(2)}}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
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
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
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
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_12_17
			__local_var_13_20 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(10)), gopurs_runtime.Value{})
			_ = __local_var_13_20
			input_13_19 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_13_20, gopurs_runtime.Value{})
			_ = input_13_19
			_dollar___unused_14_21 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_14()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_15())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_operationResults(input_13_19.IntVal), []int64{int64(11), int64(12)}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_14_21
			_dollar___unused_15_22 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_skipLoop(input_13_19.IntVal), int64(10)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_15_22
			_dollar___unused_16_23 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_16()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_17())), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						after  []int64
						before []int64
						/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
					}
					expected struct {
						after  []int64
						before []int64
						/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_freezeSnapshot(input_13_19.IntVal), struct {
					after  []int64
					before []int64
					/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
				}{[]int64{int64(10), int64(11)}, []int64{int64(10)}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						arr := orig.after
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := orig.before
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						arr := orig.after
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := orig.before
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_16_23
			__st_ref25 := input_13_19
			_ = __st_ref25
			__st_ref25 = gopurs_runtime.Int(gopurs_runtime.IntAdd(input_13_19.IntVal, int64(1)))
			__st_value26 := __st_ref25
			_ = __st_value26
			_dollar___unused_17_24 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{__st_ref25.IntVal, int64(11)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_17_24
			_dollar___unused_18_27 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply(Call_Main_snapshotRead(input_13_19.IntVal), gopurs_runtime.Int(int64(5))).IntVal, int64(15)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_18_27
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_nextIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_failedIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_failedIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_nextIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_nextIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_failedIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_failedIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_nextIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_nextIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_failedIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_failedIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_nextIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_nextIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_failedIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_failedIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_nextIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_nextIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_failedIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_failedIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_nextIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), gopurs_runtime.Value{}, Get_Main_nextIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_failedIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_failedIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_nextIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_11
}

var cache_Main___cached_dict_12 gopurs_runtime.Value
var once_Main___cached_dict_12 sync.Once

func Get_Main___cached_dict_12() gopurs_runtime.Value {
	once_Main___cached_dict_12.Do(func() {
		cache_Main___cached_dict_12 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), gopurs_runtime.Value{}, Get_Main_nextIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_failedIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_12
}

var cache_Main___cached_dict_13 gopurs_runtime.Value
var once_Main___cached_dict_13 sync.Once

func Get_Main___cached_dict_13() gopurs_runtime.Value {
	once_Main___cached_dict_13.Do(func() {
		cache_Main___cached_dict_13 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_failedIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_nextIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_13
}

var cache_Main___cached_dict_14 gopurs_runtime.Value
var once_Main___cached_dict_14 sync.Once

func Get_Main___cached_dict_14() gopurs_runtime.Value {
	once_Main___cached_dict_14.Do(func() {
		cache_Main___cached_dict_14 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_14
}

var cache_Main___cached_dict_15 gopurs_runtime.Value
var once_Main___cached_dict_15 sync.Once

func Get_Main___cached_dict_15() gopurs_runtime.Value {
	once_Main___cached_dict_15.Do(func() {
		cache_Main___cached_dict_15 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_15
}

var cache_Main___cached_dict_16 gopurs_runtime.Value
var once_Main___cached_dict_16 sync.Once

func Get_Main___cached_dict_16() gopurs_runtime.Value {
	once_Main___cached_dict_16.Do(func() {
		cache_Main___cached_dict_16 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_beforeIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_afterIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))
	})
	return cache_Main___cached_dict_16
}

var cache_Main___cached_dict_17 gopurs_runtime.Value
var once_Main___cached_dict_17 sync.Once

func Get_Main___cached_dict_17() gopurs_runtime.Value {
	once_Main___cached_dict_17.Do(func() {
		cache_Main___cached_dict_17 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_afterIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_beforeIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))
	})
	return cache_Main___cached_dict_17
}

type Constructor_Main_Emit[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 int64
}

type Constructor_Main_Stop[T_a any] struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_Failed[T_a any] struct {
	Rc uint32
	V0 int64
}

func Call_Main_snapshotRead(initial_0_loop int64) gopurs_runtime.Value {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	__st_ref0 := int64(initial_0)
	_ = __st_ref0
	__st_value1 := int64(__st_ref0)
	_ = __st_value1
	__st_ref0 = gopurs_runtime.IntAdd(initial_0, int64(1))
	__st_value2 := int64(__st_ref0)
	_ = __st_value2
	return gopurs_runtime.Func(func(offset_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(gopurs_runtime.IntAdd(__st_value1, offset_4.IntVal))
	})
}

func Call_Main_skipLoop(initial_0_loop int64) int64 {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	__st_ref0 := int64(initial_0)
	_ = __st_ref0
	for {
		if !(false) {
			break
		} else {

		}
		__st_value1 := int64(__st_ref0)
		_ = __st_value1
		__st_ref0 = gopurs_runtime.IntAdd(__st_value1, int64(1))
		__st_value2 := int64(__st_ref0)
		_ = __st_value2
		_ = Get_Data_Unit_unit()
	}
	__st_value3 := gopurs_runtime.Value{}
	_ = __st_value3
	return __st_ref0
}

func Call_Main_operationResults(initial_0_loop int64) []int64 {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	__st_ref0 := int64(initial_0)
	_ = __st_ref0
	__st_array1 := []gopurs_runtime.Value{}
	_ = __st_array1
	__st_ref0 = gopurs_runtime.IntAdd(initial_0, int64(1))
	__st_value2 := int64(__st_ref0)
	_ = __st_value2
	__st_array1 = append(__st_array1, gopurs_runtime.Int(__st_value2))
	__st_value3 := int64(int64(len(__st_array1)))
	_ = __st_value3
	__st_value4 := int64(__st_ref0)
	_ = __st_value4
	__st_array1 = append(__st_array1, gopurs_runtime.Int(gopurs_runtime.IntAdd(__st_value4, __st_value3)))
	__st_value5 := int64(int64(len(__st_array1)))
	_ = __st_value5
	__st_frozen6 := make([]gopurs_runtime.Value, len(__st_array1))
	_ = __st_frozen6
	copy(__st_frozen6, __st_array1)
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(__st_frozen6).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Main_freezeSnapshot(initial_0_loop int64) struct {
	after  []int64
	before []int64
	/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
} {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	return func() struct {
		after  []int64
		before []int64
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	} {
		orig := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			values_1_0 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
			_ = values_1_0
			_dollar___unused_2_1 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Int(initial_0), values_1_0)
			_ = _dollar___unused_2_1
			before_3_2 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_freezeImpl(), values_1_0)
			_ = before_3_2
			_dollar___unused_4_3 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Int(gopurs_runtime.IntAdd(initial_0, int64(1))), values_1_0)
			_ = _dollar___unused_4_3
			after_5_4 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), values_1_0)
			_ = after_5_4
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, after_5_4, before_3_2)
		}))
		_ = orig
		clone := struct {
			after  []int64
			before []int64
			/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
		}{}
		clone.after = func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "after").UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()
		clone.before = func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "before").UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()
		return clone
	}()
}

func Call_Main_collect(step_0_loop gopurs_runtime.Value, start_1_loop int64) struct {
	failed bool
	next   int64
	values []gopurs_runtime.Value
	/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
} {
	var step_0 gopurs_runtime.Value = step_0_loop
	_ = step_0
	var start_1 int64 = start_1_loop
	_ = start_1
	__st_array0 := []gopurs_runtime.Value{}
	_ = __st_array0
	__st_ref1 := int64(start_1)
	_ = __st_ref1
	__st_ref2 := true
	_ = __st_ref2
	__st_ref3 := gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(false), gopurs_runtime.Int(start_1), gopurs_runtime.Array([]gopurs_runtime.Value{}))
	_ = __st_ref3
	for {
		if !(__st_ref2) {
			break
		} else {

		}
		__st_value4 := int64(__st_ref1)
		_ = __st_value4
		__st_value5 := gopurs_runtime.Apply(step_0, gopurs_runtime.Int(__st_value4))
		_ = __st_value5
		var __st_branch6 gopurs_runtime.Value
		if __st_value5.Type == 9 && __st_value5.IntVal == 621001451 {
			__st_array0 = append(__st_array0, (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0)
			__st_value13 := int64(int64(len(__st_array0)))
			_ = __st_value13
			__st_ref1 = (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V1
			__st_value14 := int64(__st_ref1)
			_ = __st_value14
			__st_branch6 = Get_Data_Unit_unit()
		} else {
			if __st_value5.Type == 9 && __st_value5.IntVal == 325197958 {
				__st_frozen9 := make([]gopurs_runtime.Value, len(__st_array0))
				_ = __st_frozen9
				copy(__st_frozen9, __st_array0)
				__st_value10 := __st_frozen9
				_ = __st_value10
				__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(false), gopurs_runtime.Int((*Constructor_Main_Stop[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), gopurs_runtime.Array(__st_value10))
				__st_value11 := __st_ref3
				_ = __st_value11
				__st_ref2 = false
				__st_value12 := __st_ref2
				_ = __st_value12
				__st_branch6 = Get_Data_Unit_unit()
			} else {
				if __st_value5.Type == 9 && __st_value5.IntVal == 270449853 {
					__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(true), gopurs_runtime.Int((*Constructor_Main_Failed[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), gopurs_runtime.Array([]gopurs_runtime.Value{}))
					__st_value7 := __st_ref3
					_ = __st_value7
					__st_ref2 = false
					__st_value8 := __st_ref2
					_ = __st_value8
					__st_branch6 = Get_Data_Unit_unit()
				} else {
					__st_branch6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			}
		}
		_ = __st_branch6
	}
	__st_value15 := gopurs_runtime.Value{}
	_ = __st_value15
	return func() struct {
		failed bool
		next   int64
		values []gopurs_runtime.Value
		/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
	} {
		orig := __st_ref3
		_ = orig
		clone := struct {
			failed bool
			next   int64
			values []gopurs_runtime.Value
			/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
		}{}
		clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
		clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
		clone.values = (*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)))
		return clone
	}()
}

func Call_Main_aliasRead(initial_0_loop int64) int64 {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	__st_ref0 := int64(initial_0)
	_ = __st_ref0
	__st_ref0 = gopurs_runtime.IntAdd(initial_0, int64(1))
	__st_value1 := int64(__st_ref0)
	_ = __st_value1
	return __st_ref0
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

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
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

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
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

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys1{K0: "reflectSymbol"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys2{K0: "after", K1: "before"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys3{K0: "failed", K1: "next", K2: "values"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}
