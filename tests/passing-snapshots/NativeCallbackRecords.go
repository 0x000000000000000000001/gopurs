package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Main_add
}

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = Get_Control_Bind_bind__3603852628()
	})
	return cache_Main_bind
}

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = Get_Control_Bind_bind__2080121076()
	})
	return cache_Main_bind1
}

var cache_Main_greaterThan gopurs_runtime.Value
var once_Main_greaterThan sync.Once

func Get_Main_greaterThan() gopurs_runtime.Value {
	once_Main_greaterThan.Do(func() {
		cache_Main_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Main_greaterThan
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__2052799636()
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
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__374790306()
	})
	return cache_Main_assertEqual
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__4104969108()
	})
	return cache_Main_bind3
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__1234338868()
	})
	return cache_Main_bind4
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

var cache_Main_valueIsSymbol gopurs_runtime.Value
var once_Main_valueIsSymbol sync.Once

func Get_Main_valueIsSymbol() gopurs_runtime.Value {
	once_Main_valueIsSymbol.Do(func() {
		cache_Main_valueIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("value", 5)
		}))
	})
	return cache_Main_valueIsSymbol
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

var cache_Main_eqRowCons1 gopurs_runtime.Value
var once_Main_eqRowCons1 sync.Once

func Get_Main_eqRowCons1() gopurs_runtime.Value {
	once_Main_eqRowCons1.Do(func() {
		cache_Main_eqRowCons1 = Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_positionIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))
	})
	return cache_Main_eqRowCons1
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
		cache_Main_showRecordFieldsCons = Call_Data_Show_showRecordFieldsCons(Get_Main_positionIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))
	})
	return cache_Main_showRecordFieldsCons
}

var cache_Main_assertEqual1 gopurs_runtime.Value
var once_Main_assertEqual1 sync.Once

func Get_Main_assertEqual1() gopurs_runtime.Value {
	once_Main_assertEqual1.Do(func() {
		cache_Main_assertEqual1 = Get_Test_Assert_assertEqual__640907996()
	})
	return cache_Main_assertEqual1
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Bind_bind__1912541172()
	})
	return cache_Main_bind5
}

var cache_Main_bind6 gopurs_runtime.Value
var once_Main_bind6 sync.Once

func Get_Main_bind6() gopurs_runtime.Value {
	once_Main_bind6.Do(func() {
		cache_Main_bind6 = Get_Control_Bind_bind__533498260()
	})
	return cache_Main_bind6
}

var cache_Main_bind7 gopurs_runtime.Value
var once_Main_bind7 sync.Once

func Get_Main_bind7() gopurs_runtime.Value {
	once_Main_bind7.Do(func() {
		cache_Main_bind7 = Get_Control_Bind_bind__1573548948()
	})
	return cache_Main_bind7
}

var cache_Main_bind8 gopurs_runtime.Value
var once_Main_bind8 sync.Once

func Get_Main_bind8() gopurs_runtime.Value {
	once_Main_bind8.Do(func() {
		cache_Main_bind8 = Get_Control_Bind_bind__88048852()
	})
	return cache_Main_bind8
}

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = Get_Test_Assert_assertEqual__2127306453()
	})
	return cache_Main_assertEqual2
}

var cache_Main_bind9 gopurs_runtime.Value
var once_Main_bind9 sync.Once

func Get_Main_bind9() gopurs_runtime.Value {
	once_Main_bind9.Do(func() {
		cache_Main_bind9 = Get_Control_Bind_bind__2730670804()
	})
	return cache_Main_bind9
}

var cache_Main_bind10 gopurs_runtime.Value
var once_Main_bind10 sync.Once

func Get_Main_bind10() gopurs_runtime.Value {
	once_Main_bind10.Do(func() {
		cache_Main_bind10 = Get_Control_Bind_bind__711607892()
	})
	return cache_Main_bind10
}

var cache_Main_assertEqual3 gopurs_runtime.Value
var once_Main_assertEqual3 sync.Once

func Get_Main_assertEqual3() gopurs_runtime.Value {
	once_Main_assertEqual3.Do(func() {
		cache_Main_assertEqual3 = Get_Test_Assert_assertEqual__2867988789()
	})
	return cache_Main_assertEqual3
}

