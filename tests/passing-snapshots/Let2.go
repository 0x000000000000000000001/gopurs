package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = func() gopurs_runtime.Value {
			var g_0_0_0 gopurs_runtime.Value
			_ = g_0_0_0
			var f_0_1_1 gopurs_runtime.Value
			_ = f_0_1_1
			g_0_0_0 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 bool
				{
					if (v_1.FloatVal()) == (0.0) {
						__t2 = true
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = (gopurs_runtime.Apply(f_0_1_1, gopurs_runtime.Float((v_1.FloatVal())-(1.0))).IntVal) != (0)
				}
			end_branch_2:
				return gopurs_runtime.Bool(__t2)
			})
			f_0_1_1 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 bool
				{
					if (v_1.FloatVal()) == (0.0) {
						__t3 = false
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = (gopurs_runtime.Apply(g_0_0_0, gopurs_runtime.Float((v_1.FloatVal())-(1.0))).IntVal) != (0)
				}
			end_branch_3:
				return gopurs_runtime.Bool(__t3)
			})
			return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0_1_1, gopurs_runtime.Float(1.0)).IntVal) != (0)) != (true))
		}()
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t1 string
			{
				if (gopurs_runtime.Bool((Get_Main_test().IntVal) != (0)).IntVal) != (0) {
					__t1 = "true"
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = "false"
			}
		end_branch_1:
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t1))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_2
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}
