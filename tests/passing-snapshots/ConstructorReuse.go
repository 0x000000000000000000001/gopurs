package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_go__append gopurs_runtime.Value
var once_Main_go__append sync.Once

func Get_Main_go__append() gopurs_runtime.Value {
	once_Main_go__append.Do(func() {
		cache_Main_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Main_go__append
}

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Get_Data_Show_show__1986550628()
	})
	return cache_Main_show
}

var cache_Main_show1 gopurs_runtime.Value
var once_Main_show1 sync.Once

func Get_Main_show1() gopurs_runtime.Value {
	once_Main_show1.Do(func() {
		cache_Main_show1 = Get_Data_Show_show__2877300819()
	})
	return cache_Main_show1
}

var cache_Main_lessThan gopurs_runtime.Value
var once_Main_lessThan sync.Once

func Get_Main_lessThan() gopurs_runtime.Value {
	once_Main_lessThan.Do(func() {
		cache_Main_lessThan = Get_Data_Ord_lessThan__756116985()
	})
	return cache_Main_lessThan
}

var cache_Main_div gopurs_runtime.Value
var once_Main_div sync.Once

func Get_Main_div() gopurs_runtime.Value {
	once_Main_div.Do(func() {
		cache_Main_div = Get_Data_EuclideanRing_div__266314078()
	})
	return cache_Main_div
}

var cache_Main_assertEqual gopurs_runtime.Value
var once_Main_assertEqual sync.Once

func Get_Main_assertEqual() gopurs_runtime.Value {
	once_Main_assertEqual.Do(func() {
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__1066131861()
	})
	return cache_Main_assertEqual
}

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = Get_Control_Bind_bind__2849971988()
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

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = Get_Control_Bind_bind__345900052()
	})
	return cache_Main_bind1
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__519557204()
	})
	return cache_Main_bind2
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__2646313556()
	})
	return cache_Main_bind3
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__3990910100()
	})
	return cache_Main_bind4
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Bind_bind__2862182516()
	})
	return cache_Main_bind5
}

var cache_Main_bind6 gopurs_runtime.Value
var once_Main_bind6 sync.Once

func Get_Main_bind6() gopurs_runtime.Value {
	once_Main_bind6.Do(func() {
		cache_Main_bind6 = Get_Control_Bind_bind__1687364500()
	})
	return cache_Main_bind6
}

var cache_Main_bind7 gopurs_runtime.Value
var once_Main_bind7 sync.Once

func Get_Main_bind7() gopurs_runtime.Value {
	once_Main_bind7.Do(func() {
		cache_Main_bind7 = Get_Control_Bind_bind__3960688980()
	})
	return cache_Main_bind7
}

var cache_Main_negate gopurs_runtime.Value
var once_Main_negate sync.Once

func Get_Main_negate() gopurs_runtime.Value {
	once_Main_negate.Do(func() {
		cache_Main_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_negate(a_0_box.FloatVal()))
		})
	})
	return cache_Main_negate
}

var cache_Main_bind8 gopurs_runtime.Value
var once_Main_bind8 sync.Once

func Get_Main_bind8() gopurs_runtime.Value {
	once_Main_bind8.Do(func() {
		cache_Main_bind8 = Get_Control_Bind_bind__2743947764()
	})
	return cache_Main_bind8
}

var cache_Main_assertEqual1 gopurs_runtime.Value
var once_Main_assertEqual1 sync.Once

func Get_Main_assertEqual1() gopurs_runtime.Value {
	once_Main_assertEqual1.Do(func() {
		cache_Main_assertEqual1 = Get_Test_Assert_assertEqual__2127306453()
	})
	return cache_Main_assertEqual1
}

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = Get_Test_Assert_assertEqual__2867988789()
	})
	return cache_Main_assertEqual2
}

var cache_Main_NumberCell gopurs_runtime.Value
var once_Main_NumberCell sync.Once

func Get_Main_NumberCell() gopurs_runtime.Value {
	once_Main_NumberCell.Do(func() {
		cache_Main_NumberCell = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1758940859, UnsafePtr: unsafe.Pointer((&Constructor_Main_NumberCell{1, value0.FloatVal(), value1.IntVal}))}
			})
		})
	})
	return cache_Main_NumberCell
}

var cache_Main_Red gopurs_runtime.Value
var once_Main_Red sync.Once

func Get_Main_Red() gopurs_runtime.Value {
	once_Main_Red.Do(func() {
		cache_Main_Red = gopurs_runtime.Value{Type: 9, IntVal: int64(1227005933), UnsafePtr: nil}
	})
	return cache_Main_Red
}

var cache_Main_Black gopurs_runtime.Value
var once_Main_Black sync.Once

func Get_Main_Black() gopurs_runtime.Value {
	once_Main_Black.Do(func() {
		cache_Main_Black = gopurs_runtime.Value{Type: 9, IntVal: int64(939353081), UnsafePtr: nil}
	})
	return cache_Main_Black
}

var cache_Main_First gopurs_runtime.Value
var once_Main_First sync.Once

