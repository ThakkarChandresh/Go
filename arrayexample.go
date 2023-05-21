package main

import "fmt"

func main() {

	var array = [3]int{3, 2, 1}
	//pass by value and new copy of array is made
	var array2 = array
	array2[0] = 1

	//both will be different address
	fmt.Printf("array: %p\n", &array)
	fmt.Printf("array2: %p\n", &array2)

	//array4 := []int{2, 1, 1, 3, 2}
	//array2 = array4 //syntax error

}
