package main

import "fmt"

func main() {
	values := make([]interface{}, 10)

	for index := 0; index < len(values); index++ {
		//values[index] += values[index]
	}
	fmt.Println(values)
}
