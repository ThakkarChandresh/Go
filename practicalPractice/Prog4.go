package main

import "fmt"

const (
	x = 7
	y
	z
)

const (
	a = "A"
	b
	c = "c"
	d
	e = 3
	f
)

var (
	aa = 10
	//bb
	//cc
)

func main() {

	const (
		n = 10
		m
		r = false
		s
	)

	fmt.Println(x, y, z)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(n, m, r, s)
}
