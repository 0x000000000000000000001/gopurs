package Data_Argonaut_Core

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var jsonZero = gopurs_runtime.Apply(fromNumber, gopurs_runtime.Value{})
var jsonTrue = gopurs_runtime.Apply(fromBoolean, gopurs_runtime.Value{})
var jsonSingletonObject = gopurs_runtime.Value{}
var jsonSingletonArray = gopurs_runtime.Value{}
var jsonFalse = gopurs_runtime.Apply(fromBoolean, gopurs_runtime.Value{})
var jsonEmptyString = gopurs_runtime.Apply(fromString, gopurs_runtime.Str(""))
var jsonEmptyObject = gopurs_runtime.Apply(fromObject, empty)
var jsonEmptyArray = gopurs_runtime.Apply(fromArray, gopurs_runtime.Value{})

var ordJson = gopurs_runtime.Value{}
var eqJson = gopurs_runtime.Value{}

var caseJsonString = gopurs_runtime.Value{}
var isString = gopurs_runtime.Apply(caseJsonString, gopurs_runtime.Value{})
var toString = gopurs_runtime.Apply(caseJsonString, gopurs_runtime.Value{})
var caseJsonObject = gopurs_runtime.Value{}
var isObject = gopurs_runtime.Apply(caseJsonObject, gopurs_runtime.Value{})
var toObject = gopurs_runtime.Apply(caseJsonObject, gopurs_runtime.Value{})
var caseJsonNumber = gopurs_runtime.Value{}
var isNumber = gopurs_runtime.Apply(caseJsonNumber, gopurs_runtime.Value{})
var toNumber = gopurs_runtime.Apply(caseJsonNumber, gopurs_runtime.Value{})
var caseJsonNull = gopurs_runtime.Value{}
var isNull = gopurs_runtime.Apply(caseJsonNull, gopurs_runtime.Value{})
var toNull = gopurs_runtime.Apply(caseJsonNull, gopurs_runtime.Value{})
var caseJsonBoolean = gopurs_runtime.Value{}
var isBoolean = gopurs_runtime.Apply(caseJsonBoolean, gopurs_runtime.Value{})
var toBoolean = gopurs_runtime.Apply(caseJsonBoolean, gopurs_runtime.Value{})
var caseJsonArray = gopurs_runtime.Value{}
var isArray = gopurs_runtime.Apply(caseJsonArray, gopurs_runtime.Value{})
var toArray = gopurs_runtime.Apply(caseJsonArray, gopurs_runtime.Value{})
var caseJson = gopurs_runtime.Value{}