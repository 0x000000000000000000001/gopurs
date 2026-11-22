package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Database_SQLite3_bind gopurs_runtime.Value
var once_Database_SQLite3_bind sync.Once

func Get_Database_SQLite3_bind() gopurs_runtime.Value {
	once_Database_SQLite3_bind.Do(func() {
		cache_Database_SQLite3_bind = Get_Effect_Aff__bind()
	})
	return cache_Database_SQLite3_bind
}

var cache_Database_SQLite3_discard gopurs_runtime.Value
var once_Database_SQLite3_discard sync.Once

func Get_Database_SQLite3_discard() gopurs_runtime.Value {
	once_Database_SQLite3_discard.Do(func() {
		cache_Database_SQLite3_discard = Get_Effect_Aff__bind()
	})
	return cache_Database_SQLite3_discard
}

var cache_Database_SQLite3_pure gopurs_runtime.Value
var once_Database_SQLite3_pure sync.Once

func Get_Database_SQLite3_pure() gopurs_runtime.Value {
	once_Database_SQLite3_pure.Do(func() {
		cache_Database_SQLite3_pure = Get_Effect_Aff__pure()
	})
	return cache_Database_SQLite3_pure
}

var cache_Database_SQLite3_rollback gopurs_runtime.Value
var once_Database_SQLite3_rollback sync.Once

func Get_Database_SQLite3_rollback() gopurs_runtime.Value {
	once_Database_SQLite3_rollback.Do(func() {
		cache_Database_SQLite3_rollback = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Database_SQLite3_rollbackImpl())
	})
	return cache_Database_SQLite3_rollback
}

var cache_Database_SQLite3_queryTransaction gopurs_runtime.Value
var once_Database_SQLite3_queryTransaction sync.Once

func Get_Database_SQLite3_queryTransaction() gopurs_runtime.Value {
	once_Database_SQLite3_queryTransaction.Do(func() {
		cache_Database_SQLite3_queryTransaction = gopurs_runtime.Func3(func(tx_0_box gopurs_runtime.Value, sql_1_box gopurs_runtime.Value, params_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Database_SQLite3_queryTransaction(tx_0_box, sql_1_box.StrVal(), *((*[]gopurs_runtime.Value)(params_2_box.UnsafePtr)))
		})
	})
	return cache_Database_SQLite3_queryTransaction
}

var cache_Database_SQLite3_query gopurs_runtime.Value
var once_Database_SQLite3_query sync.Once

func Get_Database_SQLite3_query() gopurs_runtime.Value {
	once_Database_SQLite3_query.Do(func() {
		cache_Database_SQLite3_query = gopurs_runtime.Func3(func(db_0_box gopurs_runtime.Value, sql_1_box gopurs_runtime.Value, params_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Database_SQLite3_query(db_0_box, sql_1_box.StrVal(), *((*[]gopurs_runtime.Value)(params_2_box.UnsafePtr)))
		})
	})
	return cache_Database_SQLite3_query
}

var cache_Database_SQLite3_open gopurs_runtime.Value
var once_Database_SQLite3_open sync.Once

func Get_Database_SQLite3_open() gopurs_runtime.Value {
	once_Database_SQLite3_open.Do(func() {
		cache_Database_SQLite3_open = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Database_SQLite3_openImpl())
	})
	return cache_Database_SQLite3_open
}

var cache_Database_SQLite3_integerText gopurs_runtime.Value
var once_Database_SQLite3_integerText sync.Once

func Get_Database_SQLite3_integerText() gopurs_runtime.Value {
	once_Database_SQLite3_integerText.Do(func() {
		cache_Database_SQLite3_integerText = gopurs_runtime.Apply2(Get_Database_SQLite3_integerTextImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Database_SQLite3_742090555_3094389156(Rebox_Database_SQLite3_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))})
	})
	return cache_Database_SQLite3_integerText
}

var cache_Database_SQLite3_integer64 gopurs_runtime.Value
var once_Database_SQLite3_integer64 sync.Once

