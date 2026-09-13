package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_x(a_0_box))
		})
	})
	return cache_Main_x
}

var cache_Main_x__3313475196 gopurs_runtime.Value
var once_Main_x__3313475196 sync.Once

func Get_Main_x__3313475196() gopurs_runtime.Value {
	once_Main_x__3313475196.Do(func() {
		cache_Main_x__3313475196 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_x__3313475196(a_0_box.IntVal))
		})
	})
	return cache_Main_x__3313475196
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("Test")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_x(a_0_loop gopurs_runtime.Value) string {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("Test")).StrVal()
}

func Call_Main_x__3313475196(a_0_loop int64) string {
x__3313475196:
	for {
		if false {
			continue x__3313475196
		}
		var a_0 int64 = a_0_loop
		_ = a_0
		return gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str("Test")).StrVal()
	}
}
