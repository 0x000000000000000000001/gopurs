package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_testZ gopurs_runtime.Value
var once_Main_testZ sync.Once

func Get_Main_testZ() gopurs_runtime.Value {
	once_Main_testZ.Do(func() {
		cache_Main_testZ = gopurs_runtime.Value{Type: 9, IntVal: int64(2089233011), UnsafePtr: nil}
	})
	return cache_Main_testZ
}

var cache_Main_testX gopurs_runtime.Value
var once_Main_testX sync.Once

func Get_Main_testX() gopurs_runtime.Value {
	once_Main_testX.Do(func() {
		cache_Main_testX = gopurs_runtime.Value{Type: 9, IntVal: int64(2407905777), UnsafePtr: nil}
	})
	return cache_Main_testX
}

var cache_Main_testFoo gopurs_runtime.Value
var once_Main_testFoo sync.Once

func Get_Main_testFoo() gopurs_runtime.Value {
	once_Main_testFoo.Do(func() {
		cache_Main_testFoo = gopurs_runtime.Int(0)
	})
	return cache_Main_testFoo
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
