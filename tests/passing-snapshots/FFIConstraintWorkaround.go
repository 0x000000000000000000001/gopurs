package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3574847586_1386611502(Rebox_Main_1386611502_3574847586(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		})), Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("b")
		})), Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("c")
		})), Call_Data_Show_showRecordFieldsConsNil(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("e")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showChar()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main_showRecord
}

var cache_Main_showFFI gopurs_runtime.Value
var once_Main_showFFI sync.Once

func Get_Main_showFFI() gopurs_runtime.Value {
	once_Main_showFFI.Do(func() {
		cache_Main_showFFI = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showFFI(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box))
		})
	})
	return cache_Main_showFFI
}

var cache_Main_showFFI__1878234284 gopurs_runtime.Value
var once_Main_showFFI__1878234284 sync.Once

func Get_Main_showFFI__1878234284() gopurs_runtime.Value {
	once_Main_showFFI__1878234284.Do(func() {
		cache_Main_showFFI__1878234284 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_showFFI__1878234284(__eta_norm_0_0_box.IntVal))
		})
	})
	return cache_Main_showFFI__1878234284
}

var cache_Main_showFFI__460344716 gopurs_runtime.Value
var once_Main_showFFI__460344716 sync.Once

func Get_Main_showFFI__460344716() gopurs_runtime.Value {
	once_Main_showFFI__460344716.Do(func() {
		cache_Main_showFFI__460344716 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_showFFI__460344716(func() struct {
				a int64
				b bool
				c string
				e float64
			} {
				orig := __eta_norm_0_0_box
				_ = orig
				clone := struct {
					a int64
					b bool
					c string
					e float64
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
				clone.b = (gopurs_runtime.RecordGet(orig, "b").IntVal) != (0)
				clone.c = gopurs_runtime.RecordGet(orig, "c").StrVal()
				clone.e = gopurs_runtime.RecordGet(orig, "e").FloatVal()
				return clone
			}()))
		})
	})
	return cache_Main_showFFI__460344716
}

var cache_Main_showFFI__2030676012 gopurs_runtime.Value
var once_Main_showFFI__2030676012 sync.Once

func Get_Main_showFFI__2030676012() gopurs_runtime.Value {
	once_Main_showFFI__2030676012.Do(func() {
		cache_Main_showFFI__2030676012 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_showFFI__2030676012(__eta_norm_0_0_box.StrVal()))
		})
	})
	return cache_Main_showFFI__2030676012
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_showFFI__460344716(struct {
			a int64
			b bool
			c string
			e float64
		}{int64(1), true, "d", 4.0}))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Showing Int is correct"), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Main_showImpl(), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), gopurs_runtime.Int(int64(4))).StrVal()) == ("4"))), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Showing String is correct"), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Main_showImpl(), Call_Data_Show_show(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), gopurs_runtime.Str("string")).StrVal()) == ("\"string\""))), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Showing Record is correct"), gopurs_runtime.Bool((Call_Main_showFFI__460344716(struct {
						a int64
						b bool
						c string
						e float64
					}{int64(1), true, "d", 4.0})) == ("{ a: 1, b: true, c: 'd', e: 4.0 }"))), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main_showFFI(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Apply(Get_Main_showImpl(), Call_Data_Show_show(dictShow_0))
}

func Call_Main_showFFI__1878234284(__eta_norm_0_0_loop int64) string {
showFFI__1878234284:
	for {
		if false {
			continue showFFI__1878234284
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply2(Get_Main_showImpl(), Call_Data_Show_show(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), gopurs_runtime.Int(__eta_norm_0_0)).StrVal()
	}
}

func Call_Main_showFFI__460344716(__eta_norm_0_0_loop struct {
	a int64
	b bool
	c string
	e float64
}) string {
showFFI__460344716:
	for {
		if false {
			continue showFFI__460344716
		}
		var __eta_norm_0_0 struct {
			a int64
			b bool
			c string
			e float64
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply2(Get_Main_showImpl(), Call_Data_Show_show(Rebox_Main_3574847586_1386611502(Rebox_Main_1386611502_3574847586(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		})), Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("b")
		})), Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("c")
		})), Call_Data_Show_showRecordFieldsConsNil(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("e")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showChar()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))), func() gopurs_runtime.Value {
			orig := __eta_norm_0_0
			_ = orig
			return gopurs_runtime.RecordDict4("a", "b", "c", "e", gopurs_runtime.Int(orig.a), gopurs_runtime.Bool(orig.b), gopurs_runtime.Str(orig.c), gopurs_runtime.Float(orig.e))
		}()).StrVal()
	}
}

func Call_Main_showFFI__2030676012(__eta_norm_0_0_loop string) string {
showFFI__2030676012:
	for {
		if false {
			continue showFFI__2030676012
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply2(Get_Main_showImpl(), Call_Data_Show_show(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), gopurs_runtime.Str(__eta_norm_0_0)).StrVal()
	}
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_3574847586(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	a int64
	b bool
	c string
	e float64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		a int64
		b bool
		c string
		e float64
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3574847586_1386611502(in *Constructor_Data_Show_Show[struct {
	a int64
	b bool
	c string
	e float64
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Get_Main_showImpl() gopurs_runtime.Value {
	return _Gopurs_Main_ShowImpl
}
