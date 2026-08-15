package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_wildcard_prime_ gopurs_runtime.Value
var once_Main_wildcard_prime_ sync.Once

func Get_Main_wildcard_prime_() gopurs_runtime.Value {
	once_Main_wildcard_prime_.Do(func() {
		cache_Main_wildcard_prime_ = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_wildcard_prime_(v_0_box))
		})
	})
	return cache_Main_wildcard_prime_
}

var cache_Main_wildcard gopurs_runtime.Value
var once_Main_wildcard sync.Once

func Get_Main_wildcard() gopurs_runtime.Value {
	once_Main_wildcard.Do(func() {
		cache_Main_wildcard = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wildcard(v_0_box)
		})
	})
	return cache_Main_wildcard
}

var cache_Main_quux_prime_ gopurs_runtime.Value
var once_Main_quux_prime_ sync.Once

func Get_Main_quux_prime_() gopurs_runtime.Value {
	once_Main_quux_prime_.Do(func() {
		cache_Main_quux_prime_ = gopurs_runtime.RecordDict5("q", "q'", "x", "y", "z", gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0))
	})
	return cache_Main_quux_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_id_prime_ gopurs_runtime.Value
var once_Main_id_prime_ sync.Once

func Get_Main_id_prime_() gopurs_runtime.Value {
	once_Main_id_prime_.Do(func() {
		cache_Main_id_prime_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_id_prime_(x_0_box)
		})
	})
	return cache_Main_id_prime_
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0))
	})
	return cache_Main_foo
}

var cache_Main_foo_prime_ gopurs_runtime.Value
var once_Main_foo_prime_ sync.Once

func Get_Main_foo_prime_() gopurs_runtime.Value {
	once_Main_foo_prime_.Do(func() {
		cache_Main_foo_prime_ = Get_Main_foo()
	})
	return cache_Main_foo_prime_
}

var cache_Main_quux gopurs_runtime.Value
var once_Main_quux sync.Once

func Get_Main_quux() gopurs_runtime.Value {
	once_Main_quux.Do(func() {
		cache_Main_quux = gopurs_runtime.RecordDict5("f", "q", "x", "y", "z", Get_Main_foo(), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0))
	})
	return cache_Main_quux
}

var cache_Main_baz gopurs_runtime.Value
var once_Main_baz sync.Once

func Get_Main_baz() gopurs_runtime.Value {
	once_Main_baz.Do(func() {
		cache_Main_baz = gopurs_runtime.RecordDict4("w", "x", "y", "z", gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0))
	})
	return cache_Main_baz
}

var cache_Main_bar gopurs_runtime.Value
var once_Main_bar sync.Once

func Get_Main_bar() gopurs_runtime.Value {
	once_Main_bar.Do(func() {
		cache_Main_bar = gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0))
	})
	return cache_Main_bar
}

var cache_Main_bar_prime_ gopurs_runtime.Value
var once_Main_bar_prime_ sync.Once

func Get_Main_bar_prime_() gopurs_runtime.Value {
	once_Main_bar_prime_.Do(func() {
		cache_Main_bar_prime_ = Get_Main_bar()
	})
	return cache_Main_bar_prime_
}

func Call_Main_wildcard_prime_(v_0_loop gopurs_runtime.Value) float64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.RecordGet(v_0, "q").FloatVal()
}

func Call_Main_wildcard(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.RecordDict4("w", "x", "y", "z", gopurs_runtime.Float(gopurs_runtime.RecordGet(v_0, "w").FloatVal()), gopurs_runtime.Float(gopurs_runtime.RecordGet(v_0, "w").FloatVal()), gopurs_runtime.Float(gopurs_runtime.RecordGet(v_0, "w").FloatVal()), gopurs_runtime.Float(gopurs_runtime.RecordGet(v_0, "w").FloatVal()))
}

func Call_Main_id_prime_(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
