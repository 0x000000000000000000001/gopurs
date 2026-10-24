package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main__foo gopurs_runtime.Value
var once_Main__foo sync.Once

func Get_Main__foo() gopurs_runtime.Value {
	once_Main__foo.Do(func() {
		cache_Main__foo = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main__foo
}

var cache_Main__bar gopurs_runtime.Value
var once_Main__bar sync.Once

func Get_Main__bar() gopurs_runtime.Value {
	once_Main__bar.Do(func() {
		cache_Main__bar = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main__bar
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
