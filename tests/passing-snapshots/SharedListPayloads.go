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

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__1631783380()
	})
	return cache_Main_bind2
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__3339484372()
	})
	return cache_Main_bind3
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__398126548()
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
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__1066131861()
	})
	return cache_Main_assertEqual
}

var cache_Main_rowIndexIsSymbol gopurs_runtime.Value
var once_Main_rowIndexIsSymbol sync.Once

func Get_Main_rowIndexIsSymbol() gopurs_runtime.Value {
	once_Main_rowIndexIsSymbol.Do(func() {
		cache_Main_rowIndexIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("rowIndex", 8)
		}))
	})
	return cache_Main_rowIndexIsSymbol
}

var cache_Main_byteOffsetIsSymbol gopurs_runtime.Value
var once_Main_byteOffsetIsSymbol sync.Once

func Get_Main_byteOffsetIsSymbol() gopurs_runtime.Value {
	once_Main_byteOffsetIsSymbol.Do(func() {
		cache_Main_byteOffsetIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("byteOffset", 10)
		}))
	})
	return cache_Main_byteOffsetIsSymbol
}

var cache_Main_assertEqual1 gopurs_runtime.Value
var once_Main_assertEqual1 sync.Once

func Get_Main_assertEqual1() gopurs_runtime.Value {
	once_Main_assertEqual1.Do(func() {
		cache_Main_assertEqual1 = Get_Test_Assert_assertEqual__1422072883()
	})
	return cache_Main_assertEqual1
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Bind_bind__1096139284()
	})
	return cache_Main_bind5
}

var cache_Main_bind6 gopurs_runtime.Value
var once_Main_bind6 sync.Once

func Get_Main_bind6() gopurs_runtime.Value {
	once_Main_bind6.Do(func() {
		cache_Main_bind6 = Get_Control_Bind_bind__3303654932()
	})
	return cache_Main_bind6
}

var cache_Main_bind7 gopurs_runtime.Value
var once_Main_bind7 sync.Once

func Get_Main_bind7() gopurs_runtime.Value {
	once_Main_bind7.Do(func() {
		cache_Main_bind7 = Get_Control_Bind_bind__3007475220()
	})
	return cache_Main_bind7
}

var cache_Main_bind8 gopurs_runtime.Value
var once_Main_bind8 sync.Once

func Get_Main_bind8() gopurs_runtime.Value {
	once_Main_bind8.Do(func() {
		cache_Main_bind8 = Get_Control_Bind_bind__1814472276()
	})
	return cache_Main_bind8
}

var cache_Main_bind9 gopurs_runtime.Value
var once_Main_bind9 sync.Once

func Get_Main_bind9() gopurs_runtime.Value {
	once_Main_bind9.Do(func() {
		cache_Main_bind9 = Get_Control_Bind_bind__611530900()
	})
	return cache_Main_bind9
}

var cache_Main_bind10 gopurs_runtime.Value
var once_Main_bind10 sync.Once

func Get_Main_bind10() gopurs_runtime.Value {
	once_Main_bind10.Do(func() {
		cache_Main_bind10 = Get_Control_Bind_bind__1358679252()
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
		cache_Main_bind11 = Get_Control_Bind_bind__3505160468()
	})
	return cache_Main_bind11
}

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = Get_Test_Assert_assertEqual__2867988789()
	})
	return cache_Main_assertEqual2
}

var cache_Main_assertEqual3 gopurs_runtime.Value
var once_Main_assertEqual3 sync.Once

func Get_Main_assertEqual3() gopurs_runtime.Value {
	once_Main_assertEqual3.Do(func() {
		cache_Main_assertEqual3 = Get_Test_Assert_assertEqual__2127306453()
	})
	return cache_Main_assertEqual3
}

var cache_Main_coordinateSum gopurs_runtime.Value
var once_Main_coordinateSum sync.Once

