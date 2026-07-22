package B

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var X gopurs_runtime.Value
var once_X sync.Once
func Get_X() gopurs_runtime.Value {
	once_X.Do(func() {
		X = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("X")})
	})
	return X
}

var Y gopurs_runtime.Value
var once_Y sync.Once
func Get_Y() gopurs_runtime.Value {
	once_Y.Do(func() {
		Y = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Y")})
	})
	return Y
}


