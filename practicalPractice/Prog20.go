package main

import "fmt"

type Mapper struct {
	slice []int
}

func (mapper *Mapper) test() {
	mapper.slice = append(mapper.slice, 2, 3, 4, 5, 6, 7, 8)
}

func main() {
	obj := Mapper{}
	obj.test()
	fmt.Println(len(obj.slice))
}
