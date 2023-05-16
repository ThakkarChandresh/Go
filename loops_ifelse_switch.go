package main

import (
	"fmt"
)

func main() {

	name := "yash"

	// if-else

	if x := 100; x <= 500 {

		fmt.Println("less than equal to 500")

	} else {

		fmt.Println("more than 500")

	}

	if name == "yash" {

		fmt.Println("My name is Yash")

	} else {
		fmt.Println("My name is Ayush")
	}

	// switch case

	var colour string

	fmt.Scanln("Enter any colour : ", &colour)

	switch colour {

	case "red":
		fmt.Println("Switch : red")

	case "blue":
		fmt.Println("Switch : blue")
		fallthrough

	case "green":
		fmt.Println("Switch : green")

	case "white":
		fmt.Println("Switch :white")

	default:
		fmt.Println("Switch default : black")
	}

	x := 1

	for x <= 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("-------------------------")

	k := 1

	for ; k < 5; k++ {
		fmt.Println(k)
	}

	fmt.Println("-------------------------")

	fmt.Println("Printing colours slice using for loop")

	colours := []string{"red", "green", "blue", "white"}
	prac(colours[0:1])
	for colour := range colours {
		fmt.Println(colour)
	}

	fmt.Println(colours)

	fmt.Println("-------------------------")

	fmt.Println("Infinite loop\n")

	x = 5

	for {
		fmt.Println("Hello")
		if x == 10 {
			break
		}
		x++
	}
}

func prac(colours []string) {
	fmt.Println(colours)
	fmt.Println("Length:", len(colours))

	collar := append(colours, "smart", "yash", "mihir")

	for col := range collar {
		fmt.Println("color:", collar[col])
	}
}
