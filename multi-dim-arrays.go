package main

import "fmt"

func main() {
	multi := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12},
	}

	// traversing multi dimensional array

	for _, num := range multi {
		for _, value := range num {
			fmt.Print(value)
		}

		fmt.Println(" ")
	}
}
