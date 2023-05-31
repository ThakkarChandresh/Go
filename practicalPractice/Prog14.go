package main

import "fmt"

func main() {
	mapper := make(map[int]struct{})
	prepare(mapper)
	fmt.Println(mapper)
}

func prepare(mapper map[int]struct{}) {
	for i := 0; i <= 5; i++ {
		mapper[i] = struct{}{}
	}
}
