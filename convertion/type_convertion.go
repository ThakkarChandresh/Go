package main

import (
	"fmt"
	"strconv"
)

func main() {

	// NUMBER TO STRING
	s := string(99)

	fmt.Println(s)

	s1 := fmt.Sprintf("%d %f", 38383, 32.3)
	s2 := string(38383)

	fmt.Printf("%T %v\n", s1, s1)
	fmt.Printf("%T %v\n", s2, s2)

	// STRING TO NUMBER
	var f1, err = strconv.ParseFloat("3.14", 64)
	_ = err
	fmt.Printf("%T %v\n", f1, f1)

	var ii, err1 = strconv.Atoi("67")
	_ = err1
	fmt.Printf("%T %v\n", ii, ii)

}
