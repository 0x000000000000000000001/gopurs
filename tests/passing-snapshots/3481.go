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
			orig := func() *struct {
				X_0 *struct {
					X_1 string
				}
			} {
				orig := gopurs_runtime.RecordDict1("0", func() gopurs_runtime.Value {
					orig := func() *struct {
						X_1 string
					} {
						orig := gopurs_runtime.RecordDict1("1", gopurs_runtime.Str("Done"))
						_ = orig
						clone := struct {
							X_1 string
						}{}
						clone.X_1 = gopurs_runtime.RecordGet(orig, "1").StrVal()
						return &clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict([]string{"1"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.X_1)})
				}())
				_ = orig
				clone := struct {
					X_0 *struct {
						X_1 string
					}
				}{}
				clone.X_0 = func() *struct {
					X_1 string
				} {
					orig := gopurs_runtime.RecordGet(orig, "0")
					_ = orig
					clone := struct {
						X_1 string
					}{}
					clone.X_1 = gopurs_runtime.RecordGet(orig, "1").StrVal()
					return &clone
				}()
				return &clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"0"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
				orig := orig.X_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"1"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.X_1)})
			}()})
		}()
	})
	return cache_Main_message
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(func() *struct {
			X_0 *struct {
				X_1 string
			}
		} {
			orig := Get_Main_message()
			_ = orig
			clone := struct {
				X_0 *struct {
					X_1 string
				}
			}{}
			clone.X_0 = func() *struct {
				X_1 string
			} {
				orig := gopurs_runtime.RecordGet(orig, "0")
				_ = orig
				clone := struct {
					X_1 string
				}{}
				clone.X_1 = gopurs_runtime.RecordGet(orig, "1").StrVal()
				return &clone
			}()
			return &clone
		}().X_0.X_1))
	})
	return cache_Main_main
}
