package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(5)))
			_ = __local_var_0_0
			seedRef_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = seedRef_1_1
			seed_2_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), seedRef_1_1, gopurs_runtime.Value{})
			_ = seed_2_2
			// TAST (Let): original_3_3 shape=Other bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			original_3_3 := (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, seed_2_2, gopurs_runtime.Int(int64(2))), gopurs_runtime.Value{Type: 9, IntVal: int64(2187435972), UnsafePtr: nil}}))}, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Int(int64(-4)), gopurs_runtime.Int(int64(7))), gopurs_runtime.Value{Type: 9, IntVal: int64(277954537), UnsafePtr: nil}}))}, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Int(int64(9)), gopurs_runtime.Int(int64(-3))), gopurs_runtime.Value{Type: 9, IntVal: int64(2859257932), UnsafePtr: nil}}))}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})})})
			_ = original_3_3
			originalRef_4_4 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(original_3_3)}, gopurs_runtime.Value{})
			_ = originalRef_4_4
			callbackRef_5_5 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Worker_shift(), gopurs_runtime.Int(int64(10))), gopurs_runtime.Value{})
			_ = callbackRef_5_5
			callback_6_6 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), callbackRef_5_5, gopurs_runtime.Value{})
			_ = callback_6_6
			// TAST (Let): changed_7_7 shape=App(Var) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			changed_7_7 := Call_Worker_invoke(callback_6_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(original_3_3)})
			_ = changed_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_summarize(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](changed_7_7)), "15:2:B;-4:7:M;9:-3:E;"}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_summarize(original_3_3), "5:2:B;-4:7:M;9:-3:E;"}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_summarize(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_Worker_invoke(callback_6_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))), ""}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			// TAST (Let): echoed_11_11 shape=App(Var) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			echoed_11_11 := gopurs_runtime.Apply(Get_Main_echoPayload(), changed_7_7)
			_ = echoed_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](echoed_11_11), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](changed_7_7)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.actual)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.expected)})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Main_echoPayload(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.actual)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.expected)})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			mappedRef_14_14 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Worker_mapStack(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.RecordUpdate1((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V0, "rowIndex", gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V0, "rowIndex").IntVal, int64(1)))), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_14.UnsafePtr).V1}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			})), gopurs_runtime.Value{})
			_ = mappedRef_14_14
			mapped_15_15 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), mappedRef_14_14, gopurs_runtime.Value{})
			_ = mapped_15_15
			// TAST (Let): moved_16_16 shape=App(Var) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int] Empty)), (ADT ["Worker","Phase"] [])])])
			moved_16_16 := Call_Worker_invoke(mapped_15_15, echoed_11_11)
			_ = moved_16_16
			_dollar___unused_17_17 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_summarize(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](moved_16_16)), "15:3:B;-4:8:M;9:-2:E;"}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_17_17
			savedRef_18_18 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Get_Worker_keep(), moved_16_16), gopurs_runtime.Value{})
			_ = savedRef_18_18
			saved_19_19 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), savedRef_18_18, gopurs_runtime.Value{})
			_ = saved_19_19
			_dollar___unused_20_20 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_4()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_5())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(saved_19_19, Get_Data_Unit_unit())), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](moved_16_16)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.actual)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.expected)})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_20_20
			_dollar___unused_21_21 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_summarize(Call_Worker_prepend(struct {
					byteOffset int64
					rowIndex   int64
					/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
				}{int64(-8), int64(12)}, 2859257932, original_3_3)), "-8:12:E;5:2:B;-4:7:M;9:-3:E;"}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_21_21
			_dollar___unused_22_22 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Worker_summarize(Call_Worker_prepend(struct {
					byteOffset int64
					rowIndex   int64
					/*gopurs:record-v1:2:0062007900740065004f00660066007300650074,0072006f00770049006e006400650078*/
				}{int64(0), int64(0)}, 277954537, original_3_3)), "0:0:M;5:2:B;-4:7:M;9:-3:E;"}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_22_22
			preserved_23_23 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), originalRef_4_4, gopurs_runtime.Value{})
			_ = preserved_23_23
			_dollar___unused_24_24 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_6()), gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_7())), func() gopurs_runtime.Value {
				orig := struct {
					actual   *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					expected *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](preserved_23_23), original_3_3}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.actual)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.expected)})
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_24_24
			cell_25_25 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(3)), gopurs_runtime.Value{})
			_ = cell_25_25
			widerRef_26_26 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_1737577620_849153993((&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
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
			}{int64(4), cell_25_25, "retained", int64(8)}, 2187435972}), (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
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
			}{int64(4), cell_25_25, "retained", int64(8)}, 2859257932}), (*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				byteOffset int64
				cell       gopurs_runtime.Value
				label      string
				rowIndex   int64
				/*gopurs:record-v1:4:0062007900740065004f00660066007300650074,00630065006c006c,006c006100620065006c,0072006f00770049006e006400650078*/
			}, uint32]])(nil)})})))}, gopurs_runtime.Value{})
			_ = widerRef_26_26
			wider_27_27 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), widerRef_26_26, gopurs_runtime.Value{})
			_ = wider_27_27
			genericRef_28_28 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{})
			_ = genericRef_28_28
			generic_29_29 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), genericRef_28_28, gopurs_runtime.Value{})
			_ = generic_29_29
			// TAST (Let): roundtrip_30_30 shape=App(Var) bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [byteOffset: Int, rowIndex: Int, cell: (ADT ["Effect","Ref","Ref"] [Int]), label: String] Empty)), (ADT ["Worker","Phase"] [])])])
			roundtrip_30_30 := gopurs_runtime.Apply(Get_Main_echoPayload(), Call_Worker_invoke(generic_29_29, wider_27_27))
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
						_dollar___unused_31_36 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(19)), gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V0.UnsafePtr).V0, "cell"), gopurs_runtime.Value{})
						_ = _dollar___unused_31_36
						observed_32_37 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V1).V0.UnsafePtr).V0, "cell"), gopurs_runtime.Value{})
						_ = observed_32_37
						_dollar___unused_33_38 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
							orig := struct {
								actual   int64
								expected int64
								/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
							}{observed_32_37.IntVal, int64(19)}
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
						}(), gopurs_runtime.Value{})
						_ = _dollar___unused_33_38
						_dollar___unused_34_39 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
							orig := struct {
								actual   string
								expected string
								/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
							}{gopurs_runtime.RecordGet((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(roundtrip_30_30.UnsafePtr).V0.UnsafePtr).V0, "label").StrVal(), "retained"}
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
						}(), gopurs_runtime.Value{})
						_ = _dollar___unused_34_39
						return gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
							orig := struct {
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
							}()), int64(12)}
							_ = orig
							return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
						}(), gopurs_runtime.Value{})
					})
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
					return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
				}())
			}
		end_branch_40:
			_dollar___unused_31_31 := gopurs_runtime.Apply(__t40, gopurs_runtime.Value{})
			_ = _dollar___unused_31_31
			originalCell_32_41 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), cell_25_25, gopurs_runtime.Value{})
			_ = originalCell_32_41
			_dollar___unused_33_42 := gopurs_runtime.Apply2(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{originalCell_32_41.IntVal, int64(19)}
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}(), gopurs_runtime.Value{})
			_ = _dollar___unused_33_42
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_List_Types_eqList(Call_Data_Tuple_eqTuple(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Worker_eqPhase()))))}))
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = Call_Data_List_Types_showList(gopurs_runtime.Apply(Call_Data_Tuple_showTuple(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Worker_showPhase()))))}))
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = Call_Data_List_Types_eqList(Call_Data_Tuple_eqTuple(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Worker_eqPhase()))))}))
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = Call_Data_List_Types_showList(gopurs_runtime.Apply(Call_Data_Tuple_showTuple(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Worker_showPhase()))))}))
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = Call_Data_List_Types_eqList(Call_Data_Tuple_eqTuple(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Worker_eqPhase()))))}))
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = Call_Data_List_Types_showList(gopurs_runtime.Apply(Call_Data_Tuple_showTuple(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Worker_showPhase()))))}))
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = Call_Data_List_Types_eqList(Call_Data_Tuple_eqTuple(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(gopurs_runtime.Apply(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Worker_eqPhase()))))}))
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = Call_Data_List_Types_showList(gopurs_runtime.Apply(Call_Data_Tuple_showTuple(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Apply2(Call_Data_Show_showRecordFieldsCons(Get_Main_byteOffsetIsSymbol()), gopurs_runtime.Apply(Call_Data_Show_showRecordFieldsConsNil(Get_Main_rowIndexIsSymbol()), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Worker_showPhase()))))}))
	})
	return cache_Main___cached_dict_7
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

func Rebox_Main_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1209612131(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[uint32])(unsafe.Pointer(in))
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
		return gopurs_runtime.RecordShape4(&GopursRecordKeys_4d61696e_3, gopurs_runtime.Int(orig.byteOffset), orig.cell, gopurs_runtime.Str(orig.label), gopurs_runtime.Int(orig.rowIndex))
	}()
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	return out
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

func Rebox_Main_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
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

func Rebox_Main_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[uint32])(unsafe.Pointer(in))
}

func Get_Main_echoPayload() gopurs_runtime.Value {
	return _Gopurs_Main_EchoPayload
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys1{K0: "reflectSymbol"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys2{K0: "byteOffset", K1: "rowIndex"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys2{K0: "actual", K1: "expected"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys4{K0: "byteOffset", K1: "cell", K2: "label", K3: "rowIndex"}
