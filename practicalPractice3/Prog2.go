package main

import (
	"fmt"
)

func main() {
	var v1 int32
	var v2 int64
	v1 = int32(12)
	v2 = int64(14)
	//v3:=v1+v2
	fmt.Println(v1, v2)
	//fmt.Println(reflect.TypeOf(v3).Name())
}
