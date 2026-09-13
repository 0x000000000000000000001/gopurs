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

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(v_0_box))
		})
	})
	return cache_Main_g
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = func() gopurs_runtime.Value {
			var __t0 int64
			{
				if (gopurs_runtime.Apply(gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return y_0
				}), gopurs_runtime.Bool(true)).IntVal) != (0) {
					__t0 = gopurs_runtime.Apply(gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
						return y_0
					}), gopurs_runtime.Int(int64(0))).IntVal
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Apply(gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return y_0
				}), gopurs_runtime.Int(int64(1))).IntVal
			}
		end_branch_0:
			return gopurs_runtime.Int(__t0)
		}()
	})
	return cache_Main_test2
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(v_0_box)
		})
	})
	return cache_Main_f
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_f(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})), gopurs_runtime.Int(int64(1))).IntVal)
	})
	return cache_Main_test1
}

func Call_Main_g(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return y_1
	})).IntVal
}

func Call_Main_f(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, v_0)
}
