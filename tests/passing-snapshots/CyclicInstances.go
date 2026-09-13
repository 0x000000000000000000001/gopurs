package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_BIsSymbol gopurs_runtime.Value
var once_Main_BIsSymbol sync.Once

func Get_Main_BIsSymbol() gopurs_runtime.Value {
	once_Main_BIsSymbol.Do(func() {
		cache_Main_BIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("B")
		}))
	})
	return cache_Main_BIsSymbol
}

var cache_Main_genericShowConstructor gopurs_runtime.Value
var once_Main_genericShowConstructor sync.Once

func Get_Main_genericShowConstructor() gopurs_runtime.Value {
	once_Main_genericShowConstructor.Do(func() {
		cache_Main_genericShowConstructor = gopurs_runtime.Apply(Get_Data_Show_Generic_genericShowConstructor(), gopurs_runtime.Value{Type: 9, IntVal: 1968625250, UnsafePtr: unsafe.Pointer(Rebox_Main_127695515_242610358(Rebox_Main_242610358_127695515(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShowArgs[gopurs_runtime.Value]](Get_Data_Show_Generic_genericShowArgsNoArguments()))))})
	})
	return cache_Main_genericShowConstructor
}

var cache_Main_genericShowConstructor1 gopurs_runtime.Value
var once_Main_genericShowConstructor1 sync.Once

func Get_Main_genericShowConstructor1() gopurs_runtime.Value {
	once_Main_genericShowConstructor1.Do(func() {
		cache_Main_genericShowConstructor1 = Call_Data_Show_Generic_genericShowConstructor(gopurs_runtime.Value{Type: 9, IntVal: 1968625250, UnsafePtr: unsafe.Pointer(Rebox_Main_127695515_242610358(Rebox_Main_242610358_127695515(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShowArgs[gopurs_runtime.Value]](Get_Data_Show_Generic_genericShowArgsNoArguments()))))}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Z")
		})))
	})
	return cache_Main_genericShowConstructor1
}

var cache_Main_B2IsSymbol gopurs_runtime.Value
var once_Main_B2IsSymbol sync.Once

func Get_Main_B2IsSymbol() gopurs_runtime.Value {
	once_Main_B2IsSymbol.Do(func() {
		cache_Main_B2IsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("B2")
		}))
	})
	return cache_Main_B2IsSymbol
}

var cache_Main_genericShowConstructor2 gopurs_runtime.Value
var once_Main_genericShowConstructor2 sync.Once

func Get_Main_genericShowConstructor2() gopurs_runtime.Value {
	once_Main_genericShowConstructor2.Do(func() {
		cache_Main_genericShowConstructor2 = Call_Data_Show_Generic_genericShowConstructor(gopurs_runtime.Value{Type: 9, IntVal: 1968625250, UnsafePtr: unsafe.Pointer(Rebox_Main_127695515_242610358(Rebox_Main_242610358_127695515(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Generic_GenericShowArgs[gopurs_runtime.Value]](Get_Data_Show_Generic_genericShowArgsNoArguments()))))}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Z2")
		})))
	})
	return cache_Main_genericShowConstructor2
}

var cache_Main_A2 gopurs_runtime.Value
var once_Main_A2 sync.Once

func Get_Main_A2() gopurs_runtime.Value {
	once_Main_A2.Do(func() {
		cache_Main_A2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_A2(func() struct {
					x *Constructor_Main_B2
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						x *Constructor_Main_B2
					}{}
					clone.x = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](gopurs_runtime.RecordGet(orig, "x"))
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict1("x", gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer(orig.x)})
			}()
		})
	})
	return cache_Main_A2
}

var cache_Main_B2 gopurs_runtime.Value
var once_Main_B2 sync.Once

func Get_Main_B2() gopurs_runtime.Value {
	once_Main_B2.Do(func() {
		cache_Main_B2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer((&Constructor_Main_B2{1, func() struct {
				x *Constructor_Main_B2
			} {
				orig := value0
				_ = orig
				clone := struct {
					x *Constructor_Main_B2
				}{}
				clone.x = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](gopurs_runtime.RecordGet(orig, "x"))
				return clone
			}()}))}
		})
	})
	return cache_Main_B2
}

var cache_Main_Z2 gopurs_runtime.Value
var once_Main_Z2 sync.Once

