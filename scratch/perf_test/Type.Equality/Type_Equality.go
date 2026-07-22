package Type_Equality

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var refl gopurs_runtime.Value
var once_refl sync.Once
func Get_refl() gopurs_runtime.Value {
	once_refl.Do(func() {
		refl = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"proof": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}), "Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return refl
}

var proof gopurs_runtime.Value
var once_proof sync.Once
func Get_proof() gopurs_runtime.Value {
	once_proof.Do(func() {
		proof = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["proof"]
})
	})
	return proof
}

var to gopurs_runtime.Value
var once_to sync.Once
func Get_to() gopurs_runtime.Value {
	once_to.Do(func() {
		to = gopurs_runtime.Func(func(dictTypeEquals_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictTypeEquals_0.PtrVal.(map[string]gopurs_runtime.Value)["proof"], gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
})
	})
	return to
}

var from gopurs_runtime.Value
var once_from sync.Once
func Get_from() gopurs_runtime.Value {
	once_from.Do(func() {
		from = gopurs_runtime.Func(func(dictTypeEquals_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictTypeEquals_0.PtrVal.(map[string]gopurs_runtime.Value)["proof"], gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
})
	})
	return from
}


