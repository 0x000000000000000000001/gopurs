package Data_List

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var any = gopurs_runtime.Apply(any, foldableList)
var Pattern = gopurs_runtime.Value{}

var updateAt = gopurs_runtime.Value{}

var unzip = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var uncons = gopurs_runtime.Value{}
var toUnfoldable = gopurs_runtime.Value{}
var tail = gopurs_runtime.Value{}
var stripPrefix = gopurs_runtime.Value{}

var span = gopurs_runtime.Value{}

var snoc = gopurs_runtime.Value{}
var singleton = gopurs_runtime.Value{}
var sortBy = gopurs_runtime.Value{}
var sort = gopurs_runtime.Value{}



var showPattern = gopurs_runtime.Value{}
var reverse = gopurs_runtime.Value{}
var take = gopurs_runtime.Value{}
var takeWhile = gopurs_runtime.Value{}
var unsnoc = gopurs_runtime.Value{}
var zipWith = gopurs_runtime.Value{}
var zip = gopurs_runtime.Value{}
var zipWithA = gopurs_runtime.Value{}
var range = gopurs_runtime.Value{}
var partition = gopurs_runtime.Value{}
var null = gopurs_runtime.Value{}
var nubBy = gopurs_runtime.Value{}
var nub = gopurs_runtime.Value{}
var newtypePattern = gopurs_runtime.Value{}
var mapMaybe = gopurs_runtime.Value{}
var manyRec = gopurs_runtime.Value{}
var someRec = gopurs_runtime.Value{}

var some = gopurs_runtime.Value{}
var many = gopurs_runtime.Value{}

var length = gopurs_runtime.Value{}

var last = gopurs_runtime.Value{}

var insertBy = gopurs_runtime.Value{}

var insertAt = gopurs_runtime.Value{}

var insert = gopurs_runtime.Value{}
var init = gopurs_runtime.Value{}

var index = gopurs_runtime.Value{}

var head = gopurs_runtime.Value{}

var transpose = gopurs_runtime.Value{}

var groupBy = gopurs_runtime.Value{}

var groupAllBy = gopurs_runtime.Value{}
var group = gopurs_runtime.Value{}
var groupAll = gopurs_runtime.Value{}
var fromFoldable = gopurs_runtime.Value{}

var foldM = gopurs_runtime.Value{}

var findIndex = gopurs_runtime.Value{}
var findLastIndex = gopurs_runtime.Value{}

var filterM = gopurs_runtime.Value{}

var filter = gopurs_runtime.Value{}
var intersectBy = gopurs_runtime.Value{}
var intersect = gopurs_runtime.Value{}

var nubByEq = gopurs_runtime.Value{}

var nubEq = gopurs_runtime.Value{}
var eqPattern = gopurs_runtime.Value{}
var ordPattern = gopurs_runtime.Value{}
var elemLastIndex = gopurs_runtime.Value{}
var elemIndex = gopurs_runtime.Value{}
var dropWhile = gopurs_runtime.Value{}
var dropEnd = gopurs_runtime.Value{}

var drop = gopurs_runtime.Value{}

var slice = gopurs_runtime.Value{}
var takeEnd = gopurs_runtime.Value{}

var deleteBy = gopurs_runtime.Value{}

var unionBy = gopurs_runtime.Value{}
var union = gopurs_runtime.Value{}

var deleteAt = gopurs_runtime.Value{}

var delete = gopurs_runtime.Value{}
var difference = gopurs_runtime.Value{}
var concatMap = gopurs_runtime.Value{}
var concat = gopurs_runtime.Value{}
var catMaybes = gopurs_runtime.Value{}

var alterAt = gopurs_runtime.Value{}

var modifyAt = gopurs_runtime.Value{}