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

var cache_Main_fib gopurs_runtime.Value
var once_Main_fib sync.Once

func Get_Main_fib() gopurs_runtime.Value {
	once_Main_fib.Do(func() {
		cache_Main_fib = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_fib(n_0_box.FloatVal()))
		})
	})
	return cache_Main_fib
}

func Call_Main_fib(n_0_loop float64) float64 {
fib:
	for {
		if false {
			continue fib
		}
		var n_0 float64 = n_0_loop
		_ = n_0
		var __t0 float64
		{
			if (n_0) == (0.0) {
				__t0 = 1.0
				goto end_branch_0
			} else {

			}
		}
		{
			if (n_0) == (1.0) {
				__t0 = 1.0
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (Call_Main_fib((n_0) - (1.0))) + (Call_Main_fib((n_0) - (2.0)))
		}
	end_branch_0:
		return __t0
	}
}
