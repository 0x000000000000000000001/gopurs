package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Vacant gopurs_runtime.Value
var once_Main_Vacant sync.Once

func Get_Main_Vacant() gopurs_runtime.Value {
	once_Main_Vacant.Do(func() {
		cache_Main_Vacant = gopurs_runtime.Value{Type: 9, IntVal: 834710033, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Main_Vacant
}

var cache_Main_Payload gopurs_runtime.Value
var once_Main_Payload sync.Once

func Get_Main_Payload() gopurs_runtime.Value {
	once_Main_Payload.Do(func() {
		cache_Main_Payload = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1603531984, UnsafePtr: unsafe.Pointer((&Constructor_Main_Payload[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Payload
}

var cache_Main_Tagged gopurs_runtime.Value
var once_Main_Tagged sync.Once

func Get_Main_Tagged() gopurs_runtime.Value {
	once_Main_Tagged.Do(func() {
		cache_Main_Tagged = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2502599306, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tagged[gopurs_runtime.Value]{1, value0.StrVal(), value1}))}
			})
		})
	})
	return cache_Main_Tagged
}

var cache_Main_score gopurs_runtime.Value
var once_Main_score sync.Once

func Get_Main_score() gopurs_runtime.Value {
	once_Main_score.Do(func() {
		cache_Main_score = gopurs_runtime.Func(func(packet_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_score(func() struct {
				enabled bool
				left    struct {
					active bool
					column int64
					line   int64
					note   gopurs_runtime.Value
					/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
				}
				payload gopurs_runtime.Value
				right   struct {
					active bool
					label  string
					score  int64
					/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
				}
				/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
			} {
				orig := packet_0_box
				_ = orig
				clone := struct {
					enabled bool
					left    struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					}
					payload gopurs_runtime.Value
					right   struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					}
					/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
				}{}
				clone.enabled = (gopurs_runtime.RecordGet(orig, "enabled").IntVal) != (0)
				clone.left = func() struct {
					active bool
					column int64
					line   int64
					note   gopurs_runtime.Value
					/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "left")
					_ = orig
					clone := struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					}{}
					clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
					clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
					clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
					clone.note = gopurs_runtime.RecordGet(orig, "note")
					return clone
				}()
				clone.payload = gopurs_runtime.RecordGet(orig, "payload")
				clone.right = func() struct {
					active bool
					label  string
					score  int64
					/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
				} {
					orig := gopurs_runtime.RecordGet(orig, "right")
					_ = orig
					clone := struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					}{}
					clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
					clone.score = gopurs_runtime.RecordGet(orig, "score").IntVal
					return clone
				}()
				return clone
			}()))
		})
	})
	return cache_Main_score
}

var cache_Main_rareIdentity gopurs_runtime.Value
var once_Main_rareIdentity sync.Once

func Get_Main_rareIdentity() gopurs_runtime.Value {
	once_Main_rareIdentity.Do(func() {
		cache_Main_rareIdentity = gopurs_runtime.Func(func(packet_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_rareIdentity(func() struct {
					enabled bool
					left    struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					}
					payload gopurs_runtime.Value
					right   struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					}
					/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
				} {
					orig := packet_0_box
					_ = orig
					clone := struct {
						enabled bool
						left    struct {
							active bool
							column int64
							line   int64
							note   gopurs_runtime.Value
							/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
						}
						payload gopurs_runtime.Value
						right   struct {
							active bool
							label  string
							score  int64
							/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
						}
						/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
					}{}
					clone.enabled = (gopurs_runtime.RecordGet(orig, "enabled").IntVal) != (0)
					clone.left = func() struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "left")
						_ = orig
						clone := struct {
							active bool
							column int64
							line   int64
							note   gopurs_runtime.Value
							/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
						}{}
						clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
						clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
						clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
						clone.note = gopurs_runtime.RecordGet(orig, "note")
						return clone
					}()
					clone.payload = gopurs_runtime.RecordGet(orig, "payload")
					clone.right = func() struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "right")
						_ = orig
						clone := struct {
							active bool
							label  string
							score  int64
							/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
						}{}
						clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						clone.score = gopurs_runtime.RecordGet(orig, "score").IntVal
						return clone
					}()
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.enabled), func() gopurs_runtime.Value {
					orig := orig.left
					_ = orig
					return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line), orig.note)
				}(), orig.payload, func() gopurs_runtime.Value {
					orig := orig.right
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.active), gopurs_runtime.Str(orig.label), gopurs_runtime.Int(orig.score))
				}())
			}()
		})
	})
	return cache_Main_rareIdentity
}

