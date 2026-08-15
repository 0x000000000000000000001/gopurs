package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_NatPlus_dollar_Dict gopurs_runtime.Value
var once_Main_NatPlus_dollar_Dict sync.Once

func Get_Main_NatPlus_dollar_Dict() gopurs_runtime.Value {
	once_Main_NatPlus_dollar_Dict.Do(func() {
		cache_Main_NatPlus_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NatPlus_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_NatPlus_dollar_Dict
}

var cache_Main_NatMult_dollar_Dict gopurs_runtime.Value
var once_Main_NatMult_dollar_Dict sync.Once

func Get_Main_NatMult_dollar_Dict() gopurs_runtime.Value {
	once_Main_NatMult_dollar_Dict.Do(func() {
		cache_Main_NatMult_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NatMult_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_NatMult_dollar_Dict
}

var cache_Main_natPlusZ gopurs_runtime.Value
var once_Main_natPlusZ sync.Once

func Get_Main_natPlusZ() gopurs_runtime.Value {
	once_Main_natPlusZ.Do(func() {
		cache_Main_natPlusZ = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_natPlusZ
}

var cache_Main_natPlusS gopurs_runtime.Value
var once_Main_natPlusS sync.Once

func Get_Main_natPlusS() gopurs_runtime.Value {
	once_Main_natPlusS.Do(func() {
		cache_Main_natPlusS = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_natPlusS(_dollar___unused_0_box)
		})
	})
	return cache_Main_natPlusS
}

var cache_Main_natMultZ gopurs_runtime.Value
var once_Main_natMultZ sync.Once

func Get_Main_natMultZ() gopurs_runtime.Value {
	once_Main_natMultZ.Do(func() {
		cache_Main_natMultZ = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_natMultZ
}

var cache_Main_natMultS gopurs_runtime.Value
var once_Main_natMultS sync.Once

func Get_Main_natMultS() gopurs_runtime.Value {
	once_Main_natMultS.Do(func() {
		cache_Main_natMultS = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_natMultS(_dollar___unused_0_box, _dollar___unused_1_box)
		})
	})
	return cache_Main_natMultS
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fsingleton gopurs_runtime.Value
var once_Main_fsingleton sync.Once

func Get_Main_fsingleton() gopurs_runtime.Value {
	once_Main_fsingleton.Do(func() {
		cache_Main_fsingleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fsingleton(x_0_box)
		})
	})
	return cache_Main_fsingleton
}

var cache_Main_fflatten gopurs_runtime.Value
var once_Main_fflatten sync.Once

func Get_Main_fflatten() gopurs_runtime.Value {
	once_Main_fflatten.Do(func() {
		cache_Main_fflatten = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fflatten(uint32(_dollar___unused_0_box.IntVal))
		})
	})
	return cache_Main_fflatten
}

var cache_Main_fappend gopurs_runtime.Value
var once_Main_fappend sync.Once

func Get_Main_fappend() gopurs_runtime.Value {
	once_Main_fappend.Do(func() {
		cache_Main_fappend = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fappend(uint32(_dollar___unused_0_box.IntVal))
		})
	})
	return cache_Main_fappend
}

var cache_Main_fexample gopurs_runtime.Value
var once_Main_fexample sync.Once

func Get_Main_fexample() gopurs_runtime.Value {
	once_Main_fexample.Do(func() {
		cache_Main_fexample = gopurs_runtime.Apply2(Get_Main_fappendImpl(), gopurs_runtime.Apply2(Get_Main_fappendImpl(), gopurs_runtime.Apply2(Get_Main_fcons(), gopurs_runtime.Int(1), gopurs_runtime.Apply2(Get_Main_fcons(), gopurs_runtime.Int(2), Get_Main_fnil())), gopurs_runtime.Apply2(Get_Main_fcons(), gopurs_runtime.Int(3), Get_Main_fnil())), gopurs_runtime.Apply2(Get_Main_fcons(), gopurs_runtime.Int(4), gopurs_runtime.Apply2(Get_Main_fcons(), gopurs_runtime.Int(5), Get_Main_fnil())))
	})
	return cache_Main_fexample
}

