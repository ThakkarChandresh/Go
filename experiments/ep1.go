package main

import (
	"fmt"
	"reflect"
)

func main() {

	data := map[int]interface{}{}

	data[0] = map[string]map[string]map[string]int{}

	data[1] = true

	data[2] = 10

	fmt.Println(reflect.TypeOf(data[1]))

	value := data[0]

	//delete(value, "1")

	fmt.Println(fmt.Sprintf(" key is %v", reflect.TypeOf(value).Key()))

	fmt.Println(fmt.Sprintf(" value is %v", reflect.TypeOf(value).Elem()))

	var recievedMap = reflect.MapOf(reflect.TypeOf(value).Key(), reflect.TypeOf(value).Elem())
	fmt.Println("Recieved Map: ", recievedMap)

	fmt.Println(reflect.ValueOf(data[0]))

}
