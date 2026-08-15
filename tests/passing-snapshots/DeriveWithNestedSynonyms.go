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
			return gopurs_runtime.Value{Type: 9, IntVal: 1714575428, UnsafePtr: unsafe.Pointer(&Constructor_Main_Z{1, value0})}
		})
	})
	return cache_Main_Z
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1682951303, UnsafePtr: unsafe.Pointer(&Constructor_Main_Y{1, value0})}
		})
	})
	return cache_Main_Y
}

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(&Constructor_Main_X{1, value0})}
		})
	})
	return cache_Main_X
}

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_T(x_0_box)
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
		cache_Main_eqZ = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})})}
	})
	return cache_Main_eqZ
}

var cache_Main_eqY gopurs_runtime.Value
var once_Main_eqY sync.Once

func Get_Main_eqY() gopurs_runtime.Value {
	once_Main_eqY.Do(func() {
		cache_Main_eqY = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})})}
	})
	return cache_Main_eqY
}

var cache_Main_eqX gopurs_runtime.Value
var once_Main_eqX sync.Once

func Get_Main_eqX() gopurs_runtime.Value {
	once_Main_eqX.Do(func() {
		cache_Main_eqX = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})})}
	})
	return cache_Main_eqX
}

var cache_Main_eqT gopurs_runtime.Value
var once_Main_eqT sync.Once

func Get_Main_eqT() gopurs_runtime.Value {
	once_Main_eqT.Do(func() {
		cache_Main_eqT = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "baz"), "foo").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "baz"), "foo").StrVal()))
			})
		})})}
	})
	return cache_Main_eqT
}

var cache_Main_ordT gopurs_runtime.Value
var once_Main_ordT sync.Once

func Get_Main_ordT() gopurs_runtime.Value {
	once_Main_ordT.Do(func() {
		cache_Main_ordT = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqT()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "baz"), "foo").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "baz"), "foo").StrVal())).IntVal)), UnsafePtr: nil}
			})
		})})}
	})
	return cache_Main_ordT
}

type Constructor_Main_Z struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Y struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_X struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_T(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
