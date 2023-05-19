package main

import "fmt"

func main() {
	var colour string

	fmt.Scanln(&colour)

	switch colour {

	case "red":
		fmt.Println("Switch : red")

	case "blue":
		fmt.Println("Switch : blue")

	case "green":
		fmt.Println("Switch : green")

	case "white":
		fmt.Println("Switch :white")
	default:
		fmt.Println("Switch default : black")
	}

}
