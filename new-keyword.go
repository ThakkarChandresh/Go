package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	salary := new(int)

	bera := new(person)

	bera.name = "Dhaval"

	bera.age = 21

	*salary = 15000

	fmt.Println(bera.age, " ", bera.name, " ", *salary)

}