func Get_Main_Z2() gopurs_runtime.Value {
	once_Main_Z2.Do(func() {
		cache_Main_Z2 = gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer((*Constructor_Main_B2)(nil))}
	})
	return cache_Main_Z2
}

var cache_Main_C2 gopurs_runtime.Value
var once_Main_C2 sync.Once

func Get_Main_C2() gopurs_runtime.Value {
	once_Main_C2.Do(func() {
		cache_Main_C2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_C2(func() struct {
					x *Constructor_Main_B2
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						x *Constructor_Main_B2
					}{}
					clone.x = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](gopurs_runtime.RecordGet(orig, "x"))
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict1("x", gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer(orig.x)})
			}()
		})
	})
	return cache_Main_C2
}

var cache_Main_A gopurs_runtime.Value
var once_Main_A sync.Once

func Get_Main_A() gopurs_runtime.Value {
	once_Main_A.Do(func() {
		cache_Main_A = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(Call_Main_A(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0_box)))}
		})
	})
	return cache_Main_A
}

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((&Constructor_Main_B{1, gopurs_runtime.CoerceToStruct[Constructor_Main_B](value0)}))}
		})
	})
	return cache_Main_B
}

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((*Constructor_Main_B)(nil))}
	})
	return cache_Main_Z
}

var cache_Main_C gopurs_runtime.Value
var once_Main_C sync.Once

func Get_Main_C() gopurs_runtime.Value {
	once_Main_C.Do(func() {
		cache_Main_C = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(Call_Main_C(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0_box)))}
		})
	})
	return cache_Main_C
}

var cache_Main_genericC_ gopurs_runtime.Value
var once_Main_genericC_ sync.Once

func Get_Main_genericC_() gopurs_runtime.Value {
	once_Main_genericC_.Do(func() {
		cache_Main_genericC_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_1009504592_2818661616((&Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B, *Constructor_Main_B]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0))}
		})})))}
	})
	return cache_Main_genericC_
}

var cache_Main_genericC2_ gopurs_runtime.Value
var once_Main_genericC2_ sync.Once

func Get_Main_genericC2_() gopurs_runtime.Value {
	once_Main_genericC2_.Do(func() {
		cache_Main_genericC2_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_1291651792_2818661616((&Constructor_Data_Generic_Rep_Generic[struct {
			x *Constructor_Main_B2
		}, struct {
			x *Constructor_Main_B2
		}]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := func() struct {
					x *Constructor_Main_B2
				} {
					orig := x_0
					_ = orig
					clone := struct {
						x *Constructor_Main_B2
					}{}
					clone.x = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](gopurs_runtime.RecordGet(orig, "x"))
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict1("x", gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer(orig.x)})
			}()
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := func() struct {
					x *Constructor_Main_B2
				} {
					orig := x_0
					_ = orig
					clone := struct {
						x *Constructor_Main_B2
					}{}
					clone.x = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](gopurs_runtime.RecordGet(orig, "x"))
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict1("x", gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer(orig.x)})
			}()
		})})))}
	})
	return cache_Main_genericC2_
}

var cache_Main_genericB_ gopurs_runtime.Value
var once_Main_genericB_ sync.Once

func Get_Main_genericB_() gopurs_runtime.Value {
	once_Main_genericB_.Do(func() {
		cache_Main_genericB_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_2147772913_2818661616((&Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				var __t_tag_0 *Constructor_Main_B = gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0)
				_ = __t_tag_0
				if __t_tag_0 != nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_3895662118_1323331594((&Constructor_Data_Generic_Rep_Inl[*Constructor_Main_B, uint32]{1, (*Constructor_Main_B)(x_0.UnsafePtr).V0})))}
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 *Constructor_Main_B = gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0)
				_ = __t_tag_1
				if __t_tag_1 == nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_3658796408_2687169876((&Constructor_Data_Generic_Rep_Inr[*Constructor_Main_B, uint32]{1, 1454898258})))}
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 *Constructor_Main_B
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t3 = (&Constructor_Main_B{1, gopurs_runtime.CoerceToStruct[Constructor_Main_B]((*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0)})
					goto end_branch_3
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t3 = (*Constructor_Main_B)(nil)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() *Constructor_Main_B { panic("Failed pattern match") }()
			}
		end_branch_3:
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(__t3)}
		})})))}
	})
	return cache_Main_genericB_
}

