package main

import "fmt"

func main() {
	// var i1 int8 = -129
	var i1 int8 = -128
	fmt.Printf("%T %v\n", i1, i1)

	// var i2 uint16 = 65536
	var i2 uint16 = 65535
	fmt.Printf("%T %v\n", i2, i2)

	var f1, f2 = .2, 4.
	fmt.Printf("%T %v , %T %v\n", f1, f1, f2, f2)

	var c1 int32 = 'F'
	var c2 rune = 'D'
	fmt.Printf("%T %v , %T %v\n", c1, c1, c2, c2)

	var b1 bool = true
	fmt.Printf("%T %v\n", b1, b1)

	var s1 string = "Hello GO!"
	fmt.Printf("%T %v\n", s1, s1)

	var arr = [5]int{4, 5, 6}
	fmt.Printf("%T %v\n", arr, arr)
	fmt.Printf("%T %v\n", arr[0], arr[0])

	var slice = []int{4, 5, 6}
	fmt.Printf("%T %v\n", slice, slice)
	fmt.Printf("%T %v\n", slice[0], slice[0])

	var mapp = map[string]float64{
		"USD": 232.2,
		"EUR": 12.34,
	}

	fmt.Printf("%T %v\n", mapp, mapp)
	fmt.Printf("%T %v\n", mapp["USD"], mapp["USD"])

	type Car struct {
		name  string
		price int
	}

	var innova Car

	fmt.Printf("%T %v\n", innova, innova)
	fmt.Printf("%T %v\n", innova.name, innova.price)

	var v int = 20
	ptr := &v

	fmt.Printf("%T %v\n", ptr, ptr)

	fmt.Printf("%T \n", f)
}
func f() {

}
