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
