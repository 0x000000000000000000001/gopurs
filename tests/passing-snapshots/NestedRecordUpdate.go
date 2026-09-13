package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_go__init gopurs_runtime.Value
var once_Main_go__init sync.Once

func Get_Main_go__init() gopurs_runtime.Value {
	once_Main_go__init.Do(func() {
		cache_Main_go__init = func() gopurs_runtime.Value {
			orig := struct {
				bar struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}
				foo int64
			}{struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			}{int64(2), struct {
				lhs int64
				rhs int64
			}{int64(3), int64(4)}}, int64(1)}
			_ = orig
			return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
				orig := orig.bar
				_ = orig
				return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), func() gopurs_runtime.Value {
					orig := orig.qux
					_ = orig
					return gopurs_runtime.RecordDict2("lhs", "rhs", gopurs_runtime.Int(orig.lhs), gopurs_runtime.Int(orig.rhs))
				}())
			}(), gopurs_runtime.Int(orig.foo))
		}()
	})
	return cache_Main_go__init
}

var cache_Main_updated gopurs_runtime.Value
var once_Main_updated sync.Once

func Get_Main_updated() gopurs_runtime.Value {
	once_Main_updated.Do(func() {
		cache_Main_updated = func() gopurs_runtime.Value {
			orig := func() struct {
				bar struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}
				foo int64
			} {
				clone := func() struct {
					bar struct {
						baz int64
						qux struct {
							lhs int64
							rhs int64
						}
					}
					foo int64
				} {
					orig := Get_Main_go__init()
					_ = orig
					clone := struct {
						bar struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						}
						foo int64
					}{}
					clone.bar = func() struct {
						baz int64
						qux struct {
							lhs int64
							rhs int64
						}
					} {
						orig := gopurs_runtime.RecordGet(orig, "bar")
						_ = orig
						clone := struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						}{}
						clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
						clone.qux = func() struct {
							lhs int64
							rhs int64
						} {
							orig := gopurs_runtime.RecordGet(orig, "qux")
							_ = orig
							clone := struct {
								lhs int64
								rhs int64
							}{}
							clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
							clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
							return clone
						}()
						return clone
					}()
					clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
					return clone
				}()
				clone.foo = int64(10)
				clone.bar = func() struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				} {
					clone := func() struct {
						bar struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						}
						foo int64
					} {
						orig := Get_Main_go__init()
						_ = orig
						clone := struct {
							bar struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							}
							foo int64
						}{}
						clone.bar = func() struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						} {
							orig := gopurs_runtime.RecordGet(orig, "bar")
							_ = orig
							clone := struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							}{}
							clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
							clone.qux = func() struct {
								lhs int64
								rhs int64
							} {
								orig := gopurs_runtime.RecordGet(orig, "qux")
								_ = orig
								clone := struct {
									lhs int64
									rhs int64
								}{}
								clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
								clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
								return clone
							}()
							return clone
						}()
						clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
						return clone
					}().bar
					clone.baz = int64(20)
					clone.qux = func() struct {
						lhs int64
						rhs int64
					} {
						clone := func() struct {
							bar struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							}
							foo int64
						} {
							orig := Get_Main_go__init()
							_ = orig
							clone := struct {
								bar struct {
									baz int64
									qux struct {
										lhs int64
										rhs int64
									}
								}
								foo int64
							}{}
							clone.bar = func() struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							} {
								orig := gopurs_runtime.RecordGet(orig, "bar")
								_ = orig
								clone := struct {
									baz int64
									qux struct {
										lhs int64
										rhs int64
									}
								}{}
								clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
								clone.qux = func() struct {
									lhs int64
									rhs int64
								} {
									orig := gopurs_runtime.RecordGet(orig, "qux")
									_ = orig
									clone := struct {
										lhs int64
										rhs int64
									}{}
									clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
									clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
									return clone
								}()
								return clone
							}()
							clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
							return clone
						}().bar.qux
						clone.lhs = int64(30)
						clone.rhs = int64(40)
						return clone
					}()
					return clone
				}()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
				orig := orig.bar
				_ = orig
				return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), func() gopurs_runtime.Value {
					orig := orig.qux
					_ = orig
					return gopurs_runtime.RecordDict2("lhs", "rhs", gopurs_runtime.Int(orig.lhs), gopurs_runtime.Int(orig.rhs))
				}())
			}(), gopurs_runtime.Int(orig.foo))
		}()
	})
	return cache_Main_updated
}

