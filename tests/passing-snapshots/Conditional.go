package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_not gopurs_runtime.Value
var once_Main_not sync.Once

func Get_Main_not() gopurs_runtime.Value {
	once_Main_not.Do(func() {
		cache_Main_not = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_not((x_0_box.IntVal) != (0)))
		})
	})
	return cache_Main_not
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fns gopurs_runtime.Value
var once_Main_fns sync.Once

func Get_Main_fns() gopurs_runtime.Value {
	once_Main_fns.Do(func() {
		cache_Main_fns = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fns(f_0_box)
		})
	})
	return cache_Main_fns
}

func Call_Main_not(x_0_loop bool) bool {
	var x_0 bool = x_0_loop
	_ = x_0
	return (x_0) != (true)
}

func Call_Main_fns(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var __t0 gopurs_runtime.Value
	{
		if (gopurs_runtime.Apply(f_0, gopurs_runtime.Bool(true)).IntVal) != (0) {
			__t0 = f_0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((x_1.IntVal) != (0))
		})
	}
end_branch_0:
	return __t0
}
