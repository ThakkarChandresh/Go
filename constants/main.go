package main

import "fmt"

func main() {
	const day int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234
	fmt.Printf("Duration In Seconds: %v\n", duration*secondsInHour)

	// x, y := 5, 0
	// fmt.Println(x / y)

	// const x1, y1 = 5, 0
	// fmt.Println(x1 / y1)

	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3)
}
