package Data_Profunctor_Strong

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var strongFn gopurs_runtime.Value
var once_strongFn sync.Once
func Get_strongFn() gopurs_runtime.Value {
	once_strongFn.Do(func() {
		strongFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"first": gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(a2b_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
}), "second": pkg_Data_Tuple.Get_functorTuple().PtrVal.(map[string]gopurs_runtime.Value)["map"], "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
})})
	})
	return strongFn
}

var second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		second = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["second"]
})
	})
	return second
}

var first gopurs_runtime.Value
var once_first sync.Once
func Get_first() gopurs_runtime.Value {
	once_first.Do(func() {
		first = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["first"]
})
	})
	return first
}

var splitStrong gopurs_runtime.Value
var once_splitStrong sync.Once
func Get_splitStrong() gopurs_runtime.Value {
	once_splitStrong.Do(func() {
		splitStrong = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictStrong_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupoid_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], gopurs_runtime.Apply(dictStrong_1.PtrVal.(map[string]gopurs_runtime.Value)["second"], r_3)), gopurs_runtime.Apply(dictStrong_1.PtrVal.(map[string]gopurs_runtime.Value)["first"], l_2))
})
})
})
})
	})
	return splitStrong
}

var fanout gopurs_runtime.Value
var once_fanout sync.Once
func Get_fanout() gopurs_runtime.Value {
	once_fanout.Do(func() {
		fanout = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictStrong_1 gopurs_runtime.Value) gopurs_runtime.Value {
lcmap_2_0 := gopurs_runtime.Apply(pkg_Data_Profunctor.Get_lcmap(), gopurs_runtime.Apply(dictStrong_1.PtrVal.(map[string]gopurs_runtime.Value)["Profunctor0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(lcmap_2_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_5, "value1": a_5})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupoid_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], gopurs_runtime.Apply(dictStrong_1.PtrVal.(map[string]gopurs_runtime.Value)["second"], r_4)), gopurs_runtime.Apply(dictStrong_1.PtrVal.(map[string]gopurs_runtime.Value)["first"], l_3)))
})
})
})
})
	})
	return fanout
}


