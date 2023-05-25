package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
}

// declaring 2 struct types that represent geometrical shapes: rectangle and circle

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// method that calculates circle's area
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// method that calculates rectangle's area
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) whoAmI() string {
	return "I am Rectangle!"
}

// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
}

func main() {

	//############ Basic Polymorphism ###############

	// declaring a circle and a rectangle values
	c1 := circle{radius: 5}
	r1 := rectangle{width: 3, height: 2}

	// circle and rectangle both implements the geometry interface  because they implement all methods of the interface
	// an interface is implicitly implemented in Go
	print(c1)
	print(r1)

	// we can do this both
	//circle.area(c1)
	//shape.area(c1)

	//############ Using A Variable Type Of Interface ###############

	var s shape
	fmt.Printf("%T.....%#v\n", s, s) //<nil>.....<nil>

	s = circle{radius: 10.2}
	fmt.Printf("%T.....%#v\n", s, s) //main.circle.....main.circle{radius:10.2}
	print(s)

	s = rectangle{
		width:  10,
		height: 20,
	}
	fmt.Printf("%T.....%#v\n", s, s) //main.rectangle.....main.rectangle{width:10, height:20}
	print(s)

	s1 := r1
	fmt.Printf("%T.....%#v\n", s1, s1)
	s1.area()

	//############ Type Assertion ###############

	//s.whoAmI() // this will raise an error bcz whoAmI() is not defined in the interface so somehow we have to extract the dynamic value hold by that variable

	//s = c1
	//fmt.Println(s.(rectangle).whoAmI()) //this a valid way but what if the asserted type is not the actual type hold by that variable?
	//Can give panic interface conversion: main.shape is main.circle, not main.rectangle

	//This is recommended!
	s = c1
	rect, ok := s.(rectangle)
	if ok {
		fmt.Println(rect.whoAmI())
	} else {
		fmt.Printf("Wrong Type Assertion\n")
	}

	//############ Type Switch ###############
	switch val := s.(type) {
	case circle:
		fmt.Printf("This is a type of circle\n")
		fmt.Println(val.area())
	case rectangle:
		fmt.Printf("This is a rectangle type of value")
		fmt.Printf(val.whoAmI())
		//case shape: what is the use if this particular thing???
	}
}
