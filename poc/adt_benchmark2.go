package main

import (
	"fmt"
	"time"
)

const (
	TypeInt         = 1
	TypeRecord      = 8
	TypeConstructor = 9
)

type Value struct {
	Type   uint8
	IntVal int
	StrVal string
	PtrVal interface{}
}

// ==========================================
// 1. Proposed struct + Runtime ConstructorGet
// ==========================================

type ConstructorData4 struct {
	V0, V1, V2, V3 Value
}

func Constructor4(tag string, v0, v1, v2, v3 Value) Value {
	return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData4{v0, v1, v2, v3}}
}

func ConstructorGet(obj Value, index int) Value {
	switch r := obj.PtrVal.(type) {
	case *ConstructorData4:
		switch index {
		case 0: return r.V0
		case 1: return r.V1
		case 2: return r.V2
		case 3: return r.V3
		}
	}
	return Value{}
}

func runStructGetSimulation(iterations int) int {
	var current Value
	current = Constructor4("T", Value{}, Value{}, Value{}, Value{})

	sum := 0
	for i := 0; i < iterations; i++ {
		tag := current.StrVal
		if tag == "T" {
			sum++
			_ = ConstructorGet(current, 0)
			_ = ConstructorGet(current, 1)
			_ = ConstructorGet(current, 2)
			_ = ConstructorGet(current, 3)
			
			current = Constructor4("T", Value{}, Value{}, Value{}, Value{})
		}
	}
	return sum
}

func main() {
	iterations := 10_000_000

	start2 := time.Now()
	res2 := runStructGetSimulation(iterations)
	time2 := time.Since(start2)
	fmt.Printf("⏱️ Nouvelle approche avec ConstructorGet() switch : en %v (res %d)\n", time2, res2)
}
