package Control_Monad_Gen_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var sized gopurs_runtime.Value
var once_sized sync.Once
func Get_sized() gopurs_runtime.Value {
	once_sized.Do(func() {
		sized = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["sized"]
})
	})
	return sized
}

var resize gopurs_runtime.Value
var once_resize sync.Once
func Get_resize() gopurs_runtime.Value {
	once_resize.Do(func() {
		resize = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["resize"]
})
	})
	return resize
}

var chooseInt gopurs_runtime.Value
var once_chooseInt sync.Once
func Get_chooseInt() gopurs_runtime.Value {
	once_chooseInt.Do(func() {
		chooseInt = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"]
})
	})
	return chooseInt
}

var chooseFloat gopurs_runtime.Value
var once_chooseFloat sync.Once
func Get_chooseFloat() gopurs_runtime.Value {
	once_chooseFloat.Do(func() {
		chooseFloat = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseFloat"]
})
	})
	return chooseFloat
}

var chooseBool gopurs_runtime.Value
var once_chooseBool sync.Once
func Get_chooseBool() gopurs_runtime.Value {
	once_chooseBool.Do(func() {
		chooseBool = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseBool"]
})
	})
	return chooseBool
}


