package Type_RowList

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var rowListSetImpl gopurs_runtime.Value
var once_rowListSetImpl sync.Once
func Get_rowListSetImpl() gopurs_runtime.Value {
	once_rowListSetImpl.Do(func() {
		rowListSetImpl = gopurs_runtime.Func(func(dictTypeEquals_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTypeEquals1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
})
	})
	return rowListSetImpl
}

var rowListRemoveNil gopurs_runtime.Value
var once_rowListRemoveNil sync.Once
func Get_rowListRemoveNil() gopurs_runtime.Value {
	once_rowListRemoveNil.Do(func() {
		rowListRemoveNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return rowListRemoveNil
}

var rowListRemoveCons gopurs_runtime.Value
var once_rowListRemoveCons sync.Once
func Get_rowListRemoveCons() gopurs_runtime.Value {
	once_rowListRemoveCons.Do(func() {
		rowListRemoveCons = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
})
	})
	return rowListRemoveCons
}

var rowListNubNil gopurs_runtime.Value
var once_rowListNubNil sync.Once
func Get_rowListNubNil() gopurs_runtime.Value {
	once_rowListNubNil.Do(func() {
		rowListNubNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return rowListNubNil
}

var rowListNubCons gopurs_runtime.Value
var once_rowListNubCons sync.Once
func Get_rowListNubCons() gopurs_runtime.Value {
	once_rowListNubCons.Do(func() {
		rowListNubCons = gopurs_runtime.Func(func(dictTypeEquals_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTypeEquals1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTypeEquals2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
})
})
})
	})
	return rowListNubCons
}

var rowListAppendNil gopurs_runtime.Value
var once_rowListAppendNil sync.Once
func Get_rowListAppendNil() gopurs_runtime.Value {
	once_rowListAppendNil.Do(func() {
		rowListAppendNil = gopurs_runtime.Func(func(dictTypeEquals_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
	})
	return rowListAppendNil
}

var rowListAppendCons gopurs_runtime.Value
var once_rowListAppendCons sync.Once
func Get_rowListAppendCons() gopurs_runtime.Value {
	once_rowListAppendCons.Do(func() {
		rowListAppendCons = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTypeEquals_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
	})
	return rowListAppendCons
}

var listToRowNil gopurs_runtime.Value
var once_listToRowNil sync.Once
func Get_listToRowNil() gopurs_runtime.Value {
	once_listToRowNil.Do(func() {
		listToRowNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return listToRowNil
}

var listToRowCons gopurs_runtime.Value
var once_listToRowCons sync.Once
func Get_listToRowCons() gopurs_runtime.Value {
	once_listToRowCons.Do(func() {
		listToRowCons = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
	})
	return listToRowCons
}


