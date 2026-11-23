package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4))
	})
	return cache_Main_main
}

var cache_Main_fib gopurs_runtime.Value
var once_Main_fib sync.Once

func Get_Main_fib() gopurs_runtime.Value {
	once_Main_fib.Do(func() {
		cache_Main_fib = func() gopurs_runtime.Value {
			__st_ref0 := float64(1.0)
			_ = __st_ref0
			__st_ref1 := float64(1.0)
			_ = __st_ref1
			for {
				__st_value2 := float64(__st_ref0)
				_ = __st_value2
				if !((1000.0) > (__st_value2)) {
					break
				} else {

				}
				__st_value3 := float64(__st_ref0)
				_ = __st_value3
				__st_value4 := float64(__st_ref1)
				_ = __st_value4
				__st_ref1 = (__st_value3) + (__st_value4)
				__st_value5 := float64(__st_ref1)
				_ = __st_value5
				__st_ref0 = __st_value4
				_ = __st_ref0
			}
			__st_value6 := gopurs_runtime.Value{}
			_ = __st_value6
			return gopurs_runtime.Float(__st_ref1)
		}()
	})
	return cache_Main_fib
}
