package main

import (
	"fmt"
	"strings"
)

func main() {
	var flag bool
	defer func() {
		if flag {
			fmt.Println("equal")
		} else {
			fmt.Println("not equal")
		}
	}()
	v1 := "motadata"
	if strings.EqualFold(v1, "MoTaDaTa") {
		flag = true
	} else {
		flag = false
	}
}
