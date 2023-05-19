package main

import (
	"fmt"
	"reflect"
)

func main() {
	numbers := [5]int{1, 2}

	numberSlice := numbers[:2]

	fmt.Println(numbers)
	_ = numberSlice

	numberSlice = append(numberSlice, 20)
	numberSlice = append(numberSlice, 21)
	numberSlice = append(numberSlice, 23)

	numberSlice = append(numberSlice, 24)
	numberSlice = append(numberSlice, 25)

	fmt.Println(numbers)

	fmt.Printf("%#v\n", numberSlice)

	newSlice := make([]string, 10, 30)
	fmt.Println(reflect.TypeOf(newSlice).Kind())

	n := []int{100, 200}
	numberSlice = append(numberSlice, n...)
	fmt.Printf("%#v\n", numberSlice)

	src := []int{10, 20, 30}
	dest := []int{1, 2, 3}
	nodes := copy(dest, src)
	_ = nodes
	fmt.Printf("%#v\n", dest)

	cars := []string{"Ford", "Hundai"}
	fmt.Println(len(cars), cap(cars))

	s1 := []int{10, 20, 30, 40, 50}
	ints := s1[1:3]
	s1[1] = 3344
	fmt.Printf("%#v\n", ints)
	fmt.Println(len(ints), cap(ints))

	intss := s1[1:5]
	fmt.Printf("%#v\n", intss)
	fmt.Println(len(intss), cap(intss))
}
