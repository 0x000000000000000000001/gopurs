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

var cache_Main_F gopurs_runtime.Value
var once_Main_F sync.Once

func Get_Main_F() gopurs_runtime.Value {
	once_Main_F.Do(func() {
		cache_Main_F = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_F
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fproxy gopurs_runtime.Value
var once_Main_fproxy sync.Once

func Get_Main_fproxy() gopurs_runtime.Value {
	once_Main_fproxy.Do(func() {
		cache_Main_fproxy = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_fproxy(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_fproxy
}

var cache_Main_fproxy__2844018335 gopurs_runtime.Value
var once_Main_fproxy__2844018335 sync.Once

func Get_Main_fproxy__2844018335() gopurs_runtime.Value {
	once_Main_fproxy__2844018335.Do(func() {
		cache_Main_fproxy__2844018335 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_fproxy__2844018335(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_fproxy__2844018335
}

var cache_Main_fproxy__1744545119 gopurs_runtime.Value
var once_Main_fproxy__1744545119 sync.Once

func Get_Main_fproxy__1744545119() gopurs_runtime.Value {
	once_Main_fproxy__1744545119.Do(func() {
		cache_Main_fproxy__1744545119 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_fproxy__1744545119(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_fproxy__1744545119
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = gopurs_runtime.Func(func(v1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_a(uint32(v1_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_a
}

var cache_Main_b gopurs_runtime.Value
var once_Main_b sync.Once

func Get_Main_b() gopurs_runtime.Value {
	once_Main_b.Do(func() {
		cache_Main_b = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_b
}

var cache_Main_c gopurs_runtime.Value
var once_Main_c sync.Once

func Get_Main_c() gopurs_runtime.Value {
	once_Main_c.Do(func() {
		cache_Main_c = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_c
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_F[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_fproxy(v_0_loop uint32, v1_1_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 uint32 = v1_1_loop
	_ = v1_1
	return 227768594
}

func Call_Main_fproxy__2844018335(v_0_loop uint32, v1_1_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 uint32 = v1_1_loop
	_ = v1_1
	return 227768594
}

func Call_Main_fproxy__1744545119(v_0_loop uint32, v1_1_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 uint32 = v1_1_loop
	_ = v1_1
	return 227768594
}

func Call_Main_a(v1_0_loop uint32) uint32 {
	var v1_0 uint32 = v1_0_loop
	_ = v1_0
	return 227768594
}
