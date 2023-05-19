package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// minimum buffer size is set to 16 even if you give less than 16
	reader := bufio.NewReaderSize(os.Stdin, 10)

	// way to take inout even if buffer gets full
	text, prefix, _ := reader.ReadLine()
	for prefix {
		fmt.Println(string(text))
		text, prefix, _ = reader.ReadLine()
	}

	fmt.Println(string(text))
	fmt.Println(reader.Size())

}
