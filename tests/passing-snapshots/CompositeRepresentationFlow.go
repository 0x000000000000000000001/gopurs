package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = Get_Control_Bind_bind__1741690964()
	})
	return cache_Main_bind
}

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = Get_Control_Bind_bind__2695808948()
	})
	return cache_Main_bind1
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__3137154580()
	})
	return cache_Main_bind2
}

var cache_Main_discard gopurs_runtime.Value
var once_Main_discard sync.Once

func Get_Main_discard() gopurs_runtime.Value {
	once_Main_discard.Do(func() {
		cache_Main_discard = Get_Control_Bind_discard__143037513()
	})
	return cache_Main_discard
}

var cache_Main_assertEqual gopurs_runtime.Value
var once_Main_assertEqual sync.Once

func Get_Main_assertEqual() gopurs_runtime.Value {
	once_Main_assertEqual.Do(func() {
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__2867988789()
	})
	return cache_Main_assertEqual
}

var cache_Main_assertEqual1 gopurs_runtime.Value
var once_Main_assertEqual1 sync.Once

func Get_Main_assertEqual1() gopurs_runtime.Value {
	once_Main_assertEqual1.Do(func() {
		cache_Main_assertEqual1 = Get_Test_Assert_assertEqual__2127306453()
	})
	return cache_Main_assertEqual1
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
		cache_Main_eqRowCons = gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRowCons
}

var cache_Main_streamIsSymbol gopurs_runtime.Value
var once_Main_streamIsSymbol sync.Once

func Get_Main_streamIsSymbol() gopurs_runtime.Value {
	once_Main_streamIsSymbol.Do(func() {
		cache_Main_streamIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("stream", 6)
		}))
	})
	return cache_Main_streamIsSymbol
}

var cache_Main_eqList gopurs_runtime.Value
var once_Main_eqList sync.Once

func Get_Main_eqList() gopurs_runtime.Value {
	once_Main_eqList.Do(func() {
		cache_Main_eqList = Call_Data_List_Types_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main_eqList
}

var cache_Main_errorsIsSymbol gopurs_runtime.Value
var once_Main_errorsIsSymbol sync.Once

func Get_Main_errorsIsSymbol() gopurs_runtime.Value {
	once_Main_errorsIsSymbol.Do(func() {
		cache_Main_errorsIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("errors", 6)
		}))
	})
	return cache_Main_errorsIsSymbol
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

var cache_Main_errorIsSymbol gopurs_runtime.Value
var once_Main_errorIsSymbol sync.Once

func Get_Main_errorIsSymbol() gopurs_runtime.Value {
	once_Main_errorIsSymbol.Do(func() {
		cache_Main_errorIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("error", 5)
		}))
	})
	return cache_Main_errorIsSymbol
}

var cache_Main_eqRec1 gopurs_runtime.Value
var once_Main_eqRec1 sync.Once

func Get_Main_eqRec1() gopurs_runtime.Value {
	once_Main_eqRec1.Do(func() {
		cache_Main_eqRec1 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_positionIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), gopurs_runtime.Value{}, Get_Main_errorIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))
	})
	return cache_Main_eqRec1
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = Call_Data_Eq_eqArray(Get_Main_eqRec1())
	})
	return cache_Main_eqArray
}

var cache_Main_consumedIsSymbol gopurs_runtime.Value
var once_Main_consumedIsSymbol sync.Once

func Get_Main_consumedIsSymbol() gopurs_runtime.Value {
	once_Main_consumedIsSymbol.Do(func() {
		cache_Main_consumedIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("consumed", 8)
		}))
	})
	return cache_Main_consumedIsSymbol
}

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Apply2(Get_Data_Show_showRecord(), gopurs_runtime.Value{}, gopurs_runtime.Value{})
	})
	return cache_Main_showRecord
}

var cache_Main_showRecordFieldsCons gopurs_runtime.Value
var once_Main_showRecordFieldsCons sync.Once

func Get_Main_showRecordFieldsCons() gopurs_runtime.Value {
	once_Main_showRecordFieldsCons.Do(func() {
		cache_Main_showRecordFieldsCons = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsCons(), Get_Main_consumedIsSymbol())
	})
	return cache_Main_showRecordFieldsCons
}

