package main

import (
	"gopurs/output/Test1110"
	"gopurs/output/gopurs_runtime"
)

func main() {
	gopurs_runtime.Apply(Test1110.Main, gopurs_runtime.Value{})
}
