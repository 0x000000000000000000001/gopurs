package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_f
}

var cache_Main_e gopurs_runtime.Value
var once_Main_e sync.Once

func Get_Main_e() gopurs_runtime.Value {
	once_Main_e.Do(func() {
		cache_Main_e = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_e
}

var cache_Main_d gopurs_runtime.Value
var once_Main_d sync.Once

func Get_Main_d() gopurs_runtime.Value {
	once_Main_d.Do(func() {
		cache_Main_d = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_d
}

var cache_Main_c gopurs_runtime.Value
var once_Main_c sync.Once

func Get_Main_c() gopurs_runtime.Value {
	once_Main_c.Do(func() {
		cache_Main_c = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_c
}

var cache_Main_b gopurs_runtime.Value
var once_Main_b sync.Once

func Get_Main_b() gopurs_runtime.Value {
	once_Main_b.Do(func() {
		cache_Main_b = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_b
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_a
}

type Constructor_Main_Proxy struct {
	Rc uint32
}