func Get_Database_SQLite3_integer64() gopurs_runtime.Value {
	once_Database_SQLite3_integer64.Do(func() {
		cache_Database_SQLite3_integer64 = gopurs_runtime.Apply2(Get_Database_SQLite3_integer64Impl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Database_SQLite3_integer64
}

var cache_Database_SQLite3_executeTransaction gopurs_runtime.Value
var once_Database_SQLite3_executeTransaction sync.Once

func Get_Database_SQLite3_executeTransaction() gopurs_runtime.Value {
	once_Database_SQLite3_executeTransaction.Do(func() {
		cache_Database_SQLite3_executeTransaction = gopurs_runtime.Func3(func(tx_0_box gopurs_runtime.Value, sql_1_box gopurs_runtime.Value, params_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Database_SQLite3_executeTransaction(tx_0_box, sql_1_box.StrVal(), *((*[]gopurs_runtime.Value)(params_2_box.UnsafePtr)))
		})
	})
	return cache_Database_SQLite3_executeTransaction
}

var cache_Database_SQLite3_execute gopurs_runtime.Value
var once_Database_SQLite3_execute sync.Once

func Get_Database_SQLite3_execute() gopurs_runtime.Value {
	once_Database_SQLite3_execute.Do(func() {
		cache_Database_SQLite3_execute = gopurs_runtime.Func3(func(db_0_box gopurs_runtime.Value, sql_1_box gopurs_runtime.Value, params_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Database_SQLite3_execute(db_0_box, sql_1_box.StrVal(), *((*[]gopurs_runtime.Value)(params_2_box.UnsafePtr)))
		})
	})
	return cache_Database_SQLite3_execute
}

var cache_Database_SQLite3_commit gopurs_runtime.Value
var once_Database_SQLite3_commit sync.Once

func Get_Database_SQLite3_commit() gopurs_runtime.Value {
	once_Database_SQLite3_commit.Do(func() {
		cache_Database_SQLite3_commit = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Database_SQLite3_commitImpl())
	})
	return cache_Database_SQLite3_commit
}

var cache_Database_SQLite3_go__close gopurs_runtime.Value
var once_Database_SQLite3_go__close sync.Once

func Get_Database_SQLite3_go__close() gopurs_runtime.Value {
	once_Database_SQLite3_go__close.Do(func() {
		cache_Database_SQLite3_go__close = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Database_SQLite3_closeImpl())
	})
	return cache_Database_SQLite3_go__close
}

var cache_Database_SQLite3_withDatabase gopurs_runtime.Value
var once_Database_SQLite3_withDatabase sync.Once

func Get_Database_SQLite3_withDatabase() gopurs_runtime.Value {
	once_Database_SQLite3_withDatabase.Do(func() {
		cache_Database_SQLite3_withDatabase = gopurs_runtime.Func(func(path_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Database_SQLite3_withDatabase(path_0_box.StrVal())
		})
	})
	return cache_Database_SQLite3_withDatabase
}

var cache_Database_SQLite3_withDatabase__3802547991 gopurs_runtime.Value
var once_Database_SQLite3_withDatabase__3802547991 sync.Once

func Get_Database_SQLite3_withDatabase__3802547991() gopurs_runtime.Value {
	once_Database_SQLite3_withDatabase__3802547991.Do(func() {
		cache_Database_SQLite3_withDatabase__3802547991 = gopurs_runtime.Func2(func(path_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Database_SQLite3_withDatabase__3802547991(path_0_box.StrVal(), __eta_norm_0_1_box)
		})
	})
	return cache_Database_SQLite3_withDatabase__3802547991
}

var cache_Database_SQLite3_begin gopurs_runtime.Value
var once_Database_SQLite3_begin sync.Once

func Get_Database_SQLite3_begin() gopurs_runtime.Value {
	once_Database_SQLite3_begin.Do(func() {
		cache_Database_SQLite3_begin = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Database_SQLite3_beginImpl())
	})
	return cache_Database_SQLite3_begin
}

var cache_Database_SQLite3_withTransaction gopurs_runtime.Value
var once_Database_SQLite3_withTransaction sync.Once

func Get_Database_SQLite3_withTransaction() gopurs_runtime.Value {
	once_Database_SQLite3_withTransaction.Do(func() {
		cache_Database_SQLite3_withTransaction = gopurs_runtime.Func2(func(db_0_box gopurs_runtime.Value, use_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Database_SQLite3_withTransaction(db_0_box, use_1_box)
		})
	})
	return cache_Database_SQLite3_withTransaction
}

func Call_Database_SQLite3_queryTransaction(tx_0_loop gopurs_runtime.Value, sql_1_loop string, params_2_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var tx_0 gopurs_runtime.Value = tx_0_loop
	_ = tx_0
	var sql_1 string = sql_1_loop
	_ = sql_1
	var params_2 []gopurs_runtime.Value = params_2_loop
	_ = params_2
	return Call_Effect_Aff_Compat_fromEffectFnAff__2546930558(gopurs_runtime.Apply3(Get_Database_SQLite3_queryTransactionImpl(), tx_0, gopurs_runtime.Str(sql_1), gopurs_runtime.Array(params_2)))
}

func Call_Database_SQLite3_query(db_0_loop gopurs_runtime.Value, sql_1_loop string, params_2_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var db_0 gopurs_runtime.Value = db_0_loop
	_ = db_0
	var sql_1 string = sql_1_loop
	_ = sql_1
	var params_2 []gopurs_runtime.Value = params_2_loop
	_ = params_2
	return Call_Effect_Aff_Compat_fromEffectFnAff__2546930558(gopurs_runtime.Apply3(Get_Database_SQLite3_queryImpl(), db_0, gopurs_runtime.Str(sql_1), gopurs_runtime.Array(params_2)))
}

