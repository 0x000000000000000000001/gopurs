package Control_Bind

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		discard = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["discard"]
})
	})
	return discard
}

var bindProxy gopurs_runtime.Value
var once_bindProxy sync.Once
func Get_bindProxy() gopurs_runtime.Value {
	once_bindProxy.Do(func() {
		bindProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
})})
	})
	return bindProxy
}

var bindFn gopurs_runtime.Value
var once_bindFn sync.Once
func Get_bindFn() gopurs_runtime.Value {
	once_bindFn.Do(func() {
		bindFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(m_0, x_2)), x_2)
})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
})})
	})
	return bindFn
}

var bindArray gopurs_runtime.Value
var once_bindArray sync.Once
func Get_bindArray() gopurs_runtime.Value {
	once_bindArray.Do(func() {
		bindArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": Get_arrayBind(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
})})
	})
	return bindArray
}

var bind gopurs_runtime.Value
var once_bind sync.Once
func Get_bind() gopurs_runtime.Value {
	once_bind.Do(func() {
		bind = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"]
})
	})
	return bind
}

var bindFlipped gopurs_runtime.Value
var once_bindFlipped sync.Once
func Get_bindFlipped() gopurs_runtime.Value {
	once_bindFlipped.Do(func() {
		bindFlipped = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_2), b_1)
})
})
})
	})
	return bindFlipped
}

var composeKleisliFlipped gopurs_runtime.Value
var once_composeKleisliFlipped sync.Once
func Get_composeKleisliFlipped() gopurs_runtime.Value {
	once_composeKleisliFlipped.Do(func() {
		composeKleisliFlipped = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(g_2, a_3)), f_1)
})
})
})
})
	})
	return composeKleisliFlipped
}

var composeKleisli gopurs_runtime.Value
var once_composeKleisli sync.Once
func Get_composeKleisli() gopurs_runtime.Value {
	once_composeKleisli.Do(func() {
		composeKleisli = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(f_1, a_3)), g_2)
})
})
})
})
	})
	return composeKleisli
}

var discardProxy gopurs_runtime.Value
var once_discardProxy sync.Once
func Get_discardProxy() gopurs_runtime.Value {
	once_discardProxy.Do(func() {
		discardProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"discard": gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"]
})})
	})
	return discardProxy
}

var discardUnit gopurs_runtime.Value
var once_discardUnit sync.Once
func Get_discardUnit() gopurs_runtime.Value {
	once_discardUnit.Do(func() {
		discardUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"discard": gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"]
})})
	})
	return discardUnit
}

var ifM gopurs_runtime.Value
var once_ifM sync.Once
func Get_ifM() gopurs_runtime.Value {
	once_ifM.Do(func() {
		ifM = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cond_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], cond_1), gopurs_runtime.Func(func(cond_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (cond_prime_4).IntVal != 0 {
__t0 = t_2
} else {
__t0 = f_3
}
return __t0
}))
})
})
})
})
	})
	return ifM
}

var join gopurs_runtime.Value
var once_join sync.Once
func Get_join() gopurs_runtime.Value {
	once_join.Do(func() {
		join = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], m_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return join
}

func Get_arrayBind() gopurs_runtime.Value {
	return ArrayBind
}
