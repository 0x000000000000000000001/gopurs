package Data_Array_ST_Iterator

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var void = gopurs_runtime.Apply(map_, gopurs_runtime.Value{})
var Iterator = gopurs_runtime.Value{}
var peek = gopurs_runtime.Value{}
var next = gopurs_runtime.Value{}
var pushWhile = gopurs_runtime.Value{}
var pushAll = gopurs_runtime.Apply(pushWhile, gopurs_runtime.Value{})
var iterator = gopurs_runtime.Value{}
var iterate = gopurs_runtime.Value{}
var exhausted = gopurs_runtime.Value{}