var cache_Main_fexample2 gopurs_runtime.Value
var once_Main_fexample2 sync.Once

func Get_Main_fexample2() gopurs_runtime.Value {
	once_Main_fexample2.Do(func() {
		cache_Main_fexample2 = gopurs_runtime.Apply2(Get_Main_fappendImpl(), gopurs_runtime.Apply2(Get_Main_fappendImpl(), Get_Main_fexample(), Get_Main_fexample()), Get_Main_fexample())
	})
	return cache_Main_fexample2
}

var cache_Main_fexample3 gopurs_runtime.Value
var once_Main_fexample3 sync.Once

func Get_Main_fexample3() gopurs_runtime.Value {
	once_Main_fexample3.Do(func() {
		cache_Main_fexample3 = gopurs_runtime.Apply2(Get_Main_fappendImpl(), gopurs_runtime.Apply2(Get_Main_fappendImpl(), gopurs_runtime.Apply2(Get_Main_fcons(), Get_Main_fexample(), Get_Main_fnil()), gopurs_runtime.Apply2(Get_Main_fcons(), Get_Main_fexample(), Get_Main_fnil())), gopurs_runtime.Apply2(Get_Main_fcons(), Get_Main_fexample(), Get_Main_fnil()))
	})
	return cache_Main_fexample3
}

var cache_Main_fexample4 gopurs_runtime.Value
var once_Main_fexample4 sync.Once

func Get_Main_fexample4() gopurs_runtime.Value {
	once_Main_fexample4.Do(func() {
		cache_Main_fexample4 = gopurs_runtime.Apply(Get_Main_fflattenImpl(), Get_Main_fexample3())
	})
	return cache_Main_fexample4
}

var cache_Main_fappend__4108296857 gopurs_runtime.Value
var once_Main_fappend__4108296857 sync.Once

func Get_Main_fappend__4108296857() gopurs_runtime.Value {
	once_Main_fappend__4108296857.Do(func() {
		cache_Main_fappend__4108296857 = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fappend__4108296857(uint32(_dollar___unused_0_box.IntVal))
		})
	})
	return cache_Main_fappend__4108296857
}

var cache_Main_fflatten__2162348848 gopurs_runtime.Value
var once_Main_fflatten__2162348848 sync.Once

func Get_Main_fflatten__2162348848() gopurs_runtime.Value {
	once_Main_fflatten__2162348848.Do(func() {
		cache_Main_fflatten__2162348848 = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fflatten__2162348848(uint32(_dollar___unused_0_box.IntVal))
		})
	})
	return cache_Main_fflatten__2162348848
}

type Constructor_Main_NatPlus struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3417617535] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_NatPlus)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_NatPlus: " + key)
		}
	}
}

type Constructor_Main_NatMult struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2508830309] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_NatMult)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_NatMult: " + key)
		}
	}
}

func Call_Main_NatPlus_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_NatMult_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_natPlusS(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
}

func Call_Main_natMultS(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
}

func Call_Main_fsingleton(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return gopurs_runtime.Apply2(Get_Main_fcons(), x_0, Get_Main_fnil())
}

func Call_Main_fflatten(_dollar___unused_0_loop uint32) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Main_fflattenImpl()
}

func Call_Main_fappend(_dollar___unused_0_loop uint32) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Main_fappendImpl()
}

func Call_Main_fappend__4108296857(_dollar___unused_0_loop uint32) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Main_fappendImpl()
}

func Call_Main_fflatten__2162348848(_dollar___unused_0_loop uint32) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Main_fflattenImpl()
}

func Get_Main_fappendImpl() gopurs_runtime.Value {
	return _Gopurs_Main_FappendImpl
}

func Get_Main_fcons() gopurs_runtime.Value {
	return _Gopurs_Main_Fcons
}

func Get_Main_fflattenImpl() gopurs_runtime.Value {
	return _Gopurs_Main_FflattenImpl
}

func Get_Main_fnil() gopurs_runtime.Value {
	return _Gopurs_Main_Fnil
}

func Get_Main_ftoArray() gopurs_runtime.Value {
	return _Gopurs_Main_FtoArray
}
