package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Just
}

var cache_Main_Just__1283418637 gopurs_runtime.Value
var once_Main_Just__1283418637 sync.Once

func Get_Main_Just__1283418637() gopurs_runtime.Value {
	once_Main_Just__1283418637.Do(func() {
		cache_Main_Just__1283418637 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2656528265_4188394610(Call_Main_Just__1283418637(__eta_norm_0_0_box.IntVal)))}
		})
	})
	return cache_Main_Just__1283418637
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2656528265_4188394610(Rebox_Main_4188394610_2656528265(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}))))}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2656528265_4188394610((&Constructor_Main_Just[int64]{1, int64(10)})))}
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_Just__1283418637(__eta_norm_0_0_loop int64) *Constructor_Main_Just[int64] {
Just__1283418637:
	for {
		if false {
			continue Just__1283418637
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_Just[int64]{1, __eta_norm_0_0})
	}
}

func Rebox_Main_2656528265_4188394610(in *Constructor_Main_Just[int64]) *Constructor_Main_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_4188394610_2656528265(in *Constructor_Main_Just[gopurs_runtime.Value]) *Constructor_Main_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}
