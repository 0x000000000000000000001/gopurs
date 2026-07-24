package main

import (
	"fmt"
	"unsafe"
)

type ConstructorData4 struct {
	V0, V1, V2, V3 int64
}

func main() {
	c := &ConstructorData4{10, 20, 30, 40}
	ptr := unsafe.Pointer(c)
	
	arr := (*[100]int64)(ptr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
}
