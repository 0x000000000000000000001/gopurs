package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_cIsSymbol gopurs_runtime.Value
var once_Main_cIsSymbol sync.Once

func Get_Main_cIsSymbol() gopurs_runtime.Value {
	once_Main_cIsSymbol.Do(func() {
		cache_Main_cIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("c")
		}))
	})
	return cache_Main_cIsSymbol
}

var cache_Main_bIsSymbol gopurs_runtime.Value
var once_Main_bIsSymbol sync.Once

func Get_Main_bIsSymbol() gopurs_runtime.Value {
	once_Main_bIsSymbol.Do(func() {
		cache_Main_bIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("b")
		}))
	})
	return cache_Main_bIsSymbol
}

var cache_Main_aIsSymbol gopurs_runtime.Value
var once_Main_aIsSymbol sync.Once

func Get_Main_aIsSymbol() gopurs_runtime.Value {
	once_Main_aIsSymbol.Do(func() {
		cache_Main_aIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		}))
	})
	return cache_Main_aIsSymbol
}

var cache_Main_gDecodeJsonCons gopurs_runtime.Value
var once_Main_gDecodeJsonCons sync.Once

func Get_Main_gDecodeJsonCons() gopurs_runtime.Value {
	once_Main_gDecodeJsonCons.Do(func() {
		cache_Main_gDecodeJsonCons = Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))}))
	})
	return cache_Main_gDecodeJsonCons
}

var cache_Main_headIsSymbol gopurs_runtime.Value
var once_Main_headIsSymbol sync.Once

func Get_Main_headIsSymbol() gopurs_runtime.Value {
	once_Main_headIsSymbol.Do(func() {
		cache_Main_headIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("head")
		}))
	})
	return cache_Main_headIsSymbol
}

var cache_Main_eqEither gopurs_runtime.Value
var once_Main_eqEither sync.Once

func Get_Main_eqEither() gopurs_runtime.Value {
	once_Main_eqEither.Do(func() {
		cache_Main_eqEither = gopurs_runtime.Apply(Get_Data_Either_eqEither(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))})
	})
	return cache_Main_eqEither
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

var cache_Main_showEither gopurs_runtime.Value
var once_Main_showEither sync.Once

func Get_Main_showEither() gopurs_runtime.Value {
	once_Main_showEither.Do(func() {
		cache_Main_showEither = gopurs_runtime.Apply(Get_Data_Either_showEither(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))})
	})
	return cache_Main_showEither
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
		cache_Main_showRecordFieldsCons = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsCons(), Get_Main_aIsSymbol())
	})
	return cache_Main_showRecordFieldsCons
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_eqEither1 gopurs_runtime.Value
var once_Main_eqEither1 sync.Once

func Get_Main_eqEither1() gopurs_runtime.Value {
	once_Main_eqEither1.Do(func() {
		cache_Main_eqEither1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_bIsSymbol(), Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_aIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqEither1
}

var cache_Main_showEither1 gopurs_runtime.Value
var once_Main_showEither1 sync.Once

func Get_Main_showEither1() gopurs_runtime.Value {
	once_Main_showEither1.Do(func() {
		cache_Main_showEither1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_aIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_bIsSymbol(), Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showEither1
}

var cache_Main_decodeRecord gopurs_runtime.Value
var once_Main_decodeRecord sync.Once

func Get_Main_decodeRecord() gopurs_runtime.Value {
	once_Main_decodeRecord.Do(func() {
		cache_Main_decodeRecord = gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_2334244318_1277236813(Rebox_Main_1277236813_2334244318(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_bIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_aIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main_decodeRecord
}

var cache_Main_outerIsSymbol gopurs_runtime.Value
var once_Main_outerIsSymbol sync.Once

func Get_Main_outerIsSymbol() gopurs_runtime.Value {
	once_Main_outerIsSymbol.Do(func() {
		cache_Main_outerIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("outer")
		}))
	})
	return cache_Main_outerIsSymbol
}

var cache_Main_requiredIsSymbol gopurs_runtime.Value
var once_Main_requiredIsSymbol sync.Once

func Get_Main_requiredIsSymbol() gopurs_runtime.Value {
	once_Main_requiredIsSymbol.Do(func() {
		cache_Main_requiredIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("required")
		}))
	})
	return cache_Main_requiredIsSymbol
}

var cache_Main_eqEither2 gopurs_runtime.Value
var once_Main_eqEither2 sync.Once

func Get_Main_eqEither2() gopurs_runtime.Value {
	once_Main_eqEither2.Do(func() {
		cache_Main_eqEither2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_outerIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_requiredIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))))}
	})
	return cache_Main_eqEither2
}

