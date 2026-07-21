package Effect

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var monadEffect = gopurs_runtime.Value{}
var bindEffect = gopurs_runtime.Value{}
var applyEffect = gopurs_runtime.Value{}
var applicativeEffect = gopurs_runtime.Value{}
var functorEffect = gopurs_runtime.Value{}

var lift2 = gopurs_runtime.Value{}
var semigroupEffect = gopurs_runtime.Value{}
var monoidEffect = gopurs_runtime.Value{}