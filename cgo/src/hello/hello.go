// Package hello show how to call c function by c source file
package hello

/*
#include <stdlib.h>
#include "hello.h"
*/
import "C"
import (
	"unsafe"
)

// Hello must call C.free for C variable
func Hello(from, to string) {
	cfrom := C.CString(from)
	cto := C.CString(to)
	defer C.free(unsafe.Pointer(cfrom))
	defer C.free(unsafe.Pointer(cto))
	C.SayHello(cfrom, cto)
}
