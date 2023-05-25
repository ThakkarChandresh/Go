package main

import "fmt"

type Project struct {
	name, clientBase string
}
type Employee struct {
	name             string
	salary           int
	currentlyWorking Project
}

func main() {
	var emp1 = new(Employee)

	fmt.Printf("%T\n\n", emp1)
	emp1.name = "Chandrsh"
	emp1.salary = 10_00_000
	emp1.currentlyWorking.name = "Low-Code"
	emp1.currentlyWorking.clientBase = "US"

	fmt.Printf("%+v", *emp1)

	//type fullName struct {
	//} can be used inside this function only
}
