package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_errorSemigroupMaybeMaybe gopurs_runtime.Value
var once_Main_errorSemigroupMaybeMaybe sync.Once

func Get_Main_errorSemigroupMaybeMaybe() gopurs_runtime.Value {
	once_Main_errorSemigroupMaybeMaybe.Do(func() {
		cache_Main_errorSemigroupMaybeMaybe = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_errorSemigroupMaybeMaybe
}

var cache_Main_errorSemigroupIdentityIde gopurs_runtime.Value
var once_Main_errorSemigroupIdentityIde sync.Once

func Get_Main_errorSemigroupIdentityIde() gopurs_runtime.Value {
	once_Main_errorSemigroupIdentityIde.Do(func() {
		cache_Main_errorSemigroupIdentityIde = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_errorSemigroupIdentityIde
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_ErrorSemigroup[T_o any, T_m any, T_w any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004500720072006f007200530065006d006900670072006f00750070
func init() {
	gopurs_runtime.StructGetters[246409291] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ErrorSemigroup[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_ErrorSemigroup: " + key)
		}
	}
}
