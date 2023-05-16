package main

import "fmt"

func main() {
	i := 10
	j := 30

	j, i = i, j

	fmt.Println("I is", i)

	fmt.Println("J is", j)
}
