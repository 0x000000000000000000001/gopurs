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

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Apply2(Get_Data_Show_showRecord(), gopurs_runtime.Value{}, gopurs_runtime.Value{})
	})
	return cache_Main_showRecord
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

var cache_Main_showRecordFieldsCons gopurs_runtime.Value
var once_Main_showRecordFieldsCons sync.Once

func Get_Main_showRecordFieldsCons() gopurs_runtime.Value {
	once_Main_showRecordFieldsCons.Do(func() {
		cache_Main_showRecordFieldsCons = Call_Data_Show_showRecordFieldsCons(Get_Main_columnIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_lineIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main_showRecordFieldsCons
}

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Get_Data_Show_show__1618093805()
	})
	return cache_Main_show
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main_showArray
}

var cache_Main_show1 gopurs_runtime.Value
var once_Main_show1 sync.Once

func Get_Main_show1() gopurs_runtime.Value {
	once_Main_show1.Do(func() {
		cache_Main_show1 = Get_Data_Show_show__4045533208()
	})
	return cache_Main_show1
}

var cache_Main_show2 gopurs_runtime.Value
var once_Main_show2 sync.Once

func Get_Main_show2() gopurs_runtime.Value {
	once_Main_show2.Do(func() {
		cache_Main_show2 = Get_Data_Show_show__3343036312()
	})
	return cache_Main_show2
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

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Main_add
}

var cache_Main_mul gopurs_runtime.Value
var once_Main_mul sync.Once

func Get_Main_mul() gopurs_runtime.Value {
	once_Main_mul.Do(func() {
		cache_Main_mul = Get_Data_Semiring_mul__4086903470()
	})
	return cache_Main_mul
}

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = Get_Control_Bind_bind__2828832660()
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
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__1338237930()
	})
	return cache_Main_assertEqual
}

var cache_Main_negate1 gopurs_runtime.Value
var once_Main_negate1 sync.Once

func Get_Main_negate1() gopurs_runtime.Value {
	once_Main_negate1.Do(func() {
		cache_Main_negate1 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_negate1(a_0_box.IntVal))
		})
	})
	return cache_Main_negate1
}

var cache_Main_assertEqual1 gopurs_runtime.Value
var once_Main_assertEqual1 sync.Once

func Get_Main_assertEqual1() gopurs_runtime.Value {
	once_Main_assertEqual1.Do(func() {
		cache_Main_assertEqual1 = Get_Test_Assert_assertEqual__674562613()
	})
	return cache_Main_assertEqual1
}

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = Get_Test_Assert_assertEqual__2813600053()
	})
	return cache_Main_assertEqual2
}

var cache_Main_assertEqual3 gopurs_runtime.Value
var once_Main_assertEqual3 sync.Once

func Get_Main_assertEqual3() gopurs_runtime.Value {
	once_Main_assertEqual3.Do(func() {
		cache_Main_assertEqual3 = Get_Test_Assert_assertEqual__3136838421()
	})
	return cache_Main_assertEqual3
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main_eqArray
}

var cache_Main_showArray1 gopurs_runtime.Value
var once_Main_showArray1 sync.Once