var cache_Main_bind11 gopurs_runtime.Value
var once_Main_bind11 sync.Once

func Get_Main_bind11() gopurs_runtime.Value {
	once_Main_bind11.Do(func() {
		cache_Main_bind11 = Get_Control_Bind_bind__1623554388()
	})
	return cache_Main_bind11
}

var cache_Main_bind12 gopurs_runtime.Value
var once_Main_bind12 sync.Once

func Get_Main_bind12() gopurs_runtime.Value {
	once_Main_bind12.Do(func() {
		cache_Main_bind12 = Get_Control_Bind_bind__4020646868()
	})
	return cache_Main_bind12
}

var cache_Main_assertEqual4 gopurs_runtime.Value
var once_Main_assertEqual4 sync.Once

func Get_Main_assertEqual4() gopurs_runtime.Value {
	once_Main_assertEqual4.Do(func() {
		cache_Main_assertEqual4 = Get_Test_Assert_assertEqual__2832848779()
	})
	return cache_Main_assertEqual4
}

var cache_Main_not gopurs_runtime.Value
var once_Main_not sync.Once

func Get_Main_not() gopurs_runtime.Value {
	once_Main_not.Do(func() {
		cache_Main_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Main_not
}

var cache_Main_namedCallback gopurs_runtime.Value
var once_Main_namedCallback sync.Once

func Get_Main_namedCallback() gopurs_runtime.Value {
	once_Main_namedCallback.Do(func() {
		cache_Main_namedCallback = gopurs_runtime.Func(func(input_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_namedCallback(func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_0_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Main_namedCallback
}

var cache_Main_consume gopurs_runtime.Value
var once_Main_consume sync.Once

func Get_Main_consume() gopurs_runtime.Value {
	once_Main_consume.Do(func() {
		cache_Main_consume = gopurs_runtime.Func2(func(callback_0_box gopurs_runtime.Value, input_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_consume(callback_0_box, func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_1_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Main_consume
}

var cache_Main_consume__3961633801 gopurs_runtime.Value
var once_Main_consume__3961633801 sync.Once

func Get_Main_consume__3961633801() gopurs_runtime.Value {
	once_Main_consume__3961633801.Do(func() {
		cache_Main_consume__3961633801 = gopurs_runtime.Func2(func(callback_0_box gopurs_runtime.Value, input_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_consume__3961633801(callback_0_box, func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_1_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Main_consume__3961633801
}

var cache_Main_consume__186119419 gopurs_runtime.Value
var once_Main_consume__186119419 sync.Once

func Get_Main_consume__186119419() gopurs_runtime.Value {
	once_Main_consume__186119419.Do(func() {
		cache_Main_consume__186119419 = gopurs_runtime.Func2(func(callback_0_box gopurs_runtime.Value, input_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_consume__186119419(callback_0_box, func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_1_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Main_consume__186119419
}

var cache_Main_consume__775892937 gopurs_runtime.Value
var once_Main_consume__775892937 sync.Once

func Get_Main_consume__775892937() gopurs_runtime.Value {
	once_Main_consume__775892937.Do(func() {
		cache_Main_consume__775892937 = gopurs_runtime.Func2(func(callback_0_box gopurs_runtime.Value, input_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_consume__775892937(callback_0_box, func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_1_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Main_consume__775892937
}

var cache_Main_consume__3591009769 gopurs_runtime.Value
var once_Main_consume__3591009769 sync.Once

func Get_Main_consume__3591009769() gopurs_runtime.Value {
	once_Main_consume__3591009769.Do(func() {
		cache_Main_consume__3591009769 = gopurs_runtime.Func2(func(callback_0_box gopurs_runtime.Value, input_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_consume__3591009769(callback_0_box, func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_1_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
							orig := orig.position
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
						}(), gopurs_runtime.Int(orig.value))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Main_consume__3591009769
}

var cache_Main_choose gopurs_runtime.Value
var once_Main_choose sync.Once

func Get_Main_choose() gopurs_runtime.Value {
	once_Main_choose.Do(func() {
		cache_Main_choose = gopurs_runtime.Func3(func(flag_0_box gopurs_runtime.Value, left_1_box gopurs_runtime.Value, right_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_choose((flag_0_box.IntVal) != (0), func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := left_1_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}(), func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := right_2_box
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
					orig := orig.position
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
				}(), gopurs_runtime.Int(orig.value))
			}()
		})
	})
	return cache_Main_choose
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__callback_fn_1 := Get_Effect_Ref__new()
			_ = __callback_fn_1
			__callback_arg_2 := struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}{struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}{int64(2), int64(3)}, int64(7)}
			_ = __callback_arg_2
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := func(__native_callback func(struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
				if __native_callback_ok {
					return __native_callback(__callback_arg_2)
				} else {

				}
				return gopurs_runtime.Apply(__callback_fn_1, func() gopurs_runtime.Value {
					orig := __callback_arg_2
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), gopurs_runtime.Int(orig.value))
				}())
			}(gopurs_runtime.FunctionData[func(struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value](__callback_fn_1))
			_ = __local_var_0_0
			__local_var_1_3 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_3
			__local_var_2_4 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_1_3, gopurs_runtime.Value{})
			_ = __local_var_2_4
			__local_var_3_5 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(token_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t6 gopurs_runtime.Value
				{
					if (gopurs_runtime.RecordGet(token_3, "value").IntVal) > (int64(0)) {
						__t6 = func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(token_3, "value").IntVal, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(token_3, "position"), "column").IntVal)), true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}
			end_branch_6:
				return __t6
			}), func(token_3 struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value {
				var __t7 gopurs_runtime.Value
				{
					if (token_3.value) > (int64(0)) {
						__t7 = func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Int(gopurs_runtime.IntAdd(token_3.value, token_3.position.column)), true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()
						goto end_branch_7
					} else {

					}
				}
				{
					__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}
			end_branch_7:
				return __t7
			}), gopurs_runtime.Value{})
			_ = __local_var_3_5
			__local_var_4_8 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_3_5, gopurs_runtime.Value{})
			_ = __local_var_4_8
			__callback_fn_10 := __local_var_4_8
			_ = __callback_fn_10
			__callback_arg_11 := __local_var_2_4
			_ = __callback_arg_11
			__local_var_5_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1255103439(Rebox_Main_1386611502_2818770644(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_0())), "", struct {
				actual   *Constructor_Data_Maybe_Just[int64]
				expected *Constructor_Data_Maybe_Just[int64]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__callback_fn_10, __callback_arg_11))), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(int64(9)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}), gopurs_runtime.Value{})
			_ = __local_var_5_9
			__callback_fn_13 := __local_var_4_8
			_ = __callback_fn_13
			__callback_arg_14 := gopurs_runtime.RecordUpdate1(__local_var_2_4, "value", gopurs_runtime.Int(int64(0)))
			_ = __callback_arg_14
			__local_var_6_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1255103439(Rebox_Main_1386611502_2818770644(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), "", struct {
				actual   *Constructor_Data_Maybe_Just[int64]
				expected *Constructor_Data_Maybe_Just[int64]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__callback_fn_13, __callback_arg_14))), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}), gopurs_runtime.Value{})
			_ = __local_var_6_12
			__local_var_7_15 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(token_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{token_7, true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
			}), func(token_7 struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{func() gopurs_runtime.Value {
						orig := token_7
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
							orig := orig.position
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
						}(), gopurs_runtime.Int(orig.value))
					}(), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
			}), gopurs_runtime.Value{})
			_ = __local_var_7_15
			__local_var_8_16 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_7_15, gopurs_runtime.Value{})
			_ = __local_var_8_16
			__callback_fn_18 := __local_var_8_16
			_ = __callback_fn_18
			__callback_arg_19 := __local_var_2_4
			_ = __callback_arg_19
			__local_var_9_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___4119699256(Rebox_Main_1386611502_2844336510(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_2())), "", struct {
				actual *Constructor_Data_Maybe_Just[struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}]
				expected *Constructor_Data_Maybe_Just[struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Rebox_Main_3094389156_820964629(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__callback_fn_18, __callback_arg_19))), Rebox_Main_3094389156_820964629(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}
					V1 bool
				}{func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := __local_var_2_4
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: func() gopurs_runtime.Value {
						orig := _v.V0
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
							orig := orig.position
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
						}(), gopurs_runtime.Int(orig.value))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}), gopurs_runtime.Value{})
			_ = __local_var_9_17
			__local_var_10_20 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), Get_Main_namedCallback(), gopurs_runtime.Value{})
			_ = __local_var_10_20
			__local_var_11_21 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_10_20, gopurs_runtime.Value{})
			_ = __local_var_11_21
			__local_var_12_22 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(input_12 gopurs_runtime.Value) gopurs_runtime.Value {
				__callback_fn_23 := __local_var_11_21
				_ = __callback_fn_23
				__callback_arg_24 := input_12
				_ = __callback_arg_24
				return gopurs_runtime.Apply(__callback_fn_23, __callback_arg_24)
			}), func(input_12 struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value {
				__callback_fn_25 := __local_var_11_21
				_ = __callback_fn_25
				__callback_arg_26 := input_12
				_ = __callback_arg_26
				return func(__native_callback func(struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
					if __native_callback_ok {
						return __native_callback(__callback_arg_26)
					} else {

					}
					return gopurs_runtime.Apply(__callback_fn_25, func() gopurs_runtime.Value {
						orig := __callback_arg_26
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
							orig := orig.position
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
						}(), gopurs_runtime.Int(orig.value))
					}())
				}(gopurs_runtime.FunctionData[func(struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}) gopurs_runtime.Value](__callback_fn_25))
			}), gopurs_runtime.Value{})
			_ = __local_var_12_22
			__local_var_13_27 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_12_22, gopurs_runtime.Value{})
			_ = __local_var_13_27
			__callback_fn_29 := __local_var_13_27
			_ = __callback_fn_29
			__callback_arg_30 := __local_var_2_4
			_ = __callback_arg_30
			__local_var_14_28 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1255103439(Rebox_Main_1386611502_2818770644(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())), "", struct {
				actual   *Constructor_Data_Maybe_Just[int64]
				expected *Constructor_Data_Maybe_Just[int64]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__callback_fn_29, __callback_arg_30))), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Int(int64(10)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}), gopurs_runtime.Value{})
			_ = __local_var_14_28
			__local_var_15_31 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(token_15 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Func(func(offset_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(token_15, "value").IntVal, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(token_15, "position"), "line").IntVal), offset_16.IntVal))
					}), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
			}), func(token_15 struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Func(func(offset_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(token_15.value, token_15.position.line), offset_16.IntVal))
					}), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
			}), gopurs_runtime.Value{})
			_ = __local_var_15_31
			__local_var_16_32 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_15_31, gopurs_runtime.Value{})
			_ = __local_var_16_32
			__callback_fn_34 := __local_var_16_32
			_ = __callback_fn_34
			__callback_arg_35 := __local_var_2_4
			_ = __callback_arg_35
			// TAST (Let): v_17_33 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] Int)])
			v_17_33 := gopurs_runtime.Apply(__callback_fn_34, __callback_arg_35)
			_ = v_17_33
			var __t40 gopurs_runtime.Value
			{
				if v_17_33.Type == 9 && v_17_33.IntVal == 930809136 && v_17_33.UnsafePtr == nil {
					__t40 = Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
						actual   bool
						expected bool
						/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
					}{false, true})
					goto end_branch_40
				} else {

				}
			}
			{
				if v_17_33.Type == 9 && v_17_33.IntVal == 930809136 && v_17_33.UnsafePtr != nil {
					__t40 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_18_37 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_17_33.UnsafePtr).V0, gopurs_runtime.Value{})
						_ = __local_var_18_37
						__local_var_19_38 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_18_37, gopurs_runtime.Value{})
						_ = __local_var_19_38
						__local_var_20_39 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
							actual   int64
							expected int64
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{gopurs_runtime.Apply(__local_var_19_38, gopurs_runtime.Int(int64(4))).IntVal, int64(14)}), gopurs_runtime.Value{})
						_ = __local_var_20_39
						return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
							actual   int64
							expected int64
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{gopurs_runtime.Apply(__local_var_19_38, gopurs_runtime.Int(int64(5))).IntVal, int64(15)}), gopurs_runtime.Value{})
					})
					goto end_branch_40
				} else {

				}
			}
			{
				__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_40:
			__local_var_18_36 := gopurs_runtime.Apply(__t40, gopurs_runtime.Value{})
			_ = __local_var_18_36
			__local_var_19_41 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Bool(true), gopurs_runtime.Value{})
			_ = __local_var_19_41
			__local_var_20_42 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_19_41, gopurs_runtime.Value{})
			_ = __local_var_20_42
			// TAST (Let): other_21_43 shape=Other bindingType=(Record (Row [value: Int, position: (Record (Row [line: Int, column: Int] Empty))] Empty))
			other_21_43 := gopurs_runtime.RecordUpdate1(__local_var_2_4, "value", gopurs_runtime.Int(int64(12)))
			_ = other_21_43
			__local_var_22_44 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3781821398(Rebox_Main_1386611502_2250985407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_4())), "", struct {
				actual struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}
				expected struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_choose((__local_var_20_42.IntVal) != (0), func() struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			} {
				orig := __local_var_2_4
				_ = orig
				clone := struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}{}
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
				clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
				return clone
			}(), func() struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			} {
				orig := other_21_43
				_ = orig
				clone := struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}{}
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
				clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
				return clone
			}()), func() struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			} {
				orig := __local_var_2_4
				_ = orig
				clone := struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}{}
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
				clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
				return clone
			}()}), gopurs_runtime.Value{})
			_ = __local_var_22_44
			__local_var_23_45 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3781821398(Rebox_Main_1386611502_2250985407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_5())), "", struct {
				actual struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}
				expected struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_choose(((__local_var_20_42.IntVal) != (0)) != (true), func() struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			} {
				orig := __local_var_2_4
				_ = orig
				clone := struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}{}
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
				clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
				return clone
			}(), func() struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			} {
				orig := other_21_43
				_ = orig
				clone := struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}{}
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
				clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
				return clone
			}()), func() struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			} {
				orig := other_21_43
				_ = orig
				clone := struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}{}
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
				clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
				return clone
			}()}), gopurs_runtime.Value{})
			_ = __local_var_23_45
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = Call_Data_Maybe_showMaybe(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Get_Main_showRecordFieldsCons()))
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Get_Main_showRecordFieldsCons())
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Get_Main_showRecordFieldsCons())
	})
	return cache_Main___cached_dict_5
}

