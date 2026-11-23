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
		cache_Main_bind = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind
}

var cache_Main_bind1 gopurs_runtime.Value
var once_Main_bind1 sync.Once

func Get_Main_bind1() gopurs_runtime.Value {
	once_Main_bind1.Do(func() {
		cache_Main_bind1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind1
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Main_add
}

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Main_pure
}

var cache_Main_bind2 gopurs_runtime.Value
var once_Main_bind2 sync.Once

func Get_Main_bind2() gopurs_runtime.Value {
	once_Main_bind2.Do(func() {
		cache_Main_bind2 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind2
}

var cache_Main_discard gopurs_runtime.Value
var once_Main_discard sync.Once

func Get_Main_discard() gopurs_runtime.Value {
	once_Main_discard.Do(func() {
		cache_Main_discard = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_discard
}

var cache_Main_pure1 gopurs_runtime.Value
var once_Main_pure1 sync.Once

func Get_Main_pure1() gopurs_runtime.Value {
	once_Main_pure1.Do(func() {
		cache_Main_pure1 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Main_pure1
}

var cache_Main_bind3 gopurs_runtime.Value
var once_Main_bind3 sync.Once

func Get_Main_bind3() gopurs_runtime.Value {
	once_Main_bind3.Do(func() {
		cache_Main_bind3 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind3
}

var cache_Main_void gopurs_runtime.Value
var once_Main_void sync.Once

func Get_Main_void() gopurs_runtime.Value {
	once_Main_void.Do(func() {
		cache_Main_void = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_void(__local_var_0_box)
		})
	})
	return cache_Main_void
}

var cache_Main_bind4 gopurs_runtime.Value
var once_Main_bind4 sync.Once

func Get_Main_bind4() gopurs_runtime.Value {
	once_Main_bind4.Do(func() {
		cache_Main_bind4 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind4
}

var cache_Main_bind5 gopurs_runtime.Value
var once_Main_bind5 sync.Once

func Get_Main_bind5() gopurs_runtime.Value {
	once_Main_bind5.Do(func() {
		cache_Main_bind5 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind5
}

var cache_Main_bind6 gopurs_runtime.Value
var once_Main_bind6 sync.Once

func Get_Main_bind6() gopurs_runtime.Value {
	once_Main_bind6.Do(func() {
		cache_Main_bind6 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind6
}

var cache_Main_bind7 gopurs_runtime.Value
var once_Main_bind7 sync.Once

func Get_Main_bind7() gopurs_runtime.Value {
	once_Main_bind7.Do(func() {
		cache_Main_bind7 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind7
}

var cache_Main_bind8 gopurs_runtime.Value
var once_Main_bind8 sync.Once

func Get_Main_bind8() gopurs_runtime.Value {
	once_Main_bind8.Do(func() {
		cache_Main_bind8 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind8
}

var cache_Main_bind9 gopurs_runtime.Value
var once_Main_bind9 sync.Once

func Get_Main_bind9() gopurs_runtime.Value {
	once_Main_bind9.Do(func() {
		cache_Main_bind9 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind9
}

var cache_Main_pure2 gopurs_runtime.Value
var once_Main_pure2 sync.Once

func Get_Main_pure2() gopurs_runtime.Value {
	once_Main_pure2.Do(func() {
		cache_Main_pure2 = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Main_pure2
}

var cache_Main_bind10 gopurs_runtime.Value
var once_Main_bind10 sync.Once

func Get_Main_bind10() gopurs_runtime.Value {
	once_Main_bind10.Do(func() {
		cache_Main_bind10 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind10
}

var cache_Main_bind11 gopurs_runtime.Value
var once_Main_bind11 sync.Once

func Get_Main_bind11() gopurs_runtime.Value {
	once_Main_bind11.Do(func() {
		cache_Main_bind11 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind11
}

var cache_Main_bind12 gopurs_runtime.Value
var once_Main_bind12 sync.Once

func Get_Main_bind12() gopurs_runtime.Value {
	once_Main_bind12.Do(func() {
		cache_Main_bind12 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind12
}

var cache_Main_bind13 gopurs_runtime.Value
var once_Main_bind13 sync.Once

func Get_Main_bind13() gopurs_runtime.Value {
	once_Main_bind13.Do(func() {
		cache_Main_bind13 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind13
}

var cache_Main_discard1 gopurs_runtime.Value
var once_Main_discard1 sync.Once

func Get_Main_discard1() gopurs_runtime.Value {
	once_Main_discard1.Do(func() {
		cache_Main_discard1 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_discard1
}

var cache_Main_bind14 gopurs_runtime.Value
var once_Main_bind14 sync.Once

func Get_Main_bind14() gopurs_runtime.Value {
	once_Main_bind14.Do(func() {
		cache_Main_bind14 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind14
}

var cache_Main_void1 gopurs_runtime.Value
var once_Main_void1 sync.Once

func Get_Main_void1() gopurs_runtime.Value {
	once_Main_void1.Do(func() {
		cache_Main_void1 = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_void1(__local_var_0_box)
		})
	})
	return cache_Main_void1
}

var cache_Main_bind15 gopurs_runtime.Value
var once_Main_bind15 sync.Once

func Get_Main_bind15() gopurs_runtime.Value {
	once_Main_bind15.Do(func() {
		cache_Main_bind15 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind15
}

var cache_Main_bind16 gopurs_runtime.Value
var once_Main_bind16 sync.Once

func Get_Main_bind16() gopurs_runtime.Value {
	once_Main_bind16.Do(func() {
		cache_Main_bind16 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind16
}

var cache_Main_void2 gopurs_runtime.Value
var once_Main_void2 sync.Once

func Get_Main_void2() gopurs_runtime.Value {
	once_Main_void2.Do(func() {
		cache_Main_void2 = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_void2(__local_var_0_box)
		})
	})
	return cache_Main_void2
}

var cache_Main_bind17 gopurs_runtime.Value
var once_Main_bind17 sync.Once

func Get_Main_bind17() gopurs_runtime.Value {
	once_Main_bind17.Do(func() {
		cache_Main_bind17 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind17
}

var cache_Main_bind18 gopurs_runtime.Value
var once_Main_bind18 sync.Once

func Get_Main_bind18() gopurs_runtime.Value {
	once_Main_bind18.Do(func() {
		cache_Main_bind18 = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Main_bind18
}

var cache_Main_lessThan gopurs_runtime.Value
var once_Main_lessThan sync.Once

func Get_Main_lessThan() gopurs_runtime.Value {
	once_Main_lessThan.Do(func() {
		cache_Main_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Main_lessThan
}

var cache_Main_mul gopurs_runtime.Value
var once_Main_mul sync.Once

func Get_Main_mul() gopurs_runtime.Value {
	once_Main_mul.Do(func() {
		cache_Main_mul = Get_Data_Semiring_mul__4086903470()
	})
	return cache_Main_mul
}

var cache_Main_bind19 gopurs_runtime.Value
var once_Main_bind19 sync.Once

func Get_Main_bind19() gopurs_runtime.Value {
	once_Main_bind19.Do(func() {
		cache_Main_bind19 = Get_Control_Bind_bind__333752852()
	})
	return cache_Main_bind19
}

var cache_Main_bind20 gopurs_runtime.Value
var once_Main_bind20 sync.Once

func Get_Main_bind20() gopurs_runtime.Value {
	once_Main_bind20.Do(func() {
		cache_Main_bind20 = Get_Control_Bind_bind__3692032020()
	})
	return cache_Main_bind20
}

var cache_Main_bind21 gopurs_runtime.Value
var once_Main_bind21 sync.Once

func Get_Main_bind21() gopurs_runtime.Value {
	once_Main_bind21.Do(func() {
		cache_Main_bind21 = Get_Control_Bind_bind__2880717012()
	})
	return cache_Main_bind21
}

var cache_Main_discard2 gopurs_runtime.Value
var once_Main_discard2 sync.Once

func Get_Main_discard2() gopurs_runtime.Value {
	once_Main_discard2.Do(func() {
		cache_Main_discard2 = Get_Control_Bind_discard__143037513()
	})
	return cache_Main_discard2
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

var cache_Main_eqRowCons1 gopurs_runtime.Value
var once_Main_eqRowCons1 sync.Once

func Get_Main_eqRowCons1() gopurs_runtime.Value {
	once_Main_eqRowCons1.Do(func() {
		cache_Main_eqRowCons1 = gopurs_runtime.Apply(Get_Main_eqRowCons(), Get_Main_valuesIsSymbol())
	})
	return cache_Main_eqRowCons1
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main_eqArray
}

var cache_Main_nextIsSymbol gopurs_runtime.Value
var once_Main_nextIsSymbol sync.Once

func Get_Main_nextIsSymbol() gopurs_runtime.Value {
	once_Main_nextIsSymbol.Do(func() {
		cache_Main_nextIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("next", 4)
		}))
	})
	return cache_Main_nextIsSymbol
}

var cache_Main_failedIsSymbol gopurs_runtime.Value
var once_Main_failedIsSymbol sync.Once

func Get_Main_failedIsSymbol() gopurs_runtime.Value {
	once_Main_failedIsSymbol.Do(func() {
		cache_Main_failedIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("failed", 6)
		}))
	})
	return cache_Main_failedIsSymbol
}

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Apply2(Get_Data_Show_showRecord(), gopurs_runtime.Value{}, gopurs_runtime.Value{})
	})
	return cache_Main_showRecord
}

var cache_Main_showRecordFieldsCons gopurs_runtime.Value
var once_Main_showRecordFieldsCons sync.Once

func Get_Main_showRecordFieldsCons() gopurs_runtime.Value {
	once_Main_showRecordFieldsCons.Do(func() {
		cache_Main_showRecordFieldsCons = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsCons(), Get_Main_failedIsSymbol())
	})
	return cache_Main_showRecordFieldsCons
}

var cache_Main_showRecordFieldsCons1 gopurs_runtime.Value
var once_Main_showRecordFieldsCons1 sync.Once

func Get_Main_showRecordFieldsCons1() gopurs_runtime.Value {
	once_Main_showRecordFieldsCons1.Do(func() {
		cache_Main_showRecordFieldsCons1 = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsCons(), Get_Main_nextIsSymbol())
	})
	return cache_Main_showRecordFieldsCons1
}

var cache_Main_showRecordFieldsConsNil gopurs_runtime.Value
var once_Main_showRecordFieldsConsNil sync.Once

func Get_Main_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_Main_showRecordFieldsConsNil.Do(func() {
		cache_Main_showRecordFieldsConsNil = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsConsNil(), Get_Main_valuesIsSymbol())
	})
	return cache_Main_showRecordFieldsConsNil
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main_showArray
}

var cache_Main_assertEqual gopurs_runtime.Value
var once_Main_assertEqual sync.Once

func Get_Main_assertEqual() gopurs_runtime.Value {
	once_Main_assertEqual.Do(func() {
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__2478740706()
	})
	return cache_Main_assertEqual
}

var cache_Main_assertEqual1 gopurs_runtime.Value
var once_Main_assertEqual1 sync.Once

func Get_Main_assertEqual1() gopurs_runtime.Value {
	once_Main_assertEqual1.Do(func() {
		cache_Main_assertEqual1 = Get_Test_Assert_assertEqual__2891126901()
	})
	return cache_Main_assertEqual1
}

var cache_Main_assertEqual2 gopurs_runtime.Value
var once_Main_assertEqual2 sync.Once

func Get_Main_assertEqual2() gopurs_runtime.Value {
	once_Main_assertEqual2.Do(func() {
		cache_Main_assertEqual2 = Get_Test_Assert_assertEqual__1123177538()
	})
	return cache_Main_assertEqual2
}

var cache_Main_bind22 gopurs_runtime.Value
var once_Main_bind22 sync.Once

func Get_Main_bind22() gopurs_runtime.Value {
	once_Main_bind22.Do(func() {
		cache_Main_bind22 = Get_Control_Bind_bind__2595690612()
	})
	return cache_Main_bind22
}

var cache_Main_bind23 gopurs_runtime.Value
var once_Main_bind23 sync.Once

func Get_Main_bind23() gopurs_runtime.Value {
	once_Main_bind23.Do(func() {
		cache_Main_bind23 = Get_Control_Bind_bind__1922180980()
	})
	return cache_Main_bind23
}

var cache_Main_assertEqual3 gopurs_runtime.Value
var once_Main_assertEqual3 sync.Once

func Get_Main_assertEqual3() gopurs_runtime.Value {
	once_Main_assertEqual3.Do(func() {
		cache_Main_assertEqual3 = Get_Test_Assert_assertEqual__946142709()
	})
	return cache_Main_assertEqual3
}

var cache_Main_assertEqual4 gopurs_runtime.Value
var once_Main_assertEqual4 sync.Once

func Get_Main_assertEqual4() gopurs_runtime.Value {
	once_Main_assertEqual4.Do(func() {
		cache_Main_assertEqual4 = Get_Test_Assert_assertEqual__2867988789()
	})
	return cache_Main_assertEqual4
}

var cache_Main_beforeIsSymbol gopurs_runtime.Value
var once_Main_beforeIsSymbol sync.Once

func Get_Main_beforeIsSymbol() gopurs_runtime.Value {
	once_Main_beforeIsSymbol.Do(func() {
		cache_Main_beforeIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("before", 6)
		}))
	})
	return cache_Main_beforeIsSymbol
}

var cache_Main_afterIsSymbol gopurs_runtime.Value
var once_Main_afterIsSymbol sync.Once

func Get_Main_afterIsSymbol() gopurs_runtime.Value {
	once_Main_afterIsSymbol.Do(func() {
		cache_Main_afterIsSymbol = gopurs_runtime.RecordShape1(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("after", 5)
		}))
	})
	return cache_Main_afterIsSymbol
}

var cache_Main_assertEqual5 gopurs_runtime.Value
var once_Main_assertEqual5 sync.Once

func Get_Main_assertEqual5() gopurs_runtime.Value {
	once_Main_assertEqual5.Do(func() {
		cache_Main_assertEqual5 = Get_Test_Assert_assertEqual__2101451337()
	})
	return cache_Main_assertEqual5
}

var cache_Main_Emit gopurs_runtime.Value
var once_Main_Emit sync.Once

func Get_Main_Emit() gopurs_runtime.Value {
	once_Main_Emit.Do(func() {
		cache_Main_Emit = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, value0, value1.IntVal}))}
			})
		})
	})
	return cache_Main_Emit
}

var cache_Main_Emit__1834327443 gopurs_runtime.Value
var once_Main_Emit__1834327443 sync.Once

func Get_Main_Emit__1834327443() gopurs_runtime.Value {
	once_Main_Emit__1834327443.Do(func() {
		cache_Main_Emit__1834327443 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Emit__1834327443(func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_1_0_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), __eta_norm_0_1_box.IntVal)
		})
	})
	return cache_Main_Emit__1834327443
}

var cache_Main_Emit__1315694163 gopurs_runtime.Value
var once_Main_Emit__1315694163 sync.Once

func Get_Main_Emit__1315694163() gopurs_runtime.Value {
	once_Main_Emit__1315694163.Do(func() {
		cache_Main_Emit__1315694163 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Emit__1315694163(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.IntVal)
		})
	})
	return cache_Main_Emit__1315694163
}

var cache_Main_Emit__203891187 gopurs_runtime.Value
var once_Main_Emit__203891187 sync.Once

func Get_Main_Emit__203891187() gopurs_runtime.Value {
	once_Main_Emit__203891187.Do(func() {
		cache_Main_Emit__203891187 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Emit__203891187(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box.IntVal)
		})
	})
	return cache_Main_Emit__203891187
}

var cache_Main_Stop gopurs_runtime.Value
var once_Main_Stop sync.Once

func Get_Main_Stop() gopurs_runtime.Value {
	once_Main_Stop.Do(func() {
		cache_Main_Stop = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, value0.IntVal}))}
		})
	})
	return cache_Main_Stop
}

var cache_Main_Stop__1606680011 gopurs_runtime.Value
var once_Main_Stop__1606680011 sync.Once

func Get_Main_Stop__1606680011() gopurs_runtime.Value {
	once_Main_Stop__1606680011.Do(func() {
		cache_Main_Stop__1606680011 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Stop__1606680011(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_Stop__1606680011
}

var cache_Main_Stop__2802585741 gopurs_runtime.Value
var once_Main_Stop__2802585741 sync.Once

func Get_Main_Stop__2802585741() gopurs_runtime.Value {
	once_Main_Stop__2802585741.Do(func() {
		cache_Main_Stop__2802585741 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Stop__2802585741(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_Stop__2802585741
}

var cache_Main_Stop__916476971 gopurs_runtime.Value
var once_Main_Stop__916476971 sync.Once

func Get_Main_Stop__916476971() gopurs_runtime.Value {
	once_Main_Stop__916476971.Do(func() {
		cache_Main_Stop__916476971 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Stop__916476971(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_Stop__916476971
}

var cache_Main_Failed gopurs_runtime.Value
var once_Main_Failed sync.Once

func Get_Main_Failed() gopurs_runtime.Value {
	once_Main_Failed.Do(func() {
		cache_Main_Failed = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 270449853, UnsafePtr: unsafe.Pointer((&Constructor_Main_Failed[gopurs_runtime.Value]{1, value0.IntVal}))}
		})
	})
	return cache_Main_Failed
}

var cache_Main_Failed__2802585741 gopurs_runtime.Value
var once_Main_Failed__2802585741 sync.Once

func Get_Main_Failed__2802585741() gopurs_runtime.Value {
	once_Main_Failed__2802585741.Do(func() {
		cache_Main_Failed__2802585741 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Failed__2802585741(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_Failed__2802585741
}

var cache_Main_snapshotRead gopurs_runtime.Value
var once_Main_snapshotRead sync.Once

func Get_Main_snapshotRead() gopurs_runtime.Value {
	once_Main_snapshotRead.Do(func() {
		cache_Main_snapshotRead = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_snapshotRead(initial_0_box.IntVal)
		})
	})
	return cache_Main_snapshotRead
}

var cache_Main_skipLoop gopurs_runtime.Value
var once_Main_skipLoop sync.Once

func Get_Main_skipLoop() gopurs_runtime.Value {
	once_Main_skipLoop.Do(func() {
		cache_Main_skipLoop = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_skipLoop(initial_0_box.IntVal))
		})
	})
	return cache_Main_skipLoop
}

var cache_Main_operationResults gopurs_runtime.Value
var once_Main_operationResults sync.Once

func Get_Main_operationResults() gopurs_runtime.Value {
	once_Main_operationResults.Do(func() {
		cache_Main_operationResults = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_operationResults(initial_0_box.IntVal)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_operationResults
}

var cache_Main_freezeSnapshot gopurs_runtime.Value
var once_Main_freezeSnapshot sync.Once

func Get_Main_freezeSnapshot() gopurs_runtime.Value {
	once_Main_freezeSnapshot.Do(func() {
		cache_Main_freezeSnapshot = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_freezeSnapshot(initial_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, func() gopurs_runtime.Value {
					arr := orig.after
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.before
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()
		})
	})
	return cache_Main_freezeSnapshot
}

var cache_Main_collect gopurs_runtime.Value
var once_Main_collect sync.Once

func Get_Main_collect() gopurs_runtime.Value {
	once_Main_collect.Do(func() {
		cache_Main_collect = gopurs_runtime.Func2(func(step_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_collect(step_0_box, start_1_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), gopurs_runtime.Array(orig.values))
			}()
		})
	})
	return cache_Main_collect
}

var cache_Main_collect__2173325217 gopurs_runtime.Value
var once_Main_collect__2173325217 sync.Once

func Get_Main_collect__2173325217() gopurs_runtime.Value {
	once_Main_collect__2173325217.Do(func() {
		cache_Main_collect__2173325217 = gopurs_runtime.Func2(func(step_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_collect__2173325217(step_0_box, start_1_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
					arr := orig.values
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
				}())
			}()
		})
	})
	return cache_Main_collect__2173325217
}

var cache_Main_collect__110832033 gopurs_runtime.Value
var once_Main_collect__110832033 sync.Once

func Get_Main_collect__110832033() gopurs_runtime.Value {
	once_Main_collect__110832033.Do(func() {
		cache_Main_collect__110832033 = gopurs_runtime.Func2(func(step_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_collect__110832033(step_0_box, start_1_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()
		})
	})
	return cache_Main_collect__110832033
}

var cache_Main_collect__1788204577 gopurs_runtime.Value
var once_Main_collect__1788204577 sync.Once

func Get_Main_collect__1788204577() gopurs_runtime.Value {
	once_Main_collect__1788204577.Do(func() {
		cache_Main_collect__1788204577 = gopurs_runtime.Func2(func(step_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_collect__1788204577(step_0_box, start_1_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()
		})
	})
	return cache_Main_collect__1788204577
}

var cache_Main_aliasRead gopurs_runtime.Value
var once_Main_aliasRead sync.Once

func Get_Main_aliasRead() gopurs_runtime.Value {
	once_Main_aliasRead.Do(func() {
		cache_Main_aliasRead = gopurs_runtime.Func(func(initial_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_aliasRead(initial_0_box.IntVal))
		})
	})
	return cache_Main_aliasRead
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if (n_0.IntVal) < (int64(4)) {
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.IntMul(n_0.IntVal, int64(3))), gopurs_runtime.IntAdd(n_0.IntVal, int64(1))}))}
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, gopurs_runtime.IntAdd(n_0.IntVal, int64(7))}))}
				}
			end_branch_1:
				return __t1
			}), gopurs_runtime.Value{})
			_ = __local_var_0_0
			__local_var_1_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_2
			__callback_fn_4 := Get_Effect_Ref__new()
			_ = __callback_fn_4
			__callback_arg_5 := Call_Main_collect__110832033(__local_var_1_2, int64(1))
			_ = __callback_arg_5
			__local_var_2_3 := gopurs_runtime.Apply(func(__native_callback func(struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}) gopurs_runtime.Value, __native_callback_ok bool) gopurs_runtime.Value {
				if __native_callback_ok {
					return __native_callback(__callback_arg_5)
				} else {

				}
				return gopurs_runtime.Apply(__callback_fn_4, func() gopurs_runtime.Value {
					orig := __callback_arg_5
					_ = orig
					return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
						arr := orig.values
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}())
			}(gopurs_runtime.FunctionData[func(struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}) gopurs_runtime.Value](__callback_fn_4)), gopurs_runtime.Value{})
			_ = __local_var_2_3
			__local_var_3_6 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_2_3, gopurs_runtime.Value{})
			_ = __local_var_3_6
			__local_var_4_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__2478740706(struct {
				actual struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				expected struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			} {
				orig := __local_var_3_6
				_ = orig
				clone := struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{}
				clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
				clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
				clone.values = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}(), struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, int64(11), []int64{int64(3), int64(6), int64(9)}}}), gopurs_runtime.Value{})
			_ = __local_var_4_7
			__local_var_5_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__2478740706(struct {
				actual struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				expected struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_collect__110832033(__local_var_1_2, int64(3)), struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, int64(11), []int64{int64(9)}}}), gopurs_runtime.Value{})
			_ = __local_var_5_8
			__local_var_6_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__2478740706(struct {
				actual struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				expected struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_collect__110832033(__local_var_1_2, int64(5)), struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, int64(12), []int64{}}}), gopurs_runtime.Value{})
			_ = __local_var_6_9
			__local_var_7_10 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_2_3, gopurs_runtime.Value{})
			_ = __local_var_7_10
			__local_var_8_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__2478740706(struct {
				actual struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				expected struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			} {
				orig := __local_var_7_10
				_ = orig
				clone := struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}{}
				clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
				clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
				clone.values = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				return clone
			}(), struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, int64(11), []int64{int64(3), int64(6), int64(9)}}}), gopurs_runtime.Value{})
			_ = __local_var_8_11
			__local_var_9_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__2478740706(struct {
				actual struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				expected struct {
					failed bool
					next   int64
					values []int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_collect__110832033(gopurs_runtime.Func(func(n_9 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t13 gopurs_runtime.Value
				{
					if (n_9.IntVal) < (int64(2)) {
						__t13 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, n_9, gopurs_runtime.IntAdd(n_9.IntVal, int64(1))}))}
						goto end_branch_13
					} else {

					}
				}
				{
					__t13 = gopurs_runtime.Value{Type: 9, IntVal: 270449853, UnsafePtr: unsafe.Pointer((&Constructor_Main_Failed[gopurs_runtime.Value]{1, int64(91)}))}
				}
			end_branch_13:
				return __t13
			}), int64(0)), struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{true, int64(91), []int64{}}}), gopurs_runtime.Value{})
			_ = __local_var_9_12
			__local_var_10_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__2891126901(struct {
				actual struct {
					failed bool
					next   int64
					values []string
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				expected struct {
					failed bool
					next   int64
					values []string
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_collect__1788204577(gopurs_runtime.Func(func(n_10 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t15 gopurs_runtime.Value
				{
					if (n_10.IntVal) < (int64(2)) {
						__t15 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, gopurs_runtime.StrWithCodeUnits("λ😀", 3), gopurs_runtime.IntAdd(n_10.IntVal, int64(1))}))}
						goto end_branch_15
					} else {

					}
				}
				{
					__t15 = gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, n_10.IntVal}))}
				}
			end_branch_15:
				return __t15
			}), int64(0)), struct {
				failed bool
				next   int64
				values []string
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, int64(2), []string{"λ😀", "λ😀"}}}), gopurs_runtime.Value{})
			_ = __local_var_10_14
			__local_var_11_16 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__1123177538(struct {
				actual struct {
					failed bool
					next   int64
					values [][]int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				expected struct {
					failed bool
					next   int64
					values [][]int64
					/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_collect__2173325217(gopurs_runtime.Func(func(n_11 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t17 gopurs_runtime.Value
				{
					if (n_11.IntVal) < (int64(2)) {
						__t17 = gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer((&Constructor_Main_Emit[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
							arr := []int64{n_11.IntVal, gopurs_runtime.IntAdd(n_11.IntVal, int64(1))}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}(), gopurs_runtime.IntAdd(n_11.IntVal, int64(1))}))}
						goto end_branch_17
					} else {

					}
				}
				{
					__t17 = gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stop[gopurs_runtime.Value]{1, n_11.IntVal}))}
				}
			end_branch_17:
				return __t17
			}), int64(0)), struct {
				failed bool
				next   int64
				values [][]int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, int64(2), [][]int64{[]int64{int64(0), int64(1)}, []int64{int64(1), int64(2)}}}}), gopurs_runtime.Value{})
			_ = __local_var_11_16
			__local_var_12_19 := gopurs_runtime.Apply2(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(10)), gopurs_runtime.Value{})
			_ = __local_var_12_19
			__local_var_12_18 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_12_19, gopurs_runtime.Value{})
			_ = __local_var_12_18
			__local_var_13_20 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___3119790530(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_0())), "", struct {
				actual   []int64
				expected []int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_operationResults(__local_var_12_18.IntVal), []int64{int64(11), int64(12)}}), gopurs_runtime.Value{})
			_ = __local_var_13_20
			__local_var_14_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_skipLoop(__local_var_12_18.IntVal), int64(10)}), gopurs_runtime.Value{})
			_ = __local_var_14_21
			__local_var_15_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual__2101451337(struct {
				actual struct {
					after  []int64
					before []int64
					/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
				}
				expected struct {
					after  []int64
					before []int64
					/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_freezeSnapshot(__local_var_12_18.IntVal), struct {
				after  []int64
				before []int64
				/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
			}{[]int64{int64(10), int64(11)}, []int64{int64(10)}}}), gopurs_runtime.Value{})
			_ = __local_var_15_22
			__local_var_16_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_aliasRead(__local_var_12_18.IntVal), int64(11)}), gopurs_runtime.Value{})
			_ = __local_var_16_23
			__local_var_17_24 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(Call_Main_snapshotRead(__local_var_12_18.IntVal), gopurs_runtime.Int(int64(5))).IntVal, int64(15)}), gopurs_runtime.Value{})
			_ = __local_var_17_24
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

