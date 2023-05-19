package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	variadic("yash")
	variadic("yash", "tiwari")
	variadic("yash", "tiwari", strconv.Itoa(21))

	variadic1(1, 2, 3, 4, 5, 6)

	variadic2(1, "Ayush", 21.5, true, []string{"cpp", "java", "javascript"})

	var str, str2, str3 string

	fmt.Scan(&str, &str2, &str3)

	fmt.Println(str, str2, str3)
}

func variadic(s ...string) {
	fmt.Println(s)
}

func variadic1(i ...int) {
	fmt.Println(i)
}

func variadic2(inter ...interface{}) {
	for i := 0; i < len(inter); i++ {
		fmt.Println(reflect.TypeOf(inter[i]).Kind(), " ", inter[i])
	}
}