var cache_Main_showB gopurs_runtime.Value
var once_Main_showB sync.Once

func Get_Main_showB() gopurs_runtime.Value {
	once_Main_showB.Do(func() {
		cache_Main_showB = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_576757679_1386611502((&Constructor_Data_Show_Show[*Constructor_Main_B]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Show_Generic_genericShow__1161896515(gopurs_runtime.CoerceToStruct[Constructor_Main_B](x_0)))
		})})))}
	})
	return cache_Main_showB
}

var cache_Main_showA gopurs_runtime.Value
var once_Main_showA sync.Once

func Get_Main_showA() gopurs_runtime.Value {
	once_Main_showA.Do(func() {
		cache_Main_showA = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_576757679_1386611502(Rebox_Main_1386611502_576757679(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showB()))))}
	})
	return cache_Main_showA
}

var cache_Main_showC gopurs_runtime.Value
var once_Main_showC sync.Once

func Get_Main_showC() gopurs_runtime.Value {
	once_Main_showC.Do(func() {
		cache_Main_showC = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_576757679_1386611502((&Constructor_Data_Show_Show[*Constructor_Main_B]{1, gopurs_runtime.Apply2(Get_Data_Show_Generic_genericShow(), gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_1009504592_2818661616(Rebox_Main_2818661616_1009504592(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_genericC_()))))}, Call_Data_Show_Generic_genericShowConstructor(Call_Data_Show_Generic_genericShowArgsArgument(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_576757679_1386611502(Rebox_Main_1386611502_576757679(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showA()))))}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("C")
		}))))})))}
	})
	return cache_Main_showC
}

var cache_Main_genericB2_ gopurs_runtime.Value
var once_Main_genericB2_ sync.Once

func Get_Main_genericB2_() gopurs_runtime.Value {
	once_Main_genericB2_.Do(func() {
		cache_Main_genericB2_ = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_3587705155_2818661616((&Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B2, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				var __t_tag_0 *Constructor_Main_B2 = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](x_0)
				_ = __t_tag_0
				if __t_tag_0 != nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_615526420_1323331594((&Constructor_Data_Generic_Rep_Inl[struct {
						x *Constructor_Main_B2
					}, uint32]{1, (*Constructor_Main_B2)(x_0.UnsafePtr).V0})))}
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 *Constructor_Main_B2 = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](x_0)
				_ = __t_tag_1
				if __t_tag_1 == nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_3485313226_2687169876((&Constructor_Data_Generic_Rep_Inr[struct {
						x *Constructor_Main_B2
					}, uint32]{1, 1454898258})))}
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return __t2
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 *Constructor_Main_B2
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t3 = (&Constructor_Main_B2{1, func() struct {
						x *Constructor_Main_B2
					} {
						orig := (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0
						_ = orig
						clone := struct {
							x *Constructor_Main_B2
						}{}
						clone.x = gopurs_runtime.CoerceToStruct[Constructor_Main_B2](gopurs_runtime.RecordGet(orig, "x"))
						return clone
					}()})
					goto end_branch_3
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					__t3 = (*Constructor_Main_B2)(nil)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() *Constructor_Main_B2 { panic("Failed pattern match") }()
			}
		end_branch_3:
			return gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer(__t3)}
		})})))}
	})
	return cache_Main_genericB2_
}

var cache_Main_showB2 gopurs_runtime.Value
var once_Main_showB2 sync.Once

func Get_Main_showB2() gopurs_runtime.Value {
	once_Main_showB2.Do(func() {
		cache_Main_showB2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2461310173_1386611502((&Constructor_Data_Show_Show[*Constructor_Main_B2]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Show_Generic_genericShow__3507161152(gopurs_runtime.CoerceToStruct[Constructor_Main_B2](x_0)))
		})})))}
	})
	return cache_Main_showB2
}

var cache_Main_showA2 gopurs_runtime.Value
var once_Main_showA2 sync.Once

func Get_Main_showA2() gopurs_runtime.Value {
	once_Main_showA2.Do(func() {
		cache_Main_showA2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_4853725_1386611502(Rebox_Main_1386611502_4853725(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("x")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2461310173_1386611502(Rebox_Main_1386611502_2461310173(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showB2()))))}))))))}
	})
	return cache_Main_showA2
}

