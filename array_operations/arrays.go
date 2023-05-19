package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	numbers[0] = 1

	fmt.Printf("%#v\n", numbers)

	// numbers[3] = 100

	//range is just a language keyword not a function
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	fmt.Println(strings.Repeat("#", 5))

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	multi := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		[4]int{},
	}
	fmt.Printf("%#v\n", multi)

}
