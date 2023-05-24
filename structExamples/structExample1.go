package main

import "fmt"

type Vertex struct {
	X    int
	Y    int
	Name string
	Flag bool
}

func main() {
	vertex := Vertex{10, 20, "test", false}

	fmt.Println(vertex)

	if vertex.Flag {
		fmt.Println(vertex.Name, vertex.X)
	} else {
		fmt.Println(vertex.Name, vertex.X)
	}
}
