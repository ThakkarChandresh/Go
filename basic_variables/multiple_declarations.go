package main

import "fmt"

func main() {
	car, cost := "Audi", 50000

	_ = car
	_ = cost

	car, year := "BMW", 12

	fmt.Println(car, year)

	var opened = true

	opened, file := false, "a.txt"

	_, _ = opened, file

	var (
		salary    int
		firstname string
		gender    bool
	)

	println(salary, firstname, gender)

	var a, b, c int

	fmt.Println(a, b, c)

	sum := 10 + 5.6
	fmt.Println(sum)

}
