package Data_Semigroup_Last

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Last gopurs_runtime.Value
var once_Last sync.Once
func Get_Last() gopurs_runtime.Value {
	once_Last.Do(func() {
		Last = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Last
}

var showLast gopurs_runtime.Value
var once_showLast sync.Once
func Get_showLast() gopurs_runtime.Value {
	once_showLast.Do(func() {
		showLast = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Last ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showLast
}

var semigroupLast gopurs_runtime.Value
var once_semigroupLast sync.Once
func Get_semigroupLast() gopurs_runtime.Value {
	once_semigroupLast.Do(func() {
		semigroupLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
})})
	})
	return semigroupLast
}

var ordLast gopurs_runtime.Value
var once_ordLast sync.Once
func Get_ordLast() gopurs_runtime.Value {
	once_ordLast.Do(func() {
		ordLast = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordLast
}

var functorLast gopurs_runtime.Value
var once_functorLast sync.Once
func Get_functorLast() gopurs_runtime.Value {
	once_functorLast.Do(func() {
		functorLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})
	})
	return functorLast
}

var eqLast gopurs_runtime.Value
var once_eqLast sync.Once
func Get_eqLast() gopurs_runtime.Value {
	once_eqLast.Do(func() {
		eqLast = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqLast
}

var eq1Last gopurs_runtime.Value
var once_eq1Last sync.Once
func Get_eq1Last() gopurs_runtime.Value {
	once_eq1Last.Do(func() {
		eq1Last = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1Last
}

var ord1Last gopurs_runtime.Value
var once_ord1Last sync.Once
func Get_ord1Last() gopurs_runtime.Value {
	once_ord1Last.Do(func() {
		ord1Last = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Last()
})})
	})
	return ord1Last
}

var boundedLast gopurs_runtime.Value
var once_boundedLast sync.Once
func Get_boundedLast() gopurs_runtime.Value {
	once_boundedLast.Do(func() {
		boundedLast = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedLast
}

var applyLast gopurs_runtime.Value
var once_applyLast sync.Once
func Get_applyLast() gopurs_runtime.Value {
	once_applyLast.Do(func() {
		applyLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLast()
})})
	})
	return applyLast
}

var bindLast gopurs_runtime.Value
var once_bindLast sync.Once
func Get_bindLast() gopurs_runtime.Value {
	once_bindLast.Do(func() {
		bindLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLast()
})})
	})
	return bindLast
}

var applicativeLast gopurs_runtime.Value
var once_applicativeLast sync.Once
func Get_applicativeLast() gopurs_runtime.Value {
	once_applicativeLast.Do(func() {
		applicativeLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_Last(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLast()
})})
	})
	return applicativeLast
}

var monadLast gopurs_runtime.Value
var once_monadLast sync.Once
func Get_monadLast() gopurs_runtime.Value {
	once_monadLast.Do(func() {
		monadLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeLast()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindLast()
})})
	})
	return monadLast
}


