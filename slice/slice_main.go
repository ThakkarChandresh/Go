package main

import "fmt"

func main() {
	var cities []string
	fmt.Println(cities == nil)
	fmt.Printf("%#v\n", cities)
	fmt.Println(len(cities))

	// cities[0] = "AHD"

	numbers := []int{2, 3, 4, 5}
	fmt.Printf("%#v\n", numbers)

	// numbers[4] = 20

	nums := make([]int, 5)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Chandresh", "Thakkar"}
	name := make(names, 4)
	fmt.Printf("%#v\n", friends)
	fmt.Printf("%#v\n", name)

	for _, value := range friends {
		fmt.Println(value)
	}

	//slices with same element type can be assigned to each other

	var n []int
	n = numbers
	fmt.Printf("%#v\n", n)
}
