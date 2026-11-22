package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Posix_go__append gopurs_runtime.Value
var once_Data_Posix_go__append sync.Once

func Get_Data_Posix_go__append() gopurs_runtime.Value {
	once_Data_Posix_go__append.Do(func() {
		cache_Data_Posix_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Posix_go__append
}

var cache_Data_Posix_show gopurs_runtime.Value
var once_Data_Posix_show sync.Once

func Get_Data_Posix_show() gopurs_runtime.Value {
	once_Data_Posix_show.Do(func() {
		cache_Data_Posix_show = Get_Data_Show_show__1986550628()
	})
	return cache_Data_Posix_show
}

var cache_Data_Posix_Uid gopurs_runtime.Value
var once_Data_Posix_Uid sync.Once

func Get_Data_Posix_Uid() gopurs_runtime.Value {
	once_Data_Posix_Uid.Do(func() {
		cache_Data_Posix_Uid = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Posix_Uid(x_0_box.IntVal))
		})
	})
	return cache_Data_Posix_Uid
}

var cache_Data_Posix_Pid gopurs_runtime.Value
var once_Data_Posix_Pid sync.Once

func Get_Data_Posix_Pid() gopurs_runtime.Value {
	once_Data_Posix_Pid.Do(func() {
		cache_Data_Posix_Pid = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Posix_Pid(x_0_box.IntVal))
		})
	})
	return cache_Data_Posix_Pid
}

var cache_Data_Posix_Gid gopurs_runtime.Value
var once_Data_Posix_Gid sync.Once

func Get_Data_Posix_Gid() gopurs_runtime.Value {
	once_Data_Posix_Gid.Do(func() {
		cache_Data_Posix_Gid = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Posix_Gid(x_0_box.IntVal))
		})
	})
	return cache_Data_Posix_Gid
}

var cache_Data_Posix_showUid gopurs_runtime.Value
var once_Data_Posix_showUid sync.Once

func Get_Data_Posix_showUid() gopurs_runtime.Value {
	once_Data_Posix_showUid.Do(func() {
		cache_Data_Posix_showUid = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Uid ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})})))}
	})
	return cache_Data_Posix_showUid
}

var cache_Data_Posix_showPid gopurs_runtime.Value
var once_Data_Posix_showPid sync.Once

func Get_Data_Posix_showPid() gopurs_runtime.Value {
	once_Data_Posix_showPid.Do(func() {
		cache_Data_Posix_showPid = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Pid ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})})))}
	})
	return cache_Data_Posix_showPid
}

var cache_Data_Posix_showGid gopurs_runtime.Value
var once_Data_Posix_showGid sync.Once

func Get_Data_Posix_showGid() gopurs_runtime.Value {
	once_Data_Posix_showGid.Do(func() {
		cache_Data_Posix_showGid = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Gid ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})})))}
	})
	return cache_Data_Posix_showGid
}

var cache_Data_Posix_ordUid gopurs_runtime.Value
var once_Data_Posix_ordUid sync.Once

func Get_Data_Posix_ordUid() gopurs_runtime.Value {
	once_Data_Posix_ordUid.Do(func() {
		cache_Data_Posix_ordUid = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_3308271157_4177771502(Rebox_Data_Posix_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Posix_ordUid
}

var cache_Data_Posix_ordPid gopurs_runtime.Value
var once_Data_Posix_ordPid sync.Once

func Get_Data_Posix_ordPid() gopurs_runtime.Value {
	once_Data_Posix_ordPid.Do(func() {
		cache_Data_Posix_ordPid = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_3308271157_4177771502(Rebox_Data_Posix_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Posix_ordPid
}

var cache_Data_Posix_ordGid gopurs_runtime.Value
var once_Data_Posix_ordGid sync.Once

func Get_Data_Posix_ordGid() gopurs_runtime.Value {
	once_Data_Posix_ordGid.Do(func() {
		cache_Data_Posix_ordGid = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_3308271157_4177771502(Rebox_Data_Posix_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Posix_ordGid
}

var cache_Data_Posix_newtypeUid gopurs_runtime.Value
var once_Data_Posix_newtypeUid sync.Once

func Get_Data_Posix_newtypeUid() gopurs_runtime.Value {
	once_Data_Posix_newtypeUid.Do(func() {
		cache_Data_Posix_newtypeUid = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_3903242120_385277032((&Constructor_Data_Newtype_Newtype[int64, int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Posix_newtypeUid
}

var cache_Data_Posix_newtypePid gopurs_runtime.Value
var once_Data_Posix_newtypePid sync.Once

func Get_Data_Posix_newtypePid() gopurs_runtime.Value {
	once_Data_Posix_newtypePid.Do(func() {
		cache_Data_Posix_newtypePid = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_3903242120_385277032((&Constructor_Data_Newtype_Newtype[int64, int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Posix_newtypePid
}

var cache_Data_Posix_newtypeGid gopurs_runtime.Value
var once_Data_Posix_newtypeGid sync.Once

func Get_Data_Posix_newtypeGid() gopurs_runtime.Value {
	once_Data_Posix_newtypeGid.Do(func() {
		cache_Data_Posix_newtypeGid = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_3903242120_385277032((&Constructor_Data_Newtype_Newtype[int64, int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Posix_newtypeGid
}

var cache_Data_Posix_eqUid gopurs_runtime.Value
var once_Data_Posix_eqUid sync.Once

func Get_Data_Posix_eqUid() gopurs_runtime.Value {
	once_Data_Posix_eqUid.Do(func() {
		cache_Data_Posix_eqUid = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_1053099733_3790796878(Rebox_Data_Posix_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Posix_eqUid
}

var cache_Data_Posix_eqPid gopurs_runtime.Value
var once_Data_Posix_eqPid sync.Once

func Get_Data_Posix_eqPid() gopurs_runtime.Value {
	once_Data_Posix_eqPid.Do(func() {
		cache_Data_Posix_eqPid = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_1053099733_3790796878(Rebox_Data_Posix_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Posix_eqPid
}

var cache_Data_Posix_eqGid gopurs_runtime.Value
var once_Data_Posix_eqGid sync.Once

func Get_Data_Posix_eqGid() gopurs_runtime.Value {
	once_Data_Posix_eqGid.Do(func() {
		cache_Data_Posix_eqGid = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_1053099733_3790796878(Rebox_Data_Posix_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Posix_eqGid
}

func Call_Data_Posix_Uid(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Posix_Pid(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Posix_Gid(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Data_Posix_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_3903242120_385277032(in *Constructor_Data_Newtype_Newtype[int64, int64]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}
