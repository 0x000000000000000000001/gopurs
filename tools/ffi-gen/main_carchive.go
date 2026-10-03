//go:build carchive

package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Lengths preserve embedded NULs so malformed sources reach the same parser and
// diagnostics as the native Go host. Rust owns neither Go memory nor Go pointers.
//
//export GopursParseFFI
func GopursParseFFI(content *C.uchar, contentLen C.size_t, prefix *C.uchar, prefixLen C.size_t, prepare C.int, failed *C.int) (result *C.char) {
	defer func() {
		if err := recover(); err != nil {
			*failed = 1
			result = C.CString(fmt.Sprint(err))
		}
	}()
	source := string(unsafe.Slice((*byte)(unsafe.Pointer(content)), int(contentLen)))
	var output string
	var err error
	if prepare != 0 {
		name := string(unsafe.Slice((*byte)(unsafe.Pointer(prefix)), int(prefixLen)))
		output, err = Prepare(name, source)
	} else {
		output, err = Extract(source)
	}
	if err != nil {
		*failed = 1
		return C.CString(err.Error())
	}
	*failed = 0
	return C.CString(output)
}

//export GopursFreeFFI
func GopursFreeFFI(result *C.char) {
	C.free(unsafe.Pointer(result))
}

func main() {}