var cache_Main_expected gopurs_runtime.Value
var once_Main_expected sync.Once

func Get_Main_expected() gopurs_runtime.Value {
	once_Main_expected.Do(func() {
		cache_Main_expected = func() gopurs_runtime.Value {
			orig := struct {
				bar struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}
				foo int64
			}{struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			}{int64(20), struct {
				lhs int64
				rhs int64
			}{int64(30), int64(40)}}, int64(10)}
			_ = orig
			return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
				orig := orig.bar
				_ = orig
				return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), func() gopurs_runtime.Value {
					orig := orig.qux
					_ = orig
					return gopurs_runtime.RecordDict2("lhs", "rhs", gopurs_runtime.Int(orig.lhs), gopurs_runtime.Int(orig.rhs))
				}())
			}(), gopurs_runtime.Int(orig.foo))
		}()
	})
	return cache_Main_expected
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func6(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value, dictEq2_2_box gopurs_runtime.Value, dictEq3_3_box gopurs_runtime.Value, l_4_box gopurs_runtime.Value, r_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_check(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq2_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq3_3_box), l_4_box, r_5_box))
		})
	})
	return cache_Main_check
}

var cache_Main_check__3900443344 gopurs_runtime.Value
var once_Main_check__3900443344 sync.Once

func Get_Main_check__3900443344() gopurs_runtime.Value {
	once_Main_check__3900443344.Do(func() {
		cache_Main_check__3900443344 = gopurs_runtime.Func2(func(l_unused_0_box gopurs_runtime.Value, r_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_check__3900443344(l_unused_0_box, r_unused_1_box))
		})
	})
	return cache_Main_check__3900443344
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
				if Call_Main_check__3900443344(func() gopurs_runtime.Value {
					orig := func() struct {
						bar struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						}
						foo int64
					} {
						orig := Get_Main_updated()
						_ = orig
						clone := struct {
							bar struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							}
							foo int64
						}{}
						clone.bar = func() struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						} {
							orig := gopurs_runtime.RecordGet(orig, "bar")
							_ = orig
							clone := struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							}{}
							clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
							clone.qux = func() struct {
								lhs int64
								rhs int64
							} {
								orig := gopurs_runtime.RecordGet(orig, "qux")
								_ = orig
								clone := struct {
									lhs int64
									rhs int64
								}{}
								clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
								clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
								return clone
							}()
							return clone
						}()
						clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
						orig := orig.bar
						_ = orig
						return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), func() gopurs_runtime.Value {
							orig := orig.qux
							_ = orig
							return gopurs_runtime.RecordDict2("lhs", "rhs", gopurs_runtime.Int(orig.lhs), gopurs_runtime.Int(orig.rhs))
						}())
					}(), gopurs_runtime.Int(orig.foo))
				}(), func() gopurs_runtime.Value {
					orig := func() struct {
						bar struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						}
						foo int64
					} {
						orig := Get_Main_expected()
						_ = orig
						clone := struct {
							bar struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							}
							foo int64
						}{}
						clone.bar = func() struct {
							baz int64
							qux struct {
								lhs int64
								rhs int64
							}
						} {
							orig := gopurs_runtime.RecordGet(orig, "bar")
							_ = orig
							clone := struct {
								baz int64
								qux struct {
									lhs int64
									rhs int64
								}
							}{}
							clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
							clone.qux = func() struct {
								lhs int64
								rhs int64
							} {
								orig := gopurs_runtime.RecordGet(orig, "qux")
								_ = orig
								clone := struct {
									lhs int64
									rhs int64
								}{}
								clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
								clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
								return clone
							}()
							return clone
						}()
						clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
						return clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict2("bar", "foo", func() gopurs_runtime.Value {
						orig := orig.bar
						_ = orig
						return gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(orig.baz), func() gopurs_runtime.Value {
							orig := orig.qux
							_ = orig
							return gopurs_runtime.RecordDict2("lhs", "rhs", gopurs_runtime.Int(orig.lhs), gopurs_runtime.Int(orig.rhs))
						}())
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

func Call_Main_check(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], dictEq1_1_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], dictEq2_2_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], dictEq3_3_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], l_4_loop gopurs_runtime.Value, r_5_loop gopurs_runtime.Value) bool {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq1_1_loop
	_ = dictEq1_1
	var dictEq2_2 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq2_2_loop
	_ = dictEq2_2
	var dictEq3_3 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq3_3_loop
	_ = dictEq3_3
	var l_4 gopurs_runtime.Value = l_4_loop
	_ = l_4
	var r_5 gopurs_runtime.Value = r_5_loop
	_ = r_5
	return ((gopurs_runtime.Apply2(dictEq_0.V0, gopurs_runtime.RecordGet(l_4, "foo"), gopurs_runtime.RecordGet(r_5, "foo")).IntVal) != (0)) && (((gopurs_runtime.Apply2(dictEq1_1.V0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_4, "bar"), "baz"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_5, "bar"), "baz")).IntVal) != (0)) && (((gopurs_runtime.Apply2(dictEq2_2.V0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_4, "bar"), "qux"), "lhs"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_5, "bar"), "qux"), "lhs")).IntVal) != (0)) && ((gopurs_runtime.Apply2(dictEq3_3.V0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_4, "bar"), "qux"), "rhs"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_5, "bar"), "qux"), "rhs")).IntVal) != (0))))
}

