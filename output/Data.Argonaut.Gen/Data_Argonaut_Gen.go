package Data_Argonaut_Gen

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var foldable1NonEmpty = gopurs_runtime.Apply(foldable1NonEmpty, foldableArray)
var min = gopurs_runtime.Value{}

var genJson = gopurs_runtime.Value{}