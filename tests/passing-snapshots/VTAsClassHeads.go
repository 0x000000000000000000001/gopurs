package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Superclass_dollar_Dict gopurs_runtime.Value
var once_Main_Superclass_dollar_Dict sync.Once

func Get_Main_Superclass_dollar_Dict() gopurs_runtime.Value {
	once_Main_Superclass_dollar_Dict.Do(func() {
		cache_Main_Superclass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Superclass_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Superclass_dollar_Dict
}

var cache_Main_Singleton_dollar_Dict gopurs_runtime.Value
var once_Main_Singleton_dollar_Dict sync.Once

func Get_Main_Singleton_dollar_Dict() gopurs_runtime.Value {
	once_Main_Singleton_dollar_Dict.Do(func() {
		cache_Main_Singleton_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Singleton_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Singleton_dollar_Dict
}

var cache_Main_MultiWithFDs_dollar_Dict gopurs_runtime.Value
var once_Main_MultiWithFDs_dollar_Dict sync.Once

func Get_Main_MultiWithFDs_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultiWithFDs_dollar_Dict.Do(func() {
		cache_Main_MultiWithFDs_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultiWithFDs_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultiWithFDs_dollar_Dict
}

var cache_Main_MultiWithBidiFDs_dollar_Dict gopurs_runtime.Value
var once_Main_MultiWithBidiFDs_dollar_Dict sync.Once

func Get_Main_MultiWithBidiFDs_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultiWithBidiFDs_dollar_Dict.Do(func() {
		cache_Main_MultiWithBidiFDs_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultiWithBidiFDs_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultiWithBidiFDs_dollar_Dict
}

var cache_Main_MultiNoFDs_dollar_Dict gopurs_runtime.Value
var once_Main_MultiNoFDs_dollar_Dict sync.Once

func Get_Main_MultiNoFDs_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultiNoFDs_dollar_Dict.Do(func() {
		cache_Main_MultiNoFDs_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultiNoFDs_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultiNoFDs_dollar_Dict
}

var cache_Main_MultiCoveringSets_dollar_Dict gopurs_runtime.Value
var once_Main_MultiCoveringSets_dollar_Dict sync.Once

func Get_Main_MultiCoveringSets_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultiCoveringSets_dollar_Dict.Do(func() {
		cache_Main_MultiCoveringSets_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultiCoveringSets_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultiCoveringSets_dollar_Dict
}

var cache_Main_MainClass_dollar_Dict gopurs_runtime.Value
var once_Main_MainClass_dollar_Dict sync.Once

func Get_Main_MainClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_MainClass_dollar_Dict.Do(func() {
		cache_Main_MainClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MainClass_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MainClass_dollar_Dict
}

var cache_Main_ConflictingIdentSynonym_dollar_Dict gopurs_runtime.Value
var once_Main_ConflictingIdentSynonym_dollar_Dict sync.Once

func Get_Main_ConflictingIdentSynonym_dollar_Dict() gopurs_runtime.Value {
	once_Main_ConflictingIdentSynonym_dollar_Dict.Do(func() {
		cache_Main_ConflictingIdentSynonym_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ConflictingIdentSynonym_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ConflictingIdentSynonym_dollar_Dict
}

var cache_Main_ConflictingIdent_dollar_Dict gopurs_runtime.Value
var once_Main_ConflictingIdent_dollar_Dict sync.Once

func Get_Main_ConflictingIdent_dollar_Dict() gopurs_runtime.Value {
	once_Main_ConflictingIdent_dollar_Dict.Do(func() {
		cache_Main_ConflictingIdent_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ConflictingIdent_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ConflictingIdent_dollar_Dict
}

var cache_Main_B2 gopurs_runtime.Value
var once_Main_B2 sync.Once

func Get_Main_B2() gopurs_runtime.Value {
	once_Main_B2.Do(func() {
		cache_Main_B2 = gopurs_runtime.Value{Type: 9, IntVal: int64(4102534158), UnsafePtr: nil}
	})
	return cache_Main_B2
}

var cache_Main_A2 gopurs_runtime.Value
var once_Main_A2 sync.Once

func Get_Main_A2() gopurs_runtime.Value {
	once_Main_A2.Do(func() {
		cache_Main_A2 = gopurs_runtime.Value{Type: 9, IntVal: int64(2540403533), UnsafePtr: nil}
	})
	return cache_Main_A2
}

var cache_Main_superclassB2 gopurs_runtime.Value
var once_Main_superclassB2 sync.Once

func Get_Main_superclassB2() gopurs_runtime.Value {
	once_Main_superclassB2.Do(func() {
		cache_Main_superclassB2 = gopurs_runtime.Value{Type: 9, IntVal: 3942833201, UnsafePtr: unsafe.Pointer((&Constructor_Main_Superclass[uint32]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4102534158), UnsafePtr: nil}}))}
	})
	return cache_Main_superclassB2
}

var cache_Main_superclassA2 gopurs_runtime.Value
var once_Main_superclassA2 sync.Once

func Get_Main_superclassA2() gopurs_runtime.Value {
	once_Main_superclassA2.Do(func() {
		cache_Main_superclassA2 = gopurs_runtime.Value{Type: 9, IntVal: 3942833201, UnsafePtr: unsafe.Pointer((&Constructor_Main_Superclass[uint32]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2540403533), UnsafePtr: nil}}))}
	})
	return cache_Main_superclassA2
}

var cache_Main_singletonString gopurs_runtime.Value
var once_Main_singletonString sync.Once

func Get_Main_singletonString() gopurs_runtime.Value {
	once_Main_singletonString.Do(func() {
		cache_Main_singletonString = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer((&Constructor_Main_Singleton[string]{1, "string"}))}
	})
	return cache_Main_singletonString
}

var cache_Main_singletonInt gopurs_runtime.Value
var once_Main_singletonInt sync.Once

func Get_Main_singletonInt() gopurs_runtime.Value {
	once_Main_singletonInt.Do(func() {
		cache_Main_singletonInt = gopurs_runtime.Value{Type: 9, IntVal: 3773262993, UnsafePtr: unsafe.Pointer((&Constructor_Main_Singleton[int64]{1, "int"}))}
	})
	return cache_Main_singletonInt
}

var cache_Main_multiWithFDsStringInt gopurs_runtime.Value
var once_Main_multiWithFDsStringInt sync.Once

func Get_Main_multiWithFDsStringInt() gopurs_runtime.Value {
	once_Main_multiWithFDsStringInt.Do(func() {
		cache_Main_multiWithFDsStringInt = gopurs_runtime.Value{Type: 9, IntVal: 4172703812, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiWithFDs[string, int64]{1, 1}))}
	})
	return cache_Main_multiWithFDsStringInt
}

var cache_Main_multiWithFDsIntInt gopurs_runtime.Value
var once_Main_multiWithFDsIntInt sync.Once

func Get_Main_multiWithFDsIntInt() gopurs_runtime.Value {
	once_Main_multiWithFDsIntInt.Do(func() {
		cache_Main_multiWithFDsIntInt = gopurs_runtime.Value{Type: 9, IntVal: 4172703812, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiWithFDs[int64, int64]{1, 0}))}
	})
	return cache_Main_multiWithFDsIntInt
}

var cache_Main_multiWithBidiFDsStringStr gopurs_runtime.Value
var once_Main_multiWithBidiFDsStringStr sync.Once

func Get_Main_multiWithBidiFDsStringStr() gopurs_runtime.Value {
	once_Main_multiWithBidiFDsStringStr.Do(func() {
		cache_Main_multiWithBidiFDsStringStr = gopurs_runtime.Value{Type: 9, IntVal: 3391744610, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiWithBidiFDs[string, string]{1, 1}))}
	})
	return cache_Main_multiWithBidiFDsStringStr
}

