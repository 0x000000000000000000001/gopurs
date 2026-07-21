package Node_FS_Perms

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var eq = gopurs_runtime.Value{}
var compare = gopurs_runtime.Value{}
var write = gopurs_runtime.Value{}
var semiringPerm = gopurs_runtime.Value{}
var read = gopurs_runtime.Value{}
var permToInt = gopurs_runtime.Value{}
var permsToString = gopurs_runtime.Value{}
var permsToInt = gopurs_runtime.Apply(_unsafePartial, gopurs_runtime.Value{})
var none = gopurs_runtime.Value{}
var mkPerms = gopurs_runtime.Value{}
var mkPerm = gopurs_runtime.Value{}
var execute = gopurs_runtime.Value{}
var permFromChar = gopurs_runtime.Value{}
var permsFromString = gopurs_runtime.Value{}
var eqPerm = gopurs_runtime.Value{}
var eqPerms = gopurs_runtime.Value{}
var ordPerm = gopurs_runtime.Value{}
var compare1 = gopurs_runtime.Value{}
var ordPerms = gopurs_runtime.Value{}
var all = gopurs_runtime.Value{}
var permsAll = gopurs_runtime.Value{}
var permsReadWrite = gopurs_runtime.Value{}
var showPerm = gopurs_runtime.Value{}
var showPerms = gopurs_runtime.Value{}