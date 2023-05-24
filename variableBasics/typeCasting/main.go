package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var a = "100"

	var b, err = strconv.Atoi(a)

	_ = b

	//fmt.Print(b, " Error ", err)

	intVar, err := strconv.ParseInt(a, 2, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(a, 16, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(a, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(a, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	s := "64.8" //"30000.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))

}