var cache_Main_showEither2 gopurs_runtime.Value
var once_Main_showEither2 sync.Once

func Get_Main_showEither2() gopurs_runtime.Value {
	once_Main_showEither2.Do(func() {
		cache_Main_showEither2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_outerIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_requiredIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))))}
	})
	return cache_Main_showEither2
}

var cache_Main_decodeRecord1 gopurs_runtime.Value
var once_Main_decodeRecord1 sync.Once

func Get_Main_decodeRecord1() gopurs_runtime.Value {
	once_Main_decodeRecord1.Do(func() {
		cache_Main_decodeRecord1 = gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_252997862_1277236813(Rebox_Main_1277236813_252997862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_requiredIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{}))), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_outerIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))))}
	})
	return cache_Main_decodeRecord1
}

var cache_Main_tailIsSymbol gopurs_runtime.Value
var once_Main_tailIsSymbol sync.Once

func Get_Main_tailIsSymbol() gopurs_runtime.Value {
	once_Main_tailIsSymbol.Do(func() {
		cache_Main_tailIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("tail")
		}))
	})
	return cache_Main_tailIsSymbol
}

var cache_Main_eqRowCons1 gopurs_runtime.Value
var once_Main_eqRowCons1 sync.Once

func Get_Main_eqRowCons1() gopurs_runtime.Value {
	once_Main_eqRowCons1.Do(func() {
		cache_Main_eqRowCons1 = Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})
	})
	return cache_Main_eqRowCons1
}

var cache_Main_eqEither3 gopurs_runtime.Value
var once_Main_eqEither3 sync.Once

func Get_Main_eqEither3() gopurs_runtime.Value {
	once_Main_eqEither3.Do(func() {
		cache_Main_eqEither3 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_headIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqEither3
}

var cache_Main_showRecordFieldsConsNil gopurs_runtime.Value
var once_Main_showRecordFieldsConsNil sync.Once

func Get_Main_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_Main_showRecordFieldsConsNil.Do(func() {
		cache_Main_showRecordFieldsConsNil = Call_Data_Show_showRecordFieldsConsNil(Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})
	})
	return cache_Main_showRecordFieldsConsNil
}

var cache_Main_showEither3 gopurs_runtime.Value
var once_Main_showEither3 sync.Once

func Get_Main_showEither3() gopurs_runtime.Value {
	once_Main_showEither3.Do(func() {
		cache_Main_showEither3 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_headIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showEither3
}

var cache_Main_eqRec1 gopurs_runtime.Value
var once_Main_eqRec1 sync.Once

func Get_Main_eqRec1() gopurs_runtime.Value {
	once_Main_eqRec1.Do(func() {
		cache_Main_eqRec1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_381738941_3790796878(Rebox_Main_3790796878_381738941(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main_eqRec1
}

var cache_Main_showRecord1 gopurs_runtime.Value
var once_Main_showRecord1 sync.Once

func Get_Main_showRecord1() gopurs_runtime.Value {
	once_Main_showRecord1.Do(func() {
		cache_Main_showRecord1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2815936861_1386611502(Rebox_Main_1386611502_2815936861(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main_showRecord1
}

var cache_Main_Probe gopurs_runtime.Value
var once_Main_Probe sync.Once

func Get_Main_Probe() gopurs_runtime.Value {
	once_Main_Probe.Do(func() {
		cache_Main_Probe = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_Probe(x_0_box.IntVal))
		})
	})
	return cache_Main_Probe
}

var cache_Main_showProbe gopurs_runtime.Value
var once_Main_showProbe sync.Once

func Get_Main_showProbe() gopurs_runtime.Value {
	once_Main_showProbe.Do(func() {
		cache_Main_showProbe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}
	})
	return cache_Main_showProbe
}

var cache_Main_showEither4 gopurs_runtime.Value
var once_Main_showEither4 sync.Once

func Get_Main_showEither4() gopurs_runtime.Value {
	once_Main_showEither4.Do(func() {
		cache_Main_showEither4 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Either_showEither(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_showJsonDecodeError()))}, Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_aIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_bIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showEither4
}

var cache_Main_sharedTail gopurs_runtime.Value
var once_Main_sharedTail sync.Once

func Get_Main_sharedTail() gopurs_runtime.Value {
	once_Main_sharedTail.Do(func() {
		cache_Main_sharedTail = func() gopurs_runtime.Value {
			orig := struct {
				tail int64
			}{int64(41)}
			_ = orig
			return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
		}()
	})
	return cache_Main_sharedTail
}

var cache_Main_object gopurs_runtime.Value
var once_Main_object sync.Once

func Get_Main_object() gopurs_runtime.Value {
	once_Main_object.Do(func() {
		cache_Main_object = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Argonaut_Core_fromObject(), gopurs_runtime.Apply(Get_Foreign_Object_fromFoldable(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}))
	})
	return cache_Main_object
}

var cache_Main_eqProbe gopurs_runtime.Value
var once_Main_eqProbe sync.Once

func Get_Main_eqProbe() gopurs_runtime.Value {
	once_Main_eqProbe.Do(func() {
		cache_Main_eqProbe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878((&Constructor_Data_Eq_Eq[int64]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((x_0.IntVal) == (y_1.IntVal))
		})})))}
	})
	return cache_Main_eqProbe
}

