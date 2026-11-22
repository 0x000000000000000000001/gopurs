package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Buffer_Class_MutableBuffer_dollar_Dict gopurs_runtime.Value
var once_Node_Buffer_Class_MutableBuffer_dollar_Dict sync.Once

func Get_Node_Buffer_Class_MutableBuffer_dollar_Dict() gopurs_runtime.Value {
	once_Node_Buffer_Class_MutableBuffer_dollar_Dict.Do(func() {
		cache_Node_Buffer_Class_MutableBuffer_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1692970589, UnsafePtr: unsafe.Pointer(Call_Node_Buffer_Class_MutableBuffer_dollar_Dict(func() struct {
				Monad0                               gopurs_runtime.Value
				concat                               gopurs_runtime.Value
				gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
				copy                                 gopurs_runtime.Value
				create                               gopurs_runtime.Value
				fill                                 gopurs_runtime.Value
				freeze                               gopurs_runtime.Value
				fromArray                            gopurs_runtime.Value
				fromArrayBuffer                      gopurs_runtime.Value
				fromString                           gopurs_runtime.Value
				getAtOffset                          gopurs_runtime.Value
				read                                 gopurs_runtime.Value
				readString                           gopurs_runtime.Value
				setAtOffset                          gopurs_runtime.Value
				size                                 gopurs_runtime.Value
				slice                                gopurs_runtime.Value
				thaw                                 gopurs_runtime.Value
				toArray                              gopurs_runtime.Value
				toArrayBuffer                        gopurs_runtime.Value
				toString                             gopurs_runtime.Value
				unsafeFreeze                         gopurs_runtime.Value
				unsafeThaw                           gopurs_runtime.Value
				write                                gopurs_runtime.Value
				writeString                          gopurs_runtime.Value
				/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0                               gopurs_runtime.Value
					concat                               gopurs_runtime.Value
					gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
					copy                                 gopurs_runtime.Value
					create                               gopurs_runtime.Value
					fill                                 gopurs_runtime.Value
					freeze                               gopurs_runtime.Value
					fromArray                            gopurs_runtime.Value
					fromArrayBuffer                      gopurs_runtime.Value
					fromString                           gopurs_runtime.Value
					getAtOffset                          gopurs_runtime.Value
					read                                 gopurs_runtime.Value
					readString                           gopurs_runtime.Value
					setAtOffset                          gopurs_runtime.Value
					size                                 gopurs_runtime.Value
					slice                                gopurs_runtime.Value
					thaw                                 gopurs_runtime.Value
					toArray                              gopurs_runtime.Value
					toArrayBuffer                        gopurs_runtime.Value
					toString                             gopurs_runtime.Value
					unsafeFreeze                         gopurs_runtime.Value
					unsafeThaw                           gopurs_runtime.Value
					write                                gopurs_runtime.Value
					writeString                          gopurs_runtime.Value
					/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.concat = gopurs_runtime.RecordGet(orig, "concat")
				clone.gopurs_field_99_111_110_99_97_116_39 = gopurs_runtime.RecordGet(orig, "concat'")
				clone.copy = gopurs_runtime.RecordGet(orig, "copy")
				clone.create = gopurs_runtime.RecordGet(orig, "create")
				clone.fill = gopurs_runtime.RecordGet(orig, "fill")
				clone.freeze = gopurs_runtime.RecordGet(orig, "freeze")
				clone.fromArray = gopurs_runtime.RecordGet(orig, "fromArray")
				clone.fromArrayBuffer = gopurs_runtime.RecordGet(orig, "fromArrayBuffer")
				clone.fromString = gopurs_runtime.RecordGet(orig, "fromString")
				clone.getAtOffset = gopurs_runtime.RecordGet(orig, "getAtOffset")
				clone.read = gopurs_runtime.RecordGet(orig, "read")
				clone.readString = gopurs_runtime.RecordGet(orig, "readString")
				clone.setAtOffset = gopurs_runtime.RecordGet(orig, "setAtOffset")
				clone.size = gopurs_runtime.RecordGet(orig, "size")
				clone.slice = gopurs_runtime.RecordGet(orig, "slice")
				clone.thaw = gopurs_runtime.RecordGet(orig, "thaw")
				clone.toArray = gopurs_runtime.RecordGet(orig, "toArray")
				clone.toArrayBuffer = gopurs_runtime.RecordGet(orig, "toArrayBuffer")
				clone.toString = gopurs_runtime.RecordGet(orig, "toString")
				clone.unsafeFreeze = gopurs_runtime.RecordGet(orig, "unsafeFreeze")
				clone.unsafeThaw = gopurs_runtime.RecordGet(orig, "unsafeThaw")
				clone.write = gopurs_runtime.RecordGet(orig, "write")
				clone.writeString = gopurs_runtime.RecordGet(orig, "writeString")
				return clone
			}()))}
		})
	})
	return cache_Node_Buffer_Class_MutableBuffer_dollar_Dict
}

