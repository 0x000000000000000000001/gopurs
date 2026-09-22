package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_nameIsSymbol gopurs_runtime.Value
var once_Main_nameIsSymbol sync.Once

func Get_Main_nameIsSymbol() gopurs_runtime.Value {
	once_Main_nameIsSymbol.Do(func() {
		cache_Main_nameIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("name")
		}))
	})
	return cache_Main_nameIsSymbol
}

var cache_Main_idIsSymbol gopurs_runtime.Value
var once_Main_idIsSymbol sync.Once

func Get_Main_idIsSymbol() gopurs_runtime.Value {
	once_Main_idIsSymbol.Do(func() {
		cache_Main_idIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("id")
		}))
	})
	return cache_Main_idIsSymbol
}

var cache_Main_extraIsSymbol gopurs_runtime.Value
var once_Main_extraIsSymbol sync.Once

func Get_Main_extraIsSymbol() gopurs_runtime.Value {
	once_Main_extraIsSymbol.Do(func() {
		cache_Main_extraIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("extra")
		}))
	})
	return cache_Main_extraIsSymbol
}

var cache_Main_activeIsSymbol gopurs_runtime.Value
var once_Main_activeIsSymbol sync.Once

func Get_Main_activeIsSymbol() gopurs_runtime.Value {
	once_Main_activeIsSymbol.Do(func() {
		cache_Main_activeIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("active")
		}))
	})
	return cache_Main_activeIsSymbol
}

var cache_Main_eqRec1 gopurs_runtime.Value
var once_Main_eqRec1 sync.Once

func Get_Main_eqRec1() gopurs_runtime.Value {
	once_Main_eqRec1.Do(func() {
		cache_Main_eqRec1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2758774104_3790796878(Rebox_Main_3790796878_2758774104(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_nameIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), gopurs_runtime.Value{}, Get_Main_idIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_extraIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_activeIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main_eqRec1
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
		cache_Main_showRecordFieldsCons = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsCons(), Get_Main_activeIsSymbol())
	})
	return cache_Main_showRecordFieldsCons
}

var cache_Main_showRecordFieldsCons1 gopurs_runtime.Value
var once_Main_showRecordFieldsCons1 sync.Once

func Get_Main_showRecordFieldsCons1() gopurs_runtime.Value {
	once_Main_showRecordFieldsCons1.Do(func() {
		cache_Main_showRecordFieldsCons1 = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsCons(), Get_Main_idIsSymbol())
	})
	return cache_Main_showRecordFieldsCons1
}

var cache_Main_showRecord1 gopurs_runtime.Value
var once_Main_showRecord1 sync.Once

func Get_Main_showRecord1() gopurs_runtime.Value {
	once_Main_showRecord1.Do(func() {
		cache_Main_showRecord1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1994653624_1386611502(Rebox_Main_1386611502_1994653624(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_activeIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_extraIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_idIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_nameIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main_showRecord1
}

var cache_Main_nestedIsSymbol gopurs_runtime.Value
var once_Main_nestedIsSymbol sync.Once

func Get_Main_nestedIsSymbol() gopurs_runtime.Value {
	once_Main_nestedIsSymbol.Do(func() {
		cache_Main_nestedIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("nested")
		}))
	})
	return cache_Main_nestedIsSymbol
}

var cache_Main_retainedIsSymbol gopurs_runtime.Value
var once_Main_retainedIsSymbol sync.Once

func Get_Main_retainedIsSymbol() gopurs_runtime.Value {
	once_Main_retainedIsSymbol.Do(func() {
		cache_Main_retainedIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("retained")
		}))
	})
	return cache_Main_retainedIsSymbol
}

var cache_Main_eqRowCons1 gopurs_runtime.Value
var once_Main_eqRowCons1 sync.Once

func Get_Main_eqRowCons1() gopurs_runtime.Value {
	once_Main_eqRowCons1.Do(func() {
		cache_Main_eqRowCons1 = Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
	})
	return cache_Main_eqRowCons1
}

