package main

import "fmt"

type Obj1 struct{ X, Y int }

var (
	v1 = Obj1{1, 2}
	v2 = Obj1{X: 1}
	v3 = Obj1{}
	p  = &Obj1{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