func Call_Database_SQLite3_executeTransaction(tx_0_loop gopurs_runtime.Value, sql_1_loop string, params_2_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var tx_0 gopurs_runtime.Value = tx_0_loop
	_ = tx_0
	var sql_1 string = sql_1_loop
	_ = sql_1
	var params_2 []gopurs_runtime.Value = params_2_loop
	_ = params_2
	return Call_Effect_Aff_Compat_fromEffectFnAff__560502078(gopurs_runtime.Apply3(Get_Database_SQLite3_executeTransactionImpl(), tx_0, gopurs_runtime.Str(sql_1), gopurs_runtime.Array(params_2)))
}

func Call_Database_SQLite3_execute(db_0_loop gopurs_runtime.Value, sql_1_loop string, params_2_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var db_0 gopurs_runtime.Value = db_0_loop
	_ = db_0
	var sql_1 string = sql_1_loop
	_ = sql_1
	var params_2 []gopurs_runtime.Value = params_2_loop
	_ = params_2
	return Call_Effect_Aff_Compat_fromEffectFnAff__560502078(gopurs_runtime.Apply3(Get_Database_SQLite3_executeImpl(), db_0, gopurs_runtime.Str(sql_1), gopurs_runtime.Array(params_2)))
}

func Call_Database_SQLite3_withDatabase(path_0_loop string) gopurs_runtime.Value {
	var path_0 string = path_0_loop
	_ = path_0
	return gopurs_runtime.Apply2(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(Get_Database_SQLite3_open(), gopurs_runtime.Str(path_0)), func() gopurs_runtime.Value {
		orig := struct {
			completed gopurs_runtime.Value
			failed    gopurs_runtime.Value
			killed    gopurs_runtime.Value
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
		}{gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Database_SQLite3_go__close()
		}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Database_SQLite3_go__close()
		}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Database_SQLite3_go__close()
		})}
		_ = orig
		return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
	}())
}

func Call_Database_SQLite3_withDatabase__3802547991(path_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
withDatabase__3802547991:
	for {
		if false {
			continue withDatabase__3802547991
		}
		var path_0 string = path_0_loop
		_ = path_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Effect_Aff_bracket__939099628(gopurs_runtime.Apply(Get_Database_SQLite3_open(), gopurs_runtime.Str(path_0)), Get_Database_SQLite3_go__close(), __eta_norm_0_1)
	}
}

func Call_Database_SQLite3_withTransaction(db_0_loop gopurs_runtime.Value, use_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var db_0 gopurs_runtime.Value = db_0_loop
	_ = db_0
	var use_1 gopurs_runtime.Value = use_1_loop
	_ = use_1
	return gopurs_runtime.Apply3(Get_Effect_Aff_generalBracket(), gopurs_runtime.Apply(Get_Database_SQLite3_begin(), db_0), func() gopurs_runtime.Value {
		orig := struct {
			completed gopurs_runtime.Value
			failed    gopurs_runtime.Value
			killed    gopurs_runtime.Value
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,006600610069006c00650064,006b0069006c006c00650064*/
		}{gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Database_SQLite3_rollback()
		}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Database_SQLite3_rollback()
		}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Database_SQLite3_rollback()
		})}
		_ = orig
		return gopurs_runtime.RecordDict3("completed", "failed", "killed", orig.completed, orig.failed, orig.killed)
	}(), gopurs_runtime.Func(func(tx_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(use_1, tx_2), gopurs_runtime.Func(func(result_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Database_SQLite3_commit(), tx_2), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Aff__pure(), result_3)
			}))
		}))
	}))
}

func Rebox_Database_SQLite3_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Database_SQLite3_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Get_Database_SQLite3_beginImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_BeginImpl
}

func Get_Database_SQLite3_closeImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_CloseImpl
}

func Get_Database_SQLite3_commitImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_CommitImpl
}

func Get_Database_SQLite3_executeImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_ExecuteImpl
}

func Get_Database_SQLite3_executeTransactionImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_ExecuteTransactionImpl
}

func Get_Database_SQLite3_integer64Impl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_Integer64Impl
}

func Get_Database_SQLite3_integerTextImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_IntegerTextImpl
}

func Get_Database_SQLite3_null() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_Null
}

func Get_Database_SQLite3_openImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_OpenImpl
}

func Get_Database_SQLite3_queryImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_QueryImpl
}

func Get_Database_SQLite3_queryTransactionImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_QueryTransactionImpl
}

func Get_Database_SQLite3_rollbackImpl() gopurs_runtime.Value {
	return _Gopurs_Database_SQLite3_RollbackImpl
}
