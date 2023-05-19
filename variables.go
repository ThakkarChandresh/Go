package main

import (
	"fmt"
	"math/cmplx"
	"reflect"
)

var colour = "red" // package level. Will not give any error if you don't use it

func main() {

	/*Naming Conventions for Golang Variables
	These are the following rules for naming a Golang variable:

	A name must begin with a letter, and can have any number of additional letters and numbers.
	A variable name cannot start with a number.
	A variable name cannot contain spaces.
	If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
	If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
	If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
	Variable names are case-sensitive (car, Car and CAR are three different variables).
	*/

	/*	Variables declared without an explicit initial value are given their zero value.

		The zero value is:

			0 for numeric types,
		    false for the boolean type, and
			"" (the empty string) for strings.
	*/

	colour := "blue" // short hand declaration

	//this is var blocks which can give better readability

	var (
		married      bool       = false
		intMaxRange  uint64     = 1<<64 - 1
		complexRange complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Println("Are you Married ? ", married)
	fmt.Println("Int max range", intMaxRange)
	fmt.Println("Complex example", complexRange)

	fmt.Println("Colour ", colour)

	if married {
		fmt.Println("Married")
	} else {
		fmt.Println("Unmarried")
	}

	// using reflect package for runtime inspections
	fmt.Println(reflect.TypeOf(married))

	demo := "yash" //var demo = "y

	demo, demo2 := "mihir", "nikunj"

	fmt.Println(demo, " ", demo2)
}
