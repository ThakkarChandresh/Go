package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 'a'

	fmt.Printf("%q \n%v", a, reflect.TypeOf(a))
	//var s = new([6]int)[0:6]
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // 6 6

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // 0 6
	//|------

	// Extend its length.
	s = s[:4]
	printSlice(s) // 4 6
	//|------

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // 2 4
	//--|----

	s = s[1:]
	printSlice(s) // 2 6
	//------|

	var fruits = make([]string, 0, 10)

	fruits = append(fruits, "apple", "mango", "leechi", "watermelon")

	fmt.Println(itemExists(fruits, "mango"))

	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := []int{10, 20, 30, 40, 50}

	stringSlice := []string{"yash", "mihir", "dhaval"}
	stringSlice2 := []string{"yash", "mihir", "dhava"}

	compareSlice(slice1, slice2)

	compareSlice(stringSlice, stringSlice2)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func compareSlice(slice1 interface{}, slice2 interface{}) {

	value1 := reflect.ValueOf(slice1)

	value2 := reflect.ValueOf(slice2)

	if value1.Kind() != value2.Kind() {
		panic("Invalid type")
	}

	if value1.Len() != value2.Len() {
		fmt.Println("Both slices are not equal")
	} else {
		equal := false

		for i := 0; i < value1.Len(); i++ {
			if value1.Index(i).Equal(value2.Index(i)) {
				equal = true
			} else {
				equal = false
			}
		}

		if equal {
			fmt.Println("Both slices are equal")
		} else {
			fmt.Println("Both slices are not equal")
		}
	}
}
