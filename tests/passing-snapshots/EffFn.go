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
			str_0_0 := gopurs_runtime.UncurriedApp3(Get_Main_add3(), gopurs_runtime.Str("a"), gopurs_runtime.Str("b"), gopurs_runtime.Str("c"))
			_ = str_0_0
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((str_0_0.StrVal()) == ("abc"))), gopurs_runtime.Value{})
		})
	})
	return cache_Main_testRunFn
}

var cache_Main_testBothWays gopurs_runtime.Value
var once_Main_testBothWays sync.Once

func Get_Main_testBothWays() gopurs_runtime.Value {
	once_Main_testBothWays.Do(func() {
		cache_Main_testBothWays = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true))
		})
	})
	return cache_Main_testBothWays
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			_dollar___unused_0_0 := gopurs_runtime.Apply(Get_Main_testBothWays(), gopurs_runtime.Value{})
			_ = _dollar___unused_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(Get_Main_testRunFn(), gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Get_Main_add3() gopurs_runtime.Value {
	return _Gopurs_Main_Add3
}