var cache_Main_makePacket gopurs_runtime.Value
var once_Main_makePacket sync.Once

func Get_Main_makePacket() gopurs_runtime.Value {
	once_Main_makePacket.Do(func() {
		cache_Main_makePacket = gopurs_runtime.Func(func(seed_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makePacket(seed_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.enabled), func() gopurs_runtime.Value {
					orig := orig.left
					_ = orig
					return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line), orig.note)
				}(), orig.payload, func() gopurs_runtime.Value {
					orig := orig.right
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.active), gopurs_runtime.Str(orig.label), gopurs_runtime.Int(orig.score))
				}())
			}()
		})
	})
	return cache_Main_makePacket
}

var cache_Main_bump gopurs_runtime.Value
var once_Main_bump sync.Once

func Get_Main_bump() gopurs_runtime.Value {
	once_Main_bump.Do(func() {
		cache_Main_bump = gopurs_runtime.Func2(func(amount_0_box gopurs_runtime.Value, packet_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_bump(amount_0_box.IntVal, func() struct {
					enabled bool
					left    struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					}
					payload gopurs_runtime.Value
					right   struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					}
					/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
				} {
					orig := packet_1_box
					_ = orig
					clone := struct {
						enabled bool
						left    struct {
							active bool
							column int64
							line   int64
							note   gopurs_runtime.Value
							/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
						}
						payload gopurs_runtime.Value
						right   struct {
							active bool
							label  string
							score  int64
							/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
						}
						/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
					}{}
					clone.enabled = (gopurs_runtime.RecordGet(orig, "enabled").IntVal) != (0)
					clone.left = func() struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "left")
						_ = orig
						clone := struct {
							active bool
							column int64
							line   int64
							note   gopurs_runtime.Value
							/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
						}{}
						clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
						clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
						clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
						clone.note = gopurs_runtime.RecordGet(orig, "note")
						return clone
					}()
					clone.payload = gopurs_runtime.RecordGet(orig, "payload")
					clone.right = func() struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "right")
						_ = orig
						clone := struct {
							active bool
							label  string
							score  int64
							/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
						}{}
						clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						clone.score = gopurs_runtime.RecordGet(orig, "score").IntVal
						return clone
					}()
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.enabled), func() gopurs_runtime.Value {
					orig := orig.left
					_ = orig
					return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line), orig.note)
				}(), orig.payload, func() gopurs_runtime.Value {
					orig := orig.right
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.active), gopurs_runtime.Str(orig.label), gopurs_runtime.Int(orig.score))
				}())
			}()
		})
	})
	return cache_Main_bump
}

var cache_Main_run gopurs_runtime.Value
var once_Main_run sync.Once

