package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = func() gopurs_runtime.Value {
			orig := struct {
				baz string
			}{"baz"}
			_ = orig
			return gopurs_runtime.RecordDict1("baz", gopurs_runtime.Str(orig.baz))
		}()
	})
	return cache_Main_x
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_blah gopurs_runtime.Value
var once_Main_blah sync.Once

func Get_Main_blah() gopurs_runtime.Value {
	once_Main_blah.Do(func() {
		cache_Main_blah = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_blah(x1_0_box)
		})
	})
	return cache_Main_blah
}

var cache_Main_blah__2107698554 gopurs_runtime.Value
var once_Main_blah__2107698554 sync.Once

func Get_Main_blah__2107698554() gopurs_runtime.Value {
	once_Main_blah__2107698554.Do(func() {
		cache_Main_blah__2107698554 = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_blah__2107698554(x1_0_box)
		})
	})
	return cache_Main_blah__2107698554
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = func() gopurs_runtime.Value {
			orig := func() struct {
				baz string
			} {
				orig := func() gopurs_runtime.Value {
					orig := func() struct {
						baz string
					} {
						clone := func() struct {
							baz string
						} {
							orig := Get_Main_x()
							_ = orig
							clone := struct {
								baz string
							}{}
							clone.baz = gopurs_runtime.RecordGet(orig, "baz").StrVal()
							return clone
						}()
						clone.baz = "blah"
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict1("baz", gopurs_runtime.Str(orig.baz))
				}()
				_ = orig
				clone := struct {
					baz string
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").StrVal()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict1("baz", gopurs_runtime.Str(orig.baz))
		}()
	})
	return cache_Main_test
}

func Call_Main_blah(x1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x1_0 gopurs_runtime.Value = x1_0_loop
	_ = x1_0
	return x1_0
}

func Call_Main_blah__2107698554(x1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
blah__2107698554:
	for {
		if false {
			continue blah__2107698554
		}
		var x1_0 gopurs_runtime.Value = x1_0_loop
		_ = x1_0
		return x1_0
	}
}
