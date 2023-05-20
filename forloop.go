package main

import "fmt"

func main() {
	k := 1
	for ; k <= 10; k++ {
		fmt.Printf("%d ", k)
	}
	fmt.Println()

	k = 1
	for k <= 10 {
		fmt.Printf("%d ", k)
		k++
	}
	fmt.Println()

	for k := 1; ; k++ {
		fmt.Printf("%d ", k)
		if k == 10 {
			break
		}
	}
	fmt.Println()

	var names = []string{"I", "am", "the", "best"}
	for index, name := range names {
		fmt.Println(index, name)
	}

}