var cache_Main_showList gopurs_runtime.Value
var once_Main_showList sync.Once

func Get_Main_showList() gopurs_runtime.Value {
	once_Main_showList.Do(func() {
		cache_Main_showList = Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main_showList
}

var cache_Main_showRecord1 gopurs_runtime.Value
var once_Main_showRecord1 sync.Once

func Get_Main_showRecord1() gopurs_runtime.Value {
	once_Main_showRecord1.Do(func() {
		cache_Main_showRecord1 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_errorIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_positionIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}))
	})
	return cache_Main_showRecord1
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = Call_Data_Show_showArray(Get_Main_showRecord1())
	})
	return cache_Main_showArray
}

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = Get_Test_Assert_assertEqual__3132330058()
	})
	return cache_Main_assertEqual2
}

var cache_Main_assertEqual3 gopurs_runtime.Value
var once_Main_assertEqual3 sync.Once

func Get_Main_assertEqual3() gopurs_runtime.Value {
	once_Main_assertEqual3.Do(func() {
		cache_Main_assertEqual3 = Get_Test_Assert_assertEqual__1511484059()
	})
	return cache_Main_assertEqual3
}

var cache_Main_assertEqual4 gopurs_runtime.Value
var once_Main_assertEqual4 sync.Once

func Get_Main_assertEqual4() gopurs_runtime.Value {
	once_Main_assertEqual4.Do(func() {
		cache_Main_assertEqual4 = Get_Test_Assert_assertEqual__1105254914()
	})
	return cache_Main_assertEqual4
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__3909457492()
	})
	return cache_Main_bind3
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__3285793908()
	})
	return cache_Main_bind4
}

var cache_Main_assertEqual5 gopurs_runtime.Value
var once_Main_assertEqual5 sync.Once

func Get_Main_assertEqual5() gopurs_runtime.Value {
	once_Main_assertEqual5.Do(func() {
		cache_Main_assertEqual5 = Get_Test_Assert_assertEqual__1435614722()
	})
	return cache_Main_assertEqual5
}

var cache_Main_extraIsSymbol gopurs_runtime.Value
var once_Main_extraIsSymbol sync.Once

func Get_Main_extraIsSymbol() gopurs_runtime.Value {
	once_Main_extraIsSymbol.Do(func() {
		cache_Main_extraIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("extra", 5)
		}))
	})
	return cache_Main_extraIsSymbol
}

var cache_Main_assertEqual6 gopurs_runtime.Value
var once_Main_assertEqual6 sync.Once

func Get_Main_assertEqual6() gopurs_runtime.Value {
	once_Main_assertEqual6.Do(func() {
		cache_Main_assertEqual6 = Get_Test_Assert_assertEqual__3267767652()
	})
	return cache_Main_assertEqual6
}

var cache_Main_throughCallback gopurs_runtime.Value
var once_Main_throughCallback sync.Once

func Get_Main_throughCallback() gopurs_runtime.Value {
	once_Main_throughCallback.Do(func() {
		cache_Main_throughCallback = gopurs_runtime.Func2(func(transform_0_box gopurs_runtime.Value, original_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_throughCallback(transform_0_box, func() struct {
					consumed bool
					errors   []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}
					stream *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
				} {
					orig := original_1_box
					_ = orig
					clone := struct {
						consumed bool
						errors   []struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						}
						stream *Constructor_Data_List_Types_Cons[int64]
						/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
					}{}
					clone.consumed = (gopurs_runtime.RecordGet(orig, "consumed").IntVal) != (0)
					clone.errors = func() []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					} {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "errors").UnsafePtr)
						unboxed := make([]struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						}, len(arr))
						for i, v := range arr {
							unboxed[i] = func() struct {
								error    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
							} {
								orig := v
								_ = orig
								clone := struct {
									error    string
									position struct {
										column int64
										line   int64
										/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
									}
									/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
								}{}
								clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
								return clone
							}()
						}
						return unboxed
					}()
					clone.stream = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "stream")))
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.consumed), func() gopurs_runtime.Value {
					arr := orig.errors
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.error), func() gopurs_runtime.Value {
								orig := orig.position
								_ = orig
								return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
							}())
						}()
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(orig.stream))})
			}()
		})
	})
	return cache_Main_throughCallback
}

