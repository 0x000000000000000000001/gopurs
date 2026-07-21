package Data_Array_ST

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var unshiftAll = gopurs_runtime.Apply(runSTFn2, unshiftAllImpl)
var unshift = gopurs_runtime.Value{}
var unsafeThaw = gopurs_runtime.Apply(runSTFn1, unsafeThawImpl)
var unsafeFreeze = gopurs_runtime.Apply(runSTFn1, unsafeFreezeImpl)
var toAssocArray = gopurs_runtime.Apply(runSTFn1, toAssocArrayImpl)
var thaw = gopurs_runtime.Apply(runSTFn1, thawImpl)
var withArray = gopurs_runtime.Value{}
var splice = gopurs_runtime.Apply(runSTFn4, spliceImpl)
var sortBy = gopurs_runtime.Value{}
var sortWith = gopurs_runtime.Value{}
var sort = gopurs_runtime.Value{}
var shift = gopurs_runtime.Apply(runSTFn3, shiftImpl)
var run = gopurs_runtime.Value{}
var pushAll = gopurs_runtime.Apply(runSTFn2, pushAllImpl)
var push = gopurs_runtime.Apply(runSTFn2, pushImpl)
var pop = gopurs_runtime.Apply(runSTFn3, popImpl)
var poke = gopurs_runtime.Apply(runSTFn3, pokeImpl)
var peek = gopurs_runtime.Apply(runSTFn4, peekImpl)
var modify = gopurs_runtime.Value{}
var length = gopurs_runtime.Apply(runSTFn1, lengthImpl)
var freeze = gopurs_runtime.Apply(runSTFn1, freezeImpl)
var clone = gopurs_runtime.Apply(runSTFn1, cloneImpl)