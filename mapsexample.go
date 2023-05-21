package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var m map[string]int
	// below generates panic as map is nil
	//m["some"] = 1
	fmt.Println(m)

	//generates non nil empty map
	p := make(map[string]int)
	p["some"] = 1
	p["thing"] = 2
	p["some"] = 3 //overwritten value of some to 3

	fmt.Println("p:", p)

	fmt.Println("some's key value: ", p["some"])

	//below prints default value of int
	fmt.Println("unknown key's value: ", p["unknownkey"])
	delete(p, "unknownkey")

	fmt.Println("after deleting key which don't exist: ", p)

	//telling go to put aside space for 1000 key, value pairs
	q := make(map[string]string, 1000)
	q["mihir"] = "agrawal"

	/*
		go gets default value if key
		is not present in map, so we can't be
		sure if key was present or not
		To deal with this, go provides
		another lookup function which returns key value
		and bool value, where bool indicating
		whether key was present or not
	*/
	if val, ok := p["something"]; ok {
		fmt.Println("Key 'something' found with value: ", val)
	} else {
		fmt.Println("Key 'something' not found")
	}

	words := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	//scan words and not entire line
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	fmt.Println(words)

	//declaring struct of type pair
	type pair struct {
		key string
		val int
	}

	//declaring slice of pair type
	var mapPairs []pair
	for k, v := range words {
		mapPairs = append(mapPairs, pair{k, v})
	}
	fmt.Println("Array of pairs: ", mapPairs)

	//sort the map pairs by keys decreasing
	//func takes indexes i,j
	sort.Slice(mapPairs, func(i, j int) bool {
		return mapPairs[i].key > mapPairs[j].key
	})
	fmt.Println("Array of pairs sorted decreasing by key: ", mapPairs)

	//sort the map pairs by value
	sort.Slice(mapPairs, func(i, j int) bool {
		return mapPairs[i].val > mapPairs[j].val
	})
	fmt.Println("Array of pairs sorted decreasing by value: ", mapPairs)
}
