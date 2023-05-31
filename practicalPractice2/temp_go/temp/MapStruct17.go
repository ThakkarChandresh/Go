package main

import "fmt"

func main() {
	mapper := make(map[int]struct{})

	mapper[0] = struct{}{}
	mapper[1] = struct{}{}
	mapper[2] = struct{}{}
	mapper[3] = struct{}{}
	mapper[4] = struct{}{}
	mapper[5] = struct{}{}

	for key := range mapper {
		mapper[key+1] = struct{}{}
		delete(mapper, key)
	}
	fmt.Println("finished")
}
