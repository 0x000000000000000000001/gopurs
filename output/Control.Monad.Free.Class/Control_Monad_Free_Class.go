package Control_Monad_Free_Class

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var wrapFree = gopurs_runtime.Value{}
var monadFreeWriterT = gopurs_runtime.Value{}
var monadFreeStateT = gopurs_runtime.Value{}
var monadFreeReaderT = gopurs_runtime.Value{}
var monadFreeMaybeT = gopurs_runtime.Value{}
var monadFreeFree = gopurs_runtime.Value{}
var monadFreeExceptT = gopurs_runtime.Value{}