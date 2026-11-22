package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Nullable_compose gopurs_runtime.Value
var once_Data_Nullable_compose sync.Once

func Get_Data_Nullable_compose() gopurs_runtime.Value {
	once_Data_Nullable_compose.Do(func() {
		cache_Data_Nullable_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Nullable_compose
}

var cache_Data_Nullable_toNullable gopurs_runtime.Value
var once_Data_Nullable_toNullable sync.Once

func Get_Data_Nullable_toNullable() gopurs_runtime.Value {
	once_Data_Nullable_toNullable.Do(func() {
		cache_Data_Nullable_toNullable = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Nullable_toNullable(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_0_box))
		})
	})
	return cache_Data_Nullable_toNullable
}

var cache_Data_Nullable_toMaybe gopurs_runtime.Value
var once_Data_Nullable_toMaybe sync.Once

func Get_Data_Nullable_toMaybe() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe.Do(func() {
		cache_Data_Nullable_toMaybe = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe(n_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe
}

var cache_Data_Nullable_toMaybe__1335646348 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__1335646348 sync.Once

func Get_Data_Nullable_toMaybe__1335646348() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__1335646348.Do(func() {
		cache_Data_Nullable_toMaybe__1335646348 = gopurs_runtime.Func(func(n_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__1335646348(n_unused_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__1335646348
}

var cache_Data_Nullable_toMaybe__4230544212 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__4230544212 sync.Once

func Get_Data_Nullable_toMaybe__4230544212() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__4230544212.Do(func() {
		cache_Data_Nullable_toMaybe__4230544212 = gopurs_runtime.Func(func(n_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__4230544212(n_unused_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__4230544212
}

var cache_Data_Nullable_toMaybe__2363327311 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__2363327311 sync.Once

func Get_Data_Nullable_toMaybe__2363327311() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__2363327311.Do(func() {
		cache_Data_Nullable_toMaybe__2363327311 = gopurs_runtime.Func(func(n_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__2363327311(n_unused_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__2363327311
}

var cache_Data_Nullable_toMaybe__3477274022 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__3477274022 sync.Once

func Get_Data_Nullable_toMaybe__3477274022() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__3477274022.Do(func() {
		cache_Data_Nullable_toMaybe__3477274022 = gopurs_runtime.Func(func(n_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__3477274022(n_unused_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__3477274022
}

var cache_Data_Nullable_toMaybe__1352391578 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__1352391578 sync.Once

func Get_Data_Nullable_toMaybe__1352391578() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__1352391578.Do(func() {
		cache_Data_Nullable_toMaybe__1352391578 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__1352391578(n_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__1352391578
}

var cache_Data_Nullable_toMaybe__3221728186 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__3221728186 sync.Once

func Get_Data_Nullable_toMaybe__3221728186() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__3221728186.Do(func() {
		cache_Data_Nullable_toMaybe__3221728186 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__3221728186(n_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__3221728186
}

var cache_Data_Nullable_toMaybe__1844236282 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__1844236282 sync.Once

func Get_Data_Nullable_toMaybe__1844236282() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__1844236282.Do(func() {
		cache_Data_Nullable_toMaybe__1844236282 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__1844236282(n_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__1844236282
}

var cache_Data_Nullable_toMaybe__1451648826 gopurs_runtime.Value
var once_Data_Nullable_toMaybe__1451648826 sync.Once

func Get_Data_Nullable_toMaybe__1451648826() gopurs_runtime.Value {
	once_Data_Nullable_toMaybe__1451648826.Do(func() {
		cache_Data_Nullable_toMaybe__1451648826 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Nullable_toMaybe__1451648826(n_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Nullable_toMaybe__1451648826
}

var cache_Data_Nullable_showNullable gopurs_runtime.Value
var once_Data_Nullable_showNullable sync.Once

func Get_Data_Nullable_showNullable() gopurs_runtime.Value {
	once_Data_Nullable_showNullable.Do(func() {
		cache_Data_Nullable_showNullable = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Nullable_showNullable(dictShow_0_box)
		})
	})
	return cache_Data_Nullable_showNullable
}

var cache_Data_Nullable_eqNullable gopurs_runtime.Value
var once_Data_Nullable_eqNullable sync.Once

func Get_Data_Nullable_eqNullable() gopurs_runtime.Value {
	once_Data_Nullable_eqNullable.Do(func() {
		cache_Data_Nullable_eqNullable = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Nullable_eqNullable(dictEq_0_box)
		})
	})
	return cache_Data_Nullable_eqNullable
}

var cache_Data_Nullable_ordNullable gopurs_runtime.Value
var once_Data_Nullable_ordNullable sync.Once

func Get_Data_Nullable_ordNullable() gopurs_runtime.Value {
	once_Data_Nullable_ordNullable.Do(func() {
		cache_Data_Nullable_ordNullable = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Nullable_ordNullable(dictOrd_0_box)
		})
	})
	return cache_Data_Nullable_ordNullable
}

var cache_Data_Nullable_eq1Nullable gopurs_runtime.Value
var once_Data_Nullable_eq1Nullable sync.Once

func Get_Data_Nullable_eq1Nullable() gopurs_runtime.Value {
	once_Data_Nullable_eq1Nullable.Do(func() {
		cache_Data_Nullable_eq1Nullable = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Nullable_eqNullable(dictEq_0)))
		})}))}
	})
	return cache_Data_Nullable_eq1Nullable
}

var cache_Data_Nullable_ord1Nullable gopurs_runtime.Value
var once_Data_Nullable_ord1Nullable sync.Once

func Get_Data_Nullable_ord1Nullable() gopurs_runtime.Value {
	once_Data_Nullable_ord1Nullable.Do(func() {
		cache_Data_Nullable_ord1Nullable = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Nullable_eq1Nullable()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Nullable_ordNullable(dictOrd_0)))
		})}))}
	})
	return cache_Data_Nullable_ord1Nullable
}

func Call_Data_Nullable_toNullable(v2_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
	var v2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_0_loop
	_ = v2_0
	var __t0 gopurs_runtime.Value
	{
		if v2_0 == nil {
			__t0 = Get_Data_Nullable_null()
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_0 != nil {
			__t0 = gopurs_runtime.Apply(Get_Data_Nullable_notNull(), (v2_0).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Nullable_toMaybe(n_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var n_0 gopurs_runtime.Value = n_0_loop
	_ = n_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), n_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Nullable_toMaybe__1335646348(n_unused_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__1335646348:
	for {
		if false {
			continue toMaybe__1335646348
		}
		var n_unused_0 gopurs_runtime.Value = n_unused_0_loop
		_ = n_unused_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_abortImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_toMaybe__4230544212(n_unused_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__4230544212:
	for {
		if false {
			continue toMaybe__4230544212
		}
		var n_unused_0 gopurs_runtime.Value = n_unused_0_loop
		_ = n_unused_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_channelRefImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_toMaybe__2363327311(n_unused_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__2363327311:
	for {
		if false {
			continue toMaybe__2363327311
		}
		var n_unused_0 gopurs_runtime.Value = n_unused_0_loop
		_ = n_unused_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_channelUnrefImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_toMaybe__3477274022(n_unused_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__3477274022:
	for {
		if false {
			continue toMaybe__3477274022
		}
		var n_unused_0 gopurs_runtime.Value = n_unused_0_loop
		_ = n_unused_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_disconnectImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_toMaybe__1352391578(n_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__1352391578:
	for {
		if false {
			continue toMaybe__1352391578
		}
		var n_0 gopurs_runtime.Value = n_0_loop
		_ = n_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), n_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_toMaybe__3221728186(n_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__3221728186:
	for {
		if false {
			continue toMaybe__3221728186
		}
		var n_0 gopurs_runtime.Value = n_0_loop
		_ = n_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), n_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_toMaybe__1844236282(n_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__1844236282:
	for {
		if false {
			continue toMaybe__1844236282
		}
		var n_0 gopurs_runtime.Value = n_0_loop
		_ = n_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), n_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_toMaybe__1451648826(n_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toMaybe__1451648826:
	for {
		if false {
			continue toMaybe__1451648826
		}
		var n_0 gopurs_runtime.Value = n_0_loop
		_ = n_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Nullable_1170268447_3094389156(Rebox_Data_Nullable_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), n_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Nullable_1170268447_3094389156(Rebox_Data_Nullable_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}, Get_Data_Maybe_Just())))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Nullable_showNullable(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope15)] String)
	__local_var_1_0 := Call_Data_Show_show(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil {
				__t1 = gopurs_runtime.StrWithCodeUnits("null", 4)
				goto end_branch_1
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil {
				__t1 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}), Get_Data_Nullable_toMaybe())}))}
}

func Call_Data_Nullable_eqNullable(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Maybe_eqMaybe(dictEq_0)))
	_ = __local_var_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_1_0, func() gopurs_runtime.Value {
			_v := Call_Data_Nullable_toMaybe(x_2)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}(), func() gopurs_runtime.Value {
			_v := Call_Data_Nullable_toMaybe(y_3)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	})}))}
}

func Call_Data_Nullable_ordNullable(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqNullable1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Nullable","Nullable"] [(TypeVar a$scope19)])])
	eqNullable1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Nullable_eqNullable(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
	_ = eqNullable1_1_0
	// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=Any
	__local_var_2_1 := Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Maybe_ordMaybe(dictOrd_0)))
	_ = __local_var_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNullable1_1_0)}
	}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_2_1, func() gopurs_runtime.Value {
			_v := Call_Data_Nullable_toMaybe(x_3)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}(), func() gopurs_runtime.Value {
			_v := Call_Data_Nullable_toMaybe(y_4)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	})}))}
}

func Rebox_Data_Nullable_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Nullable_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Get_Data_Nullable_notNull() gopurs_runtime.Value {
	return _Gopurs_Data_Nullable_NotNull
}

func Get_Data_Nullable_null() gopurs_runtime.Value {
	return _Gopurs_Data_Nullable_Null
}

func Get_Data_Nullable_nullable() gopurs_runtime.Value {
	return _Gopurs_Data_Nullable_Nullable
}
