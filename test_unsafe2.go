package main

import (
	"fmt"
	"unsafe"
)

func main() {
	args := []int64{10, 20, 30, 40}
	ptr := unsafe.Pointer(&args[0])
	
	arr := (*[100]int64)(ptr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
}
