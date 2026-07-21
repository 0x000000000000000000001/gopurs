package Gopurs_CodeGen

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var fromFoldable = gopurs_runtime.Apply(runFn2, fromFoldableImpl)

var translateExpr = gopurs_runtime.Value{}

var translateBinding = gopurs_runtime.Value{}
var translateBindingGroup = gopurs_runtime.Value{}
var translate = gopurs_runtime.Value{}