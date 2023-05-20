package main

import "fmt"

func printProduct(x, y int) {
	fmt.Println(x * y)
}

// returns multiple values
func getMultipleValue() (int, string, float64) {
	return 10, "mihir", 45.45
}

func sum(a, b int) any {
	defer func() {
		a = a + b
		fmt.Printf("%p\n", &a)
	}()
	fmt.Printf("%p\n", &a)
	return a
}

func update(a *int) {
	*a = 314 // defrencing pointer address
	return
}

func main() {

	fmt.Println("sum: ", sum(10, 20))
	printProduct(10, 20)
	fmt.Println(getMultipleValue())

	var arr [3]int
	fmt.Println("Before:", arr)

	update(&arr[0])

	fmt.Println("After :", arr)
}