var cache_Main_showC2 gopurs_runtime.Value
var once_Main_showC2 sync.Once

func Get_Main_showC2() gopurs_runtime.Value {
	once_Main_showC2.Do(func() {
		cache_Main_showC2 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_4853725_1386611502((&Constructor_Data_Show_Show[struct {
			x *Constructor_Main_B2
		}]{1, gopurs_runtime.Apply2(Get_Data_Show_Generic_genericShow(), gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_1291651792_2818661616(Rebox_Main_2818661616_1291651792(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Main_genericC2_()))))}, Call_Data_Show_Generic_genericShowConstructor(Call_Data_Show_Generic_genericShowArgsArgument(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_4853725_1386611502(Rebox_Main_1386611502_4853725(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main_showA2()))))}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("C2")
		}))))})))}
	})
	return cache_Main_showC2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_B2 struct {
	Rc uint32
	V0 struct {
		x *Constructor_Main_B2
	}
}

type Constructor_Main_Z2 struct {
	Rc uint32
}

type Constructor_Main_B struct {
	Rc uint32
	V0 *Constructor_Main_B
}

type Constructor_Main_Z struct {
	Rc uint32
}

func Call_Main_A2(x_0_loop struct {
	x *Constructor_Main_B2
}) struct {
	x *Constructor_Main_B2
} {
	var x_0 struct {
		x *Constructor_Main_B2
	} = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C2(x_0_loop struct {
	x *Constructor_Main_B2
}) struct {
	x *Constructor_Main_B2
} {
	var x_0 struct {
		x *Constructor_Main_B2
	} = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_A(x_0_loop *Constructor_Main_B) *Constructor_Main_B {
	var x_0 *Constructor_Main_B = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C(x_0_loop *Constructor_Main_B) *Constructor_Main_B {
	var x_0 *Constructor_Main_B = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Main_1009504592_2818661616(in *Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B, *Constructor_Main_B]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_127695515_242610358(in *Constructor_Data_Show_Generic_GenericShowArgs[uint32]) *Constructor_Data_Show_Generic_GenericShowArgs[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Generic_GenericShowArgs[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1291651792_2818661616(in *Constructor_Data_Generic_Rep_Generic[struct {
	x *Constructor_Main_B2
}, struct {
	x *Constructor_Main_B2
}]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_1386611502_2461310173(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Main_B2] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[*Constructor_Main_B2]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_4853725(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	x *Constructor_Main_B2
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		x *Constructor_Main_B2
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_576757679(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Main_B] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[*Constructor_Main_B]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2147772913_2818661616(in *Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_242610358_127695515(in *Constructor_Data_Show_Generic_GenericShowArgs[gopurs_runtime.Value]) *Constructor_Data_Show_Generic_GenericShowArgs[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Generic_GenericShowArgs[uint32]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2461310173_1386611502(in *Constructor_Data_Show_Show[*Constructor_Main_B2]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2818661616_1009504592(in *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B, *Constructor_Main_B] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B, *Constructor_Main_B]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2818661616_1291651792(in *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[struct {
	x *Constructor_Main_B2
}, struct {
	x *Constructor_Main_B2
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Generic[struct {
		x *Constructor_Main_B2
	}, struct {
		x *Constructor_Main_B2
	}]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3485313226_2687169876(in *Constructor_Data_Generic_Rep_Inr[struct {
	x *Constructor_Main_B2
}, uint32]) *Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Main_3587705155_2818661616(in *Constructor_Data_Generic_Rep_Generic[*Constructor_Main_B2, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3658796408_2687169876(in *Constructor_Data_Generic_Rep_Inr[*Constructor_Main_B, uint32]) *Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Main_3895662118_1323331594(in *Constructor_Data_Generic_Rep_Inl[*Constructor_Main_B, uint32]) *Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Main_4853725_1386611502(in *Constructor_Data_Show_Show[struct {
	x *Constructor_Main_B2
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_576757679_1386611502(in *Constructor_Data_Show_Show[*Constructor_Main_B]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_615526420_1323331594(in *Constructor_Data_Generic_Rep_Inl[struct {
	x *Constructor_Main_B2
}, uint32]) *Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict1("x", gopurs_runtime.Value{Type: 9, IntVal: 4102534158, UnsafePtr: unsafe.Pointer(orig.x)})
	}()
	return out
}
