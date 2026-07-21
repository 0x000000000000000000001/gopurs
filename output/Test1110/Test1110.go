package Test1110

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var X = gopurs_runtime.Value{}
var x = gopurs_runtime.Value{}
var test = gopurs_runtime.Value{}
var Main = gopurs_runtime.Apply(gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fmt.Println(x.StrVal)
return gopurs_runtime.Value{}
})
}), gopurs_runtime.Str("Done"))
var cA = gopurs_runtime.Value{}
var c = gopurs_runtime.Value{}
var test2 = gopurs_runtime.Value{}
