package main

import "fmt"

func main() {
	s := map[string]any{}

	fmt.Println(s["Chandresh"]) //<nil>
	s["Chandresh"] = 40

	fmt.Println(s["Chandresh"]) //40
	a := s["Chandresh"]

	a = a.(int) + 10 //we must do type assertion inorder to perform any int specific operations or methods

	fmt.Printf("%T......%#v\n", s["Chandresh"], s["Chandresh"])
	fmt.Printf("%T......%#v\n", a, a)
}
