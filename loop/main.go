package main

import "fmt"

func main() {
	i := 8
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 10

	for j < 20 {
		fmt.Println()
	}

	sum := 0
	for {
		sum++
	}
	fmt.Println(sum) //unreachable statement
}
