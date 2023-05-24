package main

import "fmt"

func Magic[S any, T any](a S, b T) {
	fmt.Println(a, b)
}

func main() {
	Magic("Harsh", 21)
	Magic(21, "Harsh")
}
