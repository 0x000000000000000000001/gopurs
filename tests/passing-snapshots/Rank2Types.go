package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test1(f_0_box))
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_forever gopurs_runtime.Value
var once_Main_forever sync.Once

func Get_Main_forever() gopurs_runtime.Value {
	once_Main_forever.Do(func() {
		cache_Main_forever = gopurs_runtime.Func2(func(bind_0_box gopurs_runtime.Value, action_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_forever(bind_0_box, action_1_box)
		})
	})
	return cache_Main_forever
}

var cache_Main_forever__2758867366 gopurs_runtime.Value
var once_Main_forever__2758867366 sync.Once

func Get_Main_forever__2758867366() gopurs_runtime.Value {
	once_Main_forever__2758867366.Do(func() {
		cache_Main_forever__2758867366 = gopurs_runtime.Func2(func(bind_0_box gopurs_runtime.Value, action_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_forever__2758867366(bind_0_box, action_1_box)
		})
	})
	return cache_Main_forever__2758867366
}

func Call_Main_test1(f_0_loop gopurs_runtime.Value) float64 {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Float(0.0)).FloatVal()
}

func Call_Main_forever(bind_0_loop gopurs_runtime.Value, action_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
forever:
	for {
		if false {
			continue forever
		}
		var bind_0 gopurs_runtime.Value = bind_0_loop
		_ = bind_0
		var action_1 gopurs_runtime.Value = action_1_loop
		_ = action_1
		return gopurs_runtime.Apply2(bind_0, action_1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_forever(bind_0, action_1)
		}))
	}
}

func Call_Main_forever__2758867366(bind_0_loop gopurs_runtime.Value, action_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var bind_0 gopurs_runtime.Value = bind_0_loop
	_ = bind_0
	var action_1 gopurs_runtime.Value = action_1_loop
	_ = action_1
	return gopurs_runtime.Apply2(bind_0, action_1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(bind_0, action_1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_forever(bind_0, action_1)
		}))
	}))
}
