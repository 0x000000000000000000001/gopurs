package Lib

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Tuple gopurs_runtime.Value
var once_Tuple sync.Once
func Get_Tuple() gopurs_runtime.Value {
	once_Tuple.Do(func() {
		Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": value0, "value1": value1})
})
})
	})
	return Tuple
}

var testInt gopurs_runtime.Value
var once_testInt sync.Once
func Get_testInt() gopurs_runtime.Value {
	once_testInt.Do(func() {
		testInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"runTest": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("4")
})})
	})
	return testInt
}

var runTest gopurs_runtime.Value
var once_runTest sync.Once
func Get_runTest() gopurs_runtime.Value {
	once_runTest.Do(func() {
		runTest = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["runTest"]
})
	})
	return runTest
}

var test/\ gopurs_runtime.Value
var once_test/\ sync.Once
func Get_test/\() gopurs_runtime.Value {
	once_test/\.Do(func() {
		test/\ = gopurs_runtime.Func(func(dictTest_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTest1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"runTest": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictTest_0.PtrVal.(map[string]gopurs_runtime.Value)["runTest"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(dictTest1_1.PtrVal.(map[string]gopurs_runtime.Value)["runTest"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})})
})
})
	})
	return test/\
}


