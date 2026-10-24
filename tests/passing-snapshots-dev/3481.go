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
			orig := struct {
				gopurs_field_48 struct {
					gopurs_field_49 string
					/*gopurs:record-v1:1:0031*/
				}
				/*gopurs:record-v1:1:0030*/
			}{struct {
				gopurs_field_49 string
				/*gopurs:record-v1:1:0031*/
			}{"Done"}}
			_ = orig
			return gopurs_runtime.RecordDict1("0", func() gopurs_runtime.Value {
				orig := orig.gopurs_field_48
				_ = orig
				return gopurs_runtime.RecordDict1("1", gopurs_runtime.Str(orig.gopurs_field_49))
			}())
		}()
	})
	return cache_Main_message
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(func() struct {
			gopurs_field_48 struct {
				gopurs_field_49 string
				/*gopurs:record-v1:1:0031*/
			}
			/*gopurs:record-v1:1:0030*/
		} {
			orig := Get_Main_message()
			_ = orig
			clone := struct {
				gopurs_field_48 struct {
					gopurs_field_49 string
					/*gopurs:record-v1:1:0031*/
				}
				/*gopurs:record-v1:1:0030*/
			}{}
			clone.gopurs_field_48 = func() struct {
				gopurs_field_49 string
				/*gopurs:record-v1:1:0031*/
			} {
				orig := gopurs_runtime.RecordGet(orig, "0")
				_ = orig
				clone := struct {
					gopurs_field_49 string
					/*gopurs:record-v1:1:0031*/
				}{}
				clone.gopurs_field_49 = gopurs_runtime.RecordGet(orig, "1").StrVal()
				return clone
			}()
			return clone
		}().gopurs_field_48.gopurs_field_49))
	})
	return cache_Main_main
}
