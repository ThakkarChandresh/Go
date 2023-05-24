package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "window":
		fmt.Println("Window machine")
	case "linux":
		fmt.Println("Linux machine")
	default:
		fmt.Println("Unrecognized machine")
	}
}
