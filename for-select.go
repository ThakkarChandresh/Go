package main

import "fmt"

func main() {

	var input = make(chan int)

	var output = make(chan int)

	var done = make(chan bool)

	go func() {
		for {
			select {

			case val := <-input:
				output <- val + 100

			case signal := <-done:

				fmt.Println(signal)
				close(output)
				break
			}
		}
	}()

	for i := 0; i < 5; i++ {
		input <- i
	}

	close(input)

	done <- true

	for {
		data, ok := <-output

		if !ok {

			fmt.Println("Channel closed")

			break

		} else {
			fmt.Println(data)
		}
	}

}
