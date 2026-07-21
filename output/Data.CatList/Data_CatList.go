package Data_CatList

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var CatNil = gopurs_runtime.Value{}
var CatCons = gopurs_runtime.Value{}

var showCatList = gopurs_runtime.Value{}

var null = gopurs_runtime.Value{}
var link = gopurs_runtime.Value{}
var foldr = gopurs_runtime.Value{}
var uncons = gopurs_runtime.Value{}

var foldableCatList = gopurs_runtime.Value{}

var length = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})

var foldMap = gopurs_runtime.Value{}

var empty = gopurs_runtime.Value{}
var append = link
var cons = gopurs_runtime.Value{}

var functorCatList = gopurs_runtime.Value{}

var singleton = gopurs_runtime.Value{}

var traversableCatList = gopurs_runtime.Value{}

var semigroupCatList = gopurs_runtime.Value{}
var monoidCatList = gopurs_runtime.Value{}

var monadCatList = gopurs_runtime.Value{}
var bindCatList = gopurs_runtime.Value{}
var applyCatList = gopurs_runtime.Value{}
var applicativeCatList = gopurs_runtime.Value{}

var fromFoldable = gopurs_runtime.Value{}
var snoc = gopurs_runtime.Value{}
var unfoldable1CatList = gopurs_runtime.Value{}
var unfoldableCatList = gopurs_runtime.Value{}
var altCatList = gopurs_runtime.Value{}
var plusCatList = gopurs_runtime.Value{}
var alternativeCatList = gopurs_runtime.Value{}
var monadPlusCatList = gopurs_runtime.Value{}