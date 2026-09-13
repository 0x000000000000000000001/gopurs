package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_update gopurs_runtime.Value
var once_Main_update sync.Once

func Get_Main_update() gopurs_runtime.Value {
	once_Main_update.Do(func() {
		cache_Main_update = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_update(v_0_box, v1_1_box, v2_2_box, v3_3_box)
		})
	})
	return cache_Main_update
}

var cache_Main_go__init gopurs_runtime.Value
var once_Main_go__init sync.Once

func Get_Main_go__init() gopurs_runtime.Value {
	once_Main_go__init.Do(func() {
		cache_Main_go__init = func() gopurs_runtime.Value {
			orig := struct {
				bar struct {
					baz int64
					qux int64
				}
				foo int64
			}{struct {
				baz int64
				qux int64
			}{int64(2), int64(3)}, int64(1)}
			_ = orig
			return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
				orig := orig.bar
				_ = orig
				return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), gopurs_runtime.Int(orig.qux))
			}(), gopurs_runtime.Int(orig.foo))
		}()
	})
	return cache_Main_go__init
}

var cache_Main_expected gopurs_runtime.Value
var once_Main_expected sync.Once

func Get_Main_expected() gopurs_runtime.Value {
	once_Main_expected.Do(func() {
		cache_Main_expected = func() gopurs_runtime.Value {
			orig := struct {
				bar struct {
					baz int64
					qux int64
				}
				foo int64
			}{struct {
				baz int64
				qux int64
			}{int64(20), int64(30)}, int64(10)}
			_ = orig
			return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
				orig := orig.bar
				_ = orig
				return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), gopurs_runtime.Int(orig.qux))
			}(), gopurs_runtime.Int(orig.foo))
		}()
	})
	return cache_Main_expected
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func5(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value, dictEq2_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, r_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_check(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq2_2_box), l_3_box, r_4_box))
		})
	})
	return cache_Main_check
}

var cache_Main_after gopurs_runtime.Value
var once_Main_after sync.Once

func Get_Main_after() gopurs_runtime.Value {
	once_Main_after.Do(func() {
		cache_Main_after = func() gopurs_runtime.Value {
			orig := func() struct {
				bar struct {
					baz int64
					qux int64
				}
				foo int64
			} {
				orig := Call_Main_update(func() gopurs_runtime.Value {
					orig := func() struct {
						bar struct {
							baz int64
							qux int64
						}
						foo int64
					} {
						orig := Get_Main_go__init()
						_ = orig
						clone := struct {
							bar struct {
								baz int64
								qux int64
							}
							foo int64
						}{}
						clone.bar = func() struct {
							baz int64
							qux int64
						} {
							orig := gopurs_runtime.RecordGet(orig, "bar")
							_ = orig
							clone := struct {
								baz int64
								qux int64
							}{}
							clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
							clone.qux = gopurs_runtime.RecordGet(orig, "qux").IntVal
							return clone
						}()
						clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
						orig := orig.bar
						_ = orig
						return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), gopurs_runtime.Int(orig.qux))
					}(), gopurs_runtime.Int(orig.foo))
				}(), gopurs_runtime.Int(int64(10)), gopurs_runtime.Int(int64(20)), gopurs_runtime.Int(int64(30)))
				_ = orig
				clone := struct {
					bar struct {
						baz int64
						qux int64
					}
					foo int64
				}{}
				clone.bar = func() struct {
					baz int64
					qux int64
				} {
					orig := gopurs_runtime.RecordGet(orig, "bar")
					_ = orig
					clone := struct {
						baz int64
						qux int64
					}{}
					clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
					clone.qux = gopurs_runtime.RecordGet(orig, "qux").IntVal
					return clone
				}()
				clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
				orig := orig.bar
				_ = orig
				return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), gopurs_runtime.Int(orig.qux))
			}(), gopurs_runtime.Int(orig.foo))
		}()
	})
	return cache_Main_after
}

var cache_Main_check__1055591803 gopurs_runtime.Value
var once_Main_check__1055591803 sync.Once

