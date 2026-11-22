package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ord_eqRec gopurs_runtime.Value
var once_Data_Ord_eqRec sync.Once

func Get_Data_Ord_eqRec() gopurs_runtime.Value {
	once_Data_Ord_eqRec.Do(func() {
		cache_Data_Ord_eqRec = gopurs_runtime.Apply(Get_Data_Eq_eqRec(), gopurs_runtime.Value{})
	})
	return cache_Data_Ord_eqRec
}

var cache_Data_Ord_negate gopurs_runtime.Value
var once_Data_Ord_negate sync.Once

func Get_Data_Ord_negate() gopurs_runtime.Value {
	once_Data_Ord_negate.Do(func() {
		cache_Data_Ord_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Ord_negate(a_0_box.IntVal))
		})
	})
	return cache_Data_Ord_negate
}

var cache_Data_Ord_notEq gopurs_runtime.Value
var once_Data_Ord_notEq sync.Once

func Get_Data_Ord_notEq() gopurs_runtime.Value {
	once_Data_Ord_notEq.Do(func() {
		cache_Data_Ord_notEq = Get_Data_Eq_notEq__2123013075()
	})
	return cache_Data_Ord_notEq
}

var cache_Data_Ord_OrdRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Ord_OrdRecord_dollar_Dict sync.Once

func Get_Data_Ord_OrdRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ord_OrdRecord_dollar_Dict.Do(func() {
		cache_Data_Ord_OrdRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(Call_Data_Ord_OrdRecord_dollar_Dict(func() struct {
				EqRecord0     gopurs_runtime.Value
				compareRecord gopurs_runtime.Value
				/*gopurs:record-v1:2:00450071005200650063006f007200640030,0063006f006d0070006100720065005200650063006f00720064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					EqRecord0     gopurs_runtime.Value
					compareRecord gopurs_runtime.Value
					/*gopurs:record-v1:2:00450071005200650063006f007200640030,0063006f006d0070006100720065005200650063006f00720064*/
				}{}
				clone.EqRecord0 = gopurs_runtime.RecordGet(orig, "EqRecord0")
				clone.compareRecord = gopurs_runtime.RecordGet(orig, "compareRecord")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_OrdRecord_dollar_Dict
}

var cache_Data_Ord_Ord_dollar_Dict gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict sync.Once

func Get_Data_Ord_Ord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict
}

var cache_Data_Ord_Ord_dollar_Dict__3648059431 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__3648059431 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__3648059431() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__3648059431.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__3648059431 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__3648059431(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__3648059431
}

var cache_Data_Ord_Ord_dollar_Dict__2520386791 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2520386791 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2520386791() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2520386791.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2520386791 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__2520386791(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2520386791
}

var cache_Data_Ord_Ord_dollar_Dict__3832589959 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__3832589959 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__3832589959() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__3832589959.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__3832589959 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__3832589959(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__3832589959
}

var cache_Data_Ord_Ord_dollar_Dict__4094323175 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__4094323175 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__4094323175() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__4094323175.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__4094323175 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1888627657_4177771502(Call_Data_Ord_Ord_dollar_Dict__4094323175(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__4094323175
}

var cache_Data_Ord_Ord_dollar_Dict__757207815 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__757207815 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__757207815() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__757207815.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__757207815 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Call_Data_Ord_Ord_dollar_Dict__757207815(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__757207815
}

var cache_Data_Ord_Ord_dollar_Dict__2432464903 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2432464903 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2432464903() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2432464903.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2432464903 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Call_Data_Ord_Ord_dollar_Dict__2432464903(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2432464903
}

var cache_Data_Ord_Ord_dollar_Dict__3848172583 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__3848172583 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__3848172583() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__3848172583.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__3848172583 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_758368489_4177771502(Call_Data_Ord_Ord_dollar_Dict__3848172583(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__3848172583
}

var cache_Data_Ord_Ord_dollar_Dict__3034920007 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__3034920007 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__3034920007() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__3034920007.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__3034920007 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Call_Data_Ord_Ord_dollar_Dict__3034920007(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__3034920007
}

var cache_Data_Ord_Ord_dollar_Dict__2010913863 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2010913863 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2010913863() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2010913863.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2010913863 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Call_Data_Ord_Ord_dollar_Dict__2010913863(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2010913863
}

var cache_Data_Ord_Ord_dollar_Dict__3994504999 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__3994504999 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__3994504999() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__3994504999.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__3994504999 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__3994504999(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__3994504999
}

var cache_Data_Ord_Ord_dollar_Dict__2123630727 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2123630727 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2123630727() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2123630727.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2123630727 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__2123630727(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2123630727
}

var cache_Data_Ord_Ord_dollar_Dict__1187641767 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__1187641767 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__1187641767() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__1187641767.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__1187641767 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Call_Data_Ord_Ord_dollar_Dict__1187641767(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__1187641767
}

var cache_Data_Ord_Ord_dollar_Dict__1570991783 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__1570991783 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__1570991783() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__1570991783.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__1570991783 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Call_Data_Ord_Ord_dollar_Dict__1570991783(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__1570991783
}

var cache_Data_Ord_Ord_dollar_Dict__2690096935 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2690096935 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2690096935() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2690096935.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2690096935 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3383090825_4177771502(Call_Data_Ord_Ord_dollar_Dict__2690096935(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2690096935
}

var cache_Data_Ord_Ord_dollar_Dict__832340135 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__832340135 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__832340135() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__832340135.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__832340135 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__832340135(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__832340135
}

var cache_Data_Ord_Ord_dollar_Dict__1238669767 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__1238669767 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__1238669767() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__1238669767.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__1238669767 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__1238669767(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__1238669767
}

var cache_Data_Ord_Ord_dollar_Dict__1954751335 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__1954751335 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__1954751335() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__1954751335.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__1954751335 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__1954751335(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__1954751335
}

var cache_Data_Ord_Ord_dollar_Dict__4182661159 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__4182661159 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__4182661159() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__4182661159.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__4182661159 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Call_Data_Ord_Ord_dollar_Dict__4182661159(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__4182661159
}

var cache_Data_Ord_Ord_dollar_Dict__2336892295 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2336892295 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2336892295() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2336892295.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2336892295 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_219188042_4177771502(Call_Data_Ord_Ord_dollar_Dict__2336892295(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2336892295
}

var cache_Data_Ord_Ord_dollar_Dict__4108678119 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__4108678119 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__4108678119() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__4108678119.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__4108678119 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2406510097_4177771502(Call_Data_Ord_Ord_dollar_Dict__4108678119(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__4108678119
}

var cache_Data_Ord_Ord_dollar_Dict__2468619815 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2468619815 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2468619815() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2468619815.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2468619815 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3308271157_4177771502(Call_Data_Ord_Ord_dollar_Dict__2468619815(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2468619815
}

var cache_Data_Ord_Ord_dollar_Dict__3472323847 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__3472323847 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__3472323847() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__3472323847.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__3472323847 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3047586294_4177771502(Call_Data_Ord_Ord_dollar_Dict__3472323847(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__3472323847
}

var cache_Data_Ord_Ord_dollar_Dict__2744908871 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2744908871 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2744908871() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2744908871.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2744908871 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2406510097_4177771502(Call_Data_Ord_Ord_dollar_Dict__2744908871(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}())))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2744908871
}

var cache_Data_Ord_Ord_dollar_Dict__2297783463 gopurs_runtime.Value
var once_Data_Ord_Ord_dollar_Dict__2297783463 sync.Once

func Get_Data_Ord_Ord_dollar_Dict__2297783463() gopurs_runtime.Value {
	once_Data_Ord_Ord_dollar_Dict__2297783463.Do(func() {
		cache_Data_Ord_Ord_dollar_Dict__2297783463 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord_dollar_Dict__2297783463(func() struct {
				Eq0     gopurs_runtime.Value
				compare gopurs_runtime.Value
				/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq0     gopurs_runtime.Value
					compare gopurs_runtime.Value
					/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
				}{}
				clone.Eq0 = gopurs_runtime.RecordGet(orig, "Eq0")
				clone.compare = gopurs_runtime.RecordGet(orig, "compare")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord_dollar_Dict__2297783463
}

var cache_Data_Ord_Ord1_dollar_Dict gopurs_runtime.Value
var once_Data_Ord_Ord1_dollar_Dict sync.Once

func Get_Data_Ord_Ord1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ord_Ord1_dollar_Dict.Do(func() {
		cache_Data_Ord_Ord1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Ord1_dollar_Dict(func() struct {
				Eq10     gopurs_runtime.Value
				compare1 gopurs_runtime.Value
				/*gopurs:record-v1:2:0045007100310030,0063006f006d00700061007200650031*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Eq10     gopurs_runtime.Value
					compare1 gopurs_runtime.Value
					/*gopurs:record-v1:2:0045007100310030,0063006f006d00700061007200650031*/
				}{}
				clone.Eq10 = gopurs_runtime.RecordGet(orig, "Eq10")
				clone.compare1 = gopurs_runtime.RecordGet(orig, "compare1")
				return clone
			}()))}
		})
	})
	return cache_Data_Ord_Ord1_dollar_Dict
}

var cache_Data_Ord_ordVoid gopurs_runtime.Value
var once_Data_Ord_ordVoid sync.Once

func Get_Data_Ord_ordVoid() gopurs_runtime.Value {
	once_Data_Ord_ordVoid.Do(func() {
		cache_Data_Ord_ordVoid = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqVoid()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_Ord_ordVoid
}

var cache_Data_Ord_ordUnit gopurs_runtime.Value
var once_Data_Ord_ordUnit sync.Once

func Get_Data_Ord_ordUnit() gopurs_runtime.Value {
	once_Data_Ord_ordUnit.Do(func() {
		cache_Data_Ord_ordUnit = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqUnit()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_Ord_ordUnit
}

var cache_Data_Ord_ordString gopurs_runtime.Value
var once_Data_Ord_ordString sync.Once

func Get_Data_Ord_ordString() gopurs_runtime.Value {
	once_Data_Ord_ordString.Do(func() {
		cache_Data_Ord_ordString = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1140313009_3790796878(Rebox_Data_Ord_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}
		}), gopurs_runtime.Apply3(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordString
}

var cache_Data_Ord_ordRecordNil gopurs_runtime.Value
var once_Data_Ord_ordRecordNil sync.Once

func Get_Data_Ord_ordRecordNil() gopurs_runtime.Value {
	once_Data_Ord_ordRecordNil.Do(func() {
		cache_Data_Ord_ordRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Eq_eqRowNil()
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_Ord_ordRecordNil
}

var cache_Data_Ord_ordProxy gopurs_runtime.Value
var once_Data_Ord_ordProxy sync.Once

func Get_Data_Ord_ordProxy() gopurs_runtime.Value {
	once_Data_Ord_ordProxy.Do(func() {
		cache_Data_Ord_ordProxy = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Eq_eqProxy()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})})))}
	})
	return cache_Data_Ord_ordProxy
}

var cache_Data_Ord_ordOrdering gopurs_runtime.Value
var once_Data_Ord_ordOrdering sync.Once

func Get_Data_Ord_ordOrdering() gopurs_runtime.Value {
	once_Data_Ord_ordOrdering.Do(func() {
		cache_Data_Ord_ordOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3768443459_3790796878(Rebox_Data_Ord_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Ordering_eqOrdering()))))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t11 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1527465420 {
					var __t2 gopurs_runtime.Value
					{
						var __t_tag_1 uint32 = uint32(v1_1.IntVal)
						_ = __t_tag_1
						if uint32(__t_tag_1) == 1527465420 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_2:
					__t11 = __t2
					goto end_branch_11
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 902936544 {
					var __t7 gopurs_runtime.Value
					{
						var __t_tag_4 uint32 = uint32(v1_1.IntVal)
						_ = __t_tag_4
						if uint32(__t_tag_4) == 902936544 {
							__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_7
						} else {

						}
					}
					{
						var __t_tag_5 uint32 = uint32(v1_1.IntVal)
						_ = __t_tag_5
						if uint32(__t_tag_5) == 1527465420 {
							__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
							goto end_branch_7
						} else {

						}
					}
					{
						var __t_tag_6 uint32 = uint32(v1_1.IntVal)
						_ = __t_tag_6
						if uint32(__t_tag_6) == 380165415 {
							__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
					}
				end_branch_7:
					__t11 = __t7
					goto end_branch_11
				} else {

				}
			}
			{
				var __t_tag_8 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_8
				if uint32(__t_tag_8) == 380165415 {
					var __t10 gopurs_runtime.Value
					{
						var __t_tag_9 uint32 = uint32(v1_1.IntVal)
						_ = __t_tag_9
						if uint32(__t_tag_9) == 380165415 {
							__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					}
				end_branch_10:
					__t11 = __t10
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_11:
			return __t11
		})})))}
	})
	return cache_Data_Ord_ordOrdering
}