var cache_Main_beforeIsSymbol gopurs_runtime.Value
var once_Main_beforeIsSymbol sync.Once

func Get_Main_beforeIsSymbol() gopurs_runtime.Value {
	once_Main_beforeIsSymbol.Do(func() {
		cache_Main_beforeIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("before")
		}))
	})
	return cache_Main_beforeIsSymbol
}

var cache_Main_eqRec2 gopurs_runtime.Value
var once_Main_eqRec2 sync.Once

func Get_Main_eqRec2() gopurs_runtime.Value {
	once_Main_eqRec2.Do(func() {
		cache_Main_eqRec2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2543566495_3790796878(Rebox_Main_3790796878_2543566495(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_nestedIsSymbol(), Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))), gopurs_runtime.Value{}, Get_Main_nameIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), gopurs_runtime.Value{}, Get_Main_idIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_beforeIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_687527510_3790796878(Rebox_Main_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}), gopurs_runtime.Value{}, Get_Main_activeIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main_eqRec2
}

var cache_Main_showRecordFieldsCons2 gopurs_runtime.Value
var once_Main_showRecordFieldsCons2 sync.Once

func Get_Main_showRecordFieldsCons2() gopurs_runtime.Value {
	once_Main_showRecordFieldsCons2.Do(func() {
		cache_Main_showRecordFieldsCons2 = gopurs_runtime.Apply(Get_Data_Show_showRecordFieldsCons(), Get_Main_nameIsSymbol())
	})
	return cache_Main_showRecordFieldsCons2
}

var cache_Main_showRecordFieldsConsNil gopurs_runtime.Value
var once_Main_showRecordFieldsConsNil sync.Once

func Get_Main_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_Main_showRecordFieldsConsNil.Do(func() {
		cache_Main_showRecordFieldsConsNil = Call_Data_Show_showRecordFieldsConsNil(Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})
	})
	return cache_Main_showRecordFieldsConsNil
}

var cache_Main_showRecord2 gopurs_runtime.Value
var once_Main_showRecord2 sync.Once

func Get_Main_showRecord2() gopurs_runtime.Value {
	once_Main_showRecord2.Do(func() {
		cache_Main_showRecord2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_872707775_1386611502(Rebox_Main_1386611502_872707775(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_activeIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_beforeIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_idIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_nameIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_nestedIsSymbol(), Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main_showRecord2
}

var cache_Main_zIsSymbol gopurs_runtime.Value
var once_Main_zIsSymbol sync.Once

func Get_Main_zIsSymbol() gopurs_runtime.Value {
	once_Main_zIsSymbol.Do(func() {
		cache_Main_zIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("z")
		}))
	})
	return cache_Main_zIsSymbol
}

var cache_Main_valuesIsSymbol gopurs_runtime.Value
var once_Main_valuesIsSymbol sync.Once

func Get_Main_valuesIsSymbol() gopurs_runtime.Value {
	once_Main_valuesIsSymbol.Do(func() {
		cache_Main_valuesIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("values")
		}))
	})
	return cache_Main_valuesIsSymbol
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

var cache_Main_eqRec3 gopurs_runtime.Value
var once_Main_eqRec3 sync.Once

func Get_Main_eqRec3() gopurs_runtime.Value {
	once_Main_eqRec3.Do(func() {
		cache_Main_eqRec3 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2267788589_3790796878(Rebox_Main_3790796878_2267788589(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_zIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), gopurs_runtime.Value{}, Get_Main_valuesIsSymbol(), Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})), gopurs_runtime.Value{}, Get_Main_tailIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}), gopurs_runtime.Value{}, Get_Main_nameIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), gopurs_runtime.Value{}, Get_Main_idIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_activeIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main_eqRec3
}

var cache_Main_showRecord3 gopurs_runtime.Value
var once_Main_showRecord3 sync.Once

