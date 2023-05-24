package main

import "fmt"

func swap(A, B string) (string, string) {
	return B, A
}

func main() {
	b, a := swap("first", "second")
	fmt.Println(b, a)
}
