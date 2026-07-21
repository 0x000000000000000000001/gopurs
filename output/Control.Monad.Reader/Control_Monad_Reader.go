package Control_Monad_Reader

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var withReader = withReaderT
var runReader = gopurs_runtime.Value{}
var mapReader = gopurs_runtime.Value{}