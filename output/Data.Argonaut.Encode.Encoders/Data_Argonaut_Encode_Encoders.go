package Data_Argonaut_Encode_Encoders

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var toUnfoldable = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var toUnfoldable1 = gopurs_runtime.Value{}
var toUnfoldable2 = gopurs_runtime.Apply(toUnfoldable, unfoldableList)
var fromFoldable = gopurs_runtime.Apply(fromFoldable, foldableList)
var extend = gopurs_runtime.Value{}
var extendOptional = gopurs_runtime.Value{}
var encodeVoid = absurd
var encodeUnit = gopurs_runtime.Value{}
var encodeTuple = gopurs_runtime.Value{}
var encodeString = fromString
var encodeNumber = fromNumber
var encodeNonEmptyString = gopurs_runtime.Value{}
var encodeMaybe = gopurs_runtime.Value{}
var encodeList = gopurs_runtime.Value{}
var encodeMap = gopurs_runtime.Value{}
var encodeNonEmptyList = gopurs_runtime.Value{}
var encodeNonEmpty_List = gopurs_runtime.Value{}
var encodeSet = gopurs_runtime.Value{}
var encodeInt = gopurs_runtime.Value{}
var encodeIdentity = gopurs_runtime.Value{}
var encodeForeignObject = gopurs_runtime.Value{}
var encodeEither = gopurs_runtime.Value{}
var encodeCodePoint = gopurs_runtime.Value{}
var encodeChar = gopurs_runtime.Value{}
var encodeBoolean = fromBoolean
var encodeArray = gopurs_runtime.Value{}
var encodeNonEmptyArray = gopurs_runtime.Value{}
var encodeNonEmpty_Array = gopurs_runtime.Value{}
var assocOptional = gopurs_runtime.Value{}
var assoc = gopurs_runtime.Value{}