var cache_Main_multiWithBidiFDsIntInt gopurs_runtime.Value
var once_Main_multiWithBidiFDsIntInt sync.Once

func Get_Main_multiWithBidiFDsIntInt() gopurs_runtime.Value {
	once_Main_multiWithBidiFDsIntInt.Do(func() {
		cache_Main_multiWithBidiFDsIntInt = gopurs_runtime.Value{Type: 9, IntVal: 3391744610, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiWithBidiFDs[int64, int64]{1, 0}))}
	})
	return cache_Main_multiWithBidiFDsIntInt
}

var cache_Main_multiNoFDsStringInt gopurs_runtime.Value
var once_Main_multiNoFDsStringInt sync.Once

func Get_Main_multiNoFDsStringInt() gopurs_runtime.Value {
	once_Main_multiNoFDsStringInt.Do(func() {
		cache_Main_multiNoFDsStringInt = gopurs_runtime.Value{Type: 9, IntVal: 2354658663, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiNoFDs[string, int64]{1, 1}))}
	})
	return cache_Main_multiNoFDsStringInt
}

var cache_Main_multiNoFDsIntInt gopurs_runtime.Value
var once_Main_multiNoFDsIntInt sync.Once

func Get_Main_multiNoFDsIntInt() gopurs_runtime.Value {
	once_Main_multiNoFDsIntInt.Do(func() {
		cache_Main_multiNoFDsIntInt = gopurs_runtime.Value{Type: 9, IntVal: 2354658663, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiNoFDs[int64, int64]{1, 0}))}
	})
	return cache_Main_multiNoFDsIntInt
}