var cache_Data_Ord_ordNumber gopurs_runtime.Value
var once_Data_Ord_ordNumber sync.Once

func Get_Data_Ord_ordNumber() gopurs_runtime.Value {
	once_Data_Ord_ordNumber.Do(func() {
		cache_Data_Ord_ordNumber = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3047586294_4177771502((&Constructor_Data_Ord_Ord[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_687527510_3790796878(Rebox_Data_Ord_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}
		}), gopurs_runtime.Apply3(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordNumber
}

var cache_Data_Ord_ordInt gopurs_runtime.Value
var once_Data_Ord_ordInt sync.Once

func Get_Data_Ord_ordInt() gopurs_runtime.Value {
	once_Data_Ord_ordInt.Do(func() {
		cache_Data_Ord_ordInt = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3308271157_4177771502((&Constructor_Data_Ord_Ord[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1053099733_3790796878(Rebox_Data_Ord_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
		}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordInt
}

var cache_Data_Ord_ordChar gopurs_runtime.Value
var once_Data_Ord_ordChar sync.Once

func Get_Data_Ord_ordChar() gopurs_runtime.Value {
	once_Data_Ord_ordChar.Do(func() {
		cache_Data_Ord_ordChar = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1140313009_3790796878(Rebox_Data_Ord_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqChar()))))}
		}), gopurs_runtime.Apply3(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordChar
}

var cache_Data_Ord_ordBoolean gopurs_runtime.Value
var once_Data_Ord_ordBoolean sync.Once

func Get_Data_Ord_ordBoolean() gopurs_runtime.Value {
	once_Data_Ord_ordBoolean.Do(func() {
		cache_Data_Ord_ordBoolean = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_219188042_4177771502((&Constructor_Data_Ord_Ord[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_2737952170_3790796878(Rebox_Data_Ord_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}
		}), gopurs_runtime.Apply3(Get_Data_Ord_ordBooleanImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})))}
	})
	return cache_Data_Ord_ordBoolean
}

var cache_Data_Ord_compareRecord gopurs_runtime.Value
var once_Data_Ord_compareRecord sync.Once

func Get_Data_Ord_compareRecord() gopurs_runtime.Value {
	once_Data_Ord_compareRecord.Do(func() {
		cache_Data_Ord_compareRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compareRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Ord_compareRecord
}

var cache_Data_Ord_ordRecord gopurs_runtime.Value
var once_Data_Ord_ordRecord sync.Once

func Get_Data_Ord_ordRecord() gopurs_runtime.Value {
	once_Data_Ord_ordRecord.Do(func() {
		cache_Data_Ord_ordRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictOrdRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_ordRecord(_dollar___unused_0_box, dictOrdRecord_1_box)
		})
	})
	return cache_Data_Ord_ordRecord
}

var cache_Data_Ord_compare1 gopurs_runtime.Value
var once_Data_Ord_compare1 sync.Once

func Get_Data_Ord_compare1() gopurs_runtime.Value {
	once_Data_Ord_compare1.Do(func() {
		cache_Data_Ord_compare1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Ord_compare1
}

var cache_Data_Ord_compare gopurs_runtime.Value
var once_Data_Ord_compare sync.Once

func Get_Data_Ord_compare() gopurs_runtime.Value {
	once_Data_Ord_compare.Do(func() {
		cache_Data_Ord_compare = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Ord_compare
}

var cache_Data_Ord_compare__3334414965 gopurs_runtime.Value
var once_Data_Ord_compare__3334414965 sync.Once

func Get_Data_Ord_compare__3334414965() gopurs_runtime.Value {
	once_Data_Ord_compare__3334414965.Do(func() {
		cache_Data_Ord_compare__3334414965 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3334414965(__eta0_0_box, __eta1_1_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__3334414965
}

var cache_Data_Ord_compare__2271344155 gopurs_runtime.Value
var once_Data_Ord_compare__2271344155 sync.Once

func Get_Data_Ord_compare__2271344155() gopurs_runtime.Value {
	once_Data_Ord_compare__2271344155.Do(func() {
		cache_Data_Ord_compare__2271344155 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__2271344155(__eta_norm_1_0_box, __eta_norm_0_1_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__2271344155
}

var cache_Data_Ord_compare__3891544906 gopurs_runtime.Value
var once_Data_Ord_compare__3891544906 sync.Once

func Get_Data_Ord_compare__3891544906() gopurs_runtime.Value {
	once_Data_Ord_compare__3891544906.Do(func() {
		cache_Data_Ord_compare__3891544906 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3891544906(__eta_norm_1_0_box, __eta_norm_0_1_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__3891544906
}

var cache_Data_Ord_compare__4139028341 gopurs_runtime.Value
var once_Data_Ord_compare__4139028341 sync.Once

func Get_Data_Ord_compare__4139028341() gopurs_runtime.Value {
	once_Data_Ord_compare__4139028341.Do(func() {
		cache_Data_Ord_compare__4139028341 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__4139028341(uint32(__eta0_0_box.IntVal), uint32(__eta1_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__4139028341
}

var cache_Data_Ord_compare__1150078980 gopurs_runtime.Value
var once_Data_Ord_compare__1150078980 sync.Once

func Get_Data_Ord_compare__1150078980() gopurs_runtime.Value {
	once_Data_Ord_compare__1150078980.Do(func() {
		cache_Data_Ord_compare__1150078980 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__1150078980(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](__eta0_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](__eta1_1_box))), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__1150078980
}

var cache_Data_Ord_compare__3124366677 gopurs_runtime.Value
var once_Data_Ord_compare__3124366677 sync.Once

func Get_Data_Ord_compare__3124366677() gopurs_runtime.Value {
	once_Data_Ord_compare__3124366677.Do(func() {
		cache_Data_Ord_compare__3124366677 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3124366677(uint32(__eta0_0_box.IntVal), uint32(__eta1_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__3124366677
}

var cache_Data_Ord_compare__78251971 gopurs_runtime.Value
var once_Data_Ord_compare__78251971 sync.Once

func Get_Data_Ord_compare__78251971() gopurs_runtime.Value {
	once_Data_Ord_compare__78251971.Do(func() {
		cache_Data_Ord_compare__78251971 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__78251971(__eta0_0_box, __eta1_1_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__78251971
}

var cache_Data_Ord_compare__4050261087 gopurs_runtime.Value
var once_Data_Ord_compare__4050261087 sync.Once

func Get_Data_Ord_compare__4050261087() gopurs_runtime.Value {
	once_Data_Ord_compare__4050261087.Do(func() {
		cache_Data_Ord_compare__4050261087 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__4050261087(__eta0_0_box, __eta1_1_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__4050261087
}

var cache_Data_Ord_compare__3685570754 gopurs_runtime.Value
var once_Data_Ord_compare__3685570754 sync.Once

func Get_Data_Ord_compare__3685570754() gopurs_runtime.Value {
	once_Data_Ord_compare__3685570754.Do(func() {
		cache_Data_Ord_compare__3685570754 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3685570754(Rebox_Data_Ord_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)), Rebox_Data_Ord_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__3685570754
}

var cache_Data_Ord_compare__1384306305 gopurs_runtime.Value
var once_Data_Ord_compare__1384306305 sync.Once

func Get_Data_Ord_compare__1384306305() gopurs_runtime.Value {
	once_Data_Ord_compare__1384306305.Do(func() {
		cache_Data_Ord_compare__1384306305 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__1384306305(Rebox_Data_Ord_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)), Rebox_Data_Ord_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__1384306305
}

var cache_Data_Ord_compare__749224996 gopurs_runtime.Value
var once_Data_Ord_compare__749224996 sync.Once

func Get_Data_Ord_compare__749224996() gopurs_runtime.Value {
	once_Data_Ord_compare__749224996.Do(func() {
		cache_Data_Ord_compare__749224996 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__749224996(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](__eta0_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](__eta1_1_box))), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__749224996
}

var cache_Data_Ord_compare__3024553627 gopurs_runtime.Value
var once_Data_Ord_compare__3024553627 sync.Once

func Get_Data_Ord_compare__3024553627() gopurs_runtime.Value {
	once_Data_Ord_compare__3024553627.Do(func() {
		cache_Data_Ord_compare__3024553627 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3024553627(__eta_norm_1_0_box, __eta_norm_0_1_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__3024553627
}

var cache_Data_Ord_compare__2273577763 gopurs_runtime.Value
var once_Data_Ord_compare__2273577763 sync.Once

func Get_Data_Ord_compare__2273577763() gopurs_runtime.Value {
	once_Data_Ord_compare__2273577763.Do(func() {
		cache_Data_Ord_compare__2273577763 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__2273577763(__eta0_0_box.StrVal(), __eta1_1_box.StrVal())), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__2273577763
}

var cache_Data_Ord_compare__1544596384 gopurs_runtime.Value
var once_Data_Ord_compare__1544596384 sync.Once

func Get_Data_Ord_compare__1544596384() gopurs_runtime.Value {
	once_Data_Ord_compare__1544596384.Do(func() {
		cache_Data_Ord_compare__1544596384 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__1544596384(__eta0_0_box.IntVal, __eta1_1_box.IntVal)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__1544596384
}

var cache_Data_Ord_compare__1452131379 gopurs_runtime.Value
var once_Data_Ord_compare__1452131379 sync.Once

func Get_Data_Ord_compare__1452131379() gopurs_runtime.Value {
	once_Data_Ord_compare__1452131379.Do(func() {
		cache_Data_Ord_compare__1452131379 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__1452131379(__eta0_0_box.IntVal, __eta1_1_box.IntVal)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__1452131379
}

var cache_Data_Ord_compare__1925727043 gopurs_runtime.Value
var once_Data_Ord_compare__1925727043 sync.Once

func Get_Data_Ord_compare__1925727043() gopurs_runtime.Value {
	once_Data_Ord_compare__1925727043.Do(func() {
		cache_Data_Ord_compare__1925727043 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__1925727043(__eta0_0_box.IntVal, __eta1_1_box.IntVal)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__1925727043
}

var cache_Data_Ord_compare__205334653 gopurs_runtime.Value
var once_Data_Ord_compare__205334653 sync.Once

func Get_Data_Ord_compare__205334653() gopurs_runtime.Value {
	once_Data_Ord_compare__205334653.Do(func() {
		cache_Data_Ord_compare__205334653 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__205334653(__eta0_0_box.IntVal, __eta1_1_box.IntVal)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__205334653
}

var cache_Data_Ord_compare__1652089056 gopurs_runtime.Value
var once_Data_Ord_compare__1652089056 sync.Once

func Get_Data_Ord_compare__1652089056() gopurs_runtime.Value {
	once_Data_Ord_compare__1652089056.Do(func() {
		cache_Data_Ord_compare__1652089056 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__1652089056(__eta0_0_box.IntVal, __eta1_1_box.IntVal)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__1652089056
}

var cache_Data_Ord_compare__100490739 gopurs_runtime.Value
var once_Data_Ord_compare__100490739 sync.Once

func Get_Data_Ord_compare__100490739() gopurs_runtime.Value {
	once_Data_Ord_compare__100490739.Do(func() {
		cache_Data_Ord_compare__100490739 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__100490739(__eta0_0_box.IntVal, __eta1_1_box.IntVal)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__100490739
}

var cache_Data_Ord_compare__276935885 gopurs_runtime.Value
var once_Data_Ord_compare__276935885 sync.Once

func Get_Data_Ord_compare__276935885() gopurs_runtime.Value {
	once_Data_Ord_compare__276935885.Do(func() {
		cache_Data_Ord_compare__276935885 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__276935885(__eta0_0_box.IntVal, __eta1_1_box.IntVal)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__276935885
}

var cache_Data_Ord_compare__2789112131 gopurs_runtime.Value
var once_Data_Ord_compare__2789112131 sync.Once

func Get_Data_Ord_compare__2789112131() gopurs_runtime.Value {
	once_Data_Ord_compare__2789112131.Do(func() {
		cache_Data_Ord_compare__2789112131 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__2789112131(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal())), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__2789112131
}

var cache_Data_Ord_compare__3089094802 gopurs_runtime.Value
var once_Data_Ord_compare__3089094802 sync.Once

func Get_Data_Ord_compare__3089094802() gopurs_runtime.Value {
	once_Data_Ord_compare__3089094802.Do(func() {
		cache_Data_Ord_compare__3089094802 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3089094802(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal())), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__3089094802
}

var cache_Data_Ord_compare__3009265731 gopurs_runtime.Value
var once_Data_Ord_compare__3009265731 sync.Once

func Get_Data_Ord_compare__3009265731() gopurs_runtime.Value {
	once_Data_Ord_compare__3009265731.Do(func() {
		cache_Data_Ord_compare__3009265731 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__3009265731(__eta0_0_box.StrVal(), __eta1_1_box.StrVal())), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__3009265731
}

var cache_Data_Ord_compare__1263564168 gopurs_runtime.Value
var once_Data_Ord_compare__1263564168 sync.Once

func Get_Data_Ord_compare__1263564168() gopurs_runtime.Value {
	once_Data_Ord_compare__1263564168.Do(func() {
		cache_Data_Ord_compare__1263564168 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_compare__1263564168(__eta0_0_box.StrVal(), __eta1_1_box.StrVal())), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_compare__1263564168
}

var cache_Data_Ord_compare2 gopurs_runtime.Value
var once_Data_Ord_compare2 sync.Once

func Get_Data_Ord_compare2() gopurs_runtime.Value {
	once_Data_Ord_compare2.Do(func() {
		cache_Data_Ord_compare2 = Get_Data_Ord_compare__1925727043()
	})
	return cache_Data_Ord_compare2
}

var cache_Data_Ord_comparing gopurs_runtime.Value
var once_Data_Ord_comparing sync.Once

func Get_Data_Ord_comparing() gopurs_runtime.Value {
	once_Data_Ord_comparing.Do(func() {
		cache_Data_Ord_comparing = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_comparing(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Ord_comparing
}

var cache_Data_Ord_greaterThan gopurs_runtime.Value
var once_Data_Ord_greaterThan sync.Once

func Get_Data_Ord_greaterThan() gopurs_runtime.Value {
	once_Data_Ord_greaterThan.Do(func() {
		cache_Data_Ord_greaterThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThan(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
		})
	})
	return cache_Data_Ord_greaterThan
}

var cache_Data_Ord_greaterThan__2273689272 gopurs_runtime.Value
var once_Data_Ord_greaterThan__2273689272 sync.Once

func Get_Data_Ord_greaterThan__2273689272() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__2273689272.Do(func() {
		cache_Data_Ord_greaterThan__2273689272 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__2273689272(Rebox_Data_Ord_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)), Rebox_Data_Ord_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box))))
		})
	})
	return cache_Data_Ord_greaterThan__2273689272
}

var cache_Data_Ord_greaterThan__3348776441 gopurs_runtime.Value
var once_Data_Ord_greaterThan__3348776441 sync.Once

func Get_Data_Ord_greaterThan__3348776441() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__3348776441.Do(func() {
		cache_Data_Ord_greaterThan__3348776441 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__3348776441(__eta0_0_box.IntVal, __eta1_1_box.IntVal))
		})
	})
	return cache_Data_Ord_greaterThan__3348776441
}

var cache_Data_Ord_greaterThan__756116985 gopurs_runtime.Value
var once_Data_Ord_greaterThan__756116985 sync.Once

func Get_Data_Ord_greaterThan__756116985() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__756116985.Do(func() {
		cache_Data_Ord_greaterThan__756116985 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__756116985(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal()))
		})
	})
	return cache_Data_Ord_greaterThan__756116985
}

var cache_Data_Ord_greaterThan__3845833352 gopurs_runtime.Value
var once_Data_Ord_greaterThan__3845833352 sync.Once

func Get_Data_Ord_greaterThan__3845833352() gopurs_runtime.Value {
	once_Data_Ord_greaterThan__3845833352.Do(func() {
		cache_Data_Ord_greaterThan__3845833352 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThan__3845833352(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal()))
		})
	})
	return cache_Data_Ord_greaterThan__3845833352
}

var cache_Data_Ord_greaterThanOrEq gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq sync.Once

func Get_Data_Ord_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq.Do(func() {
		cache_Data_Ord_greaterThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
		})
	})
	return cache_Data_Ord_greaterThanOrEq
}

var cache_Data_Ord_greaterThanOrEq__4238969209 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__4238969209 sync.Once

func Get_Data_Ord_greaterThanOrEq__4238969209() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__4238969209.Do(func() {
		cache_Data_Ord_greaterThanOrEq__4238969209 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__4238969209(__eta0_0_box.StrVal(), __eta1_1_box.StrVal()))
		})
	})
	return cache_Data_Ord_greaterThanOrEq__4238969209
}

var cache_Data_Ord_greaterThanOrEq__3348776441 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__3348776441 sync.Once

func Get_Data_Ord_greaterThanOrEq__3348776441() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__3348776441.Do(func() {
		cache_Data_Ord_greaterThanOrEq__3348776441 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__3348776441(__eta0_0_box.IntVal, __eta1_1_box.IntVal))
		})
	})
	return cache_Data_Ord_greaterThanOrEq__3348776441
}

var cache_Data_Ord_greaterThanOrEq__756116985 gopurs_runtime.Value
var once_Data_Ord_greaterThanOrEq__756116985 sync.Once

func Get_Data_Ord_greaterThanOrEq__756116985() gopurs_runtime.Value {
	once_Data_Ord_greaterThanOrEq__756116985.Do(func() {
		cache_Data_Ord_greaterThanOrEq__756116985 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_greaterThanOrEq__756116985(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal()))
		})
	})
	return cache_Data_Ord_greaterThanOrEq__756116985
}

var cache_Data_Ord_lessThan gopurs_runtime.Value
var once_Data_Ord_lessThan sync.Once

func Get_Data_Ord_lessThan() gopurs_runtime.Value {
	once_Data_Ord_lessThan.Do(func() {
		cache_Data_Ord_lessThan = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThan(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
		})
	})
	return cache_Data_Ord_lessThan
}

var cache_Data_Ord_lessThan__3348776441 gopurs_runtime.Value
var once_Data_Ord_lessThan__3348776441 sync.Once

func Get_Data_Ord_lessThan__3348776441() gopurs_runtime.Value {
	once_Data_Ord_lessThan__3348776441.Do(func() {
		cache_Data_Ord_lessThan__3348776441 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThan__3348776441(__eta0_0_box.IntVal, __eta1_1_box.IntVal))
		})
	})
	return cache_Data_Ord_lessThan__3348776441
}

var cache_Data_Ord_lessThan__756116985 gopurs_runtime.Value
var once_Data_Ord_lessThan__756116985 sync.Once

func Get_Data_Ord_lessThan__756116985() gopurs_runtime.Value {
	once_Data_Ord_lessThan__756116985.Do(func() {
		cache_Data_Ord_lessThan__756116985 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThan__756116985(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal()))
		})
	})
	return cache_Data_Ord_lessThan__756116985
}

var cache_Data_Ord_lessThan__3845833352 gopurs_runtime.Value
var once_Data_Ord_lessThan__3845833352 sync.Once

func Get_Data_Ord_lessThan__3845833352() gopurs_runtime.Value {
	once_Data_Ord_lessThan__3845833352.Do(func() {
		cache_Data_Ord_lessThan__3845833352 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThan__3845833352(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal()))
		})
	})
	return cache_Data_Ord_lessThan__3845833352
}

var cache_Data_Ord_signum gopurs_runtime.Value
var once_Data_Ord_signum sync.Once

func Get_Data_Ord_signum() gopurs_runtime.Value {
	once_Data_Ord_signum.Do(func() {
		cache_Data_Ord_signum = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_signum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dictRing_1_box))
		})
	})
	return cache_Data_Ord_signum
}

var cache_Data_Ord_lessThanOrEq gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq sync.Once

func Get_Data_Ord_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq.Do(func() {
		cache_Data_Ord_lessThanOrEq = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
		})
	})
	return cache_Data_Ord_lessThanOrEq
}

var cache_Data_Ord_lessThanOrEq__4238969209 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__4238969209 sync.Once

func Get_Data_Ord_lessThanOrEq__4238969209() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__4238969209.Do(func() {
		cache_Data_Ord_lessThanOrEq__4238969209 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__4238969209(__eta0_0_box.StrVal(), __eta1_1_box.StrVal()))
		})
	})
	return cache_Data_Ord_lessThanOrEq__4238969209
}

var cache_Data_Ord_lessThanOrEq__3348776441 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__3348776441 sync.Once

func Get_Data_Ord_lessThanOrEq__3348776441() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__3348776441.Do(func() {
		cache_Data_Ord_lessThanOrEq__3348776441 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__3348776441(__eta0_0_box.IntVal, __eta1_1_box.IntVal))
		})
	})
	return cache_Data_Ord_lessThanOrEq__3348776441
}

