package main

import (
	"fmt"
	"protobuf/pb"
)

func main() {
	t1 := pb.Test{
		Num1: 10,
		Num2: 1073741824,
	}
	t2 := pb.Test2{
		Name: "Steven",
	}

	t3 := pb.Test3{
		Num: []int32{10, 100, 1000},
	}

	var encoded []byte
	var err error
	if encoded, err = t1.XXX_Marshal(nil, false); err != nil {
		panic(err)
	}
	fmt.Printf("encoded: %x\n", encoded)
	// output
	// encoded: 080a1500000040

	if encoded, err = t2.XXX_Marshal(nil, false); err != nil {
		panic(err)
	}
	fmt.Printf("encoded: %x\n", encoded)
	// output
	// encoded: 0a0653746576656e

	if encoded, err = t3.XXX_Marshal(nil, false); err != nil {
		panic(err)
	}
	fmt.Printf("encoded: %x\n", encoded)
	// output
	// encoded: 0a040a64e807
}
