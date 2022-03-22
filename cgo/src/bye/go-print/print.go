package main

import "C"
import "fmt"

//go:generate go build -buildmode=c-shared -o ../print.so print.go

//export  GoPrint
func GoPrint(from, to *C.char) {
	fmt.Printf("%s: Bye %s.\n", C.GoString(from), C.GoString(to))
}

// main函数是必须的 有main函数才能让cgo编译器去把包编译成C的库
func main() {
}

//1、第11行 这里go代码中的main函数是必须的，有main函数才能让cgo编译器去把包编译成c的库
//
//2、第3行 import “C”是必须的，如果没有import “C” 将只会build出一个.a文件，而缺少.h文件
//
//3、第6行 //export GoPrint  这里的SayBye要和下面的的go函数名一致，并且下面一行即为要导出的go函数
//
//4、命令执行完毕后会生成两个文件 nautilus.a nautilus.h
