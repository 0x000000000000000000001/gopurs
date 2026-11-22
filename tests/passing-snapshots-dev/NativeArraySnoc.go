package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_columnIsSymbol gopurs_runtime.Value
var once_Main_columnIsSymbol sync.Once

func Get_Main_columnIsSymbol() gopurs_runtime.Value {
	once_Main_columnIsSymbol.Do(func() {
		cache_Main_columnIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("column", 6)
		}))
	})
	return cache_Main_columnIsSymbol
}

var cache_Main_lineIsSymbol gopurs_runtime.Value
var once_Main_lineIsSymbol sync.Once

func Get_Main_lineIsSymbol() gopurs_runtime.Value {
	once_Main_lineIsSymbol.Do(func() {
		cache_Main_lineIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("line", 4)
		}))
	})
	return cache_Main_lineIsSymbol
}

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

var cache_Main_positionIsSymbol gopurs_runtime.Value
var once_Main_positionIsSymbol sync.Once

func Get_Main_positionIsSymbol() gopurs_runtime.Value {
	once_Main_positionIsSymbol.Do(func() {
		cache_Main_positionIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("position", 8)
		}))
	})
	return cache_Main_positionIsSymbol
}

var cache_Main_labelIsSymbol gopurs_runtime.Value
var once_Main_labelIsSymbol sync.Once

func Get_Main_labelIsSymbol() gopurs_runtime.Value {
	once_Main_labelIsSymbol.Do(func() {
		cache_Main_labelIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("label", 5)
		}))
	})
	return cache_Main_labelIsSymbol
}

var cache_Main_Empty gopurs_runtime.Value
var once_Main_Empty sync.Once

func Get_Main_Empty() gopurs_runtime.Value {
	once_Main_Empty.Do(func() {
		cache_Main_Empty = gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((*Constructor_Main_Item)(nil))}
	})
	return cache_Main_Empty
}

var cache_Main_Item gopurs_runtime.Value
var once_Main_Item sync.Once

func Get_Main_Item() gopurs_runtime.Value {
	once_Main_Item.Do(func() {
		cache_Main_Item = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((&Constructor_Main_Item{1, func() struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				} {
					orig := value0
					_ = orig
					clone := struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}{}
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.position = func() struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "position")
						_ = orig
						clone := struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}{}
						clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
						clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
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
				}(), Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](value1))}))}
			})
		})
	})
	return cache_Main_Item
}

var cache_Main_describe gopurs_runtime.Value
var once_Main_describe sync.Once

func Get_Main_describe() gopurs_runtime.Value {
	once_Main_describe.Do(func() {
		cache_Main_describe = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_describe(gopurs_runtime.CoerceToStruct[Constructor_Main_Item](v_0_box)))
		})
	})
	return cache_Main_describe
}

var cache_Main_applyAppend gopurs_runtime.Value
var once_Main_applyAppend sync.Once

func Get_Main_applyAppend() gopurs_runtime.Value {
	once_Main_applyAppend.Do(func() {
		cache_Main_applyAppend = gopurs_runtime.Func3(func(append1_0_box gopurs_runtime.Value, values_1_box gopurs_runtime.Value, value_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_applyAppend(append1_0_box, (*((*[]gopurs_runtime.Value)(values_1_box.UnsafePtr))), value_2_box))
		})
	})
	return cache_Main_applyAppend
}

var cache_Main_appendItem gopurs_runtime.Value
var once_Main_appendItem sync.Once

func Get_Main_appendItem() gopurs_runtime.Value {
	once_Main_appendItem.Do(func() {
		cache_Main_appendItem = Get_Data_Array_snoc()
	})
	return cache_Main_appendItem
}

var cache_Main_appendInt gopurs_runtime.Value
var once_Main_appendInt sync.Once