var cache_Main_eqEither4 gopurs_runtime.Value
var once_Main_eqEither4 sync.Once

func Get_Main_eqEither4() gopurs_runtime.Value {
	once_Main_eqEither4.Do(func() {
		cache_Main_eqEither4 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Either_eqEither(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_eqJsonDecodeError()))}, Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_cIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_bIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))}), gopurs_runtime.Value{}, Get_Main_aIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqProbe()))))})))))}
	})
	return cache_Main_eqEither4
}

var cache_Main_calls gopurs_runtime.Value
var once_Main_calls sync.Once

func Get_Main_calls() gopurs_runtime.Value {
	once_Main_calls.Do(func() {
		cache_Main_calls = gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			arr := []int64{}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()))
	})
	return cache_Main_calls
}

var cache_Main_observe gopurs_runtime.Value
var once_Main_observe sync.Once

func Get_Main_observe() gopurs_runtime.Value {
	once_Main_observe.Do(func() {
		cache_Main_observe = gopurs_runtime.Func(func(value_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_observe(value_0_box.IntVal))
		})
	})
	return cache_Main_observe
}

var cache_Main_decodeProbe gopurs_runtime.Value
var once_Main_decodeProbe sync.Once

func Get_Main_decodeProbe() gopurs_runtime.Value {
	once_Main_decodeProbe.Do(func() {
		cache_Main_decodeProbe = gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813((&Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64]{1, gopurs_runtime.Func(func(json_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
			__local_var_1_0 := gopurs_runtime.Apply3(Get_Data_Argonaut_Decode_Internal_Record_typedInt(), func() gopurs_runtime.Value {
				orig := func() struct {
					atIndex      gopurs_runtime.Value
					atKey        gopurs_runtime.Value
					isRight      gopurs_runtime.Value
					just         gopurs_runtime.Value
					left         gopurs_runtime.Value
					leftOf       gopurs_runtime.Value
					missingValue gopurs_runtime.Value
					named        gopurs_runtime.Value
					nothing      *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					right        gopurs_runtime.Value
					rightValue   gopurs_runtime.Value
					typeMismatch gopurs_runtime.Value
				} {
					orig := Get_Data_Argonaut_Decode_Class_recordErrorSupport()
					_ = orig
					clone := struct {
						atIndex      gopurs_runtime.Value
						atKey        gopurs_runtime.Value
						isRight      gopurs_runtime.Value
						just         gopurs_runtime.Value
						left         gopurs_runtime.Value
						leftOf       gopurs_runtime.Value
						missingValue gopurs_runtime.Value
						named        gopurs_runtime.Value
						nothing      *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
						right        gopurs_runtime.Value
						rightValue   gopurs_runtime.Value
						typeMismatch gopurs_runtime.Value
					}{}
					clone.atIndex = gopurs_runtime.RecordGet(orig, "atIndex")
					clone.atKey = gopurs_runtime.RecordGet(orig, "atKey")
					clone.isRight = gopurs_runtime.RecordGet(orig, "isRight")
					clone.just = gopurs_runtime.RecordGet(orig, "just")
					clone.left = gopurs_runtime.RecordGet(orig, "left")
					clone.leftOf = gopurs_runtime.RecordGet(orig, "leftOf")
					clone.missingValue = gopurs_runtime.RecordGet(orig, "missingValue")
					clone.named = gopurs_runtime.RecordGet(orig, "named")
					clone.nothing = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "nothing"))
					clone.right = gopurs_runtime.RecordGet(orig, "right")
					clone.rightValue = gopurs_runtime.RecordGet(orig, "rightValue")
					clone.typeMismatch = gopurs_runtime.RecordGet(orig, "typeMismatch")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"atIndex", "atKey", "isRight", "just", "left", "leftOf", "missingValue", "named", "nothing", "right", "rightValue", "typeMismatch"}, []gopurs_runtime.Value{orig.atIndex, orig.atKey, orig.isRight, orig.just, orig.left, orig.leftOf, orig.missingValue, orig.named, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.nothing)}, orig.right, orig.rightValue, orig.typeMismatch})
			}(), Get_Data_Argonaut_Decode_Decoders_decodeInt(), json_0)
			_ = __local_var_1_0
			var __t5 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 3711209382 {
					// TAST (Let): __local_var_2_1 shape=Other bindingType=(TypeVar e$scope161)
					__local_var_2_1 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
					_ = __local_var_2_1
					__t5 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{__local_var_2_1, gopurs_runtime.Value{}, false}
					goto end_branch_5
				} else {

				}
			}
			{
				if __local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 2465973597 {
					// TAST (Let): __local_var_2_2 shape=Other bindingType=(TypeVar b$scope119)
					__local_var_2_2 := (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
					_ = __local_var_2_2
					// TAST (Let): observed_7_3 shape=App(Var) bindingType=Int
					observed_7_3 := Call_Main_observe(__local_var_2_2.IntVal)
					_ = observed_7_3
					var __t4 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}
					{
						if (observed_7_3) < (int64(0)) {
							__t4 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
								V2 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "negative probe"}))}, gopurs_runtime.Value{}, false}
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{gopurs_runtime.Value{}, gopurs_runtime.Int(observed_7_3), true}
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
					V2 bool
				}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
			}
		end_branch_5:
			return func() gopurs_runtime.Value {
				_v := __t5
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})})))}
	})
	return cache_Main_decodeProbe
}

