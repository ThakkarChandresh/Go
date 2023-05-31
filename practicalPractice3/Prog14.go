package main

import "sync"

func main() {
	g := sync.WaitGroup{}
	g.Add(1)
	go func(g *sync.WaitGroup) {
		g.Done()
		g.Wait()
		g.Done()
	}(&g)
	g.Wait()
}
