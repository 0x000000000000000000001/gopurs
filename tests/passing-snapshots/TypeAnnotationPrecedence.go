package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_appendAndLog gopurs_runtime.Value
var once_Main_appendAndLog sync.Once

func Get_Main_appendAndLog() gopurs_runtime.Value {
	once_Main_appendAndLog.Do(func() {
		cache_Main_appendAndLog = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_appendAndLog(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[string, string]](x_0_box))
		})
	})
	return cache_Main_appendAndLog
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_appendAndLog(x_0_loop *Constructor_Data_Tuple_Tuple[string, string]) gopurs_runtime.Value {
	var x_0 *Constructor_Data_Tuple_Tuple[string, string] = x_0_loop
	_ = x_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(((x_0).V0.StrVal())+((x_0).V1.StrVal())))
}
