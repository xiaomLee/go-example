package say_cgo

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L../output -lsay
#include <stdlib.h>
#include "say.h"
*/
import "C"
import "unsafe"

func Say(name, sth string) {
	cname := C.CString(name)
	csth := C.CString(sth)
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(csth))
	C.SaySomething(cname, csth)
}
