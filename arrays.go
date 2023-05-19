package main

import (
	"fmt"
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

	test(unsafe.Pointer(&arr[0]))

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
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(ptr) + 8)))
}
