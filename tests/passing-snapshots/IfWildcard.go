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
		cache_Main_what = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_what((v_0_box.IntVal) != (0))), UnsafePtr: nil}
		})
	})
	return cache_Main_what
}

var cache_Main_cond gopurs_runtime.Value
var once_Main_cond sync.Once

func Get_Main_cond() gopurs_runtime.Value {
	once_Main_cond.Do(func() {
		cache_Main_cond = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cond((v_0_box.IntVal) != (0), v1_1_box, v2_2_box)
		})
	})
	return cache_Main_cond
}

var cache_Main_cond__1906969362 gopurs_runtime.Value
var once_Main_cond__1906969362 sync.Once

func Get_Main_cond__1906969362() gopurs_runtime.Value {
	once_Main_cond__1906969362.Do(func() {
		cache_Main_cond__1906969362 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cond__1906969362((v_0_box.IntVal) != (0), v1_1_box, v2_2_box)
		})
	})
	return cache_Main_cond__1906969362
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

func Call_Main_what(v_0_loop bool) uint32 {
	var v_0 bool = v_0_loop
	_ = v_0
	var __t0 uint32
	{
		if v_0 {
			__t0 = 1409933510
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = 1682951303
	}
end_branch_0:
	return __t0
}

func Call_Main_cond(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 bool = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t0 gopurs_runtime.Value
	{
		if v_0 {
			__t0 = v1_1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = v2_2
	}
end_branch_0:
	return __t0
}

func Call_Main_cond__1906969362(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 bool = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t0 gopurs_runtime.Value
	{
		if v_0 {
			__t0 = v1_1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = v2_2
	}
end_branch_0:
	return __t0
}
