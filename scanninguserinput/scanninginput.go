package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var str, str2 string
	//Scanf requires you to write entire sentence like for below if we write:-
	//Name: Mihir then "Mihir" will be extracted from entire string
	n, err := fmt.Scanf("Name: %s", &str)
	fmt.Println(n, "<>", err)
	fmt.Println("1) Your name: ", str)

	//Scanln will read line until the new line and will take values seperated by spaces
	//so if "A B C" is given as input then we first 'str'
	n, err = fmt.Scanln(&str)
	fmt.Println(n, "<>", err)
	fmt.Println("2) Your name: ", str, str2)

	n, err = fmt.Scan(&str, &str2, &str)
	fmt.Println(n, "<>", err)
	fmt.Println("3) Your name: ", str, str2)

	r, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	fmt.Println("Your statement: ", string(r))

}
