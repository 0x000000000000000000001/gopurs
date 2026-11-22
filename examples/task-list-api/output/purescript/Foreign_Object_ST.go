package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Foreign_Object_ST_peek gopurs_runtime.Value
var once_Foreign_Object_ST_peek sync.Once

func Get_Foreign_Object_ST_peek() gopurs_runtime.Value {
	once_Foreign_Object_ST_peek.Do(func() {
		cache_Foreign_Object_ST_peek = gopurs_runtime.Apply2(Get_Foreign_Object_ST_peekImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	return cache_Foreign_Object_ST_peek
}

var cache_Foreign_Object_ST_go__new gopurs_runtime.Value
var once_Foreign_Object_ST_go__new sync.Once

func Get_Foreign_Object_ST_go__new() gopurs_runtime.Value {
	once_Foreign_Object_ST_go__new.Do(func() {
		cache_Foreign_Object_ST_go__new = Get_Foreign_Object_ST_newImpl()
	})
	return cache_Foreign_Object_ST_go__new
}

var cache_Foreign_Object_ST_go__delete gopurs_runtime.Value
var once_Foreign_Object_ST_go__delete sync.Once

func Get_Foreign_Object_ST_go__delete() gopurs_runtime.Value {
	once_Foreign_Object_ST_go__delete.Do(func() {
		cache_Foreign_Object_ST_go__delete = Get_Foreign_Object_ST_deleteImpl()
	})
	return cache_Foreign_Object_ST_go__delete
}

func Get_Foreign_Object_ST_deleteImpl() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_ST_DeleteImpl
}

func Get_Foreign_Object_ST_newImpl() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_ST_NewImpl
}

func Get_Foreign_Object_ST_peekImpl() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_ST_PeekImpl
}

func Get_Foreign_Object_ST_poke() gopurs_runtime.Value {
	return _Gopurs_Foreign_Object_ST_Poke
}
