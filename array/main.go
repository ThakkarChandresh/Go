package main

import "fmt"

func main() {
	var numbers [5]int

	fmt.Printf("%T\n%v\n%#v\n", numbers, numbers, numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [4]float64{-10, 1., .3, 6.}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"ch", "an", "dr", "es"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"Chan", "dresh"}
	fmt.Printf("%#v\n", a4)

	//ellipsis operator
	a5 := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Printf("%#v\n", a5)
	fmt.Println(len(a5))

	a6 := [...]string{
		"hello",
		"world",
		"my",
		"name", //this is compulsory
	}
	fmt.Printf("%#v\n", a6)
}