func Get_Main_showArray1() gopurs_runtime.Value {
	once_Main_showArray1.Do(func() {
		cache_Main_showArray1 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main_showArray1
}

var cache_Main_assertEqual4 gopurs_runtime.Value
var once_Main_assertEqual4 sync.Once

func Get_Main_assertEqual4() gopurs_runtime.Value {
	once_Main_assertEqual4.Do(func() {
		cache_Main_assertEqual4 = Get_Test_Assert_assertEqual__3154990355()
	})
	return cache_Main_assertEqual4
}

var cache_Main_assertEqual5 gopurs_runtime.Value
var once_Main_assertEqual5 sync.Once

func Get_Main_assertEqual5() gopurs_runtime.Value {
	once_Main_assertEqual5.Do(func() {
		cache_Main_assertEqual5 = Get_Test_Assert_assertEqual__2313317602()
	})
	return cache_Main_assertEqual5
}

var cache_Main_go__map gopurs_runtime.Value
var once_Main_go__map sync.Once

func Get_Main_go__map() gopurs_runtime.Value {
	once_Main_go__map.Do(func() {
		cache_Main_go__map = Get_Data_Functor_map__2485353908()
	})
	return cache_Main_go__map
}

var cache_Main_sub gopurs_runtime.Value
var once_Main_sub sync.Once

func Get_Main_sub() gopurs_runtime.Value {
	once_Main_sub.Do(func() {
		cache_Main_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Main_sub
}

var cache_Main_eqRec gopurs_runtime.Value
var once_Main_eqRec sync.Once

func Get_Main_eqRec() gopurs_runtime.Value {
	once_Main_eqRec.Do(func() {
		cache_Main_eqRec = gopurs_runtime.Apply(Get_Data_Eq_eqRec(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRec
}

var cache_Main_eqRowCons gopurs_runtime.Value
var once_Main_eqRowCons sync.Once

func Get_Main_eqRowCons() gopurs_runtime.Value {
	once_Main_eqRowCons.Do(func() {
		cache_Main_eqRowCons = gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRowCons
}

var cache_Main_valuesIsSymbol gopurs_runtime.Value
var once_Main_valuesIsSymbol sync.Once

func Get_Main_valuesIsSymbol() gopurs_runtime.Value {
	once_Main_valuesIsSymbol.Do(func() {
		cache_Main_valuesIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("values", 6)
		}))
	})
	return cache_Main_valuesIsSymbol
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

var cache_Main_labelIsSymbol gopurs_runtime.Value
var once_Main_labelIsSymbol sync.Once

func Get_Main_labelIsSymbol() gopurs_runtime.Value {
	once_Main_labelIsSymbol.Do(func() {
		cache_Main_labelIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("label", 5)
		}))
	})
	return cache_Main_labelIsSymbol
}

var cache_Main_assertEqual6 gopurs_runtime.Value
var once_Main_assertEqual6 sync.Once

func Get_Main_assertEqual6() gopurs_runtime.Value {
	once_Main_assertEqual6.Do(func() {
		cache_Main_assertEqual6 = Get_Test_Assert_assertEqual__3636626607()
	})
	return cache_Main_assertEqual6
}

var cache_Main_map1 gopurs_runtime.Value
var once_Main_map1 sync.Once

func Get_Main_map1() gopurs_runtime.Value {
	once_Main_map1.Do(func() {
		cache_Main_map1 = Get_Data_Functor_map__859611156()
	})
	return cache_Main_map1
}

var cache_Main_assertEqual7 gopurs_runtime.Value
var once_Main_assertEqual7 sync.Once

func Get_Main_assertEqual7() gopurs_runtime.Value {
	once_Main_assertEqual7.Do(func() {
		cache_Main_assertEqual7 = Get_Test_Assert_assertEqual__2867988789()
	})
	return cache_Main_assertEqual7
}

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = Get_Control_Bind_bind__2708175764()
	})
	return cache_Main_bind1
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Bind_bind__731960532()
	})
	return cache_Main_bind2
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Bind_bind__4105335956()
	})
	return cache_Main_bind3
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Bind_bind__4082872532()
	})
	return cache_Main_bind4
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Bind_bind__1088188436()
	})
	return cache_Main_bind5
}

var cache_Main_bind6 gopurs_runtime.Value
var once_Main_bind6 sync.Once

func Get_Main_bind6() gopurs_runtime.Value {
	once_Main_bind6.Do(func() {
		cache_Main_bind6 = Get_Control_Bind_bind__3874655604()
	})
	return cache_Main_bind6
}

var cache_Main_bind7 gopurs_runtime.Value
var once_Main_bind7 sync.Once

func Get_Main_bind7() gopurs_runtime.Value {
	once_Main_bind7.Do(func() {
		cache_Main_bind7 = Get_Control_Bind_bind__1264300116()
	})
	return cache_Main_bind7
}

var cache_Main_bind8 gopurs_runtime.Value
var once_Main_bind8 sync.Once

func Get_Main_bind8() gopurs_runtime.Value {
	once_Main_bind8.Do(func() {
		cache_Main_bind8 = Get_Control_Bind_bind__3443383764()
	})
	return cache_Main_bind8
}

var cache_Main_bind9 gopurs_runtime.Value
var once_Main_bind9 sync.Once

func Get_Main_bind9() gopurs_runtime.Value {
	once_Main_bind9.Do(func() {
		cache_Main_bind9 = Get_Control_Bind_bind__3089214420()
	})
	return cache_Main_bind9
}

var cache_Main_bind10 gopurs_runtime.Value
var once_Main_bind10 sync.Once

func Get_Main_bind10() gopurs_runtime.Value {
	once_Main_bind10.Do(func() {
		cache_Main_bind10 = Get_Control_Bind_bind__2595690612()
	})
	return cache_Main_bind10
}

var cache_Main_Empty gopurs_runtime.Value
var once_Main_Empty sync.Once

func Get_Main_Empty() gopurs_runtime.Value {
	once_Main_Empty.Do(func() {
		cache_Main_Empty = gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((*Constructor_Main_Item)(nil))}
	})
	return cache_Main_Empty
}

var cache_Main_Item gopurs_runtime.Value
var once_Main_Item sync.Once

func Get_Main_Item() gopurs_runtime.Value {
	once_Main_Item.Do(func() {
		cache_Main_Item = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((&Constructor_Main_Item{1, func() struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				} {
					orig := value0
					_ = orig
					clone := struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}{}
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
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
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}(), Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](value1))}))}
			})
		})
	})
	return cache_Main_Item
}

