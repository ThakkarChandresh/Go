package main

import "fmt"

func main() {
	if a, b := 10, 20; a < 5 {
		fmt.Println("a is less than 5")
		fmt.Println(b)
	} else if a = 20; a == 20 {
		fmt.Println("value of a = ", a)
		a = 400
		fmt.Println(a)
	}
}
