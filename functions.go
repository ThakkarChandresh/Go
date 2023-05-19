package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func data(a *int, b *int) (x, y int) {
	temp := *a

	*a = *b

	*b = temp

	fmt.Println(a)

	fmt.Println(b)

	x = *a
	y = *b

	return
}

func main() {

	var x int = 10

	var y int = 20

	fmt.Println("Addition of", x, " + ", y, " is ", add(x, y))

	fmt.Println("After Swap : ")

	fmt.Println(data(&x, &y))

	// anonymous function

	length := 20
	breadth := 40

	area := func(length, breadth int) int {
		return length * breadth
	}(length, breadth)

	fmt.Println(area)

	func() {
		var area int
		area = length * length
		fmt.Println(area)
	}()

}
