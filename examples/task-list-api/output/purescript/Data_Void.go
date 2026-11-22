package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Void_Void gopurs_runtime.Value
var once_Data_Void_Void sync.Once

func Get_Data_Void_Void() gopurs_runtime.Value {
	once_Data_Void_Void.Do(func() {
		cache_Data_Void_Void = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Void_Void(x_0_box)
		})
	})
	return cache_Data_Void_Void
}

var cache_Data_Void_absurd gopurs_runtime.Value
var once_Data_Void_absurd sync.Once

func Get_Data_Void_absurd() gopurs_runtime.Value {
	once_Data_Void_absurd.Do(func() {
		cache_Data_Void_absurd = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Void_absurd(a_0_box)
		})
	})
	return cache_Data_Void_absurd
}

var cache_Data_Void_absurd__1978526906 gopurs_runtime.Value
var once_Data_Void_absurd__1978526906 sync.Once

func Get_Data_Void_absurd__1978526906() gopurs_runtime.Value {
	once_Data_Void_absurd__1978526906.Do(func() {
		cache_Data_Void_absurd__1978526906 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Void_absurd__1978526906(a_0_box)), UnsafePtr: nil}
		})
	})
	return cache_Data_Void_absurd__1978526906
}

var cache_Data_Void_absurd__2332778112 gopurs_runtime.Value
var once_Data_Void_absurd__2332778112 sync.Once

func Get_Data_Void_absurd__2332778112() gopurs_runtime.Value {
	once_Data_Void_absurd__2332778112.Do(func() {
		cache_Data_Void_absurd__2332778112 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Void_absurd__2332778112(a_0_box))
		})
	})
	return cache_Data_Void_absurd__2332778112
}

func Call_Data_Void_Void(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Void_absurd(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var Call_local_Data_Void_spin__1769020947_1_0_0 func(gopurs_runtime.Value) gopurs_runtime.Value
	_ = Call_local_Data_Void_spin__1769020947_1_0_0
	var spin__1769020947_1_0_0 gopurs_runtime.Value
	_ = spin__1769020947_1_0_0
	var Call_local_Data_Void_spin_1_1_1 func(gopurs_runtime.Value) gopurs_runtime.Value
	_ = Call_local_Data_Void_spin_1_1_1
	var spin_1_1_1 gopurs_runtime.Value
	_ = spin_1_1_1
	Call_local_Data_Void_spin__1769020947_1_0_0 = func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	spin__1769020947_1_0_0:
		for {
			if false {
				continue spin__1769020947_1_0_0
			}
			var v_2 gopurs_runtime.Value = v_2_loop
			_ = v_2
			v_2_loop = v_2
			continue spin__1769020947_1_0_0
			return func() gopurs_runtime.Value { panic("unreachable") }()
		}
	}
	spin__1769020947_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Data_Void_spin__1769020947_1_0_0(v_2_loop_val)
	})
	Call_local_Data_Void_spin_1_1_1 = func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	spin_1_1_1:
		for {
			if false {
				continue spin_1_1_1
			}
			var v_2 gopurs_runtime.Value = v_2_loop
			_ = v_2
			return Call_local_Data_Void_spin__1769020947_1_0_0(v_2)
		}
	}
	spin_1_1_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Data_Void_spin_1_1_1(v_2_loop_val)
	})
	return Call_local_Data_Void_spin__1769020947_1_0_0(a_0)
}

func Call_Data_Void_absurd__1978526906(a_0_loop gopurs_runtime.Value) uint32 {
absurd__1978526906:
	for {
		if false {
			continue absurd__1978526906
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		var Call_local_Data_Void_spin__1133204955_1_0_2 func(gopurs_runtime.Value) uint32
		_ = Call_local_Data_Void_spin__1133204955_1_0_2
		var spin__1133204955_1_0_2 gopurs_runtime.Value
		_ = spin__1133204955_1_0_2
		var Call_local_Data_Void_spin_1_1_3 func(gopurs_runtime.Value) uint32
		_ = Call_local_Data_Void_spin_1_1_3
		var spin_1_1_3 gopurs_runtime.Value
		_ = spin_1_1_3
		Call_local_Data_Void_spin__1133204955_1_0_2 = func(v_2_loop gopurs_runtime.Value) uint32 {
		spin__1133204955_1_0_2:
			for {
				if false {
					continue spin__1133204955_1_0_2
				}
				var v_2 gopurs_runtime.Value = v_2_loop
				_ = v_2
				v_2_loop = v_2
				continue spin__1133204955_1_0_2
				return func() uint32 { panic("unreachable") }()
			}
		}
		spin__1133204955_1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Void_spin__1133204955_1_0_2(v_2_loop_val)), UnsafePtr: nil}
		})
		Call_local_Data_Void_spin_1_1_3 = func(v_2_loop gopurs_runtime.Value) uint32 {
		spin_1_1_3:
			for {
				if false {
					continue spin_1_1_3
				}
				var v_2 gopurs_runtime.Value = v_2_loop
				_ = v_2
				return Call_local_Data_Void_spin__1133204955_1_0_2(v_2)
			}
		}
		spin_1_1_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_Void_spin_1_1_3(v_2_loop_val)), UnsafePtr: nil}
		})
		return Call_local_Data_Void_spin__1133204955_1_0_2(a_0)
	}
}

func Call_Data_Void_absurd__2332778112(a_0_loop gopurs_runtime.Value) bool {
absurd__2332778112:
	for {
		if false {
			continue absurd__2332778112
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		var Call_local_Data_Void_spin__2182866177_1_0_4 func(gopurs_runtime.Value) bool
		_ = Call_local_Data_Void_spin__2182866177_1_0_4
		var spin__2182866177_1_0_4 gopurs_runtime.Value
		_ = spin__2182866177_1_0_4
		var Call_local_Data_Void_spin_1_1_5 func(gopurs_runtime.Value) bool
		_ = Call_local_Data_Void_spin_1_1_5
		var spin_1_1_5 gopurs_runtime.Value
		_ = spin_1_1_5
		Call_local_Data_Void_spin__2182866177_1_0_4 = func(v_2_loop gopurs_runtime.Value) bool {
		spin__2182866177_1_0_4:
			for {
				if false {
					continue spin__2182866177_1_0_4
				}
				var v_2 gopurs_runtime.Value = v_2_loop
				_ = v_2
				v_2_loop = v_2
				continue spin__2182866177_1_0_4
				return func() bool { panic("unreachable") }()
			}
		}
		spin__2182866177_1_0_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_local_Data_Void_spin__2182866177_1_0_4(v_2_loop_val))
		})
		Call_local_Data_Void_spin_1_1_5 = func(v_2_loop gopurs_runtime.Value) bool {
		spin_1_1_5:
			for {
				if false {
					continue spin_1_1_5
				}
				var v_2 gopurs_runtime.Value = v_2_loop
				_ = v_2
				return Call_local_Data_Void_spin__2182866177_1_0_4(v_2)
			}
		}
		spin_1_1_5 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_local_Data_Void_spin_1_1_5(v_2_loop_val))
		})
		return Call_local_Data_Void_spin__2182866177_1_0_4(a_0)
	}
}
