package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_N gopurs_runtime.Value
var once_Main_N sync.Once

func Get_Main_N() gopurs_runtime.Value {
	once_Main_N.Do(func() {
		cache_Main_N = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_N(x_0_box)
		})
	})
	return cache_Main_N
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fst gopurs_runtime.Value
var once_Main_fst sync.Once

func Get_Main_fst() gopurs_runtime.Value {
	once_Main_fst.Do(func() {
		cache_Main_fst = gopurs_runtime.RecordDict2("get", "set", gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordGet(p_0, "fst")
		}), gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("fst", "snd", a_1, gopurs_runtime.RecordGet(p_0, "snd"))
			})
		}))
	})
	return cache_Main_fst
}

var cache_Main_fst__562322518 gopurs_runtime.Value
var once_Main_fst__562322518 sync.Once

func Get_Main_fst__562322518() gopurs_runtime.Value {
	once_Main_fst__562322518.Do(func() {
		cache_Main_fst__562322518 = gopurs_runtime.RecordDict2("get", "set", gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordGet(p_0, "fst")
		}), gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("fst", "snd", a_1, gopurs_runtime.RecordGet(p_0, "snd"))
			})
		}))
	})
	return cache_Main_fst__562322518
}

var cache_Main_fst__3521275305 gopurs_runtime.Value
var once_Main_fst__3521275305 sync.Once

func Get_Main_fst__3521275305() gopurs_runtime.Value {
	once_Main_fst__3521275305.Do(func() {
		cache_Main_fst__3521275305 = gopurs_runtime.RecordDict2("get", "set", gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordGet(p_0, "fst")
		}), gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("fst", "snd", a_1, gopurs_runtime.RecordGet(p_0, "snd"))
			})
		}))
	})
	return cache_Main_fst__3521275305
}

var cache_Main_composeLenses gopurs_runtime.Value
var once_Main_composeLenses sync.Once

func Get_Main_composeLenses() gopurs_runtime.Value {
	once_Main_composeLenses.Do(func() {
		cache_Main_composeLenses = gopurs_runtime.Func2(func(l1_0_box gopurs_runtime.Value, l2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_composeLenses(l1_0_box, l2_1_box)
		})
	})
	return cache_Main_composeLenses
}

var cache_Main_composeLenses__1870392786 gopurs_runtime.Value
var once_Main_composeLenses__1870392786 sync.Once

func Get_Main_composeLenses__1870392786() gopurs_runtime.Value {
	once_Main_composeLenses__1870392786.Do(func() {
		cache_Main_composeLenses__1870392786 = gopurs_runtime.Func2(func(l1_0_box gopurs_runtime.Value, l2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_composeLenses__1870392786(l1_0_box, l2_1_box)
		})
	})
	return cache_Main_composeLenses__1870392786
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.RecordDict2("get", "set", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(a_0, "fst"), "fst")
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("fst", "snd", gopurs_runtime.RecordDict2("fst", "snd", c_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(a_0, "fst"), "snd")), gopurs_runtime.RecordGet(a_0, "snd"))
			})
		}))
	})
	return cache_Main_test1
}

func Call_Main_N(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_composeLenses(l1_0_loop gopurs_runtime.Value, l2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var l1_0 gopurs_runtime.Value = l1_0_loop
	_ = l1_0
	var l2_1 gopurs_runtime.Value = l2_1_loop
	_ = l2_1
	return gopurs_runtime.RecordDict2("get", "set", gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(l2_1, "get"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(l1_0, "get"), a_2))
	}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(l1_0, "set"), a_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(l2_1, "set"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(l1_0, "get"), a_2), c_3))
		})
	}))
}

func Call_Main_composeLenses__1870392786(l1_0_loop gopurs_runtime.Value, l2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var l1_0 gopurs_runtime.Value = l1_0_loop
	_ = l1_0
	var l2_1 gopurs_runtime.Value = l2_1_loop
	_ = l2_1
	return gopurs_runtime.RecordDict2("get", "set", gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(l2_1, "get"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(l1_0, "get"), a_2))
	}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(l1_0, "set"), a_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(l2_1, "set"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(l1_0, "get"), a_2), c_3))
		})
	}))
}
