package Data_String_NonEmpty_Internal

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var NonEmptyString = gopurs_runtime.Value{}
var NonEmptyReplacement = gopurs_runtime.Value{}
var toUpper = gopurs_runtime.Value{}
var toString = gopurs_runtime.Value{}
var toLower = gopurs_runtime.Value{}
var showNonEmptyString = gopurs_runtime.Value{}
var showNonEmptyReplacement = gopurs_runtime.Value{}
var semigroupNonEmptyString = semigroupString
var semigroupNonEmptyReplacement = semigroupString
var replaceAll = gopurs_runtime.Value{}
var replace = gopurs_runtime.Value{}
var prependString = gopurs_runtime.Value{}
var ordNonEmptyString = ordString
var ordNonEmptyReplacement = ordString
var nonEmptyNonEmpty = gopurs_runtime.Value{}
var nes = gopurs_runtime.Value{}
var makeNonEmptyBad = gopurs_runtime.Value{}
var localeCompare = gopurs_runtime.Value{}
var liftS = gopurs_runtime.Value{}
var joinWith1 = gopurs_runtime.Value{}
var joinWith = gopurs_runtime.Value{}
var join1With = gopurs_runtime.Value{}
var fromString = gopurs_runtime.Value{}
var stripPrefix = gopurs_runtime.Value{}
var stripSuffix = gopurs_runtime.Value{}
var trim = gopurs_runtime.Value{}
var unsafeFromString = gopurs_runtime.Value{}
var eqNonEmptyString = eqString
var eqNonEmptyReplacement = eqString
var contains = gopurs_runtime.Value{}
var appendString = gopurs_runtime.Value{}