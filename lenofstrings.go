package main

import "fmt"

func main() {

	//len prints the number of bytes variable value takes

	//each character is ascii
	var name string = "mihir"
	fmt.Println("Length: ", len(name)) //print '5'

	//string where in character takes more than byte
	//here visually string is of 5 length but underneath
	//each character takes 2 bytes and hence go prints
	//10 since it always prints the actual memory taken
	//by string in bytes
	name = "èèèçÿ"
	fmt.Println("Length: ", len(name)) //prints '10'

	//how to get actual number of characters
	//here casting to array of 'rune' datatype will help
	//prints 5
	fmt.Println("Actual characters no: ", len([]rune(name)))

}