func Get_Main_First() gopurs_runtime.Value {
	once_Main_First.Do(func() {
		cache_Main_First = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2444816644, UnsafePtr: unsafe.Pointer((&Constructor_Main_First{1, uint32(value0.IntVal), value1.IntVal}))}
			})
		})
	})
	return cache_Main_First
}

var cache_Main_Second gopurs_runtime.Value
var once_Main_Second sync.Once

func Get_Main_Second() gopurs_runtime.Value {
	once_Main_Second.Do(func() {
		cache_Main_Second = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1617734670, UnsafePtr: unsafe.Pointer((&Constructor_Main_Second{1, uint32(value0.IntVal), value1.IntVal}))}
			})
		})
	})
	return cache_Main_Second
}

var cache_Main_Empty gopurs_runtime.Value
var once_Main_Empty sync.Once

func Get_Main_Empty() gopurs_runtime.Value {
	once_Main_Empty.Do(func() {
		cache_Main_Empty = gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer((*Constructor_Main_Node)(nil))}
	})
	return cache_Main_Empty
}

var cache_Main_Node gopurs_runtime.Value
var once_Main_Node sync.Once

func Get_Main_Node() gopurs_runtime.Value {
	once_Main_Node.Do(func() {
		cache_Main_Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer((&Constructor_Main_Node{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_Node](value1), value2.IntVal, (value3.IntVal) != (0), gopurs_runtime.CoerceToStruct[Constructor_Main_Node](value4)}))}
						})
					})
				})
			})
		})
	})
	return cache_Main_Node
}

var cache_Main_toSecond gopurs_runtime.Value
var once_Main_toSecond sync.Once

func Get_Main_toSecond() gopurs_runtime.Value {
	once_Main_toSecond.Do(func() {
		cache_Main_toSecond = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_toSecond(v_0_box)
		})
	})
	return cache_Main_toSecond
}

var cache_Main_swapChildren gopurs_runtime.Value
var once_Main_swapChildren sync.Once

func Get_Main_swapChildren() gopurs_runtime.Value {
	once_Main_swapChildren.Do(func() {
		cache_Main_swapChildren = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_swapChildren(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v_0_box)))}
		})
	})
	return cache_Main_swapChildren
}

var cache_Main_setTrue gopurs_runtime.Value
var once_Main_setTrue sync.Once

func Get_Main_setTrue() gopurs_runtime.Value {
	once_Main_setTrue.Do(func() {
		cache_Main_setTrue = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v_0_box)))}
		})
	})
	return cache_Main_setTrue
}

var cache_Main_setSeven gopurs_runtime.Value
var once_Main_setSeven sync.Once

func Get_Main_setSeven() gopurs_runtime.Value {
	once_Main_setSeven.Do(func() {
		cache_Main_setSeven = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v_0_box)))}
		})
	})
	return cache_Main_setSeven
}

var cache_Main_setBlackTrue gopurs_runtime.Value
var once_Main_setBlackTrue sync.Once

func Get_Main_setBlackTrue() gopurs_runtime.Value {
	once_Main_setBlackTrue.Do(func() {
		cache_Main_setBlackTrue = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setBlackTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v_0_box)))}
		})
	})
	return cache_Main_setBlackTrue
}

var cache_Main_setBlack gopurs_runtime.Value
var once_Main_setBlack sync.Once

func Get_Main_setBlack() gopurs_runtime.Value {
	once_Main_setBlack.Do(func() {
		cache_Main_setBlack = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v_0_box)))}
		})
	})
	return cache_Main_setBlack
}

var cache_Main_positiveZero gopurs_runtime.Value
var once_Main_positiveZero sync.Once

func Get_Main_positiveZero() gopurs_runtime.Value {
	once_Main_positiveZero.Do(func() {
		cache_Main_positiveZero = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1758940859, UnsafePtr: unsafe.Pointer(Call_Main_positiveZero(gopurs_runtime.CoerceToStruct[Constructor_Main_NumberCell](v_0_box)))}
		})
	})
	return cache_Main_positiveZero
}

var cache_Main_numberMarker gopurs_runtime.Value
var once_Main_numberMarker sync.Once

func Get_Main_numberMarker() gopurs_runtime.Value {
	once_Main_numberMarker.Do(func() {
		cache_Main_numberMarker = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_numberMarker(gopurs_runtime.CoerceToStruct[Constructor_Main_NumberCell](v_0_box)))
		})
	})
	return cache_Main_numberMarker
}

var cache_Main_nodeText gopurs_runtime.Value
var once_Main_nodeText sync.Once

func Get_Main_nodeText() gopurs_runtime.Value {
	once_Main_nodeText.Do(func() {
		cache_Main_nodeText = gopurs_runtime.Func5(func(color_0_box gopurs_runtime.Value, key_1_box gopurs_runtime.Value, active_2_box gopurs_runtime.Value, left_3_box gopurs_runtime.Value, right_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_nodeText(color_0_box.StrVal(), key_1_box.IntVal, (active_2_box.IntVal) != (0), left_3_box.StrVal(), right_4_box.StrVal()))
		})
	})
	return cache_Main_nodeText
}

