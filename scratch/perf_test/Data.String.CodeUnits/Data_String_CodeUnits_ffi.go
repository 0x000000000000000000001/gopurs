package Data_String_CodeUnits

import (
	"strings"
	"gopurs/output/gopurs_runtime"
)

var FromCharArray = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	arr := a.PtrVal.([]gopurs_runtime.Value)
	var b strings.Builder
	for _, v := range arr {
		b.WriteString(v.StrVal)
	}
	return gopurs_runtime.Str(b.String())
})

var ToCharArray = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	str := s.StrVal
	arr := make([]gopurs_runtime.Value, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = gopurs_runtime.Str(string(str[i]))
	}
	return gopurs_runtime.Array(arr)
})

var Singleton = gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
	return c
})

var X_CharAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
				idx := int(i.IntVal)
				str := s.StrVal
				if idx >= 0 && idx < len(str) {
					return gopurs_runtime.Apply(just, gopurs_runtime.Str(string(str[idx])))
				}
				return nothing
			})
		})
	})
})

var X_ToChar = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
			str := s.StrVal
			if len(str) == 1 {
				return gopurs_runtime.Apply(just, s)
			}
			return nothing
		})
	})
})

var Length = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Int(len(s.StrVal))
})

var CountPrefix = gopurs_runtime.Func(func(p gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
		str := s.StrVal
		i := 0
		for i < len(str) {
			charVal := gopurs_runtime.Str(string(str[i]))
			if gopurs_runtime.Apply(p, charVal).IntVal != 0 {
				i++
			} else {
				break
			}
		}
		return gopurs_runtime.Int(i)
	})
})

var X_IndexOf = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
				idx := strings.Index(s.StrVal, x.StrVal)
				if idx == -1 {
					return nothing
				}
				return gopurs_runtime.Apply(just, gopurs_runtime.Int(idx))
			})
		})
	})
})

var X_IndexOfStartingAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(startAt gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
					startIdx := int(startAt.IntVal)
					str := s.StrVal
					if startIdx < 0 || startIdx > len(str) {
						return nothing
					}
					idx := strings.Index(str[startIdx:], x.StrVal)
					if idx == -1 {
						return nothing
					}
					return gopurs_runtime.Apply(just, gopurs_runtime.Int(idx+startIdx))
				})
			})
		})
	})
})

var X_LastIndexOf = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
				idx := strings.LastIndex(s.StrVal, x.StrVal)
				if idx == -1 {
					return nothing
				}
				return gopurs_runtime.Apply(just, gopurs_runtime.Int(idx))
			})
		})
	})
})

var X_LastIndexOfStartingAt = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(startAt gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
					startIdx := int(startAt.IntVal)
					str := s.StrVal
					if startIdx < 0 || startIdx >= len(str) {
						return nothing
					}
					idx := strings.LastIndex(str[:startIdx+len(x.StrVal)], x.StrVal)
					if idx == -1 {
						return nothing
					}
					return gopurs_runtime.Apply(just, gopurs_runtime.Int(idx))
				})
			})
		})
	})
})

var Take = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
		str := s.StrVal
		idx := int(n.IntVal)
		if idx < 0 { idx = 0 }
		if idx > len(str) { idx = len(str) }
		return gopurs_runtime.Str(str[:idx])
	})
})

var Drop = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
		str := s.StrVal
		idx := int(n.IntVal)
		if idx < 0 { idx = 0 }
		if idx > len(str) { idx = len(str) }
		return gopurs_runtime.Str(str[idx:])
	})
})

var Slice = gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
			str := s.StrVal
			start := int(b.IntVal)
			end := int(e.IntVal)
			if start < 0 { start = len(str) + start }
			if end < 0 { end = len(str) + end }
			if start < 0 { start = 0 }
			if end < 0 { end = 0 }
			if start > len(str) { start = len(str) }
			if end > len(str) { end = len(str) }
			if start > end { return gopurs_runtime.Str("") }
			return gopurs_runtime.Str(str[start:end])
		})
	})
})

var SplitAt = gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
		str := s.StrVal
		idx := int(i.IntVal)
		if idx < 0 { idx = 0 }
		if idx > len(str) { idx = len(str) }
		rec := make(map[string]gopurs_runtime.Value)
		rec["before"] = gopurs_runtime.Str(str[:idx])
		rec["after"] = gopurs_runtime.Str(str[idx:])
		return gopurs_runtime.Record(rec)
	})
})