var cache_Main_multiCoveringSetsIntIntSt gopurs_runtime.Value
var once_Main_multiCoveringSetsIntIntSt sync.Once

func Get_Main_multiCoveringSetsIntIntSt() gopurs_runtime.Value {
	once_Main_multiCoveringSetsIntIntSt.Do(func() {
		cache_Main_multiCoveringSetsIntIntSt = gopurs_runtime.Value{Type: 9, IntVal: 2410351339, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiCoveringSets[int64, int64, string, string, bool, bool]{1, 2, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0.IntVal)).StrVal()), gopurs_runtime.Str("2"))
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 string
			{
				if (f_0.IntVal) != (0) {
					__t0 = "true"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "false"
			}
		end_branch_0:
			return gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Str(__t0), gopurs_runtime.Str("2"))
		})}))}
	})
	return cache_Main_multiCoveringSetsIntIntSt
}

var cache_Main_multiCoveringSetsBooleanB gopurs_runtime.Value
var once_Main_multiCoveringSetsBooleanB sync.Once

func Get_Main_multiCoveringSetsBooleanB() gopurs_runtime.Value {
	once_Main_multiCoveringSetsBooleanB.Do(func() {
		cache_Main_multiCoveringSetsBooleanB = gopurs_runtime.Value{Type: 9, IntVal: 2410351339, UnsafePtr: unsafe.Pointer((&Constructor_Main_MultiCoveringSets[bool, bool, string, string, int64, int64]{1, 1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 string
			{
				if (a_0.IntVal) != (0) {
					__t0 = "101"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "100"
			}
		end_branch_0:
			return gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Str(__t0), gopurs_runtime.Str("1"))
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(f_0.IntVal)).StrVal()), gopurs_runtime.Str("1"))
		})}))}
	})
	return cache_Main_multiCoveringSetsBooleanB
}

var cache_Main_mainClassB2 gopurs_runtime.Value
var once_Main_mainClassB2 sync.Once

func Get_Main_mainClassB2() gopurs_runtime.Value {
	once_Main_mainClassB2.Do(func() {
		cache_Main_mainClassB2 = gopurs_runtime.Value{Type: 9, IntVal: 1673006587, UnsafePtr: unsafe.Pointer((&Constructor_Main_MainClass[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3942833201, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Superclass[uint32]](Get_Main_superclassB2()))}
		}), 3}))}
	})
	return cache_Main_mainClassB2
}

var cache_Main_mainClassA2 gopurs_runtime.Value
var once_Main_mainClassA2 sync.Once

func Get_Main_mainClassA2() gopurs_runtime.Value {
	once_Main_mainClassA2.Do(func() {
		cache_Main_mainClassA2 = gopurs_runtime.Value{Type: 9, IntVal: 1673006587, UnsafePtr: unsafe.Pointer((&Constructor_Main_MainClass[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3942833201, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Superclass[uint32]](Get_Main_superclassA2()))}
		}), 0}))}
	})
	return cache_Main_mainClassA2
}

var cache_Main_eqB2 gopurs_runtime.Value
var once_Main_eqB2 sync.Once

func Get_Main_eqB2() gopurs_runtime.Value {
	once_Main_eqB2.Do(func() {
		cache_Main_eqB2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})}))}
	})
	return cache_Main_eqB2
}

var cache_Main_eqA2 gopurs_runtime.Value
var once_Main_eqA2 sync.Once

func Get_Main_eqA2() gopurs_runtime.Value {
	once_Main_eqA2.Do(func() {
		cache_Main_eqA2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})}))}
	})
	return cache_Main_eqA2
}

