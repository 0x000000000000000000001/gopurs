package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_message gopurs_runtime.Value
var once_Main_message sync.Once
func Get_Main_message() gopurs_runtime.Value {
	once_Main_message.Do(func() {
		cache_Main_message = func() gopurs_runtime.Value {
				orig := func() *struct{
	0 *struct{
	1 string
}
} {
					orig := gopurs_runtime.RecordDict1("0", func() gopurs_runtime.Value {
				orig := func() *struct{
	1 string
} {
					orig := gopurs_runtime.RecordDict1("1", gopurs_runtime.Str("Done"))
					_ = orig
					clone := struct{
	1 string
}{}
					clone.1 = gopurs_runtime.RecordGet(orig, "1").StrVal()
					return &clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"1"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.1)})
				}())
					_ = orig
					clone := struct{
	0 *struct{
	1 string
}
}{}
					clone.0 = func() *struct{
	1 string
} {
					orig := gopurs_runtime.RecordGet(orig, "0")
					_ = orig
					clone := struct{
	1 string
}{}
					clone.1 = gopurs_runtime.RecordGet(orig, "1").StrVal()
					return &clone
				}()
					return &clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"0"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
				orig := orig.0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"1"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.1)})
				}()})
				}()
	})
	return cache_Main_message
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once
func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(func() *struct{
	0 *struct{
	1 string
}
} {
					orig := Get_Main_message()
					_ = orig
					clone := struct{
	0 *struct{
	1 string
}
}{}
					clone.0 = func() *struct{
	1 string
} {
					orig := gopurs_runtime.RecordGet(orig, "0")
					_ = orig
					clone := struct{
	1 string
}{}
					clone.1 = gopurs_runtime.RecordGet(orig, "1").StrVal()
					return &clone
				}()
					return &clone
				}().0.1))
	})
	return cache_Main_main
}




