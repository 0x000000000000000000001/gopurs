package Record_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
)



func Get_unsafeDelete() gopurs_runtime.Value {
	return UnsafeDelete
}

func Get_unsafeGet() gopurs_runtime.Value {
	return UnsafeGet
}

func Get_unsafeHas() gopurs_runtime.Value {
	return UnsafeHas
}

func Get_unsafeSet() gopurs_runtime.Value {
	return UnsafeSet
}
