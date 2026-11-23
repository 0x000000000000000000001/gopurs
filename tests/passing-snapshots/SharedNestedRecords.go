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
		cache_Main_bind = Get_Control_Bind_bind__1264300116()
	})
	return cache_Main_bind
}

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = Get_Control_Bind_bind__2595690612()
	})
	return cache_Main_bind1
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__3833962260()
	})
	return cache_Main_bind2
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__1860431828()
	})
	return cache_Main_bind3
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__3040259796()
	})
	return cache_Main_bind4
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

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = Get_Test_Assert_assertEqual__1066131861()
	})
	return cache_Main_assertEqual2
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Bind_bind__4184743060()
	})
	return cache_Main_bind5
}

var cache_Main_bind6 gopurs_runtime.Value
var once_Main_bind6 sync.Once

func Get_Main_bind6() gopurs_runtime.Value {
	once_Main_bind6.Do(func() {
		cache_Main_bind6 = Get_Control_Bind_bind__4247668692()
	})
	return cache_Main_bind6
}

var cache_Main_negate gopurs_runtime.Value
var once_Main_negate sync.Once

func Get_Main_negate() gopurs_runtime.Value {
	once_Main_negate.Do(func() {
		cache_Main_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_negate(a_0_box.IntVal))
		})
	})
	return cache_Main_negate
}

var cache_Main_bind7 gopurs_runtime.Value
var once_Main_bind7 sync.Once

func Get_Main_bind7() gopurs_runtime.Value {
	once_Main_bind7.Do(func() {
		cache_Main_bind7 = Get_Control_Bind_bind__3692096660()
	})
	return cache_Main_bind7
}

var cache_Main_bind8 gopurs_runtime.Value
var once_Main_bind8 sync.Once

func Get_Main_bind8() gopurs_runtime.Value {
	once_Main_bind8.Do(func() {
		cache_Main_bind8 = Get_Control_Bind_bind__462331540()
	})
	return cache_Main_bind8
}

var cache_Main_bind9 gopurs_runtime.Value
var once_Main_bind9 sync.Once

func Get_Main_bind9() gopurs_runtime.Value {
	once_Main_bind9.Do(func() {
		cache_Main_bind9 = Get_Control_Bind_bind__1240467924()
	})
	return cache_Main_bind9
}

var cache_Main_bind10 gopurs_runtime.Value
var once_Main_bind10 sync.Once

func Get_Main_bind10() gopurs_runtime.Value {
	once_Main_bind10.Do(func() {
		cache_Main_bind10 = Get_Control_Bind_bind__1666388308()
	})
	return cache_Main_bind10
}

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

var cache_Main_bind11 gopurs_runtime.Value
var once_Main_bind11 sync.Once

