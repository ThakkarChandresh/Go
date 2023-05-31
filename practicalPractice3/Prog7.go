package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]struct{})
	m["v1"] = struct{}{}
	m["v2"] = struct{}{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(m map[string]struct{}, g *sync.WaitGroup) {
		m["v3"] = struct{}{}
		m["v4"] = struct{}{}
		fmt.Println("len: ", len(m))
		g.Done()
	}(m, &wg)
	wg.Wait()
	fmt.Println("LEN: ", len(m))
}
