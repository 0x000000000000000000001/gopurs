package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X{1, value0.IntVal}))}
		})
	})
	return cache_Main_X
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X)(nil))}
	})
	return cache_Main_Y
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_x(gopurs_runtime.CoerceToStruct[Constructor_Main_X](v_0_box)))
		})
	})
	return cache_Main_x
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_X struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_Y struct {
	Rc uint32
}

func Call_Main_x(v_0_loop *Constructor_Main_X) int64 {
	var v_0 *Constructor_Main_X = v_0_loop
	_ = v_0
	var __t1 int64
	{
		if v_0 == nil {
			__t1 = 0
			goto end_branch_1
		} else {

		}
	}
	{
		if v_0 != nil {
			var __t0 int64
			{
				if ((v_0).V0) == (1) {
					__t0 = 1
					goto end_branch_0
				} else {

				}
			}
			{
				if v_0 != nil {
					__t0 = 2
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
			}
		end_branch_0:
			__t1 = __t0
			goto end_branch_1
		} else {

		}
	}
	{
		if v_0 != nil {
			__t1 = 2
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
	}
end_branch_1:
	return __t1
}