var cache_Main_setConsumed gopurs_runtime.Value
var once_Main_setConsumed sync.Once

func Get_Main_setConsumed() gopurs_runtime.Value {
	once_Main_setConsumed.Do(func() {
		cache_Main_setConsumed = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_setConsumed(row_0_box)
		})
	})
	return cache_Main_setConsumed
}

var cache_Main_setConsumed__1062648954 gopurs_runtime.Value
var once_Main_setConsumed__1062648954 sync.Once

func Get_Main_setConsumed__1062648954() gopurs_runtime.Value {
	once_Main_setConsumed__1062648954.Do(func() {
		cache_Main_setConsumed__1062648954 = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_setConsumed__1062648954(row_0_box)
		})
	})
	return cache_Main_setConsumed__1062648954
}

var cache_Main_nativeResult gopurs_runtime.Value
var once_Main_nativeResult sync.Once

func Get_Main_nativeResult() gopurs_runtime.Value {
	once_Main_nativeResult.Do(func() {
		cache_Main_nativeResult = gopurs_runtime.Func(func(state_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3162159231_138441832(Call_Main_nativeResult(func() struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			} {
				orig := state_0_box
				_ = orig
				clone := struct {
					consumed bool
					errors   []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}
					stream *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
				}{}
				clone.consumed = (gopurs_runtime.RecordGet(orig, "consumed").IntVal) != (0)
				clone.errors = func() []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "errors").UnsafePtr)
					unboxed := make([]struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						} {
							orig := v
							_ = orig
							clone := struct {
								error    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
							}{}
							clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
							return clone
						}()
					}
					return unboxed
				}()
				clone.stream = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "stream")))
				return clone
			}())))}
		})
	})
	return cache_Main_nativeResult
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(state_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordUpdate2(state_0, "consumed", gopurs_runtime.Bool(true), "errors", func() gopurs_runtime.Value {
					arr := Call_Data_Array_snoc__1021931931(func() []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					} {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(state_0, "errors").UnsafePtr)
						unboxed := make([]struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						}, len(arr))
						for i, v := range arr {
							unboxed[i] = func() struct {
								error    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
							} {
								orig := v
								_ = orig
								clone := struct {
									error    string
									position struct {
										column int64
										line   int64
										/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
									}
									/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
								}{}
								clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
								return clone
							}()
						}
						return unboxed
					}(), struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}{"λ😀", struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}{int64(-2147483648), int64(2147483647)}})
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.error), func() gopurs_runtime.Value {
								orig := orig.position
								_ = orig
								return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
							}())
						}()
					}
					return gopurs_runtime.Array(boxed)
				}())
			}), func(state_0 struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			}) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					orig := func() struct {
						consumed bool
						errors   []struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						}
						stream *Constructor_Data_List_Types_Cons[int64]
						/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
					} {
						clone := state_0
						clone.consumed = true
						clone.errors = Call_Data_Array_snoc__1021931931(state_0.errors, struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						}{"λ😀", struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}{int64(-2147483648), int64(2147483647)}})
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.consumed), func() gopurs_runtime.Value {
						arr := orig.errors
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								orig := v
								_ = orig
								return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.error), func() gopurs_runtime.Value {
									orig := orig.position
									_ = orig
									return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
								}())
							}()
						}
						return gopurs_runtime.Array(boxed)
					}(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(orig.stream))})
				}()
			}), gopurs_runtime.Value{})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__callback_fn_3 := Get_Effect_Ref__new()
			_ = __callback_fn_3
			__callback_arg_4 := struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			}{false, []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}{}, (&Constructor_Data_List_Types_Cons[int64]{1, int64(7), (&Constructor_Data_List_Types_Cons[int64]{1, int64(9), (*Constructor_Data_List_Types_Cons[int64])(nil)})})}
			_ = __callback_arg_4
			__local_var_2_2 := gopurs_runtime.Apply(func(__native_callback func(struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
				if __native_callback_ok {
					return __native_callback(__callback_arg_4)
				} else {

				}
				return gopurs_runtime.Apply(__callback_fn_3, func() gopurs_runtime.Value {
					orig := __callback_arg_4
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.consumed), func() gopurs_runtime.Value {
						arr := orig.errors
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								orig := v
								_ = orig
								return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.error), func() gopurs_runtime.Value {
									orig := orig.position
									_ = orig
									return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
								}())
							}()
						}
						return gopurs_runtime.Array(boxed)
					}(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(orig.stream))})
				}())
			}(gopurs_runtime.FunctionData[func(struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			}) gopurs_runtime.Value](__callback_fn_3)), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_5 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_2_2, gopurs_runtime.Value{})
			_ = __local_var_3_5
			// TAST (Let): result_4_6 shape=App(Var) bindingType=(Record (Row [consumed: Boolean, errors: (Array (Record (Row [error: String, position: (Record (Row [column: Int, line: Int] Empty))] Empty))), stream: (ADT ["Data","List","Types","List"] [Int])] Empty))
			result_4_6 := Call_Main_throughCallback(__local_var_1_1, func() struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			} {
				orig := __local_var_3_5
				_ = orig
				clone := struct {
					consumed bool
					errors   []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}
					stream *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
				}{}
				clone.consumed = (gopurs_runtime.RecordGet(orig, "consumed").IntVal) != (0)
				clone.errors = func() []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "errors").UnsafePtr)
					unboxed := make([]struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						} {
							orig := v
							_ = orig
							clone := struct {
								error    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
							}{}
							clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
							return clone
						}()
					}
					return unboxed
				}()
				clone.stream = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "stream")))
				return clone
			}())
			_ = result_4_6
			__local_var_5_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Int(int64(len(result_4_6.errors))).IntVal, int64(2)}), gopurs_runtime.Value{})
			_ = __local_var_5_7
			__local_var_6_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{result_4_6.consumed, true}), gopurs_runtime.Value{})
			_ = __local_var_6_8
			__local_var_7_9 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_2_2, gopurs_runtime.Value{})
			_ = __local_var_7_9
			__local_var_8_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__3132330058(struct {
				actual struct {
					consumed bool
					errors   []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}
					stream *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
				}
				expected struct {
					consumed bool
					errors   []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}
					stream *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			} {
				orig := __local_var_7_9
				_ = orig
				clone := struct {
					consumed bool
					errors   []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}
					stream *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
				}{}
				clone.consumed = (gopurs_runtime.RecordGet(orig, "consumed").IntVal) != (0)
				clone.errors = func() []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "errors").UnsafePtr)
					unboxed := make([]struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						} {
							orig := v
							_ = orig
							clone := struct {
								error    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
							}{}
							clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
							return clone
						}()
					}
					return unboxed
				}()
				clone.stream = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "stream")))
				return clone
			}(), func() struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			} {
				orig := __local_var_3_5
				_ = orig
				clone := struct {
					consumed bool
					errors   []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}
					stream *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
				}{}
				clone.consumed = (gopurs_runtime.RecordGet(orig, "consumed").IntVal) != (0)
				clone.errors = func() []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "errors").UnsafePtr)
					unboxed := make([]struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						} {
							orig := v
							_ = orig
							clone := struct {
								error    string
								position struct {
									column int64
									line   int64
									/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
								}
								/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
							}{}
							clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
							return clone
						}()
					}
					return unboxed
				}()
				clone.stream = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "stream")))
				return clone
			}()}), gopurs_runtime.Value{})
			_ = __local_var_8_10
			__local_var_9_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.RecordGet(__local_var_3_5, "consumed").IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_9_11
			__local_var_10_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___678702677(Rebox_Main_1386611502_1510340014(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_0())), "", struct {
				actual []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				expected []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			} {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_3_5, "errors").UnsafePtr)
				unboxed := make([]struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}, len(arr))
				for i, v := range arr {
					unboxed[i] = func() struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					} {
						orig := v
						_ = orig
						clone := struct {
							error    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
						}{}
						clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
						return clone
					}()
				}
				return unboxed
			}(), []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}{}}), gopurs_runtime.Value{})
			_ = __local_var_10_12
			__local_var_11_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3457568413(Rebox_Main_1386611502_284210452(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), "", struct {
				actual *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}]
				expected *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_nativeResult(result_4_6), (&Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}]{1, result_4_6.stream, result_4_6.errors})}), gopurs_runtime.Value{})
			_ = __local_var_11_13
			__local_var_12_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3457568413(Rebox_Main_1386611502_284210452(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_2())), "", struct {
				actual *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}]
				expected *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_nativeResult(struct {
				consumed bool
				errors   []struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				}
				stream *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
			}{false, []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}{}, (*Constructor_Data_List_Types_Cons[int64])(nil)}), (&Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}]{1, Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})), []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}{}})}), gopurs_runtime.Value{})
			_ = __local_var_12_14
			__local_var_13_15 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(flag_13 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordUpdate1(__local_var_3_5, "consumed", gopurs_runtime.Bool((flag_13.IntVal) != (0)))
			}), gopurs_runtime.Value{})
			_ = __local_var_13_15
			__local_var_14_16 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_13_15, gopurs_runtime.Value{})
			_ = __local_var_14_16
			__local_var_15_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_14_16, gopurs_runtime.Bool(true)), "consumed").IntVal) != (0), true}), gopurs_runtime.Value{})
			_ = __local_var_15_17
			__local_var_16_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_14_16, gopurs_runtime.Bool(false)), "consumed").IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_16_18
			__local_var_17_19 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1506863018(Rebox_Main_1386611502_2622893113(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())), "", struct {
				actual   *Constructor_Data_List_Types_Cons[int64]
				expected *Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_14_16, gopurs_runtime.Bool(true)), "stream"))), Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(__local_var_3_5, "stream")))}), gopurs_runtime.Value{})
			_ = __local_var_17_19
			__local_var_18_20 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__3267767652(struct {
				actual struct {
					consumed bool
					extra    string
					/*gopurs:record-v1:2:0063006f006e00730075006d00650064,00650078007400720061*/
				}
				expected struct {
					consumed bool
					extra    string
					/*gopurs:record-v1:2:0063006f006e00730075006d00650064,00650078007400720061*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{struct {
				consumed bool
				extra    string
				/*gopurs:record-v1:2:0063006f006e00730075006d00650064,00650078007400720061*/
			}{true, "preserved"}, struct {
				consumed bool
				extra    string
				/*gopurs:record-v1:2:0063006f006e00730075006d00650064,00650078007400720061*/
			}{true, "preserved"}}), gopurs_runtime.Value{})
			_ = __local_var_18_20
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_Show_showArray(Get_Main_showRecord1())
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = Call_Data_Tuple_showTuple(Get_Main_showList(), Get_Main_showArray())
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = Call_Data_Tuple_showTuple(Get_Main_showList(), Get_Main_showArray())
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_3
}

