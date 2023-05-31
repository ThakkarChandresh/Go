package main

import "fmt"

func main() {

	function2()
	fmt.Println("finished")
}

func function2() {
	if r := recover(); r != nil {
		fmt.Println("panic recovered")
	}
	fmt.Println("some work-1")
	panic("panic")
	fmt.Println("some work-2")
}
