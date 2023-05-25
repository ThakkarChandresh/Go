package main

import "fmt"

func main() {
	fmt.Println("Sum is :", sum(5, 40))
}

func sum(a, b int) (sum int) {
	/*defer func() {
		sum = a + b
	}()*/

	defer func() {
		a = a + 100
		//sum = a + 100
	}()
	sum = a + b
	/*if a <= 5 {
		return sum
	} else {
		return a + b
	}*/
	return a
}
