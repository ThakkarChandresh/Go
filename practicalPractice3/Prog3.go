package main

import "fmt"

func main() {
	m := make(map[string][]string)
	fruits := []string{"o", "b", "a", "k", "w"}
	m["fruit"] = fruits
	for i := range fruits {
		fruits = fruits[:i]
	}
	fmt.Println("Slice values are: ", fruits, "\tMap values are: ", m["fruit"])
}
