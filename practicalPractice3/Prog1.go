package main

import "fmt"

func main() {
	defer func() { fmt.Println("f1") }()
	defer func() { fmt.Println("f2") }()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("f6")
		}
	}()
	defer func() { fmt.Println("f3"); panic("f4") }()
	defer func() { fmt.Println("f5") }()
}
