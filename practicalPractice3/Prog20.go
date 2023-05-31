package main

import "fmt"

type Mapper struct {
	slice []int
}

func (mapper *Mapper) exec() {
	mapper.slice = append(mapper.slice, 2, 3, 4)
}

func main() {
	obj := Mapper{}
	obj.exec()
	fmt.Printf("len: %v", len(obj.slice))
}
