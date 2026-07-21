package Data_Argonaut_Decode_Decoders

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var traverse = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeEither)
var fromFoldable = gopurs_runtime.Apply(foldrArray, Cons)
var traverse1 = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeEither)
var traverse2 = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeEither)
var traverse3 = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeEither)
var traverse4 = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeEither)
var traverse5 = gopurs_runtime.Value{}
var traverseWithIndex = gopurs_runtime.Apply(gopurs_runtime.Value{}, applicativeEither)
var getFieldOptional' = gopurs_runtime.Value{}
var getFieldOptional = gopurs_runtime.Value{}
var getField = gopurs_runtime.Value{}
var decodeVoid = gopurs_runtime.Value{}
var decodeString = gopurs_runtime.Apply(caseJsonString, gopurs_runtime.Value{})
var decodeNumber = gopurs_runtime.Apply(caseJsonNumber, gopurs_runtime.Value{})
var decodeNull = gopurs_runtime.Apply(caseJsonNull, gopurs_runtime.Value{})
var decodeNonEmptyString = gopurs_runtime.Value{}
var decodeMaybe = gopurs_runtime.Value{}
var decodeJObject = gopurs_runtime.Value{}
var decodeJArray = gopurs_runtime.Value{}
var decodeList = gopurs_runtime.Value{}
var decodeSet = gopurs_runtime.Value{}
var decodeNonEmptyArray = gopurs_runtime.Value{}
var decodeNonEmptyList = gopurs_runtime.Value{}
var decodeNonEmpty_Array = gopurs_runtime.Value{}
var decodeNonEmpty_List = gopurs_runtime.Value{}
var decodeInt = gopurs_runtime.Value{}
var decodeIdentity = gopurs_runtime.Value{}
var decodeForeignObject = gopurs_runtime.Value{}
var decodeEither = gopurs_runtime.Value{}
var decodeCodePoint = gopurs_runtime.Value{}
var decodeBoolean = gopurs_runtime.Apply(caseJsonBoolean, gopurs_runtime.Value{})
var decodeArray = gopurs_runtime.Value{}
var decodeTuple = gopurs_runtime.Value{}
var decodeMap = gopurs_runtime.Value{}