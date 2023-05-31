package main

import "fmt"

func main() {
	mapper := make(map[string][]string)

	fruits := []string{"orange", "banana", "apple", "kiwi", "watermelon"} // 300 -> 100,5,5

	mapper["fruits"] = fruits // 200 -> 100,5,5

	for index := range fruits {
		if index%2 == 0 {
			fruits = fruits[:index]
		}
	}

	//tempslice := mapper["fruits"]

	fmt.Println(fmt.Sprintf("Slice Values are %v", fruits))
	fmt.Println(fmt.Sprintf("Map Values are %v", mapper["fruits"]))

	/*
		Output :

		Slice Values are [orange banana apple kiwi]
		Map Values are [orange banana apple kiwi watermelon]
	*/
}