var cache_Main_mixSources gopurs_runtime.Value
var once_Main_mixSources sync.Once

func Get_Main_mixSources() gopurs_runtime.Value {
	once_Main_mixSources.Do(func() {
		cache_Main_mixSources = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_mixSources(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v1_1_box)))}
		})
	})
	return cache_Main_mixSources
}

var cache_Main_isNegativeZero gopurs_runtime.Value
var once_Main_isNegativeZero sync.Once

func Get_Main_isNegativeZero() gopurs_runtime.Value {
	once_Main_isNegativeZero.Do(func() {
		cache_Main_isNegativeZero = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_isNegativeZero(gopurs_runtime.CoerceToStruct[Constructor_Main_NumberCell](v_0_box)))
		})
	})
	return cache_Main_isNegativeZero
}

var cache_Main_colorName gopurs_runtime.Value
var once_Main_colorName sync.Once

func Get_Main_colorName() gopurs_runtime.Value {
	once_Main_colorName.Do(func() {
		cache_Main_colorName = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_colorName(uint32(v_0_box.IntVal)))
		})
	})
	return cache_Main_colorName
}

var cache_Main_render gopurs_runtime.Value
var once_Main_render sync.Once

func Get_Main_render() gopurs_runtime.Value {
	once_Main_render.Do(func() {
		cache_Main_render = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](v_0_box)))
		})
	})
	return cache_Main_render
}

var cache_Main_expectTree gopurs_runtime.Value
var once_Main_expectTree sync.Once

func Get_Main_expectTree() gopurs_runtime.Value {
	once_Main_expectTree.Do(func() {
		cache_Main_expectTree = gopurs_runtime.Func2(func(expected_0_box gopurs_runtime.Value, actual_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_expectTree(expected_0_box.StrVal(), gopurs_runtime.CoerceToStruct[Constructor_Main_Node](actual_1_box))
		})
	})
	return cache_Main_expectTree
}

var cache_Main_renderTagged gopurs_runtime.Value
var once_Main_renderTagged sync.Once

