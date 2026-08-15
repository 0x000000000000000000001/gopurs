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
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Hello, World."))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_1
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}

func Call_Main_greet(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((((gopurs_runtime.RecordGet(v_0, "greeting").StrVal())+(", "))+(gopurs_runtime.RecordGet(v_0, "name").StrVal()))+(".")))
}