var cache_Main_gDecodeJsonCons1 gopurs_runtime.Value
var once_Main_gDecodeJsonCons1 sync.Once

func Get_Main_gDecodeJsonCons1() gopurs_runtime.Value {
	once_Main_gDecodeJsonCons1.Do(func() {
		cache_Main_gDecodeJsonCons1 = Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main_decodeProbe()))))}))
	})
	return cache_Main_gDecodeJsonCons1
}

var cache_Main_decodeTracked gopurs_runtime.Value
var once_Main_decodeTracked sync.Once

func Get_Main_decodeTracked() gopurs_runtime.Value {
	once_Main_decodeTracked.Do(func() {
		cache_Main_decodeTracked = Call_Data_Argonaut_Decode_Class_decodeJson(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main_decodeProbe()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main_decodeProbe()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Main_decodeProbe()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_cIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_bIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_aIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{})))
	})
	return cache_Main_decodeTracked
}

var cache_Main_tailResult gopurs_runtime.Value
var once_Main_tailResult sync.Once

func Get_Main_tailResult() gopurs_runtime.Value {
	once_Main_tailResult.Do(func() {
		cache_Main_tailResult = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_tailResult(v_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Main_tailResult
}

var cache_Main_customTailDecoder gopurs_runtime.Value
var once_Main_customTailDecoder sync.Once

func Get_Main_customTailDecoder() gopurs_runtime.Value {
	once_Main_customTailDecoder.Do(func() {
		cache_Main_customTailDecoder = gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(value_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_tailResult(value_0)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})}))}
	})
	return cache_Main_customTailDecoder
}

var cache_Main_gDecodeJsonCons2 gopurs_runtime.Value
var once_Main_gDecodeJsonCons2 sync.Once

func Get_Main_gDecodeJsonCons2() gopurs_runtime.Value {
	once_Main_gDecodeJsonCons2.Do(func() {
		cache_Main_gDecodeJsonCons2 = gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_customTailDecoder()))}, Get_Main_headIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})))}
	})
	return cache_Main_gDecodeJsonCons2
}

var cache_Main_decodeCustom gopurs_runtime.Value
var once_Main_decodeCustom sync.Once