func Call_Main_throughCallback(transform_0_loop gopurs_runtime.Value, original_1_loop struct {
	consumed bool
	errors   []struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	}
	stream *Constructor_Data_List_Types_Cons[int64]
	/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
}) struct {
	consumed bool
	errors   []struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	}
	stream *Constructor_Data_List_Types_Cons[int64]
	/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
} {
	var transform_0 gopurs_runtime.Value = transform_0_loop
	_ = transform_0
	var original_1 struct {
		consumed bool
		errors   []struct {
			error    string
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
		}
		stream *Constructor_Data_List_Types_Cons[int64]
		/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
	} = original_1_loop
	_ = original_1
	__callback_fn_0 := transform_0
	_ = __callback_fn_0
	__callback_fn_1 := transform_0
	_ = __callback_fn_1
	__callback_arg_2 := original_1
	_ = __callback_arg_2
	__callback_arg_3 := func(__native_callback func(struct {
		consumed bool
		errors   []struct {
			error    string
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
		}
		stream *Constructor_Data_List_Types_Cons[int64]
		/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
	}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
		if __native_callback_ok {
			return __native_callback(__callback_arg_2)
		} else {

		}
		return gopurs_runtime.Apply(__callback_fn_1, func() gopurs_runtime.Value {
			orig := __callback_arg_2
			_ = orig
			return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.consumed), func() gopurs_runtime.Value {
				arr := orig.errors
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						orig := v
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.error), func() gopurs_runtime.Value {
							orig := orig.position
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
						}())
					}()
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(orig.stream))})
		}())
	}(gopurs_runtime.FunctionData[func(struct {
		consumed bool
		errors   []struct {
			error    string
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
		}
		stream *Constructor_Data_List_Types_Cons[int64]
		/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
	}) gopurs_runtime.Value](__callback_fn_1))
	_ = __callback_arg_3
	return func() struct {
		consumed bool
		errors   []struct {
			error    string
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
		}
		stream *Constructor_Data_List_Types_Cons[int64]
		/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
	} {
		orig := gopurs_runtime.Apply(__callback_fn_0, __callback_arg_3)
		_ = orig
		clone := struct {
			consumed bool
			errors   []struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}
			stream *Constructor_Data_List_Types_Cons[int64]
			/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
		}{}
		clone.consumed = (gopurs_runtime.RecordGet(orig, "consumed").IntVal) != (0)
		clone.errors = func() []struct {
			error    string
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
		} {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "errors").UnsafePtr)
			unboxed := make([]struct {
				error    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
			}, len(arr))
			for i, v := range arr {
				unboxed[i] = func() struct {
					error    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
				} {
					orig := v
					_ = orig
					clone := struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					}{}
					clone.error = gopurs_runtime.RecordGet(orig, "error").StrVal()
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
					return clone
				}()
			}
			return unboxed
		}()
		clone.stream = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "stream")))
		return clone
	}()
}