var cache_Main_conflictingIdentSynonymSt gopurs_runtime.Value
var once_Main_conflictingIdentSynonymSt sync.Once

func Get_Main_conflictingIdentSynonymSt() gopurs_runtime.Value {
	once_Main_conflictingIdentSynonymSt.Do(func() {
		cache_Main_conflictingIdentSynonymSt = gopurs_runtime.Value{Type: 9, IntVal: 3172630699, UnsafePtr: unsafe.Pointer((&Constructor_Main_ConflictingIdentSynonym[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(1)
		})}))}
	})
	return cache_Main_conflictingIdentSynonymSt
}

var cache_Main_conflictingIdentSynonymIn gopurs_runtime.Value
var once_Main_conflictingIdentSynonymIn sync.Once

func Get_Main_conflictingIdentSynonymIn() gopurs_runtime.Value {
	once_Main_conflictingIdentSynonymIn.Do(func() {
		cache_Main_conflictingIdentSynonymIn = gopurs_runtime.Value{Type: 9, IntVal: 3172630699, UnsafePtr: unsafe.Pointer((&Constructor_Main_ConflictingIdentSynonym[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(2)
		})}))}
	})
	return cache_Main_conflictingIdentSynonymIn
}

var cache_Main_conflictingIdentString gopurs_runtime.Value
var once_Main_conflictingIdentString sync.Once

func Get_Main_conflictingIdentString() gopurs_runtime.Value {
	once_Main_conflictingIdentString.Do(func() {
		cache_Main_conflictingIdentString = gopurs_runtime.Value{Type: 9, IntVal: 1296752506, UnsafePtr: unsafe.Pointer((&Constructor_Main_ConflictingIdent[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(1)
		})}))}
	})
	return cache_Main_conflictingIdentString
}

var cache_Main_conflictingIdentInt gopurs_runtime.Value
var once_Main_conflictingIdentInt sync.Once

func Get_Main_conflictingIdentInt() gopurs_runtime.Value {
	once_Main_conflictingIdentInt.Do(func() {
		cache_Main_conflictingIdentInt = gopurs_runtime.Value{Type: 9, IntVal: 1296752506, UnsafePtr: unsafe.Pointer((&Constructor_Main_ConflictingIdent[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(2)
		})}))}
	})
	return cache_Main_conflictingIdentInt
}

var cache_Main_superClassValue gopurs_runtime.Value
var once_Main_superClassValue sync.Once

func Get_Main_superClassValue() gopurs_runtime.Value {
	once_Main_superClassValue.Do(func() {
		cache_Main_superClassValue = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_superClassValue(dict_0_box)
		})
	})
	return cache_Main_superClassValue
}

var cache_Main_singleton gopurs_runtime.Value
var once_Main_singleton sync.Once

func Get_Main_singleton() gopurs_runtime.Value {
	once_Main_singleton.Do(func() {
		cache_Main_singleton = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_singleton(dict_0_box)
		})
	})
	return cache_Main_singleton
}

var cache_Main_singletonWorks gopurs_runtime.Value
var once_Main_singletonWorks sync.Once

