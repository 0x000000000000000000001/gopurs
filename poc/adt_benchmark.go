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
// 1. CURRENT: Records (RecordData5 for 5 fields)
// ==========================================

type RecordData5 struct {
	K0, K1, K2, K3, K4 string
	V0, V1, V2, V3, V4 Value
}

func RecordDict5(k0, k1, k2, k3, k4 string, v0, v1, v2, v3, v4 Value) Value {
	return Value{Type: TypeRecord, PtrVal: &RecordData5{k0, k1, k2, k3, k4, v0, v1, v2, v3, v4}}
}

// In real compiler, _tag is often K0, but other fields are accessed by names
func RecordGet(obj Value, key string) Value {
	switch r := obj.PtrVal.(type) {
	case *RecordData5:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
		if r.K4 == key { return r.V4 }
	}
	return Value{}
}

func runRecordSimulation(iterations int) int {
	var current Value
	current = RecordDict5("_tag", "value0", "value1", "value2", "value3", Value{StrVal: "T"}, Value{}, Value{}, Value{}, Value{})

	sum := 0
	for i := 0; i < iterations; i++ {
		tag := RecordGet(current, "_tag").StrVal
		if tag == "T" {
			sum++
			_ = RecordGet(current, "value0")
			_ = RecordGet(current, "value1")
			_ = RecordGet(current, "value2")
			_ = RecordGet(current, "value3")
			
			current = RecordDict5("_tag", "value0", "value1", "value2", "value3", Value{StrVal: "T"}, Value{}, Value{}, Value{}, Value{})
		}
	}
	return sum
}

// ==========================================
// 2. PROPOSED: Struct Constructors (No keys)
// ==========================================

type ConstructorData4 struct {
	V0, V1, V2, V3 Value
}

func Constructor4(tag string, v0, v1, v2, v3 Value) Value {
	return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData4{v0, v1, v2, v3}}
}

// Direct field access via type assertion (compiler knows indices)
func runStructSimulation(iterations int) int {
	var current Value
	current = Constructor4("T", Value{}, Value{}, Value{}, Value{})

	sum := 0
	for i := 0; i < iterations; i++ {
		tag := current.StrVal // O(1) directly from Value!
		if tag == "T" {
			sum++
			r := current.PtrVal.(*ConstructorData4)
			_ = r.V0
			_ = r.V1
			_ = r.V2
			_ = r.V3
			
			current = Constructor4("T", Value{}, Value{}, Value{}, Value{})
		}
	}
	return sum
}

func main() {
	iterations := 10_000_000

	fmt.Println("Démarrage du benchmark d'allocation ADT (10M itérations)...")
	fmt.Println("-------------------------------------------------")

	start1 := time.Now()
	res1 := runRecordSimulation(iterations)
	time1 := time.Since(start1)
	fmt.Printf("⏱️ Approche actuelle (RecordDict5) : en %v (res %d)\n", time1, res1)

	start2 := time.Now()
	res2 := runStructSimulation(iterations)
	time2 := time.Since(start2)
	fmt.Printf("⏱️ Nouvelle approche (Constructor struct direct) : en %v (res %d)\n", time2, res2)

	fmt.Println("-------------------------------------------------")
	ratio := float64(time1.Nanoseconds()) / float64(time2.Nanoseconds())
	fmt.Printf("📊 Bilan :\n")
	fmt.Printf(" - La nouvelle approche est %.1f fois plus rapide !\n", ratio)
}
