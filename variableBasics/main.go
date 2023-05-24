package main

import "fmt"

func main() {
	var age int = 30

	fmt.Println("Age is: ", age)

	var name = "Nikunj"

	_ = name

	//name := "Nikunj Patel"

	//fmt.Println("Name is: ", name)

	var x int
	var y float32
	var nameVal string
	var done bool

	fmt.Println(fmt.Println("Deven"))

	println("value of x: ", x)
	println("value of y: ", y)
	println("value of name: ", nameVal)
	println("value of done: ", done)

	fmt.Printf("y is %f", 5.0)

	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)

}
