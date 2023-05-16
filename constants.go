package main

import (
	"fmt"
	"reflect"
)

func main() {

	//normal const declaration

	const NAME string = "yash"
	const AGE = 20

	fmt.Println(NAME+" ", AGE)

	//block declaration

	const (
		BLOOD_GROUP = "A+"
		BIRTH_CITY  = "Surat"
	)
	fmt.Println(BLOOD_GROUP, " ", BIRTH_CITY)

	/*untyped constants ( means at compile time type will be const
	but upon operation the type will be assigned based on literal)*/

	const x = 30
	const y = 10.4

	fmt.Println(x, y, reflect.TypeOf(y))
}
