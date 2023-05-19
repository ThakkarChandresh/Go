package main

import "fmt"

func main() {
	var numbers = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	i := 0

	j := len(numbers) - 1

	for i < j {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		i++
		j--
	}

	fmt.Println(numbers)
}
