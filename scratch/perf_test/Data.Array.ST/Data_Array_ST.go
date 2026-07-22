package Data_Array_ST

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Control_Monad_ST_Uncurried "gopurs/output/Control.Monad.ST.Uncurried"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var unshiftAll gopurs_runtime.Value
var once_unshiftAll sync.Once
func Get_unshiftAll() gopurs_runtime.Value {
	once_unshiftAll.Do(func() {
		unshiftAll = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl())
	})
	return unshiftAll
}

var unshift gopurs_runtime.Value
var once_unshift sync.Once
func Get_unshift() gopurs_runtime.Value {
	once_unshift.Do(func() {
		unshift = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_unshiftAllImpl()), gopurs_runtime.Array([]gopurs_runtime.Value{a_0}))
})
	})
	return unshift
}

var unsafeThaw gopurs_runtime.Value
var once_unsafeThaw sync.Once
func Get_unsafeThaw() gopurs_runtime.Value {
	once_unsafeThaw.Do(func() {
		unsafeThaw = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeThawImpl())
	})
	return unsafeThaw
}

var unsafeFreeze gopurs_runtime.Value
var once_unsafeFreeze sync.Once
func Get_unsafeFreeze() gopurs_runtime.Value {
	once_unsafeFreeze.Do(func() {
		unsafeFreeze = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_unsafeFreezeImpl())
	})
	return unsafeFreeze
}

var toAssocArray gopurs_runtime.Value
var once_toAssocArray sync.Once
func Get_toAssocArray() gopurs_runtime.Value {
	once_toAssocArray.Do(func() {
		toAssocArray = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_toAssocArrayImpl())
	})
	return toAssocArray
}

var thaw gopurs_runtime.Value
var once_thaw sync.Once
func Get_thaw() gopurs_runtime.Value {
	once_thaw.Do(func() {
		thaw = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_thawImpl())
	})
	return thaw
}

var withArray gopurs_runtime.Value
var once_withArray sync.Once
func Get_withArray() gopurs_runtime.Value {
	once_withArray.Do(func() {
		withArray = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(Get_thaw(), xs_1)), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(f_0, result_2)), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unsafeFreeze(), result_2)
}))
}))
})
})
	})
	return withArray
}

var splice gopurs_runtime.Value
var once_splice sync.Once
func Get_splice() gopurs_runtime.Value {
	once_splice.Do(func() {
		splice = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_spliceImpl())
	})
	return splice
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_sortByImpl()), comp_0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
} else {
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Int(0)), gopurs_runtime.Int(1))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
}))
})
	})
	return sortBy
}

var sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		sortWith = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(f_1, x_2)), gopurs_runtime.Apply(f_1, y_3))
})
}))
})
})
	})
	return sortWith
}

var sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		sort = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_sortBy(), dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"])
})
	})
	return sort
}

var shift gopurs_runtime.Value
var once_shift sync.Once
func Get_shift() gopurs_runtime.Value {
	once_shift.Do(func() {
		shift = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_shiftImpl()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return shift
}

var run gopurs_runtime.Value
var once_run sync.Once
func Get_run() gopurs_runtime.Value {
	once_run.Do(func() {
		run = gopurs_runtime.Func(func(st_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), st_0), Get_unsafeFreeze()))
})
	})
	return run
}

var pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		pushAll = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushAllImpl())
	})
	return pushAll
}

var push gopurs_runtime.Value
var once_push sync.Once
func Get_push() gopurs_runtime.Value {
	once_push.Do(func() {
		push = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_pushImpl())
	})
	return push
}

var pop gopurs_runtime.Value
var once_pop sync.Once
func Get_pop() gopurs_runtime.Value {
	once_pop.Do(func() {
		pop = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_popImpl()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return pop
}

var poke gopurs_runtime.Value
var once_poke sync.Once
func Get_poke() gopurs_runtime.Value {
	once_poke.Do(func() {
		poke = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_pokeImpl())
	})
	return poke
}

var peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		peek = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), Get_peekImpl()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return peek
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_peek(), i_0), xs_2)), gopurs_runtime.Func(func(entry_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(entry_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_poke(), i_0), gopurs_runtime.Apply(f_1, entry_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), xs_2)
} else {
if (gopurs_runtime.Bool(entry_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_pure_(), gopurs_runtime.Bool(false))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}))
})
})
})
	})
	return modify
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_lengthImpl())
	})
	return length
}

var freeze gopurs_runtime.Value
var once_freeze sync.Once
func Get_freeze() gopurs_runtime.Value {
	once_freeze.Do(func() {
		freeze = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_freezeImpl())
	})
	return freeze
}

var clone gopurs_runtime.Value
var once_clone sync.Once
func Get_clone() gopurs_runtime.Value {
	once_clone.Do(func() {
		clone = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), Get_cloneImpl())
	})
	return clone
}

func Get_cloneImpl() gopurs_runtime.Value {
	return CloneImpl
}

func Get_freezeImpl() gopurs_runtime.Value {
	return FreezeImpl
}

func Get_lengthImpl() gopurs_runtime.Value {
	return LengthImpl
}

func Get_new_() gopurs_runtime.Value {
	return New_
}

func Get_peekImpl() gopurs_runtime.Value {
	return PeekImpl
}

func Get_pokeImpl() gopurs_runtime.Value {
	return PokeImpl
}

func Get_popImpl() gopurs_runtime.Value {
	return PopImpl
}

func Get_pushAllImpl() gopurs_runtime.Value {
	return PushAllImpl
}

func Get_pushImpl() gopurs_runtime.Value {
	return PushImpl
}

func Get_shiftImpl() gopurs_runtime.Value {
	return ShiftImpl
}

func Get_sortByImpl() gopurs_runtime.Value {
	return SortByImpl
}

func Get_spliceImpl() gopurs_runtime.Value {
	return SpliceImpl
}

func Get_thawImpl() gopurs_runtime.Value {
	return ThawImpl
}

func Get_toAssocArrayImpl() gopurs_runtime.Value {
	return ToAssocArrayImpl
}

func Get_unsafeFreezeImpl() gopurs_runtime.Value {
	return UnsafeFreezeImpl
}

func Get_unsafeThawImpl() gopurs_runtime.Value {
	return UnsafeThawImpl
}

func Get_unshiftAllImpl() gopurs_runtime.Value {
	return UnshiftAllImpl
}
