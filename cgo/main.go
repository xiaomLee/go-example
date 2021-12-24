package main

import (
	"cgo/helloc"
	"cgo/say"
)

func main() {
	helloc.Hello("Tom")
	say.Something("How do you do?")
	say.Bye("Tom")
}
