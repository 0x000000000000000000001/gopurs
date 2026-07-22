package Data_Function_Uncurried

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var runFn1 gopurs_runtime.Value
var once_runFn1 sync.Once
func Get_runFn1() gopurs_runtime.Value {
	once_runFn1.Do(func() {
		runFn1 = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return f_0
})
	})
	return runFn1
}

var mkFn1 gopurs_runtime.Value
var once_mkFn1 sync.Once
func Get_mkFn1() gopurs_runtime.Value {
	once_mkFn1.Do(func() {
		mkFn1 = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return f_0
})
	})
	return mkFn1
}

func Get_mkFn0() gopurs_runtime.Value {
	return MkFn0
}

func Get_mkFn10() gopurs_runtime.Value {
	return MkFn10
}

func Get_mkFn2() gopurs_runtime.Value {
	return MkFn2
}

func Get_mkFn3() gopurs_runtime.Value {
	return MkFn3
}

func Get_mkFn4() gopurs_runtime.Value {
	return MkFn4
}

func Get_mkFn5() gopurs_runtime.Value {
	return MkFn5
}

func Get_mkFn6() gopurs_runtime.Value {
	return MkFn6
}

func Get_mkFn7() gopurs_runtime.Value {
	return MkFn7
}

func Get_mkFn8() gopurs_runtime.Value {
	return MkFn8
}

func Get_mkFn9() gopurs_runtime.Value {
	return MkFn9
}

func Get_runFn0() gopurs_runtime.Value {
	return RunFn0
}

func Get_runFn10() gopurs_runtime.Value {
	return RunFn10
}

func Get_runFn2() gopurs_runtime.Value {
	return RunFn2
}

func Get_runFn3() gopurs_runtime.Value {
	return RunFn3
}

func Get_runFn4() gopurs_runtime.Value {
	return RunFn4
}

func Get_runFn5() gopurs_runtime.Value {
	return RunFn5
}

func Get_runFn6() gopurs_runtime.Value {
	return RunFn6
}

func Get_runFn7() gopurs_runtime.Value {
	return RunFn7
}

func Get_runFn8() gopurs_runtime.Value {
	return RunFn8
}

func Get_runFn9() gopurs_runtime.Value {
	return RunFn9
}
