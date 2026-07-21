package Node_Buffer_ST

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var writeString = gopurs_runtime.Apply(unsafeCoerce, writeString)
var write = gopurs_runtime.Apply(unsafeCoerce, write)
var unsafeThaw = gopurs_runtime.Apply(unsafeCoerce, unsafeThaw)
var unsafeFreeze = gopurs_runtime.Apply(unsafeCoerce, unsafeFreeze)
var transcode = gopurs_runtime.Apply(unsafeCoerce, transcode)
var toString' = gopurs_runtime.Apply(unsafeCoerce, toString')
var toString = gopurs_runtime.Apply(unsafeCoerce, toString)
var toArrayBuffer = gopurs_runtime.Apply(unsafeCoerce, gopurs_runtime.Apply(usingFromImmutable, toArrayBuffer))
var toArray = gopurs_runtime.Apply(unsafeCoerce, gopurs_runtime.Apply(usingFromImmutable, toArray))
var thaw = gopurs_runtime.Apply(unsafeCoerce, thaw)
var swap64 = gopurs_runtime.Apply(unsafeCoerce, swap64)
var swap32 = gopurs_runtime.Apply(unsafeCoerce, swap32)
var swap16 = gopurs_runtime.Apply(unsafeCoerce, swap16)
var slice = gopurs_runtime.Apply(unsafeCoerce, slice)
var size = gopurs_runtime.Apply(unsafeCoerce, gopurs_runtime.Apply(usingFromImmutable, size))
var setAtOffset = gopurs_runtime.Apply(unsafeCoerce, setAtOffset)
var run = gopurs_runtime.Value{}
var readString = gopurs_runtime.Apply(unsafeCoerce, readString)
var read = gopurs_runtime.Apply(unsafeCoerce, read)
var getAtOffset = gopurs_runtime.Apply(unsafeCoerce, getAtOffset)
var fromString = gopurs_runtime.Apply(unsafeCoerce, fromString)
var fromArrayBuffer = gopurs_runtime.Apply(unsafeCoerce, gopurs_runtime.Apply(usingToImmutable, fromArrayBuffer))
var fromArray = gopurs_runtime.Apply(unsafeCoerce, gopurs_runtime.Apply(usingToImmutable, fromArray))
var freeze = gopurs_runtime.Apply(unsafeCoerce, freeze)
var fill = gopurs_runtime.Apply(unsafeCoerce, fill)
var copy = gopurs_runtime.Apply(unsafeCoerce, copy)
var concat' = gopurs_runtime.Apply(unsafeCoerce, concat')
var concat = gopurs_runtime.Apply(unsafeCoerce, concat)
var compareParts = gopurs_runtime.Apply(unsafeCoerce, compareParts)
var allocUnsafeSlow = gopurs_runtime.Apply(unsafeCoerce, allocUnsafeSlow)
var allocUnsafe = gopurs_runtime.Apply(unsafeCoerce, allocUnsafe)
var alloc = gopurs_runtime.Apply(unsafeCoerce, gopurs_runtime.Apply(usingToImmutable, alloc))
var create = alloc
var mutableBufferST = gopurs_runtime.Value{}