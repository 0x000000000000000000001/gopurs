package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_inner gopurs_runtime.Value
var once_Main_inner sync.Once

func Get_Main_inner() gopurs_runtime.Value {
	once_Main_inner.Do(func() {
		cache_Main_inner = gopurs_runtime.Float(0.0)
	})
	return cache_Main_inner
}

var cache_Main_outer gopurs_runtime.Value
var once_Main_outer sync.Once

func Get_Main_outer() gopurs_runtime.Value {
	once_Main_outer.Do(func() {
		cache_Main_outer = gopurs_runtime.RecordDict1("inner", gopurs_runtime.Float(0.0))
	})
	return cache_Main_outer
}
