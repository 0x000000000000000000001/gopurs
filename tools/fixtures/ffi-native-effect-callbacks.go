package fixture

var trace string

func RunInt(action func() int) func() int {
	trace += "F"
	return func() int {
		trace += "B"
		value := action()
		trace += "A"
		return value
	}
}

func RunInt64(action func() int64) func() int64 {
	trace += "F"
	return func() int64 {
		trace += "B"
		value := action()
		trace += "A"
		return value
	}
}

func RunValue(action func() any) func() any {
	trace += "F"
	return func() any {
		trace += "B"
		value := action()
		trace += "A"
		return value
	}
}

func Tick(value int64) func() int64 {
	return func() int64 {
		trace += "T"
		return value
	}
}

func Reset()           { trace = "" }
func Snapshot() string { return trace }
