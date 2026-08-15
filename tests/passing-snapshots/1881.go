package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_qux gopurs_runtime.Value
var once_Main_qux sync.Once

func Get_Main_qux() gopurs_runtime.Value {
	once_Main_qux.Do(func() {
		cache_Main_qux = gopurs_runtime.Int(3)
	})
	return cache_Main_qux
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Int(1)
	})
	return cache_Main_foo
}

var cache_Main_baz gopurs_runtime.Value
var once_Main_baz sync.Once

func Get_Main_baz() gopurs_runtime.Value {
	once_Main_baz.Do(func() {
		cache_Main_baz = gopurs_runtime.Int(3)
	})
	return cache_Main_baz
}

var cache_Main_bar gopurs_runtime.Value
var once_Main_bar sync.Once

func Get_Main_bar() gopurs_runtime.Value {
	once_Main_bar.Do(func() {
		cache_Main_bar = gopurs_runtime.Int(2)
	})
	return cache_Main_bar
}
