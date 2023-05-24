package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("len: ", len(slice), "\tcap: ", cap(slice), "\tslice: ", slice)
}

func returnArrayAndSlice() (*[10]int, []int) {
	array := [10]int{2, 4, 6, 8, 10, 12, 14}
	slice := array[3:8]
	fmt.Println("array: ", array, "\tslice: ", slice)
	fmt.Printf("Array address: %p, Slice address: %p\n", &array, &slice)
	return &array, slice
}

func sliceCreationFromArray() {
	array, slice := returnArrayAndSlice()
	fmt.Println(array, slice)
}

func sliceDefaultExamples() {
	_, slice := returnArrayAndSlice()
	printSlice(slice[2:])
	printSlice(slice[:4])
	printSlice(slice[3:])
}

func sliceAppend() {
	var array *[10]int
	var slice []int
	array, slice = returnArrayAndSlice()
	fmt.Printf("Array address: %p, Slice address: %p\n", &array, &slice)
	slice[4] = 15
	fmt.Printf("Array address: %p, Slice address: %p\n", &array, &slice)

	fmt.Println("Array: ", array, "\tSlice:", slice)
	slice = append(slice, 1, 3, 5, 7)
	fmt.Printf("Array address: %p, Slice address: %p\n", &array, &slice)

	fmt.Println("Array: ", array, "\tSlice:", slice)
	slice1 := append(slice, 9, 11, 13, 15, 17)
	fmt.Printf("Array address: %p, Slice address: %p, Slice1 address: %p\n", &array, &slice, &slice1)

	fmt.Println("Array: ", array, "\tSlice: ", slice, "\tSlice1: ", slice1)
	slice[0] = 9
	slice1[1] = 11
	fmt.Printf("Array address: %p, Slice address: %p, Slice1 address: %p\n", &array, &slice, &slice1)

	fmt.Println("Array: ", array, "\tSlice: ", slice, "\tSlice1: ", slice1)

	for i, v := range slice1 {
		fmt.Println("Slice index: ", i, "\tvalue:", v)
	}

	fmt.Println("============================================")

	for i, _ := range slice1 {
		fmt.Println("Slice index: ", i)
	}
	fmt.Println("============================================")

	for _, v := range slice1 {
		fmt.Println("Slice value:", v)
	}

}

func main() {
	print("This is main ")
	//sliceCreationFromArray()
	//sliceDefaultExamples()
	sliceAppend()
}
