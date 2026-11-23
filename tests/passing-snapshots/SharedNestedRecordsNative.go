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

var cache_Main_lessThan gopurs_runtime.Value
var once_Main_lessThan sync.Once

func Get_Main_lessThan() gopurs_runtime.Value {
	once_Main_lessThan.Do(func() {
		cache_Main_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Main_lessThan
}

var cache_Main_eq gopurs_runtime.Value
var once_Main_eq sync.Once

func Get_Main_eq() gopurs_runtime.Value {
	once_Main_eq.Do(func() {
		cache_Main_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Main_eq
}

var cache_Main_mul gopurs_runtime.Value
var once_Main_mul sync.Once

func Get_Main_mul() gopurs_runtime.Value {
	once_Main_mul.Do(func() {
		cache_Main_mul = Get_Data_Semiring_mul__4086903470()
	})
	return cache_Main_mul
}

var cache_Main_not gopurs_runtime.Value
var once_Main_not sync.Once

func Get_Main_not() gopurs_runtime.Value {
	once_Main_not.Do(func() {
		cache_Main_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Main_not
}

var cache_Main_lessThanOrEq gopurs_runtime.Value
var once_Main_lessThanOrEq sync.Once

func Get_Main_lessThanOrEq() gopurs_runtime.Value {
	once_Main_lessThanOrEq.Do(func() {
		cache_Main_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Main_lessThanOrEq
}

var cache_Main_sub gopurs_runtime.Value
var once_Main_sub sync.Once

func Get_Main_sub() gopurs_runtime.Value {
	once_Main_sub.Do(func() {
		cache_Main_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Main_sub
}

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = Get_Control_Bind_bind__2595690612()
	})
	return cache_Main_bind
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

var cache_Main_Payload__2978473649 gopurs_runtime.Value
var once_Main_Payload__2978473649 sync.Once

func Get_Main_Payload__2978473649() gopurs_runtime.Value {
	once_Main_Payload__2978473649.Do(func() {
		cache_Main_Payload__2978473649 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Payload__2978473649(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_Payload__2978473649
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

var cache_Main_Tagged__46254249 gopurs_runtime.Value
var once_Main_Tagged__46254249 sync.Once

func Get_Main_Tagged__46254249() gopurs_runtime.Value {
	once_Main_Tagged__46254249.Do(func() {
		cache_Main_Tagged__46254249 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Tagged__46254249(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box.IntVal)
		})
	})
	return cache_Main_Tagged__46254249
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
		cache_Main_run = gopurs_runtime.Func2(func(count_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_run(count_0_box.IntVal, seed_1_box.IntVal))
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
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(5)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_1_1, gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_score(Call_Main_makePacket(__local_var_2_2.IntVal)), int64(35)}), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_score(Call_Main_makePacket(int64(0))), int64(15)}), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_run(int64(3), __local_var_2_2.IntVal), int64(109)}), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_run(int64(0), __local_var_2_2.IntVal), int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_6_6
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

func Call_Main_Payload__2978473649(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
Payload__2978473649:
	for {
		if false {
			continue Payload__2978473649
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1603531984, UnsafePtr: unsafe.Pointer(Rebox_Main_1938252319_3483727428((&Constructor_Main_Payload[int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Tagged__46254249(__eta_norm_1_0_loop string, __eta_norm_0_1_loop int64) gopurs_runtime.Value {
Tagged__46254249:
	for {
		if false {
			continue Tagged__46254249
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 2502599306, UnsafePtr: unsafe.Pointer(Rebox_Main_1486975205_1539302174((&Constructor_Main_Tagged[int64]{1, __eta_norm_1_0, __eta_norm_0_1})))}
	}
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

func Call_Main_run(count_0_loop int64, seed_1_loop int64) int64 {
	var count_0 int64 = count_0_loop
	_ = count_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var Call_local_Main_go__2528240558_2_0_0 func(int64, struct {
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
	_ = Call_local_Main_go__2528240558_2_0_0
	var go__2528240558_2_0_0 gopurs_runtime.Value
	_ = go__2528240558_2_0_0
	var Call_local_Main_go__go_2_1_1 func(int64, struct {
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
	_ = Call_local_Main_go__go_2_1_1
	var go__go_2_1_1 gopurs_runtime.Value
	_ = go__go_2_1_1
	Call_local_Main_go__2528240558_2_0_0 = func(n_3_loop int64, packet_4_loop struct {
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
	}, total_5_loop int64) int64 {
	go__2528240558_2_0_0:
		for {
			if false {
				continue go__2528240558_2_0_0
			}
			var n_3 int64 = n_3_loop
			_ = n_3
			var packet_4 struct {
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
			} = packet_4_loop
			_ = packet_4
			var total_5 int64 = total_5_loop
			_ = total_5
			var __t3 int64
			{
				if (n_3) <= (int64(0)) {
					__t3 = total_5
					goto end_branch_3
				} else {

				}
			}
			{
				// TAST (Let): next_6_2 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Main","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Main","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
				next_6_2 := Call_Main_bump(int64(1), packet_4)
				_ = next_6_2
				n_3_loop = gopurs_runtime.IntSub(n_3, int64(1))
				packet_4_loop = next_6_2
				total_5_loop = gopurs_runtime.IntAdd(total_5, Call_Main_score(next_6_2))
				continue go__2528240558_2_0_0
				__t3 = func() int64 { panic("unreachable") }()
			}
		end_branch_3:
			return __t3
		}
	}
	go__2528240558_2_0_0 = gopurs_runtime.Func(func(n_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(packet_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(total_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Main_go__2528240558_2_0_0(n_3_loop_val.IntVal, func() struct {
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
					orig := packet_4_loop_val
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
				}(), total_5_loop_val.IntVal))
			})
		})
	})
	Call_local_Main_go__go_2_1_1 = func(n_3_loop int64, packet_4_loop struct {
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
	}, total_5_loop int64) int64 {
	go__go_2_1_1:
		for {
			if false {
				continue go__go_2_1_1
			}
			var n_3 int64 = n_3_loop
			_ = n_3
			var packet_4 struct {
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
			} = packet_4_loop
			_ = packet_4
			var total_5 int64 = total_5_loop
			_ = total_5
			var __t5 int64
			{
				if (n_3) <= (int64(0)) {
					__t5 = total_5
					goto end_branch_5
				} else {

				}
			}
			{
				// TAST (Let): next_6_4 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Main","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Main","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
				next_6_4 := Call_Main_bump(int64(1), packet_4)
				_ = next_6_4
				__t5 = Call_local_Main_go__2528240558_2_0_0(gopurs_runtime.IntSub(n_3, int64(1)), next_6_4, gopurs_runtime.IntAdd(total_5, Call_Main_score(next_6_4)))
			}
		end_branch_5:
			return __t5
		}
	}
	go__go_2_1_1 = gopurs_runtime.Func(func(n_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(packet_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(total_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Main_go__go_2_1_1(n_3_loop_val.IntVal, func() struct {
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
					orig := packet_4_loop_val
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
				}(), total_5_loop_val.IntVal))
			})
		})
	})
	return Call_local_Main_go__2528240558_2_0_0(count_0, Call_Main_makePacket(seed_1), int64(0))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1486975205_1539302174(in *Constructor_Main_Tagged[int64]) *Constructor_Main_Tagged[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Tagged[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_1938252319_3483727428(in *Constructor_Main_Payload[int64]) *Constructor_Main_Payload[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Payload[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys4{K0: "enabled", K1: "left", K2: "payload", K3: "right"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys4{K0: "active", K1: "column", K2: "line", K3: "note"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys3{K0: "active", K1: "label", K2: "score"}
