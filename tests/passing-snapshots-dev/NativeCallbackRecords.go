package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			__callback_fn_1 := Get_Effect_Ref_go__new()
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
			inputRef_1_3 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = inputRef_1_3
			input_2_4 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), inputRef_1_3, gopurs_runtime.Value{})
			_ = input_2_4
			callbackRef_3_5 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(token_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
			_ = callbackRef_3_5
			callback_4_8 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), callbackRef_3_5, gopurs_runtime.Value{})
			_ = callback_4_8
			_dollar___unused_5_9 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_Maybe_Just[int64]
					expected *Constructor_Data_Maybe_Just[int64]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := Call_Main_consume(callback_4_8, func() struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					} {
						orig := input_2_4
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
				}())), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(int64(9)), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.expected))})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_5_9
			_dollar___unused_6_10 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_Maybe_Just[int64]
					expected *Constructor_Data_Maybe_Just[int64]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := Call_Main_consume(callback_4_8, func() struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					} {
						orig := gopurs_runtime.RecordUpdate1(input_2_4, "value", gopurs_runtime.Int(int64(0)))
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
				}())), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.expected))})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_6_10
			identityRef_7_11 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(token_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
			_ = identityRef_7_11
			identity_8_12 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), identityRef_7_11, gopurs_runtime.Value{})
			_ = identity_8_12
			_dollar___unused_9_13 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_4()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_5())), func() gopurs_runtime.Value {
				orig := struct {
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
				}{Rebox_Main_3094389156_820964629(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := Call_Main_consume(identity_8_12, func() struct {
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						value int64
						/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
					} {
						orig := input_2_4
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
				}())), Rebox_Main_3094389156_820964629(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
						orig := input_2_4
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
				}()))}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_820964629_3094389156(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_820964629_3094389156(orig.expected))})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_9_13
			namedRef_10_14 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), Get_Main_namedCallback(), gopurs_runtime.Value{})
			_ = namedRef_10_14
			named_11_15 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), namedRef_10_14, gopurs_runtime.Value{})
			_ = named_11_15
			partialRef_12_16 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Main_consume(), named_11_15), gopurs_runtime.Value{})
			_ = partialRef_12_16
			partial_13_17 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), partialRef_12_16, gopurs_runtime.Value{})
			_ = partial_13_17
			__callback_fn_19 := partial_13_17
			_ = __callback_fn_19
			__callback_arg_20 := input_2_4
			_ = __callback_arg_20
			_dollar___unused_14_18 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_6()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_7())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_Maybe_Just[int64]
					expected *Constructor_Data_Maybe_Just[int64]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__callback_fn_19, __callback_arg_20))), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(int64(10)), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.expected))})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_14_18
			factoryRef_15_21 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(token_15 gopurs_runtime.Value) gopurs_runtime.Value {
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
			_ = factoryRef_15_21
			factory_16_22 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), factoryRef_15_21, gopurs_runtime.Value{})
			_ = factory_16_22
			// TAST (Let): v_17_23 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] Int)])
			v_17_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Main_consume(factory_16_22, func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_2_4
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
			}())
			_ = v_17_23
			var __t28 gopurs_runtime.Value
			{
				if v_17_23 == nil {
					__t28 = gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
						orig := struct {
							actual   bool
							expected bool
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{false, true}
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
					}())
					goto end_branch_28
				} else {

				}
			}
			{
				if v_17_23 != nil {
					__t28 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						closureRef_18_25 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), (v_17_23).V0, gopurs_runtime.Value{})
						_ = closureRef_18_25
						saved_19_26 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), closureRef_18_25, gopurs_runtime.Value{})
						_ = saved_19_26
						_dollar___unused_20_27 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
							orig := struct {
								actual   int64
								expected int64
								/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
							}{gopurs_runtime.Apply(saved_19_26, gopurs_runtime.Int(int64(4))).IntVal, int64(14)}
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
						}(), gopurs_runtime.Value{})
						_ = _dollar___unused_20_27
						return gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
							orig := struct {
								actual   int64
								expected int64
								/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
							}{gopurs_runtime.Apply(saved_19_26, gopurs_runtime.Int(int64(5))).IntVal, int64(15)}
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
						}(), gopurs_runtime.Value{})
					})
					goto end_branch_28
				} else {

				}
			}
			{
				__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_28:
			_dollar___unused_18_24 := gopurs_runtime.Apply(__t28, gopurs_runtime.Value{})
			_ = _dollar___unused_18_24
			flagRef_19_29 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Bool(true), gopurs_runtime.Value{})
			_ = flagRef_19_29
			flag_20_30 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), flagRef_19_29, gopurs_runtime.Value{})
			_ = flag_20_30
			// TAST (Let): other_21_31 shape=Other bindingType=(Record (Row [value: Int, position: (Record (Row [line: Int, column: Int] Empty))] Empty))
			other_21_31 := gopurs_runtime.RecordUpdate1(input_2_4, "value", gopurs_runtime.Int(int64(12)))
			_ = other_21_31
			_dollar___unused_22_32 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_8()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_9())), func() gopurs_runtime.Value {
				orig := struct {
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
				}{Call_Main_choose((flag_20_30.IntVal) != (0), func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_2_4
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
					orig := other_21_31
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
					orig := input_2_4
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
				}()}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), gopurs_runtime.Int(orig.value))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), gopurs_runtime.Int(orig.value))
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_22_32
			_dollar___unused_23_33 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_10()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_11())), func() gopurs_runtime.Value {
				orig := struct {
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
				}{Call_Main_choose(((flag_20_30.IntVal) != (0)) != (true), func() struct {
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					value int64
					/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
				} {
					orig := input_2_4
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
					orig := other_21_31
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
					orig := other_21_31
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
				}()}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), gopurs_runtime.Int(orig.value))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), gopurs_runtime.Int(orig.value))
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_23_33
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
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
		cache_Main___cached_dict_2 = Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
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
		cache_Main___cached_dict_4 = Call_Data_Maybe_eqMaybe(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_positionIsSymbol()), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = Call_Data_Maybe_showMaybe(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_positionIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_positionIsSymbol()), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_positionIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_positionIsSymbol()), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_positionIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))
	})
	return cache_Main___cached_dict_11
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

func Rebox_Main_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
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

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_820964629_3094389156(in *Constructor_Data_Maybe_Just[struct {
	position struct {
		column int64
		line   int64
		/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
	}
	value int64
	/*gopurs:record-v1:2:0070006f0073006900740069006f006e,00760061006c00750065*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
			orig := orig.position
			_ = orig
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
		}(), gopurs_runtime.Int(orig.value))
	}()
	return out
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys1{K0: "reflectSymbol"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys2{K0: "position", K1: "value"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys2{K0: "column", K1: "line"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}
