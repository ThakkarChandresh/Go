package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a space-separated string"
	str = strings.Replace(str, " ", ",", -1)
	fmt.Println(str)
}