func Get_Main_appendInt() gopurs_runtime.Value {
	once_Main_appendInt.Do(func() {
		cache_Main_appendInt = Get_Data_Array_snoc()
	})
	return cache_Main_appendInt
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				orig := struct {
					arrays   [][]int64
					booleans []bool
					entry    struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}
					functions []gopurs_runtime.Value
					ints      []int64
					lists     []*Constructor_Data_List_Types_Cons[int64]
					numbers   []float64
					strings   []string
					/*gopurs:record-v1:8:006100720072006100790073,0062006f006f006c00650061006e0073,0065006e007400720079,00660075006e006300740069006f006e0073,0069006e00740073,006c0069007300740073,006e0075006d0062006500720073,0073007400720069006e00670073*/
				}{[][]int64{[]int64{}, []int64{int64(1), int64(2)}}, []bool{true, false}, struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}{"kept", struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}{int64(7), int64(11)}, []int64{int64(13), int64(17)}}, (*((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(n_0.IntVal, int64(1)))
				}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntMul(n_0.IntVal, int64(2)))
				})}).UnsafePtr))), []int64{int64(-2147483648), int64(0), int64(2147483647)}, func() []*Constructor_Data_List_Types_Cons[int64] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, int64(3), (&Constructor_Data_List_Types_Cons[int64]{1, int64(5), (*Constructor_Data_List_Types_Cons[int64])(nil)})})))}}).UnsafePtr)
					unboxed := make([]*Constructor_Data_List_Types_Cons[int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v))
					}
					return unboxed
				}(), []float64{1.25, -2.5}, []string{"", "λ😀"}}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"arrays", "booleans", "entry", "functions", "ints", "lists", "numbers", "strings"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					arr := orig.arrays
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
				}(), func() gopurs_runtime.Value {
					arr := orig.booleans
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Bool(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					orig := orig.entry
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), gopurs_runtime.Array(orig.functions), func() gopurs_runtime.Value {
					arr := orig.ints
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.lists
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.numbers
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Float(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.strings
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}())
			_ = __local_var_0_0
			inputsRef_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = inputsRef_1_1
			input_2_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), inputsRef_1_1, gopurs_runtime.Value{})
			_ = input_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_3()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_4())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), gopurs_runtime.RecordGet(input_2_2, "ints"), gopurs_runtime.Int(int64(-7))).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(-7)}}
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
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_5()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_6())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), func() gopurs_runtime.Value {
						arr := []int64{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), gopurs_runtime.Int(int64(42))).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(42)}}
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
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_7()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_8())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []float64
					expected []float64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []float64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(input_2_2, "numbers").UnsafePtr))), gopurs_runtime.Float(3.75))).UnsafePtr)
					unboxed := make([]float64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.FloatVal()
					}
					return unboxed
				}(), []float64{1.25, -2.5, 3.75}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Float(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Float(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_9()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_10())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []string
					expected []string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(input_2_2, "strings").UnsafePtr))), gopurs_runtime.StrWithCodeUnits("last", 4))).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}(), []string{"", "λ😀", "last"}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_11()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_12())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []bool
					expected []bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []bool {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(input_2_2, "booleans").UnsafePtr))), gopurs_runtime.Bool(true))).UnsafePtr)
					unboxed := make([]bool, len(arr))
					for i, v := range arr {
						unboxed[i] = (v.IntVal) != (0)
					}
					return unboxed
				}(), []bool{true, false, true}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Bool(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Bool(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_13()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_14())), func() gopurs_runtime.Value {
				orig := struct {
					actual   [][]int64
					expected [][]int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() [][]int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(input_2_2, "arrays").UnsafePtr))), func() gopurs_runtime.Value {
						arr := []int64{int64(3)}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())).UnsafePtr)
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
				}(), [][]int64{[]int64{}, []int64{int64(1), int64(2)}, []int64{int64(3)}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
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
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
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
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_15()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_16())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []*Constructor_Data_List_Types_Cons[int64]
					expected []*Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []*Constructor_Data_List_Types_Cons[int64] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(input_2_2, "lists").UnsafePtr))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, int64(7), (*Constructor_Data_List_Types_Cons[int64])(nil)})))})).UnsafePtr)
					unboxed := make([]*Constructor_Data_List_Types_Cons[int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v))
					}
					return unboxed
				}(), func() []*Constructor_Data_List_Types_Cons[int64] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, int64(3), (&Constructor_Data_List_Types_Cons[int64]{1, int64(5), (*Constructor_Data_List_Types_Cons[int64])(nil)})})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, int64(7), (*Constructor_Data_List_Types_Cons[int64])(nil)})))}}).UnsafePtr)
					unboxed := make([]*Constructor_Data_List_Types_Cons[int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v))
					}
					return unboxed
				}()}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(v))}
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			arrayIntrinsic_arg_11 := func(f_10 gopurs_runtime.Value) int64 {
				return gopurs_runtime.Apply(f_10, gopurs_runtime.Int(int64(10))).IntVal
			}
			_ = arrayIntrinsic_arg_11
			_ = arrayIntrinsic_arg_11
			arrayIntrinsic_arg_12 := Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(input_2_2, "functions").UnsafePtr))), gopurs_runtime.Func(func(n_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntSub(n_10.IntVal, int64(3)))
			}))
			_ = arrayIntrinsic_arg_12
			_ = arrayIntrinsic_arg_12
			_dollar___unused_10_10 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_17()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_18())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr_val_arrayMap13 := arrayIntrinsic_arg_12
					_ = arr_val_arrayMap13
					_ = arr_val_arrayMap13
					arr_go_arrayMap13 := arr_val_arrayMap13
					_ = arr_go_arrayMap13
					res_go_arrayMap13 := make([]int64, len(arr_go_arrayMap13))
					_ = res_go_arrayMap13
					for i_arrayMap13, v_arrayMap13 := range arr_go_arrayMap13 {
						res_go_arrayMap13[i_arrayMap13] = arrayIntrinsic_arg_11(v_arrayMap13)
					}
					return res_go_arrayMap13
				}(), []int64{int64(11), int64(20), int64(7)}}
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
			_ = _dollar___unused_10_10
			_dollar___unused_11_13 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_19()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_20())), func() gopurs_runtime.Value {
				orig := struct {
					actual []struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}
					expected []struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []struct {
							label    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							values []int64
							/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
						}{func() struct {
							label    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							values []int64
							/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
						} {
							orig := gopurs_runtime.RecordGet(input_2_2, "entry")
							_ = orig
							clone := struct {
								label    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								values []int64
								/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
							}{}
							clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
							clone.position = func() struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							} {
								orig := gopurs_runtime.RecordGet(orig, "position")
								_ = orig
								clone := struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}{}
								clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
								clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
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
						}()}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								orig := v
								_ = orig
								return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
									orig := orig.position
									_ = orig
									return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
								}(), func() gopurs_runtime.Value {
									arr := orig.values
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}())
							}()
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr))), func() gopurs_runtime.Value {
						orig := struct {
							label    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							values []int64
							/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
						}{"new", struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}{int64(19), int64(23)}, []int64{int64(29)}}
						_ = orig
						return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
							orig := orig.position
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
						}(), func() gopurs_runtime.Value {
							arr := orig.values
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}())
					}())).UnsafePtr)
					unboxed := make([]struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							label    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							values []int64
							/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
						} {
							orig := v
							_ = orig
							clone := struct {
								label    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								values []int64
								/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
							}{}
							clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
							clone.position = func() struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							} {
								orig := gopurs_runtime.RecordGet(orig, "position")
								_ = orig
								clone := struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}{}
								clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
								clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
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
					}
					return unboxed
				}(), func() []struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(input_2_2, "entry"), func() gopurs_runtime.Value {
						orig := struct {
							label    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							values []int64
							/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
						}{"new", struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}{int64(19), int64(23)}, []int64{int64(29)}}
						_ = orig
						return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
							orig := orig.position
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
						}(), func() gopurs_runtime.Value {
							arr := orig.values
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}())
					}()}).UnsafePtr)
					unboxed := make([]struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							label    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							values []int64
							/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
						} {
							orig := v
							_ = orig
							clone := struct {
								label    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								values []int64
								/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
							}{}
							clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
							clone.position = func() struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							} {
								orig := gopurs_runtime.RecordGet(orig, "position")
								_ = orig
								clone := struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}{}
								clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
								clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
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
					}
					return unboxed
				}()}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
								orig := orig.position
								_ = orig
								return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
							}(), func() gopurs_runtime.Value {
								arr := orig.values
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}())
						}()
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
								orig := orig.position
								_ = orig
								return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
							}(), func() gopurs_runtime.Value {
								arr := orig.values
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}())
						}()
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_11_13
			arrayIntrinsic_arg_15 := Get_Main_describe()
			_ = arrayIntrinsic_arg_15
			_ = arrayIntrinsic_arg_15
			arrayIntrinsic_arg_16 := gopurs_runtime.Apply2(Get_Main_appendItem(), func() gopurs_runtime.Value {
				arr := []*Constructor_Main_Item{(*Constructor_Main_Item)(nil), (&Constructor_Main_Item{1, func() struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				} {
					orig := gopurs_runtime.RecordGet(input_2_2, "entry")
					_ = orig
					clone := struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}{}
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.position = func() struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "position")
						_ = orig
						clone := struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}{}
						clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
						clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
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
				}(), (&Constructor_Data_List_Types_Cons[int64]{1, int64(31), (*Constructor_Data_List_Types_Cons[int64])(nil)})})}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer(v)}
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((&Constructor_Main_Item{1, struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			}{"new", struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}{int64(19), int64(23)}, []int64{int64(29)}}, (*Constructor_Data_List_Types_Cons[int64])(nil)}))})
			_ = arrayIntrinsic_arg_16
			_ = arrayIntrinsic_arg_16
			arrayIntrinsic_arg_17 := Get_Main_describe()
			_ = arrayIntrinsic_arg_17
			_ = arrayIntrinsic_arg_17
			arrayIntrinsic_arg_18 := []*Constructor_Main_Item{(*Constructor_Main_Item)(nil), (&Constructor_Main_Item{1, func() struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			} {
				orig := gopurs_runtime.RecordGet(input_2_2, "entry")
				_ = orig
				clone := struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}{}
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.position = func() struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "position")
					_ = orig
					clone := struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}{}
					clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
					clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
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
			}(), (&Constructor_Data_List_Types_Cons[int64]{1, int64(31), (*Constructor_Data_List_Types_Cons[int64])(nil)})}), (&Constructor_Main_Item{1, struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			}{"new", struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}{int64(19), int64(23)}, []int64{int64(29)}}, (*Constructor_Data_List_Types_Cons[int64])(nil)})}
			_ = arrayIntrinsic_arg_18
			_ = arrayIntrinsic_arg_18
			_dollar___unused_12_14 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_21()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_22())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []string
					expected []string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []string {
					arr_val_arrayMap15 := arrayIntrinsic_arg_16
					_ = arr_val_arrayMap15
					_ = arr_val_arrayMap15
					arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
					_ = arr_go_arrayMap15
					res_go_arrayMap15 := make([]string, len(*arr_go_arrayMap15))
					_ = res_go_arrayMap15
					for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
						res_go_arrayMap15[i_arrayMap15] = Call_Main_describe(gopurs_runtime.CoerceToStruct[Constructor_Main_Item](v_arrayMap15))
					}
					return res_go_arrayMap15
				}(), func() []string {
					arr_val_arrayMap15 := arrayIntrinsic_arg_18
					_ = arr_val_arrayMap15
					_ = arr_val_arrayMap15
					arr_go_arrayMap15 := arr_val_arrayMap15
					_ = arr_go_arrayMap15
					res_go_arrayMap15 := make([]string, len(arr_go_arrayMap15))
					_ = res_go_arrayMap15
					for i_arrayMap15, v_arrayMap15 := range arr_go_arrayMap15 {
						res_go_arrayMap15[i_arrayMap15] = Call_Main_describe(v_arrayMap15)
					}
					return res_go_arrayMap15
				}()}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_12_14
			arrayIntrinsic_arg_20 := Get_Main_describe()
			_ = arrayIntrinsic_arg_20
			_ = arrayIntrinsic_arg_20
			arrayIntrinsic_arg_21 := gopurs_runtime.Apply2(Get_Main_appendItem(), func() gopurs_runtime.Value {
				arr := []*Constructor_Main_Item{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer(v)}
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((*Constructor_Main_Item)(nil))})
			_ = arrayIntrinsic_arg_21
			_ = arrayIntrinsic_arg_21
			_dollar___unused_13_19 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_23()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_24())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []string
					expected []string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []string {
					arr_val_arrayMap16 := arrayIntrinsic_arg_21
					_ = arr_val_arrayMap16
					_ = arr_val_arrayMap16
					arr_go_arrayMap16 := (*[]gopurs_runtime.Value)(arr_val_arrayMap16.UnsafePtr)
					_ = arr_go_arrayMap16
					res_go_arrayMap16 := make([]string, len(*arr_go_arrayMap16))
					_ = res_go_arrayMap16
					for i_arrayMap16, v_arrayMap16 := range *arr_go_arrayMap16 {
						res_go_arrayMap16[i_arrayMap16] = Call_Main_describe(gopurs_runtime.CoerceToStruct[Constructor_Main_Item](v_arrayMap16))
					}
					return res_go_arrayMap16
				}(), []string{"empty"}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_13_19
			_dollar___unused_14_22 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{int64(2), int64(2)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_14_22
			partialRef_15_23 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Main_appendInt(), gopurs_runtime.RecordGet(input_2_2, "ints")), gopurs_runtime.Value{})
			_ = partialRef_15_23
			partial_16_24 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), partialRef_15_23, gopurs_runtime.Value{})
			_ = partial_16_24
			_dollar___unused_17_25 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_25()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_26())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(partial_16_24, gopurs_runtime.Int(int64(5))).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5)}}
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
			_ = _dollar___unused_17_25
			_dollar___unused_18_26 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_27()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_28())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(partial_16_24, gopurs_runtime.Int(int64(9))).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(9)}}
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
			_ = _dollar___unused_18_26
			aliasRef_19_27 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), Get_Data_Array_snoc(), gopurs_runtime.Value{})
			_ = aliasRef_19_27
			alias_20_28 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), aliasRef_19_27, gopurs_runtime.Value{})
			_ = alias_20_28
			_dollar___unused_21_29 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_29()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_30())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []string
					expected []string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Main_applyAppend(alias_20_28, (*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(input_2_2, "strings").UnsafePtr))), gopurs_runtime.StrWithCodeUnits("alias", 5))).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}(), []string{"", "λ😀", "alias"}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_21_29
			_dollar___unused_22_30 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_31()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_32())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []string
					expected []string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Main_applyAppend(alias_20_28, (*((*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []string{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr))), gopurs_runtime.StrWithCodeUnits("only", 4))).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}(), []string{"only"}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_22_30
			firstRef_23_31 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply2(Get_Main_appendInt(), gopurs_runtime.RecordGet(input_2_2, "ints"), gopurs_runtime.Int(int64(5))), gopurs_runtime.Value{})
			_ = firstRef_23_31
			first_24_32 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), firstRef_23_31, gopurs_runtime.Value{})
			_ = first_24_32
			_dollar___unused_25_33 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_33()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_34())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), first_24_32, gopurs_runtime.Int(int64(6))).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5), int64(6)}}
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
			_ = _dollar___unused_25_33
			_dollar___unused_26_34 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_35()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_36())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), first_24_32, gopurs_runtime.Int(int64(8))).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5), int64(8)}}
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
			_ = _dollar___unused_26_34
			_dollar___unused_27_35 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_37()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_38())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(first_24_32.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5)}}
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
			_ = _dollar___unused_27_35
			preserved_28_36 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), inputsRef_1_1, gopurs_runtime.Value{})
			_ = preserved_28_36
			_dollar___unused_29_37 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_39()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_40())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(preserved_28_36, "ints").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}(), []int64{int64(-2147483648), int64(0), int64(2147483647)}}
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
			_ = _dollar___unused_29_37
			_dollar___unused_30_38 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_41()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_42())), func() gopurs_runtime.Value {
				orig := struct {
					actual   [][]int64
					expected [][]int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() [][]int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(preserved_28_36, "arrays").UnsafePtr)
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
				}(), [][]int64{[]int64{}, []int64{int64(1), int64(2)}}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
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
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
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
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_30_38
			_dollar___unused_31_39 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_43()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_44())), func() gopurs_runtime.Value {
				orig := struct {
					actual   []string
					expected []string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(preserved_28_36, "strings").UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}(), []string{"", "λ😀"}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_31_39
			shared_32_40 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(10)), gopurs_runtime.Value{})
			_ = shared_32_40
			other_33_41 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(20)), gopurs_runtime.Value{})
			_ = other_33_41
			cellsRef_34_42 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				arr := []struct {
					cell  gopurs_runtime.Value
					label string
					/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
				}{struct {
					cell  gopurs_runtime.Value
					label string
					/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
				}{shared_32_40, "shared"}}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						orig := v
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, orig.cell, gopurs_runtime.Str(orig.label))
					}()
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Value{})
			_ = cellsRef_34_42
			cells_35_43 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), cellsRef_34_42, gopurs_runtime.Value{})
			_ = cells_35_43
			// TAST (Let): withOther_36_44 shape=App(Var) bindingType=(Array (Record (Row [label: String, cell: (ADT ["Effect","Ref","Ref"] [Int])] Empty)))
			withOther_36_44 := func() []struct {
				cell  gopurs_runtime.Value
				label string
				/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
			} {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(cells_35_43.UnsafePtr))), func() gopurs_runtime.Value {
					orig := struct {
						cell  gopurs_runtime.Value
						label string
						/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
					}{other_33_41, "other"}
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, orig.cell, gopurs_runtime.Str(orig.label))
				}())).UnsafePtr)
				unboxed := make([]struct {
					cell  gopurs_runtime.Value
					label string
					/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
				}, len(arr))
				for i, v := range arr {
					unboxed[i] = func() struct {
						cell  gopurs_runtime.Value
						label string
						/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
					} {
						orig := v
						_ = orig
						clone := struct {
							cell  gopurs_runtime.Value
							label string
							/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
						}{}
						clone.cell = gopurs_runtime.RecordGet(orig, "cell")
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
				}
				return unboxed
			}()
			_ = withOther_36_44
			var __t47 gopurs_runtime.Value
			{
				if (gopurs_runtime.Int(int64(len(withOther_36_44))).IntVal) == (int64(2)) {
					__t47 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						_dollar___unused_37_46 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(41)), gopurs_runtime.RecordGet(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
							arr := withOther_36_44
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, orig.cell, gopurs_runtime.Str(orig.label))
								}()
							}
							return gopurs_runtime.Array(boxed)
						}(), 0), "cell"), gopurs_runtime.Value{})
						_ = _dollar___unused_37_46
						return gopurs_runtime.Apply3(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(42)), gopurs_runtime.RecordGet(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
							arr := withOther_36_44
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, orig.cell, gopurs_runtime.Str(orig.label))
								}()
							}
							return gopurs_runtime.Array(boxed)
						}(), 1), "cell"), gopurs_runtime.Value{})
					})
					goto end_branch_47
				} else {

				}
			}
			{
				__t47 = gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
					orig := struct {
						actual   int64
						expected int64
						/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
					}{gopurs_runtime.Int(int64(len(withOther_36_44))).IntVal, int64(2)}
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
				}())
			}
		end_branch_47:
			_dollar___unused_37_45 := gopurs_runtime.Apply(__t47, gopurs_runtime.Value{})
			_ = _dollar___unused_37_45
			sharedValue_38_48 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), shared_32_40, gopurs_runtime.Value{})
			_ = sharedValue_38_48
			otherValue_39_49 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), other_33_41, gopurs_runtime.Value{})
			_ = otherValue_39_49
			_dollar___unused_40_50 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{sharedValue_38_48.IntVal, int64(41)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_40_50
			_dollar___unused_41_51 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{otherValue_39_49.IntVal, int64(42)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_41_51
			_dollar___unused_42_52 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(cells_35_43))).IntVal, int64(1)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_42_52
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_687527510_3790796878(Rebox_Main_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))})
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))})
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))})
	})
	return cache_Main___cached_dict_11
}

