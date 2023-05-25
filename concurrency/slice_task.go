package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	var numbers = make([]int, 30)

	var wg sync.WaitGroup

	wg.Add(2)

	// ######### Two Goroutines Reading #########
	/*for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}
	go func() {
		defer wg.Done()
		for _, value := range numbers {
			fmt.Println(value)
		}
	}()

	go func() {
		defer wg.Done()
		for _, value := range numbers {
			fmt.Println(value)
		}
	}()*/

	// ######### First Read Second Writes #########
	/*go func() {
		defer wg.Done()
		//for _, value := range numbers {
		//	fmt.Println(value)
		//}
		for i := 0; i < len(numbers); i++ {
			fmt.Println(numbers[i])
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < len(numbers); i++ {
			numbers[i] = i
		}
	}()*/

	// ######### First Writes Second Reads #########
	go func() {
		defer wg.Done()
		for i := 0; i < len(numbers); i++ {
			numbers[i] = i
		}
	}()

	go func() {
		file, _ := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		defer file.Close()
		defer wg.Done()

		//for _, value := range numbers {
		//	fmt.Println(value)
		//}
		for i := 0; i < len(numbers); i++ {
			file.WriteString(fmt.Sprintf("%d\n", numbers[i]))
		}
	}()

	// ######### Both Writes #########
	/*go func() {
		defer wg.Done()
		for i := 0; i < len(numbers); i++ {
			numbers[i] = i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < len(numbers); i++ {
			numbers[i] = len(numbers) - i
		}
	}()

	for _, value := range numbers {
		fmt.Println(value)
	}*/
	wg.Wait()
}
