package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_testRunFn gopurs_runtime.Value
var once_Main_testRunFn sync.Once

func Get_Main_testRunFn() gopurs_runtime.Value {
	once_Main_testRunFn.Do(func() {
		cache_Main_testRunFn = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.UncurriedApp3(Get_Main_add3(), gopurs_runtime.Str("a"), gopurs_runtime.Str("b"), gopurs_runtime.Str("c"))
			_ = __local_var_0_0
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((__local_var_0_0.StrVal()) == ("abc"))), gopurs_runtime.Value{})
		})
	})
	return cache_Main_testRunFn
}

var cache_Main_testBothWays gopurs_runtime.Value
var once_Main_testBothWays sync.Once

func Get_Main_testBothWays() gopurs_runtime.Value {
	once_Main_testBothWays.Do(func() {
		cache_Main_testBothWays = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Int(int64(42))
			_ = __local_var_0_0
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((__local_var_0_0.IntVal) == (int64(42)))), gopurs_runtime.Value{})
		})
	})
	return cache_Main_testBothWays
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_testBothWays(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_testRunFn(), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}))
	})
	return cache_Main_main
}

func Get_Main_add3() gopurs_runtime.Value {
	return _Gopurs_Main_Add3
}
