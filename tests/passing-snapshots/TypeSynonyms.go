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
		cache_Main_fst = func() gopurs_runtime.Value {
			orig := struct {
				get gopurs_runtime.Value
				set gopurs_runtime.Value
			}{gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordGet(p_0, "fst")
			}), gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					orig := struct {
						fst gopurs_runtime.Value
						snd gopurs_runtime.Value
					}{a_1, gopurs_runtime.RecordGet(p_0, "snd")}
					_ = orig
					return gopurs_runtime.RecordDict2("fst", "snd", orig.fst, orig.snd)
				}()
			})}
			_ = orig
			return gopurs_runtime.RecordDict2("get", "set", orig.get, orig.set)
		}()
	})
	return cache_Main_fst
}

var cache_Main_composeLenses gopurs_runtime.Value
var once_Main_composeLenses sync.Once

func Get_Main_composeLenses() gopurs_runtime.Value {
	once_Main_composeLenses.Do(func() {
		cache_Main_composeLenses = gopurs_runtime.Func2(func(l1_0_box gopurs_runtime.Value, l2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_composeLenses(func() struct {
					get gopurs_runtime.Value
					set gopurs_runtime.Value
				} {
					orig := l1_0_box
					_ = orig
					clone := struct {
						get gopurs_runtime.Value
						set gopurs_runtime.Value
					}{}
					clone.get = gopurs_runtime.RecordGet(orig, "get")
					clone.set = gopurs_runtime.RecordGet(orig, "set")
					return clone
				}(), func() struct {
					get gopurs_runtime.Value
					set gopurs_runtime.Value
				} {
					orig := l2_1_box
					_ = orig
					clone := struct {
						get gopurs_runtime.Value
						set gopurs_runtime.Value
					}{}
					clone.get = gopurs_runtime.RecordGet(orig, "get")
					clone.set = gopurs_runtime.RecordGet(orig, "set")
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("get", "set", orig.get, orig.set)
			}()
		})
	})
	return cache_Main_composeLenses
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = func() gopurs_runtime.Value {
			orig := struct {
				get gopurs_runtime.Value
				set gopurs_runtime.Value
			}{gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(a_0, "fst"), "fst")
			}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, c_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					orig := func() struct {
						fst struct {
							fst gopurs_runtime.Value
							snd gopurs_runtime.Value
						}
						snd gopurs_runtime.Value
					} {
						orig := func() gopurs_runtime.Value {
							orig := struct {
								fst gopurs_runtime.Value
								snd gopurs_runtime.Value
							}{func() gopurs_runtime.Value {
								orig := struct {
									fst gopurs_runtime.Value
									snd gopurs_runtime.Value
								}{c_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(a_0, "fst"), "snd")}
								_ = orig
								return gopurs_runtime.RecordDict2("fst", "snd", orig.fst, orig.snd)
							}(), gopurs_runtime.RecordGet(a_0, "snd")}
							_ = orig
							return gopurs_runtime.RecordDict2("fst", "snd", orig.fst, orig.snd)
						}()
						_ = orig
						clone := struct {
							fst struct {
								fst gopurs_runtime.Value
								snd gopurs_runtime.Value
							}
							snd gopurs_runtime.Value
						}{}
						clone.fst = func() struct {
							fst gopurs_runtime.Value
							snd gopurs_runtime.Value
						} {
							orig := gopurs_runtime.RecordGet(orig, "fst")
							_ = orig
							clone := struct {
								fst gopurs_runtime.Value
								snd gopurs_runtime.Value
							}{}
							clone.fst = gopurs_runtime.RecordGet(orig, "fst")
							clone.snd = gopurs_runtime.RecordGet(orig, "snd")
							return clone
						}()
						clone.snd = gopurs_runtime.RecordGet(orig, "snd")
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict2("fst", "snd", func() gopurs_runtime.Value {
						orig := orig.fst
						_ = orig
						return gopurs_runtime.RecordDict2("fst", "snd", orig.fst, orig.snd)
					}(), orig.snd)
				}()
			})}
			_ = orig
			return gopurs_runtime.RecordDict2("get", "set", orig.get, orig.set)
		}()
	})
	return cache_Main_test1
}

func Call_Main_N(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_composeLenses(l1_0_loop struct {
	get gopurs_runtime.Value
	set gopurs_runtime.Value
}, l2_1_loop struct {
	get gopurs_runtime.Value
	set gopurs_runtime.Value
}) struct {
	get gopurs_runtime.Value
	set gopurs_runtime.Value
} {
	var l1_0 struct {
		get gopurs_runtime.Value
		set gopurs_runtime.Value
	} = l1_0_loop
	_ = l1_0
	var l2_1 struct {
		get gopurs_runtime.Value
		set gopurs_runtime.Value
	} = l2_1_loop
	_ = l2_1
	return struct {
		get gopurs_runtime.Value
		set gopurs_runtime.Value
	}{gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(l2_1.get, gopurs_runtime.Apply(l1_0.get, a_2))
	}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, c_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(l1_0.set, a_2, gopurs_runtime.Apply2(l2_1.set, gopurs_runtime.Apply(l1_0.get, a_2), c_3))
	})}
}