func Get_Main_coordinateSum() gopurs_runtime.Value {
	once_Main_coordinateSum.Do(func() {
		cache_Main_coordinateSum = gopurs_runtime.Func(func(row_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_coordinateSum(func() struct {
				byteOffset int64
				rowIndex   int64
				/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
			} {
				orig := row_0_box
				_ = orig
				clone := struct {
					byteOffset int64
					rowIndex   int64
					/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
				}{}
				clone.byteOffset = gopurs_runtime.RecordGet(orig, "byteOffset").IntVal
				clone.rowIndex = gopurs_runtime.RecordGet(orig, "rowIndex").IntVal
				return clone
			}()))
		})
	})
	return cache_Main_coordinateSum
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
			// TAST (Let): original_3_3 shape=Other bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			original_3_3 := (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, __local_var_2_2, gopurs_runtime.Int(int64(2))), gopurs_runtime.Value{Type: 9, IntVal: int64(2187435972), UnsafePtr: nil}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}(), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Int(int64(-4)), gopurs_runtime.Int(int64(7))), gopurs_runtime.Value{Type: 9, IntVal: int64(277954537), UnsafePtr: nil}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}(), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Int(int64(9)), gopurs_runtime.Int(int64(-3))), gopurs_runtime.Value{Type: 9, IntVal: int64(2859257932), UnsafePtr: nil}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}(), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})})})
			_ = original_3_3
			__local_var_4_4 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(original_3_3)}, gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Worker_shift(), gopurs_runtime.Int(int64(10))), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_5_5, gopurs_runtime.Value{})
			_ = __local_var_6_6
			// TAST (Let): changed_7_7 shape=App(Other) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			changed_7_7 := gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(original_3_3)})
			_ = changed_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_summarize(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](changed_7_7)), "15:2:B;-4:7:M;9:-3:E;"}), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_summarize(original_3_3), "5:2:B;-4:7:M;9:-3:E;"}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_summarize(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))), ""}), gopurs_runtime.Value{})
			_ = __local_var_10_10
			// TAST (Let): echoed_11_11 shape=App(Var) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			echoed_11_11 := gopurs_runtime.Apply(Get_Main_echoPayload(), changed_7_7)
			_ = echoed_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__1422072883(struct {
				actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](echoed_11_11), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](changed_7_7)}), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__1422072883(struct {
				actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})}), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Worker_mapStack(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.RecordUpdate1((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V0, "rowIndex", gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V0, "rowIndex").IntVal, int64(1)))), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V1}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			})), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_14_14, gopurs_runtime.Value{})
			_ = __local_var_15_15
			// TAST (Let): moved_16_16 shape=App(Other) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			moved_16_16 := gopurs_runtime.Apply(__local_var_15_15, echoed_11_11)
			_ = moved_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_summarize(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](moved_16_16)), "15:3:B;-4:8:M;9:-2:E;"}), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
				return moved_16_16
			}), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_18_18, gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__1422072883(struct {
				actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_19_19, Get_Data_Unit_unit())), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](moved_16_16)}), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_summarize(Call_Worker_prepend(struct {
				byteOffset int64
				rowIndex   int64
				/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
			}{int64(-8), int64(12)}, 2859257932, original_3_3)), "-8:12:E;5:2:B;-4:7:M;9:-3:E;"}), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Worker_summarize(Call_Worker_prepend(struct {
				byteOffset int64
				rowIndex   int64
				/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
			}{int64(0), int64(0)}, 277954537, original_3_3)), "0:0:M;5:2:B;-4:7:M;9:-3:E;"}), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_4_4, gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__1422072883(struct {
				actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_23_23), original_3_3}), gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(3)), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_1737577620_849153993((&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}, uint32]]{1, (&Constructor_Data_Tuple_Tuple[struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}, uint32]{1, struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}{int64(4), __local_var_25_25, "retained", int64(8)}, 2187435972}), (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}, uint32]]{1, (&Constructor_Data_Tuple_Tuple[struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}, uint32]{1, struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}{int64(4), __local_var_25_25, "retained", int64(8)}, 2859257932}), (*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}, uint32]])(nil)})})))}, gopurs_runtime.Value{})
			_ = __local_var_26_26
			__local_var_27_27 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_26_26, gopurs_runtime.Value{})
			_ = __local_var_27_27
			__local_var_28_28 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{})
			_ = __local_var_28_28
			__local_var_29_29 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_28_28, gopurs_runtime.Value{})
			_ = __local_var_29_29
			// TAST (Let): roundtrip_30_30 shape=App(Var) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int, cell: (ADT ["Effect","Ref","Ref"] [Int]), label: String] Empty)), (ADT ["Worker","Phase"] [])])])
			roundtrip_30_30 := gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Apply(__local_var_29_29, __local_var_27_27))
			_ = roundtrip_30_30
			var __t40 gopurs_runtime.Value
			{
				var __t_and_35 bool = false
				if roundtrip_30_30.Type == 9 && roundtrip_30_30.IntVal == 1358893437 && roundtrip_30_30.UnsafePtr != nil {

					var __t_tag_32 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V1
					_ = __t_tag_32
					var __t_and_34 bool = false
					if __t_tag_32 != nil {

						var __t_tag_33 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V1).V1
						_ = __t_tag_33
						__t_and_34 = (__t_tag_33 == nil)
					}
					__t_and_35 = __t_and_34
				}
				if __t_and_35 {
					__t40 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_31_36 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(19)), gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V0.UnsafePtr).V0, "cell"), gopurs_runtime.Value{})
						_ = __local_var_31_36
						__local_var_32_37 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V1).V0.UnsafePtr).V0, "cell"), gopurs_runtime.Value{})
						_ = __local_var_32_37
						__local_var_33_38 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
							actual   int64
							expected int64
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{__local_var_32_37.IntVal, int64(19)}), gopurs_runtime.Value{})
						_ = __local_var_33_38
						__local_var_34_39 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___895398247(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())), "", struct {
							actual   string
							expected string
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V0.UnsafePtr).V0, "label").StrVal(), "retained"}), gopurs_runtime.Value{})
						_ = __local_var_34_39
						return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
							actual   int64
							expected int64
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{Call_Main_coordinateSum(func() struct {
							byteOffset int64
							rowIndex   int64
							/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
						} {
							orig := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V1).V0.UnsafePtr).V0
							_ = orig
							clone := struct {
								byteOffset int64
								rowIndex   int64
								/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
							}{}
							clone.byteOffset = gopurs_runtime.RecordGet(orig, "byteOffset").IntVal
							clone.rowIndex = gopurs_runtime.RecordGet(orig, "rowIndex").IntVal
							return clone
						}()), int64(12)}), gopurs_runtime.Value{})
					})
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
			__local_var_31_31 := gopurs_runtime.Apply(__t40, gopurs_runtime.Value{})
			_ = __local_var_31_31
			__local_var_32_41 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_25_25, gopurs_runtime.Value{})
			_ = __local_var_32_41
			__local_var_33_42 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{__local_var_32_41.IntVal, int64(19)}), gopurs_runtime.Value{})
			_ = __local_var_33_42
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

