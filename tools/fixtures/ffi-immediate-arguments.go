package fixture

var trace string

func Identity(value any) func() any {
	trace += "F"
	return func() any {
		trace += "E"
		return value
	}
}

func Snapshot() string { return trace }
func Reset()           { trace = "" }
