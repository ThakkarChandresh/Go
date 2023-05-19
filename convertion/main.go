package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x = 3   //int type
	var y = 3.2 //float type

	// x = x * y //compile error ->  mismatched types

	x = x * int(y) // converting float64 to int
	fmt.Println(x) // => 9

	y = float64(x) * y //converting int to float64
	fmt.Println(y)     // => 28.8

	x = int(float64(x) * y)
	fmt.Println(x) // => 259

	//In Go types with different names are different types.
	var a int = 5   // same size as int64 or int32 (platform specific)
	var b int64 = 2 // int and int64 are not the same type
	var c int32 = 6
	// a = int(c) //it will also not work
	_ = c

	// a = b // error: cannot use b (type int64) as type int in assignment
	a = int(b) // converting int64 to int (explicit conversion required)

	// preventing unused variable error
	_ = a

	//** CONVERTING NUMBERS TO STRINGS AND STRINGS TO NUMBERS **//

	s := string(99)            // int to rune (Unicode code point)
	fmt.Println(s)             // => 99, the ascii code for symbol c
	fmt.Println(string(34234)) // => 34234 is the unicode code point for 薺

	// we cannot convert a float to a string_basics similar to an int to a string_basics
	// s1 := string_basics(65.1) // error

	// converting float to string_basics
	var myStr = fmt.Sprintf("%f", 5.12)
	fmt.Println(myStr) // => 5.120000

	// converting int to string_basics
	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1) // => 34234

	// converting string_basics to float
	var result, err = strconv.ParseFloat("3.142", 64)
	if err == nil {
		fmt.Printf("Type: %T, Value: %v\n", result, result) // => Type: float64, Value: 3.142
	} else {
		fmt.Println("Cannot convert to float64!")
	}

	// Atoi(string_basics to int) and Itoa(int to string_basics).
	i, err := strconv.Atoi("-50")
	s = strconv.Itoa(20)
	fmt.Printf("i Type is %T, i value is %v\n", i, i) // => i Type is int, i value is -50
	fmt.Printf("s Type is %T, s value is %q\n", s, s) // => s Type is string_basics, s value is "20"
}
