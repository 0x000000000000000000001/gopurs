package Data_Array_ST_Iterator

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		void = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return void
}

var Iterator gopurs_runtime.Value
var once_Iterator sync.Once
func Get_Iterator() gopurs_runtime.Value {
	once_Iterator.Do(func() {
		Iterator = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Iterator"), "value0": value0, "value1": value1})
})
})
	})
	return Iterator
}

var peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		peek = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_read(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], i_1))
}))
})
	})
	return peek
}

var next gopurs_runtime.Value
var once_next sync.Once
func Get_next() gopurs_runtime.Value {
	once_next.Do(func() {
		next = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_read(), __local_var_1_0)), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_modifyImpl(), gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), s_3), gopurs_runtime.Int(1))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": s_prime_4_1, "value": s_prime_4_1})
})), __local_var_1_0)), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], i_2))
}))
}))
})
	})
	return next
}

var pushWhile gopurs_runtime.Value
var once_pushWhile sync.Once
func Get_pushWhile() gopurs_runtime.Value {
	once_pushWhile.Do(func() {
		pushWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(iter_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(array_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_new_(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(break__3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), pkg_Data_HeytingAlgebra.Get_boolNot()), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_read(), break__3))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(Get_peek(), iter_1)), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(mx_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(p_0, mx_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), mx_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), array_2)), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(Get_next(), iter_1))
}))
} else {
__t0 = gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_write(), gopurs_runtime.Bool(true)), break__3))
}
return __t0
})))
}))
})
})
})
	})
	return pushWhile
}

var pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		pushAll = gopurs_runtime.Apply(Get_pushWhile(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return pushAll
}

var iterator gopurs_runtime.Value
var once_iterator sync.Once
func Get_iterator() gopurs_runtime.Value {
	once_iterator.Do(func() {
		iterator = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), gopurs_runtime.Apply(Get_Iterator(), f_0)), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_new_(), gopurs_runtime.Int(0)))
})
	})
	return iterator
}

var iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		iterate = gopurs_runtime.Func(func(iter_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_new_(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(break__2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), pkg_Data_HeytingAlgebra.Get_boolNot()), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_read(), break__2))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(Get_next(), iter_0)), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(mx_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_1, mx_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(mx_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_write(), gopurs_runtime.Bool(true)), break__2))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})))
}))
})
})
	})
	return iterate
}

var exhausted gopurs_runtime.Value
var once_exhausted sync.Once
func Get_exhausted() gopurs_runtime.Value {
	once_exhausted.Do(func() {
		exhausted = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_map_(), pkg_Data_Maybe.Get_isNothing())
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Apply(Get_peek(), x_1))
})
}()
	})
	return exhausted
}


