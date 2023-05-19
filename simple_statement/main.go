package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// converting string_basics to int:
	i, err := strconv.Atoi("45")

	// error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)

	}

	// simple (short) statement ->  the same effect as the above code
	// i and err are variables scoped to the if statement only
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is ", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) < 2 {
		fmt.Println("At least one argument is required!")
	} else if f, err := strconv.ParseFloat(args[1], 64); err == nil {
		fmt.Printf("Your Number Is %T......%v\n", f, f)
	} else if err != nil {
		fmt.Println("Some Error Occured! ", err)
	} else {
		fmt.Print("N/A")
	}
}
