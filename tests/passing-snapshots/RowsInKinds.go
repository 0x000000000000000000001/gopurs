package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_P gopurs_runtime.Value
var once_Main_P sync.Once

func Get_Main_P() gopurs_runtime.Value {
	once_Main_P.Do(func() {
		cache_Main_P = gopurs_runtime.Value{Type: 9, IntVal: int64(3253896398), UnsafePtr: nil}
	})
	return cache_Main_P
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_P struct {
	Rc uint32
}