func Call_Main_namedCallback(input_0_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var input_0 struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	} = input_0_loop
	_ = input_0
	return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{gopurs_runtime.Int(gopurs_runtime.IntAdd(input_0.value, input_0.position.line)), true}
}

func Call_Main_consume(callback_0_loop gopurs_runtime.Value, input_1_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var callback_0 gopurs_runtime.Value = callback_0_loop
	_ = callback_0
	var input_1 struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	} = input_1_loop
	_ = input_1
	__callback_fn_0 := callback_0
	_ = __callback_fn_0
	__callback_arg_1 := input_1
	_ = __callback_arg_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := func(__native_callback func(struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
		if __native_callback_ok {
			return __native_callback(__callback_arg_1)
		} else {

		}
		return gopurs_runtime.Apply(__callback_fn_0, func() gopurs_runtime.Value {
			orig := __callback_arg_1
			_ = orig
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
				orig := orig.position
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
			}(), gopurs_runtime.Int(orig.value))
		}())
	}(gopurs_runtime.FunctionData[func(struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	}) gopurs_runtime.Value](__callback_fn_0)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Main_consume__3961633801(callback_0_loop gopurs_runtime.Value, input_1_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
consume__3961633801:
	for {
		if false {
			continue consume__3961633801
		}
		var callback_0 gopurs_runtime.Value = callback_0_loop
		_ = callback_0
		var input_1 struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		} = input_1_loop
		_ = input_1
		__callback_fn_0 := callback_0
		_ = __callback_fn_0
		__callback_arg_1 := input_1
		_ = __callback_arg_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := func(__native_callback func(struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
			if __native_callback_ok {
				return __native_callback(__callback_arg_1)
			} else {

			}
			return gopurs_runtime.Apply(__callback_fn_0, func() gopurs_runtime.Value {
				orig := __callback_arg_1
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
					orig := orig.position
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
				}(), gopurs_runtime.Int(orig.value))
			}())
		}(gopurs_runtime.FunctionData[func(struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		}) gopurs_runtime.Value](__callback_fn_0)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Main_consume__186119419(callback_0_loop gopurs_runtime.Value, input_1_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
consume__186119419:
	for {
		if false {
			continue consume__186119419
		}
		var callback_0 gopurs_runtime.Value = callback_0_loop
		_ = callback_0
		var input_1 struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		} = input_1_loop
		_ = input_1
		__callback_fn_0 := callback_0
		_ = __callback_fn_0
		__callback_arg_1 := input_1
		_ = __callback_arg_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := func(__native_callback func(struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
			if __native_callback_ok {
				return __native_callback(__callback_arg_1)
			} else {

			}
			return gopurs_runtime.Apply(__callback_fn_0, func() gopurs_runtime.Value {
				orig := __callback_arg_1
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
					orig := orig.position
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
				}(), gopurs_runtime.Int(orig.value))
			}())
		}(gopurs_runtime.FunctionData[func(struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		}) gopurs_runtime.Value](__callback_fn_0)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Main_consume__775892937(callback_0_loop gopurs_runtime.Value, input_1_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
consume__775892937:
	for {
		if false {
			continue consume__775892937
		}
		var callback_0 gopurs_runtime.Value = callback_0_loop
		_ = callback_0
		var input_1 struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		} = input_1_loop
		_ = input_1
		__callback_fn_0 := callback_0
		_ = __callback_fn_0
		__callback_arg_1 := input_1
		_ = __callback_arg_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := func(__native_callback func(struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
			if __native_callback_ok {
				return __native_callback(__callback_arg_1)
			} else {

			}
			return gopurs_runtime.Apply(__callback_fn_0, func() gopurs_runtime.Value {
				orig := __callback_arg_1
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
					orig := orig.position
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
				}(), gopurs_runtime.Int(orig.value))
			}())
		}(gopurs_runtime.FunctionData[func(struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		}) gopurs_runtime.Value](__callback_fn_0)); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Main_consume__3591009769(callback_0_loop gopurs_runtime.Value, input_1_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}) struct {
	V0 struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	}
	V1 bool
} {
consume__3591009769:
	for {
		if false {
			continue consume__3591009769
		}
		var callback_0 gopurs_runtime.Value = callback_0_loop
		_ = callback_0
		var input_1 struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		} = input_1_loop
		_ = input_1
		__callback_fn_0 := callback_0
		_ = __callback_fn_0
		__callback_arg_1 := input_1
		_ = __callback_arg_1
		return func() struct {
			V0 struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}
			V1 bool
		} {
			_v := func(__native_callback func(struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
				if __native_callback_ok {
					return __native_callback(__callback_arg_1)
				} else {

				}
				return gopurs_runtime.Apply(__callback_fn_0, func() gopurs_runtime.Value {
					orig := __callback_arg_1
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), gopurs_runtime.Int(orig.value))
				}())
			}(gopurs_runtime.FunctionData[func(struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}) gopurs_runtime.Value](__callback_fn_0))
			if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
				return struct {
					V0 struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}
					V1 bool
				}{V0: func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0
					_ = orig
					clone := struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					}{}
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
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}(), V1: true}
			}
			return struct {
				V0 struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				}
				V1 bool
			}{V0: struct {
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				value int64
				/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
			}{}, V1: false}
		}()
	}
}

func Call_Main_choose(flag_0_loop bool, left_1_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}, right_2_loop struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}) struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
} {
	var flag_0 bool = flag_0_loop
	_ = flag_0
	var left_1 struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	} = left_1_loop
	_ = left_1
	var right_2 struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	} = right_2_loop
	_ = right_2
	var __t0 struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	}
	{
		if flag_0 {
			__t0 = left_1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = right_2
	}
end_branch_0:
	return __t0
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2250985407(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2818770644(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2844336510(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	}]])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Main_3094389156_820964629(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	}]{}
	out.V0 = func() struct {
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		value int64
		/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			position struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}
			value int64
			/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
		}{}
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
		clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
		return clone
	}()
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys1{K0: "reflectSymbol"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys2{K0: "position", K1: "value"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys2{K0: "column", K1: "line"}