func Get_Main_singletonWorks() gopurs_runtime.Value {
	once_Main_singletonWorks.Do(func() {
		cache_Main_singletonWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_singletonWorks
}

var cache_Main_partialOfFESet gopurs_runtime.Value
var once_Main_partialOfFESet sync.Once

func Get_Main_partialOfFESet() gopurs_runtime.Value {
	once_Main_partialOfFESet.Do(func() {
		cache_Main_partialOfFESet = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_partialOfFESet(gopurs_runtime.CoerceToStruct[Constructor_Main_MultiCoveringSets[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_partialOfFESet
}

var cache_Main_partialOfABSet gopurs_runtime.Value
var once_Main_partialOfABSet sync.Once

func Get_Main_partialOfABSet() gopurs_runtime.Value {
	once_Main_partialOfABSet.Do(func() {
		cache_Main_partialOfABSet = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_partialOfABSet(gopurs_runtime.CoerceToStruct[Constructor_Main_MultiCoveringSets[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_partialOfABSet
}

var cache_Main_noneOfSets gopurs_runtime.Value
var once_Main_noneOfSets sync.Once

func Get_Main_noneOfSets() gopurs_runtime.Value {
	once_Main_noneOfSets.Do(func() {
		cache_Main_noneOfSets = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_noneOfSets(dict_0_box)
		})
	})
	return cache_Main_noneOfSets
}

var cache_Main_multiWithFDs gopurs_runtime.Value
var once_Main_multiWithFDs sync.Once

func Get_Main_multiWithFDs() gopurs_runtime.Value {
	once_Main_multiWithFDs.Do(func() {
		cache_Main_multiWithFDs = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_multiWithFDs(dict_0_box)
		})
	})
	return cache_Main_multiWithFDs
}

var cache_Main_multiWithFdsWorks gopurs_runtime.Value
var once_Main_multiWithFdsWorks sync.Once

func Get_Main_multiWithFdsWorks() gopurs_runtime.Value {
	once_Main_multiWithFdsWorks.Do(func() {
		cache_Main_multiWithFdsWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_multiWithFdsWorks
}

var cache_Main_multiWithBidiFDs gopurs_runtime.Value
var once_Main_multiWithBidiFDs sync.Once

func Get_Main_multiWithBidiFDs() gopurs_runtime.Value {
	once_Main_multiWithBidiFDs.Do(func() {
		cache_Main_multiWithBidiFDs = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_multiWithBidiFDs(dict_0_box)
		})
	})
	return cache_Main_multiWithBidiFDs
}

var cache_Main_multiWithBidiFDsLeftWorks gopurs_runtime.Value
var once_Main_multiWithBidiFDsLeftWorks sync.Once

func Get_Main_multiWithBidiFDsLeftWorks() gopurs_runtime.Value {
	once_Main_multiWithBidiFDsLeftWorks.Do(func() {
		cache_Main_multiWithBidiFDsLeftWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_multiWithBidiFDsLeftWorks
}

var cache_Main_multiWithBidiFDsRightWorks gopurs_runtime.Value
var once_Main_multiWithBidiFDsRightWorks sync.Once

func Get_Main_multiWithBidiFDsRightWorks() gopurs_runtime.Value {
	once_Main_multiWithBidiFDsRightWorks.Do(func() {
		cache_Main_multiWithBidiFDsRightWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_multiWithBidiFDsRightWorks
}

var cache_Main_multiNoFds gopurs_runtime.Value
var once_Main_multiNoFds sync.Once

func Get_Main_multiNoFds() gopurs_runtime.Value {
	once_Main_multiNoFds.Do(func() {
		cache_Main_multiNoFds = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_multiNoFds(dict_0_box)
		})
	})
	return cache_Main_multiNoFds
}

var cache_Main_multiNoFdsWorks gopurs_runtime.Value
var once_Main_multiNoFdsWorks sync.Once

func Get_Main_multiNoFdsWorks() gopurs_runtime.Value {
	once_Main_multiNoFdsWorks.Do(func() {
		cache_Main_multiNoFdsWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_multiNoFdsWorks
}

var cache_Main_multiCoveringSetsWorks gopurs_runtime.Value
var once_Main_multiCoveringSetsWorks sync.Once

func Get_Main_multiCoveringSetsWorks() gopurs_runtime.Value {
	once_Main_multiCoveringSetsWorks.Do(func() {
		cache_Main_multiCoveringSetsWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if ((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(3)).StrVal()) == (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(3)).StrVal())) && ((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(20)).StrVal()) == (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(20)).StrVal())) {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[string]{1, gopurs_runtime.Str("MultiCoveringSets failed")}))}
			}
		end_branch_1:
			// TAST (Let): __local_var_0_0 shape=Branch(Other, def=Other) bindingType=Any
			__local_var_0_0 := __t1
			_ = __local_var_0_0
			return __local_var_0_0
		})
	})
	return cache_Main_multiCoveringSetsWorks
}

var cache_Main_mainClassInt gopurs_runtime.Value
var once_Main_mainClassInt sync.Once

func Get_Main_mainClassInt() gopurs_runtime.Value {
	once_Main_mainClassInt.Do(func() {
		cache_Main_mainClassInt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_mainClassInt(dict_0_box)
		})
	})
	return cache_Main_mainClassInt
}

var cache_Main_mainClassWorks gopurs_runtime.Value
var once_Main_mainClassWorks sync.Once

func Get_Main_mainClassWorks() gopurs_runtime.Value {
	once_Main_mainClassWorks.Do(func() {
		cache_Main_mainClassWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_mainClassWorks
}

var cache_Main_conflictingIdentSynonym gopurs_runtime.Value
var once_Main_conflictingIdentSynonym sync.Once

func Get_Main_conflictingIdentSynonym() gopurs_runtime.Value {
	once_Main_conflictingIdentSynonym.Do(func() {
		cache_Main_conflictingIdentSynonym = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_conflictingIdentSynonym(gopurs_runtime.CoerceToStruct[Constructor_Main_ConflictingIdentSynonym[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_conflictingIdentSynonym
}

var cache_Main_conflictingIdentSynonymWorks gopurs_runtime.Value
var once_Main_conflictingIdentSynonymWorks sync.Once

func Get_Main_conflictingIdentSynonymWorks() gopurs_runtime.Value {
	once_Main_conflictingIdentSynonymWorks.Do(func() {
		cache_Main_conflictingIdentSynonymWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_conflictingIdentSynonymWorks
}

var cache_Main_conflictingIdent gopurs_runtime.Value
var once_Main_conflictingIdent sync.Once

func Get_Main_conflictingIdent() gopurs_runtime.Value {
	once_Main_conflictingIdent.Do(func() {
		cache_Main_conflictingIdent = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_conflictingIdent(gopurs_runtime.CoerceToStruct[Constructor_Main_ConflictingIdent[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_conflictingIdent
}

var cache_Main_conflictingIdentWorks gopurs_runtime.Value
var once_Main_conflictingIdentWorks sync.Once

func Get_Main_conflictingIdentWorks() gopurs_runtime.Value {
	once_Main_conflictingIdentWorks.Do(func() {
		cache_Main_conflictingIdentWorks = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})
	})
	return cache_Main_conflictingIdentWorks
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(TypeApp (TypeVar m) [(Array (TypeVar b))])
			__local_var_0_0 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						a_prime__2_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
						_ = a_prime__2_1
						return gopurs_runtime.Apply(f_0, a_prime__2_1)
					})
				})
			}), Get_Effect_pureE(), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_0
			}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				})}).UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()))
			_ = __local_var_0_0
			arr_prime__1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = arr_prime__1_2
			// TAST (Let): __local_var_2_4 shape=App(Var) bindingType=Any
			__local_var_2_4 := gopurs_runtime.Apply2(Get_Data_Array_mapMaybe(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_2
			}), arr_prime__1_2)
			_ = __local_var_2_4
			var __t5 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
			{
				if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_2_4))).IntVal) > (0) {
					__t5 = (&Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]{1, __local_var_2_4})
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = (*Constructor_Data_Maybe_Just[[]gopurs_runtime.Value])(nil)
			}
		end_branch_5:
			// TAST (Let): v_2_3 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array String)])
			var v_2_3 *Constructor_Data_Maybe_Just[[]string] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]string]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})
			var __t7 gopurs_runtime.Value
			{
				if v_2_3 != nil {
					__t7 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("Errors...")+(gopurs_runtime.RecordGet(func() gopurs_runtime.Value {
						arr_val_foldlArray7 := func() gopurs_runtime.Value {
							arr := func() []string {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)((v_2_3).V0.UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()).UnsafePtr)
								unboxed := make([]string, len(arr))
								for i, v := range arr {
									unboxed[i] = v.StrVal()
								}
								return unboxed
							}()
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Str(v)
							}
							return gopurs_runtime.Array(boxed)
						}()
						_ = arr_val_foldlArray7
						res_go_foldlArray7 := gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(""), gopurs_runtime.Bool(true))
						_ = res_go_foldlArray7
						arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
						_ = arr_go_foldlArray7
						for _, v_foldlArray7 := range *arr_go_foldlArray7 {
							res_go_foldlArray7 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t6 gopurs_runtime.Value
									{
										if (gopurs_runtime.RecordGet(v_3, "init").IntVal) != (0) {
											__t6 = gopurs_runtime.RecordDict2("acc", "init", v1_4, gopurs_runtime.Bool(false))
											goto end_branch_6
										} else {

										}
									}
									{
										__t6 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(((gopurs_runtime.RecordGet(v_3, "acc").StrVal())+("\x0a"))+(v1_4.StrVal())), gopurs_runtime.Bool(false))
									}
								end_branch_6:
									return __t6
								})
							}), res_go_foldlArray7, v_foldlArray7)
						}
						return res_go_foldlArray7
					}(), "acc").StrVal())))
					goto end_branch_7
				} else {

				}
			}
			{
				if v_2_3 == nil {
					__t7 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_7:
			return gopurs_runtime.Apply(__t7, gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_B2 struct {
	Rc uint32
}

type Constructor_Main_A2 struct {
	Rc uint32
}

type Constructor_Main_Superclass[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3942833201] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Superclass[any])(ptr)
		_ = c
		switch key {
		case "superClassValue":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Superclass: " + key)
		}
	}
}