func Get_Main_renderTagged() gopurs_runtime.Value {
	once_Main_renderTagged.Do(func() {
		cache_Main_renderTagged = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_renderTagged(v_0_box))
		})
	})
	return cache_Main_renderTagged
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): redText_0_0 shape=App(Var) bindingType=String
			redText_0_0 := Call_Main_nodeText("R", int64(17), false, "B:-3:true[E][E]", "R:9:false[E][E]")
			_ = redText_0_0
			// TAST (Let): sevenText_1_1 shape=App(Var) bindingType=String
			sevenText_1_1 := Call_Main_nodeText("R", int64(7), false, "B:-3:true[E][E]", "R:9:false[E][E]")
			_ = sevenText_1_1
			// TAST (Let): trueText_2_2 shape=App(Var) bindingType=String
			trueText_2_2 := Call_Main_nodeText("R", int64(17), true, "B:-3:true[E][E]", "R:9:false[E][E]")
			_ = trueText_2_2
			// TAST (Let): blackText_3_3 shape=App(Var) bindingType=String
			blackText_3_3 := Call_Main_nodeText("B", int64(17), false, "B:-3:true[E][E]", "R:9:false[E][E]")
			_ = blackText_3_3
			__callback_fn_5 := Get_Effect_Ref__new()
			_ = __callback_fn_5
			__callback_arg_6 := struct {
				black   *Constructor_Main_Node
				empty   *Constructor_Main_Node
				flagged *Constructor_Main_Node
				other   *Constructor_Main_Node
				red     *Constructor_Main_Node
				seven   *Constructor_Main_Node
				/*gopurs:record-v1:6:0062006c00610063006b,0065006d007000740079,0066006c00610067006700650064,006f0074006800650072,007200650064,0073006500760065006e*/
			}{(&Constructor_Main_Node{1, 939353081, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(17), false, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})}), (*Constructor_Main_Node)(nil), (&Constructor_Main_Node{1, 1227005933, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(17), true, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})}), (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(91), true, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(42), true, (*Constructor_Main_Node)(nil)})}), (&Constructor_Main_Node{1, 1227005933, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(17), false, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})}), (&Constructor_Main_Node{1, 1227005933, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(7), false, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})})}
			_ = __callback_arg_6
			// TAST (Let): __local_var_4_4 shape=App(Var) bindingType=Any
			__local_var_4_4 := func(__native_callback func(struct {
				black   *Constructor_Main_Node
				empty   *Constructor_Main_Node
				flagged *Constructor_Main_Node
				other   *Constructor_Main_Node
				red     *Constructor_Main_Node
				seven   *Constructor_Main_Node
				/*gopurs:record-v1:6:0062006c00610063006b,0065006d007000740079,0066006c00610067006700650064,006f0074006800650072,007200650064,0073006500760065006e*/
			}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
				if __native_callback_ok {
					return __native_callback(__callback_arg_6)
				} else {

				}
				return gopurs_runtime.Apply(__callback_fn_5, func() gopurs_runtime.Value {
					orig := __callback_arg_6
					_ = orig
					return gopurs_runtime.RecordDict([]string{"black", "empty", "flagged", "other", "red", "seven"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.black)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.empty)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.flagged)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.other)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.red)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.seven)}})
				}())
			}(gopurs_runtime.FunctionData[func(struct {
				black   *Constructor_Main_Node
				empty   *Constructor_Main_Node
				flagged *Constructor_Main_Node
				other   *Constructor_Main_Node
				red     *Constructor_Main_Node
				seven   *Constructor_Main_Node
				/*gopurs:record-v1:6:0062006c00610063006b,0065006d007000740079,0066006c00610067006700650064,006f0074006800650072,007200650064,0073006500760065006e*/
			}) gopurs_runtime.Value](__callback_fn_5))
			_ = __local_var_4_4
			__local_var_5_7 := gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Value{})
			_ = __local_var_5_7
			__local_var_6_8 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_5_7, gopurs_runtime.Value{})
			_ = __local_var_6_8
			__local_var_7_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "red")))), blackText_3_3}), gopurs_runtime.Value{})
			_ = __local_var_7_9
			__local_var_8_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "black")))), blackText_3_3}), gopurs_runtime.Value{})
			_ = __local_var_8_10
			__local_var_9_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "empty")))), "E"}), gopurs_runtime.Value{})
			_ = __local_var_9_11
			__local_var_10_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "red")))), sevenText_1_1}), gopurs_runtime.Value{})
			_ = __local_var_10_12
			__local_var_11_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "seven")))), sevenText_1_1}), gopurs_runtime.Value{})
			_ = __local_var_11_13
			__local_var_12_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "red")))), trueText_2_2}), gopurs_runtime.Value{})
			_ = __local_var_12_14
			__local_var_13_15 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "flagged")))), trueText_2_2}), gopurs_runtime.Value{})
			_ = __local_var_13_15
			__local_var_14_16 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "red"))), redText_0_0}), gopurs_runtime.Value{})
			_ = __local_var_14_16
			__local_var_15_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "black"))), blackText_3_3}), gopurs_runtime.Value{})
			_ = __local_var_15_17
			__local_var_16_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "empty")))), "E"}), gopurs_runtime.Value{})
			_ = __local_var_16_18
			__local_var_17_19 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "empty")))), "E"}), gopurs_runtime.Value{})
			_ = __local_var_17_19
			__local_var_18_20 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "red"))))}, gopurs_runtime.Value{})
			_ = __local_var_18_20
			__local_var_19_21 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_18_20, gopurs_runtime.Value{})
			_ = __local_var_19_21
			__local_var_20_22 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](__local_var_19_21)))}, gopurs_runtime.Value{})
			_ = __local_var_20_22
			__local_var_21_23 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_20_22, gopurs_runtime.Value{})
			_ = __local_var_21_23
			__local_var_22_24 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](__local_var_21_23)), Call_Main_nodeText("B", int64(7), false, "B:-3:true[E][E]", "R:9:false[E][E]")}), gopurs_runtime.Value{})
			_ = __local_var_22_24
			__local_var_23_25 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](__local_var_19_21)), blackText_3_3}), gopurs_runtime.Value{})
			_ = __local_var_23_25
			__local_var_24_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "red"))), redText_0_0}), gopurs_runtime.Value{})
			_ = __local_var_24_26
			__local_var_25_27 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_swapChildren(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "black")))), Call_Main_nodeText("B", int64(17), false, "R:9:false[E][E]", "B:-3:true[E][E]")}), gopurs_runtime.Value{})
			_ = __local_var_25_27
			__local_var_26_28 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_mixSources(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "black")), gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "other")))), Call_Main_nodeText("B", int64(17), false, "B:-3:true[E][E]", "B:42:true[E][E]")}), gopurs_runtime.Value{})
			_ = __local_var_26_28
			__local_var_27_29 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "black"))), blackText_3_3}), gopurs_runtime.Value{})
			_ = __local_var_27_29
			__local_var_28_30 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(Call_Main_setBlackTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "black")))), Call_Main_nodeText("B", int64(17), true, "B:-3:true[E][E]", "R:9:false[E][E]")}), gopurs_runtime.Value{})
			_ = __local_var_28_30
			__local_var_29_31 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(__local_var_6_8, "black"))), blackText_3_3}), gopurs_runtime.Value{})
			_ = __local_var_29_31
			__local_var_30_32 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2444816644, UnsafePtr: unsafe.Pointer((&Constructor_Main_First{1, 939353081, int64(41)}))}, gopurs_runtime.Value{})
			_ = __local_var_30_32
			__local_var_31_33 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_30_32, gopurs_runtime.Value{})
			_ = __local_var_31_33
			__local_var_32_34 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_renderTagged(Call_Main_toSecond(__local_var_31_33)), "second:B:41"}), gopurs_runtime.Value{})
			_ = __local_var_32_34
			__local_var_33_35 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_renderTagged(__local_var_31_33), "first:B:41"}), gopurs_runtime.Value{})
			_ = __local_var_33_35
			__local_var_34_36 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Float(0.0), gopurs_runtime.Value{})
			_ = __local_var_34_36
			__local_var_35_37 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_34_36, gopurs_runtime.Value{})
			_ = __local_var_35_37
			__local_var_36_38 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 1758940859, UnsafePtr: unsafe.Pointer((&Constructor_Main_NumberCell{1, -(__local_var_35_37.FloatVal()), int64(5)}))}, gopurs_runtime.Value{})
			_ = __local_var_36_38
			__local_var_37_39 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_36_38, gopurs_runtime.Value{})
			_ = __local_var_37_39
			__local_var_38_40 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{((1.0) / ((*Constructor_Main_NumberCell)(__local_var_37_39.UnsafePtr).V0)) < (0.0), true}), gopurs_runtime.Value{})
			_ = __local_var_38_40
			__local_var_39_41 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 1758940859, UnsafePtr: unsafe.Pointer(Call_Main_positiveZero(gopurs_runtime.CoerceToStruct[Constructor_Main_NumberCell](__local_var_37_39)))}, gopurs_runtime.Value{})
			_ = __local_var_39_41
			__local_var_40_42 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_39_41, gopurs_runtime.Value{})
			_ = __local_var_40_42
			__local_var_41_43 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{((1.0) / ((*Constructor_Main_NumberCell)(__local_var_40_42.UnsafePtr).V0)) < (0.0), false}), gopurs_runtime.Value{})
			_ = __local_var_41_43
			__local_var_42_44 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{((1.0) / ((*Constructor_Main_NumberCell)(__local_var_37_39.UnsafePtr).V0)) < (0.0), true}), gopurs_runtime.Value{})
			_ = __local_var_42_44
			__local_var_43_45 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(*Constructor_Main_NumberCell)(__local_var_40_42.UnsafePtr).V1, int64(5)}), gopurs_runtime.Value{})
			_ = __local_var_43_45
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main___gopurs_owned_mixSources_0_consume(__arg0 *Constructor_Main_Node, __arg1 *Constructor_Main_Node, __donor *Constructor_Main_Node) *Constructor_Main_Node {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if ((__arg0) != (nil)) && ((__arg1) != (nil)) {
			__scalar_3 := uint32(939353081)
			_ = __scalar_3
			__read_4 := __arg0.V1
			_ = __read_4
			__scalar_5 := int64(__arg0.V2)
			_ = __scalar_5
			__scalar_6 := bool(__arg0.V3)
			_ = __scalar_6
			__read_7 := __arg1.V4
			_ = __read_7
			__donor_slot_10 := __donor
			_ = __donor_slot_10
			__dead_8 := __arg0
			_ = __dead_8
			__dead_9 := __arg1
			_ = __dead_9
			__dead_8.Rc = 1
			__dead_8.V0 = __scalar_3
			__dead_8.V1 = __read_4
			__dead_8.V2 = __scalar_5
			__dead_8.V3 = __scalar_6
			__dead_8.V4 = __read_7
			return __dead_8
		} else {
			__read_0 := __arg0
			_ = __read_0
			__donor_slot_2 := __donor
			_ = __donor_slot_2
			__dead_1 := __arg1
			_ = __dead_1
			return __read_0
		}
	}
}

