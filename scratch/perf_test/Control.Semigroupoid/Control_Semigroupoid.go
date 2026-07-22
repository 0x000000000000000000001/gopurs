package Control_Semigroupoid

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var semigroupoidFn gopurs_runtime.Value
var once_semigroupoidFn sync.Once
func Get_semigroupoidFn() gopurs_runtime.Value {
	once_semigroupoidFn.Do(func() {
		semigroupoidFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
})})
	})
	return semigroupoidFn
}

var compose gopurs_runtime.Value
var once_compose sync.Once
func Get_compose() gopurs_runtime.Value {
	once_compose.Do(func() {
		compose = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"]
})
	})
	return compose
}

var composeFlipped gopurs_runtime.Value
var once_composeFlipped sync.Once
func Get_composeFlipped() gopurs_runtime.Value {
	once_composeFlipped.Do(func() {
		composeFlipped = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupoid_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], g_2), f_1)
})
})
})
	})
	return composeFlipped
}


