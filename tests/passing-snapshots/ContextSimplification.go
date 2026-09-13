package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_shout gopurs_runtime.Value
var once_Main_shout sync.Once

func Get_Main_shout() gopurs_runtime.Value {
	once_Main_shout.Do(func() {
		cache_Main_shout = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_shout(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box))
		})
	})
	return cache_Main_shout
}

var cache_Main_usesShowTwice gopurs_runtime.Value
var once_Main_usesShowTwice sync.Once

func Get_Main_usesShowTwice() gopurs_runtime.Value {
	once_Main_usesShowTwice.Do(func() {
		cache_Main_usesShowTwice = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_usesShowTwice(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), (v_1_box.IntVal) != (0))
		})
	})
	return cache_Main_usesShowTwice
}

var cache_Main_usesShowTwice__2069622520 gopurs_runtime.Value
var once_Main_usesShowTwice__2069622520 sync.Once

func Get_Main_usesShowTwice__2069622520() gopurs_runtime.Value {
	once_Main_usesShowTwice__2069622520.Do(func() {
		cache_Main_usesShowTwice__2069622520 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_usesShowTwice__2069622520((v_0_box.IntVal) != (0), __eta_norm_0_1_box.StrVal())
		})
	})
	return cache_Main_usesShowTwice__2069622520
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Call_Main_shout(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), gopurs_runtime.Str("Test")), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_shout(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Console_log(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str((v_1.StrVal()) + ("!"))
	}), Call_Data_Show_show(dictShow_0)))
}

func Call_Main_usesShowTwice(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], v_1_loop bool) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var v_1 bool = v_1_loop
	_ = v_1
	var __t0 gopurs_runtime.Value
	{
		if v_1 {
			__t0 = Call_Main_shout(dictShow_0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply(Get_Effect_Console_logShow(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(dictShow_0)})
	}
end_branch_0:
	return __t0
}

func Call_Main_usesShowTwice__2069622520(v_0_loop bool, __eta_norm_0_1_loop string) gopurs_runtime.Value {
usesShowTwice__2069622520:
	for {
		if false {
			continue usesShowTwice__2069622520
		}
		var v_0 bool = v_0_loop
		_ = v_0
		var __eta_norm_0_1 string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var __t0 gopurs_runtime.Value
		{
			if v_0 {
				__t0 = gopurs_runtime.Apply(Call_Main_shout(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), gopurs_runtime.Str(__eta_norm_0_1))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(__eta_norm_0_1)).StrVal()))
		}
	end_branch_0:
		return __t0
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

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
