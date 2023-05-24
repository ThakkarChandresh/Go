package main

import "fmt"

func main() {
	fmt.Print(add(10, 20))
	a := 60
	b := 40
	update(&a, &b)

	fmt.Println("new value of a: ", a, " | new value of b: ", b)

	var sum = func(a int, b int) int {
		return a + b
	}

	fmt.Println("Sum: ", sum)

	tem := 100

	println("value:-  ", func(a int, b int) int {
		return a + b + tem
	}(10, 30))

	a1 := func(a int, b int) int {
		return a + b
	}

	fmt.Printf("%d", a1(60, 80))

	partial := partialSum(10)

	println("Higher order function: ", partial(24))
	println("Higher order function2: ", partialSum(24)(16))
}

func function() {
	fmt.Println("We are inside function")
}

func add(x int, y int) int {
	total := x + y
	return total
}

func update(a *int, b *int) {
	*a = *a + 9
}

func sum(a, b int) int {
	return a + b
}

func partialSum(a int) func(int) int {
	return func(b int) int {
		return sum(a, b)
	}
}
