package Main

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var x = gopurs_runtime.Value{}
var Main = gopurs_runtime.Apply(gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
resA := gopurs_runtime.Apply(a, gopurs_runtime.Value{})
resB := gopurs_runtime.Apply(f, resA)
return gopurs_runtime.Apply(resB, gopurs_runtime.Value{})
})
})
}), gopurs_runtime.Apply(gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fmt.Println(x.StrVal)
return gopurs_runtime.Value{}
})
}), gopurs_runtime.Apply(gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(fmt.Sprintf("%q", s.StrVal))
}), gopurs_runtime.Str("Test"))))