var cache_Main___cached_dict_12 gopurs_runtime.Value
var once_Main___cached_dict_12 sync.Once

func Get_Main___cached_dict_12() gopurs_runtime.Value {
	once_Main___cached_dict_12.Do(func() {
		cache_Main___cached_dict_12 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))})
	})
	return cache_Main___cached_dict_12
}

var cache_Main___cached_dict_13 gopurs_runtime.Value
var once_Main___cached_dict_13 sync.Once

func Get_Main___cached_dict_13() gopurs_runtime.Value {
	once_Main___cached_dict_13.Do(func() {
		cache_Main___cached_dict_13 = Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))
	})
	return cache_Main___cached_dict_13
}

var cache_Main___cached_dict_14 gopurs_runtime.Value
var once_Main___cached_dict_14 sync.Once

func Get_Main___cached_dict_14() gopurs_runtime.Value {
	once_Main___cached_dict_14.Do(func() {
		cache_Main___cached_dict_14 = Call_Data_Show_showArray(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))
	})
	return cache_Main___cached_dict_14
}

var cache_Main___cached_dict_15 gopurs_runtime.Value
var once_Main___cached_dict_15 sync.Once

func Get_Main___cached_dict_15() gopurs_runtime.Value {
	once_Main___cached_dict_15.Do(func() {
		cache_Main___cached_dict_15 = Call_Data_Eq_eqArray(Call_Data_List_Types_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))
	})
	return cache_Main___cached_dict_15
}