func Get_Main_run() gopurs_runtime.Value {
	once_Main_run.Do(func() {
		cache_Main_run = gopurs_runtime.Func3(func(callback_0_box gopurs_runtime.Value, count_1_box gopurs_runtime.Value, seed_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_run(callback_0_box, count_1_box.IntVal, seed_2_box.IntVal))
		})
	})
	return cache_Main_run
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(5)))
			_ = __local_var_0_0
			seedRef_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = seedRef_1_1
			seed_2_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), seedRef_1_1, gopurs_runtime.Value{})
			_ = seed_2_2
			callbackRef_3_3 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), Get_Main_rareIdentity(), gopurs_runtime.Value{})
			_ = callbackRef_3_3
			callback_4_4 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), callbackRef_3_3, gopurs_runtime.Value{})
			_ = callback_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_run(callback_4_4, int64(3), seed_2_2.IntVal), int64(109)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_run(callback_4_4, int64(0), seed_2_2.IntVal), int64(0)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Vacant[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Payload[T_a any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_Tagged[T_a any] struct {
	Rc uint32
	V0 string
	V1 T_a
}

func Call_Main_score(packet_0_loop struct {
	enabled bool
	left    struct {
		active bool
		column int64
		line   int64
		note   gopurs_runtime.Value
		/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
	}
	payload gopurs_runtime.Value
	right   struct {
		active bool
		label  string
		score  int64
		/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
	}
	/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
}) int64 {
	var packet_0 struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	} = packet_0_loop
	_ = packet_0
	var __t0 int64
	{
		if packet_0.left.active {
			__t0 = int64(1)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = int64(0)
	}
end_branch_0:
	var __t1 int64
	{
		if packet_0.right.active {
			__t1 = int64(2)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = int64(0)
	}
end_branch_1:
	var __t2 int64
	{
		if packet_0.enabled {
			__t2 = int64(4)
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = int64(0)
	}
end_branch_2:
	var __t6 int64
	{
		var __t_tag_3 gopurs_runtime.Value = packet_0.payload
		_ = __t_tag_3
		if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 834710033 {
			__t6 = int64(0)
			goto end_branch_6
		} else {

		}
	}
	{
		var __t_tag_4 gopurs_runtime.Value = packet_0.payload
		_ = __t_tag_4
		if __t_tag_4.Type == 9 && __t_tag_4.IntVal == 1603531984 {
			__t6 = (*Constructor_Main_Payload[gopurs_runtime.Value])(packet_0.payload.UnsafePtr).V0.IntVal
			goto end_branch_6
		} else {

		}
	}
	{
		var __t_tag_5 gopurs_runtime.Value = packet_0.payload
		_ = __t_tag_5
		if __t_tag_5.Type == 9 && __t_tag_5.IntVal == 2502599306 {
			__t6 = (*Constructor_Main_Tagged[gopurs_runtime.Value])(packet_0.payload.UnsafePtr).V1.IntVal
			goto end_branch_6
		} else {

		}
	}
	{
		__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
	}
end_branch_6:
	return gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(packet_0.left.column, packet_0.left.line), packet_0.right.score), __t0), __t1), __t2), __t6)
}

func Call_Main_rareIdentity(packet_0_loop struct {
	enabled bool
	left    struct {
		active bool
		column int64
		line   int64
		note   gopurs_runtime.Value
		/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
	}
	payload gopurs_runtime.Value
	right   struct {
		active bool
		label  string
		score  int64
		/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
	}
	/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
}) struct {
	enabled bool
	left    struct {
		active bool
		column int64
		line   int64
		note   gopurs_runtime.Value
		/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
	}
	payload gopurs_runtime.Value
	right   struct {
		active bool
		label  string
		score  int64
		/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
	}
	/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
} {
	var packet_0 struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	} = packet_0_loop
	_ = packet_0
	return packet_0
}

func Call_Main_makePacket(seed_0_loop int64) struct {
	enabled bool
	left    struct {
		active bool
		column int64
		line   int64
		note   gopurs_runtime.Value
		/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
	}
	payload gopurs_runtime.Value
	right   struct {
		active bool
		label  string
		score  int64
		/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
	}
	/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
} {
	var seed_0 int64 = seed_0_loop
	_ = seed_0
	var __t1 gopurs_runtime.Value
	{
		if (seed_0) < (int64(0)) {
			__t1 = gopurs_runtime.Value{Type: 9, IntVal: 834710033, UnsafePtr: unsafe.Pointer(nil)}
			goto end_branch_1
		} else {

		}
	}
	{
		var __t0 gopurs_runtime.Value
		{
			if (seed_0) == (int64(0)) {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2502599306, UnsafePtr: unsafe.Pointer((&Constructor_Main_Tagged[gopurs_runtime.Value]{1, "zero", gopurs_runtime.Int(int64(0))}))}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1603531984, UnsafePtr: unsafe.Pointer((&Constructor_Main_Payload[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.IntMul(seed_0, int64(2)))}))}
		}
	end_branch_0:
		__t1 = __t0
	}
end_branch_1:
	return struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	}{true, struct {
		active bool
		column int64
		line   int64
		note   gopurs_runtime.Value
		/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
	}{true, seed_0, int64(3), gopurs_runtime.Value{Type: 9, IntVal: 1603531984, UnsafePtr: unsafe.Pointer((&Constructor_Main_Payload[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.IntAdd(seed_0, int64(1)))}))}}, __t1, struct {
		active bool
		label  string
		score  int64
		/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
	}{false, "keep", gopurs_runtime.IntAdd(seed_0, int64(7))}}
}