var cache_Data_Ord_lessThanOrEq__756116985 gopurs_runtime.Value
var once_Data_Ord_lessThanOrEq__756116985 sync.Once

func Get_Data_Ord_lessThanOrEq__756116985() gopurs_runtime.Value {
	once_Data_Ord_lessThanOrEq__756116985.Do(func() {
		cache_Data_Ord_lessThanOrEq__756116985 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_lessThanOrEq__756116985(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal()))
		})
	})
	return cache_Data_Ord_lessThanOrEq__756116985
}

var cache_Data_Ord_max gopurs_runtime.Value
var once_Data_Ord_max sync.Once

func Get_Data_Ord_max() gopurs_runtime.Value {
	once_Data_Ord_max.Do(func() {
		cache_Data_Ord_max = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_max(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
		})
	})
	return cache_Data_Ord_max
}

var cache_Data_Ord_max__4009478382 gopurs_runtime.Value
var once_Data_Ord_max__4009478382 sync.Once

func Get_Data_Ord_max__4009478382() gopurs_runtime.Value {
	once_Data_Ord_max__4009478382.Do(func() {
		cache_Data_Ord_max__4009478382 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Ord_max__4009478382(__eta0_0_box.IntVal, __eta1_1_box.IntVal))
		})
	})
	return cache_Data_Ord_max__4009478382
}

