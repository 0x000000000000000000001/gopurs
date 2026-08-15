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

var cache_Main_Pair gopurs_runtime.Value
var once_Main_Pair sync.Once

func Get_Main_Pair() gopurs_runtime.Value {
	once_Main_Pair.Do(func() {
		cache_Main_Pair = gopurs_runtime.Value{Type: 9, IntVal: int64(893478516), UnsafePtr: nil}
	})
	return cache_Main_Pair
}

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Value{Type: 9, IntVal: int64(4250879068), UnsafePtr: nil}
	})
	return cache_Main_B
}

var cache_Main_A gopurs_runtime.Value
var once_Main_A sync.Once

func Get_Main_A() gopurs_runtime.Value {
	once_Main_A.Do(func() {
		cache_Main_A = gopurs_runtime.Value{Type: 9, IntVal: int64(4219254943), UnsafePtr: nil}
	})
	return cache_Main_A
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_k gopurs_runtime.Value
var once_Main_k sync.Once

func Get_Main_k() gopurs_runtime.Value {
	once_Main_k.Do(func() {
		cache_Main_k = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_k(uint32(v_0_box.IntVal)))
		})
	})
	return cache_Main_k
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Int(42)
	})
	return cache_Main_test
}

type Constructor_Main_Proxy struct {
	Rc uint32
}

type Constructor_Main_Pair struct {
	Rc uint32
}

type Constructor_Main_B struct {
	Rc uint32
}

type Constructor_Main_A struct {
	Rc uint32
}

func Call_Main_k(v_0_loop uint32) int64 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 42
}
