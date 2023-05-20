package main

import "fmt"

func addd() int {
	return 20
}

func main() {

	b := 20
	b += addd()
	switch a := b; {
	case a <= 5 && a >= 10:
		fmt.Println("a")
	case a < 10 && a >= 20:
		fmt.Println("b")
	}

}
