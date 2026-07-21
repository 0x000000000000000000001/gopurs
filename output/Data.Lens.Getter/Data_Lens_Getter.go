package Data_Lens_Getter

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var fanout = gopurs_runtime.Apply(fanout, semigroupoidFn)
var view = gopurs_runtime.Value{}
var viewOn = gopurs_runtime.Value{}
var use = gopurs_runtime.Value{}
var to = gopurs_runtime.Value{}
var takeBoth = gopurs_runtime.Value{}
var iview = gopurs_runtime.Value{}
var iuse = gopurs_runtime.Value{}
var cloneGetter = gopurs_runtime.Value{}
