package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_greet gopurs_runtime.Value
var once_Main_greet sync.Once

func Get_Main_greet() gopurs_runtime.Value {
	once_Main_greet.Do(func() {
		cache_Main_greet = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_greet(v_0_box)
		})
	})
	return cache_Main_greet
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_greet(func() gopurs_runtime.Value {
			orig := struct {
				greeting string
				name     string
			}{"Hello", "World"}
			_ = orig
			return gopurs_runtime.RecordDict2("greeting", "name", gopurs_runtime.Str(orig.greeting), gopurs_runtime.Str(orig.name))
		}()), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_greet(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((((gopurs_runtime.RecordGet(v_0, "greeting").StrVal())+(", "))+(gopurs_runtime.RecordGet(v_0, "name").StrVal()))+(".")))
}
