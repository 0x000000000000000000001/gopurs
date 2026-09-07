package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			// TAST (Let): __local_var_4_4 shape=App(Var) bindingType=Any
			__local_var_4_4 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					black   *Constructor_Main_Node
					empty   *Constructor_Main_Node
					flagged *Constructor_Main_Node
					other   *Constructor_Main_Node
					red     *Constructor_Main_Node
					seven   *Constructor_Main_Node
				}{(&Constructor_Main_Node{1, 939353081, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(17), false, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})}), (*Constructor_Main_Node)(nil), (&Constructor_Main_Node{1, 1227005933, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(17), true, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})}), (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(91), true, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(42), true, (*Constructor_Main_Node)(nil)})}), (&Constructor_Main_Node{1, 1227005933, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(17), false, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})}), (&Constructor_Main_Node{1, 1227005933, (&Constructor_Main_Node{1, 939353081, (*Constructor_Main_Node)(nil), int64(-3), true, (*Constructor_Main_Node)(nil)}), int64(7), false, (&Constructor_Main_Node{1, 1227005933, (*Constructor_Main_Node)(nil), int64(9), false, (*Constructor_Main_Node)(nil)})})}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"black", "empty", "flagged", "other", "red", "seven"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.black)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.empty)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.flagged)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.other)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.red)}, gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(orig.seven)}})
			}())
			_ = __local_var_4_4
			inputsRef_5_5 := gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Value{})
			_ = inputsRef_5_5
			inputs_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), inputsRef_5_5), gopurs_runtime.Value{})
			_ = inputs_6_6
			__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "red"))))), gopurs_runtime.Str(blackText_3_3))), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "black"))))), gopurs_runtime.Str(blackText_3_3))), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "empty"))))), gopurs_runtime.Str("E"))), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "red"))))), gopurs_runtime.Str(sevenText_1_1))), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "seven"))))), gopurs_runtime.Str(sevenText_1_1))), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "red"))))), gopurs_runtime.Str(trueText_2_2))), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "flagged"))))), gopurs_runtime.Str(trueText_2_2))), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "red")))), gopurs_runtime.Str(redText_0_0))), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "black")))), gopurs_runtime.Str(blackText_3_3))), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "empty"))))), gopurs_runtime.Str("E"))), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "empty"))))), gopurs_runtime.Str("E"))), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setBlack(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "red"))))}), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_18), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 4068147934, UnsafePtr: unsafe.Pointer(Call_Main_setSeven(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](__local_var_19_19)))}), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_20_20), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](__local_var_21_21))), gopurs_runtime.Str(Call_Main_nodeText("B", int64(7), false, "B:-3:true[E][E]", "R:9:false[E][E]")))), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](__local_var_19_19))), gopurs_runtime.Str(blackText_3_3))), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "red")))), gopurs_runtime.Str(redText_0_0))), gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_swapChildren(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "black"))))), gopurs_runtime.Str(Call_Main_nodeText("B", int64(17), false, "R:9:false[E][E]", "B:-3:true[E][E]")))), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_mixSources(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "black")), gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "other"))))), gopurs_runtime.Str(Call_Main_nodeText("B", int64(17), false, "B:-3:true[E][E]", "B:42:true[E][E]")))), gopurs_runtime.Value{})
			_ = __local_var_26_26
			__local_var_27_27 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "black")))), gopurs_runtime.Str(blackText_3_3))), gopurs_runtime.Value{})
			_ = __local_var_27_27
			__local_var_28_28 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(Call_Main_setBlackTrue(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "black"))))), gopurs_runtime.Str(Call_Main_nodeText("B", int64(17), true, "B:-3:true[E][E]", "R:9:false[E][E]")))), gopurs_runtime.Value{})
			_ = __local_var_28_28
			__local_var_29_29 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Node](gopurs_runtime.RecordGet(inputs_6_6, "black")))), gopurs_runtime.Str(blackText_3_3))), gopurs_runtime.Value{})
			_ = __local_var_29_29
			__local_var_30_30 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2444816644, UnsafePtr: unsafe.Pointer((&Constructor_Main_First{1, 939353081, int64(41)}))}), gopurs_runtime.Value{})
			_ = __local_var_30_30
			__local_var_31_31 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_30_30), gopurs_runtime.Value{})
			_ = __local_var_31_31
			__local_var_32_32 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_renderTagged(Call_Main_toSecond(__local_var_31_31))), gopurs_runtime.Str("second:B:41"))), gopurs_runtime.Value{})
			_ = __local_var_32_32
			__local_var_33_33 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_renderTagged(__local_var_31_31)), gopurs_runtime.Str("first:B:41"))), gopurs_runtime.Value{})
			_ = __local_var_33_33
			__local_var_34_34 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Float(0.0)), gopurs_runtime.Value{})
			_ = __local_var_34_34
			__local_var_35_35 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_34_34), gopurs_runtime.Value{})
			_ = __local_var_35_35
			__local_var_36_36 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 1758940859, UnsafePtr: unsafe.Pointer((&Constructor_Main_NumberCell{1, -(__local_var_35_35.FloatVal()), int64(5)}))}), gopurs_runtime.Value{})
			_ = __local_var_36_36
			__local_var_37_37 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_36_36), gopurs_runtime.Value{})
			_ = __local_var_37_37
			var __t_tag_39 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((1.0)/((*Constructor_Main_NumberCell)(__local_var_37_37.UnsafePtr).V0)), gopurs_runtime.Float(0.0))
			__local_var_38_38 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[bool]](Get_Data_Eq_eqBoolean())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[bool]](Get_Data_Show_showBoolean())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool((uint32(__t_tag_39.IntVal) == 1527465420)), gopurs_runtime.Bool(true))), gopurs_runtime.Value{})
			_ = __local_var_38_38
			__local_var_39_40 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 1758940859, UnsafePtr: unsafe.Pointer(Call_Main_positiveZero(gopurs_runtime.CoerceToStruct[Constructor_Main_NumberCell](__local_var_37_37)))}), gopurs_runtime.Value{})
			_ = __local_var_39_40
			__local_var_40_41 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_39_40), gopurs_runtime.Value{})
			_ = __local_var_40_41
			var __t_tag_43 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((1.0)/((*Constructor_Main_NumberCell)(__local_var_40_41.UnsafePtr).V0)), gopurs_runtime.Float(0.0))
			__local_var_41_42 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[bool]](Get_Data_Eq_eqBoolean())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[bool]](Get_Data_Show_showBoolean())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool((uint32(__t_tag_43.IntVal) == 1527465420)), gopurs_runtime.Bool(false))), gopurs_runtime.Value{})
			_ = __local_var_41_42
			var __t_tag_45 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((1.0)/((*Constructor_Main_NumberCell)(__local_var_37_37.UnsafePtr).V0)), gopurs_runtime.Float(0.0))
			__local_var_42_44 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[bool]](Get_Data_Eq_eqBoolean())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[bool]](Get_Data_Show_showBoolean())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool((uint32(__t_tag_45.IntVal) == 1527465420)), gopurs_runtime.Bool(true))), gopurs_runtime.Value{})
			_ = __local_var_42_44
			__local_var_43_46 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int((*Constructor_Main_NumberCell)(__local_var_40_41.UnsafePtr).V1), gopurs_runtime.Int(int64(5)))), gopurs_runtime.Value{})
			_ = __local_var_43_46
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
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
	return (((((((((color_0) + (":")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(key_1)).StrVal())) + (":")) + (__t0)) + ("[")) + (left_3)) + ("][")) + (right_4)) + ("]")
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
	var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((1.0)/((v_0).V0)), gopurs_runtime.Float(0.0))
	return (uint32(__t_tag_0.IntVal) == 1527465420)
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
					if uint32(__t_tag_0) == 1227005933 {
						__t2 = "R"
						goto end_branch_2
					} else {

					}
				}
				{
					var __t_tag_1 uint32 = (v_0).V0
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
	return gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(Call_Main_render(actual_1)), gopurs_runtime.Str(expected_0)))
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
				if uint32(__t_tag_0) == 1227005933 {
					__t2 = "first:R:"
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = (*Constructor_Main_First)(v_0.UnsafePtr).V0
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
			__t6 = (__t2) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Main_First)(v_0.UnsafePtr).V1)).StrVal())
			goto end_branch_6
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 1617734670 {
			var __t5 string
			{
				var __t_tag_3 uint32 = (*Constructor_Main_Second)(v_0.UnsafePtr).V0
				if uint32(__t_tag_3) == 1227005933 {
					__t5 = "second:R:"
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = (*Constructor_Main_Second)(v_0.UnsafePtr).V0
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
			__t6 = (__t5) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Main_Second)(v_0.UnsafePtr).V1)).StrVal())
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

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
