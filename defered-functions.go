package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func third() {
	fmt.Println("Third")
}

func output() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func main() {
	// LIFO
	defer third()
	first()
	output()
	defer second()

	fmt.Printf("y is %d", 5.6)

	fmt.Printf("y is %f", 5)

}
