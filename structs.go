package main

import "fmt"

type student struct {
	id      int
	name    string
	gender  string
	contact int64

	result struct {
		percentage float32
		marks      int
		remarks    string
	}
}

func main() {

	var student1 student

	student1.id = 1

	student1.name = "Yash"

	student1.gender = "male"

	student1.contact = 7016183012

	student1.result.marks = 480

	student1.result.percentage = 80.00

	student1.result.remarks = "Very Good"

	fmt.Println(student1)

	updateStudent(&student1)

	fmt.Println(student1)
}

func updateStudent(student1 *student) {
	student1.contact = 9377892723
}
