package main

import (
	"fmt"
)
import "reflect"

func main() {
	var intArray [5]int
	//var stringArray []string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	//fmt.Println(reflect.ValueOf(stringArray).Kind())

	x := [5]int{1, 2, 3, 4, 5} // initialized with values
	var y = [5]int{10, 20, 30} //Partial assignment

	/*for _, value := range x {
		fmt.Println(strconv.Itoa(value))
	}*/

	fmt.Println(x)
	fmt.Println(y)

	a := [...]int{10, 20, 30, 40, 50, 60}

	fmt.Println(a)

	b := [...]int{4: 90}

	fmt.Println(b)

	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	//intArray = [5]int{10, 20, 30, 40, 50}
	/*fmt.Println(itemExists(&strArray, "Canada"))
	fmt.Println(itemExists(&strArray, "Africa"))*/
	fmt.Println(itemExists(strArray, "Africa"))

	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy", "USA"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[len(countries)-1])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

	fmt.Println("Capacity: ", cap(countries))

}

/*func itemExists(str *[5]string, value string) bool {

	println(str[1])

	return false
}*/

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	fmt.Println("arr: ", arr)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	fmt.Printf("%T\n\n", arr)

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
