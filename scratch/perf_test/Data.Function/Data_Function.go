package Data_Function

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var on gopurs_runtime.Value
var once_on sync.Once
func Get_on() gopurs_runtime.Value {
	once_on.Do(func() {
		on = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2)), gopurs_runtime.Apply(g_1, y_3))
})
})
})
})
	})
	return on
}

var flip gopurs_runtime.Value
var once_flip sync.Once
func Get_flip() gopurs_runtime.Value {
	once_flip.Do(func() {
		flip = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, a_2), b_1)
})
})
})
	})
	return flip
}

var const_ gopurs_runtime.Value
var once_const_ sync.Once
func Get_const_() gopurs_runtime.Value {
	once_const_.Do(func() {
		const_ = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
})
	})
	return const_
}

var applyN gopurs_runtime.Value
var once_applyN sync.Once
func Get_applyN() gopurs_runtime.Value {
	once_applyN.Do(func() {
		applyN = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(n_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_2), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t1 = acc_3
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), n_2), gopurs_runtime.Int(1))), gopurs_runtime.Apply(f_0, acc_3))
}
return __t1
})
})
return go__1_0
})
	})
	return applyN
}

var applyFlipped gopurs_runtime.Value
var once_applyFlipped sync.Once
func Get_applyFlipped() gopurs_runtime.Value {
	once_applyFlipped.Do(func() {
		applyFlipped = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_0)
})
})
	})
	return applyFlipped
}

var apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		apply = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
})
})
	})
	return apply
}


