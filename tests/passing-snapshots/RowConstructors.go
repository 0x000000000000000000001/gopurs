package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_wildcard_prime_ gopurs_runtime.Value
var once_Main_wildcard_prime_ sync.Once

func Get_Main_wildcard_prime_() gopurs_runtime.Value {
	once_Main_wildcard_prime_.Do(func() {
		cache_Main_wildcard_prime_ = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_wildcard_prime_(v_0_box))
		})
	})
	return cache_Main_wildcard_prime_
}

var cache_Main_wildcard gopurs_runtime.Value
var once_Main_wildcard sync.Once

func Get_Main_wildcard() gopurs_runtime.Value {
	once_Main_wildcard.Do(func() {
		cache_Main_wildcard = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_wildcard(v_0_box)
				_ = orig
				return gopurs_runtime.RecordDict4("w", "x", "y", "z", gopurs_runtime.Float(orig.w), gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
			}()
		})
	})
	return cache_Main_wildcard
}

var cache_Main_quux_prime_ gopurs_runtime.Value
var once_Main_quux_prime_ sync.Once

func Get_Main_quux_prime_() gopurs_runtime.Value {
	once_Main_quux_prime_.Do(func() {
		cache_Main_quux_prime_ = func() gopurs_runtime.Value {
			orig := struct {
				q        float64
				q_prime_ float64
				x        float64
				y        float64
				z        float64
			}{0.0, 0.0, 0.0, 0.0, 0.0}
			_ = orig
			return gopurs_runtime.RecordDict5("q", "q'", "x", "y", "z", gopurs_runtime.Float(orig.q), gopurs_runtime.Float(orig.q_prime_), gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
		}()
	})
	return cache_Main_quux_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_id_prime_ gopurs_runtime.Value
var once_Main_id_prime_ sync.Once

func Get_Main_id_prime_() gopurs_runtime.Value {
	once_Main_id_prime_.Do(func() {
		cache_Main_id_prime_ = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_id_prime_
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = func() gopurs_runtime.Value {
			orig := struct {
				x float64
				y float64
				z float64
			}{0.0, 0.0, 0.0}
			_ = orig
			return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
		}()
	})
	return cache_Main_foo
}

var cache_Main_foo_prime_ gopurs_runtime.Value
var once_Main_foo_prime_ sync.Once

func Get_Main_foo_prime_() gopurs_runtime.Value {
	once_Main_foo_prime_.Do(func() {
		cache_Main_foo_prime_ = func() gopurs_runtime.Value {
			orig := func() struct {
				x float64
				y float64
				z float64
			} {
				orig := gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), func() gopurs_runtime.Value {
					orig := func() struct {
						x float64
						y float64
						z float64
					} {
						orig := Get_Main_foo()
						_ = orig
						clone := struct {
							x float64
							y float64
							z float64
						}{}
						clone.x = gopurs_runtime.RecordGet(orig, "x").FloatVal()
						clone.y = gopurs_runtime.RecordGet(orig, "y").FloatVal()
						clone.z = gopurs_runtime.RecordGet(orig, "z").FloatVal()
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
				}())
				_ = orig
				clone := struct {
					x float64
					y float64
					z float64
				}{}
				clone.x = gopurs_runtime.RecordGet(orig, "x").FloatVal()
				clone.y = gopurs_runtime.RecordGet(orig, "y").FloatVal()
				clone.z = gopurs_runtime.RecordGet(orig, "z").FloatVal()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
		}()
	})
	return cache_Main_foo_prime_
}

var cache_Main_quux gopurs_runtime.Value
var once_Main_quux sync.Once

func Get_Main_quux() gopurs_runtime.Value {
	once_Main_quux.Do(func() {
		cache_Main_quux = func() gopurs_runtime.Value {
			orig := struct {
				f struct {
					x float64
					y float64
					z float64
				}
				q float64
				x float64
				y float64
				z float64
			}{func() struct {
				x float64
				y float64
				z float64
			} {
				orig := Get_Main_foo_prime_()
				_ = orig
				clone := struct {
					x float64
					y float64
					z float64
				}{}
				clone.x = gopurs_runtime.RecordGet(orig, "x").FloatVal()
				clone.y = gopurs_runtime.RecordGet(orig, "y").FloatVal()
				clone.z = gopurs_runtime.RecordGet(orig, "z").FloatVal()
				return clone
			}(), 0.0, 0.0, 0.0, 0.0}
			_ = orig
			return gopurs_runtime.RecordDict5("f", "q", "x", "y", "z", func() gopurs_runtime.Value {
				orig := orig.f
				_ = orig
				return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
			}(), gopurs_runtime.Float(orig.q), gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
		}()
	})
	return cache_Main_quux
}

var cache_Main_baz gopurs_runtime.Value
var once_Main_baz sync.Once

func Get_Main_baz() gopurs_runtime.Value {
	once_Main_baz.Do(func() {
		cache_Main_baz = func() gopurs_runtime.Value {
			orig := struct {
				w float64
				x float64
				y float64
				z float64
			}{0.0, 0.0, 0.0, 0.0}
			_ = orig
			return gopurs_runtime.RecordDict4("w", "x", "y", "z", gopurs_runtime.Float(orig.w), gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
		}()
	})
	return cache_Main_baz
}

var cache_Main_bar gopurs_runtime.Value
var once_Main_bar sync.Once

func Get_Main_bar() gopurs_runtime.Value {
	once_Main_bar.Do(func() {
		cache_Main_bar = func() gopurs_runtime.Value {
			orig := struct {
				x float64
				y float64
				z float64
			}{0.0, 0.0, 0.0}
			_ = orig
			return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
		}()
	})
	return cache_Main_bar
}

var cache_Main_bar_prime_ gopurs_runtime.Value
var once_Main_bar_prime_ sync.Once

func Get_Main_bar_prime_() gopurs_runtime.Value {
	once_Main_bar_prime_.Do(func() {
		cache_Main_bar_prime_ = func() gopurs_runtime.Value {
			orig := func() struct {
				x float64
				y float64
				z float64
			} {
				orig := gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), func() gopurs_runtime.Value {
					orig := func() struct {
						x float64
						y float64
						z float64
					} {
						orig := Get_Main_bar()
						_ = orig
						clone := struct {
							x float64
							y float64
							z float64
						}{}
						clone.x = gopurs_runtime.RecordGet(orig, "x").FloatVal()
						clone.y = gopurs_runtime.RecordGet(orig, "y").FloatVal()
						clone.z = gopurs_runtime.RecordGet(orig, "z").FloatVal()
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
				}())
				_ = orig
				clone := struct {
					x float64
					y float64
					z float64
				}{}
				clone.x = gopurs_runtime.RecordGet(orig, "x").FloatVal()
				clone.y = gopurs_runtime.RecordGet(orig, "y").FloatVal()
				clone.z = gopurs_runtime.RecordGet(orig, "z").FloatVal()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y), gopurs_runtime.Float(orig.z))
		}()
	})
	return cache_Main_bar_prime_
}

func Call_Main_wildcard_prime_(v_0_loop gopurs_runtime.Value) float64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.RecordGet(v_0, "q").FloatVal()
}

func Call_Main_wildcard(v_0_loop gopurs_runtime.Value) struct {
	w float64
	x float64
	y float64
	z float64
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return struct {
		w float64
		x float64
		y float64
		z float64
	}{gopurs_runtime.RecordGet(v_0, "w").FloatVal(), gopurs_runtime.RecordGet(v_0, "w").FloatVal(), gopurs_runtime.RecordGet(v_0, "w").FloatVal(), gopurs_runtime.RecordGet(v_0, "w").FloatVal()}
}
