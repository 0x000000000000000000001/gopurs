package Control_MonadPlus

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Monad "gopurs/output/Control.Monad"
)

var monadPlusArray gopurs_runtime.Value
var once_monadPlusArray sync.Once
func Get_monadPlusArray() gopurs_runtime.Value {
	once_monadPlusArray.Do(func() {
		monadPlusArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadArray()
}), "Alternative1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alternative.Get_alternativeArray()
})})
	})
	return monadPlusArray
}


