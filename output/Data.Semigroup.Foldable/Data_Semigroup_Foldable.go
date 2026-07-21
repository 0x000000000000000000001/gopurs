package Data_Semigroup_Foldable

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var FoldRight1 = gopurs_runtime.Value{}
var mkFoldRight1 = gopurs_runtime.Apply(FoldRight1, const)
var foldr1 = gopurs_runtime.Value{}
var foldl1 = gopurs_runtime.Value{}
var maximumBy = gopurs_runtime.Value{}
var minimumBy = gopurs_runtime.Value{}
var foldableTuple = gopurs_runtime.Value{}
var foldableMultiplicative = gopurs_runtime.Value{}
var foldableIdentity = gopurs_runtime.Value{}
var foldableDual = gopurs_runtime.Value{}
var foldRight1Semigroup = gopurs_runtime.Value{}
var semigroupDual = gopurs_runtime.Value{}
var foldMap1DefaultR = gopurs_runtime.Value{}
var foldMap1DefaultL = gopurs_runtime.Value{}
var foldMap1 = gopurs_runtime.Value{}
var foldl1Default = gopurs_runtime.Value{}
var foldr1Default = gopurs_runtime.Value{}
var intercalateMap = gopurs_runtime.Value{}
var intercalate = gopurs_runtime.Value{}
var maximum = gopurs_runtime.Value{}
var minimum = gopurs_runtime.Value{}
var traverse1_ = gopurs_runtime.Value{}
var for1_ = gopurs_runtime.Value{}
var sequence1_ = gopurs_runtime.Value{}
var fold1 = gopurs_runtime.Value{}