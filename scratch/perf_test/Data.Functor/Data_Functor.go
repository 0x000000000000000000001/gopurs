package Data_Functor

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		map_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["map"]
})
	})
	return map_
}

var mapFlipped gopurs_runtime.Value
var once_mapFlipped sync.Once
func Get_mapFlipped() gopurs_runtime.Value {
	once_mapFlipped.Do(func() {
		mapFlipped = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fa_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), fa_1)
})
})
})
	})
	return mapFlipped
}

var void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		void = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
})
	})
	return void
}

var voidLeft gopurs_runtime.Value
var once_voidLeft sync.Once
func Get_voidLeft() gopurs_runtime.Value {
	once_voidLeft.Do(func() {
		voidLeft = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
})), f_1)
})
})
})
	})
	return voidLeft
}

var voidRight gopurs_runtime.Value
var once_voidRight sync.Once
func Get_voidRight() gopurs_runtime.Value {
	once_voidRight.Do(func() {
		voidRight = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
})
})
	})
	return voidRight
}

var functorProxy gopurs_runtime.Value
var once_functorProxy sync.Once
func Get_functorProxy() gopurs_runtime.Value {
	once_functorProxy.Do(func() {
		functorProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})})
	})
	return functorProxy
}

var functorFn gopurs_runtime.Value
var once_functorFn sync.Once
func Get_functorFn() gopurs_runtime.Value {
	once_functorFn.Do(func() {
		functorFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": pkg_Control_Semigroupoid.Get_semigroupoidFn().PtrVal.(map[string]gopurs_runtime.Value)["compose"]})
	})
	return functorFn
}

var functorArray gopurs_runtime.Value
var once_functorArray sync.Once
func Get_functorArray() gopurs_runtime.Value {
	once_functorArray.Do(func() {
		functorArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": Get_arrayMap()})
	})
	return functorArray
}

var flap gopurs_runtime.Value
var once_flap sync.Once
func Get_flap() gopurs_runtime.Value {
	once_flap.Do(func() {
		flap = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ff_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_2)
})), ff_1)
})
})
})
	})
	return flap
}

func Get_arrayMap() gopurs_runtime.Value {
	return ArrayMap
}
