package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{17, 12, 18, 9, 24, 42, 64, 12, 68, 82, 57, 32, 9, 2, 12, 82, 52, 34, 92, 36}

	c := make(chan int)
	for i := 0; i < len(a); i = i + 5 {
		go sum(a[i:i+5], c)
	}
	output := make([]int, 5)
	for i := 0; i < 4; i++ {
		output[i] = <-c
	}
	close(c)

	fmt.Println(output)
}
