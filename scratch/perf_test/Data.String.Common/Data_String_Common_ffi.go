package Data_String_Common

import (
	"strings"
	"gopurs/output/gopurs_runtime"
)

var X_LocaleCompare = gopurs_runtime.Func(func(lt gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(eq gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(gt gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s2 gopurs_runtime.Value) gopurs_runtime.Value {
					cmp := strings.Compare(s1.StrVal, s2.StrVal)
					if cmp < 0 {
						return lt
					} else if cmp > 0 {
						return gt
					}
					return eq
				})
			})
		})
	})
})

var Replace = gopurs_runtime.Func(func(s1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(s3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(strings.Replace(s3.StrVal, s1.StrVal, s2.StrVal, 1))
		})
	})
})

var ReplaceAll = gopurs_runtime.Func(func(s1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(s3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(strings.ReplaceAll(s3.StrVal, s1.StrVal, s2.StrVal))
		})
	})
})

var Split = gopurs_runtime.Func(func(sep gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
		parts := strings.Split(s.StrVal, sep.StrVal)
		arr := make([]gopurs_runtime.Value, len(parts))
		for i, p := range parts {
			arr[i] = gopurs_runtime.Str(p)
		}
		return gopurs_runtime.Array(arr)
	})
})

var ToLower = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str(strings.ToLower(s.StrVal))
})

var ToUpper = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str(strings.ToUpper(s.StrVal))
})

var Trim = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str(strings.TrimSpace(s.StrVal))
})

var JoinWith = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
		arr := xs.PtrVal.([]gopurs_runtime.Value)
		var strs []string
		for _, v := range arr {
			strs = append(strs, v.StrVal)
		}
		return gopurs_runtime.Str(strings.Join(strs, s.StrVal))
	})
})
