package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Left gopurs_runtime.Value
var once_Main_Left sync.Once

func Get_Main_Left() gopurs_runtime.Value {
	once_Main_Left.Do(func() {
		cache_Main_Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 913637797, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left{1, value0}))}
		})
	})
	return cache_Main_Left
}

var cache_Main_Right gopurs_runtime.Value
var once_Main_Right sync.Once

func Get_Main_Right() gopurs_runtime.Value {
	once_Main_Right.Do(func() {
		cache_Main_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2535318782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Right{1, value0}))}
		})
	})
	return cache_Main_Right
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_functorEither gopurs_runtime.Value
var once_Main_functorEither sync.Once

func Get_Main_functorEither() gopurs_runtime.Value {
	once_Main_functorEither.Do(func() {
		cache_Main_functorEither = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if v1_1.Type == 9 && v1_1.IntVal == 913637797 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 913637797, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left{1, (*Constructor_Main_Left)(v1_1.UnsafePtr).V0}))}
						goto end_branch_0
					} else {

					}
				}
				{
					if v1_1.Type == 9 && v1_1.IntVal == 2535318782 {
						__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2535318782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Right{1, gopurs_runtime.Apply(v_0, (*Constructor_Main_Right)(v1_1.UnsafePtr).V0)}))}
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			})
		})}))}
	})
	return cache_Main_functorEither
}

type Constructor_Main_Left struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Right struct {
	Rc uint32
	V0 gopurs_runtime.Value
}