func Call_Main_setConsumed(row_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var row_0 gopurs_runtime.Value = row_0_loop
	_ = row_0
	return gopurs_runtime.RecordUpdate1(row_0, "consumed", gopurs_runtime.Bool(true))
}

func Call_Main_setConsumed__1062648954(row_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
setConsumed__1062648954:
	for {
		if false {
			continue setConsumed__1062648954
		}
		var row_0 gopurs_runtime.Value = row_0_loop
		_ = row_0
		return gopurs_runtime.RecordUpdate1(row_0, "consumed", gopurs_runtime.Bool(true))
	}
}

func Call_Main_nativeResult(state_0_loop struct {
	consumed bool
	errors   []struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	}
	stream *Constructor_Data_List_Types_Cons[int64]
	/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
}) *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
	error    string
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
}] {
	var state_0 struct {
		consumed bool
		errors   []struct {
			error    string
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
		}
		stream *Constructor_Data_List_Types_Cons[int64]
		/*gopurs:record-v1:3:0063006f006e00730075006d00650064,006500720072006f00720073,00730074007200650061006d*/
	} = state_0_loop
	_ = state_0
	return (&Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	}]{1, state_0.stream, state_0.errors})
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

func Rebox_Main_1386611502_1510340014(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]struct {
	error    string
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	}])(unsafe.Pointer(in))
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

func Rebox_Main_1386611502_2622893113(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[int64]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[int64]])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_284210452(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
	error    string
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
}]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[*Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	}]])(unsafe.Pointer(in))
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

func Rebox_Main_3162159231_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
	error    string
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
}]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(in.V0))}
	out.V1 = func() gopurs_runtime.Value {
		arr := in.V1
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = func() gopurs_runtime.Value {
				orig := v
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.error), func() gopurs_runtime.Value {
					orig := orig.position
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
				}())
			}()
		}
		return gopurs_runtime.Array(boxed)
	}()
	return out
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

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys3{K0: "consumed", K1: "errors", K2: "stream"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys2{K0: "error", K1: "position"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys2{K0: "column", K1: "line"}
