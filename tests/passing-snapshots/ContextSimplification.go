package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_shout gopurs_runtime.Value
var once_Main_shout sync.Once

func Get_Main_shout() gopurs_runtime.Value {
	once_Main_shout.Do(func() {
		cache_Main_shout = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_shout(dictShow_0_box, x_1_box)
		})
	})
	return cache_Main_shout
}

var cache_Main_usesShowTwice gopurs_runtime.Value
var once_Main_usesShowTwice sync.Once

func Get_Main_usesShowTwice() gopurs_runtime.Value {
	once_Main_usesShowTwice.Do(func() {
		cache_Main_usesShowTwice = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_usesShowTwice(dictShow_0_box, v_1_box)
		})
	})
	return cache_Main_usesShowTwice
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Call_Main_usesShowTwice(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Data_Show_showString()))}, gopurs_runtime.Bool(true)), gopurs_runtime.Str("Test"))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_shout(dictShow_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), x_1).StrVal())+("!")))
}

func Call_Main_usesShowTwice(dictShow_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var __t0 gopurs_runtime.Value
	{
		if (v_1.IntVal) != (0) {
			__t0 = gopurs_runtime.Apply(Get_Main_shout(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](dictShow_0))})
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), a_2).StrVal()))
		})
	}
end_branch_0:
	return __t0
}
