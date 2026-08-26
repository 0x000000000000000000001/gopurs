package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_fn gopurs_runtime.Value
var once_Main_fn sync.Once

func Get_Main_fn() gopurs_runtime.Value {
	once_Main_fn.Do(func() {
		cache_Main_fn = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fn(func() *struct {
				b *struct {
					c *struct {
						d gopurs_runtime.Value
					}
				}
			} {
				orig := v_0_box
				_ = orig
				clone := struct {
					b *struct {
						c *struct {
							d gopurs_runtime.Value
						}
					}
				}{}
				clone.b = func() *struct {
					c *struct {
						d gopurs_runtime.Value
					}
				} {
					orig := gopurs_runtime.RecordGet(orig, "b")
					_ = orig
					clone := struct {
						c *struct {
							d gopurs_runtime.Value
						}
					}{}
					clone.c = func() *struct {
						d gopurs_runtime.Value
					} {
						orig := gopurs_runtime.RecordGet(orig, "c")
						_ = orig
						clone := struct {
							d gopurs_runtime.Value
						}{}
						clone.d = gopurs_runtime.RecordGet(orig, "d")
						return &clone
					}()
					return &clone
				}()
				return &clone
			}())
		})
	})
	return cache_Main_fn
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = func() gopurs_runtime.Value {
			orig := func() *struct {
				b *struct {
					c *struct {
						d int64
					}
				}
			} {
				orig := gopurs_runtime.RecordDict1("b", func() gopurs_runtime.Value {
					orig := func() *struct {
						c *struct {
							d int64
						}
					} {
						orig := gopurs_runtime.RecordDict1("c", func() gopurs_runtime.Value {
							orig := func() *struct {
								d int64
							} {
								orig := gopurs_runtime.RecordDict1("d", gopurs_runtime.Int(2))
								_ = orig
								clone := struct {
									d int64
								}{}
								clone.d = gopurs_runtime.RecordGet(orig, "d").IntVal
								return &clone
							}()
							_ = orig
							return gopurs_runtime.RecordDict([]string{"d"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.d)})
						}())
						_ = orig
						clone := struct {
							c *struct {
								d int64
							}
						}{}
						clone.c = func() *struct {
							d int64
						} {
							orig := gopurs_runtime.RecordGet(orig, "c")
							_ = orig
							clone := struct {
								d int64
							}{}
							clone.d = gopurs_runtime.RecordGet(orig, "d").IntVal
							return &clone
						}()
						return &clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict([]string{"c"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
						orig := orig.c
						_ = orig
						return gopurs_runtime.RecordDict([]string{"d"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.d)})
					}()})
				}())
				_ = orig
				clone := struct {
					b *struct {
						c *struct {
							d int64
						}
					}
				}{}
				clone.b = func() *struct {
					c *struct {
						d int64
					}
				} {
					orig := gopurs_runtime.RecordGet(orig, "b")
					_ = orig
					clone := struct {
						c *struct {
							d int64
						}
					}{}
					clone.c = func() *struct {
						d int64
					} {
						orig := gopurs_runtime.RecordGet(orig, "c")
						_ = orig
						clone := struct {
							d int64
						}{}
						clone.d = gopurs_runtime.RecordGet(orig, "d").IntVal
						return &clone
					}()
					return &clone
				}()
				return &clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"b"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
				orig := orig.b
				_ = orig
				return gopurs_runtime.RecordDict([]string{"c"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
					orig := orig.c
					_ = orig
					return gopurs_runtime.RecordDict([]string{"d"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.d)})
				}()})
			}()})
		}()
	})
	return cache_Main_a
}

var cache_Main_d gopurs_runtime.Value
var once_Main_d sync.Once

func Get_Main_d() gopurs_runtime.Value {
	once_Main_d.Do(func() {
		cache_Main_d = gopurs_runtime.Int(4)
	})
	return cache_Main_d
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_fn(v_0_loop *struct {
	b *struct {
		c *struct {
			d gopurs_runtime.Value
		}
	}
}) gopurs_runtime.Value {
	var v_0 *struct {
		b *struct {
			c *struct {
				d gopurs_runtime.Value
			}
		}
	} = v_0_loop
	_ = v_0
	return v_0.b.c.d
}
