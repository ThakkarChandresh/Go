package main

import "fmt"

type MyStruct struct {
	age int
}

func (m *MyStruct) getInt() {

	fmt.Println(m.age)
}

func main() {

	var ptrMyStruct *MyStruct = &MyStruct{age: 10}

	fmt.Println(ptrMyStruct.age)

	ptrMyStruct.getInt()

}
