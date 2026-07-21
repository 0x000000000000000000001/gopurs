package main

import (
	"fmt"
	"time"
)

// --- 1. VERSION STATIQUE ---
// Le compilateur sait que tout est entier. Les variables restent sur la pile (Stack).
func sumStatic(limit int) int {
	acc := 0
	for i := 0; i < limit; i++ {
		acc += i
	}
	return acc
}

// --- 2. VERSION DYNAMIQUE NAÏVE (any) ---
// Traduction littérale avec le type `any` de Go. Force l'allocation sur le Heap.
func sumDynamicNaive(limitAny any) any {
	acc := any(0)
	limit := limitAny.(int)

	for i := any(0); i.(int) < limit; i = any(i.(int) + 1) {
		acc = any(acc.(int) + i.(int))
	}
	return acc
}

// --- 3. VERSION DYNAMIQUE OPTIMISÉE (Tagged Type / Value Object) ---
// On n'utilise plus `any`. On utilise notre propre structure universelle.
// La structure reste allouée sur la pile (Stack), le GC ne se réveille jamais.

const TypeInt = 1

type Value struct {
	Type   uint8 // 1 = Int, 2 = Float, 3 = String...
	IntVal int
	// FloatVal float64
	// PtrVal unsafe.Pointer
}

func sumDynamicOptim(limitVal Value) Value {
	acc := Value{Type: TypeInt, IntVal: 0}
	limit := limitVal.IntVal

	i := Value{Type: TypeInt, IntVal: 0}
	for i.IntVal < limit {
		// On simule une addition dynamique qui renvoie une nouvelle Value
		acc = Value{Type: TypeInt, IntVal: acc.IntVal + i.IntVal}
		i.IntVal++
	}
	return acc
}

func main() {
	iterations := 1_000_000_000

	fmt.Println("Démarrage du benchmark (1 milliard d'itérations)...")
	fmt.Println("-------------------------------------------------")

	// 1. Test Statique
	startStatic := time.Now()
	resStatic := sumStatic(iterations)
	timeStatic := time.Since(startStatic)
	fmt.Printf("⏱️ Statique       : %d (en %v)\n", resStatic, timeStatic)

	// 2. Test Dynamique Naïf
	startDynamicNaive := time.Now()
	resDynamicNaive := sumDynamicNaive(iterations)
	timeDynamicNaive := time.Since(startDynamicNaive)
	fmt.Printf("⏱️ Dynamique Naïf : %v (en %v)\n", resDynamicNaive, timeDynamicNaive)

	// 3. Test Dynamique Optimisé (Value Struct)
	startDynamicOptim := time.Now()
	resDynamicOptim := sumDynamicOptim(Value{Type: TypeInt, IntVal: iterations})
	timeDynamicOptim := time.Since(startDynamicOptim)
	fmt.Printf("⏱️ Dynamique Optim: %d (en %v)\n", resDynamicOptim.IntVal, timeDynamicOptim)

	fmt.Println("-------------------------------------------------")

	ratioNaive := float64(timeDynamicNaive.Nanoseconds()) / float64(timeStatic.Nanoseconds())
	ratioOptim := float64(timeDynamicOptim.Nanoseconds()) / float64(timeStatic.Nanoseconds())
	
	fmt.Printf("📊 Bilan :\n")
	fmt.Printf(" - La version Naïve ('any') est %.1f fois plus lente que le natif.\n", ratioNaive)
	fmt.Printf(" - La version Optim (Struct) est %.1f fois plus lente que le natif.\n", ratioOptim)
}
