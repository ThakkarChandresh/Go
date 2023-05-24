package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

type shape interface {
	area() float64
	perimeter()
}

func (c circle) area() float64 {
	fmt.Println("We are inside area...")
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() {
	fmt.Println("Circle perimeter")
}

func main() {
	c1 := circle{radius: 5}

	fmt.Println(c1)

	fmt.Println("Circle Area: ", c1.area())

	var s shape = circle{radius: 200}

	fmt.Println(s)

	var slice1 = []int{1, 2, 3}

	var empty interface{}

	empty = slice1

	fmt.Println(len(empty.([]string)))
}
