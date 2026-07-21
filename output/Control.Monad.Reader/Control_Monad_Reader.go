package Control_Monad_Reader

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var withReader = withReaderT
var runReader = gopurs_runtime.Value{}
var mapReader = gopurs_runtime.Value{}
