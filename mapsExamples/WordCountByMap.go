package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	map1 := make(map[string]int)
	for _, str := range strings.Fields(s) {
		elem, ok := map1[str]
		if ok {
			map1[str] = elem + 1
		} else {
			map1[str] = 1
		}
	}

	return map1
}

func main() {
	wc.Test(WordCount)
}
