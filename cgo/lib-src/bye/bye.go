package bye

//#include <stdlib.h>
//#include "bye.h"
import "C"
import "unsafe"

func Bye(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(&cname))
	C.SayBye(cname)
}
