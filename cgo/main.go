package main

import (
	"cgo/helloc"
	"cgo/lib-src/bye"
	"cgo/say"
)

func main() {
	helloc.Hello("Tom")
	say.Something("How do you do?")
	bye.Bye("Tom")
}
