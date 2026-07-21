package Ansi_Output

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var withGraphics = gopurs_runtime.Value{}
var underline = gopurs_runtime.Value{}
var strikethrough = gopurs_runtime.Value{}
var italic = gopurs_runtime.Value{}
var inverse = gopurs_runtime.Value{}
var foreground = gopurs_runtime.Value{}
var dim = gopurs_runtime.Value{}
var bold = gopurs_runtime.Value{}
var background = gopurs_runtime.Value{}