var cache_Main_describe gopurs_runtime.Value
var once_Main_describe sync.Once

func Get_Main_describe() gopurs_runtime.Value {
	once_Main_describe.Do(func() {
		cache_Main_describe = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_describe(gopurs_runtime.CoerceToStruct[Constructor_Main_Item](v_0_box)))
		})
	})
	return cache_Main_describe
}

var cache_Main_applyAppend gopurs_runtime.Value
var once_Main_applyAppend sync.Once

func Get_Main_applyAppend() gopurs_runtime.Value {
	once_Main_applyAppend.Do(func() {
		cache_Main_applyAppend = gopurs_runtime.Func3(func(append1_0_box gopurs_runtime.Value, values_1_box gopurs_runtime.Value, value_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_applyAppend(append1_0_box, (*((*[]gopurs_runtime.Value)(values_1_box.UnsafePtr))), value_2_box))
		})
	})
	return cache_Main_applyAppend
}

var cache_Main_applyAppend__3199238329 gopurs_runtime.Value
var once_Main_applyAppend__3199238329 sync.Once

func Get_Main_applyAppend__3199238329() gopurs_runtime.Value {
	once_Main_applyAppend__3199238329.Do(func() {
		cache_Main_applyAppend__3199238329 = gopurs_runtime.Func3(func(append1_0_box gopurs_runtime.Value, values_1_box gopurs_runtime.Value, value_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_applyAppend__3199238329(append1_0_box, func() []string {
					arr := *(*[]gopurs_runtime.Value)(values_1_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}(), value_2_box.StrVal())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_applyAppend__3199238329
}

var cache_Main_appendItem gopurs_runtime.Value
var once_Main_appendItem sync.Once

func Get_Main_appendItem() gopurs_runtime.Value {
	once_Main_appendItem.Do(func() {
		cache_Main_appendItem = Get_Data_Array_snoc()
	})
	return cache_Main_appendItem
}

var cache_Main_appendInt gopurs_runtime.Value
var once_Main_appendInt sync.Once

func Get_Main_appendInt() gopurs_runtime.Value {
	once_Main_appendInt.Do(func() {
		cache_Main_appendInt = Get_Data_Array_snoc()
	})
	return cache_Main_appendInt
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					arrays   [][]int64
					booleans []bool
					entry    struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}
					functions []gopurs_runtime.Value
					ints      []int64
					lists     []*Constructor_Data_List_Types_Cons[int64]
					numbers   []float64
					strings   []string
					/*gopurs:record-v1:8:006100720072006100790073,0062006f006f006c00650061006e0073,0065006e007400720079,00660075006e006300740069006f006e0073,0069006e00740073,006c0069007300740073,006e0075006d0062006500720073,0073007400720069006e00670073*/
				}{[][]int64{[]int64{}, []int64{int64(1), int64(2)}}, []bool{true, false}, struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}{"kept", struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}{int64(7), int64(11)}, []int64{int64(13), int64(17)}}, (*((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(n_0.IntVal, int64(1)))
				}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntMul(n_0.IntVal, int64(2)))
				})}).UnsafePtr))), []int64{int64(-2147483648), int64(0), int64(2147483647)}, func() []*Constructor_Data_List_Types_Cons[int64] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, int64(3), (&Constructor_Data_List_Types_Cons[int64]{1, int64(5), (*Constructor_Data_List_Types_Cons[int64])(nil)})})))}}).UnsafePtr)
					unboxed := make([]*Constructor_Data_List_Types_Cons[int64], len(arr))
					for i, v := range arr {
						unboxed[i] = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v))
					}
					return unboxed
				}(), []float64{1.25, -2.5}, []string{"", "λ😀"}}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"arrays", "booleans", "entry", "functions", "ints", "lists", "numbers", "strings"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					arr := orig.arrays
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							arr := v
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.booleans
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Bool(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					orig := orig.entry
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), gopurs_runtime.Array(orig.functions), func() gopurs_runtime.Value {
					arr := orig.ints
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.lists
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993(v))}
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.numbers
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Float(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.strings
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()})
			}())
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_1_1, gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_0())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), gopurs_runtime.RecordGet(__local_var_2_2, "ints"), gopurs_runtime.Int(int64(-7))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(-7)}}), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), func() gopurs_runtime.Value {
					arr := []int64{}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(int64(42))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(42)}}), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3842687442(Rebox_Main_1386611502_3467499472(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_2())), "", struct {
				actual   []float64
				expected []float64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Data_Array_snoc__2031673812(func() []float64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_2_2, "numbers").UnsafePtr)
				unboxed := make([]float64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.FloatVal()
				}
				return unboxed
			}(), 3.75), []float64{1.25, -2.5, 3.75}}), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1268761764(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())), "", struct {
				actual   []string
				expected []string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Data_Array_snoc__2342910242(func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_2_2, "strings").UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}(), "last"), []string{"", "λ😀", "last"}}), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1282197621(Rebox_Main_1386611502_38358956(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_4())), "", struct {
				actual   []bool
				expected []bool
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Data_Array_snoc__1829638003(func() []bool {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_2_2, "booleans").UnsafePtr)
				unboxed := make([]bool, len(arr))
				for i, v := range arr {
					unboxed[i] = (v.IntVal) != (0)
				}
				return unboxed
			}(), true), []bool{true, false, true}}), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___165009865(Rebox_Main_1386611502_940452149(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_5())), "", struct {
				actual   [][]int64
				expected [][]int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Data_Array_snoc__3649272514(func() [][]int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_2_2, "arrays").UnsafePtr)
				unboxed := make([][]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
				}
				return unboxed
			}(), []int64{int64(3)}), [][]int64{[]int64{}, []int64{int64(1), int64(2)}, []int64{int64(3)}}}), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___680644777(Rebox_Main_1386611502_927154015(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_6())), "", struct {
				actual   []*Constructor_Data_List_Types_Cons[int64]
				expected []*Constructor_Data_List_Types_Cons[int64]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Data_Array_snoc__2640853809(func() []*Constructor_Data_List_Types_Cons[int64] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_2_2, "lists").UnsafePtr)
				unboxed := make([]*Constructor_Data_List_Types_Cons[int64], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v))
				}
				return unboxed
			}(), (&Constructor_Data_List_Types_Cons[int64]{1, int64(7), (*Constructor_Data_List_Types_Cons[int64])(nil)})), func() []*Constructor_Data_List_Types_Cons[int64] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, int64(3), (&Constructor_Data_List_Types_Cons[int64]{1, int64(5), (*Constructor_Data_List_Types_Cons[int64])(nil)})})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, int64(7), (*Constructor_Data_List_Types_Cons[int64])(nil)})))}}).UnsafePtr)
				unboxed := make([]*Constructor_Data_List_Types_Cons[int64], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Main_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v))
				}
				return unboxed
			}()}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			arrayIntrinsic_arg_11 := func(f_10 gopurs_runtime.Value) int64 {
				return gopurs_runtime.Apply(f_10, gopurs_runtime.Int(int64(10))).IntVal
			}
			_ = arrayIntrinsic_arg_11
			_ = arrayIntrinsic_arg_11
			arrayIntrinsic_arg_12 := gopurs_runtime.Array(Call_Data_Array_snoc__1029666857((*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_2_2, "functions").UnsafePtr))), gopurs_runtime.Func(func(n_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntSub(n_10.IntVal, int64(3)))
			})))
			_ = arrayIntrinsic_arg_12
			_ = arrayIntrinsic_arg_12
			__local_var_10_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_7())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr_val_arrayMap13 := arrayIntrinsic_arg_12
				_ = arr_val_arrayMap13
				_ = arr_val_arrayMap13
				arr_go_arrayMap13 := (*[]gopurs_runtime.Value)(arr_val_arrayMap13.UnsafePtr)
				_ = arr_go_arrayMap13
				res_go_arrayMap13 := make([]int64, len(*arr_go_arrayMap13))
				_ = res_go_arrayMap13
				for i_arrayMap13, v_arrayMap13 := range *arr_go_arrayMap13 {
					res_go_arrayMap13[i_arrayMap13] = arrayIntrinsic_arg_11(v_arrayMap13)
				}
				return res_go_arrayMap13
			}(), []int64{int64(11), int64(20), int64(7)}}), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__3636626607(struct {
				actual []struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}
				expected []struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Data_Array_snoc__3664849448([]struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			}{func() struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			} {
				orig := gopurs_runtime.RecordGet(__local_var_2_2, "entry")
				_ = orig
				clone := struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}{}
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
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
				clone.values = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}()}, struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			}{"new", struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}{int64(19), int64(23)}, []int64{int64(29)}}), func() []struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			} {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(__local_var_2_2, "entry"), func() gopurs_runtime.Value {
					orig := struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}{"new", struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}{int64(19), int64(23)}, []int64{int64(29)}}
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_1, gopurs_runtime.Str(orig.label), func() gopurs_runtime.Value {
						orig := orig.position
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
					}(), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}()}).UnsafePtr)
				unboxed := make([]struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}, len(arr))
				for i, v := range arr {
					unboxed[i] = func() struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					} {
						orig := v
						_ = orig
						clone := struct {
							label    string
							position struct {
								column int64
								line   int64
								/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
							}
							values []int64
							/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
						}{}
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
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
						clone.values = func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
							unboxed := make([]int64, len(arr))
							for i, v := range arr {
								unboxed[i] = v.IntVal
							}
							return unboxed
						}()
						return clone
					}()
				}
				return unboxed
			}()}), gopurs_runtime.Value{})
			_ = __local_var_11_13
			arrayIntrinsic_arg_15 := Get_Main_describe()
			_ = arrayIntrinsic_arg_15
			_ = arrayIntrinsic_arg_15
			arrayIntrinsic_arg_16 := gopurs_runtime.Apply2(Get_Main_appendItem(), func() gopurs_runtime.Value {
				arr := []*Constructor_Main_Item{(*Constructor_Main_Item)(nil), (&Constructor_Main_Item{1, func() struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				} {
					orig := gopurs_runtime.RecordGet(__local_var_2_2, "entry")
					_ = orig
					clone := struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}{}
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
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
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}(), (&Constructor_Data_List_Types_Cons[int64]{1, int64(31), (*Constructor_Data_List_Types_Cons[int64])(nil)})})}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer(v)}
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((&Constructor_Main_Item{1, struct {
				label    string
				position struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}
				values []int64
				/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
			}{"new", struct {
				column int64
				line   int64
				/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
			}{int64(19), int64(23)}, []int64{int64(29)}}, (*Constructor_Data_List_Types_Cons[int64])(nil)}))})
			_ = arrayIntrinsic_arg_16
			_ = arrayIntrinsic_arg_16
			arrayIntrinsic_arg_17 := Get_Main_describe()
			_ = arrayIntrinsic_arg_17
			_ = arrayIntrinsic_arg_17
			arrayIntrinsic_arg_18 := func() gopurs_runtime.Value {
				arr := []*Constructor_Main_Item{(*Constructor_Main_Item)(nil), (&Constructor_Main_Item{1, func() struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				} {
					orig := gopurs_runtime.RecordGet(__local_var_2_2, "entry")
					_ = orig
					clone := struct {
						label    string
						position struct {
							column int64
							line   int64
							/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
						}
						values []int64
						/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
					}{}
					clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
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
					clone.values = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
					return clone
				}(), (&Constructor_Data_List_Types_Cons[int64]{1, int64(31), (*Constructor_Data_List_Types_Cons[int64])(nil)})}), (&Constructor_Main_Item{1, struct {
					label    string
					position struct {
						column int64
						line   int64
						/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
					}
					values []int64
					/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
				}{"new", struct {
					column int64
					line   int64
					/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
				}{int64(19), int64(23)}, []int64{int64(29)}}, (*Constructor_Data_List_Types_Cons[int64])(nil)})}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer(v)}
				}
				return gopurs_runtime.Array(boxed)
			}()
			_ = arrayIntrinsic_arg_18
			_ = arrayIntrinsic_arg_18
			__local_var_12_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1268761764(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_8())), "", struct {
				actual   []string
				expected []string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []string {
				arr_val_arrayMap15 := arrayIntrinsic_arg_16
				_ = arr_val_arrayMap15
				_ = arr_val_arrayMap15
				arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
				_ = arr_go_arrayMap15
				res_go_arrayMap15 := make([]string, len(*arr_go_arrayMap15))
				_ = res_go_arrayMap15
				for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
					res_go_arrayMap15[i_arrayMap15] = Call_Main_describe(gopurs_runtime.CoerceToStruct[Constructor_Main_Item](v_arrayMap15))
				}
				return res_go_arrayMap15
			}(), func() []string {
				arr_val_arrayMap15 := arrayIntrinsic_arg_18
				_ = arr_val_arrayMap15
				_ = arr_val_arrayMap15
				arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
				_ = arr_go_arrayMap15
				res_go_arrayMap15 := make([]string, len(*arr_go_arrayMap15))
				_ = res_go_arrayMap15
				for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
					res_go_arrayMap15[i_arrayMap15] = Call_Main_describe(gopurs_runtime.CoerceToStruct[Constructor_Main_Item](v_arrayMap15))
				}
				return res_go_arrayMap15
			}()}), gopurs_runtime.Value{})
			_ = __local_var_12_14
			arrayIntrinsic_arg_20 := Get_Main_describe()
			_ = arrayIntrinsic_arg_20
			_ = arrayIntrinsic_arg_20
			arrayIntrinsic_arg_21 := gopurs_runtime.Apply2(Get_Main_appendItem(), func() gopurs_runtime.Value {
				arr := []*Constructor_Main_Item{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer(v)}
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Value{Type: 9, IntVal: 2860581675, UnsafePtr: unsafe.Pointer((*Constructor_Main_Item)(nil))})
			_ = arrayIntrinsic_arg_21
			_ = arrayIntrinsic_arg_21
			__local_var_13_19 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1268761764(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_9())), "", struct {
				actual   []string
				expected []string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []string {
				arr_val_arrayMap16 := arrayIntrinsic_arg_21
				_ = arr_val_arrayMap16
				_ = arr_val_arrayMap16
				arr_go_arrayMap16 := (*[]gopurs_runtime.Value)(arr_val_arrayMap16.UnsafePtr)
				_ = arr_go_arrayMap16
				res_go_arrayMap16 := make([]string, len(*arr_go_arrayMap16))
				_ = res_go_arrayMap16
				for i_arrayMap16, v_arrayMap16 := range *arr_go_arrayMap16 {
					res_go_arrayMap16[i_arrayMap16] = Call_Main_describe(gopurs_runtime.CoerceToStruct[Constructor_Main_Item](v_arrayMap16))
				}
				return res_go_arrayMap16
			}(), []string{"empty"}}), gopurs_runtime.Value{})
			_ = __local_var_13_19
			__local_var_14_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{int64(2), int64(2)}), gopurs_runtime.Value{})
			_ = __local_var_14_22
			__local_var_15_23 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Main_appendInt(), gopurs_runtime.RecordGet(__local_var_2_2, "ints")), gopurs_runtime.Value{})
			_ = __local_var_15_23
			__local_var_16_24 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_15_23, gopurs_runtime.Value{})
			_ = __local_var_16_24
			__local_var_17_25 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_10())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(__local_var_16_24, gopurs_runtime.Int(int64(5))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5)}}), gopurs_runtime.Value{})
			_ = __local_var_17_25
			__local_var_18_26 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_11())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(__local_var_16_24, gopurs_runtime.Int(int64(9))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(9)}}), gopurs_runtime.Value{})
			_ = __local_var_18_26
			__local_var_19_27 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), Get_Data_Array_snoc(), gopurs_runtime.Value{})
			_ = __local_var_19_27
			__local_var_20_28 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_19_27, gopurs_runtime.Value{})
			_ = __local_var_20_28
			__local_var_21_29 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1268761764(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_12())), "", struct {
				actual   []string
				expected []string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(__local_var_20_28, gopurs_runtime.RecordGet(__local_var_2_2, "strings"), gopurs_runtime.StrWithCodeUnits("alias", 5)).UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}(), []string{"", "λ😀", "alias"}}), gopurs_runtime.Value{})
			_ = __local_var_21_29
			__local_var_22_30 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1268761764(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_13())), "", struct {
				actual   []string
				expected []string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(__local_var_20_28, func() gopurs_runtime.Value {
					arr := []string{}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.StrWithCodeUnits("only", 4)).UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}(), []string{"only"}}), gopurs_runtime.Value{})
			_ = __local_var_22_30
			__local_var_23_31 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Apply2(Get_Main_appendInt(), gopurs_runtime.RecordGet(__local_var_2_2, "ints"), gopurs_runtime.Int(int64(5))), gopurs_runtime.Value{})
			_ = __local_var_23_31
			__local_var_24_32 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_23_31, gopurs_runtime.Value{})
			_ = __local_var_24_32
			__local_var_25_33 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_14())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), __local_var_24_32, gopurs_runtime.Int(int64(6))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5), int64(6)}}), gopurs_runtime.Value{})
			_ = __local_var_25_33
			__local_var_26_34 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_15())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Main_appendInt(), __local_var_24_32, gopurs_runtime.Int(int64(8))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5), int64(8)}}), gopurs_runtime.Value{})
			_ = __local_var_26_34
			__local_var_27_35 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_16())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(__local_var_24_32.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(-2147483648), int64(0), int64(2147483647), int64(5)}}), gopurs_runtime.Value{})
			_ = __local_var_27_35
			__local_var_28_36 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_1_1, gopurs_runtime.Value{})
			_ = __local_var_28_36
			__local_var_29_37 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_17())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_28_36, "ints").UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(-2147483648), int64(0), int64(2147483647)}}), gopurs_runtime.Value{})
			_ = __local_var_29_37
			__local_var_30_38 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___165009865(Rebox_Main_1386611502_940452149(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_18())), "", struct {
				actual   [][]int64
				expected [][]int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() [][]int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_28_36, "arrays").UnsafePtr)
				unboxed := make([][]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}()
				}
				return unboxed
			}(), [][]int64{[]int64{}, []int64{int64(1), int64(2)}}}), gopurs_runtime.Value{})
			_ = __local_var_30_38
			__local_var_31_39 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1268761764(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_19())), "", struct {
				actual   []string
				expected []string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_28_36, "strings").UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}(), []string{"", "λ😀"}}), gopurs_runtime.Value{})
			_ = __local_var_31_39
			__local_var_32_40 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(10)), gopurs_runtime.Value{})
			_ = __local_var_32_40
			__local_var_33_41 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(20)), gopurs_runtime.Value{})
			_ = __local_var_33_41
			__local_var_34_42 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				arr := []struct {
					cell  gopurs_runtime.Value
					label string
					/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
				}{struct {
					cell  gopurs_runtime.Value
					label string
					/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
				}{__local_var_32_40, "shared"}}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						orig := v
						_ = orig
						return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, orig.cell, gopurs_runtime.Str(orig.label))
					}()
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Value{})
			_ = __local_var_34_42
			__local_var_35_43 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_34_42, gopurs_runtime.Value{})
			_ = __local_var_35_43
			// TAST (Let): withOther_36_44 shape=App(Var) bindingType=(Array (Record (Row [label: String, cell: (ADT ["Effect","Ref","Ref"] [Int])] Empty)))
			withOther_36_44 := Call_Data_Array_snoc__3427008165(func() []struct {
				cell  gopurs_runtime.Value
				label string
				/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
			} {
				arr := *(*[]gopurs_runtime.Value)(__local_var_35_43.UnsafePtr)
				unboxed := make([]struct {
					cell  gopurs_runtime.Value
					label string
					/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
				}, len(arr))
				for i, v := range arr {
					unboxed[i] = func() struct {
						cell  gopurs_runtime.Value
						label string
						/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
					} {
						orig := v
						_ = orig
						clone := struct {
							cell  gopurs_runtime.Value
							label string
							/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
						}{}
						clone.cell = gopurs_runtime.RecordGet(orig, "cell")
						clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
						return clone
					}()
				}
				return unboxed
			}(), struct {
				cell  gopurs_runtime.Value
				label string
				/*gopurs:record-v1:2:00630065006c006c,006c006100620065006c*/
			}{__local_var_33_41, "other"})
			_ = withOther_36_44
			var __t47 gopurs_runtime.Value
			{
				if (gopurs_runtime.Int(int64(len(withOther_36_44))).IntVal) == (int64(2)) {
					__t47 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_37_46 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(41)), gopurs_runtime.RecordGet(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
							arr := withOther_36_44
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, orig.cell, gopurs_runtime.Str(orig.label))
								}()
							}
							return gopurs_runtime.Array(boxed)
						}(), 0), "cell"), gopurs_runtime.Value{})
						_ = __local_var_37_46
						return gopurs_runtime.Apply3(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(42)), gopurs_runtime.RecordGet(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
							arr := withOther_36_44
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_3, orig.cell, gopurs_runtime.Str(orig.label))
								}()
							}
							return gopurs_runtime.Array(boxed)
						}(), 1), "cell"), gopurs_runtime.Value{})
					})
					goto end_branch_47
				} else {

				}
			}
			{
				__t47 = Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Int(int64(len(withOther_36_44))).IntVal, int64(2)})
			}
		end_branch_47:
			__local_var_37_45 := gopurs_runtime.Apply(__t47, gopurs_runtime.Value{})
			_ = __local_var_37_45
			__local_var_38_48 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_32_40, gopurs_runtime.Value{})
			_ = __local_var_38_48
			__local_var_39_49 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_33_41, gopurs_runtime.Value{})
			_ = __local_var_39_49
			__local_var_40_50 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{__local_var_38_48.IntVal, int64(41)}), gopurs_runtime.Value{})
			_ = __local_var_40_50
			__local_var_41_51 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{__local_var_39_49.IntVal, int64(42)}), gopurs_runtime.Value{})
			_ = __local_var_41_51
			__local_var_42_52 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_35_43))).IntVal, int64(1)}), gopurs_runtime.Value{})
			_ = __local_var_42_52
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))})
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))})
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = Call_Data_Show_showArray(Get_Main_showArray1())
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = Call_Data_Show_showArray(Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_11
}

