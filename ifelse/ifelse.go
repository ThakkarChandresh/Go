package main

import (
	"fmt"
	"strconv"
)

func badMethod(message string) {
	panic(message)
}

func main() {

	//notice that strconv.Atoi() is inside 'if' statement
	//then also it can be accessed in further 'else ifs' scope
	if num, _ := strconv.Atoi("102"); num < 10 {
		fmt.Println("Number is smaller than 10: ", num)
	} else if num < 30 && num > 20 {
		fmt.Println("Number is between 20 and 30: ", num)
	} else if num > 30 && num < 40 {
		fmt.Println("Number is between 30 and 40: ", num)
	} else {
		fmt.Println("Number is greater than 100")
	}

}
