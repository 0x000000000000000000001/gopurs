package Node_FS

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var FileLink = gopurs_runtime.Value{}
var DirLink = gopurs_runtime.Value{}
var JunctionLink = gopurs_runtime.Value{}
var symlinkTypeToNode = gopurs_runtime.Value{}
var showSymlinkType = gopurs_runtime.Value{}
var eqSymlinkType = gopurs_runtime.Value{}