func Get_Main_bind11() gopurs_runtime.Value {
	once_Main_bind11.Do(func() {
		cache_Main_bind11 = Get_Control_Bind_bind__3855431764()
	})
	return cache_Main_bind11
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
			// TAST (Let): original_3_3 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			original_3_3 := Call_Worker_makePacket(__local_var_2_2.IntVal)
			_ = original_3_3
			__local_var_4_4 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), original_3_3, gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Worker_advance(), gopurs_runtime.Int(int64(3))), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_5_5, gopurs_runtime.Value{})
			_ = __local_var_6_6
			// TAST (Let): changed_7_7 shape=App(Other) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			changed_7_7 := gopurs_runtime.Apply(__local_var_6_6, original_3_3)
			_ = changed_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_score(original_3_3), int64(35)}), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_score(Call_Worker_makePacket(int64(0))), int64(15)}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_score(changed_7_7), int64(49)}), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "column").IntVal, int64(5)}), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "right"), "score").IntVal, int64(12)}), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.RecordGet(original_3_3, "enabled").IntVal) != (0), true}), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(changed_7_7, "left"), "column").IntVal, int64(8)}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(changed_7_7, "right"), "score").IntVal, int64(24)}), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.RecordGet(changed_7_7, "enabled").IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(changed_7_7, "right"), "label").StrVal(), "keep"}), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
				return changed_7_7
			}), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_18_18, gopurs_runtime.Value{})
			_ = __local_var_19_19
			// TAST (Let): echoed_20_20 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			echoed_20_20 := gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Apply(__local_var_19_19, Get_Data_Unit_unit()))
			_ = echoed_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_score(echoed_20_20), int64(49)}), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "left"), "column").IntVal, int64(8)}), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "left"), "active").IntVal) != (0), true}), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
				actual   bool
				expected bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "right"), "active").IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_24_24
			var __t27 gopurs_runtime.Value
			{
				var __t_tag_26 gopurs_runtime.Value = gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "left"), "note")
				_ = __t_tag_26
				if __t_tag_26.Type == 9 && __t_tag_26.IntVal == 4153609677 {
					__t27 = Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
						actual   int64
						expected int64
						/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
					}{(*Constructor_Worker_Payload[gopurs_runtime.Value])(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(echoed_20_20, "left"), "note").UnsafePtr).V0.IntVal, int64(6)})
					goto end_branch_27
				} else {

				}
			}
			{
				__t27 = Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{false, true})
			}
		end_branch_27:
			__local_var_25_25 := gopurs_runtime.Apply(__t27, gopurs_runtime.Value{})
			_ = __local_var_25_25
			// TAST (Let): other_26_28 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			other_26_28 := Call_Worker_advance(int64(-2), original_3_3)
			_ = other_26_28
			__local_var_27_29 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_score(other_26_28), int64(39)}), gopurs_runtime.Value{})
			_ = __local_var_27_29
			__local_var_28_30 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_score(changed_7_7), int64(49)}), gopurs_runtime.Value{})
			_ = __local_var_28_30
			__local_var_29_31 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_4_4, gopurs_runtime.Value{})
			_ = __local_var_29_31
			__local_var_30_32 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_score(__local_var_29_31), int64(35)}), gopurs_runtime.Value{})
			_ = __local_var_30_32
			__local_var_31_33 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(other_26_28, "left"), "column").IntVal, int64(3)}), gopurs_runtime.Value{})
			_ = __local_var_31_33
			__local_var_32_34 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(7)), gopurs_runtime.Value{})
			_ = __local_var_32_34
			__local_var_33_35 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(v_33 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_0, gopurs_runtime.RecordGet(original_3_3, "enabled"), gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "active"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "column"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(original_3_3, "left"), "line"), gopurs_runtime.Value{Type: 9, IntVal: 4153609677, UnsafePtr: unsafe.Pointer((&Constructor_Worker_Payload[gopurs_runtime.Value]{1, __local_var_32_34}))}), gopurs_runtime.Value{Type: 9, IntVal: 4153609677, UnsafePtr: unsafe.Pointer((&Constructor_Worker_Payload[gopurs_runtime.Value]{1, __local_var_32_34}))}, gopurs_runtime.RecordGet(original_3_3, "right"))
			}), gopurs_runtime.Value{})
			_ = __local_var_33_35
			__local_var_34_36 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_33_35, gopurs_runtime.Value{})
			_ = __local_var_34_36
			// TAST (Let): returned_35_37 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [(ADT ["Effect","Ref","Ref"] [Int])]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [(ADT ["Effect","Ref","Ref"] [Int])])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean] Empty))
			returned_35_37 := gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Apply(__local_var_34_36, Get_Data_Unit_unit()))
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
				__t40 = Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{false, true})
			}
		end_branch_40:
			__local_var_36_38 := gopurs_runtime.Apply(__t40, gopurs_runtime.Value{})
			_ = __local_var_36_38
			var __t44 gopurs_runtime.Value
			{
				var __t_tag_42 gopurs_runtime.Value = gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(returned_35_37, "left"), "note")
				_ = __t_tag_42
				if __t_tag_42.Type == 9 && __t_tag_42.IntVal == 4153609677 {
					__t44 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_37_43 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), (*Constructor_Worker_Payload[gopurs_runtime.Value])(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(returned_35_37, "left"), "note").UnsafePtr).V0, gopurs_runtime.Value{})
						_ = __local_var_37_43
						return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
							actual   int64
							expected int64
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{__local_var_37_43.IntVal, int64(23)}), gopurs_runtime.Value{})
					})
					goto end_branch_44
				} else {

				}
			}
			{
				__t44 = Call_Test_Assert_assertEqual_prime___1094418358(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())), "", struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{false, true})
			}
		end_branch_44:
			__local_var_37_41 := gopurs_runtime.Apply(__t44, gopurs_runtime.Value{})
			_ = __local_var_37_41
			__local_var_38_45 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_32_34, gopurs_runtime.Value{})
			_ = __local_var_38_45
			__local_var_39_46 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{__local_var_38_45.IntVal, int64(23)}), gopurs_runtime.Value{})
			_ = __local_var_39_46
			__local_var_40_47 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{})
			_ = __local_var_40_47
			__local_var_41_48 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_40_47, gopurs_runtime.Value{})
			_ = __local_var_41_48
			// TAST (Let): roundtrip_42_49 shape=App(Var) bindingType=(Record (Row [payload: (ADT ["Worker","Payload"] [Int]), left: (Record (Row [column: Int, line: Int, active: Boolean, note: (ADT ["Worker","Payload"] [Int])] Empty)), right: (Record (Row [score: Int, label: String, active: Boolean] Empty)), enabled: Boolean, extra: String] Empty))
			roundtrip_42_49 := gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Apply(__local_var_41_48, gopurs_runtime.RecordShape5(&GopursRecordKeys_4d61696e_2, gopurs_runtime.RecordGet(original_3_3, "enabled"), gopurs_runtime.StrWithCodeUnits("retained", 8), gopurs_runtime.RecordGet(original_3_3, "left"), gopurs_runtime.RecordGet(original_3_3, "payload"), gopurs_runtime.RecordGet(original_3_3, "right"))))
			_ = roundtrip_42_49
			__local_var_43_50 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(roundtrip_42_49, "left"), "column").IntVal, int64(5)}), gopurs_runtime.Value{})
			_ = __local_var_43_50
			__local_var_44_51 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.RecordGet(roundtrip_42_49, "extra").StrVal(), "retained"}), gopurs_runtime.Value{})
			_ = __local_var_44_51
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
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

func Get_Main_echoPayload() gopurs_runtime.Value {
	return _Gopurs_Main_EchoPayload
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys4{K0: "enabled", K1: "left", K2: "payload", K3: "right"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys4{K0: "active", K1: "column", K2: "line", K3: "note"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys5{K0: "enabled", K1: "extra", K2: "left", K3: "payload", K4: "right"}
