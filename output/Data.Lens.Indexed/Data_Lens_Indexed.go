package Data_Lens_Indexed

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var applicativeStateT = gopurs_runtime.Apply(applicativeStateT, monadIdentity)
var applicativeCompose = gopurs_runtime.Apply(applicativeCompose, applicativeStateT)
var applyStateT = gopurs_runtime.Apply(applyStateT, monadIdentity)
var monadStateStateT = gopurs_runtime.Apply(monadStateStateT, monadIdentity)
var get = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Value{})
var unIndex = gopurs_runtime.Value{}
var reindexed = gopurs_runtime.Value{}
var iwander = gopurs_runtime.Value{}
var positions = gopurs_runtime.Value{}
var itraversed = gopurs_runtime.Value{}
var asIndex = gopurs_runtime.Value{}
