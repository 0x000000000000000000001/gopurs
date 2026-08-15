package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_proxy1Add2Is3 gopurs_runtime.Value
var once_Main_proxy1Add2Is3 sync.Once

func Get_Main_proxy1Add2Is3() gopurs_runtime.Value {
	once_Main_proxy1Add2Is3.Do(func() {
		cache_Main_proxy1Add2Is3 = gopurs_runtime.Value{Type: 9, IntVal: int64(828635455), UnsafePtr: nil}
	})
	return cache_Main_proxy1Add2Is3
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
