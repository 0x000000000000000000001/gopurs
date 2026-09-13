package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_doIt gopurs_runtime.Value
var once_Main_doIt sync.Once

func Get_Main_doIt() gopurs_runtime.Value {
	once_Main_doIt.Do(func() {
		cache_Main_doIt = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_doIt
}

var cache_Main_set gopurs_runtime.Value
var once_Main_set sync.Once

func Get_Main_set() gopurs_runtime.Value {
	once_Main_set.Do(func() {
		cache_Main_set = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Testing...")), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		}))
	})
	return cache_Main_set
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Get_Main_set(), gopurs_runtime.Value{})
			_ = __local_var_0_0
			var __t1 gopurs_runtime.Value
			{
				if (__local_var_0_0.IntVal) != (0) {
					__t1 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("Failed"))
			}
		end_branch_1:
			return gopurs_runtime.Apply(__t1, gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}
