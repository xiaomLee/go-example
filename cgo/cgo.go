// #cgo CFLAGS: -I.     // 头文件的位置，相对于源文件是当前目录，所以是 ./include，头文件在多个目录时写多个 #cgo CFLAGS
// #cgo LDFLAGS: -L../lib -lany -Wl,-rpath,lib  // 从哪里加载动态库，位置与文件名，-lany 加载 libany.so 文件
// #cgo LDFLAGS: -lany  //不指定从哪加载动态库，只指定需要加载的库，编译时需export CGO_LDFLAGS =-L$(PWD)/libxxxa -L$(PWD)/libxxxb
// 此方式适用需动态指定加载路径的编译,同时运行时需export LD_LIBRARY_PATH=$(PWD)/libxxxa:$(PWD)/libxxxb
package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lsay -lbye -lprint
#include <stdlib.h>
#include "say/say.h"
#include "bye/bye.h"
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

func Bye(from, to string) {
	cfrom := C.CString(from)
	cto := C.CString(to)
	defer C.free(unsafe.Pointer(cfrom))
	defer C.free(unsafe.Pointer(cto))
	C.SayBye(cfrom, cto)
}
