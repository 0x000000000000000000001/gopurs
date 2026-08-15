package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_y gopurs_runtime.Value
var once_Main_y sync.Once

func Get_Main_y() gopurs_runtime.Value {
	once_Main_y.Do(func() {
		cache_Main_y = gopurs_runtime.Float(gopurs_runtime.Int(1).FloatVal())
	})
	return cache_Main_y
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Float(gopurs_runtime.Int(1).FloatVal())
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
