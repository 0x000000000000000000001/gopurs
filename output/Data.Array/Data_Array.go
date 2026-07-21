package Data_Array

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var traverse_ = gopurs_runtime.Apply(traverse_, applicativeST)
var void = gopurs_runtime.Apply(map_, gopurs_runtime.Value{})
var intercalate1 = gopurs_runtime.Value{}
var zipWith = gopurs_runtime.Apply(runFn3, zipWithImpl)
var zipWithA = gopurs_runtime.Value{}
var zip = gopurs_runtime.Apply(zipWith, Tuple)
var updateAtIndices = gopurs_runtime.Value{}
var updateAt = gopurs_runtime.Apply(runFn5, _updateAt)
var unsafeIndex = gopurs_runtime.Value{}
var unsafeIndex1 = gopurs_runtime.Apply(runFn2, unsafeIndexImpl)
var uncons = gopurs_runtime.Apply(runFn3, unconsImpl)
var toUnfoldable = gopurs_runtime.Value{}
var tail = gopurs_runtime.Apply(runFn3, unconsImpl)
var sortBy = gopurs_runtime.Value{}
var sortWith = gopurs_runtime.Value{}
var sort = gopurs_runtime.Value{}
var snoc = gopurs_runtime.Value{}
var slice = gopurs_runtime.Apply(runFn3, sliceImpl)
var splitAt = gopurs_runtime.Value{}
var take = gopurs_runtime.Value{}
var singleton = gopurs_runtime.Value{}
var scanr = gopurs_runtime.Apply(runFn3, scanrImpl)
var scanl = gopurs_runtime.Apply(runFn3, scanlImpl)
var replicate = gopurs_runtime.Apply(runFn2, replicateImpl)
var range = gopurs_runtime.Apply(runFn2, rangeImpl)
var partition = gopurs_runtime.Apply(runFn2, partitionImpl)
var null = gopurs_runtime.Value{}
var modifyAtIndices = gopurs_runtime.Value{}
var mapWithIndex = mapWithIndexArray
var intersperse = gopurs_runtime.Value{}
var intercalate = gopurs_runtime.Value{}
var insertAt = gopurs_runtime.Apply(runFn5, _insertAt)
var init = gopurs_runtime.Value{}
var index = gopurs_runtime.Apply(runFn4, indexImpl)
var last = gopurs_runtime.Value{}
var unsnoc = gopurs_runtime.Value{}
var modifyAt = gopurs_runtime.Value{}
var span = gopurs_runtime.Value{}
var takeWhile = gopurs_runtime.Value{}
var unzip = gopurs_runtime.Value{}
var head = gopurs_runtime.Value{}
var nubBy = gopurs_runtime.Value{}
var nub = gopurs_runtime.Value{}
var groupBy = gopurs_runtime.Value{}
var groupAllBy = gopurs_runtime.Value{}
var groupAll = gopurs_runtime.Value{}
var group = gopurs_runtime.Value{}
var fromFoldable = gopurs_runtime.Value{}
var foldr = foldrArray
var foldl = foldlArray
var transpose = gopurs_runtime.Value{}
var foldRecM = gopurs_runtime.Value{}
var foldMap = gopurs_runtime.Value{}

var foldM = gopurs_runtime.Value{}

var fold = gopurs_runtime.Value{}
var findMap = gopurs_runtime.Apply(runFn4, findMapImpl)
var findLastIndex = gopurs_runtime.Apply(runFn4, findLastIndexImpl)
var insertBy = gopurs_runtime.Value{}
var insert = gopurs_runtime.Value{}
var findIndex = gopurs_runtime.Apply(runFn4, findIndexImpl)
var find = gopurs_runtime.Value{}
var filter = gopurs_runtime.Apply(runFn2, filterImpl)
var intersectBy = gopurs_runtime.Value{}
var intersect = gopurs_runtime.Value{}
var elemLastIndex = gopurs_runtime.Value{}
var elemIndex = gopurs_runtime.Value{}
var notElem = gopurs_runtime.Value{}
var elem = gopurs_runtime.Value{}
var dropWhile = gopurs_runtime.Value{}
var dropEnd = gopurs_runtime.Value{}
var drop = gopurs_runtime.Value{}
var takeEnd = gopurs_runtime.Value{}
var deleteAt = gopurs_runtime.Apply(runFn4, _deleteAt)
var deleteBy = gopurs_runtime.Value{}
var delete = gopurs_runtime.Value{}
var difference = gopurs_runtime.Value{}
var cons = gopurs_runtime.Value{}

var some = gopurs_runtime.Value{}
var many = gopurs_runtime.Value{}

var concatMap = gopurs_runtime.Value{}
var mapMaybe = gopurs_runtime.Value{}
var filterA = gopurs_runtime.Value{}
var catMaybes = gopurs_runtime.Apply(mapMaybe, gopurs_runtime.Value{})
var any = gopurs_runtime.Apply(runFn2, anyImpl)
var nubByEq = gopurs_runtime.Value{}
var nubEq = gopurs_runtime.Value{}
var unionBy = gopurs_runtime.Value{}
var union = gopurs_runtime.Value{}
var alterAt = gopurs_runtime.Value{}
var all = gopurs_runtime.Apply(runFn2, allImpl)