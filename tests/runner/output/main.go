package main

import (
	"gopurs/output/Main"
	"gopurs/output/gopurs_runtime"
)

func main() {
	gopurs_runtime.Apply(Main.Main, gopurs_runtime.Value{})
}
