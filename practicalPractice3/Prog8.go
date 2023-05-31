package main

import "fmt"

func main() {
	m := make(map[int]struct{})
	m[0] = struct{}{}
	m[1] = struct{}{}
	m[2] = struct{}{}
	m[3] = struct{}{}
	m[4] = struct{}{}
	m[5] = struct{}{}
	for k := range m {
		m[k+1] = struct{}{}
		delete(m, k)
	}
	fmt.Println("finished")
}
