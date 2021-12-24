// Package helloc example show how to call c function by c source file
package helloc

/*
#include <stdlib.h>
#include "hello.h"
*/
import "C"
import "unsafe"

// Hello must call C.free for C variable
func Hello(s string) {
	cname := C.CString(s)
	defer C.free(unsafe.Pointer(cname))
	C.SayHello(cname)
}
