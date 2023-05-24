package main

import "fmt"

func main() {

	var object interface{} = "Yash"

	objectString := object.(string)

	fmt.Println(objectString)

	object = map[int]string{}

	objectMap := object.(map[int]string)

	objectMap[0] = "yash"

	objectMap[1] = "mihir"

	objectMap[2] = "dhaval"

	fmt.Println(objectMap)

}
