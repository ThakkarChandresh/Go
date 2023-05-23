package main

import (
	"encoding/json"
	"fmt"
)

func badMethod(message string) {
	panic(message)
}

func main() {

	var a = 10

	defer func(b int) {
		fmt.Println("A: ", b)
	}(a)

	defer func() {
		fmt.Println("A: ", a)
	}()

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("Error: ", err)
	//	}
	//}()

	defer func(b *int) {
		fmt.Println("A")
		*b++
	}(&a)

	fmt.Println("foo")

	defer func(b *int) {
		fmt.Println("B")
		*b++
	}(&a)

	fmt.Println("bar")

	defer func(b *int) {
		fmt.Println("C")
		*b++
	}(&a)

	b := 10

	//here a's value is
	//copied at function call
	//parameter and so
	//when atlast fmt.Println(a)
	//is called 'a' value is 10
	//and not 11
	defer fmt.Println("println B: ", b)
	//same principle don't applied to body of defer func
	//hence it will print updated b which is increment
	defer func() {
		fmt.Println("B in defer: ", b)
	}()

	b++

	fmt.Println(json.Marshal(make([]int, 0, 0)))
	fmt.Println(json.Marshal([]int{}))
	c := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &c)
}
