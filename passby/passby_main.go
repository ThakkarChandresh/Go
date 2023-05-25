package main

import "fmt"

func main() {
	var a int
	var b, c = &a, &a

	fmt.Println(&a, &a)
	fmt.Println(b, c)
	fmt.Println(&b, &c)

	fmt.Printf("%T......%#v\n", a, a)
	fmt.Printf("%T......%#v\n", b, b)
	fmt.Printf("%T......%#v\n\n", c, c)

	sl := []string{"Nikunj"}
	fmt.Printf("%T......%#v\n", sl, sl)
	sliceModify(sl)
	fmt.Printf("%T......%#v\n", sl, sl)
}

func sliceModify(sl []string) {
	sl[0] = "Chandresh"
}
