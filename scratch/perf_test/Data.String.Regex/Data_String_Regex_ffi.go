package Data_String_Regex
import "gopurs/output/gopurs_runtime"
func dummyFn(name string) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		panic("Not implemented: " + name)
	})
}
var ShowRegexImpl = dummyFn("ShowRegexImpl")
var RegexImpl = dummyFn("RegexImpl")
var Source = dummyFn("Source")
var FlagsImpl = dummyFn("FlagsImpl")
var Test = dummyFn("Test")
var X_Match = dummyFn("X_Match")
var Replace = dummyFn("Replace")
var X_ReplaceBy = dummyFn("X_ReplaceBy")
var X_Search = dummyFn("X_Search")
var Split = dummyFn("Split")
