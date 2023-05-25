package main

import "fmt"

type Person struct {
	name   string
	age    int
	colors []string
}

func main() {

	// ############# Exercise 1 #################
	me := Person{"Chandresh", 10, []string{"Red", "Blue"}}

	fmt.Printf("%+v\n\n", me)

	you := Person{"You", 0, []string{"You Know"}}

	fmt.Printf("%+v\n\n", you)

	// ############# Exercise 3 #################
	for i, v := range me.colors {
		fmt.Println(i, v)
	}
}
