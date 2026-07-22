package Data_Functor_Coproduct_Inject

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var prj gopurs_runtime.Value
var once_prj sync.Once
func Get_prj() gopurs_runtime.Value {
	once_prj.Do(func() {
		prj = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["prj"]
})
	})
	return prj
}

var injectReflexive gopurs_runtime.Value
var once_injectReflexive sync.Once
func Get_injectReflexive() gopurs_runtime.Value {
	once_injectReflexive.Do(func() {
		injectReflexive = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"inj": pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"], "prj": pkg_Data_Maybe.Get_Just()})
	})
	return injectReflexive
}

var injectLeft gopurs_runtime.Value
var once_injectLeft sync.Once
func Get_injectLeft() gopurs_runtime.Value {
	once_injectLeft.Do(func() {
		injectLeft = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"inj": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_0})
}), "prj": gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})})
	})
	return injectLeft
}

var inj gopurs_runtime.Value
var once_inj sync.Once
func Get_inj() gopurs_runtime.Value {
	once_inj.Do(func() {
		inj = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["inj"]
})
	})
	return inj
}

var injectRight gopurs_runtime.Value
var once_injectRight sync.Once
func Get_injectRight() gopurs_runtime.Value {
	once_injectRight.Do(func() {
		injectRight = gopurs_runtime.Func(func(dictInject_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"inj": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(dictInject_0.PtrVal.(map[string]gopurs_runtime.Value)["inj"], x_1)})
}), "prj": gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(dictInject_0.PtrVal.(map[string]gopurs_runtime.Value)["prj"], v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})})
})
	})
	return injectRight
}


