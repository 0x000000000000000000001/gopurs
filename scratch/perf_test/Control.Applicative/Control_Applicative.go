package Control_Applicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var pure gopurs_runtime.Value
var once_pure sync.Once
func Get_pure() gopurs_runtime.Value {
	once_pure.Do(func() {
		pure = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"]
})
	})
	return pure
}

var unless gopurs_runtime.Value
var once_unless sync.Once
func Get_unless() gopurs_runtime.Value {
	once_unless.Do(func() {
		unless = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.IntVal == 0)).IntVal != 0 {
__t0 = v1_2
} else {
if (v_1).IntVal != 0 {
__t0 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit())
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})
	})
	return unless
}

var when gopurs_runtime.Value
var once_when sync.Once
func Get_when() gopurs_runtime.Value {
	once_when.Do(func() {
		when = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (v_1).IntVal != 0 {
__t0 = v1_2
} else {
__t0 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit())
}
return __t0
})
})
})
	})
	return when
}

var liftA1 gopurs_runtime.Value
var once_liftA1 sync.Once
func Get_liftA1() gopurs_runtime.Value {
	once_liftA1.Do(func() {
		liftA1 = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], f_1)), a_2)
})
})
})
	})
	return liftA1
}

var applicativeProxy gopurs_runtime.Value
var once_applicativeProxy sync.Once
func Get_applicativeProxy() gopurs_runtime.Value {
	once_applicativeProxy.Do(func() {
		applicativeProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
})})
	})
	return applicativeProxy
}

var applicativeFn gopurs_runtime.Value
var once_applicativeFn sync.Once
func Get_applicativeFn() gopurs_runtime.Value {
	once_applicativeFn.Do(func() {
		applicativeFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
})})
	})
	return applicativeFn
}

var applicativeArray gopurs_runtime.Value
var once_applicativeArray sync.Once
func Get_applicativeArray() gopurs_runtime.Value {
	once_applicativeArray.Do(func() {
		applicativeArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{x_0})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
})})
	})
	return applicativeArray
}


