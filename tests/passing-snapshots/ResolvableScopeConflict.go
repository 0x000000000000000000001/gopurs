package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_what gopurs_runtime.Value
var once_Main_what sync.Once

func Get_Main_what() gopurs_runtime.Value {
	once_Main_what.Do(func() {
		cache_Main_what = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_what((v_0_box.IntVal) != (0)))
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

func Call_Main_what(v_0_loop bool) int64 {
	var v_0 bool = v_0_loop
	_ = v_0
	var __t0 int64
	{
		if v_0 {
			__t0 = 1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = 3
	}
end_branch_0:
	return __t0
}
