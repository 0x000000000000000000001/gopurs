package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_utf8multibyte gopurs_runtime.Value
var once_Main_utf8multibyte sync.Once

func Get_Main_utf8multibyte() gopurs_runtime.Value {
	once_Main_utf8multibyte.Do(func() {
		cache_Main_utf8multibyte = gopurs_runtime.Str("Hello λ→ world!!")
	})
	return cache_Main_utf8multibyte
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
