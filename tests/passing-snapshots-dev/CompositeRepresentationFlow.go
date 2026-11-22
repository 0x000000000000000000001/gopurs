package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Func(func(state_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordUpdate2(state_0, "consumed", gopurs_runtime.Bool(true), "errors", func() gopurs_runtime.Value {
					arr := func() []struct {
						error    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
					} {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_snoc((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(state_0, "errors").UnsafePtr))), func() gopurs_runtime.Value {
							orig := struct {
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
							}{int64(-2147483648), int64(2147483647)}}
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.error), func() gopurs_runtime.Value {
								orig := orig.position
								_ = orig
								return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
							}())
						}())).UnsafePtr)
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
			}))
			_ = __local_var_0_0
			transformRef_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = transformRef_1_1
			transform_2_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), transformRef_1_1, gopurs_runtime.Value{})
			_ = transform_2_2
			originalRef_3_3 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				orig := struct {
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
			}(), gopurs_runtime.Value{})
			_ = originalRef_3_3
			original_4_4 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), originalRef_3_3, gopurs_runtime.Value{})
			_ = original_4_4
			// TAST (Let): result_5_5 shape=App(Var) bindingType=(Record (Row [consumed: Boolean, errors: (Array (Record (Row [error: String, position: (Record (Row [column: Int, line: Int] Empty))] Empty))), stream: (ADT ["Data","List","Types","List"] [Int])] Empty))
			result_5_5 := Call_Main_throughCallback(transform_2_2, func() struct {
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
				orig := original_4_4
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
			_ = result_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Int(int64(len(result_5_5.errors))).IntVal, int64(2)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{result_5_5.consumed, true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			preserved_8_8 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), originalRef_3_3, gopurs_runtime.Value{})
			_ = preserved_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), func() gopurs_runtime.Value {
				orig := struct {
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
					orig := preserved_8_8
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
					orig := original_4_4
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
				}()}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, func() gopurs_runtime.Value {
					orig := orig.actual
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
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
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
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.RecordGet(original_4_4, "consumed").IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())), func() gopurs_runtime.Value {
				orig := struct {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(original_4_4, "errors").UnsafePtr)
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
				}{}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, func() gopurs_runtime.Value {
					arr := orig.actual
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
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
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
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_4()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_5())), func() gopurs_runtime.Value {
				orig := struct {
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
				}{Call_Main_nativeResult(result_5_5), Rebox_Main_138441832_3162159231(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(result_5_5.stream))}, func() gopurs_runtime.Value {
						arr := result_5_5.errors
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
					}()}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3162159231_138441832(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3162159231_138441832(orig.expected))})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_6()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_7())), func() gopurs_runtime.Value {
				orig := struct {
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
				}{}, (*Constructor_Data_List_Types_Cons[int64])(nil)}), Rebox_Main_138441832_3162159231(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array([]gopurs_runtime.Value{})}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3162159231_138441832(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3162159231_138441832(orig.expected))})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			callbackRef_14_14 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Func(func(flag_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordUpdate1(original_4_4, "consumed", gopurs_runtime.Bool((flag_14.IntVal) != (0)))
			}), gopurs_runtime.Value{})
			_ = callbackRef_14_14
			callback_15_15 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), callbackRef_14_14, gopurs_runtime.Value{})
			_ = callback_15_15
			_dollar___unused_16_16 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.RecordGet(gopurs_runtime.Apply(callback_15_15, gopurs_runtime.Bool(true)), "consumed").IntVal) != (0), true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_16_16
			_dollar___unused_17_17 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.RecordGet(gopurs_runtime.Apply(callback_15_15, gopurs_runtime.Bool(false)), "consumed").IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_17_17
			_dollar___unused_18_18 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_8()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_9())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_List_Types_Cons[int64]
					expected *Constructor_Data_List_Types_Cons[int64]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply(callback_15_15, gopurs_runtime.Bool(true)), "stream"))), Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(original_4_4, "stream")))}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(orig.expected))})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_18_18
			_dollar___unused_19_19 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_10()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_11())), func() gopurs_runtime.Value {
				orig := struct {
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
				}{true, "preserved"}}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_4, func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_5, gopurs_runtime.Bool(orig.consumed), gopurs_runtime.Str(orig.extra))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_5, gopurs_runtime.Bool(orig.consumed), gopurs_runtime.Str(orig.extra))
				}())
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_19_19
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_streamIsSymbol(), Call_Data_List_Types_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_errorsIsSymbol(), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_positionIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), gopurs_runtime.Value{}, Get_Main_errorIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))), gopurs_runtime.Value{}, Get_Main_consumedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_consumedIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_errorsIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_streamIsSymbol(), Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), Call_Data_Show_showArray(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_errorIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_positionIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_positionIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), gopurs_runtime.Value{}, Get_Main_errorIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = Call_Data_Show_showArray(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_errorIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_positionIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = Call_Data_Tuple_eqTuple(Call_Data_List_Types_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_positionIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), gopurs_runtime.Value{}, Get_Main_errorIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = Call_Data_Tuple_showTuple(Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), Call_Data_Show_showArray(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_errorIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_positionIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}))))
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = Call_Data_Tuple_eqTuple(Call_Data_List_Types_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), Call_Data_Eq_eqArray(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_positionIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_columnIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))), gopurs_runtime.Value{}, Get_Main_errorIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = Call_Data_Tuple_showTuple(Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), Call_Data_Show_showArray(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_errorIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_positionIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}))))
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = Call_Data_List_Types_eqList(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_extraIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), gopurs_runtime.Value{}, Get_Main_consumedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_consumedIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_extraIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))
	})
	return cache_Main___cached_dict_11
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
		orig := gopurs_runtime.Apply(transform_0, gopurs_runtime.Apply(transform_0, func() gopurs_runtime.Value {
			orig := original_1
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
		}()))
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
	return Rebox_Main_138441832_3162159231(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(state_0.stream))}, func() gopurs_runtime.Value {
			arr := state_0.errors
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
		}()}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))
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

func Rebox_Main_138441832_3162159231(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
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
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[int64], []struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	}]{}
	out.V0 = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0))
	out.V1 = func() []struct {
		error    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		/*gopurs:record-v1:2:006500720072006f0072,0070006f0073006900740069006f006e*/
	} {
		arr := *(*[]gopurs_runtime.Value)(in.V1.UnsafePtr)
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
	return out
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

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
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

var GopursRecordKeys_4d61696e_4 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}

var GopursRecordKeys_4d61696e_5 = gopurs_runtime.RecordKeys2{K0: "consumed", K1: "extra"}
