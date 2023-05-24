package main

import "fmt"

func main() {
	//println("start: This is nikunj patel")
	//println("start: This is nikunj patel")

	//defer println("panic padi gyo")

	//panic("this is a panic")
	result := sum(10, 20)

	fmt.Printf("%T\n", result)

	fmt.Printf("result: %v\n", result)
	//println("end: This is nikunj patel")
}

/*func function() any {
	println("We are inside function")

	defer func() any {
		if err := recover(); err != nil {
			fmt.Printf("%v", err)
			return err
		}
		return ""
	}()

	panic("Panic inside function")
}*/

func sum(a, b int) int {

	return a + b

}
