package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqT gopurs_runtime.Value
var once_Main_eqT sync.Once

func Get_Main_eqT() gopurs_runtime.Value {
	once_Main_eqT.Do(func() {
		cache_Main_eqT = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Main_Just[int64]]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 bool
				{
					if x_0.Type == 9 && x_0.IntVal == 3271839782 && x_0.UnsafePtr != nil {
						__t0 = (y_1.Type == 9 && y_1.IntVal == 3271839782 && y_1.UnsafePtr != nil) && (((*Constructor_Main_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0.IntVal) == ((*Constructor_Main_Just[gopurs_runtime.Value])(y_1.UnsafePtr).V0.IntVal))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (x_0.Type == 9 && x_0.IntVal == 3271839782 && x_0.UnsafePtr == nil) && (y_1.Type == 9 && y_1.IntVal == 3271839782 && y_1.UnsafePtr == nil)
				}
			end_branch_0:
				return gopurs_runtime.Bool(__t0)
			})
		})}))}
	})
	return cache_Main_eqT
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}