type Constructor_Main_Emit[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 int64
}

type Constructor_Main_Stop[T_a any] struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_Failed[T_a any] struct {
	Rc uint32
	V0 int64
}

func Call_Main_void(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Main_void1(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Main_void2(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return Get_Data_Unit_unit()
	})
}

func Call_Main_Emit__1834327443(__eta_norm_1_0_loop []int64, __eta_norm_0_1_loop int64) gopurs_runtime.Value {
Emit__1834327443:
	for {
		if false {
			continue Emit__1834327443
		}
		var __eta_norm_1_0 []int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer(Rebox_Main_2837667778_3860057311((&Constructor_Main_Emit[[]int64]{1, __eta_norm_1_0, __eta_norm_0_1})))}
	}
}

func Call_Main_Emit__1315694163(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop int64) gopurs_runtime.Value {
Emit__1315694163:
	for {
		if false {
			continue Emit__1315694163
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer(Rebox_Main_1033515812_3860057311((&Constructor_Main_Emit[int64]{1, __eta_norm_1_0, __eta_norm_0_1})))}
	}
}

func Call_Main_Emit__203891187(__eta_norm_1_0_loop string, __eta_norm_0_1_loop int64) gopurs_runtime.Value {
Emit__203891187:
	for {
		if false {
			continue Emit__203891187
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Value{Type: 9, IntVal: 621001451, UnsafePtr: unsafe.Pointer(Rebox_Main_452751936_3860057311((&Constructor_Main_Emit[string]{1, __eta_norm_1_0, __eta_norm_0_1})))}
	}
}

func Call_Main_Stop__1606680011(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
Stop__1606680011:
	for {
		if false {
			continue Stop__1606680011
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer(Rebox_Main_4180764687_2803447442((&Constructor_Main_Stop[[]int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Stop__2802585741(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
Stop__2802585741:
	for {
		if false {
			continue Stop__2802585741
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer(Rebox_Main_2491663465_2803447442((&Constructor_Main_Stop[int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Stop__916476971(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
Stop__916476971:
	for {
		if false {
			continue Stop__916476971
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 325197958, UnsafePtr: unsafe.Pointer(Rebox_Main_4258924941_2803447442((&Constructor_Main_Stop[string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Failed__2802585741(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
Failed__2802585741:
	for {
		if false {
			continue Failed__2802585741
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 270449853, UnsafePtr: unsafe.Pointer(Rebox_Main_2585605810_280377929((&Constructor_Main_Failed[int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_snapshotRead(initial_0_loop int64) gopurs_runtime.Value {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(initial_0))
		_ = __local_var_1_0
		cell_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = cell_2_1
		value_3_2 := gopurs_runtime.ReadSTRef(cell_2_1)
		_ = value_3_2
		_dollar___unused_4_3 := gopurs_runtime.WriteSTRef(cell_2_1, gopurs_runtime.Int(gopurs_runtime.IntAdd(initial_0, int64(1))))
		_ = _dollar___unused_4_3
		return gopurs_runtime.Func(func(offset_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_3_2.IntVal, offset_5.IntVal))
		})
	}))
}

func Call_Main_skipLoop(initial_0_loop int64) int64 {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(initial_0))
		_ = __local_var_1_0
		cell_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = cell_2_1
		_dollar___unused_3_2 := gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(false)
		}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			current_3_3 := gopurs_runtime.ReadSTRef(cell_2_1)
			_ = current_3_3
			__local_var_4_4 := gopurs_runtime.WriteSTRef(cell_2_1, gopurs_runtime.Int(gopurs_runtime.IntAdd(current_3_3.IntVal, int64(1))))
			_ = __local_var_4_4
			return Get_Data_Unit_unit()
		}), gopurs_runtime.Value{})
		_ = _dollar___unused_3_2
		return gopurs_runtime.ReadSTRef(cell_2_1)
	})).IntVal
}

func Call_Main_operationResults(initial_0_loop int64) []int64 {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
			__local_var_1_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(initial_0))
			_ = __local_var_1_0
			cell_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
			_ = cell_2_1
			values_3_2 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
			_ = values_3_2
			assigned_4_3 := gopurs_runtime.WriteSTRef(cell_2_1, gopurs_runtime.Int(gopurs_runtime.IntAdd(initial_0, int64(1))))
			_ = assigned_4_3
			count_5_4 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), assigned_4_3, values_3_2)
			_ = count_5_4
			current_6_5 := gopurs_runtime.ReadSTRef(cell_2_1)
			_ = current_6_5
			_dollar___unused_7_6 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Int(gopurs_runtime.IntAdd(current_6_5.IntVal, count_5_4.IntVal)), values_3_2)
			_ = _dollar___unused_7_6
			return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), values_3_2)
		})).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Main_freezeSnapshot(initial_0_loop int64) struct {
	after  []int64
	before []int64
	/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
} {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	return func() struct {
		after  []int64
		before []int64
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	} {
		orig := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			values_1_0 := gopurs_runtime.Apply(Get_Data_Array_ST_go__new(), gopurs_runtime.Value{})
			_ = values_1_0
			_dollar___unused_2_1 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Int(initial_0), values_1_0)
			_ = _dollar___unused_2_1
			before_3_2 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_freezeImpl(), values_1_0)
			_ = before_3_2
			_dollar___unused_4_3 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Int(gopurs_runtime.IntAdd(initial_0, int64(1))), values_1_0)
			_ = _dollar___unused_4_3
			after_5_4 := gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), values_1_0)
			_ = after_5_4
			return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_1, after_5_4, before_3_2)
		}))
		_ = orig
		clone := struct {
			after  []int64
			before []int64
			/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
		}{}
		clone.after = func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "after").UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()
		clone.before = func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "before").UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()
		return clone
	}()
}

