package main

import (
	"cgo/src/hello"
)

func main() {
	hello.Hello("Tom", "Jack")
	Say("Tom", "what a fucking day!")
	Say("Jack", "yes.")
	Bye("Jack", "Tom")
}
