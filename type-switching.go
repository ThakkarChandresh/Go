package main

import "fmt"

func main() {

	var object interface{} = 10

	switch tp := object.(type) {

	// int64 is not equal to int

	case int8, int16, int32, int64, int:
		fmt.Println("Int type : ", tp)

	case string:
		fmt.Println("String type : ", tp)

	case float64:
		fmt.Println("Float 64 type : ", tp)

	default:
		fmt.Println("Unknown type")
	}
}
