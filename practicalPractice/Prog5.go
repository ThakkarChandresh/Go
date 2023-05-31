package main

import (
	"fmt"
	"time"
)

func main() {
	var flag bool
	go func(flag *bool) {
		defer func() {
			if *flag {
				fmt.Println("Flag is true")
			} else {
				fmt.Println("Flag is false")
			}
		}()
		mapper := make(map[string]interface{})
		mapper["v1"] = "abc"
		mapper["v2"] = 123
		mapper["v3"] = []string{"abc", "def"}
		if _, ok := mapper["v1"]; ok {
			//flag = true
			*flag = true
		}
	}(&flag)

	time.Sleep(time.Second * 3)

	fmt.Println(flag)

	if flag {
		fmt.Println("Flag is true")
	} else {
		fmt.Println("Flag is false")
	}
}
