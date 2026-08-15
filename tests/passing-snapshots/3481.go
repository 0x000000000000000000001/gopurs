package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_message gopurs_runtime.Value
var once_Main_message sync.Once

func Get_Main_message() gopurs_runtime.Value {
	once_Main_message.Do(func() {
		cache_Main_message = gopurs_runtime.RecordDict1("0", gopurs_runtime.RecordDict1("1", gopurs_runtime.Str("Done")))
	})
	return cache_Main_message
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_message(), "0"), "1").StrVal()))
	})
	return cache_Main_main
}