type Constructor_Main_Singleton[T_x any] struct {
	Rc uint32
	V0 string
}

func init() {
	gopurs_runtime.StructGetters[3773262993] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Singleton[any])(ptr)
		_ = c
		switch key {
		case "singleton":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Singleton: " + key)
		}
	}
}

type Constructor_Main_MultiWithFDs[T_a any, T_b any] struct {
	Rc uint32
	V0 int64
}

func init() {
	gopurs_runtime.StructGetters[4172703812] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultiWithFDs[any, any])(ptr)
		_ = c
		switch key {
		case "multiWithFDs":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_MultiWithFDs: " + key)
		}
	}
}

type Constructor_Main_MultiWithBidiFDs[T_a any, T_b any] struct {
	Rc uint32
	V0 int64
}

func init() {
	gopurs_runtime.StructGetters[3391744610] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultiWithBidiFDs[any, any])(ptr)
		_ = c
		switch key {
		case "multiWithBidiFDs":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_MultiWithBidiFDs: " + key)
		}
	}
}

type Constructor_Main_MultiNoFDs[T_a any, T_b any] struct {
	Rc uint32
	V0 int64
}

func init() {
	gopurs_runtime.StructGetters[2354658663] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultiNoFDs[any, any])(ptr)
		_ = c
		switch key {
		case "multiNoFds":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_MultiNoFDs: " + key)
		}
	}
}

