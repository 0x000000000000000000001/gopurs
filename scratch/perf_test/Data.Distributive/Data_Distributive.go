package Data_Distributive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var distributiveIdentity gopurs_runtime.Value
var once_distributiveIdentity sync.Once
func Get_distributiveIdentity() gopurs_runtime.Value {
	once_distributiveIdentity.Do(func() {
		distributiveIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"distribute": gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Unsafe_Coerce.Get_unsafeCoerce())
}), "collect": gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(f_1, x_2))
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
})})
	})
	return distributiveIdentity
}

var distribute gopurs_runtime.Value
var once_distribute sync.Once
func Get_distribute() gopurs_runtime.Value {
	once_distribute.Do(func() {
		distribute = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["distribute"]
})
	})
	return distribute
}

var distributiveFunction gopurs_runtime.Value
var once_distributiveFunction sync.Once
func Get_distributiveFunction() gopurs_runtime.Value {
	once_distributiveFunction.Do(func() {
		distributiveFunction = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"distribute": gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, e_2)
})), a_1)
})
})
}), "collect": gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(Get_distributiveFunction().PtrVal.(map[string]gopurs_runtime.Value)["distribute"], dictFunctor_0)
__local_var_3_1 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
})})
	})
	return distributiveFunction
}

var cotraverse gopurs_runtime.Value
var once_cotraverse sync.Once
func Get_cotraverse() gopurs_runtime.Value {
	once_cotraverse.Do(func() {
		cotraverse = gopurs_runtime.Func(func(dictDistributive_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
distribute2_2_0 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["distribute"], dictFunctor_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(distribute2_2_0, x_5))
})
})
})
})
	})
	return cotraverse
}

var collectDefault gopurs_runtime.Value
var once_collectDefault sync.Once
func Get_collectDefault() gopurs_runtime.Value {
	once_collectDefault.Do(func() {
		collectDefault = gopurs_runtime.Func(func(dictDistributive_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
distribute2_2_0 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["distribute"], dictFunctor_1)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(dictFunctor_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
})
})
	})
	return collectDefault
}

var distributiveTuple gopurs_runtime.Value
var once_distributiveTuple sync.Once
func Get_distributiveTuple() gopurs_runtime.Value {
	once_distributiveTuple.Do(func() {
		distributiveTuple = gopurs_runtime.Func(func(dictTypeEquals_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
from_1_0 := gopurs_runtime.Apply(dictTypeEquals_0.PtrVal.(map[string]gopurs_runtime.Value)["proof"], gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"collect": gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
distribute2_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_distributiveTuple(), dictTypeEquals_0).PtrVal.(map[string]gopurs_runtime.Value)["distribute"], dictFunctor_2)
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(dictFunctor_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_3_1, gopurs_runtime.Apply(__local_var_5_2, x_6))
})
})
}), "distribute": gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_3 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Apply(from_1_0, pkg_Data_Unit.Get_unit()))
__local_var_4_4 := gopurs_runtime.Apply(dictFunctor_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_snd())
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Apply(__local_var_4_4, x_5))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
})})
}
}()
})
	})
	return distributiveTuple
}

var collect gopurs_runtime.Value
var once_collect sync.Once
func Get_collect() gopurs_runtime.Value {
	once_collect.Do(func() {
		collect = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["collect"]
})
	})
	return collect
}

var distributeDefault gopurs_runtime.Value
var once_distributeDefault sync.Once
func Get_distributeDefault() gopurs_runtime.Value {
	once_distributeDefault.Do(func() {
		distributeDefault = gopurs_runtime.Func(func(dictDistributive_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["collect"], dictFunctor_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return distributeDefault
}


