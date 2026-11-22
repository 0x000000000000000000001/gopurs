package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Ref_go__new gopurs_runtime.Value
var once_Effect_Ref_go__new sync.Once

func Get_Effect_Ref_go__new() gopurs_runtime.Value {
	once_Effect_Ref_go__new.Do(func() {
		cache_Effect_Ref_go__new = Get_Effect_Ref__new()
	})
	return cache_Effect_Ref_go__new
}

var cache_Effect_Ref_new__3819798246 gopurs_runtime.Value
var once_Effect_Ref_new__3819798246 sync.Once

func Get_Effect_Ref_new__3819798246() gopurs_runtime.Value {
	once_Effect_Ref_new__3819798246.Do(func() {
		cache_Effect_Ref_new__3819798246 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Ref_new__3819798246(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Effect_Ref_new__3819798246
}

var cache_Effect_Ref_new__4125901334 gopurs_runtime.Value
var once_Effect_Ref_new__4125901334 sync.Once

func Get_Effect_Ref_new__4125901334() gopurs_runtime.Value {
	once_Effect_Ref_new__4125901334.Do(func() {
		cache_Effect_Ref_new__4125901334 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Ref_new__4125901334(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Effect_Ref_new__4125901334
}

var cache_Effect_Ref_new__3830549650 gopurs_runtime.Value
var once_Effect_Ref_new__3830549650 sync.Once

func Get_Effect_Ref_new__3830549650() gopurs_runtime.Value {
	once_Effect_Ref_new__3830549650.Do(func() {
		cache_Effect_Ref_new__3830549650 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Ref_new__3830549650((__eta_norm_0_0_box.IntVal) != (0))
		})
	})
	return cache_Effect_Ref_new__3830549650
}

var cache_Effect_Ref_new__1316289842 gopurs_runtime.Value
var once_Effect_Ref_new__1316289842 sync.Once

func Get_Effect_Ref_new__1316289842() gopurs_runtime.Value {
	once_Effect_Ref_new__1316289842.Do(func() {
		cache_Effect_Ref_new__1316289842 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Ref_new__1316289842(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Effect_Ref_new__1316289842
}

var cache_Effect_Ref_modify_prime_ gopurs_runtime.Value
var once_Effect_Ref_modify_prime_ sync.Once

func Get_Effect_Ref_modify_prime_() gopurs_runtime.Value {
	once_Effect_Ref_modify_prime_.Do(func() {
		cache_Effect_Ref_modify_prime_ = Get_Effect_Ref_modifyImpl()
	})
	return cache_Effect_Ref_modify_prime_
}

var cache_Effect_Ref_modify_prime___3131127740 gopurs_runtime.Value
var once_Effect_Ref_modify_prime___3131127740 sync.Once

func Get_Effect_Ref_modify_prime___3131127740() gopurs_runtime.Value {
	once_Effect_Ref_modify_prime___3131127740.Do(func() {
		cache_Effect_Ref_modify_prime___3131127740 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Ref_modify_prime___3131127740(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Ref_modify_prime___3131127740
}

var cache_Effect_Ref_modify gopurs_runtime.Value
var once_Effect_Ref_modify sync.Once

func Get_Effect_Ref_modify() gopurs_runtime.Value {
	once_Effect_Ref_modify.Do(func() {
		cache_Effect_Ref_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Ref_modify(f_0_box)
		})
	})
	return cache_Effect_Ref_modify
}

var cache_Effect_Ref_modify__1373575713 gopurs_runtime.Value
var once_Effect_Ref_modify__1373575713 sync.Once

func Get_Effect_Ref_modify__1373575713() gopurs_runtime.Value {
	once_Effect_Ref_modify__1373575713.Do(func() {
		cache_Effect_Ref_modify__1373575713 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Ref_modify__1373575713(f_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Effect_Ref_modify__1373575713
}

func Call_Effect_Ref_new__3819798246(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
new__3819798246:
	for {
		if false {
			continue new__3819798246
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Node_Stream_Aff_pure(), Get_Data_Unit_unit()))
	}
}

func Call_Effect_Ref_new__4125901334(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
new__4125901334:
	for {
		if false {
			continue new__4125901334
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Node_Stream_Aff_mempty())
	}
}

func Call_Effect_Ref_new__3830549650(__eta_norm_0_0_loop bool) gopurs_runtime.Value {
new__3830549650:
	for {
		if false {
			continue new__3830549650
		}
		var __eta_norm_0_0 bool = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Bool(__eta_norm_0_0))
	}
}

func Call_Effect_Ref_new__1316289842(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
new__1316289842:
	for {
		if false {
			continue new__1316289842
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(__eta_norm_0_0))
	}
}

func Call_Effect_Ref_modify_prime___3131127740(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
modify_prime___3131127740:
	for {
		if false {
			continue modify_prime___3131127740
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(Get_Effect_Ref_modifyImpl(), __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Effect_Ref_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(Get_Effect_Ref_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): s_prime__2_0 shape=App(Other) bindingType=(TypeVar s$scope10)
		s_prime__2_0 := gopurs_runtime.Apply(f_0, s_1)
		_ = s_prime__2_0
		return gopurs_runtime.RecordDict2("state", "value", s_prime__2_0, s_prime__2_0)
	}))
}

func Call_Effect_Ref_modify__1373575713(f_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
modify__1373575713:
	for {
		if false {
			continue modify__1373575713
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(Get_Effect_Ref_modifyImpl(), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): s_prime__3_0 shape=App(Other) bindingType=Int
			s_prime__3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Int(s_2.IntVal)).IntVal
			_ = s_prime__3_0
			return gopurs_runtime.RecordDict2("state", "value", gopurs_runtime.Int(s_prime__3_0), gopurs_runtime.Int(s_prime__3_0))
		}), __eta_norm_0_1)
	}
}

func Get_Effect_Ref__new() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref__New
}

func Get_Effect_Ref_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_ModifyImpl
}

func Get_Effect_Ref_modify_() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_Modify_
}

func Get_Effect_Ref_newWithSelf() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_NewWithSelf
}

func Get_Effect_Ref_read() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_Read
}

func Get_Effect_Ref_write() gopurs_runtime.Value {
	return _Gopurs_Effect_Ref_Write
}
