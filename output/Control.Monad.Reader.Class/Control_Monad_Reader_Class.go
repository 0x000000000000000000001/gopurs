package Control_Monad_Reader_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var monadAskFun = gopurs_runtime.Value{}
var monadReaderFun = gopurs_runtime.Value{}
var local = gopurs_runtime.Value{}
var ask = gopurs_runtime.Value{}
var asks = gopurs_runtime.Value{}