package Test_Assert

import (
	"gopurs/output/gopurs_runtime"
)

var AssertImpl = gopurs_runtime.Func(func(messageVal gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(successVal gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			message := messageVal.StrVal
			success := successVal.IntVal != 0
			if !success {
				panic(message)
			}
			return gopurs_runtime.Value{}
		})
	})
})

var CheckThrows = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) (res gopurs_runtime.Value) {
		res = gopurs_runtime.Int(0)
		defer func() {
			if r := recover(); r != nil {
				res = gopurs_runtime.Int(1)
			}
		}()
		gopurs_runtime.Apply(fn, gopurs_runtime.Value{})
		return res
	})
})