func Call_Main_collect(step_0_loop gopurs_runtime.Value, start_1_loop int64) struct {
	failed bool
	next   int64
	values []gopurs_runtime.Value
	/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
} {
	var step_0 gopurs_runtime.Value = step_0_loop
	_ = step_0
	var start_1 int64 = start_1_loop
	_ = start_1
	__st_array0 := []gopurs_runtime.Value{}
	_ = __st_array0
	__st_ref1 := int64(start_1)
	_ = __st_ref1
	__st_ref2 := true
	_ = __st_ref2
	__st_ref3 := func() gopurs_runtime.Value {
		orig := struct {
			failed bool
			next   int64
			values []gopurs_runtime.Value
			/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
		}{false, start_1, (*((*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)))}
		_ = orig
		return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), gopurs_runtime.Array(orig.values))
	}()
	_ = __st_ref3
	for {
		if !(__st_ref2) {
			break
		} else {

		}
		__st_value4 := int64(__st_ref1)
		_ = __st_value4
		__st_value5 := gopurs_runtime.Apply(step_0, gopurs_runtime.Int(__st_value4))
		_ = __st_value5
		var __st_branch6 gopurs_runtime.Value
		if __st_value5.Type == 9 && __st_value5.IntVal == 621001451 {
			__st_array0 = append(__st_array0, (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0)
			__st_value13 := int64(int64(len(__st_array0)))
			_ = __st_value13
			__st_ref1 = (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V1
			__st_value14 := int64(__st_ref1)
			_ = __st_value14
			__st_branch6 = Get_Data_Unit_unit()
		} else {
			if __st_value5.Type == 9 && __st_value5.IntVal == 325197958 {
				__st_frozen9 := make([]gopurs_runtime.Value, len(__st_array0))
				_ = __st_frozen9
				copy(__st_frozen9, __st_array0)
				__st_value10 := __st_frozen9
				_ = __st_value10
				__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(false), gopurs_runtime.Int((*Constructor_Main_Stop[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), gopurs_runtime.Array(__st_value10))
				__st_value11 := __st_ref3
				_ = __st_value11
				__st_ref2 = false
				__st_value12 := __st_ref2
				_ = __st_value12
				__st_branch6 = Get_Data_Unit_unit()
			} else {
				if __st_value5.Type == 9 && __st_value5.IntVal == 270449853 {
					__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(true), gopurs_runtime.Int((*Constructor_Main_Failed[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), gopurs_runtime.Array([]gopurs_runtime.Value{}))
					__st_value7 := __st_ref3
					_ = __st_value7
					__st_ref2 = false
					__st_value8 := __st_ref2
					_ = __st_value8
					__st_branch6 = Get_Data_Unit_unit()
				} else {
					__st_branch6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			}
		}
		_ = __st_branch6
	}
	__st_value15 := gopurs_runtime.Value{}
	_ = __st_value15
	return func() struct {
		failed bool
		next   int64
		values []gopurs_runtime.Value
		/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
	} {
		orig := __st_ref3
		_ = orig
		clone := struct {
			failed bool
			next   int64
			values []gopurs_runtime.Value
			/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
		}{}
		clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
		clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
		clone.values = (*((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)))
		return clone
	}()
}

func Call_Main_collect__2173325217(step_0_loop gopurs_runtime.Value, start_1_loop int64) struct {
	failed bool
	next   int64
	values [][]int64
	/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
} {
collect__2173325217:
	for {
		if false {
			continue collect__2173325217
		}
		var step_0 gopurs_runtime.Value = step_0_loop
		_ = step_0
		var start_1 int64 = start_1_loop
		_ = start_1
		__st_array0 := []gopurs_runtime.Value{}
		_ = __st_array0
		__st_ref1 := int64(start_1)
		_ = __st_ref1
		__st_ref2 := true
		_ = __st_ref2
		__st_ref3 := func() gopurs_runtime.Value {
			orig := struct {
				failed bool
				next   int64
				values [][]int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, start_1, [][]int64{}}
			_ = orig
			return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
				arr := orig.values
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
			}())
		}()
		_ = __st_ref3
		for {
			if !(__st_ref2) {
				break
			} else {

			}
			__st_value4 := int64(__st_ref1)
			_ = __st_value4
			__st_value5 := gopurs_runtime.Apply(step_0, gopurs_runtime.Int(__st_value4))
			_ = __st_value5
			var __st_branch6 gopurs_runtime.Value
			if __st_value5.Type == 9 && __st_value5.IntVal == 621001451 {
				__st_array0 = append(__st_array0, (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0)
				__st_value13 := int64(int64(len(__st_array0)))
				_ = __st_value13
				__st_ref1 = (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V1
				__st_value14 := int64(__st_ref1)
				_ = __st_value14
				__st_branch6 = Get_Data_Unit_unit()
			} else {
				if __st_value5.Type == 9 && __st_value5.IntVal == 325197958 {
					__st_frozen9 := make([]gopurs_runtime.Value, len(__st_array0))
					_ = __st_frozen9
					copy(__st_frozen9, __st_array0)
					__st_value10 := __st_frozen9
					_ = __st_value10
					__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(false), gopurs_runtime.Int((*Constructor_Main_Stop[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), gopurs_runtime.Array(__st_value10))
					__st_value11 := __st_ref3
					_ = __st_value11
					__st_ref2 = false
					__st_value12 := __st_ref2
					_ = __st_value12
					__st_branch6 = Get_Data_Unit_unit()
				} else {
					if __st_value5.Type == 9 && __st_value5.IntVal == 270449853 {
						__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(true), gopurs_runtime.Int((*Constructor_Main_Failed[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), func() gopurs_runtime.Value {
							arr := [][]int64{}
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
						}())
						__st_value7 := __st_ref3
						_ = __st_value7
						__st_ref2 = false
						__st_value8 := __st_ref2
						_ = __st_value8
						__st_branch6 = Get_Data_Unit_unit()
					} else {
						__st_branch6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				}
			}
			_ = __st_branch6
		}
		__st_value15 := gopurs_runtime.Value{}
		_ = __st_value15
		return func() struct {
			failed bool
			next   int64
			values [][]int64
			/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
		} {
			orig := __st_ref3
			_ = orig
			clone := struct {
				failed bool
				next   int64
				values [][]int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{}
			clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
			clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
			clone.values = func() [][]int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
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
			}()
			return clone
		}()
	}
}

func Call_Main_collect__110832033(step_0_loop gopurs_runtime.Value, start_1_loop int64) struct {
	failed bool
	next   int64
	values []int64
	/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
} {
collect__110832033:
	for {
		if false {
			continue collect__110832033
		}
		var step_0 gopurs_runtime.Value = step_0_loop
		_ = step_0
		var start_1 int64 = start_1_loop
		_ = start_1
		__st_array0 := []gopurs_runtime.Value{}
		_ = __st_array0
		__st_ref1 := int64(start_1)
		_ = __st_ref1
		__st_ref2 := true
		_ = __st_ref2
		__st_ref3 := func() gopurs_runtime.Value {
			orig := struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, start_1, []int64{}}
			_ = orig
			return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
				arr := orig.values
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())
		}()
		_ = __st_ref3
		for {
			if !(__st_ref2) {
				break
			} else {

			}
			__st_value4 := int64(__st_ref1)
			_ = __st_value4
			__st_value5 := gopurs_runtime.Apply(step_0, gopurs_runtime.Int(__st_value4))
			_ = __st_value5
			var __st_branch6 gopurs_runtime.Value
			if __st_value5.Type == 9 && __st_value5.IntVal == 621001451 {
				__st_array0 = append(__st_array0, (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0)
				__st_value13 := int64(int64(len(__st_array0)))
				_ = __st_value13
				__st_ref1 = (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V1
				__st_value14 := int64(__st_ref1)
				_ = __st_value14
				__st_branch6 = Get_Data_Unit_unit()
			} else {
				if __st_value5.Type == 9 && __st_value5.IntVal == 325197958 {
					__st_frozen9 := make([]gopurs_runtime.Value, len(__st_array0))
					_ = __st_frozen9
					copy(__st_frozen9, __st_array0)
					__st_value10 := __st_frozen9
					_ = __st_value10
					__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(false), gopurs_runtime.Int((*Constructor_Main_Stop[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), gopurs_runtime.Array(__st_value10))
					__st_value11 := __st_ref3
					_ = __st_value11
					__st_ref2 = false
					__st_value12 := __st_ref2
					_ = __st_value12
					__st_branch6 = Get_Data_Unit_unit()
				} else {
					if __st_value5.Type == 9 && __st_value5.IntVal == 270449853 {
						__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(true), gopurs_runtime.Int((*Constructor_Main_Failed[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), func() gopurs_runtime.Value {
							arr := []int64{}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}())
						__st_value7 := __st_ref3
						_ = __st_value7
						__st_ref2 = false
						__st_value8 := __st_ref2
						_ = __st_value8
						__st_branch6 = Get_Data_Unit_unit()
					} else {
						__st_branch6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				}
			}
			_ = __st_branch6
		}
		__st_value15 := gopurs_runtime.Value{}
		_ = __st_value15
		return func() struct {
			failed bool
			next   int64
			values []int64
			/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
		} {
			orig := __st_ref3
			_ = orig
			clone := struct {
				failed bool
				next   int64
				values []int64
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{}
			clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
			clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
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
}

func Call_Main_collect__1788204577(step_0_loop gopurs_runtime.Value, start_1_loop int64) struct {
	failed bool
	next   int64
	values []string
	/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
} {
collect__1788204577:
	for {
		if false {
			continue collect__1788204577
		}
		var step_0 gopurs_runtime.Value = step_0_loop
		_ = step_0
		var start_1 int64 = start_1_loop
		_ = start_1
		__st_array0 := []gopurs_runtime.Value{}
		_ = __st_array0
		__st_ref1 := int64(start_1)
		_ = __st_ref1
		__st_ref2 := true
		_ = __st_ref2
		__st_ref3 := func() gopurs_runtime.Value {
			orig := struct {
				failed bool
				next   int64
				values []string
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{false, start_1, []string{}}
			_ = orig
			return gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(orig.failed), gopurs_runtime.Int(orig.next), func() gopurs_runtime.Value {
				arr := orig.values
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}())
		}()
		_ = __st_ref3
		for {
			if !(__st_ref2) {
				break
			} else {

			}
			__st_value4 := int64(__st_ref1)
			_ = __st_value4
			__st_value5 := gopurs_runtime.Apply(step_0, gopurs_runtime.Int(__st_value4))
			_ = __st_value5
			var __st_branch6 gopurs_runtime.Value
			if __st_value5.Type == 9 && __st_value5.IntVal == 621001451 {
				__st_array0 = append(__st_array0, (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0)
				__st_value13 := int64(int64(len(__st_array0)))
				_ = __st_value13
				__st_ref1 = (*Constructor_Main_Emit[gopurs_runtime.Value])(__st_value5.UnsafePtr).V1
				__st_value14 := int64(__st_ref1)
				_ = __st_value14
				__st_branch6 = Get_Data_Unit_unit()
			} else {
				if __st_value5.Type == 9 && __st_value5.IntVal == 325197958 {
					__st_frozen9 := make([]gopurs_runtime.Value, len(__st_array0))
					_ = __st_frozen9
					copy(__st_frozen9, __st_array0)
					__st_value10 := __st_frozen9
					_ = __st_value10
					__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(false), gopurs_runtime.Int((*Constructor_Main_Stop[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), gopurs_runtime.Array(__st_value10))
					__st_value11 := __st_ref3
					_ = __st_value11
					__st_ref2 = false
					__st_value12 := __st_ref2
					_ = __st_value12
					__st_branch6 = Get_Data_Unit_unit()
				} else {
					if __st_value5.Type == 9 && __st_value5.IntVal == 270449853 {
						__st_ref3 = gopurs_runtime.RecordShape3(&GopursRecordKeys_4d61696e_2, gopurs_runtime.Bool(true), gopurs_runtime.Int((*Constructor_Main_Failed[gopurs_runtime.Value])(__st_value5.UnsafePtr).V0), func() gopurs_runtime.Value {
							arr := []string{}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Str(v)
							}
							return gopurs_runtime.Array(boxed)
						}())
						__st_value7 := __st_ref3
						_ = __st_value7
						__st_ref2 = false
						__st_value8 := __st_ref2
						_ = __st_value8
						__st_branch6 = Get_Data_Unit_unit()
					} else {
						__st_branch6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				}
			}
			_ = __st_branch6
		}
		__st_value15 := gopurs_runtime.Value{}
		_ = __st_value15
		return func() struct {
			failed bool
			next   int64
			values []string
			/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
		} {
			orig := __st_ref3
			_ = orig
			clone := struct {
				failed bool
				next   int64
				values []string
				/*gopurs:record-v1:3:006600610069006c00650064,006e006500780074,00760061006c007500650073*/
			}{}
			clone.failed = (gopurs_runtime.RecordGet(orig, "failed").IntVal) != (0)
			clone.next = gopurs_runtime.RecordGet(orig, "next").IntVal
			clone.values = func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}()
			return clone
		}()
	}
}

func Call_Main_aliasRead(initial_0_loop int64) int64 {
	var initial_0 int64 = initial_0_loop
	_ = initial_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(initial_0))
		_ = __local_var_1_0
		cell_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = cell_2_1
		_dollar___unused_3_2 := gopurs_runtime.WriteSTRef(cell_2_1, gopurs_runtime.Int(gopurs_runtime.IntAdd(initial_0, int64(1))))
		_ = _dollar___unused_3_2
		return gopurs_runtime.ReadSTRef(cell_2_1)
	})).IntVal
}

func Rebox_Main_1033515812_3860057311(in *Constructor_Main_Emit[int64]) *Constructor_Main_Emit[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Emit[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = in.V1
	return out
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

func Rebox_Main_2491663465_2803447442(in *Constructor_Main_Stop[int64]) *Constructor_Main_Stop[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Stop[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2585605810_280377929(in *Constructor_Main_Failed[int64]) *Constructor_Main_Failed[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Failed[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2837667778_3860057311(in *Constructor_Main_Emit[[]int64]) *Constructor_Main_Emit[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Emit[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		arr := in.V0
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Int(v)
		}
		return gopurs_runtime.Array(boxed)
	}()
	out.V1 = in.V1
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_4180764687_2803447442(in *Constructor_Main_Stop[[]int64]) *Constructor_Main_Stop[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Stop[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4258924941_2803447442(in *Constructor_Main_Stop[string]) *Constructor_Main_Stop[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Stop[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_452751936_3860057311(in *Constructor_Main_Emit[string]) *Constructor_Main_Emit[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Emit[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys1{K0: "reflectSymbol"}

var GopursRecordKeys_4d61696e_1 = gopurs_runtime.RecordKeys2{K0: "after", K1: "before"}

var GopursRecordKeys_4d61696e_2 = gopurs_runtime.RecordKeys3{K0: "failed", K1: "next", K2: "values"}