var cache_Data_Ord_min gopurs_runtime.Value
var once_Data_Ord_min sync.Once

func Get_Data_Ord_min() gopurs_runtime.Value {
	once_Data_Ord_min.Do(func() {
		cache_Data_Ord_min = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_min(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, y_2_box)
		})
	})
	return cache_Data_Ord_min
}

var cache_Data_Ord_min__4009478382 gopurs_runtime.Value
var once_Data_Ord_min__4009478382 sync.Once

func Get_Data_Ord_min__4009478382() gopurs_runtime.Value {
	once_Data_Ord_min__4009478382.Do(func() {
		cache_Data_Ord_min__4009478382 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Ord_min__4009478382(__eta0_0_box.IntVal, __eta1_1_box.IntVal))
		})
	})
	return cache_Data_Ord_min__4009478382
}

var cache_Data_Ord_ordArray gopurs_runtime.Value
var once_Data_Ord_ordArray sync.Once

func Get_Data_Ord_ordArray() gopurs_runtime.Value {
	once_Data_Ord_ordArray.Do(func() {
		cache_Data_Ord_ordArray = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_ordArray(dictOrd_0_box)
		})
	})
	return cache_Data_Ord_ordArray
}

var cache_Data_Ord_ord1Array gopurs_runtime.Value
var once_Data_Ord_ord1Array sync.Once

