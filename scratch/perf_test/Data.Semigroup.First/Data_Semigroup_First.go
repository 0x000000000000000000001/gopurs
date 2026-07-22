package Data_Semigroup_First

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var First gopurs_runtime.Value
var once_First sync.Once
func Get_First() gopurs_runtime.Value {
	once_First.Do(func() {
		First = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return First
}

var showFirst gopurs_runtime.Value
var once_showFirst sync.Once
func Get_showFirst() gopurs_runtime.Value {
	once_showFirst.Do(func() {
		showFirst = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(First ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showFirst
}

var semigroupFirst gopurs_runtime.Value
var once_semigroupFirst sync.Once
func Get_semigroupFirst() gopurs_runtime.Value {
	once_semigroupFirst.Do(func() {
		semigroupFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
})})
	})
	return semigroupFirst
}

var ordFirst gopurs_runtime.Value
var once_ordFirst sync.Once
func Get_ordFirst() gopurs_runtime.Value {
	once_ordFirst.Do(func() {
		ordFirst = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordFirst
}

var functorFirst gopurs_runtime.Value
var once_functorFirst sync.Once
func Get_functorFirst() gopurs_runtime.Value {
	once_functorFirst.Do(func() {
		functorFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})
	})
	return functorFirst
}

var eqFirst gopurs_runtime.Value
var once_eqFirst sync.Once
func Get_eqFirst() gopurs_runtime.Value {
	once_eqFirst.Do(func() {
		eqFirst = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqFirst
}

var eq1First gopurs_runtime.Value
var once_eq1First sync.Once
func Get_eq1First() gopurs_runtime.Value {
	once_eq1First.Do(func() {
		eq1First = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1First
}

var ord1First gopurs_runtime.Value
var once_ord1First sync.Once
func Get_ord1First() gopurs_runtime.Value {
	once_ord1First.Do(func() {
		ord1First = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1First()
})})
	})
	return ord1First
}

var boundedFirst gopurs_runtime.Value
var once_boundedFirst sync.Once
func Get_boundedFirst() gopurs_runtime.Value {
	once_boundedFirst.Do(func() {
		boundedFirst = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedFirst
}

var applyFirst gopurs_runtime.Value
var once_applyFirst sync.Once
func Get_applyFirst() gopurs_runtime.Value {
	once_applyFirst.Do(func() {
		applyFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorFirst()
})})
	})
	return applyFirst
}

var bindFirst gopurs_runtime.Value
var once_bindFirst sync.Once
func Get_bindFirst() gopurs_runtime.Value {
	once_bindFirst.Do(func() {
		bindFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFirst()
})})
	})
	return bindFirst
}

var applicativeFirst gopurs_runtime.Value
var once_applicativeFirst sync.Once
func Get_applicativeFirst() gopurs_runtime.Value {
	once_applicativeFirst.Do(func() {
		applicativeFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_First(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFirst()
})})
	})
	return applicativeFirst
}

var monadFirst gopurs_runtime.Value
var once_monadFirst sync.Once
func Get_monadFirst() gopurs_runtime.Value {
	once_monadFirst.Do(func() {
		monadFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeFirst()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindFirst()
})})
	})
	return monadFirst
}


