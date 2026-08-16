package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
		})
	})
	return cache_Main_bind
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			message_0_0 := gopurs_runtime.Str("Done")
			_ = message_0_0
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(message_0_0.StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_bind__2601835655 gopurs_runtime.Value
var once_Main_bind__2601835655 sync.Once

func Get_Main_bind__2601835655() gopurs_runtime.Value {
	once_Main_bind__2601835655.Do(func() {
		cache_Main_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bind__2601835655(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
		})
	})
	return cache_Main_bind__2601835655
}

func Call_Main_bind(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_bind__2601835655(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}
