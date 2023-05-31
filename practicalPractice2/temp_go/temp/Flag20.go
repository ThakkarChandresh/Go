package main

import (
	"fmt"
	"time"
)

func main() {
	var flag bool

	go func(flag bool) {
		defer func() {
			if flag {
				fmt.Println("Flag is true")
			} else {
				fmt.Println("flag is false")
			}
		}()

		mapper := make(map[string]interface{})

		mapper["value1"] = "abc"
		mapper["value1"] = "123"
		mapper["value1"] = []string{"abc", "def"}

		if _, ok := mapper["value1"]; ok {

			fmt.Println("inside go if")
			flag = true
		}
	}(flag)

	time.Sleep(time.Second * 5)

	if flag {
		fmt.Println("flag is true")
	} else {
		fmt.Println("flag is false")
	}
}
