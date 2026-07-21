package Node_FS_Aff

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var toAff = gopurs_runtime.Value{}
var toAff1 = gopurs_runtime.Value{}
var unlink = gopurs_runtime.Apply(toAff1, unlink)
var toAff2 = gopurs_runtime.Value{}
var truncate = gopurs_runtime.Apply(toAff2, truncate)
var writeFile = gopurs_runtime.Apply(toAff2, writeFile)
var toAff3 = gopurs_runtime.Value{}
var utimes = gopurs_runtime.Apply(toAff3, utimes)
var writeTextFile = gopurs_runtime.Apply(toAff3, writeTextFile)
var toAff5 = gopurs_runtime.Value{}
var symlink = gopurs_runtime.Apply(toAff3, symlink)
var stat = gopurs_runtime.Apply(toAff1, stat)
var rmdir' = gopurs_runtime.Apply(toAff2, rmdir')
var rmdir = gopurs_runtime.Apply(toAff1, rmdir)
var rm' = gopurs_runtime.Apply(toAff2, rm')
var rm = gopurs_runtime.Apply(toAff1, rm)
var rename = gopurs_runtime.Apply(toAff2, rename)
var realpath' = gopurs_runtime.Apply(toAff2, realpath')
var realpath = gopurs_runtime.Apply(toAff1, realpath)
var readlink = gopurs_runtime.Apply(toAff1, readlink)
var readdir = gopurs_runtime.Apply(toAff1, readdir)
var readTextFile = gopurs_runtime.Apply(toAff2, readTextFile)
var readFile = gopurs_runtime.Apply(toAff1, readFile)
var mkdtemp' = gopurs_runtime.Apply(toAff2, mkdtemp')
var mkdtemp = gopurs_runtime.Apply(toAff1, mkdtemp)
var mkdir' = gopurs_runtime.Apply(toAff2, mkdir')
var mkdir = gopurs_runtime.Apply(toAff1, mkdir)
var lstat = gopurs_runtime.Apply(toAff1, lstat)
var link = gopurs_runtime.Apply(toAff2, link)
var fdWrite = gopurs_runtime.Apply(toAff5, fdWrite)
var fdRead = gopurs_runtime.Apply(toAff5, fdRead)
var fdOpen = gopurs_runtime.Apply(toAff3, fdOpen)
var fdNext = gopurs_runtime.Apply(toAff2, fdNext)
var fdClose = gopurs_runtime.Apply(toAff1, fdClose)
var fdAppend = gopurs_runtime.Apply(toAff2, fdAppend)
var copyFile' = gopurs_runtime.Apply(toAff3, copyFile')
var copyFile = gopurs_runtime.Apply(toAff2, copyFile)
var chown = gopurs_runtime.Apply(toAff3, chown)
var chmod = gopurs_runtime.Apply(toAff2, chmod)
var appendTextFile = gopurs_runtime.Apply(toAff3, appendTextFile)
var appendFile = gopurs_runtime.Apply(toAff2, appendFile)
var access' = gopurs_runtime.Value{}
var access = gopurs_runtime.Value{}