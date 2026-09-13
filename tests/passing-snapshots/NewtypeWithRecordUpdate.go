package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_NewType gopurs_runtime.Value
var once_Main_NewType sync.Once

func Get_Main_NewType() gopurs_runtime.Value {
	once_Main_NewType.Do(func() {
		cache_Main_NewType = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NewType(x_0_box)
		})
	})
	return cache_Main_NewType
}

var cache_Main_NewType__2196289690 gopurs_runtime.Value
var once_Main_NewType__2196289690 sync.Once

func Get_Main_NewType__2196289690() gopurs_runtime.Value {
	once_Main_NewType__2196289690.Do(func() {
		cache_Main_NewType__2196289690 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NewType__2196289690(x_0_box)
		})
	})
	return cache_Main_NewType__2196289690
}

var cache_Main_rec1 gopurs_runtime.Value
var once_Main_rec1 sync.Once

func Get_Main_rec1() gopurs_runtime.Value {
	once_Main_rec1.Do(func() {
		cache_Main_rec1 = func() gopurs_runtime.Value {
			orig := struct {
				a float64
				b float64
				c float64
			}{0.0, 0.0, 0.0}
			_ = orig
			return gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Float(orig.a), gopurs_runtime.Float(orig.b), gopurs_runtime.Float(orig.c))
		}()
	})
	return cache_Main_rec1
}

var cache_Main_rec2 gopurs_runtime.Value
var once_Main_rec2 sync.Once

func Get_Main_rec2() gopurs_runtime.Value {
	once_Main_rec2.Do(func() {
		cache_Main_rec2 = func() gopurs_runtime.Value {
			orig := func() struct {
				a float64
				b float64
				c float64
			} {
				clone := func() struct {
					a float64
					b float64
					c float64
				} {
					orig := Get_Main_rec1()
					_ = orig
					clone := struct {
						a float64
						b float64
						c float64
					}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").FloatVal()
					clone.b = gopurs_runtime.RecordGet(orig, "b").FloatVal()
					clone.c = gopurs_runtime.RecordGet(orig, "c").FloatVal()
					return clone
				}()
				clone.a = 1.0
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Float(orig.a), gopurs_runtime.Float(orig.b), gopurs_runtime.Float(orig.c))
		}()
	})
	return cache_Main_rec2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_NewType(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_NewType__2196289690(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
NewType__2196289690:
	for {
		if false {
			continue NewType__2196289690
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}