func Call_Main_check__3900443344(l_unused_0_loop gopurs_runtime.Value, r_unused_1_loop gopurs_runtime.Value) bool {
check__3900443344:
	for {
		if false {
			continue check__3900443344
		}
		var l_unused_0 gopurs_runtime.Value = l_unused_0_loop
		_ = l_unused_0
		var r_unused_1 gopurs_runtime.Value = r_unused_1_loop
		_ = r_unused_1
		return ((func() struct {
			bar struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			}
			foo int64
		} {
			orig := Get_Main_updated()
			_ = orig
			clone := struct {
				bar struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}
				foo int64
			}{}
			clone.bar = func() struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			} {
				orig := gopurs_runtime.RecordGet(orig, "bar")
				_ = orig
				clone := struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
				clone.qux = func() struct {
					lhs int64
					rhs int64
				} {
					orig := gopurs_runtime.RecordGet(orig, "qux")
					_ = orig
					clone := struct {
						lhs int64
						rhs int64
					}{}
					clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
					clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
					return clone
				}()
				return clone
			}()
			clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
			return clone
		}().foo) == (int64(10))) && (((func() struct {
			bar struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			}
			foo int64
		} {
			orig := Get_Main_updated()
			_ = orig
			clone := struct {
				bar struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}
				foo int64
			}{}
			clone.bar = func() struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			} {
				orig := gopurs_runtime.RecordGet(orig, "bar")
				_ = orig
				clone := struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
				clone.qux = func() struct {
					lhs int64
					rhs int64
				} {
					orig := gopurs_runtime.RecordGet(orig, "qux")
					_ = orig
					clone := struct {
						lhs int64
						rhs int64
					}{}
					clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
					clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
					return clone
				}()
				return clone
			}()
			clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
			return clone
		}().bar.baz) == (int64(20))) && (((func() struct {
			bar struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			}
			foo int64
		} {
			orig := Get_Main_updated()
			_ = orig
			clone := struct {
				bar struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}
				foo int64
			}{}
			clone.bar = func() struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			} {
				orig := gopurs_runtime.RecordGet(orig, "bar")
				_ = orig
				clone := struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
				clone.qux = func() struct {
					lhs int64
					rhs int64
				} {
					orig := gopurs_runtime.RecordGet(orig, "qux")
					_ = orig
					clone := struct {
						lhs int64
						rhs int64
					}{}
					clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
					clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
					return clone
				}()
				return clone
			}()
			clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
			return clone
		}().bar.qux.lhs) == (int64(30))) && ((func() struct {
			bar struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			}
			foo int64
		} {
			orig := Get_Main_updated()
			_ = orig
			clone := struct {
				bar struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}
				foo int64
			}{}
			clone.bar = func() struct {
				baz int64
				qux struct {
					lhs int64
					rhs int64
				}
			} {
				orig := gopurs_runtime.RecordGet(orig, "bar")
				_ = orig
				clone := struct {
					baz int64
					qux struct {
						lhs int64
						rhs int64
					}
				}{}
				clone.baz = gopurs_runtime.RecordGet(orig, "baz").IntVal
				clone.qux = func() struct {
					lhs int64
					rhs int64
				} {
					orig := gopurs_runtime.RecordGet(orig, "qux")
					_ = orig
					clone := struct {
						lhs int64
						rhs int64
					}{}
					clone.lhs = gopurs_runtime.RecordGet(orig, "lhs").IntVal
					clone.rhs = gopurs_runtime.RecordGet(orig, "rhs").IntVal
					return clone
				}()
				return clone
			}()
			clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
			return clone
		}().bar.qux.rhs) == (int64(40)))))
	}
}