func Get_Main_check__1055591803() gopurs_runtime.Value {
	once_Main_check__1055591803.Do(func() {
		cache_Main_check__1055591803 = gopurs_runtime.Func2(func(l_unused_0_box gopurs_runtime.Value, r_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_check__1055591803(l_unused_0_box, r_unused_1_box))
		})
	})
	return cache_Main_check__1055591803
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			_ = __local_var_0_0
			var __t1 gopurs_runtime.Value
			{
				if Call_Main_check__1055591803(func() gopurs_runtime.Value {
					orig := func() struct {
						bar struct {
							baz int64
							qux int64
						}
						foo int64
					} {
						orig := Get_Main_after()
						_ = orig
						clone := struct {
							bar struct {
								baz int64
								qux int64
							}
							foo int64
						}{}
						clone.bar = func() struct {
							baz int64
							qux int64
						} {
							orig := gopurs_runtime.RecordGet(orig, "bar")
							_ = orig
							clone := struct {
								baz int64
								qux int64
							}{}
							clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
							clone.qux = gopurs_runtime.RecordGet(orig, "qux").IntVal
							return clone
						}()
						clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
						orig := orig.bar
						_ = orig
						return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), gopurs_runtime.Int(orig.qux))
					}(), gopurs_runtime.Int(orig.foo))
				}(), func() gopurs_runtime.Value {
					orig := func() struct {
						bar struct {
							baz int64
							qux int64
						}
						foo int64
					} {
						orig := Get_Main_expected()
						_ = orig
						clone := struct {
							bar struct {
								baz int64
								qux int64
							}
							foo int64
						}{}
						clone.bar = func() struct {
							baz int64
							qux int64
						} {
							orig := gopurs_runtime.RecordGet(orig, "bar")
							_ = orig
							clone := struct {
								baz int64
								qux int64
							}{}
							clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
							clone.qux = gopurs_runtime.RecordGet(orig, "qux").IntVal
							return clone
						}()
						clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
						orig := orig.bar
						_ = orig
						return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), gopurs_runtime.Int(orig.qux))
					}(), gopurs_runtime.Int(orig.foo))
				}()) {
					__t1 = __local_var_0_0
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
			}
		end_branch_1:
			return __t1
		}()
	})
	return cache_Main_main
}

func Call_Main_update(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value, v3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var v3_3 gopurs_runtime.Value = v3_3_loop
	_ = v3_3
	return gopurs_runtime.RecordUpdate2(v_0, "foo", v1_1, "bar", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(v_0, "bar"), "baz", v2_2, "qux", v3_3))
}

func Call_Main_check(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], dictEq1_1_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], dictEq2_2_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], l_3_loop gopurs_runtime.Value, r_4_loop gopurs_runtime.Value) bool {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq1_1_loop
	_ = dictEq1_1
	var dictEq2_2 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq2_2_loop
	_ = dictEq2_2
	var l_3 gopurs_runtime.Value = l_3_loop
	_ = l_3
	var r_4 gopurs_runtime.Value = r_4_loop
	_ = r_4
	return ((gopurs_runtime.Apply2(dictEq_0.V0, gopurs_runtime.RecordGet(l_3, "foo"), gopurs_runtime.RecordGet(r_4, "foo")).IntVal) != (0)) && (((gopurs_runtime.Apply2(dictEq1_1.V0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_3, "bar"), "baz"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_4, "bar"), "baz")).IntVal) != (0)) && ((gopurs_runtime.Apply2(dictEq2_2.V0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_3, "bar"), "qux"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_4, "bar"), "qux")).IntVal) != (0)))
}

func Call_Main_check__1055591803(l_unused_0_loop gopurs_runtime.Value, r_unused_1_loop gopurs_runtime.Value) bool {
check__1055591803:
	for {
		if false {
			continue check__1055591803
		}
		var l_unused_0 gopurs_runtime.Value = l_unused_0_loop
		_ = l_unused_0
		var r_unused_1 gopurs_runtime.Value = r_unused_1_loop
		_ = r_unused_1
		return ((func() struct {
			bar struct {
				baz int64
				qux int64
			}
			foo int64
		} {
			orig := Get_Main_after()
			_ = orig
			clone := struct {
				bar struct {
					baz int64
					qux int64
				}
				foo int64
			}{}
			clone.bar = func() struct {
				baz int64
				qux int64
			} {
				orig := gopurs_runtime.RecordGet(orig, "bar")
				_ = orig
				clone := struct {
					baz int64
					qux int64
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
				clone.qux = gopurs_runtime.RecordGet(orig, "qux").IntVal
				return clone
			}()
			clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
			return clone
		}().foo) == (int64(10))) && (((func() struct {
			bar struct {
				baz int64
				qux int64
			}
			foo int64
		} {
			orig := Get_Main_after()
			_ = orig
			clone := struct {
				bar struct {
					baz int64
					qux int64
				}
				foo int64
			}{}
			clone.bar = func() struct {
				baz int64
				qux int64
			} {
				orig := gopurs_runtime.RecordGet(orig, "bar")
				_ = orig
				clone := struct {
					baz int64
					qux int64
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
				clone.qux = gopurs_runtime.RecordGet(orig, "qux").IntVal
				return clone
			}()
			clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
			return clone
		}().bar.baz) == (int64(20))) && ((func() struct {
			bar struct {
				baz int64
				qux int64
			}
			foo int64
		} {
			orig := Get_Main_after()
			_ = orig
			clone := struct {
				bar struct {
					baz int64
					qux int64
				}
				foo int64
			}{}
			clone.bar = func() struct {
				baz int64
				qux int64
			} {
				orig := gopurs_runtime.RecordGet(orig, "bar")
				_ = orig
				clone := struct {
					baz int64
					qux int64
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
				clone.qux = gopurs_runtime.RecordGet(orig, "qux").IntVal
				return clone
			}()
			clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
			return clone
		}().bar.qux) == (int64(30))))
	}
}
