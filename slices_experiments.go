package main

import (
	"fmt"
	"reflect"
)

// function overloading not possible in go
//func add(int a,int b){
//
//}

func add() (b int, a int) {
	a = 1
	b = 10
	c := 100
	return c, c
}

func main() {

	a, b := add()
	fmt.Println("Add: ", a, b)

	var slice1 = []int{1, 2}
	slice2 := slice1

	//both are different slice descriptors
	fmt.Printf("Slice1: %p\n", &slice1)
	fmt.Printf("Slice2: %p\n", &slice2)

	//but have same array pointed underneath
	fmt.Printf("Slice1: %p\n", &slice1[0])
	fmt.Printf("Slice2: %p\n", &slice2[0])

	var emptySlice []string
	for i := 0; i < 5; i++ {
		emptySlice = append(emptySlice, "A")
		fmt.Print(cap(emptySlice), " ")
	}
	fmt.Println()

	var slice = make([]string, 0, 5) //[-->A]
	for i := 0; i < 5; i++ {
		slice = append(slice, "A")
		fmt.Println(len(slice))
	}
	fmt.Println("old size: ", cap(slice))
	fmt.Println(reflect.TypeOf(slice))

	//appending 6th element
	slice = append(slice, "D")
	fmt.Println("new slice: ", cap(slice))

	slice = append(slice, "D")
	fmt.Println("new slice: ", cap(slice))
	//slice: pointer_start , 2, cap [-->A,B,C,D,-]
	//slice1: pointer_start, len+1, cap
	fmt.Println(slice)

	var arr2 [5]int
	for i := 0; i < 3; i++ {
		arr2[i] = i
	}
	//0,1,2,-,-
	fmt.Println("Cap: ", cap(arr2))
	fmt.Printf("%p\n", &arr2)

	//0,1,2
	arr3 := arr2[:3]

	//0,1,2,3,-
	arr3 = append(arr3, 3)
	arr3 = append(arr3, 4)
	arr3 = append(arr3, 5)

	for i := 0; i < 5; i++ {
		fmt.Printf("%p\n ", &arr2[i])
	}
	fmt.Println()
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("%p\n", &arr3[i])
	}
	fmt.Println()

	fmt.Println("Cap: ", cap(arr3))
	fmt.Println("arr3: ", reflect.TypeOf(arr3))

}
