package main

import (
	"fmt"
)

//var wg = sync.WaitGroup{}

func goThrough(s []int, c chan int) {
	fmt.Println("going")
	for index, value := range s {
		fmt.Println(index, value)
	}
	c <- 9
	//wg.Done()
}

func updateThrough(s []int, c chan int) {
	fmt.Println("updating")
	for index, value := range s {
		s[index] = value + 1
	}
	c <- 99
	//wg.Done()
}

func main() {
	//wg.Add(2)
	c := make(chan int)
	s := []int{20, 40, 30, 10, 50, 70, 60, 90, 80}
	go goThrough(s, c)
	fmt.Println(<-c)
	go updateThrough(s, c)
	fmt.Println(<-c)
	//time.Sleep(time.Second * 2)
	//wg.Wait()
}
