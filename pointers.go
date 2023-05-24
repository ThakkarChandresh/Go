package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := 10

	p := &a

	// single pointer

	fmt.Println(*p)

	// double pointers

	var b **int = &p

	fmt.Println(**b)

	fmt.Println(unsafe.Sizeof(p))

}
