package Node_Buffer

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var writeString = gopurs_runtime.Value{}
var write = gopurs_runtime.Value{}
var unsafeThaw = gopurs_runtime.Value{}
var usingToImmutable = gopurs_runtime.Value{}
var unsafeFreeze = gopurs_runtime.Value{}
var usingFromImmutable = gopurs_runtime.Value{}
var transcode = gopurs_runtime.Value{}
var toString' = gopurs_runtime.Value{}
var toString = gopurs_runtime.Value{}
var toArrayBuffer = gopurs_runtime.Apply(usingFromImmutable, toArrayBuffer)
var toArray = gopurs_runtime.Apply(usingFromImmutable, toArray)
var thaw = gopurs_runtime.Apply(runEffectFn1, thawImpl)
var swap64 = gopurs_runtime.Value{}
var swap32 = gopurs_runtime.Value{}
var swap16 = gopurs_runtime.Value{}
var slice = gopurs_runtime.Apply(unsafeCoerce, slice)
var size = gopurs_runtime.Apply(usingFromImmutable, size)
var setPoolSize = gopurs_runtime.Value{}
var setAtOffset = gopurs_runtime.Value{}
var readString = gopurs_runtime.Value{}
var read = gopurs_runtime.Value{}
var getAtOffset = gopurs_runtime.Value{}
var fromString = gopurs_runtime.Value{}
var fromArrayBuffer = gopurs_runtime.Apply(usingToImmutable, fromArrayBuffer)
var fromArray = gopurs_runtime.Apply(usingToImmutable, fromArray)
var freeze = gopurs_runtime.Apply(runEffectFn1, freezeImpl)
var fill = gopurs_runtime.Value{}
var copy = gopurs_runtime.Value{}
var concat' = gopurs_runtime.Value{}
var concat = gopurs_runtime.Value{}
var compareParts = gopurs_runtime.Value{}
var allocUnsafeSlow = gopurs_runtime.Value{}
var allocUnsafe = gopurs_runtime.Value{}
var alloc = gopurs_runtime.Apply(usingToImmutable, alloc)
var create = gopurs_runtime.Apply(usingToImmutable, alloc)
var mutableBufferEffect = gopurs_runtime.Value{}