package main

import (
	"fmt"
)

func main() {
	var student = make(map[int]string, 3)

	fmt.Println("Empty Map : ", student)

	// insertion in map

	student[1] = "mihir"
	student[3] = "deven"
	student[0] = "yash"
	student[2] = "dhaval"

	fmt.Println("\nMap data after insertion : ", student)

	fmt.Println("Map length after insertion", len(student))

	// updation

	student[3] = "chandresh"

	fmt.Println("\nAfter updating name of student having id = 3 :", student)

	//deletion

	delete(student, 3)

	fmt.Println("\nAfter deleting key 3 ", student)

	fmt.Println("Length after deleting key 3 ", student)

	fmt.Println("Traversing a student map")

	// traversing map in student

	for id, name := range student {
		fmt.Println("Id:", id, "=>", "Age:", name)
	}

	// sorting map using keys

	keys := make([]int, len(student))

	for key := range student {
		keys = append(keys, key)
	}

	//sort.Ints(keys)

	for key := range keys {
		fmt.Println(student[key])
	}

	// truncating a map

	// 1st way
	for id, _ := range student {
		delete(student, id)
	}

	//2nd way

	student = make(map[int]string)
}