func Call_Main___gopurs_owned_mixSources_0(__arg0 *Constructor_Main_Node, __arg1 *Constructor_Main_Node) *Constructor_Main_Node {
	return Call_Main___gopurs_owned_mixSources_0_consume(__arg0, __arg1, nil)
}

func Call_Main___gopurs_owned_setBlack_0_consume(__arg0 *Constructor_Main_Node, __donor *Constructor_Main_Node) *Constructor_Main_Node {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) != (nil) {
			__scalar_0 := uint32(939353081)
			_ = __scalar_0
			__read_1 := __arg0.V1
			_ = __read_1
			__scalar_2 := int64(__arg0.V2)
			_ = __scalar_2
			__scalar_3 := bool(__arg0.V3)
			_ = __scalar_3
			__read_4 := __arg0.V4
			_ = __read_4
			__donor_slot_6 := __donor
			_ = __donor_slot_6
			__dead_5 := __arg0
			_ = __dead_5
			__dead_5.Rc = 1
			__dead_5.V0 = __scalar_0
			__dead_5.V1 = __read_1
			__dead_5.V2 = __scalar_2
			__dead_5.V3 = __scalar_3
			__dead_5.V4 = __read_4
			return __dead_5
		} else {
			if (__arg0) == (nil) {
				__donor_slot_8 := __donor
				_ = __donor_slot_8
				__dead_7 := __arg0
				_ = __dead_7
				return nil
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_setBlack_0(__arg0 *Constructor_Main_Node) *Constructor_Main_Node {
	return Call_Main___gopurs_owned_setBlack_0_consume(__arg0, nil)
}

func Call_Main___gopurs_owned_setBlackTrue_0_consume(__arg0 *Constructor_Main_Node, __donor *Constructor_Main_Node) *Constructor_Main_Node {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) != (nil) {
			__scalar_0 := uint32(939353081)
			_ = __scalar_0
			__read_1 := __arg0.V1
			_ = __read_1
			__scalar_2 := int64(__arg0.V2)
			_ = __scalar_2
			__scalar_3 := bool(true)
			_ = __scalar_3
			__read_4 := __arg0.V4
			_ = __read_4
			__donor_slot_6 := __donor
			_ = __donor_slot_6
			__dead_5 := __arg0
			_ = __dead_5
			__dead_5.Rc = 1
			__dead_5.V0 = __scalar_0
			__dead_5.V1 = __read_1
			__dead_5.V2 = __scalar_2
			__dead_5.V3 = __scalar_3
			__dead_5.V4 = __read_4
			return __dead_5
		} else {
			if (__arg0) == (nil) {
				__donor_slot_8 := __donor
				_ = __donor_slot_8
				__dead_7 := __arg0
				_ = __dead_7
				return nil
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_setBlackTrue_0(__arg0 *Constructor_Main_Node) *Constructor_Main_Node {
	return Call_Main___gopurs_owned_setBlackTrue_0_consume(__arg0, nil)
}

func Call_Main___gopurs_owned_setSeven_0_consume(__arg0 *Constructor_Main_Node, __donor *Constructor_Main_Node) *Constructor_Main_Node {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) != (nil) {
			__scalar_0 := uint32(__arg0.V0)
			_ = __scalar_0
			__read_1 := __arg0.V1
			_ = __read_1
			__scalar_2 := int64(int64(7))
			_ = __scalar_2
			__scalar_3 := bool(__arg0.V3)
			_ = __scalar_3
			__read_4 := __arg0.V4
			_ = __read_4
			__donor_slot_6 := __donor
			_ = __donor_slot_6
			__dead_5 := __arg0
			_ = __dead_5
			__dead_5.Rc = 1
			__dead_5.V0 = __scalar_0
			__dead_5.V1 = __read_1
			__dead_5.V2 = __scalar_2
			__dead_5.V3 = __scalar_3
			__dead_5.V4 = __read_4
			return __dead_5
		} else {
			if (__arg0) == (nil) {
				__donor_slot_8 := __donor
				_ = __donor_slot_8
				__dead_7 := __arg0
				_ = __dead_7
				return nil
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_setSeven_0(__arg0 *Constructor_Main_Node) *Constructor_Main_Node {
	return Call_Main___gopurs_owned_setSeven_0_consume(__arg0, nil)
}

func Call_Main___gopurs_owned_setTrue_0_consume(__arg0 *Constructor_Main_Node, __donor *Constructor_Main_Node) *Constructor_Main_Node {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) != (nil) {
			__scalar_0 := uint32(__arg0.V0)
			_ = __scalar_0
			__read_1 := __arg0.V1
			_ = __read_1
			__scalar_2 := int64(__arg0.V2)
			_ = __scalar_2
			__scalar_3 := bool(true)
			_ = __scalar_3
			__read_4 := __arg0.V4
			_ = __read_4
			__donor_slot_6 := __donor
			_ = __donor_slot_6
			__dead_5 := __arg0
			_ = __dead_5
			__dead_5.Rc = 1
			__dead_5.V0 = __scalar_0
			__dead_5.V1 = __read_1
			__dead_5.V2 = __scalar_2
			__dead_5.V3 = __scalar_3
			__dead_5.V4 = __read_4
			return __dead_5
		} else {
			if (__arg0) == (nil) {
				__donor_slot_8 := __donor
				_ = __donor_slot_8
				__dead_7 := __arg0
				_ = __dead_7
				return nil
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_setTrue_0(__arg0 *Constructor_Main_Node) *Constructor_Main_Node {
	return Call_Main___gopurs_owned_setTrue_0_consume(__arg0, nil)
}

func Call_Main___gopurs_owned_swapChildren_0_consume(__arg0 *Constructor_Main_Node, __donor *Constructor_Main_Node) *Constructor_Main_Node {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) != (nil) {
			__scalar_0 := uint32(939353081)
			_ = __scalar_0
			__read_1 := __arg0.V4
			_ = __read_1
			__scalar_2 := int64(__arg0.V2)
			_ = __scalar_2
			__scalar_3 := bool(__arg0.V3)
			_ = __scalar_3
			__read_4 := __arg0.V1
			_ = __read_4
			__donor_slot_6 := __donor
			_ = __donor_slot_6
			__dead_5 := __arg0
			_ = __dead_5
			__dead_5.Rc = 1
			__dead_5.V0 = __scalar_0
			__dead_5.V1 = __read_1
			__dead_5.V2 = __scalar_2
			__dead_5.V3 = __scalar_3
			__dead_5.V4 = __read_4
			return __dead_5
		} else {
			if (__arg0) == (nil) {
				__donor_slot_8 := __donor
				_ = __donor_slot_8
				__dead_7 := __arg0
				_ = __dead_7
				return nil
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_swapChildren_0(__arg0 *Constructor_Main_Node) *Constructor_Main_Node {
	return Call_Main___gopurs_owned_swapChildren_0_consume(__arg0, nil)
}

type Constructor_Main_NumberCell struct {
	Rc uint32
	V0 float64
	V1 int64
}

type Constructor_Main_Red struct {
	Rc uint32
}

type Constructor_Main_Black struct {
	Rc uint32
}

type Constructor_Main_First struct {
	Rc uint32
	V0 uint32
	V1 int64
}

type Constructor_Main_Second struct {
	Rc uint32
	V0 uint32
	V1 int64
}

type Constructor_Main_Empty struct {
	Rc uint32
}

type Constructor_Main_Node struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_Main_Node
	V2 int64
	V3 bool
	V4 *Constructor_Main_Node
}

func Call_Main_negate(a_0_loop float64) float64 {
	var a_0 float64 = a_0_loop
	_ = a_0
	return -(a_0)
}

func Call_Main_toSecond(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 *Constructor_Main_Second
	{
		if v_0.Type == 9 && v_0.IntVal == 2444816644 {
			__t0 = (&Constructor_Main_Second{1, 939353081, (*Constructor_Main_First)(v_0.UnsafePtr).V1})
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 1617734670 {
			__t0 = (&Constructor_Main_Second{1, 939353081, (*Constructor_Main_Second)(v_0.UnsafePtr).V1})
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Second](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_0:
	return gopurs_runtime.Value{Type: 9, IntVal: 1617734670, UnsafePtr: unsafe.Pointer(__t0)}
}

func Call_Main_swapChildren(v_0_loop *Constructor_Main_Node) *Constructor_Main_Node {
	var v_0 *Constructor_Main_Node = v_0_loop
	_ = v_0
	var __t0 *Constructor_Main_Node
	{
		if v_0 != nil {
			__t0 = (&Constructor_Main_Node{1, 939353081, (v_0).V4, (v_0).V2, (v_0).V3, (v_0).V1})
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == nil {
			__t0 = (*Constructor_Main_Node)(nil)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Main_Node { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_setTrue(v_0_loop *Constructor_Main_Node) *Constructor_Main_Node {
	var v_0 *Constructor_Main_Node = v_0_loop
	_ = v_0
	var __t1 *Constructor_Main_Node
	{
		if v_0 != nil {
			var __reuse_0 *Constructor_Main_Node
			if ((v_0) != (nil)) && (((v_0).V3) == (true)) {
				__reuse_0 = v_0
			} else {
				__reuse_0 = (&Constructor_Main_Node{1, (v_0).V0, (v_0).V1, (v_0).V2, true, (v_0).V4})
			}
			__t1 = __reuse_0
			goto end_branch_1
		} else {

		}
	}
	{
		if v_0 == nil {
			__t1 = (*Constructor_Main_Node)(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Main_Node { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Main_setSeven(v_0_loop *Constructor_Main_Node) *Constructor_Main_Node {
	var v_0 *Constructor_Main_Node = v_0_loop
	_ = v_0
	var __t1 *Constructor_Main_Node
	{
		if v_0 != nil {
			var __reuse_0 *Constructor_Main_Node
			if ((v_0) != (nil)) && (((v_0).V2) == (int64(7))) {
				__reuse_0 = v_0
			} else {
				__reuse_0 = (&Constructor_Main_Node{1, (v_0).V0, (v_0).V1, int64(7), (v_0).V3, (v_0).V4})
			}
			__t1 = __reuse_0
			goto end_branch_1
		} else {

		}
	}
	{
		if v_0 == nil {
			__t1 = (*Constructor_Main_Node)(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Main_Node { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Main_setBlackTrue(v_0_loop *Constructor_Main_Node) *Constructor_Main_Node {
	var v_0 *Constructor_Main_Node = v_0_loop
	_ = v_0
	var __t0 *Constructor_Main_Node
	{
		if v_0 != nil {
			__t0 = (&Constructor_Main_Node{1, 939353081, (v_0).V1, (v_0).V2, true, (v_0).V4})
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == nil {
			__t0 = (*Constructor_Main_Node)(nil)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Main_Node { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_setBlack(v_0_loop *Constructor_Main_Node) *Constructor_Main_Node {
	var v_0 *Constructor_Main_Node = v_0_loop
	_ = v_0
	var __t1 *Constructor_Main_Node
	{
		if v_0 != nil {
			var __reuse_0 *Constructor_Main_Node
			if ((v_0) != (nil)) && (((v_0).V0) == (939353081)) {
				__reuse_0 = v_0
			} else {
				__reuse_0 = (&Constructor_Main_Node{1, 939353081, (v_0).V1, (v_0).V2, (v_0).V3, (v_0).V4})
			}
			__t1 = __reuse_0
			goto end_branch_1
		} else {

		}
	}
	{
		if v_0 == nil {
			__t1 = (*Constructor_Main_Node)(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Main_Node { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Main_positiveZero(v_0_loop *Constructor_Main_NumberCell) *Constructor_Main_NumberCell {
	var v_0 *Constructor_Main_NumberCell = v_0_loop
	_ = v_0
	return (&Constructor_Main_NumberCell{1, 0.0, (v_0).V1})
}

func Call_Main_numberMarker(v_0_loop *Constructor_Main_NumberCell) int64 {
	var v_0 *Constructor_Main_NumberCell = v_0_loop
	_ = v_0
	return (v_0).V1
}

func Call_Main_nodeText(color_0_loop string, key_1_loop int64, active_2_loop bool, left_3_loop string, right_4_loop string) string {
	var color_0 string = color_0_loop
	_ = color_0
	var key_1 int64 = key_1_loop
	_ = key_1
	var active_2 bool = active_2_loop
	_ = active_2
	var left_3 string = left_3_loop
	_ = left_3
	var right_4 string = right_4_loop
	_ = right_4
	var __t0 string
	{
		if active_2 {
			__t0 = "true"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = "false"
	}
end_branch_0:
	return gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(color_0, ":"), Data_Show_ShowIntImpl(key_1)), ":"), __t0), "["), left_3), "]["), right_4), "]")
}

func Call_Main_mixSources(v_0_loop *Constructor_Main_Node, v1_1_loop *Constructor_Main_Node) *Constructor_Main_Node {
	var v_0 *Constructor_Main_Node = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Main_Node = v1_1_loop
	_ = v1_1
	var __t0 *Constructor_Main_Node
	{
		if (v_0 != nil) && (v1_1 != nil) {
			__t0 = (&Constructor_Main_Node{1, 939353081, (v_0).V1, (v_0).V2, (v_0).V3, (v1_1).V4})
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = v_0
	}
end_branch_0:
	return __t0
}

func Call_Main_isNegativeZero(v_0_loop *Constructor_Main_NumberCell) bool {
	var v_0 *Constructor_Main_NumberCell = v_0_loop
	_ = v_0
	return ((1.0) / ((v_0).V0)) < (0.0)
}

func Call_Main_colorName(v_0_loop uint32) string {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var __t0 string
	{
		if v_0 == 1227005933 {
			__t0 = "R"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 939353081 {
			__t0 = "B"
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

func Call_Main_render(v_0_loop *Constructor_Main_Node) string {
render:
	for {
		if false {
			continue render
		}
		var v_0 *Constructor_Main_Node = v_0_loop
		_ = v_0
		var __t3 string
		{
			if v_0 == nil {
				__t3 = "E"
				goto end_branch_3
			} else {

			}
		}
		{
			if v_0 != nil {
				var __t2 string
				{
					var __t_tag_0 uint32 = (v_0).V0
					_ = __t_tag_0
					if uint32(__t_tag_0) == 1227005933 {
						__t2 = "R"
						goto end_branch_2
					} else {

					}
				}
				{
					var __t_tag_1 uint32 = (v_0).V0
					_ = __t_tag_1
					if uint32(__t_tag_1) == 939353081 {
						__t2 = "B"
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() string { panic("Failed pattern match") }()
				}
			end_branch_2:
				__t3 = Call_Main_nodeText(__t2, (v_0).V2, (v_0).V3, Call_Main_render((v_0).V1), Call_Main_render((v_0).V4))
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() string { panic("Failed pattern match") }()
		}
	end_branch_3:
		return __t3
	}
}

func Call_Main_expectTree(expected_0_loop string, actual_1_loop *Constructor_Main_Node) gopurs_runtime.Value {
	var expected_0 string = expected_0_loop
	_ = expected_0
	var actual_1 *Constructor_Main_Node = actual_1_loop
	_ = actual_1
	return Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
		actual   string
		expected string
		/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
	}{Call_Main_render(actual_1), expected_0})
}

func Call_Main_renderTagged(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t6 string
	{
		if v_0.Type == 9 && v_0.IntVal == 2444816644 {
			var __t2 string
			{
				var __t_tag_0 uint32 = (*Constructor_Main_First)(v_0.UnsafePtr).V0
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1227005933 {
					__t2 = "first:R:"
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = (*Constructor_Main_First)(v_0.UnsafePtr).V0
				_ = __t_tag_1
				if uint32(__t_tag_1) == 939353081 {
					__t2 = "first:B:"
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_2:
			__t6 = gopurs_runtime.ConcatString(__t2, Data_Show_ShowIntImpl((*Constructor_Main_First)(v_0.UnsafePtr).V1))
			goto end_branch_6
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 1617734670 {
			var __t5 string
			{
				var __t_tag_3 uint32 = (*Constructor_Main_Second)(v_0.UnsafePtr).V0
				_ = __t_tag_3
				if uint32(__t_tag_3) == 1227005933 {
					__t5 = "second:R:"
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = (*Constructor_Main_Second)(v_0.UnsafePtr).V0
				_ = __t_tag_4
				if uint32(__t_tag_4) == 939353081 {
					__t5 = "second:B:"
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_5:
			__t6 = gopurs_runtime.ConcatString(__t5, Data_Show_ShowIntImpl((*Constructor_Main_Second)(v_0.UnsafePtr).V1))
			goto end_branch_6
		} else {

		}
	}
	{
		__t6 = func() string { panic("Failed pattern match") }()
	}
end_branch_6:
	return __t6
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
