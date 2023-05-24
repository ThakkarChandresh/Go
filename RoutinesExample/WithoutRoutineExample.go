package main

import (
	"fmt"
	"time"
)

func say(something string) {
	for i := 0; i < 5; i = i + 1 {
		time.Sleep((3000) * time.Microsecond)
		fmt.Println(something, i)
	}
}

func main() {
	say("Hello")
	say("Bye")
}
