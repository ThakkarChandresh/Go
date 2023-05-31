package main

import (
	"fmt"
	"time"
)

func main() {
	size := 100000000
	values1 := make([]int32, size)
	start1 := time.Now().UnixMilli()
	for i := 0; i < size; i++ {
		values1[i] = int32(i)
	}
	end1 := time.Now().UnixMilli()
	t1 := end1 - start1
	var values2 []int32
	start2 := time.Now().UnixMilli()
	for i := 0; i < size; i++ {
		values2 = append(values2, int32(i))
	}
	end2 := time.Now().UnixMilli()
	t2 := end2 - start2
	fmt.Println(t1, t2)
}