func Get_Main_showRecord3() gopurs_runtime.Value {
	once_Main_showRecord3.Do(func() {
		cache_Main_showRecord3 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1517968077_1386611502(Rebox_Main_1386611502_1517968077(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_activeIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_idIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_nameIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_tailIsSymbol(), Call_Data_Show_showRecordFieldsCons(Get_Main_valuesIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_zIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main_showRecord3
}

var cache_Main_eqRec4 gopurs_runtime.Value
var once_Main_eqRec4 sync.Once

func Get_Main_eqRec4() gopurs_runtime.Value {
	once_Main_eqRec4.Do(func() {
		cache_Main_eqRec4 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1830338253_3790796878(Rebox_Main_3790796878_1830338253(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))))}
	})
	return cache_Main_eqRec4
}

var cache_Main_showRecord4 gopurs_runtime.Value
var once_Main_showRecord4 sync.Once

func Get_Main_showRecord4() gopurs_runtime.Value {
	once_Main_showRecord4.Do(func() {
		cache_Main_showRecord4 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_920949869_1386611502(Rebox_Main_1386611502_920949869(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}))))))}
	})
	return cache_Main_showRecord4
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

var cache_Main_makeC gopurs_runtime.Value
var once_Main_makeC sync.Once

func Get_Main_makeC() gopurs_runtime.Value {
	once_Main_makeC.Do(func() {
		cache_Main_makeC = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeC(n_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"active", "id", "name", "tail", "values", "z"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name), gopurs_runtime.Bool(orig.tail), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Str(orig.z)})
			}()
		})
	})
	return cache_Main_makeC
}

var cache_Main_runC gopurs_runtime.Value
var once_Main_runC sync.Once

func Get_Main_runC() gopurs_runtime.Value {
	once_Main_runC.Do(func() {
		cache_Main_runC = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runC(n_0_box.IntVal))
		})
	})
	return cache_Main_runC
}

var cache_Main_makeB gopurs_runtime.Value
var once_Main_makeB sync.Once

func Get_Main_makeB() gopurs_runtime.Value {
	once_Main_makeB.Do(func() {
		cache_Main_makeB = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeB(n_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict5("active", "before", "id", "name", "nested", gopurs_runtime.Bool(orig.active), gopurs_runtime.Float(orig.before), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name), func() gopurs_runtime.Value {
					orig := orig.nested
					_ = orig
					return gopurs_runtime.RecordDict1("retained", gopurs_runtime.Str(orig.retained))
				}())
			}()
		})
	})
	return cache_Main_makeB
}

var cache_Main_runB gopurs_runtime.Value
var once_Main_runB sync.Once

func Get_Main_runB() gopurs_runtime.Value {
	once_Main_runB.Do(func() {
		cache_Main_runB = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runB(n_0_box.IntVal))
		})
	})
	return cache_Main_runB
}

var cache_Main_makeA gopurs_runtime.Value
var once_Main_makeA sync.Once

func Get_Main_makeA() gopurs_runtime.Value {
	once_Main_makeA.Do(func() {
		cache_Main_makeA = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeA(n_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict4("active", "extra", "id", "name", gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.extra), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name))
			}()
		})
	})
	return cache_Main_makeA
}

var cache_Main_runA gopurs_runtime.Value
var once_Main_runA sync.Once

