package main

import "fmt"

// In go a function can return another function and can also take a function as an argument
func main() {
	l := 20
	b := 30

	fmt.Println("Let's Create New Anonymous Function That performs Sum: ")

	func(a, b int) {
		fmt.Println("Sum Is:", a+b)
	}(10, 20)

	func() {
		var area int
		l = 30
		area = l * b
		fmt.Println(area)
	}()

	fmt.Println(l)

	mulFunc := func(a, b int) (mul int) {
		//defer mul = a * b //Expression in defer must be a function call
		mul = a * b
		fmt.Println("Multiplication Is:", mul)
		return
	}

	defer mulFunc(10, 20)
	fmt.Println("Done!!!!")
}
