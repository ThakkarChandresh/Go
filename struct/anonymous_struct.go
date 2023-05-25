package main

import (
	"fmt"
)

type temp struct {
	name    string
	contact int
}

func main() {
	var superTemp temp
	temp := struct {
		name    string
		contact int
	}{
		name:    "Chandresh",
		contact: 83011,
	}

	fmt.Printf("%T\n%v\n\n", temp, temp)
	fmt.Printf("%T\n%v", superTemp, superTemp)
}
