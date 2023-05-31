package main

import "fmt"

func main() {
	mapper := make(map[string][]string)
	fruits := []string{"orange", "banana", "apple", "kiwi", "watermelon"}
	mapper["fruit"] = fruits
	for index := range fruits {
		if index%2 == 0 {
			fruits = fruits[:index]
		}
	}
	fmt.Printf("slice values are: %v", fruits)
	fmt.Printf("map values are: %v", mapper["fruit"])
}