func Get_Main_runA() gopurs_runtime.Value {
	once_Main_runA.Do(func() {
		cache_Main_runA = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runA(n_0_box.IntVal))
		})
	})
	return cache_Main_runA
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
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := Call_Main_makeA(__local_var_2_2.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict4("active", "extra", "id", "name", gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.extra), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name))
			}()), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := Call_Main_makeB(__local_var_2_2.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict5("active", "before", "id", "name", "nested", gopurs_runtime.Bool(orig.active), gopurs_runtime.Float(orig.before), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name), func() gopurs_runtime.Value {
					orig := orig.nested
					_ = orig
					return gopurs_runtime.RecordDict1("retained", gopurs_runtime.Str(orig.retained))
				}())
			}()), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := Call_Main_makeC(__local_var_2_2.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"active", "id", "name", "tail", "values", "z"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name), gopurs_runtime.Bool(orig.tail), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Str(orig.z)})
			}()), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_3), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_4), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_5), gopurs_runtime.Value{})
			_ = __local_var_8_8
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Worker_score(func(record struct {
				active bool
				extra  int64
				id     int64
				name   string
			}) struct {
				active bool
				id     int64
				name   string
			} {
				return struct {
					active bool
					id     int64
					name   string
				}{record.active, record.id, record.name}
			}(func() struct {
				active bool
				extra  int64
				id     int64
				name   string
			} {
				orig := __local_var_6_6
				_ = orig
				clone := struct {
					active bool
					extra  int64
					id     int64
					name   string
				}{}
				clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
				clone.extra = gopurs_runtime.RecordGet(orig, "extra").IntVal
				clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
				clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
				return clone
			}())), int64(2851)}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{Call_Worker_score(func(record struct {
					active bool
					before float64
					id     int64
					name   string
					nested struct {
						retained string
					}
				}) struct {
					active bool
					id     int64
					name   string
				} {
					return struct {
						active bool
						id     int64
						name   string
					}{record.active, record.id, record.name}
				}(func() struct {
					active bool
					before float64
					id     int64
					name   string
					nested struct {
						retained string
					}
				} {
					orig := __local_var_7_7
					_ = orig
					clone := struct {
						active bool
						before float64
						id     int64
						name   string
						nested struct {
							retained string
						}
					}{}
					clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
					clone.before = gopurs_runtime.RecordGet(orig, "before").FloatVal()
					clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
					clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
					clone.nested = func() struct {
						retained string
					} {
						orig := gopurs_runtime.RecordGet(orig, "nested")
						_ = orig
						clone := struct {
							retained string
						}{}
						clone.retained = gopurs_runtime.RecordGet(orig, "retained").StrVal()
						return clone
					}()
					return clone
				}())), int64(-195)}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{Call_Worker_score(func(record struct {
						active bool
						id     int64
						name   string
						tail   bool
						values []int64
						z      string
					}) struct {
						active bool
						id     int64
						name   string
					} {
						return struct {
							active bool
							id     int64
							name   string
						}{record.active, record.id, record.name}
					}(func() struct {
						active bool
						id     int64
						name   string
						tail   bool
						values []int64
						z      string
					} {
						orig := __local_var_8_8
						_ = orig
						clone := struct {
							active bool
							id     int64
							name   string
							tail   bool
							values []int64
							z      string
						}{}
						clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
						clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
						clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
						clone.tail = (gopurs_runtime.RecordGet(orig, "tail").IntVal) != (0)
						clone.values = func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
							unboxed := make([]int64, len(arr))
							for i, v := range arr {
								unboxed[i] = v.IntVal
							}
							return unboxed
						}()
						clone.z = gopurs_runtime.RecordGet(orig, "z").StrVal()
						return clone
					}())), int64(2851)}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): a2_12_9 shape=Other bindingType=(Record (Row [id: Int, active: Boolean, extra: Int, name: String] Empty))
						a2_12_9 := func() struct {
							active bool
							extra  int64
							id     int64
							name   string
						} {
							orig := gopurs_runtime.RecordUpdate1(__local_var_6_6, "id", gopurs_runtime.Int(int64(9)))
							_ = orig
							clone := struct {
								active bool
								extra  int64
								id     int64
								name   string
							}{}
							clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
							clone.extra = gopurs_runtime.RecordGet(orig, "extra").IntVal
							clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
							clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
							return clone
						}()
						_ = a2_12_9
						// TAST (Let): b2_13_10 shape=Other bindingType=(Record (Row [name: String, active: Boolean, before: Number, id: Int, nested: (Record (Row [retained: String] Empty))] Empty))
						b2_13_10 := func() struct {
							active bool
							before float64
							id     int64
							name   string
							nested struct {
								retained string
							}
						} {
							orig := gopurs_runtime.RecordUpdate1(__local_var_7_7, "name", gopurs_runtime.Str("alpha"))
							_ = orig
							clone := struct {
								active bool
								before float64
								id     int64
								name   string
								nested struct {
									retained string
								}
							}{}
							clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
							clone.before = gopurs_runtime.RecordGet(orig, "before").FloatVal()
							clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
							clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
							clone.nested = func() struct {
								retained string
							} {
								orig := gopurs_runtime.RecordGet(orig, "nested")
								_ = orig
								clone := struct {
									retained string
								}{}
								clone.retained = gopurs_runtime.RecordGet(orig, "retained").StrVal()
								return clone
							}()
							return clone
						}()
						_ = b2_13_10
						// TAST (Let): c2_14_11 shape=Other bindingType=(Record (Row [active: Boolean, id: Int, name: String, tail: Boolean, values: (Array Int), z: String] Empty))
						c2_14_11 := func() struct {
							active bool
							id     int64
							name   string
							tail   bool
							values []int64
							z      string
						} {
							orig := gopurs_runtime.RecordUpdate1(__local_var_8_8, "active", gopurs_runtime.Bool(false))
							_ = orig
							clone := struct {
								active bool
								id     int64
								name   string
								tail   bool
								values []int64
								z      string
							}{}
							clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
							clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
							clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
							clone.tail = (gopurs_runtime.RecordGet(orig, "tail").IntVal) != (0)
							clone.values = func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							clone.z = gopurs_runtime.RecordGet(orig, "z").StrVal()
							return clone
						}()
						_ = c2_14_11
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{Call_Worker_score(func(record struct {
							active bool
							extra  int64
							id     int64
							name   string
						}) struct {
							active bool
							id     int64
							name   string
						} {
							return struct {
								active bool
								id     int64
								name   string
							}{record.active, record.id, record.name}
						}(a2_12_9)), int64(4435)}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
								actual   int64
								expected int64
							}{Call_Worker_score(func(record struct {
								active bool
								before float64
								id     int64
								name   string
								nested struct {
									retained string
								}
							}) struct {
								active bool
								id     int64
								name   string
							} {
								return struct {
									active bool
									id     int64
									name   string
								}{record.active, record.id, record.name}
							}(b2_13_10)), int64(1057)}), gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
									actual   int64
									expected int64
								}{Call_Worker_score(func(record struct {
									active bool
									id     int64
									name   string
									tail   bool
									values []int64
									z      string
								}) struct {
									active bool
									id     int64
									name   string
								} {
									return struct {
										active bool
										id     int64
										name   string
									}{record.active, record.id, record.name}
								}(c2_14_11)), int64(1057)}), gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_18_12 shape=App(Var) bindingType=Any
										__local_var_18_12 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_3)
										_ = __local_var_18_12
										__local_var_19_13 := gopurs_runtime.Apply(__local_var_18_12, gopurs_runtime.Value{})
										_ = __local_var_19_13
										__local_var_20_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_4), gopurs_runtime.Value{})
										_ = __local_var_20_14
										__local_var_21_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_5), gopurs_runtime.Value{})
										_ = __local_var_21_15
										return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___2008691784("", struct {
											actual struct {
												active bool
												extra  int64
												id     int64
												name   string
											}
											expected struct {
												active bool
												extra  int64
												id     int64
												name   string
											}
										}{func() struct {
											active bool
											extra  int64
											id     int64
											name   string
										} {
											orig := __local_var_19_13
											_ = orig
											clone := struct {
												active bool
												extra  int64
												id     int64
												name   string
											}{}
											clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
											clone.extra = gopurs_runtime.RecordGet(orig, "extra").IntVal
											clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
											clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
											return clone
										}(), func() struct {
											active bool
											extra  int64
											id     int64
											name   string
										} {
											orig := __local_var_6_6
											_ = orig
											clone := struct {
												active bool
												extra  int64
												id     int64
												name   string
											}{}
											clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
											clone.extra = gopurs_runtime.RecordGet(orig, "extra").IntVal
											clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
											clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
											return clone
										}()}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___2892873416("", struct {
												actual struct {
													active bool
													before float64
													id     int64
													name   string
													nested struct {
														retained string
													}
												}
												expected struct {
													active bool
													before float64
													id     int64
													name   string
													nested struct {
														retained string
													}
												}
											}{func() struct {
												active bool
												before float64
												id     int64
												name   string
												nested struct {
													retained string
												}
											} {
												orig := __local_var_20_14
												_ = orig
												clone := struct {
													active bool
													before float64
													id     int64
													name   string
													nested struct {
														retained string
													}
												}{}
												clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
												clone.before = gopurs_runtime.RecordGet(orig, "before").FloatVal()
												clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
												clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
												clone.nested = func() struct {
													retained string
												} {
													orig := gopurs_runtime.RecordGet(orig, "nested")
													_ = orig
													clone := struct {
														retained string
													}{}
													clone.retained = gopurs_runtime.RecordGet(orig, "retained").StrVal()
													return clone
												}()
												return clone
											}(), func() struct {
												active bool
												before float64
												id     int64
												name   string
												nested struct {
													retained string
												}
											} {
												orig := __local_var_7_7
												_ = orig
												clone := struct {
													active bool
													before float64
													id     int64
													name   string
													nested struct {
														retained string
													}
												}{}
												clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
												clone.before = gopurs_runtime.RecordGet(orig, "before").FloatVal()
												clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
												clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
												clone.nested = func() struct {
													retained string
												} {
													orig := gopurs_runtime.RecordGet(orig, "nested")
													_ = orig
													clone := struct {
														retained string
													}{}
													clone.retained = gopurs_runtime.RecordGet(orig, "retained").StrVal()
													return clone
												}()
												return clone
											}()}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1234170632("", struct {
													actual struct {
														active bool
														id     int64
														name   string
														tail   bool
														values []int64
														z      string
													}
													expected struct {
														active bool
														id     int64
														name   string
														tail   bool
														values []int64
														z      string
													}
												}{func() struct {
													active bool
													id     int64
													name   string
													tail   bool
													values []int64
													z      string
												} {
													orig := __local_var_21_15
													_ = orig
													clone := struct {
														active bool
														id     int64
														name   string
														tail   bool
														values []int64
														z      string
													}{}
													clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
													clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
													clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
													clone.tail = (gopurs_runtime.RecordGet(orig, "tail").IntVal) != (0)
													clone.values = func() []int64 {
														arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
														unboxed := make([]int64, len(arr))
														for i, v := range arr {
															unboxed[i] = v.IntVal
														}
														return unboxed
													}()
													clone.z = gopurs_runtime.RecordGet(orig, "z").StrVal()
													return clone
												}(), func() struct {
													active bool
													id     int64
													name   string
													tail   bool
													values []int64
													z      string
												} {
													orig := __local_var_8_8
													_ = orig
													clone := struct {
														active bool
														id     int64
														name   string
														tail   bool
														values []int64
														z      string
													}{}
													clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
													clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
													clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
													clone.tail = (gopurs_runtime.RecordGet(orig, "tail").IntVal) != (0)
													clone.values = func() []int64 {
														arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "values").UnsafePtr)
														unboxed := make([]int64, len(arr))
														for i, v := range arr {
															unboxed[i] = v.IntVal
														}
														return unboxed
													}()
													clone.z = gopurs_runtime.RecordGet(orig, "z").StrVal()
													return clone
												}()}), gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
														actual   int64
														expected int64
													}{a2_12_9.extra, int64(42)}), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___3008771400("", struct {
															actual struct {
																retained string
															}
															expected struct {
																retained string
															}
														}{b2_13_10.nested, struct {
															retained string
														}{"keep"}}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1772858129("", struct {
																actual   []int64
																expected []int64
															}{c2_14_11.values, []int64{int64(1), int64(2), int64(3)}}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																	actual   bool
																	expected bool
																}{c2_14_11.tail, false}), gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																		actual   string
																		expected string
																	}{c2_14_11.z, "last"}), gopurs_runtime.Func(func(_dollar___unused_29 gopurs_runtime.Value) gopurs_runtime.Value {
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
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_makeC(n_0_loop int64) struct {
	active bool
	id     int64
	name   string
	tail   bool
	values []int64
	z      string
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	return struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
	}{true, n_0, "alpha", false, []int64{int64(1), int64(2), int64(3)}, "last"}
}

func Call_Main_runC(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	return Call_Worker_score(func(record struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
	}) struct {
		active bool
		id     int64
		name   string
	} {
		return struct {
			active bool
			id     int64
			name   string
		}{record.active, record.id, record.name}
	}(Call_Main_makeC(n_0)))
}

