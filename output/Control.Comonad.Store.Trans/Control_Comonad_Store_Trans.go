package Control_Comonad_Store_Trans

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var StoreT = gopurs_runtime.Value{}
var runStoreT = gopurs_runtime.Value{}
var newtypeStoreT = gopurs_runtime.Value{}
var functorStoreT = gopurs_runtime.Value{}
var extendStoreT = gopurs_runtime.Value{}
var comonadTransStoreT = gopurs_runtime.Value{}
var comonadStoreT = gopurs_runtime.Value{}