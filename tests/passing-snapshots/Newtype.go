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
			return gopurs_runtime.Str(Call_Main_Thing(x_0_box.StrVal()))
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

var cache_Main_Box__75739034 gopurs_runtime.Value
var once_Main_Box__75739034 sync.Once

func Get_Main_Box__75739034() gopurs_runtime.Value {
	once_Main_Box__75739034.Do(func() {
		cache_Main_Box__75739034 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_Box__75739034(x_0_box.FloatVal()))
		})
	})
	return cache_Main_Box__75739034
}

var cache_Main_showThing gopurs_runtime.Value
var once_Main_showThing sync.Once

func Get_Main_showThing() gopurs_runtime.Value {
	once_Main_showThing.Do(func() {
		cache_Main_showThing = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(("Thing ") + (Data_Show_ShowStringImpl(v_0.StrVal())))
		})})))}
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
		cache_Main_logShow = gopurs_runtime.Apply(Get_Effect_Console_logShow(), Call_Main_showBox(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}))
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
			__local_var_0_0 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showThing()))), gopurs_runtime.Str("hello")), gopurs_runtime.Value{})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(Call_Effect_Console_logShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Main_showBox(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))})), gopurs_runtime.Float(42.0)), gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Effect_Console_logShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Main_showBox(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))})), gopurs_runtime.Float(Call_Main_apply(Get_Main_Box(), gopurs_runtime.Float(9000.0)).FloatVal())), gopurs_runtime.Value{})
			_ = __local_var_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_Thing(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Box(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Box__75739034(x_0_loop float64) float64 {
Box__75739034:
	for {
		if false {
			continue Box__75739034
		}
		var x_0 float64 = x_0_loop
		_ = x_0
		return x_0
	}
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

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[float64])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
