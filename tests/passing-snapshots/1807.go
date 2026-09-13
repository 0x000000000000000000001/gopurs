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
			return Call_Main_fn(v_0_box)
		})
	})
	return cache_Main_fn
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = func() gopurs_runtime.Value {
			orig := struct {
				b struct {
					c struct {
						d int64
					}
				}
			}{struct {
				c struct {
					d int64
				}
			}{struct {
				d int64
			}{int64(2)}}}
			_ = orig
			return gopurs_runtime.RecordDict1("b", func() gopurs_runtime.Value {
				orig := orig.b
				_ = orig
				return gopurs_runtime.RecordDict1("c", func() gopurs_runtime.Value {
					orig := orig.c
					_ = orig
					return gopurs_runtime.RecordDict1("d", gopurs_runtime.Int(orig.d))
				}())
			}())
		}()
	})
	return cache_Main_a
}

var cache_Main_d gopurs_runtime.Value
var once_Main_d sync.Once

func Get_Main_d() gopurs_runtime.Value {
	once_Main_d.Do(func() {
		cache_Main_d = gopurs_runtime.Int((Call_Main_fn(func() gopurs_runtime.Value {
			orig := func() struct {
				b struct {
					c struct {
						d int64
					}
				}
			} {
				orig := Get_Main_a()
				_ = orig
				clone := struct {
					b struct {
						c struct {
							d int64
						}
					}
				}{}
				clone.b = func() struct {
					c struct {
						d int64
					}
				} {
					orig := gopurs_runtime.RecordGet(orig, "b")
					_ = orig
					clone := struct {
						c struct {
							d int64
						}
					}{}
					clone.c = func() struct {
						d int64
					} {
						orig := gopurs_runtime.RecordGet(orig, "c")
						_ = orig
						clone := struct {
							d int64
						}{}
						clone.d = gopurs_runtime.RecordGet(orig, "d").IntVal
						return clone
					}()
					return clone
				}()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict1("b", func() gopurs_runtime.Value {
				orig := orig.b
				_ = orig
				return gopurs_runtime.RecordDict1("c", func() gopurs_runtime.Value {
					orig := orig.c
					_ = orig
					return gopurs_runtime.RecordDict1("d", gopurs_runtime.Int(orig.d))
				}())
			}())
		}()).IntVal) + (int64(2)))
	})
	return cache_Main_d
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if ((Call_Main_fn(func() gopurs_runtime.Value {
					orig := func() struct {
						b struct {
							c struct {
								d int64
							}
						}
					} {
						orig := Get_Main_a()
						_ = orig
						clone := struct {
							b struct {
								c struct {
									d int64
								}
							}
						}{}
						clone.b = func() struct {
							c struct {
								d int64
							}
						} {
							orig := gopurs_runtime.RecordGet(orig, "b")
							_ = orig
							clone := struct {
								c struct {
									d int64
								}
							}{}
							clone.c = func() struct {
								d int64
							} {
								orig := gopurs_runtime.RecordGet(orig, "c")
								_ = orig
								clone := struct {
									d int64
								}{}
								clone.d = gopurs_runtime.RecordGet(orig, "d").IntVal
								return clone
							}()
							return clone
						}()
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict1("b", func() gopurs_runtime.Value {
						orig := orig.b
						_ = orig
						return gopurs_runtime.RecordDict1("c", func() gopurs_runtime.Value {
							orig := orig.c
							_ = orig
							return gopurs_runtime.RecordDict1("d", gopurs_runtime.Int(orig.d))
						}())
					}())
				}()).IntVal) + (int64(2))) == (int64(4)) {
					__t0 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail"))
			}
		end_branch_0:
			return __t0
		}()
	})
	return cache_Main_main
}

func Call_Main_fn(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "b"), "c"), "d")
}
