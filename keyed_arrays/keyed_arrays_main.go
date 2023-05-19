package main

import "fmt"

func main() {
	grades := [3]int{
		2: 30,
		1: 2,
		0: 10,
	}

	fmt.Printf("%#v\n", grades)

	accounts := [3]int{2: 40}

	fmt.Printf("%#v\n", accounts)

	names := [...]string{
		5: "Chan",
		"Dresh",
		0: "Jay Swaminarayan",
	}

	fmt.Printf("%#v\n", names)
}
