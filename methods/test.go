package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Hour
	fmt.Printf("%T\n", day)

	fmt.Println(day.Seconds())

	//var day1 time.Duration

	//fmt.Println(day1.Seconds())

	var n int64 = 973917839
	fmt.Println(n)
	fmt.Println(time.Duration(n))
}
