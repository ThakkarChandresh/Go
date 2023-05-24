package main

import (
	"fmt"
	"math"
)

type Obj struct {
	X, Y float64
}

func (v Obj) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Obj) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Obj{5, 12}
	v.Scale(10)
	fmt.Println(v.Abs())
}
