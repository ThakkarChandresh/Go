package main

import (
	"fmt"
	"math"
)

type Obj2 struct {
	X, Y float64
}

func Abs(v Obj2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Obj2, f float64) { // try removing *
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Obj2{5, 12}
	Scale(&v, 10) // try removing &
	fmt.Println(Abs(v))
}
