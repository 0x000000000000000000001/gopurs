package Data_String_CodePoints
import "gopurs/output/gopurs_runtime"

func dummyFn(name string) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		panic("Not implemented: " + name)
	})
}
var X_UnsafeCodePointAt0 = dummyFn("X_UnsafeCodePointAt0")
var X_CodePointAt = dummyFn("X_CodePointAt")
var X_CountPrefix = dummyFn("X_CountPrefix")
var X_FromCodePointArray = dummyFn("X_FromCodePointArray")
var X_Singleton = dummyFn("X_Singleton")
var X_Take = dummyFn("X_Take")
var X_ToCodePointArray = dummyFn("X_ToCodePointArray")
