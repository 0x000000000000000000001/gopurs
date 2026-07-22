package CustomAssert

import (
	"fmt"
	"gopurs/output/gopurs_runtime"
)

var AssertThrowsImpl = gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) (ret gopurs_runtime.Value) {
			defer func() {
				if r := recover(); r != nil {
					ret = gopurs_runtime.Str(fmt.Sprint(r))
				}
			}()
			
			gopurs_runtime.Apply(f, arg)
			panic("Assertion failed: An error should have been thrown")
		})
	})
})
