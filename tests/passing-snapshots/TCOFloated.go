package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_ordRecord gopurs_runtime.Value
var once_Main_ordRecord sync.Once

func Get_Main_ordRecord() gopurs_runtime.Value {
	once_Main_ordRecord.Do(func() {
		cache_Main_ordRecord = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_290402411_4177771502(Rebox_Main_4177771502_290402411(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Ord_ordRecord(gopurs_runtime.Value{}, gopurs_runtime.Apply3(Call_Data_Ord_ordRecordCons(Get_Data_Ord_ordRecordNil()), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("foo")
		})), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_3308271157_4177771502(Rebox_Main_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}))))))}
	})
	return cache_Main_ordRecord
}

var cache_Main_looper gopurs_runtime.Value
var once_Main_looper sync.Once

func Get_Main_looper() gopurs_runtime.Value {
	once_Main_looper.Do(func() {
		cache_Main_looper = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_looper(func() struct {
				foo int64
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					foo int64
				}{}
				clone.foo = gopurs_runtime.RecordGet(orig, "foo").IntVal
				return clone
			}()))
		})
	})
	return cache_Main_looper
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_looper(struct {
			foo int64
		}{int64(100000)})))
	})
	return cache_Main_main
}

func Call_Main_looper(x_0_loop struct {
	foo int64
}) string {
looper:
	for {
		if false {
			continue looper
		}
		var x_0 struct {
			foo int64
		} = x_0_loop
		_ = x_0
		var __t1 string
		{
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Ord_ordRecord(gopurs_runtime.Value{}, gopurs_runtime.Apply3(Call_Data_Ord_ordRecordCons(Get_Data_Ord_ordRecordNil()), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("foo")
			})), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_3308271157_4177771502(Rebox_Main_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))})), "compare"), func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int(orig.foo))
			}(), gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int(int64(0))))
			_ = __t_tag_0
			if (uint32(__t_tag_0.IntVal) == 380165415) != (true) {
				__t1 = "Done"
				goto end_branch_1
			} else {

			}
		}
		{
			x_0_loop = struct {
				foo int64
			}{(x_0.foo) - (int64(1))}
			continue looper
			__t1 = func() string { panic("unreachable") }()
		}
	end_branch_1:
		return __t1
	}
}

func Rebox_Main_290402411_4177771502(in *Constructor_Data_Ord_Ord[struct {
	foo int64
}]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_4177771502_290402411(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[struct {
	foo int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Ord_Ord[struct {
		foo int64
	}]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Ord_Ord[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}