func Get_Data_Ord_ord1Array() gopurs_runtime.Value {
	once_Data_Ord_ord1Array.Do(func() {
		cache_Data_Ord_ord1Array = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Ord_ordArray(dictOrd_0)))
		})}))}
	})
	return cache_Data_Ord_ord1Array
}

var cache_Data_Ord_ordRecordCons gopurs_runtime.Value
var once_Data_Ord_ordRecordCons sync.Once

func Get_Data_Ord_ordRecordCons() gopurs_runtime.Value {
	once_Data_Ord_ordRecordCons.Do(func() {
		cache_Data_Ord_ordRecordCons = gopurs_runtime.Func(func(dictOrdRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_ordRecordCons(dictOrdRecord_0_box)
		})
	})
	return cache_Data_Ord_ordRecordCons
}

var cache_Data_Ord_clamp gopurs_runtime.Value
var once_Data_Ord_clamp sync.Once

func Get_Data_Ord_clamp() gopurs_runtime.Value {
	once_Data_Ord_clamp.Do(func() {
		cache_Data_Ord_clamp = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_clamp(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), low_1_box, hi_2_box, x_3_box)
		})
	})
	return cache_Data_Ord_clamp
}

var cache_Data_Ord_clamp__1405043522 gopurs_runtime.Value
var once_Data_Ord_clamp__1405043522 sync.Once

func Get_Data_Ord_clamp__1405043522() gopurs_runtime.Value {
	once_Data_Ord_clamp__1405043522.Do(func() {
		cache_Data_Ord_clamp__1405043522 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Ord_clamp__1405043522(__eta0_0_box.IntVal, __eta1_1_box.IntVal, __eta2_2_box.IntVal))
		})
	})
	return cache_Data_Ord_clamp__1405043522
}

var cache_Data_Ord_between gopurs_runtime.Value
var once_Data_Ord_between sync.Once

func Get_Data_Ord_between() gopurs_runtime.Value {
	once_Data_Ord_between.Do(func() {
		cache_Data_Ord_between = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, low_1_box gopurs_runtime.Value, hi_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Ord_between(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), low_1_box, hi_2_box, x_3_box))
		})
	})
	return cache_Data_Ord_between
}

var cache_Data_Ord_abs gopurs_runtime.Value
var once_Data_Ord_abs sync.Once

func Get_Data_Ord_abs() gopurs_runtime.Value {
	once_Data_Ord_abs.Do(func() {
		cache_Data_Ord_abs = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_abs(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dictRing_1_box))
		})
	})
	return cache_Data_Ord_abs
}

type Constructor_Data_Ord_OrdRecord[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004f00720064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004f00720064005f004f00720064005200650063006f00720064
func init() {
	gopurs_runtime.StructGetters[4162894775] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "EqRecord0":
			return gopurs_runtime.Box(c.V0)
		case "compareRecord":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Ord_OrdRecord: " + key)
		}
	}
}

type Constructor_Data_Ord_Ord[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004f00720064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004f00720064005f004f00720064
func init() {
	gopurs_runtime.StructGetters[1435789946] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Eq0":
			return gopurs_runtime.Box(c.V0)
		case "compare":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Ord_Ord: " + key)
		}
	}
}

type Constructor_Data_Ord_Ord1[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004f00720064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004f00720064005f004f007200640031
func init() {
	gopurs_runtime.StructGetters[1632188299] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_Ord1[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Eq10":
			return gopurs_runtime.Box(c.V0)
		case "compare1":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Ord_Ord1: " + key)
		}
	}
}