type Constructor_Main_MultiCoveringSets[T_a any, T_b any, T_c any, T_d any, T_e any, T_f any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2410351339] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultiCoveringSets[any, any, any, any, any, any])(ptr)
		_ = c
		switch key {
		case "noneOfSets":
			return gopurs_runtime.Box(c.V0)
		case "partialOfABSet":
			return gopurs_runtime.Box(c.V1)
		case "partialOfFESet":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Main_MultiCoveringSets: " + key)
		}
	}
}

type Constructor_Main_MainClass[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 int64
}

func init() {
	gopurs_runtime.StructGetters[1673006587] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MainClass[any])(ptr)
		_ = c
		switch key {
		case "Superclass0":
			return gopurs_runtime.Box(c.V0)
		case "mainClassInt":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_MainClass: " + key)
		}
	}
}

type Constructor_Main_ConflictingIdentSynonym[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3172630699] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ConflictingIdentSynonym[any])(ptr)
		_ = c
		switch key {
		case "conflictingIdentSynonym":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ConflictingIdentSynonym: " + key)
		}
	}
}

type Constructor_Main_ConflictingIdent[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1296752506] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ConflictingIdent[any])(ptr)
		_ = c
		switch key {
		case "conflictingIdent":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ConflictingIdent: " + key)
		}
	}
}

func Call_Main_Superclass_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Singleton_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultiWithFDs_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultiWithBidiFDs_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultiNoFDs_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultiCoveringSets_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MainClass_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_ConflictingIdentSynonym_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_ConflictingIdent_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_superClassValue(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "superClassValue")
}

func Call_Main_singleton(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "singleton")
}

func Call_Main_partialOfFESet(dict_0_loop *Constructor_Main_MultiCoveringSets[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MultiCoveringSets[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V2)
}

func Call_Main_partialOfABSet(dict_0_loop *Constructor_Main_MultiCoveringSets[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_MultiCoveringSets[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_noneOfSets(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "noneOfSets")
}

func Call_Main_multiWithFDs(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "multiWithFDs")
}

func Call_Main_multiWithBidiFDs(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "multiWithBidiFDs")
}

func Call_Main_multiNoFds(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "multiNoFds")
}

func Call_Main_mainClassInt(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "mainClassInt")
}

func Call_Main_conflictingIdentSynonym(dict_0_loop *Constructor_Main_ConflictingIdentSynonym[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ConflictingIdentSynonym[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_conflictingIdent(dict_0_loop *Constructor_Main_ConflictingIdent[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ConflictingIdent[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
