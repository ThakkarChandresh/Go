package main

import "fmt"

type Mapper struct {
	slice []int
}

func (mapper Mapper) test() {
	mapper.slice = append(mapper.slice, 2, 3, 4)
}

func main() {

	obj := Mapper{}

	obj.test()

	fmt.Println(fmt.Sprintf("%v", len(obj.slice)))
}