var cache_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957 gopurs_runtime.Value
var once_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957 sync.Once

func Get_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957() gopurs_runtime.Value {
	once_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957.Do(func() {
		cache_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1692970589, UnsafePtr: unsafe.Pointer(Call_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957(func() struct {
				Monad0                               gopurs_runtime.Value
				concat                               gopurs_runtime.Value
				gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
				copy                                 gopurs_runtime.Value
				create                               gopurs_runtime.Value
				fill                                 gopurs_runtime.Value
				freeze                               gopurs_runtime.Value
				fromArray                            gopurs_runtime.Value
				fromArrayBuffer                      gopurs_runtime.Value
				fromString                           gopurs_runtime.Value
				getAtOffset                          gopurs_runtime.Value
				read                                 gopurs_runtime.Value
				readString                           gopurs_runtime.Value
				setAtOffset                          gopurs_runtime.Value
				size                                 gopurs_runtime.Value
				slice                                gopurs_runtime.Value
				thaw                                 gopurs_runtime.Value
				toArray                              gopurs_runtime.Value
				toArrayBuffer                        gopurs_runtime.Value
				toString                             gopurs_runtime.Value
				unsafeFreeze                         gopurs_runtime.Value
				unsafeThaw                           gopurs_runtime.Value
				write                                gopurs_runtime.Value
				writeString                          gopurs_runtime.Value
				/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0                               gopurs_runtime.Value
					concat                               gopurs_runtime.Value
					gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
					copy                                 gopurs_runtime.Value
					create                               gopurs_runtime.Value
					fill                                 gopurs_runtime.Value
					freeze                               gopurs_runtime.Value
					fromArray                            gopurs_runtime.Value
					fromArrayBuffer                      gopurs_runtime.Value
					fromString                           gopurs_runtime.Value
					getAtOffset                          gopurs_runtime.Value
					read                                 gopurs_runtime.Value
					readString                           gopurs_runtime.Value
					setAtOffset                          gopurs_runtime.Value
					size                                 gopurs_runtime.Value
					slice                                gopurs_runtime.Value
					thaw                                 gopurs_runtime.Value
					toArray                              gopurs_runtime.Value
					toArrayBuffer                        gopurs_runtime.Value
					toString                             gopurs_runtime.Value
					unsafeFreeze                         gopurs_runtime.Value
					unsafeThaw                           gopurs_runtime.Value
					write                                gopurs_runtime.Value
					writeString                          gopurs_runtime.Value
					/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.concat = gopurs_runtime.RecordGet(orig, "concat")
				clone.gopurs_field_99_111_110_99_97_116_39 = gopurs_runtime.RecordGet(orig, "concat'")
				clone.copy = gopurs_runtime.RecordGet(orig, "copy")
				clone.create = gopurs_runtime.RecordGet(orig, "create")
				clone.fill = gopurs_runtime.RecordGet(orig, "fill")
				clone.freeze = gopurs_runtime.RecordGet(orig, "freeze")
				clone.fromArray = gopurs_runtime.RecordGet(orig, "fromArray")
				clone.fromArrayBuffer = gopurs_runtime.RecordGet(orig, "fromArrayBuffer")
				clone.fromString = gopurs_runtime.RecordGet(orig, "fromString")
				clone.getAtOffset = gopurs_runtime.RecordGet(orig, "getAtOffset")
				clone.read = gopurs_runtime.RecordGet(orig, "read")
				clone.readString = gopurs_runtime.RecordGet(orig, "readString")
				clone.setAtOffset = gopurs_runtime.RecordGet(orig, "setAtOffset")
				clone.size = gopurs_runtime.RecordGet(orig, "size")
				clone.slice = gopurs_runtime.RecordGet(orig, "slice")
				clone.thaw = gopurs_runtime.RecordGet(orig, "thaw")
				clone.toArray = gopurs_runtime.RecordGet(orig, "toArray")
				clone.toArrayBuffer = gopurs_runtime.RecordGet(orig, "toArrayBuffer")
				clone.toString = gopurs_runtime.RecordGet(orig, "toString")
				clone.unsafeFreeze = gopurs_runtime.RecordGet(orig, "unsafeFreeze")
				clone.unsafeThaw = gopurs_runtime.RecordGet(orig, "unsafeThaw")
				clone.write = gopurs_runtime.RecordGet(orig, "write")
				clone.writeString = gopurs_runtime.RecordGet(orig, "writeString")
				return clone
			}()))}
		})
	})
	return cache_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957
}

