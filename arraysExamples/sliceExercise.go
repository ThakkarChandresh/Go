package main

import (
	"golang.org/x/tour/pic"
	//"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	retSlice := make([][]uint8, dy)
	for i, _ := range retSlice {
		retSlice[i] = make([]uint8, dx)
	}
	return retSlice
}

func main() {
	//fmt.Println(Pic(10,10));
	pic.Show(Pic)

}
