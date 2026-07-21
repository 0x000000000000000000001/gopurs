package Data_Interval_Duration_Iso

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var lookup = gopurs_runtime.Value{}
var foldMap1 = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidList)
var foldMap2 = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var fold = gopurs_runtime.Value{}
var toUnfoldable = gopurs_runtime.Value{}
var IsEmpty = gopurs_runtime.Value{}
var InvalidWeekComponentUsage = gopurs_runtime.Value{}
var ContainsNegativeValue = gopurs_runtime.Value{}
var InvalidFractionalUse = gopurs_runtime.Value{}
var unIsoDuration = gopurs_runtime.Value{}
var showIsoDuration = gopurs_runtime.Value{}
var showError = gopurs_runtime.Value{}
var prettyError = gopurs_runtime.Value{}
var eqIsoDuration = gopurs_runtime.Value{}
var ordIsoDuration = gopurs_runtime.Value{}
var eqError = gopurs_runtime.Value{}
var ordError = gopurs_runtime.Value{}
var checkWeekUsage = gopurs_runtime.Value{}
var checkNegativeValues = gopurs_runtime.Value{}
var checkFractionalUse = gopurs_runtime.Value{}
var checkEmptiness = gopurs_runtime.Value{}
var checkValidIsoDuration = gopurs_runtime.Value{}
var mkIsoDuration = gopurs_runtime.Value{}