func Get_Main_decodeCustom() gopurs_runtime.Value {
	once_Main_decodeCustom.Do(func() {
		cache_Main_decodeCustom = gopurs_runtime.Func(func(value_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Main_decodeCustom(value_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Main_decodeCustom
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
			arr := []int64{}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), Get_Main_calls()), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1953230001("", struct {
				actual   gopurs_runtime.Value
				expected gopurs_runtime.Value
			}{gopurs_runtime.Apply(Get_Main_decodeTracked(), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
				arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Str("c"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(3.0))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Str("a"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(1.0))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Str("b"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(2.0))}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
				}
				return gopurs_runtime.Array(boxed)
			}())), func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 struct {
						a int64
						b int64
						c int64
					}
					V2 bool
				}{gopurs_runtime.Value{}, struct {
					a int64
					b int64
					c int64
				}{int64(1), int64(2), int64(3)}, true}
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b), gopurs_runtime.Int(orig.c))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
				__local_var_2_0 := gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls())
				_ = __local_var_2_0
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
					_ = __local_var_3_1
					return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
						actual   []int64
						expected []int64
					}{func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(__local_var_3_1.UnsafePtr)
						unboxed := make([]int64, len(arr))
						for i, v := range arr {
							unboxed[i] = v.IntVal
						}
						return unboxed
					}(), []int64{int64(1), int64(2), int64(3)}}), gopurs_runtime.Value{})
				}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
						arr := []int64{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), Get_Main_calls()), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1953230001("", struct {
							actual   gopurs_runtime.Value
							expected gopurs_runtime.Value
						}{gopurs_runtime.Apply(Get_Main_decodeTracked(), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
							arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{gopurs_runtime.Str("a"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(1.0))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{gopurs_runtime.Str("b"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(-2.0))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 gopurs_runtime.Value
								}{gopurs_runtime.Str("c"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(-3.0))}
								return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
							}()))}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
							}
							return gopurs_runtime.Array(boxed)
						}())), func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 struct {
									a int64
									b int64
									c int64
								}
								V2 bool
							}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "b", gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "negative probe"}))}}))}, struct {
								a int64
								b int64
								c int64
							}{}, false}
							if _v.V2 {
								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
									orig := _v.V1
									_ = orig
									return gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b), gopurs_runtime.Int(orig.c))
								}()})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
						}()}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=Any
							__local_var_6_2 := gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls())
							_ = __local_var_6_2
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_7_3 := gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Value{})
								_ = __local_var_7_3
								return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
									actual   []int64
									expected []int64
								}{func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(__local_var_7_3.UnsafePtr)
									unboxed := make([]int64, len(arr))
									for i, v := range arr {
										unboxed[i] = v.IntVal
									}
									return unboxed
								}(), []int64{int64(1), int64(-2)}}), gopurs_runtime.Value{})
							}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
									arr := []int64{}
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), Get_Main_calls()), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1953230001("", struct {
										actual   gopurs_runtime.Value
										expected gopurs_runtime.Value
									}{gopurs_runtime.Apply(Get_Main_decodeTracked(), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
										arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
											_v := struct {
												V0 gopurs_runtime.Value
												V1 gopurs_runtime.Value
											}{gopurs_runtime.Str("a"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(1.0))}
											return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
										}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
											_v := struct {
												V0 gopurs_runtime.Value
												V1 gopurs_runtime.Value
											}{gopurs_runtime.Str("c"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(3.0))}
											return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
										}()))}
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
										}
										return gopurs_runtime.Array(boxed)
									}())), func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 struct {
												a int64
												b int64
												c int64
											}
											V2 bool
										}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "b", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}, struct {
											a int64
											b int64
											c int64
										}{}, false}
										if _v.V2 {
											return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
												orig := _v.V1
												_ = orig
												return gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b), gopurs_runtime.Int(orig.c))
											}()})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
									}()}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_10_4 shape=App(Var) bindingType=Any
										__local_var_10_4 := gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls())
										_ = __local_var_10_4
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											__local_var_11_5 := gopurs_runtime.Apply(__local_var_10_4, gopurs_runtime.Value{})
											_ = __local_var_11_5
											return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
												actual   []int64
												expected []int64
											}{func() []int64 {
												arr := *(*[]gopurs_runtime.Value)(__local_var_11_5.UnsafePtr)
												unboxed := make([]int64, len(arr))
												for i, v := range arr {
													unboxed[i] = v.IntVal
												}
												return unboxed
											}(), []int64{int64(1)}}), gopurs_runtime.Value{})
										}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1547446417("", struct {
												actual   gopurs_runtime.Value
												expected gopurs_runtime.Value
											}{gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_bIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_aIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{}), "decodeJson"), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
												arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
													_v := struct {
														V0 gopurs_runtime.Value
														V1 gopurs_runtime.Value
													}{gopurs_runtime.Str("a"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(7.0))}
													return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
												}()))}
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
												}
												return gopurs_runtime.Array(boxed)
											}())), func() gopurs_runtime.Value {
												_v := struct {
													V0 gopurs_runtime.Value
													V1 struct {
														a int64
														b *Constructor_Data_Maybe_Just[int64]
													}
													V2 bool
												}{gopurs_runtime.Value{}, struct {
													a int64
													b *Constructor_Data_Maybe_Just[int64]
												}{int64(7), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, true}
												if _v.V2 {
													return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
														orig := _v.V1
														_ = orig
														return gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(orig.a), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.b))})
													}()})}
												}
												return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
											}()}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1547446417("", struct {
													actual   gopurs_runtime.Value
													expected gopurs_runtime.Value
												}{gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_bIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), Get_Main_aIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{}), "decodeJson"), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
													arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 gopurs_runtime.Value
														}{gopurs_runtime.Str("a"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(7.0))}
														return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
													}())), Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 gopurs_runtime.Value
														}{gopurs_runtime.Str("b"), Get_Data_Argonaut_Core_jsonNull()}
														return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
													}()))}
													boxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
													}
													return gopurs_runtime.Array(boxed)
												}())), func() gopurs_runtime.Value {
													_v := struct {
														V0 gopurs_runtime.Value
														V1 struct {
															a int64
															b *Constructor_Data_Maybe_Just[int64]
														}
														V2 bool
													}{gopurs_runtime.Value{}, struct {
														a int64
														b *Constructor_Data_Maybe_Just[int64]
													}{int64(7), Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, true}
													if _v.V2 {
														return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
															orig := _v.V1
															_ = orig
															return gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(orig.a), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_1170268447_3094389156(orig.b))})
														}()})}
													}
													return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
												}()}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1064098801("", struct {
														actual   gopurs_runtime.Value
														expected gopurs_runtime.Value
													}{gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Apply(Call_Data_Argonaut_Decode_Class_decodeRecord(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_requiredIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{}))), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_outerIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{})), gopurs_runtime.Value{}), "decodeJson"), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
														arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{Rebox_Main_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
															_v := struct {
																V0 gopurs_runtime.Value
																V1 gopurs_runtime.Value
															}{gopurs_runtime.Str("outer"), gopurs_runtime.Apply(Get_Main_object(), func() gopurs_runtime.Value {
																arr := []*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{}
																boxed := make([]gopurs_runtime.Value, len(arr))
																for i, v := range arr {
																	boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
																}
																return gopurs_runtime.Array(boxed)
															}())}
															return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
														}()))}
														boxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3103076375_138441832(v))}
														}
														return gopurs_runtime.Array(boxed)
													}())), func() gopurs_runtime.Value {
														_v := struct {
															V0 gopurs_runtime.Value
															V1 struct {
																outer struct {
																	required int64
																}
															}
															V2 bool
														}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "outer", gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "required", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}}))}, struct {
															outer struct {
																required int64
															}
														}{}, false}
														if _v.V2 {
															return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																orig := _v.V1
																_ = orig
																return gopurs_runtime.RecordDict1("outer", func() gopurs_runtime.Value {
																	orig := orig.outer
																	_ = orig
																	return gopurs_runtime.RecordDict1("required", gopurs_runtime.Int(orig.required))
																}())
															}()})}
														}
														return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
													}()}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
															arr := []int64{}
															boxed := make([]gopurs_runtime.Value, len(arr))
															for i, v := range arr {
																boxed[i] = gopurs_runtime.Int(v)
															}
															return gopurs_runtime.Array(boxed)
														}(), Get_Main_calls()), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																__local_var_16_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
																	_v := Call_Main_decodeCustom(gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST())), Get_Foreign_Object_ST_go__new(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str("head"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(1.0))))))
																	if _v.V2 {
																		return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																			orig := _v.V1
																			_ = orig
																			return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
																		}()})}
																	}
																	return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																}()), gopurs_runtime.Value{})
																_ = __local_var_16_6
																__local_var_17_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
																	_v := Call_Main_decodeCustom(gopurs_runtime.Apply(Get_Foreign_Object_runST(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_bindST())), Get_Foreign_Object_ST_go__new(), gopurs_runtime.Apply2(Get_Foreign_Object_ST_poke(), gopurs_runtime.Str("head"), gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Float(2.0))))))
																	if _v.V2 {
																		return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																			orig := _v.V1
																			_ = orig
																			return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
																		}()})}
																	}
																	return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																}()), gopurs_runtime.Value{})
																_ = __local_var_17_7
																// TAST (Let): __local_var_18_8 shape=App(Var) bindingType=Any
																__local_var_18_8 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_6)
																_ = __local_var_18_8
																return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_19_9 := gopurs_runtime.Apply(__local_var_18_8, gopurs_runtime.Value{})
																	_ = __local_var_19_9
																	return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1570739249("", struct {
																		actual   gopurs_runtime.Value
																		expected gopurs_runtime.Value
																	}{__local_var_19_9, func() gopurs_runtime.Value {
																		_v := struct {
																			V0 gopurs_runtime.Value
																			V1 struct {
																				head int64
																				tail int64
																			}
																			V2 bool
																		}{gopurs_runtime.Value{}, struct {
																			head int64
																			tail int64
																		}{int64(1), int64(41)}, true}
																		if _v.V2 {
																			return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																				orig := _v.V1
																				_ = orig
																				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
																			}()})}
																		}
																		return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																	}()}), gopurs_runtime.Value{})
																}), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
																	// TAST (Let): __local_var_20_10 shape=App(Var) bindingType=Any
																	__local_var_20_10 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_17_7)
																	_ = __local_var_20_10
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																		__local_var_21_11 := gopurs_runtime.Apply(__local_var_20_10, gopurs_runtime.Value{})
																		_ = __local_var_21_11
																		return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1570739249("", struct {
																			actual   gopurs_runtime.Value
																			expected gopurs_runtime.Value
																		}{__local_var_21_11, func() gopurs_runtime.Value {
																			_v := struct {
																				V0 gopurs_runtime.Value
																				V1 struct {
																					head int64
																					tail int64
																				}
																				V2 bool
																			}{gopurs_runtime.Value{}, struct {
																				head int64
																				tail int64
																			}{int64(2), int64(41)}, true}
																			if _v.V2 {
																				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																					orig := _v.V1
																					_ = orig
																					return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
																				}()})}
																			}
																			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																		}()}), gopurs_runtime.Value{})
																	}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1981780840("", struct {
																			actual struct {
																				tail int64
																			}
																			expected struct {
																				tail int64
																			}
																		}{func() struct {
																			tail int64
																		} {
																			orig := Get_Main_sharedTail()
																			_ = orig
																			clone := struct {
																				tail int64
																			}{}
																			clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
																			return clone
																		}(), struct {
																			tail int64
																		}{int64(41)}}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeHas(), gopurs_runtime.Str("head"), func() gopurs_runtime.Value {
																				orig := func() struct {
																					tail int64
																				} {
																					orig := Get_Main_sharedTail()
																					_ = orig
																					clone := struct {
																						tail int64
																					}{}
																					clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
																					return clone
																				}()
																				_ = orig
																				return gopurs_runtime.RecordDict1("tail", gopurs_runtime.Int(orig.tail))
																			}()).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																				// TAST (Let): __local_var_24_12 shape=App(Var) bindingType=Any
																				__local_var_24_12 := gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls())
																				_ = __local_var_24_12
																				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																					__local_var_25_13 := gopurs_runtime.Apply(__local_var_24_12, gopurs_runtime.Value{})
																					_ = __local_var_25_13
																					return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
																						actual   []int64
																						expected []int64
																					}{func() []int64 {
																						arr := *(*[]gopurs_runtime.Value)(__local_var_25_13.UnsafePtr)
																						unboxed := make([]int64, len(arr))
																						for i, v := range arr {
																							unboxed[i] = v.IntVal
																						}
																						return unboxed
																					}(), []int64{int64(100), int64(100)}}), gopurs_runtime.Value{})
																				}), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
																						arr := []int64{}
																						boxed := make([]gopurs_runtime.Value, len(arr))
																						for i, v := range arr {
																							boxed[i] = gopurs_runtime.Int(v)
																						}
																						return gopurs_runtime.Array(boxed)
																					}(), Get_Main_calls()), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1570739249("", struct {
																							actual   gopurs_runtime.Value
																							expected gopurs_runtime.Value
																						}{func() gopurs_runtime.Value {
																							_v := Call_Main_decodeCustom(Get_Foreign_Object_empty())
																							if _v.V2 {
																								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																									orig := _v.V1
																									_ = orig
																									return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
																								}()})}
																							}
																							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																						}(), func() gopurs_runtime.Value {
																							_v := struct {
																								V0 gopurs_runtime.Value
																								V1 struct {
																									head int64
																									tail int64
																								}
																								V2 bool
																							}{gopurs_runtime.Value{Type: 9, IntVal: 1896025177, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_AtKey{1, "head", gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}}))}, struct {
																								head int64
																								tail int64
																							}{}, false}
																							if _v.V2 {
																								return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																									orig := _v.V1
																									_ = orig
																									return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Int(orig.tail))
																								}()})}
																							}
																							return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																						}()}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																							// TAST (Let): __local_var_28_14 shape=App(Var) bindingType=Any
																							__local_var_28_14 := gopurs_runtime.Apply(Get_Effect_Ref_read(), Get_Main_calls())
																							_ = __local_var_28_14
																							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																								__local_var_29_15 := gopurs_runtime.Apply(__local_var_28_14, gopurs_runtime.Value{})
																								_ = __local_var_29_15
																								return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
																									actual   []int64
																									expected []int64
																								}{func() []int64 {
																									arr := *(*[]gopurs_runtime.Value)(__local_var_29_15.UnsafePtr)
																									unboxed := make([]int64, len(arr))
																									for i, v := range arr {
																										unboxed[i] = v.IntVal
																									}
																									return unboxed
																								}(), []int64{}}), gopurs_runtime.Value{})
																							}), gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
																								return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																							}))
																						}))
																					}))
																				}))
																			}))
																		}))
																	}))
																})), gopurs_runtime.Value{})
															})
														}))
													}))
												}))
											}))
										}))
									}))
								}))
							}))
						}))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main_Probe(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_observe(value_0_loop int64) int64 {
	var value_0 int64 = value_0_loop
	_ = value_0
	return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), seen_1, func() gopurs_runtime.Value {
					arr := []int64{value_0}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr))).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	}), Get_Main_calls()), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(value_0)
		})
	}))).IntVal
}