var cache_Main___cached_dict_16 gopurs_runtime.Value
var once_Main___cached_dict_16 sync.Once

func Get_Main___cached_dict_16() gopurs_runtime.Value {
	once_Main___cached_dict_16.Do(func() {
		cache_Main___cached_dict_16 = Call_Data_Show_showArray(Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))
	})
	return cache_Main___cached_dict_16
}

var cache_Main___cached_dict_17 gopurs_runtime.Value
var once_Main___cached_dict_17 sync.Once

func Get_Main___cached_dict_17() gopurs_runtime.Value {
	once_Main___cached_dict_17.Do(func() {
		cache_Main___cached_dict_17 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_17
}

var cache_Main___cached_dict_18 gopurs_runtime.Value
var once_Main___cached_dict_18 sync.Once

func Get_Main___cached_dict_18() gopurs_runtime.Value {
	once_Main___cached_dict_18.Do(func() {
		cache_Main___cached_dict_18 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_18
}

var cache_Main___cached_dict_19 gopurs_runtime.Value
var once_Main___cached_dict_19 sync.Once

func Get_Main___cached_dict_19() gopurs_runtime.Value {
	once_Main___cached_dict_19.Do(func() {
		cache_Main___cached_dict_19 = Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol()), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_positionIsSymbol()), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), gopurs_runtime.Value{}, Get_Main_labelIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))
	})
	return cache_Main___cached_dict_19
}

