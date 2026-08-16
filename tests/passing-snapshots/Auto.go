package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Auto gopurs_runtime.Value
var once_Main_Auto sync.Once

func Get_Main_Auto() gopurs_runtime.Value {
	once_Main_Auto.Do(func() {
		cache_Main_Auto = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Auto
}

var cache_Main_run gopurs_runtime.Value
var once_Main_run sync.Once

func Get_Main_run() gopurs_runtime.Value {
	once_Main_run.Do(func() {
		cache_Main_run = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_run(s_0_box, i_1_box)
		})
	})
	return cache_Main_run
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_exists gopurs_runtime.Value
var once_Main_exists sync.Once

func Get_Main_exists() gopurs_runtime.Value {
	once_Main_exists.Do(func() {
		cache_Main_exists = gopurs_runtime.Func3(func(state_0_box gopurs_runtime.Value, step_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_exists(state_0_box, step_1_box, f_2_box)
		})
	})
	return cache_Main_exists
}

type Constructor_Main_Auto struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_run(s_0_loop gopurs_runtime.Value, i_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var s_0 gopurs_runtime.Value = s_0_loop
	_ = s_0
	var i_1 gopurs_runtime.Value = i_1_loop
	_ = i_1
	return gopurs_runtime.Apply(s_0, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(a_2, "step"), gopurs_runtime.RecordGet(a_2, "state"), i_1)
	}))
}

func Call_Main_exists(state_0_loop gopurs_runtime.Value, step_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var state_0 gopurs_runtime.Value = state_0_loop
	_ = state_0
	var step_1 gopurs_runtime.Value = step_1_loop
	_ = step_1
	var f_2 gopurs_runtime.Value = f_2_loop
	_ = f_2
	return gopurs_runtime.Apply(f_2, gopurs_runtime.RecordDict2("state", "step", state_0, step_1))
}
