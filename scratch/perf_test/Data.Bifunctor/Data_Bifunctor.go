package Data_Bifunctor

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var bimap gopurs_runtime.Value
var once_bimap sync.Once
func Get_bimap() gopurs_runtime.Value {
	once_bimap.Do(func() {
		bimap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"]
})
	})
	return bimap
}

var bivoid gopurs_runtime.Value
var once_bivoid sync.Once
func Get_bivoid() gopurs_runtime.Value {
	once_bivoid.Do(func() {
		bivoid = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
})
	})
	return bivoid
}

var lmap gopurs_runtime.Value
var once_lmap sync.Once
func Get_lmap() gopurs_runtime.Value {
	once_lmap.Do(func() {
		lmap = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBifunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return lmap
}

var rmap gopurs_runtime.Value
var once_rmap sync.Once
func Get_rmap() gopurs_runtime.Value {
	once_rmap.Do(func() {
		rmap = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictBifunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["bimap"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return rmap
}

var bifunctorTuple gopurs_runtime.Value
var once_bifunctorTuple sync.Once
func Get_bifunctorTuple() gopurs_runtime.Value {
	once_bifunctorTuple.Do(func() {
		bifunctorTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(g_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})})
	})
	return bifunctorTuple
}

var bifunctorEither gopurs_runtime.Value
var once_bifunctorEither sync.Once
func Get_bifunctorEither() gopurs_runtime.Value {
	once_bifunctorEither.Do(func() {
		bifunctorEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})})
	})
	return bifunctorEither
}

var bifunctorConst gopurs_runtime.Value
var once_bifunctorConst sync.Once
func Get_bifunctorConst() gopurs_runtime.Value {
	once_bifunctorConst.Do(func() {
		bifunctorConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
})})
	})
	return bifunctorConst
}


