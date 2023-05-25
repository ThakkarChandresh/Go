package main

import (
	"fmt"
	"time"
)

type Integer interface {
	increment() chandreshInt
}

type chandreshInt int

func (c *chandreshInt) increment() chandreshInt {
	*c += 1
	return *c
}

func main() {
	cInt := chandreshInt(1)

	cInt.increment()
	fmt.Println(cInt)

	newInt := Integer.increment(&cInt)
	fmt.Println(newInt)
	fmt.Println(cInt)

	time.Sleep(10_000 * time.Second)
}
