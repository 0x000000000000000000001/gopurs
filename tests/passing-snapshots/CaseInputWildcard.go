package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Value{Type: 9, IntVal: int64(1409933510), UnsafePtr: nil}
	})
	return cache_Main_X
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Value{Type: 9, IntVal: int64(1682951303), UnsafePtr: nil}
	})
	return cache_Main_Y
}

var cache_Main_what gopurs_runtime.Value
var once_Main_what sync.Once

func Get_Main_what() gopurs_runtime.Value {
	once_Main_what.Do(func() {
		cache_Main_what = gopurs_runtime.Func3(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_what(uint32(x_0_box.IntVal), v_1_box.IntVal, (v1_2_box.IntVal) != (0))), UnsafePtr: nil}
		})
	})
	return cache_Main_what
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
}

type Constructor_Main_Y struct {
	Rc uint32
}

func Call_Main_what(x_0_loop uint32, v_1_loop int64, v1_2_loop bool) uint32 {
	var x_0 uint32 = x_0_loop
	_ = x_0
	var v_1 int64 = v_1_loop
	_ = v_1
	var v1_2 bool = v1_2_loop
	_ = v1_2
	var __t1 uint32
	{
		if ((v_1) == (0)) && (v1_2) {
			var __t0 uint32
			{
				if x_0 == 1409933510 {
					__t0 = 1409933510
					goto end_branch_0
				} else {

				}
			}
			{
				if x_0 == 1682951303 {
					__t0 = 1409933510
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = 1682951303
			}
		end_branch_0:
			__t1 = __t0
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = 1682951303
	}
end_branch_1:
	return __t1
}
