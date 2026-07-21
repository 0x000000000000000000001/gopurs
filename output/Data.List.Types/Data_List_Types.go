package Data_List_Types

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var Nil = gopurs_runtime.Value{}
var Cons = gopurs_runtime.Value{}
var NonEmptyList = gopurs_runtime.Value{}
var toList = gopurs_runtime.Value{}
var newtypeNonEmptyList = gopurs_runtime.Value{}
var nelCons = gopurs_runtime.Value{}
var listMap = gopurs_runtime.Value{}
var functorList = gopurs_runtime.Value{}
var functorNonEmptyList = gopurs_runtime.Value{}

var foldableList = gopurs_runtime.Value{}

var intercalate = gopurs_runtime.Value{}
var foldableNonEmptyList = gopurs_runtime.Value{}

var foldableWithIndexList = gopurs_runtime.Value{}

var foldableWithIndexNonEmpty = gopurs_runtime.Apply(foldableWithIndexNonEmpty, foldableWithIndexList)
var foldableWithIndexNonEmptyList = gopurs_runtime.Value{}
var functorWithIndexList = gopurs_runtime.Value{}
var mapWithIndex = gopurs_runtime.Value{}
var functorWithIndexNonEmptyList = gopurs_runtime.Value{}
var semigroupList = gopurs_runtime.Value{}
var monoidList = gopurs_runtime.Value{}
var semigroupNonEmptyList = gopurs_runtime.Value{}
var showList = gopurs_runtime.Value{}
var showNonEmptyList = gopurs_runtime.Value{}

var traversableList = gopurs_runtime.Value{}

var traversableNonEmptyList = gopurs_runtime.Apply(traversableNonEmpty, traversableList)
var traversableWithIndexList = gopurs_runtime.Value{}
var traverseWithIndex = gopurs_runtime.Value{}
var traversableWithIndexNonEmptyList = gopurs_runtime.Value{}
var unfoldable1List = gopurs_runtime.Value{}
var unfoldableList = gopurs_runtime.Value{}
var unfoldable1NonEmptyList = gopurs_runtime.Value{}
var foldable1NonEmptyList = gopurs_runtime.Apply(foldable1NonEmpty, foldableList)
var extendNonEmptyList = gopurs_runtime.Value{}
var extendList = gopurs_runtime.Value{}
var eq1List = gopurs_runtime.Value{}
var eq1NonEmptyList = gopurs_runtime.Value{}
var eqList = gopurs_runtime.Value{}
var eqNonEmptyList = gopurs_runtime.Value{}
var ord1List = gopurs_runtime.Value{}
var ordNonEmpty = gopurs_runtime.Apply(ordNonEmpty, ord1List)
var ord1NonEmptyList = gopurs_runtime.Apply(ord1NonEmpty, ord1List)
var ordList = gopurs_runtime.Value{}
var ordNonEmptyList = gopurs_runtime.Value{}
var comonadNonEmptyList = gopurs_runtime.Value{}

var applyList = gopurs_runtime.Value{}

var applyNonEmptyList = gopurs_runtime.Value{}

var bindList = gopurs_runtime.Value{}

var bindNonEmptyList = gopurs_runtime.Value{}
var applicativeList = gopurs_runtime.Value{}
var monadList = gopurs_runtime.Value{}
var altNonEmptyList = gopurs_runtime.Value{}
var altList = gopurs_runtime.Value{}
var plusList = gopurs_runtime.Value{}
var alternativeList = gopurs_runtime.Value{}
var monadPlusList = gopurs_runtime.Value{}
var applicativeNonEmptyList = gopurs_runtime.Value{}
var monadNonEmptyList = gopurs_runtime.Value{}

var traversable1NonEmptyList = gopurs_runtime.Value{}