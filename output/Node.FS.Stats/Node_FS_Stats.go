package Node_FS_Stats

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var uid = gopurs_runtime.Value{}
var statusChangedTimeMs = gopurs_runtime.Value{}
var statusChangedTime = gopurs_runtime.Value{}
var size = gopurs_runtime.Value{}
var showStats = gopurs_runtime.Value{}
var rdev = gopurs_runtime.Value{}
var nlink = gopurs_runtime.Value{}
var modifiedTimeMs = gopurs_runtime.Value{}
var modifiedTime = gopurs_runtime.Value{}
var mode = gopurs_runtime.Value{}
var isSymbolicLink = gopurs_runtime.Value{}
var isSocket = gopurs_runtime.Value{}
var isFile = gopurs_runtime.Value{}
var isFIFO = gopurs_runtime.Value{}
var isDirectory = gopurs_runtime.Value{}
var isCharacterDevice = gopurs_runtime.Value{}
var isBlockDevice = gopurs_runtime.Value{}
var inode = gopurs_runtime.Value{}
var gid = gopurs_runtime.Value{}
var dev = gopurs_runtime.Value{}
var blocks = gopurs_runtime.Value{}
var blkSize = gopurs_runtime.Value{}
var birthtimeMs = gopurs_runtime.Value{}
var birthTime = gopurs_runtime.Value{}
var accessedTimeMs = gopurs_runtime.Value{}
var accessedTime = gopurs_runtime.Value{}