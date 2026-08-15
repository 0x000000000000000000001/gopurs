package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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

var cache_Main_use gopurs_runtime.Value
var once_Main_use sync.Once

func Get_Main_use() gopurs_runtime.Value {
	once_Main_use.Do(func() {
		cache_Main_use = gopurs_runtime.Str((("{ asdf: ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("asdf")).StrVal())) + (" }"))
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
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictReflectable_0, "reflectType"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
}
