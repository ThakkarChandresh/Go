package main

import (
	"awesomeProject/checkingpackagevisibility"
	"fmt"
	"reflect"
)

func main() {
	//declaration and initialization on same line
	var country string = "India"
	fmt.Println(reflect.TypeOf(country))

	//declaration with type inference
	var x = 10.0
	fmt.Println(reflect.TypeOf(x))

	//short declaration
	marks := 40
	var grace float64 = 4
	fmt.Println(float64(marks)*(1+grace/100), " ", reflect.TypeOf(marks))

	fmt.Println("Price after applying discount:", checkingpackagevisibility.PublicGetDiscount(100))
	fmt.Println(age)
}
