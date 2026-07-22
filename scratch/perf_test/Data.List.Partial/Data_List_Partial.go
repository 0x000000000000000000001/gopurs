package Data_List_Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t0
})
})
	})
	return tail
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_last(), gopurs_runtime.Value{}), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t0
}
}()
})
})
	})
	return last
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_init_(), gopurs_runtime.Value{}), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t0
}
}()
})
})
	})
	return init_
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t0 = v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t0
})
})
	})
	return head
}


