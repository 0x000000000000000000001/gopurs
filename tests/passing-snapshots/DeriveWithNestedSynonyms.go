package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Z
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Y
}

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_X
}

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_T(func() struct {
					baz struct {
						foo string
						/*gopurs:record-v1:1:0066006f006f*/
					}
					/*gopurs:record-v1:1:00620061007a*/
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						baz struct {
							foo string
							/*gopurs:record-v1:1:0066006f006f*/
						}
						/*gopurs:record-v1:1:00620061007a*/
					}{}
					clone.baz = func() struct {
						foo string
						/*gopurs:record-v1:1:0066006f006f*/
					} {
						orig := gopurs_runtime.RecordGet(orig, "baz")
						_ = orig
						clone := struct {
							foo string
							/*gopurs:record-v1:1:0066006f006f*/
						}{}
						clone.foo = gopurs_runtime.RecordGet(orig, "foo").StrVal()
						return clone
					}()
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict1("baz", func() gopurs_runtime.Value {
					orig := orig.baz
					_ = orig
					return gopurs_runtime.RecordDict1("foo", gopurs_runtime.Str(orig.foo))
				}())
			}()
		})
	})
	return cache_Main_T
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqZ gopurs_runtime.Value
var once_Main_eqZ sync.Once

func Get_Main_eqZ() gopurs_runtime.Value {
	once_Main_eqZ.Do(func() {
		cache_Main_eqZ = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Main_eqZ
}

var cache_Main_eqY gopurs_runtime.Value
var once_Main_eqY sync.Once

func Get_Main_eqY() gopurs_runtime.Value {
	once_Main_eqY.Do(func() {
		cache_Main_eqY = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Main_eqY
}

var cache_Main_eqX gopurs_runtime.Value
var once_Main_eqX sync.Once

func Get_Main_eqX() gopurs_runtime.Value {
	once_Main_eqX.Do(func() {
		cache_Main_eqX = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Main_eqX
}

var cache_Main_eqT gopurs_runtime.Value
var once_Main_eqT sync.Once

func Get_Main_eqT() gopurs_runtime.Value {
	once_Main_eqT.Do(func() {
		cache_Main_eqT = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3002535950_3790796878((&Constructor_Data_Eq_Eq[struct {
			baz struct {
				foo string
				/*gopurs:record-v1:1:0066006f006f*/
			}
			/*gopurs:record-v1:1:00620061007a*/
		}]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "baz"), "foo").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "baz"), "foo").StrVal()))
		})})))}
	})
	return cache_Main_eqT
}

var cache_Main_ordT gopurs_runtime.Value
var once_Main_ordT sync.Once

func Get_Main_ordT() gopurs_runtime.Value {
	once_Main_ordT.Do(func() {
		cache_Main_ordT = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_2553746606_4177771502((&Constructor_Data_Ord_Ord[struct {
			baz struct {
				foo string
				/*gopurs:record-v1:1:0066006f006f*/
			}
			/*gopurs:record-v1:1:00620061007a*/
		}]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3002535950_3790796878(Rebox_Main_3790796878_3002535950(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqT()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "baz"), "foo"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "baz"), "foo"))
		})})))}
	})
	return cache_Main_ordT
}

type Constructor_Main_Z struct {
	Rc uint32
	V0 struct {

		/*gopurs:record-v1:0:*/
	}
}

type Constructor_Main_Y struct {
	Rc uint32
	V0 struct {
		foo struct {

			/*gopurs:record-v1:0:*/
		}
		/*gopurs:record-v1:1:0066006f006f*/
	}
}

type Constructor_Main_X struct {
	Rc uint32
	V0 struct {

		/*gopurs:record-v1:0:*/
	}
}

func Call_Main_T(x_0_loop struct {
	baz struct {
		foo string
		/*gopurs:record-v1:1:0066006f006f*/
	}
	/*gopurs:record-v1:1:00620061007a*/
}) struct {
	baz struct {
		foo string
		/*gopurs:record-v1:1:0066006f006f*/
	}
	/*gopurs:record-v1:1:00620061007a*/
} {
	var x_0 struct {
		baz struct {
			foo string
			/*gopurs:record-v1:1:0066006f006f*/
		}
		/*gopurs:record-v1:1:00620061007a*/
	} = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Main_2553746606_4177771502(in *Constructor_Data_Ord_Ord[struct {
	baz struct {
		foo string
		/*gopurs:record-v1:1:0066006f006f*/
	}
	/*gopurs:record-v1:1:00620061007a*/
}]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3002535950_3790796878(in *Constructor_Data_Eq_Eq[struct {
	baz struct {
		foo string
		/*gopurs:record-v1:1:0066006f006f*/
	}
	/*gopurs:record-v1:1:00620061007a*/
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_3002535950(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	baz struct {
		foo string
		/*gopurs:record-v1:1:0066006f006f*/
	}
	/*gopurs:record-v1:1:00620061007a*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[struct {
		baz struct {
			foo string
			/*gopurs:record-v1:1:0066006f006f*/
		}
		/*gopurs:record-v1:1:00620061007a*/
	}])(unsafe.Pointer(in))
}
