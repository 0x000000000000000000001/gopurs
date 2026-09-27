package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(go__const_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(go__const_0_box, v_1_box)
		})
	})
	return cache_Main_test
}

var cache_Main_test__1801063706 gopurs_runtime.Value
var once_Main_test__1801063706 sync.Once

func Get_Main_test__1801063706() gopurs_runtime.Value {
	once_Main_test__1801063706.Do(func() {
		cache_Main_test__1801063706 = gopurs_runtime.Func2(func(go__const_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test__1801063706(go__const_0_box.StrVal(), func() struct {
			} {
				orig := v_1_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}()))
		})
	})
	return cache_Main_test__1801063706
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Not done"), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_test(go__const_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var go__const_0 gopurs_runtime.Value = go__const_0_loop
	_ = go__const_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return go__const_0
}

func Call_Main_test__1801063706(go__const_0_loop string, v_1_loop struct {
}) string {
test__1801063706:
	for {
		if false {
			continue test__1801063706
		}
		var go__const_0 string = go__const_0_loop
		_ = go__const_0
		var v_1 struct {
		} = v_1_loop
		_ = v_1
		return go__const_0
	}
}