func Call_Main_coordinateSum(row_0_loop struct {
	byteOffset int64
	rowIndex   int64
	/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
}) int64 {
	var row_0 struct {
		byteOffset int64
		rowIndex   int64
		/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
	} = row_0_loop
	_ = row_0
	return gopurs_runtime.IntAdd(row_0.byteOffset, row_0.rowIndex)
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

func Rebox_Main_1737577620_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
	byteOffset int64
	cell       gopurs_runtime.Value
	label      string
	rowIndex   int64
	/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
}, uint32]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_2325946424_138441832(in.V0))}
	out.V1 = Rebox_Main_1737577620_849153993(in.V1)
	return out
}

func Rebox_Main_2325946424_138441832(in *Constructor_Data_Tuple_Tuple[struct {
	byteOffset int64
	cell       gopurs_runtime.Value
	label      string
	rowIndex   int64
	/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
}, uint32]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.byteOffset), orig.cell, gopurs_runtime.Str(orig.label), gopurs_runtime.Int(orig.rowIndex))
	}()
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	return out
}

func Get_Main_echoPayload() gopurs_runtime.Value {
	return _Gopurs_Main_EchoPayload
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys1{K0: "reflectSymbol"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys2{K0: "byteOffset", K1: "rowIndex"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys4{K0: "byteOffset", K1: "cell", K2: "label", K3: "rowIndex"}
