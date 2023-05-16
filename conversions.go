package main

import (
	"fmt"
	"strconv"
)

func main() {

	// convert string to integer
	var age, _ = strconv.Atoi("21")

	var plotNo, _ = strconv.ParseInt("87", 10, 64)

	fmt.Println("String to integer ", age)

	fmt.Println("String to integer ", plotNo)

	// convert string to Float

	var height, _ = strconv.ParseFloat("5.9", 64)

	fmt.Println("String to float ", height)

	// convert string to Boolean

	married, _ := strconv.ParseBool("true")

	fmt.Println("string to Boolean", married)

	// convert Integer to string

	var salary string = strconv.FormatInt(100000, 10)

	fmt.Println("Integer to string", salary)

	// convert Boolean to string

	var educated string = strconv.FormatBool(true)

	fmt.Println("Boolean to string", educated)

	// convert Float to string

	// prec(3rd arg) = -1 to specify to set smallest number of digits

	var weight string = strconv.FormatFloat(64.8, 'E', -1, 64)

	fmt.Println("Float to string", weight)

	// convert int16 to int32

	var i int8 = 127

	fmt.Println("Integer 8 to Integer 32 ", int32(i))

	// convert int to float

	var index int = 10

	fmt.Println("Integer to FLoat ", float64(index))

}
