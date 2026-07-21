package Data_Posix

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var Uid = gopurs_runtime.Value{}
var Pid = gopurs_runtime.Value{}
var Gid = gopurs_runtime.Value{}
var showUid = gopurs_runtime.Value{}
var showPid = gopurs_runtime.Value{}
var showGid = gopurs_runtime.Value{}
var ordUid = ordInt
var ordPid = ordInt
var ordGid = ordInt
var newtypeUid = gopurs_runtime.Value{}
var newtypePid = gopurs_runtime.Value{}
var newtypeGid = gopurs_runtime.Value{}
var eqUid = eqInt
var eqPid = eqInt
var eqGid = eqInt