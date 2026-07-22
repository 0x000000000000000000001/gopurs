package Either

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Left gopurs_runtime.Value
var once_Left sync.Once
func Get_Left() gopurs_runtime.Value {
	once_Left.Do(func() {
		Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": value0})
})
	})
	return Left
}

var Right gopurs_runtime.Value
var once_Right sync.Once
func Get_Right() gopurs_runtime.Value {
	once_Right.Do(func() {
		Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": value0})
})
	})
	return Right
}


