package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(record_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 string
			{
				if (gopurs_runtime.RecordGet(record_0, "b").IntVal) != (0) {
					__t0 = (" b: ") + ("true")
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (" b: ") + ("false")
			}
		end_branch_0:
			return gopurs_runtime.Str((((((((("{ a: ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(record_0, "a")).StrVal())) + (",")) + (__t0)) + (", c: ")) + (gopurs_runtime.Apply(Get_Data_Show_showCharImpl(), gopurs_runtime.RecordGet(record_0, "c")).StrVal())) + (", e: ")) + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.RecordGet(record_0, "e")).StrVal())) + (" }"))
		})}))}
	})
	return cache_Main_showRecord
}

var cache_Main_showFFI gopurs_runtime.Value
var once_Main_showFFI sync.Once

func Get_Main_showFFI() gopurs_runtime.Value {
	once_Main_showFFI.Do(func() {
		cache_Main_showFFI = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showFFI(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box))
		})
	})
	return cache_Main_showFFI
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Main_showImpl(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showRecord()).V0), gopurs_runtime.RecordDict4("a", "b", "c", "e", gopurs_runtime.Int(1), gopurs_runtime.Bool(true), gopurs_runtime.Str("d"), gopurs_runtime.Float(4.0))).StrVal()))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_1
				_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Showing Int is correct"), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Main_showImpl(), Get_Data_Show_showIntImpl(), gopurs_runtime.Int(4)).StrVal()) == ("4"))), gopurs_runtime.Value{})
				_ = _dollar___unused_2_2
				_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Showing String is correct"), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Main_showImpl(), Get_Data_Show_showStringImpl(), gopurs_runtime.Str("string")).StrVal()) == ("\"string\""))), gopurs_runtime.Value{})
				_ = _dollar___unused_3_3
				_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Showing Record is correct"), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Main_showImpl(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showRecord()).V0), gopurs_runtime.RecordDict4("a", "b", "c", "e", gopurs_runtime.Int(1), gopurs_runtime.Bool(true), gopurs_runtime.Str("d"), gopurs_runtime.Float(4.0))).StrVal()) == ("{ a: 1, b: true, c: 'd', e: 4.0 }"))), gopurs_runtime.Value{})
				_ = _dollar___unused_4_4
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}

var cache_Main_showFFI__2742601362 gopurs_runtime.Value
var once_Main_showFFI__2742601362 sync.Once

func Get_Main_showFFI__2742601362() gopurs_runtime.Value {
	once_Main_showFFI__2742601362.Do(func() {
		cache_Main_showFFI__2742601362 = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showFFI__2742601362(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0_box))
		})
	})
	return cache_Main_showFFI__2742601362
}

func Call_Main_showFFI(dictShow_0_loop *Constructor_Data_Show_Show) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Apply(Get_Main_showImpl(), gopurs_runtime.Box(dictShow_0.V0))
}

func Call_Main_showFFI__2742601362(dictShow_0_loop *Constructor_Data_Show_Show) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Apply(Get_Main_showImpl(), gopurs_runtime.Box(dictShow_0.V0))
}

func Get_Main_showImpl() gopurs_runtime.Value {
	return _Gopurs_Main_ShowImpl
}
