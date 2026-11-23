package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			// TAST (Let): original_3_3 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			original_3_3 := Call_Worker_makePacket(seed_2_2.IntVal)
			_ = original_3_3
			originalRef_4_4 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), original_3_3, gopurs_runtime.Value{})
			_ = originalRef_4_4
			callbackRef_5_5 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Worker_advance(), gopurs_runtime.Int(int64(3))), gopurs_runtime.Value{})
			_ = callbackRef_5_5
			callback_6_6 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), callbackRef_5_5, gopurs_runtime.Value{})
			_ = callback_6_6
			// TAST (Let): changed_7_7 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			changed_7_7 := Call_Worker_invoke(callback_6_6, original_3_3)
			_ = changed_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_score(original_3_3), int64(35)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_score(Call_Worker_makePacket(int64(0))), int64(15)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_score(changed_7_7), int64(49)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "column").IntVal, int64(5)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "right"), "score").IntVal, int64(12)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.RecordGet(original_3_3, "enabled").IntVal) != (0), true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			_dollar___unused_14_14 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(changed_7_7, "left"), "column").IntVal, int64(8)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_14_14
			_dollar___unused_15_15 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(changed_7_7, "right"), "score").IntVal, int64(24)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_15_15
			_dollar___unused_16_16 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.RecordGet(changed_7_7, "enabled").IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_16_16
			_dollar___unused_17_17 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(changed_7_7, "right"), "label").StrVal(), "keep"}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_17_17
			savedRef_18_18 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Worker_keep(), changed_7_7), gopurs_runtime.Value{})
			_ = savedRef_18_18
			saved_19_19 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), savedRef_18_18, gopurs_runtime.Value{})
			_ = saved_19_19
			// TAST (Let): echoed_20_20 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			echoed_20_20 := gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Apply(saved_19_19, Get_Data_Unit_unit()))
			_ = echoed_20_20
			_dollar___unused_21_21 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_score(echoed_20_20), int64(49)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_21_21
			_dollar___unused_22_22 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_readColumn(echoed_20_20), int64(8)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_22_22
			_dollar___unused_23_23 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "left"), "active").IntVal) != (0), true}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_23_23
			_dollar___unused_24_24 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "right"), "active").IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_24_24
			var __t27 gopurs_runtime.Value
			{
				var __t_tag_26 gopurs_runtime.Value = gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "left"), "note")
				_ = __t_tag_26
				if __t_tag_26.Type == 9 && __t_tag_26.IntVal == 4153609677 {
					__t27 = gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
						orig := struct {
							actual   int64
							expected int64
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{(*Constructor_Worker_Payload[gopurs_runtime.Value])(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "left"), "note").UnsafePtr).V0.IntVal, int64(6)}
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
					}())
					goto end_branch_27
				} else {

				}
			}
			{
				__t27 = gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
					orig := struct {
						actual   bool
						expected bool
						/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
					}{false, true}
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
				}())
			}
		end_branch_27:
			_dollar___unused_25_25 := gopurs_runtime.Apply(__t27, gopurs_runtime.Value{})
			_ = _dollar___unused_25_25
			// TAST (Let): other_26_28 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			other_26_28 := Call_Worker_invoke(gopurs_runtime.Apply(Get_Worker_advance(), gopurs_runtime.Int(int64(-2))), original_3_3)
			_ = other_26_28
			_dollar___unused_27_29 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_score(other_26_28), int64(39)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_27_29
			_dollar___unused_28_30 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_score(changed_7_7), int64(49)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_28_30
			preserved_29_31 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), originalRef_4_4, gopurs_runtime.Value{})
			_ = preserved_29_31
			_dollar___unused_30_32 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_score(preserved_29_31), int64(35)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_30_32
			_dollar___unused_31_33 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(other_26_28, "left"), "column").IntVal, int64(3)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_31_33
			cell_32_34 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(7)), gopurs_runtime.Value{})
			_ = cell_32_34
			retainedRef_33_35 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Worker_keep(), gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_1, gopurs_runtime.RecordGet(original_3_3, "enabled"), gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "active"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "column"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "line"), gopurs_runtime.Value{Type: 9, IntVal: 4153609677, UnsafePtr: unsafe.Pointer((&Constructor_Worker_Payload[gopurs_runtime.Value]{1, cell_32_34}))}), gopurs_runtime.Value{Type: 9, IntVal: 4153609677, UnsafePtr: unsafe.Pointer((&Constructor_Worker_Payload[gopurs_runtime.Value]{1, cell_32_34}))}, gopurs_runtime.RecordGet(original_3_3, "right"))), gopurs_runtime.Value{})
			_ = retainedRef_33_35
			retained_34_36 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), retainedRef_33_35, gopurs_runtime.Value{})
			_ = retained_34_36
			// TAST (Let): returned_35_37 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [(ADT ["Effect","Ref","Ref"] [Int])]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [(ADT ["Effect","Ref","Ref"] [Int])])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			returned_35_37 := gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Apply(retained_34_36, Get_Data_Unit_unit()))
			_ = returned_35_37
			var __t40 gopurs_runtime.Value
			{
				var __t_tag_39 gopurs_runtime.Value = gopurs_runtime.RecordGet(returned_35_37, "payload")
				_ = __t_tag_39
				if __t_tag_39.Type == 9 && __t_tag_39.IntVal == 4153609677 {
					__t40 = gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(23)), (*Constructor_Worker_Payload[gopurs_runtime.Value])(gopurs_runtime.RecordGet(returned_35_37, "payload").UnsafePtr).V0)
					goto end_branch_40
				} else {

				}
			}
			{
				__t40 = gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
					orig := struct {
						actual   bool
						expected bool
						/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
					}{false, true}
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
				}())
			}
		end_branch_40:
			_dollar___unused_36_38 := gopurs_runtime.Apply(__t40, gopurs_runtime.Value{})
			_ = _dollar___unused_36_38
			var __t44 gopurs_runtime.Value
			{
				var __t_tag_42 gopurs_runtime.Value = gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(returned_35_37, "left"), "note")
				_ = __t_tag_42
				if __t_tag_42.Type == 9 && __t_tag_42.IntVal == 4153609677 {
					__t44 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						sharedValue_37_43 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), (*Constructor_Worker_Payload[gopurs_runtime.Value])(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(returned_35_37, "left"), "note").UnsafePtr).V0, gopurs_runtime.Value{})
						_ = sharedValue_37_43
						return gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
							orig := struct {
								actual   int64
								expected int64
								/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
							}{sharedValue_37_43.IntVal, int64(23)}
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
						}(), gopurs_runtime.Value{})
					})
					goto end_branch_44
				} else {

				}
			}
			{
				__t44 = gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
					orig := struct {
						actual   bool
						expected bool
						/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
					}{false, true}
					_ = orig
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
				}())
			}
		end_branch_44:
			_dollar___unused_37_41 := gopurs_runtime.Apply(__t44, gopurs_runtime.Value{})
			_ = _dollar___unused_37_41
			observed_38_45 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), cell_32_34, gopurs_runtime.Value{})
			_ = observed_38_45
			_dollar___unused_39_46 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{observed_38_45.IntVal, int64(23)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_39_46
			genericRef_40_47 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{})
			_ = genericRef_40_47
			generic_41_48 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), genericRef_40_47, gopurs_runtime.Value{})
			_ = generic_41_48
			// TAST (Let): roundtrip_42_49 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean, extra: String] Empty))
			roundtrip_42_49 := gopurs_runtime.Apply(Get_Main_echoPayload(), Call_Worker_invoke(generic_41_48, gopurs_runtime.RecordShape5(&GopursRecordKeys_4d61696e_3, gopurs_runtime.RecordGet(original_3_3, "enabled"), gopurs_runtime.StrWithCodeUnits("retained", 8), gopurs_runtime.RecordGet(original_3_3, "left"), gopurs_runtime.RecordGet(original_3_3, "payload"), gopurs_runtime.RecordGet(original_3_3, "right"))))
			_ = roundtrip_42_49
			_dollar___unused_43_50 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_readColumn(roundtrip_42_49), int64(5)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_43_50
			_dollar___unused_44_51 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.RecordGet(roundtrip_42_49, "extra").StrVal(), "retained"}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_44_51
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
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

func Get_Main_echoPayload() gopurs_runtime.Value {
	return _Gopurs_Main_EchoPayload
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys4{K0: "enabled", K1: "left", K2: "payload", K3: "right"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys4{K0: "active", K1: "column", K2: "line", K3: "note"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys5{K0: "enabled", K1: "extra", K2: "left", K3: "payload", K4: "right"}