var cache_Main___cached_dict_12 gopurs_runtime.Value
var once_Main___cached_dict_12 sync.Once

func Get_Main___cached_dict_12() gopurs_runtime.Value {
	once_Main___cached_dict_12.Do(func() {
		cache_Main___cached_dict_12 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_12
}

var cache_Main___cached_dict_13 gopurs_runtime.Value
var once_Main___cached_dict_13 sync.Once

func Get_Main___cached_dict_13() gopurs_runtime.Value {
	once_Main___cached_dict_13.Do(func() {
		cache_Main___cached_dict_13 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_13
}

var cache_Main___cached_dict_14 gopurs_runtime.Value
var once_Main___cached_dict_14 sync.Once

func Get_Main___cached_dict_14() gopurs_runtime.Value {
	once_Main___cached_dict_14.Do(func() {
		cache_Main___cached_dict_14 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_14
}

var cache_Main___cached_dict_15 gopurs_runtime.Value
var once_Main___cached_dict_15 sync.Once

func Get_Main___cached_dict_15() gopurs_runtime.Value {
	once_Main___cached_dict_15.Do(func() {
		cache_Main___cached_dict_15 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_15
}

var cache_Main___cached_dict_16 gopurs_runtime.Value
var once_Main___cached_dict_16 sync.Once

func Get_Main___cached_dict_16() gopurs_runtime.Value {
	once_Main___cached_dict_16.Do(func() {
		cache_Main___cached_dict_16 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_16
}

var cache_Main___cached_dict_17 gopurs_runtime.Value
var once_Main___cached_dict_17 sync.Once

func Get_Main___cached_dict_17() gopurs_runtime.Value {
	once_Main___cached_dict_17.Do(func() {
		cache_Main___cached_dict_17 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main___cached_dict_17
}

var cache_Main___cached_dict_18 gopurs_runtime.Value
var once_Main___cached_dict_18 sync.Once

func Get_Main___cached_dict_18() gopurs_runtime.Value {
	once_Main___cached_dict_18.Do(func() {
		cache_Main___cached_dict_18 = Call_Data_Show_showArray(Get_Main_showArray1())
	})
	return cache_Main___cached_dict_18
}

var cache_Main___cached_dict_19 gopurs_runtime.Value
var once_Main___cached_dict_19 sync.Once

func Get_Main___cached_dict_19() gopurs_runtime.Value {
	once_Main___cached_dict_19.Do(func() {
		cache_Main___cached_dict_19 = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main___cached_dict_19
}

type Constructor_Main_Empty struct {
	Rc uint32
}

type Constructor_Main_Item struct {
	Rc uint32
	V0 struct {
		label    string
		position struct {
			column int64
			line   int64
			/*gopurs:record-v1:2:0063006f006c0075006d006e,006c0069006e0065*/
		}
		values []int64
		/*gopurs:record-v1:3:006c006100620065006c,0070006f0073006900740069006f006e,00760061006c007500650073*/
	}
	V1 *Constructor_Data_List_Types_Cons[int64]
}

func Call_Main_negate(a_0_loop float64) float64 {
	var a_0 float64 = a_0_loop
	_ = a_0
	return -(a_0)
}

func Call_Main_negate1(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Main_describe(v_0_loop *Constructor_Main_Item) string {
	var v_0 *Constructor_Main_Item = v_0_loop
	_ = v_0
	var __t0 string
	{
		if v_0 == nil {
			__t0 = "empty"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 != nil {
			__t0 = gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString((v_0).V0.label, ":"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Get_Main_showRecordFieldsCons()), "show"), func() gopurs_runtime.Value {
				orig := (v_0).V0.position
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Int(orig.column), gopurs_runtime.Int(orig.line))
			}()).StrVal()), ":"), gopurs_runtime.Apply(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showArray())).V0, func() gopurs_runtime.Value {
				arr := (v_0).V0.values
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).StrVal()), ":"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_List_Types_showList(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), "show"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((v_0).V1))}).StrVal())
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

func Call_Main_applyAppend(append1_0_loop gopurs_runtime.Value, values_1_loop []gopurs_runtime.Value, value_2_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var append1_0 gopurs_runtime.Value = append1_0_loop
	_ = append1_0
	var values_1 []gopurs_runtime.Value = values_1_loop
	_ = values_1
	var value_2 gopurs_runtime.Value = value_2_loop
	_ = value_2
	return (*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(append1_0, gopurs_runtime.Array(values_1), value_2).UnsafePtr)))
}

func Call_Main_applyAppend__3199238329(append1_0_loop gopurs_runtime.Value, values_1_loop []string, value_2_loop string) []string {
applyAppend__3199238329:
	for {
		if false {
			continue applyAppend__3199238329
		}
		var append1_0 gopurs_runtime.Value = append1_0_loop
		_ = append1_0
		var values_1 []string = values_1_loop
		_ = values_1
		var value_2 string = value_2_loop
		_ = value_2
		return func() []string {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(append1_0, func() gopurs_runtime.Value {
				arr := values_1
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), gopurs_runtime.Str(value_2)).UnsafePtr)
			unboxed := make([]string, len(arr))
			for i, v := range arr {
				unboxed[i] = v.StrVal()
			}
			return unboxed
		}()
	}
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]int64])(unsafe.Pointer(in))
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

func Rebox_Main_1386611502_1953100407(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]string])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[float64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_3467499472(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]float64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_38358956(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]bool])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_927154015(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]*Constructor_Data_List_Types_Cons[int64]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]*Constructor_Data_List_Types_Cons[int64]])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_940452149(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[][]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[][]int64])(unsafe.Pointer(in))
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

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
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

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys3{K0: "label", K1: "position", K2: "values"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys2{K0: "column", K1: "line"}

var GopursRecordKeys_4d61696e_3 = gopurs_runtime.RecordKeys2{K0: "cell", K1: "label"}