var cache_Main___cached_dict_20 gopurs_runtime.Value
var once_Main___cached_dict_20 sync.Once

func Get_Main___cached_dict_20() gopurs_runtime.Value {
	once_Main___cached_dict_20.Do(func() {
		cache_Main___cached_dict_20 = Call_Data_Show_showArray(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_labelIsSymbol()), gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_positionIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valuesIsSymbol()), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))
	})
	return cache_Main___cached_dict_20
}

var cache_Main___cached_dict_21 gopurs_runtime.Value
var once_Main___cached_dict_21 sync.Once

func Get_Main___cached_dict_21() gopurs_runtime.Value {
	once_Main___cached_dict_21.Do(func() {
		cache_Main___cached_dict_21 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main___cached_dict_21
}

var cache_Main___cached_dict_22 gopurs_runtime.Value
var once_Main___cached_dict_22 sync.Once

func Get_Main___cached_dict_22() gopurs_runtime.Value {
	once_Main___cached_dict_22.Do(func() {
		cache_Main___cached_dict_22 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_22
}

var cache_Main___cached_dict_23 gopurs_runtime.Value
var once_Main___cached_dict_23 sync.Once

func Get_Main___cached_dict_23() gopurs_runtime.Value {
	once_Main___cached_dict_23.Do(func() {
		cache_Main___cached_dict_23 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main___cached_dict_23
}

var cache_Main___cached_dict_24 gopurs_runtime.Value
var once_Main___cached_dict_24 sync.Once

func Get_Main___cached_dict_24() gopurs_runtime.Value {
	once_Main___cached_dict_24.Do(func() {
		cache_Main___cached_dict_24 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_24
}

var cache_Main___cached_dict_25 gopurs_runtime.Value
var once_Main___cached_dict_25 sync.Once

func Get_Main___cached_dict_25() gopurs_runtime.Value {
	once_Main___cached_dict_25.Do(func() {
		cache_Main___cached_dict_25 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_25
}

var cache_Main___cached_dict_26 gopurs_runtime.Value
var once_Main___cached_dict_26 sync.Once

func Get_Main___cached_dict_26() gopurs_runtime.Value {
	once_Main___cached_dict_26.Do(func() {
		cache_Main___cached_dict_26 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_26
}

var cache_Main___cached_dict_27 gopurs_runtime.Value
var once_Main___cached_dict_27 sync.Once

func Get_Main___cached_dict_27() gopurs_runtime.Value {
	once_Main___cached_dict_27.Do(func() {
		cache_Main___cached_dict_27 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_27
}

var cache_Main___cached_dict_28 gopurs_runtime.Value
var once_Main___cached_dict_28 sync.Once

func Get_Main___cached_dict_28() gopurs_runtime.Value {
	once_Main___cached_dict_28.Do(func() {
		cache_Main___cached_dict_28 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_28
}

var cache_Main___cached_dict_29 gopurs_runtime.Value
var once_Main___cached_dict_29 sync.Once

func Get_Main___cached_dict_29() gopurs_runtime.Value {
	once_Main___cached_dict_29.Do(func() {
		cache_Main___cached_dict_29 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main___cached_dict_29
}

var cache_Main___cached_dict_30 gopurs_runtime.Value
var once_Main___cached_dict_30 sync.Once

func Get_Main___cached_dict_30() gopurs_runtime.Value {
	once_Main___cached_dict_30.Do(func() {
		cache_Main___cached_dict_30 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_30
}

var cache_Main___cached_dict_31 gopurs_runtime.Value
var once_Main___cached_dict_31 sync.Once

func Get_Main___cached_dict_31() gopurs_runtime.Value {
	once_Main___cached_dict_31.Do(func() {
		cache_Main___cached_dict_31 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main___cached_dict_31
}

var cache_Main___cached_dict_32 gopurs_runtime.Value
var once_Main___cached_dict_32 sync.Once

func Get_Main___cached_dict_32() gopurs_runtime.Value {
	once_Main___cached_dict_32.Do(func() {
		cache_Main___cached_dict_32 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_32
}

var cache_Main___cached_dict_33 gopurs_runtime.Value
var once_Main___cached_dict_33 sync.Once

func Get_Main___cached_dict_33() gopurs_runtime.Value {
	once_Main___cached_dict_33.Do(func() {
		cache_Main___cached_dict_33 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_33
}

var cache_Main___cached_dict_34 gopurs_runtime.Value
var once_Main___cached_dict_34 sync.Once

func Get_Main___cached_dict_34() gopurs_runtime.Value {
	once_Main___cached_dict_34.Do(func() {
		cache_Main___cached_dict_34 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_34
}

var cache_Main___cached_dict_35 gopurs_runtime.Value
var once_Main___cached_dict_35 sync.Once

func Get_Main___cached_dict_35() gopurs_runtime.Value {
	once_Main___cached_dict_35.Do(func() {
		cache_Main___cached_dict_35 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_35
}

var cache_Main___cached_dict_36 gopurs_runtime.Value
var once_Main___cached_dict_36 sync.Once

func Get_Main___cached_dict_36() gopurs_runtime.Value {
	once_Main___cached_dict_36.Do(func() {
		cache_Main___cached_dict_36 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_36
}

var cache_Main___cached_dict_37 gopurs_runtime.Value
var once_Main___cached_dict_37 sync.Once

func Get_Main___cached_dict_37() gopurs_runtime.Value {
	once_Main___cached_dict_37.Do(func() {
		cache_Main___cached_dict_37 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_37
}

var cache_Main___cached_dict_38 gopurs_runtime.Value
var once_Main___cached_dict_38 sync.Once

func Get_Main___cached_dict_38() gopurs_runtime.Value {
	once_Main___cached_dict_38.Do(func() {
		cache_Main___cached_dict_38 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_38
}

var cache_Main___cached_dict_39 gopurs_runtime.Value
var once_Main___cached_dict_39 sync.Once

func Get_Main___cached_dict_39() gopurs_runtime.Value {
	once_Main___cached_dict_39.Do(func() {
		cache_Main___cached_dict_39 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_39
}

var cache_Main___cached_dict_40 gopurs_runtime.Value
var once_Main___cached_dict_40 sync.Once

func Get_Main___cached_dict_40() gopurs_runtime.Value {
	once_Main___cached_dict_40.Do(func() {
		cache_Main___cached_dict_40 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_40
}

var cache_Main___cached_dict_41 gopurs_runtime.Value
var once_Main___cached_dict_41 sync.Once

func Get_Main___cached_dict_41() gopurs_runtime.Value {
	once_Main___cached_dict_41.Do(func() {
		cache_Main___cached_dict_41 = Call_Data_Eq_eqArray(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))
	})
	return cache_Main___cached_dict_41
}

var cache_Main___cached_dict_42 gopurs_runtime.Value
var once_Main___cached_dict_42 sync.Once

func Get_Main___cached_dict_42() gopurs_runtime.Value {
	once_Main___cached_dict_42.Do(func() {
		cache_Main___cached_dict_42 = Call_Data_Show_showArray(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))
	})
	return cache_Main___cached_dict_42
}

var cache_Main___cached_dict_43 gopurs_runtime.Value
var once_Main___cached_dict_43 sync.Once

func Get_Main___cached_dict_43() gopurs_runtime.Value {
	once_Main___cached_dict_43.Do(func() {
		cache_Main___cached_dict_43 = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main___cached_dict_43
}

var cache_Main___cached_dict_44 gopurs_runtime.Value
var once_Main___cached_dict_44 sync.Once

func Get_Main___cached_dict_44() gopurs_runtime.Value {
	once_Main___cached_dict_44.Do(func() {
		cache_Main___cached_dict_44 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_44
}

type Constructor_Main_Empty struct {
	Rc uint32
}

type Constructor_Main_Item struct {
	Rc uint32
	V0 struct {
		label    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		values []int64
		/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
	}
	V1 *Constructor_Data_List_Types_Cons[int64]
}

func Call_Main_describe(v_0_loop *Constructor_Main_Item) string {
	var v_0 *Constructor_Main_Item = v_0_loop
	_ = v_0
	var __t0 string
	{
		if v_0 == nil {
			__t0 = "empty"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 != nil {
			__t0 = gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString((v_0).V0.label, ":"), gopurs_runtime.Apply(Call_Data_Show_show(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_0())), func() gopurs_runtime.Value {
				orig := (v_0).V0.position
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
			}()).StrVal()), ":"), gopurs_runtime.Apply(Call_Data_Show_show(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), func() gopurs_runtime.Value {
				arr := (v_0).V0.values
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).StrVal()), ":"), gopurs_runtime.Apply(Call_Data_Show_show(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_2())), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((v_0).V1))}).StrVal())
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_applyAppend(append1_0_loop gopurs_runtime.Value, values_1_loop []gopurs_runtime.Value, value_2_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var append1_0 gopurs_runtime.Value = append1_0_loop
	_ = append1_0
	var values_1 []gopurs_runtime.Value = values_1_loop
	_ = values_1
	var value_2 gopurs_runtime.Value = value_2_loop
	_ = value_2
	return (*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(append1_0, gopurs_runtime.Array(values_1), value_2).UnsafePtr)))
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

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[float64])(unsafe.Pointer(in))
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

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3704040722_849153993(in *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = Rebox_Main_3704040722_849153993(in.V1)
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

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_687527510(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[float64])(unsafe.Pointer(in))
}

func Rebox_Main_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_849153993_3704040722(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Main_849153993_3704040722(in.V1)
	return out
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys1{K0: "reflectSymbol"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys3{K0: "label", K1: "position", K2: "values"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys2{K0: "column", K1: "line"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}

var GopursRecordKeys_4d61696e_4 = gopurs_runtime.RecordKeys2{K0: "cell", K1: "label"}
