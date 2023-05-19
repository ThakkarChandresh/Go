package main

import "fmt"

func main() {
	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1, 2}, []int{1, 2, 3}
	_, _ = a, b
	//fmt.Println(a == b) //slice can only be compared to nil

	var eq bool

	if len(a) == len(b) {
		eq = true
		for i, value := range b {
			if value != a[i] {
				eq = false
				break
			}
		}
	}

	if eq {
		fmt.Println("Both Slices Are Equal")
	} else {
		fmt.Println("Slices Are Not Equal!")
	}
}
