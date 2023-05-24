package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//partial assignment

	colours := [5]string{"red", "green", "blue"}

	fmt.Println(colours)

	colours[3] = "black"

	colours[4] = "white"

	fmt.Println(colours)

	name := [5]string{1: "yash", 4: "mihir"}

	fmt.Println(name)

	arr := [5]int{10, 20, 30, 40, 50}

	passByValue(arr)

	fmt.Println("After pass by value ")

	for _, num := range arr {
		fmt.Println(num)
	}

	passByReference(&arr)

	fmt.Println("After pass by reference ")

	for _, num := range arr {
		fmt.Println(num)
	}

	//copy array

	//copying by value

	array2 := arr

	arr[4] = 100

	fmt.Println("After copying by value array 2 = ", array2)

	//copy by reference

	array3 := &arr

	arr[4] = 200

	fmt.Println("After copying by value array 3 = ", array3)

	//floatArray := []float64{20.5, 10, 60.6, 8, 100.0000}

	// unsafe pointers is array

	test(unsafe.Pointer(&arr[0]))

	numbers1 := [5]int{1, 2, 3, 4, 5}

	numbers2 := [5]int{1, 2, 3, 4, 5}

	compareArrays(numbers1, numbers2)

}

func passByValue(array [5]int) {
	array[0] = 70
	array[1] = 80
}

func passByReference(array *[5]int) {
	array[0] = 70
	array[1] = 80
}

func test(ptr unsafe.Pointer) {
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(ptr) + 16)))
}

func compareArrays(array1 [5]int, array2 [5]int) {

	value1 := reflect.ValueOf(array1)

	value2 := reflect.ValueOf(array2)

	if value1.Len() != value2.Len() {
		fmt.Println("Both types are not equal")
	} else {

		if array1 == array2 {
			fmt.Println("Both arrays are equal")
		} else {
			fmt.Println("Both arrays are not equal")
		}
	}

}
