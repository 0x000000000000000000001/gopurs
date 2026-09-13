package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_appendAndLog gopurs_runtime.Value
var once_Main_appendAndLog sync.Once

func Get_Main_appendAndLog() gopurs_runtime.Value {
	once_Main_appendAndLog.Do(func() {
		cache_Main_appendAndLog = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Console_log(), gopurs_runtime.Apply(Get_Data_Tuple_uncurry__1230334089(), Call_Data_Semigroup_go__append(Rebox_Main_443971153_4179793454(Rebox_Main_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))))
	})
	return cache_Main_appendAndLog
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Main_appendAndLog(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_2080100712_138441832(Rebox_Main_138441832_2080100712(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Str("Do"), gopurs_runtime.Str("ne")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))})
	})
	return cache_Main_main
}

func Rebox_Main_138441832_2080100712(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, string]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Main_2080100712_138441832(in *Constructor_Data_Tuple_Tuple[string, string]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Main_4179793454_443971153(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semigroup_Semigroup[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
