package Main

import "gopurs/output/gopurs_runtime"

func ReadWhileDraining(avar gopurs_runtime.Value) func() bool {
	return func() bool {
		identity := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { return value })
		calls := 0
		second := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				if value.IntVal == 42 {
					calls++
				}
				return gopurs_runtime.Any(nil)
			})
		})
		first := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				calls++
				// Queue another read while the first read's callback is draining.
				gopurs_runtime.Apply(Effect_AVar__ReadVar(identity, identity, avar, second), gopurs_runtime.Any(nil))
				return gopurs_runtime.Any(nil)
			})
		})
		gopurs_runtime.Apply(Effect_AVar__ReadVar(identity, identity, avar, first), gopurs_runtime.Any(nil))
		return calls == 2
	}
}
