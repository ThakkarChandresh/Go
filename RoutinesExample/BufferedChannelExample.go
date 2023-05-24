package main

import (
	"fmt"
	"time"
)

func operation(c chan string, str string) {
	var retStr = ""
	for _, char := range str {
		retStr = string(char) + retStr
	}
	c <- retStr
}
func operation2(str string) string {
	var retStr = ""
	for _, char := range str {
		retStr = string(char) + retStr
	}
	return retStr
}

func main() {
	start := time.Now()
	c := make(chan string, 2)
	go operation(c, "Harsh")
	go operation(c, "Kumar")
	go operation(c, "This is go language")
	go operation(c, "How is this go routine and go channels are working")

	fmt.Println("go - ", <-c)
	fmt.Println("go - ", <-c)
	fmt.Println("go - ", <-c)
	fmt.Println("go - ", <-c)

	fmt.Println("go time elapsed: ", time.Since(start))
	start = time.Now()
	fmt.Println(operation2("Harsh"))
	fmt.Println(operation2("Kumar"))
	fmt.Println(operation2("This is go language"))
	fmt.Println(operation2("How is this go routine and go channels are working"))
	fmt.Println("time elapsed: ", time.Since(start))
	//fmt.Println("go - ", <-c)
	//fmt.Println("go - ", <-c)
	//fmt.Println("go - ", <-c)
	//fmt.Println("go - ", <-c)
}
