package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_point gopurs_runtime.Value
var once_Main_point sync.Once

func Get_Main_point() gopurs_runtime.Value {
	once_Main_point.Do(func() {
		cache_Main_point = func() gopurs_runtime.Value {
			orig := struct {
				x float64
				y float64
			}{1.0, 0.0}
			_ = orig
			return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y))
		}()
	})
	return cache_Main_point
}

var cache_Main_getX gopurs_runtime.Value
var once_Main_getX sync.Once

func Get_Main_getX() gopurs_runtime.Value {
	once_Main_getX.Do(func() {
		cache_Main_getX = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_getX(v_0_box)
		})
	})
	return cache_Main_getX
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))), gopurs_runtime.Float(Call_Main_getX(func() gopurs_runtime.Value {
			orig := func() struct {
				x float64
				y float64
			} {
				orig := Get_Main_point()
				_ = orig
				clone := struct {
					x float64
					y float64
				}{}
				clone.x = gopurs_runtime.RecordGet(orig, "x").FloatVal()
				clone.y = gopurs_runtime.RecordGet(orig, "y").FloatVal()
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y))
		}()).FloatVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("OK")), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						orig := func() struct {
							y string
						} {
							orig := gopurs_runtime.RecordGet(v_2, "x")
							_ = orig
							clone := struct {
								y string
							}{}
							clone.y = gopurs_runtime.RecordGet(orig, "y").StrVal()
							return clone
						}()
						_ = orig
						return gopurs_runtime.RecordDict1("y", gopurs_runtime.Str(orig.y))
					}()
				}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str(gopurs_runtime.RecordGet(v_2, "y").StrVal())
				})), func() gopurs_runtime.Value {
					orig := struct {
						x struct {
							y string
						}
					}{struct {
						y string
					}{"Nested"}}
					_ = orig
					return gopurs_runtime.RecordDict1("x", func() gopurs_runtime.Value {
						orig := orig.x
						_ = orig
						return gopurs_runtime.RecordDict1("y", gopurs_runtime.Str(orig.y))
					}())
				}()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
				}))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main_getX(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.RecordGet(v_0, "x")
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[float64]{}
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