func Call_Main_tailResult(v_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		tail int64
	}
	V2 bool
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 struct {
		V0 gopurs_runtime.Value
		V1 struct {
			tail int64
		}
		V2 bool
	}
	{
		if (Call_Main_observe(int64(100))) == (int64(100)) {
			__t0 = struct {
				V0 gopurs_runtime.Value
				V1 struct {
					tail int64
				}
				V2 bool
			}{gopurs_runtime.Value{}, func() struct {
				tail int64
			} {
				orig := Get_Main_sharedTail()
				_ = orig
				clone := struct {
					tail int64
				}{}
				clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
				return clone
			}(), true}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = struct {
			V0 gopurs_runtime.Value
			V1 struct {
				tail int64
			}
			V2 bool
		}{gopurs_runtime.Value{Type: 9, IntVal: 3199441748, UnsafePtr: unsafe.Pointer(nil)}, struct {
			tail int64
		}{}, false}
	}
end_branch_0:
	return __t0
}

func Call_Main_decodeCustom(value_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		head int64
		tail int64
	}
	V2 bool
} {
	var value_0 gopurs_runtime.Value = value_0_loop
	_ = value_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 struct {
			head int64
			tail int64
		}
		V2 bool
	} {
		_v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply4(Call_Data_Argonaut_Decode_Class_gDecodeJsonCons(Call_Data_Argonaut_Decode_Class_decodeFieldId(gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Main_1405234230_1277236813(Rebox_Main_1277236813_1405234230(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_customTailDecoder()))}, Get_Main_headIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), "gDecodeJson"), value_0, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
		if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 struct {
					head int64
					tail int64
				}
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: func() struct {
				head int64
				tail int64
			} {
				orig := (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0
				_ = orig
				clone := struct {
					head int64
					tail int64
				}{}
				clone.head = gopurs_runtime.RecordGet(orig, "head").IntVal
				clone.tail = gopurs_runtime.RecordGet(orig, "tail").IntVal
				return clone
			}(), V2: true}
		}
		return struct {
			V0 gopurs_runtime.Value
			V1 struct {
				head int64
				tail int64
			}
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: struct {
			head int64
			tail int64
		}{}, V2: false}
	}()
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_1277236813_1405234230(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1277236813_2334244318(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	a int64
	b *Constructor_Data_Maybe_Just[int64]
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		a int64
		b *Constructor_Data_Maybe_Just[int64]
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1277236813_252997862(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	outer struct {
		required int64
	}
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		outer struct {
			required int64
		}
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_138441832_3103076375(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1
	return out
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2815936861(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	tail int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		tail int64
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1405234230_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[int64]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
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

func Rebox_Main_2334244318_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	a int64
	b *Constructor_Data_Maybe_Just[int64]
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_252997862_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	outer struct {
		required int64
	}
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2815936861_1386611502(in *Constructor_Data_Show_Show[struct {
	tail int64
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Main_3103076375_138441832(in *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_381738941(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	tail int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		tail int64
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_381738941_3790796878(in *Constructor_Data_Eq_Eq[struct {
	tail int64
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