func Call_Data_Ord_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Data_Ord_OrdRecord_dollar_Dict(x_0_loop struct {
	EqRecord0     gopurs_runtime.Value
	compareRecord gopurs_runtime.Value
	/*gopurs:record-v1:2:00450071005200650063006f007200640030,0063006f006d0070006100720065005200650063006f00720064*/
}) *Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		EqRecord0     gopurs_runtime.Value
		compareRecord gopurs_runtime.Value
		/*gopurs:record-v1:2:00450071005200650063006f007200640030,0063006f006d0070006100720065005200650063006f00720064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", orig.EqRecord0, orig.compareRecord)
	}())
}

func Call_Data_Ord_Ord_dollar_Dict(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	var x_0 struct {
		Eq0     gopurs_runtime.Value
		compare gopurs_runtime.Value
		/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
	}())
}

func Call_Data_Ord_Ord_dollar_Dict__3648059431(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__3648059431:
	for {
		if false {
			continue Ord_dollar_Dict__3648059431
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2520386791(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__2520386791:
	for {
		if false {
			continue Ord_dollar_Dict__2520386791
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__3832589959(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__3832589959:
	for {
		if false {
			continue Ord_dollar_Dict__3832589959
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__4094323175(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime] {
Ord_dollar_Dict__4094323175:
	for {
		if false {
			continue Ord_dollar_Dict__4094323175
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__757207815(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[uint32] {
Ord_dollar_Dict__757207815:
	for {
		if false {
			continue Ord_dollar_Dict__757207815
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2432464903(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[uint32] {
Ord_dollar_Dict__2432464903:
	for {
		if false {
			continue Ord_dollar_Dict__2432464903
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__3848172583(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date] {
Ord_dollar_Dict__3848172583:
	for {
		if false {
			continue Ord_dollar_Dict__3848172583
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__3034920007(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[uint32] {
Ord_dollar_Dict__3034920007:
	for {
		if false {
			continue Ord_dollar_Dict__3034920007
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2010913863(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[uint32] {
Ord_dollar_Dict__2010913863:
	for {
		if false {
			continue Ord_dollar_Dict__2010913863
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__3994504999(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__3994504999:
	for {
		if false {
			continue Ord_dollar_Dict__3994504999
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2123630727(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__2123630727:
	for {
		if false {
			continue Ord_dollar_Dict__2123630727
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__1187641767(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[uint32] {
Ord_dollar_Dict__1187641767:
	for {
		if false {
			continue Ord_dollar_Dict__1187641767
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__1570991783(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[uint32] {
Ord_dollar_Dict__1570991783:
	for {
		if false {
			continue Ord_dollar_Dict__1570991783
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2690096935(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time] {
Ord_dollar_Dict__2690096935:
	for {
		if false {
			continue Ord_dollar_Dict__2690096935
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__832340135(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__832340135:
	for {
		if false {
			continue Ord_dollar_Dict__832340135
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__1238669767(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__1238669767:
	for {
		if false {
			continue Ord_dollar_Dict__1238669767
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__1954751335(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__1954751335:
	for {
		if false {
			continue Ord_dollar_Dict__1954751335
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__4182661159(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[uint32] {
Ord_dollar_Dict__4182661159:
	for {
		if false {
			continue Ord_dollar_Dict__4182661159
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2336892295(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[bool] {
Ord_dollar_Dict__2336892295:
	for {
		if false {
			continue Ord_dollar_Dict__2336892295
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[bool]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__4108678119(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[string] {
Ord_dollar_Dict__4108678119:
	for {
		if false {
			continue Ord_dollar_Dict__4108678119
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2468619815(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[int64] {
Ord_dollar_Dict__2468619815:
	for {
		if false {
			continue Ord_dollar_Dict__2468619815
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__3472323847(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[float64] {
Ord_dollar_Dict__3472323847:
	for {
		if false {
			continue Ord_dollar_Dict__3472323847
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2744908871(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[string] {
Ord_dollar_Dict__2744908871:
	for {
		if false {
			continue Ord_dollar_Dict__2744908871
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord_dollar_Dict__2297783463(x_0_loop struct {
	Eq0     gopurs_runtime.Value
	compare gopurs_runtime.Value
	/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
}) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
Ord_dollar_Dict__2297783463:
	for {
		if false {
			continue Ord_dollar_Dict__2297783463
		}
		var x_0 struct {
			Eq0     gopurs_runtime.Value
			compare gopurs_runtime.Value
			/*gopurs:record-v1:2:004500710030,0063006f006d0070006100720065*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Eq0", "compare", orig.Eq0, orig.compare)
		}())
	}
}

func Call_Data_Ord_Ord1_dollar_Dict(x_0_loop struct {
	Eq10     gopurs_runtime.Value
	compare1 gopurs_runtime.Value
	/*gopurs:record-v1:2:0045007100310030,0063006f006d00700061007200650031*/
}) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	var x_0 struct {
		Eq10     gopurs_runtime.Value
		compare1 gopurs_runtime.Value
		/*gopurs:record-v1:2:0045007100310030,0063006f006d00700061007200650031*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Eq10", "compare1", orig.Eq10, orig.compare1)
	}())
}

func Call_Data_Ord_compareRecord(dict_0_loop *Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Ord_ordRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictOrdRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictOrdRecord_1 gopurs_runtime.Value = dictOrdRecord_1_loop
	_ = dictOrdRecord_1
	// TAST (Let): eqRec1_2_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(Record (Row [] (TypeVar row$scope20)))])
	eqRec1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "EqRecord0"), gopurs_runtime.Value{})))
	_ = eqRec1_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqRec1_2_0)}
	}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Ord_compare1(dict_0_loop *Constructor_Data_Ord_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Ord_compare(dict_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Ord_compare__3334414965(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) uint32 {
compare__3334414965:
	for {
		if false {
			continue compare__3334414965
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Argonaut_Core_ordJson()).V1, __eta0_0, __eta1_1).IntVal)
	}
}

func Call_Data_Ord_compare__2271344155(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) uint32 {
compare__2271344155:
	for {
		if false {
			continue compare__2271344155
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Argonaut_Core_ordJson()).V1, __eta_norm_1_0, __eta_norm_0_1).IntVal)
	}
}

func Call_Data_Ord_compare__3891544906(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) uint32 {
compare__3891544906:
	for {
		if false {
			continue compare__3891544906
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Error_ordJsonDecodeError()).V1, __eta_norm_1_0, __eta_norm_0_1).IntVal)
	}
}

func Call_Data_Ord_compare__4139028341(__eta0_0_loop uint32, __eta1_1_loop uint32) uint32 {
compare__4139028341:
	for {
		if false {
			continue compare__4139028341
		}
		var __eta0_0 uint32 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 uint32 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_Component_ordMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta0_0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta1_1), UnsafePtr: nil}).IntVal)
	}
}

func Call_Data_Ord_compare__1150078980(__eta0_0_loop *Constructor_Data_Date_Date, __eta1_1_loop *Constructor_Data_Date_Date) uint32 {
compare__1150078980:
	for {
		if false {
			continue compare__1150078980
		}
		var __eta0_0 *Constructor_Data_Date_Date = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Date_Date = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_758368489(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_ordDate())).V1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__eta1_1)}).IntVal)
	}
}

func Call_Data_Ord_compare__3124366677(__eta0_0_loop uint32, __eta1_1_loop uint32) uint32 {
compare__3124366677:
	for {
		if false {
			continue compare__3124366677
		}
		var __eta0_0 uint32 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 uint32 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Interval_Duration_ordDurationComponent())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta0_0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta1_1), UnsafePtr: nil}).IntVal)
	}
}

func Call_Data_Ord_compare__78251971(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) uint32 {
compare__78251971:
	for {
		if false {
			continue compare__78251971
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Map_Internal_ordMap(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3730953251_4177771502(Rebox_Data_Ord_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Interval_Duration_ordDurationComponent()))))}, gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3047586294_4177771502(Rebox_Data_Ord_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}), "compare"), __eta0_0, __eta1_1).IntVal)
	}
}

func Call_Data_Ord_compare__4050261087(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) uint32 {
compare__4050261087:
	for {
		if false {
			continue compare__4050261087
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Interval_Duration_ordDuration()).V1, __eta0_0, __eta1_1).IntVal)
	}
}

func Call_Data_Ord_compare__3685570754(__eta0_0_loop *Constructor_Data_Maybe_Just[int64], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) uint32 {
compare__3685570754:
	for {
		if false {
			continue compare__3685570754
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[int64] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_Maybe_ordMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3308271157_4177771502(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_Component_ordDay()))))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(__eta0_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(__eta1_1))}).IntVal)
	}
}

func Call_Data_Ord_compare__1384306305(__eta0_0_loop *Constructor_Data_Maybe_Just[int64], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) uint32 {
compare__1384306305:
	for {
		if false {
			continue compare__1384306305
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[int64] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_Maybe_ordMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3308271157_4177771502(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(__eta0_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(__eta1_1))}).IntVal)
	}
}

func Call_Data_Ord_compare__749224996(__eta0_0_loop *Constructor_Data_Time_Time, __eta1_1_loop *Constructor_Data_Time_Time) uint32 {
compare__749224996:
	for {
		if false {
			continue compare__749224996
		}
		var __eta0_0 *Constructor_Data_Time_Time = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Time_Time = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3383090825(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_ordTime())).V1, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(__eta1_1)}).IntVal)
	}
}

func Call_Data_Ord_compare__3024553627(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) uint32 {
compare__3024553627:
	for {
		if false {
			continue compare__3024553627
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return uint32(gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Foreign_ordForeignError()).V1, __eta_norm_1_0, __eta_norm_0_1).IntVal)
	}
}

func Call_Data_Ord_compare__2273577763(__eta0_0_loop string, __eta1_1_loop string) uint32 {
compare__2273577763:
	for {
		if false {
			continue compare__2273577763
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar())).V1, gopurs_runtime.Str(__eta0_0), gopurs_runtime.Str(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__1544596384(__eta0_0_loop int64, __eta1_1_loop int64) uint32 {
compare__1544596384:
	for {
		if false {
			continue compare__1544596384
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_Component_ordDay())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__1452131379(__eta0_0_loop int64, __eta1_1_loop int64) uint32 {
compare__1452131379:
	for {
		if false {
			continue compare__1452131379
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_Component_ordYear())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__1925727043(__eta0_0_loop int64, __eta1_1_loop int64) uint32 {
compare__1925727043:
	for {
		if false {
			continue compare__1925727043
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__205334653(__eta0_0_loop int64, __eta1_1_loop int64) uint32 {
compare__205334653:
	for {
		if false {
			continue compare__205334653
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_Component_ordHour())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__1652089056(__eta0_0_loop int64, __eta1_1_loop int64) uint32 {
compare__1652089056:
	for {
		if false {
			continue compare__1652089056
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_Component_ordMillisecond())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__100490739(__eta0_0_loop int64, __eta1_1_loop int64) uint32 {
compare__100490739:
	for {
		if false {
			continue compare__100490739
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_Component_ordMinute())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__276935885(__eta0_0_loop int64, __eta1_1_loop int64) uint32 {
compare__276935885:
	for {
		if false {
			continue compare__276935885
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_Component_ordSecond())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__2789112131(__eta0_0_loop float64, __eta1_1_loop float64) uint32 {
compare__2789112131:
	for {
		if false {
			continue compare__2789112131
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber())).V1, gopurs_runtime.Float(__eta0_0), gopurs_runtime.Float(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__3089094802(__eta0_0_loop float64, __eta1_1_loop float64) uint32 {
compare__3089094802:
	for {
		if false {
			continue compare__3089094802
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_Duration_ordMilliseconds())).V1, gopurs_runtime.Float(__eta0_0), gopurs_runtime.Float(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__3009265731(__eta0_0_loop string, __eta1_1_loop string) uint32 {
compare__3009265731:
	for {
		if false {
			continue compare__3009265731
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.Str(__eta0_0), gopurs_runtime.Str(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_compare__1263564168(__eta0_0_loop string, __eta1_1_loop string) uint32 {
compare__1263564168:
	for {
		if false {
			continue compare__1263564168
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		return uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_String_NonEmpty_Internal_ordNonEmptyString())).V1, gopurs_runtime.Str(__eta0_0), gopurs_runtime.Str(__eta1_1)).IntVal)
	}
}

func Call_Data_Ord_comparing(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	var y_3 gopurs_runtime.Value = y_3_loop
	_ = y_3
	return uint32(gopurs_runtime.Apply2(dictOrd_0.V1, gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}

func Call_Data_Ord_greaterThan(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var a1_1 gopurs_runtime.Value = a1_1_loop
	_ = a1_1
	var a2_2 gopurs_runtime.Value = a2_2_loop
	_ = a2_2
	var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2).IntVal)
	_ = __t_tag_0
	return (uint32(__t_tag_0) == 380165415)
}

func Call_Data_Ord_greaterThan__2273689272(__eta0_0_loop *Constructor_Data_Maybe_Just[int64], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
greaterThan__2273689272:
	for {
		if false {
			continue greaterThan__2273689272
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[int64] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_Maybe_ordMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_3308271157_4177771502(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_Component_ordDay()))))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(__eta0_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1170268447_3094389156(__eta1_1))})
		_ = __t_tag_0
		return (uint32(__t_tag_0.IntVal) == 380165415)
	}
}

func Call_Data_Ord_greaterThan__3348776441(__eta0_0_loop int64, __eta1_1_loop int64) bool {
greaterThan__3348776441:
	for {
		if false {
			continue greaterThan__3348776441
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) > (__eta1_1)
	}
}

func Call_Data_Ord_greaterThan__756116985(__eta0_0_loop float64, __eta1_1_loop float64) bool {
greaterThan__756116985:
	for {
		if false {
			continue greaterThan__756116985
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) > (__eta1_1)
	}
}

func Call_Data_Ord_greaterThan__3845833352(__eta0_0_loop float64, __eta1_1_loop float64) bool {
greaterThan__3845833352:
	for {
		if false {
			continue greaterThan__3845833352
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_Duration_ordMilliseconds())).V1, gopurs_runtime.Float(__eta0_0), gopurs_runtime.Float(__eta1_1))
		_ = __t_tag_0
		return (uint32(__t_tag_0.IntVal) == 380165415)
	}
}

func Call_Data_Ord_greaterThanOrEq(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var a1_1 gopurs_runtime.Value = a1_1_loop
	_ = a1_1
	var a2_2 gopurs_runtime.Value = a2_2_loop
	_ = a2_2
	var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2).IntVal)
	_ = __t_tag_0
	return (uint32(__t_tag_0) == 1527465420) != (true)
}

func Call_Data_Ord_greaterThanOrEq__4238969209(__eta0_0_loop string, __eta1_1_loop string) bool {
greaterThanOrEq__4238969209:
	for {
		if false {
			continue greaterThanOrEq__4238969209
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) >= (__eta1_1)
	}
}

func Call_Data_Ord_greaterThanOrEq__3348776441(__eta0_0_loop int64, __eta1_1_loop int64) bool {
greaterThanOrEq__3348776441:
	for {
		if false {
			continue greaterThanOrEq__3348776441
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) >= (__eta1_1)
	}
}

func Call_Data_Ord_greaterThanOrEq__756116985(__eta0_0_loop float64, __eta1_1_loop float64) bool {
greaterThanOrEq__756116985:
	for {
		if false {
			continue greaterThanOrEq__756116985
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) >= (__eta1_1)
	}
}

func Call_Data_Ord_lessThan(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var a1_1 gopurs_runtime.Value = a1_1_loop
	_ = a1_1
	var a2_2 gopurs_runtime.Value = a2_2_loop
	_ = a2_2
	var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2).IntVal)
	_ = __t_tag_0
	return (uint32(__t_tag_0) == 1527465420)
}

func Call_Data_Ord_lessThan__3348776441(__eta0_0_loop int64, __eta1_1_loop int64) bool {
lessThan__3348776441:
	for {
		if false {
			continue lessThan__3348776441
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) < (__eta1_1)
	}
}

func Call_Data_Ord_lessThan__756116985(__eta0_0_loop float64, __eta1_1_loop float64) bool {
lessThan__756116985:
	for {
		if false {
			continue lessThan__756116985
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) < (__eta1_1)
	}
}

func Call_Data_Ord_lessThan__3845833352(__eta0_0_loop float64, __eta1_1_loop float64) bool {
lessThan__3845833352:
	for {
		if false {
			continue lessThan__3845833352
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_Duration_ordMilliseconds())).V1, gopurs_runtime.Float(__eta0_0), gopurs_runtime.Float(__eta1_1))
		_ = __t_tag_0
		return (uint32(__t_tag_0.IntVal) == 1527465420)
	}
}

func Call_Data_Ord_signum(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictRing_1_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var dictRing_1 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dictRing_1_loop
	_ = dictRing_1
	// TAST (Let): Semiring0_2_0 shape=App(Other) bindingType=Any
	Semiring0_2_0 := gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{})
	_ = Semiring0_2_0
	// TAST (Let): zero_3_1 shape=App(Var) bindingType=(TypeVar a$scope39)
	zero_3_1 := Call_Data_Semiring_zero(Semiring0_2_0)
	_ = zero_3_1
	// TAST (Let): zero_4_2 shape=Other bindingType=(TypeVar a$scope39)
	zero_4_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "zero")
	_ = zero_4_2
	// TAST (Let): one_5_3 shape=Other bindingType=(TypeVar a$scope39)
	one_5_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "one")
	_ = one_5_3
	// TAST (Let): one1_6_4 shape=App(Var) bindingType=(TypeVar a$scope39)
	one1_6_4 := Call_Data_Semiring_one(Semiring0_2_0)
	_ = one1_6_4
	return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t8 gopurs_runtime.Value
		{
			var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_7, zero_3_1)
			_ = __t_tag_7
			if uint32(__t_tag_7.IntVal) == 1527465420 {
				__t8 = gopurs_runtime.Apply2(dictRing_1.V1, zero_4_2, one_5_3)
				goto end_branch_8
			} else {

			}
		}
		{
			var __t6 gopurs_runtime.Value
			{
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_7, zero_3_1)
				_ = __t_tag_5
				if uint32(__t_tag_5.IntVal) == 380165415 {
					__t6 = one1_6_4
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = x_7
			}
		end_branch_6:
			__t8 = __t6
		}
	end_branch_8:
		return __t8
	})
}

func Call_Data_Ord_lessThanOrEq(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var a1_1 gopurs_runtime.Value = a1_1_loop
	_ = a1_1
	var a2_2 gopurs_runtime.Value = a2_2_loop
	_ = a2_2
	var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2).IntVal)
	_ = __t_tag_0
	return (uint32(__t_tag_0) == 380165415) != (true)
}

func Call_Data_Ord_lessThanOrEq__4238969209(__eta0_0_loop string, __eta1_1_loop string) bool {
lessThanOrEq__4238969209:
	for {
		if false {
			continue lessThanOrEq__4238969209
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) <= (__eta1_1)
	}
}

func Call_Data_Ord_lessThanOrEq__3348776441(__eta0_0_loop int64, __eta1_1_loop int64) bool {
lessThanOrEq__3348776441:
	for {
		if false {
			continue lessThanOrEq__3348776441
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) <= (__eta1_1)
	}
}

func Call_Data_Ord_lessThanOrEq__756116985(__eta0_0_loop float64, __eta1_1_loop float64) bool {
lessThanOrEq__756116985:
	for {
		if false {
			continue lessThanOrEq__756116985
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) <= (__eta1_1)
	}
}

func Call_Data_Ord_max(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var y_2 gopurs_runtime.Value = y_2_loop
	_ = y_2
	// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
	v_3_0 := uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2).IntVal)
	_ = v_3_0
	var __t1 gopurs_runtime.Value
	{
		if v_3_0 == 1527465420 {
			__t1 = y_2
			goto end_branch_1
		} else {

		}
	}
	{
		if v_3_0 == 902936544 {
			__t1 = x_1
			goto end_branch_1
		} else {

		}
	}
	{
		if v_3_0 == 380165415 {
			__t1 = x_1
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Data_Ord_max__4009478382(__eta0_0_loop int64, __eta1_1_loop int64) int64 {
max__4009478382:
	for {
		if false {
			continue max__4009478382
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
		v_2_0 := uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
		_ = v_2_0
		var __t1 int64
		{
			if v_2_0 == 1527465420 {
				__t1 = __eta1_1
				goto end_branch_1
			} else {

			}
		}
		{
			if v_2_0 == 902936544 {
				__t1 = __eta0_0
				goto end_branch_1
			} else {

			}
		}
		{
			if v_2_0 == 380165415 {
				__t1 = __eta0_0
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() int64 { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_Ord_min(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var y_2 gopurs_runtime.Value = y_2_loop
	_ = y_2
	// TAST (Let): v_3_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
	v_3_0 := uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_1, y_2).IntVal)
	_ = v_3_0
	var __t1 gopurs_runtime.Value
	{
		if v_3_0 == 1527465420 {
			__t1 = x_1
			goto end_branch_1
		} else {

		}
	}
	{
		if v_3_0 == 902936544 {
			__t1 = x_1
			goto end_branch_1
		} else {

		}
	}
	{
		if v_3_0 == 380165415 {
			__t1 = y_2
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Data_Ord_min__4009478382(__eta0_0_loop int64, __eta1_1_loop int64) int64 {
min__4009478382:
	for {
		if false {
			continue min__4009478382
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
		v_2_0 := uint32(gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(__eta0_0), gopurs_runtime.Int(__eta1_1)).IntVal)
		_ = v_2_0
		var __t1 int64
		{
			if v_2_0 == 1527465420 {
				__t1 = __eta0_0
				goto end_branch_1
			} else {

			}
		}
		{
			if v_2_0 == 902936544 {
				__t1 = __eta0_0
				goto end_branch_1
			} else {

			}
		}
		{
			if v_2_0 == 380165415 {
				__t1 = __eta1_1
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() int64 { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Data_Ord_ordArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqArray_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(Array (TypeVar a$scope50))])
	eqArray_1_0 := Rebox_Data_Ord_3790796878_1939691112(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))))
	_ = eqArray_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_4088624520_4177771502((&Constructor_Data_Ord_Ord[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_1939691112_3790796878(eqArray_1_0))}
	}), gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Rebox_Data_Ord_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Data_Ord_ordArrayImpl(), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_6_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
			v_6_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, y_5).IntVal)
			_ = v_6_1
			var __t2 gopurs_runtime.Value
			{
				if v_6_1 == 902936544 {
					__t2 = gopurs_runtime.Int(int64(0))
					goto end_branch_2
				} else {

				}
			}
			{
				if v_6_1 == 1527465420 {
					__t2 = gopurs_runtime.Int(int64(1))
					goto end_branch_2
				} else {

				}
			}
			{
				if v_6_1 == 380165415 {
					__t2 = gopurs_runtime.Int(int64(-1))
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
			}
		end_branch_2:
			return __t2
		}), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(xs_2.UnsafePtr))), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(ys_3.UnsafePtr)))).IntVal))
	})})))}
}

func Call_Data_Ord_ordRecordCons(dictOrdRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrdRecord_0 gopurs_runtime.Value = dictOrdRecord_0_loop
	_ = dictOrdRecord_0
	// TAST (Let): eqRowCons_1_0 shape=App(Var) bindingType=Any
	eqRowCons_1_0 := gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), gopurs_runtime.Value{})
	_ = eqRowCons_1_0
	return gopurs_runtime.Func2(func(_dollar___unused_2 gopurs_runtime.Value, dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eqRowCons1_4_1 shape=App(Other) bindingType=Any
		eqRowCons1_4_1 := gopurs_runtime.Apply(eqRowCons_1_0, dictIsSymbol_3)
		_ = eqRowCons1_4_1
		return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): eqRowCons2_6_2 shape=App(Other) bindingType=(TypeApp (ADT ["Data","Eq","EqRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key$scope60), (TypeVar focus$scope61), (TypeVar rowlistTail$scope58)]), (TypeVar row$scope59)])
			eqRowCons2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(eqRowCons1_4_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})))
			_ = eqRowCons2_6_2
			return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(eqRowCons2_6_2)}
			}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, ra_8 gopurs_runtime.Value, rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): key_10_3 shape=App(Other) bindingType=String
				key_10_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_3, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
				_ = key_10_3
				// TAST (Let): left_11_4 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
				left_11_4 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_5, "compare"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_10_3), ra_8), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_10_3), rb_9)).IntVal)
				_ = left_11_4
				var __t5 gopurs_runtime.Value
				{
					if (left_11_4 == 902936544) != (true) {
						__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(left_11_4), UnsafePtr: nil}
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9)
				}
			end_branch_5:
				return __t5
			})}))}
		})
	})
}

func Call_Data_Ord_clamp(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var low_1 gopurs_runtime.Value = low_1_loop
	_ = low_1
	var hi_2 gopurs_runtime.Value = hi_2_loop
	_ = hi_2
	var x_3 gopurs_runtime.Value = x_3_loop
	_ = x_3
	// TAST (Let): v_4_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
	v_4_1 := uint32(gopurs_runtime.Apply2(dictOrd_0.V1, low_1, x_3).IntVal)
	_ = v_4_1
	var __t2 gopurs_runtime.Value
	{
		if v_4_1 == 1527465420 {
			__t2 = x_3
			goto end_branch_2
		} else {

		}
	}
	{
		if v_4_1 == 902936544 {
			__t2 = low_1
			goto end_branch_2
		} else {

		}
	}
	{
		if v_4_1 == 380165415 {
			__t2 = low_1
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_2:
	// TAST (Let): __local_var_4_0 shape=Let(Branch(Other, Other, Other, def=Other)) bindingType=(TypeVar a$scope73)
	__local_var_4_0 := __t2
	_ = __local_var_4_0
	// TAST (Let): v_5_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
	v_5_3 := uint32(gopurs_runtime.Apply2(dictOrd_0.V1, hi_2, __local_var_4_0).IntVal)
	_ = v_5_3
	var __t4 gopurs_runtime.Value
	{
		if v_5_3 == 1527465420 {
			__t4 = hi_2
			goto end_branch_4
		} else {

		}
	}
	{
		if v_5_3 == 902936544 {
			__t4 = hi_2
			goto end_branch_4
		} else {

		}
	}
	{
		if v_5_3 == 380165415 {
			__t4 = __local_var_4_0
			goto end_branch_4
		} else {

		}
	}
	{
		__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_4:
	return __t4
}

func Call_Data_Ord_clamp__1405043522(__eta0_0_loop int64, __eta1_1_loop int64, __eta2_2_loop int64) int64 {
clamp__1405043522:
	for {
		if false {
			continue clamp__1405043522
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		var __eta2_2 int64 = __eta2_2_loop
		_ = __eta2_2
		return Call_Data_Ord_min__4009478382(__eta1_1, Call_Data_Ord_max__4009478382(__eta0_0, __eta2_2))
	}
}

func Call_Data_Ord_between(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], low_1_loop gopurs_runtime.Value, hi_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) bool {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var low_1 gopurs_runtime.Value = low_1_loop
	_ = low_1
	var hi_2 gopurs_runtime.Value = hi_2_loop
	_ = hi_2
	var x_3 gopurs_runtime.Value = x_3_loop
	_ = x_3
	var __t2 bool
	{
		var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_3, low_1)
		_ = __t_tag_1
		if uint32(__t_tag_1.IntVal) == 1527465420 {
			__t2 = false
			goto end_branch_2
		} else {

		}
	}
	{
		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_3, hi_2)
		_ = __t_tag_0
		__t2 = (uint32(__t_tag_0.IntVal) == 380165415) != (true)
	}
end_branch_2:
	return __t2
}

func Call_Data_Ord_abs(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], dictRing_1_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var dictRing_1 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dictRing_1_loop
	_ = dictRing_1
	// TAST (Let): zero_2_0 shape=App(Var) bindingType=(TypeVar a$scope77)
	zero_2_0 := Call_Data_Semiring_zero(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}))
	_ = zero_2_0
	// TAST (Let): zero_3_1 shape=Other bindingType=(TypeVar a$scope77)
	zero_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_1.V0, gopurs_runtime.Value{}), "zero")
	_ = zero_3_1
	return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 gopurs_runtime.Value
		{
			var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, x_4, zero_2_0)
			_ = __t_tag_2
			if (uint32(__t_tag_2.IntVal) == 1527465420) != (true) {
				__t3 = x_4
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.Apply2(dictRing_1.V1, zero_3_1, x_4)
		}
	end_branch_3:
		return __t3
	})
}

func Rebox_Data_Ord_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Ord_1888627657_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_1939691112_3790796878(in *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Ord_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3383090825_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3790796878_1939691112(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_3790796878_687527510(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_4088624520_4177771502(in *Constructor_Data_Ord_Ord[[]gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_4177771502_3047586294(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_4177771502_3383090825(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_4177771502_758368489(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ord_758368489_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Data_Ord_ordArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdArrayImpl
}

func Get_Data_Ord_ordBooleanImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdBooleanImpl
}

func Get_Data_Ord_ordCharImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdCharImpl
}

func Get_Data_Ord_ordIntImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdIntImpl
}

func Get_Data_Ord_ordNumberImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdNumberImpl
}

func Get_Data_Ord_ordStringImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Ord_OrdStringImpl
}
