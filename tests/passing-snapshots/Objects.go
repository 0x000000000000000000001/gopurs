package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_typed gopurs_runtime.Value
var once_Main_typed sync.Once

func Get_Main_typed() gopurs_runtime.Value {
	once_Main_typed.Do(func() {
		cache_Main_typed = gopurs_runtime.RecordDict1("foo", gopurs_runtime.Float(0.0))
	})
	return cache_Main_typed
}

var cache_Main_test7 gopurs_runtime.Value
var once_Main_test7 sync.Once

func Get_Main_test7() gopurs_runtime.Value {
	once_Main_test7.Do(func() {
		cache_Main_test7 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test7(v_0_box)
		})
	})
	return cache_Main_test7
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Float(1.0)
	})
	return cache_Main_test6
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Float(1.0)
	})
	return cache_Main_test5
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Float(0.0)
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test2(x_0_box)
		})
	})
	return cache_Main_test2
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Float(1.0)
	})
	return cache_Main_test4
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test(x_0_box))
		})
	})
	return cache_Main_test
}

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
		cache_Main_g = gopurs_runtime.Float(2.0)
	})
	return cache_Main_g
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Float(1.0)
	})
	return cache_Main_f
}

var cache_Main_go__append gopurs_runtime.Value
var once_Main_go__append sync.Once

func Get_Main_go__append() gopurs_runtime.Value {
	once_Main_go__append.Do(func() {
		cache_Main_go__append = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_go__append(o_0_box)
		})
	})
	return cache_Main_go__append
}

var cache_Main_apTest gopurs_runtime.Value
var once_Main_apTest sync.Once

func Get_Main_apTest() gopurs_runtime.Value {
	once_Main_apTest.Do(func() {
		cache_Main_apTest = gopurs_runtime.RecordDict2("bar", "foo", gopurs_runtime.Float(1.0), gopurs_runtime.Str("Foo"))
	})
	return cache_Main_apTest
}

func Call_Main_test7(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.RecordGet(v_0, "b")
}

func Call_Main_test2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return gopurs_runtime.RecordGet(x_0, "!@#")
}

func Call_Main_test(x_0_loop gopurs_runtime.Value) float64 {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return ((gopurs_runtime.RecordGet(x_0, "foo").FloatVal()) + (gopurs_runtime.RecordGet(x_0, "bar").FloatVal())) + (1.0)
}

func Call_Main_go__append(o_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return gopurs_runtime.RecordDict2("bar", "foo", gopurs_runtime.Float(1.0), gopurs_runtime.RecordGet(o_0, "foo"))
}
