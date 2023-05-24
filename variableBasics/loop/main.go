package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	for key := range strDict {
		fmt.Println(key)
	}

	for _, value := range strDict {
		fmt.Println(value)
	}

	for range "hello" {
		fmt.Println("chandu")
	}
}