func Call_Main_makeB(n_0_loop int64) struct {
	active bool
	before float64
	id     int64
	name   string
	nested struct {
		retained string
	}
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	return struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
		}
	}{false, 1.25, n_0, "beta", struct {
		retained string
	}{"keep"}}
}

func Call_Main_runB(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	return Call_Worker_score(func(record struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
		}
	}) struct {
		active bool
		id     int64
		name   string
	} {
		return struct {
			active bool
			id     int64
			name   string
		}{record.active, record.id, record.name}
	}(Call_Main_makeB(n_0)))
}

func Call_Main_makeA(n_0_loop int64) struct {
	active bool
	extra  int64
	id     int64
	name   string
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	return struct {
		active bool
		extra  int64
		id     int64
		name   string
	}{true, int64(42), n_0, "alpha"}
}

func Call_Main_runA(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	return Call_Worker_score(func(record struct {
		active bool
		extra  int64
		id     int64
		name   string
	}) struct {
		active bool
		id     int64
		name   string
	} {
		return struct {
			active bool
			id     int64
			name   string
		}{record.active, record.id, record.name}
	}(Call_Main_makeA(n_0)))
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

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1517968077(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	active bool
	id     int64
	name   string
	tail   bool
	values []int64
	z      string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
	}]{}
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

