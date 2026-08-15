package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Cons gopurs_runtime.Value
var once_Main_Cons sync.Once

func Get_Main_Cons() gopurs_runtime.Value {
	once_Main_Cons.Do(func() {
		cache_Main_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](value1)}))}
			})
		})
	})
	return cache_Main_Cons
}

var cache_Main_step gopurs_runtime.Value
var once_Main_step sync.Once

func Get_Main_step() gopurs_runtime.Value {
	once_Main_step.Do(func() {
		cache_Main_step = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Call_Main_step(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](v_0_box)))}
		})
	})
	return cache_Main_step
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_head gopurs_runtime.Value
var once_Main_head sync.Once

func Get_Main_head() gopurs_runtime.Value {
	once_Main_head.Do(func() {
		cache_Main_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_head(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](xs_0_box))
		})
	})
	return cache_Main_head
}

var cache_Main_step__1567000704 gopurs_runtime.Value
var once_Main_step__1567000704 sync.Once

func Get_Main_step__1567000704() gopurs_runtime.Value {
	once_Main_step__1567000704.Do(func() {
		cache_Main_step__1567000704 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Call_Main_step__1567000704(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](v_0_box)))}
		})
	})
	return cache_Main_step__1567000704
}

type Constructor_Main_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_Cons
}

func Call_Main_step(v_0_loop *Constructor_Main_Cons) *Constructor_Main_Cons {
	var v_0 *Constructor_Main_Cons = v_0_loop
	_ = v_0
	return (v_0).V1
}

func Call_Main_head(xs_0_loop *Constructor_Main_Cons) gopurs_runtime.Value {
	var xs_0 *Constructor_Main_Cons = xs_0_loop
	_ = xs_0
	return ((xs_0).V1).V0
}

func Call_Main_step__1567000704(v_0_loop *Constructor_Main_Cons) *Constructor_Main_Cons {
	var v_0 *Constructor_Main_Cons = v_0_loop
	_ = v_0
	return (v_0).V1
}
