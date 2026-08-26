package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_empty gopurs_runtime.Value
var once_Main_empty sync.Once

func Get_Main_empty() gopurs_runtime.Value {
	once_Main_empty.Do(func() {
		cache_Main_empty = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_empty(r_0_box, f_1_box)
		})
	})
	return cache_Main_empty
}

var cache_Main_empty__2426865526 gopurs_runtime.Value
var once_Main_empty__2426865526 sync.Once

func Get_Main_empty__2426865526() gopurs_runtime.Value {
	once_Main_empty__2426865526.Do(func() {
		cache_Main_empty__2426865526 = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_empty__2426865526(r_0_box, f_1_box)
		})
	})
	return cache_Main_empty__2426865526
}

var cache_Main_cons gopurs_runtime.Value
var once_Main_cons sync.Once

func Get_Main_cons() gopurs_runtime.Value {
	once_Main_cons.Do(func() {
		cache_Main_cons = gopurs_runtime.Func4(func(a_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons(a_0_box, l_1_box, r_2_box, f_3_box)
		})
	})
	return cache_Main_cons
}

var cache_Main_cons__3906071881 gopurs_runtime.Value
var once_Main_cons__3906071881 sync.Once

func Get_Main_cons__3906071881() gopurs_runtime.Value {
	once_Main_cons__3906071881.Do(func() {
		cache_Main_cons__3906071881 = gopurs_runtime.Func4(func(a_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons__3906071881(a_0_box, l_1_box, r_2_box, f_3_box)
		})
	})
	return cache_Main_cons__3906071881
}

var cache_Main_go__append gopurs_runtime.Value
var once_Main_go__append sync.Once

func Get_Main_go__append() gopurs_runtime.Value {
	once_Main_go__append.Do(func() {
		cache_Main_go__append = gopurs_runtime.Func4(func(l1_0_box gopurs_runtime.Value, l2_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_go__append(l1_0_box, l2_1_box, r_2_box, f_3_box)
		})
	})
	return cache_Main_go__append
}

var cache_Main_append__3802550593 gopurs_runtime.Value
var once_Main_append__3802550593 sync.Once

func Get_Main_append__3802550593() gopurs_runtime.Value {
	once_Main_append__3802550593.Do(func() {
		cache_Main_append__3802550593 = gopurs_runtime.Func4(func(l1_0_box gopurs_runtime.Value, l2_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_append__3802550593(l1_0_box, l2_1_box, r_2_box, f_3_box)
		})
	})
	return cache_Main_append__3802550593
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(r_0_box, f_1_box)
		})
	})
	return cache_Main_test
}

func Call_Main_empty(r_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return r_0
}

func Call_Main_empty__2426865526(r_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return r_0
}

func Call_Main_cons(a_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var l_1 gopurs_runtime.Value = l_1_loop
	_ = l_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var f_3 gopurs_runtime.Value = f_3_loop
	_ = f_3
	return gopurs_runtime.Apply2(f_3, a_0, gopurs_runtime.Apply2(l_1, r_2, f_3))
}

func Call_Main_cons__3906071881(a_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var l_1 gopurs_runtime.Value = l_1_loop
	_ = l_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var f_3 gopurs_runtime.Value = f_3_loop
	_ = f_3
	return gopurs_runtime.Apply2(f_3, a_0, gopurs_runtime.Apply2(l_1, r_2, f_3))
}

func Call_Main_go__append(l1_0_loop gopurs_runtime.Value, l2_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var l1_0 gopurs_runtime.Value = l1_0_loop
	_ = l1_0
	var l2_1 gopurs_runtime.Value = l2_1_loop
	_ = l2_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var f_3 gopurs_runtime.Value = f_3_loop
	_ = f_3
	return gopurs_runtime.Apply2(l2_1, gopurs_runtime.Apply2(l1_0, r_2, f_3), f_3)
}

func Call_Main_append__3802550593(l1_0_loop gopurs_runtime.Value, l2_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var l1_0 gopurs_runtime.Value = l1_0_loop
	_ = l1_0
	var l2_1 gopurs_runtime.Value = l2_1_loop
	_ = l2_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	var f_3 gopurs_runtime.Value = f_3_loop
	_ = f_3
	return gopurs_runtime.Apply2(l2_1, gopurs_runtime.Apply2(l1_0, r_2, f_3), f_3)
}

func Call_Main_test(r_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply2(f_1, gopurs_runtime.Int(2), gopurs_runtime.Apply2(f_1, gopurs_runtime.Int(1), r_0))
}