var cache_Node_Buffer_Class_writeString gopurs_runtime.Value
var once_Node_Buffer_Class_writeString sync.Once

func Get_Node_Buffer_Class_writeString() gopurs_runtime.Value {
	once_Node_Buffer_Class_writeString.Do(func() {
		cache_Node_Buffer_Class_writeString = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_writeString(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_writeString
}

var cache_Node_Buffer_Class_write gopurs_runtime.Value
var once_Node_Buffer_Class_write sync.Once

func Get_Node_Buffer_Class_write() gopurs_runtime.Value {
	once_Node_Buffer_Class_write.Do(func() {
		cache_Node_Buffer_Class_write = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_write(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_write
}

var cache_Node_Buffer_Class_unsafeThaw gopurs_runtime.Value
var once_Node_Buffer_Class_unsafeThaw sync.Once

func Get_Node_Buffer_Class_unsafeThaw() gopurs_runtime.Value {
	once_Node_Buffer_Class_unsafeThaw.Do(func() {
		cache_Node_Buffer_Class_unsafeThaw = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_unsafeThaw(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_unsafeThaw
}

var cache_Node_Buffer_Class_unsafeFreeze gopurs_runtime.Value
var once_Node_Buffer_Class_unsafeFreeze sync.Once

func Get_Node_Buffer_Class_unsafeFreeze() gopurs_runtime.Value {
	once_Node_Buffer_Class_unsafeFreeze.Do(func() {
		cache_Node_Buffer_Class_unsafeFreeze = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_unsafeFreeze(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_unsafeFreeze
}

var cache_Node_Buffer_Class_toString gopurs_runtime.Value
var once_Node_Buffer_Class_toString sync.Once

func Get_Node_Buffer_Class_toString() gopurs_runtime.Value {
	once_Node_Buffer_Class_toString.Do(func() {
		cache_Node_Buffer_Class_toString = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_toString(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_toString
}

var cache_Node_Buffer_Class_toArrayBuffer gopurs_runtime.Value
var once_Node_Buffer_Class_toArrayBuffer sync.Once

func Get_Node_Buffer_Class_toArrayBuffer() gopurs_runtime.Value {
	once_Node_Buffer_Class_toArrayBuffer.Do(func() {
		cache_Node_Buffer_Class_toArrayBuffer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_toArrayBuffer(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_toArrayBuffer
}

var cache_Node_Buffer_Class_toArray gopurs_runtime.Value
var once_Node_Buffer_Class_toArray sync.Once

func Get_Node_Buffer_Class_toArray() gopurs_runtime.Value {
	once_Node_Buffer_Class_toArray.Do(func() {
		cache_Node_Buffer_Class_toArray = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_toArray(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_toArray
}

var cache_Node_Buffer_Class_thaw gopurs_runtime.Value
var once_Node_Buffer_Class_thaw sync.Once

func Get_Node_Buffer_Class_thaw() gopurs_runtime.Value {
	once_Node_Buffer_Class_thaw.Do(func() {
		cache_Node_Buffer_Class_thaw = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_thaw(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_thaw
}

var cache_Node_Buffer_Class_slice gopurs_runtime.Value
var once_Node_Buffer_Class_slice sync.Once

func Get_Node_Buffer_Class_slice() gopurs_runtime.Value {
	once_Node_Buffer_Class_slice.Do(func() {
		cache_Node_Buffer_Class_slice = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_slice(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_slice
}

var cache_Node_Buffer_Class_size gopurs_runtime.Value
var once_Node_Buffer_Class_size sync.Once

func Get_Node_Buffer_Class_size() gopurs_runtime.Value {
	once_Node_Buffer_Class_size.Do(func() {
		cache_Node_Buffer_Class_size = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_size(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_size
}

var cache_Node_Buffer_Class_setAtOffset gopurs_runtime.Value
var once_Node_Buffer_Class_setAtOffset sync.Once

func Get_Node_Buffer_Class_setAtOffset() gopurs_runtime.Value {
	once_Node_Buffer_Class_setAtOffset.Do(func() {
		cache_Node_Buffer_Class_setAtOffset = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_setAtOffset(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_setAtOffset
}

var cache_Node_Buffer_Class_readString gopurs_runtime.Value
var once_Node_Buffer_Class_readString sync.Once

func Get_Node_Buffer_Class_readString() gopurs_runtime.Value {
	once_Node_Buffer_Class_readString.Do(func() {
		cache_Node_Buffer_Class_readString = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_readString(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_readString
}

var cache_Node_Buffer_Class_read gopurs_runtime.Value
var once_Node_Buffer_Class_read sync.Once

func Get_Node_Buffer_Class_read() gopurs_runtime.Value {
	once_Node_Buffer_Class_read.Do(func() {
		cache_Node_Buffer_Class_read = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_read(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_read
}

var cache_Node_Buffer_Class_getAtOffset gopurs_runtime.Value
var once_Node_Buffer_Class_getAtOffset sync.Once

func Get_Node_Buffer_Class_getAtOffset() gopurs_runtime.Value {
	once_Node_Buffer_Class_getAtOffset.Do(func() {
		cache_Node_Buffer_Class_getAtOffset = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_getAtOffset(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_getAtOffset
}

var cache_Node_Buffer_Class_fromString gopurs_runtime.Value
var once_Node_Buffer_Class_fromString sync.Once

func Get_Node_Buffer_Class_fromString() gopurs_runtime.Value {
	once_Node_Buffer_Class_fromString.Do(func() {
		cache_Node_Buffer_Class_fromString = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_fromString(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_fromString
}

var cache_Node_Buffer_Class_fromArrayBuffer gopurs_runtime.Value
var once_Node_Buffer_Class_fromArrayBuffer sync.Once

func Get_Node_Buffer_Class_fromArrayBuffer() gopurs_runtime.Value {
	once_Node_Buffer_Class_fromArrayBuffer.Do(func() {
		cache_Node_Buffer_Class_fromArrayBuffer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_fromArrayBuffer(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_fromArrayBuffer
}

var cache_Node_Buffer_Class_fromArray gopurs_runtime.Value
var once_Node_Buffer_Class_fromArray sync.Once

func Get_Node_Buffer_Class_fromArray() gopurs_runtime.Value {
	once_Node_Buffer_Class_fromArray.Do(func() {
		cache_Node_Buffer_Class_fromArray = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_fromArray(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_fromArray
}

var cache_Node_Buffer_Class_freeze gopurs_runtime.Value
var once_Node_Buffer_Class_freeze sync.Once

func Get_Node_Buffer_Class_freeze() gopurs_runtime.Value {
	once_Node_Buffer_Class_freeze.Do(func() {
		cache_Node_Buffer_Class_freeze = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_freeze(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_freeze
}

var cache_Node_Buffer_Class_fill gopurs_runtime.Value
var once_Node_Buffer_Class_fill sync.Once

func Get_Node_Buffer_Class_fill() gopurs_runtime.Value {
	once_Node_Buffer_Class_fill.Do(func() {
		cache_Node_Buffer_Class_fill = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_fill(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_fill
}

var cache_Node_Buffer_Class_create gopurs_runtime.Value
var once_Node_Buffer_Class_create sync.Once

func Get_Node_Buffer_Class_create() gopurs_runtime.Value {
	once_Node_Buffer_Class_create.Do(func() {
		cache_Node_Buffer_Class_create = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_create(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_create
}

var cache_Node_Buffer_Class_copy gopurs_runtime.Value
var once_Node_Buffer_Class_copy sync.Once

func Get_Node_Buffer_Class_copy() gopurs_runtime.Value {
	once_Node_Buffer_Class_copy.Do(func() {
		cache_Node_Buffer_Class_copy = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_copy(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_copy
}

var cache_Node_Buffer_Class_concat_prime_ gopurs_runtime.Value
var once_Node_Buffer_Class_concat_prime_ sync.Once

func Get_Node_Buffer_Class_concat_prime_() gopurs_runtime.Value {
	once_Node_Buffer_Class_concat_prime_.Do(func() {
		cache_Node_Buffer_Class_concat_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_concat_prime_(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_concat_prime_
}

var cache_Node_Buffer_Class_concat gopurs_runtime.Value
var once_Node_Buffer_Class_concat sync.Once

func Get_Node_Buffer_Class_concat() gopurs_runtime.Value {
	once_Node_Buffer_Class_concat.Do(func() {
		cache_Node_Buffer_Class_concat = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Class_concat(gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Node_Buffer_Class_concat
}

type Constructor_Node_Buffer_Class_MutableBuffer[T_buf any, T_m any] struct {
	Rc  uint32
	V0  gopurs_runtime.Value
	V1  gopurs_runtime.Value
	V2  gopurs_runtime.Value
	V3  gopurs_runtime.Value
	V4  gopurs_runtime.Value
	V5  gopurs_runtime.Value
	V6  gopurs_runtime.Value
	V7  gopurs_runtime.Value
	V8  gopurs_runtime.Value
	V9  gopurs_runtime.Value
	V10 gopurs_runtime.Value
	V11 gopurs_runtime.Value
	V12 gopurs_runtime.Value
	V13 gopurs_runtime.Value
	V14 gopurs_runtime.Value
	V15 gopurs_runtime.Value
	V16 gopurs_runtime.Value
	V17 gopurs_runtime.Value
	V18 gopurs_runtime.Value
	V19 gopurs_runtime.Value
	V20 gopurs_runtime.Value
	V21 gopurs_runtime.Value
	V22 gopurs_runtime.Value
	V23 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004e006f00640065005f004200750066006600650072005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004e006f00640065005f004200750066006600650072005f0043006c006100730073005f004d0075007400610062006c0065004200750066006600650072
func init() {
	gopurs_runtime.StructGetters[1692970589] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "concat":
			return gopurs_runtime.Box(c.V1)
		case "concat'":
			return gopurs_runtime.Box(c.V2)
		case "copy":
			return gopurs_runtime.Box(c.V3)
		case "create":
			return gopurs_runtime.Box(c.V4)
		case "fill":
			return gopurs_runtime.Box(c.V5)
		case "freeze":
			return gopurs_runtime.Box(c.V6)
		case "fromArray":
			return gopurs_runtime.Box(c.V7)
		case "fromArrayBuffer":
			return gopurs_runtime.Box(c.V8)
		case "fromString":
			return gopurs_runtime.Box(c.V9)
		case "getAtOffset":
			return gopurs_runtime.Box(c.V10)
		case "read":
			return gopurs_runtime.Box(c.V11)
		case "readString":
			return gopurs_runtime.Box(c.V12)
		case "setAtOffset":
			return gopurs_runtime.Box(c.V13)
		case "size":
			return gopurs_runtime.Box(c.V14)
		case "slice":
			return gopurs_runtime.Box(c.V15)
		case "thaw":
			return gopurs_runtime.Box(c.V16)
		case "toArray":
			return gopurs_runtime.Box(c.V17)
		case "toArrayBuffer":
			return gopurs_runtime.Box(c.V18)
		case "toString":
			return gopurs_runtime.Box(c.V19)
		case "unsafeFreeze":
			return gopurs_runtime.Box(c.V20)
		case "unsafeThaw":
			return gopurs_runtime.Box(c.V21)
		case "write":
			return gopurs_runtime.Box(c.V22)
		case "writeString":
			return gopurs_runtime.Box(c.V23)
		default:
			panic("Key not found in dictionary Constructor_Node_Buffer_Class_MutableBuffer: " + key)
		}
	}
}

func Call_Node_Buffer_Class_MutableBuffer_dollar_Dict(x_0_loop struct {
	Monad0                               gopurs_runtime.Value
	concat                               gopurs_runtime.Value
	gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
	copy                                 gopurs_runtime.Value
	create                               gopurs_runtime.Value
	fill                                 gopurs_runtime.Value
	freeze                               gopurs_runtime.Value
	fromArray                            gopurs_runtime.Value
	fromArrayBuffer                      gopurs_runtime.Value
	fromString                           gopurs_runtime.Value
	getAtOffset                          gopurs_runtime.Value
	read                                 gopurs_runtime.Value
	readString                           gopurs_runtime.Value
	setAtOffset                          gopurs_runtime.Value
	size                                 gopurs_runtime.Value
	slice                                gopurs_runtime.Value
	thaw                                 gopurs_runtime.Value
	toArray                              gopurs_runtime.Value
	toArrayBuffer                        gopurs_runtime.Value
	toString                             gopurs_runtime.Value
	unsafeFreeze                         gopurs_runtime.Value
	unsafeThaw                           gopurs_runtime.Value
	write                                gopurs_runtime.Value
	writeString                          gopurs_runtime.Value
	/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
}) *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Monad0                               gopurs_runtime.Value
		concat                               gopurs_runtime.Value
		gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
		copy                                 gopurs_runtime.Value
		create                               gopurs_runtime.Value
		fill                                 gopurs_runtime.Value
		freeze                               gopurs_runtime.Value
		fromArray                            gopurs_runtime.Value
		fromArrayBuffer                      gopurs_runtime.Value
		fromString                           gopurs_runtime.Value
		getAtOffset                          gopurs_runtime.Value
		read                                 gopurs_runtime.Value
		readString                           gopurs_runtime.Value
		setAtOffset                          gopurs_runtime.Value
		size                                 gopurs_runtime.Value
		slice                                gopurs_runtime.Value
		thaw                                 gopurs_runtime.Value
		toArray                              gopurs_runtime.Value
		toArrayBuffer                        gopurs_runtime.Value
		toString                             gopurs_runtime.Value
		unsafeFreeze                         gopurs_runtime.Value
		unsafeThaw                           gopurs_runtime.Value
		write                                gopurs_runtime.Value
		writeString                          gopurs_runtime.Value
		/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict([]string{"Monad0", "concat", "concat'", "copy", "create", "fill", "freeze", "fromArray", "fromArrayBuffer", "fromString", "getAtOffset", "read", "readString", "setAtOffset", "size", "slice", "thaw", "toArray", "toArrayBuffer", "toString", "unsafeFreeze", "unsafeThaw", "write", "writeString"}, []gopurs_runtime.Value{orig.Monad0, orig.concat, orig.gopurs_field_99_111_110_99_97_116_39, orig.copy, orig.create, orig.fill, orig.freeze, orig.fromArray, orig.fromArrayBuffer, orig.fromString, orig.getAtOffset, orig.read, orig.readString, orig.setAtOffset, orig.size, orig.slice, orig.thaw, orig.toArray, orig.toArrayBuffer, orig.toString, orig.unsafeFreeze, orig.unsafeThaw, orig.write, orig.writeString})
	}())
}

func Call_Node_Buffer_Class_MutableBuffer_dollar_Dict__1994236957(x_0_loop struct {
	Monad0                               gopurs_runtime.Value
	concat                               gopurs_runtime.Value
	gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
	copy                                 gopurs_runtime.Value
	create                               gopurs_runtime.Value
	fill                                 gopurs_runtime.Value
	freeze                               gopurs_runtime.Value
	fromArray                            gopurs_runtime.Value
	fromArrayBuffer                      gopurs_runtime.Value
	fromString                           gopurs_runtime.Value
	getAtOffset                          gopurs_runtime.Value
	read                                 gopurs_runtime.Value
	readString                           gopurs_runtime.Value
	setAtOffset                          gopurs_runtime.Value
	size                                 gopurs_runtime.Value
	slice                                gopurs_runtime.Value
	thaw                                 gopurs_runtime.Value
	toArray                              gopurs_runtime.Value
	toArrayBuffer                        gopurs_runtime.Value
	toString                             gopurs_runtime.Value
	unsafeFreeze                         gopurs_runtime.Value
	unsafeThaw                           gopurs_runtime.Value
	write                                gopurs_runtime.Value
	writeString                          gopurs_runtime.Value
	/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
}) *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] {
MutableBuffer_dollar_Dict__1994236957:
	for {
		if false {
			continue MutableBuffer_dollar_Dict__1994236957
		}
		var x_0 struct {
			Monad0                               gopurs_runtime.Value
			concat                               gopurs_runtime.Value
			gopurs_field_99_111_110_99_97_116_39 gopurs_runtime.Value
			copy                                 gopurs_runtime.Value
			create                               gopurs_runtime.Value
			fill                                 gopurs_runtime.Value
			freeze                               gopurs_runtime.Value
			fromArray                            gopurs_runtime.Value
			fromArrayBuffer                      gopurs_runtime.Value
			fromString                           gopurs_runtime.Value
			getAtOffset                          gopurs_runtime.Value
			read                                 gopurs_runtime.Value
			readString                           gopurs_runtime.Value
			setAtOffset                          gopurs_runtime.Value
			size                                 gopurs_runtime.Value
			slice                                gopurs_runtime.Value
			thaw                                 gopurs_runtime.Value
			toArray                              gopurs_runtime.Value
			toArrayBuffer                        gopurs_runtime.Value
			toString                             gopurs_runtime.Value
			unsafeFreeze                         gopurs_runtime.Value
			unsafeThaw                           gopurs_runtime.Value
			write                                gopurs_runtime.Value
			writeString                          gopurs_runtime.Value
			/*gopurs:record-v1:24:004d006f006e006100640030,0063006f006e006300610074,0063006f006e0063006100740027,0063006f00700079,006300720065006100740065,00660069006c006c,0066007200650065007a0065,00660072006f006d00410072007200610079,00660072006f006d00410072007200610079004200750066006600650072,00660072006f006d0053007400720069006e0067,00670065007400410074004f00660066007300650074,0072006500610064,00720065006100640053007400720069006e0067,00730065007400410074004f00660066007300650074,00730069007a0065,0073006c006900630065,0074006800610077,0074006f00410072007200610079,0074006f00410072007200610079004200750066006600650072,0074006f0053007400720069006e0067,0075006e00730061006600650046007200650065007a0065,0075006e00730061006600650054006800610077,00770072006900740065,007700720069007400650053007400720069006e0067*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict([]string{"Monad0", "concat", "concat'", "copy", "create", "fill", "freeze", "fromArray", "fromArrayBuffer", "fromString", "getAtOffset", "read", "readString", "setAtOffset", "size", "slice", "thaw", "toArray", "toArrayBuffer", "toString", "unsafeFreeze", "unsafeThaw", "write", "writeString"}, []gopurs_runtime.Value{orig.Monad0, orig.concat, orig.gopurs_field_99_111_110_99_97_116_39, orig.copy, orig.create, orig.fill, orig.freeze, orig.fromArray, orig.fromArrayBuffer, orig.fromString, orig.getAtOffset, orig.read, orig.readString, orig.setAtOffset, orig.size, orig.slice, orig.thaw, orig.toArray, orig.toArrayBuffer, orig.toString, orig.unsafeFreeze, orig.unsafeThaw, orig.write, orig.writeString})
		}())
	}
}

func Call_Node_Buffer_Class_writeString(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V23
}

func Call_Node_Buffer_Class_write(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V22
}

func Call_Node_Buffer_Class_unsafeThaw(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V21
}

func Call_Node_Buffer_Class_unsafeFreeze(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V20
}

func Call_Node_Buffer_Class_toString(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V19
}

func Call_Node_Buffer_Class_toArrayBuffer(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V18
}

func Call_Node_Buffer_Class_toArray(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V17
}

func Call_Node_Buffer_Class_thaw(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V16
}

func Call_Node_Buffer_Class_slice(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V15
}

func Call_Node_Buffer_Class_size(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V14
}

func Call_Node_Buffer_Class_setAtOffset(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V13
}

func Call_Node_Buffer_Class_readString(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V12
}

func Call_Node_Buffer_Class_read(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V11
}

func Call_Node_Buffer_Class_getAtOffset(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V10
}

func Call_Node_Buffer_Class_fromString(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V9
}

func Call_Node_Buffer_Class_fromArrayBuffer(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V8
}

func Call_Node_Buffer_Class_fromArray(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V7
}

func Call_Node_Buffer_Class_freeze(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V6
}

func Call_Node_Buffer_Class_fill(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V5
}

func Call_Node_Buffer_Class_create(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V4
}

func Call_Node_Buffer_Class_copy(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Node_Buffer_Class_concat_prime_(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Node_Buffer_Class_concat(dict_0_loop *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}
