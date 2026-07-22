package Control_Monad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
)

var whenM gopurs_runtime.Value
var once_whenM sync.Once
func Get_whenM() gopurs_runtime.Value {
	once_whenM.Do(func() {
		whenM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
when_1_0 := gopurs_runtime.Apply(pkg_Control_Applicative.Get_when(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(mb_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], mb_2), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(when_1_0, b_4), m_3)
}))
})
})
})
	})
	return whenM
}

var unlessM gopurs_runtime.Value
var once_unlessM sync.Once
func Get_unlessM() gopurs_runtime.Value {
	once_unlessM.Do(func() {
		unlessM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
unless_1_0 := gopurs_runtime.Apply(pkg_Control_Applicative.Get_unless(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(mb_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], mb_2), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(unless_1_0, b_4), m_3)
}))
})
})
})
	})
	return unlessM
}

var monadProxy gopurs_runtime.Value
var once_monadProxy sync.Once
func Get_monadProxy() gopurs_runtime.Value {
	once_monadProxy.Do(func() {
		monadProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeProxy()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindProxy()
})})
	})
	return monadProxy
}

var monadFn gopurs_runtime.Value
var once_monadFn sync.Once
func Get_monadFn() gopurs_runtime.Value {
	once_monadFn.Do(func() {
		monadFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeFn()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindFn()
})})
	})
	return monadFn
}

var monadArray gopurs_runtime.Value
var once_monadArray sync.Once
func Get_monadArray() gopurs_runtime.Value {
	once_monadArray.Do(func() {
		monadArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindArray()
})})
	})
	return monadArray
}

var liftM1 gopurs_runtime.Value
var once_liftM1 sync.Once
func Get_liftM1() gopurs_runtime.Value {
	once_liftM1.Do(func() {
		liftM1 = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_2), gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(f_1, a_prime_3))
}))
})
})
})
	})
	return liftM1
}

var ap gopurs_runtime.Value
var once_ap sync.Once
func Get_ap() gopurs_runtime.Value {
	once_ap.Do(func() {
		ap = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], f_2), gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_3), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
})
})
	})
	return ap
}