func Call_Main_bump(amount_0_loop int64, packet_1_loop struct {
	enabled bool
	left    struct {
		active bool
		column int64
		line   int64
		note   gopurs_runtime.Value
		/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
	}
	payload gopurs_runtime.Value
	right   struct {
		active bool
		label  string
		score  int64
		/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
	}
	/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
}) struct {
	enabled bool
	left    struct {
		active bool
		column int64
		line   int64
		note   gopurs_runtime.Value
		/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
	}
	payload gopurs_runtime.Value
	right   struct {
		active bool
		label  string
		score  int64
		/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
	}
	/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
} {
	var amount_0 int64 = amount_0_loop
	_ = amount_0
	var packet_1 struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	} = packet_1_loop
	_ = packet_1
	return func() struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	} {
		clone := packet_1
		clone.left = func() struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		} {
			clone := packet_1.left
			clone.column = gopurs_runtime.IntAdd(packet_1.left.column, amount_0)
			return clone
		}()
		clone.right = func() struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		} {
			clone := packet_1.right
			clone.score = gopurs_runtime.IntAdd(packet_1.right.score, amount_0)
			return clone
		}()
		clone.enabled = (packet_1.enabled) != (true)
		return clone
	}()
}

func Call_Main_run(callback_0_loop gopurs_runtime.Value, count_1_loop int64, seed_2_loop int64) int64 {
	var callback_0 gopurs_runtime.Value = callback_0_loop
	_ = callback_0
	var count_1 int64 = count_1_loop
	_ = count_1
	var seed_2 int64 = seed_2_loop
	_ = seed_2
	var Call_local_Main_go__go_3_0_0 func(int64, struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	}, int64) int64
	_ = Call_local_Main_go__go_3_0_0
	var go__go_3_0_0 gopurs_runtime.Value
	_ = go__go_3_0_0
	Call_local_Main_go__go_3_0_0 = func(n_4_loop int64, packet_5_loop struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	}, total_6_loop int64) int64 {
	go__go_3_0_0:
		for {
			if false {
				continue go__go_3_0_0
			}
			var n_4 int64 = n_4_loop
			_ = n_4
			var packet_5 struct {
				enabled bool
				left    struct {
					active bool
					column int64
					line   int64
					note   gopurs_runtime.Value
					/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
				}
				payload gopurs_runtime.Value
				right   struct {
					active bool
					label  string
					score  int64
					/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
				}
				/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
			} = packet_5_loop
			_ = packet_5
			var total_6 int64 = total_6_loop
			_ = total_6
			var __t2 int64
			{
				if (n_4) <= (int64(0)) {
					__t2 = total_6
					goto end_branch_2
				} else {

				}
			}
			{
				// TAST (Let): next_7_1 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Main","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Main","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
				next_7_1 := Call_Main_bump(int64(1), packet_5)
				_ = next_7_1
				n_4_loop = gopurs_runtime.IntSub(n_4, int64(1))
				packet_5_loop = next_7_1
				total_6_loop = gopurs_runtime.IntAdd(total_6, Call_Main_score(next_7_1))
				continue go__go_3_0_0
				__t2 = func() int64 { panic("unreachable") }()
			}
		end_branch_2:
			return __t2
		}
	}
	go__go_3_0_0 = gopurs_runtime.Func(func(n_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(packet_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(total_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Main_go__go_3_0_0(n_4_loop_val.IntVal, func() struct {
					enabled bool
					left    struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					}
					payload gopurs_runtime.Value
					right   struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					}
					/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
				} {
					orig := packet_5_loop_val
					_ = orig
					clone := struct {
						enabled bool
						left    struct {
							active bool
							column int64
							line   int64
							note   gopurs_runtime.Value
							/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
						}
						payload gopurs_runtime.Value
						right   struct {
							active bool
							label  string
							score  int64
							/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
						}
						/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
					}{}
					clone.enabled = (gopurs_runtime.RecordGet(orig, "enabled").IntVal) != (0)
					clone.left = func() struct {
						active bool
						column int64
						line   int64
						note   gopurs_runtime.Value
						/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "left")
						_ = orig
						clone := struct {
							active bool
							column int64
							line   int64
							note   gopurs_runtime.Value
							/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
						}{}
						clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
						clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
						clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
						clone.note = gopurs_runtime.RecordGet(orig, "note")
						return clone
					}()
					clone.payload = gopurs_runtime.RecordGet(orig, "payload")
					clone.right = func() struct {
						active bool
						label  string
						score  int64
						/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "right")
						_ = orig
						clone := struct {
							active bool
							label  string
							score  int64
							/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
						}{}
						clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						clone.score = gopurs_runtime.RecordGet(orig, "score").IntVal
						return clone
					}()
					return clone
				}(), total_6_loop_val.IntVal))
			})
		})
	})
	__callback_fn_3 := callback_0
	_ = __callback_fn_3
	__callback_arg_4 := Call_Main_makePacket(seed_2)
	_ = __callback_arg_4
	return Call_local_Main_go__go_3_0_0(count_1, func() struct {
		enabled bool
		left    struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		}
		payload gopurs_runtime.Value
		right   struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		}
		/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
	} {
		orig := func(__native_callback func(struct {
			enabled bool
			left    struct {
				active bool
				column int64
				line   int64
				note   gopurs_runtime.Value
				/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
			}
			payload gopurs_runtime.Value
			right   struct {
				active bool
				label  string
				score  int64
				/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
			}
			/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
		}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
			if __native_callback_ok {
				return __native_callback(__callback_arg_4)
			} else {

			}
			return gopurs_runtime.Apply(__callback_fn_3, func() gopurs_runtime.Value {
				orig := __callback_arg_4
				_ = orig
				return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.enabled), func() gopurs_runtime.Value {
					orig := orig.left
					_ = orig
					return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line), orig.note)
				}(), orig.payload, func() gopurs_runtime.Value {
					orig := orig.right
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.active), gopurs_runtime.Str(orig.label), gopurs_runtime.Int(orig.score))
				}())
			}())
		}(gopurs_runtime.FunctionData[func(struct {
			enabled bool
			left    struct {
				active bool
				column int64
				line   int64
				note   gopurs_runtime.Value
				/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
			}
			payload gopurs_runtime.Value
			right   struct {
				active bool
				label  string
				score  int64
				/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
			}
			/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
		}) gopurs_runtime.Value](__callback_fn_3))
		_ = orig
		clone := struct {
			enabled bool
			left    struct {
				active bool
				column int64
				line   int64
				note   gopurs_runtime.Value
				/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
			}
			payload gopurs_runtime.Value
			right   struct {
				active bool
				label  string
				score  int64
				/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
			}
			/*gopurs:record-v1:4:0065006e00610062006c00650064,006c006500660074,007000610079006c006f00610064,00720069006700680074*/
		}{}
		clone.enabled = (gopurs_runtime.RecordGet(orig, "enabled").IntVal) != (0)
		clone.left = func() struct {
			active bool
			column int64
			line   int64
			note   gopurs_runtime.Value
			/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
		} {
			orig := gopurs_runtime.RecordGet(orig, "left")
			_ = orig
			clone := struct {
				active bool
				column int64
				line   int64
				note   gopurs_runtime.Value
				/*gopurs:record-v1:4:006100630074006900760065,0063006f006c0075006d006e,006c0069006e0065,006e006f00740065*/
			}{}
			clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
			clone.column = gopurs_runtime.RecordGet(orig, "column").IntVal
			clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
			clone.note = gopurs_runtime.RecordGet(orig, "note")
			return clone
		}()
		clone.payload = gopurs_runtime.RecordGet(orig, "payload")
		clone.right = func() struct {
			active bool
			label  string
			score  int64
			/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
		} {
			orig := gopurs_runtime.RecordGet(orig, "right")
			_ = orig
			clone := struct {
				active bool
				label  string
				score  int64
				/*gopurs:record-v1:3:006100630074006900760065,006c006100620065006c,00730063006f00720065*/
			}{}
			clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
			clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
			clone.score = gopurs_runtime.RecordGet(orig, "score").IntVal
			return clone
		}()
		return clone
	}(), int64(0))
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

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys4{K0: "enabled", K1: "left", K2: "payload", K3: "right"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys4{K0: "active", K1: "column", K2: "line", K3: "note"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys3{K0: "active", K1: "label", K2: "score"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}
