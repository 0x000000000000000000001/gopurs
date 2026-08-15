package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_k gopurs_runtime.Value
var once_Main_k sync.Once

func Get_Main_k() gopurs_runtime.Value {
	once_Main_k.Do(func() {
		cache_Main_k = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_k(x_0_box.StrVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_k
}

var cache_Main_iterate gopurs_runtime.Value
var once_Main_iterate sync.Once

func Get_Main_iterate() gopurs_runtime.Value {
	once_Main_iterate.Do(func() {
		cache_Main_iterate = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_iterate(v_0_box.FloatVal(), v1_1_box, v2_2_box)
		})
	})
	return cache_Main_iterate
}

var cache_Main_iterate__1459115933 gopurs_runtime.Value
var once_Main_iterate__1459115933 sync.Once

func Get_Main_iterate__1459115933() gopurs_runtime.Value {
	once_Main_iterate__1459115933.Do(func() {
		cache_Main_iterate__1459115933 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_iterate__1459115933(v_0_box.FloatVal(), v1_1_box, v2_2_box)
		})
	})
	return cache_Main_iterate__1459115933
}

func Call_Main_k(x_0_loop string, y_1_loop float64) string {
	var x_0 string = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	return x_0
}

func Call_Main_iterate(v_0_loop float64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
iterate:
	for {
		if false {
			continue iterate
		}
		var v_0 float64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (0.0) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (1.0)
			v1_1_loop = v1_1
			v2_2_loop = gopurs_runtime.Apply(v1_1, v2_2)
			continue iterate
			__t0 = gopurs_runtime.Value{}
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_iterate__1459115933(v_0_loop float64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t0 gopurs_runtime.Value
	{
		if (v_0) == (0.0) {
			__t0 = v2_2
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Call_Main_iterate((v_0)-(1.0), v1_1, gopurs_runtime.Apply(v1_1, v2_2))
	}
end_branch_0:
	return __t0
}
