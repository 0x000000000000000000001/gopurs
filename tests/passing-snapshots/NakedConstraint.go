package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Nil gopurs_runtime.Value
var once_Main_Nil sync.Once

func Get_Main_Nil() gopurs_runtime.Value {
	once_Main_Nil.Do(func() {
		cache_Main_Nil = gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))}
	})
	return cache_Main_Nil
}

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
		cache_Main_head = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_head(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](v_1_box)))
		})
	})
	return cache_Main_head
}

type Constructor_Main_Nil struct {
	Rc uint32
}

type Constructor_Main_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_Cons
}

func Call_Main_head(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Main_Cons) int64 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 *Constructor_Main_Cons = v_1_loop
	_ = v_1
	var __t0 int64
	{
		if v_1 != nil {
			__t0 = (v_1).V0.IntVal
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
	}
end_branch_0:
	return gopurs_runtime.Int(__t0).IntVal
}
