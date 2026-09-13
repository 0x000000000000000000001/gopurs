package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_otherwise gopurs_runtime.Value
var once_Main_otherwise sync.Once

func Get_Main_otherwise() gopurs_runtime.Value {
	once_Main_otherwise.Do(func() {
		cache_Main_otherwise = gopurs_runtime.Bool(false)
	})
	return cache_Main_otherwise
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("1.0")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}
