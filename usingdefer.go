package main

import "fmt"

func badMethod(message string) {
	panic(message)
}

func main() {

	//fmt is package for printing function

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	defer func() {
		//if err := recover(); err != nil {
		fmt.Println("Error 2: ")
		//}
	}()

	badMethod("Yash the conquerer")

	fmt.Println("yoo")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error 3: ", err)
		}
	}()

	badMethod("Yash the conquerer")

}
