package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Thing gopurs_runtime.Value
var once_Main_Thing sync.Once

func Get_Main_Thing() gopurs_runtime.Value {
	once_Main_Thing.Do(func() {
		cache_Main_Thing = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Thing(x_0_box)
		})
	})
	return cache_Main_Thing
}

var cache_Main_Box gopurs_runtime.Value
var once_Main_Box sync.Once

func Get_Main_Box() gopurs_runtime.Value {
	once_Main_Box.Do(func() {
		cache_Main_Box = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Box(x_0_box)
		})
	})
	return cache_Main_Box
}

var cache_Main_showThing gopurs_runtime.Value
var once_Main_showThing sync.Once

func Get_Main_showThing() gopurs_runtime.Value {
	once_Main_showThing.Do(func() {
		cache_Main_showThing = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(("Thing ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal()))
		})}))}
	})
	return cache_Main_showThing
}

var cache_Main_showBox gopurs_runtime.Value
var once_Main_showBox sync.Once

func Get_Main_showBox() gopurs_runtime.Value {
	once_Main_showBox.Do(func() {
		cache_Main_showBox = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showBox(dictShow_0_box)
		})
	})
	return cache_Main_showBox
}

var cache_Main_logShow gopurs_runtime.Value
var once_Main_logShow sync.Once

func Get_Main_logShow() gopurs_runtime.Value {
	once_Main_logShow.Do(func() {
		cache_Main_logShow = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=LitRecord bindingType=(Record (Row [show: (Func [(TypeVar a)] String)] Any))
			__local_var_0_0 := gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(("Box ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), v_0).StrVal()))
			}))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "show"), a_1).StrVal()))
			})
		}()
	})
	return cache_Main_logShow
}

var cache_Main_apply gopurs_runtime.Value
var once_Main_apply sync.Once

func Get_Main_apply() gopurs_runtime.Value {
	once_Main_apply.Do(func() {
		cache_Main_apply = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_apply(f_0_box, x_1_box)
		})
	})
	return cache_Main_apply
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("Thing ")+(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(gopurs_runtime.Str("hello").StrVal())).StrVal())))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("Box ")+(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(42.0)).StrVal()))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("Box ")+(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.Float(9000.0).FloatVal())).StrVal()))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_Thing(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Box(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_showBox(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(("Box ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()))
	})}))}
}

func Call_Main_apply(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(f_0, x_1)
}
