package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// Atoi()
	var num1, _ = strconv.Atoi("1234")
	var num2, _ = strconv.Atoi("9876")
	fmt.Println(num1 + num2)

	fmt.Println("--------------------------------------------")

	// ParseInt()
	strVar := "102"
	intVar, err := strconv.ParseInt(strVar, 2, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
	strVar = "fa"
	intVar, err = strconv.ParseInt(strVar, 16, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	fmt.Println("--------------------------------------------")

	//ParseFloat
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))
	s1 := "-3.141"
	f1, err := strconv.ParseFloat(s1, 8)
	fmt.Println(f1, err, reflect.TypeOf(f1))
	s2 := "A-3141X"
	f2, err := strconv.ParseFloat(s2, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f2, err, reflect.TypeOf(f2))
	}

	fmt.Println("--------------------------------------------")

	// ParseBool()
	bools1 := "true"
	b1, _ := strconv.ParseBool(bools1)
	fmt.Printf("%T, %v\n", b1, b1)

	bools2 := "t"
	b2, _ := strconv.ParseBool(bools2)
	fmt.Printf("%T, %v\n", b2, b2)

	bools3 := "0"
	b3, _ := strconv.ParseBool(bools3)
	fmt.Printf("%T, %v\n", b3, b3)

	bools4 := "F"
	b4, _ := strconv.ParseBool(bools4)
	fmt.Printf("%T, %v\n", b4, b4)

	fmt.Println("--------------------------------------------")

	// FormatBool(), Sprintf()

	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var boolstring string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(boolstring))

	fmt.Println("--------------------------------------------")

	// FormatFloat()

	var floatVar float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(floatVar))
	fmt.Println(floatVar)

	var floatToString string = strconv.FormatFloat(floatVar, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(floatToString))
	fmt.Println(floatToString)

	fmt.Println("--------------------------------------------")

	// FormatFloat()

	var intVar2 int64 = 125
	fmt.Println(reflect.TypeOf(intVar2))
	fmt.Println(intVar2)

	var intToString string = strconv.FormatInt(intVar2, 10)
	fmt.Println(reflect.TypeOf(intToString))

	fmt.Println("Base 10 value of s:", intToString)

	intToString = strconv.FormatInt(intVar2, 8)
	fmt.Println("Base 8 value of s:", intToString)

	intToString = strconv.FormatInt(intVar2, 16)
	fmt.Println("Base 16 value of s:", intToString)

	intToString = strconv.FormatInt(intVar2, 32)
	fmt.Println("Base 32 value of s:", intToString)

}
