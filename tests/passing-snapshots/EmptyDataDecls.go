package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_ArrayBox gopurs_runtime.Value
var once_Main_ArrayBox sync.Once

func Get_Main_ArrayBox() gopurs_runtime.Value {
	once_Main_ArrayBox.Do(func() {
		cache_Main_ArrayBox = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_ArrayBox
}

var cache_Main_nil gopurs_runtime.Value
var once_Main_nil sync.Once

func Get_Main_nil() gopurs_runtime.Value {
	once_Main_nil.Do(func() {
		cache_Main_nil = gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())
	})
	return cache_Main_nil
}

var cache_Main_cons_prime_ gopurs_runtime.Value
var once_Main_cons_prime_ sync.Once

func Get_Main_cons_prime_() gopurs_runtime.Value {
	once_Main_cons_prime_.Do(func() {
		cache_Main_cons_prime_ = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons_prime_(x_0_box, v_1_box)
		})
	})
	return cache_Main_cons_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_cons_prime___278004513 gopurs_runtime.Value
var once_Main_cons_prime___278004513 sync.Once

func Get_Main_cons_prime___278004513() gopurs_runtime.Value {
	once_Main_cons_prime___278004513.Do(func() {
		cache_Main_cons_prime___278004513 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons_prime___278004513(x_0_box, v_1_box)
		})
	})
	return cache_Main_cons_prime___278004513
}

var cache_Main_cons_prime___1420788263 gopurs_runtime.Value
var once_Main_cons_prime___1420788263 sync.Once

func Get_Main_cons_prime___1420788263() gopurs_runtime.Value {
	once_Main_cons_prime___1420788263.Do(func() {
		cache_Main_cons_prime___1420788263 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons_prime___1420788263(x_0_box, v_1_box)
		})
	})
	return cache_Main_cons_prime___1420788263
}

var cache_Main_nil__527179238 gopurs_runtime.Value
var once_Main_nil__527179238 sync.Once

func Get_Main_nil__527179238() gopurs_runtime.Value {
	once_Main_nil__527179238.Do(func() {
		cache_Main_nil__527179238 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())
	})
	return cache_Main_nil__527179238
}

type Constructor_Main_ArrayBox struct {
	Rc uint32
	V0 []gopurs_runtime.Value
}

func Call_Main_cons_prime_(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Array(func() []gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), v_1).UnsafePtr)
		unboxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = v
		}
		return unboxed
	}())
}

func Call_Main_cons_prime___278004513(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Array(func() []gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), v_1).UnsafePtr)
		unboxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = v
		}
		return unboxed
	}())
}

func Call_Main_cons_prime___1420788263(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Array(func() []gopurs_runtime.Value {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), v_1).UnsafePtr)
		unboxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			unboxed[i] = v
		}
		return unboxed
	}())
}
