package main

import (
	"awesomeProject/sayinghello"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Println(sayinghello.Say(os.Args[1]))
	} else {
		fmt.Println(sayinghello.Say("user"))
	}
}
