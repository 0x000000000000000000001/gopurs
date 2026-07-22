package Foreign_Object_ST

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		peek = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_peekImpl(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return peek
}

func Get_delete_() gopurs_runtime.Value {
	return Delete_
}

func Get_new_() gopurs_runtime.Value {
	return New_
}

func Get_peekImpl() gopurs_runtime.Value {
	return PeekImpl
}

func Get_poke() gopurs_runtime.Value {
	return Poke
}
