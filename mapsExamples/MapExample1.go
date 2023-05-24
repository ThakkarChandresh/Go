package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["abc"] = "123"
	map1["def"] = "456"
	fmt.Println(map1)
	fmt.Println(map1["def"])
}
