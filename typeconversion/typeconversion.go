package main

import (
	"fmt"
	"strconv"
)

func main() {

	//ascii to integer
	ans, err := strconv.Atoi("10000.0")
	if err != nil {
		fmt.Println("Custom Error: ", err)
	} else {
		fmt.Println("Converted Integer: ", ans)
	}

	if ans2, err2 := strconv.Atoi("3929.40"); err2 == nil {
		fmt.Println("Ans: ", ans2)
	} else {
		fmt.Println("Error: ", err2)
	}

	var val float64 = 24.30003004939239434828439429
	fmt.Println(int32(val))

	var val2 uint8 = 1
	fmt.Println(val2 << 7)
}
