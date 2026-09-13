package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_reflect gopurs_runtime.Value
var once_Main_reflect sync.Once

func Get_Main_reflect() gopurs_runtime.Value {
	once_Main_reflect.Do(func() {
		cache_Main_reflect = gopurs_runtime.Func(func(dictReflectable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_reflect(dictReflectable_0_box)
		})
	})
	return cache_Main_reflect
}

var cache_Main_reflect__824790448 gopurs_runtime.Value
var once_Main_reflect__824790448 sync.Once

func Get_Main_reflect__824790448() gopurs_runtime.Value {
	once_Main_reflect__824790448.Do(func() {
		cache_Main_reflect__824790448 = gopurs_runtime.Func(func(dictReflectable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_reflect__824790448(dictReflectable_0_box)
		})
	})
	return cache_Main_reflect__824790448
}

var cache_Main_use gopurs_runtime.Value
var once_Main_use sync.Once

func Get_Main_use() gopurs_runtime.Value {
	once_Main_use.Do(func() {
		cache_Main_use = gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("asdf")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})), "show"), func() gopurs_runtime.Value {
			orig := struct {
				asdf gopurs_runtime.Value
			}{Call_Main_reflect(gopurs_runtime.Value{Type: 9, IntVal: 19771322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("asdf")
			})}))})}
			_ = orig
			return gopurs_runtime.RecordDict1("asdf", orig.asdf)
		}()).StrVal())
	})
	return cache_Main_use
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_reflect(dictReflectable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictReflectable_0 gopurs_runtime.Value = dictReflectable_0_loop
	_ = dictReflectable_0
	return gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(gopurs_runtime.CoerceToStruct[Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]](dictReflectable_0)), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)), UnsafePtr: nil})
}

func Call_Main_reflect__824790448(dictReflectable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
reflect__824790448:
	for {
		if false {
			continue reflect__824790448
		}
		var dictReflectable_0 gopurs_runtime.Value = dictReflectable_0_loop
		_ = dictReflectable_0
		return gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(gopurs_runtime.CoerceToStruct[Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]](dictReflectable_0)), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)), UnsafePtr: nil})
	}
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
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