func Rebox_Main_1386611502_1994653624(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	active bool
	extra  int64
	id     int64
	name   string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		active bool
		extra  int64
		id     int64
		name   string
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_872707775(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	active bool
	before float64
	id     int64
	name   string
	nested struct {
		retained string
	}
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
		}
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_920949869(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	retained string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		retained string
	}]{}
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

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1517968077_1386611502(in *Constructor_Data_Show_Show[struct {
	active bool
	id     int64
	name   string
	tail   bool
	values []int64
	z      string
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
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

func Rebox_Main_1830338253_3790796878(in *Constructor_Data_Eq_Eq[struct {
	retained string
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1994653624_1386611502(in *Constructor_Data_Show_Show[struct {
	active bool
	extra  int64
	id     int64
	name   string
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2267788589_3790796878(in *Constructor_Data_Eq_Eq[struct {
	active bool
	id     int64
	name   string
	tail   bool
	values []int64
	z      string
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2543566495_3790796878(in *Constructor_Data_Eq_Eq[struct {
	active bool
	before float64
	id     int64
	name   string
	nested struct {
		retained string
	}
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
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

func Rebox_Main_2758774104_3790796878(in *Constructor_Data_Eq_Eq[struct {
	active bool
	extra  int64
	id     int64
	name   string
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
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

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1830338253(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	retained string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		retained string
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2267788589(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	active bool
	id     int64
	name   string
	tail   bool
	values []int64
	z      string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2543566495(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	active bool
	before float64
	id     int64
	name   string
	nested struct {
		retained string
	}
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
		}
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2758774104(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	active bool
	extra  int64
	id     int64
	name   string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		active bool
		extra  int64
		id     int64
		name   string
	}]{}
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

func Rebox_Main_3790796878_687527510(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_872707775_1386611502(in *Constructor_Data_Show_Show[struct {
	active bool
	before float64
	id     int64
	name   string
	nested struct {
		retained string
	}
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_920949869_1386611502(in *Constructor_Data_Show_Show[struct {
	retained string
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
