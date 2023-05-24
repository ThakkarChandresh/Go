package main

import (
	"fmt"
	"strconv"
)

func printTable(c chan string, n int) {
	for i := 1; i <= 10; i++ {
		c <- "" + strconv.Itoa(n) + "x" + strconv.Itoa(i) + "=" + strconv.Itoa(n*i)
	}
	close(c)
}

func first() {
	c := make(chan string)
	go printTable(c, 7)
	for item := range c {
		fmt.Println(item)
	}
}

func printTableLine(c chan string, n, i int) {
	c <- "" + strconv.Itoa(n) + "x" + strconv.Itoa(i) + "=" + strconv.Itoa(n*i)
}

func second() {
	c := make(chan string)
	for i := range [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		go printTableLine(c, 8, i+1)
	}
	for item := range c {
		fmt.Println(item)
	}
}

func main() {
	//first()
	second()
}
