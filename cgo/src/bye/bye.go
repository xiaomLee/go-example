package bye

//#include <stdlib.h>
//#include "bye.h"
import "C"
import "unsafe"

func Bye(from, to string) {
	cfrom := C.CString(from)
	cto := C.CString(to)
	defer C.free(unsafe.Pointer(cfrom))
	defer C.free(unsafe.Pointer(cto))
	C.SayBye(cfrom, cto)
}
