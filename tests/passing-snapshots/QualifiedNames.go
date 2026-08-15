package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_identity(x_0_box.StrVal()))
		})
	})
	return cache_Main_identity
}

var cache_Main_either gopurs_runtime.Value
var once_Main_either sync.Once

func Get_Main_either() gopurs_runtime.Value {
	once_Main_either.Do(func() {
		cache_Main_either = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_either(v_0_box, v1_1_box, v2_2_box)
		})
	})
	return cache_Main_either
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal()))
	})
	return cache_Main_main
}

var cache_Main_either__1292901638 gopurs_runtime.Value
var once_Main_either__1292901638 sync.Once

func Get_Main_either__1292901638() gopurs_runtime.Value {
	once_Main_either__1292901638.Do(func() {
		cache_Main_either__1292901638 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_either__1292901638(v_0_box, v1_1_box, v2_2_box)
		})
	})
	return cache_Main_either__1292901638
}

func Call_Main_identity(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return gopurs_runtime.Str(x_0).StrVal()
}

func Call_Main_either(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t0 gopurs_runtime.Value
	{
		if v2_2.Type == 9 && v2_2.IntVal == 1485529257 {
			__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Either_Left)(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2.Type == 9 && v2_2.IntVal == 3726768370 {
			__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Either_Right)(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_either__1292901638(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t0 gopurs_runtime.Value
	{
		if v2_2.Type == 9 && v2_2.IntVal == 1485529257 {
			__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Either_Left)(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2.Type == 9 && v2_2.IntVal == 3726768370 {
			__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Either_Right)(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}
