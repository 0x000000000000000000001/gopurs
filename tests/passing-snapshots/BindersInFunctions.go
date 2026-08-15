package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_snd gopurs_runtime.Value
var once_Main_snd sync.Once

func Get_Main_snd() gopurs_runtime.Value {
	once_Main_snd.Do(func() {
		cache_Main_snd = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_snd(_dollar___unused_0_box, func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v_1_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})
	})
	return cache_Main_snd
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result from 'snd'."), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_1
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}

var cache_Main_snd__2357233531 gopurs_runtime.Value
var once_Main_snd__2357233531 sync.Once

func Get_Main_snd__2357233531() gopurs_runtime.Value {
	once_Main_snd__2357233531.Do(func() {
		cache_Main_snd__2357233531 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_snd__2357233531(_dollar___unused_0_box, func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v_1_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})
	})
	return cache_Main_snd__2357233531
}

func Call_Main_snd(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 []gopurs_runtime.Value = v_1_loop
	_ = v_1
	var __t0 gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(v_1))).IntVal) == (2) {
			__t0 = gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v_1), 1)
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

func Call_Main_snd__2357233531(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 []gopurs_runtime.Value = v_1_loop
	_ = v_1
	var __t0 gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(v_1))).IntVal) == (2) {
			__t0 = gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v_1), 1)
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
