package main

import (
	"fmt"
)

func main() {
	printFunc("Nikunj", "Chandresh", "Yash", "Harsh")
	intList(4, 3, 1, 3, 5, 6, 1)
}

func printFunc(s ...string) {
	fmt.Println(s)
}

func intList(list ...int) {
	for index, value := range list {
		//fmt.Printf(strconv.Itoa(index), " - ", value)
		println(index, " - ", value)
	}
}
