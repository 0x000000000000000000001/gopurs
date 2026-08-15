package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_testY gopurs_runtime.Value
var once_Main_testY sync.Once

func Get_Main_testY() gopurs_runtime.Value {
	once_Main_testY.Do(func() {
		cache_Main_testY = gopurs_runtime.Value{Type: 9, IntVal: int64(314700016), UnsafePtr: nil}
	})
	return cache_Main_testY
}

var cache_Main_testX gopurs_runtime.Value
var once_Main_testX sync.Once

func Get_Main_testX() gopurs_runtime.Value {
	once_Main_testX.Do(func() {
		cache_Main_testX = gopurs_runtime.Value{Type: 9, IntVal: int64(2407905777), UnsafePtr: nil}
	})
	return cache_Main_testX
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
