package main

import (
	"fmt"
	"sync"
)

func main() {

	map1 := map[string]string{}

	var wait = sync.WaitGroup{}
	wait.Add(2)

	go func() {
		map1["name"] = "chandresh"
		wait.Done()
	}()

	go func() {
		fmt.Println(map1["name"])
		wait.Done()
	}()

	wait.Wait()
}
